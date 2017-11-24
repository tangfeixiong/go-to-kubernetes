#!/bin/sh

source_home="https://apt.dockerproject.org/repo"

sub_path="dists/debian-jessie"

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

target_dir="$base_dir/$sub_path"

mkdir -p $target_dir/main/binary-amd64

pkgs=" \
    InRelease \
    Release \
    Release.gpg \
    main/Contents-amd64 \
    main/Contents-amd64.bz2 \
    main/Contents-amd64.gz \
    main/filelist \
    main/binary-amd64/InRelease \
    main/binary-amd64/Packages \
    main/binary-amd64/Packages.bz2 \
    main/binary-amd64/Packages.gz \
    main/binary-amd64/Release \
    main/binary-amd64/Release.gpg \
"

for i in $pkgs; do
    curl -jkSL $source_home/$sub_path/$i -o $target_dir/$i
done

