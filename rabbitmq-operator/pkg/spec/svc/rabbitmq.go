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

	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/spec/artifact"
)

var (
	asset_default string = "template/rabbitmq-service.yaml.tpl"

	image_default string = "docker.io/rabbitmq:3.7"

	name_default string = "demo-rabbitmq-cluster"

	amqp_client_name string = "amqp"
	mgmt_http_name   string = "http"
)

type RabbitmqRecipient struct {
	Name, Namespace              string
	CustomResourceName           string
	AmqpClientName, MgmtHttpName string
	AmqpClientPort, MgmtHttpPort int
}

func NewRabbitmqRecipient(name, ns, crn string) *RabbitmqRecipient {
	recipe := &RabbitmqRecipient{
		Name:               name,
		Namespace:          ns,
		CustomResourceName: crn,
	}

	if name == "" {
		recipe.Name = strings.Join([]string{name_default, rand.String(8)}, "-")
	}
	if ns == "" {
		recipe.Namespace = "default"
	}
	if crn == "" {
		recipe.CustomResourceName = name_default
	}

	return recipe
}

func (recipe *RabbitmqRecipient) Parse() (*bytes.Buffer, error) {
	var text string

	data, err := artifact.Asset(asset_default)
	if err != nil {
		glog.Errorf("Get asset failed, error: %s", err.Error())
		return nil, err
	} else {
		text = string(data)
	}

	te := template.Must(template.New("svc").Parse(text))
	b := &bytes.Buffer{}

	err = te.Execute(b, recipe)
	if err != nil {
		return nil, fmt.Errorf("Execute template failed: %v\n", err)
	}

	glog.V(5).Infoln("Parsed artifact:", b.String())
	return b, nil
}

func (recipe *RabbitmqRecipient) Generate() (*apiv1.Service, error) {
	b, err := recipe.Parse()
	if err != nil {
		return nil, err
	}

	obj := new(apiv1.Service)

	err = yaml.Unmarshal(b.Bytes(), obj)
	if err != nil {
		glog.Errorf("Deserialize artifact failed: %v", err)
		return nil, fmt.Errorf("Perform deserializtion failed: %v\n", err)
	}

	obj.Namespace = recipe.Namespace

	//	if ok, sp := findServicePort(obj, client_default); ok {
	//		sp.Name = recipe.PortName
	//		sp.Port = int32(recipe.ServicePort)
	//	} else {
	//		sp := &apiv1.ServicePort{
	//			Name:     recipe.PortName,
	//			Port:     int32(recipe.ServicePort),
	//			Protocol: apiv1.ProtocolTCP,
	//		}
	//		obj.Spec.Ports = []apiv1.ServicePort{*sp}
	//	}

	//	if ok, sp := findServicePort(obj, sync_default); ok {
	//		sp.Name = recipe.SyncName
	//		sp.Port = int32(recipe.ServicePort + 10000)
	//	} else {
	//		sp := &apiv1.ServicePort{
	//			Name:     recipe.SyncName,
	//			Port:     int32(recipe.ServicePort + 10000),
	//			Protocol: apiv1.ProtocolTCP,
	//		}
	//		obj.Spec.Ports = []apiv1.ServicePort{*sp}
	//	}

	return obj, nil
}

func findServicePort(svc *apiv1.Service, name string) (bool, *apiv1.ServicePort) {
	var sp *apiv1.ServicePort

	for i := 0; i < len(svc.Spec.Ports); i++ {
		if svc.Spec.Ports[i].Name == name {
			sp = &svc.Spec.Ports[i]
			return true, sp
		}
	}
	return false, nil
}
