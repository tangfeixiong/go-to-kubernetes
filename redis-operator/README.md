# Redis Operator under CRD and API extensions

## Demo

Step 1st
```
[vagrant@rookdev-172-17-4-59 redis-operator]$ kubectl create -f redis-crd.yaml 
customresourcedefinition "clusters.example.com" created
customresourcedefinition "redises.example.com" created
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get crd
NAME                           AGE
clusters.example.com           4h
redises.example.com            4h
```

Step 2nd
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl create -f redis-operator.yaml 
namespace "example-system" created
clusterrole "my-redis-operator-example" created
serviceaccount "redis-operator" created
clusterrolebinding "redis-operator" created
deployment "redis-operator" created
service "redis-operator" created
```

```
[vagrant@rookdev-172-17-4-59 redis-operator]$ kubectl get ns
NAME             STATUS    AGE
default          Active    38d
example-system   Active    1m
kube-public      Active    38d
kube-system      Active    38d
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get clusterrole my-redis-operator-example
NAME                        AGE
my-redis-operator-example   1h
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get sa -n example-system
NAME             SECRETS   AGE
default          1         1h
redis-operator   1         1h
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get clusterrolebindings/redis-operator -o yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  creationTimestamp: 2018-01-25T10:43:19Z
  name: redis-operator
  resourceVersion: "2004619"
  selfLink: /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/redis-operator
  uid: 8ef7ab0d-01bc-11e8-a011-525400224e72
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: my-redis-operator-example
subjects:
- kind: ServiceAccount
  name: redis-operator
  namespace: example-system
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl -n example-system get deploy
NAME             DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
redis-operator   1         1         1            1           1h
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl --namespace example-system get pods -o wide
NAME                              READY     STATUS    RESTARTS   AGE       IP             NODE
redis-operator-75d95bf5b4-vfs4k   1/1       Running   0          1m        10.244.2.187   rookdev-172-17-4-61
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl --namespace example-system logs redis-operator-75d95bf5b4-vfs4k
I0125 10:45:31.278351       1 controller.go:154] Setting up event handlers
time="2018-01-25T10:45:31Z" level=info msg="Starting controller" pkg=controller
I0125 10:45:31.280089       1 controller.go:229] Waiting for informer caches to sync
time="2018-01-25T10:45:32Z" level=info msg="Sync completed" pkg=controller
I0125 10:45:32.085241       1 controller.go:235] Starting workers
I0125 10:45:32.085261       1 controller.go:241] Started workers
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl -n example-system get svc                       
NAME             TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)     AGE
redis-operator   ClusterIP   None         <none>        10002/TCP   1h
```

Step 3rd
```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator$ kubectl create -f redis-provision-4or3_2.yaml 
cluster "demo-redis-sentinels" created
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator$ kubectl get pods -o wide
NAME                                    READY     STATUS    RESTARTS   AGE       IP             NODE
demo-redis-sentinels-0                  1/1       Running   0          1m        10.244.3.208   rookdev-172-17-4-63
demo-redis-sentinels-1                  1/1       Running   0          1m        10.244.2.190   rookdev-172-17-4-61
demo-redis-sentinels-549fcdccb4-45svm   1/1       Running   0          1m        10.244.2.189   rookdev-172-17-4-61
demo-redis-sentinels-549fcdccb4-74kb4   1/1       Running   0          1m        10.244.3.209   rookdev-172-17-4-63
demo-redis-sentinels-549fcdccb4-flw9t   1/1       Running   0          1m        10.244.2.188   rookdev-172-17-4-61
```

Step 4th
```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator$ kubectl exec -ti demo-redis-sentinels-0 -- redis-cli INFO REPLICATION
# Replication
role:master
connected_slaves:1
slave0:ip=10.244.2.190,port=6379,state=online,offset=52357,lag=0
master_replid:05dcba4bb566aacc7ab394ce9e8436861ba0e374
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:52661
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:52661
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator$ kubectl exec -ti demo-redis-sentinels-1 -- redis-cli INFO REPLICATION
# Replication
role:slave
master_host:10.244.3.208
master_port:6379
master_link_status:up
master_last_io_seconds_ago:0
master_sync_in_progress:0
slave_repl_offset:54955
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:05dcba4bb566aacc7ab394ce9e8436861ba0e374
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:54955
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:54955
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator$ kubectl exec -ti demo-redis-sentinels-549fcdccb4-45svm -- redis-cli -p 26379 SENTINEL masters
1)  1) "name"
    2) "demo-redis-sentinels"
    3) "ip"
    4) "10.244.3.208"
    5) "port"
    6) "6379"
    7) "runid"
    8) "e5a7b7fabfd90c74037e2c6129ce1c71795560b0"
    9) "flags"
   10) "master"
   11) "link-pending-commands"
   12) "0"
   13) "link-refcount"
   14) "1"
   15) "last-ping-sent"
   16) "0"
   17) "last-ok-ping-reply"
   18) "663"
   19) "last-ping-reply"
   20) "663"
   21) "down-after-milliseconds"
   22) "60000"
   23) "info-refresh"
   24) "2050"
   25) "role-reported"
   26) "master"
   27) "role-reported-time"
   28) "333268"
   29) "config-epoch"
   30) "0"
   31) "num-slaves"
   32) "1"
   33) "num-other-sentinels"
   34) "2"
   35) "quorum"
   36) "2"
   37) "failover-timeout"
   38) "180000"
   39) "parallel-syncs"
   40) "1"
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator$ kubectl exec -ti demo-redis-sentinels-549fcdccb4-74kb4 -- redis-cli -p 26379 SENTINEL master demo-redis-sentinels
 1) "name"
 2) "demo-redis-sentinels"
 3) "ip"
 4) "10.244.3.208"
 5) "port"
 6) "6379"
 7) "runid"
 8) "e5a7b7fabfd90c74037e2c6129ce1c71795560b0"
 9) "flags"
