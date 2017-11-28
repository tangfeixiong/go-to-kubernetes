# instruction

Using previous deployed nodes: [centos 7.2](./kubeadm-deploy-k8s-into-centos7.2.md), [ubuntu xenial](./kubeadm-deploy-k8s-into-ubuntu16.04.md), [fedora 23](./kubeadm-deploy-k8s-into-fedora23.md)

__notice__  may not work in CentOS 7.2 with docker engine v1.12

![]()

## Networking

CentOS 7.2
```
[vagrant@kubedev-10-64-33-82 ~]$ ip a show enp0s8; ip a show enp0s9
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 08:00:27:33:cc:33 brd ff:ff:ff:ff:ff:ff
    inet 10.64.33.82/24 brd 10.64.33.255 scope global enp0s8
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe33:cc33/64 scope link 
       valid_lft forever preferred_lft forever
4: enp0s9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 08:00:27:48:b2:a2 brd ff:ff:ff:ff:ff:ff
    inet 172.28.128.4/24 brd 172.28.128.255 scope global dynamic enp0s9
       valid_lft 836sec preferred_lft 836sec
    inet6 fe80::a00:27ff:fe48:b2a2/64 scope link 
       valid_lft forever preferred_lft forever
```

Fedora 23
```
[vagrant@kubedev-10-64-33-199 ~]$ ip a show eth1; ip a show eth2
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:3f:88:a3 brd ff:ff:ff:ff:ff:ff
    inet 10.64.33.199/24 brd 10.64.33.255 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe3f:88a3/64 scope link 
       valid_lft forever preferred_lft forever
4: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:27:95:e9 brd ff:ff:ff:ff:ff:ff
    inet 172.28.128.3/24 brd 172.28.128.255 scope global dynamic eth2
       valid_lft 991sec preferred_lft 991sec
    inet6 fe80::a00:27ff:fe27:95e9/64 scope link 
       valid_lft forever preferred_lft forever
```

Ubuntu 16.04
```
ubuntu@kubedev-10-64-33-195:~$ ip a show enp0s8; ip a show enp0s9
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:c4:54:fb brd ff:ff:ff:ff:ff:ff
    inet 10.64.33.195/24 brd 10.64.33.255 scope global enp0s8
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fec4:54fb/64 scope link 
       valid_lft forever preferred_lft forever
4: enp0s9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:e7:a0:16 brd ff:ff:ff:ff:ff:ff
    inet 172.28.128.6/24 brd 172.28.128.255 scope global enp0s9
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fee7:a016/64 scope link 
       valid_lft forever preferred_lft forever
```

### VIP

Configure 172.28.128.50 as internal VIP into CentOS 7.2 for private network
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo ip addr add 10.64.33.50/24 dev enp0s8
```

Optinally, configure 172.28.128.50 as external VIP into CentOS 7.2 if designed for public network
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo ip addr add 172.28.128.50/24 dev enp0s9
```

## Clean previous standalone mode

Export ConfigMap from any single master
```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl get configmap/kubeadm-config --output=yaml --namespace=kube-system
apiVersion: v1
data:
  MasterConfiguration: |
    api:
      advertiseAddress: 10.64.33.82
      bindPort: 443
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
    nodeName: kubedev-10-64-33-82
    token: ""
    tokenTTL: 24h0m0s
    unifiedControlPlaneImage: ""
kind: ConfigMap
metadata:
  creationTimestamp: 2017-11-22T06:38:59Z
  name: kubeadm-config
  namespace: kube-system
  resourceVersion: "24835"
  selfLink: /api/v1/namespaces/kube-system/configmaps/kubeadm-config
  uid: d270f3d4-cf4f-11e7-ab20-080027539284
```

### Stop kubelet and etcd data, clean configuration

__Stop kubelet service__

CentOS 7
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo systemctl stop kubelet.service
```

Fedora 23
```
[vagrant@kubedev-10-64-33-199 ~]$ sudo systemctl stop kubelet.service
```

Ubuntu 16.04
```
ubuntu@kubedev-10-64-33-195:~$ sudo systemctl stop kubelet.service
```

__Delete configuration but etcd.yaml__

CentOS 7
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo mv /etc/kubernetes/manifests/etcd.yaml ./etcd-single.yaml && sudo rm -rf /etc/kubernetes/manifests/; sudo rm -rf /etc/kubernetes/pki/; sudo rm -rf /etc/kubernetes/*.conf
```

Fedora 23
```
[vagrant@kubedev-10-64-33-199 ~]$ sudo mv /etc/kubernetes/manifests/etcd.yaml ./etcd-single.yaml && sudo rm -rf /etc/kubernetes/manifests/; sudo rm -rf /etc/kubernetes/pki/; sudo rm -rf /etc/kubernetes/*.conf
```

Ubuntu 16.04
```
ubuntu@kubedev-10-64-33-195:~$ sudo mv /etc/kubernetes/manifests/etcd.yaml ./etcd-single.yaml && sudo rm -rf /etc/kubernetes/manifests/; sudo rm -rf /etc/kubernetes/pki/; sudo rm -rf /etc/kubernetes/*.conf
```

__Delete etcd pervious data__

