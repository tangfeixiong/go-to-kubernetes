
package client

import (
	"fmt"
	"io/ioutil"
    "log"
    "strings"

	"k8s.io/kubernetes/pkg/api"
	apiunversioned "k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	clientcmdapi "k8s.io/kubernetes/pkg/client/unversioned/clientcmd/api"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/yaml"
    
)

type K8sClientConfig struct {
	ClusterID string
	NameSpace string
	// Server is the address of the kubernetes cluster (https://hostname:port).
	Server string `json:"server"`
	// APIVersion is the preferred api version for communicating with the kubernetes cluster (v1, v2, etc).
	APIVersion string `json:"api-version,omitempty"`
	// InsecureSkipTLSVerify skips the validity check for the server's certificate. This will make your HTTPS connections insecure.
	InsecureSkipTLSVerify bool `json:"insecure-skip-tls-verify,omitempty"`
	// CertificateAuthority is the path to a cert file for the certificate authority.
	CertificateAuthority string `json:"certificate-authority,omitempty"`
	// CertificateAuthorityData contains PEM-encoded certificate authority certificates. Overrides CertificateAuthority
	CertificateAuthorityData []byte `json:"certificate-authority-data,omitempty"`

	// ClientCertificate is the path to a client cert file for TLS.
	ClientCertificate string `json:"client-certificate,omitempty"`
	// ClientCertificateData contains PEM-encoded data from a client cert file for TLS. Overrides ClientCertificate
	ClientCertificateData []byte `json:"client-certificate-data,omitempty"`
	// ClientKey is the path to a client key file for TLS.
	ClientKey string `json:"client-key,omitempty"`
	// ClientKeyData contains PEM-encoded data from a client key file for TLS. Overrides ClientKey
	ClientKeyData []byte `json:"client-key-data,omitempty"`
	config        *clientcmdapi.Config
	client        *unversioned.Client
}

func (k8scc *K8sClientConfig) UnversionedClient() *unversioned.Client {
	config := clientcmdapi.NewConfig()
	config.Clusters[k8scc.ClusterID] = &clientcmdapi.Cluster{
		InsecureSkipTLSVerify:    k8scc.InsecureSkipTLSVerify,
		Server:                   k8scc.Server,
		CertificateAuthority:     k8scc.CertificateAuthority,
		CertificateAuthorityData: k8scc.CertificateAuthorityData,
	}
	config.AuthInfos[k8scc.ClusterID] = &clientcmdapi.AuthInfo{
		ClientCertificate:     k8scc.ClientCertificate,
		ClientKey:             k8scc.ClientKey,
		ClientCertificateData: k8scc.ClientCertificateData,
		ClientKeyData:         k8scc.ClientKeyData,
	}
	config.Contexts[k8scc.ClusterID] = &clientcmdapi.Context{
		Cluster:  k8scc.ClusterID,
		AuthInfo: k8scc.ClusterID,
	}
	config.CurrentContext = k8scc.ClusterID

	clientBuilder := clientcmd.NewNonInteractiveClientConfig(*config, k8scc.ClusterID, &clientcmd.ConfigOverrides{})

	clientConfig, err := clientBuilder.ClientConfig()
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	client, err := unversioned.New(clientConfig)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	
	return client
}

func (conf *K8sClientConfig) Init() {
	conf.client = conf.UnversionedClient()
}


func (conf *K8sClientConfig) CreateRc(fileName string) (rc *api.ReplicationController, err error) {
	if conf.client == nil {
		conf.Init()
	}
	if !strings.HasSuffix(fileName, ".yaml") && !strings.HasSuffix(fileName, ".json") {
		err = fmt.Errorf("Only Support Yaml and Json")
		return
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unexpected error while reading file: %v", err)
		return
	}
	jsonData := data
	if strings.HasSuffix(fileName, ".yaml") {
		jsonData, err = yaml.ToJSON(data)
		if err != nil {
			log.Fatalf("Unexpected error while changing Yaml to Json: %v", err)
			return
		}
	}
	//log.Println(string(jsonData))
	ctrl := &api.ReplicationController{}
	if err = runtime.DecodeInto(api.Codecs.LegacyCodec(apiunversioned.GroupVersion{}), jsonData, ctrl); err != nil {
		log.Fatalf("Unexpected error decoding rc: %v", err)
	}
	//log.Println(ctrl.Spec.Template.)
	rc, err = conf.createReplicationControllers(ctrl)
	return
}
func (conf *K8sClientConfig) CreateRcByInput(data []byte) (rc *api.ReplicationController, err error) {
	if conf.client == nil {
		conf.Init()
	}
	jsonData := data
	tmp, err := yaml.ToJSON(data)
	if err == nil {
		jsonData = tmp
	}
	ctrl := &api.ReplicationController{}
	if err = runtime.DecodeInto(api.Codecs.LegacyCodec(apiunversioned.GroupVersion{}), jsonData, ctrl); err != nil {
		log.Fatalf("Unexpected error decoding rc: %v", err)
	}
	rc, err = conf.createReplicationControllers(ctrl)
	return
}
func (conf *K8sClientConfig) GetRc(name string) (rc *api.ReplicationController, err error) {
	if conf.client == nil {
		conf.Init()
	}
	rc, err = conf.client.ReplicationControllers(conf.NameSpace).Get(name)
	return

}

func (conf *K8sClientConfig) ListRc(opts api.ListOptions) (rc *api.ReplicationControllerList, err error) {
	if conf.client == nil {
		conf.Init()
	}
	rc, err = conf.client.ReplicationControllers(conf.NameSpace).List(opts)
	return

}

func (conf *K8sClientConfig) DeleteRc(name string) (err error) {
	if conf.client == nil {
		conf.Init()
	}
	err = conf.client.ReplicationControllers(conf.NameSpace).Delete(name)
	return

}
func (conf *K8sClientConfig) ScaleRc(name string, num int) (rc *api.ReplicationController, err error) {
	if conf.client == nil {
		conf.Init()
	}
	ctrl, err := conf.client.ReplicationControllers(conf.NameSpace).Get(name)
	if err != nil {
		log.Fatalf("Unexpected error : %v", err)
		return
	}
	ctrl.Spec.Replicas = num
	rc, err = conf.updateReplicationControllers(ctrl)
	return
}

func (conf *K8sClientConfig) createReplicationControllers(ctrl *api.ReplicationController) (rc *api.ReplicationController, err error) {
	rc, err = conf.client.ReplicationControllers(conf.NameSpace).Create(ctrl)
	return

}

func (conf *K8sClientConfig) updateReplicationControllers(ctrl *api.ReplicationController) (rc *api.ReplicationController, err error) {
	rc, err = conf.client.ReplicationControllers(conf.NameSpace).Update(ctrl)
	return

}