10) "master"
11) "link-pending-commands"
12) "0"
13) "link-refcount"
14) "1"
15) "last-ping-sent"
16) "0"
17) "last-ok-ping-reply"
18) "798"
19) "last-ping-reply"
20) "798"
21) "down-after-milliseconds"
22) "60000"
23) "info-refresh"
24) "1720"
25) "role-reported"
26) "master"
27) "role-reported-time"
28) "443411"
29) "config-epoch"
30) "0"
31) "num-slaves"
32) "1"
33) "num-other-sentinels"
34) "2"
35) "quorum"
36) "2"
37) "failover-timeout"
38) "180000"
39) "parallel-syncs"
40) "1"
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator$ kubectl exec -ti demo-redis-sentinels-549fcdccb4-flw9t -- redis-cli -p 26379 SENTINEL slaves demo-redis-sentinels
1)  1) "name"
    2) "10.244.2.190:6379"
    3) "ip"
    4) "10.244.2.190"
    5) "port"
    6) "6379"
    7) "runid"
    8) "6fa0bd157fd95a3f7862323a90d220c37cbb2834"
    9) "flags"
   10) "slave"
   11) "link-pending-commands"
   12) "0"
   13) "link-refcount"
   14) "1"
   15) "last-ping-sent"
   16) "0"
   17) "last-ok-ping-reply"
   18) "259"
   19) "last-ping-reply"
   20) "259"
   21) "down-after-milliseconds"
   22) "60000"
   23) "info-refresh"
   24) "1174"
   25) "role-reported"
   26) "slave"
   27) "role-reported-time"
   28) "513155"
   29) "master-link-down-time"
   30) "0"
   31) "master-link-status"
   32) "ok"
   33) "master-host"
   34) "10.244.3.208"
   35) "master-port"
   36) "6379"
   37) "slave-priority"
   38) "100"
   39) "slave-repl-offset"
   40) "117529"
```

## Development

### Kubernetes API extensions code-generator

[./k8s-api-ext-code-generator.md](./k8s-api-ext-code-generator.md)

Generated
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ls pkg/apis pkg/client     
pkg/apis:
example.com

pkg/client:
clientset  informers  listers
```

### go-bindata code generator

Refer to https://github.com/jteeuwen/go-bindata

Generate
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ make go-bindata-spec
```

Generated
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ls pkg/spec/artifact/             
artifacts.go
```

### Protobuf and gRPC code generater

Protobuf: refer to https://github.com/google/protobuf and https://github.com/golang/protobuf

gRPC: refer to https://github.com/grpc/grpc-go and https://github.com/grpc-ecosystem/grpc-gateway

Generate
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ make protoc-grpc
```

Generated
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ls pb/*.go
pb/message.pb.go  pb/service.pb.go  pb/service.pb.gw.go
```

### Build or install

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:$HOME/go make go-install
### snippets ###
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/operator
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/server
github.com/tangfeixiong/go-to-kubernetes/redis-operator/cmd
github.com/tangfeixiong/go-to-kubernetes/redis-operator
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ls -lh $GOPATH/bin/redis-operator 
-rwxr-xr-x. 1 vagrant vagrant 66M Jan  2 22:40 /Users/fanhongling/Downloads/workspace/bin/redis-operator
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ GOPATH=$GOPATH:$HOME/go make go-build
runtime/internal/sys
runtime/internal/atomic
runtime
errors
internal/race
sync/atomic
sync
io
internal/cpu
math
syscall
time
internal/poll
os
unicode/utf8
strconv
unicode
reflect
fmt
sort
flag
bytes
bufio
strings
encoding/csv
context
internal/nettrace
internal/singleflight
math/rand
net
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/spf13/pflag
### snippets ###
github.com/tangfeixiong/go-to-kubernetes/redis-operator
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ bin/redis-operator 
bin/redis-operator flag redefined: log_dir
panic: bin/redis-operator flag redefined: log_dir

goroutine 1 [running]:
flag.(*FlagSet).Var(0xc420054060, 0x2a03ce0, 0xc42000f3b0, 0x1c3b807, 0x7, 0x1c69c37, 0x2f)
	/opt/go/src/flag/flag.go:793 +0x5e1
flag.(*FlagSet).StringVar(0xc420054060, 0xc42000f3b0, 0x1c3b807, 0x7, 0x0, 0x0, 0x1c69c37, 0x2f)
	/opt/go/src/flag/flag.go:696 +0x8b
flag.(*FlagSet).String(0xc420054060, 0x1c3b807, 0x7, 0x0, 0x0, 0x1c69c37, 0x2f, 0xc42000f300)
	/opt/go/src/flag/flag.go:709 +0x8b
