# Kubernetes Devel Env

## V1.17.x

Feb 7, 2020
+ Ubuntu bionic
+ kubernetes v1.17.2 master
+ VirtualBox/Vagrant
    - https://github.com/tangfeixiong/lang-learn/tree/master/python/scikit-learn-lab

Feb 27, 2020
+ CentOS 7.7
+ Kubernetes v1.17.3 minon
+ VirtualBox/Vagrant
    - https://github.com/tangfeixiong/go-to-kubernetes/blob/master/Vagrantfile
    - https://github.com/tangfeixiong/go-to-kubernetes/blob/master/vagrantup.md

### Feb 27 2020

Host
```
[vagrant@localhost ~]$ sudo hostnamectl
   Static hostname: localhost.localdomain
         Icon name: computer-vm
           Chassis: vm
        Machine ID: 4ebecf3991fb44aa9bd4f2456a477cf5
           Boot ID: 2fd487d37d7544bf9b45b8f887851ed3
    Virtualization: kvm
  Operating System: CentOS Linux 7 (Core)
       CPE OS Name: cpe:/o:centos:centos:7
            Kernel: Linux 3.10.0-1062.9.1.el7.x86_64
      Architecture: x86-64
```

VirtualBox private networking
```
[vagrant@localhost ~]$ ip addr show eth1 | grep inet | head -1 | awk '{print $2}' | cut -d/ -f1
172.28.128.3
```

__Docker__

see in ./vagrantup.md

version
```
[vagrant@localhost ~]$ docker version
Client: Docker Engine - Community
 Version:           19.03.6
 API version:       1.40
 Go version:        go1.12.16
 Git commit:        369ce74a3c
 Built:             Thu Feb 13 01:29:29 2020
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.6
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.12.16
  Git commit:       369ce74a3c
  Built:            Thu Feb 13 01:28:07 2020
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.10
  GitCommit:        b34a5c8af56e510852c35414db4c1f4fa6172339
 runc:
  Version:          1.0.0-rc8+dev
  GitCommit:        3e425f80a8c931f88e6d94a8c831b9d5aa481657
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683
```

info
```
[vagrant@localhost ~]$ docker info
Client:
 Debug Mode: false

Server:
 Containers: 0
  Running: 0
  Paused: 0
  Stopped: 0
 Images: 0
 Server Version: 19.03.6
 Storage Driver: overlay2
  Backing Filesystem: xfs
  Supports d_type: true
  Native Overlay Diff: true
 Logging Driver: json-file
 Cgroup Driver: cgroupfs
 Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
 Swarm: inactive
 Runtimes: runc
 Default Runtime: runc
 Init Binary: docker-init
 containerd version: b34a5c8af56e510852c35414db4c1f4fa6172339
 runc version: 3e425f80a8c931f88e6d94a8c831b9d5aa481657
 init version: fec3683
 Security Options:
  seccomp
   Profile: default
 Kernel Version: 3.10.0-1062.9.1.el7.x86_64
 Operating System: CentOS Linux 7 (Core)
 OSType: linux
 Architecture: x86_64
 CPUs: 1
 Total Memory: 1.795GiB
 Name: localhost.localdomain
 ID: XKOA:Z5HV:EDRR:FBWD:PVBW:5P2M:BB5N:NGN4:V7X6:BTCR:PLXE:UBIU
 Docker Root Dir: /var/lib/docker
 Debug Mode: false
 Registry: https://index.docker.io/v1/
 Labels:
 Experimental: false
 Insecure Registries:
  127.0.0.0/8
 Live Restore Enabled: false

WARNING: bridge-nf-call-iptables is disabled
WARNING: bridge-nf-call-ip6tables is disabled
```

firewall
```
[vagrant@localhost ~]$ sudo firewall-cmd --state
not running
```

iptables
```
[vagrant@localhost ~]$ sudo iptables -L -vn
Chain INPUT (policy ACCEPT 40528 packets, 60M bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain FORWARD (policy DROP 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DOCKER-USER  all  --  *      *       0.0.0.0/0            0.0.0.0/0           
    0     0 DOCKER-ISOLATION-STAGE-1  all  --  *      *       0.0.0.0/0            0.0.0.0/0           
    0     0 ACCEPT     all  --  *      docker0  0.0.0.0/0            0.0.0.0/0            ctstate RELATED,ESTABLISHED
    0     0 DOCKER     all  --  *      docker0  0.0.0.0/0            0.0.0.0/0           
    0     0 ACCEPT     all  --  docker0 !docker0  0.0.0.0/0            0.0.0.0/0           
    0     0 ACCEPT     all  --  docker0 docker0  0.0.0.0/0            0.0.0.0/0           

Chain OUTPUT (policy ACCEPT 25241 packets, 1186K bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain DOCKER (1 references)
 pkts bytes target     prot opt in     out     source               destination         

Chain DOCKER-ISOLATION-STAGE-1 (1 references)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DOCKER-ISOLATION-STAGE-2  all  --  docker0 !docker0  0.0.0.0/0            0.0.0.0/0           
    0     0 RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0           

Chain DOCKER-ISOLATION-STAGE-2 (1 references)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DROP       all  --  *      docker0  0.0.0.0/0            0.0.0.0/0           
    0     0 RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0           

Chain DOCKER-USER (1 references)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0           
```

NAT
```
[vagrant@localhost ~]$ sudo iptables -t nat -L -vn
Chain PREROUTING (policy ACCEPT 24 packets, 9914 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    2   620 DOCKER     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL

Chain INPUT (policy ACCEPT 24 packets, 9914 bytes)
 pkts bytes target     prot opt in     out     source               destination         

Chain OUTPUT (policy ACCEPT 341 packets, 28871 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 DOCKER     all  --  *      *       0.0.0.0/0           !127.0.0.0/8          ADDRTYPE match dst-type LOCAL

Chain POSTROUTING (policy ACCEPT 341 packets, 28871 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 MASQUERADE  all  --  *      !docker0  172.17.0.0/16        0.0.0.0/0           

Chain DOCKER (2 references)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 RETURN     all  --  docker0 *       0.0.0.0/0            0.0.0.0/0           
```

__reference__

https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

Google kubernetes CentOS 7 repository
```
[vagrant@localhost ~]$ cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
> [kubernetes]
> name=Kubernetes
> baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
> enabled=1
> gpgcheck=1
> repo_gpgcheck=1
> gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
> EOF
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
```

package, note Google required accessable
```
[vagrant@localhost ~]$ sudo yum list kubeadm kubectl kubelet
Failed to set locale, defaulting to C
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirror.bit.edu.cn
 * extras: mirrors.huaweicloud.com
 * updates: mirror.bit.edu.cn
Available Packages
kubeadm.x86_64                                                           1.17.3-0                                                            kubernetes
kubectl.x86_64                                                           1.17.3-0                                                            kubernetes
kubelet.x86_64                                                           1.17.3-0                                                            kubernetes
```

selinux
```
[vagrant@localhost ~]$ sudo sed -i 's/^SELINUX=enforcing$/# &\nSELINUX=permissive/' /etc/selinux/config
```

disable SELINUX now
```
[vagrant@localhost ~]$ sudo setenforce 0
[vagrant@localhost ~]$ sudo getenforce 
Permissive
```

bridge_net_filter is disabled
```
[vagrant@localhost ~]$ grep -r 'net.bridge' /usr/lib/sysctl.d/
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-ip6tables = 0
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-iptables = 0
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-arptables = 0
```

modify
```
[vagrant@localhost ~]$ cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
> net.bridge.bridge-nf-call-ip6tables = 1
> net.bridge.bridge-nf-call-iptables = 1
> EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
```

and enable
```
[vagrant@localhost ~]$ sudo sysctl --system
* Applying /usr/lib/sysctl.d/00-system.conf ...
net.bridge.bridge-nf-call-ip6tables = 0
net.bridge.bridge-nf-call-iptables = 0
net.bridge.bridge-nf-call-arptables = 0
* Applying /usr/lib/sysctl.d/10-default-yama-scope.conf ...
kernel.yama.ptrace_scope = 0
* Applying /usr/lib/sysctl.d/50-default.conf ...
kernel.sysrq = 16
kernel.core_uses_pid = 1
net.ipv4.conf.default.rp_filter = 1
net.ipv4.conf.all.rp_filter = 1
net.ipv4.conf.default.accept_source_route = 0
net.ipv4.conf.all.accept_source_route = 0
net.ipv4.conf.default.promote_secondaries = 1
net.ipv4.conf.all.promote_secondaries = 1
fs.protected_hardlinks = 1
fs.protected_symlinks = 1
* Applying /etc/sysctl.d/99-sysctl.conf ...
* Applying /etc/sysctl.d/k8s.conf ...
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
* Applying /etc/sysctl.conf ...
```

module
```
[vagrant@localhost ~]$ sudo lsmod | grep br_netfilter
br_netfilter           22256  0 
bridge                151336  1 br_netfilter
```

or
```
[vagrant@localhost ~]$ sudo cat /proc/modules | grep br_netfilter
br_netfilter 22256 0 - Live 0xffffffffc04e7000
bridge 151336 1 br_netfilter, Live 0xffffffffc04b5000
```

