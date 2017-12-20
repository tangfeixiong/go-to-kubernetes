# DevOps

## Prerequisites

1. [Virtual machine](vagrantup.md)
1. [Docker](../docker-1.13-for-ubuntu16.04.md)
1. [Dockerized file server] (../dockerized-file-server.md)
1. [Hostname, hosts, and firewall](../host-name-and-security-for-kubernetes.md)


## Kubernetes release archive

Already extracted into Fedora 26, refer to [README](./README.md)

and `kubectl` will be installed by `kubeadm`

### Images from archive

Through tar over ssh Fedora 26
```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "ls /opt/kubernetes/server/bin"
The authenticity of host '172.17.4.59 (172.17.4.59)' can't be established.
ECDSA key fingerprint is SHA256:VV533NeJHFykYUa3CBNDm0zDI4HSnEaMKcQSKwtvLGw.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.17.4.59' (ECDSA) to the list of known hosts.
vagrant@172.17.4.59's password: 
apiextensions-apiserver
cloud-controller-manager
cloud-controller-manager.docker_tag
cloud-controller-manager.tar
hyperkube
kubeadm
kube-aggregator
kube-aggregator.docker_tag
kube-aggregator.tar
kube-apiserver
kube-apiserver.docker_tag
kube-apiserver.tar
kube-controller-manager
kube-controller-manager.docker_tag
kube-controller-manager.tar
kubectl
kubelet
kube-proxy
kube-proxy.docker_tag
kube-proxy.tar
kube-scheduler
kube-scheduler.docker_tag
kube-scheduler.tar
mounter
```

```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "cat /opt/kubernetes/server/bin/kube-apiserver.tar" | docker load
vagrant@172.17.4.59's password: 
0271b8eebde3: Loading layer [==================================================>] 1.338 MB/1.338 MB
9ccc9fba4253: Loading layer [==================================================>] 209.2 MB/209.2 MB
Loaded image: gcr.io/google_containers/kube-apiserver:v1.9.0
```

```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "cat /opt/kubernetes/server/bin/kube-controller-manager.tar" | docker load
vagrant@172.17.4.59's password: 
50a426d115f8: Loading layer [==================================================>] 136.6 MB/136.6 MB
Loaded image: gcr.io/google_containers/kube-controller-manager:v1.9.0
```

```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "cat /opt/kubernetes/server/bin/kube-scheduler.tar" | docker load
vagrant@172.17.4.59's password: 
f733b8f8af29: Loading layer [==================================================>] 61.57 MB/61.57 MB
Loaded image: gcr.io/google_containers/kube-scheduler:v1.9.0
```

```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "cat /opt/kubernetes/server/bin/kube-proxy.tar" | docker load
vagrant@172.17.4.59's password: 
Permission denied, please try again.
vagrant@172.17.4.59's password: 
684c19bf2c27: Loading layer [==================================================>]  44.2 MB/44.2 MB
deb4ca39ea31: Loading layer [==================================================>] 3.358 MB/3.358 MB
9c44b0d51ed1: Loading layer [==================================================>] 63.38 MB/63.38 MB
Loaded image: gcr.io/google_containers/kube-proxy:v1.9.0
```

```
ubuntu@ubuntu-xenial:~$ docker images gcr.io/google_containers/*
REPOSITORY                                         TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-proxy                v1.9.0              f6f363e6e98e        3 days ago          109 MB
gcr.io/google_containers/kube-apiserver            v1.9.0              7bff5aa286d7        3 days ago          210 MB
gcr.io/google_containers/kube-controller-manager   v1.9.0              3bb172f9452c        3 days ago          138 MB
gcr.io/google_containers/kube-scheduler            v1.9.0              5ceb21996307        3 days ago          62.7 MB
```

