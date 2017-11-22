
## releases

get binary packages
```
fanhonglingdeMacBook-Pro:v1.8.4 fanhongling$ curl -jkSL -C- https://dl.k8s.io/v1.8.4/kubernetes-server-linux-amd64.tar.gz -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     27      0  0:00:05  0:00:05 --:--:--    48
100  385M  100  385M    0     0  85228      0  1:18:59  1:18:59 --:--:-- 40587
```

```
fanhonglingdeMacBook-Pro:v1.8.4 fanhongling$ tar -ztf kubernetes-server-linux-amd64.tar.gz 
kubernetes/
kubernetes/server/
kubernetes/server/bin/
kubernetes/server/bin/kube-proxy.tar
kubernetes/server/bin/kubectl
kubernetes/server/bin/kube-apiserver.tar
kubernetes/server/bin/apiextensions-apiserver
kubernetes/server/bin/kubeadm
kubernetes/server/bin/kube-apiserver.docker_tag
kubernetes/server/bin/kube-aggregator.docker_tag
kubernetes/server/bin/hyperkube
kubernetes/server/bin/kube-aggregator.tar
kubernetes/server/bin/cloud-controller-manager
kubernetes/server/bin/kube-scheduler.docker_tag
kubernetes/server/bin/kubefed
kubernetes/server/bin/kube-controller-manager
kubernetes/server/bin/kube-proxy
kubernetes/server/bin/kube-scheduler
kubernetes/server/bin/cloud-controller-manager.tar
kubernetes/server/bin/kube-scheduler.tar
kubernetes/server/bin/kubelet
kubernetes/server/bin/kube-controller-manager.docker_tag
kubernetes/server/bin/kube-apiserver
kubernetes/server/bin/kube-aggregator
kubernetes/server/bin/kube-controller-manager.tar
kubernetes/server/bin/cloud-controller-manager.docker_tag
kubernetes/server/bin/kube-proxy.docker_tag
kubernetes/LICENSES
kubernetes/kubernetes-src.tar.gz
kubernetes/addons/
```

get addons manifests
```
fanhonglingdeMacBook-Pro:v1.8.4 fanhongling$ curl -jkSL -C- https://dl.k8s.io/v1.8.4/kubernetes.tar.gz -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     17      0  0:00:09  0:00:09 --:--:--    36
100 4029k  100 4029k    0     0  87544      0  0:00:47  0:00:47 --:--:--  189k
```

```
fanhonglingdeMacBook-Pro:v1.8.4 fanhongling$ tar -ztf kubernetes.tar.gz 
### snippets ###
kubernetes/cluster/addons/
kubernetes/cluster/addons/dns/
kubernetes/cluster/addons/dns/kubedns-controller.yaml.in
kubernetes/cluster/addons/dns/kubedns-svc.yaml.base
kubernetes/cluster/addons/dns/OWNERS
kubernetes/cluster/addons/dns/kubedns-controller.yaml.base
kubernetes/cluster/addons/dns/Makefile
kubernetes/cluster/addons/dns/kubedns-controller.yaml.sed
kubernetes/cluster/addons/dns/kubedns-cm.yaml
kubernetes/cluster/addons/dns/transforms2salt.sed
kubernetes/cluster/addons/dns/kubedns-svc.yaml.sed
kubernetes/cluster/addons/dns/transforms2sed.sed
kubernetes/cluster/addons/dns/kubedns-sa.yaml
kubernetes/cluster/addons/dns/kubedns-svc.yaml.in
kubernetes/cluster/addons/dns/README.md
### snippets ###
```

