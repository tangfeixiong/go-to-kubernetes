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

	DefaultDomainName = "svc.cluster.local"
)

type Config struct {
	Kubeconfig string
	Kind       string
	Name       string
	Namespace  string
	Clustering ClusterType
	Port       int
	Dir        string
	DomainName string
	LogLevel   int
}

func BuildGaleraCnf(cfg *Config) {
	if cfg.Namespace == "" {
		cfg.Namespace = os.Getenv("MY_POD_NAMESPACE")
	}
	if cfg.Dir == "" {
		cfg.Dir = "/etc/mysql/mariadb.conf.d"
	}
	if cfg.DomainName == "" {
		cfg.DomainName = DefaultDomainName
	}

	clientset, err := k8sClientset(cfg.Kubeconfig)
	if err != nil {
		glog.Fatal("Get kubernetes clientset failed:", err.Error())
	}

	//glog.V(5).Infof("With init parameters: %v\n", cfg)
	glog.Infof("With init parameters: %v\n", cfg)
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
			glog.Exitf("Get statefulset failed: %s\n", err.Error())
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
			glog.Exitf("Get statefulset failed: %s\n", err.Error())
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
	if len(list.Items) > 0 {
		//mypodname, ok := os.LookupEnv("MY_POD_NAME")
		if mypodname == "" || !ok {
			glog.Error("Get POD name failed or Environment POD name not found")
		}
		myname := mypodname
		// myname, err := os.Hostname()
		myip := ""
		//		for _, pod := range list.Items {
		//			if pod.Name == myname {
		//				myip = pod.Status.PodIP
		//				break
		//			}
		//		}
		myip, ok = os.LookupEnv("MY_POD_IP")
		if myip == "" || !ok {
			glog.Error("Get POD IP failed or environment POD IP not found")
		}

		for i := 0; i < len(list.Items); i++ {
			//addr := fmt.Sprintf("%s-%d.%s.%s.%s", obj.Name, i, obj.Name, obj.Namespace, cfg.DomainName)
			pod := list.Items[i]
			if pod.Status.PodIP != "" && pod.Status.PodIP != myip {
				gcomm = append(gcomm, pod.Status.PodIP)
			}
		}

		var recipe *galera.Recipient
		var err error
		if len(gcomm) == 0 {
			recipe, err = galera.NewRecipient(false, galera.ClusterNameSetter(obj.Name),
				galera.WsrepProviderOptionsSetter("pc.npvo=TRUE", "pc.recovery=FALSE"),
				galera.ThisNodeSetter(myip, myname))
		} else {
			recipe, err = galera.NewRecipient(false, galera.ClusterNameSetter(obj.Name),
				galera.ClusterAddressesSetter(gcomm...),
				galera.WsrepProviderOptionsSetter("pc.wait_prim_timeout=PT300S", "pc.npvo=TRUE", "pc.recovery=FALSE"),
				galera.ThisNodeSetter(myip, myname))
		}
		if err != nil {
			glog.Error("New recipe failed:", err)
		} else if err = recipe.Generate(path); err != nil {
			glog.Fatal("Write galera.cnf failed:", err)
		} else {
			glog.Infof("Bootstrap: %v", gcomm)
		}
		return
	}

	for i := 0; i < int(*obj.Spec.Replicas); i++ {
		addr := fmt.Sprintf("%s-%d.%s.%s.%s", obj.Name, i, obj.Name, obj.Namespace, cfg.DomainName)
		gcomm = append(gcomm, addr)
	}
	if recipe, err := galera.NewRecipient(false, galera.ClusterNameSetter(obj.Name),
		galera.ClusterAddressesSetter(gcomm...), galera.WsrepProviderOptionsSetter("pc.wait_prim=FALSE")); err != nil {
		glog.Error("New recipe failed:", err)
	} else if err := recipe.Generate(path); err != nil {
		glog.Fatal("Write galera.cnf failed:", err)
	} else {
		glog.Infoln("Bootstrap without thisNodeAddress and thisNodeName")
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
