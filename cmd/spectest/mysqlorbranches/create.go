package mysqlorbranches

import (
	"fmt"

	"github.com/helm/helm-classic/codec"

	apiv1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/artifact"
	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/sts"
	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/svc"
	"github.com/tangfeixiong/go-to-kubernetes/pkg/spec/crd"
)

func CreateCRD(apiextensionsclientset apiextensionsclient.Interface) {
	recipe, err := crd.NewRecipient("example.com", "v1", "mysqlfamilies", "mysqlfamily")

	customeresourcedefinition, err := recipe.DecodeArtifact(nil)
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

func CreateLocalStorage(clientset kubernetes.Interface) {
	asset, err := artifact.Asset("template/local-storage-provision.yaml")
	if err != nil {
		panic(err)
	}

	decoders, err := codec.YAML.Decode(asset).All()
	if err != nil {
		panic(err)
	}
	for _, decoder := range decoders {
		tm := new(metav1.TypeMeta)
		if err := decoder.Object(tm); err != nil {
			panic(err)
		}
		switch {
		case tm.Kind == "StorageClass":
			sc := new(storagev1.StorageClass)
			if err := decoder.Object(sc); err != nil {
				panic(err)
			}

			storageClient := clientset.StorageV1().StorageClasses()

			result, err := storageClient.Create(sc)
			if errors.IsAlreadyExists(err) {
				fmt.Printf("Already exists StorageClass %q.\n", sc.Name)
				continue
			}
			if err != nil {
				panic(err)
			}
			fmt.Printf("Created StorageClass %q.\n", result.Name)
		case tm.Kind == "PersistentVolume":
			pv := new(apiv1.PersistentVolume)
			if err := decoder.Object(pv); err != nil {
				panic(err)
			}

			pvClient := clientset.CoreV1().PersistentVolumes()

			result, err := pvClient.Create(pv)
			if errors.IsAlreadyExists(err) {
				fmt.Printf("Already exists PV %q.\n", pv.Name)
				continue
			}
			if err != nil {
				panic(err)
			}
			fmt.Printf("Created PersistentVolume %q.\n", result.Name)
		case tm.Kind == "PersistentVolumeClaim":
			pvc := new(apiv1.PersistentVolumeClaim)
			if err := decoder.Object(pvc); err != nil {
				panic(err)
			}

			pvcClient := clientset.CoreV1().PersistentVolumeClaims(apiv1.NamespaceDefault)

			result, err := pvcClient.Create(pvc)
			if errors.IsAlreadyExists(err) {
				fmt.Printf("Already exists PVC %q.\n", pvc.Name)
				continue
			}
			if err != nil {
				panic(err)
			}
			fmt.Printf("Created PersistentVolumeCalim %q.\n", result.Name)
		default:
			data, err := decoder.YAML()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Unkown Object: %s", string(data))
		}
	}
}

func CreateMariadbSvc(clientset kubernetes.Interface) {
	recipe := svc.NewMariadbGaleraRecipient("mariadb-galera", "default", "mariadb-galera")
	service, err := recipe.Create()
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

func CreateMariadbSts(clientset kubernetes.Interface) {
	count := uint8(3)
	recipe := sts.NewMariadbGaleraRecipient("mariadb-galera", "", "mariadb-galera", "", &count, nil)
	statefulset, err := recipe.Create()
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
