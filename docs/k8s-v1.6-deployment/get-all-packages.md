# Deploy K8s single under Kubelet － 在CentOS/Fedora部署Kubernetes v1.6.x

__Abstract__

Kubernetes v1.6.4

__Tips__

另可参考[K8s v1.5.7安装](../k8s-v1.5.7-deployment)


### Vagrant synced_folders into VirtualBox

Mac
```
    "sync_folders"      : 
        {
            ".":
              {
                "to"      : "/data/src/github.com/openshift/origin",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/99-mirror":
              {
                "to"      : "/Users/fanhongling/Downloads/99-mirror",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/99-mirror/linux-bin/k8s-v1.6.4/kubernetes":
              {
                "to"      : "/opt/kubernetes",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/go-kubernetes/pkg":
              {
                "to"      : "/Users/fanhongling/Downloads/go-kubernetes/pkg",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/go-kubernetes/src":
              {
                "to"      : "/Users/fanhongling/Downloads/go-kubernetes/src",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/go-openshift/pkg":
              {
                "to"      : "/Users/fanhongling/Downloads/go-openshift/pkg",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/go-openshift/src":
              {
                "to"      : "/Users/fanhongling/Downloads/go-openshift/src",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/workspace":
              {
                "to"      : "/Users/fanhongling/Downloads/workspace",
                "exclude" : null
              },
            "/Users/fanhongling/go/pkg":
              {
                "to"      : "/Users/fanhongling/go/pkg",
                "exclude" : null
              },
            "/Users/fanhongling/go/src":
              {
                "to"      : "/Users/fanhongling/go/src",
                "exclude" : null
              }
        }
```

### Shutdown Firewall & Selinux
```
[tangfx@localhost ~]$ sudo systemctl stop firewalld.service
[tangfx@localhost ~]$ sudo systemctl disable firewalld.service
[vagrant@localhost ~]$ systemctl is-active firewalld.service
unknown
[tangfx@localhost ~]$ sudo setenforce 0
[tangfx@localhost ~]$ sudo sed -i 's/SELINUX=.*/SELINUE=permissive/' /etc/sysconfig/selinux
[tangfx@localhost ~]$ sudo getenforce
Permissive
```

### 安装Docker Engine 1.10+

__CentOS 7.3__

从Windows Cygwin登录CentOS虚拟机
```
tangf@DESKTOP-H68OQDV ~
$ ssh tangfx@10.64.33.80
```

安装Docker engine
```
[tangfx@localhost ~]$ sudo yum install -y docker
[tangfx@localhost ~]$ sudo systemctl start docker.service
[tangfx@localhost ~]$ sudo systemctl -l status docker.service
[tangfx@localhost ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit docker.service
[tangfx@localhost ~]$ sudo systemctl enable docker.service
[tangfx@localhost ~]$ sudo usermod -aG root tangfx
[tangfx@localhost ~]$ exit
```

Log out然后Log in
```
tangf@DESKTOP-H68OQDV ~
$ ssh tangfx@10.64.33.80
```

查看Docker版本和信息(CentOS7.3)
```
[tangfx@localhost ~]$ docker version
Client:
 Version:         1.12.6
 API version:     1.24
 Package version: docker-common-1.12.6-16.el7.centos.x86_64
 Go version:      go1.7.4
 Git commit:      3a094bd/1.12.6
 Built:           Fri Apr 14 13:46:13 2017
 OS/Arch:         linux/amd64

Server:
 Version:         1.12.6
 API version:     1.24
 Package version: docker-common-1.12.6-16.el7.centos.x86_64
 Go version:      go1.7.4
 Git commit:      3a094bd/1.12.6
 Built:           Fri Apr 14 13:46:13 2017
 OS/Arch:         linux/amd64
[tangfx@localhost ~]$ docker info
[tangfx@localhost ~]$ ls /usr/lib/systemd/system/docker.service
[tangfx@localhost ~]$ ls /etc/sysconfig/docker
[tangfx@localhost ~]$ ls /etc/docker
[tangfx@localhost ~]$ ls /var/lib/docker
```

__Fedora 23__

Update _docker engine_
```
[vagrant@localhost ~]$ docker version
Client:
 Version:         1.9.1
 API version:     1.21
 Package version: docker-1.9.1-6.git6ec29ef.fc23.x86_64
 Go version:      go1.5.3
 Git commit:      7206621
 Built:           Mon Jan 25 20:52:22 UTC 2016
 OS/Arch:         linux/amd64

Server:
 Version:         1.9.1
 API version:     1.21
 Package version: docker-1.9.1-6.git6ec29ef.fc23.x86_64
 Go version:      go1.5.3
 Git commit:      7206621
 Built:           Mon Jan 25 20:52:22 UTC 2016
 OS/Arch:         linux/amd64
[vagrant@localhost ~]$ sudo dnf --showduplicates list docker
上次元数据过期检查在 0:04:16 前执行于 Sat Jun 10 18:56:05 2017。
已安装的软件包
docker.x86_64                                                  1:1.9.1-6.git6ec29ef.fc23                                                    @updates
可安装的软件包
docker.x86_64                                                  1:1.7.0-22.gitdcff4e1.fc23                                                   fedora  
docker.x86_64                                                  1:1.9.1-6.git6ec29ef.fc23                                                    @updates
docker.x86_64                                                  2:1.10.3-45.gite03ddb8.fc23                                                  updates 
[vagrant@localhost ~]$ sudo dnf update docker
上次元数据过期检查在 0:05:15 前执行于 Sat Jun 10 18:56:05 2017。
依赖关系解决。
====================================================================================================================================================
 Package                                   架构                     版本                                            仓库                       大小
====================================================================================================================================================
安装:
 docker-forward-journald                   x86_64                   2:1.10.3-45.gite03ddb8.fc23                     updates                   993 k
升级:
 docker                                    x86_64                   2:1.10.3-45.gite03ddb8.fc23                     updates                   9.2 M
 docker-selinux                            x86_64                   2:1.10.3-45.gite03ddb8.fc23                     updates                    71 k

事务概要
====================================================================================================================================================
安装  1 Package
升级  2 Packages

总下载：10 M
确定吗？[y/N]： y
下载软件包：
(1/3): docker-selinux-1.10.3-45.gite03ddb8.fc23.x86_64.rpm                                                           24 kB/s |  71 kB     00:02    
(2/3): docker-forward-journald-1.10.3-45.gite03ddb8.fc23.x86_64.rpm                                                 304 kB/s | 993 kB     00:03    
(3/3): docker-1.10.3-45.gite03ddb8.fc23.x86_64.rpm                                                                  1.9 MB/s | 9.2 MB     00:04    
----------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                                1.2 MB/s |  10 MB     00:08     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  升级: docker-selinux-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                                       1/5 
  安装: docker-forward-journald-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                              2/5 
  升级: docker-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                                               3/5 
  清理: docker-1:1.9.1-6.git6ec29ef.fc23.x86_64                                                                                                 4/5 
  清理: docker-selinux-1:1.9.1-6.git6ec29ef.fc23.x86_64                                                                                         5/5 
  验证: docker-forward-journald-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                              1/5 
  验证: docker-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                                               2/5 
  验证: docker-selinux-2:1.10.3-45.gite03ddb8.fc23.x86_64                                                                                       3/5 
  验证: docker-selinux-1:1.9.1-6.git6ec29ef.fc23.x86_64                                                                                         4/5 
  验证: docker-1:1.9.1-6.git6ec29ef.fc23.x86_64                                                                                                 5/5 

已安装:
  docker-forward-journald.x86_64 2:1.10.3-45.gite03ddb8.fc23                                                                                        

已升级:
  docker.x86_64 2:1.10.3-45.gite03ddb8.fc23                            docker-selinux.x86_64 2:1.10.3-45.gite03ddb8.fc23                           

完毕！
[vagrant@localhost ~]$ sudo systemctl start docker
[vagrant@localhost ~]$ docker version
Client:
 Version:         1.10.3
 API version:     1.22
 Package version: docker-1.10.3-45.gite03ddb8.fc23.x86_64
 Go version:      go1.5.4
 Git commit:      e03ddb8/1.10.3
 Built:           
 OS/Arch:         linux/amd64

Server:
 Version:         1.10.3
 API version:     1.22
 Package version: docker-1.10.3-45.gite03ddb8.fc23.x86_64
 Go version:      go1.5.4
 Git commit:      e03ddb8/1.10.3
 Built:           
 OS/Arch:         linux/amd64
```

### 下载K8s社区分发包

Download (CentOS7.3)
```
[tangfx@localhost v1.6.2]$ curl -jkSLOC - https://dl.k8s.io/v1.6.2/kubernetes-server-linux-amd64.tar.gz
** Resuming transfer from byte position 187525609
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     22      0  0:00:07  0:00:07 --:--:--    34
100  168M  100  168M    0     0  3269k      0  0:00:52  0:00:52 --:--:-- 2333k
[tangfx@localhost ~]$ curl -jkSL -O https://dl.k8s.io/v1.6.2/kubernetes.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     45      0  0:00:03  0:00:03 --:--:--    45
100 5942k  100 5942k    0     0   755k      0  0:00:07  0:00:07 --:--:-- 1444k
```

## Installation

Extract (Mac)
```
fanhonglingdeMacBook-Pro:linux-bin fanhongling$ mkdir k8s-v1.6.4
fanhonglingdeMacBook-Pro:linux-bin fanhongling$ tar -C ./k8s-v1.6.4 -zxf ../https0x3A0x2F0x2Fdl.k8s.io/v1.6.4/kubernetes-server-linux-amd64.tar.gz 
fanhonglingdeMacBook-Pro:linux-bin fanhongling$ ls k8s-v1.6.4/
kubernetes
```

Extracted (Fedora23)
```
[vagrant@localhost ~]$ ls -1R /opt/kubernetes
/opt/kubernetes:
addons
kubernetes-src.tar.gz
LICENSES
server

/opt/kubernetes/addons:

/opt/kubernetes/server:
bin

/opt/kubernetes/server/bin:
cloud-controller-manager
hyperkube
kubeadm
kube-aggregator
kube-aggregator.docker_tag
kube-aggregator.tar
kube-apiserver
kube-apiserver.docker_tag
kube-apiserver.tar
kube-controller-manager
kube-controller-manager.docker_tag
kube-controller-manager.tar
kubectl
kubefed
kubelet
kube-proxy
kube-proxy.docker_tag
kube-proxy.tar
kube-scheduler
kube-scheduler.docker_tag
kube-scheduler.tar
```

### Images

Load components image
```
[vagrant@localhost ~]$ tag=$(cat /opt/kubernetes/server/bin/kube-aggregator.docker_tag); docker load -i /opt/kubernetes/server/bin/kube-aggregator.tar; img=$(docker images gcr.io/google_containers/kube-aggregator | awk 'NR==2 {print $1}'); echo $img:$tag; docker images gcr.io/google_containers/kube-aggregator
gcr.io/google_containers/kube-aggregator:acc93b50bf7612199b908db762ae5361
REPOSITORY                                 TAG                                IMAGE ID            CREATED             VIRTUAL SIZE
gcr.io/google_containers/kube-aggregator   acc93b50bf7612199b908db762ae5361   03a126322d2e        3 weeks ago         56.36 MB
[vagrant@localhost ~]$ tag=$(cat /opt/kubernetes/server/bin/kube-apiserver.docker_tag); docker load -i /opt/kubernetes/server/bin/kube-apiserver.tar; img=$(docker images gcr.io/google_containers/kube-apiserver | awk 'NR==2 {print $1}'); echo $img:$tag; docker images gcr.io/google_containers/kube-apiserver
gcr.io/google_containers/kube-apiserver:41a3581e4847878bbcb1722636ab51ad
REPOSITORY                                TAG                                IMAGE ID            CREATED             VIRTUAL SIZE
gcr.io/google_containers/kube-apiserver   41a3581e4847878bbcb1722636ab51ad   d3e2b7788577        3 weeks ago         150.6 MB
[vagrant@localhost ~]$ tag=$(cat /opt/kubernetes/server/bin/kube-controller-manager.docker_tag); docker load -i /opt/kubernetes/server/bin/kube-controller-manager.tar; img=$(docker images gcr.io/google_containers/kube-controller-manager | awk 'NR==2 {print $1}'); echo $img:$tag; docker images gcr.io/google_containers/kube-controller-manager
gcr.io/google_containers/kube-controller-manager:6b1e075eaab8773444890c5cc17e0dcf
REPOSITORY                                         TAG                                IMAGE ID            CREATED             VIRTUAL SIZE
gcr.io/google_containers/kube-controller-manager   6b1e075eaab8773444890c5cc17e0dcf   71000af673a1        3 weeks ago         132.8 MB
[vagrant@localhost ~]$ tag=$(cat /opt/kubernetes/server/bin/kube-proxy.docker_tag); docker load -i /opt/kubernetes/server/bin/kube-proxy.tar; img=$(docker images gcr.io/google_containers/kube-proxy | awk 'NR==2 {print $1}'); echo $img:$tag; docker images gcr.io/google_containers/kube-proxy
gcr.io/google_containers/kube-proxy:4c6ae16e62496a67d3e9cc622c1006a3
REPOSITORY                            TAG                                IMAGE ID            CREATED             VIRTUAL SIZE
gcr.io/google_containers/kube-proxy   4c6ae16e62496a67d3e9cc622c1006a3   ed2aedfc02c5        3 weeks ago         109.2 MB
[vagrant@localhost ~]$ tag=$(cat /opt/kubernetes/server/bin/kube-scheduler.docker_tag); docker load -i /opt/kubernetes/server/bin/kube-scheduler.tar; img=$(docker images gcr.io/google_containers/kube-scheduler | awk 'NR==2 {print $1}'); echo $img:$tag; docker images gcr.io/google_containers/kube-scheduler
gcr.io/google_containers/kube-scheduler:c8ed221003bb194f37ef4221727bce1c
REPOSITORY                                TAG                                IMAGE ID            CREATED             VIRTUAL SIZE
gcr.io/google_containers/kube-scheduler   c8ed221003bb194f37ef4221727bce1c   e883cb967cd8        3 weeks ago         76.75 MB
```

