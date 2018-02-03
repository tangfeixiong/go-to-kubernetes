/*
  For example:
    [vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go test -v ./pkg/spec/rabbitmq/
*/

package rabbitmq

import (
	"os"
	"path/filepath"
	"testing"

	sampleapiv1alpha1 "github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/apis/example.com/v1alpha1"
	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/spec/artifact"
)

func TestParse(t *testing.T) {
	recipe, err := NewRabbitmqRecipient(DomainNameSetter("cluster.local"))
	if err != nil {
		t.Fatal(err)
	}

	asset, err := artifact.Asset("template/rabbitmq.conf.tpl")
	if err != nil {
		t.Fatal(err)
	}

	b, err := recipe.parse(asset)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(b.String())
	}

	b, err = recipe.Parse("template/enabled_plugins.tpl")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(b.String())
	}
}

func TestFile(t *testing.T) {
	recipe, err := NewRabbitmqRecipient(ClusteringModeSetter(string(sampleapiv1alpha1.PeerDiscovery_K8s)), NamespaceSetter("example-system"), TargetServiceSetter("my-rabbitmq-cluster"))
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(os.Getenv("HOME"))
	err = recipe.Generate(path)
	if err != nil {
		t.Fatal(err)
	}
}
