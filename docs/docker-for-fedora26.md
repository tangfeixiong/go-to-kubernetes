# instruction

docker
```
[vagrant@localhost ~]$ sudo dnf install -y docker
Failed to set locale, defaulting to C
Last metadata expiration check: 1:47:00 ago on Sun Dec 17 13:53:12 2017.
Dependencies resolved.
================================================================================================================================================
 Package                                       Arch                   Version                                     Repository               Size
================================================================================================================================================
Installing:
 docker                                        x86_64                 2:1.13.1-44.git584d391.fc26                 updates                  20 M
Installing dependencies:
 atomic-registries                             x86_64                 1.20.1-6.fc26                               updates                  38 k
 container-selinux                             noarch                 2:2.36-1.fc26                               updates                  35 k
 container-storage-setup                       noarch                 0.8.0-2.git1d27ecf.fc26                     updates                  36 k
 device-mapper-event                           x86_64                 1.02.137-6.fc26                             fedora                  241 k
 device-mapper-event-libs                      x86_64                 1.02.137-6.fc26                             fedora                  242 k
 device-mapper-persistent-data                 x86_64                 0.6.3-5.fc26                                fedora                  434 k
 docker-common                                 x86_64                 2:1.13.1-44.git584d391.fc26                 updates                  83 k
 docker-rhel-push-plugin                       x86_64                 2:1.13.1-44.git584d391.fc26                 updates                 1.6 M
 gnupg                                         x86_64                 1.4.22-1.fc26                               updates                 1.3 M
 iptables                                      x86_64                 1.6.1-2.fc26                                fedora                  434 k
 libaio                                        x86_64                 0.3.110-7.fc26                              fedora                   28 k
 libnet                                        x86_64                 1.1.6-12.fc26                               fedora                   63 k
 libnetfilter_conntrack                        x86_64                 1.0.6-2.fc26                                fedora                   61 k
 libnfnetlink                                  x86_64                 1.0.1-9.fc26                                fedora                   30 k
 libusb                                        x86_64                 1:0.1.5-8.fc26                              fedora                   39 k
 lvm2                                          x86_64                 2.02.168-6.fc26                             fedora                  1.2 M
 lvm2-libs                                     x86_64                 2.02.168-6.fc26                             fedora                  1.0 M
 oci-umount                                    x86_64                 2:2.3.0-1.git51e7c50.fc26                   updates                  34 k
 policycoreutils-python-utils                  x86_64                 2.6-5.fc26                                  fedora                  218 k
 protobuf-c                                    x86_64                 1.2.1-4.fc26                                fedora                   34 k
 python-rhsm-certificates                      x86_64                 1.20.2-1.fc26                               updates                  44 k
 python3-pytoml                                noarch                 0.1.14-1.git7dea353.fc26                    updates                  23 k
 skopeo-containers                             x86_64                 0.1.27-1.git93876ac.fc26                    updates                  16 k
 systemd-container                             x86_64                 233-6.fc26                                  fedora                  416 k
 yajl                                          x86_64                 2.1.0-6.fc26                                fedora                   36 k
Installing weak dependencies:
 criu                                          x86_64                 3.6-1.fc26                                  updates                 455 k
 oci-register-machine                          x86_64                 0-5.11.gitcd1e331.fc26                      updates                 1.0 M
 oci-systemd-hook                              x86_64                 1:0.1.13-1.gitafe4b4a.fc26                  updates                  35 k

Transaction Summary
================================================================================================================================================
Install  29 Packages

Total download size: 29 M
Installed size: 93 M
Downloading Packages:
(1/29): atomic-registries-1.20.1-6.fc26.x86_64.rpm                                                              119 kB/s |  38 kB     00:00    
(2/29): docker-common-1.13.1-44.git584d391.fc26.x86_64.rpm                                                      169 kB/s |  83 kB     00:00    
(3/29): oci-umount-2.3.0-1.git51e7c50.fc26.x86_64.rpm                                                           223 kB/s |  34 kB     00:00    
(4/29): skopeo-containers-0.1.27-1.git93876ac.fc26.x86_64.rpm                                                    40 kB/s |  16 kB     00:00    
(5/29): docker-rhel-push-plugin-1.13.1-44.git584d391.fc26.x86_64.rpm                                            831 kB/s | 1.6 MB     00:02    
(6/29): lvm2-2.02.168-6.fc26.x86_64.rpm                                                                         554 kB/s | 1.2 MB     00:02    
(7/29): yajl-2.1.0-6.fc26.x86_64.rpm                                                                             38 kB/s |  36 kB     00:00    
(8/29): device-mapper-event-libs-1.02.137-6.fc26.x86_64.rpm                                                     708 kB/s | 242 kB     00:00    
(9/29): device-mapper-persistent-data-0.6.3-5.fc26.x86_64.rpm                                                   655 kB/s | 434 kB     00:00    
(10/29): libaio-0.3.110-7.fc26.x86_64.rpm                                                                        56 kB/s |  28 kB     00:00    
(11/29): lvm2-libs-2.02.168-6.fc26.x86_64.rpm                                                                   1.1 MB/s | 1.0 MB     00:00    
(12/29): device-mapper-event-1.02.137-6.fc26.x86_64.rpm                                                         657 kB/s | 241 kB     00:00    
(13/29): container-selinux-2.36-1.fc26.noarch.rpm                                                               112 kB/s |  35 kB     00:00    
(14/29): policycoreutils-python-utils-2.6-5.fc26.x86_64.rpm                                                     655 kB/s | 218 kB     00:00    
(15/29): container-storage-setup-0.8.0-2.git1d27ecf.fc26.noarch.rpm                                              60 kB/s |  36 kB     00:00    
(16/29): libusb-0.1.5-8.fc26.x86_64.rpm                                                                         178 kB/s |  39 kB     00:00    
(17/29): python-rhsm-certificates-1.20.2-1.fc26.x86_64.rpm                                                      226 kB/s |  44 kB     00:00    
(18/29): iptables-1.6.1-2.fc26.x86_64.rpm                                                                       917 kB/s | 434 kB     00:00    
(19/29): libnetfilter_conntrack-1.0.6-2.fc26.x86_64.rpm                                                         261 kB/s |  61 kB     00:00    
(20/29): libnfnetlink-1.0.1-9.fc26.x86_64.rpm                                                                   172 kB/s |  30 kB     00:00    
(21/29): python3-pytoml-0.1.14-1.git7dea353.fc26.noarch.rpm                                                     114 kB/s |  23 kB     00:00    
(22/29): docker-1.13.1-44.git584d391.fc26.x86_64.rpm                                                            2.6 MB/s |  20 MB     00:07    
(23/29): criu-3.6-1.fc26.x86_64.rpm                                                                             682 kB/s | 455 kB     00:00    
(24/29): libnet-1.1.6-12.fc26.x86_64.rpm                                                                        470 kB/s |  63 kB     00:00    
(25/29): oci-systemd-hook-0.1.13-1.gitafe4b4a.fc26.x86_64.rpm                                                   271 kB/s |  35 kB     00:00    
(26/29): gnupg-1.4.22-1.fc26.x86_64.rpm                                                                         487 kB/s | 1.3 MB     00:02    
(27/29): protobuf-c-1.2.1-4.fc26.x86_64.rpm                                                                      48 kB/s |  34 kB     00:00    
(28/29): oci-register-machine-0-5.11.gitcd1e331.fc26.x86_64.rpm                                                 1.5 MB/s | 1.0 MB     00:00    
(29/29): systemd-container-233-6.fc26.x86_64.rpm                                                                768 kB/s | 416 kB     00:00    
------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                           2.4 MB/s |  29 MB     00:11     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                        1/1 
  Installing       : device-mapper-event-libs-1.02.137-6.fc26.x86_64                                                                       1/29 
  Running scriptlet: device-mapper-event-libs-1.02.137-6.fc26.x86_64                                                                       1/29 
  Installing       : libnfnetlink-1.0.1-9.fc26.x86_64                                                                                      2/29 
  Running scriptlet: libnfnetlink-1.0.1-9.fc26.x86_64                                                                                      2/29 
  Installing       : yajl-2.1.0-6.fc26.x86_64                                                                                              3/29 
  Running scriptlet: yajl-2.1.0-6.fc26.x86_64                                                                                              3/29 
  Installing       : oci-umount-2:2.3.0-1.git51e7c50.fc26.x86_64                                                                           4/29 
  Installing       : libnetfilter_conntrack-1.0.6-2.fc26.x86_64                                                                            5/29 
  Running scriptlet: libnetfilter_conntrack-1.0.6-2.fc26.x86_64                                                                            5/29 
  Installing       : iptables-1.6.1-2.fc26.x86_64                                                                                          6/29 
  Running scriptlet: iptables-1.6.1-2.fc26.x86_64                                                                                          6/29 
  Installing       : device-mapper-event-1.02.137-6.fc26.x86_64                                                                            7/29 
  Running scriptlet: device-mapper-event-1.02.137-6.fc26.x86_64                                                                            7/29 
  Installing       : lvm2-libs-2.02.168-6.fc26.x86_64                                                                                      8/29 
  Running scriptlet: lvm2-libs-2.02.168-6.fc26.x86_64                                                                                      8/29 
  Installing       : systemd-container-233-6.fc26.x86_64                                                                                   9/29 
  Installing       : protobuf-c-1.2.1-4.fc26.x86_64                                                                                       10/29 
  Running scriptlet: protobuf-c-1.2.1-4.fc26.x86_64                                                                                       10/29 
  Installing       : libnet-1.1.6-12.fc26.x86_64                                                                                          11/29 
  Running scriptlet: libnet-1.1.6-12.fc26.x86_64                                                                                          11/29 
  Installing       : python3-pytoml-0.1.14-1.git7dea353.fc26.noarch                                                                       12/29 
  Installing       : atomic-registries-1.20.1-6.fc26.x86_64                                                                               13/29 
  Installing       : python-rhsm-certificates-1.20.2-1.fc26.x86_64                                                                        14/29 
  Installing       : libusb-1:0.1.5-8.fc26.x86_64                                                                                         15/29 
  Running scriptlet: libusb-1:0.1.5-8.fc26.x86_64                                                                                         15/29 
  Installing       : gnupg-1.4.22-1.fc26.x86_64                                                                                           16/29 
  Running scriptlet: gnupg-1.4.22-1.fc26.x86_64                                                                                           16/29 
  Installing       : policycoreutils-python-utils-2.6-5.fc26.x86_64                                                                       17/29 
  Installing       : container-selinux-2:2.36-1.fc26.noarch                                                                               18/29 
  Running scriptlet: container-selinux-2:2.36-1.fc26.noarch                                                                               18/29 
  Installing       : libaio-0.3.110-7.fc26.x86_64                                                                                         19/29 
  Running scriptlet: libaio-0.3.110-7.fc26.x86_64                                                                                         19/29 
  Installing       : device-mapper-persistent-data-0.6.3-5.fc26.x86_64                                                                    20/29 
  Installing       : lvm2-2.02.168-6.fc26.x86_64                                                                                          21/29 
  Running scriptlet: lvm2-2.02.168-6.fc26.x86_64                                                                                          21/29 
  Installing       : container-storage-setup-0.8.0-2.git1d27ecf.fc26.noarch                                                               22/29 
  Installing       : skopeo-containers-0.1.27-1.git93876ac.fc26.x86_64                                                                    23/29 
  Installing       : docker-rhel-push-plugin-2:1.13.1-44.git584d391.fc26.x86_64                                                           24/29 
  Running scriptlet: docker-rhel-push-plugin-2:1.13.1-44.git584d391.fc26.x86_64                                                           24/29 
  Installing       : docker-common-2:1.13.1-44.git584d391.fc26.x86_64                                                                     25/29 
  Installing       : docker-2:1.13.1-44.git584d391.fc26.x86_64                                                                            26/29 
  Running scriptlet: docker-2:1.13.1-44.git584d391.fc26.x86_64                                                                            26/29 
  Installing       : criu-3.6-1.fc26.x86_64                                                                                               27/29 
  Running scriptlet: criu-3.6-1.fc26.x86_64                                                                                               27/29 
  Installing       : oci-register-machine-0-5.11.gitcd1e331.fc26.x86_64                                                                   28/29 
  Installing       : oci-systemd-hook-1:0.1.13-1.gitafe4b4a.fc26.x86_64                                                                   29/29 
  Running scriptlet: docker-2:1.13.1-44.git584d391.fc26.x86_64                                                                            29/29 
  Running scriptlet: oci-systemd-hook-1:0.1.13-1.gitafe4b4a.fc26.x86_64                                                                   29/29 
  Verifying        : docker-2:1.13.1-44.git584d391.fc26.x86_64                                                                             1/29 
  Verifying        : atomic-registries-1.20.1-6.fc26.x86_64                                                                                2/29 
  Verifying        : docker-common-2:1.13.1-44.git584d391.fc26.x86_64                                                                      3/29 
  Verifying        : docker-rhel-push-plugin-2:1.13.1-44.git584d391.fc26.x86_64                                                            4/29 
  Verifying        : oci-umount-2:2.3.0-1.git51e7c50.fc26.x86_64                                                                           5/29 
  Verifying        : skopeo-containers-0.1.27-1.git93876ac.fc26.x86_64                                                                     6/29 
  Verifying        : lvm2-2.02.168-6.fc26.x86_64                                                                                           7/29 
  Verifying        : yajl-2.1.0-6.fc26.x86_64                                                                                              8/29 
  Verifying        : device-mapper-event-libs-1.02.137-6.fc26.x86_64                                                                       9/29 
  Verifying        : device-mapper-persistent-data-0.6.3-5.fc26.x86_64                                                                    10/29 
  Verifying        : lvm2-libs-2.02.168-6.fc26.x86_64                                                                                     11/29 
  Verifying        : libaio-0.3.110-7.fc26.x86_64                                                                                         12/29 
  Verifying        : device-mapper-event-1.02.137-6.fc26.x86_64                                                                           13/29 
  Verifying        : container-selinux-2:2.36-1.fc26.noarch                                                                               14/29 
  Verifying        : policycoreutils-python-utils-2.6-5.fc26.x86_64                                                                       15/29 
  Verifying        : container-storage-setup-0.8.0-2.git1d27ecf.fc26.noarch                                                               16/29 
  Verifying        : gnupg-1.4.22-1.fc26.x86_64                                                                                           17/29 
  Verifying        : libusb-1:0.1.5-8.fc26.x86_64                                                                                         18/29 
  Verifying        : python-rhsm-certificates-1.20.2-1.fc26.x86_64                                                                        19/29 
  Verifying        : iptables-1.6.1-2.fc26.x86_64                                                                                         20/29 
  Verifying        : libnetfilter_conntrack-1.0.6-2.fc26.x86_64                                                                           21/29 
  Verifying        : libnfnetlink-1.0.1-9.fc26.x86_64                                                                                     22/29 
  Verifying        : python3-pytoml-0.1.14-1.git7dea353.fc26.noarch                                                                       23/29 
  Verifying        : criu-3.6-1.fc26.x86_64                                                                                               24/29 
  Verifying        : libnet-1.1.6-12.fc26.x86_64                                                                                          25/29 
  Verifying        : protobuf-c-1.2.1-4.fc26.x86_64                                                                                       26/29 
  Verifying        : oci-systemd-hook-1:0.1.13-1.gitafe4b4a.fc26.x86_64                                                                   27/29 
  Verifying        : oci-register-machine-0-5.11.gitcd1e331.fc26.x86_64                                                                   28/29 
  Verifying        : systemd-container-233-6.fc26.x86_64                                                                                  29/29 

Installed:
  docker.x86_64 2:1.13.1-44.git584d391.fc26                            criu.x86_64 3.6-1.fc26                                                  
  oci-register-machine.x86_64 0-5.11.gitcd1e331.fc26                   oci-systemd-hook.x86_64 1:0.1.13-1.gitafe4b4a.fc26                      
  atomic-registries.x86_64 1.20.1-6.fc26                               container-selinux.noarch 2:2.36-1.fc26                                  
  container-storage-setup.noarch 0.8.0-2.git1d27ecf.fc26               device-mapper-event.x86_64 1.02.137-6.fc26                              
  device-mapper-event-libs.x86_64 1.02.137-6.fc26                      device-mapper-persistent-data.x86_64 0.6.3-5.fc26                       
  docker-common.x86_64 2:1.13.1-44.git584d391.fc26                     docker-rhel-push-plugin.x86_64 2:1.13.1-44.git584d391.fc26              
  gnupg.x86_64 1.4.22-1.fc26                                           iptables.x86_64 1.6.1-2.fc26                                            
  libaio.x86_64 0.3.110-7.fc26                                         libnet.x86_64 1.1.6-12.fc26                                             
  libnetfilter_conntrack.x86_64 1.0.6-2.fc26                           libnfnetlink.x86_64 1.0.1-9.fc26                                        
  libusb.x86_64 1:0.1.5-8.fc26                                         lvm2.x86_64 2.02.168-6.fc26                                             
  lvm2-libs.x86_64 2.02.168-6.fc26                                     oci-umount.x86_64 2:2.3.0-1.git51e7c50.fc26                             
  policycoreutils-python-utils.x86_64 2.6-5.fc26                       protobuf-c.x86_64 1.2.1-4.fc26                                          
  python-rhsm-certificates.x86_64 1.20.2-1.fc26                        python3-pytoml.noarch 0.1.14-1.git7dea353.fc26                          
  skopeo-containers.x86_64 0.1.27-1.git93876ac.fc26                    systemd-container.x86_64 233-6.fc26                                     
  yajl.x86_64 2.1.0-6.fc26                                            

Complete!
```

