package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var logger *log.Logger = log.New(os.Stderr, fmt.Sprintf("[%s] ", filepath.Base(os.Args[0])), log.LstdFlags|log.Lshortfile)

/*
  Inspired by:
    https://github.com/kubernetes/client-go/blob/master/util/homedir/homedir.go
*/
// HomeDir returns the home directory for the current user
func HomeDir() string {
	if runtime.GOOS == "windows" {

		// First prefer the HOME environmental variable
		if home := os.Getenv("HOME"); len(home) > 0 {
			if _, err := os.Stat(home); err == nil {
				return home
			}
		}
		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			homeDir := homeDrive + homePath
			if _, err := os.Stat(homeDir); err == nil {
				return homeDir
			}
		}
		if userProfile := os.Getenv("USERPROFILE"); len(userProfile) > 0 {
			if _, err := os.Stat(userProfile); err == nil {
				return userProfile
			}
		}
	}
	return os.Getenv("HOME")
}

func K8sClientset(kubeconfig string) (kubernetes.Interface, *rest.Config, error) {
	var config *rest.Config
	var err error
	var clientset *kubernetes.Clientset

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			//panic(err)
			logger.Println("Read Kubernetes config failed:", err.Error())
			return nil, nil, err
		}
	} else {
		// creates the in-cluster config
		config, err = rest.InClusterConfig()
		if err != nil {
			// panic(err.Error())
			logger.Println("Read Kubernetes in-cluster config failed:", err.Error())
			return nil, nil, err
		}
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		// panic(err.Error())
		logger.Println("Fetch Kubernetes clientset failed:", err.Error())
		return nil, config, err
	}
	return clientset, config, nil
}
