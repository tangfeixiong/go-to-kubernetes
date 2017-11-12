# Instruction

## Build 1st

### 2017-11-09

Project
```
[vagrant@localhost github.com]$ git clone --depth=1 https://github.com/bazelbuild/bazel bazelbuild/bazel
正克隆到 'bazelbuild/bazel'...
remote: Counting objects: 10211, done.
remote: Compressing objects: 100% (7679/7679), done.
remote: Total 10211 (delta 3168), reused 5593 (delta 2105), pack-reused 0
接收对象中: 100% (10211/10211), 76.69 MiB | 1.11 MiB/s, 完成.
处理 delta 中: 100% (3168/3168), 完成.
检查连接... 完成。
正在检出文件: 100% (9016/9016), 完成.
```

Issue
```
[vagrant@localhost bazel]$ ./compile.sh 
🍃  Building Bazel from scratch
ERROR: Must specify PROTOC if not bootstrapping from the distribution artifact

--------------------------------------------------------------------------------
NOTE: This failure is likely occuring if you are trying to bootstrap bazel from
a developer checkout. Those checkouts do not include the generated output of
the protoc compiler (as we prefer not to version generated files).

* To build a developer version of bazel, do

    bazel build //src:bazel

* To bootstrap your first bazel binary, please download a dist archive from our
  release page at https://github.com/bazelbuild/bazel/releases and run
  compile.sh on the unpacked archive.

The full install instructions to install a release version of bazel can be found
at https://docs.bazel.build/install-compile-source.html
For a rationale, why the bootstrap process is organized in this way, see
https://bazel.build/designs/2016/10/11/distribution-artifact.html
--------------------------------------------------------------------------------

```

Resolve
```
[vagrant@localhost bazel]$ /usr/local/bin/protoc --version
libprotoc 3.3.0
[vagrant@localhost bazel]$ export PROTOC=/usr/local/bin/protoc
```

Issue
```
[vagrant@localhost bazel]$ ./compile.sh 
🍃  Building Bazel from scratch
ERROR: Must specify GRPC_JAVA_PLUGIN if not bootstrapping from the distribution artifact

--------------------------------------------------------------------------------
NOTE: This failure is likely occuring if you are trying to bootstrap bazel from
a developer checkout. Those checkouts do not include the generated output of
the protoc compiler (as we prefer not to version generated files).

* To build a developer version of bazel, do

    bazel build //src:bazel

* To bootstrap your first bazel binary, please download a dist archive from our
  release page at https://github.com/bazelbuild/bazel/releases and run
  compile.sh on the unpacked archive.

The full install instructions to install a release version of bazel can be found
at https://docs.bazel.build/install-compile-source.html
For a rationale, why the bootstrap process is organized in this way, see
https://bazel.build/designs/2016/10/11/distribution-artifact.html
--------------------------------------------------------------------------------

```

plugin
```
[vagrant@localhost bazel]$ ldd ../../grpc/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java 
	linux-vdso.so.1 (0x00007ffcd25f4000)
	libpthread.so.0 => /lib64/libpthread.so.0 (0x00007fb691199000)
	libm.so.6 => /lib64/libm.so.6 (0x00007fb690e97000)
	libc.so.6 => /lib64/libc.so.6 (0x00007fb690ad5000)
	/lib64/ld-linux-x86-64.so.2 (0x0000562ed664f000)
[vagrant@localhost bazel]$ ldd ../../grpc/grpc-java/compiler/build/artifacts/java_plugin/protoc-gen-grpc-java.exe 
	linux-vdso.so.1 (0x00007ffd06d6d000)
	libpthread.so.0 => /lib64/libpthread.so.0 (0x00007f2558557000)
	libm.so.6 => /lib64/libm.so.6 (0x00007f2558255000)
	libc.so.6 => /lib64/libc.so.6 (0x00007f2557e93000)
	/lib64/ld-linux-x86-64.so.2 (0x00005559550bc000)
[vagrant@localhost grpc-java]$ ldd ~/.m2/repository/io/grpc/protoc-gen-grpc-java/1.4.0-SNAPSHOT/protoc-gen-grpc-java-1.4.0-SNAPSHOT-linux-x86_64.exe
	linux-vdso.so.1 (0x00007ffead5e6000)
	libpthread.so.0 => /lib64/libpthread.so.0 (0x00007fd43a7b5000)
	libm.so.6 => /lib64/libm.so.6 (0x00007fd43a4b3000)
	libc.so.6 => /lib64/libc.so.6 (0x00007fd43a0f1000)
	/lib64/ld-linux-x86-64.so.2 (0x000055c7dcf9d000)
[vagrant@localhost bazel]$ export GRPC_JAVA_PLUGIN=/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/compiler/build/artifacts/java_plugin/protoc-gen-grpc-java.exe
```

Issue
```
[vagrant@localhost bazel]$ ./compile.sh 
🍃  Building Bazel from scratch./usr/local/bin/protoc -Isrc/main/protobuf/ -Isrc/main/java/com/google/devtools/build/lib/buildeventstream/proto/ --java_out=/tmp/bazel_6j6UtjtB/src --plugin=protoc-gen-grpc=/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/compiler/build/artifacts/java_plugin/protoc-gen-grpc-java.exe --grpc_out=/tmp/bazel_6j6UtjtB/src src/main/java/com/google/devtools/build/lib/buildeventstream/proto/build_event_stream.proto
src/main/protobuf/invocation_policy.proto: File not found.
src/main/protobuf/command_line.proto: File not found.
build_event_stream.proto: Import "src/main/protobuf/invocation_policy.proto" was not found or had errors.
build_event_stream.proto: Import "src/main/protobuf/command_line.proto" was not found or had errors.
build_event_stream.proto:306:3: "blaze.invocation_policy.InvocationPolicy" is not defined.
build_event_stream.proto:573:5: "command_line.CommandLine" is not defined.
```

Edit _scripts/bootstrap/compile.sh_

