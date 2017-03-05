Install ETCD 安装ETCD配置管理服务
===============================

Install
--------

About etcd2 package

    [vagrant@localhost ~]$ sudo yum -v list etcd
    Loading "fastestmirror" plugin
    Config time: 0.005
    Yum version: 3.4.3
    rpmdb time: 0.000
    Setting up Package Sacks
    Loading mirror speeds from cached hostfile
     * base: mirrors.zju.edu.cn
     * extras: mirrors.zju.edu.cn
     * updates: mirrors.zju.edu.cn
    pkgsack time: 0.007
    Available Packages
    etcd.x86_64                          2.3.7-4.el7                          extras

    [vagrant@localhost ~]$ sudo yum deplist etcd
    Loaded plugins: fastestmirror
    Loading mirror speeds from cached hostfile
     * base: mirrors.zju.edu.cn
     * extras: mirrors.zju.edu.cn
     * updates: mirrors.zju.edu.cn
    package: etcd.x86_64 2.3.7-4.el7
      dependency: /bin/sh
       provider: bash.x86_64 4.2.46-20.el7_2
      dependency: shadow-utils
       provider: shadow-utils.x86_64 2:4.1.5.1-18.el7
      dependency: systemd
       provider: systemd.x86_64 219-19.el7_2.13

Using _offline_ repo

```
[root@localhost ecp]# yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline install etcd
已加载插件：fastestmirror
c7-media                                                                                       | 3.6 kB  00:00:00     
offline                                                                                        | 3.4 kB  00:00:00     
Loading mirror speeds from cached hostfile
 * c7-media: 
 * offline: 
正在解决依赖关系
--> 正在检查事务
---> 软件包 etcd.x86_64.0.2.3.7-4.el7 将被 安装
--> 解决依赖关系完成

依赖关系解决

======================================================================================================================
 Package                 架构                      版本                              源                          大小
======================================================================================================================
正在安装:
 etcd                    x86_64                    2.3.7-4.el7                       offline                    6.5 M

事务概要
======================================================================================================================
安装  1 软件包

总下载量：6.5 M
安装大小：31 M
Is this ok [y/d/N]: y
Downloading packages:
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在安装    : etcd-2.3.7-4.el7.x86_64                                                                           1/1 
  验证中      : etcd-2.3.7-4.el7.x86_64                                                                           1/1 

已安装:
  etcd.x86_64 0:2.3.7-4.el7                                                                                           

完毕！

```

Deep dive into etcd2

    [vagrant@localhost ~]$ sudo systemctl list-unit-files etcd.service
    UNIT FILE    STATE
    etcd.service disabled

    1 unit files listed.

    [vagrant@localhost ~]$ cat /usr/lib/systemd/system/etcd.service
    [Unit]
    Description=Etcd Server
    After=network.target
    After=network-online.target
    Wants=network-online.target

    [Service]
    Type=notify
    WorkingDirectory=/var/lib/etcd/
    EnvironmentFile=-/etc/etcd/etcd.conf
    User=etcd
    # set GOMAXPROCS to number of processors
    ExecStart=/bin/bash -c "GOMAXPROCS=$(nproc) /usr/bin/etcd --name=\"${ETCD_NAME}\" --data-dir=\"${ETCD_DATA_DIR}\" --listen-client-urls=\"${ETCD_LISTEN_CLIENT_URLS}\""
    Restart=on-failure
    LimitNOFILE=65536

    [Install]
    WantedBy=multi-user.target

    [vagrant@localhost ~]$ ls -ld /var/lib/etcd
    drwxr-xr-x. 2 etcd etcd 6 Sep 16 13:18 /var/lib/etcd

    [vagrant@localhost ~]$ cat /etc/passwd | grep etcd
    etcd:x:994:990:etcd user:/var/lib/etcd:/sbin/nologin

Daemon option

```
[root@localhost ecp]# vi /etc/etcd/etcd.conf 

```

