## DevOps

CRD

Operator

Persistent Volume



Validate
```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pvc -l example.com/hadoop-operator=demo-hdfs-classic        
NAME                           STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS       AGE
hostpath-demo-hdfs-classic-0   Bound     pvc-9091d170-0bee-11e8-8484-525400224e72   80Mi       RWO            example-hostpath   18m
hostpath-demo-hdfs-classic-1   Bound     pvc-0ba28605-0bef-11e8-8484-525400224e72   80Mi       RWO            example-hostpath   15m
hostpath-demo-hdfs-classic-2   Bound     pvc-5a1a9b4e-0bef-11e8-8484-525400224e72   80Mi       RWO            example-hostpath   12m
hostpath-demo-hdfs-classic-3   Bound     pvc-85d0aea7-0bef-11e8-8484-525400224e72   80Mi       RWO            example-hostpath   11m
hostpath-demo-hdfs-classic-4   Bound     pvc-918b9025-0bef-11e8-8484-525400224e72   80Mi       RWO            example-hostpath   11m
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get pods -l example.com/hadoop-operator=demo-hdfs-classic -o wide
NAME                  READY     STATUS    RESTARTS   AGE       IP             NODE
demo-hdfs-classic-0   1/1       Running   0          8m        10.244.3.105   rookdev-172-17-4-63
demo-hdfs-classic-1   1/1       Running   0          4m        10.244.2.187   rookdev-172-17-4-61
demo-hdfs-classic-2   1/1       Running   0          2m        10.244.3.106   rookdev-172-17-4-63
demo-hdfs-classic-3   1/1       Running   0          1m        10.244.2.188   rookdev-172-17-4-61
demo-hdfs-classic-4   1/1       Running   0          55s       10.244.3.107   rookdev-172-17-4-63
```

Expose namenode web into _NodePort_
```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl expose pod demo-hdfs-classic-0 --port=9870 --type=NodePort
service "demo-hdfs-classic-0" exposed
```

Find _NodePort_
```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get svc
NAME                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
demo-hdfs-classic     ClusterIP   None            <none>        9000/TCP,9870/TCP   13m
demo-hdfs-classic-0   NodePort    10.110.35.141   <none>        9870:30538/TCP      39s
kubernetes            ClusterIP   10.96.0.1       <none>        443/TCP             51d
```

![屏幕快照 2018-02-07 上午2.20.49.png](./屏幕快照%202018-02-07%20上午2.20.49.png)

CI/CD
```
[vagrant@kubedev-172-17-4-59 hadoop-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go make go-bindata-spec     
```

```
[vagrant@kubedev-172-17-4-59 hadoop-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go build ./pkg/apis/...
```

```
[vagrant@kubedev-172-17-4-59 hadoop-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go build ./pkg/client/...
```


