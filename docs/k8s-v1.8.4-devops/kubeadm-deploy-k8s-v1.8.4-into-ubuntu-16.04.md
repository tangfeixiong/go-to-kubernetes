# Instruction

hostname
```
ubuntu@ubuntu-xenial:~$ echo "kubedev-10-64-33-195" | sudo tee /etc/hostname
kubedev-10-64-33-195
```

```
ubuntu@ubuntu-xenial:~$ sudo hostname kubedev-10-64-33-195
ubuntu@ubuntu-xenial:~$ hostname
kubedev-10-64-33-195
```

```
ubuntu@ubuntu-xenial:~$ echo "10.64.33.195    kubedev-10-64-33-195" | sudo tee -a /etc/hosts
sudo: unable to resolve host kubedev-10-64-33-195
10.64.33.195    kubedev-10-64-33-195
```


Firwall
```
ubuntu@ubuntu-xenial:~$ sudo ufw status
Status: inactive
```

APT repository, Using [http file server on previous node](../k8s-v1.8.2-devops)
```
ubuntu@ubuntu-xenial:~$ echo -e "deb http://apt.kubernetes.io/ kubernetes-xenial main\ndeb [ arch=amd64 ] file:///vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main\ndeb [ arch=amd64 ] http://10.64.33.82:48080/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
deb [ arch=amd64 ] file:///vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main
deb [ arch=amd64 ] http://10.64.33.82:48080/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main
```

```
ubuntu@ubuntu-xenial:~$ sudo apt-key add /vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/doc/apt-key.gpg 
OK
```

```
ubuntu@ubuntu-xenial:~$ sudo rm -rf /var/lib/apt/lists/*
```

```
ubuntu@ubuntu-xenial:~$ sudo apt-get --ignore-missing update
Get:1 file:/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,942 B]
Get:1 file:/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,942 B]
Get:2 file:/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [7,460 B]
Get:3 http://10.64.33.82:48080/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,942 B]
Get:4 http://10.64.33.82:48080/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [7,460 B]
### snippets ###
```

docker.io
```
ubuntu@ubuntu-xenial:~$ sudo apt-cache show docker.io
Package: docker.io
Architecture: amd64
Version: 1.13.1-0ubuntu1~16.04.2
Built-Using: glibc (= 2.23-0ubuntu9)
Priority: optional
Section: universe/admin
Origin: Ubuntu
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Original-Maintainer: Paul Tagliamonte <paultag@debian.org>
Bugs: https://bugs.launchpad.net/ubuntu/+filebug
Installed-Size: 61245
Depends: adduser, containerd (>= 0.2.5~), iptables, runc (>= 1.0.0~rc2-0ubuntu1~), init-system-helpers (>= 1.18~), lsb-base (>= 4.1+Debian11ubuntu7), libapparmor1 (>= 2.6~devel), libc6 (>= 2.14), libdevmapper1.02.1 (>= 2:1.02.97), libseccomp2 (>= 2.1.0)
Recommends: ca-certificates, cgroupfs-mount | cgroup-lite, git, ubuntu-fan, xz-utils, apparmor
Suggests: aufs-tools, btrfs-tools, debootstrap, docker-doc, rinse, zfs-fuse | zfsutils
Breaks: docker (<< 1.5~)
Replaces: docker (<< 1.5~)
Filename: pool/universe/d/docker.io/docker.io_1.13.1-0ubuntu1~16.04.2_amd64.deb
Size: 11895572
MD5sum: 30f482b26348edfc7ab2d3cdd428623b
SHA1: b548b0b5038b4ce7fc21137c45e8c657c0034d5e
SHA256: 7069eb25c18e2b3b099db236a020eb924280f8d1b90c40f4d9dbc1d070560709
Homepage: https://dockerproject.org
Description-en: Linux container runtime
 Docker complements kernel namespacing with a high-level API which operates at
 the process level. It runs unix processes with strong guarantees of isolation
 and repeatability across servers.
 .
 Docker is a great building block for automating distributed systems:
 large-scale web deployments, database clusters, continuous deployment systems,
 private PaaS, service-oriented architectures, etc.
 .
 This package contains the daemon and client. Using docker.io on non-amd64 hosts
 is not supported at this time. Please be careful when using it on anything
 besides amd64.
 .
 Also, note that kernel version 3.8 or above is required for proper operation of
 the daemon process, and that any lower versions may have subtle and/or glaring
 issues.
Description-md5: 05dc9eba68f3bf418e6a0cf29d555878
### snippets ###
```

