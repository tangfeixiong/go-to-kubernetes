

## Content

Inspired from https://github.com/Yolean/kubernetes-kafka

## Inspired from https://github.com/Yolean/kubernetes-kafka

![屏幕快照 2017-05-05 下午11.23.18.png](./屏幕快照%202017-05-05%20下午11.23.18.png)

### zookeeper

Download _zookeeper_ manifests
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples$ cd kafka/
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ mkdir -p https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/zookeeper/00namespace.yml -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper/00namespace.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    59  100    59    0     0      3      0  0:00:19  0:00:17  0:00:02    15
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/zookeeper/20zoo-service.yml -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper/20zoo-service.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   205  100   205    0     0    132      0  0:00:01  0:00:01 --:--:--   132
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/zookeeper/30service.yml -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper/30service.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   220  100   220    0     0    133      0  0:00:01  0:00:01 --:--:--   133
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/zookeeper/50zoo.yml -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper/50zoo.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1206  100  1206    0     0    232      0  0:00:05  0:00:05 --:--:--   258
```

Do some modification
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ cd https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper/
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper$ cp 50zoo.yml 50zoo.yml.origin
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper$ sed -i 's/.zoo:/.zoo.kafka.svc.cluster.local:/g' 50zoo.yml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper$ mv test.sh test.sh.origin
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper$ sed 's%-- /opt/zookeeper/bin/zkCli.sh%--namespace=kafka -- /zookeeper-3.4.9/bin/zkCli.sh%g' test.sh.origin >> test.sh
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ chmod a+x https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper/test.sh
```

Deploy _zookeeper_ as _statefulset_
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ sudo /opt/kubernetes/server/bin/kubectl create -f ./https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper/
namespace "kafka" created
service "zoo" created
service "zookeeper" created
statefulset "zoo" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper$ sudo /opt/kubernetes/server/bin/kubectl get po/zoo-0 po/zoo-1 po/zoo-2 po/zoo-3 po/zoo-4 svc/zoo svc/zookeeper statefulsets/zoo --namespace=kafka 
NAME       READY     STATUS    RESTARTS   AGE
po/zoo-0   1/1       Running   0          4m
po/zoo-1   1/1       Running   0          4m
po/zoo-2   1/1       Running   0          4m
po/zoo-3   1/1       Running   0          4m
po/zoo-4   1/1       Running   0          4m

NAME            CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
svc/zoo         None            <none>        2888/TCP,3888/TCP   16h
svc/zookeeper   10.123.248.11   <none>        2181/TCP            16h

