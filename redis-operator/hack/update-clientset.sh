#!/usr/bin/env bash

PROJECT=github.com/tangfeixiong/go-to-kubernetes
GOPATH="$(cd $(dirname ${BASH_SOURCE})/../../../../../.. && pwd)"

GOPATH=$GOPATH client-gen \
  --input-base ${PROJECT}/redis-operator/pkg/apis \
  --clientset-path ${PROJECT}/redis-operator/pkg/client \
  --input example.com/v1 \
  --clientset-name clientset \
  --fake-clientset=false \
  --logtostderr \
  --v 2