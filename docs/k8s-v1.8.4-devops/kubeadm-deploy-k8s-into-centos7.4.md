# instruction

networking
```
[vagrant@kubedev-10-64-33-82 ~]$ ip a show enp0s8
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 08:00:27:ff:85:a2 brd ff:ff:ff:ff:ff:ff
    inet 10.64.33.82/24 brd 10.64.33.255 scope global enp0s8
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:feff:85a2/64 scope link 
       valid_lft forever preferred_lft forever
```

## Prerequisites

### firewall and selinux

firewall
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo systemctl is-active firewalld.service
unknown
```

selinux
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo getenforce
Permissive
```

### hostname and hosts

hostname
```
[vagrant@openshiftdev ~]$ echo "kubedev-10-64-33-82" | sudo tee /etc/hostname
kubedev-10-64-33-82
```

```
[vagrant@openshiftdev ~]$ sudo hostname --file /etc/hostname
```

```
[vagrant@openshiftdev ~]$ hostname
kubedev-10-64-33-82
```

hosts
```
[vagrant@openshiftdev ~]$ echo "10.64.33.82    kubedev-10-64-33-82" | sudo tee -a /etc/hosts
10.64.33.82    kubedev-10-64-33-82
```

```
[vagrant@openshiftdev ~]$ cat /etc/hosts
127.0.0.1   openshiftdev.local openshiftdev localhost localhost.localdomain localhost4 localhost4.localdomain4
::1         localhost localhost.localdomain localhost6 localhost6.localdomain6
10.64.33.82    kubedev-10-64-33-82
```

## install kubeadm

### local kubernetes yum Repository

add
```
[vagrant@kubedev-10-64-33-82 ~]$ cat /etc/yum.repos.d/kubernetes.repo 
[kubernetes]
name=Kubernetes
#baseurl=http://yum.kubernetes.io/repos/kubernetes-el7-x86_64
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
        https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg

[kubernetes-local-files]
name=Kubernetes local files
baseurl=file:///vagrant_f/16-mirror/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/
enabled=0
gpgcheck=0

[kubernetes-local-http]
name=Kubernetes local http
baseurl=http://10.64.33.195:48080/vagrant_f/16-mirror/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/
enabled=0
gpgcheck=0
```

### install

kubeadm v1.8.4
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo yum --disablerepo=kubernetes --enablerepo=kubernetes-local-files search kubeadm
Loaded plugins: fastestmirror
base                                                                                                         | 3.6 kB  00:00:00     
dl.google.com_linux_chrome_rpm_stable_x86_64                                                                 |  951 B  00:00:00     
dockerrepo                                                                                                   | 2.9 kB  00:00:00     
epel/x86_64/metalink                                                                                         | 5.9 kB  00:00:00     
epel                                                                                                         | 4.7 kB  00:00:00     
extras                                                                                                       | 3.4 kB  00:00:00     
google-chrome                                                                                                |  951 B  00:00:00     
kubernetes-local-files                                                                                       | 1.4 kB  00:00:00     
openshift-deps                                                                                               | 2.9 kB  00:00:00     
origin-deps-rhel7                                                                                            | 2.9 kB  00:00:00     
updates                                                                                                      | 3.4 kB  00:00:00     
(1/9): base/7/x86_64/group_gz                                                                                | 156 kB  00:00:00     
(2/9): epel/x86_64/group_gz                                                                                  | 266 kB  00:00:02     
(3/9): extras/7/x86_64/primary_db                                                                            | 130 kB  00:00:00     
(4/9): kubernetes-local-files/primary                                                                        |  13 kB  00:00:00     
(5/9): epel/x86_64/updateinfo                                                                                | 851 kB  00:00:03     
(6/9): base/7/x86_64/primary_db                                                                              | 5.7 MB  00:00:04     
(7/9): updates/7/x86_64/primary_db                                                                           | 3.6 MB  00:00:03     
(8/9): dockerrepo/7/primary_db                                                                               |  34 kB  00:00:06     
(9/9): epel/x86_64/primary_db                                                                                | 6.1 MB  00:00:14     
(1/2): dl.google.com_linux_chrome_rpm_stable_x86_64/primary                                                  | 2.0 kB  00:00:00     
(2/2): google-chrome/primary                                                                                 | 2.0 kB  00:00:00     
Determining fastest mirrors
 * base: mirrors.sohu.com
 * epel: mirrors.ustc.edu.cn
 * extras: mirrors.sohu.com
 * updates: mirrors.sohu.com
