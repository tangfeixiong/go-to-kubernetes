package preboot

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	//"strconv"
	"strings"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
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
	commonutil "github.com/tangfeixiong/go-to-kubernetes/pkg/util"
)

type ClusterType string

const (
	MariadbGalera ClusterType = "galera"
	MySQLGPC      ClusterType = "mysql"

	DefaultDomainName = "svc.cluster.local"
	MariadbConfDir    = "/etc/mysql/mariadb.conf.d"
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

func PreBootGalera(cfg *Config) {
	if cfg.Namespace == "" {
		cfg.Namespace = os.Getenv("MY_POD_NAMESPACE")
	}
	if cfg.Dir == "" {
		cfg.Dir = MariadbConfDir
	}
	if cfg.DomainName == "" {
		cfg.DomainName = DefaultDomainName
	}

	//clientset, err := k8sClientset(cfg.Kubeconfig)
	clientset, _, err := commonutil.K8sClientset(cfg.Kubeconfig)
	if err != nil {
		glog.Fatal(err.Error())
	}

	glog.V(5).Infof("With parameters: %v\n", cfg)
	//glog.Infof("With parameters: %v\n", cfg)
	cfg.CreateGaleraCnf(clientset)

	copies := []string{"/tini", "/mysql-operator"}
	for _, src := range copies {
		dest := filepath.Join("/operator-entrypoint/", src)
		if fi, err := os.Stat(dest); err != nil {
			if os.IsNotExist(err) {
				in, err := os.Open(src)
				if err != nil {
					glog.Fatalf("Open %s failed: %s", src, err)
				}
				defer in.Close()

				out, err := os.Create(dest)
				if err != nil {
					glog.Fatalf("Create %s failed: %s", dest, err.Error())
				}
				defer out.Close()

				if _, err := io.Copy(out, in); err != nil {
					glog.Fatalf("Write %s failed: %s", dest, err)
				}

				if err := os.Chmod(dest, 0755); err != nil {
					glog.Fatal(err)
				}
			} else {
				glog.Fatalf("Stats %s failed: %s. %q", dest, err.Error(), fi)
			}
		}
	}
}

func StartMysqldAgent(cfg *Config) {
	ns := corev1.NamespaceAll
	if cfg.Namespace == "" {
		cfg.Namespace = os.Getenv("MY_POD_NAMESPACE")
	}
	ns = cfg.Namespace
	if ns == "" {
		ns = corev1.NamespaceDefault
	}
	if cfg.Dir == "" {
		cfg.Dir = MariadbConfDir
	}
	if cfg.DomainName == "" {
		cfg.DomainName = DefaultDomainName
	}
	clientset, _, err := commonutil.K8sClientset(cfg.Kubeconfig)
	if err != nil {
		glog.Fatal(err.Error())
	}

	stsClient := clientset.AppsV1beta2().StatefulSets(ns)

	mypodname, ok := os.LookupEnv("MY_POD_NAME")
	if cfg.Name == "" {
		if !ok || mypodname == "" {
			glog.Exitln("Read Pod name environment failed")
		}
		pos := strings.LastIndex(mypodname, "-")
		if pos == -1 {
			glog.Exitln("Read Pod name format failed")
		}
		cfg.Name = mypodname[:pos]
	}

	statefulSet, err := stsClient.Get(cfg.Name, metav1.GetOptions{})
	if err != nil {
		glog.Exitf("Reap statefulset failed: %s\n", err.Error())
	}

	selector := statefulSet.Spec.Selector
	if selector == nil {
		// If empty, defaulted to labels on the pod template.
		selector = &metav1.LabelSelector{
			MatchLabels: statefulSet.Spec.Template.Labels,
		}
	}

	podClient := clientset.CoreV1().Pods(ns)

	pods, err := podClient.List(metav1.ListOptions{
		LabelSelector: metav1.FormatLabelSelector(selector),
	})
	if err != nil {
		glog.Exit("List pods failed:", err.Error())
	}

	gcomm := []string{}
	for i := 0; i < len(pods.Items); i++ {
		//addr := fmt.Sprintf("%s-%d.%s.%s.%s", statefulSet.Name, i, statefulSet.Name, statefulSet.Namespace, cfg.DomainName)
		pod := pods.Items[i]
		if pod.Status.PodIP != "" && pod.Name != mypodname {
			// Create the database handle, confirm driver is present
			username := "root"
			password := "ROOT_PASS"
			host := pod.Status.PodIP
			port := 3306
			timeout := "30s"
			db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/?timeout=%s", username, password, host, port, timeout))
			defer db.Close()

			var fieldName, wsrepStatus string
			// SHOW GLOBAL STATUS LIKE 'wsrep_ready';
			err := db.QueryRow("SHOW GLOBAL STATUS LIKE 'wsrep_connected';").Scan(&fieldName, &wsrepStatus)
			if err != nil {
				glog.Exitf("Check galera failed: %s", err.Error())
			}
			if wsrepStatus == "ON" {
				gcomm = append(gcomm, pod.Status.PodIP)
				break
			}
		}
	}
	if len(gcomm) == 0 && len(pods.Items) > 1 {
		glog.Exitf("Determine GCOMM failed")
	}

	/*
	   $ mysqld --wsrep_new_cluster
	   $ mysqld --wsrep_cluster_address=gcomm://192.168.0.1  # DNS names work as well
	*/
	var bin string
	var args []string
	if len(gcomm) == 0 {
		bin = "sh"
		args = []string{"sh", "-c", "'/usr/local/bin/docker-entrypoint.sh mysqld --wsrep_new_cluster'"}
	} else {
		bin = "mysqld"
		args = []string{"mysqld", "--wsrep_cluster_address=gcomm://" + gcomm[0]}
	}
	binary, lookErr := exec.LookPath(bin)
	if lookErr != nil {
		glog.Fatalf("Locate bin failed: %s", lookErr)
	}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		glog.Fatalf("Boot mysqld failed: %s", execErr)
	}
}

func (cfg *Config) CreateGaleraCnf(clientset kubernetes.Interface) {
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
