# 基于Maven repository的Java开发环境

## Table of content

* Maven及本地仓库
* Sonatype Nexus创建私有项目仓库
* Gradle
* 示例项目

## Maven及本地仓库

### JDK安装

* 使用Linux发行的OpenJDK，通过YUM（Redhat，CentOS，Fedora，OracleLinux等系列）安装
* 使用Oracle JDK安装, [参考链接](../get-oracle-java-intro.md)

### Maven安装

* [参考链接](../get-maven-and-gradle.md)

### 默认仓库位置

__本地仓库位置__

Windows下在*\\用户\\你\\.m2*
```
[vagrant@localhost go-to-openstack-bootcamp]$ ls ~/.m2/repository/
antlr                     commons-beanutils      commons-pool                                                        junit           ring
aopalliance               commons-chain          commons-validator                                                   kr              ring-cors
archetype-catalog.xml     commons-cli            compojure                                                           log4j           sslext
asm                       commons-codec          de                                                                  logkit          stax
avalon-framework          commons-collections    dk                                                                  mysql           uk
backport-util-concurrent  commons-configuration  dnsjava                                                             nekohtml        wsdl4j
biz                       commons-daemon         dom4j                                                               net             xalan
bouncycastle              commons-dbcp           github-dot-com                                                      nz              xerces
bsh                       commons-dbutils        hiccup                                                              ognl            xml-apis
cglib                     commons-digester       https0x3A0x2F0x2Fgithub0x2Ecom0x2Ftangfeixiong0x2Fgo-to-kubernetes  org             xmlenc
ch                        commons-fileupload     io                                                                  oro             xmlpull
classworlds               commons-httpclient     jakarta-regexp                                                      pl              xml-resolver
clj-stacktrace            commons-io             javax                                                               plexus          xpp3
clj-time                  commons-jxpath         jdom                                                                redis
clout                     commons-lang           jline                                                               regexp
co                        commons-logging        joda-time                                                           repository.xml
com                       commons-net            jtidy                                                               rhino
```

__默认远程仓库位置__

https://repo1.maven.org/maven2/

* 使用客户端工具mvn或eclipse m2插件编译时，mvn首先检查远程仓库的artifact的hash是否与本地一致，否则downloading
* 本地没有的到pom中指定的远程仓库 to download
* 指定的远程仓库中没有的，到默认远程仓库 to download

有关默认仓库，请参考SuperPOM https://maven.apache.org/guides/introduction/introduction-to-the-pom.html#Super_POM

__设置多个远程仓库__

参考 https://maven.apache.org/guides/mini/guide-multiple-repositories.html

一些远程仓库

```
    <repositories>
	    <repository>
		    <id>confluent</id>
			<name>confluent</name>
			<url>http://packages.confluent.io/maven/</url>
	    </repository>
		<repository>
			<id>twitter4j</id>
			<url>http://twitter4j.org/maven2</url>
		</repository>
        <repository>
            <id>clojars</id>
            <name>clojars</name>
            <url>http://clojars.org/repo/</url>
        </repository>
	    <repository>
		    <id>spring-libs-release</id>
			<name>spring-libs-release</name>
			<url>https://repo.spring.io/libs-release</url>
	    </repository>
	    <repository>
		    <id>spring-libs-milestone</id>
			<name>spring-libs-milestone</name>
			<url>https://repo.spring.io/libs-milestone</url>
	    </repository>
        <repository>
		    <id>spring-plugins-release</id>
			<name>spring-plugins-release</name>
			<url>https://repo.spring.io/plugins-release</url>
	    </repository>
        <repository>
		    <id>spring-plugins-snapshot</id>
			<name>spring-plugins-snapshot</name>
			<url>https://repo.spring.io/plugins-snapshot</url>
	    </repository>
		<repository>
		    <id>sonatype-snapshots</id>
			<name>sonatype-snapshots</name>
			<url>https://oss.sonatype.org/content/repositories/snapshots/</url>
		</repository>	
    </repositories>
```

## Sonatype Nexus创建私有项目仓库