dl.google.com_linux_chrome_rpm_stable_x86_64                                                                                    3/3
google-chrome                                                                                                                   3/3
kubernetes-local-files                                                                                                        84/84
======================================================= N/S matched: kubeadm =======================================================
kubeadm.x86_64 : Command-line utility for administering a Kubernetes cluster.

  Name and summary matches only, use "search all" for everything.
[vagrant@kubedev-10-64-33-82 ~]$ sudo yum --disablerepo=kubernetes --enablerepo=kubernetes-local-files install -y kubeadm
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.sohu.com
 * epel: mirrors.ustc.edu.cn
 * extras: mirrors.sohu.com
 * updates: mirrors.sohu.com
Resolving Dependencies
--> Running transaction check
---> Package kubeadm.x86_64 0:1.8.4-0 will be installed
--> Processing Dependency: kubelet >= 1.6.0 for package: kubeadm-1.8.4-0.x86_64
--> Processing Dependency: kubectl >= 1.6.0 for package: kubeadm-1.8.4-0.x86_64
--> Processing Dependency: kubernetes-cni for package: kubeadm-1.8.4-0.x86_64
--> Running transaction check
---> Package kubectl.x86_64 0:1.8.4-0 will be installed
---> Package kubelet.x86_64 0:1.8.4-0 will be installed
---> Package kubernetes-cni.x86_64 0:0.5.1-1 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

====================================================================================================================================
 Package                          Arch                     Version                   Repository                                Size
====================================================================================================================================
Installing:
 kubeadm                          x86_64                   1.8.4-0                   kubernetes-local-files                    15 M
Installing for dependencies:
 kubectl                          x86_64                   1.8.4-0                   kubernetes-local-files                   7.3 M
 kubelet                          x86_64                   1.8.4-0                   kubernetes-local-files                    16 M
 kubernetes-cni                   x86_64                   0.5.1-1                   kubernetes-local-files                   7.4 M

Transaction Summary
====================================================================================================================================
Install  1 Package (+3 Dependent packages)

Total download size: 45 M
Installed size: 242 M
Downloading packages:
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                               4.9 MB/s |  45 MB  00:00:09     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : kubernetes-cni-0.5.1-1.x86_64                                                                                    1/4 
  Installing : kubelet-1.8.4-0.x86_64                                                                                           2/4 
  Installing : kubectl-1.8.4-0.x86_64                                                                                           3/4 
  Installing : kubeadm-1.8.4-0.x86_64                                                                                           4/4 
  Verifying  : kubelet-1.8.4-0.x86_64                                                                                           1/4 
  Verifying  : kubernetes-cni-0.5.1-1.x86_64                                                                                    2/4 
  Verifying  : kubeadm-1.8.4-0.x86_64                                                                                           3/4 
  Verifying  : kubectl-1.8.4-0.x86_64                                                                                           4/4 

Installed:
  kubeadm.x86_64 0:1.8.4-0                                                                                                          

Dependency Installed:
  kubectl.x86_64 0:1.8.4-0                 kubelet.x86_64 0:1.8.4-0                 kubernetes-cni.x86_64 0:0.5.1-1                

Complete!
```

### installed

kubelet
```
[vagrant@kubedev-10-64-33-82 ~]$ ls /etc/systemd/system/kubelet.service*
/etc/systemd/system/kubelet.service

/etc/systemd/system/kubelet.service.d:
10-kubeadm.conf
```

```
[vagrant@kubedev-10-64-33-82 ~]$ ls -R /etc/kubernetes/
/etc/kubernetes/:
manifests