Search images
```
[vagrant@localhost ~]$ docker images gcr.io/google_containers/kube-*
REPOSITORY                                         TAG                                IMAGE ID            CREATED             VIRTUAL SIZE
gcr.io/google_containers/kube-apiserver            41a3581e4847878bbcb1722636ab51ad   d3e2b7788577        3 weeks ago         150.6 MB
gcr.io/google_containers/kube-proxy                4c6ae16e62496a67d3e9cc622c1006a3   ed2aedfc02c5        3 weeks ago         109.2 MB
gcr.io/google_containers/kube-controller-manager   6b1e075eaab8773444890c5cc17e0dcf   71000af673a1        3 weeks ago         132.8 MB
gcr.io/google_containers/kube-scheduler            c8ed221003bb194f37ef4221727bce1c   e883cb967cd8        3 weeks ago         76.75 MB
gcr.io/google_containers/kube-aggregator           acc93b50bf7612199b908db762ae5361   03a126322d2e        3 weeks ago         56.36 MB
```

Extract __saltbase__
```
[vagrant@localhost ~]$ tar -C /opt -zxf /Users/fanhongling/Downloads/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/v1.6.4/kubernetes.tar.gz 
[vagrant@localhost ~]$ tar -C /opt/kubernetes/cluster/ -zxf /opt/kubernetes/server/kubernetes-salt.tar.gz --strip-components=1
[vagrant@localhost ~]$ ls -1 /opt/kubernetes/cluster/saltbase/
BUILD
install.sh
pillar
reactor
README.md
salt
```

Pull image of __etcd__
```
[vagrant@localhost ~]$ tag=$(cat /opt/kubernetes/cluster/saltbase/salt/etcd/etcd.manifest | grep "gcr.io/google_containers/etcd:" | sed "s/.*'etcd_docker_tag', '\([0-9.]\+\)'.*/\1/");echo $tag;docker pull gcr.io/google_containers/etcd:$tag
3.0.17
Trying to pull repository gcr.io/google_containers/etcd ... 3.0.17: Pulling from google_containers/etcd
c366cffde3c9: Pull complete 
48645f8250fe: Pull complete 
ebdd1c183fbb: Pull complete 
25c0d1d39747: Pull complete 
444e5d46c375: Pull complete 
Digest: sha256:214e1f9b18a48aac9b42b64854d084f07b40870bc7790c6ee2092f567a20b42e
Status: Downloaded newer image for gcr.io/google_containers/etcd:3.0.17
[vagrant@localhost ~]$ docker images gcr.io/google_containers/etcd:3.0.17
REPOSITORY                      TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/etcd   3.0.17              e6a6589433f4        3 months ago        168.9 MB
```

__POD__ image
```
[vagrant@localhost ~]$ docker pull gcr.io/google_containers/pause-amd64:3.0
Trying to pull repository gcr.io/google_containers/pause-amd64 ... 3.0: Pulling from google_containers/pause-amd64
Digest: sha256:163ac025575b775d1c0f9bf0bdd0f086883171eb475b5068e7defa4ca9e76516
Status: Image is up to date for gcr.io/google_containers/pause-amd64:3.0
[vagrant@localhost ~]$ docker images gcr.io/google_containers/pause-amd64:3.0
REPOSITORY                             TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/pause-amd64   3.0                 99e59f495ffa        13 months ago       746.9 kB
```

Create directory for system __manifests__ and __addons__
```
[vagrant@localhost ~]$ sudo mkdir -p /etc/kubernetes/manifests
[vagrant@localhost ~]$ sudo mkdir -p /etc/kubernetes/addons
```

### Certificates