flag.String(0x1c3b807, 0x7, 0x0, 0x0, 0x1c69c37, 0x2f, 0xc42000f0c0)
	/opt/go/src/flag/flag.go:716 +0x69
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/jw-s/redis-operator/vendor/github.com/golang/glog.init()
	/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/jw-s/redis-operator/vendor/github.com/golang/glog/glog_file.go:41 +0x14a
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/jw-s/redis-operator/pkg/generated/clientset.init()
	<autogenerated>:1 +0x44
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/operator.init()
	<autogenerated>:1 +0x66
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/server.init()
	<autogenerated>:1 +0xbb
github.com/tangfeixiong/go-to-kubernetes/redis-operator/cmd.init()
	<autogenerated>:1 +0x62
main.init()
	<autogenerated>:1 +0x5d
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ find ../vendor -name glog
../vendor/github.com/golang/glog
../vendor/github.com/jw-s/redis-operator/vendor/github.com/golang/glog
../vendor/k8s.io/heapster/vendor/github.com/golang/glog
../vendor/k8s.io/kubernetes/vendor/github.com/golang/glog
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ rm -rf ../vendor/github.com/jw-s/redis-operator/vendor/github.com/golang/glog/ ../vendor/k8s.io/heapster/vendor/github.com/golang/glog/ ../vendor/k8s.io/kubernetes/vendor/github.com/golang/glog/
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ docker build --no-cache --force-rm -t docker.io/tangfeixiong/redis-operator:latest -f Dockerfile.busybox ./
Sending build context to Docker daemon 41.85 MB
Step 1/7 : FROM busybox
Trying to pull repository docker.io/library/busybox ... 
sha256:bbc3a03235220b170ba48a157dd097dd1379299370e1ed99ce976df0355d24f0: Pulling from docker.io/library/busybox
0ffadd58f2a6: Pull complete 
Digest: sha256:bbc3a03235220b170ba48a157dd097dd1379299370e1ed99ce976df0355d24f0
Status: Downloaded newer image for docker.io/busybox:latest
 ---> 6ad733544a63
Step 2/7 : LABEL "maintainer" "tangfeixiong <tangfx128@gmail.com>" "project" "https://github.com/tangfeixiong/go-to-kubernetes" "name" "redis-operator" "version" "0.1" "created-by" '{"go":"v1.9.2","kubernetes":"v1.8","docker":"1.13"}'
 ---> Running in 0806f6be1bb2
 ---> df5131223310
Removing intermediate container 0806f6be1bb2
Step 3/7 : COPY bin/redis-operator /
 ---> 1a64689e5b9f
Removing intermediate container cc9da498955d
Step 4/7 : ENV DOCKER_API_VERSION '1.12' DOCKER_CONFIG_JSON '{"auths": {"localhost:5000": {"auth": "","email": ""}}}' REGISTRY_CERTS_JSON '{"localhost:5000": {"ca_base64": "", "crt_base64": "", "key_base64": ""}}'
 ---> Running in efaf0df895aa
 ---> 86f61d43660c
Removing intermediate container efaf0df895aa
Step 5/7 : EXPOSE 10002 10001
 ---> Running in e0ef44783ebb
 ---> e1751cc19ad2
Removing intermediate container e0ef44783ebb
Step 6/7 : ENTRYPOINT /redis-operator serve
 ---> Running in 2d53e4470be6
 ---> 79b5a7da1a98
Removing intermediate container 2d53e4470be6
Step 7/7 : CMD -v 2 --logtostderr=true
 ---> Running in 0832f0434fd6
 ---> e08845d7fedc
