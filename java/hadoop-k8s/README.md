## Development

Hadoop HDFS for Kubernetes

1. hdfs java
1. hdfs springboot
1. hdfs spring cloud kubernetes

## CI/CD

### HDFS SpringBoot

Generate _jar_ for instance, Maven should shows if everything have been downloaded)
```
fanhonglingdeMacBook-Pro:hadoop-k8s fanhongling$ mvn clean package spring-boot:repackage 
[INFO] Scanning for projects...
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Build Order:
[INFO] 
[INFO] hadoop Kubernetes :: Applications
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
[INFO] hadoop Kubernetes :: Applications .................. SUCCESS [  0.622 s]
[INFO] hdfs-springboot .................................... SUCCESS [  4.392 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 6.918 s
[INFO] Finished at: 2018-02-08T12:30:06-08:00
[INFO] Final Memory: 45M/316M
[INFO] ------------------------------------------------------------------------
```

Run in Windows


Build docker and deploy into Kubernetes


## Reference

### Maven-Wrapper

Refer to https://github.com/takari/maven-wrapper

```
fanhonglingdeMacBook-Pro:hadoop-k8s fanhongling$ mvn -N io.takari:maven:wrapper -Dmaven=3.5.2
```

```
fanhonglingdeMacBook-Pro:hadoop-k8s fanhongling$ mvn -N io.takari:maven:wrapper -Dmaven=3.3.9
```

### Spring-Hadoop

Refer to http://projects.spring.io/spring-hadoop/

Samples repo
```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/spring-projects$ git clone https://github.com/spring-projects/spring-hadoop-samples spring-hadoop-samples
Cloning into 'spring-hadoop-samples'...
remote: Counting objects: 1699, done.
remote: Total 1699 (delta 0), reused 0 (delta 0), pack-reused 1699
Receiving objects: 100% (1699/1699), 3.20 MiB | 362.00 KiB/s, done.
Resolving deltas: 100% (506/506), done.
Checking out files: 100% (239/239), done.
```

Project repo
```
ubuntu@rookdev-172-17-4-61:/Users/fanhongling/Downloads/workspace/src/github.com/spring-projects$ git clone https://github.com/spring-projects/spring-hadoop spring-hadoop
Cloning into 'spring-hadoop'...
remote: Counting objects: 27018, done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 27018 (delta 0), reused 0 (delta 0), pack-reused 27015
Receiving objects: 100% (27018/27018), 7.69 MiB | 234.00 KiB/s, done.
Resolving deltas: 100% (10283/10283), done.
Checking out files: 100% (1321/1321), done.
```