Detail option
```
[root@localhost ecp]# cat /etc/etcd/etcd.conf 
# [member]
ETCD_NAME=default
ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
#ETCD_WAL_DIR=""
#ETCD_SNAPSHOT_COUNT="10000"
#ETCD_HEARTBEAT_INTERVAL="100"
#ETCD_ELECTION_TIMEOUT="1000"
# ETCD_LISTEN_PEER_URLS="http://localhost:2380"
ETCD_LISTEN_PEER_URLS="http://192.168.2.20:2380"
# ETCD_LISTEN_CLIENT_URLS="http://localhost:2379"
ETCD_LISTEN_CLIENT_URLS="http://localhost:2379,http://192.168.2.20:2379"
#ETCD_MAX_SNAPSHOTS="5"
#ETCD_MAX_WALS="5"
#ETCD_CORS=""
#
#[cluster]
# ETCD_INITIAL_ADVERTISE_PEER_URLS="http://localhost:2380"
ETCD_INITIAL_ADVERTISE_PEER_URLS="http://192.168.2.20:2380"
# if you use different ETCD_NAME (e.g. test), set ETCD_INITIAL_CLUSTER value for this name, i.e. "test=http://..."
# ETCD_INITIAL_CLUSTER="default=http://localhost:2380"
ETCD_INITIAL_CLUSTER="default=http://192.168.2.20:2380"
#ETCD_INITIAL_CLUSTER_STATE="new"
#ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
# ETCD_ADVERTISE_CLIENT_URLS="http://localhost:2379"
ETCD_ADVERTISE_CLIENT_URLS="http://192.168.2.20:2379"
#ETCD_DISCOVERY=""
#ETCD_DISCOVERY_SRV=""
#ETCD_DISCOVERY_FALLBACK="proxy"
#ETCD_DISCOVERY_PROXY=""
#ETCD_STRICT_RECONFIG_CHECK="false"
#
#[proxy]
#ETCD_PROXY="off"
#ETCD_PROXY_FAILURE_WAIT="5000"
#ETCD_PROXY_REFRESH_INTERVAL="30000"
#ETCD_PROXY_DIAL_TIMEOUT="1000"
#ETCD_PROXY_WRITE_TIMEOUT="5000"
#ETCD_PROXY_READ_TIMEOUT="0"
#
#[security]
#ETCD_CERT_FILE=""
#ETCD_KEY_FILE=""
#ETCD_CLIENT_CERT_AUTH="false"
#ETCD_TRUSTED_CA_FILE=""
#ETCD_PEER_CERT_FILE=""
#ETCD_PEER_KEY_FILE=""
#ETCD_PEER_CLIENT_CERT_AUTH="false"
#ETCD_PEER_TRUSTED_CA_FILE=""
#
#[logging]
#ETCD_DEBUG="false"
# examples for -log-package-levels etcdserver=WARNING,security=DEBUG
#ETCD_LOG_PACKAGE_LEVELS=""
#
#[profiling]
#ETCD_ENABLE_PPROF="false"

```

Start Daemon

```
[root@localhost ecp]# systemctl enable etcd.service
Created symlink from /etc/systemd/system/multi-user.target.wants/etcd.service to /usr/lib/systemd/system/etcd.service.

[root@localhost ecp]# systemctl start etcd.service

[root@localhost ecp]# systemctl -l status etcd.service
● etcd.service - Etcd Server
   Loaded: loaded (/usr/lib/systemd/system/etcd.service; enabled; vendor preset: disabled)
   Active: active (running) since 三 2017-02-22 12:47:52 CST; 1min 7s ago
 Main PID: 11648 (etcd)
   Memory: 2.9M
   CGroup: /system.slice/etcd.service
           └─11648 /usr/bin/etcd --name=default --data-dir=/var/lib/etcd/default.etcd --listen-client-urls=http://localhost:2379,http://192.168.2.20:2379

2月 22 12:47:52 localhost.localdomain systemd[1]: Started Etcd Server.
2月 22 12:47:52 localhost.localdomain etcd[11648]: added local member 9bfd26ae5d19811c [http://192.168.2.20:2380] to cluster aad0dfda8ae36082
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c is starting a new election at term 1
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c became candidate at term 2
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c received vote from 9bfd26ae5d19811c at term 2
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c became leader at term 2
2月 22 12:47:52 localhost.localdomain etcd[11648]: raft.node: 9bfd26ae5d19811c elected leader 9bfd26ae5d19811c at term 2
2月 22 12:47:52 localhost.localdomain etcd[11648]: published {Name:default ClientURLs:[http://192.168.2.20:2379]} to cluster aad0dfda8ae36082
2月 22 12:47:52 localhost.localdomain etcd[11648]: setting up the initial cluster version to 2.3
2月 22 12:47:52 localhost.localdomain etcd[11648]: set the initial cluster version to 2.3

```

CLI

```
[root@localhost ecp]# curl http://127.0.0.1:2379/version
{"etcdserver":"2.3.7","etcdcluster":"2.3.0"}

[root@localhost ecp]# which etcdctl
/usr/bin/etcdctl

[root@localhost ecp]# etcdctl ls

[root@localhost ecp]# etcdctl member list
9bfd26ae5d19811c: name=default peerURLs=http://192.168.2.20:2380 clientURLs=http://192.168.2.20:2379 isLeader=true

[root@localhost ecp]# etcdctl cluster-health
member 9bfd26ae5d19811c is healthy: got healthy result from http://192.168.2.20:2379
cluster is healthy

```

Journal