CentOS
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo rm -rf /var/lib/etcd/member/
```

Fedora
```
[vagrant@kubedev-10-64-33-199 ~]$ sudo rm -rf /var/lib/etcd/member/
```

Ubuntu
```
ubuntu@kubedev-10-64-33-195:~$ sudo rm -rf /var/lib/etcd/member/
```

## Create etcd cluster

### external mode

CentOS
```
[vagrant@kubedev-10-64-33-82 ~]$ docker run -d --name etcd --net host -v /srv/etcd/etcd0.etcd:/etcd0.etcd -v /srv/etcd/data:/var/lib/etcd --restart always gcr.io/google_containers/etcd-amd64:3.0.17 etcd --listen-client-urls=http://10.64.33.82:2379,http://127.0.0.1:2379 --advertise-client-urls=http://10.64.33.82:2379 --name=etcd0 --initial-advertise-peer-urls=http://10.64.33.82:2380 --listen-peer-urls=http://10.64.33.82:2380 --initial-cluster-token=etcd-cluster-token --initial-cluster=etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380 -initial-cluster-state=new
313248567a0fc25b68314ad830ecc21806e43eefa4d98e22a68771a3bc895530
```

```
[vagrant@kubedev-10-64-33-82 ~]$ docker ps -f name=etcd
CONTAINER ID        IMAGE                                        COMMAND                  CREATED             STATUS              PORTS               NAMES
313248567a0f        gcr.io/google_containers/etcd-amd64:3.0.17   "etcd --listen-client"   16 seconds ago      Up 16 seconds                           etcd
```

Fedora
```
[vagrant@kubedev-10-64-33-199 ~]$ docker run -d --name etcd --net host -v /srv/etcd/etcd1.etcd:/etcd1.etcd -v /srv/etcd/data:/var/lib/etcd --restart always gcr.io/google_containers/etcd-amd64:3.0.17 etcd --listen-client-urls=http://10.64.33.199:2379,http://127.0.0.1:2379 --advertise-client-urls=http://10.64.33.199:2379 --name=etcd1 --initial-advertise-peer-urls=http://10.64.33.199:2380 --listen-peer-urls=http://10.64.33.199:2380 --initial-cluster-token=etcd-cluster-token --initial-cluster=etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380 -initial-cluster-state=new
f13b89af6c64437c43eea8414077f94dde80368b67743259290e797e4729e219
```

```
[vagrant@kubedev-10-64-33-199 ~]$ docker ps -f name=etcd
CONTAINER ID        IMAGE                                        COMMAND                  CREATED             STATUS              PORTS               NAMES
f13b89af6c64        gcr.io/google_containers/etcd-amd64:3.0.17   "etcd --listen-client"   13 seconds ago      Up 12 seconds                           etcd
```

Ubuntu
```
ubuntu@kubedev-10-64-33-195:~$ docker run -d --name etcd --net host -v /srv/etcd/etcd2.etcd:/etcd2.etcd -v /srv/etcd/data:/var/lib/etcd --restart always gcr.io/google_containers/etcd-amd64:3.0.17 etcd --listen-client-urls=http://10.64.33.195:2379,http://127.0.0.1:2379 --advertise-client-urls=http://10.64.33.195:2379 --name=etcd2 --initial-advertise-peer-urls=http://10.64.33.195:2380 --listen-peer-urls=http://10.64.33.195:2380 --initial-cluster-token=etcd-cluster-token --initial-cluster=etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380 -initial-cluster-state=new
55ea2b550443d979f63ca662cfb94a70cd33b3a96c64924dbc9f814c39ac0832
```

```
ubuntu@kubedev-10-64-33-195:~$ docker ps -f name=etcd
CONTAINER ID        IMAGE                                        COMMAND                  CREATED             STATUS              PORTS               NAMES
55ea2b550443        gcr.io/google_containers/etcd-amd64:3.0.17   "etcd --listen-cli..."   20 seconds ago      Up 20 seconds                           etcd
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo du -sh /srv/etcd/etcd0.etcd/
123M	/srv/etcd/etcd0.etcd/
[vagrant@kubedev-10-64-33-199 ~]$ sudo du -sh /srv/etcd/etcd1.etcd/
123M	/srv/etcd/etcd1.etcd/
ubuntu@kubedev-10-64-33-195:~$ sudo du -sh /srv/etcd/etcd2.etcd/
123M	/srv/etcd/etcd2.etcd/
```

__Verify from every one__

v2
```
[vagrant@kubedev-10-64-33-82 ~]$ docker exec -ti etcd etcdctl member list
b54cb34b5858144: name=etcd0 peerURLs=http://10.64.33.82:2380 clientURLs=http://10.64.33.82:2379 isLeader=true
be7241c169024363: name=etcd1 peerURLs=http://10.64.33.199:2380 clientURLs=http://10.64.33.199:2379 isLeader=false
f4fb4b08483d8f4d: name=etcd2 peerURLs=http://10.64.33.195:2380 clientURLs=http://10.64.33.195:2379 isLeader=false
```

```
[vagrant@kubedev-10-64-33-82 ~]$ docker exec -ti etcd etcdctl cluster-health
member b54cb34b5858144 is healthy: got healthy result from http://10.64.33.82:2379
member be7241c169024363 is healthy: got healthy result from http://10.64.33.199:2379
member f4fb4b08483d8f4d is healthy: got healthy result from http://10.64.33.195:2379
cluster is healthy
```

Or v3
```
ubuntu@kubedev-10-64-33-195:~$ docker exec -ti etcd ash -c "ETCDCTL_API=3 etcdctl member list"
b54cb34b5858144, started, etcd0, http://10.64.33.82:2380, http://10.64.33.82:2379
be7241c169024363, started, etcd1, http://10.64.33.199:2380, http://10.64.33.199:2379
f4fb4b08483d8f4d, started, etcd2, http://10.64.33.195:2380, http://10.64.33.195:2379
```

### in kubelet standalone (TBC)

__Generate etcd.yaml into each master node__

CentOS
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo sed 's%listen-client-urls=http://127.0.0.1:2379%listen-client-urls=http://$ip:2379,http://127.0.0.1:2379%;s%advertise-client-urls=http://127.0.0.1:2379%advertise-client-urls=http://$ip:2379%;s%\(data-dir=/var/lib/etcd\)%\1\n    - --name=$name\n    - --initial-advertise-peer-urls=http://$ip:2380\n    - --listen-peer-urls=http://$ip:2380\n    - --initial-cluster-token=etcd-cluster-token\n    - --initial-cluster=$cluster\n    - --initial-cluster-state new%' etcd-single.yaml | env name=etcd0 ip=10.64.33.82 cluster="etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380" envsubst | tee etcd-cluster.yaml
apiVersion: v1
kind: Pod
metadata:
  annotations:
    scheduler.alpha.kubernetes.io/critical-pod: ""
  creationTimestamp: null
  labels:
    component: etcd
    tier: control-plane
  name: etcd
  namespace: kube-system
spec:
  containers:
  - command:
    - etcd
    - --listen-client-urls=http://10.64.33.82:2379,http://127.0.0.1:2379
    - --advertise-client-urls=http://10.64.33.82:2379
    - --data-dir=/var/lib/etcd
    - --name=etcd0
    - --initial-advertise-peer-urls=http://10.64.33.82:2380
    - --listen-peer-urls=http://10.64.33.82:2380
    - --initial-cluster-token=etcd-cluster-token
    - --initial-cluster=etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380
    - --initial-cluster-state new
    image: gcr.io/google_containers/etcd-amd64:3.0.17
    livenessProbe:
      failureThreshold: 8
      httpGet:
        host: 127.0.0.1
        path: /health
        port: 2379
        scheme: HTTP
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: etcd
    resources: {}
    volumeMounts:
    - mountPath: /var/lib/etcd
      name: etcd
  hostNetwork: true
  volumes:
  - hostPath:
      path: /var/lib/etcd
      type: DirectoryOrCreate
    name: etcd
status: {}

```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo mkdir -p /etc/kubernetes/manifests && sudo cp etcd-cluster.yaml /etc/kubernetes/manifests/etcd.yaml
```

Fedora 23
```
[vagrant@kubedev-10-64-33-199 ~]$ sudo sed 's%listen-client-urls=http://127.0.0.1:2379%listen-client-urls=http://$ip:2379,http://127.0.0.1:2379%;s%advertise-client-urls=http://127.0.0.1:2379%advertise-client-urls=http://$ip:2379%;s%\(data-dir=/var/lib/etcd\)%\1\n    - --name=$name\n    - --initial-advertise-peer-urls=http://$ip:2380\n    - --listen-peer-urls=http://$ip:2380\n    - --initial-cluster-token=etcd-cluster-token\n    - --initial-cluster=$cluster\n    - --initial-cluster-state new%' etcd-single.yaml | env name=etcd1 ip=10.64.33.199 cluster="etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380" envsubst | tee etcd-cluster.yaml
apiVersion: v1
kind: Pod
metadata:
  annotations:
    scheduler.alpha.kubernetes.io/critical-pod: ""
  creationTimestamp: null
  labels:
    component: etcd
    tier: control-plane
  name: etcd
  namespace: kube-system
spec:
  containers:
  - command:
    - etcd
    - --listen-client-urls=http://10.64.33.199:2379,http://127.0.0.1:2379
    - --advertise-client-urls=http://10.64.33.199:2379
    - --data-dir=/var/lib/etcd
    - --name=etcd1
    - --initial-advertise-peer-urls=http://10.64.33.199:2380
    - --listen-peer-urls=http://10.64.33.199:2380
    - --initial-cluster-token=etcd-cluster-token
    - --initial-cluster=etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380
    - --initial-cluster-state new
    image: gcr.io/google_containers/etcd-amd64:3.0.17
    livenessProbe:
      failureThreshold: 8
      httpGet:
        host: 127.0.0.1
        path: /health
        port: 2379
        scheme: HTTP
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: etcd
    resources: {}
    volumeMounts:
    - mountPath: /var/lib/etcd
      name: etcd
  hostNetwork: true
  volumes:
  - hostPath:
      path: /var/lib/etcd
      type: DirectoryOrCreate
    name: etcd
status: {}

```

```
[vagrant@kubedev-10-64-33-199 ~]$ sudo mkdir -p /etc/kubernetes/manifests && sudo cp etcd-cluster.yaml /etc/kubernetes/manifests/etcd.yaml
```

Ubuntu
```
ubuntu@kubedev-10-64-33-195:~$ sudo sed 's%listen-client-urls=http://127.0.0.1:2379%listen-client-urls=http://$ip:2379,http://127.0.0.1:2379%;s%advertise-client-urls=http://127.0.0.1:2379%advertise-client-urls=http://$ip:2379%;s%\(data-dir=/var/lib/etcd\)%\1\n    - --name=$name\n    - --initial-advertise-peer-urls=http://$ip:2380\n    - --listen-peer-urls=http://$ip:2380\n    - --initial-cluster-token=etcd-cluster-token\n    - --initial-cluster=$cluster\n    - --initial-cluster-state new%' etcd-single.yaml | env name=etcd2 ip=10.64.33.195 cluster="etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380" envsubst | tee etcd-cluster.yaml
apiVersion: v1
kind: Pod
metadata:
  annotations:
    scheduler.alpha.kubernetes.io/critical-pod: ""
  creationTimestamp: null
  labels:
    component: etcd
    tier: control-plane
  name: etcd
  namespace: kube-system
