# Instruction - gRPC-Java development environment

## Nov 9

Fedora 23

Protobuf
```
```
[vagrant@localhost bazel]$ protoc --version
libprotoc 3.0.0
```

Project
```
[vagrant@localhost grpc-java]$ git pull
remote: Counting objects: 11183, done.
remote: Compressing objects: 100% (8/8), done.
remote: Total 11183 (delta 3507), reused 3512 (delta 3507), pack-reused 7668
接收对象中: 100% (11183/11183), 4.77 MiB | 998.00 KiB/s, 完成.
处理 delta 中: 100% (5407/5407), 完成 593 个本地对象.
来自 https://github.com/grpc/grpc-java
   1f7c9d5..bbb98fe  master     -> origin/master
   e17ec28..f289112  gh-pages   -> origin/gh-pages
   f1d1d69..4135e7a  v1.0.x     -> origin/v1.0.x
   b4277d5..27d0c60  v1.1.x     -> origin/v1.1.x
   00fe9e9..ccdd213  v1.2.x     -> origin/v1.2.x
   c03fb61..44ce204  v1.3.x     -> origin/v1.3.x
 * [新分支]          v1.4.x     -> origin/v1.4.x
 * [新分支]          v1.5.x     -> origin/v1.5.x
 * [新分支]          v1.6.x     -> origin/v1.6.x
 * [新分支]          v1.7.x     -> origin/v1.7.x
 * [新分支]          v1.8.x     -> origin/v1.8.x
 * [新tag]           v1.3.1     -> v1.3.1
 * [新tag]           v1.4.0     -> v1.4.0
 * [新tag]           v1.5.0     -> v1.5.0
 * [新tag]           v1.6.1     -> v1.6.1
 * [新tag]           v1.7.0     -> v1.7.0
更新 1f7c9d5..bbb98fe
```

Gradle
```
[vagrant@localhost grpc-java]$ ./gradlew build
Downloading https://services.gradle.org/distributions/gradle-4.3-bin.zip
...........................
Unzipping /home/vagrant/.gradle/wrapper/dists/gradle-4.3-bin/452wx51oxqsia28686mgqhot6/gradle-4.3-bin.zip to /home/vagrant/.gradle/wrapper/dists/gradle-4.3-bin/452wx51oxqsia28686mgqhot6
Set executable permissions for: /home/vagrant/.gradle/wrapper/dists/gradle-4.3-bin/452wx51oxqsia28686mgqhot6/gradle-4.3/bin/gradle
Starting a Gradle Daemon (subsequent builds will be faster)
*** Skipping the build of codegen and compilation of proto files because skipCodegen=true
Download https://repo1.maven.org/maven2/ru/vyarus/gradle-animalsniffer-plugin/1.4.0/gradle-animalsniffer-plugin-1.4.0.pom
Download https://plugins.gradle.org/m2/me/champeau/gradle/jmh-gradle-plugin/0.4.4/jmh-gradle-plugin-0.4.4.pom
Download https://plugins.gradle.org/m2/net/ltgt/gradle/gradle-errorprone-plugin/0.0.11/gradle-errorprone-plugin-0.0.11.pom
Download https://plugins.gradle.org/m2/me/champeau/gradle/japicmp-gradle-plugin/0.2.5/japicmp-gradle-plugin-0.2.5.pom
Download https://repo1.maven.org/maven2/org/openjdk/jmh/jmh-core/1.19/jmh-core-1.19.pom
Download https://repo1.maven.org/maven2/org/openjdk/jmh/jmh-parent/1.19/jmh-parent-1.19.pom
Download https://repo1.maven.org/maven2/com/github/siom79/japicmp/japicmp/0.10.0/japicmp-0.10.0.pom
Download https://repo1.maven.org/maven2/com/github/siom79/japicmp/japicmp-base/0.10.0/japicmp-base-0.10.0.pom
Download https://repo1.maven.org/maven2/io/airlift/airline/0.7/airline-0.7.pom
Download https://repo1.maven.org/maven2/io/airlift/airbase/28/airbase-28.pom
Download https://repo1.maven.org/maven2/com/google/code/findbugs/annotations/2.0.3/annotations-2.0.3.pom
Download https://repo1.maven.org/maven2/ru/vyarus/gradle-animalsniffer-plugin/1.4.0/gradle-animalsniffer-plugin-1.4.0.jar
Download https://plugins.gradle.org/m2/me/champeau/gradle/jmh-gradle-plugin/0.4.4/jmh-gradle-plugin-0.4.4.jar
Download https://plugins.gradle.org/m2/me/champeau/gradle/japicmp-gradle-plugin/0.2.5/japicmp-gradle-plugin-0.2.5.jar
Download https://repo1.maven.org/maven2/org/openjdk/jmh/jmh-core/1.19/jmh-core-1.19.jar
Download https://repo1.maven.org/maven2/com/github/siom79/japicmp/japicmp/0.10.0/japicmp-0.10.0.jar
Download https://repo1.maven.org/maven2/io/airlift/airline/0.7/airline-0.7.jar
Download https://repo1.maven.org/maven2/com/google/code/findbugs/annotations/2.0.3/annotations-2.0.3.jar
Download https://plugins.gradle.org/m2/net/ltgt/gradle/gradle-errorprone-plugin/0.0.11/gradle-errorprone-plugin-0.0.11.jar
Download https://repo1.maven.org/maven2/io/grpc/grpc-auth/1.6.1/grpc-auth-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-auth/1.6.1/grpc-auth-1.6.1.jar
Download https://repo1.maven.org/maven2/io/grpc/grpc-context/1.6.1/grpc-context-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-context/1.6.1/grpc-context-1.6.1.jar-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-core/1.6.1/grpc-core-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-core/1.6.1/grpc-core-1.6.1.jar
Download https://repo1.maven.org/maven2/io/grpc/grpc-grpclb/1.6.1/grpc-grpclb-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-grpclb/1.6.1/grpc-grpclb-1.6.1.jar
Download https://repo1.maven.org/maven2/io/grpc/grpc-netty/1.6.1/grpc-netty-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-netty/1.6.1/grpc-netty-1.6.1.jar-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-okhttp/1.6.1/grpc-okhttp-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-okhttp/1.6.1/grpc-okhttp-1.6.1.jar
Download https://repo1.maven.org/maven2/io/grpc/grpc-protobuf/1.6.1/grpc-protobuf-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-protobuf/1.6.1/grpc-protobuf-1.6.1.jar
Download https://repo1.maven.org/maven2/io/grpc/grpc-protobuf-lite/1.6.1/grpc-protobuf-lite-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-protobuf-lite/1.6.1/grpc-protobuf-lite-1.6.1.jar
Download https://repo1.maven.org/maven2/io/grpc/grpc-protobuf-nano/1.6.1/grpc-protobuf-nano-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-protobuf-nano/1.6.1/grpc-protobuf-nano-1.6.1.jar
Download https://repo1.maven.org/maven2/io/grpc/grpc-stub/1.6.1/grpc-stub-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-stub/1.6.1/grpc-stub-1.6.1.jar
Download https://repo1.maven.org/maven2/io/grpc/grpc-testing/1.6.1/grpc-testing-1.6.1.pom
Download https://repo1.maven.org/maven2/io/grpc/grpc-testing/1.6.1/grpc-testing-1.6.1.jar
Download https://repo1.maven.org/maven2/com/google/protobuf/protobuf-gradle-plugin/0.8.1/protobuf-gradle-plugin-0.8.1.pom
Download https://repo1.maven.org/maven2/com/google/protobuf/protobuf-gradle-plugin/0.8.1/protobuf-gradle-plugin-0.8.1.jar
Download https://jcenter.bintray.com/com/google/cloud/tools/appengine-gradle-plugin/1.3.3/appengine-gradle-plugin-1.3.3.pom
Download https://jcenter.bintray.com/com/google/cloud/tools/appengine-plugins-core/0.3.2/appengine-plugins-core-0.3.2.pom
Download https://jcenter.bintray.com/com/google/cloud/tools/appengine-gradle-plugin/1.3.3/appengine-gradle-plugin-1.3.3.jar
Download https://jcenter.bintray.com/com/google/cloud/tools/appengine-plugins-core/0.3.2/appengine-plugins-core-0.3.2.jar
Download https://repo1.maven.org/maven2/com/google/errorprone/error_prone_core/2.0.21/error_prone_core-2.0.21.pom
Download https://repo1.maven.org/maven2/com/google/errorprone/error_prone_parent/2.0.21/error_prone_parent-2.0.21.pom
Download https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotation/2.0.21/error_prone_annotation-2.0.21.pom
Download https://repo1.maven.org/maven2/com/google/errorprone/javac/9-dev-r4023-2/javac-9-dev-r4023-2.pom
Download https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.0.21/error_prone_annotations-2.0.21.pom
Download https://repo1.maven.org/maven2/com/google/errorprone/error_prone_check_api/2.0.21/error_prone_check_api-2.0.21.pom
Download https://repo1.maven.org/maven2/com/github/kevinstern/software-and-algorithms/1.0/software-and-algorithms-1.0.pom
Download https://repo1.maven.org/maven2/com/google/errorprone/error_prone_check_api/2.0.21/error_prone_check_api-2.0.21.jar
Download https://repo1.maven.org/maven2/com/google/errorprone/error_prone_core/2.0.21/error_prone_core-2.0.21.jar
Download https://repo1.maven.org/maven2/com/google/errorprone/javac/9-dev-r4023-2/javac-9-dev-r4023-2.jar
Download https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.0.21/error_prone_annotations-2.0.21.jar
Download https://repo1.maven.org/maven2/com/github/kevinstern/software-and-algorithms/1.0/software-and-algorithms-1.0.jar
Download https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotation/2.0.21/error_prone_annotation-2.0.21.jar
Download https://repo1.maven.org/maven2/com/google/errorprone/javac/9-dev-r4023-2/javac-9-dev-r4023-2.jar

