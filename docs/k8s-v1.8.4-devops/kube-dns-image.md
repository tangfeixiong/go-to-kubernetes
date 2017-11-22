__HELP__

If environment is blocking to access google, Refer to [./kubernetes-pod-base-image.md](./kubernetes-pod-base-image.md) for same idea

__For v1.8.4__

Pull
```
[vagrant@localhost go-to-kubernetes]$ docker pull gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.5
Trying to pull repository gcr.io/google_containers/k8s-dns-kube-dns-amd64 ... 
1.14.5: Pulling from gcr.io/google_containers/k8s-dns-kube-dns-amd64
280aca6ddce2: Pull complete 
a46f95c56b32: Pull complete 
Digest: sha256:1a3fc069de481ae690188f6f1ba4664b5cc7760af37120f70c86505c79eea61d
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.5
```

```
[vagrant@localhost go-to-kubernetes]$ docker pull gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
Trying to pull repository gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64 ... 
1.14.5: Pulling from gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64
280aca6ddce2: Already exists 
5062407518d7: Pull complete 
4e650c6bfcbc: Pull complete 
358bdb7dd123: Pull complete 
62106dacdb76: Pull complete 
Digest: sha256:46b933bb70270c8a02fa6b6f87d440f6f1fce1a5a2a719e164f83f7b109f7544
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
```

```
[vagrant@localhost ~]$ docker pull gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.5
Trying to pull repository gcr.io/google_containers/k8s-dns-sidecar-amd64 ... 
1.14.5: Pulling from gcr.io/google_containers/k8s-dns-sidecar-amd64
280aca6ddce2: Already exists 
44cb8dd9693b: Pull complete 
Digest: sha256:9aab42bf6a2a068b797fe7d91a5d8d915b10dbbc3d6f2b10492848debfba6044
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.5
```

Load from archive
```
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-kube-dns-amd640x3A1.14.5.tar 
a1a7a797fc0e: Loading layer [==================================================>] 45.42 MB/45.42 MB
Loaded image: gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.5
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-dnsmasq-nanny-amd640x3A1.14.5.tar 
e620d0ac6539: Loading layer [==================================================>]  2.56 kB/2.56 kB
9f4f5a30979e: Loading layer [==================================================>]   362 kB/362 kB
fd7319ac31de: Loading layer [==================================================>] 3.072 kB/3.072 kB
b23d166217e1: Loading layer [==================================================>]  37.1 MB/37.1 MB
Loaded image: gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-sidecar-amd640x3A1.14.5.tar 
acfc450a47fa: Loading layer [==================================================>] 37.86 MB/37.86 MB
Loaded image: gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.5
```