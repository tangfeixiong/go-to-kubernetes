Mount CentOS Complete ISO Media as YUM Repositories - 配置CentOS光盘到YUM安装仓库
===============================================================================

Tables of Content
-------------------

* CentOS Everything ISO

* Search sample software packages into meida repo

* Install software from media repo

Using Everything ISO as OS-Media Repo - 下载CentOS Everything ISO作为离线包仓库
-----------------------------------------------------------------------------

Execute __Linux__ command _wget_ to download - 从[阿里云镜像站](http://mirrors.aliyun.com)下载

```
tangf@DESKTOP-H68OQDV /cygdrive/f/99-mirror/http%3A%2F%2Fmirror.centos.org/centos%2F7%2E3%2E1611%2Fisos%2Fx86_64
$ wget -c http://mirrors.aliyun.com/centos/7.3.1611/isos/x86_64/CentOS-7-x86_64-Everything-1611.iso
--2017-02-19 03:51:52--  http://mirrors.aliyun.com/centos/7.3.1611/isos/x86_64/CentOS-7-x86_64-Everything-1611.iso
正在解析主机 mirrors.aliyun.com (mirrors.aliyun.com)... 115.28.122.210, 112.124.140.210
正在连接 mirrors.aliyun.com (mirrors.aliyun.com)|115.28.122.210|:80... 已连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：8280604672 (7.7G) [application/octet-stream]
正在保存至: “CentOS-7-x86_64-Everything-1611.iso”

 CentOS-7-x86_64-Ev  10%[=>                  ] 807.68M  1.02MB/s    eta 1h 58m

```

The _yum_ repo for CentOS 7 Everything ISO

```
[vagrant@openshiftdev ~]$ cat /etc/yum.repos.d/CentOS-Media.repo
# CentOS-Media.repo
#
#  This repo can be used with mounted DVD media, verify the mount point for
#  CentOS-7.  You can use this repo and yum to install items directly off the
#  DVD ISO that we release.
#
# To use this repo, put in your DVD and use it with the other repos too:
#  yum --enablerepo=c7-media [command]
#
# or for ONLY the media repo, do this:
#
#  yum --disablerepo=\* --enablerepo=c7-media [command]

[c7-media]
name=CentOS-$releasever - Media
baseurl=file:///media/CentOS/
        file:///media/cdrom/
        file:///media/cdrecorder/
gpgcheck=1
enabled=0
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7

```

Then command _mount_ ISO into file system

```
[vagrant@openshiftdev ~]$ sudo mkdir -p /media/cdrom
[vagrant@openshiftdev ~]$ sudo mount -t iso9660 -o loop /vagrant_data/99-mirror/http%3A%2F%2Fmirror.centos.org/centos%2F7%2Fisos%2Fx86_64/CentOS-7-x86_64-Everything-1611.iso /media/cdrom/
mount: /dev/loop0 is write-protected, mounting read-only
[vagrant@openshiftdev ~]$ ls /media/cdrom/
CentOS_BuildTag  GPL       LiveOS    RPM-GPG-KEY-CentOS-7
EFI              images    Packages  RPM-GPG-KEY-CentOS-Testing-7
EULA             isolinux  repodata  TRANS.TBL
[vagrant@openshiftdev ~]$ ls -1 /media/cdrom/Packages/ | wc -l
9364

```

And list and search from __media__ repo, for instance: __nginx__ and __mysql__

```
[vagrant@openshiftdev ~]$ yum --disablerepo=\* --enablerepo=c7-media search nginx
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * c7-media:
============================== N/S matched: nginx ==============================
pcp-pmda-nginx.x86_64 : Performance Co-Pilot (PCP) metrics for the Nginx
                      : Webserver

  Name and summary matches only, use "search all" for everything.
[vagrant@openshiftdev ~]$ yum --disablerepo=\* --enablerepo=c7-media list | grep nginx
pcp-pmda-nginx.x86_64                     3.11.3-4.el7                 c7-media

[vagrant@openshiftdev ~]$ yum --disablerepo=\* --enablerepo=c7-media list | grep mysql
akonadi-mysql.x86_64                      1.9.2-4.el7                  c7-media
apr-util-mysql.x86_64                     1.5.2-6.el7                  c7-media
dovecot-mysql.x86_64                      1:2.2.10-7.el7               c7-media
freeradius-mysql.x86_64                   3.0.4-6.el7                  c7-media
libdbi-dbd-mysql.x86_64                   0.8.3-16.el7                 c7-media
mysql-connector-java.noarch               1:5.1.25-3.el7               c7-media
mysql-connector-odbc.x86_64               5.2.5-6.el7                  c7-media
pcp-pmda-mysql.x86_64                     3.11.3-4.el7                 c7-media
php-mysql.x86_64                          5.4.16-42.el7                c7-media
php-mysqlnd.x86_64                        5.4.16-42.el7                c7-media
qt-mysql.i686                             1:4.8.5-13.el7               c7-media
qt-mysql.x86_64                           1:4.8.5-13.el7               c7-media
qt5-qtbase-mysql.i686                     5.6.1-10.el7                 c7-media
qt5-qtbase-mysql.x86_64                   5.6.1-10.el7                 c7-media
redland-mysql.x86_64                      1.0.16-6.el7                 c7-media
rsyslog-mysql.x86_64                      7.4.7-16.el7                 c7-media

```

For example, install **net_tools** and __wget__

```
[vagrant@localhost ~]$ sudo yum --disablerepo=\* --enablerepo=c7-media install net-tools
Loaded plugins: fastestmirror
c7-media                                                 | 3.6 kB     00:00
Not using downloaded repomd.xml because it is older than what we have:
  Current   : Wed Dec  9 23:14:09 2015
  Downloaded: Wed Dec  9 22:35:45 2015
Loading mirror speeds from cached hostfile
 * c7-media:
Resolving Dependencies
--> Running transaction check
---> Package net-tools.x86_64 0:2.0-0.17.20131004git.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package        Arch        Version                         Repository     Size
================================================================================
Installing:
 net-tools      x86_64      2.0-0.17.20131004git.el7        c7-media      304 k

Transaction Summary
================================================================================
Install  1 Package

[vagrant@localhost ~]$ sudo netstat -tpnl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 127.0.0.1:2380          0.0.0.0:*               LISTEN      998/etcd
tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      1225/kube-apiserver
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      993/sshd
tcp        0      0 127.0.0.1:25            0.0.0.0:*               LISTEN      1430/master
tcp        0      0 127.0.0.1:7001          0.0.0.0:*               LISTEN      998/etcd
tcp6       0      0 :::6443                 :::*                    LISTEN      1225/kube-apiserver
tcp6       0      0 :::2379                 :::*                    LISTEN      998/etcd
tcp6       0      0 :::22                   :::*                    LISTEN      993/sshd
tcp6       0      0 ::1:25                  :::*                    LISTEN      1430/master


[vagrant@localhost ~]$ sudo yum --disablerepo=\* --enablerepo=c7-media install wget
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * c7-media:
Resolving Dependencies
--> Running transaction check
---> Package wget.x86_64 0:1.14-10.el7_0.1 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package       Arch            Version                  Repository         Size
================================================================================
Installing:
 wget          x86_64          1.14-10.el7_0.1          c7-media          545 k

Transaction Summary
================================================================================
Install  1 Package

[vagrant@localhost ~]$ wget
wget: missing URL
Usage: wget [OPTION]... [URL]...

Try `wget --help' for more options.

```

Optional, install __nginx__ as file server

```


```

