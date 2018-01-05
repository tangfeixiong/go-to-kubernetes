package operator

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	client "github.com/jw-s/redis-operator/pkg/generated/clientset"
	redisinformers "github.com/jw-s/redis-operator/pkg/generated/informers/externalversions"
	"github.com/jw-s/redis-operator/pkg/operator/controller"
	"github.com/jw-s/redis-operator/pkg/operator/util"
	opkit "github.com/rook/operator-kit"
	rookop "github.com/rook/rook/pkg/operator"
	"github.com/sirupsen/logrus"

	"k8s.io/api/extensions/v1beta1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	errorsUtil "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/util/version"

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pb"
	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/crd"
)

/*
  https://github.com/rook/rook/blob/master/pkg/operator/operator.go
*/
type Operator struct {
	config               *rest.Config
	kubeClient           kubernetes.Interface
	apiExtClientset      apiextensionsclient.Interface
	redisClient          client.Interface
	controllerConfig     controller.Config
	redisInformerFactory redisinformers.SharedInformerFactory
	informerFactory      informers.SharedInformerFactory
	c                    controller.Controller
	operator             *rookop.Operator
	Interval             time.Duration
	Timeout              time.Duration
	agents               map[string]*struct{}
}

func Run() (*Operator, error) {
	config, err := util.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("Could not create In-cluster config: " + err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Could not create kube client: " + err.Error())
	}

	apiExtClientset, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s API extension clientset. %+v", err)
	}

	redisClient, err := client.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Could not create redis client: " + err.Error())
	}

	controllerConfig := controller.NewConfig(config, Resync)

	redisInformerFactory := redisinformers.NewSharedInformerFactory(redisClient, Resync)

	informerFactory := informers.NewSharedInformerFactory(kubeClient, Resync)

	c := controller.New(controllerConfig,
		kubeClient,
		redisClient.RedisV1(),
		redisInformerFactory.Redis().V1().Redises(),
		informerFactory.Core().V1().Pods(),
		informerFactory.Apps().V1beta1().Deployments(),
		informerFactory.Core().V1().Services(),
		informerFactory.Core().V1().Endpoints(),
		informerFactory.Core().V1().ConfigMaps(),
		informerFactory.Apps().V1beta1().StatefulSets())

	doneChan := make(chan struct{})

	go c.Run(doneChan)

	go informerFactory.Start(doneChan)

	go redisInformerFactory.Start(doneChan)

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		for {
			select {
			case <-signalChan:
				logrus.Info("Shutdown signal received, exiting...")
				close(doneChan)
				// os.Exit(0)
				runtime.Goexit()
			}
		}
	}()

	return &Operator{
		config:               config,
		kubeClient:           kubeClient,
		apiExtClientset:      apiExtClientset,
		redisClient:          redisClient,
		controllerConfig:     controllerConfig,
		redisInformerFactory: redisInformerFactory,
		informerFactory:      informerFactory,
		c:                    c,
	}, nil
}

func (op *Operator) RunOutCluster(kubeconfig string) error {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("Unable to config kubernetes client:", err.Error())
		return err
	}
	op.config = config
	return err
}

func (op *Operator) CreateCRD(recipe *pb.CrdRecipient) error {
	r := crd.Recipient{
		Name:     recipe.Name,
		Group:    recipe.Group,
		Version:  recipe.Version,
		Scope:    recipe.Scope.String(),
		Plural:   recipe.Plural,
		Singular: recipe.Singular,
		Kind:     recipe.Kind,
	}
	return op.CreateCRDorTPR(r)
}

func (op *Operator) CreateCRDorTPR(recipe crd.Recipient) error {
	// CRD is available on v1.7.0 and above. TPR became deprecated on v1.7.0
	serverVersion, err := op.kubeClient.Discovery().ServerVersion()
	if err != nil {
		return fmt.Errorf("Error getting server version: %v", err)
	}
	kubeVersion := version.MustParseSemantic(serverVersion.GitVersion)

	const (
		serverVersionV170 = "v1.7.0"
	)

	var lastErr error
	if kubeVersion.AtLeast(version.MustParseSemantic(serverVersionV170)) {
		err = op.createCRD(recipe)
		if err != nil {
			lastErr = err
		}

		if err := op.waitForCRDInit(recipe); err != nil {
			lastErr = err
		}
	} else {
		context := opkit.Context{
			Clientset:             op.kubeClient,
			APIExtensionClientset: op.apiExtClientset,
			Interval:              500 * time.Millisecond,
			Timeout:               60 * time.Second,
		}
		scheme := opkit.CustomResource{
			Name:    recipe.Name,
			Plural:  recipe.Plural,
			Group:   recipe.Group,
			Version: recipe.Version,
			Scope:   apiextensionsv1beta1.NamespaceScoped,
			Kind:    recipe.Kind,
		}
		if recipe.Scope == string(apiextensionsv1beta1.ClusterScoped) || recipe.Scope == string(apiextensionsv1beta1.NamespaceScoped) {
			scheme.Scope = apiextensionsv1beta1.ResourceScope(recipe.Scope)
		}
		resources := []opkit.CustomResource{scheme}

		// Create and wait for TPR resources
		for _, resource := range resources {
			err = createTPR(context, resource)
			if err != nil {
				lastErr = err
			}
		}

		for _, resource := range resources {
			if err := waitForTPRInit(context, resource); err != nil {
				lastErr = err
			}
		}
	}
	return lastErr
}