__kubernetes__

bin
```
[vagrant@localhost ~]$ sudo yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
Failed to set locale, defaulting to C
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirror.bit.edu.cn
 * extras: mirrors.huaweicloud.com
 * updates: mirror.bit.edu.cn
Resolving Dependencies
--> Running transaction check
---> Package kubeadm.x86_64 0:1.17.3-0 will be installed
--> Processing Dependency: kubernetes-cni >= 0.7.5 for package: kubeadm-1.17.3-0.x86_64
--> Processing Dependency: cri-tools >= 1.13.0 for package: kubeadm-1.17.3-0.x86_64
---> Package kubectl.x86_64 0:1.17.3-0 will be installed
---> Package kubelet.x86_64 0:1.17.3-0 will be installed
--> Processing Dependency: socat for package: kubelet-1.17.3-0.x86_64
--> Processing Dependency: conntrack for package: kubelet-1.17.3-0.x86_64
--> Running transaction check
---> Package conntrack-tools.x86_64 0:1.4.4-5.el7_7.2 will be installed
--> Processing Dependency: libnetfilter_cttimeout.so.1(LIBNETFILTER_CTTIMEOUT_1.1)(64bit) for package: conntrack-tools-1.4.4-5.el7_7.2.x86_64
--> Processing Dependency: libnetfilter_cttimeout.so.1(LIBNETFILTER_CTTIMEOUT_1.0)(64bit) for package: conntrack-tools-1.4.4-5.el7_7.2.x86_64
--> Processing Dependency: libnetfilter_cthelper.so.0(LIBNETFILTER_CTHELPER_1.0)(64bit) for package: conntrack-tools-1.4.4-5.el7_7.2.x86_64
--> Processing Dependency: libnetfilter_queue.so.1()(64bit) for package: conntrack-tools-1.4.4-5.el7_7.2.x86_64
--> Processing Dependency: libnetfilter_cttimeout.so.1()(64bit) for package: conntrack-tools-1.4.4-5.el7_7.2.x86_64
--> Processing Dependency: libnetfilter_cthelper.so.0()(64bit) for package: conntrack-tools-1.4.4-5.el7_7.2.x86_64
---> Package cri-tools.x86_64 0:1.13.0-0 will be installed
---> Package kubernetes-cni.x86_64 0:0.7.5-0 will be installed
---> Package socat.x86_64 0:1.7.3.2-2.el7 will be installed
--> Running transaction check
---> Package libnetfilter_cthelper.x86_64 0:1.0.0-10.el7_7.1 will be installed
---> Package libnetfilter_cttimeout.x86_64 0:1.0.0-6.el7_7.1 will be installed
---> Package libnetfilter_queue.x86_64 0:1.0.2-2.el7_2 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

=======================================================================================================================================================
 Package                                     Arch                        Version                                 Repository                       Size
=======================================================================================================================================================
Installing:
 kubeadm                                     x86_64                      1.17.3-0                                kubernetes                      8.7 M
 kubectl                                     x86_64                      1.17.3-0                                kubernetes                      9.4 M
 kubelet                                     x86_64                      1.17.3-0                                kubernetes                       20 M
Installing for dependencies:
 conntrack-tools                             x86_64                      1.4.4-5.el7_7.2                         updates                         187 k
 cri-tools                                   x86_64                      1.13.0-0                                kubernetes                      5.1 M
 kubernetes-cni                              x86_64                      0.7.5-0                                 kubernetes                       10 M
 libnetfilter_cthelper                       x86_64                      1.0.0-10.el7_7.1                        updates                          18 k
 libnetfilter_cttimeout                      x86_64                      1.0.0-6.el7_7.1                         updates                          18 k
 libnetfilter_queue                          x86_64                      1.0.2-2.el7_2                           base                             23 k
 socat                                       x86_64                      1.7.3.2-2.el7                           base                            290 k

Transaction Summary
=======================================================================================================================================================
Install  3 Packages (+7 Dependent packages)

Total download size: 54 M
Installed size: 243 M
Downloading packages:
(1/10): conntrack-tools-1.4.4-5.el7_7.2.x86_64.rpm                                                                              | 187 kB  00:00:01     
warning: /var/cache/yum/x86_64/7/kubernetes/packages/14bfe6e75a9efc8eca3f638eb22c7e2ce759c67f95b43b16fae4ebabde1549f3-cri-tools-1.13.0-0.x86_64.rpm: Header V4 RSA/SHA512 Signature, key ID 3e1ba8d5: NOKEY
Public key for 14bfe6e75a9efc8eca3f638eb22c7e2ce759c67f95b43b16fae4ebabde1549f3-cri-tools-1.13.0-0.x86_64.rpm is not installed
(2/10): 14bfe6e75a9efc8eca3f638eb22c7e2ce759c67f95b43b16fae4ebabde1549f3-cri-tools-1.13.0-0.x86_64.rpm                          | 5.1 MB  00:00:23     
(3/10): d0c889d5fae925a9266aa521f6aed361ea53792c86a51c38f5d92cd3f03e2d30-kubeadm-1.17.3-0.x86_64.rpm                            | 8.7 MB  00:00:53     
(4/10): 35625b6ab1da6c58ce4946742181c0dcf9ac9b6c2b5bea2c13eed4876024c342-kubectl-1.17.3-0.x86_64.rpm                            | 9.4 MB  00:00:36     
(5/10): libnetfilter_cthelper-1.0.0-10.el7_7.1.x86_64.rpm                                                                       |  18 kB  00:00:00     
(6/10): libnetfilter_queue-1.0.2-2.el7_2.x86_64.rpm                                                                             |  23 kB  00:00:01     
(7/10): libnetfilter_cttimeout-1.0.0-6.el7_7.1.x86_64.rpm                                                                       |  18 kB  00:00:01     
(8/10): socat-1.7.3.2-2.el7.x86_64.rpm                                                                                          | 290 kB  00:00:30     
(9/10): 548a0dcd865c16a50980420ddfa5fbccb8b59621179798e6dc905c9bf8af3b34-kubernetes-cni-0.7.5-0.x86_64.rpm                      |  10 MB  00:00:54     
(10/10): 3f1db71d0bb6d72bc956d788ffee737714e5717c629b26355a2dcf1dba4ad231-kubelet-1.17.3-0.x86_64.rpm                           |  20 MB  00:01:29     
-------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                  391 kB/s |  54 MB  00:02:22     
Retrieving key from https://packages.cloud.google.com/yum/doc/yum-key.gpg
Importing GPG key 0xA7317B0F:
 Userid     : "Google Cloud Packages Automatic Signing Key <gc-team@google.com>"
 Fingerprint: d0bc 747f d8ca f711 7500 d6fa 3746 c208 a731 7b0f
 From       : https://packages.cloud.google.com/yum/doc/yum-key.gpg
Retrieving key from https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
Importing GPG key 0x3E1BA8D5:
 Userid     : "Google Cloud Packages RPM Signing Key <gc-team@google.com>"
 Fingerprint: 3749 e1ba 95a8 6ce0 5454 6ed2 f09c 394c 3e1b a8d5
 From       : https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : libnetfilter_cttimeout-1.0.0-6.el7_7.1.x86_64                                                                                      1/10 
  Installing : socat-1.7.3.2-2.el7.x86_64                                                                                                         2/10 
  Installing : kubectl-1.17.3-0.x86_64                                                                                                            3/10 
  Installing : cri-tools-1.13.0-0.x86_64                                                                                                          4/10 
  Installing : libnetfilter_queue-1.0.2-2.el7_2.x86_64                                                                                            5/10 
  Installing : libnetfilter_cthelper-1.0.0-10.el7_7.1.x86_64                                                                                      6/10 
  Installing : conntrack-tools-1.4.4-5.el7_7.2.x86_64                                                                                             7/10 
  Installing : kubernetes-cni-0.7.5-0.x86_64                                                                                                      8/10 
  Installing : kubelet-1.17.3-0.x86_64                                                                                                            9/10 
  Installing : kubeadm-1.17.3-0.x86_64                                                                                                           10/10 
  Verifying  : libnetfilter_cthelper-1.0.0-10.el7_7.1.x86_64                                                                                      1/10 
  Verifying  : conntrack-tools-1.4.4-5.el7_7.2.x86_64                                                                                             2/10 
  Verifying  : kubelet-1.17.3-0.x86_64                                                                                                            3/10 
  Verifying  : kubeadm-1.17.3-0.x86_64                                                                                                            4/10 
  Verifying  : libnetfilter_queue-1.0.2-2.el7_2.x86_64                                                                                            5/10 
  Verifying  : cri-tools-1.13.0-0.x86_64                                                                                                          6/10 
  Verifying  : kubectl-1.17.3-0.x86_64                                                                                                            7/10 
  Verifying  : kubernetes-cni-0.7.5-0.x86_64                                                                                                      8/10 
  Verifying  : socat-1.7.3.2-2.el7.x86_64                                                                                                         9/10 
  Verifying  : libnetfilter_cttimeout-1.0.0-6.el7_7.1.x86_64                                                                                     10/10 

Installed:
  kubeadm.x86_64 0:1.17.3-0                         kubectl.x86_64 0:1.17.3-0                         kubelet.x86_64 0:1.17.3-0                        

Dependency Installed:
  conntrack-tools.x86_64 0:1.4.4-5.el7_7.2            cri-tools.x86_64 0:1.13.0-0                         kubernetes-cni.x86_64 0:0.7.5-0              
  libnetfilter_cthelper.x86_64 0:1.0.0-10.el7_7.1     libnetfilter_cttimeout.x86_64 0:1.0.0-6.el7_7.1     libnetfilter_queue.x86_64 0:1.0.2-2.el7_2    
  socat.x86_64 0:1.7.3.2-2.el7                       

Complete!
[vagrant@localhost ~]$ kubectl version --client
```

