# instruction

Fedora
```
[vagrant@localhost ~]$ cat /etc/fedora-release 
Fedora release 23 (Twenty Three)
```

Docker
```
[vagrant@localhost ~]$ docker version
Client:
 Version:         1.10.3
 API version:     1.22
 Package version: docker-1.10.3-44.git8b7fa4a.fc23.x86_64
 Go version:      go1.5.4
 Git commit:      8b7fa4a/1.10.3
 Built:           
 OS/Arch:         linux/amd64

Server:
 Version:         1.10.3
 API version:     1.22
 Package version: docker-1.10.3-44.git8b7fa4a.fc23.x86_64
 Go version:      go1.5.4
 Git commit:      8b7fa4a/1.10.3
 Built:           
 OS/Arch:         linux/amd64
```

Upgrade docker
```
[vagrant@localhost ~]$ sudo dnf upgrade docker
Last metadata expiration check: 1:05:08 ago on Fri Nov 24 02:55:53 2017.
Dependencies resolved.
====================================================================================================================================
 Package                               Arch                 Version                                     Repository             Size
====================================================================================================================================
Upgrading:
 docker                                x86_64               2:1.10.3-45.gite03ddb8.fc23                 updates               9.2 M
 docker-forward-journald               x86_64               2:1.10.3-45.gite03ddb8.fc23                 updates               993 k
 docker-selinux                        x86_64               2:1.10.3-45.gite03ddb8.fc23                 updates                71 k

Transaction Summary
====================================================================================================================================
Upgrade  3 Packages

Total download size: 10 M
Is this ok [y/N]: y
Downloading Packages:
(1/3): docker-selinux-1.10.3-45.gite03ddb8.fc23.x86_64.rpm                                          9.4 kB/s |  71 kB     00:07    
(2/3): docker-forward-journald-1.10.3-45.gite03ddb8.fc23.x86_64.rpm                                  95 kB/s | 993 kB     00:10    
(3/3): docker-1.10.3-45.gite03ddb8.fc23.x86_64.rpm                                                  441 kB/s | 9.2 MB     00:21    
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                               354 kB/s |  10 MB     00:29     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Upgrading   : docker-selinux-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                               1/6 
  Upgrading   : docker-forward-journald-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                      2/6 
  Upgrading   : docker-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                       3/6 
  Cleanup     : docker-2:1.10.3-44.git8b7fa4a.fc23.x86_64                                                                       4/6 
  Cleanup     : docker-selinux-2:1.10.3-44.git8b7fa4a.fc23.x86_64                                                               5/6 
  Cleanup     : docker-forward-journald-2:1.10.3-44.git8b7fa4a.fc23.x86_64                                                      6/6 
  Verifying   : docker-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                       1/6 
  Verifying   : docker-forward-journald-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                      2/6 
  Verifying   : docker-selinux-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                               3/6 
  Verifying   : docker-2:1.10.3-44.git8b7fa4a.fc23.x86_64                                                                       4/6 
  Verifying   : docker-forward-journald-2:1.10.3-44.git8b7fa4a.fc23.x86_64                                                      5/6 
  Verifying   : docker-selinux-2:1.10.3-44.git8b7fa4a.fc23.x86_64                                                               6/6 

Upgraded:
  docker.x86_64 2:1.10.3-45.gite03ddb8.fc23                    docker-forward-journald.x86_64 2:1.10.3-45.gite03ddb8.fc23           
  docker-selinux.x86_64 2:1.10.3-45.gite03ddb8.fc23           

Complete!
```

Current repositories
```
[vagrant@localhost ~]$ ls -1 /etc/yum.repos.d/
dl.google.com_linux_chrome_rpm_stable_x86_64.repo
fedora.repo
fedora-updates.repo
fedora-updates-testing.repo
google-chrome.repo
openshift-rhel7-dependencies.repo
```

## Prerequiestes

### firewall and selinux

fireall
```
[vagrant@localhost ~]$ sudo systemctl is-active firewalld.service
unknown
```

selinux
```
[vagrant@localhost ~]$ sudo getenforce
Permissive
```

### hostname and hosts

hostname
```
[vagrant@localhost ~]$ echo "kubedev-10-64-33-199" | sudo tee /etc/hostname 
kubedev-10-64-33-199
```

```
[vagrant@localhost ~]$ sudo hostname --file /etc/hostname
```

hosts
```
[vagrant@localhost ~]$ hostname
kubedev-10-64-33-199
```

```
[vagrant@localhost ~]$ echo "10.64.33.199    kubedev-10-64-33-199" | sudo tee -a /etc/hosts
10.64.33.199    kubedev-10-64-33-199
[vagrant@localhost ~]$ cat /etc/hosts
127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
::1         localhost localhost.localdomain localhost6 localhost6.localdomain6
10.64.33.199    kubedev-10-64-33-199
```

## Install kubeadm

### local kubernetes repository

files and local http
```
[vagrant@localhost ~]$ echo " 
> [kubernetes]
> name=Kubernetes
> #baseurl=http://yum.kubernetes.io/repos/kubernetes-el7-x86_64
> baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
> enabled=1
> gpgcheck=1
> repo_gpgcheck=1
> gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
>         https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
> 
> [kubernetes-offline]
> name=Kubernetes-offline
> baseurl=file:///vagrant_drive_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
> enabled=0
> gpgcheck=0
> 
> [kubernetes-10-64-33-82]
> name=Kubernetes-10-64-33-82
> baseurl=http://10.64.33.82:48080/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
> enabled=0
> gpgcheck=0
> " | sudo tee /etc/yum.repos.d/kubernetes.repo
 
[kubernetes]
name=Kubernetes
#baseurl=http://yum.kubernetes.io/repos/kubernetes-el7-x86_64
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
        https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg

[kubernetes-offline]
name=Kubernetes-offline
baseurl=file:///vagrant_drive_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=0
gpgcheck=0

[kubernetes-10-64-33-82]
name=Kubernetes-10-64-33-82
baseurl=http://10.64.33.82:48080/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=0
gpgcheck=0

```

```
[vagrant@localhost ~]$ sudo dnf --disablerepo=kubernetes --enablerepo=kubernetes-\* repolist
Kubernetes-10-64-33-82                                                                              319 kB/s |  18 kB     00:00    
Kubernetes-offline                                                                                   11 MB/s |  18 kB     00:00    
repo id                                              repo name                                                                status
dl.google.com_linux_chrome_rpm_stable_x86_64         added from: http://dl.google.com/linux/chrome/rpm/stable/x86_64               3
*fedora                                              Fedora 23 - x86_64                                                       46,074
google-chrome                                        google-chrome                                                                 3
kubernetes-10-64-33-82                               Kubernetes-10-64-33-82                                                       84
kubernetes-offline                                   Kubernetes-offline                                                           84
origin-deps-rhel7                                    OpenShift Origin Dependency repo for RHEL 7                                   6
*updates                                             Fedora 23 - x86_64 - Updates                                             21,532
```

### install kubeadm from local repository

Using local repo
```
[vagrant@localhost ~]$ sudo dnf --disablerepo=kubernetes --enablerepo=kubernetes-\* install kubeadm
Last metadata expiration check: 0:00:48 ago on Fri Nov 24 04:14:08 2017.
Dependencies resolved.
====================================================================================================================================
 Package                          Arch                     Version                   Repository                                Size
====================================================================================================================================
Installing:
 kubeadm                          x86_64                   1.8.4-0                   kubernetes-10-64-33-82                    15 M
 kubectl                          x86_64                   1.8.4-0                   kubernetes-10-64-33-82                   7.3 M
 kubelet                          x86_64                   1.8.4-0                   kubernetes-10-64-33-82                    16 M
 kubernetes-cni                   x86_64                   0.5.1-1                   kubernetes-10-64-33-82                   7.4 M

Transaction Summary
====================================================================================================================================
Install  4 Packages

Total download size: 45 M
Installed size: 242 M
Is this ok [y/N]: y
Downloading Packages:
(1/4): a9db28728641ddbf7f025b8b496804d82a396d0ccb178fffd124623fb2f999ea-kubectl-1.8.4-0.x86_64.rpm   13 MB/s | 7.3 MB     00:00    
(2/4): aeaad1e283c54876b759a089f152228d7cd4c049f271125c23623995b8e76f96-kubeadm-1.8.4-0.x86_64.rpm   12 MB/s |  15 MB     00:01    
(3/4): 79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_  10 MB/s | 7.4 MB     00:00    
(4/4): 1acca81eb5cf99453f30466876ff03146112b7f12c625cb48f12508684e02665-kubelet-1.8.4-0.x86_64.rpm   11 MB/s |  16 MB     00:01    
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                31 MB/s |  45 MB     00:01     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Installing  : kubernetes-cni-0.5.1-1.x86_64                                                                                   1/4 
  Installing  : kubelet-1.8.4-0.x86_64                                                                                          2/4 
  Installing  : kubectl-1.8.4-0.x86_64                                                                                          3/4 
  Installing  : kubeadm-1.8.4-0.x86_64                                                                                          4/4 
  Verifying   : kubeadm-1.8.4-0.x86_64                                                                                          1/4 
  Verifying   : kubectl-1.8.4-0.x86_64                                                                                          2/4 
  Verifying   : kubelet-1.8.4-0.x86_64                                                                                          3/4 
  Verifying   : kubernetes-cni-0.5.1-1.x86_64                                                                                   4/4 

Installed:
  kubeadm.x86_64 1.8.4-0         kubectl.x86_64 1.8.4-0         kubelet.x86_64 1.8.4-0         kubernetes-cni.x86_64 0.5.1-1        

Complete!
```

### installed

kubelet
```
[vagrant@localhost ~]$ which kubelet
/bin/kubelet
```

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
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS
```

```
[vagrant@localhost ~]$ sudo systemctl is-active kubelet.service
unknown
```

```
[vagrant@localhost ~]$ sudo systemctl is-active kubelet.service
unknown
[vagrant@localhost ~]$ sudo systemctl status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: inactive (dead)
     Docs: http://kubernetes.io/docs/
```

kubectl
```
[vagrant@localhost ~]$ kubectl version --client
Client Version: version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.4", GitCommit:"9befc2b8928a9426501d3bf62f72849d5cbcd5a3", GitTreeState:"clean", BuildDate:"2017-11-20T05:28:34Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
```

cni
```
[vagrant@localhost ~]$ ls /opt/cni/bin/
bridge  cnitool  dhcp  flannel  host-local  ipvlan  loopback  macvlan  noop  ptp  tuning
```

## Load docker images

Via [http file server on previous node](../k8s-v1.8.2-devops)

### kubernetes

kube-apiserver
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/home/vagrant/kubernetes/server/bin/kube-apiserver.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  185M  100  185M    0     0  79.6M      0  0:00:02  0:00:02 --:--:-- 79.6M
```

```
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/kube-apiserver:v1.8.4 gcr.io/google_containers/kube-apiserver-amd64:v1.8.4
```

kube-controller-manager
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/home/vagrant/kubernetes/server/bin/kube-controller-manager.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  123M  100  123M    0     0  72.6M      0  0:00:01  0:00:01 --:--:-- 72.6M
```

```
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/kube-controller-manager:v1.8.4 gcr.io/google_containers/kube-controller-manager-amd64:v1.8.4
```

kube-scheduler
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/home/vagrant/kubernetes/server/bin/kube-scheduler.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 52.6M  100 52.6M    0     0  60.0M      0 --:--:-- --:--:-- --:--:-- 60.0M
```

```
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/kube-scheduler:v1.8.4 gcr.io/google_containers/kube-scheduler-amd64:v1.8.4
```

kube-proxy
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/home/vagrant/kubernetes/server/bin/kube-proxy.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.6M  100 90.6M    0     0  98.6M      0 --:--:-- --:--:-- --:--:-- 98.6M
```

```
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/kube-proxy:v1.8.4 gcr.io/google_containers/kube-proxy-amd64:v1.8.4
```

### etcd

version 3.0.17
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fetcd0x2E3.0.17.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  161M  100  161M    0     0  55.1M      0  0:00:02  0:00:02 --:--:-- 55.1M
```

```
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/etcd:3.0.17 gcr.io/google_containers/etcd-amd64:3.0.17
```

### addons

dns version 1.14.5
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-dnsmasq-nanny-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 39.7M  100 39.7M    0     0  28.9M      0  0:00:01  0:00:01 --:--:-- 28.9M
```

```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-kube-dns-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 47.3M  100 47.3M    0     0  40.7M      0  0:00:01  0:00:01 --:--:-- 40.7M
```

```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-sidecar-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 40.1M  100 40.1M    0     0  49.4M      0 --:--:-- --:--:-- --:--:-- 49.4M
```

### Pod base

pause v3
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  745k  100  745k    0     0  6151k      0 --:--:-- --:--:-- --:--:-- 6161k
Error response from daemon: invalid argument
```

