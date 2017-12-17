# Instruction

## standalone

For kubectl
```
[vagrant@localhost go-to-kubernetes]$ kubectl get nodes
NAME          STATUS    ROLES     AGE       VERSION
172.17.4.50   Ready     <none>    188d      v1.7.4
```

### rook

_Note: disable RBAC in yaml manifest_

Deploy rook
```
[vagrant@localhost ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-operator-no-rbac.yaml 
namespace "rook-system" created
serviceaccount "rook-operator" created
clusterrolebinding "rook-operator" created
deployment "rook-operator" created
Error from server (Forbidden): error when creating "rook0x2Fcluster/examples/kubernetes/rook-operator-no-rbac.yaml": clusterroles.rbac.authorization.k8s.io "rook-operator" is forbidden: attempt to grant extra privileges: [PolicyRule{Resources:["namespaces"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["namespaces"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["namespaces"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["namespaces"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["namespaces"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["namespaces"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["namespaces"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["serviceaccounts"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["serviceaccounts"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["serviceaccounts"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["serviceaccounts"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["serviceaccounts"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["serviceaccounts"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["serviceaccounts"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["secrets"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["secrets"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["secrets"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["secrets"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["secrets"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["secrets"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["secrets"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["pods"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["pods"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["pods"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["pods"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["pods"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["pods"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["pods"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["services"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["services"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["services"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["services"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["services"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["services"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["services"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["nodes"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["nodes"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["nodes"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["nodes"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["nodes"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["nodes"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["nodes"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["nodes/proxy"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["nodes/proxy"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["nodes/proxy"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["nodes/proxy"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["nodes/proxy"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["nodes/proxy"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["nodes/proxy"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["configmaps"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["configmaps"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["configmaps"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["configmaps"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["configmaps"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["configmaps"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["configmaps"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["events"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["events"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["events"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["events"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["events"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["events"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["events"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["persistentvolumes"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["persistentvolumes"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["persistentvolumes"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["persistentvolumes"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["persistentvolumes"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["persistentvolumes"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["persistentvolumes"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["persistentvolumeclaims"], APIGroups:[""], Verbs:["get"]} PolicyRule{Resources:["persistentvolumeclaims"], APIGroups:[""], Verbs:["list"]} PolicyRule{Resources:["persistentvolumeclaims"], APIGroups:[""], Verbs:["watch"]} PolicyRule{Resources:["persistentvolumeclaims"], APIGroups:[""], Verbs:["patch"]} PolicyRule{Resources:["persistentvolumeclaims"], APIGroups:[""], Verbs:["create"]} PolicyRule{Resources:["persistentvolumeclaims"], APIGroups:[""], Verbs:["update"]} PolicyRule{Resources:["persistentvolumeclaims"], APIGroups:[""], Verbs:["delete"]} PolicyRule{Resources:["thirdpartyresources"], APIGroups:["extensions"], Verbs:["get"]} PolicyRule{Resources:["thirdpartyresources"], APIGroups:["extensions"], Verbs:["list"]} PolicyRule{Resources:["thirdpartyresources"], APIGroups:["extensions"], Verbs:["watch"]} PolicyRule{Resources:["thirdpartyresources"], APIGroups:["extensions"], Verbs:["create"]} PolicyRule{Resources:["thirdpartyresources"], APIGroups:["extensions"], Verbs:["delete"]} PolicyRule{Resources:["deployments"], APIGroups:["extensions"], Verbs:["get"]} PolicyRule{Resources:["deployments"], APIGroups:["extensions"], Verbs:["list"]} PolicyRule{Resources:["deployments"], APIGroups:["extensions"], Verbs:["watch"]} PolicyRule{Resources:["deployments"], APIGroups:["extensions"], Verbs:["create"]} PolicyRule{Resources:["deployments"], APIGroups:["extensions"], Verbs:["delete"]} PolicyRule{Resources:["daemonsets"], APIGroups:["extensions"], Verbs:["get"]} PolicyRule{Resources:["daemonsets"], APIGroups:["extensions"], Verbs:["list"]} PolicyRule{Resources:["daemonsets"], APIGroups:["extensions"], Verbs:["watch"]} PolicyRule{Resources:["daemonsets"], APIGroups:["extensions"], Verbs:["create"]} PolicyRule{Resources:["daemonsets"], APIGroups:["extensions"], Verbs:["delete"]} PolicyRule{Resources:["replicasets"], APIGroups:["extensions"], Verbs:["get"]} PolicyRule{Resources:["replicasets"], APIGroups:["extensions"], Verbs:["list"]} PolicyRule{Resources:["replicasets"], APIGroups:["extensions"], Verbs:["watch"]} PolicyRule{Resources:["replicasets"], APIGroups:["extensions"], Verbs:["create"]} PolicyRule{Resources:["replicasets"], APIGroups:["extensions"], Verbs:["delete"]} PolicyRule{Resources:["customresourcedefinitions"], APIGroups:["apiextensions.k8s.io"], Verbs:["get"]} PolicyRule{Resources:["customresourcedefinitions"], APIGroups:["apiextensions.k8s.io"], Verbs:["list"]} PolicyRule{Resources:["customresourcedefinitions"], APIGroups:["apiextensions.k8s.io"], Verbs:["watch"]} PolicyRule{Resources:["customresourcedefinitions"], APIGroups:["apiextensions.k8s.io"], Verbs:["create"]} PolicyRule{Resources:["customresourcedefinitions"], APIGroups:["apiextensions.k8s.io"], Verbs:["delete"]} PolicyRule{Resources:["clusterroles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["get"]} PolicyRule{Resources:["clusterroles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["list"]} PolicyRule{Resources:["clusterroles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["watch"]} PolicyRule{Resources:["clusterroles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["create"]} PolicyRule{Resources:["clusterroles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["update"]} PolicyRule{Resources:["clusterroles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["delete"]} PolicyRule{Resources:["clusterrolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["get"]} PolicyRule{Resources:["clusterrolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["list"]} PolicyRule{Resources:["clusterrolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["watch"]} PolicyRule{Resources:["clusterrolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["create"]} PolicyRule{Resources:["clusterrolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["update"]} PolicyRule{Resources:["clusterrolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["delete"]} PolicyRule{Resources:["roles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["get"]} PolicyRule{Resources:["roles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["list"]} PolicyRule{Resources:["roles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["watch"]} PolicyRule{Resources:["roles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["create"]} PolicyRule{Resources:["roles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["update"]} PolicyRule{Resources:["roles"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["delete"]} PolicyRule{Resources:["rolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["get"]} PolicyRule{Resources:["rolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["list"]} PolicyRule{Resources:["rolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["watch"]} PolicyRule{Resources:["rolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["create"]} PolicyRule{Resources:["rolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["update"]} PolicyRule{Resources:["rolebindings"], APIGroups:["rbac.authorization.k8s.io"], Verbs:["delete"]} PolicyRule{Resources:["storageclasses"], APIGroups:["storage.k8s.io"], Verbs:["get"]} PolicyRule{Resources:["storageclasses"], APIGroups:["storage.k8s.io"], Verbs:["list"]} PolicyRule{Resources:["storageclasses"], APIGroups:["storage.k8s.io"], Verbs:["watch"]} PolicyRule{Resources:["storageclasses"], APIGroups:["storage.k8s.io"], Verbs:["delete"]} PolicyRule{Resources:["*"], APIGroups:["rook.io"], Verbs:["*"]}] user=&{kubecfg  [system:authenticated] map[]} ownerrules=[] ruleResolutionErrors=[]
```