> Task :grpc-context:compileJava 
The SimpleWorkResult type has been deprecated and is scheduled to be removed in Gradle 5.0. Please use WorkResults.didWork() instead.

Download https://repo1.maven.org/maven2/io/opencensus/opencensus-api/0.8.0/opencensus-api-0.8.0.pom
Download https://repo1.maven.org/maven2/io/opencensus/opencensus-contrib-grpc-metrics/0.8.0/opencensus-contrib-grpc-metrics-0.8.0.pom
Download https://repo1.maven.org/maven2/com/google/instrumentation/instrumentation-api/0.4.3/instrumentation-api-0.4.3.po
Download https://repo1.maven.org/maven2/com/google/instrumentation/instrumentation-api/0.4.3/instrumentation-api-0.4.3.ja
Download https://repo1.maven.org/maven2/io/opencensus/opencensus-contrib-grpc-metrics/0.8.0/opencensus-contrib-grpc-metrics-0.8.0.jar
Download https://repo1.maven.org/maven2/io/opencensus/opencensus-api/0.8.0/opencensus-api-0.8.0.jar

> Task :grpc-core:compileJava 
The SimpleWorkResult type has been deprecated and is scheduled to be removed in Gradle 5.0. Please use WorkResults.didWork() instead.

Download https://repo1.maven.org/maven2/io/netty/netty-codec-http2/maven-metadata.xml
Download https://repo1.maven.org/maven2/io/netty/netty-codec-http2/4.1.16.Final/netty-codec-http2-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-parent/4.1.16.Final/netty-parent-4.1.16.Final.pom
Download https://oss.sonatype.org/content/repositories/snapshots/io/netty/netty-codec-http2/maven-metadata.xml
Download https://repo1.maven.org/maven2/io/netty/netty-handler-proxy/4.1.16.Final/netty-handler-proxy-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-codec-http/4.1.16.Final/netty-codec-http-4.1.16.Final.pom.pom
Download https://repo1.maven.org/maven2/io/netty/netty-handler/4.1.16.Final/netty-handler-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-codec-socks/4.1.16.Final/netty-codec-socks-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-transport/4.1.16.Final/netty-transport-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-codec/4.1.16.Final/netty-codec-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-buffer/4.1.16.Final/netty-buffer-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-resolver/4.1.16.Final/netty-resolver-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-common/4.1.16.Final/netty-common-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-codec-http2/4.1.16.Final/netty-codec-http2-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/io/netty/netty-handler-proxy/4.1.16.Final/netty-handler-proxy-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/io/netty/netty-handler/4.1.16.Final/netty-handler-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/io/netty/netty-codec-socks/4.1.16.Final/netty-codec-socks-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/io/netty/netty-codec/4.1.16.Final/netty-codec-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/io/netty/netty-buffer/4.1.16.Final/netty-buffer-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/io/netty/netty-transport/4.1.16.Final/netty-transport-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/io/netty/netty-resolver/4.1.16.Final/netty-resolver-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/io/netty/netty-common/4.1.16.Final/netty-common-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/io/netty/netty-codec-http/4.1.16.Final/netty-codec-http-4.1.16.Final.jar
Download https://repo1.maven.org/maven2/com/google/protobuf/protobuf-java-util/3.4.0/protobuf-java-util-3.4.0.pom
Download https://repo1.maven.org/maven2/com/google/protobuf/protobuf-java/3.4.0/protobuf-java-3.4.0.pom
Download https://repo1.maven.org/maven2/com/google/protobuf/protobuf-parent/3.4.0/protobuf-parent-3.4.0.pom
Download https://repo1.maven.org/maven2/com/google/protobuf/protobuf-java-util/3.4.0/protobuf-java-util-3.4.0.jar
Download https://repo1.maven.org/maven2/com/google/protobuf/protobuf-java/3.4.0/protobuf-java-3.4.0.jar

> Task :grpc-protobuf-nano:compileJava 
The SimpleWorkResult type has been deprecated and is scheduled to be removed in Gradle 5.0. Please use WorkResults.didWork() instead.

Download https://repo1.maven.org/maven2/com/google/auth/google-auth-library-oauth2-http/0.7.0/google-auth-library-oauth2-http-0.7.0.pom
Download https://repo1.maven.org/maven2/com/google/auth/google-auth-library-parent/0.7.0/google-auth-library-parent-0.7.0.pom
Download https://repo1.maven.org/maven2/com/google/auth/google-auth-library-credentials/0.7.0/google-auth-library-credentials-0.7.0.pom
Download https://repo1.maven.org/maven2/com/google/auth/google-auth-library-credentials/0.7.0/google-auth-library-credentials-0.7.0.jar
Download https://repo1.maven.org/maven2/com/google/auth/google-auth-library-oauth2-http/0.7.0/google-auth-library-oauth2-http-0.7.0.jar

> Task :grpc-auth:compileTestJava 
The SimpleWorkResult type has been deprecated and is scheduled to be removed in Gradle 5.0. Please use WorkResults.didWork() instead.

Download https://repo1.maven.org/maven2/com/puppycrawl/tools/checkstyle/6.17/checkstyle-6.17.pom
Download https://repo1.maven.org/maven2/org/antlr/antlr4-runtime/4.5.2-1/antlr4-runtime-4.5.2-1.pom
Download https://repo1.maven.org/maven2/org/antlr/antlr4-master/4.5.2-1/antlr4-master-4.5.2-1.pom
Download https://repo1.maven.org/maven2/com/puppycrawl/tools/checkstyle/6.17/checkstyle-6.17.jar
Download https://repo1.maven.org/maven2/org/antlr/antlr4-runtime/4.5.2-1/antlr4-runtime-4.5.2-1.jar
Download https://repo1.maven.org/maven2/org/jacoco/org.jacoco.agent/0.7.9/org.jacoco.agent-0.7.9.jar
Download https://repo1.maven.org/maven2/io/netty/netty-tcnative-boringssl-static/2.0.6.Final/netty-tcnative-boringssl-static-2.0.6.Final.pom
Download https://repo1.maven.org/maven2/org/hdrhistogram/HdrHistogram/2.1.10/HdrHistogram-2.1.10.pom
Download https://repo1.maven.org/maven2/io/netty/netty-tcnative-parent/2.0.6.Final/netty-tcnative-parent-2.0.6.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-transport-native-epoll/4.1.16.Final/netty-transport-native-epoll-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-transport-native-unix-common/4.1.16.Final/netty-transport-native-unix-common-4.1.16.Final.pom
Download https://repo1.maven.org/maven2/io/netty/netty-tcnative-boringssl-static/2.0.6.Final/netty-tcnative-boringssl-static-2.0.6.Final.jar
Download https://repo1.maven.org/maven2/org/hdrhistogram/HdrHistogram/2.1.10/HdrHistogram-2.1.10.jar
Download https://repo1.maven.org/maven2/io/netty/netty-transport-native-epoll/4.1.16.Final/netty-transport-native-epoll-4.1.16.Final-linux-x86_64.jar
Download https://repo1.maven.org/maven2/io/netty/netty-transport-native-unix-common/4.1.16.Final/netty-transport-native-unix-common-4.1.16.Final.jar

> Task :grpc-testing:compileJava 
The SimpleWorkResult type has been deprecated and is scheduled to be removed in Gradle 5.0. Please use WorkResults.didWork() instead.

Download https://repo1.maven.org/maven2/org/openjdk/jmh/jmh-generator-bytecode/1.19/jmh-generator-bytecode-1.19.pom
Download https://repo1.maven.org/maven2/org/openjdk/jmh/jmh-generator-asm/1.19/jmh-generator-asm-1.19.pom
Download https://repo1.maven.org/maven2/org/openjdk/jmh/jmh-generator-reflection/1.19/jmh-generator-reflection-1.19.pom
Download https://repo1.maven.org/maven2/org/openjdk/jmh/jmh-generator-asm/1.19/jmh-generator-asm-1.19.jar
Download https://repo1.maven.org/maven2/org/openjdk/jmh/jmh-generator-bytecode/1.19/jmh-generator-bytecode-1.19.jar
Download https://repo1.maven.org/maven2/org/openjdk/jmh/jmh-generator-reflection/1.19/jmh-generator-reflection-1.19.jar
Download https://repo1.maven.org/maven2/com/google/appengine/appengine/maven-metadata.xml
Download https://repo1.maven.org/maven2/com/google/appengine/appengine/1.9.59/appengine-1.9.59.pom
Download https://repo1.maven.org/maven2/org/sonatype/oss/oss-parent/4/oss-parent-4.pom
Download https://jcenter.bintray.com/com/google/appengine/appengine/maven-metadata.xml
Download https://repo1.maven.org/maven2/com/google/api/api-common/1.0.0-rc1/api-common-1.0.0-rc1.pom
Download https://repo1.maven.org/maven2/io/opencensus/opencensus-impl/0.8.0/opencensus-impl-0.8.0.pom
Download https://repo1.maven.org/maven2/com/google/auto/value/auto-value/1.1/auto-value-1.1.pom
Download https://jcenter.bintray.com/io/netty/netty-codec-http2/maven-metadata.xml
Download https://repo1.maven.org/maven2/com/lmax/disruptor/3.3.6/disruptor-3.3.6.pom
Download https://repo1.maven.org/maven2/io/opencensus/opencensus-impl-core/0.8.0/opencensus-impl-core-0.8.0.pom
Download https://repo1.maven.org/maven2/com/google/api/api-common/1.0.0-rc1/api-common-1.0.0-rc1.jar
Download https://repo1.maven.org/maven2/io/opencensus/opencensus-impl-core/0.8.0/opencensus-impl-core-0.8.0.jar
Download https://repo1.maven.org/maven2/com/google/auto/value/auto-value/1.1/auto-value-1.1.jar
Download https://repo1.maven.org/maven2/com/lmax/disruptor/3.3.6/disruptor-3.3.6.jar
Download https://repo1.maven.org/maven2/io/opencensus/opencensus-impl/0.8.0/opencensus-impl-0.8.0.jareClasspath

