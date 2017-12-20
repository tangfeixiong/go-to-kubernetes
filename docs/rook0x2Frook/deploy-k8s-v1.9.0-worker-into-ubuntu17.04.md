# DevOps


ubuntu@ubuntu-zesty:~$ sudo cat /etc/network/interfaces
```

auto enp0s8
iface enp0s8 inet static
  address 172.17.4.61
  netmask 255.255.255.0

auto enp0s9
iface enp0s9 inet dhcp
```

```
ubuntu@ubuntu-zesty:~$ sudo systemctl restart networking.service
```

```
ubuntu@ubuntu-zesty:~$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 02:f2:d5:cb:ed:dc brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global enp0s3
       valid_lft forever preferred_lft forever
    inet6 fe80::f2:d5ff:fecb:eddc/64 scope link 
       valid_lft forever preferred_lft forever
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:4a:74:4e brd ff:ff:ff:ff:ff:ff
    inet 172.17.4.61/24 brd 172.17.4.255 scope global enp0s8
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe4a:744e/64 scope link 
       valid_lft forever preferred_lft forever
4: enp0s9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:26:cd:70 brd ff:ff:ff:ff:ff:ff
    inet 172.28.128.7/24 brd 172.28.128.255 scope global enp0s9
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe26:cd70/64 scope link 
       valid_lft forever preferred_lft forever
```

```
ubuntu@ubuntu-zesty:~$ hostname -I
10.0.2.15 172.17.4.61 172.28.128.7 
ubuntu@ubuntu-zesty:~$ addr=$(hostname -I | awk '{print $2}'); echo $addr; name="rookdev-${addr//\./-}; echo $name
```

```
ubuntu@ubuntu-zesty:~$ addr=$(hostname -I | awk '{print $2}'); echo $addr; name="rookdev-${addr//\./-}"; echo $name
172.17.4.61
rookdev-172-17-4-61
```

```
ubuntu@ubuntu-zesty:~$ sudo hostnamectl set-hostname $name
ubuntu@ubuntu-zesty:~$ hostname
rookdev-172-17-4-61
ubuntu@ubuntu-zesty:~$ cat /etc/hostname
rookdev-172-17-4-61
```

```
ubuntu@ubuntu-zesty:~$ echo "$addr    $name" | sudo tee -a /etc/hosts
172.17.4.61    rookdev-172-17-4-61
``` 

```
ubuntu@ubuntu-zesty:~$ sudo ufw status
Status: inactive
```

```
ubuntu@ubuntu-zesty:~$ uname -r
4.10.0-42-generic
```

__issue__
```
ubuntu@ubuntu-zesty:~$ sudo sysctl net.bridge
sysctl: cannot stat /proc/sys/net/bridge: No such file or directory
```

refer to https://askubuntu.com/questions/645638/directory-proc-sys-net-bridge-missing

```
ubuntu@ubuntu-zesty:~$ sudo modprobe br_netfilter
```

```
ubuntu@ubuntu-zesty:~$ sudo sysctl net.bridge
net.bridge.bridge-nf-call-arptables = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-filter-pppoe-tagged = 0
net.bridge.bridge-nf-filter-vlan-tagged = 0
net.bridge.bridge-nf-pass-vlan-input-dev = 0
```

```
ubuntu@ubuntu-zesty:~$ sudo apt-get update 
Hit:1 http://security.ubuntu.com/ubuntu zesty-security InRelease                                         
Hit:2 http://archive.ubuntu.com/ubuntu zesty InRelease                                                   
Hit:3 http://archive.ubuntu.com/ubuntu zesty-updates InRelease
Hit:4 http://archive.ubuntu.com/ubuntu zesty-backports InRelease
Reading package lists... Done
```

```
ubuntu@ubuntu-zesty:~$ sudo apt install apt-transport-https software-properties-common 
Reading package lists... Done
Building dependency tree       
Reading state information... Done
software-properties-common is already the newest version (0.96.24.13).
apt-transport-https is already the newest version (1.4.6~17.04.1).
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
```

```
ubuntu@ubuntu-zesty:~$ sudo add-apt-repository "deb https://apt.kubernetes.io/ kubernetes-xenial main"
```