```
[vagrant@localhost ~]$ kubectl -n rook-system get pods
NAME                             READY     STATUS    RESTARTS   AGE
rook-agent-q6jzw                 1/1       Running   0          7m
rook-operator-1569969384-knrw2   1/1       Running   0          7m
```

Deploy cluster
```
[vagrant@localhost ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-cluster.yaml 
namespace "rook" created
cluster "rook" created
```

```
[vagrant@localhost ~]$ kubectl -n rook get pods
NAME                              READY     STATUS    RESTARTS   AGE
rook-api-877605387-ng72q          1/1       Running   0          16m
rook-ceph-mgr0-2792268363-xnl49   1/1       Running   0          16m
rook-ceph-mgr1-3797254331-rmgzf   1/1       Running   0          16m
rook-ceph-mon0-zrzt6              1/1       Running   0          17m
rook-ceph-mon1-62064              1/1       Running   0          17m
rook-ceph-mon2-2w720              1/1       Running   0          17m
rook-ceph-osd-bph19               1/1       Running   0          16m
```

Deploy cephfs
```
[vagrant@localhost ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-filesystem.yaml 
filesystem "myfs" created
```

```
[vagrant@localhost ~]$ kubectl -n rook get pods
NAME                                  READY     STATUS    RESTARTS   AGE
rook-api-877605387-ng72q              1/1       Running   0          45m
rook-ceph-mds-myfs-3914257510-6cxvv   1/1       Running   0          54s
rook-ceph-mds-myfs-3914257510-d55nd   1/1       Running   0          54s
rook-ceph-mgr0-2792268363-xnl49       1/1       Running   0          45m
rook-ceph-mgr1-3797254331-rmgzf       1/1       Running   0          45m
rook-ceph-mon0-zrzt6                  1/1       Running   0          45m
rook-ceph-mon1-62064                  1/1       Running   0          45m
rook-ceph-mon2-2w720                  1/1       Running   0          45m
rook-ceph-osd-bph19                   1/1       Running   0          45m
```

