#!/usr/bin/env bash

PROJECT=github.com/tangfeixiong/go-to-kubernetes
GOPATH="$(cd $(dirname ${BASH_SOURCE})/../../../../../.. && pwd)"

GOPATH=$GOPATH deepcopy-gen \
  --input-dirs ${PROJECT}/redis-operator/pkg/apis/example.com/v1 \
  --output-file-base zz_generated.deepcopy \
  --bounding-dirs ${PROJECT}/redis-operator/pkg/apis \
  --go-header-file /dev/null \
  --logtostderr --v 2