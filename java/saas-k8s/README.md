# Help

A Spring Boot web app with Application Catalog Provisioner for Kubernetes 

## Development

Build parent only
```
[vagrant@kubedev-172-17-4-59 saas-k8s]$ mvn install --non-recursive
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building MOOC for Kubernetes 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-install-plugin:2.4:install (default-install) @ mooc-parent ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/pom.xml to /home/vagrant/.m2/repository/https0x3A0x2F0x2Fgithub0x2Ecom0x2Fstackdocker/mooc-k8s/mooc-parent/0.0.1-SNAPSHOT/mooc-parent-0.0.1-SNAPSHOT.pom
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 3.338 s
[INFO] Finished at: 2018-06-28T15:40:20Z
[INFO] Final Memory: 6M/31M
[INFO] ------------------------------------------------------------------------
```

Web server
```
[vagrant@kubedev-172-17-4-59 saas-k8s]$ mvn package spring-boot:run --projects web-server
[INFO] Scanning for projects...
[INFO] 
[INFO] ------------------------------------------------------------------------
[INFO] Building Web Server 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-resources-plugin:2.7:resources (default-resources) @ web-server ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/src/main/webapp
[INFO] Copying 8 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:compile (default-compile) @ web-server ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 9 source files to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/target/classes
[INFO] 
[INFO] --- maven-resources-plugin:2.7:testResources (default-testResources) @ web-server ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:testCompile (default-testCompile) @ web-server ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-surefire-plugin:2.22.0:test (default-test) @ web-server ---
[INFO] Tests are skipped.
[INFO] 
[INFO] --- maven-jar-plugin:2.4:jar (default-jar) @ web-server ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/target/web-server-0.0.1-SNAPSHOT.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:2.0.3.RELEASE:repackage (re-packaging) @ web-server ---
[INFO] 
[INFO] >>> spring-boot-maven-plugin:2.0.3.RELEASE:run (default-cli) > test-compile @ web-server >>>
[INFO] 
[INFO] --- maven-resources-plugin:2.7:resources (default-resources) @ web-server ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/src/main/webapp
[INFO] Copying 8 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:compile (default-compile) @ web-server ---
[INFO] Nothing to compile - all classes are up to date
[INFO] 
[INFO] --- maven-resources-plugin:2.7:testResources (default-testResources) @ web-server ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.7.0:testCompile (default-testCompile) @ web-server ---
[INFO] No sources to compile
[INFO] 
[INFO] <<< spring-boot-maven-plugin:2.0.3.RELEASE:run (default-cli) < test-compile @ web-server <<<
[INFO] 
[INFO] 
[INFO] --- spring-boot-maven-plugin:2.0.3.RELEASE:run (default-cli) @ web-server ---
[INFO] Attaching agents: []
SLF4J: Class path contains multiple SLF4J bindings.
SLF4J: Found binding in [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/ch/qos/logback/logback-classic/1.2.3/logback-classic-1.2.3.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: Found binding in [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/org/slf4j/slf4j-log4j12/1.7.25/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
SLF4J: See http://www.slf4j.org/codes.html#multiple_bindings for an explanation.
15:47:10,890 |-INFO in ch.qos.logback.classic.LoggerContext[default] - Could NOT find resource [logback-test.xml]
15:47:10,891 |-INFO in ch.qos.logback.classic.LoggerContext[default] - Could NOT find resource [logback.groovy]
15:47:10,892 |-INFO in ch.qos.logback.classic.LoggerContext[default] - Found resource [logback.xml] at [file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/target/classes/logback.xml]
15:47:10,894 |-WARN in ch.qos.logback.classic.LoggerContext[default] - Resource [logback.xml] occurs multiple times on the classpath.
15:47:10,894 |-WARN in ch.qos.logback.classic.LoggerContext[default] - Resource [logback.xml] occurs at [file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/target/classes/logback.xml]
15:47:10,894 |-WARN in ch.qos.logback.classic.LoggerContext[default] - Resource [logback.xml] occurs at [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/io/kubernetes/client-java/2.0.0-beta1/client-java-2.0.0-beta1.jar!/logback.xml]
15:47:11,157 |-INFO in ch.qos.logback.classic.joran.action.ConfigurationAction - debug attribute not set
15:47:11,178 |-INFO in ch.qos.logback.core.joran.util.ConfigurationWatchListUtil@1794d431 - Adding [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/org/springframework/boot/spring-boot/2.0.3.RELEASE/spring-boot-2.0.3.RELEASE.jar!/org/springframework/boot/logging/logback/base.xml] to configuration watch list.
15:47:11,178 |-INFO in ch.qos.logback.core.joran.spi.ConfigurationWatchList@42e26948 - URL [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/org/springframework/boot/spring-boot/2.0.3.RELEASE/spring-boot-2.0.3.RELEASE.jar!/org/springframework/boot/logging/logback/base.xml] is not of type file
15:47:11,189 |-INFO in ch.qos.logback.core.joran.util.ConfigurationWatchListUtil@1794d431 - Adding [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/org/springframework/boot/spring-boot/2.0.3.RELEASE/spring-boot-2.0.3.RELEASE.jar!/org/springframework/boot/logging/logback/defaults.xml] to configuration watch list.
15:47:11,189 |-INFO in ch.qos.logback.core.joran.spi.ConfigurationWatchList@42e26948 - URL [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/org/springframework/boot/spring-boot/2.0.3.RELEASE/spring-boot-2.0.3.RELEASE.jar!/org/springframework/boot/logging/logback/defaults.xml] is not of type file
15:47:11,194 |-INFO in ch.qos.logback.core.joran.action.ConversionRuleAction - registering conversion word clr with class [org.springframework.boot.logging.logback.ColorConverter]
15:47:11,194 |-INFO in ch.qos.logback.core.joran.action.ConversionRuleAction - registering conversion word wex with class [org.springframework.boot.logging.logback.WhitespaceThrowableProxyConverter]
15:47:11,194 |-INFO in ch.qos.logback.core.joran.action.ConversionRuleAction - registering conversion word wEx with class [org.springframework.boot.logging.logback.ExtendedWhitespaceThrowableProxyConverter]
15:47:11,211 |-INFO in ch.qos.logback.classic.joran.action.LoggerAction - Setting level of logger [org.apache.catalina.startup.DigesterFactory] to ERROR
15:47:11,211 |-INFO in ch.qos.logback.classic.joran.action.LoggerAction - Setting level of logger [org.apache.catalina.util.LifecycleBase] to ERROR
15:47:11,212 |-INFO in ch.qos.logback.classic.joran.action.LoggerAction - Setting level of logger [org.apache.coyote.http11.Http11NioProtocol] to WARN
15:47:11,212 |-INFO in ch.qos.logback.classic.joran.action.LoggerAction - Setting level of logger [org.apache.sshd.common.util.SecurityUtils] to WARN
15:47:11,213 |-INFO in ch.qos.logback.classic.joran.action.LoggerAction - Setting level of logger [org.apache.tomcat.util.net.NioSelectorPool] to WARN
15:47:11,213 |-INFO in ch.qos.logback.classic.joran.action.LoggerAction - Setting level of logger [org.eclipse.jetty.util.component.AbstractLifeCycle] to ERROR
15:47:11,213 |-INFO in ch.qos.logback.classic.joran.action.LoggerAction - Setting level of logger [org.hibernate.validator.internal.util.Version] to WARN
15:47:11,215 |-INFO in ch.qos.logback.core.joran.util.ConfigurationWatchListUtil@1794d431 - Adding [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/org/springframework/boot/spring-boot/2.0.3.RELEASE/spring-boot-2.0.3.RELEASE.jar!/org/springframework/boot/logging/logback/console-appender.xml] to configuration watch list.
15:47:11,215 |-INFO in ch.qos.logback.core.joran.spi.ConfigurationWatchList@42e26948 - URL [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/org/springframework/boot/spring-boot/2.0.3.RELEASE/spring-boot-2.0.3.RELEASE.jar!/org/springframework/boot/logging/logback/console-appender.xml] is not of type file
15:47:11,219 |-INFO in ch.qos.logback.core.joran.action.AppenderAction - About to instantiate appender of type [ch.qos.logback.core.ConsoleAppender]
15:47:11,237 |-INFO in ch.qos.logback.core.joran.action.AppenderAction - Naming appender as [CONSOLE]
15:47:11,261 |-INFO in ch.qos.logback.core.joran.action.NestedComplexPropertyIA - Assuming default type [ch.qos.logback.classic.encoder.PatternLayoutEncoder] for [encoder] property
15:47:11,443 |-INFO in ch.qos.logback.core.joran.util.ConfigurationWatchListUtil@1794d431 - Adding [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/org/springframework/boot/spring-boot/2.0.3.RELEASE/spring-boot-2.0.3.RELEASE.jar!/org/springframework/boot/logging/logback/file-appender.xml] to configuration watch list.
15:47:11,443 |-INFO in ch.qos.logback.core.joran.spi.ConfigurationWatchList@42e26948 - URL [jar:file:/Users/fanhongling/Downloads/99-mirror/0x2Em20x2Frepository/org/springframework/boot/spring-boot/2.0.3.RELEASE/spring-boot-2.0.3.RELEASE.jar!/org/springframework/boot/logging/logback/file-appender.xml] is not of type file
15:47:11,447 |-INFO in ch.qos.logback.core.joran.action.AppenderAction - About to instantiate appender of type [ch.qos.logback.core.rolling.RollingFileAppender]
15:47:11,456 |-INFO in ch.qos.logback.core.joran.action.AppenderAction - Naming appender as [FILE]
15:47:11,460 |-INFO in ch.qos.logback.core.joran.action.NestedComplexPropertyIA - Assuming default type [ch.qos.logback.classic.encoder.PatternLayoutEncoder] for [encoder] property
15:47:11,482 |-INFO in c.q.l.core.rolling.SizeAndTimeBasedRollingPolicy@1471868639 - Archive files will be limited to [10 MB] each.
15:47:11,488 |-INFO in c.q.l.core.rolling.SizeAndTimeBasedRollingPolicy@1471868639 - Will use gz compression
15:47:11,491 |-INFO in c.q.l.core.rolling.SizeAndTimeBasedRollingPolicy@1471868639 - Will use the pattern /tmp/spring.log.%d{yyyy-MM-dd}.%i for the active file
15:47:11,503 |-INFO in ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP@343f4d3d - The date pattern is 'yyyy-MM-dd' from file name pattern '/tmp/spring.log.%d{yyyy-MM-dd}.%i.gz'.
15:47:11,503 |-INFO in ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP@343f4d3d - Roll-over at midnight.
15:47:11,508 |-INFO in ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP@343f4d3d - Setting initial period to Wed Jun 27 02:40:17 UTC 2018
15:47:11,525 |-INFO in ch.qos.logback.core.rolling.RollingFileAppender[FILE] - Active log file name: /tmp/spring.log
15:47:11,525 |-INFO in ch.qos.logback.core.rolling.RollingFileAppender[FILE] - File property is set to [/tmp/spring.log]
15:47:11,552 |-INFO in ch.qos.logback.classic.joran.action.RootLoggerAction - Setting level of ROOT logger to INFO
15:47:11,552 |-INFO in ch.qos.logback.core.joran.action.AppenderRefAction - Attaching appender named [CONSOLE] to Logger[ROOT]
15:47:11,557 |-INFO in ch.qos.logback.core.joran.action.AppenderRefAction - Attaching appender named [FILE] to Logger[ROOT]
15:47:11,557 |-INFO in ch.qos.logback.classic.joran.action.ConfigurationAction - End of configuration.
15:47:11,559 |-INFO in ch.qos.logback.classic.joran.JoranConfigurator@53b32d7 - Registering current configuration as safe fallback point

SLF4J: Actual binding is of type [ch.qos.logback.classic.util.ContextSelectorStaticBinder]

  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::        (v2.0.3.RELEASE)

2018-06-28 15:47:14.694  INFO 23371 --- [  restartedMain] io.stackdocker.moocsaas.webapp.App       : Starting App on kubedev-172-17-4-59 with PID 23371 (/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/target/classes started by vagrant in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s)
2018-06-28 15:47:14.733  INFO 23371 --- [  restartedMain] io.stackdocker.moocsaas.webapp.App       : No active profile set, falling back to default profiles: default
2018-06-28 15:47:15.132  INFO 23371 --- [  restartedMain] ConfigServletWebServerApplicationContext : Refreshing org.springframework.boot.web.servlet.context.AnnotationConfigServletWebServerApplicationContext@5d6fe687: startup date [Thu Jun 28 15:47:15 UTC 2018]; root of context hierarchy
2018-06-28 15:47:20.690  INFO 23371 --- [  restartedMain] o.s.b.f.s.DefaultListableBeanFactory     : Overriding bean definition for bean 'catalogRepository' with a different definition: replacing [Generic bean: class [io.stackdocker.moocsaas.webapp.service.CatalogRepository]; scope=singleton; abstract=false; lazyInit=false; autowireMode=0; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=null; factoryMethodName=null; initMethodName=null; destroyMethodName=null; defined in file [/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/java/saas-k8s/web-server/target/classes/io/stackdocker/moocsaas/webapp/service/CatalogRepository.class]] with [Root bean: class [null]; scope=; abstract=false; lazyInit=false; autowireMode=3; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=app; factoryMethodName=catalogRepository; initMethodName=null; destroyMethodName=(inferred); defined in io.stackdocker.moocsaas.webapp.App]
2018-06-28 15:47:25.036  INFO 23371 --- [  restartedMain] o.s.b.w.embedded.tomcat.TomcatWebServer  : Tomcat initialized with port(s): 8080 (http)
2018-06-28 15:47:25.148  INFO 23371 --- [  restartedMain] o.apache.catalina.core.StandardService   : Starting service [Tomcat]
2018-06-28 15:47:25.150  INFO 23371 --- [  restartedMain] org.apache.catalina.core.StandardEngine  : Starting Servlet Engine: Apache Tomcat/8.5.31
2018-06-28 15:47:25.177  INFO 23371 --- [ost-startStop-1] o.a.catalina.core.AprLifecycleListener   : The APR based Apache Tomcat Native library which allows optimal performance in production environments was not found on the java.library.path: [/usr/java/packages/lib/amd64:/usr/lib64:/lib64:/lib:/usr/lib]
2018-06-28 15:47:25.931  INFO 23371 --- [ost-startStop-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring embedded WebApplicationContext
2018-06-28 15:47:25.933  INFO 23371 --- [ost-startStop-1] o.s.web.context.ContextLoader            : Root WebApplicationContext: initialization completed in 10801 ms
2018-06-28 15:47:26.590  INFO 23371 --- [ost-startStop-1] o.s.b.w.servlet.ServletRegistrationBean  : Servlet dispatcherServlet mapped to [/]
2018-06-28 15:47:26.605  INFO 23371 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'characterEncodingFilter' to: [/*]
2018-06-28 15:47:26.606  INFO 23371 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'hiddenHttpMethodFilter' to: [/*]
2018-06-28 15:47:26.608  INFO 23371 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'httpPutFormContentFilter' to: [/*]
2018-06-28 15:47:26.609  INFO 23371 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'requestContextFilter' to: [/*]
2018-06-28 15:47:29.582  INFO 23371 --- [  restartedMain] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/**/favicon.ico] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2018-06-28 15:47:31.362  INFO 23371 --- [  restartedMain] s.w.s.m.m.a.RequestMappingHandlerAdapter : Looking for @ControllerAdvice: org.springframework.boot.web.servlet.context.AnnotationConfigServletWebServerApplicationContext@5d6fe687: startup date [Thu Jun 28 15:47:15 UTC 2018]; root of context hierarchy
###...snippets...
2018-06-28 16:17:06.285  INFO 29000 --- [  restartedMain] o.s.b.w.embedded.tomcat.TomcatWebServer  : Tomcat started on port(s): 8080 (http) with context path ''
2018-06-28 16:17:06.325  INFO 29000 --- [  restartedMain] io.stackdocker.moocsaas.webapp.App       : Started App in 25.211 seconds (JVM running for 28.456)
```

