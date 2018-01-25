#!/bin/bash

# Inspired by:
#
#   https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/
#   https://github.com/openshift-evangelists/crd-code-generation/blob/master/hack/update-codegen.sh

#   https://github.com/kubernetes/sample-controller/blob/master/hack/update-codegen.sh
#   https://github.com/kubernetes/code-generator/blob/master/hack/update-codegen.sh
#   https://github.com/kubernetes/code-generator/blob/master/generate-groups.sh

#   https://github.com/rook/rook/blob/master/build/codegen/codegen.sh

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
#GOPATH=${GOPATH:-$(go env GOPATH)}
GOPATH=$(cd ${SCRIPT_ROOT}/../../../../.. && pwd)
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ../vendor/k8s.io/code-generator 2>/dev/null || echo ${GOPATH}/src/k8s.io/code-generator)}

if [[ ! -e $GOPATH/src/k8s.io ]]; then
    ln -s $(go env GOPATH)/src/k8s.io $GOPATH/src/k8s.io
fi

cd $SCRIPT_ROOT/../ && GOPATH=$GOPATH CODEGEN_PKG=vendor/k8s.io/code-generator \
vendor/k8s.io/code-generator/generate-groups.sh all \
  github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis \
  example.com:v1 
