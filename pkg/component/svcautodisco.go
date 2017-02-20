package component

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	kapi "k8s.io/kubernetes/pkg/api"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubelet/envvars"
	"k8s.io/kubernetes/pkg/labels"

	"github.com/tangfeixiong/go-to-kubernetes/pkg/api"
	"github.com/tangfeixiong/go-to-kubernetes/pkg/client"
)

const (
	service_host_suffix     = "_SERVICE_HOST"
	service_port_suffix     = "_SERVICE_PORT"
	service_name_kubernetes = "kubernetes"
)

/*
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=registry-2041433494-17v6m
REGISTRY_HTTP_TLS_CERTIFICATE=/certs/tls.crt
REGISTRY_HTTP_TLS_KEY=/certs/tls.key
DOCKER_REGISTRY_CONFIG=/etc/docker/registry/config.yml
REGISTRY_STORAGE_FILESYSTEM_ROOTDIRECTORY=/var/lib/registry
REGISTRY_PORT_5000_TCP_PORT=5000
OPENSHIFT_PORT_8443_TCP=tcp://10.3.0.30:8443
REGISTRY_PORT_5000_TCP_ADDR=10.3.0.205
OPENSHIFT_SERVICE_HOST=10.3.0.30
REGISTRY_PORT=tcp://10.3.0.205:5000
OPENSHIFT_SERVICE_PORT_OPENSHIFT=8443
OPENSHIFT_PORT_8443_TCP_ADDR=10.3.0.30
REGISTRY_SERVICE_PORT=5000
REGISTRY_SERVICE_PORT_REGISTRY=5000
REGISTRY_PORT_5000_TCP=tcp://10.3.0.205:5000
OPENSHIFT_SERVICE_PORT=8443
REGISTRY_PORT_5000_TCP_PROTO=tcp
OPENSHIFT_PORT_8443_TCP_PROTO=tcp
OPENSHIFT_PORT=tcp://10.3.0.30:8443
OPENSHIFT_PORT_8443_TCP_PORT=8443
REGISTRY_SERVICE_HOST=10.3.0.205
KUBERNETES_SERVICE_HOST=10.3.0.1
KUBERNETES_SERVICE_PORT=443
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_PORT=tcp://10.3.0.1:443
KUBERNETES_PORT_443_TCP=tcp://10.3.0.1:443
KUBERNETES_PORT_443_TCP_PROTO=tcp
KUBERNETES_PORT_443_TCP_ADDR=10.3.0.1
KUBERNETES_PORT_443_TCP_PORT=443
ETCD_V3_SINGLE_SERVICE_HOST=10.3.0.212
ETCD_V3_SINGLE_SERVICE_PORT=2379
ETCD_V3_SINGLE_SERVICE_PORT_CLIENT=2379
ETCD_V3_SINGLE_PORT=tcp://10.3.0.212:2379
ETCD_V3_SINGLE_PORT_2379_TCP=tcp://10.3.0.212:2379
ETCD_V3_SINGLE_PORT_2379_TCP_PROTO=tcp
ETCD_V3_SINGLE_PORT_2379_TCP_ADDR=10.3.0.212
ETCD_V3_SINGLE_PORT_2379_TCP_PORT=2379
HOME=/root
*/

type serivceAutoDiscovery struct {
	cf  *client.CmdUtilFactory
	err error
}

func NewServiceAutoDiscovery() *serivceAutoDiscovery {
	sad := &serivceAutoDiscovery{}
	sad.cf, sad.err = client.NewCmdUtilFactory("", "", "")
	return sad
}

func (sad *serivceAutoDiscovery) EnvVar(namespace string) ([]api.ServiceAutoDiscovery, error) {
	if sad.err != nil {
		return nil, sad.err
	}
	return ServiceAutoDiscoveryEnvVar(sad.cf.Factory, namespace)
}

func ServiceAutoDiscoveryEnvVar(f *kcmdutil.Factory, namespace string) ([]api.ServiceAutoDiscovery, error) {
	c, err := f.Client()
	if err != nil {
		return nil, err
	}
	return serviceAutoDiscoveryEnvVar(c, namespace)
}

