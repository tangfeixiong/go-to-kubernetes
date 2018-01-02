Build

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ pwd
/Users/fanhongling/go/src/github.com/jw-s/redis-operator
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ GOPATH=/Users/fanhongling/go go build -o bin/redis-operator ./cmd/operator
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ docker build -t joelws/redis-operator -f Dockerfile_prod .
Sending build context to Docker daemon 79.15 MB
Step 1/4 : FROM scratch
 ---> 
Step 2/4 : ADD bin/redis-operator /redis-operator
 ---> 823960b61483
Removing intermediate container ed5977824763
Step 3/4 : ENTRYPOINT /redis-operator
 ---> Running in 6ccb58cea168
 ---> d29b19cd2b4c
Removing intermediate container 6ccb58cea168
Step 4/4 : VOLUME /tmp
 ---> Running in 7b7793c731bd
 ---> 072f6f4f8676
Removing intermediate container 7b7793c731bd
Successfully built 072f6f4f8676
```

Image

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ docker images joelws/*
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
joelws/redis-operator   latest              072f6f4f8676        3 minutes ago       36.7 MB
```

Kubernetes

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl version  
Client Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T21:07:38Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T20:55:30Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get nodes
NAME                  STATUS                        ROLES     AGE       VERSION
kubedev-172-17-4-55   NotReady,SchedulingDisabled   <none>    10d       v1.9.0
kubedev-172-17-4-59   Ready                         master    11d       v1.9.0
kubedev-172-17-4-65   NotReady,SchedulingDisabled   <none>    10d       v1.9.0
rookdev-172-17-4-61   Ready                         <none>    9d        v1.9.0
rookdev-172-17-4-63   Ready                         <none>    10d       v1.9.0
```


Deploy

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl create -f contrib/kube-redis/redis-cr.yml 
customresourcedefinition "redises.operator.joelws.com" created
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get redises   
No resources found.
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get redis  
No resources found.
```


```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl create -f contrib/kube-redis/deployment.yml 
deployment "redis-operator" created
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pods -o wide
NAME                              READY     STATUS              RESTARTS   AGE       IP            NODE
gofileserver-7df6f9ccb8-kn5fx     1/1       Running             1          10d       10.244.2.24   rookdev-172-17-4-61
hello-6cb8784cb-2gzbc             1/1       Running             1          12d       10.244.0.6    kubedev-172-17-4-59
redis-operator-7dd566cfb6-w7mnr   0/1       ErrImageNeverPull   0          35s       <none>        rookdev-172-17-4-63
```

```
[vagrant@rookdev-172-17-4-63 ~]$ ssh vagrant@172.17.4.59 "docker save joelws/redis-operator" | docker load
vagrant@172.17.4.59's password: 
5a9ea3f4d6e3: Loading layer [==================================================>] 36.69 MB/36.69 MB
Loaded image: joelws/redis-operator:latest
```

```
ubuntu@rookdev-172-17-4-61:~$ ssh vagrant@172.17.4.59 "docker save joelws/redis-operator" | docker load
vagrant@172.17.4.59's password: 
5a9ea3f4d6e3: Loading layer [==================================================>] 36.69 MB/36.69 MB
Loaded image: joelws/redis-operator:latest
```

Issue

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pods -o wide
NAME                              READY     STATUS             RESTARTS   AGE       IP            NODE
gofileserver-7df6f9ccb8-kn5fx     1/1       Running            1          10d       10.244.2.24   rookdev-172-17-4-61
hello-6cb8784cb-2gzbc             1/1       Running            1          12d       10.244.0.6    kubedev-172-17-4-59
redis-operator-7dd566cfb6-c97qq   0/1       CrashLoopBackOff   2          32s       10.244.3.71   rookdev-172-17-4-63
```

```
ubuntu@rookdev-172-17-4-61:~$ docker run -ti --rm --name=redis-operator joelws/redis-operator
standard_init_linux.go:178: exec user process caused "no such file or directory"
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ GOPATH=/Users/fanhongling/go CGO_ENABLED=0 go build --installsuffix cgo -o bin/redis-operator ./cmd/operator
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ docker build -t joelws/redis-operator -f Dockerfile_prod --force-rm --no-cache .
Sending build context to Docker daemon 79.06 MB
Step 1/4 : FROM scratch
 ---> 
