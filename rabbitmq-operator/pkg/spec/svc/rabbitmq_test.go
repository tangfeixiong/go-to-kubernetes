/*
  [vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ go test -v ./pkg/spec/svc/ -args --logtostderr --v=5
*/

package svc

import (
	"testing"
)

func TestExample(t *testing.T) {
	recipe := NewRabbitmqRecipient("my-rabbitmq", "default", "my-rabbitmq")

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
