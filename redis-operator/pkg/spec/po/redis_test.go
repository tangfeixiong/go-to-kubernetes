package po

import (
	"testing"
)

var (
	recipe Recipient = Recipient{
		ClusterName: "my-redis",
		Name:        "redis",
		Namespace:   "example-system",
	}
)

func TestPOD_decode(t *testing.T) {
	result, err := Deserialize(recipe)
	if err != nil {
		t.Error("deserilize POD from go template:", err)
	} else {
		t.Logf("%+v", result)
	}
}