func (op *Operator) createCRD(recipe crd.Recipient) error {
	rediscrd, err := crd.Deserialize(recipe)

	_, err = op.apiExtClientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(rediscrd)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return fmt.Errorf("failed to create %s CRD. %+v", recipe.Name, err)
		}
	}
	return nil
}

func (op *Operator) waitForCRDInit(recipe crd.Recipient) error {
	crdName := fmt.Sprintf("%s.%s", recipe.Name, recipe.Group)
	return wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		crdRedis, err := op.apiExtClientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crdName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, cond := range crdRedis.Status.Conditions {
			switch cond.Type {
			case apiextensionsv1beta1.Established:
				if cond.Status == apiextensionsv1beta1.ConditionTrue {
					return true, nil
				}
			case apiextensionsv1beta1.NamesAccepted:
				if cond.Status == apiextensionsv1beta1.ConditionFalse {
					return false, fmt.Errorf("Name conflict: %v\n", cond.Reason)
				}
			}
		}
		return false, nil
	})
}

func createTPR(context opkit.Context, resource opkit.CustomResource) error {
	tprName := fmt.Sprintf("%s.%s", resource.Name, resource.Group)
	tpr := &v1beta1.ThirdPartyResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: tprName,
		},
		Versions: []v1beta1.APIVersion{
			{Name: resource.Version},
		},
		Description: fmt.Sprintf("ThirdPartyResource for %s", resource.Name),
	}
	_, err := context.Clientset.ExtensionsV1beta1().ThirdPartyResources().Create(tpr)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return fmt.Errorf("failed to create %s TPR. %+v", resource.Name, err)
		}
	}
	return nil
}

func waitForTPRInit(context opkit.Context, resource opkit.CustomResource) error {
	// wait for TPR being established
	restcli := context.Clientset.CoreV1().RESTClient()
	uri := fmt.Sprintf("apis/%s/%s/%s", resource.Group, resource.Version, resource.Plural)
	tprName := fmt.Sprintf("%s.%s", resource.Name, resource.Group)

	err := wait.Poll(context.Interval, context.Timeout, func() (bool, error) {
		_, err := restcli.Get().RequestURI(uri).DoRaw()
		if err != nil {
			if errors.IsNotFound(err) {
				return false, nil
			}
			return false, err
		}
		return true, nil

	})
	if err != nil {
		deleteErr := context.Clientset.ExtensionsV1beta1().ThirdPartyResources().Delete(tprName, nil)
		if deleteErr != nil {
			return errorsUtil.NewAggregate([]error{err, deleteErr})
		}
		return err
	}
	return nil
}

/*
  https://github.com/jw-s/redis-operator/blob/master/cmd/operator/main.go
*/

var (
	Resync time.Duration = time.Minute
)

func main() {

	doneChan := make(chan struct{})

	config, err := util.InClusterConfig()
	if err != nil {
		panic("Could not create In-cluster config: " + err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic("Could not create kube client: " + err.Error())
	}

	redisClient, err := client.NewForConfig(config)

	if err != nil {
		panic("Could not create redis client: " + err.Error())
	}

	controllerConfig := controller.NewConfig(config, Resync)

	redisInformerFactory := redisinformers.NewSharedInformerFactory(redisClient, Resync)

	informerFactory := informers.NewSharedInformerFactory(kubeClient, Resync)

	c := controller.New(controllerConfig,
		kubeClient,
		redisClient.RedisV1(),
		redisInformerFactory.Redis().V1().Redises(),
		informerFactory.Core().V1().Pods(),
		informerFactory.Apps().V1beta1().Deployments(),
		informerFactory.Core().V1().Services(),
		informerFactory.Core().V1().Endpoints(),
		informerFactory.Core().V1().ConfigMaps(),
		informerFactory.Apps().V1beta1().StatefulSets())

	go c.Run(doneChan)

	go informerFactory.Start(doneChan)

	go redisInformerFactory.Start(doneChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-signalChan:
			logrus.Info("Shutdown signal received, exiting...")
			close(doneChan)
			os.Exit(0)
		}
	}

}