```
ubuntu@ubuntu-xenial:~$ sudo apt-cache show kubeadm=1.8.4-00
Package: kubeadm
Version: 1.8.4-00
Installed-Size: 133266
Maintainer: Kubernetes Authors <kubernetes-dev+release@googlegroups.com>
Architecture: amd64
Depends: kubelet (>= 1.6.0), kubectl (>= 1.6.0)
Description: Kubernetes Cluster Bootstrapping Tool
 The Kubernetes command line tool for bootstrapping a Kubernetes cluster.
Description-md5: bb3c7836839894793de38af875e01b30
Homepage: https://kubernetes.io
Filename: pool/kubeadm_1.8.4-00_amd64_0088836fbb451bc49ece82f34c035f50f2e1dd4dea78f6d585574d585da11e8e.deb
Priority: optional
SHA256: 0088836fbb451bc49ece82f34c035f50f2e1dd4dea78f6d585574d585da11e8e
Section: misc
Size: 18136438

ubuntu@ubuntu-xenial:~$ sudo apt-cache show kubectl=1.8.4-00
Package: kubectl
Version: 1.8.4-00
Installed-Size: 51125
Maintainer: Kubernetes Authors <kubernetes-dev+release@googlegroups.com>
Architecture: amd64
Description: Kubernetes Command Line Tool
 The Kubernetes command line tool for interacting with the Kubernetes API.
Description-md5: fb58ab85a9089d0257cb8f7cda7d5a09
Homepage: https://kubernetes.io
Filename: pool/kubectl_1.8.4-00_amd64_b48511a481ddcfbf935ad935bc6c3992c1c4315fcd8f3f794e367b9b26b775be.deb
Priority: optional
SHA256: b48511a481ddcfbf935ad935bc6c3992c1c4315fcd8f3f794e367b9b26b775be
Section: misc
Size: 8611520

ubuntu@ubuntu-xenial:~$ sudo apt-cache show kubelet=1.8.4-00
Package: kubelet
Version: 1.8.4-00
Installed-Size: 134422
Maintainer: Kubernetes Authors <kubernetes-dev+release@googlegroups.com>
Architecture: amd64
Depends: iptables (>= 1.4.21), kubernetes-cni (>= 0.5.1), iproute2, socat, util-linux, mount, ebtables, ethtool, init-system-helpers (>= 1.18~)
Description: Kubernetes Node Agent
 The node agent of Kubernetes, the container cluster manager
Description-md5: 0bad32a8f0183e699705f1585a26664c
Homepage: https://kubernetes.io
Filename: pool/kubelet_1.8.4-00_amd64_601882506070723b643552ae98325c849840b09b1fc1666de74c7b69a07f06df.deb
Priority: optional
SHA256: 601882506070723b643552ae98325c849840b09b1fc1666de74c7b69a07f06df
Section: misc
Size: 19208628

```

### Install

