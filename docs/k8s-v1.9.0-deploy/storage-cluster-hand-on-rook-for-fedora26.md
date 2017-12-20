# DevOps

Image
```
[vagrant@kubedev-172-17-4-59 go-to-kubernetes]$ docker pull docker.io/rook/rook:master
Trying to pull repository docker.io/rook/rook ... 
sha256:bb88d929d233b2cc4cf8b4765715fd05642de1002010b8d178fd04f9367503c4: Pulling from docker.io/rook/rook
cd52aae0f560: Pull complete 
46152aaa3769: Pull complete 
8a49792ee32b: Pull complete 
a43db90f1073: Pull complete 
cfb724a24cc5: Pull complete 
3df4aa5ae86f: Pull complete 
e033543339fe: Pull complete 
7924ea32c0a7: Pull complete 
Digest: sha256:bb88d929d233b2cc4cf8b4765715fd05642de1002010b8d178fd04f9367503c4
Status: Downloaded newer image for docker.io/rook/rook:master
```

```
[vagrant@kubedev-172-17-4-59 go-to-kubernetes]$ docker pull docker.io/rook/toolbox:master
Trying to pull repository docker.io/rook/toolbox ... 
sha256:f2b581df6718e525c6fa9e72d76dbe9b725a07c09760dd48f6aa8be14d9d1c4c: Pulling from docker.io/rook/toolbox
cd52aae0f560: Already exists 
46152aaa3769: Already exists 
8a49792ee32b: Already exists 
a43db90f1073: Already exists 
cfb724a24cc5: Already exists 
3df4aa5ae86f: Already exists 
e033543339fe: Already exists 
fe1856b939e6: Pull complete 
966cfe7613cf: Pull complete 
3eba886796b6: Pull complete 
Digest: sha256:f2b581df6718e525c6fa9e72d76dbe9b725a07c09760dd48f6aa8be14d9d1c4c
Status: Downloaded newer image for docker.io/rook/toolbox:master
```

## Before

worker node required
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get nodes
NAME                  STATUS    ROLES     AGE       VERSION
kubedev-172-17-4-55   Ready     <none>    1h        v1.9.0
kubedev-172-17-4-59   Ready     master    19h       v1.9.0
```

modify to deploy operator into master
```
[vagrant@kubedev-172-17-4-59 ~]$ mkdir rook0x2Fcluster && cp -r /Users/fanhongling/go/src/github.com/rook/rook/cluster/examples rook0x2Fcluster
```

append for rook-operator.yaml
```
      tolerations:
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
```

## rook

### operator

create
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-operator.yaml 
namespace "rook-system" created
clusterrole "rook-operator" created
serviceaccount "rook-operator" created
clusterrolebinding "rook-operator" created
deployment "rook-operator" created
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl -n rook-system get pods
NAME                            READY     STATUS    RESTARTS   AGE
rook-operator-cccddbbf9-2qlj7   1/1       Running   0          21s
```

### cluster

