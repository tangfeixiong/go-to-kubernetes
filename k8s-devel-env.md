# Kubernetes Devel Env

## V1.17.2

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

### Feb 7 2020

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
+ https://medium.com/@anilkreddyr/kubernetes-with-flannel-understanding-the-networking-part-2-78b53e5364c7

Apply
```
vagrant@ubuntu-bionic:~$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml | sed "s|10.244.0.0/16|10.10.0.0/16|;s/- --kube-subnet-mgr/&\n        - --iface=enp0s8/" | kubectl apply -f -
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
vagrant@ubuntu-bionic:~$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml | sed "s|10.244.0.0/16|10.10.0.0/16|;s/- --kube-subnet-mgr/&\n        - --iface=enp0s8/" | tee flannel-10.10.0.0_16-enp0s8.yaml
```

Then
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes$ kubectl apply -f flannel.yaml 
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
kube-flannel-ds-amd64-t7w6q             1/1     Running   0          91s     10.0.2.15   ubuntu-bionic   <none>           <none>
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
FLANNEL_NETWORK=10.10.0.0/24
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

