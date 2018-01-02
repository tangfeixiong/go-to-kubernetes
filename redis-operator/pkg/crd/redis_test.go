package crd

import (
	"bytes"
	"testing"
	"text/template"
)

var (
	recipe Recipient = Recipient{
		Name:     "redises",
		Group:    "operator.joelws.com",
		Version:  "v1",
		Scope:    "Namespaced",
		Plural:   "redises",
		Singular: "redis",
		Kind:     "Redis",
	}
)

func TestCRD_gen(t *testing.T) {
	te := template.Must(template.New("RedisCRD").Parse(tpl))
	b := &bytes.Buffer{}

	err := te.Execute(b, recipe)
	if err != nil {
		t.Error("executing template:", err)
	} else {
		t.Log(b.String())
	}
}

func TestCRD_decode(t *testing.T) {
	result, err := Deserialize(recipe)
	if err != nil {
		t.Error("deserilize CRD:", err)
	} else {
		t.Logf("%+v", result)
	}
}
