# Instruction

## Table of contents

Ubuntu 16.04 cloud image

Fedora 27 cloud base (trouble)

## Ubuntu 16.04

Prerequisites

1. Launch from vagrant, [install from VBoxGuestAdditions.ISO](../../examples/rook/vagrantup.md)
1. Install docker
1. Master cluster: 10.64.33.82, 10.64.33.199, 10.64.33.199
1. File server for kubernetes image archive, local YUM repository, master configurations

Hostname
```
[vagrant@localhost ~]$ sudo hostname kubedev-10-64-33-200
```

```
[vagrant@localhost ~]$ echo "kubedev-10-64-33-200" | sudo tee /etc/hostname 
kubedev-10-64-33-200
```

```
[vagrant@localhost ~]$ echo -e "\n10.64.33.200    kubedev-10-64-33-200" | sudo tee -a /etc/hosts

10.64.33.200    kubedev-10-64-33-201
```


pki
```
ubuntu@kubedev-10-64-33-200:~$ scp -r ubuntu@10.64.33.195:/home/ubuntu/etc0x2Fkubernetes .
The authenticity of host '10.64.33.195 (10.64.33.195)' can't be established.
ECDSA key fingerprint is SHA256:QQBkXG5+IgQTMn7rh2aTbJi+IQLE106NbIBXJqKIoPI.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '10.64.33.195' (ECDSA) to the list of known hosts.
kubelet.conf                                                                                                  100% 5539     5.4KB/s   00:00    
controller-manager.conf                                                                                       100% 5487     5.4KB/s   00:00    
sa.key                                                                                                        100% 1679     1.6KB/s   00:00    
apiserver.crt                                                                                                 100% 1387     1.4KB/s   00:00    
front-proxy-ca.crt                                                                                            100% 1025     1.0KB/s   00:00    
front-proxy-client.crt                                                                                        100% 1050     1.0KB/s   00:00    
apiserver.key                                                                                                 100% 1679     1.6KB/s   00:00    
ca.key                                                                                                        100% 1679     1.6KB/s   00:00    
ca.crt                                                                                                        100% 1025     1.0KB/s   00:00    
apiserver-kubelet-client.crt                                                                                  100% 1099     1.1KB/s   00:00    
apiserver-kubelet-client.key                                                                                  100% 1679     1.6KB/s   00:00    
front-proxy-client.key                                                                                        100% 1675     1.6KB/s   00:00    
front-proxy-ca.key                                                                                            100% 1675     1.6KB/s   00:00    
sa.pub                                                                                                        100%  451     0.4KB/s   00:00    
kube-scheduler.yaml                                                                                           100%  991     1.0KB/s   00:00    
kube-apiserver.yaml                                                                                           100% 2619     2.6KB/s   00:00    
kube-controller-manager.yaml                                                                                  100% 2232     2.2KB/s   00:00    
scheduler.conf                                                                                                100% 5435     5.3KB/s   00:00    
admin.conf                                                                                                    100% 5447     5.3KB/s   00:00    
```

images
```
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-apiserver.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  185M  100  185M    0     0  82.1M      0  0:00:02  0:00:02 --:--:-- 82.1M
0271b8eebde3: Loading layer [==================================================>] 1.338 MB/1.338 MB
f72e92775dd7: Loading layer [==================================================>] 193.2 MB/193.2 MB
Loaded image: gcr.io/google_containers/kube-apiserver:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-controller-manager.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  123M  100  123M    0     0  94.1M      0  0:00:01  0:00:01 --:--:-- 94.1M
4f92df6d8677: Loading layer [==================================================>] 128.2 MB/128.2 MB
Loaded image: gcr.io/google_containers/kube-controller-manager:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-scheduler.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 52.6M  100 52.6M    0     0  94.1M      0 --:--:-- --:--:-- --:--:-- 94.0M
da6851a1e488: Loading layer [==================================================>] 53.85 MB/53.85 MB
Loaded image: gcr.io/google_containers/kube-scheduler:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-proxy.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.6M  100 90.6M    0     0  73.2M      0  0:00:01  0:00:01 --:--:-- 73.3M
08ae86c4c4c9: Loading layer [==================================================>] 42.05 MB/42.05 MB
48a108f012f8: Loading layer [==================================================>] 5.045 MB/5.045 MB
9fc6ccb688b9: Loading layer [==================================================>] 47.93 MB/47.93 MB
Loaded image: gcr.io/google_containers/kube-proxy:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-kube0x3A2.1.1.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.8M  100 90.8M    0     0  10.0M      0  0:00:09  0:00:09 --:--:-- 6349k
5bef08742407: Loading layer [==================================================>] 4.221 MB/4.221 MB
a740a6d19f48: Loading layer [==================================================>] 18.88 MB/18.88 MB
d86a8ab219a9: Loading layer [==================================================>]    27 MB/27 MB
f680fc950fce: Loading layer [==================================================>] 10.26 MB/10.26 MB
4f89cf228c9b: Loading layer [==================================================>] 2.048 kB/2.048 kB
802751b045ac: Loading layer [==================================================>]   277 kB/277 kB
d56c780427c7: Loading layer [==================================================>] 34.62 MB/34.62 MB
Loaded image: weaveworks/weave-kube:2.1.1
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-npc0x3A2.1.1.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 44.7M  100 44.7M    0     0  41.1M      0  0:00:01  0:00:01 --:--:-- 41.2M
8dccfe2dec8c: Loading layer [==================================================>] 2.811 MB/2.811 MB
11565a7c0730: Loading layer [==================================================>] 39.91 MB/39.91 MB
f56503d36fe6: Loading layer [==================================================>]  2.56 kB/2.56 kB
Loaded image: weaveworks/weave-npc:2.1.1
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fetcd0x2E3.0.17.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  161M  100  161M    0     0  32.5M      0  0:00:04  0:00:04 --:--:-- 33.3M
38ac8d0f5bb3: Loading layer [==================================================>] 1.312 MB/1.312 MB
c872b2c2ac77: Loading layer [==================================================>] 136.7 MB/136.7 MB
be71b2a80bbd: Loading layer [==================================================>] 31.16 MB/31.16 MB
Loaded image: gcr.io/google_containers/etcd:3.0.17
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-dnsmasq-nanny-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 39.7M  100 39.7M    0     0  47.2M      0 --:--:-- --:--:-- --:--:-- 47.1M
e620d0ac6539: Loading layer [==================================================>]  2.56 kB/2.56 kB
9f4f5a30979e: Loading layer [==================================================>]   362 kB/362 kB
fd7319ac31de: Loading layer [==================================================>] 3.072 kB/3.072 kB
b23d166217e1: Loading layer [==================================================>]  37.1 MB/37.1 MB
Loaded image: gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-kube-dns-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 47.3M  100 47.3M    0     0  47.8M      0 --:--:-- --:--:-- --:--:-- 47.8M
a1a7a797fc0e: Loading layer [==================================================>] 45.42 MB/45.42 MB
Loaded image: gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.5
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-sidecar-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 40.1M  100 40.1M    0     0  43.9M      0 --:--:-- --:--:-- --:--:-- 43.9M
acfc450a47fa: Loading layer [==================================================>] 37.86 MB/37.86 MB
Loaded image: gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.5
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  747k  100  747k    0     0  7623k      0 --:--:-- --:--:-- --:--:-- 7706k
5f70bf18a086: Loading layer [==================================================>] 1.024 kB/1.024 kB
41ff149e94f2: Loading layer [==================================================>] 748.5 kB/748.5 kB
Loaded image: gcr.io/google_containers/pause-amd64:3.0
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/quay.io0x2Fcoreos0x2Fflannel0x3Av0.9.0-amd64.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 49.7M  100 49.7M    0     0  29.3M      0  0:00:01  0:00:01 --:--:-- 29.4M
f439636ab0f0: Loading layer [==================================================>] 6.797 MB/6.797 MB
91b6f6ead101: Loading layer [==================================================>] 4.414 MB/4.414 MB
39f837629582: Loading layer [==================================================>] 34.49 MB/34.49 MB
d3e99a0118c5: Loading layer [==================================================>] 2.225 MB/2.225 MB
32adca76eade: Loading layer [==================================================>]  5.12 kB/5.12 kB
Loaded image: quay.io/coreos/flannel:v0.9.0-amd64
ubuntu@kubedev-10-64-33-200:~$ docker images
REPOSITORY                                             TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver                v1.8.4              10a052dccbc5        5 days ago          194 MB
gcr.io/google_containers/kube-controller-manager       v1.8.4              7058ac4d4af5        5 days ago          129 MB
gcr.io/google_containers/kube-proxy                    v1.8.4              65a61c14e8c2        5 days ago          93.2 MB
gcr.io/google_containers/kube-scheduler                v1.8.4              0d985fed7f95        5 days ago          55 MB
weaveworks/weave-npc                                   2.1.1               70a3a81a2ad5        8 days ago          46.6 MB
weaveworks/weave-kube                                  2.1.1               bc65281cfd26        8 days ago          92.6 MB
gcr.io/google_containers/k8s-dns-sidecar-amd64         1.14.5              fed89e8b4248        8 weeks ago         41.8 MB
gcr.io/google_containers/k8s-dns-kube-dns-amd64        1.14.5              512cd7425a73        8 weeks ago         49.4 MB
gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64   1.14.5              459944ce8cc4        8 weeks ago         41.4 MB
quay.io/coreos/flannel                                 v0.9.0-amd64        4c600a64a18a        2 months ago        51.3 MB
gcr.io/google_containers/etcd                          3.0.17              243830dae7dd        9 months ago        169 MB
gcr.io/google_containers/pause-amd64                   3.0                 99e59f495ffa        19 months ago       747 kB
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/kube-apiserver:v1.8.4 gcr.io/google_containers/kube-apiserver-amd64:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/kube-controller-manager:v1.8.4 gcr.io/google_containers/kube-controller-manager-amd64:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/kube-scheduler:v1.8.4 gcr.io/google_containers/kube-scheduler-amd64:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/kube-proxy:v1.8.4 gcr.io/google_containers/kube-proxy-amd64:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/etcd:3.0.17 gcr.io/google_containers/etcd-amd64:3.0.17
```