/etc/kubernetes/manifests:
```

```
[vagrant@kubedev-10-64-33-82 ~]$ which kubelet
/bin/kubelet
```

```
[vagrant@kubedev-10-64-33-82 ~]$ which kubectl
/bin/kubectl
```

CNI
```
[vagrant@kubedev-10-64-33-82 ~]$ ls /opt/cni/bin/
bridge  cnitool  dhcp  flannel  host-local  ipvlan  loopback  macvlan  noop  ptp  tuning
```

## docker images

kubernetes
```
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-apiserver.tar
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-controller-manager.tar 
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-scheduler.tar 
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-proxy.tar 
```

```
[vagrant@kubedev-10-64-33-82 ~]$ docker tag gcr.io/google_containers/kube-apiserver:v1.8.4 gcr.io/google_containers/kube-apiserver-amd64:v1.8.4
[vagrant@kubedev-10-64-33-82 ~]$ docker tag gcr.io/google_containers/kube-controller-manager:v1.8.4 gcr.io/google_containers/kube-controller-manager-amd64:v1.8.4
[vagrant@kubedev-10-64-33-82 ~]$ docker tag gcr.io/google_containers/kube-scheduler:v1.8.4 gcr.io/google_containers/kube-scheduler-amd64:v1.8.4
[vagrant@kubedev-10-64-33-82 ~]$ docker tag gcr.io/google_containers/kube-proxy:v1.8.4 gcr.io/google_containers/kube-proxy-amd64:v1.8.4
```

ETCD
```
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fetcd0x2E3.0.17.tar 
```

```
[vagrant@kubedev-10-64-33-82 ~]$ docker tag gcr.io/google_containers/etcd:3.0.17 gcr.io/google_containers/etcd-amd64:3.0.17
```

k9s addons
```
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-dnsmasq-nanny-amd640x3A1.14.5.tar 
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-kube-dns-amd640x3A1.14.5.tar 
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-sidecar-amd640x3A1.14.5.tar 
```

POD base
```
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar 
```

CNI
```
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/quay.io0x2Fcoreos0x2Fflannel0x3Av0.9.0-amd64.tar 
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-kube0x3A2.1.1.tar 
[vagrant@kubedev-10-64-33-82 ~]$ docker load -i /vagrant_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-npc0x3A2.1.1.tar 
```

## Deploy k8s

before
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo sed -i 's/\(ExecStart=.*ARGS\)/\1 --fail-swap-on=false/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

init
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
```

```
[vagrant@kubedev-10-64-33-82 ~]$ cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
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
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false --hostname-override=kubedev-10-64-33-82 --node-ip=10.64.33.82 --pod-cidr=10.244.0.0/16
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo systemctl start kubelet.service
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo chmod o+r /etc/kubernetes/admin.conf 
[vagrant@kubedev-10-64-33-82 ~]$ export KUBECONFIG=/etc/kubernetes/admin.conf 
```

```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl get all --namespace=kube-system
NAME                 DESIRED   CURRENT   READY     UP-TO-DATE   AVAILABLE   NODE SELECTOR                   AGE
ds/kube-flannel-ds   1         1         0         1            0           beta.kubernetes.io/arch=amd64   1h
ds/kube-flannel-ds   1         1         0         1            0           beta.kubernetes.io/arch=amd64   1h

NAME                                             READY     STATUS                  RESTARTS   AGE
po/etcd-kubedev-10-64-33-82                      1/1       Running                 4          2m
po/kube-apiserver-kubedev-10-64-33-82            1/1       Running                 3          2m
po/kube-controller-manager-kubedev-10-64-33-82   1/1       Running                 4          2m
po/kube-scheduler-kubedev-10-64-33-82            1/1       Running                 1          14m
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo kubeadm init --apiserver-advertise-address 10.64.33.82 --apiserver-bind-port 443 --pod-network-cidr 10.244.0.0/16 --kubernetes-version v1.8.4 --skip-preflight-checks
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
[etcd] Wrote Static Pod manifest for a local etcd instance to "/etc/kubernetes/manifests/etcd.yaml"
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests"
[init] This often takes around a minute; or longer if the control plane images have to be pulled.
[apiclient] All control plane components are healthy after 13.550759 seconds
[uploadconfig] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[markmaster] Will mark node kubedev-10-64-33-82 as master by adding a label and a taint
[markmaster] Master kubedev-10-64-33-82 tainted and labelled with key/value: node-role.kubernetes.io/master=""
[bootstraptoken] Using token: fbdadf.ae6a434e584ede62
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

  kubeadm join --token fbdadf.ae6a434e584ede62 10.64.33.82:443 --discovery-token-ca-cert-hash sha256:e2435ed26a2ba502b68108dfdfe6f6d47034c2433fb0b93966451fba4c09d43a

```

```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl get configmap --namespace=kube-system
NAME                                 DATA      AGE
extension-apiserver-authentication   6         6h
kube-flannel-cfg                     2         1h
kube-proxy                           1         5m
kubeadm-config                       1         5m
```

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

