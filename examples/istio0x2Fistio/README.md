# Instruction

Bazel
```
[vagrant@localhost bazel]$ sudo dnf list | grep bazel
bazel.src                                0.7.0-1.el7.centos              vbatts-bazel
bazel.x86_64                             0.7.0-1.el7.centos              vbatts-bazel
[vagrant@localhost bazel]$ sudo dnf install -y bazel
Copr repo for bazel owned by vbatts                                                       920  B/s | 1.4 kB     00:01    
上次元数据过期检查在 0:00:00 前执行于 Sat Nov 11 08:59:31 2017。
依赖关系解决。
==========================================================================================================================
 Package                架构                    版本                                  仓库                           大小
==========================================================================================================================
安装:
 bazel                  x86_64                  0.7.0-1.el7.centos                    vbatts-bazel                   86 M

事务概要
==========================================================================================================================
安装  1 Package

总下载：86 M
安装大小：88 M
下载软件包：
bazel-0.7.0-1.el7.centos.x86_64.rpm                                                       145 kB/s |  86 MB     10:06    
--------------------------------------------------------------------------------------------------------------------------
总计                                                                                      145 kB/s |  86 MB     10:06     
警告：/var/cache/dnf/vbatts-bazel-dfddf3a1bf312c3c/packages/bazel-0.7.0-1.el7.centos.x86_64.rpm: 头V3 RSA/SHA1 Signature, 密钥 ID eb2be214: NOKEY
导入 GPG 公钥 0xEB2BE214:
 Userid: "vbatts_bazel (None) <vbatts#bazel@copr.fedorahosted.org>"
 指纹: 090F 9C8B BDB6 3200 807E 16C2 978A 4B98 EB2B E214
 来自: https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/pubkey.gpg
导入公钥成功
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: bazel-0.7.0-1.el7.centos.x86_64                                                                               1/1 
  验证: bazel-0.7.0-1.el7.centos.x86_64                                                                               1/1 

已安装:
  bazel.x86_64 0.7.0-1.el7.centos                                                                                         

完毕！
[vagrant@localhost bazel]$ bazel version
Extracting Bazel installation...
.................
Build label: 0.7.0- (@non-git)
Build target: bazel-out/local-opt/bin/src/main/java/com/google/devtools/build/lib/bazel/BazelServer_deploy.jar
Build time: Wed Oct 18 15:28:47 2017 (1508340527)
Build timestamp: 1508340527
Build timestamp as int: 1508340527
[vagrant@localhost bazel]$ bazel help
                                               [bazel release 0.7.0- (@non-git)]
Usage: bazel <command> <options> ...

Available commands:
  analyze-profile     Analyzes build profile data.
  build               Builds the specified targets.
  canonicalize-flags  Canonicalizes a list of bazel options.
  clean               Removes output files and optionally stops the server.
  coverage            Generates code coverage report for specified test targets.
  dump                Dumps the internal state of the bazel server process.
  fetch               Fetches external repositories that are prerequisites to the targets.
  help                Prints help for commands, or the index.
  info                Displays runtime info about the bazel server.
  license             Prints the license of this software.
  mobile-install      Installs targets to mobile devices.
  print_action        Prints the command line args for compiling a file.
  query               Executes a dependency graph query.
  run                 Runs the specified target.
  shutdown            Stops the bazel server.
  test                Builds and runs the specified test targets.
  version             Prints version information for bazel.

Getting more help:
  bazel help <command>
                   Prints help and options for <command>.
  bazel help startup_options
                   Options for the JVM hosting bazel.
  bazel help target-syntax
                   Explains the syntax for specifying targets.
  bazel help info-keys
                   Displays a list of keys used by the info command.
```

Golang
```
[vagrant@localhost istio]$ go version
go version go1.9.2 linux/amd64
```

## Build

Issue
```
[vagrant@localhost istio]$ make build
ln -s ~/.kube/config pilot/platform/kube/
bazel  build  //...
.....................................................
ERROR: /Users/fanhongling/go/src/github.com/istio/istio/mixer/tools/codegen/pkg/inventory/BUILD:3:1: no such package '@org_golang_x_tools//imports': java.io.IOException: thread interrupted and referenced by '//mixer/tools/codegen/pkg/inventory:go_default_library'.
ERROR: Analysis of target '//mixer/tools/codegen/cmd/mixgeninventory:go_default_library' failed; build aborted: no such package '@org_golang_x_tools//imports': java.io.IOException: thread interrupted.
INFO: Elapsed time: 477.959s
Makefile:45: recipe for target 'build' failed
make: *** [build] Error 1
```