```
ubuntu@ubuntu-zesty:~$ curl -jkSL https://apt.kubernetes.io/doc/apt-key.gpg | sudo apt-key add -
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0    149      0  0:00:01  0:00:01 --:--:--   148
100   663  100   663    0     0    428      0  0:00:01  0:00:01 --:--:--   428
OK
```

```
ubuntu@ubuntu-zesty:~$ sudo apt-get update && sudo apt-cache show kubeadm=1.9.0-00
Hit:1 http://security.ubuntu.com/ubuntu zesty-security InRelease
Hit:2 http://archive.ubuntu.com/ubuntu zesty InRelease
Hit:3 http://archive.ubuntu.com/ubuntu zesty-updates InRelease
Hit:5 http://archive.ubuntu.com/ubuntu zesty-backports InRelease
Hit:4 https://packages.cloud.google.com/apt kubernetes-xenial InRelease
Reading package lists... Done
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

```
ubuntu@ubuntu-zesty:~$ sudo apt install -y --no-install-recommends docker.io kubeadm
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  containerd kubectl kubelet kubernetes-cni runc socat
Suggested packages:
  aufs-tools debootstrap docker-doc rinse zfs-fuse | zfsutils
Recommended packages:
  cgroupfs-mount | cgroup-lite ubuntu-fan
The following NEW packages will be installed:
  containerd docker.io kubeadm kubectl kubelet kubernetes-cni runc socat
0 upgraded, 8 newly installed, 0 to remove and 0 not upgraded.
Need to get 74.5 MB of archives.
After this operation, 496 MB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu zesty-updates/universe amd64 runc amd64 1.0.0~rc2+docker1.13.1-0ubuntu1~17.04.1 [1,501 kB]
Get:2 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubernetes-cni amd64 0.6.0-00 [5,910 kB]
Get:3 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubelet amd64 1.9.0-00 [20.5 MB]                                      
Get:4 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubectl amd64 1.9.0-00 [10.5 MB]                                      
Get:6 http://archive.ubuntu.com/ubuntu zesty/universe amd64 containerd amd64 0.2.5-0ubuntu1 [3,820 kB]                                         
Get:5 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.9.0-00 [20.0 MB]                                      
Get:7 http://archive.ubuntu.com/ubuntu zesty-updates/universe amd64 docker.io amd64 1.13.1-0ubuntu1~17.04.1 [11.9 MB]                          
Get:8 http://archive.ubuntu.com/ubuntu zesty/universe amd64 socat amd64 1.7.3.1-2 [337 kB]                                                     
Fetched 74.5 MB in 1min 28s (838 kB/s)                                                                                                         
Selecting previously unselected package runc.
(Reading database ... 56519 files and directories currently installed.)
Preparing to unpack .../0-runc_1.0.0~rc2+docker1.13.1-0ubuntu1~17.04.1_amd64.deb ...
Unpacking runc (1.0.0~rc2+docker1.13.1-0ubuntu1~17.04.1) ...
Selecting previously unselected package containerd.
Preparing to unpack .../1-containerd_0.2.5-0ubuntu1_amd64.deb ...
Unpacking containerd (0.2.5-0ubuntu1) ...
Selecting previously unselected package docker.io.
Preparing to unpack .../2-docker.io_1.13.1-0ubuntu1~17.04.1_amd64.deb ...
Unpacking docker.io (1.13.1-0ubuntu1~17.04.1) ...
Selecting previously unselected package kubernetes-cni.
Preparing to unpack .../3-kubernetes-cni_0.6.0-00_amd64.deb ...
Unpacking kubernetes-cni (0.6.0-00) ...
Selecting previously unselected package socat.
Preparing to unpack .../4-socat_1.7.3.1-2_amd64.deb ...
Unpacking socat (1.7.3.1-2) ...
Selecting previously unselected package kubelet.
Preparing to unpack .../5-kubelet_1.9.0-00_amd64.deb ...
Unpacking kubelet (1.9.0-00) ...
Selecting previously unselected package kubectl.
Preparing to unpack .../6-kubectl_1.9.0-00_amd64.deb ...
Unpacking kubectl (1.9.0-00) ...
Selecting previously unselected package kubeadm.
Preparing to unpack .../7-kubeadm_1.9.0-00_amd64.deb ...
Unpacking kubeadm (1.9.0-00) ...
Setting up kubernetes-cni (0.6.0-00) ...
Setting up runc (1.0.0~rc2+docker1.13.1-0ubuntu1~17.04.1) ...
Processing triggers for ureadahead (0.100.0-19) ...
Setting up socat (1.7.3.1-2) ...
Setting up containerd (0.2.5-0ubuntu1) ...
Setting up kubelet (1.9.0-00) ...
Created symlink /etc/systemd/system/multi-user.target.wants/kubelet.service → /lib/systemd/system/kubelet.service.
Processing triggers for systemd (232-21ubuntu7.1) ...
Setting up kubectl (1.9.0-00) ...
Processing triggers for man-db (2.7.6.1-2) ...
Setting up kubeadm (1.9.0-00) ...
Setting up docker.io (1.13.1-0ubuntu1~17.04.1) ...
Adding group `docker' (GID 117) ...
Done.
Created symlink /etc/systemd/system/sockets.target.wants/docker.socket → /lib/systemd/system/docker.socket.
Processing triggers for ureadahead (0.100.0-19) ...
Processing triggers for systemd (232-21ubuntu7.1) ...
```