```
[vagrant@localhost ~]$ docker images
REPOSITORY                                               TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver-amd64            v1.8.4              10a052dccbc5        3 days ago          194.4 MB
gcr.io/google_containers/kube-apiserver                  v1.8.4              10a052dccbc5        3 days ago          194.4 MB
gcr.io/google_containers/kube-controller-manager-amd64   v1.8.4              7058ac4d4af5        3 days ago          129.3 MB
gcr.io/google_containers/kube-controller-manager         v1.8.4              7058ac4d4af5        3 days ago          129.3 MB
gcr.io/google_containers/kube-proxy-amd64                v1.8.4              65a61c14e8c2        3 days ago          93.2 MB
gcr.io/google_containers/kube-proxy                      v1.8.4              65a61c14e8c2        3 days ago          93.2 MB
gcr.io/google_containers/kube-scheduler-amd64            v1.8.4              0d985fed7f95        3 days ago          54.98 MB
gcr.io/google_containers/kube-scheduler                  v1.8.4              0d985fed7f95        3 days ago          54.98 MB
gcr.io/google_containers/k8s-dns-sidecar-amd64           1.14.5              fed89e8b4248        8 weeks ago         41.81 MB
gcr.io/google_containers/k8s-dns-kube-dns-amd64          1.14.5              512cd7425a73        8 weeks ago         49.38 MB
gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64     1.14.5              459944ce8cc4        8 weeks ago         41.42 MB
gcr.io/google_containers/etcd-amd64                      3.0.17              243830dae7dd        9 months ago        168.9 MB
gcr.io/google_containers/etcd                            3.0.17              243830dae7dd        9 months ago        168.9 MB
<none>                                                   <none>              99e59f495ffa        18 months ago       746.9 kB
```

```
[vagrant@localhost ~]$ docker tag 99e59f495ffa gcr.io/google_containers/pause-amd64:3.0
```

### network

coreos flannel
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/quay.io0x2Fcoreos0x2Fflannel0x3Av0.9.0-amd64.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 49.7M  100 49.7M    0     0  48.9M      0  0:00:01  0:00:01 --:--:-- 49.0M
```

weaveworks weave
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-kube0x3A2.1.1.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.8M  100 90.8M    0     0  61.6M      0  0:00:01  0:00:01 --:--:-- 61.6M
```

```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/vagrant_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-npc0x3A2.1.1.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 44.7M  100 44.7M    0     0  55.6M      0 --:--:-- --:--:-- --:--:-- 55.7M
```

## Copy manifests from previous http file server

flannel
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/home/vagrant/kube-flannel.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2849  100  2849    0     0   107k      0 --:--:-- --:--:-- --:--:--  111k
```

weave
```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/home/vagrant/weave-daemonset-k8s-1.6.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  4579  100  4579    0     0   127k      0 --:--:-- --:--:-- --:--:--  131k
```

```
[vagrant@localhost ~]$ curl http://10.64.33.82:48080/home/vagrant/weave-daemonset-k8s-1.7.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  5064  100  5064    0     0   360k      0 --:--:-- --:--:-- --:--:--  380k
```

## Deploy k8s 

networks
```
[vagrant@localhost ~]$ ip a show eth1; ip a show eth2
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:3f:88:a3 brd ff:ff:ff:ff:ff:ff
    inet 10.64.33.199/24 brd 10.64.33.255 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe3f:88a3/64 scope link 
       valid_lft forever preferred_lft forever
4: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:27:95:e9 brd ff:ff:ff:ff:ff:ff
    inet 172.28.128.3/24 brd 172.28.128.255 scope global dynamic eth2
       valid_lft 1092sec preferred_lft 1092sec
    inet6 fe80::a00:27ff:fe27:95e9/64 scope link 
       valid_lft forever preferred_lft forever
```

### start kubelet service

Must manually start service in Fedora23
```
[vagrant@localhost ~]$ sudo systemctl start kubelet.service
```

Because not initialized
```
[vagrant@localhost ~]$ sudo systemctl status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: activating (auto-restart) (Result: exit-code) since Fri 2017-11-24 05:22:48 UTC; 7s ago
     Docs: http://kubernetes.io/docs/
  Process: 22196 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS (code=exited, status=1/FAILURE)
 Main PID: 22196 (code=exited, status=1/FAILURE)

Nov 24 05:22:48 kubedev-10-64-33-199 systemd[1]: kubelet.service: Main process exited, code=exited, status=1/FAILURE
Nov 24 05:22:48 kubedev-10-64-33-199 systemd[1]: kubelet.service: Unit entered failed state.
Nov 24 05:22:48 kubedev-10-64-33-199 systemd[1]: kubelet.service: Failed with result 'exit-code'.
```

Service maybe restarted in 10 seconds
```
Nov 24 05:23:39 kubedev-10-64-33-199 systemd[1]: kubelet.service: Service hold-off time over, scheduling restart.
Nov 24 05:23:39 kubedev-10-64-33-199 systemd[1]: Started kubelet: The Kubernetes Node Agent.
Nov 24 05:23:39 kubedev-10-64-33-199 systemd[1]: Starting kubelet: The Kubernetes Node Agent...
Nov 24 05:23:40 kubedev-10-64-33-199 kubelet[22224]: I1124 05:23:40.111323   22224 feature_gate.go:156] feature gates: map[]
Nov 24 05:23:40 kubedev-10-64-33-199 kubelet[22224]: I1124 05:23:40.111618   22224 controller.go:114] kubelet config controller: starting controller
Nov 24 05:23:40 kubedev-10-64-33-199 kubelet[22224]: I1124 05:23:40.111756   22224 controller.go:118] kubelet config controller: validating combination of defaults and flags
Nov 24 05:23:40 kubedev-10-64-33-199 kubelet[22224]: error: unable to load client CA file /etc/kubernetes/pki/ca.crt: open /etc/kubernetes/pki/ca.crt: no such file or directory
Nov 24 05:23:40 kubedev-10-64-33-199 systemd[1]: kubelet.service: Main process exited, code=exited, status=1/FAILURE
Nov 24 05:23:40 kubedev-10-64-33-199 systemd[1]: kubelet.service: Unit entered failed state.
Nov 24 05:23:40 kubedev-10-64-33-199 systemd[1]: kubelet.service: Failed with result 'exit-code'.
```

Init with kubeadm
```
[vagrant@localhost ~]$ sudo kubeadm init --apiserver-advertise-address 10.64.33.199 --apiserver-bind-port 443 --apiserver-cert-extra-sans 10.64.33.199,172.28.128.3 --pod-network-cidr 10.244.0.0/16 --kubernetes-version v1.8.4 --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.8.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [kubedev-10-64-33-199 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.64.33.199 10.64.33.199 172.28.128.3]
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
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz/syncloop' failed with error: Get http://localhost:10255/healthz/syncloop: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz/syncloop' failed with error: Get http://localhost:10255/healthz/syncloop: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz/syncloop' failed with error: Get http://localhost:10255/healthz/syncloop: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz/syncloop' failed with error: Get http://localhost:10255/healthz/syncloop: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.

Unfortunately, an error has occurred:
	timed out waiting for the condition

This error is likely caused by that:
	- The kubelet is not running
	- The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)
	- There is no internet connection; so the kubelet can't pull the following control plane images:
		- gcr.io/google_containers/kube-apiserver-amd64:v1.8.4
		- gcr.io/google_containers/kube-controller-manager-amd64:v1.8.4
		- gcr.io/google_containers/kube-scheduler-amd64:v1.8.4

You can troubleshoot this for example with the following commands if you're on a systemd-powered system:
	- 'systemctl status kubelet'
	- 'journalctl -xeu kubelet'
couldn't initialize a Kubernetes cluster
```

Issue: same as CentOS 7.2, swap not supported
```
Nov 24 05:29:44 kubedev-10-64-33-199 systemd[1]: kubelet.service: Service hold-off time over, scheduling restart.
Nov 24 05:29:44 kubedev-10-64-33-199 systemd[1]: Started kubelet: The Kubernetes Node Agent.
Nov 24 05:29:44 kubedev-10-64-33-199 systemd[1]: Starting kubelet: The Kubernetes Node Agent...
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.325812   22829 feature_gate.go:156] feature gates: map[]
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.326107   22829 controller.go:114] kubelet config controller: starting controller
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.326234   22829 controller.go:118] kubelet config controller: validating combination of defaults and flags
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.330695   22829 client.go:75] Connecting to docker on unix:///var/run/docker.sock
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.330841   22829 client.go:95] Start docker client with request timeout=2m0s
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: W1124 05:29:44.341021   22829 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.361294   22829 feature_gate.go:156] feature gates: map[]
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: W1124 05:29:44.361561   22829 server.go:289] --cloud-provider=auto-detect is deprecated. The desired cloud provider should be set explicitly
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.378397   22829 certificate_manager.go:361] Requesting new certificate.
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: E1124 05:29:44.378791   22829 certificate_manager.go:284] Failed while requesting a signed certificate from the master: cannot create certificate signing request: Post https://10.64.33.199:443/apis/certificates.k8s.io/v1beta1/certificatesigningrequests: dial tcp 10.64.33.199:443: getsockopt: connection refused
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.379473   22829 manager.go:149] cAdvisor running in container: "/sys/fs/cgroup/cpu,cpuacct/system.slice/kubelet.service"
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: W1124 05:29:44.488615   22829 manager.go:157] unable to connect to Rkt api service: rkt: cannot tcp Dial rkt api service: dial tcp [::1]:15441: getsockopt: connection refused
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: W1124 05:29:44.488660   22829 manager.go:166] unable to connect to CRI-O api service: Get http://%2Fvar%2Frun%2Fcrio.sock/info: dial unix /var/run/crio.sock: connect: no such file or directory
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.596510   22829 fs.go:139] Filesystem UUIDs: map[79d1f0d8-29ec-443f-a26f-3ed5d754e024:/dev/dm-1 c159b294-c140-4a5c-a57d-ec5ec9199fe0:/dev/sda1 cc7b1e9a-86ec-4771-bd2a-650ed6dee8d4:/dev/dm-0 e88c9b60-eb5e-49b7-8cdb-3e228484090a:/dev/dm-8 6a582259-0c00-445a-93bd-d61f458bb287:/dev/dm-4]
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.596719   22829 fs.go:140] Filesystem partitions: map[/dev/mapper/vg_vagrant-lv_root:{mountpoint:/ major:253 minor:0 fsType:ext4 blockSize:0} /dev/sda1:{mountpoint:/boot major:8 minor:1 fsType:ext4 blockSize:0} /dev/mapper/vg_vagrant-openshift--xfs--vol--dir:{mountpoint:/mnt/openshift-xfs-vol-dir major:253 minor:4 fsType:xfs blockSize:0} shm:{mountpoint:/var/lib/docker/containers/2bd0c057ae94f73cce613fbca4f49e46f3d8364292810e18ac086a1ecebd1122/shm major:0 minor:40 fsType:tmpfs blockSize:0} docker-253:0-1311012-pool:{mountpoint: major:253 minor:2 fsType:devicemapper blockSize:128} tmpfs:{mountpoint:/dev/shm major:0 minor:19 fsType:tmpfs blockSize:0}]
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.602555   22829 manager.go:216] Machine: {NumCores:1 CpuFrequency:3393181 MemoryCapacity:3155668992 HugePages:[{PageSize:2048 NumPages:0}] MachineID:5c949bb3146241e09f7e671b0704d4fb SystemUUID:2F77DA33-8080-412E-B2A5-BCE1A2305391 BootID:9a8368b9-fb38-4497-9d1f-a904bb329483 Filesystems:[{Device:tmpfs DeviceMajor:0 DeviceMinor:19 Capacity:1577832448 Type:vfs Inodes:385213 HasInodes:true} {Device:/dev/mapper/vg_vagrant-lv_root DeviceMajor:253 DeviceMinor:0 Capacity:40522645504 Type:vfs Inodes:2523136 HasInodes:true} {Device:/dev/sda1 DeviceMajor:8 DeviceMinor:1 Capacity:499355648 Type:vfs Inodes:128016 HasInodes:true} {Device:/dev/mapper/vg_vagrant-openshift--xfs--vol--dir DeviceMajor:253 DeviceMinor:4 Capacity:10693378048 Type:vfs Inodes:5226496 HasInodes:true} {Device:shm DeviceMajor:0 DeviceMinor:40 Capacity:67108864 Type:vfs Inodes:385213 HasInodes:true} {Device:docker-253:0-1311012-pool DeviceMajor:253 DeviceMinor:2 Capacity:30064771072 Type:devicemapper Inodes:0 HasInodes:false}] DiskMap:map[253:0:{Name:dm-0 Major:253 Minor:0 Size:41305505792 Scheduler:none} 253:3:{Name:dm-3 Major:253 Minor:3 Size:2181038080 Scheduler:none} 253:6:{Name:dm-6 Major:253 Minor:6 Size:10737418240 Scheduler:none} 253:7:{Name:dm-7 Major:253 Minor:7 Size:10737418240 Scheduler:none} 253:8:{Name:dm-8 Major:253 Minor:8 Size:10737418240 Scheduler:none} 8:0:{Name:sda Major:8 Minor:0 Size:85899345920 Scheduler:cfq} 253:1:{Name:dm-1 Major:253 Minor:1 Size:1107296256 Scheduler:none} 253:2:{Name:dm-2 Major:253 Minor:2 Size:30064771072 Scheduler:none} 253:4:{Name:dm-4 Major:253 Minor:4 Size:10703863808 Scheduler:none} 253:5:{Name:dm-5 Major:253 Minor:5 Size:30064771072 Scheduler:none}] NetworkDevices:[{Name:br-4dde2dd32915 MacAddress:02:42:c2:8f:d9:c4 Speed:0 Mtu:1500} {Name:eth0 MacAddress:08:00:27:24:23:96 Speed:1000 Mtu:1500} {Name:eth1 MacAddress:08:00:27:3f:88:a3 Speed:1000 Mtu:1500} {Name:eth2 MacAddress:08:00:27:27:95:e9 Speed:1000 Mtu:1500}] Topology:[{Id:0 Memory:3155668992 Cores:[{Id:0 Threads:[0] Caches:[{Size:32768 Typ
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: e:Data Level:1} {Size:32768 Type:Instruction Level:1} {Size:262144 Type:Unified Level:2}]}] Caches:[{Size:6291456 Type:Unified Level:3}]}] CloudProvider:Unknown InstanceType:Unknown InstanceID:None}
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.619600   22829 manager.go:222] Version: {KernelVersion:4.2.3-300.fc23.x86_64 ContainerOsVersion:Fedora 23 (Twenty Three) DockerVersion:1.10.3 DockerAPIVersion:1.22 CadvisorVersion: CadvisorRevision:}
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: I1124 05:29:44.619833   22829 server.go:422] --cgroups-per-qos enabled, but --cgroup-root was not specified.  defaulting to /
Nov 24 05:29:44 kubedev-10-64-33-199 kubelet[22829]: error: failed to run Kubelet: Running with swap on is not supported, please disable swap! or set --fail-swap-on flag to false. /proc/swaps contained: [Filename                                Type                Size        Used        Priority /dev/dm-1                               partition        1081340        0        -1]
Nov 24 05:29:44 kubedev-10-64-33-199 systemd[1]: kubelet.service: Main process exited, code=exited, status=1/FAILURE
Nov 24 05:29:44 kubedev-10-64-33-199 systemd[1]: kubelet.service: Unit entered failed state.
Nov 24 05:29:44 kubedev-10-64-33-199 systemd[1]: kubelet.service: Failed with result 'exit-code'.
```

```
[vagrant@localhost ~]$ sudo sed -i 's/\(ExecStart=.*ARGS\)/\1 --fail-swap-on=false/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

