#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/apt"

mkdir -p $download_dir/{dists,doc,pool,src}

wget --timestamping --recursive --level=inf --no-remove-listing \
    --no-check-certificate \
    --accept=kubernetes-xenial,kubernetes-trusty,InRelease,Release,Release.gpg,main,binary-amd64,Packages,Packages.gz,Release \
    --reject=binary-arm64,binary-armhf,binary-i386,binary-ppc64el,binary-s390x,source \
    --directory-prefix=$download_dir/dists \
    --no-host-directories \
    --cut-dirs=2 \
    --base=https://packages.cloud.google.com/apt/dists \
    https://packages.cloud.google.com/apt/dists
