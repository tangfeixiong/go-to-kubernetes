Intruduction of *docker exec* and *kubectl exec* command 简述docker和kubectl的exec命令机制
=======================================================================================

1. *docker exec* command help and example

Help

    [vagrant@localhost go-to-kubernetes]$ docker exec --help
    
    Usage:	docker exec [OPTIONS] CONTAINER COMMAND [ARG...]
    
    Run a command in a running container
    
      -d, --detach=false         Detached mode: run command in the background
      --help=false               Print usage
      -i, --interactive=false    Keep STDIN open even if not attached
      --privileged=false         Give extended privileges to the command
      -t, --tty=false            Allocate a pseudo-TTY
      -u, --user=                Username or UID (format: <name|uid>[:<group|gid>])
    
这里，-i 参数要求docker engine创建一个stdin设备，容器内执行的shell或命令从该设备读入数据。-t参数要求docker engine创建一个伪终端设备，从该设备进行io交互

    [vagrant@localhost go-to-kubernetes]$ docker ps -l
    CONTAINER ID        IMAGE                      COMMAND             CREATED             STATUS              PORTS               NAMES
    9601f55a2faa        fatherlinux/centos4-base   "/bin/bash"         3 days ago          Up 3 days                               centos4-jdk1-5

    [vagrant@localhost go-to-kubernetes]$ docker exec -t -i centos4-jdk1-5 bash
    [root@9601f55a2faa /]# ls -ltr /dev
    total 0
    drwxrwxrwt  2 root root       40 Jul 15 20:36 mqueue
    drwxrwxrwt  2 root root       40 Jul 22 17:03 shm
    crw-------  1 root root 136,   1 Jul 22 17:03 console
    drwxr-xr-x  2 root root        0 Jul 22 17:03 pts
    crw-rw-rw-  1 root root   1,   5 Jul 22 17:03 zero
    crw-rw-rw-  1 root root   1,   9 Jul 22 17:03 urandom
    crw-rw-rw-  1 root root   5,   0 Jul 22 17:03 tty
    crw-rw-rw-  1 root root   1,   8 Jul 22 17:03 random
    lrwxrwxrwx  1 root root        8 Jul 22 17:03 ptmx -> pts/ptmx
    crw-rw-rw-  1 root root   1,   3 Jul 22 17:03 null
    c---------  1 root root  10, 229 Jul 22 17:03 fuse
    crw-rw-rw-  1 root root   1,   7 Jul 22 17:03 full
    lrwxrwxrwx  1 root root       15 Jul 22 17:03 stdout -> /proc/self/fd/1
    lrwxrwxrwx  1 root root       15 Jul 22 17:03 stdin -> /proc/self/fd/0
    lrwxrwxrwx  1 root root       15 Jul 22 17:03 stderr -> /proc/self/fd/2
    lrwxrwxrwx  1 root root       11 Jul 22 17:03 kcore -> /proc/kcore
    lrwxrwxrwx  1 root root       13 Jul 22 17:03 fd -> /proc/self/fd

    [root@9601f55a2faa /]# exit
    exit
    
显示host的file descripter和std

    [vagrant@localhost go-to-kubernetes]$ ls -l /dev/fd
    lrwxrwxrwx. 1 root root 13 7月  16 00:36 /dev/fd -> /proc/self/fd

    [vagrant@localhost go-to-kubernetes]$ ls -ltr /dev/std*
    lrwxrwxrwx. 1 root root 15 7月  16 00:36 /dev/stdout -> /proc/self/fd/1
    lrwxrwxrwx. 1 root root 15 7月  16 00:36 /dev/stdin -> /proc/self/fd/0
    lrwxrwxrwx. 1 root root 15 7月  16 00:36 /dev/stderr -> /proc/self/fd/2

2. *kubectl exec* command helm and example

help

    [vagrant@localhost go-to-kubernetes]$ kubectl exec --help
    Execute a command in a container.
    
    ...
    
    Flags:
      -c, --container="": Container name. If omitted, the first container in the pod will be chosen
      -p, --pod="": Pod name
      -i, --stdin[=false]: Pass stdin to the container
      -t, --tty[=false]: Stdin is a TTY
    
    ...

