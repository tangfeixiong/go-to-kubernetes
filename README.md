# Go-to-Kubernetes

## Tables of Content

[Documentation](./docs)

* CentOS DevOps - CentOS Everything ISO media as offline repository
* CentOS DevOps - CentOS YUM repository mirror
* CentOS HandsOn - Docker installation
* Deploy a minimal file server - Docker image serving HTTP worked as a local YUM repository mirror server
* Kubernetes v1.5.7 DevOps - bash部署
* Kubernetes v1.6 DevOps - bash部署
* Kubernetes v1.7 DevOps - bash部署
* Kubernetes v1.8 DevOps - kubeadm部署

[Kubernetes yum/apt repository mirror](./http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm)

[Clustering examples in Kubernetes](./examples)

* 示例K8s的ElasticSearch日志分析平台 —— ElasticSearch With StorageClass and StatefulSet [Hand-on](./example/elasticsearch)
* 示例k8s编排Kafka集群 —— Deploying Kafka into Kubernetes - [Hand-on](./examples/kafka)
* 以及Kafka在k8s上使用Prometheus和Grafana去度量和监控 —— Kafka metrics with Prometheus and Grafana into Kubernetes

Others

* [CentOS DevOps - GlusterFS](./docs/glusterfs)
* [CentOS DevOps - Hadoop HDFS](./docs/hadoop)
* Java示例io.fabric8:kubernetes-client项目的example - [Java example of Fabric8 kubernetes-client project](./java-devel)


## Golang

Version and Env

    $ go version
    go version go1.6.2 linux/amd64

    $ echo $GOPATH; echo $GOBIN
    /data:/go:/work
    /data/bin

## Vendoring

Using `godep`:

    $ go get https://github.com/tools/godep
    
Currently

    $ GOPATH=/work:/go:/data godep save -v -d ./cmd/c3/
    godep: Go Version: go1.6
    godep: Finding dependencies for [./cmd/c3/]
    godep: Found package: github.com/tangfeixiong/go-to-kubernetes/cmd/c3
    godep: 	Deps: bufio bytes compress/gzip compress/zlib container/heap container/list crypto crypto/hmac crypto/md5 crypto/rand crypto/rsa crypto/sha1 crypto/sha256 crypto/tls crypto/x509 crypto/x509/pkix database/sql/driver encoding encoding/base64 encoding/binary encoding/csv encoding/hex encoding/json encoding/pem encoding/xml errors expvar flag fmt github.com/golang/build/kubernetes github.com/golang/glog github.com/inconshreveable/mousetrap github.com/spf13/cobra github.com/spf13/pflag github.com/tangfeixiong/go-to-kubernetes/pkg/api github.com/tangfeixiong/go-to-kubernetes/pkg/client github.com/tangfeixiong/go-to-kubernetes/pkg/cmd github.com/tangfeixiong/go-to-kubernetes/pkg/component go/ast go/doc go/format go/parser go/token golang.org/x/build/kubernetes/api golang.org/x/net/context golang.org/x/net/context/ctxhttp gopkg.in/inf.v0 hash hash/adler32 hash/fnv io io/ioutil k8s.io/kubernetes/federation/apis/federation k8s.io/kubernetes/federation/apis/federation/install k8s.io/kubernetes/federation/apis/federation/v1beta1
    ...
    k8s.io/kubernetes/vendor/golang.org/x/net/websocket k8s.io/kubernetes/vendor/golang.org/x/oauth2 k8s.io/kubernetes/vendor/golang.org/x/oauth2/google k8s.io/kubernetes/vendor/golang.org/x/oauth2/internal k8s.io/kubernetes/vendor/golang.org/x/oauth2/jws k8s.io/kubernetes/vendor/golang.org/x/oauth2/jwt k8s.io/kubernetes/vendor/google.golang.org/cloud/compute/metadata k8s.io/kubernetes/vendor/google.golang.org/cloud/internal k8s.io/kubernetes/vendor/gopkg.in/inf.v0 k8s.io/kubernetes/vendor/gopkg.in/yaml.v2 log log/syslog math math/big math/rand mime net net/http net/http/httputil net/mail net/rpc net/textproto net/url os os/exec os/signal os/user path path/filepath reflect regexp runtime runtime/cgo runtime/debug sort strconv strings sync sync/atomic syscall text/tabwriter text/template time unicode unicode/utf16 unicode/utf8 unsafe
    godep: Computing new Godeps.json file
    ...
    godep: Rewriting paths (if necessary)

