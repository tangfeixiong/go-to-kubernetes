package rabbitmq

import (
	"fmt"
	"testing"

	sampleapiv1alpha1 "github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/apis/example.com/v1alpha1"
)

func TestExampe(t *testing.T) {
	m := sampleapiv1alpha1.PeerDiscovery_K8s

	switch m {
	case sampleapiv1alpha1.PeerDiscovery_K8s:
		fmt.Println("Expected")
	default:
		t.Fatal("Unexpected")
	}
}