create
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-cluster.yaml 
namespace "rook" created
cluster "rook" created
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl -n rook get cluster
NAME      AGE
rook      18h
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl -n rook-system get pods -o wide
NAME                            READY     STATUS    RESTARTS   AGE       IP            NODE
rook-agent-9d5mv                1/1       Running   0          54m       172.17.4.55   kubedev-172-17-4-55
rook-operator-cccddbbf9-2qlj7   1/1       Running   0          18h       10.244.0.4    kubedev-172-17-4-59
```

```
ubuntu@ubuntu-xenial:~$ kubectl get crd
NAME                        AGE
clusters.rook.io            18h
filesystems.rook.io         18h
objectstores.rook.io        18h
pools.rook.io               18h
volumeattachments.rook.io   18h
```

worker node
```
ubuntu@ubuntu-xenial:~$ ls /usr/libexec/kubernetes/kubelet-plugins/volume/exec/
rook.io~rook
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl -n rook-system logs rook-agent-9d5mv             
2017-12-19 02:38:53.503192 I | rook: starting Rook v0.6.0-98.gc1a87fa with arguments '/usr/local/bin/rook agent'
2017-12-19 02:38:53.503372 I | rook: flag values: --help=false, --log-level=INFO
2017-12-19 02:38:53.504152 I | rook: starting rook agent
2017-12-19 02:38:53.513581 I | exec: Running command: modinfo -F parm rbd
2017-12-19 02:38:53.521290 I | exec: Running command: modprobe rbd single_major=Y
2017-12-19 02:38:53.608975 I | flexvolume: Rook Flexvolume configured
2017-12-19 02:38:53.609069 I | flexvolume: Listening on unix socket for Kubernetes volume attach commands.
2017-12-19 02:38:53.611079 I | agent-cluster: start watching cluster resources
```

[vagrant@kubedev-172-17-4-59 ~]$ kubectl -n rook-system logs rook-operator-cccddbbf9-2qlj7
```
2017-12-19 03:44:01.072434 I | rook: starting Rook v0.6.0-96.g043cd8b with arguments '/usr/local/bin/rook operator'
2017-12-19 03:44:01.074603 I | rook: flag values: --help=false, --log-level=INFO, --mon-healthcheck-interval=45s, --mon-out-timeout=5m0s
2017-12-19 03:44:01.076526 I | rook: starting operator
2017-12-19 03:44:03.741572 I | op-k8sutil: cluster role rook-agent already exists. Updating if needed.
2017-12-19 03:44:03.766212 I | op-agent: getting flexvolume dir path from FLEXVOLUME_DIR_PATH env var
2017-12-19 03:44:03.766338 I | op-agent: flexvolume dir path env var FLEXVOLUME_DIR_PATH is not provided. Defaulting to: /usr/libexec/kubernetes/kubelet-plugins/volume/exec/
2017-12-19 03:44:03.766368 I | op-agent: discovered flexvolume dir path from source default. value: /usr/libexec/kubernetes/kubelet-plugins/volume/exec/
2017-12-19 03:44:03.780320 I | op-agent: rook-agent daemonset already exists
2017-12-19 03:44:03.797809 I | operator: rook-provisioner started
2017-12-19 03:44:03.797903 I | op-cluster: start watching clusters in all namespaces
ERROR: logging before flag.Parse: I1219 03:44:03.798876       5 controller.go:407] Starting provisioner controller db7a1885-e46e-11e7-bc62-0a580af40004!
2017-12-19 03:44:03.829575 I | op-cluster: starting cluster namespace rook
2017-12-19 03:44:09.837361 I | op-mon: start running mons
2017-12-19 03:44:09.842568 I | cephmon: parsing mon endpoints: 
2017-12-19 03:44:09.842634 W | cephmon: ignoring invalid monitor 
2017-12-19 03:44:09.842689 I | op-mon: loaded: maxMonID=-1, mons=map[], mapping=&{Node:map[] Port:map[]}
2017-12-19 03:44:09.848673 I | op-mon: saved mon endpoints to config map map[maxMonId:-1 mapping:{"node":{},"port":{}} data:]
2017-12-19 03:44:09.850159 I | cephmon: writing config file /var/lib/rook/rook/rook.config
2017-12-19 03:44:09.851510 I | cephmon: generated admin config in /var/lib/rook/rook
2017-12-19 03:44:09.871027 I | op-mon: Found 1 running nodes without mons
2017-12-19 03:44:09.910066 I | op-mon: mon rook-ceph-mon0 running at 10.101.233.25:6790
2017-12-19 03:44:09.947943 I | op-mon: saved mon endpoints to config map map[mapping:{"node":{"rook-ceph-mon0":{"Name":"kubedev-172-17-4-55","Address":"172.17.4.55"},"rook-ceph-mon1":{"Name":"kubedev-172-17-4-55","Address":"172.17.4.55"},"rook-ceph-mon2":{"Name":"kubedev-172-17-4-55","Address":"172.17.4.55"}},"port":{}} data:rook-ceph-mon0=10.101.233.25:6790 maxMonId:2]
2017-12-19 03:44:09.948959 I | cephmon: writing config file /var/lib/rook/rook/rook.config
2017-12-19 03:44:09.949109 I | cephmon: generated admin config in /var/lib/rook/rook
2017-12-19 03:44:09.957215 I | cephmon: writing config file /var/lib/rook/rook/rook.config
2017-12-19 03:44:09.958313 I | cephmon: generated admin config in /var/lib/rook/rook
2017-12-19 03:44:10.012837 I | op-mon: mons created: 1
2017-12-19 03:44:10.014014 I | op-mon: waiting for mon quorum
2017-12-19 03:44:10.042922 I | exec: Running command: ceph mon_status --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/600093975
2017-12-19 03:44:26.537702 I | op-mon: Ceph monitors formed quorum
2017-12-19 03:44:26.552417 I | op-mon: mon rook-ceph-mon0 running at 10.101.233.25:6790
2017-12-19 03:44:26.559899 I | op-mon: mon rook-ceph-mon1 running at 10.105.74.10:6790
2017-12-19 03:44:26.579937 I | op-mon: saved mon endpoints to config map map[data:rook-ceph-mon0=10.101.233.25:6790,rook-ceph-mon1=10.105.74.10:6790 maxMonId:2 mapping:{"node":{"rook-ceph-mon0":{"Name":"kubedev-172-17-4-55","Address":"172.17.4.55"},"rook-ceph-mon1":{"Name":"kubedev-172-17-4-55","Address":"172.17.4.55"},"rook-ceph-mon2":{"Name":"kubedev-172-17-4-55","Address":"172.17.4.55"}},"port":{}}]
2017-12-19 03:44:26.581934 I | cephmon: writing config file /var/lib/rook/rook/rook.config
2017-12-19 03:44:26.582106 I | cephmon: generated admin config in /var/lib/rook/rook
2017-12-19 03:44:26.594456 I | cephmon: writing config file /var/lib/rook/rook/rook.config
2017-12-19 03:44:26.594713 I | cephmon: generated admin config in /var/lib/rook/rook
2017-12-19 03:44:26.611994 I | op-mon: replicaset rook-ceph-mon0 already exists
2017-12-19 03:44:26.628310 I | op-mon: mons created: 2
2017-12-19 03:44:26.628389 I | op-mon: waiting for mon quorum
2017-12-19 03:44:26.628512 I | exec: Running command: ceph mon_status --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/483544970
2017-12-19 03:44:26.884997 W | op-mon: failed to find initial monitor rook-ceph-mon1 in mon map
2017-12-19 03:44:31.885636 I | exec: Running command: ceph mon_status --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/128790369
2017-12-19 03:44:32.092255 W | op-mon: failed to find initial monitor rook-ceph-mon1 in mon map
2017-12-19 03:44:37.093281 I | exec: Running command: ceph mon_status --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/934719308
2017-12-19 03:44:37.355549 W | op-mon: failed to find initial monitor rook-ceph-mon1 in mon map
2017-12-19 03:44:42.356004 I | exec: Running command: ceph mon_status --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/551176251
2017-12-19 03:44:47.556272 I | op-mon: Ceph monitors formed quorum
2017-12-19 03:44:47.572583 I | op-mon: mon rook-ceph-mon0 running at 10.101.233.25:6790
2017-12-19 03:44:47.593284 I | op-mon: mon rook-ceph-mon1 running at 10.105.74.10:6790
2017-12-19 03:44:47.605599 I | op-mon: mon rook-ceph-mon2 running at 10.109.229.27:6790
2017-12-19 03:44:47.625378 I | op-mon: saved mon endpoints to config map map[data:rook-ceph-mon0=10.101.233.25:6790,rook-ceph-mon1=10.105.74.10:6790,rook-ceph-mon2=10.109.229.27:6790 maxMonId:2 mapping:{"node":{"rook-ceph-mon0":{"Name":"kubedev-172-17-4-55","Address":"172.17.4.55"},"rook-ceph-mon1":{"Name":"kubedev-172-17-4-55","Address":"172.17.4.55"},"rook-ceph-mon2":{"Name":"kubedev-172-17-4-55","Address":"172.17.4.55"}},"port":{}}]
2017-12-19 03:44:47.625756 I | cephmon: writing config file /var/lib/rook/rook/rook.config
2017-12-19 03:44:47.625881 I | cephmon: generated admin config in /var/lib/rook/rook
2017-12-19 03:44:47.633638 I | cephmon: writing config file /var/lib/rook/rook/rook.config
2017-12-19 03:44:47.635029 I | cephmon: generated admin config in /var/lib/rook/rook
2017-12-19 03:44:47.646534 I | op-mon: replicaset rook-ceph-mon0 already exists
2017-12-19 03:44:47.652477 I | op-mon: replicaset rook-ceph-mon1 already exists
2017-12-19 03:44:47.669466 I | op-mon: mons created: 3
2017-12-19 03:44:47.669558 I | op-mon: waiting for mon quorum
2017-12-19 03:44:47.670359 I | exec: Running command: ceph mon_status --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/907535966
2017-12-19 03:44:48.001677 W | op-mon: failed to find initial monitor rook-ceph-mon2 in mon map
2017-12-19 03:44:53.003512 I | exec: Running command: ceph mon_status --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/010295845
2017-12-19 03:44:53.207038 W | op-mon: failed to find initial monitor rook-ceph-mon2 in mon map
2017-12-19 03:44:58.207445 I | exec: Running command: ceph mon_status --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/785791808
2017-12-19 03:44:58.472321 W | op-mon: failed to find initial monitor rook-ceph-mon2 in mon map
2017-12-19 03:45:03.473156 I | exec: Running command: ceph mon_status --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/797161375
2017-12-19 03:45:08.598297 I | op-mon: Ceph monitors formed quorum
2017-12-19 03:45:08.601197 I | op-cluster: creating initial crushmap
2017-12-19 03:45:08.601307 I | cephclient: setting crush tunables to firefly
2017-12-19 03:45:08.601408 I | exec: Running command: ceph osd crush tunables firefly --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format plain --out-file /tmp/265265266
2017-12-19 03:45:08.838683 I | exec: adjusted tunables profile to firefly
2017-12-19 03:45:08.838835 I | cephclient: succeeded setting crush tunables to profile firefly: 
2017-12-19 03:45:08.838956 I | exec: Running command: crushtool -c /tmp/862894889 -o /tmp/885291124
2017-12-19 03:45:08.866278 I | exec: Running command: ceph osd setcrushmap -i /tmp/885291124 --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/496725827
2017-12-19 03:45:10.017721 I | exec: 3
2017-12-19 03:45:10.017917 I | op-cluster: created initial crushmap
2017-12-19 03:45:10.021873 I | op-mgr: start running mgr
2017-12-19 03:45:10.026400 I | exec: Running command: ceph auth get-or-create-key mgr.rook-ceph-mgr0 mon allow * --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/510374854
2017-12-19 03:45:10.311975 I | op-mgr: rook-ceph-mgr0 deployment started
2017-12-19 03:45:10.324901 I | exec: Running command: ceph auth get-or-create-key mgr.rook-ceph-mgr1 mon allow * --cluster=rook --conf=/var/lib/rook/rook/rook.config --keyring=/var/lib/rook/rook/client.admin.keyring --format json --out-file /tmp/280046189
2017-12-19 03:45:10.807485 I | op-mgr: rook-ceph-mgr1 deployment started
2017-12-19 03:45:10.810797 I | op-api: starting the Rook api
2017-12-19 03:45:10.898672 I | op-api: API service running at 10.101.96.44:8124
2017-12-19 03:45:10.951246 I | op-k8sutil: creating role rook-api in namespace rook
2017-12-19 03:45:10.993410 I | op-api: api deployment started
2017-12-19 03:45:10.995943 I | op-osd: start running osds in namespace rook
2017-12-19 03:45:11.053810 I | op-k8sutil: creating role rook-ceph-osd in namespace rook
2017-12-19 03:45:11.146732 I | op-osd: osd daemon set started
2017-12-19 03:45:11.150974 I | op-cluster: Done creating rook instance in namespace rook
2017-12-19 03:45:11.151078 I | op-pool: start watching pool resources in namespace rook
2017-12-19 03:45:11.151113 I | op-rgw: start watching object store resources in namespace rook
2017-12-19 03:45:11.151144 I | op-mds: start watching filesystem resource in namespace rook
```

if ceph does not work, restart container
```
[vagrant@kubedev-172-17-4-59 ~]$ docker ps --filter "label=io.kubernetes.container.name=rook-operator"
CONTAINER ID        IMAGE                                                                                         COMMAND                  CREATED             STATUS              PORTS               NAMES
87c02788e147        docker.io/rook/rook@sha256:bb88d929d233b2cc4cf8b4765715fd05642de1002010b8d178fd04f9367503c4   "/tini -- /usr/loc..."   18 hours ago        Up 9 minutes                            k8s_rook-operator_rook-operator-cccddbbf9-2qlj7_rook-system_90a44bd7-e3d3-11e7-8db4-525400224e72_0
```

```
[vagrant@kubedev-172-17-4-59 ~]$ docker restart 87c02788e147
```

```
ubuntu@ubuntu-xenial:~$ kubectl -n rook get pods
NAME                              READY     STATUS    RESTARTS   AGE
rook-api-848df956bf-28pkt         1/1       Running   0          17s
rook-ceph-mgr0-75954cf8c7-dkklq   1/1       Running   0          18s
rook-ceph-mgr1-5d8b54f9d9-85rdv   1/1       Running   0          18s
rook-ceph-mon0-74bsm              1/1       Running   0          1m
rook-ceph-mon1-xs6cs              1/1       Running   0          1m
rook-ceph-mon2-jz5qh              1/1       Running   0          41s
rook-ceph-osd-mdhk8               1/1       Running   0          17s
```

### file system

create 
```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-filesystem.yaml 
filesystem "myfs" created
```

```
ubuntu@ubuntu-xenial:~$ kubectl -n rook get filesystem
NAME      AGE
myfs      24s
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl -n rook get pods -o wide
NAME                                  READY     STATUS    RESTARTS   AGE       IP            NODE
rook-api-848df956bf-28pkt             1/1       Running   0          13m       10.244.1.7    kubedev-172-17-4-55
rook-ceph-mds-myfs-7d59fdfcf4-9x7jn   1/1       Running   0          1m        10.244.1.10   kubedev-172-17-4-55
rook-ceph-mds-myfs-7d59fdfcf4-h69fr   1/1       Running   0          1m        10.244.1.9    kubedev-172-17-4-55
rook-ceph-mgr0-75954cf8c7-dkklq       1/1       Running   0          13m       10.244.1.5    kubedev-172-17-4-55
rook-ceph-mgr1-5d8b54f9d9-85rdv       1/1       Running   0          13m       10.244.1.6    kubedev-172-17-4-55
rook-ceph-mon0-74bsm                  1/1       Running   0          14m       10.244.1.2    kubedev-172-17-4-55
rook-ceph-mon1-xs6cs                  1/1       Running   0          14m       10.244.1.3    kubedev-172-17-4-55
rook-ceph-mon2-jz5qh                  1/1       Running   0          13m       10.244.1.4    kubedev-172-17-4-55
rook-ceph-osd-mdhk8                   1/1       Running   0          13m       10.244.1.8    kubedev-172-17-4-55
```

```
ubuntu@ubuntu-xenial:~$ kubectl -n rook get service -o wide
NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE       SELECTOR
rook-api         ClusterIP   10.101.96.44    <none>        8124/TCP   1h        app=rook-api,rook_cluster=rook
rook-ceph-mon0   ClusterIP   10.101.233.25   <none>        6790/TCP   1h        app=rook-ceph-mon,mon=rook-ceph-mon0,mon_cluster=rook
rook-ceph-mon1   ClusterIP   10.105.74.10    <none>        6790/TCP   1h        app=rook-ceph-mon,mon=rook-ceph-mon1,mon_cluster=rook
rook-ceph-mon2   ClusterIP   10.109.229.27   <none>        6790/TCP   1h        app=rook-ceph-mon,mon=rook-ceph-mon2,mon_cluster=rook
```

Using such files for test
```
ubuntu@ubuntu-xenial:~$ ls -l /tmp/kubeadm*
lrwxrwxrwx 1 root root  63 Dec 19 02:36 /tmp/kubeadm.INFO -> kubeadm.kubedev-172-17-4-55.root.log.INFO.20171219-023657.12445
-rw-r--r-- 1 root root 195 Dec 19 02:33 /tmp/kubeadm.kubedev-172-17-4-55.root.log.INFO.20171219-023332.12218
-rw-r--r-- 1 root root 195 Dec 19 02:36 /tmp/kubeadm.kubedev-172-17-4-55.root.log.INFO.20171219-023641.12405
-rw-r--r-- 1 root root 195 Dec 19 02:36 /tmp/kubeadm.kubedev-172-17-4-55.root.log.INFO.20171219-023657.12445
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/docs/rook0x2Frook/gofileserver-deployment.yaml 
deployment "gofileserver" created
```

__issue__
```
ubuntu@ubuntu-xenial:~$ kubectl get pods
NAME                            READY     STATUS     RESTARTS   AGE
gofileserver-5c5ccf45f4-sr66m   0/1       Init:0/1   0          10s
hello-6cb8784cb-2gzbc           1/1       Running    0          20h
ubuntu@ubuntu-xenial:~$ kubectl describe pod gofileserver-5c5ccf45f4-sr66m
Name:           gofileserver-5c5ccf45f4-sr66m
Namespace:      default
Node:           kubedev-172-17-4-55/172.17.4.55
Start Time:     Tue, 19 Dec 2017 04:11:41 +0000
Labels:         app=gofileserver
                pod-template-hash=1717790190