kubectl
```
Client Version: version.Info{Major:"1", Minor:"17", GitVersion:"v1.17.3", GitCommit:"06ad960bfd03b39c8310aaf92d1e7c12ce618213", GitTreeState:"clean", BuildDate:"2020-02-11T18:14:22Z", GoVersion:"go1.13.6", Compiler:"gc", Platform:"linux/amd64"}
```

enable auto reboot
```
[vagrant@localhost ~]$ sudo systemctl enable --now kubelet
```

__Clustering__

https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/

Trouble 1#
```
[vagrant@localhost ~]$ sudo kubeadm join 172.28.128.4:6443 --token p3hpp3.ehka0529vafadvax --discovery-token-ca-cert-hash sha256:c86e295d23dc1189911b9a14ba19245c4bc33191c0b0de2a9576a45e3c9e2e9e 
W0226 22:20:40.004590   31972 join.go:346] [preflight] WARNING: JoinControlPane.controlPlane settings will be ignored when control-plane flag is not set.
[preflight] Running pre-flight checks
	[WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
error execution phase preflight: [preflight] Some fatal errors occurred:
	[ERROR Swap]: running with swap on is not supported. Please disable swap
[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
To see the stack trace of this error execute with --v=5 or higher
```

https://github.com/kubernetes/kubeadm/issues/610
```
[vagrant@localhost ~]$ sudo swapoff -a
```

succeeded
```
[vagrant@localhost ~]$ sudo kubeadm join 172.28.128.4:6443 --token p3hpp3.ehka0529vafadvax --discovery-token-ca-cert-hash sha256:c86e295d23dc1189911b9a14ba19245c4bc33191c0b0de2a9576a45e3c9e2e9e --ignore-preflight-errors Swap
W0226 22:41:34.843088     353 join.go:346] [preflight] WARNING: JoinControlPane.controlPlane settings will be ignored when control-plane flag is not set.
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

verify
```
[vagrant@localhost ~]$ sudo systemctl status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/usr/lib/systemd/system/kubelet.service; enabled; vendor preset: disabled)
  Drop-In: /usr/lib/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Wed 2020-02-26 22:41:35 UTC; 2min 17s ago
     Docs: https://kubernetes.io/docs/
 Main PID: 496 (kubelet)
    Tasks: 14
   Memory: 33.7M
   CGroup: /system.slice/kubelet.service
           └─496 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --config=/...

Feb 26 22:43:31 localhost.localdomain kubelet[496]: W0226 22:43:31.162496     496 cni.go:237] Unable to update cni config: no networks found ...i/net.d
Feb 26 22:43:31 localhost.localdomain kubelet[496]: E0226 22:43:31.797137     496 kubelet.go:2183] Container runtime network not ready: Netwo...ialized
Feb 26 22:43:36 localhost.localdomain kubelet[496]: W0226 22:43:36.163088     496 cni.go:237] Unable to update cni config: no networks found ...i/net.d
Feb 26 22:43:36 localhost.localdomain kubelet[496]: E0226 22:43:36.806101     496 kubelet.go:2183] Container runtime network not ready: Netwo...ialized
Feb 26 22:43:41 localhost.localdomain kubelet[496]: W0226 22:43:41.163355     496 cni.go:237] Unable to update cni config: no networks found ...i/net.d
Feb 26 22:43:41 localhost.localdomain kubelet[496]: E0226 22:43:41.814674     496 kubelet.go:2183] Container runtime network not ready: Netwo...ialized
Feb 26 22:43:46 localhost.localdomain kubelet[496]: W0226 22:43:46.163874     496 cni.go:237] Unable to update cni config: no networks found ...i/net.d
Feb 26 22:43:46 localhost.localdomain kubelet[496]: E0226 22:43:46.824190     496 kubelet.go:2183] Container runtime network not ready: Netwo...ialized
Feb 26 22:43:51 localhost.localdomain kubelet[496]: W0226 22:43:51.165022     496 cni.go:237] Unable to update cni config: no networks found ...i/net.d
Feb 26 22:43:51 localhost.localdomain kubelet[496]: E0226 22:43:51.833792     496 kubelet.go:2183] Container runtime network not ready: Netwo...ialized
Hint: Some lines were ellipsized, use -l to show in full.
```

conf
```
[vagrant@localhost ~]$ sudo cat /lib/systemd/system/kubelet.service.d/10-kubeadm.conf 
# Note: This dropin only works with kubeadm and kubelet v1.11+
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_CONFIG_ARGS=--config=/var/lib/kubelet/config.yaml"
# This is a file that "kubeadm init" and "kubeadm join" generates at runtime, populating the KUBELET_KUBEADM_ARGS variable dynamically
EnvironmentFile=-/var/lib/kubelet/kubeadm-flags.env
# This is a file that the user can use for overrides of the kubelet args as a last resort. Preferably, the user should use
# the .NodeRegistration.KubeletExtraArgs object in the configuration files instead. KUBELET_EXTRA_ARGS should be sourced from this file.
EnvironmentFile=-/etc/sysconfig/kubelet
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_CONFIG_ARGS $KUBELET_KUBEADM_ARGS $KUBELET_EXTRA_ARGS
```

__investigation__

On master
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl get nodes
NAME                    STATUS   ROLES    AGE     VERSION
localhost.localdomain   Ready    <none>   6m27s   v1.17.3
ubuntu-bionic           Ready    master   19d     v1.17.2
```

networking
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get pods -o wide
NAME                                    READY   STATUS    RESTARTS   AGE     IP           NODE                    NOMINATED NODE   READINESS GATES
coredns-6955765f44-8n5f6                1/1     Running   22         19d     10.10.0.74   ubuntu-bionic           <none>           <none>
coredns-6955765f44-jsz85                1/1     Running   22         19d     10.10.0.75   ubuntu-bionic           <none>           <none>
etcd-ubuntu-bionic                      1/1     Running   22         19d     10.0.2.15    ubuntu-bionic           <none>           <none>
kube-apiserver-ubuntu-bionic            1/1     Running   22         19d     10.0.2.15    ubuntu-bionic           <none>           <none>
kube-controller-manager-ubuntu-bionic   1/1     Running   22         19d     10.0.2.15    ubuntu-bionic           <none>           <none>
kube-flannel-ds-amd64-jdmzd             1/1     Running   0          33m     10.0.2.15    ubuntu-bionic           <none>           <none>
kube-flannel-ds-amd64-pl7w9             1/1     Running   0          5m53s   10.0.2.15    localhost.localdomain   <none>           <none>
kube-proxy-kplmt                        1/1     Running   0          5m53s   10.0.2.15    localhost.localdomain   <none>           <none>
kube-proxy-ptzp4                        1/1     Running   22         19d     10.0.2.15    ubuntu-bionic           <none>           <none>
kube-scheduler-ubuntu-bionic            1/1     Running   22         19d     10.0.2.15    ubuntu-bionic           <none>           <none>
kubernetes-dashboard-7c54d59f66-76zq8   1/1     Running   1          2d16h   10.10.0.76   ubuntu-bionic           <none>           <none>
```

trouble
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system logs pods/kube-flannel-ds-amd64-pl7w9
Error from server (NotFound): the server could not find the requested resource ( pods/log kube-flannel-ds-amd64-pl7w9)
```

https://github.com/kubernetes/kubernetes/issues/60835#issuecomment-395583606

https://github.com/kubernetes/kubernetes/issues/60835#issuecomment-395931644

on worker
```
[vagrant@localhost ~]$ sudo sed -i "s/^KUBELET_EXTRA_ARGS=.*/& --node-ip=172.28.128.3/" /etc/sysconfig/kubelet 
```

restart kubelet
```
[vagrant@localhost ~]$ sudo systemctl restart kubelet.service
```