NAME               DESIRED   CURRENT   AGE
statefulsets/zoo   5         5         4m
```

Run test
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/zookeeper$ ./test.sh 
Connecting to localhost:2181
2017-05-05 23:21:42,889 [myid:] - INFO  [main:Environment@100] - Client environment:zookeeper.version=3.4.9-1757313, built on 08/23/2016 06:50 GMT
2017-05-05 23:21:42,894 [myid:] - INFO  [main:Environment@100] - Client environment:host.name=zoo-0.zoo.kafka.svc.
2017-05-05 23:21:42,894 [myid:] - INFO  [main:Environment@100] - Client environment:java.version=1.8.0_111-internal
2017-05-05 23:21:42,897 [myid:] - INFO  [main:Environment@100] - Client environment:java.vendor=Oracle Corporation
2017-05-05 23:21:42,898 [myid:] - INFO  [main:Environment@100] - Client environment:java.home=/usr/lib/jvm/java-1.8-openjdk/jre
2017-05-05 23:21:42,898 [myid:] - INFO  [main:Environment@100] - Client environment:java.class.path=/zookeeper-3.4.9/bin/../build/classes:/zookeeper-3.4.9/bin/../build/lib/*.jar:/zookeeper-3.4.9/bin/../lib/slf4j-log4j12-1.6.1.jar:/zookeeper-3.4.9/bin/../lib/slf4j-api-1.6.1.jar:/zookeeper-3.4.9/bin/../lib/netty-3.10.5.Final.jar:/zookeeper-3.4.9/bin/../lib/log4j-1.2.16.jar:/zookeeper-3.4.9/bin/../lib/jline-0.9.94.jar:/zookeeper-3.4.9/bin/../zookeeper-3.4.9.jar:/zookeeper-3.4.9/bin/../src/java/lib/*.jar:/conf:
2017-05-05 23:21:42,898 [myid:] - INFO  [main:Environment@100] - Client environment:java.library.path=/usr/lib/jvm/java-1.8-openjdk/jre/lib/amd64/server:/usr/lib/jvm/java-1.8-openjdk/jre/lib/amd64:/usr/lib/jvm/java-1.8-openjdk/jre/../lib/amd64:/usr/java/packages/lib/amd64:/usr/lib64:/lib64:/lib:/usr/lib
2017-05-05 23:21:42,898 [myid:] - INFO  [main:Environment@100] - Client environment:java.io.tmpdir=/tmp
2017-05-05 23:21:42,899 [myid:] - INFO  [main:Environment@100] - Client environment:java.compiler=<NA>
2017-05-05 23:21:42,899 [myid:] - INFO  [main:Environment@100] - Client environment:os.name=Linux
2017-05-05 23:21:42,899 [myid:] - INFO  [main:Environment@100] - Client environment:os.arch=amd64
2017-05-05 23:21:42,899 [myid:] - INFO  [main:Environment@100] - Client environment:os.version=3.13.0-76-generic
2017-05-05 23:21:42,900 [myid:] - INFO  [main:Environment@100] - Client environment:user.name=root
2017-05-05 23:21:42,900 [myid:] - INFO  [main:Environment@100] - Client environment:user.home=/root
2017-05-05 23:21:42,900 [myid:] - INFO  [main:Environment@100] - Client environment:user.dir=/zookeeper-3.4.9
2017-05-05 23:21:42,904 [myid:] - INFO  [main:ZooKeeper@438] - Initiating client connection, connectString=localhost:2181 sessionTimeout=30000 watcher=org.apache.zookeeper.ZooKeeperMain$MyWatcher@69d0a921
2017-05-05 23:21:42,959 [myid:] - INFO  [main-SendThread(localhost:2181):ClientCnxn$SendThread@1032] - Opening socket connection to server localhost/0:0:0:0:0:0:0:1:2181. Will not attempt to authenticate using SASL (unknown error)
2017-05-05 23:21:43,079 [myid:] - INFO  [main-SendThread(localhost:2181):ClientCnxn$SendThread@876] - Socket connection established to localhost/0:0:0:0:0:0:0:1:2181, initiating session
2017-05-05 23:21:43,248 [myid:] - INFO  [main-SendThread(localhost:2181):ClientCnxn$SendThread@1299] - Session establishment complete on server localhost/0:0:0:0:0:0:0:1:2181, sessionid = 0x15bdadffe560000, negotiated timeout = 30000

WATCHER::

WatchedEvent state:SyncConnected type:None path:null
Created /foo
Connecting to localhost:2181
2017-05-05 23:21:44,093 [myid:] - INFO  [main:Environment@100] - Client environment:zookeeper.version=3.4.9-1757313, built on 08/23/2016 06:50 GMT
2017-05-05 23:21:44,097 [myid:] - INFO  [main:Environment@100] - Client environment:host.name=zoo-2.zoo.kafka.svc.
2017-05-05 23:21:44,098 [myid:] - INFO  [main:Environment@100] - Client environment:java.version=1.8.0_111-internal
2017-05-05 23:21:44,104 [myid:] - INFO  [main:Environment@100] - Client environment:java.vendor=Oracle Corporation
2017-05-05 23:21:44,104 [myid:] - INFO  [main:Environment@100] - Client environment:java.home=/usr/lib/jvm/java-1.8-openjdk/jre
2017-05-05 23:21:44,104 [myid:] - INFO  [main:Environment@100] - Client environment:java.class.path=/zookeeper-3.4.9/bin/../build/classes:/zookeeper-3.4.9/bin/../build/lib/*.jar:/zookeeper-3.4.9/bin/../lib/slf4j-log4j12-1.6.1.jar:/zookeeper-3.4.9/bin/../lib/slf4j-api-1.6.1.jar:/zookeeper-3.4.9/bin/../lib/netty-3.10.5.Final.jar:/zookeeper-3.4.9/bin/../lib/log4j-1.2.16.jar:/zookeeper-3.4.9/bin/../lib/jline-0.9.94.jar:/zookeeper-3.4.9/bin/../zookeeper-3.4.9.jar:/zookeeper-3.4.9/bin/../src/java/lib/*.jar:/conf:
2017-05-05 23:21:44,105 [myid:] - INFO  [main:Environment@100] - Client environment:java.library.path=/usr/lib/jvm/java-1.8-openjdk/jre/lib/amd64/server:/usr/lib/jvm/java-1.8-openjdk/jre/lib/amd64:/usr/lib/jvm/java-1.8-openjdk/jre/../lib/amd64:/usr/java/packages/lib/amd64:/usr/lib64:/lib64:/lib:/usr/lib
2017-05-05 23:21:44,105 [myid:] - INFO  [main:Environment@100] - Client environment:java.io.tmpdir=/tmp
2017-05-05 23:21:44,106 [myid:] - INFO  [main:Environment@100] - Client environment:java.compiler=<NA>
2017-05-05 23:21:44,106 [myid:] - INFO  [main:Environment@100] - Client environment:os.name=Linux
2017-05-05 23:21:44,107 [myid:] - INFO  [main:Environment@100] - Client environment:os.arch=amd64
2017-05-05 23:21:44,107 [myid:] - INFO  [main:Environment@100] - Client environment:os.version=3.13.0-76-generic
2017-05-05 23:21:44,108 [myid:] - INFO  [main:Environment@100] - Client environment:user.name=root
2017-05-05 23:21:44,109 [myid:] - INFO  [main:Environment@100] - Client environment:user.home=/root
2017-05-05 23:21:44,109 [myid:] - INFO  [main:Environment@100] - Client environment:user.dir=/zookeeper-3.4.9
2017-05-05 23:21:44,112 [myid:] - INFO  [main:ZooKeeper@438] - Initiating client connection, connectString=localhost:2181 sessionTimeout=30000 watcher=org.apache.zookeeper.ZooKeeperMain$MyWatcher@69d0a921
2017-05-05 23:21:44,147 [myid:] - INFO  [main-SendThread(localhost:2181):ClientCnxn$SendThread@1032] - Opening socket connection to server localhost/0:0:0:0:0:0:0:1:2181. Will not attempt to authenticate using SASL (unknown error)
2017-05-05 23:21:44,268 [myid:] - INFO  [main-SendThread(localhost:2181):ClientCnxn$SendThread@876] - Socket connection established to localhost/0:0:0:0:0:0:0:1:2181, initiating session
2017-05-05 23:21:44,409 [myid:] - INFO  [main-SendThread(localhost:2181):ClientCnxn$SendThread@1299] - Session establishment complete on server localhost/0:0:0:0:0:0:0:1:2181, sessionid = 0x35bdadfe93c0000, negotiated timeout = 30000

WATCHER::

WatchedEvent state:SyncConnected type:None path:null
bar
cZxid = 0x100000002
ctime = Fri May 05 23:21:43 GMT 2017
mZxid = 0x100000002
mtime = Fri May 05 23:21:43 GMT 2017
pZxid = 0x100000002
cversion = 0
dataVersion = 0
aclVersion = 0
ephemeralOwner = 0x0
dataLength = 3
numChildren = 0
```