The address is listed at the last 2nd line in Spring Console log

### Docker

Image
```
[vagrant@kubedev-172-17-4-59 saas-k8s]$ docker build --force-rm --no-cache -t docker.io/tangfeixiong/mooc-k8s -f ./web-server/Dockerfile.centos7 ./web-server/
Sending build context to Docker daemon 49.55 MB
Step 1/8 : FROM docker.io/centos:centos7
 ---> ff426288ea90
Step 2/8 : LABEL maintainer "tangfeixiong <tangfx128@gmail.com>" project "https://github.com/tangfeixiong/go-to-kubernetes" name "api" namespace "stackdocker0x2Eio" annotation '{"stackdocker.io/created-by":"n/a"}' tag "centos java springboot tomcat jsp shiro restTemplate"
 ---> Running in 2cbb20546437
 ---> 3f9a902c5bd7
Removing intermediate container 2cbb20546437
Step 3/8 : ARG jarTgt
 ---> Running in 78eda782a823
 ---> 0ba321c236da
Removing intermediate container 78eda782a823
Step 4/8 : ARG javaOpt
 ---> Running in e5b4c8ec7d36
 ---> a7790a4beae2
Removing intermediate container e5b4c8ec7d36
Step 5/8 : COPY ${jarTgt:-/target/web-server*.jar} /web-server.jar
 ---> c45c90835144
Removing intermediate container bc64596aad87
Step 6/8 : ENV JAVA_OPTIONS "${javaOpt:-'-Xms128m -Xmx512m -XX:PermSize=128m -XX:MaxPermSize=256m'}" APISERVER_ADDRESS "http://127.0.0.1:8090" SERVER_PORT "8080"
 ---> Running in 4423a4e15b1f
 ---> 6921f3d13636
Removing intermediate container 4423a4e15b1f
Step 7/8 : RUN set -x     && install_Pkgs="         tar         unzip         bc         which         lsof         java-1.8.0-openjdk-headless     "     && yum install -y $install_Pkgs     && yum clean all -y     && echo
 ---> Running in 12eda5c935b6
+ install_Pkgs='         tar         unzip         bc         which         lsof         java-1.8.0-openjdk-headless     '
+ yum install -y tar unzip bc which lsof java-1.8.0-openjdk-headless
Loaded plugins: fastestmirror, ovl
Determining fastest mirrors
 * base: mirrors.163.com
 * extras: mirrors.nju.edu.cn
 * updates: mirrors.nju.edu.cn
Resolving Dependencies
--> Running transaction check
---> Package bc.x86_64 0:1.06.95-13.el7 will be installed
---> Package java-1.8.0-openjdk-headless.x86_64 1:1.8.0.171-8.b10.el7_5 will be installed
--> Processing Dependency: tzdata-java >= 2015d for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: nss-softokn(x86-64) >= 3.36.0 for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: nss(x86-64) >= 3.36.0 for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: copy-jdk-configs >= 2.2 for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: lksctp-tools(x86-64) for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: libjpeg.so.62(LIBJPEG_6.2)(64bit) for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: jpackage-utils for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: libjpeg.so.62()(64bit) for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
--> Processing Dependency: libfreetype.so.6()(64bit) for package: 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_64
---> Package lsof.x86_64 0:4.87-5.el7 will be installed
---> Package tar.x86_64 2:1.26-32.el7 will be updated
---> Package tar.x86_64 2:1.26-34.el7 will be an update
---> Package unzip.x86_64 0:6.0-19.el7 will be installed
---> Package which.x86_64 0:2.20-7.el7 will be installed
--> Running transaction check
---> Package copy-jdk-configs.noarch 0:3.3-10.el7_5 will be installed
---> Package freetype.x86_64 0:2.4.11-15.el7 will be installed
---> Package javapackages-tools.noarch 0:3.4.1-11.el7 will be installed
--> Processing Dependency: python-javapackages = 3.4.1-11.el7 for package: javapackages-tools-3.4.1-11.el7.noarch
--> Processing Dependency: libxslt for package: javapackages-tools-3.4.1-11.el7.noarch
---> Package libjpeg-turbo.x86_64 0:1.2.90-5.el7 will be installed
---> Package lksctp-tools.x86_64 0:1.0.17-2.el7 will be installed
---> Package nss.x86_64 0:3.28.4-15.el7_4 will be updated
--> Processing Dependency: nss = 3.28.4-15.el7_4 for package: nss-sysinit-3.28.4-15.el7_4.x86_64
--> Processing Dependency: nss(x86-64) = 3.28.4-15.el7_4 for package: nss-tools-3.28.4-15.el7_4.x86_64
---> Package nss.x86_64 0:3.36.0-5.el7_5 will be an update
--> Processing Dependency: nss-util >= 3.36.0-1 for package: nss-3.36.0-5.el7_5.x86_64
--> Processing Dependency: nspr >= 4.19.0 for package: nss-3.36.0-5.el7_5.x86_64
--> Processing Dependency: libnssutil3.so(NSSUTIL_3.31)(64bit) for package: nss-3.36.0-5.el7_5.x86_64
---> Package nss-softokn.x86_64 0:3.28.3-8.el7_4 will be updated
---> Package nss-softokn.x86_64 0:3.36.0-5.el7_5 will be an update
--> Processing Dependency: nss-softokn-freebl(x86-64) >= 3.36.0-5.el7_5 for package: nss-softokn-3.36.0-5.el7_5.x86_64
---> Package tzdata-java.noarch 0:2018e-3.el7 will be installed
--> Running transaction check
---> Package libxslt.x86_64 0:1.1.28-5.el7 will be installed
---> Package nspr.x86_64 0:4.13.1-1.0.el7_3 will be updated
---> Package nspr.x86_64 0:4.19.0-1.el7_5 will be an update
---> Package nss-softokn-freebl.x86_64 0:3.28.3-8.el7_4 will be updated
---> Package nss-softokn-freebl.x86_64 0:3.36.0-5.el7_5 will be an update
---> Package nss-sysinit.x86_64 0:3.28.4-15.el7_4 will be updated
---> Package nss-sysinit.x86_64 0:3.36.0-5.el7_5 will be an update
---> Package nss-tools.x86_64 0:3.28.4-15.el7_4 will be updated
---> Package nss-tools.x86_64 0:3.36.0-5.el7_5 will be an update
---> Package nss-util.x86_64 0:3.28.4-3.el7 will be updated
---> Package nss-util.x86_64 0:3.36.0-1.el7_5 will be an update
---> Package python-javapackages.noarch 0:3.4.1-11.el7 will be installed
--> Processing Dependency: python-lxml for package: python-javapackages-3.4.1-11.el7.noarch
--> Running transaction check
---> Package python-lxml.x86_64 0:3.2.1-4.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package                      Arch    Version                    Repository
                                                                           Size
================================================================================
Installing:
 bc                           x86_64  1.06.95-13.el7             base     115 k
 java-1.8.0-openjdk-headless  x86_64  1:1.8.0.171-8.b10.el7_5    updates   32 M
 lsof                         x86_64  4.87-5.el7                 base     331 k
 unzip                        x86_64  6.0-19.el7                 base     170 k
 which                        x86_64  2.20-7.el7                 base      41 k
Updating:
 tar                          x86_64  2:1.26-34.el7              base     845 k
Installing for dependencies:
 copy-jdk-configs             noarch  3.3-10.el7_5               updates   21 k
 freetype                     x86_64  2.4.11-15.el7              base     392 k
 javapackages-tools           noarch  3.4.1-11.el7               base      73 k
 libjpeg-turbo                x86_64  1.2.90-5.el7               base     134 k
 libxslt                      x86_64  1.1.28-5.el7               base     242 k
 lksctp-tools                 x86_64  1.0.17-2.el7               base      88 k
 python-javapackages          noarch  3.4.1-11.el7               base      31 k
 python-lxml                  x86_64  3.2.1-4.el7                base     758 k
 tzdata-java                  noarch  2018e-3.el7                updates  185 k
Updating for dependencies:
 nspr                         x86_64  4.19.0-1.el7_5             updates  127 k
 nss                          x86_64  3.36.0-5.el7_5             updates  835 k
 nss-softokn                  x86_64  3.36.0-5.el7_5             updates  315 k
 nss-softokn-freebl           x86_64  3.36.0-5.el7_5             updates  222 k
 nss-sysinit                  x86_64  3.36.0-5.el7_5             updates   62 k
 nss-tools                    x86_64  3.36.0-5.el7_5             updates  514 k
 nss-util                     x86_64  3.36.0-1.el7_5             updates   78 k

Transaction Summary
================================================================================
Install  5 Packages (+9 Dependent packages)
Upgrade  1 Package  (+7 Dependent packages)

Total download size: 37 M
Downloading packages:
Delta RPMs disabled because /usr/bin/applydeltarpm not installed.
warning: /var/cache/yum/x86_64/7/base/packages/javapackages-tools-3.4.1-11.el7.noarch.rpm: Header V3 RSA/SHA256 Signature, key ID f4a80eb5: NOKEY
Public key for javapackages-tools-3.4.1-11.el7.noarch.rpm is not installed
Public key for nspr-4.19.0-1.el7_5.x86_64.rpm is not installed
--------------------------------------------------------------------------------
Total                                              4.0 MB/s |  37 MB  00:09     
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Importing GPG key 0xF4A80EB5:
 Userid     : "CentOS-7 Key (CentOS 7 Official Signing Key) <security@centos.org>"
 Fingerprint: 6341 ab27 53d7 8a78 a7c2 7bb1 24c6 a8a7 f4a8 0eb5
 Package    : centos-release-7-4.1708.el7.centos.x86_64 (@CentOS)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Updating   : nspr-4.19.0-1.el7_5.x86_64                                  1/30 
  Updating   : nss-util-3.36.0-1.el7_5.x86_64                              2/30 
  Installing : libxslt-1.1.28-5.el7.x86_64                                 3/30 
  Installing : python-lxml-3.2.1-4.el7.x86_64                              4/30 
  Installing : python-javapackages-3.4.1-11.el7.noarch                     5/30 
  Installing : javapackages-tools-3.4.1-11.el7.noarch                      6/30 
  Updating   : nss-softokn-freebl-3.36.0-5.el7_5.x86_64                    7/30 
  Updating   : nss-softokn-3.36.0-5.el7_5.x86_64                           8/30 
  Updating   : nss-3.36.0-5.el7_5.x86_64                                   9/30 
  Updating   : nss-sysinit-3.36.0-5.el7_5.x86_64                          10/30 
  Installing : tzdata-java-2018e-3.el7.noarch                             11/30 
  Installing : lksctp-tools-1.0.17-2.el7.x86_64                           12/30 
  Installing : copy-jdk-configs-3.3-10.el7_5.noarch                       13/30 
  Installing : freetype-2.4.11-15.el7.x86_64                              14/30 
  Installing : libjpeg-turbo-1.2.90-5.el7.x86_64                          15/30 
  Installing : 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_   16/30 
  Updating   : nss-tools-3.36.0-5.el7_5.x86_64                            17/30 
  Installing : unzip-6.0-19.el7.x86_64                                    18/30 
  Updating   : 2:tar-1.26-34.el7.x86_64                                   19/30 
  Installing : lsof-4.87-5.el7.x86_64                                     20/30 
  Installing : which-2.20-7.el7.x86_64                                    21/30 
install-info: No such file or directory for /usr/share/info/which.info.gz
  Installing : bc-1.06.95-13.el7.x86_64                                   22/30 
  Cleanup    : nss-tools-3.28.4-15.el7_4.x86_64                           23/30 
  Cleanup    : nss-sysinit-3.28.4-15.el7_4.x86_64                         24/30 
  Cleanup    : nss-3.28.4-15.el7_4.x86_64                                 25/30 
  Cleanup    : nss-softokn-3.28.3-8.el7_4.x86_64                          26/30 
  Cleanup    : nss-util-3.28.4-3.el7.x86_64                               27/30 
  Cleanup    : nss-softokn-freebl-3.28.3-8.el7_4.x86_64                   28/30 
  Cleanup    : nspr-4.13.1-1.0.el7_3.x86_64                               29/30 
  Cleanup    : 2:tar-1.26-32.el7.x86_64                                   30/30 
  Verifying  : 1:java-1.8.0-openjdk-headless-1.8.0.171-8.b10.el7_5.x86_    1/30 
  Verifying  : python-javapackages-3.4.1-11.el7.noarch                     2/30 
  Verifying  : libjpeg-turbo-1.2.90-5.el7.x86_64                           3/30 
  Verifying  : nss-tools-3.36.0-5.el7_5.x86_64                             4/30 
  Verifying  : python-lxml-3.2.1-4.el7.x86_64                              5/30 
  Verifying  : freetype-2.4.11-15.el7.x86_64                               6/30 
  Verifying  : bc-1.06.95-13.el7.x86_64                                    7/30 
  Verifying  : nss-util-3.36.0-1.el7_5.x86_64                              8/30 
  Verifying  : copy-jdk-configs-3.3-10.el7_5.noarch                        9/30 
  Verifying  : which-2.20-7.el7.x86_64                                    10/30 
  Verifying  : lsof-4.87-5.el7.x86_64                                     11/30 
  Verifying  : nss-3.36.0-5.el7_5.x86_64                                  12/30 
  Verifying  : lksctp-tools-1.0.17-2.el7.x86_64                           13/30 
  Verifying  : libxslt-1.1.28-5.el7.x86_64                                14/30 
  Verifying  : javapackages-tools-3.4.1-11.el7.noarch                     15/30 
  Verifying  : nss-softokn-freebl-3.36.0-5.el7_5.x86_64                   16/30 
  Verifying  : 2:tar-1.26-34.el7.x86_64                                   17/30 
  Verifying  : nspr-4.19.0-1.el7_5.x86_64                                 18/30 
  Verifying  : tzdata-java-2018e-3.el7.noarch                             19/30 
  Verifying  : unzip-6.0-19.el7.x86_64                                    20/30 
  Verifying  : nss-softokn-3.36.0-5.el7_5.x86_64                          21/30 
  Verifying  : nss-sysinit-3.36.0-5.el7_5.x86_64                          22/30 
  Verifying  : nspr-4.13.1-1.0.el7_3.x86_64                               23/30 
  Verifying  : nss-util-3.28.4-3.el7.x86_64                               24/30 
  Verifying  : nss-softokn-3.28.3-8.el7_4.x86_64                          25/30 
  Verifying  : 2:tar-1.26-32.el7.x86_64                                   26/30 
  Verifying  : nss-sysinit-3.28.4-15.el7_4.x86_64                         27/30 
  Verifying  : nss-3.28.4-15.el7_4.x86_64                                 28/30 
  Verifying  : nss-tools-3.28.4-15.el7_4.x86_64                           29/30 
  Verifying  : nss-softokn-freebl-3.28.3-8.el7_4.x86_64                   30/30 

Installed:
  bc.x86_64 0:1.06.95-13.el7                                                    
  java-1.8.0-openjdk-headless.x86_64 1:1.8.0.171-8.b10.el7_5                    
  lsof.x86_64 0:4.87-5.el7                                                      
  unzip.x86_64 0:6.0-19.el7                                                     
  which.x86_64 0:2.20-7.el7                                                     

Dependency Installed:
  copy-jdk-configs.noarch 0:3.3-10.el7_5    freetype.x86_64 0:2.4.11-15.el7    
  javapackages-tools.noarch 0:3.4.1-11.el7  libjpeg-turbo.x86_64 0:1.2.90-5.el7
  libxslt.x86_64 0:1.1.28-5.el7             lksctp-tools.x86_64 0:1.0.17-2.el7 
  python-javapackages.noarch 0:3.4.1-11.el7 python-lxml.x86_64 0:3.2.1-4.el7   
  tzdata-java.noarch 0:2018e-3.el7         

Updated:
  tar.x86_64 2:1.26-34.el7                                                      

Dependency Updated:
  nspr.x86_64 0:4.19.0-1.el7_5                                                  
  nss.x86_64 0:3.36.0-5.el7_5                                                   
  nss-softokn.x86_64 0:3.36.0-5.el7_5                                           
  nss-softokn-freebl.x86_64 0:3.36.0-5.el7_5                                    
  nss-sysinit.x86_64 0:3.36.0-5.el7_5                                           
  nss-tools.x86_64 0:3.36.0-5.el7_5                                             
  nss-util.x86_64 0:3.36.0-1.el7_5                                              

Complete!
+ yum clean all -y
Loaded plugins: fastestmirror, ovl
Cleaning repos: base extras updates
Cleaning up everything
Maybe you want: rm -rf /var/cache/yum, to also free up space taken by orphaned data from disabled or removed repos
Cleaning up list of fastest mirrors
+ echo

 ---> 319eee5d398c
Removing intermediate container 12eda5c935b6
Step 8/8 : CMD java -Djava.security.egd=file:/dev/./urandom $JAVA_OPTIONS -jar /web-server.jar
 ---> Running in 82520e397c90
 ---> 1857311a4041
Removing intermediate container 82520e397c90
Successfully built 1857311a4041
```

```
[vagrant@kubedev-172-17-4-59 saas-k8s]$ docker images docker.io/tangfeixiong/mooc-k8s
REPOSITORY                        TAG                 IMAGE ID            CREATED             SIZE
docker.io/tangfeixiong/mooc-k8s   latest              1857311a4041        14 minutes ago      420 MB
```

### Kubernetes

__TBC__