on master
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system logs kube-flannel-ds-amd64-pl7w9
I0226 22:45:21.885295       1 main.go:210] Could not find valid interface matching enp0s8: error looking up interface enp0s8: route ip+net: no such network interface
I0226 22:45:21.885581       1 main.go:527] Using interface with name eth1 and address 172.28.128.3
I0226 22:45:21.885593       1 main.go:544] Defaulting external address to interface address (172.28.128.3)
I0226 22:45:21.989651       1 kube.go:126] Waiting 10m0s for node controller to sync
I0226 22:45:21.989756       1 kube.go:309] Starting kube subnet manager
I0226 22:45:22.990686       1 kube.go:133] Node controller sync successful
I0226 22:45:22.990714       1 main.go:244] Created subnet manager: Kubernetes Subnet Manager - localhost.localdomain
I0226 22:45:22.990719       1 main.go:247] Installing signal handlers
I0226 22:45:22.990842       1 main.go:386] Found network config - Backend type: vxlan
I0226 22:45:22.990883       1 vxlan.go:120] VXLAN config: VNI=1 Port=0 GBP=false DirectRouting=false
I0226 22:45:23.014535       1 main.go:351] Current network or subnet (10.10.0.0/16, 10.10.1.0/24) is not equal to previous one (0.0.0.0/0, 0.0.0.0/0), trying to recycle old iptables rules
I0226 22:45:23.021184       1 iptables.go:167] Deleting iptables rule: -s 0.0.0.0/0 -d 0.0.0.0/0 -j RETURN
I0226 22:45:23.080360       1 iptables.go:167] Deleting iptables rule: -s 0.0.0.0/0 ! -d 224.0.0.0/4 -j MASQUERADE --random-fully
I0226 22:45:23.081180       1 iptables.go:167] Deleting iptables rule: ! -s 0.0.0.0/0 -d 0.0.0.0/0 -j RETURN
I0226 22:45:23.081830       1 iptables.go:167] Deleting iptables rule: ! -s 0.0.0.0/0 -d 0.0.0.0/0 -j MASQUERADE --random-fully
I0226 22:45:23.083487       1 main.go:317] Wrote subnet file to /run/flannel/subnet.env
I0226 22:45:23.083498       1 main.go:321] Running backend.
I0226 22:45:23.083507       1 main.go:339] Waiting for all goroutines to exit
I0226 22:45:23.083525       1 vxlan_network.go:60] watching for new subnet leases
I0226 22:45:23.086113       1 iptables.go:145] Some iptables rules are missing; deleting and recreating rules
I0226 22:45:23.086128       1 iptables.go:167] Deleting iptables rule: -s 10.10.0.0/16 -d 10.10.0.0/16 -j RETURN
I0226 22:45:23.086884       1 iptables.go:167] Deleting iptables rule: -s 10.10.0.0/16 ! -d 224.0.0.0/4 -j MASQUERADE --random-fully
I0226 22:45:23.087584       1 iptables.go:167] Deleting iptables rule: ! -s 10.10.0.0/16 -d 10.10.1.0/24 -j RETURN
I0226 22:45:23.088267       1 iptables.go:167] Deleting iptables rule: ! -s 10.10.0.0/16 -d 10.10.0.0/16 -j MASQUERADE --random-fully
I0226 22:45:23.088971       1 iptables.go:155] Adding iptables rule: -s 10.10.0.0/16 -d 10.10.0.0/16 -j RETURN
I0226 22:45:23.090709       1 iptables.go:155] Adding iptables rule: -s 10.10.0.0/16 ! -d 224.0.0.0/4 -j MASQUERADE --random-fully
I0226 22:45:23.183894       1 iptables.go:145] Some iptables rules are missing; deleting and recreating rules
I0226 22:45:23.183919       1 iptables.go:167] Deleting iptables rule: -s 10.10.0.0/16 -j ACCEPT
I0226 22:45:23.185360       1 iptables.go:167] Deleting iptables rule: -d 10.10.0.0/16 -j ACCEPT
I0226 22:45:23.186340       1 iptables.go:155] Adding iptables rule: -s 10.10.0.0/16 -j ACCEPT
I0226 22:45:23.189299       1 iptables.go:155] Adding iptables rule: -d 10.10.0.0/16 -j ACCEPT
I0226 22:45:23.281704       1 iptables.go:155] Adding iptables rule: ! -s 10.10.0.0/16 -d 10.10.1.0/24 -j RETURN
I0226 22:45:23.283910       1 iptables.go:155] Adding iptables rule: ! -s 10.10.0.0/16 -d 10.10.0.0/16 -j MASQUERADE --random-fully
```

host route
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ ip r
default via 10.0.2.2 dev enp0s3 proto dhcp src 10.0.2.15 metric 100 
10.0.2.0/24 dev enp0s3 proto kernel scope link src 10.0.2.15 
10.0.2.2 dev enp0s3 proto dhcp scope link src 10.0.2.15 metric 100 
10.10.0.0/24 dev cni0 proto kernel scope link src 10.10.0.1 
10.10.1.0/24 via 10.10.1.0 dev flannel.1 onlink 
172.17.0.0/16 dev docker0 proto kernel scope link src 172.17.0.1 linkdown 
172.28.128.0/24 dev enp0s8 proto kernel scope link src 172.28.128.4 
```

arp
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ ip neigh show 10.10.1.0
10.10.1.0 dev flannel.1 lladdr de:80:b1:b5:d3:32 PERMANENT
```

fdb
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ bridge fdb show | grep de:80:b1:b5:d3:32
de:80:b1:b5:d3:32 dev flannel.1 dst 172.28.128.3 self permanent
```

On worker
```
[vagrant@localhost ~]$ docker ps
CONTAINER ID        IMAGE                   COMMAND                  CREATED             STATUS              PORTS               NAMES
6f8fc6a5e3dd        ff281650a721            "/opt/bin/flanneld -…"   8 minutes ago       Up 8 minutes                            k8s_kube-flannel_kube-flannel-ds-amd64-pl7w9_kube-system_7e807ac2-9f10-4544-a8d1-39d09fe0fec4_0
10872a5fcd53        k8s.gcr.io/kube-proxy   "/usr/local/bin/kube…"   9 minutes ago       Up 9 minutes                            k8s_kube-proxy_kube-proxy-kplmt_kube-system_f804210a-fcb1-4f74-939b-4591f7501873_0
9b6deca11c85        k8s.gcr.io/pause:3.1    "/pause"                 12 minutes ago      Up 12 minutes                           k8s_POD_kube-flannel-ds-amd64-pl7w9_kube-system_7e807ac2-9f10-4544-a8d1-39d09fe0fec4_0
958c9aae5431        k8s.gcr.io/pause:3.1    "/pause"                 12 minutes ago      Up 12 minutes                           k8s_POD_kube-proxy-kplmt_kube-system_f804210a-fcb1-4f74-939b-4591f7501873_0
```

logs
```
[vagrant@localhost ~]$ docker logs 6f8fc6a5e3dd
I0226 22:45:21.885295       1 main.go:210] Could not find valid interface matching enp0s8: error looking up interface enp0s8: route ip+net: no such network interface
I0226 22:45:21.885581       1 main.go:527] Using interface with name eth1 and address 172.28.128.3
I0226 22:45:21.885593       1 main.go:544] Defaulting external address to interface address (172.28.128.3)
I0226 22:45:21.989651       1 kube.go:126] Waiting 10m0s for node controller to sync
I0226 22:45:21.989756       1 kube.go:309] Starting kube subnet manager
I0226 22:45:22.990686       1 kube.go:133] Node controller sync successful
I0226 22:45:22.990714       1 main.go:244] Created subnet manager: Kubernetes Subnet Manager - localhost.localdomain
I0226 22:45:22.990719       1 main.go:247] Installing signal handlers
I0226 22:45:22.990842       1 main.go:386] Found network config - Backend type: vxlan
I0226 22:45:22.990883       1 vxlan.go:120] VXLAN config: VNI=1 Port=0 GBP=false DirectRouting=false
I0226 22:45:23.014535       1 main.go:351] Current network or subnet (10.10.0.0/16, 10.10.1.0/24) is not equal to previous one (0.0.0.0/0, 0.0.0.0/0), trying to recycle old iptables rules
I0226 22:45:23.021184       1 iptables.go:167] Deleting iptables rule: -s 0.0.0.0/0 -d 0.0.0.0/0 -j RETURN
I0226 22:45:23.080360       1 iptables.go:167] Deleting iptables rule: -s 0.0.0.0/0 ! -d 224.0.0.0/4 -j MASQUERADE --random-fully
I0226 22:45:23.081180       1 iptables.go:167] Deleting iptables rule: ! -s 0.0.0.0/0 -d 0.0.0.0/0 -j RETURN
I0226 22:45:23.081830       1 iptables.go:167] Deleting iptables rule: ! -s 0.0.0.0/0 -d 0.0.0.0/0 -j MASQUERADE --random-fully
I0226 22:45:23.083487       1 main.go:317] Wrote subnet file to /run/flannel/subnet.env
I0226 22:45:23.083498       1 main.go:321] Running backend.
I0226 22:45:23.083507       1 main.go:339] Waiting for all goroutines to exit
I0226 22:45:23.083525       1 vxlan_network.go:60] watching for new subnet leases
I0226 22:45:23.086113       1 iptables.go:145] Some iptables rules are missing; deleting and recreating rules
I0226 22:45:23.086128       1 iptables.go:167] Deleting iptables rule: -s 10.10.0.0/16 -d 10.10.0.0/16 -j RETURN
I0226 22:45:23.086884       1 iptables.go:167] Deleting iptables rule: -s 10.10.0.0/16 ! -d 224.0.0.0/4 -j MASQUERADE --random-fully
I0226 22:45:23.087584       1 iptables.go:167] Deleting iptables rule: ! -s 10.10.0.0/16 -d 10.10.1.0/24 -j RETURN
I0226 22:45:23.088267       1 iptables.go:167] Deleting iptables rule: ! -s 10.10.0.0/16 -d 10.10.0.0/16 -j MASQUERADE --random-fully
I0226 22:45:23.088971       1 iptables.go:155] Adding iptables rule: -s 10.10.0.0/16 -d 10.10.0.0/16 -j RETURN
I0226 22:45:23.090709       1 iptables.go:155] Adding iptables rule: -s 10.10.0.0/16 ! -d 224.0.0.0/4 -j MASQUERADE --random-fully
I0226 22:45:23.183894       1 iptables.go:145] Some iptables rules are missing; deleting and recreating rules
I0226 22:45:23.183919       1 iptables.go:167] Deleting iptables rule: -s 10.10.0.0/16 -j ACCEPT
I0226 22:45:23.185360       1 iptables.go:167] Deleting iptables rule: -d 10.10.0.0/16 -j ACCEPT
I0226 22:45:23.186340       1 iptables.go:155] Adding iptables rule: -s 10.10.0.0/16 -j ACCEPT
I0226 22:45:23.189299       1 iptables.go:155] Adding iptables rule: -d 10.10.0.0/16 -j ACCEPT
I0226 22:45:23.281704       1 iptables.go:155] Adding iptables rule: ! -s 10.10.0.0/16 -d 10.10.1.0/24 -j RETURN
I0226 22:45:23.283910       1 iptables.go:155] Adding iptables rule: ! -s 10.10.0.0/16 -d 10.10.0.0/16 -j MASQUERADE --random-fully
```