### kafka

Download _kafka_ manifests
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/00namespace.yml -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/00namespace.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    59    0    59    0     0      0      0 --:--:--  0:01:27 --:--:--    12
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/10pvc.yml -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/10pvc.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   765  100   765    0     0     17      0  0:00:45  0:00:44  0:00:01   175
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/20dns.yml -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/20dns.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   240  100   240    0     0     60      0  0:00:04  0:00:03  0:00:01    60
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/50kafka.yml -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/50kafka.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   858  100   858    0     0     52      0  0:00:16  0:00:16 --:--:--   195
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/README.md -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/README.md
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  3167  100  3167    0     0    513      0  0:00:06  0:00:06 --:--:--   870
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ mkdir -p https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/bootstrap
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/bootstrap/pv-template.yml -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/bootstrap/pv-template.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   723  100   723    0     0     81      0  0:00:08  0:00:08 --:--:--   145
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/bootstrap/pv.sh -o https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/bootstrap/pv.sh
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   438  100   438    0     0     72      0  0:00:06  0:00:06 --:--:--    96
```

Do some modification
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ sed -i 's/storage: 200Gi/storage: 100Mi/g' https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/10pvc.yml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ sed -i 's/--override/--override zookeeper.connect=zookeeper.kafka.svc.cluster.local:2181 --override/;s/storage: 200Gi/storage: 100Mi/g' https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/50kafka.yml
```