> Task :grpc-interop-testing:test 
[jetty-alpn-agent] Using: alpn-boot-8.1.7.v20160121.jar

Download https://repo1.maven.org/maven2/org/codehaus/mojo/signature/java17/1.0/java17-1.0.pom
Download https://repo1.maven.org/maven2/org/codehaus/mojo/signature/signatures-parent/1.2/signatures-parent-1.2.pom
Download https://repo1.maven.org/maven2/org/codehaus/mojo/mojo-parent/31/mojo-parent-31.pom

> Task :grpc-okhttp:test 

io.grpc.okhttp.OkHttpTransportTest > earlyServerClose_serverFailure_withClientCancelOnListenerClosed FAILED
    java.lang.AssertionError: Timed out waiting for server stream

155 tests completed, 1 failed, 1 skipped


FAILURE: Build failed with an exception.

* What went wrong:
Execution failed for task ':grpc-okhttp:test'.
> There were failing tests. See the report at: file:///Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/okhttp/build/reports/tests/test/index.html

* Try:
Run with --stacktrace option to get the stack trace. Run with --info or --debug option to get more log output.

* Get more help at https://help.gradle.org

BUILD FAILED in 37s
157 actionable tasks: 1 executed, 156 up-to-date
```

Rollback
```
[vagrant@localhost grpc-java]$ git checkout 1f7c9d571620e4d61dd89c8c85256122e81fd25f
正在检出文件: 100% (801/801), 完成.
Note: checking out '1f7c9d571620e4d61dd89c8c85256122e81fd25f'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by performing another checkout.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -b with the checkout command again. Example:

  git checkout -b <new-branch-name>

HEAD 目前位于 1f7c9d5... core: document ServerBuilder return values and method history
[vagrant@localhost grpc-java]$ ./gradlew build
Starting a Gradle Daemon (subsequent builds will be faster)
*** Skipping the build of codegen and compilation of proto files because skipCodegen=true
Download https://plugins.gradle.org/m2/net/ltgt/gradle/gradle-errorprone-plugin/0.0.9/gradle-errorprone-plugin-0.0.9.pom
Download https://plugins.gradle.org/m2/net/ltgt/gradle/gradle-errorprone-plugin/0.0.9/gradle-errorprone-plugin-0.0.9.jar
Download https://plugins.gradle.org/m2/gradle/plugin/me/champeau/gradle/jmh-gradle-plugin/0.3.0/jmh-gradle-plugin-0.3.0.pom
Download https://plugins.gradle.org/m2/gradle/plugin/me/champeau/gradle/jmh-gradle-plugin/0.3.0/jmh-gradle-plugin-0.3.0.jar
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/core/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/core/build/resources/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/auth/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/protobuf/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/benchmarks/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/grpclb/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/grpclb/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/stub/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/services/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/services/build/resources/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/protobuf/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/protobuf-lite/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/testing-proto/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/testing-proto/build/resources/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/testing/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/testing/build/resources/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/interop-testing/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/interop-testing/build/resources/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/context/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/protobuf-lite/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/protobuf-lite/build/resources/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/context/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/core/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/core/build/resources/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/interop-testing/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/protobuf-nano/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/netty/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/benchmarks/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/thrift/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/netty/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/netty/build/resources/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/protobuf-nano/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/protobuf-nano/build/resources/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/stub/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/testing/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/services/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/services/build/resources/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/okhttp/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/okhttp/build/resources/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/benchmarks/build/classes/jmh'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/thrift/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/okhttp/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/auth/build/classes/main'
:grpc-context:compileJava
:grpc-context:processResources NO-SOURCE
:grpc-context:classes
:grpc-context:jar
:grpc-core:compileJava
:grpc-core:processResources
:grpc-core:classes
:grpc-core:jar
:grpc-auth:compileJava
The message received from the daemon indicates that the daemon has disappeared.
Build request sent: Build{id=1228d849-a66f-40a1-aedc-44e43d6f70f8.1, currentDir=/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java}
Attempting to read last messages from the daemon log...
Daemon pid: 22278
  log file: /home/vagrant/.gradle/daemon/3.4.1/daemon-22278.out.log