spec:
  containers:
  - command:
    - etcd
    - --listen-client-urls=http://10.64.33.195:2379,http://127.0.0.1:2379
    - --advertise-client-urls=http://10.64.33.195:2379
    - --data-dir=/var/lib/etcd
    - --name=etcd2
    - --initial-advertise-peer-urls=http://10.64.33.195:2380
    - --listen-peer-urls=http://10.64.33.195:2380
    - --initial-cluster-token=etcd-cluster-token
    - --initial-cluster=etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380
    - --initial-cluster-state new
    image: gcr.io/google_containers/etcd-amd64:3.0.17
    livenessProbe:
      failureThreshold: 8
      httpGet:
        host: 127.0.0.1
        path: /health
        port: 2379
        scheme: HTTP
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: etcd
    resources: {}
    volumeMounts:
    - mountPath: /var/lib/etcd
      name: etcd
  hostNetwork: true
  volumes:
  - hostPath:
      path: /var/lib/etcd
      type: DirectoryOrCreate
    name: etcd
status: {}

```

```
ubuntu@kubedev-10-64-33-195:~$ sudo mkdir -p /etc/kubernetes/manifests && sudo cp etcd-cluster.yaml /etc/kubernetes/manifests/etcd.yaml
```

__restart docker and start kubelet__

CentOS
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo systemctl restart docker.service && sudo systemctl start kubelet.service
```

Fedora
```
[vagrant@kubedev-10-64-33-199 ~]$ sudo systemctl restart docker.service && sudo systemctl start kubelet.service
```

Ubuntu
```
ubuntu@kubedev-10-64-33-195:~$ sudo systemctl restart docker.service && sudo systemctl start kubelet.service
```

## Create multi master configuration

Refert to https://github.com/kubernetes/kubernetes/bolb/master/cmd/kubeadm/app/apis/kubeadm/v1alpha1/types.go

### Create multi-masters-configuration.yaml based upon MasterConfiguration

Generate into CentOS
```
[vagrant@kubedev-10-64-33-82 ~]$ cat multi-masters-configuration.yaml 
apiVersion: kubeadm.k8s.io/v1alpha1
kind: MasterConfiguration
api:
  advertiseAddress: 10.64.33.50
  bindPort: 443
etcd:
  endpoints:
  - http://10.64.33.82:2379
  - http://10.64.33.199:2379
  - http://10.64.33.195:2379
  caFile: ""
  certFile: ""
  keyFile: ""
  dataDir: /var/lib/etcd
  image: ""
  extraArgs: null
kubeletConfiguration:
  baseConfig: null
networking:
  serviceSubnet: 10.96.0.0/12
  podSubnet: 10.244.0.0/16
  dnsDomain: cluster.local
kubernetesVersion: v1.8.4
cloudProvider: ""
nodeName: ""
authorizationModes:
- Node
- RBAC
token: ""
tokenTTL: 24h0m0s
apiServerExtraArgs:
  "admission-control": "Initializers,NamespaceLifecycle,LimitRanger,ServiceAccount,PersistentVolumeLabel,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,ResourceQuota"
controllerManagerExtraArgs: null
SchedulerExtraArgs: null
apiServerExtraVolumes: null
controllerManagerExtraVolumes: null
schedulerExtraVolumes: null
apiServerCertSANs:
- 10.64.33.82
- 10.64.33.199
- 10.64.33.195
- 10.64.33.224
- 10.64.33.240
- kubedev-10-64-33-82
- kubedev-10-64-33-199
- kubedev-10-64-33-195
certificatesDir: /etc/kubernetes/pki
imageRepository: gcr.io/google_containers
unifiedControlPlaneImage: ""
featureGates: null
```

Copy into Fedora for example via local http
```
[vagrant@kubedev-10-64-33-199 ~]$ curl http://10.64.33.82:48080/home/vagrant/multi-masters-configuration.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1184  100  1184    0     0   120k      0 --:--:-- --:--:-- --:--:--  144k
```

Copy into Ubuntu or via local http
```
ubuntu@kubedev-10-64-33-195:~$ curl http://10.64.33.82:48080/home/vagrant/multi-masters-configuration.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1184  100  1184    0     0   110k      0 --:--:-- --:--:-- --:--:--  115k
```

### Init masters

CentOS
```
[vagrant@localhost ~]$ sudo systemctl start kubelet.service; sudo kubeadm init --config multi-masters-configuration.yaml --skip-preflight-checks
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
[kubeconfig] Wrote KubeConfig file to disk: "admin.conf"
[kubeconfig] Wrote KubeConfig file to disk: "kubelet.conf"
[kubeconfig] Wrote KubeConfig file to disk: "controller-manager.conf"
[kubeconfig] Wrote KubeConfig file to disk: "scheduler.conf"
[controlplane] Wrote Static Pod manifest for component kube-apiserver to "/etc/kubernetes/manifests/kube-apiserver.yaml"
[controlplane] Wrote Static Pod manifest for component kube-controller-manager to "/etc/kubernetes/manifests/kube-controller-manager.yaml"
[controlplane] Wrote Static Pod manifest for component kube-scheduler to "/etc/kubernetes/manifests/kube-scheduler.yaml"
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

The front work is doing like
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo kubeadm init --apiserver-advertise-address 10.64.33.50 --apiserver-bind-port 443 --apiserver-cert-extra-sans 10.64.33.82,10.64.33.199,10.64.33.195.10.64.33.224,10.64.33.240 --pod-network-cidr 10.244.0.0/16 --kubernetes-version v1.8.4 --skip-preflight-checks
```

docker v1.11 need cgroufs
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo journalctl --no-tail --pager-end --no-pager --unit kubelet.service
### snippets ###
Nov 26 00:33:39 kubedev-10-64-33-82 kubelet[4879]: error: failed to run Kubelet: failed to create kubelet: misconfiguration: kubelet cgroup driver: "systemd" is different from docker cgroup driver: "cgroupfs"
Nov 26 00:33:39 kubedev-10-64-33-82 systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 26 00:33:39 kubedev-10-64-33-82 systemd[1]: Unit kubelet.service entered failed state.
Nov 26 00:33:39 kubedev-10-64-33-82 systemd[1]: kubelet.service failed.
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo sed -i 's/\(Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"\)/#\1\nEnvironment="KUBELET_CGROUP_ARGS=--cgroup-driver=cgroupfs"/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

```
[vagrant@localhost ~]$ sudo systemctl restart kubelet.service; sudo kubeadm init --config multi-masters-configuration.yaml --skip-preflight-checks
Warning: kubelet.service changed on disk. Run 'systemctl daemon-reload' to reload units.
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
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests"
[init] This often takes around a minute; or longer if the control plane images have to be pulled.
[apiclient] All control plane components are healthy after 93.003519 seconds
[uploadconfig] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[markmaster] Will mark node kubedev-10-64-33-82 as master by adding a label and a taint
[markmaster] Master kubedev-10-64-33-82 tainted and labelled with key/value: node-role.kubernetes.io/master=""
[bootstraptoken] Using token: 90d3e2.adc22e3d43b3faa2
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

  kubeadm join --token 90d3e2.adc22e3d43b3faa2 10.64.33.82:443 --discovery-token-ca-cert-hash sha256:92202ed5b88a724e1206b95dff886ddce6cc67628886b2eeafb95be4a5d888b6

```

### configuration for other nodes

CentOS
```
[vagrant@kubedev-10-64-33-82 ~]$ mkdir -p etc0x2Fkubernetes && sudo rsync -avzq /etc/kubernetes etc0x2Fkubernetes && sudo chown -R vagrant: etc0x2Fkubernetes/ && rm -f etc0x2Fkubernetes/manifests/etcd.yaml 
```

Fedora
```
[vagrant@kubedev-10-64-33-199 ~]$ rsync -avzq -e ssh vagrant@10.64.33.82:/home/vagrant/etc0x2Fkubernetes .
vagrant@10.64.33.82's password: 
```

Ubuntu
```
ubuntu@kubedev-10-64-33-195:~$ rsync -avzq -e ssh vagrant@10.64.33.82:/home/vagrant/etc0x2Fkubernetes .
vagrant@10.64.33.82's password: 
```

### init other masters

Fedora
```
[vagrant@kubedev-10-64-33-199 ~]$ sudo mkdir -p /etc/kubernetes/ && sudo cp -r etc0x2Fkubernetes/pki /etc/kubernetes/ && sudo cp etc0x2Fkubernetes/*.conf /etc/kubernetes/
```

```
[vagrant@kubedev-10-64-33-199 ~]$ sudo systemctl is-active kubelet.service
active
```