Hadoop Docker image
```
[vagrant@kubedev-172-17-4-59 hadoop-operator]$ make hadoop-docker
Sending build context to Docker daemon 129.5 kB
Step 1/8 : FROM openjdk:8-jre-slim
Trying to pull repository docker.io/library/openjdk ... 
sha256:c7d59dd74f9d6cd277544c4319cc33ce6b22f07ba0f88dcaf9b19b94f6219368: Pulling from docker.io/library/openjdk
e7bb522d92ff: Pull complete 
acf3a7df1b51: Pull complete 
c1c98005fcff: Pull complete 
39dcc90226db: Pull complete 
693e7ee43844: Pull complete 
cb083d9c54cf: Pull complete 
Digest: sha256:c7d59dd74f9d6cd277544c4319cc33ce6b22f07ba0f88dcaf9b19b94f6219368
Status: Downloaded newer image for docker.io/openjdk:8-jre-slim
 ---> 66d599bf7861
Step 2/8 : LABEL "maintainer" "tangfeixiong <tangfx128@gmail.com>" "project" "https://github.com/tangfeixiong/go-to-kubernetes" "name" "hadoop-operator" "version" "0.1-alpha" "created-by" '{"linux":"4.11.8-300.fc26.x86_64","docker":"1.13.1","hadoop":"3.0.0"}'
 ---> Running in 1ea6b899a892
 ---> 9f7e57d6edf5
Removing intermediate container 1ea6b899a892
Step 3/8 : ARG hadoop_dist_mirror=https://mirrors.tuna.tsinghua.edu.cn/apache/hadoop/common/hadoop-3.0.0/hadoop-3.0.0.tar.gz
 ---> Running in bb555dfe16c4
 ---> de13607bd4f5
Removing intermediate container bb555dfe16c4
Step 4/8 : RUN set -eux; 	apt-get update; 	apt-get install -y --no-install-recommends 		gnupg 		dirmngr 	; 	rm -rf /var/lib/apt/lists/*
 ---> Running in d08fa5311655
+ apt-get update
Get:2 http://security.debian.org stretch/updates InRelease [63.0 kB]
Ign:1 http://cdn-fastly.deb.debian.org/debian stretch InRelease
Get:3 http://cdn-fastly.deb.debian.org/debian stretch-updates InRelease [91.0 kB]
Get:4 http://cdn-fastly.deb.debian.org/debian stretch Release [118 kB]
Get:5 http://cdn-fastly.deb.debian.org/debian stretch-updates/main amd64 Packages [8384 B]
Get:6 http://security.debian.org stretch/updates/main amd64 Packages [338 kB]
Get:7 http://cdn-fastly.deb.debian.org/debian stretch Release.gpg [2434 B]
Get:8 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 Packages [9531 kB]
Fetched 10.2 MB in 19s (518 kB/s)
Reading package lists...
+ apt-get install -y --no-install-recommends gnupg dirmngr
Reading package lists...
Building dependency tree...
Reading state information...
The following additional packages will be installed:
  gnupg-agent libassuan0 libksba8 libldap-2.4-2 libldap-common libnpth0
  libreadline7 libsasl2-2 libsasl2-modules-db pinentry-curses readline-common
Suggested packages:
  dbus-user-session libpam-systemd pinentry-gnome3 tor parcimonie xloadimage
  scdaemon pinentry-doc readline-doc
Recommended packages:
  gnupg-l10n libsasl2-modules
The following NEW packages will be installed:
  dirmngr gnupg gnupg-agent libassuan0 libksba8 libldap-2.4-2 libldap-common
  libnpth0 libreadline7 libsasl2-2 libsasl2-modules-db pinentry-curses
  readline-common
0 upgraded, 13 newly installed, 0 to remove and 1 not upgraded.
Need to get 3178 kB of archives.
After this operation, 6602 kB of additional disk space will be used.
Get:1 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libassuan0 amd64 2.4.3-2 [42.5 kB]
Get:2 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 pinentry-curses amd64 1.0.0-2 [50.5 kB]
Get:3 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libnpth0 amd64 1.3-1 [14.6 kB]
Get:4 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 readline-common all 7.0-3 [70.4 kB]
Get:5 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libreadline7 amd64 7.0-3 [151 kB]
Get:6 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 gnupg-agent amd64 2.1.18-8~deb9u1 [553 kB]
Get:7 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libksba8 amd64 1.3.5-2 [99.7 kB]
Get:8 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 gnupg amd64 2.1.18-8~deb9u1 [1124 kB]
Get:9 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libsasl2-modules-db amd64 2.1.27~101-g0780600+dfsg-3 [68.2 kB]
Get:10 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libsasl2-2 amd64 2.1.27~101-g0780600+dfsg-3 [105 kB]
Get:11 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libldap-common all 2.4.44+dfsg-5+deb9u1 [85.4 kB]
Get:12 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libldap-2.4-2 amd64 2.4.44+dfsg-5+deb9u1 [219 kB]
Get:13 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 dirmngr amd64 2.1.18-8~deb9u1 [595 kB]
debconf: delaying package configuration, since apt-utils is not installed
Fetched 3178 kB in 7s (418 kB/s)
Selecting previously unselected package libassuan0:amd64.
(Reading database ... 7792 files and directories currently installed.)
Preparing to unpack .../00-libassuan0_2.4.3-2_amd64.deb ...
Unpacking libassuan0:amd64 (2.4.3-2) ...
Selecting previously unselected package pinentry-curses.
Preparing to unpack .../01-pinentry-curses_1.0.0-2_amd64.deb ...
Unpacking pinentry-curses (1.0.0-2) ...
Selecting previously unselected package libnpth0:amd64.
Preparing to unpack .../02-libnpth0_1.3-1_amd64.deb ...
Unpacking libnpth0:amd64 (1.3-1) ...
Selecting previously unselected package readline-common.
Preparing to unpack .../03-readline-common_7.0-3_all.deb ...
Unpacking readline-common (7.0-3) ...
Selecting previously unselected package libreadline7:amd64.
Preparing to unpack .../04-libreadline7_7.0-3_amd64.deb ...
Unpacking libreadline7:amd64 (7.0-3) ...
Selecting previously unselected package gnupg-agent.
Preparing to unpack .../05-gnupg-agent_2.1.18-8~deb9u1_amd64.deb ...
Unpacking gnupg-agent (2.1.18-8~deb9u1) ...
Selecting previously unselected package libksba8:amd64.
Preparing to unpack .../06-libksba8_1.3.5-2_amd64.deb ...
Unpacking libksba8:amd64 (1.3.5-2) ...
Selecting previously unselected package gnupg.
Preparing to unpack .../07-gnupg_2.1.18-8~deb9u1_amd64.deb ...
Unpacking gnupg (2.1.18-8~deb9u1) ...
Selecting previously unselected package libsasl2-modules-db:amd64.
Preparing to unpack .../08-libsasl2-modules-db_2.1.27~101-g0780600+dfsg-3_amd64.deb ...
Unpacking libsasl2-modules-db:amd64 (2.1.27~101-g0780600+dfsg-3) ...
Selecting previously unselected package libsasl2-2:amd64.
Preparing to unpack .../09-libsasl2-2_2.1.27~101-g0780600+dfsg-3_amd64.deb ...
Unpacking libsasl2-2:amd64 (2.1.27~101-g0780600+dfsg-3) ...
Selecting previously unselected package libldap-common.
Preparing to unpack .../10-libldap-common_2.4.44+dfsg-5+deb9u1_all.deb ...
Unpacking libldap-common (2.4.44+dfsg-5+deb9u1) ...
Selecting previously unselected package libldap-2.4-2:amd64.
Preparing to unpack .../11-libldap-2.4-2_2.4.44+dfsg-5+deb9u1_amd64.deb ...
Unpacking libldap-2.4-2:amd64 (2.4.44+dfsg-5+deb9u1) ...
Selecting previously unselected package dirmngr.
Preparing to unpack .../12-dirmngr_2.1.18-8~deb9u1_amd64.deb ...
Unpacking dirmngr (2.1.18-8~deb9u1) ...
Setting up libnpth0:amd64 (1.3-1) ...
Setting up readline-common (7.0-3) ...
Setting up libldap-common (2.4.44+dfsg-5+deb9u1) ...
Setting up libreadline7:amd64 (7.0-3) ...
Setting up libsasl2-modules-db:amd64 (2.1.27~101-g0780600+dfsg-3) ...
Setting up libsasl2-2:amd64 (2.1.27~101-g0780600+dfsg-3) ...
Setting up libksba8:amd64 (1.3.5-2) ...
Processing triggers for libc-bin (2.24-11+deb9u1) ...
Setting up libldap-2.4-2:amd64 (2.4.44+dfsg-5+deb9u1) ...
Setting up libassuan0:amd64 (2.4.3-2) ...
Setting up pinentry-curses (1.0.0-2) ...
Setting up gnupg-agent (2.1.18-8~deb9u1) ...
Setting up dirmngr (2.1.18-8~deb9u1) ...
Setting up gnupg (2.1.18-8~deb9u1) ...
Processing triggers for libc-bin (2.24-11+deb9u1) ...
+ rm -rf /var/lib/apt/lists/deb.debian.org_debian_dists_stretch-updates_InRelease /var/lib/apt/lists/deb.debian.org_debian_dists_stretch-updates_main_binary-amd64_Packages.lz4 /var/lib/apt/lists/deb.debian.org_debian_dists_stretch_Release /var/lib/apt/lists/deb.debian.org_debian_dists_stretch_Release.gpg /var/lib/apt/lists/deb.debian.org_debian_dists_stretch_main_binary-amd64_Packages.lz4 /var/lib/apt/lists/lock /var/lib/apt/lists/partial /var/lib/apt/lists/security.debian.org_dists_stretch_updates_InRelease /var/lib/apt/lists/security.debian.org_dists_stretch_updates_main_binary-amd64_Packages.lz4
 ---> 7f16979d454b
Removing intermediate container d08fa5311655
Step 5/8 : RUN set -eux; 		fetchDeps=' 		ca-certificates         curl 		wget 	'; 	apt-get update; 	apt-get install -y --no-install-recommends $fetchDeps; 	rm -rf /var/lib/apt/lists/*
 ---> Running in dedd898f203c
+ fetchDeps= 		ca-certificates         curl 		wget 	
+ apt-get update
Get:1 http://security.debian.org stretch/updates InRelease [63.0 kB]
Get:2 http://security.debian.org stretch/updates/main amd64 Packages [338 kB]
Ign:3 http://cdn-fastly.deb.debian.org/debian stretch InRelease
Get:4 http://cdn-fastly.deb.debian.org/debian stretch-updates InRelease [91.0 kB]
Get:5 http://cdn-fastly.deb.debian.org/debian stretch Release [118 kB]
Get:6 http://cdn-fastly.deb.debian.org/debian stretch-updates/main amd64 Packages [8384 B]
Get:7 http://cdn-fastly.deb.debian.org/debian stretch Release.gpg [2434 B]
Get:8 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 Packages [9531 kB]
Fetched 10.2 MB in 6s (1511 kB/s)
Reading package lists...
+ apt-get install -y --no-install-recommends ca-certificates curl wget
Reading package lists...
Building dependency tree...
Reading state information...
ca-certificates is already the newest version (20161130+nmu1).
ca-certificates set to manually installed.
The following additional packages will be installed:
  libcurl3 libidn2-0 libnghttp2-14 libpsl5 librtmp1 libssh2-1 libssl1.0.2
  libunistring0
Recommended packages:
  publicsuffix
The following NEW packages will be installed:
  curl libcurl3 libidn2-0 libnghttp2-14 libpsl5 librtmp1 libssh2-1 libssl1.0.2
  libunistring0 wget
0 upgraded, 10 newly installed, 0 to remove and 1 not upgraded.
Need to get 3270 kB of archives.
After this operation, 9478 kB of additional disk space will be used.
Get:1 http://security.debian.org stretch/updates/main amd64 libssl1.0.2 amd64 1.0.2l-2+deb9u2 [1294 kB]
Get:9 http://security.debian.org stretch/updates/main amd64 libcurl3 amd64 7.52.1-5+deb9u4 [291 kB]
Get:10 http://security.debian.org stretch/updates/main amd64 curl amd64 7.52.1-5+deb9u4 [227 kB]
Get:2 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libunistring0 amd64 0.9.6+really0.9.3-0.1 [279 kB]
Get:3 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libidn2-0 amd64 0.16-1+deb9u1 [60.7 kB]
Get:4 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libpsl5 amd64 0.17.0-3 [41.8 kB]
Get:5 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 wget amd64 1.18-5+deb9u1 [800 kB]
Get:6 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libnghttp2-14 amd64 1.18.1-1 [79.1 kB]
Get:7 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 librtmp1 amd64 2.4+20151223.gitfa8646d.1-1+b1 [60.4 kB]
Get:8 http://cdn-fastly.deb.debian.org/debian stretch/main amd64 libssh2-1 amd64 1.7.0-1 [138 kB]
debconf: delaying package configuration, since apt-utils is not installed
Fetched 3270 kB in 45s (71.5 kB/s)
Selecting previously unselected package libssl1.0.2:amd64.
(Reading database ... 7985 files and directories currently installed.)
Preparing to unpack .../0-libssl1.0.2_1.0.2l-2+deb9u2_amd64.deb ...
Unpacking libssl1.0.2:amd64 (1.0.2l-2+deb9u2) ...
Selecting previously unselected package libunistring0:amd64.
Preparing to unpack .../1-libunistring0_0.9.6+really0.9.3-0.1_amd64.deb ...
Unpacking libunistring0:amd64 (0.9.6+really0.9.3-0.1) ...
Selecting previously unselected package libidn2-0:amd64.
Preparing to unpack .../2-libidn2-0_0.16-1+deb9u1_amd64.deb ...
Unpacking libidn2-0:amd64 (0.16-1+deb9u1) ...
Selecting previously unselected package libpsl5:amd64.
Preparing to unpack .../3-libpsl5_0.17.0-3_amd64.deb ...
Unpacking libpsl5:amd64 (0.17.0-3) ...
Selecting previously unselected package wget.
Preparing to unpack .../4-wget_1.18-5+deb9u1_amd64.deb ...
Unpacking wget (1.18-5+deb9u1) ...
Selecting previously unselected package libnghttp2-14:amd64.
Preparing to unpack .../5-libnghttp2-14_1.18.1-1_amd64.deb ...
Unpacking libnghttp2-14:amd64 (1.18.1-1) ...
Selecting previously unselected package librtmp1:amd64.
Preparing to unpack .../6-librtmp1_2.4+20151223.gitfa8646d.1-1+b1_amd64.deb ...
Unpacking librtmp1:amd64 (2.4+20151223.gitfa8646d.1-1+b1) ...
Selecting previously unselected package libssh2-1:amd64.
Preparing to unpack .../7-libssh2-1_1.7.0-1_amd64.deb ...
Unpacking libssh2-1:amd64 (1.7.0-1) ...
Selecting previously unselected package libcurl3:amd64.
Preparing to unpack .../8-libcurl3_7.52.1-5+deb9u4_amd64.deb ...
Unpacking libcurl3:amd64 (7.52.1-5+deb9u4) ...
Selecting previously unselected package curl.
Preparing to unpack .../9-curl_7.52.1-5+deb9u4_amd64.deb ...
Unpacking curl (7.52.1-5+deb9u4) ...
Setting up libnghttp2-14:amd64 (1.18.1-1) ...
Setting up librtmp1:amd64 (2.4+20151223.gitfa8646d.1-1+b1) ...
Setting up libssl1.0.2:amd64 (1.0.2l-2+deb9u2) ...
debconf: unable to initialize frontend: Dialog
debconf: (TERM is not set, so the dialog frontend is not usable.)
debconf: falling back to frontend: Readline
debconf: unable to initialize frontend: Readline
debconf: (Can't locate Term/ReadLine.pm in @INC (you may need to install the Term::ReadLine module) (@INC contains: /etc/perl /usr/local/lib/x86_64-linux-gnu/perl/5.24.1 /usr/local/share/perl/5.24.1 /usr/lib/x86_64-linux-gnu/perl5/5.24 /usr/share/perl5 /usr/lib/x86_64-linux-gnu/perl/5.24 /usr/share/perl/5.24 /usr/local/lib/site_perl /usr/lib/x86_64-linux-gnu/perl-base .) at /usr/share/perl5/Debconf/FrontEnd/Readline.pm line 7.)
debconf: falling back to frontend: Teletype
Setting up libssh2-1:amd64 (1.7.0-1) ...
Processing triggers for libc-bin (2.24-11+deb9u1) ...
Setting up libunistring0:amd64 (0.9.6+really0.9.3-0.1) ...
Setting up libidn2-0:amd64 (0.16-1+deb9u1) ...
Setting up libpsl5:amd64 (0.17.0-3) ...
Setting up wget (1.18-5+deb9u1) ...
Setting up libcurl3:amd64 (7.52.1-5+deb9u4) ...
Setting up curl (7.52.1-5+deb9u4) ...
Processing triggers for libc-bin (2.24-11+deb9u1) ...
+ rm -rf /var/lib/apt/lists/deb.debian.org_debian_dists_stretch-updates_InRelease /var/lib/apt/lists/deb.debian.org_debian_dists_stretch-updates_main_binary-amd64_Packages.lz4 /var/lib/apt/lists/deb.debian.org_debian_dists_stretch_Release /var/lib/apt/lists/deb.debian.org_debian_dists_stretch_Release.gpg /var/lib/apt/lists/deb.debian.org_debian_dists_stretch_main_binary-amd64_Packages.lz4 /var/lib/apt/lists/lock /var/lib/apt/lists/partial /var/lib/apt/lists/security.debian.org_dists_stretch_updates_InRelease /var/lib/apt/lists/security.debian.org_dists_stretch_updates_main_binary-amd64_Packages.lz4
 ---> f33a5178d2da
Removing intermediate container dedd898f203c
Step 6/8 : RUN curl -jkSL $hadoop_dist_mirror | tar -C / -zx -f-     hadoop-3.0.0/sbin/     hadoop-3.0.0/include/     hadoop-3.0.0/LICENSE.txt     hadoop-3.0.0/libexec/     hadoop-3.0.0/README.txt     hadoop-3.0.0/NOTICE.txt     hadoop-3.0.0/lib/     hadoop-3.0.0/share/hadoop/     hadoop-3.0.0/bin/     hadoop-3.0.0/etc/
 ---> Running in cd1bf043aa6f
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  292M  100  292M    0     0  1578k      0  0:03:09  0:03:09 --:--:--  531k
 ---> 98ba99b5fce6
Removing intermediate container cd1bf043aa6f
Step 7/8 : ENV DOCKER_API_VERSION '1.12' DOCKER_CONFIG_JSON '{"auths": {"localhost:5000": {"auth": "","email": ""}}}' REGISTRY_CERTS_JSON '{"localhost:5000": {"ca_base64": "", "crt_base64": "", "key_base64": ""}}' HADOOP_VERSION 3.0.0
 ---> Running in b9b5f06b2013
 ---> 387ed1dfa0c5
Removing intermediate container b9b5f06b2013
Step 8/8 : EXPOSE 9000 50010 50020 50070 50075 50090 9871 9870 9820 9869 9868 9867 9866 9865 9864 19888 8030 8031 8032 8033 8040 8042 8088 8188 49707 2122
 ---> Running in 9d1b5b5baf0b
 ---> f2e0aa59ad40
Removing intermediate container 9d1b5b5baf0b
Successfully built f2e0aa59ad40
```

