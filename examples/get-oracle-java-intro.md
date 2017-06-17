## Introduction

### Maven

* Browse __Apache Maven Dist__

https://archive.apache.org/dist/maven/maven-3/

__清华大学开源软件镜像站Maven镜像__

https://mirrors.tuna.tsinghua.edu.cn/apache/maven/maven-3/

* Comand __curl__ into File Server

For Dockernized File Server, Refer to [gofileserver.md](./gofileserver.md)

[vagrant@localhost ~]$ dest_dir=/Users/fanhongling/Downloads/99-mirror/https0x3A0x2F0x2Farchive.apache.org0x2Fdist/maven/maven-30x2F3.5.00x2Fbinaries; mkdir -p $dest_dir && pushd $dest_dir; curl --retry 10 --connect-timeout 60 -jkSLO -C- https://mirrors.tuna.tsinghua.edu.cn/apache/maven/maven-3/3.5.0/binaries/apache-maven-3.5.0-bin.tar.gz; popd
~/Downloads/99-mirror/https0x3A0x2F0x2Farchive.apache.org0x2Fdist/maven/maven-30x2F3.5.00x2Fbinaries ~/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 8334k  100 8334k    0     0   315k      0  0:00:26  0:00:26 --:--:--  544k
~/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples

### Gradle

* Browse

https://services.gradle.org/distributions/

* Alternative command __wget__ into File Server

For Dockernized File Server, Refer to [gofileserver.md](./gofileserver.md)

[vagrant@localhost ~]$ dest_dir=/Users/fanhongling/Downloads/99-mirror/https0x3A0x2F0x2Fservices.gradle.org0x2Fdistributions; mkdir -p $dest_dir && wget --no-clobber --tries=10 -c --no-check-certificate --no-cookies https://services.gradle.org/distributions/gradle-2.14-bin.zip --directory-prefix=$dest_dir
[vagrant@localhost 99-mirror]$ dest_dir=/Users/fanhongling/Downloads/99-mirror/https0x3A0x2F0x2Fservices.gradle.org0x2Fdistributions; mkdir -p $dest_dir && wget --no-clobber --tries=10 -c --no-check-certificate --no-cookies https://services.gradle.org/distributions/gradle-2.14-bin.zip --directory-prefix=$dest_dir
--2017-06-16 17:21:18--  https://services.gradle.org/distributions/gradle-2.14-bin.zip
正在解析主机 services.gradle.org (services.gradle.org)... 104.16.172.166, 104.16.174.166, 104.16.173.166, ...
正在连接 services.gradle.org (services.gradle.org)|104.16.172.166|:443... 已连接。
已发出 HTTP 请求，正在等待回应... 301 Moved Permanently
位置：https://downloads.gradle.org/distributions/gradle-2.14-bin.zip [跟随至新的 URL]
--2017-06-16 17:21:19--  https://downloads.gradle.org/distributions/gradle-2.14-bin.zip
正在解析主机 downloads.gradle.org (downloads.gradle.org)... 104.16.171.166, 104.16.174.166, 104.16.172.166, ...
正在连接 downloads.gradle.org (downloads.gradle.org)|104.16.171.166|:443... 已连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：45033194 (43M) [application/zip]
正在保存至: “/Users/fanhongling/Downloads/99-mirror/https0x3A0x2F0x2Fservices.gradle.org0x2Fdistributions/gradle-2.14-bin.zip”

gradle-2.14-bin.zip                    100%[=========================================================================>]  42.95M   297KB/s    in 2m 0s   

2017-06-16 17:23:21 (365 KB/s) - 已保存 “/Users/fanhongling/Downloads/99-mirror/https0x3A0x2F0x2Fservices.gradle.org0x2Fdistributions/gradle-2.14-bin.zip” [45033194/45033194])