deploy
```
ubuntu@kubedev-10-64-33-200:~$ sudo kubeadm join --token 90d3e2.adc22e3d43b3faa2 10.64.33.82:443 --discovery-token-ca-cert-hash sha256:92202ed5b88a724e1206b95dff886ddce6cc67628886b2eeafb95be4a5d888b6 --skip-preflight-checks 
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[preflight] Skipping pre-flight checks
[discovery] Trying to connect to API Server "10.64.33.82:443"
[discovery] Created cluster-info discovery client, requesting info from "https://10.64.33.82:443"
[discovery] Requesting info from "https://10.64.33.82:443" again to validate TLS against the pinned public key
[discovery] Cluster info signature and contents are valid and TLS certificate validates against pinned roots, will use API Server "10.64.33.82:443"
[discovery] Successfully established connection with API Server "10.64.33.82:443"
[bootstrap] Detected server version: v1.8.4
[bootstrap] The server supports the Certificates API (certificates.k8s.io/v1beta1)

Node join complete:
* Certificate signing request sent to master and response
  received.
* Kubelet informed of new secure connection details.

Run 'kubectl get nodes' on the master to see this machine join.
```

```
ubuntu@kubedev-10-64-33-200:~$ export KUBECONFIG=/home/ubuntu/etc0x2Fkubernetes/admin.conf 
ubuntu@kubedev-10-64-33-200:~$ kubectl get nodes
NAME                   STATUS    ROLES     AGE       VERSION
kubedev-10-64-33-195   Ready     master    7h        v1.8.4
kubedev-10-64-33-199   Ready     master    8h        v1.8.4
kubedev-10-64-33-200   Ready     <none>    57s       v1.8.4
kubedev-10-64-33-82    Ready     master    1h        v1.8.4
```

```
ubuntu@kubedev-10-64-33-200:~$ kubectl get pods -o wide -n kube-system | grep kubedev-10-64-33-200
kube-flannel-ds-r74c6                          0/1       CrashLoopBackOff   5          5m        10.64.33.200   kubedev-10-64-33-200
kube-proxy-nz8dz                               1/1       Running            0          5m        10.64.33.200   kubedev-10-64-33-200
```


## Fedora27

Prerequisites

1. Launch from vagrant, [install from VBoxGuestAdditions.ISO](../../examples/rook/vagrantup.md)
1. Install docker
1. Master cluster: 10.64.33.82, 10.64.33.199, 10.64.33.199
1. File server for kubernetes image archive, local YUM repository, master configurations

Hostname
```
[vagrant@localhost ~]$ sudo hostname rookdev-10-64-33-201
```

```
[vagrant@localhost ~]$ echo "rookdev-10-64-33-201" | sudo tee /etc/hostname 
rookdev-10-64-33-201
```

```
[vagrant@localhost ~]$ echo -e "\n10.64.33.201    rookdev-10-64-33-201" | sudo tee -a /etc/hosts

10.64.33.201    rookdev-10-64-33-201
```

SELinux
```
[vagrant@localhost ~]$ sudo setenforce 0
```

```
[vagrant@localhost ~]$ getenforce 
Permissive
```

To enable permanently, modify /etc/sysconfig/selinux, like
```
$ sudo sed -i 's/SELINUX=enforcing/SELINUX=permissive/' /etc/sysconfig/selinux
```

_Note, fedora 27 seems to trouble that it is alway Enforcing after reboot_

Kubernetes local YUM repository
```
[vagrant@localhost ~]$ echo -e "[kubernetes-local-http]\nname=Kubernetes local http\nbaseurl=http://10.64.33.82:48080/windows10_drive_f/16-mirror/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/\nenabled=0\ngpgcheck=0\n\n[kubernetes-local-files]\nname=Kubernetes local files\nbaseurl=file:///windows10_drive_f/16-mirror/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/\nenabled=0\ngpgcheck=0\n" | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes-local-http]
name=Kubernetes local http
baseurl=http://10.64.33.82:48080/windows10_drive_f/16-mirror/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/
enabled=0
gpgcheck=0

[kubernetes-local-files]
name=Kubernetes local files
baseurl=file:///windows10_drive_f/16-mirror/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/
enabled=0
gpgcheck=0

```

```
[vagrant@localhost ~]$ sudo dnf search --enablerepo=kubernetes\* kubeadm kubelet kubectl
Failed to synchronize cache for repo 'kubernetes-local-files', disabling.
Last metadata expiration check: 0:00:13 ago on Sat 16 Dec 2017 02:42:56 AM UTC.
======================================================== Name Exactly Matched: kubelet =========================================================
kubelet.x86_64 : Container cluster management
======================================================== Name Exactly Matched: kubectl =========================================================
kubectl.x86_64 : Command-line utility for interacting with a Kubernetes cluster.
======================================================== Name Exactly Matched: kubeadm =========================================================
kubeadm.x86_64 : Command-line utility for administering a Kubernetes cluster.
============================================================ Name Matched: kubeadm =============================================================
kubernetes-kubeadm.x86_64 : Kubernetes tool for standing up clusters
```

