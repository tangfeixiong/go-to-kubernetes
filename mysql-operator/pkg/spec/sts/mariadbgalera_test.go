/*
  [vagrant@kubedev-172-17-4-59 mysql-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go test -v -run Galera ./pkg/spec/sts/ -args --logtostderr --v=5
*/

package sts

import (
	"testing"
)

func TestGalera(t *testing.T) {
	recipe := NewMariadbGaleraRecipient("my-galera", "", "my-galera", "", nil, nil)

	//	b, err := recipe.parseTpl()
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	t.Log(b.String())

	s, err := recipe.Create()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%q", s)
	}
}