Sonatype nexus是基于[Maven的repository manager规范](https://maven.apache.org/repository-management.html)的私有仓库管理器

### 安装

[参考链接](./docker-run.md)

### 配置项目仓库

如图，将`./m2/repository`中的内容作为项目仓库

[屏幕快照 2017-06-08 上午11.46.09.png](./屏幕快照 2017-06-08 上午11.46.09.png)

### 离线环境下复制artifacts到项目仓库

__查artifacts__

https://mvnrepository.com/

__在上网机器上__

用mvn从互联网下载，如：
```
fanhonglingdeMacBook-Pro:99-mirror fanhongling$ mvn dependency:get -DremoteRepositories=https://repo1.maven.org/maven2/ -Dartifact=io.fabric8.ipaas.platform.packages:ipaas-platform:1.0.32:yml:kubernetes
```

然后从该机的~/.m2/repository/...复制到内网Nexus的相同位置

__或者复制整个remote repo__

到 https://repo1.maven.org/maven2/ 下载

aliyun

http://maven.aliyun.com/nexus/content/groups/public/

其它

Spring下的maven repositories

https://repo.spring.io

## Gradle

Gadle可使用Maven远程和本地仓库

### 安装

* [参考链接](../get-maven-and-gradle.md)

## 示例项目

https://github.com/tangfeixiong/osev3-exampls/tree/cicd/spring-boot/sample-microservices-springboot

创建新分支
```
[vagrant@localhost sample-microservices-springboot]$ git checkout -b cicd
```

### 使用mvn

__settings.xml__

Copy
```
[vagrant@localhost go-to-openstack-bootcamp]$ cp /opt/apache-maven-3.3.9/conf/settings.xml ~/.m2/
```

Add followings
```
    <mirrors>
        <mirror>
            <id>nexus</id>
            <mirrorOf>*</mirrorOf>
            <url>http://172.17.4.50:8081/content/groups/public</url>
        </mirror>
    </mirrors>
    
    <servers>
        <server>
            <id>my-snapshots</id>
            <username>admin</username>
            <password>admin123</password>
        </server>
        <server>
            <id>my-releases</id>
            <username>admin</username>
            <password>admin123</password>
        </server>
    </servers>

    <profiles>
        <profile>
            <id>staging</id>
            <repositories>
                <repository>
                    <id>nexusLocalSync</id>
                    <url>http://172.17.4.50:8081/content/repositories/allsynced/</url>
                    <releases>
                        <enabled>true</enabled>
                    </releases>
                    <snapshots>
                        <enabled>true</enabled>
                    </snapshots>
                </repository>
                <repository>
                    <id>nexusLocalRepo</id>
                    <url>http://172.17.4.50:8081/content/repositories/releases/</url>
                    <releases>
                        <enabled>true</enabled>
                    </releases>
                    <snapshots>
                        <enabled>true</enabled>
                    </snapshots>
                </repository>
            </repositories>

            <pluginRepositories>
                <pluginRepository>
                    <id>nexusLocalSync</id>
                    <url>http://172.17.4.50:8081/content/repositories/allsynced/</url>
                    <releases>
                        <enabled>true</enabled>
                    </releases>
                    <snapshots>
                        <enabled>true</enabled>
                    </snapshots>
                </pluginRepository>
                <pluginRepository>
                    <id>nexusLocalRepo</id>
                    <url>http://172.17.4.50:8081/content/repositories/snapshots/</url>
                    <releases>
                        <enabled>true</enabled>
                    </releases>
                    <snapshots>
                        <enabled>true</enabled>
                    </snapshots>
                </pluginRepository>
            </pluginRepositories>
        </profile>
    </profiles>

    <activeProfiles>
        <activeProfile>staging</activeProfile>
    </activeProfiles>
```

__Deployment with the Nexus Staging Maven Plugin__

Refer to https://books.sonatype.com/nexus-book/reference/staging-deployment.html

Add followings in project pom.xml
```
    <repositories>
        <repository>
            <id>my-synced</id>
            <url>http://172.17.4.50:8081/content/repositories/allsynced</url>
            <releases>
			    <enabled>true</enabled>
		    </releases>
            <snapshots>
			    <enabled>true</enabled>
			</snapshots>
        </repository>
        <repository>
            <id>my-mirror</id>
            <url>http://172.17.4.50:8081/content/groups/public</url>
		</repository>
    </repositories>

    <pluginRepositories>
        <pluginRepository>
            <id>my-synced</id>
            <url>http://172.17.4.50:8081/content/repositories/allsynced</url>
            <releases>
			    <enabled>true</enabled>
			</releases>
            <snapshots>
			    <enabled>true</enabled>
			</snapshots>
        </pluginRepository>
        <pluginRepository>
            <id>my-mirror</id>
            <url>http://172.17.4.50:8081/content/groups/public</url>
		</pluginRepository>
    </pluginRepositories>
	
	<distributionManagement>
	    <snapshotRepository>
	        <id>my-snapshots</id>
	        <name>My internal repository</name>
	        <url>http://172.17.4.50:8081/content/repositories/snapshots</url>
	    </snapshotRepository>
	
	    <repository>
	        <id>my-releases</id>
	        <name>My internal repository</name>
	        <url>http://172.17.4.50:8081/content/repositories/releases</url>
	    </repository>
	</distributionManagement>
	
	<build>
	  <plugins>
		<plugin>
		   <artifactId>maven-deploy-plugin</artifactId>
		   <version>2.8.2</version>
		   <executions>
		      <execution>
		         <id>default-deploy</id>
		         <phase>deploy</phase>
		         <goals>
		            <goal>deploy</goal>
		         </goals>
		      </execution>
		   </executions>
	       <configuration>
	         <skip>true</skip>
	       </configuration>
		</plugin>	
	
	    <plugin>
	      <groupId>org.sonatype.plugins</groupId>
	      <artifactId>nexus-staging-maven-plugin</artifactId>
		  <version>1.6.8</version>
	      <executions>
	        <execution>
	          <id>default-deploy</id>
	          <phase>deploy</phase>
	          <goals>
	            <goal>deploy</goal>
	          </goals>
	        </execution>
	      </executions>
	      <configuration>
	        <serverId>nexus</serverId>
	        <nexusUrl>http://localhost:8081/nexus/</nexusUrl>
	        <!-- explicit matching using the staging profile id -->
	        <stagingProfileId>staging</stagingProfileId>
	      </configuration>
	    </plugin>
	  </plugins>
	</build>
```

Run
```
[vagrant@localhost sample-microservices-springboot]$ mvn clean deploy
[INFO] Scanning for projects...
Downloading: http://172.17.4.50:8081/content/groups/public/org/springframework/spring-framework-bom/4.1.7.RELEASE/spring-framework-bom-4.1.7.RELEASE.pom
Downloaded: http://172.17.4.50:8081/content/groups/public/org/springframework/spring-framework-bom/4.1.7.RELEASE/spring-framework-bom-4.1.7.RELEASE.pom (0 B at 0.0 KB/sec)
Downloading: ...
...snipped...
[INFO] Deploying remotely...
[INFO] Bulk deploying locally gathered artifacts from directory: 
[INFO]  * Bulk deploying locally gathered snapshot artifacts
Downloading: http://172.17.4.50:8081/content/repositories/snapshots/com/openshift/evangelists/microservices-demo-springboot/0.0.1-SNAPSHOT/maven-metadata.xml
Uploading: http://172.17.4.50:8081/content/repositories/snapshots/com/openshift/evangelists/microservices-demo-springboot/0.0.1-SNAPSHOT/microservices-demo-springboot-0.0.1-20170627.171950-1.pom
Uploaded: http://172.17.4.50:8081/content/repositories/snapshots/com/openshift/evangelists/microservices-demo-springboot/0.0.1-SNAPSHOT/microservices-demo-springboot-0.0.1-20170627.171950-1.pom (4 KB at 35.1 KB/sec)
Downloading: http://172.17.4.50:8081/content/repositories/snapshots/com/openshift/evangelists/microservices-demo-springboot/maven-metadata.xml
Uploading: http://172.17.4.50:8081/content/repositories/snapshots/com/openshift/evangelists/microservices-demo-springboot/0.0.1-SNAPSHOT/maven-metadata.xml
Uploaded: http://172.17.4.50:8081/content/repositories/snapshots/com/openshift/evangelists/microservices-demo-springboot/0.0.1-SNAPSHOT/maven-metadata.xml (631 B at 16.2 KB/sec)
Uploading: http://172.17.4.50:8081/content/repositories/snapshots/com/openshift/evangelists/microservices-demo-springboot/maven-metadata.xml
Uploaded: http://172.17.4.50:8081/content/repositories/snapshots/com/openshift/evangelists/microservices-demo-springboot/maven-metadata.xml (314 B at 8.8 KB/sec)
[INFO]  * Bulk deploy of locally gathered snapshot artifacts finished.
[INFO] Remote deploy finished with success.
[INFO] ------------------------------------------------------------------------
[INFO] Reactor Summary:
[INFO] 
[INFO] web ................................................ SUCCESS [ 14.749 s]
[INFO] repositories-mem ................................... SUCCESS [  7.673 s]
[INFO] microservices-demo ................................. SUCCESS [  2.800 s]
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 25.966 s
[INFO] Finished at: 2017-06-27T17:19:51+00:00
[INFO] Final Memory: 31M/226M
[INFO] ------------------------------------------------------------------------
```

![屏幕快照 2017-06-27 上午11.47.40.png](./屏幕快照%202017-06-27%20上午11.47.40.png)

### 使用gradle

wrapper
```
[vagrant@localhost sample-microservices-springboot]$ gradle wrapper --gradle-version 3.5
:wrapper

BUILD SUCCESSFUL

Total time: 6.283 secs

This build could be faster, please consider using the Gradle Daemon: http://gradle.org/docs/2.5/userguide/gradle_daemon.html
```
