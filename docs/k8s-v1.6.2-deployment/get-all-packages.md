Deploy K8s cluster under Kubelet － 在CentOS/Fedora部署Kubernetes v1.6.2
================================ 

__Tips__

另可参考[K8s v1.5.7安装](../k8s-v1.5.7-deployment)

__安装Docker Engine 1.12.6__

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

查看Docker版本和信息
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

### 下载K8s社区分发包

Download
```
[tangfx@localhost ~]$ curl -jkSLOC - https://dl.k8s.io/v1.6.2/kubernetes-server-linux-amd64.tar.gz
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

Extract
```
[tangfx@localhost ~]$ sudo tar -C /opt -zxf kubernetes-server-linux-amd64.tar.gz -v
kubernetes/
kubernetes/kubernetes-src.tar.gz
kubernetes/server/
kubernetes/server/bin/
kubernetes/server/bin/kubefed
kubernetes/server/bin/kube-apiserver.docker_tag
kubernetes/server/bin/kubelet
kubernetes/server/bin/kube-proxy.tar
kubernetes/server/bin/kube-apiserver.tar
kubernetes/server/bin/cloud-controller-manager
kubernetes/server/bin/kube-controller-manager.tar
kubernetes/server/bin/kube-aggregator
kubernetes/server/bin/kube-scheduler
kubernetes/server/bin/kubectl
kubernetes/server/bin/kube-proxy.docker_tag
kubernetes/server/bin/kube-scheduler.docker_tag
kubernetes/server/bin/kubeadm
kubernetes/server/bin/kube-controller-manager.docker_tag
kubernetes/server/bin/hyperkube
kubernetes/server/bin/kube-proxy
kubernetes/server/bin/kube-controller-manager
kubernetes/server/bin/kube-scheduler.tar
kubernetes/server/bin/kube-apiserver
kubernetes/server/bin/kube-aggregator.tar
kubernetes/server/bin/kube-aggregator.docker_tag
kubernetes/LICENSES
kubernetes/addons/
```

K8s images
```
[tangfx@localhost ~]$ docker load -i  /opt/kubernetes/server/bin/kube-apiserver.tar
c0de73ac9968: Loading layer [==================================================>] 1.312 MB/1.312 MB
65299b1581f3: Loading layer [==================================================>] 149.4 MB/149.4 MB
Loaded image: gcr.io/google_containers/kube-apiserver:5a24975398aa6a912695413b8acb65a8
[tangfx@localhost ~]$ tag=$(cat  /opt/kubernetes/server/bin/kube-apiserver.docker_tag ); echo $tag
5a24975398aa6a912695413b8acb65a8
[tangfx@localhost ~]$ img=$(docker images gcr.io/google_containers/kube-apiserver | awk 'NR==2 {print $1}'); echo $img
gcr.io/google_containers/kube-apiserver
[tangfx@localhost ~]$ docker load -i  /opt/kubernetes/server/bin/kube-controller-manager.tar
e3b49630145f: Loading layer [==================================================>] 131.6 MB/131.6 MB
Loaded image: gcr.io/google_containers/kube-controller-manager:82ba6120de2a37d3bbd84e8e2e6f2881
[tangfx@localhost ~]$ docker load -i  /opt/kubernetes/server/bin/kube-aggregator.tar
957a49fa0792: Loading layer [==================================================>] 55.26 MB/55.26 MB
Loaded image: gcr.io/google_containers/kube-aggregator:e10230754305ccf06f897f1c4047d0da
[tangfx@localhost ~]$ docker load -i  /opt/kubernetes/server/bin/kube-apiserver.tar 
Loaded image: gcr.io/google_containers/kube-apiserver:5a24975398aa6a912695413b8acb65a8
[tangfx@localhost ~]$ docker load -i  /opt/kubernetes/server/bin/kube-proxy.tar 
d30ca2ff3955: Loading layer [==================================================>]  42.2 MB/42.2 MB
dc588c8efd57: Loading layer [==================================================>] 4.743 MB/4.743 MB
3cedb9d0c8ed: Loading layer [==================================================>] 64.02 MB/64.02 MB
Loaded image: gcr.io/google_containers/kube-proxy:7e9fc4499ced8a0514f9403f8ffddfd7
[tangfx@localhost ~]$ docker load -i  /opt/kubernetes/server/bin/kube-scheduler.tar
7120ceaf7698: Loading layer [==================================================>] 75.65 MB/75.65 MB
Loaded image: gcr.io/google_containers/kube-scheduler:61551485f6f21737288567d9390634eb
[tangfx@localhost ~]$ docker images
REPOSITORY                                         TAG                                IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-proxy                7e9fc4499ced8a0514f9403f8ffddfd7   7a1b61b8f5d4        2 weeks ago         109.2 MB
gcr.io/google_containers/kube-apiserver            5a24975398aa6a912695413b8acb65a8   e14b1d5ee474        2 weeks ago         150.5 MB
gcr.io/google_containers/kube-controller-manager   82ba6120de2a37d3bbd84e8e2e6f2881   c7ad09fe3b82        2 weeks ago         132.7 MB
gcr.io/google_containers/kube-aggregator           e10230754305ccf06f897f1c4047d0da   72b975069762        2 weeks ago         56.37 MB
gcr.io/google_containers/kube-scheduler            61551485f6f21737288567d9390634eb   b55f2a2481b9        2 weeks ago         76.76 MB
```

Extract
```
[tangfx@localhost ~]$ sudo tar -C /opt -zxf kubernetes.tar.gz
[tangfx@localhost ~]$ sudo tar -C /opt -zxf  /opt/kubernetes/server/kubernetes-salt.tar.gz 
```

__Kubelet__ _systemd_ file
```
[tangfx@localhost ~]$ sudo sed -e 's%\(ExecStart=/usr/local/bin/kubelet.*\)%# \1\nExecStart=/opt/kubernetes/server/bin/kubelet "$DAEMON_ARGS"%' /opt/kubernetes/saltbase/salt/kubelet/kubelet.service | sudo tee /etc/systemd/system/kubelet.service
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