```
[root@localhost ecp]# journalctl --no-pager --no-tail --pager-end -u etcd.service
-- Logs begin at 二 2017-02-21 11:23:50 CST, end at 三 2017-02-22 12:47:52 CST. --
2月 22 12:47:51 localhost.localdomain systemd[1]: Starting Etcd Server...
2月 22 12:47:52 localhost.localdomain etcd[11648]: recognized and used environment variable ETCD_ADVERTISE_CLIENT_URLS=http://192.168.2.20:2379
2月 22 12:47:52 localhost.localdomain etcd[11648]: recognized and used environment variable ETCD_INITIAL_ADVERTISE_PEER_URLS=http://192.168.2.20:2380
2月 22 12:47:52 localhost.localdomain etcd[11648]: recognized and used environment variable ETCD_INITIAL_CLUSTER=default=http://192.168.2.20:2380
2月 22 12:47:52 localhost.localdomain etcd[11648]: recognized and used environment variable ETCD_LISTEN_PEER_URLS=http://192.168.2.20:2380
2月 22 12:47:52 localhost.localdomain etcd[11648]: recognized environment variable ETCD_NAME, but unused: shadowed by corresponding flag 
2月 22 12:47:52 localhost.localdomain etcd[11648]: recognized environment variable ETCD_DATA_DIR, but unused: shadowed by corresponding flag 
2月 22 12:47:52 localhost.localdomain etcd[11648]: recognized environment variable ETCD_LISTEN_CLIENT_URLS, but unused: shadowed by corresponding flag 
2月 22 12:47:52 localhost.localdomain etcd[11648]: etcd Version: 2.3.7
2月 22 12:47:52 localhost.localdomain etcd[11648]: Git SHA: fd17c91
2月 22 12:47:52 localhost.localdomain etcd[11648]: Go Version: go1.6.3
2月 22 12:47:52 localhost.localdomain etcd[11648]: Go OS/Arch: linux/amd64
2月 22 12:47:52 localhost.localdomain etcd[11648]: setting maximum number of CPUs to 2, total number of available CPUs is 2
2月 22 12:47:52 localhost.localdomain etcd[11648]: listening for peers on http://192.168.2.20:2380
2月 22 12:47:52 localhost.localdomain etcd[11648]: listening for client requests on http://192.168.2.20:2379
2月 22 12:47:52 localhost.localdomain etcd[11648]: listening for client requests on http://localhost:2379
2月 22 12:47:52 localhost.localdomain etcd[11648]: name = default
2月 22 12:47:52 localhost.localdomain etcd[11648]: data dir = /var/lib/etcd/default.etcd
2月 22 12:47:52 localhost.localdomain etcd[11648]: member dir = /var/lib/etcd/default.etcd/member
2月 22 12:47:52 localhost.localdomain etcd[11648]: heartbeat = 100ms
2月 22 12:47:52 localhost.localdomain etcd[11648]: election = 1000ms
2月 22 12:47:52 localhost.localdomain etcd[11648]: snapshot count = 10000
2月 22 12:47:52 localhost.localdomain etcd[11648]: advertise client URLs = http://192.168.2.20:2379
2月 22 12:47:52 localhost.localdomain etcd[11648]: initial advertise peer URLs = http://192.168.2.20:2380
2月 22 12:47:52 localhost.localdomain etcd[11648]: initial cluster = default=http://192.168.2.20:2380
2月 22 12:47:52 localhost.localdomain etcd[11648]: starting member 9bfd26ae5d19811c in cluster aad0dfda8ae36082
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c became follower at term 0
2月 22 12:47:52 localhost.localdomain etcd[11648]: newRaft 9bfd26ae5d19811c [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c became follower at term 1
2月 22 12:47:52 localhost.localdomain etcd[11648]: starting server... [version: 2.3.7, cluster version: to_be_decided]
2月 22 12:47:52 localhost.localdomain systemd[1]: Started Etcd Server.
2月 22 12:47:52 localhost.localdomain etcd[11648]: added local member 9bfd26ae5d19811c [http://192.168.2.20:2380] to cluster aad0dfda8ae36082
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c is starting a new election at term 1
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c became candidate at term 2
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c received vote from 9bfd26ae5d19811c at term 2
2月 22 12:47:52 localhost.localdomain etcd[11648]: 9bfd26ae5d19811c became leader at term 2
2月 22 12:47:52 localhost.localdomain etcd[11648]: raft.node: 9bfd26ae5d19811c elected leader 9bfd26ae5d19811c at term 2
2月 22 12:47:52 localhost.localdomain etcd[11648]: published {Name:default ClientURLs:[http://192.168.2.20:2379]} to cluster aad0dfda8ae36082
2月 22 12:47:52 localhost.localdomain etcd[11648]: setting up the initial cluster version to 2.3
2月 22 12:47:52 localhost.localdomain etcd[11648]: set the initial cluster version to 2.3

```