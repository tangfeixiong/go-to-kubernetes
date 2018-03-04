# Table of Contents

console-k8s

hadoop-k8s

## Kubernetes client for Java项目io.fabric8：kubernetes-client示例

Github仓库 －－ https://github.com/fabric8io/kubernetes-client

### 环境

Java

* Mac
* Kubernetes client config文件：/Users/fanhongling/Downloads/kubeconfig

Kubernetes

* Mac的VirtualBox linux (kube-apiserver为https://172.17.4.200:6443)
* Kubernetes client config文件：/home/vagrant/.kube/config

其它

* workspace以vboxfs，Mac和VM共享
* m2 repositories也为共享

### 在Mac

从Github获取项目
```
fanhonglingdeMacBook-Pro:~ fanhongling$ git clone https://github.com/fabric8io/kubernetes-client Downloads/workspace/src/github.com/fabric8io/kubernetes-client
fanhonglingdeMacBook-Pro:~ fanhongling$ cd Downloads/workspace/src/github.com/fabric8io/kubernetes-client/
```

Build项目到m2 local
```
fanhonglingdeMacBook-Pro:kubernetes-client fanhongling$ mvn clean install
[INFO] Scanning for projects...
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Build Order:
[INFO] 
[INFO] Fabric8 :: Kubernetes :: Project
[INFO] Fabric8 :: Kubernetes :: Java Client
[INFO] Fabric8 :: Kubernetes :: Server Mock
[INFO] Fabric8 :: Openshift :: Java Client
[INFO] Fabric8 :: Openshift :: Server Mock
[INFO] Fabric8 :: Kubernetes :: Client Examples
[INFO] Fabric8 :: Kubernetes :: Platforms
[INFO] Fabric8 :: Kubernetes :: Platforms :: Karaf
[INFO] Fabric8 :: Kubernetes :: Platforms :: Karaf :: Features
[INFO] Fabric8 :: Kubernetes :: Platforms :: Karaf :: Tests
[INFO] Fabric8 :: Kubernetes :: Tests
[INFO] Fabric8 :: Kubernetes and Openshift :: UberJar
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Fabric8 :: Kubernetes :: Project 2.2-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.5:clean (default-clean) @ kubernetes-client-project ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/github.com/fabric8io/kubernetes-client/target
[INFO] 
[INFO] --- maven-enforcer-plugin:1.2:enforce (enforce-maven) @ kubernetes-client-project ---
[INFO] 
[INFO] --- sundr-maven-plugin:0.3.12:generate-bom (default) @ kubernetes-client-project ---
Downloading: ###下载依赖###
[INFO] --- maven-install-plugin:2.4:install (default-install) @ kubernetes-openshift-uberjar ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/fabric8io/kubernetes-client/uberjar/target/kubernetes-openshift-uberjar-2.2-SNAPSHOT.jar to /Users/fanhongling/.m2/repository/io/fabric8/kubernetes-openshift-uberjar/2.2-SNAPSHOT/kubernetes-openshift-uberjar-2.2-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/fabric8io/kubernetes-client/uberjar/dependency-reduced-pom.xml to /Users/fanhongling/.m2/repository/io/fabric8/kubernetes-openshift-uberjar/2.2-SNAPSHOT/kubernetes-openshift-uberjar-2.2-SNAPSHOT.pom
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/fabric8io/kubernetes-client/uberjar/target/kubernetes-openshift-uberjar-2.2-SNAPSHOT-versioned.jar to /Users/fanhongling/.m2/repository/io/fabric8/kubernetes-openshift-uberjar/2.2-SNAPSHOT/kubernetes-openshift-uberjar-2.2-SNAPSHOT-versioned.jar
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Summary:
[INFO] 
[INFO] Fabric8 :: Kubernetes :: Project ................... SUCCESS [09:28 min]
[INFO] Fabric8 :: Kubernetes :: Java Client ............... SUCCESS [02:38 min]
[INFO] Fabric8 :: Kubernetes :: Server Mock ............... SUCCESS [ 16.352 s]
[INFO] Fabric8 :: Openshift :: Java Client ................ SUCCESS [ 12.429 s]
[INFO] Fabric8 :: Openshift :: Server Mock ................ SUCCESS [  0.291 s]
[INFO] Fabric8 :: Kubernetes :: Client Examples ........... SUCCESS [  0.554 s]
[INFO] Fabric8 :: Kubernetes :: Platforms ................. SUCCESS [  0.016 s]
[INFO] Fabric8 :: Kubernetes :: Platforms :: Karaf ........ SUCCESS [  0.060 s]
[INFO] Fabric8 :: Kubernetes :: Platforms :: Karaf :: Features SUCCESS [  5.080 s]
[INFO] Fabric8 :: Kubernetes :: Platforms :: Karaf :: Tests SUCCESS [16:22 min]
[INFO] Fabric8 :: Kubernetes :: Tests ..................... SUCCESS [04:39 min]
[INFO] Fabric8 :: Kubernetes and Openshift :: UberJar ..... SUCCESS [ 11.027 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 33:56 min
[INFO] Finished at: 2017-05-07T00:20:33-07:00
[INFO] Final Memory: 55M/337M
[INFO] ------------------------------------------------------------------------
```

运行示例
```
fanhonglingdeMacBook-Pro:~ fanhongling$ cd kubernetes-examples
fanhonglingdeMacBook-Pro:kubernetes-examples fanhongling$ KUBECONFIG=/Users/fanhongling/Downloads/kubeconfig mvn clean compile exec:java -Dexec.mainClass="io.fabric8.kubernetes.examples.ListExamples" -Dexec.args="https://172.17.4.200:6443" -e
[INFO] Error stacktraces are turned on.
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Fabric8 :: Kubernetes :: Client Examples 2.2-SNAPSHOT
[INFO] ------------------------------------------------------------------------
### 下载，clean和其它依赖插件 ###
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ kubernetes-examples ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 18 source files to /Users/fanhongling/Downloads/workspace/src/github.com/fabric8io/kubernetes-client/kubernetes-examples/target/classes
[INFO] 
[INFO] --- maven-clean-plugin:2.5:clean (default-clean) @ kubernetes-examples ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/github.com/fabric8io/kubernetes-client/kubernetes-examples/target
[INFO] 
[INFO] --- exec-maven-plugin:1.4.0:java (default-cli) @ kubernetes-examples ---
00:50:37.655 [io.fabric8.kubernetes.examples.ListExamples.main()] DEBUG io.fabric8.kubernetes.client.Config - Trying to configure client from Kubernetes config...
00:50:37.665 [io.fabric8.kubernetes.examples.ListExamples.main()] DEBUG io.fabric8.kubernetes.client.Config - Found for Kubernetes config at: [/Users/fanhongling/Downloads/kubeconfig].
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: --> GET https://172.17.4.200:6443/api/v1/namespaces http/1.1
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Host: 172.17.4.200:6443
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Connection: Keep-Alive
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Accept-Encoding: gzip
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: User-Agent: okhttp/3.6.0
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: --> END GET
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: <-- 200 OK https://172.17.4.200:6443/api/v1/namespaces (40ms)
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Content-Type: application/json
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Date: Sun, 07 May 2017 07:41:38 GMT
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Content-Length: 873
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: 
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: {"kind":"NamespaceList","apiVersion":"v1","metadata":{"selfLink":"/api/v1/namespaces","resourceVersion":"228108"},"items":[{"metadata":{"name":"default","selfLink":"/api/v1/namespacesdefault","uid":"2e95dea2-2fac-11e7-84fa-080027360e7e","resourceVersion":"10","creationTimestamp":"2017-05-03T02:57:02Z"},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}},{"metadata":{"name":"kafka","selfLink":"/api/v1/namespaceskafka","uid":"1f0d9656-315d-11e7-93cb-080027360e7e","resourceVersion":"99591","creationTimestamp":"2017-05-05T06:36:08Z"},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}},{"metadata":{"name":"kube-system","selfLink":"/api/v1/namespaceskube-system","uid":"2e9df2f1-2fac-11e7-84fa-080027360e7e","resourceVersion":"24","creationTimestamp":"2017-05-03T02:57:02Z"},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}]}

五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: <-- END HTTP (873-byte body)
NamespaceList(apiVersion=v1, items=[Namespace(apiVersion=v1, kind=Namespace, metadata=ObjectMeta(annotations=null, clusterName=null, creationTimestamp=2017-05-03T02:57:02Z, deletionGracePeriodSeconds=null, deletionTimestamp=null, finalizers=[], generateName=null, generation=null, labels=null, name=default, namespace=null, ownerReferences=[], resourceVersion=10, selfLink=/api/v1/namespacesdefault, uid=2e95dea2-2fac-11e7-84fa-080027360e7e, additionalProperties={}), spec=NamespaceSpec(finalizers=[kubernetes], additionalProperties={}), status=NamespaceStatus(phase=Active, additionalProperties={}), additionalProperties={}), Namespace(apiVersion=v1, kind=Namespace, metadata=ObjectMeta(annotations=null, clusterName=null, creationTimestamp=2017-05-05T06:36:08Z, deletionGracePeriodSeconds=null, deletionTimestamp=null, finalizers=[], generateName=null, generation=null, labels=null, name=kafka, namespace=null, ownerReferences=[], resourceVersion=99591, selfLink=/api/v1/namespaceskafka, uid=1f0d9656-315d-11e7-93cb-080027360e7e, additionalProperties={}), spec=NamespaceSpec(finalizers=[kubernetes], additionalProperties={}), status=NamespaceStatus(phase=Active, additionalProperties={}), additionalProperties={}), Namespace(apiVersion=v1, kind=Namespace, metadata=ObjectMeta(annotations=null, clusterName=null, creationTimestamp=2017-05-03T02:57:02Z, deletionGracePeriodSeconds=null, deletionTimestamp=null, finalizers=[], generateName=null, generation=null, labels=null, name=kube-system, namespace=null, ownerReferences=[], resourceVersion=24, selfLink=/api/v1/namespaceskube-system, uid=2e9df2f1-2fac-11e7-84fa-080027360e7e, additionalProperties={}), spec=NamespaceSpec(finalizers=[kubernetes], additionalProperties={}), status=NamespaceStatus(phase=Active, additionalProperties={}), additionalProperties={})], kind=NamespaceList, metadata=ListMeta(resourceVersion=228108, selfLink=/api/v1/namespaces, additionalProperties={}), additionalProperties={})
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: --> GET https://172.17.4.200:6443/api/v1/namespaces?labelSelector=this%3Dworks http/1.1
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Host: 172.17.4.200:6443
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Connection: Keep-Alive
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Accept-Encoding: gzip
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: User-Agent: okhttp/3.6.0
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: --> END GET
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: <-- 200 OK https://172.17.4.200:6443/api/v1/namespaces?labelSelector=this%3Dworks (5ms)
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Content-Type: application/json
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Date: Sun, 07 May 2017 07:41:38 GMT
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Content-Length: 126
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: 
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: {"kind":"NamespaceList","apiVersion":"v1","metadata":{"selfLink":"/api/v1/namespaces","resourceVersion":"228108"},"items":[]}

五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: <-- END HTTP (126-byte body)
NamespaceList(apiVersion=v1, items=[], kind=NamespaceList, metadata=ListMeta(resourceVersion=228108, selfLink=/api/v1/namespaces, additionalProperties={}), additionalProperties={})
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: --> GET https://172.17.4.200:6443/api/v1/namespaces/default/pods?labelSelector=this%3Dworks http/1.1
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Host: 172.17.4.200:6443
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Connection: Keep-Alive
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Accept-Encoding: gzip
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: User-Agent: okhttp/3.6.0
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: --> END GET
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: <-- 200 OK https://172.17.4.200:6443/api/v1/namespaces/default/pods?labelSelector=this%3Dworks (4ms)
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Content-Type: application/json
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Date: Sun, 07 May 2017 07:41:38 GMT
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: Content-Length: 133
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: 
五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: {"kind":"PodList","apiVersion":"v1","metadata":{"selfLink":"/api/v1/namespaces/default/pods","resourceVersion":"228108"},"items":[]}

五月 07, 2017 12:50:38 上午 okhttp3.internal.platform.Platform log
信息: <-- END HTTP (133-byte body)
PodList(apiVersion=v1, items=[], kind=PodList, metadata=ListMeta(resourceVersion=228108, selfLink=/api/v1/namespaces/default/pods, additionalProperties={}), additionalProperties={})
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: --> GET https://172.17.4.200:6443/api/v1/namespaces/test/pods?labelSelector=this%3Dworks http/1.1
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Host: 172.17.4.200:6443
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Connection: Keep-Alive
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Accept-Encoding: gzip
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: User-Agent: okhttp/3.6.0
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: --> END GET
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: <-- 200 OK https://172.17.4.200:6443/api/v1/namespaces/test/pods?labelSelector=this%3Dworks (4ms)
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Content-Type: application/json
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Date: Sun, 07 May 2017 07:41:39 GMT
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Content-Length: 130
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: 
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: {"kind":"PodList","apiVersion":"v1","metadata":{"selfLink":"/api/v1/namespaces/test/pods","resourceVersion":"228108"},"items":[]}

五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: <-- END HTTP (130-byte body)
PodList(apiVersion=v1, items=[], kind=PodList, metadata=ListMeta(resourceVersion=228108, selfLink=/api/v1/namespaces/test/pods, additionalProperties={}), additionalProperties={})
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: --> GET https://172.17.4.200:6443/api/v1/namespaces/test/pods/testing http/1.1
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Host: 172.17.4.200:6443
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Connection: Keep-Alive
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Accept-Encoding: gzip
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: User-Agent: okhttp/3.6.0
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: --> END GET
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: <-- 404 Not Found https://172.17.4.200:6443/api/v1/namespaces/test/pods/testing (3ms)
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Content-Type: application/json
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Date: Sun, 07 May 2017 07:41:39 GMT
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: Content-Length: 182
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: 
五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: {"kind":"Status","apiVersion":"v1","metadata":{},"status":"Failure","message":"pods \"testing\" not found","reason":"NotFound","details":{"name":"testing","kind":"pods"},"code":404}

五月 07, 2017 12:50:39 上午 okhttp3.internal.platform.Platform log
信息: <-- END HTTP (182-byte body)
null
[WARNING] thread Thread[OkHttp ConnectionPool,5,io.fabric8.kubernetes.examples.ListExamples] was interrupted but is still alive after waiting at least 15000msecs
[WARNING] thread Thread[OkHttp ConnectionPool,5,io.fabric8.kubernetes.examples.ListExamples] will linger despite being asked to die via interruption
[WARNING] thread Thread[Okio Watchdog,5,io.fabric8.kubernetes.examples.ListExamples] will linger despite being asked to die via interruption
[WARNING] NOTE: 2 thread(s) did not finish despite being asked to  via interruption. This is not a problem with exec:java, it is a problem with the running code. Although not serious, it should be remedied.
[WARNING] Couldn't destroy threadgroup org.codehaus.mojo.exec.ExecJavaMojo$IsolatedThreadGroup[name=io.fabric8.kubernetes.examples.ListExamples,maxpri=10]
java.lang.IllegalThreadStateException
	at java.lang.ThreadGroup.destroy(ThreadGroup.java:778)
	at org.codehaus.mojo.exec.ExecJavaMojo.execute(ExecJavaMojo.java:328)
	at org.apache.maven.plugin.DefaultBuildPluginManager.executeMojo(DefaultBuildPluginManager.java:134)
	at org.apache.maven.lifecycle.internal.MojoExecutor.execute(MojoExecutor.java:207)
	at org.apache.maven.lifecycle.internal.MojoExecutor.execute(MojoExecutor.java:153)
	at org.apache.maven.lifecycle.internal.MojoExecutor.execute(MojoExecutor.java:145)
	at org.apache.maven.lifecycle.internal.LifecycleModuleBuilder.buildProject(LifecycleModuleBuilder.java:116)
	at org.apache.maven.lifecycle.internal.LifecycleModuleBuilder.buildProject(LifecycleModuleBuilder.java:80)
	at org.apache.maven.lifecycle.internal.builder.singlethreaded.SingleThreadedBuilder.build(SingleThreadedBuilder.java:51)
	at org.apache.maven.lifecycle.internal.LifecycleStarter.execute(LifecycleStarter.java:128)
	at org.apache.maven.DefaultMaven.doExecute(DefaultMaven.java:307)
	at org.apache.maven.DefaultMaven.doExecute(DefaultMaven.java:193)
	at org.apache.maven.DefaultMaven.execute(DefaultMaven.java:106)
	at org.apache.maven.cli.MavenCli.execute(MavenCli.java:863)
	at org.apache.maven.cli.MavenCli.doMain(MavenCli.java:288)
	at org.apache.maven.cli.MavenCli.main(MavenCli.java:199)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:498)
	at org.codehaus.plexus.classworlds.launcher.Launcher.launchEnhanced(Launcher.java:289)
	at org.codehaus.plexus.classworlds.launcher.Launcher.launch(Launcher.java:229)
	at org.codehaus.plexus.classworlds.launcher.Launcher.mainWithExitCode(Launcher.java:415)
	at org.codehaus.plexus.classworlds.launcher.Launcher.main(Launcher.java:356)
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 22.707 s
[INFO] Finished at: 2017-05-07T00:50:54-07:00
[INFO] Final Memory: 31M/322M
[INFO] ------------------------------------------------------------------------
```

### 在Linux VM

以Mac中的build运行examples
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/fabric8io/kubernetes-client/kubernetes-examples$ mvn exec:java -Dexec.mainClass="io.fabric8.kubernetes.examples.FullExample" -Dexec.args="https://172.17.4.200:6443"
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Fabric8 :: Kubernetes :: Client Examples 2.2-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- exec-maven-plugin:1.4.0:java (default-cli) @ kubernetes-examples ---
08:17:46.798 [io.fabric8.kubernetes.examples.FullExample.main()] DEBUG io.fabric8.kubernetes.client.Config - Trying to configure client from Kubernetes config...
08:17:46.830 [io.fabric8.kubernetes.examples.FullExample.main()] DEBUG io.fabric8.kubernetes.client.Config - Found for Kubernetes config at: [/home/vagrant/.kube/config].
08:17:47.835 [io.fabric8.kubernetes.examples.FullExample.main()] DEBUG io.fabric8.kubernetes.client.dsl.internal.WatchConnectionManager - Connecting websocket ... io.fabric8.kubernetes.client.dsl.internal.WatchConnectionManager@557f14c5
08:17:48.392 [OkHttp https://172.17.4.200:6443/...] DEBUG io.fabric8.kubernetes.client.dsl.internal.WatchConnectionManager - WebSocket successfully opened
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: --> POST https://172.17.4.200:6443/api/v1/namespaces http/1.1
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: Content-Type: application/json; charset=utf-8
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: Content-Length: 152
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: Host: 172.17.4.200:6443
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: Connection: Keep-Alive
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: Accept-Encoding: gzip
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: User-Agent: okhttp/3.6.0
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: 
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: {"apiVersion":"v1","kind":"Namespace","metadata":{"annotations":{},"finalizers":[],"labels":{"this":"rocks"},"name":"thisisatest","ownerReferences":[]}}
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: --> END POST (152-byte body)
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: <-- 201 Created https://172.17.4.200:6443/api/v1/namespaces (57ms)
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: Content-Type: application/json
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: Date: Sun, 07 May 2017 08:17:48 GMT
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: Content-Length: 322
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: 
May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: {"kind":"Namespace","apiVersion":"v1","metadata":{"name":"thisisatest","selfLink":"/api/v1/namespacesthisisatest","uid":"a82099ea-32fd-11e7-8e74-080027360e7e","resourceVersion":"230496","creationTimestamp":"2017-05-07T08:17:48Z","labels":{"this":"rocks"}},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}

May 07, 2017 8:17:48 AM okhttp3.internal.platform.Platform log
INFO: <-- END HTTP (322-byte body)
### 省略......
08:17:49.914 [io.fabric8.kubernetes.examples.FullExample.main()] INFO  io.fabric8.kubernetes.examples.FullExample - Get RC by name in namespace: ReplicationController(apiVersion=v1, kind=ReplicationController, metadata=ObjectMeta(annotations=null, clusterName=null, creationTimestamp=2017-05-07T08:17:49Z, deletionGracePeriodSeconds=null, deletionTimestamp=null, finalizers=[], generateName=null, generation=1, labels={server=nginx}, name=nginx-controller, namespace=thisisatest, ownerReferences=[], resourceVersion=230519, selfLink=/api/v1/namespaces/thisisatest/replicationcontrollers/nginx-controller, uid=a89cae64-32fd-11e7-8e74-080027360e7e, additionalProperties={}), spec=ReplicationControllerSpec(minReadySeconds=null, replicas=3, selector={server=nginx}, template=PodTemplateSpec(metadata=ObjectMeta(annotations=null, clusterName=null, creationTimestamp=null, deletionGracePeriodSeconds=null, deletionTimestamp=null, finalizers=[], generateName=null, generation=null, labels={server=nginx}, name=null, namespace=null, ownerReferences=[], resourceVersion=null, selfLink=null, uid=null, additionalProperties={}), spec=PodSpec(activeDeadlineSeconds=null, containers=[Container(args=[], command=[], env=[], image=nginx, imagePullPolicy=Always, lifecycle=null, livenessProbe=null, name=nginx, ports=[ContainerPort(containerPort=80, hostIP=null, hostPort=null, name=null, protocol=TCP, additionalProperties={})], readinessProbe=null, resources=ResourceRequirements(limits=null, requests=null, additionalProperties={}), securityContext=null, stdin=null, stdinOnce=null, terminationMessagePath=/dev/termination-log, tty=null, volumeMounts=[], workingDir=null, additionalProperties={})], dnsPolicy=ClusterFirst, hostIPC=null, hostNetwork=null, hostPID=null, hostname=null, imagePullSecrets=[], nodeName=null, nodeSelector=null, restartPolicy=Always, securityContext=PodSecurityContext(fsGroup=null, runAsNonRoot=null, runAsUser=null, seLinuxOptions=null, supplementalGroups=[], additionalProperties={}), serviceAccount=null, serviceAccountName=null, subdomain=null, terminationGracePeriodSeconds=30, volumes=[], additionalProperties={}), additionalProperties={}), additionalProperties={}), status=ReplicationControllerStatus(availableReplicas=null, conditions=[], fullyLabeledReplicas=3, observedGeneration=1, readyReplicas=null, replicas=3, additionalProperties={}), additionalProperties={})
### 省略......
08:17:50.112 [io.fabric8.kubernetes.examples.FullExample.main()] INFO  io.fabric8.kubernetes.examples.FullExample - Dump RC as YAML: ---
apiVersion: "v1"
kind: "ReplicationController"
metadata:
  creationTimestamp: "2017-05-07T08:17:49Z"
  finalizers: []
  generation: 1
  labels:
    server: "nginx"
  name: "nginx-controller"
  namespace: "thisisatest"
  ownerReferences: []
  resourceVersion: "230519"
  selfLink: "/api/v1/namespaces/thisisatest/replicationcontrollers/nginx-controller"
  uid: "a89cae64-32fd-11e7-8e74-080027360e7e"
spec:
  replicas: 3
  selector:
    server: "nginx"
  template:
    metadata:
      finalizers: []
      labels:
        server: "nginx"
      ownerReferences: []
    spec:
      containers:
      - args: []
        command: []
        env: []
        image: "nginx"
        imagePullPolicy: "Always"
        name: "nginx"
        ports:
        - containerPort: 80
          protocol: "TCP"
        resources: {}
        terminationMessagePath: "/dev/termination-log"
        volumeMounts: []
      dnsPolicy: "ClusterFirst"
      imagePullSecrets: []
      restartPolicy: "Always"
      securityContext:
        supplementalGroups: []
      terminationGracePeriodSeconds: 30
      volumes: []
status:
  conditions: []
  fullyLabeledReplicas: 3
  observedGeneration: 1
  replicas: 3

08:17:50.215 [io.fabric8.kubernetes.examples.FullExample.main()] INFO  io.fabric8.kubernetes.examples.FullExample - Dump RC as YAML without state: ---
apiVersion: "v1"
kind: "ReplicationController"
metadata:
  finalizers: []
  labels:
    server: "nginx"
  name: "nginx-controller"
  namespace: "thisisatest"
  ownerReferences: []
spec:
  replicas: 3
  selector:
    server: "nginx"
  template:
    metadata:
      finalizers: []
      labels:
        server: "nginx"
      ownerReferences: []
    spec:
      containers:
      - args: []
        command: []
        env: []
        image: "nginx"
        imagePullPolicy: "Always"
        name: "nginx"
        ports:
        - containerPort: 80
          protocol: "TCP"
        resources: {}
        terminationMessagePath: "/dev/termination-log"
        volumeMounts: []
      dnsPolicy: "ClusterFirst"
      imagePullSecrets: []
      restartPolicy: "Always"
      securityContext:
        supplementalGroups: []
      terminationGracePeriodSeconds: 30
      volumes: []

May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: --> GET https://172.17.4.200:6443/api/v1/replicationcontrollers?labelSelector=server%3Dnginx http/1.1
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: Host: 172.17.4.200:6443
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: Connection: Keep-Alive
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: Accept-Encoding: gzip
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: User-Agent: okhttp/3.6.0
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: --> END GET
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: <-- 200 OK https://172.17.4.200:6443/api/v1/replicationcontrollers?labelSelector=server%3Dnginx (8ms)
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: Content-Type: application/json
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: Date: Sun, 07 May 2017 08:17:50 GMT
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: Content-Length: 1758
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: 
May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: {"kind":"ReplicationControllerList","apiVersion":"v1","metadata":{"selfLink":"/api/v1/replicationcontrollers","resourceVersion":"230524"},"items":[{"metadata":{"name":"nginx-controller","namespace":"thisisatest","selfLink":"/api/v1/namespaces/thisisatest/replicationcontrollers/nginx-controller","uid":"a89cae64-32fd-11e7-8e74-080027360e7e","resourceVersion":"230519","generation":1,"creationTimestamp":"2017-05-07T08:17:49Z","labels":{"server":"nginx"}},"spec":{"replicas":3,"selector":{"server":"nginx"},"template":{"metadata":{"creationTimestamp":null,"labels":{"server":"nginx"}},"spec":{"containers":[{"name":"nginx","image":"nginx","ports":[{"containerPort":80,"protocol":"TCP"}],"resources":{},"terminationMessagePath":"/dev/termination-log","imagePullPolicy":"Always"}],"restartPolicy":"Always","terminationGracePeriodSeconds":30,"dnsPolicy":"ClusterFirst","securityContext":{}}}},"status":{"replicas":3,"fullyLabeledReplicas":3,"observedGeneration":1}},{"metadata":{"name":"nginx2-controller","namespace":"thisisatest","selfLink":"/api/v1/namespaces/thisisatest/replicationcontrollers/nginx2-controller","uid":"a8c60ca4-32fd-11e7-8e74-080027360e7e","resourceVersion":"230522","generation":1,"creationTimestamp":"2017-05-07T08:17:49Z","labels":{"server":"nginx"}},"spec":{"replicas":0,"selector":{"server":"nginx2"},"template":{"metadata":{"creationTimestamp":null,"labels":{"server":"nginx2"}},"spec":{"containers":[{"name":"nginx","image":"nginx","ports":[{"containerPort":80,"protocol":"TCP"}],"resources":{},"terminationMessagePath":"/dev/termination-log","imagePullPolicy":"Always"}],"restartPolicy":"Always","terminationGracePeriodSeconds":30,"dnsPolicy":"ClusterFirst","securityContext":{}}}},"status":{"replicas":0,"observedGeneration":1}}]}

May 07, 2017 8:17:50 AM okhttp3.internal.platform.Platform log
INFO: <-- END HTTP (1758-byte body)
### 省略......
^C
```

从Mac上执行`kubectl`命令
```
 /Users/fanhongling/Downloads/99-mirror/https%3A%2F%2Fdl.k8s.io/v1.6.2/kubernetes-client-darwin-amd64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     51      0  0:00:03  0:00:03 --:--:--    51
100 28.7M  100 28.7M    0     0   856k      0  0:00:34  0:00:34 --:--:-- 1410k
fanhonglingdeMacBook-Pro:kubernetes-examples fanhongling$ tar -ztf /Users/fanhongling/Downloads/99-mirror/https%3A%2F%2Fdl.k8s.io/v1.6.2/kubernetes-client-darwin-amd64.tar.gz
kubernetes/
kubernetes/client/
kubernetes/client/bin/
kubernetes/client/bin/kubefed
kubernetes/client/bin/kubectl
fanhonglingdeMacBook-Pro:kubernetes-examples fanhongling$ tar -C /Users/fanhongling/Downloads/workspace/bin -zxf /Users/fanhongling/Downloads/99-mirror/https%3A%2F%2Fdl.k8s.io/v1.6.2/kubernetes-client-darwin-amd64.tar.gz --strip-components=3
fanhonglingdeMacBook-Pro:kubernetes-examples fanhongling$ KUBECONFIG=/Users/fanhongling/Downloads/kubeconfig kubectl get ns
NAME          STATUS    AGE
default       Active    4d
kafka         Active    2d
kube-system   Active    4d
thisisatest   Active    15m
fanhonglingdeMacBook-Pro:kubernetes-examples fanhongling$ KUBECONFIG=/Users/fanhongling/Downloads/kubeconfig kubectl get all --namespace=thisisatest
NAME                                                         READY     STATUS             RESTARTS   AGE
po/nginx-controller-20k8x                                    0/1       ImagePullBackOff   0          15m
po/nginx-controller-28vq2                                    0/1       ImagePullBackOff   0          15m
po/nginx-controller-6hqkf                                    0/1       ImagePullBackOff   0          15m
po/nginx-controller-955c0340244fe5573ff7f61adfa2cb57-h98mk   0/1       ImagePullBackOff   0          15m
po/nginx-controller-bf0v5                                    0/1       ImagePullBackOff   0          15m
po/nginx-controller-hjql2                                    0/1       ImagePullBackOff   0          15m
po/nginx-controller-lk3q0                                    0/1       ImagePullBackOff   0          15m
po/nginx-controller-nqkbd                                    0/1       ImagePullBackOff   0          15m
po/nginx-controller-q0jgv                                    0/1       ImagePullBackOff   0          15m
po/nginx-controller-zr4p2                                    0/1       ImagePullBackOff   0          15m

NAME                                                   DESIRED   CURRENT   READY     AGE
rc/nginx-controller                                    8         8         0         15m
rc/nginx-controller-955c0340244fe5573ff7f61adfa2cb57   1         1         0         15m
rc/nginx2-controller                                   0         0         0         15m
```

### TBD
To be determined
```
fanhonglingdeMacBook-Pro:kubernetes-examples fanhongling$ mvn clean exec:java -Dexec.mainClass="io.fabric8.kubernetes.examples.ListExamples" -Dexec.args="https://172.17.4.200:6443" versions:set -DnewVersion=2.2.14
```
