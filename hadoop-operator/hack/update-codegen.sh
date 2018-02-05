#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# Inspired by:
#   https://github.com/kubernetes/code-generator/blob/master/generate-groups.sh
#
#   https://github.com/jw-s/redis-operator/tree/master/hack

PROJECT='github.com/tangfeixiong/go-to-kubernetes'

OPERATOR='hadoop-operator'
GROUP='example.com'
VERSION='v1alpha1'

GOPATH="$(cd $(dirname ${BASH_SOURCE})/../../../../../.. && pwd)"
#if [[ -L ${HOME}/go ]]; then
#  GOPATH+=":$(readlink $HOME/go)"
#else
#  GOPATH+=":$HOME/go"
#fi
#echo $GOPATH

echo "--------deepcopy--------"
GOPATH=$GOPATH deepcopy-gen \
  --input-dirs ${PROJECT}/${OPERATOR}/pkg/apis/${GROUP}/${VERSION} \
  --output-file-base zz_generated.deepcopy \
  --bounding-dirs ${PROJECT}/${OPERATOR}/pkg/apis \
  --go-header-file /dev/null \
  --logtostderr --v 2

echo "--------clientset--------"
GOPATH=$GOPATH client-gen \
  --clientset-name versioned \
  --input-base "" \
  --input ${PROJECT}/${OPERATOR}/pkg/apis/${GROUP}/${VERSION} \
  --output-package ${PROJECT}/${OPERATOR}/pkg/client/clientset \
  --go-header-file /dev/null \
  --fake-clientset=false \
  --logtostderr --v 2

echo "--------listers--------"
GOPATH=$GOPATH lister-gen \
  --input-dirs ${PROJECT}/${OPERATOR}/pkg/apis/${GROUP}/${VERSION} \
  --output-package ${PROJECT}/${OPERATOR}/pkg/client/listers \
  --go-header-file /dev/null \
  --logtostderr --v 2

echo "--------informers--------"
GOPATH=$GOPATH informer-gen \
  --input-dirs ${PROJECT}/${OPERATOR}/pkg/apis/${GROUP}/${VERSION} \
  --versioned-clientset-package ${PROJECT}/${OPERATOR}/pkg/client/clientset/versioned \
  --listers-package ${PROJECT}/${OPERATOR}/pkg/client/listers \
  --output-package ${PROJECT}/${OPERATOR}/pkg/client/informers \
  --go-header-file /dev/null \
  --logtostderr --v 2