Lab

```
root@480e60ebed55:/# hadoop-3.0.0/bin/hdfs namenode -format demo
2018-02-04 12:27:45,403 INFO namenode.NameNode: STARTUP_MSG: 
/************************************************************
STARTUP_MSG: Starting NameNode
STARTUP_MSG:   host = 480e60ebed55/172.18.0.3
STARTUP_MSG:   args = [-format, demo]
STARTUP_MSG:   version = 3.0.0
STARTUP_MSG:   classpath = /hadoop-3.0.0/etc/hadoop:/hadoop-3.0.0/share/hadoop/common/lib/jersey-core-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/netty-3.10.5.Final.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-server-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-webapp-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-io-2.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/zookeeper-3.4.9.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-common-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/avro-1.7.7.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-jaxrs-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsr311-api-1.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/guava-11.0.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-asn1-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-databind-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-beanutils-1.9.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-http-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-core-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-lang-2.6.jar:/hadoop-3.0.0/share/hadoop/common/lib/httpcore-4.4.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-json-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/json-smart-2.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/snappy-java-1.0.5.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-admin-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-cli-1.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/jettison-1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/httpclient-4.5.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/jcip-annotations-1.0-1.jar:/hadoop-3.0.0/share/hadoop/common/lib/xz-1.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-server-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-client-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-xml-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-codec-1.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/asm-5.0.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/woodstox-core-5.0.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-io-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-servlet-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/jaxb-api-2.2.11.jar:/hadoop-3.0.0/share/hadoop/common/lib/accessors-smart-1.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/protobuf-java-2.5.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-recipes-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-util-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/paranamer-2.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/hadoop-annotations-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-math3-3.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-config-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-core-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-xdr-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/nimbus-jose-jwt-4.41.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jaxb-impl-2.2.3-1.jar:/hadoop-3.0.0/share/hadoop/common/lib/gson-2.2.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-lang3-3.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-servlet-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/re2j-1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-collections-3.2.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-logging-1.1.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-framework-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-simplekdc-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-net-3.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-compress-1.4.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/slf4j-api-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-xc-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsch-0.1.54.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-pkix-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-crypto-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jul-to-slf4j-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/log4j-1.2.17.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/junit-4.11.jar:/hadoop-3.0.0/share/hadoop/common/lib/mockito-all-1.8.5.jar:/hadoop-3.0.0/share/hadoop/common/lib/metrics-core-3.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/stax2-api-3.1.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsp-api-2.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/slf4j-log4j12-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/htrace-core4-4.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/common/lib/hadoop-auth-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsr305-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-mapper-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-configuration2-2.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-core-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-client-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/hamcrest-core-1.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-identity-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-security-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/javax.servlet-api-3.1.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-server-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/token-provider-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-nfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-common-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-kms-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-core-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/netty-3.10.5.Final.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-server-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-webapp-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-io-2.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/zookeeper-3.4.9.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-common-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/avro-1.7.7.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-jaxrs-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsr311-api-1.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/guava-11.0.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-asn1-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-databind-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-beanutils-1.9.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-http-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-core-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-lang-2.6.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/httpcore-4.4.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-json-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/leveldbjni-all-1.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-daemon-1.0.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/json-smart-2.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/snappy-java-1.0.5.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-admin-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-cli-1.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jettison-1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/httpclient-4.5.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jcip-annotations-1.0-1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/xz-1.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-server-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-client-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-xml-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-codec-1.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/asm-5.0.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/woodstox-core-5.0.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-util-ajax-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-io-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-servlet-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jaxb-api-2.2.11.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/accessors-smart-1.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/protobuf-java-2.5.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-recipes-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-util-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/paranamer-2.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/hadoop-annotations-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/okio-1.4.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-math3-3.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-config-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-core-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-xdr-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/nimbus-jose-jwt-4.41.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jaxb-impl-2.2.3-1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/json-simple-1.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/gson-2.2.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-lang3-3.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-servlet-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/re2j-1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-collections-3.2.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-logging-1.1.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-framework-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-simplekdc-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-net-3.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-compress-1.4.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-xc-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsch-0.1.54.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-pkix-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-crypto-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/log4j-1.2.17.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/stax2-api-3.1.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/htrace-core4-4.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/hadoop-auth-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsr305-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-mapper-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-configuration2-2.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-core-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-client-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/netty-all-4.0.23.Final.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-identity-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-security-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/okhttp-2.4.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/javax.servlet-api-3.1.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-server-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/token-provider-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-native-client-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-nfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-native-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-httpfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-client-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-jobclient-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-shuffle-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-examples-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-hs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-nativetask-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-jobclient-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-hs-plugins-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-core-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-app-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-module-jaxb-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-csv-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-math-2.2.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/guice-4.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/json-io-2.5.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jcodings-1.0.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/java-util-1.9.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-jaxrs-json-provider-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/HikariCP-java7-2.4.12.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-prefix-tree-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/mssql-jdbc-6.2.1.jre7.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-common-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/joni-2.1.2.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/findbugs-annotations-1.3.9-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-server-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-el-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jamon-runtime-2.4.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/guice-servlet-4.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/disruptor-3.3.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jersey-client-1.19.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/htrace-core-3.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jsp-api-2.1-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/javax.inject-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-jaxrs-base-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-procedure-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-protocol-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/geronimo-jcache_1.0_spec-1.0-alpha-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/fst-2.50.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-hadoop2-compat-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-hadoop-compat-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-client-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-annotations-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jasper-compiler-5.5.23.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/aopalliance-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/metrics-core-3.0.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jasper-runtime-5.5.23.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jsp-2.1-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/ehcache-3.3.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/servlet-api-2.5-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/metrics-core-2.2.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jersey-guice-1.19.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-httpclient-3.1.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-registry-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-router-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-resourcemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-tests-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-applicationhistoryservice-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-sharedcachemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-api-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-nodemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-hbase-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-hbase-tests-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timeline-pluginstorage-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-web-proxy-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-applications-unmanaged-am-launcher-3.0.0.jar
STARTUP_MSG:   build = https://git-wip-us.apache.org/repos/asf/hadoop.git -r c25427ceca461ee979d30edd7a4b0f50718e6533; compiled by 'andrew' on 2017-12-08T19:16Z
STARTUP_MSG:   java = 1.8.0_151
************************************************************/
2018-02-04 12:27:45,426 INFO namenode.NameNode: registered UNIX signal handlers for [TERM, HUP, INT]
2018-02-04 12:27:45,430 INFO namenode.NameNode: createNameNode [-format, demo]
Formatting using clusterid: CID-af1d7218-0a19-4815-be03-bbaa5c3abcba
2018-02-04 12:27:46,776 INFO namenode.FSEditLog: Edit logging is async:true
2018-02-04 12:27:46,835 INFO namenode.FSNamesystem: KeyProvider: null
2018-02-04 12:27:46,841 INFO namenode.FSNamesystem: fsLock is fair: true
2018-02-04 12:27:46,852 INFO namenode.FSNamesystem: Detailed lock hold time metrics enabled: false
2018-02-04 12:27:46,880 INFO namenode.FSNamesystem: fsOwner             = root (auth:SIMPLE)
2018-02-04 12:27:46,880 INFO namenode.FSNamesystem: supergroup          = supergroup
2018-02-04 12:27:46,881 INFO namenode.FSNamesystem: isPermissionEnabled = true
2018-02-04 12:27:46,881 INFO namenode.FSNamesystem: HA Enabled: false
2018-02-04 12:27:46,999 INFO common.Util: dfs.datanode.fileio.profiling.sampling.percentage set to 0. Disabling file IO profiling
2018-02-04 12:27:47,042 INFO blockmanagement.DatanodeManager: dfs.block.invalidate.limit: configured=1000, counted=60, effected=1000
2018-02-04 12:27:47,042 INFO blockmanagement.DatanodeManager: dfs.namenode.datanode.registration.ip-hostname-check=true
2018-02-04 12:27:47,049 INFO blockmanagement.BlockManager: dfs.namenode.startup.delay.block.deletion.sec is set to 000:00:00:00.000
2018-02-04 12:27:47,059 INFO blockmanagement.BlockManager: The block deletion will start around 2018 Feb 04 12:27:47
2018-02-04 12:27:47,073 INFO util.GSet: Computing capacity for map BlocksMap
2018-02-04 12:27:47,074 INFO util.GSet: VM type       = 64-bit
2018-02-04 12:27:47,092 INFO util.GSet: 2.0% max memory 483.4 MB = 9.7 MB
2018-02-04 12:27:47,092 INFO util.GSet: capacity      = 2^20 = 1048576 entries
2018-02-04 12:27:47,180 INFO blockmanagement.BlockManager: dfs.block.access.token.enable = false
2018-02-04 12:27:47,195 INFO Configuration.deprecation: No unit for dfs.namenode.safemode.extension(30000) assuming MILLISECONDS
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManagerSafeMode: dfs.namenode.safemode.threshold-pct = 0.9990000128746033
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManagerSafeMode: dfs.namenode.safemode.min.datanodes = 0
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManagerSafeMode: dfs.namenode.safemode.extension = 30000
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManager: defaultReplication         = 3
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManager: maxReplication             = 512
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManager: minReplication             = 1
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManager: maxReplicationStreams      = 2
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManager: redundancyRecheckInterval  = 3000ms
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManager: encryptDataTransfer        = false
2018-02-04 12:27:47,195 INFO blockmanagement.BlockManager: maxNumBlocksToLog          = 1000
2018-02-04 12:27:47,288 INFO util.GSet: Computing capacity for map INodeMap
2018-02-04 12:27:47,289 INFO util.GSet: VM type       = 64-bit
2018-02-04 12:27:47,289 INFO util.GSet: 1.0% max memory 483.4 MB = 4.8 MB
2018-02-04 12:27:47,289 INFO util.GSet: capacity      = 2^19 = 524288 entries
2018-02-04 12:27:47,290 INFO namenode.FSDirectory: ACLs enabled? false
2018-02-04 12:27:47,290 INFO namenode.FSDirectory: POSIX ACL inheritance enabled? true
2018-02-04 12:27:47,290 INFO namenode.FSDirectory: XAttrs enabled? true
2018-02-04 12:27:47,290 INFO namenode.NameNode: Caching file names occurring more than 10 times
2018-02-04 12:27:47,303 INFO snapshot.SnapshotManager: Loaded config captureOpenFiles: false, skipCaptureAccessTimeOnlyChange: false, snapshotDiffAllowSnapRootDescendant: true
2018-02-04 12:27:47,313 INFO util.GSet: Computing capacity for map cachedBlocks
2018-02-04 12:27:47,319 INFO util.GSet: VM type       = 64-bit
2018-02-04 12:27:47,319 INFO util.GSet: 0.25% max memory 483.4 MB = 1.2 MB
2018-02-04 12:27:47,319 INFO util.GSet: capacity      = 2^17 = 131072 entries
2018-02-04 12:27:47,338 INFO metrics.TopMetrics: NNTop conf: dfs.namenode.top.window.num.buckets = 10
2018-02-04 12:27:47,338 INFO metrics.TopMetrics: NNTop conf: dfs.namenode.top.num.users = 10
2018-02-04 12:27:47,338 INFO metrics.TopMetrics: NNTop conf: dfs.namenode.top.windows.minutes = 1,5,25
2018-02-04 12:27:47,354 INFO namenode.FSNamesystem: Retry cache on namenode is enabled
2018-02-04 12:27:47,356 INFO namenode.FSNamesystem: Retry cache will use 0.03 of total heap and retry cache entry expiry time is 600000 millis
2018-02-04 12:27:47,364 INFO util.GSet: Computing capacity for map NameNodeRetryCache
2018-02-04 12:27:47,364 INFO util.GSet: VM type       = 64-bit
2018-02-04 12:27:47,364 INFO util.GSet: 0.029999999329447746% max memory 483.4 MB = 148.5 KB
2018-02-04 12:27:47,364 INFO util.GSet: capacity      = 2^14 = 16384 entries
2018-02-04 12:27:47,451 INFO namenode.FSImage: Allocated new BlockPoolId: BP-2074120163-172.18.0.3-1517747267431
2018-02-04 12:27:47,493 INFO common.Storage: Storage directory /tmp/hadoop-root/dfs/name has been successfully formatted.
2018-02-04 12:27:47,539 INFO namenode.FSImageFormatProtobuf: Saving image file /tmp/hadoop-root/dfs/name/current/fsimage.ckpt_0000000000000000000 using no compression
2018-02-04 12:27:47,711 INFO namenode.FSImageFormatProtobuf: Image file /tmp/hadoop-root/dfs/name/current/fsimage.ckpt_0000000000000000000 of size 389 bytes saved in 0 seconds.
2018-02-04 12:27:47,734 INFO namenode.NNStorageRetentionManager: Going to retain 1 images with txid >= 0
2018-02-04 12:27:47,743 INFO namenode.NameNode: SHUTDOWN_MSG: 
/************************************************************
SHUTDOWN_MSG: Shutting down NameNode at 480e60ebed55/172.18.0.3
************************************************************/
```