Issue
```
[vagrant@localhost bazel]$ PROTOC=/usr/local/bin/protoc GRPC_JAVA_PLUGIN=/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java ./compile.sh 
🍃  Building Bazel from scratch../usr/lib/jvm/java-1.8.0-openjdk-1.8.0.91-6.b14.fc23.x86_64/bin/javac -classpath third_party/aether/aether-api-1.0.0.v20140518.jar:third_party/aether/aether-connector-basic-1.0.0.v20140518.jar:third_party/aether/aether-impl-1.0.0.v20140518.jar:third_party/aether/aether-spi-1.0.0.v20140518.jar:third_party/aether/aether-transport-classpath-1.0.0.v20140518.jar:third_party/aether/aether-transport-file-1.0.0.v20140518.jar:third_party/aether/aether-transport-http-1.0.0.v20140518.jar:third_party/aether/aether-transport-wagon-1.0.0.v20140518.jar:third_party/aether/aether-util-1.0.0.v20140518.jar:third_party/allocation_instrumenter/java-allocation-instrumenter-3.0.1.jar:third_party/android_common/com.android.tools.build_builder-model_2.0.0.jar:third_party/android_common/com.android.tools.build_builder-test-api_2.0.0.jar:third_party/android_common/com.android.tools.build_builder_2.0.0.jar:third_party/android_common/com.android.tools.build_manifest-merger_25.0.0.jar:third_party/android_common/com.android.tools.external.lombok_lombok-ast_0.2.3.jar:third_party/android_common/com.android.tools.layoutlib_layoutlib_25.0.0.jar:third_party/android_common/com.android.tools.lint_lint-api_25.0.0.jar:third_party/android_common/com.android.tools.lint_lint-checks_25.0.0.jar:third_party/android_common/com.android.tools_common_25.0.0.jar:third_party/android_common/com.android.tools_ddmlib_25.0.0.jar:third_party/android_common/com.android.tools_dvlib_25.0.0.jar:third_party/android_common/com.android.tools_repository_25.0.0.jar:third_party/android_common/com.android.tools_sdk-common_25.0.0.jar:third_party/android_common/com.android.tools_sdklib_25.0.0.jar:third_party/android_common/com.android_annotations_25.0.0.jar:third_party/apache_commons_codec/commons-codec-1.9.jar:third_party/apache_commons_collections/commons-collections-3.2.2.jar:third_party/apache_commons_compress/apache-commons-compress-1.9.jar:third_party/apache_commons_lang/commons-lang-2.6.jar:third_party/apache_commons_logging/commons-logging-1.1.1.jar:third_party/apache_commons_pool2/commons-pool2-2.3.jar:third_party/apache_httpclient/httpclient-4.5.3.jar:third_party/apache_httpcore/httpcore-4.4.6.jar:third_party/apache_velocity/velocity-1.7.jar:third_party/api_client/google-api-client-1.22.0.jar:third_party/api_client/google-api-client-jackson2-1.22.0.jar:third_party/api_client/google-http-client-1.22.0.jar:third_party/api_client/google-http-client-jackson2-1.22.0.jar:third_party/asm/asm-5.1.jar:third_party/asm/asm-analysis-5.1.jar:third_party/asm/asm-commons-5.1.jar:third_party/asm/asm-tree-5.1.jar:third_party/asm/asm-util-5.1.jar:third_party/auth/google-auth-library-credentials-0.6.0.jar:third_party/auth/google-auth-library-oauth2-http-0.6.0.jar:third_party/auto/auto-common-0.3.jar:third_party/auto/auto-service-1.0-rc2.jar:third_party/auto/auto-value-1.4.jar:third_party/bytebuddy/byte-buddy-dep-0.7-rc6.jar:third_party/error_prone/error_prone_annotation-2.1.1.jar:third_party/error_prone/error_prone_annotations-2.1.1.jar:third_party/error_prone/error_prone_check_api-2.1.1.jar:third_party/guava/guava-testlib-23.1-jre.jar:third_party/gson/gson-2.2.4.jar:third_party/hamcrest/hamcrest-core-1.3.jar:third_party/hazelcast/hazelcast-3.6.4.jar:third_party/hazelcast/hazelcast-client-3.6.4.jar:third_party/hungarian_algorithm/software-and-algorithms-1.0-src.jar:third_party/hungarian_algorithm/software-and-algorithms-1.0.jar:third_party/ijar/test/libwrongcentraldir.jar:third_party/instrumentation/instrumentation-api-0.4.3.jar:third_party/jackson2/jackson-core-2.8.6.jar:third_party/java/android_databinding/v2_3_1/exec.jar:third_party/java/jacoco/jacocoagent.jar:third_party/java/jacoco/org.jacoco.agent-0.7.5.201505241946-src.jar:third_party/java/jacoco/org.jacoco.agent-0.7.5.201505241946.jar:third_party/java/jacoco/org.jacoco.core-0.7.5.201505241946-src.jar:third_party/java/jacoco/org.jacoco.core-0.7.5.201505241946.jar:third_party/java/jacoco/org.jacoco.report-0.7.5.201505241946-src.jar:third_party/java/jacoco/org.jacoco.report-0.7.5.201505241946.jar:third_party/java/jarjar/gradle/wrapper/gradle-wrapper.jar:third_party/java/jarjar/jarjar-1.4.jar:third_party/java/jarjar/jarjar-core/src/test/resources/enumtest.jar:third_party/java/javapoet/javapoet-1.8.0.jar:third_party/java/jcommander/jcommander-1.48.jar:third_party/java/jdk/langtools/javac-9-dev-r4023-3.jar:third_party/java/jdk/langtools/javac.jar:third_party/java/proguard/proguard5.3.3/examples/annotations/lib/annotations.jar:third_party/java/proguard/proguard5.3.3/lib/proguard.jar:third_party/java/proguard/proguard5.3.3/lib/proguardgui.jar:third_party/java/proguard/proguard5.3.3/lib/retrace.jar:third_party/jcip_annotations/jcip-annotations-1.0-1.jar:third_party/jgit/org.eclipse.jgit-4.0.1.201506240215-r.jar:third_party/jimfs/jimfs-1.1.jar:third_party/joda_time/joda-time-2.3.jar:third_party/jsch/jsch-0.1.51.jar:third_party/jsr305/jsr-305.jar:third_party/jsr330_inject/javax.inject.jar:third_party/junit/junit-4.11.jar:third_party/maven/maven-builder-support-3.3.3.jar:third_party/maven/maven-settings-3.3.3.jar:third_party/maven/maven-settings-builder-3.3.3.jar:third_party/maven_model/maven-aether-provider-3.2.3.jar:third_party/maven_model/maven-model-3.2.3.jar:third_party/maven_model/maven-model-builder-3.2.3.jar:third_party/maven_model/maven-repository-metadata-3.2.3.jar:third_party/mockito/mockito-all-1.10.19.jar:third_party/netty/netty-all-4.1.14.Final.jar:third_party/netty_tcnative/netty-tcnative-boringssl-static-2.0.5.Final.jar:third_party/opencensus/opencensus-api-0.5.1.jar:third_party/pcollections/pcollections-2.1.2.jar:third_party/plexus_component_annotations/plexus-component-annotations-1.6.jar:third_party/plexus_interpolation/plexus-interpolation-1.22.jar:third_party/plexus_utils/plexus-utils-3.0.21.jar:third_party/protobuf/3.4.0/libprotobuf_java.jar:third_party/protobuf/3.4.0/libprotobuf_java_util.jar:third_party/slf4j/slf4j-api-1.7.7.jar:third_party/slf4j/slf4j-jdk14-1.7.7.jar:third_party/tomcat_annotations_api/tomcat-annotations-api-8.0.5.jar:third_party/truth/truth-0.36.jar:third_party/truth8/truth-java8-extension-0.36.jar:third_party/turbine/turbine-0.1-20170828.jar:third_party/xz/xz-1.5.jar:third_party/grpc/grpc-auth-1.6.1.jar:third_party/grpc/grpc-context-1.6.1.jar:third_party/grpc/grpc-core-1.6.1.jar:third_party/grpc/grpc-netty-1.6.1.jar:third_party/grpc/grpc-protobuf-1.6.1.jar:third_party/grpc/grpc-protobuf-lite-1.6.1.jar:third_party/grpc/grpc-stub-1.6.1.jar:third_party/guava/guava-23.1-jre.jar:third_party/error_prone/error_prone_core-2.1.1.jar:/tmp/bazel_5f3TDAii -sourcepath src/java_tools/singlejar/java/com/google/devtools/build/zip:src/main/java:src/tools/xcode-common/java/com/google/devtools/build/xcode/common:src/tools/xcode-common/java/com/google/devtools/build/xcode/util:third_party/java/dd_plist/java:/tmp/bazel_5f3TDAii/src -d /tmp/bazel_5f3TDAii/classes -source 1.8 -target 1.8 -encoding UTF-8 @/tmp/bazel_PwfcGnuo/param
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:17: 警告: Signal是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Signal;
               ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:18: 警告: SignalHandler是内部专用 API, 可能会在未来发行版中删除
import sun.misc.SignalHandler;
               ^
src/main/java/com/google/devtools/build/lib/skyframe/serialization/strings/FastStringCodec.java:27: 警告: Unsafe是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Unsafe;
               ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:17: 警告: Signal是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Signal;
               ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:18: 警告: SignalHandler是内部专用 API, 可能会在未来发行版中删除
import sun.misc.SignalHandler;
               ^
src/main/java/com/google/devtools/build/lib/skyframe/serialization/strings/FastStringCodec.java:27: 警告: Unsafe是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Unsafe;
               ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:32: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Function;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:33: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Line;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:34: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Profile;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:35: 错误: 程序包com.google.perftools.profiles.ProfileProto.Profile不存在
import com.google.perftools.profiles.ProfileProto.Profile.Builder;
                                                         ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:36: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Sample;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:37: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.ValueType;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:237: 错误: 找不到符号
  Profile buildMemoryProfile() {
  ^
  符号:   类 Profile
  位置: 类 AllocationTracker
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:22: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import static com.google.devtools.build.v1.BuildEvent.EventCase.COMPONENT_STREAM_FINISHED;
                                                     ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:22: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildEvent.EventCase.COMPONENT_STREAM_FINISHED;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:23: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import static com.google.devtools.build.v1.BuildStatus.Result.COMMAND_FAILED;
                                                      ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:23: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildStatus.Result.COMMAND_FAILED;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:24: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import static com.google.devtools.build.v1.BuildStatus.Result.COMMAND_SUCCEEDED;
                                                      ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:24: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildStatus.Result.COMMAND_SUCCEEDED;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:25: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import static com.google.devtools.build.v1.BuildStatus.Result.UNKNOWN_STATUS;
                                                      ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:25: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildStatus.Result.UNKNOWN_STATUS;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:53: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import com.google.devtools.build.v1.BuildStatus.Result;
                                               ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:54: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:55: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamResponse;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:56: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishLifecycleEventRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:19: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:20: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamResponse;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:21: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishLifecycleEventRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:17: 错误: 程序包com.google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished不存在
import static com.google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished.FinishType.FINISHED;
                                                                                  ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:17: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished.FinishType.FINISHED;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:22: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.BuildEvent;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:23: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:24: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.BuildEnqueued;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:25: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.BuildFinished;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:26: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.EventCase;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:27: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.InvocationAttemptFinished;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:28: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.InvocationAttemptStarted;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:29: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.BuildStatus;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:30: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import com.google.devtools.build.v1.BuildStatus.Result;
                                               ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:31: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.OrderedBuildEvent;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:32: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:33: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishLifecycleEventRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:34: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.StreamId;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:35: 错误: 程序包com.google.devtools.build.v1.StreamId不存在
import com.google.devtools.build.v1.StreamId.BuildComponent;
                                            ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:100: 错误: 找不到符号
  private ConcurrentLinkedDeque<PublishBuildToolEventStreamRequest> pendingAck;
                                ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:102: 错误: 找不到符号
  private final BlockingDeque<PublishBuildToolEventStreamRequest> pendingSend;
                              ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:104: 错误: 找不到符号
  private Result invocationResult;
          ^
  符号:   类 Result
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:364: 错误: 找不到符号
      PublishBuildToolEventStreamRequest serialisedEvent) {
      ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:376: 错误: 找不到符号
  private synchronized Result getInvocationResult() {
                       ^
  符号:   类 Result
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:444: 错误: 找不到符号
  private Status publishLifecycleEvent(PublishLifecycleEventRequest request) throws Exception {
                                       ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:469: 错误: 找不到符号
      final ConcurrentLinkedDeque<PublishBuildToolEventStreamRequest> pendingAck,
                                  ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:470: 错误: 找不到符号
      final BlockingDeque<PublishBuildToolEventStreamRequest> pendingSend,
                          ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:529: 错误: 找不到符号
  private static boolean isLastEvent(@Nullable PublishBuildToolEventStreamRequest event) {
                                               ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:536: 错误: 找不到符号
      final Deque<PublishBuildToolEventStreamRequest> pendingAck,
                  ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:535: 错误: 找不到符号
  private Function<PublishBuildToolEventStreamResponse, Void> ackCallback(
                   ^
  符号:   类 PublishBuildToolEventStreamResponse
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:33: 错误: 找不到符号
  Status publish(PublishLifecycleEventRequest lifecycleEvent) throws Exception;
                 ^
  符号:   类 PublishLifecycleEventRequest
  位置: 接口 BuildEventServiceClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:45: 错误: 找不到符号
      Function<PublishBuildToolEventStreamResponse, Void> ackCallback) throws Exception;
               ^
  符号:   类 PublishBuildToolEventStreamResponse
  位置: 接口 BuildEventServiceClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:53: 错误: 找不到符号
  void sendOverStream(PublishBuildToolEventStreamRequest buildEvent) throws Exception;
                      ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 接口 BuildEventServiceClient
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:69: 错误: 找不到符号
  public PublishLifecycleEventRequest buildEnqueued() {
         ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:77: 错误: 找不到符号
  public PublishLifecycleEventRequest buildFinished(Result result) {
                                                    ^
  符号:   类 Result
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:77: 错误: 找不到符号
  public PublishLifecycleEventRequest buildFinished(Result result) {
         ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:87: 错误: 找不到符号
  public PublishLifecycleEventRequest invocationStarted() {
         ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:96: 错误: 找不到符号
  public PublishLifecycleEventRequest invocationFinished(Result result) {
                                                         ^
  符号:   类 Result
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:96: 错误: 找不到符号
  public PublishLifecycleEventRequest invocationFinished(Result result) {
         ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:110: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest streamFinished() {
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:117: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest bazelEvent(Any packedEvent) {
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:122: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest bazelEvent(int sequenceNumber, Any packedEvent) {
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:129: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest streamFinished(int sequenceNumber) {
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:139: 错误: 程序包BuildEvent不存在
      int sequenceNumber, BuildEvent.Builder besEvent) {
                                    ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:138: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest publishBuildToolEventStreamRequest(
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:156: 错误: 程序包BuildEvent不存在
      int sequenceNumber, BuildEvent.Builder lifecycleEvent) {
                                    ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:155: 错误: 程序包PublishLifecycleEventRequest不存在
  public PublishLifecycleEventRequest.Builder lifecycleEvent(@Nullable String projectId,
                                     ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:171: 错误: 找不到符号
  public StreamId streamId(EventCase eventCase) {
                           ^
  符号:   类 EventCase
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:171: 错误: 找不到符号
  public StreamId streamId(EventCase eventCase) {
         ^
  符号:   类 StreamId
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:426: 错误: 找不到符号
    private void publishInvocationFinishedEvent(final Result result) throws Exception {
                                                      ^
  符号:   类 Result
  位置: 类 BuildEventServiceTransport.BuildEventServiceUpload
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:434: 错误: 找不到符号
    private void publishBuildFinishedEvent(final Result result) throws Exception {
                                                 ^
  符号:   类 Result
  位置: 类 BuildEventServiceTransport.BuildEventServiceUpload
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:25: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildEventGrpc;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:26: 错误: 程序包com.google.devtools.build.v1.PublishBuildEventGrpc不存在
import com.google.devtools.build.v1.PublishBuildEventGrpc.PublishBuildEventBlockingStub;
                                                         ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:27: 错误: 程序包com.google.devtools.build.v1.PublishBuildEventGrpc不存在
import com.google.devtools.build.v1.PublishBuildEventGrpc.PublishBuildEventStub;
                                                         ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:28: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:29: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamResponse;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:30: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishLifecycleEventRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:47: 错误: 找不到符号
  private final PublishBuildEventStub besAsync;
                ^
  符号:   类 PublishBuildEventStub
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:48: 错误: 找不到符号
  private final PublishBuildEventBlockingStub besBlocking;
                ^
  符号:   类 PublishBuildEventBlockingStub
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:50: 错误: 找不到符号
  private final AtomicReference<StreamObserver<PublishBuildToolEventStreamRequest>> streamReference;
                                               ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:70: 错误: 找不到符号
  public Status publish(PublishLifecycleEventRequest lifecycleEvent) throws Exception {
                        ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:85: 错误: 找不到符号
      Function<PublishBuildToolEventStreamResponse, Void> ack)
               ^
  符号:   类 PublishBuildToolEventStreamResponse
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:95: 错误: 找不到符号
      final Function<PublishBuildToolEventStreamResponse, Void> ack,
                     ^
  符号:   类 PublishBuildToolEventStreamResponse
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:94: 错误: 找不到符号
  private StreamObserver<PublishBuildToolEventStreamRequest> createStream(
                         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:125: 错误: 找不到符号
  public void sendOverStream(PublishBuildToolEventStreamRequest buildEvent) throws Exception {
                             ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:19: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Function;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:20: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Line;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:21: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Location;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:22: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Profile;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:23: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Sample;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:24: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.ValueType;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:148: 错误: 找不到符号
  static Profile createProfile() {
         ^
  符号:   类 Profile
  位置: 类 Callcounts
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:174: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:178: 错误: 程序包Profile不存在
    StringTable(Profile.Builder profile) {
                       ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:194: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:199: 错误: 程序包Profile不存在
    FunctionTable(Profile.Builder profile, StringTable stringTable) {
                         ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:218: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:223: 错误: 程序包Profile不存在
    LocationTable(Profile.Builder profile, FunctionTable functionTable) {
                         ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:293: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:297: 错误: 程序包Profile不存在
    StringTable(Profile.Builder profile) {
                       ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:313: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:17: 警告: Signal是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Signal;
               ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:18: 警告: SignalHandler是内部专用 API, 可能会在未来发行版中删除
import sun.misc.SignalHandler;
               ^
src/main/java/com/google/devtools/build/lib/skyframe/serialization/strings/FastStringCodec.java:27: 警告: Unsafe是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Unsafe;
               ^
100 个错误
9 个警告
```

### 2017-11-11