```
ubuntu@ubuntu-zesty:~$ sudo usermod -aG docker ubuntu
```

```
ubuntu@ubuntu-zesty:~$ sudo systemctl is-active docker.service kubelet.service
active
activating
```

```
ubuntu@ubuntu-zesty:~$ mkdir .kube; scp vagrant@172.17.4.59:/home/vagrant/.kube/config .kube
The authenticity of host '172.17.4.59 (172.17.4.59)' can't be established.
ECDSA key fingerprint is SHA256:VV533NeJHFykYUa3CBNDm0zDI4HSnEaMKcQSKwtvLGw.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.17.4.59' (ECDSA) to the list of known hosts.
vagrant@172.17.4.59's password: 
config                                                                                                        100% 5451     4.8MB/s   00:00    
```

```
ubuntu@ubuntu-zesty:~$ kubectl version
Client Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T21:07:38Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T20:55:30Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
```

```
ubuntu@ubuntu-zesty:~$ exit
logout
Connection to 127.0.0.1 closed.
```

```
fanhonglingdeMacBook-Pro:zesty fanhongling$ ssh -i ~/.ssh/vagrant ubuntu@172.17.4.61
The authenticity of host '172.17.4.61 (172.17.4.61)' can't be established.
RSA key fingerprint is de:4e:67:93:85:1e:c0:05:36:5e:9a:e4:85:4b:ad:db.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.17.4.61' (RSA) to the list of known hosts.
Welcome to Ubuntu 17.04 (GNU/Linux 4.10.0-42-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage


  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

0 packages can be updated.
0 updates are security updates.

New release '17.10' available.
Run 'do-release-upgrade' to upgrade to it.


Last login: Tue Dec 19 19:13:43 2017 from 10.0.2.2
```

```
ubuntu@rookdev-172-17-4-61:~$ docker info
Containers: 0
 Running: 0
 Paused: 0
 Stopped: 0
Images: 0
Server Version: 1.13.1
Storage Driver: aufs
 Root Dir: /var/lib/docker/aufs
 Backing Filesystem: extfs
 Dirs: 0
 Dirperm1 Supported: true
Logging Driver: json-file
Cgroup Driver: cgroupfs
Plugins: 
 Volume: local
 Network: bridge host macvlan null overlay
Swarm: inactive
Runtimes: runc
Default Runtime: runc
Init Binary: docker-init
containerd version:  (expected: aa8187dbd3b7ad67d8e5e3a15115d3eef43a7ed1)
runc version: N/A (expected: 9df8b306d01f59d3a8029be411de015b7304dd8f)
init version: N/A (expected: 949e6facb77383876aeff8a6944dde66b3089574)
Security Options:
 apparmor
 seccomp
  Profile: default
Kernel Version: 4.10.0-42-generic
Operating System: Ubuntu 17.04
OSType: linux
Architecture: x86_64
CPUs: 1
Total Memory: 1.951 GiB
Name: rookdev-172-17-4-61
ID: D6CX:4TWQ:VGDR:7X5Q:PNZH:F43O:5RCS:OYDD:GVBM:EOYJ:EKGI:JCYT
Docker Root Dir: /var/lib/docker
Debug Mode (client): false
Debug Mode (server): false
Registry: https://index.docker.io/v1/
WARNING: No swap limit support
Experimental: false
Insecure Registries:
 127.0.0.0/8
Live Restore Enabled: false
```