Step 2/4 : ADD bin/redis-operator /redis-operator
 ---> 856f290194d7
Removing intermediate container 1460dda5328b
Step 3/4 : ENTRYPOINT /redis-operator
 ---> Running in 15182ec53b3f
 ---> ad7ea5f2fbbd
Removing intermediate container 15182ec53b3f
Step 4/4 : VOLUME /tmp
 ---> Running in 2438ae2bc033
 ---> 7a75adb49d9b
Removing intermediate container 2438ae2bc033
Successfully built 7a75adb49d9b
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ docker run -ti --rm --name=redis-operator joelws/redis-operator
panic: lookup kubernetes.default.svc on 10.0.2.3:53: no such host

goroutine 1 [running]:
github.com/jw-s/redis-operator/pkg/operator/util.InClusterConfig(0xf88ae0, 0x0, 0xc42004c960)
	/Users/fanhongling/go/src/github.com/jw-s/redis-operator/pkg/operator/util/k8s.go:19 +0x149
main.main()
	/Users/fanhongling/go/src/github.com/jw-s/redis-operator/cmd/operator/main.go:26 +0x5d
```

```
[vagrant@rookdev-172-17-4-63 ~]$ ssh vagrant@172.17.4.59 "docker save joelws/redis-operator" | docker load
vagrant@172.17.4.59's password: 
9f5bf7b294af: Loading layer [==================================================>]  36.6 MB/36.6 MB
The image joelws/redis-operator:latest already exists, renaming the old one with ID sha256:072f6f4f8676af9aa7c124d846daf1b4179d6dd441d11f5f91c6c207afe16de9 to empty string
Loaded image: joelws/redis-operator:latest
```

```
ubuntu@rookdev-172-17-4-61:~$ ssh vagrant@172.17.4.59 "docker save joelws/redis-operator" | docker load
vagrant@172.17.4.59's password: 
9f5bf7b294af: Loading layer [==================================================>]  36.6 MB/36.6 MB
The image joelws/redis-operator:latest already exists, renaming the old one with ID sha256:072f6f4f8676af9aa7c124d846daf1b4179d6dd441d11f5f91c6c207afe16de9 to empty string
Loaded image: joelws/redis-operator:latest
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl delete pod redis-operator-7dd566cfb6-c97qq
pod "redis-operator-7dd566cfb6-c97qq" deleted
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pods -o wide
NAME                              READY     STATUS        RESTARTS   AGE       IP            NODE
gofileserver-7df6f9ccb8-kn5fx     1/1       Running       1          10d       10.244.2.24   rookdev-172-17-4-61
hello-6cb8784cb-2gzbc             1/1       Running       1          12d       10.244.0.6    kubedev-172-17-4-59
redis-operator-7dd566cfb6-8vtnk   1/1       Running       0          26s       10.244.2.26   rookdev-172-17-4-61
redis-operator-7dd566cfb6-c97qq   0/1       Terminating   9          41m       10.244.3.71   rookdev-172-17-4-63
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl create -f contrib/kube-redis/redis-server.yml 
redis "my-redis" created
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get redis
NAME       AGE
my-redis   15s
```

RBAC issue
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl logs po/redis-operator-7dd566cfb6-8vtnk | head
time="2017-12-30T10:52:09Z" level=info msg="Starting controller" pkg=controller
E1230 10:52:09.598053       1 reflector.go:205] github.com/jw-s/redis-operator/vendor/k8s.io/client-go/informers/factory.go:73: Failed to list *v1.Service: services is forbidden: User "system:serviceaccount:default:default" cannot list services at the cluster scope
E1230 10:52:09.602053       1 reflector.go:205] github.com/jw-s/redis-operator/vendor/k8s.io/client-go/informers/factory.go:73: Failed to list *v1.Endpoints: endpoints is forbidden: User "system:serviceaccount:default:default" cannot list endpoints at the cluster scope
E1230 10:52:09.606128       1 reflector.go:205] github.com/jw-s/redis-operator/vendor/k8s.io/client-go/informers/factory.go:73: Failed to list *v1beta1.Deployment: deployments.apps is forbidden: User "system:serviceaccount:default:default" cannot list deployments.apps at the cluster scope
E1230 10:52:09.607428       1 reflector.go:205] github.com/jw-s/redis-operator/vendor/k8s.io/client-go/informers/factory.go:73: Failed to list *v1.ConfigMap: configmaps is forbidden: User "system:serviceaccount:default:default" cannot list configmaps at the cluster scope
E1230 10:52:09.608586       1 reflector.go:205] github.com/jw-s/redis-operator/pkg/generated/informers/externalversions/factory.go:45: Failed to list *v1.Redis: redises.operator.joelws.com is forbidden: User "system:serviceaccount:default:default" cannot list redises.operator.joelws.com at the cluster scope
E1230 10:52:09.649314       1 reflector.go:205] github.com/jw-s/redis-operator/vendor/k8s.io/client-go/informers/factory.go:73: Failed to list *v1.Pod: pods is forbidden: User "system:serviceaccount:default:default" cannot list pods at the cluster scope
E1230 10:52:09.649565       1 reflector.go:205] github.com/jw-s/redis-operator/vendor/k8s.io/client-go/informers/factory.go:73: Failed to list *v1beta1.StatefulSet: statefulsets.apps is forbidden: User "system:serviceaccount:default:default" cannot list statefulsets.apps at the cluster scope
E1230 10:52:10.601634       1 reflector.go:205] github.com/jw-s/redis-operator/vendor/k8s.io/client-go/informers/factory.go:73: Failed to list *v1.Service: services is forbidden: User "system:serviceaccount:default:default" cannot list services at the cluster scope
E1230 10:52:10.603845       1 reflector.go:205] github.com/jw-s/redis-operator/vendor/k8s.io/client-go/informers/factory.go:73: Failed to list *v1.Endpoints: endpoints is forbidden: User "system:serviceaccount:default:default" cannot list endpoints at the cluster scope
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl delete -f contrib/kube-redis/redis-server.yml 
redis "my-redis" deleted
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl delete -f contrib/kube-redis/deployment.yml 
deployment "redis-operator" deleted
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/docs/jw-s0x2Fredis-operator/redis-operator.yaml 
clusterrole "redis-operator" created
serviceaccount "redis-operator" created
clusterrolebinding "redis-operator" created
deployment "redis-operator" created
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pods -o wide -l app=redis-operator
NAME                              READY     STATUS    RESTARTS   AGE       IP            NODE
redis-operator-5979b5bd8b-25k5t   1/1       Running   0          7s        10.244.3.73   rookdev-172-17-4-63
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl logs redis-operator-5979b5bd8b-25k5t 
time="2017-12-30T19:41:46Z" level=info msg="Starting controller" pkg=controller
time="2017-12-30T19:41:47Z" level=info msg="Sync completed" pkg=controller
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/docs/jw-s0x2Fredis-operator/redis-server.yaml   
redis "my-redis" created
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get redis
NAME       AGE
my-redis   22s
```