tag 
```
ubuntu@ubuntu-xenial:~$ docker tag gcr.io/google_containers/kube-apiserver:v1.9.0 gcr.io/google_containers/kube-apiserver-amd64:v1.9.0
ubuntu@ubuntu-xenial:~$ docker tag gcr.io/google_containers/kube-controller-manager:v1.9.0 gcr.io/google_containers/kube-controller-manager-amd64:v1.9.0
ubuntu@ubuntu-xenial:~$ docker tag gcr.io/google_containers/kube-scheduler:v1.9.0 gcr.io/google_containers/kube-scheduler-amd64:v1.9.0
ubuntu@ubuntu-xenial:~$ docker tag gcr.io/google_containers/kube-proxy:v1.9.0 gcr.io/google_containers/kube-proxy-amd64:v1.9.0
```

## Other images of POD base, etcd, kube-dns, cni flannel

Through docker over ssh Fedora 26
```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "docker images"
vagrant@172.17.4.59's password: 
REPOSITORY                                               TAG                 IMAGE ID            CREATED             SIZE
docker.io/rook/toolbox                                   master              78b8570f3f57        3 days ago          415 MB
docker.io/rook/rook                                      master              8d9558dd80c0        3 days ago          340 MB
gcr.io/google_containers/kube-proxy-amd64                v1.9.0              f6f363e6e98e        3 days ago          109 MB
gcr.io/google_containers/kube-proxy                      v1.9.0              f6f363e6e98e        3 days ago          109 MB
gcr.io/google_containers/cloud-controller-manager        v1.9.0              4c938e6fc795        3 days ago          118 MB
gcr.io/google_containers/kube-apiserver-amd64            v1.9.0              7bff5aa286d7        3 days ago          210 MB
gcr.io/google_containers/kube-apiserver                  v1.9.0              7bff5aa286d7        3 days ago          210 MB
gcr.io/google_containers/kube-controller-manager-amd64   v1.9.0              3bb172f9452c        3 days ago          138 MB
gcr.io/google_containers/kube-controller-manager         v1.9.0              3bb172f9452c        3 days ago          138 MB
gcr.io/google_containers/kube-aggregator                 v1.9.0              e08db9577664        3 days ago          55.8 MB
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
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "docker save gcr.io/google_containers/pause-amd64:3.0" | docker load
vagrant@172.17.4.59's password: 
5f70bf18a086: Loading layer [==================================================>] 1.024 kB/1.024 kB
41ff149e94f2: Loading layer [==================================================>] 748.5 kB/748.5 kB
Loaded image: gcr.io/google_containers/pause-amd64:3.0
```

```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "docker save gcr.io/google_containers/etcd-amd64:3.1.10" | docker load
vagrant@172.17.4.59's password: 
6a749002dd6a: Loading layer [==================================================>] 1.338 MB/1.338 MB
bbd07ea14872: Loading layer [==================================================>] 159.2 MB/159.2 MB
611a3394df5d: Loading layer [==================================================>] 32.44 MB/32.44 MB
Loaded image: gcr.io/google_containers/etcd-amd64:3.1.10
```

```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "docker save gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.7" | docker load
vagrant@172.17.4.59's password: 
5bef08742407: Loading layer [==================================================>] 4.221 MB/4.221 MB
b87261cc1ccb: Loading layer [==================================================>]  2.56 kB/2.56 kB
ac66a5c581a8: Loading layer [==================================================>]   362 kB/362 kB
22f71f461ac8: Loading layer [==================================================>] 3.072 kB/3.072 kB
686a085da152: Loading layer [==================================================>] 36.63 MB/36.63 MB
Loaded image: gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.7
```

```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "docker save gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.7" | docker load
vagrant@172.17.4.59's password: 
cd69fdcd7591: Loading layer [==================================================>] 46.31 MB/46.31 MB
Loaded image: gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.7
```

```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "docker save gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.7" | docker load
vagrant@172.17.4.59's password: 
bd94706d2c63: Loading layer [==================================================>] 38.07 MB/38.07 MB
Loaded image: gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.7
```

CoreDNS
```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "docker save docker.io/coredns/coredns:0.9.10" | docker load
vagrant@172.17.4.59's password: 
034ea01e4aa0: Loading layer [==================================================>] 9.324 MB/9.324 MB
1ea1b16efe91: Loading layer [==================================================>] 31.82 MB/31.82 MB
Loaded image: coredns/coredns:0.9.10
```

