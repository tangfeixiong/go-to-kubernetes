# Instruction

Refer to [./README.md](./README.md) about install kubeadm 

kubeadm
```
[vagrant@openshiftdev ~]$ which kubeadm
/bin/kubeadm
```

```
[vagrant@openshiftdev ~]$ kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.2", GitCommit:"bdaeafa71f6c7c04636251031f93464384d54963", GitTreeState:"clean", BuildDate:"2017-10-24T19:38:10Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
```

```
[vagrant@openshiftdev ~]$ kubeadm init --apiserver-bind-port 443 --kubernetes-version v1.6.4
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
this version of kubeadm only supports deploying clusters with the control plane version >= 1.7.0. Current version: v1.6.4
```

## Filewall and Selinux

refer to [./firewalld-and-selinux.md](./firewalld-and-selinux.md)

## K8s v1.7.4

Download server package
```
fanhonglingdeMacBook-Pro:v1.7.4 fanhongling$ curl -jkSL https://dl.k8s.io/v1.7.4/kubernetes-server-linux-amd64.tar.gz -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0    110      0  0:00:01  0:00:01 --:--:--   110
100  417M  100  417M    0     0  3889k      0  0:01:49  0:01:49 --:--:-- 4259k
```

```
[vagrant@openshiftdev ~]$ tar -C /home/vagrant -zxf /vagrant_f/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/v1.7.4/kubernetes-server-linux-amd64.tar.gz 
```

```
[vagrant@openshiftdev ~]$ ls kubernetes/server/bin/
apiextensions-apiserver              kube-aggregator             kube-controller-manager             kube-proxy
cloud-controller-manager             kube-aggregator.docker_tag  kube-controller-manager.docker_tag  kube-proxy.docker_tag
cloud-controller-manager.docker_tag  kube-aggregator.tar         kube-controller-manager.tar         kube-proxy.tar
cloud-controller-manager.tar         kube-apiserver              kubectl                             kube-scheduler
hyperkube                            kube-apiserver.docker_tag   kubefed                             kube-scheduler.docker_tag
kubeadm                              kube-apiserver.tar          kubelet                             kube-scheduler.tar
```

Load docker images
```
[vagrant@openshiftdev ~]$ bindir=./kubernetes/server/bin; for i in $(ls $bindir/*.tar); do docker load -i $i; name=$(basename $i .tar); tag=$(cat $bindir/$name.docker_tag); repo=gcr.io/google_containers/$name; img=$(docker images -q $repo:$tag); [[ -n $img ]] && docker images $repo:$tag || (echo "...failed to load" && docker images $repo); done
REPOSITORY                                          TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/cloud-controller-manager   v1.7.4              a42a0fd85571        3 months ago        125.3 MB
REPOSITORY                                 TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-aggregator   v1.7.4              c1752dd09c50        3 months ago        59.16 MB
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver   v1.7.4              5260ecb5129c        3 months ago        186.1 MB
REPOSITORY                                         TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-controller-manager   v1.7.4              d2adddc4b1cb        3 months ago        138 MB
REPOSITORY                            TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-proxy   v1.7.4              0f3bf654ec61        3 months ago        114.7 MB
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-scheduler   v1.7.4              b1cd468ba656        3 months ago        77.2 MB
```

```
[vagrant@openshiftdev ~]$ docker images gcr.io/google_containers/*
REPOSITORY                                          TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/cloud-controller-manager   v1.7.4              a42a0fd85571        3 months ago        125.3 MB
gcr.io/google_containers/kube-controller-manager    v1.7.4              d2adddc4b1cb        3 months ago        138 MB
gcr.io/google_containers/kube-apiserver             v1.7.4              5260ecb5129c        3 months ago        186.1 MB
gcr.io/google_containers/kube-proxy                 v1.7.4              0f3bf654ec61        3 months ago        114.7 MB
gcr.io/google_containers/kube-aggregator            v1.7.4              c1752dd09c50        3 months ago        59.16 MB
gcr.io/google_containers/kube-scheduler             v1.7.4              b1cd468ba656        3 months ago        77.2 MB
```

### Offline install

Refer to [./using-yum-mirror](./using-yum-mirror) to enable local repository

### Issue

#### docker 1.10