Annotations:    <none>
Status:         Pending
IP:             
Controlled By:  ReplicaSet/gofileserver-5c5ccf45f4
Init Containers:
  sidecar:
    Container ID:  
    Image:         busybox
    Image ID:      
    Port:          <none>
    Command:
      cp
    Args:
      -rf
      /tmp/test/kubeadm*
      /mnt/test
    State:          Waiting
      Reason:       PodInitializing
    Ready:          False
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /mnt/test from download-store (rw)
      /tmp/test from download-sidecar (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-67l4p (ro)
Containers:
  gofileserver:
    Container ID:   
    Image:          tangfeixiong/gofileserver
    Image ID:       
    Port:           48080/TCP
    State:          Waiting
      Reason:       PodInitializing
    Ready:          False
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /mnt from download-store (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-67l4p (ro)
Conditions:
  Type           Status
  Initialized    False 
  Ready          False 
  PodScheduled   True 
Volumes:
  download-sidecar:
    Type:          HostPath (bare host directory volume)
    Path:          /tmp
    HostPathType:  
  download-store:
    Type:    FlexVolume (a generic volume resource that is provisioned/attached using an exec based plugin)
    Driver:      Options:  %v

    FSType:     rook.io/rook
    SecretRef:  ceph
    ReadOnly:   <nil>
%!(EXTRA bool=false, map[string]string=map[clusterName:rook fsName:myfs])  default-token-67l4p:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-67l4p
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type     Reason                 Age   From                          Message
  ----     ------                 ----  ----                          -------
  Normal   Scheduled              24s   default-scheduler             Successfully assigned gofileserver-5c5ccf45f4-sr66m to kubedev-172-17-4-55
  Normal   SuccessfulMountVolume  24s   kubelet, kubedev-172-17-4-55  MountVolume.SetUp succeeded for volume "download-sidecar"
  Normal   SuccessfulMountVolume  24s   kubelet, kubedev-172-17-4-55  MountVolume.SetUp succeeded for volume "default-token-67l4p"
  Warning  FailedMount            23s   kubelet, kubedev-172-17-4-55  MountVolume.SetUp failed for volume "download-store" : mount command failed, status: Failure, reason: failed to mount filesystem myfs to /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store with monitor 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/ and options [name=admin secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ==]: mount failed: exit status 32
Mounting command: systemd-run
Mounting arguments: --description=Kubernetes transient mount for /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store --scope -- mount -t ceph -o name=admin,secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ== 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/ /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store
Output: Running scope as unit run-r593eee49048747dd98005b871fe28585.scope.
mount: 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/: can't read superblock
  Warning  FailedMount  23s  kubelet, kubedev-172-17-4-55  MountVolume.SetUp failed for volume "download-store" : mount command failed, status: Failure, reason: failed to mount filesystem myfs to /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store with monitor 10.109.229.27:6790,10.101.233.25:6790,10.105.74.10:6790:/ and options [name=admin secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ==]: mount failed: exit status 32
Mounting command: systemd-run
Mounting arguments: --description=Kubernetes transient mount for /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store --scope -- mount -t ceph -o name=admin,secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ== 10.109.229.27:6790,10.101.233.25:6790,10.105.74.10:6790:/ /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store
Output: Running scope as unit run-redb9aa4b51a64c549b6ba05f49ad0480.scope.
mount: 10.109.229.27:6790,10.101.233.25:6790,10.105.74.10:6790:/: can't read superblock
  Warning  FailedMount  21s  kubelet, kubedev-172-17-4-55  MountVolume.SetUp failed for volume "download-store" : mount command failed, status: Failure, reason: failed to mount filesystem myfs to /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store with monitor 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/ and options [name=admin secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ==]: mount failed: exit status 32
Mounting command: systemd-run
Mounting arguments: --description=Kubernetes transient mount for /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store --scope -- mount -t ceph -o name=admin,secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ== 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/ /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store
Output: Running scope as unit run-refa4ac0bf2b749b782397b9ede98df8f.scope.
mount: 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/: can't read superblock
  Warning  FailedMount  19s  kubelet, kubedev-172-17-4-55  MountVolume.SetUp failed for volume "download-store" : mount command failed, status: Failure, reason: failed to mount filesystem myfs to /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store with monitor 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/ and options [name=admin secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ==]: mount failed: exit status 32
Mounting command: systemd-run
Mounting arguments: --description=Kubernetes transient mount for /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store --scope -- mount -t ceph -o name=admin,secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ== 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/ /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store
Output: Running scope as unit run-ra536f6c891154fba861bba69a85e78f8.scope.
mount: 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/: can't read superblock
  Warning  FailedMount  15s  kubelet, kubedev-172-17-4-55  MountVolume.SetUp failed for volume "download-store" : mount command failed, status: Failure, reason: failed to mount filesystem myfs to /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store with monitor 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/ and options [name=admin secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ==]: mount failed: exit status 32
Mounting command: systemd-run
Mounting arguments: --description=Kubernetes transient mount for /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store --scope -- mount -t ceph -o name=admin,secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ== 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/ /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store
Output: Running scope as unit run-r4019c580283242d99ae8e376892c967a.scope.
mount: 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/: can't read superblock
  Warning  FailedMount  6s  kubelet, kubedev-172-17-4-55  MountVolume.SetUp failed for volume "download-store" : mount command failed, status: Failure, reason: failed to mount filesystem myfs to /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store with monitor 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/ and options [name=admin secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ==]: mount failed: exit status 32
Mounting command: systemd-run
Mounting arguments: --description=Kubernetes transient mount for /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store --scope -- mount -t ceph -o name=admin,secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ== 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/ /var/lib/kubelet/pods/b770dafd-e472-11e7-8db4-525400224e72/volumes/rook.io~rook/download-store
Output: Running scope as unit run-rb938823ea6dd433d8282df5540ae1a43.scope.
mount: 10.105.74.10:6790,10.109.229.27:6790,10.101.233.25:6790:/: can't read superblock
```

Refert to https://github.com/rook/rook/issues/1044

https://github.com/rook/rook/issues/1220

and https://github.com/rook/rook/blob/master/Documentation/k8s-filesystem.md#kernel-version-requirement

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl create -f rook0x2Fcluster/examples/kubernetes/rook-tools.yaml 
pod "rook-tools" created
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl get pods -n rook rook-tools        
NAME         READY     STATUS    RESTARTS   AGE
rook-tools   1/1       Running   0          3m
```

```
[vagrant@kubedev-172-17-4-59 ~]$ kubectl exec -n rook -ti rook-tools bash
```

```
root@rook-tools:/# ls /usr/local/bin/
apt-ng-host-discover  ceph-brag      ceph-crush-location  ceph-mgr  ceph-post-file  ceph-run    rados          rbd              rookctl
ceph                  ceph-clsinfo   ceph-debugpack       ceph-mon  ceph-rbdnamer   crushtool   radosgw        rbd-replay-many  toolbox.sh
ceph-authtool         ceph-coverage  ceph-mds             ceph-osd  ceph-rest-api   monmaptool  radosgw-admin  rbdmap
```

```
root@rook-tools:/# ceph df
GLOBAL:
    SIZE       AVAIL     RAW USED     %RAW USED 
    10885M     8835M        2050M         18.84 
POOLS:
    NAME              ID     USED      %USED     MAX AVAIL     OBJECTS 
    myfs-metadata     1      14126         0         8290M          21 
    myfs-data0        2          0         0         8290M           0 
```

```
root@rook-tools:/# rados df
POOL_NAME     USED  OBJECTS CLONES COPIES MISSING_ON_PRIMARY UNFOUND DEGRADED RD_OPS RD    WR_OPS WR   
myfs-data0        0       0      0      0                  0       0        0      0     0      0    0 
myfs-metadata 14456      21      0     21                  0       0        0   5043 2547k    141 107k 

total_objects    21
total_used       2050M
total_avail      8835M
total_space      10885M
```

```
root@rook-tools:/# ceph status
  cluster:
    id:     2f34d4af-51c6-4def-8416-e0283862f1e5
    health: HEALTH_OK
 
  services:
    mon: 3 daemons, quorum rook-ceph-mon0,rook-ceph-mon1,rook-ceph-mon2
    mgr: rook-ceph-mgr0(active), standbys: rook-ceph-mgr1
    mds: myfs-1/1/1 up  {0=mh69fr=up:active}, 1 up:standby-replay
    osd: 1 osds: 1 up, 1 in
 
  data:
    pools:   2 pools, 200 pgs
    objects: 21 objects, 19406 bytes
    usage:   2050 MB used, 8835 MB / 10885 MB avail
    pgs:     200 active+clean
 
  io:
    client:   1279 B/s rd, 2 op/s rd, 0 op/s wr
 
```

```
root@rook-tools:/# rookctl status
OVERALL STATUS: OK

USAGE:
TOTAL       USED       DATA        AVAILABLE
10.63 GiB   2.00 GiB   14.12 KiB   8.63 GiB

MONITORS:
NAME             ADDRESS                IN QUORUM   STATUS
rook-ceph-mon0   10.101.233.25:6790/0   true        OK
rook-ceph-mon1   10.105.74.10:6790/0    true        OK
rook-ceph-mon2   10.109.229.27:6790/0   true        OK

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

```
root@rook-tools:/# rookctl filesystem mount --name myfs --path /tmp/myfs
2017-12-19 05:01:53.221284 I | mount 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/: mount: wrong fs type, bad option, bad superblock on 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/,
2017-12-19 05:01:53.221415 I | mount 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/:        missing codepage or helper program, or other error
2017-12-19 05:01:53.221439 I | mount 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/: 
2017-12-19 05:01:53.221463 I | mount 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/:        In some cases useful info is found in syslog - try
2017-12-19 05:01:53.221490 I | mount 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/:        dmesg | tail or so.
command mount 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/ failed: Failed to complete mount 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/: exit status 32
```

```
root@rook-tools:/# dmesg | tail
[26350.888617]  front: 000000e0: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
[26350.888618]  front: 000000f0: 00 00 00 00 02 00 00 00 00 00 00 00 04 00 00 00  ................
[26350.888619]  front: 00000100: 00 00 00 00 01 00 00 00 00 00 00 00 00 00 00 00  ................
[26350.888619]  front: 00000110: 00 00 00 00 00 00 00 00 30 00 00 00 01 00 00 00  ........0.......
[26350.888620]  front: 00000120: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
[26350.888621]  front: 00000130: 00 00 00 00 01 00 00 00 00 00 00 00 01 00 00 00  ................
[26350.888622]  front: 00000140: 00 00 00 00 00 00 00 00 00 00 00 00              ............
[26350.888623] footer: 00000000: 65 99 ea 51 00 00 00 00 00 00 00 00 6b e4 9b 81  e..Q........k...
[26350.888624] footer: 00000010: 4f 3f 96 18 05                                   O?...
[26390.514388] libceph: bad option at 'mds_namespace=myfs'
```

```
root@rook-tools:/# ceph-authtool /etc/ceph/keyring --list
[client.admin]
	key = AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ==
```

```
root@rook-tools:/# mount -t ceph -o name=admin,secret=AQAUhzdaEFrlHRAAOQrYHiWgJGXA/XTqUdlYmQ== 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/ /tmp/myfs
mount: 10.101.233.25:6790,10.105.74.10:6790,10.109.229.27:6790:/: can't read superblock
```

