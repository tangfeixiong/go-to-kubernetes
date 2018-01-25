package deploy

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"

	//"k8s.io/api/apps/v1beta1"
	"k8s.io/api/apps/v1beta2"
	apiv1 "k8s.io/api/core/v1"
	//"k8s.io/api/extensions/v1beta1"
	// "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/artifact"
)

const (
	SENTINEL_PORT_DEFAULT = 26379
	DEFAULT_NAMESPACE     = "default"
	DEFAULT_IMAGE         = "docker.io/redis:4.0-alpine"
	DEFAULT_QUORUM        = 2
)

var (
	sentinel_tpl      string = "template/sentinel-deployment.yaml"
	name_default      string = "my-redis"
	replicas_default  int    = 3
	container_default string = "sentinel"
	port_default      string = "client"
)

type SentinelRecipient struct {
	Name, Namespace          string
	ProvisioningName         string
	Image, PortName          string
	Replications, ClientPort int
	Quorum                   uint8
}

func NewSentinelRecipient(name, namespace, resourceProvision string, quorum int) *SentinelRecipient {
	var replications int
	if quorum%2 > 0 {
		replications = (quorum - 1) * 2
	} else {
		replications = (quorum-1)*2 + 1
	}

	recipe := &SentinelRecipient{
		Name:             name,
		Namespace:        namespace,
		ProvisioningName: resourceProvision,
		Replications:     replications,
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
	//	if portname == "" {
	//		recipe.PortName = port_default
	//	}
	recipe.PortName = port_default
	//	if sentinelport <= 1024 || sentinelport >= 65535 {
	//		recipe.ClientPort = SENTINEL_PORT_DEFAULT
	//	}
	recipe.ClientPort = SENTINEL_PORT_DEFAULT

	if recipe.Replications < 3 {
		recipe.Replications = replicas_default
	}

	return recipe
}

func (recipe *SentinelRecipient) parseTpl() (*bytes.Buffer, error) {
	var text string

	data, err := artifact.Asset(sentinel_tpl)
	if err != nil {
		glog.Errorf("Cloud not find spec binding, error: %s", err.Error())
		return nil, err
	} else {
		text = string(data)
	}

	te := template.Must(template.New("deploy").Parse(text))
	b := &bytes.Buffer{}

	err = te.Execute(b, recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse spec binding: %v\n", err)
	}

	glog.V(5).Infoln("Deployment artifact for Sentinel:", b.String())
	return b, nil
}

func (recipe *SentinelRecipient) GenerateArtifact() (*v1beta2.Deployment, error) {
	b, err := recipe.parseTpl()
	if err != nil {
		return nil, err
	}

	result := new(v1beta2.Deployment)

	err = yaml.Unmarshal(b.Bytes(), &result)
	if err != nil {
		glog.Errorf("Could not decode Deployment artifact: %v", err)
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

	return result, nil
}

func findContainerPort(d *v1beta2.Deployment, name string) (bool, *apiv1.Container, *apiv1.ContainerPort) {
	var c *apiv1.Container = nil
	var p *apiv1.ContainerPort

	for i := 0; i < len(d.Spec.Template.Spec.Containers); i++ {
		if d.Spec.Template.Spec.Containers[i].Name == container_default {
			c = &d.Spec.Template.Spec.Containers[i]
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