Docker
```
ubuntu@ubuntu-xenial:~$ sudo apt install -y docker.io
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  bridge-utils cgroupfs-mount containerd runc ubuntu-fan
Suggested packages:
  mountall aufs-tools debootstrap docker-doc rinse zfs-fuse | zfsutils
The following NEW packages will be installed:
  bridge-utils cgroupfs-mount containerd docker.io runc ubuntu-fan
0 upgraded, 6 newly installed, 0 to remove and 231 not upgraded.
Need to get 17.5 MB of archives.
After this operation, 90.5 MB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu xenial/main amd64 bridge-utils amd64 1.5-9ubuntu1 [28.6 kB]
Get:2 http://archive.ubuntu.com/ubuntu xenial/universe amd64 cgroupfs-mount all 1.2 [4,970 B]
Get:3 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 runc amd64 1.0.0~rc2+docker1.13.1-0ubuntu1~16.04.1 [1,488 kB]
Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 containerd amd64 0.2.5-0ubuntu1~16.04.1 [4,041 kB]            
Get:5 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 docker.io amd64 1.13.1-0ubuntu1~16.04.2 [11.9 MB]             
Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 ubuntu-fan all 0.9.2 [30.7 kB]                                    
Fetched 17.5 MB in 24s (714 kB/s)                                                                                                  
Selecting previously unselected package bridge-utils.
(Reading database ... 79718 files and directories currently installed.)
Preparing to unpack .../bridge-utils_1.5-9ubuntu1_amd64.deb ...
Unpacking bridge-utils (1.5-9ubuntu1) ...
Selecting previously unselected package cgroupfs-mount.
Preparing to unpack .../cgroupfs-mount_1.2_all.deb ...
Unpacking cgroupfs-mount (1.2) ...
Selecting previously unselected package runc.
Preparing to unpack .../runc_1.0.0~rc2+docker1.13.1-0ubuntu1~16.04.1_amd64.deb ...
Unpacking runc (1.0.0~rc2+docker1.13.1-0ubuntu1~16.04.1) ...
Selecting previously unselected package containerd.
Preparing to unpack .../containerd_0.2.5-0ubuntu1~16.04.1_amd64.deb ...
Unpacking containerd (0.2.5-0ubuntu1~16.04.1) ...
Selecting previously unselected package docker.io.
Preparing to unpack .../docker.io_1.13.1-0ubuntu1~16.04.2_amd64.deb ...
Unpacking docker.io (1.13.1-0ubuntu1~16.04.2) ...
Selecting previously unselected package ubuntu-fan.
Preparing to unpack .../ubuntu-fan_0.9.2_all.deb ...
Unpacking ubuntu-fan (0.9.2) ...
Processing triggers for man-db (2.7.5-1) ...
Processing triggers for ureadahead (0.100.0-19) ...
Processing triggers for systemd (229-4ubuntu11) ...
Setting up bridge-utils (1.5-9ubuntu1) ...
Setting up cgroupfs-mount (1.2) ...
Setting up runc (1.0.0~rc2+docker1.13.1-0ubuntu1~16.04.1) ...
Setting up containerd (0.2.5-0ubuntu1~16.04.1) ...
Setting up docker.io (1.13.1-0ubuntu1~16.04.2) ...
Adding group `docker' (GID 123) ...
Done.
Setting up ubuntu-fan (0.9.2) ...
Processing triggers for systemd (229-4ubuntu11) ...
Processing triggers for ureadahead (0.100.0-19) ...
```

```
ubuntu@ubuntu-xenial:~$ sudo usermod -aG docker ubuntu
```

logout and login
```
ubuntu@kubedev-10-64-33-195:~$ docker version
Client:
 Version:      1.13.1
 API version:  1.26
 Go version:   go1.6.2
 Git commit:   092cba3
 Built:        Thu Nov  2 20:40:23 2017
 OS/Arch:      linux/amd64

Server:
 Version:      1.13.1
 API version:  1.26 (minimum version 1.12)
 Go version:   go1.6.2
 Git commit:   092cba3
 Built:        Thu Nov  2 20:40:23 2017
 OS/Arch:      linux/amd64
 Experimental: false
```

Kubeadm
```
ubuntu@kubedev-10-64-33-195:~$ sudo apt install -y kubeadm
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  ebtables kubectl kubelet kubernetes-cni socat
The following NEW packages will be installed:
  ebtables kubeadm kubectl kubelet kubernetes-cni socat
