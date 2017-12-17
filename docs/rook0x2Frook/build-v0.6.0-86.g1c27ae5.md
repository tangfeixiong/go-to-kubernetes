# Build

## Fedora23

Helm download failure
```
[vagrant@localhost rook]$ make
build/makelib/golang.mk:106: WARNING: the source directory is not relative to the GOPATH at /home/vagrant/go:/Users/fanhongling/Downloads/workspace:/Users/fanhongling/Downloads/go-kubernetes:/Users/fanhongling/Downloads/go-openshift:/data or you are you using symlinks. The build might run into issue. Please move the source directory to be at /home/vagrant/go:/Users/fanhongling/Downloads/workspace:/Users/fanhongling/Downloads/go-kubernetes:/Users/fanhongling/Downloads/go-openshift:/data/src/github.com/rook/rook
=== installing helm

gzip: stdin: unexpected end of file
tar: Child returned status 1
tar: Error is not recoverable: exiting now
build/makelib/helm.mk:32: recipe for target '/Users/fanhongling/go/src/github.com/rook/rook/.cache/tools/linux_amd64/helm-v2.5.1' failed
make: *** [/Users/fanhongling/go/src/github.com/rook/rook/.cache/tools/linux_amd64/helm-v2.5.1] Error 2
```

Helm is archived in https://storage.googleapis.com/kubernetes-helm/

Refer to `./build/makelib/helm.mk` line 34

But networking can not arrived, required like VPN
```
[vagrant@localhost rook]$ export GOPATH=/Users/fanhongling/go
[vagrant@localhost rook]$ make -j4
=== installing helm
Creating /Users/fanhongling/go/src/github.com/rook/rook/.cache/helm 
Creating /Users/fanhongling/go/src/github.com/rook/rook/.cache/helm/repository 
Creating /Users/fanhongling/go/src/github.com/rook/rook/.cache/helm/repository/cache 
Creating /Users/fanhongling/go/src/github.com/rook/rook/.cache/helm/repository/local 
Creating /Users/fanhongling/go/src/github.com/rook/rook/.cache/helm/plugins 
Creating /Users/fanhongling/go/src/github.com/rook/rook/.cache/helm/starters 
Creating /Users/fanhongling/go/src/github.com/rook/rook/.cache/helm/cache/archive 
Creating /Users/fanhongling/go/src/github.com/rook/rook/.cache/helm/repository/repositories.yaml 
$HELM_HOME has been configured at /Users/fanhongling/go/src/github.com/rook/rook/.cache/helm.
Not installing tiller due to 'client-only' flag having been set
Happy Helming!
=== helm package rook
==> Linting /Users/fanhongling/go/src/github.com/rook/rook/cluster/charts/rook
Lint OK

1 chart(s) linted, no failures
Successfully packaged chart and saved it to: /Users/fanhongling/go/src/github.com/rook/rook/_output/charts/rook-v0.6.0-86.g1c27ae5.tgz
=== helm index
=== installing dep
waiting for lockfile /Users/fanhongling/go/pkg/dep/sm.lock: Lockfile created, but doesn't exist
waiting for lockfile /Users/fanhongling/go/pkg/dep/sm.lock: Lockfile created, but doesn't exist
waiting for lockfile /Users/fanhongling/go/pkg/dep/sm.lock: Lockfile created, but doesn't exist
waiting for lockfile /Users/fanhongling/go/pkg/dep/sm.lock: Lockfile created, but doesn't exist
waiting for lockfile /Users/fanhongling/go/pkg/dep/sm.lock: Lockfile created, but doesn't exist
waiting for lockfile /Users/fanhongling/go/pkg/dep/sm.lock: Lockfile created, but doesn't exist
^Cbuild/makelib/golang.mk:163: recipe for target 'go.vendor.lite' failed
make[1]: *** [go.vendor.lite] 中断
Makefile:104: recipe for target 'build.common' failed
make: *** [build.common] 中断

```

https://github.com/golang/dep/issues/947

Locking fails in vboxfs because hard links are not possible
```
[vagrant@localhost rook]$ export DEPNOLOCK=1
[vagrant@localhost rook]$ make -j4
=== updating vendor dependencies
grouped write of manifest, lock and vendor: error while writing out vendor tree: failed to write dep tree: failed to export golang.org/x/text: https://go.googlesource.com/text does not exist in the local cache and fetching failed: unable to get repository: : command failed: [git clone --recursive -v --progress https://go.googlesource.com/text /Users/fanhongling/go/pkg/dep/sources/https---go.googlesource.com-text]: exit status 128
build/makelib/golang.mk:163: recipe for target 'go.vendor.lite' failed
make[1]: *** [go.vendor.lite] Error 1
Makefile:104: recipe for target 'build.common' failed
make: *** [build.common] Error 2
```

VPN connection trouble...

```
[vagrant@localhost rook]$ make -j4
=== updating vendor dependencies
=== go vet
=== go build linux_amd64
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
encoding/binary
sort
strings
context
internal/nettrace
internal/singleflight
math/rand
path/filepath
net
io/ioutil
log
encoding
encoding/base64
unicode/utf16
encoding/json
github.com/rook/rook/vendor/github.com/gogo/protobuf/proto
github.com/rook/rook/vendor/github.com/coreos/go-systemd/journal
log/syslog
github.com/rook/rook/vendor/github.com/coreos/pkg/capnslog
github.com/rook/rook/vendor/github.com/gogo/protobuf/sortkeys
math/bits
compress/flate
hash
hash/crc32
github.com/rook/rook/vendor/github.com/mailru/easyjson/jlexer
compress/gzip
github.com/rook/rook/vendor/github.com/mailru/easyjson/buffer
regexp/syntax
github.com/rook/rook/vendor/github.com/mailru/easyjson/jwriter
container/list
crypto/subtle
crypto/cipher
regexp
crypto/internal/cipherhw
crypto/aes
internal/syscall/unix
math/big
github.com/rook/rook/vendor/gopkg.in/yaml.v2
crypto/rand
crypto
crypto/des
crypto/elliptic
crypto/sha512
encoding/asn1
crypto/hmac
crypto/md5
crypto/rc4
crypto/rsa
crypto/ecdsa
crypto/sha1
crypto/sha256
crypto/dsa
crypto/x509/pkix
encoding/hex
encoding/pem
vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20
crypto/x509
vendor/golang_org/x/crypto/poly1305
vendor/golang_org/x/crypto/chacha20poly1305
vendor/golang_org/x/crypto/curve25519
vendor/golang_org/x/net/http2/hpack
vendor/golang_org/x/text/transform
crypto/tls
vendor/golang_org/x/text/unicode/bidi
vendor/golang_org/x/text/secure/bidirule
vendor/golang_org/x/text/unicode/norm
vendor/golang_org/x/net/idna
net/url
vendor/golang_org/x/net/lex/httplex
vendor/golang_org/x/net/proxy
mime
mime/quotedprintable
net/textproto
net/http/httptrace
net/http/internal
mime/multipart
path
github.com/rook/rook/vendor/github.com/PuerkitoBio/urlesc
github.com/rook/rook/vendor/golang.org/x/text/transform
net/http
github.com/rook/rook/vendor/golang.org/x/text/unicode/bidi
github.com/rook/rook/vendor/golang.org/x/text/secure/bidirule
github.com/rook/rook/vendor/golang.org/x/text/unicode/norm
github.com/rook/rook/vendor/golang.org/x/net/idna
github.com/rook/rook/vendor/golang.org/x/text/width
github.com/rook/rook/vendor/github.com/PuerkitoBio/purell
encoding/csv
flag
github.com/rook/rook/vendor/github.com/spf13/pflag
github.com/rook/rook/vendor/gopkg.in/inf.v0
hash/adler32
compress/zlib
github.com/rook/rook/vendor/github.com/go-openapi/swag
encoding/xml
github.com/rook/rook/vendor/github.com/go-openapi/jsonpointer
github.com/rook/rook/vendor/github.com/go-openapi/jsonreference
github.com/rook/rook/vendor/github.com/go-openapi/spec
github.com/rook/rook/vendor/github.com/emicklei/go-restful/log
github.com/rook/rook/vendor/github.com/emicklei/go-restful
github.com/rook/rook/vendor/github.com/google/gofuzz
github.com/rook/rook/vendor/k8s.io/kube-openapi/pkg/common
github.com/rook/rook/vendor/k8s.io/apimachinery/third_party/forked/golang/reflect
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/api/resource
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/conversion
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/selection
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/fields
os/user
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/sets
github.com/rook/rook/vendor/github.com/golang/glog
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/errors
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/validation/field
go/token
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/validation
go/scanner
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/labels
go/ast
text/template/parse
go/parser
text/template
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/conversion/queryparams
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/runtime/schema
go/doc
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/types
runtime/debug
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/intstr
github.com/rook/rook/vendor/golang.org/x/net/http2/hpack
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/runtime
github.com/rook/rook/vendor/golang.org/x/net/lex/httplex
github.com/rook/rook/vendor/golang.org/x/net/http2
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/runtime
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/wait
github.com/rook/rook/vendor/github.com/spf13/cobra
github.com/rook/rook/pkg/util/flags
text/tabwriter
github.com/rook/rook/pkg/daemon/ceph/util
github.com/rook/rook/pkg/util/display
os/exec
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/net
github.com/rook/rook/pkg/util/exec
database/sql/driver
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/watch
github.com/rook/rook/vendor/github.com/google/uuid
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/apis/meta/v1
github.com/rook/rook/pkg/util/sys
github.com/rook/rook/pkg/version
github.com/rook/rook/cmd/rookctl/version
github.com/rook/rook/vendor/k8s.io/api/core/v1
github.com/rook/rook/pkg/model
github.com/rook/rook/cmd/rookctl/client
github.com/rook/rook/pkg/rook/client
github.com/rook/rook/cmd/rookctl/rook
github.com/rook/rook/cmd/rookctl/block
github.com/rook/rook/cmd/rookctl/pool
github.com/rook/rook/cmd/rookctl/filesystem
github.com/rook/rook/cmd/rookctl/node
github.com/rook/rook/cmd/rookctl/object
github.com/rook/rook/cmd/rookctl/status
github.com/rook/rook/cmd/rookctl
github.com/rook/rook/vendor/github.com/go-ini/ini
github.com/rook/rook/pkg/apis/rook.io
github.com/rook/rook/pkg/apis/rook.io/v1alpha1
github.com/rook/rook/vendor/github.com/ghodss/yaml
github.com/rook/rook/vendor/github.com/json-iterator/go
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/runtime/serializer/recognizer
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/framer
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/yaml
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/runtime/serializer/protobuf
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/runtime/serializer/versioning
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/api/errors
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/runtime/serializer/streaming
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/version
github.com/rook/rook/vendor/k8s.io/client-go/pkg/version
github.com/rook/rook/vendor/k8s.io/client-go/rest/watch
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/runtime/serializer/json
github.com/rook/rook/vendor/k8s.io/client-go/tools/clientcmd/api
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/runtime/serializer
github.com/rook/rook/vendor/k8s.io/client-go/tools/metrics
net/http/httputil
github.com/rook/rook/pkg/client/clientset/versioned/scheme
github.com/rook/rook/vendor/github.com/gregjones/httpcache
github.com/rook/rook/vendor/github.com/google/btree
github.com/rook/rook/vendor/github.com/peterbourgon/diskv
github.com/rook/rook/vendor/k8s.io/client-go/util/cert
github.com/rook/rook/vendor/github.com/juju/ratelimit
github.com/rook/rook/vendor/github.com/gregjones/httpcache/diskcache
github.com/rook/rook/vendor/k8s.io/client-go/transport
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/clock
github.com/rook/rook/vendor/k8s.io/client-go/util/integer
github.com/rook/rook/vendor/k8s.io/client-go/util/flowcontrol
github.com/rook/rook/vendor/github.com/emicklei/go-restful-swagger12
github.com/rook/rook/vendor/k8s.io/client-go/rest
github.com/rook/rook/vendor/github.com/golang/protobuf/proto
github.com/rook/rook/pkg/client/clientset/versioned/typed/rook/v1alpha1
github.com/rook/rook/vendor/github.com/golang/protobuf/ptypes/any
github.com/rook/rook/vendor/github.com/golang/protobuf/ptypes/duration
github.com/rook/rook/vendor/github.com/golang/protobuf/ptypes/timestamp
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/api/equality
github.com/rook/rook/vendor/github.com/golang/protobuf/ptypes
github.com/rook/rook/vendor/github.com/davecgh/go-spew/spew
github.com/rook/rook/vendor/github.com/googleapis/gnostic/extensions
github.com/rook/rook/vendor/github.com/googleapis/gnostic/compiler
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/diff
github.com/rook/rook/vendor/github.com/googleapis/gnostic/OpenAPIv2
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/json
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/conversion/unstructured
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/apis/meta/v1/unstructured
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/apis/meta/v1alpha1
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/api/meta
github.com/rook/rook/vendor/k8s.io/api/admissionregistration/v1alpha1
github.com/rook/rook/vendor/k8s.io/api/apps/v1beta1
github.com/rook/rook/vendor/k8s.io/api/apps/v1beta2
github.com/rook/rook/vendor/k8s.io/api/authentication/v1
github.com/rook/rook/vendor/k8s.io/api/authentication/v1beta1
github.com/rook/rook/vendor/k8s.io/api/authorization/v1
github.com/rook/rook/vendor/k8s.io/api/authorization/v1beta1
github.com/rook/rook/vendor/k8s.io/api/autoscaling/v1
github.com/rook/rook/vendor/k8s.io/api/autoscaling/v2beta1
github.com/rook/rook/vendor/k8s.io/api/batch/v1
github.com/rook/rook/vendor/k8s.io/api/certificates/v1beta1
github.com/rook/rook/vendor/k8s.io/api/extensions/v1beta1
github.com/rook/rook/vendor/k8s.io/api/batch/v1beta1
github.com/rook/rook/vendor/k8s.io/api/batch/v2alpha1
github.com/rook/rook/vendor/k8s.io/api/networking/v1
github.com/rook/rook/vendor/k8s.io/api/policy/v1beta1
github.com/rook/rook/vendor/k8s.io/api/rbac/v1
github.com/rook/rook/vendor/k8s.io/api/rbac/v1alpha1
github.com/rook/rook/vendor/k8s.io/api/rbac/v1beta1
github.com/rook/rook/vendor/k8s.io/api/scheduling/v1alpha1
github.com/rook/rook/vendor/k8s.io/api/settings/v1alpha1
github.com/rook/rook/vendor/k8s.io/api/storage/v1
github.com/rook/rook/vendor/k8s.io/api/storage/v1beta1
github.com/rook/rook/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/scheme
github.com/rook/rook/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1
github.com/rook/rook/vendor/k8s.io/client-go/discovery
github.com/rook/rook/vendor/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme
github.com/rook/rook/vendor/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1
github.com/rook/rook/pkg/client/clientset/versioned
github.com/rook/rook/vendor/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/apps/v1beta1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/apps/v1beta2
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/authentication/v1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/authentication/v1beta1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/authorization/v1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/authorization/v1beta1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/autoscaling/v1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/batch/v1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/batch/v1beta1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/batch/v2alpha1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/certificates/v1beta1
github.com/rook/rook/vendor/k8s.io/client-go/tools/reference
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/extensions/v1beta1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/core/v1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/networking/v1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/policy/v1beta1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/rbac/v1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/rbac/v1alpha1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/rbac/v1beta1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/settings/v1alpha1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/storage/v1
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes/typed/storage/v1beta1
container/heap
github.com/rook/rook/vendor/golang.org/x/net/context
github.com/rook/rook/vendor/github.com/hashicorp/golang-lru/simplelru
github.com/rook/rook/vendor/github.com/hashicorp/golang-lru
github.com/rook/rook/vendor/k8s.io/client-go/kubernetes
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/cache
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/apis/meta/internalversion
github.com/rook/rook/vendor/k8s.io/client-go/tools/pager
github.com/rook/rook/vendor/k8s.io/client-go/tools/cache
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/util/version
github.com/rook/rook/pkg/util/kvstore
github.com/rook/rook/pkg/clusterd
github.com/rook/rook/vendor/github.com/rook/operator-kit
github.com/rook/rook/pkg/operator/k8sutil
github.com/rook/rook/pkg/daemon/ceph/client
github.com/rook/rook/pkg/daemon/agent/flexvolume/attachment
github.com/rook/rook/pkg/util
github.com/rook/rook/pkg/daemon/ceph/mon
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/apimachinery
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/kubelet/apis
github.com/rook/rook/pkg/daemon/ceph/mds
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/apimachinery/registered
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/apimachinery/announced
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/api
github.com/rook/rook/pkg/operator/pool
github.com/rook/rook/pkg/daemon/ceph/rgw
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/apis/componentconfig
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/api/helper
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/api/v1/helper
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/kubelet/types
github.com/rook/rook/pkg/operator/cluster/mon
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/api/v1/helper/qos
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/kubelet/qos
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/master/ports
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/util/pointer
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1
encoding/gob
github.com/rook/rook/pkg/operator/cluster/api
github.com/rook/rook/pkg/operator/cluster/mgr
github.com/rook/rook/pkg/operator/cluster/osd
github.com/rook/rook/pkg/operator/mds
github.com/rook/rook/pkg/operator/rgw
html
html/template
github.com/rook/rook/pkg/operator/cluster
net/rpc
github.com/rook/rook/pkg/daemon/agent/flexvolume/manager/ceph
os/signal
github.com/rook/rook/vendor/github.com/gorilla/mux
expvar
github.com/rook/rook/vendor/github.com/beorn7/perks/quantile
github.com/rook/rook/vendor/github.com/prometheus/client_model/go
github.com/rook/rook/vendor/github.com/matttproud/golang_protobuf_extensions/pbutil
github.com/rook/rook/vendor/github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg
github.com/rook/rook/vendor/github.com/prometheus/common/model
github.com/rook/rook/pkg/operator/agent
github.com/rook/rook/vendor/github.com/prometheus/common/expfmt
github.com/rook/rook/vendor/github.com/prometheus/procfs/xfs
github.com/rook/rook/vendor/github.com/prometheus/procfs
github.com/rook/rook/pkg/daemon/agent/flexvolume
github.com/rook/rook/vendor/github.com/prometheus/client_golang/prometheus
github.com/rook/rook/vendor/github.com/prometheus/client_golang/prometheus/promhttp
github.com/rook/rook/pkg/daemon/ceph/collectors
github.com/rook/rook/pkg/daemon/agent/cluster
github.com/rook/rook/pkg/daemon/api
github.com/rook/rook/pkg/daemon/agent
github.com/rook/rook/pkg/daemon/ceph/mgr
github.com/rook/rook/vendor/github.com/jbw976/go-ps
github.com/rook/rook/pkg/util/proc
github.com/rook/rook/pkg/daemon/ceph/osd
github.com/rook/rook/vendor/github.com/golang/groupcache/lru
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/mergepatch
github.com/rook/rook/vendor/k8s.io/apimachinery/third_party/forked/golang/json
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/strategicpatch
github.com/rook/rook/vendor/k8s.io/client-go/tools/record
github.com/rook/rook/vendor/github.com/pborman/uuid
github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/util/uuid
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/util/goroutinemap/exponentialbackoff
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/util/goroutinemap
github.com/rook/rook/pkg/operator/provisioner/controller/leaderelection/resourcelock
github.com/rook/rook/pkg/operator/provisioner/controller/leaderelection
github.com/rook/rook/pkg/operator/provisioner/controller
github.com/rook/rook/pkg/operator/provisioner
github.com/rook/rook/pkg/operator
github.com/rook/rook/cmd/rook
github.com/rook/rook/vendor/golang.org/x/sys/unix
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/util/io
github.com/rook/rook/vendor/k8s.io/utils/exec
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/util/mount
github.com/rook/rook/vendor/k8s.io/kubernetes/pkg/volume/util
github.com/rook/rook/cmd/rookflex/cmd
github.com/rook/rook/cmd/rookflex
database/sql
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/awserr
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/internal/shareddefaults
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/credentials
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/endpoints
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/client/metadata
github.com/rook/rook/vendor/github.com/jmespath/go-jmespath
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/awsutil
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws
github.com/rook/rook/vendor/github.com/go-sql-driver/mysql
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/request
github.com/rook/rook/vendor/github.com/corpix/uarand
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/client
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/corehandlers
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/ec2metadata
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/rest
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/query/queryutil
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/signer/v4
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/query
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/service/sts
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/credentials/endpointcreds
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/defaults
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/credentials/stscreds
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/restxml
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/session
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/service/s3
github.com/rook/rook/vendor/github.com/icrowley/fake
github.com/rook/rook/vendor/github.com/pmezard/go-difflib/difflib
net/http/httptest
github.com/rook/rook/vendor/github.com/stretchr/testify/assert
github.com/rook/rook/vendor/github.com/stretchr/testify/require
runtime/trace
testing
github.com/rook/rook/tests/framework/objects
github.com/rook/rook/vendor/github.com/stretchr/testify/suite
runtime/pprof
testing/internal/testdeps
github.com/rook/rook/tests/framework/utils
github.com/rook/rook/tests/framework/contracts
github.com/rook/rook/tests/framework/installer
github.com/rook/rook/tests/framework/clients
runtime/pprof
runtime/trace
testing
testing/internal/testdeps
database/sql
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/awserr
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/internal/shareddefaults
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/credentials
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/endpoints
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/client/metadata
github.com/rook/rook/vendor/github.com/jmespath/go-jmespath
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/awsutil
github.com/rook/rook/vendor/github.com/go-sql-driver/mysql
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/request
github.com/rook/rook/vendor/github.com/corpix/uarand
github.com/rook/rook/vendor/github.com/icrowley/fake
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/client
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/corehandlers
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/rest
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/query/queryutil
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/signer/v4
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/query
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/ec2metadata
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/service/sts
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/credentials/endpointcreds
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/credentials/stscreds
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/defaults
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/private/protocol/restxml
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/aws/session
github.com/rook/rook/vendor/github.com/aws/aws-sdk-go/service/s3
github.com/rook/rook/vendor/github.com/pmezard/go-difflib/difflib
net/http/httptest
github.com/rook/rook/vendor/github.com/stretchr/testify/assert
github.com/rook/rook/vendor/github.com/stretchr/testify/require
github.com/rook/rook/tests/framework/objects
github.com/rook/rook/vendor/github.com/stretchr/testify/suite
github.com/rook/rook/tests/framework/utils
github.com/rook/rook/tests/framework/contracts
github.com/rook/rook/tests/framework/installer
github.com/rook/rook/tests/framework/clients
github.com/rook/rook/tests/integration
testmain
github.com/rook/rook/tests/longhaul
testmain
=== docker build build-295cc74b/rootfs-builder-amd64
sha256:d459ce2d4f4e5374b9979a884a6a30384f7aa55b94663904f21caa3164168a94
=== docker build build-295cc74b/rootfs-amd64-d459ce2d4f4e
6b597271595233476785d19cd4e2b0760b710aa7147a1f38a5ea651759573c20
sha256:c7f1b0195b6e39010c5c6802ddfe130dd8fb8bb06cfbda46ba44d904d9775c13
=== docker build build-295cc74b/base-amd64
sha256:caf35b84e1a6417d060040e92530c638782a38486da666c40452860aaeb1d03f
=== caching image build-295cc74b/base-amd64
=== caching image build-295cc74b/rootfs-builder-amd64
=== caching image build-295cc74b/rootfs-amd64-d459ce2d4f4e
=== docker build build-295cc74b/cross-gnu-amd64
sha256:56e4e165cca43ae0a55b89692c17c4ba49cff0d6688daae8117425f9d4f9f06b
=== caching image build-295cc74b/cross-gnu-amd64
=== docker build build-295cc74b/ceph-builder-amd64-tcmalloc-release
Sending build context to Docker daemon 14.85 kB
Step 1 : FROM build-295cc74b/cross-gnu-amd64
 ---> 56e4e165cca4
Step 2 : ARG CEPH_GIT_REPO
 ---> Running in 8c13c180cdf8
 ---> c7a31b72694e
Removing intermediate container 8c13c180cdf8
Step 3 : RUN git clone ${CEPH_GIT_REPO} ceph
 ---> Running in 1916f6ecd0eb
Cloning into 'ceph'...
 ---> 55c752753c92
Removing intermediate container 1916f6ecd0eb
Step 4 : ARG ARCH
 ---> Running in ae9dc7510f26
 ---> c5ba7a0596ce
Removing intermediate container ae9dc7510f26
Step 5 : ARG CROSS_TRIPLE
 ---> Running in b648911953b3
 ---> d401f747a50d
Removing intermediate container b648911953b3
Step 6 : RUN DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         python-setuptools         gperf                 cython:${ARCH}         python:${ARCH}         libaio-dev:${ARCH}         libatomic-ops-dev:${ARCH}         libbabeltrace-dev:${ARCH}         libblkid-dev:${ARCH}         libboost-context-dev:${ARCH}         libboost-coroutine-dev:${ARCH}         libboost-date-time-dev:${ARCH}         libboost-filesystem-dev:${ARCH}         libboost-iostreams-dev:${ARCH}         libboost-program-options-dev:${ARCH}         libboost-python-dev:${ARCH}         libboost-random-dev:${ARCH}         libboost-regex-dev:${ARCH}         libboost-system-dev:${ARCH}         libboost-thread-dev:${ARCH}         libcurl4-gnutls-dev:${ARCH}         libexpat1-dev:${ARCH}         libgoogle-perftools-dev:${ARCH}         libgoogle-perftools4:${ARCH}         libibverbs-dev:${ARCH}         libjemalloc-dev:${ARCH}         libkeyutils-dev:${ARCH}         libldap2-dev:${ARCH}         libnss3-dev:${ARCH}         libpython-dev:${ARCH}         libsnappy-dev:${ARCH}         libssl-dev:${ARCH}         libtcmalloc-minimal4:${ARCH}         libudev-dev:${ARCH}         libunwind-dev:${ARCH}         zlib1g-dev:${ARCH} &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Running in 554996be9533
Get:1 http://archive.ubuntu.com/ubuntu zesty InRelease [243 kB]
Get:2 http://ports.ubuntu.com zesty InRelease [243 kB]
Get:3 http://archive.ubuntu.com/ubuntu zesty-updates InRelease [89.2 kB]
Get:4 http://ports.ubuntu.com zesty-updates InRelease [89.2 kB]
Get:5 http://ports.ubuntu.com zesty-security InRelease [89.2 kB]
Get:6 http://ports.ubuntu.com zesty-backports InRelease [89.2 kB]
Get:7 http://archive.ubuntu.com/ubuntu zesty-security InRelease [89.2 kB]
Get:8 http://archive.ubuntu.com/ubuntu zesty-backports InRelease [89.2 kB]
Get:9 http://ports.ubuntu.com zesty/main arm64 Packages [1523 kB]
Get:10 http://archive.ubuntu.com/ubuntu zesty/main amd64 Packages [1574 kB]
Get:11 http://archive.ubuntu.com/ubuntu zesty/restricted amd64 Packages [14.3 kB]
Get:12 http://archive.ubuntu.com/ubuntu zesty/universe amd64 Packages [10.5 MB]
Get:13 http://ports.ubuntu.com zesty/restricted armhf Packages [9242 B]
Get:14 http://ports.ubuntu.com zesty/universe armhf Packages [10.2 MB]
Get:15 http://ports.ubuntu.com zesty/main armhf Packages [1525 kB]
Get:16 http://ports.ubuntu.com zesty/universe arm64 Packages [10.2 MB]
Get:17 http://archive.ubuntu.com/ubuntu zesty-updates/restricted amd64 Packages [3604 B]
Get:18 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 Packages [321 kB]
Get:19 http://archive.ubuntu.com/ubuntu zesty-updates/universe amd64 Packages [217 kB]
Get:20 http://archive.ubuntu.com/ubuntu zesty-security/universe amd64 Packages [112 kB]
Get:21 http://archive.ubuntu.com/ubuntu zesty-security/restricted amd64 Packages [3221 B]
Get:22 http://archive.ubuntu.com/ubuntu zesty-security/main amd64 Packages [220 kB]
Get:23 http://archive.ubuntu.com/ubuntu zesty-backports/main amd64 Packages [1698 B]
Get:24 http://archive.ubuntu.com/ubuntu zesty-backports/universe amd64 Packages [4043 B]
Get:25 http://ports.ubuntu.com zesty-updates/universe armhf Packages [222 kB]
Get:26 http://ports.ubuntu.com zesty-updates/main armhf Packages [303 kB]
Get:27 http://ports.ubuntu.com zesty-updates/universe arm64 Packages [235 kB]
Get:28 http://ports.ubuntu.com zesty-updates/main arm64 Packages [300 kB]
Get:29 http://ports.ubuntu.com zesty-updates/restricted armhf Packages [2953 B]
Get:30 http://ports.ubuntu.com zesty-security/universe armhf Packages [121 kB]
Get:31 http://ports.ubuntu.com zesty-security/main arm64 Packages [203 kB]
Get:32 http://ports.ubuntu.com zesty-security/universe arm64 Packages [133 kB]
Get:33 http://ports.ubuntu.com zesty-security/restricted armhf Packages [2953 B]
Get:34 http://ports.ubuntu.com zesty-security/main armhf Packages [206 kB]
Get:35 http://ports.ubuntu.com zesty-backports/main armhf Packages [1702 B]
Get:36 http://ports.ubuntu.com zesty-backports/main arm64 Packages [1704 B]
Get:37 http://ports.ubuntu.com zesty-backports/universe arm64 Packages [3786 B]
Get:38 http://ports.ubuntu.com zesty-backports/universe armhf Packages [4036 B]
Fetched 39.3 MB in 1min 57s (335 kB/s)
Reading package lists...
Reading package lists...
Building dependency tree...
Reading state information...
The following additional packages will be installed:
  dh-python icu-devtools libaio1 libbabeltrace-ctf-dev libbabeltrace-ctf1
  libbabeltrace1 libboost-atomic1.62-dev libboost-atomic1.62.0
  libboost-chrono1.62-dev libboost-chrono1.62.0 libboost-context1.62-dev
  libboost-context1.62.0 libboost-coroutine1.62-dev libboost-coroutine1.62.0
  libboost-date-time1.62-dev libboost-date-time1.62.0
  libboost-filesystem1.62-dev libboost-filesystem1.62.0
  libboost-iostreams1.62-dev libboost-iostreams1.62.0
  libboost-program-options1.62-dev libboost-program-options1.62.0
  libboost-python1.62-dev libboost-python1.62.0 libboost-random1.62-dev
  libboost-random1.62.0 libboost-regex1.62-dev libboost-regex1.62.0
  libboost-serialization1.62-dev libboost-serialization1.62.0
  libboost-system1.62-dev libboost-system1.62.0 libboost-thread1.62-dev
  libboost-thread1.62.0 libboost1.62-dev libdw1 libelf1 libglib2.0-0
  libibverbs1 libicu-dev libjemalloc1 libmpdec2 libnl-3-200 libnl-route-3-200
  libnspr4 libnspr4-dev libnss3 libpython-stdlib libpython2.7 libpython2.7-dev
  libpython2.7-minimal libpython2.7-stdlib libpython3-dev libpython3-stdlib
  libpython3.5 libpython3.5-dev libpython3.5-minimal libpython3.5-stdlib
  libsnappy1v5 libunwind8 libunwind8-dev mime-support python-dev
  python-minimal python-pkg-resources python2.7 python2.7-dev
  python2.7-minimal python3 python3-dev python3-minimal python3.5
  python3.5-dev python3.5-minimal uuid-dev
Suggested packages:
  cython-doc libboost1.62-doc gccxml libboost-exception1.62-dev
  libboost-fiber1.62-dev libboost-graph1.62-dev
  libboost-graph-parallel1.62-dev libboost-locale1.62-dev libboost-log1.62-dev
  libboost-math1.62-dev libboost-mpi1.62-dev libboost-mpi-python1.62-dev
  libboost-signals1.62-dev libboost-test1.62-dev libboost-timer1.62-dev
  libboost-type-erasure1.62-dev libboost-wave1.62-dev libboost1.62-tools-dev
  libmpfrc++-dev libntl-dev libcurl4-doc libcurl3-dbg libgnutls28-dev
  libidn11-dev libkrb5-dev librtmp-dev pkg-config icu-doc python-doc python-tk
  python-setuptools-doc python2.7-doc binfmt-support python3-doc python3-tk
  python3-venv python3.5-venv python3.5-doc
Recommended packages:
  libglib2.0-data shared-mime-info xdg-user-dirs libssl-doc
The following NEW packages will be installed:
  cython dh-python gperf icu-devtools libaio-dev libaio1 libatomic-ops-dev
  libbabeltrace-ctf-dev libbabeltrace-ctf1 libbabeltrace-dev libbabeltrace1
  libblkid-dev libboost-atomic1.62-dev libboost-atomic1.62.0
  libboost-chrono1.62-dev libboost-chrono1.62.0 libboost-context-dev
  libboost-context1.62-dev libboost-context1.62.0 libboost-coroutine-dev
  libboost-coroutine1.62-dev libboost-coroutine1.62.0 libboost-date-time-dev
  libboost-date-time1.62-dev libboost-date-time1.62.0 libboost-filesystem-dev
  libboost-filesystem1.62-dev libboost-filesystem1.62.0 libboost-iostreams-dev
  libboost-iostreams1.62-dev libboost-iostreams1.62.0
  libboost-program-options-dev libboost-program-options1.62-dev
  libboost-program-options1.62.0 libboost-python-dev libboost-python1.62-dev
  libboost-python1.62.0 libboost-random-dev libboost-random1.62-dev
  libboost-random1.62.0 libboost-regex-dev libboost-regex1.62-dev
  libboost-regex1.62.0 libboost-serialization1.62-dev
  libboost-serialization1.62.0 libboost-system-dev libboost-system1.62-dev
  libboost-system1.62.0 libboost-thread-dev libboost-thread1.62-dev
  libboost-thread1.62.0 libboost1.62-dev libcurl4-gnutls-dev libdw1 libelf1
  libexpat1-dev libglib2.0-0 libgoogle-perftools-dev libgoogle-perftools4
  libibverbs-dev libibverbs1 libicu-dev libjemalloc-dev libjemalloc1
  libkeyutils-dev libldap2-dev libmpdec2 libnl-3-200 libnl-route-3-200
  libnspr4 libnspr4-dev libnss3 libnss3-dev libpython-dev libpython-stdlib
  libpython2.7 libpython2.7-dev libpython2.7-minimal libpython2.7-stdlib
  libpython3-dev libpython3-stdlib libpython3.5 libpython3.5-dev
  libpython3.5-minimal libpython3.5-stdlib libsnappy-dev libsnappy1v5
  libssl-dev libtcmalloc-minimal4 libudev-dev libunwind-dev libunwind8
  libunwind8-dev mime-support python python-dev python-minimal
  python-pkg-resources python-setuptools python2.7 python2.7-dev
  python2.7-minimal python3 python3-dev python3-minimal python3.5
  python3.5-dev python3.5-minimal uuid-dev zlib1g-dev
0 upgraded, 110 newly installed, 0 to remove and 0 not upgraded.
Need to get 113 MB of archives.
After this operation, 418 MB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython3.5-minimal amd64 3.5.3-1ubuntu0~17.04.2 [527 kB]
Get:2 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 python3.5-minimal amd64 3.5.3-1ubuntu0~17.04.2 [1636 kB]
Get:3 http://archive.ubuntu.com/ubuntu zesty/main amd64 python3-minimal amd64 3.5.3-1 [23.4 kB]
Get:4 http://archive.ubuntu.com/ubuntu zesty/main amd64 mime-support all 3.60ubuntu1 [30.1 kB]
Get:5 http://archive.ubuntu.com/ubuntu zesty/main amd64 libmpdec2 amd64 2.4.2-1 [82.6 kB]
Get:6 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython3.5-stdlib amd64 3.5.3-1ubuntu0~17.04.2 [2163 kB]
Get:7 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 python3.5 amd64 3.5.3-1ubuntu0~17.04.2 [175 kB]
Get:8 http://archive.ubuntu.com/ubuntu zesty/main amd64 libpython3-stdlib amd64 3.5.3-1 [6838 B]
Get:9 http://archive.ubuntu.com/ubuntu zesty/main amd64 dh-python all 2.20170125 [83.7 kB]
Get:10 http://archive.ubuntu.com/ubuntu zesty/main amd64 python3 amd64 3.5.3-1 [8696 B]
Get:11 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libudev-dev amd64 232-21ubuntu7.1 [19.0 kB]
Get:12 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython2.7-minimal amd64 2.7.13-2ubuntu0.1 [338 kB]
Get:13 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 python2.7-minimal amd64 2.7.13-2ubuntu0.1 [1323 kB]
Get:14 http://archive.ubuntu.com/ubuntu zesty/main amd64 python-minimal amd64 2.7.13-2 [28.2 kB]
Get:15 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython2.7-stdlib amd64 2.7.13-2ubuntu0.1 [1899 kB]
Get:16 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 python2.7 amd64 2.7.13-2ubuntu0.1 [229 kB]
Get:17 http://archive.ubuntu.com/ubuntu zesty/main amd64 libpython-stdlib amd64 2.7.13-2 [7774 B]
Get:18 http://archive.ubuntu.com/ubuntu zesty/main amd64 python amd64 2.7.13-2 [139 kB]
Get:19 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 uuid-dev amd64 2.29-1ubuntu2.1 [26.7 kB]
Get:20 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libblkid-dev amd64 2.29-1ubuntu2.1 [138 kB]
Get:21 http://archive.ubuntu.com/ubuntu zesty/main amd64 libelf1 amd64 0.166-2ubuntu1 [42.9 kB]
Get:22 http://archive.ubuntu.com/ubuntu zesty/main amd64 libglib2.0-0 amd64 2.52.0-1 [1144 kB]
Get:23 http://archive.ubuntu.com/ubuntu zesty/universe amd64 cython amd64 0.25.2-1 [1835 kB]
Get:24 http://archive.ubuntu.com/ubuntu zesty/universe amd64 gperf amd64 3.0.4-2 [102 kB]
Get:25 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 icu-devtools amd64 57.1-5ubuntu0.2 [168 kB]
Get:26 http://archive.ubuntu.com/ubuntu zesty/main amd64 libaio1 amd64 0.3.110-3 [6382 B]
Get:27 http://archive.ubuntu.com/ubuntu zesty/main amd64 libaio-dev amd64 0.3.110-3 [12.8 kB]
Get:28 http://archive.ubuntu.com/ubuntu zesty/main amd64 libatomic-ops-dev amd64 7.4.4-3 [70.8 kB]
Get:29 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost1.62-dev amd64 1.62.0+dfsg-4 [6986 kB]
Get:30 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-atomic1.62.0 amd64 1.62.0+dfsg-4 [7336 B]
Get:31 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-atomic1.62-dev amd64 1.62.0+dfsg-4 [5060 B]
Get:32 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-system1.62.0 amd64 1.62.0+dfsg-4 [9398 B]
Get:33 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-chrono1.62.0 amd64 1.62.0+dfsg-4 [11.8 kB]
Get:34 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-chrono1.62-dev amd64 1.62.0+dfsg-4 [12.1 kB]
Get:35 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-context1.62.0 amd64 1.62.0+dfsg-4 [8298 B]
Get:36 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-context1.62-dev amd64 1.62.0+dfsg-4 [6364 B]
Get:37 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-context-dev amd64 1.62.0.1 [3062 B]
Get:38 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-thread1.62.0 amd64 1.62.0+dfsg-4 [46.9 kB]
Get:39 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-coroutine1.62.0 amd64 1.62.0+dfsg-4 [19.7 kB]
Get:40 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-coroutine1.62-dev amd64 1.62.0+dfsg-4 [20.1 kB]
Get:41 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-coroutine-dev amd64 1.62.0.1 [3126 B]
Get:42 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-date-time1.62.0 amd64 1.62.0+dfsg-4 [20.0 kB]
Get:43 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-serialization1.62.0 amd64 1.62.0+dfsg-4 [101 kB]
Get:44 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-serialization1.62-dev amd64 1.62.0+dfsg-4 [139 kB]
Get:45 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-date-time1.62-dev amd64 1.62.0+dfsg-4 [25.8 kB]
Get:46 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-date-time-dev amd64 1.62.0.1 [2876 B]
Get:47 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-filesystem1.62.0 amd64 1.62.0+dfsg-4 [38.6 kB]
Get:48 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-system1.62-dev amd64 1.62.0+dfsg-4 [9540 B]
Get:49 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-filesystem1.62-dev amd64 1.62.0+dfsg-4 [48.6 kB]
Get:50 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-filesystem-dev amd64 1.62.0.1 [2900 B]
Get:51 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-regex1.62.0 amd64 1.62.0+dfsg-4 [263 kB]
Get:52 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libicu-dev amd64 57.1-5ubuntu0.2 [16.5 MB]
Get:53 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-regex1.62-dev amd64 1.62.0+dfsg-4 [309 kB]
Get:54 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-iostreams1.62.0 amd64 1.62.0+dfsg-4 [27.5 kB]
Get:55 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-iostreams1.62-dev amd64 1.62.0+dfsg-4 [32.7 kB]
Get:56 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-iostreams-dev amd64 1.62.0.1 [2860 B]
Get:57 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-program-options1.62.0 amd64 1.62.0+dfsg-4 [136 kB]
Get:58 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-program-options1.62-dev amd64 1.62.0+dfsg-4 [159 kB]
Get:59 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-program-options-dev amd64 1.62.0.1 [2886 B]
Get:60 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-python1.62.0 amd64 1.62.0+dfsg-4 [113 kB]
Get:61 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython2.7 amd64 2.7.13-2ubuntu0.1 [1072 kB]
Get:62 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libexpat1-dev amd64 2.2.0-2ubuntu0.1 [122 kB]
Get:63 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython2.7-dev amd64 2.7.13-2ubuntu0.1 [28.2 MB]
Get:64 http://archive.ubuntu.com/ubuntu zesty/main amd64 libpython-dev amd64 2.7.13-2 [7842 B]
Get:65 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 python2.7-dev amd64 2.7.13-2ubuntu0.1 [276 kB]
Get:66 http://archive.ubuntu.com/ubuntu zesty/main amd64 python-dev amd64 2.7.13-2 [1160 B]
Get:67 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython3.5 amd64 3.5.3-1ubuntu0~17.04.2 [1372 kB]
Get:68 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython3.5-dev amd64 3.5.3-1ubuntu0~17.04.2 [37.7 MB]
Get:69 http://archive.ubuntu.com/ubuntu zesty/main amd64 libpython3-dev amd64 3.5.3-1 [6932 B]
Get:70 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 python3.5-dev amd64 3.5.3-1ubuntu0~17.04.2 [413 kB]
Get:71 http://archive.ubuntu.com/ubuntu zesty/main amd64 python3-dev amd64 3.5.3-1 [1190 B]
Get:72 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-python1.62-dev amd64 1.62.0+dfsg-4 [142 kB]
Get:73 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-python-dev amd64 1.62.0.1 [3254 B]
Get:74 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-random1.62.0 amd64 1.62.0+dfsg-4 [11.1 kB]
Get:75 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-random1.62-dev amd64 1.62.0+dfsg-4 [10.5 kB]
Get:76 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-random-dev amd64 1.62.0.1 [2866 B]
Get:77 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-regex-dev amd64 1.62.0.1 [3120 B]
Get:78 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-system-dev amd64 1.62.0.1 [3006 B]
Get:79 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-thread1.62-dev amd64 1.62.0+dfsg-4 [46.7 kB]
Get:80 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-thread-dev amd64 1.62.0.1 [2882 B]
Get:81 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libcurl4-gnutls-dev amd64 7.52.1-4ubuntu1.4 [270 kB]
Get:82 http://archive.ubuntu.com/ubuntu zesty/main amd64 libdw1 amd64 0.166-2ubuntu1 [191 kB]
Get:83 http://archive.ubuntu.com/ubuntu zesty/main amd64 libtcmalloc-minimal4 amd64 2.5-0ubuntu4 [88.3 kB]
Get:84 http://archive.ubuntu.com/ubuntu zesty/main amd64 libunwind8 amd64 1.1-4.1ubuntu2 [46.3 kB]
Get:85 http://archive.ubuntu.com/ubuntu zesty/main amd64 libgoogle-perftools4 amd64 2.5-0ubuntu4 [183 kB]
Get:86 http://archive.ubuntu.com/ubuntu zesty/main amd64 libunwind-dev amd64 1.1-4.1ubuntu2 [417 kB]
Get:87 http://archive.ubuntu.com/ubuntu zesty/main amd64 libunwind8-dev amd64 1.1-4.1ubuntu2 [2610 B]
Get:88 http://archive.ubuntu.com/ubuntu zesty/main amd64 libgoogle-perftools-dev amd64 2.5-0ubuntu4 [193 kB]
Get:89 http://archive.ubuntu.com/ubuntu zesty/main amd64 libkeyutils-dev amd64 1.5.9-9ubuntu1 [32.6 kB]
Get:90 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnl-3-200 amd64 3.2.29-0ubuntu2.1 [52.9 kB]
Get:91 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnl-route-3-200 amd64 3.2.29-0ubuntu2.1 [146 kB]
Get:92 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnspr4 amd64 2:4.13.1-0ubuntu0.17.04.1 [112 kB]
Get:93 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnspr4-dev amd64 2:4.13.1-0ubuntu0.17.04.1 [212 kB]
Get:94 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnss3 amd64 2:3.28.4-0ubuntu0.17.04.3 [1143 kB]
Get:95 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnss3-dev amd64 2:3.28.4-0ubuntu0.17.04.3 [220 kB]
Get:96 http://archive.ubuntu.com/ubuntu zesty/main amd64 zlib1g-dev amd64 1:1.2.11.dfsg-0ubuntu1 [173 kB]
Get:97 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libssl-dev amd64 1.0.2g-1ubuntu11.4 [1352 kB]
Get:98 http://archive.ubuntu.com/ubuntu zesty/main amd64 python-pkg-resources all 33.1.1-1 [127 kB]
Get:99 http://archive.ubuntu.com/ubuntu zesty/main amd64 python-setuptools all 33.1.1-1 [258 kB]
Get:100 http://archive.ubuntu.com/ubuntu zesty/main amd64 libbabeltrace1 amd64 1.5.1-1 [34.6 kB]
Get:101 http://archive.ubuntu.com/ubuntu zesty/main amd64 libbabeltrace-ctf1 amd64 1.5.1-1 [124 kB]
Get:102 http://archive.ubuntu.com/ubuntu zesty/main amd64 libbabeltrace-ctf-dev amd64 1.5.1-1 [152 kB]
Get:103 http://archive.ubuntu.com/ubuntu zesty/main amd64 libbabeltrace-dev amd64 1.5.1-1 [40.3 kB]
Get:104 http://archive.ubuntu.com/ubuntu zesty/main amd64 libibverbs1 amd64 1.2.1-2ubuntu1 [30.8 kB]
Get:105 http://archive.ubuntu.com/ubuntu zesty/main amd64 libibverbs-dev amd64 1.2.1-2ubuntu1 [97.5 kB]
Get:106 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libjemalloc1 amd64 3.6.0-9.1 [79.9 kB]
Get:107 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libjemalloc-dev amd64 3.6.0-9.1 [156 kB]
Get:108 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libldap2-dev amd64 2.4.44+dfsg-3ubuntu2.1 [261 kB]
Get:109 http://archive.ubuntu.com/ubuntu zesty/main amd64 libsnappy1v5 amd64 1.1.4-1 [16.0 kB]
Get:110 http://archive.ubuntu.com/ubuntu zesty/main amd64 libsnappy-dev amd64 1.1.4-1 [27.9 kB]
debconf: delaying package configuration, since apt-utils is not installed
Fetched 113 MB in 5min 48s (325 kB/s)
Selecting previously unselected package libpython3.5-minimal:amd64.
(Reading database ... 18014 files and directories currently installed.)
Preparing to unpack .../0-libpython3.5-minimal_3.5.3-1ubuntu0~17.04.2_amd64.deb ...
Unpacking libpython3.5-minimal:amd64 (3.5.3-1ubuntu0~17.04.2) ...
Selecting previously unselected package python3.5-minimal.
Preparing to unpack .../1-python3.5-minimal_3.5.3-1ubuntu0~17.04.2_amd64.deb ...
Unpacking python3.5-minimal (3.5.3-1ubuntu0~17.04.2) ...
Selecting previously unselected package python3-minimal.
Preparing to unpack .../2-python3-minimal_3.5.3-1_amd64.deb ...
Unpacking python3-minimal (3.5.3-1) ...
Selecting previously unselected package mime-support.
Preparing to unpack .../3-mime-support_3.60ubuntu1_all.deb ...
Unpacking mime-support (3.60ubuntu1) ...
Selecting previously unselected package libmpdec2:amd64.
Preparing to unpack .../4-libmpdec2_2.4.2-1_amd64.deb ...
Unpacking libmpdec2:amd64 (2.4.2-1) ...
Selecting previously unselected package libpython3.5-stdlib:amd64.
Preparing to unpack .../5-libpython3.5-stdlib_3.5.3-1ubuntu0~17.04.2_amd64.deb ...
Unpacking libpython3.5-stdlib:amd64 (3.5.3-1ubuntu0~17.04.2) ...
Selecting previously unselected package python3.5.
Preparing to unpack .../6-python3.5_3.5.3-1ubuntu0~17.04.2_amd64.deb ...
Unpacking python3.5 (3.5.3-1ubuntu0~17.04.2) ...
Selecting previously unselected package libpython3-stdlib:amd64.
Preparing to unpack .../7-libpython3-stdlib_3.5.3-1_amd64.deb ...
Unpacking libpython3-stdlib:amd64 (3.5.3-1) ...
Selecting previously unselected package dh-python.
Preparing to unpack .../8-dh-python_2.20170125_all.deb ...
Unpacking dh-python (2.20170125) ...
Setting up libpython3.5-minimal:amd64 (3.5.3-1ubuntu0~17.04.2) ...
Setting up python3.5-minimal (3.5.3-1ubuntu0~17.04.2) ...
Setting up python3-minimal (3.5.3-1) ...
Selecting previously unselected package python3.
(Reading database ... 18958 files and directories currently installed.)
Preparing to unpack .../0-python3_3.5.3-1_amd64.deb ...
Unpacking python3 (3.5.3-1) ...
Selecting previously unselected package libudev-dev:amd64.
Preparing to unpack .../1-libudev-dev_232-21ubuntu7.1_amd64.deb ...
Unpacking libudev-dev:amd64 (232-21ubuntu7.1) ...
Selecting previously unselected package libpython2.7-minimal:amd64.
Preparing to unpack .../2-libpython2.7-minimal_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking libpython2.7-minimal:amd64 (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package python2.7-minimal.
Preparing to unpack .../3-python2.7-minimal_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking python2.7-minimal (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package python-minimal.
Preparing to unpack .../4-python-minimal_2.7.13-2_amd64.deb ...
Unpacking python-minimal (2.7.13-2) ...
Selecting previously unselected package libpython2.7-stdlib:amd64.
Preparing to unpack .../5-libpython2.7-stdlib_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking libpython2.7-stdlib:amd64 (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package python2.7.
Preparing to unpack .../6-python2.7_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking python2.7 (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package libpython-stdlib:amd64.
Preparing to unpack .../7-libpython-stdlib_2.7.13-2_amd64.deb ...
Unpacking libpython-stdlib:amd64 (2.7.13-2) ...
Setting up libpython2.7-minimal:amd64 (2.7.13-2ubuntu0.1) ...
Setting up python2.7-minimal (2.7.13-2ubuntu0.1) ...
Setting up python-minimal (2.7.13-2) ...
Selecting previously unselected package python.
(Reading database ... 19795 files and directories currently installed.)
Preparing to unpack .../00-python_2.7.13-2_amd64.deb ...
Unpacking python (2.7.13-2) ...
Selecting previously unselected package uuid-dev:amd64.
Preparing to unpack .../01-uuid-dev_2.29-1ubuntu2.1_amd64.deb ...
Unpacking uuid-dev:amd64 (2.29-1ubuntu2.1) ...
Selecting previously unselected package libblkid-dev:amd64.
Preparing to unpack .../02-libblkid-dev_2.29-1ubuntu2.1_amd64.deb ...
Unpacking libblkid-dev:amd64 (2.29-1ubuntu2.1) ...
Selecting previously unselected package libelf1:amd64.
Preparing to unpack .../03-libelf1_0.166-2ubuntu1_amd64.deb ...
Unpacking libelf1:amd64 (0.166-2ubuntu1) ...
Selecting previously unselected package libglib2.0-0:amd64.
Preparing to unpack .../04-libglib2.0-0_2.52.0-1_amd64.deb ...
Unpacking libglib2.0-0:amd64 (2.52.0-1) ...
Selecting previously unselected package cython.
Preparing to unpack .../05-cython_0.25.2-1_amd64.deb ...
Unpacking cython (0.25.2-1) ...
Selecting previously unselected package gperf.
Preparing to unpack .../06-gperf_3.0.4-2_amd64.deb ...
Unpacking gperf (3.0.4-2) ...
Selecting previously unselected package icu-devtools.
Preparing to unpack .../07-icu-devtools_57.1-5ubuntu0.2_amd64.deb ...
Unpacking icu-devtools (57.1-5ubuntu0.2) ...
Selecting previously unselected package libaio1:amd64.
Preparing to unpack .../08-libaio1_0.3.110-3_amd64.deb ...
Unpacking libaio1:amd64 (0.3.110-3) ...
Selecting previously unselected package libaio-dev.
Preparing to unpack .../09-libaio-dev_0.3.110-3_amd64.deb ...
Unpacking libaio-dev (0.3.110-3) ...
Selecting previously unselected package libatomic-ops-dev.
Preparing to unpack .../10-libatomic-ops-dev_7.4.4-3_amd64.deb ...
Unpacking libatomic-ops-dev (7.4.4-3) ...
Selecting previously unselected package libboost1.62-dev:amd64.
Preparing to unpack .../11-libboost1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-atomic1.62.0:amd64.
Preparing to unpack .../12-libboost-atomic1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-atomic1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-atomic1.62-dev:amd64.
Preparing to unpack .../13-libboost-atomic1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-atomic1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-system1.62.0:amd64.
Preparing to unpack .../14-libboost-system1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-system1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-chrono1.62.0:amd64.
Preparing to unpack .../15-libboost-chrono1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-chrono1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-chrono1.62-dev:amd64.
Preparing to unpack .../16-libboost-chrono1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-chrono1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-context1.62.0:amd64.
Preparing to unpack .../17-libboost-context1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-context1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-context1.62-dev:amd64.
Preparing to unpack .../18-libboost-context1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-context1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-context-dev:amd64.
Preparing to unpack .../19-libboost-context-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-context-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libboost-thread1.62.0:amd64.
Preparing to unpack .../20-libboost-thread1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-thread1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-coroutine1.62.0:amd64.
Preparing to unpack .../21-libboost-coroutine1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-coroutine1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-coroutine1.62-dev:amd64.
Preparing to unpack .../22-libboost-coroutine1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-coroutine1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-coroutine-dev:amd64.
Preparing to unpack .../23-libboost-coroutine-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-coroutine-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libboost-date-time1.62.0:amd64.
Preparing to unpack .../24-libboost-date-time1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-date-time1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-serialization1.62.0:amd64.
Preparing to unpack .../25-libboost-serialization1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-serialization1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-serialization1.62-dev:amd64.
Preparing to unpack .../26-libboost-serialization1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-serialization1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-date-time1.62-dev:amd64.
Preparing to unpack .../27-libboost-date-time1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-date-time1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-date-time-dev:amd64.
Preparing to unpack .../28-libboost-date-time-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-date-time-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libboost-filesystem1.62.0:amd64.
Preparing to unpack .../29-libboost-filesystem1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-filesystem1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-system1.62-dev:amd64.
Preparing to unpack .../30-libboost-system1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-system1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-filesystem1.62-dev:amd64.
Preparing to unpack .../31-libboost-filesystem1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-filesystem1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-filesystem-dev:amd64.
Preparing to unpack .../32-libboost-filesystem-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-filesystem-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libboost-regex1.62.0:amd64.
Preparing to unpack .../33-libboost-regex1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-regex1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libicu-dev.
Preparing to unpack .../34-libicu-dev_57.1-5ubuntu0.2_amd64.deb ...
Unpacking libicu-dev (57.1-5ubuntu0.2) ...
Selecting previously unselected package libboost-regex1.62-dev:amd64.
Preparing to unpack .../35-libboost-regex1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-regex1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-iostreams1.62.0:amd64.
Preparing to unpack .../36-libboost-iostreams1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-iostreams1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-iostreams1.62-dev:amd64.
Preparing to unpack .../37-libboost-iostreams1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-iostreams1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-iostreams-dev:amd64.
Preparing to unpack .../38-libboost-iostreams-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-iostreams-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libboost-program-options1.62.0:amd64.
Preparing to unpack .../39-libboost-program-options1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-program-options1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-program-options1.62-dev:amd64.
Preparing to unpack .../40-libboost-program-options1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-program-options1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-program-options-dev:amd64.
Preparing to unpack .../41-libboost-program-options-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-program-options-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libboost-python1.62.0.
Preparing to unpack .../42-libboost-python1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-python1.62.0 (1.62.0+dfsg-4) ...
Selecting previously unselected package libpython2.7:amd64.
Preparing to unpack .../43-libpython2.7_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking libpython2.7:amd64 (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package libexpat1-dev:amd64.
Preparing to unpack .../44-libexpat1-dev_2.2.0-2ubuntu0.1_amd64.deb ...
Unpacking libexpat1-dev:amd64 (2.2.0-2ubuntu0.1) ...
Selecting previously unselected package libpython2.7-dev:amd64.
Preparing to unpack .../45-libpython2.7-dev_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking libpython2.7-dev:amd64 (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package libpython-dev:amd64.
Preparing to unpack .../46-libpython-dev_2.7.13-2_amd64.deb ...
Unpacking libpython-dev:amd64 (2.7.13-2) ...
Selecting previously unselected package python2.7-dev.
Preparing to unpack .../47-python2.7-dev_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking python2.7-dev (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package python-dev.
Preparing to unpack .../48-python-dev_2.7.13-2_amd64.deb ...
Unpacking python-dev (2.7.13-2) ...
Selecting previously unselected package libpython3.5:amd64.
Preparing to unpack .../49-libpython3.5_3.5.3-1ubuntu0~17.04.2_amd64.deb ...
Unpacking libpython3.5:amd64 (3.5.3-1ubuntu0~17.04.2) ...
Selecting previously unselected package libpython3.5-dev:amd64.
Preparing to unpack .../50-libpython3.5-dev_3.5.3-1ubuntu0~17.04.2_amd64.deb ...
Unpacking libpython3.5-dev:amd64 (3.5.3-1ubuntu0~17.04.2) ...
Selecting previously unselected package libpython3-dev:amd64.
Preparing to unpack .../51-libpython3-dev_3.5.3-1_amd64.deb ...
Unpacking libpython3-dev:amd64 (3.5.3-1) ...
Selecting previously unselected package python3.5-dev.
Preparing to unpack .../52-python3.5-dev_3.5.3-1ubuntu0~17.04.2_amd64.deb ...
Unpacking python3.5-dev (3.5.3-1ubuntu0~17.04.2) ...
Selecting previously unselected package python3-dev.
Preparing to unpack .../53-python3-dev_3.5.3-1_amd64.deb ...
Unpacking python3-dev (3.5.3-1) ...
Selecting previously unselected package libboost-python1.62-dev.
Preparing to unpack .../54-libboost-python1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-python1.62-dev (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-python-dev.
Preparing to unpack .../55-libboost-python-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-python-dev (1.62.0.1) ...
Selecting previously unselected package libboost-random1.62.0:amd64.
Preparing to unpack .../56-libboost-random1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-random1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-random1.62-dev:amd64.
Preparing to unpack .../57-libboost-random1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-random1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-random-dev:amd64.
Preparing to unpack .../58-libboost-random-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-random-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libboost-regex-dev:amd64.
Preparing to unpack .../59-libboost-regex-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-regex-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libboost-system-dev:amd64.
Preparing to unpack .../60-libboost-system-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-system-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libboost-thread1.62-dev:amd64.
Preparing to unpack .../61-libboost-thread1.62-dev_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-thread1.62-dev:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-thread-dev:amd64.
Preparing to unpack .../62-libboost-thread-dev_1.62.0.1_amd64.deb ...
Unpacking libboost-thread-dev:amd64 (1.62.0.1) ...
Selecting previously unselected package libcurl4-gnutls-dev:amd64.
Preparing to unpack .../63-libcurl4-gnutls-dev_7.52.1-4ubuntu1.4_amd64.deb ...
Unpacking libcurl4-gnutls-dev:amd64 (7.52.1-4ubuntu1.4) ...
Selecting previously unselected package libdw1:amd64.
Preparing to unpack .../64-libdw1_0.166-2ubuntu1_amd64.deb ...
Unpacking libdw1:amd64 (0.166-2ubuntu1) ...
Selecting previously unselected package libtcmalloc-minimal4.
Preparing to unpack .../65-libtcmalloc-minimal4_2.5-0ubuntu4_amd64.deb ...
Unpacking libtcmalloc-minimal4 (2.5-0ubuntu4) ...
Selecting previously unselected package libunwind8.
Preparing to unpack .../66-libunwind8_1.1-4.1ubuntu2_amd64.deb ...
Unpacking libunwind8 (1.1-4.1ubuntu2) ...
Selecting previously unselected package libgoogle-perftools4.
Preparing to unpack .../67-libgoogle-perftools4_2.5-0ubuntu4_amd64.deb ...
Unpacking libgoogle-perftools4 (2.5-0ubuntu4) ...
Selecting previously unselected package libunwind-dev.
Preparing to unpack .../68-libunwind-dev_1.1-4.1ubuntu2_amd64.deb ...
Unpacking libunwind-dev (1.1-4.1ubuntu2) ...
Selecting previously unselected package libunwind8-dev.
Preparing to unpack .../69-libunwind8-dev_1.1-4.1ubuntu2_amd64.deb ...
Unpacking libunwind8-dev (1.1-4.1ubuntu2) ...
Selecting previously unselected package libgoogle-perftools-dev.
Preparing to unpack .../70-libgoogle-perftools-dev_2.5-0ubuntu4_amd64.deb ...
Unpacking libgoogle-perftools-dev (2.5-0ubuntu4) ...
Selecting previously unselected package libkeyutils-dev:amd64.
Preparing to unpack .../71-libkeyutils-dev_1.5.9-9ubuntu1_amd64.deb ...
Unpacking libkeyutils-dev:amd64 (1.5.9-9ubuntu1) ...
Selecting previously unselected package libnl-3-200:amd64.
Preparing to unpack .../72-libnl-3-200_3.2.29-0ubuntu2.1_amd64.deb ...
Unpacking libnl-3-200:amd64 (3.2.29-0ubuntu2.1) ...
Selecting previously unselected package libnl-route-3-200:amd64.
Preparing to unpack .../73-libnl-route-3-200_3.2.29-0ubuntu2.1_amd64.deb ...
Unpacking libnl-route-3-200:amd64 (3.2.29-0ubuntu2.1) ...
Selecting previously unselected package libnspr4:amd64.
Preparing to unpack .../74-libnspr4_2%3a4.13.1-0ubuntu0.17.04.1_amd64.deb ...
Unpacking libnspr4:amd64 (2:4.13.1-0ubuntu0.17.04.1) ...
Selecting previously unselected package libnspr4-dev.
Preparing to unpack .../75-libnspr4-dev_2%3a4.13.1-0ubuntu0.17.04.1_amd64.deb ...
Unpacking libnspr4-dev (2:4.13.1-0ubuntu0.17.04.1) ...
Selecting previously unselected package libnss3:amd64.
Preparing to unpack .../76-libnss3_2%3a3.28.4-0ubuntu0.17.04.3_amd64.deb ...
Unpacking libnss3:amd64 (2:3.28.4-0ubuntu0.17.04.3) ...
Selecting previously unselected package libnss3-dev:amd64.
Preparing to unpack .../77-libnss3-dev_2%3a3.28.4-0ubuntu0.17.04.3_amd64.deb ...
Unpacking libnss3-dev:amd64 (2:3.28.4-0ubuntu0.17.04.3) ...
Selecting previously unselected package zlib1g-dev:amd64.
Preparing to unpack .../78-zlib1g-dev_1%3a1.2.11.dfsg-0ubuntu1_amd64.deb ...
Unpacking zlib1g-dev:amd64 (1:1.2.11.dfsg-0ubuntu1) ...
Selecting previously unselected package libssl-dev:amd64.
Preparing to unpack .../79-libssl-dev_1.0.2g-1ubuntu11.4_amd64.deb ...
Unpacking libssl-dev:amd64 (1.0.2g-1ubuntu11.4) ...
Selecting previously unselected package python-pkg-resources.
Preparing to unpack .../80-python-pkg-resources_33.1.1-1_all.deb ...
Unpacking python-pkg-resources (33.1.1-1) ...
Selecting previously unselected package python-setuptools.
Preparing to unpack .../81-python-setuptools_33.1.1-1_all.deb ...
Unpacking python-setuptools (33.1.1-1) ...
Selecting previously unselected package libbabeltrace1:amd64.
Preparing to unpack .../82-libbabeltrace1_1.5.1-1_amd64.deb ...
Unpacking libbabeltrace1:amd64 (1.5.1-1) ...
Selecting previously unselected package libbabeltrace-ctf1:amd64.
Preparing to unpack .../83-libbabeltrace-ctf1_1.5.1-1_amd64.deb ...
Unpacking libbabeltrace-ctf1:amd64 (1.5.1-1) ...
Selecting previously unselected package libbabeltrace-ctf-dev.
Preparing to unpack .../84-libbabeltrace-ctf-dev_1.5.1-1_amd64.deb ...
Unpacking libbabeltrace-ctf-dev (1.5.1-1) ...
Selecting previously unselected package libbabeltrace-dev.
Preparing to unpack .../85-libbabeltrace-dev_1.5.1-1_amd64.deb ...
Unpacking libbabeltrace-dev (1.5.1-1) ...
Selecting previously unselected package libibverbs1.
Preparing to unpack .../86-libibverbs1_1.2.1-2ubuntu1_amd64.deb ...
Unpacking libibverbs1 (1.2.1-2ubuntu1) ...
Selecting previously unselected package libibverbs-dev.
Preparing to unpack .../87-libibverbs-dev_1.2.1-2ubuntu1_amd64.deb ...
Unpacking libibverbs-dev (1.2.1-2ubuntu1) ...
Selecting previously unselected package libjemalloc1.
Preparing to unpack .../88-libjemalloc1_3.6.0-9.1_amd64.deb ...
Unpacking libjemalloc1 (3.6.0-9.1) ...
Selecting previously unselected package libjemalloc-dev.
Preparing to unpack .../89-libjemalloc-dev_3.6.0-9.1_amd64.deb ...
Unpacking libjemalloc-dev (3.6.0-9.1) ...
Selecting previously unselected package libldap2-dev:amd64.
Preparing to unpack .../90-libldap2-dev_2.4.44+dfsg-3ubuntu2.1_amd64.deb ...
Unpacking libldap2-dev:amd64 (2.4.44+dfsg-3ubuntu2.1) ...
Selecting previously unselected package libsnappy1v5:amd64.
Preparing to unpack .../91-libsnappy1v5_1.1.4-1_amd64.deb ...
Unpacking libsnappy1v5:amd64 (1.1.4-1) ...
Selecting previously unselected package libsnappy-dev:amd64.
Preparing to unpack .../92-libsnappy-dev_1.1.4-1_amd64.deb ...
Unpacking libsnappy-dev:amd64 (1.1.4-1) ...
Setting up libldap2-dev:amd64 (2.4.44+dfsg-3ubuntu2.1) ...
Setting up libboost-serialization1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-date-time1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up mime-support (3.60ubuntu1) ...
Setting up libboost1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libjemalloc1 (3.6.0-9.1) ...
Setting up libboost-program-options1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libelf1:amd64 (0.166-2ubuntu1) ...
Setting up libglib2.0-0:amd64 (2.52.0-1) ...
No schema files found: doing nothing.
Setting up libboost-regex1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-program-options1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up uuid-dev:amd64 (2.29-1ubuntu2.1) ...
Setting up gperf (3.0.4-2) ...
Setting up libjemalloc-dev (3.6.0-9.1) ...
Setting up libnspr4:amd64 (2:4.13.1-0ubuntu0.17.04.1) ...
Setting up libboost-serialization1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libnspr4-dev (2:4.13.1-0ubuntu0.17.04.1) ...
Setting up libboost-iostreams1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libtcmalloc-minimal4 (2.5-0ubuntu4) ...
Setting up libunwind8 (1.1-4.1ubuntu2) ...
Processing triggers for libc-bin (2.24-9ubuntu2.2) ...
Setting up libaio1:amd64 (0.3.110-3) ...
Setting up libkeyutils-dev:amd64 (1.5.9-9ubuntu1) ...
Setting up libsnappy1v5:amd64 (1.1.4-1) ...
Setting up libunwind-dev (1.1-4.1ubuntu2) ...
Setting up libboost-atomic1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libatomic-ops-dev (7.4.4-3) ...
Setting up libexpat1-dev:amd64 (2.2.0-2ubuntu0.1) ...
Setting up libboost-python1.62.0 (1.62.0+dfsg-4) ...
Setting up icu-devtools (57.1-5ubuntu0.2) ...
Setting up libaio-dev (0.3.110-3) ...
Setting up libnl-3-200:amd64 (3.2.29-0ubuntu2.1) ...
Setting up libcurl4-gnutls-dev:amd64 (7.52.1-4ubuntu1.4) ...
Setting up libudev-dev:amd64 (232-21ubuntu7.1) ...
Setting up libboost-context1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libpython2.7-stdlib:amd64 (2.7.13-2ubuntu0.1) ...
Setting up libbabeltrace1:amd64 (1.5.1-1) ...
Setting up libmpdec2:amd64 (2.4.2-1) ...
Setting up libboost-system1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-program-options-dev:amd64 (1.62.0.1) ...
Setting up zlib1g-dev:amd64 (1:1.2.11.dfsg-0ubuntu1) ...
Setting up libboost-atomic1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-date-time1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-context1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libunwind8-dev (1.1-4.1ubuntu2) ...
Setting up libdw1:amd64 (0.166-2ubuntu1) ...
Setting up libboost-thread1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-random1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libssl-dev:amd64 (1.0.2g-1ubuntu11.4) ...
Setting up libpython3.5-stdlib:amd64 (3.5.3-1ubuntu0~17.04.2) ...
Setting up libblkid-dev:amd64 (2.29-1ubuntu2.1) ...
Setting up python2.7 (2.7.13-2ubuntu0.1) ...
Setting up libnss3:amd64 (2:3.28.4-0ubuntu0.17.04.3) ...
Setting up libnl-route-3-200:amd64 (3.2.29-0ubuntu2.1) ...
Setting up libicu-dev (57.1-5ubuntu0.2) ...
Setting up libnss3-dev:amd64 (2:3.28.4-0ubuntu0.17.04.3) ...
Setting up libboost-date-time-dev:amd64 (1.62.0.1) ...
Setting up libpython-stdlib:amd64 (2.7.13-2) ...
Setting up libgoogle-perftools4 (2.5-0ubuntu4) ...
Setting up libsnappy-dev:amd64 (1.1.4-1) ...
Setting up libboost-system1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-filesystem1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libpython2.7:amd64 (2.7.13-2ubuntu0.1) ...
Setting up libboost-context-dev:amd64 (1.62.0.1) ...
Setting up libpython2.7-dev:amd64 (2.7.13-2ubuntu0.1) ...
Setting up python2.7-dev (2.7.13-2ubuntu0.1) ...
Setting up libbabeltrace-ctf1:amd64 (1.5.1-1) ...
Setting up libbabeltrace-ctf-dev (1.5.1-1) ...
Setting up python (2.7.13-2) ...
Setting up libboost-chrono1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-random1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libgoogle-perftools-dev (2.5-0ubuntu4) ...
Setting up python3.5 (3.5.3-1ubuntu0~17.04.2) ...
Setting up libpython3-stdlib:amd64 (3.5.3-1) ...
Setting up libibverbs1 (1.2.1-2ubuntu1) ...
Setting up libpython-dev:amd64 (2.7.13-2) ...
Setting up libboost-filesystem1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-coroutine1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-regex1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libpython3.5:amd64 (3.5.3-1ubuntu0~17.04.2) ...
Setting up python-dev (2.7.13-2) ...
Setting up libboost-system-dev:amd64 (1.62.0.1) ...
Setting up libboost-iostreams1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libbabeltrace-dev (1.5.1-1) ...
Setting up python-pkg-resources (33.1.1-1) ...
Setting up libboost-filesystem-dev:amd64 (1.62.0.1) ...
Setting up libboost-chrono1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libpython3.5-dev:amd64 (3.5.3-1ubuntu0~17.04.2) ...
Setting up cython (0.25.2-1) ...
Setting up libboost-random-dev:amd64 (1.62.0.1) ...
Setting up libboost-iostreams-dev:amd64 (1.62.0.1) ...
Setting up libboost-thread1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-coroutine1.62-dev:amd64 (1.62.0+dfsg-4) ...
Setting up libboost-regex-dev:amd64 (1.62.0.1) ...
Setting up libboost-coroutine-dev:amd64 (1.62.0.1) ...
Setting up python3.5-dev (3.5.3-1ubuntu0~17.04.2) ...
Setting up libibverbs-dev (1.2.1-2ubuntu1) ...
Setting up libpython3-dev:amd64 (3.5.3-1) ...
Setting up python-setuptools (33.1.1-1) ...
Setting up libboost-thread-dev:amd64 (1.62.0.1) ...
Setting up python3 (3.5.3-1) ...
Setting up python3-dev (3.5.3-1) ...
Setting up dh-python (2.20170125) ...
Setting up libboost-python1.62-dev (1.62.0+dfsg-4) ...
Setting up libboost-python-dev (1.62.0.1) ...
Processing triggers for libc-bin (2.24-9ubuntu2.2) ...
Reading package lists...
Building dependency tree...
Reading state information...
Calculating upgrade...
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
Reading package lists...
Building dependency tree...
Reading state information...
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
 ---> fee5046a3b49
Removing intermediate container 554996be9533
Step 7 : ARG CEPH_GIT_COMMIT
 ---> Running in 3c2025101ccb
 ---> e54d72560e34
Removing intermediate container 3c2025101ccb
Step 8 : RUN cd /build/ceph &&     git fetch --all --prune &&     git checkout -b ceph-builder ${CEPH_GIT_COMMIT} &&     git submodule update --init --recursive
 ---> Running in 19e6424305ec
Fetching origin
Switched to a new branch 'ceph-builder'
Submodule 'ceph-erasure-code-corpus' (https://github.com/ceph/ceph-erasure-code-corpus.git) registered for path 'ceph-erasure-code-corpus'
Submodule 'ceph-object-corpus' (https://github.com/ceph/ceph-object-corpus.git) registered for path 'ceph-object-corpus'
Submodule 'src/Beast' (https://github.com/ceph/Beast.git) registered for path 'src/Beast'
Submodule 'src/blkin' (https://github.com/ceph/blkin) registered for path 'src/blkin'
Submodule 'src/civetweb' (https://github.com/ceph/civetweb) registered for path 'src/civetweb'
Submodule 'src/crypto/isa-l/isa-l_crypto' (https://github.com/01org/isa-l_crypto) registered for path 'src/crypto/isa-l/isa-l_crypto'
Submodule 'src/dpdk' (https://github.com/ceph/dpdk) registered for path 'src/dpdk'
Submodule 'src/erasure-code/jerasure/gf-complete' (https://github.com/ceph/gf-complete.git) registered for path 'src/erasure-code/jerasure/gf-complete'
Submodule 'src/erasure-code/jerasure/jerasure' (https://github.com/ceph/jerasure.git) registered for path 'src/erasure-code/jerasure/jerasure'
Submodule 'src/googletest' (https://github.com/ceph/googletest) registered for path 'src/googletest'
Submodule 'src/isa-l' (https://github.com/ceph/isa-l) registered for path 'src/isa-l'
Submodule 'src/lua' (https://github.com/ceph/lua.git) registered for path 'src/lua'
Submodule 'src/rapidjson' (https://github.com/ceph/rapidjson) registered for path 'src/rapidjson'
Submodule 'src/rocksdb' (https://github.com/ceph/rocksdb) registered for path 'src/rocksdb'
Submodule 'src/spdk' (https://github.com/ceph/spdk.git) registered for path 'src/spdk'
Submodule 'src/xxHash' (https://github.com/ceph/xxHash.git) registered for path 'src/xxHash'
Submodule 'src/zstd' (https://github.com/facebook/zstd) registered for path 'src/zstd'
Cloning into '/build/ceph/ceph-erasure-code-corpus'...
Cloning into '/build/ceph/ceph-object-corpus'...
Cloning into '/build/ceph/src/Beast'...
Cloning into '/build/ceph/src/blkin'...
Cloning into '/build/ceph/src/civetweb'...
Cloning into '/build/ceph/src/crypto/isa-l/isa-l_crypto'...
Cloning into '/build/ceph/src/dpdk'...
Cloning into '/build/ceph/src/erasure-code/jerasure/gf-complete'...
Cloning into '/build/ceph/src/erasure-code/jerasure/jerasure'...
Cloning into '/build/ceph/src/googletest'...
Cloning into '/build/ceph/src/isa-l'...
Cloning into '/build/ceph/src/lua'...
Cloning into '/build/ceph/src/rapidjson'...
Cloning into '/build/ceph/src/rocksdb'...
Cloning into '/build/ceph/src/spdk'...
Cloning into '/build/ceph/src/xxHash'...
Cloning into '/build/ceph/src/zstd'...
Submodule path 'ceph-erasure-code-corpus': checked out '2d7d78b9cc52e8a9529d8cc2d2954c7d375d5dd7'
Submodule path 'ceph-object-corpus': checked out 'f0ba19f9d8792d308b3d2f746cff85436f87c4ae'
Submodule path 'src/Beast': checked out 'd8db5f1a0d607aa78e6e857daa0410b0d5326978'
Submodule path 'src/blkin': checked out 'f24ceec055ea236a093988237a9821d145f5f7c8'
Submodule path 'src/civetweb': checked out 'de23828cea740032428ce6f288ad061b429ce68e'
Submodule path 'src/crypto/isa-l/isa-l_crypto': checked out '603529a4e06ac8a1662c13d6b31f122e21830352'
Submodule path 'src/dpdk': checked out 'd3bfeaaabfd37ecc7d7d2edf2a3b60cc2d913cc9'
Submodule path 'src/erasure-code/jerasure/gf-complete': checked out '7e61b44404f0ed410c83cfd3947a52e88ae044e1'
Submodule path 'src/erasure-code/jerasure/jerasure': checked out '96c76b89d661c163f65a014b8042c9354ccf7f31'
Submodule path 'src/googletest': checked out 'fdb850479284e2aae047b87df6beae84236d0135'
Submodule path 'src/isa-l': checked out '7e1a337433a340bc0974ed0f04301bdaca374af6'
Submodule path 'src/lua': checked out '1fce39c6397056db645718b8f5821571d97869a4'
Submodule path 'src/rapidjson': checked out 'f54b0e47a08782a6131cc3d60f94d038fa6e0a51'
Submodule 'thirdparty/gtest' (https://github.com/google/googletest.git) registered for path 'src/rapidjson/thirdparty/gtest'
Cloning into '/build/ceph/src/rapidjson/thirdparty/gtest'...
Submodule path 'src/rapidjson/thirdparty/gtest': checked out '0a439623f75c029912728d80cb7f1b8b48739ca4'
Submodule path 'src/rocksdb': checked out 'bac0af364a2c742ea9bb63199ddeee0edacdb2ac'
Submodule path 'src/spdk': checked out '5742e9b9e782b9666e10c9224389e4d015c3cdee'
Submodule path 'src/xxHash': checked out '1f40c6511fa8dd9d2e337ca8c9bc18b3e87663c9'
Submodule path 'src/zstd': checked out 'dc9931205a87fdf6b993b3e215139011d3ef9fee'
devmapper: Thin Pool has 50085 free data blocks which is less than minimum required 52428 free data blocks. Create more free space in thin pool or use dm.min_free_space option to change behavior
Makefile:59: recipe for target 'do.build' failed
make[2]: *** [do.build] Error 1
Makefile:30: recipe for target 'ceph.linux_amd64' failed
make[1]: *** [ceph.linux_amd64] Error 2
Makefile:113: recipe for target 'build' failed
make: *** [build] Error 2
```

Docker storage have not enough space

## Fedora 27 cloud base

[vagrant up](./vagrantup.md), Vagrantfile for Windows10 reside under rook repository

no VPN: copy rook repository (vendor directory generated) into windows 10.

go
```
[vagrant@localhost ~]$ sudo tar -C /usr/local -xzf /windows10_drive_f/99-mirror/https0x3A0x2F0x2Fgolang.org0x2Fdl/go1.9.2.linux-amd64.tar.gz 
```

```
[vagrant@localhost ~]$ go version
go version go1.9.2 linux/amd64
```

```
[vagrant@localhost ~]$ go env
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/windows10_drive_g/work"
GORACE=""
GOROOT="/usr/local/go"
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build256069898=/tmp/go-build -gno-record-gcc-switches"
CXX="g++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
```

Docker
```
[vagrant@localhost ~]$ sudo dnf install -y docker
Last metadata expiration check: 0:53:00 ago on Thu 14 Dec 2017 08:28:03 AM UTC.
Dependencies resolved.
================================================================================================================================================
 Package                                       Arch                   Version                                     Repository               Size
================================================================================================================================================
Installing:
 docker                                        x86_64                 2:1.13.1-42.git4402c09.fc27                 updates                  20 M
Installing dependencies:
 atomic-registries                             x86_64                 1.20.1-6.fc27                               updates                  38 k
 container-selinux                             noarch                 2:2.36-1.fc27                               updates                  36 k
 container-storage-setup                       noarch                 0.8.0-2.git1d27ecf.fc27                     updates                  36 k
 device-mapper-event                           x86_64                 1.02.142-4.fc27                             fedora                  251 k
 device-mapper-event-libs                      x86_64                 1.02.142-4.fc27                             fedora                  251 k
 device-mapper-persistent-data                 x86_64                 0.7.5-1.fc27                                updates                 429 k
 docker-common                                 x86_64                 2:1.13.1-42.git4402c09.fc27                 updates                  85 k
 docker-rhel-push-plugin                       x86_64                 2:1.13.1-42.git4402c09.fc27                 updates                 1.7 M
 gnupg                                         x86_64                 1.4.22-3.fc27                               fedora                  1.3 M
 iptables                                      x86_64                 1.6.1-4.fc27                                fedora                  471 k
 libaio                                        x86_64                 0.3.110-9.fc27                              fedora                   29 k
 libnet                                        x86_64                 1.1.6-14.fc27                               fedora                   65 k
 libnetfilter_conntrack                        x86_64                 1.0.6-4.fc27                                fedora                   62 k
 libnfnetlink                                  x86_64                 1.0.1-11.fc27                               fedora                   31 k
 libusb                                        x86_64                 1:0.1.5-10.fc27                             fedora                   40 k
 lvm2                                          x86_64                 2.02.173-4.fc27                             fedora                  1.4 M
 lvm2-libs                                     x86_64                 2.02.173-4.fc27                             fedora                  1.1 M
 oci-umount                                    x86_64                 2:2.3.0-1.git51e7c50.fc27                   updates                  35 k
 policycoreutils-python-utils                  x86_64                 2.7-1.fc27                                  fedora                  223 k
 protobuf-c                                    x86_64                 1.2.1-7.fc27                                fedora                   34 k
 python-rhsm-certificates                      x86_64                 1.20.1-3.fc27                               fedora                   44 k
 python3-pytoml                                noarch                 0.1.14-2.git7dea353.fc27                    fedora                   23 k
 skopeo-containers                             x86_64                 0.1.27-1.git93876ac.fc27                    updates                  16 k
 systemd-container                             x86_64                 234-8.fc27                                  fedora                  422 k
 yajl                                          x86_64                 2.1.0-8.fc27                                fedora                   38 k
Installing weak dependencies:
 criu                                          x86_64                 3.6-1.fc27                                  updates                 456 k
 oci-register-machine                          x86_64                 0-5.11.gitcd1e331.fc27                      fedora                  1.0 M
 oci-systemd-hook                              x86_64                 1:0.1.13-1.gitafe4b4a.fc27                  fedora                   36 k

Transaction Summary
================================================================================================================================================
Install  29 Packages

Total download size: 30 M
Installed size: 95 M
Downloading Packages:
(1/29): atomic-registries-1.20.1-6.fc27.x86_64.rpm                                                               81 kB/s |  38 kB     00:00    
(2/29): docker-common-1.13.1-42.git4402c09.fc27.x86_64.rpm                                                      148 kB/s |  85 kB     00:00    
(3/29): skopeo-containers-0.1.27-1.git93876ac.fc27.x86_64.rpm                                                   343 kB/s |  16 kB     00:00    
(4/29): docker-rhel-push-plugin-1.13.1-42.git4402c09.fc27.x86_64.rpm                                            2.0 MB/s | 1.7 MB     00:00    
(5/29): python-rhsm-certificates-1.20.1-3.fc27.x86_64.rpm                                                       137 kB/s |  44 kB     00:00    
(6/29): python3-pytoml-0.1.14-2.git7dea353.fc27.noarch.rpm                                                       80 kB/s |  23 kB     00:00    
(7/29): libusb-0.1.5-10.fc27.x86_64.rpm                                                                         136 kB/s |  40 kB     00:00    
(8/29): gnupg-1.4.22-3.fc27.x86_64.rpm                                                                          696 kB/s | 1.3 MB     00:01    
(9/29): lvm2-libs-2.02.173-4.fc27.x86_64.rpm                                                                    1.1 MB/s | 1.1 MB     00:01    
(10/29): device-mapper-event-1.02.142-4.fc27.x86_64.rpm                                                         519 kB/s | 251 kB     00:00    
(11/29): device-mapper-event-libs-1.02.142-4.fc27.x86_64.rpm                                                    843 kB/s | 251 kB     00:00    
(12/29): device-mapper-persistent-data-0.7.5-1.fc27.x86_64.rpm                                                  683 kB/s | 429 kB     00:00    
(13/29): libaio-0.3.110-9.fc27.x86_64.rpm                                                                       343 kB/s |  29 kB     00:00    
(14/29): container-selinux-2.36-1.fc27.noarch.rpm                                                               103 kB/s |  36 kB     00:00    
(15/29): lvm2-2.02.173-4.fc27.x86_64.rpm                                                                        386 kB/s | 1.4 MB     00:03    
(16/29): policycoreutils-python-utils-2.7-1.fc27.x86_64.rpm                                                     693 kB/s | 223 kB     00:00    
(17/29): container-storage-setup-0.8.0-2.git1d27ecf.fc27.noarch.rpm                                             472 kB/s |  36 kB     00:00    
(18/29): oci-umount-2.3.0-1.git51e7c50.fc27.x86_64.rpm                                                          251 kB/s |  35 kB     00:00    
(19/29): yajl-2.1.0-8.fc27.x86_64.rpm                                                                           404 kB/s |  38 kB     00:00    
(20/29): libnetfilter_conntrack-1.0.6-4.fc27.x86_64.rpm                                                         608 kB/s |  62 kB     00:00    
(21/29): libnfnetlink-1.0.1-11.fc27.x86_64.rpm                                                                  124 kB/s |  31 kB     00:00    
(22/29): docker-1.13.1-42.git4402c09.fc27.x86_64.rpm                                                            3.1 MB/s |  20 MB     00:06    
(23/29): iptables-1.6.1-4.fc27.x86_64.rpm                                                                       703 kB/s | 471 kB     00:00    
(24/29): libnet-1.1.6-14.fc27.x86_64.rpm                                                                        408 kB/s |  65 kB     00:00    
(25/29): criu-3.6-1.fc27.x86_64.rpm                                                                             1.3 MB/s | 456 kB     00:00    
(26/29): protobuf-c-1.2.1-7.fc27.x86_64.rpm                                                                     246 kB/s |  34 kB     00:00    
(27/29): oci-systemd-hook-0.1.13-1.gitafe4b4a.fc27.x86_64.rpm                                                   182 kB/s |  36 kB     00:00    
(28/29): systemd-container-234-8.fc27.x86_64.rpm                                                                997 kB/s | 422 kB     00:00    
(29/29): oci-register-machine-0-5.11.gitcd1e331.fc27.x86_64.rpm                                                 1.2 MB/s | 1.0 MB     00:00    
------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                           1.9 MB/s |  30 MB     00:15     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                        1/1 
  Installing       : device-mapper-event-libs-1.02.142-4.fc27.x86_64                                                                       1/29 
  Running scriptlet: device-mapper-event-libs-1.02.142-4.fc27.x86_64                                                                       1/29 
  Installing       : libnfnetlink-1.0.1-11.fc27.x86_64                                                                                     2/29 
  Running scriptlet: libnfnetlink-1.0.1-11.fc27.x86_64                                                                                     2/29 
  Installing       : yajl-2.1.0-8.fc27.x86_64                                                                                              3/29 
  Running scriptlet: yajl-2.1.0-8.fc27.x86_64                                                                                              3/29 
  Installing       : oci-umount-2:2.3.0-1.git51e7c50.fc27.x86_64                                                                           4/29 
  Installing       : libnetfilter_conntrack-1.0.6-4.fc27.x86_64                                                                            5/29 
  Running scriptlet: libnetfilter_conntrack-1.0.6-4.fc27.x86_64                                                                            5/29 
  Installing       : iptables-1.6.1-4.fc27.x86_64                                                                                          6/29 
  Running scriptlet: iptables-1.6.1-4.fc27.x86_64                                                                                          6/29 
  Installing       : device-mapper-event-1.02.142-4.fc27.x86_64                                                                            7/29 
  Running scriptlet: device-mapper-event-1.02.142-4.fc27.x86_64                                                                            7/29 
  Installing       : lvm2-libs-2.02.173-4.fc27.x86_64                                                                                      8/29 
  Running scriptlet: lvm2-libs-2.02.173-4.fc27.x86_64                                                                                      8/29 
  Installing       : systemd-container-234-8.fc27.x86_64                                                                                   9/29 
  Installing       : protobuf-c-1.2.1-7.fc27.x86_64                                                                                       10/29 
  Running scriptlet: protobuf-c-1.2.1-7.fc27.x86_64                                                                                       10/29 
  Installing       : libnet-1.1.6-14.fc27.x86_64                                                                                          11/29 
  Running scriptlet: libnet-1.1.6-14.fc27.x86_64                                                                                          11/29 
  Installing       : policycoreutils-python-utils-2.7-1.fc27.x86_64                                                                       12/29 
  Installing       : container-selinux-2:2.36-1.fc27.noarch                                                                               13/29 
  Running scriptlet: container-selinux-2:2.36-1.fc27.noarch                                                                               13/29 
  Installing       : libaio-0.3.110-9.fc27.x86_64                                                                                         14/29 
  Running scriptlet: libaio-0.3.110-9.fc27.x86_64                                                                                         14/29 
  Installing       : device-mapper-persistent-data-0.7.5-1.fc27.x86_64                                                                    15/29 
  Installing       : lvm2-2.02.173-4.fc27.x86_64                                                                                          16/29 
  Running scriptlet: lvm2-2.02.173-4.fc27.x86_64                                                                                          16/29 
  Installing       : container-storage-setup-0.8.0-2.git1d27ecf.fc27.noarch                                                               17/29 
  Installing       : libusb-1:0.1.5-10.fc27.x86_64                                                                                        18/29 
  Running scriptlet: libusb-1:0.1.5-10.fc27.x86_64                                                                                        18/29 
  Installing       : gnupg-1.4.22-3.fc27.x86_64                                                                                           19/29 
  Running scriptlet: gnupg-1.4.22-3.fc27.x86_64                                                                                           19/29 
  Installing       : python3-pytoml-0.1.14-2.git7dea353.fc27.noarch                                                                       20/29 
  Installing       : atomic-registries-1.20.1-6.fc27.x86_64                                                                               21/29 
  Installing       : python-rhsm-certificates-1.20.1-3.fc27.x86_64                                                                        22/29 
  Installing       : skopeo-containers-0.1.27-1.git93876ac.fc27.x86_64                                                                    23/29 
  Installing       : docker-rhel-push-plugin-2:1.13.1-42.git4402c09.fc27.x86_64                                                           24/29 
  Running scriptlet: docker-rhel-push-plugin-2:1.13.1-42.git4402c09.fc27.x86_64                                                           24/29 
  Installing       : docker-common-2:1.13.1-42.git4402c09.fc27.x86_64                                                                     25/29 
  Installing       : docker-2:1.13.1-42.git4402c09.fc27.x86_64                                                                            26/29 
  Running scriptlet: docker-2:1.13.1-42.git4402c09.fc27.x86_64                                                                            26/29 
  Installing       : criu-3.6-1.fc27.x86_64                                                                                               27/29 
  Running scriptlet: criu-3.6-1.fc27.x86_64                                                                                               27/29 
  Installing       : oci-register-machine-0-5.11.gitcd1e331.fc27.x86_64                                                                   28/29 
  Installing       : oci-systemd-hook-1:0.1.13-1.gitafe4b4a.fc27.x86_64                                                                   29/29 
  Running scriptlet: docker-2:1.13.1-42.git4402c09.fc27.x86_64                                                                            29/29 
  Running scriptlet: oci-systemd-hook-1:0.1.13-1.gitafe4b4a.fc27.x86_64                                                                   29/29 
  Verifying        : docker-2:1.13.1-42.git4402c09.fc27.x86_64                                                                             1/29 
  Verifying        : atomic-registries-1.20.1-6.fc27.x86_64                                                                                2/29 
  Verifying        : docker-common-2:1.13.1-42.git4402c09.fc27.x86_64                                                                      3/29 
  Verifying        : docker-rhel-push-plugin-2:1.13.1-42.git4402c09.fc27.x86_64                                                            4/29 
  Verifying        : skopeo-containers-0.1.27-1.git93876ac.fc27.x86_64                                                                     5/29 
  Verifying        : gnupg-1.4.22-3.fc27.x86_64                                                                                            6/29 
  Verifying        : python-rhsm-certificates-1.20.1-3.fc27.x86_64                                                                         7/29 
  Verifying        : python3-pytoml-0.1.14-2.git7dea353.fc27.noarch                                                                        8/29 
  Verifying        : libusb-1:0.1.5-10.fc27.x86_64                                                                                         9/29 
  Verifying        : lvm2-2.02.173-4.fc27.x86_64                                                                                          10/29 
  Verifying        : lvm2-libs-2.02.173-4.fc27.x86_64                                                                                     11/29 
  Verifying        : device-mapper-event-1.02.142-4.fc27.x86_64                                                                           12/29 
  Verifying        : device-mapper-event-libs-1.02.142-4.fc27.x86_64                                                                      13/29 
  Verifying        : device-mapper-persistent-data-0.7.5-1.fc27.x86_64                                                                    14/29 
  Verifying        : libaio-0.3.110-9.fc27.x86_64                                                                                         15/29 
  Verifying        : container-selinux-2:2.36-1.fc27.noarch                                                                               16/29 
  Verifying        : policycoreutils-python-utils-2.7-1.fc27.x86_64                                                                       17/29 
  Verifying        : container-storage-setup-0.8.0-2.git1d27ecf.fc27.noarch                                                               18/29 
  Verifying        : oci-umount-2:2.3.0-1.git51e7c50.fc27.x86_64                                                                          19/29 
  Verifying        : yajl-2.1.0-8.fc27.x86_64                                                                                             20/29 
  Verifying        : iptables-1.6.1-4.fc27.x86_64                                                                                         21/29 
  Verifying        : libnetfilter_conntrack-1.0.6-4.fc27.x86_64                                                                           22/29 
  Verifying        : libnfnetlink-1.0.1-11.fc27.x86_64                                                                                    23/29 
  Verifying        : criu-3.6-1.fc27.x86_64                                                                                               24/29 
  Verifying        : libnet-1.1.6-14.fc27.x86_64                                                                                          25/29 
  Verifying        : protobuf-c-1.2.1-7.fc27.x86_64                                                                                       26/29 
  Verifying        : oci-register-machine-0-5.11.gitcd1e331.fc27.x86_64                                                                   27/29 
  Verifying        : oci-systemd-hook-1:0.1.13-1.gitafe4b4a.fc27.x86_64                                                                   28/29 
  Verifying        : systemd-container-234-8.fc27.x86_64                                                                                  29/29 

Installed:
  docker.x86_64 2:1.13.1-42.git4402c09.fc27                            criu.x86_64 3.6-1.fc27                                                  
  oci-register-machine.x86_64 0-5.11.gitcd1e331.fc27                   oci-systemd-hook.x86_64 1:0.1.13-1.gitafe4b4a.fc27                      
  atomic-registries.x86_64 1.20.1-6.fc27                               container-selinux.noarch 2:2.36-1.fc27                                  
  container-storage-setup.noarch 0.8.0-2.git1d27ecf.fc27               device-mapper-event.x86_64 1.02.142-4.fc27                              
  device-mapper-event-libs.x86_64 1.02.142-4.fc27                      device-mapper-persistent-data.x86_64 0.7.5-1.fc27                       
  docker-common.x86_64 2:1.13.1-42.git4402c09.fc27                     docker-rhel-push-plugin.x86_64 2:1.13.1-42.git4402c09.fc27              
  gnupg.x86_64 1.4.22-3.fc27                                           iptables.x86_64 1.6.1-4.fc27                                            
  libaio.x86_64 0.3.110-9.fc27                                         libnet.x86_64 1.1.6-14.fc27                                             
  libnetfilter_conntrack.x86_64 1.0.6-4.fc27                           libnfnetlink.x86_64 1.0.1-11.fc27                                       
  libusb.x86_64 1:0.1.5-10.fc27                                        lvm2.x86_64 2.02.173-4.fc27                                             
  lvm2-libs.x86_64 2.02.173-4.fc27                                     oci-umount.x86_64 2:2.3.0-1.git51e7c50.fc27                             
  policycoreutils-python-utils.x86_64 2.7-1.fc27                       protobuf-c.x86_64 1.2.1-7.fc27                                          
  python-rhsm-certificates.x86_64 1.20.1-3.fc27                        python3-pytoml.noarch 0.1.14-2.git7dea353.fc27                          
  skopeo-containers.x86_64 0.1.27-1.git93876ac.fc27                    systemd-container.x86_64 234-8.fc27                                     
  yajl.x86_64 2.1.0-8.fc27                                            

Complete!
```

```
[vagrant@localhost ~]$ sudo systemctl start docker.service
[vagrant@localhost ~]$ sudo systemctl status docker.service
● docker.service - Docker Application Container Engine
   Loaded: loaded (/usr/lib/systemd/system/docker.service; disabled; vendor preset: disabled)
   Active: active (running) since Thu 2017-12-14 09:24:41 UTC; 7s ago
     Docs: http://docs.docker.com
 Main PID: 3466 (dockerd-current)
    Tasks: 7 (limit: 8192)
   CGroup: /system.slice/docker.service
           └─3466 /usr/bin/dockerd-current --add-runtime oci=/usr/libexec/docker/docker-runc-current --default-runtime=oci --authorization-plugi

Dec 14 09:24:41 localhost dockerd-current[3466]: time="2017-12-14T09:24:41.076848970Z" level=warning msg="Your kernel does not support cgroup rt
Dec 14 09:24:41 localhost dockerd-current[3466]: time="2017-12-14T09:24:41.077271266Z" level=info msg="Loading containers: start."
Dec 14 09:24:41 localhost dockerd-current[3466]: time="2017-12-14T09:24:41.111410266Z" level=info msg="Firewalld running: false"
Dec 14 09:24:41 localhost dockerd-current[3466]: time="2017-12-14T09:24:41.159277271Z" level=info msg="Default bridge (docker0) is assigned with
Dec 14 09:24:41 localhost dockerd-current[3466]: time="2017-12-14T09:24:41.191622587Z" level=info msg="Loading containers: done."
Dec 14 09:24:41 localhost dockerd-current[3466]: time="2017-12-14T09:24:41.200427424Z" level=warning msg="failed to retrieve docker-init version
Dec 14 09:24:41 localhost dockerd-current[3466]: time="2017-12-14T09:24:41.200740784Z" level=info msg="Daemon has completed initialization"
Dec 14 09:24:41 localhost dockerd-current[3466]: time="2017-12-14T09:24:41.200865437Z" level=info msg="Docker daemon" commit=fbadd78-unsupported
Dec 14 09:24:41 localhost systemd[1]: Started Docker Application Container Engine.
Dec 14 09:24:41 localhost dockerd-current[3466]: time="2017-12-14T09:24:41.215892760Z" level=info msg="API listen on /var/run/docker.sock"
[vagrant@localhost ~]$ sudo systemctl enable docker.service
Created symlink /etc/systemd/system/multi-user.target.wants/docker.service → /usr/lib/systemd/system/docker.service.
```

```
[vagrant@localhost ~]$ sudo usermod -aG root vagrant
[vagrant@localhost ~]$ exit
logout
Connection to 10.64.33.198 closed.

tangf@DESKTOP-H68OQDV ~
$ ssh vagrant@10.64.33.198
vagrant@10.64.33.198's password: 
Last login: Thu Dec 14 09:18:13 2017 from 10.64.33.1
[vagrant@localhost ~]$ docker info
Containers: 0
 Running: 0
 Paused: 0
 Stopped: 0
Images: 0
Server Version: 1.13.1
Storage Driver: overlay2
 Backing Filesystem: extfs
 Supports d_type: true
 Native Overlay Diff: true
Logging Driver: journald
Cgroup Driver: systemd
Plugins: 
 Volume: local
 Network: bridge host macvlan null overlay
 Authorization: rhel-push-plugin
Swarm: inactive
Runtimes: oci runc
Default Runtime: oci
Init Binary: /usr/libexec/docker/docker-init-current
containerd version: fbadd789ddf86a4be9d6905528b7486c61e52612 (expected: aa8187dbd3b7ad67d8e5e3a15115d3eef43a7ed1)
runc version: fbadd789ddf86a4be9d6905528b7486c61e52612-dirty (expected: 9df8b306d01f59d3a8029be411de015b7304dd8f)
init version: N/A (expected: 949e6facb77383876aeff8a6944dde66b3089574)
Security Options:
 seccomp
  WARNING: You're not using the default seccomp profile
  Profile: /etc/docker/seccomp.json
 selinux
Kernel Version: 4.13.9-300.fc27.x86_64
Operating System: Fedora 27 (Cloud Edition)
OSType: linux
Architecture: x86_64
Number of Docker Hooks: 3
CPUs: 1
Total Memory: 4.841 GiB
Name: localhost
ID: EC6N:7VB6:5O3P:DCLT:L6R5:OPUC:LIIL:DIC2:LC6G:ARKQ:BSIH:IVE6
Docker Root Dir: /var/lib/docker
Debug Mode (client): false
Debug Mode (server): false
Registry: https://index.docker.io/v1/
Experimental: false
Insecure Registries:
 127.0.0.0/8
Live Restore Enabled: false
Registries: docker.io (secure), registry.fedoraproject.org (secure), registry.access.redhat.com (secure), docker.io (secure)
```

```
[vagrant@localhost rook]$ sudo dnf install -yq git
```

```
[vagrant@localhost rook]$ mkdir /home/vagrant/bin
```

copy `/bin/shasum` from old Fedora 23
```
[vagrant@localhost rook]$ sudo dnf install -yq perl-CPAN
```

```
=== go vet
=== go build linux_amd64
github.com/rook/rook/cmd/rookctl
github.com/rook/rook/cmd/rook
github.com/rook/rook/cmd/rookflex
github.com/rook/rook/tests/integration
testmain
github.com/rook/rook/tests/longhaul
testmain
=== docker build build-be9394dd/rootfs-builder-amd64
sha256:98e231c352a93ac7ba5c9c5ba6e9e0152ca0bc4ca4f4433274705f757cee1303
=== docker build build-be9394dd/base-amd64
Sending build context to Docker daemon 3.072 kB
Step 1/7 : FROM build-be9394dd/rootfs-amd64-98e231c352a9
 ---> 521831cd5775
Step 2/7 : RUN HOST_IP=$(route -n | awk '/^0.0.0.0/ {print $2}') &&     echo "#!/bin/bash" > /usr/local/bin/apt-ng-host-discover &&     echo "if nc -w1 -z $HOST_IP 3142; then printf http://$HOST_IP:3142; else printf DIRECT; fi" >> /usr/local/bin/apt-ng-host-discover &&     chmod +x /usr/local/bin/apt-ng-host-discover &&     echo 'Acquire::http::ProxyAutoDetect "/usr/local/bin/apt-ng-host-discover";' > /etc/apt/apt.conf.d/30proxy
 ---> Using cache
 ---> dadcf422b385
Step 3/7 : ARG ARCH
 ---> Using cache
 ---> 0507640f4b14
Step 4/7 : ADD https://github.com/krallin/tini/releases/download/v0.14.0/tini-${ARCH} /tini
Get https://github-production-release-asset-2e65be.s3.amazonaws.com/31432573/c21bad02-eba8-11e6-9853-18020eb6f4a7?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20171214%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20171214T105948Z&X-Amz-Expires=300&X-Amz-Signature=8c4fae81baa6149e7a5489a4a857e4200d3e115945058c5e5bd272fbb91fc9d7&X-Amz-SignedHeaders=host&actor_id=0&response-content-disposition=attachment%3B%20filename%3Dtini-amd64&response-content-type=application%2Foctet-stream: dial tcp 52.216.80.120:443: i/o timeout
make[2]: *** [Makefile:54: do.build] Error 1
make[1]: *** [Makefile:27: base.linux_amd64] Error 2
make: *** [Makefile:114: build] Error 2
```

amazonaws!!!

run mini http
```
[vagrant@localhost /] /windows10_drive_f/99-mirror/linux-bin/gofileserver
Listening at :48080
```

modify Dockerfile of ./images/base line 28
```
#ADD https://github.com/krallin/tini/releases/download/v0.14.0/tini-${ARCH} /tini
ADD http://localhost:48080/windows10_drive_f/99-mirror/https0x3A0x2F0x2Fgithub.com0x2Fkrallin0x2Ftini0x2Freleases0x2Fdownload/v0.14.0/tini-${ARCH} /tini
```

```
[vagrant@localhost rook]$ make -j4
=== go vet
=== go build linux_amd64
github.com/rook/rook/cmd/rookctl
github.com/rook/rook/cmd/rook
github.com/rook/rook/cmd/rookflex
github.com/rook/rook/tests/integration
testmain
github.com/rook/rook/tests/longhaul
testmain
=== docker build build-be9394dd/rootfs-builder-amd64
sha256:8b040d7215165585e313697c8b484243132462a835973fca53e2880cacb6bcf0
=== docker build build-be9394dd/rootfs-amd64-8b040d721516
3824242a6ed8e61f25b4c56080253e431238083389d49761bcea6fb3438de908
sha256:2e3b912145097e5be9582284c3b04fd91bba7a35be074fa4a68010815869cc31
=== docker build build-be9394dd/base-amd64
sha256:bdc3253aab817a7ef3a1187e858f460cd6c48cc105cb8c7832cdd6fa282b7f23
=== caching image build-be9394dd/base-amd64
=== caching image build-be9394dd/rootfs-builder-amd64
=== caching image build-be9394dd/rootfs-amd64-8b040d721516
=== docker build build-be9394dd/cross-gnu-amd64
sha256:c4a654332a52cdbe64b1ba5ea19bbfc5e12c66bf8883708a17339a68662210f4
=== caching image build-be9394dd/cross-gnu-amd64
=== docker build build-be9394dd/ceph-builder-amd64-tcmalloc-release
sha256:b7d0de780ed5063c8235da394dac8d2e3e30e33d6e87bb490dd499ed1c5c7849
=== building ceph from build-be9394dd/ceph-builder-amd64-tcmalloc-release
make: the '-j' option requires a positive integer argument
Usage: make [options] [target] ...
Options:
  -b, -m                      Ignored for compatibility.
  -B, --always-make           Unconditionally make all targets.
  -C DIRECTORY, --directory=DIRECTORY
                              Change to DIRECTORY before doing anything.
  -d                          Print lots of debugging information.
  --debug[=FLAGS]             Print various types of debugging information.
  -e, --environment-overrides
                              Environment variables override makefiles.
  --eval=STRING               Evaluate STRING as a makefile statement.
  -f FILE, --file=FILE, --makefile=FILE
                              Read FILE as a makefile.
  -h, --help                  Print this message and exit.
  -i, --ignore-errors         Ignore errors from recipes.
  -I DIRECTORY, --include-dir=DIRECTORY
                              Search DIRECTORY for included makefiles.
  -j [N], --jobs[=N]          Allow N jobs at once; infinite jobs with no arg.
  -k, --keep-going            Keep going when some targets can't be made.
  -l [N], --load-average[=N], --max-load[=N]
                              Don't start multiple jobs unless load is below N.
  -L, --check-symlink-times   Use the latest mtime between symlinks and target.
  -n, --just-print, --dry-run, --recon
                              Don't actually run any recipe; just print them.
  -o FILE, --old-file=FILE, --assume-old=FILE
                              Consider FILE to be very old and don't remake it.
  -O[TYPE], --output-sync[=TYPE]
                              Synchronize output of parallel jobs by TYPE.
  -p, --print-data-base       Print make's internal database.
  -q, --question              Run no recipe; exit status says if up to date.
  -r, --no-builtin-rules      Disable the built-in implicit rules.
  -R, --no-builtin-variables  Disable the built-in variable settings.
  -s, --silent, --quiet       Don't echo recipes.
  -S, --no-keep-going, --stop
                              Turns off -k.
  -t, --touch                 Touch targets instead of remaking them.
  --trace                     Print tracing information.
  -v, --version               Print the version number of make and exit.
  -w, --print-directory       Print the current directory.
  --no-print-directory        Turn off -w, even if it was turned on implicitly.
  -W FILE, --what-if=FILE, --new-file=FILE, --assume-new=FILE
                              Consider FILE to be infinitely new.
  --warn-undefined-variables  Warn when an undefined variable is referenced.

This program built for x86_64-pc-linux-gnu
Report bugs to <bug-make@gnu.org>
make[4]: *** [Makefile:76: ceph] Error 2
make[3]: *** [../image.mk:134: cache.lookup] Error 2
make[2]: *** [Makefile:63: do.build] Error 2
make[1]: *** [Makefile:30: ceph.linux_amd64] Error 2
make: *** [Makefile:114: build] Error 2
```

modify scipts of ./images/ceph/build-ceph.sh line 20
```
#makeargs=-j$(( $(nproc) - 1 ))
makeargs=-j1
```

Issue: SELinux

Enter directory of images
```
[vagrant@localhost images]$ make V=1 build
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/rootfs-builder-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/3 : FROM ubuntu:zesty
Trying to pull repository docker.io/library/ubuntu ... 
sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56: Pulling from docker.io/library/ubuntu
Digest: sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56
Status: Image is up to date for docker.io/ubuntu:zesty
 ---> fe1cc5b91830
Step 2/3 : RUN sed -i -e 's/^deb-src/#deb-src/' /etc/apt/sources.list &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy --no-install-recommends         net-tools         netcat &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -yy --no-install-recommends &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean -y
 ---> Using cache
 ---> 179e64f482d5
Step 3/3 : RUN tar -czf /usr/share/copyrights.tar.xz /usr/share/doc/*/copyright &&     rm -rf         /usr/share/doc         /usr/share/man         /usr/share/info         /usr/share/locale         /var/lib/apt/lists/*         /var/log/*         /var/cache/debconf/*         /etc/systemd         /lib/lsb         /lib/udev         /usr/lib/x86_64-linux-gnu/gconv/IBM*         /usr/lib/x86_64-linux-gnu/gconv/EBC* &&     mkdir -p /usr/share/man/man{1..8}
 ---> Using cache
 ---> 8b040d721516
Successfully built 8b040d721516
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/base-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/7 : FROM build-be9394dd/rootfs-amd64-8b040d721516
 ---> 2e3b91214509
Step 2/7 : RUN HOST_IP=$(route -n | awk '/^0.0.0.0/ {print $2}') &&     echo "#!/bin/bash" > /usr/local/bin/apt-ng-host-discover &&     echo "if nc -w1 -z $HOST_IP 3142; then printf http://$HOST_IP:3142; else printf DIRECT; fi" >> /usr/local/bin/apt-ng-host-discover &&     chmod +x /usr/local/bin/apt-ng-host-discover &&     echo 'Acquire::http::ProxyAutoDetect "/usr/local/bin/apt-ng-host-discover";' > /etc/apt/apt.conf.d/30proxy
 ---> Using cache
 ---> e3eb9ca4c9ff
Step 3/7 : ARG ARCH
 ---> Using cache
 ---> cb5e0f44ebbb
Step 4/7 : ADD http://localhost:48080/windows10_drive_f/99-mirror/https0x3A0x2F0x2Fgithub.com0x2Fkrallin0x2Ftini0x2Freleases0x2Fdownload/v0.14.0/tini-${ARCH} /tini
Downloading [==================================================>] 19.89 kB/19.89 kB
 ---> Using cache
 ---> 39a370fa6cb6
Step 5/7 : RUN chmod +x /tini
 ---> Using cache
 ---> 34fd3135dcaf
Step 6/7 : ENTRYPOINT /tini -g --
 ---> Using cache
 ---> 4c72f52634fd
Step 7/7 : CMD /bin/bash
 ---> Using cache
 ---> bdc3253aab81
Successfully built bdc3253aab81
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== caching image build-be9394dd/base-amd64
=== caching image build-be9394dd/rootfs-builder-amd64
=== caching image build-be9394dd/rootfs-amd64-8b040d721516
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== docker build build-be9394dd/cross-gnu-amd64
Sending build context to Docker daemon 13.82 kB
Step 1/5 : FROM build-be9394dd/base-amd64
 ---> bdc3253aab81
Step 2/5 : RUN echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty main universe restricted" > /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-backports main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-backports main universe restricted" >> /etc/apt/sources.list &&    dpkg --add-architecture armhf &&     dpkg --add-architecture arm64 &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         build-essential         ca-certificates         ccache         cmake         crossbuild-essential-arm64         crossbuild-essential-armhf         curl         git         jq         patch         yasm         zip &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> b4f4549a106a
Step 3/5 : ADD toolchain /usr/local/toolchain
 ---> Using cache
 ---> 3c2fc8b13df0
Step 4/5 : WORKDIR /build
 ---> Using cache
 ---> 5af7d7101090
Step 5/5 : ENV TERM xterm
 ---> Using cache
 ---> c4a654332a52
Successfully built c4a654332a52
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== caching image build-be9394dd/cross-gnu-amd64
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== docker build build-be9394dd/ceph-builder-amd64-tcmalloc-release
Sending build context to Docker daemon 14.85 kB
Step 1/16 : FROM build-be9394dd/cross-gnu-amd64
 ---> c4a654332a52
Step 2/16 : ARG CEPH_GIT_REPO
 ---> Using cache
 ---> d8a28e9b2948
Step 3/16 : RUN git clone ${CEPH_GIT_REPO} ceph
 ---> Using cache
 ---> ce6f3e2a9e94
Step 4/16 : ARG ARCH
 ---> Using cache
 ---> 10b0906f2885
Step 5/16 : ARG CROSS_TRIPLE
 ---> Using cache
 ---> 18fb3b2e6d58
Step 6/16 : RUN DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         python-setuptools         gperf                 cython:${ARCH}         python:${ARCH}         libaio-dev:${ARCH}         libatomic-ops-dev:${ARCH}         libbabeltrace-dev:${ARCH}         libblkid-dev:${ARCH}         libboost-context-dev:${ARCH}         libboost-coroutine-dev:${ARCH}         libboost-date-time-dev:${ARCH}         libboost-filesystem-dev:${ARCH}         libboost-iostreams-dev:${ARCH}         libboost-program-options-dev:${ARCH}         libboost-python-dev:${ARCH}         libboost-random-dev:${ARCH}         libboost-regex-dev:${ARCH}         libboost-system-dev:${ARCH}         libboost-thread-dev:${ARCH}         libcurl4-gnutls-dev:${ARCH}         libexpat1-dev:${ARCH}         libgoogle-perftools-dev:${ARCH}         libgoogle-perftools4:${ARCH}         libibverbs-dev:${ARCH}         libjemalloc-dev:${ARCH}         libkeyutils-dev:${ARCH}         libldap2-dev:${ARCH}         libnss3-dev:${ARCH}         libpython-dev:${ARCH}         libsnappy-dev:${ARCH}         libssl-dev:${ARCH}         libtcmalloc-minimal4:${ARCH}         libudev-dev:${ARCH}         libunwind-dev:${ARCH}         zlib1g-dev:${ARCH} &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> 8dc3989141e0
Step 7/16 : ARG CEPH_GIT_COMMIT
 ---> Using cache
 ---> 9d5d869cd549
Step 8/16 : RUN cd /build/ceph &&     git fetch --all --prune &&     git checkout -b ceph-builder ${CEPH_GIT_COMMIT} &&     git submodule update --init --recursive
 ---> Using cache
 ---> 1c0a3108f9d6
Step 9/16 : RUN mv /build/ceph/CMakeLists.txt /build/ceph/CMakeLists.original.txt
 ---> Using cache
 ---> 70dfa574cf25
Step 10/16 : ADD CMakeLists.txt /build/ceph
 ---> Using cache
 ---> 3b3c3ec08145
Step 11/16 : ARG CEPH_BUILD_TYPE
 ---> Using cache
 ---> 3f5087dd3915
Step 12/16 : ARG CEPH_ALLOCATOR
 ---> Using cache
 ---> 225c822b412c
Step 13/16 : RUN mkdir -p /build/ceph/build && cd /build/ceph/build &&     cmake     -DCMAKE_SKIP_INSTALL_ALL_DEPENDENCY=ON     -DCMAKE_INSTALL_PREFIX=/usr/local     -DCMAKE_BUILD_TYPE=${CEPH_BUILD_TYPE}     -DCMAKE_TOOLCHAIN_FILE=/usr/local/toolchain/${CROSS_TRIPLE}.cmake     -DALLOCATOR=${CEPH_ALLOCATOR}     -DWITH_SYSTEM_BOOST=ON     -DWITH_EMBEDDED=OFF     -DWITH_FUSE=OFF     -DWITH_LEVELDB=OFF     -DWITH_LTTNG=OFF     -DWITH_MANPAGE=OFF     -DWITH_PROFILER=OFF     -DWITH_PYTHON3=OFF     -DWITH_RADOSGW_FCGI_FRONTEND=OFF     ..
 ---> Using cache
 ---> 95367828ae35
Step 14/16 : WORKDIR /build/ceph/build
 ---> Using cache
 ---> 8c46aca7ea3c
Step 15/16 : ADD build-ceph.sh /build/
 ---> d9e5a61fdc5c
Removing intermediate container 568cb04c0318
Step 16/16 : CMD /build/build-ceph.sh
 ---> Running in 6eebe764c7cb
 ---> c19572dcbeb0
Removing intermediate container 6eebe764c7cb
Successfully built c19572dcbeb0
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[3]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== building ceph from build-be9394dd/ceph-builder-amd64-tcmalloc-release
Scanning dependencies of target civetweb_h
[  0%] keep civetweb.h up-to-date
[  0%] Built target civetweb_h
Scanning dependencies of target common_utf8
[  0%] Building C object src/CMakeFiles/common_utf8.dir/common/utf8.c.o
ccache: error: Failed to create directory /root/.ccache/tmp: Permission denied
src/CMakeFiles/common_utf8.dir/build.make:62: recipe for target 'src/CMakeFiles/common_utf8.dir/common/utf8.c.o' failed
make[3]: *** [src/CMakeFiles/common_utf8.dir/common/utf8.c.o] Error 1
CMakeFiles/Makefile2:640: recipe for target 'src/CMakeFiles/common_utf8.dir/all' failed
make[2]: *** [src/CMakeFiles/common_utf8.dir/all] Error 2
CMakeFiles/Makefile2:374: recipe for target 'CMakeFiles/rook.dir/rule' failed
make[1]: *** [CMakeFiles/rook.dir/rule] Error 2
Makefile:225: recipe for target 'rook' failed
make: *** [rook] Error 2
make[3]: *** [Makefile:76: ceph] Error 2
make[3]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[2]: *** [../image.mk:134: cache.lookup] Error 2
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[1]: *** [Makefile:63: do.build] Error 2
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make: *** [Makefile:30: ceph.linux_amd64] Error 2
```

Or 
```
[vagrant@localhost rook]$ make V=1 build
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook'
=== go vet
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rook/agent.go cmd/rook/api.go cmd/rook/main.go cmd/rook/mds.go cmd/rook/mgr.go cmd/rook/mon.go cmd/rook/operator.go cmd/rook/osd.go cmd/rook/rgw.go cmd/rook/version.go cmd/rook/mon_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/main.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/block/block.go cmd/rookctl/block/create.go cmd/rookctl/block/delete.go cmd/rookctl/block/list.go cmd/rookctl/block/log.go cmd/rookctl/block/map.go cmd/rookctl/block/unmap.go cmd/rookctl/block/block_test.go cmd/rookctl/block/create_test.go cmd/rookctl/block/delete_test.go cmd/rookctl/block/list_test.go cmd/rookctl/block/map_test.go cmd/rookctl/block/unmap_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/client/client.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/filesystem/create.go cmd/rookctl/filesystem/delete.go cmd/rookctl/filesystem/filesystem.go cmd/rookctl/filesystem/list.go cmd/rookctl/filesystem/mount.go cmd/rookctl/filesystem/unmount.go cmd/rookctl/filesystem/create_test.go cmd/rookctl/filesystem/delete_test.go cmd/rookctl/filesystem/list_test.go cmd/rookctl/filesystem/mount_test.go cmd/rookctl/filesystem/unmount_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/node/list.go cmd/rookctl/node/node.go cmd/rookctl/node/list_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/object/bucket.go cmd/rookctl/object/connection.go cmd/rookctl/object/create.go cmd/rookctl/object/delete.go cmd/rookctl/object/list.go cmd/rookctl/object/object.go cmd/rookctl/object/user.go cmd/rookctl/object/connection_test.go cmd/rookctl/object/create_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/pool/create.go cmd/rookctl/pool/delete.go cmd/rookctl/pool/list.go cmd/rookctl/pool/pool.go cmd/rookctl/pool/create_test.go cmd/rookctl/pool/list_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/rook/rook.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/status/status.go cmd/rookctl/status/status_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookctl/version/version.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookflex/main.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  cmd/rookflex/cmd/init.go cmd/rookflex/cmd/mount.go cmd/rookflex/cmd/root.go cmd/rookflex/cmd/unmount.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/apis/rook.io/register.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/apis/rook.io/v1alpha1/doc.go pkg/apis/rook.io/v1alpha1/placement.go pkg/apis/rook.io/v1alpha1/pool.go pkg/apis/rook.io/v1alpha1/register.go pkg/apis/rook.io/v1alpha1/storage.go pkg/apis/rook.io/v1alpha1/types.go pkg/apis/rook.io/v1alpha1/volumeattachment.go pkg/apis/rook.io/v1alpha1/zz_generated.deepcopy.go pkg/apis/rook.io/v1alpha1/placement_test.go pkg/apis/rook.io/v1alpha1/storage_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/client/clientset/versioned/clientset.go pkg/client/clientset/versioned/doc.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/client/clientset/versioned/fake/clientset_generated.go pkg/client/clientset/versioned/fake/doc.go pkg/client/clientset/versioned/fake/register.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/client/clientset/versioned/scheme/doc.go pkg/client/clientset/versioned/scheme/register.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/client/clientset/versioned/typed/rook/v1alpha1/cluster.go pkg/client/clientset/versioned/typed/rook/v1alpha1/doc.go pkg/client/clientset/versioned/typed/rook/v1alpha1/filesystem.go pkg/client/clientset/versioned/typed/rook/v1alpha1/generated_expansion.go pkg/client/clientset/versioned/typed/rook/v1alpha1/objectstore.go pkg/client/clientset/versioned/typed/rook/v1alpha1/pool.go pkg/client/clientset/versioned/typed/rook/v1alpha1/rook_client.go pkg/client/clientset/versioned/typed/rook/v1alpha1/volumeattachment.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/client/clientset/versioned/typed/rook/v1alpha1/fake/doc.go pkg/client/clientset/versioned/typed/rook/v1alpha1/fake/fake_cluster.go pkg/client/clientset/versioned/typed/rook/v1alpha1/fake/fake_filesystem.go pkg/client/clientset/versioned/typed/rook/v1alpha1/fake/fake_objectstore.go pkg/client/clientset/versioned/typed/rook/v1alpha1/fake/fake_pool.go pkg/client/clientset/versioned/typed/rook/v1alpha1/fake/fake_rook_client.go pkg/client/clientset/versioned/typed/rook/v1alpha1/fake/fake_volumeattachment.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/clusterd/context.go pkg/clusterd/disk.go pkg/clusterd/network.go pkg/clusterd/disk_test.go pkg/clusterd/network_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/agent/agent.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/agent/cluster/controller.go pkg/daemon/agent/cluster/controller_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/agent/flexvolume/controller.go pkg/daemon/agent/flexvolume/fake.go pkg/daemon/agent/flexvolume/server.go pkg/daemon/agent/flexvolume/types.go pkg/daemon/agent/flexvolume/controller_test.go pkg/daemon/agent/flexvolume/server_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/agent/flexvolume/attachment/crd.go pkg/daemon/agent/flexvolume/attachment/fake.go pkg/daemon/agent/flexvolume/attachment/resource.go pkg/daemon/agent/flexvolume/attachment/tpr.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/agent/flexvolume/manager/fake.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/agent/flexvolume/manager/ceph/manager.go pkg/daemon/agent/flexvolume/manager/ceph/manager_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/api/client.go pkg/daemon/api/daemon.go pkg/daemon/api/filesystem.go pkg/daemon/api/handlers.go pkg/daemon/api/image.go pkg/daemon/api/logger.go pkg/daemon/api/metrics.go pkg/daemon/api/node.go pkg/daemon/api/object.go pkg/daemon/api/pool.go pkg/daemon/api/router.go pkg/daemon/api/routes.go pkg/daemon/api/status.go pkg/daemon/api/filesystem_test.go pkg/daemon/api/handlers_test.go pkg/daemon/api/image_test.go pkg/daemon/api/object_test.go pkg/daemon/api/status_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/client/auth.go pkg/daemon/ceph/client/crush.go pkg/daemon/ceph/client/erasure-code-profile.go pkg/daemon/ceph/client/filesystem.go pkg/daemon/ceph/client/image.go pkg/daemon/ceph/client/log.go pkg/daemon/ceph/client/mon.go pkg/daemon/ceph/client/osd.go pkg/daemon/ceph/client/pool.go pkg/daemon/ceph/client/status.go pkg/daemon/ceph/client/usage.go pkg/daemon/ceph/client/crush_test.go pkg/daemon/ceph/client/erasure-code-profile_test.go pkg/daemon/ceph/client/filesystem_test.go pkg/daemon/ceph/client/image_test.go pkg/daemon/ceph/client/pool_test.go pkg/daemon/ceph/client/status_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/client/test/mon.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/collectors/health.go pkg/daemon/ceph/collectors/log.go pkg/daemon/ceph/collectors/mon.go pkg/daemon/ceph/collectors/osd.go pkg/daemon/ceph/collectors/pool.go pkg/daemon/ceph/collectors/usage.go pkg/daemon/ceph/collectors/health_test.go pkg/daemon/ceph/collectors/mon_test.go pkg/daemon/ceph/collectors/osd_test.go pkg/daemon/ceph/collectors/pool_test.go pkg/daemon/ceph/collectors/usage_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/mds/daemon.go pkg/daemon/ceph/mds/filesystem.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/mgr/daemon.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/mon/config.go pkg/daemon/ceph/mon/daemon.go pkg/daemon/ceph/mon/info.go pkg/daemon/ceph/mon/config_test.go pkg/daemon/ceph/mon/daemon_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/osd/agent.go pkg/daemon/ceph/osd/config.go pkg/daemon/ceph/osd/daemon.go pkg/daemon/ceph/osd/device.go pkg/daemon/ceph/osd/filesystem.go pkg/daemon/ceph/osd/health.go pkg/daemon/ceph/osd/scheme.go pkg/daemon/ceph/osd/agent_test.go pkg/daemon/ceph/osd/crushmap_test.go pkg/daemon/ceph/osd/daemon_test.go pkg/daemon/ceph/osd/device_test.go pkg/daemon/ceph/osd/filesystem_test.go pkg/daemon/ceph/osd/health_test.go pkg/daemon/ceph/osd/scheme_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/rgw/admin.go pkg/daemon/ceph/rgw/agent.go pkg/daemon/ceph/rgw/bucket.go pkg/daemon/ceph/rgw/daemon.go pkg/daemon/ceph/rgw/mime.go pkg/daemon/ceph/rgw/objectstore.go pkg/daemon/ceph/rgw/user.go pkg/daemon/ceph/rgw/daemon_test.go pkg/daemon/ceph/rgw/objectstore_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/test/info.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/daemon/ceph/util/util.go pkg/daemon/ceph/util/util_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/model/block-image.go pkg/model/client.go pkg/model/filesystem.go pkg/model/node.go pkg/model/object.go pkg/model/pool.go pkg/model/status.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/operator.go pkg/operator/operator_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/agent/agent.go pkg/operator/agent/types.go pkg/operator/agent/agent_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/cluster/controller.go pkg/operator/cluster/controller_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/cluster/api/api.go pkg/operator/cluster/api/api_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/cluster/mgr/mgr.go pkg/operator/cluster/mgr/mgr_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/cluster/mon/health.go pkg/operator/cluster/mon/mon.go pkg/operator/cluster/mon/spec.go pkg/operator/cluster/mon/util.go pkg/operator/cluster/mon/health_test.go pkg/operator/cluster/mon/mon_test.go pkg/operator/cluster/mon/spec_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/cluster/osd/osd.go pkg/operator/cluster/osd/osd_test.go pkg/operator/cluster/osd/spec_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/k8sutil/k8sutil.go pkg/operator/k8sutil/kvstore.go pkg/operator/k8sutil/pod.go pkg/operator/k8sutil/rbac.go pkg/operator/k8sutil/resources.go pkg/operator/k8sutil/volume.go pkg/operator/k8sutil/kvstore_test.go pkg/operator/k8sutil/pod_test.go pkg/operator/k8sutil/rbac_test.go pkg/operator/k8sutil/resources_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/mds/controller.go pkg/operator/mds/mds.go pkg/operator/mds/controller_test.go pkg/operator/mds/mds_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/pool/controller.go pkg/operator/pool/controller_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/provisioner/provisioner.go pkg/operator/provisioner/provisioner_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/provisioner/controller/controller.go pkg/operator/provisioner/controller/volume.go pkg/operator/provisioner/controller/controller_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/provisioner/controller/leaderelection/leaderelection.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/provisioner/controller/leaderelection/resourcelock/interface.go pkg/operator/provisioner/controller/leaderelection/resourcelock/provisionpvclock.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/rgw/controller.go pkg/operator/rgw/rgw.go pkg/operator/rgw/controller_test.go pkg/operator/rgw/rgw_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/operator/test/client.go pkg/operator/test/info.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/rook/client/block.go pkg/rook/client/client.go pkg/rook/client/filesystem.go pkg/rook/client/node.go pkg/rook/client/object.go pkg/rook/client/pool.go pkg/rook/client/status.go pkg/rook/client/client_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/rook/test/mockclient.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/test/info.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/util/file.go pkg/util/set.go pkg/util/set_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/util/display/bytes.go pkg/util/display/numbers.go pkg/util/display/bytes_test.go pkg/util/display/numbers_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/util/exec/error.go pkg/util/exec/exec.go pkg/util/exec/log.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/util/exec/test/mockexec.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/util/flags/flags.go pkg/util/flags/flags_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/util/kvstore/kvstore.go pkg/util/kvstore/mockkvstore.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/util/proc/log.go pkg/util/proc/monitoredproc.go pkg/util/proc/procmanager.go pkg/util/proc/monitoredproc_test.go pkg/util/proc/procmanager_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/util/sys/device.go pkg/util/sys/kmod.go pkg/util/sys/log.go pkg/util/sys/parse.go pkg/util/sys/device_test.go pkg/util/sys/parse_test.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  pkg/version/version.go
/usr/local/go/pkg/tool/linux_amd64/vet -tags  tests/integration/base_block_test.go tests/integration/base_deploy_test.go tests/integration/base_file_test.go tests/integration/base_object_test.go tests/integration/block_createBlock_api_test.go tests/integration/block_createBlock_k8s_test.go tests/integration/helm_test.go tests/integration/mon_test.go tests/integration/multi_cluster_deploy_test.go tests/integration/pool_test.go tests/integration/smoke_test.go
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook'
=== go build linux_amd64
WORK=/tmp/go-build347445997
github.com/rook/rook/cmd/rookctl
mkdir -p $WORK/github.com/rook/rook/cmd/rookctl/_obj/
mkdir -p $WORK/github.com/rook/rook/cmd/rookctl/_obj/exe/
cd /windows10_drive_g/work/src/github.com/rook/rook/cmd/rookctl
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/github.com/rook/rook/cmd/rookctl.a -trimpath $WORK -goversion go1.9.2 -p main -complete -installsuffix static -buildid f6354b09596f83afa94584fa5609c39975903192 -D _/windows10_drive_g/work/src/github.com/rook/rook/cmd/rookctl -I $WORK -I /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -pack ./main.go
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/github.com/rook/rook/cmd/rookctl/_obj/exe/a.out -L $WORK -L /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -installsuffix static -extld=gcc -buildmode=exe -buildid=f6354b09596f83afa94584fa5609c39975903192 -v -n -s -w -X github.com/rook/rook/pkg/version.Version=v0.6.0-86.g1c27ae5.dirty $WORK/github.com/rook/rook/cmd/rookctl.a
# github.com/rook/rook/cmd/rookctl
HEADER = -H4 -T0x401000 -D0x0 -R0x1000
searching for runtime.a in $WORK/runtime.a
searching for runtime.a in /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static/runtime.a
 0.00 deadcode
 0.07 pclntab=2967785 bytes, funcdata total 515575 bytes
 0.08 dodata
 0.11 reloc
 0.15 asmb
 0.15 codeblk
 0.15 rodatblk
 0.16 datblk
 0.16 headr
 0.17 cpu time
172216 symbols
300648 liveness data
mkdir -p /windows10_drive_g/work/src/github.com/rook/rook/_output/bin/linux_amd64/
cp $WORK/github.com/rook/rook/cmd/rookctl/_obj/exe/a.out /windows10_drive_g/work/src/github.com/rook/rook/_output/bin/linux_amd64/rookctl
WORK=/tmp/go-build197029419
github.com/rook/rook/cmd/rook
mkdir -p $WORK/github.com/rook/rook/cmd/rook/_obj/
mkdir -p $WORK/github.com/rook/rook/cmd/rook/_obj/exe/
cd /windows10_drive_g/work/src/github.com/rook/rook/cmd/rook
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/github.com/rook/rook/cmd/rook.a -trimpath $WORK -goversion go1.9.2 -p main -complete -installsuffix static -buildid 611a2e3a7bd98d5a558b67e7225ac4f0693b6e71 -importmap github.com/coreos/pkg/capnslog=github.com/rook/rook/vendor/github.com/coreos/pkg/capnslog -importmap github.com/go-ini/ini=github.com/rook/rook/vendor/github.com/go-ini/ini -importmap github.com/spf13/cobra=github.com/rook/rook/vendor/github.com/spf13/cobra -importmap github.com/spf13/pflag=github.com/rook/rook/vendor/github.com/spf13/pflag -importmap k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset=github.com/rook/rook/vendor/k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset -importmap k8s.io/client-go/kubernetes=github.com/rook/rook/vendor/k8s.io/client-go/kubernetes -importmap k8s.io/client-go/rest=github.com/rook/rook/vendor/k8s.io/client-go/rest -D _/windows10_drive_g/work/src/github.com/rook/rook/cmd/rook -I $WORK -I /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -pack ./agent.go ./api.go ./main.go ./mds.go ./mgr.go ./mon.go ./operator.go ./osd.go ./rgw.go ./version.go
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/github.com/rook/rook/cmd/rook/_obj/exe/a.out -L $WORK -L /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -installsuffix static -extld=gcc -buildmode=exe -buildid=611a2e3a7bd98d5a558b67e7225ac4f0693b6e71 -v -n -s -w -X github.com/rook/rook/pkg/version.Version=v0.6.0-86.g1c27ae5.dirty $WORK/github.com/rook/rook/cmd/rook.a
# github.com/rook/rook/cmd/rook
HEADER = -H4 -T0x401000 -D0x0 -R0x1000
searching for runtime.a in $WORK/runtime.a
searching for runtime.a in /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static/runtime.a
 0.00 deadcode
 0.14 pclntab=6903753 bytes, funcdata total 1210944 bytes
 0.16 dodata
 0.24 reloc
 0.28 asmb
 0.28 codeblk
 0.29 rodatblk
 0.31 datblk
 0.31 headr
 0.32 cpu time
312601 symbols
652820 liveness data
mkdir -p /windows10_drive_g/work/src/github.com/rook/rook/_output/bin/linux_amd64/
cp $WORK/github.com/rook/rook/cmd/rook/_obj/exe/a.out /windows10_drive_g/work/src/github.com/rook/rook/_output/bin/linux_amd64/rook
WORK=/tmp/go-build240532014
github.com/rook/rook/cmd/rookflex
mkdir -p $WORK/github.com/rook/rook/cmd/rookflex/_obj/
mkdir -p $WORK/github.com/rook/rook/cmd/rookflex/_obj/exe/
cd /windows10_drive_g/work/src/github.com/rook/rook/cmd/rookflex
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/github.com/rook/rook/cmd/rookflex.a -trimpath $WORK -goversion go1.9.2 -p main -complete -installsuffix static -buildid d62fd78731a3d1dad081b2264f84b637a8503909 -D _/windows10_drive_g/work/src/github.com/rook/rook/cmd/rookflex -I $WORK -I /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -pack ./main.go
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/github.com/rook/rook/cmd/rookflex/_obj/exe/a.out -L $WORK -L /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -installsuffix static -extld=gcc -buildmode=exe -buildid=d62fd78731a3d1dad081b2264f84b637a8503909 -v -n -s -w -X github.com/rook/rook/pkg/version.Version=v0.6.0-86.g1c27ae5.dirty $WORK/github.com/rook/rook/cmd/rookflex.a
# github.com/rook/rook/cmd/rookflex
HEADER = -H4 -T0x401000 -D0x0 -R0x1000
searching for runtime.a in $WORK/runtime.a
searching for runtime.a in /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static/runtime.a
 0.00 deadcode
 0.13 pclntab=6005816 bytes, funcdata total 1046956 bytes
 0.15 dodata
 0.19 reloc
 0.22 asmb
 0.22 codeblk
 0.23 rodatblk
 0.25 datblk
 0.25 headr
 0.26 cpu time
307474 symbols
552852 liveness data
mkdir -p /windows10_drive_g/work/src/github.com/rook/rook/_output/bin/linux_amd64/
cp $WORK/github.com/rook/rook/cmd/rookflex/_obj/exe/a.out /windows10_drive_g/work/src/github.com/rook/rook/_output/bin/linux_amd64/rookflex
WORK=/tmp/go-build664582842
WORK=/tmp/go-build992226513
mkdir -p $WORK/github.com/rook/rook/tests/integration/_test/github.com/rook/rook/tests/
github.com/rook/rook/tests/integration
mkdir -p $WORK/github.com/rook/rook/tests/integration/_test/_obj_test/
cd /windows10_drive_g/work/src/github.com/rook/rook/tests/integration
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/github.com/rook/rook/tests/integration/_test/github.com/rook/rook/tests/integration.a -trimpath $WORK -goversion go1.9.2 -p github.com/rook/rook/tests/integration -complete -installsuffix static -buildid 759ee2c7d018c8ef170bbfcb8c761af1acc8ee10 -importmap github.com/coreos/pkg/capnslog=github.com/rook/rook/vendor/github.com/coreos/pkg/capnslog -importmap github.com/stretchr/testify/assert=github.com/rook/rook/vendor/github.com/stretchr/testify/assert -importmap github.com/stretchr/testify/require=github.com/rook/rook/vendor/github.com/stretchr/testify/require -importmap github.com/stretchr/testify/suite=github.com/rook/rook/vendor/github.com/stretchr/testify/suite -importmap k8s.io/api/core/v1=github.com/rook/rook/vendor/k8s.io/api/core/v1 -importmap k8s.io/apimachinery/pkg/apis/meta/v1=github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/apis/meta/v1 -D _/windows10_drive_g/work/src/github.com/rook/rook/tests/integration -I $WORK -I /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -pack ./base_block_test.go ./base_deploy_test.go ./base_file_test.go ./base_object_test.go ./block_createBlock_api_test.go ./block_createBlock_k8s_test.go ./helm_test.go ./mon_test.go ./multi_cluster_deploy_test.go ./pool_test.go ./smoke_test.go
testmain
mkdir -p $WORK/github.com/rook/rook/tests/integration/_test/
cd $WORK/github.com/rook/rook/tests/integration/_test
/usr/local/go/pkg/tool/linux_amd64/compile -o ./main.a -trimpath $WORK -goversion go1.9.2 -p main -complete -installsuffix static -D "" -I . -I $WORK -I /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -pack ./_testmain.go
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/github.com/rook/rook/tests/integration/_test/integration.test -L $WORK/github.com/rook/rook/tests/integration/_test -L $WORK -L /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -installsuffix static -extld=gcc -buildmode=exe -v -n -s -w -X github.com/rook/rook/pkg/version.Version=v0.6.0-86.g1c27ae5.dirty $WORK/github.com/rook/rook/tests/integration/_test/main.a
# testmain
HEADER = -H4 -T0x401000 -D0x0 -R0x1000
searching for runtime.a in $WORK/github.com/rook/rook/tests/integration/_test/runtime.a
searching for runtime.a in $WORK/runtime.a
searching for runtime.a in /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static/runtime.a
 0.00 deadcode
 0.39 pclntab=6941320 bytes, funcdata total 1325065 bytes
 0.41 dodata
 0.47 reloc
 0.50 asmb
 0.50 codeblk
 0.51 rodatblk
 0.53 datblk
 0.53 headr
 0.54 cpu time
313415 symbols
786264 liveness data
mkdir -p /windows10_drive_g/work/src/github.com/rook/rook/_output/tests/linux_amd64/
cp $WORK/github.com/rook/rook/tests/integration/_test/integration.test /windows10_drive_g/work/src/github.com/rook/rook/_output/tests/linux_amd64/integration
WORK=/tmp/go-build477059898
WORK=/tmp/go-build025668945
mkdir -p $WORK/github.com/rook/rook/tests/longhaul/_test/github.com/rook/rook/tests/
github.com/rook/rook/tests/longhaul
mkdir -p $WORK/github.com/rook/rook/tests/longhaul/_test/_obj_test/
cd /windows10_drive_g/work/src/github.com/rook/rook/tests/longhaul
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/github.com/rook/rook/tests/longhaul/_test/github.com/rook/rook/tests/longhaul.a -trimpath $WORK -goversion go1.9.2 -p github.com/rook/rook/tests/longhaul -complete -installsuffix static -buildid 759ee2c7d018c8ef170bbfcb8c761af1acc8ee10 -importmap github.com/coreos/pkg/capnslog=github.com/rook/rook/vendor/github.com/coreos/pkg/capnslog -importmap github.com/icrowley/fake=github.com/rook/rook/vendor/github.com/icrowley/fake -importmap github.com/stretchr/testify/assert=github.com/rook/rook/vendor/github.com/stretchr/testify/assert -importmap github.com/stretchr/testify/require=github.com/rook/rook/vendor/github.com/stretchr/testify/require -importmap github.com/stretchr/testify/suite=github.com/rook/rook/vendor/github.com/stretchr/testify/suite -importmap k8s.io/apimachinery/pkg/apis/meta/v1=github.com/rook/rook/vendor/k8s.io/apimachinery/pkg/apis/meta/v1 -D _/windows10_drive_g/work/src/github.com/rook/rook/tests/longhaul -I $WORK -I /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -pack ./base_block_test.go ./base_object_test.go ./block_test.go ./block_with_fencing_test.go ./object_test.go
testmain
mkdir -p $WORK/github.com/rook/rook/tests/longhaul/_test/
cd $WORK/github.com/rook/rook/tests/longhaul/_test
/usr/local/go/pkg/tool/linux_amd64/compile -o ./main.a -trimpath $WORK -goversion go1.9.2 -p main -complete -installsuffix static -D "" -I . -I $WORK -I /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -pack ./_testmain.go
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/github.com/rook/rook/tests/longhaul/_test/longhaul.test -L $WORK/github.com/rook/rook/tests/longhaul/_test -L $WORK -L /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static -installsuffix static -extld=gcc -buildmode=exe -v -n -s -w -X github.com/rook/rook/pkg/version.Version=v0.6.0-86.g1c27ae5.dirty $WORK/github.com/rook/rook/tests/longhaul/_test/main.a
# testmain
HEADER = -H4 -T0x401000 -D0x0 -R0x1000
searching for runtime.a in $WORK/github.com/rook/rook/tests/longhaul/_test/runtime.a
searching for runtime.a in $WORK/runtime.a
searching for runtime.a in /windows10_drive_g/work/src/github.com/rook/rook/.work/pkg/linux_amd64_static/runtime.a
 0.00 deadcode
 0.14 pclntab=6947469 bytes, funcdata total 1320265 bytes
 0.16 dodata
 0.21 reloc
 0.24 asmb
 0.24 codeblk
 0.26 rodatblk
 0.27 datblk
 0.28 headr
 0.29 cpu time
312035 symbols
781240 liveness data
mkdir -p /windows10_drive_g/work/src/github.com/rook/rook/_output/tests/linux_amd64/
cp $WORK/github.com/rook/rook/tests/longhaul/_test/longhaul.test /windows10_drive_g/work/src/github.com/rook/rook/_output/tests/linux_amd64/longhaul
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images'
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/rootfs-builder-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/3 : FROM ubuntu:zesty
Trying to pull repository docker.io/library/ubuntu ... 
sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56: Pulling from docker.io/library/ubuntu
Digest: sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56
Status: Image is up to date for docker.io/ubuntu:zesty
 ---> fe1cc5b91830
Step 2/3 : RUN sed -i -e 's/^deb-src/#deb-src/' /etc/apt/sources.list &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy --no-install-recommends         net-tools         netcat &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -yy --no-install-recommends &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean -y
 ---> Using cache
 ---> 179e64f482d5
Step 3/3 : RUN tar -czf /usr/share/copyrights.tar.xz /usr/share/doc/*/copyright &&     rm -rf         /usr/share/doc         /usr/share/man         /usr/share/info         /usr/share/locale         /var/lib/apt/lists/*         /var/log/*         /var/cache/debconf/*         /etc/systemd         /lib/lsb         /lib/udev         /usr/lib/x86_64-linux-gnu/gconv/IBM*         /usr/lib/x86_64-linux-gnu/gconv/EBC* &&     mkdir -p /usr/share/man/man{1..8}
 ---> Using cache
 ---> 8b040d721516
Successfully built 8b040d721516
make[3]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[3]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/base-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/7 : FROM build-be9394dd/rootfs-amd64-8b040d721516
 ---> 2e3b91214509
Step 2/7 : RUN HOST_IP=$(route -n | awk '/^0.0.0.0/ {print $2}') &&     echo "#!/bin/bash" > /usr/local/bin/apt-ng-host-discover &&     echo "if nc -w1 -z $HOST_IP 3142; then printf http://$HOST_IP:3142; else printf DIRECT; fi" >> /usr/local/bin/apt-ng-host-discover &&     chmod +x /usr/local/bin/apt-ng-host-discover &&     echo 'Acquire::http::ProxyAutoDetect "/usr/local/bin/apt-ng-host-discover";' > /etc/apt/apt.conf.d/30proxy
 ---> Using cache
 ---> e3eb9ca4c9ff
Step 3/7 : ARG ARCH
 ---> Using cache
 ---> cb5e0f44ebbb
Step 4/7 : ADD http://localhost:48080/windows10_drive_f/99-mirror/https0x3A0x2F0x2Fgithub.com0x2Fkrallin0x2Ftini0x2Freleases0x2Fdownload/v0.14.0/tini-${ARCH} /tini
Downloading [==================================================>] 19.89 kB/19.89 kB
 ---> Using cache
 ---> 39a370fa6cb6
Step 5/7 : RUN chmod +x /tini
 ---> Using cache
 ---> 34fd3135dcaf
Step 6/7 : ENTRYPOINT /tini -g --
 ---> Using cache
 ---> 4c72f52634fd
Step 7/7 : CMD /bin/bash
 ---> Using cache
 ---> bdc3253aab81
Successfully built bdc3253aab81
make[3]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== caching image build-be9394dd/base-amd64
=== caching image build-be9394dd/rootfs-builder-amd64
=== caching image build-be9394dd/rootfs-amd64-8b040d721516
make[3]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== docker build build-be9394dd/cross-gnu-amd64
Sending build context to Docker daemon 13.82 kB
Step 1/5 : FROM build-be9394dd/base-amd64
 ---> bdc3253aab81
Step 2/5 : RUN echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty main universe restricted" > /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-backports main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-backports main universe restricted" >> /etc/apt/sources.list &&    dpkg --add-architecture armhf &&     dpkg --add-architecture arm64 &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         build-essential         ca-certificates         ccache         cmake         crossbuild-essential-arm64         crossbuild-essential-armhf         curl         git         jq         patch         yasm         zip &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> b4f4549a106a
Step 3/5 : ADD toolchain /usr/local/toolchain
 ---> Using cache
 ---> 3c2fc8b13df0
Step 4/5 : WORKDIR /build
 ---> Using cache
 ---> 5af7d7101090
Step 5/5 : ENV TERM xterm
 ---> Using cache
 ---> c4a654332a52
Successfully built c4a654332a52
make[3]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== caching image build-be9394dd/cross-gnu-amd64
make[3]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== docker build build-be9394dd/ceph-builder-amd64-tcmalloc-release
Sending build context to Docker daemon 14.85 kB
Step 1/16 : FROM build-be9394dd/cross-gnu-amd64
 ---> c4a654332a52
Step 2/16 : ARG CEPH_GIT_REPO
 ---> Using cache
 ---> d8a28e9b2948
Step 3/16 : RUN git clone ${CEPH_GIT_REPO} ceph
 ---> Using cache
 ---> ce6f3e2a9e94
Step 4/16 : ARG ARCH
 ---> Using cache
 ---> 10b0906f2885
Step 5/16 : ARG CROSS_TRIPLE
 ---> Using cache
 ---> 18fb3b2e6d58
Step 6/16 : RUN DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         python-setuptools         gperf                 cython:${ARCH}         python:${ARCH}         libaio-dev:${ARCH}         libatomic-ops-dev:${ARCH}         libbabeltrace-dev:${ARCH}         libblkid-dev:${ARCH}         libboost-context-dev:${ARCH}         libboost-coroutine-dev:${ARCH}         libboost-date-time-dev:${ARCH}         libboost-filesystem-dev:${ARCH}         libboost-iostreams-dev:${ARCH}         libboost-program-options-dev:${ARCH}         libboost-python-dev:${ARCH}         libboost-random-dev:${ARCH}         libboost-regex-dev:${ARCH}         libboost-system-dev:${ARCH}         libboost-thread-dev:${ARCH}         libcurl4-gnutls-dev:${ARCH}         libexpat1-dev:${ARCH}         libgoogle-perftools-dev:${ARCH}         libgoogle-perftools4:${ARCH}         libibverbs-dev:${ARCH}         libjemalloc-dev:${ARCH}         libkeyutils-dev:${ARCH}         libldap2-dev:${ARCH}         libnss3-dev:${ARCH}         libpython-dev:${ARCH}         libsnappy-dev:${ARCH}         libssl-dev:${ARCH}         libtcmalloc-minimal4:${ARCH}         libudev-dev:${ARCH}         libunwind-dev:${ARCH}         zlib1g-dev:${ARCH} &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> 8dc3989141e0
Step 7/16 : ARG CEPH_GIT_COMMIT
 ---> Using cache
 ---> 9d5d869cd549
Step 8/16 : RUN cd /build/ceph &&     git fetch --all --prune &&     git checkout -b ceph-builder ${CEPH_GIT_COMMIT} &&     git submodule update --init --recursive
 ---> Using cache
 ---> 1c0a3108f9d6
Step 9/16 : RUN mv /build/ceph/CMakeLists.txt /build/ceph/CMakeLists.original.txt
 ---> Using cache
 ---> 70dfa574cf25
Step 10/16 : ADD CMakeLists.txt /build/ceph
 ---> Using cache
 ---> 3b3c3ec08145
Step 11/16 : ARG CEPH_BUILD_TYPE
 ---> Using cache
 ---> 3f5087dd3915
Step 12/16 : ARG CEPH_ALLOCATOR
 ---> Using cache
 ---> 225c822b412c
Step 13/16 : RUN mkdir -p /build/ceph/build && cd /build/ceph/build &&     cmake     -DCMAKE_SKIP_INSTALL_ALL_DEPENDENCY=ON     -DCMAKE_INSTALL_PREFIX=/usr/local     -DCMAKE_BUILD_TYPE=${CEPH_BUILD_TYPE}     -DCMAKE_TOOLCHAIN_FILE=/usr/local/toolchain/${CROSS_TRIPLE}.cmake     -DALLOCATOR=${CEPH_ALLOCATOR}     -DWITH_SYSTEM_BOOST=ON     -DWITH_EMBEDDED=OFF     -DWITH_FUSE=OFF     -DWITH_LEVELDB=OFF     -DWITH_LTTNG=OFF     -DWITH_MANPAGE=OFF     -DWITH_PROFILER=OFF     -DWITH_PYTHON3=OFF     -DWITH_RADOSGW_FCGI_FRONTEND=OFF     ..
 ---> Using cache
 ---> 95367828ae35
Step 14/16 : WORKDIR /build/ceph/build
 ---> Using cache
 ---> 8c46aca7ea3c
Step 15/16 : ADD build-ceph.sh /build/
 ---> Using cache
 ---> d9e5a61fdc5c
Step 16/16 : CMD /build/build-ceph.sh
 ---> Using cache
 ---> c19572dcbeb0
Successfully built c19572dcbeb0
make[3]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[4]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== building ceph from build-be9394dd/ceph-builder-amd64-tcmalloc-release
Scanning dependencies of target civetweb_h
[  0%] keep civetweb.h up-to-date
[  0%] Built target civetweb_h
Scanning dependencies of target common_utf8
[  0%] Building C object src/CMakeFiles/common_utf8.dir/common/utf8.c.o
ccache: error: Failed to create directory /root/.ccache/tmp: Permission denied
src/CMakeFiles/common_utf8.dir/build.make:62: recipe for target 'src/CMakeFiles/common_utf8.dir/common/utf8.c.o' failed
make[3]: *** [src/CMakeFiles/common_utf8.dir/common/utf8.c.o] Error 1
CMakeFiles/Makefile2:640: recipe for target 'src/CMakeFiles/common_utf8.dir/all' failed
make[2]: *** [src/CMakeFiles/common_utf8.dir/all] Error 2
CMakeFiles/Makefile2:374: recipe for target 'CMakeFiles/rook.dir/rule' failed
make[1]: *** [CMakeFiles/rook.dir/rule] Error 2
Makefile:225: recipe for target 'rook' failed
make: *** [rook] Error 2
make[4]: *** [Makefile:76: ceph] Error 2
make[4]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[3]: *** [../image.mk:134: cache.lookup] Error 2
make[3]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[2]: *** [Makefile:63: do.build] Error 2
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[1]: *** [Makefile:30: ceph.linux_amd64] Error 2
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images'
make: *** [Makefile:114: build] Error 2
```

```
[vagrant@localhost rook]$ sudo setenforce 0
[vagrant@localhost rook]$ getenforce
Permissive
[vagrant@localhost rook]$ sudo sed -i 's/SELINUX=enforcing/SELINUX=permissive/' /etc/sysconfig/selinux 
```

Build in docker require more resources, change VM memory size 8192, 2 cores
```
[vagrant@localhost images]$ make V=1 build
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/rootfs-builder-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/3 : FROM ubuntu:zesty
Trying to pull repository docker.io/library/ubuntu ... 
sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56: Pulling from docker.io/library/ubuntu
Digest: sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56
Status: Image is up to date for docker.io/ubuntu:zesty
 ---> fe1cc5b91830
Step 2/3 : RUN sed -i -e 's/^deb-src/#deb-src/' /etc/apt/sources.list &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy --no-install-recommends         net-tools         netcat &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -yy --no-install-recommends &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean -y
 ---> Using cache
 ---> 179e64f482d5
Step 3/3 : RUN tar -czf /usr/share/copyrights.tar.xz /usr/share/doc/*/copyright &&     rm -rf         /usr/share/doc         /usr/share/man         /usr/share/info         /usr/share/locale         /var/lib/apt/lists/*         /var/log/*         /var/cache/debconf/*         /etc/systemd         /lib/lsb         /lib/udev         /usr/lib/x86_64-linux-gnu/gconv/IBM*         /usr/lib/x86_64-linux-gnu/gconv/EBC* &&     mkdir -p /usr/share/man/man{1..8}
 ---> Using cache
 ---> 8b040d721516
Successfully built 8b040d721516
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/base-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/7 : FROM build-be9394dd/rootfs-amd64-8b040d721516
 ---> 2e3b91214509
Step 2/7 : RUN HOST_IP=$(route -n | awk '/^0.0.0.0/ {print $2}') &&     echo "#!/bin/bash" > /usr/local/bin/apt-ng-host-discover &&     echo "if nc -w1 -z $HOST_IP 3142; then printf http://$HOST_IP:3142; else printf DIRECT; fi" >> /usr/local/bin/apt-ng-host-discover &&     chmod +x /usr/local/bin/apt-ng-host-discover &&     echo 'Acquire::http::ProxyAutoDetect "/usr/local/bin/apt-ng-host-discover";' > /etc/apt/apt.conf.d/30proxy
 ---> Using cache
 ---> e3eb9ca4c9ff
Step 3/7 : ARG ARCH
 ---> Using cache
 ---> cb5e0f44ebbb
Step 4/7 : ADD http://localhost:48080/windows10_drive_f/99-mirror/https0x3A0x2F0x2Fgithub.com0x2Fkrallin0x2Ftini0x2Freleases0x2Fdownload/v0.14.0/tini-${ARCH} /tini
Downloading [==================================================>] 19.89 kB/19.89 kB
 ---> Using cache
 ---> 39a370fa6cb6
Step 5/7 : RUN chmod +x /tini
 ---> Using cache
 ---> 34fd3135dcaf
Step 6/7 : ENTRYPOINT /tini -g --
 ---> Using cache
 ---> 4c72f52634fd
Step 7/7 : CMD /bin/bash
 ---> Using cache
 ---> bdc3253aab81
Successfully built bdc3253aab81
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== caching image build-be9394dd/base-amd64
=== caching image build-be9394dd/rootfs-builder-amd64
=== caching image build-be9394dd/rootfs-amd64-8b040d721516
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== docker build build-be9394dd/cross-gnu-amd64
Sending build context to Docker daemon 13.82 kB
Step 1/5 : FROM build-be9394dd/base-amd64
 ---> bdc3253aab81
Step 2/5 : RUN echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty main universe restricted" > /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-backports main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-backports main universe restricted" >> /etc/apt/sources.list &&    dpkg --add-architecture armhf &&     dpkg --add-architecture arm64 &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         build-essential         ca-certificates         ccache         cmake         crossbuild-essential-arm64         crossbuild-essential-armhf         curl         git         jq         patch         yasm         zip &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> b4f4549a106a
Step 3/5 : ADD toolchain /usr/local/toolchain
 ---> Using cache
 ---> 3c2fc8b13df0
Step 4/5 : WORKDIR /build
 ---> Using cache
 ---> 5af7d7101090
Step 5/5 : ENV TERM xterm
 ---> Using cache
 ---> c4a654332a52
Successfully built c4a654332a52
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== caching image build-be9394dd/cross-gnu-amd64
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== docker build build-be9394dd/ceph-builder-amd64-tcmalloc-release
Sending build context to Docker daemon 14.85 kB
Step 1/16 : FROM build-be9394dd/cross-gnu-amd64
 ---> c4a654332a52
Step 2/16 : ARG CEPH_GIT_REPO
 ---> Using cache
 ---> d8a28e9b2948
Step 3/16 : RUN git clone ${CEPH_GIT_REPO} ceph
 ---> Using cache
 ---> ce6f3e2a9e94
Step 4/16 : ARG ARCH
 ---> Using cache
 ---> 10b0906f2885
Step 5/16 : ARG CROSS_TRIPLE
 ---> Using cache
 ---> 18fb3b2e6d58
Step 6/16 : RUN DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         python-setuptools         gperf                 cython:${ARCH}         python:${ARCH}         libaio-dev:${ARCH}         libatomic-ops-dev:${ARCH}         libbabeltrace-dev:${ARCH}         libblkid-dev:${ARCH}         libboost-context-dev:${ARCH}         libboost-coroutine-dev:${ARCH}         libboost-date-time-dev:${ARCH}         libboost-filesystem-dev:${ARCH}         libboost-iostreams-dev:${ARCH}         libboost-program-options-dev:${ARCH}         libboost-python-dev:${ARCH}         libboost-random-dev:${ARCH}         libboost-regex-dev:${ARCH}         libboost-system-dev:${ARCH}         libboost-thread-dev:${ARCH}         libcurl4-gnutls-dev:${ARCH}         libexpat1-dev:${ARCH}         libgoogle-perftools-dev:${ARCH}         libgoogle-perftools4:${ARCH}         libibverbs-dev:${ARCH}         libjemalloc-dev:${ARCH}         libkeyutils-dev:${ARCH}         libldap2-dev:${ARCH}         libnss3-dev:${ARCH}         libpython-dev:${ARCH}         libsnappy-dev:${ARCH}         libssl-dev:${ARCH}         libtcmalloc-minimal4:${ARCH}         libudev-dev:${ARCH}         libunwind-dev:${ARCH}         zlib1g-dev:${ARCH} &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> 8dc3989141e0
Step 7/16 : ARG CEPH_GIT_COMMIT
 ---> Using cache
 ---> 9d5d869cd549
Step 8/16 : RUN cd /build/ceph &&     git fetch --all --prune &&     git checkout -b ceph-builder ${CEPH_GIT_COMMIT} &&     git submodule update --init --recursive
 ---> Using cache
 ---> 1c0a3108f9d6
Step 9/16 : RUN mv /build/ceph/CMakeLists.txt /build/ceph/CMakeLists.original.txt
 ---> Using cache
 ---> 70dfa574cf25
Step 10/16 : ADD CMakeLists.txt /build/ceph
 ---> Using cache
 ---> 3b3c3ec08145
Step 11/16 : ARG CEPH_BUILD_TYPE
 ---> Using cache
 ---> 3f5087dd3915
Step 12/16 : ARG CEPH_ALLOCATOR
 ---> Using cache
 ---> 225c822b412c
Step 13/16 : RUN mkdir -p /build/ceph/build && cd /build/ceph/build &&     cmake     -DCMAKE_SKIP_INSTALL_ALL_DEPENDENCY=ON     -DCMAKE_INSTALL_PREFIX=/usr/local     -DCMAKE_BUILD_TYPE=${CEPH_BUILD_TYPE}     -DCMAKE_TOOLCHAIN_FILE=/usr/local/toolchain/${CROSS_TRIPLE}.cmake     -DALLOCATOR=${CEPH_ALLOCATOR}     -DWITH_SYSTEM_BOOST=ON     -DWITH_EMBEDDED=OFF     -DWITH_FUSE=OFF     -DWITH_LEVELDB=OFF     -DWITH_LTTNG=OFF     -DWITH_MANPAGE=OFF     -DWITH_PROFILER=OFF     -DWITH_PYTHON3=OFF     -DWITH_RADOSGW_FCGI_FRONTEND=OFF     ..
 ---> Using cache
 ---> 95367828ae35
Step 14/16 : WORKDIR /build/ceph/build
 ---> Using cache
 ---> 8c46aca7ea3c
Step 15/16 : ADD build-ceph.sh /build/
 ---> Using cache
 ---> d9e5a61fdc5c
Step 16/16 : CMD /build/build-ceph.sh
 ---> Using cache
 ---> c19572dcbeb0
Successfully built c19572dcbeb0
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[3]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== building ceph from build-be9394dd/ceph-builder-amd64-tcmalloc-release
Scanning dependencies of target civetweb_h
[  0%] keep civetweb.h up-to-date
[  0%] Built target civetweb_h
Scanning dependencies of target common_utf8
[  0%] Building C object src/CMakeFiles/common_utf8.dir/common/utf8.c.o
[  0%] Linking C static library ../lib/libcommon_utf8.a
[  0%] Built target common_utf8
Scanning dependencies of target compressor_objs
[  0%] Building CXX object src/compressor/CMakeFiles/compressor_objs.dir/Compressor.cc.o
[  1%] Building CXX object src/compressor/CMakeFiles/compressor_objs.dir/AsyncCompressor.cc.o
[  1%] Built target compressor_objs
Scanning dependencies of target common-objs
[  1%] Building C object src/CMakeFiles/common-objs.dir/ceph_ver.c.o
[  1%] Building CXX object src/CMakeFiles/common-objs.dir/common/AsyncOpTracker.cc.o
[  1%] Building CXX object src/CMakeFiles/common-objs.dir/common/DecayCounter.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/LogClient.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/LogEntry.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/PrebufferedStreambuf.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/BackTrace.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/perf_counters.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/perf_histogram.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/mutex_debug.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/Mutex.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/OutputDataSocket.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/admin_socket.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/admin_socket_client.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/bloom_filter.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/Readahead.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/cmdparse.cc.o
[  3%] Building C object src/CMakeFiles/common-objs.dir/common/escape.c.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/url_escape.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/io_priority.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/Clock.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_time.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/mempool.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/Throttle.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/Timer.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/Finisher.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/environment.cc.o
[  5%] Building C object src/CMakeFiles/common-objs.dir/common/sctp_crc32.c.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/crc32c.cc.o
[  5%] Building C object src/CMakeFiles/common-objs.dir/common/crc32c_intel_baseline.c.o
[  5%] Building C object src/CMakeFiles/common-objs.dir/xxHash/xxhash.c.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/assert.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/run_cmd.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/WorkQueue.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/ConfUtils.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/MemoryModel.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/fd.cc.o
[  5%] Building C object src/CMakeFiles/common-objs.dir/common/xattr.c.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/str_list.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/str_map.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/snap_types.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/errno.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/TrackedOp.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/SloppyCRCMap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/types.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/iso_8601.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/log/Log.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/log/SubsystemMap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mon/MonCap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mon/MonClient.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mon/MonMap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mgr/MgrClient.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/msg/simple/Accepter.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/msg/DispatchQueue.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/msg/Message.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mon/PGMap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mgr/ServiceMap.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/osd/ECMsgTypes.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/osd/HitSet.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/common/RefCountedObj.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/Messenger.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/simple/Pipe.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/simple/PipeConnection.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/simple/SimpleMessenger.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/AsyncConnection.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/AsyncMessenger.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/Event.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/EventSelect.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/Stack.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/PosixStack.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/net_handler.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/QueueStrategy.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/rdma/Infiniband.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/rdma/RDMAConnectedSocketImpl.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/rdma/RDMAServerSocketImpl.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/rdma/RDMAStack.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/msg/msg_types.cc.o
[ 11%] Building C object src/CMakeFiles/common-objs.dir/common/reverse.c.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/hobject.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/osd/OSDMap.cc.o
/build/ceph/src/osd/OSDMap.cc: In member function 'int OSDMap::validate_crush_rules(CrushWrapper*, std::ostream*) const':
/build/ceph/src/osd/OSDMap.cc:3317:25: warning: comparison between signed and unsigned integer expressions [-Wsign-compare]
     if (pool.get_size() < (int)newcrush->get_rule_mask_min_size(ruleno) ||
         ~~~~~~~~~~~~~~~~^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
/build/ceph/src/osd/OSDMap.cc:3318:18: warning: comparison between signed and unsigned integer expressions [-Wsign-compare]
  pool.get_size() > (int)newcrush->get_rule_mask_max_size(ruleno)) {
  ~~~~~~~~~~~~~~~~^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/osd/OSDMapMapping.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/histogram.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/osd/osd_types.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/osd/OpRequest.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/blkdev.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/common_init.cc.o
[ 11%] Building C object src/CMakeFiles/common-objs.dir/common/pipe.c.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_argparse.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_context.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/code_environment.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/dout.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/signal.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/Thread.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/Formatter.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/HTMLFormatter.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/HeartbeatMap.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/PluginRegistry.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_fs.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_hash.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_strings.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_frag.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/options.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/config.cc.o
[ 13%] Building C object src/CMakeFiles/common-objs.dir/common/utf8.c.o
[ 13%] Building C object src/CMakeFiles/common-objs.dir/common/mime.c.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/strtol.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/page.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/lockdep.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/version.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/hex.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/entity_name.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_crypto.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_crypto_cms.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_json.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/ipaddr.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/pick_address.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/address_helper.cc.o
[ 15%] Building C object src/CMakeFiles/common-objs.dir/common/linux_version.c.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/TracepointProvider.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/Cycles.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/scrub_types.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/bit_str.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/osdc/Striper.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/osdc/Objecter.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/Graylog.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/fs_types.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/dns_resolve.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/hostname.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/common/util.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/arch/probe.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/AuthAuthorizeHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/AuthClientHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/AuthSessionHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/AuthMethodList.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/cephx/CephxAuthorizeHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/cephx/CephxClientHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/cephx/CephxProtocol.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/cephx/CephxSessionHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/none/AuthNoneAuthorizeHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/unknown/AuthUnknownAuthorizeHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/Crypto.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/KeyRing.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/RotatingKeyRing.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/mds/MDSMap.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/mds/FSMap.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/mds/FSMapUser.cc.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/mds/inode_backtrace.cc.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/mds/mdstypes.cc.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/mds/flock.cc.o
[ 18%] Building C object src/CMakeFiles/common-objs.dir/arch/intel.c.o
[ 18%] Building C object src/CMakeFiles/common-objs.dir/common/crc32c_intel_fast.c.o
[ 18%] Building ASM object src/CMakeFiles/common-objs.dir/common/crc32c_intel_fast_asm.s.o
[ 18%] Building ASM object src/CMakeFiles/common-objs.dir/common/crc32c_intel_fast_zero_asm.s.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/EventEpoll.cc.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/perfglue/disabled_stubs.cc.o
[ 18%] Built target common-objs
Scanning dependencies of target crush_objs
[ 18%] Building C object src/CMakeFiles/crush_objs.dir/crush/builder.c.o
[ 18%] Building C object src/CMakeFiles/crush_objs.dir/crush/mapper.c.o
[ 18%] Building C object src/CMakeFiles/crush_objs.dir/crush/crush.c.o
[ 18%] Building C object src/CMakeFiles/crush_objs.dir/crush/hash.c.o
[ 18%] Building CXX object src/CMakeFiles/crush_objs.dir/crush/CrushWrapper.cc.o
[ 18%] Building CXX object src/CMakeFiles/crush_objs.dir/crush/CrushCompiler.cc.o
[ 18%] Building CXX object src/CMakeFiles/crush_objs.dir/crush/CrushTester.cc.o
[ 18%] Building CXX object src/CMakeFiles/crush_objs.dir/crush/CrushLocation.cc.o
[ 18%] Built target crush_objs
Scanning dependencies of target common_mountcephfs_objs
[ 18%] Building C object src/CMakeFiles/common_mountcephfs_objs.dir/common/armor.c.o
[ 18%] Building C object src/CMakeFiles/common_mountcephfs_objs.dir/common/safe_io.c.o
[ 18%] Building C object src/CMakeFiles/common_mountcephfs_objs.dir/common/module.c.o
[ 18%] Building C object src/CMakeFiles/common_mountcephfs_objs.dir/common/addr_parsing.c.o
[ 18%] Built target common_mountcephfs_objs
Scanning dependencies of target common_buffer_obj
[ 18%] Building CXX object src/CMakeFiles/common_buffer_obj.dir/common/buffer.cc.o
[ 18%] Built target common_buffer_obj
Scanning dependencies of target common_texttable_obj
[ 18%] Building CXX object src/CMakeFiles/common_texttable_obj.dir/common/TextTable.cc.o
[ 18%] Built target common_texttable_obj
Scanning dependencies of target json_spirit
[ 18%] Building CXX object src/json_spirit/CMakeFiles/json_spirit.dir/json_spirit_reader.cpp.o
[ 18%] Building CXX object src/json_spirit/CMakeFiles/json_spirit.dir/json_spirit_writer.cpp.o
[ 18%] Linking CXX static library ../../lib/libjson_spirit.a
[ 18%] Built target json_spirit
Scanning dependencies of target global_common_objs
[ 18%] Building CXX object src/global/CMakeFiles/global_common_objs.dir/global_context.cc.o
[ 18%] Built target global_common_objs
Scanning dependencies of target erasure_code
[ 18%] Building CXX object src/erasure-code/CMakeFiles/erasure_code.dir/ErasureCodePlugin.cc.o
[ 18%] Linking CXX static library ../../lib/liberasure_code.a
[ 18%] Built target erasure_code
Scanning dependencies of target ceph-common
[ 18%] Linking CXX shared library ../lib/libceph-common.so
[ 18%] Built target ceph-common
Scanning dependencies of target libglobal_objs
[ 18%] Building CXX object src/global/CMakeFiles/libglobal_objs.dir/global_init.cc.o
[ 18%] Building CXX object src/global/CMakeFiles/libglobal_objs.dir/pidfile.cc.o
[ 18%] Building CXX object src/global/CMakeFiles/libglobal_objs.dir/signal_handler.cc.o
[ 18%] Built target libglobal_objs
Scanning dependencies of target global
[ 18%] Linking CXX static library ../../lib/libglobal.a
[ 18%] Built target global
Scanning dependencies of target cls_rgw_client
[ 18%] Building CXX object src/cls/CMakeFiles/cls_rgw_client.dir/rgw/cls_rgw_client.cc.o
[ 18%] Building CXX object src/cls/CMakeFiles/cls_rgw_client.dir/rgw/cls_rgw_types.cc.o
[ 18%] Building CXX object src/cls/CMakeFiles/cls_rgw_client.dir/rgw/cls_rgw_ops.cc.o
[ 18%] Linking CXX static library ../../lib/libcls_rgw_client.a
[ 18%] Built target cls_rgw_client
Scanning dependencies of target cls_lock_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_lock_client.dir/lock/cls_lock_client.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_lock_client.dir/lock/cls_lock_types.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_lock_client.dir/lock/cls_lock_ops.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_lock_client.a
[ 20%] Built target cls_lock_client
Scanning dependencies of target cls_statelog_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_statelog_client.dir/statelog/cls_statelog_client.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_statelog_client.a
[ 20%] Built target cls_statelog_client
Scanning dependencies of target cls_log_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_log_client.dir/log/cls_log_client.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_log_client.a
[ 20%] Built target cls_log_client
Scanning dependencies of target cls_refcount_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_refcount_client.dir/refcount/cls_refcount_client.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_refcount_client.dir/refcount/cls_refcount_ops.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_refcount_client.a
[ 20%] Built target cls_refcount_client
Scanning dependencies of target cls_timeindex_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_timeindex_client.dir/timeindex/cls_timeindex_client.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_timeindex_client.a
[ 20%] Built target cls_timeindex_client
Scanning dependencies of target cls_version_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_version_client.dir/version/cls_version_client.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_version_client.dir/version/cls_version_types.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_version_client.a
[ 20%] Built target cls_version_client
Scanning dependencies of target cls_replica_log_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_replica_log_client.dir/replica_log/cls_replica_log_types.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_replica_log_client.dir/replica_log/cls_replica_log_ops.cc.o
[ 22%] Building CXX object src/cls/CMakeFiles/cls_replica_log_client.dir/replica_log/cls_replica_log_client.cc.o
[ 22%] Linking CXX static library ../../lib/libcls_replica_log_client.a
[ 22%] Built target cls_replica_log_client
Scanning dependencies of target cls_user_client
[ 22%] Building CXX object src/cls/CMakeFiles/cls_user_client.dir/user/cls_user_client.cc.o
[ 22%] Building CXX object src/cls/CMakeFiles/cls_user_client.dir/user/cls_user_types.cc.o
[ 22%] Building CXX object src/cls/CMakeFiles/cls_user_client.dir/user/cls_user_ops.cc.o
[ 22%] Linking CXX static library ../../lib/libcls_user_client.a
[ 22%] Built target cls_user_client
Scanning dependencies of target osdc
[ 22%] Building CXX object src/osdc/CMakeFiles/osdc.dir/Filer.cc.o
[ 22%] Building CXX object src/osdc/CMakeFiles/osdc.dir/ObjectCacher.cc.o
[ 22%] Building CXX object src/osdc/CMakeFiles/osdc.dir/Objecter.cc.o
[ 22%] Building CXX object src/osdc/CMakeFiles/osdc.dir/Striper.cc.o
[ 22%] Linking CXX static library ../../lib/libosdc.a
[ 22%] Built target osdc
Scanning dependencies of target librados_api_obj
[ 22%] Building CXX object src/librados/CMakeFiles/librados_api_obj.dir/librados.cc.o
[ 22%] Built target librados_api_obj
Scanning dependencies of target librados_objs
[ 22%] Building CXX object src/librados/CMakeFiles/librados_objs.dir/IoCtxImpl.cc.o
[ 22%] Building CXX object src/librados/CMakeFiles/librados_objs.dir/RadosXattrIter.cc.o
[ 22%] Building CXX object src/librados/CMakeFiles/librados_objs.dir/RadosClient.cc.o
[ 22%] Built target librados_objs
Scanning dependencies of target librados
[ 22%] Linking CXX shared library ../../lib/librados.so
[ 22%] Built target librados
[ 22%] Generate rgw_iam_policy_keywords.frag.cc
Scanning dependencies of target rgw_a
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_acl.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_acl_s3.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_acl_swift.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_auth.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_auth_keystone.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_auth_s3.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_basic_types.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_bucket.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_cache.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_client_io.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_common.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_compression.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_cors.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_cors_s3.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_dencoder.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_env.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_es_query.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_formats.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_frontend.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_gc.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_http_client.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_json_enc.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_keystone.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_ldap.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_loadgen.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_log.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_lc.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_lc_s3.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_metadata.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_multi.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_multi_del.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_data_sync.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync_module.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync_module_es.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync_module_es_rest.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync_module_log.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_period_history.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_period_puller.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_period_pusher.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_realm_reloader.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_realm_watcher.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_reshard.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_coroutine.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_cr_rados.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_object_expirer_core.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_op.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_os_lib.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_policy_s3.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_process.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_quota.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rados.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_replica_log.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_request.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_resolve.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_bucket.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_client.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_config.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_conn.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_log.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_metadata.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_opstate.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_realm.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_replica_log.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_role.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_s3.cc.o
/build/ceph/src/rgw/rgw_rest_s3.cc:3333:13: warning: 'void init_anon_user(req_state*)' defined but not used [-Wunused-function]
 static void init_anon_user(struct req_state *s)
             ^~~~~~~~~~~~~~
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_swift.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_usage.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_user.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_role.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_string.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_swift_auth.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_tag.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_tag_s3.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_tools.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_usage.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_user.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_website.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_xml.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_xml_enc.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_torrent.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_crypt.cc.o
/build/ceph/src/rgw/rgw_crypt.cc:38:2: warning: #warning "TODO: move this code to auth/Crypto for others to reuse." [-Wcpp]
 #warning "TODO: move this code to auth/Crypto for others to reuse."
  ^~~~~~~
/build/ceph/src/rgw/rgw_crypt.cc:247:2: warning: #warning "TODO: use auth/Crypto instead of reimplementing." [-Wcpp]
 #warning "TODO: use auth/Crypto instead of reimplementing."
  ^~~~~~~
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_crypt_sanitize.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_iam_policy.cc.o
[ 32%] Linking CXX static library ../../lib/librgw_a.a
[ 32%] Built target rgw_a
Scanning dependencies of target radosgw-admin
[ 32%] Building CXX object src/rgw/CMakeFiles/radosgw-admin.dir/rgw_admin.cc.o
[ 32%] Building CXX object src/rgw/CMakeFiles/radosgw-admin.dir/rgw_orphan.cc.o
[ 32%] Linking CXX executable ../../bin/radosgw-admin
[ 32%] Built target radosgw-admin
Scanning dependencies of target common
[ 32%] Linking CXX static library ../lib/libcommon.a
[ 32%] Built target common
Scanning dependencies of target global-static
[ 32%] Linking CXX static library ../../lib/libglobal-static.a
[ 32%] Built target global-static
Scanning dependencies of target liblua
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lapi.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lcode.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lctype.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ldebug.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ldo.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ldump.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lfunc.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lgc.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/llex.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lmem.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lobject.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lopcodes.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lparser.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lstate.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lstring.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ltable.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ltm.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lundump.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lvm.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lzio.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lauxlib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lbaselib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lbitlib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lcorolib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/ldblib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/liolib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lmathlib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/loslib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lstrlib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/ltablib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/linit.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lutf8lib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/loadlib_rel.c.o
[ 35%] Linking C static library ../../lib/liblua.a
[ 35%] Built target liblua
Scanning dependencies of target heap_profiler_objs
[ 35%] Building CXX object src/CMakeFiles/heap_profiler_objs.dir/perfglue/heap_profiler.cc.o
[ 35%] Built target heap_profiler_objs
Scanning dependencies of target mds
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/Capability.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSDaemon.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSRank.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/Beacon.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/flock.cc.o
[ 35%] Building C object src/mds/CMakeFiles/mds.dir/locks.c.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/journal.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/Server.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/Mutation.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/MDCache.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/RecoveryQueue.cc.o
In file included from /usr/include/boost/throw_exception.hpp:42:0,
                 from /usr/include/boost/optional/optional.hpp:30,
                 from /usr/include/boost/optional.hpp:15,
                 from /build/ceph/src/include/denc.h:38,
                 from /build/ceph/src/include/encoding.h:25,
                 from /build/ceph/src/include/uuid.h:8,
                 from /build/ceph/src/include/types.h:21,
                 from /build/ceph/src/msg/msg_types.h:21,
                 from /build/ceph/src/common/entity_name.h:20,
                 from /build/ceph/src/common/config.h:19,
                 from /build/ceph/src/mds/CInode.h:20,
                 from /build/ceph/src/mds/RecoveryQueue.cc:15:
/usr/include/boost/exception/exception.hpp:479:13: internal compiler error: Segmentation fault
             ~clone_impl() throw()
             ^
Please submit a full bug report,
with preprocessed source if appropriate.
See <file:///usr/share/doc/gcc-6/README.Bugs> for instructions.
src/mds/CMakeFiles/mds.dir/build.make:302: recipe for target 'src/mds/CMakeFiles/mds.dir/RecoveryQueue.cc.o' failed
make[3]: *** [src/mds/CMakeFiles/mds.dir/RecoveryQueue.cc.o] Error 1
CMakeFiles/Makefile2:4356: recipe for target 'src/mds/CMakeFiles/mds.dir/all' failed
make[2]: *** [src/mds/CMakeFiles/mds.dir/all] Error 2
CMakeFiles/Makefile2:374: recipe for target 'CMakeFiles/rook.dir/rule' failed
make[1]: *** [CMakeFiles/rook.dir/rule] Error 2
Makefile:225: recipe for target 'rook' failed
make: *** [rook] Error 2
make[3]: *** [Makefile:76: ceph] Error 2
make[3]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[2]: *** [../image.mk:134: cache.lookup] Error 2
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[1]: *** [Makefile:63: do.build] Error 2
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make: *** [Makefile:30: ceph.linux_amd64] Error 2
```

Success
```
[vagrant@localhost images]$ make V=1 ceph.linux_amd64
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/rootfs-builder-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/3 : FROM ubuntu:zesty
Trying to pull repository docker.io/library/ubuntu ... 
sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56: Pulling from docker.io/library/ubuntu
Digest: sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56
Status: Image is up to date for docker.io/ubuntu:zesty
 ---> fe1cc5b91830
Step 2/3 : RUN sed -i -e 's/^deb-src/#deb-src/' /etc/apt/sources.list &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy --no-install-recommends         net-tools         netcat &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -yy --no-install-recommends &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean -y
 ---> Using cache
 ---> 179e64f482d5
Step 3/3 : RUN tar -czf /usr/share/copyrights.tar.xz /usr/share/doc/*/copyright &&     rm -rf         /usr/share/doc         /usr/share/man         /usr/share/info         /usr/share/locale         /var/lib/apt/lists/*         /var/log/*         /var/cache/debconf/*         /etc/systemd         /lib/lsb         /lib/udev         /usr/lib/x86_64-linux-gnu/gconv/IBM*         /usr/lib/x86_64-linux-gnu/gconv/EBC* &&     mkdir -p /usr/share/man/man{1..8}
 ---> Using cache
 ---> 8b040d721516
Successfully built 8b040d721516
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/base-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/7 : FROM build-be9394dd/rootfs-amd64-8b040d721516
 ---> 2e3b91214509
Step 2/7 : RUN HOST_IP=$(route -n | awk '/^0.0.0.0/ {print $2}') &&     echo "#!/bin/bash" > /usr/local/bin/apt-ng-host-discover &&     echo "if nc -w1 -z $HOST_IP 3142; then printf http://$HOST_IP:3142; else printf DIRECT; fi" >> /usr/local/bin/apt-ng-host-discover &&     chmod +x /usr/local/bin/apt-ng-host-discover &&     echo 'Acquire::http::ProxyAutoDetect "/usr/local/bin/apt-ng-host-discover";' > /etc/apt/apt.conf.d/30proxy
 ---> Using cache
 ---> e3eb9ca4c9ff
Step 3/7 : ARG ARCH
 ---> Using cache
 ---> cb5e0f44ebbb
Step 4/7 : ADD http://localhost:48080/windows10_drive_f/99-mirror/https0x3A0x2F0x2Fgithub.com0x2Fkrallin0x2Ftini0x2Freleases0x2Fdownload/v0.14.0/tini-${ARCH} /tini
Downloading [==================================================>] 19.89 kB/19.89 kB
 ---> Using cache
 ---> 39a370fa6cb6
Step 5/7 : RUN chmod +x /tini
 ---> Using cache
 ---> 34fd3135dcaf
Step 6/7 : ENTRYPOINT /tini -g --
 ---> Using cache
 ---> 4c72f52634fd
Step 7/7 : CMD /bin/bash
 ---> Using cache
 ---> bdc3253aab81
Successfully built bdc3253aab81
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== caching image build-be9394dd/base-amd64
=== caching image build-be9394dd/rootfs-builder-amd64
=== caching image build-be9394dd/rootfs-amd64-8b040d721516
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== docker build build-be9394dd/cross-gnu-amd64
Sending build context to Docker daemon 13.82 kB
Step 1/5 : FROM build-be9394dd/base-amd64
 ---> bdc3253aab81
Step 2/5 : RUN echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty main universe restricted" > /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-backports main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-backports main universe restricted" >> /etc/apt/sources.list &&    dpkg --add-architecture armhf &&     dpkg --add-architecture arm64 &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         build-essential         ca-certificates         ccache         cmake         crossbuild-essential-arm64         crossbuild-essential-armhf         curl         git         jq         patch         yasm         zip &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> b4f4549a106a
Step 3/5 : ADD toolchain /usr/local/toolchain
 ---> Using cache
 ---> 3c2fc8b13df0
Step 4/5 : WORKDIR /build
 ---> Using cache
 ---> 5af7d7101090
Step 5/5 : ENV TERM xterm
 ---> Using cache
 ---> c4a654332a52
Successfully built c4a654332a52
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== caching image build-be9394dd/cross-gnu-amd64
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== docker build build-be9394dd/ceph-builder-amd64-tcmalloc-release
Sending build context to Docker daemon 14.85 kB
Step 1/16 : FROM build-be9394dd/cross-gnu-amd64
 ---> c4a654332a52
Step 2/16 : ARG CEPH_GIT_REPO
 ---> Using cache
 ---> d8a28e9b2948
Step 3/16 : RUN git clone ${CEPH_GIT_REPO} ceph
 ---> Using cache
 ---> ce6f3e2a9e94
Step 4/16 : ARG ARCH
 ---> Using cache
 ---> 10b0906f2885
Step 5/16 : ARG CROSS_TRIPLE
 ---> Using cache
 ---> 18fb3b2e6d58
Step 6/16 : RUN DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         python-setuptools         gperf                 cython:${ARCH}         python:${ARCH}         libaio-dev:${ARCH}         libatomic-ops-dev:${ARCH}         libbabeltrace-dev:${ARCH}         libblkid-dev:${ARCH}         libboost-context-dev:${ARCH}         libboost-coroutine-dev:${ARCH}         libboost-date-time-dev:${ARCH}         libboost-filesystem-dev:${ARCH}         libboost-iostreams-dev:${ARCH}         libboost-program-options-dev:${ARCH}         libboost-python-dev:${ARCH}         libboost-random-dev:${ARCH}         libboost-regex-dev:${ARCH}         libboost-system-dev:${ARCH}         libboost-thread-dev:${ARCH}         libcurl4-gnutls-dev:${ARCH}         libexpat1-dev:${ARCH}         libgoogle-perftools-dev:${ARCH}         libgoogle-perftools4:${ARCH}         libibverbs-dev:${ARCH}         libjemalloc-dev:${ARCH}         libkeyutils-dev:${ARCH}         libldap2-dev:${ARCH}         libnss3-dev:${ARCH}         libpython-dev:${ARCH}         libsnappy-dev:${ARCH}         libssl-dev:${ARCH}         libtcmalloc-minimal4:${ARCH}         libudev-dev:${ARCH}         libunwind-dev:${ARCH}         zlib1g-dev:${ARCH} &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> 8dc3989141e0
Step 7/16 : ARG CEPH_GIT_COMMIT
 ---> Using cache
 ---> 9d5d869cd549
Step 8/16 : RUN cd /build/ceph &&     git fetch --all --prune &&     git checkout -b ceph-builder ${CEPH_GIT_COMMIT} &&     git submodule update --init --recursive
 ---> Using cache
 ---> 1c0a3108f9d6
Step 9/16 : RUN mv /build/ceph/CMakeLists.txt /build/ceph/CMakeLists.original.txt
 ---> Using cache
 ---> 70dfa574cf25
Step 10/16 : ADD CMakeLists.txt /build/ceph
 ---> Using cache
 ---> 3b3c3ec08145
Step 11/16 : ARG CEPH_BUILD_TYPE
 ---> Using cache
 ---> 3f5087dd3915
Step 12/16 : ARG CEPH_ALLOCATOR
 ---> Using cache
 ---> 225c822b412c
Step 13/16 : RUN mkdir -p /build/ceph/build && cd /build/ceph/build &&     cmake     -DCMAKE_SKIP_INSTALL_ALL_DEPENDENCY=ON     -DCMAKE_INSTALL_PREFIX=/usr/local     -DCMAKE_BUILD_TYPE=${CEPH_BUILD_TYPE}     -DCMAKE_TOOLCHAIN_FILE=/usr/local/toolchain/${CROSS_TRIPLE}.cmake     -DALLOCATOR=${CEPH_ALLOCATOR}     -DWITH_SYSTEM_BOOST=ON     -DWITH_EMBEDDED=OFF     -DWITH_FUSE=OFF     -DWITH_LEVELDB=OFF     -DWITH_LTTNG=OFF     -DWITH_MANPAGE=OFF     -DWITH_PROFILER=OFF     -DWITH_PYTHON3=OFF     -DWITH_RADOSGW_FCGI_FRONTEND=OFF     ..
 ---> Using cache
 ---> 95367828ae35
Step 14/16 : WORKDIR /build/ceph/build
 ---> Using cache
 ---> 8c46aca7ea3c
Step 15/16 : ADD build-ceph.sh /build/
 ---> Using cache
 ---> d9e5a61fdc5c
Step 16/16 : CMD /build/build-ceph.sh
 ---> Using cache
 ---> c19572dcbeb0
Successfully built c19572dcbeb0
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[3]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== building ceph from build-be9394dd/ceph-builder-amd64-tcmalloc-release
Scanning dependencies of target civetweb_h
[  0%] keep civetweb.h up-to-date
[  0%] Built target civetweb_h
Scanning dependencies of target common_utf8
[  0%] Building C object src/CMakeFiles/common_utf8.dir/common/utf8.c.o
[  0%] Linking C static library ../lib/libcommon_utf8.a
[  0%] Built target common_utf8
Scanning dependencies of target compressor_objs
[  0%] Building CXX object src/compressor/CMakeFiles/compressor_objs.dir/Compressor.cc.o
[  1%] Building CXX object src/compressor/CMakeFiles/compressor_objs.dir/AsyncCompressor.cc.o
[  1%] Built target compressor_objs
Scanning dependencies of target common-objs
[  1%] Building C object src/CMakeFiles/common-objs.dir/ceph_ver.c.o
[  1%] Building CXX object src/CMakeFiles/common-objs.dir/common/AsyncOpTracker.cc.o
[  1%] Building CXX object src/CMakeFiles/common-objs.dir/common/DecayCounter.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/LogClient.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/LogEntry.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/PrebufferedStreambuf.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/BackTrace.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/perf_counters.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/perf_histogram.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/mutex_debug.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/Mutex.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/OutputDataSocket.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/admin_socket.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/admin_socket_client.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/bloom_filter.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/Readahead.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/cmdparse.cc.o
[  3%] Building C object src/CMakeFiles/common-objs.dir/common/escape.c.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/url_escape.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/io_priority.cc.o
[  3%] Building CXX object src/CMakeFiles/common-objs.dir/common/Clock.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_time.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/mempool.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/Throttle.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/Timer.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/Finisher.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/environment.cc.o
[  5%] Building C object src/CMakeFiles/common-objs.dir/common/sctp_crc32.c.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/crc32c.cc.o
[  5%] Building C object src/CMakeFiles/common-objs.dir/common/crc32c_intel_baseline.c.o
[  5%] Building C object src/CMakeFiles/common-objs.dir/xxHash/xxhash.c.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/assert.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/run_cmd.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/WorkQueue.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/ConfUtils.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/MemoryModel.cc.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/fd.cc.o
[  5%] Building C object src/CMakeFiles/common-objs.dir/common/xattr.c.o
[  5%] Building CXX object src/CMakeFiles/common-objs.dir/common/str_list.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/str_map.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/snap_types.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/errno.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/TrackedOp.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/SloppyCRCMap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/types.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/common/iso_8601.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/log/Log.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/log/SubsystemMap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mon/MonCap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mon/MonClient.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mon/MonMap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mgr/MgrClient.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/msg/simple/Accepter.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/msg/DispatchQueue.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/msg/Message.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mon/PGMap.cc.o
[  7%] Building CXX object src/CMakeFiles/common-objs.dir/mgr/ServiceMap.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/osd/ECMsgTypes.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/osd/HitSet.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/common/RefCountedObj.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/Messenger.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/simple/Pipe.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/simple/PipeConnection.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/simple/SimpleMessenger.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/AsyncConnection.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/AsyncMessenger.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/Event.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/EventSelect.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/Stack.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/PosixStack.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/net_handler.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/QueueStrategy.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/rdma/Infiniband.cc.o
[  9%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/rdma/RDMAConnectedSocketImpl.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/rdma/RDMAServerSocketImpl.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/rdma/RDMAStack.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/msg/msg_types.cc.o
[ 11%] Building C object src/CMakeFiles/common-objs.dir/common/reverse.c.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/hobject.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/osd/OSDMap.cc.o
/build/ceph/src/osd/OSDMap.cc: In member function 'int OSDMap::validate_crush_rules(CrushWrapper*, std::ostream*) const':
/build/ceph/src/osd/OSDMap.cc:3317:25: warning: comparison between signed and unsigned integer expressions [-Wsign-compare]
     if (pool.get_size() < (int)newcrush->get_rule_mask_min_size(ruleno) ||
         ~~~~~~~~~~~~~~~~^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
/build/ceph/src/osd/OSDMap.cc:3318:18: warning: comparison between signed and unsigned integer expressions [-Wsign-compare]
  pool.get_size() > (int)newcrush->get_rule_mask_max_size(ruleno)) {
  ~~~~~~~~~~~~~~~~^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/osd/OSDMapMapping.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/histogram.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/osd/osd_types.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/osd/OpRequest.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/blkdev.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/common_init.cc.o
[ 11%] Building C object src/CMakeFiles/common-objs.dir/common/pipe.c.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_argparse.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_context.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/code_environment.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/dout.cc.o
[ 11%] Building CXX object src/CMakeFiles/common-objs.dir/common/signal.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/Thread.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/Formatter.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/HTMLFormatter.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/HeartbeatMap.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/PluginRegistry.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_fs.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_hash.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_strings.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_frag.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/options.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/config.cc.o
[ 13%] Building C object src/CMakeFiles/common-objs.dir/common/utf8.c.o
[ 13%] Building C object src/CMakeFiles/common-objs.dir/common/mime.c.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/strtol.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/page.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/lockdep.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/version.cc.o
[ 13%] Building CXX object src/CMakeFiles/common-objs.dir/common/hex.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/entity_name.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_crypto.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_crypto_cms.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/ceph_json.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/ipaddr.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/pick_address.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/address_helper.cc.o
[ 15%] Building C object src/CMakeFiles/common-objs.dir/common/linux_version.c.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/TracepointProvider.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/Cycles.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/scrub_types.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/bit_str.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/osdc/Striper.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/osdc/Objecter.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/Graylog.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/fs_types.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/dns_resolve.cc.o
[ 15%] Building CXX object src/CMakeFiles/common-objs.dir/common/hostname.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/common/util.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/arch/probe.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/AuthAuthorizeHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/AuthClientHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/AuthSessionHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/AuthMethodList.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/cephx/CephxAuthorizeHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/cephx/CephxClientHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/cephx/CephxProtocol.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/cephx/CephxSessionHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/none/AuthNoneAuthorizeHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/unknown/AuthUnknownAuthorizeHandler.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/Crypto.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/KeyRing.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/auth/RotatingKeyRing.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/mds/MDSMap.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/mds/FSMap.cc.o
[ 16%] Building CXX object src/CMakeFiles/common-objs.dir/mds/FSMapUser.cc.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/mds/inode_backtrace.cc.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/mds/mdstypes.cc.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/mds/flock.cc.o
[ 18%] Building C object src/CMakeFiles/common-objs.dir/arch/intel.c.o
[ 18%] Building C object src/CMakeFiles/common-objs.dir/common/crc32c_intel_fast.c.o
[ 18%] Building ASM object src/CMakeFiles/common-objs.dir/common/crc32c_intel_fast_asm.s.o
[ 18%] Building ASM object src/CMakeFiles/common-objs.dir/common/crc32c_intel_fast_zero_asm.s.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/msg/async/EventEpoll.cc.o
[ 18%] Building CXX object src/CMakeFiles/common-objs.dir/perfglue/disabled_stubs.cc.o
[ 18%] Built target common-objs
Scanning dependencies of target crush_objs
[ 18%] Building C object src/CMakeFiles/crush_objs.dir/crush/builder.c.o
[ 18%] Building C object src/CMakeFiles/crush_objs.dir/crush/mapper.c.o
[ 18%] Building C object src/CMakeFiles/crush_objs.dir/crush/crush.c.o
[ 18%] Building C object src/CMakeFiles/crush_objs.dir/crush/hash.c.o
[ 18%] Building CXX object src/CMakeFiles/crush_objs.dir/crush/CrushWrapper.cc.o
[ 18%] Building CXX object src/CMakeFiles/crush_objs.dir/crush/CrushCompiler.cc.o
[ 18%] Building CXX object src/CMakeFiles/crush_objs.dir/crush/CrushTester.cc.o
[ 18%] Building CXX object src/CMakeFiles/crush_objs.dir/crush/CrushLocation.cc.o
[ 18%] Built target crush_objs
Scanning dependencies of target common_mountcephfs_objs
[ 18%] Building C object src/CMakeFiles/common_mountcephfs_objs.dir/common/armor.c.o
[ 18%] Building C object src/CMakeFiles/common_mountcephfs_objs.dir/common/safe_io.c.o
[ 18%] Building C object src/CMakeFiles/common_mountcephfs_objs.dir/common/module.c.o
[ 18%] Building C object src/CMakeFiles/common_mountcephfs_objs.dir/common/addr_parsing.c.o
[ 18%] Built target common_mountcephfs_objs
Scanning dependencies of target common_buffer_obj
[ 18%] Building CXX object src/CMakeFiles/common_buffer_obj.dir/common/buffer.cc.o
[ 18%] Built target common_buffer_obj
Scanning dependencies of target common_texttable_obj
[ 18%] Building CXX object src/CMakeFiles/common_texttable_obj.dir/common/TextTable.cc.o
[ 18%] Built target common_texttable_obj
Scanning dependencies of target json_spirit
[ 18%] Building CXX object src/json_spirit/CMakeFiles/json_spirit.dir/json_spirit_reader.cpp.o
[ 18%] Building CXX object src/json_spirit/CMakeFiles/json_spirit.dir/json_spirit_writer.cpp.o
[ 18%] Linking CXX static library ../../lib/libjson_spirit.a
[ 18%] Built target json_spirit
Scanning dependencies of target global_common_objs
[ 18%] Building CXX object src/global/CMakeFiles/global_common_objs.dir/global_context.cc.o
[ 18%] Built target global_common_objs
Scanning dependencies of target erasure_code
[ 18%] Building CXX object src/erasure-code/CMakeFiles/erasure_code.dir/ErasureCodePlugin.cc.o
[ 18%] Linking CXX static library ../../lib/liberasure_code.a
[ 18%] Built target erasure_code
Scanning dependencies of target ceph-common
[ 18%] Linking CXX shared library ../lib/libceph-common.so
[ 18%] Built target ceph-common
Scanning dependencies of target libglobal_objs
[ 18%] Building CXX object src/global/CMakeFiles/libglobal_objs.dir/global_init.cc.o
[ 18%] Building CXX object src/global/CMakeFiles/libglobal_objs.dir/pidfile.cc.o
[ 18%] Building CXX object src/global/CMakeFiles/libglobal_objs.dir/signal_handler.cc.o
[ 18%] Built target libglobal_objs
Scanning dependencies of target global
[ 18%] Linking CXX static library ../../lib/libglobal.a
[ 18%] Built target global
Scanning dependencies of target cls_rgw_client
[ 18%] Building CXX object src/cls/CMakeFiles/cls_rgw_client.dir/rgw/cls_rgw_client.cc.o
[ 18%] Building CXX object src/cls/CMakeFiles/cls_rgw_client.dir/rgw/cls_rgw_types.cc.o
[ 18%] Building CXX object src/cls/CMakeFiles/cls_rgw_client.dir/rgw/cls_rgw_ops.cc.o
[ 18%] Linking CXX static library ../../lib/libcls_rgw_client.a
[ 18%] Built target cls_rgw_client
Scanning dependencies of target cls_lock_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_lock_client.dir/lock/cls_lock_client.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_lock_client.dir/lock/cls_lock_types.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_lock_client.dir/lock/cls_lock_ops.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_lock_client.a
[ 20%] Built target cls_lock_client
Scanning dependencies of target cls_statelog_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_statelog_client.dir/statelog/cls_statelog_client.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_statelog_client.a
[ 20%] Built target cls_statelog_client
Scanning dependencies of target cls_log_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_log_client.dir/log/cls_log_client.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_log_client.a
[ 20%] Built target cls_log_client
Scanning dependencies of target cls_refcount_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_refcount_client.dir/refcount/cls_refcount_client.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_refcount_client.dir/refcount/cls_refcount_ops.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_refcount_client.a
[ 20%] Built target cls_refcount_client
Scanning dependencies of target cls_timeindex_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_timeindex_client.dir/timeindex/cls_timeindex_client.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_timeindex_client.a
[ 20%] Built target cls_timeindex_client
Scanning dependencies of target cls_version_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_version_client.dir/version/cls_version_client.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_version_client.dir/version/cls_version_types.cc.o
[ 20%] Linking CXX static library ../../lib/libcls_version_client.a
[ 20%] Built target cls_version_client
Scanning dependencies of target cls_replica_log_client
[ 20%] Building CXX object src/cls/CMakeFiles/cls_replica_log_client.dir/replica_log/cls_replica_log_types.cc.o
[ 20%] Building CXX object src/cls/CMakeFiles/cls_replica_log_client.dir/replica_log/cls_replica_log_ops.cc.o
[ 22%] Building CXX object src/cls/CMakeFiles/cls_replica_log_client.dir/replica_log/cls_replica_log_client.cc.o
[ 22%] Linking CXX static library ../../lib/libcls_replica_log_client.a
[ 22%] Built target cls_replica_log_client
Scanning dependencies of target cls_user_client
[ 22%] Building CXX object src/cls/CMakeFiles/cls_user_client.dir/user/cls_user_client.cc.o
[ 22%] Building CXX object src/cls/CMakeFiles/cls_user_client.dir/user/cls_user_types.cc.o
[ 22%] Building CXX object src/cls/CMakeFiles/cls_user_client.dir/user/cls_user_ops.cc.o
[ 22%] Linking CXX static library ../../lib/libcls_user_client.a
[ 22%] Built target cls_user_client
Scanning dependencies of target osdc
[ 22%] Building CXX object src/osdc/CMakeFiles/osdc.dir/Filer.cc.o
[ 22%] Building CXX object src/osdc/CMakeFiles/osdc.dir/ObjectCacher.cc.o
[ 22%] Building CXX object src/osdc/CMakeFiles/osdc.dir/Objecter.cc.o
[ 22%] Building CXX object src/osdc/CMakeFiles/osdc.dir/Striper.cc.o
[ 22%] Linking CXX static library ../../lib/libosdc.a
[ 22%] Built target osdc
Scanning dependencies of target librados_api_obj
[ 22%] Building CXX object src/librados/CMakeFiles/librados_api_obj.dir/librados.cc.o
[ 22%] Built target librados_api_obj
Scanning dependencies of target librados_objs
[ 22%] Building CXX object src/librados/CMakeFiles/librados_objs.dir/IoCtxImpl.cc.o
[ 22%] Building CXX object src/librados/CMakeFiles/librados_objs.dir/RadosXattrIter.cc.o
[ 22%] Building CXX object src/librados/CMakeFiles/librados_objs.dir/RadosClient.cc.o
[ 22%] Built target librados_objs
Scanning dependencies of target librados
[ 22%] Linking CXX shared library ../../lib/librados.so
[ 22%] Built target librados
[ 22%] Generate rgw_iam_policy_keywords.frag.cc
Scanning dependencies of target rgw_a
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_acl.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_acl_s3.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_acl_swift.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_auth.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_auth_keystone.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_auth_s3.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_basic_types.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_bucket.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_cache.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_client_io.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_common.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_compression.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_cors.cc.o
[ 22%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_cors_s3.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_dencoder.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_env.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_es_query.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_formats.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_frontend.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_gc.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_http_client.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_json_enc.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_keystone.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_ldap.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_loadgen.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_log.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_lc.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_lc_s3.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_metadata.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_multi.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_multi_del.cc.o
[ 24%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_data_sync.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync_module.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync_module_es.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync_module_es_rest.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_sync_module_log.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_period_history.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_period_puller.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_period_pusher.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_realm_reloader.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_realm_watcher.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_reshard.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_coroutine.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_cr_rados.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_object_expirer_core.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_op.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_os_lib.cc.o
[ 26%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_policy_s3.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_process.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_quota.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rados.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_replica_log.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_request.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_resolve.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_bucket.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_client.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_config.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_conn.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_log.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_metadata.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_opstate.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_realm.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_replica_log.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_role.cc.o
[ 28%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_s3.cc.o
/build/ceph/src/rgw/rgw_rest_s3.cc:3333:13: warning: 'void init_anon_user(req_state*)' defined but not used [-Wunused-function]
 static void init_anon_user(struct req_state *s)
             ^~~~~~~~~~~~~~
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_swift.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_usage.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_rest_user.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_role.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_string.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_swift_auth.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_tag.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_tag_s3.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_tools.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_usage.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_user.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_website.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_xml.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_xml_enc.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_torrent.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_crypt.cc.o
/build/ceph/src/rgw/rgw_crypt.cc:38:2: warning: #warning "TODO: move this code to auth/Crypto for others to reuse." [-Wcpp]
 #warning "TODO: move this code to auth/Crypto for others to reuse."
  ^~~~~~~
/build/ceph/src/rgw/rgw_crypt.cc:247:2: warning: #warning "TODO: use auth/Crypto instead of reimplementing." [-Wcpp]
 #warning "TODO: use auth/Crypto instead of reimplementing."
  ^~~~~~~
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_crypt_sanitize.cc.o
[ 30%] Building CXX object src/rgw/CMakeFiles/rgw_a.dir/rgw_iam_policy.cc.o
[ 32%] Linking CXX static library ../../lib/librgw_a.a
[ 32%] Built target rgw_a
Scanning dependencies of target radosgw-admin
[ 32%] Building CXX object src/rgw/CMakeFiles/radosgw-admin.dir/rgw_admin.cc.o
[ 32%] Building CXX object src/rgw/CMakeFiles/radosgw-admin.dir/rgw_orphan.cc.o
[ 32%] Linking CXX executable ../../bin/radosgw-admin
[ 32%] Built target radosgw-admin
Scanning dependencies of target common
[ 32%] Linking CXX static library ../lib/libcommon.a
[ 32%] Built target common
Scanning dependencies of target global-static
[ 32%] Linking CXX static library ../../lib/libglobal-static.a
[ 32%] Built target global-static
Scanning dependencies of target liblua
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lapi.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lcode.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lctype.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ldebug.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ldo.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ldump.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lfunc.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lgc.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/llex.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lmem.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lobject.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lopcodes.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lparser.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lstate.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lstring.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ltable.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/ltm.c.o
[ 33%] Building C object src/lua/CMakeFiles/liblua.dir/src/lundump.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lvm.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lzio.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lauxlib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lbaselib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lbitlib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lcorolib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/ldblib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/liolib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lmathlib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/loslib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lstrlib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/ltablib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/linit.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/lutf8lib.c.o
[ 35%] Building C object src/lua/CMakeFiles/liblua.dir/src/loadlib_rel.c.o
[ 35%] Linking C static library ../../lib/liblua.a
[ 35%] Built target liblua
Scanning dependencies of target heap_profiler_objs
[ 35%] Building CXX object src/CMakeFiles/heap_profiler_objs.dir/perfglue/heap_profiler.cc.o
[ 35%] Built target heap_profiler_objs
Scanning dependencies of target mds
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/Capability.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSDaemon.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSRank.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/Beacon.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/flock.cc.o
[ 35%] Building C object src/mds/CMakeFiles/mds.dir/locks.c.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/journal.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/Server.cc.o
[ 35%] Building CXX object src/mds/CMakeFiles/mds.dir/Mutation.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/MDCache.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/RecoveryQueue.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/StrayManager.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/PurgeQueue.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/Locker.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/Migrator.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/MDBalancer.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/CDentry.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/CDir.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/CInode.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/LogEvent.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSTable.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/InoTable.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/JournalPointer.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSTableClient.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSTableServer.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/ScrubStack.cc.o
[ 37%] Building CXX object src/mds/CMakeFiles/mds.dir/DamageTable.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/SimpleLock.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/SnapRealm.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/SnapServer.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/snap.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/SessionMap.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSContext.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSAuthCaps.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/MDLog.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/MDSCacheObject.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/Mantle.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/__/common/TrackedOp.cc.o
[ 39%] Building CXX object src/mds/CMakeFiles/mds.dir/__/osdc/Journaler.cc.o
[ 39%] Linking CXX static library ../../lib/libmds.a
[ 39%] Built target mds
Scanning dependencies of target ceph-mds
[ 39%] Building CXX object src/CMakeFiles/ceph-mds.dir/ceph_mds.cc.o
[ 41%] Linking CXX executable ../bin/ceph-mds
[ 41%] Built target ceph-mds
Scanning dependencies of target client
[ 41%] Building CXX object src/client/CMakeFiles/client.dir/Client.cc.o
/build/ceph/src/client/Client.cc: In member function 'void Client::trim_caps(MetaSession*, int)':
/build/ceph/src/client/Client.cc:4116:22: warning: comparison between signed and unsigned integer expressions [-Wsign-compare]
   if (s->caps.size() > max)
       ~~~~~~~~~~~~~~~^~~~~
[ 41%] Building CXX object src/client/CMakeFiles/client.dir/Dentry.cc.o
[ 41%] Building CXX object src/client/CMakeFiles/client.dir/Fh.cc.o
[ 41%] Building CXX object src/client/CMakeFiles/client.dir/Inode.cc.o
[ 41%] Building CXX object src/client/CMakeFiles/client.dir/MetaRequest.cc.o
[ 41%] Building CXX object src/client/CMakeFiles/client.dir/ClientSnapRealm.cc.o
[ 41%] Building CXX object src/client/CMakeFiles/client.dir/MetaSession.cc.o
[ 41%] Building CXX object src/client/CMakeFiles/client.dir/Trace.cc.o
[ 43%] Building CXX object src/client/CMakeFiles/client.dir/posix_acl.cc.o
[ 43%] Linking CXX static library ../../lib/libclient.a
[ 43%] Built target client
Scanning dependencies of target ceph-mgr
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/ceph_mgr.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mon/PGMap.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/DaemonState.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/DaemonServer.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/ClusterState.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/ActivePyModules.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/StandbyPyModules.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/PyModuleRegistry.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/PyModuleRunner.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/PyFormatter.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/PyOSDMap.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/BaseMgrModule.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/BaseMgrStandbyModule.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/ActivePyModule.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/MgrStandby.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/Mgr.cc.o
[ 43%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/Gil.cc.o
[ 45%] Building CXX object src/CMakeFiles/ceph-mgr.dir/mgr/mgr_commands.cc.o
[ 45%] Linking CXX executable ../bin/ceph-mgr
[ 45%] Built target ceph-mgr
Scanning dependencies of target erasure_code_objs
[ 45%] Building CXX object src/erasure-code/CMakeFiles/erasure_code_objs.dir/ErasureCode.cc.o
[ 45%] Built target erasure_code_objs
Scanning dependencies of target ec_isa
[ 45%] Building C object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/ec_base.c.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_2vect_dot_prod_sse.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_3vect_dot_prod_sse.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_4vect_dot_prod_sse.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_5vect_dot_prod_sse.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_6vect_dot_prod_sse.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_dot_prod_sse.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_2vect_mad_avx2.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_3vect_mad_avx2.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_4vect_mad_avx2.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_5vect_mad_avx2.asm.o
[ 45%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_6vect_mad_avx2.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_mad_avx2.asm.o
[ 47%] Building C object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/ec_highlevel_func.c.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_2vect_mad_avx.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_3vect_mad_avx.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_4vect_mad_avx.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_5vect_mad_avx.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_6vect_mad_avx.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_mad_avx.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/ec_multibinary.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_2vect_mad_sse.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_3vect_mad_sse.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_4vect_mad_sse.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_5vect_mad_sse.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_6vect_mad_sse.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_mad_sse.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_2vect_dot_prod_avx2.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_3vect_dot_prod_avx2.asm.o
[ 47%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_4vect_dot_prod_avx2.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_5vect_dot_prod_avx2.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_6vect_dot_prod_avx2.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_dot_prod_avx2.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_mul_avx.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_2vect_dot_prod_avx.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_3vect_dot_prod_avx.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_4vect_dot_prod_avx.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_5vect_dot_prod_avx.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_6vect_dot_prod_avx.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_dot_prod_avx.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_mul_sse.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_2vect_dot_prod_avx512.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_2vect_mad_avx512.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_3vect_dot_prod_avx512.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_3vect_mad_avx512.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_4vect_dot_prod_avx512.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_4vect_mad_avx512.asm.o
[ 49%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_dot_prod_avx512.asm.o
[ 50%] Building ASM object src/erasure-code/isa/CMakeFiles/ec_isa.dir/__/__/isa-l/erasure_code/gf_vect_mad_avx512.asm.o
[ 50%] Building CXX object src/erasure-code/isa/CMakeFiles/ec_isa.dir/ErasureCodeIsa.cc.o
[ 50%] Building CXX object src/erasure-code/isa/CMakeFiles/ec_isa.dir/ErasureCodeIsaTableCache.cc.o
[ 50%] Building CXX object src/erasure-code/isa/CMakeFiles/ec_isa.dir/ErasureCodePluginIsa.cc.o
[ 50%] Building CXX object src/erasure-code/isa/CMakeFiles/ec_isa.dir/xor_op.cc.o
[ 50%] Linking CXX shared library ../../../lib/libec_isa.so
[ 50%] Built target ec_isa
Scanning dependencies of target jerasure_utils
[ 50%] Building CXX object src/erasure-code/jerasure/CMakeFiles/jerasure_utils.dir/ErasureCodePluginJerasure.cc.o
[ 50%] Building CXX object src/erasure-code/jerasure/CMakeFiles/jerasure_utils.dir/ErasureCodeJerasure.cc.o
[ 50%] Built target jerasure_utils
Scanning dependencies of target jerasure_objs
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/jerasure_objs.dir/jerasure/src/cauchy.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/jerasure_objs.dir/jerasure/src/galois.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/jerasure_objs.dir/jerasure/src/jerasure.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/jerasure_objs.dir/jerasure/src/liberation.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/jerasure_objs.dir/jerasure/src/reed_sol.c.o
[ 50%] Building CXX object src/erasure-code/jerasure/CMakeFiles/jerasure_objs.dir/jerasure_init.cc.o
[ 50%] Built target jerasure_objs
Scanning dependencies of target gf-complete_objs
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_cpu.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_wgen.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_w16.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_w32.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_w64.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_w128.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_general.c.o
[ 50%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_w4.c.o
[ 52%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_rand.c.o
[ 52%] Building C object src/erasure-code/jerasure/CMakeFiles/gf-complete_objs.dir/gf-complete/src/gf_w8.c.o
[ 52%] Built target gf-complete_objs
Scanning dependencies of target ec_jerasure_generic
[ 52%] Linking CXX shared library ../../../lib/libec_jerasure_generic.so
[ 52%] Built target ec_jerasure_generic
Scanning dependencies of target ec_jerasure_sse4
[ 52%] Linking CXX shared library ../../../lib/libec_jerasure_sse4.so
[ 52%] Built target ec_jerasure_sse4
Scanning dependencies of target ec_jerasure_sse3
[ 52%] Linking CXX shared library ../../../lib/libec_jerasure_sse3.so
[ 52%] Built target ec_jerasure_sse3
Scanning dependencies of target ec_jerasure
[ 52%] Linking CXX shared library ../../../lib/libec_jerasure.so
[ 52%] Built target ec_jerasure
Scanning dependencies of target ec_lrc
[ 52%] Building CXX object src/erasure-code/lrc/CMakeFiles/ec_lrc.dir/ErasureCodePluginLrc.cc.o
[ 52%] Building CXX object src/erasure-code/lrc/CMakeFiles/ec_lrc.dir/ErasureCodeLrc.cc.o
[ 52%] Building CXX object src/erasure-code/lrc/CMakeFiles/ec_lrc.dir/__/__/common/str_map.cc.o
[ 52%] Linking CXX shared library ../../../lib/libec_lrc.so
[ 52%] Built target ec_lrc
Scanning dependencies of target shec_utils
[ 52%] Building CXX object src/erasure-code/shec/CMakeFiles/shec_utils.dir/__/ErasureCode.cc.o
[ 52%] Building CXX object src/erasure-code/shec/CMakeFiles/shec_utils.dir/ErasureCodePluginShec.cc.o
[ 52%] Building CXX object src/erasure-code/shec/CMakeFiles/shec_utils.dir/ErasureCodeShec.cc.o
[ 52%] Building CXX object src/erasure-code/shec/CMakeFiles/shec_utils.dir/ErasureCodeShecTableCache.cc.o
[ 52%] Building C object src/erasure-code/shec/CMakeFiles/shec_utils.dir/determinant.c.o
[ 52%] Built target shec_utils
Scanning dependencies of target ec_shec_sse3
[ 52%] Linking CXX shared library ../../../lib/libec_shec_sse3.so
[ 52%] Built target ec_shec_sse3
Scanning dependencies of target ec_shec_generic
[ 52%] Linking CXX shared library ../../../lib/libec_shec_generic.so
[ 52%] Built target ec_shec_generic
Scanning dependencies of target ec_shec_sse4
[ 52%] Linking CXX shared library ../../../lib/libec_shec_sse4.so
[ 52%] Built target ec_shec_sse4
Scanning dependencies of target ec_shec
[ 54%] Linking CXX shared library ../../../lib/libec_shec.so
[ 54%] Built target ec_shec
Scanning dependencies of target erasure_code_plugins
[ 54%] Built target erasure_code_plugins
Scanning dependencies of target rocksdb_ext
[ 54%] Creating directories for 'rocksdb_ext'
[ 54%] No download step for 'rocksdb_ext'
[ 54%] No patch step for 'rocksdb_ext'
[ 54%] No update step for 'rocksdb_ext'
[ 54%] Performing configure step for 'rocksdb_ext'
-- The C compiler identification is GNU 6.3.0
-- The CXX compiler identification is GNU 6.3.0
-- Check for working C compiler: /usr/bin/cc
-- Check for working C compiler: /usr/bin/cc -- works
-- Detecting C compiler ABI info
-- Detecting C compiler ABI info - done
-- Detecting C compile features
-- Detecting C compile features - done
-- Check for working CXX compiler: /usr/lib/ccache/x86_64-linux-gnu-g++
-- Check for working CXX compiler: /usr/lib/ccache/x86_64-linux-gnu-g++ -- works
-- Detecting CXX compiler ABI info
-- Detecting CXX compiler ABI info - done
-- Detecting CXX compile features
-- Detecting CXX compile features - done
-- Found Git: /usr/bin/git (found version "2.11.0") 
-- Performing Test HAVE_OMIT_LEAF_FRAME_POINTER
-- Performing Test HAVE_OMIT_LEAF_FRAME_POINTER - Success
-- Performing Test HAVE_FALLOCATE
-- Performing Test HAVE_FALLOCATE - Success
-- Looking for malloc_usable_size
-- Looking for malloc_usable_size - found
-- Looking for pthread.h
-- Looking for pthread.h - found
-- Looking for pthread_create
-- Looking for pthread_create - not found
-- Looking for pthread_create in pthreads
-- Looking for pthread_create in pthreads - not found
-- Looking for pthread_create in pthread
-- Looking for pthread_create in pthread - found
-- Found Threads: TRUE  
-- JNI library is disabled
-- Configuring done
-- Generating done
-- Build files have been written to: /build/ceph/build/src/rocksdb
[ 54%] Performing forcebuild step for 'rocksdb_ext'
[ 54%] Performing build step for 'rocksdb_ext'
Scanning dependencies of target build_version
[  0%] Building CXX object CMakeFiles/build_version.dir/build_version.cc.o
[  0%] Built target build_version
Scanning dependencies of target rocksdb
[  0%] Building CXX object CMakeFiles/rocksdb.dir/cache/clock_cache.cc.o
[  0%] Building CXX object CMakeFiles/rocksdb.dir/cache/lru_cache.cc.o
[  3%] Building CXX object CMakeFiles/rocksdb.dir/cache/sharded_cache.cc.o
[  3%] Building CXX object CMakeFiles/rocksdb.dir/db/builder.cc.o
[  3%] Building CXX object CMakeFiles/rocksdb.dir/db/c.cc.o
[  3%] Building CXX object CMakeFiles/rocksdb.dir/db/column_family.cc.o
[  3%] Building CXX object CMakeFiles/rocksdb.dir/db/compacted_db_impl.cc.o
[  3%] Building CXX object CMakeFiles/rocksdb.dir/db/compaction.cc.o
[  3%] Building CXX object CMakeFiles/rocksdb.dir/db/compaction_iterator.cc.o
[  6%] Building CXX object CMakeFiles/rocksdb.dir/db/compaction_job.cc.o
[  6%] Building CXX object CMakeFiles/rocksdb.dir/db/compaction_picker.cc.o
[  6%] Building CXX object CMakeFiles/rocksdb.dir/db/compaction_picker_universal.cc.o
[  6%] Building CXX object CMakeFiles/rocksdb.dir/db/convenience.cc.o
[  6%] Building CXX object CMakeFiles/rocksdb.dir/db/db_filesnapshot.cc.o
[  6%] Building CXX object CMakeFiles/rocksdb.dir/db/db_impl.cc.o
[  6%] Building CXX object CMakeFiles/rocksdb.dir/db/db_impl_write.cc.o
[ 10%] Building CXX object CMakeFiles/rocksdb.dir/db/db_impl_compaction_flush.cc.o
[ 10%] Building CXX object CMakeFiles/rocksdb.dir/db/db_impl_files.cc.o
[ 10%] Building CXX object CMakeFiles/rocksdb.dir/db/db_impl_open.cc.o
[ 10%] Building CXX object CMakeFiles/rocksdb.dir/db/db_impl_debug.cc.o
[ 10%] Building CXX object CMakeFiles/rocksdb.dir/db/db_impl_experimental.cc.o
[ 10%] Building CXX object CMakeFiles/rocksdb.dir/db/db_impl_readonly.cc.o
[ 10%] Building CXX object CMakeFiles/rocksdb.dir/db/db_info_dumper.cc.o
[ 13%] Building CXX object CMakeFiles/rocksdb.dir/db/db_iter.cc.o
[ 13%] Building CXX object CMakeFiles/rocksdb.dir/db/dbformat.cc.o
[ 13%] Building CXX object CMakeFiles/rocksdb.dir/db/event_helpers.cc.o
[ 13%] Building CXX object CMakeFiles/rocksdb.dir/db/experimental.cc.o
[ 13%] Building CXX object CMakeFiles/rocksdb.dir/db/external_sst_file_ingestion_job.cc.o
[ 13%] Building CXX object CMakeFiles/rocksdb.dir/db/file_indexer.cc.o
[ 17%] Building CXX object CMakeFiles/rocksdb.dir/db/flush_job.cc.o
[ 17%] Building CXX object CMakeFiles/rocksdb.dir/db/flush_scheduler.cc.o
[ 17%] Building CXX object CMakeFiles/rocksdb.dir/db/forward_iterator.cc.o
[ 17%] Building CXX object CMakeFiles/rocksdb.dir/db/internal_stats.cc.o
[ 17%] Building CXX object CMakeFiles/rocksdb.dir/db/log_reader.cc.o
[ 17%] Building CXX object CMakeFiles/rocksdb.dir/db/log_writer.cc.o
[ 17%] Building CXX object CMakeFiles/rocksdb.dir/db/managed_iterator.cc.o
[ 20%] Building CXX object CMakeFiles/rocksdb.dir/db/memtable.cc.o
[ 20%] Building CXX object CMakeFiles/rocksdb.dir/db/memtable_list.cc.o
[ 20%] Building CXX object CMakeFiles/rocksdb.dir/db/merge_helper.cc.o
[ 20%] Building CXX object CMakeFiles/rocksdb.dir/db/merge_operator.cc.o
[ 20%] Building CXX object CMakeFiles/rocksdb.dir/db/range_del_aggregator.cc.o
[ 20%] Building CXX object CMakeFiles/rocksdb.dir/db/repair.cc.o
[ 20%] Building CXX object CMakeFiles/rocksdb.dir/db/snapshot_impl.cc.o
[ 24%] Building CXX object CMakeFiles/rocksdb.dir/db/table_cache.cc.o
[ 24%] Building CXX object CMakeFiles/rocksdb.dir/db/table_properties_collector.cc.o
[ 24%] Building CXX object CMakeFiles/rocksdb.dir/db/transaction_log_impl.cc.o
[ 24%] Building CXX object CMakeFiles/rocksdb.dir/db/version_builder.cc.o
[ 24%] Building CXX object CMakeFiles/rocksdb.dir/db/version_edit.cc.o
[ 24%] Building CXX object CMakeFiles/rocksdb.dir/db/version_set.cc.o
[ 24%] Building CXX object CMakeFiles/rocksdb.dir/db/wal_manager.cc.o
[ 27%] Building CXX object CMakeFiles/rocksdb.dir/db/write_batch.cc.o
[ 27%] Building CXX object CMakeFiles/rocksdb.dir/db/write_batch_base.cc.o
[ 27%] Building CXX object CMakeFiles/rocksdb.dir/db/write_controller.cc.o
[ 27%] Building CXX object CMakeFiles/rocksdb.dir/db/write_thread.cc.o
[ 27%] Building CXX object CMakeFiles/rocksdb.dir/env/env.cc.o
[ 27%] Building CXX object CMakeFiles/rocksdb.dir/env/env_chroot.cc.o
[ 31%] Building CXX object CMakeFiles/rocksdb.dir/env/env_hdfs.cc.o
[ 31%] Building CXX object CMakeFiles/rocksdb.dir/env/memenv.cc.o
[ 31%] Building CXX object CMakeFiles/rocksdb.dir/memtable/hash_cuckoo_rep.cc.o
[ 31%] Building CXX object CMakeFiles/rocksdb.dir/memtable/hash_linklist_rep.cc.o
[ 31%] Building CXX object CMakeFiles/rocksdb.dir/memtable/hash_skiplist_rep.cc.o
[ 31%] Building CXX object CMakeFiles/rocksdb.dir/memtable/memtable_allocator.cc.o
[ 31%] Building CXX object CMakeFiles/rocksdb.dir/memtable/skiplistrep.cc.o
[ 34%] Building CXX object CMakeFiles/rocksdb.dir/memtable/vectorrep.cc.o
[ 34%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/histogram.cc.o
[ 34%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/histogram_windowing.cc.o
[ 34%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/instrumented_mutex.cc.o
[ 34%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/iostats_context.cc.o
[ 34%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/perf_context.cc.o
[ 34%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/perf_level.cc.o
[ 37%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/statistics.cc.o
[ 37%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/thread_status_impl.cc.o
[ 37%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/thread_status_updater.cc.o
[ 37%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/thread_status_util.cc.o
[ 37%] Building CXX object CMakeFiles/rocksdb.dir/monitoring/thread_status_util_debug.cc.o
[ 37%] Building CXX object CMakeFiles/rocksdb.dir/options/cf_options.cc.o
[ 37%] Building CXX object CMakeFiles/rocksdb.dir/options/db_options.cc.o
[ 41%] Building CXX object CMakeFiles/rocksdb.dir/options/options.cc.o
[ 41%] Building CXX object CMakeFiles/rocksdb.dir/options/options_helper.cc.o
[ 41%] Building CXX object CMakeFiles/rocksdb.dir/options/options_parser.cc.o
[ 41%] Building CXX object CMakeFiles/rocksdb.dir/options/options_sanity_check.cc.o
[ 41%] Building CXX object CMakeFiles/rocksdb.dir/port/stack_trace.cc.o
[ 41%] Building CXX object CMakeFiles/rocksdb.dir/table/adaptive_table_factory.cc.o
[ 44%] Building CXX object CMakeFiles/rocksdb.dir/table/block.cc.o
[ 44%] Building CXX object CMakeFiles/rocksdb.dir/table/block_based_filter_block.cc.o
[ 44%] Building CXX object CMakeFiles/rocksdb.dir/table/block_based_table_builder.cc.o
[ 44%] Building CXX object CMakeFiles/rocksdb.dir/table/block_based_table_factory.cc.o
[ 44%] Building CXX object CMakeFiles/rocksdb.dir/table/block_based_table_reader.cc.o
[ 44%] Building CXX object CMakeFiles/rocksdb.dir/table/block_builder.cc.o
[ 44%] Building CXX object CMakeFiles/rocksdb.dir/table/block_prefix_index.cc.o
[ 48%] Building CXX object CMakeFiles/rocksdb.dir/table/bloom_block.cc.o
[ 48%] Building CXX object CMakeFiles/rocksdb.dir/table/cuckoo_table_builder.cc.o
[ 48%] Building CXX object CMakeFiles/rocksdb.dir/table/cuckoo_table_factory.cc.o
[ 48%] Building CXX object CMakeFiles/rocksdb.dir/table/cuckoo_table_reader.cc.o
[ 48%] Building CXX object CMakeFiles/rocksdb.dir/table/flush_block_policy.cc.o
[ 48%] Building CXX object CMakeFiles/rocksdb.dir/table/format.cc.o
[ 48%] Building CXX object CMakeFiles/rocksdb.dir/table/full_filter_block.cc.o
[ 51%] Building CXX object CMakeFiles/rocksdb.dir/table/get_context.cc.o
[ 51%] Building CXX object CMakeFiles/rocksdb.dir/table/index_builder.cc.o
[ 51%] Building CXX object CMakeFiles/rocksdb.dir/table/iterator.cc.o
[ 51%] Building CXX object CMakeFiles/rocksdb.dir/table/merging_iterator.cc.o
[ 51%] Building CXX object CMakeFiles/rocksdb.dir/table/meta_blocks.cc.o
[ 51%] Building CXX object CMakeFiles/rocksdb.dir/table/partitioned_filter_block.cc.o
[ 51%] Building CXX object CMakeFiles/rocksdb.dir/table/persistent_cache_helper.cc.o
[ 55%] Building CXX object CMakeFiles/rocksdb.dir/table/plain_table_builder.cc.o
[ 55%] Building CXX object CMakeFiles/rocksdb.dir/table/plain_table_factory.cc.o
[ 55%] Building CXX object CMakeFiles/rocksdb.dir/table/plain_table_index.cc.o
[ 55%] Building CXX object CMakeFiles/rocksdb.dir/table/plain_table_key_coding.cc.o
[ 55%] Building CXX object CMakeFiles/rocksdb.dir/table/plain_table_reader.cc.o
[ 55%] Building CXX object CMakeFiles/rocksdb.dir/table/sst_file_writer.cc.o
[ 58%] Building CXX object CMakeFiles/rocksdb.dir/table/table_properties.cc.o
[ 58%] Building CXX object CMakeFiles/rocksdb.dir/table/two_level_iterator.cc.o
[ 58%] Building CXX object CMakeFiles/rocksdb.dir/tools/db_bench_tool.cc.o
[ 58%] Building CXX object CMakeFiles/rocksdb.dir/tools/dump/db_dump_tool.cc.o
[ 58%] Building CXX object CMakeFiles/rocksdb.dir/tools/ldb_cmd.cc.o
[ 58%] Building CXX object CMakeFiles/rocksdb.dir/tools/ldb_tool.cc.o
[ 58%] Building CXX object CMakeFiles/rocksdb.dir/tools/sst_dump_tool.cc.o
[ 62%] Building CXX object CMakeFiles/rocksdb.dir/util/arena.cc.o
[ 62%] Building CXX object CMakeFiles/rocksdb.dir/util/auto_roll_logger.cc.o
[ 62%] Building CXX object CMakeFiles/rocksdb.dir/util/bloom.cc.o
[ 62%] Building CXX object CMakeFiles/rocksdb.dir/util/coding.cc.o
[ 62%] Building CXX object CMakeFiles/rocksdb.dir/util/compaction_job_stats_impl.cc.o
[ 62%] Building CXX object CMakeFiles/rocksdb.dir/util/comparator.cc.o
[ 62%] Building CXX object CMakeFiles/rocksdb.dir/util/concurrent_arena.cc.o
[ 65%] Building CXX object CMakeFiles/rocksdb.dir/util/crc32c.cc.o
[ 65%] Building CXX object CMakeFiles/rocksdb.dir/util/delete_scheduler.cc.o
[ 65%] Building CXX object CMakeFiles/rocksdb.dir/util/dynamic_bloom.cc.o
[ 65%] Building CXX object CMakeFiles/rocksdb.dir/util/event_logger.cc.o
[ 65%] Building CXX object CMakeFiles/rocksdb.dir/util/file_reader_writer.cc.o
[ 65%] Building CXX object CMakeFiles/rocksdb.dir/util/file_util.cc.o
[ 65%] Building CXX object CMakeFiles/rocksdb.dir/util/filename.cc.o
[ 68%] Building CXX object CMakeFiles/rocksdb.dir/util/filter_policy.cc.o
[ 68%] Building CXX object CMakeFiles/rocksdb.dir/util/hash.cc.o
[ 68%] Building CXX object CMakeFiles/rocksdb.dir/util/log_buffer.cc.o
[ 68%] Building CXX object CMakeFiles/rocksdb.dir/util/murmurhash.cc.o
[ 68%] Building CXX object CMakeFiles/rocksdb.dir/util/random.cc.o
[ 68%] Building CXX object CMakeFiles/rocksdb.dir/util/rate_limiter.cc.o
[ 68%] Building CXX object CMakeFiles/rocksdb.dir/util/slice.cc.o
[ 72%] Building CXX object CMakeFiles/rocksdb.dir/util/sst_file_manager_impl.cc.o
[ 72%] Building CXX object CMakeFiles/rocksdb.dir/util/status.cc.o
[ 72%] Building CXX object CMakeFiles/rocksdb.dir/util/status_message.cc.o
[ 72%] Building CXX object CMakeFiles/rocksdb.dir/util/string_util.cc.o
[ 72%] Building CXX object CMakeFiles/rocksdb.dir/util/sync_point.cc.o
[ 72%] Building CXX object CMakeFiles/rocksdb.dir/util/testutil.cc.o
[ 75%] Building CXX object CMakeFiles/rocksdb.dir/util/thread_local.cc.o
[ 75%] Building CXX object CMakeFiles/rocksdb.dir/util/threadpool_imp.cc.o
[ 75%] Building CXX object CMakeFiles/rocksdb.dir/util/transaction_test_util.cc.o
[ 75%] Building CXX object CMakeFiles/rocksdb.dir/util/xxhash.cc.o
[ 75%] Building CXX object CMakeFiles/rocksdb.dir/utilities/backupable/backupable_db.cc.o
[ 75%] Building CXX object CMakeFiles/rocksdb.dir/utilities/blob_db/blob_db.cc.o
[ 75%] Building CXX object CMakeFiles/rocksdb.dir/utilities/checkpoint/checkpoint_impl.cc.o
[ 79%] Building CXX object CMakeFiles/rocksdb.dir/utilities/col_buf_decoder.cc.o
[ 79%] Building CXX object CMakeFiles/rocksdb.dir/utilities/col_buf_encoder.cc.o
[ 79%] Building CXX object CMakeFiles/rocksdb.dir/utilities/column_aware_encoding_util.cc.o
[ 79%] Building CXX object CMakeFiles/rocksdb.dir/utilities/compaction_filters/remove_emptyvalue_compactionfilter.cc.o
[ 79%] Building CXX object CMakeFiles/rocksdb.dir/utilities/date_tiered/date_tiered_db_impl.cc.o
[ 79%] Building CXX object CMakeFiles/rocksdb.dir/utilities/document/document_db.cc.o
[ 79%] Building CXX object CMakeFiles/rocksdb.dir/utilities/document/json_document.cc.o
[ 82%] Building CXX object CMakeFiles/rocksdb.dir/utilities/document/json_document_builder.cc.o
[ 82%] Building CXX object CMakeFiles/rocksdb.dir/utilities/env_mirror.cc.o
[ 82%] Building CXX object CMakeFiles/rocksdb.dir/utilities/env_timed.cc.o
[ 82%] Building CXX object CMakeFiles/rocksdb.dir/utilities/geodb/geodb_impl.cc.o
[ 82%] Building CXX object CMakeFiles/rocksdb.dir/utilities/leveldb_options/leveldb_options.cc.o
[ 82%] Building CXX object CMakeFiles/rocksdb.dir/utilities/lua/rocks_lua_compaction_filter.cc.o
[ 82%] Building CXX object CMakeFiles/rocksdb.dir/utilities/memory/memory_util.cc.o
[ 86%] Building CXX object CMakeFiles/rocksdb.dir/utilities/merge_operators/max.cc.o
[ 86%] Building CXX object CMakeFiles/rocksdb.dir/utilities/merge_operators/put.cc.o
[ 86%] Building CXX object CMakeFiles/rocksdb.dir/utilities/merge_operators/string_append/stringappend.cc.o
[ 86%] Building CXX object CMakeFiles/rocksdb.dir/utilities/merge_operators/string_append/stringappend2.cc.o
[ 86%] Building CXX object CMakeFiles/rocksdb.dir/utilities/merge_operators/uint64add.cc.o
[ 86%] Building CXX object CMakeFiles/rocksdb.dir/utilities/option_change_migration/option_change_migration.cc.o
[ 89%] Building CXX object CMakeFiles/rocksdb.dir/utilities/options/options_util.cc.o
[ 89%] Building CXX object CMakeFiles/rocksdb.dir/utilities/persistent_cache/block_cache_tier.cc.o
[ 89%] Building CXX object CMakeFiles/rocksdb.dir/utilities/persistent_cache/block_cache_tier_file.cc.o
[ 89%] Building CXX object CMakeFiles/rocksdb.dir/utilities/persistent_cache/block_cache_tier_metadata.cc.o
[ 89%] Building CXX object CMakeFiles/rocksdb.dir/utilities/persistent_cache/persistent_cache_tier.cc.o
[ 89%] Building CXX object CMakeFiles/rocksdb.dir/utilities/persistent_cache/volatile_tier_impl.cc.o
[ 89%] Building CXX object CMakeFiles/rocksdb.dir/utilities/redis/redis_lists.cc.o
[ 93%] Building CXX object CMakeFiles/rocksdb.dir/utilities/simulator_cache/sim_cache.cc.o
[ 93%] Building CXX object CMakeFiles/rocksdb.dir/utilities/spatialdb/spatial_db.cc.o
[ 93%] Building CXX object CMakeFiles/rocksdb.dir/utilities/table_properties_collectors/compact_on_deletion_collector.cc.o
[ 93%] Building CXX object CMakeFiles/rocksdb.dir/utilities/transactions/optimistic_transaction_db_impl.cc.o
[ 93%] Building CXX object CMakeFiles/rocksdb.dir/utilities/transactions/optimistic_transaction_impl.cc.o
[ 93%] Building CXX object CMakeFiles/rocksdb.dir/utilities/transactions/transaction_base.cc.o
[ 93%] Building CXX object CMakeFiles/rocksdb.dir/utilities/transactions/transaction_db_impl.cc.o
[ 96%] Building CXX object CMakeFiles/rocksdb.dir/utilities/transactions/transaction_db_mutex_impl.cc.o
[ 96%] Building CXX object CMakeFiles/rocksdb.dir/utilities/transactions/transaction_impl.cc.o
[ 96%] Building CXX object CMakeFiles/rocksdb.dir/utilities/transactions/transaction_lock_mgr.cc.o
[ 96%] Building CXX object CMakeFiles/rocksdb.dir/utilities/transactions/transaction_util.cc.o
[ 96%] Building CXX object CMakeFiles/rocksdb.dir/utilities/ttl/db_ttl_impl.cc.o
[ 96%] Building CXX object CMakeFiles/rocksdb.dir/utilities/write_batch_with_index/write_batch_with_index.cc.o
[ 96%] Building CXX object CMakeFiles/rocksdb.dir/utilities/write_batch_with_index/write_batch_with_index_internal.cc.o
[100%] Building CXX object CMakeFiles/rocksdb.dir/port/port_posix.cc.o
[100%] Building CXX object CMakeFiles/rocksdb.dir/env/env_posix.cc.o
[100%] Building CXX object CMakeFiles/rocksdb.dir/env/io_posix.cc.o
[100%] Linking CXX static library librocksdb.a
[100%] Built target rocksdb
[ 54%] Performing install step for 'rocksdb_ext'
[ 54%] Completed 'rocksdb_ext'
[ 54%] Built target rocksdb_ext
Scanning dependencies of target kv_objs
[ 54%] Building CXX object src/kv/CMakeFiles/kv_objs.dir/KeyValueDB.cc.o
[ 54%] Building CXX object src/kv/CMakeFiles/kv_objs.dir/MemDB.cc.o
[ 54%] Building CXX object src/kv/CMakeFiles/kv_objs.dir/RocksDBStore.cc.o
[ 54%] Built target kv_objs
Scanning dependencies of target mon
[ 54%] Building CXX object src/mon/CMakeFiles/mon.dir/__/auth/cephx/CephxKeyServer.cc.o
[ 54%] Building CXX object src/mon/CMakeFiles/mon.dir/__/auth/cephx/CephxServiceHandler.cc.o
[ 54%] Building CXX object src/mon/CMakeFiles/mon.dir/__/auth/AuthServiceHandler.cc.o
[ 54%] Building CXX object src/mon/CMakeFiles/mon.dir/Paxos.cc.o
[ 54%] Building CXX object src/mon/CMakeFiles/mon.dir/PaxosService.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/OSDMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/MDSMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/FSCommands.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/MgrMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/MgrStatMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/Monitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/MonmapMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/LogMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/AuthMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/Elector.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/HealthMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/OldHealthMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/DataHealthService.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/PGMonitor.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/PGMap.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/ConfigKeyService.cc.o
[ 56%] Building CXX object src/mon/CMakeFiles/mon.dir/__/mgr/mgr_commands.cc.o
[ 56%] Linking CXX static library ../../lib/libmon.a
[ 56%] Built target mon
Scanning dependencies of target kv
[ 56%] Linking CXX static library ../../lib/libkv.a
[ 56%] Built target kv
Scanning dependencies of target ceph_crypto_isal
[ 56%] Building CXX object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isal_crypto_accel.cc.o
[ 56%] Building CXX object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isal_crypto_plugin.cc.o
[ 56%] Building C object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_pre.c.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_multibinary.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/keyexp_128.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/keyexp_192.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/keyexp_256.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/keyexp_multibinary.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_dec_128_x4_sse.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_dec_128_x8_avx.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_dec_192_x4_sse.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_dec_192_x8_avx.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_dec_256_x4_sse.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_dec_256_x8_avx.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_enc_128_x4_sb.asm.o
[ 56%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_enc_128_x8_sb.asm.o
[ 58%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_enc_192_x4_sb.asm.o
[ 58%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_enc_192_x8_sb.asm.o
[ 58%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_enc_256_x4_sb.asm.o
[ 58%] Building ASM object src/crypto/isa-l/CMakeFiles/ceph_crypto_isal.dir/isa-l_crypto/aes/cbc_enc_256_x8_sb.asm.o
[ 58%] Linking CXX shared library ../../../lib/libceph_crypto_isal.so
[ 58%] Built target ceph_crypto_isal
Scanning dependencies of target crypto_plugins
[ 58%] Built target crypto_plugins
Scanning dependencies of target zstd_ext
[ 58%] Creating directories for 'zstd_ext'
[ 58%] No download step for 'zstd_ext'
[ 58%] No patch step for 'zstd_ext'
[ 58%] No update step for 'zstd_ext'
[ 58%] Performing configure step for 'zstd_ext'
-- The C compiler identification is GNU 6.3.0
-- The CXX compiler identification is GNU 6.3.0
-- Check for working C compiler: /usr/lib/ccache/x86_64-linux-gnu-gcc
-- Check for working C compiler: /usr/lib/ccache/x86_64-linux-gnu-gcc -- works
-- Detecting C compiler ABI info
-- Detecting C compiler ABI info - done
-- Detecting C compile features
-- Detecting C compile features - done
-- Check for working CXX compiler: /usr/lib/ccache/x86_64-linux-gnu-g++
-- Check for working CXX compiler: /usr/lib/ccache/x86_64-linux-gnu-g++ -- works
-- Detecting CXX compiler ABI info
-- Detecting CXX compiler ABI info - done
-- Detecting CXX compile features
-- Detecting CXX compile features - done
-- ZSTD_LEGACY_SUPPORT not defined!
ZSTD VERSION 1.1.1024
the variable PREFIX=/usr/local
-- Performing Test POSITION_INDEPENDENT_CODE_FLAG_ALLOWED
-- Performing Test POSITION_INDEPENDENT_CODE_FLAG_ALLOWED - Success
Compiler flag -fPIC allowed
-- Performing Test WARNING_UNDEF_ALLOWED
-- Performing Test WARNING_UNDEF_ALLOWED - Success
Compiler flag -Wundef allowed
-- Performing Test WARNING_SHADOW_ALLOWED
-- Performing Test WARNING_SHADOW_ALLOWED - Success
Compiler flag -Wshadow allowed
-- Performing Test WARNING_CAST_ALIGN_ALLOWED
-- Performing Test WARNING_CAST_ALIGN_ALLOWED - Success
Compiler flag -Wcast-align allowed
-- Performing Test WARNING_CAST_QUAL_ALLOWED
-- Performing Test WARNING_CAST_QUAL_ALLOWED - Success
Compiler flag -Wcast-qual allowed
-- Performing Test WARNING_STRICT_PROTOTYPES_ALLOWED
-- Performing Test WARNING_STRICT_PROTOTYPES_ALLOWED - Success
Compiler flag -Wstrict-prototypes allowed
-- Performing Test WARNING_ALL_ALLOWED
-- Performing Test WARNING_ALL_ALLOWED - Success
Compiler flag -Wall allowed
-- Performing Test WARNING_EXTRA_ALLOWED
-- Performing Test WARNING_EXTRA_ALLOWED - Success
Compiler flag -Wextra allowed
-- Performing Test WARNING_FLOAT_EQUAL_ALLOWED
-- Performing Test WARNING_FLOAT_EQUAL_ALLOWED - Success
Compiler flag -Wfloat-equal allowed
-- Performing Test WARNING_SIGN_CONVERSION_ALLOWED
-- Performing Test WARNING_SIGN_CONVERSION_ALLOWED - Success
Compiler flag -Wsign-conversion allowed
-- Configuring done
-- Generating done
-- Build files have been written to: /build/ceph/build/src/compressor/zstd/libzstd
[ 60%] Performing forcebuild step for 'zstd_ext'
[ 60%] Performing build step for 'zstd_ext'
Scanning dependencies of target libzstd_static
[  6%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/common/entropy_common.c.o
[ 13%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/common/zstd_common.c.o
[ 20%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/common/error_private.c.o
[ 26%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/common/xxhash.c.o
[ 33%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/common/fse_decompress.c.o
[ 40%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/compress/fse_compress.c.o
[ 46%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/compress/huf_compress.c.o
[ 53%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/compress/zstd_compress.c.o
[ 60%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/decompress/huf_decompress.c.o
[ 66%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/decompress/zstd_decompress.c.o
[ 73%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/dictBuilder/divsufsort.c.o
[ 80%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/dictBuilder/zdict.c.o
[ 86%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/deprecated/zbuff_compress.c.o
[ 93%] Building C object lib/CMakeFiles/libzstd_static.dir/build/ceph/src/zstd/lib/deprecated/zbuff_decompress.c.o
[100%] Linking C static library libzstd.a
[100%] Built target libzstd_static
[ 60%] Performing install step for 'zstd_ext'
[ 60%] Completed 'zstd_ext'
[ 60%] Built target zstd_ext
Scanning dependencies of target ceph_zstd
[ 60%] Building CXX object src/compressor/zstd/CMakeFiles/ceph_zstd.dir/CompressionPluginZstd.cc.o
[ 60%] Linking CXX shared library ../../../lib/libceph_zstd.so
[ 60%] Built target ceph_zstd
Scanning dependencies of target ceph_snappy
[ 60%] Building CXX object src/compressor/snappy/CMakeFiles/ceph_snappy.dir/CompressionPluginSnappy.cc.o
[ 60%] Linking CXX shared library ../../../lib/libceph_snappy.so
[ 60%] Built target ceph_snappy
Scanning dependencies of target ceph_zlib
[ 60%] Building CXX object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/CompressionPluginZlib.cc.o
[ 60%] Building CXX object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/ZlibCompressor.cc.o
[ 60%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/crc32_gzip.asm.o
[ 60%] Building C object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/crc32_gzip_base.c.o
[ 60%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/detect_repeated_char.asm.o
[ 60%] Building C object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/encode_df.c.o
[ 60%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/encode_df_04.asm.o
[ 60%] Building C object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/flatten_ll.c.o
[ 62%] Building C object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/huff_codes.c.o
[ 62%] Building C object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/hufftables_c.c.o
[ 62%] Building C object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip.c.o
[ 62%] Building C object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_base.c.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_body_01.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_body_02.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_body_04.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_finish.asm.o
[ 62%] Building C object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_icf_base.c.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_icf_body_01.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_icf_body_02.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_icf_body_04.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_icf_finish.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_multibinary.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_update_histogram_01.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/igzip_update_histogram_04.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/proc_heap.asm.o
[ 62%] Building ASM object src/compressor/zlib/CMakeFiles/ceph_zlib.dir/__/__/isa-l/igzip/rfc1951_lookup.asm.o
[ 64%] Linking CXX shared library ../../../lib/libceph_zlib.so
[ 64%] Built target ceph_zlib
Scanning dependencies of target compressor_plugins
[ 64%] Built target compressor_plugins
Scanning dependencies of target os
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/ObjectStore.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/Transaction.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/chain_xattr.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/BtrfsFileStoreBackend.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/DBObjectMap.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/FileJournal.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/FileStore.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/JournalThrottle.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/GenericFileStoreBackend.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/JournalingObjectStore.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/HashIndex.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/IndexManager.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/LFNIndex.cc.o
[ 64%] Building CXX object src/os/CMakeFiles/os.dir/filestore/WBThrottle.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/memstore/MemStore.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/kstore/KStore.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/kstore/kstore_types.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/fs/FS.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/Allocator.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/BitmapFreelistManager.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/BlockDevice.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/BlueFS.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/bluefs_types.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/BlueRocksEnv.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/BlueStore.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/bluestore_types.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/FreelistManager.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/KernelDevice.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/StupidAllocator.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/BitMapAllocator.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/BitAllocator.cc.o
[ 66%] Building CXX object src/os/CMakeFiles/os.dir/bluestore/aio.cc.o
[ 67%] Linking CXX static library ../../lib/libos.a
[ 67%] Built target os
Scanning dependencies of target ceph-mon
[ 67%] Building CXX object src/CMakeFiles/ceph-mon.dir/ceph_mon.cc.o
[ 67%] Linking CXX executable ../bin/ceph-mon
[ 67%] Built target ceph-mon
Scanning dependencies of target cephfs
[ 67%] Building CXX object src/CMakeFiles/cephfs.dir/libcephfs.cc.o
[ 67%] Linking CXX shared library ../lib/libcephfs.so
[ 67%] Built target cephfs
Scanning dependencies of target dmclock
[ 67%] Building CXX object src/dmclock/src/CMakeFiles/dmclock.dir/dmclock_util.cc.o
[ 67%] Building CXX object src/dmclock/src/CMakeFiles/dmclock.dir/__/support/src/run_every.cc.o
[ 67%] Linking CXX static library ../../../lib/libdmclock.a
[ 67%] Built target dmclock
Scanning dependencies of target cls_references_objs
[ 67%] Building CXX object src/CMakeFiles/cls_references_objs.dir/objclass/class_api.cc.o
[ 67%] Built target cls_references_objs
Scanning dependencies of target osd
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/OSD.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/Watch.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/ClassHandler.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/PG.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/PGLog.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/PrimaryLogPG.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/ReplicatedBackend.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/ECBackend.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/ECTransaction.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/PGBackend.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/OSDCap.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/Session.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/SnapMapper.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/ScrubStore.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/osd_types.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/ECUtil.cc.o
[ 67%] Building CXX object src/osd/CMakeFiles/osd.dir/ExtentCache.cc.o
[ 69%] Building CXX object src/osd/CMakeFiles/osd.dir/mClockOpClassQueue.cc.o
[ 69%] Building CXX object src/osd/CMakeFiles/osd.dir/mClockClientQueue.cc.o
[ 69%] Building CXX object src/osd/CMakeFiles/osd.dir/PGQueueable.cc.o
[ 69%] Building CXX object src/osd/CMakeFiles/osd.dir/__/common/TrackedOp.cc.o
[ 69%] Building CXX object src/osd/CMakeFiles/osd.dir/__/osdc/Objecter.cc.o
[ 69%] Building CXX object src/osd/CMakeFiles/osd.dir/__/osdc/Striper.cc.o
[ 69%] Linking CXX static library ../../lib/libosd.a
[ 69%] Built target osd
Scanning dependencies of target ceph-osd
[ 69%] Building CXX object src/CMakeFiles/ceph-osd.dir/ceph_osd.cc.o
[ 69%] Linking CXX executable ../bin/ceph-osd
[ 69%] Built target ceph-osd
Scanning dependencies of target cls_rbd
[ 71%] Building CXX object src/cls/CMakeFiles/cls_rbd.dir/rbd/cls_rbd.cc.o
[ 71%] Building CXX object src/cls/CMakeFiles/cls_rbd.dir/rbd/cls_rbd_types.cc.o
[ 71%] Linking CXX shared library ../../lib/libcls_rbd.so
[ 71%] Built target cls_rbd
Scanning dependencies of target monmaptool
[ 73%] Building CXX object src/tools/CMakeFiles/monmaptool.dir/monmaptool.cc.o
[ 73%] Linking CXX executable ../../bin/monmaptool
[ 73%] Built target monmaptool
Scanning dependencies of target crushtool
[ 73%] Building CXX object src/tools/CMakeFiles/crushtool.dir/crushtool.cc.o
[ 73%] Linking CXX executable ../../bin/crushtool
[ 73%] Built target crushtool
Scanning dependencies of target ceph-authtool
[ 73%] Building CXX object src/tools/CMakeFiles/ceph-authtool.dir/ceph_authtool.cc.o
[ 73%] Linking CXX executable ../../bin/ceph-authtool
[ 73%] Built target ceph-authtool
Scanning dependencies of target radosstriper
[ 75%] Building CXX object src/libradosstriper/CMakeFiles/radosstriper.dir/libradosstriper.cc.o
[ 75%] Building CXX object src/libradosstriper/CMakeFiles/radosstriper.dir/RadosStriperImpl.cc.o
/build/ceph/src/libradosstriper/RadosStriperImpl.cc: In constructor '{anonymous}::WriteCompletionData::WriteCompletionData(libradosstriper::RadosStriperImpl*, const string&, const string&, librados::AioCompletionImpl*, int)':
/build/ceph/src/libradosstriper/RadosStriperImpl.cc:253:28: warning: '{anonymous}::WriteCompletionData::m_unlockCompletion' will be initialized after [-Wreorder]
   librados::AioCompletion *m_unlockCompletion;
                            ^~~~~~~~~~~~~~~~~~
/build/ceph/src/libradosstriper/RadosStriperImpl.cc:251:7: warning:   'int {anonymous}::WriteCompletionData::m_writeRc' [-Wreorder]
   int m_writeRc;
       ^~~~~~~~~
/build/ceph/src/libradosstriper/RadosStriperImpl.cc:270:1: warning:   when initialized here [-Wreorder]
 WriteCompletionData::WriteCompletionData
 ^~~~~~~~~~~~~~~~~~~
[ 75%] Building CXX object src/libradosstriper/CMakeFiles/radosstriper.dir/MultiAioCompletionImpl.cc.o
[ 75%] Linking CXX shared library ../../lib/libradosstriper.so
[ 75%] Built target radosstriper
Scanning dependencies of target rados
[ 75%] Building CXX object src/tools/CMakeFiles/rados.dir/rados/rados.cc.o
[ 75%] Building CXX object src/tools/CMakeFiles/rados.dir/RadosDump.cc.o
[ 77%] Building CXX object src/tools/CMakeFiles/rados.dir/rados/RadosImport.cc.o
[ 77%] Building CXX object src/tools/CMakeFiles/rados.dir/rados/PoolDump.cc.o
[ 77%] Building CXX object src/tools/CMakeFiles/rados.dir/__/common/util.cc.o
[ 77%] Building CXX object src/tools/CMakeFiles/rados.dir/__/common/obj_bencher.cc.o
[ 77%] Linking CXX executable ../../bin/rados
[ 77%] Built target rados
Scanning dependencies of target rbd_types
[ 77%] Building CXX object src/librbd/CMakeFiles/rbd_types.dir/journal/Types.cc.o
[ 77%] Building CXX object src/librbd/CMakeFiles/rbd_types.dir/mirroring_watcher/Types.cc.o
[ 77%] Building CXX object src/librbd/CMakeFiles/rbd_types.dir/watcher/Types.cc.o
[ 77%] Building CXX object src/librbd/CMakeFiles/rbd_types.dir/WatchNotifyTypes.cc.o
[ 77%] Linking CXX static library ../../lib/librbd_types.a
[ 77%] Built target rbd_types
Scanning dependencies of target parse_secret_objs
[ 77%] Building C object src/CMakeFiles/parse_secret_objs.dir/common/secret.c.o
[ 77%] Built target parse_secret_objs
Scanning dependencies of target krbd
[ 77%] Building CXX object src/CMakeFiles/krbd.dir/krbd.cc.o
[ 77%] Linking CXX static library ../lib/libkrbd.a
[ 77%] Built target krbd
Scanning dependencies of target cls_journal_client
[ 77%] Building CXX object src/cls/CMakeFiles/cls_journal_client.dir/journal/cls_journal_client.cc.o
[ 77%] Building CXX object src/cls/CMakeFiles/cls_journal_client.dir/journal/cls_journal_types.cc.o
[ 77%] Linking CXX static library ../../lib/libcls_journal_client.a
[ 77%] Built target cls_journal_client
Scanning dependencies of target cls_rbd_client
[ 77%] Building CXX object src/cls/CMakeFiles/cls_rbd_client.dir/rbd/cls_rbd_client.cc.o
[ 77%] Building CXX object src/cls/CMakeFiles/cls_rbd_client.dir/rbd/cls_rbd_types.cc.o
[ 77%] Linking CXX static library ../../lib/libcls_rbd_client.a
[ 77%] Built target cls_rbd_client
Scanning dependencies of target journal
[ 77%] Building CXX object src/journal/CMakeFiles/journal.dir/Entry.cc.o
[ 77%] Building CXX object src/journal/CMakeFiles/journal.dir/Future.cc.o
[ 77%] Building CXX object src/journal/CMakeFiles/journal.dir/FutureImpl.cc.o
[ 77%] Building CXX object src/journal/CMakeFiles/journal.dir/Journaler.cc.o
[ 77%] Building CXX object src/journal/CMakeFiles/journal.dir/JournalMetadata.cc.o
[ 77%] Building CXX object src/journal/CMakeFiles/journal.dir/JournalPlayer.cc.o
[ 77%] Building CXX object src/journal/CMakeFiles/journal.dir/JournalRecorder.cc.o
[ 77%] Building CXX object src/journal/CMakeFiles/journal.dir/JournalTrimmer.cc.o
[ 79%] Building CXX object src/journal/CMakeFiles/journal.dir/ObjectPlayer.cc.o
[ 79%] Building CXX object src/journal/CMakeFiles/journal.dir/ObjectRecorder.cc.o
[ 79%] Building CXX object src/journal/CMakeFiles/journal.dir/Utils.cc.o
[ 79%] Linking CXX static library ../../lib/libjournal.a
[ 79%] Built target journal
Scanning dependencies of target rados_snap_set_diff_obj
[ 79%] Building CXX object src/CMakeFiles/rados_snap_set_diff_obj.dir/librados/snap_set_diff.cc.o
[ 79%] Built target rados_snap_set_diff_obj
Scanning dependencies of target rbd_internal
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/AsyncObjectThrottle.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/AsyncRequest.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/ExclusiveLock.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/ImageCtx.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/ImageState.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/ImageWatcher.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/internal.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/Journal.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/LibrbdAdminSocketHook.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/LibrbdWriteback.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/ManagedLock.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/MirroringWatcher.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/ObjectMap.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/Operations.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/Utils.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/Watcher.cc.o
[ 81%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/api/DiffIterate.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/api/Image.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/api/Mirror.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/cache/ImageWriteback.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/cache/PassthroughImageCache.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/exclusive_lock/AutomaticPolicy.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/exclusive_lock/PreAcquireRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/exclusive_lock/PostAcquireRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/exclusive_lock/PreReleaseRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/exclusive_lock/StandardPolicy.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image/CloneRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image/CloseRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image/CreateRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image/OpenRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image/RefreshParentRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image/RefreshRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image/RemoveRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image/SetFlagsRequest.cc.o
[ 83%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image/SetSnapRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/image_watcher/NotifyLockOwner.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/io/AioCompletion.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/io/AsyncOperation.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/io/CopyupRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/io/ImageRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/io/ImageRequestWQ.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/io/ObjectRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/io/ReadResult.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/journal/CreateRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/journal/DemoteRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/journal/OpenRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/journal/PromoteRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/journal/RemoveRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/journal/Replay.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/journal/StandardPolicy.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/journal/Utils.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/managed_lock/AcquireRequest.cc.o
[ 84%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/managed_lock/BreakRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/managed_lock/GetLockerRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/managed_lock/ReacquireRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/managed_lock/ReleaseRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/managed_lock/Utils.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/mirror/DemoteRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/mirror/DisableRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/mirror/EnableRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/mirror/GetInfoRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/mirror/GetStatusRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/mirror/PromoteRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/CreateRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/InvalidateRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/LockRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/RefreshRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/RemoveRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/Request.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/ResizeRequest.cc.o
[ 86%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/SnapshotCreateRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/SnapshotRemoveRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/SnapshotRollbackRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/UnlockRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/object_map/UpdateRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/DisableFeaturesRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/EnableFeaturesRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/FlattenRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/MetadataRemoveRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/MetadataSetRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/ObjectMapIterate.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/RebuildObjectMapRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/RenameRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/Request.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/ResizeRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/SnapshotCreateRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/SnapshotProtectRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/SnapshotRemoveRequest.cc.o
[ 88%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/SnapshotRenameRequest.cc.o
[ 90%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/SnapshotRollbackRequest.cc.o
[ 90%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/SnapshotUnprotectRequest.cc.o
[ 90%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/SnapshotLimitRequest.cc.o
[ 90%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/operation/TrimRequest.cc.o
[ 90%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/watcher/Notifier.cc.o
[ 90%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/watcher/RewatchRequest.cc.o
[ 90%] Building CXX object src/librbd/CMakeFiles/rbd_internal.dir/__/common/ContextCompletion.cc.o
[ 90%] Linking CXX static library ../../lib/librbd_internal.a
[ 90%] Built target rbd_internal
Scanning dependencies of target librbd
[ 90%] Building CXX object src/librbd/CMakeFiles/librbd.dir/librbd.cc.o
[ 90%] Linking CXX shared library ../../lib/librbd.so
[ 90%] Built target librbd
Scanning dependencies of target rbd
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/rbd.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/ArgumentTypes.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/IndentStream.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/OptionPrinter.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/Shell.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/Utils.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Bench.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Children.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Clone.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Copy.cc.o
[ 90%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Create.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Diff.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/DiskUsage.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Export.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Feature.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Flatten.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/ImageMeta.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Import.cc.o
/build/ceph/src/tools/rbd/action/Import.cc: In function 'int rbd::action::import::execute(const boost::program_options::variables_map&)':
/build/ceph/src/tools/rbd/action/Import.cc:600:31: warning: 'length' may be used uninitialized in this function [-Wmaybe-uninitialized]
       r = skip_tag(fd, length);
                               ^
/build/ceph/src/tools/rbd/action/Import.cc:583:14: note: 'length' was declared here
     uint64_t length;
              ^~~~~~
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Info.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Journal.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Kernel.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/List.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Lock.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/MergeDiff.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/MirrorPool.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/MirrorImage.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Nbd.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/ObjectMap.cc.o
[ 92%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Pool.cc.o
[ 94%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Remove.cc.o
[ 94%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Rename.cc.o
[ 94%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Resize.cc.o
[ 94%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Snap.cc.o
[ 94%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Status.cc.o
[ 94%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Trash.cc.o
[ 94%] Building CXX object src/tools/rbd/CMakeFiles/rbd.dir/action/Watch.cc.o
[ 94%] Linking CXX executable ../../../bin/rbd
[ 94%] Built target rbd
Scanning dependencies of target civetweb_common_objs
[ 94%] Building C object src/CMakeFiles/civetweb_common_objs.dir/civetweb/src/civetweb.c.o
[ 94%] Built target civetweb_common_objs
Scanning dependencies of target radosgw_a
[ 94%] Building CXX object src/rgw/CMakeFiles/radosgw_a.dir/rgw_loadgen_process.cc.o
[ 94%] Building CXX object src/rgw/CMakeFiles/radosgw_a.dir/rgw_civetweb.cc.o
[ 94%] Building CXX object src/rgw/CMakeFiles/radosgw_a.dir/rgw_civetweb_frontend.cc.o
[ 94%] Building CXX object src/rgw/CMakeFiles/radosgw_a.dir/rgw_civetweb_log.cc.o
[ 94%] Building CXX object src/rgw/CMakeFiles/radosgw_a.dir/rgw_asio_client.cc.o
[ 94%] Building CXX object src/rgw/CMakeFiles/radosgw_a.dir/rgw_asio_frontend.cc.o
In file included from /usr/include/boost/coroutine/standard_stack_allocator.hpp:21:0,
                 from /usr/include/boost/coroutine/stack_allocator.hpp:16,
                 from /usr/include/boost/coroutine/attributes.hpp:15,
                 from /usr/include/boost/coroutine/all.hpp:10,
                 from /usr/include/boost/asio/spawn.hpp:19,
                 from /build/ceph/src/rgw/rgw_asio_frontend.cc:10:
/usr/include/boost/coroutine/detail/config.hpp:17:4: warning: #warning "Boost.Coroutine is now deprecated. Please switch to Boost.Coroutine2. To disable this warning message, define BOOST_COROUTINE_NO_DEPRECATION_WARNING." [-Wcpp]
 #  warning                  "Boost.Coroutine is now deprecated. Please switch to Boost.Coroutine2. To disable this warning message, define BOOST_COROUTINE_NO_DEPRECATION_WARNING."
    ^~~~~~~
[ 94%] Linking CXX static library ../../lib/libradosgw_a.a
[ 94%] Built target radosgw_a
Scanning dependencies of target cls_rgw
[ 94%] Building CXX object src/cls/CMakeFiles/cls_rgw.dir/rgw/cls_rgw.cc.o
[ 94%] Building CXX object src/cls/CMakeFiles/cls_rgw.dir/rgw/cls_rgw_ops.cc.o
[ 94%] Building CXX object src/cls/CMakeFiles/cls_rgw.dir/rgw/cls_rgw_types.cc.o
[ 94%] Building CXX object src/cls/CMakeFiles/cls_rgw.dir/__/common/ceph_json.cc.o
[ 94%] Linking CXX shared library ../../lib/libcls_rgw.so
[ 94%] Built target cls_rgw
Scanning dependencies of target cls_refcount
[ 94%] Building CXX object src/cls/CMakeFiles/cls_refcount.dir/refcount/cls_refcount.cc.o
[ 94%] Building CXX object src/cls/CMakeFiles/cls_refcount.dir/refcount/cls_refcount_ops.cc.o
[ 94%] Building CXX object src/cls/CMakeFiles/cls_refcount.dir/__/common/ceph_json.cc.o
[ 94%] Linking CXX shared library ../../lib/libcls_refcount.so
[ 94%] Built target cls_refcount
Scanning dependencies of target cls_user
[ 94%] Building CXX object src/cls/CMakeFiles/cls_user.dir/user/cls_user.cc.o
[ 94%] Linking CXX shared library ../../lib/libcls_user.so
[ 94%] Built target cls_user
Scanning dependencies of target cls_lock
[ 94%] Building CXX object src/cls/CMakeFiles/cls_lock.dir/lock/cls_lock.cc.o
[ 94%] Linking CXX shared library ../../lib/libcls_lock.so
[ 94%] Built target cls_lock
Scanning dependencies of target cls_replica_log
[ 94%] Building CXX object src/cls/CMakeFiles/cls_replica_log.dir/replica_log/cls_replica_log.cc.o
[ 94%] Linking CXX shared library ../../lib/libcls_replica_log.so
[ 94%] Built target cls_replica_log
Scanning dependencies of target cls_log
[ 94%] Building CXX object src/cls/CMakeFiles/cls_log.dir/log/cls_log.cc.o
[ 94%] Linking CXX shared library ../../lib/libcls_log.so
[ 94%] Built target cls_log
Scanning dependencies of target cls_version
[ 94%] Building CXX object src/cls/CMakeFiles/cls_version.dir/version/cls_version.cc.o
[ 94%] Linking CXX shared library ../../lib/libcls_version.so
[ 94%] Built target cls_version
Scanning dependencies of target cls_timeindex
[ 94%] Building CXX object src/cls/CMakeFiles/cls_timeindex.dir/timeindex/cls_timeindex.cc.o
[ 96%] Linking CXX shared library ../../lib/libcls_timeindex.so
[ 96%] Built target cls_timeindex
Scanning dependencies of target cls_statelog
[ 96%] Building CXX object src/cls/CMakeFiles/cls_statelog.dir/statelog/cls_statelog.cc.o
[ 96%] Linking CXX shared library ../../lib/libcls_statelog.so
[ 96%] Built target cls_statelog
Scanning dependencies of target radosgw
[ 98%] Building CXX object src/rgw/CMakeFiles/radosgw.dir/rgw_main.cc.o
[ 98%] Linking CXX executable ../../bin/radosgw
[ 98%] Built target radosgw
Scanning dependencies of target rgw
[100%] Building CXX object src/rgw/CMakeFiles/rgw.dir/librgw.cc.o
[100%] Building CXX object src/rgw/CMakeFiles/rgw.dir/rgw_file.cc.o
[100%] Linking CXX shared library ../../lib/librgw.so
[100%] Built target rgw
Scanning dependencies of target rook
[100%] Built target rook
Install the project...
-- Install configuration: "release"
-- Installing: /build/ceph-install/usr/local/lib/ceph/libceph-common.so.0
-- Installing: /build/ceph-install/usr/local/lib/ceph/libceph-common.so
-- Installing: /build/ceph-install/usr/local/bin/ceph-mgr
-- Set runtime path of "/build/ceph-install/usr/local/bin/ceph-mgr" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/ceph-mon
-- Set runtime path of "/build/ceph-install/usr/local/bin/ceph-mon" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/ceph-osd
-- Set runtime path of "/build/ceph-install/usr/local/bin/ceph-osd" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/ceph-mds
-- Set runtime path of "/build/ceph-install/usr/local/bin/ceph-mds" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/ceph-debugpack
-- Installing: /build/ceph-install/usr/local/bin/ceph-coverage
-- Installing: /build/ceph-install/usr/local/bin/ceph
-- Installing: /build/ceph-install/usr/local/bin/ceph-crush-location
-- Installing: /build/ceph-install/usr/local/bin/ceph-post-file
-- Installing: /build/ceph-install/usr/local/bin/ceph-run
-- Installing: /build/ceph-install/usr/local/bin/ceph-rest-api
-- Installing: /build/ceph-install/usr/local/bin/ceph-clsinfo
-- Installing: /build/ceph-install/usr/local/etc/init.d/ceph
-- Installing: /build/ceph-install/usr/local/share/ceph/id_rsa_drop.ceph.com
-- Installing: /build/ceph-install/usr/local/share/ceph/id_rsa_drop.ceph.com.pub
-- Installing: /build/ceph-install/usr/local/share/ceph/known_hosts_drop.ceph.com
-- Installing: /build/ceph-install/usr/local/libexec/ceph/ceph_common.sh
-- Installing: /build/ceph-install/usr/local/libexec/ceph/ceph-osd-prestart.sh
-- Installing: /build/ceph-install/usr/local/sbin/ceph-create-keys
-- Installing: /build/ceph-install/usr/local/lib/libcephfs.so.2.0.0
-- Installing: /build/ceph-install/usr/local/lib/libcephfs.so.2
-- Installing: /build/ceph-install/usr/local/lib/libcephfs.so
-- Set runtime path of "/build/ceph-install/usr/local/lib/libcephfs.so.2.0.0" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/include/cephfs
-- Installing: /build/ceph-install/usr/local/include/cephfs/libcephfs.h
-- Installing: /build/ceph-install/usr/local/include/cephfs/ceph_statx.h
-- Installing: /build/ceph-install/usr/local/bin/ceph-rbdnamer
-- Installing: /build/ceph-install/usr/local/bin/rbd-replay-many
-- Installing: /build/ceph-install/usr/local/bin/rbdmap
-- Installing: /build/ceph-install/usr/local/share/doc/ceph/sample.ceph.conf
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_rbd.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_rbd.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_rbd.so
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_lock.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_lock.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_lock.so
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_refcount.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_refcount.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_refcount.so
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_version.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_version.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_version.so
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_log.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_log.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_log.so
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_statelog.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_statelog.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_statelog.so
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_timeindex.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_timeindex.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_timeindex.so
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_replica_log.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_replica_log.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_replica_log.so
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_user.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_user.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_user.so
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_rgw.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_rgw.so.1
-- Installing: /build/ceph-install/usr/local/lib/rados-classes/libcls_rgw.so
-- Installing: /build/ceph-install/usr/local/include/rados/librados.h
-- Installing: /build/ceph-install/usr/local/include/rados/rados_types.h
-- Installing: /build/ceph-install/usr/local/include/rados/rados_types.hpp
-- Installing: /build/ceph-install/usr/local/include/rados/librados.hpp
-- Installing: /build/ceph-install/usr/local/include/rados/buffer.h
-- Installing: /build/ceph-install/usr/local/include/rados/buffer_fwd.h
-- Installing: /build/ceph-install/usr/local/include/rados/inline_memory.h
-- Installing: /build/ceph-install/usr/local/include/rados/memory.h
-- Installing: /build/ceph-install/usr/local/include/rados/page.h
-- Installing: /build/ceph-install/usr/local/include/rados/crc32c.h
-- Installing: /build/ceph-install/usr/local/include/rados/objclass.h
-- Installing: /build/ceph-install/usr/local/include/radosstriper/libradosstriper.h
-- Installing: /build/ceph-install/usr/local/include/radosstriper/libradosstriper.hpp
-- Installing: /build/ceph-install/usr/local/include/rbd/features.h
-- Installing: /build/ceph-install/usr/local/include/rbd/librbd.h
-- Installing: /build/ceph-install/usr/local/include/rbd/librbd.hpp
-- Installing: /build/ceph-install/usr/local/include/rados/librgw.h
-- Installing: /build/ceph-install/usr/local/include/rados/rgw_file.h
-- Installing: /build/ceph-install/usr/local/lib/librados.so.2.0.0
-- Installing: /build/ceph-install/usr/local/lib/librados.so.2
-- Installing: /build/ceph-install/usr/local/lib/librados.so
-- Set runtime path of "/build/ceph-install/usr/local/lib/librados.so.2.0.0" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/lib/libradosstriper.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/libradosstriper.so.1
-- Installing: /build/ceph-install/usr/local/lib/libradosstriper.so
-- Set runtime path of "/build/ceph-install/usr/local/lib/libradosstriper.so.1.0.0" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/lib/python2.7/dist-packages/ceph_argparse.py
-- Installing: /build/ceph-install/usr/local/lib/python2.7/dist-packages/ceph_daemon.py
-- Installing: /build/ceph-install/usr/local/lib/python2.7/dist-packages/ceph_volume_client.py
-- Installing: /build/ceph-install/usr/local/lib/python2.7/dist-packages/ceph_rest_api.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/influx
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/influx/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/influx/module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/balancer
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/balancer/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/balancer/module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/status
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/status/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/status/module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/decorators.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api/config.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api/crush.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api/server.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api/request.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api/osd.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api/mon.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api/doc.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/api/pool.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/hooks.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/restful/common.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/localpool
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/localpool/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/localpool/module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/mgr_module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/selftest
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/selftest/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/selftest/module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/prometheus
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/prometheus/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/prometheus/module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/cephfs_clients.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/clients.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/osds.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/HACKING.rst
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/rbd_iscsi.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/rbd_pool.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/README.rst
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/remote_view_cache.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/base.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/osd_perf.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/rbd_iscsi.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/servers.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/standby.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/filesystem.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/rbd_ls.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/rbd_mirroring.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/.jshintrc
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/LICENSE
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ionslider
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ionslider/ion.rangeSlider.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ionslider/img
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ionslider/img/sprite-skin-flat.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ionslider/img/sprite-skin-nice.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ionslider/ion.rangeSlider.skinNice.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ionslider/ion.rangeSlider.skinFlat.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ionslider/ion.rangeSlider.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/sparkline
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/sparkline/jquery.sparkline.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/sparkline/jquery.sparkline.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/chartjs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/chartjs/Chart.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/chartjs/Chart.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jQuery
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jQuery/jquery-2.2.3.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/jquery.dataTables.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/images
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/images/sort_desc_disabled.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/images/sort_asc_disabled.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/images/sort_desc.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/images/sort_both.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/images/sort_asc.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/jquery.dataTables_themeroller.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/dataTables.bootstrap.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/jquery.dataTables.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/jquery.dataTables.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/dataTables.bootstrap.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/jquery.dataTables.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/License.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/child-rows
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/child-rows/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/child-rows/custom-renderer.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/child-rows/column-control.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/child-rows/right-column.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/child-rows/disable-child-rows.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/child-rows/whole-row-control.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/display-control
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/display-control/fixedHeader.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/display-control/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/display-control/auto.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/display-control/complexHeader.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/display-control/classes.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/display-control/init-classes.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/initialisation
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/initialisation/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/initialisation/option.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/initialisation/className.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/initialisation/new.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/initialisation/default.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/initialisation/ajax.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/styling
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/styling/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/styling/foundation.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/styling/scrolling.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/styling/bootstrap.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/examples/styling/compact.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/Readme.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/css/dataTables.responsive.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/css/dataTables.responsive.scss
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/js/dataTables.responsive.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Responsive/js/dataTables.responsive.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/examples/events.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/examples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/examples/scrolling.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/examples/simple.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/examples/html.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/Readme.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/css/dataTables.keyTable.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/css/dataTables.keyTable.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/js/dataTables.keyTable.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/KeyTable/js/dataTables.keyTable.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/swf
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/swf/copy_csv_xls.swf
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/swf/copy_csv_xls_pdf.swf
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/alter_buttons.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/pdf_message.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/swf_path.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/multiple_tables.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/new_init.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/select_single.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/bootstrap.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/simple.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/select_os.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/multi_instance.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/button_text.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/plug-in.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/ajax.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/select_column.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/defaults.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/jqueryui.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/collection.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/examples/select_multi.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/csv_hover.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/csv.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/collection.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/xls_hover.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/psd
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/psd/collection.psd
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/psd/file_types.psd
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/psd/printer.psd
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/psd/copy document.psd
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/copy.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/print.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/print_hover.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/pdf.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/pdf_hover.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/xls.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/collection_hover.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/images/copy_hover.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/Readme.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/css/dataTables.tableTools.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/css/dataTables.tableTools.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/js/dataTables.tableTools.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/TableTools/js/dataTables.tableTools.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/examples/zIndexes.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/examples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/examples/simple.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/examples/two_tables.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/examples/header_footer.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/examples/top_left_right.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/Readme.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/css/dataTables.fixedHeader.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/css/dataTables.fixedHeader.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/js/dataTables.fixedHeader.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedHeader/js/dataTables.fixedHeader.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/examples/fill-horizontal.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/examples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/examples/scrolling.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/examples/step-callback.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/examples/simple.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/examples/columns.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/examples/complete-callback.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/examples/fill-both.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/Readme.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/images
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/images/filler.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/css/dataTables.autoFill.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/css/dataTables.autoFill.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/js/dataTables.autoFill.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/AutoFill/js/dataTables.autoFill.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/License.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/restore.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/new_init.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/two_tables_identical.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/exclude_columns.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/button_order.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/text.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/simple.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/title_callback.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/two_tables.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/mouseover.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/jqueryui.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/examples/group_columns.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/Readme.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/css/dataTables.colVis.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/css/dataTables.colVis.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/css/dataTables.colvis.jqueryui.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/js/dataTables.colVis.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColVis/js/dataTables.colVis.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/License.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/server_side.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/fixedheader.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/new_init.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/fixedcolumns.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/scrolling.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/reset.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/simple.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/alt_insert.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/jqueryui.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/predefined.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/col_filter.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/state_save.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/colvis.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/examples/realtime.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/images
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/images/insert.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/Readme.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/css/dataTables.colReorder.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/css/dataTables.colReorder.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/js/dataTables.colReorder.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/ColReorder/js/dataTables.colReorder.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples/data
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples/data/ssp.php
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples/data/2500.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples/large_js_source.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples/server-side_processing.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples/api_scrolling.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples/simple.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/examples/state_saving.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/Readme.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/images
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/images/loading-background.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/css/dataTables.scroller.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/css/dataTables.scroller.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/js/dataTables.scroller.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/Scroller/js/dataTables.scroller.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/License.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/two_columns.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/size_fluid.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/size_fixed.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/bootstrap.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/left_right_columns.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/index_column.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/server-side-processing.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/css_size.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/simple.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/right_column.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/col_filter.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/colvis.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/examples/rowspan.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/Readme.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/css/dataTables.fixedColumns.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/css/dataTables.fixedColumns.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/js/dataTables.fixedColumns.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/extensions/FixedColumns/js/dataTables.fixedColumns.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datatables/dataTables.bootstrap.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/img
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/img/hue.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/img/alpha-horizontal.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/img/saturation.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/img/hue-horizontal.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/img/alpha.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/bootstrap-colorpicker.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/bootstrap-colorpicker.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/bootstrap-colorpicker.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/colorpicker/bootstrap-colorpicker.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/fullcalendar
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/fullcalendar/fullcalendar.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/fullcalendar/fullcalendar.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/fullcalendar/fullcalendar.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/fullcalendar/fullcalendar.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/fullcalendar/fullcalendar.print.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/bootstrap-datepicker.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.pt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.mk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.ca.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.zh-TW.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.ja.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.is.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.rs.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.cy.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.vi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.nl-BE.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.sw.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.el.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.cs.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.ua.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.sv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.nb.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.he.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.sq.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.az.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.ka.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.ms.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.th.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.lv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.gl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.no.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.ar.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.kk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.bg.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.tr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.sk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.sl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.de.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.fr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.lt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.nl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.hu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.hr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.da.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.pt-BR.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.rs-latin.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.id.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.et.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.it.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.es.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.fa.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.ru.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.zh-CN.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.ro.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.fi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.kr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/locales/bootstrap-datepicker.pl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/datepicker/datepicker3.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/timepicker
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/timepicker/bootstrap-timepicker.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/timepicker/bootstrap-timepicker.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/timepicker/bootstrap-timepicker.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/timepicker/bootstrap-timepicker.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/morris
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/morris/morris.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/morris/morris.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/morris/morris.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/select2.full.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/ro.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/et.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/id.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/vi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/ms.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/sv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/he.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/eu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/km.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/ca.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/ar.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/fr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/sr-Cyrl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/nb.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/ru.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/tr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/mk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/hi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/sk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/ko.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/nl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/hu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/el.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/ja.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/pt-BR.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/en.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/fa.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/zh-TW.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/es.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/hr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/cs.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/pt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/lt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/gl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/is.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/lv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/th.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/bg.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/pl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/it.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/de.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/da.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/sr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/fi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/az.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/zh-CN.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/i18n/uk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/select2.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/select2.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/select2.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/select2.full.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/select2/select2.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/pace
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/pace/pace.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/pace/pace.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/pace/pace.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/pace/pace.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jQueryUI
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jQueryUI/jquery-ui.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jQueryUI/jquery-ui.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/fastclick
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/fastclick/fastclick.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/fastclick/fastclick.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/bootstrap-wysihtml5
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/bootstrap-wysihtml5/bootstrap3-wysihtml5.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/bootstrap-wysihtml5/bootstrap3-wysihtml5.all.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/bootstrap-wysihtml5/bootstrap3-wysihtml5.all.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/bootstrap-wysihtml5/bootstrap3-wysihtml5.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/ckeditor.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/CHANGES.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/styles.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/adapters
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/adapters/jquery.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/icons_hidpi.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/tabletools
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/tabletools/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/tabletools/dialogs/tableCell.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/a11yhelp.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/ug.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/eo.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/fr-ca.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/ro.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/et.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/id.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/en-gb.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/vi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/sv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/he.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/eu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/km.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/ca.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/ar.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/fr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/pt-br.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/nb.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/ru.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/tr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/mk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/gu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/hi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/sk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/no.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/fo.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/ko.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/nl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/hu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/sr-latn.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/cy.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/el.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/ja.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/en.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/zh.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/fa.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/si.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/es.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/hr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/af.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/cs.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/zh-cn.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/sl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/tt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/pt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/lt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/gl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/lv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/th.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/_translationstatus.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/bg.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/pl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/it.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/de.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/da.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/de-ch.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/sr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/fi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/sq.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/mn.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/ku.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/a11yhelp/dialogs/lang/uk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/table
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/table/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/table/dialogs/table.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/link
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/link/images
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/link/images/hidpi
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/link/images/hidpi/anchor.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/link/images/anchor.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/link/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/link/dialogs/anchor.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/link/dialogs/link.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/dialog
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/dialog/dialogDefinition.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/image
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/image/images
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/image/images/noimage.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/image/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/image/dialogs/image.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/about
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/about/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/about/dialogs/hidpi
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/about/dialogs/hidpi/logo_ckeditor.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/about/dialogs/logo_ckeditor.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/about/dialogs/about.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/icons.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/specialchar.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/ug.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/eo.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/fr-ca.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/et.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/id.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/en-gb.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/vi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/sv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/he.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/eu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/km.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/ca.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/ar.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/fr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/pt-br.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/nb.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/ru.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/tr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/sk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/no.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/ko.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/nl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/hu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/cy.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/el.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/ja.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/en.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/zh.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/fa.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/si.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/es.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/hr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/af.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/cs.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/zh-cn.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/sl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/tt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/pt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/lt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/gl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/lv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/th.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/_translationstatus.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/bg.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/pl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/it.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/de.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/da.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/de-ch.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/fi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/sq.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/ku.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/specialchar/dialogs/lang/uk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/wsc
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/wsc/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/wsc/dialogs/wsc_ie.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/wsc/dialogs/ciframe.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/wsc/dialogs/wsc.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/wsc/dialogs/wsc.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/wsc/dialogs/tmpFrameset.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/wsc/LICENSE.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/wsc/README.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/scayt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/scayt/CHANGELOG.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/scayt/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/scayt/dialogs/options.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/scayt/dialogs/toolbar.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/scayt/LICENSE.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/scayt/README.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/clipboard
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/clipboard/dialogs
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/clipboard/dialogs/paste.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/magicline
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/magicline/images
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/magicline/images/icon-rtl.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/magicline/images/hidpi
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/magicline/images/hidpi/icon-rtl.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/magicline/images/hidpi/icon.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/magicline/images/icon.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/pastefromword
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/pastefromword/filter
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/plugins/pastefromword/filter/default.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/lib
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/lib/codemirror
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/lib/codemirror/neo.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/lib/codemirror/show-hint.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/lib/codemirror/LICENSE
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/lib/codemirror/show-hint.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/lib/codemirror/javascript.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/lib/codemirror/codemirror.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/lib/codemirror/codemirror.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/font
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/font/fontello.woff
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/font/LICENSE.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/font/fontello.ttf
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/font/config.json
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/font/fontello.eot
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/font/fontello.svg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/css/fontello.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/js/toolbarmodifier.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/js/fulltoolbareditor.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/js/abstracttoolbarmodifier.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/toolbarconfigurator/js/toolbartextmodifier.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/img
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/img/header-separator.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/img/logo.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/img/github-top.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/img/navigation-tip.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/img/header-bg.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/sample.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/htmlwriter
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/htmlwriter/outputhtml.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/htmlwriter/assets
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/htmlwriter/assets/outputforflash
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/htmlwriter/assets/outputforflash/swfobject.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/htmlwriter/assets/outputforflash/outputforflash.swf
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/htmlwriter/assets/outputforflash/outputforflash.fla
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/htmlwriter/outputforflash.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/enterkey
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/enterkey/enterkey.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/tabindex.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/inlineall.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/datafiltering.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/toolbar
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/toolbar/toolbar.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/uilanguages.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/jquery.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/dialog
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/dialog/dialog.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/dialog/assets
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/dialog/assets/my_dialog.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/replacebyclass.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/appendto.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/sample_posteddata.php
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/wysiwygarea
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/wysiwygarea/fullpage.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/uicolor.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/inlinebycode.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/sample.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/assets
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/assets/posteddata.php
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/assets/outputxhtml
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/assets/outputxhtml/outputxhtml.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/assets/uilanguages
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/assets/uilanguages/languages.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/assets/sample.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/assets/inlineall
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/assets/inlineall/logo.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/api.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/ajax.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/replacebycode.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/inlinetextarea.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/magicline
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/magicline/magicline.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/readonly.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/xhtmlstyle.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/old/divreplace.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/css/samples.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/js/sample.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/samples/js/sf.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/config.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/editor_iequirks.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/icons_hidpi.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/dialog_iequirks.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/editor_ie8.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/editor_gecko.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/readme.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/editor_ie.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/dialog_ie8.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/editor_ie7.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/dialog_ie7.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/editor.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/dialog.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/lock-open.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/hidpi
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/hidpi/lock-open.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/hidpi/close.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/hidpi/refresh.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/hidpi/lock.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/close.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/refresh.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/lock.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/spinner.gif
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/images/arrow.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/dialog_ie.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/skins/moono/icons.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/LICENSE.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ug.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/bs.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/eo.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/fr-ca.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ro.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/et.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/id.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/en-gb.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/vi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ms.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/sv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/he.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/eu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/km.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ca.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ar.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/fr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/pt-br.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/nb.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ru.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/tr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/mk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/gu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/hi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/sk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/no.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/fo.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ko.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/nl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/hu.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/sr-latn.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/en-ca.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/cy.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/el.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ja.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/en.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/zh.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/fa.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/bn.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/si.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/es.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/hr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/af.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/cs.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/zh-cn.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/sl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/tt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/pt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/en-au.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/lt.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/gl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/is.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/lv.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/th.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/bg.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/pl.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/it.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/de.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/da.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/de-ch.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/sr.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/fi.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/sq.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ka.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/mn.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/ku.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/lang/uk.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/contents.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/build-config.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/ckeditor/README.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/jquery.inputmask.numeric.extensions.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/jquery.inputmask.extensions.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/jquery.inputmask.regex.extensions.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/jquery.inputmask.phone.extensions.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/jquery.inputmask.date.extensions.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/jquery.inputmask.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/phone-codes
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/phone-codes/phone-codes.json
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/phone-codes/phone-be.json
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/input-mask/phone-codes/readme.txt
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jvectormap
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jvectormap/jquery-jvectormap-1.2.2.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jvectormap/jquery-jvectormap-usa-en.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jvectormap/jquery-jvectormap-world-mill-en.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/jvectormap/jquery-jvectormap-1.2.2.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/excanvas.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.categories.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.threshold.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.resize.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.crosshair.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.colorhelpers.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.categories.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.canvas.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.stack.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.colorhelpers.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.crosshair.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.symbol.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.errorbars.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.selection.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.selection.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.fillbetween.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.time.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.navigate.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.threshold.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.fillbetween.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.errorbars.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.image.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.pie.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.image.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.resize.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.canvas.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/excanvas.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.time.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.navigate.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.symbol.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.stack.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/flot/jquery.flot.pie.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/bootstrap-slider
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/bootstrap-slider/bootstrap-slider.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/bootstrap-slider/slider.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/daterangepicker
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/daterangepicker/daterangepicker.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/daterangepicker/moment.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/daterangepicker/daterangepicker.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/daterangepicker/moment.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/knob
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/knob/jquery.knob.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/_all.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/pink.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/line.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/purple.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/green.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/grey.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/blue.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/line.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/yellow.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/red.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/orange.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/aero.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/line/line@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/icheck.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/green@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/red.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/purple@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/grey.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/_all.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/pink.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/pink@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/orange@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/orange.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/flat.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/purple.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/flat@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/green.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/yellow.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/green.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/grey.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/blue.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/flat.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/yellow.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/yellow@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/red.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/orange.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/aero.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/blue.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/pink.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/aero.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/grey@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/red@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/blue@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/purple.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/flat/aero@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/polaris
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/polaris/polaris@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/polaris/polaris.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/polaris/polaris.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/icheck.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/all.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/green@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/red.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/purple@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/grey.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/_all.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/pink.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/pink@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/orange@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/orange.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/purple.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/green.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/yellow.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/square.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/green.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/grey.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/blue.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/yellow.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/yellow@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/red.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/orange.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/aero.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/blue.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/pink.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/aero.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/grey@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/red@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/blue@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/purple.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/aero@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/square@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/square/square.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/futurico
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/futurico/futurico.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/futurico/futurico@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/futurico/futurico.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/green@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/red.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/purple@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/grey.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/_all.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/pink.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/pink@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/minimal.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/orange@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/orange.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/purple.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/green.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/yellow.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/minimal@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/green.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/grey.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/blue.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/yellow.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/yellow@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/red.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/orange.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/aero.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/blue.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/pink.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/aero.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/minimal.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/grey@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/red@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/blue@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/purple.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/iCheck/minimal/aero@2x.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/slimScroll
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/slimScroll/jquery.slimscroll.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/plugins/slimScroll/jquery.slimscroll.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/fonts
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/fonts/glyphicons-halflings-regular.woff2
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/fonts/glyphicons-halflings-regular.woff
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/fonts/glyphicons-halflings-regular.svg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/fonts/glyphicons-halflings-regular.eot
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/fonts/glyphicons-halflings-regular.ttf
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/css/bootstrap.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/css/bootstrap.min.css.map
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/css/bootstrap.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/css/bootstrap.css.map
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/js/bootstrap.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/js/bootstrap.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bootstrap/js/npm.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/boxed-bg.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/default-50x50.gif
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/icons.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/boxed-bg.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/photo3.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/avatar04.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/user4-128x128.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/user3-128x128.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/avatar3.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/user5-128x128.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/avatar2.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/avatar5.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/user8-128x128.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/photo2.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/user6-128x128.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/avatar.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/user7-128x128.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/photo4.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/user1-128x128.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/credit
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/credit/visa.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/credit/american-express.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/credit/cirrus.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/credit/paypal.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/credit/paypal2.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/credit/mastercard.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/credit/mestro.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/photo1.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/img/user2-160x160.jpg
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/AdminLTE.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-red-light.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-blue-light.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-purple.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-blue.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-black-light.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-yellow-light.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-purple-light.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-black.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/_all-skins.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-green.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-yellow.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-red.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-green-light.min.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-black-light.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-purple-light.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-green.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-yellow.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-yellow-light.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-green-light.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/_all-skins.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-black.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-red-light.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-blue.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-purple.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-red.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/skins/skin-blue-light.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/css/AdminLTE.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/js/app.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/js/demo.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/js/app.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/js/pages
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/js/pages/dashboard2.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/dist/js/pages/dashboard.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/.gitignore
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/README.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/bower.json
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/docs.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/style.css
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/download.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/dependencies.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/implementations.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/license.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/components.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/advice.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/plugins.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/layout.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/browsers.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/introduction.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/upgrade.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/faq.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/include/adminlte-options.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/documentation/build/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/index.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/info-box.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/mixins.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/labels.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/login_and_register.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/table.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/modal.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/buttons.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/social-widgets.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/control-sidebar.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/miscellaneous.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/fullcalendar.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/bootstrap-social.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/progress-bars.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/core.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/sidebar.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/timeline.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/navs.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/404_500_errors.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/mailbox.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/profile.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/dropdown.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/direct-chat.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/sidebar-mini.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/carousel.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/invoice.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/lockscreen.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/forms.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/variables.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-red-light.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-green-light.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-black.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-green.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-black-light.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-blue-light.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-blue.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/_all-skins.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-yellow-light.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-purple-light.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-purple.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-red.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/skins/skin-yellow.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/alerts.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/print.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/small-box.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/.csslintrc
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/header.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/callout.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/boxes.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/users-list.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/products.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/select2.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/less/AdminLTE.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/list-group.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/labels.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/image.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/reset-filter.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/opacity.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/buttons.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/nav-divider.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/grid.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/panels.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/reset-text.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/size.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/center-block.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/grid-framework.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/text-emphasis.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/responsive-visibility.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/nav-vertical-align.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/background-variant.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/hide-text.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/forms.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/alerts.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/border-radius.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/resize.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/text-overflow.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/pagination.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/table-row.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/tab-focus.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/clearfix.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/progress-bar.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/gradients.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/mixins/vendor-prefixes.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/build/bootstrap-less/variables.less
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/changelog.md
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/composer.json
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/Gruntfile.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/index2.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/charts
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/charts/flot.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/charts/morris.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/charts/inline.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/charts/chartjs.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/mailbox
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/mailbox/read-mail.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/mailbox/compose.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/mailbox/mailbox.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/tables
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/tables/simple.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/tables/data.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/login.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/invoice.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/register.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/blank.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/500.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/invoice-print.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/profile.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/404.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/pace.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/examples/lockscreen.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/layout
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/layout/fixed.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/layout/boxed.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/layout/top-nav.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/layout/collapsed-sidebar.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/widgets.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/UI
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/UI/modals.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/UI/sliders.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/UI/buttons.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/UI/timeline.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/UI/icons.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/UI/general.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/calendar.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/forms
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/forms/advanced.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/forms/general.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/pages/forms/editors.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/starter.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/AdminLTE-2.3.7/package.json
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/favicon.ico
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/Ceph_Logo_Standard_RGB_White_120411_fa.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/logo-mini.png
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/rivets.bundled.min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/static/underscore-min.js
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/rbd_mirroring.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/health.html
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/dashboard/types.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/.gitignore
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/zabbix
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/zabbix/__init__.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/zabbix/module.py
-- Installing: /build/ceph-install/usr/local/lib/ceph/mgr/zabbix/zabbix_template.xml
running build
running build_ext
cythoning rados.pyx to /build/ceph/build/src/pybind/rados/pyrex/rados.c
creating /build/ceph/build/src/pybind/rados/pyrex
building 'rados' extension
creating /build/ceph/build/src/pybind/rados/build
creating /build/ceph/build/src/pybind/rados/build/ceph
creating /build/ceph/build/src/pybind/rados/build/ceph/build
creating /build/ceph/build/src/pybind/rados/build/ceph/build/src
creating /build/ceph/build/src/pybind/rados/build/ceph/build/src/pybind
creating /build/ceph/build/src/pybind/rados/build/ceph/build/src/pybind/rados
creating /build/ceph/build/src/pybind/rados/build/ceph/build/src/pybind/rados/pyrex
/usr/lib/ccache/x86_64-linux-gnu-gcc -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -fno-strict-aliasing -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -iquote/build/ceph/src/include -fPIC -I/usr/include/python2.7 -I/usr/include/x86_64-linux-gnu/python2.7 -I/usr/include/python2.7 -c /build/ceph/build/src/pybind/rados/pyrex/rados.c -o /build/ceph/build/src/pybind/rados/build/ceph/build/src/pybind/rados/pyrex/rados.o -fno-strict-aliasing -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -lpython2.7 -lpthread -ldl -lutil -lm -Xlinker -export-dynamic -Wl,-O1 -Wl,-Bsymbolic-functions
creating /build/ceph/build/lib/cython_modules
creating /build/ceph/build/lib/cython_modules/lib.2
x86_64-linux-gnu-gcc -pthread -shared -Wl,-O1 -Wl,-Bsymbolic-functions -Wl,-Bsymbolic-functions -Wl,-z,relro -fno-strict-aliasing -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -L/build/ceph/build/lib -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -iquote/build/ceph/src/include /build/ceph/build/src/pybind/rados/build/ceph/build/src/pybind/rados/pyrex/rados.o -L/usr/lib/python2.7/config-x86_64-linux-gnu -lrados -lpython2.7 -lpthread -ldl -lutil -lm -o /build/ceph/build/lib/cython_modules/lib.2/rados.so
running install
running install_lib
copying /build/ceph/build/lib/cython_modules/lib.2/rados.so -> /build/ceph-install/usr/local/lib/python2.7/dist-packages
running install_egg_info
running egg_info
creating /build/ceph/build/src/pybind/rados/rados.egg-info
writing /build/ceph/build/src/pybind/rados/rados.egg-info/PKG-INFO
writing top-level names to /build/ceph/build/src/pybind/rados/rados.egg-info/top_level.txt
writing dependency_links to /build/ceph/build/src/pybind/rados/rados.egg-info/dependency_links.txt
writing manifest file '/build/ceph/build/src/pybind/rados/rados.egg-info/SOURCES.txt'
reading manifest file '/build/ceph/build/src/pybind/rados/rados.egg-info/SOURCES.txt'
reading manifest template 'MANIFEST.in'
writing manifest file '/build/ceph/build/src/pybind/rados/rados.egg-info/SOURCES.txt'
Copying /build/ceph/build/src/pybind/rados/rados.egg-info to /build/ceph-install/usr/local/lib/python2.7/dist-packages/rados-2.0.0.egg-info
Skipping SOURCES.txt
running install_scripts
writing list of installed files to '/dev/null'
running build
running build_ext
cythoning rbd.pyx to /build/ceph/build/src/pybind/rbd/pyrex/rbd.c
creating /build/ceph/build/src/pybind/rbd/pyrex
building 'rbd' extension
creating /build/ceph/build/src/pybind/rbd/build
creating /build/ceph/build/src/pybind/rbd/build/ceph
creating /build/ceph/build/src/pybind/rbd/build/ceph/build
creating /build/ceph/build/src/pybind/rbd/build/ceph/build/src
creating /build/ceph/build/src/pybind/rbd/build/ceph/build/src/pybind
creating /build/ceph/build/src/pybind/rbd/build/ceph/build/src/pybind/rbd
creating /build/ceph/build/src/pybind/rbd/build/ceph/build/src/pybind/rbd/pyrex
/usr/lib/ccache/x86_64-linux-gnu-gcc -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -fno-strict-aliasing -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -iquote/build/ceph/src/include -fPIC -I/usr/include/python2.7 -I/usr/include/x86_64-linux-gnu/python2.7 -I/usr/include/python2.7 -c /build/ceph/build/src/pybind/rbd/pyrex/rbd.c -o /build/ceph/build/src/pybind/rbd/build/ceph/build/src/pybind/rbd/pyrex/rbd.o -fno-strict-aliasing -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -lpython2.7 -lpthread -ldl -lutil -lm -Xlinker -export-dynamic -Wl,-O1 -Wl,-Bsymbolic-functions
x86_64-linux-gnu-gcc -pthread -shared -Wl,-O1 -Wl,-Bsymbolic-functions -Wl,-Bsymbolic-functions -Wl,-z,relro -fno-strict-aliasing -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -L/build/ceph/build/lib -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -iquote/build/ceph/src/include /build/ceph/build/src/pybind/rbd/build/ceph/build/src/pybind/rbd/pyrex/rbd.o -L/usr/lib/python2.7/config-x86_64-linux-gnu -lrbd -lrados -lpython2.7 -lpthread -ldl -lutil -lm -o /build/ceph/build/lib/cython_modules/lib.2/rbd.so
running install
running install_lib
copying /build/ceph/build/lib/cython_modules/lib.2/rbd.so -> /build/ceph-install/usr/local/lib/python2.7/dist-packages
running install_egg_info
running egg_info
creating /build/ceph/build/src/pybind/rbd/rbd.egg-info
writing /build/ceph/build/src/pybind/rbd/rbd.egg-info/PKG-INFO
writing top-level names to /build/ceph/build/src/pybind/rbd/rbd.egg-info/top_level.txt
writing dependency_links to /build/ceph/build/src/pybind/rbd/rbd.egg-info/dependency_links.txt
writing manifest file '/build/ceph/build/src/pybind/rbd/rbd.egg-info/SOURCES.txt'
reading manifest file '/build/ceph/build/src/pybind/rbd/rbd.egg-info/SOURCES.txt'
reading manifest template 'MANIFEST.in'
writing manifest file '/build/ceph/build/src/pybind/rbd/rbd.egg-info/SOURCES.txt'
Copying /build/ceph/build/src/pybind/rbd/rbd.egg-info to /build/ceph-install/usr/local/lib/python2.7/dist-packages/rbd-2.0.0.egg-info
Skipping SOURCES.txt
running install_scripts
writing list of installed files to '/dev/null'
running build
running build_ext
cythoning cephfs.pyx to /build/ceph/build/src/pybind/cephfs/pyrex/cephfs.c
creating /build/ceph/build/src/pybind/cephfs/pyrex
building 'cephfs' extension
creating /build/ceph/build/src/pybind/cephfs/build
creating /build/ceph/build/src/pybind/cephfs/build/ceph
creating /build/ceph/build/src/pybind/cephfs/build/ceph/build
creating /build/ceph/build/src/pybind/cephfs/build/ceph/build/src
creating /build/ceph/build/src/pybind/cephfs/build/ceph/build/src/pybind
creating /build/ceph/build/src/pybind/cephfs/build/ceph/build/src/pybind/cephfs
creating /build/ceph/build/src/pybind/cephfs/build/ceph/build/src/pybind/cephfs/pyrex
/usr/lib/ccache/x86_64-linux-gnu-gcc -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -fno-strict-aliasing -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -iquote/build/ceph/src/include -fPIC -I/usr/include/python2.7 -I/usr/include/x86_64-linux-gnu/python2.7 -I/usr/include/python2.7 -c /build/ceph/build/src/pybind/cephfs/pyrex/cephfs.c -o /build/ceph/build/src/pybind/cephfs/build/ceph/build/src/pybind/cephfs/pyrex/cephfs.o -fno-strict-aliasing -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -lpython2.7 -lpthread -ldl -lutil -lm -Xlinker -export-dynamic -Wl,-O1 -Wl,-Bsymbolic-functions
x86_64-linux-gnu-gcc -pthread -shared -Wl,-O1 -Wl,-Bsymbolic-functions -Wl,-Bsymbolic-functions -Wl,-z,relro -fno-strict-aliasing -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -L/build/ceph/build/lib -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -iquote/build/ceph/src/include /build/ceph/build/src/pybind/cephfs/build/ceph/build/src/pybind/cephfs/pyrex/cephfs.o -L/usr/lib/python2.7/config-x86_64-linux-gnu -lcephfs -lpython2.7 -lpthread -ldl -lutil -lm -o /build/ceph/build/lib/cython_modules/lib.2/cephfs.so
running install
running install_lib
copying /build/ceph/build/lib/cython_modules/lib.2/cephfs.so -> /build/ceph-install/usr/local/lib/python2.7/dist-packages
running install_egg_info
running egg_info
creating /build/ceph/build/src/pybind/cephfs/cephfs.egg-info
writing /build/ceph/build/src/pybind/cephfs/cephfs.egg-info/PKG-INFO
writing top-level names to /build/ceph/build/src/pybind/cephfs/cephfs.egg-info/top_level.txt
writing dependency_links to /build/ceph/build/src/pybind/cephfs/cephfs.egg-info/dependency_links.txt
writing manifest file '/build/ceph/build/src/pybind/cephfs/cephfs.egg-info/SOURCES.txt'
reading manifest file '/build/ceph/build/src/pybind/cephfs/cephfs.egg-info/SOURCES.txt'
reading manifest template 'MANIFEST.in'
writing manifest file '/build/ceph/build/src/pybind/cephfs/cephfs.egg-info/SOURCES.txt'
Copying /build/ceph/build/src/pybind/cephfs/cephfs.egg-info to /build/ceph-install/usr/local/lib/python2.7/dist-packages/cephfs-2.0.0.egg-info
Skipping SOURCES.txt
running install_scripts
writing list of installed files to '/dev/null'
running build
running build_ext
cythoning rgw.pyx to /build/ceph/build/src/pybind/rgw/pyrex/rgw.c
creating /build/ceph/build/src/pybind/rgw/pyrex
building 'rgw' extension
creating /build/ceph/build/src/pybind/rgw/build
creating /build/ceph/build/src/pybind/rgw/build/ceph
creating /build/ceph/build/src/pybind/rgw/build/ceph/build
creating /build/ceph/build/src/pybind/rgw/build/ceph/build/src
creating /build/ceph/build/src/pybind/rgw/build/ceph/build/src/pybind
creating /build/ceph/build/src/pybind/rgw/build/ceph/build/src/pybind/rgw
creating /build/ceph/build/src/pybind/rgw/build/ceph/build/src/pybind/rgw/pyrex
/usr/lib/ccache/x86_64-linux-gnu-gcc -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -fno-strict-aliasing -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -iquote/build/ceph/src/include -fPIC -I/usr/include/python2.7 -I/usr/include/x86_64-linux-gnu/python2.7 -I/usr/include/python2.7 -c /build/ceph/build/src/pybind/rgw/pyrex/rgw.c -o /build/ceph/build/src/pybind/rgw/build/ceph/build/src/pybind/rgw/pyrex/rgw.o -fno-strict-aliasing -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -lpython2.7 -lpthread -ldl -lutil -lm -Xlinker -export-dynamic -Wl,-O1 -Wl,-Bsymbolic-functions
/build/ceph/build/src/pybind/rgw/pyrex/rgw.c: In function '__pyx_pf_3rgw_8LibRGWFS_28readdir':
/build/ceph/build/src/pybind/rgw/pyrex/rgw.c:6849:93: warning: passing argument 4 of 'rgw_readdir' from incompatible pointer type [-Wincompatible-pointer-types]
         __pyx_t_5 = rgw_readdir(__pyx_v_self->fs, __pyx_v__dir_handler, (&__pyx_v__offset), (&__pyx_f_3rgw_readdir_cb), ((void *)__pyx_v_iterate_cb), (&__pyx_v__eof), __pyx_v__flags); if (unlikely(__pyx_t_5 == -9000 && __Pyx_ErrOccurredWithGIL())) __PYX_ERR(0, 528, __pyx_L4_error)
                                                                                             ^
In file included from /build/ceph/build/src/pybind/rgw/pyrex/rgw.c:441:0:
/build/ceph/src/include/rados/rgw_file.h:219:5: note: expected 'rgw_readdir_cb {aka _Bool (*)(const char *, void *, long unsigned int,  unsigned int)}' but argument is of type '_Bool (*)(const char *, void *, uint64_t) {aka _Bool (*)(const char *, void *, long unsigned int)}'
 int rgw_readdir(struct rgw_fs *rgw_fs,
     ^~~~~~~~~~~
x86_64-linux-gnu-gcc -pthread -shared -Wl,-O1 -Wl,-Bsymbolic-functions -Wl,-Bsymbolic-functions -Wl,-z,relro -fno-strict-aliasing -DNDEBUG -g -fwrapv -O2 -Wall -Wstrict-prototypes -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -L/build/ceph/build/lib -Wdate-time -D_FORTIFY_SOURCE=2 -g -fdebug-prefix-map=/build/python2.7-IY_Jmw/python2.7-2.7.13=. -fstack-protector-strong -Wformat -Werror=format-security -iquote/build/ceph/src/include /build/ceph/build/src/pybind/rgw/build/ceph/build/src/pybind/rgw/pyrex/rgw.o -L/usr/lib/python2.7/config-x86_64-linux-gnu -lrados -lrgw -lpython2.7 -lpthread -ldl -lutil -lm -o /build/ceph/build/lib/cython_modules/lib.2/rgw.so
running install
running install_lib
copying /build/ceph/build/lib/cython_modules/lib.2/rgw.so -> /build/ceph-install/usr/local/lib/python2.7/dist-packages
running install_egg_info
running egg_info
creating /build/ceph/build/src/pybind/rgw/rgw.egg-info
writing /build/ceph/build/src/pybind/rgw/rgw.egg-info/PKG-INFO
writing top-level names to /build/ceph/build/src/pybind/rgw/rgw.egg-info/top_level.txt
writing dependency_links to /build/ceph/build/src/pybind/rgw/rgw.egg-info/dependency_links.txt
writing manifest file '/build/ceph/build/src/pybind/rgw/rgw.egg-info/SOURCES.txt'
reading manifest file '/build/ceph/build/src/pybind/rgw/rgw.egg-info/SOURCES.txt'
reading manifest template 'MANIFEST.in'
writing manifest file '/build/ceph/build/src/pybind/rgw/rgw.egg-info/SOURCES.txt'
Copying /build/ceph/build/src/pybind/rgw/rgw.egg-info to /build/ceph-install/usr/local/lib/python2.7/dist-packages/rgw-2.0.0.egg-info
Skipping SOURCES.txt
running install_scripts
writing list of installed files to '/dev/null'
/usr/bin/python2.7: can't open file 'setup.py': [Errno 2] No such file or directory
/usr/bin/python2.7: can't open file 'setup.py': [Errno 2] No such file or directory
/usr/bin/python2.7: can't open file 'setup.py': [Errno 2] No such file or directory
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_jerasure.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_jerasure_generic.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_jerasure_sse3.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_jerasure_sse4.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_lrc.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_shec.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_shec_generic.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_shec_sse3.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_shec_sse4.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/erasure-code/libec_isa.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/compressor/libceph_snappy.so.2.0.0
-- Installing: /build/ceph-install/usr/local/lib/ceph/compressor/libceph_snappy.so.2
-- Installing: /build/ceph-install/usr/local/lib/ceph/compressor/libceph_snappy.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/compressor/libceph_zlib.so.2.0.0
-- Installing: /build/ceph-install/usr/local/lib/ceph/compressor/libceph_zlib.so.2
-- Installing: /build/ceph-install/usr/local/lib/ceph/compressor/libceph_zlib.so
-- Installing: /build/ceph-install/usr/local/lib/ceph/compressor/libceph_zstd.so.2.0.0
-- Installing: /build/ceph-install/usr/local/lib/ceph/compressor/libceph_zstd.so.2
-- Installing: /build/ceph-install/usr/local/lib/ceph/compressor/libceph_zstd.so
-- Set runtime path of "/build/ceph-install/usr/local/lib/ceph/compressor/libceph_zstd.so.2.0.0" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/rados
-- Set runtime path of "/build/ceph-install/usr/local/bin/rados" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/lib/ceph/ceph-monstore-update-crush.sh
-- Installing: /build/ceph-install/usr/local/bin/crushtool
-- Set runtime path of "/build/ceph-install/usr/local/bin/crushtool" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/monmaptool
-- Set runtime path of "/build/ceph-install/usr/local/bin/monmaptool" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/ceph-authtool
-- Set runtime path of "/build/ceph-install/usr/local/bin/ceph-authtool" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/rbd
-- Set runtime path of "/build/ceph-install/usr/local/bin/rbd" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/lib/ceph/crypto/libceph_crypto_isal.so.1.0.0
-- Installing: /build/ceph-install/usr/local/lib/ceph/crypto/libceph_crypto_isal.so.1
-- Installing: /build/ceph-install/usr/local/lib/ceph/crypto/libceph_crypto_isal.so
-- Set runtime path of "/build/ceph-install/usr/local/lib/ceph/crypto/libceph_crypto_isal.so.1.0.0" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/etc/bash_completion.d/ceph
-- Installing: /build/ceph-install/usr/local/etc/bash_completion.d/rados
-- Installing: /build/ceph-install/usr/local/etc/bash_completion.d/rbd
-- Installing: /build/ceph-install/usr/local/etc/bash_completion.d/radosgw-admin
-- Installing: /build/ceph-install/usr/local/lib/librbd.so.1.12.0
-- Installing: /build/ceph-install/usr/local/lib/librbd.so.1
-- Installing: /build/ceph-install/usr/local/lib/librbd.so
-- Set runtime path of "/build/ceph-install/usr/local/lib/librbd.so.1.12.0" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/radosgw
-- Set runtime path of "/build/ceph-install/usr/local/bin/radosgw" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/radosgw-admin
-- Set runtime path of "/build/ceph-install/usr/local/bin/radosgw-admin" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/lib/librgw.so.2.0.0
-- Installing: /build/ceph-install/usr/local/lib/librgw.so.2
-- Installing: /build/ceph-install/usr/local/lib/librgw.so
-- Set runtime path of "/build/ceph-install/usr/local/lib/librgw.so.2.0.0" to "/usr/local/lib/ceph"
-- Installing: /build/ceph-install/usr/local/bin/ceph-brag
f4355289a0cdd12ae193a015b388537f18a537af7850e629b5424725065c8409
Sending build context to Docker daemon 156.3 MB
Step 1/2 : FROM scratch
 ---> 
Step 2/2 : COPY ceph-install.tar /
 ---> 62d5472affcf
Removing intermediate container 87f4c36a4233
Successfully built 62d5472affcf
make[3]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== caching image build-be9394dd/ceph-builder-amd64-tcmalloc-release
=== caching image build-be9394dd/ceph-amd64-tcmalloc-release-c19572dcbeb0
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
```

```
[vagrant@localhost images]$ make V=1 rook.linux_amd64
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/rootfs-builder-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/3 : FROM ubuntu:zesty
Trying to pull repository docker.io/library/ubuntu ... 
sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56: Pulling from docker.io/library/ubuntu
Digest: sha256:86a4a4fb5680baae304ac57337ffc8920343f8ebc019090acd390a4149622c56
Status: Image is up to date for docker.io/ubuntu:zesty
 ---> fe1cc5b91830
Step 2/3 : RUN sed -i -e 's/^deb-src/#deb-src/' /etc/apt/sources.list &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy --no-install-recommends         net-tools         netcat &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -yy --no-install-recommends &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean -y
 ---> Using cache
 ---> 179e64f482d5
Step 3/3 : RUN tar -czf /usr/share/copyrights.tar.xz /usr/share/doc/*/copyright &&     rm -rf         /usr/share/doc         /usr/share/man         /usr/share/info         /usr/share/locale         /var/lib/apt/lists/*         /var/log/*         /var/cache/debconf/*         /etc/systemd         /lib/lsb         /lib/udev         /usr/lib/x86_64-linux-gnu/gconv/IBM*         /usr/lib/x86_64-linux-gnu/gconv/EBC* &&     mkdir -p /usr/share/man/man{1..8}
 ---> Using cache
 ---> 8b040d721516
Successfully built 8b040d721516
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== docker build build-be9394dd/base-amd64
Sending build context to Docker daemon 3.584 kB
Step 1/7 : FROM build-be9394dd/rootfs-amd64-8b040d721516
 ---> 2e3b91214509
Step 2/7 : RUN HOST_IP=$(route -n | awk '/^0.0.0.0/ {print $2}') &&     echo "#!/bin/bash" > /usr/local/bin/apt-ng-host-discover &&     echo "if nc -w1 -z $HOST_IP 3142; then printf http://$HOST_IP:3142; else printf DIRECT; fi" >> /usr/local/bin/apt-ng-host-discover &&     chmod +x /usr/local/bin/apt-ng-host-discover &&     echo 'Acquire::http::ProxyAutoDetect "/usr/local/bin/apt-ng-host-discover";' > /etc/apt/apt.conf.d/30proxy
 ---> Using cache
 ---> e3eb9ca4c9ff
Step 3/7 : ARG ARCH
 ---> Using cache
 ---> cb5e0f44ebbb
Step 4/7 : ADD http://localhost:48080/windows10_drive_f/99-mirror/https0x3A0x2F0x2Fgithub.com0x2Fkrallin0x2Ftini0x2Freleases0x2Fdownload/v0.14.0/tini-${ARCH} /tini
Downloading [==================================================>] 19.89 kB/19.89 kB
 ---> Using cache
 ---> 39a370fa6cb6
Step 5/7 : RUN chmod +x /tini
 ---> Using cache
 ---> 34fd3135dcaf
Step 6/7 : ENTRYPOINT /tini -g --
 ---> Using cache
 ---> 4c72f52634fd
Step 7/7 : CMD /bin/bash
 ---> Using cache
 ---> bdc3253aab81
Successfully built bdc3253aab81
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
=== caching image build-be9394dd/base-amd64
=== caching image build-be9394dd/rootfs-builder-amd64
=== caching image build-be9394dd/rootfs-amd64-8b040d721516
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/base'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== docker build build-be9394dd/cross-gnu-amd64
Sending build context to Docker daemon 13.82 kB
Step 1/5 : FROM build-be9394dd/base-amd64
 ---> bdc3253aab81
Step 2/5 : RUN echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty main universe restricted" > /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=amd64] http://archive.ubuntu.com/ubuntu zesty-backports main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-updates main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-security main universe restricted" >> /etc/apt/sources.list &&    echo "deb [arch=armhf,arm64] http://ports.ubuntu.com/ zesty-backports main universe restricted" >> /etc/apt/sources.list &&    dpkg --add-architecture armhf &&     dpkg --add-architecture arm64 &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         build-essential         ca-certificates         ccache         cmake         crossbuild-essential-arm64         crossbuild-essential-armhf         curl         git         jq         patch         yasm         zip &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> b4f4549a106a
Step 3/5 : ADD toolchain /usr/local/toolchain
 ---> Using cache
 ---> 3c2fc8b13df0
Step 4/5 : WORKDIR /build
 ---> Using cache
 ---> 5af7d7101090
Step 5/5 : ENV TERM xterm
 ---> Using cache
 ---> c4a654332a52
Successfully built c4a654332a52
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
=== caching image build-be9394dd/cross-gnu-amd64
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/cross-gnu'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== docker build build-be9394dd/ceph-builder-amd64-tcmalloc-release
Sending build context to Docker daemon 14.85 kB
Step 1/16 : FROM build-be9394dd/cross-gnu-amd64
 ---> c4a654332a52
Step 2/16 : ARG CEPH_GIT_REPO
 ---> Using cache
 ---> d8a28e9b2948
Step 3/16 : RUN git clone ${CEPH_GIT_REPO} ceph
 ---> Using cache
 ---> ce6f3e2a9e94
Step 4/16 : ARG ARCH
 ---> Using cache
 ---> 10b0906f2885
Step 5/16 : ARG CROSS_TRIPLE
 ---> Using cache
 ---> 18fb3b2e6d58
Step 6/16 : RUN DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         python-setuptools         gperf                 cython:${ARCH}         python:${ARCH}         libaio-dev:${ARCH}         libatomic-ops-dev:${ARCH}         libbabeltrace-dev:${ARCH}         libblkid-dev:${ARCH}         libboost-context-dev:${ARCH}         libboost-coroutine-dev:${ARCH}         libboost-date-time-dev:${ARCH}         libboost-filesystem-dev:${ARCH}         libboost-iostreams-dev:${ARCH}         libboost-program-options-dev:${ARCH}         libboost-python-dev:${ARCH}         libboost-random-dev:${ARCH}         libboost-regex-dev:${ARCH}         libboost-system-dev:${ARCH}         libboost-thread-dev:${ARCH}         libcurl4-gnutls-dev:${ARCH}         libexpat1-dev:${ARCH}         libgoogle-perftools-dev:${ARCH}         libgoogle-perftools4:${ARCH}         libibverbs-dev:${ARCH}         libjemalloc-dev:${ARCH}         libkeyutils-dev:${ARCH}         libldap2-dev:${ARCH}         libnss3-dev:${ARCH}         libpython-dev:${ARCH}         libsnappy-dev:${ARCH}         libssl-dev:${ARCH}         libtcmalloc-minimal4:${ARCH}         libudev-dev:${ARCH}         libunwind-dev:${ARCH}         zlib1g-dev:${ARCH} &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> 8dc3989141e0
Step 7/16 : ARG CEPH_GIT_COMMIT
 ---> Using cache
 ---> 9d5d869cd549
Step 8/16 : RUN cd /build/ceph &&     git fetch --all --prune &&     git checkout -b ceph-builder ${CEPH_GIT_COMMIT} &&     git submodule update --init --recursive
 ---> Using cache
 ---> 1c0a3108f9d6
Step 9/16 : RUN mv /build/ceph/CMakeLists.txt /build/ceph/CMakeLists.original.txt
 ---> Using cache
 ---> 70dfa574cf25
Step 10/16 : ADD CMakeLists.txt /build/ceph
 ---> Using cache
 ---> 3b3c3ec08145
Step 11/16 : ARG CEPH_BUILD_TYPE
 ---> Using cache
 ---> 3f5087dd3915
Step 12/16 : ARG CEPH_ALLOCATOR
 ---> Using cache
 ---> 225c822b412c
Step 13/16 : RUN mkdir -p /build/ceph/build && cd /build/ceph/build &&     cmake     -DCMAKE_SKIP_INSTALL_ALL_DEPENDENCY=ON     -DCMAKE_INSTALL_PREFIX=/usr/local     -DCMAKE_BUILD_TYPE=${CEPH_BUILD_TYPE}     -DCMAKE_TOOLCHAIN_FILE=/usr/local/toolchain/${CROSS_TRIPLE}.cmake     -DALLOCATOR=${CEPH_ALLOCATOR}     -DWITH_SYSTEM_BOOST=ON     -DWITH_EMBEDDED=OFF     -DWITH_FUSE=OFF     -DWITH_LEVELDB=OFF     -DWITH_LTTNG=OFF     -DWITH_MANPAGE=OFF     -DWITH_PROFILER=OFF     -DWITH_PYTHON3=OFF     -DWITH_RADOSGW_FCGI_FRONTEND=OFF     ..
 ---> Using cache
 ---> 95367828ae35
Step 14/16 : WORKDIR /build/ceph/build
 ---> Using cache
 ---> 8c46aca7ea3c
Step 15/16 : ADD build-ceph.sh /build/
 ---> Using cache
 ---> d9e5a61fdc5c
Step 16/16 : CMD /build/build-ceph.sh
 ---> Using cache
 ---> c19572dcbeb0
Successfully built c19572dcbeb0
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== caching image build-be9394dd/ceph-builder-amd64-tcmalloc-release
=== caching image build-be9394dd/ceph-amd64-tcmalloc-release-c19572dcbeb0
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[1]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/rook'
=== building ceph for build-be9394dd/rook-amd64
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== docker build build-be9394dd/ceph-builder-amd64-tcmalloc-release
Sending build context to Docker daemon 14.85 kB
Step 1/16 : FROM build-be9394dd/cross-gnu-amd64
 ---> c4a654332a52
Step 2/16 : ARG CEPH_GIT_REPO
 ---> Using cache
 ---> d8a28e9b2948
Step 3/16 : RUN git clone ${CEPH_GIT_REPO} ceph
 ---> Using cache
 ---> ce6f3e2a9e94
Step 4/16 : ARG ARCH
 ---> Using cache
 ---> 10b0906f2885
Step 5/16 : ARG CROSS_TRIPLE
 ---> Using cache
 ---> 18fb3b2e6d58
Step 6/16 : RUN DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         python-setuptools         gperf                 cython:${ARCH}         python:${ARCH}         libaio-dev:${ARCH}         libatomic-ops-dev:${ARCH}         libbabeltrace-dev:${ARCH}         libblkid-dev:${ARCH}         libboost-context-dev:${ARCH}         libboost-coroutine-dev:${ARCH}         libboost-date-time-dev:${ARCH}         libboost-filesystem-dev:${ARCH}         libboost-iostreams-dev:${ARCH}         libboost-program-options-dev:${ARCH}         libboost-python-dev:${ARCH}         libboost-random-dev:${ARCH}         libboost-regex-dev:${ARCH}         libboost-system-dev:${ARCH}         libboost-thread-dev:${ARCH}         libcurl4-gnutls-dev:${ARCH}         libexpat1-dev:${ARCH}         libgoogle-perftools-dev:${ARCH}         libgoogle-perftools4:${ARCH}         libibverbs-dev:${ARCH}         libjemalloc-dev:${ARCH}         libkeyutils-dev:${ARCH}         libldap2-dev:${ARCH}         libnss3-dev:${ARCH}         libpython-dev:${ARCH}         libsnappy-dev:${ARCH}         libssl-dev:${ARCH}         libtcmalloc-minimal4:${ARCH}         libudev-dev:${ARCH}         libunwind-dev:${ARCH}         zlib1g-dev:${ARCH} &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean &&     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Using cache
 ---> 8dc3989141e0
Step 7/16 : ARG CEPH_GIT_COMMIT
 ---> Using cache
 ---> 9d5d869cd549
Step 8/16 : RUN cd /build/ceph &&     git fetch --all --prune &&     git checkout -b ceph-builder ${CEPH_GIT_COMMIT} &&     git submodule update --init --recursive
 ---> Using cache
 ---> 1c0a3108f9d6
Step 9/16 : RUN mv /build/ceph/CMakeLists.txt /build/ceph/CMakeLists.original.txt
 ---> Using cache
 ---> 70dfa574cf25
Step 10/16 : ADD CMakeLists.txt /build/ceph
 ---> Using cache
 ---> 3b3c3ec08145
Step 11/16 : ARG CEPH_BUILD_TYPE
 ---> Using cache
 ---> 3f5087dd3915
Step 12/16 : ARG CEPH_ALLOCATOR
 ---> Using cache
 ---> 225c822b412c
Step 13/16 : RUN mkdir -p /build/ceph/build && cd /build/ceph/build &&     cmake     -DCMAKE_SKIP_INSTALL_ALL_DEPENDENCY=ON     -DCMAKE_INSTALL_PREFIX=/usr/local     -DCMAKE_BUILD_TYPE=${CEPH_BUILD_TYPE}     -DCMAKE_TOOLCHAIN_FILE=/usr/local/toolchain/${CROSS_TRIPLE}.cmake     -DALLOCATOR=${CEPH_ALLOCATOR}     -DWITH_SYSTEM_BOOST=ON     -DWITH_EMBEDDED=OFF     -DWITH_FUSE=OFF     -DWITH_LEVELDB=OFF     -DWITH_LTTNG=OFF     -DWITH_MANPAGE=OFF     -DWITH_PROFILER=OFF     -DWITH_PYTHON3=OFF     -DWITH_RADOSGW_FCGI_FRONTEND=OFF     ..
 ---> Using cache
 ---> 95367828ae35
Step 14/16 : WORKDIR /build/ceph/build
 ---> Using cache
 ---> 8c46aca7ea3c
Step 15/16 : ADD build-ceph.sh /build/
 ---> Using cache
 ---> d9e5a61fdc5c
Step 16/16 : CMD /build/build-ceph.sh
 ---> Using cache
 ---> c19572dcbeb0
Successfully built c19572dcbeb0
make[3]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[3]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
make[3]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== caching image build-be9394dd/ceph-builder-amd64-tcmalloc-release
=== caching image build-be9394dd/ceph-amd64-tcmalloc-release-c19572dcbeb0
make[3]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
b4735a060dd672251a033db144c40f58cef343e5f10b5c698da56fdb1f0eaf53
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/ceph'
=== docker build build-be9394dd/rook-base-amd64
Sending build context to Docker daemon 156.3 MB
Step 1/6 : FROM build-be9394dd/base-amd64
 ---> bdc3253aab81
Step 2/6 : ARG EXTRALIBS
 ---> Running in 588116654a7d
 ---> c7a446fe2507
Removing intermediate container 588116654a7d
Step 3/6 : RUN BOOST_VERSION=1.62.0 &&     DEBIAN_FRONTEND=noninteractive apt-get update &&     DEBIAN_FRONTEND=noninteractive apt-get install -yy -q --no-install-recommends         ca-certificates         gdisk         kmod         libaio1         libboost-context${BOOST_VERSION}         libboost-coroutine${BOOST_VERSION}         libboost-date-time${BOOST_VERSION}         libboost-filesystem${BOOST_VERSION}         libboost-iostreams${BOOST_VERSION}         libboost-program-options${BOOST_VERSION}         libboost-python${BOOST_VERSION}         libboost-random${BOOST_VERSION}         libboost-regex${BOOST_VERSION}         libboost-system${BOOST_VERSION}         libboost-thread${BOOST_VERSION}         libcurl3-gnutls         libexpat1         libibverbs1         libldap-2.4-2         libnss3         libpython2.7         libsnappy1v5         python-prettytable         python2.7-minimal         ${EXTRALIBS} &&     DEBIAN_FRONTEND=noninteractive apt-get upgrade -y &&     DEBIAN_FRONTEND=noninteractive apt-get autoremove -y &&     DEBIAN_FRONTEND=noninteractive apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
 ---> Running in 4f2248544f34
Get:1 http://archive.ubuntu.com/ubuntu zesty InRelease [243 kB]
Get:2 http://security.ubuntu.com/ubuntu zesty-security InRelease [89.2 kB]
Get:3 http://security.ubuntu.com/ubuntu zesty-security/main amd64 Packages [220 kB]
Get:4 http://archive.ubuntu.com/ubuntu zesty-updates InRelease [89.2 kB]
Get:5 http://security.ubuntu.com/ubuntu zesty-security/universe amd64 Packages [112 kB]
Get:6 http://security.ubuntu.com/ubuntu zesty-security/restricted amd64 Packages [3221 B]
Get:7 http://security.ubuntu.com/ubuntu zesty-security/multiverse amd64 Packages [3102 B]
Get:8 http://archive.ubuntu.com/ubuntu zesty-backports InRelease [89.2 kB]
Get:9 http://archive.ubuntu.com/ubuntu zesty/main amd64 Packages [1574 kB]
Get:10 http://archive.ubuntu.com/ubuntu zesty/universe amd64 Packages [10.5 MB]
Get:11 http://archive.ubuntu.com/ubuntu zesty/restricted amd64 Packages [14.3 kB]
Get:12 http://archive.ubuntu.com/ubuntu zesty/multiverse amd64 Packages [189 kB]
Get:13 http://archive.ubuntu.com/ubuntu zesty-updates/restricted amd64 Packages [3604 B]
Get:14 http://archive.ubuntu.com/ubuntu zesty-updates/universe amd64 Packages [217 kB]
Get:15 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 Packages [321 kB]
Get:16 http://archive.ubuntu.com/ubuntu zesty-updates/multiverse amd64 Packages [13.2 kB]
Get:17 http://archive.ubuntu.com/ubuntu zesty-backports/main amd64 Packages [1698 B]
Get:18 http://archive.ubuntu.com/ubuntu zesty-backports/universe amd64 Packages [4043 B]
Fetched 13.7 MB in 1min 17s (178 kB/s)
Reading package lists...
Reading package lists...
Building dependency tree...
Reading state information...
The following additional packages will be installed:
  libasn1-8-heimdal libffi6 libgmp10 libgnutls30 libgssapi-krb5-2
  libgssapi3-heimdal libhcrypto4-heimdal libheimbase1-heimdal
  libheimntlm0-heimdal libhogweed4 libhx509-5-heimdal libicu57 libidn11
  libidn2-0 libk5crypto3 libkeyutils1 libkmod2 libkrb5-26-heimdal libkrb5-3
  libkrb5support0 libldap-common libnettle6 libnl-3-200 libnl-route-3-200
  libnspr4 libp11-kit0 libpopt0 libpsl5 libpython-stdlib libpython2.7-minimal
  libpython2.7-stdlib libroken18-heimdal librtmp1 libsasl2-2
  libsasl2-modules-db libssl1.0.0 libtasn1-6 libtcmalloc-minimal4
  libunistring0 libunwind8 libwind0-heimdal mime-support openssl python
  python-minimal python2.7
Suggested packages:
  python3 gnutls-bin krb5-doc krb5-user python-doc python-tk python2.7-doc
  binutils binfmt-support
Recommended packages:
  groff-base krb5-locales publicsuffix libsasl2-modules bzip2 file xz-utils
The following NEW packages will be installed:
  ca-certificates gdisk kmod libaio1 libasn1-8-heimdal libboost-context1.62.0
  libboost-coroutine1.62.0 libboost-date-time1.62.0 libboost-filesystem1.62.0
  libboost-iostreams1.62.0 libboost-program-options1.62.0
  libboost-python1.62.0 libboost-random1.62.0 libboost-regex1.62.0
  libboost-system1.62.0 libboost-thread1.62.0 libcurl3-gnutls libexpat1
  libffi6 libgmp10 libgnutls30 libgoogle-perftools4 libgssapi-krb5-2
  libgssapi3-heimdal libhcrypto4-heimdal libheimbase1-heimdal
  libheimntlm0-heimdal libhogweed4 libhx509-5-heimdal libibverbs1 libicu57
  libidn11 libidn2-0 libk5crypto3 libkeyutils1 libkmod2 libkrb5-26-heimdal
  libkrb5-3 libkrb5support0 libldap-2.4-2 libldap-common libnettle6
  libnl-3-200 libnl-route-3-200 libnspr4 libnss3 libp11-kit0 libpopt0 libpsl5
  libpython-stdlib libpython2.7 libpython2.7-minimal libpython2.7-stdlib
  libroken18-heimdal librtmp1 libsasl2-2 libsasl2-modules-db libsnappy1v5
  libssl1.0.0 libtasn1-6 libtcmalloc-minimal4 libunistring0 libunwind8
  libwind0-heimdal mime-support openssl python python-minimal
  python-prettytable python2.7 python2.7-minimal
0 upgraded, 71 newly installed, 0 to remove and 0 not upgraded.
Need to get 20.9 MB of archives.
After this operation, 79.4 MB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu zesty/main amd64 libpopt0 amd64 1.16-10 [26.0 kB]
Get:2 http://archive.ubuntu.com/ubuntu zesty/main amd64 libunistring0 amd64 0.9.3-5.2ubuntu1 [279 kB]
Get:3 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython2.7-minimal amd64 2.7.13-2ubuntu0.1 [338 kB]
Get:4 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 python2.7-minimal amd64 2.7.13-2ubuntu0.1 [1323 kB]
Get:5 http://archive.ubuntu.com/ubuntu zesty/main amd64 python-minimal amd64 2.7.13-2 [28.2 kB]
Get:6 http://archive.ubuntu.com/ubuntu zesty/main amd64 mime-support all 3.60ubuntu1 [30.1 kB]
Get:7 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libexpat1 amd64 2.2.0-2ubuntu0.1 [72.3 kB]
Get:8 http://archive.ubuntu.com/ubuntu zesty/main amd64 libffi6 amd64 3.2.1-6 [17.7 kB]
Get:9 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libssl1.0.0 amd64 1.0.2g-1ubuntu11.4 [1080 kB]
Get:10 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython2.7-stdlib amd64 2.7.13-2ubuntu0.1 [1899 kB]
Get:11 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 python2.7 amd64 2.7.13-2ubuntu0.1 [229 kB]
Get:12 http://archive.ubuntu.com/ubuntu zesty/main amd64 libpython-stdlib amd64 2.7.13-2 [7774 B]
Get:13 http://archive.ubuntu.com/ubuntu zesty/main amd64 python amd64 2.7.13-2 [139 kB]
Get:14 http://archive.ubuntu.com/ubuntu zesty/main amd64 libkmod2 amd64 22-1.1ubuntu1 [39.9 kB]
Get:15 http://archive.ubuntu.com/ubuntu zesty/main amd64 kmod amd64 22-1.1ubuntu1 [87.7 kB]
Get:16 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libroken18-heimdal amd64 7.1.0+dfsg-9ubuntu1.1 [40.9 kB]
Get:17 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libasn1-8-heimdal amd64 7.1.0+dfsg-9ubuntu1.1 [176 kB]
Get:18 http://archive.ubuntu.com/ubuntu zesty/main amd64 libgmp10 amd64 2:6.1.2+dfsg-1 [240 kB]
Get:19 http://archive.ubuntu.com/ubuntu zesty/main amd64 libnettle6 amd64 3.3-1 [92.4 kB]
Get:20 http://archive.ubuntu.com/ubuntu zesty/main amd64 libhogweed4 amd64 3.3-1 [135 kB]
Get:21 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libidn11 amd64 1.33-1ubuntu0.1 [45.1 kB]
Get:22 http://archive.ubuntu.com/ubuntu zesty/main amd64 libp11-kit0 amd64 0.23.3-5 [107 kB]
Get:23 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libtasn1-6 amd64 4.10-1ubuntu0.1 [35.5 kB]
Get:24 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libgnutls30 amd64 3.5.6-4ubuntu4.3 [627 kB]
Get:25 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libheimbase1-heimdal amd64 7.1.0+dfsg-9ubuntu1.1 [29.3 kB]
Get:26 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libhcrypto4-heimdal amd64 7.1.0+dfsg-9ubuntu1.1 [87.3 kB]
Get:27 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libwind0-heimdal amd64 7.1.0+dfsg-9ubuntu1.1 [48.1 kB]
Get:28 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libhx509-5-heimdal amd64 7.1.0+dfsg-9ubuntu1.1 [107 kB]
Get:29 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libkrb5-26-heimdal amd64 7.1.0+dfsg-9ubuntu1.1 [206 kB]
Get:30 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libheimntlm0-heimdal amd64 7.1.0+dfsg-9ubuntu1.1 [15.1 kB]
Get:31 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libgssapi3-heimdal amd64 7.1.0+dfsg-9ubuntu1.1 [96.9 kB]
Get:32 http://archive.ubuntu.com/ubuntu zesty/main amd64 libsasl2-modules-db amd64 2.1.27~101-g0780600+dfsg-2ubuntu1 [14.3 kB]
Get:33 http://archive.ubuntu.com/ubuntu zesty/main amd64 libsasl2-2 amd64 2.1.27~101-g0780600+dfsg-2ubuntu1 [48.8 kB]
Get:34 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libldap-common all 2.4.44+dfsg-3ubuntu2.1 [16.8 kB]
Get:35 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libldap-2.4-2 amd64 2.4.44+dfsg-3ubuntu2.1 [154 kB]
Get:36 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 openssl amd64 1.0.2g-1ubuntu11.4 [491 kB]
Get:37 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 ca-certificates all 20170717~17.04.1 [167 kB]
Get:38 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libkrb5support0 amd64 1.15-1ubuntu0.1 [31.9 kB]
Get:39 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libk5crypto3 amd64 1.15-1ubuntu0.1 [84.8 kB]
Get:40 http://archive.ubuntu.com/ubuntu zesty/main amd64 libkeyutils1 amd64 1.5.9-9ubuntu1 [9512 B]
Get:41 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libkrb5-3 amd64 1.15-1ubuntu0.1 [275 kB]
Get:42 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libgssapi-krb5-2 amd64 1.15-1ubuntu0.1 [120 kB]
Get:43 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libidn2-0 amd64 0.16-1ubuntu0.1 [47.6 kB]
Get:44 http://archive.ubuntu.com/ubuntu zesty/main amd64 libpsl5 amd64 0.17.0-4 [39.4 kB]
Get:45 http://archive.ubuntu.com/ubuntu zesty/main amd64 librtmp1 amd64 2.4+20151223.gitfa8646d.1-1 [54.2 kB]
Get:46 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libcurl3-gnutls amd64 7.52.1-4ubuntu1.4 [192 kB]
Get:47 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libicu57 amd64 57.1-5ubuntu0.2 [7687 kB]
Get:48 http://archive.ubuntu.com/ubuntu zesty/main amd64 libaio1 amd64 0.3.110-3 [6382 B]
Get:49 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-context1.62.0 amd64 1.62.0+dfsg-4 [8298 B]
Get:50 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-system1.62.0 amd64 1.62.0+dfsg-4 [9398 B]
Get:51 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-thread1.62.0 amd64 1.62.0+dfsg-4 [46.9 kB]
Get:52 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-coroutine1.62.0 amd64 1.62.0+dfsg-4 [19.7 kB]
Get:53 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-date-time1.62.0 amd64 1.62.0+dfsg-4 [20.0 kB]
Get:54 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-filesystem1.62.0 amd64 1.62.0+dfsg-4 [38.6 kB]
Get:55 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-iostreams1.62.0 amd64 1.62.0+dfsg-4 [27.5 kB]
Get:56 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-program-options1.62.0 amd64 1.62.0+dfsg-4 [136 kB]
Get:57 http://archive.ubuntu.com/ubuntu zesty/universe amd64 libboost-python1.62.0 amd64 1.62.0+dfsg-4 [113 kB]
Get:58 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-random1.62.0 amd64 1.62.0+dfsg-4 [11.1 kB]
Get:59 http://archive.ubuntu.com/ubuntu zesty/main amd64 libboost-regex1.62.0 amd64 1.62.0+dfsg-4 [263 kB]
Get:60 http://archive.ubuntu.com/ubuntu zesty/main amd64 libtcmalloc-minimal4 amd64 2.5-0ubuntu4 [88.3 kB]
Get:61 http://archive.ubuntu.com/ubuntu zesty/main amd64 libunwind8 amd64 1.1-4.1ubuntu2 [46.3 kB]
Get:62 http://archive.ubuntu.com/ubuntu zesty/main amd64 libgoogle-perftools4 amd64 2.5-0ubuntu4 [183 kB]
Get:63 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnl-3-200 amd64 3.2.29-0ubuntu2.1 [52.9 kB]
Get:64 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnl-route-3-200 amd64 3.2.29-0ubuntu2.1 [146 kB]
Get:65 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnspr4 amd64 2:4.13.1-0ubuntu0.17.04.1 [112 kB]
Get:66 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libnss3 amd64 2:3.28.4-0ubuntu0.17.04.3 [1143 kB]
Get:67 http://archive.ubuntu.com/ubuntu zesty-updates/main amd64 libpython2.7 amd64 2.7.13-2ubuntu0.1 [1072 kB]
Get:68 http://archive.ubuntu.com/ubuntu zesty/main amd64 python-prettytable all 0.7.2-3 [19.6 kB]
Get:69 http://archive.ubuntu.com/ubuntu zesty/main amd64 gdisk amd64 1.0.1-1build1 [194 kB]
Get:70 http://archive.ubuntu.com/ubuntu zesty/main amd64 libibverbs1 amd64 1.2.1-2ubuntu1 [30.8 kB]
Get:71 http://archive.ubuntu.com/ubuntu zesty/main amd64 libsnappy1v5 amd64 1.1.4-1 [16.0 kB]
debconf: delaying package configuration, since apt-utils is not installed
Fetched 20.9 MB in 1min 52s (186 kB/s)
Selecting previously unselected package libpopt0:amd64.
(Reading database ... 4165 files and directories currently installed.)
Preparing to unpack .../00-libpopt0_1.16-10_amd64.deb ...
Unpacking libpopt0:amd64 (1.16-10) ...
Selecting previously unselected package libunistring0:amd64.
Preparing to unpack .../01-libunistring0_0.9.3-5.2ubuntu1_amd64.deb ...
Unpacking libunistring0:amd64 (0.9.3-5.2ubuntu1) ...
Selecting previously unselected package libpython2.7-minimal:amd64.
Preparing to unpack .../02-libpython2.7-minimal_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking libpython2.7-minimal:amd64 (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package python2.7-minimal.
Preparing to unpack .../03-python2.7-minimal_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking python2.7-minimal (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package python-minimal.
Preparing to unpack .../04-python-minimal_2.7.13-2_amd64.deb ...
Unpacking python-minimal (2.7.13-2) ...
Selecting previously unselected package mime-support.
Preparing to unpack .../05-mime-support_3.60ubuntu1_all.deb ...
Unpacking mime-support (3.60ubuntu1) ...
Selecting previously unselected package libexpat1:amd64.
Preparing to unpack .../06-libexpat1_2.2.0-2ubuntu0.1_amd64.deb ...
Unpacking libexpat1:amd64 (2.2.0-2ubuntu0.1) ...
Selecting previously unselected package libffi6:amd64.
Preparing to unpack .../07-libffi6_3.2.1-6_amd64.deb ...
Unpacking libffi6:amd64 (3.2.1-6) ...
Selecting previously unselected package libssl1.0.0:amd64.
Preparing to unpack .../08-libssl1.0.0_1.0.2g-1ubuntu11.4_amd64.deb ...
Unpacking libssl1.0.0:amd64 (1.0.2g-1ubuntu11.4) ...
Selecting previously unselected package libpython2.7-stdlib:amd64.
Preparing to unpack .../09-libpython2.7-stdlib_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking libpython2.7-stdlib:amd64 (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package python2.7.
Preparing to unpack .../10-python2.7_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking python2.7 (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package libpython-stdlib:amd64.
Preparing to unpack .../11-libpython-stdlib_2.7.13-2_amd64.deb ...
Unpacking libpython-stdlib:amd64 (2.7.13-2) ...
Setting up libpython2.7-minimal:amd64 (2.7.13-2ubuntu0.1) ...
Setting up python2.7-minimal (2.7.13-2ubuntu0.1) ...
Setting up python-minimal (2.7.13-2) ...
Selecting previously unselected package python.
(Reading database ... 4985 files and directories currently installed.)
Preparing to unpack .../00-python_2.7.13-2_amd64.deb ...
Unpacking python (2.7.13-2) ...
Selecting previously unselected package libkmod2:amd64.
Preparing to unpack .../01-libkmod2_22-1.1ubuntu1_amd64.deb ...
Unpacking libkmod2:amd64 (22-1.1ubuntu1) ...
Selecting previously unselected package kmod.
Preparing to unpack .../02-kmod_22-1.1ubuntu1_amd64.deb ...
Unpacking kmod (22-1.1ubuntu1) ...
Selecting previously unselected package libroken18-heimdal:amd64.
Preparing to unpack .../03-libroken18-heimdal_7.1.0+dfsg-9ubuntu1.1_amd64.deb ...
Unpacking libroken18-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Selecting previously unselected package libasn1-8-heimdal:amd64.
Preparing to unpack .../04-libasn1-8-heimdal_7.1.0+dfsg-9ubuntu1.1_amd64.deb ...
Unpacking libasn1-8-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Selecting previously unselected package libgmp10:amd64.
Preparing to unpack .../05-libgmp10_2%3a6.1.2+dfsg-1_amd64.deb ...
Unpacking libgmp10:amd64 (2:6.1.2+dfsg-1) ...
Selecting previously unselected package libnettle6:amd64.
Preparing to unpack .../06-libnettle6_3.3-1_amd64.deb ...
Unpacking libnettle6:amd64 (3.3-1) ...
Selecting previously unselected package libhogweed4:amd64.
Preparing to unpack .../07-libhogweed4_3.3-1_amd64.deb ...
Unpacking libhogweed4:amd64 (3.3-1) ...
Selecting previously unselected package libidn11:amd64.
Preparing to unpack .../08-libidn11_1.33-1ubuntu0.1_amd64.deb ...
Unpacking libidn11:amd64 (1.33-1ubuntu0.1) ...
Selecting previously unselected package libp11-kit0:amd64.
Preparing to unpack .../09-libp11-kit0_0.23.3-5_amd64.deb ...
Unpacking libp11-kit0:amd64 (0.23.3-5) ...
Selecting previously unselected package libtasn1-6:amd64.
Preparing to unpack .../10-libtasn1-6_4.10-1ubuntu0.1_amd64.deb ...
Unpacking libtasn1-6:amd64 (4.10-1ubuntu0.1) ...
Selecting previously unselected package libgnutls30:amd64.
Preparing to unpack .../11-libgnutls30_3.5.6-4ubuntu4.3_amd64.deb ...
Unpacking libgnutls30:amd64 (3.5.6-4ubuntu4.3) ...
Selecting previously unselected package libheimbase1-heimdal:amd64.
Preparing to unpack .../12-libheimbase1-heimdal_7.1.0+dfsg-9ubuntu1.1_amd64.deb ...
Unpacking libheimbase1-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Selecting previously unselected package libhcrypto4-heimdal:amd64.
Preparing to unpack .../13-libhcrypto4-heimdal_7.1.0+dfsg-9ubuntu1.1_amd64.deb ...
Unpacking libhcrypto4-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Selecting previously unselected package libwind0-heimdal:amd64.
Preparing to unpack .../14-libwind0-heimdal_7.1.0+dfsg-9ubuntu1.1_amd64.deb ...
Unpacking libwind0-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Selecting previously unselected package libhx509-5-heimdal:amd64.
Preparing to unpack .../15-libhx509-5-heimdal_7.1.0+dfsg-9ubuntu1.1_amd64.deb ...
Unpacking libhx509-5-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Selecting previously unselected package libkrb5-26-heimdal:amd64.
Preparing to unpack .../16-libkrb5-26-heimdal_7.1.0+dfsg-9ubuntu1.1_amd64.deb ...
Unpacking libkrb5-26-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Selecting previously unselected package libheimntlm0-heimdal:amd64.
Preparing to unpack .../17-libheimntlm0-heimdal_7.1.0+dfsg-9ubuntu1.1_amd64.deb ...
Unpacking libheimntlm0-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Selecting previously unselected package libgssapi3-heimdal:amd64.
Preparing to unpack .../18-libgssapi3-heimdal_7.1.0+dfsg-9ubuntu1.1_amd64.deb ...
Unpacking libgssapi3-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Selecting previously unselected package libsasl2-modules-db:amd64.
Preparing to unpack .../19-libsasl2-modules-db_2.1.27~101-g0780600+dfsg-2ubuntu1_amd64.deb ...
Unpacking libsasl2-modules-db:amd64 (2.1.27~101-g0780600+dfsg-2ubuntu1) ...
Selecting previously unselected package libsasl2-2:amd64.
Preparing to unpack .../20-libsasl2-2_2.1.27~101-g0780600+dfsg-2ubuntu1_amd64.deb ...
Unpacking libsasl2-2:amd64 (2.1.27~101-g0780600+dfsg-2ubuntu1) ...
Selecting previously unselected package libldap-common.
Preparing to unpack .../21-libldap-common_2.4.44+dfsg-3ubuntu2.1_all.deb ...
Unpacking libldap-common (2.4.44+dfsg-3ubuntu2.1) ...
Selecting previously unselected package libldap-2.4-2:amd64.
Preparing to unpack .../22-libldap-2.4-2_2.4.44+dfsg-3ubuntu2.1_amd64.deb ...
Unpacking libldap-2.4-2:amd64 (2.4.44+dfsg-3ubuntu2.1) ...
Selecting previously unselected package openssl.
Preparing to unpack .../23-openssl_1.0.2g-1ubuntu11.4_amd64.deb ...
Unpacking openssl (1.0.2g-1ubuntu11.4) ...
Selecting previously unselected package ca-certificates.
Preparing to unpack .../24-ca-certificates_20170717~17.04.1_all.deb ...
Unpacking ca-certificates (20170717~17.04.1) ...
Selecting previously unselected package libkrb5support0:amd64.
Preparing to unpack .../25-libkrb5support0_1.15-1ubuntu0.1_amd64.deb ...
Unpacking libkrb5support0:amd64 (1.15-1ubuntu0.1) ...
Selecting previously unselected package libk5crypto3:amd64.
Preparing to unpack .../26-libk5crypto3_1.15-1ubuntu0.1_amd64.deb ...
Unpacking libk5crypto3:amd64 (1.15-1ubuntu0.1) ...
Selecting previously unselected package libkeyutils1:amd64.
Preparing to unpack .../27-libkeyutils1_1.5.9-9ubuntu1_amd64.deb ...
Unpacking libkeyutils1:amd64 (1.5.9-9ubuntu1) ...
Selecting previously unselected package libkrb5-3:amd64.
Preparing to unpack .../28-libkrb5-3_1.15-1ubuntu0.1_amd64.deb ...
Unpacking libkrb5-3:amd64 (1.15-1ubuntu0.1) ...
Selecting previously unselected package libgssapi-krb5-2:amd64.
Preparing to unpack .../29-libgssapi-krb5-2_1.15-1ubuntu0.1_amd64.deb ...
Unpacking libgssapi-krb5-2:amd64 (1.15-1ubuntu0.1) ...
Selecting previously unselected package libidn2-0:amd64.
Preparing to unpack .../30-libidn2-0_0.16-1ubuntu0.1_amd64.deb ...
Unpacking libidn2-0:amd64 (0.16-1ubuntu0.1) ...
Selecting previously unselected package libpsl5:amd64.
Preparing to unpack .../31-libpsl5_0.17.0-4_amd64.deb ...
Unpacking libpsl5:amd64 (0.17.0-4) ...
Selecting previously unselected package librtmp1:amd64.
Preparing to unpack .../32-librtmp1_2.4+20151223.gitfa8646d.1-1_amd64.deb ...
Unpacking librtmp1:amd64 (2.4+20151223.gitfa8646d.1-1) ...
Selecting previously unselected package libcurl3-gnutls:amd64.
Preparing to unpack .../33-libcurl3-gnutls_7.52.1-4ubuntu1.4_amd64.deb ...
Unpacking libcurl3-gnutls:amd64 (7.52.1-4ubuntu1.4) ...
Selecting previously unselected package libicu57:amd64.
Preparing to unpack .../34-libicu57_57.1-5ubuntu0.2_amd64.deb ...
Unpacking libicu57:amd64 (57.1-5ubuntu0.2) ...
Selecting previously unselected package libaio1:amd64.
Preparing to unpack .../35-libaio1_0.3.110-3_amd64.deb ...
Unpacking libaio1:amd64 (0.3.110-3) ...
Selecting previously unselected package libboost-context1.62.0:amd64.
Preparing to unpack .../36-libboost-context1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-context1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-system1.62.0:amd64.
Preparing to unpack .../37-libboost-system1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-system1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-thread1.62.0:amd64.
Preparing to unpack .../38-libboost-thread1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-thread1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-coroutine1.62.0:amd64.
Preparing to unpack .../39-libboost-coroutine1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-coroutine1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-date-time1.62.0:amd64.
Preparing to unpack .../40-libboost-date-time1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-date-time1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-filesystem1.62.0:amd64.
Preparing to unpack .../41-libboost-filesystem1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-filesystem1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-iostreams1.62.0:amd64.
Preparing to unpack .../42-libboost-iostreams1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-iostreams1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-program-options1.62.0:amd64.
Preparing to unpack .../43-libboost-program-options1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-program-options1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-python1.62.0.
Preparing to unpack .../44-libboost-python1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-python1.62.0 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-random1.62.0:amd64.
Preparing to unpack .../45-libboost-random1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-random1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libboost-regex1.62.0:amd64.
Preparing to unpack .../46-libboost-regex1.62.0_1.62.0+dfsg-4_amd64.deb ...
Unpacking libboost-regex1.62.0:amd64 (1.62.0+dfsg-4) ...
Selecting previously unselected package libtcmalloc-minimal4.
Preparing to unpack .../47-libtcmalloc-minimal4_2.5-0ubuntu4_amd64.deb ...
Unpacking libtcmalloc-minimal4 (2.5-0ubuntu4) ...
Selecting previously unselected package libunwind8.
Preparing to unpack .../48-libunwind8_1.1-4.1ubuntu2_amd64.deb ...
Unpacking libunwind8 (1.1-4.1ubuntu2) ...
Selecting previously unselected package libgoogle-perftools4.
Preparing to unpack .../49-libgoogle-perftools4_2.5-0ubuntu4_amd64.deb ...
Unpacking libgoogle-perftools4 (2.5-0ubuntu4) ...
Selecting previously unselected package libnl-3-200:amd64.
Preparing to unpack .../50-libnl-3-200_3.2.29-0ubuntu2.1_amd64.deb ...
Unpacking libnl-3-200:amd64 (3.2.29-0ubuntu2.1) ...
Selecting previously unselected package libnl-route-3-200:amd64.
Preparing to unpack .../51-libnl-route-3-200_3.2.29-0ubuntu2.1_amd64.deb ...
Unpacking libnl-route-3-200:amd64 (3.2.29-0ubuntu2.1) ...
Selecting previously unselected package libnspr4:amd64.
Preparing to unpack .../52-libnspr4_2%3a4.13.1-0ubuntu0.17.04.1_amd64.deb ...
Unpacking libnspr4:amd64 (2:4.13.1-0ubuntu0.17.04.1) ...
Selecting previously unselected package libnss3:amd64.
Preparing to unpack .../53-libnss3_2%3a3.28.4-0ubuntu0.17.04.3_amd64.deb ...
Unpacking libnss3:amd64 (2:3.28.4-0ubuntu0.17.04.3) ...
Selecting previously unselected package libpython2.7:amd64.
Preparing to unpack .../54-libpython2.7_2.7.13-2ubuntu0.1_amd64.deb ...
Unpacking libpython2.7:amd64 (2.7.13-2ubuntu0.1) ...
Selecting previously unselected package python-prettytable.
Preparing to unpack .../55-python-prettytable_0.7.2-3_all.deb ...
Unpacking python-prettytable (0.7.2-3) ...
Selecting previously unselected package gdisk.
Preparing to unpack .../56-gdisk_1.0.1-1build1_amd64.deb ...
Unpacking gdisk (1.0.1-1build1) ...
Selecting previously unselected package libibverbs1.
Preparing to unpack .../57-libibverbs1_1.2.1-2ubuntu1_amd64.deb ...
Unpacking libibverbs1 (1.2.1-2ubuntu1) ...
Selecting previously unselected package libsnappy1v5:amd64.
Preparing to unpack .../58-libsnappy1v5_1.1.4-1_amd64.deb ...
Unpacking libsnappy1v5:amd64 (1.1.4-1) ...
Setting up libnettle6:amd64 (3.3-1) ...
Setting up libpopt0:amd64 (1.16-10) ...
Setting up libexpat1:amd64 (2.2.0-2ubuntu0.1) ...
Setting up libboost-date-time1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libssl1.0.0:amd64 (1.0.2g-1ubuntu11.4) ...
Setting up mime-support (3.60ubuntu1) ...
Setting up libldap-common (2.4.44+dfsg-3ubuntu2.1) ...
Setting up libboost-program-options1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libsasl2-modules-db:amd64 (2.1.27~101-g0780600+dfsg-2ubuntu1) ...
Setting up libsasl2-2:amd64 (2.1.27~101-g0780600+dfsg-2ubuntu1) ...
Setting up libroken18-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Setting up libicu57:amd64 (57.1-5ubuntu0.2) ...
Setting up libkrb5support0:amd64 (1.15-1ubuntu0.1) ...
Setting up libnspr4:amd64 (2:4.13.1-0ubuntu0.17.04.1) ...
Setting up libkmod2:amd64 (22-1.1ubuntu1) ...
Setting up libtasn1-6:amd64 (4.10-1ubuntu0.1) ...
Setting up libboost-iostreams1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libtcmalloc-minimal4 (2.5-0ubuntu4) ...
Setting up libgmp10:amd64 (2:6.1.2+dfsg-1) ...
Setting up libunwind8 (1.1-4.1ubuntu2) ...
Processing triggers for libc-bin (2.24-9ubuntu2.2) ...
Setting up libaio1:amd64 (0.3.110-3) ...
Setting up libunistring0:amd64 (0.9.3-5.2ubuntu1) ...
Setting up libheimbase1-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Setting up libsnappy1v5:amd64 (1.1.4-1) ...
Setting up gdisk (1.0.1-1build1) ...
Setting up openssl (1.0.2g-1ubuntu11.4) ...
Setting up libboost-python1.62.0 (1.62.0+dfsg-4) ...
Setting up libffi6:amd64 (3.2.1-6) ...
Setting up libkeyutils1:amd64 (1.5.9-9ubuntu1) ...
Setting up libnl-3-200:amd64 (3.2.29-0ubuntu2.1) ...
Setting up ca-certificates (20170717~17.04.1) ...
Updating certificates in /etc/ssl/certs...
148 added, 0 removed; done.
Setting up libboost-context1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libpython2.7-stdlib:amd64 (2.7.13-2ubuntu0.1) ...
Setting up libidn11:amd64 (1.33-1ubuntu0.1) ...
Setting up libboost-system1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up kmod (22-1.1ubuntu1) ...
Setting up libk5crypto3:amd64 (1.15-1ubuntu0.1) ...
Setting up libidn2-0:amd64 (0.16-1ubuntu0.1) ...
Setting up libwind0-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Setting up libboost-thread1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libpsl5:amd64 (0.17.0-4) ...
Setting up libboost-random1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libasn1-8-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Setting up libboost-regex1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libhcrypto4-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Setting up python2.7 (2.7.13-2ubuntu0.1) ...
Setting up libnss3:amd64 (2:3.28.4-0ubuntu0.17.04.3) ...
Setting up libnl-route-3-200:amd64 (3.2.29-0ubuntu2.1) ...
Setting up libhx509-5-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Setting up libhogweed4:amd64 (3.3-1) ...
Setting up libpython-stdlib:amd64 (2.7.13-2) ...
Setting up libgoogle-perftools4 (2.5-0ubuntu4) ...
Setting up libp11-kit0:amd64 (0.23.3-5) ...
Setting up libboost-filesystem1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libpython2.7:amd64 (2.7.13-2ubuntu0.1) ...
Setting up libkrb5-3:amd64 (1.15-1ubuntu0.1) ...
Setting up libkrb5-26-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Setting up libheimntlm0-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Setting up python (2.7.13-2) ...
Setting up libibverbs1 (1.2.1-2ubuntu1) ...
Setting up python-prettytable (0.7.2-3) ...
Setting up libboost-coroutine1.62.0:amd64 (1.62.0+dfsg-4) ...
Setting up libgnutls30:amd64 (3.5.6-4ubuntu4.3) ...
Setting up librtmp1:amd64 (2.4+20151223.gitfa8646d.1-1) ...
Setting up libgssapi-krb5-2:amd64 (1.15-1ubuntu0.1) ...
Setting up libgssapi3-heimdal:amd64 (7.1.0+dfsg-9ubuntu1.1) ...
Setting up libldap-2.4-2:amd64 (2.4.44+dfsg-3ubuntu2.1) ...
Setting up libcurl3-gnutls:amd64 (7.52.1-4ubuntu1.4) ...
Processing triggers for libc-bin (2.24-9ubuntu2.2) ...
Processing triggers for ca-certificates (20170717~17.04.1) ...
Updating certificates in /etc/ssl/certs...
0 added, 0 removed; done.
Running hooks in /etc/ca-certificates/update.d...
done.
Reading package lists...
Building dependency tree...
Reading state information...
Calculating upgrade...
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
Reading package lists...
Building dependency tree...
Reading state information...
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
 ---> 7f808c2dee1f
Removing intermediate container 4f2248544f34
Step 4/6 : ADD ceph-install.tar /
 ---> 92dcf764da51
Removing intermediate container 796f8e5f2e0e
Step 5/6 : RUN ldconfig
 ---> Running in 67d75b0e4ee6
 ---> 459fb8456b58
Removing intermediate container 67d75b0e4ee6
Step 6/6 : ENV PYTHONPATH $PYTHONPATH:/usr/local/lib/python2.7/site-packages
 ---> Running in 054012fac726
 ---> e6f70c591548
Removing intermediate container 054012fac726
Successfully built e6f70c591548
=== docker build build-be9394dd/rook-amd64
Sending build context to Docker daemon 202.5 MB
Step 1/4 : FROM build-be9394dd/rook-base-amd64
 ---> e6f70c591548
Step 2/4 : ADD rook rookflex /usr/local/bin/
 ---> 2c5a2d58a206
Removing intermediate container d258639037b0
Step 3/4 : ENTRYPOINT /tini -- /usr/local/bin/rook
 ---> Running in 0461c6224602
 ---> 40a40173dae7
Removing intermediate container 0461c6224602
Step 4/4 : CMD 
 ---> Running in 781a50bdbc89
 ---> c6801d49dcde
Removing intermediate container 781a50bdbc89
Successfully built c6801d49dcde
=== saving image build-be9394dd/rook-amd64
make[2]: Entering directory '/windows10_drive_g/work/src/github.com/rook/rook/images/rook'
=== caching image build-be9394dd/rook-base-amd64
make[2]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/rook'
make[1]: Leaving directory '/windows10_drive_g/work/src/github.com/rook/rook/images/rook'
```

Include build for legacy ubuntu
```
[vagrant@localhost images]$ docker images
REPOSITORY                                                TAG                 IMAGE ID            CREATED             SIZE
build-be9394dd/rook-amd64                                 latest              c6801d49dcde        4 minutes ago       340 MB
build-be9394dd/rook-base-amd64                            latest              e6f70c591548        4 minutes ago       294 MB
cache/rook-base-amd64                                     2017-12-15.023131   e6f70c591548        4 minutes ago       294 MB
cache/ceph-amd64-tcmalloc-release-c19572dcbeb0            2017-12-15.022732   62d5472affcf        13 minutes ago      156 MB
build-be9394dd/ceph-amd64-tcmalloc-release-c19572dcbeb0   latest              62d5472affcf        13 minutes ago      156 MB
<none>                                                    <none>              3ba2d1f3021c        About an hour ago   1.94 GB
build-be9394dd/ceph-builder-amd64-tcmalloc-release        latest              c19572dcbeb0        3 hours ago         1.94 GB
cache/ceph-builder-amd64-tcmalloc-release                 2017-12-15.022732   c19572dcbeb0        3 hours ago         1.94 GB
<none>                                                    <none>              e3b8ec25d6e1        3 hours ago         1.94 GB
<none>                                                    <none>              b7d0de780ed5        3 hours ago         1.94 GB
cache/cross-gnu-amd64                                     2017-12-15.022730   c4a654332a52        4 hours ago         579 MB
build-be9394dd/cross-gnu-amd64                            latest              c4a654332a52        4 hours ago         579 MB
build-be9394dd/base-amd64                                 latest              bdc3253aab81        4 hours ago         54.8 MB
cache/base-amd64                                          2017-12-15.022729   bdc3253aab81        4 hours ago         54.8 MB
build-be9394dd/rootfs-amd64-8b040d721516                  latest              2e3b91214509        4 hours ago         54.7 MB
cache/rootfs-amd64-8b040d721516                           2017-12-15.022729   2e3b91214509        4 hours ago         54.7 MB
build-be9394dd/rootfs-builder-amd64                       latest              8b040d721516        4 hours ago         119 MB
cache/rootfs-builder-amd64                                2017-12-15.022729   8b040d721516        4 hours ago         119 MB
docker.io/ubuntu                                          zesty               fe1cc5b91830        5 hours ago         95.6 MB
<none>                                                    <none>              733201cea3e1        15 hours ago        1.94 GB
cache/cross-gnu-amd64                                     2017-12-14.122812   bb75f009f7af        16 hours ago        579 MB
cache/base-amd64                                          2017-12-14.122812   c9138ff4018e        16 hours ago        54.8 MB
build-be9394dd/rootfs-amd64-98e231c352a9                  latest              521831cd5775        16 hours ago        54.7 MB
cache/rootfs-amd64-98e231c352a9                           2017-12-14.122812   521831cd5775        16 hours ago        54.7 MB
cache/rootfs-builder-amd64                                2017-12-14.122812   98e231c352a9        16 hours ago        120 MB
docker.io/ubuntu                                          <none>              93fd6b1bc1dc        3 weeks ago         95.4 MB
```
