#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum"

mkdir -p $download_dir/{doc,pool,repos}

### The earlier RPMs

pkgs=" \
    pool/81158f40764a08356242a53ef4bf7e3c219f48f364c55260db571cae51ce6eec-kubeadm-1.6.2-0.x86_64.rpm \
    pool/ff72cbf0dfa986c36097a5c3ac2301ecb73ed28df8db86e13641a79e9df9b3ea-kubectl-1.6.2-0.x86_64.rpm \
    pool/a8a2b37431da90798331a67b8b452572a72173414b1543368786e356f23bc4ce-kubelet-1.6.2-0.x86_64.rpm \
"
pkgs=" \
    pool/0c37012e387fae5fd0d66b5d4664f495ab19fbb709df655533b2b2ecb1a05cd8-kubeadm-1.6.4-0.x86_64.rpm \
    pool/9fcf36d1acf5e17c4aec783e956b8b70ab0e91693d394e6268c16533fe8a71a9-kubectl-1.6.4-0.x86_64.rpm \
    pool/08706c8f6fbce961cd8994043cd27204f2bd3b47afd54cbd8755dae4b6cc4279-kubelet-1.6.4-0.x86_64.rpm \
"

pkgs=" \
    pool/f0a51fcde5e3b329050d7a6cf70f04a6cdf09eacfbad55f4324bfa2ea4312d0e-kubeadm-1.7.4-0.x86_64.rpm \
    pool/041d5a6813dab590b160865fea7259bc2db762a9667379d03aca8d4596a3cccd-kubectl-1.7.4-0.x86_64.rpm \
    pool/4f60c17a926175fb9abcfdd487cebafbbbce0e2248d2b99c189ae0877376b88d-kubelet-1.7.4-0.x86_64.rpm \
"
pkgs=" \
    pool/4ff875dc8857b85c490b42b750527ba20a154a49a8dacd256d16cbbf5e708dfd-kubeadm-1.7.6-1.x86_64.rpm \
    pool/71aa78fc7472de3664511c88f9d58d9a9c6922f26d67323869b5a661b106e0d0-kubectl-1.7.6-1.x86_64.rpm \
    pool/f049d9a0a9172b00aa2725c86a0de6f4ee51571105344b31b0b2523be9fda635-kubelet-1.7.6-1.x86_64.rpm \
"

pkgs=" \
    pool/d64bc1d0ca27196030c6f574252a0872b998b29be90c6cb7e97b91cd0bc78fed-kubeadm-1.8.2-0.x86_64.rpm \
    pool/3cc05eb1b1565779e8033743f1a2b9c8fb4e3b432421719ec56a3024d33dfccc-kubectl-1.8.2-0.x86_64.rpm \
    pool/7d976686554e598267d8c5eb030f2a8f4e575f47015ecf94459913b80c9e5bb8-kubelet-1.8.2-0.x86_64.rpm \
"
pkgs=" \
    pool/aeaad1e283c54876b759a089f152228d7cd4c049f271125c23623995b8e76f96-kubeadm-1.8.4-0.x86_64.rpm \
    pool/a9db28728641ddbf7f025b8b496804d82a396d0ccb178fffd124623fb2f999ea-kubectl-1.8.4-0.x86_64.rpm \
    pool/1acca81eb5cf99453f30466876ff03146112b7f12c625cb48f12508684e02665-kubelet-1.8.4-0.x86_64.rpm \
"

pkgs=" \
    pool/aa9948f82e7af317c97a242f3890985159c09c183b46ac8aab19d2ad307e6970-kubeadm-1.9.0-0.x86_64.rpm \
    pool/bc390a3d43256791bfb844696e7215fd7ad8a09f70a42667dab4997415a6ba75-kubectl-1.9.0-0.x86_64.rpm \
    pool/8f507de9e1cc26e5b0043e334e26d62001c171d8e54d839128e9bade25ecda95-kubelet-1.9.0-0.x86_64.rpm \
    pool/fe33057ffe95bfae65e2f269e1b05e99308853176e24a4d027bc082b471a07c0-kubernetes-cni-0.6.0-0.x86_64.rpm \
"

### The latest version!

pkgs=" \
    pool/b754a6990af7d7012189610b0dc69e6e950c13a8c415b9ebea8d56352e9719fd-kubeadm-1.10.2-0.x86_64.rpm \
    pool/32e8bd812a3944ccf07750d52088a118fa11493d34e009e2873317e0f0b0dfd2-kubectl-1.10.2-0.x86_64.rpm \
    pool/bdee083331998c4631bf6653454c584fb796944fe97271906acbaacbf340e1d5-kubelet-1.10.2-0.x86_64.rpm \
    pool/fe33057ffe95bfae65e2f269e1b05e99308853176e24a4d027bc082b471a07c0-kubernetes-cni-0.6.0-0.x86_64.rpm \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/yum/$i -o $download_dir/$i
done