func serviceAutoDiscoveryEnvVar(kc *kclient.Client, namespace string) ([]api.ServiceAutoDiscovery, error) {
	var logger *log.Logger = log.New(os.Stdout, "[component/ServiceAutoDiscoveryEnv] ", log.LstdFlags|log.Lshortfile)

	ns := "default"
	if len(namespace) > 0 {
		ns = namespace
	}
	var err error

	var services *kapi.ServiceList
	services, err = kc.Services(ns).List(kapi.ListOptions{})
	if err != nil {
		logger.Printf("Failed to access kubernetes for service list: %+v", err)
		return nil, err
	}
	//logger.Printf("Services retrieved (%d): %+v", len(services.Items), services)

	var all []api.ServiceAutoDiscovery
	var env []kapi.EnvVar
	var name string
	var pods *kapi.PodList
	for i := range services.Items {
		result := api.ServiceAutoDiscoveryEnv{}

		service := &services.Items[i]

		env = envvars.FromServices(&kapi.ServiceList{
			Items: []kapi.Service{*service},
		})

		for j := range env {
			switch {
			case strings.HasSuffix(env[j].Name, service_host_suffix):
				result.ServiceHostEnvVar = env[j]
			case strings.HasSuffix(env[j].Name, service_port_suffix):
				result.FirstPortEnvVar = env[j]
			}
		}

		all = append(all, api.ServiceAutoDiscovery{
			Env: result,
		})
		result = all[i].Env

		name = makeEnvVariableName(service.Name) + service_port_suffix
		// All named ports (only the first may be unnamed, checked in validation)
		for i := range service.Spec.Ports {
			sp := &service.Spec.Ports[i]
			if sp.Name != "" {
				pn := name + "_" + makeEnvVariableName(sp.Name)
				result.ServicePortsEnvVar = append(result.ServicePortsEnvVar, kapi.EnvVar{Name: pn, Value: strconv.Itoa(int(sp.Port))})
			}
		}
		// Docker-compatible vars.
		//result = append(result, makeLinkVariables(service)...)
		result.DockerCompatialEnv = makeLinkVariables(service)

		// For HOSTNAME, need investigation
		if strings.EqualFold(service_name_kubernetes, service.Name) {
			continue
		}

		pods, err = kc.Pods(ns).List(kapi.ListOptions{
			LabelSelector: labels.Set(service.Spec.Selector).AsSelector(),
		})
		if err != nil {
			logger.Printf("Failed to access kubernetes for pod list: %+v", err)
			continue
		}
		if len(pods.Items) == 0 {
			logger.Printf("Unexpected as nothind of pod list: %+v", err)
			continue
		}
		for j := range pods.Items {
			result.HostnameEnvVar = append(result.HostnameEnvVar, kapi.EnvVar{
				Name:  "HOSTNAME",
				Value: pods.Items[j].Name,
			})
		}
	}
	return all, nil
}

// to see: k8s.io/kubernetes/pkg/kubelet/envvars/envvars.go
//
func makeEnvVariableName(str string) string {
	// TODO: If we simplify to "all names are DNS1123Subdomains" this
	// will need two tweaks:
	//   1) Handle leading digits
	//   2) Handle dots
	return strings.ToUpper(strings.Replace(str, "-", "_", -1))
}

func makeLinkVariables(service *kapi.Service) /*[]api.EnvVar*/ api.DockerCompatialEnv {
	prefix := makeEnvVariableName(service.Name)
	//all := []api.EnvVar{}
	all := api.DockerCompatialEnv{
		DockerServicesEnvVar: make([]kapi.EnvVar, 0),
	}
	for i := range service.Spec.Ports {
		sp := &service.Spec.Ports[i]

		protocol := string(kapi.ProtocolTCP)
		if sp.Protocol != "" {
			protocol = string(sp.Protocol)
		}
		if i == 0 {
			// Docker special-cases the first port.
			//all = append(all, api.EnvVar{
			//	Name:  prefix + "_PORT",
			//	Value: fmt.Sprintf("%s://%s:%d", strings.ToLower(protocol), service.Spec.ClusterIP, sp.Port),
			//})
			all.DockerFirstServiceEnvVar = kapi.EnvVar{
				Name:  prefix + "_PORT",
				Value: fmt.Sprintf("%s://%s:%d", strings.ToLower(protocol), service.Spec.ClusterIP, sp.Port),
			}
		}
		portPrefix := fmt.Sprintf("%s_PORT_%d_%s", prefix, sp.Port, strings.ToUpper(protocol))
		all.DockerServicesEnvVar = append(all.DockerServicesEnvVar, []kapi.EnvVar{
			{
				Name:  portPrefix,
				Value: fmt.Sprintf("%s://%s:%d", strings.ToLower(protocol), service.Spec.ClusterIP, sp.Port),
			},
			{
				Name:  portPrefix + "_PROTO",
				Value: strings.ToLower(protocol),
			},
			{
				Name:  portPrefix + "_PORT",
				Value: strconv.Itoa(int(sp.Port)),
			},
			{
				Name:  portPrefix + "_ADDR",
				Value: service.Spec.ClusterIP,
			},
		}...)
	}
	return all
}
