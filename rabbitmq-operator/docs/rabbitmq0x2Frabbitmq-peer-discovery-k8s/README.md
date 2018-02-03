
```
[vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ kubectl create namespace test-rabbitmq
namespace "test-rabbitmq" created
```

```
[vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ kubectl create -f docs/rabbitmq0x2Frabbitmq-peer-discovery-k8s/rabbitmq_rbac.yaml 
serviceaccount "rabbitmq" created
role "endpoint-reader" created
rolebinding "endpoint-reader" created
```

```
[vagrant@kubedev-172-17-4-59 rabbitmq-operator]$ kubectl create -f docs/rabbitmq0x2Frabbitmq-peer-discovery-k8s/rabbitmq_statefulsets.yaml 
service "rabbitmq" created
configmap "rabbitmq-config" created
statefulset "rabbitmq" created
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl get pods -n test-rabbitmq -o wide
NAME         READY     STATUS    RESTARTS   AGE         IP             NODE
rabbitmq-0   1/1       Running   0          <invalid>   10.244.3.73    rookdev-172-17-4-63
rabbitmq-1   1/1       Running   0          <invalid>   10.244.2.136   rookdev-172-17-4-61
rabbitmq-2   1/1       Running   0          <invalid>   10.244.3.74    rookdev-172-17-4-63
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl -n test-rabbitmq exec -ti rabbitmq-0 -- rabbitmqctl cluster_status
Cluster status of node rabbit@10.244.3.73 ...
[{nodes,[{disc,['rabbit@10.244.2.136','rabbit@10.244.3.73',
                'rabbit@10.244.3.74']}]},
 {running_nodes,['rabbit@10.244.3.74','rabbit@10.244.2.136',
                 'rabbit@10.244.3.73']},
 {cluster_name,<<"rabbit@rabbitmq-0.rabbitmq.test-rabbitmq.svc.cluster.local">>},
 {partitions,[]},
 {alarms,[{'rabbit@10.244.3.74',[]},
          {'rabbit@10.244.2.136',[]},
          {'rabbit@10.244.3.73',[]}]}]
```