link
```
[vagrant@localhost ~]$ ip -d link show
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00 promiscuity 0 addrgenmode eui64 numtxqueues 1 numrxqueues 1 gso_max_size 65536 gso_max_segs 65535 
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether 52:54:00:fd:32:21 brd ff:ff:ff:ff:ff:ff promiscuity 0 addrgenmode eui64 numtxqueues 1 numrxqueues 1 gso_max_size 65536 gso_max_segs 65535 
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether 08:00:27:7b:dd:51 brd ff:ff:ff:ff:ff:ff promiscuity 0 addrgenmode eui64 numtxqueues 1 numrxqueues 1 gso_max_size 65536 gso_max_segs 65535 
4: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN mode DEFAULT group default 
    link/ether 02:42:33:38:54:af brd ff:ff:ff:ff:ff:ff promiscuity 0 
    bridge forward_delay 1500 hello_time 200 max_age 2000 ageing_time 30000 stp_state 0 priority 32768 vlan_filtering 0 vlan_protocol 802.1Q bridge_id 8000.2:42:33:38:54:af designated_root 8000.2:42:33:38:54:af root_port 0 root_path_cost 0 topology_change 0 topology_change_detected 0 hello_timer    0.00 tcn_timer    0.00 topology_change_timer    0.00 gc_timer  237.12 vlan_default_pvid 1 vlan_stats_enabled 0 group_fwd_mask 0 group_address 01:80:c2:00:00:00 mcast_snooping 1 mcast_router 1 mcast_query_use_ifaddr 0 mcast_querier 0 mcast_hash_elasticity 4 mcast_hash_max 512 mcast_last_member_count 2 mcast_startup_query_count 2 mcast_last_member_interval 100 mcast_membership_interval 26000 mcast_querier_interval 25500 mcast_query_interval 12500 mcast_query_response_interval 1000 mcast_startup_query_interval 3125 mcast_stats_enabled 0 mcast_igmp_version 2 mcast_mld_version 1 nf_call_iptables 0 nf_call_ip6tables 0 nf_call_arptables 0 addrgenmode eui64 numtxqueues 1 numrxqueues 1 gso_max_size 65536 gso_max_segs 65535 
5: flannel.1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UNKNOWN mode DEFAULT group default 
    link/ether de:80:b1:b5:d3:32 brd ff:ff:ff:ff:ff:ff promiscuity 0 
    vxlan id 1 local 172.28.128.3 dev eth1 srcport 0 0 dstport 8472 nolearning ageing 300 noudpcsum noudp6zerocsumtx noudp6zerocsumrx addrgenmode eui64 numtxqueues 1 numrxqueues 1 gso_max_size 65536 gso_max_segs 65535 
```

```
[vagrant@localhost ~]$ ip addr show
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 52:54:00:fd:32:21 brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global noprefixroute dynamic eth0
       valid_lft 77719sec preferred_lft 77719sec
    inet6 fe80::5054:ff:fefd:3221/64 scope link 
       valid_lft forever preferred_lft forever
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:7b:dd:51 brd ff:ff:ff:ff:ff:ff
    inet 172.28.128.3/24 brd 172.28.128.255 scope global noprefixroute dynamic eth1
       valid_lft 776sec preferred_lft 776sec
    inet6 fe80::a00:27ff:fe7b:dd51/64 scope link 
       valid_lft forever preferred_lft forever
4: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default 
    link/ether 02:42:33:38:54:af brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
5: flannel.1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UNKNOWN group default 
    link/ether de:80:b1:b5:d3:32 brd ff:ff:ff:ff:ff:ff
    inet 10.10.1.0/32 scope global flannel.1
       valid_lft forever preferred_lft forever
    inet6 fe80::dc80:b1ff:feb5:d332/64 scope link 
       valid_lft forever preferred_lft forever
```

subnet
```
[vagrant@localhost ~]$ sudo cat /var/run/flannel/subnet.env 
FLANNEL_NETWORK=10.10.0.0/16
FLANNEL_SUBNET=10.10.1.1/24
FLANNEL_MTU=1450
FLANNEL_IPMASQ=true
```

host route
```
[vagrant@localhost ~]$ ip r
default via 10.0.2.2 dev eth0 proto dhcp metric 101 
10.0.2.0/24 dev eth0 proto kernel scope link src 10.0.2.15 metric 101 
10.10.0.0/24 via 10.10.0.0 dev flannel.1 onlink 
172.17.0.0/16 dev docker0 proto kernel scope link src 172.17.0.1 
172.28.128.0/24 dev eth1 proto kernel scope link src 172.28.128.3 metric 100 
```

arp
```
[vagrant@localhost ~]$ sudo ip neigh show 10.10.0.0
10.10.0.0 dev flannel.1 lladdr 76:c4:b3:5a:59:d5 PERMANENT
```

fdb
```
[vagrant@localhost ~]$ sudo bridge fdb show | grep 76:c4:b3:5a:59:d5
76:c4:b3:5a:59:d5 dev flannel.1 dst 172.28.128.4 self permanent
```

### Feb 7 2020

Host
```
vagrant@ubuntu-bionic:~$ sudo hostnamectl
   Static hostname: ubuntu-bionic
         Icon name: computer-vm
           Chassis: vm
        Machine ID: 07010dbe21e349dc9ae0459dcde1e7f7
           Boot ID: 6c1b1f128da04d1e9b4a70141a576264
    Virtualization: oracle
  Operating System: Ubuntu 18.04.1 LTS
            Kernel: Linux 4.15.0-76-generic
      Architecture: x86-64
```

VirtualBox Networking: NAT+Private
```
vagrant@ubuntu-bionic:~$ ip addr show enp0s8 | grep inet | head -1 | awk '{print $2}' | cut -d/ -f1
172.28.128.4
```

Reference:
+ https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/

Google Kubernetes xenial repository
```
vagrant@ubuntu-bionic:~$ cat << EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
> deb https://apt.kubernetes.io/ kubernetes-xenial main
> EOF
deb https://apt.kubernetes.io/ kubernetes-xenial main
```

Kubernetes GPG key for Ubuntu Bionic
```
vagrant@ubuntu-bionic:~$ sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 6A030B21BA07F4FB
Executing: /tmp/apt-key-gpghome.fgUO1XBNnj/gpg.1.sh --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 6A030B21BA07F4FB
gpg: key 6A030B21BA07F4FB: public key "Google Cloud Packages Automatic Signing Key <gc-team@google.com>" imported
gpg: Total number processed: 1
gpg:               imported: 1
```

