# DevOps

## Table of contents
1. Archive of release
1. Images from archive
1. Other images from local file server [](./pod-etcd-dns-and-cni-images.md)
1. Install kubeadm from google repository [](#install-kubeadmin-from-google-repository)
>* Detailed background 
1. Offline repository [](./kubernetes-offline-repo.md)

## Prerequisites
* [Virtual machine for CentOS 7 or Fedora 26](/vagrantup.md)
* Docker [for CentOS 7](../docker-for-centos7.md) or [for Fedora 26](../docker-for-fedora26.md)
* Dockerized [file server for CentOS](../dockerized-file-server-for-centos7.md) or [for Fedora 26](../dockerized-file-server-for-fedora26.md) 

## Archive of release

Download archive into local file server
```
fanhonglingdeMacBook-Pro:~ fanhongling$ curl -jkSL -c - https://dl.k8s.io/v1.9.0/kubernetes-server-linux-amd64.tar.gz -o ./Downloads/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/kubernetes-server-linux-amd64_1.9.0.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     61      0  0:00:02  0:00:02 --:--:--    61
100  399M  100  399M    0     0  1047k      0  0:06:30  0:06:30 --:--:-- 1282k
```

```
fanhonglingdeMacBook-Pro:~ fanhongling$ curl -jkSL -c - https://dl.k8s.io/v1.9.0/kubernetes.tar.gz -o ./Downloads/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/kubernetes_1.9.0.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     61      0  0:00:02  0:00:02 --:--:--    61
100 2805k  100 2805k    0     0   380k      0  0:00:07  0:00:07 --:--:--  701k
```

Extract archive into Fedora 26
```
[vagrant@kubedev-172-17-4-59 ~]$ curl -L http://172.17.4.59:48080/Downloads/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/kubernetes-server-linux-amd64_1.9.0.tar.gz | sudo tar -C /opt -xz  
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  399M  100  399M    0     0  14.9M      0  0:00:26  0:00:26 --:--:-- 18.3M
```

```
[vagrant@kubedev-172-17-4-59 ~]$ curl -L http://172.17.4.59:48080/Downloads/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/kubernetes_1.9.0.tar.gz | sudo tar -C /opt -xz 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 2805k  100 2805k    0     0  10.6M      0 --:--:-- --:--:-- --:--:-- 10.7M
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo tar -C /opt/kubernetes/ -xzf /opt/kubernetes/kubernetes-src.tar.gz cluster/saltbase/
```

Show
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo ls /opt/kubernetes/server/bin
apiextensions-apiserver		     kube-aggregator		 kube-apiserver.tar		     kube-proxy.docker_tag	kubeadm
cloud-controller-manager	     kube-aggregator.docker_tag  kube-controller-manager	     kube-proxy.tar		kubectl
cloud-controller-manager.docker_tag  kube-aggregator.tar	 kube-controller-manager.docker_tag  kube-scheduler		kubelet
cloud-controller-manager.tar	     kube-apiserver		 kube-controller-manager.tar	     kube-scheduler.docker_tag	mounter
hyperkube			     kube-apiserver.docker_tag	 kube-proxy			     kube-scheduler.tar
```

Bin path
```
[vagrant@kubedev-172-17-4-59 ~]$ mkdir bin
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo chgrp -R vagrant /opt/kubernetes/                  
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo ln -sf /opt/kubernetes/server/bin/kubectl `pwd`/bin/
```

Check version
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl version --client
Client Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T21:07:38Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
```

## Images from archive

Load into Fedora 26
```
[vagrant@kubedev-172-17-4-59 ~]$ bd=/opt/kubernetes/server/bin; reg=gcr.io/google_containers; for ia in $(ls $bd/*.tar); do name=$(basename $ia .tar); tag=$(cat $bd/$name.docker_tag); echo "$name..."; docker load -i $ia; img=$(docker images -q $reg/$name:$tag); [[ -n $img ]] && echo "...$img" || (echo "...failed" && docker images $reg/$name:$tag); done; docker images $reg/*
cloud-controller-manager...
0271b8eebde3: Loading layer [==================================================>] 1.338 MB/1.338 MB
7dd6b91138c0: Loading layer [==================================================>] 116.7 MB/116.7 MB
Loaded image: gcr.io/google_containers/cloud-controller-manager:v1.9.0
...4c938e6fc795
kube-aggregator...
98350adc2da8: Loading layer [==================================================>] 54.69 MB/54.69 MB
Loaded image: gcr.io/google_containers/kube-aggregator:v1.9.0
...e08db9577664
kube-apiserver...
9ccc9fba4253: Loading layer [==================================================>] 209.2 MB/209.2 MB
Loaded image: gcr.io/google_containers/kube-apiserver:v1.9.0
...7bff5aa286d7
kube-controller-manager...
50a426d115f8: Loading layer [==================================================>] 136.6 MB/136.6 MB
Loaded image: gcr.io/google_containers/kube-controller-manager:v1.9.0
...3bb172f9452c
kube-proxy...
684c19bf2c27: Loading layer [==================================================>]  44.2 MB/44.2 MB
deb4ca39ea31: Loading layer [==================================================>] 3.358 MB/3.358 MB
9c44b0d51ed1: Loading layer [==================================================>] 63.38 MB/63.38 MB
Loaded image: gcr.io/google_containers/kube-proxy:v1.9.0
...f6f363e6e98e
kube-scheduler...
f733b8f8af29: Loading layer [==================================================>] 61.57 MB/61.57 MB
Loaded image: gcr.io/google_containers/kube-scheduler:v1.9.0
...5ceb21996307
REPOSITORY                                          TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-proxy                 v1.9.0              f6f363e6e98e        45 hours ago        109 MB
gcr.io/google_containers/cloud-controller-manager   v1.9.0              4c938e6fc795        45 hours ago        118 MB
gcr.io/google_containers/kube-apiserver             v1.9.0              7bff5aa286d7        45 hours ago        210 MB
gcr.io/google_containers/kube-aggregator            v1.9.0              e08db9577664        45 hours ago        55.8 MB
gcr.io/google_containers/kube-controller-manager    v1.9.0              3bb172f9452c        45 hours ago        138 MB
gcr.io/google_containers/kube-scheduler             v1.9.0              5ceb21996307        45 hours ago        62.7 MB
```

## Install kubeadm from google repository

Packages
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo dnf search -y kubeadm kubelet kubectl
Failed to set locale, defaulting to C
Importing GPG key 0xA7317B0F:
 Userid     : "Google Cloud Packages Automatic Signing Key <gc-team@google.com>"
 Fingerprint: D0BC 747F D8CA F711 7500 D6FA 3746 C208 A731 7B0F
 From       : https://packages.cloud.google.com/yum/doc/yum-key.gpg
Importing GPG key 0x3E1BA8D5:
 Userid     : "Google Cloud Packages RPM Signing Key <gc-team@google.com>"
 Fingerprint: 3749 E1BA 95A8 6CE0 5454 6ED2 F09C 394C 3E1B A8D5
 From       : https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
Kubernetes                                                                                                       33 kB/s |  29 kB     00:00    
Last metadata expiration check: 0:00:22 ago on Mon Dec 18 06:26:52 2017.
======================================================== Name Exactly Matched: kubelet =========================================================
kubelet.x86_64 : Container cluster management
======================================================== Name Exactly Matched: kubectl =========================================================
kubectl.x86_64 : Command-line utility for interacting with a Kubernetes cluster.
======================================================== Name Exactly Matched: kubeadm =========================================================
kubeadm.x86_64 : Command-line utility for administering a Kubernetes cluster.
============================================================ Name Matched: kubeadm =============================================================
kubernetes-kubeadm.x86_64 : Kubernetes tool for standing up clusters
```

Install
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo dnf install -y kubeadm 
Failed to set locale, defaulting to C
Last metadata expiration check: 0:06:12 ago on Mon Dec 18 06:26:52 2017.
Dependencies resolved.
================================================================================================================================================
 Package                              Arch                         Version                               Repository                        Size
================================================================================================================================================
Installing:
 kubeadm                              x86_64                       1.9.0-0                               kubernetes                        16 M
Installing dependencies:
 ebtables                             x86_64                       2.0.10-22.fc26                        fedora                           125 k
 ethtool                              x86_64                       2:4.13-1.fc26                         updates                          137 k
 kubectl                              x86_64                       1.9.0-0                               kubernetes                       8.9 M
 kubelet                              x86_64                       1.9.0-0                               kubernetes                        17 M
 kubernetes-cni                       x86_64                       0.6.0-0                               kubernetes                       8.6 M
 socat                                x86_64                       1.7.3.2-2.fc26                        fedora                           290 k

Transaction Summary
================================================================================================================================================
Install  7 Packages

Total download size: 51 M
Installed size: 275 M
Downloading Packages:
(1/7): aa9948f82e7af317c97a242f3890985159c09c183b46ac8aab19d2ad307e6970-kubeadm-1.9.0-0.x86_64.rpm              719 kB/s |  16 MB     00:23    
(2/7): ebtables-2.0.10-22.fc26.x86_64.rpm                                                                       189 kB/s | 125 kB     00:00    
(3/7): socat-1.7.3.2-2.fc26.x86_64.rpm                                                                          266 kB/s | 290 kB     00:01    
(4/7): 8f507de9e1cc26e5b0043e334e26d62001c171d8e54d839128e9bade25ecda95-kubelet-1.9.0-0.x86_64.rpm              506 kB/s |  17 MB     00:33    
(5/7): ethtool-4.13-1.fc26.x86_64.rpm                                                                           146 kB/s | 137 kB     00:00    
(6/7): fe33057ffe95bfae65e2f269e1b05e99308853176e24a4d027bc082b471a07c0-kubernetes-cni-0.6.0-0.x86_64.rpm       761 kB/s | 8.6 MB     00:11    
(7/7): bc390a3d43256791bfb844696e7215fd7ad8a09f70a42667dab4997415a6ba75-kubectl-1.9.0-0.x86_64.rpm              191 kB/s | 8.9 MB     00:47    
------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                           1.0 MB/s |  51 MB     00:50     
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
  Installing       : ethtool-2:4.13-1.fc26.x86_64                                                                                           1/7 
  Installing       : socat-1.7.3.2-2.fc26.x86_64                                                                                            2/7 
  Installing       : ebtables-2.0.10-22.fc26.x86_64                                                                                         3/7 
  Running scriptlet: ebtables-2.0.10-22.fc26.x86_64                                                                                         3/7 
  Installing       : kubernetes-cni-0.6.0-0.x86_64                                                                                          4/7 
  Installing       : kubelet-1.9.0-0.x86_64                                                                                                 5/7 
  Installing       : kubectl-1.9.0-0.x86_64                                                                                                 6/7 
  Installing       : kubeadm-1.9.0-0.x86_64                                                                                                 7/7 
  Running scriptlet: kubeadm-1.9.0-0.x86_64                                                                                                 7/7 
  Verifying        : kubeadm-1.9.0-0.x86_64                                                                                                 1/7 
  Verifying        : kubectl-1.9.0-0.x86_64                                                                                                 2/7 
  Verifying        : kubelet-1.9.0-0.x86_64                                                                                                 3/7 
  Verifying        : ebtables-2.0.10-22.fc26.x86_64                                                                                         4/7 
  Verifying        : socat-1.7.3.2-2.fc26.x86_64                                                                                            5/7 
  Verifying        : kubernetes-cni-0.6.0-0.x86_64                                                                                          6/7 
  Verifying        : ethtool-2:4.13-1.fc26.x86_64                                                                                           7/7 

Installed:
  kubeadm.x86_64 1.9.0-0         ebtables.x86_64 2.0.10-22.fc26  ethtool.x86_64 2:4.13-1.fc26  kubectl.x86_64 1.9.0-0  kubelet.x86_64 1.9.0-0 
  kubernetes-cni.x86_64 0.6.0-0  socat.x86_64 1.7.3.2-2.fc26    

Complete!
```

### Detailed background

check kubelet.service
```
[vagrant@kubedev-172-17-4-59 ~]$ cat /etc/systemd/system/kubelet.service
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

check service override
```
[vagrant@kubedev-172-17-4-59 ~]$ cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
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
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS
```

check conf location
```
[vagrant@kubedev-172-17-4-59 ~]$ ls -R /etc/kubernetes/          
/etc/kubernetes/:
manifests

/etc/kubernetes/manifests:
```

check cni
```
[vagrant@kubedev-172-17-4-59 ~]$ ls /opt/cni/bin/
bridge  dhcp  flannel  host-local  ipvlan  loopback  macvlan  portmap  ptp  sample  tuning  vlan
```