Deploy _kafka_ with _PVC_ as _statefulset_
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ sudo /opt/kubernetes/server/bin/kubectl create -f ./https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster
persistentvolumeclaim "datadir-kafka-0" created
persistentvolumeclaim "datadir-kafka-1" created
persistentvolumeclaim "datadir-kafka-2" created
service "broker" created
statefulset "kafka" created
Error from server (AlreadyExists): error when creating "https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/00namespace.yml": namespaces "kafka" already exists
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ sudo /opt/kubernetes/server/bin/kubectl create -f ./https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/bootstrap/pv-template.yml 
persistentvolume "datadir-kafka-0" created
persistentvolume "datadir-kafka-1" created
persistentvolume "datadir-kafka-2" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ sudo /opt/kubernetes/server/bin/kubectl get pvc --namespace=kafka
NAME              STATUS    VOLUME            CAPACITY   ACCESSMODES   AGE
datadir-kafka-0   Bound     datadir-kafka-0   100Mi      RWO           3m
datadir-kafka-1   Bound     datadir-kafka-1   100Mi      RWO           3m
datadir-kafka-2   Bound     datadir-kafka-2   100Mi      RWO           3m
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ sudo /opt/kubernetes/server/bin/kubectl get pv --namespace=kafka
NAME              CAPACITY   ACCESSMODES   RECLAIMPOLICY   STATUS    CLAIM                   REASON    AGE
datadir-kafka-0   100Mi      RWO           Retain          Bound     kafka/datadir-kafka-0             1m
datadir-kafka-1   100Mi      RWO           Retain          Bound     kafka/datadir-kafka-1             1m
datadir-kafka-2   100Mi      RWO           Retain          Bound     kafka/datadir-kafka-2             1m
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster$ kubectl --namespace=kafka get po/kafka-0 po/kafka-1 po/kafka-2 svc/broker statefulsets/kafka
NAME         READY     STATUS    RESTARTS   AGE
po/kafka-0   1/1       Running   0          2m
po/kafka-1   1/1       Running   0          2m
po/kafka-2   1/1       Running   0          2m

NAME         CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
svc/broker   None         <none>        9092/TCP   17h

NAME                 DESIRED   CURRENT   AGE
statefulsets/kafka   3         3         2m
```

__Validation__

create topic
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/test/11topic-create-test1.yml -o 11topic-create-test1.yml.origin
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   494  100   494    0     0     78      0  0:00:06  0:00:06 --:--:--   130
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ sed 's/:0.10.0.1/-persistent:0.10.1\n        imagePullPolicy: Never/;s%\(zookeeper\):2181%\1.kafka.svc.cluster.local:2181%' 11topic-create-test1.yml.origin >> 11topic-create-test1.yml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ kubectl create -f 11topic-create-test1.yml
job "topic-create-test1" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ kubectl --namespace=kafka get jobs -l job-name=topic-create-test1
NAME                 DESIRED   SUCCESSFUL   AGE
topic-create-test1   1         1            44m
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ kubectl --namespace=kafka get pods -a -l job-name=topic-create-test1
NAME                       READY     STATUS      RESTARTS   AGE
topic-create-test1-x2zqm   0/1       Completed   0          6m
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ kubectl --namespace=kafka logs $(kubectl --namespace=kafka get pods -a -l job-name=topic-create-test1 | awk 'NR==2 {print $1}')
Created topic "test1".
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ kubectl --namespace=kafka exec -ti kafka-1 -- /opt/kafka/bin/kafka-topics.sh --topic test1 --zookeeper zookeeper.kafka.svc.cluster.local:2181 --list
test1
```

Produce & Consume

```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster$ kubectl --namespace=kafka exec -ti kafka-1 -- /opt/kafka/bin/kafka-console-producer.sh --topic test1 --broker-list broker.kafka.svc.cluster.local:9092
This is a message!
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster$ kubectl --namespace=kafka exec -t kafka-0 -- /opt/kafka/bin/kafka-console-consumer.sh --topic test1 --bootstrap-server broker.kafka.svc.cluster.local:9092 --from-beginning
This is a message!
This is a message
This is another message
This is a message
This is another message
This is a message!
^C
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster$ kubectl --namespace=kafka exec -ti kafka-2 -- /opt/kafka/bin/kafka-console-consumer.sh --topic test1 --bootstrap-server broker.kafka.svc.cluster.local:9092 --from-beginning
This is a message!
This is a message
This is another message
This is a message
This is another message
This is a message!
^CProcessed a total of 6 messages
```