----- Last  20 lines from daemon log file - daemon-22278.out.log -----
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/protobuf-nano/build/resources/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/stub/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/testing/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/services/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/services/build/resources/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/okhttp/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/okhttp/build/resources/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/benchmarks/build/classes/jmh'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/thrift/build/classes/main'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/okhttp/build/classes/test'
Cleaned up directory '/Users/fanhongling/Downloads/workspace/src/github.com/grpc/grpc-java/auth/build/classes/main'
:grpc-context:compileJava
:grpc-context:processResources NO-SOURCE
:grpc-context:classes
:grpc-context:jar
:grpc-core:compileJava
:grpc-core:processResources
:grpc-core:classes
:grpc-core:jar
:grpc-auth:compileJava
:grpc-auth:processResources NO-SOURCE
:grpc-auth:classes
:grpc-auth:jar
:grpc-netty:compileJava
:grpc-netty:processResources
:grpc-netty:classes
:grpc-netty:jar
:grpc-okhttp:compileJava
:grpc-okhttp:processResources
:grpc-okhttp:classes
:grpc-okhttp:jar
:grpc-protobuf-lite:extractIncludeProto
:grpc-protobuf-lite:extractProto
:grpc-protobuf-lite:generateProto NO-SOURCE
:grpc-protobuf-lite:compileJava
:grpc-protobuf-lite:processResources NO-SOURCE
:grpc-protobuf-lite:classes
:grpc-protobuf-lite:jar
:grpc-protobuf:compileJava
:grpc-protobuf:processResources NO-SOURCE
:grpc-protobuf:classes
:grpc-protobuf:jar
:grpc-protobuf-nano:compileJava
:grpc-protobuf-nano:processResources NO-SOURCE
:grpc-protobuf-nano:classes
:grpc-protobuf-nano:jar
:grpc-stub:compileJava
:grpc-stub:processResources NO-SOURCE
:grpc-stub:classes
:grpc-stub:jar
:grpc-all:compileJava NO-SOURCE
:grpc-all:processResources NO-SOURCE
:grpc-all:classes UP-TO-DATE
:grpc-all:jar
:grpc-context:javadoc
:grpc-core:javadoc
:grpc-auth:javadoc
:grpc-netty:javadoc
:grpc-okhttp:javadoc
:grpc-protobuf-lite:javadoc
:grpc-protobuf:javadoc
:grpc-protobuf-nano:javadoc
:grpc-stub:javadoc
:grpc-all:javadoc
:grpc-all:javadocJar
:grpc-all:sourcesJar
:grpc-all:signArchives SKIPPED
:grpc-all:assemble
:grpc-all:animalsnifferMain NO-SOURCE
:grpc-all:compileTestJava NO-SOURCE
:grpc-all:processTestResources NO-SOURCE
:grpc-all:testClasses UP-TO-DATE
:grpc-all:animalsnifferTest NO-SOURCE
:grpc-all:checkstyleMain NO-SOURCE
:grpc-all:checkstyleTest NO-SOURCE
:grpc-all:test NO-SOURCE
:grpc-all:check UP-TO-DATE
:grpc-all:build
:grpc-auth:javadocJar
:grpc-auth:sourcesJar
:grpc-auth:signArchives SKIPPED
:grpc-auth:assemble
:grpc-auth:animalsnifferMain
:grpc-auth:compileTestJava
:grpc-auth:processTestResources NO-SOURCE
:grpc-auth:testClasses
:grpc-auth:animalsnifferTest
:grpc-auth:checkstyleMain
:grpc-auth:checkstyleTest
:grpc-auth:test
:grpc-auth:check
:grpc-auth:build
:grpc-testing:compileJava
:grpc-testing:processResources
:grpc-testing:classes
:grpc-testing:jar
:grpc-benchmarks:compileJava
Download https://repo1.maven.org/maven2/io/netty/netty-transport-native-epoll/4.1.8.Final/netty-transport-native-epoll-4.1.8.Final-linux-x86_64.jar
:grpc-benchmarks:processResources NO-SOURCE
:grpc-benchmarks:classes
:grpc-benchmarks:jar
:grpc-benchmarks:benchmark_worker
:grpc-benchmarks:openloop_client
:grpc-benchmarks:qps_client
:grpc-benchmarks:qps_server
:grpc-benchmarks:startScripts SKIPPED
:grpc-benchmarks:distTar
:grpc-benchmarks:distZip
:grpc-testing:javadoc
:grpc-benchmarks:javadoc
:grpc-benchmarks:javadocJar
:grpc-benchmarks:sourcesJar
:grpc-benchmarks:signArchives SKIPPED
:grpc-benchmarks:assemble
:grpc-benchmarks:compileTestJava
:grpc-benchmarks:processTestResources NO-SOURCE
:grpc-benchmarks:testClasses
:grpc-benchmarks:compileJmhJava
:grpc-benchmarks:processJmhResources NO-SOURCE
:grpc-benchmarks:jmhClasses
:grpc-benchmarks:animalsnifferJmh
:grpc-benchmarks:animalsnifferMain
:grpc-benchmarks:animalsnifferTest
:grpc-benchmarks:checkstyleJmh
:grpc-benchmarks:checkstyleMain
:grpc-benchmarks:checkstyleTest
:grpc-benchmarks:test
:grpc-benchmarks:check
:grpc-benchmarks:build
:grpc-context:javadocJar
:grpc-context:sourcesJar
:grpc-context:signArchives SKIPPED
:grpc-context:assemble
:grpc-context:animalsnifferMain
:grpc-context:compileTestJava
:grpc-context:processTestResources NO-SOURCE
:grpc-context:testClasses
:grpc-context:animalsnifferTest
:grpc-context:checkstyleMain
:grpc-context:checkstyleTest
:grpc-context:test
:grpc-context:check
:grpc-context:build
:grpc-core:javadocJar
:grpc-core:sourcesJar
:grpc-core:signArchives SKIPPED
:grpc-core:assemble
:grpc-core:animalsnifferMain
:grpc-core:compileTestJava
:grpc-core:processTestResources
:grpc-core:testClasses
:grpc-core:animalsnifferTest
:grpc-core:checkstyleMain
:grpc-core:checkstyleTest
:grpc-core:test
:grpc-core:check
:grpc-core:build
:grpc-grpclb:compileJava
:grpc-grpclb:processResources NO-SOURCE
:grpc-grpclb:classes
:grpc-grpclb:jar
:grpc-grpclb:javadoc
:grpc-grpclb:javadocJar
:grpc-grpclb:sourcesJar
:grpc-grpclb:signArchives SKIPPED
:grpc-grpclb:assemble
:grpc-grpclb:animalsnifferMain
:grpc-grpclb:compileTestJava
:grpc-grpclb:processTestResources NO-SOURCE
:grpc-grpclb:testClasses
:grpc-grpclb:animalsnifferTest
:grpc-grpclb:checkstyleMain
:grpc-grpclb:checkstyleTest
:grpc-grpclb:test
:grpc-grpclb:check
:grpc-grpclb:build
:grpc-testing-proto:compileJava
:grpc-testing-proto:processResources NO-SOURCE
:grpc-testing-proto:classes
:grpc-testing-proto:jar
:grpc-interop-testing:compileJava
:grpc-interop-testing:processResources
:grpc-interop-testing:classes
:grpc-interop-testing:jar
:grpc-interop-testing:http2_client
:grpc-interop-testing:reconnect_test_client
:grpc-interop-testing:startScripts SKIPPED
:grpc-interop-testing:stresstest_client
:grpc-interop-testing:test_client
:grpc-interop-testing:test_server
:grpc-interop-testing:distTar
:grpc-interop-testing:distZip
:grpc-testing-proto:javadoc
:grpc-interop-testing:javadoc
:grpc-interop-testing:javadocJar
:grpc-interop-testing:sourcesJar
:grpc-interop-testing:signArchives SKIPPED
:grpc-interop-testing:assemble
:grpc-interop-testing:animalsnifferMain
:grpc-interop-testing:compileTestJava
:grpc-interop-testing:processTestResources NO-SOURCE
:grpc-interop-testing:testClasses
:grpc-interop-testing:animalsnifferTest
:grpc-interop-testing:checkstyleMain
:grpc-interop-testing:checkstyleTest
:grpc-interop-testing:test
[jetty-alpn-agent] Using: alpn-boot-8.1.7.v20160121.jar
:grpc-interop-testing:check
:grpc-interop-testing:build
:grpc-netty:javadocJar
:grpc-netty:sourcesJar
:grpc-netty:signArchives SKIPPED
:grpc-netty:assemble
:grpc-netty:animalsnifferMain
:grpc-netty:compileTestJava
:grpc-netty:processTestResources NO-SOURCE
:grpc-netty:testClasses
:grpc-netty:animalsnifferTest
:grpc-netty:checkstyleMain
:grpc-netty:checkstyleTest
:grpc-netty:test
:grpc-netty:check
:grpc-netty:build
:grpc-okhttp:javadocJar
:grpc-okhttp:sourcesJar
:grpc-okhttp:signArchives SKIPPED
:grpc-okhttp:assemble
:grpc-okhttp:animalsnifferMain
:grpc-okhttp:compileTestJava
:grpc-okhttp:processTestResources NO-SOURCE
:grpc-okhttp:testClasses
:grpc-okhttp:animalsnifferTest
:grpc-okhttp:checkstyleMain
:grpc-okhttp:checkstyleTest
:grpc-okhttp:test
:grpc-okhttp:check
:grpc-okhttp:build
:grpc-protobuf:javadocJar
:grpc-protobuf:sourcesJar
:grpc-protobuf:signArchives SKIPPED
:grpc-protobuf:assemble
:grpc-protobuf:animalsnifferMain
:grpc-protobuf:compileTestJava
:grpc-protobuf:processTestResources NO-SOURCE
:grpc-protobuf:testClasses
:grpc-protobuf:animalsnifferTest
:grpc-protobuf:checkstyleMain
:grpc-protobuf:checkstyleTest
:grpc-protobuf:test
:grpc-protobuf:check
:grpc-protobuf:build
:grpc-protobuf-lite:javadocJar
:grpc-protobuf-lite:sourcesJar
:grpc-protobuf-lite:signArchives SKIPPED
:grpc-protobuf-lite:assemble
:grpc-protobuf-lite:animalsnifferMain
:grpc-protobuf-lite:extractIncludeTestProto
:grpc-protobuf-lite:extractTestProto
:grpc-protobuf-lite:generateTestProto
Download https://repo1.maven.org/maven2/com/google/protobuf/protoc-gen-javalite/3.0.0/protoc-gen-javalite-3.0.0.pom
Download https://repo1.maven.org/maven2/com/google/protobuf/protoc-gen-javalite/3.0.0/protoc-gen-javalite-3.0.0-linux-x86_64.exe
:grpc-protobuf-lite:compileTestJava
注: 某些输入文件使用了未经检查或不安全的操作。
注: 有关详细信息, 请使用 -Xlint:unchecked 重新编译。
:grpc-protobuf-lite:processTestResources
:grpc-protobuf-lite:testClasses
:grpc-protobuf-lite:animalsnifferTest
:grpc-protobuf-lite:checkstyleMain
:grpc-protobuf-lite:checkstyleTest
:grpc-protobuf-lite:test
:grpc-protobuf-lite:check
:grpc-protobuf-lite:build
:grpc-protobuf-nano:javadocJar
:grpc-protobuf-nano:sourcesJar
:grpc-protobuf-nano:signArchives SKIPPED
:grpc-protobuf-nano:assemble
:grpc-protobuf-nano:animalsnifferMain
:grpc-protobuf-nano:compileTestJava
:grpc-protobuf-nano:processTestResources NO-SOURCE
:grpc-protobuf-nano:testClasses
:grpc-protobuf-nano:animalsnifferTest
:grpc-protobuf-nano:checkstyleMain
:grpc-protobuf-nano:checkstyleTest
:grpc-protobuf-nano:test
:grpc-protobuf-nano:check
:grpc-protobuf-nano:build
:grpc-services:compileJava
:grpc-services:processResources NO-SOURCE
:grpc-services:classes
:grpc-services:jar
:grpc-services:javadoc
:grpc-services:javadocJar
:grpc-services:sourcesJar
:grpc-services:signArchives SKIPPED
:grpc-services:assemble
:grpc-services:animalsnifferMain
:grpc-services:compileTestJava
:grpc-services:processTestResources NO-SOURCE
:grpc-services:testClasses
:grpc-services:animalsnifferTest
:grpc-services:checkstyleMain
:grpc-services:checkstyleTest
:grpc-services:test
:grpc-services:check
:grpc-services:build
:grpc-stub:javadocJar
:grpc-stub:sourcesJar
:grpc-stub:signArchives SKIPPED
:grpc-stub:assemble
:grpc-stub:animalsnifferMain
:grpc-stub:compileTestJava
:grpc-stub:processTestResources NO-SOURCE
:grpc-stub:testClasses
:grpc-stub:animalsnifferTest
:grpc-stub:checkstyleMain
:grpc-stub:checkstyleTest
:grpc-stub:test
:grpc-stub:check
:grpc-stub:build
:grpc-testing:javadocJar
:grpc-testing:sourcesJar
:grpc-testing:signArchives SKIPPED
:grpc-testing:assemble
:grpc-testing:animalsnifferMain
:grpc-testing:compileTestJava
:grpc-testing:processTestResources NO-SOURCE
:grpc-testing:testClasses
:grpc-testing:animalsnifferTest
:grpc-testing:checkstyleMain
:grpc-testing:checkstyleTest
:grpc-testing:test
:grpc-testing:check
:grpc-testing:build
:grpc-testing-proto:javadocJar
:grpc-testing-proto:sourcesJar
:grpc-testing-proto:signArchives SKIPPED
:grpc-testing-proto:assemble
:grpc-testing-proto:animalsnifferMain
:grpc-testing-proto:compileTestJava NO-SOURCE
:grpc-testing-proto:processTestResources NO-SOURCE
:grpc-testing-proto:testClasses UP-TO-DATE
:grpc-testing-proto:animalsnifferTest NO-SOURCE
:grpc-testing-proto:checkstyleMain NO-SOURCE
:grpc-testing-proto:checkstyleTest NO-SOURCE
:grpc-testing-proto:test NO-SOURCE
:grpc-testing-proto:check
:grpc-testing-proto:build
:grpc-thrift:compileJava
:grpc-thrift:processResources NO-SOURCE
:grpc-thrift:classes
:grpc-thrift:jar
:grpc-thrift:javadoc
:grpc-thrift:javadocJar
:grpc-thrift:sourcesJar
:grpc-thrift:signArchives SKIPPED
:grpc-thrift:assemble
:grpc-thrift:animalsnifferMain
:grpc-thrift:compileTestJava
:grpc-thrift:processTestResources NO-SOURCE
:grpc-thrift:testClasses
:grpc-thrift:animalsnifferTest
:grpc-thrift:checkstyleMain
:grpc-thrift:checkstyleTest
:grpc-thrift:test
:grpc-thrift:check
:grpc-thrift:build

