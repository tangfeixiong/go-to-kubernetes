### Netflix OSS



## Service Discovery


### Client and Server - https://github.com/Netflix/eureka

Source repository
```
fanhonglingdeMacBook-Pro:github.com fanhongling$ git clone https://github.com/Netflix/eureka Netflix/eureka
Cloning into 'Netflix/eureka'...
remote: Counting objects: 50903, done.
remote: Total 50903 (delta 0), reused 0 (delta 0), pack-reused 50903
Receiving objects: 100% (50903/50903), 11.02 MiB | 60.00 KiB/s, done.
Resolving deltas: 100% (20012/20012), done.
Checking connectivity... done.
```

Build into gradle
```
fanhonglingdeMacBook-Pro:github.com fanhongling$ cd Netflix/eureka/
fanhonglingdeMacBook-Pro:eureka fanhongling$ gradle --version

------------------------------------------------------------
Gradle 3.5
------------------------------------------------------------

Build time:   2017-04-10 13:37:25 UTC
Revision:     b762622a185d59ce0cfc9cbc6ab5dd22469e18a6

Groovy:       2.4.10
Ant:          Apache Ant(TM) version 1.9.6 compiled on June 29 2015
JVM:          1.8.0_91 (Oracle Corporation 25.91-b14)
OS:           Mac OS X 10.10.4 x86_64

fanhonglingdeMacBook-Pro:eureka fanhongling$ gradle build
Starting a Gradle Daemon (subsequent builds will be faster)
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-extra-configurations-plugin/2.2.2/gradle-extra-configurations-plugin-2.2.2.pom
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-netflixoss-project-plugin/3.2.3/gradle-netflixoss-project-plugin-3.2.3.pom
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-contacts-plugin/3.0.1/gradle-contacts-plugin-3.0.1.pom
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-dependency-lock-plugin/4.2.0/gradle-dependency-lock-plugin-4.2.0.pom
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-info-plugin/3.0.3/gradle-info-plugin-3.0.3.pom
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-bintray-plugin/3.3.4/nebula-bintray-plugin-3.3.4.pom
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-core/3.0.1/nebula-core-3.0.1.pom
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-project-plugin/3.0.3/nebula-project-plugin-3.0.3.pom
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-publishing-plugin/4.4.4/nebula-publishing-plugin-4.4.4.pom
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-release-plugin/4.0.1/nebula-release-plugin-4.0.1.pom
Download https://jcenter.bintray.com/nl/javadude/gradle/plugins/license-gradle-plugin/0.11.0/license-gradle-plugin-0.11.0.pom
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-scm-plugin/2.2.0/gradle-scm-plugin-2.2.0.pom
Download https://jcenter.bintray.com/com/perforce/p4java-jfrog/2011.1.297684/p4java-jfrog-2011.1.297684.pom
Download https://jcenter.bintray.com/org/eclipse/jgit/org.eclipse.jgit/3.2.0.201312181205-r/org.eclipse.jgit-3.2.0.201312181205-r.pom
Download https://jcenter.bintray.com/org/eclipse/jgit/org.eclipse.jgit-parent/3.2.0.201312181205-r/org.eclipse.jgit-parent-3.2.0.201312181205-r.poDownload https://jcenter.bintray.com/com/jfrog/bintray/gradle/gradle-bintray-plugin/1.5/gradle-bintray-plugin-1.5.pom
Download https://jcenter.bintray.com/org/jfrog/buildinfo/build-info-extractor-gradle/4.0.0/build-info-extractor-gradle-4.0.0.pom
Download https://jcenter.bintray.com/org/ajoberstar/gradle-git/1.4.1/gradle-git-1.4.1.pom
Download https://jcenter.bintray.com/com/mycila/maven-license-plugin/maven-license-plugin/1.10.b1/maven-license-plugin-1.10.b1.pom
Download https://jcenter.bintray.com/com/mycila/parent-pom/5/parent-pom-5.pom
Download https://jcenter.bintray.com/org/codehaus/groovy/modules/http-builder/http-builder/0.7.2/http-builder-0.7.2.pom
Download https://jcenter.bintray.com/commons-io/commons-io/2.0.1/commons-io-2.0.1.pom
Download https://jcenter.bintray.com/org/apache/commons/commons-parent/15/commons-parent-15.pom
Download https://jcenter.bintray.com/org/apache/ivy/ivy/2.2.0/ivy-2.2.0.pom
Download https://jcenter.bintray.com/org/jfrog/buildinfo/build-info-extractor/2.5.4/build-info-extractor-2.5.4.pom
Download https://jcenter.bintray.com/org/ajoberstar/grgit/1.5.0/grgit-1.5.0.pom
Download https://jcenter.bintray.com/com/github/zafarkhaja/java-semver/0.9.0/java-semver-0.9.0.pom
Download https://jcenter.bintray.com/org/apache/maven/maven-plugin-api/3.0.1/maven-plugin-api-3.0.1.pom
Download https://jcenter.bintray.com/org/apache/maven/maven/3.0.1/maven-3.0.1.pom
Download https://jcenter.bintray.com/com/mycila/xmltool/xmltool/3.3/xmltool-3.3.pom
Download https://jcenter.bintray.com/org/apache/maven/maven-project/3.0-alpha-2/maven-project-3.0-alpha-2.pom
Download https://jcenter.bintray.com/org/apache/maven/maven/3.0-alpha-2/maven-3.0-alpha-2.pom
Download https://jcenter.bintray.com/org/jfrog/buildinfo/build-info-client/2.5.4/build-info-client-2.5.4.pom
Download https://jcenter.bintray.com/org/eclipse/jgit/org.eclipse.jgit.ui/4.1.1.201511131810-r/org.eclipse.jgit.ui-4.1.1.201511131810-r.pom
Download https://jcenter.bintray.com/org/eclipse/jgit/org.eclipse.jgit-parent/4.1.1.201511131810-r/org.eclipse.jgit-parent-4.1.1.201511131810-r.poDownload https://jcenter.bintray.com/org/apache/maven/maven-model/3.0.1/maven-model-3.0.1.pom
Download https://jcenter.bintray.com/org/apache/maven/maven-artifact/3.0.1/maven-artifact-3.0.1.pom
Download https://jcenter.bintray.com/org/sonatype/sisu/sisu-inject-plexus/1.4.3.1/sisu-inject-plexus-1.4.3.1.pom
Download https://jcenter.bintray.com/org/sonatype/sisu/inject/guice-plexus/1.4.3.1/guice-plexus-1.4.3.1.pom
Download https://jcenter.bintray.com/org/sonatype/sisu/inject/guice-bean/1.4.3.1/guice-bean-1.4.3.1.pom
Download https://jcenter.bintray.com/org/sonatype/sisu/sisu-inject/1.4.3.1/sisu-inject-1.4.3.1.pom
Download https://jcenter.bintray.com/org/sonatype/sisu/sisu-parent/1.4.3.1/sisu-parent-1.4.3.1.pom
Download https://jcenter.bintray.com/org/apache/maven/maven-compat/3.0-alpha-2/maven-compat-3.0-alpha-2.pom
Download https://jcenter.bintray.com/org/codehaus/plexus/plexus-container-default/1.0-beta-3.0.5/plexus-container-default-1.0-beta-3.0.5.pom
Download https://jcenter.bintray.com/org/codehaus/plexus/plexus-containers/1.0-beta-3.0.5/plexus-containers-1.0-beta-3.0.5.pom
Download https://jcenter.bintray.com/org/codehaus/woodstox/wstx-asl/3.2.6/wstx-asl-3.2.6.pom
Download https://jcenter.bintray.com/org/sonatype/spice/model-builder/1.3/model-builder-1.3.pom
Download https://jcenter.bintray.com/org/apache/maven/maven-project-builder/3.0-alpha-2/maven-project-builder-3.0-alpha-2.pom
Download https://jcenter.bintray.com/xerces/xercesImpl/2.9.1/xercesImpl-2.9.1.pom
Download https://jcenter.bintray.com/org/jfrog/buildinfo/build-info-api/2.5.4/build-info-api-2.5.4.pom
Download https://jcenter.bintray.com/org/sonatype/sisu/sisu-inject-bean/1.4.3.1/sisu-inject-bean-1.4.3.1.pom
Download https://jcenter.bintray.com/org/apache/maven/wagon/wagon-provider-api/1.0-beta-4/wagon-provider-api-1.0-beta-4.pom
Download https://jcenter.bintray.com/org/apache/maven/wagon/wagon/1.0-beta-4/wagon-1.0-beta-4.pom
Download https://jcenter.bintray.com/com/google/code/google-collections/google-collect/snapshot-20080530/google-collect-snapshot-20080530.pom
Download https://jcenter.bintray.com/org/sonatype/sisu/sisu-guice/2.9.1/sisu-guice-2.9.1.pom
Download https://jcenter.bintray.com/org/sonatype/sisu/inject/guice-parent/2.9.1/guice-parent-2.9.1.pom
Download https://jcenter.bintray.com/com/jcraft/jzlib/1.0.7/jzlib-1.0.7.pom
Download https://jcenter.bintray.com/org/eclipse/jgit/org.eclipse.jgit/4.1.1.201511131810-r/org.eclipse.jgit-4.1.1.201511131810-r.pom
Download https://jcenter.bintray.com/org/eclipse/jdt/org.eclipse.jdt.annotation/1.1.0/org.eclipse.jdt.annotation-1.1.0.pom
Download https://jcenter.bintray.com/commons-codec/commons-codec/1.8/commons-codec-1.8.pom
Download https://jcenter.bintray.com/org/apache/httpcomponents/httpcore/4.2.5/httpcore-4.2.5.pom
Download https://jcenter.bintray.com/org/apache/httpcomponents/httpcomponents-core/4.2.5/httpcomponents-core-4.2.5.pom
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-extra-configurations-plugin/2.2.2/gradle-extra-configurations-plugin-2.2.2.jar
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-netflixoss-project-plugin/3.2.3/gradle-netflixoss-project-plugin-3.2.3.jar
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-contacts-plugin/3.0.1/gradle-contacts-plugin-3.0.1.jar
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-dependency-lock-plugin/4.2.0/gradle-dependency-lock-plugin-4.2.0.jar
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-info-plugin/3.0.3/gradle-info-plugin-3.0.3.jar
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-bintray-plugin/3.3.4/nebula-bintray-plugin-3.3.4.jar
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-core/3.0.1/nebula-core-3.0.1.jar
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-project-plugin/3.0.3/nebula-project-plugin-3.0.3.jar
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-publishing-plugin/4.4.4/nebula-publishing-plugin-4.4.4.jar
Download https://jcenter.bintray.com/com/netflix/nebula/nebula-release-plugin/4.0.1/nebula-release-plugin-4.0.1.jar
Download https://jcenter.bintray.com/nl/javadude/gradle/plugins/license-gradle-plugin/0.11.0/license-gradle-plugin-0.11.0.jar
Download https://jcenter.bintray.com/com/netflix/nebula/gradle-scm-plugin/2.2.0/gradle-scm-plugin-2.2.0.jar
Download https://jcenter.bintray.com/com/perforce/p4java-jfrog/2011.1.297684/p4java-jfrog-2011.1.297684.jar
Download https://jcenter.bintray.com/com/jfrog/bintray/gradle/gradle-bintray-plugin/1.5/gradle-bintray-plugin-1.5.jar
Download https://jcenter.bintray.com/org/jfrog/buildinfo/build-info-extractor-gradle/4.0.0/build-info-extractor-gradle-4.0.0.jar
Download https://jcenter.bintray.com/org/ajoberstar/gradle-git/1.4.1/gradle-git-1.4.1.jar
Download https://jcenter.bintray.com/com/mycila/maven-license-plugin/maven-license-plugin/1.10.b1/maven-license-plugin-1.10.b1.jar
Download https://jcenter.bintray.com/org/codehaus/groovy/modules/http-builder/http-builder/0.7.2/http-builder-0.7.2.jar
Download https://jcenter.bintray.com/commons-io/commons-io/2.0.1/commons-io-2.0.1.jar
Download https://jcenter.bintray.com/org/apache/ivy/ivy/2.2.0/ivy-2.2.0.jar
Download https://jcenter.bintray.com/org/jfrog/buildinfo/build-info-extractor/2.5.4/build-info-extractor-2.5.4.jar
Download https://jcenter.bintray.com/org/ajoberstar/grgit/1.5.0/grgit-1.5.0.jar
Download https://jcenter.bintray.com/com/github/zafarkhaja/java-semver/0.9.0/java-semver-0.9.0.jar
Download https://jcenter.bintray.com/org/apache/maven/maven-plugin-api/3.0.1/maven-plugin-api-3.0.1.jar
Download https://jcenter.bintray.com/com/mycila/xmltool/xmltool/3.3/xmltool-3.3.jar
Download https://jcenter.bintray.com/org/apache/maven/maven-project/3.0-alpha-2/maven-project-3.0-alpha-2.jar
Download https://jcenter.bintray.com/org/jfrog/buildinfo/build-info-client/2.5.4/build-info-client-2.5.4.jar
Download https://jcenter.bintray.com/org/eclipse/jgit/org.eclipse.jgit.ui/4.1.1.201511131810-r/org.eclipse.jgit.ui-4.1.1.201511131810-r.jar
Download https://jcenter.bintray.com/org/apache/maven/maven-model/3.0.1/maven-model-3.0.1.jar
Download https://jcenter.bintray.com/org/apache/maven/maven-artifact/3.0.1/maven-artifact-3.0.1.jar
Download https://jcenter.bintray.com/org/sonatype/sisu/sisu-inject-plexus/1.4.3.1/sisu-inject-plexus-1.4.3.1.jar
Download https://jcenter.bintray.com/org/apache/maven/maven-compat/3.0-alpha-2/maven-compat-3.0-alpha-2.jar
Download https://jcenter.bintray.com/org/codehaus/plexus/plexus-container-default/1.0-beta-3.0.5/plexus-container-default-1.0-beta-3.0.5.jar
Download https://jcenter.bintray.com/org/codehaus/woodstox/wstx-asl/3.2.6/wstx-asl-3.2.6.jar
Download https://jcenter.bintray.com/org/sonatype/spice/model-builder/1.3/model-builder-1.3.jar
Download https://jcenter.bintray.com/org/apache/maven/maven-project-builder/3.0-alpha-2/maven-project-builder-3.0-alpha-2.jar
Download https://jcenter.bintray.com/xerces/xercesImpl/2.9.1/xercesImpl-2.9.1.jar
Download https://jcenter.bintray.com/org/jfrog/buildinfo/build-info-api/2.5.4/build-info-api-2.5.4.jar
Download https://jcenter.bintray.com/org/sonatype/sisu/sisu-inject-bean/1.4.3.1/sisu-inject-bean-1.4.3.1.jar
Download https://jcenter.bintray.com/org/apache/maven/wagon/wagon-provider-api/1.0-beta-4/wagon-provider-api-1.0-beta-4.jar
Download https://jcenter.bintray.com/com/google/code/google-collections/google-collect/snapshot-20080530/google-collect-snapshot-20080530.jar
Download https://jcenter.bintray.com/org/sonatype/sisu/sisu-guice/2.9.1/sisu-guice-2.9.1-noaop.jar
Download https://jcenter.bintray.com/com/jcraft/jzlib/1.0.7/jzlib-1.0.7.jar
Download https://jcenter.bintray.com/org/eclipse/jgit/org.eclipse.jgit/4.1.1.201511131810-r/org.eclipse.jgit-4.1.1.201511131810-r.jar
Download https://jcenter.bintray.com/org/eclipse/jdt/org.eclipse.jdt.annotation/1.1.0/org.eclipse.jdt.annotation-1.1.0.jar
Download https://jcenter.bintray.com/commons-codec/commons-codec/1.8/commons-codec-1.8.jar
Inferred project: eureka, version: 1.6.3-SNAPSHOT

FAILURE: Build failed with an exception.

* Where:
Build file '/Users/fanhongling/Downloads/workspace/src/github.com/Netflix/eureka/build.gradle' line: 50

* What went wrong:
A problem occurred evaluating root project 'eureka'.
> org/gradle/api/internal/project/AbstractProject

* Try:
Run with --stacktrace option to get the stack trace. Run with --info or --debug option to get more log output.

BUILD FAILED

Total time: 9 mins 38.097 secs
```

