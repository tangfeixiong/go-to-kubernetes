/*
  go test -v -run Sentinel ./pkg/spec/deploy/ -args --logtostderr --v=5
*/

package deploy

import (
	"testing"
)

func TestSentinel(t *testing.T) {
	recipe := &SentinelRecipient{
		Name:             "my-redis",
		ProvisioningName: "my-redis-ha",
		Replications:     3,
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
