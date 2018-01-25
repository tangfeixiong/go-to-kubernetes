/*
  go test -v -run Sentinel ./pkg/spec/svc/ -args --logtostderr --v=5
*/

package svc

import (
	"testing"
)

func TestSentinel(t *testing.T) {
	recipe := &SentinelRecipient{
		Name:         "my-redis-ha",
		ResourceName: "my-redis",
	}

	//	b, err := recipe.parseTpl()
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	t.Log(b.String())

	s, err := recipe.GenerateArtifact()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%q", s)
	}
}