```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl get nodes
NAME                  STATUS     ROLES     AGE       VERSION
kubedev-10-64-33-82   NotReady   master    2h        v1.8.4
```

[vagrant@kubedev-10-64-33-82 ~]$ kubectl get pods --all-namespaces
NAMESPACE     NAME                                          READY     STATUS    RESTARTS   AGE
kube-system   etcd-kubedev-10-64-33-82                      1/1       Running   7          2h
kube-system   kube-apiserver-kubedev-10-64-33-82            1/1       Running   3          2h
kube-system   kube-controller-manager-kubedev-10-64-33-82   1/1       Running   3          2h
kube-system   kube-dns-545bc4bfd4-2xk7v                     3/3       Running   0          2h
kube-system   kube-proxy-7fdvp                              1/1       Running   3          2h
kube-system   kube-scheduler-kubedev-10-64-33-82            1/1       Running   4          2h
```

```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl get nodes
NAME                  STATUS    ROLES     AGE       VERSION
kubedev-10-64-33-82   Ready     master    5h        v1.8.4
```

## Docker version

docker-engine v1.11
```
[vagrant@localhost ~]$ sudo yum --enablerepo=dockerrepo\* --showduplicates list docker-engine
Loaded plugins: fastestmirror
dockerrepo-local-http                                                                                        | 2.9 kB  00:00:00     
dockerrepo-local-http/primary_db                                                                             |  34 kB  00:00:00     
Loading mirror speeds from cached hostfile
 * base: mirrors.zju.edu.cn
 * extras: ftp.sjtu.edu.cn
 * updates: mirrors.tuna.tsinghua.edu.cn
Available Packages
docker-engine.x86_64                                  1.7.0-1.el7.centos                                       dockerrepo-local-http
docker-engine.x86_64                                  1.7.1-1.el7.centos                                       dockerrepo-local-http
docker-engine.x86_64                                  1.8.0-1.el7.centos                                       dockerrepo-local-http
docker-engine.x86_64                                  1.8.1-1.el7.centos                                       dockerrepo-local-http
docker-engine.x86_64                                  1.8.2-1.el7.centos                                       dockerrepo-local-http
docker-engine.x86_64                                  1.8.3-1.el7.centos                                       dockerrepo-local-http
docker-engine.x86_64                                  1.9.0-1.el7.centos                                       dockerrepo-local-http
docker-engine.x86_64                                  1.9.1-1.el7.centos                                       dockerrepo-local-http
docker-engine.x86_64                                  1.10.0-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.10.1-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.10.2-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.10.3-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.11.0-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.11.1-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.11.2-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.12.0-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.12.1-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.12.2-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.12.3-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.12.4-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.12.5-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.12.6-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.13.0-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  1.13.1-1.el7.centos                                      dockerrepo-local-http
docker-engine.x86_64                                  17.03.0.ce-1.el7.centos                                  dockerrepo-local-http
docker-engine.x86_64                                  17.03.1.ce-1.el7.centos                                  dockerrepo-local-http
docker-engine.x86_64                                  17.04.0.ce-1.el7.centos                                  dockerrepo-local-http
docker-engine.x86_64                                  17.05.0.ce-1.el7.centos                                  dockerrepo-local-http
```

```
[vagrant@localhost ~]$ sudo yum --enablerepo=dockerrepo\* --showduplicates install docker-engine-1.11.2-1.el7.centos
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.zju.edu.cn
 * extras: ftp.sjtu.edu.cn
 * updates: mirrors.tuna.tsinghua.edu.cn
