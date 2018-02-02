# Development Guide

## Prerequisites

https://github.com/kubernetes-incubator/external-storage/tree/master/local-volume#option-3-baremetal-environments

Follow [../docs/k8s-v1.9.0-deploy](../docs/k8s-v1.9.0-deploy)

__Master: Fedora26__

1. Update kubelet.service

From
```
[vagrant@kubedev-172-17-4-59 ~]$ cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
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

To
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo sed -i 's/\(^ExecStart=$\)/Environment="KUBELET_FEATURE_GATES=--feature-gates=PersistentLocalVolumes=true,VolumeScheduling=true,MountPropagation=true"\n\1/;s/\(^ExecStart=.\+$\)/\1 $KUBELET_FEATURE_GATES/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf && cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"
Environment="KUBELET_CERTIFICATE_ARGS=--rotate-certificates=true --cert-dir=/var/lib/kubelet/pki"
Environment="KUBELET_FEATURE_GATES=--feature-gates=PersistentLocalVolumes=true,VolumeScheduling=true,MountPropagation=true"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS $KUBELET_FEATURE_GATES
```

1. Update kube-apiserver.yaml

bash
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo sed -i 's/^    - kube-apiserver$/&\n    - --feature-gates=PersistentLocalVolumes=true,VolumeScheduling=true,MountPropagation=true/' /etc/kubernetes/manifests/kube-apiserver.yaml    
```
1. Update kube-controller-manager.yaml

bash
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo sed -i 's/^    - kube-controller-manager$/&\n    - --feature-gates=PersistentLocalVolumes=true,VolumeScheduling=true,MountPropagation=true/' /etc/kubernetes/manifests/kube-controller-manager.yaml
```

1. Update kube-scheduler.yaml

bash
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ sudo sed -i 's/^    - kube-scheduler$/&\n    - --feature-gates=PersistentLocalVolumes=true,VolumeScheduling=true,MountPropagation=true/' /etc/kubernetes/manifests/kube-scheduler.yaml
```

__Worker: Fedora27__

1. Update kubelet.service

bash
```
[vagrant@rookdev-172-17-4-63 ~]$ sudo sed -i 's/\(^ExecStart=$\)/Environment="KUBELET_FEATURE_GATES=--feature-gates=PersistentLocalVolumes=true,VolumeScheduling=true,MountPropagation=true"\n\1/;s/\(^ExecStart=.\+$\)/\1 $KUBELET_FEATURE_GATES/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf && cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"
Environment="KUBELET_CERTIFICATE_ARGS=--rotate-certificates=true --cert-dir=/var/lib/kubelet/pki"
Environment="KUBELET_FEATURE_GATES=--feature-gates=PersistentLocalVolumes=true,VolumeScheduling=true,MountPropagation=true"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS $KUBELET_FEATURE_GATES
```

__Worker: Ubuntu17.04__

1. Update kubelet.service

bash
```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ sudo sed -i 's/\(^ExecStart=$\)/Environment="KUBELET_FEATURE_GATES=--feature-gates=PersistentLocalVolumes=true,VolumeScheduling=true,MountPropagation=true"\n\1/;s/\(^ExecStart=.\+$\)/\1 $KUBELET_FEATURE_GATES/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf && sudo cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CERTIFICATE_ARGS=--rotate-certificates=true --cert-dir=/var/lib/kubelet/pki"
Environment="KUBELET_FEATURE_GATES=--feature-gates=PersistentLocalVolumes=true,VolumeScheduling=true,MountPropagation=true"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS $KUBELET_FEATURE_GATES
```

### Restart master and slaves

Master
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ sudo systemctl daemon-reload && sudo systemctl restart docker.service kubelet.service 
```

Worker1
```
[vagrant@rookdev-172-17-4-63 ~]$ sudo systemctl daemon-reload && sudo systemctl restart docker.service kubelet.service
```

Worker2
```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ sudo systemctl daemon-reload && sudo systemctl restart docker.service kubelet.service
```

## DevOps

Mount disk at each node

1. hostpath-provisioner

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl create -f template/hostpath-provisioner.yaml 
```

1. localvolume-provisioner

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl create -f template/local-storage-provision.yaml 
```

Provision pre-defined volume for namespace __example-system__

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl create -f mysql-volume-example.yaml 
persistentvolumeclaim "hostpath" created
persistentvolume "example-local-storage" created
persistentvolumeclaim "example-local-claim" created
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl -n example-system get pvc
NAME                             STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
example-local-claim              Bound     example-local-storage                      10Mi       RWO            local-storage      6h
hostpath                         Bound     pvc-9081b705-07ba-11e8-beb1-525400224e72   1Mi        RWX            example-hostpath   6h
```

Provison CRD

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/mysql-operator/mysql-crd.yaml 
customresourcedefinition "mariadbs.example.com" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get crd mariadbs.example.com
NAME                   AGE
mariadbs.example.com   1h
```

Provison Operator

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl create -f mysql-operator-deployment.yaml 
clusterrole "mysql-operator" created
serviceaccount "mysql-operator" created
clusterrolebinding "mysql-operator" created
deployment "mysql-operator" created
service "mysql-operator" created
Error from server (AlreadyExists): error when creating "mysql-operator-deployment.yaml": namespaces "example-system" already exists
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl -n example-system get deploy 
NAME             DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
mysql-operator   1         1         1            1           40s
redis-operator   1         1         1            1           7d
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl -n example-system get pods -o wide -l app=mysql-operator
NAME                             READY     STATUS    RESTARTS   AGE       IP            NODE
mysql-operator-f499f4c99-rqd2c   1/1       Running   0          1m        10.244.3.12   rookdev-172-17-4-63
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl -n example-system logs mysql-operator-f499f4c99-rqd2c
I0202 01:12:34.782777       1 controller.go:155] Setting up event handlers
time="2018-02-02T01:12:34Z" level=info msg="Starting controller" pkg=controller
I0202 01:12:34.784916       1 controller.go:230] Waiting for informer caches to sync
time="2018-02-02T01:12:35Z" level=info msg="Sync completed" pkg=controller
I0202 01:12:35.590405       1 controller.go:236] Starting workers
I0202 01:12:35.590679       1 controller.go:242] Started workers
```

Provison MariaDB resource

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl create -f mysql-provision_mariadb_10_2.yaml 
mariadb "demo-mariadb-galera" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl -n example-system get svc -l example.com/mysql-operator=demo-mariadb-galera
NAME                  TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                                        AGE
demo-mariadb-galera   ClusterIP   None         <none>        3306/TCP,4567/TCP,4567/UDP,4568/TCP,4444/TCP   28s
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl -n example-system get pvc -l example.com/mysql-operator=demo-mariadb-galera
NAME                             STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
hostpath-demo-mariadb-galera-0   Bound     pvc-3ab110e9-0851-11e8-a030-525400224e72   50Mi       RWO            example-hostpath   14m
hostpath-demo-mariadb-galera-1   Bound     pvc-44150f60-0851-11e8-a030-525400224e72   50Mi       RWO            example-hostpath   14m
hostpath-demo-mariadb-galera-2   Bound     pvc-5199ea5d-0851-11e8-a030-525400224e72   50Mi       RWO            example-hostpath   14m
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl -n example-system get pods -o wide -l example.com/mysql-operator=demo-mariadb-galera
NAME                    READY     STATUS    RESTARTS   AGE       IP            NODE
demo-mariadb-galera-0   1/1       Running   1          12m       10.244.3.57   rookdev-172-17-4-63
demo-mariadb-galera-1   1/1       Running   1          12m       10.244.2.34   rookdev-172-17-4-61
demo-mariadb-galera-2   1/1       Running   0          11m       10.244.3.58   rookdev-172-17-4-63
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl -n example-system exec -ti demo-mariadb-galera-1 -- mysql --user=testuser --password=testpass -e "SHOW STATUS LIKE 'wsrep_%';"
+------------------------------+----------------------------------------------------+
| Variable_name                | Value                                              |
+------------------------------+----------------------------------------------------+
| wsrep_apply_oooe             | 0.000000                                           |
| wsrep_apply_oool             | 0.000000                                           |
| wsrep_apply_window           | 1.000000                                           |
| wsrep_causal_reads           | 0                                                  |
| wsrep_cert_deps_distance     | 0.000000                                           |
| wsrep_cert_index_size        | 0                                                  |
| wsrep_cert_interval          | 0.000000                                           |
| wsrep_cluster_conf_id        | 7                                                  |
| wsrep_cluster_size           | 3                                                  |
| wsrep_cluster_state_uuid     | 4769aaa7-0851-11e8-81c7-935ba6ccd5d0               |
| wsrep_cluster_status         | Primary                                            |
| wsrep_commit_oooe            | 0.000000                                           |
| wsrep_commit_oool            | 0.000000                                           |
| wsrep_commit_window          | 1.000000                                           |
| wsrep_connected              | ON                                                 |
| wsrep_desync_count           | 0                                                  |
| wsrep_evs_delayed            |                                                    |
| wsrep_evs_evict_list         |                                                    |
| wsrep_evs_repl_latency       | 0/0/0/0/0                                          |
| wsrep_evs_state              | OPERATIONAL                                        |
| wsrep_flow_control_paused    | 0.000000                                           |
| wsrep_flow_control_paused_ns | 0                                                  |
| wsrep_flow_control_recv      | 0                                                  |
| wsrep_flow_control_sent      | 0                                                  |
| wsrep_gcomm_uuid             | d23b43f3-0855-11e8-8507-fa23577e1608               |
| wsrep_incoming_addresses     | 10.244.3.57:3306,10.244.3.58:3306,10.244.2.50:3306 |
| wsrep_last_committed         | 10507                                              |
| wsrep_local_bf_aborts        | 0                                                  |
| wsrep_local_cached_downto    | 18446744073709551615                               |
| wsrep_local_cert_failures    | 0                                                  |
| wsrep_local_commits          | 0                                                  |
| wsrep_local_index            | 2                                                  |
| wsrep_local_recv_queue       | 0                                                  |
| wsrep_local_recv_queue_avg   | 0.000000                                           |
| wsrep_local_recv_queue_max   | 1                                                  |
| wsrep_local_recv_queue_min   | 0                                                  |
| wsrep_local_replays          | 0                                                  |
| wsrep_local_send_queue       | 0                                                  |
| wsrep_local_send_queue_avg   | 0.000000                                           |
| wsrep_local_send_queue_max   | 1                                                  |
| wsrep_local_send_queue_min   | 0                                                  |
| wsrep_local_state            | 4                                                  |
| wsrep_local_state_comment    | Synced                                             |
| wsrep_local_state_uuid       | 4769aaa7-0851-11e8-81c7-935ba6ccd5d0               |
| wsrep_protocol_version       | 7                                                  |
| wsrep_provider_name          | Galera                                             |
| wsrep_provider_vendor        | Codership Oy <info@codership.com>                  |
| wsrep_provider_version       | 25.3.22(r3764)                                     |
| wsrep_ready                  | ON                                                 |
| wsrep_received               | 3                                                  |
| wsrep_received_bytes         | 332                                                |
| wsrep_repl_data_bytes        | 0                                                  |
| wsrep_repl_keys              | 0                                                  |
| wsrep_repl_keys_bytes        | 0                                                  |
| wsrep_repl_other_bytes       | 0                                                  |
| wsrep_replicated             | 0                                                  |
| wsrep_replicated_bytes       | 0                                                  |
| wsrep_thread_count           | 2                                                  |
+------------------------------+----------------------------------------------------+
```

