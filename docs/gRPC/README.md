Build gRPC-Java development environment
=======================================

Environment
-----------

Ubuntu 14.04 VBox

Protobuf development environment

Get repository
--------------

With `git`
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com$ git clone https://github.com/grpc/grpc-java grpc/grpc-java
```

Build and Install
-----------------

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