如果docker版本时1.10.x, 需使用命令选项 --skip-preflight-checks
```
[vagrant@openshiftdev ~]$ sudo kubeadm init --apiserver-bind-port 443 --kubernetes-version v1.7.4
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.7.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Running pre-flight checks
[preflight] The system verification failed. Printing the output from the verification:
OS: Linux
KERNEL_VERSION: 3.10.0-229.el7.x86_64
CONFIG_NAMESPACES: enabled
CONFIG_NET_NS: enabled
CONFIG_PID_NS: enabled
CONFIG_IPC_NS: enabled
CONFIG_UTS_NS: enabled
CONFIG_CGROUPS: enabled
CONFIG_CGROUP_CPUACCT: enabled
CONFIG_CGROUP_DEVICE: enabled
CONFIG_CGROUP_FREEZER: enabled
CONFIG_CGROUP_SCHED: enabled
CONFIG_CPUSETS: enabled
CONFIG_MEMCG: enabled
CONFIG_INET: enabled
CONFIG_EXT4_FS: enabled (as module)
CONFIG_PROC_FS: enabled
CONFIG_NETFILTER_XT_TARGET_REDIRECT: enabled (as module)
CONFIG_NETFILTER_XT_MATCH_COMMENT: enabled (as module)
CONFIG_OVERLAY_FS: enabled (as module)
CONFIG_AUFS_FS: not set - Required for aufs.
CONFIG_BLK_DEV_DM: enabled (as module)
CGROUPS_CPU: enabled
CGROUPS_CPUACCT: enabled
CGROUPS_CPUSET: enabled
CGROUPS_DEVICES: enabled
CGROUPS_FREEZER: enabled
CGROUPS_MEMORY: enabled
DOCKER_VERSION: 1.10.3
[preflight] WARNING: kubelet service is not enabled, please run 'systemctl enable kubelet.service'
[preflight] WARNING: Running with swap on is not supported. Please disable swap or set kubelet's --fail-swap-on flag to false.
[preflight] Some fatal errors occurred:
	unsupported docker version: 1.10.3
	/proc/sys/net/bridge/bridge-nf-call-iptables contents are not set to 1
[preflight] If you know what you are doing, you can skip pre-flight checks with `--skip-preflight-checks`
```

#### swap

Swap is not supported
```
[vagrant@openshiftdev ~]$ sudo kubeadm init --apiserver-bind-port 443 --kubernetes-version v1.7.4 --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.7.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [openshiftdev.local kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.0.2.15]
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
		- gcr.io/google_containers/kube-apiserver-amd64:v1.7.4
		- gcr.io/google_containers/kube-controller-manager-amd64:v1.7.4
		- gcr.io/google_containers/kube-scheduler-amd64:v1.7.4

You can troubleshoot this for example with the following commands if you're on a systemd-powered system:
	- 'systemctl status kubelet'
	- 'journalctl -xeu kubelet'
couldn't initialize a Kubernetes cluster
```

```
[vagrant@openshiftdev ~]$ cat /etc/systemd/system/kubelet.service
[Unit]
Description=kubelet: The Kubernetes Node Agent
Documentation=http://kubernetes.io/docs/

[Service]
ExecStart=/usr/bin/kubelet
Restart=always
StartLimitInterval=0
RestartSec=10

[Install]
WantedBy=multi-user.target
```

```
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
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS
```

```
[vagrant@openshiftdev ~]$ ls /etc/kubernetes/
admin.conf  controller-manager.conf  kubelet.conf  manifests  pki  scheduler.conf
```

```
[vagrant@openshiftdev ~]$ sudo systemctl start kubelet.service
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: activating (auto-restart) (Result: exit-code) since Sun 2017-11-19 20:44:47 UTC; 6s ago
     Docs: http://kubernetes.io/docs/
  Process: 7102 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS (code=exited, status=1/FAILURE)
 Main PID: 7102 (code=exited, status=1/FAILURE)

Nov 19 20:44:47 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 19 20:44:47 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
Nov 19 20:44:47 openshiftdev.local systemd[1]: kubelet.service failed.
```

```
[vagrant@openshiftdev ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service
-- Logs begin at Sat 2017-11-18 19:57:25 UTC, end at Sun 2017-11-19 20:53:17 UTC. --
Nov 19 20:46:13 openshiftdev.local systemd[1]: kubelet.service failed.
Nov 19 20:46:24 openshiftdev.local systemd[1]: kubelet.service holdoff time over, scheduling restart.
Nov 19 20:46:24 openshiftdev.local systemd[1]: Started kubelet: The Kubernetes Node Agent.
Nov 19 20:46:24 openshiftdev.local systemd[1]: Starting kubelet: The Kubernetes Node Agent...
### snippets ###
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: error: failed to run Kubelet: Running with swap on is not supported, please disable swap! or set --fail-swap-on flag to false. /proc/swaps contained: [Filename                                Type                Size        Used        Priority /dev/dm-1                               partition        1048572        1988        -1]
Nov 19 20:46:24 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 19 20:46:24 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
### snippets ###
```

