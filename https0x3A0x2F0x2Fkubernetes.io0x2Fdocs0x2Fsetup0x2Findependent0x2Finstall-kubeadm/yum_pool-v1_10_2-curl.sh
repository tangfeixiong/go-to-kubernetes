#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum"

mkdir -p $download_dir/{doc,pool,repos}

pkgs=" \
    pool/32e8bd812a3944ccf07750d52088a118fa11493d34e009e2873317e0f0b0dfd2-kubectl-1.10.2-0.x86_64.rpm \
    pool/b754a6990af7d7012189610b0dc69e6e950c13a8c415b9ebea8d56352e9719fd-kubeadm-1.10.2-0.x86_64.rpm \
    pool/bdee083331998c4631bf6653454c584fb796944fe97271906acbaacbf340e1d5-kubelet-1.10.2-0.x86_64.rpm \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/yum/$i -o $download_dir/$i
done
