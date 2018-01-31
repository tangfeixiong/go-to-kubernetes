/*
  For example:
    [vagrant@kubedev-172-17-4-59 go-to-kubernetes]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go test -v ./pkg/spec/crd/
*/

package crd

import (
	"testing"
)

func TestParser(t *testing.T) {
	recipe, err := NewRecipient("example.com", "v1", "foos", "foo")
	if err != nil {
		t.Error("creating recipe:", err)
	}

	b, err := recipe.ParseFrom([]byte{})
	if err != nil {
		t.Errorf("Executing template failed: %v", err)
	} else {
		t.Log(b.String())
	}
}

func TestDecoder(t *testing.T) {
	recipe, err := NewRecipient("example.com", "v1", "foos", "foo")
	obj, err := recipe.DecodeArtifact(nil)
	if err != nil {
		t.Error("Deserilizing CRD failed:", err)
	} else {
		t.Logf("%q", obj)
	}
}