```
[vagrant@localhost ~]$ sudo systemctl daemon-reload 
```

```
[vagrant@localhost ~]$ sudo systemctl restart kubelet.service
```

```
[vagrant@localhost ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: activating (auto-restart) (Result: exit-code) since Fri 2017-11-24 05:31:58 UTC; 7s ago
     Docs: http://kubernetes.io/docs/
  Process: 23182 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false (code=exited, status=1/FAILURE)
 Main PID: 23182 (code=exited, status=1/FAILURE)

Nov 24 05:31:58 kubedev-10-64-33-199 kubelet[23182]: error: failed to run Kubelet: failed to create kubelet: misconfiguration: kubelet cgroup driver: "systemd" is different from docker cgroup driver: "cgroupfs"
Nov 24 05:31:58 kubedev-10-64-33-199 systemd[1]: kubelet.service: Main process exited, code=exited, status=1/FAILURE
Nov 24 05:31:58 kubedev-10-64-33-199 systemd[1]: kubelet.service: Unit entered failed state.
Nov 24 05:31:58 kubedev-10-64-33-199 systemd[1]: kubelet.service: Failed with result 'exit-code'.
```

Issue: in docker version 1.10, cgroup driver is "cgroupfs", not "systemd"
```
Nov 24 05:33:01 kubedev-10-64-33-199 systemd[1]: kubelet.service: Service hold-off time over, scheduling restart.
Nov 24 05:33:01 kubedev-10-64-33-199 systemd[1]: Started kubelet: The Kubernetes Node Agent.
Nov 24 05:33:01 kubedev-10-64-33-199 systemd[1]: Starting kubelet: The Kubernetes Node Agent...
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.077280   23369 feature_gate.go:156] feature gates: map[]
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.077796   23369 controller.go:114] kubelet config controller: starting controller
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.077933   23369 controller.go:118] kubelet config controller: validating combination of defaults and flags
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.082300   23369 client.go:75] Connecting to docker on unix:///var/run/docker.sock
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.082390   23369 client.go:95] Start docker client with request timeout=2m0s
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.092536   23369 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.112461   23369 feature_gate.go:156] feature gates: map[]
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.112727   23369 server.go:289] --cloud-provider=auto-detect is deprecated. The desired cloud provider should be set explicitly
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.129189   23369 certificate_manager.go:361] Requesting new certificate.
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: E1124 05:33:02.129756   23369 certificate_manager.go:284] Failed while requesting a signed certificate from the master: cannot create certificate signing request: Post https://10.64.33.199:443/apis/certificates.k8s.io/v1beta1/certificatesigningrequests: dial tcp 10.64.33.199:443: getsockopt: connection refused
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.130233   23369 manager.go:149] cAdvisor running in container: "/sys/fs/cgroup/cpu,cpuacct/system.slice/kubelet.service"
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.239521   23369 manager.go:157] unable to connect to Rkt api service: rkt: cannot tcp Dial rkt api service: dial tcp [::1]:15441: getsockopt: connection refused
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.239590   23369 manager.go:166] unable to connect to CRI-O api service: Get http://%2Fvar%2Frun%2Fcrio.sock/info: dial unix /var/run/crio.sock: connect: no such file or directory
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.357334   23369 fs.go:139] Filesystem UUIDs: map[6a582259-0c00-445a-93bd-d61f458bb287:/dev/dm-4 79d1f0d8-29ec-443f-a26f-3ed5d754e024:/dev/dm-1 c159b294-c140-4a5c-a57d-ec5ec9199fe0:/dev/sda1 cc7b1e9a-86ec-4771-bd2a-650ed6dee8d4:/dev/dm-0 e88c9b60-eb5e-49b7-8cdb-3e228484090a:/dev/dm-8]
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.357574   23369 fs.go:140] Filesystem partitions: map[tmpfs:{mountpoint:/dev/shm major:0 minor:19 fsType:tmpfs blockSize:0} /dev/mapper/vg_vagrant-lv_root:{mountpoint:/ major:253 minor:0 fsType:ext4 blockSize:0} /dev/sda1:{mountpoint:/boot major:8 minor:1 fsType:ext4 blockSize:0} /dev/mapper/vg_vagrant-openshift--xfs--vol--dir:{mountpoint:/mnt/openshift-xfs-vol-dir major:253 minor:4 fsType:xfs blockSize:0} shm:{mountpoint:/var/lib/docker/containers/2bd0c057ae94f73cce613fbca4f49e46f3d8364292810e18ac086a1ecebd1122/shm major:0 minor:40 fsType:tmpfs blockSize:0} docker-253:0-1311012-pool:{mountpoint: major:253 minor:2 fsType:devicemapper blockSize:128}]
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.361676   23369 manager.go:216] Machine: {NumCores:1 CpuFrequency:3393181 MemoryCapacity:3155668992 HugePages:[{PageSize:2048 NumPages:0}] MachineID:5c949bb3146241e09f7e671b0704d4fb SystemUUID:2F77DA33-8080-412E-B2A5-BCE1A2305391 BootID:9a8368b9-fb38-4497-9d1f-a904bb329483 Filesystems:[{Device:/dev/mapper/vg_vagrant-lv_root DeviceMajor:253 DeviceMinor:0 Capacity:40522645504 Type:vfs Inodes:2523136 HasInodes:true} {Device:/dev/sda1 DeviceMajor:8 DeviceMinor:1 Capacity:499355648 Type:vfs Inodes:128016 HasInodes:true} {Device:/dev/mapper/vg_vagrant-openshift--xfs--vol--dir DeviceMajor:253 DeviceMinor:4 Capacity:10693378048 Type:vfs Inodes:5226496 HasInodes:true} {Device:shm DeviceMajor:0 DeviceMinor:40 Capacity:67108864 Type:vfs Inodes:385213 HasInodes:true} {Device:docker-253:0-1311012-pool DeviceMajor:253 DeviceMinor:2 Capacity:30064771072 Type:devicemapper Inodes:0 HasInodes:false} {Device:tmpfs DeviceMajor:0 DeviceMinor:19 Capacity:1577832448 Type:vfs Inodes:385213 HasInodes:true}] DiskMap:map[253:3:{Name:dm-3 Major:253 Minor:3 Size:2181038080 Scheduler:none} 253:4:{Name:dm-4 Major:253 Minor:4 Size:10703863808 Scheduler:none} 253:6:{Name:dm-6 Major:253 Minor:6 Size:10737418240 Scheduler:none} 253:7:{Name:dm-7 Major:253 Minor:7 Size:10737418240 Scheduler:none} 8:0:{Name:sda Major:8 Minor:0 Size:85899345920 Scheduler:cfq} 253:0:{Name:dm-0 Major:253 Minor:0 Size:41305505792 Scheduler:none} 253:1:{Name:dm-1 Major:253 Minor:1 Size:1107296256 Scheduler:none} 253:8:{Name:dm-8 Major:253 Minor:8 Size:10737418240 Scheduler:none} 253:2:{Name:dm-2 Major:253 Minor:2 Size:30064771072 Scheduler:none} 253:5:{Name:dm-5 Major:253 Minor:5 Size:30064771072 Scheduler:none}] NetworkDevices:[{Name:br-4dde2dd32915 MacAddress:02:42:c2:8f:d9:c4 Speed:0 Mtu:1500} {Name:eth0 MacAddress:08:00:27:24:23:96 Speed:1000 Mtu:1500} {Name:eth1 MacAddress:08:00:27:3f:88:a3 Speed:1000 Mtu:1500} {Name:eth2 MacAddress:08:00:27:27:95:e9 Speed:1000 Mtu:1500}] Topology:[{Id:0 Memory:3155668992 Cores:[{Id:0 Threads:[0] Caches:[{Size:32768 Typ
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: e:Data Level:1} {Size:32768 Type:Instruction Level:1} {Size:262144 Type:Unified Level:2}]}] Caches:[{Size:6291456 Type:Unified Level:3}]}] CloudProvider:Unknown InstanceType:Unknown InstanceID:None}
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.376118   23369 manager.go:222] Version: {KernelVersion:4.2.3-300.fc23.x86_64 ContainerOsVersion:Fedora 23 (Twenty Three) DockerVersion:1.10.3 DockerAPIVersion:1.22 CadvisorVersion: CadvisorRevision:}
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.380104   23369 server.go:422] --cgroups-per-qos enabled, but --cgroup-root was not specified.  defaulting to /
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.381344   23369 container_manager_linux.go:252] container manager verified user specified cgroup-root exists: /
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.381506   23369 container_manager_linux.go:257] Creating Container Manager object based on Node Config: {RuntimeCgroupsName: SystemCgroupsName: KubeletCgroupsName: ContainerRuntime:docker CgroupsPerQOS:true CgroupRoot:/ CgroupDriver:systemd ProtectKernelDefaults:false NodeAllocatableConfig:{KubeReservedCgroupName: SystemReservedCgroupName: EnforceNodeAllocatable:map[pods:{}] KubeReserved:map[] SystemReserved:map[] HardEvictionThresholds:[{Signal:memory.available Operator:LessThan Value:{Quantity:100Mi Percentage:0} GracePeriod:0s MinReclaim:<nil>} {Signal:nodefs.available Operator:LessThan Value:{Quantity:<nil> Percentage:0.1} GracePeriod:0s MinReclaim:<nil>} {Signal:nodefs.inodesFree Operator:LessThan Value:{Quantity:<nil> Percentage:0.05} GracePeriod:0s MinReclaim:<nil>}]} ExperimentalQOSReserved:map[] ExperimentalCPUManagerPolicy:none ExperimentalCPUManagerReconcilePeriod:10s}
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.382131   23369 container_manager_linux.go:288] Creating device plugin handler: false
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.382329   23369 kubelet.go:273] Adding manifest file: /etc/kubernetes/manifests
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.382476   23369 kubelet.go:283] Watching apiserver
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: E1124 05:33:02.411258   23369 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.64.33.199:443/api/v1/nodes?fieldSelector=metadata.name%3Dkubedev-10-64-33-199&resourceVersion=0: dial tcp 10.64.33.199:443: getsockopt: connection refused
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: E1124 05:33:02.411556   23369 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.64.33.199:443/api/v1/services?resourceVersion=0: dial tcp 10.64.33.199:443: getsockopt: connection refused
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: E1124 05:33:02.411732   23369 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.64.33.199:443/api/v1/pods?fieldSelector=spec.nodeName%3Dkubedev-10-64-33-199&resourceVersion=0: dial tcp 10.64.33.199:443: getsockopt: connection refused
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.426019   23369 kubelet_network.go:69] Hairpin mode set to "promiscuous-bridge" but kubenet is not enabled, falling back to "hairpin-veth"
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.426270   23369 kubelet.go:517] Hairpin mode set to "hairpin-veth"
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.427104   23369 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.434669   23369 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.437124   23369 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.437304   23369 docker_service.go:207] Docker cri networking managed by cni
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: I1124 05:33:02.531892   23369 docker_service.go:212] Docker Info: &{ID:YHFD:22Z2:JM6H:AUT2:XA72:6POC:DFYQ:JTFZ:KE3Y:QPUM:ZY33:DOVO Containers:3 ContainersRunning:3 ContainersPaused:0 ContainersStopped:0 Images:19 Driver:devicemapper DriverStatus:[[Pool Name docker-253:0-1311012-pool] [Pool Blocksize 65.54 kB] [Base Device Size 10.74 GB] [Backing Filesystem xfs] [Data file /dev/vg_vagrant/docker-data] [Metadata file /dev/vg_vagrant/docker-metadata] [Data Space Used 3.16 GB] [Data Space Total 30.06 GB] [Data Space Available 26.9 GB] [Metadata Space Used 3.535 MB] [Metadata Space Total 2.181 GB] [Metadata Space Available 2.178 GB] [Udev Sync Supported true] [Deferred Removal Enabled true] [Deferred Deletion Enabled true] [Deferred Deleted Device Count 0] [Library Version 1.02.109 (2015-09-22)]] SystemStatus:[] Plugins:{Volume:[local] Network:[host bridge null] Authorization:[] Log:[]} MemoryLimit:true SwapLimit:true KernelMemory:false CPUCfsPeriod:true CPUCfsQuota:true CPUShares:true CPUSet:true IPv4Forwarding:true BridgeNfIptables:true BridgeNfIP6tables:true Debug:false NFd:39 OomKillDisable:true NGoroutines:62 SystemTime:2017-11-24T05:33:02.53142935Z LoggingDriver:journald CgroupDriver: NEventsListener:0 KernelVersion:4.2.3-300.fc23.x86_64 OperatingSystem:Fedora 23 (Twenty Three) OSType:linux Architecture:x86_64 IndexServerAddress:https://index.docker.io/v1/ RegistryConfig:0xc4202c5340 NCPU:1 MemTotal:3155668992 GenericResources:[] DockerRootDir:/var/lib/docker HTTPProxy: HTTPSProxy: NoProxy: Name:kubedev-10-64-33-199 Labels:[] ExperimentalBuild:false ServerVersion:1.10.3 ClusterStore: ClusterAdvertise: Runtimes:map[] DefaultRuntime: Swarm:{NodeID: NodeAddr: LocalNodeState: ControlAvailable:false Error: RemoteManagers:[] Nodes:0 Managers:0 Cluster:<nil>} LiveRestoreEnabled:false Isolation: InitBinary: ContainerdCommit:{ID: Expected:} RuncCommit:{ID: Expected:} InitCommit:{ID: Expected:} SecurityOptions:[]}
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.531947   23369 docker_service.go:217] No cgroup driver is set in Docker
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: W1124 05:33:02.531952   23369 docker_service.go:218] Falling back to use the default driver: "cgroupfs"
Nov 24 05:33:02 kubedev-10-64-33-199 kubelet[23369]: error: failed to run Kubelet: failed to create kubelet: misconfiguration: kubelet cgroup driver: "systemd" is different from docker cgroup driver: "cgroupfs"
Nov 24 05:33:02 kubedev-10-64-33-199 systemd[1]: kubelet.service: Main process exited, code=exited, status=1/FAILURE
Nov 24 05:33:02 kubedev-10-64-33-199 systemd[1]: kubelet.service: Unit entered failed state.
Nov 24 05:33:02 kubedev-10-64-33-199 systemd[1]: kubelet.service: Failed with result 'exit-code'.
```