Flannel
```
ubuntu@ubuntu-xenial:~$ ssh vagrant@172.17.4.59 "docker save quay.io/coreos/flannel:v0.9.1-amd64" | docker load
vagrant@172.17.4.59's password: 
f439636ab0f0: Loading layer [==================================================>] 6.797 MB/6.797 MB
91b6f6ead101: Loading layer [==================================================>] 4.414 MB/4.414 MB
fc3c053505e6: Loading layer [==================================================>] 34.49 MB/34.49 MB
032657ac7c4a: Loading layer [==================================================>] 2.225 MB/2.225 MB
fd713c7c81af: Loading layer [==================================================>]  5.12 kB/5.12 kB
```

## Install kubeadm from repository mirror

gpg
```
ubuntu@ubuntu-xenial:~$ curl -L http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   663  100   663    0     0  46981      0 --:--:-- --:--:-- --:--:-- 51000
OK
```

repo
```
ubuntu@ubuntu-xenial:~$ curl -L http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/etc0x2Fapt0x2Fsources.list.d/kubernetes.list | sudo tee /etc/apt/sources.list.d/kubernetes.list
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   852  100   852    0     0  75384      0 --:--:-- --:--:-- --:--:-- 85200
### ubuntu
#deb http://apt.kubernetes.io/ kubernetes-trusty main
deb http://apt.kubernetes.io/ kubernetes-xenial main

### --- Modify URL and path first ---
deb [arch=amd64] http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main

deb [arch=amd64] file:///Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main

### debian
#deb http://apt.kubernetes.io/ kubernetes-wheezy main
#deb http://apt.kubernetes.io/ kubernetes-jessie main
#deb http://apt.kubernetes.io/ kubernetes-stretch main
```

__issue__
```
ubuntu@ubuntu-xenial:~$ sudo apt-get --ignore-missing update
Get:1 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Get:1 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Get:2 http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Ign:1 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease
Get:3 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [12.1 kB]
Ign:3 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
Get:3 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [94.4 kB]
Err:3 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
  Could not open file /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages - open (13: Permission denied)
Ign:2 http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease
Ign:4 http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
Err:4 http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
  Could not open file /var/lib/apt/lists/partial/127.0.0.1:48080_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages.gz - open (36: File name too long)
Hit:6 http://archive.ubuntu.com/ubuntu xenial InRelease                                                                                    
Get:7 http://security.ubuntu.com/ubuntu xenial-security InRelease [102 kB]
Hit:8 http://archive.ubuntu.com/ubuntu xenial-updates InRelease                   
Hit:9 http://archive.ubuntu.com/ubuntu xenial-backports InRelease                  
Err:5 https://packages.cloud.google.com/apt kubernetes-xenial InRelease                                                                        
  Connection timed out after 120001 milliseconds
Fetched 111 kB in 2min 0s (920 B/s)
Reading package lists... Done
N: Can't drop privileges for downloading as file '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/InRelease' couldn't be accessed by user '_apt'. - pkgAcquire::Run (13: Permission denied)
W: GPG error: file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease: The following signatures couldn't be verified because the public key is not available: NO_PUBKEY 3746C208A7317B0F
W: The repository 'file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease' is not signed.
N: Data from such a repository can't be authenticated and is therefore potentially dangerous to use.
N: See apt-secure(8) manpage for repository creation and user configuration details.
W: Problem unlinking the file /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages.gz - PrepareFiles (36: File name too long)
W: Problem unlinking the file /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages - PrepareFiles (36: File name too long)
W: Problem unlinking the file /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages - PrepareFiles (36: File name too long)
W: Symlinking file /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages to /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages failed - pkgAcqIndex::StageDownloadDone (36: File name too long)
W: Problem unlinking the file /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages - PrepareFiles (36: File name too long)
W: GPG error: http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease: The following signatures couldn't be verified because the public key is not available: NO_PUBKEY 3746C208A7317B0F
W: The repository 'http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease' is not signed.
N: Data from such a repository can't be authenticated and is therefore potentially dangerous to use.
N: See apt-secure(8) manpage for repository creation and user configuration details.
W: Problem unlinking the file /var/lib/apt/lists/partial/127.0.0.1:48080_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages.gz - PrepareFiles (36: File name too long)
W: Problem unlinking the file /var/lib/apt/lists/partial/127.0.0.1:48080_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages - PrepareFiles (36: File name too long)
W: Failed to fetch http://apt.kubernetes.io/dists/kubernetes-xenial/InRelease  Connection timed out after 120001 milliseconds
E: Failed to fetch store:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages  Could not open file /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages - open (13: Permission denied)
E: Failed to fetch http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages  Could not open file /var/lib/apt/lists/partial/127.0.0.1:48080_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages.gz - open (36: File name too long)
W: Some index files failed to download. They have been ignored, or old ones used instead.
```