```
[vagrant@localhost ~]$ sudo groupadd docker
```

```
[vagrant@localhost ~]$ sudo usermod -aG docker vagrant
```

```
[vagrant@localhost ~]$ exit
logout
Connection to 172.17.4.59 closed.
```

```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ ssh vagrant@172.17.4.59
vagrant@172.17.4.59's password: 
Last login: Sun Dec 17 14:48:44 2017 from 172.17.4.1
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo systemctl start docker.service && sudo systemctl enable docker.service
Created symlink /etc/systemd/system/multi-user.target.wants/docker.service → /usr/lib/systemd/system/docker.service.
```

```
[vagrant@kubedev-172-17-4-59 ~]$ docker info
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
containerd version:  (expected: aa8187dbd3b7ad67d8e5e3a15115d3eef43a7ed1)
runc version: N/A (expected: 9df8b306d01f59d3a8029be411de015b7304dd8f)
init version: N/A (expected: 949e6facb77383876aeff8a6944dde66b3089574)
Security Options:
 seccomp
  WARNING: You're not using the default seccomp profile
  Profile: /etc/docker/seccomp.json
 selinux
Kernel Version: 4.11.8-300.fc26.x86_64
Operating System: Fedora 26 (Cloud Edition)
OSType: linux
Architecture: x86_64
Number of Docker Hooks: 3
CPUs: 1
Total Memory: 1.952 GiB
Name: kubedev-172-17-4-59
ID: C43B:D67P:IIPC:XSS7:ELYF:25LC:E72L:FLSQ:PUF7:Y5K7:I7BV:LZEU
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

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo systemctl -l status --no-pager docker.service
● docker.service - Docker Application Container Engine
   Loaded: loaded (/usr/lib/systemd/system/docker.service; disabled; vendor preset: disabled)
   Active: active (running) since Sun 2017-12-17 15:45:25 UTC; 1min 39s ago
     Docs: http://docs.docker.com
 Main PID: 2976 (dockerd-current)
    Tasks: 8 (limit: 8192)
   CGroup: /system.slice/docker.service
           └─2976 /usr/bin/dockerd-current --add-runtime oci=/usr/libexec/docker/docker-runc-current --default-runtime=oci --authorization-plugin=rhel-push-plugin --containerd /run/containerd.sock --exec-opt native.cgroupdriver=systemd --userland-proxy-path=/usr/libexec/docker/docker-proxy-current --init-path=/usr/libexec/docker/docker-init-current --seccomp-profile=/etc/docker/seccomp.json --selinux-enabled --log-driver=journald --storage-driver overlay2 --add-registry docker.io --add-registry registry.fedoraproject.org --add-registry registry.access.redhat.com

Dec 17 15:45:25 kubedev-172-17-4-59 dockerd-current[2976]: time="2017-12-17T15:45:25.602457868Z" level=info msg="Default bridge (docker0) is assigned with an IP address 172.18.0.0/16. Daemon option --bip can be used to set a preferred IP address"
Dec 17 15:45:25 kubedev-172-17-4-59 dockerd-current[2976]: time="2017-12-17T15:45:25.654222422Z" level=info msg="Loading containers: done."
Dec 17 15:45:25 kubedev-172-17-4-59 dockerd-current[2976]: time="2017-12-17T15:45:25.665488931Z" level=warning msg="failed to retrieve docker-runc version: unknown output format: runc version 1.0.0-rc2\nspec: 1.0.0-rc2-dev\n"
Dec 17 15:45:25 kubedev-172-17-4-59 dockerd-current[2976]: time="2017-12-17T15:45:25.666432708Z" level=warning msg="failed to retrieve docker-init version: unknown output format: tini version 0.16.1\n"
Dec 17 15:45:25 kubedev-172-17-4-59 dockerd-current[2976]: time="2017-12-17T15:45:25.666837544Z" level=info msg="Daemon has completed initialization"
Dec 17 15:45:25 kubedev-172-17-4-59 dockerd-current[2976]: time="2017-12-17T15:45:25.667001582Z" level=info msg="Docker daemon" commit="584d391/1.13.1" graphdriver=overlay2 version=1.13.1
Dec 17 15:45:25 kubedev-172-17-4-59 systemd[1]: Started Docker Application Container Engine.
Dec 17 15:45:25 kubedev-172-17-4-59 dockerd-current[2976]: time="2017-12-17T15:45:25.688257084Z" level=info msg="API listen on /var/run/docker.sock"
Dec 17 15:45:27 kubedev-172-17-4-59 dockerd-current[2976]: time="2017-12-17T15:45:27.210189292Z" level=warning msg="failed to retrieve docker-runc version: unknown output format: runc version 1.0.0-rc2\nspec: 1.0.0-rc2-dev\n"
Dec 17 15:45:27 kubedev-172-17-4-59 dockerd-current[2976]: time="2017-12-17T15:45:27.211440394Z" level=warning msg="failed to retrieve docker-init version: unknown output format: tini version 0.16.1\n"
```
