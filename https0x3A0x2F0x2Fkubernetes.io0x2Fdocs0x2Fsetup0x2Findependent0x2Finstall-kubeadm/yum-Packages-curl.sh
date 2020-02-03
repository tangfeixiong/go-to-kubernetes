#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64"

if [[ ! -d $download_dir/Packages ]]; then mkdir -p $download_dir/Packages; fi

#; The earlier RPMs

pkgs=" \
    Packages/0c37012e387fae5fd0d66b5d4664f495ab19fbb709df655533b2b2ecb1a05cd8-kubeadm-1.6.4-0.x86_64.rpm \
    Packages/9fcf36d1acf5e17c4aec783e956b8b70ab0e91693d394e6268c16533fe8a71a9-kubectl-1.6.4-0.x86_64.rpm \
    Packages/08706c8f6fbce961cd8994043cd27204f2bd3b47afd54cbd8755dae4b6cc4279-kubelet-1.6.4-0.x86_64.rpm \
"

pkgs=" \
    Packages/4ff875dc8857b85c490b42b750527ba20a154a49a8dacd256d16cbbf5e708dfd-kubeadm-1.7.6-1.x86_64.rpm \
    Packages/71aa78fc7472de3664511c88f9d58d9a9c6922f26d67323869b5a661b106e0d0-kubectl-1.7.6-1.x86_64.rpm \
    Packages/f049d9a0a9172b00aa2725c86a0de6f4ee51571105344b31b0b2523be9fda635-kubelet-1.7.6-1.x86_64.rpm \
"

pkgs=" \
    Packages/0c923b191423caccc65c550ef07ce61b97f991ad54785dab70dc07a5763f4222-kubernetes-cni-0.3.0.1-0.07a8a2.x86_64.rpm \
    Packages/e7a4403227dd24036f3b0615663a371c4e07a95be5fee53505e647fd8ae58aa6-kubernetes-cni-0.5.1-0.x86_64.rpm \
    Packages/aeaad1e283c54876b759a089f152228d7cd4c049f271125c23623995b8e76f96-kubeadm-1.8.4-0.x86_64.rpm \
    Packages/a9db28728641ddbf7f025b8b496804d82a396d0ccb178fffd124623fb2f999ea-kubectl-1.8.4-0.x86_64.rpm \
    Packages/1acca81eb5cf99453f30466876ff03146112b7f12c625cb48f12508684e02665-kubelet-1.8.4-0.x86_64.rpm \
"

pkgs=" \
    Packages/aa9948f82e7af317c97a242f3890985159c09c183b46ac8aab19d2ad307e6970-kubeadm-1.9.0-0.x86_64.rpm \
    Packages/bc390a3d43256791bfb844696e7215fd7ad8a09f70a42667dab4997415a6ba75-kubectl-1.9.0-0.x86_64.rpm \
    Packages/8f507de9e1cc26e5b0043e334e26d62001c171d8e54d839128e9bade25ecda95-kubelet-1.9.0-0.x86_64.rpm \
    Packages/79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm \
"

pkgs=" \
    Packages/b754a6990af7d7012189610b0dc69e6e950c13a8c415b9ebea8d56352e9719fd-kubeadm-1.10.2-0.x86_64.rpm \
    Packages/32e8bd812a3944ccf07750d52088a118fa11493d34e009e2873317e0f0b0dfd2-kubectl-1.10.2-0.x86_64.rpm \
    Packages/bdee083331998c4631bf6653454c584fb796944fe97271906acbaacbf340e1d5-kubelet-1.10.2-0.x86_64.rpm \
    Packages/fe33057ffe95bfae65e2f269e1b05e99308853176e24a4d027bc082b471a07c0-kubernetes-cni-0.6.0-0.x86_64.rpm \
"

#; 2019-12-17