Path too long
```
ubuntu@ubuntu-xenial:~$ sudo ln -s /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com /opt/https0x3A0x2F0x2Fpackages.cloud.google.com
```

modify kubernetes.list
```
deb [arch=amd64] file:///opt/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main
```

__issue__
```
ubuntu@ubuntu-xenial:~$ sudo apt-get --ignore-missing update
Get:1 file:/opt/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Get:1 file:/opt/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Get:2 http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Ign:3 http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
Err:3 http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
  Could not open file /var/lib/apt/lists/partial/127.0.0.1:48080_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages.gz - open (36: File name too long)
Hit:4 http://archive.ubuntu.com/ubuntu xenial InRelease                                                                                        
Get:6 http://security.ubuntu.com/ubuntu xenial-security InRelease [102 kB]                                                                     
Get:7 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [102 kB]           
Get:8 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [102 kB]
Err:5 https://packages.cloud.google.com/apt kubernetes-xenial InRelease                                                                        
  Connection timed out after 120001 milliseconds
Fetched 315 kB in 2min 0s (2,613 B/s)
Reading package lists... Done
N: Can't drop privileges for downloading as file '/opt/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/InRelease' couldn't be accessed by user '_apt'. - pkgAcquire::Run (13: Permission denied)
W: Problem unlinking the file /var/lib/apt/lists/partial/127.0.0.1:48080_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages.gz - PrepareFiles (36: File name too long)
W: Problem unlinking the file /var/lib/apt/lists/partial/127.0.0.1:48080_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages - PrepareFiles (36: File name too long)
W: Failed to fetch http://apt.kubernetes.io/dists/kubernetes-xenial/InRelease  Connection timed out after 120001 milliseconds
E: Failed to fetch http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages  Could not open file /var/lib/apt/lists/partial/127.0.0.1:48080_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages.gz - open (36: File name too long)
W: Some index files failed to download. They have been ignored, or old ones used instead.
```

Refer to https://askubuntu.com/questions/771936/permission-error-when-installing-ttf-mscorefonts-installer

restart gofileserver
```
ubuntu@ubuntu-xenial:~$ docker stop gofileserver && docker rm gofileserver && docker run -d --name=gofileserver -v /Users/fanhongling:/mnt -v /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com:/mnt/https0x3A0x2F0x2Fpackages.cloud.google.com -p 48080:48080 docker.io/tangfeixiong/gofileserver
gofileserver
gofileserver
7d9869c596272f2b3755b8b7b77c8956033018eb79b9e476540d3eb86493022f
```

modify kubernetes.list
```
deb [arch=amd64] http://127.0.0.1:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main
```

