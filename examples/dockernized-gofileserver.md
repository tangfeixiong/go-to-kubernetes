### Run fileserver in none Internet

Image
```
[vagrant@localhost ~]$ docker pull docker.io/tangfeixiong/gofileserver
Using default tag: latest
Trying to pull repository docker.io/tangfeixiong/gofileserver ... 
latest: Pulling from docker.io/tangfeixiong/gofileserver
c0cb142e4345: Already exists 
bd94e5529d32: Already exists 
Digest: sha256:b05cb512410b314f3aae3ac1932faec2c8fc892221cc773778f4efdabd3684c8
Status: Image is up to date for docker.io/tangfeixiong/gofileserver:latest
```

Run
```
[vagrant@localhost alpine-gofileserver]$ docker run -v /Users/fanhongling/Downloads:/mnt:ro -p 48080:48080 --name=gofileserver -d docker.io/tangfeixiong/gofileserver
b1e178fa584eea379eb60a2f0947ec21247aba96378e8e073d82ed5d393a301c
```

Check
```
[vagrant@localhost alpine-gofileserver]$ curl -L localhost:48080/
<pre>
<a href="99-mirror/">99-mirror/</a>
<a href="go-kubernetes/">go-kubernetes/</a>
<a href="go-openshift/">go-openshift/</a>
<a href="workspace/">workspace/</a>
</pre>
```

Serve

* RPM - CentOS YUM repository: _curl -L localhost:48080/99-mirror/http0x3A0x2F0x2Fmirror.centos.org0x2Fcentos/7/_
* Java - JDK, JRE: _curl -L localhost:48080/99-mirror/http0x3A0x2F0x2Fdownload.oracle.com/_
* Maven packages: _curl -L localhost:48080/99-mirror/0x2Em20x2Fwrapper/_
* Gradle: _curl -L localhost:48080/99-mirror/0x2Egradle0x2Fwrapper/_
* Maven repository: _curl -L localhost:48080/99-mirror/0x2Em20x2Frepository/_
* Nodejs repository: _curl -L localhost:48080/99-mirror/0x2Enpm/_
* ...
