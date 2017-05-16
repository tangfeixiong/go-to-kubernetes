# Go-to-Kubernetes

## Tables of Content

Documentation

* DevOps动手部署Kubernetes v1.5.7 － Hand-on: Deploying Kubernetes v1.5.7

* HandsOn - CentOS Installation

* [CentOS DevOps - CentOS Everything ISO media repo](./docs/centos-devops-iso-repo.md)

* [CentOS DevOps - YUM repo mirror](./docs/centos-devops-sync-repo.md)

* [CentOS HandsOn - Docker installation](./docs/centos-devops-install-docker.md)

* [CentOS DevOps - Etcd2 installation](./docs/kubernetes-devops-install-etcd.md)

* [CentOS DevOps - Kubernetes master installation](./docs/kubernetes-devops-install-master.md)

* [CentOS DevOps - GlusterFS](./docs/glusterfs)

* [CentOS DevOps - Hadoop HDFS](./docs/hadoop)

* Java示例io.fabric8:kubernetes-client项目的example - [Java example of Fabric8 kubernetes-client project](./java-devel)

Examples

* 示例k8s编排Kafka集群 － [Hand-on: Deploying Kafka into Kubernetes](./examples/kafka)

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

## Intruduction of *docker exec* and *kubectl exec* command 简述docker exec和kubectl exec的远程命令机制

[terminal-emulator.md](./terminal-emulator.md)

## Web based terminal via websocket protool 基于websocket协议的web终端