pkgs=" \
    Packages/f70d8cb23c7b91c0509292f4862060367edabce8788b855c38a7c174f014b9e2-cri-tools-1.11.1-0.x86_64.rpm \
    Packages/53edc739a0e51a4c17794de26b13ee5df939bd3161b37f503fe2af8980b41a89-cri-tools-1.12.0-0.x86_64.rpm \
    Packages/14bfe6e75a9efc8eca3f638eb22c7e2ce759c67f95b43b16fae4ebabde1549f3-cri-tools-1.13.0-0.x86_64.rpm \
    Packages/5bc6edbce5a26d13aad44dbecde701caf1cc40b2747a5255dae1afdbf879443e-kubeadm-1.10.13-0.x86_64.rpm \
    Packages/fa3dfb16bbf2e91381debe50eeef2040d4af723fa799c12107b413e1f0ca2491-kubeadm-1.11.10-0.x86_64.rpm \
    Packages/4943c168a879e38d88e2d9eb4be4ac0d462ec88ead8f6eb2e25fa44cb137415b-kubeadm-1.12.10-0.x86_64.rpm \
    Packages/e8a8c1b525d40f9f2bbf1ec1c53e8545665e23692a00f805198439dac3fd3f7b-kubeadm-1.13.12-0.x86_64.rpm \
    Packages/e2de8688286f5f4516813d1346654a50d64c4edf68a7bbeae98946cd99eb41a7-kubeadm-1.14.10-0.x86_64.rpm \
    Packages/82c26855d0f055599563d1b46866a7b3f1b009b18f60f91dfbbeb5e078c1852c-kubeadm-1.15.7-0.x86_64.rpm \
    Packages/937e48b992617a0eeae40c93ad2105005ebf8491b4cbc1a43d2bcb35e8a9e865-kubeadm-1.16.4-0.x86_64.rpm \
    Packages/4815089dde3fc771734b222016132caa8c8a2848048aa2a2169d319436e8d032-kubectl-1.10.13-0.x86_64.rpm \
    Packages/5ce34f8aadc7cf77445262e9e8ddc530753c01c1ad79a87022da10bbc7607193-kubectl-1.11.10-0.x86_64.rpm \
    Packages/1ac30b454049ba3e7896a648e20dab138de45ea7a162d7d1d918e676cfe29d63-kubectl-1.12.10-0.x86_64.rpm \
    Packages/fbc121a458bbd7db9a547282e5b9fb8a15117d86d1df21e7e034c17f3fda839b-kubectl-1.13.12-0.x86_64.rpm \
    Packages/0c130e4f7834bdc9e6cf088f30331520d424f667e6598c3e7a8c1b403c9a6e74-kubectl-1.14.10-0.x86_64.rpm \
    Packages/8922dc33ce7258b1dc00786108640d17739523794ced9cee27f12898bfeb6f80-kubectl-1.15.7-0.x86_64.rpm \
    Packages/70b9657194f6d201244df9e111e9270d162241c014a196125e0c399ad45613e7-kubectl-1.16.4-0.x86_64.rpm \
    Packages/7b77927bb6aeecca5a16db120ec131820d8dba6f468cfde31d98263032d93c5c-kubelet-1.10.13-0.x86_64.rpm \
    Packages/8c80b05ea419461398ae0fc9faaad0a4215a6313c4a7804f08fb25393c9aed49-kubelet-1.11.10-0.x86_64.rpm \
    Packages/8cfb58f4ca56a955a65d28a99f7e0e4178fc0ac889936ce1cb6e53eda58437ed-kubelet-1.12.10-0.x86_64.rpm \
    Packages/fe8b2abe201059a84229ba14798d41a82fd547dfe02eea65e19444a571ecce36-kubelet-1.13.12-0.x86_64.rpm \
    Packages/7081b7854fee4f93b44d3fc523fc946ceb4f69b29bedcdc1a4194638acca7598-kubelet-1.14.10-0.x86_64.rpm \
    Packages/b3220d8f199d4e5fc1b12433db3ea3766ce6e5b827c352c95efcd3824125fa20-kubelet-1.15.7-0.x86_64.rpm \
    Packages/d13f20c41037ec4923cc02f0aa48a67061284f3f77c3a05c2d19df72b89bbc98-kubelet-1.16.4-0.x86_64.rpm \
    Packages/548a0dcd865c16a50980420ddfa5fbccb8b59621179798e6dc905c9bf8af3b34-kubernetes-cni-0.7.5-0.x86_64.rpm \
"

#; Feb 3 2020
pkgs=" \
Packages/105d89f0607c7baf91305ba352e78000bd20aad5cdf706bffff3b31cd546dbf3-kubeadm-1.17.2-0.x86_64.rpm
Packages/b44630896c69cd411db53be1d5cb5ae899a40aba7c0766317ea904390fcfc45b-kubectl-1.17.2-0.x86_64.rpm
Packages/3ee7f2dff78e6fbb3ac3af8acb1a907f4bec1b1ef4cf627cbe02fa553707f2e9-kubelet-1.17.2-0.x86_64.rpm
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/$i -o $download_dir/$i
done