Using `gradlew` wrapper
```
fanhonglingdeMacBook-Pro:eureka fanhongling$ ./gradlew build
Downloading https://services.gradle.org/distributions/gradle-2.10-bin.zip
..............................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................
Unzipping /Users/fanhongling/.gradle/wrapper/dists/gradle-2.10-bin/baigpnfu14tdk6ztbfwcl8275/gradle-2.10-bin.zip to /Users/fanhongling/.gradle/wrapper/dists/gradle-2.10-bin/baigpnfu14tdk6ztbfwcl8275
Set executable permissions for: /Users/fanhongling/.gradle/wrapper/dists/gradle-2.10-bin/baigpnfu14tdk6ztbfwcl8275/gradle-2.10/bin/gradle
Inferred project: eureka, version: 1.6.3-SNAPSHOT
The testJar task is deprecated.  Please place common test harness code in its own project and publish separately.
The testJar task is deprecated.  Please place common test harness code in its own project and publish separately.
The testJar task is deprecated.  Please place common test harness code in its own project and publish separately.
Publication nebula not found in project :.
[buildinfo] Not using buildInfo properties file for this build.
Publication named 'nebula' does not exist for project ':' in task ':artifactoryPublish'.
:eureka-client:compileJava
Download https://jcenter.bintray.com/javax/ws/rs/jsr311-api/1.1.1/jsr311-api-1.1.1.pom
Download https://jcenter.bintray.com/com/sun/jersey/jersey-project/1.19.1/jersey-project-1.19.1.pom
Download https://jcenter.bintray.com/com/google/inject/guice/4.1.0/guice-4.1.0.pom
Download https://jcenter.bintray.com/com/google/inject/guice-parent/4.1.0/guice-parent-4.1.0.pom
Download https://jcenter.bintray.com/com/netflix/archaius/archaius-core/0.7.3/archaius-core-0.7.3.jar
Download https://jcenter.bintray.com/org/apache/httpcomponents/httpclient/4.3.4/httpclient-4.3.4.jar
Download https://jcenter.bintray.com/com/google/inject/guice/4.1.0/guice-4.1.0.jar
警告: [options] 未与 -source 1.7 一起设置引导类路径
注: 某些输入文件使用或覆盖了已过时的 API。
注: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
注: 某些输入文件使用了未经检查或不安全的操作。
注: 有关详细信息, 请使用 -Xlint:unchecked 重新编译。
1 个警告
:eureka-client:processResources UP-TO-DATE
:eureka-client:classes
:eureka-client:writeManifestProperties
:eureka-client:jar
:eureka-client:assemble
:eureka-client-archaius2:compileJava
Download https://jcenter.bintray.com/com/netflix/archaius/archaius2-core/2.1.7/archaius2-core-2.1.7.pom
Download https://jcenter.bintray.com/com/netflix/archaius/archaius2-api/2.1.7/archaius2-api-2.1.7.pom
Download https://jcenter.bintray.com/com/netflix/archaius/archaius2-core/2.1.7/archaius2-core-2.1.7.jar
Download https://jcenter.bintray.com/com/netflix/archaius/archaius2-api/2.1.7/archaius2-api-2.1.7.jar
警告: [options] 未与 -source 1.7 一起设置引导类路径
注: 某些输入文件使用或覆盖了已过时的 API。
注: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
1 个警告
:eureka-client-archaius2:processResources UP-TO-DATE
:eureka-client-archaius2:classes
:eureka-client-archaius2:writeManifestProperties
:eureka-client-archaius2:jar
:eureka-client-archaius2:assemble
:eureka-client-jersey2:compileJava
Download https://jcenter.bintray.com/org/glassfish/jersey/core/jersey-client/2.23.1/jersey-client-2.23.1.pom
Download https://jcenter.bintray.com/org/glassfish/jersey/project/2.23.1/project-2.23.1.pom
Download https://jcenter.bintray.com/org/glassfish/jersey/connectors/jersey-apache-connector/2.23.1/jersey-apache-connector-2.23.1.pom
Download https://jcenter.bintray.com/org/glassfish/jersey/connectors/project/2.23.1/project-2.23.1.pom
Download https://jcenter.bintray.com/javax/mail/mail/1.4.7/mail-1.4.7.pom
Download https://jcenter.bintray.com/com/sun/mail/all/1.4.7/all-1.4.7.pom
Download https://jcenter.bintray.com/org/glassfish/jersey/core/jersey-common/2.23.1/jersey-common-2.23.1.pom
Download https://jcenter.bintray.com/org/glassfish/hk2/hk2-parent/2.4.0-b34/hk2-parent-2.4.0-b34.pom
Download https://jcenter.bintray.com/org/glassfish/jersey/bundles/repackaged/jersey-guava/2.23.1/jersey-guava-2.23.1.pom
Download https://jcenter.bintray.com/org/glassfish/jersey/bundles/repackaged/project/2.23.1/project-2.23.1.pom
Download https://jcenter.bintray.com/org/glassfish/jersey/bundles/project/2.23.1/project-2.23.1.pom
Download https://jcenter.bintray.com/org/glassfish/jersey/core/jersey-client/2.23.1/jersey-client-2.23.1.jar
Download https://jcenter.bintray.com/org/glassfish/jersey/connectors/jersey-apache-connector/2.23.1/jersey-apache-connector-2.23.1.jar
Download https://jcenter.bintray.com/javax/mail/mail/1.4.7/mail-1.4.7.jar
Download https://jcenter.bintray.com/org/glassfish/jersey/core/jersey-common/2.23.1/jersey-common-2.23.1.jar
Download https://jcenter.bintray.com/org/glassfish/jersey/bundles/repackaged/jersey-guava/2.23.1/jersey-guava-2.23.1.jar
警告: [options] 未与 -source 1.7 一起设置引导类路径
注: 某些输入文件使用或覆盖了已过时的 API。
注: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
注: /Users/fanhongling/Downloads/workspace/src/github.com/Netflix/eureka/eureka-client-jersey2/src/main/java/com/netflix/discovery/shared/transport/jersey2/EurekaJersey2ClientImpl.java使用了未经检查或不安全的操作。
注: 有关详细信息, 请使用 -Xlint:unchecked 重新编译。
1 个警告
:eureka-client-jersey2:processResources UP-TO-DATE
:eureka-client-jersey2:classes
:eureka-client-jersey2:writeManifestProperties
:eureka-client-jersey2:jar
:eureka-client-jersey2:assemble
:eureka-core:compileJava
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-core/1.11.9/aws-java-sdk-core-1.11.9.pom
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-pom/1.11.9/aws-java-sdk-pom-1.11.9.pom
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-ec2/1.11.9/aws-java-sdk-ec2-1.11.9.pom
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-autoscaling/1.11.9/aws-java-sdk-autoscaling-1.11.9.pom
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-sts/1.11.9/aws-java-sdk-sts-1.11.9.pom
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-route53/1.11.9/aws-java-sdk-route53-1.11.9.pom
Download https://jcenter.bintray.com/joda-time/joda-time/2.8.1/joda-time-2.8.1.pom
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-core/1.11.9/aws-java-sdk-core-1.11.9.jar
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-ec2/1.11.9/aws-java-sdk-ec2-1.11.9.jar
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-autoscaling/1.11.9/aws-java-sdk-autoscaling-1.11.9.jar
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-sts/1.11.9/aws-java-sdk-sts-1.11.9.jar
Download https://jcenter.bintray.com/com/amazonaws/aws-java-sdk-route53/1.11.9/aws-java-sdk-route53-1.11.9.jar
Download https://jcenter.bintray.com/joda-time/joda-time/2.8.1/joda-time-2.8.1.jar
Download https://jcenter.bintray.com/com/fasterxml/jackson/core/jackson-databind/2.6.6/jackson-databind-2.6.6.jar
警告: [options] 未与 -source 1.7 一起设置引导类路径
/Users/fanhongling/Downloads/workspace/src/github.com/Netflix/eureka/eureka-core/src/main/java/com/netflix/eureka/DefaultEurekaServerConfig.java:501: 警告: 最后一个参数使用了不准确的变量类型的 varargs 方法的非 varargs 调用; 
                        new String[]{propName, remoteRegionUrlWithNamePair, pairSplitChar});
                        ^
  对于 varargs 调用, 应使用 Object
  对于非 varargs 调用, 应使用 Object[], 这样也可以抑制此警告
注: 某些输入文件使用或覆盖了已过时的 API。
注: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
注: 某些输入文件使用了未经检查或不安全的操作。
注: 有关详细信息, 请使用 -Xlint:unchecked 重新编译。
2 个警告
:eureka-core:processResources UP-TO-DATE
:eureka-core:classes
:eureka-core:writeManifestProperties
:eureka-core:jar
:eureka-core:assemble
:eureka-core-jersey2:compileJava
警告: [options] 未与 -source 1.7 一起设置引导类路径
1 个警告
:eureka-core-jersey2:processResources UP-TO-DATE
:eureka-core-jersey2:classes
:eureka-core-jersey2:writeManifestProperties
:eureka-core-jersey2:jar
:eureka-core-jersey2:assemble
:eureka-examples:compileJava
警告: [options] 未与 -source 1.7 一起设置引导类路径
1 个警告
:eureka-examples:processResources UP-TO-DATE
:eureka-examples:classes
:eureka-examples:writeManifestProperties
:eureka-examples:jar
:eureka-examples:ExampleClientStartScript
Download https://jcenter.bintray.com/org/slf4j/slf4j-simple/1.7.7/slf4j-simple-1.7.7.pom
Download https://jcenter.bintray.com/org/slf4j/slf4j-simple/1.7.7/slf4j-simple-1.7.7.jar
:eureka-examples:ExampleGovernatedServiceStartScript
:eureka-examples:ExampleServiceStartScript
:eureka-examples:startScripts SKIPPED
:eureka-examples:distTar
:eureka-examples:distZip
:eureka-examples:assemble
:eureka-resources:compileJava UP-TO-DATE
:eureka-resources:processResources
:eureka-resources:classes
:eureka-resources:writeManifestProperties
:eureka-resources:jar
:eureka-resources:assemble
:eureka-server:compileJava UP-TO-DATE
:eureka-server:processResources
:eureka-server:classes
:eureka-server:writeManifestProperties
:eureka-server:war
Download https://jcenter.bintray.com/com/sun/jersey/jersey-server/1.19.1/jersey-server-1.19.1.pom
Download https://jcenter.bintray.com/com/sun/jersey/jersey-servlet/1.19.1/jersey-servlet-1.19.1.pom
Download https://jcenter.bintray.com/xerces/xercesImpl/2.4.0/xercesImpl-2.4.0.pom
Download https://jcenter.bintray.com/com/sun/jersey/jersey-server/1.19.1/jersey-server-1.19.1.jar
Download https://jcenter.bintray.com/com/sun/jersey/jersey-servlet/1.19.1/jersey-servlet-1.19.1.jar
Download https://jcenter.bintray.com/xerces/xercesImpl/2.4.0/xercesImpl-2.4.0.jar
:eureka-server:assemble
:eureka-server-governator:compileJava
Download https://jcenter.bintray.com/com/netflix/governator/governator-servlet/1.12.10/governator-servlet-1.12.10.pom
Download https://jcenter.bintray.com/com/sun/jersey/contribs/jersey-guice/1.19.1/jersey-guice-1.19.1.pom
Download https://jcenter.bintray.com/com/google/inject/extensions/guice-servlet/4.0/guice-servlet-4.0.pom
Download https://jcenter.bintray.com/com/netflix/governator/governator-servlet/1.12.10/governator-servlet-1.12.10.jar
Download https://jcenter.bintray.com/com/sun/jersey/contribs/jersey-guice/1.19.1/jersey-guice-1.19.1.jar
Download https://jcenter.bintray.com/com/google/inject/extensions/guice-servlet/4.0/guice-servlet-4.0.jar
警告: [options] 未与 -source 1.7 一起设置引导类路径
1 个警告
:eureka-server-governator:processResources
:eureka-server-governator:classes
:eureka-server-governator:writeManifestProperties
:eureka-server-governator:war
:eureka-server-governator:assemble
:eureka-test-utils:compileJava
警告: [options] 未与 -source 1.7 一起设置引导类路径
注: 某些输入文件使用或覆盖了已过时的 API。
注: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
1 个警告
:eureka-test-utils:processResources UP-TO-DATE
:eureka-test-utils:classes
:eureka-test-utils:writeManifestProperties
:eureka-test-utils:jar
:eureka-test-utils:assemble
:collectNetflixOSS
:eureka-client:writeLicenseHeader
:eureka-client:licenseMain
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/AbstractEurekaIdentity.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/AmazonInfoConfig.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/Archaius1AmazonInfoConfig.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/EurekaAccept.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/EurekaClientIdentity.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/HealthCheckCallbackToHandlerBridge.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/PropertyBasedAmazonInfoConfigConstants.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/MyDataCenterInfo.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/HealthCheckHandler.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/PropertyBasedInstanceConfigConstants.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/providers/CloudInstanceConfigProvider.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/providers/VipAddressResolver.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/providers/Archaius1VipAddressResolver.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/providers/MyDataCenterInstanceConfigProvider.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/providers/EurekaConfigBasedInstanceInfoProvider.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/UniqueIdentifier.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/RefreshableAmazonInfoProvider.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/AbstractAzToRegionMapper.java
Missing header in: eureka-client/src/main/java/com/netflix/appinfo/RefreshableInstanceConfig.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/AzToRegionMapper.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/CommonConstants.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/CacheRefreshedEvent.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/AbstractDiscoveryClientOptionalArgs.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/jackson/DataCenterTypeInfoResolver.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/EurekaJacksonCodec.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/jackson/AbstractEurekaJacksonCodec.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/jackson/EurekaJacksonJsonModifiers.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/jackson/mixin/DataCenterInfoXmlMixIn.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/wrappers/CodecWrapperBase.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/KeyFormatter.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/wrappers/CodecWrapper.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/jackson/mixin/MiniInstanceInfoMixIn.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/DiscoveryEvent.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/wrappers/EncoderWrapper.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/wrappers/DecoderWrapper.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/converters/wrappers/CodecWrappers.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/DNSBasedAzToRegionMapper.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/endpoint/DnsResolver.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/EurekaClient.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/endpoint/EndpointUtils.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/EurekaEvent.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/guice/EurekaModule.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/EurekaUpStatusResolver.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/EurekaIdentityHeaderFilter.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/EurekaNamespace.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/InstanceInfoReplicator.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/EurekaEventListener.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/InstanceRegionChecker.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/internal/util/AmazonInfoUtils.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/InternalEurekaStatusModule.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/internal/util/Archaius1Utils.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/PropertyBasedAzToRegionMapper.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/NotImplementedRegistryImpl.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/providers/DefaultEurekaClientConfigProvider.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/MonitoredConnectionManager.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/NamedConnectionPool.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/resolver/aws/ApplicationsResolver.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/PropertyBasedClientConfigConstants.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/resolver/aws/EurekaHttpResolver.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/resolver/aws/ConfigClusterResolver.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/resolver/aws/AwsEndpoint.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/resolver/AsyncResolver.java
Unknown file extension: eureka-client/src/main/java/com/netflix/discovery/shared/resolver/README.md
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/resolver/ClosableResolver.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/DefaultEurekaTransportConfig.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/EurekaClientFactoryBuilder.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/EurekaHttpClient.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/EurekaTransportConfig.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/jersey/EurekaJerseyClient.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/jersey/AbstractJerseyEurekaHttpClient.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/jersey/EurekaJerseyClientImpl.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/jersey/Jersey1TransportClientFactories.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/PropertyBasedTransportConfigConstants.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/jersey/SSLSocketFactoryAdapter.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/jersey/Jersey1DiscoveryClientOptionalArgs.java
Unknown file extension: eureka-client/src/main/java/com/netflix/discovery/shared/transport/README.md
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/jersey/TransportClientFactories.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/TimedSupervisorTask.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/util/DiscoveryBuildInfo.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/StatusChangeEvent.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/shared/transport/TransportClientFactory.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/util/EurekaEntityComparators.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/util/EurekaUtils.java
Missing header in: eureka-client/src/main/java/com/netflix/discovery/util/StringCache.java
:eureka-client:licenseTest
Missing header in: eureka-client/src/test/java/com/netflix/appinfo/CloudInstanceConfigTest.java
Missing header in: eureka-client/src/test/java/com/netflix/appinfo/InstanceInfoTest.java
Missing header in: eureka-client/src/test/java/com/netflix/appinfo/ApplicationInfoManagerTest.java
Missing header in: eureka-client/src/test/java/com/netflix/appinfo/AmazonInfoTest.java
Missing header in: eureka-client/src/test/java/com/netflix/appinfo/RefreshableAmazonInfoProviderTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/AbstractDiscoveryClientTester.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/BackUpRegistryTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/BaseDiscoveryClientTester.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/converters/CodecLoadTester.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/converters/EurekaCodecCompatibilityTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/converters/EurekaJacksonCodecTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/converters/wrappers/CodecWrappersTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/converters/EurekaJsonAndXmlJacksonCodecTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/converters/XmlCodecTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/DiscoveryClientCloseJerseyThreadTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/converters/StringCacheTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/DiscoveryClientHealthTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/DiscoveryClientRedirectTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/DiscoveryClientDisableRegistryTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/DiscoveryClientRegisterUpdateTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/DiscoveryClientRegistryTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/DiscoveryClientEventBusTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/InstanceInfoReplicatorTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/guice/EurekaModuleTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/EurekaEventListenerTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/EurekaClientLifecycleTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/Jersey1DiscoveryClientOptionalArgsTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/InstanceRegionCheckerTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/MockBackupRegistry.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/MockRemoteEurekaServer.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/shared/ApplicationsTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/shared/resolver/AsyncResolverTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/shared/resolver/aws/ApplicationsResolverTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/shared/resolver/aws/TestEurekaHttpResolver.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/shared/resolver/aws/EurekaHttpResolverTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/shared/resolver/aws/ConfigClusterResolverTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/TimedSupervisorTaskTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/util/DiscoveryBuildInfoTest.java
Missing header in: eureka-client/src/test/java/com/netflix/discovery/util/EurekaUtilsTest.java
:eureka-client:license
:eureka-client:compileTestJava
Download https://jcenter.bintray.com/org/mortbay/jetty/jetty/6.1H.22/jetty-6.1H.22.pom
Download https://jcenter.bintray.com/org/mortbay/jetty/project/6.1H.22/project-6.1H.22.pom
Download https://jcenter.bintray.com/org/mortbay/jetty/jetty-parent/8/jetty-parent-8.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-parent/9/jetty-parent-9.pom
Download https://jcenter.bintray.com/org/mock-server/mockserver-netty/3.9.2/mockserver-netty-3.9.2.pom
Download https://jcenter.bintray.com/org/mock-server/mockserver/3.9.2/mockserver-3.9.2.pom
Download https://jcenter.bintray.com/org/mortbay/jetty/jetty-util/6.1H.22/jetty-util-6.1H.22.pom
Download https://jcenter.bintray.com/org/mock-server/mockserver-client-java/3.9.2/mockserver-client-java-3.9.2.pom
Download https://jcenter.bintray.com/org/mock-server/mockserver-core/3.9.2/mockserver-core-3.9.2.pom
Download https://jcenter.bintray.com/io/netty/netty-buffer/4.0.24.Final/netty-buffer-4.0.24.Final.pom
Download https://jcenter.bintray.com/io/netty/netty-parent/4.0.24.Final/netty-parent-4.0.24.Final.pom
Download https://jcenter.bintray.com/io/netty/netty-codec/4.0.24.Final/netty-codec-4.0.24.Final.pom
Download https://jcenter.bintray.com/io/netty/netty-codec-http/4.0.24.Final/netty-codec-http-4.0.24.Final.pom
Download https://jcenter.bintray.com/io/netty/netty-common/4.0.24.Final/netty-common-4.0.24.Final.pom
Download https://jcenter.bintray.com/io/netty/netty-handler/4.0.24.Final/netty-handler-4.0.24.Final.pom
Download https://jcenter.bintray.com/io/netty/netty-transport/4.0.24.Final/netty-transport-4.0.24.Final.pom
Download https://jcenter.bintray.com/org/bouncycastle/bcprov-jdk15on/1.51/bcprov-jdk15on-1.51.pom
Download https://jcenter.bintray.com/io/netty/netty-codec-socks/4.0.24.Final/netty-codec-socks-4.0.24.Final.pom
Download https://jcenter.bintray.com/org/mortbay/jetty/jetty/6.1H.22/jetty-6.1H.22.jar
Download https://jcenter.bintray.com/org/mock-server/mockserver-netty/3.9.2/mockserver-netty-3.9.2.jar
Download https://jcenter.bintray.com/org/mortbay/jetty/jetty-util/6.1H.22/jetty-util-6.1H.22.jar
Download https://jcenter.bintray.com/org/mock-server/mockserver-client-java/3.9.2/mockserver-client-java-3.9.2.jar
Download https://jcenter.bintray.com/org/mock-server/mockserver-core/3.9.2/mockserver-core-3.9.2.jar
Download https://jcenter.bintray.com/io/netty/netty-buffer/4.0.24.Final/netty-buffer-4.0.24.Final.jar
Download https://jcenter.bintray.com/io/netty/netty-codec/4.0.24.Final/netty-codec-4.0.24.Final.jar
Download https://jcenter.bintray.com/io/netty/netty-codec-http/4.0.24.Final/netty-codec-http-4.0.24.Final.jar
Download https://jcenter.bintray.com/io/netty/netty-common/4.0.24.Final/netty-common-4.0.24.Final.jar
Download https://jcenter.bintray.com/io/netty/netty-handler/4.0.24.Final/netty-handler-4.0.24.Final.jar
Download https://jcenter.bintray.com/io/netty/netty-transport/4.0.24.Final/netty-transport-4.0.24.Final.jar
Download https://jcenter.bintray.com/org/bouncycastle/bcprov-jdk15on/1.51/bcprov-jdk15on-1.51.jar
Download https://jcenter.bintray.com/io/netty/netty-codec-socks/4.0.24.Final/netty-codec-socks-4.0.24.Final.jar
警告: [options] 未与 -source 1.7 一起设置引导类路径
注: 某些输入文件使用或覆盖了已过时的 API。
注: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
注: 某些输入文件使用了未经检查或不安全的操作。
注: 有关详细信息, 请使用 -Xlint:unchecked 重新编译。
1 个警告
:eureka-client:processTestResources UP-TO-DATE
:eureka-client:testClasses
:eureka-client:test
:eureka-client:check
:eureka-client:build
:eureka-client-archaius2:writeLicenseHeader
:eureka-client-archaius2:licenseMain
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/appinfo/Archaius2AmazonInfoConfig.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/appinfo/Ec2EurekaArchaius2InstanceConfig.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/appinfo/providers/Archaius2VipAddressResolver.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/appinfo/EurekaArchaius2InstanceConfig.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/appinfo/providers/CompositeInstanceConfigFactory.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/appinfo/providers/EurekaInstanceConfigFactory.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/discovery/EurekaArchaius2ClientConfig.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/discovery/guice/InternalEurekaClientModule.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/discovery/guice/EurekaClientModule.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/discovery/internal/util/InternalPrefixedConfig.java
Missing header in: eureka-client-archaius2/src/main/java/com/netflix/discovery/shared/transport/EurekaArchaius2TransportConfig.java
:eureka-client-archaius2:licenseTest
Missing header in: eureka-client-archaius2/src/test/java/com/netflix/appinfo/Ec2EurekaArchaius2InstanceConfigTest.java
Missing header in: eureka-client-archaius2/src/test/java/com/netflix/discovery/guice/EurekaClientModuleConfigurationTest.java
Missing header in: eureka-client-archaius2/src/test/java/com/netflix/discovery/guice/Ec2EurekaClientModuleTest.java
Missing header in: eureka-client-archaius2/src/test/java/com/netflix/discovery/guice/NonEc2EurekaClientModuleTest.java
Missing header in: eureka-client-archaius2/src/test/java/com/netflix/discovery/internal/util/InternalPrefixedConfigTest.java
:eureka-client-archaius2:license
:eureka-client-archaius2:compileTestJava
Download https://jcenter.bintray.com/com/netflix/archaius/archaius2-guice/2.1.7/archaius2-guice-2.1.7.pom
Download https://jcenter.bintray.com/com/netflix/archaius/archaius2-guice/2.1.7/archaius2-guice-2.1.7.jar
警告: [options] 未与 -source 1.7 一起设置引导类路径
注: 某些输入文件使用或覆盖了已过时的 API。
注: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
1 个警告
:eureka-client-archaius2:processTestResources UP-TO-DATE
:eureka-client-archaius2:testClasses
:eureka-client-archaius2:test
:eureka-client-archaius2:check
:eureka-client-archaius2:build
:eureka-client-jersey2:writeLicenseHeader
:eureka-client-jersey2:licenseMain
Missing header in: eureka-client-jersey2/src/main/java/com/netflix/discovery/guice/Jersey2EurekaModule.java
Missing header in: eureka-client-jersey2/src/main/java/com/netflix/discovery/shared/transport/jersey2/EurekaIdentityHeaderFilter.java
Missing header in: eureka-client-jersey2/src/main/java/com/netflix/discovery/Jersey2DiscoveryClientOptionalArgs.java
Missing header in: eureka-client-jersey2/src/main/java/com/netflix/discovery/shared/transport/jersey2/EurekaJersey2Client.java
Missing header in: eureka-client-jersey2/src/main/java/com/netflix/discovery/shared/transport/jersey2/Jersey2ApplicationClient.java
Missing header in: eureka-client-jersey2/src/main/java/com/netflix/discovery/shared/transport/jersey2/Jersey2TransportClientFactories.java
Missing header in: eureka-client-jersey2/src/main/java/com/netflix/discovery/shared/transport/jersey2/Jersey2EurekaIdentityHeaderFilter.java
Missing header in: eureka-client-jersey2/src/main/java/com/netflix/discovery/shared/transport/jersey2/EurekaJersey2ClientImpl.java
:eureka-client-jersey2:licenseTest
Missing header in: eureka-client-jersey2/src/test/java/com/netflix/discovery/guice/Jersey2EurekaModuleTest.java
:eureka-client-jersey2:license
:eureka-client-jersey2:compileTestJava
警告: [options] 未与 -source 1.7 一起设置引导类路径
注: /Users/fanhongling/Downloads/workspace/src/github.com/Netflix/eureka/eureka-client-jersey2/src/test/java/com/netflix/discovery/guice/Jersey2EurekaModuleTest.java使用或覆盖了已过时的 API。
注: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
1 个警告
:eureka-client-jersey2:processTestResources UP-TO-DATE
:eureka-client-jersey2:testClasses
:eureka-client-jersey2:test
:eureka-client-jersey2:check
:eureka-client-jersey2:build
:eureka-core:writeLicenseHeader
:eureka-core:licenseMain
Missing header in: eureka-core/src/main/java/com/netflix/eureka/aws/AwsBinder.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/aws/AwsBinderDelegate.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/aws/AwsBindingStrategy.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/AsgReplicationTask.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/aws/ElasticNetworkInterfaceBinder.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/aws/Route53Binder.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/HttpReplicationClient.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/InstanceReplicationTask.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/protocol/ReplicationInstance.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/protocol/ReplicationInstanceResponse.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/PeerEurekaNodes.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/protocol/ReplicationList.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/DynamicGZIPContentEncodingFilter.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/ReplicationTask.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/protocol/ReplicationListResponse.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/cluster/ReplicationTaskProcessor.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/GzipEncodingEnforcingFilter.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/EurekaServerIdentity.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/Key.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/InstanceRegistry.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/rule/AsgEnabledRule.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/rule/DownOrStartingRule.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/rule/AlwaysMatchInstanceStatusRule.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/rule/InstanceStatusOverrideRule.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/ResponseCache.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/rule/StatusOverrideResult.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/rule/FirstMatchWinsCompositeRule.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/resources/DefaultServerCodecs.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/resources/ServerCodecs.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/resources/ServerInfoResource.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/rule/OverrideExistsRule.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/ServerRequestAuthFilter.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/transport/JerseyReplicationClient.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/util/batcher/TaskDispatchers.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/util/batcher/AcceptorExecutor.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/registry/rule/LeaseExistsRule.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/util/batcher/TaskProcessor.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/util/batcher/TaskExecutors.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/util/batcher/TaskHolder.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/util/StatusUtil.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/util/batcher/TaskDispatcher.java
Missing header in: eureka-core/src/main/java/com/netflix/eureka/util/StatusInfo.java
:eureka-core:licenseTest
Missing header in: eureka-core/src/test/java/com/netflix/eureka/cluster/JerseyReplicationClientTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/AbstractTester.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/cluster/PeerEurekaNodesTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/cluster/PeerEurekaNodeTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/cluster/protocol/JacksonEncodingTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/DefaultEurekaServerConfigTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/cluster/TestableHttpReplicationClient.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/cluster/TestableInstanceReplicationTask.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/registry/AwsInstanceRegistryTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/cluster/ReplicationTaskProcessorTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/registry/ResponseCacheTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/RemoteRegionSoftDependencyTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/mock/MockRemoteEurekaServer.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/resources/ApplicationsResourceTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/registry/InstanceRegistryTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/resources/InstanceResourceTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/resources/PeerReplicationResourceTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/resources/ApplicationResourceTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/resources/ReplicationConcurrencyTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/resources/AbstractVIPResourceTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/util/AwsAsgUtilTest.java
Missing header in: eureka-core/src/test/java/com/netflix/eureka/util/StatusUtilTest.java
:eureka-core:license
:eureka-core:compileTestJava
警告: [options] 未与 -source 1.7 一起设置引导类路径
注: 某些输入文件使用或覆盖了已过时的 API。
注: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
1 个警告
:eureka-core:processTestResources UP-TO-DATE
:eureka-core:testClasses
:eureka-core:test
Download https://jcenter.bintray.com/org/slf4j/slf4j-simple/1.7.10/slf4j-simple-1.7.10.pom
Download https://jcenter.bintray.com/org/slf4j/slf4j-simple/1.7.10/slf4j-simple-1.7.10.jar
:eureka-core:check
:eureka-core:build
:eureka-core-jersey2:writeLicenseHeader
:eureka-core-jersey2:licenseMain
Missing header in: eureka-core-jersey2/src/main/java/com/netflix/eureka/resources/EurekaServerContextBinder.java
Missing header in: eureka-core-jersey2/src/main/java/com/netflix/eureka/Jersey2EurekaBootStrap.java
Missing header in: eureka-core-jersey2/src/main/java/com/netflix/eureka/cluster/Jersey2PeerEurekaNodes.java
Missing header in: eureka-core-jersey2/src/main/java/com/netflix/eureka/transport/Jersey2DynamicGZIPContentEncodingFilter.java
Missing header in: eureka-core-jersey2/src/main/java/com/netflix/eureka/transport/Jersey2ReplicationClient.java
:eureka-core-jersey2:licenseTest
Missing header in: eureka-core-jersey2/src/test/java/com/netflix/eureka/transport/Jersey2ReplicationClientTest.java
:eureka-core-jersey2:license
:eureka-core-jersey2:compileTestJava
警告: [options] 未与 -source 1.7 一起设置引导类路径
1 个警告
:eureka-core-jersey2:processTestResources UP-TO-DATE
:eureka-core-jersey2:testClasses
:eureka-core-jersey2:test
:eureka-core-jersey2:check
:eureka-core-jersey2:build
:eureka-examples:writeLicenseHeader
:eureka-examples:licenseMain
Missing header in: eureka-examples/src/main/java/com/netflix/eureka/ExampleEurekaGovernatedService.java
Missing header in: eureka-examples/src/main/java/com/netflix/eureka/ExampleServiceBase.java
:eureka-examples:licenseTest UP-TO-DATE
:eureka-examples:license
:eureka-examples:compileTestJava UP-TO-DATE
:eureka-examples:processTestResources UP-TO-DATE
:eureka-examples:testClasses UP-TO-DATE
:eureka-examples:test UP-TO-DATE
:eureka-examples:check
:eureka-examples:build
:eureka-resources:writeLicenseHeader
:eureka-resources:licenseMain
Missing header in: eureka-resources/src/main/resources/css/main.css
Missing header in: eureka-resources/src/main/resources/css/jquery-ui-1.7.2.custom.css
Missing header in: eureka-resources/src/main/resources/jsp/lastN.jsp
Missing header in: eureka-resources/src/main/resources/jsp/navbar.jsp
Missing header in: eureka-resources/src/main/resources/jsp/status.jsp
Missing header in: eureka-resources/src/main/resources/jsp/header.jsp
Missing header in: eureka-resources/src/main/resources/js/jquery.dataTables.js
:eureka-resources:licenseTest UP-TO-DATE
:eureka-resources:license
:eureka-resources:compileTestJava UP-TO-DATE
:eureka-resources:processTestResources UP-TO-DATE
:eureka-resources:testClasses UP-TO-DATE
:eureka-resources:test UP-TO-DATE
:eureka-resources:check
:eureka-resources:build
:eureka-server:writeLicenseHeader
:eureka-server:licenseMain
Missing header in: eureka-server/src/main/resources/log4j.properties
Missing header in: eureka-server/src/main/resources/eureka-server.properties
Missing header in: eureka-server/src/main/resources/eureka-client.properties
:eureka-server:licenseTest
Missing header in: eureka-server/src/test/java/com/netflix/eureka/resources/EurekaClientServerRestIntegrationTest.java
:eureka-server:license
:eureka-server:compileTestJava
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-server/7.2.0.v20101020/jetty-server-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-project/7.2.0.v20101020/jetty-project-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-parent/16/jetty-parent-16.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-webapp/7.2.0.v20101020/jetty-webapp-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-continuation/7.2.0.v20101020/jetty-continuation-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-http/7.2.0.v20101020/jetty-http-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-xml/7.2.0.v20101020/jetty-xml-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-servlet/7.2.0.v20101020/jetty-servlet-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-io/7.2.0.v20101020/jetty-io-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-util/7.2.0.v20101020/jetty-util-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-security/7.2.0.v20101020/jetty-security-7.2.0.v20101020.pom
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-server/7.2.0.v20101020/jetty-server-7.2.0.v20101020.jar
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-webapp/7.2.0.v20101020/jetty-webapp-7.2.0.v20101020.jar
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-continuation/7.2.0.v20101020/jetty-continuation-7.2.0.v20101020.jar
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-http/7.2.0.v20101020/jetty-http-7.2.0.v20101020.jar
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-xml/7.2.0.v20101020/jetty-xml-7.2.0.v20101020.jar
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-servlet/7.2.0.v20101020/jetty-servlet-7.2.0.v20101020.jar
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-io/7.2.0.v20101020/jetty-io-7.2.0.v20101020.jar
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-util/7.2.0.v20101020/jetty-util-7.2.0.v20101020.jar
Download https://jcenter.bintray.com/org/eclipse/jetty/jetty-security/7.2.0.v20101020/jetty-security-7.2.0.v20101020.jar
警告: [options] 未与 -source 1.7 一起设置引导类路径
1 个警告
:eureka-server:processTestResources UP-TO-DATE
:eureka-server:testClasses
:eureka-server:test
:eureka-server:check
:eureka-server:build
:eureka-server-governator:writeLicenseHeader
:eureka-server-governator:licenseMain
Missing header in: eureka-server-governator/src/main/java/com/netflix/eureka/EurekaContextListener.java
Missing header in: eureka-server-governator/src/main/resources/eureka-client.properties
Missing header in: eureka-server-governator/src/main/java/com/netflix/eureka/EurekaInjectorCreator.java
Missing header in: eureka-server-governator/src/main/java/com/netflix/eureka/guice/LocalDevEurekaServerModule.java
Missing header in: eureka-server-governator/src/main/resources/eureka-server.properties
Missing header in: eureka-server-governator/src/main/java/com/netflix/eureka/guice/Ec2EurekaServerModule.java
Missing header in: eureka-server-governator/src/main/resources/log4j.properties
:eureka-server-governator:licenseTest UP-TO-DATE
:eureka-server-governator:license
:eureka-server-governator:compileTestJava UP-TO-DATE
:eureka-server-governator:processTestResources UP-TO-DATE
:eureka-server-governator:testClasses UP-TO-DATE
:eureka-server-governator:test UP-TO-DATE
:eureka-server-governator:check
:eureka-server-governator:build
:eureka-test-utils:writeLicenseHeader
:eureka-test-utils:licenseMain
Missing header in: eureka-test-utils/src/main/java/com/netflix/discovery/junit/resource/DiscoveryClientResource.java
Missing header in: eureka-test-utils/src/main/java/com/netflix/discovery/shared/transport/ClusterSampleData.java
Missing header in: eureka-test-utils/src/main/java/com/netflix/discovery/util/ApplicationFunctions.java
Missing header in: eureka-test-utils/src/main/java/com/netflix/discovery/util/InstanceInfoGenerator.java
Missing header in: eureka-test-utils/src/main/java/com/netflix/discovery/util/DiagnosticClient.java
:eureka-test-utils:licenseTest UP-TO-DATE
:eureka-test-utils:license
:eureka-test-utils:compileTestJava
警告: [options] 未与 -source 1.7 一起设置引导类路径
1 个警告
:eureka-test-utils:processTestResources UP-TO-DATE
:eureka-test-utils:testClasses
:eureka-test-utils:test
:eureka-test-utils:check
:eureka-test-utils:build

BUILD SUCCESSFUL

Total time: 32 mins 5.164 secs

This build could be faster, please consider using the Gradle Daemon: https://docs.gradle.org/2.10/userguide/gradle_daemon.html
```