```
[vagrant@openshiftdev ~]$ sudo sed -i 's/\(ExecStart=.*ARGS\)/\1 --fail-swap-on=false/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

```
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
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false
```

#### cgroup

CentOS Cgroup settings other than Ubuntu
```
[vagrant@openshiftdev ~]$ sudo systemctl daemon-reload 
[vagrant@openshiftdev ~]$ sudo systemctl restart kubelet.service
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: activating (auto-restart) (Result: exit-code) since Sun 2017-11-19 21:08:57 UTC; 6s ago
     Docs: http://kubernetes.io/docs/
  Process: 10596 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false (code=exited, status=1/FAILURE)
 Main PID: 10596 (code=exited, status=1/FAILURE)

Nov 19 21:08:57 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 19 21:08:57 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
Nov 19 21:08:57 openshiftdev.local systemd[1]: kubelet.service failed.
```

[vagrant@openshiftdev ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service
```
### snippets ###
Nov 19 21:09:29 openshiftdev.local systemd[1]: kubelet.service holdoff time over, scheduling restart.
Nov 19 21:09:29 openshiftdev.local systemd[1]: Started kubelet: The Kubernetes Node Agent.
Nov 19 21:09:29 openshiftdev.local systemd[1]: Starting kubelet: The Kubernetes Node Agent...
### snippets ###
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.650250   10689 docker_service.go:216] No cgroup driver is set in Docker
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.650322   10689 docker_service.go:217] Falling back to use the default driver: "cgroupfs"
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: error: failed to run Kubelet: failed to create kubelet: misconfiguration: kubelet cgroup driver: "systemd" is different from docker cgroup driver: "cgroupfs"
Nov 19 21:09:30 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 19 21:09:30 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
Nov 19 21:09:30 openshiftdev.local systemd[1]: kubelet.service failed.
```

```
[vagrant@openshiftdev ~]$ sudo sed -i 's/\(Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"\)/#\1\nEnvironment="KUBELET_CGROUP_ARGS=--cgroup-driver=cgroupfs"/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

#### Default networking is on eth0, but it is vagrant console with VirtualBox NAT mode

API Endpoint unexpected
```
[vagrant@openshiftdev ~]$ sudo systemctl daemon-reload 
[vagrant@openshiftdev ~]$ sudo systemctl restart kubelet.service
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Sun 2017-11-19 21:16:27 UTC; 3s ago
     Docs: http://kubernetes.io/docs/
 Main PID: 11932 (kubelet)
   Memory: 25.0M
   CGroup: /system.slice/kubelet.service
           └─11932 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin --cluster-dns=10.96.0.10 --cluster-domain=cluster.local --authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt --cadvisor-port=0 --cgroup-driver=cgroupfs --rotate-certificates=true --cert-dir=/var/lib/kubelet/pki --fail-swap-on=false

Nov 19 21:16:29 openshiftdev.local kubelet[11932]: E1119 21:16:29.910453   11932 kubelet_node_status.go:107] Unable to register node "openshiftdev.local" with API server: Post https://10.0.2.15:443/api/v1/nodes: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: E1119 21:16:30.382097   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: E1119 21:16:30.405038   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: E1119 21:16:30.417405   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.0.2.15:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: I1119 21:16:30.710880   11932 kubelet_node_status.go:280] Setting node annotation to enable volume controller attach/detach
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: I1119 21:16:30.811518   11932 kubelet_node_status.go:83] Attempting to register node openshiftdev.local
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: E1119 21:16:30.811775   11932 kubelet_node_status.go:107] Unable to register node "openshiftdev.local" with API server: Post https://10.0.2.15:443/api/v1/nodes: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:31 openshiftdev.local kubelet[11932]: E1119 21:16:31.383195   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:31 openshiftdev.local kubelet[11932]: E1119 21:16:31.405884   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:31 openshiftdev.local kubelet[11932]: E1119 21:16:31.418781   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.0.2.15:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
```

```
[vagrant@openshiftdev ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service | grep E1119 | tail
Nov 19 21:30:00 openshiftdev.local kubelet[11932]: E1119 21:30:00.616706   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:00 openshiftdev.local kubelet[11932]: E1119 21:30:00.619308   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:01 openshiftdev.local kubelet[11932]: E1119 21:30:01.616746   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.0.2.15:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:01 openshiftdev.local kubelet[11932]: E1119 21:30:01.623751   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:01 openshiftdev.local kubelet[11932]: E1119 21:30:01.625465   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:02 openshiftdev.local kubelet[11932]: E1119 21:30:02.624423   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.0.2.15:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:02 openshiftdev.local kubelet[11932]: E1119 21:30:02.625036   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:02 openshiftdev.local kubelet[11932]: E1119 21:30:02.628671   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:02 openshiftdev.local kubelet[11932]: E1119 21:30:02.951025   11932 kubelet.go:1612] Failed creating a mirror pod for "kube-controller-manager-openshiftdev.local_kube-system(68a0ca8ea0a05b769d4c2c3355e47ae5)": Post https://10.0.2.15:443/api/v1/namespaces/kube-system/pods: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:03 openshiftdev.local kubelet[11932]: E1119 21:30:03.146769   11932 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
```

