# Help

## No problem accessing google container registry

do
```
[vagrant@localhost go-to-kubernetes]$ docker pull gcr.io/google_containers/pause-amd64:3.0
```

then
```
[vagrant@localhost go-to-kubernetes]$ docker images gcr.io/google_containers/pause-amd64
REPOSITORY                             TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/pause-amd64   3.0                 99e59f495ffa        18 months ago       746.9 kB
```

## If not

find
```
[vagrant@localhost go-to-kubernetes]$ docker search pause-amd64
INDEX       NAME                                  DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
docker.io   docker.io/kubernetes/pause                                                            15                   
docker.io   docker.io/mritd/pause-amd64           kunernetes image pause-amd64                    3                    [OK]
docker.io   docker.io/huangyj/pause-amd64         gcr.io/google_containers/pause-amd64            1                    [OK]
docker.io   docker.io/ist0ne/pause-amd64          https://gcr.io/google_containers/pause-amd64    1                    [OK]
docker.io   docker.io/lhcalibur/pause-amd64       pause-amd64                                     1                    [OK]
docker.io   docker.io/sailsxu/pause-amd64         pause-amd64                                     1                    [OK]
docker.io   docker.io/wangyanbin/pause-amd64      mirror from gcr.io/google_containers/pause...   1                    [OK]
docker.io   docker.io/4admin2root/pause-amd64     google_containers/pause-amd64                   0                    [OK]
docker.io   docker.io/chasontang/pause-amd64      Mirror of google pause-amd64                    0                    [OK]
docker.io   docker.io/empiregeneral/pause-amd64   pause-amd64                                     0                    [OK]
docker.io   docker.io/huanwei/pause-amd64         pause-amd64                                     0                    [OK]
docker.io   docker.io/ibmcom/pause                Docker Image for IBM Cloud private-CE (Com...   0                    
docker.io   docker.io/jicki/pause-amd64                                                           0                    
docker.io   docker.io/kdcos/pause-amd64           Mirror of gcr.io/google_containers/pause-a...   0                    
docker.io   docker.io/lvanneo/pause-amd64         pause-amd64                                     0                    [OK]
docker.io   docker.io/pigletfly/pause-amd64       pause-amd64                                     0                    [OK]
docker.io   docker.io/rancher/pause-amd64                                                         0                    
docker.io   docker.io/siriuszg/pause-amd64        pause-amd64:3.0                                 0                    [OK]
docker.io   docker.io/stackato/pause-amd64                                                        0                    
docker.io   docker.io/stackatodev/pause-amd64                                                     0                    
docker.io   docker.io/tggreene/pause-amd64        Mirror of gcr.io/google_containers/pause-a...   0                    
docker.io   docker.io/visenzek8s/pause-amd64                                                      0                    
docker.io   docker.io/warrior/pause-amd64         pause-amd64:3.0                                 0                    [OK]
docker.io   docker.io/wymr/pause-amd64            pause-amd64                                     0                    [OK]
docker.io   docker.io/zengxin/pause-amd64         pause-amd64                                     0                    [OK]
```

and
```
$ docker pull docker.io/siriuszg/pause-amd64:3.0
```

then
```
$ docker tag docker.io/siriuszg/pause-amd64:3.0 gcr.io/google_containers/pause-amd64:3.0
```

## Archiving

Save to export
```
$ docker save -o gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar gcr.io/google_containers/pause-amd64:3.0
```

Copy anywhere to import
```
$ docker load -i gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar
```

Or via local HTTP
```
$ curl -jkSL http://<local.http.addr>/<path/to/archive>.tar | docker load
```

## Issue

Maybe occur in some old version, in experience of Fedora/CentOS, for example
```
[vagrant@openshiftdev ~]$ docker save -o /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar gcr.io/google_containers/pause-amd64:3.0
```

But
```
[vagrant@localhost ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar 
Error response from daemon: invalid argument
```

Find like
```
[vagrant@localhost ~]$ docker images -f dangling=true
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
<none>              <none>              99e59f495ffa        18 months ago       746.9 kB
```

Correct
```
[vagrant@localhost ~]$ docker tag 99e59f495ffa gcr.io/google_containers/pause-amd64:3.0
```

## History

[pause:2.0](https://github.com/kubernetes/kubernetes/tree/release-1.2/build/pause)

[pause:0.8.0](https://github.com/kubernetes/kubernetes/tree/release-1.1/build/pause)