```
[vagrant@localhost ~]$ sudo sed -i 's/\(Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"\)/#\1\nEnvironment="KUBELET_CGROUP_ARGS=--cgroup-driver=cgroupfs"/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

```
[vagrant@localhost ~]$ sudo systemctl daemon-reload 
```

```
[vagrant@localhost ~]$ sudo systemctl restart kubelet.service
```

Actually kubelet v1.8 can not worked with docker version below 1.11
```
[vagrant@localhost ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Fri 2017-11-24 05:34:48 UTC; 1min 22s ago
     Docs: http://kubernetes.io/docs/
 Main PID: 23709 (kubelet)
   Memory: 25.6M
      CPU: 2.114s
   CGroup: /system.slice/kubelet.service
           └─23709 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin --cluster-dns=10.96.0.10 --cluster-domain=cluster.local --authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt --cadvisor-port=0 --cgroup-driver=cgroupfs --rotate-certificates=true --cert-dir=/var/lib/kubelet/pki --fail-swap-on=false

Nov 24 05:36:09 kubedev-10-64-33-199 kubelet[23709]: E1124 05:36:09.601832   23709 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Nov 24 05:36:09 kubedev-10-64-33-199 kubelet[23709]: E1124 05:36:09.764557   23709 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.64.33.199:443/api/v1/nodes?fieldSelector=metadata.name%3Dkubedev-10-64-33-199&resourceVersion=0: dial tcp 10.64.33.199:443: getsockopt: connection refused
Nov 24 05:36:09 kubedev-10-64-33-199 kubelet[23709]: E1124 05:36:09.783178   23709 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.64.33.199:443/api/v1/pods?fieldSelector=spec.nodeName%3Dkubedev-10-64-33-199&resourceVersion=0: dial tcp 10.64.33.199:443: getsockopt: connection refused
Nov 24 05:36:09 kubedev-10-64-33-199 kubelet[23709]: E1124 05:36:09.786683   23709 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.64.33.199:443/api/v1/services?resourceVersion=0: dial tcp 10.64.33.199:443: getsockopt: connection refused
Nov 24 05:36:09 kubedev-10-64-33-199 kubelet[23709]: I1124 05:36:09.807996   23709 kubelet_node_status.go:280] Setting node annotation to enable volume controller attach/detach
Nov 24 05:36:09 kubedev-10-64-33-199 kubelet[23709]: E1124 05:36:09.870935   23709 kubelet.go:1612] Failed creating a mirror pod for "etcd-kubedev-10-64-33-199_kube-system(d76e26fba3bf2bfd215eb29011d55250)": Post https://10.64.33.199:443/api/v1/namespaces/kube-system/pods: dial tcp 10.64.33.199:443: getsockopt: connection refused
Nov 24 05:36:10 kubedev-10-64-33-199 kubelet[23709]: E1124 05:36:10.188040   23709 remote_runtime.go:92] RunPodSandbox from runtime service failed: rpc error: code = Unknown desc = failed to create a sandbox for pod "etcd-kubedev-10-64-33-199": Error response from daemon: cgroup-parent for systemd cgroup should be a valid slice named as "xxx.slice"
Nov 24 05:36:10 kubedev-10-64-33-199 kubelet[23709]: E1124 05:36:10.188129   23709 kuberuntime_sandbox.go:54] CreatePodSandbox for pod "etcd-kubedev-10-64-33-199_kube-system(d76e26fba3bf2bfd215eb29011d55250)" failed: rpc error: code = Unknown desc = failed to create a sandbox for pod "etcd-kubedev-10-64-33-199": Error response from daemon: cgroup-parent for systemd cgroup should be a valid slice named as "xxx.slice"
Nov 24 05:36:10 kubedev-10-64-33-199 kubelet[23709]: E1124 05:36:10.188166   23709 kuberuntime_manager.go:633] createPodSandbox for pod "etcd-kubedev-10-64-33-199_kube-system(d76e26fba3bf2bfd215eb29011d55250)" failed: rpc error: code = Unknown desc = failed to create a sandbox for pod "etcd-kubedev-10-64-33-199": Error response from daemon: cgroup-parent for systemd cgroup should be a valid slice named as "xxx.slice"
Nov 24 05:36:10 kubedev-10-64-33-199 kubelet[23709]: E1124 05:36:10.188275   23709 pod_workers.go:182] Error syncing pod d76e26fba3bf2bfd215eb29011d55250 ("etcd-kubedev-10-64-33-199_kube-system(d76e26fba3bf2bfd215eb29011d55250)"), skipping: failed to "CreatePodSandbox" for "etcd-kubedev-10-64-33-199_kube-system(d76e26fba3bf2bfd215eb29011d55250)" with CreatePodSandboxError: "CreatePodSandbox for pod \"etcd-kubedev-10-64-33-199_kube-system(d76e26fba3bf2bfd215eb29011d55250)\" failed: rpc error: code = Unknown desc = failed to create a sandbox for pod \"etcd-kubedev-10-64-33-199\": Error response from daemon: cgroup-parent for systemd cgroup should be a valid slice named as \"xxx.slice\""
```

upgrade Kernel
```
[vagrant@localhost ~]$ sudo dnf upgrade
Failed to synchronize cache for repo 'kubernetes', disabling.
Last metadata expiration check: 2:51:27 ago on Fri Nov 24 02:55:53 2017.
Dependencies resolved.
====================================================================================================================================
 Package                           Arch         Version                    Repository                                          Size
====================================================================================================================================
Installing:
 kernel                            x86_64       4.8.13-100.fc23            updates                                             66 k
 kernel-core                       x86_64       4.8.13-100.fc23            updates                                             20 M
 kernel-devel                      x86_64       4.8.13-100.fc23            updates                                             11 M
 kernel-modules                    x86_64       4.8.13-100.fc23            updates                                             22 M
Upgrading:
 bind                              x86_64       32:9.10.4-2.P4.fc23        updates                                            1.9 M
 bind-libs                         x86_64       32:9.10.4-2.P4.fc23        updates                                            152 k
 bind-libs-lite                    x86_64       32:9.10.4-2.P4.fc23        updates                                            1.0 M
 bind-license                      noarch       32:9.10.4-2.P4.fc23        updates                                             89 k
 bind-utils                        x86_64       32:9.10.4-2.P4.fc23        updates                                            413 k
 bind99-libs                       x86_64       9.9.9-4.P4.fc23            updates                                            664 k
 bind99-license                    noarch       9.9.9-4.P4.fc23            updates                                             12 k
 dbus                              x86_64       1:1.10.14-1.fc23           updates                                            245 k
 dbus-libs                         x86_64       1:1.10.14-1.fc23           updates                                            173 k
 dmidecode                         x86_64       1:3.0-6.fc23               updates                                             88 k
 dnsmasq                           x86_64       2.76-2.fc23                updates                                            297 k
 firefox                           x86_64       50.1.0-1.fc23              updates                                             82 M
 ghostscript                       x86_64       9.20-5.fc23                updates                                             45 k
 ghostscript-core                  x86_64       9.20-5.fc23                updates                                            4.5 M
 ghostscript-x11                   x86_64       9.20-5.fc23                updates                                             73 k
 hwdata                            noarch       0.294-1.fc23               updates                                            1.4 M
 kernel-headers                    x86_64       4.8.13-100.fc23            updates                                            1.0 M
 libpng                            x86_64       2:1.6.26-1.fc23            updates                                            119 k
 ntp                               x86_64       4.2.6p5-43.fc23            updates                                            554 k
 ntpdate                           x86_64       4.2.6p5-43.fc23            updates                                             90 k
 pciutils                          x86_64       3.5.2-1.fc23               updates                                            100 k
 pciutils-libs                     x86_64       3.5.2-1.fc23               updates                                             50 k
 perl                              x86_64       4:5.22.2-355.fc23          updates                                            6.9 M
 perl-Compress-Bzip2               x86_64       2.25-1.fc23                updates                                             69 k
 perl-IO-Zlib                      noarch       1:1.10-355.fc23            updates                                             65 k
 perl-Locale-Maketext-Simple       noarch       1:0.21-355.fc23            updates                                             64 k
 perl-Math-BigInt                  noarch       1.9997-355.fc23            updates                                            188 k
 perl-Module-CoreList              noarch       1:5.20161120-1.fc23        updates                                             76 k
 perl-Time-Local                   noarch       1:1.250-1.fc23             updates                                             30 k
 perl-devel                        x86_64       4:5.22.2-355.fc23          updates                                            551 k
 perl-libs                         x86_64       4:5.22.2-355.fc23          updates                                            823 k
 perl-macros                       x86_64       4:5.22.2-355.fc23          updates                                             57 k
 python3-sssdconfig                noarch       1.14.2-2.fc23              updates                                            102 k
 quota                             x86_64       1:4.02-6.fc23              updates                                            169 k
 quota-nls                         noarch       1:4.02-6.fc23              updates                                             94 k
 sudo                              x86_64       1.8.18p1-1.fc23            updates                                            725 k
Skipping packages with broken dependencies:
 google-chrome-stable              x86_64       62.0.3202.94-1             dl.google.com_linux_chrome_rpm_stable_x86_64        47 M

Transaction Summary
====================================================================================================================================
Install   4 Packages
Upgrade  36 Packages
Skip      1 Package