Remove directory /etc/kubernetes
```
[vagrant@openshiftdev ~]$ sudo rm -rf /etc/kubernetes/*
```

```
[vagrant@openshiftdev ~]$ sudo kubeadm init --apiserver-advertise-address 10.64.33.82 --apiserver-bind-port 443 --kubernetes-version v1.7.4 --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.7.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [openshiftdev.local kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.64.33.82]
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
		- gcr.io/google_containers/kube-apiserver-amd64:v1.7.4
		- gcr.io/google_containers/kube-controller-manager-amd64:v1.7.4
		- gcr.io/google_containers/kube-scheduler-amd64:v1.7.4

You can troubleshoot this for example with the following commands if you're on a systemd-powered system:
	- 'systemctl status kubelet'
	- 'journalctl -xeu kubelet'
couldn't initialize a Kubernetes cluster
```

```
[vagrant@openshiftdev ~]$ sudo systemctl restart kubelet.service
```

### Docker is old

cgroup-parent for systemd cgroup should be a valid slice named as "xxx.slice"
```
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Tue 2017-11-21 21:59:58 UTC; 7min ago
     Docs: http://kubernetes.io/docs/
 Main PID: 6842 (kubelet)
   Memory: 31.5M
   CGroup: /system.slice/kubelet.service
           └─6842 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin --cluster-dns=10.96.0.10 --cluster-domain=cluster.local --authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt --cadvisor-port=0 --cgroup-driver=cgroupfs --rotate-certificates=true --cert-dir=/var/lib/kubelet/pki --fail-swap-on=false

Nov 21 22:07:49 openshiftdev.local kubelet[6842]: E1121 22:07:49.325245    6842 kubelet.go:1612] Failed creating a mirror pod for "kube-controller-manager-openshiftdev.local_kube-system(7206a0a1a117d27a758a8708ae01ada2)": Post https://10.64.33.82:443/api/v1/namespaces/kube-system/pods: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 21 22:07:49 openshiftdev.local kubelet[6842]: E1121 22:07:49.367573    6842 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.64.33.82:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 21 22:07:49 openshiftdev.local kubelet[6842]: E1121 22:07:49.369698    6842 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.64.33.82:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 21 22:07:49 openshiftdev.local kubelet[6842]: E1121 22:07:49.370758    6842 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.64.33.82:443/api/v1/services?resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 21 22:07:49 openshiftdev.local kubelet[6842]: E1121 22:07:49.631419    6842 remote_runtime.go:92] RunPodSandbox from runtime service failed: rpc error: code = Unknown desc = failed to create a sandbox for pod "kube-controller-manager-openshiftdev.local": Error response from daemon: cgroup-parent for systemd cgroup should be a valid slice named as "xxx.slice"
Nov 21 22:07:49 openshiftdev.local kubelet[6842]: E1121 22:07:49.631419    6842 kuberuntime_sandbox.go:54] CreatePodSandbox for pod "kube-controller-manager-openshiftdev.local_kube-system(7206a0a1a117d27a758a8708ae01ada2)" failed: rpc error: code = Unknown desc = failed to create a sandbox for pod "kube-controller-manager-openshiftdev.local": Error response from daemon: cgroup-parent for systemd cgroup should be a valid slice named as "xxx.slice"
Nov 21 22:07:49 openshiftdev.local kubelet[6842]: E1121 22:07:49.631419    6842 kuberuntime_manager.go:632] createPodSandbox for pod "kube-controller-manager-openshiftdev.local_kube-system(7206a0a1a117d27a758a8708ae01ada2)" failed: rpc error: code = Unknown desc = failed to create a sandbox for pod "kube-controller-manager-openshiftdev.local": Error response from daemon: cgroup-parent for systemd cgroup should be a valid slice named as "xxx.slice"
Nov 21 22:07:49 openshiftdev.local kubelet[6842]: E1121 22:07:49.631419    6842 pod_workers.go:182] Error syncing pod 7206a0a1a117d27a758a8708ae01ada2 ("kube-controller-manager-openshiftdev.local_kube-system(7206a0a1a117d27a758a8708ae01ada2)"), skipping: failed to "CreatePodSandbox" for "kube-controller-manager-openshiftdev.local_kube-system(7206a0a1a117d27a758a8708ae01ada2)" with CreatePodSandboxError: "CreatePodSandbox for pod \"kube-controller-manager-openshiftdev.local_kube-system(7206a0a1a117d27a758a8708ae01ada2)\" failed: rpc error: code = Unknown desc = failed to create a sandbox for pod \"kube-controller-manager-openshiftdev.local\": Error response from daemon: cgroup-parent for systemd cgroup should be a valid slice named as \"xxx.slice\""
Nov 21 22:07:49 openshiftdev.local kubelet[6842]: E1121 22:07:49.703584    6842 eviction_manager.go:238] eviction manager: unexpected err: failed to get node info: node 'openshiftdev.local' not found
Nov 21 22:07:50 openshiftdev.local kubelet[6842]: E1121 22:07:50.003588    6842 event.go:209] Unable to write event: 'Post https://10.64.33.82:443/api/v1/namespaces/default/events: dial tcp 10.64.33.82:443: getsockopt: connection refused' (may retry after sleeping)
```

