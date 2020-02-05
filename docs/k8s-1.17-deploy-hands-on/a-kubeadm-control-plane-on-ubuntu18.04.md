# A kubernetes control-plane on Ubuntu 18.04 Instance

Reference

+ https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/#initializing-your-control-plane-node
+ https://www.linode.com/docs/kubernetes/getting-started-with-kubernetes/
+ https://computingforgeeks.com/manually-pull-container-images-used-by-kubernetes-kubeadm/
+ https://github.com/kubernetes/minikube/issues/4224#issuecomment-490529544


## Docker and Apt resources

Launch instance by OpenStack

The dest addr is OpenStack floating IP, sometimes may failed to login
```
root@wei-ThinkPad-X1-Extreme:~# ssh -i keypair-train.pem ubuntu@192.168.192.123
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
IT IS POSSIBLE THAT SOMEONE IS DOING SOMETHING NASTY!
Someone could be eavesdropping on you right now (man-in-the-middle attack)!
It is also possible that a host key has just been changed.
The fingerprint for the ECDSA key sent by the remote host is
SHA256:OGsYUn9fnp955KMYc1Zs3Ze55Z2HwEBKT9/mNJJuI4o.
Please contact your system administrator.
Add correct host key in /root/.ssh/known_hosts to get rid of this message.
Offending ECDSA key in /root/.ssh/known_hosts:3
  remove with:
  ssh-keygen -f "/root/.ssh/known_hosts" -R "192.168.192.123"
ECDSA host key for 192.168.192.123 has changed and you have requested strict checking.
Host key verification failed.
```

```
root@wei-ThinkPad-X1-Extreme:~# ssh-keygen -f .ssh/known_hosts -R 192.168.192.123
# Host 192.168.192.123 found: line 3
.ssh/known_hosts updated.
Original contents retained as .ssh/known_hosts.old
```

Normal login
```
root@wei-ThinkPad-X1-Extreme:~# ssh -i keypair-train.pem ubuntu@192.168.192.123
The authenticity of host '192.168.192.123 (192.168.192.123)' can't be established.
ECDSA key fingerprint is SHA256:OGsYUn9fnp955KMYc1Zs3Ze55Z2HwEBKT9/mNJJuI4o.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '192.168.192.123' (ECDSA) to the list of known hosts.
Welcome to Ubuntu 18.04.3 LTS (GNU/Linux 4.15.0-76-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  System information as of Sun Feb  2 21:03:18 UTC 2020

  System load:  0.01              Processes:           83
  Usage of /:   6.7% of 14.37GB   Users logged in:     0
  Memory usage: 5%                IP address for ens3: 10.20.30.56
  Swap usage:   0%

0 packages can be updated.
0 updates are security updates.


The programs included with the Ubuntu system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Ubuntu comes with ABSOLUTELY NO WARRANTY, to the extent permitted by
applicable law.

To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.
```

```
ubuntu@masterk8s:~$ uname -a
Linux masterk8s 4.15.0-76-generic #86-Ubuntu SMP Fri Jan 17 17:24:28 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
```

```
ubuntu@masterk8s:~$ cat /etc/os-release 
NAME="Ubuntu"
VERSION="18.04.3 LTS (Bionic Beaver)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 18.04.3 LTS"
VERSION_ID="18.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=bionic
UBUNTU_CODENAME=bionic
```


### Trouble of nameservers

The OpenStack tenant network always config DHCP
```
wei@wei-ThinkPad-X1-Extreme:~$ ip netns
qrouter-f46292f9-3e0d-427e-83dc-13d803537360 (id: 2)
qdhcp-13d1dd1d-71d5-4b2e-8876-45158bdfe7f8 (id: 0)
qdhcp-cb0c8be8-b36b-4bb6-935d-4024d44c38bc (id: 1)
```

It worked as local nameserver, but has no way to know Internet DNS
```
wei@wei-ThinkPad-X1-Extreme:~$ sudo ip netns exec qdhcp-cb0c8be8-b36b-4bb6-935d-4024d44c38bc ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: ns-acc343dc-8c@if17: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default qlen 1000
    link/ether fa:16:3e:17:41:93 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 10.20.30.2/24 brd 10.20.30.255 scope global ns-acc343dc-8c
       valid_lft forever preferred_lft forever
    inet 169.254.169.254/16 brd 169.254.255.255 scope global ns-acc343dc-8c
       valid_lft forever preferred_lft forever
    inet6 fe80::f816:3eff:fe17:4193/64 scope link 
       valid_lft forever preferred_lft forever 
```

One by one to setup nameservers in this instance

Note, the correct way is modify OpenStack neutron-dhcp-agent to add related config options, link to https://github.com/tangfeixiong/go-to-openstack-bootcamp/blob/master/docs/openstack-train-devops_continued.md#issue

后来，通过检查Kubernetes的运行日志，这种单机配置方法是不可取的。 因为kubernetes上也有coredns域名服务，而规范上最多只能有5个域名服务地址。
coredns从该master主机上无法继承全部4个服务地址，因为其内置了2个（1个用于内部cluster service，一个也是8.8.8.8），最终丢弃了最前的2个。

In new systemd system, the resolve file could not modified
```
ubuntu@masterk8s:~$ ls -l /etc/resolv.conf 
lrwxrwxrwx 1 root root 39 Jan 29 21:15 /etc/resolv.conf -> ../run/systemd/resolve/stub-resolv.conf
```

Alternatively new _netplan_ command is prompt
```
ubuntu@masterk8s:~$ sudo vi /etc/netplan/50-cloud-init.yaml
```

Now add proper nameservers like
```
ubuntu@masterk8s:~$ cat /etc/netplan/50-cloud-init.yaml 
# This file is generated from information provided by the datasource.  Changes
# to it will not persist across an instance reboot.  To disable cloud-init's
# network configuration capabilities, write a file
# /etc/cloud/cloud.cfg.d/99-disable-network-config.cfg with the following:
# network: {config: disabled}
network:
    ethernets:
        ens3:
            dhcp4: true
            match:
                macaddress: fa:16:3e:ae:bc:0b
            set-name: ens3
            nameservers:
                addresses: [114.114.114.114, 8.8.4.4, 8.8.8.8]
    version: 2
```

Apply mod
```
ubuntu@masterk8s:~$ sudo netplan apply
```

Now 4 nameservers were used
```
ubuntu@masterk8s:~$ sudo systemd-resolve --status
Global
          DNSSEC NTA: 10.in-addr.arpa
                      16.172.in-addr.arpa
                      168.192.in-addr.arpa
                      17.172.in-addr.arpa
                      18.172.in-addr.arpa
                      19.172.in-addr.arpa
                      20.172.in-addr.arpa
                      21.172.in-addr.arpa
                      22.172.in-addr.arpa
                      23.172.in-addr.arpa
                      24.172.in-addr.arpa
                      25.172.in-addr.arpa
                      26.172.in-addr.arpa
                      27.172.in-addr.arpa
                      28.172.in-addr.arpa
                      29.172.in-addr.arpa
                      30.172.in-addr.arpa
                      31.172.in-addr.arpa
                      corp
                      d.f.ip6.arpa
                      home
                      internal
                      intranet
                      lan
                      local
                      private
                      test

Link 2 (ens3)
      Current Scopes: DNS
       LLMNR setting: yes
MulticastDNS setting: no
      DNSSEC setting: no
    DNSSEC supported: no
         DNS Servers: 114.114.114.114
                      8.8.4.4
                      8.8.8.8
                      10.20.30.2
          DNS Domain: openstacklocal
```


## Ubuntu local archive repository

The image is from https://cloud-images.ubuntu.com/releases/