```
ubuntu@ubuntu-xenial:~$ sudo apt-get --ignore-missing update
Get:1 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Get:1 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Get:2 http://127.0.0.1:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Get:3 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [12.1 kB]
Ign:3 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
Get:3 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [94.4 kB]
Err:3 file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
  Could not open file /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages - open (13: Permission denied)
Get:4 http://127.0.0.1:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [12.1 kB]                    
Get:5 http://security.ubuntu.com/ubuntu xenial-security InRelease [102 kB]                                                
Hit:6 http://archive.ubuntu.com/ubuntu xenial InRelease                                                                   
Get:7 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [102 kB]                                  
Get:9 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [102 kB]                
Err:8 https://packages.cloud.google.com/apt kubernetes-xenial InRelease                                                                        
  Connection timed out after 120000 milliseconds
Fetched 327 kB in 2min 0s (2,707 B/s)
Reading package lists... Done
N: Can't drop privileges for downloading as file '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/InRelease' couldn't be accessed by user '_apt'. - pkgAcquire::Run (13: Permission denied)
W: Problem unlinking the file /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages.gz - PrepareFiles (36: File name too long)
W: Problem unlinking the file /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages - PrepareFiles (36: File name too long)
W: Problem unlinking the file /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages - PrepareFiles (36: File name too long)
W: Symlinking file /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages to /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages failed - pkgAcqIndex::StageDownloadDone (36: File name too long)
W: Problem unlinking the file /var/lib/apt/lists/partial/_Users_fanhongling_Downloads_workspace_src_github.com_tangfeixiong_go-to-kubernetes_https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm_https0x3A0x2F0x2Fpackages.cloud.google.com_apt_dists_kubernetes-xenial_main_binary-amd64_Packages - PrepareFiles (36: File name too long)
W: Failed to fetch http://apt.kubernetes.io/dists/kubernetes-xenial/InRelease  Connection timed out after 120000 milliseconds
E: Failed to fetch store:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages  Could not open file /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages - open (13: Permission denied)
W: Some index files failed to download. They have been ignored, or old ones used instead.
```


```
ubuntu@ubuntu-xenial:~$ sudo apt-cache show kubeadm=1.9.0-00
Package: kubeadm
Version: 1.9.0-00
Installed-Size: 147193
Maintainer: Kubernetes Authors <kubernetes-dev+release@googlegroups.com>
Architecture: amd64
Depends: kubelet (>= 1.6.0), kubectl (>= 1.6.0)
Description: Kubernetes Cluster Bootstrapping Tool
 The Kubernetes command line tool for bootstrapping a Kubernetes cluster.
Description-md5: bb3c7836839894793de38af875e01b30
Homepage: https://kubernetes.io
Filename: pool/kubeadm_1.9.0-00_amd64_191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911.deb
Priority: optional
SHA256: 191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911
Section: misc
Size: 20044720
```

### install 

__issue__
```
ubuntu@ubuntu-xenial:~$ sudo apt install -y kubeadm
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  ebtables kubectl kubelet kubernetes-cni socat
The following NEW packages will be installed:
  ebtables kubeadm kubectl kubelet kubernetes-cni socat
0 upgraded, 6 newly installed, 0 to remove and 8 not upgraded.
Need to get 400 kB/57.3 MB of archives.
After this operation, 413 MB of additional disk space will be used.
Get:1 file:/opt/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubernetes-cni amd64 0.6.0-00 [5,910 kB]
Get:2 file:/opt/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubelet amd64 1.9.0-00 [20.5 MB]
Get:3 file:/opt/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubectl amd64 1.9.0-00 [10.5 MB]
Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 ebtables amd64 2.0.10.4-3.4ubuntu2 [79.4 kB]
Get:5 file:/opt/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.9.0-00 [20.0 MB]
Get:6 http://archive.ubuntu.com/ubuntu xenial/universe amd64 socat amd64 1.7.3.1-1 [321 kB]
Fetched 400 kB in 3s (130 kB/s)
Selecting previously unselected package ebtables.
(Reading database ... 54164 files and directories currently installed.)
Preparing to unpack .../ebtables_2.0.10.4-3.4ubuntu2_amd64.deb ...
Unpacking ebtables (2.0.10.4-3.4ubuntu2) ...
Selecting previously unselected package kubernetes-cni.
Preparing to unpack .../kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb ...
Unpacking kubernetes-cni (0.6.0-00) ...
Selecting previously unselected package socat.
Preparing to unpack .../socat_1.7.3.1-1_amd64.deb ...
Unpacking socat (1.7.3.1-1) ...
Selecting previously unselected package kubelet.
Preparing to unpack .../kubelet_1.9.0-00_amd64_100df9788226fe76ce828503cf24b8c4c6f9705f15504093844c9138e0b2a97f.deb ...
Unpacking kubelet (1.9.0-00) ...
Selecting previously unselected package kubectl.
Preparing to unpack .../kubectl_1.9.0-00_amd64_8ea712c18d89d56090c8753a9630d22fd8ae5cb03d4ec79a1fe6d262c8b4eb36.deb ...
Unpacking kubectl (1.9.0-00) ...
Selecting previously unselected package kubeadm.
Preparing to unpack .../kubeadm_1.9.0-00_amd64_191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911.deb ...
Unpacking kubeadm (1.9.0-00) ...
Processing triggers for systemd (229-4ubuntu21) ...
Processing triggers for ureadahead (0.100.0-19) ...
Processing triggers for man-db (2.7.5-1) ...
Setting up ebtables (2.0.10.4-3.4ubuntu2) ...
update-rc.d: warning: start and stop actions are no longer supported; falling back to defaults
Setting up kubernetes-cni (0.6.0-00) ...
Setting up socat (1.7.3.1-1) ...
Setting up kubelet (1.9.0-00) ...
Setting up kubectl (1.9.0-00) ...
Setting up kubeadm (1.9.0-00) ...
Processing triggers for systemd (229-4ubuntu21) ...
Processing triggers for ureadahead (0.100.0-19) ...
N: Can't drop privileges for downloading as file '/opt/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/pool/kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb' couldn't be accessed by user '_apt'. - pkgAcquire::Run (13: Permission denied)
```