Refer to https://github.com/kubernetes/kubernetes/issues/45555

### Docker 1.12 cgroup driver is using systemd

Upgrade docker to v1.12, refer to [./upgrade-docker-engine.md](./upgrade-docker-engine.md)
```
[vagrant@openshiftdev ~]$ sudo systemctl restart kubelet.service
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: activating (auto-restart) (Result: exit-code) since Tue 2017-11-21 23:46:51 UTC; 3s ago
     Docs: http://kubernetes.io/docs/
  Process: 22189 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false (code=exited, status=1/FAILURE)
 Main PID: 22189 (code=exited, status=1/FAILURE)

Nov 21 23:46:51 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 21 23:46:51 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
Nov 21 23:46:51 openshiftdev.local systemd[1]: kubelet.service failed.
```

```
[vagrant@openshiftdev ~]$ sudo journalctl --no-pager --no-tail --pager-end --follow --unit kubelet.service
### snippets ###
Nov 21 23:47:53 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 21 23:47:53 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
Nov 21 23:47:53 openshiftdev.local systemd[1]: kubelet.service failed.
Nov 21 23:47:53 openshiftdev.local kubelet[22315]: error: failed to run Kubelet: failed to create kubelet: misconfiguration: kubelet cgroup driver: "cgroupfs" is different from docker cgroup driver: "systemd"
### snippets ###
```

Change back to 'Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"' in /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
```
[vagrant@openshiftdev ~]$ sudo systemctl daemon-reload
[vagrant@openshiftdev ~]$ sudo systemctl restart kubelet.service
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Tue 2017-11-21 23:48:14 UTC; 56s ago
     Docs: http://kubernetes.io/docs/
 Main PID: 22380 (kubelet)
   Memory: 2.9M
   CGroup: /system.slice/kubelet.service
           └─22380 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin --cluster-dns=10.96.0.10 --cluster-domain=cluster.local --authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt --cadvisor-port=0 --cgroup-driver=systemd --rotate-certificates=true --cert-dir=/var/lib/kubelet/pki --fail-swap-on=false

Nov 21 23:49:09 openshiftdev.local kubelet[22380]: E1121 23:49:09.442823   22380 remote_image.go:108] PullImage "gcr.io/google_containers/kube-scheduler-amd64:v1.7.4" from image service failed: rpc error: code = Unknown desc = Get https://gcr.io/v1/_ping: dial tcp 108.177.97.82:443: i/o timeout
Nov 21 23:49:09 openshiftdev.local kubelet[22380]: E1121 23:49:09.442887   22380 kuberuntime_image.go:50] Pull image "gcr.io/google_containers/kube-scheduler-amd64:v1.7.4" failed: rpc error: code = Unknown desc = Get https://gcr.io/v1/_ping: dial tcp 108.177.97.82:443: i/o timeout
Nov 21 23:49:09 openshiftdev.local kubelet[22380]: E1121 23:49:09.443007   22380 kuberuntime_manager.go:714] container start failed: ErrImagePull: rpc error: code = Unknown desc = Get https://gcr.io/v1/_ping: dial tcp 108.177.97.82:443: i/o timeout
Nov 21 23:49:09 openshiftdev.local kubelet[22380]: E1121 23:49:09.443072   22380 pod_workers.go:182] Error syncing pod 5a10d702e0d21ab01a0132a2e233b601 ("kube-scheduler-openshiftdev.local_kube-system(5a10d702e0d21ab01a0132a2e233b601)"), skipping: failed to "StartContainer" for "kube-scheduler" with ErrImagePull: "rpc error: code = Unknown desc = Get https://gcr.io/v1/_ping: dial tcp 108.177.97.82:443: i/o timeout"
Nov 21 23:49:10 openshiftdev.local kubelet[22380]: I1121 23:49:10.250177   22380 kubelet_node_status.go:280] Setting node annotation to enable volume controller attach/detach
Nov 21 23:49:10 openshiftdev.local kubelet[22380]: W1121 23:49:10.254979   22380 status_manager.go:431] Failed to get status for pod "kube-scheduler-openshiftdev.local_kube-system(5a10d702e0d21ab01a0132a2e233b601)": Get https://10.64.33.82:443/api/v1/namespaces/kube-system/pods/kube-scheduler-openshiftdev.local: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 21 23:49:10 openshiftdev.local kubelet[22380]: E1121 23:49:10.255059   22380 kubelet.go:1612] Failed creating a mirror pod for "kube-scheduler-openshiftdev.local_kube-system(5a10d702e0d21ab01a0132a2e233b601)": Post https://10.64.33.82:443/api/v1/namespaces/kube-system/pods: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 21 23:49:10 openshiftdev.local kubelet[22380]: E1121 23:49:10.307679   22380 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.64.33.82:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 21 23:49:10 openshiftdev.local kubelet[22380]: E1121 23:49:10.309454   22380 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.64.33.82:443/api/v1/services?resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 21 23:49:10 openshiftdev.local kubelet[22380]: E1121 23:49:10.313140   22380 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.64.33.82:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
```