__Kubelet__ bin args
```
[tangfx@localhost ~]$ sudo sed '/^# test_args.*$/,$!d;s/{{daemon_args}}//;s%{{api_servers_with_port}}%--api-servers=http://127.0.0.1:8080%;s/{{debugging_handlers}}/--enable-debugging-handlers=true/;s/{{hostname_override}}/--hostname-override=10.64.33.80/;s/{{cloud_provider}} {{cloud_config}}//;s%{{config}} {{manifest_url}}%--pod-manifest-path=/etc/kubernetes/manifests%;s/{{pillar\[\x27allow_privileged\x27\]}}/true/;s/{{log_level}}/--v=2/;s/{{cluster_dns}} {{cluster_domain}}/--cluster-dns=10.123.240.10 --cluster-domain=cluster.local/;s%{{docker_root}} {{kubelet_root}}%--root-dir=/var/lib/kubelet%;s%{{non_masquerade_cidr}}%--non-masquerade-cidr=10.0.0.0/8%;s%{{pod_cidr}}%--pod-cidr=10.120.0.0/14%;s%{{ master_kubelet_args }}%--address=0.0.0.0 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --node-ip=10.64.33.80%;s/{{cpu_cfs_quota}}/--cpu-cfs-quota=true/;s/{{network_plugin}}//;s/{{kubelet_port}}/--port=10250/;s/{{ hairpin_mode }} {{enable_custom_metrics}}/--enable-custom-metrics=true/;s/{{node_labels}}//;s/{{babysit_daemons}}//;s/{{eviction_hard}} {{kubelet_auth}} {{feature_gates}}/--eviction-hard=memory.available<100Mi --anonymous-auth=true --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true/;s/{{test_args}}/--cgroup-driver=systemd --authorization-mode=AlwaysAllow/' /opt/kubernetes/saltbase/salt/kubelet/default | sudo tee /etc/sysconfig/kubelet
# test_args has to be kept at the end, so they'll overwrite any prior configuration
DAEMON_ARGS=" --api-servers=http://127.0.0.1:8080 --enable-debugging-handlers=true --hostname-override=10.64.33.80  --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --v=2 --cluster-dns=10.123.240.10 --cluster-domain=cluster.local --root-dir=/var/lib/kubelet  --non-masquerade-cidr=10.0.0.0/8 --pod-cidr=10.120.0.0/14 --address=0.0.0.0 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --node-ip=10.64.33.80 --cpu-cfs-quota=true  --port=10250 --enable-custom-metrics=true   --eviction-hard=memory.available<100Mi --anonymous-auth=true --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --cgroup-driver=systemd --authorization-mode=AlwaysAllow"
```

Check _kubelet_ service
```
[tangfx@localhost ~]$ sudo systemctl enable kubelet.service
[tangfx@localhost ~]$ sudo systemctl start kubelet.service
[tangfx@localhost ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service
[tangfx@localhost ~]$ sudo systemctl restart kubelet.service
[tangfx@localhost ~]$ sudo systemctl -l status kubelet.service
```