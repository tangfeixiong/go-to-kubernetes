### Tools for Kubernetes API extensions code-generator

**Note: refer to https://github.com/kubernetes/code-generator **

### Clone then build

client
```
[vagrant@kubedev-172-17-4-59 code-generator]$ GOPATH=/Users/fanhongling/go go install -v ./cmd/client-gen/
k8s.io/code-generator/vendor/github.com/golang/glog
k8s.io/code-generator/vendor/github.com/spf13/pflag
k8s.io/code-generator/vendor/k8s.io/gengo/types
k8s.io/code-generator/vendor/k8s.io/gengo/namer
k8s.io/code-generator/cmd/client-gen/types
k8s.io/code-generator/pkg/util
k8s.io/code-generator/vendor/golang.org/x/tools/go/ast/astutil
k8s.io/code-generator/vendor/golang.org/x/tools/imports
k8s.io/code-generator/vendor/k8s.io/gengo/parser
k8s.io/code-generator/vendor/k8s.io/gengo/generator
k8s.io/code-generator/vendor/k8s.io/gengo/args
k8s.io/code-generator/cmd/client-gen/args
k8s.io/code-generator/cmd/client-gen/path
k8s.io/code-generator/cmd/client-gen/generators/scheme
k8s.io/code-generator/cmd/client-gen/generators/util
k8s.io/code-generator/cmd/client-gen/generators/fake
k8s.io/code-generator/cmd/client-gen/generators
k8s.io/code-generator/cmd/client-gen
```

deepcopy
```
[vagrant@kubedev-172-17-4-59 code-generator]$ GOPATH=/Users/fanhongling/go go install -v ./cmd/deepcopy-gen/
k8s.io/code-generator/vendor/k8s.io/gengo/examples/set-gen/sets
k8s.io/code-generator/vendor/k8s.io/gengo/examples/deepcopy-gen/generators
k8s.io/code-generator/cmd/deepcopy-gen/args
k8s.io/code-generator/cmd/deepcopy-gen
```

informer
```
[vagrant@kubedev-172-17-4-59 code-generator]$ GOPATH=/Users/fanhongling/go go install -v ./cmd/informer-gen/
k8s.io/code-generator/cmd/informer-gen/args
k8s.io/code-generator/cmd/informer-gen/generators
k8s.io/code-generator/cmd/informer-gen
```

lister
```
[vagrant@kubedev-172-17-4-59 code-generator]$ GOPATH=/Users/fanhongling/go go install -v ./cmd/lister-gen/  
k8s.io/code-generator/cmd/lister-gen/args
k8s.io/code-generator/cmd/lister-gen/generators
k8s.io/code-generator/cmd/lister-gen
```

### Definitions