### Dashboard - https://github.com/Netflix/eureka-ui

```
fanhonglingdeMacBook-Pro:eureka fanhongling$ cd ..
fanhonglingdeMacBook-Pro:Netflix fanhongling$ git clone https://github.com/Netflix/eureka-ui eureka-ui
Cloning into 'eureka-ui'...
remote: Counting objects: 155, done.
remote: Total 155 (delta 0), reused 0 (delta 0), pack-reused 155
Receiving objects: 100% (155/155), 81.15 KiB | 47.00 KiB/s, done.
Resolving deltas: 100% (59/59), done.
Checking connectivity... done.
```

### Docker image - https://hub.docker.com/r/netflixoss/eureka

Inspired
```
fanhonglingdeMacBook-Pro:99-mirror fanhongling$ cd ../workspace/src/github.com/
fanhonglingdeMacBook-Pro:github.com fanhongling$ git clone https://github.com/Netflix-Skunkworks/zerotodocker Netflix-Skunkworks/zerotodocker
Cloning into 'Netflix-Skunkworks/zerotodocker'...
remote: Counting objects: 538, done.
remote: Total 538 (delta 0), reused 0 (delta 0), pack-reused 538
Receiving objects: 100% (538/538), 10.03 MiB | 57.00 KiB/s, done.
Resolving deltas: 100% (244/244), done.
Checking connectivity... done.
```

