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
[root@localhost yum.repos.d]# yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline install docker
已加载插件：fastestmirror
Loading mirror speeds from cached hostfile
 * c7-media: 
 * offline: 
正在解决依赖关系
--> 正在检查事务
---> 软件包 docker.x86_64.2.1.10.3-59.el7.centos 将被 安装
--> 正在处理依赖关系 docker-common = 2:1.10.3-59.el7.centos，它被软件包 2:docker-1.10.3-59.el7.centos.x86_64 需要
--> 正在处理依赖关系 oci-systemd-hook >= 1:0.1.4-5，它被软件包 2:docker-1.10.3-59.el7.centos.x86_64 需要
--> 正在处理依赖关系 oci-register-machine >= 1:0-1.8，它被软件包 2:docker-1.10.3-59.el7.centos.x86_64 需要
--> 正在处理依赖关系 container-selinux >= 2:1.10.3-59.el7.centos，它被软件包 2:docker-1.10.3-59.el7.centos.x86_64 需要
--> 正在处理依赖关系 libseccomp.so.2()(64bit)，它被软件包 2:docker-1.10.3-59.el7.centos.x86_64 需要
--> 正在检查事务
---> 软件包 container-selinux.x86_64.2.1.10.3-59.el7.centos 将被 安装
--> 正在处理依赖关系 policycoreutils-python，它被软件包 2:container-selinux-1.10.3-59.el7.centos.x86_64 需要
---> 软件包 docker-common.x86_64.2.1.10.3-59.el7.centos 将被 安装
---> 软件包 libseccomp.x86_64.0.2.3.1-2.el7 将被 安装
---> 软件包 oci-register-machine.x86_64.1.0-1.10.gitfcdbff0.el7 将被 安装
---> 软件包 oci-systemd-hook.x86_64.1.0.1.4-6.git337078c.el7 将被 安装
--> 正在处理依赖关系 libyajl.so.2()(64bit)，它被软件包 1:oci-systemd-hook-0.1.4-6.git337078c.el7.x86_64 需要
--> 正在检查事务
---> 软件包 policycoreutils-python.x86_64.0.2.5-8.el7 将被 安装
--> 正在处理依赖关系 setools-libs >= 3.3.8-1，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 libsemanage-python >= 2.5-4，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 audit-libs-python >= 2.1.3-4，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 python-IPy，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 libqpol.so.1(VERS_1.4)(64bit)，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 libqpol.so.1(VERS_1.2)(64bit)，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 libcgroup，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 libapol.so.4(VERS_4.0)(64bit)，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 checkpolicy，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 libqpol.so.1()(64bit)，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
--> 正在处理依赖关系 libapol.so.4()(64bit)，它被软件包 policycoreutils-python-2.5-8.el7.x86_64 需要
---> 软件包 yajl.x86_64.0.2.0.4-4.el7 将被 安装
--> 正在检查事务
---> 软件包 audit-libs-python.x86_64.0.2.6.5-3.el7 将被 安装
---> 软件包 checkpolicy.x86_64.0.2.5-4.el7 将被 安装
---> 软件包 libcgroup.x86_64.0.0.41-11.el7 将被 安装
---> 软件包 libsemanage-python.x86_64.0.2.5-4.el7 将被 安装
---> 软件包 python-IPy.noarch.0.0.75-6.el7 将被 安装
---> 软件包 setools-libs.x86_64.0.3.3.8-1.1.el7 将被 安装
--> 解决依赖关系完成

依赖关系解决

====================================================================================================================
 Package                           架构              版本                                 源                   大小
====================================================================================================================
正在安装:
 docker                            x86_64            2:1.10.3-59.el7.centos               offline              12 M
为依赖而安装:
 audit-libs-python                 x86_64            2.6.5-3.el7                          c7-media             70 k
 checkpolicy                       x86_64            2.5-4.el7                            c7-media            290 k
 container-selinux                 x86_64            2:1.10.3-59.el7.centos               offline              80 k
 docker-common                     x86_64            2:1.10.3-59.el7.centos               offline              63 k
 libcgroup                         x86_64            0.41-11.el7                          c7-media             65 k
 libseccomp                        x86_64            2.3.1-2.el7                          c7-media             56 k
 libsemanage-python                x86_64            2.5-4.el7                            c7-media            103 k
 oci-register-machine              x86_64            1:0-1.10.gitfcdbff0.el7              offline             1.1 M
 oci-systemd-hook                  x86_64            1:0.1.4-6.git337078c.el7             offline              28 k
 policycoreutils-python            x86_64            2.5-8.el7                            c7-media            444 k
 python-IPy                        noarch            0.75-6.el7                           c7-media             32 k
 setools-libs                      x86_64            3.3.8-1.1.el7                        c7-media            612 k
 yajl                              x86_64            2.0.4-4.el7                          c7-media             39 k

事务概要
====================================================================================================================
安装  1 软件包 (+13 依赖软件包)

