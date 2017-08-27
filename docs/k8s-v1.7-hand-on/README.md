# Instruction

## Abstract

Version v1.7.4

### Feature

* Single Node
* Systemd `kubelet`
* Container 'etcd, kube-apiserver, kube-controller-manager, kube-scheduler, kube-proxy'

### Environment

* Fedora23 openshift virtualbox
* Docker 1.10.3

## PoC

### Prerequisites

Download server package
```
fanhonglingdeMacBook-Pro:v1.7 fanhongling$ curl -jkSL https://dl.k8s.io/v1.7.4/kubernetes-server-linux-amd64.tar.gz -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0    110      0  0:00:01  0:00:01 --:--:--   110
100  417M  100  417M    0     0  3889k      0  0:01:49  0:01:49 --:--:-- 4259k
```

Download confiuration management package
```
fanhonglingdeMacBook-Pro:v1.7.4 fanhongling$ curl -jkSLO https://dl.k8s.io/v1.7.4/kubernetes.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     23      0  0:00:07  0:00:06  0:00:01    48
100 3833k  100 3833k    0     0   377k      0  0:00:10  0:00:10 --:--:-- 1508k
```

Download client package
```
fanhonglingdeMacBook-Pro:v1.7.4 fanhongling$ curl -jkSLO https://dl.k8s.io/v1.7.4/kubernetes-client-darwin-amd64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     51      0  0:00:03  0:00:03 --:--:--    51
100 31.5M  100 31.5M    0     0  1771k      0  0:00:18  0:00:18 --:--:--  522k
fanhonglingdeMacBook-Pro:v1.7.4 fanhongling$ curl -jkSLO https://dl.k8s.io/v1.7.4/kubernetes-client-windows-amd64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0    101      0  0:00:01  0:00:01 --:--:--   101
100 31.7M  100 31.7M    0     0  1177k      0  0:00:27  0:00:27 --:--:-- 2216k
```

### Extract binaries and images and configuration base

The binaries and images
```
[vagrant@localhost ~]$ tar -ztf /Users/fanhongling/Downloads/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/v1.7.4/kubernetes-server-linux-amd64.tar.gz 
kubernetes/
kubernetes/server/
kubernetes/server/bin/
kubernetes/server/bin/cloud-controller-manager
kubernetes/server/bin/kube-aggregator.tar
kubernetes/server/bin/kube-proxy.tar
kubernetes/server/bin/kube-proxy
kubernetes/server/bin/kube-controller-manager.tar
kubernetes/server/bin/kube-controller-manager
kubernetes/server/bin/kube-apiserver
kubernetes/server/bin/kube-aggregator.docker_tag
kubernetes/server/bin/kube-controller-manager.docker_tag
kubernetes/server/bin/kubefed
kubernetes/server/bin/kube-scheduler.tar
kubernetes/server/bin/kube-apiserver.tar
kubernetes/server/bin/kubeadm
kubernetes/server/bin/kube-scheduler.docker_tag
kubernetes/server/bin/hyperkube
kubernetes/server/bin/kube-scheduler
kubernetes/server/bin/cloud-controller-manager.tar
kubernetes/server/bin/kubelet
kubernetes/server/bin/kube-proxy.docker_tag
kubernetes/server/bin/kube-apiserver.docker_tag
kubernetes/server/bin/kubectl
kubernetes/server/bin/apiextensions-apiserver
kubernetes/server/bin/cloud-controller-manager.docker_tag
kubernetes/server/bin/kube-aggregator
kubernetes/LICENSES
kubernetes/addons/
kubernetes/kubernetes-src.tar.gz
```

Destination
```
[vagrant@localhost ~]$ mkdir -p /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4
```

Extract
```
[vagrant@localhost ~]$ tar -C /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4 -zxvf /Users/fanhongling/Downloads/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/v1.7.4/kubernetes-server-linux-amd64.tar.gz 
kubernetes/
kubernetes/server/
kubernetes/server/bin/
kubernetes/server/bin/cloud-controller-manager
kubernetes/server/bin/kube-aggregator.tar
kubernetes/server/bin/kube-proxy.tar
kubernetes/server/bin/kube-proxy
kubernetes/server/bin/kube-controller-manager.tar
kubernetes/server/bin/kube-controller-manager
kubernetes/server/bin/kube-apiserver
kubernetes/server/bin/kube-aggregator.docker_tag
kubernetes/server/bin/kube-controller-manager.docker_tag
kubernetes/server/bin/kubefed
kubernetes/server/bin/kube-scheduler.tar
kubernetes/server/bin/kube-apiserver.tar
kubernetes/server/bin/kubeadm
kubernetes/server/bin/kube-scheduler.docker_tag
kubernetes/server/bin/hyperkube
kubernetes/server/bin/kube-scheduler
kubernetes/server/bin/cloud-controller-manager.tar
kubernetes/server/bin/kubelet
kubernetes/server/bin/kube-proxy.docker_tag
kubernetes/server/bin/kube-apiserver.docker_tag
kubernetes/server/bin/kubectl
kubernetes/server/bin/apiextensions-apiserver
kubernetes/server/bin/cloud-controller-manager.docker_tag
kubernetes/server/bin/kube-aggregator
kubernetes/LICENSES
kubernetes/addons/
kubernetes/kubernetes-src.tar.gz
```

For convenient way
```
[vagrant@localhost ~]$ sudo ln -sf /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/kubernetes /opt/kubernetes
```

The binaries
```
[vagrant@localhost ~]$ ls /opt/kubernetes/server/bin/
apiextensions-apiserver              kubeadm                     kube-apiserver.docker_tag           kubectl                kube-proxy.tar
cloud-controller-manager             kube-aggregator             kube-apiserver.tar                  kubefed                kube-scheduler
cloud-controller-manager.docker_tag  kube-aggregator.docker_tag  kube-controller-manager             kubelet                kube-scheduler.docker_tag
cloud-controller-manager.tar         kube-aggregator.tar         kube-controller-manager.docker_tag  kube-proxy             kube-scheduler.tar
hyperkube                            kube-apiserver              kube-controller-manager.tar         kube-proxy.docker_tag
```

The images
```
[vagrant@localhost ~]$ for i in $(ls /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/kubernetes/server/bin/*.docker_tag); do echo $(basename $i .docker_tag); done
cloud-controller-manager
kube-aggregator
kube-apiserver
kube-controller-manager
kube-proxy
kube-scheduler
```

Load
```
[vagrant@localhost ~]$ bindir=/Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/kubernetes/server/bin; for i in $(ls $bindir/*.tar); do docker load -i $i; name=$(basename $i .tar); tag=$(cat $bindir/$name.docker_tag); repo=gcr.io/google_containers/$name; img=$(docker images -q $repo:$tag); [[ -n $img ]] && docker images $repo:$tag || (echo "...failed to load" && docker images $repo); done
REPOSITORY                                          TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/cloud-controller-manager   v1.7.4              a42a0fd85571        9 days ago          125.3 MB
REPOSITORY                                 TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-aggregator   v1.7.4              c1752dd09c50        9 days ago          59.16 MB
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver   v1.7.4              5260ecb5129c        9 days ago          186.1 MB
REPOSITORY                                         TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-controller-manager   v1.7.4              d2adddc4b1cb        9 days ago          138 MB
REPOSITORY                            TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-proxy   v1.7.4              0f3bf654ec61        9 days ago          114.7 MB
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-scheduler   v1.7.4              b1cd468ba656        9 days ago          77.2 MB
```

View
```
[vagrant@localhost ~]$ docker images gcr.io/google_containers/*
REPOSITORY                                             TAG                                IMAGE ID            CREATED             SIZE
gcr.io/google_containers/cloud-controller-manager      v1.7.4                             a42a0fd85571        9 days ago          125.3 MB
gcr.io/google_containers/kube-controller-manager       v1.7.4                             d2adddc4b1cb        9 days ago          138 MB
gcr.io/google_containers/kube-apiserver                v1.7.4                             5260ecb5129c        9 days ago          186.1 MB
gcr.io/google_containers/kube-aggregator               v1.7.4                             c1752dd09c50        9 days ago          59.16 MB
gcr.io/google_containers/kube-proxy                    v1.7.4                             0f3bf654ec61        9 days ago          114.7 MB
gcr.io/google_containers/kube-scheduler                v1.7.4                             b1cd468ba656        9 days ago          77.2 MB
gcr.io/google_containers/nginx-slim-amd64              0.21                               fe6b90978d65        6 weeks ago         95.31 MB
gcr.io/google_containers/kube-apiserver                41a3581e4847878bbcb1722636ab51ad   9196766fca5a        3 months ago        150.6 MB
gcr.io/google_containers/kube-proxy                    4c6ae16e62496a67d3e9cc622c1006a3   aba0fdd437d0        3 months ago        109.2 MB
gcr.io/google_containers/kube-controller-manager       6b1e075eaab8773444890c5cc17e0dcf   3c2198b97c22        3 months ago        132.8 MB
gcr.io/google_containers/kube-scheduler                c8ed221003bb194f37ef4221727bce1c   36723c5b4c70        3 months ago        76.75 MB
gcr.io/google_containers/kubernetes-dashboard-amd64    v1.6.0                             416701f962f2        5 months ago        108.6 MB
gcr.io/google_containers/k8s-dns-sidecar-amd64         1.14.1                             fc5e302d8309        5 months ago        44.52 MB
gcr.io/google_containers/k8s-dns-kube-dns-amd64        1.14.1                             f8363dbf447b        5 months ago        52.36 MB
gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64   1.14.1                             1091847716ec        5 months ago        44.84 MB
gcr.io/google_containers/etcd                          3.0.17                             e6a6589433f4        6 months ago        168.9 MB
gcr.io/google_containers/kubernetes-dashboard-amd64    v1.4.2                             c0e4ba8968ee        9 months ago        89.49 MB
gcr.io/google_containers/kubedns-amd64                 1.8                                4fc1ff90ae46        11 months ago       57.89 MB
gcr.io/google_containers/kube-dnsmasq-amd64            1.4                                3ec65756a89b        11 months ago       5.126 MB
gcr.io/google_containers/kubernetes-dashboard-amd64    v1.4.0                             436faaeba2e2        11 months ago       86.27 MB
gcr.io/google_containers/kubernetes-dashboard-amd64    v1.1.1                             f739d2414b14        12 months ago       55.83 MB
gcr.io/google_containers/heapster                      v1.1.0                             1769fa9dfaca        14 months ago       121.9 MB
gcr.io/google_containers/kube-dnsmasq-amd64            1.3                                9a15e39d0db8        14 months ago       5.126 MB
gcr.io/google_containers/pause-amd64                   3.0                                99e59f495ffa        15 months ago       746.9 kB
gcr.io/google_containers/exechealthz-amd64             1.0                                82a141f5d06d        17 months ago       7.116 MB
gcr.io/google_containers/elasticsearch                 2.2.0                              8d457c1db41b        18 months ago       346.2 MB
gcr.io/google_containers/pause                         2.0                                2b58359142b0        22 months ago       350.2 kB
gcr.io/google_containers/pause                         0.8.0                              bf595365a558        2 years ago         241.7 kB
```