Kubernetes
```
vagrant@ubuntu-bionic:~$ sudo apt-get install kubeadm kubectl kubelet
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  conntrack cri-tools kubernetes-cni socat
The following NEW packages will be installed:
  conntrack cri-tools kubeadm kubectl kubelet kubernetes-cni socat
0 upgraded, 7 newly installed, 0 to remove and 143 not upgraded.
Need to get 51.6 MB of archives.
After this operation, 272 MB of additional disk space will be used.
Do you want to continue? [Y/n] y
Get:3 http://archive.ubuntu.com/ubuntu bionic/main amd64 conntrack amd64 1:1.4.4+snapshot20161117-6ubuntu2 [30.6 kB]
Get:1 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 cri-tools amd64 1.13.0-00 [8776 kB]      
Get:7 http://archive.ubuntu.com/ubuntu bionic/main amd64 socat amd64 1.7.3.2-2ubuntu2 [342 kB]
Get:2 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubernetes-cni amd64 0.7.5-00 [6473 kB]                                      
Get:4 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubelet amd64 1.17.2-00 [19.2 MB]                                            
Get:5 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubectl amd64 1.17.2-00 [8738 kB]                                            
Get:6 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.17.2-00 [8061 kB]                                            
Fetched 51.6 MB in 1min 22s (628 kB/s)                                                                                                                
Selecting previously unselected package conntrack.
(Reading database ... 139102 files and directories currently installed.)
Preparing to unpack .../0-conntrack_1%3a1.4.4+snapshot20161117-6ubuntu2_amd64.deb ...
Unpacking conntrack (1:1.4.4+snapshot20161117-6ubuntu2) ...
Selecting previously unselected package cri-tools.
Preparing to unpack .../1-cri-tools_1.13.0-00_amd64.deb ...
Unpacking cri-tools (1.13.0-00) ...
Selecting previously unselected package kubernetes-cni.
Preparing to unpack .../2-kubernetes-cni_0.7.5-00_amd64.deb ...
Unpacking kubernetes-cni (0.7.5-00) ...
Selecting previously unselected package socat.
Preparing to unpack .../3-socat_1.7.3.2-2ubuntu2_amd64.deb ...
Unpacking socat (1.7.3.2-2ubuntu2) ...
Selecting previously unselected package kubelet.
Preparing to unpack .../4-kubelet_1.17.2-00_amd64.deb ...
Unpacking kubelet (1.17.2-00) ...
Selecting previously unselected package kubectl.
Preparing to unpack .../5-kubectl_1.17.2-00_amd64.deb ...
Unpacking kubectl (1.17.2-00) ...
Selecting previously unselected package kubeadm.
Preparing to unpack .../6-kubeadm_1.17.2-00_amd64.deb ...
Unpacking kubeadm (1.17.2-00) ...
Setting up conntrack (1:1.4.4+snapshot20161117-6ubuntu2) ...
Setting up kubernetes-cni (0.7.5-00) ...
Setting up cri-tools (1.13.0-00) ...
Setting up socat (1.7.3.2-2ubuntu2) ...
Setting up kubelet (1.17.2-00) ...
Created symlink /etc/systemd/system/multi-user.target.wants/kubelet.service → /lib/systemd/system/kubelet.service.
Setting up kubectl (1.17.2-00) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
Setting up kubeadm (1.17.2-00) ...
```

CNI
```
vagrant@ubuntu-bionic:~$ ls /opt/cni/bin
bridge  dhcp  flannel  host-device  host-local  ipvlan  loopback  macvlan  portmap  ptp  sample  tuning  vlan
```

Service
```
vagrant@ubuntu-bionic:~$ sudo systemctl status kubelet --no-pager
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/lib/systemd/system/kubelet.service; enabled; vendor preset: enabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Thu 2020-02-06 10:19:33 UTC; 37min ago
     Docs: https://kubernetes.io/docs/home/
 Main PID: 26296 (kubelet)
    Tasks: 17 (limit: 4702)
   CGroup: /system.slice/kubelet.service
           └─26296 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --config…nf

Feb 06 10:56:19 ubuntu-bionic kubelet[26296]: E0206 10:56:19.900412   26296 kubelet.go:2183] Container runtime network not ready: NetworkRea…nitialized
Feb 06 10:56:24 ubuntu-bionic kubelet[26296]: W0206 10:56:24.084628   26296 cni.go:237] Unable to update cni config: no networks found in /e…/cni/net.d
Feb 06 10:56:24 ubuntu-bionic kubelet[26296]: E0206 10:56:24.902988   26296 kubelet.go:2183] Container runtime network not ready: NetworkRea…nitialized
Feb 06 10:56:29 ubuntu-bionic kubelet[26296]: W0206 10:56:29.086779   26296 cni.go:237] Unable to update cni config: no networks found in /e…/cni/net.d
Feb 06 10:56:29 ubuntu-bionic kubelet[26296]: E0206 10:56:29.905917   26296 kubelet.go:2183] Container runtime network not ready: NetworkRea…nitialized
Feb 06 10:56:34 ubuntu-bionic kubelet[26296]: W0206 10:56:34.087028   26296 cni.go:237] Unable to update cni config: no networks found in /e…/cni/net.d
Feb 06 10:56:34 ubuntu-bionic kubelet[26296]: E0206 10:56:34.907507   26296 kubelet.go:2183] Container runtime network not ready: NetworkRea…nitialized
Feb 06 10:56:39 ubuntu-bionic kubelet[26296]: W0206 10:56:39.087920   26296 cni.go:237] Unable to update cni config: no networks found in /e…/cni/net.d
Feb 06 10:56:39 ubuntu-bionic kubelet[26296]: E0206 10:56:39.909059   26296 kubelet.go:2183] Container runtime network not ready: NetworkRea…nitialized
Feb 06 10:56:44 ubuntu-bionic kubelet[26296]: W0206 10:56:44.088373   26296 cni.go:237] Unable to update cni config: no networks found in /e…/cni/net.d
Hint: Some lines were ellipsized, use -l to show in full.
```

Bootstrap
```
vagrant@ubuntu-bionic:~$ sudo kubeadm init --control-plane-endpoint=172.28.128.4 --apiserver-advertise-address=172.28.128.4 --pod-network-cidr=10.10.0.0/16 --ignore-preflight-errors=NumCPU --token-ttl=0
W0207 22:01:30.652429    5113 validation.go:28] Cannot validate kube-proxy config - no validator is available
W0207 22:01:30.652505    5113 validation.go:28] Cannot validate kubelet config - no validator is available
[init] Using Kubernetes version: v1.17.2
[preflight] Running pre-flight checks
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [ubuntu-bionic kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 172.28.128.4 172.28.128.4]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [ubuntu-bionic localhost] and IPs [172.28.128.4 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [ubuntu-bionic localhost] and IPs [172.28.128.4 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
W0207 22:01:34.907114    5113 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[control-plane] Creating static Pod manifest for "kube-scheduler"
W0207 22:01:34.908407    5113 manifests.go:214] the default kube-apiserver authorization-mode is "Node,RBAC"; using "Node,RBAC"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 22.502847 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.17" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node ubuntu-bionic as control-plane by adding the label "node-role.kubernetes.io/master=''"
[mark-control-plane] Marking the node ubuntu-bionic as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: p3hpp3.ehka0529vafadvax
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

You can now join any number of control-plane nodes by copying certificate authorities
and service account keys on each node and then running the following as root:

  kubeadm join 172.28.128.4:6443 --token p3hpp3.ehka0529vafadvax \
    --discovery-token-ca-cert-hash sha256:c86e295d23dc1189911b9a14ba19245c4bc33191c0b0de2a9576a45e3c9e2e9e \
    --control-plane 

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 172.28.128.4:6443 --token p3hpp3.ehka0529vafadvax \
    --discovery-token-ca-cert-hash sha256:c86e295d23dc1189911b9a14ba19245c4bc33191c0b0de2a9576a45e3c9e2e9e 
```

kubeconfig
```
vagrant@ubuntu-bionic:~$ mkdir -p $HOME/.kube; sudo cp /etc/kubernetes/admin.conf $HOME/.kube/config && sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

node
```
vagrant@ubuntu-bionic:~$ kubectl get nodes -o wide
NAME            STATUS   ROLES    AGE     VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION      CONTAINER-RUNTIME
ubuntu-bionic   Ready    master   2m30s   v1.17.2   10.0.2.15     <none>        Ubuntu 18.04.1 LTS   4.15.0-76-generic   docker://18.9.7
```

kube-system. Not coredns are waiting cluster networking addons
```
vagrant@ubuntu-bionic:~$ kubectl -n kube-system get pods -o wide
NAME                                    READY   STATUS    RESTARTS   AGE   IP          NODE            NOMINATED NODE   READINESS GATES
coredns-6955765f44-8n5f6                0/1     Pending   0          23m   <none>      <none>          <none>           <none>
coredns-6955765f44-jsz85                0/1     Pending   0          23m   <none>      <none>          <none>           <none>
etcd-ubuntu-bionic                      1/1     Running   0          24m   10.0.2.15   ubuntu-bionic   <none>           <none>
kube-apiserver-ubuntu-bionic            1/1     Running   0          24m   10.0.2.15   ubuntu-bionic   <none>           <none>
kube-controller-manager-ubuntu-bionic   1/1     Running   0          24m   10.0.2.15   ubuntu-bionic   <none>           <none>
kube-proxy-kx5vp                        1/1     Running   0          23m   10.0.2.15   ubuntu-bionic   <none>           <none>
kube-scheduler-ubuntu-bionic            1/1     Running   0          24m   10.0.2.15   ubuntu-bionic   <none>           <none>
```

Cluster DNS
```
vagrant@ubuntu-bionic:~$ kubectl -n kube-system get svc -o wide
NAME       TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE   SELECTOR
kube-dns   ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   27m   k8s-app=kube-dns
```

API
```
vagrant@ubuntu-bionic:~$ kubectl get svc -o wide
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE   SELECTOR
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   27m   <none>
```

curl
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ curl -jkSL https://10.96.0.1/version
{
  "major": "1",
  "minor": "17",
  "gitVersion": "v1.17.2",
  "gitCommit": "59603c6e503c87169aea6106f57b9f242f64df89",
  "gitTreeState": "clean",
  "buildDate": "2020-01-18T23:22:30Z",
  "goVersion": "go1.13.5",
  "compiler": "gc",
  "platform": "linux/amd64"
}
```