0 upgraded, 6 newly installed, 0 to remove and 231 not upgraded.
Need to get 400 kB/51.9 MB of archives.
After this operation, 370 MB of additional disk space will be used.
Get:1 file:/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubernetes-cni amd64 0.5.1-00 [5,560 kB]
Get:2 file:/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubelet amd64 1.8.4-00 [19.2 MB]
Get:3 file:/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubectl amd64 1.8.4-00 [8,612 kB]
Get:4 file:/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.8.4-00 [18.1 MB]
Get:5 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 ebtables amd64 2.0.10.4-3.4ubuntu2 [79.4 kB]
Get:6 http://archive.ubuntu.com/ubuntu xenial/universe amd64 socat amd64 1.7.3.1-1 [321 kB]                                        
Fetched 400 kB in 8s (49.1 kB/s)                                                                                                   
Selecting previously unselected package ebtables.
(Reading database ... 79910 files and directories currently installed.)
Preparing to unpack .../ebtables_2.0.10.4-3.4ubuntu2_amd64.deb ...
Unpacking ebtables (2.0.10.4-3.4ubuntu2) ...
Selecting previously unselected package kubernetes-cni.
Preparing to unpack .../kubernetes-cni_0.5.1-00_amd64_08cbe5c42366ec3184cc91a4353f6e066f2d7325b4db1cb4f87c7dcc8c0eb620.deb ...
Unpacking kubernetes-cni (0.5.1-00) ...
Selecting previously unselected package socat.
Preparing to unpack .../socat_1.7.3.1-1_amd64.deb ...
Unpacking socat (1.7.3.1-1) ...
Selecting previously unselected package kubelet.
Preparing to unpack .../kubelet_1.8.4-00_amd64_601882506070723b643552ae98325c849840b09b1fc1666de74c7b69a07f06df.deb ...
Unpacking kubelet (1.8.4-00) ...
Selecting previously unselected package kubectl.
Preparing to unpack .../kubectl_1.8.4-00_amd64_b48511a481ddcfbf935ad935bc6c3992c1c4315fcd8f3f794e367b9b26b775be.deb ...
Unpacking kubectl (1.8.4-00) ...
Selecting previously unselected package kubeadm.
Preparing to unpack .../kubeadm_1.8.4-00_amd64_0088836fbb451bc49ece82f34c035f50f2e1dd4dea78f6d585574d585da11e8e.deb ...
Unpacking kubeadm (1.8.4-00) ...
Processing triggers for systemd (229-4ubuntu11) ...
Processing triggers for ureadahead (0.100.0-19) ...
Processing triggers for man-db (2.7.5-1) ...
Setting up ebtables (2.0.10.4-3.4ubuntu2) ...
update-rc.d: warning: start and stop actions are no longer supported; falling back to defaults
Setting up kubernetes-cni (0.5.1-00) ...
Setting up socat (1.7.3.1-1) ...
Setting up kubelet (1.8.4-00) ...
Setting up kubectl (1.8.4-00) ...
Setting up kubeadm (1.8.4-00) ...
Processing triggers for systemd (229-4ubuntu11) ...
Processing triggers for ureadahead (0.100.0-19) ...
```

```
ubuntu@kubedev-10-64-33-195:~$ kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.4", GitCommit:"9befc2b8928a9426501d3bf62f72849d5cbcd5a3", GitTreeState:"clean", BuildDate:"2017-11-20T05:17:43Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
```

```
ubuntu@kubedev-10-64-33-195:~$ sudo cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CERTIFICATE_ARGS=--rotate-certificates=true --cert-dir=/var/lib/kubelet/pki"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS
```

```
ubuntu@kubedev-10-64-33-195:~$ sudo cat /lib/systemd/system/kubelet.service 
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
ubuntu@kubedev-10-64-33-195:~$ ls /etc/kubernetes/
manifests
ubuntu@kubedev-10-64-33-195:~$ ls /etc/kubernetes/manifests/
```

### Copy from previous http node

K8s images
```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/home/vagrant/kubernetes/server/bin/kube-apiserver.tar | docker load 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  185M  100  185M    0     0  58.3M      0  0:00:03  0:00:03 --:--:-- 58.3M
0271b8eebde3: Loading layer [==================================================>] 1.338 MB/1.338 MB
f72e92775dd7: Loading layer [==================================================>] 193.2 MB/193.2 MB
Loaded image: gcr.io/google_containers/kube-apiserver:v1.8.4
ubuntu@kubedev-10-64-33-195:~$ docker tag gcr.io/google_containers/kube-apiserver:v1.8.4 gcr.io/google_containers/kube-apiserver-amd64:v1.8.4
ubuntu@kubedev-10-64-33-195:~$ docker rmi gcr.io/google_containers/kube-apiserver:v1.8.4
Untagged: gcr.io/google_containers/kube-apiserver:v1.8.4
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/home/vagrant/kubernetes/server/bin/kube-controller-manager.tar | docker load 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  123M  100  123M    0     0  45.2M      0  0:00:02  0:00:02 --:--:-- 45.2M
4f92df6d8677: Loading layer [==================================================>] 128.2 MB/128.2 MB
Loaded image: gcr.io/google_containers/kube-controller-manager:v1.8.4
ubuntu@kubedev-10-64-33-195:~$ docker tag gcr.io/google_containers/kube-controller-manager:v1.8.4 gcr.io/google_containers/kube-controller-manager-amd64:v1.8.4
ubuntu@kubedev-10-64-33-195:~$ docker rmi gcr.io/google_containers/kube-controller-manager:v1.8.4
Untagged: gcr.io/google_containers/kube-controller-manager:v1.8.4
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/home/vagrant/kubernetes/server/bin/kube-scheduler.tar | docker load 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 52.6M  100 52.6M    0     0  47.5M      0  0:00:01  0:00:01 --:--:-- 47.6M
da6851a1e488: Loading layer [==================================================>] 53.85 MB/53.85 MB
Loaded image: gcr.io/google_containers/kube-scheduler:v1.8.4
ubuntu@kubedev-10-64-33-195:~$ docker tag gcr.io/google_containers/kube-scheduler:v1.8.4 gcr.io/google_containers/kube-scheduler-amd64:v1.8.4
ubuntu@kubedev-10-64-33-195:~$ docker rmi gcr.io/google_containers/kube-scheduler:v1.8.4
Untagged: gcr.io/google_containers/kube-scheduler:v1.8.4
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/home/vagrant/kubernetes/server/bin/kube-proxy.tar | docker load 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.6M  100 90.6M    0     0  41.0M      0  0:00:02  0:00:02 --:--:-- 41.0M
08ae86c4c4c9: Loading layer [==================================================>] 42.05 MB/42.05 MB
48a108f012f8: Loading layer [==================================================>] 5.045 MB/5.045 MB
9fc6ccb688b9: Loading layer [==================================================>] 47.93 MB/47.93 MB
Loaded image: gcr.io/google_containers/kube-proxy:v1.8.4
ubuntu@kubedev-10-64-33-195:~$ docker tag gcr.io/google_containers/kube-proxy:v1.8.4 gcr.io/google_containers/kube-proxy-amd64:v1.8.4
ubuntu@kubedev-10-64-33-195:~$ docker rmi gcr.io/google_containers/kube-proxy:v1.8.4
Untagged: gcr.io/google_containers/kube-proxy:v1.8.4
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fetcd0x2E3.0.17.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  161M  100  161M    0     0  53.6M      0  0:00:03  0:00:03 --:--:-- 53.6M
38ac8d0f5bb3: Loading layer [==================================================>] 1.312 MB/1.312 MB
c872b2c2ac77: Loading layer [==================================================>] 136.7 MB/136.7 MB
be71b2a80bbd: Loading layer [==================================================>] 31.16 MB/31.16 MB
Loaded image: gcr.io/google_containers/etcd:3.0.17
ubuntu@kubedev-10-64-33-195:~$ docker tag gcr.io/google_containers/etcd:3.0.17 gcr.io/google_containers/etcd-amd64:3.0.17
ubuntu@kubedev-10-64-33-195:~$ docker rmi gcr.io/google_containers/etcd:3.0.17
Untagged: gcr.io/google_containers/etcd:3.0.17
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-dnsmasq-nanny-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 39.7M  100 39.7M    0     0  32.6M      0  0:00:01  0:00:01 --:--:-- 32.6M
e620d0ac6539: Loading layer [==================================================>]  2.56 kB/2.56 kB
9f4f5a30979e: Loading layer [==================================================>]   362 kB/362 kB
fd7319ac31de: Loading layer [==================================================>] 3.072 kB/3.072 kB
b23d166217e1: Loading layer [==================================================>]  37.1 MB/37.1 MB
Loaded image: gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-kube-dns-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 47.3M  100 47.3M    0     0  32.0M      0  0:00:01  0:00:01 --:--:-- 32.1M
a1a7a797fc0e: Loading layer [==================================================>] 45.42 MB/45.42 MB
Loaded image: gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.5
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-sidecar-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 40.1M  100 40.1M    0     0  32.0M      0  0:00:01  0:00:01 --:--:-- 32.1M
acfc450a47fa: Loading layer [==================================================>] 37.86 MB/37.86 MB
Loaded image: gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.5
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  745k  100  745k    0     0  5993k      0 --:--:-- --:--:-- --:--:-- 6161k
e10b0b2a65fb: Loading layer [==================================================>] 1.024 kB/1.024 kB
b7402af6e348: Loading layer [==================================================>] 748.5 kB/748.5 kB
b2091bb33512: Loading layer [==================================================>] 1.024 kB/1.024 kB
open /var/lib/docker/tmp/docker-import-259200595/repositories: no such file or directory
ubuntu@kubedev-10-64-33-195:~$ docker images
REPOSITORY                                               TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver-amd64            v1.8.4              10a052dccbc5        2 days ago          194 MB
gcr.io/google_containers/kube-controller-manager-amd64   v1.8.4              7058ac4d4af5        2 days ago          129 MB
gcr.io/google_containers/kube-proxy-amd64                v1.8.4              65a61c14e8c2        2 days ago          93.2 MB
gcr.io/google_containers/kube-scheduler-amd64            v1.8.4              0d985fed7f95        2 days ago          55 MB
gcr.io/google_containers/etcd-amd64                      3.0.17              243830dae7dd        9 months ago        169 MB
<none>                                                   <none>              99e59f495ffa        18 months ago       747 kB
ubuntu@kubedev-10-64-33-195:~$ docker tag 99e59f495ffa gcr.io/google_containers/pause-amd64:3.0
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/quay.io0x2Fcoreos0x2Fflannel0x3Av0.9.0-amd64.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 49.7M  100 49.7M    0     0  49.5M      0  0:00:01  0:00:01 --:--:-- 49.6M
5bef08742407: Loading layer [==================================================>] 4.221 MB/4.221 MB
f439636ab0f0: Loading layer [==================================================>] 6.797 MB/6.797 MB
91b6f6ead101: Loading layer [==================================================>] 4.414 MB/4.414 MB
39f837629582: Loading layer [==================================================>] 34.49 MB/34.49 MB
d3e99a0118c5: Loading layer [==================================================>] 2.225 MB/2.225 MB
32adca76eade: Loading layer [==================================================>]  5.12 kB/5.12 kB
Loaded image: quay.io/coreos/flannel:v0.9.0-amd64
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/home/vagrant/kube-flannel.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2849  100  2849    0     0  19116      0 --:--:-- --:--:-- --:--:-- 19250
```

```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/home/vagrant/weave-daemonset-k8s-1.6.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  4579  100  4579    0     0   164k      0 --:--:-- --:--:-- --:--:--  171k
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/home/vagrant/weave-daemonset-k8s-1.7.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  5064  100  5064    0     0   452k      0 --:--:-- --:--:-- --:--:--  494k
```

```
ubuntu@kubedev-10-64-33-195:~$ docker images
REPOSITORY                                               TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver-amd64            v1.8.4              10a052dccbc5        2 days ago          194 MB
gcr.io/google_containers/kube-controller-manager-amd64   v1.8.4              7058ac4d4af5        2 days ago          129 MB
gcr.io/google_containers/kube-proxy-amd64                v1.8.4              65a61c14e8c2        2 days ago          93.2 MB
gcr.io/google_containers/kube-scheduler-amd64            v1.8.4              0d985fed7f95        2 days ago          55 MB
gcr.io/google_containers/k8s-dns-sidecar-amd64           1.14.5              fed89e8b4248        8 weeks ago         41.8 MB
gcr.io/google_containers/k8s-dns-kube-dns-amd64          1.14.5              512cd7425a73        8 weeks ago         49.4 MB
gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64     1.14.5              459944ce8cc4        8 weeks ago         41.4 MB
quay.io/coreos/flannel                                   v0.9.0-amd64        4c600a64a18a        2 months ago        51.3 MB
gcr.io/google_containers/etcd-amd64                      3.0.17              243830dae7dd        9 months ago        169 MB
gcr.io/google_containers/pause-amd64                     3.0                 99e59f495ffa        18 months ago       747 kB
```

### Deploy k8s

networking
```
ubuntu@kubedev-10-64-33-195:~$ ip a show enp0s8; ip a show enp0s9
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:c4:54:fb brd ff:ff:ff:ff:ff:ff
    inet 10.64.33.195/24 brd 10.64.33.255 scope global enp0s8
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fec4:54fb/64 scope link 
       valid_lft forever preferred_lft forever