Configuration Management base package
```
[vagrant@localhost ~]$ tar -C /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/ -zxvf /Users/fanhongling/Downloads/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/v1.7.4/kubernetes.tar.gz kubernetes/server/kubernetes-salt.tar.gz
kubernetes/server/kubernetes-salt.tar.gz
```

Destination _cluster_
```
[vagrant@localhost ~]$ mkdir -p /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/kubernetes/cluster
```

Extract _saltbase_
```
[vagrant@localhost ~]$ tar -C /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/kubernetes/cluster -zxf /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/kubernetes/server/kubernetes-salt.tar.gz --strip-components=1
```

The _saltbase_ directory
```
[vagrant@localhost ~]$ ls /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/kubernetes/cluster/
saltbase
[vagrant@localhost ~]$ ls /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/kubernetes/cluster/saltbase/
BUILD  install.sh  pillar  reactor  README.md  salt
```

### Install

For __ETCD__, Refer to [v1.5.7 hand-on](../k8s-v1.5.7-deployment)
```
[vagrant@localhost ~]$ tag=$(cat /opt/kubernetes/cluster/saltbase/salt/etcd/etcd.manifest | grep "gcr.io/google_containers/etcd:" | sed "s/.*'etcd_docker_tag', '\([0-9.]\+\)'.*/\1/");echo $tag
3.0.17
```

The `etcd` image
```
[vagrant@localhost ~]$ docker images gcr.io/google_containers/etcd:3.0.17
REPOSITORY                      TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/etcd   3.0.17              e6a6589433f4        6 months ago        168.9 MB
```

For __POD__ image
```
[vagrant@localhost ~]$ docker images gcr.io/google_containers/pause-amd64:3.0
REPOSITORY                             TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/pause-amd64   3.0                 99e59f495ffa        15 months ago       746.9 kB
```

Make sure to create __manifests__ and __addons__ directory
```
[vagrant@localhost ~]$ ls /etc/kubernetes/
addons  manifests
```

### X509 cacerts

Refert [v1.6.4 hand-on](../k8s-v1.6.4-deployment)
```
[vagrant@localhost ~]$ ls /srv/kubernetes/
ca.crt  etcd-ca.crt  etcd-peer.cert  etcd-peer.key  kubeapiserver.cert  kubeapiserver.key  kubecfg.crt  kubecfg.key  server.cert  server.key
```

Server certificate
```
[vagrant@localhost ~]$ openssl x509 -in /srv/kubernetes/server.cert -text -noout
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 1 (0x1)
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN=172.17.4.50@1462900056
        Validity
            Not Before: May 10 17:07:36 2016 GMT
            Not After : May  8 17:07:36 2026 GMT
        Subject: CN=kubernetes-master
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:c0:dc:a2:20:38:75:e9:e9:21:72:4a:d4:a0:91:
                    6c:28:34:bf:e2:a7:d3:54:ca:a4:8c:bd:bc:81:14:
                    6b:2f:91:e5:83:68:bf:81:0e:4e:42:08:ff:87:32:
                    ef:67:d8:78:51:55:df:e6:dd:62:d5:47:c2:e4:fb:
                    d5:65:f3:f6:e9:21:a4:b7:82:81:59:3c:b3:90:ad:
                    65:e4:b9:42:14:76:e5:cd:6d:24:fd:55:c5:92:60:
                    b0:22:50:5e:e8:3b:2d:30:6f:1a:7e:eb:41:e0:9b:
                    e0:0f:6e:e5:53:03:bf:64:6c:10:a1:c0:7a:05:78:
                    3a:e6:d3:5b:2d:50:32:92:5e:80:1a:07:43:8a:ed:
                    2b:a9:19:4b:1c:e1:00:11:c7:e6:58:27:d7:a7:8b:
                    a7:b8:87:df:f2:f3:bf:17:f1:ab:b4:f0:e3:7e:65:
                    cf:58:d0:9b:47:53:e6:30:f1:44:84:d0:08:9b:79:
                    e8:02:33:e3:3c:89:af:8d:7c:38:d9:97:ff:31:ac:
                    38:eb:dd:11:ff:8c:df:22:35:6e:cc:37:2d:56:4e:
                    30:dd:f0:9f:96:89:c2:2e:be:c5:5c:ba:06:20:02:
                    ec:ca:53:0c:c0:9d:e6:13:93:2b:de:f5:04:1a:70:
                    b0:3f:9c:bd:4c:13:65:be:83:fb:62:55:c2:aa:f1:
                    8e:2d
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: 
                CA:FALSE
            X509v3 Subject Key Identifier: 
                15:55:EA:81:AF:62:3B:4D:90:9F:BF:5D:CF:FC:1A:E4:99:30:68:13
            X509v3 Authority Key Identifier: 
                keyid:F9:E3:13:4D:74:3A:E7:15:94:ED:F2:15:69:EB:49:E4:2E:21:DF:7B
                DirName:/CN=172.17.4.50@1462900056
                serial:B3:5F:07:24:3D:C7:0C:BA

            X509v3 Extended Key Usage: 
                TLS Web Server Authentication
            X509v3 Key Usage: 
                Digital Signature, Key Encipherment
            X509v3 Subject Alternative Name: 
                IP Address:172.17.4.50, IP Address:172.17.4.50, IP Address:10.3.0.1, DNS:kubernetes, DNS:kubernetes.default, DNS:kubernetens.default.svc, DNS:kubernetes.default.svc.cluster.local
    Signature Algorithm: sha256WithRSAEncryption
         07:bc:83:47:9c:e3:0d:ac:30:d8:eb:87:6f:87:66:72:22:a2:
         09:78:ae:b9:91:07:f9:fb:95:ea:98:db:1e:cf:22:6a:5d:f1:
         c7:c2:66:1f:11:48:d2:eb:22:0d:9a:f5:b1:8a:31:c6:16:64:
         37:2e:5b:f9:a1:49:69:ae:28:a6:7a:58:ea:f8:34:87:6a:b7:
         41:14:7b:be:d2:9f:5d:6b:e4:67:17:d7:70:7d:4d:5d:d4:35:
         96:61:c3:f2:18:bb:50:55:50:b8:dd:f6:7d:cb:4f:81:bb:47:
         d0:68:1c:be:bf:47:68:7f:ec:a5:2a:37:b9:6f:f9:8c:42:9f:
         9d:14:26:e3:5d:3d:6f:54:b2:2f:20:ef:42:00:d8:3f:2b:8a:
         28:28:bf:1f:d8:23:bf:23:0b:7a:0a:7b:86:14:7a:ad:cd:eb:
         fa:4b:74:3e:27:a0:3d:8e:4d:2c:bc:ec:bd:00:aa:4f:87:da:
         8d:56:0a:f3:cc:62:3f:04:c6:47:f3:90:7c:ca:1d:7f:3e:50:
         59:1c:d0:32:b1:86:4f:28:59:15:ec:c8:6f:8a:0d:ec:02:d7:
         80:e3:19:58:e9:03:ce:a3:16:ed:1a:8d:12:4e:75:e2:2d:46:
         a1:3e:82:7f:b8:76:1a:15:37:b0:98:da:db:a9:92:f7:cb:c8:
         f9:68:c4:1b
```

### Daemon

