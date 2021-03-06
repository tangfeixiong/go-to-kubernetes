package sts

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"

	"k8s.io/api/apps/v1beta2"
	apiv1 "k8s.io/api/core/v1"
	// "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/artifact"
)

const (
	DEFAULT_NAMESPACE  = "default"
	DEFAULT_IMAGE      = "docker.io/redis:4.0-alpine"
	REDIS_PORT_DEFAULT = 6379
)

var (
	redis_tpl         string = "template/redis-statefulset.yaml"
	name_default      string = "my-redis"
	replicas_default  int    = 3
	container_default string = "redis"
	port_default      string = "client"
	sync_default      string = "gossip"
)

type RedisRecipient struct {
	Name, Namespace               string
	ProvisioningName, ServiceName string
	Image, PortName, SyncName     string
	Replications, ClientPort      int
}

func NewRedisRecipient(name, namespace, resourceProvision, serviceName string, slaves int) *RedisRecipient {
	recipe := &RedisRecipient{
		Name:             name,
		Namespace:        namespace,
		ProvisioningName: resourceProvision,
		ServiceName:      serviceName,
		Replications:     slaves + 1,
	}

	if name == "" {
		recipe.Name = strings.Join([]string{name_default, rand.String(8)}, "-")
	}
	if namespace == "" {
		recipe.Namespace = DEFAULT_NAMESPACE
	}
	if resourceProvision == "" {
		recipe.ProvisioningName = name_default
	}
	if serviceName == "" {
		recipe.ServiceName = recipe.ProvisioningName
	}
	//	if portname == "" {
	//		recipe.PortName = port_default
	//	}
	recipe.PortName = port_default
	//	if reidsport <= 1024 || reidsport >= 65535 {
	//		recipe.ClientPort = REDIS_PORT_DEFAULT
	//	}
	recipe.ClientPort = REDIS_PORT_DEFAULT
	recipe.SyncName = sync_default

	if recipe.Replications < 2 {
		recipe.Replications = replicas_default
	}

	return recipe
}

func (recipe *RedisRecipient) parseTpl() (*bytes.Buffer, error) {
	var text string

	data, err := artifact.Asset(redis_tpl)
	if err != nil {
		glog.Errorf("Cloud not find spec binding, error: %s", err.Error())
		return nil, err
	} else {
		text = string(data)
	}

	te := template.Must(template.New("sts").Parse(text))
	b := &bytes.Buffer{}

	err = te.Execute(b, recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse spec binding: %v\n", err)
	}

	glog.V(5).Infoln("StatefulSet artifact for Redis:", b.String())
	return b, nil
}

func (recipe *RedisRecipient) GenerateArtifact() (*v1beta2.StatefulSet, error) {
	b, err := recipe.parseTpl()
	if err != nil {
		return nil, err
	}

	result := new(v1beta2.StatefulSet)

	err = yaml.Unmarshal(b.Bytes(), &result)
	if err != nil {
		glog.Errorf("Could not decode StatefulSet artifact: %v", err)
		return nil, fmt.Errorf("Failed to perform deserializtion: %v\n", err)
	}

	result.Namespace = recipe.Namespace
	if ok, c, p := findContainerPort(result, port_default); ok {
		p.Name = recipe.PortName
		p.ContainerPort = int32(recipe.ClientPort)
	} else if c == nil {
		glog.Infoln("Could not find Container to set port name")
	} else {
		p := &apiv1.ContainerPort{
			Name:          recipe.PortName,
			ContainerPort: int32(recipe.ClientPort),
			Protocol:      apiv1.ProtocolTCP,
			/*TargetPort:,
			  NodePort:*/
		}
		// c.Ports = []apiv1.ContainerPort{*p}
		c.Ports = append(c.Ports, *p)
	}

	if ok, c, p := findContainerPort(result, sync_default); ok {
		p.Name = recipe.SyncName
		p.ContainerPort = int32(recipe.ClientPort + 10000)
	} else if c == nil {
		glog.Infoln("Could not find Container to set port name")
	} else {
		p := &apiv1.ContainerPort{
			Name:          recipe.SyncName,
			ContainerPort: int32(recipe.ClientPort + 10000),
			Protocol:      apiv1.ProtocolTCP,
			/*TargetPort:,
			  NodePort:*/
		}
		// c.Ports = []apiv1.ContainerPort{*p}
		c.Ports = append(c.Ports, *p)
	}

	return result, nil
}

func findContainerPort(s *v1beta2.StatefulSet, name string) (bool, *apiv1.Container, *apiv1.ContainerPort) {
	var c *apiv1.Container = nil
	var p *apiv1.ContainerPort

	for i := 0; i < len(s.Spec.Template.Spec.Containers); i++ {
		if s.Spec.Template.Spec.Containers[i].Name == container_default {
			c = &s.Spec.Template.Spec.Containers[i]
			for j := 0; j < len(c.Ports); j++ {
				if c.Ports[j].Name == name {
					p = &c.Ports[j]
					return true, c, p
				}
			}
		}
	}
	return false, c, nil
}
