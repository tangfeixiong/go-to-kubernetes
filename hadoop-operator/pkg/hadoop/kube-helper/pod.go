/*
	"github.com/rook/rook/pkg/operator/k8sutil"
*/
package kubehelper

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetContainerImage(clientset kubernetes.Interface) (string, error) {

	podName := os.Getenv(PodNameEnvVar)
	if podName == "" {
		return "", fmt.Errorf("cannot detect the pod name. Please provide it using the downward API in the manifest file")
	}
	podNamespace := os.Getenv(PodNamespaceEnvVar)
	if podName == "" {
		return "", fmt.Errorf("cannot detect the pod namespace. Please provide it using the downward API in the manifest file")
	}

	pod, err := clientset.Core().Pods(podNamespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	if len(pod.Spec.Containers) != 1 {
		return "", fmt.Errorf("failed to get container image. There should only be exactly one container in this pod")
	}

	return pod.Spec.Containers[0].Image, nil
}
