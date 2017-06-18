# Hand-on: SpringBoot example for private Docker Registry

## Contents of table

* java
* 

## Build Ship and Run

Clone sample repository
```
[vagrant@localhost ~]$ git clone https://github.com/tangfeixiong/osev3-examples /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples
```

## Java

### Build

Package with `mvn`
```
[vagrant@localhost sample-microservices-springboot]$ mvn clean package
[INFO] Scanning for projects...
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Build Order:
[INFO] 
[INFO] web
[INFO] repositories-mem
[INFO] microservices-demo
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building web 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.5:clean (default-clean) @ web ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/web/target
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ web ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 1 resource
[INFO] Copying 10 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ web ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 5 source files to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/web/target/classes
[WARNING] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/web/src/main/java/com/openshift/evangelists/microservices/web/client/RestMessageRepositoryClient.java: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/web/src/main/java/com/openshift/evangelists/microservices/web/client/RestMessageRepositoryClient.java使用了未经检查或不安全的操作。
[WARNING] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/web/src/main/java/com/openshift/evangelists/microservices/web/client/RestMessageRepositoryClient.java: 有关详细信息, 请使用 -Xlint:unchecked 重新编译。
[INFO] 
[INFO] --- maven-resources-plugin:2.6:testResources (default-testResources) @ web ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/web/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:testCompile (default-testCompile) @ web ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-surefire-plugin:2.17:test (default-test) @ web ---
[INFO] No tests to run.
[INFO] 
[INFO] --- maven-jar-plugin:2.5:jar (default-jar) @ web ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/web/target/web.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.2.5.RELEASE:repackage (default) @ web ---
Downloading: https://repo.maven.apache.org/maven2/org/springframework/springloaded/1.2.4.RELEASE/springloaded-1.2.4.RELEASE.pom
Downloaded: https://repo.maven.apache.org/maven2/org/springframework/springloaded/1.2.4.RELEASE/springloaded-1.2.4.RELEASE.pom (2 KB at 0.6 KB/sec)
Downloading: https://repo.maven.apache.org/maven2/org/springframework/boot/spring-boot-loader-tools/1.2.5.RELEASE/spring-boot-loader-tools-1.2.5.RELEASE.pom
Downloaded: https://repo.maven.apache.org/maven2/org/springframework/boot/spring-boot-loader-tools/1.2.5.RELEASE/spring-boot-loader-tools-1.2.5.RELEASE.pom (4 KB at 6.0 KB/sec)
Downloading: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-api/0.9.1.v20140329/aether-api-0.9.1.v20140329.pom
Downloaded: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-api/0.9.1.v20140329/aether-api-0.9.1.v20140329.pom (2 KB at 3.0 KB/sec)
Downloading: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether/0.9.1.v20140329/aether-0.9.1.v20140329.pom
Downloaded: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether/0.9.1.v20140329/aether-0.9.1.v20140329.pom (30 KB at 28.4 KB/sec)
Downloading: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-spi/0.9.1.v20140329/aether-spi-0.9.1.v20140329.pom
Downloaded: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-spi/0.9.1.v20140329/aether-spi-0.9.1.v20140329.pom (3 KB at 3.2 KB/sec)
Downloading: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-util/0.9.1.v20140329/aether-util-0.9.1.v20140329.pom
Downloaded: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-util/0.9.1.v20140329/aether-util-0.9.1.v20140329.pom (3 KB at 3.5 KB/sec)
Downloading: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-impl/0.9.1.v20140329/aether-impl-0.9.1.v20140329.pom
Downloaded: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-impl/0.9.1.v20140329/aether-impl-0.9.1.v20140329.pom (4 KB at 6.4 KB/sec)
Downloading: https://repo.maven.apache.org/maven2/org/springframework/springloaded/1.2.4.RELEASE/springloaded-1.2.4.RELEASE.jar
Downloading: https://repo.maven.apache.org/maven2/org/springframework/boot/spring-boot-loader-tools/1.2.5.RELEASE/spring-boot-loader-tools-1.2.5.RELEASE.jar
Downloading: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-spi/0.9.1.v20140329/aether-spi-0.9.1.v20140329.jar
Downloading: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-impl/0.9.1.v20140329/aether-impl-0.9.1.v20140329.jar
Downloading: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-api/0.9.1.v20140329/aether-api-0.9.1.v20140329.jar
Downloaded: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-spi/0.9.1.v20140329/aether-spi-0.9.1.v20140329.jar (32 KB at 13.6 KB/sec)
Downloading: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-util/0.9.1.v20140329/aether-util-0.9.1.v20140329.jar
Downloaded: https://repo.maven.apache.org/maven2/org/springframework/boot/spring-boot-loader-tools/1.2.5.RELEASE/spring-boot-loader-tools-1.2.5.RELEASE.jar (125 KB at 44.1 KB/sec)
Downloaded: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-api/0.9.1.v20140329/aether-api-0.9.1.v20140329.jar (134 KB at 40.9 KB/sec)
Downloaded: https://repo.maven.apache.org/maven2/org/springframework/springloaded/1.2.4.RELEASE/springloaded-1.2.4.RELEASE.jar (418 KB at 123.9 KB/sec)
Downloaded: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-impl/0.9.1.v20140329/aether-impl-0.9.1.v20140329.jar (172 KB at 50.8 KB/sec)
Downloaded: https://repo.maven.apache.org/maven2/org/eclipse/aether/aether-util/0.9.1.v20140329/aether-util-0.9.1.v20140329.jar (143 KB at 38.1 KB/sec)
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building repositories-mem 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.5:clean (default-clean) @ repositories-mem ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/repositories-mem/target
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ repositories-mem ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 1 resource
[INFO] Copying 0 resource
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ repositories-mem ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 5 source files to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/repositories-mem/target/classes
[INFO] 
[INFO] --- maven-resources-plugin:2.6:testResources (default-testResources) @ repositories-mem ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/repositories-mem/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:testCompile (default-testCompile) @ repositories-mem ---
[INFO] No sources to compile
[INFO] 
[INFO] --- maven-surefire-plugin:2.17:test (default-test) @ repositories-mem ---
[INFO] No tests to run.
[INFO] 
[INFO] --- maven-jar-plugin:2.5:jar (default-jar) @ repositories-mem ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/repositories-mem/target/repositories-mem.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.2.5.RELEASE:repackage (default) @ repositories-mem ---
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building microservices-demo 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.5:clean (default-clean) @ microservices-demo-springboot ---
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Summary:
[INFO] 
[INFO] web ................................................ SUCCESS [ 25.963 s]
[INFO] repositories-mem ................................... SUCCESS [  7.348 s]
[INFO] microservices-demo ................................. SUCCESS [  0.009 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 33.976 s
[INFO] Finished at: 2017-06-16T12:55:06+00:00
[INFO] Final Memory: 31M/253M
[INFO] ------------------------------------------------------------------------
```

### Try

