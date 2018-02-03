package sts

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"

	"k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/spec/artifact"
)

var (
	asset_default string = "template/rabbitmq-statefulset.yaml.tpl"

	name_default     string = "demo-rabbitmq-cluster"
	replicas_default int32  = 3
	container_name   string = "rabbitmq"
	amqp_client      string = "amqp"
	mgmt_http        string = "http"
)

type RabbitmqRecipient struct {
	Name, Namespace                     string
	CustomResourceName, ServiceName     string
	Image, AmqpClientName, MgmtHttpName string
	AmqpClientPort, MgmtHttpPort        int
	Count                               int
	ClusterName                         string
}

func NewRabbitmqRecipient(name, namespace, crn, service string, count *int32) *RabbitmqRecipient {
	num := replicas_default
	if count != nil && *count > 0 {
		//num = (*count-1)*2 + 1
		num = *count
	}

	recipe := &RabbitmqRecipient{
		Name:               name,
		Namespace:          namespace,
		CustomResourceName: crn,
		ServiceName:        service,
		Count:              int(num),
	}

	if name == "" {
		recipe.Name = strings.Join([]string{name_default, rand.String(8)}, "-")
	}
	if namespace == "" {
		recipe.Namespace = corev1.NamespaceDefault
	}
	if crn == "" {
		recipe.CustomResourceName = name_default
	}
	if service == "" {
		recipe.ServiceName = recipe.CustomResourceName
	}

	//	if port == nil {
	//		recipe.ClientPort = DEFAULT_CLIENT_PORT
	//	} else if 1024 < *port && *port < 65535 {
	//		recipe.ClientPort = *port
	//	} else {
	//      ???
	//  }

	return recipe
}

func (recipe *RabbitmqRecipient) Parse() (*bytes.Buffer, error) {
	var text string

	data, err := artifact.Asset(asset_default)
	if err != nil {
		glog.Errorf("Reading asset failed: %s", err.Error())
		return nil, err
	} else {
		text = string(data)
	}

	te := template.Must(template.New("sts").Parse(text))
	b := &bytes.Buffer{}

	err = te.Execute(b, recipe)
	if err != nil {
		return nil, fmt.Errorf("Executing template failed: %v\n", err)
	}

	glog.V(5).Infoln("StatefulSet artifact YAML:", b.String())
	return b, nil
}

func (recipe *RabbitmqRecipient) Generate() (*v1beta2.StatefulSet, error) {
	b, err := recipe.Parse()
	if err != nil {
		return nil, err
	}

	result := new(v1beta2.StatefulSet)

	err = yaml.Unmarshal(b.Bytes(), &result)
	if err != nil {
		glog.Errorf("Deserializing artifact failed: %v", err)
		return nil, fmt.Errorf("Perform deserializtion failed: %v\n", err)
	}

	result.Namespace = recipe.Namespace

	//	if ok, c, p := findContainerPort(result, port_default); ok {
	//		p.Name = recipe.PortName
	//		p.ContainerPort = int32(recipe.ClientPort)
	//	} else if c == nil {
	//		glog.Infoln("Could not find Container to set port name")
	//	} else {
	//		p := &corev1.ContainerPort{
	//			Name:          recipe.PortName,
	//			ContainerPort: int32(recipe.ClientPort),
	//			Protocol:      corev1.ProtocolTCP,
	//		}
	//		c.Ports = append(c.Ports, *p)
	//	}

	//	if ok, c, p := findContainerPort(result, sync_default); ok {
	//		p.Name = recipe.SyncName
	//		p.ContainerPort = int32(recipe.ClientPort + 10000)
	//	} else if c == nil {
	//		glog.Infoln("Could not find Container to set port name")
	//	} else {
	//		p := &corev1.ContainerPort{
	//			Name:          recipe.SyncName,
	//			ContainerPort: int32(recipe.ClientPort + 10000),
	//			Protocol:      corev1.ProtocolTCP,
	//		}
	//		c.Ports = append(c.Ports, *p)
	//	}

	return result, nil
}

func findContainerPort(s *v1beta2.StatefulSet, name string) (bool, *corev1.Container, *corev1.ContainerPort) {
	var c *corev1.Container = nil
	var p *corev1.ContainerPort

	for i := 0; i < len(s.Spec.Template.Spec.Containers); i++ {
		if s.Spec.Template.Spec.Containers[i].Name == container_name {
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
