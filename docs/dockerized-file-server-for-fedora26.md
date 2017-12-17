### HTTP-based file server serving offline docker image archive and YUM repository mirror

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