If environment is blocking to access google, Refer to [./kubernetes-pod-base-image.md](./kubernetes-pod-base-image.md) for same idea

Load from archive
```
[vagrant@openshiftdev ~]$ docker load -i /vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fetcd0x2E3.0.17.tar 
38ac8d0f5bb3: Loading layer [==================================================>] 1.312 MB/1.312 MB
c872b2c2ac77: Loading layer [==================================================>] 136.7 MB/136.7 MB
be71b2a80bbd: Loading layer [==================================================>] 31.16 MB/31.16 MB
Loaded image: gcr.io/google_containers/etcd:3.0.17
```