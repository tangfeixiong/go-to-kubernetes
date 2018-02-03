# Devlopment Help

## DevOps

1. Persistent Volume Claim

```
[vagrant@kubedev-172-17-4-59 rabbit-operator]$ kubectl create -f hostpath-provisioner.yaml 
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl -n example-system get pvc
NAME                             STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
hostpath                         Bound     pvc-9081b705-07ba-11e8-beb1-525400224e72   1Mi        RWX            example-hostpath   1m
```

1. CRD

```
[vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ kubectl create -f rabbitmq-crd.yaml 
customresourcedefinition "rabbitmqs.example.com" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get crd rabbitmqs.example.com
NAME                    AGE
rabbitmqs.example.com   2m
```

1. Operator

```
[vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ kubectl create -f rabbitmq-operator-deployment.yaml 
clusterrole "rabbitmq-operator" created
serviceaccount "rabbitmq-operator" created
clusterrolebinding "rabbitmq-operator" created
deployment "rabbitmq-operator" created
service "rabbitmq-operator" created
Error from server (AlreadyExists): error when creating "rabbitmq-operator-deployment.yaml": namespaces "example-system" already exists
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl -n example-system get pods -l app=rabbitmq-operator
NAME                                 READY     STATUS    RESTARTS   AGE
rabbitmq-operator-5bd55fc4bc-rfsxh   1/1       Running   0          44s
```

1. RabbitMQ

```
[vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ kubectl create -f rabbitmq_3_7-resource.yaml        
rabbitmq "demo-rabbitmq-cluster" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl -n example-system get rabbitmqs
NAME                    AGE
demo-rabbitmq-cluster   3m
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl -n example-system get svc -l example.com/rabbitmq-operator=demo-rabbitmq-cluster
NAME                    TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)              AGE
demo-rabbitmq-cluster   ClusterIP   None         <none>        5672/TCP,15672/TCP   4m
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pods -l example.com/rabbitmq-operator=demo-rabbitmq-cluster -o wide
NAME                      READY     STATUS    RESTARTS   AGE       IP             NODE
demo-rabbitmq-cluster-0   1/1       Running   0          1m        10.244.3.77    rookdev-172-17-4-63
demo-rabbitmq-cluster-1   1/1       Running   0          27s       10.244.2.140   rookdev-172-17-4-61
demo-rabbitmq-cluster-2   1/1       Running   0          18s       10.244.3.78    rookdev-172-17-4-63
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl get pvc -l example.com/rabbitmq-operator=demo-rabbitmq-cluster
NAME                               STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
hostpath-demo-rabbitmq-cluster-0   Bound     pvc-626cb408-091d-11e8-a030-525400224e72   30Mi       RWO            example-hostpath   <invalid>
hostpath-demo-rabbitmq-cluster-1   Bound     pvc-75d9b32a-091d-11e8-a030-525400224e72   30Mi       RWO            example-hostpath   <invalid>
hostpath-demo-rabbitmq-cluster-2   Bound     pvc-7b3af68a-091d-11e8-a030-525400224e72   30Mi       RWO            example-hostpath   <invalid>
```

1. Verify

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti demo-rabbitmq-cluster-2 -- rabbitmqctl cluster_status
Cluster status of node rabbit@10.244.3.78 ...
[{nodes,[{disc,['rabbit@10.244.2.140','rabbit@10.244.3.77',
                'rabbit@10.244.3.78']}]},
 {running_nodes,['rabbit@10.244.3.77','rabbit@10.244.2.140',
                 'rabbit@10.244.3.78']},
 {cluster_name,<<"rabbit@demo-rabbitmq-cluster-0.demo-rabbitmq-cluster.default.svc.cluster.local">>},
 {partitions,[]},
 {alarms,[{'rabbit@10.244.3.77',[]},
          {'rabbit@10.244.2.140',[]},
          {'rabbit@10.244.3.78',[]}]}]