4: enp0s9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:e7:a0:16 brd ff:ff:ff:ff:ff:ff
    inet 172.28.128.6/24 brd 172.28.128.255 scope global enp0s9
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fee7:a016/64 scope link 
       valid_lft forever preferred_lft forever
```

```
ubuntu@kubedev-10-64-33-195:~$ sudo kubeadm init --apiserver-advertise-address 10.64.33.195 --apiserver-bind-port 443 --apiserver-cert-extra-sans 10.64.33.195,172.28.128.6 --pod-network-cidr 10.244.0.0/16 --kubernetes-version v1.8.4 --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.8.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [kubedev-10-64-33-195 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.64.33.195 10.64.33.195 172.28.128.6]
[certificates] Generated apiserver-kubelet-client certificate and key.
[certificates] Generated sa key and public key.
[certificates] Generated front-proxy-ca certificate and key.
[certificates] Generated front-proxy-client certificate and key.
[certificates] Valid certificates and keys now exist in "/etc/kubernetes/pki"
[kubeconfig] Wrote KubeConfig file to disk: "admin.conf"
[kubeconfig] Wrote KubeConfig file to disk: "kubelet.conf"
[kubeconfig] Wrote KubeConfig file to disk: "controller-manager.conf"
[kubeconfig] Wrote KubeConfig file to disk: "scheduler.conf"
[controlplane] Wrote Static Pod manifest for component kube-apiserver to "/etc/kubernetes/manifests/kube-apiserver.yaml"
[controlplane] Wrote Static Pod manifest for component kube-controller-manager to "/etc/kubernetes/manifests/kube-controller-manager.yaml"
[controlplane] Wrote Static Pod manifest for component kube-scheduler to "/etc/kubernetes/manifests/kube-scheduler.yaml"
[etcd] Wrote Static Pod manifest for a local etcd instance to "/etc/kubernetes/manifests/etcd.yaml"
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests"
[init] This often takes around a minute; or longer if the control plane images have to be pulled.
[apiclient] All control plane components are healthy after 27.503272 seconds
[uploadconfig] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[markmaster] Will mark node kubedev-10-64-33-195 as master by adding a label and a taint
[markmaster] Master kubedev-10-64-33-195 tainted and labelled with key/value: node-role.kubernetes.io/master=""
[bootstraptoken] Using token: b41473.c2b56645c2cec228
[bootstraptoken] Configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstraptoken] Configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstraptoken] Configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstraptoken] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[addons] Applied essential addon: kube-dns
[addons] Applied essential addon: kube-proxy