```
[vagrant@kubedev-10-64-33-199 ~]$ sudo systemctl start kubelet.service; sudo kubeadm init --config multi-masters-configuration.yaml --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.8.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [kubedev-10-64-33-199 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local kubedev-10-64-33-82 kubedev-10-64-33-199 kubedev-10-64-33-195] and IPs [10.96.0.1 10.64.33.199 10.64.33.82 10.64.33.199 10.64.33.195 10.64.33.224 10.64.33.240 10.64.33.50 172.28.128.50]
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
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests"
[init] This often takes around a minute; or longer if the control plane images have to be pulled.
[apiclient] All control plane components are healthy after 36.005562 seconds
[uploadconfig] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[markmaster] Will mark node kubedev-10-64-33-199 as master by adding a label and a taint
[markmaster] Master kubedev-10-64-33-199 tainted and labelled with key/value: node-role.kubernetes.io/master=""
[bootstraptoken] Using token: 572ca1.69711fceb480d27c
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

  kubeadm join --token 572ca1.69711fceb480d27c 10.64.33.199:443 --discovery-token-ca-cert-hash sha256:92202ed5b88a724e1206b95dff886ddce6cc67628886b2eeafb95be4a5d888b6

```

Ubuntu
```
ubuntu@kubedev-10-64-33-195:~$ sudo mkdir -p /etc/kubernetes/ && sudo cp -r etc0x2Fkubernetes/pki /etc/kubernetes/ && sudo cp etc0x2Fkubernetes/*.conf /etc/kubernetes/
```

```
ubuntu@kubedev-10-64-33-195:~$ sudo systemctl is-active kubelet.service
failed
```

```
ubuntu@kubedev-10-64-33-195:~$ sudo systemctl start kubelet.service; sudo kubeadm init --config multi-masters-configuration.yaml --skip-preflight-checks
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
[kubeconfig] Wrote KubeConfig file to disk: "admin.conf"
[kubeconfig] Wrote KubeConfig file to disk: "kubelet.conf"
[kubeconfig] Wrote KubeConfig file to disk: "controller-manager.conf"
[kubeconfig] Wrote KubeConfig file to disk: "scheduler.conf"
[controlplane] Wrote Static Pod manifest for component kube-apiserver to "/etc/kubernetes/manifests/kube-apiserver.yaml"
[controlplane] Wrote Static Pod manifest for component kube-controller-manager to "/etc/kubernetes/manifests/kube-controller-manager.yaml"
[controlplane] Wrote Static Pod manifest for component kube-scheduler to "/etc/kubernetes/manifests/kube-scheduler.yaml"
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests"
[init] This often takes around a minute; or longer if the control plane images have to be pulled.
[apiclient] All control plane components are healthy after 22.502495 seconds
[uploadconfig] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[markmaster] Will mark node kubedev-10-64-33-195 as master by adding a label and a taint
[markmaster] Master kubedev-10-64-33-195 tainted and labelled with key/value: node-role.kubernetes.io/master=""
[bootstraptoken] Using token: ac8d54.a8e5e072b71f78c0
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

  kubeadm join --token ac8d54.a8e5e072b71f78c0 10.64.33.195:443 --discovery-token-ca-cert-hash sha256:92202ed5b88a724e1206b95dff886ddce6cc67628886b2eeafb95be4a5d888b6

```

### etcdctl

Download
```
[vagrant@kubedev-10-64-33-82 ~]$ curl -jkSL https://github.com/coreos/etcd/releases/download/v3.2.10/etcd-v3.2.10-linux-amd64.tar.gz | tar -C /vagrant_f/99-mirror/linux-bin/ -zx
```

Verify
```
[vagrant@kubedev-10-64-33-82 ~]$ /vagrant_f/99-mirror/linux-bin/etcdctl --endpoints= http://10.64.33.82:2379,http://10.64.33.199:2379,http://10.64.33.195:2379 member list
b54cb34b5858144: name=etcd0 peerURLs=http://10.64.33.82:2380 clientURLs=http://10.64.33.82:2379 isLeader=false
be7241c169024363: name=etcd1 peerURLs=http://10.64.33.199:2380 clientURLs=http://10.64.33.199:2379 isLeader=true
f4fb4b08483d8f4d: name=etcd2 peerURLs=http://10.64.33.195:2380 clientURLs=http://10.64.33.195:2379 isLeader=false
```

```
ubuntu@kubedev-10-64-33-195:~$ /vagrant_f/99-mirror/linux-bin/etcdctl --endpoints= http://10.64.33.82:2379,http://10.64.33.199:2379,http://10.64.33.195:2379 cluster-health
member b54cb34b5858144 is healthy: got healthy result from http://10.64.33.82:2379
member be7241c169024363 is healthy: got healthy result from http://10.64.33.199:2379
member f4fb4b08483d8f4d is healthy: got healthy result from http://10.64.33.195:2379
cluster is healthy
```

### investigation with insecure

kube-flannel.yaml
```
        command: [ "/opt/bin/flanneld", "--ip-masq", "--kube-subnet-mgr", "--iface=enp0s8", "--iface=eth1", "--iface-regex=10\\.64\\.33\\.[0-9]+", "--kube-api-url=http://10.64.33.50:8080" ]
```

to set up cni
```
[vagrant@kubedev-10-64-33-199 ~]$ kubectl -n kube-system get pods -o wide -l app=flannel
NAME                    READY     STATUS    RESTARTS   AGE       IP             NODE
kube-flannel-ds-2w4xb   1/1       Running   0          27m       10.64.33.195   kubedev-10-64-33-195
kube-flannel-ds-kh97l   1/1       Running   0          27m       10.64.33.199   kubedev-10-64-33-199
kube-flannel-ds-v8lpx   1/1       Running   1          27m       10.64.33.200   kubedev-10-64-33-200
kube-flannel-ds-vs7xj   1/1       Running   0          27m       10.64.33.82    kubedev-10-64-33-82
```

kube-dns.yaml
```
      - args:
        ###
        - --kube-master-url=http://10.64.33.50:8080
        ###
```

to test dns ok
```
[vagrant@kubedev-10-64-33-199 ~]$ kubectl -n kube-system get pods -o wide -l k8s-app=kube-dns
NAME                        READY     STATUS    RESTARTS   AGE       IP           NODE
kube-dns-5f99695d95-5k27s   3/3       Running   3          38m       10.244.3.4   kubedev-10-64-33-200
kube-dns-5f99695d95-7zk7n   3/3       Running   0          38m       10.244.2.9   kubedev-10-64-33-82
kube-dns-5f99695d95-j4sb6   3/3       Running   0          36m       10.244.0.8   kubedev-10-64-33-199
kube-dns-5f99695d95-nczf2   3/3       Running   0          38m       10.244.1.7   kubedev-10-64-33-195
```

### dns issue

__Only 2 containers working__
```
ubuntu@kubedev-10-64-33-195:~$ kubectl get pods -n kube-system -l k8s-app=kube-dns -o wide
NAME                                           READY     STATUS             RESTARTS   AGE       IP             NODE
kube-dns-545bc4bfd4-5b8qr                      2/3       Running            553        23h       10.244.0.3     kubedev-10-64-33-199
```

refer to https://github.com/kubernetes/kubernetes/issues/33798

CentOS
```
[vagrant@kubedev-10-64-33-82 ~]$ grep net.bridge /usr/lib/sysctl.d/*
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-ip6tables = 0
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-iptables = 0
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-arptables = 0
```

override and update
```
[vagrant@kubedev-10-64-33-82 ~]$ echo -e "net.bridge.bridge-nf-call-ip6tables = 1\nnet.bridge.bridge-nf-call-iptables = 1\nnet.bridge.bridge-nf-call-arptables = 1" | sudo tee /etc/sysctl.d/00-system.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-arptables = 1
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo sysctl -w net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-iptables = 1
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo sysctl -w net.bridge.bridge-nf-call-ip6tables=1
net.bridge.bridge-nf-call-ip6tables = 1
```

Fedora
```
[vagrant@kubedev-10-64-33-199 ~]$ grep net.bridge /usr/lib/sysctl.d/*
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-ip6tables = 0
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-iptables = 0
/usr/lib/sysctl.d/00-system.conf:net.bridge.bridge-nf-call-arptables = 0
```

override and update
```
[vagrant@kubedev-10-64-33-199 ~]$ echo -e "net.bridge.bridge-nf-call-ip6tables = 1\nnet.bridge.bridge-nf-call-iptables = 1\nnet.bridge.bridge-nf-call-arptables = 1" | sudo tee /etc/sysctl.d/00-system.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-arptables = 1
```