Resolving Dependencies
--> Running transaction check
---> Package docker-engine.x86_64 0:1.11.2-1.el7.centos will be installed
--> Processing Dependency: docker-engine-selinux >= 1.11.2-1.el7.centos for package: docker-engine-1.11.2-1.el7.centos.x86_64
--> Processing Dependency: libcgroup for package: docker-engine-1.11.2-1.el7.centos.x86_64
--> Processing Dependency: libltdl.so.7()(64bit) for package: docker-engine-1.11.2-1.el7.centos.x86_64
--> Running transaction check
---> Package docker-engine-selinux.noarch 0:17.05.0.ce-1.el7.centos will be installed
--> Processing Dependency: policycoreutils-python for package: docker-engine-selinux-17.05.0.ce-1.el7.centos.noarch
---> Package libcgroup.x86_64 0:0.41-13.el7 will be installed
---> Package libtool-ltdl.x86_64 0:2.4.2-22.el7_3 will be installed
--> Running transaction check
---> Package policycoreutils-python.x86_64 0:2.5-17.1.el7 will be installed
--> Processing Dependency: setools-libs >= 3.3.8-1 for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Processing Dependency: libsemanage-python >= 2.5-5 for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Processing Dependency: audit-libs-python >= 2.1.3-4 for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Processing Dependency: python-IPy for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Processing Dependency: libqpol.so.1(VERS_1.4)(64bit) for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Processing Dependency: libqpol.so.1(VERS_1.2)(64bit) for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Processing Dependency: libapol.so.4(VERS_4.0)(64bit) for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Processing Dependency: checkpolicy for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Processing Dependency: libqpol.so.1()(64bit) for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Processing Dependency: libapol.so.4()(64bit) for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Running transaction check
---> Package audit-libs-python.x86_64 0:2.7.6-3.el7 will be installed
---> Package checkpolicy.x86_64 0:2.5-4.el7 will be installed
---> Package libsemanage-python.x86_64 0:2.5-8.el7 will be installed
---> Package python-IPy.noarch 0:0.75-6.el7 will be installed
---> Package setools-libs.x86_64 0:3.3.8-1.1.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

====================================================================================================================================
 Package                            Arch               Version                              Repository                         Size
====================================================================================================================================
Installing:
 docker-engine                      x86_64             1.11.2-1.el7.centos                  dockerrepo-local-http              13 M
Installing for dependencies:
 audit-libs-python                  x86_64             2.7.6-3.el7                          base                               73 k
 checkpolicy                        x86_64             2.5-4.el7                            base                              290 k
 docker-engine-selinux              noarch             17.05.0.ce-1.el7.centos              dockerrepo-local-http              28 k
 libcgroup                          x86_64             0.41-13.el7                          base                               65 k
 libsemanage-python                 x86_64             2.5-8.el7                            base                              104 k
 libtool-ltdl                       x86_64             2.4.2-22.el7_3                       base                               49 k
 policycoreutils-python             x86_64             2.5-17.1.el7                         base                              446 k
 python-IPy                         noarch             0.75-6.el7                           base                               32 k
 setools-libs                       x86_64             3.3.8-1.1.el7                        base                              612 k

Transaction Summary
====================================================================================================================================
Install  1 Package (+9 Dependent packages)

Total download size: 15 M
Installed size: 18 M
Is this ok [y/d/N]: y
Downloading packages:
(1/10): docker-engine-selinux-17.05.0.ce-1.el7.centos.noarch.rpm                                             |  28 kB  00:00:00     
warning: /var/cache/yum/x86_64/7/base/packages/audit-libs-python-2.7.6-3.el7.x86_64.rpm: Header V3 RSA/SHA256 Signature, key ID f4a80eb5: NOKEY
Public key for audit-libs-python-2.7.6-3.el7.x86_64.rpm is not installed
(2/10): audit-libs-python-2.7.6-3.el7.x86_64.rpm                                                             |  73 kB  00:00:00     
(3/10): docker-engine-1.11.2-1.el7.centos.x86_64.rpm                                                         |  13 MB  00:00:00     
(4/10): libsemanage-python-2.5-8.el7.x86_64.rpm                                                              | 104 kB  00:00:00     
(5/10): libtool-ltdl-2.4.2-22.el7_3.x86_64.rpm                                                               |  49 kB  00:00:00     
(6/10): policycoreutils-python-2.5-17.1.el7.x86_64.rpm                                                       | 446 kB  00:00:00     
(7/10): libcgroup-0.41-13.el7.x86_64.rpm                                                                     |  65 kB  00:00:00     
(8/10): python-IPy-0.75-6.el7.noarch.rpm                                                                     |  32 kB  00:00:00     
(9/10): setools-libs-3.3.8-1.1.el7.x86_64.rpm                                                                | 612 kB  00:00:00     
(10/10): checkpolicy-2.5-4.el7.x86_64.rpm                                                                    | 290 kB  00:00:02     
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                               5.3 MB/s |  15 MB  00:00:02     
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Importing GPG key 0xF4A80EB5:
 Userid     : "CentOS-7 Key (CentOS 7 Official Signing Key) <security@centos.org>"
 Fingerprint: 6341 ab27 53d7 8a78 a7c2 7bb1 24c6 a8a7 f4a8 0eb5
 Package    : centos-release-7-4.1708.el7.centos.x86_64 (@anaconda)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Is this ok [y/N]: y
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : libcgroup-0.41-13.el7.x86_64                                                                                    1/10 
  Installing : setools-libs-3.3.8-1.1.el7.x86_64                                                                               2/10 
  Installing : checkpolicy-2.5-4.el7.x86_64                                                                                    3/10 
  Installing : libtool-ltdl-2.4.2-22.el7_3.x86_64                                                                              4/10 
  Installing : python-IPy-0.75-6.el7.noarch                                                                                    5/10 
  Installing : audit-libs-python-2.7.6-3.el7.x86_64                                                                            6/10 
  Installing : libsemanage-python-2.5-8.el7.x86_64                                                                             7/10 
  Installing : policycoreutils-python-2.5-17.1.el7.x86_64                                                                      8/10 
  Installing : docker-engine-selinux-17.05.0.ce-1.el7.centos.noarch                                                            9/10 