```

## CI/CD

K8s API extensions for CRD
```
[vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ ./hack/update-codegen.sh 
--------deepcopy--------
### snip...
I0203 17:06:42.287896   27978 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/apis/example.com/v1alpha1/zz_generated.deepcopy.go"
I0203 17:06:42.302960   27978 main.go:81] Completed successfully.
--------clientset--------
### snip...
I0203 17:06:50.222116   28026 execute.go:217] Processing package "versioned", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned"
I0203 17:06:50.231549   28026 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/doc.go"
I0203 17:06:50.238348   28026 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/clientset.go"
I0203 17:06:50.243701   28026 execute.go:217] Processing package "scheme", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/scheme"
I0203 17:06:50.246202   28026 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/scheme/doc.go"
I0203 17:06:50.248248   28026 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/scheme/register.go"
I0203 17:06:50.251352   28026 execute.go:217] Processing package "v1alpha1", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/typed/example.com/v1alpha1"
I0203 17:06:50.258197   28026 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/typed/example.com/v1alpha1/amqp.go"
I0203 17:06:50.264431   28026 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/typed/example.com/v1alpha1/rabbitmq.go"
I0203 17:06:50.270850   28026 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/typed/example.com/v1alpha1/example.com_client.go"
I0203 17:06:50.282279   28026 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/typed/example.com/v1alpha1/generated_expansion.go"
I0203 17:06:50.285578   28026 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/typed/example.com/v1alpha1/doc.go"
--------listers--------
### snip...
I0203 17:06:58.232922   28061 execute.go:217] Processing package "v1alpha1", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/listers/example.com/v1alpha1"
I0203 17:06:58.250739   28061 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/listers/example.com/v1alpha1/expansion_generated.go"
I0203 17:06:58.261496   28061 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/listers/example.com/v1alpha1/amqp.go"
I0203 17:06:58.338936   28061 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/listers/example.com/v1alpha1/rabbitmq.go"
I0203 17:06:58.401548   28061 main.go:56] Completed successfully.
--------informers--------
### snip...
I0203 17:07:05.589554   28073 execute.go:217] Processing package "v1alpha1", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/example.com/v1alpha1"
I0203 17:07:05.628040   28073 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/example.com/v1alpha1/interface.go"
I0203 17:07:05.638240   28073 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/example.com/v1alpha1/amqp.go"
I0203 17:07:05.645611   28073 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/example.com/v1alpha1/rabbitmq.go"
I0203 17:07:05.652180   28073 execute.go:217] Processing package "internalinterfaces", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/internalinterfaces"
I0203 17:07:05.654366   28073 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/internalinterfaces/factory_interfaces.go"
I0203 17:07:05.657603   28073 execute.go:217] Processing package "externalversions", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions"
I0203 17:07:05.662566   28073 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/factory.go"
I0203 17:07:05.668602   28073 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/generic.go"
I0203 17:07:05.672870   28073 execute.go:217] Processing package "example", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/example.com"
I0203 17:07:05.675435   28073 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/informers/externalversions/example.com/interface.go"
I0203 17:07:05.678781   28073 main.go:59] Completed successfully.
```

Artifacts template Assets
```
[vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go make go-bindata-artifacts 
```

CI/CD
```
[vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go make docker-push          
#@CGO_ENABLED=0 go build -a -v -o ./bin/rabbitmq-operator --installsuffix cgo -ldflags "-s" ./
runtime/internal/sys
runtime/internal/atomic
runtime
errors
### snip...
github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/controller
github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/operator
github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/ui/data/webapp
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/websocket
github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/server
github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/cmd
github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator
#@docker build --no-cache --force-rm -t docker.io/tangfeixiong/rabbitmq-operator:1802031046-gitd004677 ./
Sending build context to Docker daemon 47.32 MB
Step 1/7 : FROM busybox
 ---> 6ad733544a63
Step 2/7 : LABEL "maintainer" "tangfeixiong <tangfx128@gmail.com>" "project" "https://github.com/tangfeixiong/go-to-kubernetes" "name" "rabbitmq-operator" "version" "0.1-alpha" "created-by" '{"go":"v1.9.2","kubernetes":"v1.9.0","docker":"1.13.1"}'
 ---> Running in 5ee60e81ce6b
 ---> 7bb7a478faaf
Removing intermediate container 5ee60e81ce6b
Step 3/7 : COPY bin/rabbitmq-operator /
 ---> 2f2621e73029
Removing intermediate container ab7aed5b37b9
Step 4/7 : ENV DOCKER_API_VERSION '1.12' DOCKER_CONFIG_JSON '{"auths": {"localhost:5000": {"auth": "","email": ""}}}' REGISTRY_CERTS_JSON '{"localhost:5000": {"ca_base64": "", "crt_base64": "", "key_base64": ""}}'
 ---> Running in 0c3d695143e6
 ---> 253d799b613a
Removing intermediate container 0c3d695143e6
Step 5/7 : EXPOSE 10002 10001
 ---> Running in 9500fa3e0682
 ---> 22777fea5f1e
Removing intermediate container 9500fa3e0682
Step 6/7 : ENTRYPOINT /rabbitmq-operator serve
 ---> Running in dc1bd2bf617a
 ---> ca9755060c85
Removing intermediate container dc1bd2bf617a
Step 7/7 : CMD -v 2 --logtostderr=true
 ---> Running in 348912e0443e
 ---> 0f7053bd949e
Removing intermediate container 348912e0443e
Successfully built 0f7053bd949e
The push refers to a repository [docker.io/tangfeixiong/rabbitmq-operator]
ba643fc680fd: Pushed 
0271b8eebde3: Layer already exists 
latest: digest: sha256:d076fef5eb235fbc22f6c47c1ccce5b5299a6be40fa9e5b50f96234820aec17f size: 739
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pods -l app=rabbitmq,example.com/rabbitmq-operator=my-rabbitmq,github.com/go-to-kubernetes=rabbitmq-operator  -o wide
NAME            READY     STATUS    RESTARTS   AGE       IP             NODE
my-rabbitmq-0   1/1       Running   5          5m        10.244.2.137   rookdev-172-17-4-61
my-rabbitmq-1   1/1       Running   4          4m        10.244.3.75    rookdev-172-17-4-63
my-rabbitmq-2   1/1       Running   4          3m        10.244.2.138   rookdev-172-17-4-61
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti my-rabbitmq-0 -- rabbitmqctl cluster_status
Cluster status of node rabbit@10.244.2.137 ...
[{nodes,[{disc,['rabbit@10.244.2.137','rabbit@10.244.2.138',
                'rabbit@10.244.3.75']}]},
 {running_nodes,['rabbit@10.244.3.75','rabbit@10.244.2.138',
                 'rabbit@10.244.2.137']},
 {cluster_name,<<"rabbit@my-rabbitmq-1.my-rabbitmq.default.svc.cluster.local">>},
 {partitions,[]},
 {alarms,[{'rabbit@10.244.3.75',[]},
          {'rabbit@10.244.2.138',[]},
          {'rabbit@10.244.2.137',[]}]}]
```

## Cluster Lab

https://www.rabbitmq.com/rabbitmqctl.8.html

rabbitmqctl OPTIONS
```
-n node
Default node is “rabbit@server”, where server is the local host. On a host named “myserver.example.com”, the node name of the RabbitMQ Erlang node will usually be “rabbit@myserver” (unless RABBITMQ_NODENAME has been set to some non-default value at broker startup time). The output of “hostname -s” is usually the correct suffix to use after the “@” sign. See rabbitmq-server(8) for details of configuring the RabbitMQ broker.
-t timeout, --timeout timeout
Operation timeout in seconds. Only applicable to “list” commands. Default is infinity.
-l, --longnames
Use longnames for erlang distribution. If RabbitMQ broker uses long node names for erlang distribution, the option must be specified.
--erlang-cookie cookie
Erlang distribution cookie. If RabbitMQ node is using a custom erlang cookie value, the cookie value must be set vith this parameter.
```

3 standalones
```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/lang-learn$ kubectl get pods -o wide -l example.com/rabbitmq-operator=my-rabbitmq
NAME            READY     STATUS    RESTARTS   AGE       IP             NODE
my-rabbitmq-0   1/1       Running   0          9m        10.244.2.103   rookdev-172-17-4-61
my-rabbitmq-1   1/1       Running   0          7m        10.244.3.62    rookdev-172-17-4-63
my-rabbitmq-2   1/1       Running   0          5m        10.244.2.104   rookdev-172-17-4-61
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti my-rabbitmq-0 -- bash
```

```
root@my-rabbitmq-0:/# rabbitmqctl -n rabbit@10.244.2.103 cluster_status
Cluster status of node rabbit@10.244.2.103 ...
[{nodes,[{disc,['rabbit@10.244.2.103']}]},
 {running_nodes,['rabbit@10.244.2.103']},
 {cluster_name,<<"rabbit@my-rabbitmq-0.my-rabbitmq.default.svc.cluster.local">>},
 {partitions,[]},
 {alarms,[{'rabbit@10.244.2.103',[]}]}]
```

```
root@my-rabbitmq-0:/# rabbitmqctl -n rabbit@10.244.3.62 stop_app
Stopping rabbit application on node rabbit@10.244.3.62 ...
```

```
root@my-rabbitmq-0:/# rabbitmqctl -n rabbit@10.244.3.62 join_cluster rabbit@10.244.2.103
Clustering node rabbit@10.244.3.62 with rabbit@10.244.2.103
```

```
root@my-rabbitmq-0:/# rabbitmqctl -n rabbit@10.244.3.62 start_app
Starting node rabbit@10.244.3.62 ...
 completed with 5 plugins.
```

```
root@my-rabbitmq-0:/# rabbitmqctl -n rabbit@10.244.3.62 cluster_status
Cluster status of node rabbit@10.244.3.62 ...
[{nodes,[{disc,['rabbit@10.244.2.103','rabbit@10.244.3.62']}]},
 {running_nodes,['rabbit@10.244.2.103','rabbit@10.244.3.62']},
 {cluster_name,<<"rabbit@my-rabbitmq-0.my-rabbitmq.default.svc.cluster.local">>},
 {partitions,[]},
 {alarms,[{'rabbit@10.244.2.103',[]},{'rabbit@10.244.3.62',[]}]}]
```

```
root@my-rabbitmq-0:/# rabbitmqctl -n rabbit@10.244.2.104 stop_app
Stopping rabbit application on node rabbit@10.244.2.104 ...
```

```
root@my-rabbitmq-0:/# rabbitmqctl -n rabbit@10.244.2.104 join_cluster rabbit@10.244.2.103
Clustering node rabbit@10.244.2.104 with rabbit@10.244.2.103
```

```
root@my-rabbitmq-0:/# rabbitmqctl -n rabbit@10.244.2.104 start_app
Starting node rabbit@10.244.2.104 ...
 completed with 5 plugins.
```

```
root@my-rabbitmq-0:/# rabbitmqctl -n rabbit@10.244.2.104 cluster_status
Cluster status of node rabbit@10.244.2.104 ...
[{nodes,[{disc,['rabbit@10.244.2.103','rabbit@10.244.2.104',
                'rabbit@10.244.3.62']}]},
 {running_nodes,['rabbit@10.244.2.103','rabbit@10.244.3.62',
                 'rabbit@10.244.2.104']},
 {cluster_name,<<"rabbit@my-rabbitmq-0.my-rabbitmq.default.svc.cluster.local">>},
 {partitions,[]},
 {alarms,[{'rabbit@10.244.2.103',[]},
          {'rabbit@10.244.3.62',[]},
          {'rabbit@10.244.2.104',[]}]}]
```