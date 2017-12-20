# DevOps

Using _kubeadm join_ [cluster by k8s-v1.9.0-deploy](../k8s-v1.9.0-deploy)

## Table of content

1. Prerequisites
   * docker
   * dockerized file server
   * hostname
   * hosts
   * selinux
   * firewalld
1. Images
   * Kubernetes
   * POD base
   * etcd
   * kube dns 
   * flannel
1. Install kubeadm from YUM mirror


## Prerequisites

docker
```
[vagrant@localhost ~]$ sudo dnf install -y docker
Failed to set locale, defaulting to C
Last metadata expiration check: 0:27:48 ago on Tue Dec 19 08:07:22 2017.
Dependencies resolved.
================================================================================================================================================
 Package                                       Arch                   Version                                     Repository               Size
================================================================================================================================================
Installing:
 docker                                        x86_64                 2:1.13.1-42.git4402c09.fc27                 updates                  20 M
Installing dependencies:
 atomic-registries                             x86_64                 1.20.1-6.fc27                               updates                  38 k
 container-selinux                             noarch                 2:2.36-1.fc27                               updates                  36 k
 container-storage-setup                       noarch                 0.8.0-2.git1d27ecf.fc27                     updates                  36 k
 device-mapper-event                           x86_64                 1.02.142-4.fc27                             fedora                  251 k
 device-mapper-event-libs                      x86_64                 1.02.142-4.fc27                             fedora                  251 k
 device-mapper-persistent-data                 x86_64                 0.7.5-1.fc27                                updates                 429 k
 docker-common                                 x86_64                 2:1.13.1-42.git4402c09.fc27                 updates                  85 k
 docker-rhel-push-plugin                       x86_64                 2:1.13.1-42.git4402c09.fc27                 updates                 1.7 M
 gnupg                                         x86_64                 1.4.22-3.fc27                               fedora                  1.3 M
 iptables                                      x86_64                 1.6.1-4.fc27                                fedora                  471 k
 libaio                                        x86_64                 0.3.110-9.fc27                              fedora                   29 k
 libnet                                        x86_64                 1.1.6-14.fc27                               fedora                   65 k
 libnetfilter_conntrack                        x86_64                 1.0.6-4.fc27                                fedora                   62 k
 libnfnetlink                                  x86_64                 1.0.1-11.fc27                               fedora                   31 k
 libusb                                        x86_64                 1:0.1.5-10.fc27                             fedora                   40 k
 lvm2                                          x86_64                 2.02.173-4.fc27                             fedora                  1.4 M
 lvm2-libs                                     x86_64                 2.02.173-4.fc27                             fedora                  1.1 M
 oci-umount                                    x86_64                 2:2.3.0-1.git51e7c50.fc27                   updates                  35 k
 policycoreutils-python-utils                  x86_64                 2.7-1.fc27                                  fedora                  223 k
 protobuf-c                                    x86_64                 1.2.1-7.fc27                                fedora                   34 k
 python-rhsm-certificates                      x86_64                 1.20.1-3.fc27                               fedora                   44 k
 python3-pytoml                                noarch                 0.1.14-2.git7dea353.fc27                    fedora                   23 k
 skopeo-containers                             x86_64                 0.1.27-1.git93876ac.fc27                    updates                  16 k
 systemd-container                             x86_64                 234-8.fc27                                  fedora                  422 k
 yajl                                          x86_64                 2.1.0-8.fc27                                fedora                   38 k
Installing weak dependencies:
 criu                                          x86_64                 3.6-1.fc27                                  updates                 456 k
 oci-register-machine                          x86_64                 0-5.11.gitcd1e331.fc27                      fedora                  1.0 M
 oci-systemd-hook                              x86_64                 1:0.1.13-1.gitafe4b4a.fc27                  fedora                   36 k

Transaction Summary
================================================================================================================================================
Install  29 Packages

Total download size: 30 M
Installed size: 95 M
Downloading Packages:
(1/29): atomic-registries-1.20.1-6.fc27.x86_64.rpm                                                               61 kB/s |  38 kB     00:00    
(2/29): docker-common-1.13.1-42.git4402c09.fc27.x86_64.rpm                                                       86 kB/s |  85 kB     00:00    
(3/29): skopeo-containers-0.1.27-1.git93876ac.fc27.x86_64.rpm                                                    94 kB/s |  16 kB     00:00    
(4/29): gnupg-1.4.22-3.fc27.x86_64.rpm                                                                          843 kB/s | 1.3 MB     00:01    
(5/29): python-rhsm-certificates-1.20.1-3.fc27.x86_64.rpm                                                       262 kB/s |  44 kB     00:00    
(6/29): python3-pytoml-0.1.14-2.git7dea353.fc27.noarch.rpm                                                      176 kB/s |  23 kB     00:00    
(7/29): libusb-0.1.5-10.fc27.x86_64.rpm                                                                         156 kB/s |  40 kB     00:00    
(8/29): lvm2-2.02.173-4.fc27.x86_64.rpm                                                                         481 kB/s | 1.4 MB     00:02    
(9/29): lvm2-libs-2.02.173-4.fc27.x86_64.rpm                                                                    463 kB/s | 1.1 MB     00:02    
(10/29): device-mapper-event-1.02.142-4.fc27.x86_64.rpm                                                         576 kB/s | 251 kB     00:00    
(11/29): device-mapper-event-libs-1.02.142-4.fc27.x86_64.rpm                                                    690 kB/s | 251 kB     00:00    
(12/29): device-mapper-persistent-data-0.7.5-1.fc27.x86_64.rpm                                                  156 kB/s | 429 kB     00:02    
(13/29): libaio-0.3.110-9.fc27.x86_64.rpm                                                                       121 kB/s |  29 kB     00:00    
(14/29): docker-rhel-push-plugin-1.13.1-42.git4402c09.fc27.x86_64.rpm                                           146 kB/s | 1.7 MB     00:12    
(15/29): container-selinux-2.36-1.fc27.noarch.rpm                                                                99 kB/s |  36 kB     00:00    
(16/29): policycoreutils-python-utils-2.7-1.fc27.x86_64.rpm                                                     458 kB/s | 223 kB     00:00    
(17/29): container-storage-setup-0.8.0-2.git1d27ecf.fc27.noarch.rpm                                              90 kB/s |  36 kB     00:00    
(18/29): oci-umount-2.3.0-1.git51e7c50.fc27.x86_64.rpm                                                          129 kB/s |  35 kB     00:00    
(19/29): yajl-2.1.0-8.fc27.x86_64.rpm                                                                            88 kB/s |  38 kB     00:00    
(20/29): libnetfilter_conntrack-1.0.6-4.fc27.x86_64.rpm                                                         236 kB/s |  62 kB     00:00    
(21/29): libnfnetlink-1.0.1-11.fc27.x86_64.rpm                                                                  237 kB/s |  31 kB     00:00    
(22/29): iptables-1.6.1-4.fc27.x86_64.rpm                                                                       300 kB/s | 471 kB     00:01    
(23/29): libnet-1.1.6-14.fc27.x86_64.rpm                                                                        153 kB/s |  65 kB     00:00    
(24/29): protobuf-c-1.2.1-7.fc27.x86_64.rpm                                                                     209 kB/s |  34 kB     00:00    
(25/29): criu-3.6-1.fc27.x86_64.rpm                                                                             132 kB/s | 456 kB     00:03    
(26/29): oci-systemd-hook-0.1.13-1.gitafe4b4a.fc27.x86_64.rpm                                                   181 kB/s |  36 kB     00:00    
(27/29): oci-register-machine-0-5.11.gitcd1e331.fc27.x86_64.rpm                                                 436 kB/s | 1.0 MB     00:02    
(28/29): systemd-container-234-8.fc27.x86_64.rpm                                                                461 kB/s | 422 kB     00:00    
(29/29): docker-1.13.1-42.git4402c09.fc27.x86_64.rpm                                                            556 kB/s |  20 MB     00:36    
------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                           754 kB/s |  30 MB     00:40     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                        1/1 
  Installing       : device-mapper-event-libs-1.02.142-4.fc27.x86_64                                                                       1/29 
  Running scriptlet: device-mapper-event-libs-1.02.142-4.fc27.x86_64                                                                       1/29 
  Installing       : libnfnetlink-1.0.1-11.fc27.x86_64                                                                                     2/29 
  Running scriptlet: libnfnetlink-1.0.1-11.fc27.x86_64                                                                                     2/29 
  Installing       : yajl-2.1.0-8.fc27.x86_64                                                                                              3/29 
  Running scriptlet: yajl-2.1.0-8.fc27.x86_64                                                                                              3/29 
  Installing       : oci-umount-2:2.3.0-1.git51e7c50.fc27.x86_64                                                                           4/29 
  Installing       : libnetfilter_conntrack-1.0.6-4.fc27.x86_64                                                                            5/29 
  Running scriptlet: libnetfilter_conntrack-1.0.6-4.fc27.x86_64                                                                            5/29 
  Installing       : iptables-1.6.1-4.fc27.x86_64                                                                                          6/29 
  Running scriptlet: iptables-1.6.1-4.fc27.x86_64                                                                                          6/29 
  Installing       : device-mapper-event-1.02.142-4.fc27.x86_64                                                                            7/29 
  Running scriptlet: device-mapper-event-1.02.142-4.fc27.x86_64                                                                            7/29 
  Installing       : lvm2-libs-2.02.173-4.fc27.x86_64                                                                                      8/29 
  Running scriptlet: lvm2-libs-2.02.173-4.fc27.x86_64                                                                                      8/29 
  Installing       : systemd-container-234-8.fc27.x86_64                                                                                   9/29 
  Installing       : protobuf-c-1.2.1-7.fc27.x86_64                                                                                       10/29 
  Running scriptlet: protobuf-c-1.2.1-7.fc27.x86_64                                                                                       10/29 
  Installing       : libnet-1.1.6-14.fc27.x86_64                                                                                          11/29 
  Running scriptlet: libnet-1.1.6-14.fc27.x86_64                                                                                          11/29 
  Installing       : policycoreutils-python-utils-2.7-1.fc27.x86_64                                                                       12/29 
  Installing       : container-selinux-2:2.36-1.fc27.noarch                                                                               13/29 
  Running scriptlet: container-selinux-2:2.36-1.fc27.noarch                                                                               13/29 
  Installing       : libaio-0.3.110-9.fc27.x86_64                                                                                         14/29 
  Running scriptlet: libaio-0.3.110-9.fc27.x86_64                                                                                         14/29 
  Installing       : device-mapper-persistent-data-0.7.5-1.fc27.x86_64                                                                    15/29 
  Installing       : lvm2-2.02.173-4.fc27.x86_64                                                                                          16/29 
  Running scriptlet: lvm2-2.02.173-4.fc27.x86_64                                                                                          16/29 
  Installing       : container-storage-setup-0.8.0-2.git1d27ecf.fc27.noarch                                                               17/29 
  Installing       : libusb-1:0.1.5-10.fc27.x86_64                                                                                        18/29 
  Running scriptlet: libusb-1:0.1.5-10.fc27.x86_64                                                                                        18/29 
  Installing       : gnupg-1.4.22-3.fc27.x86_64                                                                                           19/29 
  Running scriptlet: gnupg-1.4.22-3.fc27.x86_64                                                                                           19/29 
  Installing       : python3-pytoml-0.1.14-2.git7dea353.fc27.noarch                                                                       20/29 
  Installing       : atomic-registries-1.20.1-6.fc27.x86_64                                                                               21/29 
  Installing       : python-rhsm-certificates-1.20.1-3.fc27.x86_64                                                                        22/29 
  Installing       : skopeo-containers-0.1.27-1.git93876ac.fc27.x86_64                                                                    23/29 
  Installing       : docker-rhel-push-plugin-2:1.13.1-42.git4402c09.fc27.x86_64                                                           24/29 
  Running scriptlet: docker-rhel-push-plugin-2:1.13.1-42.git4402c09.fc27.x86_64                                                           24/29 
  Installing       : docker-common-2:1.13.1-42.git4402c09.fc27.x86_64                                                                     25/29 
  Installing       : docker-2:1.13.1-42.git4402c09.fc27.x86_64                                                                            26/29 
  Running scriptlet: docker-2:1.13.1-42.git4402c09.fc27.x86_64                                                                            26/29 
  Installing       : criu-3.6-1.fc27.x86_64                                                                                               27/29 
  Running scriptlet: criu-3.6-1.fc27.x86_64                                                                                               27/29 
  Installing       : oci-register-machine-0-5.11.gitcd1e331.fc27.x86_64                                                                   28/29 
  Installing       : oci-systemd-hook-1:0.1.13-1.gitafe4b4a.fc27.x86_64                                                                   29/29 
  Running scriptlet: docker-2:1.13.1-42.git4402c09.fc27.x86_64                                                                            29/29 
  Running scriptlet: oci-systemd-hook-1:0.1.13-1.gitafe4b4a.fc27.x86_64                                                                   29/29 
  Verifying        : docker-2:1.13.1-42.git4402c09.fc27.x86_64                                                                             1/29 
  Verifying        : atomic-registries-1.20.1-6.fc27.x86_64                                                                                2/29 
  Verifying        : docker-common-2:1.13.1-42.git4402c09.fc27.x86_64                                                                      3/29 
  Verifying        : docker-rhel-push-plugin-2:1.13.1-42.git4402c09.fc27.x86_64                                                            4/29 
  Verifying        : skopeo-containers-0.1.27-1.git93876ac.fc27.x86_64                                                                     5/29 
  Verifying        : gnupg-1.4.22-3.fc27.x86_64                                                                                            6/29 
  Verifying        : python-rhsm-certificates-1.20.1-3.fc27.x86_64                                                                         7/29 
  Verifying        : python3-pytoml-0.1.14-2.git7dea353.fc27.noarch                                                                        8/29 
  Verifying        : libusb-1:0.1.5-10.fc27.x86_64                                                                                         9/29 
  Verifying        : lvm2-2.02.173-4.fc27.x86_64                                                                                          10/29 
  Verifying        : lvm2-libs-2.02.173-4.fc27.x86_64                                                                                     11/29 
  Verifying        : device-mapper-event-1.02.142-4.fc27.x86_64                                                                           12/29 
  Verifying        : device-mapper-event-libs-1.02.142-4.fc27.x86_64                                                                      13/29 
  Verifying        : device-mapper-persistent-data-0.7.5-1.fc27.x86_64                                                                    14/29 
  Verifying        : libaio-0.3.110-9.fc27.x86_64                                                                                         15/29 
  Verifying        : container-selinux-2:2.36-1.fc27.noarch                                                                               16/29 
  Verifying        : policycoreutils-python-utils-2.7-1.fc27.x86_64                                                                       17/29 
  Verifying        : container-storage-setup-0.8.0-2.git1d27ecf.fc27.noarch                                                               18/29 
  Verifying        : oci-umount-2:2.3.0-1.git51e7c50.fc27.x86_64                                                                          19/29 
  Verifying        : yajl-2.1.0-8.fc27.x86_64                                                                                             20/29 
  Verifying        : iptables-1.6.1-4.fc27.x86_64                                                                                         21/29 
  Verifying        : libnetfilter_conntrack-1.0.6-4.fc27.x86_64                                                                           22/29 
  Verifying        : libnfnetlink-1.0.1-11.fc27.x86_64                                                                                    23/29 
  Verifying        : criu-3.6-1.fc27.x86_64                                                                                               24/29 
  Verifying        : libnet-1.1.6-14.fc27.x86_64                                                                                          25/29 
  Verifying        : protobuf-c-1.2.1-7.fc27.x86_64                                                                                       26/29 
  Verifying        : oci-register-machine-0-5.11.gitcd1e331.fc27.x86_64                                                                   27/29 
  Verifying        : oci-systemd-hook-1:0.1.13-1.gitafe4b4a.fc27.x86_64                                                                   28/29 
  Verifying        : systemd-container-234-8.fc27.x86_64                                                                                  29/29 

Installed:
  docker.x86_64 2:1.13.1-42.git4402c09.fc27                            criu.x86_64 3.6-1.fc27                                                  
  oci-register-machine.x86_64 0-5.11.gitcd1e331.fc27                   oci-systemd-hook.x86_64 1:0.1.13-1.gitafe4b4a.fc27                      
  atomic-registries.x86_64 1.20.1-6.fc27                               container-selinux.noarch 2:2.36-1.fc27                                  
  container-storage-setup.noarch 0.8.0-2.git1d27ecf.fc27               device-mapper-event.x86_64 1.02.142-4.fc27                              
  device-mapper-event-libs.x86_64 1.02.142-4.fc27                      device-mapper-persistent-data.x86_64 0.7.5-1.fc27                       
  docker-common.x86_64 2:1.13.1-42.git4402c09.fc27                     docker-rhel-push-plugin.x86_64 2:1.13.1-42.git4402c09.fc27              
  gnupg.x86_64 1.4.22-3.fc27                                           iptables.x86_64 1.6.1-4.fc27                                            
  libaio.x86_64 0.3.110-9.fc27                                         libnet.x86_64 1.1.6-14.fc27                                             
  libnetfilter_conntrack.x86_64 1.0.6-4.fc27                           libnfnetlink.x86_64 1.0.1-11.fc27                                       
  libusb.x86_64 1:0.1.5-10.fc27                                        lvm2.x86_64 2.02.173-4.fc27                                             
  lvm2-libs.x86_64 2.02.173-4.fc27                                     oci-umount.x86_64 2:2.3.0-1.git51e7c50.fc27                             
  policycoreutils-python-utils.x86_64 2.7-1.fc27                       protobuf-c.x86_64 1.2.1-7.fc27                                          
  python-rhsm-certificates.x86_64 1.20.1-3.fc27                        python3-pytoml.noarch 0.1.14-2.git7dea353.fc27                          
  skopeo-containers.x86_64 0.1.27-1.git93876ac.fc27                    systemd-container.x86_64 234-8.fc27                                     
  yajl.x86_64 2.1.0-8.fc27                                            

Complete!
```