The archive location is default
```
ubuntu@masterk8s:~$ cat /etc/apt/sources.list
## Note, this file is written by cloud-init on first boot of an instance
## modifications made here will not survive a re-bundle.
## if you wish to make changes you can:
## a.) add 'apt_preserve_sources_list: true' to /etc/cloud/cloud.cfg
##     or do the same in user-data
## b.) add sources in /etc/apt/sources.list.d
## c.) make changes to template file /etc/cloud/templates/sources.list.tmpl

# See http://help.ubuntu.com/community/UpgradeNotes for how to upgrade to
# newer versions of the distribution.
deb http://archive.ubuntu.com/ubuntu bionic main restricted
# deb-src http://archive.ubuntu.com/ubuntu bionic main restricted

## Major bug fix updates produced after the final release of the
## distribution.
deb http://archive.ubuntu.com/ubuntu bionic-updates main restricted
# deb-src http://archive.ubuntu.com/ubuntu bionic-updates main restricted

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team. Also, please note that software in universe WILL NOT receive any
## review or updates from the Ubuntu security team.
deb http://archive.ubuntu.com/ubuntu bionic universe
# deb-src http://archive.ubuntu.com/ubuntu bionic universe
deb http://archive.ubuntu.com/ubuntu bionic-updates universe
# deb-src http://archive.ubuntu.com/ubuntu bionic-updates universe

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team, and may not be under a free licence. Please satisfy yourself as to
## your rights to use the software. Also, please note that software in
## multiverse WILL NOT receive any review or updates from the Ubuntu
## security team.
deb http://archive.ubuntu.com/ubuntu bionic multiverse
# deb-src http://archive.ubuntu.com/ubuntu bionic multiverse
deb http://archive.ubuntu.com/ubuntu bionic-updates multiverse
# deb-src http://archive.ubuntu.com/ubuntu bionic-updates multiverse

## N.B. software from this repository may not have been tested as
## extensively as that contained in the main release, although it includes
## newer versions of some applications which may provide useful features.
## Also, please note that software in backports WILL NOT receive any review
## or updates from the Ubuntu security team.
deb http://archive.ubuntu.com/ubuntu bionic-backports main restricted universe multiverse
# deb-src http://archive.ubuntu.com/ubuntu bionic-backports main restricted universe multiverse

## Uncomment the following two lines to add software from Canonical's
## 'partner' repository.
## This software is not part of Ubuntu, but is offered by Canonical and the
## respective vendors as a service to Ubuntu users.
# deb http://archive.canonical.com/ubuntu bionic partner
# deb-src http://archive.canonical.com/ubuntu bionic partner

deb http://security.ubuntu.com/ubuntu bionic-security main restricted
# deb-src http://security.ubuntu.com/ubuntu bionic-security main restricted
deb http://security.ubuntu.com/ubuntu bionic-security universe
# deb-src http://security.ubuntu.com/ubuntu bionic-security universe
deb http://security.ubuntu.com/ubuntu bionic-security multiverse
# deb-src http://security.ubuntu.com/ubuntu bionic-security multiverse
```

The physical machine is same Ubuntu version, Fetched ISO format from https://mirrors.tuna.tsinghua.edu.cn/ubuntu-releases/

And its archive is local
```
ubuntu@masterk8s:~$ curl -jkSL http://192.168.0.106:48081/etc0x2Fapt0x2Fsources.list
# deb cdrom:[Ubuntu 18.04.3 LTS _Bionic Beaver_ - Release amd64 (20190805)]/ bionic main restricted

# See http://help.ubuntu.com/community/UpgradeNotes for how to upgrade to
# newer versions of the distribution.
deb http://cn.archive.ubuntu.com/ubuntu/ bionic main restricted
# deb-src http://cn.archive.ubuntu.com/ubuntu/ bionic main restricted

## Major bug fix updates produced after the final release of the
## distribution.
deb http://cn.archive.ubuntu.com/ubuntu/ bionic-updates main restricted
# deb-src http://cn.archive.ubuntu.com/ubuntu/ bionic-updates main restricted

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team. Also, please note that software in universe WILL NOT receive any
## review or updates from the Ubuntu security team.
deb http://cn.archive.ubuntu.com/ubuntu/ bionic universe
# deb-src http://cn.archive.ubuntu.com/ubuntu/ bionic universe
deb http://cn.archive.ubuntu.com/ubuntu/ bionic-updates universe
# deb-src http://cn.archive.ubuntu.com/ubuntu/ bionic-updates universe

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu 
## team, and may not be under a free licence. Please satisfy yourself as to 
## your rights to use the software. Also, please note that software in 
## multiverse WILL NOT receive any review or updates from the Ubuntu
## security team.
deb http://cn.archive.ubuntu.com/ubuntu/ bionic multiverse
# deb-src http://cn.archive.ubuntu.com/ubuntu/ bionic multiverse
deb http://cn.archive.ubuntu.com/ubuntu/ bionic-updates multiverse
# deb-src http://cn.archive.ubuntu.com/ubuntu/ bionic-updates multiverse

## N.B. software from this repository may not have been tested as
## extensively as that contained in the main release, although it includes
## newer versions of some applications which may provide useful features.
## Also, please note that software in backports WILL NOT receive any review
## or updates from the Ubuntu security team.
deb http://cn.archive.ubuntu.com/ubuntu/ bionic-backports main restricted universe multiverse
# deb-src http://cn.archive.ubuntu.com/ubuntu/ bionic-backports main restricted universe multiverse

## Uncomment the following two lines to add software from Canonical's
## 'partner' repository.
## This software is not part of Ubuntu, but is offered by Canonical and the
## respective vendors as a service to Ubuntu users.
# deb http://archive.canonical.com/ubuntu bionic partner
# deb-src http://archive.canonical.com/ubuntu bionic partner

deb http://security.ubuntu.com/ubuntu bionic-security main restricted
# deb-src http://security.ubuntu.com/ubuntu bionic-security main restricted
deb http://security.ubuntu.com/ubuntu bionic-security universe
# deb-src http://security.ubuntu.com/ubuntu bionic-security universe
deb http://security.ubuntu.com/ubuntu bionic-security multiverse
# deb-src http://security.ubuntu.com/ubuntu bionic-security multiverse
```

backup
```
ubuntu@masterk8s:~$ sudo cp /etc/apt/sources.list /etc/apt/sources.list.orig
```

overwrite
```
ubuntu@masterk8s:~$ sudo curl -jkSL http://192.168.0.106:48081/etc0x2Fapt0x2Fsources.list.ubuntu18.04cn -o /etc/apt/sources.list
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2905  100  2905    0     0   567k      0 --:--:-- --:--:-- --:--:--  567k
```

