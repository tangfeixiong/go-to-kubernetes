# instruction

Upgrade from previous v1.8.2 (CentOS 7.2)
```
[vagrant@openshiftdev ~]$ sudo kubeadm upgrade apply v1.8.4
[preflight] Running pre-flight checks
[upgrade] Making sure the cluster is healthy:
[upgrade/health] Checking API Server health: Healthy
[upgrade/health] Checking Node health: More than one Node unhealthy
[upgrade/health] FATAL: The cluster is not in an upgradeable state due to: there are NotReady Nodes in the cluster: [kubedev-10-64-33-82]
```

Re-install
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo rm -rf /etc/kubernetes/*
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