Build types.go
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go build -a -v ./pkg/apis/example.com/...
```

### Generate code

1st step
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ./hack/standalone-update-deepcopy.sh 
I0119 21:10:33.323764   27269 parse.go:418] type checker: /opt/go/src/runtime/trace.go:150:30: cannot use (traceBufHeader literal) (value of type traceBufHeader) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.327394   27269 parse.go:418] type checker: /opt/go/src/runtime/trace.go:150:7: array length 64 << 10 - unsafe.Sizeof((traceBufHeader literal)) (value of type uintptr) must be constant
I0119 21:10:33.328125   27269 parse.go:418] type checker: /opt/go/src/runtime/mgcwork.go:303:37: cannot use (workbufhdr literal) (value of type workbufhdr) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.328664   27269 parse.go:418] type checker: /opt/go/src/runtime/mgcwork.go:303:7: array length (_WorkbufSize - unsafe.Sizeof((workbufhdr literal))) / sys.PtrSize (value of type uintptr) must be constant
I0119 21:10:33.329844   27269 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:28: cannot use (hchan literal) (value of type hchan) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.330327   27269 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:66: cannot use (hchan literal) (value of type hchan) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.331164   27269 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:14: unsafe.Sizeof((hchan literal)) + uintptr(-int(unsafe.Sizeof((hchan literal))) & (maxAlign - 1)) (value of type uintptr) is not constant
I0119 21:10:33.333354   27269 parse.go:418] type checker: /opt/go/src/runtime/hashmap.go:80:31: cannot use (struct{b bmap; v int64} literal).v (value of type int64) as unsafe.ArbitraryType value in argument to unsafe.Offsetof
I0119 21:10:33.334568   27269 parse.go:418] type checker: /opt/go/src/runtime/hashmap.go:80:15: unsafe.Offsetof((struct{b bmap; v int64} literal).v) (value of type uintptr) is not constant
I0119 21:10:33.337845   27269 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:148:47: cannot use (mcentral literal) (value of type mcentral) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.338906   27269 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:148:13: array length sys.CacheLineSize - unsafe.Sizeof((mcentral literal)) % sys.CacheLineSize (value of type uintptr) must be constant
I0119 21:10:33.340831   27269 parse.go:418] type checker: /opt/go/src/runtime/mfinal.go:28:65: cannot use (finalizer literal) (value of type finalizer) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.341643   27269 parse.go:418] type checker: /opt/go/src/runtime/mfinal.go:28:11: array length (_FinBlockSize - 2 * sys.PtrSize - 2 * 4) / unsafe.Sizeof((finalizer literal)) (value of type uintptr) must be constant
I0119 21:10:33.343488   27269 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:1535:41: cannot use (gcBitsHeader literal) (value of type gcBitsHeader) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.344729   27269 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:1535:27: unsafe.Sizeof((gcBitsHeader literal)) (value of type uintptr) is not constant
I0119 21:10:33.346395   27269 parse.go:418] type checker: /opt/go/src/runtime/mstats.go:436:39: cannot use memstats.by_size (variable of type [67]struct{size uint32; nmalloc uint64; nfree uint64}) as unsafe.ArbitraryType value in argument to unsafe.Offsetof
I0119 21:10:33.347206   27269 parse.go:418] type checker: /opt/go/src/runtime/mstats.go:436:76: cannot use memstats.by_size[0] (variable of type struct{size uint32; nmalloc uint64; nfree uint64}) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.348681   27269 parse.go:418] type checker: /opt/go/src/runtime/panic.go:120:34: cannot use (_defer literal) (value of type _defer) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.349538   27269 parse.go:418] type checker: /opt/go/src/runtime/panic.go:120:20: unsafe.Sizeof((_defer literal)) (value of type uintptr) is not constant
I0119 21:10:33.351730   27269 parse.go:418] type checker: /opt/go/src/runtime/sema.go:51:42: cannot use (semaRoot literal) (value of type semaRoot) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.353334   27269 parse.go:418] type checker: /opt/go/src/runtime/sema.go:51:8: array length sys.CacheLineSize - unsafe.Sizeof((semaRoot literal)) (value of type uintptr) must be constant
I0119 21:10:33.363145   27269 parse.go:367] type checking encountered some issues in "runtime", but ignoring.
I0119 21:10:33.367451   27269 parse.go:418] type checker: /opt/go/src/sync/map.go:69:31: cannot convert new(interface{}) (value of type *interface{}) to unsafe.Pointer
I0119 21:10:33.368760   27269 parse.go:418] type checker: /opt/go/src/sync/pool.go:68:27: cannot use (poolLocalInternal literal) (value of type poolLocalInternal) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:33.371963   27269 parse.go:418] type checker: /opt/go/src/sync/pool.go:68:7: array length 128 - unsafe.Sizeof((poolLocalInternal literal)) % 128 (value of type uintptr) must be constant
I0119 21:10:33.373186   27269 parse.go:367] type checking encountered some issues in "sync", but ignoring.
I0119 21:10:36.044726   27269 parse.go:418] type checker: /opt/go/src/crypto/cipher/xor.go:12:36: cannot use uintptr(0) (constant 0 of type uintptr) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:10:36.046344   27269 parse.go:418] type checker: /opt/go/src/crypto/cipher/xor.go:12:18: int(unsafe.Sizeof(uintptr(0))) (value of type int) is not constant
I0119 21:10:36.047180   27269 parse.go:367] type checking encountered some issues in "crypto/cipher", but ignoring.
I0119 21:10:39.948173   27269 execute.go:217] Processing package "v1", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1"
I0119 21:10:39.951238   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.AddToScheme is not copyable
I0119 21:10:39.952184   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.ClusterMode is not copyable
I0119 21:10:39.952864   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.ConfigMap is not copyable
I0119 21:10:39.953106   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.Kind is not copyable
I0119 21:10:39.953223   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.Resource is not copyable
I0119 21:10:39.954524   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.SchemeBuilder is not copyable
I0119 21:10:39.955325   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.SchemeGroupVersion is not copyable
I0119 21:10:39.956100   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.ServerConditionType is not copyable
I0119 21:10:39.957288   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.ServerPhase is not copyable
I0119 21:10:39.958368   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.addKnownTypes is not copyable
I0119 21:10:39.959376   27269 deepcopy.go:270] Type github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1.localSchemeBuilder is not copyable
I0119 21:10:39.979444   27269 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1/zz_generated.deepcopy.go"
I0119 21:10:40.005153   27269 main.go:81] Completed successfully.
```

