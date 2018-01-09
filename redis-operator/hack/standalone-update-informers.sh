#!/usr/bin/env bash

PROJECT=github.com/tangfeixiong/go-to-kubernetes
GOPATH="$(cd $(dirname ${BASH_SOURCE})/../../../../../.. && pwd)"

GOPATH=$GOPATH informer-gen \
  --input-dirs ${PROJECT}/redis-operator/pkg/apis/example.com/v1 \
  --versioned-clientset-package ${PROJECT}/redis-operator/pkg/client/clientset/versioned \
  --listers-package ${PROJECT}/redis-operator/pkg/client/listers \
  --output-package ${PROJECT}/redis-operator/pkg/client/informers \
  --go-header-file /dev/null \
  --logtostderr --v 2