```
[vagrant@localhost ~]$ sudo systemctl is-active docker.service
inactive
```

```
[vagrant@localhost ~]$ sudo systemctl start docker.service && sudo systemctl enable docker.service
Created symlink /etc/systemd/system/multi-user.target.wants/docker.service → /usr/lib/systemd/system/docker.service.
```

```
[vagrant@localhost ~]$ sudo systemctl -l status --no-pager docker.service
● docker.service - Docker Application Container Engine
   Loaded: loaded (/usr/lib/systemd/system/docker.service; enabled; vendor preset: disabled)
   Active: active (running) since Tue 2017-12-19 08:39:15 UTC; 31s ago
     Docs: http://docs.docker.com
 Main PID: 3249 (dockerd-current)
    Tasks: 7 (limit: 8192)
   CGroup: /system.slice/docker.service
           └─3249 /usr/bin/dockerd-current --add-runtime oci=/usr/libexec/docker/docker-runc-current --default-runtime=oci --authorization-plugin=rhel-push-plugin --containerd /run/containerd.sock --exec-opt native.cgroupdriver=systemd --userland-proxy-path=/usr/libexec/docker/docker-proxy-current --init-path=/usr/libexec/docker/docker-init-current --seccomp-profile=/etc/docker/seccomp.json --selinux-enabled --log-driver=journald --storage-driver overlay2 --add-registry docker.io --add-registry registry.fedoraproject.org --add-registry registry.access.redhat.com

Dec 19 08:39:15 localhost.localdomain dockerd-current[3249]: time="2017-12-19T08:39:15.087615636Z" level=warning msg="Your kernel does not support cgroup rt runtime"
Dec 19 08:39:15 localhost.localdomain dockerd-current[3249]: time="2017-12-19T08:39:15.088186152Z" level=info msg="Loading containers: start."
Dec 19 08:39:15 localhost.localdomain dockerd-current[3249]: time="2017-12-19T08:39:15.140612947Z" level=info msg="Firewalld running: false"
Dec 19 08:39:15 localhost.localdomain dockerd-current[3249]: time="2017-12-19T08:39:15.213379573Z" level=info msg="Default bridge (docker0) is assigned with an IP address 172.18.0.0/16. Daemon option --bip can be used to set a preferred IP address"
Dec 19 08:39:15 localhost.localdomain dockerd-current[3249]: time="2017-12-19T08:39:15.258648536Z" level=info msg="Loading containers: done."
Dec 19 08:39:15 localhost.localdomain dockerd-current[3249]: time="2017-12-19T08:39:15.272298368Z" level=warning msg="failed to retrieve docker-init version: unknown output format: tini version 0.16.1\n"
Dec 19 08:39:15 localhost.localdomain dockerd-current[3249]: time="2017-12-19T08:39:15.272741333Z" level=info msg="Daemon has completed initialization"
Dec 19 08:39:15 localhost.localdomain dockerd-current[3249]: time="2017-12-19T08:39:15.273031227Z" level=info msg="Docker daemon" commit=fbadd78-unsupported graphdriver=overlay2 version=1.13.1
Dec 19 08:39:15 localhost.localdomain systemd[1]: Started Docker Application Container Engine.
Dec 19 08:39:15 localhost.localdomain dockerd-current[3249]: time="2017-12-19T08:39:15.307576278Z" level=info msg="API listen on /var/run/docker.sock"
```

