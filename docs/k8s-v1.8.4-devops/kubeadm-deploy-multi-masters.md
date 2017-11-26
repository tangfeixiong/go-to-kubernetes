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
[vagrant@kubedev-10-64-33-82 ~]$ docker run -d --name etcd --net host -v /srv/etcd/data:/var/lib/etcd --restart always gcr.io/google_containers/etcd-amd64:3.0.17 etcd --listen-client-urls=http://10.64.33.82:2379,http://127.0.0.1:2379 --advertise-client-urls=http://10.64.33.82:2379 --name=etcd0 --initial-advertise-peer-urls=http://10.64.33.82:2380 --listen-peer-urls=http://10.64.33.82:2380 --initial-cluster-token=etcd-cluster-token --initial-cluster=etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380 -initial-cluster-state new
882adca68ea7006de9d13836f0bed105b4be731e9627d70c77d701cd36f8696d
```

```
[vagrant@kubedev-10-64-33-82 ~]$ docker ps -f name=etcd
CONTAINER ID        IMAGE                                        COMMAND                  CREATED             STATUS              PORTS               NAMES
a73207b5165d        gcr.io/google_containers/etcd-amd64:3.0.17   "etcd --listen-client"   37 seconds ago      Up 35 seconds                           etcd
```

Fedora
```
[vagrant@kubedev-10-64-33-199 ~]$ docker run -d --name etcd --net host -v /srv/etcd/data:/var/lib/etcd --restart always gcr.io/google_containers/etcd-amd64:3.0.17 etcd --listen-client-urls=http://10.64.33.199:2379,http://127.0.0.1:2379 --advertise-client-urls=http://10.64.33.199:2379 --name=etcd1 --initial-advertise-peer-urls=http://10.64.33.199:2380 --listen-peer-urls=http://10.64.33.199:2380 --initial-cluster-token=etcd-cluster-token --initial-cluster=etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380 -initial-cluster-state new
9fbfb9596f88e87127163a27bc5a5f1b3ddbe4e1e409bd7fe7b2a520c9667604
```

```
[vagrant@kubedev-10-64-33-199 ~]$ docker ps -f name=etcd
CONTAINER ID        IMAGE                                        COMMAND                  CREATED             STATUS              PORTS               NAMES
9fbfb9596f88        gcr.io/google_containers/etcd-amd64:3.0.17   "etcd --listen-client"   19 seconds ago      Up 17 seconds                           etcd
```

Ubuntu
```
ubuntu@kubedev-10-64-33-195:~$ docker run -d --name etcd --net host -v /srv/etcd/data:/var/lib/etcd --restart always gcr.io/google_containers/etcd-amd64:3.0.17 etcd --listen-client-urls=http://10.64.33.195:2379,http://127.0.0.1:2379 --advertise-client-urls=http://10.64.33.195:2379 --name=etcd2 --initial-advertise-peer-urls=http://10.64.33.195:2380 --listen-peer-urls=http://10.64.33.195:2380 --initial-cluster-token=etcd-cluster-token --initial-cluster=etcd0=http://10.64.33.82:2380,etcd1=http://10.64.33.199:2380,etcd2=http://10.64.33.195:2380 -initial-cluster-state new
40c8443c79c38e4f8e63b5bb1f6e5407c6675c394ebed5692fec13fad2e1251b
```

```
ubuntu@kubedev-10-64-33-195:~$ docker ps -f name=etcd
CONTAINER ID        IMAGE                                        COMMAND                  CREATED             STATUS              PORTS               NAMES
40c8443c79c3        gcr.io/google_containers/etcd-amd64:3.0.17   "etcd --listen-cli..."   19 seconds ago      Up 18 seconds                           etcd
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

### cluster

nodes
```
ubuntu@kubedev-10-64-33-195:~$ kubectl get nodes
NAME                   STATUS    ROLES     AGE       VERSION
kubedev-10-64-33-195   Ready     master    6h        v1.8.4
kubedev-10-64-33-199   Ready     master    6h        v1.8.4
kubedev-10-64-33-82    Ready     master    4m        v1.8.4
```



node 0
```
[vagrant@kubedev-10-64-33-82 ~]$ docker ps
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS                      NAMES
e092af6dc1ad        10a052dccbc5                               "kube-apiserver --req"   13 minutes ago      Up 12 minutes                                  k8s_kube-apiserver_kube-apiserver-kubedev-10-64-33-82_kube-system_43316dc04d0cdf0fb805e70c8e426dca_25
86bba53c8b38        243830dae7dd                               "etcd --data-dir=/var"   33 minutes ago      Up 33 minutes                                  k8s_etcd_etcd-kubedev-10-64-33-82_kube-system_06db8cec81dfe6c9bf7afc5511f3458f_3
8dc4253ebd29        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 39 minutes ago      Up 39 minutes                                  k8s_POD_etcd-kubedev-10-64-33-82_kube-system_06db8cec81dfe6c9bf7afc5511f3458f_0
64f238f87001        0d985fed7f95                               "kube-scheduler --lea"   About an hour ago   Up About an hour                               k8s_kube-scheduler_kube-scheduler-kubedev-10-64-33-82_kube-system_ca97fd23ad8837acfa829af8dfc86a7e_0
fb18aafa2440        7058ac4d4af5                               "kube-controller-mana"   About an hour ago   Up About an hour                               k8s_kube-controller-manager_kube-controller-manager-kubedev-10-64-33-82_kube-system_603280533b28a5ccd299bfade799f2ac_0
decb26b435a1        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 About an hour ago   Up About an hour                               k8s_POD_kube-apiserver-kubedev-10-64-33-82_kube-system_43316dc04d0cdf0fb805e70c8e426dca_0
4e24a3e43c9e        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 About an hour ago   Up About an hour                               k8s_POD_kube-scheduler-kubedev-10-64-33-82_kube-system_ca97fd23ad8837acfa829af8dfc86a7e_0
e23dcbbf85e2        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 About an hour ago   Up About an hour                               k8s_POD_kube-controller-manager-kubedev-10-64-33-82_kube-system_603280533b28a5ccd299bfade799f2ac_0
```