事实上，kubectl将请求发送到api server，再由api server发送到pod所在host上的kubelet，然后由kubelet调用go-dockerclient api

    [vagrant@localhost go-to-kubernetes]$ kubectl get pods
    NAME                   READY     STATUS    RESTARTS   AGE
    etcd-v3-single-69ihw   1/1       Running   0          8d
    netcat-simple          1/1       Running   0          4d
    rancher-server-jxwc8   1/1       Running   0          9d
    
    [vagrant@localhost go-to-kubernetes]$ kubectl exec -t -i netcat-simple ash
    / # ls -ltr /dev
    total 0
    drwxrwxrwt    2 root     root            40 Jul 16 00:36 mqueue
    -rw-r--r--    1 root     root             0 Jul 21 21:25 termination-log
    drwxrwxrwt    2 root     root            40 Jul 21 21:25 shm
    crw-rw-rw-    1 root     root        1,   5 Jul 21 21:25 zero
    crw-rw-rw-    1 root     root        1,   9 Jul 21 21:25 urandom
    crw-rw-rw-    1 root     root        5,   0 Jul 21 21:25 tty
    lrwxrwxrwx    1 root     root            15 Jul 21 21:25 stdout -> /proc/self/fd/1
    lrwxrwxrwx    1 root     root            15 Jul 21 21:25 stdin -> /proc/self/fd/0
    lrwxrwxrwx    1 root     root            15 Jul 21 21:25 stderr -> /proc/self/fd/2
    crw-rw-rw-    1 root     root        1,   8 Jul 21 21:25 random
    drwxr-xr-x    2 root     root             0 Jul 21 21:25 pts
    lrwxrwxrwx    1 root     root             8 Jul 21 21:25 ptmx -> pts/ptmx
    crw-rw-rw-    1 root     root        1,   3 Jul 21 21:25 null
    lrwxrwxrwx    1 root     root            11 Jul 21 21:25 kcore -> /proc/kcore
    c---------    1 root     root       10, 229 Jul 21 21:25 fuse
    crw-rw-rw-    1 root     root        1,   7 Jul 21 21:25 full
    lrwxrwxrwx    1 root     root            13 Jul 21 21:25 fd -> /proc/self/fd
    / # exit

TTY，PTY
--------

简言之，TTY是上世纪读书和工作用过的字符终端

    [vagrant@localhost go-to-kubernetes]$ ls -l /dev/tty
    crw-rw-rw-. 1 root tty 5, 0 7月  25 10:15 /dev/tty

后来，计算机显示使用图形化窗口，如linux的gnome。终端变成了一个终端窗口

    [vagrant@localhost go-to-kubernetes]$ ls -l /dev/ptmx /dev/pts
    crw-rw-rw-. 1 root tty  5, 2 7月  26 01:31 /dev/ptmx
    
    /dev/pts:
    总用量 0
    crw--w----. 1 vagrant tty  136,  0 7月  26 01:31 0
    crw-------. 1 root    root 136,  1 7月  22 21:03 1
    crw-------. 1 root    root 136, 10 7月  24 23:45 10
    crw-------. 1 root    root 136, 11 7月  24 23:51 11
    crw-------. 1 root    root 136, 12 7月  25 09:23 12
    crw-------. 1 root    root 136, 13 7月  25 10:12 13
    crw-------. 1 root    root 136,  2 7月  22 22:23 2
    crw-------. 1 root    root 136,  3 7月  25 09:38 3
    crw-------. 1 root    root 136,  4 7月  24 23:15 4
    crw-------. 1 root    root 136,  5 7月  24 23:17 5
    crw-------. 1 root    root 136,  6 7月  24 23:17 6
    crw-------. 1 root    root 136,  7 7月  24 23:20 7
    crw-------. 1 root    root 136,  8 7月  24 23:20 8
    crw-------. 1 root    root 136,  9 7月  24 23:30 9
    c---------. 1 root    root   5,  2 7月  16 00:36 ptmx

于是，linux提供了一个pseduo tty技术，[ptmx 和 pts](http://linux.die.net/man/4/ptmx)

当创建一个pty时，hal为你提供了一个ptmx下的file descriptor和一个pts的id

也就如像在gnome下打开了一个终端窗口

典型应用
-------

* Gnome terminal

* SSH server

sshd 在remote host上与ssh client进行tcp 4层通信

sshd将收到的数据写入ptm，linux应用这边，数据从ptm到pts作为应用程序的stdin

sshd从ptm读到的数据发送到网络，linux应用这边，应用程序print到stdout，即pts，再stream到ptm

Kubelet和kube api-server
------------------------

Docker engine扮演类似sshd的角色

Kublet 通过go-dockerclient 作为docker engine的代理

Kube api-server 扮演ssh client的角色，通过ip网络获取terminal内容

Kubectl
-------

kubectl是kube api server的客户端

kubectl将访问的pod和容器，及命令通过http发送到kube api server

kubectl与kube api-server之间是[spdy（即http2）通信](https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/server/remotecommand/constants.go)

* cmd call

https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/exec.go#L243

Pods().Get()

Pods().SubResource().Post()

* client api

Public method: Stream

https://github.com/kubernetes/kubernetes/blob/master/pkg/client/unversioned/remotecommand/remotecommand.go#L155

private metond: stream

https://github.com/kubernetes/kubernetes/blob/master/pkg/client/unversioned/remotecommand/v2.go#L189