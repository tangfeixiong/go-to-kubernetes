package crd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

const tpl = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: {{.Plural}}.{{.Group}}
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
    # singular name to be used as an alias on the CLI and for display. Defaults to lowercased <kind>
    singular: {{.Singular}}
    # shortNames allow shorter string to match your resource on the CLI. It must be all lowercase.
    #ShortNames []string
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: {{.Kind}}
    # ListKind is the serialized kind of the list for this resource.  Defaults to <kind>List.
    #ListKind string
`

type Recipient struct {
	Group, Version, Scope, Plural, Singular, Kind, ListKind string
	ShortNames                                              []string
}

func NewRecipient(group, version, plural, singular string, options ...SetFieldFunc) (*Recipient, error) {
	recipe := &Recipient{
		Group:    group,
		Version:  version,
		Scope:    "Namespaced",
		Plural:   plural,
		Singular: singular,
		Kind:     strings.Title(singular),
	}

	for _, option := range options {
		if err := option(recipe); err != nil {
			return nil, err
		}
	}

	return recipe, nil
}

var logger = log.New(os.Stderr, fmt.Sprintf("[%s] ", filepath.Base(os.Args[0])), log.LstdFlags|log.Lshortfile)

type SetFieldFunc func(*Recipient) error

func GroupVersionSetter(group, version string) SetFieldFunc {
	return func(recipe *Recipient) error {
		if match, err := regexp.MatchString(`[a-z][a-z0-9]*(\.[a-z0-9]+){0,}`, group); err == nil && match {
			recipe.Group = group
		} else if err == nil {
			return fmt.Errorf("Group should be up to maximum length of 253 characters and consist of lower case alphanumeric characters, -, and . ")
		} else {
			logger.Panic(err)
		}

		if version == "" {
			return fmt.Errorf("Version must specified")
		}
		recipe.Version = version

		return nil
	}
}

func ScopeSetter(isNamespaced bool) SetFieldFunc {
	return func(recipe *Recipient) error {
		if isNamespaced {
			recipe.Scope = "Namespaced"
		} else {
			recipe.Scope = "Cluster"
		}
		return nil
	}
}

func (recipe *Recipient) ParseFrom(crdtpl []byte) (*bytes.Buffer, error) {
	text := tpl
	if len(crdtpl) > 0 {
		text = string(crdtpl)
	}

	te := template.Must(template.New("crd").Parse(text))
	b := &bytes.Buffer{}
	if err := te.Execute(b, recipe); err != nil {
		return nil, fmt.Errorf("Executing template failed: %v\n", err)
	}

	return b, nil
}

func (recipe *Recipient) DecodeArtifact(crdtpl []byte) (*v1beta1.CustomResourceDefinition, error) {
	b, err := recipe.ParseFrom(crdtpl)
	if err != nil {
		logger.Println(err.Error())
		return nil, err
	}

	res := new(v1beta1.CustomResourceDefinition)

	err = yaml.Unmarshal(b.Bytes(), res)
	if err != nil {
		return nil, fmt.Errorf("Deserializing CRD failed: %v\n", err)
	}

	//	scheme := runtime.NewScheme()
	//	var _ runtime.ObjectTyper = scheme

	//	codecs := serializer.NewCodecFactory(scheme)
	//	codec := codecs.LegacyCodec(v1beta1.SchemeGroupVersion)
	//	obj2, err := runtime.Decode(codec, b.Bytes())

	//	info, _ := runtime.SerializerInfoForMediaType(codecs.SupportedMediaTypes(), runtime.ContentTypeJSON)
	//	jsonserializer := info.Serializer
	//	obj4, err := runtime.Decode(jsonserializer, b.Bytes())

	return res, nil
}
