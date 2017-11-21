#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum"

mkdir -p $download_dir/{doc,pool,repos}

targets="kubernetes-el6-x86_64/repodata/filelists.xml \
    kubernetes-el6-x86_64/repodata/filelists.xml.gz \
    kubernetes-el6-x86_64/repodata/other.xml \
    kubernetes-el6-x86_64/repodata/other.xml.gz \
    kubernetes-el6-x86_64/repodata/primary.xml \
    kubernetes-el6-x86_64/repodata/primary.xml.gz \
    kubernetes-el6-x86_64/repodata/repomd.xml \
    kubernetes-el6-x86_64/repodata/repomd.xml.asc \
    kubernetes-el6-x86_64-unstable/repodata/filelists.xml \
    kubernetes-el6-x86_64-unstable/repodata/filelists.xml.gz \
    kubernetes-el6-x86_64-unstable/repodata/other.xml \
    kubernetes-el6-x86_64-unstable/repodata/other.xml.gz \
    kubernetes-el6-x86_64-unstable/repodata/primary.xml \
    kubernetes-el6-x86_64-unstable/repodata/primary.xml.gz \
    kubernetes-el6-x86_64-unstable/repodata/repomd.xml \
    kubernetes-el6-x86_64-unstable/repodata/repomd.xml.asc \
"

targets="kubernetes-el7-x86_64/repodata/filelists.xml \
    kubernetes-el7-x86_64/repodata/filelists.xml.gz \
    kubernetes-el7-x86_64/repodata/other.xml \
    kubernetes-el7-x86_64/repodata/other.xml.gz \
    kubernetes-el7-x86_64/repodata/primary.xml \
    kubernetes-el7-x86_64/repodata/primary.xml.gz \
    kubernetes-el7-x86_64/repodata/repomd.xml \
    kubernetes-el7-x86_64/repodata/repomd.xml.asc \
    kubernetes-el7-x86_64-unstable/repodata/filelists.xml \
    kubernetes-el7-x86_64-unstable/repodata/filelists.xml.gz \
    kubernetes-el7-x86_64-unstable/repodata/other.xml \
    kubernetes-el7-x86_64-unstable/repodata/other.xml.gz \
    kubernetes-el7-x86_64-unstable/repodata/primary.xml \
    kubernetes-el7-x86_64-unstable/repodata/primary.xml.gz \
    kubernetes-el7-x86_64-unstable/repodata/repomd.xml \
    kubernetes-el7-x86_64-unstable/repodata/repomd.xml.asc \
"

targets="kubernetes-el7-x86_64/repodata/filelists.xml \
    kubernetes-el7-x86_64/repodata/filelists.xml.gz \
    kubernetes-el7-x86_64/repodata/other.xml \
    kubernetes-el7-x86_64/repodata/other.xml.gz \
    kubernetes-el7-x86_64/repodata/primary.xml \
    kubernetes-el7-x86_64/repodata/primary.xml.gz \
    kubernetes-el7-x86_64/repodata/repomd.xml \
    kubernetes-el7-x86_64/repodata/repomd.xml.asc \
"

for i in $targets; do
    mkdir -p $download_dir/repos/$(dirname $i)
    curl -kjSL https://packages.cloud.google.com/yum/repos/$i -o $download_dir/repos/$i
done