```
[vagrant@kubedev-10-64-33-199 ~]$ sudo sysctl -w net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-iptables = 1
```

```
[vagrant@kubedev-10-64-33-199 ~]$ sudo sysctl -w net.bridge.bridge-nf-call-ip6tables=1
net.bridge.bridge-nf-call-ip6tables = 1
```

Ubuntu, modify /etc/sysctl.conf and update with sysctl -w
```
net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-ip6tables=1
```

```
ubuntu@kubedev-10-64-33-195:~$ sudo sysctl -w net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-iptables = 1
```

Refer to https://github.com/kubernetes/kubeadm/issues/193

CentOS7
```
[vagrant@kubedev-10-64-33-82 ~]$ grep "net.ipv4.conf.all" /usr/lib/sysctl.d/*
/usr/lib/sysctl.d/50-default.conf:net.ipv4.conf.all.rp_filter = 1
/usr/lib/sysctl.d/50-default.conf:net.ipv4.conf.all.accept_source_route = 0
/usr/lib/sysctl.d/50-default.conf:net.ipv4.conf.all.promote_secondaries = 1
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo sysctl -w net.ipv4.conf.all.forwarding=1
net.ipv4.conf.all.forwarding = 1
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo iptables -P FORWARD ACCEPT
```

refer to https://github.com/kubernetes/contrib/issues/2249
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo iptables -t nat -A PREROUTING -d 10.96.0.1 -p tcp --dport 443 -j DNAT --to-destination 127.0.0.1:443
```

log
```
ubuntu@kubedev-10-64-33-195:~$ kubectl logs pods/kube-dns-545bc4bfd4-5b8qr --namespace=kube-system -c kubedns
I1126 10:04:05.271905       1 dns.go:48] version: 1.14.4-2-g5584e04
I1126 10:04:05.275150       1 server.go:70] Using configuration read from directory: /kube-dns-config with period 10s
I1126 10:04:05.275493       1 server.go:113] FLAG: --alsologtostderr="false"
I1126 10:04:05.275678       1 server.go:113] FLAG: --config-dir="/kube-dns-config"
I1126 10:04:05.275835       1 server.go:113] FLAG: --config-map=""
I1126 10:04:05.275982       1 server.go:113] FLAG: --config-map-namespace="kube-system"
I1126 10:04:05.276128       1 server.go:113] FLAG: --config-period="10s"
I1126 10:04:05.276362       1 server.go:113] FLAG: --dns-bind-address="0.0.0.0"
I1126 10:04:05.276512       1 server.go:113] FLAG: --dns-port="10053"
I1126 10:04:05.276657       1 server.go:113] FLAG: --domain="cluster.local."
I1126 10:04:05.276800       1 server.go:113] FLAG: --federations=""
I1126 10:04:05.276939       1 server.go:113] FLAG: --healthz-port="8081"
I1126 10:04:05.277078       1 server.go:113] FLAG: --initial-sync-timeout="1m0s"
I1126 10:04:05.277472       1 server.go:113] FLAG: --kube-master-url=""
I1126 10:04:05.277637       1 server.go:113] FLAG: --kubecfg-file=""
I1126 10:04:05.277782       1 server.go:113] FLAG: --log-backtrace-at=":0"
I1126 10:04:05.277793       1 server.go:113] FLAG: --log-dir=""
I1126 10:04:05.277797       1 server.go:113] FLAG: --log-flush-frequency="5s"
I1126 10:04:05.277801       1 server.go:113] FLAG: --logtostderr="true"
I1126 10:04:05.277804       1 server.go:113] FLAG: --nameservers=""
I1126 10:04:05.277806       1 server.go:113] FLAG: --stderrthreshold="2"
I1126 10:04:05.277810       1 server.go:113] FLAG: --v="2"
I1126 10:04:05.277813       1 server.go:113] FLAG: --version="false"
I1126 10:04:05.277819       1 server.go:113] FLAG: --vmodule=""
I1126 10:04:05.282119       1 server.go:176] Starting SkyDNS server (0.0.0.0:10053)
I1126 10:04:05.282629       1 server.go:198] Skydns metrics enabled (/metrics:10055)
I1126 10:04:05.282819       1 dns.go:147] Starting endpointsController
I1126 10:04:05.282967       1 dns.go:150] Starting serviceController
I1126 10:04:05.296442       1 logs.go:41] skydns: ready for queries on cluster.local. for tcp://0.0.0.0:10053 [rcache 0]
I1126 10:04:05.296644       1 logs.go:41] skydns: ready for queries on cluster.local. for udp://0.0.0.0:10053 [rcache 0]
I1126 10:04:05.783434       1 dns.go:174] Waiting for services and endpoints to be initialized from apiserver...
I1126 10:04:06.283362       1 dns.go:174] Waiting for services and endpoints to be initialized from apiserver...
E1126 10:04:06.365273       1 reflector.go:199] k8s.io/dns/vendor/k8s.io/client-go/tools/cache/reflector.go:94: Failed to list *v1.Endpoints: Get https://10.96.0.1:443/api/v1/endpoints?resourceVersion=0: dial tcp 10.96.0.1:443: getsockopt: connection refused
E1126 10:04:06.365647       1 reflector.go:199] k8s.io/dns/vendor/k8s.io/client-go/tools/cache/reflector.go:94: Failed to list *v1.Service: Get https://10.96.0.1:443/api/v1/services?resourceVersion=0: dial tcp 10.96.0.1:443: getsockopt: connection refused
I1126 10:04:06.784595       1 dns.go:174] Waiting for services and endpoints to be initialized from apiserver...
I1126 10:04:07.283395       1 dns.go:174] Waiting for services and endpoints to be initialized from apiserver...
I1126 10:04:07.783340       1 dns.go:174] Waiting for services and endpoints to be initialized from apiserver...
I1126 10:04:08.283336       1 dns.go:174] Waiting for services and endpoints to be initialized from apiserver...
### snippets ###
```

refer to https://github.com/kubernetes/kubernetes/issues/54910

```
ubuntu@kubedev-10-64-33-195:~$ kubectl logs pods/kube-dns-545bc4bfd4-5b8qr --namespace=kube-system -c dnsmasq
I1126 10:08:25.846895       1 main.go:76] opts: {{/usr/sbin/dnsmasq [-k --cache-size=1000 --log-facility=- --server=/cluster.local/127.0.0.1#10053 --server=/in-addr.arpa/127.0.0.1#10053 --server=/ip6.arpa/127.0.0.1#10053] true} /etc/k8s/dns/dnsmasq-nanny 10000000000}
I1126 10:08:25.847591       1 nanny.go:86] Starting dnsmasq [-k --cache-size=1000 --log-facility=- --server=/cluster.local/127.0.0.1#10053 --server=/in-addr.arpa/127.0.0.1#10053 --server=/ip6.arpa/127.0.0.1#10053]
I1126 10:08:25.944380       1 nanny.go:111] 
W1126 10:08:25.944784       1 nanny.go:112] Got EOF from stdout
I1126 10:08:25.944967       1 nanny.go:108] dnsmasq[8]: started, version 2.78-security-prerelease cachesize 1000
I1126 10:08:25.945144       1 nanny.go:108] dnsmasq[8]: compile time options: IPv6 GNU-getopt no-DBus no-i18n no-IDN DHCP DHCPv6 no-Lua TFTP no-conntrack ipset auth no-DNSSEC loop-detect inotify
I1126 10:08:25.945309       1 nanny.go:108] dnsmasq[8]: using nameserver 127.0.0.1#10053 for domain ip6.arpa 
I1126 10:08:25.945466       1 nanny.go:108] dnsmasq[8]: using nameserver 127.0.0.1#10053 for domain in-addr.arpa 
I1126 10:08:25.945624       1 nanny.go:108] dnsmasq[8]: using nameserver 127.0.0.1#10053 for domain cluster.local 
I1126 10:08:25.947250       1 nanny.go:108] dnsmasq[8]: reading /etc/resolv.conf
I1126 10:08:25.947430       1 nanny.go:108] dnsmasq[8]: using nameserver 127.0.0.1#10053 for domain ip6.arpa 
I1126 10:08:25.947589       1 nanny.go:108] dnsmasq[8]: using nameserver 127.0.0.1#10053 for domain in-addr.arpa 
I1126 10:08:25.947600       1 nanny.go:108] dnsmasq[8]: using nameserver 127.0.0.1#10053 for domain cluster.local 
I1126 10:08:25.947607       1 nanny.go:108] dnsmasq[8]: using nameserver 10.0.2.3#53
I1126 10:08:25.947641       1 nanny.go:108] dnsmasq[8]: read /etc/hosts - 7 addresses
```

```
ubuntu@kubedev-10-64-33-195:~$ kubectl logs pods/kube-dns-545bc4bfd4-5b8qr --namespace=kube-system -c sidecar
ERROR: logging before flag.Parse: I1125 18:20:59.932030       1 main.go:48] Version v1.14.4-2-g5584e04
ERROR: logging before flag.Parse: I1125 18:20:59.932588       1 server.go:45] Starting server (options {DnsMasqPort:53 DnsMasqAddr:127.0.0.1 DnsMasqPollIntervalMs:5000 Probes:[{Label:kubedns Server:127.0.0.1:10053 Name:kubernetes.default.svc.cluster.local. Interval:5s Type:1} {Label:dnsmasq Server:127.0.0.1:53 Name:kubernetes.default.svc.cluster.local. Interval:5s Type:1}] PrometheusAddr:0.0.0.0 PrometheusPort:10054 PrometheusPath:/metrics PrometheusNamespace:kubedns})
ERROR: logging before flag.Parse: I1125 18:20:59.933426       1 dnsprobe.go:75] Starting dnsProbe {Label:kubedns Server:127.0.0.1:10053 Name:kubernetes.default.svc.cluster.local. Interval:5s Type:1}
ERROR: logging before flag.Parse: I1125 18:20:59.933674       1 dnsprobe.go:75] Starting dnsProbe {Label:dnsmasq Server:127.0.0.1:53 Name:kubernetes.default.svc.cluster.local. Interval:5s Type:1}
ERROR: logging before flag.Parse: W1125 18:23:10.017303       1 server.go:64] Error getting metrics from dnsmasq: read udp 127.0.0.1:59977->127.0.0.1:53: read: connection refused
ERROR: logging before flag.Parse: W1125 18:25:25.084836       1 server.go:64] Error getting metrics from dnsmasq: read udp 127.0.0.1:60673->127.0.0.1:53: read: connection refused
### snippets ###
```

```
ubuntu@kubedev-10-64-33-195:~$ sudo sysctl -w net.bridge.bridge-nf-call-ip6tables=1
net.bridge.bridge-nf-call-ip6tables = 1
```

__Crash__
```
[vagrant@kubedev-10-64-33-199 ~]$ kubectl get pods -n kube-system -l k8s-app=kube-dns -o wide
NAME                        READY     STATUS             RESTARTS   AGE       IP           NODE
kube-dns-545bc4bfd4-5b8qr   1/3       CrashLoopBackOff   571        1d        10.244.0.3   kubedev-10-64-33-199
```

bridge is created
```
[vagrant@kubedev-10-64-33-199 ~]$ brctl show
bridge name	bridge id		STP enabled	interfaces
cni0		8000.0a580af40001	no		vethf69ae17d
docker0		8000.02429c9d92dd	no		
```

but not managed by docker
```
ubuntu@kubedev-10-64-33-195:~$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
8af67210633e        bridge              bridge              local
332afc44c03b        host                host                local
fef0a0a326c9        none                null                local
```

```
[vagrant@kubedev-10-64-33-199 ~]$ ip a show cni0
33: cni0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default 
    link/ether 0a:58:0a:f4:00:01 brd ff:ff:ff:ff:ff:ff
    inet 10.244.0.1/24 scope global cni0
       valid_lft forever preferred_lft forever
    inet6 fe80::3c10:65ff:fedc:c35b/64 scope link 
       valid_lft forever preferred_lft forever
