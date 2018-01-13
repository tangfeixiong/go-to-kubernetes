package crd

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/artifact"
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
	var content string

	//content = tpl
	data, err := artifact.Asset("template/custom-resource-definition.yaml")
	if err != nil {
		glog.Fatalf("Cloud not find spec data, error: %s", err.Error())
		return nil, err
	} else {
		content = string(data)
		glog.V(5).Infoln("CRD artifact:", content)
	}

	te := template.Must(template.New("RedisCRD").Parse(content))
	b := &bytes.Buffer{}

	err = te.Execute(b, recipe)
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