Install kubeadm
```
[vagrant@localhost ~]$ sudo dnf install --enablerepo=kubernetes\* -y kubeadm
Failed to synchronize cache for repo 'kubernetes-local-files', disabling.
Last metadata expiration check: 0:01:09 ago on Sat 16 Dec 2017 02:42:56 AM UTC.
Dependencies resolved.
================================================================================================================================================
 Package                           Arch                      Version                             Repository                                Size
================================================================================================================================================
Installing:
 kubeadm                           x86_64                    1.8.4-0                             kubernetes-local-http                     15 M
Installing dependencies:
 ebtables                          x86_64                    2.0.10-24.fc27                      fedora                                   134 k
 ethtool                           x86_64                    2:4.13-1.fc27                       updates                                  138 k
 kubectl                           x86_64                    1.8.4-0                             kubernetes-local-http                    7.3 M
 kubelet                           x86_64                    1.8.4-0                             kubernetes-local-http                     16 M
 kubernetes-cni                    x86_64                    0.5.1-1                             kubernetes-local-http                    7.4 M
 socat                             x86_64                    1.7.3.2-4.fc27                      fedora                                   296 k

Transaction Summary
================================================================================================================================================
Install  7 Packages

Total download size: 46 M
Installed size: 244 M
Downloading Packages:
(1/7): a9db28728641ddbf7f025b8b496804d82a396d0ccb178fffd124623fb2f999ea-kubectl-1.8.4-0.x86_64.rpm              8.0 MB/s | 7.3 MB     00:00    
(2/7): aeaad1e283c54876b759a089f152228d7cd4c049f271125c23623995b8e76f96-kubeadm-1.8.4-0.x86_64.rpm               11 MB/s |  15 MB     00:01    
(3/7): 1acca81eb5cf99453f30466876ff03146112b7f12c625cb48f12508684e02665-kubelet-1.8.4-0.x86_64.rpm               11 MB/s |  16 MB     00:01    
(4/7): ebtables-2.0.10-24.fc27.x86_64.rpm                                                                       198 kB/s | 134 kB     00:00    
(5/7): socat-1.7.3.2-4.fc27.x86_64.rpm                                                                          719 kB/s | 296 kB     00:00    
(6/7): ethtool-4.13-1.fc27.x86_64.rpm                                                                           790 kB/s | 138 kB     00:00    
(7/7): 79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm       6.9 MB/s | 7.4 MB     00:01    
------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                           4.2 MB/s |  46 MB     00:10     
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
  Installing       : kubernetes-cni-0.5.1-1.x86_64                                                                                          4/7 
  Installing       : kubelet-1.8.4-0.x86_64                                                                                                 5/7 
  Installing       : kubectl-1.8.4-0.x86_64                                                                                                 6/7 
  Installing       : kubeadm-1.8.4-0.x86_64                                                                                                 7/7 
  Running scriptlet: kubeadm-1.8.4-0.x86_64                                                                                                 7/7 
  Verifying        : kubeadm-1.8.4-0.x86_64                                                                                                 1/7 
  Verifying        : kubectl-1.8.4-0.x86_64                                                                                                 2/7 
  Verifying        : kubelet-1.8.4-0.x86_64                                                                                                 3/7 
  Verifying        : ebtables-2.0.10-24.fc27.x86_64                                                                                         4/7 
  Verifying        : socat-1.7.3.2-4.fc27.x86_64                                                                                            5/7 
  Verifying        : kubernetes-cni-0.5.1-1.x86_64                                                                                          6/7 
  Verifying        : ethtool-2:4.13-1.fc27.x86_64                                                                                           7/7 

Installed:
  kubeadm.x86_64 1.8.4-0         ebtables.x86_64 2.0.10-24.fc27  ethtool.x86_64 2:4.13-1.fc27  kubectl.x86_64 1.8.4-0  kubelet.x86_64 1.8.4-0 
  kubernetes-cni.x86_64 0.5.1-1  socat.x86_64 1.7.3.2-4.fc27    

Complete!
```

_Upon installation have dependencies from fedora and updates repository, it is also installed from mirror, e.g. ISO_


Kubernetes image archive from local file server
```
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-apiserver.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  185M  100  185M    0     0  20.6M      0  0:00:09  0:00:09 --:--:-- 19.3M
0271b8eebde3: Loading layer [==================================================>] 1.338 MB/1.338 MB
f72e92775dd7: Loading layer [==================================================>] 193.2 MB/193.2 MB
Loaded image: gcr.io/google_containers/kube-apiserver:v1.8.4
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-controller-manager.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  123M  100  123M    0     0  12.3M      0  0:00:10  0:00:10 --:--:-- 14.7M
4f92df6d8677: Loading layer [==================================================>] 128.2 MB/128.2 MB
Loaded image: gcr.io/google_containers/kube-controller-manager:v1.8.4
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-scheduler.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 52.6M  100 52.6M    0     0  10.5M      0  0:00:05  0:00:05 --:--:-- 9627k
da6851a1e488: Loading layer [==================================================>] 53.85 MB/53.85 MB
Loaded image: gcr.io/google_containers/kube-scheduler:v1.8.4
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-proxy.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.6M  100 90.6M    0     0  22.6M      0  0:00:04  0:00:04 --:--:-- 21.0M
08ae86c4c4c9: Loading layer [==================================================>] 42.05 MB/42.05 MB
48a108f012f8: Loading layer [==================================================>] 5.045 MB/5.045 MB
9fc6ccb688b9: Loading layer [==================================================>] 47.93 MB/47.93 MB
Loaded image: gcr.io/google_containers/kube-proxy:v1.8.4
```

POD image archive
```
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  747k  100  747k    0     0   747k      0  0:00:01 --:--:--  0:00:01 2942k
5f70bf18a086: Loading layer [==================================================>] 1.024 kB/1.024 kB
41ff149e94f2: Loading layer [==================================================>] 748.5 kB/748.5 kB
Loaded image: gcr.io/google_containers/pause-amd64:3.0
```

Etcd image archive
```
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fetcd0x3A3.0.17.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  161M  100  161M    0     0  20.1M      0  0:00:08  0:00:08 --:--:-- 23.8M
38ac8d0f5bb3: Loading layer [==================================================>] 1.312 MB/1.312 MB
c872b2c2ac77: Loading layer [==================================================>] 136.7 MB/136.7 MB
be71b2a80bbd: Loading layer [==================================================>] 31.16 MB/31.16 MB
Loaded image: gcr.io/google_containers/etcd:3.0.17
```

Addon kube-dns image archive
```
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-dnsmasq-nanny-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 39.7M  100 39.7M    0     0  13.2M      0  0:00:03  0:00:03 --:--:--  9.9M
5bef08742407: Loading layer [==================================================>] 4.221 MB/4.221 MB
e620d0ac6539: Loading layer [==================================================>]  2.56 kB/2.56 kB
9f4f5a30979e: Loading layer [==================================================>]   362 kB/362 kB
fd7319ac31de: Loading layer [==================================================>] 3.072 kB/3.072 kB
b23d166217e1: Loading layer [==================================================>]  37.1 MB/37.1 MB
Loaded image: gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
```

```
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-kube-dns-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 47.3M  100 47.3M    0     0  11.8M      0  0:00:04  0:00:04 --:--:-- 10.9M
a1a7a797fc0e: Loading layer [==================================================>] 45.42 MB/45.42 MB
Loaded image: gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.5
```

```
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-sidecar-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 40.1M  100 40.1M    0     0  13.3M      0  0:00:03  0:00:03 --:--:-- 10.7M
acfc450a47fa: Loading layer [==================================================>] 37.86 MB/37.86 MB
Loaded image: gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.5
```

CNI flannel image archive
```
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/docker-images/quay.io0x2Fcoreos0x2Fflannel0x3Av0.9.0-amd64.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 49.7M  100 49.7M    0     0   9.9M      0  0:00:05  0:00:05 --:--:-- 8150k
f439636ab0f0: Loading layer [==================================================>] 6.797 MB/6.797 MB
91b6f6ead101: Loading layer [==================================================>] 4.414 MB/4.414 MB
39f837629582: Loading layer [==================================================>] 34.49 MB/34.49 MB
d3e99a0118c5: Loading layer [==================================================>] 2.225 MB/2.225 MB
32adca76eade: Loading layer [==================================================>]  5.12 kB/5.12 kB
Loaded image: quay.io/coreos/flannel:v0.9.0-amd64
```

CNI weave image archive
```
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-kube0x3A2.1.1.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.8M  100 90.8M    0     0  11.3M      0  0:00:08  0:00:08 --:--:-- 11.0M
a740a6d19f48: Loading layer [==================================================>] 18.88 MB/18.88 MB
d86a8ab219a9: Loading layer [==================================================>]    27 MB/27 MB
f680fc950fce: Loading layer [==================================================>] 10.26 MB/10.26 MB
4f89cf228c9b: Loading layer [==================================================>] 2.048 kB/2.048 kB
802751b045ac: Loading layer [==================================================>]   277 kB/277 kB
d56c780427c7: Loading layer [==================================================>] 34.62 MB/34.62 MB
Loaded image: docker.io/weaveworks/weave-kube:2.1.1
```

```
[vagrant@localhost ~]$ curl -L http://10.64.33.82:48080/windows10_drive_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-npc0x3A2.1.1.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 44.7M  100 44.7M    0     0  9173k      0  0:00:05  0:00:05 --:--:-- 7348k
8dccfe2dec8c: Loading layer [==================================================>] 2.811 MB/2.811 MB
11565a7c0730: Loading layer [==================================================>] 39.91 MB/39.91 MB
f56503d36fe6: Loading layer [==================================================>]  2.56 kB/2.56 kB
Loaded image: docker.io/weaveworks/weave-npc:2.1.1
```

Tag image
```
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/kube-apiserver:v1.8.4 gcr.io/google_containers/kube-apiserver-amd64:v1.8.4
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/kube-controller-manager:v1.8.4 gcr.io/google_containers/kube-controller-manager-amd64:v1.8.4
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/kube-scheduler:v1.8.4 gcr.io/google_containers/kube-scheduler-amd64:v1.8.4
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/kube-proxy:v1.8.4 gcr.io/google_containers/kube-proxy-amd64:v1.8.4
[vagrant@localhost ~]$ docker tag gcr.io/google_containers/etcd:3.0.17 gcr.io/google_containers/etcd-amd64:3.0.17
```

