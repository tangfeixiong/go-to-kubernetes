package client

import (
	_ "bufio"
	_ "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "log"
	"net/http"
	_ "net/url"
	"regexp"
	"strings"

	"github.com/golang/build/kubernetes"
	k8sapi "golang.org/x/build/kubernetes/api"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"k8s.io/kubernetes/pkg/api"
	unversionedapi "k8s.io/kubernetes/pkg/api/unversioned"
	unversionedclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
)

type ClientWrapper struct {
	k8sConfig *unversionedclient.Config
	k8sClient *unversionedclient.Client
	*kubernetes.Client
	apiUrl     string
	httpClient *http.Client
}

const (
	apiVersion = "v1"
)

func NewClientWrapper(apiHost string, userName string, passWord string) (*ClientWrapper, error) {
	conf := &unversionedclient.Config{
		Host:   apiHost,
		Prefix: "",
		//Version:    "v1",
		//GroupVersion: &unversionedapi.GroupVersion{"", "v1"},
		QPS:      5.0,
		Burst:    10,
		Username: userName,
		Password: passWord,
		Insecure: true,
	}

	wrapper := &ClientWrapper{
		k8sConfig: conf,
	}

	if client, err := unversionedclient.New(conf); err == nil {
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

func (c *ClientWrapper) RunPod(pod *k8sapi.Pod) (*k8sapi.Pod, error) {
	ctx := context.TODO()

	return c.Client.RunPod(ctx, pod)
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
			TypeMeta: unversionedapi.TypeMeta{
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
			TypeMeta: unversionedapi.TypeMeta{
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