```
root@480e60ebed55:/# hadoop-3.0.0/bin/hdfs namenode
2018-02-04 12:33:57,868 INFO namenode.NameNode: STARTUP_MSG: 
/************************************************************
STARTUP_MSG: Starting NameNode
STARTUP_MSG:   host = 480e60ebed55/172.18.0.3
STARTUP_MSG:   args = []
STARTUP_MSG:   version = 3.0.0
STARTUP_MSG:   classpath = /hadoop-3.0.0/etc/hadoop:/hadoop-3.0.0/share/hadoop/common/lib/jersey-core-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/netty-3.10.5.Final.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-server-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-webapp-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-io-2.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/zookeeper-3.4.9.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-common-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/avro-1.7.7.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-jaxrs-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsr311-api-1.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/guava-11.0.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-asn1-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-databind-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-beanutils-1.9.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-http-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-core-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-lang-2.6.jar:/hadoop-3.0.0/share/hadoop/common/lib/httpcore-4.4.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-json-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/json-smart-2.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/snappy-java-1.0.5.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-admin-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-cli-1.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/jettison-1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/httpclient-4.5.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/jcip-annotations-1.0-1.jar:/hadoop-3.0.0/share/hadoop/common/lib/xz-1.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-server-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-client-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-xml-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-codec-1.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/asm-5.0.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/woodstox-core-5.0.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-io-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-servlet-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/jaxb-api-2.2.11.jar:/hadoop-3.0.0/share/hadoop/common/lib/accessors-smart-1.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/protobuf-java-2.5.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-recipes-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-util-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/paranamer-2.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/hadoop-annotations-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-math3-3.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-config-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-core-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-xdr-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/nimbus-jose-jwt-4.41.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jaxb-impl-2.2.3-1.jar:/hadoop-3.0.0/share/hadoop/common/lib/gson-2.2.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-lang3-3.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-servlet-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/re2j-1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-collections-3.2.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-logging-1.1.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-framework-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-simplekdc-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-net-3.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-compress-1.4.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/slf4j-api-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-xc-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsch-0.1.54.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-pkix-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-crypto-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jul-to-slf4j-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/log4j-1.2.17.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/junit-4.11.jar:/hadoop-3.0.0/share/hadoop/common/lib/mockito-all-1.8.5.jar:/hadoop-3.0.0/share/hadoop/common/lib/metrics-core-3.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/stax2-api-3.1.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsp-api-2.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/slf4j-log4j12-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/htrace-core4-4.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/common/lib/hadoop-auth-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsr305-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-mapper-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-configuration2-2.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-core-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-client-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/hamcrest-core-1.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-identity-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-security-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/javax.servlet-api-3.1.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-server-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/token-provider-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-nfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-common-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-kms-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-core-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/netty-3.10.5.Final.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-server-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-webapp-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-io-2.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/zookeeper-3.4.9.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-common-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/avro-1.7.7.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-jaxrs-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsr311-api-1.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/guava-11.0.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-asn1-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-databind-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-beanutils-1.9.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-http-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-core-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-lang-2.6.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/httpcore-4.4.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-json-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/leveldbjni-all-1.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-daemon-1.0.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/json-smart-2.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/snappy-java-1.0.5.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-admin-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-cli-1.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jettison-1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/httpclient-4.5.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jcip-annotations-1.0-1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/xz-1.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-server-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-client-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-xml-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-codec-1.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/asm-5.0.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/woodstox-core-5.0.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-util-ajax-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-io-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-servlet-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jaxb-api-2.2.11.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/accessors-smart-1.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/protobuf-java-2.5.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-recipes-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-util-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/paranamer-2.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/hadoop-annotations-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/okio-1.4.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-math3-3.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-config-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-core-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-xdr-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/nimbus-jose-jwt-4.41.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jaxb-impl-2.2.3-1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/json-simple-1.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/gson-2.2.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-lang3-3.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-servlet-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/re2j-1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-collections-3.2.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-logging-1.1.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-framework-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-simplekdc-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-net-3.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-compress-1.4.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-xc-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsch-0.1.54.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-pkix-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-crypto-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/log4j-1.2.17.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/stax2-api-3.1.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/htrace-core4-4.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/hadoop-auth-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsr305-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-mapper-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-configuration2-2.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-core-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-client-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/netty-all-4.0.23.Final.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-identity-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-security-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/okhttp-2.4.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/javax.servlet-api-3.1.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-server-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/token-provider-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-native-client-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-nfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-native-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-httpfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-client-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-jobclient-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-shuffle-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-examples-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-hs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-nativetask-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-jobclient-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-hs-plugins-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-core-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-app-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-module-jaxb-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-csv-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-math-2.2.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/guice-4.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/json-io-2.5.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jcodings-1.0.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/java-util-1.9.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-jaxrs-json-provider-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/HikariCP-java7-2.4.12.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-prefix-tree-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/mssql-jdbc-6.2.1.jre7.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-common-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/joni-2.1.2.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/findbugs-annotations-1.3.9-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-server-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-el-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jamon-runtime-2.4.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/guice-servlet-4.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/disruptor-3.3.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jersey-client-1.19.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/htrace-core-3.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jsp-api-2.1-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/javax.inject-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-jaxrs-base-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-procedure-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-protocol-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/geronimo-jcache_1.0_spec-1.0-alpha-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/fst-2.50.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-hadoop2-compat-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-hadoop-compat-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-client-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-annotations-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jasper-compiler-5.5.23.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/aopalliance-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/metrics-core-3.0.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jasper-runtime-5.5.23.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jsp-2.1-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/ehcache-3.3.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/servlet-api-2.5-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/metrics-core-2.2.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jersey-guice-1.19.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-httpclient-3.1.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-registry-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-router-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-resourcemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-tests-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-applicationhistoryservice-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-sharedcachemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-api-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-nodemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-hbase-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-hbase-tests-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timeline-pluginstorage-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-web-proxy-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-applications-unmanaged-am-launcher-3.0.0.jar
STARTUP_MSG:   build = https://git-wip-us.apache.org/repos/asf/hadoop.git -r c25427ceca461ee979d30edd7a4b0f50718e6533; compiled by 'andrew' on 2017-12-08T19:16Z
STARTUP_MSG:   java = 1.8.0_151
************************************************************/
2018-02-04 12:33:57,898 INFO namenode.NameNode: registered UNIX signal handlers for [TERM, HUP, INT]
2018-02-04 12:33:57,903 INFO namenode.NameNode: createNameNode []
2018-02-04 12:33:58,372 INFO beanutils.FluentPropertyBeanIntrospector: Error when creating PropertyDescriptor for public final void org.apache.commons.configuration2.AbstractConfiguration.setProperty(java.lang.String,java.lang.Object)! Ignoring this property.
2018-02-04 12:33:58,407 INFO impl.MetricsConfig: loaded properties from hadoop-metrics2.properties
2018-02-04 12:33:58,561 INFO impl.MetricsSystemImpl: Scheduled Metric snapshot period at 10 second(s).
2018-02-04 12:33:58,561 INFO impl.MetricsSystemImpl: NameNode metrics system started
2018-02-04 12:33:58,585 INFO namenode.NameNode: fs.defaultFS is file:///
2018-02-04 12:33:58,800 ERROR namenode.NameNode: Failed to start namenode.
java.lang.IllegalArgumentException: Invalid URI for NameNode address (check fs.defaultFS): file:/// has no authority.
	at org.apache.hadoop.hdfs.DFSUtilClient.getNNAddress(DFSUtilClient.java:668)
	at org.apache.hadoop.hdfs.DFSUtilClient.getNNAddressCheckLogical(DFSUtilClient.java:697)
	at org.apache.hadoop.hdfs.DFSUtilClient.getNNAddress(DFSUtilClient.java:659)
	at org.apache.hadoop.hdfs.server.namenode.NameNode.getRpcServerAddress(NameNode.java:562)
	at org.apache.hadoop.hdfs.server.namenode.NameNode.loginAsNameNodeUser(NameNode.java:693)
	at org.apache.hadoop.hdfs.server.namenode.NameNode.initialize(NameNode.java:713)
	at org.apache.hadoop.hdfs.server.namenode.NameNode.<init>(NameNode.java:950)
	at org.apache.hadoop.hdfs.server.namenode.NameNode.<init>(NameNode.java:929)
	at org.apache.hadoop.hdfs.server.namenode.NameNode.createNameNode(NameNode.java:1653)
	at org.apache.hadoop.hdfs.server.namenode.NameNode.main(NameNode.java:1720)
2018-02-04 12:33:58,805 INFO util.ExitUtil: Exiting with status 1: java.lang.IllegalArgumentException: Invalid URI for NameNode address (check fs.defaultFS): file:/// has no authority.
2018-02-04 12:33:58,807 INFO namenode.NameNode: SHUTDOWN_MSG: 
/************************************************************
SHUTDOWN_MSG: Shutting down NameNode at 480e60ebed55/172.18.0.3
************************************************************/
```


