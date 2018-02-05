package operator

import (
	"fmt"
	//"os"
	//"os/signal"
	//"runtime"
	//"syscall"
	"time"

	opkit "github.com/rook/operator-kit"
	rookop "github.com/rook/rook/pkg/operator"
	//"github.com/sirupsen/logrus"

	"k8s.io/api/extensions/v1beta1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	errorsUtil "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/util/version"

	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pb"
	clientset "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned"
	informers "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/informers/externalversions"
	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/controller"
	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/hadoop"
	"github.com/tangfeixiong/go-to-kubernetes/pkg/signals"
	"github.com/tangfeixiong/go-to-kubernetes/pkg/spec/crd"
)

/*
  Inspired by:
    https://github.com/rook/rook/blob/master/pkg/operator/operator.go
*/
type Operator struct {
	cfg                   *rest.Config
	kubeClientset         kubernetes.Interface
	kubeApiExtClientset   apiextensionsclient.Interface
	sampleClientset       clientset.Interface
	sampleInformerFactory informers.SharedInformerFactory
	kubeInformerFactory   kubeinformers.SharedInformerFactory
	operator              *rookop.Operator
	Interval              time.Duration
	Timeout               time.Duration
	agents                map[string]*struct{}
	c                     *controller.Controller
}

func Run(conf hadoop.Config) (*Operator, error) {
	var cfg *rest.Config
	var err error
	if conf.Kubeconfig != "" {
		cfg, err = clientcmd.BuildConfigFromFlags("", conf.Kubeconfig)
		if err != nil {
			fmt.Println("Read kubernetes config failed:", err.Error())
			return nil, err
		}
	} else {
		cfg, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("Read Kubernetes in-cluster config failed: " + err.Error())
		}
	}

	op := new(Operator)
	ch := make(chan error)

	go op.main(cfg, ch)

	err = <-ch
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (op *Operator) main(cfg *rest.Config, errCh chan<- error) {
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	//	config, err := util.InClusterConfig()
	//	if err != nil {
	//		return nil, fmt.Errorf("Could not create In-cluster config: " + err.Error())
	//	}

	kubeClientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		//glog.Fatalf("Error building kubeconfig: %s", err.Error())
		errCh <- fmt.Errorf("Error building kubeconfig: %s", err.Error())
		return
		//return nil, fmt.Errorf("Could not create kube client: " + err.Error())
	}

	kubeApiExtClientset, err := apiextensionsclient.NewForConfig(cfg)
	if err != nil {
		errCh <- fmt.Errorf("Error building kubernetes api extensions clientset: %s", err.Error())
		return
		//return nil, fmt.Errorf("failed to create k8s API extension clientset. %+v", err)
	}

	exampleClientset, err := clientset.NewForConfig(cfg)
	if err != nil {
		//glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
		errCh <- fmt.Errorf("Error building kubernetes clientset: %s", err.Error())
		return
		//return nil, fmt.Errorf("Could not create redis CRD clientset: " + err.Error())
	}

	//	controllerConfig := controller.NewConfig(config, Resync)

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClientset, time.Second*30)
	exampleInformerFactory := informers.NewSharedInformerFactory(exampleClientset, time.Second*30)

	controller := controller.NewController(kubeClientset, exampleClientset, kubeInformerFactory, exampleInformerFactory)
	controller.Config = &struct {
		*rest.Config
		DefaultResync time.Duration
	}{cfg, Resync}

	go kubeInformerFactory.Start(stopCh)
	go exampleInformerFactory.Start(stopCh)

	op.cfg = cfg
	op.c = controller
	op.kubeClientset = kubeClientset
	op.kubeApiExtClientset = kubeApiExtClientset
	op.sampleClientset = exampleClientset
	op.sampleInformerFactory = exampleInformerFactory
	op.kubeInformerFactory = kubeInformerFactory

	if err = controller.Run( /*2*/ 1, stopCh); err != nil {
		//glog.Fatalf("Error running controller: %s", err.Error())
		errCh <- fmt.Errorf("Error running controller: %s", err.Error())
		return
	}
}

func (op *Operator) CreateCRD(recipe *pb.CrdRecipient) error {
	r := crd.Recipient{
		Group:    recipe.Group,
		Version:  recipe.Version,
		Scope:    recipe.Scope,
		Plural:   recipe.Plural,
		Singular: recipe.Singular,
		Kind:     recipe.Kind,
	}
	return op.CreateCRDorTPR(r)
}

func (op *Operator) CreateCRDorTPR(recipe crd.Recipient) error {
	// CRD is available on v1.7.0 and above. TPR became deprecated on v1.7.0
	serverVersion, err := op.kubeClientset.Discovery().ServerVersion()
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
			Clientset:             op.kubeClientset,
			APIExtensionClientset: op.kubeApiExtClientset,
			Interval:              500 * time.Millisecond,
			Timeout:               60 * time.Second,
		}
		scheme := opkit.CustomResource{
			Name:    fmt.Sprintf("%s.%s", recipe.Plural, recipe.Group),
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
	result, err := recipe.Generate()

	_, err = op.kubeApiExtClientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(result)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return fmt.Errorf("Create CRD failed. %+v", err)
		}
	}
	return nil
}

func (op *Operator) waitForCRDInit(recipe crd.Recipient) error {
	crdName := fmt.Sprintf("%s.%s", recipe.Plural, recipe.Group)
	return wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		crdRedis, err := op.kubeApiExtClientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crdName, metav1.GetOptions{})
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
  Inspired by:
    https://github.com/jw-s/redis-operator/blob/master/cmd/operator/main.go
*/

var (
	Resync time.Duration = time.Minute
)
