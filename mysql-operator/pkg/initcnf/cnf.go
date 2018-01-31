package initcnf

import (
	//"database/sql"
	"fmt"
	"os"
	"path/filepath"
	//"strconv"
	"strings"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"

	appsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	//appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	//sampleapis "github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/apis/example.com"
	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/galera"
)

type ClusterType string

const (
	MariadbGalera ClusterType = "galera"
	MySQLGPC      ClusterType = "mysql"
)

type Config struct {
	Kubeconfig string
	Kind       string
	Name       string
	Namespace  string
	Clustering ClusterType
	Port       int
	Dir        string
	LogLevel   int
}

func BuildGaleraCnf(cfg *Config) {
	if cfg.Namespace == "" {
		cfg.Namespace = os.Getenv("MY_POD_NAMESPACE")
	}
	if cfg.Dir == "" {
		cfg.Dir = "/etc/mysql/conf.d"
	}

	clientset, err := k8sClientset(cfg.Kubeconfig)
	if err != nil {
		glog.Fatal("Get kubernetes clientset failed:", err.Error())
	}

	glog.V(5).Infof("With init parameters: %v\n", cfg)
	cfg.buildGaleraCnf(clientset)
}

func (cfg *Config) buildGaleraCnf(clientset kubernetes.Interface) {
	var ns string = corev1.NamespaceAll
	if cfg.Namespace == "" {
		ns = corev1.NamespaceDefault
	} else {
		ns = cfg.Namespace
	}

	stsClient := clientset.AppsV1beta2().StatefulSets(ns)

	var obj *appsv1beta2.StatefulSet = nil
	var err error

	mypodname, ok := os.LookupEnv("MY_POD_NAME")
	if cfg.Name == "" {
		if !ok {
			glog.Error("Environment POD name not found")
		}
		lbs := labels.Set{"app": "mariadb", "component": "mariadb-galera"}
		ls := metav1.SetAsLabelSelector(lbs)
		opts := metav1.ListOptions{
			LabelSelector: metav1.FormatLabelSelector(ls),
		}
		var list *appsv1beta2.StatefulSetList
		list, err = stsClient.List(opts)
		if err != nil {
			glog.Exit("Get statefulset failed: %+v\n", err)
		}
		for _, sts := range list.Items {
			if strings.HasPrefix(mypodname, sts.Name) && strings.HasPrefix(mypodname[len(sts.Name):], "-") {
				obj = &sts
				break
			}
		}
	} else {
		obj, err = stsClient.Get(cfg.Name, metav1.GetOptions{})
		if err != nil {
			glog.Exitf("Get statefulset failed: %+v\n", err)
		}
	}
	if obj == nil {
		glog.Exit("Object not found")
	}

	selector := obj.Spec.Selector
	if selector == nil {
		// If empty, defaulted to labels on the pod template.
		selector = &metav1.LabelSelector{
			MatchLabels: obj.Spec.Template.Labels,
		}
	}

	podClient := clientset.CoreV1().Pods(ns)

	list, err := podClient.List(metav1.ListOptions{
		LabelSelector: metav1.FormatLabelSelector(selector),
	})
	if err != nil {
		glog.Exit("List pods failed:", err.Error())
	}

	path := filepath.Join(cfg.Dir, "galera.cnf")

	gcomm := []string{}
	for i := 0; i < int(*obj.Spec.Replicas); i++ {
		addr := fmt.Sprintf("%s-%d.%s.%s", obj.Name, i, obj.Name, obj.Namespace)
		gcomm = append(gcomm, addr)
	}

	if len(list.Items) > 0 {
		//mypodname, ok := os.LookupEnv("MY_POD_NAME")
		if !ok {
			glog.Error("Environment POD name not found")
		}
		myname := mypodname
		myip := ""
		//		for _, pod := range list.Items {
		//			if pod.Name == myname {
		//				myip = pod.Status.PodIP
		//				break
		//			}
		//		}
		myip, ok = os.LookupEnv("MY_POD_NAME")
		if myip == "" || !ok {
			glog.Error("Get POD IP failed or environment POD IP not found")
		}
		for _, pod := range list.Items {
			if pod.Name != myname {
				if recipe, err := galera.NewRecipient(false, galera.ClusterNameSetter(obj.Name), galera.ClusterAddressesSetter(pod.Status.PodIP), galera.ThisNodeSetter(myip, myname)); err != nil {
					glog.Error("New recipe failed:", err)
				} else if err := recipe.Generate(path); err != nil {
					glog.Fatal("Write galera.cnf failed:", err)
				}
				return
			}
		}
		if recipe, err := galera.NewRecipient(false, galera.ClusterNameSetter(obj.Name), galera.ClusterAddressesSetter(gcomm...), galera.ThisNodeSetter(myip, myname)); err != nil {
			glog.Error("New recipe failed:", err)
		} else if err := recipe.Generate(path); err != nil {
			glog.Fatal("Write galera.cnf failed:", err)
		}
		return
	}

	if recipe, err := galera.NewRecipient(false, galera.ClusterNameSetter(obj.Name), galera.ClusterAddressesSetter(gcomm...)); err != nil {
		glog.Error("New recipe failed:", err)
	} else if err := recipe.Generate(path); err != nil {
		glog.Fatal("Write galera.cnf failed:", err)
	}
}

func k8sClientset(kubeconfig string) (kubernetes.Interface, error) {
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