group
```
[vagrant@localhost ~]$ sudo groupadd docker
```

```
[vagrant@localhost ~]$ sudo usermod -aG docker vagrant
```

re-enter
```
[vagrant@localhost ~]$ exit
logout
Connection to 172.17.4.63 closed.
```

```
fanhonglingdeMacBook-Pro:rook0x2Frook fanhongling$ ssh vagrant@172.17.4.63
vagrant@172.17.4.63's password: 
Permission denied, please try again.
vagrant@172.17.4.63's password: 
Last failed login: Tue Dec 19 08:40:55 UTC 2017 from 172.17.4.1 on ssh:notty
There was 1 failed login attempt since the last successful login.
Last login: Tue Dec 19 08:15:37 2017 from 172.17.4.1
```

```
[vagrant@localhost ~]$ docker info
Containers: 0
 Running: 0
 Paused: 0
 Stopped: 0
Images: 0
Server Version: 1.13.1
Storage Driver: overlay2
 Backing Filesystem: extfs
 Supports d_type: true
 Native Overlay Diff: true
Logging Driver: journald
Cgroup Driver: systemd
Plugins: 
 Volume: local
 Network: bridge host macvlan null overlay
 Authorization: rhel-push-plugin
Swarm: inactive
Runtimes: oci runc
Default Runtime: oci
Init Binary: /usr/libexec/docker/docker-init-current
containerd version: fbadd789ddf86a4be9d6905528b7486c61e52612 (expected: aa8187dbd3b7ad67d8e5e3a15115d3eef43a7ed1)
runc version: fbadd789ddf86a4be9d6905528b7486c61e52612-dirty (expected: 9df8b306d01f59d3a8029be411de015b7304dd8f)
init version: N/A (expected: 949e6facb77383876aeff8a6944dde66b3089574)
Security Options:
 seccomp
  WARNING: You're not using the default seccomp profile
  Profile: /etc/docker/seccomp.json
 selinux
Kernel Version: 4.13.9-300.fc27.x86_64
Operating System: Fedora 27 (Cloud Edition)
OSType: linux
Architecture: x86_64
Number of Docker Hooks: 3
CPUs: 1
Total Memory: 1.951 GiB
Name: localhost.localdomain
ID: YISK:YAEI:BZDP:SZ7Y:SBXB:XMCT:75Q7:EITG:7LSQ:CIDF:TSPV:SPYR
Docker Root Dir: /var/lib/docker
Debug Mode (client): false
Debug Mode (server): false
Registry: https://index.docker.io/v1/
Experimental: false
Insecure Registries:
 127.0.0.0/8
Live Restore Enabled: false
Registries: docker.io (secure), registry.fedoraproject.org (secure), registry.access.redhat.com (secure), docker.io (secure)
```