BUILD SUCCESSFUL

Total time: 19 mins 5.692 secs
[vagrant@localhost grpc-java]$ ./gradlew install
*** Skipping the build of codegen and compilation of proto files because skipCodegen=true
:grpc-context:compileJava UP-TO-DATE
:grpc-context:processResources NO-SOURCE
:grpc-context:classes UP-TO-DATE
:grpc-context:jar UP-TO-DATE
:grpc-core:compileJava UP-TO-DATE
:grpc-core:processResources UP-TO-DATE
:grpc-core:classes UP-TO-DATE
:grpc-core:jar UP-TO-DATE
:grpc-auth:compileJava UP-TO-DATE
:grpc-auth:processResources NO-SOURCE
:grpc-auth:classes UP-TO-DATE
:grpc-auth:jar UP-TO-DATE
:grpc-netty:compileJava UP-TO-DATE
:grpc-netty:processResources UP-TO-DATE
:grpc-netty:classes UP-TO-DATE
:grpc-netty:jar UP-TO-DATE
:grpc-okhttp:compileJava UP-TO-DATE
:grpc-okhttp:processResources UP-TO-DATE
:grpc-okhttp:classes UP-TO-DATE
:grpc-okhttp:jar UP-TO-DATE
:grpc-protobuf-lite:extractIncludeProto UP-TO-DATE
:grpc-protobuf-lite:extractProto UP-TO-DATE
:grpc-protobuf-lite:generateProto NO-SOURCE
:grpc-protobuf-lite:compileJava UP-TO-DATE
:grpc-protobuf-lite:processResources NO-SOURCE
:grpc-protobuf-lite:classes UP-TO-DATE
:grpc-protobuf-lite:jar UP-TO-DATE
:grpc-protobuf:compileJava UP-TO-DATE
:grpc-protobuf:processResources NO-SOURCE
:grpc-protobuf:classes UP-TO-DATE
:grpc-protobuf:jar UP-TO-DATE
:grpc-protobuf-nano:compileJava UP-TO-DATE
:grpc-protobuf-nano:processResources NO-SOURCE
:grpc-protobuf-nano:classes UP-TO-DATE
:grpc-protobuf-nano:jar UP-TO-DATE
:grpc-stub:compileJava UP-TO-DATE
:grpc-stub:processResources NO-SOURCE
:grpc-stub:classes UP-TO-DATE
:grpc-stub:jar UP-TO-DATE
:grpc-all:compileJava NO-SOURCE
:grpc-all:processResources NO-SOURCE
:grpc-all:classes UP-TO-DATE
:grpc-all:jar UP-TO-DATE
:grpc-context:javadoc UP-TO-DATE
:grpc-core:javadoc UP-TO-DATE
:grpc-auth:javadoc UP-TO-DATE
:grpc-netty:javadoc UP-TO-DATE
:grpc-okhttp:javadoc UP-TO-DATE
:grpc-protobuf-lite:javadoc UP-TO-DATE
:grpc-protobuf:javadoc UP-TO-DATE
:grpc-protobuf-nano:javadoc UP-TO-DATE
:grpc-stub:javadoc UP-TO-DATE
:grpc-all:javadoc UP-TO-DATE
:grpc-all:javadocJar UP-TO-DATE
:grpc-all:sourcesJar UP-TO-DATE
:grpc-all:signArchives SKIPPED
:grpc-all:install
:grpc-auth:javadocJar UP-TO-DATE
:grpc-auth:sourcesJar UP-TO-DATE
:grpc-auth:signArchives SKIPPED
:grpc-auth:install
:grpc-testing:compileJava UP-TO-DATE
:grpc-testing:processResources UP-TO-DATE
:grpc-testing:classes UP-TO-DATE
:grpc-testing:jar UP-TO-DATE
:grpc-benchmarks:compileJava UP-TO-DATE
:grpc-benchmarks:processResources NO-SOURCE
:grpc-benchmarks:classes UP-TO-DATE
:grpc-benchmarks:jar UP-TO-DATE
:grpc-benchmarks:benchmark_worker UP-TO-DATE
:grpc-benchmarks:openloop_client UP-TO-DATE
:grpc-benchmarks:qps_client UP-TO-DATE
:grpc-benchmarks:qps_server UP-TO-DATE
:grpc-benchmarks:startScripts SKIPPED
:grpc-benchmarks:distTar
:grpc-benchmarks:distZip
:grpc-testing:javadoc UP-TO-DATE
:grpc-benchmarks:javadoc UP-TO-DATE
:grpc-benchmarks:javadocJar UP-TO-DATE
:grpc-benchmarks:sourcesJar UP-TO-DATE
:grpc-benchmarks:signArchives SKIPPED
:grpc-benchmarks:install
:grpc-context:javadocJar UP-TO-DATE
:grpc-context:sourcesJar UP-TO-DATE
:grpc-context:signArchives SKIPPED
:grpc-context:install
:grpc-core:javadocJar UP-TO-DATE
:grpc-core:sourcesJar UP-TO-DATE
:grpc-core:signArchives SKIPPED
:grpc-core:install
:grpc-grpclb:compileJava UP-TO-DATE
:grpc-grpclb:processResources NO-SOURCE
:grpc-grpclb:classes UP-TO-DATE
:grpc-grpclb:jar UP-TO-DATE
:grpc-grpclb:javadoc UP-TO-DATE
:grpc-grpclb:javadocJar UP-TO-DATE
:grpc-grpclb:sourcesJar UP-TO-DATE
:grpc-grpclb:signArchives SKIPPED
:grpc-grpclb:install
:grpc-testing-proto:compileJava UP-TO-DATE
:grpc-testing-proto:processResources NO-SOURCE
:grpc-testing-proto:classes UP-TO-DATE
:grpc-testing-proto:jar UP-TO-DATE
:grpc-interop-testing:compileJava UP-TO-DATE
:grpc-interop-testing:processResources UP-TO-DATE
:grpc-interop-testing:classes UP-TO-DATE
:grpc-interop-testing:jar UP-TO-DATE
:grpc-interop-testing:http2_client UP-TO-DATE
:grpc-interop-testing:reconnect_test_client UP-TO-DATE
:grpc-interop-testing:startScripts SKIPPED
:grpc-interop-testing:stresstest_client UP-TO-DATE
:grpc-interop-testing:test_client UP-TO-DATE
:grpc-interop-testing:test_server UP-TO-DATE
:grpc-interop-testing:distTar
:grpc-interop-testing:distZip
:grpc-testing-proto:javadoc UP-TO-DATE
:grpc-interop-testing:javadoc UP-TO-DATE
:grpc-interop-testing:javadocJar UP-TO-DATE
:grpc-interop-testing:sourcesJar UP-TO-DATE
:grpc-interop-testing:signArchives SKIPPED
:grpc-interop-testing:install
:grpc-netty:javadocJar UP-TO-DATE
:grpc-netty:sourcesJar UP-TO-DATE
:grpc-netty:signArchives SKIPPED
:grpc-netty:install
:grpc-okhttp:javadocJar UP-TO-DATE
:grpc-okhttp:sourcesJar UP-TO-DATE
:grpc-okhttp:signArchives SKIPPED
:grpc-okhttp:install
:grpc-protobuf:javadocJar UP-TO-DATE
:grpc-protobuf:sourcesJar UP-TO-DATE
:grpc-protobuf:signArchives SKIPPED
:grpc-protobuf:install
:grpc-protobuf-lite:javadocJar UP-TO-DATE
:grpc-protobuf-lite:sourcesJar UP-TO-DATE
:grpc-protobuf-lite:signArchives SKIPPED
:grpc-protobuf-lite:install
:grpc-protobuf-nano:javadocJar UP-TO-DATE
:grpc-protobuf-nano:sourcesJar UP-TO-DATE
:grpc-protobuf-nano:signArchives SKIPPED
:grpc-protobuf-nano:install
:grpc-services:compileJava UP-TO-DATE
:grpc-services:processResources NO-SOURCE
:grpc-services:classes UP-TO-DATE
:grpc-services:jar UP-TO-DATE
:grpc-services:javadoc UP-TO-DATE
:grpc-services:javadocJar UP-TO-DATE
:grpc-services:sourcesJar UP-TO-DATE
:grpc-services:signArchives SKIPPED
:grpc-services:install
:grpc-stub:javadocJar UP-TO-DATE
:grpc-stub:sourcesJar UP-TO-DATE
:grpc-stub:signArchives SKIPPED
:grpc-stub:install
:grpc-testing:javadocJar UP-TO-DATE
:grpc-testing:sourcesJar UP-TO-DATE
:grpc-testing:signArchives SKIPPED
:grpc-testing:install
:grpc-testing-proto:javadocJar UP-TO-DATE
:grpc-testing-proto:sourcesJar UP-TO-DATE
:grpc-testing-proto:signArchives SKIPPED
:grpc-testing-proto:install
:grpc-thrift:compileJava UP-TO-DATE
:grpc-thrift:processResources NO-SOURCE
:grpc-thrift:classes UP-TO-DATE
:grpc-thrift:jar UP-TO-DATE
:grpc-thrift:javadoc UP-TO-DATE
:grpc-thrift:javadocJar UP-TO-DATE
:grpc-thrift:sourcesJar UP-TO-DATE
:grpc-thrift:signArchives SKIPPED
:grpc-thrift:install

