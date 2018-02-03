package rabbitmq

import (
	"fmt"

	//"github.com/helm/helm-classic/codec"

	apiv1 "k8s.io/api/core/v1"
	//storagev1 "k8s.io/api/storage/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/tangfeixiong/go-to-kubernetes/pkg/spec/crd"
	//"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/spec/artifact"
	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/spec/sts"
	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/spec/svc"
)

func CreateCRD(apiextensionsclientset apiextensionsclient.Interface) {
	recipe, err := crd.NewRecipient("example.com", "v1", "rabbitmqs", "rabbitmq")

	customeresourcedefinition, err := recipe.Generate()
	if err != nil {
		panic(err)
	}

	crdClient := apiextensionsclientset.ApiextensionsV1beta1().CustomResourceDefinitions()

	result, err := crdClient.Create(customeresourcedefinition)
	if errors.IsAlreadyExists(err) {
		fmt.Printf("Already exists crd %q.\n", customeresourcedefinition.Name)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created CRD %q.\n", result.Name)
}

func CreateSvc(clientset kubernetes.Interface) {
	recipe := svc.NewRabbitmqRecipient("my-rabbitmq", "default", "my-rabbitmq")
	service, err := recipe.Generate()
	if err != nil {
		panic(err)
	}

	serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)

	result, err := serviceClient.Create(service)
	if errors.IsAlreadyExists(err) {
		fmt.Printf("Already exists SVC %q.\n", service.Name)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created service %q.\n", result.GetObjectMeta().GetName())
}

func CreateSts(clientset kubernetes.Interface) {
	count := int32(3)
	recipe := sts.NewRabbitmqRecipient("my-rabbitmq", "", "my-rabbitmq", "", &count)
	statefulset, err := recipe.Generate()
	if err != nil {
		panic(err)
	}

	statefulsetsClient := clientset.AppsV1beta2().StatefulSets(apiv1.NamespaceDefault)

	result, err := statefulsetsClient.Create(statefulset)
	if errors.IsAlreadyExists(err) {
		fmt.Printf("Already exists StatefulSet %q.\n", statefulset.Name)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created statefulset %q.\n", result.GetObjectMeta().GetName())

}