总下载量：15 M
安装大小：64 M
Is this ok [y/d/N]: y
Downloading packages:
--------------------------------------------------------------------------------------------------------------------
总计                                                                                 49 MB/s |  15 MB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在安装    : setools-libs-3.3.8-1.1.el7.x86_64                                                              1/14 
  正在安装    : libsemanage-python-2.5-4.el7.x86_64                                                            2/14 
  正在安装    : checkpolicy-2.5-4.el7.x86_64                                                                   3/14 
  正在安装    : 1:oci-register-machine-0-1.10.gitfcdbff0.el7.x86_64                                            4/14 
  正在安装    : yajl-2.0.4-4.el7.x86_64                                                                        5/14 
  正在安装    : 1:oci-systemd-hook-0.1.4-6.git337078c.el7.x86_64                                               6/14 
  正在安装    : libcgroup-0.41-11.el7.x86_64                                                                   7/14 
  正在安装    : python-IPy-0.75-6.el7.noarch                                                                   8/14 
  正在安装    : audit-libs-python-2.6.5-3.el7.x86_64                                                           9/14 
  正在安装    : policycoreutils-python-2.5-8.el7.x86_64                                                       10/14 
  正在安装    : 2:container-selinux-1.10.3-59.el7.centos.x86_64                                               11/14 
libsemanage.semanage_direct_remove_key: Unable to remove module docker at priority 200. (No such file or directory).
libsemanage.semanage_direct_remove_key: Unable to remove module docker at priority 400. (No such file or directory).
  正在安装    : 2:docker-common-1.10.3-59.el7.centos.x86_64                                                   12/14 
  正在安装    : libseccomp-2.3.1-2.el7.x86_64                                                                 13/14 
  正在安装    : 2:docker-1.10.3-59.el7.centos.x86_64                                                          14/14 
  验证中      : libseccomp-2.3.1-2.el7.x86_64                                                                  1/14 
  验证中      : 2:docker-common-1.10.3-59.el7.centos.x86_64                                                    2/14 
  验证中      : 2:docker-1.10.3-59.el7.centos.x86_64                                                           3/14 
  验证中      : policycoreutils-python-2.5-8.el7.x86_64                                                        4/14 
  验证中      : audit-libs-python-2.6.5-3.el7.x86_64                                                           5/14 
  验证中      : python-IPy-0.75-6.el7.noarch                                                                   6/14 
  验证中      : 2:container-selinux-1.10.3-59.el7.centos.x86_64                                                7/14 
  验证中      : libcgroup-0.41-11.el7.x86_64                                                                   8/14 
  验证中      : yajl-2.0.4-4.el7.x86_64                                                                        9/14 
  验证中      : 1:oci-register-machine-0-1.10.gitfcdbff0.el7.x86_64                                           10/14 
  验证中      : checkpolicy-2.5-4.el7.x86_64                                                                  11/14 
  验证中      : 1:oci-systemd-hook-0.1.4-6.git337078c.el7.x86_64                                              12/14 
  验证中      : libsemanage-python-2.5-4.el7.x86_64                                                           13/14 
  验证中      : setools-libs-3.3.8-1.1.el7.x86_64                                                             14/14 

已安装:
  docker.x86_64 2:1.10.3-59.el7.centos                                                                              

作为依赖被安装:
  audit-libs-python.x86_64 0:2.6.5-3.el7                  checkpolicy.x86_64 0:2.5-4.el7                            
  container-selinux.x86_64 2:1.10.3-59.el7.centos         docker-common.x86_64 2:1.10.3-59.el7.centos               
  libcgroup.x86_64 0:0.41-11.el7                          libseccomp.x86_64 0:2.3.1-2.el7                           
  libsemanage-python.x86_64 0:2.5-4.el7                   oci-register-machine.x86_64 1:0-1.10.gitfcdbff0.el7       
  oci-systemd-hook.x86_64 1:0.1.4-6.git337078c.el7        policycoreutils-python.x86_64 0:2.5-8.el7                 
  python-IPy.noarch 0:0.75-6.el7                          setools-libs.x86_64 0:3.3.8-1.1.el7                       
  yajl.x86_64 0:2.0.4-4.el7                              

完毕！

[root@localhost yum.repos.d]# usermod -aG dockerroot ecp
[root@localhost yum.repos.d]# cat /etc/group | grep dockerroot
dockerroot:x:993:ecp
[vagrant@openshiftdev yum.repos.d]$ logout
[ecp@localhost ~]$ logout
Connection to 192.168.2.20 closed.
localhost:puhua-yanyong fanhongling$ ssh ecp@192.168.2.20
ecp@192.168.2.20's password: 
Last login: Tue Feb 21 16:51:42 2017 from 192.168.2.54
[ecp@localhost ~]$ sudo -i
[sudo] password for ecp: 
[ecp@localhost ~]$ sudo systemctl enable docker.service
[sudo] password for ecp: 
Created symlink from /etc/systemd/system/multi-user.target.wants/docker.service to /usr/lib/systemd/system/docker.service.
[ecp@localhost ~]$ sudo systemctl start docker.service

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