node 1
```
[vagrant@kubedev-10-64-33-199 ~]$ docker ps
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
96135334ba0f        10a052dccbc5                               "kube-apiserver --req"   22 minutes ago      Up 22 minutes                           k8s_kube-apiserver_kube-apiserver-kubedev-10-64-33-199_kube-system_43316dc04d0cdf0fb805e70c8e426dca_7
b4f09fe539cf        7058ac4d4af5                               "kube-controller-mana"   36 minutes ago      Up 36 minutes                           k8s_kube-controller-manager_kube-controller-manager-kubedev-10-64-33-199_kube-system_603280533b28a5ccd299bfade799f2ac_0
b41d1b9a2416        243830dae7dd                               "etcd --data-dir=/var"   36 minutes ago      Up 36 minutes                           k8s_etcd_etcd-kubedev-10-64-33-199_kube-system_1cee9d3646a16086edb3e60fb63052dd_0
dee95026f92a        0d985fed7f95                               "kube-scheduler --lea"   36 minutes ago      Up 36 minutes                           k8s_kube-scheduler_kube-scheduler-kubedev-10-64-33-199_kube-system_ca97fd23ad8837acfa829af8dfc86a7e_0
427ef08b3c8c        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 36 minutes ago      Up 36 minutes                           k8s_POD_kube-controller-manager-kubedev-10-64-33-199_kube-system_603280533b28a5ccd299bfade799f2ac_0
96baf05922d1        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 36 minutes ago      Up 36 minutes                           k8s_POD_kube-scheduler-kubedev-10-64-33-199_kube-system_ca97fd23ad8837acfa829af8dfc86a7e_0
b452bb14698b        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 36 minutes ago      Up 36 minutes                           k8s_POD_etcd-kubedev-10-64-33-199_kube-system_1cee9d3646a16086edb3e60fb63052dd_0
b0e3962ed45f        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 36 minutes ago      Up 36 minutes                           k8s_POD_kube-apiserver-kubedev-10-64-33-199_kube-system_43316dc04d0cdf0fb805e70c8e426dca_0
```

node 2
```
ubuntu@kubedev-10-64-33-195:~$ docker ps
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
44b227ca3ef8        10a052dccbc5                               "kube-apiserver --..."   17 minutes ago      Up 17 minutes                           k8s_kube-apiserver_kube-apiserver-kubedev-10-64-33-195_kube-system_43316dc04d0cdf0fb805e70c8e426dca_3
27a73ee38857        0d985fed7f95                               "kube-scheduler --..."   17 minutes ago      Up 17 minutes                           k8s_kube-scheduler_kube-scheduler-kubedev-10-64-33-195_kube-system_ca97fd23ad8837acfa829af8dfc86a7e_1
4ec47883339a        243830dae7dd                               "etcd --data-dir=/..."   17 minutes ago      Up 17 minutes                           k8s_etcd_etcd-kubedev-10-64-33-195_kube-system_d16c77b60609a7a41f85ca5cd57b8a42_2
cabc861b35e2        7058ac4d4af5                               "kube-controller-m..."   17 minutes ago      Up 17 minutes                           k8s_kube-controller-manager_kube-controller-manager-kubedev-10-64-33-195_kube-system_603280533b28a5ccd299bfade799f2ac_1
42ab44b64e18        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 17 minutes ago      Up 17 minutes                           k8s_POD_kube-scheduler-kubedev-10-64-33-195_kube-system_ca97fd23ad8837acfa829af8dfc86a7e_1
96ff0d4837d6        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 17 minutes ago      Up 17 minutes                           k8s_POD_kube-controller-manager-kubedev-10-64-33-195_kube-system_603280533b28a5ccd299bfade799f2ac_1
19c2ff54e4ba        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 17 minutes ago      Up 17 minutes                           k8s_POD_kube-apiserver-kubedev-10-64-33-195_kube-system_43316dc04d0cdf0fb805e70c8e426dca_1
3861d6c083f0        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 17 minutes ago      Up 17 minutes                           k8s_POD_etcd-kubedev-10-64-33-195_kube-system_d16c77b60609a7a41f85ca5cd57b8a42_1
```