The `kubelet` systemd service, refer to [v1.6.4 hand-on](../k8s-v1.6-deployment)
```
[vagrant@localhost ~]$ sed -e 's%\(\(ExecStart=\)/usr/local/bin/kubelet\( "$DAEMON_ARGS"\)\)%# \1\n\2/opt/kubernetes/server/bin/kubelet\3%' /opt/kubernetes/cluster/saltbase/salt/kubelet/kubelet.service
[Unit]
Description=Kubernetes Kubelet Server
Documentation=https://github.com/kubernetes/kubernetes

[Service]
EnvironmentFile=/etc/sysconfig/kubelet
# ExecStart=/usr/local/bin/kubelet "$DAEMON_ARGS"
ExecStart=/opt/kubernetes/server/bin/kubelet "$DAEMON_ARGS"
Restart=always
RestartSec=2s
StartLimitInterval=0
KillMode=process

[Install]
WantedBy=multi-user.target
```

Networking
```
[vagrant@localhost ~]$ ip addr show dev eth1
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:5a:1a:a4 brd ff:ff:ff:ff:ff:ff
    inet 172.17.4.50/24 brd 172.17.4.255 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe5a:1aa4/64 scope link 
       valid_lft forever preferred_lft forever
```

Environment
```
[vagrant@localhost ~]$ HOSTNAME_OVERRIDE=$(hostname -I | awk '{print $2}')
[vagrant@localhost ~]$ echo $HOSTNAME_OVERRIDE
172.17.4.50
```

`kubelet` args
```
[vagrant@localhost ~]$ sed '/^# test_args.*$/,$!d;s/{{daemon_args}}//;s%{{api_servers}}%--api-servers=http://127.0.0.1:8080%;s/{{debugging_handlers}}/--enable-debugging-handlers=true/;s/{{hostname_override}}/--hostname-override=${KUBE_NODE_IP}/;s/{{cloud_provider}}//;s/{{cloud_config}}//;s%{{config}}%--pod-manifest-path=/etc/kubernetes/manifests%;s%{{manifest_url}}%%;s/{{pillar\[\x27allow_privileged\x27\]}}/true/;s/{{log_level}}/--v=2/;s/{{cluster_dns}}/--cluster-dns=10.3.0.10/;s/{{cluster_domain}}/--cluster-domain=cluster.local/;s%{{docker_root}}%%;s%{{kubelet_root}}%--root-dir=/var/lib/kubelet%;s%{{non_masquerade_cidr}}%--non-masquerade-cidr=10.0.0.0/8%;s%{{cgroup_root}}%%;s%{{system_container}}%%;s%{{pod_cidr}}%--pod-cidr=10.0.0.0/14%;s%{{ master_kubelet_args }}%--address=0.0.0.0 --authorization-mode=AlwaysAllow --node-ip=${KUBE_NODE_IP} --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0%;s/{{cpu_cfs_quota}}/--cpu-cfs-quota=true/;s/{{network_plugin}}//;s/{{kubelet_port}}/--port=10250/;s/{{ hairpin_mode }}//;s/{{enable_custom_metrics}}/--enable-custom-metrics=true/;s%{{runtime_container}}%%;s%{{kubelet_container}}%%;s/{{node_labels}}//;s/{{node_taints}}//;s/{{babysit_daemons}}//;s/{{eviction_hard}}/--eviction-hard=memory.available<100Mi/;s%{{kubelet_auth}}%--anonymous-auth=false --authorization-mode=Webhook --client-ca-file=/var/lib/kubelet/ca.crt%;s%{{pki}}%--cert-dir=/var/lib/kubelet/pki%;s/{{feature_gates}}/--feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true/;s/{{test_args}}//' /opt/kubernetes/cluster/saltbase/salt/kubelet/default | env KUBE_NODE_IP=$HOSTNAME_OVERRIDE envsubst | sudo tee /etc/sysconfig/kubelet
# test_args has to be kept at the end, so they'll overwrite any prior configuration
DAEMON_ARGS=" --api-servers=http://127.0.0.1:8080 --enable-debugging-handlers=true --hostname-override=172.17.4.50   --pod-manifest-path=/etc/kubernetes/manifests  --allow-privileged=true --v=2 --cluster-dns=10.3.0.10 --cluster-domain=cluster.local  --root-dir=/var/lib/kubelet  --non-masquerade-cidr=10.0.0.0/8   --pod-cidr=10.0.0.0/14 --address=0.0.0.0 --authorization-mode=AlwaysAllow --node-ip=172.17.4.50 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --cpu-cfs-quota=true  --port=10250  --enable-custom-metrics=true     --eviction-hard=memory.available<100Mi --anonymous-auth=false --authorization-mode=Webhook --client-ca-file=/var/lib/kubelet/ca.crt --cert-dir=/var/lib/kubelet/pki --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true "
```

Fix
```
sed '/^# test_args.*$/,$!d;s/{{daemon_args}}//;s%{{api_servers}}%--api-servers=http://127.0.0.1:8080%;s/{{debugging_handlers}}/--enable-debugging-handlers=true/;s/{{hostname_override}}/--hostname-override=${KUBE_NODE_IP}/;s/{{cloud_provider}}//;s/{{cloud_config}}//;s%{{config}}%--pod-manifest-path=/etc/kubernetes/manifests%;s%{{manifest_url}}%%;s/{{pillar\[\x27allow_privileged\x27\]}}/true/;s/{{log_level}}/--v=2/;s/{{cluster_dns}}/--cluster-dns=10.3.0.10/;s/{{cluster_domain}}/--cluster-domain=cluster.local/;s%{{docker_root}}%%;s%{{kubelet_root}}%--root-dir=/var/lib/kubelet%;s%{{non_masquerade_cidr}}%--non-masquerade-cidr=10.0.0.0/8%;s%{{cgroup_root}}%%;s%{{system_container}}%%;s%{{pod_cidr}}%--pod-cidr=10.0.0.0/14%;s%{{ master_kubelet_args }}%--address=0.0.0.0 --authorization-mode=AlwaysAllow --node-ip=${KUBE_NODE_IP} --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0%;s/{{cpu_cfs_quota}}/--cpu-cfs-quota=true/;s/{{network_plugin}}//;s/{{kubelet_port}}/--port=10250/;s/{{ hairpin_mode }}//;s/{{enable_custom_metrics}}/--enable-custom-metrics=true/;s%{{runtime_container}}%%;s%{{kubelet_container}}%%;s/{{node_labels}}//;s/{{node_taints}}//;s/{{babysit_daemons}}//;s/{{eviction_hard}}/--eviction-hard=memory.available<100Mi/;s%{{kubelet_auth}}%--anonymous-auth=true --client-ca-file=/var/lib/kubelet/ca.crt%;s%{{pki}}%--cert-dir=/var/lib/kubelet/pki%;s/{{feature_gates}}/--feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=false,DynamicKubeletConfig=false,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=false,RotateKubeletServerCertificate=false,TaintBasedEvictions=true/;s/{{test_args}}//' /opt/kubernetes/cluster/saltbase/salt/kubelet/default | env KUBE_NODE_IP=$HOSTNAME_OVERRIDE envsubst | sudo tee /etc/sysconfig/kubelet
```

cacerts
```
[vagrant@localhost ~]$ sudo cp /srv/kubernetes/ca.crt /var/lib/kubelet
[vagrant@localhost ~]$ sudo mkdir -p /var/lib/kubelet/pki
# [vagrant@localhost ~]$ sudo cp /srv/kubernetes/server.cert /var/lib/kubelet/pki/kubelet.cert
# [vagrant@localhost ~]$ sudo cp /srv/kubernetes/server.key /var/lib/kubelet/pki/kubelet.key
```

Restart
```
[vagrant@localhost ~]$ sudo systemctl restart kubelet.service
```

Status
```
[vagrant@localhost ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - Kubernetes Kubelet Server
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
   Active: active (running) since 日 2017-08-27 03:20:44 UTC; 8min ago
     Docs: https://github.com/kubernetes/kubernetes
 Main PID: 21587 (kubelet)
   CGroup: /system.slice/kubelet.service
           ├─14977 journalctl -k -f
           └─21587 /opt/kubernetes/server/bin/kubelet --api-servers=http://127.0.0.1:8080 --enable-debugging-handlers=true --hostname-override=172.17.4.50 --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --v=2 --cluster-dns=10.3.0.10 --cluster-domain=cluster.local --root-dir=/var/lib/kubelet --non-masquerade-cidr=10.0.0.0/8 --pod-cidr=10.0.0.0/14 --address=0.0.0.0 --authorization-mode=AlwaysAllow --node-ip=172.17.4.50 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --cpu-cfs-quota=true --port=10250 --enable-custom-metrics=true --eviction-hard=memory.available<100Mi --anonymous-auth=false --authorization-mode=Webhook --client-ca-file=/var/lib/kubelet/ca.crt --cert-dir=/var/lib/kubelet/pki --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true

8月 27 03:27:15 localhost.localdomain kubelet[21587]: I0827 03:27:15.167351   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:27:45 localhost.localdomain kubelet[21587]: I0827 03:27:45.159757   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:27:45 localhost.localdomain kubelet[21587]: W0827 03:27:45.159926   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:27:45 localhost.localdomain kubelet[21587]: I0827 03:27:45.163499   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:28:15 localhost.localdomain kubelet[21587]: I0827 03:28:15.160523   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:28:15 localhost.localdomain kubelet[21587]: W0827 03:28:15.160607   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:28:15 localhost.localdomain kubelet[21587]: I0827 03:28:15.164310   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:28:45 localhost.localdomain kubelet[21587]: I0827 03:28:45.159965   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:28:45 localhost.localdomain kubelet[21587]: W0827 03:28:45.160093   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:28:45 localhost.localdomain kubelet[21587]: I0827 03:28:45.172084   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
```