## CI/CD

_Note: re-deploying must clean old data in hostpath, or GALERA may run incorrect way_

Inside update operator with alternative image tag
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl -n example-system set image deployment mysql-operator mysql-operator=docker.io/tangfeixiong/mysql-operator       
deployment "mysql-operator" image updated
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl -n example-system set image deployment mysql-operator mysql-operator=docker.io/tangfeixiong/mysql-operator:latest       
deployment "mysql-operator" image updated
```

Outside investation
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ ./bin/mysql-operator serve --kubeconfig --logtostderr -v 5
I0201 23:20:47.567128   29861 controller.go:155] Setting up event handlers
INFO[0000] Starting controller                           pkg=controller
I0201 23:20:47.569090   29861 controller.go:230] Waiting for informer caches to sync
E0201 23:20:47.624381   29861 reflector.go:205] github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/client/informers/externalversions/factory.go:58: Failed to list *v1alpha1.Mariadb: the server could not find the requested resource (get mariadbs.example.com)
E0201 23:20:48.626490   29861 reflector.go:205] github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/client/informers/externalversions/factory.go:58: Failed to list *v1alpha1.Mariadb: the server could not find the requested resource (get mariadbs.example.com)
### snip... after create CRD
INFO[0172] Sync completed                                pkg=controller
I0201 23:23:40.070037   29861 controller.go:236] Starting workers
I0201 23:23:40.070421   29861 controller.go:242] Started workers
### after provison custome resource
I0201 23:25:30.352523   29861 controller.go:382] service "demo-mariadb-galera" not found
E0201 23:25:30.352552   29861 controller.go:320] error syncing 'default/demo-mariadb-galera': service "demo-mariadb-galera" not found
I0201 23:25:40.038939   29861 controller.go:303] Successfully synced 'default/demo-mariadb-galera'
I0201 23:25:40.045086   29861 event.go:218] Event(v1.ObjectReference{Kind:"Mariadb", Namespace:"default", Name:"demo-mariadb-galera", UID:"314b14eb-07a7-11e8-beb1-525400224e72", APIVersion:"example.com", ResourceVersion:"2437109", FieldPath:""}): type: 'Normal' reason: 'Synced' Foo synced successfully
I0201 23:26:10.018212   29861 controller.go:303] Successfully synced 'default/demo-mariadb-galera'
I0201 23:26:10.019597   29861 event.go:218] Event(v1.ObjectReference{Kind:"Mariadb", Namespace:"default", Name:"demo-mariadb-galera", UID:"314b14eb-07a7-11e8-beb1-525400224e72", APIVersion:"example.com", ResourceVersion:"2437109", FieldPath:""}): type: 'Normal' reason: 'Synced' Foo synced successfully
### snip...
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/mysql-operator/mysql-crd.yaml 
customresourcedefinition "mariadbs.example.com" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/mysql-operator/mysql-provision_mariadb_10_2.yaml 
mariadb "demo-mariadb-galera" created
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl get pvc -l example.com/mysql-operator=demo-mariadb-galera
NAME                             STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
hostpath-demo-mariadb-galera-0   Bound     pvc-ecadc4ff-07b1-11e8-beb1-525400224e72   50Mi       RWX            example-hostpath   <invalid>
hostpath-demo-mariadb-galera-1   Bound     pvc-ff078c8c-07b1-11e8-beb1-525400224e72   50Mi       RWX            example-hostpath   <invalid>
hostpath-demo-mariadb-galera-2   Bound     pvc-346b4631-07b2-11e8-beb1-525400224e72   50Mi       RWX            example-hostpath   <invalid>
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl get pods -l example.com/mysql-operator=demo-mariadb-galera -o wide
NAME                    READY     STATUS    RESTARTS   AGE         IP             NODE
demo-mariadb-galera-0   1/1       Running   0          52s         10.244.3.10    rookdev-172-17-4-63
demo-mariadb-galera-1   1/1       Running   1          21s         10.244.2.229   rookdev-172-17-4-61
demo-mariadb-galera-2   1/1       Running   1          <invalid>   10.244.3.11    rookdev-172-17-4-63
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl exec -ti demo-mariadb-galera-2 -- mysql --user=testuser --password=testpass -e "SHOW STATUS LIKE 'wsrep_%';"
+------------------------------+-----------------------------------------------------+
| Variable_name                | Value                                               |
+------------------------------+-----------------------------------------------------+
| wsrep_apply_oooe             | 0.000000                                            |
| wsrep_apply_oool             | 0.000000                                            |
| wsrep_apply_window           | 0.000000                                            |
| wsrep_causal_reads           | 0                                                   |
| wsrep_cert_deps_distance     | 0.000000                                            |
| wsrep_cert_index_size        | 0                                                   |
| wsrep_cert_interval          | 0.000000                                            |
| wsrep_cluster_conf_id        | 7                                                   |
| wsrep_cluster_size           | 3                                                   |
| wsrep_cluster_state_uuid     | f9468a17-07b1-11e8-a36d-b314242c9b01                |
| wsrep_cluster_status         | Primary                                             |
| wsrep_commit_oooe            | 0.000000                                            |
| wsrep_commit_oool            | 0.000000                                            |
| wsrep_commit_window          | 0.000000                                            |
| wsrep_connected              | ON                                                  |
| wsrep_desync_count           | 0                                                   |
| wsrep_evs_delayed            |                                                     |
| wsrep_evs_evict_list         |                                                     |
| wsrep_evs_repl_latency       | 0/0/0/0/0                                           |
| wsrep_evs_state              | OPERATIONAL                                         |
| wsrep_flow_control_paused    | 0.000000                                            |
| wsrep_flow_control_paused_ns | 0                                                   |
| wsrep_flow_control_recv      | 0                                                   |
| wsrep_flow_control_sent      | 0                                                   |
| wsrep_gcomm_uuid             | 462d66b7-07b2-11e8-a6ed-3f07c0dd7852                |
| wsrep_incoming_addresses     | 10.244.3.10:3306,10.244.3.11:3306,10.244.2.229:3306 |
| wsrep_last_committed         | 7100                                                |
| wsrep_local_bf_aborts        | 0                                                   |
| wsrep_local_cached_downto    | 18446744073709551615                                |
| wsrep_local_cert_failures    | 0                                                   |
| wsrep_local_commits          | 0                                                   |
| wsrep_local_index            | 1                                                   |
| wsrep_local_recv_queue       | 0                                                   |
| wsrep_local_recv_queue_avg   | 0.000000                                            |
| wsrep_local_recv_queue_max   | 1                                                   |
| wsrep_local_recv_queue_min   | 0                                                   |
| wsrep_local_replays          | 0                                                   |
| wsrep_local_send_queue       | 0                                                   |
| wsrep_local_send_queue_avg   | 0.000000                                            |
| wsrep_local_send_queue_max   | 1                                                   |
| wsrep_local_send_queue_min   | 0                                                   |
| wsrep_local_state            | 4                                                   |
| wsrep_local_state_comment    | Synced                                              |
| wsrep_local_state_uuid       | f9468a17-07b1-11e8-a36d-b314242c9b01                |
| wsrep_protocol_version       | 7                                                   |
| wsrep_provider_name          | Galera                                              |
| wsrep_provider_vendor        | Codership Oy <info@codership.com>                   |
| wsrep_provider_version       | 25.3.22(r3764)                                      |
| wsrep_ready                  | ON                                                  |
| wsrep_received               | 6                                                   |
| wsrep_received_bytes         | 1030                                                |
| wsrep_repl_data_bytes        | 0                                                   |
| wsrep_repl_keys              | 0                                                   |
| wsrep_repl_keys_bytes        | 0                                                   |
| wsrep_repl_other_bytes       | 0                                                   |
| wsrep_replicated             | 0                                                   |
| wsrep_replicated_bytes       | 0                                                   |
| wsrep_thread_count           | 2                                                   |
+------------------------------+-----------------------------------------------------+
```

Clean before re-deploying
```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl delete pvc -l example.com/mysql-operator=demo-mariadb-galera
persistentvolumeclaim "hostpath-demo-mariadb-galera-0" deleted
persistentvolumeclaim "hostpath-demo-mariadb-galera-1" deleted
persistentvolumeclaim "hostpath-demo-mariadb-galera-2" deleted
```

