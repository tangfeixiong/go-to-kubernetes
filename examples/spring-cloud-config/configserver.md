### Spring Cloud Centrized Config

Maven build Docker image
```
[vagrant@localhost configserver]$ mvn install spring-boot:repackage docker:build -Dmaven.test.skip=true
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Config Server 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- git-commit-id-plugin:2.1.11:revision (default) @ configserver ---
[info] dotGitDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-config/configserver/.git
[info] git.build.user.name "Spencer Gibb"
[info] git.build.user.email "spencer@gibb.us"
[info] git.branch master
[info] git.commit.id.describe 
[info] git.commit.id 911f79dccaabecbec7594bbecc8072abddb692b3
[info] git.commit.id.abbrev 911f79d
[info] git.commit.user.name "GitHub"
[info] git.commit.user.email "noreply@github.com"
[info] git.commit.message.full "Merge pull request #19 from markfisher/rabbit-instead-of-redisrabbitmq instead of redis in docker-compose.yml"
[info] git.commit.message.short "Merge pull request #19 from markfisher/rabbit-instead-of-redis"
[info] git.commit.time "2016-11-07 14:29:33 -0500"
[info] git.remote.origin.url https://github.com/spring-cloud-samples/configserver
[info] git.tags grafted,,->,origin/master,origin/
[info] git.build.time 2017-07-21T16:40:16+0000
[info] git.commit.id.describe-short 
[info] found property git.tags
[info] found property git.commit.id.abbrev
[info] found property git.commit.user.email
[info] found property git.commit.message.full
[info] found property git.commit.id
[info] found property git.commit.id.describe-short
[info] found property git.commit.message.short
[info] found property git.commit.user.name
[info] found property git.build.user.name
[info] found property git.commit.id.describe
[info] found property git.build.user.email
[info] found property git.branch
[info] found property git.commit.time
[info] found property git.build.time
[info] found property git.remote.origin.url
[info] Writing properties file to [ /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-config/configserver/target/classes/git.properties ] (for module  Config Server )...
[info] Config Server ] project Config Server
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ configserver ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 1 resource
[INFO] Copying 3 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ configserver ---
[INFO] Nothing to compile - all classes are up to date
[INFO] 
[INFO] --- maven-resources-plugin:2.6:testResources (default-testResources) @ configserver ---
[INFO] Not copying test resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:testCompile (default-testCompile) @ configserver ---
[INFO] Not compiling test sources
[INFO] 
[INFO] --- maven-surefire-plugin:2.18.1:test (default-test) @ configserver ---
[INFO] Tests are skipped.
[INFO] 
[INFO] --- maven-jar-plugin:2.6:jar (default-jar) @ configserver ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-config/configserver/target/configserver-0.0.1-SNAPSHOT.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.4.1.BUILD-SNAPSHOT:repackage (default) @ configserver ---
[INFO] 
[INFO] --- maven-install-plugin:2.5.2:install (default-install) @ configserver ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-config/configserver/target/configserver-0.0.1-SNAPSHOT.jar to /home/vagrant/.m2/repository/org/demo/configserver/0.0.1-SNAPSHOT/configserver-0.0.1-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-config/configserver/pom.xml to /home/vagrant/.m2/repository/org/demo/configserver/0.0.1-SNAPSHOT/configserver-0.0.1-SNAPSHOT.pom
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.4.1.BUILD-SNAPSHOT:repackage (default-cli) @ configserver ---
[INFO] 
[INFO] --- docker-maven-plugin:0.2.3:build (default-cli) @ configserver ---
[INFO] Copying /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-config/configserver/target/configserver-0.0.1-SNAPSHOT.jar -> /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-config/configserver/target/docker/configserver-0.0.1-SNAPSHOT.jar
[INFO] Copying src/main/docker/Dockerfile -> /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-config/configserver/target/docker/Dockerfile
[INFO] Building image springcloud/configserver
Step 1 : FROM openjdk:8-jre
---> d18ae8f18728
Step 2 : ADD configserver-0.0.1-SNAPSHOT.jar app.jar
---> 1d1215c4ab31
Removing intermediate container 553973162ca2
Step 3 : VOLUME /tmp
---> Running in d756c67e4472
---> 98ef723d8536
Removing intermediate container d756c67e4472
Step 4 : VOLUME /target
---> Running in d25f0187eb19
---> d718edf6a7e1
Removing intermediate container d25f0187eb19
Step 5 : RUN bash -c 'touch /app.jar'
---> Running in 130ebc7737d0
---> 8580dfb2bd33
Removing intermediate container 130ebc7737d0
Step 6 : EXPOSE 8888
---> Running in af983a2aebaa
---> edbd693b96d9
Removing intermediate container af983a2aebaa
Step 7 : ENTRYPOINT java -Djava.security.egd=file:/dev/./urandom -jar /app.jar
---> Running in a2cae83e1c6f
---> 6c4a5ba30a1d
Removing intermediate container a2cae83e1c6f
Successfully built 6c4a5ba30a1d
[INFO] Built springcloud/configserver
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 44.484 s
[INFO] Finished at: 2017-07-21T16:40:52+00:00
[INFO] Final Memory: 31M/248M
[INFO] ------------------------------------------------------------------------
```

