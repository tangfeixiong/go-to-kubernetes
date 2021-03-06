package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type ExampleV1alpha1Interface interface {
	RESTClient() rest.Interface
	HadoopsGetter
	HadoopHdfsesGetter
	HadoopMapreducesGetter
	HadoopNodesGetter
	HadoopYarnsGetter
}

// ExampleV1alpha1Client is used to interact with features provided by the example.com group.
type ExampleV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ExampleV1alpha1Client) Hadoops(namespace string) HadoopInterface {
	return newHadoops(c, namespace)
}

func (c *ExampleV1alpha1Client) HadoopHdfses(namespace string) HadoopHdfsInterface {
	return newHadoopHdfses(c, namespace)
}

func (c *ExampleV1alpha1Client) HadoopMapreduces(namespace string) HadoopMapreduceInterface {
	return newHadoopMapreduces(c, namespace)
}

func (c *ExampleV1alpha1Client) HadoopNodes(namespace string) HadoopNodeInterface {
	return newHadoopNodes(c, namespace)
}

func (c *ExampleV1alpha1Client) HadoopYarns(namespace string) HadoopYarnInterface {
	return newHadoopYarns(c, namespace)
}

// NewForConfig creates a new ExampleV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ExampleV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ExampleV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ExampleV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ExampleV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ExampleV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ExampleV1alpha1Client {
	return &ExampleV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ExampleV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
