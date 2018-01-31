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

	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/artifact"
)

const (
	DEFAULT_CLIENT_PORT  = 3306
	DEFAULT_TRAFFIC_PORT = 4567
	DEFAULT_IST_PORT     = 4568
	DEFAULT_SST_PORT     = 4444
)

var (
	asset_default string = "template/mariadb-service.yaml.tpl"

	name_default string = "demo-mariadb-galera"

	client_port_name string = "client"
	traffic_name_tcp string = "traffic"
	traffic_name_udp string = "trafficu"
	ist_port_name    string = "ist"
	sst_port_name    string = "sst"
)

type MariadbGaleraRecipient struct {
	Name, Namespace                                              string
	CustomResourceName                                           string
	ClientName, TrafficTCP, TrafficUDP, IstName, SstName         string
	ServicePort, ReplicaTcp, ReplicaUdp, StatePort, SnapshotPort int
}

func NewMariadbGaleraRecipient(name, ns, crn string) *MariadbGaleraRecipient {
	recipe := &MariadbGaleraRecipient{name, ns, crn, client_port_name, traffic_name_tcp, traffic_name_udp, ist_port_name, sst_port_name, DEFAULT_CLIENT_PORT, DEFAULT_TRAFFIC_PORT, DEFAULT_TRAFFIC_PORT, DEFAULT_IST_PORT, DEFAULT_SST_PORT}

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

func (recipe *MariadbGaleraRecipient) Parse() (*bytes.Buffer, error) {
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

func (recipe *MariadbGaleraRecipient) Create() (*apiv1.Service, error) {
	b, err := recipe.Parse()
	if err != nil {
		return nil, err
	}

	obj := new(apiv1.Service)

	err = yaml.Unmarshal(b.Bytes(), obj)
	if err != nil {
		glog.Errorf("Deserialize artifact failed: %v", err)
		return nil, fmt.Errorf("Failed to perform deserializtion: %v\n", err)
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