Build SpringBoot jar
```
[vagrant@localhost configserver]$ mvn install spring-boot:repackage
[INFO] Scanning for projects...
...snip...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Config Server 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
...snip...
[INFO] 
[INFO] --- git-commit-id-plugin:2.1.11:revision (default) @ configserver ---
[info] dotGitDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/.git
[info] git.build.user.name "Spencer Gibb"
[info] git.build.user.email "spencer@gibb.us"
[info] git.branch master
[info] git.commit.id.describe 
[info] git.commit.id 911f79dccaabecbec7594bbecc8072abddb692b3
[info] git.commit.id.abbrev 911f79d
[info] git.commit.user.name "GitHub"
[info] git.commit.user.email "noreply@github.com"
[info] git.commit.message.full "Merge pull request #19 from markfisher/rabbit-instead-of-redisrabbitmq instead of redis in docker-compose.yml"
[info] git.commit.message.short "Merge pull request #19 from markfisher/rabbit-instead-of-redis"
[info] git.commit.time "2016-11-07 14:29:33 -0500"
[info] git.remote.origin.url https://github.com/spring-cloud-samples/configserver
[info] git.tags grafted,,->,origin/master,origin/
[info] git.build.time 2017-07-21T08:27:43+0000
[info] git.commit.id.describe-short 
[info] found property git.tags
[info] found property git.commit.id.abbrev
[info] found property git.commit.user.email
[info] found property git.commit.message.full
[info] found property git.commit.id
[info] found property git.commit.id.describe-short
[info] found property git.commit.message.short
[info] found property git.commit.user.name
[info] found property git.build.user.name
[info] found property git.commit.id.describe
[info] found property git.build.user.email
[info] found property git.branch
[info] found property git.commit.time
[info] found property git.build.time
[info] found property git.remote.origin.url
[info] Writing properties file to [ /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/target/classes/git.properties ] (for module  Config Server )...
[info] Config Server ] project Config Server
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ configserver ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] Copying 1 resource
[INFO] Copying 3 resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:compile (default-compile) @ configserver ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 1 source file to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/target/classes
[INFO] 
[INFO] --- maven-resources-plugin:2.6:testResources (default-testResources) @ configserver ---
[INFO] Using 'UTF-8' encoding to copy filtered resources.
[INFO] skip non existing resourceDirectory /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/src/test/resources
[INFO] 
[INFO] --- maven-compiler-plugin:3.1:testCompile (default-testCompile) @ configserver ---
[INFO] Changes detected - recompiling the module!
[INFO] Compiling 1 source file to /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/target/test-classes
[WARNING] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/src/test/java/demo/ApplicationTests.java: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/src/test/java/demo/ApplicationTests.java使用或覆盖了已过时的 API。
[WARNING] /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/src/test/java/demo/ApplicationTests.java: 有关详细信息, 请使用 -Xlint:deprecation 重新编译。
[INFO] 
[INFO] --- maven-surefire-plugin:2.18.1:test (default-test) @ configserver ---
[INFO] Surefire report directory: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/target/surefire-reports

-------------------------------------------------------
 T E S T S
-------------------------------------------------------
08:27:53.030 [main] DEBUG org.springframework.test.context.junit4.SpringJUnit4ClassRunner - SpringJUnit4ClassRunner constructor called with [class demo.ApplicationTests]
08:27:53.072 [main] DEBUG org.springframework.test.context.BootstrapUtils - Instantiating CacheAwareContextLoaderDelegate from class [org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate]
08:27:53.109 [main] DEBUG org.springframework.test.context.BootstrapUtils - Instantiating BootstrapContext using constructor [public org.springframework.test.context.support.DefaultBootstrapContext(java.lang.Class,org.springframework.test.context.CacheAwareContextLoaderDelegate)]
08:27:53.160 [main] DEBUG org.springframework.test.context.BootstrapUtils - Instantiating TestContextBootstrapper for test class [demo.ApplicationTests] from class [org.springframework.test.context.web.WebTestContextBootstrapper]
08:27:53.231 [main] DEBUG org.springframework.test.context.web.WebTestContextBootstrapper - Found explicit ContextLoader class [org.springframework.boot.test.SpringApplicationContextLoader] for context configuration attributes [ContextConfigurationAttributes@57855c9a declaringClass = 'demo.ApplicationTests', classes = '{class demo.ConfigServerApplication}', locations = '{}', inheritLocations = true, initializers = '{}', inheritInitializers = true, name = [null], contextLoaderClass = 'org.springframework.boot.test.SpringApplicationContextLoader']
08:27:53.293 [main] DEBUG org.springframework.test.context.support.AbstractContextLoader - Did not detect default resource location for test class [demo.ApplicationTests]: class path resource [demo/ApplicationTests-context.xml] does not exist
08:27:53.299 [main] DEBUG org.springframework.test.context.support.AbstractContextLoader - Did not detect default resource location for test class [demo.ApplicationTests]: class path resource [demo/ApplicationTestsContext.groovy] does not exist
08:27:53.299 [main] INFO org.springframework.test.context.support.AbstractContextLoader - Could not detect default resource locations for test class [demo.ApplicationTests]: no resource found for suffixes {-context.xml, Context.groovy}.
08:27:53.526 [main] DEBUG org.springframework.test.context.support.ActiveProfilesUtils - Could not find an 'annotation declaring class' for annotation type [org.springframework.test.context.ActiveProfiles] and class [demo.ApplicationTests]
08:27:53.575 [main] INFO org.springframework.test.context.web.WebTestContextBootstrapper - Using TestExecutionListeners: [org.springframework.boot.test.IntegrationTestPropertiesListener@24a35978, org.springframework.test.context.support.DependencyInjectionTestExecutionListener@16f7c8c1, org.springframework.test.context.support.DirtiesContextTestExecutionListener@2f0a87b3, org.springframework.test.context.transaction.TransactionalTestExecutionListener@319b92f3, org.springframework.test.context.jdbc.SqlScriptsTestExecutionListener@fcd6521]
08:27:53.588 [main] DEBUG org.springframework.test.annotation.ProfileValueUtils - Retrieved @ProfileValueSourceConfiguration [null] for test class [demo.ApplicationTests]
08:27:53.591 [main] DEBUG org.springframework.test.annotation.ProfileValueUtils - Retrieved ProfileValueSource type [class org.springframework.test.annotation.SystemProfileValueSource] for class [demo.ApplicationTests]
Running demo.ApplicationTests
08:27:53.600 [main] DEBUG org.springframework.test.context.junit4.SpringJUnit4ClassRunner - SpringJUnit4ClassRunner constructor called with [class demo.ApplicationTests]
08:27:53.601 [main] DEBUG org.springframework.test.context.BootstrapUtils - Instantiating CacheAwareContextLoaderDelegate from class [org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate]
08:27:53.601 [main] DEBUG org.springframework.test.context.BootstrapUtils - Instantiating BootstrapContext using constructor [public org.springframework.test.context.support.DefaultBootstrapContext(java.lang.Class,org.springframework.test.context.CacheAwareContextLoaderDelegate)]
08:27:53.603 [main] DEBUG org.springframework.test.context.BootstrapUtils - Instantiating TestContextBootstrapper for test class [demo.ApplicationTests] from class [org.springframework.test.context.web.WebTestContextBootstrapper]
08:27:53.605 [main] DEBUG org.springframework.test.context.web.WebTestContextBootstrapper - Found explicit ContextLoader class [org.springframework.boot.test.SpringApplicationContextLoader] for context configuration attributes [ContextConfigurationAttributes@24b1d79b declaringClass = 'demo.ApplicationTests', classes = '{class demo.ConfigServerApplication}', locations = '{}', inheritLocations = true, initializers = '{}', inheritInitializers = true, name = [null], contextLoaderClass = 'org.springframework.boot.test.SpringApplicationContextLoader']
08:27:53.609 [main] DEBUG org.springframework.test.context.support.AbstractContextLoader - Did not detect default resource location for test class [demo.ApplicationTests]: class path resource [demo/ApplicationTests-context.xml] does not exist
08:27:53.612 [main] DEBUG org.springframework.test.context.support.AbstractContextLoader - Did not detect default resource location for test class [demo.ApplicationTests]: class path resource [demo/ApplicationTestsContext.groovy] does not exist
08:27:53.613 [main] INFO org.springframework.test.context.support.AbstractContextLoader - Could not detect default resource locations for test class [demo.ApplicationTests]: no resource found for suffixes {-context.xml, Context.groovy}.
08:27:53.642 [main] DEBUG org.springframework.test.context.support.ActiveProfilesUtils - Could not find an 'annotation declaring class' for annotation type [org.springframework.test.context.ActiveProfiles] and class [demo.ApplicationTests]
08:27:53.644 [main] INFO org.springframework.test.context.web.WebTestContextBootstrapper - Using TestExecutionListeners: [org.springframework.boot.test.IntegrationTestPropertiesListener@1cab0bfb, org.springframework.test.context.support.DependencyInjectionTestExecutionListener@5e955596, org.springframework.test.context.support.DirtiesContextTestExecutionListener@50de0926, org.springframework.test.context.transaction.TransactionalTestExecutionListener@2473b9ce, org.springframework.test.context.jdbc.SqlScriptsTestExecutionListener@60438a68]
08:27:53.644 [main] DEBUG org.springframework.test.annotation.ProfileValueUtils - Retrieved @ProfileValueSourceConfiguration [null] for test class [demo.ApplicationTests]
08:27:53.645 [main] DEBUG org.springframework.test.annotation.ProfileValueUtils - Retrieved ProfileValueSource type [class org.springframework.test.annotation.SystemProfileValueSource] for class [demo.ApplicationTests]
08:27:53.649 [main] DEBUG org.springframework.test.annotation.ProfileValueUtils - Retrieved @ProfileValueSourceConfiguration [null] for test class [demo.ApplicationTests]
08:27:53.649 [main] DEBUG org.springframework.test.annotation.ProfileValueUtils - Retrieved ProfileValueSource type [class org.springframework.test.annotation.SystemProfileValueSource] for class [demo.ApplicationTests]
08:27:53.657 [main] DEBUG org.springframework.test.annotation.ProfileValueUtils - Retrieved @ProfileValueSourceConfiguration [null] for test class [demo.ApplicationTests]
08:27:53.658 [main] DEBUG org.springframework.test.annotation.ProfileValueUtils - Retrieved ProfileValueSource type [class org.springframework.test.annotation.SystemProfileValueSource] for class [demo.ApplicationTests]
08:27:53.681 [main] DEBUG org.springframework.test.util.ReflectionTestUtils - Getting field 'mergedContextConfiguration' from target object [[DefaultTestContext@77a57272 testClass = ApplicationTests, testInstance = demo.ApplicationTests@7181ae3f, testMethod = [null], testException = [null], mergedContextConfiguration = [WebMergedContextConfiguration@46238e3f testClass = ApplicationTests, locations = '{}', classes = '{class demo.ConfigServerApplication}', contextInitializerClasses = '[]', activeProfiles = '{}', propertySourceLocations = '{}', propertySourceProperties = '{}', contextCustomizers = set[org.springframework.boot.test.context.filter.ExcludeFilterContextCustomizer@87f383f, org.springframework.boot.test.mock.mockito.MockitoContextCustomizer@0, org.springframework.boot.test.autoconfigure.properties.PropertyMappingContextCustomizer@0, org.springframework.boot.test.autoconfigure.web.servlet.WebDriverContextCustomizerFactory$Customizer@c81cdd1, org.springframework.test.context.web.socket.MockServerContainerContextCustomizer@2a70a3d8], resourceBasePath = 'src/main/webapp', contextLoader = 'org.springframework.boot.test.SpringApplicationContextLoader', parent = [null]]]] or target class [class org.springframework.test.context.support.DefaultTestContext]
08:27:53.688 [main] DEBUG org.springframework.test.util.ReflectionTestUtils - Setting field 'propertySourceProperties' of type [null] on target object [[WebMergedContextConfiguration@46238e3f testClass = ApplicationTests, locations = '{}', classes = '{class demo.ConfigServerApplication}', contextInitializerClasses = '[]', activeProfiles = '{}', propertySourceLocations = '{}', propertySourceProperties = '{}', contextCustomizers = set[org.springframework.boot.test.context.filter.ExcludeFilterContextCustomizer@87f383f, org.springframework.boot.test.mock.mockito.MockitoContextCustomizer@0, org.springframework.boot.test.autoconfigure.properties.PropertyMappingContextCustomizer@0, org.springframework.boot.test.autoconfigure.web.servlet.WebDriverContextCustomizerFactory$Customizer@c81cdd1, org.springframework.test.context.web.socket.MockServerContainerContextCustomizer@2a70a3d8], resourceBasePath = 'src/main/webapp', contextLoader = 'org.springframework.boot.test.SpringApplicationContextLoader', parent = [null]]] or target class [class org.springframework.test.context.web.WebMergedContextConfiguration] to value [[Ljava.lang.String;@76a4d6c]
08:27:53.690 [main] DEBUG org.springframework.test.context.support.DependencyInjectionTestExecutionListener - Performing dependency injection for test context [[DefaultTestContext@77a57272 testClass = ApplicationTests, testInstance = demo.ApplicationTests@7181ae3f, testMethod = [null], testException = [null], mergedContextConfiguration = [WebMergedContextConfiguration@46238e3f testClass = ApplicationTests, locations = '{}', classes = '{class demo.ConfigServerApplication}', contextInitializerClasses = '[]', activeProfiles = '{}', propertySourceLocations = '{}', propertySourceProperties = '{server.port=0, org.springframework.boot.test.WebIntegrationTest=true}', contextCustomizers = set[org.springframework.boot.test.context.filter.ExcludeFilterContextCustomizer@87f383f, org.springframework.boot.test.mock.mockito.MockitoContextCustomizer@0, org.springframework.boot.test.autoconfigure.properties.PropertyMappingContextCustomizer@0, org.springframework.boot.test.autoconfigure.web.servlet.WebDriverContextCustomizerFactory$Customizer@c81cdd1, org.springframework.test.context.web.socket.MockServerContainerContextCustomizer@2a70a3d8], resourceBasePath = 'src/main/webapp', contextLoader = 'org.springframework.boot.test.SpringApplicationContextLoader', parent = [null]]]].
08:27:54.118 [main] DEBUG org.springframework.core.env.StandardEnvironment - Adding [systemProperties] PropertySource with lowest search precedence
08:27:54.120 [main] DEBUG org.springframework.core.env.StandardEnvironment - Adding [systemEnvironment] PropertySource with lowest search precedence
08:27:54.121 [main] DEBUG org.springframework.core.env.StandardEnvironment - Initialized StandardEnvironment with PropertySources [systemProperties,systemEnvironment]
08:27:54.124 [main] DEBUG org.springframework.core.env.StandardEnvironment - Adding [integrationTest] PropertySource with search precedence immediately lower than [systemEnvironment]
2017-07-21 08:27:57.572  INFO 2550 --- [           main] s.c.a.AnnotationConfigApplicationContext : Refreshing org.springframework.context.annotation.AnnotationConfigApplicationContext@146dfe6: startup date [Fri Jul 21 08:27:57 UTC 2017]; root of context hierarchy
2017-07-21 08:27:59.096  INFO 2550 --- [           main] f.a.AutowiredAnnotationBeanPostProcessor : JSR-330 'javax.inject.Inject' annotation found and supported for autowiring
2017-07-21 08:27:59.220  INFO 2550 --- [           main] trationDelegate$BeanPostProcessorChecker : Bean 'configurationPropertiesRebinderAutoConfiguration' of type [class org.springframework.cloud.autoconfigure.ConfigurationPropertiesRebinderAutoConfiguration$$EnhancerBySpringCGLIB$$c0755e2] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)

  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/
 :: Spring Boot ::  (v1.4.1.BUILD-SNAPSHOT)

2017-07-21 08:28:01.140  INFO 2550 --- [           main] demo.ApplicationTests                    : No active profile set, falling back to default profiles: default
2017-07-21 08:28:01.177  INFO 2550 --- [           main] ationConfigEmbeddedWebApplicationContext : Refreshing org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@50b8ae8d: startup date [Fri Jul 21 08:28:01 UTC 2017]; parent: org.springframework.context.annotation.AnnotationConfigApplicationContext@146dfe6
2017-07-21 08:28:07.393  INFO 2550 --- [           main] o.s.b.f.config.PropertiesFactoryBean     : Loading properties file from URL [jar:file:/home/vagrant/.m2/repository/org/springframework/integration/spring-integration-core/4.3.2.RELEASE/spring-integration-core-4.3.2.RELEASE.jar!/META-INF/spring.integration.default.properties]
2017-07-21 08:28:07.412  INFO 2550 --- [           main] o.s.i.config.IntegrationRegistrar        : No bean named 'integrationHeaderChannelRegistry' has been explicitly defined. Therefore, a default DefaultHeaderChannelRegistry will be created.
2017-07-21 08:28:08.627 DEBUG 2550 --- [           main] o.s.cloud.context.scope.GenericScope     : Generating bean factory id from names: [$autoCreateChannelCandidates, DefaultConfiguringBeanFactoryPostProcessor, IntegrationConfigurationBeanFactoryPostProcessor, amqpAdmin, applicationContextIdFilter, archaiusEndpoint, auditEventRepository, auditListener, autoConfigurationReportEndpoint, basicErrorController, beanNameHandlerMapping, beanNameViewResolver, beansEndpoint, bindToAnnotationBeanPostProcessor, binderAwareChannelResolver, binderAwareRouterBeanPostProcessor, binderFactory, binderTypeRegistry, bindingService, bitbucketPropertyPathNotificationExtractor, busEndpoint, busJsonConverter, busPathMatcher, channelBindingServiceProperties, channelFactory, channelInitializer, characterEncodingFilter, codec, commonsFeatures, compositeMessageChannelConfigurer, compositeMessageConverterFactory, configClientHealthProperties, configClientProperties, configServerApplication, configServerHealthIndicator, configurableEnvironmentConfiguration, configurationPropertiesBeans, configurationPropertiesRebinder, configurationPropertiesReportEndpoint, contextRefresher, contextStartAfterRefreshListener, conventionErrorViewResolver, converterRegistrar, datatypeChannelMessageConverter, default.org.springframework.cloud.netflix.ribbon.RibbonAutoConfiguration.RibbonClientSpecification, default.org.springframework.cloud.netflix.ribbon.eureka.RibbonEurekaAutoConfiguration.RibbonClientSpecification, defaultMetricsTagProvider, defaultServletHandlerMapping, defaultViewResolver, discoveryClient, discoveryClientHealthIndicator, discoveryClientOptionalArgs, discoveryCompositeHealthIndicator, diskSpaceHealthIndicator, diskSpaceHealthIndicatorProperties, dispatcherServlet, dispatcherServletRegistration, dumpEndpoint, duplicateServerPropertiesDetector, dynamicDestinationsBindable, embeddedServletContainerCustomizerBeanPostProcessor, enableConfigServerMarker, encrypt-org.springframework.cloud.bootstrap.encrypt.KeyProperties, encryptionController, endpointHandlerMapping, endpoints-org.springframework.boot.actuate.endpoint.EndpointProperties, endpoints.cors-org.springframework.boot.actuate.autoconfigure.EndpointCorsProperties, endpoints.health-org.springframework.boot.actuate.autoconfigure.HealthMvcEndpointProperties, endpoints.metrics.filter-org.springframework.boot.actuate.autoconfigure.MetricFilterProperties, envInfoContributor, environmentBusEndpoint, environmentChangeListener, environmentController, environmentEncryptor, environmentEndpoint, environmentManager, environmentManagerEndpoint, environmentMvcEndpoint, environmentRepository, environmentWatch, error, errorAttributes, errorPageCustomizer, errorPageRegistrarBeanPostProcessor, eurekaApplicationInfoManager, eurekaClient, eurekaClientConfigBean, eurekaFeature, eurekaHealthIndicator, eurekaInstanceConfigBean, faviconHandlerMapping, faviconRequestHandler, featuresEndpoint, fileRegistrar, gitInfoContributor, gitProperties, githubPropertyPathNotificationExtractor, gitlabPropertyPathNotificationExtractor, globalChannelInterceptorProcessor, gson, handlerExceptionResolver, healthAggregator, healthEndpoint, healthMvcEndpoint, heapdumpMvcEndpoint, hiddenHttpMethodFilter, httpPutFormContentFilter, httpRequestHandlerAdapter, hystrixHealthIndicator, inetUtils, inetUtilsProperties, infoEndpoint, inputBindingLifecycle, integrationConversionService, integrationEvaluationContext, integrationGlobalProperties, integrationHeaderChannelRegistry, integrationLifecycleRoleController, jacksonObjectMapper, jacksonObjectMapperBuilder, jsonComponentModule, jsonPath, loadBalancedRestTemplateInitializer, loadBalancedRetryPolicyFactory, loadBalancerClient, loadBalancerRequestFactory, localeCharsetMappingsCustomizer, loggingRebinder, management.health.status-org.springframework.boot.actuate.autoconfigure.HealthIndicatorProperties, management.info-org.springframework.boot.actuate.autoconfigure.InfoContributorProperties, management.trace-org.springframework.boot.actuate.trace.TraceProperties, managementContextResolver, managementServerProperties, managementServletContext, mappingJackson2HttpMessageConverter, messageBuilderFactory, messageConverterConfigurer, messageConverters, messageHandlerMethodFactory, metricReaderPublicMetrics, metricWritersMetricExporter, metricsEndpoint, metricsFilter, metricsMvcEndpoint, monitorCache, monitorRegistry, multipartConfigElement, multipartResolver, mvcContentNegotiationManager, mvcConversionService, mvcEndpoints, mvcPathMatcher, mvcResourceUrlProvider, mvcUriComponentsContributor, mvcUrlPathHelper, mvcValidator, mvcViewResolver, observableMVCConfiguration, org.springframework.amqp.rabbit.annotation.RabbitBootstrapConfiguration, org.springframework.amqp.rabbit.config.internalRabbitListenerAnnotationProcessor, org.springframework.amqp.rabbit.config.internalRabbitListenerEndpointRegistry, org.springframework.boot.actuate.autoconfigure.AuditAutoConfiguration, org.springframework.boot.actuate.autoconfigure.AuditAutoConfiguration$AuditEventRepositoryConfiguration, org.springframework.boot.actuate.autoconfigure.EndpointAutoConfiguration, org.springframework.boot.actuate.autoconfigure.EndpointAutoConfiguration$RequestMappingEndpointConfiguration, org.springframework.boot.actuate.autoconfigure.EndpointWebMvcAutoConfiguration, org.springframework.boot.actuate.autoconfigure.EndpointWebMvcAutoConfiguration$ApplicationContextFilterConfiguration, org.springframework.boot.actuate.autoconfigure.EndpointWebMvcAutoConfiguration$EndpointWebMvcConfiguration, org.springframework.boot.actuate.autoconfigure.EndpointWebMvcManagementContextConfiguration, org.springframework.boot.actuate.autoconfigure.HealthIndicatorAutoConfiguration, org.springframework.boot.actuate.autoconfigure.HealthIndicatorAutoConfiguration$DiskSpaceHealthIndicatorConfiguration, org.springframework.boot.actuate.autoconfigure.HealthIndicatorAutoConfiguration$RabbitHealthIndicatorConfiguration, org.springframework.boot.actuate.autoconfigure.InfoContributorAutoConfiguration, org.springframework.boot.actuate.autoconfigure.ManagementServerPropertiesAutoConfiguration, org.springframework.boot.actuate.autoconfigure.MetricExportAutoConfiguration, org.springframework.boot.actuate.autoconfigure.MetricExportAutoConfiguration$MetricExportPropertiesConfiguration, org.springframework.boot.actuate.autoconfigure.MetricExportAutoConfiguration$StatsdConfiguration, org.springframework.boot.actuate.autoconfigure.MetricFilterAutoConfiguration, org.springframework.boot.actuate.autoconfigure.MetricRepositoryAutoConfiguration, org.springframework.boot.actuate.autoconfigure.PublicMetricsAutoConfiguration, org.springframework.boot.actuate.autoconfigure.PublicMetricsAutoConfiguration$TomcatMetricsConfiguration, org.springframework.boot.actuate.autoconfigure.TraceRepositoryAutoConfiguration, org.springframework.boot.actuate.autoconfigure.TraceWebFilterAutoConfiguration, org.springframework.boot.autoconfigure.AutoConfigurationPackages, org.springframework.boot.autoconfigure.PropertyPlaceholderAutoConfiguration, org.springframework.boot.autoconfigure.amqp.RabbitAnnotationDrivenConfiguration, org.springframework.boot.autoconfigure.amqp.RabbitAnnotationDrivenConfiguration$EnableRabbitConfiguration, org.springframework.boot.autoconfigure.amqp.RabbitAutoConfiguration, org.springframework.boot.autoconfigure.amqp.RabbitAutoConfiguration$MessagingTemplateConfiguration, org.springframework.boot.autoconfigure.amqp.RabbitAutoConfiguration$RabbitConnectionFactoryCreator, org.springframework.boot.autoconfigure.amqp.RabbitAutoConfiguration$RabbitTemplateConfiguration, org.springframework.boot.autoconfigure.condition.BeanTypeRegistry, org.springframework.boot.autoconfigure.context.ConfigurationPropertiesAutoConfiguration, org.springframework.boot.autoconfigure.dao.PersistenceExceptionTranslationAutoConfiguration, org.springframework.boot.autoconfigure.gson.GsonAutoConfiguration, org.springframework.boot.autoconfigure.info.ProjectInfoAutoConfiguration, org.springframework.boot.autoconfigure.integration.IntegrationAutoConfiguration, org.springframework.boot.autoconfigure.integration.IntegrationAutoConfiguration$IntegrationConfiguration, org.springframework.boot.autoconfigure.internalCachingMetadataReaderFactory, org.springframework.boot.autoconfigure.jackson.JacksonAutoConfiguration, org.springframework.boot.autoconfigure.jackson.JacksonAutoConfiguration$Jackson2ObjectMapperBuilderCustomizerConfiguration, org.springframework.boot.autoconfigure.jackson.JacksonAutoConfiguration$JacksonObjectMapperBuilderConfiguration, org.springframework.boot.autoconfigure.jackson.JacksonAutoConfiguration$JacksonObjectMapperConfiguration, org.springframework.boot.autoconfigure.web.DispatcherServletAutoConfiguration, org.springframework.boot.autoconfigure.web.DispatcherServletAutoConfiguration$DispatcherServletConfiguration, org.springframework.boot.autoconfigure.web.DispatcherServletAutoConfiguration$DispatcherServletRegistrationConfiguration, org.springframework.boot.autoconfigure.web.EmbeddedServletContainerAutoConfiguration, org.springframework.boot.autoconfigure.web.EmbeddedServletContainerAutoConfiguration$EmbeddedTomcat, org.springframework.boot.autoconfigure.web.ErrorMvcAutoConfiguration, org.springframework.boot.autoconfigure.web.ErrorMvcAutoConfiguration$WhitelabelErrorViewConfiguration, org.springframework.boot.autoconfigure.web.GsonHttpMessageConvertersConfiguration, org.springframework.boot.autoconfigure.web.HttpEncodingAutoConfiguration, org.springframework.boot.autoconfigure.web.HttpMessageConvertersAutoConfiguration, org.springframework.boot.autoconfigure.web.HttpMessageConvertersAutoConfiguration$StringHttpMessageConverterConfiguration, org.springframework.boot.autoconfigure.web.JacksonHttpMessageConvertersConfiguration, org.springframework.boot.autoconfigure.web.JacksonHttpMessageConvertersConfiguration$MappingJackson2HttpMessageConverterConfiguration, org.springframework.boot.autoconfigure.web.MultipartAutoConfiguration, org.springframework.boot.autoconfigure.web.ServerPropertiesAutoConfiguration, org.springframework.boot.autoconfigure.web.WebClientAutoConfiguration, org.springframework.boot.autoconfigure.web.WebClientAutoConfiguration$RestTemplateConfiguration, org.springframework.boot.autoconfigure.web.WebMvcAutoConfiguration, org.springframework.boot.autoconfigure.web.WebMvcAutoConfiguration$EnableWebMvcConfiguration, org.springframework.boot.autoconfigure.web.WebMvcAutoConfiguration$WebMvcAutoConfigurationAdapter, org.springframework.boot.autoconfigure.web.WebMvcAutoConfiguration$WebMvcAutoConfigurationAdapter$FaviconConfiguration, org.springframework.boot.autoconfigure.websocket.WebSocketAutoConfiguration, org.springframework.boot.autoconfigure.websocket.WebSocketAutoConfiguration$TomcatWebSocketConfiguration, org.springframework.boot.context.properties.ConfigurationPropertiesBindingPostProcessor, org.springframework.boot.context.properties.ConfigurationPropertiesBindingPostProcessor.store, org.springframework.cloud.autoconfigure.ConfigurationPropertiesRebinderAutoConfiguration, org.springframework.cloud.autoconfigure.LifecycleMvcEndpointAutoConfiguration, org.springframework.cloud.autoconfigure.RefreshAutoConfiguration, org.springframework.cloud.autoconfigure.RefreshAutoConfiguration$RefreshScopeConfiguration, org.springframework.cloud.autoconfigure.RefreshEndpointAutoConfiguration, org.springframework.cloud.autoconfigure.RefreshEndpointAutoConfiguration$RefreshEndpointConfiguration, org.springframework.cloud.autoconfigure.RefreshEndpointAutoConfiguration$RestartEndpointWithIntegration, org.springframework.cloud.bus.BusAutoConfiguration, org.springframework.cloud.bus.BusAutoConfiguration$BusEndpointConfiguration, org.springframework.cloud.bus.BusAutoConfiguration$BusEnvironmentConfiguration, org.springframework.cloud.bus.BusAutoConfiguration$BusEnvironmentConfiguration$EnvironmentBusEndpointConfiguration, org.springframework.cloud.bus.BusAutoConfiguration$BusRefreshConfiguration, org.springframework.cloud.bus.BusAutoConfiguration$BusRefreshConfiguration$BusRefreshEndpointConfiguration, org.springframework.cloud.bus.BusAutoConfiguration$MatcherConfiguration, org.springframework.cloud.bus.SpringCloudBusClient, org.springframework.cloud.bus.jackson.BusJacksonAutoConfiguration, org.springframework.cloud.client.CommonsClientAutoConfiguration, org.springframework.cloud.client.CommonsClientAutoConfiguration$ActuatorConfiguration, org.springframework.cloud.client.loadbalancer.LoadBalancerAutoConfiguration, org.springframework.cloud.client.loadbalancer.LoadBalancerAutoConfiguration$RetryAutoConfiguration, org.springframework.cloud.commons.util.UtilAutoConfiguration, org.springframework.cloud.config.client.ConfigClientAutoConfiguration, org.springframework.cloud.config.monitor.EnvironmentMonitorAutoConfiguration, org.springframework.cloud.config.monitor.EnvironmentMonitorAutoConfiguration$PropertyPathNotificationExtractorConfiguration, org.springframework.cloud.config.monitor.FileMonitorConfiguration, org.springframework.cloud.config.server.config.ConfigServerAutoConfiguration, org.springframework.cloud.config.server.config.ConfigServerConfiguration, org.springframework.cloud.config.server.config.ConfigServerEncryptionConfiguration, org.springframework.cloud.config.server.config.ConfigServerMvcConfiguration, org.springframework.cloud.config.server.config.EncryptionAutoConfiguration, org.springframework.cloud.config.server.config.EncryptionAutoConfiguration$KeyStoreConfiguration, org.springframework.cloud.config.server.config.EnvironmentRepositoryConfiguration, org.springframework.cloud.config.server.config.EnvironmentRepositoryConfiguration$DefaultEnvironmentWatch, org.springframework.cloud.config.server.config.EnvironmentRepositoryConfiguration$GitRepositoryConfiguration, org.springframework.cloud.config.server.config.ResourceRepositoryConfiguration, org.springframework.cloud.config.server.config.SingleEncryptorAutoConfiguration, org.springframework.cloud.netflix.archaius.ArchaiusAutoConfiguration, org.springframework.cloud.netflix.archaius.ArchaiusAutoConfiguration$ArchaiusEndpointConfiguration, org.springframework.cloud.netflix.archaius.ArchaiusAutoConfiguration$PropagateEventsConfiguration, org.springframework.cloud.netflix.eureka.EurekaClientAutoConfiguration, org.springframework.cloud.netflix.eureka.EurekaClientAutoConfiguration$RefreshableEurekaClientConfiguration, org.springframework.cloud.netflix.eureka.EurekaDiscoveryClientConfiguration, org.springframework.cloud.netflix.eureka.EurekaDiscoveryClientConfiguration$EurekaClientConfigurationRefresher, org.springframework.cloud.netflix.eureka.EurekaDiscoveryClientConfiguration$EurekaHealthIndicatorConfiguration, org.springframework.cloud.netflix.eureka.config.EurekaClientConfigServerAutoConfiguration, org.springframework.cloud.netflix.hystrix.HystrixAutoConfiguration, org.springframework.cloud.netflix.metrics.MetricsInterceptorConfiguration, org.springframework.cloud.netflix.metrics.MetricsInterceptorConfiguration$MetricsWebResourceConfiguration, org.springframework.cloud.netflix.metrics.servo.ServoMetricsAutoConfiguration, org.springframework.cloud.netflix.metrics.servo.ServoMetricsAutoConfiguration$MetricsTagConfiguration, org.springframework.cloud.netflix.ribbon.RibbonAutoConfiguration, org.springframework.cloud.netflix.ribbon.eureka.RibbonEurekaAutoConfiguration, org.springframework.cloud.netflix.rx.RxJavaAutoConfiguration, org.springframework.cloud.netflix.rx.RxJavaAutoConfiguration$RxJavaReturnValueHandlerConfig, org.springframework.cloud.stream.config.BinderFactoryConfiguration, org.springframework.cloud.stream.config.BindingServiceConfiguration, org.springframework.cloud.stream.config.BindingServiceConfiguration$PostProcessorConfiguration, org.springframework.cloud.stream.config.SpelExpressionConverterConfiguration, org.springframework.cloud.stream.config.codec.kryo.KryoCodecAutoConfiguration, org.springframework.context.annotation.ConfigurationClassPostProcessor.enhancedConfigurationProcessor, org.springframework.context.annotation.ConfigurationClassPostProcessor.importAwareProcessor, org.springframework.context.annotation.internalAutowiredAnnotationProcessor, org.springframework.context.annotation.internalCommonAnnotationProcessor, org.springframework.context.annotation.internalConfigurationAnnotationProcessor, org.springframework.context.annotation.internalRequiredAnnotationProcessor, org.springframework.context.annotation.internalScheduledAnnotationProcessor, org.springframework.context.event.internalEventListenerFactory, org.springframework.context.event.internalEventListenerProcessor, org.springframework.integration.expression.IntegrationEvaluationContextAwareBeanPostProcessor#0, org.springframework.integration.internalMessagingAnnotationPostProcessor, org.springframework.integration.internalPublisherAnnotationBeanPostProcessor, org.springframework.scheduling.annotation.SchedulingConfiguration, outputBindingLifecycle, pauseEndpoint, pauseMvcEndpoint, persistenceExceptionTranslationPostProcessor, preserveErrorControllerTargetClassPostProcessor, propertiesFactory, propertyAccessorBeanPostProcessor, propertyPathEndpoint, propertySourcesPlaceholderConfigurer, rabbitConnectionFactory, rabbitHealthIndicator, rabbitListenerContainerFactory, rabbitListenerContainerFactoryConfigurer, rabbitMessagingTemplate, rabbitTemplate, refreshBusEndpoint, refreshEndpoint, refreshEventListener, refreshListener, refreshMvcEndpoint, refreshScope, refreshScopeHealthIndicator, requestContextFilter, requestMappingEndpoint, requestMappingHandlerAdapter, requestMappingHandlerMapping, resourceController, resourceHandlerMapping, resourceRepository, restTemplateBuilder, restTemplateCustomizer, restartEndpoint, restartMvcEndpoint, resumeEndpoint, resumeMvcEndpoint, retryTemplate, ribbonFeature, ribbonInterceptor, rxFeature, scopedTarget.eurekaApplicationInfoManager, scopedTarget.eurekaClient, serverProperties, serviceMatcher, servoMetricNaming, servoMetricReader, servoMetricServices, servoMetricsConfig, servoMonitoringWebResourceInterceptor, servoPublicMetrics, shutdownEndpoint, simpleControllerHandlerAdapter, singleReturnValueHandler, spelConverter, spring.cloud.bus-org.springframework.cloud.bus.BusProperties, spring.cloud.codec.kryo-org.springframework.cloud.stream.config.codec.kryo.KryoCodecProperties, spring.cloud.config.server-org.springframework.cloud.config.server.config.ConfigServerProperties, spring.cloud.loadbalancer.retry-org.springframework.cloud.client.loadbalancer.LoadBalancerRetryProperties, spring.cloud.stream-org.springframework.cloud.stream.config.BindingServiceProperties, spring.http.encoding-org.springframework.boot.autoconfigure.web.HttpEncodingProperties, spring.http.multipart-org.springframework.boot.autoconfigure.web.MultipartProperties, spring.info-org.springframework.boot.autoconfigure.info.ProjectInfoProperties, spring.jackson-org.springframework.boot.autoconfigure.jackson.JacksonProperties, spring.metrics.export-org.springframework.boot.actuate.metrics.export.MetricExportProperties, spring.mvc-org.springframework.boot.autoconfigure.web.WebMvcProperties, spring.rabbitmq-org.springframework.boot.autoconfigure.amqp.RabbitProperties, spring.resources-org.springframework.boot.autoconfigure.web.ResourceProperties, springClientFactory, springCloudBusInput, springCloudBusOutput, standardJacksonObjectMapperBuilderCustomizer, stringHttpMessageConverter, systemPublicMetrics, textEncryptorLocator, toStringFriendlyJsonNodeToStringConverter, tomcatEmbeddedServletContainerFactory, tomcatPublicMetrics, traceEndpoint, traceRepository, viewControllerHandlerMapping, viewResolver, webRequestLoggingFilter, websocketContainerCustomizer, welcomePageHandlerMapping]
2017-07-21 08:28:08.635  INFO 2550 --- [           main] o.s.cloud.context.scope.GenericScope     : BeanFactory id=eae4c6e5-6d06-30c8-8b8c-fb2026cbc67b
2017-07-21 08:28:08.647  INFO 2550 --- [           main] faultConfiguringBeanFactoryPostProcessor : No bean named 'errorChannel' has been explicitly defined. Therefore, a default PublishSubscribeChannel will be created.
2017-07-21 08:28:08.664  INFO 2550 --- [           main] faultConfiguringBeanFactoryPostProcessor : No bean named 'taskScheduler' has been explicitly defined. Therefore, a default ThreadPoolTaskScheduler will be created.
2017-07-21 08:28:08.704  INFO 2550 --- [           main] f.a.AutowiredAnnotationBeanPostProcessor : JSR-330 'javax.inject.Inject' annotation found and supported for autowiring
2017-07-21 08:28:08.766  INFO 2550 --- [           main] trationDelegate$BeanPostProcessorChecker : Bean 'org.springframework.amqp.rabbit.annotation.RabbitBootstrapConfiguration' of type [class org.springframework.amqp.rabbit.annotation.RabbitBootstrapConfiguration$$EnhancerBySpringCGLIB$$b6591113] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)
2017-07-21 08:28:09.366  INFO 2550 --- [           main] o.s.b.f.config.PropertiesFactoryBean     : Loading properties file from URL [jar:file:/home/vagrant/.m2/repository/org/springframework/integration/spring-integration-core/4.3.2.RELEASE/spring-integration-core-4.3.2.RELEASE.jar!/META-INF/spring.integration.default.properties]
2017-07-21 08:28:09.368  INFO 2550 --- [           main] trationDelegate$BeanPostProcessorChecker : Bean 'integrationGlobalProperties' of type [class org.springframework.beans.factory.config.PropertiesFactoryBean] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)
2017-07-21 08:28:09.369  INFO 2550 --- [           main] trationDelegate$BeanPostProcessorChecker : Bean 'integrationGlobalProperties' of type [class java.util.Properties] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)
2017-07-21 08:28:09.559  INFO 2550 --- [           main] trationDelegate$BeanPostProcessorChecker : Bean 'org.springframework.cloud.autoconfigure.ConfigurationPropertiesRebinderAutoConfiguration' of type [class org.springframework.cloud.autoconfigure.ConfigurationPropertiesRebinderAutoConfiguration$$EnhancerBySpringCGLIB$$c0755e2] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)
2017-07-21 08:28:09.604  INFO 2550 --- [           main] trationDelegate$BeanPostProcessorChecker : Bean 'org.springframework.cloud.stream.config.BindingServiceConfiguration$PostProcessorConfiguration' of type [class org.springframework.cloud.stream.config.BindingServiceConfiguration$PostProcessorConfiguration$$EnhancerBySpringCGLIB$$b1b4f0ea] is not eligible for getting processed by all BeanPostProcessors (for example: not eligible for auto-proxying)
2017-07-21 08:28:11.395  INFO 2550 --- [           main] s.b.c.e.t.TomcatEmbeddedServletContainer : Tomcat initialized with port(s): 0 (http)
2017-07-21 08:28:11.448  INFO 2550 --- [           main] o.apache.catalina.core.StandardService   : Starting service Tomcat
2017-07-21 08:28:11.453  INFO 2550 --- [           main] org.apache.catalina.core.StandardEngine  : Starting Servlet Engine: Apache Tomcat/8.5.5
2017-07-21 08:28:12.077  INFO 2550 --- [ost-startStop-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring embedded WebApplicationContext
2017-07-21 08:28:12.078  INFO 2550 --- [ost-startStop-1] o.s.web.context.ContextLoader            : Root WebApplicationContext: initialization completed in 10901 ms
2017-07-21 08:28:13.202  INFO 2550 --- [ost-startStop-1] o.s.b.w.servlet.ServletRegistrationBean  : Mapping servlet: 'dispatcherServlet' to [/]
2017-07-21 08:28:13.233  INFO 2550 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'metricsFilter' to: [/*]
2017-07-21 08:28:13.234  INFO 2550 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'characterEncodingFilter' to: [/*]
2017-07-21 08:28:13.235  INFO 2550 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'hiddenHttpMethodFilter' to: [/*]
2017-07-21 08:28:13.237  INFO 2550 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'httpPutFormContentFilter' to: [/*]
2017-07-21 08:28:13.237  INFO 2550 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'requestContextFilter' to: [/*]
2017-07-21 08:28:13.238  INFO 2550 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'webRequestLoggingFilter' to: [/*]
2017-07-21 08:28:13.238  INFO 2550 --- [ost-startStop-1] o.s.b.w.servlet.FilterRegistrationBean   : Mapping filter: 'applicationContextIdFilter' to: [/*]
2017-07-21 08:28:16.085  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerAdapter : Looking for @ControllerAdvice: org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@50b8ae8d: startup date [Fri Jul 21 08:28:01 UTC 2017]; parent: org.springframework.context.annotation.AnnotationConfigApplicationContext@146dfe6
2017-07-21 08:28:16.630  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error],produces=[text/html]}" onto public org.springframework.web.servlet.ModelAndView org.springframework.boot.autoconfigure.web.BasicErrorController.errorHtml(javax.servlet.http.HttpServletRequest,javax.servlet.http.HttpServletResponse)
2017-07-21 08:28:16.634  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/error]}" onto public org.springframework.http.ResponseEntity<java.util.Map<java.lang.String, java.lang.Object>> org.springframework.boot.autoconfigure.web.BasicErrorController.error(javax.servlet.http.HttpServletRequest)
2017-07-21 08:28:16.645  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/monitor],methods=[POST],consumes=[application/x-www-form-urlencoded]}" onto public java.util.Set<java.lang.String> org.springframework.cloud.config.monitor.PropertyPathEndpoint.notifyByForm(org.springframework.http.HttpHeaders,java.util.List<java.lang.String>)
2017-07-21 08:28:16.645  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/monitor],methods=[POST]}" onto public java.util.Set<java.lang.String> org.springframework.cloud.config.monitor.PropertyPathEndpoint.notifyByPath(org.springframework.http.HttpHeaders,java.util.Map<java.lang.String, java.lang.Object>)
2017-07-21 08:28:16.662  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/encrypt],methods=[POST]}" onto public java.lang.String org.springframework.cloud.config.server.encryption.EncryptionController.encrypt(java.lang.String,org.springframework.http.MediaType)
2017-07-21 08:28:16.663  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/encrypt/{name}/{profiles}],methods=[POST]}" onto public java.lang.String org.springframework.cloud.config.server.encryption.EncryptionController.encrypt(java.lang.String,java.lang.String,java.lang.String,org.springframework.http.MediaType)
2017-07-21 08:28:16.664  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/decrypt/{name}/{profiles}],methods=[POST]}" onto public java.lang.String org.springframework.cloud.config.server.encryption.EncryptionController.decrypt(java.lang.String,java.lang.String,java.lang.String,org.springframework.http.MediaType)
2017-07-21 08:28:16.665  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/decrypt],methods=[POST]}" onto public java.lang.String org.springframework.cloud.config.server.encryption.EncryptionController.decrypt(java.lang.String,org.springframework.http.MediaType)
2017-07-21 08:28:16.666  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/encrypt/status],methods=[GET]}" onto public java.util.Map<java.lang.String, java.lang.Object> org.springframework.cloud.config.server.encryption.EncryptionController.status()
2017-07-21 08:28:16.667  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/key],methods=[GET]}" onto public java.lang.String org.springframework.cloud.config.server.encryption.EncryptionController.getPublicKey()
2017-07-21 08:28:16.668  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/key/{name}/{profiles}],methods=[GET]}" onto public java.lang.String org.springframework.cloud.config.server.encryption.EncryptionController.getPublicKey(java.lang.String,java.lang.String)
2017-07-21 08:28:16.682  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{name}-{profiles}.properties],methods=[GET]}" onto public org.springframework.http.ResponseEntity<java.lang.String> org.springframework.cloud.config.server.environment.EnvironmentController.properties(java.lang.String,java.lang.String,boolean) throws java.io.IOException
2017-07-21 08:28:16.683  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{name}-{profiles}.yml || /{name}-{profiles}.yaml],methods=[GET]}" onto public org.springframework.http.ResponseEntity<java.lang.String> org.springframework.cloud.config.server.environment.EnvironmentController.yaml(java.lang.String,java.lang.String,boolean) throws java.lang.Exception
2017-07-21 08:28:16.684  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{name}/{profiles:.*[^-].*}],methods=[GET]}" onto public org.springframework.cloud.config.environment.Environment org.springframework.cloud.config.server.environment.EnvironmentController.defaultLabel(java.lang.String,java.lang.String)
2017-07-21 08:28:16.685  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{label}/{name}-{profiles}.properties],methods=[GET]}" onto public org.springframework.http.ResponseEntity<java.lang.String> org.springframework.cloud.config.server.environment.EnvironmentController.labelledProperties(java.lang.String,java.lang.String,java.lang.String,boolean) throws java.io.IOException
2017-07-21 08:28:16.686  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{name}-{profiles}.json],methods=[GET]}" onto public org.springframework.http.ResponseEntity<java.lang.String> org.springframework.cloud.config.server.environment.EnvironmentController.jsonProperties(java.lang.String,java.lang.String,boolean) throws java.lang.Exception
2017-07-21 08:28:16.687  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{label}/{name}-{profiles}.json],methods=[GET]}" onto public org.springframework.http.ResponseEntity<java.lang.String> org.springframework.cloud.config.server.environment.EnvironmentController.labelledJsonProperties(java.lang.String,java.lang.String,java.lang.String,boolean) throws java.lang.Exception
2017-07-21 08:28:16.688  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{label}/{name}-{profiles}.yml || /{label}/{name}-{profiles}.yaml],methods=[GET]}" onto public org.springframework.http.ResponseEntity<java.lang.String> org.springframework.cloud.config.server.environment.EnvironmentController.labelledYaml(java.lang.String,java.lang.String,java.lang.String,boolean) throws java.lang.Exception
2017-07-21 08:28:16.689  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{name}/{profiles}/{label:.*}],methods=[GET]}" onto public org.springframework.cloud.config.environment.Environment org.springframework.cloud.config.server.environment.EnvironmentController.labelled(java.lang.String,java.lang.String,java.lang.String)
2017-07-21 08:28:16.698  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{name}/{profile}/{label}/**],methods=[GET],produces=[application/octet-stream]}" onto public synchronized byte[] org.springframework.cloud.config.server.resource.ResourceController.binary(java.lang.String,java.lang.String,java.lang.String,javax.servlet.http.HttpServletRequest) throws java.io.IOException
2017-07-21 08:28:16.700  INFO 2550 --- [           main] s.w.s.m.m.a.RequestMappingHandlerMapping : Mapped "{[/{name}/{profile}/{label}/**],methods=[GET]}" onto public java.lang.String org.springframework.cloud.config.server.resource.ResourceController.resolve(java.lang.String,java.lang.String,java.lang.String,javax.servlet.http.HttpServletRequest) throws java.io.IOException
2017-07-21 08:28:16.886  INFO 2550 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/webjars/**] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2017-07-21 08:28:16.886  INFO 2550 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/**] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2017-07-21 08:28:17.188  INFO 2550 --- [           main] o.s.w.s.handler.SimpleUrlHandlerMapping  : Mapped URL path [/**/favicon.ico] onto handler of type [class org.springframework.web.servlet.resource.ResourceHttpRequestHandler]
2017-07-21 08:28:19.899  INFO 2550 --- [           main] o.s.s.c.ThreadPoolTaskScheduler          : Initializing ExecutorService  'taskScheduler'
2017-07-21 08:28:21.971  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/resume || /admin/resume.json],methods=[POST]}" onto public java.lang.Object org.springframework.cloud.endpoint.GenericPostableMvcEndpoint.invoke()
2017-07-21 08:28:21.974  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/beans || /admin/beans.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:21.975  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/mappings || /admin/mappings.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:21.980  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/metrics/{name:.*}],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.MetricsMvcEndpoint.value(java.lang.String)
2017-07-21 08:28:21.981  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/metrics || /admin/metrics.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:21.983  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/dump || /admin/dump.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:21.984  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/archaius || /admin/archaius.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:21.986  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/pause || /admin/pause.json],methods=[POST]}" onto public java.lang.Object org.springframework.cloud.endpoint.GenericPostableMvcEndpoint.invoke()
2017-07-21 08:28:21.988  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/heapdump || /admin/heapdump.json],methods=[GET],produces=[application/octet-stream]}" onto public void org.springframework.boot.actuate.endpoint.mvc.HeapdumpMvcEndpoint.invoke(boolean,javax.servlet.http.HttpServletRequest,javax.servlet.http.HttpServletResponse) throws java.io.IOException,javax.servlet.ServletException
2017-07-21 08:28:21.991  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/configprops || /admin/configprops.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:21.993  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/features || /admin/features.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:21.998  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/trace || /admin/trace.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:22.000  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/refresh || /admin/refresh.json],methods=[POST]}" onto public java.lang.Object org.springframework.cloud.endpoint.GenericPostableMvcEndpoint.invoke()
2017-07-21 08:28:22.001  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/bus/refresh],methods=[POST]}" onto public void org.springframework.cloud.bus.endpoint.RefreshBusEndpoint.refresh(java.lang.String)
2017-07-21 08:28:22.002  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/autoconfig || /admin/autoconfig.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:22.004  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/env],methods=[POST]}" onto public java.lang.Object org.springframework.cloud.context.environment.EnvironmentManagerMvcEndpoint.value(java.util.Map<java.lang.String, java.lang.String>)
2017-07-21 08:28:22.005  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/env/reset],methods=[POST]}" onto public java.util.Map<java.lang.String, java.lang.Object> org.springframework.cloud.context.environment.EnvironmentManagerMvcEndpoint.reset()
2017-07-21 08:28:22.006  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/health || /admin/health.json],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.HealthMvcEndpoint.invoke(java.security.Principal)
2017-07-21 08:28:22.007  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/info || /admin/info.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:22.010  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/env/{name:.*}],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EnvironmentMvcEndpoint.value(java.lang.String)
2017-07-21 08:28:22.011  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/env || /admin/env.json],methods=[GET],produces=[application/json]}" onto public java.lang.Object org.springframework.boot.actuate.endpoint.mvc.EndpointMvcAdapter.invoke()
2017-07-21 08:28:22.013  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/restart || /admin/restart.json],methods=[POST]}" onto public java.lang.Object org.springframework.cloud.context.restart.RestartMvcEndpoint.invoke()
2017-07-21 08:28:22.015  INFO 2550 --- [           main] o.s.b.a.e.mvc.EndpointHandlerMapping     : Mapped "{[/admin/bus/env],methods=[POST]}" onto public void org.springframework.cloud.bus.endpoint.EnvironmentBusEndpoint.env(java.util.Map<java.lang.String, java.lang.String>,java.lang.String)
2017-07-21 08:28:23.528 DEBUG 2550 --- [           main] s.c.s.c.CompositeMessageConverterFactory : Ommitted org.springframework.integration.support.converter.DefaultDatatypeChannelMessageConverter@4b9dbf07 of type class org.springframework.integration.support.converter.DefaultDatatypeChannelMessageConverter for 'application/json' as it is not an AbstractMessageConverter
2017-07-21 08:28:23.788  INFO 2550 --- [           main] o.s.integration.channel.DirectChannel    : Channel 'configserver:0.springCloudBusInput' has 1 subscriber(s).
2017-07-21 08:28:24.577  WARN 2550 --- [           main] c.n.c.sources.URLConfigurationSource     : No URLs will be polled as dynamic configuration sources.
2017-07-21 08:28:24.577  INFO 2550 --- [           main] c.n.c.sources.URLConfigurationSource     : To enable URLs as dynamic configuration sources, define System property archaius.configurationSource.additionalUrls or make config.properties available on classpath.
2017-07-21 08:28:24.639  WARN 2550 --- [           main] c.n.c.sources.URLConfigurationSource     : No URLs will be polled as dynamic configuration sources.
2017-07-21 08:28:24.640  INFO 2550 --- [           main] c.n.c.sources.URLConfigurationSource     : To enable URLs as dynamic configuration sources, define System property archaius.configurationSource.additionalUrls or make config.properties available on classpath.
2017-07-21 08:28:25.249  INFO 2550 --- [           main] o.s.i.codec.kryo.CompositeKryoRegistrar  : configured Kryo registration [40, java.io.File] with serializer org.springframework.integration.codec.kryo.FileSerializer
2017-07-21 08:28:25.833  INFO 2550 --- [           main] o.s.c.support.DefaultLifecycleProcessor  : Starting beans in phase -2147482648
2017-07-21 08:28:25.836 DEBUG 2550 --- [           main] o.s.c.s.binding.BindableProxyFactory     : Binding outputs for :interface org.springframework.cloud.bus.SpringCloudBusClient
2017-07-21 08:28:25.836 DEBUG 2550 --- [           main] o.s.c.s.binding.BindableProxyFactory     : Binding :interface org.springframework.cloud.bus.SpringCloudBusClient:springCloudBusOutput
2017-07-21 08:28:26.346  INFO 2550 --- [           main] o.a.maven.surefire.booter.ForkedBooter   : No active profile set, falling back to default profiles: default
2017-07-21 08:28:26.378  INFO 2550 --- [           main] s.c.a.AnnotationConfigApplicationContext : Refreshing org.springframework.context.annotation.AnnotationConfigApplicationContext@cdbe995: startup date [Fri Jul 21 08:28:26 UTC 2017]; parent: org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@50b8ae8d
2017-07-21 08:28:26.549  INFO 2550 --- [           main] f.a.AutowiredAnnotationBeanPostProcessor : JSR-330 'javax.inject.Inject' annotation found and supported for autowiring
2017-07-21 08:28:26.819  INFO 2550 --- [           main] o.s.c.support.GenericApplicationContext  : Refreshing org.springframework.context.support.GenericApplicationContext@141bb6b8: startup date [Fri Jul 21 08:28:26 UTC 2017]; root of context hierarchy
2017-07-21 08:28:26.998  INFO 2550 --- [           main] o.a.maven.surefire.booter.ForkedBooter   : Started ForkedBooter in 1.087 seconds (JVM running for 35.312)
2017-07-21 08:28:27.235  WARN 2550 --- [           main] o.s.amqp.rabbit.core.RabbitAdmin         : Failed to declare exchange: Exchange [name=springCloudBus, type=topic, durable=true, autoDelete=false, internal=false, arguments={}], continuing...

org.springframework.amqp.AmqpConnectException: java.net.ConnectException: 拒绝连接
	at org.springframework.amqp.rabbit.support.RabbitExceptionTranslator.convertRabbitAccessException(RabbitExceptionTranslator.java:62) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.connection.AbstractConnectionFactory.createBareConnection(AbstractConnectionFactory.java:309) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.connection.CachingConnectionFactory.createConnection(CachingConnectionFactory.java:547) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.doExecute(RabbitTemplate.java:1387) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.execute(RabbitTemplate.java:1368) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.execute(RabbitTemplate.java:1344) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitAdmin.declareExchange(RabbitAdmin.java:156) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.declareExchange(RabbitMessageChannelBinder.java:502) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.createProducerDestinationIfNecessary(RabbitMessageChannelBinder.java:338) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.createProducerDestinationIfNecessary(RabbitMessageChannelBinder.java:89) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractMessageChannelBinder.doBindProducer(AbstractMessageChannelBinder.java:101) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractMessageChannelBinder.doBindProducer(AbstractMessageChannelBinder.java:55) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractBinder.bindProducer(AbstractBinder.java:152) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.BindingService.bindProducer(BindingService.java:125) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.BindableProxyFactory.bindOutputs(BindableProxyFactory.java:238) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.OutputBindingLifecycle.start(OutputBindingLifecycle.java:57) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.context.support.DefaultLifecycleProcessor.doStart(DefaultLifecycleProcessor.java:173) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.access$200(DefaultLifecycleProcessor.java:51) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor$LifecycleGroup.start(DefaultLifecycleProcessor.java:346) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.startBeans(DefaultLifecycleProcessor.java:149) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.onRefresh(DefaultLifecycleProcessor.java:112) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.AbstractApplicationContext.finishRefresh(AbstractApplicationContext.java:874) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.boot.context.embedded.EmbeddedWebApplicationContext.finishRefresh(EmbeddedWebApplicationContext.java:144) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.context.support.AbstractApplicationContext.refresh(AbstractApplicationContext.java:544) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.boot.context.embedded.EmbeddedWebApplicationContext.refresh(EmbeddedWebApplicationContext.java:122) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.refresh(SpringApplication.java:761) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.refreshContext(SpringApplication.java:371) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.run(SpringApplication.java:315) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.test.SpringApplicationContextLoader.loadContext(SpringApplicationContextLoader.java:103) [spring-boot-test-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate.loadContextInternal(DefaultCacheAwareContextLoaderDelegate.java:98) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate.loadContext(DefaultCacheAwareContextLoaderDelegate.java:116) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DefaultTestContext.getApplicationContext(DefaultTestContext.java:83) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DependencyInjectionTestExecutionListener.injectDependencies(DependencyInjectionTestExecutionListener.java:117) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DependencyInjectionTestExecutionListener.prepareTestInstance(DependencyInjectionTestExecutionListener.java:83) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.TestContextManager.prepareTestInstance(TestContextManager.java:230) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.createTest(SpringJUnit4ClassRunner.java:228) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner$1.runReflectiveCall(SpringJUnit4ClassRunner.java:287) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.methodBlock(SpringJUnit4ClassRunner.java:289) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.runChild(SpringJUnit4ClassRunner.java:247) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.runChild(SpringJUnit4ClassRunner.java:94) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.runners.ParentRunner$3.run(ParentRunner.java:290) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:71) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner.runChildren(ParentRunner.java:288) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner.access$000(ParentRunner.java:58) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:268) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.statements.RunBeforeTestClassCallbacks.evaluate(RunBeforeTestClassCallbacks.java:61) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.statements.RunAfterTestClassCallbacks.evaluate(RunAfterTestClassCallbacks.java:70) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.runners.ParentRunner.run(ParentRunner.java:363) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.run(SpringJUnit4ClassRunner.java:191) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.apache.maven.surefire.junit4.JUnit4Provider.execute(JUnit4Provider.java:283) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.executeWithRerun(JUnit4Provider.java:173) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.executeTestSet(JUnit4Provider.java:153) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.invoke(JUnit4Provider.java:128) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.invokeProviderInSameClassLoader(ForkedBooter.java:203) [surefire-booter-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.runSuitesInProcess(ForkedBooter.java:155) [surefire-booter-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.main(ForkedBooter.java:103) [surefire-booter-2.18.1.jar:2.18.1]
Caused by: java.net.ConnectException: 拒绝连接
	at java.net.PlainSocketImpl.socketConnect(Native Method) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.doConnect(AbstractPlainSocketImpl.java:350) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.connectToAddress(AbstractPlainSocketImpl.java:206) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.connect(AbstractPlainSocketImpl.java:188) ~[na:1.8.0_91]
	at java.net.SocksSocketImpl.connect(SocksSocketImpl.java:392) ~[na:1.8.0_91]
	at java.net.Socket.connect(Socket.java:589) ~[na:1.8.0_91]
	at com.rabbitmq.client.impl.FrameHandlerFactory.create(FrameHandlerFactory.java:32) ~[amqp-client-3.6.3.jar:na]
	at com.rabbitmq.client.ConnectionFactory.newConnection(ConnectionFactory.java:811) ~[amqp-client-3.6.3.jar:na]
	at com.rabbitmq.client.ConnectionFactory.newConnection(ConnectionFactory.java:725) ~[amqp-client-3.6.3.jar:na]
	at org.springframework.amqp.rabbit.connection.AbstractConnectionFactory.createBareConnection(AbstractConnectionFactory.java:296) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	... 55 common frames omitted

2017-07-21 08:28:27.293  INFO 2550 --- [           main] o.s.integration.channel.DirectChannel    : Channel 'configserver:0.springCloudBusOutput' has 1 subscriber(s).
2017-07-21 08:28:27.299  INFO 2550 --- [           main] o.s.c.support.DefaultLifecycleProcessor  : Starting beans in phase 0
2017-07-21 08:28:27.300  INFO 2550 --- [           main] o.s.c.c.m.FileMonitorConfiguration       : Not monitoring for local config changes
2017-07-21 08:28:27.300  INFO 2550 --- [           main] o.s.i.endpoint.EventDrivenConsumer       : Adding {logging-channel-adapter:_org.springframework.integration.errorLogger} as a subscriber to the 'errorChannel' channel
2017-07-21 08:28:27.302  INFO 2550 --- [           main] o.s.i.channel.PublishSubscribeChannel    : Channel 'configserver:0.errorChannel' has 1 subscriber(s).
2017-07-21 08:28:27.302  INFO 2550 --- [           main] o.s.i.endpoint.EventDrivenConsumer       : started _org.springframework.integration.errorLogger
2017-07-21 08:28:27.303  INFO 2550 --- [           main] o.s.c.support.DefaultLifecycleProcessor  : Starting beans in phase 2147482647
2017-07-21 08:28:27.305 DEBUG 2550 --- [           main] o.s.c.s.binding.BindableProxyFactory     : Binding inputs for :interface org.springframework.cloud.bus.SpringCloudBusClient
2017-07-21 08:28:27.308 DEBUG 2550 --- [           main] o.s.c.s.binding.BindableProxyFactory     : Binding :interface org.springframework.cloud.bus.SpringCloudBusClient:springCloudBusInput
2017-07-21 08:28:27.361  INFO 2550 --- [           main] o.s.c.s.b.r.RabbitMessageChannelBinder   : declaring queue for inbound: springCloudBus.anonymous.5iPNmDvgQGacJq6OKWfL1g, bound to: springCloudBus
2017-07-21 08:28:27.364  WARN 2550 --- [           main] o.s.amqp.rabbit.core.RabbitAdmin         : Failed to declare exchange: Exchange [name=springCloudBus, type=topic, durable=true, autoDelete=false, internal=false, arguments={}], continuing...

org.springframework.amqp.AmqpConnectException: java.net.ConnectException: 拒绝连接
	at org.springframework.amqp.rabbit.support.RabbitExceptionTranslator.convertRabbitAccessException(RabbitExceptionTranslator.java:62) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.connection.AbstractConnectionFactory.createBareConnection(AbstractConnectionFactory.java:309) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.connection.CachingConnectionFactory.createConnection(CachingConnectionFactory.java:547) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.doExecute(RabbitTemplate.java:1387) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.execute(RabbitTemplate.java:1368) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.execute(RabbitTemplate.java:1344) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitAdmin.declareExchange(RabbitAdmin.java:156) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.declareExchange(RabbitMessageChannelBinder.java:502) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.createConsumerDestinationIfNecessary(RabbitMessageChannelBinder.java:271) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.createConsumerDestinationIfNecessary(RabbitMessageChannelBinder.java:89) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractMessageChannelBinder.doBindConsumer(AbstractMessageChannelBinder.java:195) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractMessageChannelBinder.doBindConsumer(AbstractMessageChannelBinder.java:55) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractBinder.bindConsumer(AbstractBinder.java:145) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.BindingService.bindConsumer(BindingService.java:98) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.BindableProxyFactory.bindInputs(BindableProxyFactory.java:221) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.InputBindingLifecycle.start(InputBindingLifecycle.java:55) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.context.support.DefaultLifecycleProcessor.doStart(DefaultLifecycleProcessor.java:173) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.access$200(DefaultLifecycleProcessor.java:51) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor$LifecycleGroup.start(DefaultLifecycleProcessor.java:346) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.startBeans(DefaultLifecycleProcessor.java:149) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.onRefresh(DefaultLifecycleProcessor.java:112) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.AbstractApplicationContext.finishRefresh(AbstractApplicationContext.java:874) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.boot.context.embedded.EmbeddedWebApplicationContext.finishRefresh(EmbeddedWebApplicationContext.java:144) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.context.support.AbstractApplicationContext.refresh(AbstractApplicationContext.java:544) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.boot.context.embedded.EmbeddedWebApplicationContext.refresh(EmbeddedWebApplicationContext.java:122) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.refresh(SpringApplication.java:761) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.refreshContext(SpringApplication.java:371) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.run(SpringApplication.java:315) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.test.SpringApplicationContextLoader.loadContext(SpringApplicationContextLoader.java:103) [spring-boot-test-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate.loadContextInternal(DefaultCacheAwareContextLoaderDelegate.java:98) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate.loadContext(DefaultCacheAwareContextLoaderDelegate.java:116) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DefaultTestContext.getApplicationContext(DefaultTestContext.java:83) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DependencyInjectionTestExecutionListener.injectDependencies(DependencyInjectionTestExecutionListener.java:117) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DependencyInjectionTestExecutionListener.prepareTestInstance(DependencyInjectionTestExecutionListener.java:83) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.TestContextManager.prepareTestInstance(TestContextManager.java:230) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.createTest(SpringJUnit4ClassRunner.java:228) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner$1.runReflectiveCall(SpringJUnit4ClassRunner.java:287) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.methodBlock(SpringJUnit4ClassRunner.java:289) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.runChild(SpringJUnit4ClassRunner.java:247) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.runChild(SpringJUnit4ClassRunner.java:94) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.runners.ParentRunner$3.run(ParentRunner.java:290) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:71) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner.runChildren(ParentRunner.java:288) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner.access$000(ParentRunner.java:58) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:268) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.statements.RunBeforeTestClassCallbacks.evaluate(RunBeforeTestClassCallbacks.java:61) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.statements.RunAfterTestClassCallbacks.evaluate(RunAfterTestClassCallbacks.java:70) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.runners.ParentRunner.run(ParentRunner.java:363) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.run(SpringJUnit4ClassRunner.java:191) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.apache.maven.surefire.junit4.JUnit4Provider.execute(JUnit4Provider.java:283) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.executeWithRerun(JUnit4Provider.java:173) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.executeTestSet(JUnit4Provider.java:153) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.invoke(JUnit4Provider.java:128) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.invokeProviderInSameClassLoader(ForkedBooter.java:203) [surefire-booter-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.runSuitesInProcess(ForkedBooter.java:155) [surefire-booter-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.main(ForkedBooter.java:103) [surefire-booter-2.18.1.jar:2.18.1]
Caused by: java.net.ConnectException: 拒绝连接
	at java.net.PlainSocketImpl.socketConnect(Native Method) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.doConnect(AbstractPlainSocketImpl.java:350) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.connectToAddress(AbstractPlainSocketImpl.java:206) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.connect(AbstractPlainSocketImpl.java:188) ~[na:1.8.0_91]
	at java.net.SocksSocketImpl.connect(SocksSocketImpl.java:392) ~[na:1.8.0_91]
	at java.net.Socket.connect(Socket.java:589) ~[na:1.8.0_91]
	at com.rabbitmq.client.impl.FrameHandlerFactory.create(FrameHandlerFactory.java:32) ~[amqp-client-3.6.3.jar:na]
	at com.rabbitmq.client.ConnectionFactory.newConnection(ConnectionFactory.java:811) ~[amqp-client-3.6.3.jar:na]
	at com.rabbitmq.client.ConnectionFactory.newConnection(ConnectionFactory.java:725) ~[amqp-client-3.6.3.jar:na]
	at org.springframework.amqp.rabbit.connection.AbstractConnectionFactory.createBareConnection(AbstractConnectionFactory.java:296) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	... 55 common frames omitted

2017-07-21 08:28:27.386  WARN 2550 --- [           main] o.s.amqp.rabbit.core.RabbitAdmin         : Failed to declare queue: Queue [name=springCloudBus.anonymous.5iPNmDvgQGacJq6OKWfL1g, durable=false, autoDelete=true, exclusive=true, arguments=null], continuing...

org.springframework.amqp.AmqpConnectException: java.net.ConnectException: 拒绝连接
	at org.springframework.amqp.rabbit.support.RabbitExceptionTranslator.convertRabbitAccessException(RabbitExceptionTranslator.java:62) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.connection.AbstractConnectionFactory.createBareConnection(AbstractConnectionFactory.java:309) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.connection.CachingConnectionFactory.createConnection(CachingConnectionFactory.java:547) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.doExecute(RabbitTemplate.java:1387) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.execute(RabbitTemplate.java:1368) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.execute(RabbitTemplate.java:1344) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitAdmin.declareQueue(RabbitAdmin.java:206) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.declareQueue(RabbitMessageChannelBinder.java:468) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.createConsumerDestinationIfNecessary(RabbitMessageChannelBinder.java:295) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.createConsumerDestinationIfNecessary(RabbitMessageChannelBinder.java:89) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractMessageChannelBinder.doBindConsumer(AbstractMessageChannelBinder.java:195) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractMessageChannelBinder.doBindConsumer(AbstractMessageChannelBinder.java:55) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractBinder.bindConsumer(AbstractBinder.java:145) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.BindingService.bindConsumer(BindingService.java:98) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.BindableProxyFactory.bindInputs(BindableProxyFactory.java:221) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.InputBindingLifecycle.start(InputBindingLifecycle.java:55) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.context.support.DefaultLifecycleProcessor.doStart(DefaultLifecycleProcessor.java:173) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.access$200(DefaultLifecycleProcessor.java:51) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor$LifecycleGroup.start(DefaultLifecycleProcessor.java:346) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.startBeans(DefaultLifecycleProcessor.java:149) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.onRefresh(DefaultLifecycleProcessor.java:112) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.AbstractApplicationContext.finishRefresh(AbstractApplicationContext.java:874) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.boot.context.embedded.EmbeddedWebApplicationContext.finishRefresh(EmbeddedWebApplicationContext.java:144) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.context.support.AbstractApplicationContext.refresh(AbstractApplicationContext.java:544) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.boot.context.embedded.EmbeddedWebApplicationContext.refresh(EmbeddedWebApplicationContext.java:122) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.refresh(SpringApplication.java:761) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.refreshContext(SpringApplication.java:371) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.run(SpringApplication.java:315) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.test.SpringApplicationContextLoader.loadContext(SpringApplicationContextLoader.java:103) [spring-boot-test-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate.loadContextInternal(DefaultCacheAwareContextLoaderDelegate.java:98) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate.loadContext(DefaultCacheAwareContextLoaderDelegate.java:116) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DefaultTestContext.getApplicationContext(DefaultTestContext.java:83) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DependencyInjectionTestExecutionListener.injectDependencies(DependencyInjectionTestExecutionListener.java:117) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DependencyInjectionTestExecutionListener.prepareTestInstance(DependencyInjectionTestExecutionListener.java:83) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.TestContextManager.prepareTestInstance(TestContextManager.java:230) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.createTest(SpringJUnit4ClassRunner.java:228) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner$1.runReflectiveCall(SpringJUnit4ClassRunner.java:287) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.methodBlock(SpringJUnit4ClassRunner.java:289) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.runChild(SpringJUnit4ClassRunner.java:247) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.runChild(SpringJUnit4ClassRunner.java:94) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.runners.ParentRunner$3.run(ParentRunner.java:290) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:71) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner.runChildren(ParentRunner.java:288) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner.access$000(ParentRunner.java:58) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:268) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.statements.RunBeforeTestClassCallbacks.evaluate(RunBeforeTestClassCallbacks.java:61) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.statements.RunAfterTestClassCallbacks.evaluate(RunAfterTestClassCallbacks.java:70) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.runners.ParentRunner.run(ParentRunner.java:363) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.run(SpringJUnit4ClassRunner.java:191) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.apache.maven.surefire.junit4.JUnit4Provider.execute(JUnit4Provider.java:283) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.executeWithRerun(JUnit4Provider.java:173) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.executeTestSet(JUnit4Provider.java:153) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.invoke(JUnit4Provider.java:128) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.invokeProviderInSameClassLoader(ForkedBooter.java:203) [surefire-booter-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.runSuitesInProcess(ForkedBooter.java:155) [surefire-booter-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.main(ForkedBooter.java:103) [surefire-booter-2.18.1.jar:2.18.1]
Caused by: java.net.ConnectException: 拒绝连接
	at java.net.PlainSocketImpl.socketConnect(Native Method) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.doConnect(AbstractPlainSocketImpl.java:350) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.connectToAddress(AbstractPlainSocketImpl.java:206) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.connect(AbstractPlainSocketImpl.java:188) ~[na:1.8.0_91]
	at java.net.SocksSocketImpl.connect(SocksSocketImpl.java:392) ~[na:1.8.0_91]
	at java.net.Socket.connect(Socket.java:589) ~[na:1.8.0_91]
	at com.rabbitmq.client.impl.FrameHandlerFactory.create(FrameHandlerFactory.java:32) ~[amqp-client-3.6.3.jar:na]
	at com.rabbitmq.client.ConnectionFactory.newConnection(ConnectionFactory.java:811) ~[amqp-client-3.6.3.jar:na]
	at com.rabbitmq.client.ConnectionFactory.newConnection(ConnectionFactory.java:725) ~[amqp-client-3.6.3.jar:na]
	at org.springframework.amqp.rabbit.connection.AbstractConnectionFactory.createBareConnection(AbstractConnectionFactory.java:296) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	... 55 common frames omitted

2017-07-21 08:28:27.419  WARN 2550 --- [           main] o.s.amqp.rabbit.core.RabbitAdmin         : Failed to declare binding: Binding [destination=springCloudBus.anonymous.5iPNmDvgQGacJq6OKWfL1g, exchange=springCloudBus, routingKey=#], continuing...

org.springframework.amqp.AmqpConnectException: java.net.ConnectException: 拒绝连接
	at org.springframework.amqp.rabbit.support.RabbitExceptionTranslator.convertRabbitAccessException(RabbitExceptionTranslator.java:62) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.connection.AbstractConnectionFactory.createBareConnection(AbstractConnectionFactory.java:309) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.connection.CachingConnectionFactory.createConnection(CachingConnectionFactory.java:547) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.doExecute(RabbitTemplate.java:1387) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.execute(RabbitTemplate.java:1368) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitTemplate.execute(RabbitTemplate.java:1344) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.amqp.rabbit.core.RabbitAdmin.declareBinding(RabbitAdmin.java:292) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.declareBinding(RabbitMessageChannelBinder.java:574) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.notPartitionedBinding(RabbitMessageChannelBinder.java:554) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.declareConsumerBindings(RabbitMessageChannelBinder.java:519) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.createConsumerDestinationIfNecessary(RabbitMessageChannelBinder.java:297) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.rabbit.RabbitMessageChannelBinder.createConsumerDestinationIfNecessary(RabbitMessageChannelBinder.java:89) [spring-cloud-stream-binder-rabbit-1.1.4.BUILD-SNAPSHOT.jar:1.1.4.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractMessageChannelBinder.doBindConsumer(AbstractMessageChannelBinder.java:195) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractMessageChannelBinder.doBindConsumer(AbstractMessageChannelBinder.java:55) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binder.AbstractBinder.bindConsumer(AbstractBinder.java:145) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.BindingService.bindConsumer(BindingService.java:98) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.BindableProxyFactory.bindInputs(BindableProxyFactory.java:221) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.cloud.stream.binding.InputBindingLifecycle.start(InputBindingLifecycle.java:55) [spring-cloud-stream-1.1.3.BUILD-SNAPSHOT.jar:1.1.3.BUILD-SNAPSHOT]
	at org.springframework.context.support.DefaultLifecycleProcessor.doStart(DefaultLifecycleProcessor.java:173) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.access$200(DefaultLifecycleProcessor.java:51) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor$LifecycleGroup.start(DefaultLifecycleProcessor.java:346) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.startBeans(DefaultLifecycleProcessor.java:149) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.DefaultLifecycleProcessor.onRefresh(DefaultLifecycleProcessor.java:112) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.context.support.AbstractApplicationContext.finishRefresh(AbstractApplicationContext.java:874) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.boot.context.embedded.EmbeddedWebApplicationContext.finishRefresh(EmbeddedWebApplicationContext.java:144) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.context.support.AbstractApplicationContext.refresh(AbstractApplicationContext.java:544) [spring-context-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.boot.context.embedded.EmbeddedWebApplicationContext.refresh(EmbeddedWebApplicationContext.java:122) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.refresh(SpringApplication.java:761) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.refreshContext(SpringApplication.java:371) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.SpringApplication.run(SpringApplication.java:315) [spring-boot-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.boot.test.SpringApplicationContextLoader.loadContext(SpringApplicationContextLoader.java:103) [spring-boot-test-1.4.1.BUILD-SNAPSHOT.jar:1.4.1.BUILD-SNAPSHOT]
	at org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate.loadContextInternal(DefaultCacheAwareContextLoaderDelegate.java:98) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.cache.DefaultCacheAwareContextLoaderDelegate.loadContext(DefaultCacheAwareContextLoaderDelegate.java:116) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DefaultTestContext.getApplicationContext(DefaultTestContext.java:83) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DependencyInjectionTestExecutionListener.injectDependencies(DependencyInjectionTestExecutionListener.java:117) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.support.DependencyInjectionTestExecutionListener.prepareTestInstance(DependencyInjectionTestExecutionListener.java:83) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.TestContextManager.prepareTestInstance(TestContextManager.java:230) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.createTest(SpringJUnit4ClassRunner.java:228) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner$1.runReflectiveCall(SpringJUnit4ClassRunner.java:287) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.methodBlock(SpringJUnit4ClassRunner.java:289) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.runChild(SpringJUnit4ClassRunner.java:247) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.runChild(SpringJUnit4ClassRunner.java:94) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.runners.ParentRunner$3.run(ParentRunner.java:290) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner$1.schedule(ParentRunner.java:71) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner.runChildren(ParentRunner.java:288) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner.access$000(ParentRunner.java:58) [junit-4.12.jar:4.12]
	at org.junit.runners.ParentRunner$2.evaluate(ParentRunner.java:268) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.statements.RunBeforeTestClassCallbacks.evaluate(RunBeforeTestClassCallbacks.java:61) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.springframework.test.context.junit4.statements.RunAfterTestClassCallbacks.evaluate(RunAfterTestClassCallbacks.java:70) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.junit.runners.ParentRunner.run(ParentRunner.java:363) [junit-4.12.jar:4.12]
	at org.springframework.test.context.junit4.SpringJUnit4ClassRunner.run(SpringJUnit4ClassRunner.java:191) [spring-test-4.3.3.RELEASE.jar:4.3.3.RELEASE]
	at org.apache.maven.surefire.junit4.JUnit4Provider.execute(JUnit4Provider.java:283) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.executeWithRerun(JUnit4Provider.java:173) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.executeTestSet(JUnit4Provider.java:153) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.junit4.JUnit4Provider.invoke(JUnit4Provider.java:128) [surefire-junit4-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.invokeProviderInSameClassLoader(ForkedBooter.java:203) [surefire-booter-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.runSuitesInProcess(ForkedBooter.java:155) [surefire-booter-2.18.1.jar:2.18.1]
	at org.apache.maven.surefire.booter.ForkedBooter.main(ForkedBooter.java:103) [surefire-booter-2.18.1.jar:2.18.1]
Caused by: java.net.ConnectException: 拒绝连接
	at java.net.PlainSocketImpl.socketConnect(Native Method) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.doConnect(AbstractPlainSocketImpl.java:350) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.connectToAddress(AbstractPlainSocketImpl.java:206) ~[na:1.8.0_91]
	at java.net.AbstractPlainSocketImpl.connect(AbstractPlainSocketImpl.java:188) ~[na:1.8.0_91]
	at java.net.SocksSocketImpl.connect(SocksSocketImpl.java:392) ~[na:1.8.0_91]
	at java.net.Socket.connect(Socket.java:589) ~[na:1.8.0_91]
	at com.rabbitmq.client.impl.FrameHandlerFactory.create(FrameHandlerFactory.java:32) ~[amqp-client-3.6.3.jar:na]
	at com.rabbitmq.client.ConnectionFactory.newConnection(ConnectionFactory.java:811) ~[amqp-client-3.6.3.jar:na]
	at com.rabbitmq.client.ConnectionFactory.newConnection(ConnectionFactory.java:725) ~[amqp-client-3.6.3.jar:na]
	at org.springframework.amqp.rabbit.connection.AbstractConnectionFactory.createBareConnection(AbstractConnectionFactory.java:296) ~[spring-rabbit-1.6.2.RELEASE.jar:na]
	... 57 common frames omitted

2017-07-21 08:28:27.626  INFO 2550 --- [           main] o.s.i.a.i.AmqpInboundChannelAdapter      : started inbound.springCloudBus.anonymous.t8fidQ7rSj6YKAyteDKDcA
2017-07-21 08:28:27.627  INFO 2550 --- [           main] o.s.i.endpoint.EventDrivenConsumer       : Adding {message-handler:inbound.springCloudBus.default} as a subscriber to the 'bridge.springCloudBus' channel
2017-07-21 08:28:27.627  INFO 2550 --- [           main] o.s.i.endpoint.EventDrivenConsumer       : started inbound.springCloudBus.default
2017-07-21 08:28:27.637  INFO 2550 --- [           main] o.s.c.support.DefaultLifecycleProcessor  : Starting beans in phase 2147483647
2017-07-21 08:28:27.869  INFO 2550 --- [           main] s.b.c.e.t.TomcatEmbeddedServletContainer : Tomcat started on port(s): 38976 (http)
2017-07-21 08:28:27.871  INFO 2550 --- [           main] c.n.e.EurekaDiscoveryClientConfiguration : Updating port to 38976
2017-07-21 08:28:27.904  INFO 2550 --- [           main] o.s.c.n.eureka.InstanceInfoFactory       : Setting initial instance status as: STARTING
2017-07-21 08:28:30.873  INFO 2550 --- [           main] c.n.e.EurekaDiscoveryClientConfiguration : Registering application configserver with eureka with status UP
2017-07-21 08:28:30.882 DEBUG 2550 --- [           main] s.c.c.d.h.DiscoveryClientHealthIndicator : Discovery Client has been initialized
2017-07-21 08:28:30.890  INFO 2550 --- [           main] demo.ApplicationTests                    : Started ApplicationTests in 36.735 seconds (JVM running for 39.204)
2017-07-21 08:28:31.763  INFO 2550 --- [o-auto-1-exec-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring FrameworkServlet 'dispatcherServlet'
2017-07-21 08:28:31.764  INFO 2550 --- [o-auto-1-exec-1] o.s.web.servlet.DispatcherServlet        : FrameworkServlet 'dispatcherServlet': initialization started
2017-07-21 08:28:31.851  INFO 2550 --- [o-auto-1-exec-1] o.s.web.servlet.DispatcherServlet        : FrameworkServlet 'dispatcherServlet': initialization completed in 85 ms
2017-07-21 08:28:32.640  WARN 2550 --- [GacJq6OKWfL1g-1] o.s.a.r.l.SimpleMessageListenerContainer : Consumer raised exception, processing can restart if the connection factory supports it. Exception summary: org.springframework.amqp.AmqpConnectException: java.net.ConnectException: 拒绝连接
2017-07-21 08:28:32.640  INFO 2550 --- [GacJq6OKWfL1g-1] o.s.a.r.l.SimpleMessageListenerContainer : Restarting Consumer: tags=[{}], channel=null, acknowledgeMode=AUTO local queue size=0
2017-07-21 08:28:37.663  WARN 2550 --- [GacJq6OKWfL1g-2] o.s.a.r.l.SimpleMessageListenerContainer : Consumer raised exception, processing can restart if the connection factory supports it. Exception summary: org.springframework.amqp.AmqpConnectException: java.net.ConnectException: 拒绝连接
2017-07-21 08:28:37.664  INFO 2550 --- [GacJq6OKWfL1g-2] o.s.a.r.l.SimpleMessageListenerContainer : Restarting Consumer: tags=[{}], channel=null, acknowledgeMode=AUTO local queue size=0
2017-07-21 08:28:41.378 DEBUG 2550 --- [o-auto-1-exec-1] o.s.c.n.m.ServoEnvironmentPostProcessor  : Setting 'spring.aop.proxyTargetClass=true' to make spring AOP default to target class so RestTemplates can be customized
2017-07-21 08:28:41.650  INFO 2550 --- [o-auto-1-exec-1] o.s.boot.SpringApplication               : Starting application on localhost.localdomain with PID 2550 (started by vagrant in /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver)
2017-07-21 08:28:41.650  INFO 2550 --- [o-auto-1-exec-1] o.s.boot.SpringApplication               : The following profiles are active: cloud
2017-07-21 08:28:41.657  INFO 2550 --- [o-auto-1-exec-1] s.c.a.AnnotationConfigApplicationContext : Refreshing org.springframework.context.annotation.AnnotationConfigApplicationContext@2d8eedaa: startup date [Fri Jul 21 08:28:41 UTC 2017]; root of context hierarchy
2017-07-21 08:28:41.690  INFO 2550 --- [o-auto-1-exec-1] f.a.AutowiredAnnotationBeanPostProcessor : JSR-330 'javax.inject.Inject' annotation found and supported for autowiring
2017-07-21 08:28:41.745  INFO 2550 --- [o-auto-1-exec-1] o.s.boot.SpringApplication               : Started application in 1.022 seconds (JVM running for 50.059)
2017-07-21 08:28:41.753  INFO 2550 --- [o-auto-1-exec-1] o.s.c.c.s.e.NativeEnvironmentRepository  : Adding property source: file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/target/config/application.yml#cloud
2017-07-21 08:28:41.753 DEBUG 2550 --- [o-auto-1-exec-1] o.s.c.c.s.e.NativeEnvironmentRepository  : Not adding property source: classpath:/application.yml#cloud
2017-07-21 08:28:41.754  INFO 2550 --- [o-auto-1-exec-1] o.s.c.c.s.e.NativeEnvironmentRepository  : Adding property source: file:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/target/config/application.yml
2017-07-21 08:28:41.755  INFO 2550 --- [o-auto-1-exec-1] s.c.a.AnnotationConfigApplicationContext : Closing org.springframework.context.annotation.AnnotationConfigApplicationContext@2d8eedaa: startup date [Fri Jul 21 08:28:41 UTC 2017]; root of context hierarchy
Tests run: 2, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 48.62 sec - in demo.ApplicationTests
2017-07-21 08:28:42.264  INFO 2550 --- [       Thread-7] ationConfigEmbeddedWebApplicationContext : Closing org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@50b8ae8d: startup date [Fri Jul 21 08:28:01 UTC 2017]; parent: org.springframework.context.annotation.AnnotationConfigApplicationContext@146dfe6
2017-07-21 08:28:42.267  INFO 2550 --- [       Thread-7] c.n.e.EurekaDiscoveryClientConfiguration : Unregistering application configserver with eureka with status DOWN
2017-07-21 08:28:42.296  INFO 2550 --- [       Thread-7] s.c.a.AnnotationConfigApplicationContext : Closing org.springframework.context.annotation.AnnotationConfigApplicationContext@cdbe995: startup date [Fri Jul 21 08:28:26 UTC 2017]; parent: org.springframework.boot.context.embedded.AnnotationConfigEmbeddedWebApplicationContext@50b8ae8d
2017-07-21 08:28:42.305  INFO 2550 --- [       Thread-7] c.n.e.EurekaDiscoveryClientConfiguration : Unregistering application configserver with eureka with status DOWN
2017-07-21 08:28:42.309  INFO 2550 --- [       Thread-7] o.s.c.support.DefaultLifecycleProcessor  : Stopping beans in phase 2147483647
2017-07-21 08:28:42.310  INFO 2550 --- [       Thread-7] o.s.c.support.DefaultLifecycleProcessor  : Stopping beans in phase 2147482647
2017-07-21 08:28:42.315 DEBUG 2550 --- [       Thread-7] o.s.c.s.binding.BindableProxyFactory     : Unbinding inputs for :interface org.springframework.cloud.bus.SpringCloudBusClient
2017-07-21 08:28:42.315 DEBUG 2550 --- [       Thread-7] o.s.c.s.binding.BindableProxyFactory     : Unbinding :interface org.springframework.cloud.bus.SpringCloudBusClient:springCloudBusInput
2017-07-21 08:28:42.316  INFO 2550 --- [       Thread-7] o.s.a.r.l.SimpleMessageListenerContainer : Waiting for workers to finish.
2017-07-21 08:28:42.316  INFO 2550 --- [       Thread-7] o.s.a.r.l.SimpleMessageListenerContainer : Successfully waited for workers to finish.
2017-07-21 08:28:42.317  INFO 2550 --- [       Thread-7] o.s.i.a.i.AmqpInboundChannelAdapter      : stopped inbound.springCloudBus.anonymous.t8fidQ7rSj6YKAyteDKDcA
2017-07-21 08:28:42.320  INFO 2550 --- [       Thread-7] o.s.c.support.DefaultLifecycleProcessor  : Stopping beans in phase 0
2017-07-21 08:28:42.321  INFO 2550 --- [       Thread-7] o.s.i.endpoint.EventDrivenConsumer       : Removing {logging-channel-adapter:_org.springframework.integration.errorLogger} as a subscriber to the 'errorChannel' channel
2017-07-21 08:28:42.321  INFO 2550 --- [       Thread-7] o.s.i.channel.PublishSubscribeChannel    : Channel 'configserver:0.errorChannel' has 0 subscriber(s).
2017-07-21 08:28:42.322  INFO 2550 --- [       Thread-7] o.s.i.endpoint.EventDrivenConsumer       : stopped _org.springframework.integration.errorLogger
2017-07-21 08:28:42.329  INFO 2550 --- [       Thread-7] o.s.c.support.DefaultLifecycleProcessor  : Stopping beans in phase -2147482648
2017-07-21 08:28:42.330 DEBUG 2550 --- [       Thread-7] o.s.c.s.binding.BindableProxyFactory     : Unbinding outputs for :interface org.springframework.cloud.bus.SpringCloudBusClient
2017-07-21 08:28:42.330 DEBUG 2550 --- [       Thread-7] o.s.c.s.binding.BindableProxyFactory     : Binding :interface org.springframework.cloud.bus.SpringCloudBusClient:springCloudBusOutput
2017-07-21 08:28:42.342  INFO 2550 --- [       Thread-7] o.s.s.c.ThreadPoolTaskScheduler          : Shutting down ExecutorService 'taskScheduler'

Results :

Tests run: 2, Failures: 0, Errors: 0, Skipped: 0

[INFO] 
[INFO] --- maven-jar-plugin:2.6:jar (default-jar) @ configserver ---
[INFO] Building jar: /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/target/configserver-0.0.1-SNAPSHOT.jar
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.4.1.BUILD-SNAPSHOT:repackage (default) @ configserver ---
Downloading: http://172.17.4.50:8081/content/repositories/allsynced/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/maven-metadata.xml
Downloading: https://repo.spring.io/libs-snapshot-local/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/maven-metadata.xml
Downloading: http://172.17.4.50:8081/content/repositories/snapshots/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/maven-metadata.xml
Downloaded: https://repo.spring.io/libs-snapshot-local/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/maven-metadata.xml (3 KB at 2.8 KB/sec)
Downloading: http://172.17.4.50:8081/content/repositories/allsynced/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/spring-boot-loader-tools-1.4.1.BUILD-20160920.220833-116.pom
Downloading: http://172.17.4.50:8081/content/repositories/snapshots/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/spring-boot-loader-tools-1.4.1.BUILD-20160920.220833-116.pom
Downloading: https://repo.spring.io/libs-snapshot-local/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/spring-boot-loader-tools-1.4.1.BUILD-20160920.220833-116.pom
Downloaded: https://repo.spring.io/libs-snapshot-local/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/spring-boot-loader-tools-1.4.1.BUILD-20160920.220833-116.pom (5 KB at 5.9 KB/sec)
Downloading: https://repo.spring.io/libs-snapshot-local/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/spring-boot-loader-tools-1.4.1.BUILD-20160920.220833-116.jar
Downloaded: https://repo.spring.io/libs-snapshot-local/org/springframework/boot/spring-boot-loader-tools/1.4.1.BUILD-SNAPSHOT/spring-boot-loader-tools-1.4.1.BUILD-20160920.220833-116.jar (140 KB at 110.2 KB/sec)
[INFO] 
[INFO] --- maven-install-plugin:2.5.2:install (default-install) @ configserver ---
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/target/configserver-0.0.1-SNAPSHOT.jar to /home/vagrant/.m2/repository/org/demo/configserver/0.0.1-SNAPSHOT/configserver-0.0.1-SNAPSHOT.jar
[INFO] Installing /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/spring-cloud-netflix/configserver/pom.xml to /home/vagrant/.m2/repository/org/demo/configserver/0.0.1-SNAPSHOT/configserver-0.0.1-SNAPSHOT.pom
[INFO] 
[INFO] --- spring-boot-maven-plugin:1.4.1.BUILD-SNAPSHOT:repackage (default-cli) @ configserver ---
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 03:46 min
[INFO] Finished at: 2017-07-21T08:28:51+00:00
[INFO] Final Memory: 43M/266M
[INFO] ------------------------------------------------------------------------
```