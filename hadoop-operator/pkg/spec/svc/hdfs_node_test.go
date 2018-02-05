/*
  [vagrant@kubedev-172-17-4-59 hadoop-operator]$ go test -v ./pkg/spec/svc/ -args --logtostderr --v=5
*/

package svc

import (
	"testing"
)

func TestNew(t *testing.T) {
	recipe := NewHadoopHdfsNodeServiceRecipient("my-node", "examplens", "my-node")

	//	b, err := recipe.parseTpl()
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	t.Log(b.String())

	s, err := recipe.New()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%q", s)
	}
}
