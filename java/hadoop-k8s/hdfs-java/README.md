
Inspired

1. https://github.com/saagie/example-java-read-and-write-from-hdfs/blob/master/src/main/java/io/saagie/example/hdfs/Main.java


Reference

1. http://hadoop.apache.org/docs/current/hadoop-project-dist/hadoop-common/FileSystemShell.html
1. http://hadoop.apache.org/docs/current/api/index.html

__CI/CD__

From top POM directory
```
fanhonglingdeMacBook-Pro:hadoop-k8s fanhongling$ mvn clean package spring-boot:repackage 
[INFO] Scanning for projects...
Downloading from aliyun: http://maven.aliyun.com/nexus/content/groups/public/io/fabric8/docker-maven-plugin/0.24.0/docker-maven-plugin-0.24.0.pom
Downloaded from aliyun: http://maven.aliyun.com/nexus/content/groups/public/io/fabric8/docker-maven-plugin/0.24.0/docker-maven-plugin-0.24.0.pom (0 B at 0 B/s)
### snip...
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Build Order:
[INFO] 
[INFO] hadoop Kubernetes :: Applications
[INFO] hdfs-java
[INFO] hdfs-springboot
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building hadoop Kubernetes :: Applications 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.6.1:clean (default-clean) @ hadoop-Kubernetes ---
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.5.10.RELEASE:repackage (default-cli) @ hadoop-Kubernetes ---
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building hdfs-java latest
[INFO] ------------------------------------------------------------------------
Downloading from aliyun: http://maven.aliyun.com/nexus/content/groups/public/com/spotify/dockerfile-maven-plugin/1.3.7/dockerfile-maven-plugin-1.3.7.pom
Downloaded from aliyun: http://maven.aliyun.com/nexus/content/groups/public/com/spotify/dockerfile-maven-plugin/1.3.7/dockerfile-maven-plugin-1.3.7.pom (0 B at 0 B/s)
### snip...
[INFO] 
[INFO] --- maven-clean-plugin:2.6.1:clean (default-clean) @ hdfs-java ---
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:resources (default-resources) @ hdfs-java ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/src/main/resources
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/src/main/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ hdfs-java ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 1 source file to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/target/classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
### snip... Because ERRORs
```

Succeeded
```
fanhonglingdeMacBook-Pro:hadoop-k8s fanhongling$ mvn clean package spring-boot:repackage 
[INFO] Scanning for projects...
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Build Order:
[INFO] 
[INFO] hadoop Kubernetes :: Applications
[INFO] hdfs-java
[INFO] hdfs-springboot
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building hadoop Kubernetes :: Applications 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.6.1:clean (default-clean) @ hadoop-Kubernetes ---
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.5.10.RELEASE:repackage (default-cli) @ hadoop-Kubernetes ---
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building hdfs-java latest
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.6.1:clean (default-clean) @ hdfs-java ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/target
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:resources (default-resources) @ hdfs-java ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/src/main/resources
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/src/main/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ hdfs-java ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 1 source file to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/target/classes
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:testResources (default-testResources) @ hdfs-java ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:testCompile (default-testCompile) @ hdfs-java ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-surefire-plugin:2.18.1:test (default-test) @ hdfs-java ---
[INFO] No tests to run.
[INFO] 
[INFO] --- maven-jar-plugin:2.6:jar (default-jar) @ hdfs-java ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/target/hdfs-java-latest.jar
[INFO] 
[INFO] --- maven-assembly-plugin:2.6:single (make-assembly) @ hdfs-java ---
[INFO] artifact net.minidev:json-smart: checking for updates from aliyun
[INFO] artifact net.minidev:json-smart: checking for updates from clojars
[INFO] artifact net.minidev:json-smart: checking for updates from sonatype-snapshots
[INFO] artifact net.minidev:json-smart: checking for updates from central
[INFO] artifact net.minidev:json-smart: checking for updates from dynamodb-local-oregon
[INFO] artifact net.minidev:json-smart: checking for updates from apache.snapshots.https
[INFO] artifact net.minidev:json-smart: checking for updates from repository.jboss.org
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-java/target/hdfs-java-latest-jar-with-dependencies.jar
[INFO] 
[INFO] --- dockerfile-maven-plugin:1.3.7:build (default) @ hdfs-java ---
Downloading from aliyun: http://maven.aliyun.com/nexus/content/groups/public/com/spotify/docker-client/8.8.1/docker-client-8.8.1.pom
Downloaded from aliyun: http://maven.aliyun.com/nexus/content/groups/public/com/spotify/docker-client/8.8.1/docker-client-8.8.1.pom (0 B at 0 B/s)
### snip... many downloads
[INFO] Skipping execution because 'dockerfile.skip' is set
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.5.10.RELEASE:repackage (default-cli) @ hdfs-java ---
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building hdfs-springboot latest
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.6.1:clean (default-clean) @ hdfs-springboot ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-springboot/target
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:resources (default-resources) @ hdfs-springboot ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 1 resource
[INFO] Copying 1 resource
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:compile (default-compile) @ hdfs-springboot ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 1 source file to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-springboot/target/classes
[INFO] 
[INFO] --- maven-resources-plugin:3.0.2:testResources (default-testResources) @ hdfs-springboot ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-springboot/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:testCompile (default-testCompile) @ hdfs-springboot ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-surefire-plugin:2.18.1:test (default-test) @ hdfs-springboot ---
[INFO] No tests to run.
[INFO] 
[INFO] --- maven-jar-plugin:2.6:jar (default-jar) @ hdfs-springboot ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/hadoop-k8s/hdfs-springboot/target/hdfs-springboot-latest.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.5.10.RELEASE:repackage (default) @ hdfs-springboot ---
[INFO] 
[INFO] --- dockerfile-maven-plugin:1.3.7:build (default) @ hdfs-springboot ---
[INFO] Skipping execution because 'dockerfile.skip' is set
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.5.10.RELEASE:repackage (default-cli) @ hdfs-springboot ---
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Summary:
[INFO] 
[INFO] hadoop Kubernetes :: Applications .................. SUCCESS [  0.783 s]
[INFO] hdfs-java .......................................... SUCCESS [ 25.590 s]
[INFO] hdfs-springboot .................................... SUCCESS [  2.909 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 31.476 s
[INFO] Finished at: 2018-02-08T13:55:41-08:00
[INFO] Final Memory: 188M/455M
[INFO] ------------------------------------------------------------------------
```