```
ubuntu@masterk8s:~$ sudo apt-get update
Get:1 http://security.ubuntu.com/ubuntu bionic-security InRelease [88.7 kB]
Get:2 http://cn.archive.ubuntu.com/ubuntu bionic InRelease [242 kB]
Get:3 http://security.ubuntu.com/ubuntu bionic-security/universe amd64 Packages [636 kB]
Get:4 http://security.ubuntu.com/ubuntu bionic-security/universe Translation-en [215 kB]
Get:5 http://cn.archive.ubuntu.com/ubuntu bionic-updates InRelease [88.7 kB]
Get:6 http://security.ubuntu.com/ubuntu bionic-security/multiverse amd64 Packages [6328 B]
Get:7 http://security.ubuntu.com/ubuntu bionic-security/multiverse Translation-en [2640 B]
Get:8 http://cn.archive.ubuntu.com/ubuntu bionic-backports InRelease [74.6 kB] 
Get:9 http://cn.archive.ubuntu.com/ubuntu bionic/main amd64 Packages [1019 kB]
Get:10 http://cn.archive.ubuntu.com/ubuntu bionic/main Translation-en [516 kB]
Get:11 http://cn.archive.ubuntu.com/ubuntu bionic/restricted amd64 Packages [9184 B]
Get:12 http://cn.archive.ubuntu.com/ubuntu bionic/restricted Translation-en [3584 B]
Get:13 http://cn.archive.ubuntu.com/ubuntu bionic/universe amd64 Packages [8570 kB]
Get:14 http://cn.archive.ubuntu.com/ubuntu bionic/universe Translation-en [4941 kB]
Get:15 http://cn.archive.ubuntu.com/ubuntu bionic/multiverse amd64 Packages [151 kB]
Get:16 http://cn.archive.ubuntu.com/ubuntu bionic/multiverse Translation-en [108 kB]
Get:17 http://cn.archive.ubuntu.com/ubuntu bionic-updates/main amd64 Packages [843 kB]
Get:18 http://cn.archive.ubuntu.com/ubuntu bionic-updates/main Translation-en [295 kB]
Get:19 http://cn.archive.ubuntu.com/ubuntu bionic-updates/restricted amd64 Packages [29.2 kB]
Get:20 http://cn.archive.ubuntu.com/ubuntu bionic-updates/restricted Translation-en [7764 B]
Get:21 http://cn.archive.ubuntu.com/ubuntu bionic-updates/universe amd64 Packages [1046 kB]
Get:22 http://cn.archive.ubuntu.com/ubuntu bionic-updates/universe Translation-en [323 kB]
Get:23 http://cn.archive.ubuntu.com/ubuntu bionic-updates/multiverse amd64 Packages [9704 B]
Get:24 http://cn.archive.ubuntu.com/ubuntu bionic-updates/multiverse Translation-en [4572 B]
Get:25 http://cn.archive.ubuntu.com/ubuntu bionic-backports/main amd64 Packages [2512 B]
Get:26 http://cn.archive.ubuntu.com/ubuntu bionic-backports/main Translation-en [1644 B]
Get:27 http://cn.archive.ubuntu.com/ubuntu bionic-backports/universe amd64 Packages [4032 B]
Get:28 http://cn.archive.ubuntu.com/ubuntu bionic-backports/universe Translation-en [1900 B]
Fetched 19.2 MB in 9s (2046 kB/s)                                              
Reading package lists... Done
```

### Docker

To install latest docker, refer to https://docs.docker.com/install/linux/docker-ce/ubuntu/

install Ubuntu archived distribution
```
ubuntu@masterk8s:~$ sudo apt-cache madison docker.io
 docker.io | 18.09.7-0ubuntu1~18.04.4 | http://cn.archive.ubuntu.com/ubuntu bionic-updates/universe amd64 Packages
 docker.io | 18.09.7-0ubuntu1~18.04.4 | http://security.ubuntu.com/ubuntu bionic-security/universe amd64 Packages
 docker.io | 17.12.1-0ubuntu1 | http://cn.archive.ubuntu.com/ubuntu bionic/universe amd64 Packages
```

```
ubuntu@masterk8s:~$ sudo apt-get install -y docker.io
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following package was automatically installed and is no longer required:
  grub-pc-bin
Use 'sudo apt autoremove' to remove it.
The following additional packages will be installed:
  bridge-utils cgroupfs-mount containerd pigz runc ubuntu-fan
Suggested packages:
  ifupdown aufs-tools debootstrap docker-doc rinse zfs-fuse | zfsutils
The following NEW packages will be installed:
  bridge-utils cgroupfs-mount containerd docker.io pigz runc ubuntu-fan
0 upgraded, 7 newly installed, 0 to remove and 6 not upgraded.
Need to get 52.2 MB of archives.
After this operation, 257 MB of additional disk space will be used.
Get:1 http://cn.archive.ubuntu.com/ubuntu bionic/universe amd64 pigz amd64 2.4-1 [57.4 kB]
Get:2 http://cn.archive.ubuntu.com/ubuntu bionic/main amd64 bridge-utils amd64 1.5-15ubuntu1 [30.1 kB]
Get:3 http://cn.archive.ubuntu.com/ubuntu bionic/universe amd64 cgroupfs-mount all 1.4 [6320 B]
Get:4 http://cn.archive.ubuntu.com/ubuntu bionic-updates/universe amd64 runc amd64 1.0.0~rc7+git20190403.029124da-0ubuntu1~18.04.2 [1903 kB]
Get:5 http://cn.archive.ubuntu.com/ubuntu bionic-updates/universe amd64 containerd amd64 1.2.6-0ubuntu1~18.04.2 [19.4 MB]
Get:6 http://cn.archive.ubuntu.com/ubuntu bionic-updates/universe amd64 docker.io amd64 18.09.7-0ubuntu1~18.04.4 [30.7 MB]
Get:7 http://cn.archive.ubuntu.com/ubuntu bionic/main amd64 ubuntu-fan all 0.12.10 [34.7 kB]
Fetched 52.2 MB in 31s (1706 kB/s)                                             
Preconfiguring packages ...
Selecting previously unselected package pigz.
(Reading database ... 60037 files and directories currently installed.)
Preparing to unpack .../0-pigz_2.4-1_amd64.deb ...
Unpacking pigz (2.4-1) ...
Selecting previously unselected package bridge-utils.
Preparing to unpack .../1-bridge-utils_1.5-15ubuntu1_amd64.deb ...
Unpacking bridge-utils (1.5-15ubuntu1) ...
Selecting previously unselected package cgroupfs-mount.
Preparing to unpack .../2-cgroupfs-mount_1.4_all.deb ...
Unpacking cgroupfs-mount (1.4) ...
Selecting previously unselected package runc.
Preparing to unpack .../3-runc_1.0.0~rc7+git20190403.029124da-0ubuntu1~18.04.2_amd64.deb ...
Unpacking runc (1.0.0~rc7+git20190403.029124da-0ubuntu1~18.04.2) ...
Selecting previously unselected package containerd.
Preparing to unpack .../4-containerd_1.2.6-0ubuntu1~18.04.2_amd64.deb ...
Unpacking containerd (1.2.6-0ubuntu1~18.04.2) ...
Selecting previously unselected package docker.io.
Preparing to unpack .../5-docker.io_18.09.7-0ubuntu1~18.04.4_amd64.deb ...
Unpacking docker.io (18.09.7-0ubuntu1~18.04.4) ...
Selecting previously unselected package ubuntu-fan.
Preparing to unpack .../6-ubuntu-fan_0.12.10_all.deb ...
Unpacking ubuntu-fan (0.12.10) ...
Setting up runc (1.0.0~rc7+git20190403.029124da-0ubuntu1~18.04.2) ...
Setting up cgroupfs-mount (1.4) ...
Setting up containerd (1.2.6-0ubuntu1~18.04.2) ...
Created symlink /etc/systemd/system/multi-user.target.wants/containerd.service → /lib/systemd/system/containerd.service.
Setting up bridge-utils (1.5-15ubuntu1) ...
Setting up ubuntu-fan (0.12.10) ...
Created symlink /etc/systemd/system/multi-user.target.wants/ubuntu-fan.service → /lib/systemd/system/ubuntu-fan.service.
Setting up pigz (2.4-1) ...
Setting up docker.io (18.09.7-0ubuntu1~18.04.4) ...
Adding group `docker' (GID 115) ...
Done.
Created symlink /etc/systemd/system/sockets.target.wants/docker.socket → /lib/systemd/system/docker.socket.
Processing triggers for systemd (237-3ubuntu10.33) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
Processing triggers for ureadahead (0.100.0-21) ...
```

```
ubuntu@masterk8s:~$ sudo usermod -aG docker ubuntu
```

```
ubuntu@masterk8s:~$ sudo docker version
Client:
 Version:           18.09.7
 API version:       1.39
 Go version:        go1.10.1
 Git commit:        2d0083d
 Built:             Fri Aug 16 14:20:06 2019
 OS/Arch:           linux/amd64
 Experimental:      false