Journal
```
[vagrant@localhost ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service
8月 27 03:20:44 localhost.localdomain systemd[1]: Stopping Kubernetes Kubelet Server...
8月 27 03:20:44 localhost.localdomain kubelet[14891]: I0827 03:20:44.082290   14891 docker_server.go:73] Stop docker server
8月 27 03:20:44 localhost.localdomain systemd[1]: kubelet.service: Main process exited, code=exited, status=1/FAILURE
8月 27 03:20:44 localhost.localdomain systemd[1]: kubelet.service: Unit entered failed state.
8月 27 03:20:44 localhost.localdomain systemd[1]: kubelet.service: Failed with result 'exit-code'.
8月 27 03:20:44 localhost.localdomain systemd[1]: Started Kubernetes Kubelet Server.
8月 27 03:20:44 localhost.localdomain systemd[1]: Starting Kubernetes Kubelet Server...
8月 27 03:20:45 localhost.localdomain kubelet[21587]: Flag --api-servers has been deprecated, Use --kubeconfig instead. Will be removed in a future version.
8月 27 03:20:45 localhost.localdomain kubelet[21587]: Flag --non-masquerade-cidr has been deprecated, will be removed in a future version
8月 27 03:20:45 localhost.localdomain kubelet[21587]: I0827 03:20:45.114727   21587 feature_gate.go:144] feature gates: map[LocalStorageCapacityIsolation:true DynamicVolumeProvisioning:true RotateKubeletClientCertificate:true ExperimentalCriticalPodAnnotation:true PersistentLocalVolumes:true TaintBasedEvictions:true RotateKubeletServerCertificate:true DynamicKubeletConfig:true ExperimentalHostUserNamespaceDefaulting:true Accelerators:true AdvancedAuditing:true AffinityInAnnotations:true AllAlpha:true]
8月 27 03:20:45 localhost.localdomain kubelet[21587]: W0827 03:20:45.120906   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:20:45 localhost.localdomain kubelet[21587]: I0827 03:20:45.157907   21587 server.go:239] Starting Kubelet configuration sync loop
8月 27 03:20:45 localhost.localdomain kubelet[21587]: E0827 03:20:45.157953   21587 server.go:412] failed to init dynamic Kubelet configuration sync: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:20:45 localhost.localdomain kubelet[21587]: W0827 03:20:45.158282   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:20:45 localhost.localdomain kubelet[21587]: I0827 03:20:45.158972   21587 server.go:473] Starting client certificate rotation.
8月 27 03:20:45 localhost.localdomain kubelet[21587]: I0827 03:20:45.159921   21587 certificate_manager.go:192] Certificate rotation is enabled.
8月 27 03:20:45 localhost.localdomain kubelet[21587]: I0827 03:20:45.193429   21587 certificate_manager.go:355] Requesting new certificate.
8月 27 03:21:15 localhost.localdomain kubelet[21587]: I0827 03:21:15.185059   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:21:15 localhost.localdomain kubelet[21587]: W0827 03:21:15.188586   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:21:15 localhost.localdomain kubelet[21587]: I0827 03:21:15.191302   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:21:45 localhost.localdomain kubelet[21587]: I0827 03:21:45.180464   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:21:45 localhost.localdomain kubelet[21587]: W0827 03:21:45.180555   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:21:45 localhost.localdomain kubelet[21587]: I0827 03:21:45.183243   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:22:15 localhost.localdomain kubelet[21587]: I0827 03:22:15.159938   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:22:15 localhost.localdomain kubelet[21587]: W0827 03:22:15.160062   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:22:15 localhost.localdomain kubelet[21587]: I0827 03:22:15.163066   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:22:45 localhost.localdomain kubelet[21587]: I0827 03:22:45.160014   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:22:45 localhost.localdomain kubelet[21587]: W0827 03:22:45.160092   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:22:45 localhost.localdomain kubelet[21587]: I0827 03:22:45.164482   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:23:15 localhost.localdomain kubelet[21587]: I0827 03:23:15.161109   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:23:15 localhost.localdomain kubelet[21587]: W0827 03:23:15.161283   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:23:15 localhost.localdomain kubelet[21587]: I0827 03:23:15.164913   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:23:45 localhost.localdomain kubelet[21587]: I0827 03:23:45.159937   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:23:45 localhost.localdomain kubelet[21587]: W0827 03:23:45.160017   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:23:45 localhost.localdomain kubelet[21587]: I0827 03:23:45.166390   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:24:15 localhost.localdomain kubelet[21587]: I0827 03:24:15.163108   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:24:15 localhost.localdomain kubelet[21587]: W0827 03:24:15.163198   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:24:15 localhost.localdomain kubelet[21587]: I0827 03:24:15.166476   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:24:45 localhost.localdomain kubelet[21587]: I0827 03:24:45.159894   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:24:45 localhost.localdomain kubelet[21587]: W0827 03:24:45.159957   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:24:45 localhost.localdomain kubelet[21587]: I0827 03:24:45.162512   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:25:15 localhost.localdomain kubelet[21587]: I0827 03:25:15.159923   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:25:15 localhost.localdomain kubelet[21587]: W0827 03:25:15.160064   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:25:15 localhost.localdomain kubelet[21587]: I0827 03:25:15.164058   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:25:45 localhost.localdomain kubelet[21587]: I0827 03:25:45.160755   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:25:45 localhost.localdomain kubelet[21587]: W0827 03:25:45.160861   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:25:45 localhost.localdomain kubelet[21587]: I0827 03:25:45.163637   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:26:15 localhost.localdomain kubelet[21587]: I0827 03:26:15.161201   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:26:15 localhost.localdomain kubelet[21587]: W0827 03:26:15.161364   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:26:15 localhost.localdomain kubelet[21587]: I0827 03:26:15.164142   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:26:45 localhost.localdomain kubelet[21587]: I0827 03:26:45.160645   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:26:45 localhost.localdomain kubelet[21587]: W0827 03:26:45.160814   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:26:45 localhost.localdomain kubelet[21587]: I0827 03:26:45.164353   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:27:15 localhost.localdomain kubelet[21587]: I0827 03:27:15.160760   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:27:15 localhost.localdomain kubelet[21587]: W0827 03:27:15.161033   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:27:15 localhost.localdomain kubelet[21587]: I0827 03:27:15.167351   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:27:45 localhost.localdomain kubelet[21587]: I0827 03:27:45.159757   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:27:45 localhost.localdomain kubelet[21587]: W0827 03:27:45.159926   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:27:45 localhost.localdomain kubelet[21587]: I0827 03:27:45.163499   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:28:15 localhost.localdomain kubelet[21587]: I0827 03:28:15.160523   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:28:15 localhost.localdomain kubelet[21587]: W0827 03:28:15.160607   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:28:15 localhost.localdomain kubelet[21587]: I0827 03:28:15.164310   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:28:45 localhost.localdomain kubelet[21587]: I0827 03:28:45.159965   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:28:45 localhost.localdomain kubelet[21587]: W0827 03:28:45.160093   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:28:45 localhost.localdomain kubelet[21587]: I0827 03:28:45.172084   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:29:15 localhost.localdomain kubelet[21587]: I0827 03:29:15.160757   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:29:15 localhost.localdomain kubelet[21587]: W0827 03:29:15.160907   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:29:15 localhost.localdomain kubelet[21587]: I0827 03:29:15.164208   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
8月 27 03:29:45 localhost.localdomain kubelet[21587]: I0827 03:29:45.160670   21587 server.go:242] Checking API server for new Kubelet configuration.
8月 27 03:29:45 localhost.localdomain kubelet[21587]: W0827 03:29:45.160759   21587 server.go:825] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
8月 27 03:29:45 localhost.localdomain kubelet[21587]: I0827 03:29:45.163974   21587 server.go:251] Did not find a configuration for this Kubelet via API server: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
```