dockerized file server to serve repository of packages
```
[vagrant@localhost ~]$ docker run -d --name=gofileserver -v /Users/fanhongling:/mnt -v /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com:/mnt/https0x3A0x2F0x2Fpackages.cloud.google.com -p 48080:48080 docker.io/tangfeixiong/gofileserver
Unable to find image 'docker.io/tangfeixiong/gofileserver:latest' locally
Trying to pull repository docker.io/tangfeixiong/gofileserver ... 
sha256:b05cb512410b314f3aae3ac1932faec2c8fc892221cc773778f4efdabd3684c8: Pulling from docker.io/tangfeixiong/gofileserver
c0cb142e4345: Pull complete 
bd94e5529d32: Pull complete 
Digest: sha256:b05cb512410b314f3aae3ac1932faec2c8fc892221cc773778f4efdabd3684c8
Status: Downloaded newer image for docker.io/tangfeixiong/gofileserver:latest
daa084733d2c2062587a8364bdedbe2e0c8d1eb64eb96c394e87356c09477304
```

```
[vagrant@localhost ~]$ curl -L http://localhost:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/
<pre>
<a href="apt/">apt/</a>
<a href="yum/">yum/</a>
</pre>
```

eth2, create like
```
[vagrant@localhost ~]$ sudo cat /etc/sysconfig/network-scripts/ifcfg-eth2
NM_CONTROLLED=no
BOOTPROTO=dhcp
#BOOTPROTO=none
ONBOOT=yes
#IPADDR=172.17.4.63
#NETMASK=255.255.255.0
DEVICE=eth2
PEERDNS=no
```

```
[vagrant@localhost ~]$ ip addr show eth2
4: eth2: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 08:00:27:69:79:6f brd ff:ff:ff:ff:ff:ff
```

hostname
```
[vagrant@localhost ~]$ hostname -I
10.0.2.15 172.17.4.63 172.18.0.1 
```

```
[vagrant@localhost ~]$ pos=2; addr=$(hostname -I | awk -v i=$pos '{print $i}') && name="kube-${addr//./-}" && echo $name
kube-172-17-4-63
```

```
[vagrant@localhost ~]$ sudo hostnamectl set-hostname $name
```

```
[vagrant@localhost ~]$ sudo hostnamectl                   
   Static hostname: kube-172-17-4-63
         Icon name: computer-vm
           Chassis: vm
        Machine ID: 70bfdc3adb5b4b8cb5b9aab803c827aa
           Boot ID: f78df90e043547a9a2cacde420d5ba82
    Virtualization: oracle
  Operating System: Fedora 27 (Cloud Edition)
       CPE OS Name: cpe:/o:fedoraproject:fedora:27
            Kernel: Linux 4.13.9-300.fc27.x86_64
      Architecture: x86-64
```

hosts
```
[vagrant@localhost ~]$ echo "$addr    $name" | sudo tee -a /etc/hosts
172.17.4.63    kube-172-17-4-63
```

```
[vagrant@localhost ~]$ cat /etc/hosts
127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
::1         localhost localhost.localdomain localhost6 localhost6.localdomain6

172.17.4.63    kube-172-17-4-63
```

Selinux
```
[vagrant@localhost ~]$ getenforce
Enforcing
```

```
[vagrant@localhost ~]$ sudo setenforce 0
```

```
[vagrant@localhost ~]$ getenforce
Permissive
```

```
[vagrant@localhost ~]$ sudo sed -i 's/SELINUX=enforcing/SELINUX=permissive/' /etc/sysconfig/selinux 
```

```
[vagrant@localhost ~]$ sudo sed -i 's/SELINUX=enforcing/SELINUX=permissive/' /etc/selinux/config    
```

Firewall
```
[vagrant@localhost ~]$ sudo systemctl is-active firewalld.service
inactive
```

## Images

### kubernetes

ssh v1.9.0 cluster master
```
[vagrant@localhost ~]$ ssh vagrant@172.17.4.59 "ls /opt/kubernetes/server/bin"
The authenticity of host '172.17.4.59 (172.17.4.59)' can't be established.
ECDSA key fingerprint is SHA256:VV533NeJHFykYUa3CBNDm0zDI4HSnEaMKcQSKwtvLGw.
ECDSA key fingerprint is MD5:e3:32:89:1f:f1:f3:a8:b1:85:99:dc:f4:5f:24:04:67.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.17.4.59' (ECDSA) to the list of known hosts.
vagrant@172.17.4.59's password: 
apiextensions-apiserver
cloud-controller-manager
cloud-controller-manager.docker_tag
cloud-controller-manager.tar
hyperkube
kube-aggregator
kube-aggregator.docker_tag
kube-aggregator.tar
kube-apiserver
kube-apiserver.docker_tag
kube-apiserver.tar
kube-controller-manager
kube-controller-manager.docker_tag
kube-controller-manager.tar
kube-proxy
kube-proxy.docker_tag
kube-proxy.tar
kube-scheduler
kube-scheduler.docker_tag
kube-scheduler.tar
kubeadm
kubectl
kubelet
mounter
```

```
[vagrant@localhost ~]$ ssh vagrant@172.17.4.59 "cat /opt/kubernetes/server/bin/kube-apiserver.tar" | docker load
vagrant@172.17.4.59's password: 
0271b8eebde3: Loading layer [==================================================>] 1.338 MB/1.338 MB
9ccc9fba4253: Loading layer [==================================================>] 209.2 MB/209.2 MB
Loaded image: gcr.io/google_containers/kube-apiserver:v1.9.0
```

```
[vagrant@localhost ~]$ ssh vagrant@172.17.4.59 "cat /opt/kubernetes/server/bin/kube-controller-manager.tar" | docker load
vagrant@172.17.4.59's password: 
50a426d115f8: Loading layer [==================================================>] 136.6 MB/136.6 MB
Loaded image: gcr.io/google_containers/kube-controller-manager:v1.9.0
```

