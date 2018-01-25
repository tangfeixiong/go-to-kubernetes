/*
  go test -v -run Redis ./pkg/spec/svc/ -args --logtostderr --v=5
*/

package svc

import (
	"testing"
)

func TestRedis(t *testing.T) {
	recipe := &RedisRecipient{
		Name:         "my-redis",
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
