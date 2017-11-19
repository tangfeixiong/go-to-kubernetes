# 在CentOS 7离线环境安装Kubernetes

## Table of contents

制作CentOS 7本地YUM仓库

制作Kubernetes本地YUM仓库

[CentOS 7.2示例](#CentOS-7.2示例)

* 使用CentOS 7离线仓库安装Docker
* [大小仅12MB的http服务器](#./my-minimal-http-file-server)
* 安装kubeadm

[CentOS升级后](#centos-7-latest)

### Advanced

使用发行包和bash配置请参见[../k8s-v1.7-hand-on](../k8s-v1.7-hand-on)/[../k8s-v1.6-deployment](../k8s-v1.6-deployment)/[../k8s-v1.5-deployment](../k8s-v1.5.7-deployment), 基于 https://github.com/kubernetes/kubernetes/tree/master/cluster/saltbase/salt 

## Local CentOS 7 mirroring repository

脚本参见[../../examples/centos-devops-cync-repo](../../examples/centos-devops-cync-repo)

CentOS 7 exclude isos
```
[vagrant@openshiftdev ~]$ ls /vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7
atomic      cr         isos                                    os                    sclo     virt
centosplus  extras     openshift-ansible-centos-paas-sig.repo  paas                  storage
cloud       fasttrack  openshift-ansible-CentOS-SIG-PaaS       RPM-GPG-KEY-CentOS-7  updates
```

size exclude isos
```
[vagrant@openshiftdev ~]$ du -sh /vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/*
28G	/vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7
```

ISO
```
[vagrant@openshiftdev ~]$ ls /vagrant_f/99-mirror/http0x3A0x2F0x2Fmirror.centos.org/CentOS-7-x86_64-Everything-1611.iso 
/vagrant_f/99-mirror/http0x3A0x2F0x2Fmirror.centos.org/CentOS-7-x86_64-Everything-1611.iso
```

Mount as mirroring repository
```
[vagrant@openshiftdev ~]$ sudo mount -t iso9660 -o loop /vagrant_f/99-mirror/http0x3A0x2F0x2Fmirror.centos.org/CentOS-7-x86_64-Everything-1611.iso /media/cdrom/
mount: /dev/loop0 is write-protected, mounting read-only
```

list isos
```
[vagrant@openshiftdev ~]$ ls /media/cdrom
CentOS_BuildTag  EULA  images    LiveOS    repodata              RPM-GPG-KEY-CentOS-Testing-7
EFI              GPL   isolinux  Packages  RPM-GPG-KEY-CentOS-7  TRANS.TBL
```

## Local Kubernetes mirroring repository

脚本参见[../../http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm](../../http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm)

操作要求能够访问谷歌云平台Google Cloud Platform
```
[vagrant@openshiftdev ~]$ ls /vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum
doc  pool  repos
```

## CentOS 7.2示例

### Power on within vagrant up

From cygwin
```
tangf@DESKTOP-H68OQDV ~
$ ssh vagrant@10.64.33.82
The authenticity of host '10.64.33.82 (10.64.33.82)' can't be established.
ECDSA key fingerprint is SHA256:O3hRMUgb+HetWNkj/43sK6fQ12uQz6TwtJI+wo6VIac.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '10.64.33.82' (ECDSA) to the list of known hosts.
vagrant@10.64.33.82's password: 
Last login: Sun May 28 11:42:38 2017 from 10.0.2.2
```

### CentOS

version
```
[vagrant@openshiftdev ~]$ cat /etc/centos-release
CentOS Linux release 7.2.1511 (Core) 
```

systemd of running
```
[vagrant@openshiftdev ~]$ sudo systemctl list-units --type service --state running --no-pager
UNIT                     LOAD   ACTIVE SUB     DESCRIPTION
atd.service              loaded active running Job spooling tools
auditd.service           loaded active running Security Auditing Service
crond.service            loaded active running Command Scheduler
dbus.service             loaded active running D-Bus System Message Bus
docker.service           loaded active running Docker Application Container Engine
getty@tty1.service       loaded active running Getty on tty1
gssproxy.service         loaded active running GSSAPI Proxy Daemon
lvm2-lvmetad.service     loaded active running LVM2 metadata daemon
NetworkManager.service   loaded active running Network Manager
ntpd.service             loaded active running Network Time Service
polkit.service           loaded active running Authorization Manager
postfix.service          loaded active running Postfix Mail Transport Agent
rsyslog.service          loaded active running System Logging Service
sshd.service             loaded active running OpenSSH server daemon
systemd-journald.service loaded active running Journal Service
systemd-logind.service   loaded active running Login Service
systemd-udevd.service    loaded active running udev Kernel Device Manager
tuned.service            loaded active running Dynamic System Tuning Daemon
vboxadd-service.service  loaded active running LSB: VirtualBox Additions service
wpa_supplicant.service   loaded active running WPA Supplicant daemon

LOAD   = Reflects whether the unit definition was properly loaded.
ACTIVE = The high-level unit activation state, i.e. generalization of SUB.
SUB    = The low-level unit activation state, values depend on unit type.

20 loaded units listed. Pass --all to see loaded but inactive units, too.
To show all installed unit files use 'systemctl list-unit-files'.
```

include exited
```
[vagrant@openshiftdev ~]$ sudo systemctl list-units --type service --state active --no-pager
UNIT                                                   LOAD   ACTIVE SUB     DESCRIPTION
atd.service                                            loaded active running Job spooling tools
auditd.service                                         loaded active running Security Auditing Service
crond.service                                          loaded active running Command Scheduler
dbus.service                                           loaded active running D-Bus System Message Bus
docker.service                                         loaded active running Docker Application Container Engine
getty@tty1.service                                     loaded active running Getty on tty1
gssproxy.service                                       loaded active running GSSAPI Proxy Daemon
iscsi-shutdown.service                                 loaded active exited  Logout off all iSCSI sessions on shutdown
kdump.service                                          loaded active exited  Crash recovery kernel arming
kmod-static-nodes.service                              loaded active exited  Create list of required static device nodes for the cur
lvm2-lvmetad.service                                   loaded active running LVM2 metadata daemon
lvm2-monitor.service                                   loaded active exited  Monitoring of LVM2 mirrors, snapshots etc. using dmeven
lvm2-pvscan@8:2.service                                loaded active exited  LVM2 PV scan on device 8:2
network.service                                        loaded active exited  LSB: Bring up/down networking
NetworkManager.service                                 loaded active running Network Manager
nfs-config.service                                     loaded active exited  Preprocess NFS configuration
ntpd.service                                           loaded active running Network Time Service
polkit.service                                         loaded active running Authorization Manager
postfix.service                                        loaded active running Postfix Mail Transport Agent
rhel-dmesg.service                                     loaded active exited  Dump dmesg to /var/log/dmesg
rhel-import-state.service                              loaded active exited  Import network configuration from initramfs
rhel-readonly.service                                  loaded active exited  Configure read-only root support
rsyslog.service                                        loaded active running System Logging Service
sshd.service                                           loaded active running OpenSSH server daemon
sysstat.service                                        loaded active exited  Resets System Activity Logs
systemd-fsck@dev-centos-openshift\x2dxfs\x2dvol\x2ddir.service loaded active exited  File System Check on /dev/centos/openshift-xfs-vol-dir
systemd-journal-flush.service                          loaded active exited  Flush Journal to Persistent Storage
systemd-journald.service                               loaded active running Journal Service
systemd-logind.service                                 loaded active running Login Service
systemd-random-seed.service                            loaded active exited  Load/Save Random Seed
systemd-remount-fs.service                             loaded active exited  Remount Root and Kernel File Systems
systemd-sysctl.service                                 loaded active exited  Apply Kernel Variables
systemd-tmpfiles-setup-dev.service                     loaded active exited  Create Static Device Nodes in /dev
systemd-tmpfiles-setup.service                         loaded active exited  Create Volatile Files and Directories
systemd-udev-trigger.service                           loaded active exited  udev Coldplug all Devices
systemd-udevd.service                                  loaded active running udev Kernel Device Manager
systemd-update-utmp.service                            loaded active exited  Update UTMP about System Boot/Shutdown
systemd-user-sessions.service                          loaded active exited  Permit User Sessions
systemd-vconsole-setup.service                         loaded active exited  Setup Virtual Console
tuned.service                                          loaded active running Dynamic System Tuning Daemon
vboxadd-service.service                                loaded active running LSB: VirtualBox Additions service
vboxadd-x11.service                                    loaded active exited  LSB: VirtualBox Linux Additions kernel modules
vboxadd.service                                        loaded active exited  LSB: VirtualBox Linux Additions kernel modules
wpa_supplicant.service                                 loaded active running WPA Supplicant daemon

LOAD   = Reflects whether the unit definition was properly loaded.
ACTIVE = The high-level unit activation state, i.e. generalization of SUB.
SUB    = The low-level unit activation state, values depend on unit type.

44 loaded units listed. Pass --all to see loaded but inactive units, too.
To show all installed unit files use 'systemctl list-unit-files'.
```

__make sure openshift being disabled__
```
[vagrant@openshiftdev ~]$ ls /usr/lib/systemd/system/open*
/usr/lib/systemd/system/openshift.service
[vagrant@openshiftdev ~]$ ls /etc/sysconfig/open*
/etc/sysconfig/openshift
[vagrant@openshiftdev ~]$ sudo systemctl is-active openshift.service
unknown
[vagrant@openshiftdev ~]$ sudo systemctl list-unit-files openshift.service
UNIT FILE         STATE   
openshift.service disabled

```

### docker

version
```
[vagrant@openshiftdev ~]$ docker version
Client:
 Version:         1.10.3
 API version:     1.22
 Package version: docker-common-1.10.3-46.el7.centos.14.x86_64
 Go version:      go1.6.3
 Git commit:      cb079f6-unsupported
 Built:           Fri Sep 16 13:24:25 2016
 OS/Arch:         linux/amd64

Server:
 Version:         1.10.3
 API version:     1.22
 Package version: docker-common-1.10.3-46.el7.centos.14.x86_64
 Go version:      go1.6.3
 Git commit:      cb079f6-unsupported
 Built:           Fri Sep 16 13:24:25 2016
 OS/Arch:         linux/amd64
```

info
```
[vagrant@openshiftdev ~]$ docker info
Containers: 0
 Running: 0
 Paused: 0
 Stopped: 0
Images: 7
Server Version: 1.10.3
Storage Driver: devicemapper
 Pool Name: docker-253:0-203725749-pool
 Pool Blocksize: 65.54 kB
 Base Device Size: 10.74 GB
 Backing Filesystem: xfs
 Data file: /dev/centos/docker-data
 Metadata file: /dev/centos/docker-metadata
 Data Space Used: 1.476 GB
 Data Space Total: 30.09 GB
 Data Space Available: 28.62 GB
 Metadata Space Used: 1.524 MB
 Metadata Space Total: 2.189 GB
 Metadata Space Available: 2.188 GB
 Udev Sync Supported: true
 Deferred Removal Enabled: true
 Deferred Deletion Enabled: true
 Deferred Deleted Device Count: 0
 Library Version: 1.02.107-RHEL7 (2016-06-09)
Execution Driver: native-0.2
Logging Driver: journald
Plugins: 
 Volume: local
 Network: bridge null host
Kernel Version: 3.10.0-229.el7.x86_64
Operating System: CentOS Linux 7 (Core)
OSType: linux
Architecture: x86_64
Number of Docker Hooks: 2
CPUs: 1
Total Memory: 3.703 GiB
Name: openshiftdev.local
ID: YJT3:LQHU:2TGI:MEVF:RQEC:D6WA:BVEB:ZBPF:YEYX:6NZX:UIMK:4JOK
WARNING: bridge-nf-call-iptables is disabled
WARNING: bridge-nf-call-ip6tables is disabled
Registries: docker.io (secure)
```

### My minimal HTTP file server

Alternatively, choose nginx, httpd2, tomcat, ...
```
[vagrant@openshiftdev ~]$ docker pull docker.io/tangfeixiong/gofileserver
Using default tag: latest
Trying to pull repository docker.io/tangfeixiong/gofileserver ... 
latest: Pulling from docker.io/tangfeixiong/gofileserver
c0cb142e4345: Pull complete 
bd94e5529d32: Pull complete 
Digest: sha256:b05cb512410b314f3aae3ac1932faec2c8fc892221cc773778f4efdabd3684c8
Status: Downloaded newer image for docker.io/tangfeixiong/gofileserver:latest
```

Size
```
[vagrant@openshiftdev ~]$ docker images tangfeixiong/gofileserver
REPOSITORY                            TAG                 IMAGE ID            CREATED             SIZE
docker.io/tangfeixiong/gofileserver   latest              f07338e49481        5 months ago        12.62 MB
```

如果CentOS上还没有安装Docker, 可以在自己的Docker中先运行gofileserver, 然后docker cp ...
```
[vagrant@openshiftdev ~]$ docker exec -ti gofileserver ls /gofileserver
/gofileserver
[vagrant@openshiftdev ~]$ docker cp gofileserver:/gofileserver .
[vagrant@openshiftdev ~]$ ./gofileserver --help
Usage of ./gofileserver:
  -addr string
    	default listening 0.0.0.0:48080 (default ":48080")
  -dir string
    	default is working directory
```

mirroring
```
[vagrant@openshiftdev ~]$ docker run -d -p 48080:48080 -v /:/mnt --name gofileserver docker.io/tangfeixiong/gofileserver
aa310a7520157c58013d07687c2daa20873e7ac9cea151aa4e877966cbee275c
```

isos
```
[vagrant@openshiftdev ~]$ curl -jkSL http://10.64.33.82:48080/media/cdrom
<pre>
<a href=".discinfo">.discinfo</a>
<a href=".treeinfo">.treeinfo</a>
<a href="CentOS_BuildTag">CentOS_BuildTag</a>
<a href="EFI/">EFI/</a>
<a href="EULA">EULA</a>
<a href="GPL">GPL</a>
<a href="LiveOS/">LiveOS/</a>
<a href="Packages/">Packages/</a>
<a href="RPM-GPG-KEY-CentOS-7">RPM-GPG-KEY-CentOS-7</a>
<a href="RPM-GPG-KEY-CentOS-Testing-7">RPM-GPG-KEY-CentOS-Testing-7</a>
<a href="TRANS.TBL">TRANS.TBL</a>
<a href="images/">images/</a>
<a href="isolinux/">isolinux/</a>
<a href="repodata/">repodata/</a>
</pre>
```

Others
```
[vagrant@openshiftdev ~]$ curl -jkSL http://10.64.33.82:48080/vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/
<pre>
<a href="RPM-GPG-KEY-CentOS-7">RPM-GPG-KEY-CentOS-7</a>
<a href="atomic/">atomic/</a>
<a href="centosplus/">centosplus/</a>
<a href="cloud/">cloud/</a>
<a href="cr/">cr/</a>
<a href="extras/">extras/</a>
<a href="fasttrack/">fasttrack/</a>
<a href="isos/">isos/</a>
<a href="openshift-ansible-CentOS-SIG-PaaS">openshift-ansible-CentOS-SIG-PaaS</a>
<a href="openshift-ansible-centos-paas-sig.repo">openshift-ansible-centos-paas-sig.repo</a>
<a href="os/">os/</a>
<a href="paas/">paas/</a>
<a href="sclo/">sclo/</a>
<a href="storage/">storage/</a>
<a href="updates/">updates/</a>
<a href="virt/">virt/</a>
</pre>
```

Kubernetes
```
[vagrant@openshiftdev ~]$ curl -jkSL http://10.64.33.82:48080/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum
<pre>
<a href="doc/">doc/</a>
<a href="pool/">pool/</a>
<a href="repos/">repos/</a>
</pre>
```

### Install docker

CentOS Local Repository
```
[vagrant@openshiftdev ~]$ sudo cat /etc/yum.repos.d/offline.repo 
# offline.repo
#
# To use this repo, put in your DVD and use it with the other repos too:
#  yum --enablerepo=offline [command]
#  
# or for ONLY the media repo, do this:
#
#  yum --disablerepo=\* --enablerepo=offline\* [command]

[offline]
name=CentOS-$releasever - Offline
baseurl=file:///media/cdrom
gpgcheck=0
enabled=0
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7

[offline-os]
name=CentOS-$releasever - Offline os
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/os/x86_64/
gpgcheck=0
enabled=0
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7

[offline-updates]
name=CentOS-$releasever - Offline updates
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/updates/x86_64/
gpgcheck=0
enabled=0
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7

[offline-extras]
name=CentOS-$releasever - Offline extras
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/extras/x86_64/
gpgcheck=0
enabled=0
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7

[offline-centosplus]
name=CentOS-$releasever - Offline centosplus
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/centosplus/x86_64/
gpgcheck=0
enabled=0
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7

[offline-cr]
name=CentOS-$releasever - Offline cr
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/cr/x86_64/
gpgcheck=0
enabled=0
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7

[offline-fasttrack]
name=CentOS-$releasever - Offline fasttrack
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/fasttrack/x86_64/
gpgcheck=0
enabled=0
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7

[offline-paas]
name=CentOS-$releasever - Offline paas 3.6
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/paas/x86_64/openshift-origin36
gpgcheck=0
enabled=0

[offline-cloud]
name=CentOS-$releasever - Offline cloud openstack-pike
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/cloud/x86_64/openstack-pike
gpgcheck=0
enabled=0

[offline-storage-ceph-jewel]
name=CentOS-$releasever - Offline storage ceph-jewel
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/storage/x86_64/ceph-jewel
gpgcheck=0
enabled=0

[offline-storage-gluster-38]
name=CentOS-$releasever - Offline storage gluster-38
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/storage/x86_64/gluster-3.8
gpgcheck=0
enabled=0

[offline-kvm]
name=CentOS-$releasever - Offline kvm
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/virt/x86_64/kvm-common
gpgcheck=0
enabled=0

[offline-libvirt]
name=CentOS-$releasever - Offline libvirt
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/virt/x86_64/libvirt-latest
gpgcheck=0
enabled=0

[offline-ovirt-41]
name=CentOS-$releasever - Offline ovirt 4.1
baseurl=file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/virt/x86_64/ovirt-4.1
gpgcheck=0
enabled=0

```

For example, search docker
```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=\* --enablerepo=offline*,kubernetes-local search docker
Loaded plugins: fastestmirror
kubernetes-local                                                                                             | 1.4 kB  00:00:00     
offline                                                                                                      | 3.6 kB  00:00:00     
Not using downloaded repomd.xml because it is older than what we have:
  Current   : Fri Jan 13 21:38:01 2017
  Downloaded: Mon Dec  5 13:01:41 2016
offline-centosplus                                                                                           | 3.4 kB  00:00:00     
offline-cloud                                                                                                | 2.9 kB  00:00:00     
offline-cr                                                                                                   | 3.3 kB  00:00:00     
offline-extras                                                                                               | 3.4 kB  00:00:00     
offline-fasttrack                                                                                            | 3.3 kB  00:00:00     
offline-kvm                                                                                                  | 2.9 kB  00:00:00     
offline-libvirt                                                                                              | 2.9 kB  00:00:00     
offline-os                                                                                                   | 3.6 kB  00:00:00     
offline-ovirt-41                                                                                             | 2.9 kB  00:00:00     
offline-paas                                                                                                 | 2.9 kB  00:00:00     
offline-storage-ceph-jewel                                                                                   | 2.9 kB  00:00:00     
offline-storage-gluster-38                                                                                   | 2.9 kB  00:00:00     
offline-updates                                                                                              | 3.4 kB  00:00:00     
(1/14): offline-cr/primary_db                                                                                | 1.2 kB  00:00:00     
(2/14): offline-extras/primary_db                                                                            | 130 kB  00:00:00     
(3/14): offline-fasttrack/primary_db                                                                         | 1.2 kB  00:00:00     
(4/14): offline-kvm/primary_db                                                                               |  26 kB  00:00:00     
(5/14): offline-libvirt/primary_db                                                                           |  59 kB  00:00:00     
(6/14): offline-os/group_gz                                                                                  | 156 kB  00:00:00     
(7/14): offline-ovirt-41/primary_db                                                                          | 125 kB  00:00:00     
(8/14): offline-storage-ceph-jewel/primary_db                                                                |  64 kB  00:00:00     
(9/14): offline-paas/primary_db                                                                              |  30 kB  00:00:00     
(10/14): offline-storage-gluster-38/primary_db                                                               |  23 kB  00:00:00     
(11/14): offline-centosplus/primary_db                                                                       | 1.8 MB  00:00:00     
(12/14): offline-cloud/primary_db                                                                            | 840 kB  00:00:00     
(13/14): offline-updates/primary_db                                                                          | 3.6 MB  00:00:00     
(14/14): offline-os/primary_db                                                                               | 5.7 MB  00:00:00     
Loading mirror speeds from cached hostfile
====================================================== N/S matched: docker ======================================================
cockpit-docker.x86_64 : Cockpit user interface for Docker containers
docker-client.x86_64 : Client side files for Docker
docker-client-latest.x86_64 : Client side files for Docker
docker-common.x86_64 : Common files for docker and docker-latest
docker-distribution.x86_64 : Docker toolset to pack, ship, store, and deliver content
docker-latest-logrotate.x86_64 : cron job to run logrotate on Docker containers
docker-latest-v1.10-migrator.x86_64 : Calculates SHA256 checksums for docker layer content
docker-logrotate.x86_64 : cron job to run logrotate on Docker containers
docker-lvm-plugin.x86_64 : Docker volume driver for lvm volumes
docker-python.x86_64 : An API client for docker written in Python
docker-registry.noarch : Registry server for Docker
docker-registry.x86_64 : Registry server for Docker
docker-selinux.x86_64 : SELinux policies for Docker
docker-v1.10-migrator.x86_64 : Calculates SHA256 checksums for docker layer content
origin-docker-excluder.noarch : Exclude docker packages from updates
origin-dockerregistry.x86_64 : Docker Registry v2 for Origin
pcp-pmda-docker.x86_64 : Performance Co-Pilot (PCP) metrics from the Docker daemon
python-docker-py.noarch : An API client for docker written in Python
python-docker-pycreds.noarch : Python bindings for the docker credentials store API
python-heat-agent-docker-cmd.noarch : Agent for performing Docker based Heat software deployments
python2-docker.noarch : A Python library for the Docker Engine API
python2-docker-pycreds.noarch : Python bindings for the docker credentials store API
docker.x86_64 : Automates deployment of containerized applications
docker-devel.x86_64 : A golang registry for global request variables (source libraries)
docker-forward-journald.x86_64 : Forward stdin to journald
docker-latest.x86_64 : Automates deployment of containerized applications
docker-novolume-plugin.x86_64 : Block container starts with local volumes defined
docker-unit-test.x86_64 : Automates deployment of containerized applications - for running unit tests
oci-systemd-hook.x86_64 : OCI systemd hook for docker
oci-umount.x86_64 : OCI umount hook for docker
openvswitch-ovn-docker.x86_64 : Open vSwitch - Open Virtual Network support
skopeo.x86_64 : Inspect Docker images and repositories on registries

  Name and summary matches only, use "search all" for everything.
```


### Install kubernetes

Local Kubernetes Repository
```
[vagrant@openshiftdev ~]$ sudo cat /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
#baseurl=http://yum.kubernetes.io/repos/kubernetes-el7-x86_64
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
        https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg

[kubernetes-local]
name=Kubernetes offline mirror
baseurl=file:///vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=0
gpgcheck=0

[kubernetes-10-64-33-82]
name=Kubernetes local mirror
baseurl=http://10.64.33.82:48080/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=0
gpgcheck=0

```

Search
```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=\* --enablerepo=offline,kubernetes-local search --verbose kubeadm
Loading "fastestmirror" plugin
Config time: 0.006
Yum version: 3.4.3
Setting up Package Sacks
Loading mirror speeds from cached hostfile
 * offline: 
pkgsack time: 0.003
rpmdb time: 0.000
tags time: 0.000
======================================================= N/S matched: kubeadm =======================================================
kubeadm.x86_64 : Command-line utility for administering a Kubernetes cluster.
Repo        : kubernetes-local




  Name and summary matches only, use "search all" for everything.
```

或者使用本地http
```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=\* --enablerepo=offline\*,kubernetes-10-64-33-82 search kubeadm
Loaded plugins: fastestmirror
kubernetes-10-64-33-82                                                                                       | 1.4 kB  00:00:00     
kubernetes-10-64-33-82/primary                                                                               |  12 kB  00:00:00     
Loading mirror speeds from cached hostfile
kubernetes-10-64-33-82                                                                                                        78/78
======================================================= N/S matched: kubeadm =======================================================
kubeadm.x86_64 : Command-line utility for administering a Kubernetes cluster.

  Name and summary matches only, use "search all" for everything.
```

* Kubeadm

Install
```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=\* --enablerepo=offline,kubernetes-local install -y kubeadm
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * offline: 
Resolving Dependencies
--> Running transaction check
---> Package kubeadm.x86_64 0:1.8.2-0 will be installed
--> Processing Dependency: kubelet >= 1.6.0 for package: kubeadm-1.8.2-0.x86_64
--> Processing Dependency: kubectl >= 1.6.0 for package: kubeadm-1.8.2-0.x86_64
--> Processing Dependency: kubernetes-cni for package: kubeadm-1.8.2-0.x86_64
--> Running transaction check
---> Package kubectl.x86_64 0:1.8.2-0 will be installed
---> Package kubelet.x86_64 0:1.8.2-0 will be installed
---> Package kubernetes-cni.x86_64 0:0.5.1-1 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

====================================================================================================================================
 Package                           Arch                      Version                      Repository                           Size
====================================================================================================================================
Installing:
 kubeadm                           x86_64                    1.8.2-0                      kubernetes-local                     15 M
Installing for dependencies:
 kubectl                           x86_64                    1.8.2-0                      kubernetes-local                    7.3 M
 kubelet                           x86_64                    1.8.2-0                      kubernetes-local                     16 M
 kubernetes-cni                    x86_64                    0.5.1-1                      kubernetes-local                    7.4 M

Transaction Summary
====================================================================================================================================
Install  1 Package (+3 Dependent packages)

Total download size: 45 M
Installed size: 242 M
Downloading packages:
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                               118 MB/s |  45 MB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : kubelet-1.8.2-0.x86_64                                                                                           1/4 
  Installing : kubernetes-cni-0.5.1-1.x86_64                                                                                    2/4 
  Installing : kubectl-1.8.2-0.x86_64                                                                                           3/4 
  Installing : kubeadm-1.8.2-0.x86_64                                                                                           4/4 
  Verifying  : kubeadm-1.8.2-0.x86_64                                                                                           1/4 
  Verifying  : kubectl-1.8.2-0.x86_64                                                                                           2/4 
  Verifying  : kubernetes-cni-0.5.1-1.x86_64                                                                                    3/4 
  Verifying  : kubelet-1.8.2-0.x86_64                                                                                           4/4 

Installed:
  kubeadm.x86_64 0:1.8.2-0                                                                                                          

Dependency Installed:
  kubectl.x86_64 0:1.8.2-0                 kubelet.x86_64 0:1.8.2-0                 kubernetes-cni.x86_64 0:0.5.1-1                

Complete!
```

```
[vagrant@openshiftdev ~]$ which kubeadm
/bin/kubeadm
```

```
[vagrant@openshiftdev ~]$ kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.2", GitCommit:"bdaeafa71f6c7c04636251031f93464384d54963", GitTreeState:"clean", BuildDate:"2017-10-24T19:38:10Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
```




## CentOS 7 Latest

Upgrade docker
```
[vagrant@openshiftdev ~]$ sudo yum upgrade docker
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: centos.ustc.edu.cn
 * epel: mirrors.ustc.edu.cn
 * extras: centos.ustc.edu.cn
 * updates: centos.ustc.edu.cn
Resolving Dependencies
--> Running transaction check
---> Package docker.x86_64 0:1.10.3-46.el7.centos.14 will be updated
---> Package docker.x86_64 2:1.12.6-61.git85d7426.el7.centos will be an update
--> Processing Dependency: docker-common = 2:1.12.6-61.git85d7426.el7.centos for package: 2:docker-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: docker-client = 2:1.12.6-61.git85d7426.el7.centos for package: 2:docker-1.12.6-61.git85d7426.el7.centos.x86_64
--> Running transaction check
---> Package docker-client.x86_64 2:1.12.6-61.git85d7426.el7.centos will be installed
---> Package docker-common.x86_64 0:1.10.3-46.el7.centos.14 will be updated
---> Package docker-common.x86_64 2:1.12.6-61.git85d7426.el7.centos will be an update
--> Processing Dependency: oci-umount >= 2:2.0.0-1 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: oci-systemd-hook >= 1:0.1.4-9 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: oci-register-machine >= 1:0-3.10 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: container-storage-setup >= 0.7.0-1 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: container-selinux >= 2:2.21-2 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: skopeo-containers for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Running transaction check
---> Package container-selinux.noarch 2:2.28-1.git85ce147.el7 will be obsoleting
--> Processing Dependency: policycoreutils >= 2.5-11 for package: 2:container-selinux-2.28-1.git85ce147.el7.noarch
---> Package container-storage-setup.noarch 0:0.7.0-1.git4ca59c5.el7 will be installed
---> Package docker-selinux.x86_64 0:1.10.3-46.el7.centos.14 will be obsoleted
---> Package oci-register-machine.x86_64 1:0-1.8.gitaf6c129.el7 will be updated
---> Package oci-register-machine.x86_64 1:0-3.13.gitcd1e331.el7 will be an update
---> Package oci-systemd-hook.x86_64 1:0.1.4-4.git41491a3.el7 will be updated
---> Package oci-systemd-hook.x86_64 1:0.1.14-1.git1ba44c6.el7 will be an update
---> Package oci-umount.x86_64 2:2.0.0-1.git299e781.el7 will be installed
---> Package skopeo-containers.x86_64 1:0.1.24-1.dev.git28d4e08.el7 will be installed
--> Running transaction check
---> Package policycoreutils.x86_64 0:2.2.5-20.el7 will be updated
--> Processing Dependency: policycoreutils = 2.2.5-20.el7 for package: policycoreutils-python-2.2.5-20.el7.x86_64
---> Package policycoreutils.x86_64 0:2.5-17.1.el7 will be an update
--> Processing Dependency: libsepol >= 2.5-6 for package: policycoreutils-2.5-17.1.el7.x86_64
--> Processing Dependency: libselinux-utils >= 2.5-6 for package: policycoreutils-2.5-17.1.el7.x86_64
--> Processing Dependency: libsepol.so.1(LIBSEPOL_1.1)(64bit) for package: policycoreutils-2.5-17.1.el7.x86_64
--> Processing Dependency: libsepol.so.1(LIBSEPOL_1.0)(64bit) for package: policycoreutils-2.5-17.1.el7.x86_64
--> Processing Dependency: libsemanage.so.1(LIBSEMANAGE_1.1)(64bit) for package: policycoreutils-2.5-17.1.el7.x86_64
--> Running transaction check
---> Package libselinux-utils.x86_64 0:2.2.2-6.el7 will be updated
---> Package libselinux-utils.x86_64 0:2.5-11.el7 will be an update
--> Processing Dependency: libselinux(x86-64) = 2.5-11.el7 for package: libselinux-utils-2.5-11.el7.x86_64
---> Package libsemanage.x86_64 0:2.1.10-18.el7 will be updated
--> Processing Dependency: libsemanage = 2.1.10-18.el7 for package: libsemanage-python-2.1.10-18.el7.x86_64
---> Package libsemanage.x86_64 0:2.5-8.el7 will be an update
---> Package libsepol.x86_64 0:2.1.9-3.el7 will be updated
--> Processing Dependency: libsepol = 2.1.9-3.el7 for package: libsepol-devel-2.1.9-3.el7.x86_64
---> Package libsepol.x86_64 0:2.5-6.el7 will be an update
---> Package policycoreutils-python.x86_64 0:2.2.5-20.el7 will be updated
---> Package policycoreutils-python.x86_64 0:2.5-17.1.el7 will be an update
--> Processing Dependency: setools-libs >= 3.3.8-1 for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Running transaction check
---> Package libselinux.x86_64 0:2.2.2-6.el7 will be updated
--> Processing Dependency: libselinux = 2.2.2-6.el7 for package: libselinux-python-2.2.2-6.el7.x86_64
--> Processing Dependency: libselinux = 2.2.2-6.el7 for package: libselinux-devel-2.2.2-6.el7.x86_64
---> Package libselinux.x86_64 0:2.5-11.el7 will be an update
---> Package libsemanage-python.x86_64 0:2.1.10-18.el7 will be updated
---> Package libsemanage-python.x86_64 0:2.5-8.el7 will be an update
---> Package libsepol-devel.x86_64 0:2.1.9-3.el7 will be updated
---> Package libsepol-devel.x86_64 0:2.5-6.el7 will be an update
---> Package setools-libs.x86_64 0:3.3.7-46.el7 will be updated
---> Package setools-libs.x86_64 0:3.3.8-1.1.el7 will be an update
--> Running transaction check
---> Package libselinux-devel.x86_64 0:2.2.2-6.el7 will be updated
---> Package libselinux-devel.x86_64 0:2.5-11.el7 will be an update
---> Package libselinux-python.x86_64 0:2.2.2-6.el7 will be updated
---> Package libselinux-python.x86_64 0:2.5-11.el7 will be an update
--> Processing Conflict: libselinux-2.5-11.el7.x86_64 conflicts systemd < 219-20
--> Restarting Dependency Resolution with new changes.
--> Running transaction check
---> Package systemd.x86_64 0:219-19.el7_2.13 will be updated
--> Processing Dependency: systemd = 219-19.el7_2.13 for package: systemd-sysv-219-19.el7_2.13.x86_64
--> Processing Dependency: systemd = 219-19.el7_2.13 for package: systemd-devel-219-19.el7_2.13.x86_64
---> Package systemd.x86_64 0:219-42.el7_4.4 will be an update
--> Processing Dependency: systemd-libs = 219-42.el7_4.4 for package: systemd-219-42.el7_4.4.x86_64
--> Running transaction check
---> Package systemd-devel.x86_64 0:219-19.el7_2.13 will be updated
---> Package systemd-devel.x86_64 0:219-42.el7_4.4 will be an update
---> Package systemd-libs.x86_64 0:219-19.el7_2.13 will be updated
--> Processing Dependency: systemd-libs = 219-19.el7_2.13 for package: libgudev1-219-19.el7_2.13.x86_64
---> Package systemd-libs.x86_64 0:219-42.el7_4.4 will be an update
---> Package systemd-sysv.x86_64 0:219-19.el7_2.13 will be updated
---> Package systemd-sysv.x86_64 0:219-42.el7_4.4 will be an update
--> Running transaction check
---> Package libgudev1.x86_64 0:219-19.el7_2.13 will be updated
---> Package libgudev1.x86_64 0:219-42.el7_4.4 will be an update
--> Processing Conflict: systemd-219-42.el7_4.4.x86_64 conflicts dracut < 033-499
--> Restarting Dependency Resolution with new changes.
--> Running transaction check
---> Package dracut.x86_64 0:033-360.el7_2.1 will be updated
--> Processing Dependency: dracut = 033-360.el7_2.1 for package: dracut-config-rescue-033-360.el7_2.1.x86_64
--> Processing Dependency: dracut = 033-360.el7_2.1 for package: dracut-network-033-360.el7_2.1.x86_64
---> Package dracut.x86_64 0:033-502.el7 will be an update
--> Running transaction check
---> Package dracut-config-rescue.x86_64 0:033-360.el7_2.1 will be updated
---> Package dracut-config-rescue.x86_64 0:033-502.el7 will be an update
---> Package dracut-network.x86_64 0:033-360.el7_2.1 will be updated
---> Package dracut-network.x86_64 0:033-502.el7 will be an update
--> Processing Conflict: libsemanage-2.5-8.el7.x86_64 conflicts selinux-policy-base < 3.13.1-66
--> Restarting Dependency Resolution with new changes.
--> Running transaction check
---> Package selinux-policy-targeted.noarch 0:3.13.1-60.el7_2.9 will be updated
---> Package selinux-policy-targeted.noarch 0:3.13.1-166.el7_4.5 will be an update
--> Processing Dependency: selinux-policy = 3.13.1-166.el7_4.5 for package: selinux-policy-targeted-3.13.1-166.el7_4.5.noarch
--> Processing Dependency: selinux-policy = 3.13.1-166.el7_4.5 for package: selinux-policy-targeted-3.13.1-166.el7_4.5.noarch
--> Running transaction check
---> Package selinux-policy.noarch 0:3.13.1-60.el7_2.9 will be updated
---> Package selinux-policy.noarch 0:3.13.1-166.el7_4.5 will be an update
--> Finished Dependency Resolution

Dependencies Resolved

====================================================================================================================================
 Package                              Arch                Version                                        Repository            Size
====================================================================================================================================
Installing:
 container-selinux                    noarch              2:2.28-1.git85ce147.el7                        extras                29 k
     replacing  docker-selinux.x86_64 1.10.3-46.el7.centos.14
Updating:
 docker                               x86_64              2:1.12.6-61.git85d7426.el7.centos              extras                15 M
 dracut                               x86_64              033-502.el7                                    base                 321 k
 selinux-policy-targeted              noarch              3.13.1-166.el7_4.5                             updates              6.5 M
 systemd                              x86_64              219-42.el7_4.4                                 updates              5.2 M
Installing for dependencies:
 container-storage-setup              noarch              0.7.0-1.git4ca59c5.el7                         extras                32 k
 docker-client                        x86_64              2:1.12.6-61.git85d7426.el7.centos              extras               3.4 M
 oci-umount                           x86_64              2:2.0.0-1.git299e781.el7                       extras                27 k
 skopeo-containers                    x86_64              1:0.1.24-1.dev.git28d4e08.el7                  extras               8.5 k
Updating for dependencies:
 docker-common                        x86_64              2:1.12.6-61.git85d7426.el7.centos              extras                80 k
 dracut-config-rescue                 x86_64              033-502.el7                                    base                  55 k
 dracut-network                       x86_64              033-502.el7                                    base                  97 k
 libgudev1                            x86_64              219-42.el7_4.4                                 updates               83 k
 libselinux                           x86_64              2.5-11.el7                                     base                 162 k
 libselinux-devel                     x86_64              2.5-11.el7                                     base                 186 k
 libselinux-python                    x86_64              2.5-11.el7                                     base                 234 k
 libselinux-utils                     x86_64              2.5-11.el7                                     base                 151 k
 libsemanage                          x86_64              2.5-8.el7                                      base                 145 k
 libsemanage-python                   x86_64              2.5-8.el7                                      base                 104 k
 libsepol                             x86_64              2.5-6.el7                                      base                 288 k
 libsepol-devel                       x86_64              2.5-6.el7                                      base                  74 k
 oci-register-machine                 x86_64              1:0-3.13.gitcd1e331.el7                        extras               1.1 M
 oci-systemd-hook                     x86_64              1:0.1.14-1.git1ba44c6.el7                      extras                32 k
 policycoreutils                      x86_64              2.5-17.1.el7                                   base                 858 k
 policycoreutils-python               x86_64              2.5-17.1.el7                                   base                 446 k
 selinux-policy                       noarch              3.13.1-166.el7_4.5                             updates              436 k
 setools-libs                         x86_64              3.3.8-1.1.el7                                  base                 612 k
 systemd-devel                        x86_64              219-42.el7_4.4                                 updates              186 k
 systemd-libs                         x86_64              219-42.el7_4.4                                 updates              376 k
 systemd-sysv                         x86_64              219-42.el7_4.4                                 updates               70 k

Transaction Summary
====================================================================================================================================
Install  1 Package  (+ 4 Dependent packages)
Upgrade  4 Packages (+21 Dependent packages)

Total download size: 36 M
Is this ok [y/d/N]: 
```

### Kubernetes nightly