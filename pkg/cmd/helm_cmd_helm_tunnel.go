package cmd

// k8s.io/helm/cmd/helm/tunnel.go
//
import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/labels"

	"k8s.io/helm/pkg/kube"
)

// TODO refactor out this global var
var tunnel *kube.Tunnel

func newTillerPortForwarder(namespace string) (*kube.Tunnel, error) {
	kc := kube.New(nil)
	client, err := kc.Client()
	if err != nil {
		return nil, err
	}

	podName, err := getTillerPodName(client, namespace)
	if err != nil {
		return nil, err
	}
	const tillerPort = 44134
	return kc.ForwardPort(namespace, podName, tillerPort)
}

func getTillerPodName(client unversioned.PodsNamespacer, namespace string) (string, error) {
	// TODO use a const for labels
	selector := labels.Set{"app": "helm", "name": "tiller"}.AsSelector()
	pod, err := getFirstRunningPod(client, namespace, selector)
	if err != nil {
		return "", err
	}
	return pod.ObjectMeta.GetName(), nil
}

func getFirstRunningPod(client unversioned.PodsNamespacer, namespace string, selector labels.Selector) (*api.Pod, error) {
	options := api.ListOptions{LabelSelector: selector}
	pods, err := client.Pods(namespace).List(options)
	if err != nil {
		return nil, err
	}
	if len(pods.Items) < 1 {
		return nil, fmt.Errorf("could not find tiller")
	}
	for _, p := range pods.Items {
		if api.IsPodReady(&p) {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("could not find a ready tiller pod")
}