Project
```
[vagrant@localhost bazel]$ git pull --append
remote: Counting objects: 329, done.
remote: Compressing objects: 100% (86/86), done.
remote: Total 329 (delta 194), reused 303 (delta 183), pack-reused 0
接收对象中: 100% (329/329), 94.36 KiB | 0 bytes/s, 完成.
处理 delta 中: 100% (194/194), 完成 95 个本地对象.
来自 https://github.com/bazelbuild/bazel
   9135b7b..1c639ab  master     -> origin/master
更新 9135b7b..1c639ab
Fast-forward
 site/docs/platforms.md                                                        |  79 ++---
 site/docs/toolchains.md                                                       |   8 +-
 site/docs/user-manual.html                                                    |  14 -
 src/main/java/com/google/devtools/build/docgen/templates/be/be-nav.vm         |   1 +
 src/main/java/com/google/devtools/build/lib/BUILD                             |   2 +-
 src/main/java/com/google/devtools/build/lib/actions/ActionExecutedEvent.java  |  12 +
 .../com/google/devtools/build/lib/analysis/AspectAwareAttributeMapper.java    |  28 +-
 src/main/java/com/google/devtools/build/lib/analysis/Runfiles.java            |  23 +-
 .../java/com/google/devtools/build/lib/analysis/RunfilesSupplierImpl.java     |   3 +-
 src/main/java/com/google/devtools/build/lib/analysis/RunfilesSupport.java     |  10 +-
 .../java/com/google/devtools/build/lib/analysis/SourceManifestAction.java     |   2 +-
 .../devtools/build/lib/analysis/constraints/TopLevelConstraintSemantics.java  |   9 +
 .../com/google/devtools/build/lib/bazel/rules/java/BazelJavaSemantics.java    |  21 +-
 .../google/devtools/build/lib/bazel/rules/python/BazelPythonSemantics.java    |   2 +-
 .../com/google/devtools/build/lib/bazel/rules/workspace/HttpArchiveRule.java  |   4 +-
 .../google/devtools/build/lib/bazel/rules/workspace/NewHttpArchiveRule.java   |   4 +-
 .../google/devtools/build/lib/buildeventstream/proto/build_event_stream.proto |   3 +
 src/main/java/com/google/devtools/build/lib/buildtool/BuildTool.java          |  24 +-
 .../java/com/google/devtools/build/lib/packages/AbstractAttributeMapper.java  |   6 +
 .../com/google/devtools/build/lib/packages/AggregatingAttributeMapper.java    |   6 +
 src/main/java/com/google/devtools/build/lib/packages/AttributeMap.java        |   8 +-
 .../com/google/devtools/build/lib/packages/DelegatingAttributeMapper.java     |   6 +
 .../google/devtools/build/lib/rules/android/AndroidHostServiceFixture.java    |   2 +-
 .../devtools/build/lib/rules/android/LibraryRGeneratorActionBuilder.java      |   2 +-
 .../google/devtools/build/lib/rules/android/ManifestMergerActionBuilder.java  |   2 +-
 .../google/devtools/build/lib/rules/android/RClassGeneratorActionBuilder.java |   2 +-
 src/main/java/com/google/devtools/build/lib/rules/cpp/CppActionConfigs.java   |  30 +-
 src/main/java/com/google/devtools/build/lib/rules/cpp/CppConfiguration.java   |   7 -
 .../java/com/google/devtools/build/lib/rules/cpp/CppLinkActionBuilder.java    |  10 -
 src/main/java/com/google/devtools/build/lib/rules/cpp/CppOptions.java         |  16 -
 .../java/com/google/devtools/build/lib/rules/java/DeployArchiveBuilder.java   |  78 +++-
 src/main/java/com/google/devtools/build/lib/rules/java/JavaBinary.java        |   3 +
 .../java/com/google/devtools/build/lib/rules/java/JavaCompilationHelper.java  |   5 +-
 src/main/java/com/google/devtools/build/lib/rules/java/JavaCompileAction.java |  26 +-
 src/main/java/com/google/devtools/build/lib/rules/java/JavaSemantics.java     |   6 +-
 src/main/java/com/google/devtools/build/lib/rules/java/JavaToolchain.java     |   3 +-
 .../java/com/google/devtools/build/lib/rules/java/JavaToolchainProvider.java  |   6 +-
 src/main/java/com/google/devtools/build/lib/rules/java/JavaToolchainRule.java |   1 -
 .../google/devtools/build/lib/rules/java/OneVersionCheckActionBuilder.java    |   8 +-
 src/main/java/com/google/devtools/build/lib/rules/objc/AppleBinaryRule.java   |  16 +-
 .../java/com/google/devtools/build/lib/rules/objc/AppleStaticLibraryRule.java |  17 +-
 src/main/java/com/google/devtools/build/lib/rules/objc/BUILD                  |   1 +
 src/main/java/com/google/devtools/build/lib/rules/objc/ObjcRuleClasses.java   |  12 +-
 .../java/com/google/devtools/build/lib/runtime/BlazeCommandDispatcher.java    |   6 +-
 src/main/java/com/google/devtools/build/lib/runtime/BlazeOptionHandler.java   | 367 +++++++++++--------
 src/main/java/com/google/devtools/build/lib/runtime/CommonCommandOptions.java |  15 +
 src/test/java/com/google/devtools/build/android/AndroidDataMergerTest.java    |  14 +-
 src/test/java/com/google/devtools/build/android/AndroidDataWriterTest.java    |  11 +-
 .../com/google/devtools/build/android/AndroidResourceClassWriterTest.java     |  11 +-
 src/test/java/com/google/devtools/build/android/ClassPathsSubject.java        |  22 +-
 src/test/java/com/google/devtools/build/android/DataResourceXmlTest.java      |  11 +-
 src/test/java/com/google/devtools/build/android/ParsedAndroidDataSubject.java |   6 +-
 src/test/java/com/google/devtools/build/android/ParsedAndroidDataTest.java    |  13 +-
 src/test/java/com/google/devtools/build/android/PathsSubject.java             |   6 +-
 .../com/google/devtools/build/android/UnwrittenMergedAndroidDataSubject.java  |  18 +-
 .../java/com/google/devtools/build/lib/analysis/RunfilesSupplierImplTest.java |  13 -
 src/test/java/com/google/devtools/build/lib/packages/util/MOCK_OSX_CROSSTOOL  | 360 ++-----------------
 .../devtools/build/lib/rules/android/AndroidHostServiceFixtureTest.java       |   2 +-
 src/test/java/com/google/devtools/build/lib/rules/objc/AppleBinaryTest.java   |  91 +++++
 .../java/com/google/devtools/build/lib/rules/objc/AppleStaticLibraryTest.java |  91 +++++
 src/test/java/com/google/devtools/build/lib/rules/objc/ObjcRuleTestCase.java  |  40 +++
 .../java/com/google/devtools/build/lib/runtime/BlazeOptionHandlerTest.java    | 687 ++++++++++++++++++++++++++++++++++++
 .../java/com/google/devtools/build/lib/testutil/EventIterableSubject.java     |   6 +-
 .../com/google/devtools/build/lib/testutil/EventIterableSubjectFactory.java   |  15 +-
 src/test/java/com/google/devtools/build/skyframe/CycleInfoSubject.java        |   7 +-
 src/test/java/com/google/devtools/build/skyframe/CycleInfoSubjectFactory.java |  12 +-
 src/test/java/com/google/devtools/build/skyframe/ErrorInfoSubject.java        |   6 +-
 src/test/java/com/google/devtools/build/skyframe/ErrorInfoSubjectFactory.java |  12 +-
 src/test/java/com/google/devtools/build/skyframe/EvaluationResultSubject.java |   6 +-
 .../com/google/devtools/build/skyframe/EvaluationResultSubjectFactory.java    |  16 +-
 src/test/java/com/google/devtools/build/skyframe/NodeEntrySubject.java        |   6 +-
 src/test/java/com/google/devtools/build/skyframe/NodeEntrySubjectFactory.java |  12 +-
 src/test/shell/integration/build_event_stream_test.sh                         |   3 +-
 src/test/shell/integration/discard_graph_edges_test.sh                        |   8 +
 tools/build_defs/repo/java.bzl                                                |   1 -
 tools/osx/crosstool/CROSSTOOL.tpl                                             | 300 ++--------------
 76 files changed, 1590 insertions(+), 1132 deletions(-)
 create mode 100644 src/test/java/com/google/devtools/build/lib/runtime/BlazeOptionHandlerTest.java
```

Modified
```
[vagrant@localhost bazel]$ git status
位于分支 master
您的分支与上游分支 'origin/master' 一致。
尚未暂存以备提交的变更：
  （使用 "git add <文件>..." 更新要提交的内容）
  （使用 "git checkout -- <文件>..." 丢弃工作区的改动）

	修改：     scripts/bootstrap/compile.sh


耗费了 2.20 秒以枚举未跟踪的文件。'status -uno' 也许能提高速度，
但您需要小心不要忘了添加新文件（参见 'git help status'）。
修改尚未加入提交（使用 "git add" 和/或 "git commit -a"）
[vagrant@localhost bazel]$ git diff scripts/bootstrap/compile.sh
diff --git a/scripts/bootstrap/compile.sh b/scripts/bootstrap/compile.sh
index a6a411f..d61c4b0 100755
--- a/scripts/bootstrap/compile.sh
+++ b/scripts/bootstrap/compile.sh
@@ -189,7 +189,7 @@ if [ -z "${BAZEL_SKIP_JAVA_COMPILATION}" ]; then
 
         log "Compiling Java stubs for protocol buffers..."
         for f in $PROTO_FILES ; do
-            run "${PROTOC}" -Isrc/main/protobuf/ \
+            run "${PROTOC}" -I. -Isrc/main/protobuf/ \
                 -Isrc/main/java/com/google/devtools/build/lib/buildeventstream/proto/ \
                 --java_out=${OUTPUT_DIR}/src \
                 --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
```

Protobuf
```
[vagrant@localhost bazel]$ export PROTOC=/usr/local/bin/protoc
```

gRPC
```
[vagrant@localhost bazel]$ export GRPC_JAVA_PLUGIN=/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java
```