## Develop
Artifacts
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go make go-bindata-artifacts
```

Image
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go make docker-
### snip: more packages...
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/initcnf
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/server
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/cmd
github.com/tangfeixiong/go-to-kubernetes/mysql-operator
#@docker build --no-cache --force-rm -t docker.io/tangfeixiong/mysql-operator:1801311008-gitb5fa67c ./
Sending build context to Docker daemon 20.36 MB
Step 1/7 : FROM busybox
 ---> 6ad733544a63
Step 2/7 : LABEL "maintainer" "tangfeixiong <tangfx128@gmail.com>" "project" "https://github.com/tangfeixiong/go-to-kubernetes" "name" "mysql-operator" "version" "0.1" "created-by" '{"go":"v1.9.2","kubernetes":"v1.8","docker":"1.13"}'
 ---> Running in f9b873e6fcea
 ---> cd87e8a21934
Removing intermediate container f9b873e6fcea
Step 3/7 : COPY bin/mysql-operator /
 ---> 82fcd8f235c5
Removing intermediate container 44dd257836db
Step 4/7 : ENV DOCKER_API_VERSION '1.12' DOCKER_CONFIG_JSON '{"auths": {"localhost:5000": {"auth": "","email": ""}}}' REGISTRY_CERTS_JSON '{"localhost:5000": {"ca_base64": "", "crt_base64": "", "key_base64": ""}}'
 ---> Running in e8c80357cafd
 ---> d9985c1aca75
Removing intermediate container e8c80357cafd
Step 5/7 : EXPOSE 10002 10001
 ---> Running in d41cd32bd854
 ---> c67d35e977bd
Removing intermediate container d41cd32bd854
Step 6/7 : ENTRYPOINT /mysql-operator serve
 ---> Running in cb91c65618ef
 ---> ea4905a6ad86
Removing intermediate container cb91c65618ef
Step 7/7 : CMD -v 2 --logtostderr=true
 ---> Running in b2344ac901d5
 ---> 73b28b1bb617
Removing intermediate container b2344ac901d5
Successfully built 73b28b1bb617
The push refers to a repository [docker.io/tangfeixiong/mysql-operator]
bbbd56bf89d5: Pushed 
0271b8eebde3: Mounted from tangfeixiong/redis-operator 
latest: digest: sha256:367926b68253d259d306374a8aca6c8cb8a1d568ab57b017b212253bc709bd8d size: 738
```

## Outside K8s

### Investigate custome resource

Install `spectest`
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go install -v ../cmd/spectest/
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/artifact
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/sts
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/svc
github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/mysqlorbranches
github.com/tangfeixiong/go-to-kubernetes/cmd/spectest
```
crd
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ spectest create mariadb-crd -kubeconfig -v=5 -logtostderr
Created CRD "mysqlfamilies.example.com".
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get crd mysqlfamilies.example.com
NAME                        AGE
mysqlfamilies.example.com   1m
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl delete crd mysqlfamilies.example.com        
customresourcedefinition "mysqlfamilies.example.com" deleted
```

local-storage
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ spectest create mariadb-localstorage -kubeconfig -v=5 -logtostderr
Created StorageClass "local-storage".
panic: PersistentVolume "example-local-pv" is invalid: [metadata.annotations: Forbidden: Storage node affinity is disabled by feature-gate, spec.local: Forbidden: Local volumes are disabled by feature-gate]

goroutine 1 [running]:
github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/mysqlorbranches.CreateLocalStorage(0x177c7a0, 0xc4204424e0)
	/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/mysqlorbranches/create.go:76 +0x814
main.main()
	/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/main.go:113 +0x921
```
_Issue: require feature-gate args, refer to prerequisites topic_

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ spectest create mariadb-localstorage -kubeconfig -v=5 -logtostderr
Already exists StorageClass "local-storage".
Created PersistentVolume "example-local-pv".
Created PersistentVolumeCalim "example-local-claim".
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get storageclass,pv,pvc
NAME                           PROVISIONER                    AGE
storageclasses/local-storage   kubernetes.io/no-provisioner   2h

NAME                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM                                                      STORAGECLASS    REASON    AGE
pv/example-local-pv        60Mi       RWO            Retain           Bound       default/example-local-claim                                local-storage             1m

NAME                      STATUS    VOLUME             CAPACITY   ACCESS MODES   STORAGECLASS    AGE
pvc/example-local-claim   Bound     example-local-pv   60Mi       RWO            local-storage   1m
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl get storageclass
NAME            PROVISIONER                    AGE
local-storage   kubernetes.io/no-provisioner   1m
```

Mount local volume into workers
```
[vagrant@rookdev-172-17-4-63 ~]$ sudo mkdir -p /mnt/disks/vol1 /mnt/disks/vol2 /mnt/disks/vol3
[vagrant@rookdev-172-17-4-63 ~]$ sudo mkdir -p /mnt/fast-disks/vol1 /mnt/fast-disks/vol2 /mnt/fast-disks/vol3
```

```
ubuntu@rookdev-172-17-4-61:~$ sudo mkdir -p /mnt/disks/vol1 /mnt/disks/vol2 /mnt/disks/vol3
ubuntu@rookdev-172-17-4-61:~$ sudo mkdir -p /mnt/fast-disks/vol1 /mnt/fast-disks/vol2 /mnt/fast-disks/vol3
```

Local volume static provisioner
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/kubernetes-incubator/external-storage/local-volume/provisioner/deployment/kubernetes/admin_account.yaml 
serviceaccount "local-storage-admin" created
clusterrolebinding "local-storage-provisioner-pv-binding" created
clusterrolebinding "local-storage-provisioner-node-binding" created
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/kubernetes-incubator/external-storage/local-volume/provisioner/deployment/kubernetes/provisioner-daemonset.yaml 
daemonset "local-volume-provisioner" created
```
_Issue: missing ConfigMap_

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl delete -f /Users/fanhongling/Downloads/workspace/src/github.com/kubernetes-incubator/external-storage/local-volume/provisioner/deployment/kubernetes/provisioner-daemonset.yaml
daemonset "local-volume-provisioner" deleted
```

Helm
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ curl http://storage.googleapis.com/kubernetes-helm/helm-v2.7.2-linux-amd64.tar.gz  | sudo tar --strip-components 1 -C /Users/fanhongling/Downloads/99-mirror/linux-bin linux-amd64/helm -zxf -
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 11.6M  100 11.6M    0     0  3556k      0  0:00:03  0:00:03 --:--:-- 3556k
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ helm template --set engine=baremetal /Users/fanhongling/Downloads/workspace/src/github.com/kubernetes-incubator/external-storage/local-volume/helm/provisioner/
---
# Source: provisioner/templates/provisioner.yaml
 
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-provisioner-config 
  namespace: default 
data:
  storageClassMap: |     
    fast-disks:
       hostDir: /mnt/fast-disks
       mountDir:  /mnt/fast-disks  
---
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: local-volume-provisioner
  namespace: default
  labels:
    app: local-volume-provisioner