version
```
vagrant@ubuntu-bionic:~$ kubectl version
Client Version: version.Info{Major:"1", Minor:"17", GitVersion:"v1.17.2", GitCommit:"59603c6e503c87169aea6106f57b9f242f64df89", GitTreeState:"clean", BuildDate:"2020-01-18T23:30:10Z", GoVersion:"go1.13.5", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"17", GitVersion:"v1.17.2", GitCommit:"59603c6e503c87169aea6106f57b9f242f64df89", GitTreeState:"clean", BuildDate:"2020-01-18T23:22:30Z", GoVersion:"go1.13.5", Compiler:"gc", Platform:"linux/amd64"}
```

## Cluster networking

Reference
+ https://itnext.io/kubernetes-network-deep-dive-7492341e0ab5
+ https://www.objectif-libre.com/en/blog/2018/07/05/k8s-network-solutions-comparison/
+ https://itnext.io/benchmark-results-of-kubernetes-network-plugins-cni-over-10gbit-s-network-36475925a560

### Kernel

ip forward
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ cat /proc/sys/net/ipv4/ip_forward
1
```

net filter
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ cat /proc/sys/net/bridge/bridge-nf-call-iptables
1
```

### Flannel

Reference:
+ https://github.com/coreos/flannel/blob/master/Documentation/configuration.md
+ https://medium.com/@anilkreddyr/kubernetes-with-flannel-understanding-the-networking-part-2-78b53e5364c7

Apply
```
vagrant@ubuntu-bionic:~$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml | sed "s|10.244.0.0/16|10.10.0.0/16|;s/- --kube-subnet-mgr/&\n        - --iface=enp0s8\n        - --iface=eth1/" | kubectl apply -f -
...
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

Or
```
vagrant@ubuntu-bionic:~$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml | sed "s|10.244.0.0/16|10.10.0.0/16|;s/- --kube-subnet-mgr/&\n        - --iface=enp0s8\n        - --iface=eth1/" | tee flannel-10.10.0.0_16-ifindex3.yaml
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl apply -f flannel-10.10.0.0_16-ifindex3.yaml 
```

Verify
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get pods -o wide
NAME                                    READY   STATUS    RESTARTS   AGE     IP          NODE            NOMINATED NODE   READINESS GATES
coredns-6955765f44-8n5f6                1/1     Running   0          6m12s   10.10.0.3   ubuntu-bionic   <none>           <none>
coredns-6955765f44-jsz85                1/1     Running   0          6m13s   10.10.0.2   ubuntu-bionic   <none>           <none>
etcd-ubuntu-bionic                      1/1     Running   0          6m27s   10.0.2.15   ubuntu-bionic   <none>           <none>
kube-apiserver-ubuntu-bionic            1/1     Running   0          6m27s   10.0.2.15   ubuntu-bionic   <none>           <none>
kube-controller-manager-ubuntu-bionic   1/1     Running   0          6m27s   10.0.2.15   ubuntu-bionic   <none>           <none>
kube-flannel-ds-amd64-jdmzd             1/1     Running   0          91s     10.0.2.15   ubuntu-bionic   <none>           <none>
kube-proxy-ptzp4                        1/1     Running   0          6m13s   10.0.2.15   ubuntu-bionic   <none>           <none>
kube-scheduler-ubuntu-bionic            1/1     Running   0          6m26s   10.0.2.15   ubuntu-bionic   <none>           <none>
```

Interface
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ ip addr show flannel.1
13: flannel.1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UNKNOWN group default 
    link/ether ba:1e:d1:2d:24:40 brd ff:ff:ff:ff:ff:ff
    inet 10.10.0.0/32 scope global flannel.1
       valid_lft forever preferred_lft forever
    inet6 fe80::b81e:d1ff:fe2d:2440/64 scope link 
       valid_lft forever preferred_lft forever
```

bridge
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ ip addr show cni0
7: cni0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default qlen 1000
    link/ether 52:cf:d5:59:97:54 brd ff:ff:ff:ff:ff:ff
    inet 10.10.0.1/24 scope global cni0
       valid_lft forever preferred_lft forever
    inet6 fe80::50cf:d5ff:fe59:9754/64 scope link 
       valid_lft forever preferred_lft forever
```

Subnet
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ cat /run/flannel/subnet.env 
FLANNEL_NETWORK=10.10.0.0/16
FLANNEL_SUBNET=10.10.0.1/24
FLANNEL_MTU=1450
FLANNEL_IPMASQ=true
```

HostRoute, Note because the cluster is only one node, flannel.1 has not any routes
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ sudo ip r
default via 10.0.2.2 dev enp0s3 proto dhcp src 10.0.2.15 metric 100 
10.0.2.0/24 dev enp0s3 proto kernel scope link src 10.0.2.15 
10.0.2.2 dev enp0s3 proto dhcp scope link src 10.0.2.15 metric 100 
10.10.0.0/24 dev cni0 proto kernel scope link src 10.10.0.1 
172.17.0.0/16 dev docker0 proto kernel scope link src 172.17.0.1 linkdown 
172.28.128.0/24 dev enp0s8 proto kernel scope link src 172.28.128.4 
172.29.0.0/24 dev br-e65c6afeff2f proto kernel scope link src 172.29.0.1 linkdown 
```

Toplogy
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ brctl show
bridge name	bridge id		STP enabled	interfaces
br-e65c6afeff2f		8000.0242f4fbab6b	no		
cni0		8000.52cfd5599754	no		veth85badab4
							vethe2fd2859
docker0		8000.0242f09a5118	no		
```

__investigation__

Issue on VirtualBox
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get pods/kube-flannel-ds-amd64-jdmzd -o yaml | grep hostIP
  hostIP: 10.0.2.15
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get pods/kube-flannel-ds-amd64-jdmzd -o yaml | grep podIP
  podIP: 10.0.2.15
  podIPs:
```

override
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ echo "KUBELET_EXTRA_ARGS=--node-ip=172.28.128.4" | sudo tee -a /etc/default/kubelet
KUBELET_EXTRA_ARGS=--node-ip=172.28.128.4
```

verify
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get nodes -o wide
NAME                    STATUS   ROLES    AGE     VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE                KERNEL-VERSION               CONTAINER-RUNTIME
localhost.localdomain   Ready    <none>   5h20m   v1.17.3   172.28.128.3   <none>        CentOS Linux 7 (Core)   3.10.0-1062.9.1.el7.x86_64   docker://19.3.6
ubuntu-bionic           Ready    master   19d     v1.17.2   172.28.128.4   <none>        Ubuntu 18.04.1 LTS      4.15.0-88-generic            docker://18.9.7
```

```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get pods/kube-flannel-ds-amd64-jdmzd -o yaml | grep hostIP
  hostIP: 172.28.128.4
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get pods/kube-flannel-ds-amd64-jdmzd -o yaml | grep podIP
  podIP: 172.28.128.4
  podIPs:
```

### Taint

By default, your cluster will not schedule Pods on the control-plane node for security reasons. 
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl taint nodes --all node-role.kubernetes.io/master-
node/ubuntu-bionic untainted
```

### Dashboard

Reference
+ https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md

V1.10.1
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ docker images k8s.gcr.io/kubernetes-dashboard-amd64
REPOSITORY                              TAG                 IMAGE ID            CREATED             SIZE
k8s.gcr.io/kubernetes-dashboard-amd64   v1.10.1             f9aed6605b81        13 months ago       122MB
```

Apply
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ curl -jkSL https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml | kubectl apply -f-
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  4577  100  4577    0     0   8232      0 --:--:-- --:--:-- --:--:--  8232
secret/kubernetes-dashboard-certs created
serviceaccount/kubernetes-dashboard created
role.rbac.authorization.k8s.io/kubernetes-dashboard-minimal created
rolebinding.rbac.authorization.k8s.io/kubernetes-dashboard-minimal created
deployment.apps/kubernetes-dashboard created
service/kubernetes-dashboard created
```

POD
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get pods -l k8s-app=kubernetes-dashboard
NAME                                    READY   STATUS    RESTARTS   AGE
kubernetes-dashboard-7c54d59f66-kxvct   1/1     Running   0          5m58s
```

logs
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system logs -l k8s-app=kubernetes-dashboard -c kubernetes-dashboard
```

Service
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get service -lk8s-ap=kubernetes-dashboard
NAME                   TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
kubernetes-dashboard   ClusterIP   10.99.131.163   <none>        443/TCP   13m
```

