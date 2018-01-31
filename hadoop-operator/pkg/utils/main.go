/*
	"https://github.com/rook/rook/blob/master/cmd/rook/main.go"
*/

package utils

import (
	"fmt"

	rookclient "github.com/rook/rook/pkg/client/clientset/versioned"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getClientset() (kubernetes.Interface, apiextensionsclient.Interface, rookclient.Interface, error) {
	// create the k8s client
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get k8s config. %+v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create k8s clientset. %+v", err)
	}
	apiExtClientset, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create k8s API extension clientset. %+v", err)
	}
	rookClientset, err := rookclient.NewForConfig(config)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create rook clientset. %+v", err)
	}
	return clientset, apiExtClientset, rookClientset, nil
}