Exec web
```
[vagrant@localhost sample-microservices-springboot]$ java -jar web/target/web.jar --remoteHost=localhost --remotePort=9091 --server.port=8091

  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::        (v1.2.5.RELEASE)

2017-06-16 12:57:08.851  INFO 10019 --- [           main] c.o.e.m.web.SampleWebUIApplication       : Starting SampleWebUIApplication v0.0.1-SNAPSHOT on localhost.localdomain with PID 10019 (/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/web/target/web.jar started by vagrant in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot)
2017-06-16 12:57:08.992  INFO 10019 --- [           main] ationConfigEmbeddedWebApplicationContext : Refreshing org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@38644dcf: startup date [Fri Jun 16 12:57:08 UTC 2017]; root of context hierarchy
2017-06-16 12:57:10.275  INFO 10019 --- [           main] o.s.b.f.s.DefaultListableBeanFactory     : Overriding bean definition for bean 'beanNameViewResolver': replacing [Root bean: class [null]; scope=; abstract=false; lazyInit=false; autowireMode=3; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=org.springframework.boot.autoconfigure.web.ErrorMvcAutoConfiguration$WhitelabelErrorViewConfiguration; factoryMethodName=beanNameViewResolver; initMethodName=null; destroyMethodName=(inferred); defined in class path resource [org/springframework/boot/autoconfigure/web/ErrorMvcAutoConfiguration$WhitelabelErrorViewConfiguration.class]] with [Root bean: class [null]; scope=; abstract=false; lazyInit=false; autowireMode=3; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=org.springframework.boot.autoconfigure.web.WebMvcAutoConfiguration$WebMvcAutoConfigurationAdapter; factoryMethodName=beanNameViewResolver; initMethodName=null; destroyMethodName=(inferred); defined in class path resource [org/springframework/boot/autoconfigure/web/WebMvcAutoConfiguration$WebMvcAutoConfigurationAdapter.class]]
2017-06-16 12:57:12.374  INFO 10019 --- [           main] s.b.c.e.t.TomcatEmbeddedServletContainer : Tomcat initialized with port(s): 8091 (http)
2017-06-16 12:57:12.763  INFO 10019 --- [           main] o.apache.catalina.core.StandardService   : Starting service Tomcat
2017-06-16 12:57:12.768  INFO 10019 --- [           main] org.apache.catalina.core.StandardEngine  : Starting Servlet Engine: Apache Tomcat/8.0.23
2017-06-16 12:57:13.223  INFO 10019 --- [ost-startStop-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring embedded WebApplicationContext
2017-06-16 12:57:13.224  INFO 10019 --- [ost-startStop-1] o.s.web.context.ContextLoader            : Root WebApplicationContext: initialization completed in 4234 ms
2017-06-16 12:57:14.485  INFO 10019 --- [ost-startStop-1] o.s.b.c.e.ServletRegistrationBean        : Mapping servlet: 'dispatcherServlet' to [/]
2017-06-16 12:57:14.500  INFO 10019 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'characterEncodingFilter' to: [/*]
2017-06-16 12:57:14.502  INFO 10019 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'hiddenHttpMethodFilter' to: [/*]
2017-06-16 12:57:15.272  INFO 10019 --- [           main] s.w.s.m.m.a.RequestMappingHandlerAdapter : Looking for @ControllerAdvice: org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@38644dcf: startup date [Fri Jun 16 12:57:08 UTC 2017]; root of context hierarchy
2017-06-16 12:57:15.586  INFO 10019 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{id}],methods=[GET]}" onto public org.springframework.web.servlet.ModelAndView com.openshift.evangelists.microservices.web.mvc.WebMessageController.view(com.openshift.evangelists.microservices.web.api.Message)
2017-06-16 12:57:15.587  INFO 10019 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/],methods=[GET],params=[form]}" onto public java.lang.String com.openshift.evangelists.microservices.web.mvc.WebMessageController.createForm(com.openshift.evangelists.microservices.web.api.Message)
2017-06-16 12:57:15.587  INFO 10019 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/foo]}" onto public java.lang.String com.openshift.evangelists.microservices.web.mvc.WebMessageController.foo()
2017-06-16 12:57:15.588  INFO 10019 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/]}" onto public org.springframework.web.servlet.ModelAndView com.openshift.evangelists.microservices.web.mvc.WebMessageController.list()
2017-06-16 12:57:15.588  INFO 10019 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/],methods=[POST]}" onto public org.springframework.web.servlet.ModelAndView com.openshift.evangelists.microservices.web.mvc.WebMessageController.create(com.openshift.evangelists.microservices.web.api.Message,org.springframework.validation.BindingResult,org.springframework.web.servlet.mvc.support.RedirectAttributes)
2017-06-16 12:57:15.590  INFO 10019 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error],produces=[text/html]}" onto public org.springframework.web.servlet.ModelAndView org.springframework.boot.autoconfigure.web.BasicErrorController.errorHtml(javax.servlet.http.HttpServletRequest)
2017-06-16 12:57:15.590  INFO 10019 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error]}" onto public org.springframework.http.ResponseEntity<java.util.Map<java.lang.String, java.lang.Object>> org.springframework.boot.autoconfigure.web.BasicErrorController.error(javax.servlet.http.HttpServletRequest)
2017-06-16 12:57:15.637  INFO 10019 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/webjars/**] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2017-06-16 12:57:15.638  INFO 10019 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/**] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2017-06-16 12:57:15.724  INFO 10019 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/**/favicon.ico] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2017-06-16 12:57:16.115  INFO 10019 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on startup
2017-06-16 12:57:16.279  INFO 10019 --- [           main] s.b.c.e.t.TomcatEmbeddedServletContainer : Tomcat started on port(s): 8091 (http)
2017-06-16 12:57:16.284  INFO 10019 --- [           main] c.o.e.m.web.SampleWebUIApplication       : Started SampleWebUIApplication in 8.137 seconds (JVM running for 10.938)
2017-06-16 12:57:47.178  INFO 10019 --- [nio-8091-exec-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring FrameworkServlet 'dispatcherServlet'
2017-06-16 12:57:47.180  INFO 10019 --- [nio-8091-exec-1] o.s.web.servlet.DispatcherServlet        : FrameworkServlet 'dispatcherServlet': initialization started
2017-06-16 12:57:47.211  INFO 10019 --- [nio-8091-exec-1] o.s.web.servlet.DispatcherServlet        : FrameworkServlet 'dispatcherServlet': initialization completed in 29 ms
WebMessageController.list
RestMessageRepositoryClient to localhost:9091
RestMessageRepositoryClient.findAll
Remote server is not available
```

Command `curl` or Browse
```
fanhonglingdeMacBook-Pro:examples fanhongling$ curl 172.17.4.50:8091
<!DOCTYPE html>

<html>
  <head>
    <title>Messages : View all</title>
    
    <link rel="stylesheet" href="/css/bootstrap.min.css" />
  
    
  </head>
  <body>
    <div class="container">
      <div class="navbar">
        <div class="navbar-inner">
          <a class="brand" href="https://github.com/ultraq/thymeleaf-layout-dialect">
              Thymeleaf - Layout
          </a>
          <ul class="nav">
            <li>
              <a href="/">
                Messages
              </a>
            </li>
          </ul>
        </div>
      </div>
      <h1>Messages : View all</h1>
      <div class="container">
      <div class="pull-right">
        <a href="/?form">Create Message</a>
       </div>
      <table class="table table-bordered table-striped">
        <thead>
          <tr>
            <td>ID</td>
            <td>Created</td>
            <td>Summary</td>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td colspan="3">
              No messages
            </td>
          </tr>
          
        </tbody>
      </table>
    </div>
    </div>
  </body>
</html>
```

Terminate with _ctrl-c_
```
^C2017-06-16 12:58:03.133  INFO 10019 --- [       Thread-2] ationConfigEmbeddedWebApplicationContext : Closing org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@38644dcf: startup date [Fri Jun 16 12:57:08 UTC 2017]; root of context hierarchy
2017-06-16 12:58:03.137  INFO 10019 --- [       Thread-2] o.s.j.e.a.AnnotationMBeanExporter        : Unregistering JMX-exposed beans on shutdown
```

