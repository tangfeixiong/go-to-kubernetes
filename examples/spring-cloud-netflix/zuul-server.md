### Spring Cloud Netflix - Zuul Server

Maven build
```
[vagrant@localhost zuul-server]$ mvn install spring-boot:repackage
[INFO] Scanning for projects...
...snip...
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ zuul-server ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 1 resource
[INFO] Copying 1 resource
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ zuul-server ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 1 source file to /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/zuul-server/target/classes
[INFO] 
[INFO] --- maven-resources-plugin:2.6:testResources (default-testResources) @ zuul-server ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/zuul-server/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:testCompile (default-testCompile) @ zuul-server ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-surefire-plugin:2.18.1:test (default-test) @ zuul-server ---
[INFO] No tests to run.
[INFO] 
[INFO] --- maven-jar-plugin:2.6:jar (default-jar) @ zuul-server ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/zuul-server/target/zuul-server-1.0.0.BUILD-SNAPSHOT.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.4.1.RELEASE:repackage (default) @ zuul-server ---
...snip...
[INFO] 
[INFO] --- maven-install-plugin:2.5.2:install (default-install) @ zuul-server ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/zuul-server/target/zuul-server-1.0.0.BUILD-SNAPSHOT.jar to /home/vagrant/.m2/repository/org/demo/zuul-server/1.0.0.BUILD-SNAPSHOT/zuul-server-1.0.0.BUILD-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/zuul-server/pom.xml to /home/vagrant/.m2/repository/org/demo/zuul-server/1.0.0.BUILD-SNAPSHOT/zuul-server-1.0.0.BUILD-SNAPSHOT.pom
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.4.1.RELEASE:repackage (default-cli) @ zuul-server ---
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 01:52 min
[INFO] Finished at: 2017-07-21T02:47:39+00:00
[INFO] Final Memory: 38M/236M
[INFO] ------------------------------------------------------------------------
```