Issue
```
[vagrant@localhost bazel]$ ./compile.sh 
🍃  Building Bazel from scratch../usr/lib/jvm/java-1.8.0-openjdk-1.8.0.91-6.b14.fc23.x86_64/bin/javac -classpath third_party/aether/aether-api-1.0.0.v20140518.jar:third_party/aether/aether-connector-basic-1.0.0.v20140518.jar:third_party/aether/aether-impl-1.0.0.v20140518.jar:third_party/aether/aether-spi-1.0.0.v20140518.jar:third_party/aether/aether-transport-classpath-1.0.0.v20140518.jar:third_party/aether/aether-transport-file-1.0.0.v20140518.jar:third_party/aether/aether-transport-http-1.0.0.v20140518.jar:third_party/aether/aether-transport-wagon-1.0.0.v20140518.jar:third_party/aether/aether-util-1.0.0.v20140518.jar:third_party/allocation_instrumenter/java-allocation-instrumenter-3.0.1.jar:third_party/android_common/com.android.tools.build_builder-model_2.0.0.jar:third_party/android_common/com.android.tools.build_builder-test-api_2.0.0.jar:third_party/android_common/com.android.tools.build_builder_2.0.0.jar:third_party/android_common/com.android.tools.build_manifest-merger_25.0.0.jar:third_party/android_common/com.android.tools.external.lombok_lombok-ast_0.2.3.jar:third_party/android_common/com.android.tools.layoutlib_layoutlib_25.0.0.jar:third_party/android_common/com.android.tools.lint_lint-api_25.0.0.jar:third_party/android_common/com.android.tools.lint_lint-checks_25.0.0.jar:third_party/android_common/com.android.tools_common_25.0.0.jar:third_party/android_common/com.android.tools_ddmlib_25.0.0.jar:third_party/android_common/com.android.tools_dvlib_25.0.0.jar:third_party/android_common/com.android.tools_repository_25.0.0.jar:third_party/android_common/com.android.tools_sdk-common_25.0.0.jar:third_party/android_common/com.android.tools_sdklib_25.0.0.jar:third_party/android_common/com.android_annotations_25.0.0.jar:third_party/apache_commons_codec/commons-codec-1.9.jar:third_party/apache_commons_collections/commons-collections-3.2.2.jar:third_party/apache_commons_compress/apache-commons-compress-1.9.jar:third_party/apache_commons_lang/commons-lang-2.6.jar:third_party/apache_commons_logging/commons-logging-1.1.1.jar:third_party/apache_commons_pool2/commons-pool2-2.3.jar:third_party/apache_httpclient/httpclient-4.5.3.jar:third_party/apache_httpcore/httpcore-4.4.6.jar:third_party/apache_velocity/velocity-1.7.jar:third_party/api_client/google-api-client-1.22.0.jar:third_party/api_client/google-api-client-jackson2-1.22.0.jar:third_party/api_client/google-http-client-1.22.0.jar:third_party/api_client/google-http-client-jackson2-1.22.0.jar:third_party/asm/asm-5.1.jar:third_party/asm/asm-analysis-5.1.jar:third_party/asm/asm-commons-5.1.jar:third_party/asm/asm-tree-5.1.jar:third_party/asm/asm-util-5.1.jar:third_party/auth/google-auth-library-credentials-0.6.0.jar:third_party/auth/google-auth-library-oauth2-http-0.6.0.jar:third_party/auto/auto-common-0.3.jar:third_party/auto/auto-service-1.0-rc2.jar:third_party/auto/auto-value-1.4.jar:third_party/bytebuddy/byte-buddy-dep-0.7-rc6.jar:third_party/error_prone/error_prone_annotation-2.1.1.jar:third_party/error_prone/error_prone_annotations-2.1.1.jar:third_party/error_prone/error_prone_check_api-2.1.1.jar:third_party/guava/guava-testlib-23.1-jre.jar:third_party/gson/gson-2.2.4.jar:third_party/hamcrest/hamcrest-core-1.3.jar:third_party/hazelcast/hazelcast-3.6.4.jar:third_party/hazelcast/hazelcast-client-3.6.4.jar:third_party/hungarian_algorithm/software-and-algorithms-1.0-src.jar:third_party/hungarian_algorithm/software-and-algorithms-1.0.jar:third_party/ijar/test/libwrongcentraldir.jar:third_party/instrumentation/instrumentation-api-0.4.3.jar:third_party/jackson2/jackson-core-2.8.6.jar:third_party/java/android_databinding/v2_3_1/exec.jar:third_party/java/jacoco/jacocoagent.jar:third_party/java/jacoco/org.jacoco.agent-0.7.5.201505241946-src.jar:third_party/java/jacoco/org.jacoco.agent-0.7.5.201505241946.jar:third_party/java/jacoco/org.jacoco.core-0.7.5.201505241946-src.jar:third_party/java/jacoco/org.jacoco.core-0.7.5.201505241946.jar:third_party/java/jacoco/org.jacoco.report-0.7.5.201505241946-src.jar:third_party/java/jacoco/org.jacoco.report-0.7.5.201505241946.jar:third_party/java/jarjar/gradle/wrapper/gradle-wrapper.jar:third_party/java/jarjar/jarjar-1.4.jar:third_party/java/jarjar/jarjar-core/src/test/resources/enumtest.jar:third_party/java/javapoet/javapoet-1.8.0.jar:third_party/java/jcommander/jcommander-1.48.jar:third_party/java/jdk/langtools/javac-9-dev-r4023-3.jar:third_party/java/jdk/langtools/javac.jar:third_party/java/proguard/proguard5.3.3/examples/annotations/lib/annotations.jar:third_party/java/proguard/proguard5.3.3/lib/proguard.jar:third_party/java/proguard/proguard5.3.3/lib/proguardgui.jar:third_party/java/proguard/proguard5.3.3/lib/retrace.jar:third_party/jcip_annotations/jcip-annotations-1.0-1.jar:third_party/jgit/org.eclipse.jgit-4.0.1.201506240215-r.jar:third_party/jimfs/jimfs-1.1.jar:third_party/joda_time/joda-time-2.3.jar:third_party/jsch/jsch-0.1.51.jar:third_party/jsr305/jsr-305.jar:third_party/jsr330_inject/javax.inject.jar:third_party/junit/junit-4.11.jar:third_party/maven/maven-builder-support-3.3.3.jar:third_party/maven/maven-settings-3.3.3.jar:third_party/maven/maven-settings-builder-3.3.3.jar:third_party/maven_model/maven-aether-provider-3.2.3.jar:third_party/maven_model/maven-model-3.2.3.jar:third_party/maven_model/maven-model-builder-3.2.3.jar:third_party/maven_model/maven-repository-metadata-3.2.3.jar:third_party/mockito/mockito-all-1.10.19.jar:third_party/netty/netty-all-4.1.14.Final.jar:third_party/netty_tcnative/netty-tcnative-boringssl-static-2.0.5.Final.jar:third_party/opencensus/opencensus-api-0.5.1.jar:third_party/pcollections/pcollections-2.1.2.jar:third_party/plexus_component_annotations/plexus-component-annotations-1.6.jar:third_party/plexus_interpolation/plexus-interpolation-1.22.jar:third_party/plexus_utils/plexus-utils-3.0.21.jar:third_party/protobuf/3.4.0/libprotobuf_java.jar:third_party/protobuf/3.4.0/libprotobuf_java_util.jar:third_party/slf4j/slf4j-api-1.7.7.jar:third_party/slf4j/slf4j-jdk14-1.7.7.jar:third_party/tomcat_annotations_api/tomcat-annotations-api-8.0.5.jar:third_party/truth/truth-0.36.jar:third_party/truth8/truth-java8-extension-0.36.jar:third_party/turbine/turbine-0.1-20170828.jar:third_party/xz/xz-1.5.jar:third_party/grpc/grpc-auth-1.6.1.jar:third_party/grpc/grpc-context-1.6.1.jar:third_party/grpc/grpc-core-1.6.1.jar:third_party/grpc/grpc-netty-1.6.1.jar:third_party/grpc/grpc-protobuf-1.6.1.jar:third_party/grpc/grpc-protobuf-lite-1.6.1.jar:third_party/grpc/grpc-stub-1.6.1.jar:third_party/guava/guava-23.1-jre.jar:third_party/error_prone/error_prone_core-2.1.1.jar:/tmp/bazel_yTlxxH9c -sourcepath src/java_tools/singlejar/java/com/google/devtools/build/zip:src/main/java:src/tools/xcode-common/java/com/google/devtools/build/xcode/common:src/tools/xcode-common/java/com/google/devtools/build/xcode/util:third_party/java/dd_plist/java:/tmp/bazel_yTlxxH9c/src -d /tmp/bazel_yTlxxH9c/classes -source 1.8 -target 1.8 -encoding UTF-8 @/tmp/bazel_h3TKh3LV/param
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:17: 警告: Signal是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Signal;
               ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:18: 警告: SignalHandler是内部专用 API, 可能会在未来发行版中删除
import sun.misc.SignalHandler;
               ^
src/main/java/com/google/devtools/build/lib/skyframe/serialization/strings/FastStringCodec.java:27: 警告: Unsafe是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Unsafe;
               ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:17: 警告: Signal是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Signal;
               ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:18: 警告: SignalHandler是内部专用 API, 可能会在未来发行版中删除
import sun.misc.SignalHandler;
               ^
src/main/java/com/google/devtools/build/lib/skyframe/serialization/strings/FastStringCodec.java:27: 警告: Unsafe是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Unsafe;
               ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:32: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Function;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:33: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Line;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:34: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Profile;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:35: 错误: 程序包com.google.perftools.profiles.ProfileProto.Profile不存在
import com.google.perftools.profiles.ProfileProto.Profile.Builder;
                                                         ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:36: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Sample;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:37: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.ValueType;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:237: 错误: 找不到符号
  Profile buildMemoryProfile() {
  ^
  符号:   类 Profile
  位置: 类 AllocationTracker
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:22: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import static com.google.devtools.build.v1.BuildEvent.EventCase.COMPONENT_STREAM_FINISHED;
                                                     ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:22: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildEvent.EventCase.COMPONENT_STREAM_FINISHED;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:23: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import static com.google.devtools.build.v1.BuildStatus.Result.COMMAND_FAILED;
                                                      ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:23: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildStatus.Result.COMMAND_FAILED;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:24: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import static com.google.devtools.build.v1.BuildStatus.Result.COMMAND_SUCCEEDED;
                                                      ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:24: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildStatus.Result.COMMAND_SUCCEEDED;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:25: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import static com.google.devtools.build.v1.BuildStatus.Result.UNKNOWN_STATUS;
                                                      ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:25: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildStatus.Result.UNKNOWN_STATUS;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:53: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import com.google.devtools.build.v1.BuildStatus.Result;
                                               ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:54: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:55: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamResponse;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:56: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishLifecycleEventRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:19: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:20: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamResponse;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:21: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishLifecycleEventRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:17: 错误: 程序包com.google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished不存在
import static com.google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished.FinishType.FINISHED;
                                                                                  ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:17: 错误: 仅从类和接口静态导入
import static com.google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished.FinishType.FINISHED;
^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:22: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.BuildEvent;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:23: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:24: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.BuildEnqueued;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:25: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.BuildFinished;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:26: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.EventCase;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:27: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.InvocationAttemptFinished;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:28: 错误: 程序包com.google.devtools.build.v1.BuildEvent不存在
import com.google.devtools.build.v1.BuildEvent.InvocationAttemptStarted;
                                              ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:29: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.BuildStatus;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:30: 错误: 程序包com.google.devtools.build.v1.BuildStatus不存在
import com.google.devtools.build.v1.BuildStatus.Result;
                                               ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:31: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.OrderedBuildEvent;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:32: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:33: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishLifecycleEventRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:34: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.StreamId;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:35: 错误: 程序包com.google.devtools.build.v1.StreamId不存在
import com.google.devtools.build.v1.StreamId.BuildComponent;
                                            ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:100: 错误: 找不到符号
  private ConcurrentLinkedDeque<PublishBuildToolEventStreamRequest> pendingAck;
                                ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:102: 错误: 找不到符号
  private final BlockingDeque<PublishBuildToolEventStreamRequest> pendingSend;
                              ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:104: 错误: 找不到符号
  private Result invocationResult;
          ^
  符号:   类 Result
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:364: 错误: 找不到符号
      PublishBuildToolEventStreamRequest serialisedEvent) {
      ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:376: 错误: 找不到符号
  private synchronized Result getInvocationResult() {
                       ^
  符号:   类 Result
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:444: 错误: 找不到符号
  private Status publishLifecycleEvent(PublishLifecycleEventRequest request) throws Exception {
                                       ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:469: 错误: 找不到符号
      final ConcurrentLinkedDeque<PublishBuildToolEventStreamRequest> pendingAck,
                                  ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:470: 错误: 找不到符号
      final BlockingDeque<PublishBuildToolEventStreamRequest> pendingSend,
                          ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:529: 错误: 找不到符号
  private static boolean isLastEvent(@Nullable PublishBuildToolEventStreamRequest event) {
                                               ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:536: 错误: 找不到符号
      final Deque<PublishBuildToolEventStreamRequest> pendingAck,
                  ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:535: 错误: 找不到符号
  private Function<PublishBuildToolEventStreamResponse, Void> ackCallback(
                   ^
  符号:   类 PublishBuildToolEventStreamResponse
  位置: 类 BuildEventServiceTransport
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:33: 错误: 找不到符号
  Status publish(PublishLifecycleEventRequest lifecycleEvent) throws Exception;
                 ^
  符号:   类 PublishLifecycleEventRequest
  位置: 接口 BuildEventServiceClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:45: 错误: 找不到符号
      Function<PublishBuildToolEventStreamResponse, Void> ackCallback) throws Exception;
               ^
  符号:   类 PublishBuildToolEventStreamResponse
  位置: 接口 BuildEventServiceClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceClient.java:53: 错误: 找不到符号
  void sendOverStream(PublishBuildToolEventStreamRequest buildEvent) throws Exception;
                      ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 接口 BuildEventServiceClient
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:69: 错误: 找不到符号
  public PublishLifecycleEventRequest buildEnqueued() {
         ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:77: 错误: 找不到符号
  public PublishLifecycleEventRequest buildFinished(Result result) {
                                                    ^
  符号:   类 Result
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:77: 错误: 找不到符号
  public PublishLifecycleEventRequest buildFinished(Result result) {
         ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:87: 错误: 找不到符号
  public PublishLifecycleEventRequest invocationStarted() {
         ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:96: 错误: 找不到符号
  public PublishLifecycleEventRequest invocationFinished(Result result) {
                                                         ^
  符号:   类 Result
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:96: 错误: 找不到符号
  public PublishLifecycleEventRequest invocationFinished(Result result) {
         ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:110: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest streamFinished() {
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:117: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest bazelEvent(Any packedEvent) {
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:122: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest bazelEvent(int sequenceNumber, Any packedEvent) {
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:129: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest streamFinished(int sequenceNumber) {
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:139: 错误: 程序包BuildEvent不存在
      int sequenceNumber, BuildEvent.Builder besEvent) {
                                    ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:138: 错误: 找不到符号
  public PublishBuildToolEventStreamRequest publishBuildToolEventStreamRequest(
         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:156: 错误: 程序包BuildEvent不存在
      int sequenceNumber, BuildEvent.Builder lifecycleEvent) {
                                    ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:155: 错误: 程序包PublishLifecycleEventRequest不存在
  public PublishLifecycleEventRequest.Builder lifecycleEvent(@Nullable String projectId,
                                     ^
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:171: 错误: 找不到符号
  public StreamId streamId(EventCase eventCase) {
                           ^
  符号:   类 EventCase
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceProtoUtil.java:171: 错误: 找不到符号
  public StreamId streamId(EventCase eventCase) {
         ^
  符号:   类 StreamId
  位置: 类 BuildEventServiceProtoUtil
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:426: 错误: 找不到符号
    private void publishInvocationFinishedEvent(final Result result) throws Exception {
                                                      ^
  符号:   类 Result
  位置: 类 BuildEventServiceTransport.BuildEventServiceUpload
src/main/java/com/google/devtools/build/lib/buildeventservice/BuildEventServiceTransport.java:434: 错误: 找不到符号
    private void publishBuildFinishedEvent(final Result result) throws Exception {
                                                 ^
  符号:   类 Result
  位置: 类 BuildEventServiceTransport.BuildEventServiceUpload
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:25: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildEventGrpc;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:26: 错误: 程序包com.google.devtools.build.v1.PublishBuildEventGrpc不存在
import com.google.devtools.build.v1.PublishBuildEventGrpc.PublishBuildEventBlockingStub;
                                                         ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:27: 错误: 程序包com.google.devtools.build.v1.PublishBuildEventGrpc不存在
import com.google.devtools.build.v1.PublishBuildEventGrpc.PublishBuildEventStub;
                                                         ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:28: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:29: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishBuildToolEventStreamResponse;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:30: 错误: 程序包com.google.devtools.build.v1不存在
import com.google.devtools.build.v1.PublishLifecycleEventRequest;
                                   ^
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:47: 错误: 找不到符号
  private final PublishBuildEventStub besAsync;
                ^
  符号:   类 PublishBuildEventStub
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:48: 错误: 找不到符号
  private final PublishBuildEventBlockingStub besBlocking;
                ^
  符号:   类 PublishBuildEventBlockingStub
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:50: 错误: 找不到符号
  private final AtomicReference<StreamObserver<PublishBuildToolEventStreamRequest>> streamReference;
                                               ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:70: 错误: 找不到符号
  public Status publish(PublishLifecycleEventRequest lifecycleEvent) throws Exception {
                        ^
  符号:   类 PublishLifecycleEventRequest
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:85: 错误: 找不到符号
      Function<PublishBuildToolEventStreamResponse, Void> ack)
               ^
  符号:   类 PublishBuildToolEventStreamResponse
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:95: 错误: 找不到符号
      final Function<PublishBuildToolEventStreamResponse, Void> ack,
                     ^
  符号:   类 PublishBuildToolEventStreamResponse
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:94: 错误: 找不到符号
  private StreamObserver<PublishBuildToolEventStreamRequest> createStream(
                         ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/buildeventservice/client/BuildEventServiceGrpcClient.java:125: 错误: 找不到符号
  public void sendOverStream(PublishBuildToolEventStreamRequest buildEvent) throws Exception {
                             ^
  符号:   类 PublishBuildToolEventStreamRequest
  位置: 类 BuildEventServiceGrpcClient
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:19: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Function;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:20: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Line;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:21: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Location;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:22: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Profile;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:23: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.Sample;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:24: 错误: 程序包com.google.perftools.profiles.ProfileProto不存在
import com.google.perftools.profiles.ProfileProto.ValueType;
                                                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:148: 错误: 找不到符号
  static Profile createProfile() {
         ^
  符号:   类 Profile
  位置: 类 Callcounts
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:174: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:178: 错误: 程序包Profile不存在
    StringTable(Profile.Builder profile) {
                       ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:194: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:199: 错误: 程序包Profile不存在
    FunctionTable(Profile.Builder profile, StringTable stringTable) {
                         ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:218: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/profiler/callcounts/Callcounts.java:223: 错误: 程序包Profile不存在
    LocationTable(Profile.Builder profile, FunctionTable functionTable) {
                         ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:293: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:297: 错误: 程序包Profile不存在
    StringTable(Profile.Builder profile) {
                       ^
src/main/java/com/google/devtools/build/lib/profiler/memory/AllocationTracker.java:313: 错误: 程序包Profile不存在
    final Profile.Builder profile;
                 ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:17: 警告: Signal是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Signal;
               ^
src/main/java/com/google/devtools/build/lib/server/signal/InterruptSignalHandler.java:18: 警告: SignalHandler是内部专用 API, 可能会在未来发行版中删除
import sun.misc.SignalHandler;
               ^
src/main/java/com/google/devtools/build/lib/skyframe/serialization/strings/FastStringCodec.java:27: 警告: Unsafe是内部专用 API, 可能会在未来发行版中删除
import sun.misc.Unsafe;
               ^
100 个错误
9 个警告
```

Googleapis
```
[vagrant@localhost bazel]$ git diff ./scripts/bootstrap/compile.sh
diff --git a/scripts/bootstrap/compile.sh b/scripts/bootstrap/compile.sh
index a6a411f..2a74200 100755
--- a/scripts/bootstrap/compile.sh
+++ b/scripts/bootstrap/compile.sh
@@ -189,12 +189,86 @@ if [ -z "${BAZEL_SKIP_JAVA_COMPILATION}" ]; then
 
         log "Compiling Java stubs for protocol buffers..."
         for f in $PROTO_FILES ; do
-            run "${PROTOC}" -Isrc/main/protobuf/ \
+            run "${PROTOC}" -I. -Isrc/main/protobuf/ \
                 -Isrc/main/java/com/google/devtools/build/lib/buildeventstream/proto/ \
                 --java_out=${OUTPUT_DIR}/src \
                 --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
                 --grpc_out=${OUTPUT_DIR}/src "$f"
         done
+        run ${PROTOC} -I. -Ithird_party/pprof/ --java_out=${OUTPUT_DIR}/src third_party/pprof/profile.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/devtools/build/v1/build_events.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/devtools/build/v1/build_status.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/devtools/build/v1/publish_build_event.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/bytestream/bytestream.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/devtools/remoteexecution/v1test/remote_execution.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/rpc/code.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/rpc/error_details.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/rpc/status.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/longrunning/operations.proto
+        run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+            third_party/googleapis/google/watcher/v1/watch.proto
+        for i in $(ls third_party/googleapis/google/api/*.proto) ; do
+            run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+                $i
+        done
+        for i in $(ls third_party/googleapis/google/api/experimental/*.proto) ; do
+            run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+                $i
+        done
+        for i in $(ls third_party/googleapis/google/api/servicecontrol/v1/*.proto) ; do
+            run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+                $i
+        done
+        for i in $(ls third_party/googleapis/google/api/servicemanagement/v1/*.proto) ; do
+            run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+                $i
+        done
+        for i in $(ls third_party/googleapis/google/logging/type/*.proto) ; do
+            run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+                $i
+        done
+        for i in $(ls third_party/googleapis/google/logging/v2/*.proto) ; do
+            run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+                $i
+        done
+        for i in $(ls third_party/googleapis/google/type/*.proto) ; do
+            run ${PROTOC} -I. -Ithird_party/googleapis/ --java_out=${OUTPUT_DIR}/src \
+                --plugin=protoc-gen-grpc="${GRPC_JAVA_PLUGIN-}" \
+                --grpc_out=${OUTPUT_DIR}/src \
+                $i
+        done
+        
     fi
 
   java_compilation "Bazel Java" "$DIRS" "$EXCLUDE_FILES" "$LIBRARY_JARS" "${OUTPUT_DIR}"
```

