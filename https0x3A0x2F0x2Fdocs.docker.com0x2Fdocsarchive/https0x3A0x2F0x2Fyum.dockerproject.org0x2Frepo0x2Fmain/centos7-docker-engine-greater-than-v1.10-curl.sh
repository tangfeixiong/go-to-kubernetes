#!/bin/sh

source_home="https://yum.dockerproject.org/repo/main/centos/7"

sub_path="Packages"

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

target_dir="$base_dir/centos0x2F7/$sub_path"

mkdir -p $target_dir

pkgs=" \
    docker-engine-1.11.2-1.el7.centos.x86_64.rpm \
    docker-engine-1.12.6-1.el7.centos.x86_64.rpm \
    docker-engine-1.13.1-1.el7.centos.x86_64.rpm \
    docker-engine-17.05.0.ce-1.el7.centos.x86_64.rpm \
    docker-engine-debuginfo-1.11.2-1.el7.centos.x86_64.rpm \
    docker-engine-debuginfo-1.12.6-1.el7.centos.x86_64.rpm \
    docker-engine-debuginfo-1.13.1-1.el7.centos.x86_64.rpm \
    docker-engine-debuginfo-17.05.0.ce-1.el7.centos.x86_64.rpm \
    docker-engine-selinux-1.11.2-1.el7.centos.noarch.rpm \
    docker-engine-selinux-1.12.6-1.el7.centos.noarch.rpm \
    docker-engine-selinux-1.13.1-1.el7.centos.noarch.rpm \
    docker-engine-selinux-17.05.0.ce-1.el7.centos.noarch.rpm \
"

for i in $pkgs; do
    curl -jkSL $source_home/$sub_path/$i -o $target_dir/$i
done