Total download size: 157 M
Is this ok [y/N]: y
Downloading Packages:
(1/40): bind99-libs-9.9.9-2.P3.fc23_9.9.9-4.P4.fc23.x86_64.drpm                                      44 kB/s | 106 kB     00:02    
(2/40): bind99-license-9.9.9-2.P3.fc23_9.9.9-4.P4.fc23.noarch.drpm                                   10 kB/s | 5.2 kB     00:00    
[DRPM] bind99-libs-9.9.9-2.P3.fc23_9.9.9-4.P4.fc23.x86_64.drpm: done                                                               
(3/40): dbus-libs-1.10.12-1.fc23_1.10.14-1.fc23.x86_64.drpm                                          54 kB/s |  41 kB     00:00    
[DRPM] bind99-license-9.9.9-2.P3.fc23_9.9.9-4.P4.fc23.noarch.drpm: done                                                            
(4/40): ghostscript-core-9.20-2.fc23_9.20-5.fc23.x86_64.drpm                                         67 kB/s | 229 kB     00:03    
[DRPM] dbus-libs-1.10.12-1.fc23_1.10.14-1.fc23.x86_64.drpm: done                                                                   
[MIRROR] bind-libs-lite-9.10.4-2.P3.fc23_9.10.4-2.P4.fc23.x86_64.drpm: Curl error (56): Failure when receiving data from the peer for http://mirror.math.princeton.edu/pub/fedora-archive/fedora/linux/updates/23/x86_64/drpms/bind-libs-lite-9.10.4-2.P3.fc23_9.10.4-2.P4.fc23.x86_64.drpm [Recv failure: Connection reset by peer]
(5/40): hwdata-0.293-1.fc23_0.294-1.fc23.noarch.drpm                                                 55 kB/s |  43 kB     00:00    
(6/40): libpng-1.6.23-1.fc23_1.6.26-1.fc23.x86_64.drpm                                               44 kB/s |  35 kB     00:00    
(7/40): pciutils-libs-3.5.1-1.fc23_3.5.2-1.fc23.x86_64.drpm                                          37 kB/s |  20 kB     00:00    
(8/40): bind-9.10.4-2.P3.fc23_9.10.4-2.P4.fc23.x86_64.drpm                                           45 kB/s | 474 kB     00:10    
[DRPM] ghostscript-core-9.20-2.fc23_9.20-5.fc23.x86_64.drpm: done                                                                  
(9/40): perl-libs-5.22.2-354.fc23_5.22.2-355.fc23.x86_64.drpm                                        25 kB/s |  98 kB     00:03    
[DRPM] hwdata-0.293-1.fc23_0.294-1.fc23.noarch.drpm: done                                                                          
(10/40): perl-Compress-Bzip2-2.24-1.fc23_2.25-1.fc23.x86_64.drpm                                     32 kB/s |  22 kB     00:00    
[DRPM] libpng-1.6.23-1.fc23_1.6.26-1.fc23.x86_64.drpm: done                                                                        
(11/40): bind-libs-lite-9.10.4-2.P3.fc23_9.10.4-2.P4.fc23.x86_64.drpm                               9.8 kB/s | 158 kB     00:16    
[DRPM] pciutils-libs-3.5.1-1.fc23_3.5.2-1.fc23.x86_64.drpm: done                                                                   
(12/40): perl-Math-BigInt-1.9997-354.fc23_1.9997-355.fc23.noarch.drpm                                38 kB/s |  54 kB     00:01    
(13/40): perl-Time-Local-1.240-1.fc23_1.250-1.fc23.noarch.drpm                                       41 kB/s |  12 kB     00:00    
(14/40): perl-Module-CoreList-5.20161020-1.fc23_5.20161120-1.fc23.noarch.drpm                        13 kB/s |  14 kB     00:01    
[DRPM] bind-9.10.4-2.P3.fc23_9.10.4-2.P4.fc23.x86_64.drpm: done                                                                    
(15/40): quota-nls-4.02-5.fc23_4.02-6.fc23.noarch.drpm                                               26 kB/s |  21 kB     00:00    
[DRPM] perl-libs-5.22.2-354.fc23_5.22.2-355.fc23.x86_64.drpm: done                                                                 
(16/40): perl-devel-5.22.2-354.fc23_5.22.2-355.fc23.x86_64.drpm                                      49 kB/s |  72 kB     00:01    
[DRPM] perl-Compress-Bzip2-2.24-1.fc23_2.25-1.fc23.x86_64.drpm: done                                                               
(17/40): kernel-4.8.13-100.fc23.x86_64.rpm                                                           44 kB/s |  66 kB     00:01    
(18/40): perl-5.22.2-354.fc23_5.22.2-355.fc23.x86_64.drpm                                            35 kB/s | 739 kB     00:20    
[DRPM] bind-libs-lite-9.10.4-2.P3.fc23_9.10.4-2.P4.fc23.x86_64.drpm: done                                                          
(19/40): kernel-devel-4.8.13-100.fc23.x86_64.rpm                                                     25 kB/s |  11 MB     07:14    
(20/40): bind-libs-9.10.4-2.P4.fc23.x86_64.rpm                                                       12 kB/s | 152 kB     00:13    
(21/40): bind-license-9.10.4-2.P4.fc23.noarch.rpm                                                    32 kB/s |  89 kB     00:02    
(22/40): bind-utils-9.10.4-2.P4.fc23.x86_64.rpm                                                      35 kB/s | 413 kB     00:11    
(23/40): dbus-1.10.14-1.fc23.x86_64.rpm                                                              23 kB/s | 245 kB     00:10    
(24/40): dmidecode-3.0-6.fc23.x86_64.rpm                                                             21 kB/s |  88 kB     00:04    
(25/40): dnsmasq-2.76-2.fc23.x86_64.rpm                                                              33 kB/s | 297 kB     00:09    
(26/40): kernel-core-4.8.13-100.fc23.x86_64.rpm                                                      26 kB/s |  20 MB     13:04    
(27/40): ghostscript-9.20-5.fc23.x86_64.rpm                                                          22 kB/s |  45 kB     00:02    
(28/40): ghostscript-x11-9.20-5.fc23.x86_64.rpm                                                      36 kB/s |  73 kB     00:02    
(29/40): kernel-headers-4.8.13-100.fc23.x86_64.rpm                                                   34 kB/s | 1.0 MB     00:31    
(30/40): ntp-4.2.6p5-43.fc23.x86_64.rpm                                                              36 kB/s | 554 kB     00:15    
(31/40): ntpdate-4.2.6p5-43.fc23.x86_64.rpm                                                          24 kB/s |  90 kB     00:03    
(32/40): pciutils-3.5.2-1.fc23.x86_64.rpm                                                            18 kB/s | 100 kB     00:05    
(33/40): perl-IO-Zlib-1.10-355.fc23.noarch.rpm                                                       47 kB/s |  65 kB     00:01    
(34/40): perl-Locale-Maketext-Simple-0.21-355.fc23.noarch.rpm                                        46 kB/s |  64 kB     00:01    
(35/40): perl-macros-5.22.2-355.fc23.x86_64.rpm                                                      50 kB/s |  57 kB     00:01    
(36/40): python3-sssdconfig-1.14.2-2.fc23.noarch.rpm                                                 35 kB/s | 102 kB     00:02    
(37/40): quota-4.02-6.fc23.x86_64.rpm                                                                45 kB/s | 169 kB     00:03    
(38/40): kernel-modules-4.8.13-100.fc23.x86_64.rpm                                                   26 kB/s |  22 MB     14:29    
(39/40): sudo-1.8.18p1-1.fc23.x86_64.rpm                                                             43 kB/s | 725 kB     00:16    
(40/40): firefox-50.1.0-1.fc23.x86_64.rpm                                                            26 kB/s |  82 MB     53:48    
[DRPM] perl-Math-BigInt-1.9997-354.fc23_1.9997-355.fc23.noarch.drpm: done                                                          
[DRPM] perl-Time-Local-1.240-1.fc23_1.250-1.fc23.noarch.drpm: done                                                                 
[DRPM] perl-Module-CoreList-5.20161020-1.fc23_5.20161120-1.fc23.noarch.drpm: done                                                  
[DRPM] quota-nls-4.02-5.fc23_4.02-6.fc23.noarch.drpm: done                                                                         
[DRPM] perl-devel-5.22.2-354.fc23_5.22.2-355.fc23.x86_64.drpm: done                                                                
[DRPM] perl-5.22.2-354.fc23_5.22.2-355.fc23.x86_64.drpm: done                                                                      
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                38 kB/s | 141 MB     62:36     
Delta RPMs reduced 157.3 MB of updates to 140.9 MB (10.1% saved)
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Upgrading   : perl-libs-4:5.22.2-355.fc23.x86_64                                                                             1/76 
  Upgrading   : dbus-libs-1:1.10.14-1.fc23.x86_64                                                                              2/76 
  Upgrading   : bind-license-32:9.10.4-2.P4.fc23.noarch                                                                        3/76 
  Upgrading   : bind-libs-lite-32:9.10.4-2.P4.fc23.x86_64                                                                      4/76 
  Upgrading   : bind-libs-32:9.10.4-2.P4.fc23.x86_64                                                                           5/76 
  Installing  : kernel-core-4.8.13-100.fc23.x86_64                                                                             6/76 
  Installing  : kernel-modules-4.8.13-100.fc23.x86_64                                                                          7/76 
  Upgrading   : perl-macros-4:5.22.2-355.fc23.x86_64                                                                           8/76 
  Upgrading   : perl-4:5.22.2-355.fc23.x86_64                                                                                  9/76 
  Upgrading   : quota-nls-1:4.02-6.fc23.noarch                                                                                10/76 
  Upgrading   : pciutils-libs-3.5.2-1.fc23.x86_64                                                                             11/76 
  Upgrading   : ntpdate-4.2.6p5-43.fc23.x86_64                                                                                12/76 
  Upgrading   : libpng-2:1.6.26-1.fc23.x86_64                                                                                 13/76 
  Upgrading   : ghostscript-core-9.20-5.fc23.x86_64                                                                           14/76 
  Upgrading   : ghostscript-x11-9.20-5.fc23.x86_64                                                                            15/76 
  Upgrading   : ghostscript-9.20-5.fc23.x86_64                                                                                16/76 
  Upgrading   : hwdata-0.294-1.fc23.noarch                                                                                    17/76 
  Upgrading   : bind99-license-9.9.9-4.P4.fc23.noarch                                                                         18/76 
  Upgrading   : bind99-libs-9.9.9-4.P4.fc23.x86_64                                                                            19/76 
  Upgrading   : pciutils-3.5.2-1.fc23.x86_64                                                                                  20/76 
  Upgrading   : ntp-4.2.6p5-43.fc23.x86_64                                                                                    21/76 
  Upgrading   : quota-1:4.02-6.fc23.x86_64                                                                                    22/76 
  Installing  : kernel-devel-4.8.13-100.fc23.x86_64                                                                           23/76 
  Upgrading   : perl-Compress-Bzip2-2.25-1.fc23.x86_64                                                                        24/76 
  Upgrading   : perl-IO-Zlib-1:1.10-355.fc23.noarch                                                                           25/76 
  Upgrading   : perl-Locale-Maketext-Simple-1:0.21-355.fc23.noarch                                                            26/76 
  Upgrading   : perl-Math-BigInt-1.9997-355.fc23.noarch                                                                       27/76 
  Upgrading   : perl-Module-CoreList-1:5.20161120-1.fc23.noarch                                                               28/76 
  Upgrading   : perl-Time-Local-1:1.250-1.fc23.noarch                                                                         29/76 
  Upgrading   : perl-devel-4:5.22.2-355.fc23.x86_64                                                                           30/76 
  Installing  : kernel-4.8.13-100.fc23.x86_64                                                                                 31/76 
  Upgrading   : bind-32:9.10.4-2.P4.fc23.x86_64                                                                               32/76 
  Upgrading   : bind-utils-32:9.10.4-2.P4.fc23.x86_64                                                                         33/76 
  Upgrading   : dbus-1:1.10.14-1.fc23.x86_64                                                                                  34/76 
  Upgrading   : dnsmasq-2.76-2.fc23.x86_64                                                                                    35/76 
  Upgrading   : firefox-50.1.0-1.fc23.x86_64                                                                                  36/76 
  Upgrading   : sudo-1.8.18p1-1.fc23.x86_64                                                                                   37/76 
  Upgrading   : python3-sssdconfig-1.14.2-2.fc23.noarch                                                                       38/76 
  Upgrading   : kernel-headers-4.8.13-100.fc23.x86_64                                                                         39/76 
  Upgrading   : dmidecode-1:3.0-6.fc23.x86_64                                                                                 40/76 
  Cleanup     : bind-utils-32:9.10.4-2.P3.fc23.x86_64                                                                         41/76 
  Cleanup     : bind-32:9.10.4-2.P3.fc23.x86_64                                                                               42/76 
  Cleanup     : perl-Compress-Bzip2-2.24-1.fc23.x86_64                                                                        43/76 
  Cleanup     : perl-devel-4:5.22.2-354.fc23.x86_64                                                                           44/76 
  Cleanup     : perl-Math-BigInt-1.9997-354.fc23.noarch                                                                       45/76 
  Cleanup     : perl-Locale-Maketext-Simple-1:0.21-354.fc23.noarch                                                            46/76 
  Cleanup     : perl-IO-Zlib-1:1.10-354.fc23.noarch                                                                           47/76 
  Cleanup     : ghostscript-x11-9.20-2.fc23.x86_64                                                                            48/76 
  Cleanup     : ghostscript-9.20-2.fc23.x86_64                                                                                49/76 
  Cleanup     : perl-Time-Local-1:1.240-1.fc23.noarch                                                                         50/76 
  Cleanup     : perl-Module-CoreList-1:5.20161020-1.fc23.noarch                                                               51/76 
  Cleanup     : bind-libs-lite-32:9.10.4-2.P3.fc23.x86_64                                                                     52/76 
  Cleanup     : bind-libs-32:9.10.4-2.P3.fc23.x86_64                                                                          53/76 
  Cleanup     : perl-4:5.22.2-354.fc23.x86_64                                                                                 54/76 
  Cleanup     : pciutils-3.5.1-1.fc23.x86_64                                                                                  55/76 
  Cleanup     : perl-macros-4:5.22.2-354.fc23.x86_64                                                                          56/76 
  Cleanup     : ghostscript-core-9.20-2.fc23.x86_64                                                                           57/76 
  Cleanup     : quota-1:4.02-5.fc23.x86_64                                                                                    58/76 
  Cleanup     : ntp-4.2.6p5-41.fc23.x86_64                                                                                    59/76 
  Cleanup     : firefox-49.0.2-1.fc23.x86_64                                                                                  60/76 
  Cleanup     : dnsmasq-2.76-1.fc23.x86_64                                                                                    61/76 
  Cleanup     : dbus-1:1.10.12-1.fc23.x86_64                                                                                  62/76 
  Cleanup     : bind99-libs-9.9.9-2.P3.fc23.x86_64                                                                            63/76 
  Cleanup     : bind99-license-9.9.9-2.P3.fc23.noarch                                                                         64/76 
  Cleanup     : quota-nls-1:4.02-5.fc23.noarch                                                                                65/76 
  Cleanup     : hwdata-0.293-1.fc23.noarch                                                                                    66/76 
  Cleanup     : bind-license-32:9.10.4-2.P3.fc23.noarch                                                                       67/76 
  Cleanup     : python3-sssdconfig-1.14.2-1.fc23.noarch                                                                       68/76 
  Cleanup     : kernel-headers-4.7.9-100.fc23.x86_64                                                                          69/76 
  Cleanup     : dbus-libs-1:1.10.12-1.fc23.x86_64                                                                             70/76 
  Cleanup     : ntpdate-4.2.6p5-41.fc23.x86_64                                                                                71/76 
  Cleanup     : libpng-2:1.6.23-1.fc23.x86_64                                                                                 72/76 
  Cleanup     : perl-libs-4:5.22.2-354.fc23.x86_64                                                                            73/76 
  Cleanup     : pciutils-libs-3.5.1-1.fc23.x86_64                                                                             74/76 
  Cleanup     : sudo-1.8.18-1.fc23.x86_64                                                                                     75/76 
  Cleanup     : dmidecode-1:3.0-4.fc23.x86_64                                                                                 76/76 
  Verifying   : kernel-4.8.13-100.fc23.x86_64                                                                                  1/76 
  Verifying   : kernel-core-4.8.13-100.fc23.x86_64                                                                             2/76 
  Verifying   : kernel-modules-4.8.13-100.fc23.x86_64                                                                          3/76 
  Verifying   : kernel-devel-4.8.13-100.fc23.x86_64                                                                            4/76 
  Verifying   : bind-32:9.10.4-2.P4.fc23.x86_64                                                                                5/76 
  Verifying   : bind-libs-32:9.10.4-2.P4.fc23.x86_64                                                                           6/76 
  Verifying   : bind-libs-lite-32:9.10.4-2.P4.fc23.x86_64                                                                      7/76 
  Verifying   : bind-license-32:9.10.4-2.P4.fc23.noarch                                                                        8/76 
  Verifying   : bind-utils-32:9.10.4-2.P4.fc23.x86_64                                                                          9/76 
  Verifying   : bind99-libs-9.9.9-4.P4.fc23.x86_64                                                                            10/76 
  Verifying   : bind99-license-9.9.9-4.P4.fc23.noarch                                                                         11/76 
  Verifying   : dbus-1:1.10.14-1.fc23.x86_64                                                                                  12/76 
  Verifying   : dbus-libs-1:1.10.14-1.fc23.x86_64                                                                             13/76 
  Verifying   : dmidecode-1:3.0-6.fc23.x86_64                                                                                 14/76 
  Verifying   : dnsmasq-2.76-2.fc23.x86_64                                                                                    15/76 
  Verifying   : firefox-50.1.0-1.fc23.x86_64                                                                                  16/76 
  Verifying   : ghostscript-9.20-5.fc23.x86_64                                                                                17/76 
  Verifying   : ghostscript-core-9.20-5.fc23.x86_64                                                                           18/76 
  Verifying   : ghostscript-x11-9.20-5.fc23.x86_64                                                                            19/76 
  Verifying   : hwdata-0.294-1.fc23.noarch                                                                                    20/76 
  Verifying   : kernel-headers-4.8.13-100.fc23.x86_64                                                                         21/76 
  Verifying   : libpng-2:1.6.26-1.fc23.x86_64                                                                                 22/76 
  Verifying   : ntp-4.2.6p5-43.fc23.x86_64                                                                                    23/76 
  Verifying   : ntpdate-4.2.6p5-43.fc23.x86_64                                                                                24/76 
  Verifying   : pciutils-3.5.2-1.fc23.x86_64                                                                                  25/76 
  Verifying   : pciutils-libs-3.5.2-1.fc23.x86_64                                                                             26/76 
  Verifying   : perl-4:5.22.2-355.fc23.x86_64                                                                                 27/76 
  Verifying   : perl-libs-4:5.22.2-355.fc23.x86_64                                                                            28/76 
  Verifying   : perl-Compress-Bzip2-2.25-1.fc23.x86_64                                                                        29/76 
  Verifying   : perl-IO-Zlib-1:1.10-355.fc23.noarch                                                                           30/76 
  Verifying   : perl-Locale-Maketext-Simple-1:0.21-355.fc23.noarch                                                            31/76 
  Verifying   : perl-Math-BigInt-1.9997-355.fc23.noarch                                                                       32/76 
  Verifying   : perl-Module-CoreList-1:5.20161120-1.fc23.noarch                                                               33/76 
  Verifying   : perl-Time-Local-1:1.250-1.fc23.noarch                                                                         34/76 
  Verifying   : perl-devel-4:5.22.2-355.fc23.x86_64                                                                           35/76 
  Verifying   : perl-macros-4:5.22.2-355.fc23.x86_64                                                                          36/76 
  Verifying   : python3-sssdconfig-1.14.2-2.fc23.noarch                                                                       37/76 
  Verifying   : quota-1:4.02-6.fc23.x86_64                                                                                    38/76 
  Verifying   : quota-nls-1:4.02-6.fc23.noarch                                                                                39/76 
  Verifying   : sudo-1.8.18p1-1.fc23.x86_64                                                                                   40/76 
  Verifying   : ntp-4.2.6p5-41.fc23.x86_64                                                                                    41/76 
  Verifying   : ntpdate-4.2.6p5-41.fc23.x86_64                                                                                42/76 
  Verifying   : perl-Time-Local-1:1.240-1.fc23.noarch                                                                         43/76 
  Verifying   : perl-devel-4:5.22.2-354.fc23.x86_64                                                                           44/76 
  Verifying   : pciutils-3.5.1-1.fc23.x86_64                                                                                  45/76 
  Verifying   : pciutils-libs-3.5.1-1.fc23.x86_64                                                                             46/76 
  Verifying   : bind-32:9.10.4-2.P3.fc23.x86_64                                                                               47/76 
  Verifying   : bind-libs-32:9.10.4-2.P3.fc23.x86_64                                                                          48/76 
  Verifying   : bind-libs-lite-32:9.10.4-2.P3.fc23.x86_64                                                                     49/76 
  Verifying   : bind-license-32:9.10.4-2.P3.fc23.noarch                                                                       50/76 
  Verifying   : bind-utils-32:9.10.4-2.P3.fc23.x86_64                                                                         51/76 
  Verifying   : bind99-libs-9.9.9-2.P3.fc23.x86_64                                                                            52/76 
  Verifying   : bind99-license-9.9.9-2.P3.fc23.noarch                                                                         53/76 
  Verifying   : ghostscript-9.20-2.fc23.x86_64                                                                                54/76 
  Verifying   : ghostscript-core-9.20-2.fc23.x86_64                                                                           55/76 
  Verifying   : ghostscript-x11-9.20-2.fc23.x86_64                                                                            56/76 
  Verifying   : perl-4:5.22.2-354.fc23.x86_64                                                                                 57/76 
  Verifying   : perl-libs-4:5.22.2-354.fc23.x86_64                                                                            58/76 
  Verifying   : sudo-1.8.18-1.fc23.x86_64                                                                                     59/76 
  Verifying   : perl-Compress-Bzip2-2.24-1.fc23.x86_64                                                                        60/76 
  Verifying   : perl-macros-4:5.22.2-354.fc23.x86_64                                                                          61/76 
  Verifying   : dbus-1:1.10.12-1.fc23.x86_64                                                                                  62/76 
  Verifying   : dbus-libs-1:1.10.12-1.fc23.x86_64                                                                             63/76 
  Verifying   : hwdata-0.293-1.fc23.noarch                                                                                    64/76 
  Verifying   : libpng-2:1.6.23-1.fc23.x86_64                                                                                 65/76 
  Verifying   : perl-IO-Zlib-1:1.10-354.fc23.noarch                                                                           66/76 
  Verifying   : dmidecode-1:3.0-4.fc23.x86_64                                                                                 67/76 
  Verifying   : dnsmasq-2.76-1.fc23.x86_64                                                                                    68/76 
  Verifying   : perl-Locale-Maketext-Simple-1:0.21-354.fc23.noarch                                                            69/76 
  Verifying   : perl-Math-BigInt-1.9997-354.fc23.noarch                                                                       70/76 
  Verifying   : perl-Module-CoreList-1:5.20161020-1.fc23.noarch                                                               71/76 
  Verifying   : python3-sssdconfig-1.14.2-1.fc23.noarch                                                                       72/76 
  Verifying   : kernel-headers-4.7.9-100.fc23.x86_64                                                                          73/76 
  Verifying   : quota-1:4.02-5.fc23.x86_64                                                                                    74/76 
  Verifying   : quota-nls-1:4.02-5.fc23.noarch                                                                                75/76 
  Verifying   : firefox-49.0.2-1.fc23.x86_64                                                                                  76/76 