BUILD SUCCESSFUL

Total time: 32.662 secs
[vagrant@localhost grpc-java]$ ldd compiler/build/exe/java_plugin/protoc-gen-grpc-java 
	linux-vdso.so.1 (0x00007ffdc3adb000)
	libpthread.so.0 => /lib64/libpthread.so.0 (0x00007f58ad343000)
	libm.so.6 => /lib64/libm.so.6 (0x00007f58ad041000)
	libc.so.6 => /lib64/libc.so.6 (0x00007f58acc7f000)
	/lib64/ld-linux-x86-64.so.2 (0x000055f524b46000)
[vagrant@localhost grpc-java]$ ldd compiler/build/artifacts/java_plugin/protoc-gen-grpc-java.exe 
	linux-vdso.so.1 (0x00007ffed6123000)
	libpthread.so.0 => /lib64/libpthread.so.0 (0x00007f0f48be5000)
	libm.so.6 => /lib64/libm.so.6 (0x00007f0f488e3000)
	libc.so.6 => /lib64/libc.so.6 (0x00007f0f48521000)
	/lib64/ld-linux-x86-64.so.2 (0x0000564d63441000)
[vagrant@localhost grpc-java]$ ldd ~/.m2/repository/io/grpc/protoc-gen-grpc-java/1.4.0-SNAPSHOT/protoc-gen-grpc-java-1.4.0-SNAPSHOT-linux-x86_64.exe
	linux-vdso.so.1 (0x00007ffead5e6000)
	libpthread.so.0 => /lib64/libpthread.so.0 (0x00007fd43a7b5000)
	libm.so.6 => /lib64/libm.so.6 (0x00007fd43a4b3000)
	libc.so.6 => /lib64/libc.so.6 (0x00007fd43a0f1000)
	/lib64/ld-linux-x86-64.so.2 (0x000055c7dcf9d000)
```

## May 5

Ubuntu 14.04 VBox

Protobuf
```
[vagrant@localhost bazel]$ protoc --version
libprotoc 3.0.0-beta-3
```

### Get repository

With `git`
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com$ git clone https://github.com/grpc/grpc-java grpc/grpc-java
```

Revision
```
[vagrant@localhost grpc-java]$ git log -1
commit 1f7c9d571620e4d61dd89c8c85256122e81fd25f
Author: Carl Mastrangelo <notcarl@google.com>
Date:   Fri May 5 10:34:10 2017 -0700

    core: document ServerBuilder return values and method history
    
    Also, remove an erroneously added ExperimentalApi annotation.  The
    method was non experimental as of the 1.0.0 release.
```

### Build and Install


with `gradlew`
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com$ cd grpc/grpc-java/
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ ./gradlew build
Downloading https://services.gradle.org/distributions/gradle-3.4.1-bin.zip
...........................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................
Unzipping /home/vagrant/.gradle/wrapper/dists/gradle-3.4.1-bin/71zneekfcxxu7l9p7nr2sc65s/gradle-3.4.1-bin.zip to /home/vagrant/.gradle/wrapper/dists/gradle-3.4.1-bin/71zneekfcxxu7l9p7nr2sc65s
Set executable permissions for: /home/vagrant/.gradle/wrapper/dists/gradle-3.4.1-bin/71zneekfcxxu7l9p7nr2sc65s/gradle-3.4.1/bin/gradle
Starting a Gradle Daemon (subsequent builds will be faster)

BUILD SUCCESSFUL