PKI from master
```
[vagrant@localhost ~]$ for i in "apiserver.crt" "apiserver.key" "apiserver-kubelet-client.crt" "apiserver-kubelet-client.key" "ca.crt" "ca.key" "front-proxy-ca.crt" "front-proxy-ca.key" "front-proxy-client.crt" "front-proxy-client.key" "sa.key" "sa.pub"; do curl -L http://10.64.33.82:48080/etc/kubernetes/pki/$i | sudo tee /etc/kubernetes/pki/$i; done; sudo chmod -R 400 /etc/kubernetes/pki/
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1387  100  1387    0     0   1387      0  0:00:01 --:--:--  0:00:01  225k
-----BEGIN CERTIFICATE-----
MIID1DCCArygAwIBAgIIE++TmZEaUs4wDQYJKoZIhvcNAQELBQAwFTETMBEGA1UE
AxMKa3ViZXJuZXRlczAeFw0xNzExMjUxODE5MzBaFw0xODExMjUxODE5MzFaMBkx
FzAVBgNVBAMTDmt1YmUtYXBpc2VydmVyMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEAscFpT/0woQJtsdVMBU/jS4lLb5lxU7xv7NiRJOiCZV0JxkHwQNoB
QU5NKW3TT2Ln6HSuPL8X85ZbSTwKChcEQEPRJIA55o7esNy96jpHiZACo3DqsUps
vdUjbQPVDuQt4mfnieqfKn1I9DotUfEJazXfKtylpM9nYNksFFDAdcWL/tqIiiBP
/1zjpB3vBewndN9IxxrAqqHaL8b43YG8GQd6KzLRfjwSqpSJhE38LoWXE5tNYUuC
cq8HDYG0QMzZmNsGQcqlWnXLgv9ewAEwQSFo7mzu87rkDwZYGjdSQb2AyjbdRylC
+VPEVSNY55cVgDuqVMBX0oUSy6+VS+SHlwIDAQABo4IBIjCCAR4wDgYDVR0PAQH/
BAQDAgWgMBMGA1UdJQQMMAoGCCsGAQUFBwMBMIH2BgNVHREEge4wgeuCFGt1YmVk
ZXYtMTAtNjQtMzMtMTk5ggprdWJlcm5ldGVzghJrdWJlcm5ldGVzLmRlZmF1bHSC
Fmt1YmVybmV0ZXMuZGVmYXVsdC5zdmOCJGt1YmVybmV0ZXMuZGVmYXVsdC5zdmMu
Y2x1c3Rlci5sb2NhbIITa3ViZWRldi0xMC02NC0zMy04MoIUa3ViZWRldi0xMC02
NC0zMy0xOTmCFGt1YmVkZXYtMTAtNjQtMzMtMTk1hwQKYAABhwQKQCHHhwQKQCFS
hwQKQCHHhwQKQCHDhwQKQCHghwQKQCHwhwQKQCEyhwSsHIAyMA0GCSqGSIb3DQEB
CwUAA4IBAQA8nfdpikHP7cXpbMVQqcfUvhxAe8Nu6Mw162hDZSy1fCElqbpflG+F
KdX/GWbDVYNfrf83ZxNZl4Xx6NVrSMcdzOwZgZPkmXUYYb0UvaPiP2i30zKJL5rZ
ClLlm+Gw+MjYvYKWl18cdd5G2Z91Oo6yJ3qVWYPGjnNM86hpvPPuqOm/nzfDdy3Q
LVeOTm87WXSmrhEmOXSv4snMgFQ/nGMo+CX4AYmrR+Y0JtMdUwV4na6fVx/J2deL
ZTF/ZyhCAZYrPOtKfyljU9XoAD0EQXKv209+55tzQhqhBquY8R0ob3WiJTw7gmG7
GADfk5c8zq1gsxd9G95gGPkx+hsQ5knE
-----END CERTIFICATE-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1679  100  1679    0     0   1679      0  0:00:01 --:--:--  0:00:01  409k
-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEAscFpT/0woQJtsdVMBU/jS4lLb5lxU7xv7NiRJOiCZV0JxkHw
QNoBQU5NKW3TT2Ln6HSuPL8X85ZbSTwKChcEQEPRJIA55o7esNy96jpHiZACo3Dq
sUpsvdUjbQPVDuQt4mfnieqfKn1I9DotUfEJazXfKtylpM9nYNksFFDAdcWL/tqI
iiBP/1zjpB3vBewndN9IxxrAqqHaL8b43YG8GQd6KzLRfjwSqpSJhE38LoWXE5tN
YUuCcq8HDYG0QMzZmNsGQcqlWnXLgv9ewAEwQSFo7mzu87rkDwZYGjdSQb2Ayjbd
RylC+VPEVSNY55cVgDuqVMBX0oUSy6+VS+SHlwIDAQABAoIBAQCK1jpfU67L043C
PIsoCwHJWe8lt3+gm6oAIBcd8M0IRub/hzkULjKUw4fZOAK4GtRv59K/d7JJe8LJ
LK/auaeEWnPcQpr2zdh547AiZmrBs02cQvHMkAjnLLaCy196Rhbgo69hNXQkR77X
oykKu21T47vvSm/Gjoh8SKP5Kilo0hn6b72AwVPLsTih3GjS5U8DnHeEPOsYGCU6
+52auzSGSsBc/0FKQoHa6c8bmFa4gTyGYZRWeUG/THclAtD+IZGlHKr/DDhHGL8p
nSOTf8qw0ZygRDlgKmM+TiGecVPL16Dk2ADJgnX2D0No+qkacrDQizkrHoQPEdTO
QQp3S+zhAoGBANn9s1R5umj+fRJY9WnCJk+GSzPW+XGhfwaXwnxQC4Bx6PCX1ChD
c1O9zsDdRP7TypPSvsliH0dVm6XPnM+bHC2WIqJbwqAeGaBeC6N2Pbz+9V5ykzvy
iM4o8K8AJankv3vlxS573KOc9b2TMxNFiB6XGMDFF7eftcwU/YVq6fRtAoGBANC/
v8yoR41v9IpxCr17PDOg+5xEJcsKKPs6OzdnlsOtXUdho9iIe79HIqzrQP3AExF1
ZzbrHr8kSUzenzoQAJJfB8SuwyXOjJatgKie3dhwa95iUfWO7heyGwWuP7y9/3St
G0tc/0l6Y2Fbi1Fj56iyOStDzXBCYkVz81+tA8GTAoGBAJ1Uoc1iXcvRgOtxEQ7X
dVcIEdbEXRxJ9qgHd/LIwsRdckXTDuhQZjgWUdaY3Gce7KJsBA7Mps6RVtkBUqfz
xw6PwUd0+q8dillYBwZfHIf/a5BTGTi2/03r9moXrRdFZX5Cg2HdO41Kh4Hfb4mF
MY0CeHcIyOlad5dnvuJAG8KxAoGBAJhcfEeKQ79uCZTKcUZfnk8onr0t2E/b+JF1
rScBSr9D2CNzlr1C0N+Jsjt38hXDjadeg5lH/0cP2xhqHNmqOGCdT/Q0gmcj3OVV
eeSZH98MHw9Pv/Z0/j75VlPNHORL4RpuBs/47rPYypSS1prvSnSMsMFtWZMoQMXk
nawlgepVAoGBAIyKZdu0W553sk85G8cSHeU9GbGmpqshGroxIcLow6429x+an8wM
KFl6/KTGOt9o8Slsyqavo/Rs6JEXFQpjRuwCLlKfq3jaoDzB6ZrHxP5N+ZOVyje3
PbDQ4hHh7nTGD3mG2eW6jDbhq7NG8DDKeK8hWvYD4KFQv4Ij7OBMzJ89
-----END RSA PRIVATE KEY-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1099  100  1099    0     0   1099      0  0:00:01 --:--:--  0:00:01 78500
-----BEGIN CERTIFICATE-----
MIIC/zCCAeegAwIBAgIIeA5JACR+pfQwDQYJKoZIhvcNAQELBQAwFTETMBEGA1UE
AxMKa3ViZXJuZXRlczAeFw0xNzExMjUxODE5MzBaFw0xODExMjUxODE5MzFaMEEx
FzAVBgNVBAoTDnN5c3RlbTptYXN0ZXJzMSYwJAYDVQQDEx1rdWJlLWFwaXNlcnZl
ci1rdWJlbGV0LWNsaWVudDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AOIyzCPky3SapHBjV7jKURs6fyNf5MrMhNTGm85vmh0rjPyoz5p6F05YxIGXBFPN
33beGS/5GuYxA3oH97wNRZNLitsDFZS1qNK3b/1JmJhlRmxY/mLpVJbUUiRdqvZL
zn041Uep0aoZPyFyd4jT8q7r9zsjiLGu0cabcU+yGT7viDL9gQ/KrVlOiQdeW2NM
7xCHd1CfclRzP3cljdM8KDa+0kT3aQ9VxgeImPlzlzV44Y36EvryOt3qznIHEtwx
Cb5UolbrXKj7ZdsyO4YURpNBsyC8lQ5gYtn06kuAhK0RFACfMbVHm/ASMDaY84UZ
pwB8g9jP5wyiJ1rcKqZy+5cCAwEAAaMnMCUwDgYDVR0PAQH/BAQDAgWgMBMGA1Ud
JQQMMAoGCCsGAQUFBwMCMA0GCSqGSIb3DQEBCwUAA4IBAQCRydgmSOqrwax9KWXF
GskUavnyfhFzHE5YTBa0WOHzspnj3dvvooR2BIC6VGS7se4eIkrKfhEIuP8gYoUf
zxMSxNtXCLD/CwE9/ccf2qNmk8DgVP5wvkMXjN/3Tvc24nnswfNJ9QNM+/oqxjhR
rQU9VaREP1RXUPw9kYfOv3/2PxV91aLri//HPrJBEl1cNAKgsArwtyRVdOSuwQNW
nDO1VHjW+2ze/fXu20LUopb/ywCsPTrCUuitOMWiCfFk8TaNAWbDl2aRMjJLvP75
B/CLLRJc/LSnkzQmjAAFobxIT2HMTcZNUoiKJ4U9a7cdebNUciarOtYUmWs+tv8q
HzrA
-----END CERTIFICATE-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1679  100  1679    0     0   1679      0  0:00:01 --:--:--  0:00:01  327k
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA4jLMI+TLdJqkcGNXuMpRGzp/I1/kysyE1Mabzm+aHSuM/KjP
mnoXTljEgZcEU83fdt4ZL/ka5jEDegf3vA1Fk0uK2wMVlLWo0rdv/UmYmGVGbFj+
YulUltRSJF2q9kvOfTjVR6nRqhk/IXJ3iNPyruv3OyOIsa7RxptxT7IZPu+IMv2B
D8qtWU6JB15bY0zvEId3UJ9yVHM/dyWN0zwoNr7SRPdpD1XGB4iY+XOXNXjhjfoS
+vI63erOcgcS3DEJvlSiVutcqPtl2zI7hhRGk0GzILyVDmBi2fTqS4CErREUAJ8x
tUeb8BIwNpjzhRmnAHyD2M/nDKInWtwqpnL7lwIDAQABAoIBAQDCsmqYMqUVFs6X
lkq5Wht5zl4EutOZ43QqGVGe3NzQYt4V08/t7U0GOM3bpXsqDzhbfkw7yLTokiOl
9Bm+mcXbRNieLwdyWvKXfSE96Q0Zj4KQXAGKbcBLmzGg76bqHNH9Yt/hsUAwo059
oaNV+OkYuy6wjqNoHUufIcjkii3ExpYjBc8GiYjsnFJqR7uD0cJP7GcYd60Kt0XJ
QKr4TIcBPcKmwq4gkM02JgGlhn+9t03jR+ITTbHJIzgOZ8O9ZwIzMV9Zi45/0cyH
wVremor+/ayoNhEm1ayXIt+0wiH7efr6mdTO6dpenA8lfcyRfjO1RXHNZ3w7yZ5/
6y4mna3hAoGBAOvISoMKGh/XuDoBCUegmmdBnuHOxEu6WxxDQyztFAbLTz0f1zHo
J0qGTGu6Zvd3IM4/jm1AlAx68NCgHsOmjxbm8mkXj5SfVAYS9rGU1nr9B5yhopxz
yifXxvBKbJA0g7mCrig2L/rqJw7kShPwzmnI/EmaJ6PxLAonLxpIU0CpAoGBAPWY
IHi+I0ftAoNPqPGwkzmyuRnIZ0icBF3oDtVvih0OEo8bWoSrQdfqnZzls1Vqdtii
QE96hoeq3b2D7H4zqzCv4pa4oiHlnn7vNQmMNu9w1D4+Gh6FXxyy31W0ba5SNq5K
sUctI3DOodSJg/z7IEZp4PfR3JSGouCqEfKnBsI/AoGAeM20L5N7cxkpkL57LLHC
qmjaqMMxm8MtbzufLWOCwUBlpuGrdiyGBQ7uFLv6iYwFyOaTaLKPqB1NSPbhDElI
QumtBd9Gd1VyzXHl53JkubrurXe7QwGob+WEThUzrmbZnrCv3nZ0+tABlwWsVq8o
ctbBPEy5RqwIbmNJJOQtk5ECgYEAmUEVAkOkQH3l04siIa/OL3MXjr8/tE+nafIl
T4sK5JajfinTknZL4DDHqeCsx+BnyUd2WiJbUM2PEXQBX79eVENtYiBVDFH9Vbhv
oEWfAkumJjEam3EeqQn108WDrAQhmyGA/qlsNDi9Q6OMeZE70rS/beXHf+eEPC3k
hUc1MIECgYEA25cS6Y1ePULSytD3tRZSHa4IdtCf/qdyQQpyVm7/2EO+pq5iXmOc
ver95l7VKGS+goIFD4kzfX4KYpqqQB6puVGOF33fl75cXq8V4Hl4vptewhfkxc5l
/OkvLiVFzjQSxmrQ6BuJe9hVGgqwWaHAon6owhYXzk3gYx9sBV4dit0=
-----END RSA PRIVATE KEY-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1025  100  1025    0     0   1025      0  0:00:01 --:--:--  0:00:01  100k
-----BEGIN CERTIFICATE-----
MIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTE3MTEyNTE4MTkzMFoXDTI3MTEyMzE4MTkzMFowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL2O
x5luhBr98NfvhIyfAAM7A2gPZ2/KRddkaY4QfRn/Z8jWBG5m8nHmDh98EWSmoIF9
40xtT03GiR4VWtohtdQMkaUsklLijoOpJmHiuiu2CBg59qQsJ5Mkg59ep3MaHzP3
WiB1vAD55nQiW/D078bUOyMuuXsyLUJk2q3L/7ph7DZycac/LTifEcbGCWSS1ie4
2hInRGiY4hpRFmU1tR4hoR4ThJfFI2Ro+AxQcxXPNiaIAkDmYSXL/IqqkRxg1ORg
hxfASOZkOwrvY9OL3a6RUSn02p5hN0TxAVDYbdbF6CG25EdVEQfxDbsPXqDUR6QE
32f8xeBOmXuX2CIE8rECAwEAAaMjMCEwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBALgayIUpSqwc3EX5Nu2BNPyMOIPR
nZuHCR2sbOWlGuNa+L3dL/y4c/osvysgZZXUzi0zbupZCEC9hauJ1yG7SHFhdvDi
V83WTXDbS07t4nDeCB4UDYkZE8WIDrOaSGybwxLSFCHF4hG15SGHYgYTZpwPdt4j
lAaJxG6CjKa8ixKU2vrEddFuWSHsKOYbspxfu9liwJQXXz7nZnLm8qs0o7NXGnyr
VFlMsytvICywpA/rjTY7nNHVf1Y3A5phVYd2sZheXD+cVX7EF/fsoqwNCj4n5HqJ
0XWrV8OxrSF599pq8Y6AWGBgSzhl2CH/PhHlXqyPOl11FKc0rTvBNP0aBfM=
-----END CERTIFICATE-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1679  100  1679    0     0   1679      0  0:00:01 --:--:--  0:00:01  234k
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAvY7HmW6EGv3w1++EjJ8AAzsDaA9nb8pF12RpjhB9Gf9nyNYE
bmbyceYOH3wRZKaggX3jTG1PTcaJHhVa2iG11AyRpSySUuKOg6kmYeK6K7YIGDn2
pCwnkySDn16ncxofM/daIHW8APnmdCJb8PTvxtQ7Iy65ezItQmTarcv/umHsNnJx
pz8tOJ8RxsYJZJLWJ7jaEidEaJjiGlEWZTW1HiGhHhOEl8UjZGj4DFBzFc82JogC
QOZhJcv8iqqRHGDU5GCHF8BI5mQ7Cu9j04vdrpFRKfTanmE3RPEBUNht1sXoIbbk
R1URB/ENuw9eoNRHpATfZ/zF4E6Ze5fYIgTysQIDAQABAoIBAQCcyAQkzjirjo9u
OmfaWwQKp/u4TM1Ts+wyK5X2b4LZRtrTTNhbbABo3kueFOLZ6ctlbvbZ1qs5+9Sb
MjRzqGAYo9Cu2f+oWERhdz0MxxcOYj+tsftibDbLR+8pmINM+zUFqTc9E8PM1uSZ
93UuBZhwTkcfYVz+HVE70ejgyJKZ10jsT214dHXidS3aL3h3PTJNsjPu7cpSscHC
mJC3747LiNreGz7MLzMnYc0OBpKAlMGDLq0ZCad5Qu8IgFluJLPZqcKdWOlQBDMk
vLOx2ziRpTrKBOAQP7q6dwbTsAIjO/R+p8Zd6xTGDsn5kr/6J3SbUZpRgRi6hV4j
SNrprC4xAoGBAOuRrndFNEJ9d9Vwqmv43MSmPK1eQp13iItaYzlf+tVwDULX7w69
6KWB72HFWBno37PVMGd7IK4RGG0SixbkIgxDl9iZ13adj+EhnbwjJ8t4XsrVsVXY
S6y14XLgseNRdCH8s3wuK4bAUp4AVx1VuPZNrBVCIwlktEK1EDyl6yylAoGBAM3/
g0DCFghmcJzP7me0cbBfS1LX8GCuzWobW0VBVzKd6j+SOL7K9fl0S/QByF3m/7x3
HUbK+SbzYoBdFnip4B6klTh1dTacjGEpRXMCFiSKpGYrnIKcMZvo1n1CAmv/AJI+
IywSo1CIhZbuhVX4Zhck1Pao67cD/0pprxpyFBQdAoGAJlWxx5UKeRuD6ccI5MM9
4f6Pd0MUoh12qf7OkOKuphczAUK0k3iSKITpmMaNnMHJzbEOzHg8DDeZ+v3+hn4f
kFFbn5MBfjsL0/4tdhVef30aJ/X4gtjND9EVXCV2rWBu/qzp5HZ6/NXFXBMAGjT+
AY8IRvlZTEkIEKb4enYhvEUCgYBlMkeFUZNHTyO1NUSV94P+M+Ip8WDp7hWFMak1
rXKrr/36etebitMpV6NUXGfVEXzCVfWw6FdbdicGjjRS1IkAJLQVh0YsrN0dggvX
Ashj7SPz+Sqq0xZf1BNCd2H/McdTEGlw5jsKffa6fbRM1nk4l2RzQ+p//tuz75UQ
tcf9VQKBgQCg66/ZRzoU787F0T/hyiO6/tozZiKcMySxQmcHR0qq85YCQ2Q26ecs
KtaAl/k4KG6CeGrayUk9DMbL2vPdugOS4Vn1GU4LOxFblvoICZw94/LkfHwtoIYe
ioXczn60eUJ6NkcS7ATyW2/yVKHxByXzhqChxdgSvMdOAjb4KkDALA==
-----END RSA PRIVATE KEY-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1025  100  1025    0     0   1025      0  0:00:01 --:--:--  0:00:01  500k
-----BEGIN CERTIFICATE-----
MIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTE3MTEyNTE4MTkzMVoXDTI3MTEyMzE4MTkzMVowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANI2
0AI1/RPYvoi6qIu6TTegcmFf1o4Xibz4DszeI9oi63It69uEx70f9JTxHh7R4TKv
LtkPfR2tdjYtTQglKV1KtWzUPrUpgPAlnTnpKHsKvqhFXg6wQ5/cfDWGyFVnCcc+
4lcf/6ampyhGs/AgCzlFjGkxYDNbPPb+co6XHNOBf7yJV4RsDQZyzAYu57mujbA4
4NmkLrX9WEi83xFCfixVvPO9OcMdehvqtVsh6kMp9RpMwy8yff86UEtJquTcmMuf
T7eMkaY6SrPmk2k1I7iNYYuwHGUT6XYl2CBPMCOBPB6nOjYHAzaJFfWbQxGQrGS0
rerRR8/dkeTIjXVF6LECAwEAAaMjMCEwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAKh9KRu8/gb45NwwfJhSbBKmiBR/
d7nCmgvjS2w/uHPYenNh2X7Jc04Rcv1ME5KiriYpxQB21r2EK/Aw0W0bodmiManB
tIJ8N8GOohH3xsDBrHAZhkoTbRVoBMGPaoB+FAyWbTElZldNlfVK/UxoaoB9IZSa
kiekz+jA4MJyYN40EWKnZUfSg9Sh3SlXrlGtjq083flC6IyckhBU+GoLosplVUM9
keSs6bVqampuhx5oCOAg7tcNMPImLH0tg4sMT1JadiNYdnqXtf0+tW2ZBA4qKLFE
1PGi2soAmNCHZ9yIJV1i6bC8LJQIw/3D5hWnYpZ0cMNdeGSqCwHuyHP7R4I=
-----END CERTIFICATE-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1675  100  1675    0     0   1675      0  0:00:01 --:--:--  0:00:01  545k
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA0jbQAjX9E9i+iLqoi7pNN6ByYV/WjheJvPgOzN4j2iLrci3r
24THvR/0lPEeHtHhMq8u2Q99Ha12Ni1NCCUpXUq1bNQ+tSmA8CWdOekoewq+qEVe
DrBDn9x8NYbIVWcJxz7iVx//pqanKEaz8CALOUWMaTFgM1s89v5yjpcc04F/vIlX
hGwNBnLMBi7nua6NsDjg2aQutf1YSLzfEUJ+LFW88705wx16G+q1WyHqQyn1GkzD
LzJ9/zpQS0mq5NyYy59Pt4yRpjpKs+aTaTUjuI1hi7AcZRPpdiXYIE8wI4E8Hqc6
NgcDNokV9ZtDEZCsZLSt6tFHz92R5MiNdUXosQIDAQABAoIBAQCCwM1JUsOD4MFC
0rebkdoNy6LZpn8SCtoFzFWIYHP0btv8PUwDUufaX1IGgqO1j95/ZC/Dm7q7BvYM
ZMkRU+wStzF569rFXHdy6DKPntMrFVfZ2qvSfxjRCpfHw3pAhMK0SA0JHEIsKmS2
zKOHrKbOjM/1blzkFNAq4cHnp16qvzT0S243OrSQek6AP+Y7ACAwDoxM8DhFIpRi
C1J0sNt34RS8pzlahVo2IPLiT1K4eb+cu+cpnwlsAd7loS6FvTZ3D8ddTtvKHUYd
kBdBXnmh1f1jr8qmlRPDFv9o/A7oA9yMpxOwYxVO1SP2GoGRIXw/SoEs+iqrOj0J
EyzA7i0xAoGBAPSAJHN2L/DSkwKaQKZOMqTwyTrlNzQmsskmgdlnJNxxh4a3yKia
nv+DtRiTS7eJb1lYaQd2pczOvCaGQE+0hMnuxkpbahNkFF0kZLWQhfPXBjKANvdf
FKQ09zcdAxHLRK9qS8kH3r0sBP2ng/F9ng2bBLkW0mfn4W4vOqzcmB3LAoGBANwZ
2fXZYEjYq3fDz6suzc5e1O5BA5EYuynYWphXX8WyfL4/zIm2Lh1gcqXt1Gh757lI
fJci5QvA9qvyGgj6KQ7eA6jpJUZPQRTbGvGoN3H+0hJsINuIK/hCmL8BqvkdttTL
2NCdhdxfm51yZj+YGTirWZcVL8kwU04Q8LBQkMPzAoGAR8DHcb8QiKszi4L7UBMI
19D2LTuRfNOSimermcKkVbFXpZiP4bMm0DO11NV9VgqWS/EfOZp/Q+Cz5NZnTzzZ
720MckvrkG9Bbqg8ifA0739+CRoJh4U2yI3msxznhrRRnu0h+og9vOSLPmmH1joY
RUPmGe9xph6t0Kk5nftLBQsCgYAAyFK/DgPKnoKVjWuqlyMCLdIh+rcxh936extv
WTgs8sKAxqdKggxN7G7b6imMKYycWwaeXvBjh08ZM2YmTAVjWSS4MAQC6Ps5QMIq
w8dWr2o9yIahYS2O12XFoBLkideBkCv4Kec07l4WRDba2QaV7f4GjHu8AmD5cRP5
/H9XQwKBgHsePJuoNWzxsrZTkp7p2RvG0jeICDt/TvFNbt/7Y9w9A6gSD7U2a+2j
fwj50b/5SOSsQBVXAh2LaAguhi49w2cnFkONkYy40KTdRd/8sdabZzB0t9xhZRX7
PI/wiehtOcm1q2kC5pscCcwAXycIv9goXLM6immYKu2ESgq00o+v
-----END RSA PRIVATE KEY-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1050  100  1050    0     0   1050      0  0:00:01 --:--:--  0:00:01 1025k
-----BEGIN CERTIFICATE-----
MIIC2zCCAcOgAwIBAgIIA08fMC10mgEwDQYJKoZIhvcNAQELBQAwFTETMBEGA1UE
AxMKa3ViZXJuZXRlczAeFw0xNzExMjUxODE5MzFaFw0xODExMjUxODE5MzFaMB0x
GzAZBgNVBAMTEmZyb250LXByb3h5LWNsaWVudDCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBAOhvlAWxu1POJ59bJpNOdt0Ry4mgs+6j0PeLulbr73CPMkNN
kefCjPtQ1AHh0qdW+TCI4Ul0UdqTD+yvNISwGgbQZYk6ht9rP9qVW5KvWCOvvnYF
PpFWLaZM8G2R5Sl7WuRhJLRI92UnTiVhxqevHx/aE2Sy8aPovOk+jn3ZPEzvXOWK
mRpV/8Fm1ElctQzdUH3xmVfBg8h94wIkWoF18CGOakGaaxTxJ+o8cH+K7Q+TUE9G
QCkbMd/8FTecuRJuazUNtyxll80XvS8/J2YpZqh6Q1Hq8/FsOktv6n9UXOh7XUOp
AXi+M76NssLbk/U31wu3qKjsIYYMWXw3r+Fbv80CAwEAAaMnMCUwDgYDVR0PAQH/
BAQDAgWgMBMGA1UdJQQMMAoGCCsGAQUFBwMCMA0GCSqGSIb3DQEBCwUAA4IBAQBi
s/FTCOirYq6NbkDZsevRCMRG/DoTe4xKGKnyUYgrgYuHGjyg6Iff6WlQhFD66SCz
ul9ADMk3pRhOVi+OUmSNi0DpB9ds/zA8K+qTY4jGYkZratCXWdMUjIKXFiKQM5FK
wMHRvQojbHE8uZ9/f4/63Jckqjto+zGL/mRAWDTQfazb43DJt28hFQV25HxJthBV
6DqwAWNrTfDwfSYpOC4VgGvEIAXRW2AXZNTKD/nmLsZqq9qsntJ46/txV0H6XpS4
f012o+6VUuy3IyRwki0Ej/0UoR+HpAJLc5X7azp8Nnaa9PPG5Lmn3FU/sMyQjpS0
QeuAYV3NaJnJQbFZq17q
-----END CERTIFICATE-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1675  100  1675    0     0   1675      0  0:00:01 --:--:--  0:00:01  817k
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA6G+UBbG7U84nn1smk0523RHLiaCz7qPQ94u6VuvvcI8yQ02R
58KM+1DUAeHSp1b5MIjhSXRR2pMP7K80hLAaBtBliTqG32s/2pVbkq9YI6++dgU+
kVYtpkzwbZHlKXta5GEktEj3ZSdOJWHGp68fH9oTZLLxo+i86T6Ofdk8TO9c5YqZ
GlX/wWbUSVy1DN1QffGZV8GDyH3jAiRagXXwIY5qQZprFPEn6jxwf4rtD5NQT0ZA
KRsx3/wVN5y5Em5rNQ23LGWXzRe9Lz8nZilmqHpDUerz8Ww6S2/qf1Rc6HtdQ6kB
eL4zvo2ywtuT9TfXC7eoqOwhhgxZfDev4Vu/zQIDAQABAoIBABufBswWbthqqzK4
tkDTFjT3JGTagAdf2DlSlXY222FY4vQ0h54u9Gn5FPiUSWAxabNm0L7V39RFnaLW
vzr1YmJk6WwBVH3AVPo4vYUU6Th0bgHAZcsHnq0gf4HF0KG5UvfIw3/ZG+0dWp/+
mojauqwiUiOpUCXWRnzBPlWeNxg9L5cjZ+T5kL4fAPQJK1RlXcJDNTtJyUAd84m8
nkRuzMkTZvI5X+R3N1LHV5teUpXORmLXGDF8/rDql/VayuG0AMEkS8xsslpB6gTP
QZP4bkwAT/BtqDK3TONvbhXfkHlOx/7Mi89gR+7IimNnHpJume0k+bP2lvMLLaxE
PAXc7lECgYEA/4BwGwNiLzeMAF27u8J0W49n6M7Fzmg2QMLz0vIw+0NvSVoIdA9G
5C5KBX0085L7J4gXeSFCLebLx7it88xKMjNFQ7pEoWH1LEN9PmLIk8EPLnVwgsFl
lCH1LSN74hAYkDEenNzgxb/7Md1988GMN3JFcYllkQpdvCBjQHiXihcCgYEA6OOf
2XVxEA5EkldgTepFpewCA67AZDANInRN2bPrHNfN3tiSBt8JuqlzlrRNXSXJU//v
bwxpHdAfk1DSoAwxjaUu+lM2spthibsypTFU6mzrrNUxtfEPvAetcVFndSmq8XEP
XpAo7KUEEUSCxVxkWN6C86QNXx9xDnQXe5F+x7sCgYAzYySo+jQ8JOLei/ufFswK
Qjx24Wd9TQ5kVhFuAtIBJx3x1c7PcOckK38tvXHx70TePYvjC+JgkwO3RpiN6S+f
CxrKNWNqReZmoahF7N5EgVbJVEttsPH6Kh5CltlqrSwTjB2JFQ+MFkGyVMbpSKsn
167J2AV6d2hE1wq/Q8HOvwKBgC+4fMe9deZ5Vfm6p8QvbggavSruAwREiFx4An/K
izHE0q2IKRv5VfAGyuWlmnisxn1M404Y9+0g+WdhV1duSDUjFOY8aOfeSPVFA5Kv
8ht7KCupH9+NPtEjrZmbuZxbZmjE5wZUed5LzfKP3E/+p51CMwkjo43LTopMpYcv
gBADAoGBAIeVPw6wwBCKEiP/qmHUDZV5khoom9jwLorLhB1jFm9SPRXIihZvPDq9
lTQwA63NxduUjdR+xSI06nb9askm4LSRUFiR/LUWkpb65e12n8hiP1zUtbCP6XkA
knVlLdSaEgxztok5xgomeSL3YtkdAybRvCRmVFEC5WLZl279EV+u
-----END RSA PRIVATE KEY-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1679  100  1679    0     0   1679      0  0:00:01 --:--:--  0:00:01 73000
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA2jZYvsu/71vcj6Tkt1VRW0apsrxPZBK92wySWdcCyGJOfpTi
EX/mxe9M3gFGvNtmEiJKKYQXCY8X69f1HdzLx1ORaA2l8fr8bBVujQIOpWoap2Rs
zIkqdfTE2t3wQRyYluGBkMMjBLvp6tNVpXzSdeN07eNrdoGJPLFKrTQMOIuE3jy3
yppEM7NQjt4Po6r8+szaEVMRGK0dEhd2X1zMA6ztZpkPGms0Sx2Zj8RSMHQ448l8
JyQa/5nELAphlf/s+o9xzFRdKElI5qwdq1hJaapsCPs5sMqy331qobRKNLC5BImh
ID3YUVTHVxkgFPoQpg04jrve/BBP56AgEKYKcwIDAQABAoIBAHaCt5X/wGD16gbV
jiYHnyTMbtOUtCVsjFIQNNe4D1nXEjxrDZBlf2cWfQ/Bb7KYEYpdtxtvErftUBeV
dRlf1giOWVVCvFJ8dmUZCq5oZJ3vz0QEJZa8FCR16AfaK2ahajnfhWeioz0p9XPE
jSpQbVj7TID6V13IIkmlHCKPNufp1V7Yyi87LgQTLJzzgOv7++F1lZPobHJFN9rV
uLFkWcETYMBv02YL4Fk2krl4gVlIISdq6vE3FqdfRANXy+AKgd7MzUFFzUXoHe94
BCmZYU5Y5Zhp662t3YBSFhm0OeN+0d3XxoV123G2IptH9dYid1CpFUg6hi/uUATE
hCi2VaECgYEA8FOTKE9linHDiwnjkH2YRARVw9+cS9N5fri51ttk7pOkvJvN6qcb
HmuSU2EjbAjliFnkcwbFejQi2stMOxm8dIGL9emb3lXtoGcgJ6MWQlWO/mb7JvBU
s4M1JwRmm3hi7TntHb4wnB2IREcy2CMQUAULxZFGm6KZVBVlOCfXPYMCgYEA6HGP
Q8YODq8Top7zEzOH5TFnXSEqBtMEYLeA04QTZ1tEOh1opK9jobF90yOUTJkHYe+W
feT+9myrPa7h9bjQV5KFMuzVNCFmZcqdSk8FMd5FN3bkckyCB3UzJr0G4hzxg79S
ZITETf4IBRHD9xKZue6PCFEksMdpa0vJEmhl3FECgYAC08jgaYa1ST9WrGgb9A5N
houHwWKyVauRPRUanPwj/mKqwV57yzl7cyKGjpb/F1+z6fLRE5xRIrniPxAHhfb8
m+WMoBHwXKxt/aiYLsf9Qcpr3cRIof0tG3cNnXSjkZgAWHUcq3cg7wwCWaAUgN0Z
Q+XujhX5kMEjBxlovrc+hwKBgQCHM8aJPWJf2CYrqRjqSkb5YUsmOGwhnAwn1ZTx
cxl2PPfqdcCwSfRB0OPHNCLGixxoMtvuToye+YaBUn/MVbuZtHgVrvw0XlkZTElB
N8VlxEuuf1FRyn63p8wg9EJUs1EnuNlL7o3gb6pqiHLb+82/AKvhwpD84oDHEEXc
lyzogQKBgQCRDO+HbNGEhjGGApSUTtxP/tpog3+U8pinyecPMseUM9AP9eL/9C14
zRp8nt0L1UP+3ufB01dfe12RYzo4tjMhJATGqKMORD/owarspL9TEobPqpscTrPg
nNxFi8zBnC52qYO5DPhk6kxKt8wFA+JFbRL/kBHMqXItB8kwOWj1Uw==
-----END RSA PRIVATE KEY-----
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   451  100   451    0     0    451      0  0:00:01 --:--:--  0:00:01  146k
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2jZYvsu/71vcj6Tkt1VR
W0apsrxPZBK92wySWdcCyGJOfpTiEX/mxe9M3gFGvNtmEiJKKYQXCY8X69f1HdzL
x1ORaA2l8fr8bBVujQIOpWoap2RszIkqdfTE2t3wQRyYluGBkMMjBLvp6tNVpXzS
deN07eNrdoGJPLFKrTQMOIuE3jy3yppEM7NQjt4Po6r8+szaEVMRGK0dEhd2X1zM
A6ztZpkPGms0Sx2Zj8RSMHQ448l8JyQa/5nELAphlf/s+o9xzFRdKElI5qwdq1hJ
aapsCPs5sMqy331qobRKNLC5BImhID3YUVTHVxkgFPoQpg04jrve/BBP56AgEKYK
cwIDAQAB
-----END PUBLIC KEY-----
```

