Install docker engine - 动手安装docker engine
=============================================

Tables of Content
-------------------

Using OS-Media repo

Installation

The docker Repo
---------------------------------

User guide - https://docs.docker.com/engine/installation/

### Using OS-Media Repo - 离线包仓库
-----------------------------------

For detailed content, refer to [mirror-centos-repo.md](./mirror-centos-repo.md). 有关离线包的准备可参考[mirror-centos-repo.md](./mirror-centos-repo.md)

```
[vagrant@openshiftdev yum.repos.d]$ sudo yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline list | grep docker
docker.x86_64                             1.10.3-46.el7.centos.14      @extras
docker-common.x86_64                      1.10.3-46.el7.centos.14      @extras
docker-selinux.x86_64                     1.10.3-46.el7.centos.14      @extras
cockpit-docker.x86_64                     122-3.el7.centos             offline
docker.x86_64                             2:1.10.3-59.el7.centos       offline
docker-common.x86_64                      2:1.10.3-59.el7.centos       offline
docker-devel.x86_64                       1.3.2-4.el7.centos           offline
docker-distribution.x86_64                2.5.1-1.el7                  offline
docker-forward-journald.x86_64            1.10.3-44.el7.centos         offline
docker-latest.x86_64                      1.12.3-10.el7.centos         offline
docker-latest-logrotate.x86_64            1.12.3-10.el7.centos         offline
docker-latest-v1.10-migrator.x86_64       1.12.3-10.el7.centos         offline
docker-logrotate.x86_64                   2:1.10.3-59.el7.centos       offline
docker-lvm-plugin.x86_64                  2:1.10.3-59.el7.centos       offline
docker-novolume-plugin.x86_64             2:1.10.3-59.el7.centos       offline
docker-python.x86_64                      1.4.0-115.el7                offline
docker-registry.noarch                    0.6.8-8.el7                  offline
docker-registry.x86_64                    0.9.1-7.el7                  offline
docker-unit-test.x86_64                   2:1.10.3-59.el7.centos       offline
docker-v1.10-migrator.x86_64              2:1.10.3-59.el7.centos       offline
python-docker-py.noarch                   1.9.0-1.el7                  offline

```

### Installation - 安装过程
----------------------------

Run _yum_ to install and authorize user as docker client

```
[vagrant@openshiftdev yum.repos.d]$ sudo yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline install -y docker
[vagrant@openshiftdev yum.repos.d]$ sudo groupadd -aG dockerroot vagrant
[vagrant@openshiftdev yum.repos.d]$ sudo cat /etc/group | grep dockerroot
[vagrant@openshiftdev yum.repos.d]$ logout

```

Load docker image archive

```
[vagrant@openshiftdev yum.repos.d]$ docker load --help

Usage:	docker load [OPTIONS]

Load an image from a tar archive or STDIN

  --help=false       Print usage
  -i, --input=       Read from a tar archive file, instead of STDIN

[vagrant@openshiftdev yum.repos.d]$ docker load -i tangfeixiong%2Fnetcat-hello-http%3Agitrev-7257e93.tar

```

Play

```
[vagrant@openshiftdev yum.repos.d]$ docker run -d -p 80:80 --name=hello docker.io/tangfeixiong/net-hello-http:gitrev-7257e93

[vagrant@openshiftdev yum.repos.d]$ curl 127.0.0.1

```

Open browser to view example html