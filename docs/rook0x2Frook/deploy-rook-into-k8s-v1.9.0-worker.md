### DevOps

Follow [../k8s-v1.9.0-deploy/storage-cluster-hand-on-rook-for-fedora26.md](../k8s-v1.9.0-deploy/storage-cluster-hand-on-rook-for-fedora26.md)


```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get nodes        
NAME                  STATUS                        ROLES     AGE       VERSION
kubedev-172-17-4-55   NotReady,SchedulingDisabled   <none>    1d        v1.9.0
kubedev-172-17-4-59   Ready                         master    1d        v1.9.0
kubedev-172-17-4-65   NotReady,SchedulingDisabled   <none>    14h       v1.9.0
rookdev-172-17-4-61   Ready                         <none>    9h        v1.9.0
rookdev-172-17-4-63   Ready                         <none>    18h       v1.9.0
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get nodes -o wide
NAME                  STATUS                        ROLES     AGE       VERSION   EXTERNAL-IP   OS-IMAGE                       KERNEL-VERSION           CONTAINER-RUNTIME
kubedev-172-17-4-55   NotReady,SchedulingDisabled   <none>    1d        v1.9.0    <none>        Ubuntu 16.04.3 LTS             4.4.0-104-generic        docker://1.13.1
kubedev-172-17-4-59   Ready                         master    1d        v1.9.0    <none>        Fedora 26 (Cloud Edition)      4.11.8-300.fc26.x86_64   docker://1.13.1
kubedev-172-17-4-65   NotReady,SchedulingDisabled   <none>    14h       v1.9.0    <none>        Debian GNU/Linux 9 (stretch)   4.9.0-4-amd64            docker://1.13.1
rookdev-172-17-4-61   Ready                         <none>    9h        v1.9.0    <none>        Ubuntu 17.04                   4.10.0-42-generic        docker://1.13.1
rookdev-172-17-4-63   Ready                         <none>    18h       v1.9.0    <none>        Fedora 27 (Cloud Edition)      4.13.9-300.fc27.x86_64   docker://1.13.1
```




```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get nodes
NAME                  STATUS                        ROLES     AGE       VERSION
kubedev-172-17-4-55   NotReady,SchedulingDisabled   <none>    15h       v1.9.0
kubedev-172-17-4-59   Ready                         master    1d        v1.9.0
kubedev-172-17-4-65   Ready                         <none>    2h        v1.9.0
rookdev-172-17-4-63   Ready                         <none>    6h        v1.9.0
```

rook operator
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-operator.yaml 
namespace "rook-system" created
clusterrole "rook-operator" created
serviceaccount "rook-operator" created
clusterrolebinding "rook-operator" created
deployment "rook-operator" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl -n rook-system get pods -o wide
NAME                            READY     STATUS    RESTARTS   AGE       IP            NODE
rook-agent-2z9lp                0/1       Pending   0          1m        <none>        kubedev-172-17-4-55
rook-agent-7zg7w                1/1       Running   0          1m        172.17.4.63   rookdev-172-17-4-63
rook-agent-ktfsp                0/1       Pending   0          1m        <none>        kubedev-172-17-4-65
rook-agent-mm72r                1/1       Running   0          1m        172.17.4.61   rookdev-172-17-4-61
rook-operator-cccddbbf9-ms9tc   1/1       Running   0          1m        10.244.3.3    rookdev-172-17-4-63
```

ceph cluster
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-cluster.yaml 
namespace "rook" created
cluster "rook" created
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl -n rook get cluster
NAME      AGE
rook      32s
```

```
NAME                              READY     STATUS    RESTARTS   AGE       IP            NODE
rook-api-848df956bf-sw4mh         1/1       Running   0          17s       10.244.2.11   rookdev-172-17-4-61
rook-ceph-mgr0-75954cf8c7-8ff8l   1/1       Running   0          18s       10.244.2.10   rookdev-172-17-4-61
rook-ceph-mgr1-5d8b54f9d9-t98cl   1/1       Running   0          18s       10.244.3.5    rookdev-172-17-4-63
rook-ceph-mon0-j7wz9              1/1       Running   0          56s       10.244.2.8    rookdev-172-17-4-61
rook-ceph-mon1-2rbww              1/1       Running   0          46s       10.244.3.4    rookdev-172-17-4-63
rook-ceph-mon2-zxsf6              1/1       Running   0          32s       10.244.2.9    rookdev-172-17-4-61
rook-ceph-osd-dlgkl               0/1       Pending   0          17s       <none>        kubedev-172-17-4-55
rook-ceph-osd-lf8jq               1/1       Running   0          17s       10.244.3.6    rookdev-172-17-4-63
rook-ceph-osd-pdbpw               0/1       Pending   0          17s       <none>        kubedev-172-17-4-65
rook-ceph-osd-x2sx7               1/1       Running   0          17s       10.244.2.12   rookdev-172-17-4-61
```