Service repo
```
[vagrant@localhost sample-microservices-springboot]$ java -jar repositories-mem/target/repositories-mem.jar --server.port=9091

  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::        (v1.2.5.RELEASE)

2017-06-17 06:03:05.351  INFO 16589 --- [           main] c.o.e.m.r.InMemoryRepositoryApplication  : Starting InMemoryRepositoryApplication v0.0.1-SNAPSHOT on localhost.localdomain with PID 16589 (/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/repositories-mem/target/repositories-mem.jar started by vagrant in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot)
2017-06-17 06:03:05.484  INFO 16589 --- [           main] ationConfigEmbeddedWebApplicationContext : Refreshing org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@278b5329: startup date [Sat Jun 17 06:03:05 UTC 2017]; root of context hierarchy
2017-06-17 06:03:07.602  INFO 16589 --- [           main] o.s.b.f.s.DefaultListableBeanFactory     : Overriding bean definition for bean 'beanNameViewResolver': replacing [Root bean: class [null]; scope=; abstract=false; lazyInit=false; autowireMode=3; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=org.springframework.boot.autoconfigure.web.ErrorMvcAutoConfiguration$WhitelabelErrorViewConfiguration; factoryMethodName=beanNameViewResolver; initMethodName=null; destroyMethodName=(inferred); defined in class path resource [org/springframework/boot/autoconfigure/web/ErrorMvcAutoConfiguration$WhitelabelErrorViewConfiguration.class]] with [Root bean: class [null]; scope=; abstract=false; lazyInit=false; autowireMode=3; dependencyCheck=0; autowireCandidate=true; primary=false; factoryBeanName=org.springframework.boot.autoconfigure.web.WebMvcAutoConfiguration$WebMvcAutoConfigurationAdapter; factoryMethodName=beanNameViewResolver; initMethodName=null; destroyMethodName=(inferred); defined in class path resource [org/springframework/boot/autoconfigure/web/WebMvcAutoConfiguration$WebMvcAutoConfigurationAdapter.class]]
2017-06-17 06:03:10.207  INFO 16589 --- [           main] s.b.c.e.t.TomcatEmbeddedServletContainer : Tomcat initialized with port(s): 9091 (http)
2017-06-17 06:03:10.685  INFO 16589 --- [           main] o.apache.catalina.core.StandardService   : Starting service Tomcat
2017-06-17 06:03:10.691  INFO 16589 --- [           main] org.apache.catalina.core.StandardEngine  : Starting Servlet Engine: Apache Tomcat/8.0.23
2017-06-17 06:03:11.019  INFO 16589 --- [ost-startStop-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring embedded WebApplicationContext
2017-06-17 06:03:11.020  INFO 16589 --- [ost-startStop-1] o.s.web.context.ContextLoader            : Root WebApplicationContext: initialization completed in 5538 ms
2017-06-17 06:03:13.426  INFO 16589 --- [ost-startStop-1] o.s.b.c.e.ServletRegistrationBean        : Mapping servlet: 'dispatcherServlet' to [/]
2017-06-17 06:03:13.444  INFO 16589 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'metricFilter' to: [/*]
2017-06-17 06:03:13.445  INFO 16589 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'characterEncodingFilter' to: [/*]
2017-06-17 06:03:13.445  INFO 16589 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'hiddenHttpMethodFilter' to: [/*]
2017-06-17 06:03:13.446  INFO 16589 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'webRequestLoggingFilter' to: [/*]
2017-06-17 06:03:13.446  INFO 16589 --- [ost-startStop-1] o.s.b.c.embedded.FilterRegistrationBean  : Mapping filter: 'applicationContextIdFilter' to: [/*]
InMemoryMessageRepository.save[pre]: Message{id=null, text='Hello', summary='World', created=java.util.GregorianCalendar[time=1497679393645,areFieldsSet=true,areAllFieldsSet=true,lenient=true,zone=sun.util.calendar.ZoneInfo[id="UTC",offset=0,dstSavings=0,useDaylight=false,transitions=0,lastRule=null],firstDayOfWeek=1,minimalDaysInFirstWeek=1,ERA=1,YEAR=2017,MONTH=5,WEEK_OF_YEAR=24,WEEK_OF_MONTH=3,DAY_OF_MONTH=17,DAY_OF_YEAR=168,DAY_OF_WEEK=7,DAY_OF_WEEK_IN_MONTH=3,AM_PM=0,HOUR=6,HOUR_OF_DAY=6,MINUTE=3,SECOND=13,MILLISECOND=645,ZONE_OFFSET=0,DST_OFFSET=0]}
InMemoryMessageRepository.save[post]: Message{id=1, text='Hello', summary='World', created=java.util.GregorianCalendar[time=1497679393645,areFieldsSet=true,areAllFieldsSet=true,lenient=true,zone=sun.util.calendar.ZoneInfo[id="UTC",offset=0,dstSavings=0,useDaylight=false,transitions=0,lastRule=null],firstDayOfWeek=1,minimalDaysInFirstWeek=1,ERA=1,YEAR=2017,MONTH=5,WEEK_OF_YEAR=24,WEEK_OF_MONTH=3,DAY_OF_MONTH=17,DAY_OF_YEAR=168,DAY_OF_WEEK=7,DAY_OF_WEEK_IN_MONTH=3,AM_PM=0,HOUR=6,HOUR_OF_DAY=6,MINUTE=3,SECOND=13,MILLISECOND=645,ZONE_OFFSET=0,DST_OFFSET=0]}
InMemoryMessageRepository.save[pre]: Message{id=null, text='Hi', summary='Universe', created=java.util.GregorianCalendar[time=1497679393645,areFieldsSet=true,areAllFieldsSet=true,lenient=true,zone=sun.util.calendar.ZoneInfo[id="UTC",offset=0,dstSavings=0,useDaylight=false,transitions=0,lastRule=null],firstDayOfWeek=1,minimalDaysInFirstWeek=1,ERA=1,YEAR=2017,MONTH=5,WEEK_OF_YEAR=24,WEEK_OF_MONTH=3,DAY_OF_MONTH=17,DAY_OF_YEAR=168,DAY_OF_WEEK=7,DAY_OF_WEEK_IN_MONTH=3,AM_PM=0,HOUR=6,HOUR_OF_DAY=6,MINUTE=3,SECOND=13,MILLISECOND=645,ZONE_OFFSET=0,DST_OFFSET=0]}
InMemoryMessageRepository.save[post]: Message{id=2, text='Hi', summary='Universe', created=java.util.GregorianCalendar[time=1497679393645,areFieldsSet=true,areAllFieldsSet=true,lenient=true,zone=sun.util.calendar.ZoneInfo[id="UTC",offset=0,dstSavings=0,useDaylight=false,transitions=0,lastRule=null],firstDayOfWeek=1,minimalDaysInFirstWeek=1,ERA=1,YEAR=2017,MONTH=5,WEEK_OF_YEAR=24,WEEK_OF_MONTH=3,DAY_OF_MONTH=17,DAY_OF_YEAR=168,DAY_OF_WEEK=7,DAY_OF_WEEK_IN_MONTH=3,AM_PM=0,HOUR=6,HOUR_OF_DAY=6,MINUTE=3,SECOND=13,MILLISECOND=645,ZONE_OFFSET=0,DST_OFFSET=0]}
InMemoryMessageRepository.save[pre]: Message{id=null, text='Hola', summary='OpenShift', created=java.util.GregorianCalendar[time=1497679393646,areFieldsSet=true,areAllFieldsSet=true,lenient=true,zone=sun.util.calendar.ZoneInfo[id="UTC",offset=0,dstSavings=0,useDaylight=false,transitions=0,lastRule=null],firstDayOfWeek=1,minimalDaysInFirstWeek=1,ERA=1,YEAR=2017,MONTH=5,WEEK_OF_YEAR=24,WEEK_OF_MONTH=3,DAY_OF_MONTH=17,DAY_OF_YEAR=168,DAY_OF_WEEK=7,DAY_OF_WEEK_IN_MONTH=3,AM_PM=0,HOUR=6,HOUR_OF_DAY=6,MINUTE=3,SECOND=13,MILLISECOND=646,ZONE_OFFSET=0,DST_OFFSET=0]}
InMemoryMessageRepository.save[post]: Message{id=3, text='Hola', summary='OpenShift', created=java.util.GregorianCalendar[time=1497679393646,areFieldsSet=true,areAllFieldsSet=true,lenient=true,zone=sun.util.calendar.ZoneInfo[id="UTC",offset=0,dstSavings=0,useDaylight=false,transitions=0,lastRule=null],firstDayOfWeek=1,minimalDaysInFirstWeek=1,ERA=1,YEAR=2017,MONTH=5,WEEK_OF_YEAR=24,WEEK_OF_MONTH=3,DAY_OF_MONTH=17,DAY_OF_YEAR=168,DAY_OF_WEEK=7,DAY_OF_WEEK_IN_MONTH=3,AM_PM=0,HOUR=6,HOUR_OF_DAY=6,MINUTE=3,SECOND=13,MILLISECOND=646,ZONE_OFFSET=0,DST_OFFSET=0]}
2017-06-17 06:03:13.926  INFO 16589 --- [           main] o.s.b.f.config.PropertiesFactoryBean     : Loading properties file from class path resource [rest-default-messages.properties]
2017-06-17 06:03:14.204  INFO 16589 --- [           main] s.w.s.m.m.a.RequestMappingHandlerAdapter : Looking for @ControllerAdvice: org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@278b5329: startup date [Sat Jun 17 06:03:05 UTC 2017]; root of context hierarchy
2017-06-17 06:03:14.439  INFO 16589 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{id}],methods=[GET]}" onto public com.openshift.evangelists.microservices.api.Message com.openshift.evangelists.microservices.repository.rest.InMemoryRepositoryRestMessageController.getById(java.lang.Long)
2017-06-17 06:03:14.440  INFO 16589 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/],methods=[GET]}" onto public java.lang.Iterable<com.openshift.evangelists.microservices.api.Message> com.openshift.evangelists.microservices.repository.rest.InMemoryRepositoryRestMessageController.getAll()
2017-06-17 06:03:14.441  INFO 16589 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{id}],methods=[PUT]}" onto public com.openshift.evangelists.microservices.api.Message com.openshift.evangelists.microservices.repository.rest.InMemoryRepositoryRestMessageController.update(java.lang.String,com.openshift.evangelists.microservices.api.Message)
2017-06-17 06:03:14.442  INFO 16589 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{id}],methods=[DELETE]}" onto public void com.openshift.evangelists.microservices.repository.rest.InMemoryRepositoryRestMessageController.delete(java.lang.Long)
2017-06-17 06:03:14.443  INFO 16589 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/],methods=[POST]}" onto public com.openshift.evangelists.microservices.api.Message com.openshift.evangelists.microservices.repository.rest.InMemoryRepositoryRestMessageController.create(com.openshift.evangelists.microservices.api.Message)
2017-06-17 06:03:14.445  INFO 16589 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error],produces=[text/html]}" onto public org.springframework.web.servlet.ModelAndView org.springframework.boot.autoconfigure.web.BasicErrorController.errorHtml(javax.servlet.http.HttpServletRequest)
2017-06-17 06:03:14.446  INFO 16589 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error]}" onto public org.springframework.http.ResponseEntity<java.util.Map<java.lang.String, java.lang.Object>> org.springframework.boot.autoconfigure.web.BasicErrorController.error(javax.servlet.http.HttpServletRequest)
2017-06-17 06:03:14.513  INFO 16589 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/webjars/**] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2017-06-17 06:03:14.514  INFO 16589 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/**] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2017-06-17 06:03:14.616  INFO 16589 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/**/favicon.ico] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2017-06-17 06:03:15.315  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/mappings],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-06-17 06:03:15.317  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/dump],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-06-17 06:03:15.318  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/configprops],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-06-17 06:03:15.319  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/beans],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-06-17 06:03:15.320  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/info],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-06-17 06:03:15.320  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/autoconfig],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-06-17 06:03:15.321  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/env/{name:.*}],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EnvironmentMvcEndpoint.value(java.lang.String)
2017-06-17 06:03:15.322  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/env],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-06-17 06:03:15.324  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/metrics/{name:.*}],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.MetricsMvcEndpoint.value(java.lang.String)
2017-06-17 06:03:15.324  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/metrics],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-06-17 06:03:15.326  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/health]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.HealthMvcEndpoint.invoke(java.security.Principal)
2017-06-17 06:03:15.329  INFO 16589 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/trace],methods=[GET]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-06-17 06:03:15.487  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerAdapter   : Looking for @ControllerAdvice: org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@278b5329: startup date [Sat Jun 17 06:03:05 UTC 2017]; root of context hierarchy
2017-06-17 06:03:15.537  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}],methods=[GET],produces=[application/x-spring-data-compact+json || text/uri-list]}" onto public org.springframework.hateoas.Resources<?> org.springframework.data.rest.webmvc.RepositoryEntityController.getCollectionResourceCompact(org.springframework.data.rest.webmvc.RootResourceInformation,org.springframework.data.rest.webmvc.support.DefaultedPageable,org.springframework.data.domain.Sort,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler) throws org.springframework.data.rest.webmvc.ResourceNotFoundException,org.springframework.web.HttpRequestMethodNotSupportedException
2017-06-17 06:03:15.544  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}],methods=[POST]}" onto public org.springframework.http.ResponseEntity<org.springframework.hateoas.ResourceSupport> org.springframework.data.rest.webmvc.RepositoryEntityController.postCollectionResource(org.springframework.data.rest.webmvc.RootResourceInformation,org.springframework.data.rest.webmvc.PersistentEntityResource,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler) throws org.springframework.web.HttpRequestMethodNotSupportedException
2017-06-17 06:03:15.547  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}],methods=[OPTIONS]}" onto public org.springframework.http.ResponseEntity<?> org.springframework.data.rest.webmvc.RepositoryEntityController.optionsForItemResource(org.springframework.data.rest.webmvc.RootResourceInformation)
2017-06-17 06:03:15.547  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}],methods=[HEAD]}" onto public org.springframework.http.ResponseEntity<?> org.springframework.data.rest.webmvc.RepositoryEntityController.headForItemResource(org.springframework.data.rest.webmvc.RootResourceInformation,java.io.Serializable) throws org.springframework.web.HttpRequestMethodNotSupportedException
2017-06-17 06:03:15.547  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}],methods=[GET]}" onto public org.springframework.http.ResponseEntity<org.springframework.hateoas.Resource<?>> org.springframework.data.rest.webmvc.RepositoryEntityController.getItemResource(org.springframework.data.rest.webmvc.RootResourceInformation,java.io.Serializable,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler) throws org.springframework.web.HttpRequestMethodNotSupportedException
2017-06-17 06:03:15.548  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}],methods=[PUT]}" onto public org.springframework.http.ResponseEntity<? extends org.springframework.hateoas.ResourceSupport> org.springframework.data.rest.webmvc.RepositoryEntityController.putItemResource(org.springframework.data.rest.webmvc.RootResourceInformation,org.springframework.data.rest.webmvc.PersistentEntityResource,java.io.Serializable,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler) throws org.springframework.web.HttpRequestMethodNotSupportedException
2017-06-17 06:03:15.548  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}],methods=[PATCH]}" onto public org.springframework.http.ResponseEntity<org.springframework.hateoas.ResourceSupport> org.springframework.data.rest.webmvc.RepositoryEntityController.patchItemResource(org.springframework.data.rest.webmvc.RootResourceInformation,org.springframework.data.rest.webmvc.PersistentEntityResource,java.io.Serializable,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler) throws org.springframework.web.HttpRequestMethodNotSupportedException,org.springframework.data.rest.webmvc.ResourceNotFoundException
2017-06-17 06:03:15.550  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}],methods=[DELETE]}" onto public org.springframework.http.ResponseEntity<?> org.springframework.data.rest.webmvc.RepositoryEntityController.deleteItemResource(org.springframework.data.rest.webmvc.RootResourceInformation,java.io.Serializable) throws org.springframework.data.rest.webmvc.ResourceNotFoundException,org.springframework.web.HttpRequestMethodNotSupportedException
2017-06-17 06:03:15.550  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}],methods=[OPTIONS]}" onto public org.springframework.http.ResponseEntity<?> org.springframework.data.rest.webmvc.RepositoryEntityController.optionsForCollectionResource(org.springframework.data.rest.webmvc.RootResourceInformation)
2017-06-17 06:03:15.551  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}],methods=[HEAD]}" onto public org.springframework.http.ResponseEntity<?> org.springframework.data.rest.webmvc.RepositoryEntityController.headCollectionResource(org.springframework.data.rest.webmvc.RootResourceInformation) throws org.springframework.web.HttpRequestMethodNotSupportedException
2017-06-17 06:03:15.551  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}],methods=[GET]}" onto public org.springframework.hateoas.Resources<?> org.springframework.data.rest.webmvc.RepositoryEntityController.getCollectionResource(org.springframework.data.rest.webmvc.RootResourceInformation,org.springframework.data.rest.webmvc.support.DefaultedPageable,org.springframework.data.domain.Sort,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler) throws org.springframework.data.rest.webmvc.ResourceNotFoundException,org.springframework.web.HttpRequestMethodNotSupportedException
2017-06-17 06:03:15.552  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/],methods=[GET]}" onto public org.springframework.http.HttpEntity<org.springframework.data.rest.webmvc.RepositoryLinksResource> org.springframework.data.rest.webmvc.RepositoryController.listRepositories()
2017-06-17 06:03:15.552  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/],methods=[OPTIONS]}" onto public org.springframework.http.HttpEntity<?> org.springframework.data.rest.webmvc.RepositoryController.optionsForRepositories()
2017-06-17 06:03:15.553  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/],methods=[HEAD]}" onto public org.springframework.http.ResponseEntity<?> org.springframework.data.rest.webmvc.RepositoryController.headForRepositories()
2017-06-17 06:03:15.561  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}/{property}],methods=[GET]}" onto public org.springframework.http.ResponseEntity<org.springframework.hateoas.ResourceSupport> org.springframework.data.rest.webmvc.RepositoryPropertyReferenceController.followPropertyReference(org.springframework.data.rest.webmvc.RootResourceInformation,java.io.Serializable,java.lang.String,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler) throws java.lang.Exception
2017-06-17 06:03:15.561  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}/{property}/{propertyId}],methods=[GET]}" onto public org.springframework.http.ResponseEntity<org.springframework.hateoas.ResourceSupport> org.springframework.data.rest.webmvc.RepositoryPropertyReferenceController.followPropertyReference(org.springframework.data.rest.webmvc.RootResourceInformation,java.io.Serializable,java.lang.String,java.lang.String,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler) throws java.lang.Exception
2017-06-17 06:03:15.566  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}/{property}],methods=[DELETE]}" onto public org.springframework.http.ResponseEntity<? extends org.springframework.hateoas.ResourceSupport> org.springframework.data.rest.webmvc.RepositoryPropertyReferenceController.deletePropertyReference(org.springframework.data.rest.webmvc.RootResourceInformation,java.io.Serializable,java.lang.String) throws java.lang.Exception
2017-06-17 06:03:15.566  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}/{property}],methods=[GET],produces=[application/x-spring-data-compact+json || text/uri-list]}" onto public org.springframework.http.ResponseEntity<org.springframework.hateoas.ResourceSupport> org.springframework.data.rest.webmvc.RepositoryPropertyReferenceController.followPropertyReferenceCompact(org.springframework.data.rest.webmvc.RootResourceInformation,java.io.Serializable,java.lang.String,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler) throws java.lang.Exception
2017-06-17 06:03:15.567  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}/{property}],methods=[PATCH || PUT],consumes=[application/json || application/x-spring-data-compact+json || text/uri-list]}" onto public org.springframework.http.ResponseEntity<? extends org.springframework.hateoas.ResourceSupport> org.springframework.data.rest.webmvc.RepositoryPropertyReferenceController.createPropertyReference(org.springframework.data.rest.webmvc.RootResourceInformation,org.springframework.http.HttpMethod,org.springframework.hateoas.Resources<java.lang.Object>,java.io.Serializable,java.lang.String) throws java.lang.Exception
2017-06-17 06:03:15.567  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/{id}/{property}/{propertyId}],methods=[DELETE]}" onto public org.springframework.http.ResponseEntity<org.springframework.hateoas.ResourceSupport> org.springframework.data.rest.webmvc.RepositoryPropertyReferenceController.deletePropertyReferenceId(org.springframework.data.rest.webmvc.RootResourceInformation,java.io.Serializable,java.lang.String,java.lang.String) throws java.lang.Exception
2017-06-17 06:03:15.568  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/schema],methods=[GET],produces=[application/schema+json]}" onto public org.springframework.http.HttpEntity<org.springframework.data.rest.webmvc.json.JsonSchema> org.springframework.data.rest.webmvc.RepositorySchemaController.schema(org.springframework.data.rest.webmvc.RootResourceInformation)
2017-06-17 06:03:15.569  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/search],methods=[OPTIONS]}" onto public org.springframework.http.HttpEntity<?> org.springframework.data.rest.webmvc.RepositorySearchController.optionsForSearches(org.springframework.data.rest.webmvc.RootResourceInformation)
2017-06-17 06:03:15.569  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/search],methods=[GET]}" onto public org.springframework.hateoas.ResourceSupport org.springframework.data.rest.webmvc.RepositorySearchController.listSearches(org.springframework.data.rest.webmvc.RootResourceInformation)
2017-06-17 06:03:15.570  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/search/{search}],methods=[GET]}" onto public org.springframework.http.ResponseEntity<java.lang.Object> org.springframework.data.rest.webmvc.RepositorySearchController.executeSearch(org.springframework.data.rest.webmvc.RootResourceInformation,org.springframework.web.context.request.WebRequest,java.lang.String,org.springframework.data.rest.webmvc.support.DefaultedPageable,org.springframework.data.domain.Sort,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler)
2017-06-17 06:03:15.571  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/search/{search}],methods=[GET],produces=[application/x-spring-data-compact+json]}" onto public org.springframework.hateoas.ResourceSupport org.springframework.data.rest.webmvc.RepositorySearchController.executeSearchCompact(org.springframework.data.rest.webmvc.RootResourceInformation,org.springframework.web.context.request.WebRequest,java.lang.String,java.lang.String,org.springframework.data.rest.webmvc.support.DefaultedPageable,org.springframework.data.domain.Sort,org.springframework.data.rest.webmvc.PersistentEntityResourceAssembler)
2017-06-17 06:03:15.571  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/search/{search}],methods=[OPTIONS]}" onto public org.springframework.http.ResponseEntity<java.lang.Object> org.springframework.data.rest.webmvc.RepositorySearchController.optionsForSearch(org.springframework.data.rest.webmvc.RootResourceInformation,java.lang.String)
2017-06-17 06:03:15.572  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/search/{search}],methods=[HEAD]}" onto public org.springframework.http.ResponseEntity<java.lang.Object> org.springframework.data.rest.webmvc.RepositorySearchController.headForSearch(org.springframework.data.rest.webmvc.RootResourceInformation,java.lang.String)
2017-06-17 06:03:15.573  INFO 16589 --- [           main] o.s.d.r.w.RepositoryRestHandlerMapping   : Mapped "{[/{repository}/search],methods=[HEAD]}" onto public org.springframework.http.HttpEntity<?> org.springframework.data.rest.webmvc.RepositorySearchController.headForSearches(org.springframework.data.rest.webmvc.RootResourceInformation)
2017-06-17 06:03:15.579  INFO 16589 --- [           main] o.s.d.r.w.BaseUriAwareHandlerMapping     : Mapped "{[/alps || /alps/{repository}],methods=[OPTIONS]}" onto org.springframework.http.HttpEntity<?> org.springframework.data.rest.webmvc.alps.AlpsController.alpsOptions()
2017-06-17 06:03:15.580  INFO 16589 --- [           main] o.s.d.r.w.BaseUriAwareHandlerMapping     : Mapped "{[/alps],methods=[GET]}" onto org.springframework.http.HttpEntity<org.springframework.hateoas.alps.Alps> org.springframework.data.rest.webmvc.alps.AlpsController.alps()
2017-06-17 06:03:15.583  INFO 16589 --- [           main] o.s.d.r.w.BaseUriAwareHandlerMapping     : Mapped "{[/alps/{repository}],methods=[GET]}" onto org.springframework.http.HttpEntity<org.springframework.data.rest.webmvc.RootResourceInformation> org.springframework.data.rest.webmvc.alps.AlpsController.descriptor(org.springframework.data.rest.webmvc.RootResourceInformation)
2017-06-17 06:03:15.650  INFO 16589 --- [           main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on startup
2017-06-17 06:03:15.682  INFO 16589 --- [           main] o.s.c.support.DefaultLifecycleProcessor  : Starting beans in phase 0
2017-06-17 06:03:16.011  INFO 16589 --- [           main] s.b.c.e.t.TomcatEmbeddedServletContainer : Tomcat started on port(s): 9091 (http)
2017-06-17 06:03:16.020  INFO 16589 --- [           main] c.o.e.m.r.InMemoryRepositoryApplication  : Started InMemoryRepositoryApplication in 11.389 seconds (JVM running for 13.751)
```