```

For v1.11, modify docker opts with `docker daemon --bridge=cni0 ...` in /etc/sysconfig/docker ...
```
[vagrant@kubedev-10-64-33-199 ~]$ sudo systemctl daemon-reload && sudo systemctl restart docker.service
```

```
ubuntu@kubedev-10-64-33-195:~$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
c5de57a5a3a6        bridge              bridge              local
332afc44c03b        host                host                local
fef0a0a326c9        none                null                local
```

```
[vagrant@kubedev-10-64-33-199 ~]$ docker network inspect bridge
[
    {
        "Name": "bridge",
        "Id": "d2921e21bf168935652a8d24e173fd6d1a17a32faeff2ae4a4722560f174b20c",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "10.244.0.0/24",
                    "Gateway": "10.244.0.1"
                }
            ]
        },
        "Internal": false,
        "Containers": {},
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "cni0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {}
    }
]
```

```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl --namespace=kube-system get pods -o wide
NAME                                           READY     STATUS    RESTARTS   AGE       IP             NODE
kube-apiserver-kubedev-10-64-33-195            1/1       Running   0          2m        10.64.33.195   kubedev-10-64-33-195
kube-apiserver-kubedev-10-64-33-199            1/1       Running   0          58s       10.64.33.199   kubedev-10-64-33-199
kube-apiserver-kubedev-10-64-33-82             1/1       Running   0          1m        10.64.33.82    kubedev-10-64-33-82
kube-controller-manager-kubedev-10-64-33-195   1/1       Running   0          2m        10.64.33.195   kubedev-10-64-33-195
kube-controller-manager-kubedev-10-64-33-199   1/1       Running   0          55s       10.64.33.199   kubedev-10-64-33-199
kube-controller-manager-kubedev-10-64-33-82    1/1       Running   0          51s       10.64.33.82    kubedev-10-64-33-82
kube-dns-545bc4bfd4-wbdfl                      3/3       Running   0          1m        10.244.2.15    kubedev-10-64-33-82
kube-proxy-4bmmw                               1/1       Running   0          1m        10.64.33.195   kubedev-10-64-33-195
kube-proxy-nnjtz                               1/1       Running   0          1m        10.64.33.199   kubedev-10-64-33-199
kube-proxy-rv4qt                               1/1       Running   0          1m        10.64.33.82    kubedev-10-64-33-82
kube-scheduler-kubedev-10-64-33-195            1/1       Running   0          42s       10.64.33.195   kubedev-10-64-33-195
kube-scheduler-kubedev-10-64-33-199            1/1       Running   0          1m        10.64.33.199   kubedev-10-64-33-199
kube-scheduler-kubedev-10-64-33-82             1/1       Running   28         1h        10.64.33.82    kubedev-10-64-33-82
```

scale
```
ubuntu@kubedev-10-64-33-195:~$ kubectl scale deployment/kube-dns --namespace=kube-system --replicas=3
deployment "kube-dns" scaled
```

```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl -n kube-system get pods -o wide
NAME                                           READY     STATUS    RESTARTS   AGE       IP             NODE
kube-apiserver-kubedev-10-64-33-195            1/1       Running   3          1h        10.64.33.195   kubedev-10-64-33-195
kube-apiserver-kubedev-10-64-33-199            1/1       Running   10         1h        10.64.33.199   kubedev-10-64-33-199
kube-apiserver-kubedev-10-64-33-82             1/1       Running   4          1h        10.64.33.82    kubedev-10-64-33-82
kube-controller-manager-kubedev-10-64-33-195   1/1       Running   3          1h        10.64.33.195   kubedev-10-64-33-195
kube-controller-manager-kubedev-10-64-33-199   1/1       Running   10         1h        10.64.33.199   kubedev-10-64-33-199
kube-controller-manager-kubedev-10-64-33-82    1/1       Running   4          1h        10.64.33.82    kubedev-10-64-33-82
kube-dns-545bc4bfd4-mqq8j                      3/3       Running   16         1h        10.244.0.237   kubedev-10-64-33-195
kube-dns-545bc4bfd4-sflc2                      3/3       Running   11         1h        10.244.2.62    kubedev-10-64-33-199
kube-dns-545bc4bfd4-wbdfl                      3/3       Running   3          1h        10.244.1.194   kubedev-10-64-33-82
kube-proxy-4bmmw                               1/1       Running   4          1h        10.64.33.195   kubedev-10-64-33-195
kube-proxy-nnjtz                               1/1       Running   11         1h        10.64.33.199   kubedev-10-64-33-199
kube-proxy-rv4qt                               1/1       Running   4          1h        10.64.33.82    kubedev-10-64-33-82
kube-scheduler-kubedev-10-64-33-195            1/1       Running   3          1h        10.64.33.195   kubedev-10-64-33-195
kube-scheduler-kubedev-10-64-33-199            1/1       Running   11         1h        10.64.33.199   kubedev-10-64-33-199
kube-scheduler-kubedev-10-64-33-82             1/1       Running   32         3h        10.64.33.82    kubedev-10-64-33-82
```

### networking

flannel, change arguments for support destination networking interface
```
ubuntu@kubedev-10-64-33-195:~$ egrep 'command: \[' kube-flannel.yaml 
        command: [ "/opt/bin/flanneld", "--ip-masq", "--kube-subnet-mgr", "--iface=enp0s8", "--iface=eth1", "--iface-regex=10\\.64\\.33\\.[0-9]+" ]
