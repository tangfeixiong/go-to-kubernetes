# A kubernetes workloads node on Ubuntu 16.04 Instance

## Table of content

1. Ubuntu repository
1. Docker
1. Kubernetes
1. [Join Cluster](#clusting)
1. [CNI networking](#kubernetes-cni-networking)
1. [Kubernetes Dashboard](#dashboard)

__Prerequisites__

Image from https://mirrors.tuna.tsinghua.edu.cn/ubuntu-cloud-images/

Launch VM _workerk8s1_ from OpenStack Dashboard 

Private IP is 10.20.30.207
```
wei@wei-ThinkPad-X1-Extreme:~$ ip netns
qrouter-f46292f9-3e0d-427e-83dc-13d803537360 (id: 2)
qdhcp-13d1dd1d-71d5-4b2e-8876-45158bdfe7f8 (id: 0)
qdhcp-cb0c8be8-b36b-4bb6-935d-4024d44c38bc (id: 1)
```

Login through net ns
```
wei@wei-ThinkPad-X1-Extreme:~$ sudo ip netns exec qrouter-f46292f9-3e0d-427e-83dc-13d803537360 ssh -i keypair-train.pem ubuntu@10.20.30.207
[sudo] password for wei: 
The authenticity of host '10.20.30.207 (10.20.30.207)' can't be established.
ECDSA key fingerprint is SHA256:xpp5x5eOfj0ALdnLajryMh1uekUVo7WiXbu5gsWvIps.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '10.20.30.207' (ECDSA) to the list of known hosts.
Welcome to Ubuntu 16.04.6 LTS (GNU/Linux 4.4.0-173-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

0 packages can be updated.
0 updates are security updates.



The programs included with the Ubuntu system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Ubuntu comes with ABSOLUTELY NO WARRANTY, to the extent permitted by
applicable law.

To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.
```

Trouble
```
ubuntu@workerk8s1:~$ hostname -i
hostname: Name or service not known
```

```
ubuntu@workerk8s1:~$ ping -c3 workerk8s1
ping: unknown host workerk8s1
```

```
ubuntu@workerk8s1:~$ v=$(ip addr | grep inet | head -3 | tail -1 | awk '{print $2}' | cut -d/ -f1) && sudo sed -i "1 i\
> $v $(hostname)" /etc/hosts
```

## Ubuntu repository mirror

Modify to cn.archive.ubuntu.com
```
ubuntu@workerk8s1:~$ sudo cp /etc/apt/sources.list /etc/apt/sources.list.orig
```

```
ubuntu@workerk8s1:~$ sudo sed -i "s/archive.ubuntu.com/cn.archive.ubuntu.com/" /etc/apt/sources.list
```

Local Kubernetes repository
```
ubuntu@workerk8s1:~$ sudo curl -jkSL http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   659  100   659    0     0  11265      0 --:--:-- --:--:-- --:--:-- 11362
OK
```

```
ubuntu@workerk8s1:~$ sudo curl -jkSL http://192.168.0.106:48080/etc0x2Fapt0x2Fkubernetes.list -o /etc/apt/sources.list.d/kubernetes.list
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    53  100    53    0     0   8081      0 --:--:-- --:--:-- --:--:--  8833
```

```
ubuntu@workerk8s1:~$ sudo sed -i "s%deb http://apt.kubernetes.io/%# &\ndeb http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/ kubernetes-xenial main%" /etc/apt/sources.list.d/kubernetes.list 
```

Cache
```
ubuntu@workerk8s1:~$ sudo apt-get update
Get:1 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial InRelease [8993 B]
Ign:2 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages
Get:2 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [33.3 kB]
Get:3 http://security.ubuntu.com/ubuntu xenial-security InRelease [109 kB]     
Get:4 http://cn.archive.ubuntu.com/ubuntu xenial InRelease [247 kB]
Get:5 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [819 kB]
Get:6 http://cn.archive.ubuntu.com/ubuntu xenial-updates InRelease [109 kB]
Get:7 http://cn.archive.ubuntu.com/ubuntu xenial-backports InRelease [107 kB]
Get:8 http://cn.archive.ubuntu.com/ubuntu xenial/main amd64 Packages [1201 kB]
Get:9 http://cn.archive.ubuntu.com/ubuntu xenial/main Translation-en [568 kB]  
Get:10 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [313 kB]
Get:11 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [479 kB]
Get:12 http://cn.archive.ubuntu.com/ubuntu xenial/restricted amd64 Packages [8344 B]
Get:13 http://cn.archive.ubuntu.com/ubuntu xenial/restricted Translation-en [2908 B]
Get:14 http://cn.archive.ubuntu.com/ubuntu xenial/universe amd64 Packages [7532 kB]
Get:15 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [196 kB]
Get:16 http://security.ubuntu.com/ubuntu xenial-security/multiverse amd64 Packages [5728 B]
Get:17 http://security.ubuntu.com/ubuntu xenial-security/multiverse Translation-en [2708 B]
Get:18 http://cn.archive.ubuntu.com/ubuntu xenial/universe Translation-en [4354 kB]
Get:19 http://cn.archive.ubuntu.com/ubuntu xenial/multiverse amd64 Packages [144 kB]
Get:20 http://cn.archive.ubuntu.com/ubuntu xenial/multiverse Translation-en [106 kB]
Get:21 http://cn.archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [1099 kB]
Get:22 http://cn.archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [421 kB]
Get:23 http://cn.archive.ubuntu.com/ubuntu xenial-updates/restricted amd64 Packages [7616 B]
Get:24 http://cn.archive.ubuntu.com/ubuntu xenial-updates/restricted Translation-en [2272 B]
Get:25 http://cn.archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [783 kB]
Get:26 http://cn.archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [328 kB]
Get:27 http://cn.archive.ubuntu.com/ubuntu xenial-updates/multiverse amd64 Packages [16.8 kB]
Get:28 http://cn.archive.ubuntu.com/ubuntu xenial-updates/multiverse Translation-en [8468 B]
Get:29 http://cn.archive.ubuntu.com/ubuntu xenial-backports/main amd64 Packages [7280 B]
Get:30 http://cn.archive.ubuntu.com/ubuntu xenial-backports/main Translation-en [4456 B]
Get:31 http://cn.archive.ubuntu.com/ubuntu xenial-backports/universe amd64 Packages [8064 B]
Get:32 http://cn.archive.ubuntu.com/ubuntu xenial-backports/universe Translation-en [4328 B]
Fetched 19.0 MB in 8s (2117 kB/s)                                              
Reading package lists... Done
```

### Docker
```
ubuntu@workerk8s1:~$ sudo apt-cache madison docker.io
 docker.io | 18.09.7-0ubuntu1~16.04.5 | http://cn.archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages
 docker.io | 18.09.7-0ubuntu1~16.04.5 | http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages
 docker.io | 1.10.3-0ubuntu6 | http://cn.archive.ubuntu.com/ubuntu xenial/universe amd64 Packages
```

```
ubuntu@workerk8s1:~$ sudo apt-get install -y docker.io
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  bridge-utils cgroupfs-mount containerd pigz runc ubuntu-fan
Suggested packages:
  mountall aufs-tools debootstrap docker-doc rinse zfs-fuse | zfsutils
The following NEW packages will be installed:
  bridge-utils cgroupfs-mount containerd docker.io pigz runc ubuntu-fan
0 upgraded, 7 newly installed, 0 to remove and 10 not upgraded.
Need to get 52.2 MB of archives.
After this operation, 257 MB of additional disk space will be used.
Get:1 http://cn.archive.ubuntu.com/ubuntu xenial/universe amd64 pigz amd64 2.3.1-2 [61.1 kB]
Get:2 http://cn.archive.ubuntu.com/ubuntu xenial/main amd64 bridge-utils amd64 1.5-9ubuntu1 [28.6 kB]
Get:3 http://cn.archive.ubuntu.com/ubuntu xenial/universe amd64 cgroupfs-mount all 1.2 [4970 B]
Get:4 http://cn.archive.ubuntu.com/ubuntu xenial-updates/universe amd64 runc amd64 1.0.0~rc7+git20190403.029124da-0ubuntu1~16.04.4 [1890 kB]
Get:5 http://cn.archive.ubuntu.com/ubuntu xenial-updates/universe amd64 containerd amd64 1.2.6-0ubuntu1~16.04.3 [19.7 MB]
Get:6 http://cn.archive.ubuntu.com/ubuntu xenial-updates/universe amd64 docker.io amd64 18.09.7-0ubuntu1~16.04.5 [30.4 MB]
Get:7 http://cn.archive.ubuntu.com/ubuntu xenial-updates/main amd64 ubuntu-fan all 0.12.8~16.04.3 [35.1 kB]
Fetched 52.2 MB in 10s (4855 kB/s)                                             
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LC_TIME = "zh_CN.UTF-8",
	LC_MONETARY = "zh_CN.UTF-8",
	LC_ADDRESS = "zh_CN.UTF-8",
	LC_TELEPHONE = "zh_CN.UTF-8",
	LC_NAME = "zh_CN.UTF-8",
	LC_MEASUREMENT = "zh_CN.UTF-8",
	LC_IDENTIFICATION = "zh_CN.UTF-8",
	LC_NUMERIC = "zh_CN.UTF-8",
	LC_PAPER = "zh_CN.UTF-8",
	LANG = "en_US.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to a fallback locale ("en_US.UTF-8").
locale: Cannot set LC_ALL to default locale: No such file or directory
Preconfiguring packages ...
Selecting previously unselected package pigz.
(Reading database ... 54249 files and directories currently installed.)
Preparing to unpack .../pigz_2.3.1-2_amd64.deb ...
Unpacking pigz (2.3.1-2) ...
Selecting previously unselected package bridge-utils.
Preparing to unpack .../bridge-utils_1.5-9ubuntu1_amd64.deb ...
Unpacking bridge-utils (1.5-9ubuntu1) ...
Selecting previously unselected package cgroupfs-mount.
Preparing to unpack .../cgroupfs-mount_1.2_all.deb ...
Unpacking cgroupfs-mount (1.2) ...
Selecting previously unselected package runc.
Preparing to unpack .../runc_1.0.0~rc7+git20190403.029124da-0ubuntu1~16.04.4_amd64.deb ...
Unpacking runc (1.0.0~rc7+git20190403.029124da-0ubuntu1~16.04.4) ...
Selecting previously unselected package containerd.
Preparing to unpack .../containerd_1.2.6-0ubuntu1~16.04.3_amd64.deb ...
Unpacking containerd (1.2.6-0ubuntu1~16.04.3) ...
Selecting previously unselected package docker.io.
Preparing to unpack .../docker.io_18.09.7-0ubuntu1~16.04.5_amd64.deb ...
Unpacking docker.io (18.09.7-0ubuntu1~16.04.5) ...
Selecting previously unselected package ubuntu-fan.
Preparing to unpack .../ubuntu-fan_0.12.8~16.04.3_all.deb ...
Unpacking ubuntu-fan (0.12.8~16.04.3) ...
Processing triggers for man-db (2.7.5-1) ...
Processing triggers for ureadahead (0.100.0-19.1) ...
Processing triggers for systemd (229-4ubuntu21.23) ...
Setting up pigz (2.3.1-2) ...
Setting up bridge-utils (1.5-9ubuntu1) ...
locale: Cannot set LC_ALL to default locale: No such file or directory
Setting up cgroupfs-mount (1.2) ...
Setting up runc (1.0.0~rc7+git20190403.029124da-0ubuntu1~16.04.4) ...
Setting up containerd (1.2.6-0ubuntu1~16.04.3) ...
Setting up docker.io (18.09.7-0ubuntu1~16.04.5) ...
locale: Cannot set LC_ALL to default locale: No such file or directory
Adding group `docker' (GID 116) ...
Done.
Setting up ubuntu-fan (0.12.8~16.04.3) ...
Processing triggers for ureadahead (0.100.0-19.1) ...
Processing triggers for systemd (229-4ubuntu21.23) ...
```

```
ubuntu@workerk8s1:~$ sudo usermod -aG docker ubuntu
```

```
ubuntu@workerk8s1:~$ sudo docker info
Containers: 0
 Running: 0
 Paused: 0
 Stopped: 0
Images: 0
Server Version: 18.09.7
Storage Driver: overlay2
 Backing Filesystem: extfs
 Supports d_type: true
 Native Overlay Diff: true
Logging Driver: json-file
Cgroup Driver: cgroupfs
Plugins:
 Volume: local
 Network: bridge host macvlan null overlay
 Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
Swarm: inactive
Runtimes: runc
Default Runtime: runc
Init Binary: docker-init
containerd version: 
runc version: N/A
init version: v0.18.0 (expected: fec3683b971d9c3ef73f284f176672c44b448662)
Security Options:
 apparmor
 seccomp
  Profile: default
Kernel Version: 4.4.0-173-generic
Operating System: Ubuntu 16.04.6 LTS
OSType: linux
Architecture: x86_64
CPUs: 1
Total Memory: 1.953GiB
Name: workerk8s1
ID: ZNKX:C2YV:IW35:JPTN:EKOG:3RMU:6KUV:Y3PE:4NSR:B25O:PGNJ:SNJI
Docker Root Dir: /var/lib/docker
Debug Mode (client): false
Debug Mode (server): false
Registry: https://index.docker.io/v1/
Labels:
Experimental: false
Insecure Registries:
 127.0.0.0/8
Live Restore Enabled: false

WARNING: No swap limit support
```

### Kubernetes
```
ubuntu@workerk8s1:~$ sudo apt-get install -y kubectl kubeadm kubelet
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  conntrack cri-tools ebtables kubernetes-cni socat
The following NEW packages will be installed:
  conntrack cri-tools ebtables kubeadm kubectl kubelet kubernetes-cni socat
0 upgraded, 8 newly installed, 0 to remove and 10 not upgraded.
Need to get 51.7 MB of archives.
After this operation, 272 MB of additional disk space will be used.
Get:1 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 cri-tools amd64 1.13.0-00 [8776 kB]
Get:2 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubernetes-cni amd64 0.7.5-00 [6473 kB]
Get:3 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubelet amd64 1.17.2-00 [19.2 MB]
Get:4 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubectl amd64 1.17.2-00 [8738 kB]
Get:5 http://192.168.0.106:48080/https0x3A0x2F0x2Fpackages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.17.2-00 [8061 kB]
Get:6 http://cn.archive.ubuntu.com/ubuntu xenial/main amd64 conntrack amd64 1:1.4.3-3 [27.3 kB]
Get:7 http://cn.archive.ubuntu.com/ubuntu xenial-updates/main amd64 ebtables amd64 2.0.10.4-3.4ubuntu2.16.04.2 [79.9 kB]
Get:8 http://cn.archive.ubuntu.com/ubuntu xenial/universe amd64 socat amd64 1.7.3.1-1 [321 kB]
Fetched 51.7 MB in 3s (16.1 MB/s)
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LC_TIME = "zh_CN.UTF-8",
	LC_MONETARY = "zh_CN.UTF-8",
	LC_ADDRESS = "zh_CN.UTF-8",
	LC_TELEPHONE = "zh_CN.UTF-8",
	LC_NAME = "zh_CN.UTF-8",
	LC_MEASUREMENT = "zh_CN.UTF-8",
	LC_IDENTIFICATION = "zh_CN.UTF-8",
	LC_NUMERIC = "zh_CN.UTF-8",
	LC_PAPER = "zh_CN.UTF-8",
	LANG = "en_US.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to a fallback locale ("en_US.UTF-8").
locale: Cannot set LC_ALL to default locale: No such file or directory
Selecting previously unselected package conntrack.
(Reading database ... 54563 files and directories currently installed.)
Preparing to unpack .../conntrack_1%3a1.4.3-3_amd64.deb ...
Unpacking conntrack (1:1.4.3-3) ...
Selecting previously unselected package cri-tools.
Preparing to unpack .../cri-tools_1.13.0-00_amd64.deb ...
Unpacking cri-tools (1.13.0-00) ...
Selecting previously unselected package ebtables.
Preparing to unpack .../ebtables_2.0.10.4-3.4ubuntu2.16.04.2_amd64.deb ...
Unpacking ebtables (2.0.10.4-3.4ubuntu2.16.04.2) ...
Selecting previously unselected package kubernetes-cni.
Preparing to unpack .../kubernetes-cni_0.7.5-00_amd64.deb ...
Unpacking kubernetes-cni (0.7.5-00) ...
Selecting previously unselected package socat.
Preparing to unpack .../socat_1.7.3.1-1_amd64.deb ...
Unpacking socat (1.7.3.1-1) ...
Selecting previously unselected package kubelet.
Preparing to unpack .../kubelet_1.17.2-00_amd64.deb ...
Unpacking kubelet (1.17.2-00) ...
Selecting previously unselected package kubectl.
Preparing to unpack .../kubectl_1.17.2-00_amd64.deb ...
Unpacking kubectl (1.17.2-00) ...
Selecting previously unselected package kubeadm.
Preparing to unpack .../kubeadm_1.17.2-00_amd64.deb ...
Unpacking kubeadm (1.17.2-00) ...
Processing triggers for man-db (2.7.5-1) ...
Processing triggers for ureadahead (0.100.0-19.1) ...
Processing triggers for systemd (229-4ubuntu21.23) ...
Setting up conntrack (1:1.4.3-3) ...
Setting up cri-tools (1.13.0-00) ...
Setting up ebtables (2.0.10.4-3.4ubuntu2.16.04.2) ...
update-rc.d: warning: start and stop actions are no longer supported; falling back to defaults
Setting up kubernetes-cni (0.7.5-00) ...
Setting up socat (1.7.3.1-1) ...
Setting up kubelet (1.17.2-00) ...
Setting up kubectl (1.17.2-00) ...
Setting up kubeadm (1.17.2-00) ...
Processing triggers for ureadahead (0.100.0-19.1) ...
Processing triggers for systemd (229-4ubuntu21.23) ...
```

### Image

From aliyuncs
```
ubuntu@workerk8s1:~$ sudo kubeadm config images pull --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers
W0204 19:10:03.385475    5836 version.go:101] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
W0204 19:10:03.385656    5836 version.go:102] falling back to the local client version: v1.17.2
W0204 19:10:03.385927    5836 validation.go:28] Cannot validate kubelet config - no validator is available
W0204 19:10:03.385952    5836 validation.go:28] Cannot validate kube-proxy config - no validator is available
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:v1.17.2
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:v1.17.2
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:v1.17.2
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.17.2
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.1
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:3.4.3-0
[config/images] Pulled registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:1.6.5
```

## Clustering

Join kubernetes worker with private IP as apiserver advertise address
```
ubuntu@workerk8s1:~$ sudo kubeadm join 10.20.30.56:6443 --token 9udxjm.1prjluofu8mw7zyv --discovery-token-ca-cert-hash sha256:e76a4e827e734a94738acb78edf4285393e73a5eb2731090422829071356ec49 --ignore-preflight-errors=NumCPU
W0204 19:15:12.969620    6243 join.go:346] [preflight] WARNING: JoinControlPane.controlPlane settings will be ignored when control-plane flag is not set.
[preflight] Running pre-flight checks
	[WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
[preflight] Reading configuration from the cluster...
[preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -oyaml'
[kubelet-start] Downloading configuration for the kubelet from the "kubelet-config-1.17" ConfigMap in the kube-system namespace
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Starting the kubelet
[kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster.
```

Bootstrap kubernetes worker with recommended cgroup driver
```
ubuntu@workerk8s1:~$ sudo kubeadm join 10.20.30.56:6443 --token h8cu4b.gb8znn0tj4vogkgd --discovery-token-ca-cert-hash sha256:80f3952e35b6e7c0c64e843d5404ebf8102c6d90d95a1869a9cc77768b2f327b --ignore-preflight-errors=NumCPU
W0204 22:50:39.482750   13853 join.go:346] [preflight] WARNING: JoinControlPane.controlPlane settings will be ignored when control-plane flag is not set.
[preflight] Running pre-flight checks
	[WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
[preflight] Reading configuration from the cluster...
[preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -oyaml'
[kubelet-start] Downloading configuration for the kubelet from the "kubelet-config-1.17" ConfigMap in the kube-system namespace
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Starting the kubelet
[kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster.
```

__Cluster info__
```
wei@wei-ThinkPad-X1-Extreme:~$ ./kubernetes/server/bin/kubectl get nodes
NAME         STATUS     ROLES    AGE    VERSION
masterk8s    NotReady   master   8h     v1.17.2
workerk8s1   NotReady   <none>   119s   v1.17.2
```

## Kubernetes CNI Networking

### Flannel

Reference
+ https://kubernetes.io/docs/concepts/cluster-administration/addons/
+ https://github.com/coreos/flannel/blob/master/Documentation/kubernetes.md
+ https://github.com/coreos/flannel

Investigate
```
ubuntu@workerk8s1:~$ sudo vi /etc/sysctl.conf 
```

```
ubuntu@workerk8s1:~$ sudo sysctl -p /etc/sysctl.conf
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-iptables = 1
```

```
ubuntu@workerk8s1:~$ sudo ufw allow 8285/udp
Rules updated
Rules updated (v6)
```

```
ubuntu@workerk8s1:~$ sudo ufw allow 8472/udp
Rules updated
Rules updated (v6)
```

```
ubuntu@masterk8s:~$ sudo systemctl stop kubelet
ubuntu@masterk8s:~$ sudo systemctl stop docker
ubuntu@masterk8s:~$ sudo iptables --flush
ubuntu@masterk8s:~$ sudo iptables -t nat --flush
ubuntu@masterk8s:~$ sudo systemctl start docker
ubuntu@masterk8s:~$ sudo systemctl start kubelet
ubuntu@masterk8s:~$ sudo ufw default allow incoming
Default incoming policy changed to 'allow'
(be sure to update your rules accordingly)
ubuntu@masterk8s:~$ sudo ufw default allow outgoing
Default outgoing policy changed to 'allow'
(be sure to update your rules accordingly)
ubuntu@masterk8s:~$ sudo iptables -P FORWARD ACCEPT
```

__From control-plane__
```
ubuntu@masterk8s:~$ sudo ufw allow 8285/udp
Rules updated
Rules updated (v6)
```

```
ubuntu@masterk8s:~$ sudo ufw allow 8472/udp
Rules updated
Rules updated (v6)
```

```
ubuntu@masterk8s:~$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 14416  100 14416    0     0  27097      0 --:--:-- --:--:-- --:--:-- 27097
```

```
ubuntu@masterk8s:~$ sed "s%10.244.0.0/16%10.10.0.0/16%"  kube-flannel.yml > flannel.yaml
```

apply
```
ubuntu@masterk8s:~$ kubectl apply -f flannel.yaml 
podsecuritypolicy.policy/psp.flannel.unprivileged created
clusterrole.rbac.authorization.k8s.io/flannel created
clusterrolebinding.rbac.authorization.k8s.io/flannel created
serviceaccount/flannel created
configmap/kube-flannel-cfg created
daemonset.apps/kube-flannel-ds-amd64 created
daemonset.apps/kube-flannel-ds-arm64 created
daemonset.apps/kube-flannel-ds-arm created
daemonset.apps/kube-flannel-ds-ppc64le created
daemonset.apps/kube-flannel-ds-s390x created
```

```
ubuntu@masterk8s:~$ kubectl get nodes
NAME         STATUS   ROLES    AGE   VERSION
masterk8s    Ready    master   9h    v1.17.2
workerk8s1   Ready    <none>   12m   v1.17.2
```

检查flannel POD的日志有异常
```
ubuntu@masterk8s:~$ kubectl -n kube-system get pods
ubuntu@masterk8s:~$ kubectl -n kube-system logs pods/...
```

### Weave

Reference
+ https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/
+ https://www.weave.works/docs/net/latest/kubernetes/kube-addon/

Manifest
```
ubuntu@masterk8s:~$ curl -jkSL https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n') -o weave.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   132  100   132    0     0     91      0  0:00:01  0:00:01 --:--:--    91
100 10814  100 10814    0     0   6179      0  0:00:01  0:00:01 --:--:-- 10.3M
```

```
ubuntu@masterk8s:~$ kubectl apply -f weave.yaml 
serviceaccount/weave-net created
clusterrole.rbac.authorization.k8s.io/weave-net created
clusterrolebinding.rbac.authorization.k8s.io/weave-net created
role.rbac.authorization.k8s.io/weave-net created
rolebinding.rbac.authorization.k8s.io/weave-net created
daemonset.apps/weave-net created
```


### Calico

Refence
+ https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/
+ https://docs.projectcalico.org/getting-started/kubernetes/quickstart

Investigate
```
ubuntu@workerk8s1:~$ sudo systemctl stop kubelet
```

```
ubuntu@workerk8s1:~$ sudo systemctl stop docker
```

```
ubuntu@workerk8s1:~$ sudo iptables --flush
```

```
ubuntu@workerk8s1:~$ sudo iptables -t nat --flush
```

```
ubuntu@workerk8s1:~$ sudo systemctl start docker
```

```
ubuntu@workerk8s1:~$ sudo systemctl start kubelet
```

```
ubuntu@workerk8s1:~$ sudo ufw default allow incoming
Default incoming policy changed to 'allow'
(be sure to update your rules accordingly)
```

```
ubuntu@workerk8s1:~$ sudo ufw default allow outgoing
Default outgoing policy changed to 'allow'
(be sure to update your rules accordingly)
```

```
ubuntu@workerk8s1:~$ sudo iptables -P FORWARD ACCEPT
```

Download
```
ubuntu@masterk8s:~$ curl -jkSL https://docs.projectcalico.org/v3.11/manifests/calico.yaml -o calico.yaml.orig
```

Pod network
```
ubuntu@masterk8s:~$ sed "s%192.168.0.0/16%10.10.0.0/16%" calico.yaml.orig > calico.yaml
```

Deploy
```
ubuntu@masterk8s:~$ kubectl apply -f calico.yaml
configmap/calico-config created
customresourcedefinition.apiextensions.k8s.io/felixconfigurations.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/ipamblocks.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/blockaffinities.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/ipamhandles.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/ipamconfigs.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/bgppeers.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/bgpconfigurations.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/ippools.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/hostendpoints.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/clusterinformations.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/globalnetworkpolicies.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/globalnetworksets.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/networkpolicies.crd.projectcalico.org created
customresourcedefinition.apiextensions.k8s.io/networksets.crd.projectcalico.org created
clusterrole.rbac.authorization.k8s.io/calico-kube-controllers created
clusterrolebinding.rbac.authorization.k8s.io/calico-kube-controllers created
clusterrole.rbac.authorization.k8s.io/calico-node created
clusterrolebinding.rbac.authorization.k8s.io/calico-node created
daemonset.apps/calico-node created
serviceaccount/calico-node created
deployment.apps/calico-kube-controllers created
serviceaccount/calico-kube-controllers created
```

Explore
```
wei@wei-ThinkPad-X1-Extreme:~$ kubectl -n kube-system get pods
NAME                                       READY   STATUS    RESTARTS   AGE
calico-kube-controllers-5b644bc49c-2z2hh   1/1     Running   0          91m
calico-node-728pd                          0/1     Running   0          91m
calico-node-k949d                          0/1     Running   0          91m
coredns-7f9c544f75-lvhq2                   1/1     Running   0          4h39m
coredns-7f9c544f75-mwbz7                   1/1     Running   0          4h39m
etcd-masterk8s                             1/1     Running   0          4h40m
kube-apiserver-masterk8s                   1/1     Running   0          4h40m
kube-controller-manager-masterk8s          1/1     Running   0          4h40m
kube-proxy-krrj2                           1/1     Running   3          4h29m
kube-proxy-xcgcl                           1/1     Running   0          4h39m
kube-scheduler-masterk8s                   1/1     Running   0          4h40m
```

## Management

### Dashboard

Trouble
+ https://github.com/kubernetes/dashboard/issues/2799
+ https://github.com/kubernetes-sigs/kubespray/issues/5188
+ https://github.com/kubernetes/dashboard/issues/3789

__v1.10.1__
```
ubuntu@masterk8s:~$ curl -jkSL https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml -o kubernetes-dashboard-v1.10.1.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  4577  100  4577    0     0   5830      0 --:--:-- --:--:-- --:--:--  5830
```

```
ubuntu@workerk8s1:~$ curl http://192.168.0.106:48080/k8s.gcr.io0x2Fkubernetes-dashboard-amd64_v1.10.1.tar | docker load 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  116M  100  116M    0     0   155M      0 --:--:-- --:--:-- --:--:--  155M
fbdfe08b001c: Loading layer  122.3MB/122.3MB
Loaded image: k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1
```

```
ubuntu@masterk8s:~$ kubectl apply -f kubernetes-dashboard-v1.10.1.yaml 
secret/kubernetes-dashboard-certs created
serviceaccount/kubernetes-dashboard created
role.rbac.authorization.k8s.io/kubernetes-dashboard-minimal created
rolebinding.rbac.authorization.k8s.io/kubernetes-dashboard-minimal created
deployment.apps/kubernetes-dashboard created
service/kubernetes-dashboard created
```

```
wei@wei-ThinkPad-X1-Extreme:~$ kubectl -n kube-system get pods
NAME                                       READY   STATUS    RESTARTS   AGE
calico-kube-controllers-5b644bc49c-2z2hh   1/1     Running   0          91m
calico-node-728pd                          0/1     Running   0          91m
calico-node-k949d                          0/1     Running   0          91m
coredns-7f9c544f75-lvhq2                   1/1     Running   0          4h39m
coredns-7f9c544f75-mwbz7                   1/1     Running   0          4h39m
etcd-masterk8s                             1/1     Running   0          4h40m
kube-apiserver-masterk8s                   1/1     Running   0          4h40m
kube-controller-manager-masterk8s          1/1     Running   0          4h40m
kube-proxy-krrj2                           1/1     Running   3          4h29m
kube-proxy-xcgcl                           1/1     Running   0          4h39m
kube-scheduler-masterk8s                   1/1     Running   0          4h40m
kubernetes-dashboard-7c54d59f66-bnnrt      1/1     Running   0          101m
```

Expose as NodePort
```
wei@wei-ThinkPad-X1-Extreme:~$ ./kubernetes/server/bin/kubectl expose service kubernetes-dashboard -n kubernetes-dashboard --port=443 --target-port=443 --type=NodePort --name=kubernetes-dashboard-https
service/kubernetes-dashboard-https exposed
```

__Dashboard Token__
```
ubuntu@masterk8s:~$ cat dashboard-admin-user.yaml 
# https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md
# $ kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kube-system
  # namespace: kubernetes-dashboard

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kube-system
  # namespace: kubernetes-dashboard

---
```

```
ubuntu@masterk8s:~$ kubectl apply -f dashboard-admin-user.yaml 
```

Browse page snapshots
+ ./Screenshot from 2020-02-05 10-49-17.png


__v2__

Not verified
```
ubuntu@masterk8s:~$ curl -jkSL  https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-rc3/aio/deploy/recommended.yaml -o kubernetes-dashboard-v2.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  7566  100  7566    0     0  14063      0 --:--:-- --:--:-- --:--:-- 14037
```

```
ubuntu@masterk8s:~$ kubectl apply -f kubernetes-dashboard-v2.yaml 
namespace/kubernetes-dashboard created
serviceaccount/kubernetes-dashboard created
service/kubernetes-dashboard created
secret/kubernetes-dashboard-certs created
secret/kubernetes-dashboard-csrf created
secret/kubernetes-dashboard-key-holder created
configmap/kubernetes-dashboard-settings created
role.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrole.rbac.authorization.k8s.io/kubernetes-dashboard created
rolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
deployment.apps/kubernetes-dashboard created
service/dashboard-metrics-scraper created
deployment.apps/dashboard-metrics-scraper created
```

```
ubuntu@masterk8s:~$ kubectl get ns
NAME                   STATUS   AGE
default                Active   9h
kube-node-lease        Active   9h
kube-public            Active   9h
kube-system            Active   9h
kubernetes-dashboard   Active   8h
```

Remove
```
ubuntu@masterk8s:~$ kubectl delete -f kubernetes-dashboard-v2.yaml 
namespace "kubernetes-dashboard" deleted
serviceaccount "kubernetes-dashboard" deleted
service "kubernetes-dashboard" deleted
secret "kubernetes-dashboard-certs" deleted
secret "kubernetes-dashboard-csrf" deleted
secret "kubernetes-dashboard-key-holder" deleted
configmap "kubernetes-dashboard-settings" deleted
role.rbac.authorization.k8s.io "kubernetes-dashboard" deleted
clusterrole.rbac.authorization.k8s.io "kubernetes-dashboard" deleted
rolebinding.rbac.authorization.k8s.io "kubernetes-dashboard" deleted
clusterrolebinding.rbac.authorization.k8s.io "kubernetes-dashboard" deleted
deployment.apps "kubernetes-dashboard" deleted
service "dashboard-metrics-scraper" deleted
deployment.apps "dashboard-metrics-scraper" deleted
```