Removing intermediate container 0832f0434fd6
Successfully built e08845d7fedc
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ docker images tangfeixiong/redis-operator
REPOSITORY                              TAG                 IMAGE ID            CREATED             SIZE
docker.io/tangfeixiong/redis-operator   latest              e08845d7fedc        19 minutes ago      42.8 MB
```

```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/go/src/golang.org/x/net$ ssh vagrant@172.17.4.59 "docker save docker.io/tangfeixiong/redis-operator" | docker load
vagrant@172.17.4.59's password: 
871bc905972c: Loading layer [==================================================>] 29.33 MB/29.33 MB
The image tangfeixiong/redis-operator:latest already exists, renaming the old one with ID sha256:e08845d7fedcc882cd2b1cff4d7a9d9a38b7a5a082ab7e6237f8c2644c28c619 to empty string
Loaded image: tangfeixiong/redis-operator:latest
```

```
[vagrant@rookdev-172-17-4-63 ~]$ ssh vagrant@172.17.4.59 "docker save docker.io/tangfeixiong/redis-operator" | docker load
vagrant@172.17.4.59's password: 
871bc905972c: Loading layer [==================================================>] 29.33 MB/29.33 MB
The image docker.io/tangfeixiong/redis-operator:latest already exists, renaming the old one with ID sha256:e08845d7fedcc882cd2b1cff4d7a9d9a38b7a5a082ab7e6237f8c2644c28c619 to empty string
Loaded image: docker.io/tangfeixiong/redis-operator:latest
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: tangfeixiong
Password: 
Login Succeeded
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ docker push docker.io/tangfeixiong/redis-operator
The push refers to a repository [docker.io/tangfeixiong/redis-operator]
871bc905972c: Pushed 
0271b8eebde3: Layer already exists 
latest: digest: sha256:8cc09af9be104aa08a3a5f01606135c15de841901f36fce4fc6f549adb1c7c7b size: 738
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl create -f redis-operator.yaml 
namespace "example-system" created
clusterrole "redis-operator-example" created
serviceaccount "redis-operator" created
clusterrolebinding "redis-operator" created
deployment "redis-operator" created
service "redis-operator" created
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl -n example-system get pods -o wide
NAME                             READY     STATUS    RESTARTS   AGE       IP            NODE
redis-operator-d47d9dbf6-ppmfm   1/1       Running   0          10s       10.244.2.71   rookdev-172-17-4-61
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ./runtests_curl.sh create-crd
{"recipe":{"name":"redises","group":"example.com","version":"v1alpha1","plural":"redises","singular":"redis","kind":"Redis"}}
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get crd redises.example.com
NAME                  AGE
redises.example.com   44s
```

__docker__
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go make docker-push
runtime/internal/sys
runtime/internal/atomic
runtime
errors
internal/race
sync/atomic
sync
io
internal/cpu
math
syscall
time
internal/poll
os
unicode/utf8
strconv
unicode
reflect
fmt
sort
flag
bytes
bufio
strings
encoding/csv
context
internal/nettrace
internal/singleflight
math/rand
net
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/spf13/pflag
path/filepath
io/ioutil
net/url
text/template/parse
text/template
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/spf13/cobra
hash
crypto
crypto/sha1
container/list
crypto/subtle
crypto/cipher
crypto/internal/cipherhw
crypto/aes
encoding/binary
crypto/des
math/bits
math/big
crypto/elliptic
crypto/sha512
encoding/asn1
crypto/ecdsa
crypto/hmac
crypto/md5
internal/syscall/unix
crypto/rand
crypto/rc4
crypto/rsa
crypto/sha256
crypto/dsa
crypto/x509/pkix
encoding/hex
encoding/base64
encoding/pem
crypto/x509
vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20
vendor/golang_org/x/crypto/poly1305
vendor/golang_org/x/crypto/chacha20poly1305
vendor/golang_org/x/crypto/curve25519
crypto/tls
log
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-redis/redis/internal
hash/crc32
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-redis/redis/internal/consistenthash
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-redis/redis/internal/hashtag
encoding
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-redis/redis/internal/proto
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-redis/redis/internal/pool
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-redis/redis
os/user
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/glog
unicode/utf16
encoding/json
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/gogo/protobuf/proto
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/gogo/protobuf/sortkeys
compress/flate
compress/gzip
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/mailru/easyjson/jlexer
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/mailru/easyjson/buffer
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/mailru/easyjson/jwriter
regexp/syntax
regexp
github.com/tangfeixiong/go-to-kubernetes/vendor/gopkg.in/yaml.v2
vendor/golang_org/x/net/http2/hpack
vendor/golang_org/x/text/transform
vendor/golang_org/x/text/unicode/bidi
vendor/golang_org/x/text/secure/bidirule
vendor/golang_org/x/text/unicode/norm
vendor/golang_org/x/net/idna
vendor/golang_org/x/net/lex/httplex
vendor/golang_org/x/net/proxy
mime
mime/quotedprintable
net/textproto
mime/multipart
net/http/httptrace
net/http/internal
path
net/http
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-openapi/swag
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-openapi/jsonpointer
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/PuerkitoBio/urlesc
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/text/transform
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/text/unicode/bidi
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/text/secure/bidirule
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/text/unicode/norm
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/idna
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/text/width
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/PuerkitoBio/purell
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-openapi/jsonreference
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-openapi/spec
github.com/tangfeixiong/go-to-kubernetes/vendor/gopkg.in/inf.v0
hash/adler32
compress/zlib
encoding/xml
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/emicklei/go-restful/log
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/emicklei/go-restful
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kube-openapi/pkg/common
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/api/resource
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/google/gofuzz
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/third_party/forked/golang/reflect
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/conversion
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/selection
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/fields
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/sets
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/errors
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/validation/field
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/validation
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/labels
go/token
go/scanner
go/ast
go/doc
go/parser
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/conversion/queryparams
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/schema
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/runtime
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/types
runtime/debug
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/intstr
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/http2/hpack
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/lex/httplex
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/http2
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/net
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/runtime
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/wait
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/watch
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/core/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/apps/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/extensions/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/json
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/ghodss/yaml
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/json-iterator/go
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/serializer/recognizer
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/framer
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/yaml
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/serializer/json
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/serializer/protobuf
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/serializer/versioning
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/serializer
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/api/errors
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/serializer/streaming
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/version
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/pkg/version
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/rest/watch
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/clientcmd/api
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/metrics
net/http/httputil
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/gregjones/httpcache
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/google/btree
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/peterbourgon/diskv
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/gregjones/httpcache/diskcache
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/transport
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/util/cert
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/juju/ratelimit
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/clock
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/util/integer
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/util/flowcontrol
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/rest
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/emicklei/go-restful-swagger12
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/proto
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/ptypes/any
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/ptypes/duration
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/ptypes/timestamp
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/ptypes
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/googleapis/gnostic/extensions
os/exec
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/googleapis/gnostic/compiler
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/googleapis/gnostic/OpenAPIv2
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/api/equality
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/davecgh/go-spew/spew
text/tabwriter
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/diff
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/conversion/unstructured
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1/unstructured
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/api/meta
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/admissionregistration/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/apps/v1beta2
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/authentication/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/authentication/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/authorization/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/authorization/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/autoscaling/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/autoscaling/v2beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/batch/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/batch/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/batch/v2alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/certificates/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/networking/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/policy/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/rbac/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/rbac/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/rbac/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/scheduling/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/settings/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/storage/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/api/storage/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/scheme
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/discovery
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/apps/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/apps/v1beta2
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/authentication/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/authentication/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/authorization/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/authorization/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/autoscaling/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/batch/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/batch/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/batch/v2alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/certificates/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/reference
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/core/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/extensions/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/networking/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/policy/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/rbac/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/rbac/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/rbac/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/settings/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/storage/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes/typed/storage/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/kubernetes
container/heap
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/context
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/hashicorp/golang-lru/simplelru
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/hashicorp/golang-lru
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/cache
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/internalversion
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/pager
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/cache
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/util/version
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/operator-kit
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/coreos/go-systemd/journal
log/syslog
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/coreos/pkg/capnslog
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/apis/rook.io
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/model
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/apis/rook.io/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/client/clientset/versioned/scheme
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/client/clientset/versioned/typed/rook/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/client/clientset/versioned
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/util/exec
database/sql/driver
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/google/uuid
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/util/sys
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/clusterd
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/k8sutil
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/daemon/agent/flexvolume/attachment
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/daemon/ceph/client
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/go-ini/ini
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/util
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/daemon/ceph/mon
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/apimachinery
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/apimachinery/registered
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/apimachinery/announced
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/api
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/api/helper
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/api/v1/helper
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/kubelet/apis
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/cluster/ceph/mon
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/rook/client
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/cluster/api
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/cluster/ceph/mgr
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/cluster/ceph/osd
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/daemon/ceph/mds
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/pool
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/file/ceph
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/file
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/daemon/ceph/rgw
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/object/ceph
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/object
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/cluster
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/apis/componentconfig
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/api/v1/helper/qos
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/kubelet/qos
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/kubelet/types
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/master/ports
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/util/pointer
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/agent
encoding/gob
html
html/template
net/rpc
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/daemon/agent/flexvolume
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/groupcache/lru
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/mergepatch
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/third_party/forked/golang/json
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/strategicpatch
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/record
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/provisioner/controller/leaderelection/resourcelock
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/provisioner/controller/leaderelection
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/pborman/uuid
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/uuid
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/util/goroutinemap/exponentialbackoff
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/kubernetes/pkg/util/goroutinemap
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/provisioner/controller
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator/provisioner
os/signal
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/rook/rook/pkg/operator
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/ptypes/struct
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/jsonpb
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/grpc-ecosystem/grpc-gateway/runtime/internal
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/grpc-ecosystem/grpc-gateway/utilities
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/internal/timeseries
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/trace
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/codes
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/credentials
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/grpclb/grpc_lb_v1
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/grpclog
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/internal
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/keepalive
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/metadata
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/naming
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/peer
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/stats
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/genproto/googleapis/rpc/status
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/status
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/tap
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/transport
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/grpc-ecosystem/grpc-gateway/runtime
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/protoc-gen-go/descriptor
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pb
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/scheme
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/typed/example.com/v1
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/internalinterfaces
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/listers/example.com/v1
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/example.com/v1
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/example.com
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/sys/unix
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/crypto/ssh/terminal
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/sirupsen/logrus
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/artifact
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/apimachinery/pkg/util/rand
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/deploy
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/sts
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/svc
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/internalinterfaces
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/admissionregistration/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/admissionregistration/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/admissionregistration
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/apps/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/apps/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/apps/v1beta2
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/apps/v1beta2
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/apps
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/autoscaling/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/autoscaling/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/autoscaling/v2beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/autoscaling/v2beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/autoscaling
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/batch/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/batch/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/batch/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/batch/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/batch/v2alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/batch/v2alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/batch
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/certificates/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/certificates/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/certificates
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/core/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/core/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/core
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/extensions/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/extensions/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/extensions
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/networking/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/networking/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/networking
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/policy/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/policy/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/policy
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/rbac/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/rbac/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/rbac/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/rbac/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/rbac/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/rbac/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/rbac
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/scheduling/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/scheduling/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/scheduling
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/settings/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/settings/v1alpha1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/settings
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/storage/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/storage/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/listers/storage/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/storage/v1beta1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers/storage
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/informers
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/util/workqueue
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/controller
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/signals
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/crd
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/howeyc/gopass
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/imdario/mergo
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/auth
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/clientcmd/api/v1
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/clientcmd/api/latest
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/util/homedir
github.com/tangfeixiong/go-to-kubernetes/vendor/k8s.io/client-go/tools/clientcmd
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/operator
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/elazarl/go-bindata-assetfs
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/gorilla/websocket
github.com/philips/grpc-gateway-example/pkg/ui/data/swagger
github.com/tangfeixiong/go-to-bigdata/nps-wss/pkg/httpfs
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/ui/data/webapp
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/websocket
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/server
github.com/tangfeixiong/go-to-kubernetes/redis-operator/cmd
github.com/tangfeixiong/go-to-kubernetes/redis-operator
#@docker build --no-cache --force-rm -t docker.io/tangfeixiong/redis-operator:latest ./
Sending build context to Docker daemon 29.73 MB
Step 1/7 : FROM busybox
 ---> 6ad733544a63
Step 2/7 : LABEL "maintainer" "tangfeixiong <tangfx128@gmail.com>" "project" "https://github.com/tangfeixiong/go-to-kubernetes" "name" "redis-operator" "version" "0.1" "created-by" '{"go":"v1.9.2","kubernetes":"v1.8","docker":"1.13"}'
 ---> Running in af67cbe01e96
 ---> 483d302a0600
Removing intermediate container af67cbe01e96
Step 3/7 : COPY bin/redis-operator /
 ---> bd8c54d38399
Removing intermediate container 1eba37b4986c
Step 4/7 : ENV DOCKER_API_VERSION '1.12' DOCKER_CONFIG_JSON '{"auths": {"localhost:5000": {"auth": "","email": ""}}}' REGISTRY_CERTS_JSON '{"localhost:5000": {"ca_base64": "", "crt_base64": "", "key_base64": ""}}'
 ---> Running in 6130b9100659
 ---> 444ae0587cbe
Removing intermediate container 6130b9100659
Step 5/7 : EXPOSE 10002 10001
 ---> Running in f2c09bfb4506
 ---> 91f164436295
Removing intermediate container f2c09bfb4506
Step 6/7 : ENTRYPOINT /redis-operator serve
 ---> Running in fa1808f5c4b7
 ---> d76737597435
Removing intermediate container fa1808f5c4b7
Step 7/7 : CMD -v 2 --logtostderr=true
 ---> Running in a975fe908614
 ---> 3c2ea5ba5e44
Removing intermediate container a975fe908614
Successfully built 3c2ea5ba5e44
The push refers to a repository [docker.io/tangfeixiong/redis-operator]
4abf275618e6: Pushed 
0271b8eebde3: Layer already exists 
latest: digest: sha256:644b1223c8a767e1c091fcf621d1356b36b619bc5640bf3ca6cc3f13f75e76bc size: 738
```