More
```
[vagrant@localhost ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service | grep "E0827"
8月 27 02:42:51 localhost.localdomain kubelet[14891]: E0827 02:42:51.950796   14891 remote_runtime.go:277] ContainerStatus "d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922" from runtime service failed: rpc error: code = 2 desc = Error response from daemon: devmapper: Unknown device 5812bf809bd680ae77d4a6bbabf8589b61ff409403e8353b1ebe3834441ad579
8月 27 02:42:51 localhost.localdomain kubelet[14891]: E0827 02:42:51.950938   14891 kuberuntime_gc.go:127] Failed to remove container "d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922": failed to get container status "d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922": rpc error: code = 2 desc = Error response from daemon: devmapper: Unknown device 5812bf809bd680ae77d4a6bbabf8589b61ff409403e8353b1ebe3834441ad579
8月 27 02:42:51 localhost.localdomain kubelet[14891]: E0827 02:42:51.968127   14891 remote_runtime.go:277] ContainerStatus "7ecf5a41ea4c925b1a3fd1e090404904fd4dd8eace9cfb4e2c6e977cf447bd30" from runtime service failed: rpc error: code = 2 desc = Error: No such container: 7ecf5a41ea4c925b1a3fd1e090404904fd4dd8eace9cfb4e2c6e977cf447bd30
8月 27 02:42:52 localhost.localdomain kubelet[14891]: E0827 02:42:52.333841   14891 remote_runtime.go:277] ContainerStatus "d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922" from runtime service failed: rpc error: code = 2 desc = Error: No such container: d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922
8月 27 02:42:52 localhost.localdomain kubelet[14891]: E0827 02:42:52.340575   14891 kuberuntime_container.go:411] ContainerStatus for d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922 error: rpc error: code = 2 desc = Error: No such container: d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922
8月 27 02:42:52 localhost.localdomain kubelet[14891]: E0827 02:42:52.341663   14891 kuberuntime_manager.go:857] getPodContainerStatuses for pod "nexus-4015917150-lbs5w_default(fd00e463-551d-11e7-8e5e-0800274654e7)" failed: rpc error: code = 2 desc = Error: No such container: d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922
8月 27 02:42:52 localhost.localdomain kubelet[14891]: E0827 02:42:52.343493   14891 generic.go:241] PLEG: Ignoring events for pod nexus-4015917150-lbs5w/default: rpc error: code = 2 desc = Error: No such container: d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922
8月 27 02:42:52 localhost.localdomain kubelet[14891]: E0827 02:42:52.379756   14891 pod_workers.go:182] Error syncing pod fd00e463-551d-11e7-8e5e-0800274654e7 ("nexus-4015917150-lbs5w_default(fd00e463-551d-11e7-8e5e-0800274654e7)"), skipping: rpc error: code = 2 desc = Error: No such container: d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922
8月 27 02:42:52 localhost.localdomain kubelet[14891]: E0827 02:42:52.457853   14891 pod_workers.go:182] Error syncing pod fd00e463-551d-11e7-8e5e-0800274654e7 ("nexus-4015917150-lbs5w_default(fd00e463-551d-11e7-8e5e-0800274654e7)"), skipping: rpc error: code = 2 desc = Error: No such container: d05d41c8b2dab32d5db8a5bad29ff036257700b231a23e6c8ba599995b3a3922
8月 27 02:49:06 localhost.localdomain kubelet[14891]: E0827 02:49:06.048003   14891 summary.go:70] Partial failure issuing GetContainerInfoV2: partial failures: ["/kubepods/besteffort/podfd00e463-551d-11e7-8e5e-0800274654e7/9d007afc4edd8c66a3968607e563d740db2ba8a95e98442c496a43fd1bc42d25": RecentStats: unable to find data for container /kubepods/besteffort/podfd00e463-551d-11e7-8e5e-0800274654e7/9d007afc4edd8c66a3968607e563d740db2ba8a95e98442c496a43fd1bc42d25]
8月 27 02:49:15 localhost.localdomain kubelet[14891]: E0827 02:49:15.993325   14891 remote_runtime.go:277] ContainerStatus "4df42bc0ee8c625fd6748c2995a89c75efa7d97f403dff68c7985c9986866711" from runtime service failed: rpc error: code = 2 desc = Error: No such container: 4df42bc0ee8c625fd6748c2995a89c75efa7d97f403dff68c7985c9986866711
8月 27 02:49:15 localhost.localdomain kubelet[14891]: E0827 02:49:15.993452   14891 kuberuntime_container.go:411] ContainerStatus for 4df42bc0ee8c625fd6748c2995a89c75efa7d97f403dff68c7985c9986866711 error: rpc error: code = 2 desc = Error: No such container: 4df42bc0ee8c625fd6748c2995a89c75efa7d97f403dff68c7985c9986866711
8月 27 02:49:15 localhost.localdomain kubelet[14891]: E0827 02:49:15.993484   14891 kuberuntime_manager.go:857] getPodContainerStatuses for pod "nexus-4015917150-lbs5w_default(fd00e463-551d-11e7-8e5e-0800274654e7)" failed: rpc error: code = 2 desc = Error: No such container: 4df42bc0ee8c625fd6748c2995a89c75efa7d97f403dff68c7985c9986866711
8月 27 02:49:15 localhost.localdomain kubelet[14891]: E0827 02:49:15.993495   14891 generic.go:241] PLEG: Ignoring events for pod nexus-4015917150-lbs5w/default: rpc error: code = 2 desc = Error: No such container: 4df42bc0ee8c625fd6748c2995a89c75efa7d97f403dff68c7985c9986866711
8月 27 02:50:08 localhost.localdomain kubelet[14891]: E0827 02:50:08.352011   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:50:09 localhost.localdomain kubelet[14891]: E0827 02:50:09.103735   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:50:11 localhost.localdomain kubelet[14891]: E0827 02:50:11.839299   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:50:22 localhost.localdomain kubelet[14891]: E0827 02:50:22.895164   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:50:36 localhost.localdomain kubelet[14891]: E0827 02:50:36.826305   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:50:51 localhost.localdomain kubelet[14891]: E0827 02:50:51.833921   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:51:07 localhost.localdomain kubelet[14891]: E0827 02:51:07.852915   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:51:18 localhost.localdomain kubelet[14891]: E0827 02:51:18.866098   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:51:30 localhost.localdomain kubelet[14891]: E0827 02:51:30.832543   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:51:42 localhost.localdomain kubelet[14891]: E0827 02:51:42.829105   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:51:57 localhost.localdomain kubelet[14891]: E0827 02:51:57.819228   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:52:09 localhost.localdomain kubelet[14891]: E0827 02:52:09.835121   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:52:24 localhost.localdomain kubelet[14891]: E0827 02:52:24.847126   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 02:52:37 localhost.localdomain kubelet[14891]: E0827 02:52:37.828487   14891 pod_workers.go:182] Error syncing pod fce7a065-551d-11e7-8e5e-0800274654e7 ("jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"), skipping: failed to "StartContainer" for "jenkins" with CrashLoopBackOff: "Back-off 2m40s restarting failed container=jenkins pod=jenkins-4278279665-4t0qd_default(fce7a065-551d-11e7-8e5e-0800274654e7)"
8月 27 03:20:45 localhost.localdomain kubelet[21587]: E0827 03:20:45.157953   21587 server.go:412] failed to init dynamic Kubelet configuration sync: cloud provider was nil, and attempt to use hostname to find config resulted in: configmaps "kubelet-172.17.4.50" not found
```

### APIServer

kube_apiserver, refer to [v1.6.4 hand-on](../k8s-v1.6-deployment)

Configured directory
```
[vagrant@localhost ~]$ ls /srv/
kubernetes  sshproxy
```

basicauth
```
[vagrant@localhost ~]$ echo "password,admin,admin" | sudo tee -a /srv/kubernetes/basic_auth.csv
password,admin,admin
```

or
```
[vagrant@localhost ~]$ sudo curl -jkSL https://raw.githubusercontent.com/jimmycuadra/kube-up-artifacts/master/master/mnt/master-pd/srv/kubernetes/basic_auth.csv -o /srv/kubernetes/basic_auth.csv
```

tokenauth
```
[vagrant@localhost ~]$ sudo curl -jkSL https://raw.githubusercontent.com/jimmycuadra/kube-up-artifacts/master/master/mnt/master-pd/srv/kubernetes/known_tokens.csv -o /srv/kubernetes/known_tokens.csv
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   443  100   443    0     0    494      0 --:--:-- --:--:-- --:--:--   493
[vagrant@localhost ~]$ cat /srv/kubernetes/known_tokens.csv 
X4o3adoGo5Yx6oi4LUo2JB2GO8b0CatJ,kubelet,kubelet
tHDwA7KmaQ0jTLSKcPV2p3LrMhXP2e3m,kube_proxy,kube_proxy
cBvKaWGjXE5TnOsOW7juaO0JsXiNsxN1,system:scheduler,system:scheduler
2gxkVhMhQvMKydr7RwYYS66vXEJxy47s,system:controller_manager,system:controller_manager
isb1TVpwRYzGfshMVfTfSEDvwx2txvQt,system:logging,system:logging
JF5hfsUcDPg1JrTzlw9T6aKAQwn2BGn8,system:monitoring,system:monitoring
lXjgo4cUaiyjyCHcrn0OnonOxme6G4AV,system:dns,system:dns
```

or token
```
[vagrant@localhost ~]$ TOKEN=$(dd if=/dev/urandom bs=128 count=1 2>/dev/null | base64 | tr -d "=+/" | dd bs=32 count=1 2>/dev/null); echo $TOKEN
jbRc91Esa7kkuXNGqp5JL21gl3XeNMPU
```

for SSH proxy
```
[vagrant@localhost ~]$ sudo curl -jkSL https://raw.githubusercontent.com/mitchellh/vagrant/master/keys/vagrant -o /srv/sshproxy/.sshkeyfile
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1675  100  1675    0     0   1859      0 --:--:-- --:--:-- --:--:--  1859
[vagrant@localhost ~]$ sudo chmod 400 /srv/sshproxy/.sshkeyfile 
```

