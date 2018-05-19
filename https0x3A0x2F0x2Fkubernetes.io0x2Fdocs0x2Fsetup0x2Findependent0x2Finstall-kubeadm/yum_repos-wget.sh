#!/bin/bash

### Deprecated

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum"

mkdir -p $download_dir/{doc,pool,repos}

wget --mirror \
    --directory-prefix=$download_dir/pool \
    --no-host-directories \
    --cut-dirs=2 \
    --no-check-certificate \
    https://packages.cloud.google.com/yum/pool
    
#wget --spider
#    --base=https://packages.cloud.google.com/yum/pool
