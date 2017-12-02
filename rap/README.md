# Development

Run
```
[vagrant@localhost rap]$ docker run --name rap --rm -p 10072:10072 docker.io/tangfeixiong/rap:0.1
Start gRPC into [::]:10071
Start gRPC Gateway into [::]:10072
```

Test
```
[vagrant@localhost go-to-kubernetes]$ rap/runtests_curl.sh vnc
{"vnc_addr":"172.17.0.4:5900","auth_token":"not_used","token":"070a1689-5372-4278-be6f-f1049b295692"}
```

For example web address
```
[vagrant@localhost go-to-kubernetes]$ ip a show eth1
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:5a:1a:a4 brd ff:ff:ff:ff:ff:ff
    inet 172.17.4.50/24 brd 172.17.4.255 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fe5a:1aa4/64 scope link 
       valid_lft forever preferred_lft forever
```

Browse http://172.17.4.50:10072/novnc/?token=070a1689-5372-4278-be6f-f1049b295692
![屏幕快照 2017-12-01 下午4.27.33.png](./屏幕快照%202017-12-01%20下午4.27.33.png)

Active noVNC
![屏幕快照 2017-12-01 下午4.25.57.png](./屏幕快照%202017-12-01%20下午4.25.57.png)


## Build

docker image
```
[vagrant@localhost rap]$ make docker-build
runtime/internal/sys
runtime/internal/atomic
runtime
errors
internal/race
sync/atomic
internal/cpu
sync
math
io
syscall
unicode/utf8
strconv
unicode
time
reflect
internal/poll
os
bytes
bufio
fmt
sort
strings
internal/nettrace
internal/singleflight
math/rand
flag
encoding/csv
context
path/filepath
net
io/ioutil
net/url
text/template/parse
text/template
math/bits
compress/flate
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/spf13/pflag
encoding/binary
hash
hash/crc32
compress/gzip
container/list
crypto/subtle
crypto/cipher
crypto/internal/cipherhw
crypto/aes
internal/syscall/unix
math/big
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/spf13/cobra
crypto
crypto/des
crypto/sha512
crypto/hmac
crypto/md5
crypto/rc4
crypto/sha1
crypto/sha256
encoding/hex
crypto/rand
crypto/elliptic
encoding/asn1
crypto/rsa
crypto/ecdsa
crypto/dsa
crypto/x509/pkix
encoding/base64
vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20
vendor/golang_org/x/crypto/poly1305
vendor/golang_org/x/crypto/chacha20poly1305
encoding/pem
crypto/x509
vendor/golang_org/x/crypto/curve25519
vendor/golang_org/x/net/http2/hpack
vendor/golang_org/x/text/transform
log
vendor/golang_org/x/text/unicode/bidi
crypto/tls
vendor/golang_org/x/text/secure/bidirule
vendor/golang_org/x/text/unicode/norm
vendor/golang_org/x/net/idna
vendor/golang_org/x/net/proxy
mime
vendor/golang_org/x/net/lex/httplex
mime/quotedprintable
net/textproto
net/http/httptrace
net/http/internal
mime/multipart
path
os/user
net/http
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/glog
encoding
unicode/utf16
encoding/json
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/proto
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/ptypes/struct
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/jsonpb
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/grpc-ecosystem/grpc-gateway/runtime/internal
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/grpc-ecosystem/grpc-gateway/utilities
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/context
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/http2/hpack
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/text/transform
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/text/unicode/bidi
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/text/secure/bidirule
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/text/unicode/norm
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/elazarl/go-bindata-assetfs
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/internal/timeseries
html
html/template
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/idna
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/lex/httplex
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/http2
text/tabwriter
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
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/ptypes/any
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/genproto/googleapis/rpc/status
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/status
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/tap
net/http/httputil
github.com/philips/grpc-gateway-example/pkg/ui/data/swagger
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc/transport
github.com/tangfeixiong/go-to-kubernetes/pkg/ui/data/novnc
github.com/tangfeixiong/go-to-kubernetes/vendor/google.golang.org/grpc
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/grpc-ecosystem/grpc-gateway/runtime
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/golang/protobuf/protoc-gen-go/descriptor
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api
github.com/tangfeixiong/go-to-kubernetes/rap/pb
github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/pborman/uuid
github.com/tangfeixiong/go-to-kubernetes/rap/pkg/gowebsockifynovnc
github.com/tangfeixiong/go-to-kubernetes/vendor/golang.org/x/net/websocket
os/signal
github.com/tangfeixiong/go-to-kubernetes/rap/pkg/server
github.com/tangfeixiong/go-to-kubernetes/rap/cmd
github.com/tangfeixiong/go-to-kubernetes/rap
docker build -t docker.io/tangfeixiong/rap:latest ./
Sending build context to Docker daemon 14.92 MB
Step 1 : FROM busybox
 ---> d20ae45477cb
Step 2 : LABEL "maintainer" "tangfeixiong <tangfx128@gmail.com>" "project" "https://github.com/tangfeixiong/go-to-kubernetes" "name" "rap" "version" "0.1" "created-by" '{"name":"go-to-kubernetesr","namespace":"default","version":"0.1"}'
 ---> Running in ed1663626dab
 ---> 1ce503113999
Removing intermediate container ed1663626dab
Step 3 : COPY bin/rap /
 ---> 448ce0f6ef0d
Removing intermediate container c557330538a1
Step 4 : EXPOSE 10071 10072
 ---> Running in 4c5ae93b25e7
 ---> 64f3f290c8dd
Removing intermediate container 4c5ae93b25e7
Step 5 : ENTRYPOINT /rap serve
 ---> Running in 3ab475adb5ec
 ---> cdfa4b258828
Removing intermediate container 3ab475adb5ec
Step 6 : CMD -v 2 --logtostderr=true
 ---> Running in 53f217b3a263
 ---> 123370af74ea
Removing intermediate container 53f217b3a263
Successfully built 123370af74ea
```

protobuf
```
[vagrant@localhost rap]$ make protoc-grpc
```