spec:
  selector:
    matchLabels:
      app: local-volume-provisioner 
  template:
    metadata:
      labels:
        app: local-volume-provisioner
    spec:
      serviceAccountName: local-storage-admin
      containers:
        - image: "quay.io/external_storage/local-volume-provisioner:latest"
          imagePullPolicy: "Always"
          name: provisioner 
          securityContext:
            privileged: true
          env:
          - name: MY_NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
          volumeMounts:
            - mountPath: /etc/provisioner/config 
              name: provisioner-config
              readOnly: true             
            - mountPath:  /mnt/fast-disks 
              name: fast-disks 
      volumes:
        - name: provisioner-config
          configMap:
            name: local-provisioner-config         
        - name: fast-disks
          hostPath:
            path: /mnt/fast-disks 
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl create -f template/local-volume-static-provisioner.yaml 
configmap "local-provisioner-config" created
daemonset "local-volume-provisioner" created
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl get ds
NAME                       DESIRED   CURRENT   READY     UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
local-volume-provisioner   4         4         2         4            2           <none>          48s
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pods -l app=local-volume-provisioner -o wide
NAME                             READY     STATUS     RESTARTS   AGE       IP             NODE
local-volume-provisioner-hhp6c   0/1       NodeLost   0          4m        <none>         kubedev-172-17-4-65
local-volume-provisioner-r4wzt   0/1       NodeLost   0          4m        <none>         kubedev-172-17-4-55
local-volume-provisioner-rnwjr   1/1       Running    0          4m        10.244.2.213   rookdev-172-17-4-61
local-volume-provisioner-w7p9r   1/1       Running    0          4m        10.244.3.223   rookdev-172-17-4-63
```

_Issue: must be mount point_
```
E0131 21:48:59.991419       1 discovery.go:172] Path "/mnt/fast-disks/vol1" is not an actual mountpoint
E0131 21:48:59.991450       1 discovery.go:172] Path "/mnt/fast-disks/vol2" is not an actual mountpoint
E0131 21:48:59.991461       1 discovery.go:172] Path "/mnt/fast-disks/vol3" is not an actual mountpoint
```

hostpath-provisioner
```
[vagrant@rookdev-172-17-4-63 ~]$ mkdir -p /tmp/hostpath-provisioner
[ubuntu@rookdev-172-17-4-61:~$ mkdir -p /tmp/hostpath-provisioner
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl create -f template/hostpath-provisioner.yaml 
storageclass "example-hostpath" created
serviceaccount "hostpath-admin" created
clusterrolebinding "hostpath-provisioner-pv-binding" created
clusterrolebinding "hostpath-provisioner-node-binding" created
daemonset "hostpath-provisioner" created
persistentvolumeclaim "hostpath" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get storageclass example-hostpath
NAME               PROVISIONER                    AGE
example-hostpath   example.com/hostpath           52s
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get ds hostpath-provisioner
NAME                   DESIRED   CURRENT   READY     UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
hostpath-provisioner   4         4         2         4            2           <none>          3m
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl get pods -l app=hostpath-provisioner
NAME                         READY     STATUS    RESTARTS   AGE
hostpath-provisioner-pp2w6   0/1       Pending   0          2m
hostpath-provisioner-pzg6q   0/1       Pending   0          2m
hostpath-provisioner-sndsr   1/1       Running   0          2m
hostpath-provisioner-zjhrp   1/1       Running   0          2m
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl get pvc hostpath
NAME       STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
hostpath   Bound     pvc-87039137-06da-11e8-beb1-525400224e72   1Mi        RWX            example-hostpath   2m
```

Service
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ spectest create mariadb-svc -kubeconfig -v=5 -logtostderr
Created service "mariadb-galera".
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get svc -l app=mariadb
NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                        AGE
mariadb-galera   ClusterIP   10.96.201.234   <none>        3306/TCP,4567/TCP,4567/UDP,4568/TCP,4444/TCP   2m
```

StatefulSet
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ spectest create mariadb-sts -kubeconfig -v=5 -logtostderr
Created statefulset "mariadb-galera".
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get sts -l app=mariadb
NAME                   DESIRED   CURRENT   AGE
mariadb-galera         3         1         21s
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl get pvc -l app=mariadb
NAME                        STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
hostpath-mariadb-galera-0   Bound     pvc-4c66444f-06db-11e8-beb1-525400224e72   50Mi       RWX            example-hostpath   20m
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pods -l app=mariadb -o wide
NAME               READY     STATUS    RESTARTS   AGE       IP             NODE
mariadb-galera-0   1/1       Running   0          1m        10.244.3.253   rookdev-172-17-4-63
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl logs mariadb-galera-0
2018-02-01 16:42:59 139928297252736 [Note] mysqld (mysqld 10.2.12-MariaDB-10.2.12+maria~jessie) starting as process 1 ...
2018-02-01 16:42:59 139928297252736 [Note] WSREP: Read nil XID from storage engines, skipping position init
2018-02-01 16:42:59 139928297252736 [Note] WSREP: wsrep_load(): loading provider library '/usr/lib/libgalera_smm.so'
2018-02-01 16:42:59 139928297252736 [Note] WSREP: wsrep_load(): Galera 25.3.22(r3764) by Codership Oy <info@codership.com> loaded successfully.
2018-02-01 16:42:59 139928297252736 [Note] WSREP: CRC-32C: using hardware acceleration.
2018-02-01 16:42:59 139928297252736 [Note] WSREP: Found saved state: 00000000-0000-0000-0000-000000000000:-1, safe_to_bootstrap: 1
2018-02-01 16:42:59 139928297252736 [Note] WSREP: Passing config to GCS: base_dir = /var/lib/mysql/; base_host = 10.244.3.253; base_port = 4567; cert.log_conflicts = no; debug = no; evs.auto_evict = 0; evs.delay_margin = PT1S; evs.delayed_keep_period = PT30S; evs.inactive_check_period = PT0.5S; evs.inactive_timeout = PT15S; evs.join_retrans_period = PT1S; evs.max_install_timeouts = 3; evs.send_window = 4; evs.stats_report_period = PT1M; evs.suspect_timeout = PT5S; evs.user_send_window = 2; evs.view_forget_timeout = PT24H; gcache.dir = /var/lib/mysql/; gcache.keep_pages_size = 0; gcache.mem_size = 0; gcache.name = /var/lib/mysql//galera.cache; gcache.page_size = 128M; gcache.recover = no; gcache.size = 128M; gcomm.thread_prio = ; gcs.fc_debug = 0; gcs.fc_factor = 1.0; gcs.fc_limit = 16; gcs.fc_master_slave = no; gcs.max_packet_size = 64500; gcs.max_throttle = 0.25; gcs.recv_q_hard_limit = 9223372036854775807; gcs.recv_q_soft_limit = 0.25; gcs.sync_donor = no; gmcast.segment = 0; gmcast.version = 0; pc.announce_timeout = PT3S; pc.checksum = false; pc.i
2018-02-01 16:43:00 139928297252736 [Note] WSREP: GCache history reset: 00000000-0000-0000-0000-000000000000:0 -> 00000000-0000-0000-0000-000000000000:-1
2018-02-01 16:43:00 139928297252736 [Note] WSREP: Assign initial position for certification: -1, protocol version: -1
2018-02-01 16:43:00 139928297252736 [Note] WSREP: wsrep_sst_grab()
2018-02-01 16:43:00 139928297252736 [Note] WSREP: Start replication
2018-02-01 16:43:00 139928297252736 [Note] WSREP: Setting initial position to 00000000-0000-0000-0000-000000000000:-1
2018-02-01 16:43:00 139928297252736 [Note] WSREP: protonet asio version 0
2018-02-01 16:43:00 139928297252736 [Note] WSREP: Using CRC-32C for message checksums.
2018-02-01 16:43:00 139928297252736 [Note] WSREP: backend: asio
2018-02-01 16:43:00 139928297252736 [Note] WSREP: gcomm thread scheduling priority set to other:0 
2018-02-01 16:43:00 139928297252736 [Warning] WSREP: access file(/var/lib/mysql//gvwstate.dat) failed(No such file or directory)
2018-02-01 16:43:00 139928297252736 [Note] WSREP: restore pc from disk failed
2018-02-01 16:43:00 139928297252736 [Note] WSREP: GMCast version 0
2018-02-01 16:43:00 139928297252736 [Note] WSREP: (f69d2c57, 'tcp://0.0.0.0:4567') listening at tcp://0.0.0.0:4567
2018-02-01 16:43:00 139928297252736 [Note] WSREP: (f69d2c57, 'tcp://0.0.0.0:4567') multicast: , ttl: 1
2018-02-01 16:43:00 139928297252736 [Note] WSREP: EVS version 0
2018-02-01 16:43:00 139928297252736 [Note] WSREP: gcomm: connecting to group 'mariadb-galera', peer ''
2018-02-01 16:43:00 139928297252736 [Note] WSREP: start_prim is enabled, turn off pc_recovery
2018-02-01 16:43:00 139928297252736 [Note] WSREP: Node f69d2c57 state prim
2018-02-01 16:43:00 139928297252736 [Note] WSREP: view(view_id(PRIM,f69d2c57,1) memb {
	f69d2c57,0
} joined {
} left {
} partitioned {
})
2018-02-01 16:43:00 139928297252736 [Note] WSREP: save pc into disk
2018-02-01 16:43:00 139928297252736 [Note] WSREP: gcomm: connected
2018-02-01 16:43:00 139928297252736 [Note] WSREP: Changing maximum packet size to 64500, resulting msg size: 32636
2018-02-01 16:43:00 139928297252736 [Note] WSREP: Shifting CLOSED -> OPEN (TO: 0)
2018-02-01 16:43:00 139928297252736 [Note] WSREP: Opened channel 'mariadb-galera'
2018-02-01 16:43:00 139928297252736 [Note] WSREP: Waiting for SST to complete.
2018-02-01 16:43:00 139927907284736 [Note] WSREP: New COMPONENT: primary = yes, bootstrap = no, my_idx = 0, memb_num = 1
2018-02-01 16:43:00 139927907284736 [Note] WSREP: Starting new group from scratch: f69e5fe3-076e-11e8-83e0-6646cc4cd744
2018-02-01 16:43:00 139927907284736 [Note] WSREP: STATE_EXCHANGE: sent state UUID: f69e61b7-076e-11e8-a9ae-4e180f730401
2018-02-01 16:43:00 139927907284736 [Note] WSREP: STATE EXCHANGE: sent state msg: f69e61b7-076e-11e8-a9ae-4e180f730401
2018-02-01 16:43:00 139927907284736 [Note] WSREP: STATE EXCHANGE: got state msg: f69e61b7-076e-11e8-a9ae-4e180f730401 from 0 (mariadb-galera-0)
2018-02-01 16:43:00 139927907284736 [Note] WSREP: Quorum results:
	version    = 4,
	component  = PRIMARY,
	conf_id    = 0,
	members    = 1/1 (joined/total),
	act_id     = 0,
	last_appl. = -1,
	protocols  = 0/7/3 (gcs/repl/appl),
	group UUID = f69e5fe3-076e-11e8-83e0-6646cc4cd744
2018-02-01 16:43:00 139927907284736 [Note] WSREP: Flow-control interval: [16, 16]
2018-02-01 16:43:00 139927907284736 [Note] WSREP: Trying to continue unpaused monitor
2018-02-01 16:43:00 139927907284736 [Note] WSREP: Restored state OPEN -> JOINED (0)
2018-02-01 16:43:00 139927907284736 [Note] WSREP: Member 0.0 (mariadb-galera-0) synced with group.
2018-02-01 16:43:00 139927907284736 [Note] WSREP: Shifting JOINED -> SYNCED (TO: 0)
2018-02-01 16:43:00 139928295753472 [Note] WSREP: New cluster view: global state: f69e5fe3-076e-11e8-83e0-6646cc4cd744:0, view# 1: Primary, number of nodes: 1, my index: 0, protocol version 3
2018-02-01 16:43:00 139928297252736 [Note] WSREP: SST complete, seqno: 0
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: For Galera, using innodb_lock_schedule_algorithm=fcfs
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Uses event mutexes
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Compressed tables use zlib 1.2.8
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Using Linux native AIO
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Number of pools: 1
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Using SSE2 crc32 instructions
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Initializing buffer pool, total size = 256M, instances = 1, chunk size = 128M
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Completed initialization of buffer pool
2018-02-01 16:43:00 139927099475712 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Highest supported file format is Barracuda.
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: 128 out of 128 rollback segments are active.
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Creating shared tablespace for temporary tables
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: Waiting for purge to start
2018-02-01 16:43:00 139928297252736 [Note] InnoDB: 5.7.20 started; log sequence number 1619559
2018-02-01 16:43:00 139927091083008 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
2018-02-01 16:43:00 139927091083008 [Note] InnoDB: Buffer pool(s) load completed at 180201 16:43:00
2018-02-01 16:43:00 139928297252736 [Note] Plugin 'FEEDBACK' is disabled.
2018-02-01 16:43:00 139928297252736 [Note] Server socket created on IP: '0.0.0.0'.
2018-02-01 16:43:00 139928297252736 [Warning] 'proxies_priv' entry '@% root@mariadb-galera-0' ignored in --skip-name-resolve mode.
2018-02-01 16:43:00 139928295753472 [Note] WSREP: wsrep_notify_cmd is not defined, skipping notification.
2018-02-01 16:43:00 139928295753472 [Note] WSREP: REPL Protocols: 7 (3, 2)
2018-02-01 16:43:00 139928295753472 [Note] WSREP: Assign initial position for certification: 0, protocol version: 3
2018-02-01 16:43:00 139927949248256 [Note] WSREP: Service thread queue flushed.
2018-02-01 16:43:00 139928295753472 [Note] WSREP: GCache history reset: 00000000-0000-0000-0000-000000000000:0 -> f69e5fe3-076e-11e8-83e0-6646cc4cd744:0
2018-02-01 16:43:00 139928295753472 [Note] WSREP: Synchronized with group, ready for connections
2018-02-01 16:43:00 139928295753472 [Note] WSREP: wsrep_notify_cmd is not defined, skipping notification.
2018-02-01 16:43:00 139928297252736 [Note] Reading of all Master_info entries succeded
2018-02-01 16:43:00 139928297252736 [Note] Added new Master_info '' to hash table
2018-02-01 16:43:00 139928297252736 [Note] mysqld: ready for connections.
Version: '10.2.12-MariaDB-10.2.12+maria~jessie'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  mariadb.org binary distribution
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl get pvc -l app=mariadb
NAME                        STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
hostpath-mariadb-galera-0   Bound     pvc-4c66444f-06db-11e8-beb1-525400224e72   50Mi       RWX            example-hostpath   2m
hostpath-mariadb-galera-1   Bound     pvc-1f1252a8-06eb-11e8-beb1-525400224e72   50Mi       RWX            example-hostpath   17s
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pod mariadb-galera-1 -o wide
NAME               READY     STATUS    RESTARTS   AGE       IP             NODE
mariadb-galera-1   1/1       Running   0          3m        10.244.2.225   rookdev-172-17-4-61
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl logs mariadb-galera-1
2018-02-01 16:43:54 140249689663360 [Note] mysqld (mysqld 10.2.12-MariaDB-10.2.12+maria~jessie) starting as process 1 ...
2018-02-01 16:43:54 140249689663360 [Note] WSREP: Read nil XID from storage engines, skipping position init
2018-02-01 16:43:54 140249689663360 [Note] WSREP: wsrep_load(): loading provider library '/usr/lib/libgalera_smm.so'
2018-02-01 16:43:54 140249689663360 [Note] WSREP: wsrep_load(): Galera 25.3.22(r3764) by Codership Oy <info@codership.com> loaded successfully.
2018-02-01 16:43:54 140249689663360 [Note] WSREP: CRC-32C: using hardware acceleration.
2018-02-01 16:43:54 140249689663360 [Warning] WSREP: Could not open state file for reading: '/var/lib/mysql//grastate.dat'
2018-02-01 16:43:54 140249689663360 [Note] WSREP: Found saved state: 00000000-0000-0000-0000-000000000000:-1, safe_to_bootstrap: 1
2018-02-01 16:43:54 140249689663360 [Note] WSREP: Passing config to GCS: base_dir = /var/lib/mysql/; base_host = 10.244.2.225; base_port = 4567; cert.log_conflicts = no; debug = no; evs.auto_evict = 0; evs.delay_margin = PT1S; evs.delayed_keep_period = PT30S; evs.inactive_check_period = PT0.5S; evs.inactive_timeout = PT15S; evs.join_retrans_period = PT1S; evs.max_install_timeouts = 3; evs.send_window = 4; evs.stats_report_period = PT1M; evs.suspect_timeout = PT5S; evs.user_send_window = 2; evs.view_forget_timeout = PT24H; gcache.dir = /var/lib/mysql/; gcache.keep_pages_size = 0; gcache.mem_size = 0; gcache.name = /var/lib/mysql//galera.cache; gcache.page_size = 128M; gcache.recover = no; gcache.size = 128M; gcomm.thread_prio = ; gcs.fc_debug = 0; gcs.fc_factor = 1.0; gcs.fc_limit = 16; gcs.fc_master_slave = no; gcs.max_packet_size = 64500; gcs.max_throttle = 0.25; gcs.recv_q_hard_limit = 9223372036854775807; gcs.recv_q_soft_limit = 0.25; gcs.sync_donor = no; gmcast.segment = 0; gmcast.version = 0; pc.announce_timeout = PT3S; pc.checksum = false; pc.i
2018-02-01 16:43:54 140249689663360 [Note] WSREP: GCache history reset: 00000000-0000-0000-0000-000000000000:0 -> 00000000-0000-0000-0000-000000000000:-1
2018-02-01 16:43:54 140249689663360 [Note] WSREP: Assign initial position for certification: -1, protocol version: -1
2018-02-01 16:43:54 140249689663360 [Note] WSREP: wsrep_sst_grab()
2018-02-01 16:43:54 140249689663360 [Note] WSREP: Start replication
2018-02-01 16:43:54 140249689663360 [Note] WSREP: Setting initial position to 00000000-0000-0000-0000-000000000000:-1
2018-02-01 16:43:54 140249689663360 [Note] WSREP: protonet asio version 0
2018-02-01 16:43:54 140249689663360 [Note] WSREP: Using CRC-32C for message checksums.
2018-02-01 16:43:54 140249689663360 [Note] WSREP: backend: asio
2018-02-01 16:43:54 140249689663360 [Note] WSREP: gcomm thread scheduling priority set to other:0 
2018-02-01 16:43:54 140249689663360 [Warning] WSREP: access file(/var/lib/mysql//gvwstate.dat) failed(No such file or directory)
2018-02-01 16:43:54 140249689663360 [Note] WSREP: restore pc from disk failed
2018-02-01 16:43:54 140249689663360 [Note] WSREP: GMCast version 0
2018-02-01 16:43:54 140249689663360 [Note] WSREP: (16f980a0, 'tcp://0.0.0.0:4567') listening at tcp://0.0.0.0:4567
2018-02-01 16:43:54 140249689663360 [Note] WSREP: (16f980a0, 'tcp://0.0.0.0:4567') multicast: , ttl: 1
2018-02-01 16:43:54 140249689663360 [Note] WSREP: EVS version 0
2018-02-01 16:43:54 140249689663360 [Note] WSREP: gcomm: connecting to group 'mariadb-galera', peer '10.244.3.253:'
2018-02-01 16:43:54 140249689663360 [Note] WSREP: (16f980a0, 'tcp://0.0.0.0:4567') connection established to f69d2c57 tcp://10.244.3.253:4567
2018-02-01 16:43:54 140249689663360 [Note] WSREP: (16f980a0, 'tcp://0.0.0.0:4567') turning message relay requesting on, nonlive peers: 
2018-02-01 16:43:54 140249689663360 [Note] WSREP: declaring f69d2c57 at tcp://10.244.3.253:4567 stable
2018-02-01 16:43:54 140249689663360 [Note] WSREP: Node f69d2c57 state prim
2018-02-01 16:43:54 140249689663360 [Note] WSREP: view(view_id(PRIM,16f980a0,2) memb {
	16f980a0,0
	f69d2c57,0
} joined {
} left {
} partitioned {
})
2018-02-01 16:43:54 140249689663360 [Note] WSREP: save pc into disk
2018-02-01 16:43:55 140249689663360 [Note] WSREP: gcomm: connected
2018-02-01 16:43:55 140249689663360 [Note] WSREP: Changing maximum packet size to 64500, resulting msg size: 32636
2018-02-01 16:43:55 140249689663360 [Note] WSREP: Shifting CLOSED -> OPEN (TO: 0)
2018-02-01 16:43:55 140249689663360 [Note] WSREP: Opened channel 'mariadb-galera'
2018-02-01 16:43:55 140249689663360 [Note] WSREP: Waiting for SST to complete.
2018-02-01 16:43:55 140249296508672 [Note] WSREP: New COMPONENT: primary = yes, bootstrap = no, my_idx = 0, memb_num = 2
2018-02-01 16:43:55 140249296508672 [Note] WSREP: STATE_EXCHANGE: sent state UUID: 17928ebb-076f-11e8-a57b-aa39ba6d3c68
2018-02-01 16:43:55 140249296508672 [Note] WSREP: STATE EXCHANGE: sent state msg: 17928ebb-076f-11e8-a57b-aa39ba6d3c68
2018-02-01 16:43:55 140249296508672 [Note] WSREP: STATE EXCHANGE: got state msg: 17928ebb-076f-11e8-a57b-aa39ba6d3c68 from 0 (mariadb-galera-1)
2018-02-01 16:43:55 140249296508672 [Note] WSREP: STATE EXCHANGE: got state msg: 17928ebb-076f-11e8-a57b-aa39ba6d3c68 from 1 (mariadb-galera-0)
2018-02-01 16:43:55 140249296508672 [Note] WSREP: Quorum results:
	version    = 4,
	component  = PRIMARY,
	conf_id    = 1,
	members    = 1/2 (joined/total),
	act_id     = 0,
	last_appl. = -1,
	protocols  = 0/7/3 (gcs/repl/appl),
	group UUID = f69e5fe3-076e-11e8-83e0-6646cc4cd744
2018-02-01 16:43:55 140249296508672 [Note] WSREP: Flow-control interval: [23, 23]
2018-02-01 16:43:55 140249296508672 [Note] WSREP: Trying to continue unpaused monitor
2018-02-01 16:43:55 140249296508672 [Note] WSREP: Shifting OPEN -> PRIMARY (TO: 0)
2018-02-01 16:43:55 140249688164096 [Note] WSREP: State transfer required: 
	Group state: f69e5fe3-076e-11e8-83e0-6646cc4cd744:0
	Local state: 00000000-0000-0000-0000-000000000000:-1
2018-02-01 16:43:55 140249688164096 [Note] WSREP: New cluster view: global state: f69e5fe3-076e-11e8-83e0-6646cc4cd744:0, view# 2: Primary, number of nodes: 2, my index: 0, protocol version 3
2018-02-01 16:43:55 140249688164096 [Warning] WSREP: Gap in state sequence. Need state transfer.
2018-02-01 16:43:55 140249288115968 [Note] WSREP: Running: 'wsrep_sst_rsync --role 'joiner' --address '10.244.2.225' --datadir '/var/lib/mysql/'   --parent '1'  '' '
2018-02-01 16:43:55 140249688164096 [Note] WSREP: Prepared SST request: rsync|10.244.2.225:4444/rsync_sst
2018-02-01 16:43:55 140249688164096 [Note] WSREP: wsrep_notify_cmd is not defined, skipping notification.
2018-02-01 16:43:55 140249688164096 [Note] WSREP: REPL Protocols: 7 (3, 2)
2018-02-01 16:43:55 140249688164096 [Note] WSREP: Assign initial position for certification: 0, protocol version: 3
2018-02-01 16:43:55 140249338472192 [Note] WSREP: Service thread queue flushed.
2018-02-01 16:43:55 140249688164096 [Warning] WSREP: Failed to prepare for incremental state transfer: Local state UUID (00000000-0000-0000-0000-000000000000) does not match group state UUID (f69e5fe3-076e-11e8-83e0-6646cc4cd744): 1 (Operation not permitted)
	 at galera/src/replicator_str.cpp:prepare_for_IST():482. IST will be unavailable.
2018-02-01 16:43:55 140249296508672 [Note] WSREP: Member 0.0 (mariadb-galera-1) requested state transfer from '*any*'. Selected 1.0 (mariadb-galera-0)(SYNCED) as donor.
2018-02-01 16:43:55 140249296508672 [Note] WSREP: Shifting PRIMARY -> JOINER (TO: 0)
2018-02-01 16:43:55 140249688164096 [Note] WSREP: Requesting state transfer: success, donor: 1
2018-02-01 16:43:55 140249688164096 [Note] WSREP: GCache history reset: 00000000-0000-0000-0000-000000000000:0 -> f69e5fe3-076e-11e8-83e0-6646cc4cd744:0
2018-02-01 16:43:57 140249304901376 [Note] WSREP: (16f980a0, 'tcp://0.0.0.0:4567') turning message relay requesting off
2018-02-01 16:44:01 140249296508672 [Note] WSREP: 1.0 (mariadb-galera-0): State transfer to 0.0 (mariadb-galera-1) complete.
2018-02-01 16:44:01 140249296508672 [Note] WSREP: Member 1.0 (mariadb-galera-0) synced with group.
WSREP_SST: [INFO] Joiner cleanup. rsync PID: 142 (20180201 16:44:01.650)
WSREP_SST: [INFO] Joiner cleanup done. (20180201 16:44:02.159)
2018-02-01 16:44:02 140249689663360 [Note] WSREP: SST complete, seqno: 0
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: For Galera, using innodb_lock_schedule_algorithm=fcfs
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Uses event mutexes
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Compressed tables use zlib 1.2.8
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Using Linux native AIO
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Number of pools: 1
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Using SSE2 crc32 instructions
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Initializing buffer pool, total size = 256M, instances = 1, chunk size = 128M
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Completed initialization of buffer pool
2018-02-01 16:44:02 140248358000384 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Highest supported file format is Barracuda.
2018-02-01 16:44:02 140249689663360 [Note] InnoDB: Starting crash recovery from checkpoint LSN=1619742
2018-02-01 16:44:03 140249689663360 [Note] InnoDB: 128 out of 128 rollback segments are active.
2018-02-01 16:44:03 140249689663360 [Note] InnoDB: Creating shared tablespace for temporary tables
2018-02-01 16:44:03 140249689663360 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
2018-02-01 16:44:03 140249689663360 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
2018-02-01 16:44:03 140249689663360 [Note] InnoDB: Waiting for purge to start
2018-02-01 16:44:03 140249689663360 [Note] InnoDB: 5.7.20 started; log sequence number 1619751
2018-02-01 16:44:03 140248349607680 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
2018-02-01 16:44:03 140248349607680 [Note] InnoDB: Buffer pool(s) load completed at 180201 16:44:03
2018-02-01 16:44:03 140249689663360 [Note] Plugin 'FEEDBACK' is disabled.
2018-02-01 16:44:03 140249689663360 [Note] Server socket created on IP: '0.0.0.0'.
2018-02-01 16:44:03 140249689663360 [Warning] 'proxies_priv' entry '@% root@mariadb-galera-0' ignored in --skip-name-resolve mode.
2018-02-01 16:44:03 140249689663360 [Note] WSREP: Signalling provider to continue.
2018-02-01 16:44:03 140249689663360 [Note] WSREP: SST received: f69e5fe3-076e-11e8-83e0-6646cc4cd744:0
2018-02-01 16:44:03 140249296508672 [Note] WSREP: 0.0 (mariadb-galera-1): State transfer from 1.0 (mariadb-galera-0) complete.
2018-02-01 16:44:03 140249296508672 [Note] WSREP: Shifting JOINER -> JOINED (TO: 0)
2018-02-01 16:44:03 140249689663360 [Note] Reading of all Master_info entries succeded
2018-02-01 16:44:03 140249689663360 [Note] Added new Master_info '' to hash table
2018-02-01 16:44:03 140249689663360 [Note] mysqld: ready for connections.
Version: '10.2.12-MariaDB-10.2.12+maria~jessie'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  mariadb.org binary distribution
2018-02-01 16:44:03 140249296508672 [Note] WSREP: Member 0.0 (mariadb-galera-1) synced with group.
2018-02-01 16:44:03 140249296508672 [Note] WSREP: Shifting JOINED -> SYNCED (TO: 0)
2018-02-01 16:44:03 140249688164096 [Note] WSREP: Synchronized with group, ready for connections
2018-02-01 16:44:03 140249688164096 [Note] WSREP: wsrep_notify_cmd is not defined, skipping notification.
2018-02-01 16:44:39 140249304901376 [Note] WSREP: (16f980a0, 'tcp://0.0.0.0:4567') turning message relay requesting on, nonlive peers: tcp://10.244.3.254:4567 
2018-02-01 16:44:39 140249304901376 [Note] WSREP: (16f980a0, 'tcp://0.0.0.0:4567') connection established to 31b1834e tcp://10.244.3.254:4567
2018-02-01 16:44:39 140249304901376 [Note] WSREP: declaring 31b1834e at tcp://10.244.3.254:4567 stable
2018-02-01 16:44:39 140249304901376 [Note] WSREP: declaring f69d2c57 at tcp://10.244.3.253:4567 stable
2018-02-01 16:44:39 140249304901376 [Note] WSREP: Node 16f980a0 state prim
2018-02-01 16:44:39 140249304901376 [Note] WSREP: view(view_id(PRIM,16f980a0,3) memb {
	16f980a0,0
	31b1834e,0
	f69d2c57,0
} joined {
} left {
} partitioned {
})
2018-02-01 16:44:39 140249304901376 [Note] WSREP: save pc into disk
2018-02-01 16:44:39 140249296508672 [Note] WSREP: New COMPONENT: primary = yes, bootstrap = no, my_idx = 0, memb_num = 3
2018-02-01 16:44:39 140249296508672 [Note] WSREP: STATE_EXCHANGE: sent state UUID: 321ee830-076f-11e8-aec5-6e1b23326a0e
2018-02-01 16:44:39 140249296508672 [Note] WSREP: STATE EXCHANGE: sent state msg: 321ee830-076f-11e8-aec5-6e1b23326a0e
2018-02-01 16:44:39 140249296508672 [Note] WSREP: STATE EXCHANGE: got state msg: 321ee830-076f-11e8-aec5-6e1b23326a0e from 0 (mariadb-galera-1)
2018-02-01 16:44:39 140249296508672 [Note] WSREP: STATE EXCHANGE: got state msg: 321ee830-076f-11e8-aec5-6e1b23326a0e from 2 (mariadb-galera-0)
2018-02-01 16:44:39 140249296508672 [Note] WSREP: STATE EXCHANGE: got state msg: 321ee830-076f-11e8-aec5-6e1b23326a0e from 1 (mariadb-galera-2)
2018-02-01 16:44:39 140249296508672 [Note] WSREP: Quorum results:
	version    = 4,
	component  = PRIMARY,
	conf_id    = 2,
	members    = 2/3 (joined/total),
	act_id     = 0,
	last_appl. = 0,
	protocols  = 0/7/3 (gcs/repl/appl),
	group UUID = f69e5fe3-076e-11e8-83e0-6646cc4cd744
2018-02-01 16:44:39 140249296508672 [Note] WSREP: Flow-control interval: [28, 28]
2018-02-01 16:44:39 140249296508672 [Note] WSREP: Trying to continue unpaused monitor
2018-02-01 16:44:39 140249688164096 [Note] WSREP: New cluster view: global state: f69e5fe3-076e-11e8-83e0-6646cc4cd744:0, view# 3: Primary, number of nodes: 3, my index: 0, protocol version 3
2018-02-01 16:44:39 140249688164096 [Note] WSREP: wsrep_notify_cmd is not defined, skipping notification.
2018-02-01 16:44:39 140249688164096 [Note] WSREP: REPL Protocols: 7 (3, 2)
2018-02-01 16:44:39 140249688164096 [Note] WSREP: Assign initial position for certification: 0, protocol version: 3
2018-02-01 16:44:39 140249338472192 [Note] WSREP: Service thread queue flushed.
2018-02-01 16:44:40 140249296508672 [Note] WSREP: Member 1.0 (mariadb-galera-2) requested state transfer from '*any*'. Selected 0.0 (mariadb-galera-1)(SYNCED) as donor.
2018-02-01 16:44:40 140249296508672 [Note] WSREP: Shifting SYNCED -> DONOR/DESYNCED (TO: 0)
2018-02-01 16:44:40 140249688164096 [Note] WSREP: wsrep_notify_cmd is not defined, skipping notification.
2018-02-01 16:44:40 140248026961664 [Note] WSREP: Running: 'wsrep_sst_rsync --role 'donor' --address '10.244.3.254:4444/rsync_sst' --socket '/var/run/mysqld/mysqld.sock' --datadir '/var/lib/mysql/'     '' --gtid 'f69e5fe3-076e-11e8-83e0-6646cc4cd744:0' --gtid-domain-id '0''
2018-02-01 16:44:40 140249688164096 [Note] WSREP: sst_donor_thread signaled with 0
2018-02-01 16:44:40 140248026961664 [Note] WSREP: Flushing tables for SST...
2018-02-01 16:44:40 140248026961664 [Note] WSREP: Provider paused at f69e5fe3-076e-11e8-83e0-6646cc4cd744:0 (7)
2018-02-01 16:44:40 140248026961664 [Note] WSREP: Tables flushed.
2018-02-01 16:44:42 140249304901376 [Note] WSREP: (16f980a0, 'tcp://0.0.0.0:4567') turning message relay requesting off
2018-02-01 16:44:45 140248026961664 [Note] WSREP: resuming provider at 7
2018-02-01 16:44:45 140248026961664 [Note] WSREP: Provider resumed.
2018-02-01 16:44:45 140249296508672 [Note] WSREP: 0.0 (mariadb-galera-1): State transfer to 1.0 (mariadb-galera-2) complete.
2018-02-01 16:44:45 140249296508672 [Note] WSREP: Shifting DONOR/DESYNCED -> JOINED (TO: 0)
2018-02-01 16:44:45 140249296508672 [Note] WSREP: Member 0.0 (mariadb-galera-1) synced with group.
2018-02-01 16:44:45 140249296508672 [Note] WSREP: Shifting JOINED -> SYNCED (TO: 0)
2018-02-01 16:44:45 140249688164096 [Note] WSREP: Synchronized with group, ready for connections
2018-02-01 16:44:45 140249688164096 [Note] WSREP: wsrep_notify_cmd is not defined, skipping notification.
2018-02-01 16:44:47 140249296508672 [Note] WSREP: 1.0 (mariadb-galera-2): State transfer from 0.0 (mariadb-galera-1) complete.
2018-02-01 16:44:47 140249296508672 [Note] WSREP: Member 1.0 (mariadb-galera-2) synced with group.
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl get pvc -l app=mariadb
NAME                        STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
hostpath-mariadb-galera-0   Bound     pvc-4c66444f-06db-11e8-beb1-525400224e72   50Mi       RWX            example-hostpath   7m
hostpath-mariadb-galera-1   Bound     pvc-1f1252a8-06eb-11e8-beb1-525400224e72   50Mi       RWX            example-hostpath   5m
hostpath-mariadb-galera-2   Bound     pvc-d1eb5f7b-06eb-11e8-beb1-525400224e72   50Mi       RWX            example-hostpath   4m
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pods -l app=mariadb -o wide
NAME               READY     STATUS    RESTARTS   AGE       IP             NODE
mariadb-galera-0   1/1       Running   0          7m        10.244.3.253   rookdev-172-17-4-63
mariadb-galera-1   1/1       Running   0          6m        10.244.2.225   rookdev-172-17-4-61
mariadb-galera-2   1/1       Running   0          5m        10.244.3.254   rookdev-172-17-4-63
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl logs mariadb-galera-2
2018-02-01 16:44:39 140445071054720 [Note] mysqld (mysqld 10.2.12-MariaDB-10.2.12+maria~jessie) starting as process 1 ...
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Read nil XID from storage engines, skipping position init
2018-02-01 16:44:39 140445071054720 [Note] WSREP: wsrep_load(): loading provider library '/usr/lib/libgalera_smm.so'
2018-02-01 16:44:39 140445071054720 [Note] WSREP: wsrep_load(): Galera 25.3.22(r3764) by Codership Oy <info@codership.com> loaded successfully.
2018-02-01 16:44:39 140445071054720 [Note] WSREP: CRC-32C: using hardware acceleration.
2018-02-01 16:44:39 140445071054720 [Warning] WSREP: Could not open state file for reading: '/var/lib/mysql//grastate.dat'
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Found saved state: 00000000-0000-0000-0000-000000000000:-1, safe_to_bootstrap: 1
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Passing config to GCS: base_dir = /var/lib/mysql/; base_host = 10.244.3.254; base_port = 4567; cert.log_conflicts = no; debug = no; evs.auto_evict = 0; evs.delay_margin = PT1S; evs.delayed_keep_period = PT30S; evs.inactive_check_period = PT0.5S; evs.inactive_timeout = PT15S; evs.join_retrans_period = PT1S; evs.max_install_timeouts = 3; evs.send_window = 4; evs.stats_report_period = PT1M; evs.suspect_timeout = PT5S; evs.user_send_window = 2; evs.view_forget_timeout = PT24H; gcache.dir = /var/lib/mysql/; gcache.keep_pages_size = 0; gcache.mem_size = 0; gcache.name = /var/lib/mysql//galera.cache; gcache.page_size = 128M; gcache.recover = no; gcache.size = 128M; gcomm.thread_prio = ; gcs.fc_debug = 0; gcs.fc_factor = 1.0; gcs.fc_limit = 16; gcs.fc_master_slave = no; gcs.max_packet_size = 64500; gcs.max_throttle = 0.25; gcs.recv_q_hard_limit = 9223372036854775807; gcs.recv_q_soft_limit = 0.25; gcs.sync_donor = no; gmcast.segment = 0; gmcast.version = 0; pc.announce_timeout = PT3S; pc.checksum = false; pc.i
2018-02-01 16:44:39 140445071054720 [Note] WSREP: GCache history reset: 00000000-0000-0000-0000-000000000000:0 -> 00000000-0000-0000-0000-000000000000:-1
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Assign initial position for certification: -1, protocol version: -1
2018-02-01 16:44:39 140445071054720 [Note] WSREP: wsrep_sst_grab()
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Start replication
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Setting initial position to 00000000-0000-0000-0000-000000000000:-1
2018-02-01 16:44:39 140445071054720 [Note] WSREP: protonet asio version 0
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Using CRC-32C for message checksums.
2018-02-01 16:44:39 140445071054720 [Note] WSREP: backend: asio
2018-02-01 16:44:39 140445071054720 [Note] WSREP: gcomm thread scheduling priority set to other:0 
2018-02-01 16:44:39 140445071054720 [Warning] WSREP: access file(/var/lib/mysql//gvwstate.dat) failed(No such file or directory)
2018-02-01 16:44:39 140445071054720 [Note] WSREP: restore pc from disk failed
2018-02-01 16:44:39 140445071054720 [Note] WSREP: GMCast version 0
2018-02-01 16:44:39 140445071054720 [Note] WSREP: (31b1834e, 'tcp://0.0.0.0:4567') listening at tcp://0.0.0.0:4567
2018-02-01 16:44:39 140445071054720 [Note] WSREP: (31b1834e, 'tcp://0.0.0.0:4567') multicast: , ttl: 1
2018-02-01 16:44:39 140445071054720 [Note] WSREP: EVS version 0
2018-02-01 16:44:39 140445071054720 [Note] WSREP: gcomm: connecting to group 'mariadb-galera', peer '10.244.3.253:'
2018-02-01 16:44:39 140445071054720 [Note] WSREP: (31b1834e, 'tcp://0.0.0.0:4567') connection established to f69d2c57 tcp://10.244.3.253:4567
2018-02-01 16:44:39 140445071054720 [Note] WSREP: (31b1834e, 'tcp://0.0.0.0:4567') turning message relay requesting on, nonlive peers: tcp://10.244.2.225:4567 
2018-02-01 16:44:39 140445071054720 [Note] WSREP: (31b1834e, 'tcp://0.0.0.0:4567') connection established to 16f980a0 tcp://10.244.2.225:4567
2018-02-01 16:44:39 140445071054720 [Note] WSREP: declaring 16f980a0 at tcp://10.244.2.225:4567 stable
2018-02-01 16:44:39 140445071054720 [Note] WSREP: declaring f69d2c57 at tcp://10.244.3.253:4567 stable
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Node 16f980a0 state prim
2018-02-01 16:44:39 140445071054720 [Note] WSREP: view(view_id(PRIM,16f980a0,3) memb {
	16f980a0,0
	31b1834e,0
	f69d2c57,0
} joined {
} left {
} partitioned {
})
2018-02-01 16:44:39 140445071054720 [Note] WSREP: save pc into disk
2018-02-01 16:44:39 140445071054720 [Note] WSREP: gcomm: connected
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Changing maximum packet size to 64500, resulting msg size: 32636
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Shifting CLOSED -> OPEN (TO: 0)
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Opened channel 'mariadb-galera'
2018-02-01 16:44:39 140445071054720 [Note] WSREP: Waiting for SST to complete.
2018-02-01 16:44:39 140444675589888 [Note] WSREP: New COMPONENT: primary = yes, bootstrap = no, my_idx = 1, memb_num = 3
2018-02-01 16:44:39 140444675589888 [Note] WSREP: STATE EXCHANGE: Waiting for state UUID.
2018-02-01 16:44:39 140444675589888 [Note] WSREP: STATE EXCHANGE: sent state msg: 321ee830-076f-11e8-aec5-6e1b23326a0e
2018-02-01 16:44:39 140444675589888 [Note] WSREP: STATE EXCHANGE: got state msg: 321ee830-076f-11e8-aec5-6e1b23326a0e from 0 (mariadb-galera-1)
2018-02-01 16:44:39 140444675589888 [Note] WSREP: STATE EXCHANGE: got state msg: 321ee830-076f-11e8-aec5-6e1b23326a0e from 2 (mariadb-galera-0)
2018-02-01 16:44:39 140444675589888 [Note] WSREP: STATE EXCHANGE: got state msg: 321ee830-076f-11e8-aec5-6e1b23326a0e from 1 (mariadb-galera-2)
2018-02-01 16:44:39 140444675589888 [Note] WSREP: Quorum results:
	version    = 4,
	component  = PRIMARY,
	conf_id    = 2,
	members    = 2/3 (joined/total),
	act_id     = 0,
	last_appl. = -1,
	protocols  = 0/7/3 (gcs/repl/appl),
	group UUID = f69e5fe3-076e-11e8-83e0-6646cc4cd744
2018-02-01 16:44:39 140444675589888 [Note] WSREP: Flow-control interval: [28, 28]
2018-02-01 16:44:39 140444675589888 [Note] WSREP: Trying to continue unpaused monitor
2018-02-01 16:44:39 140444675589888 [Note] WSREP: Shifting OPEN -> PRIMARY (TO: 0)
2018-02-01 16:44:39 140445069555456 [Note] WSREP: State transfer required: 
	Group state: f69e5fe3-076e-11e8-83e0-6646cc4cd744:0
	Local state: 00000000-0000-0000-0000-000000000000:-1
2018-02-01 16:44:39 140445069555456 [Note] WSREP: New cluster view: global state: f69e5fe3-076e-11e8-83e0-6646cc4cd744:0, view# 3: Primary, number of nodes: 3, my index: 1, protocol version 3
2018-02-01 16:44:39 140445069555456 [Warning] WSREP: Gap in state sequence. Need state transfer.
2018-02-01 16:44:39 140444667197184 [Note] WSREP: Running: 'wsrep_sst_rsync --role 'joiner' --address '10.244.3.254' --datadir '/var/lib/mysql/'   --parent '1'  '' '
2018-02-01 16:44:39 140445069555456 [Note] WSREP: Prepared SST request: rsync|10.244.3.254:4444/rsync_sst
2018-02-01 16:44:39 140445069555456 [Note] WSREP: wsrep_notify_cmd is not defined, skipping notification.
2018-02-01 16:44:39 140445069555456 [Note] WSREP: REPL Protocols: 7 (3, 2)
2018-02-01 16:44:39 140445069555456 [Note] WSREP: Assign initial position for certification: 0, protocol version: 3
2018-02-01 16:44:39 140444925593344 [Note] WSREP: Service thread queue flushed.
2018-02-01 16:44:39 140445069555456 [Warning] WSREP: Failed to prepare for incremental state transfer: Local state UUID (00000000-0000-0000-0000-000000000000) does not match group state UUID (f69e5fe3-076e-11e8-83e0-6646cc4cd744): 1 (Operation not permitted)
	 at galera/src/replicator_str.cpp:prepare_for_IST():482. IST will be unavailable.
2018-02-01 16:44:39 140444675589888 [Note] WSREP: Member 1.0 (mariadb-galera-2) requested state transfer from '*any*'. Selected 0.0 (mariadb-galera-1)(SYNCED) as donor.
2018-02-01 16:44:39 140444675589888 [Note] WSREP: Shifting PRIMARY -> JOINER (TO: 0)
2018-02-01 16:44:39 140445069555456 [Note] WSREP: Requesting state transfer: success, donor: 0
2018-02-01 16:44:39 140445069555456 [Note] WSREP: GCache history reset: 00000000-0000-0000-0000-000000000000:0 -> f69e5fe3-076e-11e8-83e0-6646cc4cd744:0
2018-02-01 16:44:42 140444683982592 [Note] WSREP: (31b1834e, 'tcp://0.0.0.0:4567') turning message relay requesting off
2018-02-01 16:44:45 140444675589888 [Note] WSREP: 0.0 (mariadb-galera-1): State transfer to 1.0 (mariadb-galera-2) complete.
2018-02-01 16:44:45 140444675589888 [Note] WSREP: Member 0.0 (mariadb-galera-1) synced with group.
WSREP_SST: [INFO] Joiner cleanup. rsync PID: 142 (20180201 16:44:45.956)
WSREP_SST: [INFO] Joiner cleanup done. (20180201 16:44:46.462)
2018-02-01 16:44:46 140445071054720 [Note] WSREP: SST complete, seqno: 0
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: For Galera, using innodb_lock_schedule_algorithm=fcfs
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Uses event mutexes
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Compressed tables use zlib 1.2.8
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Using Linux native AIO
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Number of pools: 1
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Using SSE2 crc32 instructions
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Initializing buffer pool, total size = 256M, instances = 1, chunk size = 128M
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Completed initialization of buffer pool
2018-02-01 16:44:46 140443737081600 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Highest supported file format is Barracuda.
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Starting crash recovery from checkpoint LSN=1619779
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: 128 out of 128 rollback segments are active.
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Creating shared tablespace for temporary tables
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: Waiting for purge to start
2018-02-01 16:44:46 140445071054720 [Note] InnoDB: 5.7.20 started; log sequence number 1619788
2018-02-01 16:44:46 140443728688896 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
2018-02-01 16:44:46 140443728688896 [Note] InnoDB: Buffer pool(s) load completed at 180201 16:44:46
2018-02-01 16:44:46 140445071054720 [Note] Plugin 'FEEDBACK' is disabled.
2018-02-01 16:44:46 140445071054720 [Note] Server socket created on IP: '0.0.0.0'.
2018-02-01 16:44:46 140445071054720 [Warning] 'proxies_priv' entry '@% root@mariadb-galera-0' ignored in --skip-name-resolve mode.
2018-02-01 16:44:46 140445071054720 [Note] WSREP: Signalling provider to continue.
2018-02-01 16:44:46 140445071054720 [Note] WSREP: SST received: f69e5fe3-076e-11e8-83e0-6646cc4cd744:0
2018-02-01 16:44:46 140445071054720 [Note] Reading of all Master_info entries succeded
2018-02-01 16:44:46 140445071054720 [Note] Added new Master_info '' to hash table
2018-02-01 16:44:46 140445071054720 [Note] mysqld: ready for connections.
Version: '10.2.12-MariaDB-10.2.12+maria~jessie'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  mariadb.org binary distribution
2018-02-01 16:44:46 140444675589888 [Note] WSREP: 1.0 (mariadb-galera-2): State transfer from 0.0 (mariadb-galera-1) complete.
2018-02-01 16:44:46 140444675589888 [Note] WSREP: Shifting JOINER -> JOINED (TO: 0)
2018-02-01 16:44:46 140444675589888 [Note] WSREP: Member 1.0 (mariadb-galera-2) synced with group.
2018-02-01 16:44:46 140444675589888 [Note] WSREP: Shifting JOINED -> SYNCED (TO: 0)
2018-02-01 16:44:46 140445069555456 [Note] WSREP: Synchronized with group, ready for connections
2018-02-01 16:44:46 140445069555456 [Note] WSREP: wsrep_notify_cmd is not defined, skipping notification.
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti mariadb-galera-2 -- mysql --user=testuser --password=testpass -e "SHOW STATUS LIKE 'wsrep_%';"
+------------------------------+-------------------------------------------------------+
| Variable_name                | Value                                                 |
+------------------------------+-------------------------------------------------------+
| wsrep_apply_oooe             | 0.000000                                              |
| wsrep_apply_oool             | 0.000000                                              |
| wsrep_apply_window           | 0.000000                                              |
| wsrep_causal_reads           | 0                                                     |
| wsrep_cert_deps_distance     | 0.000000                                              |
| wsrep_cert_index_size        | 0                                                     |
| wsrep_cert_interval          | 0.000000                                              |
| wsrep_cluster_conf_id        | 3                                                     |
| wsrep_cluster_size           | 3                                                     |
| wsrep_cluster_state_uuid     | f69e5fe3-076e-11e8-83e0-6646cc4cd744                  |
| wsrep_cluster_status         | Primary                                               |
| wsrep_commit_oooe            | 0.000000                                              |
| wsrep_commit_oool            | 0.000000                                              |
| wsrep_commit_window          | 0.000000                                              |
| wsrep_connected              | ON                                                    |
| wsrep_desync_count           | 0                                                     |
| wsrep_evs_delayed            |                                                       |
| wsrep_evs_evict_list         |                                                       |
| wsrep_evs_repl_latency       | 0/0/0/0/0                                             |
| wsrep_evs_state              | OPERATIONAL                                           |
| wsrep_flow_control_paused    | 0.000000                                              |
| wsrep_flow_control_paused_ns | 0                                                     |
| wsrep_flow_control_recv      | 0                                                     |
| wsrep_flow_control_sent      | 0                                                     |
| wsrep_gcomm_uuid             | 31b1834e-076f-11e8-a108-fe33ec03b5f6                  |
| wsrep_incoming_addresses     | 10.244.2.225:3306,10.244.3.254:3306,10.244.3.253:3306 |
| wsrep_last_committed         | 0                                                     |
| wsrep_local_bf_aborts        | 0                                                     |
| wsrep_local_cached_downto    | 18446744073709551615                                  |
| wsrep_local_cert_failures    | 0                                                     |
| wsrep_local_commits          | 0                                                     |
| wsrep_local_index            | 1                                                     |
| wsrep_local_recv_queue       | 0                                                     |
| wsrep_local_recv_queue_avg   | 0.000000                                              |
| wsrep_local_recv_queue_max   | 1                                                     |
| wsrep_local_recv_queue_min   | 0                                                     |
| wsrep_local_replays          | 0                                                     |
| wsrep_local_send_queue       | 0                                                     |
| wsrep_local_send_queue_avg   | 0.000000                                              |
| wsrep_local_send_queue_max   | 1                                                     |
| wsrep_local_send_queue_min   | 0                                                     |
| wsrep_local_state            | 4                                                     |
| wsrep_local_state_comment    | Synced                                                |
| wsrep_local_state_uuid       | f69e5fe3-076e-11e8-83e0-6646cc4cd744                  |
| wsrep_protocol_version       | 7                                                     |
| wsrep_provider_name          | Galera                                                |
| wsrep_provider_vendor        | Codership Oy <info@codership.com>                     |
| wsrep_provider_version       | 25.3.22(r3764)                                        |
| wsrep_ready                  | ON                                                    |
| wsrep_received               | 3                                                     |
| wsrep_received_bytes         | 320                                                   |
| wsrep_repl_data_bytes        | 0                                                     |
| wsrep_repl_keys              | 0                                                     |
| wsrep_repl_keys_bytes        | 0                                                     |
| wsrep_repl_other_bytes       | 0                                                     |
| wsrep_replicated             | 0                                                     |
| wsrep_replicated_bytes       | 0                                                     |
| wsrep_thread_count           | 2                                                     |
+------------------------------+-------------------------------------------------------+
```

__Inspect__

When init failed
```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pods -l app=mariadb
NAME               READY     STATUS                  RESTARTS   AGE
mariadb-galera-0   0/1       Init:CrashLoopBackOff   3          1m
```

```
[vagrant@rookdev-172-17-4-63 ~]$ docker ps -a -f label="io.kubernetes.pod.name"="mariadb-galera-0"
CONTAINER ID        IMAGE                                                                                                           COMMAND                  CREATED              STATUS                        PORTS               NAMES
2652844a7741        docker.io/tangfeixiong/mysql-operator@sha256:367926b68253d259d306374a8aca6c8cb8a1d568ab57b017b212253bc709bd8d   "/mysql-operator i..."   About a minute ago   Exited (255) 59 seconds ago                       k8s_initcnf_mariadb-galera-0_default_4c6f0eaf-06e0-11e8-beb1-525400224e72_7
0b2afa1f7a82        gcr.io/google_containers/pause-amd64:3.0                                                                        "/pause"                 11 minutes ago       Up 11 minutes                                     k8s_POD_mariadb-galera-0_default_4c6f0eaf-06e0-11e8-beb1-525400224e72_0
```

```
[vagrant@rookdev-172-17-4-63 ~]$ docker logs 2652844a7741
```

### Investigate initContainer

StatefulSet required
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ mysql-operator init --kubeconfig --logtostderr -v 5 --conf_dir=$HOME --name=my-galera
F0131 09:46:04.208810   19922 cnf.go:99] Get statefulset failed: statefulsets.apps "my-galera" not found
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ bin/mysql-operator init --kubeconfig --logtostderr -v 5 --conf_dir=$HOME --name=mariadb-galera
I0201 09:51:55.588256    8215 cnf.go:59] With init parameters: &{/home/vagrant/.kube/config  mariadb-galera  galera 3306 /home/vagrant 0}
E0201 09:51:55.737256    8215 cnf.go:135] Get POD name failed or Environment POD name not found
E0201 09:51:55.737717    8215 cnf.go:148] Get POD IP failed or environment POD IP not found
[mysql-operator] 2018/02/01 09:51:55 cnf.go:138: Wrote /home/vagrant/galera.cnf
I0201 09:51:55.745934    8215 cnf.go:157] By bootstrap ip: 10.244.3.244
```


