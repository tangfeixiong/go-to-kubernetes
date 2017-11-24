#!/bin/sh

source_home="https://yum.dockerproject.org/repo/main/fedora/23"

sub_path="repodata"

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

target_dir="$base_dir/fedora0x2F23/$sub_path"

mkdir -p $target_dir

pkgs=" \
    17ba1d0122e3bad5506b6d4c8bfef4ad415854620863ca83bd39ca5b4b71746c-filelists.xml.gz \
    242385b31d067f49a9042c1f809fb856f6a5d7eb4fccaaedd0ae2776a638aab3-other.sqlite.bz2 \
    848a650f8b62f4b756edd7b2be158f87141af2834d302e846cc39b6e8829d125-other.xml.gz \
    ba79fb7ce4bfb4b9bd990ac7d46fddab975ac5707f7b48be13e21d1e6d8b930e-primary.xml.gz \
    d9355ebf02875e5624b43098a0567c244cfa63b5b80b7317be891a91db3fa8e3-filelists.sqlite.bz2 \
    f77d37a8d813fb6ae2138e8a0cb876bc068d6f9a76a8f6ac0da80853f22ba1e8-primary.sqlite.bz2 \
    repomd.xml \
    repomd.xml.asc \
"

for i in $pkgs; do
    curl -jkSL $source_home/$sub_path/$i -o $target_dir/$i
done

