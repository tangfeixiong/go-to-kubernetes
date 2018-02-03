package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	apiv1 "k8s.io/api/core/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/mysqlorbranches"
	"github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/rabbitmq"
	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/crd"
	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/deploy"
	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/sts"
	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/svc"
)

/*
  Inspired by:
    https://github.com/kubernetes/blob/master/client-go/examples/out-of-cluster-client-configuration/main.go
*/
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

var (
	kubeconfig *string
	stuff      bool
	config     *restclient.Config
	clientset  kubernetes.Interface
)

func main() {
	//var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.BoolVar(&stuff, "stuff", false, "for test")
	flag.Parse()

	var err error
	// use the current context in kubeconfig
	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	start := time.Now()
	for stuff && time.Since(start) < time.Duration(10)*time.Second {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get("example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod not found\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod\n")
		}

		time.Sleep(10 * time.Second)
	}

	switch os.Args[1] {
	case "create":
		switch os.Args[2] {
		case "redis-sts":
			createRedisStatefulSet()
		case "redis-svc":
			createRedisSvc()
		case "sentinel-deploy":
			createSentinelDeploy()
		case "sentinel-svc":
			createSentinelSvc()
		case "redis-crd":
			createRedisCrd()
			break
		case "mariadb-crd":
			apiextensionsclientset, err := apiextensionsclient.NewForConfig(config)
			if err != nil {
				panic(err)
			}
			mysqlorbranches.CreateCRD(apiextensionsclientset)
		case "mariadb-hostpath":
			mysqlorbranches.CreateHostPath(clientset)
		case "mariadb-localstorage":
			mysqlorbranches.CreateLocalStorage(clientset)
		case "mariadb-svc":
			mysqlorbranches.CreateMariadbSvc(clientset)
		case "mariadb-sts":
			mysqlorbranches.CreateMariadbSts(clientset)

		case "rabbitmq-svc":
			rabbitmq.CreateSvc(clientset)
		case "rabbitmq-sts":
			rabbitmq.CreateSts(clientset)

		default:
			fmt.Println("Unknown create option", os.Args[1])
		}
	default:
		fmt.Println("Unknown command", os.Args[1])
	}
}

func createRedisStatefulSet() {
	slaves := 1
	recipe := sts.NewRedisRecipient("my-redis", "", "my-redis", "my-redis", slaves)
	statefulset, err := recipe.GenerateArtifact()
	if err != nil {
		panic(err)
	}

	statefulsetsClient := clientset.AppsV1beta2().StatefulSets(apiv1.NamespaceDefault)

	result, err := statefulsetsClient.Create(statefulset)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created statefulset %q.\n", result.GetObjectMeta().GetName())
}

func createRedisSvc() {
	recipe := svc.NewRedisRecipient("my-redis", "", "my-redis", "client", "gossip", 6379)
	service, err := recipe.GenerateArtifact()
	if err != nil {
		panic(err)
	}

	serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)

	result, err := serviceClient.Create(service)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created service %q.\n", result.GetObjectMeta().GetName())
}

func createSentinelDeploy() {
	quorum := 2
	recipe := deploy.NewSentinelRecipient("my-redis", "", "my-redis-ha", quorum)
	deployment, err := recipe.GenerateArtifact()
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1beta2().Deployments(apiv1.NamespaceDefault)
	//deploymentsClient := clientset.ExtensionsV1beta1().Deployments(apiv1.NamespaceDefault)

	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func createSentinelSvc() {
	recipe := svc.NewSentinelRecipient("my-redis-ha", "", "my-redis-ha", "client", 26379)
	service, err := recipe.GenerateArtifact()
	if err != nil {
		panic(err)
	}

	serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)

	result, err := serviceClient.Create(service)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created service %q.\n", result.GetObjectMeta().GetName())
}

func createRedisCrd() {
	var recipe crd.Recipient = crd.Recipient{
		Name:     "redises",
		Group:    "example.com",
		Version:  "v1",
		Scope:    "Namespaced",
		Plural:   "redises",
		Singular: "redis",
		Kind:     "Redis",
	}

	customeresourcedefinition, err := crd.Deserialize(recipe)
	if err != nil {
		panic(err)
	}

	apiextensionsclientset, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	crdClient := apiextensionsclientset.ApiextensionsV1beta1().CustomResourceDefinitions()

	result, err := crdClient.Create(customeresourcedefinition)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created CRD %q.\n", result.Name)
}