Your Kubernetes master has initialized successfully!

To start using your cluster, you need to run (as a regular user):

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  http://kubernetes.io/docs/admin/addons/

You can now join any number of machines by running the following on each node
as root:

  kubeadm join --token b41473.c2b56645c2cec228 10.64.33.195:443 --discovery-token-ca-cert-hash sha256:d4ad7317762418d23c0122141a8d1734e6f82ecd8bdee191f4fc74924fbc836e

```

```
ubuntu@kubedev-10-64-33-195:~$ sudo chmod o+r /etc/kubernetes/admin.conf 
ubuntu@kubedev-10-64-33-195:~$ export KUBECONFIG=/etc/kubernetes/admin.conf 
ubuntu@kubedev-10-64-33-195:~$ kubectl version
Client Version: version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.4", GitCommit:"9befc2b8928a9426501d3bf62f72849d5cbcd5a3", GitTreeState:"clean", BuildDate:"2017-11-20T05:28:34Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.4", GitCommit:"9befc2b8928a9426501d3bf62f72849d5cbcd5a3", GitTreeState:"clean", BuildDate:"2017-11-20T05:17:43Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
ubuntu@kubedev-10-64-33-195:~$ kubectl get nodes
NAME                   STATUS     ROLES     AGE       VERSION
kubedev-10-64-33-195   NotReady   master    1m        v1.8.4
```

```
ubuntu@kubedev-10-64-33-195:~$ kubectl get all --all-namespaces
NAMESPACE     NAME            DESIRED   CURRENT   READY     UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
kube-system   ds/kube-proxy   1         1         1         1            1           <none>          2m

