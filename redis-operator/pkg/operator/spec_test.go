package operator

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/po"
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
	clientset  kubernetes.Interface
)

func init() {
	//var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.BoolVar(&stuff, "stuff", false, "for test")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
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
}

func TestStuff(t *testing.T) {
	if !stuff {
		fmt.Println("Must: go test -v -run Stuff ./pkg/operator -args --stuff")
	}
}

func TestPOD_RedisBootstrap(t *testing.T) {
	var recipe po.Recipient = po.Recipient{
		ClusterName: "my-redis",
		Name:        "redis",
	}
	artfact, err := po.Deserialize(recipe)
	if err != nil {
		t.Fatal("Could not deserilize POD from go template:", err)
	}

	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	result, err := podClient.Create(artfact)
	if err != nil {
		t.Fatal("Could not create POD:", err)
	}
	fmt.Printf("Created POD %q.\n", result.GetObjectMeta().GetName())
}

/*
  Inspred by:
    https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go
*/
func TestDepoly_RedisSentinel(t *testing.T) {

	deploymentsClient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)

	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}