Want update, for example, update [Kubernetes forks from Openshift Origin](https://github.com/openshift/kubernetes):

    $ (TBC)

## Make binary

Go build

    $ GOPATH=/work:/go:/data go build -o $GOBIN/c3 -a -v github.com/tangfeixiong/go-to-kubernetes/cmd/c3

Or

    $ GOPATH=/work:/go:/data go install -v github.com/tangfeixiong/go-to-kubernetes/cmd/c3
    
### mod

Issue
```
go: github.com/tangfeixiong/go-to-kubernetes imports
	github.com/tangfeixiong/go-to-kubernetes/cmd/k8sec imports
	github.com/tangfeixiong/go-to-kubernetes/pkg/k8sec/server imports
	github.com/tangfeixiong/go-to-kubernetes/pkg/k8sec/agent/manager imports
	github.com/google/cadvisor/container/docker imports
	github.com/opencontainers/runc/libcontainer/configs imports
	github.com/Sirupsen/logrus: github.com/Sirupsen/logrus@v1.4.2: parsing go.mod:
	module declares its path as: github.com/sirupsen/logrus
	        but was required as: github.com/Sirupsen/logrus
Makefile:225: recipe for target 'go-install' failed
make: *** [go-install] Error 1
```

https://stackoverflow.com/questions/56032544/how-to-find-dependency-causing-sirupsen-logrus-vs-sirupsen-logrus-unexpecte

```
go: finding gopkg.in/jcmturner/rpc.v1 v1.1.0
build github.com/tangfeixiong/go-to-kubernetes: cannot load github.com/containerd/containerd/dialer: module github.com/containerd/containerd@latest found (v1.3.2), but does not contain package github.com/containerd/containerd/dialer
Makefile:225: recipe for target 'go-install' failed
make: *** [go-install] Error 1
```

```
go: finding github.com/containerd/containerd v1.0.2
build github.com/tangfeixiong/go-to-kubernetes: cannot load github.com/coreos/go-systemd/dbus: no matching versions for query "latest"
Makefile:225: recipe for target 'go-install' failed
make: *** [go-install] Error 1
```

```
github.com/godbus/dbus
github.com/coreos/go-systemd/dbus
# github.com/coreos/go-systemd/dbus
/home/vagrant/go/pkg/mod/github.com/coreos/go-systemd@v0.0.0-20190321100706-95778dfbb74e/dbus/dbus.go:127:33: cannot use dbus.SystemBusPrivate (type func() (*dbus.Conn, error)) as type func(...<T>) (*dbus.Conn, error) in argument to dbusAuthHelloConnection
/home/vagrant/go/pkg/mod/github.com/coreos/go-systemd@v0.0.0-20190321100706-95778dfbb74e/dbus/dbus.go:136:33: cannot use dbus.SessionBusPrivate (type func() (*dbus.Conn, error)) as type func(...<T>) (*dbus.Conn, error) in argument to dbusAuthHelloConnection
/home/vagrant/go/pkg/mod/github.com/coreos/go-systemd@v0.0.0-20190321100706-95778dfbb74e/dbus/dbus.go:146:42: undefined: dbus.ConnOption
/home/vagrant/go/pkg/mod/github.com/coreos/go-systemd@v0.0.0-20190321100706-95778dfbb74e/dbus/dbus.go:204:48: undefined: dbus.ConnOption
/home/vagrant/go/pkg/mod/github.com/coreos/go-systemd@v0.0.0-20190321100706-95778dfbb74e/dbus/dbus.go:224:53: undefined: dbus.ConnOption
```

```
github.com/google/cadvisor/storage/influxdb
# github.com/google/cadvisor/storage/influxdb
/home/vagrant/go/pkg/mod/github.com/google/cadvisor@v0.35.0/storage/influxdb/influxdb.go:352:12: row.Err undefined (type models.Row has no field or method Err)
/home/vagrant/go/pkg/mod/github.com/google/cadvisor@v0.35.0/storage/influxdb/influxdb.go:353:33: row.Err undefined (type models.Row has no field or method Err)
```

curl http://172.28.128.4:19753/rp/cadvisor/docker/
```
failed to get docker info: Error response from daemon: client version 1.40 is too new. Maximum supported API version is 1.39
```

## Intruduction of *docker exec* and *kubectl exec* command 简述docker exec和kubectl exec的远程命令机制

[terminal-emulator.md](./terminal-emulator.md)

## Web based terminal via websocket protool 基于websocket协议的web终端