Modify kubelet.service
```
[vagrant@localhost ~]$ sudo sed -i 's/\(ExecStart=.*ARGS\)/\1 --fail-swap-on=false/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
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
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false
```

Copy master configurations
```
[vagrant@localhost ~]$ scp -i .ssh/vagrant -r vagrant@10.64.33.82:/home/vagrant/etc0x2Fkubernetes .
kubelet.conf                                                                                                  100% 5530     2.0MB/s   00:00    
controller-manager.conf                                                                                       100% 5486   558.7KB/s   00:00    
sa.key                                                                                                        100% 1679   157.4KB/s   00:00    
apiserver.crt                                                                                                 100% 1387     1.5MB/s   00:00    
front-proxy-ca.crt                                                                                            100% 1025     1.1MB/s   00:00    
front-proxy-client.crt                                                                                        100% 1050   943.8KB/s   00:00    
apiserver.key                                                                                                 100% 1679     2.6MB/s   00:00    
ca.key                                                                                                        100% 1679     2.3MB/s   00:00    
ca.crt                                                                                                        100% 1025     1.7MB/s   00:00    
apiserver-kubelet-client.crt                                                                                  100% 1099     2.0MB/s   00:00    
apiserver-kubelet-client.key                                                                                  100% 1679     2.8MB/s   00:00    
front-proxy-client.key                                                                                        100% 1675     3.1MB/s   00:00    
front-proxy-ca.key                                                                                            100% 1675     3.0MB/s   00:00    
sa.pub                                                                                                        100%  451   774.4KB/s   00:00    
kube-scheduler.yaml                                                                                           100%  991   937.1KB/s   00:00    
kube-apiserver.yaml                                                                                           100% 2726     4.7MB/s   00:00    
kube-controller-manager.yaml                                                                                  100% 2232     2.6MB/s   00:00    
scheduler.conf                                                                                                100% 5430   178.6KB/s   00:00    
admin.conf                                                                                                    100% 5446     2.9MB/s   00:00    
```