Download [OpenVPN](https://github.com/OpenVPN/easy-rsa/releases) appliance
```
vagrant@localhost:~$ mkdir kube
vagrant@localhost:~$ curl -L -o /home/vagrant/kube/easy-rsa.tar.gz https://storage.googleapis.com/kubernetes-release/easy-rsa/easy-rsa.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 43444  100 43444    0     0  35046      0  0:00:01  0:00:01 --:--:-- 35035
vagrant@localhost:~$ ls -1 kube
easy-rsa.tar.gz
```

Generate, __Note: it is copyed from v1.5.7 hand-on, must change 172.17.4.200 -> 172.17.4.50 and 10.123.240.1 -> 10.3.0.1__
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo DEBUG='true' CERT_DIR=/root CERT_GROUP=root /opt/kubernetes/saltbase/salt/generate-cert/make-ca-cert.sh 172.17.4.200 IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local
+ cert_ip=172.17.4.200
+ extra_sans=IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local
+ cert_dir=/root
+ cert_group=root
+ mkdir -p /root
+ use_cn=false
+ '[' 172.17.4.200 == _use_gce_external_ip_ ']'
+ '[' 172.17.4.200 == _use_aws_external_ip_ ']'
+ '[' 172.17.4.200 == _use_azure_dns_name_ ']'
+ sans=IP:172.17.4.200
+ [[ -n IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local ]]
+ sans=IP:172.17.4.200,IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local
++ mktemp -d -t kubernetes_cacert.XXXXXX
+ tmpdir=/tmp/kubernetes_cacert.vxpao1
+ trap 'rm -rf "${tmpdir}"' EXIT
+ cd /tmp/kubernetes_cacert.vxpao1
+ '[' -f /home/vagrant/kube/easy-rsa.tar.gz ']'
+ curl -L -O https://storage.googleapis.com/kubernetes-release/easy-rsa/easy-rsa.tar.gz
+ tar xzf easy-rsa.tar.gz
+ cd easy-rsa-master/easyrsa3
+ ./easyrsa init-pki
++ date +%s
+ ./easyrsa --batch --req-cn=172.17.4.200@1493677386 build-ca nopass
+ '[' false = true ']'
+ ./easyrsa --subject-alt-name=IP:172.17.4.200,IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local build-server-full kubernetes-master nopass
+ cp -p pki/issued/kubernetes-master.crt /root/server.cert
+ cp -p pki/private/kubernetes-master.key /root/server.key
+ ./easyrsa build-client-full kubecfg nopass
+ cp -p pki/ca.crt /root/ca.crt
+ cp -p pki/issued/kubecfg.crt /root/kubecfg.crt
+ cp -p pki/private/kubecfg.key /root/kubecfg.key
+ chgrp root /root/server.key /root/server.cert /root/ca.crt
+ chmod 660 /root/server.key /root/server.cert /root/ca.crt
+ rm -rf /tmp/kubernetes_cacert.vxpao1
vagrant@vagrant-ubuntu-trusty-64:~$
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /root
ca.crt
kube
kubecfg.crt
kubecfg.key
server.cert
server.key
```

### Kubelet

Deploy __systemd__ service descriptor (maybe `sudo`)
```
[vagrant@localhost ~]$ sed -e 's%\(\(ExecStart=\)/usr/local/bin/kubelet\( "$DAEMON_ARGS"\)\)%# \1\n\2/opt/kubernetes/server/bin/kubelet\3%' /opt/kubernetes/cluster/saltbase/salt/kubelet/kubelet.service | sudo tee /etc/systemd/system/kubelet.service
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

Bin args: _Fedora23_
```
[vagrant@localhost ~]$ sed '/^# test_args.*$/,$!d;s/{{daemon_args}}//;s%{{api_servers_with_port}}%--api-servers=http://127.0.0.1:8080%;s/{{debugging_handlers}}/--enable-debugging-handlers=true/;s/{{hostname_override}}/--hostname-override=${KUBE_NODE_IP}/;s/{{cloud_provider}}//;s/{{cloud_config}}//;s%{{config}}%--pod-manifest-path=/etc/kubernetes/manifests%;s%{{manifest_url}}%%;s/{{pillar\[\x27allow_privileged\x27\]}}/true/;s/{{log_level}}/--v=2/;s/{{cluster_dns}}/--cluster-dns=10.3.0.10/;s/{{cluster_domain}}/--cluster-domain=cluster.local/;s%{{docker_root}}%%;s%{{kubelet_root}}%--root-dir=/var/lib/kubelet%;s%{{non_masquerade_cidr}}%--non-masquerade-cidr=10.0.0.0/8%;s%{{cgroup_root}}%%;s%{{system_container}}%%;s%{{pod_cidr}}%--pod-cidr=10.0.0.0/14%;s%{{ master_kubelet_args }}%--address=0.0.0.0 --authorization-mode=AlwaysAllow --node-ip=${KUBE_NODE_IP} --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0%;s/{{cpu_cfs_quota}}/--cpu-cfs-quota=true/;s/{{network_plugin}}//;s/{{kubelet_port}}/--port=10250/;s/{{ hairpin_mode }}//;s/{{enable_custom_metrics}}/--enable-custom-metrics=true/;s%{{runtime_container}}%%;s%{{kubelet_container}}%%;s/{{node_labels}}//;s/{{babysit_daemons}}//;s/{{eviction_hard}}/--eviction-hard=memory.available<100Mi/;s/{{kubelet_auth}}/--anonymous-auth=true/;s/{{feature_gates}}/--feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true/;s/{{test_args}}//' /opt/kubernetes/cluster/saltbase/salt/kubelet/default | env KUBE_NODE_IP=172.17.4.50 envsubst | sudo tee /etc/sysconfig/kubelet
# test_args has to be kept at the end, so they'll overwrite any prior configuration
DAEMON_ARGS=" --api-servers=http://127.0.0.1:8080 --enable-debugging-handlers=true --hostname-override=172.17.4.50   --pod-manifest-path=/etc/kubernetes/manifests  --allow-privileged=true --v=2 --cluster-dns=10.3.0.10 --cluster-domain=cluster.local  --root-dir=/var/lib/kubelet  --non-masquerade-cidr=10.0.0.0/8   --pod-cidr=10.0.0.0/14 --address=0.0.0.0 --authorization-mode=AlwaysAllow --node-ip=172.17.4.50 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --cpu-cfs-quota=true  --port=10250  --enable-custom-metrics=true     --eviction-hard=memory.available<100Mi --anonymous-auth=true --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true "
```

Start _kubelet_ service
```
[vagrant@localhost ~]$ sudo systemctl daemon-reload
[tangfx@localhost ~]$ sudo systemctl enable kubelet.service
[vagrant@localhost ~]$ sudo systemctl start kubelet.service
```

Debug: _The first time as it used invalid `cgroup_driver=systemd` args_
```
[vagrant@localhost ~]$ sudo systemctl status kubelet.service
● kubelet.service - Kubernetes Kubelet Server
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
   Active: activating (auto-restart) (Result: exit-code) since 六 2017-06-10 09:50:36 UTC; 340ms ago
     Docs: https://github.com/kubernetes/kubernetes
  Process: 11915 ExecStart=/opt/kubernetes/server/bin/kubelet $DAEMON_ARGS (code=exited, status=1/FAILURE)
 Main PID: 11915 (code=exited, status=1/FAILURE)

6月 10 09:50:36 localhost.localdomain systemd[1]: kubelet.service: Main process exited, code=exited, status=1/FAILURE
6月 10 09:50:36 localhost.localdomain systemd[1]: kubelet.service: Unit entered failed state.
6月 10 09:50:36 localhost.localdomain systemd[1]: kubelet.service: Failed with result 'exit-code'.
[vagrant@localhost ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service
6月 10 09:52:30 localhost.localdomain systemd[1]: Started Kubernetes Kubelet Server.
6月 10 09:52:30 localhost.localdomain systemd[1]: Starting Kubernetes Kubelet Server...
6月 10 09:52:30 localhost.localdomain kubelet[12803]: Flag --api-servers has been deprecated, Use --kubeconfig instead. Will be removed in a future version.
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.447092   12803 feature_gate.go:144] feature gates: map[Accelerators:true AffinityInAnnotations:true AllAlpha:true DynamicKubeletConfig:true ExperimentalCriticalPodAnnotation:true TaintBasedEvictions:true DynamicVolumeProvisioning:true ExperimentalHostUserNamespaceDefaulting:true]
6月 10 09:52:30 localhost.localdomain kubelet[12803]: W0610 09:52:30.448123   12803 server.go:715] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.454794   12803 server.go:232] Starting Kubelet configuration sync loop
6月 10 09:52:30 localhost.localdomain kubelet[12803]: E0610 09:52:30.454836   12803 server.go:407] failed to init dynamic Kubelet configuration sync: cloud provider was nil, and attempt to use hostname to find config resulted in: Get http://127.0.0.1:8080/api/v1/namespaces/kube-system/configmaps/kubelet-172.17.4.50: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 09:52:30 localhost.localdomain kubelet[12803]: W0610 09:52:30.454926   12803 server.go:715] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.456192   12803 docker.go:364] Connecting to docker on unix:///var/run/docker.sock
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.456292   12803 docker.go:384] Start docker client with request timeout=2m0s
6月 10 09:52:30 localhost.localdomain kubelet[12803]: W0610 09:52:30.470780   12803 cni.go:157] Unable to update cni config: No networks found in /etc/cni/net.d
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.486634   12803 manager.go:143] cAdvisor running in container: "/"
6月 10 09:52:30 localhost.localdomain kubelet[12803]: W0610 09:52:30.672309   12803 manager.go:151] unable to connect to Rkt api service: rkt: cannot tcp Dial rkt api service: dial tcp [::1]:15441: getsockopt: connection refused
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.855252   12803 fs.go:117] Filesystem partitions: map[/dev/mapper/vg_vagrant-lv_root:{mountpoint:/ major:253 minor:0 fsType:ext4 blockSize:0} /dev/sda1:{mountpoint:/boot major:8 minor:1 fsType:ext4 blockSize:0} docker-253:0-1310998-pool:{mountpoint: major:253 minor:2 fsType:devicemapper blockSize:128}]
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.863347   12803 manager.go:198] Machine: {NumCores:2 CpuFrequency:2700107 MemoryCapacity:5201813504 MachineID:5c949bb3146241e09f7e671b0704d4fb SystemUUID:284705D4-2625-4A04-8B45-675B98E6E998 BootID:48d69df5-2b4b-4136-af94-1dfe72b1fea3 Filesystems:[{Device:/dev/mapper/vg_vagrant-lv_root Capacity:40522645504 Type:vfs Inodes:2523136 HasInodes:true} {Device:/dev/sda1 Capacity:499355648 Type:vfs Inodes:128016 HasInodes:true} {Device:docker-253:0-1310998-pool Capacity:34359738368 Type:devicemapper Inodes:0 HasInodes:false}] DiskMap:map[253:3:{Name:dm-3 Major:253 Minor:3 Size:2147483648 Scheduler:none} 253:4:{Name:dm-4 Major:253 Minor:4 Size:34359738368 Scheduler:none} 8:0:{Name:sda Major:8 Minor:0 Size:85899345920 Scheduler:cfq} 253:0:{Name:dm-0 Major:253 Minor:0 Size:41305505792 Scheduler:none} 253:1:{Name:dm-1 Major:253 Minor:1 Size:1107296256 Scheduler:none} 253:2:{Name:dm-2 Major:253 Minor:2 Size:34359738368 Scheduler:none}] NetworkDevices:[{Name:br-3ba1465afaf1 MacAddress:02:42:41:c7:e0:a2 Speed:0 Mtu:1500} {Name:eth0 MacAddress:08:00:27:46:54:e7 Speed:1000 Mtu:1500} {Name:eth1 MacAddress:08:00:27:5a:1a:a4 Speed:1000 Mtu:1500}] Topology:[{Id:0 Memory:5201813504 Cores:[{Id:0 Threads:[0] Caches:[{Size:32768 Type:Data Level:1} {Size:32768 Type:Instruction Level:1} {Size:262144 Type:Unified Level:2} {Size:3145728 Type:Unified Level:3}]} {Id:1 Threads:[1] Caches:[{Size:32768 Type:Data Level:1} {Size:32768 Type:Instruction Level:1} {Size:262144 Type:Unified Level:2} {Size:3145728 Type:Unified Level:3}]}] Caches:[]}] CloudProvider:Unknown InstanceType:Unknown InstanceID:None}
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.874899   12803 manager.go:204] Version: {KernelVersion:4.2.3-300.fc23.x86_64 ContainerOsVersion:Fedora 23 (Twenty Three) DockerVersion:1.9.1 CadvisorVersion: CadvisorRevision:}
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.877976   12803 server.go:509] --cgroups-per-qos enabled, but --cgroup-root was not specified.  defaulting to /
6月 10 09:52:30 localhost.localdomain kubelet[12803]: W0610 09:52:30.880711   12803 container_manager_linux.go:218] Running with swap on is not supported, please disable swap! This will be a fatal error by default starting in K8s v1.6! In the meantime, you can opt-in to making this a fatal error by enabling --experimental-fail-swap-on.
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.881245   12803 container_manager_linux.go:245] container manager verified user specified cgroup-root exists: /
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.881349   12803 container_manager_linux.go:250] Creating Container Manager object based on Node Config: {RuntimeCgroupsName: SystemCgroupsName: KubeletCgroupsName: ContainerRuntime:docker CgroupsPerQOS:true CgroupRoot:/ CgroupDriver:systemd ProtectKernelDefaults:false EnableCRI:true NodeAllocatableConfig:{KubeReservedCgroupName: SystemReservedCgroupName: EnforceNodeAllocatable:map[pods:{}] KubeReserved:map[] SystemReserved:map[] HardEvictionThresholds:[{Signal:memory.available Operator:LessThan Value:{Quantity:100Mi Percentage:0} GracePeriod:0s MinReclaim:<nil>}]} ExperimentalQOSReserved:map[]}
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.881662   12803 server.go:802] Using root directory: /var/lib/kubelet
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.881843   12803 kubelet.go:255] Adding manifest file: /etc/kubernetes/manifests
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.882147   12803 file.go:48] Watching path "/etc/kubernetes/manifests"
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.882153   12803 kubelet.go:265] Watching apiserver
6月 10 09:52:30 localhost.localdomain kubelet[12803]: E0610 09:52:30.884990   12803 file.go:72] unable to read config path "/etc/kubernetes/manifests": path does not exist, ignoring
6月 10 09:52:30 localhost.localdomain kubelet[12803]: E0610 09:52:30.888215   12803 reflector.go:190] k8s.io/kubernetes/pkg/kubelet/kubelet.go:390: Failed to list *v1.Node: Get http://127.0.0.1:8080/api/v1/nodes?fieldSelector=metadata.name%3D172.17.4.50&resourceVersion=0: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 09:52:30 localhost.localdomain kubelet[12803]: E0610 09:52:30.888477   12803 reflector.go:190] k8s.io/kubernetes/pkg/kubelet/kubelet.go:382: Failed to list *v1.Service: Get http://127.0.0.1:8080/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 09:52:30 localhost.localdomain kubelet[12803]: E0610 09:52:30.889620   12803 reflector.go:190] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:46: Failed to list *v1.Pod: Get http://127.0.0.1:8080/api/v1/pods?fieldSelector=spec.nodeName%3D172.17.4.50&resourceVersion=0: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.890577   12803 kubelet.go:484] Experimental host user namespace defaulting is enabled.
6月 10 09:52:30 localhost.localdomain kubelet[12803]: W0610 09:52:30.890737   12803 kubelet_network.go:70] Hairpin mode set to "promiscuous-bridge" but kubenet is not enabled, falling back to "hairpin-veth"
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.890759   12803 kubelet.go:494] Hairpin mode set to "hairpin-veth"
6月 10 09:52:30 localhost.localdomain kubelet[12803]: W0610 09:52:30.896850   12803 cni.go:157] Unable to update cni config: No networks found in /etc/cni/net.d
6月 10 09:52:30 localhost.localdomain kubelet[12803]: I0610 09:52:30.907458   12803 docker_service.go:187] Docker cri networking managed by kubernetes.io/no-op
6月 10 09:52:31 localhost.localdomain kubelet[12803]: W0610 09:52:31.083945   12803 docker_service.go:196] No cgroup driver is set in Docker
6月 10 09:52:31 localhost.localdomain kubelet[12803]: W0610 09:52:31.083965   12803 docker_service.go:197] Falling back to use the default driver: "cgroupfs"
6月 10 09:52:31 localhost.localdomain kubelet[12803]: error: failed to run Kubelet: failed to create kubelet: misconfiguration: kubelet cgroup driver: "systemd" is different from docker cgroup driver: "cgroupfs"
6月 10 09:52:31 localhost.localdomain systemd[1]: kubelet.service: Main process exited, code=exited, status=1/FAILURE
6月 10 09:52:31 localhost.localdomain systemd[1]: kubelet.service: Unit entered failed state.
6月 10 09:52:31 localhost.localdomain systemd[1]: kubelet.service: Failed with result 'exit-code'.
```

Check
```
[vagrant@localhost ~]$ sudo systemctl restart kubelet.service
[vagrant@localhost ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - Kubernetes Kubelet Server
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
   Active: active (running) since 六 2017-06-10 09:59:25 UTC; 45s ago
     Docs: https://github.com/kubernetes/kubernetes
 Main PID: 15727 (kubelet)
   CGroup: /system.slice/kubelet.service
           ├─15681 journalctl -k -f
           ├─15727 /opt/kubernetes/server/bin/kubelet --api-servers=http://127.0.0.1:8080 --enable-debugging-handlers=true --hostname-override=172.17.4.50 --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --v=2 --cluster-dns=10.123.240.10 --cluster-domain=cluster.local --root-dir=/var/lib/kubelet --non-masquerade-cidr=10.0.0.0/8 --pod-cidr=10.120.0.0/14 --address=0.0.0.0 --authorization-mode=AlwaysAllow --node-ip=172.17.4.50 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --cpu-cfs-quota=true --port=10250 --enable-custom-metrics=true --eviction-hard=memory.available<100Mi --anonymous-auth=true --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true
           └─15771 journalctl -k -f

6月 10 10:00:07 localhost.localdomain kubelet[15727]: I0610 10:00:07.922887   15727 kubelet_node_status.go:379] Recording NodeHasSufficientMemory event message for node 172.17.4.50
6月 10 10:00:07 localhost.localdomain kubelet[15727]: I0610 10:00:07.922901   15727 kubelet_node_status.go:379] Recording NodeHasNoDiskPressure event message for node 172.17.4.50
6月 10 10:00:07 localhost.localdomain kubelet[15727]: I0610 10:00:07.922919   15727 kubelet_node_status.go:77] Attempting to register node 172.17.4.50
6月 10 10:00:07 localhost.localdomain kubelet[15727]: E0610 10:00:07.923223   15727 kubelet_node_status.go:101] Unable to register node "172.17.4.50" with API server: Post http://127.0.0.1:8080/api/v1/nodes: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 10:00:08 localhost.localdomain kubelet[15727]: E0610 10:00:08.691322   15727 reflector.go:190] k8s.io/kubernetes/pkg/kubelet/kubelet.go:382: Failed to list *v1.Service: Get http://127.0.0.1:8080/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 10:00:08 localhost.localdomain kubelet[15727]: E0610 10:00:08.692569   15727 reflector.go:190] k8s.io/kubernetes/pkg/kubelet/kubelet.go:390: Failed to list *v1.Node: Get http://127.0.0.1:8080/api/v1/nodes?fieldSelector=metadata.name%3D172.17.4.50&resourceVersion=0: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 10:00:08 localhost.localdomain kubelet[15727]: E0610 10:00:08.703838   15727 reflector.go:190] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:46: Failed to list *v1.Pod: Get http://127.0.0.1:8080/api/v1/pods?fieldSelector=spec.nodeName%3D172.17.4.50&resourceVersion=0: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 10:00:09 localhost.localdomain kubelet[15727]: E0610 10:00:09.694760   15727 reflector.go:190] k8s.io/kubernetes/pkg/kubelet/kubelet.go:390: Failed to list *v1.Node: Get http://127.0.0.1:8080/api/v1/nodes?fieldSelector=metadata.name%3D172.17.4.50&resourceVersion=0: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 10:00:09 localhost.localdomain kubelet[15727]: E0610 10:00:09.695142   15727 reflector.go:190] k8s.io/kubernetes/pkg/kubelet/kubelet.go:382: Failed to list *v1.Service: Get http://127.0.0.1:8080/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:8080: getsockopt: connection refused
6月 10 10:00:09 localhost.localdomain kubelet[15727]: E0610 10:00:09.705506   15727 reflector.go:190] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:46: Failed to list *v1.Pod: Get http://127.0.0.1:8080/api/v1/pods?fieldSelector=spec.nodeName%3D172.17.4.50&resourceVersion=0: dial tcp 127.0.0.1:8080: getsockopt: connection refused
```

### Etcd

Log file 
```
[vagrant@localhost ~]$ sudo touch /var/log/etcd.log
```

Certificates
```
[vagrant@localhost ~]$ sudo cp /data/src/github.com/openshift/origin/etc/kubernetes/cacerts/server.cert /srv/kubernetes/etcd-peer.cert
[vagrant@localhost ~]$ sudo cp /data/src/github.com/openshift/origin/etc/kubernetes/cacerts/server.key /srv/kubernetes/etcd-peer.key
[vagrant@localhost ~]$ sudo cp /data/src/github.com/openshift/origin/etc/kubernetes/cacerts/ca.crt /srv/kubernetes/etcd-ca.crt
```

Manifest
```
[vagrant@localhost ~]$ sed '/{[%#].*[%#]}/d;s/{{ suffix }}//g;s/{{ pillar.get(\x27etcd_docker_tag\x27, \x27\(.*\)\x27) }}/\1/;s/{{ cpulimit }}/"250m"/;s/{{ hostname }}/${KUBE_ETCD_HOST}/g;s/{{ etcd_protocol }}/https/g;s/{{ server_port }}/2380/g;s/{{ port }}/2379/g;s/{{ quota_bytes }}/--quota-backend-bytes 4294967296/;s/{{ cluster_state }}/new/;s%{{ etcd_cluster }}%etcd-${KUBE_ETCD_HOST}=https://${KUBE_ETCD_HOST}:2380%;s%{{ etcd_creds }}%--peer-trusted-ca-file /srv/kubernetes/etcd-ca.crt --peer-cert-file /srv/kubernetes/etcd-peer.cert --peer-key-file /srv/kubernetes/etcd-peer.key -peer-client-cert-auth%;s/{{ pillar.get(\x27storage_backend\x27, \x27etcd3\x27) }}/etcd3/;s/{{ pillar.get(\x27etcd_version\x27, \x27\(.*\)\x27) }}/\1/;s%{{ srv_kube_path }}%/srv/kubernetes%' /opt/kubernetes/cluster/saltbase/salt/etcd/etcd.manifest | KUBE_ETCD_HOST=172.17.4.50 envsubst | sudo tee /etc/kubernetes/manifests/etcd.json

{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"etcd-server",
  "namespace": "kube-system"
},
"spec":{
"hostNetwork": true,
"containers":[
    {
    "name": "etcd-container",
    "image": "gcr.io/google_containers/etcd:3.0.17",
    "resources": {
      "requests": {
        "cpu": "250m"
      }
    },
    "command": [
              "/bin/sh",
              "-c",
              "if [ -e /usr/local/bin/migrate-if-needed.sh ]; then /usr/local/bin/migrate-if-needed.sh 1>>/var/log/etcd.log 2>&1; fi; /usr/local/bin/etcd --name etcd-172.17.4.50 --listen-peer-urls https://172.17.4.50:2380 --initial-advertise-peer-urls https://172.17.4.50:2380 --advertise-client-urls http://127.0.0.1:2379 --listen-client-urls http://127.0.0.1:2379 --quota-backend-bytes 4294967296 --data-dir /var/etcd/data --initial-cluster-state new --initial-cluster etcd-172.17.4.50=https://172.17.4.50:2380 --peer-trusted-ca-file /srv/kubernetes/etcd-ca.crt --peer-cert-file /srv/kubernetes/etcd-peer.cert --peer-key-file /srv/kubernetes/etcd-peer.key -peer-client-cert-auth 1>>/var/log/etcd.log 2>&1"
            ],
    "env": [
      { "name": "TARGET_STORAGE",
        "value": "etcd3"
      },
      { "name": "TARGET_VERSION",
        "value": "3.0.17"
      },
      { "name": "DATA_DIRECTORY",
        "value": "/var/etcd/data"
      }
        ],
    "livenessProbe": {
      "httpGet": {
        "host": "127.0.0.1",
        "port": 2379,
        "path": "/health"
      },
      "initialDelaySeconds": 15,
      "timeoutSeconds": 15
    },
    "ports": [
      { "name": "serverport",
        "containerPort": 2380,
        "hostPort": 2380 
      },
      { "name": "clientport",
        "containerPort": 2379,
        "hostPort": 2379
      }
        ],
    "volumeMounts": [
      { "name": "varetcd",
        "mountPath": "/var/etcd",
        "readOnly": false
      },
      { "name": "varlogetcd",
        "mountPath": "/var/log/etcd.log",
        "readOnly": false
      },
      { "name": "etc",
        "mountPath": "/srv/kubernetes",
        "readOnly": false
      }
    ]
    }
],
"volumes":[
  { "name": "varetcd",
    "hostPath": {
        "path": "/mnt/master-pd/var/etcd"}
  },
  { "name": "varlogetcd",
    "hostPath": {
        "path": "/var/log/etcd.log"}
  },
  { "name": "etc",
    "hostPath": {
        "path": "/srv/kubernetes"}
  }
]
}}
```

Restart _kubelet_ (maybe)
```
[vagrant@localhost ~]$ sudo systemctl restart kubelet.service
```

Debug: _As it indicate to update docker engine_
```
[vagrant@localhost ~]$ sudo journalctl --no-tail --no-pager --pager-end --unit kubelet.service | grep -e "--security-opt"
6月 10 12:08:26 localhost.localdomain kubelet[28846]: E0610 12:08:26.708018   28846 remote_runtime.go:86] RunPodSandbox from runtime service failed: rpc error: code = 2 desc = failed to create a sandbox for pod "etcd-server-172.17.4.50": Error response from daemon: Invalid --security-opt: "seccomp:unconfined"
6月 10 12:08:26 localhost.localdomain kubelet[28846]: E0610 12:08:26.708078   28846 kuberuntime_sandbox.go:54] CreatePodSandbox for pod "etcd-server-172.17.4.50_kube-system(92ea623866c62d0d58056ae7081c0fbe)" failed: rpc error: code = 2 desc = failed to create a sandbox for pod "etcd-server-172.17.4.50": Error response from daemon: Invalid --security-opt: "seccomp:unconfined"
6月 10 12:08:26 localhost.localdomain kubelet[28846]: E0610 12:08:26.708095   28846 kuberuntime_manager.go:619] createPodSandbox for pod "etcd-server-172.17.4.50_kube-system(92ea623866c62d0d58056ae7081c0fbe)" failed: rpc error: code = 2 desc = failed to create a sandbox for pod "etcd-server-172.17.4.50": Error response from daemon: Invalid --security-opt: "seccomp:unconfined"
6月 10 12:08:26 localhost.localdomain kubelet[28846]: E0610 12:08:26.708126   28846 pod_workers.go:182] Error syncing pod 92ea623866c62d0d58056ae7081c0fbe ("etcd-server-172.17.4.50_kube-system(92ea623866c62d0d58056ae7081c0fbe)"), skipping: failed to "CreatePodSandbox" for "etcd-server-172.17.4.50_kube-system(92ea623866c62d0d58056ae7081c0fbe)" with CreatePodSandboxError: "CreatePodSandbox for pod \"etcd-server-172.17.4.50_kube-system(92ea623866c62d0d58056ae7081c0fbe)\" failed: rpc error: code = 2 desc = failed to create a sandbox for pod \"etcd-server-172.17.4.50\": Error response from daemon: Invalid --security-opt: \"seccomp:unconfined\""
```

Check: _After update docker version_
```
[vagrant@localhost ~]$ docker ps
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
67f96a93ac02        e6a6589433f4                               "/bin/sh -c 'if [ -e "   3 minutes ago       Up 3 minutes                            k8s_etcd-container_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
d67c2281716d        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 3 minutes ago       Up 3 minutes                            k8s_POD_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
[vagrant@localhost ~]$ etcdctl member list
8d07311ffdd56e6f: name=etcd-172.17.4.50 peerURLs=https://172.17.4.50:2380 clientURLs=http://127.0.0.1:2379 isLeader=true
[vagrant@localhost ~]$ ETCDCTL_API=3 etcdctl member list
8d07311ffdd56e6f, started, etcd-172.17.4.50, https://172.17.4.50:2380, http://127.0.0.1:2379
[vagrant@localhost ~]$ ETCDCTL_API=3 etcdctl put foo bar
OK
[vagrant@localhost ~]$ ETCDCTL_API=3 etcdctl get foo
foo
bar
[vagrant@localhost ~]$ ETCDCTL_API=3 etcdctl del foo
1
```

Check log
```
[vagrant@localhost ~]$ tail /var/log/etcd.log 
2017-06-10 19:18:34.993058 I | raft: 8d07311ffdd56e6f became candidate at term 2
2017-06-10 19:18:34.993077 I | raft: 8d07311ffdd56e6f received vote from 8d07311ffdd56e6f at term 2
2017-06-10 19:18:34.993096 I | raft: 8d07311ffdd56e6f became leader at term 2
2017-06-10 19:18:34.993105 I | raft: raft.node: 8d07311ffdd56e6f elected leader 8d07311ffdd56e6f at term 2
2017-06-10 19:18:34.993426 I | etcdserver: setting up the initial cluster version to 3.0
2017-06-10 19:18:34.993912 N | membership: set the initial cluster version to 3.0
2017-06-10 19:18:34.993962 I | api: enabled capabilities for version 3.0
2017-06-10 19:18:34.994020 I | etcdserver: published {Name:etcd-172.17.4.50 ClientURLs:[http://127.0.0.1:2379]} to cluster 9b9ee104470b136d
2017-06-10 19:18:34.994029 I | etcdmain: ready to serve client requests
2017-06-10 19:18:34.994309 N | etcdmain: serving insecure client requests on 127.0.0.1:2379, this is strongly discouraged!
```

### Kube-apiserver

Log & key loacation
```
[vagrant@localhost ~]$ sudo touch /var/log/kube-apiserver.log && sudo touch /var/log/kube-apiserver-audit.log
[vagrant@localhost ~]$ sudo mkdir -p /srv/sshproxy /etc/srv/pki /etc/openssl /var/ssl
```

Certificates
```
[vagrant@localhost ~]$ sudo cp /data/src/github.com/openshift/origin/etc/kubernetes/cacerts/ca.crt /srv/kubernetes/
[vagrant@localhost ~]$ sudo cp /data/src/github.com/openshift/origin/etc/kubernetes/cacerts/server.cert /srv/kubernetes/
[vagrant@localhost ~]$ sudo cp /data/src/github.com/openshift/origin/etc/kubernetes/cacerts/server.key /srv/kubernetes/
[vagrant@localhost ~]$ sudo cp /srv/kubernetes/server.cert /srv/kubernetes/kubeapiserver.cert
[vagrant@localhost ~]$ sudo cp /srv/kubernetes/server.key /srv/kubernetes/kubeapiserver.key
```

Manifest
```
[vagrant@localhost ~]$ sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%\("image": \)".*kube_docker_registry\x27\x5d\x7d\x7d\(/kube-apiserver:\).*kube-apiserver_docker_tag.*"%\1"gcr.io/google_containers\2$KUBE_APISERVER_DOCKERTAG"%;s%{{params}}%--address=127.0.0.1 --storage-backend=etcd3 --storage-media-type=application/vnd.kubernetes.protobuf --etcd-servers=http://127.0.0.1:2379 --etcd-servers-overrides=/events#http://127.0.0.1:2379 --runtime-config=api/all=true --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true --admission-control=ResourceQuota,ServiceAccount --max-requests-inflight=400 --target-ram-mb=1024 --service-cluster-ip-range=10.3.0.0/20 --client-ca-file=/srv/kubernetes/ca.crt --min-request-timeout=1800 --enable-garbage-collector=true --etcd-quorum-read=true --audit-log-path=/var/log/kube-apiserver-audit.log --audit-log-maxage=0 --audit-log-maxbackup=0 --audit-log-maxsize=2000000000 --tls-cert-file=/srv/kubernetes/server.cert --tls-private-key-file=/srv/kubernetes/server.key --kubelet-client-certificate=/srv/kubernetes/kubeapiserver.cert --kubelet-client-key=/srv/kubernetes/kubeapiserver.key --secure-port=6443 --token-auth-file=/dev/null --bind-address=0.0.0.0 --v=2 --advertise-address=$KUBE_APISERVER_HOST --enable-swagger-ui=true%;s/{{pillar.*allow_privileged.*}}/true/;s/{{container_env}}//;s/{{secure_port}}/6443/g;s/{{cloud_config_mount}}\|{{additional_cloud_config_mount}}\|{{webhook_config_mount}}\|{{webhook_authn_config_mount}}\|{{admission_controller_config_mount}}\|{{image_policy_webhook_config_mount}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;s%{{srv_sshproxy_path}}%/srv/sshproxy%g;s/{{cloud_config_volume}}\|{{additional_cloud_config_volume}}\|{{webhook_config_volume}}\|{{webhook_authn_config_volume}}\|{{admission_controller_config_volume}}\|{{image_policy_webhook_config_volume}}//;s%"path": "/etc/srv/pki"%"path": "/etc/pki"%' /opt/kubernetes/cluster/saltbase/salt/kube-apiserver/kube-apiserver.manifest | KUBE_APISERVER_DOCKERTAG=$(cat /opt/kubernetes/server/bin/kube-apiserver.docker_tag) KUBE_APISERVER_HOST=172.17.4.50 envsubst | sudo tee /etc/kubernetes/manifests/kube-apiserver.json


{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"kube-apiserver",
  "namespace": "kube-system",
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
    "image": "gcr.io/google_containers/kube-apiserver:41a3581e4847878bbcb1722636ab51ad",
    "resources": {
      "requests": {
        "cpu": "250m"
      }
    },
    "command": [
                 "/bin/sh",
                 "-c",
                 "/usr/local/bin/kube-apiserver --address=127.0.0.1 --storage-backend=etcd3 --storage-media-type=application/vnd.kubernetes.protobuf --etcd-servers=http://127.0.0.1:2379 --etcd-servers-overrides=/events#http://127.0.0.1:2379 --runtime-config=api/all=true --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true --admission-control=ResourceQuota,ServiceAccount --max-requests-inflight=400 --target-ram-mb=1024 --service-cluster-ip-range=10.3.0.0/20 --client-ca-file=/srv/kubernetes/ca.crt --min-request-timeout=1800 --enable-garbage-collector=true --etcd-quorum-read=true --audit-log-path=/var/log/kube-apiserver-audit.log --audit-log-maxage=0 --audit-log-maxbackup=0 --audit-log-maxsize=2000000000 --tls-cert-file=/srv/kubernetes/server.cert --tls-private-key-file=/srv/kubernetes/server.key --kubelet-client-certificate=/srv/kubernetes/kubeapiserver.cert --kubelet-client-key=/srv/kubernetes/kubeapiserver.key --secure-port=6443 --token-auth-file=/dev/null --bind-address=0.0.0.0 --v=2 --advertise-address=172.17.4.50 --enable-swagger-ui=true --allow-privileged=true 1>>/var/log/kube-apiserver.log 2>&1"
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

Debug
```
[vagrant@localhost ~]$ docker ps
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
d03bfed86480        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 6 minutes ago       Up 6 minutes                            k8s_POD_kube-apiserver-172.17.4.50_kube-system_e81a8e52b045d83ee62893b4c4033354_0
67f96a93ac02        e6a6589433f4                               "/bin/sh -c 'if [ -e "   2 hours ago         Up 2 hours                              k8s_etcd-container_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
d67c2281716d        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 2 hours ago         Up 2 hours                              k8s_POD_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
[vagrant@localhost ~]$ ls -l /var/log/kube-apiserver*
-rw-r--r--. 1 root root    0 6月  10 20:02 /var/log/kube-apiserver-audit.log
-rw-r--r--. 1 root root 3395 6月  10 22:18 /var/log/kube-apiserver.log
```

After correct
```
[vagrant@localhost ~]$ sudo systemctl restart kubelet.service 
```

Check
```
[vagrant@localhost ~]$ docker ps
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
4871d8978295        9196766fca5a                               "/bin/sh -c '/usr/loc"   2 minutes ago       Up 2 minutes                            k8s_kube-apiserver_kube-apiserver-172.17.4.50_kube-system_75a081d878fbde18ca953b4426639b8d_4
575c3883f38a        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_kube-apiserver-172.17.4.50_kube-system_75a081d878fbde18ca953b4426639b8d_0
67f96a93ac02        e6a6589433f4                               "/bin/sh -c 'if [ -e "   3 hours ago         Up 3 hours                              k8s_etcd-container_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
d67c2281716d        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 3 hours ago         Up 3 hours                              k8s_POD_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
```

Check log
```
[vagrant@localhost ~]$ ls -l /var/log/kube-apiserver*
-rw-r--r--. 1 root root 414811 6月  10 23:20 /var/log/kube-apiserver-audit.log
-rw-r--r--. 1 root root  84556 6月  10 23:20 /var/log/kube-apiserver.log
[vagrant@localhost ~]$ tail /var/log/kube-apiserver.log 
[vagrant@localhost ~]$ tail /var/log/kube-apiserver-audit.log 
```

Command `curl`
```
[vagrant@localhost ~]$ sudo netstat -tpnl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 127.0.0.1:10248         0.0.0.0:*               LISTEN      5927/kubelet        
tcp        0      0 127.0.0.1:2379          0.0.0.0:*               LISTEN      22694/etcd          
tcp        0      0 172.17.4.50:2380        0.0.0.0:*               LISTEN      22694/etcd          
tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      6070/kube-apiserver 
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      735/sshd            
tcp6       0      0 :::4194                 :::*                    LISTEN      5927/kubelet        
tcp6       0      0 :::10250                :::*                    LISTEN      5927/kubelet        
tcp6       0      0 :::6443                 :::*                    LISTEN      6070/kube-apiserver 
tcp6       0      0 :::10255                :::*                    LISTEN      5927/kubelet        
tcp6       0      0 :::22                   :::*                    LISTEN      735/sshd            
[vagrant@localhost ~]$ curl http://127.0.0.1:8080
{
  "paths": [
    "/api",
    "/api/v1",
    "/apis",
    "/apis/apps",
    "/apis/apps/v1beta1",
    "/apis/authentication.k8s.io",
    "/apis/authentication.k8s.io/v1",
    "/apis/authentication.k8s.io/v1beta1",
    "/apis/authorization.k8s.io",
    "/apis/authorization.k8s.io/v1",
    "/apis/authorization.k8s.io/v1beta1",
    "/apis/autoscaling",
    "/apis/autoscaling/v1",
    "/apis/autoscaling/v2alpha1",
    "/apis/batch",
    "/apis/batch/v1",
    "/apis/batch/v2alpha1",
    "/apis/certificates.k8s.io",
    "/apis/certificates.k8s.io/v1beta1",
    "/apis/extensions",
    "/apis/extensions/v1beta1",
    "/apis/policy",
    "/apis/policy/v1beta1",
    "/apis/rbac.authorization.k8s.io",
    "/apis/rbac.authorization.k8s.io/v1alpha1",
    "/apis/rbac.authorization.k8s.io/v1beta1",
    "/apis/settings.k8s.io",
    "/apis/settings.k8s.io/v1alpha1",
    "/apis/storage.k8s.io",
    "/apis/storage.k8s.io/v1",
    "/apis/storage.k8s.io/v1beta1",
    "/healthz",
    "/healthz/ping",
    "/healthz/poststarthook/bootstrap-controller",
    "/healthz/poststarthook/ca-registration",
    "/healthz/poststarthook/extensions/third-party-resources",
    "/logs",
    "/metrics",
    "/swagger-ui/",
    "/swaggerapi/",
    "/ui/",
    "/version"
  ]
}
[vagrant@localhost ~]$ curl --capath /srv/kubernetes --cacert /srv/kubernetes/ca.crt --cert /data/src/github.com/openshift/origin/etc/kubernetes/caerts/kubecfg.crt --key /data/src/github.com/openshift/origin/etc/kubernetes/cacerts/kubecfg.key https://172.17.4.50:6443/api
{
  "kind": "APIVersions",
  "versions": [
    "v1"
  ],
  "serverAddressByClientCIDRs": [
    {
      "clientCIDR": "0.0.0.0/0",
      "serverAddress": "172.17.4.50:6443"
    }
  ]
}
[vagrant@localhost ~]$ curl http://127.0.0.1:8080/apis
{
  "kind": "APIGroupList",
  "groups": [
    {
      "name": "authentication.k8s.io",
      "versions": [
        {
          "groupVersion": "authentication.k8s.io/v1",
          "version": "v1"
        },
        {
          "groupVersion": "authentication.k8s.io/v1beta1",
          "version": "v1beta1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "authentication.k8s.io/v1",
        "version": "v1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "authorization.k8s.io",
      "versions": [
        {
          "groupVersion": "authorization.k8s.io/v1",
          "version": "v1"
        },
        {
          "groupVersion": "authorization.k8s.io/v1beta1",
          "version": "v1beta1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "authorization.k8s.io/v1",
        "version": "v1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "autoscaling",
      "versions": [
        {
          "groupVersion": "autoscaling/v1",
          "version": "v1"
        },
        {
          "groupVersion": "autoscaling/v2alpha1",
          "version": "v2alpha1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "autoscaling/v2alpha1",
        "version": "v2alpha1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "batch",
      "versions": [
        {
          "groupVersion": "batch/v1",
          "version": "v1"
        },
        {
          "groupVersion": "batch/v2alpha1",
          "version": "v2alpha1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "batch/v1",
        "version": "v1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "certificates.k8s.io",
      "versions": [
        {
          "groupVersion": "certificates.k8s.io/v1beta1",
          "version": "v1beta1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "certificates.k8s.io/v1beta1",
        "version": "v1beta1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "extensions",
      "versions": [
        {
          "groupVersion": "extensions/v1beta1",
          "version": "v1beta1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "extensions/v1beta1",
        "version": "v1beta1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "policy",
      "versions": [
        {
          "groupVersion": "policy/v1beta1",
          "version": "v1beta1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "policy/v1beta1",
        "version": "v1beta1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "rbac.authorization.k8s.io",
      "versions": [
        {
          "groupVersion": "rbac.authorization.k8s.io/v1beta1",
          "version": "v1beta1"
        },
        {
          "groupVersion": "rbac.authorization.k8s.io/v1alpha1",
          "version": "v1alpha1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "rbac.authorization.k8s.io/v1beta1",
        "version": "v1beta1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "settings.k8s.io",
      "versions": [
        {
          "groupVersion": "settings.k8s.io/v1alpha1",
          "version": "v1alpha1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "settings.k8s.io/v1alpha1",
        "version": "v1alpha1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "storage.k8s.io",
      "versions": [
        {
          "groupVersion": "storage.k8s.io/v1beta1",
          "version": "v1beta1"
        },
        {
          "groupVersion": "storage.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "storage.k8s.io/v1",
        "version": "v1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    },
    {
      "name": "apps",
      "versions": [
        {
          "groupVersion": "apps/v1beta1",
          "version": "v1beta1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "apps/v1beta1",
        "version": "v1beta1"
      },
      "serverAddressByClientCIDRs": [
        {
          "clientCIDR": "0.0.0.0/0",
          "serverAddress": "172.17.4.50:6443"
        }
      ]
    }
  ]
}
[vagrant@localhost ~]$ curl http://127.0.0.1:8080/api/v1/nodes
{
  "kind": "NodeList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/nodes",
    "resourceVersion": "888"
  },
  "items": [
    {
      "metadata": {
        "name": "172.17.4.50",
        "selfLink": "/api/v1/nodes172.17.4.50",
        "uid": "4dd88960-4e32-11e7-8689-0800274654e7",
        "resourceVersion": "888",
        "creationTimestamp": "2017-06-10T23:12:41Z",
        "labels": {
          "beta.kubernetes.io/arch": "amd64",
          "beta.kubernetes.io/os": "linux",
          "kubernetes.io/hostname": "172.17.4.50"
        },
        "annotations": {
          "volumes.kubernetes.io/controller-managed-attach-detach": "true"
        }
      },
      "spec": {
        "externalID": "172.17.4.50"
      },
      "status": {
        "capacity": {
          "alpha.kubernetes.io/nvidia-gpu": "0",
          "cpu": "2",
          "memory": "5079896Ki",
          "pods": "110"
        },
        "allocatable": {
          "alpha.kubernetes.io/nvidia-gpu": "0",
          "cpu": "2",
          "memory": "4977496Ki",
          "pods": "110"
        },
        "conditions": [
          {
            "type": "OutOfDisk",
            "status": "False",
            "lastHeartbeatTime": "2017-06-10T23:22:49Z",
            "lastTransitionTime": "2017-06-10T23:12:41Z",
            "reason": "KubeletHasSufficientDisk",
            "message": "kubelet has sufficient disk space available"
          },
          {
            "type": "MemoryPressure",
            "status": "False",
            "lastHeartbeatTime": "2017-06-10T23:22:49Z",
            "lastTransitionTime": "2017-06-10T23:12:41Z",
            "reason": "KubeletHasSufficientMemory",
            "message": "kubelet has sufficient memory available"
          },
          {
            "type": "DiskPressure",
            "status": "False",
            "lastHeartbeatTime": "2017-06-10T23:22:49Z",
            "lastTransitionTime": "2017-06-10T23:12:41Z",
            "reason": "KubeletHasNoDiskPressure",
            "message": "kubelet has no disk pressure"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastHeartbeatTime": "2017-06-10T23:22:49Z",
            "lastTransitionTime": "2017-06-10T23:12:41Z",
            "reason": "KubeletReady",
            "message": "kubelet is posting ready status"
          }
        ],
        "addresses": [
          {
            "type": "LegacyHostIP",
            "address": "172.17.4.50"
          },
          {
            "type": "InternalIP",
            "address": "172.17.4.50"
          },
          {
            "type": "Hostname",
            "address": "172.17.4.50"
          }
        ],
        "daemonEndpoints": {
          "kubeletEndpoint": {
            "Port": 10250
          }
        },
        "nodeInfo": {
          "machineID": "5c949bb3146241e09f7e671b0704d4fb",
          "systemUUID": "284705D4-2625-4A04-8B45-675B98E6E998",
          "bootID": "48d69df5-2b4b-4136-af94-1dfe72b1fea3",
          "kernelVersion": "4.2.3-300.fc23.x86_64",
          "osImage": "Fedora 23 (Twenty Three)",
          "containerRuntimeVersion": "docker://1.10.3",
          "kubeletVersion": "v1.6.4",
          "kubeProxyVersion": "v1.6.4",
          "operatingSystem": "linux",
          "architecture": "amd64"
        },
        "images": [
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 1646614695
          },
          {
            "names": [
              "tangfeixiong/jdk5u22-centos4:latest"
            ],
            "sizeBytes": 1390171421
          },
          {
            "names": [
              "docker.io/fatherlinux/centos4-base:latest"
            ],
            "sizeBytes": 1211673739
          },
          {
            "names": [
              "picc/dms:base_domain_wls1213-serverjre7-ol7"
            ],
            "sizeBytes": 1134113907
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 1109572121
          },
          {
            "names": [
              "tangfeixiong/weblogic:base_domain_12.1.3-serverjre7-centos7"
            ],
            "sizeBytes": 982083103
          },
          {
            "names": [
              "tangfeixiong/openshift-webtty:latest"
            ],
            "sizeBytes": 902998294
          },
          {
            "names": [
              "tangfeixiong/osospringbootapp:09b1a334"
            ],
            "sizeBytes": 811385452
          },
          {
            "names": [
              "172.17.4.50:30005/tangfx/osospringbootapp:latest",
              "tangfeixiong/osoms-springboot-app:09b1a334"
            ],
            "sizeBytes": 811385451
          },
          {
            "names": [
              "docker.io/tangfeixiong/springboot-sti:gitrev-1125149",
              "hub.qingyuanos.com/admin/springboot-sti:latest",
              "tangfeixiong/springboot-sti:gitcommit-1125149-0901T2236",
              "tangfeixiong/springboot-sti:gitcommit1125149-0901T2236"
            ],
            "sizeBytes": 794495025
          },
          {
            "names": [
              "docker.io/fabric8/jenkins-docker:2.2.265"
            ],
            "sizeBytes": 783029584
          },
          {
            "names": [
              "tangfeixiong/redis-commander:joeferner"
            ],
            "sizeBytes": 745876416
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 714495618
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 678556620
          },
          {
            "names": [
              "172.17.4.50:5000/clearwater_ellis:latest"
            ],
            "sizeBytes": 659340719
          },
          {
            "names": [
              "172.17.4.50:5000/clearwater_homestead:latest"
            ],
            "sizeBytes": 628614294
          },
          {
            "names": [
              "tangfeixiong/tomcat:7-serverjre7-centos7"
            ],
            "sizeBytes": 587042865
          },
          {
            "names": [
              "tangfeixiong/openshift-origin:v1.3.0-alpha.2_88b8a33-dirty"
            ],
            "sizeBytes": 570040773
          },
          {
            "names": [
              "172.17.4.50:5000/clearwater_cassandra:latest"
            ],
            "sizeBytes": 565727040
          },
          {
            "names": [
              "172.17.4.50:5000/clearwater_homer:latest"
            ],
            "sizeBytes": 559201231
          },
          {
            "names": [
              "docker.io/openshift/origin-docker-builder:v1.3.0-alpha.0"
            ],
            "sizeBytes": 543882072
          },
          {
            "names": [
              "docker.io/openshift/origin:v1.3.0-alpha.0"
            ],
            "sizeBytes": 543882072
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 538694868
          },
          {
            "names": [
              "docker.io/openshift/base-centos7:latest"
            ],
            "sizeBytes": 515313323
          },
          {
            "names": [
              "172.17.4.50:5000/clearwater_sprout:latest"
            ],
            "sizeBytes": 500787141
          },
          {
            "names": [
              "docker.io/openshift/origin-f5-router:latest"
            ],
            "sizeBytes": 490622690
          },
          {
            "names": [
              "172.17.4.50:5000/clearwater_bono:latest"
            ],
            "sizeBytes": 487662897
          },
          {
            "names": [
              "docker.io/openshift/origin-docker-builder:v1.3.0-alpha.3"
            ],
            "sizeBytes": 483533748
          },
          {
            "names": [
              "docker.io/openshift/origin:v1.3.0-alpha.3"
            ],
            "sizeBytes": 483533748
          },
          {
            "names": [
              "172.17.4.50:5000/clearwater_chronos:latest"
            ],
            "sizeBytes": 481290812
          },
          {
            "names": [
              "172.17.4.50:5000/clearwater_ralf:latest"
            ],
            "sizeBytes": 476076548
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 472640219
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 455596899
          },
          {
            "names": [
              "172.17.4.50:5000/clearwater_memcached:latest"
            ],
            "sizeBytes": 439189269
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 434431472
          },
          {
            "names": [
              "hub.qingyuanos.com/admin/ubuntu-binary-manila-scheduler:mitaka"
            ],
            "sizeBytes": 431239387
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 427303800
          },
          {
            "names": [
              "docker.io/openshift/origin-docker-builder:v1.3.0-alpha.1"
            ],
            "sizeBytes": 415088397
          },
          {
            "names": [
              "docker.io/openshift/origin:v1.3.0-alpha.1"
            ],
            "sizeBytes": 415088397
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 415030860
          },
          {
            "names": [
              "gcr.io/google_containers/hyperkube-amd64:v1.4.0-ci"
            ],
            "sizeBytes": 408356486
          },
          {
            "names": [
              "gcr.io/google_containers/kibana:1.3"
            ],
            "sizeBytes": 396842879
          },
          {
            "names": [
              "gcr.io/google-samples/cassandra:v11"
            ],
            "sizeBytes": 381549114
          },
          {
            "names": [
              "tangfeixiong/repcached:2.3.1-memcached-1.4.13"
            ],
            "sizeBytes": 378689361
          },
          {
            "names": [
              "docker.io/tangfeixiong/springboot-osev3-examples:web"
            ],
            "sizeBytes": 369714999
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 368001566
          },
          {
            "names": [
              "hub.qingyuanos.com/admin/go2paas:1610140535.gitref-57279bd",
              "tangfeixiong/gotopaas:1610140535.gitref-57279bd"
            ],
            "sizeBytes": 359090031
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 357262959
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 349978444
          },
          {
            "names": [
              "\u003cnone\u003e@\u003cnone\u003e",
              "\u003cnone\u003e:\u003cnone\u003e"
            ],
            "sizeBytes": 349095007
          }
        ]
      }
    }
  ]
}
[vagrant@localhost ~]$ curl http://127.0.0.1:8080/api/v1/namespaces
{
  "kind": "NamespaceList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/namespaces",
    "resourceVersion": "895"
  },
  "items": [
    {
      "metadata": {
        "name": "default",
        "selfLink": "/api/v1/namespacesdefault",
        "uid": "4a659add-4e32-11e7-8689-0800274654e7",
        "resourceVersion": "661",
        "creationTimestamp": "2017-06-10T23:12:36Z"
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
        "name": "kube-public",
        "selfLink": "/api/v1/namespaceskube-public",
        "uid": "4a6fb4ee-4e32-11e7-8689-0800274654e7",
        "resourceVersion": "666",
        "creationTimestamp": "2017-06-10T23:12:36Z"
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
        "selfLink": "/api/v1/namespaceskube-system",
        "uid": "4a6046e8-4e32-11e7-8689-0800274654e7",
        "resourceVersion": "658",
        "creationTimestamp": "2017-06-10T23:12:36Z"
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
[vagrant@localhost ~]$ curl http://127.0.0.1:8080/healthz
ok
```

### Client kubectl

Version
```
[vagrant@localhost ~]$ kubectl version --client
Client Version: version.Info{Major:"1", Minor:"6", GitVersion:"v1.6.4", GitCommit:"d6f433224538d4f9ca2f7ae19b252e6fcb66a3ae", GitTreeState:"clean", BuildDate:"2017-05-19T18:44:27Z", GoVersion:"go1.7.5", Compiler:"gc", Platform:"linux/amd64"}
```

Configure
```
vagrant@localhost:~$ kubectl config set-cluster kube --server=https://172.17.4.200:6443
vagrant@localhost:~$ encoded=$(sudo base64 -w0 /srv/kubernetes/ca.crt); kubectl config set clusters.kube.certificate-authority-data $encoded
Property "clusters.kube.certificate-authority-data" set.
vagrant@localhost:~$ encoded=$(sudo base64 -w0 /srv/kubernetes/kubecfg.crt); kubectl config set users.admin.client-certificate-data $encoded
Property "users.admin.client-certificate-data" set.
vagrant@localhost:~$ encoded=$(sudo base64 -w0 /srv/kubernetes/kubecfg.key); kubectl config set users.admin.client-key-data $encoded
Property "users.admin.client-key-data" set.
vagrant@localhost:~$ kubectl config set-context kube-admin --cluster=kube --user=admin
Context "kube-admin" set.
vagrant@localhost:~$ kubectl config use-context kube-admin
Switched to context "kube-admin".
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl config view
```

File
```
vagrant@vagrant-ubuntu-trusty-64:~$ cat .kube/config
```

Command
```
[vagrant@localhost ~]$ kubectl get nodes
NAME          STATUS    AGE       VERSION
172.17.4.50   Ready     3m        v1.6.4
[vagrant@localhost ~]$ kubectl get ns
NAME          STATUS    AGE
default       Active    3m
kube-public   Active    3m
kube-system   Active    3m
[vagrant@localhost ~]$ kubectl get --namespace=kube-public all
No resources found.
[vagrant@localhost ~]$ kubectl get --namespace=kube-system all
NAME                            READY     STATUS    RESTARTS   AGE
po/etcd-server-172.17.4.50      1/1       Running   0          3m
po/kube-apiserver-172.17.4.50   1/1       Running   4          3m
[vagrant@localhost ~]$ kubectl get all
NAME             CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
svc/kubernetes   10.123.240.1   <none>        443/TCP   4m
[vagrant@localhost ~]$ kubectl version
Client Version: version.Info{Major:"1", Minor:"6", GitVersion:"v1.6.4", GitCommit:"d6f433224538d4f9ca2f7ae19b252e6fcb66a3ae", GitTreeState:"clean", BuildDate:"2017-05-19T18:44:27Z", GoVersion:"go1.7.5", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"6", GitVersion:"v1.6.4", GitCommit:"d6f433224538d4f9ca2f7ae19b252e6fcb66a3ae", GitTreeState:"clean", BuildDate:"2017-05-19T18:33:17Z", GoVersion:"go1.7.5", Compiler:"gc", Platform:"linux/amd64"}
[vagrant@localhost ~]$ kubectl api-versions
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
policy/v1beta1
rbac.authorization.k8s.io/v1alpha1
rbac.authorization.k8s.io/v1beta1
settings.k8s.io/v1alpha1
storage.k8s.io/v1
storage.k8s.io/v1beta1
v1
```

Copy into Mac
```
fanhonglingdeMacBook-Pro:Downloads fanhongling$ KUBECONFIG=./kubeconfig kubectl version
Client Version: version.Info{Major:"1", Minor:"6", GitVersion:"v1.6.4", GitCommit:"d6f433224538d4f9ca2f7ae19b252e6fcb66a3ae", GitTreeState:"clean", BuildDate:"2017-05-19T18:44:27Z", GoVersion:"go1.7.5", Compiler:"gc", Platform:"darwin/amd64"}
Server Version: version.Info{Major:"1", Minor:"6", GitVersion:"v1.6.4", GitCommit:"d6f433224538d4f9ca2f7ae19b252e6fcb66a3ae", GitTreeState:"clean", BuildDate:"2017-05-19T18:33:17Z", GoVersion:"go1.7.5", Compiler:"gc", Platform:"linux/amd64"}
fanhonglingdeMacBook-Pro:Downloads fanhongling$ KUBECONFIG=./kubeconfig kubectl get nodes
NAME          STATUS    AGE       VERSION
172.17.4.50   Ready     53m       v1.6.4
fanhonglingdeMacBook-Pro:Downloads fanhongling$ KUBECONFIG=./kubeconfig kubectl get ns
NAME          STATUS    AGE
default       Active    53m
kube-public   Active    53m
kube-system   Active    53m
```

### Kube-controller-manager

Log location
```
[vagrant@localhost ~]$ sudo touch /var/log/kube-controller-manager.log
```

Manifest
```
[vagrant@localhost ~]$ sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%\("image": \)".*kube_docker_registry\x27\x5d\x7d\x7d\(/kube-controller-manager:\).*kube-controller-manager_docker_tag.*"%\1"gcr.io/google_containers\2$KUBE_CONTROLLERMANAGER_DOCKERTAG"%;s%{{params}}%--master=127.0.0.1:8080 --cluster-name=kubernetes --cluster-cidr=10.0.0.0/14 --service-cluster-ip-range=10.3.0.0/20 --terminated-pod-gc-threshold=12500 --enable-garbage-collector=true --service-account-private-key-file=/srv/kubernetes/server.key --v=2 --root-ca-file=/srv/kubernetes/ca.crt --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true --enable-hostpath-provisioner=true%;s/{{container_env}}//;s/{{cloud_config_mount}}\|{{additional_cloud_config_mount}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;s%{{cloud_config_volume}}\|{{additional_cloud_config_volume}}%%' /opt/kubernetes/cluster/saltbase/salt/kube-controller-manager/kube-controller-manager.manifest | KUBE_CONTROLLERMANAGER_DOCKERTAG=$(cat /opt/kubernetes/server/bin/kube-controller-manager.docker_tag) envsubst | sudo tee /etc/kubernetes/manifests/kube-controller-manager.json


{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"kube-controller-manager",
  "namespace": "kube-system",
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
    "image": "gcr.io/google_containers/kube-controller-manager:6b1e075eaab8773444890c5cc17e0dcf",
    "resources": {
      "requests": {
        "cpu": "200m"
      }
    },
    "command": [
                 "/bin/sh",
                 "-c",
                 "/usr/local/bin/kube-controller-manager --master=127.0.0.1:8080 --cluster-name=kubernetes --cluster-cidr=10.0.0.0/14 --service-cluster-ip-range=10.3.0.0/20 --terminated-pod-gc-threshold=12500 --enable-garbage-collector=true --service-account-private-key-file=/srv/kubernetes/server.key --v=2 --root-ca-file=/srv/kubernetes/ca.crt --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true --enable-hostpath-provisioner=true 1>>/var/log/kube-controller-manager.log 2>&1"
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

Check
```
[vagrant@localhost ~]$ kubectl --namespace=kube-system get pods
NAME                                  READY     STATUS    RESTARTS   AGE
etcd-server-172.17.4.50               1/1       Running   0          2h
kube-apiserver-172.17.4.50            1/1       Running   4          2h
kube-controller-manager-172.17.4.50   1/1       Running   5          1h
[vagrant@localhost ~]$ docker ps
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
3b87849d4c06        3c2198b97c22                               "/bin/sh -c '/usr/loc"   50 seconds ago      Up 49 seconds                           k8s_kube-controller-manager_kube-controller-manager-172.17.4.50_kube-system_a3bf90a7b790ce310f59114ed2c72d1c_5
c831bc9a3027        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 3 minutes ago       Up 3 minutes                            k8s_POD_kube-controller-manager-172.17.4.50_kube-system_a3bf90a7b790ce310f59114ed2c72d1c_0
4871d8978295        9196766fca5a                               "/bin/sh -c '/usr/loc"   About an hour ago   Up About an hour                        k8s_kube-apiserver_kube-apiserver-172.17.4.50_kube-system_75a081d878fbde18ca953b4426639b8d_4
575c3883f38a        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 About an hour ago   Up About an hour                        k8s_POD_kube-apiserver-172.17.4.50_kube-system_75a081d878fbde18ca953b4426639b8d_0
67f96a93ac02        e6a6589433f4                               "/bin/sh -c 'if [ -e "   5 hours ago         Up 5 hours                              k8s_etcd-container_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
d67c2281716d        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 5 hours ago         Up 5 hours                              k8s_POD_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
```

Check log
```
[vagrant@localhost ~]$ ls -l /var/log/kube-controller-manager.log 
-rw-r--r--. 1 root root 9083 6月  11 00:39 /var/log/kube-controller-manager.log
[vagrant@localhost ~]$ tail /var/log/kube-controller-manager.log 
I0611 00:39:03.239997       7 disruption.go:277] Sending events to api server.
I0611 00:39:03.248303       7 garbagecollector.go:116] Garbage Collector: All resource monitors have synced. Proceeding to collect garbage
I0611 00:39:03.251144       7 ttlcontroller.go:271] Changed ttl annotation for node 172.17.4.50 to 0 seconds
I0611 00:39:03.265809       7 nodecontroller.go:595] NodeController observed a new Node: "172.17.4.50"
I0611 00:39:03.265946       7 controller_utils.go:274] Recording Registered Node 172.17.4.50 in NodeController event message for node 172.17.4.50
I0611 00:39:03.265974       7 nodecontroller.go:612] Initializing eviction metric for zone: 
I0611 00:39:03.266951       7 taint_controller.go:180] Starting NoExecuteTaintManager
W0611 00:39:03.268759       7 nodecontroller.go:947] Missing timestamp for Node 172.17.4.50. Assuming now as a timestamp.
I0611 00:39:03.268955       7 nodecontroller.go:863] NodeController detected that zone  is now in state Normal.
I0611 00:39:03.269133       7 event.go:217] Event(v1.ObjectReference{Kind:"Node", Namespace:"", Name:"172.17.4.50", UID:"4dd88960-4e32-11e7-8689-0800274654e7", APIVersion:"", ResourceVersion:"", FieldPath:""}): type: 'Normal' reason: 'RegisteredNode' Node 172.17.4.50 event: Registered Node 172.17.4.50 in NodeController
```

### Kube-scheduler

Log location
```
[vagrant@localhost ~]$ sudo touch /var/log/kube-scheduler.log
```

Manifest
```
[vagrant@localhost ~]$ sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%\("image": \)".*kube_docker_registry\x27\x5d\x7d\x7d\(/kube-scheduler:\).*kube-scheduler_docker_tag.*"%\1"gcr.io/google_containers\2$KUBE_SCHEDULER_DOCKERTAG"%;s%{{params}}%--master=127.0.0.1:8080 --v=2 --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true --algorithm-provider=DefaultProvider%;s%{{srv_kube_path}}%/srv/kubernetes%g' /opt/kubernetes/cluster/saltbase/salt/kube-scheduler/kube-scheduler.manifest | KUBE_SCHEDULER_DOCKERTAG=$(cat /opt/kubernetes/server/bin/kube-scheduler.docker_tag) envsubst | sudo tee /etc/kubernetes/manifests/kube-scheduler.json

{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"kube-scheduler",
  "namespace": "kube-system",
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
    "image": "gcr.io/google_containers/kube-scheduler:c8ed221003bb194f37ef4221727bce1c",
    "resources": {
      "requests": {
        "cpu": "100m"
      }
    },
    "command": [
                 "/bin/sh",
                 "-c",
                 "/usr/local/bin/kube-scheduler --master=127.0.0.1:8080 --v=2 --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true --algorithm-provider=DefaultProvider 1>>/var/log/kube-scheduler.log 2>&1"
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

Check
```
[vagrant@localhost ~]$ kubectl --namespace=kube-system get pods
NAME                                  READY     STATUS    RESTARTS   AGE
etcd-server-172.17.4.50               1/1       Running   0          2h
kube-apiserver-172.17.4.50            1/1       Running   4          2h
kube-controller-manager-172.17.4.50   1/1       Running   5          1h
kube-scheduler-172.17.4.50            1/1       Running   0          30m
[vagrant@localhost ~]$ docker ps
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
0c317682a020        36723c5b4c70                               "/bin/sh -c '/usr/loc"   47 seconds ago      Up 46 seconds                           k8s_kube-scheduler_kube-scheduler-172.17.4.50_kube-system_a54d0d44e53453b1b04e472bf676e5ef_0
7349a4e19c38        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 48 seconds ago      Up 47 seconds                           k8s_POD_kube-scheduler-172.17.4.50_kube-system_a54d0d44e53453b1b04e472bf676e5ef_0
3b87849d4c06        3c2198b97c22                               "/bin/sh -c '/usr/loc"   29 minutes ago      Up 29 minutes                           k8s_kube-controller-manager_kube-controller-manager-172.17.4.50_kube-system_a3bf90a7b790ce310f59114ed2c72d1c_5
c831bc9a3027        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 32 minutes ago      Up 32 minutes                           k8s_POD_kube-controller-manager-172.17.4.50_kube-system_a3bf90a7b790ce310f59114ed2c72d1c_0
4871d8978295        9196766fca5a                               "/bin/sh -c '/usr/loc"   About an hour ago   Up About an hour                        k8s_kube-apiserver_kube-apiserver-172.17.4.50_kube-system_75a081d878fbde18ca953b4426639b8d_4
575c3883f38a        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 About an hour ago   Up About an hour                        k8s_POD_kube-apiserver-172.17.4.50_kube-system_75a081d878fbde18ca953b4426639b8d_0
67f96a93ac02        e6a6589433f4                               "/bin/sh -c 'if [ -e "   5 hours ago         Up 5 hours                              k8s_etcd-container_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
d67c2281716d        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 5 hours ago         Up 5 hours                              k8s_POD_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
```

Check log
```
[vagrant@localhost ~]$ ls -l /var/log/kube-scheduler.log 
-rw-r--r--. 1 root root 1464 6月  11 01:07 /var/log/kube-scheduler.log
[vagrant@localhost ~]$ tail /var/log/kube-scheduler.log 
I0611 01:07:29.331962       7 feature_gate.go:144] feature gates: map[AffinityInAnnotations:true AllAlpha:true TaintBasedEvictions:true DynamicKubeletConfig:true DynamicVolumeProvisioning:true ExperimentalCriticalPodAnnotation:true ExperimentalHostUserNamespaceDefaulting:true Accelerators:true]
I0611 01:07:29.332707       7 factory.go:300] Creating scheduler from algorithm provider 'DefaultProvider'
I0611 01:07:29.332726       7 factory.go:346] Creating scheduler with fit predicates 'map[MaxEBSVolumeCount:{} MaxGCEPDVolumeCount:{} MaxAzureDiskVolumeCount:{} NoDiskConflict:{} GeneralPredicates:{} NoVolumeZoneConflict:{} MatchInterPodAffinity:{} PodToleratesNodeTaints:{} CheckNodeMemoryPressure:{} CheckNodeDiskPressure:{}]' and priority functions 'map[TaintTolerationPriority:{} SelectorSpreadPriority:{} InterPodAffinityPriority:{} LeastRequestedPriority:{} BalancedResourceAllocation:{} NodePreferAvoidPodsPriority:{} NodeAffinityPriority:{}]
I0611 01:07:29.332875       7 leaderelection.go:179] attempting to acquire leader lease...
I0611 01:07:29.372594       7 leaderelection.go:189] successfully acquired lease kube-system/kube-scheduler
I0611 01:07:29.372874       7 event.go:217] Event(v1.ObjectReference{Kind:"Endpoints", Namespace:"kube-system", Name:"kube-scheduler", UID:"5711f520-4e42-11e7-928f-0800274654e7", APIVersion:"v1", ResourceVersion:"2517", FieldPath:""}): type: 'Normal' reason: 'LeaderElection' localhost.localdomain became leader
```

### Kube-proxy

Log location
```
[vagrant@localhost ~]$ sudo touch /var/log/kube-proxy.log
```

Working directory and _kubeconfig_
```
[vagrant@localhost ~]$ sudo mkdir -p /var/lib/kube-proxy/
[vagrant@localhost ~]$ sudo cp /data/src/github.com/openshift/origin/etc/kubeconfig /var/lib/kube-proxy/
```

Manifest
```
[vagrant@localhost ~]$ sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%\(image: \).*kube_docker_registry\x27\x5d\x7d\x7d\(/kube-proxy:\).*kube-proxy_docker_tag.*%\1gcr.io/google_containers\2$KUBE_PROXY_DOCKERTAG%;s/{{ cpurequest }}/100m/;s%{{api_servers_with_port}}%--master=https://$KUBE_APISERVER_HOST:6443%;s%{{kubeconfig}}%--kubeconfig=/var/lib/kube-proxy/kubeconfig%;s%{{cluster_cidr}}%--cluster-cidr=10.0.0.0/14%;s%{{params}}%--v=2 --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true%;s/{{container_env}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;' /opt/kubernetes/cluster/saltbase/salt/kube-proxy/kube-proxy.manifest | KUBE_PROXY_DOCKERTAG=$(cat /opt/kubernetes/server/bin/kube-proxy.docker_tag) KUBE_APISERVER_HOST=172.17.4.50 envsubst | sudo tee /etc/kubernetes/manifests/kube-proxy.yaml

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
  containers:
  - name: kube-proxy
    image: gcr.io/google_containers/kube-proxy:4c6ae16e62496a67d3e9cc622c1006a3
    resources:
      requests:
        cpu: 100m
    command:
    - /bin/sh
    - -c
    - echo -998 > /proc/$$$/oom_score_adj && kube-proxy --master=https://172.17.4.50:6443 --kubeconfig=/var/lib/kube-proxy/kubeconfig --cluster-cidr=10.0.0.0/14 --resource-container="" --v=2 --feature-gates=Accelerators=true,AffinityInAnnotations=true,AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,TaintBasedEvictions=true 1>>/var/log/kube-proxy.log 2>&1
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
```

Check
```
[vagrant@localhost ~]$ kubectl --namespace=kube-system get pods
NAME                                  READY     STATUS    RESTARTS   AGE
etcd-server-172.17.4.50               1/1       Running   0          2h
kube-apiserver-172.17.4.50            1/1       Running   4          2h
kube-controller-manager-172.17.4.50   1/1       Running   5          1h
kube-proxy-172.17.4.50                1/1       Running   6          7m
kube-scheduler-172.17.4.50            1/1       Running   0          30m
[vagrant@localhost ~]$ docker ps
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
ad0f164d068b        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 31 seconds ago      Up 30 seconds                           k8s_POD_kube-proxy-172.17.4.50_kube-system_7bc8b0e23c37832b93252f9bef47fe6d_0
0c317682a020        36723c5b4c70                               "/bin/sh -c '/usr/loc"   23 minutes ago      Up 23 minutes                           k8s_kube-scheduler_kube-scheduler-172.17.4.50_kube-system_a54d0d44e53453b1b04e472bf676e5ef_0
7349a4e19c38        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 23 minutes ago      Up 23 minutes                           k8s_POD_kube-scheduler-172.17.4.50_kube-system_a54d0d44e53453b1b04e472bf676e5ef_0
3b87849d4c06        3c2198b97c22                               "/bin/sh -c '/usr/loc"   51 minutes ago      Up 51 minutes                           k8s_kube-controller-manager_kube-controller-manager-172.17.4.50_kube-system_a3bf90a7b790ce310f59114ed2c72d1c_5
c831bc9a3027        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 54 minutes ago      Up 54 minutes                           k8s_POD_kube-controller-manager-172.17.4.50_kube-system_a3bf90a7b790ce310f59114ed2c72d1c_0
4871d8978295        9196766fca5a                               "/bin/sh -c '/usr/loc"   2 hours ago         Up 2 hours                              k8s_kube-apiserver_kube-apiserver-172.17.4.50_kube-system_75a081d878fbde18ca953b4426639b8d_4
575c3883f38a        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 2 hours ago         Up 2 hours                              k8s_POD_kube-apiserver-172.17.4.50_kube-system_75a081d878fbde18ca953b4426639b8d_0
67f96a93ac02        e6a6589433f4                               "/bin/sh -c 'if [ -e "   6 hours ago         Up 6 hours                              k8s_etcd-container_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
d67c2281716d        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 6 hours ago         Up 6 hours                              k8s_POD_etcd-server-172.17.4.50_kube-system_92ea623866c62d0d58056ae7081c0fbe_0
```

Check log
```
[vagrant@localhost ~]$ ls -l /var/log/kube-proxy.log 
-rw-r--r--. 1 root root 6461 6月  11 01:36 /var/log/kube-proxy.log
[vagrant@localhost ~]$ tail /var/log/kube-proxy.log 
Flag --resource-container has been deprecated, This feature will be removed in a later release.
I0611 01:53:20.369694       7 feature_gate.go:144] feature gates: map[ExperimentalHostUserNamespaceDefaulting:true Accelerators:true AffinityInAnnotations:true AllAlpha:true DynamicKubeletConfig:true ExperimentalCriticalPodAnnotation:true TaintBasedEvictions:true DynamicVolumeProvisioning:true]
I0611 01:53:20.401007       7 iptables.go:175] Could not connect to D-Bus system bus: dial unix /var/run/dbus/system_bus_socket: connect: no such file or directory
I0611 01:53:20.401053       7 server.go:173] setting OOM scores is unsupported in this build
I0611 01:53:20.411873       7 server.go:225] Using iptables Proxier.
W0611 01:53:20.463089       7 server.go:469] Failed to retrieve node info: nodes "localhost.localdomain" not found
W0611 01:53:20.463620       7 proxier.go:293] invalid nodeIP, initializing kube-proxy with 127.0.0.1 as nodeIP
I0611 01:53:20.463661       7 server.go:249] Tearing down userspace rules.
I0611 01:53:20.581651       7 conntrack.go:81] Set sysctl 'net/netfilter/nf_conntrack_max' to 131072
I0611 01:53:20.581862       7 conntrack.go:66] Setting conntrack hashsize to 32768
I0611 01:53:20.582131       7 conntrack.go:81] Set sysctl 'net/netfilter/nf_conntrack_tcp_timeout_established' to 86400
I0611 01:53:20.582165       7 conntrack.go:81] Set sysctl 'net/netfilter/nf_conntrack_tcp_timeout_close_wait' to 3600
I0611 01:53:20.582256       7 proxier.go:472] Adding new service "default/kubernetes:https" at 10.123.240.1:443/TCP
I0611 01:53:20.583453       7 proxier.go:741] Not syncing iptables until Services and Endpoints have been received from master
I0611 01:53:20.583482       7 proxier.go:540] Received first Endpoints update
```

## Addons

Directory
```
[vagrant@localhost ~]$ sudo mkdir -p /etc/kubernetes/addons
```

### Kube-addon-manager

Image
```
[vagrant@localhost ~]$ docker pull gcr.io/google-containers/kube-addon-manager:v6.4-beta.1
Trying to pull repository gcr.io/google-containers/kube-addon-manager ... 
v6.4-beta.1: Pulling from gcr.io/google-containers/kube-addon-manager
627beaf3eaaf: Pull complete 
3f1bedcfa9cc: Pull complete 
703edae953b1: Pull complete 
509b2e89b0ac: Pull complete 
d8236f1e8f64: Pull complete 
921d51ba264c: Pull complete 
db9bef7be833: Pull complete 
Digest: sha256:03a3aeeead17886d1a8c17c41ae56f7a1be4b7c23a281579fbfb13acb2c1eda1
Status: Downloaded newer image for gcr.io/google-containers/kube-addon-manager:v6.4-beta.1
```

Manifest
```
[vagrant@localhost ~]$ sudo cp /opt/kubernetes/cluster/saltbase/salt/kube-addons/kube-addon-manager.yaml /etc/kubernetes/manifests/
```

Check
```
[vagrant@localhost ~]$ kubectl --namespace=kube-system get all --selector=k8s-app=kube-dns
NAME                          READY     STATUS    RESTARTS   AGE
po/kube-dns-806549836-3n26k   3/3       Running   0          27m

NAME           CLUSTER-IP   EXTERNAL-IP   PORT(S)         AGE
svc/kube-dns   10.3.0.10    <none>        53/UDP,53/TCP   28m

NAME              DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/kube-dns   1         1         1            1           29m

NAME                    DESIRED   CURRENT   READY     AGE
rs/kube-dns-806549836   1         1         1         27m
```

### DNS 

Image
```
[vagrant@localhost ~]$ docker pull gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.1
Trying to pull repository gcr.io/google_containers/k8s-dns-kube-dns-amd64 ... 
1.14.1: Pulling from gcr.io/google_containers/k8s-dns-kube-dns-amd64
0a8490d0dfd3: Pull complete 
e9448fc5df71: Pull complete 
Digest: sha256:33914315e600dfb756e550828307dfa2b21fb6db24fe3fe495e33d1022f9245d
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.1
[vagrant@localhost ~]$ docker pull gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.1
Trying to pull repository gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64 ... 
1.14.1: Pulling from gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64
b7f33cc0b48e: Pull complete 
ab6ef7f2f90b: Pull complete 
3e0100af9919: Pull complete 
b582d55d1318: Pull complete 
c962ee2db05b: Pull complete 
Digest: sha256:89c9a1d3cfbf370a9c1a949f39f92c1dc2dbe8c3e6cc1802b7f2b48e4dfe9a9e
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.1
[vagrant@localhost ~]$ docker pull gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.1
Trying to pull repository gcr.io/google_containers/k8s-dns-sidecar-amd64 ... 
1.14.1: Pulling from gcr.io/google_containers/k8s-dns-sidecar-amd64
0a8490d0dfd3: Already exists 
13c5596281ba: Pull complete 
Digest: sha256:d33a91a5d65c223f410891001cd379ac734d036429e033865d700a4176e944b0
Status: Downloaded newer image for gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.1
```

Manifests
```
[vagrant@localhost ~]$ sed 's/{{ pillar.*dns_domain.*}}/cluster.local/g;s/{{ pillar.*federations_domain_map.*}}//' /opt/kubernetes/cluster/saltbase/salt/kube-addons/dns/kubedns-controller.yaml.in | sudo tee /etc/kubernetes/addons/kubedns-controller.yaml
# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Should keep target in cluster/addons/dns-horizontal-autoscaler/dns-horizontal-autoscaler.yaml
# in sync with this file.

# Warning: This is a file generated from the base underscore template file: kubedns-controller.yaml.base

apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  # replicas: not specified here:
  # 1. In order to make Addon Manager do not reconcile this replicas parameter.
  # 2. Default is 1.
  # 3. Will be tuned in real time if DNS horizontal auto-scaling is turned on.
  strategy:
    rollingUpdate:
      maxSurge: 10%
      maxUnavailable: 0
  selector:
    matchLabels:
      k8s-app: kube-dns
  template:
    metadata:
      labels:
        k8s-app: kube-dns
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      tolerations:
      - key: "CriticalAddonsOnly"
        operator: "Exists"
      volumes:
      - name: kube-dns-config
        configMap:
          name: kube-dns
          optional: true
      containers:
      - name: kubedns
        image: gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.1
        resources:
          # TODO: Set memory limits when we've profiled the container for large
          # clusters, then set request = limit to keep this container in
          # guaranteed class. Currently, this container falls into the
          # "burstable" category so the kubelet doesn't backoff from restarting it.
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        livenessProbe:
          httpGet:
            path: /healthcheck/kubedns
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        readinessProbe:
          httpGet:
            path: /readiness
            port: 8081
            scheme: HTTP
          # we poll on pod startup for the Kubernetes master service and
          # only setup the /readiness HTTP server once that's available.
          initialDelaySeconds: 3
          timeoutSeconds: 5
        args:
        - --domain=cluster.local.
        - --dns-port=10053
        - --config-dir=/kube-dns-config
        - --v=2
        
        env:
        - name: PROMETHEUS_PORT
          value: "10055"
        ports:
        - containerPort: 10053
          name: dns-local
          protocol: UDP
        - containerPort: 10053
          name: dns-tcp-local
          protocol: TCP
        - containerPort: 10055
          name: metrics
          protocol: TCP
        volumeMounts:
        - name: kube-dns-config
          mountPath: /kube-dns-config
      - name: dnsmasq
        image: gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.1
        livenessProbe:
          httpGet:
            path: /healthcheck/dnsmasq
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - -v=2
        - -logtostderr
        - -configDir=/etc/k8s/dns/dnsmasq-nanny
        - -restartDnsmasq=true
        - --
        - -k
        - --cache-size=1000
        - --log-facility=-
        - --server=/cluster.local/127.0.0.1#10053
        - --server=/in-addr.arpa/127.0.0.1#10053
        - --server=/ip6.arpa/127.0.0.1#10053
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        # see: https://github.com/kubernetes/kubernetes/issues/29055 for details
        resources:
          requests:
            cpu: 150m
            memory: 20Mi
        volumeMounts:
        - name: kube-dns-config
          mountPath: /etc/k8s/dns/dnsmasq-nanny
      - name: sidecar
        image: gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.1
        livenessProbe:
          httpGet:
            path: /metrics
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - --v=2
        - --logtostderr
        - --probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.cluster.local,5,A
        - --probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.cluster.local,5,A
        ports:
        - containerPort: 10054
          name: metrics
          protocol: TCP
        resources:
          requests:
            memory: 20Mi
            cpu: 10m
      dnsPolicy: Default  # Don't use cluster DNS.
      serviceAccountName: kube-dns
[vagrant@localhost ~]$ sed 's/{{ pillar.*dns_server.*}}/10.3.0.10/' /opt/kubernetes/cluster/saltbase/salt/kube-addons/dns/kubedns-svc.yaml.in | sudo tee /etc/kubernetes/addons/kubedns-svc.yaml
# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Warning: This is a file generated from the base underscore template file: kubedns-svc.yaml.base

apiVersion: v1
kind: Service
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/name: "KubeDNS"
spec:
  selector:
    k8s-app: kube-dns
  clusterIP: 10.3.0.10
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP
```

Check
```
[vagrant@localhost ~]$ kubectl --namespace=kube-system get all -l k8s-app=kube-dns
NAME                          READY     STATUS    RESTARTS   AGE
po/kube-dns-806549836-21l3d   2/3       Running   9          9m

NAME           CLUSTER-IP      EXTERNAL-IP   PORT(S)         AGE
svc/kube-dns   10.3.0.10       <none>        53/UDP,53/TCP   5m

NAME              DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/kube-dns   1         1         1            0           9m

NAME                    DESIRED   CURRENT   READY     AGE
rs/kube-dns-806549836   1         1         0         9m
```

### Dashboard

Image
```
[vagrant@localhost ~]$ docker pull gcr.io/google_containers/kubernetes-dashboard-amd64:v1.6.0
Trying to pull repository gcr.io/google_containers/kubernetes-dashboard-amd64 ... 
v1.6.0: Pulling from gcr.io/google_containers/kubernetes-dashboard-amd64
156c1d5571df: Pull complete 
Digest: sha256:4ad64dfa7159ff4a99a65a4f96432f2fdb6542857cf230858b3159017833a882
Status: Downloaded newer image for gcr.io/google_containers/kubernetes-dashboard-amd64:v1.6.0
```

Manifest
```
[vagrant@localhost ~]$ sudo cp /opt/kubernetes/cluster/saltbase/salt/kube-addons/dashboard/dashboard-controller.yaml /etc/kubernetes/addons/
[vagrant@localhost ~]$ sudo cp /opt/kubernetes/cluster/saltbase/salt/kube-addons/dashboard/dashboard-service.yaml /etc/kubernetes/addons/
```

Alternative
```
[vagrant@localhost ~]$ sed 's/^.*resources:$/        args:\n        - --apiserver-host=http:\/\/127.0.0.1:8080\n&/;s/^    spec:$/&\n      hostNetwork: true/' /opt/kubernetes/cluster/saltbase/salt/kube-addons/dashboard/dashboard-controller.yaml | sudo tee /etc/kubernetes/addons/dashboard-controller.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: kubernetes-dashboard
  template:
    metadata:
      labels:
        k8s-app: kubernetes-dashboard
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      hostNetwork: true
      containers:
      - name: kubernetes-dashboard
        image: gcr.io/google_containers/kubernetes-dashboard-amd64:v1.6.0
        args:
        - --apiserver-host=http://127.0.0.1:8080
        resources:
          # keep request = limit to keep this container in guaranteed class
          limits:
            cpu: 100m
            memory: 50Mi
          requests:
            cpu: 100m
            memory: 50Mi
        ports:
        - containerPort: 9090
        livenessProbe:
          httpGet:
            path: /
            port: 9090
          initialDelaySeconds: 30
          timeoutSeconds: 30
      tolerations:
      - key: "CriticalAddonsOnly"
        operator: "Exists"
```

Check
```
[vagrant@localhost ~]$ kubectl --namespace=kube-system get all -l k8s-app=kubernetes-dashboard
NAME                                       READY     STATUS    RESTARTS   AGE
po/kubernetes-dashboard-1995634392-362sf   1/1       Running   0          42s

NAME                       CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
svc/kubernetes-dashboard   10.3.8.135   <none>        80/TCP    34m

NAME                          DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/kubernetes-dashboard   1         1         1            1           2m

NAME                                 DESIRED   CURRENT   READY     AGE
rs/kubernetes-dashboard-1995634392   1         1         1         42s
```

Curl
```
[vagrant@localhost ~]$ curl 127.0.0.1:9090
 <!doctype html> <html ng-app="kubernetesDashboard"> <head> <meta charset="utf-8"> <title ng-controller="kdTitle as $ctrl" ng-bind="$ctrl.title()"></title> <link rel="icon" type="image/png" href="assets/images/kubernetes-logo.png"> <meta name="viewport" content="width=device-width"> <link rel="stylesheet" href="static/vendor.4f4b705f.css"> <link rel="stylesheet" href="static/app.93b90a74.css"> </head> <body> <!--[if lt IE 10]>
      <p class="browsehappy">You are using an <strong>outdated</strong> browser.
      Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your
      experience.</p>
    <![endif]--> <kd-chrome layout="column" layout-fill> </kd-chrome> <script src="static/vendor.6952e31e.js"></script> <script src="api/appConfig.json"></script> <script src="static/app.8a6b8127.js"></script> </body> </html> 
```

Browse
```
[vagrant@localhost ~]$ kubectl expose deployment/kubernetes-dashboard --namespace=kube-system --name=kubernetes-dashboard-ex --target-port=9090 --type=NodePort
service "kubernetes-dashboard-ex" exposed
[vagrant@localhost ~]$ kubectl --namespace=kube-system get service,ep kubernetes-dashboard-ex 
NAME                          CLUSTER-IP   EXTERNAL-IP   PORT(S)          AGE
svc/kubernetes-dashboard-ex   10.3.7.209   <nodes>       9090:32186/TCP   1m

NAME                         ENDPOINTS          AGE
ep/kubernetes-dashboard-ex   172.17.4.50:9090   1m
```

![屏幕快照 2017-06-11 上午4.25.28.png](./屏幕快照%202017-06-11%20上午4.25.28.png)

## Cluster Networking

### Flannel

