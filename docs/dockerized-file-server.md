## HTTP-based file server serving offline docker image archive and YUM/APT repository mirror

### Table of contents

1. CentOS 7
1. Fedora 26
1. Ubuntu 16.04

### CentOS 7

Pull
```
[vagrant@kubedev-10-64-33-82 ~]$ docker pull docker.io/tangfeixiong/gofileserver
Using default tag: latest
latest: Pulling from tangfeixiong/gofileserver
c0cb142e4345: Pull complete 
bd94e5529d32: Pull complete 
Digest: sha256:b05cb512410b314f3aae3ac1932faec2c8fc892221cc773778f4efdabd3684c8
Status: Downloaded newer image for tangfeixiong/gofileserver:latest
```

or save as archive for offline work
```
$ docker save -o docker.io0x2Ftangfeixiong0x2Fgofileserver.tar
```

conveniently load in offline environment
```
$ docker load -i docker.io0x2Ftangfeixiong0x2Fgofileserver.tar
```

Run
```
[vagrant@kubedev-10-64-33-82 ~]$ docker run -d --name=gofileserver -p 48080:48080 -v /:/mnt docker.io/tangfeixiong/gofileserver
daeece8d67249922dce96508e60bfc48c218b0cbdda99a20f5f0354580b69e14
```

Or
```
[vagrant@kubedev-10-64-33-82 ~]$ docker run -d --name=gofileserver -p 48080:48080 --restart=always -v /:/mnt docker.io/tangfeixiong/gofileserver
```

Or
```
[vagrant@kubedev-10-64-33-82 ~]$ docker update --restart=always gofileserver
gofileserver
```

For example, my kubernetes apt/yum repository mirror
```
[vagrant@kubedev-172-17-4-59 ~]$ curl -L http://172.17.4.59:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/
<pre>
<a href="apt/">apt/</a>
<a href="yum/">yum/</a>
</pre>
```

and CentOS extras, updates even cloud, pass and storage repository 
```
[vagrant@localhost rook]$ curl -jkSL http://10.64.33.82:48080/windows10_drive_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/
<pre>
<a href="RPM-GPG-KEY-CentOS-7">RPM-GPG-KEY-CentOS-7</a>
<a href="atomic/">atomic/</a>
<a href="centosplus/">centosplus/</a>
<a href="cloud/">cloud/</a>
<a href="configmanagement/">configmanagement/</a>
<a href="cr/">cr/</a>
<a href="dotnet/">dotnet/</a>
<a href="extras/">extras/</a>
<a href="fasttrack/">fasttrack/</a>
<a href="isos/">isos/</a>
<a href="openshift-ansible-CentOS-SIG-PaaS">openshift-ansible-CentOS-SIG-PaaS</a>
<a href="openshift-ansible-centos-paas-sig.repo">openshift-ansible-centos-paas-sig.repo</a>
<a href="opstools/">opstools/</a>
<a href="os/">os/</a>
<a href="paas/">paas/</a>
<a href="rt/">rt/</a>
<a href="sclo/">sclo/</a>
<a href="storage/">storage/</a>
<a href="updates/">updates/</a>
<a href="virt/">virt/</a>
</pre>
```

### Fedora 26

Load archive. For more detailed, refer to [dockerized-file-server-for-centos7.md](./dockerized-file-server-for-centos7.md)
```
[vagrant@kubedev-172-17-4-59 ~]$ docker load -i /Users/fanhongling/Downloads/99-mirror/docker-images/docker.io0x2Ftangfeixiong0x2Fgofileserver.tar 
9007f5987db3: Loading layer [==================================================>]  5.05 MB/5.05 MB
fc9862ed7ee4: Loading layer [==================================================>] 7.819 MB/7.819 MB
Loaded image: docker.io/tangfeixiong/gofileserver:latest
```

Run
```
[vagrant@kubedev-172-17-4-59 ~]$ docker run -d --name=gofileserver --restart=always -v /Users/fanhongling:/mnt -p 48080:48080 docker.io/tangfeixiong/gofileserver
11ce4ad76853f8211b1dfe0624fb8742c9dd2fd650bcc72b1d896687a6a6b84d
```

Test
```
[vagrant@kubedev-172-17-4-59 ~]$ curl -L http://172.17.4.59:48080/
<pre>
<a href="Downloads/">Downloads/</a>
<a href="go/">go/</a>
</pre>
```

For example, my kubernetes apt/yum repository mirror
```
[vagrant@kubedev-172-17-4-59 ~]$ curl -L http://172.17.4.59:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/
<pre>
<a href="apt/">apt/</a>
<a href="yum/">yum/</a>
</pre>
```

### Ubuntu 16.04 akka xenial

run
```
ubuntu@ubuntu-xenial:~$ docker run -d --name=gofileserver -v /Users/fanhongling:/mnt -p 48080:48080 docker.io/tangfeixiong/gofileserver
Unable to find image 'tangfeixiong/gofileserver:latest' locally
latest: Pulling from tangfeixiong/gofileserver
c0cb142e4345: Pull complete 
bd94e5529d32: Pull complete 
Digest: sha256:b05cb512410b314f3aae3ac1932faec2c8fc892221cc773778f4efdabd3684c8
Status: Downloaded newer image for tangfeixiong/gofileserver:latest
d7a979b6b4b6b275472436863dedaa4913067bf427a4c159cd1ace1d30919d02
```

```
ubuntu@ubuntu-xenial:~$ curl -L http://127.0.0.1:48080/
<pre>
<a href="Downloads/">Downloads/</a>
<a href="go/">go/</a>
</pre>
```