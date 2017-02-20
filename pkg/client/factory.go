package client

import (
	_ "bufio"
	_ "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/url"
	"os"
	"regexp"
	"strings"

	"golang.org/x/build/kubernetes"

	k8sapi "golang.org/x/build/kubernetes/api"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/restclient"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	clientcmdapi "k8s.io/kubernetes/pkg/client/unversioned/clientcmd/api"
	"k8s.io/kubernetes/pkg/fields"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/labels"
)

type CmdUtilFactory struct {
	*kcmdutil.Factory
}

func NewCmdUtilFactory(kubeconfigPath, kubeApiserver, kubectlContext string) (*CmdUtilFactory, error) {
	cc, err := customClientConfig(kubeconfigPath, kubeApiserver, kubectlContext)
	if err != nil {
		return nil, err
	}
	f := &CmdUtilFactory{
		Factory: kcmdutil.NewFactory(cc),
	}
	return f, nil
}

/*func DefaultClientConfig(flags *pflag.FlagSet) clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	flags.StringVar(&loadingRules.ExplicitPath, "kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")

	overrides := &clientcmd.ConfigOverrides{}
	flagNames := clientcmd.RecommendedConfigOverrideFlags("")
	// short flagnames are disabled by default.  These are here for compatibility with existing scripts
	flagNames.ClusterOverrideFlags.APIServer.ShortName = "s"

	clientcmd.BindOverrideFlags(overrides, flags, flagNames)
	clientConfig := clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, overrides, os.Stdin)

	return clientConfig
}*/

func customClientConfig(kubeconfigPath, kubeApiserver, kubectlContext string) (clientcmd.ClientConfig, error) {
	var logger *log.Logger = log.New(os.Stdout, "[client/customClientConfig] ", log.LstdFlags|log.Lshortfile)
	var loadingRules *clientcmd.ClientConfigLoadingRules

	var overrides *clientcmd.ConfigOverrides = &clientcmd.ConfigOverrides{}
	if len(kubeApiserver) > 0 {
		overrides.ClusterInfo.Server = kubeApiserver
	}

	var clientConfig clientcmd.ClientConfig

	if len(kubeconfigPath) == 0 {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
		clientConfig = clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, overrides, os.Stdin)
		return clientConfig, nil
	}

	var data []byte
	var err error
	data, err = ioutil.ReadFile(kubeconfigPath)
	if err != nil {
		logger.Printf("kubeconfig not accessed: %+v\n", err)
		return nil, err
	}
	logger.Printf("kubeconfig: \n%+v\n", string(data))
	var conf *clientcmdapi.Config
	conf, err = clientcmd.Load(data)
	if err != nil {
		logger.Printf("kubeconfig invalid: %v\n", err)
		return nil, err
	}

	clientConfig = clientcmd.NewNonInteractiveClientConfig(*conf, kubectlContext, overrides, loadingRules)
	return clientConfig, nil
}

type ClientWrapper struct {
	k8sConfig *restclient.Config
	k8sClient *kclient.Client
	*kubernetes.Client
	apiUrl     string
	httpClient *http.Client
}

const (
	apiVersion = "v1"
)

func NewClientWrapper(apiHost string, userName string, passWord string) (*ClientWrapper, error) {
	conf := &restclient.Config{
		Host:   apiHost,
		Prefix: "",
		//Version:    "v1",
		//GroupVersion: &unversioned.GroupVersion{"", "v1"},
		QPS:      5.0,
		Burst:    10,
		Username: userName,
		Password: passWord,
		Insecure: true,
	}

	wrapper := &ClientWrapper{
		k8sConfig: conf,
	}

	if client, err := kclient.New(conf); err == nil {
		wrapper.k8sClient = client
		return wrapper, nil
	}

	wrapper.httpClient = &http.Client{
	//Transport: &oauth2.Transport{
	//    Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "aCcessWbU3toKen"}),
	//}
	}

	baseUrl := apiHost
	if m, _ := regexp.MatchString(`(?i)https?://`, apiHost); !m {
		baseUrl = "http://" + apiHost
	}

	if c, err := kubernetes.NewClient(baseUrl, wrapper.httpClient); err == nil {
		return nil, err
	} else {
		wrapper.Client = c
	}

	wrapper.apiUrl = strings.TrimSuffix(baseUrl, "/") + "/api/" + apiVersion

	return wrapper, nil
}

func (c *ClientWrapper) RunPod(pod *k8sapi.Pod) (*k8sapi.PodStatus, error) {
	ctx := context.TODO()

	return c.Client.RunLongLivedPod(ctx, pod)
}

func (c *ClientWrapper) GetPods() ([]k8sapi.Pod, error) {
	ctx := context.TODO()

	return c.Client.GetPods(ctx)
}

// Returns all RCs in the cluster.
func (c *ClientWrapper) GetDefaultReplicationControllers() ([]k8sapi.ReplicationController, error) {
	if c.Client != nil {
		ctx := context.TODO()

		getURL := c.apiUrl + "/namespaces/default/replicationcontrollers"

		// Make request to Kubernetes API
		req, err := http.NewRequest("GET", getURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: GET %q : %v", getURL, err)
		}
		res, err := ctxhttp.Do(ctx, c.httpClient, req)
		if err != nil {
			return nil, fmt.Errorf("failed to make request: GET %q: %v", getURL, err)
		}

		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to read request body for GET %q: %v", getURL, err)
		}
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("http error %d GET %q: %q: %v", res.StatusCode, getURL, string(body), err)
		}

		var rcList k8sapi.ReplicationControllerList
		if err := json.Unmarshal(body, &rcList); err != nil {
			return nil, fmt.Errorf("failed to decode list of pod resources: %v", err)
		}
		return rcList.Items, nil
	}
	return nil, fmt.Errorf("client not configed")
}

func (c *ClientWrapper) FindReplicationControllerList() (*api.ReplicationControllerList, error) {
	if c.k8sClient != nil {
		rcFinder := c.k8sClient.ReplicationControllers("default")

		return rcFinder.List(api.ListOptions{
			TypeMeta: unversioned.TypeMeta{
				Kind:       "List", //"ReplicationControllerList",
				APIVersion: "v1",
			},
			LabelSelector: labels.Everything(),
			FieldSelector: fields.Everything(),
		})

		//return rcFinder.List(labels.Everything())
	}
	return nil, fmt.Errorf("client not configed")
}

func (c *ClientWrapper) FindReplicationController(name string) (*api.ReplicationController, error) {
	if c.k8sClient != nil {
		rcFinder := c.k8sClient.ReplicationControllers("default")
		return rcFinder.Get(name)
	}
	return nil, fmt.Errorf("client not configed")
}

func (c *ClientWrapper) FindServiceList() (*api.ServiceList, error) {
	if c.k8sClient != nil {
		svcFinder := c.k8sClient.Services("default")

		return svcFinder.List(api.ListOptions{
			TypeMeta: unversioned.TypeMeta{
				Kind:       "ServiceList",
				APIVersion: "v1",
			},
			LabelSelector: labels.Everything(),
			FieldSelector: fields.Everything(),
		})

		//return svcFinder.List(labels.Everything())
	}
	return nil, fmt.Errorf("client not configed")
}

func (c *ClientWrapper) FindService(name string) (*api.Service, error) {
	if c.k8sClient != nil {
		svcFinder := c.k8sClient.Services("default")
		return svcFinder.Get(name)
	}
	return nil, fmt.Errorf("client not configed")
}