2nd step
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ./hack/standalone-update-clientset.sh 
I0119 21:14:19.327920   27940 parse.go:418] type checker: /opt/go/src/runtime/trace.go:150:30: cannot use (traceBufHeader literal) (value of type traceBufHeader) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.330806   27940 parse.go:418] type checker: /opt/go/src/runtime/trace.go:150:7: array length 64 << 10 - unsafe.Sizeof((traceBufHeader literal)) (value of type uintptr) must be constant
I0119 21:14:19.331593   27940 parse.go:418] type checker: /opt/go/src/runtime/mgcwork.go:303:37: cannot use (workbufhdr literal) (value of type workbufhdr) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.332297   27940 parse.go:418] type checker: /opt/go/src/runtime/mgcwork.go:303:7: array length (_WorkbufSize - unsafe.Sizeof((workbufhdr literal))) / sys.PtrSize (value of type uintptr) must be constant
I0119 21:14:19.333135   27940 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:28: cannot use (hchan literal) (value of type hchan) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.333612   27940 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:66: cannot use (hchan literal) (value of type hchan) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.334187   27940 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:14: unsafe.Sizeof((hchan literal)) + uintptr(-int(unsafe.Sizeof((hchan literal))) & (maxAlign - 1)) (value of type uintptr) is not constant
I0119 21:14:19.339682   27940 parse.go:418] type checker: /opt/go/src/runtime/hashmap.go:80:31: cannot use (struct{b bmap; v int64} literal).v (value of type int64) as unsafe.ArbitraryType value in argument to unsafe.Offsetof
I0119 21:14:19.340412   27940 parse.go:418] type checker: /opt/go/src/runtime/hashmap.go:80:15: unsafe.Offsetof((struct{b bmap; v int64} literal).v) (value of type uintptr) is not constant
I0119 21:14:19.342304   27940 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:148:47: cannot use (mcentral literal) (value of type mcentral) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.342944   27940 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:148:13: array length sys.CacheLineSize - unsafe.Sizeof((mcentral literal)) % sys.CacheLineSize (value of type uintptr) must be constant
I0119 21:14:19.344211   27940 parse.go:418] type checker: /opt/go/src/runtime/mfinal.go:28:65: cannot use (finalizer literal) (value of type finalizer) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.344835   27940 parse.go:418] type checker: /opt/go/src/runtime/mfinal.go:28:11: array length (_FinBlockSize - 2 * sys.PtrSize - 2 * 4) / unsafe.Sizeof((finalizer literal)) (value of type uintptr) must be constant
I0119 21:14:19.346104   27940 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:1535:41: cannot use (gcBitsHeader literal) (value of type gcBitsHeader) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.346535   27940 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:1535:27: unsafe.Sizeof((gcBitsHeader literal)) (value of type uintptr) is not constant
I0119 21:14:19.348404   27940 parse.go:418] type checker: /opt/go/src/runtime/mstats.go:436:39: cannot use memstats.by_size (variable of type [67]struct{size uint32; nmalloc uint64; nfree uint64}) as unsafe.ArbitraryType value in argument to unsafe.Offsetof
I0119 21:14:19.349317   27940 parse.go:418] type checker: /opt/go/src/runtime/mstats.go:436:76: cannot use memstats.by_size[0] (variable of type struct{size uint32; nmalloc uint64; nfree uint64}) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.350764   27940 parse.go:418] type checker: /opt/go/src/runtime/panic.go:120:34: cannot use (_defer literal) (value of type _defer) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.351412   27940 parse.go:418] type checker: /opt/go/src/runtime/panic.go:120:20: unsafe.Sizeof((_defer literal)) (value of type uintptr) is not constant
I0119 21:14:19.352887   27940 parse.go:418] type checker: /opt/go/src/runtime/sema.go:51:42: cannot use (semaRoot literal) (value of type semaRoot) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.353723   27940 parse.go:418] type checker: /opt/go/src/runtime/sema.go:51:8: array length sys.CacheLineSize - unsafe.Sizeof((semaRoot literal)) (value of type uintptr) must be constant
I0119 21:14:19.359121   27940 parse.go:367] type checking encountered some issues in "runtime", but ignoring.
I0119 21:14:19.828966   27940 parse.go:418] type checker: /opt/go/src/sync/map.go:69:31: cannot convert new(interface{}) (value of type *interface{}) to unsafe.Pointer
I0119 21:14:19.830018   27940 parse.go:418] type checker: /opt/go/src/sync/pool.go:68:27: cannot use (poolLocalInternal literal) (value of type poolLocalInternal) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:19.830693   27940 parse.go:418] type checker: /opt/go/src/sync/pool.go:68:7: array length 128 - unsafe.Sizeof((poolLocalInternal literal)) % 128 (value of type uintptr) must be constant
I0119 21:14:19.831664   27940 parse.go:367] type checking encountered some issues in "sync", but ignoring.
I0119 21:14:22.502660   27940 parse.go:418] type checker: /opt/go/src/crypto/cipher/xor.go:12:36: cannot use uintptr(0) (constant 0 of type uintptr) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:22.503436   27940 parse.go:418] type checker: /opt/go/src/crypto/cipher/xor.go:12:18: int(unsafe.Sizeof(uintptr(0))) (value of type int) is not constant
I0119 21:14:22.504327   27940 parse.go:367] type checking encountered some issues in "crypto/cipher", but ignoring.
I0119 21:14:26.714471   27940 parse.go:418] type checker: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/davecgh/go-spew/spew/bypass.go:34:26: cannot use (*byte)(nil) (value of type *byte) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:14:26.715306   27940 parse.go:418] type checker: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/vendor/github.com/davecgh/go-spew/spew/bypass.go:34:12: unsafe.Sizeof((*byte)(nil)) (value of type uintptr) is not constant
I0119 21:14:26.716607   27940 parse.go:367] type checking encountered some issues in "github.com/davecgh/go-spew/spew", but ignoring.
I0119 21:14:27.054469   27940 execute.go:217] Processing package "versioned", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned"
I0119 21:14:27.066724   27940 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/doc.go"
I0119 21:14:27.076942   27940 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/clientset.go"
I0119 21:14:27.086512   27940 execute.go:217] Processing package "scheme", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/scheme"
I0119 21:14:27.094124   27940 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/scheme/doc.go"
I0119 21:14:27.098791   27940 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/scheme/register.go"
I0119 21:14:27.102767   27940 execute.go:217] Processing package "v1", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/typed/example.com/v1"
I0119 21:14:27.113464   27940 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/typed/example.com/v1/example.com_client.go"
I0119 21:14:27.130269   27940 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/typed/example.com/v1/generated_expansion.go"
I0119 21:14:27.134204   27940 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/typed/example.com/v1/doc.go"
I0119 21:14:27.137806   27940 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/typed/example.com/v1/cluster.go"
I0119 21:14:27.144897   27940 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/clientset/versioned/typed/example.com/v1/redis.go"
```

3rd step
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ./hack/standalone-update-listers.sh   
I0119 21:19:07.080125   28755 parse.go:418] type checker: /opt/go/src/runtime/trace.go:150:30: cannot use (traceBufHeader literal) (value of type traceBufHeader) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.083137   28755 parse.go:418] type checker: /opt/go/src/runtime/trace.go:150:7: array length 64 << 10 - unsafe.Sizeof((traceBufHeader literal)) (value of type uintptr) must be constant
I0119 21:19:07.083954   28755 parse.go:418] type checker: /opt/go/src/runtime/mgcwork.go:303:37: cannot use (workbufhdr literal) (value of type workbufhdr) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.084992   28755 parse.go:418] type checker: /opt/go/src/runtime/mgcwork.go:303:7: array length (_WorkbufSize - unsafe.Sizeof((workbufhdr literal))) / sys.PtrSize (value of type uintptr) must be constant
I0119 21:19:07.086472   28755 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:28: cannot use (hchan literal) (value of type hchan) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.087158   28755 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:66: cannot use (hchan literal) (value of type hchan) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.087930   28755 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:14: unsafe.Sizeof((hchan literal)) + uintptr(-int(unsafe.Sizeof((hchan literal))) & (maxAlign - 1)) (value of type uintptr) is not constant
I0119 21:19:07.093090   28755 parse.go:418] type checker: /opt/go/src/runtime/hashmap.go:80:31: cannot use (struct{b bmap; v int64} literal).v (value of type int64) as unsafe.ArbitraryType value in argument to unsafe.Offsetof
I0119 21:19:07.094060   28755 parse.go:418] type checker: /opt/go/src/runtime/hashmap.go:80:15: unsafe.Offsetof((struct{b bmap; v int64} literal).v) (value of type uintptr) is not constant
I0119 21:19:07.096342   28755 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:148:47: cannot use (mcentral literal) (value of type mcentral) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.097563   28755 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:148:13: array length sys.CacheLineSize - unsafe.Sizeof((mcentral literal)) % sys.CacheLineSize (value of type uintptr) must be constant
I0119 21:19:07.099434   28755 parse.go:418] type checker: /opt/go/src/runtime/mfinal.go:28:65: cannot use (finalizer literal) (value of type finalizer) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.100408   28755 parse.go:418] type checker: /opt/go/src/runtime/mfinal.go:28:11: array length (_FinBlockSize - 2 * sys.PtrSize - 2 * 4) / unsafe.Sizeof((finalizer literal)) (value of type uintptr) must be constant
I0119 21:19:07.102486   28755 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:1535:41: cannot use (gcBitsHeader literal) (value of type gcBitsHeader) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.103383   28755 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:1535:27: unsafe.Sizeof((gcBitsHeader literal)) (value of type uintptr) is not constant
I0119 21:19:07.104794   28755 parse.go:418] type checker: /opt/go/src/runtime/mstats.go:436:39: cannot use memstats.by_size (variable of type [67]struct{size uint32; nmalloc uint64; nfree uint64}) as unsafe.ArbitraryType value in argument to unsafe.Offsetof
I0119 21:19:07.105904   28755 parse.go:418] type checker: /opt/go/src/runtime/mstats.go:436:76: cannot use memstats.by_size[0] (variable of type struct{size uint32; nmalloc uint64; nfree uint64}) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.108447   28755 parse.go:418] type checker: /opt/go/src/runtime/panic.go:120:34: cannot use (_defer literal) (value of type _defer) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.109209   28755 parse.go:418] type checker: /opt/go/src/runtime/panic.go:120:20: unsafe.Sizeof((_defer literal)) (value of type uintptr) is not constant
I0119 21:19:07.111173   28755 parse.go:418] type checker: /opt/go/src/runtime/sema.go:51:42: cannot use (semaRoot literal) (value of type semaRoot) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.112216   28755 parse.go:418] type checker: /opt/go/src/runtime/sema.go:51:8: array length sys.CacheLineSize - unsafe.Sizeof((semaRoot literal)) (value of type uintptr) must be constant
I0119 21:19:07.121946   28755 parse.go:367] type checking encountered some issues in "runtime", but ignoring.
I0119 21:19:07.130509   28755 parse.go:418] type checker: /opt/go/src/sync/map.go:69:31: cannot convert new(interface{}) (value of type *interface{}) to unsafe.Pointer
I0119 21:19:07.131229   28755 parse.go:418] type checker: /opt/go/src/sync/pool.go:68:27: cannot use (poolLocalInternal literal) (value of type poolLocalInternal) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:07.131712   28755 parse.go:418] type checker: /opt/go/src/sync/pool.go:68:7: array length 128 - unsafe.Sizeof((poolLocalInternal literal)) % 128 (value of type uintptr) must be constant
I0119 21:19:07.132234   28755 parse.go:367] type checking encountered some issues in "sync", but ignoring.
I0119 21:19:10.088380   28755 parse.go:418] type checker: /opt/go/src/crypto/cipher/xor.go:12:36: cannot use uintptr(0) (constant 0 of type uintptr) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:19:10.089213   28755 parse.go:418] type checker: /opt/go/src/crypto/cipher/xor.go:12:18: int(unsafe.Sizeof(uintptr(0))) (value of type int) is not constant
I0119 21:19:10.089800   28755 parse.go:367] type checking encountered some issues in "crypto/cipher", but ignoring.
I0119 21:19:13.937463   28755 execute.go:217] Processing package "v1", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/listers/example.com/v1"
I0119 21:19:13.960435   28755 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/listers/example.com/v1/expansion_generated.go"
I0119 21:19:13.971340   28755 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/listers/example.com/v1/cluster.go"
I0119 21:19:14.049022   28755 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/listers/example.com/v1/redis.go"
I0119 21:19:14.120489   28755 main.go:56] Completed successfully.
```

