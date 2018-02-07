package hadoop

import (
	//"database/sql"
	"fmt"
	"os"
	//"path/filepath"
	//"strconv"
	"strings"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"

	appsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	//appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"

	sampleapiv1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	hadoopspec "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/spec/hadoop"
	"github.com/tangfeixiong/go-to-kubernetes/pkg/util"
)

type ConfigMode string

const (
	ConfigMap     ConfigMode = "configMap"
	InitContainer ConfigMode = "initContainer"

	cluster_example = "my-hdfs"
	domain_example  = "cluster.local"
)

type Config struct {
	mode               string
	Kubeconfig         string
	Kind               string
	CustomResourceName string
	Name               string
	resourceHostname   string
	ServiceName        string
	Namespace          string
	BaseDomain         string
	Port               int
	Dir                string
	ClusterID          string
	NodeType           string
	LogLevel           int
}

func (cfg *Config) prerequisites() (kubernetes.Interface, error) {
	if cfg.resourceHostname == "" {
		cfg.resourceHostname = os.Getenv("MY_POD_NAME")
	}
	if cfg.Namespace == "" {
		cfg.Namespace = os.Getenv("MY_POD_NAMESPACE")
	}
	if cfg.Dir == "" {
		cfg.Dir = "/hadoop-3.0.0"
	}
	if cfg.BaseDomain == "" {
		cfg.BaseDomain = domain_example
	}

	if cfg.NodeType == "" {
		cfg.NodeType = string(sampleapiv1alpha1.DataNode)
	}

	clientset, _, err := util.K8sClientset(cfg.Kubeconfig)
	return clientset, err
}

func PreBoot(cfg *Config) {
	clientset, err := cfg.prerequisites()
	if err != nil {
		glog.Fatal(err)
	}

	glog.V(5).Infof("With init args: %v\n", cfg)
	cfg.initContainer(clientset)
}

func InitMap(cfg *Config) error {
	clientset, err := cfg.prerequisites()
	if err != nil {
		return err
	}
	if cfg.Namespace == "" {
		cfg.Namespace = corev1.NamespaceDefault
	}

	glog.Infof("With config parameters: %v\n", cfg)
	_, err = cfg.InitMap(clientset)
	return err
}

func (cfg *Config) InitMap(clientset kubernetes.Interface) (*corev1.ConfigMap, error) {
	cfg.resourceHostname = fmt.Sprintf("%s-0", cfg.Name)
	recipe, err := hadoopspec.NewHadoopRecipient(
		hadoopspec.NameNodeFqdnSetter(cfg.resourceHostname, cfg.ServiceName, cfg.Namespace, "svc", cfg.BaseDomain),
		hadoopspec.ResourceNameSetter(cfg.Name),
		hadoopspec.CustomResourceNameSetter(cfg.CustomResourceName))
	if err != nil {
		return nil, fmt.Errorf("Create recipe failed:", err)
	}

	cm, err := recipe.New()
	if err != nil {
		return nil, fmt.Errorf("New configmap failed: %v", err)
	}

	configMapClient := clientset.CoreV1().ConfigMaps(cfg.Namespace)

	result, err := configMapClient.Create(cm)
	if errors.IsAlreadyExists(err) {
		glog.V(5).Infof("Already exists CM %q.\n", cm.Name)
		result, err = configMapClient.Update(cm)
	}
	if err != nil {
		return nil, err
	}
	glog.V(5).Infof("Created config map %q.\n", result.GetObjectMeta().GetName())

	return result, nil
}

func (cfg *Config) initContainer(clientset kubernetes.Interface) {
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
		lbs := labels.Set{"app": "hadoop", "github.com/go-to-kubernetes": "hadoop-operator"}
		ls := metav1.SetAsLabelSelector(lbs)
		opts := metav1.ListOptions{
			LabelSelector: metav1.FormatLabelSelector(ls),
		}
		var list *appsv1beta2.StatefulSetList
		list, err = stsClient.List(opts)
		if err != nil {
			glog.Exitf("Reading statefulset failed: %s\n", err.Error())
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
			glog.Exitf("Reading statefulset failed: %s\n", err.Error())
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

	if len(list.Items) != 0 {
		//mypodname, ok := os.LookupEnv("MY_POD_NAME")
		if mypodname == "" || !ok {
			glog.Error("Read POD name failed or Environment POD name not found")
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
			glog.Error("Read POD IP failed or environment POD IP not found")
		}

		var xmlrecipe *hadoopspec.HadoopRecipient
		cfg.resourceHostname = fmt.Sprintf("%s-0", obj.Name)
		xmlrecipe, err = hadoopspec.NewHadoopRecipient(
			hadoopspec.NameNodeFqdnSetter(cfg.resourceHostname, obj.Spec.ServiceName, obj.Namespace, "svc", cfg.BaseDomain),
			hadoopspec.ResourceNameSetter(obj.Name),
			hadoopspec.CustomResourceNameSetter(obj.ObjectMeta.Labels["example.com/hadoop-operator"]))
		if err != nil {
			glog.Fatal("Create recipe failed:", err)
		}

		err = xmlrecipe.Generate("/operator-entrypoint")
		if err != nil {
			glog.Fatal("Generate files failed: %v", err)
		}

		var recipe *hadoopspec.HadoopStartupRecipient
		node := string(sampleapiv1alpha1.DataNode)
		if strings.HasSuffix(mypodname, "-0") {
			node = string(sampleapiv1alpha1.NameNode)
		} else if strings.HasSuffix(mypodname, "-1") {
			node = string(sampleapiv1alpha1.SecondaryNameNode)
		}
		recipe, err = hadoopspec.NewHadoopStartupRecipient(hadoopspec.HadoopHomeSetter(cfg.Dir),
			hadoopspec.BinNameSetter("hdfs"), hadoopspec.CmdNameSetter(node))
		if err != nil {
			glog.Error("New recipe failed:", err)
		}
		if err = recipe.Generate("/operator-entrypoint"); err != nil {
			glog.Fatal("Write bootstrap files failed:", err)
		}

		if node == string(sampleapiv1alpha1.NameNode) {
			err = RealRunner{}.Format("/hadoop-3.0.0/bin/hadoop", node, cfg.ClusterID)
			if err != nil {
				glog.Exitf("Format DFS failed: %s", err.Error())
			}
		}
		return
	}

	glog.Infoln("No POD found")
}
