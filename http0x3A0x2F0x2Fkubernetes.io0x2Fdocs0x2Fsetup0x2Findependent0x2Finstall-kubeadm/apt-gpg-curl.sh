#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/apt"

mkdir -p $download_dir/{dists,doc,pool,src}

curl -jkSL https://packages.cloud.google.com/apt/doc/apt-key.gpg -o $download_dir/doc/apt-key.gpg

curl -jkSL https://packages.cloud.google.com/apt/doc/rpm-package-key.gpg -o $download_dir/doc/rpm-package-key.gpg

curl -kjSL https://packages.cloud.google.com/apt/doc/yum-key.gpg -o $download_dir/doc/yum-key.gpg