### 2017-11-11

Porject
```
[vagrant@localhost istio]$ git pull --append
remote: Counting objects: 286, done.
remote: Compressing objects: 100% (189/189), done.
remote: Total 286 (delta 195), reused 181 (delta 95), pack-reused 0
接收对象中: 100% (286/286), 80.52 KiB | 25.00 KiB/s, 完成.
处理 delta 中: 100% (195/195), 完成 104 个本地对象.
来自 https://github.com/istio/istio
   817f218..2606e68  master     -> origin/master
更新 817f218..2606e68
error: Your local changes to the following files would be overwritten by merge:
	generated_files
	lintconfig.json
Please, commit your changes or stash them before you can merge.
Aborting
```

```
[vagrant@localhost istio]$ git checkout -- generated_files lintconfig.json
[vagrant@localhost istio]$ git pull
更新 817f218..2606e68
Fast-forward
 .circleci/config.yml                                           |   12 +-
 CONTRIBUTING.md                                                |    2 +-
 DEV-GUIDE.md                                                   |   20 +-
 GROUPS.md                                                      |  146 ++++-
 Gopkg.lock                                                     |    6 +-
 WORKSPACE                                                      |    2 +
 broker/OWNERS                                                  |    6 +
 devel/githubContrib/Contributions.txt                          |    2 -
 generated_files                                                |    2 +
 install/kubernetes/istio-auth.yaml                             |    2 +-
 install/kubernetes/istio-one-namespace-auth.yaml               |    2 +-
 install/kubernetes/istio-one-namespace.yaml                    |    2 +-
 install/kubernetes/istio.yaml                                  |    2 +-
 install/kubernetes/templates/istio-pilot.yaml.tmpl             |    2 +-
 lintconfig.json                                                |    2 +
 mixer/adapter/kubernetes/kubernetes.go                         |   27 +-
 mixer/adapter/svcctrl/template/svcctrlreport/BUILD             |    8 +
 mixer/adapter/svcctrl/template/svcctrlreport/template.proto    |   68 +++
 mixer/istio_api.bzl                                            |   44 +-
 mixer/pkg/adapter/BUILD                                        |    2 +
 mixer/pkg/adapter/requestData.go                               |   52 ++
 mixer/pkg/adapter/requestData_test.go                          |   68 +++
 mixer/pkg/api/grpcServer.go                                    |   25 +-
 mixer/pkg/attribute/protoBag.go                                |    4 +-
 mixer/pkg/config/runtime.go                                    |    4 +-
 mixer/pkg/il/compiler/BUILD                                    |    1 +
 mixer/pkg/il/compiler/compiler_test.go                         | 1266 ++------------------------------------
 mixer/pkg/il/evaluator/BUILD                                   |    1 +
 mixer/pkg/il/evaluator/evaluator.go                            |   57 +-
 mixer/pkg/il/evaluator/evaluator_test.go                       |  223 +++----
 mixer/pkg/il/runtime/BUILD                                     |   21 +
 mixer/pkg/il/runtime/externs.go                                |   70 +++
 mixer/pkg/il/runtime/externs_test.go                           |  107 ++++
 mixer/pkg/il/testing/BUILD                                     |   12 +-
 mixer/pkg/il/testing/tests.go                                  | 1638 ++++++++++++++++++++++++++++++++++++++++++++++++++
 mixer/pkg/il/testing/util.go                                   |   32 +
 mixer/pkg/runtime/BUILD                                        |    2 +
 mixer/pkg/runtime/context.go                                   |   32 +
 mixer/pkg/runtime/context_test.go                              |   83 +++
 mixer/pkg/runtime/dispatcher.go                                |   11 +-
 mixer/pkg/runtime/dispatcher_test.go                           |   12 +-
 mixer/pkg/runtime/init.go                                      |    2 +-
 mixer/pkg/status/status.go                                     |    4 +-
 mixer/template/BUILD                                           |    8 +-
 mixer/template/template.gen.go                                 |  192 ++++++
 mixer/template/tracespan/BUILD.bazel                           |    8 +
 mixer/template/tracespan/go_default_library_handler.gen.go     |  143 +++++
 mixer/template/tracespan/go_default_library_tmpl.pb.go         | 1197 ++++++++++++++++++++++++++++++++++++
 mixer/template/tracespan/tracespan.proto                       |  103 ++++
 pilot/adapter/config/file/config.go                            |   19 +-
 pilot/adapter/config/file/config_test.go                       |    4 +-
 pilot/cmd/istioctl/inject.go                                   |   10 +-
 pilot/cmd/istioctl/main.go                                     |    9 +-
 pilot/platform/kube/inject/inject.go                           |   28 +-
 pilot/platform/kube/inject/inject_test.go                      |    5 +
 pilot/platform/kube/inject/testdata/cronjob.yaml               |   18 +
 pilot/platform/kube/inject/testdata/cronjob.yaml.injected      |  110 ++++
 pilot/proxy/envoy/BUILD                                        |    2 +
 pilot/proxy/envoy/config.go                                    |   16 +-
 pilot/proxy/envoy/ingress.go                                   |    2 +-
 pilot/proxy/envoy/mixer.go                                     |  109 +++-
 pilot/proxy/envoy/route.go                                     |    6 +-
 pilot/proxy/envoy/testdata/lds-httpproxy.json.golden           |   36 +-
 pilot/proxy/envoy/testdata/lds-ingress.json.golden             |   50 +-
 pilot/proxy/envoy/testdata/lds-router-auth.json.golden         |   75 ++-
 pilot/proxy/envoy/testdata/lds-router.json.golden              |   75 ++-
 pilot/proxy/envoy/testdata/lds-v0-egress-rule-auth.json.golden |  254 +++++++-
 pilot/proxy/envoy/testdata/lds-v0-egress-rule.json.golden      |  254 +++++++-
 pilot/proxy/envoy/testdata/lds-v0-fault-auth.json.golden       |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v0-fault.json.golden            |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v0-none-auth-optin.json.golden  |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v0-none-auth-optout.json.golden |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v0-none-auth.json.golden        |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v0-none.json.golden             |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v0-weighted-auth.json.golden    |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v0-weighted.json.golden         |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v1-egress-rule-auth.json.golden |  254 +++++++-
 pilot/proxy/envoy/testdata/lds-v1-egress-rule.json.golden      |  254 +++++++-
 pilot/proxy/envoy/testdata/lds-v1-fault-auth.json.golden       |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v1-fault.json.golden            |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v1-none-auth.json.golden        |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v1-none.json.golden             |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v1-weighted-auth.json.golden    |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-v1-weighted.json.golden         |  218 ++++++-
 pilot/proxy/envoy/testdata/lds-websocket.json.golden           |  220 ++++++-
 pilot/proxy/envoy/testdata/rds-ingress-ssl.json.golden         |    1 +
 pilot/proxy/envoy/testdata/rds-ingress-weighted.json.golden    |    1 +
 pilot/proxy/envoy/testdata/rds-ingress.json.golden             |    1 +
 prow/istio-presubmit.sh                                        |    2 +
 release/README.md                                              |    2 +
 security/.gitignore                                            |    4 +
 security/bin/e2e.sh                                            |   15 +-
 security/bin/push-docker                                       |   74 ++-
 security/cmd/node_agent/main.go                                |   21 +-
 security/cmd/node_agent/na/nafactory.go                        |    2 +-
 security/cmd/node_agent/na/nodeagent.go                        |    6 +-
 security/cmd/node_agent/na/nodeagent_test.go                   |   46 +-
 security/docker/BUILD                                          |   23 -
 security/docker/Dockerfile.istio-ca                            |    6 +
 security/docker/Dockerfile.istio-ca-test                       |   16 +
 security/docker/Dockerfile.node-agent-test                     |   11 +
 security/docker/app.js                                         |   57 ++
 security/docker/start_app.sh                                   |   18 +
 security/integration/BUILD                                     |    1 +
 security/integration/main.go                                   |  404 +++++++------
 security/integration/utils/BUILD                               |   24 +
 security/integration/utils/kubernetes.go                       |  283 +++++++++
 security/pkg/cmd/BUILD                                         |   14 +-
 security/pkg/cmd/flag_test.go                                  |   46 ++
 security/pkg/pki/ca/generate_cert.go                           |   19 +-
 security/pkg/pki/san.go                                        |   12 +-
 security/pkg/pki/san_test.go                                   |    3 +-
 security/pkg/platform/aws.go                                   |   18 +-
 security/pkg/platform/client.go                                |   19 +-
 security/pkg/platform/gcp.go                                   |   20 +-
 security/pkg/platform/gcp_test.go                              |   17 +-
 security/pkg/platform/mock/fakeclient.go                       |   23 +-
 security/pkg/platform/onprem.go                                |   31 +-
 security/pkg/platform/onprem_test.go                           |   16 +-
 security/pkg/server/grpc/authenticator.go                      |   59 +-
 security/pkg/server/grpc/authenticator_test.go                 |   28 +-
 security/pkg/server/grpc/authorizer.go                         |   23 +-
 security/pkg/server/grpc/authorizer_test.go                    |   30 +-
 security/pkg/server/grpc/server.go                             |   40 +-
 security/pkg/server/grpc/server_test.go                        |   93 +--
 {devel => tools}/githubContrib/BUILD.bazel                     |    0
 tools/githubContrib/Contributions.txt                          |    2 +
 {devel => tools}/githubContrib/githubContrib.go                |    0
 {devel => tools}/githubContrib/githubContrib_test.go           |    0
 {devel => tools}/minikube.md                                   |    0
 {devel => tools}/rules.yml                                     |    0
 {devel => tools}/setup_run                                     |    6 +-
 {devel => tools}/update_all                                    |    6 +-
 133 files changed, 10080 insertions(+), 2062 deletions(-)
 create mode 100644 broker/OWNERS
 delete mode 100644 devel/githubContrib/Contributions.txt
 create mode 100644 mixer/adapter/svcctrl/template/svcctrlreport/BUILD
 create mode 100644 mixer/adapter/svcctrl/template/svcctrlreport/template.proto
 create mode 100644 mixer/pkg/adapter/requestData.go
 create mode 100644 mixer/pkg/adapter/requestData_test.go
 create mode 100644 mixer/pkg/il/runtime/BUILD
 create mode 100644 mixer/pkg/il/runtime/externs.go
 create mode 100644 mixer/pkg/il/runtime/externs_test.go
 create mode 100644 mixer/pkg/il/testing/tests.go
 create mode 100644 mixer/pkg/il/testing/util.go
 create mode 100644 mixer/pkg/runtime/context.go
 create mode 100644 mixer/pkg/runtime/context_test.go
 create mode 100644 mixer/template/tracespan/BUILD.bazel
 create mode 100644 mixer/template/tracespan/go_default_library_handler.gen.go
 create mode 100644 mixer/template/tracespan/go_default_library_tmpl.pb.go
 create mode 100644 mixer/template/tracespan/tracespan.proto
 create mode 100644 pilot/platform/kube/inject/testdata/cronjob.yaml
 create mode 100644 pilot/platform/kube/inject/testdata/cronjob.yaml.injected
 delete mode 100644 security/docker/BUILD
 create mode 100644 security/docker/Dockerfile.istio-ca
 create mode 100644 security/docker/Dockerfile.istio-ca-test
 create mode 100644 security/docker/Dockerfile.node-agent-test
 create mode 100644 security/docker/app.js
 create mode 100644 security/docker/start_app.sh
 create mode 100644 security/integration/utils/BUILD
 create mode 100644 security/integration/utils/kubernetes.go
 create mode 100644 security/pkg/cmd/flag_test.go
 rename {devel => tools}/githubContrib/BUILD.bazel (100%)
 create mode 100644 tools/githubContrib/Contributions.txt
 rename {devel => tools}/githubContrib/githubContrib.go (100%)
 rename {devel => tools}/githubContrib/githubContrib_test.go (100%)
 rename {devel => tools}/minikube.md (100%)
 rename {devel => tools}/rules.yml (100%)
 rename {devel => tools}/setup_run (64%)
 rename {devel => tools}/update_all (60%)
```

```
[vagrant@localhost istio]$ make build
ln -s ~/.kube/config pilot/platform/kube/
ln: 无法创建符号链接"pilot/platform/kube/config": 文件已存在
Makefile:68: recipe for target 'pilot/platform/kube/config' failed
make: *** [pilot/platform/kube/config] Error 1
[vagrant@localhost istio]$ bazel build //...
.
ERROR: /Users/fanhongling/go/src/github.com/istio/istio/pilot/proxy/BUILD:3:1: no such package '@org_golang_x_time//rate': failed to fetch org_golang_x_time: 2017/11/11 11:04:33 unrecognized import path "golang.org/x/time"
 and referenced by '//pilot/proxy:go_default_library'.
ERROR: Analysis of target '//pilot/proxy:go_default_library' failed; build aborted: no such package '@org_golang_x_time//rate': failed to fetch org_golang_x_time: 2017/11/11 11:04:33 unrecognized import path "golang.org/x/time"
.
INFO: Elapsed time: 585.238s
```