Installed:
  kernel.x86_64 4.8.13-100.fc23                kernel-core.x86_64 4.8.13-100.fc23        kernel-devel.x86_64 4.8.13-100.fc23       
  kernel-modules.x86_64 4.8.13-100.fc23       

Upgraded:
  bind.x86_64 32:9.10.4-2.P4.fc23                              bind-libs.x86_64 32:9.10.4-2.P4.fc23                                 
  bind-libs-lite.x86_64 32:9.10.4-2.P4.fc23                    bind-license.noarch 32:9.10.4-2.P4.fc23                              
  bind-utils.x86_64 32:9.10.4-2.P4.fc23                        bind99-libs.x86_64 9.9.9-4.P4.fc23                                   
  bind99-license.noarch 9.9.9-4.P4.fc23                        dbus.x86_64 1:1.10.14-1.fc23                                         
  dbus-libs.x86_64 1:1.10.14-1.fc23                            dmidecode.x86_64 1:3.0-6.fc23                                        
  dnsmasq.x86_64 2.76-2.fc23                                   firefox.x86_64 50.1.0-1.fc23                                         
  ghostscript.x86_64 9.20-5.fc23                               ghostscript-core.x86_64 9.20-5.fc23                                  
  ghostscript-x11.x86_64 9.20-5.fc23                           hwdata.noarch 0.294-1.fc23                                           
  kernel-headers.x86_64 4.8.13-100.fc23                        libpng.x86_64 2:1.6.26-1.fc23                                        
  ntp.x86_64 4.2.6p5-43.fc23                                   ntpdate.x86_64 4.2.6p5-43.fc23                                       
  pciutils.x86_64 3.5.2-1.fc23                                 pciutils-libs.x86_64 3.5.2-1.fc23                                    
  perl.x86_64 4:5.22.2-355.fc23                                perl-Compress-Bzip2.x86_64 2.25-1.fc23                               
  perl-IO-Zlib.noarch 1:1.10-355.fc23                          perl-Locale-Maketext-Simple.noarch 1:0.21-355.fc23                   
  perl-Math-BigInt.noarch 1.9997-355.fc23                      perl-Module-CoreList.noarch 1:5.20161120-1.fc23                      
  perl-Time-Local.noarch 1:1.250-1.fc23                        perl-devel.x86_64 4:5.22.2-355.fc23                                  
  perl-libs.x86_64 4:5.22.2-355.fc23                           perl-macros.x86_64 4:5.22.2-355.fc23                                 
  python3-sssdconfig.noarch 1.14.2-2.fc23                      quota.x86_64 1:4.02-6.fc23                                           
  quota-nls.noarch 1:4.02-6.fc23                               sudo.x86_64 1.8.18p1-1.fc23                                          

Complete!
```

local docker project repository
```
[vagrant@localhost ~]$ echo '
> [dockerrepo]
> name=Docker Repository
> baseurl=https://yum.dockerproject.org/repo/main/fedora/$releasever/
> enabled=1
> gpgcheck=1
> gpgkey=https://yum.dockerproject.org/gpg
> 
> [dockerrepo-local-files]
> name=Docker Repository local files repo
> baseurl=file:///vagrant_drive_f/16-mirror/https0x3A0x2F0x2Fdocs.docker.com0x2Fdocsarchive/https0x3A0x2F0x2Fyum.dockerproject.org0x2Frepo0x2Fmain/fedora0x2F23
> enabled=0
> gpgcheck=0
> 
> [dockerrepo-local-http]
> name=Docker Repository local files repo
> baseurl=http://10.64.33.82:48080/vagrant_f/16-mirror/https0x3A0x2F0x2Fdocs.docker.com0x2Fdocsarchive/https0x3A0x2F0x2Fyum.dockerproject.org0x2Frepo0x2Fmain/fedora0x2F23/
> enabled=0
> gpgcheck=0
> ' | sudo tee /etc/yum.repos.d/docker.repo

[dockerrepo]
name=Docker Repository
baseurl=https://yum.dockerproject.org/repo/main/fedora/$releasever/
enabled=1
gpgcheck=1
gpgkey=https://yum.dockerproject.org/gpg

[dockerrepo-local-files]
name=Docker Repository local files repo
baseurl=file:///vagrant_drive_f/16-mirror/https0x3A0x2F0x2Fdocs.docker.com0x2Fdocsarchive/https0x3A0x2F0x2Fyum.dockerproject.org0x2Frepo0x2Fmain/fedora0x2F23
enabled=0
gpgcheck=0

[dockerrepo-local-http]
name=Docker Repository local files repo
baseurl=http://10.64.33.82:48080/vagrant_f/16-mirror/https0x3A0x2F0x2Fdocs.docker.com0x2Fdocsarchive/https0x3A0x2F0x2Fapt.dockerproject.org0x2Frepo0x2Fmain/fedora0x2F23/
enabled=0
gpgcheck=0

```

### issue

upgrade docker engine
```
[vagrant@localhost ~]$ sudo dnf --disablerepo=dockerrepo --enablerepo=dockerrepo-* list | egrep "^docker-engine"
Failed to synchronize cache for repo 'kubernetes', disabling.
docker-engine.src                        1.12.6-1.fc23                  dockerrepo-local-files
docker-engine.x86_64                     1.12.6-1.fc23                  dockerrepo-local-files
docker-engine-debuginfo.x86_64           1.12.6-1.fc23                  dockerrepo-local-files
docker-engine-selinux.noarch             1.12.6-1.fc23                  dockerrepo-local-files
docker-engine-selinux.src                1.12.6-1.fc23                  dockerrepo-local-files
```

```
[vagrant@localhost ~]$ sudo dnf --disablerepo=kubernetes --disablerepo=dockerrepo --enablerepo=dockerrepo-* --allowerasing install docker-engine
Last metadata expiration check: 0:11:02 ago on Fri Nov 24 08:45:00 2017.
Dependencies resolved.
====================================================================================================================================
 Package                           Arch              Version                                Repository                         Size
