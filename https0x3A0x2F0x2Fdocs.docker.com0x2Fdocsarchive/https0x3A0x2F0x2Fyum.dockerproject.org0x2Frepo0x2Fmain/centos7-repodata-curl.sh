#!/bin/sh

source_home="https://yum.dockerproject.org/repo/main/centos/7"

sub_path="repodata"

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

target_dir="$base_dir/centos0x2F7/$sub_path"

mkdir -p $target_dir

pkgs=" \
    0615103afda1f7dfe0638913d6c66d2a52f0e2b15bb4f5a80f9b609d098a6639-filelists.sqlite.bz2 \
    074434cadbea297f9f24da50e2bdd6aaca65f864616e3ccf3e2edd7d476449be-filelists.xml.gz \
    4f1b98c5420d4b7f7fc36f2f6cb57d7f3820b64009efe7ce1f1aa57ab7e4c2e7-primary.sqlite.bz2 \
    ac0f667babb086459d9891a07dcd89833012bf53cdf6d4da7ac92e91dd668b28-other.xml.gz \
    e1404df8a48c9646cae9af4466e0841e084e94d5d74d2df26fb03ff2190670c1-other.sqlite.bz2 \
    f4069fb0c4045248591849a1b66e8f235d6195f10eb5bd6d414355a422ca5ffd-primary.xml.gz \
    repomd.xml \
    repomd.xml.asc \
"
pkgs=" \
    4f1b98c5420d4b7f7fc36f2f6cb57d7f3820b64009efe7ce1f1aa57ab7e4c2e7-primary.sqlite.bz2 \
"

for i in $pkgs; do
    curl -jkSL $source_home/$sub_path/$i -o $target_dir/$i
done

