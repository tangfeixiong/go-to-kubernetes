package operator

import (
	"github.com/golang/glog"

	//"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func K8sClientset(kubeconfig string) (kubernetes.Interface, error) {
	var config *rest.Config
	var err error
	var clientset *kubernetes.Clientset

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			//panic(err)
			glog.V(2).Infoln("Could not get Kubernetes config:", err.Error())
			return nil, err
		}
	} else {
		// creates the in-cluster config
		config, err = rest.InClusterConfig()
		if err != nil {
			// panic(err.Error())
			glog.V(2).Infoln("Could not get Kubernetes config:", err.Error())
			return nil, err
		}
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		// panic(err.Error())
		glog.V(2).Infoln("Could not get Kubernetes clientset:", err.Error())
		return nil, err
	}
	return clientset, nil
}