NAMESPACE     NAME              DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kube-system   deploy/kube-dns   1         1         1            0           2m

NAMESPACE     NAME                     DESIRED   CURRENT   READY     AGE
kube-system   rs/kube-dns-545bc4bfd4   1         1         0         1m

NAMESPACE     NAME              DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kube-system   deploy/kube-dns   1         1         1            0           2m

NAMESPACE     NAME            DESIRED   CURRENT   READY     UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
kube-system   ds/kube-proxy   1         1         1         1            1           <none>          2m

NAMESPACE     NAME                     DESIRED   CURRENT   READY     AGE
kube-system   rs/kube-dns-545bc4bfd4   1         1         0         1m

NAMESPACE     NAME                                              READY     STATUS    RESTARTS   AGE
kube-system   po/etcd-kubedev-10-64-33-195                      1/1       Running   0          56s
kube-system   po/kube-apiserver-kubedev-10-64-33-195            1/1       Running   0          1m
kube-system   po/kube-controller-manager-kubedev-10-64-33-195   1/1       Running   0          1m
kube-system   po/kube-dns-545bc4bfd4-6dhfr                      0/3       Pending   0          1m
kube-system   po/kube-proxy-c9tfw                               1/1       Running   0          1m
kube-system   po/kube-scheduler-kubedev-10-64-33-195            1/1       Running   0          1m

NAMESPACE     NAME             TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)         AGE
default       svc/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP         2m
kube-system   svc/kube-dns     ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP   2m
```


Node not ready, so install CNI network
```
ubuntu@kubedev-10-64-33-195:~$ kubectl --namespace=kube-system create -f weave-daemonset-k8s-1.7.yaml 
serviceaccount "weave-net" created
clusterrole "weave-net" created
clusterrolebinding "weave-net" created
role "weave-net" created
rolebinding "weave-net" created
daemonset "weave-net" created
```

```
ubuntu@kubedev-10-64-33-195:~$ kubectl --namespace=kube-system get pods
NAME                                           READY     STATUS    RESTARTS   AGE
etcd-kubedev-10-64-33-195                      1/1       Running   0          8m
kube-apiserver-kubedev-10-64-33-195            1/1       Running   0          8m
kube-controller-manager-kubedev-10-64-33-195   1/1       Running   0          8m
kube-dns-545bc4bfd4-6dhfr                      2/3       Running   0          9m
kube-proxy-c9tfw                               1/1       Running   0          9m
kube-scheduler-kubedev-10-64-33-195            1/1       Running   0          8m
weave-net-qmbxs                                2/2       Running   0          1m
```