### Incorrect kube-system images name

Pods are running, but others not started

```
[vagrant@openshiftdev ~]$ docker ps -f label=io.kubernetes.pod.namespace=kube-system
CONTAINER ID        IMAGE                                      COMMAND                  CREATED              STATUS              PORTS               NAMES
499a3c6d0dc0        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 19 minutes ago       Up 19 minutes                           k8s_POD_etcd-openshiftdev.local_kube-system_d76e26fba3bf2bfd215eb29011d55250_0
6ca17cb266c1        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 19 minutes ago       Up 19 minutes                           k8s_POD_kube-controller-manager-openshiftdev.local_kube-system_7206a0a1a117d27a758a8708ae01ada2_0
c4f3d3ba418c        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 19 minutes ago       Up 19 minutes                           k8s_POD_kube-apiserver-openshiftdev.local_kube-system_c0ff30e7eb32f84064235a4d57912227_0
b75fb0ab932f        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 19 minutes ago       Up 19 minutes                           k8s_POD_kube-scheduler-openshiftdev.local_kube-system_5a10d702e0d21ab01a0132a2e233b601_0
```

```
[vagrant@openshiftdev ~]$ docker tag gcr.io/google_containers/kube-apiserver:v1.7.4 gcr.io/google_containers/kube-apiserver-amd64:v1.7.4
[vagrant@openshiftdev ~]$ docker tag gcr.io/google_containers/kube-controller-manager:v1.7.4 gcr.io/google_containers/kube-controller-manager-amd64:v1.7.4
[vagrant@openshiftdev ~]$ docker tag gcr.io/google_containers/kube-scheduler:v1.7.4 gcr.io/google_containers/kube-scheduler-amd64:v1.7.4
[vagrant@openshiftdev ~]$ docker tag gcr.io/google_containers/kube-proxy:v1.7.4 gcr.io/google_containers/kube-proxy-amd64:v1.7.4
```

```
[vagrant@openshiftdev ~]$ docker ps -f label=io.kubernetes.pod.namespace=kube-system
CONTAINER ID        IMAGE                                      COMMAND                  CREATED              STATUS              PORTS               NAMES
6c6d51eed65e        d2adddc4b1cb                               "kube-controller-mana"   18 seconds ago       Up 17 seconds                           k8s_kube-controller-manager_kube-controller-manager-openshiftdev.local_kube-system_7206a0a1a117d27a758a8708ae01ada2_0
f024aefb9e64        5260ecb5129c                               "kube-apiserver --kub"   49 seconds ago       Up 48 seconds                           k8s_kube-apiserver_kube-apiserver-openshiftdev.local_kube-system_c0ff30e7eb32f84064235a4d57912227_0
818522f49c5b        b1cd468ba656                               "kube-scheduler --lea"   About a minute ago   Up About a minute                       k8s_kube-scheduler_kube-scheduler-openshiftdev.local_kube-system_5a10d702e0d21ab01a0132a2e233b601_0
499a3c6d0dc0        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 19 minutes ago       Up 19 minutes                           k8s_POD_etcd-openshiftdev.local_kube-system_d76e26fba3bf2bfd215eb29011d55250_0
6ca17cb266c1        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 19 minutes ago       Up 19 minutes                           k8s_POD_kube-controller-manager-openshiftdev.local_kube-system_7206a0a1a117d27a758a8708ae01ada2_0
c4f3d3ba418c        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 19 minutes ago       Up 19 minutes                           k8s_POD_kube-apiserver-openshiftdev.local_kube-system_c0ff30e7eb32f84064235a4d57912227_0
b75fb0ab932f        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 19 minutes ago       Up 19 minutes                           k8s_POD_kube-scheduler-openshiftdev.local_kube-system_5a10d702e0d21ab01a0132a2e233b601_0
```

