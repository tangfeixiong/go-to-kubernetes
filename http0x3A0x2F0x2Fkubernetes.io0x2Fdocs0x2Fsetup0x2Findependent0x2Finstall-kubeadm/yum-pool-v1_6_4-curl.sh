#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum"

mkdir -p $download_dir/{doc,pool,repos}

### v1.6.2
pkgs=" \
    pool/81158f40764a08356242a53ef4bf7e3c219f48f364c55260db571cae51ce6eec-kubeadm-1.6.2-0.x86_64.rpm \
    pool/ff72cbf0dfa986c36097a5c3ac2301ecb73ed28df8db86e13641a79e9df9b3ea-kubectl-1.6.2-0.x86_64.rpm \
    pool/a8a2b37431da90798331a67b8b452572a72173414b1543368786e356f23bc4ce-kubelet-1.6.2-0.x86_64.rpm \
    pool/0c923b191423caccc65c550ef07ce61b97f991ad54785dab70dc07a5763f4222-kubernetes-cni-0.3.0.1-0.07a8a2.x86_64.rpm \
    pool/e7a4403227dd24036f3b0615663a371c4e07a95be5fee53505e647fd8ae58aa6-kubernetes-cni-0.5.1-0.x86_64.rpm \
    pool/79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm \
"

### v1.6.4
pkgs=" \
    pool/0c37012e387fae5fd0d66b5d4664f495ab19fbb709df655533b2b2ecb1a05cd8-kubeadm-1.6.4-0.x86_64.rpm \
    pool/9fcf36d1acf5e17c4aec783e956b8b70ab0e91693d394e6268c16533fe8a71a9-kubectl-1.6.4-0.x86_64.rpm \
    pool/08706c8f6fbce961cd8994043cd27204f2bd3b47afd54cbd8755dae4b6cc4279-kubelet-1.6.4-0.x86_64.rpm \
    pool/0c923b191423caccc65c550ef07ce61b97f991ad54785dab70dc07a5763f4222-kubernetes-cni-0.3.0.1-0.07a8a2.x86_64.rpm \
    pool/e7a4403227dd24036f3b0615663a371c4e07a95be5fee53505e647fd8ae58aa6-kubernetes-cni-0.5.1-0.x86_64.rpm \
    pool/79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm \
"
pkgs=" \
    pool/0c37012e387fae5fd0d66b5d4664f495ab19fbb709df655533b2b2ecb1a05cd8-kubeadm-1.6.4-0.x86_64.rpm \
    pool/9fcf36d1acf5e17c4aec783e956b8b70ab0e91693d394e6268c16533fe8a71a9-kubectl-1.6.4-0.x86_64.rpm \
    pool/08706c8f6fbce961cd8994043cd27204f2bd3b47afd54cbd8755dae4b6cc4279-kubelet-1.6.4-0.x86_64.rpm \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/yum/$i -o $download_dir/$i
done
