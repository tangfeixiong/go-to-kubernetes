# DevOps

POD base
```
[vagrant@kubedev-172-17-4-59 ~]$ /opt/kubernetes/server/bin/kubelet --help 2>&1 | egrep 'gcr.io/google_containers/pause' 
      --pod-infra-container-image string                                                                          The image whose network/ipc namespaces containers in each pod will use. (default "gcr.io/google_containers/pause-amd64:3.0")
```

```
$ img="gcr.io/google_containers/pause-amd64:3.0"; name=${img//\//0x2F}; name=${name//:/0x3A}; echo $name
```

```
[vagrant@kubedev-172-17-4-59 ~]$ docker pull gcr.io/google_containers/pause-amd64:3.0
Trying to pull repository gcr.io/google_containers/pause-amd64 ... 
sha256:163ac025575b775d1c0f9bf0bdd0f086883171eb475b5068e7defa4ca9e76516: Pulling from gcr.io/google_containers/pause-amd64
a3ed95caeb02: Pull complete 
f11233434377: Pull complete 
Digest: sha256:163ac025575b775d1c0f9bf0bdd0f086883171eb475b5068e7defa4ca9e76516
```

Or
```
[vagrant@kubedev-172-17-4-59 ~]$ img="gcr.io/google_containers/pause-amd64:3.0"; name=${img//\//0x2F}; name=${name//:/0x3A}; curl -L http://172.17.4.59:48080/Downloads/99-mirror/docker-images/${name}.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  747k  100  747k    0     0  10.9M      0 --:--:-- --:--:-- --:--:-- 11.0M
Loaded image: gcr.io/google_containers/pause-amd64:3.0
```

Etcd
```
[vagrant@kubedev-172-17-4-59 ~]$ cat /opt/kubernetes/cluster/saltbase/salt/etcd/etcd.manifest | grep "'etcd_version'" | awk -F',' '{print $2}' | cut -c3- | awk -F"'" '{print $1}'
3.1.10
```

```
[vagrant@kubedev-172-17-4-59 ~]$ docker pull gcr.io/google_containers/etcd:3.1.10
Trying to pull repository gcr.io/google_containers/etcd ... 
sha256:28cf78933de29fd26d7a879e51ebd39784cd98109568fd3da61b141257fb85a6: Pulling from gcr.io/google_containers/etcd
aab39f0bc16d: Pull complete 
02bfb792d6f2: Pull complete 
ecd49f05c6c0: Pull complete 
Digest: sha256:28cf78933de29fd26d7a879e51ebd39784cd98109568fd3da61b141257fb85a6
Status: Downloaded newer image for gcr.io/google_containers/etcd:3.1.10
```

kube-dns
```
[vagrant@kubedev-172-17-4-59 ~]$ for i in $(cat /opt/kubernetes/cluster/addons/dns/kube-dns.yaml.base | grep ' image: ' | awk '{print $2}'); do docker pull $i; done
Trying to pull repository gcr.io/google_containers/k8s-dns-kube-dns-amd64 ... 
sha256:f5bddc71efe905f4e4b96f3ca346414be6d733610c1525b98fff808f93966680: Pulling from gcr.io/google_containers/k8s-dns-kube-dns-amd64
280aca6ddce2: Pull complete 
7007640b33be: Pull complete 
Digest: sha256:f5bddc71efe905f4e4b96f3ca346414be6d733610c1525b98fff808f93966680
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.7
Trying to pull repository gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64 ... 
sha256:6cfb9f9c2756979013dbd3074e852c2d8ac99652570c5d17d152e0c0eb3321d6: Pulling from gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64
280aca6ddce2: Already exists 
2fb3ce93f3bf: Pull complete 
2db596bfc1fc: Pull complete 
b013b54fa20b: Pull complete 
7ddc81908397: Pull complete 
Digest: sha256:6cfb9f9c2756979013dbd3074e852c2d8ac99652570c5d17d152e0c0eb3321d6
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.7
Trying to pull repository gcr.io/google_containers/k8s-dns-sidecar-amd64 ... 
sha256:f80f5f9328107dc516d67f7b70054354b9367d31d4946a3bffd3383d83d7efe8: Pulling from gcr.io/google_containers/k8s-dns-sidecar-amd64
280aca6ddce2: Already exists 
90bc6fcbe2bc: Pull complete 
Digest: sha256:f80f5f9328107dc516d67f7b70054354b9367d31d4946a3bffd3383d83d7efe8
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.7
```

CoreDNS
```
[vagrant@kubedev-172-17-4-59 ~]$ for i in $(cat /opt/kubernetes/cluster/addons/dns/coredns.yaml.base | grep ' image: ' | awk '{print $2}'); do docker pull $i; done
Trying to pull repository docker.io/coredns/coredns ... 
sha256:3689b4999726949ac587dbe1d9805720abf19f97ae57baf994028f0054fe792a: Pulling from docker.io/coredns/coredns
88286f41530e: Pull complete 
ab85c1d43d5d: Pull complete 
e793982d5a0e: Pull complete 
Digest: sha256:3689b4999726949ac587dbe1d9805720abf19f97ae57baf994028f0054fe792a
Status: Downloaded newer image for docker.io/coredns/coredns:0.9.10
```

Dashboard
```
[vagrant@kubedev-172-17-4-59 ~]$ img=$(cat /opt/kubernetes/cluster/addons/dashboard/dashboard-controller.yaml | grep ' image: ' | cut -c16-); docker pull $img
Trying to pull repository gcr.io/google_containers/kubernetes-dashboard-amd64 ... 
sha256:71a0de5c6a21cb0c2fbcad71a4fef47acd3e61cd78109822d35e1742f9d8140d: Pulling from gcr.io/google_containers/kubernetes-dashboard-amd64
39e01bcdd352: Pull complete 
Digest: sha256:71a0de5c6a21cb0c2fbcad71a4fef47acd3e61cd78109822d35e1742f9d8140d
Status: Downloaded newer image for gcr.io/google_containers/kubernetes-dashboard-amd64:v1.8.0
```

Flannel, refer to [../coreos0x2Fflannel/](../coreos0x2Fflannel/)