ceph filesystem
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-filesystem.yaml 
filesystem "myfs" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl -n rook get filesystem  
NAME      AGE
myfs      35s
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl -n rook get pods -o wide
NAME                                  READY     STATUS    RESTARTS   AGE       IP            NODE
rook-api-848df956bf-sw4mh             1/1       Running   0          2m        10.244.2.11   rookdev-172-17-4-61
rook-ceph-mds-myfs-7d59fdfcf4-4rzkd   1/1       Running   0          41s       10.244.3.7    rookdev-172-17-4-63
rook-ceph-mds-myfs-7d59fdfcf4-rdkx8   1/1       Running   0          41s       10.244.2.13   rookdev-172-17-4-61
rook-ceph-mgr0-75954cf8c7-8ff8l       1/1       Running   0          2m        10.244.2.10   rookdev-172-17-4-61
rook-ceph-mgr1-5d8b54f9d9-t98cl       1/1       Running   0          2m        10.244.3.5    rookdev-172-17-4-63
rook-ceph-mon0-j7wz9                  1/1       Running   0          3m        10.244.2.8    rookdev-172-17-4-61
rook-ceph-mon1-2rbww                  1/1       Running   0          3m        10.244.3.4    rookdev-172-17-4-63
rook-ceph-mon2-zxsf6                  1/1       Running   0          3m        10.244.2.9    rookdev-172-17-4-61
rook-ceph-osd-dlgkl                   0/1       Pending   0          2m        <none>        kubedev-172-17-4-55
rook-ceph-osd-lf8jq                   1/1       Running   0          2m        10.244.3.6    rookdev-172-17-4-63
rook-ceph-osd-pdbpw                   0/1       Pending   0          2m        <none>        kubedev-172-17-4-65
rook-ceph-osd-x2sx7                   1/1       Running   0          2m        10.244.2.12   rookdev-172-17-4-61
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-tools.yaml      
pod "rook-tools" created
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/docs/rook0x2Frook/gofileserver-deployment.yaml 
deployment "gofileserver" created
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl get pods -o wide -l app=gofileserver
NAME                            READY     STATUS    RESTARTS   AGE       IP            NODE
gofileserver-7df6f9ccb8-kn5fx   1/1       Running   0          13m       10.244.2.16   rookdev-172-17-4-61
```

```
ubuntu@rookdev-172-17-4-61:~$ curl http://10.244.2.16:48080/
<pre>
</pre>
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pods -o wide -l app=gofileserver
NAME                            READY     STATUS    RESTARTS   AGE       IP            NODE
gofileserver-7df6f9ccb8-kn5fx   1/1       Running   0          50s       10.244.2.16   rookdev-172-17-4-61
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -n rook -ti rook-tools bash
```

```
root@rook-tools:/# rookctl status
OVERALL STATUS: WARNING

SUMMARY:
SEVERITY   NAME             MESSAGE
WARNING    MON_CLOCK_SKEW   clock skew detected on mon.rook-ceph-mon2, mon.rook-ceph-mon0

USAGE:
TOTAL       USED       DATA       AVAILABLE
10.63 GiB   2.00 GiB   7.03 KiB   8.63 GiB

MONITORS:
NAME             ADDRESS                 IN QUORUM   STATUS
rook-ceph-mon1   10.99.123.14:6790/0     true        OK
rook-ceph-mon2   10.100.241.37:6790/0    true        WARNING
rook-ceph-mon0   10.102.156.194:6790/0   true        WARNING

MGRs:
NAME             STATUS
rook-ceph-mgr0   Active
rook-ceph-mgr1   Standby

OSDs:
TOTAL     UP        IN        FULL      NEAR FULL
2         1         1         false     false

PLACEMENT GROUPS (200 total):
STATE          COUNT
active+clean   200
```

```
root@rook-tools:/# ceph status
  cluster:
    id:     fd2e1f2b-4bce-41ee-9a91-f8a1a30cd445
    health: HEALTH_WARN
            clock skew detected on mon.rook-ceph-mon2, mon.rook-ceph-mon0
 
  services:
    mon: 3 daemons, quorum rook-ceph-mon1,rook-ceph-mon2,rook-ceph-mon0
    mgr: rook-ceph-mgr0(active), standbys: rook-ceph-mgr1
    mds: myfs-1/1/1 up  {0=m4rzkd=up:active}, 1 up:standby-replay
    osd: 2 osds: 1 up, 1 in
 
  data:
    pools:   2 pools, 200 pgs
    objects: 21 objects, 7195 bytes
    usage:   2050 MB used, 8835 MB / 10885 MB avail
    pgs:     200 active+clean
 
  io:
    client:   852 B/s rd, 1 op/s rd, 0 op/s wr
 
```

```
root@rook-tools:/# ceph df
GLOBAL:
    SIZE       AVAIL     RAW USED     %RAW USED 
    10885M     8835M        2050M         18.84 
POOLS:
    NAME              ID     USED     %USED     MAX AVAIL     OBJECTS 
    myfs-metadata     1      7195         0         8290M          21 
    myfs-data0        2         0         0         8290M           0 
```

```
root@rook-tools:/# rados df
POOL_NAME     USED OBJECTS CLONES COPIES MISSING_ON_PRIMARY UNFOUND DEGRADED RD_OPS RD    WR_OPS WR    
myfs-data0       0       0      0      0                  0       0        0      0     0      0     0 
myfs-metadata 7195      21      0     21                  0       0        0   3395 1709k     78 45056 

total_objects    21
total_used       2050M
total_avail      8835M
total_space      10885M
```

```
root@rook-tools:/# mkdir /tmp/myfs
```

```
root@rook-tools:/# rookctl filesystem mount --name=myfs --path=/tmp/myfs
succeeded mounting shared filesystem myfs at '/tmp/myfs'
```

```
root@rook-tools:/# echo "world" > /tmp/myfs/hello
```

```
[vagrant@kubedev-172-17-4-59 ~]$ curl -L http://10.244.2.16:48080/
<pre>
<a href="hello">hello</a>
</pre>
```

```
[vagrant@kubedev-172-17-4-59 ~]$ curl -L http://10.244.2.16:48080/hello
world
```
