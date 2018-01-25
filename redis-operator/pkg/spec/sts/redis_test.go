/*
  go test -v -run Sentinel ./pkg/spec/sts/ -args --logtostderr --v=5
*/

package sts

import (
	"testing"
)

func TestSentinel(t *testing.T) {
	recipe := &RedisRecipient{
		Name:             "my-redis",
		ProvisioningName: "my-redis",
		ServiceName:      "my-redis",
		Replications:     2,
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
