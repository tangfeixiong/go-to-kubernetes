/*
  [vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go test -v ./pkg/spec/sts/ -args --logtostderr --v=5
*/

package sts

import (
	"testing"
)

func TestExample(t *testing.T) {
	recipe := NewRabbitmqRecipient("my-rabbitmq", "", "my-rabbitmq", "", nil)

	//	b, err := recipe.parseTpl()
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	t.Log(b.String())

	s, err := recipe.Generate()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%q", s)
	}
}
