#!/bin/sh
base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

curl -jkSL https://yum.dockerproject.org/gpg -o $base_dir/https0x3A0x2F0x2Fyum.dockerproject.org0x2Fgpg