```
[vagrant@localhost ~]$ ssh vagrant@172.17.4.59 "cat /opt/kubernetes/server/bin/kube-scheduler.tar" | docker load
vagrant@172.17.4.59's password: 
f733b8f8af29: Loading layer [==================================================>] 61.57 MB/61.57 MB
Loaded image: gcr.io/google_containers/kube-scheduler:v1.9.0
```

```
[vagrant@localhost ~]$ ssh vagrant@172.17.4.59 "cat /opt/kubernetes/server/bin/kube-proxy.tar" | docker load
vagrant@172.17.4.59's password: 
684c19bf2c27: Loading layer [==================================================>]  44.2 MB/44.2 MB
deb4ca39ea31: Loading layer [==================================================>] 3.358 MB/3.358 MB
9c44b0d51ed1: Loading layer [==================================================>] 63.38 MB/63.38 MB
Loaded image: gcr.io/google_containers/kube-proxy:v1.9.0
```

tag 
```
[vagrant@localhost ~]$ for name in "kube-apiserver" "kube-controller-manager" "kube-scheduler" "kube-proxy"; do docker tag "gcr.io/google_containers/$name:v1.9.0" "gcr.io/google_containers/$name-amd64:v1.9.0"; done
```

### POD base, etcd, kube dns, cni flannel

ssh v1.9.0 cluster maaster
```
[vagrant@localhost ~]$ ssh vagrant@172.17.4.59 "docker images"
vagrant@172.17.4.59's password: 
REPOSITORY                                               TAG                 IMAGE ID            CREATED             SIZE
docker.io/rook/toolbox                                   master              78b8570f3f57        3 days ago          415 MB
docker.io/rook/rook                                      master              8d9558dd80c0        3 days ago          340 MB
gcr.io/google_containers/kube-proxy                      v1.9.0              f6f363e6e98e        3 days ago          109 MB
gcr.io/google_containers/kube-proxy-amd64                v1.9.0              f6f363e6e98e        3 days ago          109 MB
gcr.io/google_containers/cloud-controller-manager        v1.9.0              4c938e6fc795        3 days ago          118 MB
gcr.io/google_containers/kube-apiserver                  v1.9.0              7bff5aa286d7        3 days ago          210 MB
gcr.io/google_containers/kube-apiserver-amd64            v1.9.0              7bff5aa286d7        3 days ago          210 MB
gcr.io/google_containers/kube-aggregator                 v1.9.0              e08db9577664        3 days ago          55.8 MB
gcr.io/google_containers/kube-controller-manager-amd64   v1.9.0              3bb172f9452c        3 days ago          138 MB
gcr.io/google_containers/kube-controller-manager         v1.9.0              3bb172f9452c        3 days ago          138 MB
gcr.io/google_containers/kube-scheduler-amd64            v1.9.0              5ceb21996307        3 days ago          62.7 MB
gcr.io/google_containers/kube-scheduler                  v1.9.0              5ceb21996307        3 days ago          62.7 MB
gcr.io/google_containers/kubernetes-dashboard-amd64      v1.8.0              55dbc28356f2        3 weeks ago         119 MB
quay.io/coreos/flannel                                   v0.9.1-amd64        2b736d06ca4c        4 weeks ago         51.3 MB
docker.io/coredns/coredns                                0.9.10              409b1e046465        6 weeks ago         44.7 MB
gcr.io/google_containers/k8s-dns-sidecar-amd64           1.14.7              db76ee297b85        8 weeks ago         42 MB
gcr.io/google_containers/k8s-dns-kube-dns-amd64          1.14.7              5d049a8c4eec        8 weeks ago         50.3 MB
gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64     1.14.7              5feec37454f4        8 weeks ago         41 MB
gcr.io/google_containers/etcd-amd64                      3.1.10              1406502a6459        3 months ago        193 MB
gcr.io/google_containers/etcd                            3.1.10              1406502a6459        3 months ago        193 MB
docker.io/tangfeixiong/gofileserver                      latest              f07338e49481        6 months ago        12.6 MB
docker.io/tangfeixiong/netcat-hello-http                 latest              29c91b3bcc05        14 months ago       12.1 MB
gcr.io/google_containers/pause-amd64                     3.0                 99e59f495ffa        19 months ago       747 kB
```

```
[vagrant@localhost ~]$ for name in "gcr.io/google_containers/pause-amd64:3.0" "gcr.io/google_containers/etcd-amd64:3.1.10" "gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.7" "gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.7" "gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.7" "docker.io/coredns/coredns:0.9.10" "quay.io/coreos/flannel:v0.9.1-amd64"; do ssh vagrant@172.17.4.59 "docker save $name" | docker load; done
vagrant@172.17.4.59's password: 
5f70bf18a086: Loading layer [==================================================>] 1.024 kB/1.024 kB
41ff149e94f2: Loading layer [==================================================>] 748.5 kB/748.5 kB
Loaded image: gcr.io/google_containers/pause-amd64:3.0
vagrant@172.17.4.59's password: 
6a749002dd6a: Loading layer [==================================================>] 1.338 MB/1.338 MB
bbd07ea14872: Loading layer [==================================================>] 159.2 MB/159.2 MB
611a3394df5d: Loading layer [==================================================>] 32.44 MB/32.44 MB
Loaded image: gcr.io/google_containers/etcd-amd64:3.1.10
vagrant@172.17.4.59's password: 
5bef08742407: Loading layer [==================================================>] 4.221 MB/4.221 MB
bd94706d2c63: Loading layer [==================================================>] 38.07 MB/38.07 MB
Loaded image: gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.7
vagrant@172.17.4.59's password: 
cd69fdcd7591: Loading layer [==================================================>] 46.31 MB/46.31 MB
Loaded image: gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.7
vagrant@172.17.4.59's password: 
b87261cc1ccb: Loading layer [==================================================>]  2.56 kB/2.56 kB
ac66a5c581a8: Loading layer [==================================================>]   362 kB/362 kB
22f71f461ac8: Loading layer [==================================================>] 3.072 kB/3.072 kB
686a085da152: Loading layer [==================================================>] 36.63 MB/36.63 MB
Loaded image: gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.7
vagrant@172.17.4.59's password: 
034ea01e4aa0: Loading layer [==================================================>] 9.324 MB/9.324 MB
1ea1b16efe91: Loading layer [==================================================>] 31.82 MB/31.82 MB
Loaded image: docker.io/coredns/coredns:0.9.10
vagrant@172.17.4.59's password: 
f439636ab0f0: Loading layer [==================================================>] 6.797 MB/6.797 MB
91b6f6ead101: Loading layer [==================================================>] 4.414 MB/4.414 MB
fc3c053505e6: Loading layer [==================================================>] 34.49 MB/34.49 MB
032657ac7c4a: Loading layer [==================================================>] 2.225 MB/2.225 MB
fd713c7c81af: Loading layer [==================================================>]  5.12 kB/5.12 kB
Loaded image: quay.io/coreos/flannel:v0.9.1-amd64
```

## Instrall kubeadm

repo
```
[vagrant@localhost ~]$ curl -L http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/etc0x2Fyum.repos.d/kubernetes.repo | sudo tee /etc/yum.repos.d/kubernetes.repo
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   853  100   853    0     0    853      0  0:00:01 --:--:--  0:00:01 37086
[kubernetes]
name=Kubernetes
#baseurl=http://yum.kubernetes.io/repos/kubernetes-el7-x86_64
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
        https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
		
### Modify URL first
[kubernetes-local-http]
name=Kubernetes offline - file server
baseurl=http://127.0.0.1:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/
enabled=0
gpgcheck=0

### Modify path first
[kubernetes-local-file]
name=Kubernetes offline - downloads
baseurl=file:///Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/
enabled=0
gpgcheck=0

```

