/*
  For example:
    [vagrant@kubedev-172-17-4-59 hadoop-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go test -v -run Script ./pkg/spec/hadoop/
*/

package hadoop

import (
	"os"
	"path/filepath"
	"testing"

	//sampleapiv1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
)

func TestScript_Parse(t *testing.T) {
	recipe, err := NewHadoopStartupRecipient(HadoopHomeSetter("/Users/fanhongling/Downloads/99-mirror/linux-bin/hadoop-3.0.0/"),
		BinNameSetter("hdfs"), CmdNameSetter("datanode"))
	if err != nil {
		t.Fatal(err)
	}

	b, err := recipe.Parse(script_asset_name)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(b.String())
	}
}

func TestScript_Generate(t *testing.T) {
	recipe, err := NewHadoopStartupRecipient(HadoopHomeSetter("/Users/fanhongling/Downloads/99-mirror/linux-bin/hadoop-3.0.0/"),
		BinNameSetter("hdfs"), CmdNameSetter("namenode"))
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(os.Getenv("HOME"))
	err = recipe.Generate(path)
	if err != nil {
		t.Fatal(err)
	}
}