Missing etcd image
```
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Tue 2017-11-21 23:48:14 UTC; 20min ago
     Docs: http://kubernetes.io/docs/
 Main PID: 22380 (kubelet)
   Memory: 11.4M
   CGroup: /system.slice/kubelet.service
           └─22380 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin --cluster-dns=10.96.0.10 --cluster-domain=cluster.local --authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt --cadvisor-port=0 --cgroup-driver=systemd --rotate-certificates=true --cert-dir=/var/lib/kubelet/pki --fail-swap-on=false

Nov 22 00:09:03 openshiftdev.local kubelet[22380]: E1122 00:09:03.335387   22380 kubelet.go:1612] Failed creating a mirror pod for "etcd-openshiftdev.local_kube-system(d76e26fba3bf2bfd215eb29011d55250)": Post https://10.64.33.82:443/api/v1/namespaces/kube-system/pods: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 22 00:09:03 openshiftdev.local kubelet[22380]: E1122 00:09:03.625608   22380 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.64.33.82:443/api/v1/services?resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 22 00:09:03 openshiftdev.local kubelet[22380]: E1122 00:09:03.628170   22380 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.64.33.82:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 22 00:09:03 openshiftdev.local kubelet[22380]: I1122 00:09:03.635737   22380 kuberuntime_manager.go:499] Container {Name:etcd Image:gcr.io/google_containers/etcd-amd64:3.0.17 Command:[etcd --listen-client-urls=http://127.0.0.1:2379 --advertise-client-urls=http://127.0.0.1:2379 --data-dir=/var/lib/etcd] Args:[] WorkingDir: Ports:[] EnvFrom:[] Env:[] Resources:{Limits:map[] Requests:map[]} VolumeMounts:[{Name:etcd ReadOnly:false MountPath:/var/lib/etcd SubPath: MountPropagation:<nil>}] LivenessProbe:&Probe{Handler:Handler{Exec:nil,HTTPGet:&HTTPGetAction{Path:/health,Port:2379,Host:127.0.0.1,Scheme:HTTP,HTTPHeaders:[],},TCPSocket:nil,},InitialDelaySeconds:15,TimeoutSeconds:15,PeriodSeconds:10,SuccessThreshold:1,FailureThreshold:8,} ReadinessProbe:nil Lifecycle:nil TerminationMessagePath:/dev/termination-log TerminationMessagePolicy:File ImagePullPolicy:IfNotPresent SecurityContext:nil Stdin:false StdinOnce:false TTY:false} is dead, but RestartPolicy says that we should restart it.
Nov 22 00:09:03 openshiftdev.local kubelet[22380]: E1122 00:09:03.638685   22380 pod_workers.go:182] Error syncing pod d76e26fba3bf2bfd215eb29011d55250 ("etcd-openshiftdev.local_kube-system(d76e26fba3bf2bfd215eb29011d55250)"), skipping: failed to "StartContainer" for "etcd" with ImagePullBackOff: "Back-off pulling image \"gcr.io/google_containers/etcd-amd64:3.0.17\""
Nov 22 00:09:03 openshiftdev.local kubelet[22380]: E1122 00:09:03.641658   22380 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.64.33.82:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 22 00:09:04 openshiftdev.local kubelet[22380]: E1122 00:09:04.521558   22380 eviction_manager.go:238] eviction manager: unexpected err: failed to get node info: node 'openshiftdev.local' not found
Nov 22 00:09:04 openshiftdev.local kubelet[22380]: E1122 00:09:04.626285   22380 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.64.33.82:443/api/v1/services?resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 22 00:09:04 openshiftdev.local kubelet[22380]: E1122 00:09:04.628737   22380 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.64.33.82:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
Nov 22 00:09:04 openshiftdev.local kubelet[22380]: E1122 00:09:04.642894   22380 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.64.33.82:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.64.33.82:443: getsockopt: connection refused
```

Refer to [../kubernetes-etcd-v3-image.md](../kubernetes-etcd-v3-image.md)
```
[vagrant@openshiftdev ~]$ docker tag gcr.io/google_containers/etcd:3.0.17 gcr.io/google_containers/etcd-amd64:3.0.17
```