```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti my-redis-0 -- redis-cli INFO replication
# Replication
role:master
connected_slaves:1
slave0:ip=10.244.3.175,port=6379,state=online,offset=154,lag=1
master_replid:33735414e49b41f60af0c06b621db6be3845e0aa
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:154
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:154
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti my-redis-1 -- redis-cli INFO replication
# Replication
role:slave
master_host:10.244.2.137
master_port:6379
master_link_status:up
master_last_io_seconds_ago:7
master_sync_in_progress:0
slave_repl_offset:168
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:33735414e49b41f60af0c06b621db6be3845e0aa
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:168
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:168
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl scale sts my-redis --replicas=3
statefulset "my-redis" scaled
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti my-redis-2 -- redis-cli INFO replication
# Replication
role:slave
master_host:10.244.2.137
master_port:6379
master_link_status:up
master_last_io_seconds_ago:10
master_sync_in_progress:0
slave_repl_offset:546
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:33735414e49b41f60af0c06b621db6be3845e0aa
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:546
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:491
repl_backlog_histlen:56
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti my-redis-0 -- redis-cli INFO replication
# Replication
role:master
connected_slaves:2
slave0:ip=10.244.3.175,port=6379,state=online,offset=574,lag=1
slave1:ip=10.244.2.138,port=6379,state=online,offset=574,lag=1
master_replid:33735414e49b41f60af0c06b621db6be3845e0aa
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:574
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:574
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti my-redis-0 -- cat redis.conf
appendonly yes
protected-mode no
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti my-redis-1 -- cat redis.conf
appendonly yes
protected-mode no
slaveof 10.244.2.137 6379
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl exec -ti my-redis-2 -- cat redis.conf
appendonly yes
protected-mode no
slaveof 10.244.2.137 6379
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ kubectl get pods -o wide
NAME                            READY     STATUS    RESTARTS   AGE       IP             NODE
my-redis-0                      1/1       Running   0          22h       10.244.2.137   rookdev-172-17-4-61
my-redis-1                      1/1       Running   0          22h       10.244.3.175   rookdev-172-17-4-63
my-redis-2                      1/1       Running   0          22h       10.244.2.138   rookdev-172-17-4-61
my-redis-cf97f57bc-78fvv        1/1       Running   0          2m        10.244.2.139   rookdev-172-17-4-61
my-redis-cf97f57bc-l62t6        1/1       Running   0          2m        10.244.3.177   rookdev-172-17-4-63
my-redis-cf97f57bc-tc9cj        1/1       Running   0          2m        10.244.3.176   rookdev-172-17-4-63
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl exec -ti my-redis-cf97f57bc-78fvv -- redis-cli -p 26379 SENTINEL masters 
1)  1) "name"
    2) "my-redis"
    3) "ip"
    4) "10.244.2.137"
    5) "port"
    6) "6379"
    7) "runid"
    8) "a064565f41ccb2f5081150538d2a22bec320865d"
    9) "flags"
   10) "master"
   11) "link-pending-commands"
   12) "0"
   13) "link-refcount"
   14) "1"
   15) "last-ping-sent"
   16) "0"
   17) "last-ok-ping-reply"
   18) "236"
   19) "last-ping-reply"
   20) "236"
   21) "down-after-milliseconds"
   22) "60000"
   23) "info-refresh"
   24) "4459"
   25) "role-reported"
   26) "master"
   27) "role-reported-time"
   28) "64756"
   29) "config-epoch"
   30) "0"
   31) "num-slaves"
   32) "2"
   33) "num-other-sentinels"
   34) "2"
   35) "quorum"
   36) "2"
   37) "failover-timeout"
   38) "180000"
   39) "parallel-syncs"
   40) "1"
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl exec -ti my-redis-cf97f57bc-l62t6 -- redis-cli -p 26379 SENTINEL masters 
1)  1) "name"
    2) "my-redis"
    3) "ip"
    4) "10.244.2.137"
    5) "port"
    6) "6379"
    7) "runid"
    8) "a064565f41ccb2f5081150538d2a22bec320865d"
    9) "flags"
   10) "master"
   11) "link-pending-commands"
   12) "0"
   13) "link-refcount"
   14) "1"
   15) "last-ping-sent"
   16) "0"
   17) "last-ok-ping-reply"
   18) "549"
   19) "last-ping-reply"
   20) "549"
   21) "down-after-milliseconds"
   22) "60000"
   23) "info-refresh"
   24) "283"
   25) "role-reported"
   26) "master"
   27) "role-reported-time"
   28) "211187"
   29) "config-epoch"
   30) "0"
   31) "num-slaves"
   32) "2"
   33) "num-other-sentinels"
   34) "2"
   35) "quorum"
   36) "2"
   37) "failover-timeout"
   38) "180000"
   39) "parallel-syncs"
   40) "1"
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl exec -ti my-redis-cf97f57bc-tc9cj -- redis-cli -p 26379 SENTINEL masters 
1)  1) "name"
    2) "my-redis"
    3) "ip"
    4) "10.244.2.137"
    5) "port"
    6) "6379"
    7) "runid"
    8) "a064565f41ccb2f5081150538d2a22bec320865d"
    9) "flags"
   10) "master"
   11) "link-pending-commands"
   12) "0"
   13) "link-refcount"
   14) "1"
   15) "last-ping-sent"
   16) "0"
   17) "last-ok-ping-reply"
   18) "802"
   19) "last-ping-reply"
   20) "802"
   21) "down-after-milliseconds"
   22) "60000"
   23) "info-refresh"
   24) "9243"
   25) "role-reported"
   26) "master"
   27) "role-reported-time"
   28) "310591"
   29) "config-epoch"
   30) "0"
   31) "num-slaves"
   32) "2"
   33) "num-other-sentinels"
   34) "2"
   35) "quorum"
   36) "2"
   37) "failover-timeout"
   38) "180000"
   39) "parallel-syncs"
   40) "1"
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl exec -ti my-redis-cf97f57bc-78fvv -- cat sentinel.conf
sentinel myid 680d72d916d0833d84800ea59a686426a9754857
sentinel monitor my-redis 10.244.2.137 6379 2
sentinel down-after-milliseconds my-redis 60000
sentinel config-epoch my-redis 0
# Generated by CONFIG REWRITE
port 26379
dir "/data"
sentinel leader-epoch my-redis 0
sentinel known-slave my-redis 10.244.3.175 6379
sentinel known-slave my-redis 10.244.2.138 6379
sentinel known-sentinel my-redis 10.244.3.176 26379 2c958fb556d42a9e78e5c82c26b771ed8887e10b
sentinel known-sentinel my-redis 10.244.3.177 26379 dc88af434430ea5b2b4dc13f302832fc42e6beec
sentinel current-epoch 0
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl exec -ti my-redis-cf97f57bc-l62t6 -- cat sentinel.conf
sentinel myid dc88af434430ea5b2b4dc13f302832fc42e6beec
sentinel monitor my-redis 10.244.2.137 6379 2
sentinel down-after-milliseconds my-redis 60000
sentinel config-epoch my-redis 0
# Generated by CONFIG REWRITE
port 26379
dir "/data"
sentinel leader-epoch my-redis 0
sentinel known-slave my-redis 10.244.2.138 6379
sentinel known-slave my-redis 10.244.3.175 6379
sentinel known-sentinel my-redis 10.244.3.176 26379 2c958fb556d42a9e78e5c82c26b771ed8887e10b
sentinel known-sentinel my-redis 10.244.2.139 26379 680d72d916d0833d84800ea59a686426a9754857
sentinel current-epoch 0
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl exec -ti my-redis-cf97f57bc-tc9cj -- cat sentinel.conf
sentinel myid 2c958fb556d42a9e78e5c82c26b771ed8887e10b
sentinel monitor my-redis 10.244.2.137 6379 2
sentinel down-after-milliseconds my-redis 60000
sentinel config-epoch my-redis 0
# Generated by CONFIG REWRITE
port 26379
dir "/data"
sentinel leader-epoch my-redis 0
sentinel known-slave my-redis 10.244.2.138 6379
sentinel known-slave my-redis 10.244.3.175 6379
sentinel known-sentinel my-redis 10.244.2.139 26379 680d72d916d0833d84800ea59a686426a9754857
sentinel known-sentinel my-redis 10.244.3.177 26379 dc88af434430ea5b2b4dc13f302832fc42e6beec
sentinel current-epoch 0
```

