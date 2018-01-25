#!/usr/bin/env bash

# Inspired by:
#   https://github.com/kubernetes/code-generator/blob/master/generate-groups.sh
#
#   https://github.com/jw-s/redis-operator/tree/master/hack

PROJECT=github.com/tangfeixiong/go-to-kubernetes
GOPATH="$(cd $(dirname ${BASH_SOURCE})/../../../../../.. && pwd)"

GOPATH=$GOPATH informer-gen \
  --input-dirs ${PROJECT}/redis-operator/pkg/apis/example.com/v1 \
  --versioned-clientset-package ${PROJECT}/redis-operator/pkg/client/clientset/versioned \
  --listers-package ${PROJECT}/redis-operator/pkg/client/listers \
  --output-package ${PROJECT}/redis-operator/pkg/client/informers \
  --go-header-file /dev/null \
  --logtostderr --v 2