Alternatively, _ctrl-z_ bring process back into _backgroud_ and unblock console
```
^Z
[1]+  已停止               java -jar repositories-mem/target/repositories-mem.jar --server.port=9091
[vagrant@localhost sample-microservices-springboot]$ bg %1
[1]+ java -jar repositories-mem/target/repositories-mem.jar --server.port=9091 &
```

Consle test
```
[vagrant@localhost sample-microservices-springboot]$ curl http://localhost:9091
2017-06-17 06:06:27.860  INFO 16589 --- [nio-9091-exec-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring FrameworkServlet 'dispatcherServlet'
2017-06-17 06:06:27.860  INFO 16589 --- [nio-9091-exec-1] o.s.web.servlet.DispatcherServlet        : FrameworkServlet 'dispatcherServlet': initialization started
2017-06-17 06:06:27.915  INFO 16589 --- [nio-9091-exec-1] o.s.web.servlet.DispatcherServlet        : FrameworkServlet 'dispatcherServlet': initialization completed in 55 ms
InMemoryRepositoryRestMessageController.getAll
InMemoryMessageRepository.findAllIt
[{"id":1,"text":"Hello","summary":"World","created":1497679393645},{"id":2,"text":"Hi","summary":"Universe","created":1497679393645},{"id":3,"text":"Hola","summary":"OpenShift","created":1497679393646}]
```

