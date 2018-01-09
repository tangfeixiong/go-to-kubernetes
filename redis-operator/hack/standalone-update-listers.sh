#!/usr/bin/env bash

PROJECT=github.com/tangfeixiong/go-to-kubernetes
GOPATH="$(cd $(dirname ${BASH_SOURCE})/../../../../../.. && pwd)"

GOPATH=$GOPATH lister-gen \
  --input-dirs ${PROJECT}/redis-operator/pkg/apis/example.com/v1 \
  --output-package ${PROJECT}/redis-operator/pkg/client/listers \
  --go-header-file /dev/null \
  --logtostderr --v 2