Issue
```
[vagrant@localhost bazel]$ ./compile.sh 
🍃  Building Bazel from scratch......
🍃  Building Bazel with Bazel.
WARNING: /tmp/bazel_R7IThW63/out/external/bazel_tools/WORKSPACE:1: Workspace name in /tmp/bazel_R7IThW63/out/external/bazel_tools/WORKSPACE (@io_bazel) does not match the name given in the repository's definition (@bazel_tools); this will cause a build error in future versions
ERROR: /Users/fanhongling/Downloads/workspace/src/github.com/bazelbuild/bazel/src/java_tools/buildjar/java/com/google/devtools/build/buildjar/BUILD:144:12: in srcs attribute of bootstrap_java_library rule //src/java_tools/buildjar/java/com/google/devtools/build/buildjar:skylark-deps: '//:bootstrap-derived-java-srcs' does not produce any bootstrap_java_library srcs files (expected .java)
ERROR: Analysis of target '//src:bazel' failed; build aborted: Analysis of target '//src/java_tools/buildjar/java/com/google/devtools/build/buildjar:skylark-deps' failed; build aborted
INFO: Elapsed time: 31.913s
FAILED: Build did NOT complete successfully (107 packages loaded)

ERROR: Could not build Bazel
```

### Install

DNF
```
[vagrant@localhost bazel]$ curl -jkSL https://copr.fedorainfracloud.org/coprs/vbatts/bazel/repo/fedora-23/vbatts-bazel-fedora-23.repo
[vbatts-bazel]
name=Copr repo for bazel owned by vbatts
baseurl=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/fedora-$releasever-$basearch/
type=rpm-md
skip_if_unavailable=True
gpgcheck=1
gpgkey=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/pubkey.gpg
repo_gpgcheck=0
enabled=1
enabled_metadata=1
[vagrant@localhost bazel]$ curl -jkSL https://copr.fedorainfracloud.org/coprs/vbatts/bazel/repo/fedora-24/vbatts-bazel-fedora-24.repo
[vbatts-bazel]
name=Copr repo for bazel owned by vbatts
baseurl=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/fedora-$releasever-$basearch/
type=rpm-md
skip_if_unavailable=True
gpgcheck=1
gpgkey=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/pubkey.gpg
repo_gpgcheck=0
enabled=1
enabled_metadata=1
[vagrant@localhost bazel]$ curl -jkSL https://copr.fedorainfracloud.org/coprs/vbatts/bazel/repo/fedora-25/vbatts-bazel-fedora-25.repo
[vbatts-bazel]
name=Copr repo for bazel owned by vbatts
baseurl=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/fedora-$releasever-$basearch/
type=rpm-md
skip_if_unavailable=True
gpgcheck=1
gpgkey=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/pubkey.gpg
repo_gpgcheck=0
enabled=1
enabled_metadata=1
[vagrant@localhost bazel]$ curl -jkSL https://copr.fedorainfracloud.org/coprs/vbatts/bazel/repo/epel-7/v.repo-bazel-epel-7 
[vbatts-bazel]
name=Copr repo for bazel owned by vbatts
baseurl=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/epel-7-$basearch/
type=rpm-md
skip_if_unavailable=True
gpgcheck=1
gpgkey=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/pubkey.gpg
repo_gpgcheck=0
enabled=1
enabled_metadata=1
```

```
[vagrant@localhost bazel]$ sudo curl -jkSL https://copr.fedorainfracloud.org/coprs/vbatts/bazel/repo/epel-7/vbatts-bazel-epel-7.repo -o /etc/yum.repos.d/vbatts-bazel-epel-7.repo
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   314  100   314    0     0     89      0  0:00:03  0:00:03 --:--:--    89
[vagrant@localhost bazel]$ sudo dnf list | grep bazel
bazel.src                                0.7.0-1.el7.centos              vbatts-bazel
bazel.x86_64                             0.7.0-1.el7.centos              vbatts-bazel
```

```
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
```

```
[vagrant@localhost bazel]$ bazel version
Extracting Bazel installation...
.................
Build label: 0.7.0- (@non-git)
Build target: bazel-out/local-opt/bin/src/main/java/com/google/devtools/build/lib/bazel/BazelServer_deploy.jar
Build time: Wed Oct 18 15:28:47 2017 (1508340527)
Build timestamp: 1508340527
Build timestamp as int: 1508340527
```

```
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

Copr (Fedora-24/25/26/27/..., lack Fedora-23)
```
[vagrant@localhost bazel]$ sudo dnf copr enable vbatts/bazel

You are about to enable a Copr repository. Please note that this
repository is not part of the main Fedora distribution, and quality
may vary.

The Fedora Project does not exercise any power over the contents of
this repository beyond the rules outlined in the Copr FAQ at
<https://fedorahosted.org/copr/wiki/UserDocs#WhatIcanbuildinCopr>, and
packages are not held to any quality or security level.

Please do not file bug reports about these packages in Fedora
Bugzilla. In case of problems, contact the owner of this repository.

