/*
  For example:
    [vagrant@kubedev-172-17-4-59 mysql-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go test -v ./pkg/spec/galera/
*/

package galera

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/artifact"
)

func TestParse(t *testing.T) {
	recipe, err := NewRecipient(true, ThisNodeSetter("127.0.0.1", ""))
	if err != nil {
		t.Fatal(err)
	}

	asset, err := artifact.Asset("template/galera.cnf.tpl")
	if err != nil {
		t.Fatal(err)
	}

	b, err := recipe.parse(asset)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(b.String())
	}
}

func TestCreate(t *testing.T) {
	recipe, err := NewRecipient(false,
		ClusterAddressesSetter("192.168.1.10", "192.168.1.11", "192.168.1.12"),
		WsrepProviderOptionsSetter("pc.wait_prim=FALSE", "pc.bootstrap=TRUE"))
	if err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(os.Getenv("HOME"), "galera.cnf")
	err = recipe.Generate(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("galera.cnf wrote into $HOME")
}