====================================================================================================================================
Installing:
 audit-libs-python                 x86_64            2.6.7-1.fc23                           updates                            96 k
 docker-engine                     x86_64            1.12.6-1.fc23                          dockerrepo-local-files             19 M
 docker-engine-selinux             noarch            1.12.6-1.fc23                          dockerrepo-local-files             32 k
 libselinux-python                 x86_64            2.4-4.fc23                             fedora                            302 k
 libsemanage-python                x86_64            2.4-4.fc23                             fedora                            109 k
 libtool-ltdl                      x86_64            2.4.6-8.fc23                           updates                            54 k
 policycoreutils-python            x86_64            2.4-21.fc23                            updates                           400 k
 python-IPy                        noarch            0.81-13.fc23                           fedora                             42 k
Removing:
 docker                            x86_64            2:1.10.3-45.gite03ddb8.fc23            @updates                           43 M
 docker-selinux                    x86_64            2:1.10.3-45.gite03ddb8.fc23            @updates                           27 k

Transaction Summary
====================================================================================================================================
Install  8 Packages
Remove   2 Packages

Total size: 20 M
Total download size: 1.0 M
Is this ok [y/N]: y
Downloading Packages:
(1/6): audit-libs-python-2.6.7-1.fc23.x86_64.rpm                                                     11 kB/s |  96 kB     00:08    
(2/6): libtool-ltdl-2.4.6-8.fc23.x86_64.rpm                                                         6.0 kB/s |  54 kB     00:09    
(3/6): policycoreutils-python-2.4-21.fc23.x86_64.rpm                                                 43 kB/s | 400 kB     00:09    
(4/6): python-IPy-0.81-13.fc23.noarch.rpm                                                            30 kB/s |  42 kB     00:01    
(5/6): libsemanage-python-2.4-4.fc23.x86_64.rpm                                                      37 kB/s | 109 kB     00:02    
(6/6): libselinux-python-2.4-4.fc23.x86_64.rpm                                                       83 kB/s | 302 kB     00:03    
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                52 kB/s | 1.0 MB     00:19     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Installing  : python-IPy-0.81-13.fc23.noarch                                                                                 1/10 
  Installing  : libsemanage-python-2.4-4.fc23.x86_64                                                                           2/10 
  Installing  : libselinux-python-2.4-4.fc23.x86_64                                                                            3/10 
  Installing  : audit-libs-python-2.6.7-1.fc23.x86_64                                                                          4/10 
  Installing  : policycoreutils-python-2.4-21.fc23.x86_64                                                                      5/10 
  Installing  : docker-engine-selinux-1.12.6-1.fc23.noarch                                                                     6/10 
  Installing  : libtool-ltdl-2.4.6-8.fc23.x86_64                                                                               7/10 
  Installing  : docker-engine-1.12.6-1.fc23.x86_64                                                                             8/10 
  Erasing     : docker-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                      9/10 
warning: /etc/sysconfig/docker-storage saved as /etc/sysconfig/docker-storage.rpmsave
warning: /etc/sysconfig/docker saved as /etc/sysconfig/docker.rpmsave
  Erasing     : docker-selinux-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                             10/10 
  Verifying   : docker-engine-1.12.6-1.fc23.x86_64                                                                             1/10 
  Verifying   : libtool-ltdl-2.4.6-8.fc23.x86_64                                                                               2/10 
  Verifying   : docker-engine-selinux-1.12.6-1.fc23.noarch                                                                     3/10 
  Verifying   : policycoreutils-python-2.4-21.fc23.x86_64                                                                      4/10 
  Verifying   : audit-libs-python-2.6.7-1.fc23.x86_64                                                                          5/10 
  Verifying   : libselinux-python-2.4-4.fc23.x86_64                                                                            6/10 
  Verifying   : libsemanage-python-2.4-4.fc23.x86_64                                                                           7/10 
  Verifying   : python-IPy-0.81-13.fc23.noarch                                                                                 8/10 
  Verifying   : docker-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                      9/10 
  Verifying   : docker-selinux-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                             10/10 

Removed:
  docker.x86_64 2:1.10.3-45.gite03ddb8.fc23                    docker-selinux.x86_64 2:1.10.3-45.gite03ddb8.fc23                   

Installed:
  audit-libs-python.x86_64 2.6.7-1.fc23        docker-engine.x86_64 1.12.6-1.fc23      docker-engine-selinux.noarch 1.12.6-1.fc23   
  libselinux-python.x86_64 2.4-4.fc23          libsemanage-python.x86_64 2.4-4.fc23    libtool-ltdl.x86_64 2.4.6-8.fc23             
  policycoreutils-python.x86_64 2.4-21.fc23    python-IPy.noarch 0.81-13.fc23         

Complete!
```

```
[vagrant@localhost ~]$ docker version
Client:
 Version:      1.12.6
 API version:  1.24
 Go version:   go1.6.4
 Git commit:   78d1802
 Built:        Tue Jan 10 20:21:11 2017
 OS/Arch:      linux/amd64
Error response from daemon: client is newer than server (client API version: 1.24, server API version: 1.22)
```

```
[vagrant@localhost ~]$ sudo systemctl -l status docker.service
● docker.service - Docker Application Container Engine
   Loaded: loaded (/usr/lib/systemd/system/docker.service; disabled; vendor preset: disabled)
   Active: inactive (dead) since Fri 2017-11-24 08:56:54 UTC; 1min 16s ago
     Docs: https://docs.docker.com
 Main PID: 3900 (code=killed, signal=TERM)
   CGroup: /system.slice/docker.service
           ├─3901 /usr/bin/docker daemon --exec-opt native.cgroupdriver=systemd --insecure-registry=172.30.0.0/16 --insecure-registry=ci.dev.openshift.redhat.com:5000 --selinux-enabled --log-driver=journald -s devicemapper --storage-opt dm.datadev=/dev/vg_vagrant/docker-data --storage-opt dm.metadatadev=/dev/vg_vagrant/docker-metadata --storage-opt dm.use_deferred_removal=true --storage-opt dm.use_deferred_deletion=true
           ├─3902 /usr/bin/forward-journald -tag docker
           └─4142 docker-proxy -proto tcp -host-ip 0.0.0.0 -host-port 8082 -container-ip 172.18.0.4 -container-port 8082

Nov 24 08:57:51 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:57:51.226773784Z" level=info msg="{Action=create, LoginUID=4294967295, PID=23709}"
Nov 24 08:57:51 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:57:51.232878102Z" level=error msg="Handler for POST /v1.22/containers/create returned error: cgroup-parent for systemd cgroup should be a valid slice named as \"xxx.slice\""
Nov 24 08:58:00 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:58:00.206824740Z" level=info msg="{Action=create, LoginUID=4294967295, PID=23709}"
Nov 24 08:58:00 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:58:00.213153183Z" level=error msg="Handler for POST /v1.22/containers/create returned error: cgroup-parent for systemd cgroup should be a valid slice named as \"xxx.slice\""
Nov 24 08:58:00 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:58:00.219694776Z" level=info msg="{Action=create, LoginUID=4294967295, PID=23709}"
Nov 24 08:58:00 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:58:00.227978639Z" level=error msg="Handler for POST /v1.22/containers/create returned error: cgroup-parent for systemd cgroup should be a valid slice named as \"xxx.slice\""
Nov 24 08:58:05 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:58:05.195449947Z" level=info msg="{Action=create, LoginUID=4294967295, PID=23709}"
Nov 24 08:58:05 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:58:05.198641172Z" level=error msg="Handler for POST /v1.22/containers/create returned error: cgroup-parent for systemd cgroup should be a valid slice named as \"xxx.slice\""
Nov 24 08:58:05 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:58:05.200918177Z" level=info msg="{Action=create, LoginUID=4294967295, PID=23709}"
Nov 24 08:58:05 kubedev-10-64-33-199 forward-journal[3902]: time="2017-11-24T08:58:05.206610037Z" level=error msg="Handler for POST /v1.22/containers/create returned error: cgroup-parent for systemd cgroup should be a valid slice named as \"xxx.slice\""
[vagrant@localhost ~]$ sudo systemctl restart docker.service
Job for docker.service failed because the control process exited with error code. See "systemctl status docker.service" and "journalctl -xe" for details.
```

### downgrade docker engine

refer to https://docs.docker.com/v1.12/engine/reference/commandline/dockerd/#/options-for-the-runtime

To v1.11
```
[vagrant@localhost ~]$ sudo dnf --showduplicates --disablerepo=kubernetes --enablerepo=dockerrepo-\* list docker-engine
Last metadata expiration check: 0:24:51 ago on Fri Nov 24 09:48:17 2017.
Installed Packages
docker-engine.x86_64                                      1.12.6-1.fc23                                      @dockerrepo-local-files
Available Packages
docker-engine.src                                         1.9.1-1.fc23                                       dockerrepo             
docker-engine.src                                         1.9.1-1.fc23                                       dockerrepo-local-http  
docker-engine.src                                         1.9.1-1.fc23                                       dockerrepo-local-files 
docker-engine.x86_64                                      1.9.1-1.fc23                                       dockerrepo             
docker-engine.x86_64                                      1.9.1-1.fc23                                       dockerrepo-local-http  
docker-engine.x86_64                                      1.9.1-1.fc23                                       dockerrepo-local-files 
docker-engine.src                                         1.9.1-2.fc23                                       dockerrepo             
docker-engine.src                                         1.9.1-2.fc23                                       dockerrepo-local-http  
docker-engine.src                                         1.9.1-2.fc23                                       dockerrepo-local-files 
docker-engine.x86_64                                      1.9.1-2.fc23                                       dockerrepo             
docker-engine.x86_64                                      1.9.1-2.fc23                                       dockerrepo-local-http  
docker-engine.x86_64                                      1.9.1-2.fc23                                       dockerrepo-local-files 
docker-engine.src                                         1.10.0-1.fc23                                      dockerrepo             
docker-engine.src                                         1.10.0-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.10.0-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.10.0-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.10.0-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.10.0-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.10.1-1.fc23                                      dockerrepo             
docker-engine.src                                         1.10.1-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.10.1-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.10.1-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.10.1-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.10.1-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.10.2-1.fc23                                      dockerrepo             
docker-engine.src                                         1.10.2-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.10.2-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.10.2-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.10.2-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.10.2-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.10.3-1.fc23                                      dockerrepo             
docker-engine.src                                         1.10.3-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.10.3-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.10.3-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.10.3-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.10.3-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.11.0-1.fc23                                      dockerrepo             
docker-engine.src                                         1.11.0-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.11.0-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.11.0-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.11.0-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.11.0-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.11.1-1.fc23                                      dockerrepo             
docker-engine.src                                         1.11.1-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.11.1-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.11.1-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.11.1-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.11.1-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.11.2-1.fc23                                      dockerrepo             
docker-engine.src                                         1.11.2-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.11.2-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.11.2-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.11.2-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.11.2-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.12.0-1.fc23                                      dockerrepo             
docker-engine.src                                         1.12.0-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.12.0-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.12.0-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.12.0-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.12.0-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.12.1-1.fc23                                      dockerrepo             
docker-engine.src                                         1.12.1-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.12.1-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.12.1-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.12.1-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.12.1-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.12.2-1.fc23                                      dockerrepo             
docker-engine.src                                         1.12.2-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.12.2-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.12.2-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.12.2-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.12.2-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.12.3-1.fc23                                      dockerrepo             
docker-engine.src                                         1.12.3-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.12.3-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.12.3-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.12.3-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.12.3-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.12.4-1.fc23                                      dockerrepo             
docker-engine.src                                         1.12.4-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.12.4-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.12.4-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.12.4-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.12.4-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.12.5-1.fc23                                      dockerrepo             
docker-engine.src                                         1.12.5-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.12.5-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.12.5-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.12.5-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.12.5-1.fc23                                      dockerrepo-local-files 
docker-engine.src                                         1.12.6-1.fc23                                      dockerrepo             
docker-engine.src                                         1.12.6-1.fc23                                      dockerrepo-local-http  
docker-engine.src                                         1.12.6-1.fc23                                      dockerrepo-local-files 
docker-engine.x86_64                                      1.12.6-1.fc23                                      @dockerrepo-local-files
docker-engine.x86_64                                      1.12.6-1.fc23                                      dockerrepo             
docker-engine.x86_64                                      1.12.6-1.fc23                                      dockerrepo-local-http  
docker-engine.x86_64                                      1.12.6-1.fc23                                      dockerrepo-local-files 
[vagrant@localhost ~]$ sudo dnf --showduplicates --disablerepo=kubernetes --enablerepo=dockerrepo-\* downgrade docker-engine-1.11.2-1.fc23
Last metadata expiration check: 0:25:58 ago on Fri Nov 24 09:48:17 2017.
Dependencies resolved.
====================================================================================================================================
 Package                          Arch                      Version                             Repository                     Size
====================================================================================================================================
Downgrading:
 docker-engine                    x86_64                    1.11.2-1.fc23                       dockerrepo                     13 M

Transaction Summary
====================================================================================================================================
Downgrade  1 Package

Total download size: 13 M
Is this ok [y/N]: y
Downloading Packages:
docker-engine-1.11.2-1.fc23.x86_64.rpm                                                               89 kB/s |  13 MB     02:31    
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                89 kB/s |  13 MB     02:32     
warning: /var/cache/dnf/dockerrepo-aceaa77cf60c58c7/packages/docker-engine-1.11.2-1.fc23.x86_64.rpm: Header V4 RSA/SHA512 Signature, key ID 2c52609d: NOKEY
Importing GPG key 0x2C52609D:
 Userid     : "Docker Release Tool (releasedocker) <docker@docker.com>"
 Fingerprint: 5811 8E89 F3A9 1289 7C07 0ADB F762 2157 2C52 609D
 From       : https://yum.dockerproject.org/gpg