Extract k8s images
```
[vagrant@openshiftdev ~]$ rm -rf ./kubernetes/*
[vagrant@openshiftdev ~]$ tar -C /home/vagrant/ -zxf /vagrant_f/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/v1.8.4/kubernetes-server-linux-amd64.tar.gz 
[vagrant@openshiftdev ~]$ ls -1 kubernetes/server/bin/
apiextensions-apiserver
cloud-controller-manager
cloud-controller-manager.docker_tag
cloud-controller-manager.tar
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

```
[vagrant@openshiftdev ~]$ bindir=./kubernetes/server/bin; for i in $(ls $bindir/*.tar); do docker load -i $i; name=$(basename $i .tar); tag=$(cat $bindir/$name.docker_tag); repo=gcr.io/google_containers/$name; img=$(docker images -q $repo:$tag); [[ -n $img ]] && docker images $repo:$tag || (echo "...failed to load" && docker images $repo); done
0271b8eebde3: Loading layer [==================================================>] 1.338 MB/1.338 MB
93df51b9594c: Loading layer [==================================================>] 109.1 MB/109.1 MB
Loaded image: gcr.io/google_containers/cloud-controller-manager:v1.8.4
REPOSITORY                                          TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/cloud-controller-manager   v1.8.4              1fe353851780        44 hours ago        110.2 MB
39aab9798d30: Loading layer [==================================================>] 53.86 MB/53.86 MB
Loaded image: gcr.io/google_containers/kube-aggregator:v1.8.4
REPOSITORY                                 TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-aggregator   v1.8.4              0a974eef74bb        44 hours ago        54.99 MB
f72e92775dd7: Loading layer [==================================================>] 193.2 MB/193.2 MB
Loaded image: gcr.io/google_containers/kube-apiserver:v1.8.4
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver   v1.8.4              10a052dccbc5        44 hours ago        194.4 MB
4f92df6d8677: Loading layer [==================================================>] 128.2 MB/128.2 MB
Loaded image: gcr.io/google_containers/kube-controller-manager:v1.8.4
REPOSITORY                                         TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-controller-manager   v1.8.4              7058ac4d4af5        44 hours ago        129.3 MB
08ae86c4c4c9: Loading layer [==================================================>] 42.05 MB/42.05 MB
48a108f012f8: Loading layer [==================================================>] 5.045 MB/5.045 MB
9fc6ccb688b9: Loading layer [==================================================>] 47.93 MB/47.93 MB
Loaded image: gcr.io/google_containers/kube-proxy:v1.8.4
REPOSITORY                            TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-proxy   v1.8.4              65a61c14e8c2        44 hours ago        93.2 MB
da6851a1e488: Loading layer [==================================================>] 53.85 MB/53.85 MB
Loaded image: gcr.io/google_containers/kube-scheduler:v1.8.4
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-scheduler   v1.8.4              0d985fed7f95        44 hours ago        54.98 MB
```

```
[vagrant@openshiftdev ~]$ docker tag gcr.io/google_containers/kube-apiserver:v1.8.4 gcr.io/google_containers/kube-apiserver-amd64:v1.8.4
[vagrant@openshiftdev ~]$ docker tag gcr.io/google_containers/kube-controller-manager:v1.8.4 gcr.io/google_containers/kube-controller-manager-amd64:v1.8.4
[vagrant@openshiftdev ~]$ docker tag gcr.io/google_containers/kube-scheduler:v1.8.4 gcr.io/google_containers/kube-scheduler-amd64:v1.8.4
[vagrant@kubedev-10-64-33-82 ~]$ docker tag gcr.io/google_containers/kube-proxy:v1.8.4 gcr.io/google_containers/kube-proxy-amd64:v1.8.4
```

```
[vagrant@openshiftdev ~]$ docker images gcr.io/google_containers/*
REPOSITORY                                               TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver-amd64            v1.8.4              10a052dccbc5        45 hours ago        194.4 MB
gcr.io/google_containers/kube-apiserver                  v1.8.4              10a052dccbc5        45 hours ago        194.4 MB
gcr.io/google_containers/cloud-controller-manager        v1.8.4              1fe353851780        45 hours ago        110.2 MB
gcr.io/google_containers/kube-controller-manager-amd64   v1.8.4              7058ac4d4af5        45 hours ago        129.3 MB
gcr.io/google_containers/kube-controller-manager         v1.8.4              7058ac4d4af5        45 hours ago        129.3 MB
gcr.io/google_containers/kube-proxy                      v1.8.4              65a61c14e8c2        45 hours ago        93.2 MB
gcr.io/google_containers/kube-aggregator                 v1.8.4              0a974eef74bb        45 hours ago        54.99 MB
gcr.io/google_containers/kube-scheduler-amd64            v1.8.4              0d985fed7f95        45 hours ago        54.98 MB
gcr.io/google_containers/kube-scheduler                  v1.8.4              0d985fed7f95        45 hours ago        54.98 MB
gcr.io/google_containers/cloud-controller-manager        v1.7.4              a42a0fd85571        3 months ago        125.3 MB
gcr.io/google_containers/kube-controller-manager-amd64   v1.7.4              d2adddc4b1cb        3 months ago        138 MB
gcr.io/google_containers/kube-controller-manager         v1.7.4              d2adddc4b1cb        3 months ago        138 MB
gcr.io/google_containers/kube-apiserver-amd64            v1.7.4              5260ecb5129c        3 months ago        186.1 MB
gcr.io/google_containers/kube-apiserver                  v1.7.4              5260ecb5129c        3 months ago        186.1 MB
gcr.io/google_containers/kube-proxy                      v1.7.4              0f3bf654ec61        3 months ago        114.7 MB
gcr.io/google_containers/kube-aggregator                 v1.7.4              c1752dd09c50        3 months ago        59.16 MB
gcr.io/google_containers/kube-scheduler                  v1.7.4              b1cd468ba656        3 months ago        77.2 MB
gcr.io/google_containers/kube-scheduler-amd64            v1.7.4              b1cd468ba656        3 months ago        77.2 MB
gcr.io/google_containers/etcd-amd64                      3.0.17              243830dae7dd        9 months ago        168.9 MB
gcr.io/google_containers/etcd                            3.0.17              243830dae7dd        9 months ago        168.9 MB
gcr.io/google_containers/pause-amd64                     3.0                 99e59f495ffa        18 months ago       746.9 kB
```

## http yum repo

docker
```
[vagrant@openshiftdev ~]$ docker ps -f name=gofileserver
CONTAINER ID        IMAGE                                 COMMAND             CREATED             STATUS              PORTS                      NAMES
aa310a752015        docker.io/tangfeixiong/gofileserver   "/gofileserver"     3 days ago          Up 2 hours          0.0.0.0:48080->48080/tcp   gofileserver
```

```
[vagrant@openshiftdev ~]$ curl http://10.64.33.82:48080/vagrant_f/16-mirror/
<pre>
<a href="centos/">centos/</a>
<a href="fedora/">fedora/</a>
<a href="http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/">http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/</a>
<a href="ubuntu/">ubuntu/</a>
</pre>
```

k8s repos
```
[vagrant@openshiftdev ~]$ cat /etc/yum.repos.d/kubernetes.repo 
[kubernetes]
name=Kubernetes
#baseurl=http://yum.kubernetes.io/repos/kubernetes-el7-x86_64
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
        https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg

[kubernetes-local]
name=Kubernetes offline mirror
baseurl=file:///vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=0
gpgcheck=0

[kubernetes-10-64-33-82]
name=Kubernetes local mirror
baseurl=http://10.64.33.82:48080/vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=0
gpgcheck=0
```

## kubeadm

upgrade
```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=\* --enablerepo=kubernetes-local --enablerepo=kubernetes-10-64-33-82 list | grep kubeadm
kubeadm.x86_64                       1.8.2-0                           @kubernetes-local
kubeadm.x86_64                       1.8.4-0                           kubernetes-10-64-33-82
```

Upgrade from v1.8.2
```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=kubernetes --enablerepo=kubernetes-local --enablerepo=kubernetes-10-64-33-82 upgrade kubeadm
Loaded plugins: fastestmirror
base                                                                                                         | 3.6 kB  00:00:00     
dl.google.com_linux_chrome_rpm_stable_x86_64                                                                 |  951 B  00:00:00     
epel/x86_64/metalink                                                                                         | 6.0 kB  00:00:00     
extras                                                                                                       | 3.4 kB  00:00:00     
google-chrome                                                                                                |  951 B  00:00:00     
openshift-deps                                                                                               | 2.9 kB  00:00:00     
origin-deps-rhel7                                                                                            | 2.9 kB  00:00:00     
updates                                                                                                      | 3.4 kB  00:00:00     
Loading mirror speeds from cached hostfile
 * base: ftp.sjtu.edu.cn
 * epel: mirror.lzu.edu.cn
 * extras: mirrors.163.com
 * updates: mirrors.163.com
Resolving Dependencies
--> Running transaction check
---> Package kubeadm.x86_64 0:1.8.2-0 will be updated
---> Package kubeadm.x86_64 0:1.8.4-0 will be an update
--> Finished Dependency Resolution

Dependencies Resolved

====================================================================================================================================
 Package                    Arch                      Version                       Repository                                 Size
====================================================================================================================================
Updating:
 kubeadm                    x86_64                    1.8.4-0                       kubernetes-10-64-33-82                     15 M

Transaction Summary
====================================================================================================================================
Upgrade  1 Package

Total download size: 15 M
Is this ok [y/d/N]: y
Downloading packages:
No Presto metadata available for kubernetes-10-64-33-82
aeaad1e283c54876b759a089f152228d7cd4c049f271125c23623995b8e76f96-kubeadm-1.8.4-0.x86_64.rpm                  |  15 MB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Updating   : kubeadm-1.8.4-0.x86_64                                                                                           1/2 
  Cleanup    : kubeadm-1.8.2-0.x86_64                                                                                           2/2 
  Verifying  : kubeadm-1.8.4-0.x86_64                                                                                           1/2 
  Verifying  : kubeadm-1.8.2-0.x86_64                                                                                           2/2 

Updated:
  kubeadm.x86_64 0:1.8.4-0                                                                                                          

Complete!
```

```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=kubernetes --enablerepo=kubernetes-local --enablerepo=kubernetes-10-64-33-82 upgrade kubelet kubectl
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: ftp.sjtu.edu.cn
 * epel: mirror.lzu.edu.cn
 * extras: mirrors.163.com
 * updates: mirrors.163.com
Resolving Dependencies
--> Running transaction check
---> Package kubectl.x86_64 0:1.8.2-0 will be updated
---> Package kubectl.x86_64 0:1.8.4-0 will be an update
---> Package kubelet.x86_64 0:1.8.2-0 will be updated
---> Package kubelet.x86_64 0:1.8.4-0 will be an update
--> Finished Dependency Resolution

Dependencies Resolved

====================================================================================================================================
 Package                    Arch                      Version                       Repository                                 Size
====================================================================================================================================
Updating:
 kubectl                    x86_64                    1.8.4-0                       kubernetes-10-64-33-82                    7.3 M
 kubelet                    x86_64                    1.8.4-0                       kubernetes-10-64-33-82                     16 M

Transaction Summary
====================================================================================================================================
Upgrade  2 Packages

Total download size: 23 M
Is this ok [y/d/N]: y
Downloading packages:
No Presto metadata available for kubernetes-10-64-33-82
(1/2): a9db28728641ddbf7f025b8b496804d82a396d0ccb178fffd124623fb2f999ea-kubectl-1.8.4-0.x86_64.rpm           | 7.3 MB  00:00:00     
(2/2): 1acca81eb5cf99453f30466876ff03146112b7f12c625cb48f12508684e02665-kubelet-1.8.4-0.x86_64.rpm           |  16 MB  00:00:00     
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                44 MB/s |  23 MB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Updating   : kubectl-1.8.4-0.x86_64                                                                                           1/4 
  Updating   : kubelet-1.8.4-0.x86_64                                                                                           2/4 
  Cleanup    : kubectl-1.8.2-0.x86_64                                                                                           3/4 
  Cleanup    : kubelet-1.8.2-0.x86_64                                                                                           4/4 
  Verifying  : kubelet-1.8.4-0.x86_64                                                                                           1/4 
  Verifying  : kubectl-1.8.4-0.x86_64                                                                                           2/4 
  Verifying  : kubectl-1.8.2-0.x86_64                                                                                           3/4 
  Verifying  : kubelet-1.8.2-0.x86_64                                                                                           4/4 

Updated:
  kubectl.x86_64 0:1.8.4-0                                         kubelet.x86_64 0:1.8.4-0                                        

Complete!
```

```
[vagrant@openshiftdev ~]$ kubectl version
Client Version: version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.4", GitCommit:"9befc2b8928a9426501d3bf62f72849d5cbcd5a3", GitTreeState:"clean", BuildDate:"2017-11-20T05:28:34Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"7", GitVersion:"v1.7.4", GitCommit:"793658f2d7ca7f064d2bdf606519f9fe1229c381", GitTreeState:"clean", BuildDate:"2017-08-17T08:30:51Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
[vagrant@openshiftdev ~]$ kubelet --version
Kubernetes v1.8.4
```


## Rename node

change hostname
```
[vagrant@openshiftdev ~]$ echo "10.64.33.82    kubedev-10-64-33-82" | sudo tee -a /etc/hosts
10.64.33.82    kubedev-10-64-33-82
[vagrant@openshiftdev ~]$ echo "kubedev-10-64-33-82" | sudo tee /etc/hostname
kubedev-10-64-33-82
[vagrant@openshiftdev ~]$ sudo hostname kubedev-10-64-33-82
```

```
[vagrant@openshiftdev ~]$ sudo sed -i 's/\(ExecStart=.*ARGS\)/\1 --hostname-override=kubedev-10-64-33-82 --node-ip=10.64.33.82/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

```
[vagrant@openshiftdev ~]$ sudo rm /etc/kubernetes/*
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo kubeadm init --apiserver-advertise-address 10.64.33.82 --apiserver-bind-port 443 --pod-network-cidr 10.244.0.0/16 --kubernetes-version v1.8.4 --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.8.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [kubedev-10-64-33-82 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.64.33.82]
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
[vagrant@openshiftdev ~]$ cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"
Environment="KUBELET_CERTIFICATE_ARGS=--rotate-certificates=true --cert-dir=/var/lib/kubelet/pki"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false --hostname-override=kubedev-10-64-33-82 --node-ip=10.64.33.82
[vagrant@openshiftdev ~]$ kubectl get nodes
error: Error loading config file "/etc/kubernetes/admin.conf": open /etc/kubernetes/admin.conf: permission denied
[vagrant@openshiftdev ~]$ sudo chmod o+r /etc/kubernetes/admin.conf 
[vagrant@openshiftdev ~]$ kubectl get nodes
Unable to connect to the server: x509: certificate signed by unknown authority (possibly because of "crypto/rsa: verification error" while trying to verify candidate authority certificate "kubernetes")
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: failed (Result: exit-code) since Wed 2017-11-22 03:46:19 UTC; 4min 56s ago
     Docs: http://kubernetes.io/docs/
  Process: 31296 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false --hostname-override=kubedev-10-64-33-82 --node-ip=10.64.33.82 (code=exited, status=1/FAILURE)
 Main PID: 31296 (code=exited, status=1/FAILURE)

Nov 22 03:46:16 kubedev-10-64-33-82 kubelet[31296]: W1122 03:46:16.654020   31296 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 22 03:46:16 kubedev-10-64-33-82 kubelet[31296]: E1122 03:46:16.654122   31296 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Nov 22 03:46:18 kubedev-10-64-33-82 kubelet[31296]: I1122 03:46:18.139428   31296 kubelet_node_status.go:280] Setting node annotation to enable volume controller attach/detach
Nov 22 03:46:18 kubedev-10-64-33-82 kubelet[31296]: I1122 03:46:18.141189   31296 kubelet_node_status.go:83] Attempting to register node kubedev-10-64-33-82
Nov 22 03:46:18 kubedev-10-64-33-82 kubelet[31296]: E1122 03:46:18.142940   31296 kubelet_node_status.go:107] Unable to register node "kubedev-10-64-33-82" with API server: nodes "kubedev-10-64-33-82" is forbidden: node openshiftdev.local cannot modify node kubedev-10-64-33-82
Nov 22 03:46:19 kubedev-10-64-33-82 systemd[1]: Stopping kubelet: The Kubernetes Node Agent...
Nov 22 03:46:19 kubedev-10-64-33-82 systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 22 03:46:19 kubedev-10-64-33-82 systemd[1]: Stopped kubelet: The Kubernetes Node Agent.
Nov 22 03:46:19 kubedev-10-64-33-82 systemd[1]: Unit kubelet.service entered failed state.
Nov 22 03:46:19 kubedev-10-64-33-82 systemd[1]: kubelet.service failed.
```

```
[vagrant@openshiftdev ~]$ sudo systemctl restart docker.service
```

```
[vagrant@openshiftdev ~]$ kubectl describe nodes/kubedev-10-64-33-82
Name:               kubedev-10-64-33-82
Roles:              <none>
Labels:             beta.kubernetes.io/arch=amd64
                    beta.kubernetes.io/os=linux
                    kubernetes.io/hostname=kubedev-10-64-33-82
Annotations:        alpha.kubernetes.io/provided-node-ip=10.64.33.82
                    volumes.kubernetes.io/controller-managed-attach-detach=true
Taints:             <none>
CreationTimestamp:  Wed, 22 Nov 2017 03:55:31 +0000
Conditions:
  Type             Status  LastHeartbeatTime                 LastTransitionTime                Reason                       Message
  ----             ------  -----------------                 ------------------                ------                       -------
  OutOfDisk        False   Wed, 22 Nov 2017 03:56:31 +0000   Wed, 22 Nov 2017 03:55:27 +0000   KubeletHasSufficientDisk     kubelet has sufficient disk space available
  MemoryPressure   False   Wed, 22 Nov 2017 03:56:31 +0000   Wed, 22 Nov 2017 03:55:27 +0000   KubeletHasSufficientMemory   kubelet has sufficient memory available
  DiskPressure     False   Wed, 22 Nov 2017 03:56:31 +0000   Wed, 22 Nov 2017 03:55:27 +0000   KubeletHasNoDiskPressure     kubelet has no disk pressure
  Ready            False   Wed, 22 Nov 2017 03:56:31 +0000   Wed, 22 Nov 2017 03:55:27 +0000   KubeletNotReady              runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Addresses:
  InternalIP:  10.64.33.82
  Hostname:    kubedev-10-64-33-82
Capacity:
 cpu:     1
 memory:  3883128Ki
 pods:    110
Allocatable:
 cpu:     1
 memory:  3780728Ki
 pods:    110
System Info:
 Machine ID:                 50e9e11ee2d84dc98dcfc8d70a476ae7
 System UUID:                A6CAFB25-3104-4470-9098-847053BB0273
 Boot ID:                    05188f76-b062-4b2c-86f1-cd66b21f26fb
 Kernel Version:             3.10.0-229.el7.x86_64
 OS Image:                   CentOS Linux 7 (Core)
 Operating System:           linux
 Architecture:               amd64
 Container Runtime Version:  docker://1.12.6
 Kubelet Version:            v1.8.4
 Kube-Proxy Version:         v1.8.4
ExternalID:                  kubedev-10-64-33-82
Non-terminated Pods:         (4 in total)
  Namespace                  Name                                           CPU Requests  CPU Limits  Memory Requests  Memory Limits
  ---------                  ----                                           ------------  ----------  ---------------  -------------
  kube-system                etcd-kubedev-10-64-33-82                       0 (0%)        0 (0%)      0 (0%)           0 (0%)
  kube-system                kube-apiserver-kubedev-10-64-33-82             250m (25%)    0 (0%)      0 (0%)           0 (0%)
  kube-system                kube-controller-manager-kubedev-10-64-33-82    200m (20%)    0 (0%)      0 (0%)           0 (0%)
  kube-system                kube-scheduler-kubedev-10-64-33-82             100m (10%)    0 (0%)      0 (0%)           0 (0%)
Allocated resources:
  (Total limits may be over 100 percent, i.e., overcommitted.)
  CPU Requests  CPU Limits  Memory Requests  Memory Limits
  ------------  ----------  ---------------  -------------
  550m (55%)    0 (0%)      0 (0%)           0 (0%)
Events:
  Type    Reason                   Age                From                          Message
  ----    ------                   ----               ----                          -------
  Normal  Starting                 11m                kubelet, kubedev-10-64-33-82  Starting kubelet.
  Normal  NodeAllocatableEnforced  11m                kubelet, kubedev-10-64-33-82  Updated Node Allocatable limit across pods
  Normal  NodeHasSufficientDisk    11m (x8 over 11m)  kubelet, kubedev-10-64-33-82  Node kubedev-10-64-33-82 status is now: NodeHasSufficientDisk
  Normal  NodeHasSufficientMemory  11m (x8 over 11m)  kubelet, kubedev-10-64-33-82  Node kubedev-10-64-33-82 status is now: NodeHasSufficientMemory
  Normal  NodeHasNoDiskPressure    11m (x7 over 11m)  kubelet, kubedev-10-64-33-82  Node kubedev-10-64-33-82 status is now: NodeHasNoDiskPressure
  Normal  NodeAllocatableEnforced  3m                 kubelet, kubedev-10-64-33-82  Updated Node Allocatable limit across pods
  Normal  NodeHasSufficientDisk    3m (x8 over 3m)    kubelet, kubedev-10-64-33-82  Node kubedev-10-64-33-82 status is now: NodeHasSufficientDisk
  Normal  NodeHasSufficientMemory  3m (x8 over 3m)    kubelet, kubedev-10-64-33-82  Node kubedev-10-64-33-82 status is now: NodeHasSufficientMemory
  Normal  NodeHasNoDiskPressure    3m (x7 over 3m)    kubelet, kubedev-10-64-33-82  Node kubedev-10-64-33-82 status is now: NodeHasNoDiskPressure
```

```
[vagrant@openshiftdev ~]$ kubectl get nodes
NAME                  STATUS     ROLES     AGE       VERSION
kubedev-10-64-33-82   NotReady   <none>    9m        v1.8.4
```

```
[vagrant@openshiftdev ~]$ kubectl get all --all-namespaces
NAMESPACE     NAME                                             READY     STATUS    RESTARTS   AGE
kube-system   po/etcd-kubedev-10-64-33-82                      1/1       Running   4          12m
kube-system   po/kube-apiserver-kubedev-10-64-33-82            1/1       Running   4          12m
kube-system   po/kube-controller-manager-kubedev-10-64-33-82   1/1       Running   6          13m
kube-system   po/kube-scheduler-kubedev-10-64-33-82            1/1       Running   4          12m
```