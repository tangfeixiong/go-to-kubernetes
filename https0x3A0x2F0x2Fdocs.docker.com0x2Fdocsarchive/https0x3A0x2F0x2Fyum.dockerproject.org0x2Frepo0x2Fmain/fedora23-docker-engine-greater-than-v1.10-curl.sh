#!/bin/sh

source_home="https://yum.dockerproject.org/repo/main/fedora/23"

sub_path="Packages"

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

target_dir="$base_dir/fedora0x2F23/$sub_path"

mkdir -p $target_dir

pkgs=" \
    docker-engine-1.11.2-1.fc23.x86_64.rpm \
    docker-engine-1.12.6-1.fc23.x86_64.rpm \
    docker-engine-debuginfo-1.11.2-1.fc23.x86_64.rpm \
    docker-engine-debuginfo-1.12.6-1.fc23.x86_64.rpm \
    docker-engine-selinux-1.11.2-1.fc23.noarch.rpm \
    docker-engine-selinux-1.12.6-1.fc23.noarch.rpm \
"

for i in $pkgs; do
    curl -jkSL $source_home/$sub_path/$i -o $target_dir/$i
done

