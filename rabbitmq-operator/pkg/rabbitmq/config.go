package rabbitmq

import (
	//"database/sql"
	//"fmt"
	"os"
	//"path/filepath"
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

	"github.com/tangfeixiong/go-to-kubernetes/pkg/util"
	sampleapiv1alpha1 "github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/apis/example.com/v1alpha1"
	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/spec/rabbitmq"
)

type ClusteringMode string

const (
	//PeerDiscovery_K8s ClusteringMode = "Peer_Discovery_K8s"
	//Native  ClusteringMode = "Native"
	//None    ClusteringMode = "None"

	DefaultDomainName = "cluster.local"
)

type Config struct {
	Kubeconfig string
	Kind       string
	Name       string
	Namespace  string
	Mode       string
	Port       int
	Dir        string
	DomainName string
	LogLevel   int
}

func BuildConfigFile(cfg *Config) {
	if cfg.Namespace == "" {
		cfg.Namespace = os.Getenv("MY_POD_NAMESPACE")
	}
	if cfg.Dir == "" {
		cfg.Dir = "/etc/rabbitmq"
	}
	if cfg.DomainName == "" {
		cfg.DomainName = DefaultDomainName
	}
	if cfg.Mode == "" {
		cfg.Mode = string(sampleapiv1alpha1.PeerDiscovery_K8s)
	}

	clientset, _, err := util.K8sClientset(cfg.Kubeconfig)
	if err != nil {
		glog.Fatal(err)
	}

	//glog.V(5).Infof("With init parameters: %v\n", cfg)
	glog.Infof("With init parameters: %v\n", cfg)
	cfg.buildConfigFile(clientset)
}

func (cfg *Config) buildConfigFile(clientset kubernetes.Interface) {
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
		lbs := labels.Set{"app": "rabbitmq", "github.com/go-to-kubernetes": "rabbitmq-operator"}
		ls := metav1.SetAsLabelSelector(lbs)
		opts := metav1.ListOptions{
			LabelSelector: metav1.FormatLabelSelector(ls),
		}
		var list *appsv1beta2.StatefulSetList
		list, err = stsClient.List(opts)
		if err != nil {
			glog.Exitf("Read statefulset failed: %s\n", err.Error())
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
			glog.Exitf("Read statefulset failed: %s\n", err.Error())
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

	if len(list.Items) > 0 {
		//mypodname, ok := os.LookupEnv("MY_POD_NAME")
		if mypodname == "" || !ok {
			glog.Error("Get POD name failed or Environment POD name not found")
		}
		// mypodname, err := os.Hostname()

		mypodip := ""
		//		for _, pod := range list.Items {
		//			if pod.Name == mypodname {
		//				mypodip = pod.Status.PodIP
		//				break
		//			}
		//		}
		mypodip, ok = os.LookupEnv("MY_POD_IP")
		if mypodip == "" || !ok {
			glog.Error("Get POD IP failed or environment POD IP not found")
		}

		var recipe *rabbitmq.RabbitmqRecipient
		var err error
		recipe, err = rabbitmq.NewRabbitmqRecipient(rabbitmq.ClusteringModeSetter(cfg.Mode),
			rabbitmq.NamespaceSetter(obj.Namespace),
			rabbitmq.TargetServiceSetter(obj.Spec.ServiceName),
			rabbitmq.DomainNameSetter(cfg.DomainName))
		if err != nil {
			glog.Error("New recipe failed:", err)
		}
		if err = recipe.Generate(cfg.Dir); err != nil {
			glog.Fatal("Write bootstrap files failed:", err)
		}
		return
	}
}