Server:
 Engine:
  Version:          18.09.7
  API version:      1.39 (minimum version 1.12)
  Go version:       go1.10.1
  Git commit:       2d0083d
  Built:            Wed Aug 14 19:41:23 2019
  OS/Arch:          linux/amd64
  Experimental:     false
```

Note, after re-login, no `sudo` required
```
ubuntu@masterk8s:~$ docker info
Containers: 0
 Running: 0
 Paused: 0
 Stopped: 0
Images: 7
Server Version: 18.09.7
Storage Driver: overlay2
 Backing Filesystem: extfs
 Supports d_type: true
 Native Overlay Diff: true
Logging Driver: json-file
Cgroup Driver: cgroupfs
Plugins:
 Volume: local
 Network: bridge host macvlan null overlay
 Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
Swarm: inactive
Runtimes: runc
Default Runtime: runc
Init Binary: docker-init
containerd version: 
runc version: N/A
init version: v0.18.0 (expected: fec3683b971d9c3ef73f284f176672c44b448662)
Security Options:
 apparmor
 seccomp
  Profile: default
Kernel Version: 4.15.0-76-generic
Operating System: Ubuntu 18.04.3 LTS
OSType: linux
Architecture: x86_64
CPUs: 2
Total Memory: 1.899GiB
Name: masterk8s
ID: E7BO:CLNZ:K4LQ:AIOZ:5PTH:3ATA:4UHU:GYFN:HTWW:F2V7:7X66:QSNY
Docker Root Dir: /var/lib/docker
Debug Mode (client): false
Debug Mode (server): false
Registry: https://index.docker.io/v1/
Labels:
Experimental: false
Insecure Registries:
 127.0.0.0/8
Live Restore Enabled: false

WARNING: No swap limit support
```

#### CGroup driver

It is suggested by kubernetes
```
ubuntu@masterk8s:~$ sudo bash -c 'cat > /etc/docker/daemon.json <<EOF
> {
>   "exec-opts": ["native.cgroupdriver=systemd"],
>   "log-driver": "json-file",
>   "log-opts": {
>     "max-size": "100m"
>   },
>   "storage-driver": "overlay2"
> }
> EOF'
```

```
ubuntu@masterk8s:~$ sudo mkdir -p /etc/systemd/system/docker.service.d
```

```
ubuntu@masterk8s:~$ sudo systemctl daemon-reload
```

```
ubuntu@masterk8s:~$ sudo systemctl restart docker
```

```
ubuntu@masterk8s:~$ sudo systemctl status docker
● docker.service - Docker Application Container Engine
   Loaded: loaded (/lib/systemd/system/docker.service; disabled; vendor preset: 
   Active: active (running) since Mon 2020-02-03 23:08:45 UTC; 6s ago
     Docs: https://docs.docker.com
 Main PID: 9736 (dockerd)
    Tasks: 9
   CGroup: /system.slice/docker.service
           └─9736 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/contain

Feb 03 23:08:45 masterk8s dockerd[9736]: time="2020-02-03T23:08:45.049778261Z" l
Feb 03 23:08:45 masterk8s dockerd[9736]: time="2020-02-03T23:08:45.049785465Z" l
Feb 03 23:08:45 masterk8s dockerd[9736]: time="2020-02-03T23:08:45.050091975Z" l
Feb 03 23:08:45 masterk8s dockerd[9736]: time="2020-02-03T23:08:45.130806543Z" l
Feb 03 23:08:45 masterk8s dockerd[9736]: time="2020-02-03T23:08:45.305514273Z" l
Feb 03 23:08:45 masterk8s dockerd[9736]: time="2020-02-03T23:08:45.313370032Z" l
Feb 03 23:08:45 masterk8s dockerd[9736]: time="2020-02-03T23:08:45.373898621Z" l
Feb 03 23:08:45 masterk8s dockerd[9736]: time="2020-02-03T23:08:45.373972169Z" l
Feb 03 23:08:45 masterk8s dockerd[9736]: time="2020-02-03T23:08:45.385315975Z" l
Feb 03 23:08:45 masterk8s systemd[1]: Started Docker Application Container Engin
```


## Local Kubernetes repository

https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

Import gpg key （该key只能用于16.04，见下）
```
ubuntu@masterk8s:~$ curl -jkSL http://192.168.0.106:48081/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   663  100   663    0     0   107k      0 --:--:-- --:--:-- --:--:--  107k
OK
```

Import local repository apt source
```
ubuntu@masterk8s:~$ sudo curl -jkSL http://192.168.0.106:48081/etc0x2Fapt0x2Fkubernetes.list -o /etc/apt/sources.list.d/kubernetes.list
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    53  100    53    0     0  10600      0 --:--:-- --:--:-- --:--:-- 10600
```

```
ubuntu@masterk8s:~$ sudo sed -i "s%$deb http://apt.kubernetes.io/ kubernetes-xenial main$%# &\ndeb http://192.168.0.106:48081/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main%" /etc/apt/sources.list.d/kubernetes.list 
```

Trouble: The import key在18.04上无效
```
ubuntu@masterk8s:~$ sudo apt-get update
Get:1 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8993 B]
Err:1 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease
  The following signatures couldn't be verified because the public key is not available: NO_PUBKEY 6A030B21BA07F4FB
Hit:2 http://cn.archive.ubuntu.com/ubuntu bionic InRelease                     
Hit:3 http://cn.archive.ubuntu.com/ubuntu bionic-updates InRelease      
Hit:4 http://cn.archive.ubuntu.com/ubuntu bionic-backports InRelease     
Hit:5 http://security.ubuntu.com/ubuntu bionic-security InRelease
Reading package lists... Done
W: GPG error: http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease: The following signatures couldn't be verified because the public key is not available: NO_PUBKEY 6A030B21BA07F4FB
E: The repository 'http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease' is not signed.
N: Updating from such a repository can't be done securely, and is therefore disabled by default.
N: See apt-secure(8) manpage for repository creation and user configuration details.
```

Link to https://stackoverflow.com/questions/49877401/apt-get-update-error-related-with-kubeadm
```
ubuntu@masterk8s:~$ sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 6A030B21BA07F4FB
Executing: /tmp/apt-key-gpghome.sKb2wVhmzt/gpg.1.sh --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 6A030B21BA07F4FB
gpg: key 6A030B21BA07F4FB: public key "Google Cloud Packages Automatic Signing Key <gc-team@google.com>" imported
gpg: Total number processed: 1
gpg:               imported: 1

ubuntu@masterk8s:~$ sudo apt-get update
Get:1 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8993 B]
Ign:2 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
Get:2 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [33.3 kB]
Hit:3 http://cn.archive.ubuntu.com/ubuntu bionic InRelease                     
Hit:4 http://cn.archive.ubuntu.com/ubuntu bionic-updates InRelease             
Hit:5 http://security.ubuntu.com/ubuntu bionic-security InRelease        
Hit:6 http://cn.archive.ubuntu.com/ubuntu bionic-backports InRelease     
Fetched 42.3 kB in 1s (36.2 kB/s)
Reading package lists... Done
```

安装
```
ubuntu@masterk8s:~$ sudo apt-get install -y kubectl kubeadm kubelet
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following package was automatically installed and is no longer required:
  grub-pc-bin
Use 'sudo apt autoremove' to remove it.
The following additional packages will be installed:
  conntrack cri-tools kubernetes-cni socat
The following NEW packages will be installed:
  conntrack cri-tools kubeadm kubectl kubelet kubernetes-cni socat