```
[vagrant@rookdev-172-17-4-63 ~]$ docker run --rm --name=test -ti --entrypoint="/bin/ash" docker.io/redis:4.0-alpine -c "redis-cli -h 10.102.38.21 -p 26379 SENTINEL masters"
1)  1) "name"
    2) "my-redis"
    3) "ip"
    4) "10.244.3.184"
    5) "port"
    6) "6379"
    7) "runid"
    8) "f1c118267716a052b01e66906a7d4339d2d2ebc6"
    9) "flags"
   10) "master"
   11) "link-pending-commands"
   12) "0"
   13) "link-refcount"
   14) "1"
   15) "last-ping-sent"
   16) "0"
   17) "last-ok-ping-reply"
   18) "226"
   19) "last-ping-reply"
   20) "226"
   21) "down-after-milliseconds"
   22) "60000"
   23) "info-refresh"
   24) "1328"
   25) "role-reported"
   26) "master"
   27) "role-reported-time"
   28) "3374108"
   29) "config-epoch"
   30) "0"
   31) "num-slaves"
   32) "1"
   33) "num-other-sentinels"
   34) "2"
   35) "quorum"
   36) "2"
   37) "failover-timeout"
   38) "180000"
   39) "parallel-syncs"
   40) "1"
```

### CI/CD

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ redis-operator serve --kubeconfig --logtostderr --v=5
I0125 08:36:10.463292   26211 controller.go:154] Setting up event handlers
INFO[0000] Starting controller                           pkg=controller
I0125 08:36:10.477105   26211 controller.go:229] Waiting for informer caches to sync
INFO[0000] Sync completed                                pkg=controller
I0125 08:36:10.782018   26211 controller.go:235] Starting workers
I0125 08:36:10.782555   26211 controller.go:241] Started workers
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/redis-crd.yaml 
customresourcedefinition "clusters.example.com" created
customresourcedefinition "redises.example.com" created
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl create -f /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/redis-provision-4or3_2.yaml 
cluster "demo-redis-sentinels" created
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl get svc -o wide
NAME                      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)              AGE       SELECTOR
demo-redis-sentinels      ClusterIP   None            <none>        6379/TCP,16379/TCP   11m       app=redis,component=redis,redis=demo-redis-sentinels
demo-redis-sentinels-ha   ClusterIP   10.109.188.95   <none>        26379/TCP            11m       app=redis,component=sentinel,sentinel=demo-redis-sentinels
kubernetes                ClusterIP   10.96.0.1       <none>        443/TCP              38d       <none>
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl get sts -o wide
NAME                   DESIRED   CURRENT   AGE       CONTAINERS   IMAGES
demo-redis-sentinels   2         2         11m       redis        docker.io/redis:4.0-alpine
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl get deploy -o wide
NAME                   DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE       CONTAINERS     IMAGES                                     SELECTOR
demo-redis-sentinels   3         3         3            3           11m       sentinel       docker.io/redis:4.0-alpine                 app=redis,component=sentinel,sentinel=demo-redis-sentinels
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl get pods -o wide
NAME                                    READY     STATUS    RESTARTS   AGE       IP             NODE
demo-redis-sentinels-0                  1/1       Running   0          12m       10.244.2.185   rookdev-172-17-4-61
demo-redis-sentinels-1                  1/1       Running   0          11m       10.244.3.206   rookdev-172-17-4-63
demo-redis-sentinels-549fcdccb4-8wdk5   1/1       Running   0          12m       10.244.3.205   rookdev-172-17-4-63
demo-redis-sentinels-549fcdccb4-9cr47   1/1       Running   0          12m       10.244.3.204   rookdev-172-17-4-63
demo-redis-sentinels-549fcdccb4-vrcrp   1/1       Running   0          12m       10.244.2.186   rookdev-172-17-4-61
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl exec -ti demo-redis-sentinels-1 -- redis-cli INFO REPLICATION
# Replication
role:slave
master_host:10.244.2.185
master_port:6379
master_link_status:up
master_last_io_seconds_ago:1
master_sync_in_progress:0
slave_repl_offset:21885
slave_priority:100
slave_read_only:1
connected_slaves:0
master_replid:38076048dddfb3d3876d0e88fe0c63a1a50d4501
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:21885
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:21885
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl exec -ti demo-redis-sentinels-0 -- redis-cli INFO REPLICATION
# Replication
role:master
connected_slaves:1
slave0:ip=10.244.3.206,port=6379,state=online,offset=24331,lag=0
master_replid:38076048dddfb3d3876d0e88fe0c63a1a50d4501
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:24331
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:24331
```

```
ubuntu@rookdev-172-17-4-61:~$ kubectl exec -ti demo-redis-sentinels-549fcdccb4-8wdk5 -- redis-cli -p 26379 SENTINEL masters
1)  1) "name"
    2) "demo-redis-sentinels"
    3) "ip"
    4) "10.244.2.185"
    5) "port"
    6) "6379"
    7) "runid"
    8) "7c2b0cb2a6661814cd8f09fc1cde6e6a43f25def"
    9) "flags"
   10) "master"
   11) "link-pending-commands"
   12) "0"
   13) "link-refcount"
   14) "1"
   15) "last-ping-sent"
   16) "0"
   17) "last-ok-ping-reply"
   18) "83"
   19) "last-ping-reply"
   20) "83"
   21) "down-after-milliseconds"
   22) "60000"
   23) "info-refresh"
   24) "4367"
   25) "role-reported"
   26) "master"
   27) "role-reported-time"
   28) "345936"
   29) "config-epoch"
   30) "0"
   31) "num-slaves"
   32) "1"
   33) "num-other-sentinels"
   34) "2"
   35) "quorum"
   36) "2"
   37) "failover-timeout"
   38) "180000"
   39) "parallel-syncs"
   40) "1"
```
