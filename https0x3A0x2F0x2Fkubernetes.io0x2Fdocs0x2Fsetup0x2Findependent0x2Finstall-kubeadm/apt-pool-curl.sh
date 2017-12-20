#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/apt"

mkdir -p $download_dir/{dists,doc,pool,src}

pkgs=" \
    pool/kubeadm_1.8.4-00_amd64_0088836fbb451bc49ece82f34c035f50f2e1dd4dea78f6d585574d585da11e8e.deb \
    pool/kubectl_1.8.4-00_amd64_b48511a481ddcfbf935ad935bc6c3992c1c4315fcd8f3f794e367b9b26b775be.deb \
    pool/kubelet_1.8.4-00_amd64_601882506070723b643552ae98325c849840b09b1fc1666de74c7b69a07f06df.deb \
"

### For me, previous pkgs have already download, edit by yourselves!

pkgs=" \
    pool/kubeadm_1.9.0-00_amd64_191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911.deb \
    pool/kubectl_1.9.0-00_amd64_8ea712c18d89d56090c8753a9630d22fd8ae5cb03d4ec79a1fe6d262c8b4eb36.deb \
    pool/kubelet_1.9.0-00_amd64_100df9788226fe76ce828503cf24b8c4c6f9705f15504093844c9138e0b2a97f.deb \
    pool/kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/apt/$i -o $download_dir/$i
done