0 upgraded, 7 newly installed, 0 to remove and 18 not upgraded.
Need to get 51.6 MB of archives.
After this operation, 272 MB of additional disk space will be used.
Get:1 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 cri-tools amd64 1.13.0-00 [8776 kB]
Get:2 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubernetes-cni amd64 0.7.5-00 [6473 kB]
Get:3 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubelet amd64 1.17.2-00 [19.2 MB]
Get:4 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubectl amd64 1.17.2-00 [8738 kB]
Get:5 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.17.2-00 [8061 kB]
Get:6 http://cn.archive.ubuntu.com/ubuntu bionic/main amd64 conntrack amd64 1:1.4.4+snapshot20161117-6ubuntu2 [30.6 kB]
Get:7 http://cn.archive.ubuntu.com/ubuntu bionic/main amd64 socat amd64 1.7.3.2-2ubuntu2 [342 kB]
Fetched 51.6 MB in 2s (24.1 MB/s)
Selecting previously unselected package conntrack.
(Reading database ... 60357 files and directories currently installed.)
Preparing to unpack .../0-conntrack_1%3a1.4.4+snapshot20161117-6ubuntu2_amd64.deb ...
Unpacking conntrack (1:1.4.4+snapshot20161117-6ubuntu2) ...
Selecting previously unselected package cri-tools.
Preparing to unpack .../1-cri-tools_1.13.0-00_amd64.deb ...
Unpacking cri-tools (1.13.0-00) ...
Selecting previously unselected package kubernetes-cni.
Preparing to unpack .../2-kubernetes-cni_0.7.5-00_amd64.deb ...
Unpacking kubernetes-cni (0.7.5-00) ...
Selecting previously unselected package socat.
Preparing to unpack .../3-socat_1.7.3.2-2ubuntu2_amd64.deb ...
Unpacking socat (1.7.3.2-2ubuntu2) ...
Selecting previously unselected package kubelet.
Preparing to unpack .../4-kubelet_1.17.2-00_amd64.deb ...
Unpacking kubelet (1.17.2-00) ...
Selecting previously unselected package kubectl.
Preparing to unpack .../5-kubectl_1.17.2-00_amd64.deb ...
Unpacking kubectl (1.17.2-00) ...
Selecting previously unselected package kubeadm.
Preparing to unpack .../6-kubeadm_1.17.2-00_amd64.deb ...
Unpacking kubeadm (1.17.2-00) ...
Setting up conntrack (1:1.4.4+snapshot20161117-6ubuntu2) ...
Setting up kubernetes-cni (0.7.5-00) ...
Setting up cri-tools (1.13.0-00) ...
Setting up socat (1.7.3.2-2ubuntu2) ...
Setting up kubelet (1.17.2-00) ...
Created symlink /etc/systemd/system/multi-user.target.wants/kubelet.service → /lib/systemd/system/kubelet.service.
Setting up kubectl (1.17.2-00) ...
Setting up kubeadm (1.17.2-00) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
```

Behind installation
```
ubuntu@masterk8s:~$ sudo cat /lib/systemd/system/kubelet.service
[Unit]
Description=kubelet: The Kubernetes Node Agent
Documentation=https://kubernetes.io/docs/home/

[Service]
ExecStart=/usr/bin/kubelet
Restart=always
StartLimitInterval=0
RestartSec=10

[Install]
WantedBy=multi-user.target
```

```
ubuntu@masterk8s:~$ ls /opt/
cni  containerd
```

```
ubuntu@masterk8s:~$ sudo ls -R /etc/kubernetes/
/etc/kubernetes/:
manifests