Is this ok [y/N]: y
Key imported successfully
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Downgrading : docker-engine-1.11.2-1.fc23.x86_64                                                                              1/2 
  Erasing     : docker-engine-1.12.6-1.fc23.x86_64                                                                              2/2 
  Verifying   : docker-engine-1.11.2-1.fc23.x86_64                                                                              1/2 
  Verifying   : docker-engine-1.12.6-1.fc23.x86_64                                                                              2/2 

Downgraded:
  docker-engine.x86_64 1.11.2-1.fc23                                                                                                

Complete!
```

```
[vagrant@localhost ~]$ docker version
Client:
 Version:      1.11.2
 API version:  1.23
 Go version:   go1.5.4
 Git commit:   b9f10c9
 Built:        Wed Jun  1 21:39:21 2016
 OS/Arch:      linux/amd64
Cannot connect to the Docker daemon. Is the docker daemon running on this host?
```

Using v1.10 config
```
[vagrant@localhost ~]$ sudo systemctl edit docker.service
```

```
[vagrant@localhost ~]$ sudo cat /etc/systemd/system/docker.service.d/override.conf 
[Service]
Environment="OPTIONS=--exec-opt native.cgroupdriver=cgroupfs --cgroup-parent=/docker --insecure-registry=172.30.0.0/16 --insecure-registry=ci.dev.openshift.redhat.com:5000 --selinux-enabled --log-driver=journald"

Environment="DOCKER_CERT_PATH=/etc/docker"

Environment="DOCKER_STORAGE_OPTIONS=-s devicemapper --storage-opt dm.datadev=/dev/vg_vagrant/docker-data --storage-opt dm.metadatadev=/dev/vg_vagrant/docker-metadata --storage-opt dm.use_deferred_removal=true --storage-opt dm.use_deferred_deletion=true"

EnvironmentFile=-/etc/sysconfig/docker.rpmsave
EnvironmentFile=-/etc/sysconfig/docker-storage.rpmsave
#EnvironmentFile=-/etc/sysconfig/docker-network
Environment=GOTRACEBACK=crash
ExecStart=
ExecStart=/usr/bin/docker daemon $OPTIONS $DOCKER_STORAGE_OPTIONS $INSECURE_REGISTRY

```

```
[vagrant@localhost ~]$ sudo systemctl daemon-reload
[vagrant@localhost ~]$ sudo systemctl start docker.service
[vagrant@localhost ~]$ sudo systemctl -l status docker.service
● docker.service - Docker Application Container Engine
   Loaded: loaded (/usr/lib/systemd/system/docker.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/docker.service.d
           └─override.conf
   Active: active (running) since Fri 2017-11-24 10:43:01 UTC; 4s ago
     Docs: https://docs.docker.com
 Main PID: 9685 (docker)
   Memory: 29.9M
      CPU: 642ms
   CGroup: /system.slice/docker.service
           ├─ 9685 /usr/bin/docker daemon --insecure-registry=172.30.0.0/16 --insecure-registry=ci.dev.openshift.redhat.com:5000 --selinux-enabled --log-driver=journald -s devicemapper --storage-opt dm.datadev=/dev/vg_vagrant/docker-data --storage-opt dm.metadatadev=/dev/vg_vagrant/docker-metadata --storage-opt dm.use_deferred_removal=true --storage-opt dm.use_deferred_deletion=true
           ├─ 9688 docker-containerd -l /var/run/docker/libcontainerd/docker-containerd.sock --runtime docker-runc --start-timeout 2m
           ├─ 9802 docker-containerd-shim 3d705da7744ea333cd6e52fe29a7c61368616d45acedcf5e12a5fb7d51d502e1 /var/run/docker/libcontainerd/3d705da7744ea333cd6e52fe29a7c61368616d45acedcf5e12a5fb7d51d502e1 docker-runc
           ├─ 9829 docker-containerd-shim 2bd0c057ae94f73cce613fbca4f49e46f3d8364292810e18ac086a1ecebd1122 /var/run/docker/libcontainerd/2bd0c057ae94f73cce613fbca4f49e46f3d8364292810e18ac086a1ecebd1122 docker-runc
           ├─ 9838 docker-proxy -proto tcp -host-ip 0.0.0.0 -host-port 8082 -container-ip 172.18.0.3 -container-port 8082
           ├─ 9942 docker-containerd-shim 92f7eb502cb8301c2cc73fabccd423100a4a74617d6f5a0c8ecb226c07c97c2a /var/run/docker/libcontainerd/92f7eb502cb8301c2cc73fabccd423100a4a74617d6f5a0c8ecb226c07c97c2a docker-runc
           ├─10138 docker-containerd-shim 79cb57fb10f81d206f1713ce90cc11695b1ecae2df9d805cd259c7769fd00b76 /var/run/docker/libcontainerd/79cb57fb10f81d206f1713ce90cc11695b1ecae2df9d805cd259c7769fd00b76 docker-runc
           └─10202 docker-containerd-shim 6d0d6170e3e6da51630cd2f73e4f3a7d0d15d7435bcfe95179b699ec0c987d3b /var/run/docker/libcontainerd/6d0d6170e3e6da51630cd2f73e4f3a7d0d15d7435bcfe95179b699ec0c987d3b docker-runc
### snippets ###
```

```
[vagrant@localhost ~]$ docker version
Client:
 Version:      1.11.2
 API version:  1.23
 Go version:   go1.5.4
 Git commit:   b9f10c9
 Built:        Wed Jun  1 21:39:21 2016
 OS/Arch:      linux/amd64

Server:
 Version:      1.11.2
 API version:  1.23
 Go version:   go1.5.4
 Git commit:   b9f10c9
 Built:        Wed Jun  1 21:39:21 2016
 OS/Arch:      linux/amd64
```

### kubelet running

but lack cni
```
[vagrant@localhost ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Fri 2017-11-24 05:34:48 UTC; 5h 12min ago
     Docs: http://kubernetes.io/docs/
 Main PID: 23709 (kubelet)
   Memory: 37.5M
      CPU: 7min 3.745s
   CGroup: /system.slice/kubelet.service
           └─23709 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin --cluster-dns=10.96.0.10 --cluster-domain=cluster.local --authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt --cadvisor-port=0 --cgroup-driver=cgroupfs --rotate-certificates=true --cert-dir=/var/lib/kubelet/pki --fail-swap-on=false

Nov 24 10:46:42 kubedev-10-64-33-199 kubelet[23709]: W1124 10:46:42.136656   23709 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 24 10:46:42 kubedev-10-64-33-199 kubelet[23709]: E1124 10:46:42.137412   23709 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Nov 24 10:46:43 kubedev-10-64-33-199 kubelet[23709]: W1124 10:46:43.415232   23709 helpers.go:847] eviction manager: no observation found for eviction signal allocatableNodeFs.available
Nov 24 10:46:47 kubedev-10-64-33-199 kubelet[23709]: W1124 10:46:47.138992   23709 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 24 10:46:47 kubedev-10-64-33-199 kubelet[23709]: E1124 10:46:47.139684   23709 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Nov 24 10:46:52 kubedev-10-64-33-199 kubelet[23709]: W1124 10:46:52.143207   23709 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 24 10:46:52 kubedev-10-64-33-199 kubelet[23709]: E1124 10:46:52.144131   23709 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Nov 24 10:46:53 kubedev-10-64-33-199 kubelet[23709]: W1124 10:46:53.459486   23709 helpers.go:847] eviction manager: no observation found for eviction signal allocatableNodeFs.available
Nov 24 10:46:57 kubedev-10-64-33-199 kubelet[23709]: W1124 10:46:57.148906   23709 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 24 10:46:57 kubedev-10-64-33-199 kubelet[23709]: E1124 10:46:57.151445   23709 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
```

Complete init
```
[vagrant@localhost ~]$ sudo kubeadm init --apiserver-advertise-address 10.64.33.199 --apiserver-bind-port 443 --apiserver-cert-extra-sans 10.64.33.199,172.28.128.3 --pod-network-cidr 10.244.0.0/16 --kubernetes-version v1.8.4 --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.8.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Using the existing ca certificate and key.
[certificates] Using the existing apiserver certificate and key.
[certificates] Using the existing apiserver-kubelet-client certificate and key.
[certificates] Using the existing sa key.
[certificates] Using the existing front-proxy-ca certificate and key.
[certificates] Using the existing front-proxy-client certificate and key.
[certificates] Valid certificates and keys now exist in "/etc/kubernetes/pki"
[kubeconfig] Using existing up-to-date KubeConfig file: "admin.conf"
[kubeconfig] Using existing up-to-date KubeConfig file: "kubelet.conf"
[kubeconfig] Using existing up-to-date KubeConfig file: "controller-manager.conf"
[kubeconfig] Using existing up-to-date KubeConfig file: "scheduler.conf"
[controlplane] Wrote Static Pod manifest for component kube-apiserver to "/etc/kubernetes/manifests/kube-apiserver.yaml"
[controlplane] Wrote Static Pod manifest for component kube-controller-manager to "/etc/kubernetes/manifests/kube-controller-manager.yaml"
[controlplane] Wrote Static Pod manifest for component kube-scheduler to "/etc/kubernetes/manifests/kube-scheduler.yaml"
[etcd] Wrote Static Pod manifest for a local etcd instance to "/etc/kubernetes/manifests/etcd.yaml"
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests"
[init] This often takes around a minute; or longer if the control plane images have to be pulled.
[apiclient] All control plane components are healthy after 19.029924 seconds
[uploadconfig] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[markmaster] Will mark node kubedev-10-64-33-199 as master by adding a label and a taint
[markmaster] Master kubedev-10-64-33-199 tainted and labelled with key/value: node-role.kubernetes.io/master=""
[bootstraptoken] Using token: 6d6b2b.baf2514312b610a3
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

  kubeadm join --token 6d6b2b.baf2514312b610a3 10.64.33.199:443 --discovery-token-ca-cert-hash sha256:cbdb31a6e38373248484ed39c30b72c3dc7c6f1a918868c202579337ee4ec651

```

```
[vagrant@localhost ~]$ export KUBECONFIG=/etc/kubernetes/admin.conf 
[vagrant@localhost ~]$ sudo chmod o+r /etc/kubernetes/admin.conf 
[vagrant@localhost ~]$ kubectl get nodes
NAME                   STATUS     ROLES     AGE       VERSION
kubedev-10-64-33-199   NotReady   master    9m        v1.8.4
```

```
[vagrant@localhost ~]$ kubectl get pods --namespace=kube-system
NAME                                           READY     STATUS    RESTARTS   AGE
etcd-kubedev-10-64-33-199                      1/1       Running   0          2h
kube-apiserver-kubedev-10-64-33-199            1/1       Running   0          2h
kube-controller-manager-kubedev-10-64-33-199   1/1       Running   0          2h
kube-dns-545bc4bfd4-8zvmp                      0/3       Pending   0          2h
kube-proxy-cdqcp                               1/1       Running   0          2h
kube-scheduler-kubedev-10-64-33-199            1/1       Running   0          2h
```

### cni flannel

change iface into kube-flannel.yaml
```
[vagrant@localhost ~]$ kubectl create -f kube-flannel.yaml 
clusterrole "flannel" created
clusterrolebinding "flannel" created
serviceaccount "flannel" created
configmap "kube-flannel-cfg" created
daemonset "kube-flannel-ds" created
```

```
[vagrant@localhost ~]$ kubectl get pods --namespace=kube-system
NAME                                           READY     STATUS    RESTARTS   AGE
etcd-kubedev-10-64-33-199                      1/1       Running   0          3h
kube-apiserver-kubedev-10-64-33-199            1/1       Running   0          2h
kube-controller-manager-kubedev-10-64-33-199   1/1       Running   0          2h
kube-dns-545bc4bfd4-8zvmp                      3/3       Running   0          2h
kube-flannel-ds-d6bb4                          1/1       Running   0          1m
kube-proxy-cdqcp                               1/1       Running   0          2h
kube-scheduler-kubedev-10-64-33-199            1/1       Running   0          2h
```

```
[vagrant@localhost ~]$ kubectl get nodes
NAME                   STATUS    ROLES     AGE       VERSION
kubedev-10-64-33-199   Ready     master    3h        v1.8.4
```

```
[vagrant@kubedev-10-64-33-199 ~]$ kubectl get cm/kubeadm-config -o yaml --namespace=kube-system
apiVersion: v1
data:
  MasterConfiguration: |
    api:
      advertiseAddress: 10.64.33.199
      bindPort: 443
    apiServerCertSANs:
    - 10.64.33.199
    - 172.28.128.3
    authorizationModes:
    - Node
    - RBAC
    certificatesDir: /etc/kubernetes/pki
    cloudProvider: ""
    etcd:
      caFile: ""
      certFile: ""
      dataDir: /var/lib/etcd
      endpoints: null
      image: ""
      keyFile: ""
    imageRepository: gcr.io/google_containers
    kubernetesVersion: v1.8.4
    networking:
      dnsDomain: cluster.local
      podSubnet: 10.244.0.0/16
      serviceSubnet: 10.96.0.0/12
    nodeName: kubedev-10-64-33-199
    token: ""
    tokenTTL: 24h0m0s
    unifiedControlPlaneImage: ""
kind: ConfigMap
metadata:
  creationTimestamp: 2017-11-24T10:51:14Z
  name: kubeadm-config
  namespace: kube-system
  resourceVersion: "831"
  selfLink: /api/v1/namespaces/kube-system/configmaps/kubeadm-config
  uid: 6482c9d1-d105-11e7-b56d-080027242396
```