```

```
ubuntu@kubedev-10-64-33-195:~$ kubectl create -f kube-flannel.yaml 
clusterrole "flannel" created
clusterrolebinding "flannel" created
serviceaccount "flannel" created
configmap "kube-flannel-cfg" created
daemonset "kube-flannel-ds" created
```

Issue
```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl logs kube-flannel-ds-5npq4 -n kube-system
I1126 19:11:12.468564       1 main.go:201] Could not find valid interface matching enp0s8: error looking up interface enp0s8: route ip+net: no such network interface
I1126 19:11:12.469458       1 main.go:483] Using interface with name eth1 and address 10.64.33.82
I1126 19:11:12.469471       1 main.go:500] Defaulting external address to interface address (10.64.33.82)
E1126 19:11:16.030754       1 main.go:232] Failed to create SubnetManager: error retrieving pod spec for 'kube-system/kube-flannel-ds-5npq4': Get https://10.96.0.1:443/api/v1/namespaces/kube-system/pods/kube-flannel-ds-5npq4: dial tcp 10.96.0.1:443: getsockopt: connection refused
```

```
[vagrant@kubedev-10-64-33-199 ~]$  kubectl -n kube-system get pods -o wide
NAME                                           READY     STATUS    RESTARTS   AGE       IP             NODE
kube-apiserver-kubedev-10-64-33-195            1/1       Running   3          1h        10.64.33.195   kubedev-10-64-33-195
kube-apiserver-kubedev-10-64-33-199            1/1       Running   10         1h        10.64.33.199   kubedev-10-64-33-199
kube-apiserver-kubedev-10-64-33-82             1/1       Running   4          1h        10.64.33.82    kubedev-10-64-33-82
kube-controller-manager-kubedev-10-64-33-195   1/1       Running   3          1h        10.64.33.195   kubedev-10-64-33-195
kube-controller-manager-kubedev-10-64-33-199   1/1       Running   10         1h        10.64.33.199   kubedev-10-64-33-199
kube-controller-manager-kubedev-10-64-33-82    1/1       Running   4          1h        10.64.33.82    kubedev-10-64-33-82
kube-dns-545bc4bfd4-mqq8j                      3/3       Running   16         1h        10.244.0.237   kubedev-10-64-33-195
kube-dns-545bc4bfd4-sflc2                      3/3       Running   11         1h        10.244.2.62    kubedev-10-64-33-199
kube-dns-545bc4bfd4-wbdfl                      3/3       Running   3          1h        10.244.1.194   kubedev-10-64-33-82
kube-flannel-ds-b5784                          1/1       Running   0          21s       10.64.33.82    kubedev-10-64-33-82
kube-flannel-ds-gk826                          1/1       Running   0          21s       10.64.33.195   kubedev-10-64-33-195
kube-flannel-ds-tbzdp                          1/1       Running   0          21s       10.64.33.199   kubedev-10-64-33-199
kube-proxy-4bmmw                               1/1       Running   4          1h        10.64.33.195   kubedev-10-64-33-195
kube-proxy-nnjtz                               1/1       Running   11         1h        10.64.33.199   kubedev-10-64-33-199
kube-proxy-rv4qt                               1/1       Running   4          1h        10.64.33.82    kubedev-10-64-33-82
kube-scheduler-kubedev-10-64-33-195            1/1       Running   3          1h        10.64.33.195   kubedev-10-64-33-195
kube-scheduler-kubedev-10-64-33-199            1/1       Running   11         1h        10.64.33.199   kubedev-10-64-33-199
kube-scheduler-kubedev-10-64-33-82             1/1       Running   32         3h        10.64.33.82    kubedev-10-64-33-82
```

### summary

nodes
```
[vagrant@kubedev-10-64-33-199 ~]$ kubectl get nodes -o wide
NAME                   STATUS    ROLES     AGE       VERSION   EXTERNAL-IP   OS-IMAGE                   KERNEL-VERSION              CONTAINER-RUNTIME
kubedev-10-64-33-195   Ready     master    3h        v1.8.4    <none>        Ubuntu 16.04.3 LTS         4.4.0-101-generic           docker://1.13.1
kubedev-10-64-33-199   Ready     master    3h        v1.8.4    <none>        Fedora 23 (Twenty Three)   4.8.13-100.fc23.x86_64      docker://1.11.2
kubedev-10-64-33-82    Ready     master    3h        v1.8.4    <none>        CentOS Linux 7 (Core)      3.10.0-693.5.2.el7.x86_64   docker://1.11.2
```

node 0
```
[vagrant@kubedev-10-64-33-82 ~]$ docker ps
CONTAINER ID        IMAGE                                        COMMAND                  CREATED              STATUS              PORTS               NAMES
3ff0f240571e        4c600a64a18a                                 "/opt/bin/flanneld --"   About a minute ago   Up About a minute                       k8s_kube-flannel_kube-flannel-ds-b5784_kube-system_a246b5e0-d3cf-11e7-b63a-024403e322ab_0
c3cb26e1f6db        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 About a minute ago   Up About a minute                       k8s_POD_kube-flannel-ds-b5784_kube-system_a246b5e0-d3cf-11e7-b63a-024403e322ab_0
7c02dea0e5cc        fed89e8b4248                                 "/sidecar --v=2 --log"   3 minutes ago        Up 3 minutes                            k8s_sidecar_kube-dns-545bc4bfd4-wbdfl_kube-system_da15b25c-d3c2-11e7-9c69-080027eae602_1
7d3d2058ed09        459944ce8cc4                                 "/dnsmasq-nanny -v=2 "   3 minutes ago        Up 3 minutes                            k8s_dnsmasq_kube-dns-545bc4bfd4-wbdfl_kube-system_da15b25c-d3c2-11e7-9c69-080027eae602_1
f75b3775061a        512cd7425a73                                 "/kube-dns --domain=c"   3 minutes ago        Up 3 minutes                            k8s_kubedns_kube-dns-545bc4bfd4-wbdfl_kube-system_da15b25c-d3c2-11e7-9c69-080027eae602_1
9f2a3204250b        65a61c14e8c2                                 "/usr/local/bin/kube-"   3 minutes ago        Up 3 minutes                            k8s_kube-proxy_kube-proxy-rv4qt_kube-system_da0658a5-d3c2-11e7-9c69-080027eae602_4
49202ebc5709        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 3 minutes ago        Up 3 minutes                            k8s_POD_kube-dns-545bc4bfd4-wbdfl_kube-system_da15b25c-d3c2-11e7-9c69-080027eae602_1
d273af23fb6f        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 3 minutes ago        Up 3 minutes                            k8s_POD_kube-proxy-rv4qt_kube-system_da0658a5-d3c2-11e7-9c69-080027eae602_4
d8e36527465b        7058ac4d4af5                                 "kube-controller-mana"   3 minutes ago        Up 3 minutes                            k8s_kube-controller-manager_kube-controller-manager-kubedev-10-64-33-82_kube-system_2abd80eadd4e71350cbaac2862cae7c7_4
5014d15c95c5        0d985fed7f95                                 "kube-scheduler --add"   3 minutes ago        Up 3 minutes                            k8s_kube-scheduler_kube-scheduler-kubedev-10-64-33-82_kube-system_e554495c6f8701f21accd04866090b05_32
41087bf6fae1        10a052dccbc5                                 "kube-apiserver --all"   3 minutes ago        Up 3 minutes                            k8s_kube-apiserver_kube-apiserver-kubedev-10-64-33-82_kube-system_32325e2febd61625afff8b2e43041954_4
af260880b247        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 3 minutes ago        Up 3 minutes                            k8s_POD_kube-controller-manager-kubedev-10-64-33-82_kube-system_2abd80eadd4e71350cbaac2862cae7c7_4
0b7378e9a168        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 3 minutes ago        Up 3 minutes                            k8s_POD_kube-scheduler-kubedev-10-64-33-82_kube-system_e554495c6f8701f21accd04866090b05_9
0fc8e028488e        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 3 minutes ago        Up 3 minutes                            k8s_POD_kube-apiserver-kubedev-10-64-33-82_kube-system_32325e2febd61625afff8b2e43041954_4
313248567a0f        gcr.io/google_containers/etcd-amd64:3.0.17   "etcd --listen-client"   3 hours ago          Up 12 minutes                           etcd
```

node 1
```
[vagrant@kubedev-10-64-33-199 ~]$ docker ps
CONTAINER ID        IMAGE                                        COMMAND                  CREATED             STATUS              PORTS                    NAMES
c6e95b833c7d        4c600a64a18a                                 "/opt/bin/flanneld --"   3 minutes ago       Up 3 minutes                                 k8s_kube-flannel_kube-flannel-ds-tbzdp_kube-system_a24726f0-d3cf-11e7-b63a-024403e322ab_0
523f5a75606b        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 3 minutes ago       Up 3 minutes                                 k8s_POD_kube-flannel-ds-tbzdp_kube-system_a24726f0-d3cf-11e7-b63a-024403e322ab_0
ee7bc696f10e        fed89e8b4248                                 "/sidecar --v=2 --log"   4 minutes ago       Up 4 minutes                                 k8s_sidecar_kube-dns-545bc4bfd4-sflc2_kube-system_09bca1bf-d3c7-11e7-a68d-024403e322ab_3
b370cdb6939f        459944ce8cc4                                 "/dnsmasq-nanny -v=2 "   4 minutes ago       Up 4 minutes                                 k8s_dnsmasq_kube-dns-545bc4bfd4-sflc2_kube-system_09bca1bf-d3c7-11e7-a68d-024403e322ab_3
5d7e88344722        65a61c14e8c2                                 "/usr/local/bin/kube-"   4 minutes ago       Up 4 minutes                                 k8s_kube-proxy_kube-proxy-nnjtz_kube-system_da068521-d3c2-11e7-9c69-080027eae602_11
b1d513be87ea        512cd7425a73                                 "/kube-dns --domain=c"   4 minutes ago       Up 4 minutes                                 k8s_kubedns_kube-dns-545bc4bfd4-sflc2_kube-system_09bca1bf-d3c7-11e7-a68d-024403e322ab_5
7115cc5d4fdd        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 4 minutes ago       Up 4 minutes                                 k8s_POD_kube-proxy-nnjtz_kube-system_da068521-d3c2-11e7-9c69-080027eae602_13
0dc1d6dbf38f        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 4 minutes ago       Up 4 minutes                                 k8s_POD_kube-dns-545bc4bfd4-sflc2_kube-system_09bca1bf-d3c7-11e7-a68d-024403e322ab_568
ba3fc8e555e9        7058ac4d4af5                                 "kube-controller-mana"   5 minutes ago       Up 4 minutes                                 k8s_kube-controller-manager_kube-controller-manager-kubedev-10-64-33-199_kube-system_d7178e08c00d6217be0c45efd8413533_10
c4e4382674fa        0d985fed7f95                                 "kube-scheduler --add"   5 minutes ago       Up 5 minutes                                 k8s_kube-scheduler_kube-scheduler-kubedev-10-64-33-199_kube-system_e554495c6f8701f21accd04866090b05_11
137af55e28d4        10a052dccbc5                                 "kube-apiserver --sec"   5 minutes ago       Up 5 minutes                                 k8s_kube-apiserver_kube-apiserver-kubedev-10-64-33-199_kube-system_32337c60b0b723900a81a9c2f48359dd_10
b122d7a49d34        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 5 minutes ago       Up 5 minutes                                 k8s_POD_kube-controller-manager-kubedev-10-64-33-199_kube-system_d7178e08c00d6217be0c45efd8413533_15
24d953946248        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 5 minutes ago       Up 5 minutes                                 k8s_POD_kube-scheduler-kubedev-10-64-33-199_kube-system_e554495c6f8701f21accd04866090b05_16
84c5619e75a1        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 5 minutes ago       Up 5 minutes                                 k8s_POD_kube-apiserver-kubedev-10-64-33-199_kube-system_32337c60b0b723900a81a9c2f48359dd_12
f13b89af6c64        gcr.io/google_containers/etcd-amd64:3.0.17   "etcd --listen-client"   3 hours ago         Up 5 minutes                                 etcd
```

node 2
```
ubuntu@kubedev-10-64-33-195:~$ docker ps
CONTAINER ID        IMAGE                                        COMMAND                  CREATED             STATUS              PORTS                      NAMES
2abb962b5e0c        4c600a64a18a                                 "/opt/bin/flanneld..."   4 minutes ago       Up 4 minutes                                   k8s_kube-flannel_kube-flannel-ds-gk826_kube-system_a2450936-d3cf-11e7-b63a-024403e322ab_0
9bd664b94540        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 4 minutes ago       Up 4 minutes                                   k8s_POD_kube-flannel-ds-gk826_kube-system_a2450936-d3cf-11e7-b63a-024403e322ab_0
5c9919257799        512cd7425a73                                 "/kube-dns --domai..."   5 minutes ago       Up 5 minutes                                   k8s_kubedns_kube-dns-545bc4bfd4-mqq8j_kube-system_5d8f8aac-d3c3-11e7-9c69-080027eae602_6
825834b2329d        fed89e8b4248                                 "/sidecar --v=2 --..."   5 minutes ago       Up 5 minutes                                   k8s_sidecar_kube-dns-545bc4bfd4-mqq8j_kube-system_5d8f8aac-d3c3-11e7-9c69-080027eae602_5
e8a0ce312c72        459944ce8cc4                                 "/dnsmasq-nanny -v..."   5 minutes ago       Up 5 minutes                                   k8s_dnsmasq_kube-dns-545bc4bfd4-mqq8j_kube-system_5d8f8aac-d3c3-11e7-9c69-080027eae602_5
62c061da780c        65a61c14e8c2                                 "/usr/local/bin/ku..."   5 minutes ago       Up 5 minutes                                   k8s_kube-proxy_kube-proxy-4bmmw_kube-system_da0599c6-d3c2-11e7-9c69-080027eae602_4
62dabffc713c        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 5 minutes ago       Up 5 minutes                                   k8s_POD_kube-dns-545bc4bfd4-mqq8j_kube-system_5d8f8aac-d3c3-11e7-9c69-080027eae602_236
f903760a482a        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 5 minutes ago       Up 5 minutes                                   k8s_POD_kube-proxy-4bmmw_kube-system_da0599c6-d3c2-11e7-9c69-080027eae602_4
57b093409179        7058ac4d4af5                                 "kube-controller-m..."   5 minutes ago       Up 5 minutes                                   k8s_kube-controller-manager_kube-controller-manager-kubedev-10-64-33-195_kube-system_bf5427c07fa2cd55ef80aee53d9204e5_3
0a4f60e432fa        10a052dccbc5                                 "kube-apiserver --..."   5 minutes ago       Up 5 minutes                                   k8s_kube-apiserver_kube-apiserver-kubedev-10-64-33-195_kube-system_09ffa3ed2853c0d43bd972690ea51a62_3
13c17a33851b        0d985fed7f95                                 "kube-scheduler --..."   5 minutes ago       Up 5 minutes                                   k8s_kube-scheduler_kube-scheduler-kubedev-10-64-33-195_kube-system_ca97fd23ad8837acfa829af8dfc86a7e_3
b365344a1327        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 5 minutes ago       Up 5 minutes                                   k8s_POD_kube-apiserver-kubedev-10-64-33-195_kube-system_09ffa3ed2853c0d43bd972690ea51a62_4
8fed036058d5        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 5 minutes ago       Up 5 minutes                                   k8s_POD_kube-scheduler-kubedev-10-64-33-195_kube-system_ca97fd23ad8837acfa829af8dfc86a7e_3
3164d559f994        gcr.io/google_containers/pause-amd64:3.0     "/pause"                 5 minutes ago       Up 5 minutes                                   k8s_POD_kube-controller-manager-kubedev-10-64-33-195_kube-system_bf5427c07fa2cd55ef80aee53d9204e5_3
55ea2b550443        gcr.io/google_containers/etcd-amd64:3.0.17   "etcd --listen-cli..."   3 hours ago         Up 8 minutes                                   etcd
```
