#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/apt"

mkdir -p $download_dir/{dists,doc,pool,src}

pkgs="pool/docker-engine_1.11.2-0~xenial_amd64_5554a8bc383e65fb10d556239c72b26457cc5d97f49e3a353c3382f8434e7d21.deb \
    pool/kubeadm_1.8.2-00_amd64_e743a9538b855d08ddaa68e7910af2dfc2bb9c1a0938d79089b0a9d3f1c19dde.deb \
    pool/kubectl_1.8.2-00_amd64_b01f6fa567e98752181a3ad057275851e00aa5fa1a8db6eb5a6d81c0f499e1ec.deb \
    pool/kubelet_1.8.2-00_amd64_9646b1262ea4c99a89e94801b150e355a4e984c462bf8fb8a2fe9a33dacd74e0.deb \
    pool/kubernetes-cni_0.5.1-00_amd64_08cbe5c42366ec3184cc91a4353f6e066f2d7325b4db1cb4f87c7dcc8c0eb620.deb \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/apt/$i -o $download_dir/$i
done