Issue
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pods -o wide -l app=redis
NAME                                      READY     STATUS              RESTARTS   AGE       IP            NODE
redis-sentinel-my-redis-d7fcc55c5-44999   1/1       Running             0          11m       10.244.3.74   rookdev-172-17-4-63
redis-sentinel-my-redis-d7fcc55c5-hj2pk   1/1       Running             0          11m       10.244.3.75   rookdev-172-17-4-63
redis-sentinel-my-redis-d7fcc55c5-wcgvm   1/1       Running             1          11m       10.244.2.33   rookdev-172-17-4-61
redis-slave-my-redis-0                    0/1       ContainerCreating   0          11m       <none>        rookdev-172-17-4-61
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl logs redis-sentinel-my-redis-d7fcc55c5-44999
1:X 30 Dec 19:47:53.458 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:X 30 Dec 19:47:53.460 # Redis version=4.0.6, bits=64, commit=00000000, modified=0, pid=1, just started
1:X 30 Dec 19:47:53.460 # Configuration loaded
1:X 30 Dec 19:47:53.461 * Running mode=sentinel, port=26379.
1:X 30 Dec 19:47:53.462 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1:X 30 Dec 19:47:53.482 # Sentinel ID is 537368f314998acd66ae7eb2ff8e8449a81290ca
1:X 30 Dec 19:47:53.483 # +monitor master my-redis 10.105.72.21 6379 quorum 2
1:X 30 Dec 19:48:23.492 # +sdown master my-redis 10.105.72.21 6379
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl logs redis-sentinel-my-redis-d7fcc55c5-hj2pk
1:X 30 Dec 19:48:04.581 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:X 30 Dec 19:48:04.581 # Redis version=4.0.6, bits=64, commit=00000000, modified=0, pid=1, just started
1:X 30 Dec 19:48:04.581 # Configuration loaded
1:X 30 Dec 19:48:04.583 * Running mode=sentinel, port=26379.
1:X 30 Dec 19:48:04.583 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1:X 30 Dec 19:48:04.595 # Sentinel ID is 08c6e67bb83ab10fa1747a1344fe1037b3d28081
1:X 30 Dec 19:48:04.595 # +monitor master my-redis 10.105.72.21 6379 quorum 2
1:X 30 Dec 19:48:34.584 # +sdown master my-redis 10.105.72.21 6379
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl logs redis-sentinel-my-redis-d7fcc55c5-wcgvm
1:X 30 Dec 19:49:48.453 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:X 30 Dec 19:49:48.453 # Redis version=4.0.6, bits=64, commit=00000000, modified=0, pid=1, just started
1:X 30 Dec 19:49:48.453 # Configuration loaded
1:X 30 Dec 19:49:48.455 * Running mode=sentinel, port=26379.
1:X 30 Dec 19:49:48.455 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1:X 30 Dec 19:49:48.476 # Sentinel ID is 1bbb6204b4cf8850ea764424956e3bd1ce33fc5c
1:X 30 Dec 19:49:48.476 # +monitor master my-redis 10.105.72.21 6379 quorum 2
1:X 30 Dec 19:50:18.460 # +sdown master my-redis 10.105.72.21 6379
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl describe pod redis-slave-my-redis-0
Name:           redis-slave-my-redis-0
Namespace:      default
Node:           rookdev-172-17-4-61/172.17.4.61
Start Time:     Sat, 30 Dec 2017 19:40:58 +0000
Labels:         app=redis
                controller-revision-hash=redis-slave-my-redis-c9d79f495
                redis_operator=my-redis
                role=slave
                statefulset.kubernetes.io/pod-name=redis-slave-my-redis-0