Total time: 2 mins 39.964 secs
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ ./gradlew install
*** Building codegen requires Protobuf version 3.2.0
*** Please refer to https://github.com/grpc/grpc-java/blob/master/COMPILING.md#how-to-build-code-generation-plugin
:grpc-context:compileJava UP-TO-DATE
:grpc-context:processResources NO-SOURCE
:grpc-context:classes UP-TO-DATE
:grpc-context:jar UP-TO-DATE
:grpc-core:compileJava UP-TO-DATE
:grpc-core:processResources UP-TO-DATE
:grpc-core:classes UP-TO-DATE
:grpc-core:jar UP-TO-DATE
:grpc-auth:compileJava UP-TO-DATE
:grpc-auth:processResources NO-SOURCE
:grpc-auth:classes UP-TO-DATE
:grpc-auth:jar UP-TO-DATE
:grpc-netty:compileJava UP-TO-DATE
:grpc-netty:processResources UP-TO-DATE
:grpc-netty:classes UP-TO-DATE
:grpc-netty:jar UP-TO-DATE
:grpc-okhttp:compileJava UP-TO-DATE
:grpc-okhttp:processResources UP-TO-DATE
:grpc-okhttp:classes UP-TO-DATE
:grpc-okhttp:jar UP-TO-DATE
:grpc-compiler:compileJava_pluginExecutableJava_pluginCpp UP-TO-DATE
:grpc-compiler:linkJava_pluginExecutable UP-TO-DATE
:grpc-compiler:java_pluginExecutable UP-TO-DATE
:grpc-protobuf:deleteGeneratedSourcemain
:grpc-protobuf-lite:extractIncludeProto UP-TO-DATE
:grpc-protobuf-lite:extractProto UP-TO-DATE
:grpc-protobuf-lite:generateProto NO-SOURCE
:grpc-protobuf-lite:compileJava UP-TO-DATE
:grpc-protobuf-lite:processResources NO-SOURCE
:grpc-protobuf-lite:classes UP-TO-DATE
:grpc-protobuf-lite:jar UP-TO-DATE
:grpc-protobuf:extractIncludeProto UP-TO-DATE
:grpc-protobuf:extractProto UP-TO-DATE
:grpc-protobuf:generateProto NO-SOURCE
:grpc-protobuf:compileJava UP-TO-DATE
:grpc-protobuf:processResources NO-SOURCE
:grpc-protobuf:classes UP-TO-DATE
:grpc-protobuf:jar UP-TO-DATE
:grpc-protobuf-nano:deleteGeneratedSourcemain
:grpc-protobuf-nano:extractIncludeProto UP-TO-DATE
:grpc-protobuf-nano:extractProto UP-TO-DATE
:grpc-protobuf-nano:generateProto NO-SOURCE
:grpc-protobuf-nano:compileJava UP-TO-DATE
:grpc-protobuf-nano:processResources NO-SOURCE
:grpc-protobuf-nano:classes UP-TO-DATE
:grpc-protobuf-nano:jar UP-TO-DATE
:grpc-stub:compileJava UP-TO-DATE
:grpc-stub:processResources NO-SOURCE
:grpc-stub:classes UP-TO-DATE
:grpc-stub:jar UP-TO-DATE
:grpc-all:compileJava NO-SOURCE
:grpc-all:processResources NO-SOURCE
:grpc-all:classes UP-TO-DATE
:grpc-all:jar UP-TO-DATE
:grpc-context:javadoc UP-TO-DATE
:grpc-core:javadoc UP-TO-DATE
:grpc-auth:javadoc UP-TO-DATE
:grpc-netty:javadoc UP-TO-DATE
:grpc-okhttp:javadoc UP-TO-DATE
:grpc-protobuf-lite:javadoc UP-TO-DATE
:grpc-protobuf:javadoc UP-TO-DATE
:grpc-protobuf-nano:javadoc UP-TO-DATE
:grpc-stub:javadoc UP-TO-DATE
:grpc-all:javadoc UP-TO-DATE
:grpc-all:javadocJar UP-TO-DATE
:grpc-all:sourcesJar UP-TO-DATE
:grpc-all:signArchives SKIPPED
:grpc-all:install
:grpc-auth:javadocJar UP-TO-DATE
:grpc-auth:sourcesJar UP-TO-DATE
:grpc-auth:signArchives SKIPPED
:grpc-auth:install
:grpc-benchmarks:deleteGeneratedSourcemain
:grpc-testing:compileJava UP-TO-DATE
:grpc-testing:processResources UP-TO-DATE
:grpc-testing:classes UP-TO-DATE
:grpc-testing:jar UP-TO-DATE
:grpc-benchmarks:extractIncludeProto UP-TO-DATE
:grpc-benchmarks:extractProto UP-TO-DATE
:grpc-benchmarks:generateProto
:grpc-benchmarks:compileJava UP-TO-DATE
:grpc-benchmarks:processResources UP-TO-DATE
:grpc-benchmarks:classes UP-TO-DATE
:grpc-benchmarks:jar UP-TO-DATE
:grpc-benchmarks:benchmark_worker UP-TO-DATE
:grpc-benchmarks:openloop_client UP-TO-DATE
:grpc-benchmarks:qps_client UP-TO-DATE
:grpc-benchmarks:qps_server UP-TO-DATE
:grpc-benchmarks:startScripts SKIPPED
:grpc-benchmarks:distTar
:grpc-benchmarks:distZip
:grpc-testing:javadoc UP-TO-DATE
:grpc-benchmarks:javadoc UP-TO-DATE
:grpc-benchmarks:javadocJar UP-TO-DATE
:grpc-benchmarks:sourcesJar UP-TO-DATE
:grpc-benchmarks:signArchives SKIPPED
:grpc-benchmarks:install
:grpc-compiler:buildArtifacts UP-TO-DATE
:grpc-compiler:extractIncludeProto UP-TO-DATE
:grpc-compiler:extractProto UP-TO-DATE
:grpc-compiler:generateProto NO-SOURCE
:grpc-compiler:compileJava NO-SOURCE
:grpc-compiler:processResources NO-SOURCE
:grpc-compiler:classes UP-TO-DATE
:grpc-compiler:jar UP-TO-DATE
:grpc-compiler:javadoc NO-SOURCE
:grpc-compiler:javadocJar UP-TO-DATE
:grpc-compiler:sourcesJar UP-TO-DATE
:grpc-compiler:signArchives SKIPPED
:grpc-compiler:install
:grpc-context:javadocJar UP-TO-DATE
:grpc-context:sourcesJar UP-TO-DATE
:grpc-context:signArchives SKIPPED
:grpc-context:install
:grpc-core:javadocJar UP-TO-DATE
:grpc-core:sourcesJar UP-TO-DATE
:grpc-core:signArchives SKIPPED
:grpc-core:install
:grpc-grpclb:deleteGeneratedSourcemain
:grpc-grpclb:extractIncludeProto UP-TO-DATE
:grpc-grpclb:extractProto UP-TO-DATE
:grpc-grpclb:generateProto
:grpc-grpclb:compileJava UP-TO-DATE
:grpc-grpclb:processResources UP-TO-DATE
:grpc-grpclb:classes UP-TO-DATE
:grpc-grpclb:jar UP-TO-DATE
:grpc-grpclb:javadoc UP-TO-DATE
:grpc-grpclb:javadocJar UP-TO-DATE
:grpc-grpclb:sourcesJar UP-TO-DATE
:grpc-grpclb:signArchives SKIPPED
:grpc-grpclb:install
:grpc-testing-proto:deleteGeneratedSourcemain
:grpc-testing-proto:extractIncludeProto UP-TO-DATE
:grpc-testing-proto:extractProto UP-TO-DATE
:grpc-testing-proto:generateProto
:grpc-testing-proto:compileJava UP-TO-DATE
:grpc-testing-proto:processResources UP-TO-DATE
:grpc-testing-proto:classes UP-TO-DATE
:grpc-testing-proto:jar UP-TO-DATE
:grpc-interop-testing:compileJava UP-TO-DATE
:grpc-interop-testing:processResources UP-TO-DATE
:grpc-interop-testing:classes UP-TO-DATE
:grpc-interop-testing:jar UP-TO-DATE
:grpc-interop-testing:http2_client UP-TO-DATE
:grpc-interop-testing:reconnect_test_client UP-TO-DATE
:grpc-interop-testing:startScripts SKIPPED
:grpc-interop-testing:stresstest_client UP-TO-DATE
:grpc-interop-testing:test_client UP-TO-DATE
:grpc-interop-testing:test_server UP-TO-DATE
:grpc-interop-testing:distTar
:grpc-interop-testing:distZip
:grpc-testing-proto:javadoc NO-SOURCE
:grpc-interop-testing:javadoc UP-TO-DATE
:grpc-interop-testing:javadocJar UP-TO-DATE
:grpc-interop-testing:sourcesJar UP-TO-DATE
:grpc-interop-testing:signArchives SKIPPED
:grpc-interop-testing:install
:grpc-netty:javadocJar UP-TO-DATE
:grpc-netty:sourcesJar UP-TO-DATE
:grpc-netty:signArchives SKIPPED
:grpc-netty:install
:grpc-okhttp:javadocJar UP-TO-DATE
:grpc-okhttp:sourcesJar UP-TO-DATE
:grpc-okhttp:signArchives SKIPPED
:grpc-okhttp:install
:grpc-protobuf:javadocJar UP-TO-DATE
:grpc-protobuf:sourcesJar UP-TO-DATE
:grpc-protobuf:signArchives SKIPPED
:grpc-protobuf:install
:grpc-protobuf-lite:javadocJar UP-TO-DATE
:grpc-protobuf-lite:sourcesJar UP-TO-DATE
:grpc-protobuf-lite:signArchives SKIPPED
:grpc-protobuf-lite:install
:grpc-protobuf-nano:javadocJar UP-TO-DATE
:grpc-protobuf-nano:sourcesJar UP-TO-DATE
:grpc-protobuf-nano:signArchives SKIPPED
:grpc-protobuf-nano:install
:grpc-services:deleteGeneratedSourcemain
:grpc-services:extractIncludeProto UP-TO-DATE
:grpc-services:extractProto UP-TO-DATE
:grpc-services:generateProto
:grpc-services:compileJava UP-TO-DATE
:grpc-services:processResources UP-TO-DATE
:grpc-services:classes UP-TO-DATE
:grpc-services:jar UP-TO-DATE
:grpc-services:javadoc UP-TO-DATE
:grpc-services:javadocJar UP-TO-DATE
:grpc-services:sourcesJar UP-TO-DATE
:grpc-services:signArchives SKIPPED
:grpc-services:install
:grpc-stub:javadocJar UP-TO-DATE
:grpc-stub:sourcesJar UP-TO-DATE
:grpc-stub:signArchives SKIPPED
:grpc-stub:install
:grpc-testing:javadocJar UP-TO-DATE
:grpc-testing:sourcesJar UP-TO-DATE
:grpc-testing:signArchives SKIPPED
:grpc-testing:install
:grpc-testing-proto:javadocJar UP-TO-DATE
:grpc-testing-proto:sourcesJar UP-TO-DATE
:grpc-testing-proto:signArchives SKIPPED
:grpc-testing-proto:install
:grpc-thrift:compileJava UP-TO-DATE
:grpc-thrift:processResources NO-SOURCE
:grpc-thrift:classes UP-TO-DATE
:grpc-thrift:jar UP-TO-DATE
:grpc-thrift:javadoc UP-TO-DATE
:grpc-thrift:javadocJar UP-TO-DATE
:grpc-thrift:sourcesJar UP-TO-DATE
:grpc-thrift:signArchives SKIPPED
:grpc-thrift:install

BUILD SUCCESSFUL

Total time: 23.412 secs
```

Check
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ ls ~/.m2/repository/io/grpc/
grpc-all		grpc-core		grpc-okhttp		grpc-services		grpc-thrift
grpc-auth		grpc-grpclb		grpc-protobuf		grpc-stub		protoc-gen-grpc-java
grpc-benchmarks		grpc-interop-testing	grpc-protobuf-lite	grpc-testing
grpc-context		grpc-netty		grpc-protobuf-nano	grpc-testing-proto
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ ls ~/.m2/repository/io/grpc/protoc-gen-grpc-java/
1.4.0-SNAPSHOT  maven-metadata-local.xml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ ls ~/.m2/repository/io/grpc/protoc-gen-grpc-java/1.4.0-SNAPSHOT/
maven-metadata-local.xml				protoc-gen-grpc-java-1.4.0-SNAPSHOT.pom
protoc-gen-grpc-java-1.4.0-SNAPSHOT-linux-x86_64.exe
```

Examples
--------

With 'gradle'
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ which javac
/opt/jdk1.8.0_112/bin/javac
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ which java
/opt/jdk1.8.0_112/bin/java
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ which gradle
/opt/gradle/gradle-3.5/bin/gradle
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ gradle --version

------------------------------------------------------------
Gradle 3.5
------------------------------------------------------------

Build time:   2017-04-10 13:37:25 UTC
Revision:     b762622a185d59ce0cfc9cbc6ab5dd22469e18a6

Groovy:       2.4.10
Ant:          Apache Ant(TM) version 1.9.6 compiled on June 29 2015
JVM:          1.8.0_112 (Oracle Corporation 25.112-b15)
OS:           Linux 3.13.0-76-generic amd64

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java$ cd examples/
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java/examples$ gradle installDist
:extractIncludeProto UP-TO-DATE
:extractProto UP-TO-DATE
:generateProto
:compileJava
:processResources
:classes
:jar
:compressingHelloWorldClient
:helloWorldClient
:helloWorldServer
:routeGuideClient
:routeGuideServer
:startScripts SKIPPED
:installDist

BUILD SUCCESSFUL

