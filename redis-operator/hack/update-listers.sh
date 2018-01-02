#!/usr/bin/env bash

PROJECT=github.com/tangfeixiong/go-to-kubernetes
GOPATH="$(cd $(dirname ${BASH_SOURCE})/../../../../../.. && pwd)"

GOPATH=$GOPATH lister-gen \
  --go-header-file /dev/null \
  --input-dirs ${PROJECT}/redis-operator/pkg/apis/example.com/v1 \
  --output-base ${OUTPUT_BASE} \
  --output-package ${PROJECT}/redis-operator/pkg/client/listers \
  --logtostderr \
  --v 2