```
[vagrant@openshiftdev ~]$ docker ps -f label=io.kubernetes.pod.namespace=kube-system
CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
e861f771ad5a        243830dae7dd                               "etcd --listen-client"   8 seconds ago       Up 6 seconds                            k8s_etcd_etcd-openshiftdev.local_kube-system_d76e26fba3bf2bfd215eb29011d55250_0
6c6d51eed65e        d2adddc4b1cb                               "kube-controller-mana"   13 minutes ago      Up 12 minutes                           k8s_kube-controller-manager_kube-controller-manager-openshiftdev.local_kube-system_7206a0a1a117d27a758a8708ae01ada2_0
818522f49c5b        b1cd468ba656                               "kube-scheduler --lea"   13 minutes ago      Up 13 minutes                           k8s_kube-scheduler_kube-scheduler-openshiftdev.local_kube-system_5a10d702e0d21ab01a0132a2e233b601_0
499a3c6d0dc0        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 32 minutes ago      Up 32 minutes                           k8s_POD_etcd-openshiftdev.local_kube-system_d76e26fba3bf2bfd215eb29011d55250_0
6ca17cb266c1        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 32 minutes ago      Up 32 minutes                           k8s_POD_kube-controller-manager-openshiftdev.local_kube-system_7206a0a1a117d27a758a8708ae01ada2_0
c4f3d3ba418c        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 32 minutes ago      Up 32 minutes                           k8s_POD_kube-apiserver-openshiftdev.local_kube-system_c0ff30e7eb32f84064235a4d57912227_0
b75fb0ab932f        gcr.io/google_containers/pause-amd64:3.0   "/pause"                 32 minutes ago      Up 32 minutes                           k8s_POD_kube-scheduler-openshiftdev.local_kube-system_5a10d702e0d21ab01a0132a2e233b601_0
```

```
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Tue 2017-11-21 23:48:14 UTC; 41min ago
     Docs: http://kubernetes.io/docs/
 Main PID: 22380 (kubelet)
   Memory: 16.1M
   CGroup: /system.slice/kubelet.service
           └─22380 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin --cluster-dns=10.96.0.10 --cluster-domain=cluster.local --authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt --cadvisor-port=0 --cgroup-driver=systemd --rotate-certificates=true --cert-dir=/var/lib/kubelet/pki --fail-swap-on=false

Nov 22 00:29:15 openshiftdev.local kubelet[22380]: W1122 00:29:15.455041   22380 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 22 00:29:15 openshiftdev.local kubelet[22380]: E1122 00:29:15.455432   22380 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Nov 22 00:29:20 openshiftdev.local kubelet[22380]: W1122 00:29:20.456078   22380 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 22 00:29:20 openshiftdev.local kubelet[22380]: E1122 00:29:20.456109   22380 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Nov 22 00:29:25 openshiftdev.local kubelet[22380]: W1122 00:29:25.457840   22380 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 22 00:29:25 openshiftdev.local kubelet[22380]: E1122 00:29:25.457976   22380 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Nov 22 00:29:30 openshiftdev.local kubelet[22380]: W1122 00:29:30.459087   22380 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 22 00:29:30 openshiftdev.local kubelet[22380]: E1122 00:29:30.459175   22380 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
Nov 22 00:29:35 openshiftdev.local kubelet[22380]: W1122 00:29:35.459985   22380 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 22 00:29:35 openshiftdev.local kubelet[22380]: E1122 00:29:35.460086   22380 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
```

### Client

Cli
```
[vagrant@openshiftdev ~]$ kubectl version --client
Client Version: version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.2", GitCommit:"bdaeafa71f6c7c04636251031f93464384d54963", GitTreeState:"clean", BuildDate:"2017-10-24T19:48:57Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
```

```
[vagrant@openshiftdev ~]$ export KUBECONFIG=/etc/kubernetes/admin.conf
[vagrant@openshiftdev ~]$ kubectl version
Client Version: version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.2", GitCommit:"bdaeafa71f6c7c04636251031f93464384d54963", GitTreeState:"clean", BuildDate:"2017-10-24T19:48:57Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
error: Error loading config file "/etc/kubernetes/admin.conf": open /etc/kubernetes/admin.conf: permission denied
```

```
[vagrant@openshiftdev ~]$ sudo chmod o+r /etc/kubernetes/admin.conf
```

```
[vagrant@openshiftdev ~]$ kubectl version
Client Version: version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.2", GitCommit:"bdaeafa71f6c7c04636251031f93464384d54963", GitTreeState:"clean", BuildDate:"2017-10-24T19:48:57Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"7", GitVersion:"v1.7.4", GitCommit:"793658f2d7ca7f064d2bdf606519f9fe1229c381", GitTreeState:"clean", BuildDate:"2017-08-17T08:30:51Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
```

```
[vagrant@openshiftdev ~]$ kubectl get pods --namespace=kube-system
NAME                                         READY     STATUS    RESTARTS   AGE
etcd-openshiftdev.local                      1/1       Running   0          5m
kube-apiserver-openshiftdev.local            1/1       Running   8          5m
kube-controller-manager-openshiftdev.local   1/1       Running   1          5m
kube-scheduler-openshiftdev.local            1/1       Running   1          4m
```

Node not ready, because of CNI net
```
[vagrant@openshiftdev ~]$ kubectl get nodes
NAME                 STATUS     ROLES     AGE       VERSION
openshiftdev.local   NotReady   <none>    5m        v1.8.2
```

