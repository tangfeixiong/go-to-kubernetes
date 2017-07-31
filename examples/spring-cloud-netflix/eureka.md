### Spring Cloud Netflix - Eureka

project
```
[vagrant@localhost github.com] git clone --depth=1 https://github.com/spring-cloud-samples/eureka spring-cloud-samples/eureka
```

Trouble, in Linux
```
[vagrant@localhost eureka]$ mvn install spring-boot:repackage docker:build -Dmaven.gitcommitid.skip=true -Dmaven.gitcommitid.useNativeGit=true
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Eureka Server 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- git-commit-id-plugin:2.1.11:revision (default) @ eureka ---
[info] dotGitDirectory /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/eureka/.git
[info] git.build.user.name tangfeixiong
[info] git.build.user.email tangfx128@gmail.com
[info] git.branch master
[info] Tag refs [ [] ]
[info] Created map: [ {} ] 
[info] HEAD is [ ff513fe311ba3ce262027f40238a66ed2a164a6a ] 
[info] Repo is in dirty state [ false ]
java.lang.RuntimeException: Unable to find commits until some tag
	at pl.project13.jgit.DescribeCommand.findCommitsUntilSomeTag(DescribeCommand.java:414)
	at pl.project13.jgit.DescribeCommand.call(DescribeCommand.java:306)
	at pl.project13.maven.git.JGitProvider.getGitDescribe(JGitProvider.java:225)
	at pl.project13.maven.git.JGitProvider.getGitDescribe(JGitProvider.java:122)
	at pl.project13.maven.git.GitDataProvider.maybePutGitDescribe(GitDataProvider.java:93)
	at pl.project13.maven.git.GitDataProvider.loadGitData(GitDataProvider.java:63)
	at pl.project13.maven.git.GitCommitIdMojo.loadGitDataWithJGit(GitCommitIdMojo.java:485)
	at pl.project13.maven.git.GitCommitIdMojo.loadGitData(GitCommitIdMojo.java:457)
	at pl.project13.maven.git.GitCommitIdMojo.execute(GitCommitIdMojo.java:315)
	at org.apache.maven.plugin.DefaultBuildPluginManager.executeMojo(DefaultBuildPluginManager.java:134)
	at org.apache.maven.lifecycle.internal.MojoExecutor.execute(MojoExecutor.java:208)
	at org.apache.maven.lifecycle.internal.MojoExecutor.execute(MojoExecutor.java:153)
	at org.apache.maven.lifecycle.internal.MojoExecutor.execute(MojoExecutor.java:145)
	at org.apache.maven.lifecycle.internal.LifecycleModuleBuilder.buildProject(LifecycleModuleBuilder.java:116)
	at org.apache.maven.lifecycle.internal.LifecycleModuleBuilder.buildProject(LifecycleModuleBuilder.java:80)
	at org.apache.maven.lifecycle.internal.builder.singlethreaded.SingleThreadedBuilder.build(SingleThreadedBuilder.java:51)
	at org.apache.maven.lifecycle.internal.LifecycleStarter.execute(LifecycleStarter.java:128)
	at org.apache.maven.DefaultMaven.doExecute(DefaultMaven.java:307)
	at org.apache.maven.DefaultMaven.doExecute(DefaultMaven.java:193)
	at org.apache.maven.DefaultMaven.execute(DefaultMaven.java:106)
	at org.apache.maven.cli.MavenCli.execute(MavenCli.java:862)
	at org.apache.maven.cli.MavenCli.doMain(MavenCli.java:286)
	at org.apache.maven.cli.MavenCli.main(MavenCli.java:197)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:498)
	at org.codehaus.plexus.classworlds.launcher.Launcher.launchEnhanced(Launcher.java:289)
	at org.codehaus.plexus.classworlds.launcher.Launcher.launch(Launcher.java:229)
	at org.codehaus.plexus.classworlds.launcher.Launcher.mainWithExitCode(Launcher.java:415)
	at org.codehaus.plexus.classworlds.launcher.Launcher.main(Launcher.java:356)
Caused by: org.eclipse.jgit.errors.RevWalkException: Walk failure.
	at org.eclipse.jgit.revwalk.RevWalk.iterator(RevWalk.java:1259)
	at pl.project13.jgit.DescribeCommand.findCommitsUntilSomeTag(DescribeCommand.java:402)
	... 30 more
Caused by: org.eclipse.jgit.errors.MissingObjectException: Missing commit f81c16d587a532de071256f449255e5be21ea8a2
	at org.eclipse.jgit.internal.storage.file.WindowCursor.open(WindowCursor.java:149)
	at org.eclipse.jgit.revwalk.RevWalk.getCachedBytes(RevWalk.java:883)
	at org.eclipse.jgit.revwalk.RevCommit.parseHeaders(RevCommit.java:145)
	at org.eclipse.jgit.revwalk.PendingGenerator.next(PendingGenerator.java:148)
	at org.eclipse.jgit.revwalk.StartGenerator.next(StartGenerator.java:183)
	at org.eclipse.jgit.revwalk.RevWalk.next(RevWalk.java:421)
	at org.eclipse.jgit.revwalk.RevWalk.iterator(RevWalk.java:1257)
	... 31 more
[INFO] ------------------------------------------------------------------------
[INFO] BUILD FAILURE
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 5.157 s
[INFO] Finished at: 2017-07-20T04:11:05+00:00
[INFO] Final Memory: 29M/222M
[INFO] ------------------------------------------------------------------------
[ERROR] Failed to execute goal pl.project13.maven:git-commit-id-plugin:2.1.11:revision (default) on project eureka: Could not complete Mojo execution... Unable to find commits until some tag: Walk failure. Missing commit f81c16d587a532de071256f449255e5be21ea8a2 -> [Help 1]
[ERROR] 
[ERROR] To see the full stack trace of the errors, re-run Maven with the -e switch.
[ERROR] Re-run Maven using the -X switch to enable full debug logging.
[ERROR] 
[ERROR] For more information about the errors and possible solutions, please read the following articles:
[ERROR] [Help 1] http://cwiki.apache.org/confluence/display/MAVEN/MojoExecutionException
```