can ignored, or refer to previously restarting gofileserver

```
ubuntu@ubuntu-xenial:~$ sudo apt install -y --allow-unauthenticated kubeadm
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following NEW packages will be installed:
  kubeadm
0 upgraded, 1 newly installed, 0 to remove and 8 not upgraded.
Need to get 20.0 MB of archives.
After this operation, 151 MB of additional disk space will be used.
Get:1 http://127.0.0.1:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.9.0-00 [20.0 MB]
Fetched 20.0 MB in 0s (37.1 MB/s)
Selecting previously unselected package kubeadm.
(Reading database ... 54256 files and directories currently installed.)
Preparing to unpack .../kubeadm_1.9.0-00_amd64.deb ...
Unpacking kubeadm (1.9.0-00) ...
Setting up kubeadm (1.9.0-00) ...
```

```
ubuntu@ubuntu-xenial:~$ sudo cat /lib/systemd/system/kubelet.service 
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
ubuntu@ubuntu-xenial:~$ ls -R /etc/kubernetes/
/etc/kubernetes/:
manifests

/etc/kubernetes/manifests:
```

```
ubuntu@ubuntu-xenial:~$ ls /opt/cni/bin/
bridge  dhcp  flannel  host-local  ipvlan  loopback  macvlan  portmap  ptp  sample  tuning  vlan
```

### copy kubectl admin.conf

kubectl
```
ubuntu@ubuntu-xenial:~$ kubectl version --client
Client Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T21:07:38Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
```

copy
```
ubuntu@ubuntu-xenial:~$ mkdir -p .kube && scp vagrant@172.17.4.59:/home/vagrant/.kube/config .kube
vagrant@172.17.4.59's password: 
config                                                                                                        100% 5451     5.3KB/s   00:00    
```

```
ubuntu@ubuntu-xenial:~$ kubectl version
Client Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T21:07:38Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T20:55:30Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
```

```
ubuntu@ubuntu-xenial:~$ kubectl get nodes
NAME                  STATUS    ROLES     AGE       VERSION
kubedev-172-17-4-59   Ready     master    18h       v1.9.0
```

### join

token
```
ubuntu@ubuntu-xenial:~$ kubeadm token list --kubeconfig=.kube/config
TOKEN                     TTL       EXPIRES                USAGES                   DESCRIPTION                                                EXTRA GROUPS
87bdd5.7aa47e764446c926   5h        2017-12-19T07:43:22Z   authentication,signing   The default bootstrap token generated by 'kubeadm init'.   system:bootstrappers:kubeadm:default-node-token
```

