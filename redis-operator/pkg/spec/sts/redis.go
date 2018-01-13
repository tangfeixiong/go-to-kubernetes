package sts

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"

	apiv1 "k8s.io/api/core/v1"
	// "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/artifact"
)

type Recipient struct {
	ClusterName, Name, Namespace, Image string
	SentinelQuorum                      uint8
}

const (
	DEFAULT_NAMESPACE       = "default"
	DEFAULT_IMAGE           = "docker.io/redis:4.0-alpine"
	DEFAULT_SENTINEL_QUORUM = 2
)

func Deserialize(recipe Recipient) (*apiv1.Pod, error) {
	if recipe.Namespace == "" {
		recipe.Namespace = DEFAULT_NAMESPACE
	}
	if recipe.Image == "" {
		recipe.Image = DEFAULT_IMAGE
	}
	if recipe.SentinelQuorum == 0 {
		recipe.SentinelQuorum = DEFAULT_SENTINEL_QUORUM
	}

	var content string

	data, err := artifact.Asset("template/redis-bootstrap.yaml")
	if err != nil {
		glog.Fatalf("Cloud not find spec data, error: %s", err.Error())
		return nil, err
	} else {
		content = string(data)
		glog.V(5).Infoln("POD artifact:", content)
	}

	te := template.Must(template.New("pod").Parse(content))
	b := &bytes.Buffer{}

	err = te.Execute(b, recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to import recipe: %v\n", err)
	}

	result := new(apiv1.Pod)

	err = yaml.Unmarshal(b.Bytes(), &result)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode into target: %v\n", err)
	}

	return result, nil
}