Do you want to continue? [y/N]: y
启用软件仓库成功。
[vagrant@localhost bazel]$ cat /etc/yum.repos.d/_copr_vbatts-bazel.repo 
[vbatts-bazel]
name=Copr repo for bazel owned by vbatts
baseurl=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/fedora-$releasever-$basearch/
type=rpm-md
skip_if_unavailable=True
gpgcheck=1
gpgkey=https://copr-be.cloud.fedoraproject.org/results/vbatts/bazel/pubkey.gpg
repo_gpgcheck=0
enabled=1
enabled_metadata=1
[vagrant@localhost bazel]$ sudo dnf copr remove vbatts/bazel
```


### gperftools

```
[vagrant@localhost github.com]$ git clone --depth=1 https://github.com/gperftools/gperftools gperftools/gperftools
正克隆到 'gperftools/gperftools'...
remote: Counting objects: 364, done.
remote: Compressing objects: 100% (338/338), done.
remote: Total 364 (delta 81), reused 115 (delta 11), pack-reused 0
接收对象中: 100% (364/364), 1.04 MiB | 47.00 KiB/s, 完成.
处理 delta 中: 100% (81/81), 完成.
检查连接... 完成。
[vagrant@localhost github.com]$ cd gperftools/gperftools/
[vagrant@localhost gperftools]$ ./autogen.sh 
libtoolize: putting auxiliary files in '.'.
libtoolize: copying file './ltmain.sh'
libtoolize: putting macros in AC_CONFIG_MACRO_DIRS, 'm4'.
libtoolize: copying file 'm4/libtool.m4'
libtoolize: copying file 'm4/ltoptions.m4'
libtoolize: copying file 'm4/ltsugar.m4'
libtoolize: copying file 'm4/ltversion.m4'
libtoolize: copying file 'm4/lt~obsolete.m4'
configure.ac:57: installing './compile'
configure.ac:20: installing './config.guess'
configure.ac:20: installing './config.sub'
configure.ac:21: installing './install-sh'
configure.ac:21: installing './missing'
Makefile.am: installing './depcomp'
parallel-tests: installing './test-driver'
[vagrant@localhost gperftools]$ ./configure
checking build system type... x86_64-unknown-linux-gnu
checking host system type... x86_64-unknown-linux-gnu
checking for a BSD-compatible install... /bin/install -c
checking whether build environment is sane... yes
checking for a thread-safe mkdir -p... /bin/mkdir -p
checking for gawk... gawk
checking whether make sets $(MAKE)... yes
checking whether make supports nested variables... yes
checking whether to enable maintainer-specific portions of Makefiles... no
checking for git... /bin/git
checking for style of include used by make... GNU
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether gcc understands -c and -o together... yes
checking dependency style of gcc... gcc3
checking for g++... g++
checking whether we are using the GNU C++ compiler... yes
checking whether g++ accepts -g... yes
checking dependency style of g++... gcc3
checking for gcc... (cached) gcc
checking whether we are using the GNU C compiler... (cached) yes
checking whether gcc accepts -g... (cached) yes
checking for gcc option to accept ISO C89... (cached) none needed
checking whether gcc understands -c and -o together... (cached) yes
checking dependency style of gcc... (cached) gcc3
checking how to run the C preprocessor... gcc -E
checking for objcopy... objcopy
checking if objcopy supports -W... no
checking how to print strings... printf
checking for a sed that does not truncate output... /bin/sed
checking for grep that handles long lines and -e... /bin/grep
checking for egrep... /bin/grep -E
checking for fgrep... /bin/grep -F
checking for ld used by gcc... /bin/ld
checking if the linker (/bin/ld) is GNU ld... yes
checking for BSD- or MS-compatible name lister (nm)... /bin/nm -B
checking the name lister (/bin/nm -B) interface... BSD nm
checking whether ln -s works... yes
checking the maximum length of command line arguments... 1572864
checking how to convert x86_64-unknown-linux-gnu file names to x86_64-unknown-linux-gnu format... func_convert_file_noop
checking how to convert x86_64-unknown-linux-gnu file names to toolchain format... func_convert_file_noop
checking for /bin/ld option to reload object files... -r
checking for objdump... objdump
checking how to recognize dependent libraries... pass_all
checking for dlltool... no
checking how to associate runtime and link libraries... printf %s\n
checking for ar... ar
checking for archiver @FILE support... @
checking for strip... strip
checking for ranlib... ranlib
checking command to parse /bin/nm -B output from gcc object... ok
checking for sysroot... no
checking for a working dd... /bin/dd
checking how to truncate binary pipes... /bin/dd bs=4096 count=1
checking for mt... no
checking if : is a manifest tool... no
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for dlfcn.h... yes
checking for objdir... .libs
checking if gcc supports -fno-rtti -fno-exceptions... no
checking for gcc option to produce PIC... -fPIC -DPIC
checking if gcc PIC flag -fPIC -DPIC works... yes
checking if gcc static flag -static works... yes
checking if gcc supports -c -o file.o... yes
checking if gcc supports -c -o file.o... (cached) yes
checking whether the gcc linker (/bin/ld -m elf_x86_64) supports shared libraries... yes
checking whether -lc should be explicitly linked in... no
checking dynamic linker characteristics... GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking whether stripping libraries is possible... yes
checking if libtool supports shared libraries... yes
checking whether to build shared libraries... yes
checking whether to build static libraries... yes
checking how to run the C++ preprocessor... g++ -E
checking for ld used by g++... /bin/ld -m elf_x86_64
checking if the linker (/bin/ld -m elf_x86_64) is GNU ld... yes
checking whether the g++ linker (/bin/ld -m elf_x86_64) supports shared libraries... yes
checking for g++ option to produce PIC... -fPIC -DPIC
checking if g++ PIC flag -fPIC -DPIC works... yes
checking if g++ static flag -static works... no
checking if g++ supports -c -o file.o... yes
checking if g++ supports -c -o file.o... (cached) yes
checking whether the g++ linker (/bin/ld -m elf_x86_64) supports shared libraries... yes
checking dynamic linker characteristics... (cached) GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking for inline... inline
checking for __attribute__... yes
checking for __attribute__((aligned(N))) on functions... yes
checking for ANSI C header files... (cached) yes
checking for __int64... no
checking for struct mallinfo... yes
checking for Elf32_Versym... yes
checking for sbrk... yes
checking for geteuid... yes
checking for fork... yes
checking features.h usability... yes
checking features.h presence... yes
checking for features.h... yes
checking malloc.h usability... yes
checking malloc.h presence... yes
checking for malloc.h... yes
checking glob.h usability... yes
checking glob.h presence... yes
checking for glob.h... yes
checking execinfo.h usability... yes
checking execinfo.h presence... yes
checking for execinfo.h... yes
checking unwind.h usability... yes
checking unwind.h presence... yes
checking for unwind.h... yes
checking sched.h usability... yes
checking sched.h presence... yes
checking for sched.h... yes
checking conflict-signal.h usability... no
checking conflict-signal.h presence... no
checking for conflict-signal.h... no
checking sys/prctl.h usability... yes
checking sys/prctl.h presence... yes
checking for sys/prctl.h... yes
checking linux/ptrace.h usability... yes
checking linux/ptrace.h presence... yes
checking for linux/ptrace.h... yes
checking sys/syscall.h usability... yes
checking sys/syscall.h presence... yes
checking for sys/syscall.h... yes
checking sys/socket.h usability... yes
checking sys/socket.h presence... yes
checking for sys/socket.h... yes
checking sys/wait.h usability... yes
checking sys/wait.h presence... yes
checking for sys/wait.h... yes
checking poll.h usability... yes
checking poll.h presence... yes
checking for poll.h... yes
checking fcntl.h usability... yes
checking fcntl.h presence... yes
checking for fcntl.h... yes
checking grp.h usability... yes
checking grp.h presence... yes
checking for grp.h... yes
checking pwd.h usability... yes
checking pwd.h presence... yes
checking for pwd.h... yes
checking sys/resource.h usability... yes
checking sys/resource.h presence... yes
checking for sys/resource.h... yes
checking valgrind.h usability... no
checking valgrind.h presence... no
checking for valgrind.h... no
checking sys/cdefs.h usability... yes
checking sys/cdefs.h presence... yes
checking for sys/cdefs.h... yes
checking for features.h... (cached) yes
checking whether cfree is declared... yes
checking whether posix_memalign is declared... yes
checking whether memalign is declared... yes
checking whether valloc is declared... yes
checking whether pvalloc is declared... yes
checking for stdlib.h... (cached) yes
checking for unistd.h... (cached) yes
checking for sys/param.h... yes
checking for getpagesize... yes
checking for working mmap... no
checking if int32_t is the same type as intptr_t... no
checking ucontext.h usability... yes
checking ucontext.h presence... yes
checking for ucontext.h... yes
checking sys/ucontext.h usability... yes
checking sys/ucontext.h presence... yes
checking for sys/ucontext.h... yes
checking cygwin/signal.h usability... no
checking cygwin/signal.h presence... no
checking for cygwin/signal.h... no
checking how to access the program counter from a struct ucontext... uc_mcontext.gregs[REG_RIP]
checking libunwind.h usability... no
checking libunwind.h presence... no
checking for libunwind.h... no
checking for x86 without frame pointers... yes
checking if the compiler supports -Wno-unused-result... yes
configure: Will not build sized deallocation operators
checking if C++ compiler supports -fsized-deallocation... yes
checking if target has _Unwind_Backtrace... yes
checking printf format code for printing a size_t and ssize_t... l
checking for __builtin_stack_pointer()... no
checking for __builtin_expect()... yes
checking for __environ... yes
checking for __thread... yes
checking if nanosleep requires any libraries... no
checking whether uname is declared... yes
checking for the pthreads library -lpthreads... no
checking whether pthreads work without any flags... no
checking whether pthreads work with -Kthread... no
checking whether pthreads work with -kthread... no
checking for the pthreads library -llthread... no
checking whether pthreads work with -pthread... yes
checking for joinable pthread attribute... PTHREAD_CREATE_JOINABLE
checking if more special flags are required for pthreads... no
checking whether to check for GCC pthread/shared inconsistencies... yes
checking whether -pthread is sufficient with -shared... yes
checking whether what we have so far is sufficient with -nostdlib... no
checking whether -lpthread saves the day... yes
checking whether pthread symbols are available in C++ without including pthread.h... no
checking whether the compiler implements namespaces... yes
checking what namespace STL code is in... std
checking for program_invocation_name... yes
checking for Linux SIGEV_THREAD_ID... yes
checking that generated files are newer than configure... done
configure: creating ./config.status
config.status: creating Makefile
config.status: creating src/gperftools/tcmalloc.h
config.status: creating src/windows/gperftools/tcmalloc.h
config.status: creating src/config.h
config.status: executing depfiles commands
config.status: executing libtool commands
configure: WARNING: No frame pointers and no libunwind. Using experimental backtrace capturing via libgcc. Expect crashy cpu profiler.
[vagrant@localhost gperftools]$ make
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_la-tcmalloc.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_la-tcmalloc.Tpo -c -o src/libtcmalloc_minimal_la-tcmalloc.lo `test -f 'src/tcmalloc.cc' || echo './'`src/tcmalloc.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_la-tcmalloc.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_la-tcmalloc.Tpo -c src/tcmalloc.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_la-tcmalloc.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_la-tcmalloc.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_la-tcmalloc.Tpo -c src/tcmalloc.cc -o src/libtcmalloc_minimal_la-tcmalloc.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_la-tcmalloc.Tpo src/.deps/libtcmalloc_minimal_la-tcmalloc.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-common.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-common.Tpo -c -o src/libtcmalloc_minimal_internal_la-common.lo `test -f 'src/common.cc' || echo './'`src/common.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-common.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-common.Tpo -c src/common.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-common.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-common.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-common.Tpo -c src/common.cc -o src/libtcmalloc_minimal_internal_la-common.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-common.Tpo src/.deps/libtcmalloc_minimal_internal_la-common.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-internal_logging.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-internal_logging.Tpo -c -o src/libtcmalloc_minimal_internal_la-internal_logging.lo `test -f 'src/internal_logging.cc' || echo './'`src/internal_logging.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-internal_logging.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-internal_logging.Tpo -c src/internal_logging.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-internal_logging.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-internal_logging.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-internal_logging.Tpo -c src/internal_logging.cc -o src/libtcmalloc_minimal_internal_la-internal_logging.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-internal_logging.Tpo src/.deps/libtcmalloc_minimal_internal_la-internal_logging.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-system-alloc.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-system-alloc.Tpo -c -o src/libtcmalloc_minimal_internal_la-system-alloc.lo `test -f 'src/system-alloc.cc' || echo './'`src/system-alloc.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-system-alloc.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-system-alloc.Tpo -c src/system-alloc.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-system-alloc.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-system-alloc.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-system-alloc.Tpo -c src/system-alloc.cc -o src/libtcmalloc_minimal_internal_la-system-alloc.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-system-alloc.Tpo src/.deps/libtcmalloc_minimal_internal_la-system-alloc.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-memfs_malloc.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-memfs_malloc.Tpo -c -o src/libtcmalloc_minimal_internal_la-memfs_malloc.lo `test -f 'src/memfs_malloc.cc' || echo './'`src/memfs_malloc.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-memfs_malloc.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-memfs_malloc.Tpo -c src/memfs_malloc.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-memfs_malloc.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-memfs_malloc.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-memfs_malloc.Tpo -c src/memfs_malloc.cc -o src/libtcmalloc_minimal_internal_la-memfs_malloc.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-memfs_malloc.Tpo src/.deps/libtcmalloc_minimal_internal_la-memfs_malloc.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-central_freelist.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-central_freelist.Tpo -c -o src/libtcmalloc_minimal_internal_la-central_freelist.lo `test -f 'src/central_freelist.cc' || echo './'`src/central_freelist.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-central_freelist.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-central_freelist.Tpo -c src/central_freelist.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-central_freelist.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-central_freelist.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-central_freelist.Tpo -c src/central_freelist.cc -o src/libtcmalloc_minimal_internal_la-central_freelist.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-central_freelist.Tpo src/.deps/libtcmalloc_minimal_internal_la-central_freelist.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-page_heap.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-page_heap.Tpo -c -o src/libtcmalloc_minimal_internal_la-page_heap.lo `test -f 'src/page_heap.cc' || echo './'`src/page_heap.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-page_heap.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-page_heap.Tpo -c src/page_heap.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-page_heap.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-page_heap.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-page_heap.Tpo -c src/page_heap.cc -o src/libtcmalloc_minimal_internal_la-page_heap.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-page_heap.Tpo src/.deps/libtcmalloc_minimal_internal_la-page_heap.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-sampler.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-sampler.Tpo -c -o src/libtcmalloc_minimal_internal_la-sampler.lo `test -f 'src/sampler.cc' || echo './'`src/sampler.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-sampler.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-sampler.Tpo -c src/sampler.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-sampler.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-sampler.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-sampler.Tpo -c src/sampler.cc -o src/libtcmalloc_minimal_internal_la-sampler.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-sampler.Tpo src/.deps/libtcmalloc_minimal_internal_la-sampler.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-span.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-span.Tpo -c -o src/libtcmalloc_minimal_internal_la-span.lo `test -f 'src/span.cc' || echo './'`src/span.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-span.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-span.Tpo -c src/span.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-span.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-span.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-span.Tpo -c src/span.cc -o src/libtcmalloc_minimal_internal_la-span.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-span.Tpo src/.deps/libtcmalloc_minimal_internal_la-span.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-stack_trace_table.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-stack_trace_table.Tpo -c -o src/libtcmalloc_minimal_internal_la-stack_trace_table.lo `test -f 'src/stack_trace_table.cc' || echo './'`src/stack_trace_table.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-stack_trace_table.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-stack_trace_table.Tpo -c src/stack_trace_table.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-stack_trace_table.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-stack_trace_table.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-stack_trace_table.Tpo -c src/stack_trace_table.cc -o src/libtcmalloc_minimal_internal_la-stack_trace_table.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-stack_trace_table.Tpo src/.deps/libtcmalloc_minimal_internal_la-stack_trace_table.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-static_vars.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-static_vars.Tpo -c -o src/libtcmalloc_minimal_internal_la-static_vars.lo `test -f 'src/static_vars.cc' || echo './'`src/static_vars.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-static_vars.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-static_vars.Tpo -c src/static_vars.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-static_vars.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-static_vars.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-static_vars.Tpo -c src/static_vars.cc -o src/libtcmalloc_minimal_internal_la-static_vars.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-static_vars.Tpo src/.deps/libtcmalloc_minimal_internal_la-static_vars.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-symbolize.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-symbolize.Tpo -c -o src/libtcmalloc_minimal_internal_la-symbolize.lo `test -f 'src/symbolize.cc' || echo './'`src/symbolize.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-symbolize.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-symbolize.Tpo -c src/symbolize.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-symbolize.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-symbolize.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-symbolize.Tpo -c src/symbolize.cc -o src/libtcmalloc_minimal_internal_la-symbolize.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-symbolize.Tpo src/.deps/libtcmalloc_minimal_internal_la-symbolize.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-thread_cache.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-thread_cache.Tpo -c -o src/libtcmalloc_minimal_internal_la-thread_cache.lo `test -f 'src/thread_cache.cc' || echo './'`src/thread_cache.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-thread_cache.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-thread_cache.Tpo -c src/thread_cache.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-thread_cache.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-thread_cache.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-thread_cache.Tpo -c src/thread_cache.cc -o src/libtcmalloc_minimal_internal_la-thread_cache.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-thread_cache.Tpo src/.deps/libtcmalloc_minimal_internal_la-thread_cache.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-malloc_hook.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-malloc_hook.Tpo -c -o src/libtcmalloc_minimal_internal_la-malloc_hook.lo `test -f 'src/malloc_hook.cc' || echo './'`src/malloc_hook.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-malloc_hook.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-malloc_hook.Tpo -c src/malloc_hook.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-malloc_hook.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-malloc_hook.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-malloc_hook.Tpo -c src/malloc_hook.cc -o src/libtcmalloc_minimal_internal_la-malloc_hook.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-malloc_hook.Tpo src/.deps/libtcmalloc_minimal_internal_la-malloc_hook.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-malloc_extension.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-malloc_extension.Tpo -c -o src/libtcmalloc_minimal_internal_la-malloc_extension.lo `test -f 'src/malloc_extension.cc' || echo './'`src/malloc_extension.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-malloc_extension.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-malloc_extension.Tpo -c src/malloc_extension.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_internal_la-malloc_extension.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_minimal_internal_la-malloc_extension.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_internal_la-malloc_extension.Tpo -c src/malloc_extension.cc -o src/libtcmalloc_minimal_internal_la-malloc_extension.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_internal_la-malloc_extension.Tpo src/.deps/libtcmalloc_minimal_internal_la-malloc_extension.Plo
depbase=`echo src/base/spinlock.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/base/spinlock.lo -MD -MP -MF $depbase.Tpo -c -o src/base/spinlock.lo src/base/spinlock.cc &&\
mv -f $depbase.Tpo $depbase.Plo
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/spinlock.lo -MD -MP -MF src/base/.deps/spinlock.Tpo -c src/base/spinlock.cc  -fPIC -DPIC -o src/base/.libs/spinlock.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/spinlock.lo -MD -MP -MF src/base/.deps/spinlock.Tpo -c src/base/spinlock.cc -o src/base/spinlock.o >/dev/null 2>&1
depbase=`echo src/base/spinlock_internal.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/base/spinlock_internal.lo -MD -MP -MF $depbase.Tpo -c -o src/base/spinlock_internal.lo src/base/spinlock_internal.cc &&\
mv -f $depbase.Tpo $depbase.Plo
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/spinlock_internal.lo -MD -MP -MF src/base/.deps/spinlock_internal.Tpo -c src/base/spinlock_internal.cc  -fPIC -DPIC -o src/base/.libs/spinlock_internal.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/spinlock_internal.lo -MD -MP -MF src/base/.deps/spinlock_internal.Tpo -c src/base/spinlock_internal.cc -o src/base/spinlock_internal.o >/dev/null 2>&1
depbase=`echo src/base/atomicops-internals-x86.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/base/atomicops-internals-x86.lo -MD -MP -MF $depbase.Tpo -c -o src/base/atomicops-internals-x86.lo src/base/atomicops-internals-x86.cc &&\
mv -f $depbase.Tpo $depbase.Plo
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/atomicops-internals-x86.lo -MD -MP -MF src/base/.deps/atomicops-internals-x86.Tpo -c src/base/atomicops-internals-x86.cc  -fPIC -DPIC -o src/base/.libs/atomicops-internals-x86.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/atomicops-internals-x86.lo -MD -MP -MF src/base/.deps/atomicops-internals-x86.Tpo -c src/base/atomicops-internals-x86.cc -o src/base/atomicops-internals-x86.o >/dev/null 2>&1
/bin/sh ./libtool  --tag=CXX   --mode=link g++ -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -no-undefined   -o libspinlock.la  src/base/spinlock.lo src/base/spinlock_internal.lo src/base/atomicops-internals-x86.lo   
libtool: link: ar cru .libs/libspinlock.a src/base/.libs/spinlock.o src/base/.libs/spinlock_internal.o src/base/.libs/atomicops-internals-x86.o 
ar: `u' modifier ignored since `D' is the default (see `U')
libtool: link: ranlib .libs/libspinlock.a
libtool: link: ( cd ".libs" && rm -f "libspinlock.la" && ln -s "../libspinlock.la" "libspinlock.la" )
depbase=`echo src/base/sysinfo.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/base/sysinfo.lo -MD -MP -MF $depbase.Tpo -c -o src/base/sysinfo.lo src/base/sysinfo.cc &&\
mv -f $depbase.Tpo $depbase.Plo
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/sysinfo.lo -MD -MP -MF src/base/.deps/sysinfo.Tpo -c src/base/sysinfo.cc  -fPIC -DPIC -o src/base/.libs/sysinfo.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/sysinfo.lo -MD -MP -MF src/base/.deps/sysinfo.Tpo -c src/base/sysinfo.cc -o src/base/sysinfo.o >/dev/null 2>&1
/bin/sh ./libtool  --tag=CXX   --mode=link g++ -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -no-undefined   -o libsysinfo.la  src/base/sysinfo.lo    
libtool: link: ar cru .libs/libsysinfo.a src/base/.libs/sysinfo.o 
ar: `u' modifier ignored since `D' is the default (see `U')
libtool: link: ranlib .libs/libsysinfo.a
libtool: link: ( cd ".libs" && rm -f "libsysinfo.la" && ln -s "../libsysinfo.la" "libsysinfo.la" )
depbase=`echo src/base/logging.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/base/logging.lo -MD -MP -MF $depbase.Tpo -c -o src/base/logging.lo src/base/logging.cc &&\
mv -f $depbase.Tpo $depbase.Plo
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/logging.lo -MD -MP -MF src/base/.deps/logging.Tpo -c src/base/logging.cc  -fPIC -DPIC -o src/base/.libs/logging.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/logging.lo -MD -MP -MF src/base/.deps/logging.Tpo -c src/base/logging.cc -o src/base/logging.o >/dev/null 2>&1
depbase=`echo src/base/dynamic_annotations.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ./libtool  --tag=CC   --mode=compile gcc -DHAVE_CONFIG_H -I. -I./src  -I./src    -g -O2 -MT src/base/dynamic_annotations.lo -MD -MP -MF $depbase.Tpo -c -o src/base/dynamic_annotations.lo src/base/dynamic_annotations.c &&\
mv -f $depbase.Tpo $depbase.Plo
libtool: compile:  gcc -DHAVE_CONFIG_H -I. -I./src -I./src -g -O2 -MT src/base/dynamic_annotations.lo -MD -MP -MF src/base/.deps/dynamic_annotations.Tpo -c src/base/dynamic_annotations.c  -fPIC -DPIC -o src/base/.libs/dynamic_annotations.o
libtool: compile:  gcc -DHAVE_CONFIG_H -I. -I./src -I./src -g -O2 -MT src/base/dynamic_annotations.lo -MD -MP -MF src/base/.deps/dynamic_annotations.Tpo -c src/base/dynamic_annotations.c -o src/base/dynamic_annotations.o >/dev/null 2>&1
/bin/sh ./libtool  --tag=CXX   --mode=link g++ -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -no-undefined   -o liblogging.la  src/base/logging.lo src/base/dynamic_annotations.lo   
libtool: link: ar cru .libs/liblogging.a src/base/.libs/logging.o src/base/.libs/dynamic_annotations.o 
ar: `u' modifier ignored since `D' is the default (see `U')
libtool: link: ranlib .libs/liblogging.a
libtool: link: ( cd ".libs" && rm -f "liblogging.la" && ln -s "../liblogging.la" "liblogging.la" )
depbase=`echo src/maybe_threads.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -MT src/maybe_threads.lo -MD -MP -MF $depbase.Tpo -c -o src/maybe_threads.lo src/maybe_threads.cc &&\
mv -f $depbase.Tpo $depbase.Plo
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/maybe_threads.lo -MD -MP -MF src/.deps/maybe_threads.Tpo -c src/maybe_threads.cc  -fPIC -DPIC -o src/.libs/maybe_threads.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/maybe_threads.lo -MD -MP -MF src/.deps/maybe_threads.Tpo -c src/maybe_threads.cc -o src/maybe_threads.o >/dev/null 2>&1
/bin/sh ./libtool  --tag=CXX   --mode=link g++ -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -no-undefined   -o libmaybe_threads.la   src/maybe_threads.lo  
libtool: link: ar cru .libs/libmaybe_threads.a src/.libs/maybe_threads.o 
ar: `u' modifier ignored since `D' is the default (see `U')
libtool: link: ranlib .libs/libmaybe_threads.a
libtool: link: ( cd ".libs" && rm -f "libmaybe_threads.la" && ln -s "../libmaybe_threads.la" "libmaybe_threads.la" )
/bin/sh ./libtool  --tag=CXX   --mode=link g++ -DNO_TCMALLOC_SAMPLES -DNO_HEAP_CHECK -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -no-undefined   -o libtcmalloc_minimal_internal.la  src/libtcmalloc_minimal_internal_la-common.lo src/libtcmalloc_minimal_internal_la-internal_logging.lo src/libtcmalloc_minimal_internal_la-system-alloc.lo src/libtcmalloc_minimal_internal_la-memfs_malloc.lo src/libtcmalloc_minimal_internal_la-central_freelist.lo src/libtcmalloc_minimal_internal_la-page_heap.lo src/libtcmalloc_minimal_internal_la-sampler.lo src/libtcmalloc_minimal_internal_la-span.lo src/libtcmalloc_minimal_internal_la-stack_trace_table.lo src/libtcmalloc_minimal_internal_la-static_vars.lo src/libtcmalloc_minimal_internal_la-symbolize.lo src/libtcmalloc_minimal_internal_la-thread_cache.lo src/libtcmalloc_minimal_internal_la-malloc_hook.lo src/libtcmalloc_minimal_internal_la-malloc_extension.lo    libspinlock.la libsysinfo.la liblogging.la libmaybe_threads.la 
libtool: link: (cd .libs/libtcmalloc_minimal_internal.lax/libspinlock.a && ar x "/Users/fanhongling/Downloads/workspace/src/github.com/gperftools/gperftools/./.libs/libspinlock.a")
libtool: link: (cd .libs/libtcmalloc_minimal_internal.lax/libsysinfo.a && ar x "/Users/fanhongling/Downloads/workspace/src/github.com/gperftools/gperftools/./.libs/libsysinfo.a")
libtool: link: (cd .libs/libtcmalloc_minimal_internal.lax/liblogging.a && ar x "/Users/fanhongling/Downloads/workspace/src/github.com/gperftools/gperftools/./.libs/liblogging.a")
libtool: link: (cd .libs/libtcmalloc_minimal_internal.lax/libmaybe_threads.a && ar x "/Users/fanhongling/Downloads/workspace/src/github.com/gperftools/gperftools/./.libs/libmaybe_threads.a")
libtool: link: ar cru .libs/libtcmalloc_minimal_internal.a src/.libs/libtcmalloc_minimal_internal_la-common.o src/.libs/libtcmalloc_minimal_internal_la-internal_logging.o src/.libs/libtcmalloc_minimal_internal_la-system-alloc.o src/.libs/libtcmalloc_minimal_internal_la-memfs_malloc.o src/.libs/libtcmalloc_minimal_internal_la-central_freelist.o src/.libs/libtcmalloc_minimal_internal_la-page_heap.o src/.libs/libtcmalloc_minimal_internal_la-sampler.o src/.libs/libtcmalloc_minimal_internal_la-span.o src/.libs/libtcmalloc_minimal_internal_la-stack_trace_table.o src/.libs/libtcmalloc_minimal_internal_la-static_vars.o src/.libs/libtcmalloc_minimal_internal_la-symbolize.o src/.libs/libtcmalloc_minimal_internal_la-thread_cache.o src/.libs/libtcmalloc_minimal_internal_la-malloc_hook.o src/.libs/libtcmalloc_minimal_internal_la-malloc_extension.o   .libs/libtcmalloc_minimal_internal.lax/libspinlock.a/atomicops-internals-x86.o .libs/libtcmalloc_minimal_internal.lax/libspinlock.a/spinlock.o .libs/libtcmalloc_minimal_internal.lax/libspinlock.a/spinlock_internal.o  .libs/libtcmalloc_minimal_internal.lax/libsysinfo.a/sysinfo.o  .libs/libtcmalloc_minimal_internal.lax/liblogging.a/dynamic_annotations.o .libs/libtcmalloc_minimal_internal.lax/liblogging.a/logging.o  .libs/libtcmalloc_minimal_internal.lax/libmaybe_threads.a/maybe_threads.o 
ar: `u' modifier ignored since `D' is the default (see `U')
libtool: link: ranlib .libs/libtcmalloc_minimal_internal.a
libtool: link: rm -fr .libs/libtcmalloc_minimal_internal.lax
libtool: link: ( cd ".libs" && rm -f "libtcmalloc_minimal_internal.la" && ln -s "../libtcmalloc_minimal_internal.la" "libtcmalloc_minimal_internal.la" )
/bin/sh ./libtool  --tag=CXX   --mode=link g++ -DNO_TCMALLOC_SAMPLES -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -g -O2 -version-info 8:5:4 -no-undefined   -o libtcmalloc_minimal.la -rpath /usr/local/lib src/libtcmalloc_minimal_la-tcmalloc.lo    libtcmalloc_minimal_internal.la 
libtool: link: g++  -fPIC -DPIC -shared -nostdlib /usr/lib/gcc/x86_64-redhat-linux/5.3.1/../../../../lib64/crti.o /usr/lib/gcc/x86_64-redhat-linux/5.3.1/crtbeginS.o  src/.libs/libtcmalloc_minimal_la-tcmalloc.o  -Wl,--whole-archive ./.libs/libtcmalloc_minimal_internal.a -Wl,--no-whole-archive  -L/usr/lib/gcc/x86_64-redhat-linux/5.3.1 -L/usr/lib/gcc/x86_64-redhat-linux/5.3.1/../../../../lib64 -L/lib/../lib64 -L/usr/lib/../lib64 -L/usr/lib/gcc/x86_64-redhat-linux/5.3.1/../../.. -lstdc++ -lm -lc -lgcc_s /usr/lib/gcc/x86_64-redhat-linux/5.3.1/crtendS.o /usr/lib/gcc/x86_64-redhat-linux/5.3.1/../../../../lib64/crtn.o  -pthread -g -O2   -pthread -Wl,-soname -Wl,libtcmalloc_minimal.so.4 -o .libs/libtcmalloc_minimal.so.4.4.5
libtool: link: (cd ".libs" && rm -f "libtcmalloc_minimal.so.4" && ln -s "libtcmalloc_minimal.so.4.4.5" "libtcmalloc_minimal.so.4")
libtool: link: (cd ".libs" && rm -f "libtcmalloc_minimal.so" && ln -s "libtcmalloc_minimal.so.4.4.5" "libtcmalloc_minimal.so")
libtool: link: (cd .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a && ar x "/Users/fanhongling/Downloads/workspace/src/github.com/gperftools/gperftools/./.libs/libtcmalloc_minimal_internal.a")
libtool: link: ar cru .libs/libtcmalloc_minimal.a  src/libtcmalloc_minimal_la-tcmalloc.o  .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/atomicops-internals-x86.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/dynamic_annotations.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-central_freelist.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-common.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-internal_logging.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-malloc_extension.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-malloc_hook.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-memfs_malloc.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-page_heap.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-sampler.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-span.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-stack_trace_table.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-static_vars.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-symbolize.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-system-alloc.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-thread_cache.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/logging.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/maybe_threads.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/spinlock.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/spinlock_internal.o .libs/libtcmalloc_minimal.lax/libtcmalloc_minimal_internal.a/sysinfo.o 
ar: `u' modifier ignored since `D' is the default (see `U')
libtool: link: ranlib .libs/libtcmalloc_minimal.a
libtool: link: rm -fr .libs/libtcmalloc_minimal.lax
libtool: link: ( cd ".libs" && rm -f "libtcmalloc_minimal.la" && ln -s "../libtcmalloc_minimal.la" "libtcmalloc_minimal.la" )
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -DNO_TCMALLOC_SAMPLES -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -DTCMALLOC_FOR_DEBUGALLOCATION -g -O2 -MT src/libtcmalloc_minimal_debug_la-debugallocation.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_debug_la-debugallocation.Tpo -c -o src/libtcmalloc_minimal_debug_la-debugallocation.lo `test -f 'src/debugallocation.cc' || echo './'`src/debugallocation.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -DTCMALLOC_FOR_DEBUGALLOCATION -g -O2 -MT src/libtcmalloc_minimal_debug_la-debugallocation.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_debug_la-debugallocation.Tpo -c src/debugallocation.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_minimal_debug_la-debugallocation.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -DNO_TCMALLOC_SAMPLES -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -DTCMALLOC_FOR_DEBUGALLOCATION -g -O2 -MT src/libtcmalloc_minimal_debug_la-debugallocation.lo -MD -MP -MF src/.deps/libtcmalloc_minimal_debug_la-debugallocation.Tpo -c src/debugallocation.cc -o src/libtcmalloc_minimal_debug_la-debugallocation.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_minimal_debug_la-debugallocation.Tpo src/.deps/libtcmalloc_minimal_debug_la-debugallocation.Plo
/bin/sh ./libtool  --tag=CXX   --mode=link g++ -DNO_TCMALLOC_SAMPLES -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER -DTCMALLOC_FOR_DEBUGALLOCATION -g -O2 -version-info 8:5:4 -no-undefined  -version-info 8:5:4  -o libtcmalloc_minimal_debug.la -rpath /usr/local/lib src/libtcmalloc_minimal_debug_la-debugallocation.lo    libtcmalloc_minimal_internal.la 
libtool: link: g++  -fPIC -DPIC -shared -nostdlib /usr/lib/gcc/x86_64-redhat-linux/5.3.1/../../../../lib64/crti.o /usr/lib/gcc/x86_64-redhat-linux/5.3.1/crtbeginS.o  src/.libs/libtcmalloc_minimal_debug_la-debugallocation.o  -Wl,--whole-archive ./.libs/libtcmalloc_minimal_internal.a -Wl,--no-whole-archive  -L/usr/lib/gcc/x86_64-redhat-linux/5.3.1 -L/usr/lib/gcc/x86_64-redhat-linux/5.3.1/../../../../lib64 -L/lib/../lib64 -L/usr/lib/../lib64 -L/usr/lib/gcc/x86_64-redhat-linux/5.3.1/../../.. -lstdc++ -lm -lc -lgcc_s /usr/lib/gcc/x86_64-redhat-linux/5.3.1/crtendS.o /usr/lib/gcc/x86_64-redhat-linux/5.3.1/../../../../lib64/crtn.o  -pthread -g -O2   -pthread -Wl,-soname -Wl,libtcmalloc_minimal_debug.so.4 -o .libs/libtcmalloc_minimal_debug.so.4.4.5
libtool: link: (cd ".libs" && rm -f "libtcmalloc_minimal_debug.so.4" && ln -s "libtcmalloc_minimal_debug.so.4.4.5" "libtcmalloc_minimal_debug.so.4")
libtool: link: (cd ".libs" && rm -f "libtcmalloc_minimal_debug.so" && ln -s "libtcmalloc_minimal_debug.so.4.4.5" "libtcmalloc_minimal_debug.so")
libtool: link: (cd .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a && ar x "/Users/fanhongling/Downloads/workspace/src/github.com/gperftools/gperftools/./.libs/libtcmalloc_minimal_internal.a")
libtool: link: ar cru .libs/libtcmalloc_minimal_debug.a  src/libtcmalloc_minimal_debug_la-debugallocation.o  .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/atomicops-internals-x86.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/dynamic_annotations.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-central_freelist.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-common.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-internal_logging.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-malloc_extension.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-malloc_hook.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-memfs_malloc.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-page_heap.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-sampler.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-span.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-stack_trace_table.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-static_vars.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-symbolize.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-system-alloc.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/libtcmalloc_minimal_internal_la-thread_cache.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/logging.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/maybe_threads.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/spinlock.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/spinlock_internal.o .libs/libtcmalloc_minimal_debug.lax/libtcmalloc_minimal_internal.a/sysinfo.o 
ar: `u' modifier ignored since `D' is the default (see `U')
libtool: link: ranlib .libs/libtcmalloc_minimal_debug.a
libtool: link: rm -fr .libs/libtcmalloc_minimal_debug.lax
libtool: link: ( cd ".libs" && rm -f "libtcmalloc_minimal_debug.la" && ln -s "../libtcmalloc_minimal_debug.la" "libtcmalloc_minimal_debug.la" )
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER   -g -O2 -MT src/libtcmalloc_la-tcmalloc.lo -MD -MP -MF src/.deps/libtcmalloc_la-tcmalloc.Tpo -c -o src/libtcmalloc_la-tcmalloc.lo `test -f 'src/tcmalloc.cc' || echo './'`src/tcmalloc.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_la-tcmalloc.lo -MD -MP -MF src/.deps/libtcmalloc_la-tcmalloc.Tpo -c src/tcmalloc.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_la-tcmalloc.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_la-tcmalloc.lo -MD -MP -MF src/.deps/libtcmalloc_la-tcmalloc.Tpo -c src/tcmalloc.cc -o src/libtcmalloc_la-tcmalloc.o >/dev/null 2>&1
mv -f src/.deps/libtcmalloc_la-tcmalloc.Tpo src/.deps/libtcmalloc_la-tcmalloc.Plo
depbase=`echo src/base/thread_lister.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ./libtool  --tag=CC   --mode=compile gcc -DHAVE_CONFIG_H -I. -I./src  -I./src    -g -O2 -MT src/base/thread_lister.lo -MD -MP -MF $depbase.Tpo -c -o src/base/thread_lister.lo src/base/thread_lister.c &&\
mv -f $depbase.Tpo $depbase.Plo
libtool: compile:  gcc -DHAVE_CONFIG_H -I. -I./src -I./src -g -O2 -MT src/base/thread_lister.lo -MD -MP -MF src/base/.deps/thread_lister.Tpo -c src/base/thread_lister.c  -fPIC -DPIC -o src/base/.libs/thread_lister.o
libtool: compile:  gcc -DHAVE_CONFIG_H -I. -I./src -I./src -g -O2 -MT src/base/thread_lister.lo -MD -MP -MF src/base/.deps/thread_lister.Tpo -c src/base/thread_lister.c -o src/base/thread_lister.o >/dev/null 2>&1
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER   -g -O2 -MT src/base/libtcmalloc_la-linuxthreads.lo -MD -MP -MF src/base/.deps/libtcmalloc_la-linuxthreads.Tpo -c -o src/base/libtcmalloc_la-linuxthreads.lo `test -f 'src/base/linuxthreads.cc' || echo './'`src/base/linuxthreads.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/libtcmalloc_la-linuxthreads.lo -MD -MP -MF src/base/.deps/libtcmalloc_la-linuxthreads.Tpo -c src/base/linuxthreads.cc  -fPIC -DPIC -o src/base/.libs/libtcmalloc_la-linuxthreads.o
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/base/libtcmalloc_la-linuxthreads.lo -MD -MP -MF src/base/.deps/libtcmalloc_la-linuxthreads.Tpo -c src/base/linuxthreads.cc -o src/base/libtcmalloc_la-linuxthreads.o >/dev/null 2>&1
mv -f src/base/.deps/libtcmalloc_la-linuxthreads.Tpo src/base/.deps/libtcmalloc_la-linuxthreads.Plo
/bin/sh ./libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I./src  -I./src   -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc  -Wno-unused-result -fsized-deallocation  -DNO_FRAME_POINTER   -g -O2 -MT src/libtcmalloc_la-heap-checker.lo -MD -MP -MF src/.deps/libtcmalloc_la-heap-checker.Tpo -c -o src/libtcmalloc_la-heap-checker.lo `test -f 'src/heap-checker.cc' || echo './'`src/heap-checker.cc
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I./src -I./src -pthread -DNDEBUG -Wall -Wwrite-strings -Woverloaded-virtual -Wno-sign-compare -fno-builtin-malloc -fno-builtin-free -fno-builtin-realloc -fno-builtin-calloc -fno-builtin-cfree -fno-builtin-memalign -fno-builtin-posix_memalign -fno-builtin-valloc -fno-builtin-pvalloc -Wno-unused-result -fsized-deallocation -DNO_FRAME_POINTER -g -O2 -MT src/libtcmalloc_la-heap-checker.lo -MD -MP -MF src/.deps/libtcmalloc_la-heap-checker.Tpo -c src/heap-checker.cc  -fPIC -DPIC -o src/.libs/libtcmalloc_la-heap-checker.o
src/heap-checker.cc: 在静态成员函数‘static void HeapLeakChecker::IgnoreLiveObjectsLocked(const char*, const char*)’中:
src/heap-checker.cc:1436:52: 错误：‘MS_ASYNC’在此作用域中尚未声明
         if (msync(const_cast<char*>(object), size, MS_ASYNC) != 0) {
                                                    ^
src/heap-checker.cc:1436:60: 错误：‘msync’在此作用域中尚未声明
         if (msync(const_cast<char*>(object), size, MS_ASYNC) != 0) {
                                                            ^
Makefile:4547: recipe for target 'src/libtcmalloc_la-heap-checker.lo' failed
make: *** [src/libtcmalloc_la-heap-checker.lo] Error 1
```

```
[vagrant@localhost bazel]$ sudo dnf list | grep perftools
gperftools.x86_64                        2.4-5.fc23                      fedora 
gperftools-devel.i686                    2.4-5.fc23                      fedora 
gperftools-devel.x86_64                  2.4-5.fc23                      fedora 
gperftools-libs.i686                     2.4-5.fc23                      fedora 
gperftools-libs.x86_64                   2.4-5.fc23                      fedora 
[vagrant@localhost bazel]$ sudo dnf install -y gperftools-devel
Kubernetes                                                                                 17 kB/s |  18 kB     00:01    
上次元数据过期检查在 0:00:00 前执行于 Sat Nov 11 06:25:51 2017。
依赖关系解决。
==========================================================================================================================
 Package                            架构                     版本                          仓库                      大小
==========================================================================================================================
安装:
 gperftools-devel                   x86_64                   2.4-5.fc23                    fedora                   302 k
 gperftools-libs                    x86_64                   2.4-5.fc23                    fedora                   273 k

事务概要
==========================================================================================================================
安装  2 Packages

总下载：574 k
安装大小：2.0 M
下载软件包：
(1/2): gperftools-libs-2.4-5.fc23.x86_64.rpm                                               31 kB/s | 273 kB     00:08    
(2/2): gperftools-devel-2.4-5.fc23.x86_64.rpm                                              30 kB/s | 302 kB     00:09    
--------------------------------------------------------------------------------------------------------------------------
总计                                                                                       51 kB/s | 574 kB     00:11     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: gperftools-libs-2.4-5.fc23.x86_64                                                                             1/2 
  安装: gperftools-devel-2.4-5.fc23.x86_64                                                                            2/2 
  验证: gperftools-devel-2.4-5.fc23.x86_64                                                                            1/2 
  验证: gperftools-libs-2.4-5.fc23.x86_64                                                                             2/2 

已安装:
  gperftools-devel.x86_64 2.4-5.fc23                           gperftools-libs.x86_64 2.4-5.fc23                          

完毕！
[vagrant@localhost bazel]$ sudo dnf install -y gperftools
上次元数据过期检查在 0:00:21 前执行于 Sat Nov 11 06:25:51 2017。
依赖关系解决。
==========================================================================================================================
 Package                                    架构              版本                               仓库                大小
==========================================================================================================================
安装:
 Xaw3d                                      x86_64            1.6.2-8.fc23                       fedora             176 k
 gperftools                                 x86_64            2.4-5.fc23                         fedora              13 k
 graphviz                                   x86_64            2.38.0-29.fc23                     fedora             3.2 M
 gts                                        x86_64            0.7.6-26.20111025.fc23             fedora             225 k
 gv                                         x86_64            3.7.4-10.fc23                      fedora             242 k
 lasi                                       x86_64            1.1.2-5.fc23                       fedora              46 k
 libXaw                                     x86_64            1.0.13-2.fc23                      fedora             191 k
 librsvg2                                   x86_64            2.40.16-1.fc23                     updates            134 k
 libtool-ltdl                               x86_64            2.4.6-8.fc23                       updates             54 k
 netpbm                                     x86_64            10.75.99-1.fc23                    updates            189 k
 pprof                                      noarch            2.4-5.fc23                         fedora              61 k
 xorg-x11-fonts-ISO8859-1-100dpi            noarch            7.5-15.fc23                        fedora             1.1 M

事务概要
==========================================================================================================================
安装  12 Packages

总下载：5.6 M
安装大小：66 M
下载软件包：
(1/12): gperftools-2.4-5.fc23.x86_64.rpm                                                   42 kB/s |  13 kB     00:00    
(2/12): pprof-2.4-5.fc23.noarch.rpm                                                        28 kB/s |  61 kB     00:02    
(3/12): Xaw3d-1.6.2-8.fc23.x86_64.rpm                                                      32 kB/s | 176 kB     00:05    
(4/12): gv-3.7.4-10.fc23.x86_64.rpm                                                        37 kB/s | 242 kB     00:06    
(5/12): lasi-1.1.2-5.fc23.x86_64.rpm                                                      6.0 kB/s |  46 kB     00:07    
(6/12): gts-0.7.6-26.20111025.fc23.x86_64.rpm                                              14 kB/s | 225 kB     00:15    
(7/12): libXaw-1.0.13-2.fc23.x86_64.rpm                                                    21 kB/s | 191 kB     00:09    
(8/12): netpbm-10.75.99-1.fc23.x86_64.rpm                                                  59 kB/s | 189 kB     00:03    
(9/12): librsvg2-2.40.16-1.fc23.x86_64.rpm                                                120 kB/s | 134 kB     00:01    
(10/12): libtool-ltdl-2.4.6-8.fc23.x86_64.rpm                                              23 kB/s |  54 kB     00:02    
(11/12): xorg-x11-fonts-ISO8859-1-100dpi-7.5-15.fc23.noarch.rpm                           101 kB/s | 1.1 MB     00:10    
(12/12): graphviz-2.38.0-29.fc23.x86_64.rpm                                                49 kB/s | 3.2 MB     01:07    
--------------------------------------------------------------------------------------------------------------------------
总计                                                                                       79 kB/s | 5.6 MB     01:12     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: libtool-ltdl-2.4.6-8.fc23.x86_64                                                                             1/12 
  安装: librsvg2-2.40.16-1.fc23.x86_64                                                                               2/12 
  安装: netpbm-10.75.99-1.fc23.x86_64                                                                                3/12 
  安装: gts-0.7.6-26.20111025.fc23.x86_64                                                                            4/12 
  安装: xorg-x11-fonts-ISO8859-1-100dpi-7.5-15.fc23.noarch                                                           5/12 
  安装: libXaw-1.0.13-2.fc23.x86_64                                                                                  6/12 
  安装: lasi-1.1.2-5.fc23.x86_64                                                                                     7/12 
  安装: graphviz-2.38.0-29.fc23.x86_64                                                                               8/12 
  安装: Xaw3d-1.6.2-8.fc23.x86_64                                                                                    9/12 
  安装: gv-3.7.4-10.fc23.x86_64                                                                                     10/12 
  安装: pprof-2.4-5.fc23.noarch                                                                                     11/12 
  安装: gperftools-2.4-5.fc23.x86_64                                                                                12/12 
  验证: gperftools-2.4-5.fc23.x86_64                                                                                 1/12 
  验证: pprof-2.4-5.fc23.noarch                                                                                      2/12 
  验证: gv-3.7.4-10.fc23.x86_64                                                                                      3/12 
  验证: Xaw3d-1.6.2-8.fc23.x86_64                                                                                    4/12 
  验证: graphviz-2.38.0-29.fc23.x86_64                                                                               5/12 
  验证: gts-0.7.6-26.20111025.fc23.x86_64                                                                            6/12 
  验证: lasi-1.1.2-5.fc23.x86_64                                                                                     7/12 
  验证: libXaw-1.0.13-2.fc23.x86_64                                                                                  8/12 
  验证: xorg-x11-fonts-ISO8859-1-100dpi-7.5-15.fc23.noarch                                                           9/12 
  验证: netpbm-10.75.99-1.fc23.x86_64                                                                               10/12 
  验证: librsvg2-2.40.16-1.fc23.x86_64                                                                              11/12 
  验证: libtool-ltdl-2.4.6-8.fc23.x86_64                                                                            12/12 

已安装:
  Xaw3d.x86_64 1.6.2-8.fc23          gperftools.x86_64 2.4-5.fc23    graphviz.x86_64 2.38.0-29.fc23                     
  gts.x86_64 0.7.6-26.20111025.fc23  gv.x86_64 3.7.4-10.fc23         lasi.x86_64 1.1.2-5.fc23                           
  libXaw.x86_64 1.0.13-2.fc23        librsvg2.x86_64 2.40.16-1.fc23  libtool-ltdl.x86_64 2.4.6-8.fc23                   
  netpbm.x86_64 10.75.99-1.fc23      pprof.noarch 2.4-5.fc23         xorg-x11-fonts-ISO8859-1-100dpi.noarch 7.5-15.fc23 

完毕！
```