gpg 
```
[vagrant@localhost ~]$ sudo rpmkeys --import http://127.0.0.1:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/doc/yum-key.gpg
```

__issue__
```
[vagrant@localhost ~]$ sudo dnf --disablerepo=kubernetes --enablerepo=kubernetes-local-http repolist
Failed to set locale, defaulting to C
Failed to synchronize cache for repo 'kubernetes-local-http', disabling.
Last metadata expiration check: 1:29:37 ago on Tue Dec 19 08:07:22 2017.
repo id                                                    repo name                                                                      status
*fedora                                                    Fedora 27 - x86_64                                                             54801
*updates                                                   Fedora 27 - x86_64 - Updates                                                    7824
```

Refer to https://bugzilla.redhat.com/show_bug.cgi?id=1494178

```
[vagrant@localhost ~]$ sudo dnf search -y kubeadm kubelet kubectl
Failed to set locale, defaulting to C
Importing GPG key 0xA7317B0F:
 Userid     : "Google Cloud Packages Automatic Signing Key <gc-team@google.com>"
 Fingerprint: D0BC 747F D8CA F711 7500 D6FA 3746 C208 A731 7B0F
 From       : https://packages.cloud.google.com/yum/doc/yum-key.gpg
Importing GPG key 0x3E1BA8D5:
 Userid     : "Google Cloud Packages RPM Signing Key <gc-team@google.com>"
 Fingerprint: 3749 E1BA 95A8 6CE0 5454 6ED2 F09C 394C 3E1B A8D5
 From       : https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
Kubernetes                                                                                                       41 kB/s |  29 kB     00:00    
Last metadata expiration check: 0:00:00 ago on Tue Dec 19 09:45:16 2017.
======================================================== Name Exactly Matched: kubelet =========================================================
kubelet.x86_64 : Container cluster management
======================================================== Name Exactly Matched: kubectl =========================================================
kubectl.x86_64 : Command-line utility for interacting with a Kubernetes cluster.
======================================================== Name Exactly Matched: kubeadm =========================================================
kubeadm.x86_64 : Command-line utility for administering a Kubernetes cluster.
============================================================ Name Matched: kubeadm =============================================================
kubernetes-kubeadm.x86_64 : Kubernetes tool for standing up clusters
```

```
[vagrant@localhost ~]$ sudo dnf install -y kubeadm
Failed to set locale, defaulting to C
Last metadata expiration check: 0:00:56 ago on Tue Dec 19 09:45:16 2017.
Dependencies resolved.
================================================================================================================================================
 Package                              Arch                         Version                               Repository                        Size
================================================================================================================================================
Installing:
 kubeadm                              x86_64                       1.9.0-0                               kubernetes                        16 M
Installing dependencies:
 ebtables                             x86_64                       2.0.10-24.fc27                        fedora                           134 k
 ethtool                              x86_64                       2:4.13-1.fc27                         updates                          138 k
 kubectl                              x86_64                       1.9.0-0                               kubernetes                       8.9 M
 kubelet                              x86_64                       1.9.0-0                               kubernetes                        17 M
 kubernetes-cni                       x86_64                       0.6.0-0                               kubernetes                       8.6 M
 socat                                x86_64                       1.7.3.2-4.fc27                        fedora                           296 k

Transaction Summary
================================================================================================================================================
Install  7 Packages

Total download size: 51 M
Installed size: 275 M
Downloading Packages:
(1/7): bc390a3d43256791bfb844696e7215fd7ad8a09f70a42667dab4997415a6ba75-kubectl-1.9.0-0.x86_64.rpm              888 kB/s | 8.9 MB     00:10    
(2/7): ebtables-2.0.10-24.fc27.x86_64.rpm                                                                       216 kB/s | 134 kB     00:00    
(3/7): socat-1.7.3.2-4.fc27.x86_64.rpm                                                                          415 kB/s | 296 kB     00:00    
(4/7): fe33057ffe95bfae65e2f269e1b05e99308853176e24a4d027bc082b471a07c0-kubernetes-cni-0.6.0-0.x86_64.rpm       687 kB/s | 8.6 MB     00:12    
(5/7): ethtool-4.13-1.fc27.x86_64.rpm                                                                           235 kB/s | 138 kB     00:00    
(6/7): aa9948f82e7af317c97a242f3890985159c09c183b46ac8aab19d2ad307e6970-kubeadm-1.9.0-0.x86_64.rpm              585 kB/s |  16 MB     00:28    
(7/7): 8f507de9e1cc26e5b0043e334e26d62001c171d8e54d839128e9bade25ecda95-kubelet-1.9.0-0.x86_64.rpm              475 kB/s |  17 MB     00:36    
------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                           1.3 MB/s |  51 MB     00:39     
warning: /var/cache/dnf/kubernetes-33343725abd9cbdc/packages/aa9948f82e7af317c97a242f3890985159c09c183b46ac8aab19d2ad307e6970-kubeadm-1.9.0-0.x86_64.rpm: Header V4 RSA/SHA1 Signature, key ID 3e1ba8d5: NOKEY
Importing GPG key 0xA7317B0F:
 Userid     : "Google Cloud Packages Automatic Signing Key <gc-team@google.com>"
 Fingerprint: D0BC 747F D8CA F711 7500 D6FA 3746 C208 A731 7B0F
 From       : https://packages.cloud.google.com/yum/doc/yum-key.gpg
Key imported successfully
Importing GPG key 0x3E1BA8D5:
 Userid     : "Google Cloud Packages RPM Signing Key <gc-team@google.com>"
 Fingerprint: 3749 E1BA 95A8 6CE0 5454 6ED2 F09C 394C 3E1B A8D5
 From       : https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
Key imported successfully
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                        1/1 
  Installing       : ethtool-2:4.13-1.fc27.x86_64                                                                                           1/7 
  Installing       : socat-1.7.3.2-4.fc27.x86_64                                                                                            2/7 
  Installing       : ebtables-2.0.10-24.fc27.x86_64                                                                                         3/7 
  Running scriptlet: ebtables-2.0.10-24.fc27.x86_64                                                                                         3/7 
  Installing       : kubernetes-cni-0.6.0-0.x86_64                                                                                          4/7 
  Installing       : kubelet-1.9.0-0.x86_64                                                                                                 5/7 
  Installing       : kubectl-1.9.0-0.x86_64                                                                                                 6/7 
  Installing       : kubeadm-1.9.0-0.x86_64                                                                                                 7/7 
  Running scriptlet: kubeadm-1.9.0-0.x86_64                                                                                                 7/7 
  Verifying        : kubeadm-1.9.0-0.x86_64                                                                                                 1/7 
  Verifying        : kubectl-1.9.0-0.x86_64                                                                                                 2/7 
  Verifying        : kubelet-1.9.0-0.x86_64                                                                                                 3/7 
  Verifying        : ebtables-2.0.10-24.fc27.x86_64                                                                                         4/7 
  Verifying        : socat-1.7.3.2-4.fc27.x86_64                                                                                            5/7 
  Verifying        : kubernetes-cni-0.6.0-0.x86_64                                                                                          6/7 
  Verifying        : ethtool-2:4.13-1.fc27.x86_64                                                                                           7/7 

Installed:
  kubeadm.x86_64 1.9.0-0         ebtables.x86_64 2.0.10-24.fc27  ethtool.x86_64 2:4.13-1.fc27  kubectl.x86_64 1.9.0-0  kubelet.x86_64 1.9.0-0 
  kubernetes-cni.x86_64 0.6.0-0  socat.x86_64 1.7.3.2-4.fc27    

Complete!
```

