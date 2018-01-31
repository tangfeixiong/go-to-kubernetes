/*
  go test -v -run Galera ./pkg/spec/svc/ -args --logtostderr --v=5
*/

package svc

import (
	"testing"
)

func TestGalera(t *testing.T) {
	recipe := NewMariadbGaleraRecipient("my-galera", "default", "my-galera")

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
