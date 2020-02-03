#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/apt"

mkdir -p $download_dir/{dists,doc,pool,src}

targets=" \
    kubernetes-jessie/InRelease \
    kubernetes-jessie/Release \
    kubernetes-jessie/Release.gpg \
    kubernetes-jessie/main/binary-amd64/Packages \
    kubernetes-jessie/main/binary-amd64/Packages.gz \
    kubernetes-jessie/main/binary-amd64/Release \
    kubernetes-jessie/main/debian-installer/binary-amd64/Packages \
    kubernetes-jessie/main/debian-installer/binary-amd64/Packages.gz \
    kubernetes-stretch/InRelease \
    kubernetes-stretch/Release \
    kubernetes-stretch/Release.gpg \
    kubernetes-stretch/main/binary-amd64/Packages \
    kubernetes-stretch/main/binary-amd64/Packages.gz \
    kubernetes-stretch/main/binary-amd64/Release \
    kubernetes-stretch/main/debian-installer/binary-amd64/Packages \
    kubernetes-stretch/main/debian-installer/binary-amd64/Packages.gz \
    kubernetes-trusty/InRelease \
    kubernetes-trusty/Release \
    kubernetes-trusty/Release.gpg \
    kubernetes-trusty/main/binary-amd64/Packages \
    kubernetes-trusty/main/binary-amd64/Packages.gz \
    kubernetes-trusty/main/binary-amd64/Release \
    kubernetes-trusty/main/debian-installer/binary-amd64/Packages \
    kubernetes-trusty/main/debian-installer/binary-amd64/Packages.gz \
    kubernetes-xenial/InRelease \
    kubernetes-xenial/Release \
    kubernetes-xenial/Release.gpg \
    kubernetes-xenial/main/binary-amd64/Packages \
    kubernetes-xenial/main/binary-amd64/Packages.gz \
    kubernetes-xenial/main/binary-amd64/Release \
    kubernetes-xenial/main/debian-installer/binary-amd64/Packages \
    kubernetes-xenial/main/debian-installer/binary-amd64/Packages.gz \
    kubernetes-yakkety/InRelease \
    kubernetes-yakkety/Release \
    kubernetes-yakkety/Release.gpg \
    kubernetes-yakkety/main/binary-amd64/Packages \
    kubernetes-yakkety/main/binary-amd64/Packages.gz \
    kubernetes-yakkety/main/binary-amd64/Release \
    kubernetes-yakkety/main/debian-installer/binary-amd64/Packages \
    kubernetes-yakkety/main/debian-installer/binary-amd64/Packages.gz \
"

targets=" \
    kubernetes-trusty/InRelease \
    kubernetes-trusty/Release \
    kubernetes-trusty/Release.gpg \
    kubernetes-trusty/main/binary-amd64/Packages \
    kubernetes-trusty/main/binary-amd64/Packages.gz \
    kubernetes-trusty/main/binary-amd64/Release \
    kubernetes-trusty/main/debian-installer/binary-amd64/Packages \
    kubernetes-trusty/main/debian-installer/binary-amd64/Packages.gz \
    kubernetes-xenial/InRelease \
    kubernetes-xenial/Release \
    kubernetes-xenial/Release.gpg \
    kubernetes-xenial/main/binary-amd64/Packages \
    kubernetes-xenial/main/binary-amd64/Packages.gz \
    kubernetes-xenial/main/binary-amd64/Release \
    kubernetes-xenial/main/debian-installer/binary-amd64/Packages \
    kubernetes-xenial/main/debian-installer/binary-amd64/Packages.gz \
"

targets=" \
    kubernetes-xenial/InRelease \
    kubernetes-xenial/Release \
    kubernetes-xenial/Release.gpg \
    kubernetes-xenial/main/binary-amd64/Packages \
    kubernetes-xenial/main/binary-amd64/Packages.gz \
    kubernetes-xenial/main/binary-amd64/Release \
    kubernetes-xenial/main/debian-installer/binary-amd64/Packages \
    kubernetes-xenial/main/debian-installer/binary-amd64/Packages.gz \
"

for i in $targets; do
    mkdir -p $download_dir/dists/$(dirname $i)
    curl -kjSL https://packages.cloud.google.com/apt/dists/$i -o $download_dir/dists/$i
done