Then bring up to _foreground_ to terminate
```
[vagrant@localhost sample-microservices-springboot]$ fg %1
java -jar repositories-mem/target/repositories-mem.jar --server.port=9091
^C2017-06-17 06:07:18.667  INFO 16589 --- [       Thread-2] ationConfigEmbeddedWebApplicationContext : Closing org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@278b5329: startup date [Sat Jun 17 06:03:05 UTC 2017]; root of context hierarchy
2017-06-17 06:07:18.673  INFO 16589 --- [       Thread-2] o.s.c.support.DefaultLifecycleProcessor  : Stopping beans in phase 0
2017-06-17 06:07:18.682  INFO 16589 --- [       Thread-2] o.s.j.e.a.AnnotationMBeanExporter        : Unregistering JMX-exposed beans on shutdown
```

## Docker image

### web prject

Build image
```
[vagrant@localhost sample-microservices-springboot]$ docker build -t docker.io/tangfeixiong/springboot-osev3-examples:web -f web/Dockerfile.oracle-server-jre1.8 web/
Sending build context to Docker daemon 17.08 MB
Step 1 : FROM docker.io/centos:centos7
 ---> 34c4689f9727
Step 2 : LABEL maintainer "tangfeixiong" mailto "tangfx128@gmail.com"
 ---> Running in 3fc89ed0596d
 ---> 47b64e39e38f
Removing intermediate container 3fc89ed0596d
Step 3 : ARG java_pkg="http://172.17.4.50:48080/99-mirror/http0x3A0x2F0x2Fdownload.oracle.com/otn-pub0x2Fjava0x2Fjdk0x2F8u112-b15/server-jre-8u112-linux-x64.tar.gz"
 ---> Running in 6dacb6afe623
 ---> ef25466fdc27
Removing intermediate container 6dacb6afe623
Step 4 : COPY /target/web.jar /opt/app/web.jar
 ---> 6f9d463aad21
Removing intermediate container 62d2068b2741
Step 5 : ENV JAVA_HOME "/opt/java" PATH $PATH:/opt/java/bin JAVA_OPTIONS "-Xms128m -Xmx512m -XX:PermSize=128m -XX:MaxPermSize=256m" REMOTE_HOST "localhost" REMOTE_PORT 9091 SERVER_PORT 8091
 ---> Running in 25a3e8e7c3b4
 ---> 9c1011a73fd2
Removing intermediate container 25a3e8e7c3b4
Step 6 : RUN set -x     temp_dir=/tmp/build     && mkdir -p $JAVA_HOME     && javaPkgDest="${java_pkg:-'http://download.oracle.com/otn-pub/java/jdk/8u112-b15/server-jre-8u112-linux-x64.tar.gz'}"     && if [[ -f $temp_dir/$javaPkgDest ]]; then         tar -C $JAVA_HOME --strip-components=1 -zxf $temp_dir/$javaPkgDest;         rm -f $temp_dir/$javaPkgDest;     elif [[ ! "$javaPkgDest" =~ https?://download\.oracle\.com/?.* ]]; then         curl -jkSL ${javaPkgDest//%/%25}             | gunzip -c -             | tar -x -C $JAVA_HOME --strip-components=1;     else         curl -jkSLH "Cookie: oraclelicense=accept-securebackup-cookie; " "$javaPkgDest"             | gunzip             | tar -x -C "$JAVA_HOME" --strip-components=1;     fi     && rm -rf $temp_dir     && export JAVA_OPTIONS="-Xms128m -Xmx256m -XX:PermSize=128m -XX:MaxPermSize=128m"
 ---> Running in f3f69f6534c7
+ mkdir -p /opt/java
+ javaPkgDest=http://172.17.4.50:48080/99-mirror/http0x3A0x2F0x2Fdownload.oracle.com/otn-pub0x2Fjava0x2Fjdk0x2F8u112-b15/server-jre-8u112-linux-x64.tar.gz
+ [[ -f /http://172.17.4.50:48080/99-mirror/http0x3A0x2F0x2Fdownload.oracle.com/otn-pub0x2Fjava0x2Fjdk0x2F8u112-b15/server-jre-8u112-linux-x64.tar.gz ]]
+ [[ ! http://172.17.4.50:48080/99-mirror/http0x3A0x2F0x2Fdownload.oracle.com/otn-pub0x2Fjava0x2Fjdk0x2F8u112-b15/server-jre-8u112-linux-x64.tar.gz =~ https?://download\.oracle\.com/?.* ]]
+ curl -jkSL http://172.17.4.50:48080/99-mirror/http0x3A0x2F0x2Fdownload.oracle.com/otn-pub0x2Fjava0x2Fjdk0x2F8u112-b15/server-jre-8u112-linux-x64.tar.gz
+ gunzip -c -
+ tar -x -C /opt/java --strip-components=1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 57.1M  100 57.1M    0     0  30.6M      0  0:00:01  0:00:01 --:--:-- 30.6M
+ rm -rf
+ export 'JAVA_OPTIONS=-Xms128m -Xmx256m -XX:PermSize=128m -XX:MaxPermSize=128m'
+ JAVA_OPTIONS='-Xms128m -Xmx256m -XX:PermSize=128m -XX:MaxPermSize=128m'
 ---> d81bdf6d3514
Removing intermediate container f3f69f6534c7
Step 7 : EXPOSE 8091
 ---> Running in b6ca98648cfc
 ---> 74fbe5d8f509
Removing intermediate container b6ca98648cfc
Step 8 : CMD java -Djava.security.egd=file:/dev/./urandom $JAVA_OPTIONS -jar /opt/app/web.jar --remoteHost=$REMOTE_HOST --remotePort=$REMOTE_PORT --server.port=$SERVER_PORT
 ---> Running in 8f5750f09e94
 ---> 6034ad36219c
Removing intermediate container 8f5750f09e94
Successfully built 6034ad36219c
```
Check
```
[vagrant@localhost sample-microservices-springboot]$ docker images tangfeixiong/springboot-osev3-examples
REPOSITORY                                         TAG                 IMAGE ID            CREATED              SIZE
docker.io/tangfeixiong/springboot-osev3-examples   web                 6034ad36219c        About a minute ago   369.7 MB
```

