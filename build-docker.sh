#!/bin/bash
set -e

if [ $# -ne 1 ]; then
	echo "Usage: $0 tag" >/dev/stderr
	exit 1
fi

if [[ ! -f $GOPATH/bin/kubectl ]]; then
    echo "kubectl required" >/dev/stderr
	exit 1
fi

if [[ ! -f $HOME/.gotty ]]; then
    echo "gotty config required" >/dev/stderr
	exit 1
fi

if [[ ! -d $HOME/ssl || ! -f $HOME/ssl/ca.pem ]]; then
    echo "kube cert required" >/dev/stderr
	exit 1
fi

HUB_REPO="qingyuanos.com/admin/caas-hterm"

tag=$1

ROOT_FS=$(dirname "${BASH_SOURCE}")/rootfs

go install -v github.com/tangfeixiong/go-to-kubernetes/cmd/exec2hterm

cp -f $GOPATH/bin/exec2hterm $ROOT_FS

cp -f $HOME/.gotty $ROOT_FS

cp -f $GOPATH/bin/kubectl $ROOT_FS

cp -rf $HOME/ssl/* $ROOT_FS

docker build -t ${HUB_REPO}:${tag} $ROOT_FS

rm $ROOT_FS/exec2hterm $ROOT_FS/.gotty $ROOT_FS/kubectl