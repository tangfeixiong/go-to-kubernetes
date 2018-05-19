#!/bin/bash

set -e

working_dir=`pwd`

base_dir=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

download_dir="$base_dir/https0x3A0x2F0x2Fpackages.cloud.google.com/yum"

if [ -x /bin/docker -o -x /usr/bin/docker -o -x /usr/local/bin/docker ]; then
    if [[ $(docker version -f {{.Client.Version}}) =~ [1-9]\.[1-9][0-9]?\.[0-9]+ ]]; then

mkdir -p $download_dir/{doc,pool,repos/kubernetes-el7-x86_64/repodata}

curl -jkSL https://packages.cloud.google.com/yum/doc/apt-key.gpg -o $download_dir/doc/apt-key.gpg

curl -jkSL https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg -o $download_dir/doc/rpm-package-key.gpg

curl -kjSL https://packages.cloud.google.com/yum/doc/yum-key.gpg -o $download_dir/doc/yum-key.gpg

curl -kjSL https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/filelists.xml -o $download_dir/repos/kubernetes-el7-x86_64/repodata/filelists.xml

curl -kjSL https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/filelists.xml.gz -o $download_dir/repos/kubernetes-el7-x86_64/repodata/filelists.xml.gz

curl -kjSL https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/other.xml -o $download_dir/repos/kubernetes-el7-x86_64/repodata/other.xml

curl -kjSL https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/other.xml.gz -o $download_dir/repos/kubernetes-el7-x86_64/repodata/other.xml.gz

curl -kjSL https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/primary.xml -o $download_dir/repos/kubernetes-el7-x86_64/repodata/primary.xml

curl -kjSL https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/primary.xml.gz -o $download_dir/repos/kubernetes-el7-x86_64/repodata/primary.xml.gz

curl -kjSL https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/repomd.xml -o $download_dir/repos/kubernetes-el7-x86_64/repodata/repomd.xml

curl -kjSL https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/repomd.xml.asc -o $download_dir/repos/kubernetes-el7-x86_64/repodata/repomd.xml.asc

install_pkg=" \
    kubeadm \
    kubectl \
    kubelet \
    kubernetes-cni \
    ###rkt
"
# yum repo-pkgs kubernetes install $install_pkg

# docker run --rm -t -v $base_dir/kubernetes-el7-x86_64:/tmp/download centos:centos7 \
#     "yum --downloadonly --downloaddir=/tmp/download install $install_pkg"

docker run --rm -t -v $base_dir/kubernetes.repo:/etc/yum.repos.d/kubernets.repo \
    -v $base_dir/kubernetes-el7-x86_64:/var/cache/yum centos:centos7 \
    yum --downloadonly install $install_pkg

# docker run --rm -t -v $base_dir/kubernetes-el7-x86_64:/tmp/download centos:centos7 \
#     /bin/bash -c " \
# 	    mkdir -p /var/cache/yum/x86_64 \
#         && ln -s /tmp/download /var/cache/yum/x86_64/7 \
#         && export install_pkg=' \
#             glibc-common \
#         ' \
# 	    && yum --downloadonly install \$install_pkg \
# 	  "

    fi
else

base_dir+="/rpm"
mkdir -p $base_dir

basearch=x86_64

baseurl=http://mirror.centos.org/centos/7/os/$basearch/Packages/
updateurl=http://mirror.centos.org/centos/7/updates/$basearch/Packages/

install_rpm=" \
	lzo-2.06-8.el7.x86_64.rpm \
	unzip-6.0-15.el7.x86_64.rpm \
" 

for i in $install_rpms; do
    [ ! -f "$base_dir/$i" ] && curl -jkSL "${baseurl}$i" -o "$base_dir/$i";
done

install_rpm=" \
	bsdtar-3.1.2-10.el7_2.x86_64.rpm \
	libarchive-3.1.2-10.el7_2.x86_64.rpm \
" 

for i in $install_rpms; do
    [ ! -f "$base_dir/$i" ] && curl -jkSL "${updateurl}$i" -o "$base_dir/$i";
done

fi

cd $working_dir
