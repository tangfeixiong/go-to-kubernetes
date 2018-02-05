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

	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/spec/artifact"
)

const (
	name_example = "hadoop-hdfs-node"
)

var (
	asset_default string = "template/hdfs-node-service.yaml.gotpl"

	client_port_name string = "hdfs"
	web_port_name    string = "http"
)

type HadoopHdfsNodeServiceRecipient struct {
	Name, Namespace                 string
	CustomResourceName              string
	ClientPortName, WebPortName     string
	DataNodeClientPort, HdfsWebPort int
}

func NewHadoopHdfsNodeServiceRecipient(name, ns, crn string) *HadoopHdfsNodeServiceRecipient {
	recipe := &HadoopHdfsNodeServiceRecipient{
		Name:               name,
		Namespace:          ns,
		CustomResourceName: crn,
	}

	if name == "" {
		recipe.Name = strings.Join([]string{name_example, rand.String(8)}, "-")
	}
	if ns == "" {
		recipe.Namespace = apiv1.NamespaceDefault
	}
	if crn == "" {
		recipe.CustomResourceName = name_example
	}

	return recipe
}

func (recipe *HadoopHdfsNodeServiceRecipient) Parse() (*bytes.Buffer, error) {
	var text string

	data, err := artifact.Asset(asset_default)
	if err != nil {
		glog.Errorf("Read asset failed, error: %s", err.Error())
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

func (recipe *HadoopHdfsNodeServiceRecipient) New() (*apiv1.Service, error) {
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