service
```
[vagrant@localhost ~]$ cat /etc/systemd/system/kubelet.service
[Unit]
Description=kubelet: The Kubernetes Node Agent
Documentation=http://kubernetes.io/docs/

[Service]
ExecStart=/usr/bin/kubelet
Restart=always
StartLimitInterval=0
RestartSec=10

[Install]
WantedBy=multi-user.target
```

```
[vagrant@localhost ~]$ cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"
Environment="KUBELET_CERTIFICATE_ARGS=--rotate-certificates=true --cert-dir=/var/lib/kubelet/pki"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBEL
```

conf
```
[vagrant@localhost ~]$ ls -R /etc/kubernetes/
/etc/kubernetes/:
manifests

/etc/kubernetes/manifests:
```

```
[vagrant@localhost ~]$ ls /opt/cni/bin/
bridge  dhcp  flannel  host-local  ipvlan  loopback  macvlan  portmap  ptp  sample  tuning  vlan
```

### install 

copy admin.conf from v1.9.0 cluster master
```
[vagrant@localhost ~]$ mkdir -p .kube && scp vagrant@172.17.4.59:/home/vagrant/.kube/config .kube
vagrant@172.17.4.59's password: 
config                                                                                                        100% 5451     4.9MB/s   00:00    
```

```
[vagrant@localhost ~]$ kubectl version
Client Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T21:07:38Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T20:55:30Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
```

token
```
[vagrant@localhost ~]$ sudo kubeadm token list --kubeconfig=.kube/config    
TOKEN     TTL       EXPIRES   USAGES    DESCRIPTION   EXTRA GROUPS
```

```
[vagrant@localhost ~]$ sudo kubeadm token generate --kubeconfig=.kube/config
86257a.d063ca0eedd9e165
```

sysctl
```
[vagrant@localhost ~]$ sudo sysctl net.bridge.bridge-nf-call-iptables
net.bridge.bridge-nf-call-iptables = 1
```

```
[vagrant@localhost ~]$ grep -R "bridge-nf-call-iptables" /usr/lib/sysctl.d/ /etc/sysctl.conf /etc/sysctl.d/   
```

join
```
[vagrant@localhost ~]$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 
[preflight] Running pre-flight checks.
	[WARNING Service-Kubelet]: kubelet service is not enabled, please run 'systemctl enable kubelet.service'
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
[preflight] Starting the kubelet service
[discovery: Required value: DiscoveryTokenAPIServers not set, discovery: Invalid value: "": using token-based discovery without DiscoveryTokenCACertHashes can be unsafe. set --discovery-token-unsafe-skip-ca-verification to continue]
```

```
[vagrant@localhost ~]$ sudo systemctl enable kubelet.service; sudo systemctl start kubelet.service
Created symlink /etc/systemd/system/multi-user.target.wants/kubelet.service → /etc/systemd/system/kubelet.service.
```

```
[vagrant@localhost ~]$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 --discovery-token-unsafe-skip-ca-verification
[preflight] Running pre-flight checks.
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
discovery: Required value: DiscoveryTokenAPIServers not set
```

```
[vagrant@localhost ~]$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 172.17.4.59:6443 --discovery-token-unsafe-skip-ca-verification
[preflight] Running pre-flight checks.
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
[discovery] Trying to connect to API Server "172.17.4.59:6443"
[discovery] Created cluster-info discovery client, requesting info from "https://172.17.4.59:6443"
[discovery] Failed to connect to API Server "172.17.4.59:6443": there is no JWS signed token in the cluster-info ConfigMap. This token id "86257a" is invalid for this cluster, can't connect
[discovery] Trying to connect to API Server "172.17.4.59:6443"
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo kubeadm token create 86257a.d063ca0eedd9e165
86257a.d063ca0eedd9e165
```

```
[vagrant@localhost ~]$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 172.17.4.59:6443 --discovery-token-unsafe-skip-ca-verification
[preflight] Running pre-flight checks.
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
[discovery] Trying to connect to API Server "172.17.4.59:6443"
[discovery] Created cluster-info discovery client, requesting info from "https://172.17.4.59:6443"
[discovery] Cluster info signature and contents are valid and no TLS pinning was specified, will use API Server "172.17.4.59:6443"
[discovery] Successfully established connection with API Server "172.17.4.59:6443"

This node has joined the cluster:
* Certificate signing request was sent to master and a response
  was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the master to see this node join the cluster.
```

```[vagrant@localhost ~]$ kubectl get nodes
NAME                  STATUS    ROLES     AGE       VERSION
kube-172-17-4-63      Ready     <none>    19s       v1.9.0
kubedev-172-17-4-55   Ready     <none>    8h        v1.9.0
kubedev-172-17-4-59   Ready     master    1d        v1.9.0
```

```
[vagrant@localhost ~]$ kubectl -n kube-system get pods -l app=flannel
NAME                    READY     STATUS    RESTARTS   AGE
kube-flannel-ds-7hb8k   1/1       Running   0          1m
kube-flannel-ds-7pvdh   1/1       Running   0          7h
kube-flannel-ds-x2l9q   1/1       Running   0          7h
```

```
[vagrant@localhost ~]$ ip a show cni0
9: cni0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 0a:58:0a:f4:02:01 brd ff:ff:ff:ff:ff:ff
    inet 10.244.2.1/24 scope global cni0
       valid_lft forever preferred_lft forever
    inet6 fe80::30d0:b3ff:fe05:8000/64 scope link 
       valid_lft forever preferred_lft forever
```

__issue__
```
Dec 19 18:01:36 rookdev-172-17-4-63 kubelet[2054]: E1219 18:01:36.909808    2054 kuberuntime_manager.go:647] createPodSandbox for pod "rook-ceph-osd-dl7f5_rook(8cd53cee-e4e5-11e7-bd70-525400224e72)" failed: rpc error: code = Unknown desc = NetworkPlugin cni failed to set up pod "rook-ceph-osd-dl7f5_rook" network: failed to allocate for range 0: no IP addresses available in range set: 10.244.3.1-10.244.3.254
```

