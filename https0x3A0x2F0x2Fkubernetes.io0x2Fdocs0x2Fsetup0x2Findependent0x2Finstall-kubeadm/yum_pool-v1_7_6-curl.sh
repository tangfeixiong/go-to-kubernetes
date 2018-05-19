#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum"

mkdir -p $download_dir/{doc,pool,repos}

pkgs=" \
    pool/4ff875dc8857b85c490b42b750527ba20a154a49a8dacd256d16cbbf5e708dfd-kubeadm-1.7.6-1.x86_64.rpm \
    pool/71aa78fc7472de3664511c88f9d58d9a9c6922f26d67323869b5a661b106e0d0-kubectl-1.7.6-1.x86_64.rpm \
    pool/f049d9a0a9172b00aa2725c86a0de6f4ee51571105344b31b0b2523be9fda635-kubelet-1.7.6-1.x86_64.rpm \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/yum/$i -o $download_dir/$i
done