Build
```
fanhonglingdeMacBook-Pro:libs fanhongling$ cp ../../../../../Netflix/eureka/eureka-server/build/libs/eureka-server-1.6.3-SNAPSHOT.war .

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/netflix$ docker build --rm=true -t tangfeixiong/netflix-eureka:1.6.3-SNAPSHOT -f /work/src/github.com/tangfeixiong/go-to-kubernetes/examples/netflix/Dockerfile.eureka /work/src/github.com/tangfeixiong/go-to-kubernetes/examples/netflix/
Sending build context to Docker daemon 19.43 MB
Step 1 : FROM tomcat:8
 ---> 3695a0fe8320
Step 2 : LABEL author "tangfeixiong <tangfx128@gmail.com>" description "Inspired of Netflix Open Source Development (https://github.com/Netflix-Skunkworks/zerotodocker)"
 ---> Running in fbce1a03c807
 ---> 564832e102c8
Removing intermediate container fbce1a03c807
Step 3 : ARG eureka-server-pkg=/github.com/Netflix/eureka/eureka-server/build/libs/eureka-server-1.6.3-SNAPSHOT.war
 ---> Running in dbfebf56dc36
 ---> 4d65c13fcaf6
Removing intermediate container dbfebf56dc36
Step 4 : COPY /libs/eureka-server-1.6.3-SNAPSHOT.war /
 ---> abc79da00301
Removing intermediate container 51ae2c4ac3a2
Step 5 : RUN cd /usr/local/tomcat/webapps &&  unzip -q -d /usr/local/tomcat/webapps/eureka /eureka-server-1.6.3-SNAPSHOT.war   && rm /eureka-server-1.6.3-SNAPSHOT.war   && echo -e "archaius.deployment.applicationId=eureka\narchaius.deployment.environment=test" > /usr/local/tomcat/webapps/eureka/WEB-INF/classes/config.properties   && echo -e "eureka.enableSelfPreservation=false\neureka.registration.enabled=false\neureka.numberRegistrySyncRetries=0\neureka.waitTimeInMsWhenSyncEmpty=0" > /usr/local/tomcat/webapps/eureka/WEB-INF/classes/eureka-server-test.properties   && echo -e "eureka.serviceUrl.default=http://localhost:8080/eureka/v2/\neureka.vipAddress=eureka" > /usr/local/tomcat/webapps/eureka/WEB-INF/classes/eureka-client-test.properties   && echo $CATALINA_OPTS
 ---> Running in 991a0e55ac8d

 ---> 7ca1497c47be
Removing intermediate container 991a0e55ac8d
Step 6 : EXPOSE 8080
 ---> Running in ce95023c54f8
 ---> 67847b3c73a0
Removing intermediate container ce95023c54f8
Step 7 : ENTRYPOINT /usr/local/tomcat/bin/catalina.sh
 ---> Running in 231dfffb9aaa
 ---> 8ec47967bd1f
Removing intermediate container 231dfffb9aaa
Step 8 : CMD run
 ---> Running in fcda4fd2dd05
 ---> 31f2488d6efa
Removing intermediate container fcda4fd2dd05
Successfully built 31f2488d6efa

```