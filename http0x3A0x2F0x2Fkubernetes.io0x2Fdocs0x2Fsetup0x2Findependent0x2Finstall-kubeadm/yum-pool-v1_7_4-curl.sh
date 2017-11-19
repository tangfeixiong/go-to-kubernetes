#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum"

mkdir -p $download_dir/{doc,pool,repos}

pkgs=" \
    pool/f0a51fcde5e3b329050d7a6cf70f04a6cdf09eacfbad55f4324bfa2ea4312d0e-kubeadm-1.7.4-0.x86_64.rpm \
    pool/041d5a6813dab590b160865fea7259bc2db762a9667379d03aca8d4596a3cccd-kubectl-1.7.4-0.x86_64.rpm \
    pool/4f60c17a926175fb9abcfdd487cebafbbbce0e2248d2b99c189ae0877376b88d-kubelet-1.7.4-0.x86_64.rpm \
    pool/0c923b191423caccc65c550ef07ce61b97f991ad54785dab70dc07a5763f4222-kubernetes-cni-0.3.0.1-0.07a8a2.x86_64.rpm \
    pool/e7a4403227dd24036f3b0615663a371c4e07a95be5fee53505e647fd8ae58aa6-kubernetes-cni-0.5.1-0.x86_64.rpm \
    pool/79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/yum/$i -o $download_dir/$i
done