/etc/kubernetes/manifests:
```

```
ubuntu@masterk8s:~$ sudo cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
# Note: This dropin only works with kubeadm and kubelet v1.11+
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_CONFIG_ARGS=--config=/var/lib/kubelet/config.yaml"
# This is a file that "kubeadm init" and "kubeadm join" generates at runtime, populating the KUBELET_KUBEADM_ARGS variable dynamically
EnvironmentFile=-/var/lib/kubelet/kubeadm-flags.env
# This is a file that the user can use for overrides of the kubelet args as a last resort. Preferably, the user should use
# the .NodeRegistration.KubeletExtraArgs object in the configuration files instead. KUBELET_EXTRA_ARGS should be sourced from this file.
EnvironmentFile=-/etc/default/kubelet
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_CONFIG_ARGS $KUBELET_KUBEADM_ARGS $KUBELET_EXTRA_ARGS
```

### Cont.

Command _kubectl_
```
ubuntu@masterk8s:~$ which kubectl
/usr/bin/kubectl
```

```
ubuntu@masterk8s:~$ kubectl version --client
Client Version: version.Info{Major:"1", Minor:"17", GitVersion:"v1.17.2", GitCommit:"59603c6e503c87169aea6106f57b9f242f64df89", GitTreeState:"clean", BuildDate:"2020-01-18T23:30:10Z", GoVersion:"go1.13.5", Compiler:"gc", Platform:"linux/amd64"}
```

Load local image tarball from file server
```
ubuntu@masterk8s:~$ curl http://192.168.0.106:48080/kubernetes/server/bin/kubernetes/server/bin/kube-apiserver.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    19  100    19    0     0   3166      0 --:--:-- --:--:-- --:--:--  3166
Error processing tar file(exit status 1): unexpected EOF
```

加载docker镜像tar失败
+ https://github.com/moby/moby/issues/37581
+ https://stackoverflow.com/questions/42784396/docker-error-error-processing-tar-fileexit-status-1-unexpected-eof

改用kubeadm配置式

Warning may ignore
```
ubuntu@masterk8s:~$ sudo kubeadm config images list
W0203 20:32:21.996977   26480 version.go:101] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
W0203 20:32:21.997144   26480 version.go:102] falling back to the local client version: v1.17.2
W0203 20:32:21.997428   26480 validation.go:28] Cannot validate kube-proxy config - no validator is available
W0203 20:32:21.997454   26480 validation.go:28] Cannot validate kubelet config - no validator is available
k8s.gcr.io/kube-apiserver:v1.17.2
k8s.gcr.io/kube-controller-manager:v1.17.2
k8s.gcr.io/kube-scheduler:v1.17.2
k8s.gcr.io/kube-proxy:v1.17.2
k8s.gcr.io/pause:3.1
k8s.gcr.io/etcd:3.4.3-0
k8s.gcr.io/coredns:1.6.5
```

从阿里镜像库下载
```
ubuntu@masterk8s:~$ sudo kubeadm config images pull --image-repository registry.cn-hangzhou.aliyuncs.com/google_containers
W0203 20:46:40.387976   28917 version.go:101] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
W0203 20:46:40.388145   28917 version.go:102] falling back to the local client version: v1.17.2
W0203 20:46:40.388342   28917 validation.go:28] Cannot validate kube-proxy config - no validator is available
W0203 20:46:40.388362   28917 validation.go:28] Cannot validate kubelet config - no validator is available
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:v1.17.2
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:v1.17.2
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:v1.17.2
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.17.2
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:3.4.3-0
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.6.5
```

```
ubuntu@masterk8s:~$ sudo docker images registry.cn-hangzhou.aliyuncs.com/google_containers/*
REPOSITORY                                                                    TAG                 IMAGE ID            CREATED             SIZE
registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy                v1.17.2             cba2a99699bd        2 weeks ago         116MB
registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver            v1.17.2             41ef50a5f06a        2 weeks ago         171MB
registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager   v1.17.2             da5fd66c4068        2 weeks ago         161MB
registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler            v1.17.2             f52d4c527ef2        2 weeks ago         94.4MB
registry.cn-hangzhou.aliyuncs.com/google_containers/coredns                   1.6.5               70f311871ae1        3 months ago        41.6MB
registry.cn-hangzhou.aliyuncs.com/google_containers/etcd                      3.4.3-0             303ce5db0e90        3 months ago        288MB
registry.cn-hangzhou.aliyuncs.com/google_containers/pause                     3.1                 da86e6ba6ca1        2 years ago         742kB
```


## Bootstrap

默认的clustre ip范围是10.96.0.0/12, cluster ip是VIP
POD ip范围不能与cluster ip范围重合, pod ip address is like Docker container IP

This is showed the Docker use default cgroup driver _cgroupfs_, and require 2 CPU cores （实际单cpu也是可以部署的，见下）
```
ubuntu@masterk8s:~$ sudo kubeadm init --pod-network-cidr=10.10.0.0/16 --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers
W0203 21:59:34.644247    8771 version.go:101] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
W0203 21:59:34.644502    8771 version.go:102] falling back to the local client version: v1.17.2
W0203 21:59:34.646792    8771 validation.go:28] Cannot validate kube-proxy config - no validator is available
W0203 21:59:34.649296    8771 validation.go:28] Cannot validate kubelet config - no validator is available
[init] Using Kubernetes version: v1.17.2
[preflight] Running pre-flight checks
	[WARNING Service-Docker]: docker service is not enabled, please run 'systemctl enable docker.service'
	[WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
error execution phase preflight: [preflight] Some fatal errors occurred:
	[ERROR NumCPU]: the number of available CPUs 1 is less than the required 2
[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
To see the stack trace of this error execute with --v=5 or higher
```

This should be deployed in baremetal machine, 这时 Kubernetes API server 仅仅监听服务器上配置的局域网IP地址
```
ubuntu@masterk8s:~$ sudo kubeadm init --pod-network-cidr=10.10.0.0/16 --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers
W0204 07:21:03.198249   11037 version.go:101] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
W0204 07:21:03.198415   11037 version.go:102] falling back to the local client version: v1.17.2
W0204 07:21:03.199061   11037 validation.go:28] Cannot validate kube-proxy config - no validator is available
W0204 07:21:03.199107   11037 validation.go:28] Cannot validate kubelet config - no validator is available
[init] Using Kubernetes version: v1.17.2
[preflight] Running pre-flight checks
	[WARNING Service-Docker]: docker service is not enabled, please run 'systemctl enable docker.service'
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [masterk8s kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.20.30.56]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [masterk8s localhost] and IPs [10.20.30.56 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [masterk8s localhost] and IPs [10.20.30.56 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
W0204 07:21:06.012738   11037 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[control-plane] Creating static Pod manifest for "kube-scheduler"
W0204 07:21:06.013705   11037 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 16.507067 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.17" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node masterk8s as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node masterk8s as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: wrt5m7.dwltjhgxb0vruusi
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 10.20.30.56:6443 --token wrt5m7.dwltjhgxb0vruusi \
    --discovery-token-ca-cert-hash sha256:173f4de9b8a979b50e497524e267b714c97c3021bd1ab816e576037ec1231d57 
```

Tear down
```
ubuntu@masterk8s:~$ sudo kubeadm reset
[reset] Reading configuration from the cluster...
[reset] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -oyaml'
[reset] WARNING: Changes made to this host by 'kubeadm init' or 'kubeadm join' will be reverted.
[reset] Are you sure you want to proceed? [y/N]: y
[preflight] Running pre-flight checks
[reset] Removing info for node "masterk8s" from the ConfigMap "kubeadm-config" in the "kube-system" Namespace
W0204 08:14:18.637386   12662 removeetcdmember.go:61] [reset] failed to remove etcd member: error syncing endpoints with etc: etcdclient: no available endpoints
.Please manually remove this etcd member using etcdctl
[reset] Stopping the kubelet service
[reset] Unmounting mounted directories in "/var/lib/kubelet"
[reset] Deleting contents of config directories: [/etc/kubernetes/manifests /etc/kubernetes/pki]
[reset] Deleting files: [/etc/kubernetes/admin.conf /etc/kubernetes/kubelet.conf /etc/kubernetes/bootstrap-kubelet.conf /etc/kubernetes/controller-manager.conf /etc/kubernetes/scheduler.conf]
[reset] Deleting contents of stateful directories: [/var/lib/etcd /var/lib/kubelet /var/lib/dockershim /var/run/kubernetes /var/lib/cni]

The reset process does not clean CNI configuration. To do so, you must remove /etc/cni/net.d

The reset process does not reset or clean up iptables rules or IPVS tables.
If you wish to reset iptables, you must do so manually by using the "iptables" command.

If your cluster was setup to utilize IPVS, run ipvsadm --clear (or similar)
to reset your system's IPVS tables.

The reset process does not clean your kubeconfig files and you must remove them manually.
Please, check the contents of the $HOME/.kube/config file.
```

init参数错误

kubeadm --apiserver-advertise-adress not working
＋ https://github.com/kubernetes/kubernetes/issues/82943
＋ https://github.com/kubernetes/kubeadm/issues/1390

因为apiserver advertise address是OpenStack floating IP (DNAT)
```
ubuntu@masterk8s:~$ sudo kubeadm init --apiserver-advertise-address=192.168.192.123 --apiserver-cert-extra-sans=10.20.30.56,masterk8s.openstacklocal --pod-network-cidr=10.10.0.0/16 --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers
W0204 08:23:25.760501   13194 version.go:101] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
W0204 08:23:25.760669   13194 version.go:102] falling back to the local client version: v1.17.2
W0204 08:23:25.760942   13194 validation.go:28] Cannot validate kube-proxy config - no validator is available
W0204 08:23:25.760981   13194 validation.go:28] Cannot validate kubelet config - no validator is available
[init] Using Kubernetes version: v1.17.2
[preflight] Running pre-flight checks
	[WARNING Service-Docker]: docker service is not enabled, please run 'systemctl enable docker.service'
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [masterk8s kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local masterk8s.openstacklocal] and IPs [10.96.0.1 192.168.192.123 10.20.30.56]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [masterk8s localhost] and IPs [192.168.192.123 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [masterk8s localhost] and IPs [192.168.192.123 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
W0204 08:23:28.472342   13194 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[control-plane] Creating static Pod manifest for "kube-scheduler"
W0204 08:23:28.533335   13194 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[kubelet-check] Initial timeout of 40s passed.

Unfortunately, an error has occurred:
	timed out waiting for the condition

This error is likely caused by:
	- The kubelet is not running
	- The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)

If you are on a systemd-powered system, you can try to troubleshoot the error with the following commands:
	- 'systemctl status kubelet'
	- 'journalctl -xeu kubelet'

Additionally, a control plane component may have crashed or exited when started by the container runtime.
To troubleshoot, list all containers using your preferred container runtimes CLI, e.g. docker.
Here is one example how you may list all Kubernetes containers running in docker:
	- 'docker ps -a | grep kube | grep -v pause'
	Once you have found the failing container, you can inspect its logs with:
	- 'docker logs CONTAINERID'
error execution phase wait-control-plane: couldn't initialize a Kubernetes cluster
To see the stack trace of this error execute with --v=5 or higher
```

把OpenStack floating IP加到自签名证书的SAN
```
ubuntu@masterk8s:~$ sudo kubeadm init --apiserver-cert-extra-sans=192.168.192.123,masterk8s.openstacklocal --pod-network-cidr=10.10.0.0/16 --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers
W0204 09:57:38.054946   16880 version.go:101] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
W0204 09:57:38.055127   16880 version.go:102] falling back to the local client version: v1.17.2
W0204 09:57:38.055518   16880 validation.go:28] Cannot validate kube-proxy config - no validator is available
W0204 09:57:38.055572   16880 validation.go:28] Cannot validate kubelet config - no validator is available
[init] Using Kubernetes version: v1.17.2
[preflight] Running pre-flight checks
	[WARNING Service-Docker]: docker service is not enabled, please run 'systemctl enable docker.service'
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [masterk8s kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local masterk8s.openstacklocal] and IPs [10.96.0.1 10.20.30.56 192.168.192.123]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [masterk8s localhost] and IPs [10.20.30.56 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [masterk8s localhost] and IPs [10.20.30.56 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
W0204 09:57:40.582350   16880 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[control-plane] Creating static Pod manifest for "kube-scheduler"
W0204 09:57:40.583017   16880 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 14.506793 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.17" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node masterk8s as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node masterk8s as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: n50n2x.0j5gz5h0d3z7euns
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 10.20.30.56:6443 --token n50n2x.0j5gz5h0d3z7euns \
    --discovery-token-ca-cert-hash sha256:76f24e6e089f9df910cd02d4003627a09f87e9df3cc2e2ffd1f08faa5cb11d0f 
```


Service status
```
ubuntu@masterk8s:~$ sudo systemctl status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/lib/systemd/system/kubelet.service; enabled; vendor preset: 
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Tue 2020-02-04 07:21:23 UTC; 1min 13s ago
     Docs: https://kubernetes.io/docs/home/
 Main PID: 11997 (kubelet)
    Tasks: 16 (limit: 2303)
   CGroup: /system.slice/kubelet.service
           └─11997 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/boots

Feb 04 07:22:24 masterk8s kubelet[11997]: W0204 07:22:24.075421   11997 cni.go:2
Feb 04 07:22:24 masterk8s kubelet[11997]: E0204 07:22:24.305041   11997 kubelet.
Feb 04 07:22:26 masterk8s kubelet[11997]: E0204 07:22:26.216006   11997 dns.go:1
Feb 04 07:22:28 masterk8s kubelet[11997]: E0204 07:22:28.209062   11997 dns.go:1
Feb 04 07:22:29 masterk8s kubelet[11997]: W0204 07:22:29.075878   11997 cni.go:2
Feb 04 07:22:29 masterk8s kubelet[11997]: E0204 07:22:29.307808   11997 kubelet.
Feb 04 07:22:32 masterk8s kubelet[11997]: E0204 07:22:32.208336   11997 dns.go:1
Feb 04 07:22:34 masterk8s kubelet[11997]: W0204 07:22:34.076279   11997 cni.go:2
Feb 04 07:22:34 masterk8s kubelet[11997]: E0204 07:22:34.309435   11997 kubelet.
Feb 04 07:22:35 masterk8s kubelet[11997]: E0204 07:22:35.207847   11997 dns.go:1
lines 1-21/21 (END)
```

Docker containers
```
ubuntu@masterk8s:~$ docker ps
CONTAINER ID        IMAGE                                                           COMMAND                  CREATED              STATUS              PORTS               NAMES
8fd7c37ce310        cba2a99699bd                                                    "/usr/local/bin/kube…"   About a minute ago   Up About a minute                       k8s_kube-proxy_kube-proxy-zkqkl_kube-system_aa57d9c7-44d3-4b63-90a5-27668dd958e4_0
a3354c53b204        registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1   "/pause"                 About a minute ago   Up About a minute                       k8s_POD_kube-proxy-zkqkl_kube-system_aa57d9c7-44d3-4b63-90a5-27668dd958e4_0
368e5e2c1ce3        f52d4c527ef2                                                    "kube-scheduler --au…"   About a minute ago   Up About a minute                       k8s_kube-scheduler_kube-scheduler-masterk8s_kube-system_92e6721871b55a1e38e273dab9a560bc_0
01f803c0c9d1        41ef50a5f06a                                                    "kube-apiserver --ad…"   About a minute ago   Up About a minute                       k8s_kube-apiserver_kube-apiserver-masterk8s_kube-system_0fee9326c22f89b300ac7e262a2e29a8_0
a4f827738e42        303ce5db0e90                                                    "etcd --advertise-cl…"   About a minute ago   Up About a minute                       k8s_etcd_etcd-masterk8s_kube-system_f4d4153d81aeb40526b564e8055a6992_0
4a430e156d43        da5fd66c4068                                                    "kube-controller-man…"   About a minute ago   Up About a minute                       k8s_kube-controller-manager_kube-controller-manager-masterk8s_kube-system_8dc74fad81ada88730b09dcc967277e1_0
9096a9d10809        registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1   "/pause"                 About a minute ago   Up About a minute                       k8s_POD_kube-apiserver-masterk8s_kube-system_0fee9326c22f89b300ac7e262a2e29a8_0
dfd685052dd4        registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1   "/pause"                 About a minute ago   Up About a minute                       k8s_POD_etcd-masterk8s_kube-system_f4d4153d81aeb40526b564e8055a6992_0
b9ff38dac077        registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1   "/pause"                 About a minute ago   Up About a minute                       k8s_POD_kube-scheduler-masterk8s_kube-system_92e6721871b55a1e38e273dab9a560bc_0
9bdfee522d76        registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1   "/pause"                 About a minute ago   Up About a minute                       k8s_POD_kube-controller-manager-masterk8s_kube-system_8dc74fad81ada88730b09dcc967277e1_0
```

kubeconfig
```
ubuntu@masterk8s:~$ mkdir -p $HOME/.kube
ubuntu@masterk8s:~$ sudo cp /etc/kubernetes/admin.conf $HOME/.kube/config
ubuntu@masterk8s:~$ sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

Command _kubectl_
```
ubuntu@masterk8s:~$ kubectl get nodes
NAME        STATUS     ROLES    AGE     VERSION
masterk8s   NotReady   master   6m45s   v1.17.2
```

```
ubuntu@masterk8s:~$ kubectl version
Client Version: version.Info{Major:"1", Minor:"17", GitVersion:"v1.17.2", GitCommit:"59603c6e503c87169aea6106f57b9f242f64df89", GitTreeState:"clean", BuildDate:"2020-01-18T23:30:10Z", GoVersion:"go1.13.5", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"17", GitVersion:"v1.17.2", GitCommit:"59603c6e503c87169aea6106f57b9f242f64df89", GitTreeState:"clean", BuildDate:"2020-01-18T23:22:30Z", GoVersion:"go1.13.5", Compiler:"gc", Platform:"linux/amd64"}
```

```
ubuntu@masterk8s:~$ kubectl api-versions
admissionregistration.k8s.io/v1
admissionregistration.k8s.io/v1beta1
apiextensions.k8s.io/v1
apiextensions.k8s.io/v1beta1
apiregistration.k8s.io/v1
apiregistration.k8s.io/v1beta1
apps/v1
authentication.k8s.io/v1
authentication.k8s.io/v1beta1
authorization.k8s.io/v1
authorization.k8s.io/v1beta1
autoscaling/v1
autoscaling/v2beta1
autoscaling/v2beta2
batch/v1
batch/v1beta1
certificates.k8s.io/v1beta1
coordination.k8s.io/v1
coordination.k8s.io/v1beta1
discovery.k8s.io/v1beta1
events.k8s.io/v1beta1
extensions/v1beta1
networking.k8s.io/v1
networking.k8s.io/v1beta1
node.k8s.io/v1beta1
policy/v1beta1
rbac.authorization.k8s.io/v1
rbac.authorization.k8s.io/v1beta1
scheduling.k8s.io/v1
scheduling.k8s.io/v1beta1
storage.k8s.io/v1
storage.k8s.io/v1beta1
v1
```

```
ubuntu@masterk8s:~$ kubectl get ns
NAME              STATUS   AGE
default           Active   2m13s
kube-node-lease   Active   2m14s
kube-public       Active   2m14s
kube-system       Active   2m14s
```

```
ubuntu@masterk8s:~$ kubectl get --all-namespaces all
NAMESPACE     NAME                                    READY   STATUS    RESTARTS   AGE
kube-system   pod/coredns-7f9c544f75-jhhq8            0/1     Pending   0          9m54s
kube-system   pod/coredns-7f9c544f75-ndkd8            0/1     Pending   0          9m54s
kube-system   pod/etcd-masterk8s                      1/1     Running   0          10m
kube-system   pod/kube-apiserver-masterk8s            1/1     Running   0          10m
kube-system   pod/kube-controller-manager-masterk8s   1/1     Running   0          10m
kube-system   pod/kube-proxy-zkqkl                    1/1     Running   0          9m54s
kube-system   pod/kube-scheduler-masterk8s            1/1     Running   0          10m

NAMESPACE     NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
default       service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP                  10m
kube-system   service/kube-dns     ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   10m

NAMESPACE     NAME                        DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                 AGE
kube-system   daemonset.apps/kube-proxy   1         1         1       1            1           beta.kubernetes.io/os=linux   10m

NAMESPACE     NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
kube-system   deployment.apps/coredns   0/2     2            0           10m

NAMESPACE     NAME                                 DESIRED   CURRENT   READY   AGE
kube-system   replicaset.apps/coredns-7f9c544f75   2         2         0       9m54s
```


## 在kubernetes所在网络之外使用kubectl

Using DNAT ip 

```
ubuntu@masterk8s:~$ sed "s/10.20.30.56/192.168.192.123/" $HOME/.kube/config > kubeconfig 
```

```
ubuntu@masterk8s:~$ kubectl --kubeconfig=kubeconfig get ns
NAME              STATUS   AGE
default           Active   7m12s
kube-node-lease   Active   7m13s
kube-public       Active   7m13s
kube-system       Active   7m13s
```

From Physical Machine
```
wei@wei-ThinkPad-X1-Extreme:~$ scp -i keypair-train.pem ubuntu@192.168.192.123:/home/ubuntu/kubeconfig $HOME/.kube/config
kubeconfig                                    100% 5455     3.4MB/s   00:00    
```

```
wei@wei-ThinkPad-X1-Extreme:~$ ./kubernetes/server/bin/kubectl get ns
NAME              STATUS   AGE
default           Active   13m
kube-node-lease   Active   13m
kube-public       Active   13m
kube-system       Active   13m
```

Or
```
wei@wei-ThinkPad-X1-Extreme:~$ mkdir -p $HOME/.kube && scp -i keypair-train.pem ubuntu@192.168.192.123:/home/ubuntu/.kube/config $HOME/kubeconfig
config                                        100% 5447     3.6MB/s   00:00    
wei@wei-ThinkPad-X1-Extreme:~$ sed -i "s/10.20.30.56/192.168.192.123/" kubeconfig 
```


部署control-plane HA集群采用control plane endpoint 
``
ubuntu@masterk8s:~$ sudo kubeadm init --control-plane-endpoint=192.168.192.123 --apiserver-cert-extra-sans=192.168.192.123,masterk8s.openstacklocal --pod-network-cidr=10.10.0.0/16 --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --ignore-preflight-errors=NumCPU --token-ttl=480h0m0s
W0204 10:22:36.895757   18875 version.go:101] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
W0204 10:22:36.896584   18875 version.go:102] falling back to the local client version: v1.17.2
W0204 10:22:36.897489   18875 validation.go:28] Cannot validate kube-proxy config - no validator is available
W0204 10:22:36.897996   18875 validation.go:28] Cannot validate kubelet config - no validator is available
[init] Using Kubernetes version: v1.17.2
[preflight] Running pre-flight checks
	[WARNING Service-Docker]: docker service is not enabled, please run 'systemctl enable docker.service'
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [masterk8s kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local masterk8s.openstacklocal] and IPs [10.96.0.1 10.20.30.56 192.168.192.123 192.168.192.123]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [masterk8s localhost] and IPs [10.20.30.56 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [masterk8s localhost] and IPs [10.20.30.56 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
W0204 10:22:40.064023   18875 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[control-plane] Creating static Pod manifest for "kube-scheduler"
W0204 10:22:40.065263   18875 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 15.002145 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.17" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node masterk8s as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node masterk8s as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: 9udxjm.1prjluofu8mw7zyv
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

You can now join any number of control-plane nodes by copying certificate authorities
and service account keys on each node and then running the following as root:

  kubeadm join 192.168.192.123:6443 --token 9udxjm.1prjluofu8mw7zyv \
    --discovery-token-ca-cert-hash sha256:e76a4e827e734a94738acb78edf4285393e73a5eb2731090422829071356ec49 \
    --control-plane 

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.192.123:6443 --token 9udxjm.1prjluofu8mw7zyv \
    --discovery-token-ca-cert-hash sha256:e76a4e827e734a94738acb78edf4285393e73a5eb2731090422829071356ec49 
```

Verify
```
ubuntu@masterk8s:~$ kubectl get ns
NAME              STATUS   AGE
default           Active   100s
kube-node-lease   Active   101s
kube-public       Active   101s
kube-system       Active   101s
```

```
ubuntu@masterk8s:~$ v=$(hostname -I | awk '{print $1}'); sed "s/192.168.192.123/$v/" $HOME/.kube/config > kubeconfig 
```

```
ubuntu@masterk8s:~$ kubectl --kubeconfig=kubeconfig get ns
NAME              STATUS   AGE
default           Active   6m8s
kube-node-lease   Active   6m9s
kube-public       Active   6m9s
kube-system       Active   6m9s
```

```
wei@wei-ThinkPad-X1-Extreme:~$ scp -i keypair-train.pem ubuntu@192.168.192.123:/home/ubuntu/.kube/config $HOME/.kube/config
config                                        100% 5455     8.0MB/s   00:00    
```

```
wei@wei-ThinkPad-X1-Extreme:~$ ./kubernetes/server/bin/kubectl get ns
NAME              STATUS   AGE
default           Active   7m5s
kube-node-lease   Active   7m6s
kube-public       Active   7m6s
kube-system       Active   7m6s
```

```
wei@wei-ThinkPad-X1-Extreme:~$ kubectl get nodes
NAME        STATUS     ROLES    AGE   VERSION
masterk8s   NotReady   master   47m   v1.17.2
```


使用--ignore-preflight-errors=NumCPU在单CPU上部署验证，和--token-ttl参数做批量机器部署
```
ubuntu@masterk8s:~$ sudo kubeadm init --apiserver-cert-extra-sans=192.168.192.123,masterk8s.openstacklocal --pod-network-cidr=10.10.0.0/16 --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --ignore-preflight-errors=NumCPU --token-ttl=480h0m0s
W0204 22:39:44.581880   31356 version.go:101] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
W0204 22:39:44.582681   31356 version.go:102] falling back to the local client version: v1.17.2
W0204 22:39:44.583694   31356 validation.go:28] Cannot validate kube-proxy config - no validator is available
W0204 22:39:44.584232   31356 validation.go:28] Cannot validate kubelet config - no validator is available
[init] Using Kubernetes version: v1.17.2
[preflight] Running pre-flight checks
	[WARNING Service-Docker]: docker service is not enabled, please run 'systemctl enable docker.service'
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [masterk8s kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local masterk8s.openstacklocal] and IPs [10.96.0.1 10.20.30.56 192.168.192.123]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [masterk8s localhost] and IPs [10.20.30.56 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [masterk8s localhost] and IPs [10.20.30.56 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
W0204 22:39:47.445872   31356 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[control-plane] Creating static Pod manifest for "kube-scheduler"
W0204 22:39:47.446577   31356 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 15.505626 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.17" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node masterk8s as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node masterk8s as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: h8cu4b.gb8znn0tj4vogkgd
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 10.20.30.56:6443 --token h8cu4b.gb8znn0tj4vogkgd \
    --discovery-token-ca-cert-hash sha256:80f3952e35b6e7c0c64e843d5404ebf8102c6d90d95a1869a9cc77768b2f327b 
```

