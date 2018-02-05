/*
  [vagrant@kubedev-172-17-4-59 hadoop-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go test -v ./pkg/spec/sts/ -args --logtostderr --v=5
*/

package sts

import (
	"testing"
)

func TestNew(t *testing.T) {
	recipe := NewHadoopHdfsNodeRecipient("my-hadoop", "", "my-hadoop-cr", "", "my-hdfs", "namenode", nil)

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
