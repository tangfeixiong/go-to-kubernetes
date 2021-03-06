// Package agent to manage Kubernetes storage attach events.
package agent

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/coreos/pkg/capnslog"
	rookalpha "github.com/rook/rook/pkg/apis/rook.io/v1alpha1"
	"github.com/rook/rook/pkg/daemon/agent/flexvolume/attachment"
	opcluster "github.com/rook/rook/pkg/operator/cluster"
	"github.com/rook/rook/pkg/operator/k8sutil"
	"k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/api/rbac/v1beta1"
	kserrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	agentDaemonsetName          = "rook-agent"
	flexvolumePathDirEnv        = "FLEXVOLUME_DIR_PATH"
	flexvolumeDefaultDirPath    = "/usr/libexec/kubernetes/kubelet-plugins/volume/exec/"
	agentDaemonsetTolerationEnv = "AGENT_TOLERATION"
)

var logger = capnslog.NewPackageLogger("github.com/rook/rook", "op-agent")

var clusterAccessRules = []v1beta1.PolicyRule{
	{
		APIGroups: []string{""},
		Resources: []string{"pods", "secrets", "configmaps", "persistentvolumes", "nodes", "nodes/proxy"},
		Verbs:     []string{"get", "list"},
	},
	{
		APIGroups: []string{rookalpha.CustomResourceGroup},
		Resources: []string{attachment.CustomResourceNamePlural},
		Verbs:     []string{"get", "list", "create", "delete", "update"},
	},
	{
		APIGroups: []string{rookalpha.CustomResourceGroup},
		Resources: []string{opcluster.CustomResourceNamePlural},
		Verbs:     []string{"get", "list", "watch"},
	},
	{
		APIGroups: []string{"storage.k8s.io"},
		Resources: []string{"storageclasses"},
		Verbs:     []string{"get"},
	},
}

// New creates an instance of Agent
func New(clientset kubernetes.Interface) *Agent {
	return &Agent{
		clientset: clientset,
	}
}

// Start the agent
func (a *Agent) Start(namespace, agentImage string) error {

	err := k8sutil.MakeClusterRole(a.clientset, namespace, agentDaemonsetName, clusterAccessRules)
	if err != nil {
		return fmt.Errorf("failed to init RBAC for rook-agents. %+v", err)
	}

	err = a.createAgentDaemonSet(namespace, agentImage)
	if err != nil {
		return fmt.Errorf("Error starting agent daemonset: %v", err)
	}
	return nil
}

func (a *Agent) createAgentDaemonSet(namespace, agentImage string) error {

	flexvolumeDirPath, source := a.discoverFlexvolumeDir()
	logger.Infof("discovered flexvolume dir path from source %s. value: %s", source, flexvolumeDirPath)

	privileged := true
	ds := &extensions.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: agentDaemonsetName,
		},
		Spec: extensions.DaemonSetSpec{
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": agentDaemonsetName,
					},
				},
				Spec: v1.PodSpec{
					ServiceAccountName: agentDaemonsetName,
					Containers: []v1.Container{
						{
							Name:  agentDaemonsetName,
							Image: agentImage,
							Args:  []string{"agent"},
							SecurityContext: &v1.SecurityContext{
								Privileged: &privileged,
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "flexvolume",
									MountPath: "/flexmnt",
								},
								{
									Name:      "dev",
									MountPath: "/dev",
								},
								{
									Name:      "sys",
									MountPath: "/sys",
								},
								{
									Name:      "libmodules",
									MountPath: "/lib/modules",
								},
							},
							Env: []v1.EnvVar{
								k8sutil.NamespaceEnvVar(),
								k8sutil.NodeEnvVar(),
							},
						},
					},
					Volumes: []v1.Volume{
						{
							Name: "flexvolume",
							VolumeSource: v1.VolumeSource{
								HostPath: &v1.HostPathVolumeSource{
									Path: flexvolumeDirPath,
								},
							},
						},
						{
							Name: "dev",
							VolumeSource: v1.VolumeSource{
								HostPath: &v1.HostPathVolumeSource{
									Path: "/dev",
								},
							},
						},
						{
							Name: "sys",
							VolumeSource: v1.VolumeSource{
								HostPath: &v1.HostPathVolumeSource{
									Path: "/sys",
								},
							},
						},
						{
							Name: "libmodules",
							VolumeSource: v1.VolumeSource{
								HostPath: &v1.HostPathVolumeSource{
									Path: "/lib/modules",
								},
							},
						},
					},
					HostNetwork: true,
				},
			},
		},
	}

	// Add toleration if any
	tolerationValue := os.Getenv(agentDaemonsetTolerationEnv)
	if tolerationValue != "" {
		ds.Spec.Template.Spec.Tolerations = []v1.Toleration{
			{
				Effect:   v1.TaintEffect(tolerationValue),
				Operator: v1.TolerationOpExists,
			},
		}
	}

	_, err := a.clientset.Extensions().DaemonSets(namespace).Create(ds)
	if err != nil {
		if !kserrors.IsAlreadyExists(err) {
			return fmt.Errorf("failed to create rook-agent daemon set. %+v", err)
		}
		logger.Infof("rook-agent daemonset already exists")
	} else {
		logger.Infof("rook-agent daemonset started")
	}
	return nil

}

