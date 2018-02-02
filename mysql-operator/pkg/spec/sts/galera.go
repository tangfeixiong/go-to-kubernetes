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

	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/artifact"
)

const (
	DEFAULT_IMAGE        = "docker.io/mariadb:10"
	DEFAULT_CLIENT_PORT  = 3306
	DEFAULT_TRAFFIC_PORT = 4567
	DEFAULT_IST_PORT     = 4568
	DEFAULT_SST_PORT     = 4444
)

var (
	asset_default string = "template/mariadb-statefulset.yaml.tpl"

	name_default               string = "demo-galera"
	replicas_default           uint8  = 3
	container_name             string = "mariadb"
	client_connection          string = "client"
	replication_traffic        string = "replica-t"
	replication_traffic_udp    string = "replica-u"
	incremental_state_transfer string = "ist"
	state_snapshot_transfer    string = "sst"
)

type MariadbGaleraRecipient struct {
	DatabaseAccount
	Name, Namespace                                                            string
	CustomResourceName, ServiceName                                            string
	Image, ClientPortName, TrafficTcpName, TrafficUdpName, IstName, SstName    string
	Replications, ClientPort, TrafficTcpPort, TrafficUdpPort, IstPort, SstPort int
	Quorum                                                                     *uint8
	Weights                                                                    []*uint8
	Count                                                                      int
	ClusterName                                                                string
}

func NewMariadbGaleraRecipient(name, namespace, crn, service string, count *uint8, dbaccount *DatabaseAccount) *MariadbGaleraRecipient {
	return NewMariadbGaleraRecipient1(name, namespace, crn, service, name, count, dbaccount)
}

func NewMariadbGaleraRecipient1(name, namespace, crn, service, cluster string, count *uint8, dbaccount *DatabaseAccount) *MariadbGaleraRecipient {
	num := replicas_default
	if count != nil && *count > 0 {
		//num = (*count-1)*2 + 1
		num = *count
	}

	recipe := &MariadbGaleraRecipient{
		Name:               name,
		Namespace:          namespace,
		CustomResourceName: crn,
		ServiceName:        service,
		ClusterName:        cluster,
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
	if cluster == "" {
		recipe.ClusterName = recipe.ServiceName
	}

	//	if port == nil {
	//		recipe.ClientPort = DEFAULT_CLIENT_PORT
	//	} else if 1024 < *port && *port < 65535 {
	//		recipe.ClientPort = *port
	//	} else {
	//      ???
	//  }

	if dbaccount == nil {
		recipe.DatabaseAccount = default_dbaccount
	} else {
		recipe.DatabaseAccount = *dbaccount
	}

	return recipe
}

func (recipe *MariadbGaleraRecipient) Parse() (*bytes.Buffer, error) {
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

	glog.V(5).Infoln("StatefulSet artifact for Redis:", b.String())
	return b, nil
}

func (recipe *MariadbGaleraRecipient) Create() (*v1beta2.StatefulSet, error) {
	b, err := recipe.Parse()
	if err != nil {
		return nil, err
	}

	result := new(v1beta2.StatefulSet)

	err = yaml.Unmarshal(b.Bytes(), &result)
	if err != nil {
		glog.Errorf("Deserializing artifact failed: %v", err)
		return nil, fmt.Errorf("Failed to perform deserializtion: %v\n", err)
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
