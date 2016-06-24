package main

import (
	"fmt"
	"log"

	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	"k8s.io/kubernetes/pkg/runtime"
)

var (
	kubeconfig = "/Users/fanhongling/Downloads/github.com/openshift/origin/kubeconfig"
)

func main() {
    kubeconfig = "/home/ubuntu/.kube/config"
    
	// k8s.io/kubernetes/pkg/client/unversioned/clientcmd/loader.go
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// k8s.io/kubernetes/pkg/client/unversioned/clientcmd/client_config.go
	clientConfig := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{})

    restConfig, err := clientConfig.ClientConfig()
    if err != nil {
        log.Fatal(err)
    }
    
	// k8s.io/kubernetes/pkg/client/unversioned/helper.go
	kclient, err := unversioned.New(restConfig)
	if err != nil {
		log.Fatal(err)
	}

	// k8s.io/kubernetes/pkg/client/unversioned/client.go
	// k8s.io/kubernetes/pkg/client/unversioned/pods.go
	result, err := kclient.Pods("default").List(kapi.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// k8s.io/kubernetes/pkg/runtime/codec.go
	// k8s.io/kubernetes/pkg/api/register.go
	// k8s.io/kubernetes/pkg/runtime/serializer/codec_factory.go
	b, err := runtime.Encode(kapi.Codecs.LegacyCodec(kapi.SchemeGroupVersion), result, kapi.SchemeGroupVersion)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output list:\n%s\n", b)
}
