### Development Reference

[Kafka APIs](https://kafka.apache.org/documentation/#api)

[Kafka Connect](https://kafka.apache.org/documentation.html#connect)

[Kafka Streams](https://kafka.apache.org/documentation/streams#streams_developer)

### Kafka Http Proxy (Rest API)


```
fanhonglingdeMacBook-Pro:github.com fanhongling$ git clone https://github.com/confluentinc/kafka-rest confluentinc/kafka-rest
Cloning into 'confluentinc/kafka-rest'...
remote: Counting objects: 4953, done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 4953 (delta 0), reused 0 (delta 0), pack-reused 4951
Receiving objects: 100% (4953/4953), 1.05 MiB | 394.00 KiB/s, done.
Resolving deltas: 100% (2532/2532), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:github.com fanhongling$ cd confluentinc/
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ git clone https://github.com/confluentinc/rest-utils rest-utils
Cloning into 'rest-utils'...
remote: Counting objects: 1745, done.
remote: Total 1745 (delta 0), reused 0 (delta 0), pack-reused 1744
Receiving objects: 100% (1745/1745), 288.01 KiB | 42.00 KiB/s, done.
Resolving deltas: 100% (628/628), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ git clone https://github.com/confluentinc/common common
Cloning into 'common'...
remote: Counting objects: 1404, done.
remote: Compressing objects: 100% (4/4), done.
remote: Total 1404 (delta 0), reused 0 (delta 0), pack-reused 1400
Receiving objects: 100% (1404/1404), 225.70 KiB | 74.00 KiB/s, done.
Resolving deltas: 100% (510/510), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ cd common 
fanhonglingdeMacBook-Pro:common fanhongling$ mvn install -DskipTests
[INFO] Scanning for projects...
Downloading: https://repo.maven.apache.org/maven2/kr/motd/maven/os-maven-plugin/1.5.0.Final/os-maven-plugin-1.5.0.Final.pom
Downloaded: https://repo.maven.apache.org/maven2/kr/motd/maven/os-maven-plugin/1.5.0.Final/os-maven-plugin-1.5.0.Final.pom (7 KB at 2.6 KB/sec)
Downloading: https://repo.maven.apache.org/maven2/kr/motd/maven/os-maven-plugin/1.5.0.Final/os-maven-plugin-1.5.0.Final.jar
Downloaded: https://repo.maven.apache.org/maven2/kr/motd/maven/os-maven-plugin/1.5.0.Final/os-maven-plugin-1.5.0.Final.jar (29 KB at 38.0 KB/sec)
[INFO] ------------------------------------------------------------------------
[INFO] Detecting the operating system and CPU architecture
[INFO] ------------------------------------------------------------------------
[INFO] os.detected.name: osx
[INFO] os.detected.arch: x86_64
[INFO] os.detected.version: 10.10
[INFO] os.detected.version.major: 10
[INFO] os.detected.version.minor: 10
[INFO] os.detected.classifier: osx-x86_64
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Build Order:
[INFO] 
[INFO] Build Tools
[INFO] common
[INFO] utils
[INFO] metrics
[INFO] config
[INFO] package
[INFO] Common
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Build Tools 3.3.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------

[INFO] --- maven-install-plugin:2.4:install (default-install) @ common ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/common/parent/pom.xml to /Users/fanhongling/.m2/repository/io/confluent/common/3.3.0-SNAPSHOT/common-3.3.0-SNAPSHOT.pom
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/common/parent/target/common-3.3.0-SNAPSHOT-tests.jar to /Users/fanhongling/.m2/repository/io/confluent/common/3.3.0-SNAPSHOT/common-3.3.0-SNAPSHOT-tests.jar
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Summary:
[INFO] 
[INFO] Build Tools ........................................ SUCCESS [  1.368 s]
[INFO] common ............................................. SUCCESS [ 13.664 s]
[INFO] utils .............................................. SUCCESS [  7.835 s]
[INFO] metrics ............................................ SUCCESS [  5.251 s]
[INFO] config ............................................. SUCCESS [  1.331 s]
[INFO] package ............................................ SUCCESS [  6.071 s]
[INFO] Common ............................................. SUCCESS [  0.511 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 40.044 s
[INFO] Finished at: 2017-05-10T10:34:44-07:00
[INFO] Final Memory: 32M/323M
[INFO] ------------------------------------------------------------------------
fanhonglingdeMacBook-Pro:common fanhongling$ cd ..
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ git clone https://github.com/confluentinc/kafka kafka
Cloning into 'kafka'...
remote: Counting objects: 99174, done.
remote: Compressing objects: 100% (42/42), done.
remote: Total 99174 (delta 21), reused 13 (delta 13), pack-reused 99119
Receiving objects: 100% (99174/99174), 49.51 MiB | 419.00 KiB/s, done.
Resolving deltas: 100% (46583/46583), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ cd kafka
fanhonglingdeMacBook-Pro:kafka fanhongling$ gradle
Starting a Gradle Daemon (subsequent builds will be faster)
Download https://jcenter.bintray.com/org/ajoberstar/grgit/1.9.2/grgit-1.9.2.pom
Download https://jcenter.bintray.com/com/github/ben-manes/gradle-versions-plugin/0.14.0/gradle-versions-plugin-0.14.0.pom
Download https://repo1.maven.org/maven2/org/scoverage/gradle-scoverage/2.1.0/gradle-scoverage-2.1.0.pom
Download https://jcenter.bintray.com/com/github/jengelman/gradle/plugins/shadow/1.2.4/shadow-1.2.4.pom
Download https://repo1.maven.org/maven2/org/eclipse/jgit/org.eclipse.jgit/4.5.2.201704071617-r/org.eclipse.jgit-4.5.2.201704071617-r.pom
Download https://repo1.maven.org/maven2/org/eclipse/jgit/org.eclipse.jgit-parent/4.5.2.201704071617-r/org.eclipse.jgit-parent-4.5.2.201704071617Download https://repo1.maven.org/maven2/org/eclipse/jgit/org.eclipse.jgit.ui/4.5.2.201704071617-r/org.eclipse.jgit.ui-4.5.2.201704071617-r.pom
Download https://repo1.maven.org/maven2/com/jcraft/jsch/0.1.54/jsch-0.1.54.pom
Download https://repo1.maven.org/maven2/org/slf4j/slf4j-api/1.7.25/slf4j-api-1.7.25.pom
Download https://repo1.maven.org/maven2/org/slf4j/slf4j-parent/1.7.25/slf4j-parent-1.7.25.pom
Download https://repo1.maven.org/maven2/org/jdom/jdom2/2.0.5/jdom2-2.0.5.pom
Download https://repo1.maven.org/maven2/org/codehaus/groovy/groovy-backports-compat23/2.4.4/groovy-backports-compat23-2.4.4.pom
Download https://jcenter.bintray.com/org/ajoberstar/grgit/1.9.2/grgit-1.9.2.jar
Download https://jcenter.bintray.com/com/github/ben-manes/gradle-versions-plugin/0.14.0/gradle-versions-plugin-0.14.0.jar
Download https://repo1.maven.org/maven2/org/scoverage/gradle-scoverage/2.1.0/gradle-scoverage-2.1.0.jar
Download https://jcenter.bintray.com/com/github/jengelman/gradle/plugins/shadow/1.2.4/shadow-1.2.4.jar
Download https://repo1.maven.org/maven2/org/eclipse/jgit/org.eclipse.jgit/4.5.2.201704071617-r/org.eclipse.jgit-4.5.2.201704071617-r.jar
Download https://repo1.maven.org/maven2/org/eclipse/jgit/org.eclipse.jgit.ui/4.5.2.201704071617-r/org.eclipse.jgit.ui-4.5.2.201704071617-r.jar
Download https://repo1.maven.org/maven2/com/jcraft/jsch/0.1.54/jsch-0.1.54.jar
Download https://repo1.maven.org/maven2/org/slf4j/slf4j-api/1.7.25/slf4j-api-1.7.25.jar
Download https://repo1.maven.org/maven2/org/jdom/jdom2/2.0.5/jdom2-2.0.5.jar
Download https://repo1.maven.org/maven2/org/codehaus/groovy/groovy-backports-compat23/2.4.4/groovy-backports-compat23-2.4.4.jar
Exception caught during execution of command '[git, config, --system, --edit]' in '/usr/bin', return code '128', error message 'fatal: Could not switch to '/Library/Developer/CommandLineTools/usr/etc/': No such file or directory
'
Building project 'core' with Scala version 2.10.6
Download https://repo1.maven.org/maven2/org/scoverage/scalac-scoverage-plugin_2.10/1.3.0/scalac-scoverage-plugin_2.10-1.3.0.pom
Download https://repo1.maven.org/maven2/org/scoverage/scalac-scoverage-runtime_2.10/1.3.0/scalac-scoverage-runtime_2.10-1.3.0.pom
Download https://repo1.maven.org/maven2/org/scoverage/scalac-scoverage-plugin_2.10/1.3.0/scalac-scoverage-plugin_2.10-1.3.0.jar
Download https://repo1.maven.org/maven2/org/scoverage/scalac-scoverage-runtime_2.10/1.3.0/scalac-scoverage-runtime_2.10-1.3.0.jar
:downloadWrapper

BUILD SUCCESSFUL

Total time: 48.278 secs
fanhonglingdeMacBook-Pro:kafka fanhongling$ ./gradlew installAll
Downloading https://services.gradle.org/distributions/gradle-3.5-bin.zip
..................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................
Unzipping /Users/fanhongling/.gradle/wrapper/dists/gradle-3.5-bin/daoimhu7k5rlo48ntmxw2ok3e/gradle-3.5-bin.zip to /Users/fanhongling/.gradle/wrapper/dists/gradle-3.5-bin/daoimhu7k5rlo48ntmxw2ok3e
Set executable permissions for: /Users/fanhongling/.gradle/wrapper/dists/gradle-3.5-bin/daoimhu7k5rlo48ntmxw2ok3e/gradle-3.5/bin/gradle

BUILD SUCCESSFUL

Total time: 18 mins 30.698 secs

fanhonglingdeMacBook-Pro:common fanhongling$ cd ../rest-utils/
fanhonglingdeMacBook-Pro:rest-utils fanhongling$ mvn install -DskipTests
[INFO] Scanning for projects...
Downloading: http://packages.confluent.io/maven/org/glassfish/jersey/jersey-bom/2.25/jersey-bom-2.25.pom
Downloading: https://repo.maven.apache.org/maven2/org/glassfish/jersey/jersey-bom/2.25/jersey-bom-2.25.pom
Downloaded: https://repo.maven.apache.org/maven2/org/glassfish/jersey/jersey-bom/2.25/jersey-bom-2.25.pom (22 KB at 10.0 KB/sec)
[INFO] ------------------------------------------------------------------------
[INFO] Detecting the operating system and CPU architecture
[INFO] ------------------------------------------------------------------------
[INFO] os.detected.name: osx
[INFO] os.detected.arch: x86_64
[INFO] os.detected.version: 10.10
[INFO] os.detected.version.major: 10
[INFO] os.detected.version.minor: 10
[INFO] os.detected.classifier: osx-x86_64
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Build Order:
[INFO] 
[INFO] rest-utils-parent
[INFO] rest-utils
[INFO] rest-utils-test
[INFO] rest-utils-example
[INFO] rest-utils-package
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building rest-utils-parent 3.3.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------

[INFO] --- maven-install-plugin:2.4:install (default-install) @ rest-utils-package ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/rest-utils/package/pom.xml to /Users/fanhongling/.m2/repository/io/confluent/rest-utils-package/3.3.0-SNAPSHOT/rest-utils-package-3.3.0-SNAPSHOT.pom
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/rest-utils/package/target/rest-utils-package-3.3.0-SNAPSHOT-tests.jar to /Users/fanhongling/.m2/repository/io/confluent/rest-utils-package/3.3.0-SNAPSHOT/rest-utils-package-3.3.0-SNAPSHOT-tests.jar
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Summary:
[INFO] 
[INFO] rest-utils-parent .................................. SUCCESS [  4.120 s]
[INFO] rest-utils ......................................... SUCCESS [ 23.467 s]
[INFO] rest-utils-test .................................... SUCCESS [  9.967 s]
[INFO] rest-utils-example ................................. SUCCESS [  0.637 s]
[INFO] rest-utils-package ................................. SUCCESS [  3.514 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 42.550 s
[INFO] Finished at: 2017-05-10T11:12:44-07:00
[INFO] Final Memory: 39M/319M
[INFO] ------------------------------------------------------------------------
fanhonglingdeMacBook-Pro:kafka-rest fanhongling$ cd ..
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ git clone https://github.com/confluentinc/schema-registry schema-registry
Cloning into 'schema-registry'...
remote: Counting objects: 8820, done.
remote: Compressing objects: 100% (75/75), done.
remote: Total 8820 (delta 28), reused 0 (delta 0), pack-reused 8730
Receiving objects: 100% (8820/8820), 1.75 MiB | 223.00 KiB/s, done.
Resolving deltas: 100% (3458/3458), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ cd schema-registry/
fanhonglingdeMacBook-Pro:schema-registry fanhongling$ mvn install -DskipTests
[INFO] Scanning for projects...
[INFO] ------------------------------------------------------------------------
[INFO] Detecting the operating system and CPU architecture
[INFO] ------------------------------------------------------------------------
[INFO] os.detected.name: osx
[INFO] os.detected.arch: x86_64
[INFO] os.detected.version: 10.10
[INFO] os.detected.version.major: 10
[INFO] os.detected.version.minor: 10
[INFO] os.detected.classifier: osx-x86_64
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Build Order:
[INFO] 
[INFO] kafka-schema-registry-parent
[INFO] kafka-schema-registry-client
[INFO] kafka-avro-serializer
[INFO] kafka-schema-registry
[INFO] kafka-json-serializer
[INFO] kafka-connect-avro-converter
[INFO] kafka-schema-registry-package
[INFO] kafka-serde-tools-package
[INFO] maven-plugin
[INFO] kafka-streams-avro-serde
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building kafka-schema-registry-parent 3.3.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------

[INFO] --- maven-install-plugin:2.4:install (default-install) @ kafka-streams-avro-serde ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/schema-registry/avro-serde/target/kafka-streams-avro-serde-3.3.0-SNAPSHOT.jar to /Users/fanhongling/.m2/repository/io/confluent/kafka-streams-avro-serde/3.3.0-SNAPSHOT/kafka-streams-avro-serde-3.3.0-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/schema-registry/avro-serde/pom.xml to /Users/fanhongling/.m2/repository/io/confluent/kafka-streams-avro-serde/3.3.0-SNAPSHOT/kafka-streams-avro-serde-3.3.0-SNAPSHOT.pom
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/schema-registry/avro-serde/target/kafka-streams-avro-serde-3.3.0-SNAPSHOT-tests.jar to /Users/fanhongling/.m2/repository/io/confluent/kafka-streams-avro-serde/3.3.0-SNAPSHOT/kafka-streams-avro-serde-3.3.0-SNAPSHOT-tests.jar
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Summary:
[INFO] 
[INFO] kafka-schema-registry-parent ....................... SUCCESS [  4.240 s]
[INFO] kafka-schema-registry-client ....................... SUCCESS [ 19.446 s]
[INFO] kafka-avro-serializer .............................. SUCCESS [  7.491 s]
[INFO] kafka-schema-registry .............................. SUCCESS [01:54 min]
[INFO] kafka-json-serializer .............................. SUCCESS [  0.838 s]
[INFO] kafka-connect-avro-converter ....................... SUCCESS [ 48.302 s]
[INFO] kafka-schema-registry-package ...................... SUCCESS [  4.251 s]
[INFO] kafka-serde-tools-package .......................... SUCCESS [  9.658 s]
[INFO] maven-plugin ....................................... SUCCESS [ 11.898 s]
[INFO] kafka-streams-avro-serde ........................... SUCCESS [ 14.063 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 03:55 min
[INFO] Finished at: 2017-05-10T11:23:40-07:00
[INFO] Final Memory: 62M/326M
[INFO] ------------------------------------------------------------------------
fanhonglingdeMacBook-Pro:github.com fanhongling$ cd ../kafka-rest/
fanhonglingdeMacBook-Pro:kafka-rest fanhongling$ mvn install -DskipTests
fanhonglingdeMacBook-Pro:kafka-rest fanhongling$ mvn install -DskipTests
[INFO] Scanning for projects...
[WARNING] 
[WARNING] Some problems were encountered while building the effective model for io.confluent:kafka-rest:jar:3.3.0-SNAPSHOT
[WARNING] The expression ${version} is deprecated. Please use ${project.version} instead.
[WARNING] 
[WARNING] It is highly recommended to fix these problems because they threaten the stability of your build.
[WARNING] 
[WARNING] For this reason, future Maven versions might no longer support building such malformed projects.
[WARNING] 
[INFO] ------------------------------------------------------------------------
[INFO] Detecting the operating system and CPU architecture
[INFO] ------------------------------------------------------------------------
[INFO] os.detected.name: osx
[INFO] os.detected.arch: x86_64
[INFO] os.detected.version: 10.10
[INFO] os.detected.version.major: 10
[INFO] os.detected.version.minor: 10
[INFO] os.detected.classifier: osx-x86_64
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building kafka-rest 3.3.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------

[INFO] --- maven-install-plugin:2.4:install (default-install) @ kafka-rest ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/kafka-rest/target/kafka-rest-3.3.0-SNAPSHOT.jar to /Users/fanhongling/.m2/repository/io/confluent/kafka-rest/3.3.0-SNAPSHOT/kafka-rest-3.3.0-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/kafka-rest/pom.xml to /Users/fanhongling/.m2/repository/io/confluent/kafka-rest/3.3.0-SNAPSHOT/kafka-rest-3.3.0-SNAPSHOT.pom
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/kafka-rest/target/kafka-rest-3.3.0-SNAPSHOT-tests.jar to /Users/fanhongling/.m2/repository/io/confluent/kafka-rest/3.3.0-SNAPSHOT/kafka-rest-3.3.0-SNAPSHOT-tests.jar
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 17.780 s
[INFO] Finished at: 2017-05-10T11:24:32-07:00
[INFO] Final Memory: 37M/286M
[INFO] ------------------------------------------------------------------------
```

For _kafka connect_
```
fanhonglingdeMacBook-Pro:kafka-rest fanhongling$ cd ..
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ git clone https://github.com/confluentinc/kafka-connect-storage-common kafka-connect-storage-common
Cloning into 'kafka-connect-storage-common'...
remote: Counting objects: 1975, done.
remote: Compressing objects: 100% (87/87), done.
remote: Total 1975 (delta 37), reused 1 (delta 1), pack-reused 1839
Receiving objects: 100% (1975/1975), 313.41 KiB | 81.00 KiB/s, done.
Resolving deltas: 100% (690/690), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:kafka-connect-storage-common fanhongling$ mvn install -DskipTests
[INFO] Scanning for projects...
[INFO] ------------------------------------------------------------------------
[INFO] Detecting the operating system and CPU architecture
[INFO] ------------------------------------------------------------------------
[INFO] os.detected.name: osx
[INFO] os.detected.arch: x86_64
[INFO] os.detected.version: 10.10
[INFO] os.detected.version.major: 10
[INFO] os.detected.version.minor: 10
[INFO] os.detected.classifier: osx-x86_64
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Build Order:
[INFO] 
[INFO] kafka-connect-storage-common-parent
[INFO] kafka-connect-storage-common
[INFO] kafka-connect-storage-core
[INFO] kafka-connect-storage-wal
[INFO] kafka-connect-storage-partitioner
[INFO] kafka-connect-storage-hive
[INFO] kafka-connect-storage-format
[INFO] kafka-connect-storage-common-package
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building kafka-connect-storage-common-parent 3.3.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------

[INFO] ------------------------------------------------------------------------
[INFO] Reactor Summary:
[INFO] 
[INFO] kafka-connect-storage-common-parent ................ SUCCESS [  3.989 s]
[INFO] kafka-connect-storage-common ....................... SUCCESS [  3.183 s]
[INFO] kafka-connect-storage-core ......................... SUCCESS [  9.024 s]
[INFO] kafka-connect-storage-wal .......................... SUCCESS [  1.246 s]
[INFO] kafka-connect-storage-partitioner .................. SUCCESS [  1.282 s]
[INFO] kafka-connect-storage-hive ......................... SUCCESS [06:16 min]
[INFO] kafka-connect-storage-format ....................... SUCCESS [ 27.888 s]
[INFO] kafka-connect-storage-common-package ............... SUCCESS [ 14.950 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 07:19 min
[INFO] Finished at: 2017-05-10T11:34:06-07:00
[INFO] Final Memory: 51M/313M
[INFO] ------------------------------------------------------------------------
fanhonglingdeMacBook-Pro:kafka-connect-storage-common fanhongling$ cd ..
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ git clone https://github.com/confluentinc/kafka-connect-jdbc kafka-connect-jdbc
Cloning into 'kafka-connect-jdbc'...
remote: Counting objects: 8555, done.
remote: Compressing objects: 100% (25/25), done.
remote: Total 8555 (delta 9), reused 0 (delta 0), pack-reused 8530
Receiving objects: 100% (8555/8555), 11.42 MiB | 221.00 KiB/s, done.
Resolving deltas: 100% (3598/3598), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ cd kafka-connect-jdbc
fanhonglingdeMacBook-Pro:kafka-connect-jdbc fanhongling$ mvn install -DskipTests
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building kafka-connect-jdbc 3.3.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------

[INFO] --- maven-install-plugin:2.4:install (default-install) @ kafka-connect-jdbc ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/kafka-connect-jdbc/target/kafka-connect-jdbc-3.3.0-SNAPSHOT.jar to /Users/fanhongling/.m2/repository/io/confluent/kafka-connect-jdbc/3.3.0-SNAPSHOT/kafka-connect-jdbc-3.3.0-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/kafka-connect-jdbc/pom.xml to /Users/fanhongling/.m2/repository/io/confluent/kafka-connect-jdbc/3.3.0-SNAPSHOT/kafka-connect-jdbc-3.3.0-SNAPSHOT.pom
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 40.331 s
[INFO] Finished at: 2017-05-10T11:37:23-07:00
[INFO] Final Memory: 29M/323M
[INFO] ------------------------------------------------------------------------
fanhonglingdeMacBook-Pro:kafka-connect-jdbc fanhongling$ cd ..
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ git clone https://github.com/confluentinc/kafka-connect-hdfs kafka-connect-hdfs
Cloning into 'kafka-connect-hdfs'...
remote: Counting objects: 2257, done.
remote: Total 2257 (delta 0), reused 0 (delta 0), pack-reused 2257
Receiving objects: 100% (2257/2257), 504.65 KiB | 407.00 KiB/s, done.
Resolving deltas: 100% (963/963), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ cd kafka-connect-hdfs/
fanhonglingdeMacBook-Pro:kafka-connect-hdfs fanhongling$ mvn install -DskipTests
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building kafka-connect-hdfs 3.3.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------

[INFO] --- maven-install-plugin:2.4:install (default-install) @ kafka-connect-hdfs ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/kafka-connect-hdfs/target/kafka-connect-hdfs-3.3.0-SNAPSHOT.jar to /Users/fanhongling/.m2/repository/io/confluent/kafka-connect-hdfs/3.3.0-SNAPSHOT/kafka-connect-hdfs-3.3.0-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/kafka-connect-hdfs/pom.xml to /Users/fanhongling/.m2/repository/io/confluent/kafka-connect-hdfs/3.3.0-SNAPSHOT/kafka-connect-hdfs-3.3.0-SNAPSHOT.pom
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 01:50 min
[INFO] Finished at: 2017-05-10T11:41:01-07:00
[INFO] Final Memory: 51M/315M
[INFO] ------------------------------------------------------------------------
fanhonglingdeMacBook-Pro:kafka-connect-hdfs fanhongling$ cd ..
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ git clone https://github.com/confluentinc/kafka-connect-elasticsearch kafka-connect-elasticsearch
Cloning into 'kafka-connect-elasticsearch'...
remote: Counting objects: 1213, done.
remote: Total 1213 (delta 0), reused 0 (delta 0), pack-reused 1213
Receiving objects: 100% (1213/1213), 213.35 KiB | 142.00 KiB/s, done.
Resolving deltas: 100% (476/476), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:confluentinc fanhongling$ cd kafka-connect-elasticsearch/
fanhonglingdeMacBook-Pro:kafka-connect-elasticsearch fanhongling$ mvn install -DskipTests
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building kafka-connect-elasticsearch 3.3.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------

[INFO] --- maven-install-plugin:2.4:install (default-install) @ kafka-connect-elasticsearch ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/kafka-connect-elasticsearch/target/kafka-connect-elasticsearch-3.3.0-SNAPSHOT.jar to /Users/fanhongling/.m2/repository/io/confluent/kafka-connect-elasticsearch/3.3.0-SNAPSHOT/kafka-connect-elasticsearch-3.3.0-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/confluentinc/kafka-connect-elasticsearch/pom.xml to /Users/fanhongling/.m2/repository/io/confluent/kafka-connect-elasticsearch/3.3.0-SNAPSHOT/kafka-connect-elasticsearch-3.3.0-SNAPSHOT.pom
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 01:33 min
[INFO] Finished at: 2017-05-10T11:44:21-07:00
[INFO] Final Memory: 34M/321M
[INFO] ------------------------------------------------------------------------
```