ABAC 
```
[vagrant@localhost ~]$ sed -e '/^.*kube_user.*$/d' /opt/kubernetes/cluster/saltbase/salt/kube-apiserver/abac-authz-policy.jsonl 
{"apiVersion": "abac.authorization.kubernetes.io/v1beta1", "kind": "Policy", "spec": {"user":"admin", "namespace": "*", "resource": "*", "apiGroup": "*", "nonResourcePath": "*"}}
{"apiVersion": "abac.authorization.kubernetes.io/v1beta1", "kind": "Policy", "spec": {"user":"kubelet", "namespace": "*", "resource": "*", "apiGroup": "*", "nonResourcePath": "*"}}
{"apiVersion": "abac.authorization.kubernetes.io/v1beta1", "kind": "Policy", "spec": {"user":"kube_proxy", "namespace": "*", "resource": "*", "apiGroup": "*", "nonResourcePath": "*"}}
{"apiVersion": "abac.authorization.kubernetes.io/v1beta1", "kind": "Policy", "spec": {"user":"kubecfg", "namespace": "*", "resource": "*", "apiGroup": "*", "nonResourcePath": "*"}}
{"apiVersion": "abac.authorization.kubernetes.io/v1beta1", "kind": "Policy", "spec": {"user":"client", "namespace": "*", "resource": "*", "apiGroup": "*", "nonResourcePath": "*"}}
{"apiVersion": "abac.authorization.kubernetes.io/v1beta1", "kind": "Policy", "spec": {"group":"system:serviceaccounts", "namespace": "*", "resource": "*", "apiGroup": "*", "nonResourcePath": "*"}}

log
```
[vagrant@localhost ~]$ ls /var/log/kube-apiserver*
/var/log/kube-apiserver-audit.log  /var/log/kube-apiserver.log
```

POD manifest
```
[vagrant@localhost ~]$ sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%\("image": \)".*kube_docker_registry\x27\x5d\x7d\x7d\(/kube-apiserver:\).*kube-apiserver_docker_tag.*"%\1"gcr.io/google_containers\2$KUBE_APISERVER_DOCKERTAG"%;s%{{params}}%--address=127.0.0.1 --storage-backend=etcd3 --storage-media-type=application/vnd.kubernetes.protobuf --etcd-servers=http://127.0.0.1:2379 --etcd-servers-overrides=/events#http://127.0.0.1:2379 --runtime-config=api/all=true --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true --admission-control=ResourceQuota,ServiceAccount --max-requests-inflight=400 --target-ram-mb=1024 --service-cluster-ip-range=10.3.0.0/20 --client-ca-file=/srv/kubernetes/ca.crt --basic-auth-file=/srv/kubernetes/basic_auth.csv --min-request-timeout=1800 --enable-garbage-collector=true --etcd-quorum-read=true --audit-log-path=/var/log/kube-apiserver-audit.log --audit-log-maxage=0 --audit-log-maxbackup=0 --audit-log-maxsize=2000000000 --audit-webhook-mode=batch --tls-cert-file=/srv/kubernetes/server.cert --tls-private-key-file=/srv/kubernetes/server.key --kubelet-client-certificate=/srv/kubernetes/kubeapiserver.cert --kubelet-client-key=/srv/kubernetes/kubeapiserver.key --secure-port=6443 --token-auth-file=/dev/null --bind-address=0.0.0.0 --v=2 --advertise-address=$KUBE_APISERVER_HOST --ssh-user=vagrant --ssh-keyfile=/srv/sshproxy/.sshkeyfile                                         --authorization-mode=ABAC --authorization-policy-file=/srv/kubernetes/abac-authz-policy.jsonl --enable-swagger-ui=true%;s/{{pillar.*allow_privileged.*}}/true/;s/{{container_env}}//;s/{{secure_port}}/6443/g;s/{{cloud_config_mount}}\|{{additional_cloud_config_mount}}\|{{webhook_config_mount}}\|{{webhook_authn_config_mount}}\|{{audit_policy_config_mount}}\|{{audit_webhook_config_mount}}\|{{admission_controller_config_mount}}\|{{image_policy_webhook_config_mount}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;s%{{srv_sshproxy_path}}%/srv/sshproxy%g;s/{{cloud_config_volume}}\|{{additional_cloud_config_volume}}\|{{webhook_config_volume}}\|{{webhook_authn_config_volume}}\|{{audit_policy_config_volume}}\|{{audit_webhook_config_volume}}\|{{admission_controller_config_volume}}\|{{image_policy_webhook_config_volume}}//;s%"path": "/etc/srv/pki"%"path": "/etc/pki"%' /opt/kubernetes/cluster/saltbase/salt/kube-apiserver/kube-apiserver.manifest | KUBE_APISERVER_DOCKERTAG=$(cat /opt/kubernetes/server/bin/kube-apiserver.docker_tag) KUBE_APISERVER_HOST=$HOSTNAME_OVERRIDE envsubst | sudo tee /etc/kubernetes/manifests/kube-apiserver.json


