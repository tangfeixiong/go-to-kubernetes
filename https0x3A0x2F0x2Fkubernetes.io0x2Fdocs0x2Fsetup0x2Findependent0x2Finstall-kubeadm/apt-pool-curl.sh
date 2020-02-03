#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/apt"

mkdir -p $download_dir/{dists,doc,pool,src}

#; The earlier DEBs

pkgs=" \
    pool/kubeadm_1.8.4-00_amd64_0088836fbb451bc49ece82f34c035f50f2e1dd4dea78f6d585574d585da11e8e.deb \
    pool/kubectl_1.8.4-00_amd64_b48511a481ddcfbf935ad935bc6c3992c1c4315fcd8f3f794e367b9b26b775be.deb \
    pool/kubelet_1.8.4-00_amd64_601882506070723b643552ae98325c849840b09b1fc1666de74c7b69a07f06df.deb \
"

pkgs=" \
    pool/kubeadm_1.9.0-00_amd64_191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911.deb \
    pool/kubectl_1.9.0-00_amd64_8ea712c18d89d56090c8753a9630d22fd8ae5cb03d4ec79a1fe6d262c8b4eb36.deb \
    pool/kubelet_1.9.0-00_amd64_100df9788226fe76ce828503cf24b8c4c6f9705f15504093844c9138e0b2a97f.deb \
    pool/kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb \
"

#; May 19 2018

pkgs=" \
    pool/kubeadm_1.10.2-00_amd64_4bdf79bcda2a820210f5c0f7e84f9f57d6cc196ba9b59943f296544624e6d743.deb \
    pool/kubectl_1.10.2-00_amd64_a3ca21b1875a20567024584f18ec6aca390e8be8a8d92e20b5b7258452e419c9.deb \
    pool/kubelet_1.10.2-00_amd64_3ff69468521886de7b64ba9b4932bab7e6ef71f7344ace36ac0855df7bd2427a.deb \
    pool/kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb \
"

#; Dec 8 2019

pkgs=" \
    pool/cri-tools_1.11.1-00_amd64_e6300f0f4ec2fb5d6967533416580e1a22be90277944370ceb2551b04d7bc1a3.deb \
    pool/cri-tools_1.12.0-00_amd64_2d9f048a50a9dfeceebd84635f1322955aca6381d9c05b4d60b3da1edb7d856c.deb \
    pool/cri-tools_1.13.0-00_amd64_6930e446a683884314deef354fbd8a7c5fc2be5c69c58903ad83b69b42529da4.deb \
    pool/kubeadm_1.10.13-00_amd64_4594e00a6f68d2c7f8c7990cf6408ca6f1582b9df05de22d3d68fd84cf05344c.deb \
    pool/kubeadm_1.11.10-00_amd64_a518251af86868412223ce6b12d0212d5783278419bf7cf403b77cce1a31cfb9.deb \
    pool/kubeadm_1.12.10-00_amd64_600513b5bdc3ff15caab99660a4e0e8d5b7d31bc907ce63f8a4fcd18e0c9c6f6.deb \
    pool/kubeadm_1.13.12-00_amd64_5791923bba94b850ca898fdfaa75916bdd3b4f5d86a50ea738d96f14230029b9.deb \
    pool/kubeadm_1.14.9-00_amd64_5c8c4bba3ba94e1349acc2148b57844ffad756bbb20e2ed76e655e0a074ee244.deb \
    pool/kubeadm_1.15.6-00_amd64_2c1d026f59d6e012a8fc7658dddcde404f73c9e1159cdeb000aea0d4de3ed3b3.deb \
    pool/kubeadm_1.16.3-00_amd64_8ad50e8252c65f968e35b1a38240b0ee66ef3b72c5d498830788e123238aa1a6.deb \
    pool/kubectl_1.10.13-00_amd64_46d296e9a00d96a4f7a2f64e19e1b8aa846a0fc187bbab7aaf8309478a988b25.deb \
    pool/kubectl_1.11.10-00_amd64_d66a02f63226306fcbb193013174821ea63999dd91a7048fc0613bf4da161efc.deb \
    pool/kubectl_1.12.10-00_amd64_896bcfcacec129e50e5dcf6121ae5c32ff6b53c3fc4c5a082226d74b5fbfb0c5.deb \
    pool/kubectl_1.13.12-00_amd64_b32ad754b6b4910ca709d76ec909bf28e227082a38867870d80205af4631467e.deb \
    pool/kubectl_1.14.9-00_amd64_0f0bb925d3a662e699a135af91406a0e1e7fda6fc8bcee6addaf8cc8b7018e13.deb \
    pool/kubectl_1.15.6-00_amd64_a82e18c2533dc96cda21e03c605f325f352ac6fabbdd09fef43cca97254b436e.deb \
    pool/kubectl_1.16.3-00_amd64_bbca932fd543846177182fcba147e8b5375875038c7d0508f520b3f5a31bf4d9.deb \
    pool/kubelet_1.10.13-00_amd64_aa7069a9b903f263a16bf54059363e75b697b6ef9529c1042f8fe2e238eaf686.deb \
    pool/kubelet_1.11.10-00_amd64_b9d8e8599b9d606930e6d31159843cd60e1c89d978ef1d2123f02985b2922564.deb \
    pool/kubelet_1.12.10-00_amd64_cfdef95a909ed3b1de9e53a4499c7bfeb2e47eed3b6e9d13c0adc8ad60d2076d.deb \
    pool/kubelet_1.13.12-00_amd64_bbb3a620583678b250fa4b64a208a20cc5592bf7ff926ccbc9f8bd92e8515f94.deb \
    pool/kubelet_1.14.9-00_amd64_0997197a365ce053050b65d0d0b88a76059b9ef3b5d0e1ef7d3c4a8d9a29e9e3.deb \
    pool/kubelet_1.15.6-00_amd64_ef88a8ae8a9474be1aee0b508ae51ab990503086061f1a9c83677b67466f15bb.deb \
    pool/kubelet_1.16.3-00_amd64_77c3d05510bc1f753834c30de6a1a7fc032ecfb906527bc270663428b371c477.deb \
    pool/kubernetes-cni_0.7.5-00_amd64_b38a324bb34f923d353203adf0e048f3b911f49fa32f1d82051a71ecfe2cd184.deb \
"

#; Feb 3 2020
pkgs=" \
pool/kubeadm_1.17.2-00_amd64_041e0e6cc71212fff2cdaadfa7fd52d992fad23efd15157d360df86a22cf7649.deb \
pool/kubectl_1.17.2-00_amd64_4cf57f0a5445409293963fddc279fbba9f48c6c280fa27479651d1908848b96a.deb \
pool/kubelet_1.17.2-00_amd64_f22c86b79e7e98d18aaca1690d96b977bca11e9541170716bab0cb664f9e583c.deb \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/apt/$i -o $download_dir/$i
done