Job (TBD)
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ cat >21consume-test1.json<<EOF
> {
> "apiVersion": "batch/v1",
> "kind": "Job",
> "metadata": {
>   "name": "consume-test1",
>   "namespace": "kafka"
>   },
> "spec": {
>   "template": {
>     "metadata": {
>       "name": "consume-test1"
>     },
>     "spec": {
>       "containers": [
>         {
>           "name": "kafka",
>           "image": "solsson/kafka-persistent:0.10.1",
>           "imagePullPolicy": "Never",
>           "command": ["./bin/kafka-console-consumer.sh", "--bootstrap-server", "broker.kafka.svc.cluster.local:9092", "--topic", "test1", "--from-beginning"]
>         }],
>       "restartPolicy": "Never"
>     }
>   }
> }}
> EOF
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ kubectl create -f 21consume-test1.json 
job "consume-test1" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ kubectl --namespace=kafka get jobs -l job-name=consume-test1
NAME            DESIRED   SUCCESSFUL   AGE
consume-test1   1         0            1m
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster/test$ kubectl --namespace=kafka get pods -a -l job-name=consume-test1
NAME                  READY     STATUS    RESTARTS   AGE
consume-test1-s4fdk   1/1       Running   0          5m
```

### Optional

Pull image
```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull solsson/zookeeper-statefulset:3.4.9
3.4.9: Pulling from solsson/zookeeper-statefulset
3690ec4760f9: Already exists 
cfdb77eb56b4: Already exists 
857cbad9cd9a: Pull complete 
711263dfc2db: Pull complete 
eb4bdb431d73: Pull complete 
45d8562ee836: Pull complete 
874864a3453a: Pull complete 
473d551ea64b: Pull complete 
Digest: sha256:d32b44b32009a69b3450a5216f459e504f1041f587596895219fc04cf22f5546
Status: Downloaded newer image for solsson/zookeeper-statefulset:3.4.9
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull solsson/kafka-persistent:0.10.1
0.10.1: Pulling from solsson/kafka-persistent
Digest: sha256:0719b4688b666490abf4b32a3cc5c5da7bb2d6276b47377b35de5429f783e9c2
Status: Downloaded newer image for solsson/kafka-persistent:0.10.1
```

### Advanced

Dockerfile
```
fanhonglingdeMacBook-Pro:https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster fanhongling$ mkdir docker-kafka-persistent
fanhonglingdeMacBook-Pro:https%3A%2F%2Fraw.githubusercontent.com%2fYolean%2Fkubernetes-kafka%2Fmaster fanhongling$ cd docker-kafka-persistent/
fanhonglingdeMacBook-Pro:docker-kafka-persistent fanhongling$ curl -jkSLO https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/docker-kafka-persistent/Dockerfile
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   556  100   556    0     0    216      0  0:00:02  0:00:02 --:--:--   216
fanhonglingdeMacBook-Pro:docker-kafka-persistent fanhongling$ mkdir config
fanhonglingdeMacBook-Pro:docker-kafka-persistent fanhongling$ curl -jkSL https://raw.githubusercontent.com/Yolean/kubernetes-kafka/master/docker-kafka-persistent/config/server.properties -o config/server.properties
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  5434  100  5434    0     0   2136      0  0:00:02  0:00:02 --:--:--  2136
```

## Links

https://github.com/CloudTrackInc/kubernetes-kafka
```
fanhonglingdeMacBook-Pro:kafka fanhongling$ curl -jkSL https://raw.githubusercontent.com/CloudTrackInc/kubernetes-kafka/master/zoo-service.yaml -o https%3A%2F%2Fraw.githubusercontent.com%2fCloudTrackInc%2Fkubernetes-kafka%2Fmaster/zoo-service.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   329  100   329    0     0    117      0  0:00:02  0:00:02 --:--:--   117
fanhonglingdeMacBook-Pro:kafka fanhongling$ curl -jkSL https://raw.githubusercontent.com/CloudTrackInc/kubernetes-kafka/master/zoo-rc.yaml -o https%3A%2F%2Fraw.githubusercontent.com%2fCloudTrackInc%2Fkubernetes-kafka%2Fmaster/zoo-rc.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   329  100   329    0     0    117      0  0:00:02  0:00:02 --:--:--   117
fanhonglingdeMacBook-Pro:kafka fanhongling$ curl -jkSL https://raw.githubusercontent.com/CloudTrackInc/kubernetes-kafka/master/kafka-service.yaml -o https%3A%2F%2Fraw.githubusercontent.com%2fCloudTrackInc%2Fkubernetes-kafka%2Fmaster/kafka-service.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   250  100   250    0     0     28      0  0:00:08  0:00:08 --:--:--    63
fanhonglingdeMacBook-Pro:kafka fanhongling$ curl -jkSL https://raw.githubusercontent.com/CloudTrackInc/kubernetes-kafka/master/kafka-rc.yaml -o https%3A%2F%2Fraw.githubusercontent.com%2fCloudTrackInc%2Fkubernetes-kafka%2Fmaster/kafka-rc.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1020  100  1020    0     0    155      0  0:00:06  0:00:06 --:--:--   255

```

http://www.defuze.org/archives/351-running-a-zookeeper-and-kafka-cluster-with-kubernetes-on-aws.html