4th
```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ./hack/standalone-update-informers.sh 
I0119 21:21:58.573857   29226 parse.go:418] type checker: /opt/go/src/runtime/trace.go:150:30: cannot use (traceBufHeader literal) (value of type traceBufHeader) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.577170   29226 parse.go:418] type checker: /opt/go/src/runtime/trace.go:150:7: array length 64 << 10 - unsafe.Sizeof((traceBufHeader literal)) (value of type uintptr) must be constant
I0119 21:21:58.578467   29226 parse.go:418] type checker: /opt/go/src/runtime/mgcwork.go:303:37: cannot use (workbufhdr literal) (value of type workbufhdr) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.579161   29226 parse.go:418] type checker: /opt/go/src/runtime/mgcwork.go:303:7: array length (_WorkbufSize - unsafe.Sizeof((workbufhdr literal))) / sys.PtrSize (value of type uintptr) must be constant
I0119 21:21:58.580838   29226 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:28: cannot use (hchan literal) (value of type hchan) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.581839   29226 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:66: cannot use (hchan literal) (value of type hchan) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.582498   29226 parse.go:418] type checker: /opt/go/src/runtime/chan.go:27:14: unsafe.Sizeof((hchan literal)) + uintptr(-int(unsafe.Sizeof((hchan literal))) & (maxAlign - 1)) (value of type uintptr) is not constant
I0119 21:21:58.585562   29226 parse.go:418] type checker: /opt/go/src/runtime/hashmap.go:80:31: cannot use (struct{b bmap; v int64} literal).v (value of type int64) as unsafe.ArbitraryType value in argument to unsafe.Offsetof
I0119 21:21:58.586644   29226 parse.go:418] type checker: /opt/go/src/runtime/hashmap.go:80:15: unsafe.Offsetof((struct{b bmap; v int64} literal).v) (value of type uintptr) is not constant
I0119 21:21:58.589529   29226 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:148:47: cannot use (mcentral literal) (value of type mcentral) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.590421   29226 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:148:13: array length sys.CacheLineSize - unsafe.Sizeof((mcentral literal)) % sys.CacheLineSize (value of type uintptr) must be constant
I0119 21:21:58.592061   29226 parse.go:418] type checker: /opt/go/src/runtime/mfinal.go:28:65: cannot use (finalizer literal) (value of type finalizer) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.592939   29226 parse.go:418] type checker: /opt/go/src/runtime/mfinal.go:28:11: array length (_FinBlockSize - 2 * sys.PtrSize - 2 * 4) / unsafe.Sizeof((finalizer literal)) (value of type uintptr) must be constant
I0119 21:21:58.594371   29226 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:1535:41: cannot use (gcBitsHeader literal) (value of type gcBitsHeader) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.595284   29226 parse.go:418] type checker: /opt/go/src/runtime/mheap.go:1535:27: unsafe.Sizeof((gcBitsHeader literal)) (value of type uintptr) is not constant
I0119 21:21:58.601750   29226 parse.go:418] type checker: /opt/go/src/runtime/mstats.go:436:39: cannot use memstats.by_size (variable of type [67]struct{size uint32; nmalloc uint64; nfree uint64}) as unsafe.ArbitraryType value in argument to unsafe.Offsetof
I0119 21:21:58.602990   29226 parse.go:418] type checker: /opt/go/src/runtime/mstats.go:436:76: cannot use memstats.by_size[0] (variable of type struct{size uint32; nmalloc uint64; nfree uint64}) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.605202   29226 parse.go:418] type checker: /opt/go/src/runtime/panic.go:120:34: cannot use (_defer literal) (value of type _defer) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.608670   29226 parse.go:418] type checker: /opt/go/src/runtime/panic.go:120:20: unsafe.Sizeof((_defer literal)) (value of type uintptr) is not constant
I0119 21:21:58.618418   29226 parse.go:418] type checker: /opt/go/src/runtime/sema.go:51:42: cannot use (semaRoot literal) (value of type semaRoot) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.619190   29226 parse.go:418] type checker: /opt/go/src/runtime/sema.go:51:8: array length sys.CacheLineSize - unsafe.Sizeof((semaRoot literal)) (value of type uintptr) must be constant
I0119 21:21:58.632771   29226 parse.go:367] type checking encountered some issues in "runtime", but ignoring.
I0119 21:21:58.635430   29226 parse.go:418] type checker: /opt/go/src/sync/map.go:69:31: cannot convert new(interface{}) (value of type *interface{}) to unsafe.Pointer
I0119 21:21:58.636348   29226 parse.go:418] type checker: /opt/go/src/sync/pool.go:68:27: cannot use (poolLocalInternal literal) (value of type poolLocalInternal) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:21:58.636815   29226 parse.go:418] type checker: /opt/go/src/sync/pool.go:68:7: array length 128 - unsafe.Sizeof((poolLocalInternal literal)) % 128 (value of type uintptr) must be constant
I0119 21:21:58.638336   29226 parse.go:367] type checking encountered some issues in "sync", but ignoring.
I0119 21:22:01.512009   29226 parse.go:418] type checker: /opt/go/src/crypto/cipher/xor.go:12:36: cannot use uintptr(0) (constant 0 of type uintptr) as unsafe.ArbitraryType value in argument to unsafe.Sizeof
I0119 21:22:01.512927   29226 parse.go:418] type checker: /opt/go/src/crypto/cipher/xor.go:12:18: int(unsafe.Sizeof(uintptr(0))) (value of type int) is not constant
I0119 21:22:01.513936   29226 parse.go:367] type checking encountered some issues in "crypto/cipher", but ignoring.
I0119 21:22:05.041225   29226 execute.go:217] Processing package "v1", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/example.com/v1"
I0119 21:22:05.051805   29226 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/example.com/v1/interface.go"
I0119 21:22:05.061164   29226 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/example.com/v1/cluster.go"
I0119 21:22:05.068459   29226 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/example.com/v1/redis.go"
I0119 21:22:05.073129   29226 execute.go:217] Processing package "internalinterfaces", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/internalinterfaces"
I0119 21:22:05.075210   29226 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/internalinterfaces/factory_interfaces.go"
I0119 21:22:05.088999   29226 execute.go:217] Processing package "externalversions", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions"
I0119 21:22:05.093212   29226 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/factory.go"
I0119 21:22:05.101981   29226 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/generic.go"
I0119 21:22:05.106102   29226 execute.go:217] Processing package "example", disk location "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/example.com"
I0119 21:22:05.108291   29226 execute.go:67] Assembling file "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/client/informers/externalversions/example.com/interface.go"
I0119 21:22:05.111668   29226 main.go:59] Completed successfully.
```