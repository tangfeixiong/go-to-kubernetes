#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum"

mkdir -p $download_dir/{doc,pool,repos}

pkgs=" \
    pool/d64bc1d0ca27196030c6f574252a0872b998b29be90c6cb7e97b91cd0bc78fed-kubeadm-1.8.2-0.x86_64.rpm \
    pool/3cc05eb1b1565779e8033743f1a2b9c8fb4e3b432421719ec56a3024d33dfccc-kubectl-1.8.2-0.x86_64.rpm \
    pool/7d976686554e598267d8c5eb030f2a8f4e575f47015ecf94459913b80c9e5bb8-kubelet-1.8.2-0.x86_64.rpm \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/yum/$i -o $download_dir/$i
done