Skip 
```
			<plugin>
				<groupId>pl.project13.maven</groupId>
				<artifactId>git-commit-id-plugin</artifactId>
                <configuration>
                  <failOnNoGitDirectory>false</failOnNoGitDirectory>
				   <skip>true</skip>
                </configuration>
			</plugin>
```

Or 
```
			<plugin>
				<groupId>pl.project13.maven</groupId>
				<artifactId>git-commit-id-plugin</artifactId>
                <configuration>
                  <failOnNoGitDirectory>false</failOnNoGitDirectory>
				   <!--<skip>true</skip>-->
				   <useNativeGit>true</useNativeGit>
                </configuration>
			</plugin>
```

Build
```
[vagrant@localhost eureka]$ mvn clean install spring-boot:repackage docker:build -Dmaven.gitcommitid.skip=true -Dmaven.gitcommitid.useNativeGit=true
[INFO] Scanning for projects...
[INFO]                                                                         
[INFO] ------------------------------------------------------------------------
[INFO] Building Eureka Server 0.0.1-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] 
[INFO] --- maven-clean-plugin:2.6.1:clean (default-clean) @ eureka ---
[INFO] Deleting /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/eureka/target
[INFO] 
[INFO] --- git-commit-id-plugin:2.1.11:revision (default) @ eureka ---
[info] dotGitDirectory /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/eureka/.git
[info] git.build.user.name "Spencer Gibb"
[info] git.build.user.email "spencer@gibb.us"
[info] git.branch master
[info] git.commit.id.describe 
[info] git.commit.id ff513fe311ba3ce262027f40238a66ed2a164a6a
[info] git.commit.id.abbrev ff513fe
[info] git.commit.user.name "Spencer Gibb"
[info] git.commit.user.email "spencer@gibb.us"
[info] git.commit.message.full "remove unneeded `@EnableDiscoveryClient`"
[info] git.commit.message.short "remove unneeded `@EnableDiscoveryClient`"
[info] git.commit.time "2017-02-28 11:14:10 -0700"
[info] git.remote.origin.url https://github.com/spring-cloud-samples/eureka
[info] git.tags grafted,,->,origin/master,origin/
[info] git.build.time 2017-07-20T04:30:48+0000
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
[info] Writing properties file to [ /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/eureka/target/classes/git.properties ] (for module  Eureka Server )...
[info] Eureka Server ] project Eureka Server
[INFO] 
[INFO] --- maven-resources-plugin:2.6:resources (default-resources) @ eureka ---
... snip ...
[INFO] Copying /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/eureka/target/eureka-0.0.1-SNAPSHOT.jar -> /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/eureka/target/docker/eureka-0.0.1-SNAPSHOT.jar
[INFO] Copying src/main/docker/Dockerfile -> /Users/fanhongling/Downloads/workspace/src/github.com/spring-cloud-samples/eureka/target/docker/Dockerfile
[INFO] Building image springcloud/eureka
Step 1 : FROM openjdk:8-jre
Trying to pull repository docker.io/library/openjdk ... 
Pulling from docker.io/library/openjdk
c75480ad9aaf: Pull complete 
18d67befbc4e: Pull complete 
1f5d2d0853c7: Pull complete 
d35d9f711df2: Pull complete 
7b9badf7c23c: Pull complete 
e73fd941ef24: Pull complete 
119baa503f8c: Pull complete 
a2fd7e5b1f07: Pull complete 
Digest: sha256:30b4c10542c4ab5a8439c4957a7eaf734d920b7a914a7d982f4649615b99ca97
Status: Downloaded newer image for docker.io/openjdk:8-jre
---> d18ae8f18728
Step 2 : VOLUME /tmp
---> Running in a9254fa0c1bc
---> 234005e57fea
Removing intermediate container a9254fa0c1bc
Step 3 : ADD eureka-0.0.1-SNAPSHOT.jar /app.jar
---> 8bf6046b14a6
Removing intermediate container 0c55828ca8c3
Step 4 : RUN bash -c 'touch /app.jar'
---> Running in 34c936118b73
---> dc0185d5a9d6
Removing intermediate container 34c936118b73
Step 5 : EXPOSE 8761
---> Running in b7a4fa9e073e
---> 3cce2681fcdf
Removing intermediate container b7a4fa9e073e
Step 6 : ENTRYPOINT java -Djava.security.egd=file:/dev/./urandom -jar /app.jar
---> Running in aea3c9d4df5c
---> 636da1f3a19a
Removing intermediate container aea3c9d4df5c
Successfully built 636da1f3a19a
[INFO] Built springcloud/eureka
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time: 12:12 min
[INFO] Finished at: 2017-07-20T15:40:40+00:00
[INFO] Final Memory: 23M/186M
[INFO] ------------------------------------------------------------------------
```