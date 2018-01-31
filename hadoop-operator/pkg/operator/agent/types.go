// Package agent to manage Kubernetes storage attach events.
package agent

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/apis/componentconfig"
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1"
)

// Agent reference to be deployed
type Agent struct {
	clientset kubernetes.Interface
}

// NodeConfigControllerManager is a reference of all the configuration for the K8S node from the controllermanager
type NodeConfigControllerManager struct {
	ComponentConfig componentconfig.KubeControllerManagerConfiguration `json:"componentconfig"`
}

// NodeConfigKubelet is a reference of all the configuration for the K8S node from kubelet
type NodeConfigKubelet struct {
	ComponentConfig kubeletconfig.KubeletConfiguration `json:"componentconfig"`
}

// KubeletConfiguration represents the response from the node config URI (configz) in Kubernetes 1.8+
type KubeletConfiguration struct {
	KubeletConfig struct {
		VolumePluginDir string `json:"volumePluginDir"`
	} `json:"kubeletconfig"`
}
