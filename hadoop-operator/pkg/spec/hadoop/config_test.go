/*
  For example:
    [vagrant@kubedev-172-17-4-59 hadoop-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go test -v -run Config ./pkg/spec/hadoop/
*/

package hadoop

import (
	"os"
	"path/filepath"
	"testing"

	//sampleapiv1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
)

func TestConfig_Parse(t *testing.T) {
	recipe, err := NewHadoopRecipient(NameNodeHostSetter("my-hadoop-0.my-hdfs"),
		ResourceNameSetter("my-hadoop"), CustomResourceNameSetter("demo-hadoop"))
	if err != nil {
		t.Fatal(err)
	}

	b, err := recipe.Parse(core_asset_name)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(b.String())
	}

	b, err = recipe.Parse(hdfs_asset_name)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(b.String())
	}

	b, err = recipe.Parse(yarn_asset_name)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(b.String())
	}

	b, err = recipe.Parse(mapred_asset_name)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(b.String())
	}
}

func TestConfig_Generate(t *testing.T) {
	recipe, err := NewHadoopRecipient(NameNodeHostSetter("127.0.0.1"),
		ResourceNameSetter("my-hadoop"), CustomResourceNameSetter("demo-hadoop"))
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(os.Getenv("HOME"))
	err = recipe.Generate(path)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConfig_New(t *testing.T) {
	recipe, err := NewHadoopRecipient(NameNodeFqdnSetter("my-hadoop-0", "my-hadoop", "example-app", "svc", "cluster.local"),
		ResourceNameSetter("my-hadoop"), CustomResourceNameSetter("demo-hadoop"))
	if err != nil {
		t.Fatal(err)
	}

	obj, err := recipe.New()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%v", obj)
	}
}
