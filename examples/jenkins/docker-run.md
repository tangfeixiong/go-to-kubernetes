
Hand-on
```
[vagrant@bogon ~]$ docker pull docker.io/fabric8/jenkins-docker:2.2.311
Trying to pull repository docker.io/fabric8/jenkins-docker ... 
2.2.311: Pulling from docker.io/fabric8/jenkins-docker
5040bd298390: Pull complete 
fce5728aad85: Pull complete 
76610ec20bf5: Pull complete 
60170fec2151: Pull complete 
e98f73de8f0d: Pull complete 
11f7af24ed9c: Pull complete 
49e2d6393f32: Pull complete 
bb9cdec9c7f3: Pull complete 
01e750d14101: Pull complete 
84fd58b8cd8d: Pull complete 
f131f10cb86f: Pull complete 
ac9c39133887: Pull complete 
e75457f9bfe8: Pull complete 
ea7f4cf742a8: Pull complete 
51cff2b7a540: Pull complete 
67b9cd9b179e: Pull complete 
4b1f366ac9d6: Pull complete 
d3b54d9fb0ac: Pull complete 
fa09c141497c: Pull complete 
04bac9eaa667: Pull complete 
330af5c46922: Pull complete 
cac5bfbceaa7: Pull complete 
8f922682c4a0: Pull complete 
17467d9b0eb1: Pull complete 
3f0ce885433c: Pull complete 
Digest: sha256:d2da5a18524d06ec9e310efa4c11d96d6cc2f77309a13f8077dfe08ba762a71e
Status: Downloaded newer image for docker.io/fabric8/jenkins-docker:2.2.311

# docker run -it -p 8080:8080 --name jenkins -e SEED_GIT_URL=https://github.com/fabric8io/default-jenkins-dsl.git -e NEXUS_SERVICE_HOST=dockerhost -e NEXUS_SERVICE_PORT=8081 fabric8/jenkins

[vagrant@localhost linux-bin]$ docker run -d -p 18080:8080 --name jenkins -e NEXUS_SERVICE_HOST=127.0.0.1:8081 -e NEXUS_SERVICE_PORT=8081 docker.io/fabric8/jenkins-docker:2.2.311
d7127155027542082be86c2b843e8a7c7f40fecaf8df8e7ed1300e86116831f9
```