#!/bin/sh

source_home="https://apt.dockerproject.org/repo"

sub_path="conf"

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

target_dir="$base_dir/$sub_path"

mkdir -p $target_dir

pkgs=" \
    apt-ftparchive.conf \
    docker-engine-release.conf \
"

for i in $pkgs; do
    curl -jkSL $source_home/$sub_path/$i -o $target_dir/$i
done

