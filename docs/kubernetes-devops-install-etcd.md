Install ETCD 安装ETCD配置管理服务
===============================

POC
----

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

* Install from local repo

Using *gofileserver*

    tangf@DESKTOP-H68OQDV /cygdrive/g/2015-12-19-repository/99-mirror/centos/rsync%3A%2F%2Fmirrors.yun-idc.com%2Fcentos%2F7
    $ cd /cygdrive/i/mirror/centos/rsync%3A%2F%2Fmirrors.yun-idc.com%2Fcentos%2F7/

    tangf@DESKTOP-H68OQDV /cygdrive/i/mirror/centos/rsync%3A%2F%2Fmirrors.yun-idc.com%2Fcentos%2F7
    $ gofileserver.exe
    Listening at  :48080

    [vagrant@localhost ~]$ sudo yum --disablerepo=extras --enablerepo=extras-mirror install -y etcd
    Loaded plugins: fastestmirror
    Loading mirror speeds from cached hostfile
     * base: mirrors.zju.edu.cn
     * updates: mirrors.zju.edu.cn
    Resolving Dependencies
    --> Running transaction check
    ---> Package etcd.x86_64 0:2.3.7-4.el7 will be installed
    --> Finished Dependency Resolution

    Dependencies Resolved

    ================================================================================
     Package      Arch           Version                Repository             Size
    ================================================================================
    Installing:
     etcd         x86_64         2.3.7-4.el7            extras-mirror         6.5 M

    Transaction Summary
    ================================================================================
    Install  1 Package

    Total download size: 6.5 M
    Installed size: 31 M
    Downloading packages:
    etcd-2.3.7-4.el7.x86_64.rpm                                | 6.5 MB   00:06
    Running transaction check
    Running transaction test
    Transaction test succeeded
    Running transaction
      Installing : etcd-2.3.7-4.el7.x86_64                                      1/1
      Verifying  : etcd-2.3.7-4.el7.x86_64                                      1/1

    Installed:
      etcd.x86_64 0:2.3.7-4.el7

    Complete!

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

    [vagrant@localhost ~]$ sudo vi /etc/etcd/etcd.conf
    # [member]
    ETCD_NAME=default
    ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
    #ETCD_WAL_DIR=""
    #ETCD_SNAPSHOT_COUNT="10000"
    #ETCD_HEARTBEAT_INTERVAL="100"
    #ETCD_ELECTION_TIMEOUT="1000"
    #ETCD_LISTEN_PEER_URLS="http://localhost:2380"
    #ETCD_LISTEN_CLIENT_URLS="http://localhost:2379"
    ETCD_LISTEN_CLIENT_URLS="http://10.64.33.81:2379"
    #ETCD_MAX_SNAPSHOTS="5"
    #ETCD_MAX_WALS="5"
    #ETCD_CORS=""
    #
    #[cluster]
    #ETCD_INITIAL_ADVERTISE_PEER_URLS="http://localhost:2380"
    # if you use different ETCD_NAME (e.g. test), set ETCD_INITIAL_CLUSTER value for this name, i.e. "test=http://..."
    #ETCD_INITIAL_CLUSTER="default=http://localhost:2380"
    #ETCD_INITIAL_CLUSTER_STATE="new"
    #ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
    #ETCD_ADVERTISE_CLIENT_URLS="http://localhost:2379"
    ETCD_ADVERTISE_CLIENT_URLS="http://0.0.0.0:2379"
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

    [vagrant@localhost ~]$ sudo systemctl enable etcd.service
    Created symlink from /etc/systemd/system/multi-user.target.wants/etcd.service to /usr/lib/systemd/system/etcd.service.

    [vagrant@localhost ~]$ ls -l /etc/systemd/system/multi-user.target.wants/etcd.service
    lrwxrwxrwx. 1 root root 36 Nov 10 20:20 /etc/systemd/system/multi-user.target.wants/etcd.service -> /usr/lib/systemd/system/etcd.service

    [vagrant@localhost ~]$ sudo systemctl start etcd.service

    [vagrant@localhost ~]$ sudo systemctl -l status etcd.service
    鈼▒ etcd.service - Etcd Server
       Loaded: loaded (/usr/lib/systemd/system/etcd.service; enabled; vendor preset: disabled)
       Active: active (running) since Thu 2016-11-10 20:25:15 UTC; 52s ago
     Main PID: 4473 (etcd)
       CGroup: /system.slice/etcd.service
               鈹斺攢4473 /usr/bin/etcd --name=default --data-dir=/var/lib/etcd/default.etcd --listen-client-urls=http://0.0.0.0:2379

    Nov 10 20:25:15 localhost.localdomain etcd[4473]: starting server... [version: 2.3.7, cluster version: to_be_decided]
    Nov 10 20:25:15 localhost.localdomain systemd[1]: Started Etcd Server.
    Nov 10 20:25:15 localhost.localdomain etcd[4473]: added local member ce2a822cea30bfca [http://localhost:2380 http://localhost:7001] to cluster 7e27652122e8b2ae
    Nov 10 20:25:15 localhost.localdomain etcd[4473]: set the initial cluster version to 2.3
    Nov 10 20:25:17 localhost.localdomain etcd[4473]: ce2a822cea30bfca is starting a new election at term 2
    Nov 10 20:25:17 localhost.localdomain etcd[4473]: ce2a822cea30bfca became candidate at term 3
    Nov 10 20:25:17 localhost.localdomain etcd[4473]: ce2a822cea30bfca received vote from ce2a822cea30bfca at term 3
    Nov 10 20:25:17 localhost.localdomain etcd[4473]: ce2a822cea30bfca became leader at term 3
    Nov 10 20:25:17 localhost.localdomain etcd[4473]: raft.node: ce2a822cea30bfca elected leader ce2a822cea30bfca at term 3
    Nov 10 20:25:17 localhost.localdomain etcd[4473]: published {Name:default ClientURLs:[http://10.64.33.81:2379]} to cluster 7e27652122e8b2ae

    [vagrant@localhost ~]$ sudo journalctl --no-pager --no-tail -u etcd.service

    [vagrant@localhost ~]$ curl http://127.0.0.1:2379/version
    {"etcdserver":"2.3.7","etcdcluster":"2.3.0"}

    [vagrant@localhost ~]$ which etcdctl
    /usr/bin/etcdctl

    [vagrant@localhost ~]$ etcdctl member list
    ce2a822cea30bfca: name=default peerURLs=http://localhost:2380,http://localhost:7001 clientURLs=http://10.64.33.81:2379 isLeader=true

    [vagrant@localhost ~]$ etcdctl cluster-health
    member ce2a822cea30bfca is healthy: got healthy result from http://10.64.33.81:2379
    cluster is healthy