curl
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ curl -jkSL https://10.99.131.163
 <!doctype html> <html ng-app="kubernetesDashboard"> <head> <meta charset="utf-8"> <title ng-controller="kdTitle as $ctrl" ng-bind="$ctrl.title()"></title> <link rel="icon" type="image/png" href="assets/images/kubernetes-logo.png"> <meta name="viewport" content="width=device-width"> <link rel="stylesheet" href="static/vendor.93db0a0d.css"> <link rel="stylesheet" href="static/app.ddd3b5ec.css"> </head> <body ng-controller="kdMain as $ctrl"> <!--[if lt IE 10]>
      <p class="browsehappy">You are using an <strong>outdated</strong> browser.
      Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your
      experience.</p>
    <![endif]--> <kd-login layout="column" layout-fill ng-if="$ctrl.isLoginState()"> </kd-login> <kd-chrome layout="column" layout-fill ng-if="!$ctrl.isLoginState()"> </kd-chrome> <script src="static/vendor.bd425c26.js"></script> <script src="api/appConfig.json"></script> <script src="static/app.91a96542.js"></script> </body> </html>
```

Expose
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system expose deployment kubernetes-dashboard --name=kubernetes-dashboard-expose --target-port=8443 --type=NodePort 
service/kubernetes-dashboard-expose exposed
```

NodePort
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system get service kubernetes-dashboard-expose
NAME                          TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kubernetes-dashboard-expose   NodePort   10.98.116.198   <none>        8443:32267/TCP   20s
```

Then open browser, navigate https://172.28.128.4:32267/ , note 172.28.128.4 is this VM address

User
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl apply -f - <<EOF
> apiVersion: v1
> kind: ServiceAccount
> metadata:
>   name: admin-user
>   namespace: kube-system
> ---
> apiVersion: rbac.authorization.k8s.io/v1
> kind: ClusterRoleBinding
> metadata:
>   name: admin-user
> roleRef:
>   apiGroup: rbac.authorization.k8s.io
>   kind: ClusterRole
>   name: cluster-admin
> subjects:
> - kind: ServiceAccount
>   name: admin-user
>   namespace: kube-system
> EOF
serviceaccount/admin-user created
clusterrolebinding.rbac.authorization.k8s.io/admin-user created
```

Token
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')
Name:         admin-user-token-vcqrf
Namespace:    kube-system
Labels:       <none>
Annotations:  kubernetes.io/service-account.name: admin-user
              kubernetes.io/service-account.uid: 17bd7a5b-8a39-4313-98b0-fc8f53158df6

Type:  kubernetes.io/service-account-token

Data
====
ca.crt:     1025 bytes
namespace:  11 bytes
token:      eyJhbGciOiJSUzI1NiIsImtpZCI6ImY2YmphX0MtX25OcjdQMndWaHhVaVBLdGFZS05NenQxWkhud3dFUDVjZ0EifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi11c2VyLXRva2VuLXZjcXJmIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFkbWluLXVzZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiIxN2JkN2E1Yi04YTM5LTQzMTMtOThiMC1mYzhmNTMxNThkZjYiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06YWRtaW4tdXNlciJ9.xKR3XJWAubcYHBoV4H1btNR3PD9NxLGp4fflWY_-yG7H-UFqhZLJHHlNuxs_S5Jl-zI7wrZkqrsuJBNOU2WIHgXLDjSG8g8qVSw7dnqCbGb_FB4uCeVMrI_Bd6FRw3nrLIyhFItEK2VWh9EyLoAHGrc635f_mWpbvfgDvy9IuoqhkybbWvIpQBec2zZlsRfkuSkaRJHn2RglLYScixUQE4ASnvmQ9G-x6iWruWyseq-3MdumcfZuvAZyPrCNbt3Yy4khKiJqYF3augG5fSdF9wdkfGo35pVE7Pl-u2DyMIxsCVVJfyFPdpdpmM2mOIAt2k4H-TE7AnMBtFIQdmXHtA
```

Then login dashboard with upon token.

PortForward is another dev option
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl -n kube-system port-forward --address=172.28.128.4 deployment/kubernetes-dashboard 8443
Forwarding from 172.28.128.4:8443 -> 8443
^C
```

Proxy is available if it is not a VM instance
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl proxy --address=172.28.128.4
Starting to serve on 172.28.128.4:8001
^C
```

## 26 Nov 2017

### V0.21.2 Docker-based Kubernetes cluster

kubernetes v1.0 documentation - Getting Started - Running locally via docker
* http://kubernetes.io/v1.0/docs/getting-started-guides/docker.html

Run etcd for cluster configuration management
```
# docker run --net=host -d gcr.io/google_containers/etcd:2.0.9 /usr/local/bin/etcd --addr=127.0.0.1:4001 --bind-addr=0.0.0.0:4001 --data-dir=/var/etcd/data
```

Run master. The master include controll, scheduler, POD, api... so enable TCP 8080 port or API will be failed
```
# docker run --net=host -d -v /var/run/docker.sock:/var/run/docker.sock  gcr.io/google_containers/hyperkube:v0.21.2 /hyperkube kubelet --api_servers=http://localhost:8080 --v=2 --address=0.0.0.0 --enable_server --hostname_override=127.0.0.1 --config=/etc/kubernetes/manifests
```

Run service proxy. It is used to expose service with endpoint for host networking
```
# docker run -d --net=host --privileged gcr.io/google_containers/hyperkube:v0.21.2 /hyperkube proxy --master=http://127.0.0.1:8080 --v=2
```

Run a nginx demo. Run nginx with replication controller
```
# kubectl -s http://localhost:8080 run nginx --image=nginx --port=80
```

Expose nginx as networing service. Enable TCP 80 port or service can not be created
```
# kubectl expose rc nginx --port=80
```

### Test it out

Via endpoint
```
# /opt/tfx/kubectl get endpoints
NAME         ENDPOINTS
kubernetes   192.168.0.25:6443
nginx        172.17.0.1:80
# curl http://172.17.0.1
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>
<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>
<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

Via networking service
```
# /opt/tfx/kubectl get services
NAME         LABELS                                    SELECTOR    IP(S)        PORT(S)
kubernetes   component=apiserver,provider=kubernetes   <none>      10.0.0.1     443/TCP
nginx        run=nginx                                 run=nginx   10.0.0.206   80/TCP
```

Or
```
# curl http://10.0.0.206
```

### kube api

curl
```
$ curl http://127.0.0.1:8080/api/v1/namespaces
{
  "kind": "NamespaceList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/namespaces",
    "resourceVersion": "39985"
  },
  "items": [
    {
      "metadata": {
        "name": "default",
        "selfLink": "/api/v1/namespaces/default",
        "uid": "a9117908-4b3d-11e6-a5b9-0800274654e7",
        "resourceVersion": "6",
        "creationTimestamp": "2016-07-16T10:11:28Z"
      },
      "spec": {
        "finalizers": [
          "kubernetes"
        ]
      },
      "status": {
        "phase": "Active"
      }
    },
    {
      "metadata": {
        "name": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system",
        "uid": "a91fb003-4b3d-11e6-a5b9-0800274654e7",
        "resourceVersion": "9",
        "creationTimestamp": "2016-07-16T10:11:28Z"
      },
      "spec": {
        "finalizers": [
          "kubernetes"
        ]
      },
      "status": {
        "phase": "Active"
      }
    }
  ]
}
```

### list containers with Docker

run Docker command
```
# docker ps
CONTAINER ID        IMAGE                                        COMMAND                CREATED             STATUS              PORTS               NAMES
58ca7b7bc89a        nginx                                        "nginx -g 'daemon of   59 minutes ago      Up 59 minutes                           k8s_nginx.d7d3eb2f_nginx-syeg7_default_fc945b19-41a1-11e5-b8c4-c4346b46de6e_24ac6b02
ffed942cfc52        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube controll   59 minutes ago      Up 59 minutes                           k8s_controller-manager.aad1ee8f_k8s-master-127.0.0.1_default_9b44830745c166dfc6d027b0fc2df36d_8441f980
ac7cbe846daf        gcr.io/google_containers/pause:0.8.0         "/pause"               59 minutes ago      Up 59 minutes                           k8s_POD.ef28e851_nginx-syeg7_default_fc945b19-41a1-11e5-b8c4-c4346b46de6e_eeac80e6
0af559d7e351        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube schedule   59 minutes ago      Up 59 minutes                           k8s_scheduler.b725e775_k8s-master-127.0.0.1_default_9b44830745c166dfc6d027b0fc2df36d_8d259834
fa8bac1ebcfb        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube apiserve   59 minutes ago      Up 59 minutes                           k8s_apiserver.70750283_k8s-master-127.0.0.1_default_9b44830745c166dfc6d027b0fc2df36d_6c3cd54a
608b86a418dd        gcr.io/google_containers/pause:0.8.0         "/pause"               59 minutes ago      Up 59 minutes                           k8s_POD.e4cc795_k8s-master-127.0.0.1_default_9b44830745c166dfc6d027b0fc2df36d_4c85ae5c
706a1dd60111        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube proxy --   10 days ago         Up 59 minutes                           sharp_bell
8f4172000232        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube kubelet    10 days ago         Up About an hour                        suspicious_albattani
cfb63e647cac        gcr.io/google_containers/etcd:2.0.9          "/usr/local/bin/etcd   10 days ago         Up About an hour                        gloomy_jones
```