```
ubuntu@rookdev-172-17-4-61:~$ kubeadm token list --kubeconfig=.kube/config
TOKEN                     TTL       EXPIRES                USAGES                   DESCRIPTION   EXTRA GROUPS
86257a.d063ca0eedd9e165   14h       2017-12-20T10:40:36Z   authentication,signing   <none>        system:bootstrappers:kubeadm:default-node-token
```

```
ubuntu@rookdev-172-17-4-61:~$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 172.17.4.59:6443 --discovery-token-unsafe-skip-ca-verification
[preflight] Running pre-flight checks.
	[WARNING Service-Docker]: docker service is not enabled, please run 'systemctl enable docker.service'
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
ubuntu@rookdev-172-17-4-61:~$ kubectl get nodes
NAME                  STATUS                        ROLES     AGE       VERSION
kubedev-172-17-4-55   NotReady,SchedulingDisabled   <none>    17h       v1.9.0
kubedev-172-17-4-59   Ready                         master    1d        v1.9.0
kubedev-172-17-4-65   NotReady,SchedulingDisabled   <none>    4h        v1.9.0
rookdev-172-17-4-61   Ready                         <none>    2m        v1.9.0
rookdev-172-17-4-63   Ready                         <none>    8h        v1.9.0
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl get pods -n kube-system -o wide
NAME                                          READY     STATUS     RESTARTS   AGE       IP            NODE
etcd-kubedev-172-17-4-59                      1/1       Running    1          1d        172.17.4.59   kubedev-172-17-4-59
kube-apiserver-kubedev-172-17-4-59            1/1       Running    1          1d        172.17.4.59   kubedev-172-17-4-59
kube-controller-manager-kubedev-172-17-4-59   1/1       Running    3          1d        172.17.4.59   kubedev-172-17-4-59
kube-dns-6f4fd4bdf-r8c7h                      3/3       Running    3          1d        10.244.0.5    kubedev-172-17-4-59
kube-flannel-ds-2znmr                         1/1       Running    0          42m       172.17.4.63   rookdev-172-17-4-63
kube-flannel-ds-d5df7                         0/1       NodeLost   0          42m       <none>        kubedev-172-17-4-65
kube-flannel-ds-dm6l6                         1/1       Running    0          3m        172.17.4.61   rookdev-172-17-4-61
kube-flannel-ds-r59bm                         1/1       Running    0          42m       172.17.4.59   kubedev-172-17-4-59
kube-flannel-ds-rrkqn                         0/1       NodeLost   0          42m       <none>        kubedev-172-17-4-55
kube-proxy-fbbdd                              1/1       Running    3          8h        172.17.4.63   rookdev-172-17-4-63
kube-proxy-g6qr8                              1/1       NodeLost   3          4h        172.17.4.65   kubedev-172-17-4-65
kube-proxy-ls4nx                              1/1       NodeLost   5          17h       172.17.4.55   kubedev-172-17-4-55
kube-proxy-s6jtz                              1/1       Running    0          3m        172.17.4.61   rookdev-172-17-4-61
kube-proxy-xj2vb                              1/1       Running    1          1d        172.17.4.59   kubedev-172-17-4-59
kube-scheduler-kubedev-172-17-4-59            1/1       Running    3          1d        172.17.4.59   kubedev-172-17-4-59
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl scale deploy hello --replicas=3
deployment "hello" scaled
```

```
ubuntu@rookdev-172-17-4-61:~$ ip a show cni0
7: cni0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default qlen 1000
    link/ether 0a:58:0a:f4:02:01 brd ff:ff:ff:ff:ff:ff
    inet 10.244.2.1/24 scope global cni0
       valid_lft forever preferred_lft forever
    inet6 fe80::f050:5dff:fe9f:5140/64 scope link 
       valid_lft forever preferred_lft forever
```