Test
```
[vagrant@localhost sample-microservices-springboot]$ docker run -d -p 8091:8091 --name osev3-examples-web tangfeixiong/springboot-osev3-examples:web
d64dda193d6e680a3e4069b6104edd565d918bbb314af54cbd3b597328ae1d0a
[vagrant@localhost sample-microservices-springboot]$ curl 127.0.0.1:8091
<!DOCTYPE html>

<html>
  <head>
    <title>Messages : View all</title>
    
    <link rel="stylesheet" href="/css/bootstrap.min.css" />
  
    
  </head>
  <body>
    <div class="container">
      <div class="navbar">
        <div class="navbar-inner">
          <a class="brand" href="https://github.com/ultraq/thymeleaf-layout-dialect">
              Thymeleaf - Layout
          </a>
          <ul class="nav">
            <li>
              <a href="/">
                Messages
              </a>
            </li>
          </ul>
        </div>
      </div>
      <h1>Messages : View all</h1>
      <div class="container">
      <div class="pull-right">
        <a href="/?form">Create Message</a>
       </div>
      <table class="table table-bordered table-striped">
        <thead>
          <tr>
            <td>ID</td>
            <td>Created</td>
            <td>Summary</td>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td colspan="3">
              No messages
            </td>
          </tr>
          
        </tbody>
      </table>
    </div>
    </div>
  </body>
</html>
[vagrant@localhost sample-microservices-springboot]$ docker stop osev3-examples-web
osev3-examples-web
[vagrant@localhost sample-microservices-springboot]$ docker rm osev3-examples-web
osev3-examples-web
```

Push
```
[vagrant@localhost sample-microservices-springboot]$ docker tag tangfeixiong/springboot-osev3-examples:web 172.17.4.50\:5000/springboot-osev3-examples:web
[vagrant@localhost sample-microservices-springboot]$ docker push 172.17.4.50\:5000/springboot-osev3-examples:web
The push refers to a repository [172.17.4.50:5000/springboot-osev3-examples]
7b64247b0f5c: Pushed 
5dde89c1fa66: Pushed 
5f70bf18a086: Pushed 
97ca462ad9ee: Pushed 
web: digest: sha256:5263b25c25a3e6a53b00d6da78239cc12ba0c282749011e03084e4182ba70bf1 size: 1549
```

### repository-mem project

