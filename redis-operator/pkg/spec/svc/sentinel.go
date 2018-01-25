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
	SENTINEL_PORT_DEFAULT = 26379
)

var (
	sentinel_tpl string = "template/sentinel-service.yaml"
	port_default string = "client"
)

type SentinelRecipient struct {
	Name, Namespace        string
	ResourceName, PortName string
	ServicePort            int
}

func NewSentinelRecipient(name, namespace, resourceName, portName string, port int) *SentinelRecipient {
	recipe := &SentinelRecipient{name, namespace, resourceName, portName, port}

	if name == "" {
		recipe.Name = strings.Join([]string{name_default, rand.String(8)}, "-")
	} else if name == resourceName {
		recipe.Name = fmt.Sprintf("%s-ha", name)
	}
	if namespace == "" {
		recipe.Namespace = DEFAULT_NAMESPACE
	}
	if resourceName == "" {
		recipe.ResourceName = name_default
	}
	if portName == "" {
		recipe.PortName = port_default
	}
	if port <= 1024 || port >= 65535 {
		recipe.ServicePort = SENTINEL_PORT_DEFAULT
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

	te := template.Must(template.New("svc").Parse(text))
	b := &bytes.Buffer{}

	err = te.Execute(b, recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse spec binding: %v\n", err)
	}

	glog.V(5).Infoln("Service artifact for Sentinel:", b.String())
	return b, nil
}

func (recipe *SentinelRecipient) GenerateArtifact() (*apiv1.Service, error) {
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
	if ok, sp := findServicePort(result, port_default); ok {
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

	return result, nil
}
