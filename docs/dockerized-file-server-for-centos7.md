# HTTP-based file server serving offline docker image archive and YUM repository mirror

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