Build image (OpenJDK)
```
[vagrant@localhost sample-microservices-springboot]$ docker build -t docker.io/tangfeixiong/springboot-osev3-examples:repo -f repositories-mem/Dockerfile.java-openjdk1.8-headless repositories-mem/
Sending build context to Docker daemon 17.82 MB
Step 1 : FROM docker.io/centos:centos7
 ---> 34c4689f9727
Step 2 : LABEL maintainer "tangfeixiong" mailto "tangfx128@gmail.com" project "osev3-examples" name "sample-microservices-springboot"
 ---> Running in 966a6925a2bf
 ---> 9a497590a830
Removing intermediate container 966a6925a2bf
Step 3 : ARG localrepo_url=http://172.17.4.50:48080/99-mirror/http0x3A0x2F0x2Fmirror.centos.org0x2Fcentos/centos7-local.repo
 ---> Running in 12ba1c6db28e
 ---> f8a4642828c0
Removing intermediate container 12ba1c6db28e
Step 4 : ENV JAVA_HOME "/opt/java" JAVA_OPTIONS "-Xms128m -Xmx512m -XX:PermSize=128m -XX:MaxPermSize=256m" SERVER_PORT 9091
 ---> Running in 4fe3943b9615
 ---> 6664ce561f3f
Removing intermediate container 4fe3943b9615
Step 5 : COPY /target/repositories-mem.jar /
 ---> 90cebbb5d4fb
Removing intermediate container e050967d5ee8
Step 6 : RUN set -x     && curl -jkSL $localrepo_url -o /etc/yum.repos.d/local.repo     && install_pkgs="         tar         unzip         bc         which         lsof         java-1.8.0-openjdk-headless     "     && yum install -y --disablerepo='\*' --enablerepo='base,localrepo-\*' $install_pkgs     && yum clean all -y     && export JAVA_OPTIONS="-Xms128m -Xmx256m -XX:PermSize=128m -XX:MaxPermSize=128m"
 ---> Running in 584ea477cc20
+ curl -jkSL http://172.17.4.50:48080/99-mirror/http0x3A0x2F0x2Fmirror.centos.org0x2Fcentos/centos7-local.repo -o /etc/yum.repos.d/local.repo
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1163  100  1163    0     0  76791      0 --:--:-- --:--:-- --:--:-- 83071
+ install_pkgs='         tar         unzip         bc         which         lsof         java-1.8.0-openjdk-headless     '
+ yum install -y '--disablerepo=\*' '--enablerepo=base,localrepo-\*' tar unzip bc which lsof java-1.8.0-openjdk-headless
Loaded plugins: fastestmirror, ovl
Determining fastest mirrors
 * base: mirrors.neusoft.edu.cn
 * extras: mirrors.btte.net
 * updates: mirror.bit.edu.cn
Resolving Dependencies
--> Running transaction check
---> Package bc.x86_64 0:1.06.95-13.el7 will be installed
---> Package java-1.8.0-openjdk-headless.x86_64 1:1.8.0.131-3.b12.el7_3 will be installed
--> Processing Dependency: tzdata-java >= 2015d for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
--> Processing Dependency: nss(x86-64) >= 3.28.4 for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
--> Processing Dependency: copy-jdk-configs >= 1.1-3 for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
--> Processing Dependency: chkconfig >= 1.7 for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
--> Processing Dependency: chkconfig >= 1.7 for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
--> Processing Dependency: lksctp-tools(x86-64) for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
--> Processing Dependency: libjpeg.so.62(LIBJPEG_6.2)(64bit) for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
--> Processing Dependency: jpackage-utils for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
--> Processing Dependency: libjpeg.so.62()(64bit) for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
--> Processing Dependency: libfreetype.so.6()(64bit) for package: 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_64
---> Package lsof.x86_64 0:4.87-4.el7 will be installed
---> Package tar.x86_64 2:1.26-29.el7 will be updated
---> Package tar.x86_64 2:1.26-31.el7 will be an update
---> Package unzip.x86_64 0:6.0-16.el7 will be installed
---> Package which.x86_64 0:2.20-7.el7 will be installed
--> Running transaction check
---> Package chkconfig.x86_64 0:1.3.61-5.el7_2.1 will be updated
---> Package chkconfig.x86_64 0:1.7.2-1.el7 will be an update
---> Package copy-jdk-configs.noarch 0:1.2-1.el7 will be installed
---> Package freetype.x86_64 0:2.4.11-12.el7 will be installed
---> Package javapackages-tools.noarch 0:3.4.1-11.el7 will be installed
--> Processing Dependency: python-javapackages = 3.4.1-11.el7 for package: javapackages-tools-3.4.1-11.el7.noarch
--> Processing Dependency: libxslt for package: javapackages-tools-3.4.1-11.el7.noarch
---> Package libjpeg-turbo.x86_64 0:1.2.90-5.el7 will be installed
---> Package lksctp-tools.x86_64 0:1.0.17-2.el7 will be installed
---> Package nss.x86_64 0:3.21.0-9.el7_2 will be updated
--> Processing Dependency: nss = 3.21.0-9.el7_2 for package: nss-sysinit-3.21.0-9.el7_2.x86_64
--> Processing Dependency: nss(x86-64) = 3.21.0-9.el7_2 for package: nss-tools-3.21.0-9.el7_2.x86_64
---> Package nss.x86_64 0:3.28.4-1.2.el7_3 will be an update
--> Processing Dependency: nss-util >= 3.28.2-1.1 for package: nss-3.28.4-1.2.el7_3.x86_64
--> Processing Dependency: nspr >= 4.13.1 for package: nss-3.28.4-1.2.el7_3.x86_64
--> Processing Dependency: libnssutil3.so(NSSUTIL_3.24)(64bit) for package: nss-3.28.4-1.2.el7_3.x86_64
---> Package tzdata-java.noarch 0:2017b-1.el7 will be installed
--> Running transaction check
---> Package libxslt.x86_64 0:1.1.28-5.el7 will be installed
---> Package nspr.x86_64 0:4.11.0-1.el7_2 will be updated
---> Package nspr.x86_64 0:4.13.1-1.0.el7_3 will be an update
---> Package nss-sysinit.x86_64 0:3.21.0-9.el7_2 will be updated
---> Package nss-sysinit.x86_64 0:3.28.4-1.2.el7_3 will be an update
---> Package nss-tools.x86_64 0:3.21.0-9.el7_2 will be updated
---> Package nss-tools.x86_64 0:3.28.4-1.2.el7_3 will be an update
---> Package nss-util.x86_64 0:3.21.0-2.2.el7_2 will be updated
---> Package nss-util.x86_64 0:3.28.4-1.0.el7_3 will be an update
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
 java-1.8.0-openjdk-headless  x86_64  1:1.8.0.131-3.b12.el7_3    updates   31 M
 lsof                         x86_64  4.87-4.el7                 base     331 k
 unzip                        x86_64  6.0-16.el7                 base     169 k
 which                        x86_64  2.20-7.el7                 base      41 k
Updating:
 tar                          x86_64  2:1.26-31.el7              base     843 k
Installing for dependencies:
 copy-jdk-configs             noarch  1.2-1.el7                  base      14 k
 freetype                     x86_64  2.4.11-12.el7              base     391 k
 javapackages-tools           noarch  3.4.1-11.el7               base      73 k
 libjpeg-turbo                x86_64  1.2.90-5.el7               base     134 k
 libxslt                      x86_64  1.1.28-5.el7               base     242 k
 lksctp-tools                 x86_64  1.0.17-2.el7               base      88 k
 python-javapackages          noarch  3.4.1-11.el7               base      31 k
 python-lxml                  x86_64  3.2.1-4.el7                base     758 k
 tzdata-java                  noarch  2017b-1.el7                updates  183 k
Updating for dependencies:
 chkconfig                    x86_64  1.7.2-1.el7                base     175 k
 nspr                         x86_64  4.13.1-1.0.el7_3           updates  126 k
 nss                          x86_64  3.28.4-1.2.el7_3           updates  872 k
 nss-sysinit                  x86_64  3.28.4-1.2.el7_3           updates   58 k
 nss-tools                    x86_64  3.28.4-1.2.el7_3           updates  496 k
 nss-util                     x86_64  3.28.4-1.0.el7_3           updates   73 k

Transaction Summary
================================================================================
Install  5 Packages (+9 Dependent packages)
Upgrade  1 Package  (+6 Dependent packages)

Total download size: 36 M
Downloading packages:
Delta RPMs disabled because /usr/bin/applydeltarpm not installed.
warning: /var/cache/yum/x86_64/7/base/packages/copy-jdk-configs-1.2-1.el7.noarch.rpm: Header V3 RSA/SHA256 Signature, key ID f4a80eb5: NOKEY
Public key for copy-jdk-configs-1.2-1.el7.noarch.rpm is not installed
Public key for nss-tools-3.28.4-1.2.el7_3.x86_64.rpm is not installed
--------------------------------------------------------------------------------
Total                                              3.0 MB/s |  36 MB  00:12     
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Importing GPG key 0xF4A80EB5:
 Userid     : "CentOS-7 Key (CentOS 7 Official Signing Key) <security@centos.org>"
 Fingerprint: 6341 ab27 53d7 8a78 a7c2 7bb1 24c6 a8a7 f4a8 0eb5
 Package    : centos-release-7-2.1511.el7.centos.2.10.x86_64 (@CentOS)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Updating   : nspr-4.13.1-1.0.el7_3.x86_64                                1/28 
  Updating   : nss-util-3.28.4-1.0.el7_3.x86_64                            2/28 
  Updating   : chkconfig-1.7.2-1.el7.x86_64                                3/28 
  Updating   : nss-sysinit-3.28.4-1.2.el7_3.x86_64                         4/28 
  Updating   : nss-3.28.4-1.2.el7_3.x86_64                                 5/28 
  Installing : libxslt-1.1.28-5.el7.x86_64                                 6/28 
  Installing : python-lxml-3.2.1-4.el7.x86_64                              7/28 
  Installing : python-javapackages-3.4.1-11.el7.noarch                     8/28 
  Installing : javapackages-tools-3.4.1-11.el7.noarch                      9/28 
  Installing : freetype-2.4.11-12.el7.x86_64                              10/28 
  Installing : tzdata-java-2017b-1.el7.noarch                             11/28 
  Installing : copy-jdk-configs-1.2-1.el7.noarch                          12/28 
  Installing : lksctp-tools-1.0.17-2.el7.x86_64                           13/28 
  Installing : libjpeg-turbo-1.2.90-5.el7.x86_64                          14/28 
  Installing : 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_   15/28 
  Updating   : nss-tools-3.28.4-1.2.el7_3.x86_64                          16/28 
  Installing : lsof-4.87-4.el7.x86_64                                     17/28 
  Updating   : 2:tar-1.26-31.el7.x86_64                                   18/28 
  Installing : bc-1.06.95-13.el7.x86_64                                   19/28 
  Installing : which-2.20-7.el7.x86_64                                    20/28 
install-info: No such file or directory for /usr/share/info/which.info.gz
  Installing : unzip-6.0-16.el7.x86_64                                    21/28 
  Cleanup    : nss-tools-3.21.0-9.el7_2.x86_64                            22/28 
  Cleanup    : nss-sysinit-3.21.0-9.el7_2.x86_64                          23/28 
  Cleanup    : nss-3.21.0-9.el7_2.x86_64                                  24/28 
  Cleanup    : nss-util-3.21.0-2.2.el7_2.x86_64                           25/28 
  Cleanup    : nspr-4.11.0-1.el7_2.x86_64                                 26/28 
  Cleanup    : chkconfig-1.3.61-5.el7_2.1.x86_64                          27/28 
  Cleanup    : 2:tar-1.26-29.el7.x86_64                                   28/28 
  Verifying  : python-javapackages-3.4.1-11.el7.noarch                     1/28 
  Verifying  : nss-tools-3.28.4-1.2.el7_3.x86_64                           2/28 
  Verifying  : libjpeg-turbo-1.2.90-5.el7.x86_64                           3/28 
  Verifying  : nss-3.28.4-1.2.el7_3.x86_64                                 4/28 
  Verifying  : python-lxml-3.2.1-4.el7.x86_64                              5/28 
  Verifying  : unzip-6.0-16.el7.x86_64                                     6/28 
  Verifying  : nss-util-3.28.4-1.0.el7_3.x86_64                            7/28 
  Verifying  : 1:java-1.8.0-openjdk-headless-1.8.0.131-3.b12.el7_3.x86_    8/28 
  Verifying  : which-2.20-7.el7.x86_64                                     9/28 
  Verifying  : bc-1.06.95-13.el7.x86_64                                   10/28 
  Verifying  : 2:tar-1.26-31.el7.x86_64                                   11/28 
  Verifying  : lksctp-tools-1.0.17-2.el7.x86_64                           12/28 
  Verifying  : copy-jdk-configs-1.2-1.el7.noarch                          13/28 
  Verifying  : nss-sysinit-3.28.4-1.2.el7_3.x86_64                        14/28 
  Verifying  : libxslt-1.1.28-5.el7.x86_64                                15/28 
  Verifying  : tzdata-java-2017b-1.el7.noarch                             16/28 
  Verifying  : chkconfig-1.7.2-1.el7.x86_64                               17/28 
  Verifying  : javapackages-tools-3.4.1-11.el7.noarch                     18/28 
  Verifying  : freetype-2.4.11-12.el7.x86_64                              19/28 
  Verifying  : lsof-4.87-4.el7.x86_64                                     20/28 
  Verifying  : nspr-4.13.1-1.0.el7_3.x86_64                               21/28 
  Verifying  : 2:tar-1.26-29.el7.x86_64                                   22/28 
  Verifying  : nspr-4.11.0-1.el7_2.x86_64                                 23/28 
  Verifying  : nss-3.21.0-9.el7_2.x86_64                                  24/28 
  Verifying  : nss-tools-3.21.0-9.el7_2.x86_64                            25/28 
  Verifying  : nss-sysinit-3.21.0-9.el7_2.x86_64                          26/28 
  Verifying  : chkconfig-1.3.61-5.el7_2.1.x86_64                          27/28 
  Verifying  : nss-util-3.21.0-2.2.el7_2.x86_64                           28/28 

Installed:
  bc.x86_64 0:1.06.95-13.el7                                                    
  java-1.8.0-openjdk-headless.x86_64 1:1.8.0.131-3.b12.el7_3                    
  lsof.x86_64 0:4.87-4.el7                                                      
  unzip.x86_64 0:6.0-16.el7                                                     
  which.x86_64 0:2.20-7.el7                                                     

Dependency Installed:
  copy-jdk-configs.noarch 0:1.2-1.el7       freetype.x86_64 0:2.4.11-12.el7    
  javapackages-tools.noarch 0:3.4.1-11.el7  libjpeg-turbo.x86_64 0:1.2.90-5.el7
  libxslt.x86_64 0:1.1.28-5.el7             lksctp-tools.x86_64 0:1.0.17-2.el7 
  python-javapackages.noarch 0:3.4.1-11.el7 python-lxml.x86_64 0:3.2.1-4.el7   
  tzdata-java.noarch 0:2017b-1.el7         

Updated:
  tar.x86_64 2:1.26-31.el7                                                      

Dependency Updated:
  chkconfig.x86_64 0:1.7.2-1.el7        nspr.x86_64 0:4.13.1-1.0.el7_3         
  nss.x86_64 0:3.28.4-1.2.el7_3         nss-sysinit.x86_64 0:3.28.4-1.2.el7_3  
  nss-tools.x86_64 0:3.28.4-1.2.el7_3   nss-util.x86_64 0:3.28.4-1.0.el7_3     

Complete!
+ yum clean all -y
Loaded plugins: fastestmirror, ovl
Cleaning repos: base extras updates
Cleaning up everything
Cleaning up list of fastest mirrors
+ export 'JAVA_OPTIONS=-Xms128m -Xmx256m -XX:PermSize=128m -XX:MaxPermSize=128m'
+ JAVA_OPTIONS='-Xms128m -Xmx256m -XX:PermSize=128m -XX:MaxPermSize=128m'
 ---> 16629d8f8b60
Removing intermediate container 584ea477cc20
Step 7 : EXPOSE 8091
 ---> Running in 329064f76337
 ---> 5a7e60452b97
Removing intermediate container 329064f76337
Step 8 : CMD java -Djava.security.egd=file:/dev/./urandom $JAVA_OPTIONS -jar /repositories-mem.jar --server.port=$SERVER_PORT
 ---> Running in c2105e7f94f6
 ---> 6d5a44d12c65
Removing intermediate container c2105e7f94f6
Successfully built 6d5a44d12c65
```

