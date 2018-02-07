package hadoop

import (
	"fmt"

	//"github.com/helm/helm-classic/codec"

	apiv1 "k8s.io/api/core/v1"
	//storagev1 "k8s.io/api/storage/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	hcm "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/hadoop"
	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/spec/sts"
	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/spec/svc"
	"github.com/tangfeixiong/go-to-kubernetes/pkg/spec/crd"
)

func CreateCRD(apiextensionsclientset apiextensionsclient.Interface) {
	recipe, err := crd.NewRecipient("example.com", "v1", "hdfses", "hdfs")

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

func CreateHdfsSvc(clientset kubernetes.Interface) {
	recipe := svc.NewHadoopHdfsNodeServiceRecipient("my-hdfs", "default", "my-hadoop")
	service, err := recipe.New()
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

func CreateHdfsCm(clientset kubernetes.Interface) {
	cfg := hcm.Config{
		Name:               "my-hdfs",
		ServiceName:        "my-hdfs",
		Namespace:          "default",
		BaseDomain:         "cluster.local",
		CustomResourceName: "my-hadoop",
	}

	result, err := cfg.InitMap(clientset)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created config map %q.\n", result.GetObjectMeta().GetName())
}

func CreateHdfsSts(clientset kubernetes.Interface) {
	count := int32(2)
	recipe := sts.NewHadoopHdfsNodeRecipient("my-hdfs", "default", "my-hadoop", "my-hdfs", "my-cluster", "namenode", &count)
	statefulset, err := recipe.New()
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
