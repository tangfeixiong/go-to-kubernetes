#!/usr/bin/env bash

PROJECT=github.com/tangfeixiong/go-to-kubernetes
GOPATH="$(cd $(dirname ${BASH_SOURCE})/../../../../../.. && pwd)"
if [[ -L ${HOME}/go ]]; then
  GOPATH+=":$(readlink $HOME/go)"
else
  GOPATH+=":$HOME/go"
fi

GOPATH=$GOPATH client-gen \
  --clientset-name versioned \
  --input-base "" \
  --input ${PROJECT}/redis-operator/pkg/apis/example.com/v1 \
  --output-package ${PROJECT}/redis-operator/pkg/client/clientset \
  --go-header-file /dev/null \
  --fake-clientset=false \
  --logtostderr --v 2