Total time: 11.795 secs
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java/examples$ ls build/install/examples/bin/
compileJava                         hello-world-client      hello-world-server.bat  route-guide-client.bat
compressing-hello-world-client      hello-world-client.bat  jar                     route-guide-server
compressing-hello-world-client.bat  hello-world-server      route-guide-client      route-guide-server.bat
```

Or, with `mvn`
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java/examples$ mvn verify 
[INFO] Scanning for projects...
[INFO] ------------------------------------------------------------------------
[INFO] Detecting the operating system and CPU architecture
[INFO] ------------------------------------------------------------------------
[INFO] os.detected.name: linux
[INFO] os.detected.arch: x86_64
[INFO] os.detected.release: ubuntu
[INFO] os.detected.release.version: 14.04
[INFO] os.detected.release.like.ubuntu: true
[INFO] os.detected.release.like.debian: true
[INFO] os.detected.classifier: linux-x86_64
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building examples 1.4.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------

-------------------------------------------------------
 T E S T S
-------------------------------------------------------
Running io.grpc.examples.header.HeaderClientInterceptorTest
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 1.535 sec
Running io.grpc.examples.header.HeaderServerInterceptorTest
May 08, 2017 2:45:52 AM io.grpc.examples.header.HeaderServerInterceptor interceptCall
INFO: header received from client:Metadata(grpc-accept-encoding=gzip)
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.046 sec
Running io.grpc.examples.helloworld.HelloWorldClientTest
May 08, 2017 2:45:52 AM io.grpc.examples.helloworld.HelloWorldClient greet
INFO: Will try to greet test name ...
May 08, 2017 2:45:52 AM io.grpc.examples.helloworld.HelloWorldClient greet
WARNING: RPC failed: Status{code=UNIMPLEMENTED, description=Method helloworld.Greeter/SayHello is unimplemented, cause=null}
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.088 sec
Running io.grpc.examples.helloworld.HelloWorldServerTest
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.013 sec
Running io.grpc.examples.routeguide.RouteGuideClientTest
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** RecordRoute
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient warning
WARNING: RecordRoute Failed: Status{code=INVALID_ARGUMENT, description=null, cause=null}
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Visiting point 0, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** RouteChat
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "First message" at 0, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Got message "First message" at 0, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "Second message" at 0, 1
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Got message "Second message" at 0, 1
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "Third message" at 1, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Got message "Third message" at 1, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "Fourth message" at 1, 1
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Got message "Fourth message" at 1, 1
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Finished RouteChat
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** RouteChat
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "First message" at 0, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "Second message" at 0, 1
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "Third message" at 1, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "Fourth message" at 1, 1
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Got message "dummy msg1" at 0, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Got message "dummy msg2" at 0, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Finished RouteChat
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** GetFeature: lat=-1 lon=-1
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Found feature called "dummyFeature" at -0, -0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** RouteChat
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "First message" at 0, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient warning
WARNING: RouteChat Failed: Status{code=PERMISSION_DENIED, description=null, cause=null}
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "Second message" at 0, 1
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "Third message" at 1, 0
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Sending message "Fourth message" at 1, 1
May 08, 2017 2:45:52 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** ListFeatures: lowLat=1 lowLon=2 hiLat=3 hiLon=4
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Result #1: name: "feature 1"

May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Result #2: name: "feature 2"

May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** RecordRoute
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Finished trip with 0 points. Passed 0 features. Travelled 0 meters. It took 0 seconds.
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient warning
WARNING: RecordRoute Failed: Status{code=CANCELLED, description=Failed to read message., cause=io.grpc.StatusRuntimeException: INTERNAL: More than one responses received for unary or client-streaming call
	at io.grpc.Status.asRuntimeException(Status.java:534)
	at io.grpc.stub.ClientCalls$StreamObserverToCallListenerAdapter.onMessage(ClientCalls.java:379)
	at io.grpc.internal.ClientCallImpl$ClientStreamListenerImpl$1MessageRead.runInContext(ClientCallImpl.java:489)
	at io.grpc.internal.ContextRunnable.run(ContextRunnable.java:52)
	at io.grpc.internal.SerializeReentrantCallsDirectExecutor.execute(SerializeReentrantCallsDirectExecutor.java:64)
	at io.grpc.internal.ClientCallImpl$ClientStreamListenerImpl.messageRead(ClientCallImpl.java:502)
	at io.grpc.inprocess.InProcessTransport$InProcessStream$InProcessServerStream.clientRequested(InProcessTransport.java:312)
	at io.grpc.inprocess.InProcessTransport$InProcessStream$InProcessServerStream.access$2000(InProcessTransport.java:262)
	at io.grpc.inprocess.InProcessTransport$InProcessStream$InProcessClientStream.request(InProcessTransport.java:456)
	at io.grpc.internal.ClientCallImpl.request(ClientCallImpl.java:345)
	at io.grpc.stub.ClientCalls.startCall(ClientCalls.java:282)
	at io.grpc.stub.ClientCalls.asyncStreamingRequestCall(ClientCalls.java:266)
	at io.grpc.stub.ClientCalls.asyncClientStreamingCall(ClientCalls.java:97)
	at io.grpc.examples.routeguide.RouteGuideGrpc$RouteGuideStub.recordRoute(RouteGuideGrpc.java:256)
	at io.grpc.examples.routeguide.RouteGuideClient.recordRoute(RouteGuideClient.java:179)
	at io.grpc.examples.routeguide.RouteGuideClientTest.recordRoute_wrongResponse(RouteGuideClientTest.java:357)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:498)
	at org.junit.runners.model.FrameworkMethod$1.runReflectiveCall(FrameworkMethod.java:47)
	at org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12)
	at org.junit.runners.model.FrameworkMethod.invokeExplosively(FrameworkMethod.java:44)
	at org.junit.internal.runners.statements.InvokeMethod.evaluate(InvokeMethod.java:17)
	at org.junit.internal.runners.statements.RunBefores.evaluate(RunBefores.java:26)
	at org.junit.internal.runners.statements.RunAfters.evaluate(RunAfters.java:27)
	at org.junit.runners.ParentRunner.runLeaf(ParentRunner.java:271)
	at org.junit.runners.BlockJUnit4ClassRunner.runChild(BlockJUnit4ClassRunner.java:70)
	at org.junit.runners.BlockJUnit4ClassRunner.runChild(BlockJUnit4ClassRunner.java:50)
	at org.junit.runners.ParentRunner$3.run(ParentRunner.java:238)
	at org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:63)
	at org.junit.runners.ParentRunner.runChildren(ParentRunner.java:236)
	at org.junit.runners.ParentRunner.access$000(ParentRunner.java:53)
	at org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:229)
	at org.junit.runners.ParentRunner.run(ParentRunner.java:309)
	at org.apache.maven.surefire.junit4.JUnit4Provider.execute(JUnit4Provider.java:252)
	at org.apache.maven.surefire.junit4.JUnit4Provider.executeTestSet(JUnit4Provider.java:141)
	at org.apache.maven.surefire.junit4.JUnit4Provider.invoke(JUnit4Provider.java:112)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:498)
	at org.apache.maven.surefire.util.ReflectionUtils.invokeMethodWithArray(ReflectionUtils.java:189)
	at org.apache.maven.surefire.booter.ProviderFactory$ProviderProxy.invoke(ProviderFactory.java:165)
	at org.apache.maven.surefire.booter.ProviderFactory.invokeProvider(ProviderFactory.java:85)
	at org.apache.maven.surefire.booter.ForkedBooter.runSuitesInProcess(ForkedBooter.java:115)
	at org.apache.maven.surefire.booter.ForkedBooter.main(ForkedBooter.java:75)
}
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Visiting point 0, 0
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** RecordRoute
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Visiting point 0, 0
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Visiting point 0, 0
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Visiting point 0, 0
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Visiting point 0, 0
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Finished trip with 7 points. Passed 8 features. Travelled 9 meters. It took 10 seconds.
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Finished RecordRoute
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** ListFeatures: lowLat=1 lowLon=2 hiLat=3 hiLon=4
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: Result #1: name: "feature 1"

May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient warning
WARNING: RPC failed: Status{code=INVALID_ARGUMENT, description=null, cause=null}
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient info
INFO: *** GetFeature: lat=-1 lon=-1
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideClient warning
WARNING: RPC failed: Status{code=DATA_LOSS, description=null, cause=null}
Tests run: 10, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.782 sec
Running io.grpc.examples.routeguide.RouteGuideServerTest
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideServer start
INFO: Server started, listening on 0
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideServer start
INFO: Server started, listening on 0
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideServer start
INFO: Server started, listening on 0
May 08, 2017 2:45:53 AM io.grpc.examples.routeguide.RouteGuideServer start
INFO: Server started, listening on 0
Tests run: 4, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.061 sec
*** shutting down gRPC server since JVM is shutting down
*** server shut down
*** shutting down gRPC server since JVM is shutting down
*** server shut down
*** shutting down gRPC server since JVM is shutting down
*** server shut down
*** shutting down gRPC server since JVM is shutting down
*** server shut down

Results :

Tests run: 18, Failures: 0, Errors: 0, Skipped: 0

[INFO] 
[INFO] --- maven-jar-plugin:2.4:jar (default-jar) @ examples ---
[INFO] Building jar: /work/src/github.com/grpc/grpc-java/examples/target/examples-1.4.0-SNAPSHOT.jar
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 35.717 s
[INFO] Finished at: 2017-05-08T02:45:57+00:00
[INFO] Final Memory: 21M/162M
[INFO] ------------------------------------------------------------------------
```

Experimental
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java/examples$ ./build/install/examples/bin/hello-world-server
May 08, 2017 2:50:29 AM io.grpc.examples.helloworld.HelloWorldServer start
INFO: Server started, listening on 50051
^Z
[1]+  Stopped                 ./build/install/examples/bin/hello-world-server
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java/examples$ bg %1
[1]+ ./build/install/examples/bin/hello-world-server &
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java/examples$ jobs
[1]+  Running                 ./build/install/examples/bin/hello-world-server &
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java/examples$ ./build/install/examples/bin/hello-world-client
May 08, 2017 2:52:00 AM io.grpc.examples.helloworld.HelloWorldClient greet
INFO: Will try to greet world ...
May 08, 2017 2:52:01 AM io.grpc.examples.helloworld.HelloWorldClient greet
INFO: Greeting: Hello world
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java/examples$ fg %1
./build/install/examples/bin/hello-world-server
^C*** shutting down gRPC server since JVM is shutting down
*** server shut down
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/grpc/grpc-java/examples$ jobs
```