func (a *Agent) discoverFlexvolumeDir() (flexvolumeDirPath, source string) {
	//copy flexvolume to flexvolume dir
	nodeName := os.Getenv(k8sutil.NodeNameEnvVar)
	if nodeName == "" {
		logger.Warningf("cannot detect the node name. Please provide using the downward API in the rook operator manifest file")
		return getDefaultFlexvolumeDir()
	}

	// determining where the path of the flexvolume dir on the node
	nodeConfigURI, err := k8sutil.NodeConfigURI()
	if err != nil {
		logger.Warningf(err.Error())
		return getDefaultFlexvolumeDir()
	}
	nodeConfig, err := a.clientset.Core().RESTClient().Get().RequestURI(nodeConfigURI).DoRaw()
	if err != nil {
		logger.Warningf("unable to query node configuration: %v", err)
		return getDefaultFlexvolumeDir()
	}

	// unmarshal to a NodeConfigKubelet
	configKubelet := NodeConfigKubelet{}
	if err := json.Unmarshal(nodeConfig, &configKubelet); err != nil {
		logger.Warningf("unable to parse node config from Kubelet: %+v", err)
	} else {
		flexvolumeDirPath = configKubelet.ComponentConfig.VolumePluginDir
	}

	if flexvolumeDirPath != "" {
		return flexvolumeDirPath, "NodeConfigKubelet"
	}

	// unmarshal to a NodeConfigControllerManager
	configControllerManager := NodeConfigControllerManager{}
	if err := json.Unmarshal(nodeConfig, &configControllerManager); err != nil {
		logger.Warningf("unable to parse node config from controller manager: %+v", err)
	} else {
		flexvolumeDirPath = configControllerManager.ComponentConfig.VolumeConfiguration.FlexVolumePluginDir
	}

	if flexvolumeDirPath != "" {
		return flexvolumeDirPath, "NodeConfigControllerManager"
	}

	// unmarshal to a KubeletConfiguration
	kubeletConfiguration := KubeletConfiguration{}
	if err := json.Unmarshal(nodeConfig, &kubeletConfiguration); err != nil {
		logger.Warningf("unable to parse node config as kubelet configuration: %+v", err)
	} else {
		flexvolumeDirPath = kubeletConfiguration.KubeletConfig.VolumePluginDir
	}

	if flexvolumeDirPath != "" {
		return flexvolumeDirPath, "KubeletConfiguration"
	}

	return getDefaultFlexvolumeDir()
}

func getDefaultFlexvolumeDir() (flexvolumeDirPath, source string) {
	logger.Infof("getting flexvolume dir path from %s env var", flexvolumePathDirEnv)
	flexvolumeDirPath = os.Getenv(flexvolumePathDirEnv)
	if flexvolumeDirPath != "" {
		return flexvolumeDirPath, "env var"
	}

	logger.Infof("flexvolume dir path env var %s is not provided. Defaulting to: %s",
		flexvolumePathDirEnv, flexvolumeDefaultDirPath)
	flexvolumeDirPath = flexvolumeDefaultDirPath

	return flexvolumeDirPath, "default"
}
