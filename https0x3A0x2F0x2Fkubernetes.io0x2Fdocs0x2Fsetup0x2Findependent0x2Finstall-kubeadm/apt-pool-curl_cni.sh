#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/apt"

mkdir -p $download_dir/{dists,doc,pool,src}

pkgs="pool/docker-engine_1.11.2-0~xenial_amd64_5554a8bc383e65fb10d556239c72b26457cc5d97f49e3a353c3382f8434e7d21.deb \
    pool/kubernetes-cni_0.3.0.1-07a8a2-00_amd64_9e41a275b2afeb51dcde86b922c056c7b6dc656b54dd66fa2f1a0bb8266e9c22.deb \
    pool/kubernetes-cni_0.5.1-00_amd64_08cbe5c42366ec3184cc91a4353f6e066f2d7325b4db1cb4f87c7dcc8c0eb620.deb \
    pool/kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb \
"

for i in $pkgs; do
    curl -jkSL https://packages.cloud.google.com/apt/$i -o $download_dir/$i
done