Annotations:    <none>
Status:         Pending
IP:             
Controlled By:  StatefulSet/redis-slave-my-redis
Containers:
  redis-slave:
    Container ID:  
    Image:         redis:4.0-alpine
    Image ID:      
    Port:          6379/TCP
    Args:
      --slaveof
      redis-master-my-redis
      6379
    State:          Waiting
      Reason:       ContainerCreating
    Ready:          False
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /data from slave-persistent-storage (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-67l4p (ro)
Conditions:
  Type           Status
  Initialized    True 
  Ready          False 
  PodScheduled   True 
Volumes:
  slave-persistent-storage:
    Type:       PersistentVolumeClaim (a reference to a PersistentVolumeClaim in the same namespace)
    ClaimName:  slave-persistent-storage-redis-slave-my-redis-0
    ReadOnly:   false
  default-token-67l4p:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-67l4p
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type     Reason                 Age                From                          Message
  ----     ------                 ----               ----                          -------
  Normal   SuccessfulMountVolume  47m                kubelet, rookdev-172-17-4-61  MountVolume.SetUp succeeded for volume "default-token-67l4p"
  Normal   Scheduled              41m                default-scheduler             Successfully assigned redis-slave-my-redis-0 to rookdev-172-17-4-61
  Warning  FailedMount            41m (x3 over 45m)  kubelet, rookdev-172-17-4-61  Unable to mount volumes for pod "redis-slave-my-redis-0_default(3392c89a-ed9a-11e7-bd70-525400224e72)": timeout expired waiting for volumes to attach/mount for pod "default"/"redis-slave-my-redis-0". list of unattached/unmounted volumes=[slave-persistent-storage]
  Normal   SuccessfulMountVolume  39m                kubelet, rookdev-172-17-4-61  MountVolume.SetUp succeeded for volume "default-token-67l4p"
  Warning  FailedMount            9m (x13 over 37m)  kubelet, rookdev-172-17-4-61  Unable to mount volumes for pod "redis-slave-my-redis-0_default(3392c89a-ed9a-11e7-bd70-525400224e72)": timeout expired waiting for volumes to attach/mount for pod "default"/"redis-slave-my-redis-0". list of unattached/unmounted volumes=[slave-persistent-storage]
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pv
No resources found.
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pvc
NAME                                              STATUS    VOLUME    CAPACITY   ACCESS MODES   STORAGECLASS   AGE
slave-persistent-storage-redis-slave-my-redis-0   Pending                                                      50m
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/docs/jw-s0x2Fredis-operator/pv.yaml 
persistentvolume "sample-local-volume" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/docs/jw-s0x2Fredis-operator/pv-1.yaml
persistentvolume "sample-local-volume-1" created
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/docs/jw-s0x2Fredis-operator/pv-2.yaml 
persistentvolume "sample-local-volume-2" created
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pods -o wide -l app=redis
NAME                                      READY     STATUS    RESTARTS   AGE       IP            NODE
redis-sentinel-my-redis-d7fcc55c5-44999   1/1       Running   0          1h        10.244.3.74   rookdev-172-17-4-63
redis-sentinel-my-redis-d7fcc55c5-hj2pk   1/1       Running   0          1h        10.244.3.75   rookdev-172-17-4-63
redis-sentinel-my-redis-d7fcc55c5-wcgvm   1/1       Running   1          1h        10.244.2.33   rookdev-172-17-4-61
redis-slave-my-redis-0                    1/1       Running   0          1h        10.244.2.39   rookdev-172-17-4-61
redis-slave-my-redis-1                    1/1       Running   0          9m        10.244.3.76   rookdev-172-17-4-63
redis-slave-my-redis-2                    1/1       Running   0          3m        10.244.2.40   rookdev-172-17-4-61
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl logs redis-slave-my-redis-0
1:C 30 Dec 20:46:10.939 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:C 30 Dec 20:46:10.939 # Redis version=4.0.6, bits=64, commit=00000000, modified=0, pid=1, just started
1:C 30 Dec 20:46:10.939 # Configuration loaded
1:S 30 Dec 20:46:10.940 * Running mode=standalone, port=6379.
1:S 30 Dec 20:46:10.940 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1:S 30 Dec 20:46:10.940 # Server initialized
1:S 30 Dec 20:46:10.940 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
1:S 30 Dec 20:46:10.940 * Ready to accept connections
1:S 30 Dec 20:46:10.940 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:46:10.943 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:47:11.227 # Timeout connecting to the MASTER...
1:S 30 Dec 20:47:11.227 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:47:11.229 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:48:12.525 # Timeout connecting to the MASTER...
1:S 30 Dec 20:48:12.525 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:48:12.527 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:49:13.833 # Timeout connecting to the MASTER...
1:S 30 Dec 20:49:13.833 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:49:13.834 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:50:14.127 # Timeout connecting to the MASTER...
1:S 30 Dec 20:50:14.127 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:50:14.129 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:51:15.452 # Timeout connecting to the MASTER...
1:S 30 Dec 20:51:15.452 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:51:15.453 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:52:16.760 # Timeout connecting to the MASTER...
1:S 30 Dec 20:52:16.760 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:52:16.761 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:53:17.042 # Timeout connecting to the MASTER...
1:S 30 Dec 20:53:17.042 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:53:17.043 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:54:18.378 # Timeout connecting to the MASTER...
1:S 30 Dec 20:54:18.378 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:54:18.379 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:55:19.661 # Timeout connecting to the MASTER...
1:S 30 Dec 20:55:19.661 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:55:19.663 * MASTER <-> SLAVE sync started
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl logs redis-slave-my-redis-1
1:C 30 Dec 21:03:51.961 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:C 30 Dec 21:03:51.963 # Redis version=4.0.6, bits=64, commit=00000000, modified=0, pid=1, just started
1:C 30 Dec 21:03:51.963 # Configuration loaded
1:S 30 Dec 21:03:51.966 * Running mode=standalone, port=6379.
1:S 30 Dec 21:03:51.966 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1:S 30 Dec 21:03:51.966 # Server initialized
1:S 30 Dec 21:03:51.967 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
1:S 30 Dec 21:03:51.967 * Ready to accept connections
1:S 30 Dec 21:03:51.967 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 21:03:51.972 * MASTER <-> SLAVE sync started
1:S 30 Dec 21:04:52.542 # Timeout connecting to the MASTER...
1:S 30 Dec 21:04:52.543 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 21:04:52.545 * MASTER <-> SLAVE sync started
1:S 30 Dec 21:05:53.117 # Timeout connecting to the MASTER...
1:S 30 Dec 21:05:53.118 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 21:05:53.120 * MASTER <-> SLAVE sync started
1:S 30 Dec 21:06:54.666 # Timeout connecting to the MASTER...
1:S 30 Dec 21:06:54.666 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 21:06:54.669 * MASTER <-> SLAVE sync started
1:S 30 Dec 21:07:55.209 # Timeout connecting to the MASTER...
1:S 30 Dec 21:07:55.210 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 21:07:55.212 * MASTER <-> SLAVE sync started
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl logs redis-slave-my-redis-2
1:C 30 Dec 20:53:06.480 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:C 30 Dec 20:53:06.483 # Redis version=4.0.6, bits=64, commit=00000000, modified=0, pid=1, just started
1:C 30 Dec 20:53:06.483 # Configuration loaded
1:S 30 Dec 20:53:06.485 * Running mode=standalone, port=6379.
1:S 30 Dec 20:53:06.485 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1:S 30 Dec 20:53:06.485 # Server initialized
1:S 30 Dec 20:53:06.485 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
1:S 30 Dec 20:53:06.486 * Ready to accept connections
1:S 30 Dec 20:53:06.486 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:53:06.490 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:54:07.797 # Timeout connecting to the MASTER...
1:S 30 Dec 20:54:07.797 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:54:07.798 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:55:08.096 # Timeout connecting to the MASTER...
1:S 30 Dec 20:55:08.097 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:55:08.098 * MASTER <-> SLAVE sync started
1:S 30 Dec 20:56:09.438 # Timeout connecting to the MASTER...
1:S 30 Dec 20:56:09.438 * Connecting to MASTER redis-master-my-redis:6379
1:S 30 Dec 20:56:09.440 * MASTER <-> SLAVE sync started
```

Issue
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pods -o wide -l app=redis
NAME                                      READY     STATUS    RESTARTS   AGE       IP            NODE
redis-sentinel-my-redis-d7fcc55c5-44999   1/1       Running   0          1h        10.244.3.74   rookdev-172-17-4-63
redis-sentinel-my-redis-d7fcc55c5-hj2pk   1/1       Running   0          1h        10.244.3.75   rookdev-172-17-4-63
redis-sentinel-my-redis-d7fcc55c5-wcgvm   1/1       Running   1          1h        10.244.2.33   rookdev-172-17-4-61
redis-slave-my-redis-0                    1/1       Running   0          1h        10.244.2.39   rookdev-172-17-4-61
redis-slave-my-redis-1                    1/1       Running   0          21m       10.244.3.76   rookdev-172-17-4-63
redis-slave-my-redis-2                    1/1       Running   0          15m       10.244.2.40   rookdev-172-17-4-61
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get svc -o wide
NAME                      TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)     AGE       SELECTOR
hello                     ClusterIP   10.104.114.17    <none>        80/TCP      12d       run=hello
kubernetes                ClusterIP   10.96.0.1        <none>        443/TCP     12d       <none>
redis-master-my-redis     ClusterIP   10.105.72.21     <none>        6379/TCP    8m        <none>
redis-sentinel-my-redis   ClusterIP   10.104.154.245   <none>        26379/TCP   8m        app=redis,redis_operator=my-redis,role=sentinel
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ echo -e "INFO\r\nQUIT\r\n" | curl telnet://10.104.154.245:26379
$1481
# Server
redis_version:4.0.6
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:55c8d64617b13f8d
redis_mode:sentinel
os:Linux 4.13.9-300.fc27.x86_64 x86_64
arch_bits:64
multiplexing_api:epoll
atomicvar_api:atomic-builtin
gcc_version:6.3.0
process_id:1
run_id:7461aeac4fb7ef8715a1aae349cba10b26d1beac
tcp_port:26379
uptime_in_seconds:5369
uptime_in_days:0
hz:17
lru_clock:4719202
executable:/data/redis-server
config_file:/data/redis.conf

# Clients
connected_clients:1
client_longest_output_list:0
client_biggest_input_buf:7
blocked_clients:0

# CPU
used_cpu_sys:5.22
used_cpu_user:2.25
used_cpu_sys_children:0.00
used_cpu_user_children:0.00

# Stats
total_connections_received:1
total_commands_processed:0
instantaneous_ops_per_sec:0
total_net_input_bytes:13
total_net_output_bytes:0
instantaneous_input_kbps:0.00
instantaneous_output_kbps:0.00
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:0
evicted_keys:0
keyspace_hits:0
keyspace_misses:0
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:0
migrate_cached_sockets:0
slave_expires_tracked_keys:0
active_defrag_hits:0
active_defrag_misses:0
active_defrag_key_hits:0
active_defrag_key_misses:0

# Sentinel
sentinel_masters:1
sentinel_tilt:0
sentinel_running_scripts:0
sentinel_scripts_queue_length:0
sentinel_simulate_failure_flags:0
master0:name=my-redis,status=sdown,address=10.105.72.21:6379,slaves=0,sentinels=1

+OK
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ echo -e "INFO\r\nQUIT\r\n" | curl telnet://10.105.72.21:6379
^C
```





Config Map
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get cm
NAME                       DATA      AGE
sentinel-config-my-redis   1         46m
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get cm -o yaml
apiVersion: v1
items:
- apiVersion: v1
  data:
    redis.conf: |2-

      dir /data
      sentinel monitor my-redis redis-master-my-redis 6379 2
      sentinel down-after-milliseconds my-redis 30000
      sentinel parallel-syncs my-redis 1
      sentinel failover-timeout my-redis 180000
  kind: ConfigMap
  metadata:
    creationTimestamp: 2017-12-30T19:46:59Z
    name: sentinel-config-my-redis
    namespace: default
    ownerReferences:
    - apiVersion: operator.joelws.com/v1
      controller: true
      kind: Redis
      name: my-redis
      uid: 20b6fd0e-ed9a-11e7-bd70-525400224e72
    resourceVersion: "604148"
    selfLink: /api/v1/namespaces/default/configmaps/sentinel-config-my-redis
    uid: 3352a9db-ed9a-11e7-bd70-525400224e72
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
```