```
ubuntu@ubuntu-xenial:~$ sudo kubeadm join --token 87bdd5.7aa47e764446c926 172.17.4.59:6443 --discovery-token-unsafe-skip-ca-verification
[preflight] Running pre-flight checks.
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

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get nodes
NAME                  STATUS    ROLES     AGE       VERSION
kubedev-172-17-4-55   Ready     <none>    30m       v1.9.0
kubedev-172-17-4-59   Ready     master    19h       v1.9.0
```



__issue__
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get pods -n=kube-system -l app=flannel
NAME                    READY     STATUS             RESTARTS   AGE
kube-flannel-ds-4wzss   0/1       CrashLoopBackOff   11         32m
kube-flannel-ds-vptkd   1/1       Running            0          19h
```

interface name error _enp0s8_  not _ens0p8_, regexp has problem 
```
ubuntu@ubuntu-xenial:~$ kubectl -n kube-system logs kube-flannel-ds-4wzss
I1219 03:18:43.941852       1 main.go:200] Could not find valid interface matching eth1: error looking up interface eth1: route ip+net: no such network interface
I1219 03:18:43.942224       1 main.go:200] Could not find valid interface matching ens0p8: error looking up interface ens0p8: route ip+net: no such network interface
I1219 03:18:43.944234       1 main.go:213] Could not find valid interface matching 172\\.17\\.4\\.[0-9]+: Could not match pattern 172\\.17\\.4\\.[0-9]+ to any of the available network interfaces
E1219 03:18:43.944296       1 main.go:224] Failed to find interface to use that matches the interfaces and/or regexes provided
```

From master
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl delete -f kube-flannel.yaml && kubectl create -f kube-flannel.yaml 
clusterrole "flannel" deleted
clusterrolebinding "flannel" deleted
serviceaccount "flannel" deleted
configmap "kube-flannel-cfg" deleted
daemonset "kube-flannel-ds" deleted
clusterrole "flannel" created
clusterrolebinding "flannel" created
serviceaccount "flannel" created
configmap "kube-flannel-cfg" created
daemonset "kube-flannel-ds" created
```

```
ubuntu@ubuntu-xenial:~$ kubectl -n kube-system get pods -o wide
NAME                                          READY     STATUS    RESTARTS   AGE       IP            NODE
etcd-kubedev-172-17-4-59                      1/1       Running   0          19h       172.17.4.59   kubedev-172-17-4-59
kube-apiserver-kubedev-172-17-4-59            1/1       Running   0          19h       172.17.4.59   kubedev-172-17-4-59
kube-controller-manager-kubedev-172-17-4-59   1/1       Running   0          19h       172.17.4.59   kubedev-172-17-4-59
kube-dns-6f4fd4bdf-r8c7h                      3/3       Running   0          19h       10.244.0.2    kubedev-172-17-4-59
kube-flannel-ds-7pvdh                         1/1       Running   0          19s       172.17.4.59   kubedev-172-17-4-59
kube-flannel-ds-x2l9q                         1/1       Running   0          19s       172.17.4.55   kubedev-172-17-4-55
kube-proxy-ls4nx                              1/1       Running   0          52m       172.17.4.55   kubedev-172-17-4-55
kube-proxy-xj2vb                              1/1       Running   0          19h       172.17.4.59   kubedev-172-17-4-59
kube-scheduler-kubedev-172-17-4-59            1/1       Running   0          19h       172.17.4.59   kubedev-172-17-4-59
```

```
ubuntu@kubedev-172-17-4-55:~$ ip a show cni0
11: cni0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default qlen 1000
    link/ether 0a:58:0a:f4:01:01 brd ff:ff:ff:ff:ff:ff
    inet 10.244.1.1/24 scope global cni0
       valid_lft forever preferred_lft forever
    inet6 fe80::444a:22ff:fefd:7522/64 scope link 
       valid_lft forever preferred_lft forever
```