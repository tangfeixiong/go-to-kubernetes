package component

import (
	"testing"
)

func TestSAD_example(t *testing.T) {
	sad := NewServiceAutoDiscovery()
	result, err := sad.EnvVar("default")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Service auto discovered (%d): %+v", len(result), result)
}