Re-declaration of type docker_t
Failed to create node
Bad type declaration at /etc/selinux/targeted/tmp/modules/400/docker/cil:1
/usr/sbin/semodule:  Failed!
restorecon:  lstat(/var/lib/docker) failed:  No such file or directory
warning: %post(docker-engine-selinux-17.05.0.ce-1.el7.centos.noarch) scriptlet failed, exit status 255
Non-fatal POSTIN scriptlet failure in rpm package docker-engine-selinux-17.05.0.ce-1.el7.centos.noarch
  Installing : docker-engine-1.11.2-1.el7.centos.x86_64                                                                       10/10 
  Verifying  : libsemanage-python-2.5-8.el7.x86_64                                                                             1/10 
  Verifying  : audit-libs-python-2.7.6-3.el7.x86_64                                                                            2/10 
  Verifying  : docker-engine-selinux-17.05.0.ce-1.el7.centos.noarch                                                            3/10 
  Verifying  : python-IPy-0.75-6.el7.noarch                                                                                    4/10 
  Verifying  : libtool-ltdl-2.4.2-22.el7_3.x86_64                                                                              5/10 
  Verifying  : policycoreutils-python-2.5-17.1.el7.x86_64                                                                      6/10 
  Verifying  : libcgroup-0.41-13.el7.x86_64                                                                                    7/10 
  Verifying  : checkpolicy-2.5-4.el7.x86_64                                                                                    8/10 
  Verifying  : docker-engine-1.11.2-1.el7.centos.x86_64                                                                        9/10 
  Verifying  : setools-libs-3.3.8-1.1.el7.x86_64                                                                              10/10 

Installed:
  docker-engine.x86_64 0:1.11.2-1.el7.centos                                                                                        

Dependency Installed:
  audit-libs-python.x86_64 0:2.7.6-3.el7                                    checkpolicy.x86_64 0:2.5-4.el7                         
  docker-engine-selinux.noarch 0:17.05.0.ce-1.el7.centos                    libcgroup.x86_64 0:0.41-13.el7                         
  libsemanage-python.x86_64 0:2.5-8.el7                                     libtool-ltdl.x86_64 0:2.4.2-22.el7_3                   
  policycoreutils-python.x86_64 0:2.5-17.1.el7                              python-IPy.noarch 0:0.75-6.el7                         
  setools-libs.x86_64 0:3.3.8-1.1.el7                                      

Complete!
```

```
[vagrant@localhost ~]$ sudo usermod -aG docker vagrant
```

```
[vagrant@localhost ~]$ exit
logout
Connection to 10.64.33.82 closed.

tangf@DESKTOP-H68OQDV ~
$ ssh -i vagrant vagrant@10.64.33.82
Last login: Sat Nov 25 23:39:14 2017 from 10.64.33.1
```

```
[vagrant@localhost ~]$ docker version
Client:
 Version:      1.11.2
 API version:  1.23
 Go version:   go1.5.4
 Git commit:   b9f10c9
 Built:        Wed Jun  1 21:23:11 2016
 OS/Arch:      linux/amd64

Server:
 Version:      1.11.2
 API version:  1.23
 Go version:   go1.5.4
 Git commit:   b9f10c9
 Built:        Wed Jun  1 21:23:11 2016
 OS/Arch:      linux/amd64
```