```
[vagrant@kubedev-172-17-4-59 hadoop-operator]$ docker run --rm --name hadoop -ti tangfeixiong/hadoop-docker bash
```

```
root@480e60ebed55:/# echo '<configuration>
>   <property>
>     <name>fs.defaultFS</name>
>     <value>hdfs://127.0.0.1:9000</value>
>   </property>
> </configuration>' > core-site.xml
```

```
root@480e60ebed55:/# cp core-site.xml hadoop-3.0.0/etc/hadoop/core-site.xml
```

```
root@480e60ebed55:/# hadoop-3.0.0/bin/hdfs namenode -format demo
2018-02-04 14:28:05,026 INFO namenode.NameNode: STARTUP_MSG: 
/************************************************************
STARTUP_MSG: Starting NameNode
STARTUP_MSG:   host = 480e60ebed55/172.18.0.3
STARTUP_MSG:   args = [-format, demo]
STARTUP_MSG:   version = 3.0.0
STARTUP_MSG:   classpath = /hadoop-3.0.0/etc/hadoop:/hadoop-3.0.0/share/hadoop/common/lib/jersey-core-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/netty-3.10.5.Final.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-server-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-webapp-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-io-2.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/zookeeper-3.4.9.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-common-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/avro-1.7.7.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-jaxrs-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsr311-api-1.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/guava-11.0.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-asn1-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-databind-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-beanutils-1.9.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-http-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-core-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-lang-2.6.jar:/hadoop-3.0.0/share/hadoop/common/lib/httpcore-4.4.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-json-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/json-smart-2.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/snappy-java-1.0.5.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-admin-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-cli-1.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/jettison-1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/httpclient-4.5.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/jcip-annotations-1.0-1.jar:/hadoop-3.0.0/share/hadoop/common/lib/xz-1.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-server-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-client-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-xml-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-codec-1.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/asm-5.0.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/woodstox-core-5.0.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-io-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-servlet-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/jaxb-api-2.2.11.jar:/hadoop-3.0.0/share/hadoop/common/lib/accessors-smart-1.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/protobuf-java-2.5.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-recipes-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-util-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/paranamer-2.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/hadoop-annotations-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-math3-3.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-config-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-core-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-xdr-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/nimbus-jose-jwt-4.41.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jaxb-impl-2.2.3-1.jar:/hadoop-3.0.0/share/hadoop/common/lib/gson-2.2.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-lang3-3.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-servlet-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/re2j-1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-collections-3.2.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-logging-1.1.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-framework-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-simplekdc-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-net-3.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-compress-1.4.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/slf4j-api-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-xc-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsch-0.1.54.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-pkix-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-crypto-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jul-to-slf4j-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/log4j-1.2.17.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/junit-4.11.jar:/hadoop-3.0.0/share/hadoop/common/lib/mockito-all-1.8.5.jar:/hadoop-3.0.0/share/hadoop/common/lib/metrics-core-3.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/stax2-api-3.1.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsp-api-2.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/slf4j-log4j12-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/htrace-core4-4.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/common/lib/hadoop-auth-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsr305-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-mapper-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-configuration2-2.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-core-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-client-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/hamcrest-core-1.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-identity-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-security-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/javax.servlet-api-3.1.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-server-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/token-provider-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-nfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-common-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-kms-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-core-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/netty-3.10.5.Final.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-server-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-webapp-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-io-2.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/zookeeper-3.4.9.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-common-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/avro-1.7.7.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-jaxrs-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsr311-api-1.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/guava-11.0.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-asn1-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-databind-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-beanutils-1.9.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-http-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-core-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-lang-2.6.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/httpcore-4.4.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-json-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/leveldbjni-all-1.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-daemon-1.0.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/json-smart-2.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/snappy-java-1.0.5.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-admin-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-cli-1.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jettison-1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/httpclient-4.5.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jcip-annotations-1.0-1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/xz-1.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-server-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-client-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-xml-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-codec-1.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/asm-5.0.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/woodstox-core-5.0.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-util-ajax-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-io-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-servlet-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jaxb-api-2.2.11.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/accessors-smart-1.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/protobuf-java-2.5.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-recipes-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-util-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/paranamer-2.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/hadoop-annotations-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/okio-1.4.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-math3-3.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-config-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-core-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-xdr-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/nimbus-jose-jwt-4.41.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jaxb-impl-2.2.3-1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/json-simple-1.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/gson-2.2.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-lang3-3.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-servlet-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/re2j-1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-collections-3.2.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-logging-1.1.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-framework-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-simplekdc-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-net-3.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-compress-1.4.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-xc-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsch-0.1.54.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-pkix-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-crypto-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/log4j-1.2.17.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/stax2-api-3.1.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/htrace-core4-4.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/hadoop-auth-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsr305-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-mapper-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-configuration2-2.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-core-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-client-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/netty-all-4.0.23.Final.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-identity-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-security-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/okhttp-2.4.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/javax.servlet-api-3.1.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-server-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/token-provider-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-native-client-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-nfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-native-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-httpfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-client-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-jobclient-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-shuffle-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-examples-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-hs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-nativetask-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-jobclient-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-hs-plugins-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-core-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-app-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-module-jaxb-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-csv-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-math-2.2.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/guice-4.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/json-io-2.5.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jcodings-1.0.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/java-util-1.9.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-jaxrs-json-provider-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/HikariCP-java7-2.4.12.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-prefix-tree-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/mssql-jdbc-6.2.1.jre7.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-common-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/joni-2.1.2.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/findbugs-annotations-1.3.9-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-server-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-el-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jamon-runtime-2.4.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/guice-servlet-4.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/disruptor-3.3.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jersey-client-1.19.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/htrace-core-3.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jsp-api-2.1-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/javax.inject-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-jaxrs-base-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-procedure-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-protocol-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/geronimo-jcache_1.0_spec-1.0-alpha-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/fst-2.50.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-hadoop2-compat-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-hadoop-compat-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-client-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-annotations-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jasper-compiler-5.5.23.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/aopalliance-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/metrics-core-3.0.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jasper-runtime-5.5.23.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jsp-2.1-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/ehcache-3.3.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/servlet-api-2.5-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/metrics-core-2.2.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jersey-guice-1.19.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-httpclient-3.1.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-registry-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-router-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-resourcemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-tests-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-applicationhistoryservice-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-sharedcachemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-api-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-nodemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-hbase-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-hbase-tests-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timeline-pluginstorage-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-web-proxy-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-applications-unmanaged-am-launcher-3.0.0.jar
STARTUP_MSG:   build = https://git-wip-us.apache.org/repos/asf/hadoop.git -r c25427ceca461ee979d30edd7a4b0f50718e6533; compiled by 'andrew' on 2017-12-08T19:16Z
STARTUP_MSG:   java = 1.8.0_151
************************************************************/
2018-02-04 14:28:05,053 INFO namenode.NameNode: registered UNIX signal handlers for [TERM, HUP, INT]
2018-02-04 14:28:05,059 INFO namenode.NameNode: createNameNode [-format, demo]
Formatting using clusterid: CID-fa61baf9-c459-4965-ba8d-cd3afdee36f1
2018-02-04 14:28:06,105 INFO namenode.FSEditLog: Edit logging is async:true
2018-02-04 14:28:06,125 INFO namenode.FSNamesystem: KeyProvider: null
2018-02-04 14:28:06,126 INFO namenode.FSNamesystem: fsLock is fair: true
2018-02-04 14:28:06,130 INFO namenode.FSNamesystem: Detailed lock hold time metrics enabled: false
2018-02-04 14:28:06,136 INFO namenode.FSNamesystem: fsOwner             = root (auth:SIMPLE)
2018-02-04 14:28:06,136 INFO namenode.FSNamesystem: supergroup          = supergroup
2018-02-04 14:28:06,136 INFO namenode.FSNamesystem: isPermissionEnabled = true
2018-02-04 14:28:06,136 INFO namenode.FSNamesystem: HA Enabled: false
2018-02-04 14:28:06,208 INFO common.Util: dfs.datanode.fileio.profiling.sampling.percentage set to 0. Disabling file IO profiling
2018-02-04 14:28:06,236 INFO blockmanagement.DatanodeManager: dfs.block.invalidate.limit: configured=1000, counted=60, effected=1000
2018-02-04 14:28:06,236 INFO blockmanagement.DatanodeManager: dfs.namenode.datanode.registration.ip-hostname-check=true
2018-02-04 14:28:06,252 INFO blockmanagement.BlockManager: dfs.namenode.startup.delay.block.deletion.sec is set to 000:00:00:00.000
2018-02-04 14:28:06,254 INFO blockmanagement.BlockManager: The block deletion will start around 2018 Feb 04 14:28:06
2018-02-04 14:28:06,258 INFO util.GSet: Computing capacity for map BlocksMap
2018-02-04 14:28:06,258 INFO util.GSet: VM type       = 64-bit
2018-02-04 14:28:06,261 INFO util.GSet: 2.0% max memory 483.4 MB = 9.7 MB
2018-02-04 14:28:06,261 INFO util.GSet: capacity      = 2^20 = 1048576 entries
2018-02-04 14:28:06,296 INFO blockmanagement.BlockManager: dfs.block.access.token.enable = false
2018-02-04 14:28:06,303 INFO Configuration.deprecation: No unit for dfs.namenode.safemode.extension(30000) assuming MILLISECONDS
2018-02-04 14:28:06,303 INFO blockmanagement.BlockManagerSafeMode: dfs.namenode.safemode.threshold-pct = 0.9990000128746033
2018-02-04 14:28:06,303 INFO blockmanagement.BlockManagerSafeMode: dfs.namenode.safemode.min.datanodes = 0
2018-02-04 14:28:06,304 INFO blockmanagement.BlockManagerSafeMode: dfs.namenode.safemode.extension = 30000
2018-02-04 14:28:06,304 INFO blockmanagement.BlockManager: defaultReplication         = 3
2018-02-04 14:28:06,304 INFO blockmanagement.BlockManager: maxReplication             = 512
2018-02-04 14:28:06,304 INFO blockmanagement.BlockManager: minReplication             = 1
2018-02-04 14:28:06,304 INFO blockmanagement.BlockManager: maxReplicationStreams      = 2
2018-02-04 14:28:06,304 INFO blockmanagement.BlockManager: redundancyRecheckInterval  = 3000ms
2018-02-04 14:28:06,304 INFO blockmanagement.BlockManager: encryptDataTransfer        = false
2018-02-04 14:28:06,304 INFO blockmanagement.BlockManager: maxNumBlocksToLog          = 1000
2018-02-04 14:28:06,382 INFO util.GSet: Computing capacity for map INodeMap
2018-02-04 14:28:06,382 INFO util.GSet: VM type       = 64-bit
2018-02-04 14:28:06,383 INFO util.GSet: 1.0% max memory 483.4 MB = 4.8 MB
2018-02-04 14:28:06,384 INFO util.GSet: capacity      = 2^19 = 524288 entries
2018-02-04 14:28:06,384 INFO namenode.FSDirectory: ACLs enabled? false
2018-02-04 14:28:06,385 INFO namenode.FSDirectory: POSIX ACL inheritance enabled? true
2018-02-04 14:28:06,385 INFO namenode.FSDirectory: XAttrs enabled? true
2018-02-04 14:28:06,385 INFO namenode.NameNode: Caching file names occurring more than 10 times
2018-02-04 14:28:06,393 INFO snapshot.SnapshotManager: Loaded config captureOpenFiles: false, skipCaptureAccessTimeOnlyChange: false, snapshotDiffAllowSnapRootDescendant: true
2018-02-04 14:28:06,398 INFO util.GSet: Computing capacity for map cachedBlocks
2018-02-04 14:28:06,399 INFO util.GSet: VM type       = 64-bit
2018-02-04 14:28:06,399 INFO util.GSet: 0.25% max memory 483.4 MB = 1.2 MB
2018-02-04 14:28:06,399 INFO util.GSet: capacity      = 2^17 = 131072 entries
2018-02-04 14:28:06,412 INFO metrics.TopMetrics: NNTop conf: dfs.namenode.top.window.num.buckets = 10
2018-02-04 14:28:06,412 INFO metrics.TopMetrics: NNTop conf: dfs.namenode.top.num.users = 10
2018-02-04 14:28:06,413 INFO metrics.TopMetrics: NNTop conf: dfs.namenode.top.windows.minutes = 1,5,25
2018-02-04 14:28:06,426 INFO namenode.FSNamesystem: Retry cache on namenode is enabled
2018-02-04 14:28:06,429 INFO namenode.FSNamesystem: Retry cache will use 0.03 of total heap and retry cache entry expiry time is 600000 millis
2018-02-04 14:28:06,436 INFO util.GSet: Computing capacity for map NameNodeRetryCache
2018-02-04 14:28:06,436 INFO util.GSet: VM type       = 64-bit
2018-02-04 14:28:06,436 INFO util.GSet: 0.029999999329447746% max memory 483.4 MB = 148.5 KB
2018-02-04 14:28:06,436 INFO util.GSet: capacity      = 2^14 = 16384 entries
Re-format filesystem in Storage Directory /tmp/hadoop-root/dfs/name ? (Y or N) y
2018-02-04 14:28:10,051 INFO namenode.FSImage: Allocated new BlockPoolId: BP-1694953607-172.18.0.3-1517754490039
2018-02-04 14:28:10,051 INFO common.Storage: Will remove files: [/tmp/hadoop-root/dfs/name/current/fsimage_0000000000000000000.md5, /tmp/hadoop-root/dfs/name/current/seen_txid, /tmp/hadoop-root/dfs/name/current/VERSION, /tmp/hadoop-root/dfs/name/current/fsimage_0000000000000000000]
2018-02-04 14:28:10,089 INFO common.Storage: Storage directory /tmp/hadoop-root/dfs/name has been successfully formatted.
2018-02-04 14:28:10,106 INFO namenode.FSImageFormatProtobuf: Saving image file /tmp/hadoop-root/dfs/name/current/fsimage.ckpt_0000000000000000000 using no compression
2018-02-04 14:28:10,238 INFO namenode.FSImageFormatProtobuf: Image file /tmp/hadoop-root/dfs/name/current/fsimage.ckpt_0000000000000000000 of size 389 bytes saved in 0 seconds.
2018-02-04 14:28:10,252 INFO namenode.NNStorageRetentionManager: Going to retain 1 images with txid >= 0
2018-02-04 14:28:10,260 INFO namenode.NameNode: SHUTDOWN_MSG: 
/************************************************************
SHUTDOWN_MSG: Shutting down NameNode at 480e60ebed55/172.18.0.3
************************************************************/
```

