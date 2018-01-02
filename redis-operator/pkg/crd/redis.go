package crd

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ghodss/yaml"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

/*
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: redises.operator.joelws.com
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: operator.joelws.com
  # version name to use for REST API: /apis/<group>/<version>
  version: v1
  # either Namespaced or Cluster
  scope: Namespaced
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: redises
    # singular name to be used as an alias on the CLI and for display
    singular: redis
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: Redis
    # shortNames allow shorter string to match your resource on the CLI
*/

const tpl = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: {{.Name}}.{{.Group}}
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: {{.Group}}
  # version name to use for REST API: /apis/<group>/<version>
  version: {{.Version}}
  # either Namespaced or Cluster
  scope: {{.Scope}}
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: {{.Plural}}
    # singular name to be used as an alias on the CLI and for display
    singular: {{.Singular}}
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: {{.Kind}}
    # shortNames allow shorter string to match your resource on the CLI
`

type Recipient struct {
	Name, Group, Version, Scope, Plural, Singular, Kind string
}

func Deserialize(recipe Recipient) (*v1beta1.CustomResourceDefinition, error) {
	te := template.Must(template.New("RedisCRD").Parse(tpl))
	b := &bytes.Buffer{}

	err := te.Execute(b, recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to import recipe: %v\n", err)
	}

	result := new(v1beta1.CustomResourceDefinition)

	err = yaml.Unmarshal(b.Bytes(), &result)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode into target: %v\n", err)
	}

	return result, nil
}
