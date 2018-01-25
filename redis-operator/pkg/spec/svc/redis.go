package svc

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"

	apiv1 "k8s.io/api/core/v1"
	// "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/artifact"
)

const (
	REDIS_PORT_DEFAULT = 6379
)

var (
	redis_tpl      string = "template/redis-service.yaml"
	client_default string = "client"
	sync_default   string = "gossip"
)

type RedisRecipient struct {
	Name, Namespace, ResourceName string
	PortName, SyncName            string
	ServicePort                   int
}

func NewRedisRecipient(name, namespace, resourceName, portName, syncPortName string, clientPort int) *RedisRecipient {
	recipe := &RedisRecipient{name, namespace, resourceName, portName, syncPortName, clientPort}

	if name == "" {
		recipe.Name = strings.Join([]string{name_default, rand.String(8)}, "-")
	}
	if namespace == "" {
		recipe.Namespace = DEFAULT_NAMESPACE
	}
	if resourceName == "" {
		recipe.ResourceName = name_default
	}
	if portName == "" {
		recipe.PortName = client_default
	}
	if syncPortName == "" {
		recipe.PortName = sync_default
	}
	if clientPort <= 1024 || clientPort >= 65535 {
		recipe.ServicePort = REDIS_PORT_DEFAULT
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

	te := template.Must(template.New("svc").Parse(text))
	b := &bytes.Buffer{}

	err = te.Execute(b, recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse spec binding: %v\n", err)
	}

	glog.V(5).Infoln("Service artifact for Redis:", b.String())
	return b, nil
}

func (recipe *RedisRecipient) GenerateArtifact() (*apiv1.Service, error) {
	b, err := recipe.parseTpl()
	if err != nil {
		return nil, err
	}

	result := new(apiv1.Service)

	err = yaml.Unmarshal(b.Bytes(), &result)
	if err != nil {
		glog.Errorf("Could not decode service artifact: %v", err)
		return nil, fmt.Errorf("Failed to perform deserializtion: %v\n", err)
	}

	result.Namespace = recipe.Namespace
	if ok, sp := findServicePort(result, client_default); ok {
		sp.Name = recipe.PortName
		sp.Port = int32(recipe.ServicePort)
	} else {
		sp := &apiv1.ServicePort{
			Name:     recipe.PortName,
			Port:     int32(recipe.ServicePort),
			Protocol: apiv1.ProtocolTCP,
			/*TargetPort:,
			  NodePort:*/
		}
		result.Spec.Ports = []apiv1.ServicePort{*sp}
	}

	if ok, sp := findServicePort(result, sync_default); ok {
		sp.Name = recipe.SyncName
		sp.Port = int32(recipe.ServicePort + 10000)
	} else {
		sp := &apiv1.ServicePort{
			Name:     recipe.SyncName,
			Port:     int32(recipe.ServicePort + 10000),
			Protocol: apiv1.ProtocolTCP,
			/*TargetPort:,
			  NodePort:*/
		}
		result.Spec.Ports = []apiv1.ServicePort{*sp}
	}

	return result, nil
}