```
[vagrant@localhost ~]$ export KUBECONFIG=/home/vagrant/etc0x2Fkubernetes/admin.conf 
```

```
[vagrant@localhost ~]$ kubectl get nodes
NAME                   STATUS     ROLES     AGE       VERSION
kubedev-10-64-33-195   Ready      master    18d       v1.8.4
kubedev-10-64-33-199   Ready      master    18d       v1.8.4
kubedev-10-64-33-82    Ready      master    18d       v1.8.4
```

bootstrap token
```
[vagrant@localhost ~]$ sudo kubeadm token create --kubeconfig=etc0x2Fkubernetes/admin.conf
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --ttl 0)
1e5b81.d491fedcb4786c5d
```

join
```
[vagrant@localhost ~]$ sudo kubeadm join --token=1e5b81.d491fedcb4786c5d 10.64.33.82:443 --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[preflight] Skipping pre-flight checks
[validation] WARNING: using token-based discovery without DiscoveryTokenCACertHashes can be unsafe (see https://kubernetes.io/docs/admin/kubeadm/#kubeadm-join).
[validation] WARNING: Pass --discovery-token-unsafe-skip-ca-verification to disable this warning. This warning will become an error in Kubernetes 1.9.
[discovery] Trying to connect to API Server "10.64.33.82:443"
[discovery] Created cluster-info discovery client, requesting info from "https://10.64.33.82:443"
[discovery] Cluster info signature and contents are valid and no TLS pinning was specified, will use API Server "10.64.33.82:443"
[discovery] Successfully established connection with API Server "10.64.33.82:443"
[bootstrap] Detected server version: v1.8.4
[bootstrap] The server supports the Certificates API (certificates.k8s.io/v1beta1)

Node join complete:
* Certificate signing request sent to master and response
  received.
* Kubelet informed of new secure connection details.

Run 'kubectl get nodes' on the master to see this machine join.
```

