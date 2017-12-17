# DevOps

## Table of contents
1. Archive of release
1. Images from archive
1. Other images from local file server [](./pod-etcd-dns-and-cni-images.md)

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