{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"kube-apiserver",
  "namespace": "kube-system",
  "annotations": {
    "scheduler.alpha.kubernetes.io/critical-pod": ""
  },
  "labels": {
    "tier": "control-plane",
    "component": "kube-apiserver"
  }
},
"spec":{
"hostNetwork": true,
"containers":[
    {
    "name": "kube-apiserver",
    "image": "gcr.io/google_containers/kube-apiserver:v1.7.4",
    "resources": {
      "requests": {
        "cpu": "250m"
      }
    },
    "command": [
                 "/bin/sh",
                 "-c",
                 "/usr/local/bin/kube-apiserver --address=127.0.0.1 --storage-backend=etcd3 --storage-media-type=application/vnd.kubernetes.protobuf --etcd-servers=http://127.0.0.1:2379 --etcd-servers-overrides=/events#http://127.0.0.1:2379 --runtime-config=api/all=true --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true --admission-control=ResourceQuota,ServiceAccount --max-requests-inflight=400 --target-ram-mb=1024 --service-cluster-ip-range=10.3.0.0/20 --client-ca-file=/srv/kubernetes/ca.crt --basic-auth-file=/srv/kubernetes/basic_auth.csv --min-request-timeout=1800 --enable-garbage-collector=true --etcd-quorum-read=true --audit-log-path=/var/log/kube-apiserver-audit.log --audit-log-maxage=0 --audit-log-maxbackup=0 --audit-log-maxsize=2000000000 --audit-webhook-mode=batch --tls-cert-file=/srv/kubernetes/server.cert --tls-private-key-file=/srv/kubernetes/server.key --kubelet-client-certificate=/srv/kubernetes/kubeapiserver.cert --kubelet-client-key=/srv/kubernetes/kubeapiserver.key --secure-port=6443 --token-auth-file=/dev/null --bind-address=0.0.0.0 --v=2 --advertise-address=172.17.4.50 --ssh-user=vagrant --ssh-keyfile=/srv/sshproxy/.sshkeyfile                                         --authorization-mode=ABAC --authorization-policy-file=/srv/kubernetes/abac-authz-policy.jsonl --enable-swagger-ui=true --allow-privileged=true 1>>/var/log/kube-apiserver.log 2>&1"
               ],
    
    "livenessProbe": {
      "httpGet": {
        "host": "127.0.0.1",
        "port": 8080,
        "path": "/healthz"
      },
      "initialDelaySeconds": 15,
      "timeoutSeconds": 15
    },
    "ports":[
      { "name": "https",
        "containerPort": 6443,
        "hostPort": 6443},{
       "name": "local",
        "containerPort": 8080,
        "hostPort": 8080}
        ],
    "volumeMounts": [
        
        
        
        
        
        
        
        
        { "name": "srvkube",
        "mountPath": "/srv/kubernetes",
        "readOnly": true},
        { "name": "logfile",
        "mountPath": "/var/log/kube-apiserver.log",
        "readOnly": false},
        { "name": "auditlogfile",
        "mountPath": "/var/log/kube-apiserver-audit.log",
        "readOnly": false},
        { "name": "etcssl",
        "mountPath": "/etc/ssl",
        "readOnly": true},
        { "name": "usrsharecacerts",
        "mountPath": "/usr/share/ca-certificates",
        "readOnly": true},
        { "name": "varssl",
        "mountPath": "/var/ssl",
        "readOnly": true},
        { "name": "etcopenssl",
        "mountPath": "/etc/openssl",
        "readOnly": true},
        { "name": "etcpki",
        "mountPath": "/etc/srv/pki",
        "readOnly": true},
        { "name": "srvsshproxy",
        "mountPath": "/srv/sshproxy",
        "readOnly": false}
      ]
    }
],
"volumes":[
  
  
  
  
  
  
  
  
  { "name": "srvkube",
    "hostPath": {
        "path": "/srv/kubernetes"}
  },
  { "name": "logfile",
    "hostPath": {
        "path": "/var/log/kube-apiserver.log"}
  },
  { "name": "auditlogfile",
    "hostPath": {
        "path": "/var/log/kube-apiserver-audit.log"}
  },
  { "name": "etcssl",
    "hostPath": {
        "path": "/etc/ssl"}
  },
  { "name": "usrsharecacerts",
    "hostPath": {
        "path": "/usr/share/ca-certificates"}
  },
  { "name": "varssl",
    "hostPath": {
        "path": "/var/ssl"}
  },
  { "name": "etcopenssl",
    "hostPath": {
        "path": "/etc/openssl"}
  },
  { "name": "etcpki",
    "hostPath": {
        "path": "/etc/pki"}
  },
  { "name": "srvsshproxy",
    "hostPath": {
        "path": "/srv/sshproxy"}
  }
]
}}
```

Fix (https://github.com/kubernetes/kubernetes/issues/14097)
```
sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%\("image": \)".*kube_docker_registry\x27\x5d\x7d\x7d\(/kube-apiserver:\).*kube-apiserver_docker_tag.*"%\1"gcr.io/google_containers\2$KUBE_APISERVER_DOCKERTAG"%;s%{{params}}%--address=127.0.0.1 --storage-backend=etcd3 --storage-media-type=application/vnd.kubernetes.protobuf --etcd-servers=http://127.0.0.1:2379 --etcd-servers-overrides=/events#http://127.0.0.1:2379 --runtime-config=api/all=true --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true --admission-control=ResourceQuota,ServiceAccount --max-requests-inflight=400 --target-ram-mb=1024 --service-cluster-ip-range=10.3.0.0/20 --client-ca-file=/srv/kubernetes/ca.crt --basic-auth-file=/srv/kubernetes/basic_auth.csv --min-request-timeout=1800 --enable-garbage-collector=true --etcd-quorum-read=true --audit-log-path=/var/log/kube-apiserver-audit.log --audit-log-maxage=0 --audit-log-maxbackup=0 --audit-log-maxsize=2000000000 --audit-webhook-mode=batch --tls-cert-file=/srv/kubernetes/server.cert --tls-private-key-file=/srv/kubernetes/server.key --secure-port=6443 --token-auth-file=/dev/null --bind-address=0.0.0.0 --v=2 --advertise-address=$KUBE_APISERVER_HOST --authorization-mode=ABAC --authorization-policy-file=/srv/kubernetes/abac-authz-policy.jsonl --enable-swagger-ui=true%;s/{{pillar.*allow_privileged.*}}/true/;s/{{container_env}}//;s/{{secure_port}}/6443/g;s/{{cloud_config_mount}}\|{{additional_cloud_config_mount}}\|{{webhook_config_mount}}\|{{webhook_authn_config_mount}}\|{{audit_policy_config_mount}}\|{{audit_webhook_config_mount}}\|{{admission_controller_config_mount}}\|{{image_policy_webhook_config_mount}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;s%{{srv_sshproxy_path}}%/srv/sshproxy%g;s/{{cloud_config_volume}}\|{{additional_cloud_config_volume}}\|{{webhook_config_volume}}\|{{webhook_authn_config_volume}}\|{{audit_policy_config_volume}}\|{{audit_webhook_config_volume}}\|{{admission_controller_config_volume}}\|{{image_policy_webhook_config_volume}}//;s%"path": "/etc/srv/pki"%"path": "/etc/pki"%' /opt/kubernetes/cluster/saltbase/salt/kube-apiserver/kube-apiserver.manifest | KUBE_APISERVER_DOCKERTAG=$(cat /opt/kubernetes/server/bin/kube-apiserver.docker_tag) KUBE_APISERVER_HOST=$HOSTNAME_OVERRIDE envsubst | sudo tee /etc/kubernetes/manifests/kube-apiserver.json
```

### Client

kubectl, refer to [v1.6.4 hand-on](../k8s-v1.6-deployment)
```
[vagrant@localhost ~]$ ln -sf /opt/kubernetes/server/bin/kubectl ./go/bin/
```

Version
```
[vagrant@localhost ~]$ kubectl version --client
Client Version: version.Info{Major:"1", Minor:"3", GitVersion:"v1.3.0+57fb9ac", GitCommit:"57fb9ac", GitTreeState:"clean", BuildDate:"2016-08-04T06:26:19Z", GoVersion:"go1.6.2", Compiler:"gc", Platform:"linux/amd64"}
[vagrant@localhost ~]$ kubectl version
Client Version: version.Info{Major:"1", Minor:"3", GitVersion:"v1.3.0+57fb9ac", GitCommit:"57fb9ac", GitTreeState:"clean", BuildDate:"2016-08-04T06:26:19Z", GoVersion:"go1.6.2", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"7", GitVersion:"v1.7.4", GitCommit:"793658f2d7ca7f064d2bdf606519f9fe1229c381", GitTreeState:"clean", BuildDate:"2017-08-17T08:30:51Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
[vagrant@localhost ~]$ kubectl api-versions
admissionregistration.k8s.io/v1alpha1
apiextensions.k8s.io/v1beta1
apiregistration.k8s.io/v1beta1
apps/v1beta1
authentication.k8s.io/v1
authentication.k8s.io/v1beta1
authorization.k8s.io/v1
authorization.k8s.io/v1beta1
autoscaling/v1
autoscaling/v2alpha1
batch/v1
batch/v2alpha1
certificates.k8s.io/v1beta1
extensions/v1beta1
networking.k8s.io/v1
policy/v1beta1
rbac.authorization.k8s.io/v1alpha1
rbac.authorization.k8s.io/v1beta1
settings.k8s.io/v1alpha1
storage.k8s.io/v1
storage.k8s.io/v1beta1
v1
```

kube-system
```
[vagrant@localhost ~]$ /Users/fanhongling/Downloads/99-mirror/linux-bin/k8s/v1.7.4/kubernetes/server/bin/kubectl get all --namespace=kube-system
NAME                                       READY     STATUS    RESTARTS   AGE
po/etcd-server-172.17.4.50                 1/1       Running   0          4h
po/kube-addon-manager-172.17.4.50          1/1       Running   0          4h
po/kube-apiserver-172.17.4.50              1/1       Running   1          17m
po/kube-controller-manager-172.17.4.50     1/1       Running   2          4h
po/kube-dns-806549836-3n26k                3/3       Running   33         76d
po/kube-proxy-172.17.4.50                  1/1       Running   0          4h
po/kube-scheduler-172.17.4.50              1/1       Running   2          4h
po/kubernetes-dashboard-1995634392-362sf   1/1       Running   5          76d

NAME                          CLUSTER-IP   EXTERNAL-IP   PORT(S)          AGE
svc/kube-dns                  10.3.0.10    <none>        53/UDP,53/TCP    76d
svc/kubernetes-dashboard      10.3.8.135   <none>        80/TCP           76d
svc/kubernetes-dashboard-ex   10.3.7.209   <nodes>       9090:32186/TCP   76d

NAME                          DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/kube-dns               1         1         1            1           76d
deploy/kubernetes-dashboard   1         1         1            1           76d

NAME                                 DESIRED   CURRENT   READY     AGE
rs/kube-dns-806549836                1         1         1         76d
rs/kubernetes-dashboard-1995634392   1         1         1         76d
```

### ControllerManager

kube-controller-manager

log
```
[vagrant@localhost ~]$ ls /var/log/kube-controller-manager.log 
/var/log/kube-controller-manager.log
```
manifest
```
[vagrant@localhost ~]$ sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%\("image": \)".*kube_docker_registry\x27\x5d\x7d\x7d\(/kube-controller-manager:\).*kube-controller-manager_docker_tag.*"%\1"gcr.io/google_containers\2$KUBE_CONTROLLERMANAGER_DOCKERTAG"%;s%{{params}}%--master=127.0.0.1:8080 --cluster-name=kubernetes --cluster-cidr=10.0.0.0/14 --service-cluster-ip-range=10.3.0.0/20 --terminated-pod-gc-threshold=12500 --enable-garbage-collector=true --service-account-private-key-file=/srv/kubernetes/server.key --v=2 --root-ca-file=/srv/kubernetes/ca.crt --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true --enable-hostpath-provisioner=true%;s/{{container_env}}//;s/{{cloud_config_mount}}\|{{additional_cloud_config_mount}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;s%{{cloud_config_volume}}\|{{additional_cloud_config_volume}}%%' /opt/kubernetes/cluster/saltbase/salt/kube-controller-manager/kube-controller-manager.manifest | KUBE_CONTROLLERMANAGER_DOCKERTAG=$(cat /opt/kubernetes/server/bin/kube-controller-manager.docker_tag) envsubst | sudo tee /etc/kubernetes/manifests/kube-controller-manager.json


{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"kube-controller-manager",
  "namespace": "kube-system",
  "annotations": {
    "scheduler.alpha.kubernetes.io/critical-pod": ""
  },
  "labels": {
    "tier": "control-plane",
    "component": "kube-controller-manager"
  }
},
"spec":{
"hostNetwork": true,
"containers":[
    {
    "name": "kube-controller-manager",
    "image": "gcr.io/google_containers/kube-controller-manager:v1.7.4",
    "resources": {
      "requests": {
        "cpu": "200m"
      }
    },
    "command": [
                 "/bin/sh",
                 "-c",
                 "/usr/local/bin/kube-controller-manager --master=127.0.0.1:8080 --cluster-name=kubernetes --cluster-cidr=10.0.0.0/14 --service-cluster-ip-range=10.3.0.0/20 --terminated-pod-gc-threshold=12500 --enable-garbage-collector=true --service-account-private-key-file=/srv/kubernetes/server.key --v=2 --root-ca-file=/srv/kubernetes/ca.crt --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true --enable-hostpath-provisioner=true 1>>/var/log/kube-controller-manager.log 2>&1"
               ],
    
    "livenessProbe": {
      "httpGet": {
        "host": "127.0.0.1",
        "port": 10252,
        "path": "/healthz"
      },
      "initialDelaySeconds": 15,
      "timeoutSeconds": 15
    },
    "volumeMounts": [
        
        
        { "name": "srvkube",
        "mountPath": "/srv/kubernetes",
        "readOnly": true},
        { "name": "logfile",
        "mountPath": "/var/log/kube-controller-manager.log",
        "readOnly": false},
        { "name": "etcssl",
        "mountPath": "/etc/ssl",
        "readOnly": true},
        { "name": "usrsharecacerts",
        "mountPath": "/usr/share/ca-certificates",
        "readOnly": true},
        { "name": "varssl",
        "mountPath": "/var/ssl",
        "readOnly": true},
        { "name": "etcopenssl",
        "mountPath": "/etc/openssl",
        "readOnly": true},
        { "name": "etcpki",
        "mountPath": "/etc/pki",
        "readOnly": true}
      ]
    }
],
"volumes":[
  
  
  { "name": "srvkube",
    "hostPath": {
        "path": "/srv/kubernetes"}
  },
  { "name": "logfile",
    "hostPath": {
        "path": "/var/log/kube-controller-manager.log"}
  },
  { "name": "etcssl",
    "hostPath": {
        "path": "/etc/ssl"}
  },
  { "name": "usrsharecacerts",
    "hostPath": {
        "path": "/usr/share/ca-certificates"}
  },
  { "name": "varssl",
    "hostPath": {
        "path": "/var/ssl"}
  },
  { "name": "etcopenssl",
    "hostPath": {
        "path": "/etc/openssl"}
  },
  { "name": "etcpki",
    "hostPath": {
        "path": "/etc/pki"}
  }
]
}}
```

### Scheduler

kube-scheduler

log
```
[vagrant@localhost ~]$ sudo touch /var/log/kube-scheduler.log 
/var/log/kube-scheduler.log
```

manifest
```
[vagrant@localhost ~]$ sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%\("image": \)".*kube_docker_registry\x27\x5d\x7d\x7d\(/kube-scheduler:\).*kube-scheduler_docker_tag.*"%\1"gcr.io/google_containers\2$KUBE_SCHEDULER_DOCKERTAG"%;s%{{params}}%--master=127.0.0.1:8080 --v=2 --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true --algorithm-provider=DefaultProvider%;s%{{srv_kube_path}}%/srv/kubernetes%g' /opt/kubernetes/cluster/saltbase/salt/kube-scheduler/kube-scheduler.manifest | KUBE_SCHEDULER_DOCKERTAG=$(cat /opt/kubernetes/server/bin/kube-scheduler.docker_tag) envsubst | sudo tee /etc/kubernetes/manifests/kube-scheduler.json