```
[vagrant@localhost ~]$ sudo systemctl start kubelet.service
[vagrant@localhost ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Sat 2017-12-16 04:53:04 UTC; 2s ago
     Docs: http://kubernetes.io/docs/
 Main PID: 3945 (kubelet)
    Tasks: 7 (limit: 4915)
   CGroup: /system.slice/kubelet.service
           └─3945 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --pod

Dec 16 04:53:04 rookdev-10-64-33-201 systemd[1]: Started kubelet: The Kubernetes Node Agent.
Dec 16 04:53:04 rookdev-10-64-33-201 kubelet[3945]: I1216 04:53:04.626010    3945 feature_gate.go:156] feature gates: map[]
Dec 16 04:53:04 rookdev-10-64-33-201 kubelet[3945]: I1216 04:53:04.626634    3945 controller.go:114] kubelet config controller: starting control
Dec 16 04:53:04 rookdev-10-64-33-201 kubelet[3945]: I1216 04:53:04.626828    3945 controller.go:118] kubelet config controller: validating combi
Dec 16 04:53:04 rookdev-10-64-33-201 kubelet[3945]: I1216 04:53:04.899902    3945 client.go:75] Connecting to docker on unix:///var/run/docker.s
Dec 16 04:53:04 rookdev-10-64-33-201 kubelet[3945]: I1216 04:53:04.900152    3945 client.go:95] Start docker client with request timeout=2m0s
Dec 16 04:53:04 rookdev-10-64-33-201 kubelet[3945]: W1216 04:53:04.903342    3945 cni.go:196] Unable to update cni config: No networks found in 
Dec 16 04:53:04 rookdev-10-64-33-201 kubelet[3945]: I1216 04:53:04.908230    3945 feature_gate.go:156] feature gates: map[]
Dec 16 04:53:04 rookdev-10-64-33-201 kubelet[3945]: W1216 04:53:04.908656    3945 server.go:289] --cloud-provider=auto-detect is deprecated. The
```

but
```
[vagrant@localhost ~]$ ls /etc/kubernetes/manifests/
```

