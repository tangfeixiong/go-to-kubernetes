#!/bin/sh

source_home="https://apt.dockerproject.org/repo"

sub_path="db"

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

target_dir="$base_dir/$sub_path"

mkdir -p $target_dir

pkgs=" \
    packages-main-amd64.db \
"

for i in $pkgs; do
    curl -jkSL $source_home/$sub_path/$i -o $target_dir/$i
done