```
[vagrant@rookdev-172-17-4-63 ~]$ sudo ls /var/lib/cni/networks/cbr0
10.244.3.10   10.244.3.123  10.244.3.147  10.244.3.170	10.244.3.194  10.244.3.217  10.244.3.240  10.244.3.35  10.244.3.59  10.244.3.82
10.244.3.100  10.244.3.124  10.244.3.148  10.244.3.171	10.244.3.195  10.244.3.218  10.244.3.241  10.244.3.36  10.244.3.6   10.244.3.83
10.244.3.101  10.244.3.125  10.244.3.149  10.244.3.172	10.244.3.196  10.244.3.219  10.244.3.242  10.244.3.37  10.244.3.60  10.244.3.84
10.244.3.102  10.244.3.126  10.244.3.15   10.244.3.173	10.244.3.197  10.244.3.22   10.244.3.243  10.244.3.38  10.244.3.61  10.244.3.85
10.244.3.103  10.244.3.127  10.244.3.150  10.244.3.174	10.244.3.198  10.244.3.220  10.244.3.244  10.244.3.39  10.244.3.62  10.244.3.86
10.244.3.104  10.244.3.128  10.244.3.151  10.244.3.175	10.244.3.199  10.244.3.221  10.244.3.245  10.244.3.4   10.244.3.63  10.244.3.87
10.244.3.105  10.244.3.129  10.244.3.152  10.244.3.176	10.244.3.2    10.244.3.222  10.244.3.246  10.244.3.40  10.244.3.64  10.244.3.88
10.244.3.106  10.244.3.13   10.244.3.153  10.244.3.177	10.244.3.20   10.244.3.223  10.244.3.247  10.244.3.41  10.244.3.65  10.244.3.89
10.244.3.107  10.244.3.130  10.244.3.154  10.244.3.178	10.244.3.200  10.244.3.224  10.244.3.248  10.244.3.42  10.244.3.66  10.244.3.9
10.244.3.108  10.244.3.131  10.244.3.155  10.244.3.179	10.244.3.201  10.244.3.225  10.244.3.249  10.244.3.43  10.244.3.67  10.244.3.90
10.244.3.109  10.244.3.132  10.244.3.156  10.244.3.18	10.244.3.202  10.244.3.226  10.244.3.25   10.244.3.44  10.244.3.68  10.244.3.91
10.244.3.11   10.244.3.133  10.244.3.157  10.244.3.180	10.244.3.203  10.244.3.227  10.244.3.250  10.244.3.45  10.244.3.69  10.244.3.92
10.244.3.110  10.244.3.134  10.244.3.158  10.244.3.181	10.244.3.204  10.244.3.228  10.244.3.251  10.244.3.46  10.244.3.7   10.244.3.93
10.244.3.111  10.244.3.135  10.244.3.159  10.244.3.182	10.244.3.205  10.244.3.229  10.244.3.252  10.244.3.47  10.244.3.70  10.244.3.94
10.244.3.112  10.244.3.136  10.244.3.16   10.244.3.183	10.244.3.206  10.244.3.23   10.244.3.253  10.244.3.48  10.244.3.71  10.244.3.95
10.244.3.113  10.244.3.137  10.244.3.160  10.244.3.184	10.244.3.207  10.244.3.230  10.244.3.254  10.244.3.49  10.244.3.72  10.244.3.96
10.244.3.114  10.244.3.138  10.244.3.161  10.244.3.185	10.244.3.208  10.244.3.231  10.244.3.26   10.244.3.5   10.244.3.73  10.244.3.97
10.244.3.115  10.244.3.139  10.244.3.162  10.244.3.186	10.244.3.209  10.244.3.232  10.244.3.27   10.244.3.50  10.244.3.74  10.244.3.98
10.244.3.116  10.244.3.14   10.244.3.163  10.244.3.187	10.244.3.21   10.244.3.233  10.244.3.28   10.244.3.51  10.244.3.75  10.244.3.99
10.244.3.117  10.244.3.140  10.244.3.164  10.244.3.188	10.244.3.210  10.244.3.234  10.244.3.29   10.244.3.52  10.244.3.76  last_reserved_ip.0
10.244.3.118  10.244.3.141  10.244.3.165  10.244.3.189	10.244.3.211  10.244.3.235  10.244.3.3	  10.244.3.53  10.244.3.77
10.244.3.119  10.244.3.142  10.244.3.166  10.244.3.19	10.244.3.212  10.244.3.236  10.244.3.30   10.244.3.54  10.244.3.78
10.244.3.12   10.244.3.143  10.244.3.167  10.244.3.190	10.244.3.213  10.244.3.237  10.244.3.31   10.244.3.55  10.244.3.79
10.244.3.120  10.244.3.144  10.244.3.168  10.244.3.191	10.244.3.214  10.244.3.238  10.244.3.32   10.244.3.56  10.244.3.8
10.244.3.121  10.244.3.145  10.244.3.169  10.244.3.192	10.244.3.215  10.244.3.239  10.244.3.33   10.244.3.57  10.244.3.80
10.244.3.122  10.244.3.146  10.244.3.17   10.244.3.193	10.244.3.216  10.244.3.24   10.244.3.34   10.244.3.58  10.244.3.81
```

```
[vagrant@rookdev-172-17-4-63 ~]$ ip a show cni0
7: cni0: <BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 00:00:00:00:00:00 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::a85c:7ff:fecc:6491/64 scope link 
       valid_lft forever preferred_lft forever
```

refer to https://github.com/kubernetes/kubernetes/issues/34278#issuecomment-254686727

```
[vagrant@rookdev-172-17-4-63 ~]$ sudo ls /var/lib/cni/networks/cbr0 | wc -l
1
```

```
[vagrant@rookdev-172-17-4-63 ~]$ sudo ls /var/lib/cni/flannel | wc -l
0
```


change node name
```
[vagrant@localhost ~]$ pos=2; addr=$(hostname -I | awk -v i=$pos '{print $i}') && name="rookdev-${addr//./-}" && echo $name
rookdev-172-17-4-63
```

```
[vagrant@localhost ~]$ sudo hostnamectl set-hostname $name && hostname
rookdev-172-17-4-63
```

```
[vagrant@localhost ~]$ sudo sed -i "s/^$addr.*$/$addr    $name/" /etc/hosts
```

```
[vagrant@localhost ~]$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 172.17.4.59:6443 --discovery-token-unsafe-skip-ca-verification
[preflight] Running pre-flight checks.
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
[preflight] Some fatal errors occurred:
	[ERROR Port-10250]: Port 10250 is in use
	[ERROR FileAvailable--etc-kubernetes-pki-ca.crt]: /etc/kubernetes/pki/ca.crt already exists
	[ERROR FileAvailable--etc-kubernetes-kubelet.conf]: /etc/kubernetes/kubelet.conf already exists
[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
[vagrant@localhost ~]$ sudo kubeadm reset
[preflight] Running pre-flight checks.
[reset] Stopping the kubelet service.
[reset] Unmounting mounted directories in "/var/lib/kubelet"
[reset] Removing kubernetes-managed containers.
[reset] No etcd manifest found in "/etc/kubernetes/manifests/etcd.yaml". Assuming external etcd.
[reset] Deleting contents of stateful directories: [/var/lib/kubelet /etc/cni/net.d /var/lib/dockershim /var/run/kubernetes]
[reset] Deleting contents of config directories: [/etc/kubernetes/manifests /etc/kubernetes/pki]
[reset] Deleting files: [/etc/kubernetes/admin.conf /etc/kubernetes/kubelet.conf /etc/kubernetes/controller-manager.conf /etc/kubernetes/scheduler.conf]
```

```
[vagrant@localhost ~]$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 172.17.4.59:6443 --discovery-token-unsafe-skip-ca-verification
[preflight] Running pre-flight checks.
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
[preflight] Starting the kubelet service
[discovery] Trying to connect to API Server "172.17.4.59:6443"
[discovery] Created cluster-info discovery client, requesting info from "https://172.17.4.59:6443"
[discovery] Cluster info signature and contents are valid and no TLS pinning was specified, will use API Server "172.17.4.59:6443"
[discovery] Successfully established connection with API Server "172.17.4.59:6443"

This node has joined the cluster:
* Certificate signing request was sent to master and a response
  was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the master to see this node join the cluster.
```

```
NAME                  STATUS    ROLES     AGE       VERSION
kubedev-172-17-4-55   Ready     <none>    8h        v1.9.0
kubedev-172-17-4-59   Ready     master    1d        v1.9.0
rookdev-172-17-4-63   Ready     <none>    26s       v1.9.0
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl -n kube-system get pods -l app=flannel
NAME                    READY     STATUS    RESTARTS   AGE
kube-flannel-ds-7pvdh   1/1       Running   0          7h
kube-flannel-ds-89hwq   1/1       Running   0          58s
kube-flannel-ds-x2l9q   1/1       Running   0          7h
```