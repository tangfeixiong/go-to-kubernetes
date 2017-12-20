# DevOps

__Issue__
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo kubeadm init --apiserver-advertise-address 10.64.33.59
[init] Using Kubernetes version: v1.9.0
[init] Using Authorization modes: [Node RBAC]
[preflight] Running pre-flight checks.
	[WARNING Service-Kubelet]: kubelet service is not enabled, please run 'systemctl enable kubelet.service'
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
[preflight] Some fatal errors occurred:
	[ERROR FileContent--proc-sys-net-bridge-bridge-nf-call-iptables]: /proc/sys/net/bridge/bridge-nf-call-iptables contents are not set to 1
[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
```

Start kubelet.service first
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo systemctl is-active kubelet.service
inactive
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo systemctl start kubelet.service && sudo systemctl enable kubelet.service
Created symlink /etc/systemd/system/multi-user.target.wants/kubelet.service → /etc/systemd/system/kubelet.service.
```

modify kernel param
```
[vagrant@kubedev-172-17-4-59 ~]$ grep -R "bridge-nf-call-iptables" /usr/lib/sysctl.d/
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-iptables = 0
```

```
[vagrant@kubedev-172-17-4-59 ~]$ grep -R "bridge-nf-call-iptables" /etc/sysctl.conf /etc/sysctl.d/
```

```
[vagrant@kubedev-172-17-4-59 ~]$ echo "net.bridge.bridge-nf-call-iptables=1" | sudo tee -a /etc/sysctl.conf 
net.bridge.bridge-nf-call-iptables=1
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo sysctl -w net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-iptables = 1
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo sysctl net.bridge
net.bridge.bridge-nf-call-arptables = 0
net.bridge.bridge-nf-call-ip6tables = 0
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-filter-pppoe-tagged = 0
net.bridge.bridge-nf-filter-vlan-tagged = 0
net.bridge.bridge-nf-pass-vlan-input-dev = 0
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo sysctl -w net.bridge.bridge-nf-call-ip6tables=1
net.bridge.bridge-nf-call-ip6tables = 1
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo sysctl -w net.bridge.bridge-nf-call-arptables=1
net.bridge.bridge-nf-call-arptables = 1
```

```
[vagrant@kubedev-172-17-4-59 ~]$ echo -e "net.bridge.bridge-nf-call-ip6tables=1\nnet.bridge.bridge-nf-call-arptables=1" | sudo tee -a /etc/sysctl.conf 
net.bridge.bridge-nf-call-ip6tables=1
net.bridge.bridge-nf-call-arptables=1
```

__Issue__
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo kubeadm init --apiserver-advertise-address 10.64.33.59
[init] Using Kubernetes version: v1.9.0
[init] Using Authorization modes: [Node RBAC]
[preflight] Running pre-flight checks.
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [kubedev-172-17-4-59 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.64.33.59]
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
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests".
[init] This might take a minute or longer if the control plane images have to be pulled.
^C
```

invalide api address
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo kubeadm reset
[preflight] Running pre-flight checks.
[reset] Stopping the kubelet service.
[reset] Unmounting mounted directories in "/var/lib/kubelet"
[reset] Removing kubernetes-managed containers.
[reset] Deleting contents of stateful directories: [/var/lib/kubelet /etc/cni/net.d /var/lib/dockershim /var/run/kubernetes /var/lib/etcd]
[reset] Deleting contents of config directories: [/etc/kubernetes/manifests /etc/kubernetes/pki]
[reset] Deleting files: [/etc/kubernetes/admin.conf /etc/kubernetes/kubelet.conf /etc/kubernetes/controller-manager.conf /etc/kubernetes/scheduler.conf]
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo kubeadm init --apiserver-advertise-address 172.17.4.59
[init] Using Kubernetes version: v1.9.0
[init] Using Authorization modes: [Node RBAC]
[preflight] Running pre-flight checks.
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [kubedev-172-17-4-59 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 172.17.4.59]
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
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests".
[init] This might take a minute or longer if the control plane images have to be pulled.
[apiclient] All control plane components are healthy after 25.503393 seconds
[uploadconfig] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[markmaster] Will mark node kubedev-172-17-4-59 as master by adding a label and a taint
[markmaster] Master kubedev-172-17-4-59 tainted and labelled with key/value: node-role.kubernetes.io/master=""
[bootstraptoken] Using token: c03eb1.a16eca29b64a1920
[bootstraptoken] Configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstraptoken] Configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstraptoken] Configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstraptoken] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[addons] Applied essential addon: kube-dns
[addons] Applied essential addon: kube-proxy

Your Kubernetes master has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

You can now join any number of machines by running the following on each node
as root:

  kubeadm join --token c03eb1.a16eca29b64a1920 172.17.4.59:6443 --discovery-token-ca-cert-hash sha256:b8c66a12251148ab7ad1c73a685668c1884874db830fb460c4f31b971e19ddc2

```

```
[vagrant@kubedev-172-17-4-59 ~]$ mkdir -p $HOME/.kube && sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config && sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

Issue
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get nodes
NAME                  STATUS     ROLES     AGE       VERSION
kubedev-172-17-4-59   NotReady   master    2m        v1.9.0
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get -n kube-system pods -o wide
NAME                                          READY     STATUS    RESTARTS   AGE       IP            NODE
etcd-kubedev-172-17-4-59                      1/1       Running   0          2m        172.17.4.59   kubedev-172-17-4-59
kube-apiserver-kubedev-172-17-4-59            1/1       Running   0          1m        172.17.4.59   kubedev-172-17-4-59
kube-controller-manager-kubedev-172-17-4-59   1/1       Running   0          2m        172.17.4.59   kubedev-172-17-4-59
kube-dns-6f4fd4bdf-swk2q                      0/3       Pending   0          2m        <none>        <none>
kube-proxy-46jgw                              1/1       Running   0          2m        172.17.4.59   kubedev-172-17-4-59
kube-scheduler-kubedev-172-17-4-59            1/1       Running   0          2m        172.17.4.59   kubedev-172-17-4-59
```

Because cni leak
```
[vagrant@kubedev-172-17-4-59 ~]$ ip a show cni0
Device "cni0" does not exist.
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/docs/coreos0x2Fflannel/kube-flannel_0.9.1.yaml 
clusterrole "flannel" created
clusterrolebinding "flannel" created
serviceaccount "flannel" created
configmap "kube-flannel-cfg" created
daemonset "kube-flannel-ds" created
```

__Issue__
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get -n kube-system pods -o wide
NAME                                          READY     STATUS              RESTARTS   AGE       IP            NODE
etcd-kubedev-172-17-4-59                      1/1       Running             0          5m        172.17.4.59   kubedev-172-17-4-59
kube-apiserver-kubedev-172-17-4-59            1/1       Running             0          5m        172.17.4.59   kubedev-172-17-4-59
kube-controller-manager-kubedev-172-17-4-59   1/1       Running             0          5m        172.17.4.59   kubedev-172-17-4-59
kube-dns-6f4fd4bdf-swk2q                      0/3       ContainerCreating   0          6m        <none>        kubedev-172-17-4-59
kube-flannel-ds-p96xd                         0/1       Error               3          1m        172.17.4.59   kubedev-172-17-4-59
kube-proxy-46jgw                              1/1       Running             0          6m        172.17.4.59   kubedev-172-17-4-59
kube-scheduler-kubedev-172-17-4-59            1/1       Running             0          5m        172.17.4.59   kubedev-172-17-4-59
```

Because interface bind incorrect
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl -n kube-system logs kube-flannel-ds-p96xd
I1218 07:15:47.139408       1 main.go:474] Determining IP address of default interface
I1218 07:15:47.140075       1 main.go:487] Using interface with name eth0 and address 10.0.2.15
I1218 07:15:47.140111       1 main.go:504] Defaulting external address to interface address (10.0.2.15)
I1218 07:15:47.151150       1 kube.go:130] Waiting 10m0s for node controller to sync
I1218 07:15:47.151240       1 kube.go:283] Starting kube subnet manager
I1218 07:15:48.151507       1 kube.go:137] Node controller sync successful
I1218 07:15:48.151630       1 main.go:234] Created subnet manager: Kubernetes Subnet Manager - kubedev-172-17-4-59
I1218 07:15:48.151657       1 main.go:237] Installing signal handlers
I1218 07:15:48.151731       1 main.go:352] Found network config - Backend type: vxlan
I1218 07:15:48.151801       1 vxlan.go:119] VXLAN config: VNI=1 Port=0 GBP=false DirectRouting=false
E1218 07:15:48.152455       1 main.go:279] Error registering network: failed to acquire lease: node "kubedev-172-17-4-59" pod cidr not assigned
I1218 07:15:48.152576       1 main.go:332] Stopping shutdownHandler...
```

```
[vagrant@kubedev-172-17-4-59 ~]$ cp /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/docs/coreos0x2Fflannel/kube-flannel_0.9.1.yaml kube-flannel.yaml
```

modify YAML like
```
ubuntu@kubedev-10-64-33-195:~$ egrep 'command: \[' kube-flannel.yaml 
        command: [ "/opt/bin/flanneld", "--ip-masq", "--kube-subnet-mgr", "--iface=eth1", "--iface=enp0s8", "--iface-regex=10\\.64\\.33\\.[0-9]+" ]
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl delete -f kube-flannel.yaml 
clusterrole "flannel" deleted
clusterrolebinding "flannel" deleted
serviceaccount "flannel" deleted
configmap "kube-flannel-cfg" deleted
daemonset "kube-flannel-ds" deleted
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f kube-flannel.yaml 
clusterrole "flannel" created
clusterrolebinding "flannel" created
serviceaccount "flannel" created
configmap "kube-flannel-cfg" created
daemonset "kube-flannel-ds" created
```

__Issue__
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get -n kube-system pods -o wide
NAME                                          READY     STATUS              RESTARTS   AGE       IP            NODE
etcd-kubedev-172-17-4-59                      1/1       Running             0          16m       172.17.4.59   kubedev-172-17-4-59
kube-apiserver-kubedev-172-17-4-59            1/1       Running             0          16m       172.17.4.59   kubedev-172-17-4-59
kube-controller-manager-kubedev-172-17-4-59   1/1       Running             0          16m       172.17.4.59   kubedev-172-17-4-59
kube-dns-6f4fd4bdf-swk2q                      0/3       ContainerCreating   0          17m       <none>        kubedev-172-17-4-59
kube-flannel-ds-v68m8                         0/1       Error               2          35s       172.17.4.59   kubedev-172-17-4-59
kube-proxy-46jgw                              1/1       Running             0          17m       172.17.4.59   kubedev-172-17-4-59
kube-scheduler-kubedev-172-17-4-59            1/1       Running             0          16m       172.17.4.59   kubedev-172-17-4-59
```

POD cidr not identical, 
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl -n kube-system logs kube-flannel-ds-v68m8
I1218 07:26:54.144562       1 main.go:487] Using interface with name eth1 and address 172.17.4.59
I1218 07:26:54.147498       1 main.go:504] Defaulting external address to interface address (172.17.4.59)
I1218 07:26:54.245579       1 kube.go:130] Waiting 10m0s for node controller to sync
I1218 07:26:54.245677       1 kube.go:283] Starting kube subnet manager
I1218 07:26:55.246049       1 kube.go:137] Node controller sync successful
I1218 07:26:55.246079       1 main.go:234] Created subnet manager: Kubernetes Subnet Manager - kubedev-172-17-4-59
I1218 07:26:55.246085       1 main.go:237] Installing signal handlers
I1218 07:26:55.246285       1 main.go:352] Found network config - Backend type: vxlan
I1218 07:26:55.246328       1 vxlan.go:119] VXLAN config: VNI=1 Port=0 GBP=false DirectRouting=false
E1218 07:26:55.246765       1 main.go:279] Error registering network: failed to acquire lease: node "kubedev-172-17-4-59" pod cidr not assigned
I1218 07:26:55.246835       1 main.go:332] Stopping shutdownHandler...
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo kubeadm config view > standalone.conf
```

modify like
```
[vagrant@kubedev-172-17-4-59 ~]$ cat standalone.conf | grep "podSubnet"
  podSubnet: "10.244.0.0/16"
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo kubeadm reset       
[preflight] Running pre-flight checks.
[reset] Stopping the kubelet service.
[reset] Unmounting mounted directories in "/var/lib/kubelet"
[reset] Removing kubernetes-managed containers.
[reset] Deleting contents of stateful directories: [/var/lib/kubelet /etc/cni/net.d /var/lib/dockershim /var/run/kubernetes /var/lib/etcd]
[reset] Deleting contents of config directories: [/etc/kubernetes/manifests /etc/kubernetes/pki]
[reset] Deleting files: [/etc/kubernetes/admin.conf /etc/kubernetes/kubelet.conf /etc/kubernetes/controller-manager.conf /etc/kubernetes/scheduler.conf]
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo systemctl restart kubelet.service
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo kubeadm init --apiserver-advertise-address 172.17.4.59 --pod-network-cidr=10.244.0.0/16
[init] Using Kubernetes version: v1.9.0
[init] Using Authorization modes: [Node RBAC]
[preflight] Running pre-flight checks.
	[WARNING FileExisting-tc]: tc not found in system path
	[WARNING FileExisting-crictl]: crictl not found in system path
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [kubedev-172-17-4-59 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 172.17.4.59]
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
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests".
[init] This might take a minute or longer if the control plane images have to be pulled.
[apiclient] All control plane components are healthy after 24.504559 seconds
[uploadconfig] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[markmaster] Will mark node kubedev-172-17-4-59 as master by adding a label and a taint
[markmaster] Master kubedev-172-17-4-59 tainted and labelled with key/value: node-role.kubernetes.io/master=""
[bootstraptoken] Using token: 87bdd5.7aa47e764446c926
[bootstraptoken] Configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstraptoken] Configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstraptoken] Configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstraptoken] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[addons] Applied essential addon: kube-dns
[addons] Applied essential addon: kube-proxy

Your Kubernetes master has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

You can now join any number of machines by running the following on each node
as root:

  kubeadm join --token 87bdd5.7aa47e764446c926 172.17.4.59:6443 --discovery-token-ca-cert-hash sha256:6b49ecf4f8c3b1eff1a34be523541454d6718b5f2d2022cad963d28e4398aefa

```

__Issue__
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f kube-flannel.yaml 
Unable to connect to the server: x509: certificate signed by unknown authority (possibly because of "crypto/rsa: verification error" while trying to verify candidate authority certificate "kubernetes")
```

CA crt changed
```
[vagrant@kubedev-172-17-4-59 ~]$ mkdir -p $HOME/.kube && sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config && sudo chown $(id -u):$(id -g) $HOME/.kube/config
cp: overwrite '/home/vagrant/.kube/config'? yes
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f kube-flannel.yaml 
clusterrole "flannel" created
clusterrolebinding "flannel" created
serviceaccount "flannel" created
configmap "kube-flannel-cfg" created
daemonset "kube-flannel-ds" created
```

__Running__
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get -n kube-system pods -o wide
NAME                                          READY     STATUS    RESTARTS   AGE       IP            NODE
etcd-kubedev-172-17-4-59                      1/1       Running   0          5m        172.17.4.59   kubedev-172-17-4-59
kube-apiserver-kubedev-172-17-4-59            1/1       Running   0          5m        172.17.4.59   kubedev-172-17-4-59
kube-controller-manager-kubedev-172-17-4-59   1/1       Running   0          5m        172.17.4.59   kubedev-172-17-4-59
kube-dns-6f4fd4bdf-r8c7h                      3/3       Running   0          6m        10.244.0.2    kubedev-172-17-4-59
kube-flannel-ds-vptkd                         1/1       Running   0          31s       172.17.4.59   kubedev-172-17-4-59
kube-proxy-xj2vb                              1/1       Running   0          6m        172.17.4.59   kubedev-172-17-4-59
kube-scheduler-kubedev-172-17-4-59            1/1       Running   0          5m        172.17.4.59   kubedev-172-17-4-59
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get -n kube-system nodes       
NAME                  STATUS    ROLES     AGE       VERSION
kubedev-172-17-4-59   Ready     master    7m        v1.9.0
```

### Test

a simple hello http
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl run hello --image=docker.io/tangfeixiong/netcat-hello-http          
deployment "hello" created
```

__Issue__
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get pods
NAME                     READY     STATUS    RESTARTS   AGE
hello-778f75f8cc-7mncg   0/1       Pending   0          44s
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl describe po/hello-778f75f8cc-7mncg
Name:           hello-778f75f8cc-7mncg
Namespace:      default
Node:           <none>
Labels:         pod-template-hash=3349319477
                run=hello
Annotations:    <none>
Status:         Pending
IP:             
Controlled By:  ReplicaSet/hello-778f75f8cc
Containers:
  hello:
    Image:        docker.io/tangfeixiong/netcat-hello-http
    Port:         <none>
    Environment:  <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-67l4p (ro)
Conditions:
  Type           Status
  PodScheduled   False 
Volumes:
  default-token-67l4p:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-67l4p
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type     Reason            Age                From               Message
  ----     ------            ----               ----               -------
  Warning  FailedScheduling  26s (x7 over 57s)  default-scheduler  0/1 nodes are available: 1 PodToleratesNodeTaints.
```

Because master node only
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get deploy 
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
hello     1         1         1            0           3m
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl edit deploy/hello
deployment "hello" edited
```

Edit like
```
  template:
    metadata:
      creationTimestamp: null
      labels:
        run: hello
    spec:
      containers:
      - image: docker.io/tangfeixiong/netcat-hello-http
        imagePullPolicy: Always
        name: hello
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      tolerations:
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get pods -o wide
NAME                    READY     STATUS    RESTARTS   AGE       IP           NODE
hello-6cb8784cb-2gzbc   1/1       Running   0          2m        10.244.0.3   kubedev-172-17-4-59
```

Service
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl expose deployment/hello --port=80
service "hello" exposed
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get service -l run=hello
NAME      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
hello     ClusterIP   10.104.114.17   <none>        80/TCP    2m
```

test
```
[vagrant@kubedev-172-17-4-59 ~]$ curl 10.244.0.3
<html><head><title>welcome</title></head><body><h1>hello world</h1></body></html>
```

```
[vagrant@kubedev-172-17-4-59 ~]$ curl 10.104.114.17
<html><head><title>welcome</title></head><body><h1>hello world</h1></body></html>
```