{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"kube-scheduler",
  "namespace": "kube-system",
  "annotations": {
    "scheduler.alpha.kubernetes.io/critical-pod": ""
  },
  "labels": {
    "tier": "control-plane",
    "component": "kube-scheduler"
  }
},
"spec":{
"hostNetwork": true,
"containers":[
    {
    "name": "kube-scheduler",
    "image": "gcr.io/google_containers/kube-scheduler:v1.7.4",
    "resources": {
      "requests": {
        "cpu": "75m"
      }
    },
    "command": [
                 "/bin/sh",
                 "-c",
                 "/usr/local/bin/kube-scheduler --master=127.0.0.1:8080 --v=2 --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true --algorithm-provider=DefaultProvider 1>>/var/log/kube-scheduler.log 2>&1"
               ],
    "livenessProbe": {
      "httpGet": {
        "host": "127.0.0.1",
        "port": 10251,
        "path": "/healthz"
      },
      "initialDelaySeconds": 15,
      "timeoutSeconds": 15
    },
    "volumeMounts": [
        {
          "name": "logfile",
          "mountPath": "/var/log/kube-scheduler.log",
          "readOnly": false
        },
        {
          "name": "srvkube",
          "mountPath": "/srv/kubernetes",
          "readOnly": true
        }
      ]
    }
],
"volumes":[
  {
    "name": "srvkube",
    "hostPath": {"path": "/srv/kubernetes"}
  },
  {
    "name": "logfile",
    "hostPath": {"path": "/var/log/kube-scheduler.log"}
  }
]
}}
```

### IPTables Proxy

kube-proxy

log
```
[vagrant@localhost ~]$ sudo touch /var/log/kube-proxy.log 
```

kubeconfig
```
[vagrant@localhost ~]$ ls /var/lib/kube-proxy/
kubeconfig
```

manifest
```
[vagrant@localhost ~]$ sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%\(image: \).*kube_docker_registry\x27\x5d\x7d\x7d\(/kube-proxy:\).*kube-proxy_docker_tag.*%\1gcr.io/google_containers\2$KUBE_PROXY_DOCKERTAG%;s/{{ cpurequest }}/100m/;s%{{api_servers_with_port}}%--master=https://$KUBE_APISERVER_HOST:6443%;s%{{kubeconfig}}%--kubeconfig=/var/lib/kube-proxy/kubeconfig%;s%{{cluster_cidr}}%--cluster-cidr=10.0.0.0/14%;s%{{params}}%--v=2 --iptables-sync-period=1m --iptables-min-sync-period=10s --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true%;s/{{container_env}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;' /opt/kubernetes/cluster/saltbase/salt/kube-proxy/kube-proxy.manifest | KUBE_PROXY_DOCKERTAG=$(cat /opt/kubernetes/server/bin/kube-proxy.docker_tag) KUBE_APISERVER_HOST=$HOSTNAME_OVERRIDE envsubst | sudo tee /etc/kubernetes/manifests/kube-proxy.yaml


# kube-proxy podspec
apiVersion: v1
kind: Pod
metadata:
  name: kube-proxy
  namespace: kube-system
  # This annotation ensures that kube-proxy does not get evicted if the node
  # supports critical pod annotation based priority scheme.
  # Note that kube-proxy runs as a static pod so this annotation does NOT have
  # any effect on rescheduler (default scheduler and rescheduler are not
  # involved in scheduling kube-proxy).
  annotations:
    scheduler.alpha.kubernetes.io/critical-pod: ''
  labels:
    tier: node
    component: kube-proxy
spec:
  hostNetwork: true
  initContainers:
  - name: touch-lock
    image: busybox
    command: ['/bin/touch', '/run/xtables.lock']
    securityContext:
      privileged: true
    volumeMounts:
    - mountPath: /run
      name: run
      readOnly: false
  containers:
  - name: kube-proxy
    image: gcr.io/google_containers/kube-proxy:v1.7.4
    resources:
      requests:
        cpu: 100m
    command:
    - /bin/sh
    - -c
    - echo -998 > /proc/$$$/oom_score_adj && kube-proxy --master=https://172.17.4.50:6443 --kubeconfig=/var/lib/kube-proxy/kubeconfig --cluster-cidr=10.0.0.0/14 --resource-container="" --v=2 --iptables-sync-period=1m --iptables-min-sync-period=10s --feature-gates=Accelerators=true,AdvancedAuditing=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,LocalStorageCapacityIsolation=true,PersistentLocalVolumes=true,RotateKubeletClientCertificate=true,RotateKubeletServerCertificate=true,TaintBasedEvictions=true 1>>/var/log/kube-proxy.log 2>&1
    
    securityContext:
      privileged: true
    volumeMounts:
    - mountPath: /etc/ssl/certs
      name: etc-ssl-certs
      readOnly: true
    - mountPath: /usr/share/ca-certificates
      name: usr-ca-certs
      readOnly: true
    - mountPath: /var/log
      name: varlog
      readOnly: false
    - mountPath: /var/lib/kube-proxy/kubeconfig
      name: kubeconfig
      readOnly: false
    - mountPath: /run/xtables.lock
      name: iptableslock
      readOnly: false
  volumes:
  - hostPath:
      path: /usr/share/ca-certificates
    name: usr-ca-certs
  - hostPath:
      path: /etc/ssl/certs
    name: etc-ssl-certs
  - hostPath:
      path: /var/lib/kube-proxy/kubeconfig
    name: kubeconfig
  - hostPath:
      path: /var/log
    name: varlog
  - hostPath:
      path: /run
    name: run
  - hostPath:
      path: /run/xtables.lock
    name: iptableslock
```

fabric8io
```
[vagrant@localhost ~]$ kubectl get ingress
NAME                      HOSTS                                                ADDRESS   PORTS     AGE
fabric8                   fabric8.default.172.17.4.50.xip.io                             80        68d
fabric8-docker-registry   fabric8-docker-registry.default.172.17.4.50.xip.io             80        68d
fabric8-forge             fabric8-forge.default.172.17.4.50.xip.io                       80        68d
gogs                      gogs.default.172.17.4.50.xip.io                                80        68d
gogs-ssh                  gogs-ssh.default.172.17.4.50.xip.io                            80        68d
jenkins                   jenkins.default.172.17.4.50.xip.io                             80        68d
jenkins-jnlp              jenkins-jnlp.default.172.17.4.50.xip.io                        80        68d
jenkinshift               jenkinshift.default.172.17.4.50.xip.io                         80        68d
nexus                     nexus.default.172.17.4.50.xip.io                               80        68d
```
![屏幕快照 2017-08-27 上午3.52.12.png](./屏幕快照%202017-08-27%20上午3.52.12.png)

## Addons

for DNS image
```
[vagrant@localhost ~]$ docker pull gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.4
Trying to pull repository gcr.io/google_containers/k8s-dns-sidecar-amd64 ... 
1.14.4: Pulling from gcr.io/google_containers/k8s-dns-sidecar-amd64
a35579ce456b: Pull complete 
17e884962b9e: Pull complete 
Digest: sha256:97074c951046e37d3cbb98b82ae85ed15704a290cce66a8314e7f846404edde9
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.4
```

and Dashboard image
```
[vagrant@localhost ~]$ docker pull gcr.io/google_containers/kubernetes-dashboard-amd64:v1.6.1
Trying to pull repository gcr.io/google_containers/kubernetes-dashboard-amd64 ... 
v1.6.1: Pulling from gcr.io/google_containers/kubernetes-dashboard-amd64
f084628ae6c6: Pull complete 
Digest: sha256:b537ce8988510607e95b8d40ac9824523b1f9029e6f9f90e9fccc663c355cf5d
Status: Downloaded newer image for gcr.io/google_containers/kubernetes-dashboard-amd64:v1.6.1
```