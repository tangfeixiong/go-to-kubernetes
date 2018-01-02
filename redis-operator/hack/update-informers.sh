#!/usr/bin/env bash

PROJECT=github.com/tangfeixiong/go-to-kubernetes
GOPATH="$(cd $(dirname ${BASH_SOURCE})/../../../../../.. && pwd)"

GOPATH=$GOPATH informer-gen \
  --go-header-file /dev/null \
  --input-dirs ${PROJECT}/redis-operator/pkg/apis/example.com/v1 \
  --output-base ${GOPATH}/src \
  --output-package ${PROJECT}/redis-operator/pkg/client/informers \
  --listers-package ${PROJECT}/redis-operator/pkg/client/listers \
  --internal-clientset-package ${PROJECT}/redis-operator/pkg/client/clientset \
  --versioned-clientset-package ${PROJECT}/redis-operator/pkg/client/clientset \
  --logtostderr \