Tools
```
[vagrant@localhost ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-tools.yaml 
pod "rook-tools" created
```

```
[vagrant@localhost ~]$ kubectl -n rook exec -ti rook-tools ceph status
  cluster:
    id:     5ef33a3d-d4cc-46dc-b6c0-383a9a315dd2
    health: HEALTH_OK
 
  services:
    mon: 3 daemons, quorum rook-ceph-mon0,rook-ceph-mon2,rook-ceph-mon1
    mgr: rook-ceph-mgr0(active), standbys: rook-ceph-mgr1
    mds: myfs-1/1/1 up  {0=md55nd=up:active}, 1 up:standby-replay
    osd: 1 osds: 1 up, 1 in
 
  data:
    pools:   2 pools, 200 pgs
    objects: 21 objects, 3770 bytes
    usage:   2050 MB used, 37618 MB / 39669 MB avail
    pgs:     200 active+clean
 
  io:
    client:   852 B/s rd, 1 op/s rd, 0 op/s wr
```
 
```
[vagrant@localhost ~]$ kubectl -n rook exec -ti rook-tools ceph df
GLOBAL:
    SIZE       AVAIL      RAW USED     %RAW USED 
    39669M     37618M        2050M          5.17 
POOLS:
    NAME              ID     USED     %USED     MAX AVAIL     OBJECTS 
    myfs-metadata     1      3770         0        35635M          21 
    myfs-data0        2         0         0        35635M           0 
```

```
[vagrant@localhost ~]$ kubectl -n rook exec -ti rook-tools rados df
POOL_NAME     USED OBJECTS CLONES COPIES MISSING_ON_PRIMARY UNFOUND DEGRADED RD_OPS RD     WR_OPS WR    
myfs-data0       0       0      0      0                  0       0        0      0      0      0     0 
myfs-metadata 3770      21      0     21                  0       0        0  28650 14332k     55 21504 

total_objects    21
total_used       2050M
total_avail      37618M
total_space      39669M
```

```
[vagrant@localhost ~]$ kubectl -n rook exec -ti rook-tools rookctl status
OVERALL STATUS: OK

USAGE:
TOTAL       USED       DATA       AVAILABLE
38.74 GiB   2.00 GiB   3.68 KiB   36.74 GiB

MONITORS:
NAME             ADDRESS              IN QUORUM   STATUS
rook-ceph-mon0   10.3.0.88:6790/0     true        OK
rook-ceph-mon2   10.3.11.34:6790/0    true        OK
rook-ceph-mon1   10.3.15.157:6790/0   true        OK

MGRs:
NAME             STATUS
rook-ceph-mgr0   Active
rook-ceph-mgr1   Standby

OSDs:
TOTAL     UP        IN        FULL      NEAR FULL
1         1         1         false     false

PLACEMENT GROUPS (200 total):
STATE          COUNT
active+clean   200
```

Test
```
[vagrant@localhost rook0x2Frook]$ kubectl create -f gofileserver-deployment.yaml 
deployment "gofileserver" created
```


__Behind__

CRD
```
[vagrant@localhost rook0x2Frook]$ kubectl get crd
NAME                        AGE
clusters.rook.io            3h
filesystems.rook.io         3h
objectstores.rook.io        3h
pools.rook.io               3h
volumeattachments.rook.io   3h
```

Flex volume plugin
```
[vagrant@localhost rook0x2Frook]$ ls /usr/libexec/kubernetes/kubelet-plugins/volume/exec/rook.io~rook/
rook
```

## K8s v1.8.4 cluster

For kubectl
```
ubuntu@kubedev-10-64-33-200:~$ export KUBECONFIG=/home/ubuntu/etc0x2Fkubernetes/admin.conf 
```

```
ubuntu@kubedev-10-64-33-200:~$ kubectl get nodes
NAME                   STATUS    ROLES     AGE       VERSION
kubedev-10-64-33-195   Ready     master    18d       v1.8.4
kubedev-10-64-33-199   Ready     master    18d       v1.8.4
kubedev-10-64-33-200   Ready     <none>    1h        v1.8.4
kubedev-10-64-33-82    Ready     master    18d       v1.8.4
```

### rook

Deploy rook
```
ubuntu@kubedev-10-64-33-200:~$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-operator.yaml 
namespace "rook-system" created
clusterrole "rook-operator" created
serviceaccount "rook-operator" created
clusterrolebinding "rook-operator" created
deployment "rook-operator" created
```

Deploy cluster
```


```
ubuntu@kubedev-10-64-33-200:~$ kubectl -n rook-system get pods -o wide
NAME                             READY     STATUS    RESTARTS   AGE       IP            NODE
rook-operator-5c5db8db4c-fs6wd   1/1       Running   23         1h        10.244.3.14   kubedev-10-64-33-200
```

trouble???