Run
```
[vagrant@localhost sample-microservices-springboot]$ docker run -d -p 9091:9091 --name osev3-examples-repo docker.io/tangfeixiong/springboot-osev3-examples:repo
2f7a782e8d5e2477e02770eaec49925b8b6dbe940175fb1ebb5876192a8ab62f
[vagrant@localhost sample-microservices-springboot]$ curl  -H "Content-type: application/json" -X POST -d '{"id":10,"text":"foo","summary":"bar"}'  http://localhost:9091/
{"id":10,"text":"foo","summary":"bar","created":1497678844930}[vagrant@localhost sample-microservices-springboot]$ 
[vagrant@localhost sample-microservices-springboot]$ curl  -H "Content-type: application/json" -X PUT -d '{"id":10,"text":"foo","summary":"baz"}'  http://localhost:9091/10
{"id":10,"text":"foo","summary":"baz","created":1497678844930}[vagrant@localhost sample-microservices-springboot]$ curl http://localhost:9091
[{"id":1,"text":"Hello","summary":"World","created":1497678837211},{"id":2,"text":"Hi","summary":"Universe","created":1497678837211},{"id":3,"text":"Hola","summary":"OpenShift","created":1497678837211},{"id":10,"text":"foo","summary":"baz","created":1497678844930}][vagrant@localhost sample-microservices-springboot]$ 
[vagrant@localhost sample-microservices-springboot]$ curl http://localhost:9091/1
{"id":1,"text":"Hello","summary":"World","created":1497678837211}[vagrant@localhost sample-microservices-springboot]$ 
[vagrant@localhost sample-microservices-springboot]$ curl -X DELETE http://localhost:9091/10
[vagrant@localhost sample-microservices-springboot]$ curl http://localhost:9091
[{"id":1,"text":"Hello","summary":"World","created":1497678837211},{"id":2,"text":"Hi","summary":"Universe","created":1497678837211},{"id":3,"text":"Hola","summary":"OpenShift","created":1497678837211}][vagrant@localhost sample-microservices-springboot]$ 
```

Drop and clear
```
[vagrant@localhost sample-microservices-springboot]$ docker stop osev3-examples-repo
osev3-examples-repo
[vagrant@localhost sample-microservices-springboot]$ docker rm osev3-examples-repo
osev3-examples-repo
```