```
root@480e60ebed55:/# hadoop-3.0.0/bin/hdfs namenode
2018-02-04 14:29:13,049 INFO namenode.NameNode: STARTUP_MSG: 
/************************************************************
STARTUP_MSG: Starting NameNode
STARTUP_MSG:   host = 480e60ebed55/172.18.0.3
STARTUP_MSG:   args = []
STARTUP_MSG:   version = 3.0.0
STARTUP_MSG:   classpath = /hadoop-3.0.0/etc/hadoop:/hadoop-3.0.0/share/hadoop/common/lib/jersey-core-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/netty-3.10.5.Final.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-server-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-webapp-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-io-2.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/zookeeper-3.4.9.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-common-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/avro-1.7.7.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-jaxrs-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsr311-api-1.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/guava-11.0.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-asn1-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-databind-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-beanutils-1.9.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-http-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-core-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-lang-2.6.jar:/hadoop-3.0.0/share/hadoop/common/lib/httpcore-4.4.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-json-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/json-smart-2.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/snappy-java-1.0.5.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-admin-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-cli-1.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/jettison-1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/httpclient-4.5.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/jcip-annotations-1.0-1.jar:/hadoop-3.0.0/share/hadoop/common/lib/xz-1.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-server-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-client-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-xml-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-codec-1.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/asm-5.0.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/woodstox-core-5.0.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-io-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/jersey-servlet-1.19.jar:/hadoop-3.0.0/share/hadoop/common/lib/jaxb-api-2.2.11.jar:/hadoop-3.0.0/share/hadoop/common/lib/accessors-smart-1.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/protobuf-java-2.5.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-recipes-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-util-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/paranamer-2.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/hadoop-annotations-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-math3-3.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-config-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-core-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-xdr-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/nimbus-jose-jwt-4.41.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jaxb-impl-2.2.3-1.jar:/hadoop-3.0.0/share/hadoop/common/lib/gson-2.2.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-lang3-3.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-servlet-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/re2j-1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-collections-3.2.2.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-logging-1.1.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/curator-framework-2.12.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-simplekdc-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-net-3.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-compress-1.4.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/slf4j-api-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-xc-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsch-0.1.54.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerby-pkix-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-crypto-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jul-to-slf4j-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/log4j-1.2.17.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/junit-4.11.jar:/hadoop-3.0.0/share/hadoop/common/lib/mockito-all-1.8.5.jar:/hadoop-3.0.0/share/hadoop/common/lib/metrics-core-3.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/stax2-api-3.1.4.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsp-api-2.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/slf4j-log4j12-1.7.25.jar:/hadoop-3.0.0/share/hadoop/common/lib/htrace-core4-4.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/common/lib/hadoop-auth-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jsr305-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-mapper-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/common/lib/commons-configuration2-2.1.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jackson-core-2.7.8.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-client-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/hamcrest-core-1.3.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-identity-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/jetty-security-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/common/lib/javax.servlet-api-3.1.0.jar:/hadoop-3.0.0/share/hadoop/common/lib/kerb-server-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/lib/token-provider-1.0.1.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-nfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-common-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-kms-3.0.0.jar:/hadoop-3.0.0/share/hadoop/common/hadoop-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-core-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/netty-3.10.5.Final.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-server-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-webapp-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-io-2.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/zookeeper-3.4.9.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-common-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/avro-1.7.7.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-jaxrs-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsr311-api-1.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/guava-11.0.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-asn1-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-databind-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-beanutils-1.9.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-http-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-core-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-lang-2.6.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/httpcore-4.4.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-json-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/leveldbjni-all-1.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-daemon-1.0.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/json-smart-2.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/snappy-java-1.0.5.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-admin-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-cli-1.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jettison-1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/httpclient-4.5.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jcip-annotations-1.0-1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/xz-1.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-server-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-client-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-xml-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-codec-1.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/asm-5.0.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/woodstox-core-5.0.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-util-ajax-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-io-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jersey-servlet-1.19.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jaxb-api-2.2.11.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/accessors-smart-1.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/protobuf-java-2.5.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-recipes-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-util-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/paranamer-2.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/hadoop-annotations-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/okio-1.4.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-math3-3.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-config-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-core-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-xdr-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/nimbus-jose-jwt-4.41.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jaxb-impl-2.2.3-1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/json-simple-1.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/gson-2.2.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-lang3-3.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-servlet-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/re2j-1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-collections-3.2.2.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-logging-1.1.3.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/curator-framework-2.12.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-simplekdc-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-net-3.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-compress-1.4.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-xc-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsch-0.1.54.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerby-pkix-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-crypto-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/log4j-1.2.17.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-util-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/stax2-api-3.1.4.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/htrace-core4-4.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/hadoop-auth-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jsr305-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-mapper-asl-1.9.13.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/commons-configuration2-2.1.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jackson-core-2.7.8.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-client-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/netty-all-4.0.23.Final.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-identity-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/jetty-security-9.3.19.v20170502.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/okhttp-2.4.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/javax.servlet-api-3.1.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/kerb-server-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/lib/token-provider-1.0.1.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-native-client-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-nfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-native-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-httpfs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/hdfs/hadoop-hdfs-client-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-jobclient-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-shuffle-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-examples-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-hs-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-nativetask-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-jobclient-3.0.0-tests.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-hs-plugins-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-core-3.0.0.jar:/hadoop-3.0.0/share/hadoop/mapreduce/hadoop-mapreduce-client-app-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-module-jaxb-annotations-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-csv-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-math-2.2.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/guice-4.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/json-io-2.5.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jcodings-1.0.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/java-util-1.9.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-jaxrs-json-provider-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/HikariCP-java7-2.4.12.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-prefix-tree-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/mssql-jdbc-6.2.1.jre7.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-common-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/joni-2.1.2.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/findbugs-annotations-1.3.9-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-server-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-el-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jamon-runtime-2.4.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/guice-servlet-4.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/disruptor-3.3.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jersey-client-1.19.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/htrace-core-3.1.0-incubating.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jsp-api-2.1-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/javax.inject-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jackson-jaxrs-base-2.7.8.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-procedure-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-protocol-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/geronimo-jcache_1.0_spec-1.0-alpha-1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/fst-2.50.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-hadoop2-compat-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-hadoop-compat-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-client-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/hbase-annotations-1.2.6.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jasper-compiler-5.5.23.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/aopalliance-1.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/metrics-core-3.0.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jasper-runtime-5.5.23.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jsp-2.1-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/ehcache-3.3.1.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/servlet-api-2.5-6.1.14.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/metrics-core-2.2.0.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/jersey-guice-1.19.jar:/hadoop-3.0.0/share/hadoop/yarn/lib/commons-httpclient-3.1.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-registry-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-applications-distributedshell-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-router-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-resourcemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-tests-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-applicationhistoryservice-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-sharedcachemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-api-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-nodemanager-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-hbase-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timelineservice-hbase-tests-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-common-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-client-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-timeline-pluginstorage-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-server-web-proxy-3.0.0.jar:/hadoop-3.0.0/share/hadoop/yarn/hadoop-yarn-applications-unmanaged-am-launcher-3.0.0.jar
STARTUP_MSG:   build = https://git-wip-us.apache.org/repos/asf/hadoop.git -r c25427ceca461ee979d30edd7a4b0f50718e6533; compiled by 'andrew' on 2017-12-08T19:16Z
STARTUP_MSG:   java = 1.8.0_151
************************************************************/
2018-02-04 14:29:13,069 INFO namenode.NameNode: registered UNIX signal handlers for [TERM, HUP, INT]
2018-02-04 14:29:13,081 INFO namenode.NameNode: createNameNode []
2018-02-04 14:29:13,566 INFO beanutils.FluentPropertyBeanIntrospector: Error when creating PropertyDescriptor for public final void org.apache.commons.configuration2.AbstractConfiguration.setProperty(java.lang.String,java.lang.Object)! Ignoring this property.
2018-02-04 14:29:13,645 INFO impl.MetricsConfig: loaded properties from hadoop-metrics2.properties
2018-02-04 14:29:13,804 INFO impl.MetricsSystemImpl: Scheduled Metric snapshot period at 10 second(s).
2018-02-04 14:29:13,804 INFO impl.MetricsSystemImpl: NameNode metrics system started
2018-02-04 14:29:13,851 INFO namenode.NameNode: fs.defaultFS is hdfs://127.0.0.1:9000
2018-02-04 14:29:13,861 INFO namenode.NameNode: Clients are to use 127.0.0.1:9000 to access this namenode/service.
2018-02-04 14:29:14,162 INFO util.JvmPauseMonitor: Starting JVM pause monitor
2018-02-04 14:29:14,291 INFO hdfs.DFSUtil: Starting Web-server for hdfs at: http://0.0.0.0:9870
2018-02-04 14:29:14,320 INFO util.log: Logging initialized @2040ms
2018-02-04 14:29:14,485 INFO server.AuthenticationFilter: Unable to initialize FileSignerSecretProvider, falling back to use random secrets.
2018-02-04 14:29:14,513 INFO http.HttpRequestLog: Http request log for http.requests.namenode is not defined
2018-02-04 14:29:14,533 INFO http.HttpServer2: Added global filter 'safety' (class=org.apache.hadoop.http.HttpServer2$QuotingInputFilter)
2018-02-04 14:29:14,545 INFO http.HttpServer2: Added filter static_user_filter (class=org.apache.hadoop.http.lib.StaticUserWebFilter$StaticUserFilter) to context hdfs
2018-02-04 14:29:14,546 INFO http.HttpServer2: Added filter static_user_filter (class=org.apache.hadoop.http.lib.StaticUserWebFilter$StaticUserFilter) to context static
2018-02-04 14:29:14,546 INFO http.HttpServer2: Added filter static_user_filter (class=org.apache.hadoop.http.lib.StaticUserWebFilter$StaticUserFilter) to context logs
2018-02-04 14:29:14,607 INFO http.HttpServer2: Added filter 'org.apache.hadoop.hdfs.web.AuthFilter' (class=org.apache.hadoop.hdfs.web.AuthFilter)
2018-02-04 14:29:14,613 INFO http.HttpServer2: addJerseyResourcePackage: packageName=org.apache.hadoop.hdfs.server.namenode.web.resources;org.apache.hadoop.hdfs.web.resources, pathSpec=/webhdfs/v1/*
2018-02-04 14:29:14,641 INFO http.HttpServer2: Jetty bound to port 9870
2018-02-04 14:29:14,643 INFO server.Server: jetty-9.3.19.v20170502
2018-02-04 14:29:14,751 INFO handler.ContextHandler: Started o.e.j.s.ServletContextHandler@5b7a7f33{/logs,file:///hadoop-3.0.0/logs/,AVAILABLE}
2018-02-04 14:29:14,764 INFO handler.ContextHandler: Started o.e.j.s.ServletContextHandler@5c7933ad{/static,file:///hadoop-3.0.0/share/hadoop/hdfs/webapps/static/,AVAILABLE}
2018-02-04 14:29:15,003 INFO handler.ContextHandler: Started o.e.j.w.WebAppContext@51a9ad5e{/,file:///hadoop-3.0.0/share/hadoop/hdfs/webapps/hdfs/,AVAILABLE}{/hdfs}
2018-02-04 14:29:15,015 INFO server.AbstractConnector: Started ServerConnector@65934e62{HTTP/1.1,[http/1.1]}{0.0.0.0:9870}
2018-02-04 14:29:15,015 INFO server.Server: Started @2738ms
2018-02-04 14:29:15,519 WARN namenode.FSNamesystem: Only one image storage directory (dfs.namenode.name.dir) configured. Beware of data loss due to lack of redundant storage directories!
2018-02-04 14:29:15,521 WARN namenode.FSNamesystem: Only one namespace edits storage directory (dfs.namenode.edits.dir) configured. Beware of data loss due to lack of redundant storage directories!
2018-02-04 14:29:15,647 INFO namenode.FSEditLog: Edit logging is async:true
2018-02-04 14:29:15,697 INFO namenode.FSNamesystem: KeyProvider: null
2018-02-04 14:29:15,700 INFO namenode.FSNamesystem: fsLock is fair: true
2018-02-04 14:29:15,703 INFO namenode.FSNamesystem: Detailed lock hold time metrics enabled: false
2018-02-04 14:29:15,720 INFO namenode.FSNamesystem: fsOwner             = root (auth:SIMPLE)
2018-02-04 14:29:15,720 INFO namenode.FSNamesystem: supergroup          = supergroup
2018-02-04 14:29:15,720 INFO namenode.FSNamesystem: isPermissionEnabled = true
2018-02-04 14:29:15,720 INFO namenode.FSNamesystem: HA Enabled: false
2018-02-04 14:29:15,773 INFO common.Util: dfs.datanode.fileio.profiling.sampling.percentage set to 0. Disabling file IO profiling
2018-02-04 14:29:15,786 INFO blockmanagement.DatanodeManager: dfs.block.invalidate.limit: configured=1000, counted=60, effected=1000
2018-02-04 14:29:15,786 INFO blockmanagement.DatanodeManager: dfs.namenode.datanode.registration.ip-hostname-check=true
2018-02-04 14:29:15,794 INFO blockmanagement.BlockManager: dfs.namenode.startup.delay.block.deletion.sec is set to 000:00:00:00.000
2018-02-04 14:29:15,794 INFO blockmanagement.BlockManager: The block deletion will start around 2018 Feb 04 14:29:15
2018-02-04 14:29:15,796 INFO util.GSet: Computing capacity for map BlocksMap
2018-02-04 14:29:15,797 INFO util.GSet: VM type       = 64-bit
2018-02-04 14:29:15,799 INFO util.GSet: 2.0% max memory 483.4 MB = 9.7 MB
2018-02-04 14:29:15,799 INFO util.GSet: capacity      = 2^20 = 1048576 entries
2018-02-04 14:29:15,810 INFO blockmanagement.BlockManager: dfs.block.access.token.enable = false
2018-02-04 14:29:15,822 INFO Configuration.deprecation: No unit for dfs.namenode.safemode.extension(30000) assuming MILLISECONDS
2018-02-04 14:29:15,823 INFO blockmanagement.BlockManagerSafeMode: dfs.namenode.safemode.threshold-pct = 0.9990000128746033
2018-02-04 14:29:15,823 INFO blockmanagement.BlockManagerSafeMode: dfs.namenode.safemode.min.datanodes = 0
2018-02-04 14:29:15,823 INFO blockmanagement.BlockManagerSafeMode: dfs.namenode.safemode.extension = 30000
2018-02-04 14:29:15,823 INFO blockmanagement.BlockManager: defaultReplication         = 3
2018-02-04 14:29:15,823 INFO blockmanagement.BlockManager: maxReplication             = 512
2018-02-04 14:29:15,823 INFO blockmanagement.BlockManager: minReplication             = 1
2018-02-04 14:29:15,823 INFO blockmanagement.BlockManager: maxReplicationStreams      = 2
2018-02-04 14:29:15,824 INFO blockmanagement.BlockManager: redundancyRecheckInterval  = 3000ms
2018-02-04 14:29:15,824 INFO blockmanagement.BlockManager: encryptDataTransfer        = false
2018-02-04 14:29:15,824 INFO blockmanagement.BlockManager: maxNumBlocksToLog          = 1000
2018-02-04 14:29:15,911 INFO util.GSet: Computing capacity for map INodeMap
2018-02-04 14:29:15,911 INFO util.GSet: VM type       = 64-bit
2018-02-04 14:29:15,914 INFO util.GSet: 1.0% max memory 483.4 MB = 4.8 MB
2018-02-04 14:29:15,914 INFO util.GSet: capacity      = 2^19 = 524288 entries
2018-02-04 14:29:15,915 INFO namenode.FSDirectory: ACLs enabled? false
2018-02-04 14:29:15,915 INFO namenode.FSDirectory: POSIX ACL inheritance enabled? true
2018-02-04 14:29:15,916 INFO namenode.FSDirectory: XAttrs enabled? true
2018-02-04 14:29:15,923 INFO namenode.NameNode: Caching file names occurring more than 10 times
2018-02-04 14:29:15,932 INFO snapshot.SnapshotManager: Loaded config captureOpenFiles: false, skipCaptureAccessTimeOnlyChange: false, snapshotDiffAllowSnapRootDescendant: true
2018-02-04 14:29:15,944 INFO util.GSet: Computing capacity for map cachedBlocks
2018-02-04 14:29:15,944 INFO util.GSet: VM type       = 64-bit
2018-02-04 14:29:15,945 INFO util.GSet: 0.25% max memory 483.4 MB = 1.2 MB
2018-02-04 14:29:15,945 INFO util.GSet: capacity      = 2^17 = 131072 entries
2018-02-04 14:29:15,962 INFO metrics.TopMetrics: NNTop conf: dfs.namenode.top.window.num.buckets = 10
2018-02-04 14:29:15,963 INFO metrics.TopMetrics: NNTop conf: dfs.namenode.top.num.users = 10
2018-02-04 14:29:15,963 INFO metrics.TopMetrics: NNTop conf: dfs.namenode.top.windows.minutes = 1,5,25
2018-02-04 14:29:15,971 INFO namenode.FSNamesystem: Retry cache on namenode is enabled
2018-02-04 14:29:15,972 INFO namenode.FSNamesystem: Retry cache will use 0.03 of total heap and retry cache entry expiry time is 600000 millis
2018-02-04 14:29:15,978 INFO util.GSet: Computing capacity for map NameNodeRetryCache
2018-02-04 14:29:15,980 INFO util.GSet: VM type       = 64-bit
2018-02-04 14:29:15,981 INFO util.GSet: 0.029999999329447746% max memory 483.4 MB = 148.5 KB
2018-02-04 14:29:15,981 INFO util.GSet: capacity      = 2^14 = 16384 entries
2018-02-04 14:29:15,998 INFO common.Storage: Lock on /tmp/hadoop-root/dfs/name/in_use.lock acquired by nodename 1085@480e60ebed55
2018-02-04 14:29:16,046 INFO namenode.FileJournalManager: Recovering unfinalized segments in /tmp/hadoop-root/dfs/name/current
2018-02-04 14:29:16,047 INFO namenode.FSImage: No edit log streams selected.
2018-02-04 14:29:16,047 INFO namenode.FSImage: Planning to load image: FSImageFile(file=/tmp/hadoop-root/dfs/name/current/fsimage_0000000000000000000, cpktTxId=0000000000000000000)
2018-02-04 14:29:16,178 INFO namenode.ErasureCodingPolicyManager: Enable the erasure coding policy RS-6-3-1024k
2018-02-04 14:29:16,191 INFO namenode.FSImageFormatPBINode: Loading 1 INodes.
2018-02-04 14:29:16,272 INFO namenode.FSImageFormatProtobuf: Loaded FSImage in 0 seconds.
2018-02-04 14:29:16,277 INFO namenode.FSImage: Loaded image for txid 0 from /tmp/hadoop-root/dfs/name/current/fsimage_0000000000000000000
2018-02-04 14:29:16,281 INFO namenode.FSNamesystem: Need to save fs image? false (staleImage=false, haEnabled=false, isRollingUpgrade=false)
2018-02-04 14:29:16,290 INFO namenode.FSEditLog: Starting log segment at 1
2018-02-04 14:29:16,496 INFO namenode.NameCache: initialized with 0 entries 0 lookups
2018-02-04 14:29:16,496 INFO namenode.FSNamesystem: Finished loading FSImage in 506 msecs
2018-02-04 14:29:16,771 INFO namenode.NameNode: RPC server is binding to localhost:9000
2018-02-04 14:29:16,783 INFO ipc.CallQueueManager: Using callQueue: class java.util.concurrent.LinkedBlockingQueue queueCapacity: 1000 scheduler: class org.apache.hadoop.ipc.DefaultRpcScheduler
2018-02-04 14:29:16,799 INFO ipc.Server: Starting Socket Reader #1 for port 9000
2018-02-04 14:29:17,108 INFO namenode.FSNamesystem: Registered FSNamesystemState, ReplicatedBlocksState and ECBlockGroupsState MBeans.
2018-02-04 14:29:17,141 INFO namenode.LeaseManager: Number of blocks under construction: 0
2018-02-04 14:29:17,171 INFO blockmanagement.BlockManager: initializing replication queues
2018-02-04 14:29:17,178 INFO hdfs.StateChange: STATE* Leaving safe mode after 0 secs
2018-02-04 14:29:17,179 INFO hdfs.StateChange: STATE* Network topology has 0 racks and 0 datanodes
2018-02-04 14:29:17,179 INFO hdfs.StateChange: STATE* UnderReplicatedBlocks has 0 blocks
2018-02-04 14:29:17,233 INFO blockmanagement.BlockManager: Total number of blocks            = 0
2018-02-04 14:29:17,235 INFO blockmanagement.BlockManager: Number of invalid blocks          = 0
2018-02-04 14:29:17,236 INFO blockmanagement.BlockManager: Number of under-replicated blocks = 0
2018-02-04 14:29:17,236 INFO blockmanagement.BlockManager: Number of  over-replicated blocks = 0
2018-02-04 14:29:17,236 INFO blockmanagement.BlockManager: Number of blocks being written    = 0
2018-02-04 14:29:17,236 INFO hdfs.StateChange: STATE* Replication Queue initialization scan for invalid, over- and under-replicated blocks completed in 55 msec
2018-02-04 14:29:17,319 INFO ipc.Server: IPC Server Responder: starting
2018-02-04 14:29:17,325 INFO ipc.Server: IPC Server listener on 9000: starting
2018-02-04 14:29:17,348 INFO namenode.NameNode: NameNode RPC up at: localhost/127.0.0.1:9000
2018-02-04 14:29:17,355 INFO namenode.FSNamesystem: Starting services required for active state
2018-02-04 14:29:17,356 INFO namenode.FSDirectory: Initializing quota with 4 thread(s)
2018-02-04 14:29:17,372 INFO namenode.FSDirectory: Quota initialization completed in 16 milliseconds
name space=1
storage space=0
storage types=RAM_DISK=0, SSD=0, DISK=0, ARCHIVE=0
2018-02-04 14:29:17,396 INFO blockmanagement.CacheReplicationMonitor: Starting CacheReplicationMonitor with interval 30000 milliseconds
^C2018-02-04 14:30:51,680 ERROR namenode.NameNode: RECEIVED SIGNAL 2: SIGINT
2018-02-04 14:30:51,682 INFO namenode.NameNode: SHUTDOWN_MSG: 
/************************************************************
SHUTDOWN_MSG: Shutting down NameNode at 480e60ebed55/172.18.0.3
************************************************************/
```
