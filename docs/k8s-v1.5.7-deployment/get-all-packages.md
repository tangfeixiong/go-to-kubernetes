Hand-on: Deploying Kubernetes Master under Kubelet －容器化部署K8s主控节点
==================================================

Table of Content －目录
----------------

Environment － 实验环境

Prerequisite － 准备

Getting K8s packages － 动手下载K8s二进制分发包

Deploying core components

* Installing _kubelet_ service － 安装kubelet为systemd或sysv服务
* Pulling POD image - 获取K8s POD镜像
* Generate TLS certificates － 为K8s主控节点生成x509证书
* Installing _etcd_ POD － 为K8s主控节点安装etcd POD
* Installing _kube-apiserver_ POD － 为K8s主控节点安装kube-apiserver POD
* Configure _kubectl_ context - 配置kubectl工具
* Installing _kube-controller-manager_ POD － 为K8s主控节点安装kube-controller-manager POD
* Installing _kube-scheduler_ POD － 为K8s主控节点安装kube-scheduler POD
* Installing _kube-proxy_ POD － 为K8s主控节点安装kube-proxy POD

Deploying addons - 安装dns和dashboard

* DNS
* Dashboard

Environment － 动手环境
-----------

Host － 主机

* Mac
* VirtualBox

VM - Linux环境

* Ubuntu 14.04
* CentOS 7.3

Prerequisite - 准备
------------

__Linux__ have installed _docker engine_ － 在Linux环境安装完成Docker－Engine


```
vagrant@vagrant-ubuntu-trusty-64:~$ docker version
Client:
 Version:      1.10.3
 API version:  1.22
 Go version:   go1.5.3
 Git commit:   20f81dd
 Built:        Thu Mar 10 15:54:52 2016
 OS/Arch:      linux/amd64

Server:
 Version:      1.10.3
 API version:  1.22
 Go version:   go1.5.3
 Git commit:   20f81dd
 Built:        Thu Mar 10 15:54:52 2016
 OS/Arch:      linux/amd64

``` 

Optional, Use _VirtualBox_ shared file system, for example:

```
vagrant@vagrant-ubuntu-trusty-64:~$ mount | grep vboxsf
go on /go type vboxsf (uid=1000,gid=1000,rw)
data on /data type vboxsf (uid=1000,gid=1000,rw)
vagrant on /vagrant type vboxsf (uid=1000,gid=1000,rw)
work_src on /work/src type vboxsf (uid=1000,gid=1000,rw)
99-mirror on /99-mirror type vboxsf (uid=1000,gid=1000,rw)

```

Getting K8s packages － 动手下载K8s二进制分发包
--------------------

The destination for downloading － 创建下载目录

```
vagrant@vagrant-ubuntu-trusty-64:~$ mkdir -p /99-mirror/https%3A%2F%2Fdl.k8s.io/v1.5.7
```

Refer to [GitHub releases](https://github.com/kubernetes/kubernetes/releases) － 使用`curl`或[`wget`](https://www.gnu.org/software/wget/manual/)下载，更多请参考社区在Github的Releases

以下是在Mac主机上操作的（99-mirror是shared file system）

```
fanhonglingdeMacBook-Pro:~ fanhongling$ cd Downloads/99-mirror/https%3A%2F%2Fdl.k8s.io/v1.5.7/
```

__The downloaded__

```
fanhonglingdeMacBook-Pro:v1.5.7 fanhongling$ curl -jkSL -O -C - https://dl.k8s.io/v1.5.7/kubernetes.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     76      0  0:00:02  0:00:02 --:--:--    76
100 6120k  100 6120k    0     0   673k      0  0:00:09  0:00:09 --:--:-- 1204k
fanhonglingdeMacBook-Pro:v1.5.7 fanhongling$ curl -jkSL -O -C - https://dl.k8s.io/v1.5.7/kubernetes-server-linux-amd64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0    102      0  0:00:01  0:00:01 --:--:--   102
100  308M  100  308M    0     0   927k      0  0:05:40  0:05:40 --:--:--  567k
fanhonglingdeMacBook-Pro:v1.5.7 fanhongling$ curl -jkSL -O -C - https://dl.k8s.io/v1.5.7/kubernetes-client-darwin-amd64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     76      0  0:00:02  0:00:02 --:--:--    76
100 22.0M  100 22.0M    0     0  1128k      0  0:00:20  0:00:20 --:--:-- 1858k
fanhonglingdeMacBook-Pro:v1.5.7 fanhongling$ curl -jkSL -O -C - https://dl.k8s.io/v1.5.7/kubernetes-client-linux-amd64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0    103      0  0:00:01  0:00:01 --:--:--   103
100 22.2M  100 22.2M    0     0  1552k      0  0:00:14  0:00:14 --:--:-- 2347k
fanhonglingdeMacBook-Pro:v1.5.7 fanhongling$ curl -jkSL -O -C - https://dl.k8s.io/v1.5.7/kubernetes-client-windows-amd64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0    102      0  0:00:01  0:00:01 --:--:--   102
100 22.2M  100 22.2M    0     0  1159k      0  0:00:19  0:00:19 --:--:-- 2370k
fanhonglingdeMacBook-Pro:v1.5.7 fanhongling$ curl -jkSL -O -C - https://dl.k8s.io/v1.5.7/kubernetes-node-linux-amd64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     98      0  0:00:01  0:00:01 --:--:--    98
100 77.1M  100 77.1M    0     0   334k      0  0:03:55  0:03:55 --:--:--  765k
fanhonglingdeMacBook-Pro:v1.5.7 fanhongling$ curl -jkSL -O -C - https://dl.k8s.io/v1.5.7/kubernetes-node-windows-amd64.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     77      0  0:00:02  0:00:02 --:--:--    77
100 75.4M  100 75.4M    0     0  1037k      0  0:01:14  0:01:14 --:--:--  765k

```

在Linux虚拟机上显示下载的内容

```
vagrant@vagrant-ubuntu-trusty-64:~$ ls -1 /99-mirror/https%3A%2F%2Fdl.k8s.io/v1.5.7/
kubernetes-client-darwin-amd64.tar.gz
kubernetes-client-linux-amd64.tar.gz
kubernetes-client-windows-amd64.tar.gz
kubernetes-node-linux-amd64.tar.gz
kubernetes-node-windows-amd64.tar.gz
kubernetes-server-linux-amd64.tar.gz
kubernetes.tar.gz

```

### K8s的二进制和Docker镜像

二进制和镜像的内容

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /99-mirror/https%3A%2F%2Fdl.k8s.io/v1.5.7/kubernetes-server-linux-amd64.tar.gz 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/server
bin
kubernetes-manifests.tar.gz
kubernetes-salt.tar.gz
README
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/server/bin
hyperkube
kubeadm
kube-apiserver
kube-apiserver.docker_tag
kube-apiserver.tar
kube-controller-manager
kube-controller-manager.docker_tag
kube-controller-manager.tar
kubectl
kube-discovery
kube-dns
kubefed
kubelet
kube-proxy
kube-proxy.docker_tag
kube-proxy.tar
kube-scheduler
kube-scheduler.docker_tag
kube-scheduler.tar

```

加载镜像到Docker engine

```
vagrant@vagrant-ubuntu-trusty-64:~$ tag=$(sudo cat /opt/kubernetes/server/bin/kube-apiserver.docker_tag);echo $tag; img=$(sudo docker load -i /opt/kubernetes/server/bin/kube-apiserver.tar); echo ${img:-...loaded}; docker images gcr.io/google_containers/kube-apiserver
9312bc9403d9a1dda49a2fe70c897deb
...loaded
REPOSITORY                                TAG                                IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver   9312bc9403d9a1dda49a2fe70c897deb   884ad1e5866e        3 days ago          132.2 MB
vagrant@vagrant-ubuntu-trusty-64:~$ tag=$(sudo cat /opt/kubernetes/server/bin/kube-controller-manager.docker_tag);echo $tag; img=$(sudo docker load -i /opt/kubernetes/server/bin/kube-controller-manager.tar); echo ${img:-...loaded}; docker images gcr.io/google_containers/kube-controller-manager
0005ed016737cb2c14929aa8efdf5be7
...loaded
REPOSITORY                                         TAG                                IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-controller-manager   0005ed016737cb2c14929aa8efdf5be7   b16ec381a53e        3 days ago          109 MB
vagrant@vagrant-ubuntu-trusty-64:~$ tag=$(sudo cat /opt/kubernetes/server/bin/kube-proxy.docker_tag);echo $tag; img=$(sudo docker load -i /opt/kubernetes/server/bin/kube-proxy.tar); echo ${img:-...loaded}; docker images gcr.io/google_containers/kube-proxy
690a68ee4bacc47b0115413489185f3d
...loaded
REPOSITORY                            TAG                                IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-proxy   690a68ee4bacc47b0115413489185f3d   0cdeac571785        3 days ago          173.5 MB
vagrant@vagrant-ubuntu-trusty-64:~$ tag=$(sudo cat /opt/kubernetes/server/bin/kube-scheduler.docker_tag);echo $tag; img=$(sudo docker load -i /opt/kubernetes/server/bin/kube-scheduler.tar); echo ${img:-...loaded}; docker images gcr.io/google_containers/kube-scheduler
5e5a1fe4ee49b94c0f3d2e745a48cb65
...loaded
REPOSITORY                                TAG                                IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-scheduler   5e5a1fe4ee49b94c0f3d2e745a48cb65   f533648f5d7f        3 days ago          58.82 MB
vagrant@vagrant-ubuntu-trusty-64:~$
vagrant@vagrant-ubuntu-trusty-64:~$ docker images | grep google_containers
gcr.io/google_containers/kube-controller-manager           0005ed016737cb2c14929aa8efdf5be7   b16ec381a53e        6 days ago          109 MB
gcr.io/google_containers/kube-apiserver                    9312bc9403d9a1dda49a2fe70c897deb   884ad1e5866e        6 days ago          132.2 MB
gcr.io/google_containers/kube-scheduler                    5e5a1fe4ee49b94c0f3d2e745a48cb65   f533648f5d7f        6 days ago          58.82 MB
gcr.io/google_containers/kube-proxy                        690a68ee4bacc47b0115413489185f3d   0cdeac571785        6 days ago          173.5 MB

```

Check `kubectl` command － 命令工具

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl version --client
Client Version: version.Info{Major:"1", Minor:"5", GitVersion:"v1.5.7", GitCommit:"8eb75a5810cba92ccad845ca360cf924f2385881", GitTreeState:"clean", BuildDate:"2017-04-27T10:00:30Z", GoVersion:"go1.7.5", Compiler:"gc", Platform:"linux/amd64"}
```

## 安装核心组件

### Installing _kubelet_ service - 安装kubelet服务

Appendix of [the help of kubelet tool](#command-kubelet) － 附录kubelet二进制执行帮助

Extract __saltbase__ － 解压服务脚本的模版

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/kubelet/
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/saltbase/salt/kubelet
default
initd
init.sls
kubeconfig
kubelet.service

```

__For _systemd_ service － 如果为CentOS 7.*，Fedora 20+系列__

安装systemd服务

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo sed -e 's%\(ExecStart=/usr/local/bin/kubelet.*\)%# \1\nExecStart=/opt/kubernetes/server/bin/kubelet "$DAEMON_ARGS"%' /opt/kubernetes/saltbase/salt/kubelet/kubelet.service | sudo tee /etc/systemd/system/kubelet.service
[Unit]
Description=Kubernetes Kubelet Server
Documentation=https://github.com/kubernetes/kubernetes

[Service]
EnvironmentFile=/etc/sysconfig/kubelet
# ExecStart=/usr/local/bin/kubelet "$DAEMON_ARGS"
ExecStart=/opt/kubernetes/server/bin/kubelet "$DAEMON_ARGS"
Restart=always
RestartSec=2s
StartLimitInterval=0
KillMode=process

[Install]
WantedBy=multi-user.target

```

Service的运行配置

（TBD）

__For _sysv_ and/or _upstart_ service － 如果为Ubuntu，Debian系列__

安装sysv服务

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo sed -e 's%\(DAEMON=/usr/local/bin/kubelet.*\)%# \1\nDAEMON=/opt/kubernetes/server/bin/kubelet%' /opt/kubernetes/saltbase/salt/kubelet/initd | sudo tee /etc/init.d/kubelet
#!/bin/bash
#
### BEGIN INIT INFO
# Provides:   kubelet 
# Required-Start:    $local_fs $network $syslog
# Required-Stop:
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: The Kubernetes node container manager
# Description:
#   The Kubernetes container manager maintains docker state against a state file.
### END INIT INFO


# PATH should only include /usr/* if it runs after the mountnfs.sh script
PATH=/sbin:/usr/sbin:/bin:/usr/bin
DESC="The Kubernetes container manager"
NAME=kubelet
# DAEMON=/usr/local/bin/kubelet
DAEMON=/opt/kubernetes/server/bin/kubelet
DAEMON_ARGS=""
DAEMON_LOG_FILE=/var/log/$NAME.log
PIDFILE=/var/run/$NAME.pid
SCRIPTNAME=/etc/init.d/$NAME
DAEMON_USER=root

# Exit if the package is not installed
[ -x "$DAEMON" ] || exit 0

# Read configuration variable file if it is present
[ -r /etc/default/$NAME ] && . /etc/default/$NAME

# Define LSB log_* functions.
# Depend on lsb-base (>= 3.2-14) to ensure that this file is present
# and status_of_proc is working.
. /lib/lsb/init-functions

#
# Function that starts the daemon/service
#
do_start()
{
        # Avoid a potential race at boot time when both monit and init.d start
        # the same service
        PIDS=$(pidof $DAEMON)
        for PID in ${PIDS}; do
            kill -9 $PID
	done

        # Return
        #   0 if daemon has been started
        #   1 if daemon was already running
        #   2 if daemon could not be started
        start-stop-daemon --start --quiet --background --no-close \
                --make-pidfile --pidfile $PIDFILE \
                --exec $DAEMON -c $DAEMON_USER --test > /dev/null \
                || return 1
        start-stop-daemon --start --quiet --background --no-close \
                --make-pidfile --pidfile $PIDFILE \
                --exec $DAEMON -c $DAEMON_USER -- \
                $DAEMON_ARGS >> $DAEMON_LOG_FILE 2>&1 \
                || return 2
}

#
# Function that stops the daemon/service
#
do_stop()
{
        # Return
        #   0 if daemon has been stopped
        #   1 if daemon was already stopped
        #   2 if daemon could not be stopped
        #   other if a failure occurred
        start-stop-daemon --stop --quiet --retry=TERM/30/KILL/5 --pidfile $PIDFILE --name $NAME
        RETVAL="$?"
        [ "$RETVAL" = 2 ] && return 2
        # Many daemons don't delete their pidfiles when they exit.
        rm -f $PIDFILE
        return "$RETVAL"
}


case "$1" in
  start)
        log_daemon_msg "Starting $DESC" "$NAME"
        do_start
        case "$?" in
                0|1) log_end_msg 0 || exit 0 ;;
                2) log_end_msg 1 || exit 1 ;;
        esac
        ;;
  stop)
        log_daemon_msg "Stopping $DESC" "$NAME"
        do_stop
        case "$?" in
                0|1) log_end_msg 0 ;;
                2) exit 1 ;;
        esac
        ;;
  status)
        status_of_proc -p $PIDFILE "$DAEMON" "$NAME" && exit 0 || exit $?
        ;;

  restart|force-reload)
        log_daemon_msg "Restarting $DESC" "$NAME"
        do_stop
        case "$?" in
          0|1)
                do_start
                case "$?" in
                        0) log_end_msg 0 ;;
                        1) log_end_msg 1 ;; # Old process is still running
                        *) log_end_msg 1 ;; # Failed to start
                esac
                ;;
          *)
                # Failed to stop
                log_end_msg 1
                ;;
        esac
        ;;
  *)
        echo "Usage: $SCRIPTNAME {start|stop|status|restart|force-reload}" >&2
        exit 3
        ;;
esac
vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo chmod +x /etc/init.d/kubelet

```

Service的运行配置

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo sed '/^# test_args.*$/,$!d;s/{{daemon_args}}/$DAEMON_ARGS/;s%{{api_servers_with_port}}%--api-servers=http://127.0.0.1:8080%;s/{{debugging_handlers}}/--enable-debugging-handlers=true/;s/{{hostname_override}}/--hostname-override=172.17.4.200/;s/{{cloud_provider}} {{cloud_config}}//;s%{{config}} {{manifest_url}}%--pod-manifest-path=/etc/kubernetes/manifests%;s/{{pillar\[\x27allow_privileged\x27\]}}/true/;s/{{log_level}}/--v=2/;s/{{cluster_dns}} {{cluster_domain}}/--cluster-dns=10.123.240.10 --cluster-domain=cluster.local/;s%{{docker_root}} {{kubelet_root}}%--root-dir=/var/lib/kubelet%;s%{{non_masquerade_cidr}}%--non-masquerade-cidr=10.0.0.0/8%;s%{{cgroup_root}} {{system_container}}%--cgroup-root=/ --system-cgroups=/system%;s%{{pod_cidr}}%--pod-cidr=10.120.0.0/14%;s%{{ master_kubelet_args }}%--address=0.0.0.0 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --node-ip=172.17.4.200%;s/{{cpu_cfs_quota}}/--cpu-cfs-quota=true/;s/{{network_plugin}}//;s/{{kubelet_port}}/--port=10250/;s/{{ hairpin_mode }} {{enable_custom_metrics}}/--enable-custom-metrics=true/;s%{{runtime_container}} {{kubelet_container}}%--runtime-cgroups=/docker-daemon --kubelet-cgroups=/kubelet%;s/{{node_labels}}//;s/{{babysit_daemons}}//;s/{{eviction_hard}} {{feature_gates}}/--eviction-hard=memory.available<100Mi --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true/;s/{{test_args}}/--anonymous-auth=true --authorization-mode=AlwaysAllow/' /opt/kubernetes/saltbase/salt/kubelet/default | sudo tee /etc/default/kubelet
# test_args has to be kept at the end, so they'll overwrite any prior configuration
DAEMON_ARGS="$DAEMON_ARGS --api-servers=http://127.0.0.1:8080 --enable-debugging-handlers=true --hostname-override=172.17.4.200  --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --v=2 --cluster-dns=10.123.240.10 --cluster-domain=cluster.local --root-dir=/var/lib/kubelet  --non-masquerade-cidr=10.0.0.0/8 --cgroup-root=/ --system-cgroups=/system --pod-cidr=10.120.0.0/14 --address=0.0.0.0 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --node-ip=172.17.4.200 --cpu-cfs-quota=true  --port=10250 --enable-custom-metrics=true --runtime-cgroups=/docker-daemon --kubelet-cgroups=/kubelet   --eviction-hard=memory.available<100Mi --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --anonymous-auth=true --authorization-mode=AlwaysAllow"

```

运行服务

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /etc/init.d/kubelet status
 * kubelet is not running
vagrant@vagrant-ubuntu-trusty-64:~$ sudo service kubelet status
 * kubelet is not running
vagrant@vagrant-ubuntu-trusty-64:~$
vagrant@vagrant-ubuntu-trusty-64:~$ sudo service kubelet start
 * Starting The Kubernetes container manager kubelet                                                                                           [ OK ] 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo service kubelet status
 * kubelet is running
vagrant@vagrant-ubuntu-trusty-64:~$
vagrant@vagrant-ubuntu-trusty-64:~$ sudo service kubelet restart
 * Restarting The Kubernetes container manager kubelet                                                                                         [ OK ] 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo service kubelet status
 * kubelet is running

```

配置Sysv服务的Autostart

（TBD）

检查kubelet的日志（这是服务不成功的记录） - Notice there was shown some default value for runtime

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo cat /var/log/kubelet.log 
Flag --api-servers has been deprecated, Use --kubeconfig instead. Will be removed in a future version.
I0430 11:03:15.113782    5680 feature_gate.go:189] feature gates: map[StreamingProxyRedirects:true ExperimentalHostUserNamespaceDefaulting:true AllAlpha:true ExperimentalCriticalPodAnnotation:true DynamicKubeletConfig:true DynamicVolumeProvisioning:true]
W0430 11:03:15.113891    5680 server.go:605] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
I0430 11:03:15.114346    5680 server.go:217] Starting Kubelet configuration sync loop
W0430 11:03:15.114409    5680 server.go:605] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
I0430 11:03:15.208128    5680 docker.go:356] Connecting to docker on unix:///var/run/docker.sock
I0430 11:03:15.208148    5680 docker.go:376] Start docker client with request timeout=2m0s
E0430 11:03:15.229011    5680 cni.go:163] error updating cni config: No networks found in /etc/cni/net.d
I0430 11:03:15.234687    5680 manager.go:143] cAdvisor running in container: "/user/1000.user/2.session"
W0430 11:03:15.862303    5680 manager.go:151] unable to connect to Rkt api service: rkt: cannot tcp Dial rkt api service: dial tcp 127.0.0.1:15441: getsockopt: connection refused
I0430 11:03:16.060671    5680 fs.go:117] Filesystem partitions: map[/dev/sda1:{mountpoint:/var/lib/docker/devicemapper major:8 minor:1 fsType:ext4 blockSize:0}]
I0430 11:03:16.063030    5680 manager.go:198] Machine: {NumCores:1 CpuFrequency:2699957 MemoryCapacity:3156168704 MachineID:225f84ab8484258f0641bdd356a933f3 SystemUUID:4C02743F-0C8F-42E5-BBD2-43A221BFA015 BootID:faf70a00-990c-4998-9bfc-dc67a52f3d49 Filesystems:[{Device:/dev/sda1 Capacity:42241163264 Type:vfs Inodes:2621440 HasInodes:true}] DiskMap:map[252:0:{Name:dm-0 Major:252 Minor:0 Size:107374182400 Scheduler:none} 252:1:{Name:dm-1 Major:252 Minor:1 Size:107374182400 Scheduler:none} 8:0:{Name:sda Major:8 Minor:0 Size:42949672960 Scheduler:deadline}] NetworkDevices:[{Name:eth0 MacAddress:08:00:27:36:0e:7e Speed:1000 Mtu:1500} {Name:eth1 MacAddress:08:00:27:9a:ac:74 Speed:1000 Mtu:1500}] Topology:[{Id:0 Memory:3156168704 Cores:[{Id:0 Threads:[0] Caches:[{Size:32768 Type:Data Level:1} {Size:32768 Type:Instruction Level:1} {Size:262144 Type:Unified Level:2}]}] Caches:[{Size:3145728 Type:Unified Level:3}]}] CloudProvider:Unknown InstanceType:Unknown InstanceID:None}
I0430 11:03:16.064527    5680 manager.go:204] Version: {KernelVersion:3.13.0-76-generic ContainerOsVersion:Ubuntu 14.04.5 LTS DockerVersion:1.10.3 CadvisorVersion: CadvisorRevision:}
I0430 11:03:16.066058    5680 server.go:700] Using root directory: /var/lib/kubelet
I0430 11:03:16.066098    5680 kubelet.go:242] Adding manifest file: /etc/kubernetes/manifests
I0430 11:03:16.066119    5680 file.go:47] Watching path "/etc/kubernetes/manifests"
I0430 11:03:16.066132    5680 kubelet.go:252] Watching apiserver
E0430 11:03:16.067349    5680 file.go:71] unable to read config path "/etc/kubernetes/manifests": path does not exist, ignoring
E0430 11:03:16.068027    5680 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:03:16.068103    5680 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:03:16.068171    5680 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
I0430 11:03:16.068500    5680 kubelet.go:467] Experimental host user namespace defaulting is enabled.
I0430 11:03:16.068517    5680 kubelet.go:477] Hairpin mode set to "promiscuous-bridge"
error: failed to run Kubelet: failed to create kubelet: Network plugin "kubenet" failed init: Failed to find nsenter binary: exec: "nsenter": executable file not found in $PATH
Flag --api-servers has been deprecated, Use --kubeconfig instead. Will be removed in a future version.
I0430 11:07:58.362224    5738 feature_gate.go:189] feature gates: map[ExperimentalHostUserNamespaceDefaulting:true ExperimentalCriticalPodAnnotation:true AllAlpha:true DynamicKubeletConfig:true DynamicVolumeProvisioning:true StreamingProxyRedirects:true]
W0430 11:07:58.362420    5738 server.go:605] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
I0430 11:07:58.362909    5738 server.go:217] Starting Kubelet configuration sync loop
W0430 11:07:58.362926    5738 server.go:605] Could not load kubeconfig file /var/lib/kubelet/kubeconfig: stat /var/lib/kubelet/kubeconfig: no such file or directory. Using default client config instead.
I0430 11:07:58.363142    5738 docker.go:356] Connecting to docker on unix:///var/run/docker.sock
I0430 11:07:58.363153    5738 docker.go:376] Start docker client with request timeout=2m0s
E0430 11:07:58.363833    5738 cni.go:163] error updating cni config: No networks found in /etc/cni/net.d
I0430 11:07:58.366898    5738 manager.go:143] cAdvisor running in container: "/user/1000.user/2.session"
W0430 11:07:58.574943    5738 manager.go:151] unable to connect to Rkt api service: rkt: cannot tcp Dial rkt api service: dial tcp 127.0.0.1:15441: getsockopt: connection refused
I0430 11:07:58.764916    5738 fs.go:117] Filesystem partitions: map[/dev/sda1:{mountpoint:/var/lib/docker/devicemapper major:8 minor:1 fsType:ext4 blockSize:0}]
I0430 11:07:58.769436    5738 manager.go:198] Machine: {NumCores:1 CpuFrequency:2699957 MemoryCapacity:3156168704 MachineID:225f84ab8484258f0641bdd356a933f3 SystemUUID:4C02743F-0C8F-42E5-BBD2-43A221BFA015 BootID:faf70a00-990c-4998-9bfc-dc67a52f3d49 Filesystems:[{Device:/dev/sda1 Capacity:42241163264 Type:vfs Inodes:2621440 HasInodes:true}] DiskMap:map[252:0:{Name:dm-0 Major:252 Minor:0 Size:107374182400 Scheduler:none} 252:1:{Name:dm-1 Major:252 Minor:1 Size:107374182400 Scheduler:none} 8:0:{Name:sda Major:8 Minor:0 Size:42949672960 Scheduler:deadline}] NetworkDevices:[{Name:eth0 MacAddress:08:00:27:36:0e:7e Speed:1000 Mtu:1500} {Name:eth1 MacAddress:08:00:27:9a:ac:74 Speed:1000 Mtu:1500}] Topology:[{Id:0 Memory:3156168704 Cores:[{Id:0 Threads:[0] Caches:[{Size:32768 Type:Data Level:1} {Size:32768 Type:Instruction Level:1} {Size:262144 Type:Unified Level:2}]}] Caches:[{Size:3145728 Type:Unified Level:3}]}] CloudProvider:Unknown InstanceType:Unknown InstanceID:None}
I0430 11:07:58.770060    5738 manager.go:204] Version: {KernelVersion:3.13.0-76-generic ContainerOsVersion:Ubuntu 14.04.5 LTS DockerVersion:1.10.3 CadvisorVersion: CadvisorRevision:}
I0430 11:07:58.771687    5738 server.go:700] Using root directory: /var/lib/kubelet
I0430 11:07:58.771721    5738 kubelet.go:242] Adding manifest file: /etc/kubernetes/manifests
I0430 11:07:58.771739    5738 file.go:47] Watching path "/etc/kubernetes/manifests"
I0430 11:07:58.771747    5738 kubelet.go:252] Watching apiserver
E0430 11:07:58.772983    5738 file.go:71] unable to read config path "/etc/kubernetes/manifests": path does not exist, ignoring
E0430 11:07:58.773726    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:07:58.773816    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:07:58.773893    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
I0430 11:07:58.774252    5738 kubelet.go:467] Experimental host user namespace defaulting is enabled.
W0430 11:07:58.774269    5738 kubelet_network.go:69] Hairpin mode set to "promiscuous-bridge" but kubenet is not enabled, falling back to "hairpin-veth"
I0430 11:07:58.774283    5738 kubelet.go:477] Hairpin mode set to "hairpin-veth"
I0430 11:07:58.971430    5738 docker_manager.go:256] Setting dockerRoot to /var/lib/docker
I0430 11:07:58.971452    5738 docker_manager.go:259] Setting cgroupDriver to 
I0430 11:07:58.971544    5738 kubelet_network.go:226] Setting Pod CIDR:  -> 10.120.0.0/14
I0430 11:07:58.971721    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/aws-ebs"
I0430 11:07:58.971737    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/empty-dir"
I0430 11:07:58.971746    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/gce-pd"
I0430 11:07:58.971754    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/git-repo"
I0430 11:07:58.971763    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/host-path"
I0430 11:07:58.971771    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/nfs"
I0430 11:07:58.971779    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/secret"
I0430 11:07:58.971787    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/iscsi"
I0430 11:07:58.971802    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/glusterfs"
I0430 11:07:58.971816    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/rbd"
I0430 11:07:58.971830    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/cinder"
I0430 11:07:58.971842    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/quobyte"
I0430 11:07:58.971849    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/cephfs"
I0430 11:07:58.971858    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/downward-api"
I0430 11:07:58.971866    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/fc"
I0430 11:07:58.971874    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/flocker"
I0430 11:07:58.971881    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/azure-file"
I0430 11:07:58.971890    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/configmap"
I0430 11:07:58.971899    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/vsphere-volume"
I0430 11:07:58.971907    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/azure-disk"
I0430 11:07:58.971915    5738 plugins.go:344] Loaded volume plugin "kubernetes.io/photon-pd"
I0430 11:07:58.972577    5738 server.go:735] Setting keys quota in /proc/sys/kernel/keys/root_maxkeys to 1000000
I0430 11:07:58.972603    5738 server.go:751] Setting keys bytes in /proc/sys/kernel/keys/root_maxbytes to 25000000
I0430 11:07:58.972619    5738 server.go:770] Started kubelet v1.5.7
E0430 11:07:58.973046    5738 kubelet.go:1145] Image garbage collection failed: unable to find data for container /
I0430 11:07:58.973247    5738 kubelet_node_status.go:204] Setting node annotation to enable volume controller attach/detach
I0430 11:07:58.973860    5738 kubelet_node_status.go:370] Using node IP: "172.17.4.200"
I0430 11:07:58.974028    5738 server.go:123] Starting to listen on 0.0.0.0:10250
I0430 11:07:58.975727    5738 server.go:140] Starting to listen read-only on 0.0.0.0:10255
E0430 11:07:58.976516    5738 event.go:208] Unable to write event: 'Post https://127.0.0.1:6443/api/v1/namespaces/default/events: dial tcp 127.0.0.1:6443: getsockopt: connection refused' (may retry after sleeping)
E0430 11:07:59.023399    5738 kubelet.go:1634] Failed to check if disk space is available for the runtime: failed to get fs info for "runtime": unable to find data for container /
E0430 11:07:59.023440    5738 kubelet.go:1642] Failed to check if disk space is available on the root partition: failed to get fs info for "root": unable to find data for container /
I0430 11:07:59.023455    5738 kubelet_node_status.go:358] Recording NodeHasSufficientDisk event message for node k8s-172-17-4-200
I0430 11:07:59.023507    5738 kubelet_node_status.go:358] Recording NodeHasSufficientMemory event message for node k8s-172-17-4-200
I0430 11:07:59.023532    5738 kubelet_node_status.go:358] Recording NodeHasNoDiskPressure event message for node k8s-172-17-4-200
I0430 11:07:59.023920    5738 container_manager_linux.go:335] Updating kernel flag: vm/overcommit_memory, expected value: 1, actual value: 0
I0430 11:07:59.023960    5738 container_manager_linux.go:335] Updating kernel flag: kernel/panic, expected value: 10, actual value: 0
I0430 11:07:59.023977    5738 container_manager_linux.go:335] Updating kernel flag: kernel/panic_on_oops, expected value: 1, actual value: 0
I0430 11:07:59.043916    5738 container_manager_linux.go:405] Configure resource-only container /docker-daemon with memory limit: 2209318092
I0430 11:07:59.043977    5738 fs_resource_analyzer.go:66] Starting FS ResourceAnalyzer
I0430 11:07:59.044012    5738 status_manager.go:129] Starting to sync pod status with apiserver
I0430 11:07:59.044044    5738 kubelet.go:1714] Starting kubelet main sync loop.
I0430 11:07:59.044081    5738 kubelet.go:1725] skipping pod synchronization - [container runtime is down]
I0430 11:07:59.046845    5738 container_manager_linux.go:769] Found 83 PIDs in root, 58 of them are not to be moved
I0430 11:07:59.046861    5738 container_manager_linux.go:776] Moving non-kernel processes: [477 482 637 771 798 801 931 962 991 998 1036 1134 1137 1141 1142 1144 1187 1188 1189 1277 1478 1653 1758 1998 2052]
I0430 11:07:59.076406    5738 volume_manager.go:240] The desired_state_of_world populator starts
I0430 11:07:59.076434    5738 volume_manager.go:242] Starting Kubelet Volume Manager
E0430 11:07:59.122112    5738 event.go:208] Unable to write event: 'Post https://127.0.0.1:6443/api/v1/namespaces/default/events: dial tcp 127.0.0.1:6443: getsockopt: connection refused' (may retry after sleeping)
I0430 11:07:59.190740    5738 container_manager_linux.go:769] Found 58 PIDs in root, 58 of them are not to be moved
I0430 11:07:59.192587    5738 kubelet_node_status.go:204] Setting node annotation to enable volume controller attach/detach
I0430 11:07:59.192735    5738 kubelet_node_status.go:370] Using node IP: "172.17.4.200"
E0430 11:07:59.277573    5738 kubelet.go:1634] Failed to check if disk space is available for the runtime: failed to get fs info for "runtime": unable to find data for container /
E0430 11:07:59.277606    5738 kubelet.go:1642] Failed to check if disk space is available on the root partition: failed to get fs info for "root": unable to find data for container /
I0430 11:07:59.277612    5738 kubelet_node_status.go:358] Recording NodeHasSufficientDisk event message for node k8s-172-17-4-200
I0430 11:07:59.277625    5738 kubelet_node_status.go:358] Recording NodeHasSufficientMemory event message for node k8s-172-17-4-200
I0430 11:07:59.277640    5738 kubelet_node_status.go:358] Recording NodeHasNoDiskPressure event message for node k8s-172-17-4-200
I0430 11:07:59.277662    5738 kubelet_node_status.go:74] Attempting to register node k8s-172-17-4-200
E0430 11:07:59.278041    5738 kubelet_node_status.go:98] Unable to register node "k8s-172-17-4-200" with API server: Post https://127.0.0.1:6443/api/v1/nodes: dial tcp 127.0.0.1:6443: getsockopt: connection refused
I0430 11:07:59.482968    5738 kubelet_node_status.go:204] Setting node annotation to enable volume controller attach/detach
I0430 11:07:59.483225    5738 kubelet_node_status.go:370] Using node IP: "172.17.4.200"
E0430 11:07:59.519956    5738 factory.go:305] devicemapper filesystem stats will not be reported: unable to find thin_ls binary
I0430 11:07:59.519977    5738 factory.go:309] Registering Docker factory
W0430 11:07:59.519986    5738 manager.go:247] Registration of the rkt container factory failed: unable to communicate with Rkt api service: rkt: cannot tcp Dial rkt api service: dial tcp 127.0.0.1:15441: getsockopt: connection refused
I0430 11:07:59.519991    5738 factory.go:54] Registering systemd factory
I0430 11:07:59.520088    5738 factory.go:86] Registering Raw factory
I0430 11:07:59.520186    5738 manager.go:1106] Started watching for new ooms in manager
I0430 11:07:59.520221    5738 oomparser.go:200] OOM parser using kernel log file: "/var/log/kern.log"
I0430 11:07:59.520521    5738 manager.go:288] Starting recovery of all containers
I0430 11:07:59.539222    5738 kubelet_node_status.go:358] Recording NodeHasSufficientDisk event message for node k8s-172-17-4-200
I0430 11:07:59.539251    5738 kubelet_node_status.go:358] Recording NodeHasSufficientMemory event message for node k8s-172-17-4-200
I0430 11:07:59.539259    5738 kubelet_node_status.go:358] Recording NodeHasNoDiskPressure event message for node k8s-172-17-4-200
I0430 11:07:59.539274    5738 kubelet_node_status.go:74] Attempting to register node k8s-172-17-4-200
E0430 11:07:59.539510    5738 kubelet_node_status.go:98] Unable to register node "k8s-172-17-4-200" with API server: Post https://127.0.0.1:6443/api/v1/nodes: dial tcp 127.0.0.1:6443: getsockopt: connection refused
I0430 11:07:59.737342    5738 manager.go:293] Recovery completed
E0430 11:07:59.747410    5738 summary.go:70] Partial failure issuing GetContainerInfoV2: partial failures: ["/docker-daemon": RecentStats: unable to find data for container /docker-daemon], ["/user": RecentStats: unable to find data for container /user], ["/kubelet": RecentStats: unable to find data for container /kubelet], ["/system": RecentStats: unable to find data for container /system]
E0430 11:07:59.747480    5738 eviction_manager.go:205] eviction manager: unexpected err: failed GetNode: node 'k8s-172-17-4-200' not found
E0430 11:07:59.775414    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:07:59.779215    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:07:59.783736    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
I0430 11:07:59.941180    5738 kubelet_node_status.go:204] Setting node annotation to enable volume controller attach/detach
I0430 11:07:59.941539    5738 kubelet_node_status.go:370] Using node IP: "172.17.4.200"
I0430 11:07:59.942256    5738 kubelet_node_status.go:358] Recording NodeHasSufficientDisk event message for node k8s-172-17-4-200
I0430 11:07:59.942278    5738 kubelet_node_status.go:358] Recording NodeHasSufficientMemory event message for node k8s-172-17-4-200
I0430 11:07:59.942287    5738 kubelet_node_status.go:358] Recording NodeHasNoDiskPressure event message for node k8s-172-17-4-200
I0430 11:07:59.942300    5738 kubelet_node_status.go:74] Attempting to register node k8s-172-17-4-200
E0430 11:07:59.942667    5738 kubelet_node_status.go:98] Unable to register node "k8s-172-17-4-200" with API server: Post https://127.0.0.1:6443/api/v1/nodes: dial tcp 127.0.0.1:6443: getsockopt: connection refused
I0430 11:08:00.743052    5738 kubelet_node_status.go:204] Setting node annotation to enable volume controller attach/detach
I0430 11:08:00.743394    5738 kubelet_node_status.go:370] Using node IP: "172.17.4.200"
I0430 11:08:00.744859    5738 kubelet_node_status.go:358] Recording NodeHasSufficientDisk event message for node k8s-172-17-4-200
I0430 11:08:00.744891    5738 kubelet_node_status.go:358] Recording NodeHasSufficientMemory event message for node k8s-172-17-4-200
I0430 11:08:00.744909    5738 kubelet_node_status.go:358] Recording NodeHasNoDiskPressure event message for node k8s-172-17-4-200
I0430 11:08:00.744927    5738 kubelet_node_status.go:74] Attempting to register node k8s-172-17-4-200
E0430 11:08:00.745279    5738 kubelet_node_status.go:98] Unable to register node "k8s-172-17-4-200" with API server: Post https://127.0.0.1:6443/api/v1/nodes: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:00.776154    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:00.780112    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:00.790138    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:01.777835    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:01.782136    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:01.791556    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
I0430 11:08:02.345607    5738 kubelet_node_status.go:204] Setting node annotation to enable volume controller attach/detach
I0430 11:08:02.345953    5738 kubelet_node_status.go:370] Using node IP: "172.17.4.200"
I0430 11:08:02.347033    5738 kubelet_node_status.go:358] Recording NodeHasSufficientDisk event message for node k8s-172-17-4-200
I0430 11:08:02.347064    5738 kubelet_node_status.go:358] Recording NodeHasSufficientMemory event message for node k8s-172-17-4-200
I0430 11:08:02.347077    5738 kubelet_node_status.go:358] Recording NodeHasNoDiskPressure event message for node k8s-172-17-4-200
I0430 11:08:02.347096    5738 kubelet_node_status.go:74] Attempting to register node k8s-172-17-4-200
E0430 11:08:02.347965    5738 kubelet_node_status.go:98] Unable to register node "k8s-172-17-4-200" with API server: Post https://127.0.0.1:6443/api/v1/nodes: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:02.778416    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:02.782864    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:02.793541    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:03.779071    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:03.783963    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:03.795601    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
I0430 11:08:04.045587    5738 kubelet.go:1781] SyncLoop (ADD, "file"): ""
E0430 11:08:04.780745    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:04.786288    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:04.797930    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
I0430 11:08:05.549653    5738 kubelet_node_status.go:204] Setting node annotation to enable volume controller attach/detach
I0430 11:08:05.549986    5738 kubelet_node_status.go:370] Using node IP: "172.17.4.200"
I0430 11:08:05.550715    5738 kubelet_node_status.go:358] Recording NodeHasSufficientDisk event message for node k8s-172-17-4-200
I0430 11:08:05.550747    5738 kubelet_node_status.go:358] Recording NodeHasSufficientMemory event message for node k8s-172-17-4-200
I0430 11:08:05.550759    5738 kubelet_node_status.go:358] Recording NodeHasNoDiskPressure event message for node k8s-172-17-4-200
I0430 11:08:05.550775    5738 kubelet_node_status.go:74] Attempting to register node k8s-172-17-4-200
E0430 11:08:05.551151    5738 kubelet_node_status.go:98] Unable to register node "k8s-172-17-4-200" with API server: Post https://127.0.0.1:6443/api/v1/nodes: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:05.782814    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:05.787747    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:05.800080    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:06.784389    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:06.788644    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:06.801618    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:07.786345    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:07.789804    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:07.802941    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:08.788440    5738 reflector.go:188] pkg/kubelet/kubelet.go:386: Failed to list *api.Node: Get https://127.0.0.1:6443/api/v1/nodes?fieldSelector=metadata.name%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:08.791626    5738 reflector.go:188] pkg/kubelet/kubelet.go:378: Failed to list *api.Service: Get https://127.0.0.1:6443/api/v1/services?resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:08.804291    5738 reflector.go:188] pkg/kubelet/config/apiserver.go:44: Failed to list *api.Pod: Get https://127.0.0.1:6443/api/v1/pods?fieldSelector=spec.nodeName%3Dk8s-172-17-4-200&resourceVersion=0: dial tcp 127.0.0.1:6443: getsockopt: connection refused
E0430 11:08:09.124069    5738 event.go:208] Unable to write event: 'Post https://127.0.0.1:6443/api/v1/namespaces/default/events: dial tcp 127.0.0.1:6443: getsockopt: connection refused' (may retry after sleeping)
E0430 11:08:09.752514    5738 eviction_manager.go:205] eviction manager: unexpected err: failed GetNode: node 'k8s-172-17-4-200' not found

```

Optional, run in foreground to validate arguments. e.g. － 调试建议，在Terminal上运行，检查参数设置是否正确

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubelet --api-servers=https://127.0.0.1:6443 --enable-debugging-handlers=true --hostname-override=k8s-172-17-4-200 --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --v=2 --cluster-dns=10.123.240.10 --root-dir=/var/lib/kubelet --non-masquerade-cidr=10.0.0.0/8 --cgroup-root=/ --system-cgroups=/system --pod-cidr=10.120.0.0/14 --address=0.0.0.0 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --node-ip=172.17.4.200 --cpu-cfs-quota=true --port=10250 --enable-custom-metrics=true --runtime-cgroups=/docker-daemon --kubelet-cgroups=/kubelet --eviction-hard=memory.available<100Mi --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --anonymous-auth=true --authorization-mode=AlwaysAllow

```

The _process_ & _networking_

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ps -ef | grep kubelet
root     31420     1  1 May03 ?        00:11:13 /opt/kubernetes/server/bin/kubelet --api-servers=http://127.0.0.1:8080 --enable-debugging-handlers=true --hostname-override=172.17.4.200 --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --v=2 --cluster-dns=10.123.240.10 --root-dir=/var/lib/kubele --non-masquerade-cidr=10.0.0.0/8 --cgroup-root=/ --system-cgroups=/system --pod-cidr=10.120.0.0/14 --address=0.0.0.0 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --node-ip=172.17.4.200 --cpu-cfs-quota=true --port=10250 --enable-custom-metrics=true --runtime-cgroups=/docker-daemon --kubelet-cgroups=/kubelet --eviction-hard=memory.available<100Mi --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --anonymous-auth=true --authorization-mode=AlwaysAllow
vagrant   5852  2329  0 11:13 pts/0    00:00:00 grep --color=auto kubelet
vagrant@vagrant-ubuntu-trusty-64:~$
vagrant@vagrant-ubuntu-trusty-64:~$ sudo netstat -tpnl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 0.0.0.0:111             0.0.0.0:*               LISTEN      771/rpcbind     
tcp        0      0 0.0.0.0:44947           0.0.0.0:*               LISTEN      798/rpc.statd   
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      1653/sshd       
tcp        0      0 127.0.0.1:10248         0.0.0.0:*               LISTEN      5738/kubelet    
tcp6       0      0 :::10255                :::*                    LISTEN      5738/kubelet    
tcp6       0      0 :::111                  :::*                    LISTEN      771/rpcbind     
tcp6       0      0 :::22                   :::*                    LISTEN      1653/sshd       
tcp6       0      0 :::4194                 :::*                    LISTEN      5738/kubelet    
tcp6       0      0 :::38663                :::*                    LISTEN      798/rpc.statd   
tcp6       0      0 :::5000                 :::*                    LISTEN      1758/docker-proxy
tcp6       0      0 :::10250                :::*                    LISTEN      5738/kubelet    

```

### Pulling K8s POD image － 获取POD的Docker镜像

在Kubelet的服务配置中指定了镜像的名称 

```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/pause-amd64:3.0
3.0: Pulling from google_containers/pause-amd64
a3ed95caeb02: Pull complete 
f11233434377: Pull complete 
Digest: sha256:163ac025575b775d1c0f9bf0bdd0f086883171eb475b5068e7defa4ca9e76516
Status: Downloaded newer image for gcr.io/google_containers/pause-amd64:3.0
vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ docker images gcr.io/google_containers/pause-amd64
REPOSITORY                             TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/pause-amd64   3.0                 99e59f495ffa        12 months ago       746.9 kB

```

### Generate TLS certificates － 为K8s主控节点生成x509证书

首先，下载OpenVPN提供的生成器（[Github仓库](https://github.com/OpenVPN/easy-rsa/releases)），如果在root下操作，则如后2行命令，在root目录进行

```
vagrant@vagrant-ubuntu-trusty-64:~$ mkdir kube
vagrant@vagrant-ubuntu-trusty-64:~$ curl -L -o /home/vagrant/kube/easy-rsa.tar.gz https://storage.googleapis.com/kubernetes-release/easy-rsa/easy-rsa.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 43444  100 43444    0     0  35046      0  0:00:01  0:00:01 --:--:-- 35035
vagrant@vagrant-ubuntu-trusty-64:~$ ls -1 kube
easy-rsa.tar.gz
vagrant@vagrant-ubuntu-trusty-64:~$ sudo mkdir -p /root/kube
vagrant@vagrant-ubuntu-trusty-64:~$ sudo curl -L -o /root/kube/easy-rsa.tar.gz https://storage.googleapis.com/kubernetes-release/easy-rsa/easy-rsa.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 43444  100 43444    0     0  34209      0  0:00:01  0:00:01 --:--:-- 34261

```

使用saltbase的模版

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/generate-cert
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/saltbase/salt/generate-cert
init.sls
make-ca-cert.sh
make-cert.sh

```

生成ca和key pair

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo DEBUG='true' CERT_DIR=/root CERT_GROUP=root /opt/kubernetes/saltbase/salt/generate-cert/make-ca-cert.sh 172.17.4.200 IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local
+ cert_ip=172.17.4.200
+ extra_sans=IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local
+ cert_dir=/root
+ cert_group=root
+ mkdir -p /root
+ use_cn=false
+ '[' 172.17.4.200 == _use_gce_external_ip_ ']'
+ '[' 172.17.4.200 == _use_aws_external_ip_ ']'
+ '[' 172.17.4.200 == _use_azure_dns_name_ ']'
+ sans=IP:172.17.4.200
+ [[ -n IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local ]]
+ sans=IP:172.17.4.200,IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local
++ mktemp -d -t kubernetes_cacert.XXXXXX
+ tmpdir=/tmp/kubernetes_cacert.vxpao1
+ trap 'rm -rf "${tmpdir}"' EXIT
+ cd /tmp/kubernetes_cacert.vxpao1
+ '[' -f /home/vagrant/kube/easy-rsa.tar.gz ']'
+ curl -L -O https://storage.googleapis.com/kubernetes-release/easy-rsa/easy-rsa.tar.gz
+ tar xzf easy-rsa.tar.gz
+ cd easy-rsa-master/easyrsa3
+ ./easyrsa init-pki
++ date +%s
+ ./easyrsa --batch --req-cn=172.17.4.200@1493677386 build-ca nopass
+ '[' false = true ']'
+ ./easyrsa --subject-alt-name=IP:172.17.4.200,IP:10.123.240.1,DNS:kubernetes,DNS:kubernetes.default,DNS:kubernetes.default.svc,DNS:kubernetes.default.svc.cluster.local build-server-full kubernetes-master nopass
+ cp -p pki/issued/kubernetes-master.crt /root/server.cert
+ cp -p pki/private/kubernetes-master.key /root/server.key
+ ./easyrsa build-client-full kubecfg nopass
+ cp -p pki/ca.crt /root/ca.crt
+ cp -p pki/issued/kubecfg.crt /root/kubecfg.crt
+ cp -p pki/private/kubecfg.key /root/kubecfg.key
+ chgrp root /root/server.key /root/server.cert /root/ca.crt
+ chmod 660 /root/server.key /root/server.cert /root/ca.crt
+ rm -rf /tmp/kubernetes_cacert.vxpao1
vagrant@vagrant-ubuntu-trusty-64:~$
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /root
ca.crt
kube
kubecfg.crt
kubecfg.key
server.cert
server.key

```

复制到社区版指定的位置

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo cp /root/server.cert /root/server.key /root/ca.crt /root/kubecfg.crt /root/kubecfg.key /srv/kubernetes
vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo chmod g-r /srv/kubernetes/*
vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -l /srv/kubernetes
total 56
-rw------- 1 root root 1220 May  1 23:14 ca.crt
-rw------- 1 root root 1220 May  2 02:48 etcd-ca.crt
-rw------- 1 root root 4841 May  2 02:50 etcd-peer.crt
-rw------- 1 root root 1708 May  2 02:50 etcd-peer.key
-rw------- 1 root root 4841 May  2 03:10 kubeapiserver.cert
-rw------- 1 root root 1708 May  2 03:10 kubeapiserver.key
-rw------- 1 root root 4417 May  1 23:14 kubecfg.crt
-rw------- 1 root root 1704 May  1 23:14 kubecfg.key
-rw------- 1 root root 4841 May  1 22:30 server.cert
-rw------- 1 root root 1708 May  1 22:30 server.key

```

### 安装 _etcd_ v3 __POD__

解压 _saltbase_ 的template

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/etcd
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/saltbase/salt/etcd
etcd.manifest
init.sls

```

根据`etcd.manifest`，获取镜像

```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/etcd:3.0.17
3.0.17: Pulling from google_containers/etcd
4b0bc1c4050b: Pull complete 
4adaf48a4994: Pull complete 
bef496bb3347: Pull complete 
Digest: sha256:d83d3545e06fb035db8512e33bd44afb55dea007a3abd7b17742d3ac6d235940
Status: Downloaded newer image for gcr.io/google_containers/etcd:3.0.17
vagrant@vagrant-ubuntu-trusty-64:~$
vagrant@vagrant-ubuntu-trusty-64:~$ docker images gcr.io/google_containers/etcd:3.0.17
REPOSITORY                      TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/etcd   3.0.17              243830dae7dd        9 weeks ago         168.9 MB

```

配置Config

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo touch /var/log/etcd.log
vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo mkdir -p /etc/kubernetes/manifests; sudo sed '/{[%#].*[%#]}/d;s/{{ suffix }}//g;s/{{ pillar.get(\x27etcd_docker_tag\x27, \x272.2.1\x27) }}/3.0.17/;s/{{ cpulimit }}/"250m"/;s/{{ hostname }}/172.17.4.200/g;s/{{ etcd_protocol }}/https/g;s/{{ server_port }}/2380/g;s/{{ port }}/2379/g;s/{{ quota_bytes }}/--quota-backend-bytes 4294967296/;s/{{ cluster_state }}/new/;s%{{ etcd_cluster }}%etcd-172.17.4.200=https://172.17.4.200:2380%;s%{{ etcd_creds }}%--peer-trusted-ca-file /srv/kubernetes/etcd-ca.crt --peer-cert-file /srv/kubernetes/etcd-peer.crt --peer-key-file /srv/kubernetes/etcd-peer.key -peer-client-cert-auth%;s/{{ storage_backend }}/etcd3/;s/{{ pillar.get(\x27etcd_version\x27, \x272.2.1\x27) }}/3.0.17/;s%{{ srv_kube_path }}%/srv/kubernetes%' /opt/kubernetes/saltbase/salt/etcd/etcd.manifest | sudo tee /etc/kubernetes/manifests/etcd.json

{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"etcd-server",
  "namespace": "kube-system"
},
"spec":{
"hostNetwork": true,
"containers":[
    {
    "name": "etcd-container",
    "image": "gcr.io/google_containers/etcd:3.0.17",
    "resources": {
      "requests": {
        "cpu": "250m"
      }
    },
    "command": [
              "/bin/sh",
              "-c",
              "if [ -e /usr/local/bin/migrate-if-needed.sh ]; then /usr/local/bin/migrate-if-needed.sh 1>>/var/log/etcd.log 2>&1; fi; /usr/local/bin/etcd --name etcd-172.17.4.200 --listen-peer-urls https://172.17.4.200:2380 --initial-advertise-peer-urls https://172.17.4.200:2380 --advertise-client-urls http://127.0.0.1:2379 --listen-client-urls http://127.0.0.1:2379 --quota-backend-bytes 4294967296 --data-dir /var/etcd/data --initial-cluster-state new --initial-cluster etcd-172.17.4.200=https://172.17.4.200:2380 --peer-trusted-ca-file /srv/kubernetes/etcd-ca.crt --peer-cert-file /srv/kubernetes/etcd-peer.crt --peer-key-file /srv/kubernetes/etcd-peer.key -peer-client-cert-auth 1>>/var/log/etcd.log 2>&1"
            ],
    "env": [
      { "name": "TARGET_STORAGE",
        "value": "etcd3"
      },
      { "name": "TARGET_VERSION",
        "value": "3.0.17"
      },
      { "name": "DATA_DIRECTORY",
        "value": "/var/etcd/data"
      }
        ],
    "livenessProbe": {
      "httpGet": {
        "host": "127.0.0.1",
        "port": 2379,
        "path": "/health"
      },
      "initialDelaySeconds": 15,
      "timeoutSeconds": 15
    },
    "ports": [
      { "name": "serverport",
        "containerPort": 2380,
        "hostPort": 2380 
      },
      { "name": "clientport",
        "containerPort": 2379,
        "hostPort": 2379
      }
        ],
    "volumeMounts": [
      { "name": "varetcd",
        "mountPath": "/var/etcd",
        "readOnly": false
      },
      { "name": "varlogetcd",
        "mountPath": "/var/log/etcd.log",
        "readOnly": false
      },
      { "name": "etc",
        "mountPath": "/srv/kubernetes",
        "readOnly": false
      }
    ]
    }
],
"volumes":[
  { "name": "varetcd",
    "hostPath": {
        "path": "/mnt/master-pd/var/etcd"}
  },
  { "name": "varlogetcd",
    "hostPath": {
        "path": "/var/log/etcd.log"}
  },
  { "name": "etc",
    "hostPath": {
        "path": "/srv/kubernetes"}
  }
]
}}

```

Kubele会t自动创建etcd POD (另一个registry容器不是本文的内容), Verify:

```
vagrant@vagrant-ubuntu-trusty-64:~$ docker ps
CONTAINER ID        IMAGE                                                                               COMMAND                  CREATED             STATUS              PORTS                    NAMES
2797a1bf0d69        gcr.io/google_containers/etcd:3.0.17                                                "/bin/sh -c 'if [ -e "   28 hours ago        Up 28 hours                                  k8s_etcd-container.b6dcac6c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_9033973a
819a6fe57007        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 28 hours ago        Up 28 hours                                  k8s_POD.d8dbe16c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_e2cda727
85a4fb1e1dcf        registry:2.3.1                                                                      "/bin/registry /etc/d"   11 months ago       Up 30 hours         0.0.0.0:5000->5000/tcp   registry

```

下载Etcd客户端命令工具，请参考[该文档](./get-etcd-v3.md)。Optional，如果etcd容器长时间没有创建，一般还是因为运行参数没有配置正确，可直接在Terminal下检查运行，如：

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /home/vagrant/bin/etcd --name etcd-172.17.4.200 --listen-peer-urls https://172.17.4.200:2380 --initial-advertise-peer-urls https://172.17.4.200:2380 --advertise-client-urls http://127.0.0.1:2379 --listen-client-urls http://127.0.0.1:2379 --quota-backend-bytes 4294967296 --data-dir /var/etcd/data --initial-cluster-state new --initial-cluster etcd-172.17.4.200=https://172.17.4.200:2380 --peer-trusted-ca-file /srv/kubernetes/etcd-ca.crt --peer-cert-file /srv/kubernetes/etcd-peer.crt --peer-key-file /srv/kubernetes/etcd-peer.key -peer-client-cert-auth

```

检查日志

```
vagrant@vagrant-ubuntu-trusty-64:~$ cat /var/log/etcd.log
2017-05-02 04:56:59 Detecting if migration is needed
/var/etcd/data is empty - skipping migration
2017-05-02 04:57:00.040789 I | etcdmain: etcd Version: 3.0.17
2017-05-02 04:57:00.040840 I | etcdmain: Git SHA: cc198e2
2017-05-02 04:57:00.040845 I | etcdmain: Go Version: go1.6.4
2017-05-02 04:57:00.040849 I | etcdmain: Go OS/Arch: linux/amd64
2017-05-02 04:57:00.040854 I | etcdmain: setting maximum number of CPUs to 1, total number of available CPUs is 1
2017-05-02 04:57:00.040882 W | etcdmain: found invalid file/dir version.txt under data dir /var/etcd/data (Ignore this if you are upgrading etcd)
2017-05-02 04:57:00.040914 I | etcdmain: peerTLS: cert = /srv/kubernetes/etcd-peer.crt, key = /srv/kubernetes/etcd-peer.key, ca = , trusted-ca = /srv/kubernetes/etcd-ca.crt, client-cert-auth = true
2017-05-02 04:57:00.041712 I | etcdmain: listening for peers on https://172.17.4.200:2380
2017-05-02 04:57:00.042155 I | etcdmain: listening for client requests on 127.0.0.1:2379
2017-05-02 04:57:00.043937 I | etcdserver: name = etcd-172.17.4.200
2017-05-02 04:57:00.043954 I | etcdserver: data dir = /var/etcd/data
2017-05-02 04:57:00.043960 I | etcdserver: member dir = /var/etcd/data/member
2017-05-02 04:57:00.043964 I | etcdserver: heartbeat = 100ms
2017-05-02 04:57:00.043967 I | etcdserver: election = 1000ms
2017-05-02 04:57:00.043971 I | etcdserver: snapshot count = 10000
2017-05-02 04:57:00.043977 I | etcdserver: advertise client URLs = http://127.0.0.1:2379
2017-05-02 04:57:00.043981 I | etcdserver: initial advertise peer URLs = https://172.17.4.200:2380
2017-05-02 04:57:00.043987 I | etcdserver: initial cluster = etcd-172.17.4.200=https://172.17.4.200:2380
2017-05-02 04:57:00.045344 I | etcdserver: starting member 41dbe89be3c85a2c in cluster 3017c3cafd87fdf
2017-05-02 04:57:00.045396 I | raft: 41dbe89be3c85a2c became follower at term 0
2017-05-02 04:57:00.045415 I | raft: newRaft 41dbe89be3c85a2c [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
2017-05-02 04:57:00.045421 I | raft: 41dbe89be3c85a2c became follower at term 1
2017-05-02 04:57:00.049756 I | etcdserver: starting server... [version: 3.0.17, cluster version: to_be_decided]
2017-05-02 04:57:00.050584 I | membership: added member 41dbe89be3c85a2c [https://172.17.4.200:2380] to cluster 3017c3cafd87fdf
2017-05-02 04:57:01.045729 I | raft: 41dbe89be3c85a2c is starting a new election at term 1
2017-05-02 04:57:01.045775 I | raft: 41dbe89be3c85a2c became candidate at term 2
2017-05-02 04:57:01.045783 I | raft: 41dbe89be3c85a2c received vote from 41dbe89be3c85a2c at term 2
2017-05-02 04:57:01.045802 I | raft: 41dbe89be3c85a2c became leader at term 2
2017-05-02 04:57:01.045811 I | raft: raft.node: 41dbe89be3c85a2c elected leader 41dbe89be3c85a2c at term 2
2017-05-02 04:57:01.046004 I | etcdserver: setting up the initial cluster version to 3.0
2017-05-02 04:57:01.046313 N | membership: set the initial cluster version to 3.0
2017-05-02 04:57:01.046343 I | api: enabled capabilities for version 3.0
2017-05-02 04:57:01.046388 I | etcdserver: published {Name:etcd-172.17.4.200 ClientURLs:[http://127.0.0.1:2379]} to cluster 3017c3cafd87fdf
2017-05-02 04:57:01.046398 I | etcdmain: ready to serve client requests
2017-05-02 04:57:01.047898 N | etcdmain: serving insecure client requests on 127.0.0.1:2379, this is strongly discouraged!

```

测试写和读

```
vagrant@vagrant-ubuntu-trusty-64:~$ ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 put foo "bar"
OK
vagrant@vagrant-ubuntu-trusty-64:~$ ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 get foo
foo
bar
vagrant@vagrant-ubuntu-trusty-64:~$
vagrant@vagrant-ubuntu-trusty-64:~$ curl 127.0.0.1:2379/v2/keys
{"action":"get","node":{"dir":true,"nodes":[{"key":"/foo","value":"bar","modifiedIndex":4,"createdIndex":4},{"key":"/registry","dir":true,"modifiedIndex":5,"createdIndex":5}]}}

```

### Deploy `kube-apiserver` POD － 部署api server容器

使用_saltbase_ template

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/kube-apiserver
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/saltbase/salt/kube-apiserver
abac-authz-policy.jsonl
init.sls
kube-apiserver.manifest

```

Configure配置。[附录command kube-apiserver](#command-kube-apiserver)

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo touch /var/log/kube-apiserver-audit.log
vagrant@vagrant-ubuntu-trusty-64:~$ sudo touch /var/log/kube-apiserver.log
vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo mkdir -p /etc/kubernetes/manifests; sudo sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%{{pillar.*/kube-apiserver:.*}}%gcr.io/google_containers/kube-apiserver:9312bc9403d9a1dda49a2fe70c897deb%;s%{{params}}%--address=127.0.0.1 --storage-backend=etcd3  --etcd-servers=127.0.0.1:2379 --runtime-config=api/all=true --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --admission-control=ResourceQuota,ServiceAccount --max-requests-inflight=400 --target-ram-mb=1024 --service-cluster-ip-range=10.123.240.0/20 --min-request-timeout=1800 --enable-garbage-collector=true --etcd-quorum-read=true --audit-log-path=/var/log/kube-apiserver-audit.log --audit-log-maxage=0 --audit-log-maxbackup=0 --audit-log-maxsize=2000000000 --tls-cert-file=/srv/kubernetes/server.cert --tls-private-key-file=/srv/kubernetes/server.key --kubelet-client-certificate=/srv/kubernetes/kubeapiserver.cert --kubelet-client-key=/srv/kubernetes/kubeapiserver.key --client-ca-file=/srv/kubernetes/ca.crt --secure-port=6443 --token-auth-file=/dev/null --bind-address=0.0.0.0 --v=2 --advertise-address=172.17.4.200 --enable-swagger-ui=true%;s/{{pillar.*}}/true/;s/{{secure_port}}/6443/g;s/{{container_env}}//;s/{{cloud_config_\(mount\|volume\)}}//;s/{{additional_cloud_config_\(mount\|volume\)}}//;s/{{webhook_config_\(mount\|volume\)}}//;s/{{webhook_authn_config_\(mount\|volume\)}}//;s/{{admission_controller_config_\(mount\|volume\)}}//;s/{{image_policy_webhook_config_\(mount\|volume\)}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;s%{{srv_sshproxy_path}}%/srv/sshproxy%g' /opt/kubernetes/saltbase/salt/kube-apiserver/kube-apiserver.manifest | sudo tee /etc/kubernetes/manifests/kube-apiserver.json

{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"kube-apiserver",
  "namespace": "kube-system",
  "labels": {
    "tier": "control-plane",
    "component": "kube-apiserver"
  }
},
"spec":{
"hostNetwork": true,
"containers":[
    {
    "name": "kube-apiserver",
    "image": "gcr.io/google_containers/kube-apiserver:9312bc9403d9a1dda49a2fe70c897deb",
    "resources": {
      "requests": {
        "cpu": "250m"
      }
    },
    "command": [
                 "/bin/sh",
                 "-c",
                 "/usr/local/bin/kube-apiserver --address=127.0.0.1 --storage-backend=etcd3  --etcd-servers=127.0.0.1:2379 --runtime-config=api/all=true --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --admission-control=ResourceQuota,ServiceAccount --max-requests-inflight=400 --target-ram-mb=1024 --service-cluster-ip-range=10.123.240.0/20 --min-request-timeout=1800 --enable-garbage-collector=true --etcd-quorum-read=true --audit-log-path=/var/log/kube-apiserver-audit.log --audit-log-maxage=0 --audit-log-maxbackup=0 --audit-log-maxsize=2000000000 --tls-cert-file=/srv/kubernetes/server.cert --tls-private-key-file=/srv/kubernetes/server.key --kubelet-client-certificate=/srv/kubernetes/kubeapiserver.cert --kubelet-client-key=/srv/kubernetes/kubeapiserver.key --client-ca-file=/srv/kubernetes/ca.crt --secure-port=6443 --token-auth-file=/dev/null --bind-address=0.0.0.0 --v=2 --advertise-address=172.17.4.200 --enable-swagger-ui=true --allow-privileged=true 1>>/var/log/kube-apiserver.log 2>&1"
               ],
    "livenessProbe": {
      "httpGet": {
        "host": "127.0.0.1",
        "port": 8080,
        "path": "/healthz"
      },
      "initialDelaySeconds": 15,
      "timeoutSeconds": 15
    },
    "ports":[
      { "name": "https",
        "containerPort": 6443,
        "hostPort": 6443},{
       "name": "local",
        "containerPort": 8080,
        "hostPort": 8080}
        ],
    "volumeMounts": [
        
        
        
        
        
        
        { "name": "srvkube",
        "mountPath": "/srv/kubernetes",
        "readOnly": true},
        { "name": "logfile",
        "mountPath": "/var/log/kube-apiserver.log",
        "readOnly": false},
        { "name": "auditlogfile",
        "mountPath": "/var/log/kube-apiserver-audit.log",
        "readOnly": false},
        { "name": "etcssl",
        "mountPath": "/etc/ssl",
        "readOnly": true},
        { "name": "varssl",
        "mountPath": "/var/ssl",
        "readOnly": true},
        { "name": "etcopenssl",
        "mountPath": "/etc/openssl",
        "readOnly": true},
        { "name": "etcpki",
        "mountPath": "/etc/srv/pki",
        "readOnly": true},
        { "name": "srvsshproxy",
        "mountPath": "/srv/sshproxy",
        "readOnly": false}
      ]
    }
],
"volumes":[
  
  
  
  
  
  
  { "name": "srvkube",
    "hostPath": {
        "path": "/srv/kubernetes"}
  },
  { "name": "logfile",
    "hostPath": {
        "path": "/var/log/kube-apiserver.log"}
  },
  { "name": "auditlogfile",
    "hostPath": {
        "path": "/var/log/kube-apiserver-audit.log"}
  },
  { "name": "etcssl",
    "hostPath": {
        "path": "/etc/ssl"}
  },
  { "name": "varssl",
    "hostPath": {
        "path": "/var/ssl"}
  },
  { "name": "etcopenssl",
    "hostPath": {
        "path": "/etc/openssl"}
  },
  { "name": "etcpki",
    "hostPath": {
        "path": "/etc/srv/pki"}
  },
  { "name": "srvsshproxy",
    "hostPath": {
        "path": "/srv/sshproxy"}
  }
]
}}

```

一旦保存，即自动创建, Verify:

```
vagrant@vagrant-ubuntu-trusty-64:~$ docker ps
CONTAINER ID        IMAGE                                                                               COMMAND                  CREATED             STATUS              PORTS                    NAMES
1e64747c3b04        gcr.io/google_containers/kube-apiserver:9312bc9403d9a1dda49a2fe70c897deb            "/bin/sh -c '/usr/loc"   5 hours ago         Up 5 hours                                   k8s_kube-apiserver.a08b4be_kube-apiserver-k8s-172-17-4-200_kube-system_6ab64262680e36c450e681d6c8a0f8b8_77c484e7
66c4a6ef7943        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 5 hours ago         Up 5 hours                                   k8s_POD.d8dbe16c_kube-apiserver-k8s-172-17-4-200_kube-system_6ab64262680e36c450e681d6c8a0f8b8_3043972b
2797a1bf0d69        gcr.io/google_containers/etcd:3.0.17                                                "/bin/sh -c 'if [ -e "   28 hours ago        Up 28 hours                                  k8s_etcd-container.b6dcac6c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_9033973a
819a6fe57007        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 28 hours ago        Up 28 hours                                  k8s_POD.d8dbe16c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_e2cda727
85a4fb1e1dcf        registry:2.3.1                                                                      "/bin/registry /etc/d"   11 months ago       Up 30 hours         0.0.0.0:5000->5000/tcp   registry
vagrant@vagrant-ubuntu-trusty-64:~$ 

```

查看日志

```
vagrant@vagrant-ubuntu-trusty-64:~$ tail /var/log/kube-apiserver.log
I0503 09:31:07.944884       6 panics.go:76] PUT /api/v1/namespaces/kube-system/endpoints/kube-controller-manager: (2.056628ms) 200 [[kube-controller-manager/v1.5.7 (linux/amd64) kubernetes/8eb75a5/leader-election] 127.0.0.1:45645]
I0503 09:31:07.975557       6 panics.go:76] GET /api/v1/namespaces/kube-system/endpoints/kube-scheduler: (1.552807ms) 200 [[kube-scheduler/v1.5.7 (linux/amd64) kubernetes/8eb75a5/leader-election] 127.0.0.1:41428]
I0503 09:31:07.977872       6 panics.go:76] PUT /api/v1/namespaces/kube-system/endpoints/kube-scheduler: (1.952104ms) 200 [[kube-scheduler/v1.5.7 (linux/amd64) kubernetes/8eb75a5/leader-election] 127.0.0.1:41428]
I0503 09:31:09.747413       6 panics.go:76] GET /api/v1/namespaces/default/services/kubernetes: (1.489936ms) 200 [[kube-apiserver/v1.5.7 (linux/amd64) kubernetes/8eb75a5] 127.0.0.1:39590]
I0503 09:31:09.749925       6 panics.go:76] GET /api/v1/namespaces/default/endpoints/kubernetes: (1.990057ms) 200 [[kube-apiserver/v1.5.7 (linux/amd64) kubernetes/8eb75a5] 127.0.0.1:39590]
I0503 09:31:09.948046       6 panics.go:76] GET /api/v1/namespaces/kube-system/endpoints/kube-controller-manager: (849.312µs) 200 [[kube-controller-manager/v1.5.7 (linux/amd64) kubernetes/8eb75a5/leader-election] 127.0.0.1:45645]
I0503 09:31:09.950790       6 panics.go:76] PUT /api/v1/namespaces/kube-system/endpoints/kube-controller-manager: (2.345275ms) 200 [[kube-controller-manager/v1.5.7 (linux/amd64) kubernetes/8eb75a5/leader-election] 127.0.0.1:45645]
I0503 09:31:09.982930       6 panics.go:76] GET /api/v1/namespaces/kube-system/endpoints/kube-scheduler: (2.876124ms) 200 [[kube-scheduler/v1.5.7 (linux/amd64) kubernetes/8eb75a5/leader-election] 127.0.0.1:41428]
I0503 09:31:09.986233       6 panics.go:76] PUT /api/v1/namespaces/kube-system/endpoints/kube-scheduler: (1.500931ms) 200 [[kube-scheduler/v1.5.7 (linux/amd64) kubernetes/8eb75a5/leader-election] 127.0.0.1:41428]
I0503 09:31:10.160924       6 panics.go:76] GET /healthz: (71.214µs) 200 [[Go-http-client/1.1] 127.0.0.1:45653]

```

使用curl验证

```
vagrant@vagrant-ubuntu-trusty-64:~$ curl http://127.0.0.1:8080
{
  "paths": [
    "/api",
    "/api/v1",
    "/apis",
    "/apis/apps",
    "/apis/apps/v1beta1",
    "/apis/authentication.k8s.io",
    "/apis/authentication.k8s.io/v1beta1",
    "/apis/authorization.k8s.io",
    "/apis/authorization.k8s.io/v1beta1",
    "/apis/autoscaling",
    "/apis/autoscaling/v1",
    "/apis/batch",
    "/apis/batch/v1",
    "/apis/batch/v2alpha1",
    "/apis/certificates.k8s.io",
    "/apis/certificates.k8s.io/v1alpha1",
    "/apis/extensions",
    "/apis/extensions/v1beta1",
    "/apis/policy",
    "/apis/policy/v1beta1",
    "/apis/rbac.authorization.k8s.io",
    "/apis/rbac.authorization.k8s.io/v1alpha1",
    "/apis/storage.k8s.io",
    "/apis/storage.k8s.io/v1beta1",
    "/healthz",
    "/healthz/poststarthook/bootstrap-controller",
    "/healthz/poststarthook/extensions/third-party-resources",
    "/healthz/poststarthook/rbac/bootstrap-roles",
    "/logs",
    "/metrics",
    "/swagger-ui/",
    "/swaggerapi/",
    "/ui/",
    "/version"
  ]
}
vagrant@vagrant-ubuntu-trusty-64:~$ curl http://127.0.0.1:8080/api
{
  "kind": "APIVersions",
  "versions": [
    "v1"
  ],
  "serverAddressByClientCIDRs": [
    {
      "clientCIDR": "0.0.0.0/0",
      "serverAddress": "172.17.4.200:6443"
    }
  ]
}
vagrant@vagrant-ubuntu-trusty-64:~$ curl http://127.0.0.1:8080/healthz
ok

```

### 配置kubectl到该Kubernetes集群

为kubectl配置当前kube-apiserver的上下文

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl config set-cluster kube --server=https://172.17.4.200:6443
vagrant@vagrant-ubuntu-trusty-64:~$ encoded=$(sudo base64 -w0 /srv/kubernetes/ca.crt); sudo /opt/kubernetes/server/bin/kubectl config set clusters.kube.certificate-authority-data $encoded
Property "clusters.kube.certificate-authority-data" set.
vagrant@vagrant-ubuntu-trusty-64:~$ encoded=$(sudo base64 -w0 /srv/kubernetes/kubecfg.crt); sudo /opt/kubernetes/server/bin/kubectl config set users.admin.client-certificate-data $encoded
Property "users.admin.client-certificate-data" set.
vagrant@vagrant-ubuntu-trusty-64:~$ encoded=$(sudo base64 -w0 /srv/kubernetes/kubecfg.key); sudo /opt/kubernetes/server/bin/kubectl config set users.admin.client-key-data $encoded
Property "users.admin.client-key-data" set.
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl config set-context kube-admin --cluster=kube --user=admin
Context "kube-admin" set.
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl config use-context kube-admin
Switched to context "kube-admin".

```

配置文件

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl config view
vagrant@vagrant-ubuntu-trusty-64:~$ cat .kube/config

```

使用kubectl查看

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get all --all-namespaces
NAMESPACE   NAME             CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
default     svc/kubernetes   10.123.240.1   <none>        443/TCP   1h

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get all --all-namespaces
NAMESPACE     NAME                                 READY     STATUS    RESTARTS   AGE
kube-system   po/etcd-server-k8s-172-17-4-200      1/1       Running   0          2m
kube-system   po/kube-apiserver-k8s-172-17-4-200   1/1       Running   17         2m

NAMESPACE   NAME             CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
default     svc/kubernetes   10.123.240.1   <none>        443/TCP   1h
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get nodes
NAME               STATUS    AGE
k8s-172-17-4-200   Ready     3m

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get all --all-namespaces
NAMESPACE     NAME                                          READY     STATUS    RESTARTS   AGE
kube-system   po/etcd-server-k8s-172-17-4-200               1/1       Running   0          5h
kube-system   po/kube-apiserver-k8s-172-17-4-200            1/1       Running   17         5h

NAMESPACE   NAME             CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
default     svc/kubernetes   10.123.240.1   <none>        443/TCP   6h

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get nodes
NAME               STATUS    AGE
k8s-172-17-4-200   Ready     5h

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl api-versions
apps/v1beta1
authentication.k8s.io/v1beta1
authorization.k8s.io/v1beta1
autoscaling/v1
batch/v1
batch/v2alpha1
certificates.k8s.io/v1alpha1
extensions/v1beta1
policy/v1beta1
rbac.authorization.k8s.io/v1alpha1
storage.k8s.io/v1beta1
v1
 

```

### Deploy `kube-controller-manager` POD

Extract _saltbase_ template

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/kube-controller-manager
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/saltbase/salt/kube-controller-manager
init.sls
kube-controller-manager.manifest

```

Configure配置。[附录command kube-controller-manager](#command-kube-controller-manager)

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo touch /var/log/kube-controller-manager.log
vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo mkdir -p /etc/kubernetes/manifests; sudo sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%{{pillar.*/kube-controller-manager:.*}}%gcr.io/google_containers/kube-controller-manager:0005ed016737cb2c14929aa8efdf5be7%;s%{{params}}%--cluster-name=kubernetes --cluster-cidr=10.120.0.0/14 --service-cluster-ip-range=10.123.240.0/20 --terminated-pod-gc-threshold=12500 --enable-garbage-collector=true --master=127.0.0.1:8080 --address=127.0.0.1 --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --root-ca-file=/srv/kubernetes/ca.crt --service-account-private-key-file=/srv/kubernetes/server.key --v=2 --enable-hostpath-provisioner=true%;s/{{container_env}}//;s/{{cloud_config_\(mount\|volume\)}}//;s/{{additional_cloud_config_\(mount\|volume\)}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;' /opt/kubernetes/saltbase/salt/kube-controller-manager/kube-controller-manager.manifest | sudo tee /etc/kubernetes/manifests/kube-controller-manager.json

{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"kube-controller-manager",
  "namespace": "kube-system",
  "labels": {
    "tier": "control-plane",
    "component": "kube-controller-manager"
  }
},
"spec":{
"hostNetwork": true,
"containers":[
    {
    "name": "kube-controller-manager",
    "image": "gcr.io/google_containers/kube-controller-manager:0005ed016737cb2c14929aa8efdf5be7",
    "resources": {
      "requests": {
        "cpu": "200m"
      }
    },
    "command": [
                 "/bin/sh",
                 "-c",
                 "/usr/local/bin/kube-controller-manager --cluster-name=kubernetes --cluster-cidr=10.120.0.0/14 --service-cluster-ip-range=10.123.240.0/20 --terminated-pod-gc-threshold=12500 --enable-garbage-collector=true --master=127.0.0.1:8080 --address=127.0.0.1 --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --root-ca-file=/srv/kubernetes/ca.crt --service-account-private-key-file=/srv/kubernetes/server.key --v=2 --enable-hostpath-provisioner=true 1>>/var/log/kube-controller-manager.log 2>&1"
               ],
    "livenessProbe": {
      "httpGet": {
        "host": "127.0.0.1",
        "port": 10252,
        "path": "/healthz"
      },
      "initialDelaySeconds": 15,
      "timeoutSeconds": 15
    },
    "volumeMounts": [
        
        
        { "name": "srvkube",
        "mountPath": "/srv/kubernetes",
        "readOnly": true},
        { "name": "logfile",
        "mountPath": "/var/log/kube-controller-manager.log",
        "readOnly": false},
        { "name": "etcssl",
        "mountPath": "/etc/ssl",
        "readOnly": true},
        { "name": "varssl",
        "mountPath": "/var/ssl",
        "readOnly": true},
        { "name": "etcopenssl",
        "mountPath": "/etc/openssl",
        "readOnly": true},
        { "name": "etcpki",
        "mountPath": "/etc/pki",
        "readOnly": true}
      ]
    }
],
"volumes":[
  
  
  { "name": "srvkube",
    "hostPath": {
        "path": "/srv/kubernetes"}
  },
  { "name": "logfile",
    "hostPath": {
        "path": "/var/log/kube-controller-manager.log"}
  },
  { "name": "etcssl",
    "hostPath": {
        "path": "/etc/ssl"}
  },
  { "name": "varssl",
    "hostPath": {
        "path": "/var/ssl"}
  },
  { "name": "etcopenssl",
    "hostPath": {
        "path": "/etc/openssl"}
  },
  { "name": "etcpki",
    "hostPath": {
        "path": "/etc/pki"}
  }
]
}}

```

一旦保存即自动加载, Verify:

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get all --all-namespaces
NAMESPACE     NAME                                          READY     STATUS    RESTARTS   AGE
kube-system   po/etcd-server-k8s-172-17-4-200               1/1       Running   0          1h
kube-system   po/kube-apiserver-k8s-172-17-4-200            1/1       Running   17         1h
kube-system   po/kube-controller-manager-k8s-172-17-4-200   1/1       Running   0          2m

NAMESPACE   NAME             CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
default     svc/kubernetes   10.123.240.1   <none>        443/TCP   3h

```

查看容器

```
vagrant@vagrant-ubuntu-trusty-64:~$ docker ps
CONTAINER ID        IMAGE                                                                               COMMAND                  CREATED              STATUS              PORTS                    NAMES
a127610b4081        gcr.io/google_containers/kube-controller-manager:0005ed016737cb2c14929aa8efdf5be7   "/bin/sh -c '/usr/loc"   About a minute ago   Up About a minute                            k8s_kube-controller-manager.af9e739e_kube-controller-manager-k8s-172-17-4-200_kube-system_3005a1f4178a1bdcad6689354a58bb93_19ebb494
9657ea9f05ed        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 About a minute ago   Up About a minute                            k8s_POD.d8dbe16c_kube-controller-manager-k8s-172-17-4-200_kube-system_3005a1f4178a1bdcad6689354a58bb93_48a5ad56
1e64747c3b04        gcr.io/google_containers/kube-apiserver:9312bc9403d9a1dda49a2fe70c897deb            "/bin/sh -c '/usr/loc"   About an hour ago    Up About an hour                             k8s_kube-apiserver.a08b4be_kube-apiserver-k8s-172-17-4-200_kube-system_6ab64262680e36c450e681d6c8a0f8b8_77c484e7
66c4a6ef7943        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 2 hours ago          Up 2 hours                                   k8s_POD.d8dbe16c_kube-apiserver-k8s-172-17-4-200_kube-system_6ab64262680e36c450e681d6c8a0f8b8_3043972b
2797a1bf0d69        gcr.io/google_containers/etcd:3.0.17                                                "/bin/sh -c 'if [ -e "   25 hours ago         Up 25 hours                                  k8s_etcd-container.b6dcac6c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_9033973a
819a6fe57007        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 25 hours ago         Up 25 hours                                  k8s_POD.d8dbe16c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_e2cda727
85a4fb1e1dcf        registry:2.3.1                                                                      "/bin/registry /etc/d"   11 months ago        Up 27 hours         0.0.0.0:5000->5000/tcp   registry

```

查看日志

```
vagrant@vagrant-ubuntu-trusty-64:~$ tail /var/log/kube-controller-manager.log
I0503 06:11:41.232443       7 serviceaccounts_controller.go:120] Starting ServiceAccount controller
I0503 06:11:41.241516       7 garbagecollector.go:766] Garbage Collector: Initializing
E0503 06:11:41.289548       7 actual_state_of_world.go:462] Failed to set statusUpdateNeeded to needed true because nodeName="k8s-172-17-4-200"  does not exist
I0503 06:11:41.381400       7 nodecontroller.go:419] NodeController observed a new Node: "k8s-172-17-4-200"
I0503 06:11:41.381418       7 controller_utils.go:274] Recording Registered Node k8s-172-17-4-200 in NodeController event message for node k8s-172-17-4-200
I0503 06:11:41.381439       7 nodecontroller.go:429] Initializing eviction metric for zone: 
W0503 06:11:41.381457       7 nodecontroller.go:678] Missing timestamp for Node k8s-172-17-4-200. Assuming now as a timestamp.
I0503 06:11:41.381497       7 nodecontroller.go:608] NodeController detected that zone  is now in state Normal.
I0503 06:11:41.381632       7 event.go:217] Event(api.ObjectReference{Kind:"Node", Namespace:"", Name:"k8s-172-17-4-200", UID:"5de1e986-2fb8-11e7-b234-080027360e7e", APIVersion:"", ResourceVersion:"", FieldPath:""}): type: 'Normal' reason: 'RegisteredNode' Node k8s-172-17-4-200 event: Registered Node k8s-172-17-4-200 in NodeController
I0503 06:11:51.242872       7 garbagecollector.go:780] Garbage Collector: All monitored resources synced. Proceeding to collect garbage

```

### Deploy `kube-scheduler` POD

Extract _saltbase_ template

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/kube-scheduler
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/saltbase/salt/kube-scheduler
init.sls
kube-scheduler.manifest

```

Configure配置。[附录command kube-scheduler](#command-kube-scheduler)

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo touch /var/log/kube-scheduler.log
vagrant@vagrant-ubuntu-trusty-64:~$
vagrant@vagrant-ubuntu-trusty-64:~$ sudo mkdir -p /etc/kubernetes/manifests; sudo sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%{{pillar.*/kube-scheduler:.*}}%gcr.io/google_containers/kube-scheduler:5e5a1fe4ee49b94c0f3d2e745a48cb65%;s%{{params}}%--algorithm-provider=DefaultProvider --master=127.0.0.1:8080 --address=127.0.0.1 --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --v=2%;s/{{container_env}}//;s/{{cloud_config_\(mount\|volume\)}}//;s/{{additional_cloud_config_\(mount\|volume\)}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;' /opt/kubernetes/saltbase/salt/kube-scheduler/kube-scheduler.manifest | sudo tee /etc/kubernetes/manifests/kube-scheduler.json

{
"apiVersion": "v1",
"kind": "Pod",
"metadata": {
  "name":"kube-scheduler",
  "namespace": "kube-system",
  "labels": {
    "tier": "control-plane",
    "component": "kube-scheduler"
  }
},
"spec":{
"hostNetwork": true,
"containers":[
    {
    "name": "kube-scheduler",
    "image": "gcr.io/google_containers/kube-scheduler:5e5a1fe4ee49b94c0f3d2e745a48cb65",
    "resources": {
      "requests": {
        "cpu": "100m"
      }
    },
    "command": [
                 "/bin/sh",
                 "-c",
                 "/usr/local/bin/kube-scheduler --master=127.0.0.1:8080 --algorithm-provider=DefaultProvider --master=127.0.0.1:8080 --address=127.0.0.1 --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --v=2 1>>/var/log/kube-scheduler.log 2>&1"
               ],
    "livenessProbe": {
      "httpGet": {
        "host": "127.0.0.1",
        "port": 10251,
        "path": "/healthz"
      },
      "initialDelaySeconds": 15,
      "timeoutSeconds": 15
    },
    "volumeMounts": [
        {
          "name": "logfile",
          "mountPath": "/var/log/kube-scheduler.log",
          "readOnly": false
        }
      ]
    }
],
"volumes":[
  { "name": "logfile",
    "hostPath": {
        "path": "/var/log/kube-scheduler.log"}
  }
]
}}
```

Verify验证

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get all --all-namespaces
NAMESPACE     NAME                                          READY     STATUS    RESTARTS   AGE
kube-system   po/etcd-server-k8s-172-17-4-200               1/1       Running   0          2h
kube-system   po/kube-apiserver-k8s-172-17-4-200            1/1       Running   17         2h
kube-system   po/kube-controller-manager-k8s-172-17-4-200   1/1       Running   0          26m
kube-system   po/kube-scheduler-k8s-172-17-4-200            1/1       Running   0          41s

NAMESPACE   NAME             CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
default     svc/kubernetes   10.123.240.1   <none>        443/TCP   3h

```

查看容器

```
vagrant@vagrant-ubuntu-trusty-64:~$ docker ps
CONTAINER ID        IMAGE                                                                               COMMAND                  CREATED             STATUS              PORTS                    NAMES
0dc677a3509a        gcr.io/google_containers/kube-scheduler:5e5a1fe4ee49b94c0f3d2e745a48cb65            "/bin/sh -c '/usr/loc"   54 seconds ago      Up 53 seconds                                k8s_kube-scheduler.70369ff_kube-scheduler-k8s-172-17-4-200_kube-system_e200995c695a2afb88f6810e52cc6d39_50060874
fbe6949140be        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 54 seconds ago      Up 53 seconds                                k8s_POD.d8dbe16c_kube-scheduler-k8s-172-17-4-200_kube-system_e200995c695a2afb88f6810e52cc6d39_e6e9e3dd
a127610b4081        gcr.io/google_containers/kube-controller-manager:0005ed016737cb2c14929aa8efdf5be7   "/bin/sh -c '/usr/loc"   26 minutes ago      Up 26 minutes                                k8s_kube-controller-manager.af9e739e_kube-controller-manager-k8s-172-17-4-200_kube-system_3005a1f4178a1bdcad6689354a58bb93_19ebb494
9657ea9f05ed        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 26 minutes ago      Up 26 minutes                                k8s_POD.d8dbe16c_kube-controller-manager-k8s-172-17-4-200_kube-system_3005a1f4178a1bdcad6689354a58bb93_48a5ad56
1e64747c3b04        gcr.io/google_containers/kube-apiserver:9312bc9403d9a1dda49a2fe70c897deb            "/bin/sh -c '/usr/loc"   2 hours ago         Up 2 hours                                   k8s_kube-apiserver.a08b4be_kube-apiserver-k8s-172-17-4-200_kube-system_6ab64262680e36c450e681d6c8a0f8b8_77c484e7
66c4a6ef7943        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 2 hours ago         Up 2 hours                                   k8s_POD.d8dbe16c_kube-apiserver-k8s-172-17-4-200_kube-system_6ab64262680e36c450e681d6c8a0f8b8_3043972b
2797a1bf0d69        gcr.io/google_containers/etcd:3.0.17                                                "/bin/sh -c 'if [ -e "   25 hours ago        Up 25 hours                                  k8s_etcd-container.b6dcac6c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_9033973a
819a6fe57007        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 25 hours ago        Up 25 hours                                  k8s_POD.d8dbe16c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_e2cda727
85a4fb1e1dcf        registry:2.3.1                                                                      "/bin/registry /etc/d"   11 months ago       Up 27 hours         0.0.0.0:5000->5000/tcp   registry

```

日志

```
vagrant@vagrant-ubuntu-trusty-64:~$ tail /var/log/kube-scheduler.log
I0503 06:37:24.103439       6 feature_gate.go:189] feature gates: map[AllAlpha:true DynamicKubeletConfig:true DynamicVolumeProvisioning:true StreamingProxyRedirects:true ExperimentalHostUserNamespaceDefaulting:true ExperimentalCriticalPodAnnotation:true]
I0503 06:37:24.104177       6 factory.go:300] Creating scheduler from algorithm provider 'DefaultProvider'
I0503 06:37:24.104204       6 factory.go:346] creating scheduler with fit predicates 'map[CheckNodeMemoryPressure:{} CheckNodeDiskPressure:{} NoVolumeZoneConflict:{} MatchInterPodAffinity:{} NoDiskConflict:{} PodToleratesNodeTaints:{} MaxEBSVolumeCount:{} MaxGCEPDVolumeCount:{} GeneralPredicates:{}]' and priority functions 'map[NodeAffinityPriority:{} TaintTolerationPriority:{} SelectorSpreadPriority:{} InterPodAffinityPriority:{} LeastRequestedPriority:{} BalancedResourceAllocation:{} NodePreferAvoidPodsPriority:{}]
I0503 06:37:24.137202       6 leaderelection.go:188] sucessfully acquired lease kube-system/kube-scheduler
I0503 06:37:24.137476       6 event.go:217] Event(api.ObjectReference{Kind:"Endpoints", Namespace:"kube-system", Name:"kube-scheduler", UID:"f78e8885-2fca-11e7-924e-080027360e7e", APIVersion:"v1", ResourceVersion:"1704", FieldPath:""}): type: 'Normal' reason: 'LeaderElection' vagrant-ubuntu-trusty-64 became leader

```

### Deploy `kube-proxy` POD

Extract _saltbase_ template

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/kube-proxy
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/saltbase/salt/kube-proxy
init.sls
kubeconfig
kube-proxy.manifest

```

提供容器网络服务，须在每个work节点上运行，以client方式访问kube-apiserver, 需要kubeconfig，从kubectl节点复制

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo mkdir -p /var/lib/kube-proxy/ && sudo cp .kube/config /var/lib/kube-proxy/kubeconfig

```

Configure配置。[附录command kube-proxy](#command-kube-proxy)

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo touch /var/log/kube-proxy.log

vagrant@vagrant-ubuntu-trusty-64:~$ sudo mkdir -p /etc/kubernetes/manifests; sudo sed '/{[%#].*[%#]}/d;1,/^# test_args.*$/d;s%{{pillar.*/kube-proxy:.*}}%gcr.io/google_containers/kube-proxy:690a68ee4bacc47b0115413489185f3d%;s/{{ cpurequest }}/100m/;s%{{api_servers_with_port}}%--master=https://172.17.4.200:6443%;s%{{kubeconfig}}%--kubeconfig=/var/lib/kube-proxy/kubeconfig%;s%{{cluster_cidr}}%10.120.0.0/14%;s%{{params}}%--feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --v=2%;s/{{container_env}}//;s%{{srv_kube_path}}%/srv/kubernetes%g;' /opt/kubernetes/saltbase/salt/kube-proxy/kube-proxy.manifest | sudo tee /etc/kubernetes/manifests/kube-proxy.yaml

# kube-proxy podspec
apiVersion: v1
kind: Pod
metadata:
  name: kube-proxy
  namespace: kube-system
  # This annotation lowers the possibility that kube-proxy gets evicted when the
  # node is under memory pressure, and prioritizes it for admission, even if
  # the node is under memory pressure.
  # Note that kube-proxy runs as a static pod so this annotation does NOT have
  # any effect on rescheduler (default scheduler and rescheduler are not
  # involved in scheduling kube-proxy).
  annotations:
    scheduler.alpha.kubernetes.io/critical-pod: ''
  labels:
    tier: node
    component: kube-proxy
spec:
  hostNetwork: true
  containers:
  - name: kube-proxy
    image: gcr.io/google_containers/kube-proxy:690a68ee4bacc47b0115413489185f3d
    resources:
      requests:
        cpu: 100m
    command:
    - /bin/sh
    - -c
    - kube-proxy --master=https://172.17.4.200:6443 --kubeconfig=/var/lib/kube-proxy/kubeconfig 10.120.0.0/14 --resource-container="" --feature-gates=AllAlpha=true,DynamicKubeletConfig=true,ExperimentalCriticalPodAnnotation=true,ExperimentalHostUserNamespaceDefaulting=true,StreamingProxyRedirects=true --v=2 1>>/var/log/kube-proxy.log 2>&1
    securityContext:
      privileged: true
    volumeMounts:
    - mountPath: /etc/ssl/certs
      name: ssl-certs-host
      readOnly: true
    - mountPath: /var/log
      name: varlog
      readOnly: false
    - mountPath: /var/lib/kube-proxy/kubeconfig
      name: kubeconfig
      readOnly: false
  volumes:
  - hostPath:
      path: /usr/share/ca-certificates
    name: ssl-certs-host
  - hostPath:
      path: /var/lib/kube-proxy/kubeconfig
    name: kubeconfig
  - hostPath:
      path: /var/log
    name: varlog
```

保存后即自动加载, Verify:

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get all --all-namespaces
NAMESPACE     NAME                                          READY     STATUS    RESTARTS   AGE
kube-system   po/etcd-server-k8s-172-17-4-200               1/1       Running   0          5h
kube-system   po/kube-apiserver-k8s-172-17-4-200            1/1       Running   17         5h
kube-system   po/kube-controller-manager-k8s-172-17-4-200   1/1       Running   0          3h
kube-system   po/kube-proxy-k8s-172-17-4-200                1/1       Running   0          30s
kube-system   po/kube-scheduler-k8s-172-17-4-200            1/1       Running   0          2h

NAMESPACE   NAME             CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
default     svc/kubernetes   10.123.240.1   <none>        443/TCP   6h

```

查看容器

```
vagrant@vagrant-ubuntu-trusty-64:~$ docker ps
CONTAINER ID        IMAGE                                                                               COMMAND                  CREATED              STATUS              PORTS                    NAMES
99b699d28acd        gcr.io/google_containers/kube-proxy:690a68ee4bacc47b0115413489185f3d                "/bin/sh -c 'kube-pro"   About a minute ago   Up About a minute                            k8s_kube-proxy.f64956ce_kube-proxy-k8s-172-17-4-200_kube-system_c0c223ef854d23a0bf73614cc14a17df_a16131c3
328ebdb97ffb        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 About a minute ago   Up About a minute                            k8s_POD.d8dbe16c_kube-proxy-k8s-172-17-4-200_kube-system_c0c223ef854d23a0bf73614cc14a17df_20bc0244
0dc677a3509a        gcr.io/google_containers/kube-scheduler:5e5a1fe4ee49b94c0f3d2e745a48cb65            "/bin/sh -c '/usr/loc"   2 hours ago          Up 2 hours                                   k8s_kube-scheduler.70369ff_kube-scheduler-k8s-172-17-4-200_kube-system_e200995c695a2afb88f6810e52cc6d39_50060874
fbe6949140be        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 2 hours ago          Up 2 hours                                   k8s_POD.d8dbe16c_kube-scheduler-k8s-172-17-4-200_kube-system_e200995c695a2afb88f6810e52cc6d39_e6e9e3dd
a127610b4081        gcr.io/google_containers/kube-controller-manager:0005ed016737cb2c14929aa8efdf5be7   "/bin/sh -c '/usr/loc"   3 hours ago          Up 3 hours                                   k8s_kube-controller-manager.af9e739e_kube-controller-manager-k8s-172-17-4-200_kube-system_3005a1f4178a1bdcad6689354a58bb93_19ebb494
9657ea9f05ed        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 3 hours ago          Up 3 hours                                   k8s_POD.d8dbe16c_kube-controller-manager-k8s-172-17-4-200_kube-system_3005a1f4178a1bdcad6689354a58bb93_48a5ad56
1e64747c3b04        gcr.io/google_containers/kube-apiserver:9312bc9403d9a1dda49a2fe70c897deb            "/bin/sh -c '/usr/loc"   5 hours ago          Up 5 hours                                   k8s_kube-apiserver.a08b4be_kube-apiserver-k8s-172-17-4-200_kube-system_6ab64262680e36c450e681d6c8a0f8b8_77c484e7
66c4a6ef7943        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 5 hours ago          Up 5 hours                                   k8s_POD.d8dbe16c_kube-apiserver-k8s-172-17-4-200_kube-system_6ab64262680e36c450e681d6c8a0f8b8_3043972b
2797a1bf0d69        gcr.io/google_containers/etcd:3.0.17                                                "/bin/sh -c 'if [ -e "   28 hours ago         Up 28 hours                                  k8s_etcd-container.b6dcac6c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_9033973a
819a6fe57007        gcr.io/google_containers/pause-amd64:3.0                                            "/pause"                 28 hours ago         Up 28 hours                                  k8s_POD.d8dbe16c_etcd-server-k8s-172-17-4-200_kube-system_f89f83ed14e46c54fcfedac316dfad40_e2cda727
85a4fb1e1dcf        registry:2.3.1                                                                      "/bin/registry /etc/d"   11 months ago        Up 30 hours         0.0.0.0:5000->5000/tcp   registry

```

查看日志

```
vagrant@vagrant-ubuntu-trusty-64:~$ tail /var/log/kube-proxy.log
W0503 09:25:47.570371       5 proxier.go:249] invalid nodeIP, initialize kube-proxy with 127.0.0.1 as nodeIP
W0503 09:25:47.570379       5 proxier.go:254] clusterCIDR not specified, unable to distinguish between internal and external traffic
I0503 09:25:47.570393       5 server.go:227] Tearing down userspace rules.
I0503 09:25:47.571887       5 healthcheck.go:119] Initializing kube-proxy health checker
I0503 09:25:47.583074       5 conntrack.go:81] Set sysctl 'net/netfilter/nf_conntrack_max' to 131072
I0503 09:25:47.583323       5 conntrack.go:66] Setting conntrack hashsize to 32768
I0503 09:25:47.583452       5 conntrack.go:81] Set sysctl 'net/netfilter/nf_conntrack_tcp_timeout_established' to 86400
I0503 09:25:47.583466       5 conntrack.go:81] Set sysctl 'net/netfilter/nf_conntrack_tcp_timeout_close_wait' to 3600
I0503 09:25:47.597923       5 proxier.go:475] Adding new service "default/kubernetes:https" at 10.123.240.1:443/TCP
I0503 09:25:47.597923       5 proxier.go:840] Not syncing iptables until Services and Endpoints have been received from master

```

## Addons

Packages 

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /99-mirror/https%3A%2F%2Fdl.k8s.io/v1.5.7/kubernetes.tar.gz 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/
addons
client
cluster
docs
examples
federation
kubernetes-src.tar.gz
LICENSES
README.md
server
third_party
Vagrantfile
version

```

### DNS

Extract模版

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /99-mirror/https%3A%2F%2Fdl.k8s.io/v1.5.7/kubernetes.tar.gz kubernetes/cluster/addons/dns

```

Pull镜像

```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/dnsmasq-metrics-amd64:1.0.1
1.0.1: Pulling from google_containers/dnsmasq-metrics-amd64
627beaf3eaaf: Pull complete 
dad6277b0212: Pull complete 
Digest: sha256:6453b3ee4f5455657133ada25c858ba695d2f90db69f3e8e69b3d9a2f6392a66
Status: Downloaded newer image for gcr.io/google_containers/dnsmasq-metrics-amd64:1.0.1
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/exechealthz-amd64:1.2
1.2: Pulling from google_containers/exechealthz-amd64
8ddc19f16526: Pull complete 
a3ed95caeb02: Pull complete 
7d1ee54af137: Pull complete 
476d09449781: Pull complete 
Digest: sha256:503e158c3f65ed7399f54010571c7c977ade7fe59010695f48d9650d83488c0a
Status: Downloaded newer image for gcr.io/google_containers/exechealthz-amd64:1.2
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/kube-dnsmasq-amd64:1.4.1
1.4.1: Pulling from google_containers/kube-dnsmasq-amd64
709515475419: Pull complete 
aec729922351: Pull complete 
fae2a7190364: Pull complete 
Digest: sha256:6732224a130aec34784396ecfa3472c7997162f94274f316d9ecd48607b0dc84
Status: Downloaded newer image for gcr.io/google_containers/kube-dnsmasq-amd64:1.4.1
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull  gcr.io/google_containers/kubedns-amd64:1.9
1.9: Pulling from google_containers/kubedns-amd64
8ddc19f16526: Already exists 
a3ed95caeb02: Already exists 
acc68ed00435: Pull complete 
Digest: sha256:3d3d67f519300af646e00adcf860b2f380d35ed4364e550d74002dadace20ead
Status: Downloaded newer image for gcr.io/google_containers/kubedns-amd64:1.9

```

Wordload _dns-deployment_

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo sed "s/\$DNS_DOMAIN/cluster.local/g" /opt/kubernetes/cluster/addons/dns/skydns-rc.yaml.sed | sudo tee /root/cluster%2Faddons%2Fdns%2Fskydns-rc.yaml
# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# TODO - At some point, we need to rename all skydns-*.yaml.* files to kubedns-*.yaml.*
# Should keep target in cluster/addons/dns-horizontal-autoscaler/dns-horizontal-autoscaler.yaml
# in sync with this file.

# Warning: This is a file generated from the base underscore template file: skydns-rc.yaml.base

apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
spec:
  # replicas: not specified here:
  # 1. In order to make Addon Manager do not reconcile this replicas parameter.
  # 2. Default is 1.
  # 3. Will be tuned in real time if DNS horizontal auto-scaling is turned on.
  strategy:
    rollingUpdate:
      maxSurge: 10%
      maxUnavailable: 0
  selector:
    matchLabels:
      k8s-app: kube-dns
  template:
    metadata:
      labels:
        k8s-app: kube-dns
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
        scheduler.alpha.kubernetes.io/tolerations: '[{"key":"CriticalAddonsOnly", "operator":"Exists"}]'
    spec:
      containers:
      - name: kubedns
        image: gcr.io/google_containers/kubedns-amd64:1.9
        resources:
          # TODO: Set memory limits when we've profiled the container for large
          # clusters, then set request = limit to keep this container in
          # guaranteed class. Currently, this container falls into the
          # "burstable" category so the kubelet doesn't backoff from restarting it.
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        livenessProbe:
          httpGet:
            path: /healthz-kubedns
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        readinessProbe:
          httpGet:
            path: /readiness
            port: 8081
            scheme: HTTP
          # we poll on pod startup for the Kubernetes master service and
          # only setup the /readiness HTTP server once that's available.
          initialDelaySeconds: 3
          timeoutSeconds: 5
        args:
        - --domain=cluster.local.
        - --dns-port=10053
        - --config-map=kube-dns
        # This should be set to v=2 only after the new image (cut from 1.5) has
        # been released, otherwise we will flood the logs.
        - --v=0
        env:
        - name: PROMETHEUS_PORT
          value: "10055"
        ports:
        - containerPort: 10053
          name: dns-local
          protocol: UDP
        - containerPort: 10053
          name: dns-tcp-local
          protocol: TCP
        - containerPort: 10055
          name: metrics
          protocol: TCP
      - name: dnsmasq
        image: gcr.io/google_containers/kube-dnsmasq-amd64:1.4.1
        livenessProbe:
          httpGet:
            path: /healthz-dnsmasq
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - --cache-size=1000
        - --no-resolv
        - --server=127.0.0.1#10053
        - --log-facility=-
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        # see: https://github.com/kubernetes/kubernetes/issues/29055 for details
        resources:
          requests:
            cpu: 150m
            memory: 10Mi
      - name: dnsmasq-metrics
        image: gcr.io/google_containers/dnsmasq-metrics-amd64:1.0.1
        livenessProbe:
          httpGet:
            path: /metrics
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - --v=2
        - --logtostderr
        ports:
        - containerPort: 10054
          name: metrics
          protocol: TCP
        resources:
          requests:
            memory: 10Mi
      - name: healthz
        image: gcr.io/google_containers/exechealthz-amd64:1.2
        resources:
          limits:
            memory: 50Mi
          requests:
            cpu: 10m
            # Note that this container shouldn't really need 50Mi of memory. The
            # limits are set higher than expected pending investigation on #29688.
            # The extra memory was stolen from the kubedns container to keep the
            # net memory requested by the pod constant.
            memory: 50Mi
        args:
        - --cmd=nslookup kubernetes.default.svc.cluster.local 127.0.0.1 >/dev/null
        - --url=/healthz-dnsmasq
        - --cmd=nslookup kubernetes.default.svc.cluster.local 127.0.0.1:10053 >/dev/null
        - --url=/healthz-kubedns
        - --port=8080
        - --quiet
        ports:
        - containerPort: 8080
          protocol: TCP
      dnsPolicy: Default  # Don't use cluster DNS.

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl create -f /root/cluster%2Faddons%2Fdns%2Fskydns-rc.yaml
deployment "kube-dns" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get deployment --namespace=kube-system
NAME       DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kube-dns   1         1         1            0           1m

```

Workload _dns-service_

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo sed "s/\$DNS_SERVER_IP/10.123.240.10/g" /opt/kubernetes/cluster/addons/dns/skydns-svc.yaml.sed | sudo tee /root/cluster%2Faddons%2Fdns%2Fskydns-svc.yaml
# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# TODO - At some point, we need to rename all skydns-*.yaml.* files to kubedns-*.yaml.*

# Warning: This is a file generated from the base underscore template file: skydns-svc.yaml.base

apiVersion: v1
kind: Service
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: "KubeDNS"
spec:
  selector:
    k8s-app: kube-dns
  clusterIP: 10.123.240.10
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl create -f /root/cluster%2Faddons%2Fdns%2Fskydns-svc.yaml
service "kube-dns" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get service --namespace=kube-system
NAME       CLUSTER-IP      EXTERNAL-IP   PORT(S)         AGE
kube-dns   10.123.240.10   <none>        53/UDP,53/TCP   16s

```

### Dashboard

Extract模版

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /99-mirror/https%3A%2F%2Fdl.k8s.io/v1.5.7/kubernetes.tar.gz kubernetes/cluster/addons/dashboard
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes/cluster/addons
addon-manager
BUILD
calico-policy-controller
cluster-loadbalancing
cluster-monitoring
dashboard
dns
dns-horizontal-autoscaler
etcd-empty-dir-cleanup
fluentd-elasticsearch
fluentd-gcp
node-problem-detector
podsecuritypolicies
python-image
README.md
registry

```

Pull镜像

```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/kubernetes-dashboard-amd64:v1.5.0
v1.5.0: Pulling from google_containers/kubernetes-dashboard-amd64
33a12e3650c8: Pull complete 
Digest: sha256:3bccb9256e8b14ae895d40d829ea45992389af3c1767a21eefbd4b3bf723f325
Status: Downloaded newer image for gcr.io/google_containers/kubernetes-dashboard-amd64:v1.5.0

```

Workload of dashboard

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo cp /opt/kubernetes/cluster/addons/dashboard/dashboard-controller.yaml /root/cluster%2Faddons%2Fdashboard%2Fdashboard-controller.yaml
vagrant@vagrant-ubuntu-trusty-64:~$ sudo cp /opt/kubernetes/cluster/addons/dashboard/dashboard-service.yaml /root/cluster%2Faddons%2Fdashboard%2Fdashboard-service.yaml

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl create -f /root/cluster%2Faddons%2Fdashboard%2Fdashboard-controller.yaml -f /root/cluster%2Faddons%2Fdashboard%2Fdashboard-service.yaml
deployment "kubernetes-dashboard" created
service "kubernetes-dashboard" created

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get deployments,services --namespace=kube-system
NAME                          DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/kube-dns               1         1         1            0           20m
deploy/kubernetes-dashboard   1         1         1            1           45s

NAME                       CLUSTER-IP       EXTERNAL-IP   PORT(S)         AGE
svc/kube-dns               10.123.240.10    <none>        53/UDP,53/TCP   14m
svc/kubernetes-dashboard   10.123.251.224   <none>        80/TCP          45s

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl expose services/kubernetes-dashboard --port=30000 --target-port=9090 --type=NodePort --name=kubernetes-dashboard-ex --namespace=kube-system
service "kubernetes-dashboard-ex" exposed
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get services --namespace=kube-system
NAME                      CLUSTER-IP       EXTERNAL-IP   PORT(S)         AGE
kube-dns                  10.123.240.10    <none>        53/UDP,53/TCP   31m
kubernetes-dashboard      10.123.251.224   <none>        80/TCP          17m
kubernetes-dashboard-ex   10.123.246.236   <nodes>       80:30700/TCP    22s
vagrant@vagrant-ubuntu-trusty-64:~$ curl 10.123.246.236
 <!doctype html> <html ng-app="kubernetesDashboard"> <head> <meta charset="utf-8"> <title>Kubernetes Dashboard</title> <link rel="icon" type="image/png" href="assets/images/kubernetes-logo.png"> <meta name="viewport" content="width=device-width"> <link rel="stylesheet" href="static/vendor.36bb79bb.css"> <link rel="stylesheet" href="static/app.968d5cf5.css"> </head> <body> <!--[if lt IE 10]>
      <p class="browsehappy">You are using an <strong>outdated</strong> browser.
      Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your
      experience.</p>
    <![endif]--> <kd-chrome layout="column" layout-fill> </kd-chrome> <script src="static/vendor.633c6c7a.js"></script> <script src="api/appConfig.json"></script> <script src="static/app.29a51112.js"></script> </body> </html> vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ curl 172.17.4.200:30700
 <!doctype html> <html ng-app="kubernetesDashboard"> <head> <meta charset="utf-8"> <title>Kubernetes Dashboard</title> <link rel="icon" type="image/png" href="assets/images/kubernetes-logo.png"> <meta name="viewport" content="width=device-width"> <link rel="stylesheet" href="static/vendor.36bb79bb.css"> <link rel="stylesheet" href="static/app.968d5cf5.css"> </head> <body> <!--[if lt IE 10]>
      <p class="browsehappy">You are using an <strong>outdated</strong> browser.
      Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your
      experience.</p>
    <![endif]--> <kd-chrome layout="column" layout-fill> </kd-chrome> <script src="static/vendor.633c6c7a.js"></script> <script src="api/appConfig.json"></script> <script src="static/app.29a51112.js"></script> </body> </html> 

```

## Appendix

### Kubernetes-manifests

Extract manifests

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-manifests.tar.gz kubernetes/kube-apiserver.yaml
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-manifests.tar.gz kubernetes/kube-controller-manager.yaml
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-manifests.tar.gz kubernetes/kube-proxy.manifest
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-manifests.tar.gz kubernetes/kube-scheduler.yaml
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxf /opt/kubernetes/server/kubernetes-manifests.tar.gz kubernetes/kube-system.json
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls -1 /opt/kubernetes | egrep "yaml|json|manifest"
kube-apiserver.yaml
kube-controller-manager.yaml
kube-proxy.manifest
kube-scheduler.yaml
kube-system.json

```

### Command kubelet

Help of kubelet

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubelet --help
Usage of /opt/kubernetes/server/bin/kubelet:
      --address ip                                              The IP address for the Kubelet to serve on (set to 0.0.0.0 for all interfaces) (default 0.0.0.0)
      --allow-privileged                                        If true, allow containers to request privileged mode. [default=false]
      --alsologtostderr                                         log to standard error as well as files
      --anonymous-auth                                          Enables anonymous requests to the Kubelet server. Requests that are not rejected by another authentication method are treated as anonymous requests. Anonymous requests have a username of system:anonymous, and a group name of system:unauthenticated. (default true)
      --application-metrics-count-limit int                     Max number of application metrics to store (per container) (default 100)
      --authentication-token-webhook                            Use the TokenReview API to determine authentication for bearer tokens.
      --authentication-token-webhook-cache-ttl duration         The duration to cache responses from the webhook token authenticator. (default 2m0s)
      --authorization-mode string                               Authorization mode for Kubelet server. Valid options are AlwaysAllow or Webhook. Webhook mode uses the SubjectAccessReview API to determine authorization. (default "AlwaysAllow")
      --authorization-webhook-cache-authorized-ttl duration     The duration to cache 'authorized' responses from the webhook authorizer. (default 5m0s)
      --authorization-webhook-cache-unauthorized-ttl duration   The duration to cache 'unauthorized' responses from the webhook authorizer. (default 30s)
      --azure-container-registry-config string                  Path to the file container Azure container registry configuration information.
      --boot-id-file string                                     Comma-separated list of files to check for boot-id. Use the first one that exists. (default "/proc/sys/kernel/random/boot_id")
      --cadvisor-port int32                                     The port of the localhost cAdvisor endpoint (default 4194)
      --cert-dir string                                         The directory where the TLS certs are located (by default /var/run/kubernetes). If --tls-cert-file and --tls-private-key-file are provided, this flag will be ignored. (default "/var/run/kubernetes")
      --cgroup-driver string                                    Driver that the kubelet uses to manipulate cgroups on the host.  Possible values: 'cgroupfs', 'systemd' (default "cgroupfs")
      --cgroup-root string                                      Optional root cgroup to use for pods. This is handled by the container runtime on a best effort basis. Default: '', which means use the container runtime default.
      --chaos-chance float                                      If > 0.0, introduce random client errors and latency. Intended for testing. [default=0.0]
      --client-ca-file string                                   If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate.
      --cloud-config string                                     The path to the cloud provider configuration file.  Empty string for no configuration file.
      --cloud-provider string                                   The provider for cloud services. By default, kubelet will attempt to auto-detect the cloud provider. Specify empty string for running with no cloud provider. [default=auto-detect] (default "auto-detect")
      --cluster-dns string                                      IP address for a cluster DNS server.  This value is used for containers' DNS server in case of Pods with "dnsPolicy=ClusterFirst"
      --cluster-domain string                                   Domain for this cluster.  If set, kubelet will configure all containers to search this domain in addition to the host's search domains
      --cni-bin-dir string                                      <Warning: Alpha feature> The full path of the directory in which to search for CNI plugin binaries. Default: /opt/cni/bin
      --cni-conf-dir string                                     <Warning: Alpha feature> The full path of the directory in which to search for CNI config files. Default: /etc/cni/net.d
      --container-hints string                                  location of the container hints file (default "/etc/cadvisor/container_hints.json")
      --container-runtime string                                The container runtime to use. Possible values: 'docker', 'rkt'. Default: 'docker'. (default "docker")
      --container-runtime-endpoint string                       [Experimental] The unix socket endpoint of remote runtime service. The endpoint is used only when CRI integration is enabled (--experimental-cri)
      --containerized                                           Experimental support for running kubelet in a container.  Intended for testing. [default=false]
      --cpu-cfs-quota                                           Enable CPU CFS quota enforcement for containers that specify CPU limits (default true)
      --docker string                                           docker endpoint (default "unix:///var/run/docker.sock")
      --docker-endpoint string                                  Use this for the docker endpoint to communicate with (default "unix:///var/run/docker.sock")
      --docker-env-metadata-whitelist string                    a comma-separated list of environment variable keys that needs to be collected for docker containers
      --docker-exec-handler string                              Handler to use when executing a command in a container. Valid values are 'native' and 'nsenter'. Defaults to 'native'. (default "native")
      --docker-only                                             Only report docker containers in addition to root stats
      --docker-root string                                      DEPRECATED: docker root is read from docker info (this is a fallback, default: /var/lib/docker) (default "/var/lib/docker")
      --enable-controller-attach-detach                         Enables the Attach/Detach controller to manage attachment/detachment of volumes scheduled to this node, and disables kubelet from executing any attach/detach operations (default true)
      --enable-custom-metrics                                   Support for gathering custom metrics.
      --enable-debugging-handlers                               Enables server endpoints for log collection and local running of containers and commands (default true)
      --enable-load-reader                                      Whether to enable cpu load reader
      --enable-server                                           Enable the Kubelet's server (default true)
      --event-burst int32                                       Maximum size of a bursty event records, temporarily allows event records to burst to this number, while still not exceeding event-qps. Only used if --event-qps > 0 (default 10)
      --event-qps int32                                         If > 0, limit event creations per second to this value. If 0, unlimited. (default 5)
      --event-storage-age-limit string                          Max length of time for which to store events (per type). Value is a comma separated list of key values, where the keys are event types (e.g.: creation, oom) or "default" and the value is a duration. Default is applied to all non-specified event types (default "default=0")
      --event-storage-event-limit string                        Max number of events to store (per type). Value is a comma separated list of key values, where the keys are event types (e.g.: creation, oom) or "default" and the value is an integer. Default is applied to all non-specified event types (default "default=0")
      --eviction-hard string                                    A set of eviction thresholds (e.g. memory.available<1Gi) that if met would trigger a pod eviction. (default "memory.available<100Mi")
      --eviction-max-pod-grace-period int32                     Maximum allowed grace period (in seconds) to use when terminating pods in response to a soft eviction threshold being met.  If negative, defer to pod specified value.
      --eviction-minimum-reclaim string                         A set of minimum reclaims (e.g. imagefs.available=2Gi) that describes the minimum amount of resource the kubelet will reclaim when performing a pod eviction if that resource is under pressure.
      --eviction-pressure-transition-period duration            Duration for which the kubelet has to wait before transitioning out of an eviction pressure condition. (default 5m0s)
      --eviction-soft string                                    A set of eviction thresholds (e.g. memory.available<1.5Gi) that if met over a corresponding grace period would trigger a pod eviction.
      --eviction-soft-grace-period string                       A set of eviction grace periods (e.g. memory.available=1m30s) that correspond to how long a soft eviction threshold must hold before triggering a pod eviction.
      --exit-on-lock-contention                                 Whether kubelet should exit upon lock-file contention.
      --experimental-allowed-unsafe-sysctls stringSlice         Comma-separated whitelist of unsafe sysctls or unsafe sysctl patterns (ending in *). Use these at your own risk.
      --experimental-bootstrap-kubeconfig string                <Warning: Experimental feature> Path to a kubeconfig file that will be used to get client certificate for kubelet. If the file specified by --kubeconfig does not exist, the bootstrap kubeconfig is used to request a client certificate from the API server. On success, a kubeconfig file referencing the generated key and obtained certificate is written to the path specified by --kubeconfig. The certificate and key file will be stored in the directory pointed by --cert-dir.
      --experimental-cgroups-per-qos                            Enable creation of QoS cgroup hierarchy, if true top level QoS and pod cgroups are created.
      --experimental-check-node-capabilities-before-mount       [Experimental] if set true, the kubelet will check the underlying node for required componenets (binaries, etc.) before performing the mount
      --experimental-cri                                        [Experimental] Enable the Container Runtime Interface (CRI) integration. If --container-runtime is set to "remote", Kubelet will communicate with the runtime/image CRI server listening on the endpoint specified by --remote-runtime-endpoint/--remote-image-endpoint. If --container-runtime is set to "docker", Kubelet will launch a in-process CRI server on behalf of docker, and communicate over a default endpoint.
      --experimental-fail-swap-on                               Makes the Kubelet fail to start if swap is enabled on the node. This is a temporary opton to maintain legacy behavior, failing due to swap enabled will happen by default in v1.6.
      --experimental-kernel-memcg-notification                  If enabled, the kubelet will integrate with the kernel memcg notification to determine if memory eviction thresholds are crossed rather than polling.
      --experimental-mounter-path string                        [Experimental] Path of mounter binary. Leave empty to use the default mount.
      --experimental-nvidia-gpus int32                          Number of NVIDIA GPU devices on this node. Only 0 (default) and 1 are currently supported.
      --feature-gates string                                    A set of key=value pairs that describe feature gates for alpha/experimental features. Options are:
AllAlpha=true|false (ALPHA - default=false)
AllowExtTrafficLocalEndpoints=true|false (BETA - default=true)
AppArmor=true|false (BETA - default=true)
DynamicKubeletConfig=true|false (ALPHA - default=false)
DynamicVolumeProvisioning=true|false (ALPHA - default=true)
ExperimentalCriticalPodAnnotation=true|false (ALPHA - default=false)
ExperimentalHostUserNamespaceDefaulting=true|false (ALPHA - default=false)
StreamingProxyRedirects=true|false (ALPHA - default=false)
      --file-check-frequency duration                           Duration between checking config files for new data (default 20s)
      --global-housekeeping-interval duration                   Interval between global housekeepings (default 1m0s)
      --google-json-key string                                  The Google Cloud Platform Service Account JSON Key to use for authentication.
      --hairpin-mode string                                     How should the kubelet setup hairpin NAT. This allows endpoints of a Service to loadbalance back to themselves if they should try to access their own Service. Valid values are "promiscuous-bridge", "hairpin-veth" and "none". (default "promiscuous-bridge")
      --healthz-bind-address ip                                 The IP address for the healthz server to serve on, defaulting to 127.0.0.1 (set to 0.0.0.0 for all interfaces) (default 127.0.0.1)
      --healthz-port int32                                      The port of the localhost healthz endpoint (default 10248)
      --host-ipc-sources stringSlice                            Comma-separated list of sources from which the Kubelet allows pods to use the host ipc namespace. [default="*"] (default [*])
      --host-network-sources stringSlice                        Comma-separated list of sources from which the Kubelet allows pods to use of host network. [default="*"] (default [*])
      --host-pid-sources stringSlice                            Comma-separated list of sources from which the Kubelet allows pods to use the host pid namespace. [default="*"] (default [*])
      --hostname-override string                                If non-empty, will use this string as identification instead of the actual hostname.
      --housekeeping-interval duration                          Interval between container housekeepings (default 10s)
      --http-check-frequency duration                           Duration between checking http for new data (default 20s)
      --image-gc-high-threshold int32                           The percent of disk usage after which image garbage collection is always run. Default: 90% (default 90)
      --image-gc-low-threshold int32                            The percent of disk usage before which image garbage collection is never run. Lowest disk usage to garbage collect to. Default: 80% (default 80)
      --image-service-endpoint string                           [Experimental] The unix socket endpoint of remote image service. If not specified, it will be the same with container-runtime-endpoint by default. The endpoint is used only when CRI integration is enabled (--experimental-cri)
      --iptables-drop-bit int32                                 The bit of the fwmark space to mark packets for dropping. Must be within the range [0, 31]. (default 15)
      --iptables-masquerade-bit int32                           The bit of the fwmark space to mark packets for SNAT. Must be within the range [0, 31]. Please match this parameter with corresponding parameter in kube-proxy. (default 14)
      --kube-api-burst int32                                    Burst to use while talking with kubernetes apiserver (default 10)
      --kube-api-content-type string                            Content type of requests sent to apiserver. (default "application/vnd.kubernetes.protobuf")
      --kube-api-qps int32                                      QPS to use while talking with kubernetes apiserver (default 5)
      --kube-reserved mapStringString                           A set of ResourceName=ResourceQuantity (e.g. cpu=200m,memory=150G) pairs that describe resources reserved for kubernetes system components. Currently only cpu and memory are supported. See http://kubernetes.io/docs/user-guide/compute-resources for more detail. [default=none]
      --kubeconfig string                                       Path to a kubeconfig file, specifying how to connect to the API server. --api-servers will be used for the location unless --require-kubeconfig is set. (default "/var/lib/kubelet/kubeconfig")
      --kubelet-cgroups string                                  Optional absolute name of cgroups to create and run the Kubelet in.
      --lock-file string                                        <Warning: Alpha feature> The path to file for kubelet to use as a lock file.
      --log-backtrace-at traceLocation                          when logging hits line file:N, emit a stack trace (default :0)
      --log-cadvisor-usage                                      Whether to log the usage of the cAdvisor container
      --log-dir string                                          If non-empty, write log files in this directory
      --log-flush-frequency duration                            Maximum number of seconds between log flushes (default 5s)
      --logtostderr                                             log to standard error instead of files (default true)
      --low-diskspace-threshold-mb int32                        The absolute free disk space, in MB, to maintain. When disk space falls below this threshold, new pods would be rejected. Default: 256 (default 256)
      --machine-id-file string                                  Comma-separated list of files to check for machine-id. Use the first one that exists. (default "/etc/machine-id,/var/lib/dbus/machine-id")
      --make-iptables-util-chains                               If true, kubelet will ensure iptables utility rules are present on host. (default true)
      --manifest-url string                                     URL for accessing the container manifest
      --manifest-url-header string                              HTTP header to use when accessing the manifest URL, with the key separated from the value with a ':', as in 'key:value'
      --master-service-namespace string                         The namespace from which the kubernetes master services should be injected into pods (default "default")
      --max-open-files int                                      Number of files that can be opened by Kubelet process. [default=1000000] (default 1000000)
      --max-pods int32                                          Number of Pods that can run on this Kubelet. (default 110)
      --minimum-image-ttl-duration duration                     Minimum age for an unused image before it is garbage collected.  Examples: '300ms', '10s' or '2h45m'. Default: '2m' (default 2m0s)
      --network-plugin string                                   <Warning: Alpha feature> The name of the network plugin to be invoked for various events in kubelet/pod lifecycle
      --network-plugin-dir string                               <Warning: Alpha feature> The full path of the directory in which to search for network plugins or CNI config
      --network-plugin-mtu int32                                <Warning: Alpha feature> The MTU to be passed to the network plugin, to override the default. Set to 0 to use the default 1460 MTU.
      --node-ip string                                          IP address of the node. If set, kubelet will use this IP address for the node
      --node-labels mapStringString                             <Warning: Alpha feature> Labels to add when registering the node in the cluster.  Labels must be key=value pairs separated by ','.
      --node-status-update-frequency duration                   Specifies how often kubelet posts node status to master. Note: be cautious when changing the constant, it must work with nodeMonitorGracePeriod in nodecontroller. Default: 10s (default 10s)
      --non-masquerade-cidr string                              Traffic to IPs outside this range will use IP masquerade. (default "10.0.0.0/8")
      --oom-score-adj int32                                     The oom-score-adj value for kubelet process. Values must be within the range [-1000, 1000] (default -999)
      --outofdisk-transition-frequency duration                 Duration for which the kubelet has to wait before transitioning out of out-of-disk node condition status. Default: 5m0s (default 5m0s)
      --pod-cidr string                                         The CIDR to use for pod IP addresses, only used in standalone mode.  In cluster mode, this is obtained from the master.
      --pod-infra-container-image string                        The image whose network/ipc namespaces containers in each pod will use. (default "gcr.io/google_containers/pause-amd64:3.0")
      --pod-manifest-path string                                Path to to the directory containing pod manifest files to run, or the path to a single pod manifest file.
      --pods-per-core int32                                     Number of Pods per core that can run on this Kubelet. The total number of Pods on this Kubelet cannot exceed max-pods, so max-pods will be used if this calculation results in a larger number of Pods allowed on the Kubelet. A value of 0 disables this limit.
      --port int32                                              The port for the Kubelet to serve on. (default 10250)
      --protect-kernel-defaults                                 Default kubelet behaviour for kernel tuning. If set, kubelet errors if any of kernel tunables is different than kubelet defaults.
      --read-only-port int32                                    The read-only port for the Kubelet to serve on with no authentication/authorization (set to 0 to disable) (default 10255)
      --really-crash-for-testing                                If true, when panics occur crash. Intended for testing.
      --register-node                                           Register the node with the apiserver (defaults to true if --api-servers is set) (default true)
      --register-schedulable                                    Register the node as schedulable. Won't have any effect if register-node is false. [default=true] (default true)
      --registry-burst int32                                    Maximum size of a bursty pulls, temporarily allows pulls to burst to this number, while still not exceeding registry-qps.  Only used if --registry-qps > 0 (default 10)
      --registry-qps int32                                      If > 0, limit registry pull QPS to this value.  If 0, unlimited. [default=5.0] (default 5)
      --require-kubeconfig                                      If true the Kubelet will exit if there are configuration errors, and will ignore the value of --api-servers in favor of the server defined in the kubeconfig file.
      --resolv-conf string                                      Resolver configuration file used as the basis for the container DNS resolution configuration. (default "/etc/resolv.conf")
      --rkt-api-endpoint string                                 The endpoint of the rkt API service to communicate with. Only used if --container-runtime='rkt'. (default "localhost:15441")
      --rkt-path string                                         Path of rkt binary. Leave empty to use the first rkt in $PATH.  Only used if --container-runtime='rkt'.
      --root-dir string                                         Directory path for managing kubelet files (volume mounts,etc). (default "/var/lib/kubelet")
      --runonce                                                 If true, exit after spawning pods from local manifests or remote urls. Exclusive with --api-servers, and --enable-server
      --runtime-cgroups string                                  Optional absolute name of cgroups to create and run the runtime in.
      --runtime-request-timeout duration                        Timeout of all runtime requests except long running request - pull, logs, exec and attach. When timeout exceeded, kubelet will cancel the request, throw out an error and retry later. Default: 2m0s (default 2m0s)
      --seccomp-profile-root string                             Directory path for seccomp profiles. (default "/var/lib/kubelet/seccomp")
      --serialize-image-pulls                                   Pull images one at a time. We recommend *not* changing the default value on nodes that run docker daemon with version < 1.9 or an Aufs storage backend. Issue #10959 has more details. [default=true] (default true)
      --stderrthreshold severity                                logs at or above this threshold go to stderr (default 2)
      --storage-driver-buffer-duration duration                 Writes in the storage driver will be buffered for this duration, and committed to the non memory backends as a single transaction (default 1m0s)
      --storage-driver-db string                                database name (default "cadvisor")
      --storage-driver-host string                              database host:port (default "localhost:8086")
      --storage-driver-password string                          database password (default "root")
      --storage-driver-secure                                   use secure connection with database
      --storage-driver-table string                             table name (default "stats")
      --storage-driver-user string                              database username (default "root")
      --streaming-connection-idle-timeout duration              Maximum time a streaming connection can be idle before the connection is automatically closed. 0 indicates no timeout. Example: '5m' (default 4h0m0s)
      --sync-frequency duration                                 Max period between synchronizing running containers and config (default 1m0s)
      --system-cgroups /                                        Optional absolute name of cgroups in which to place all non-kernel processes that are not already inside a cgroup under /. Empty for no container. Rolling back the flag requires a reboot. (Default: "").
      --system-reserved mapStringString                         A set of ResourceName=ResourceQuantity (e.g. cpu=200m,memory=150G) pairs that describe resources reserved for non-kubernetes components. Currently only cpu and memory are supported. See http://kubernetes.io/docs/user-guide/compute-resources for more detail. [default=none]
      --tls-cert-file string                                    File containing x509 Certificate for HTTPS.  (CA cert, if any, concatenated after server cert). If --tls-cert-file and --tls-private-key-file are not provided, a self-signed certificate and key are generated for the public address and saved to the directory passed to --cert-dir.
      --tls-private-key-file string                             File containing x509 private key matching --tls-cert-file.
  -v, --v Level                                                 log level for V logs
      --version version[=true]                                  Print version information and quit
      --vmodule moduleSpec                                      comma-separated list of pattern=N settings for file-filtered logging
      --volume-plugin-dir string                                <Warning: Alpha feature> The full path of the directory in which to search for additional third party volume plugins (default "/usr/libexec/kubernetes/kubelet-plugins/volume/exec/")
      --volume-stats-agg-period duration                        Specifies interval for kubelet to calculate and cache the volume disk usage for all pods and volumes.  To disable volume calculations, set to 0.  Default: '1m' (default 1m0s)
```

### Command kube-apiserver

Help of kube-apiserver

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kube-apiserver --help
Usage of /opt/kubernetes/server/bin/kube-apiserver:
      --admission-control string                                Ordered list of plug-ins to do admission control of resources into cluster. Comma-delimited list of: AlwaysAdmit, AlwaysDeny, AlwaysPullImages, DefaultStorageClass, DenyEscalatingExec, DenyExecOnPrivileged, ImagePolicyWebhook, InitialResources, LimitPodHardAntiAffinityTopology, LimitRanger, NamespaceAutoProvision, NamespaceExists, NamespaceLifecycle, OwnerReferencesPermissionEnforcement, PersistentVolumeLabel, PodNodeSelector, PodSecurityPolicy, ResourceQuota, SecurityContextDeny, ServiceAccount. (default "AlwaysAdmit")
      --admission-control-config-file string                    File with admission control configuration.
      --advertise-address ip                                    The IP address on which to advertise the apiserver to members of the cluster. This address must be reachable by the rest of the cluster. If blank, the --bind-address will be used. If --bind-address is unspecified, the host's default interface will be used.
      --allow-privileged                                        If true, allow privileged containers.
      --alsologtostderr                                         log to standard error as well as files
      --anonymous-auth                                          Enables anonymous requests to the secure port of the API server. Requests that are not rejected by another authentication method are treated as anonymous requests. Anonymous requests have a username of system:anonymous, and a group name of system:unauthenticated. 
      --apiserver-count int                                     The number of apiservers running in the cluster. (default 1)
      --audit-log-maxage int                                    The maximum number of days to retain old audit log files based on the timestamp encoded in their filename.
      --audit-log-maxbackup int                                 The maximum number of old audit log files to retain.
      --audit-log-maxsize int                                   The maximum size in megabytes of the audit log file before it gets rotated. Defaults to 100MB.
      --audit-log-path string                                   If set, all requests coming to the apiserver will be logged to this file.
      --authentication-token-webhook-cache-ttl duration         The duration to cache responses from the webhook token authenticator. Default is 2m. (default 2m0s)
      --authentication-token-webhook-config-file string         File with webhook configuration for token authentication in kubeconfig format. The API server will query the remote service to determine authentication for bearer tokens.
      --authorization-mode string                               Ordered list of plug-ins to do authorization on secure port. Comma-delimited list of: AlwaysAllow,AlwaysDeny,ABAC,Webhook,RBAC. (default "AlwaysAllow")
      --authorization-policy-file string                        File with authorization policy in csv format, used with --authorization-mode=ABAC, on the secure port.
      --authorization-rbac-super-user string                    If specified, a username which avoids RBAC authorization checks and role binding privilege escalation checks, to be used with --authorization-mode=RBAC.
      --authorization-webhook-cache-authorized-ttl duration     The duration to cache 'authorized' responses from the webhook authorizer. Default is 5m. (default 5m0s)
      --authorization-webhook-cache-unauthorized-ttl duration   The duration to cache 'unauthorized' responses from the webhook authorizer. Default is 30s. (default 30s)
      --authorization-webhook-config-file string                File with webhook configuration in kubeconfig format, used with --authorization-mode=Webhook. The API server will query the remote service to determine access on the API server's secure port.
      --basic-auth-file string                                  If set, the file that will be used to admit requests to the secure port of the API server via http basic authentication.
      --bind-address ip                                         The IP address on which to listen for the --secure-port port. The associated interface(s) must be reachable by the rest of the cluster, and by CLI/web clients. If blank, all interfaces will be used (0.0.0.0). (default 0.0.0.0)
      --cert-dir string                                         The directory where the TLS certs are located (by default /var/run/kubernetes). If --tls-cert-file and --tls-private-key-file are provided, this flag will be ignored. (default "/var/run/kubernetes")
      --client-ca-file string                                   If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate.
      --cloud-config string                                     The path to the cloud provider configuration file. Empty string for no configuration file.
      --cloud-provider string                                   The provider for cloud services. Empty string for no provider.
      --contention-profiling                                    Enable contention profiling. Requires --profiling to be set to work.
      --cors-allowed-origins stringSlice                        List of allowed origins for CORS, comma separated.  An allowed origin can be a regular expression to support subdomain matching. If this list is empty CORS will not be enabled.
      --delete-collection-workers int                           Number of workers spawned for DeleteCollection call. These are used to speed up namespace cleanup. (default 1)
      --deserialization-cache-size int                          Number of deserialized json objects to cache in memory.
      --enable-garbage-collector                                Enables the generic garbage collector. MUST be synced with the corresponding flag of the kube-controller-manager. (default true)
      --enable-swagger-ui                                       Enables swagger ui on the apiserver at /swagger-ui
      --etcd-cafile string                                      SSL Certificate Authority file used to secure etcd communication.
      --etcd-certfile string                                    SSL certification file used to secure etcd communication.
      --etcd-keyfile string                                     SSL key file used to secure etcd communication.
      --etcd-prefix string                                      The prefix for all resource paths in etcd. (default "/registry")
      --etcd-quorum-read                                        If true, enable quorum read.
      --etcd-servers stringSlice                                List of etcd servers to connect with (scheme://ip:port), comma separated.
      --etcd-servers-overrides stringSlice                      Per-resource etcd servers overrides, comma separated. The individual override format: group/resource#servers, where servers are http://ip:port, semicolon separated.
      --event-ttl duration                                      Amount of time to retain events. Default is 1h. (default 1h0m0s)
      --experimental-keystone-ca-file string                    If set, the Keystone server's certificate will be verified by one of the authorities in the experimental-keystone-ca-file, otherwise the host's root CA set will be used.
      --experimental-keystone-url string                        If passed, activates the keystone authentication plugin.
      --external-hostname string                                The hostname to use when generating externalized URLs for this master (e.g. Swagger API Docs).
      --feature-gates mapStringBool                             A set of key=value pairs that describe feature gates for alpha/experimental features. Options are:
AllAlpha=true|false (ALPHA - default=false)
AllowExtTrafficLocalEndpoints=true|false (BETA - default=true)
AppArmor=true|false (BETA - default=true)
DynamicKubeletConfig=true|false (ALPHA - default=false)
DynamicVolumeProvisioning=true|false (ALPHA - default=true)
ExperimentalCriticalPodAnnotation=true|false (ALPHA - default=false)
ExperimentalHostUserNamespaceDefaulting=true|false (ALPHA - default=false)
StreamingProxyRedirects=true|false (ALPHA - default=false)
      --insecure-allow-any-token username/group1,group2         If set, your server will be INSECURE.  Any token will be allowed and user information will be parsed from the token as username/group1,group2
      --insecure-bind-address ip                                The IP address on which to serve the --insecure-port (set to 0.0.0.0 for all interfaces). Defaults to localhost. (default 127.0.0.1)
      --insecure-port int                                       The port on which to serve unsecured, unauthenticated access. Default 8080. It is assumed that firewall rules are set up such that this port is not reachable from outside of the cluster and that port 443 on the cluster's public address is proxied to this port. This is performed by nginx in the default setup. (default 8080)
      --ir-data-source string                                   Data source used by InitialResources. Supported options: influxdb, gcm. (default "influxdb")
      --ir-dbname string                                        InfluxDB database name which contains metrics required by InitialResources (default "k8s")
      --ir-hawkular string                                      Hawkular configuration URL
      --ir-influxdb-host string                                 Address of InfluxDB which contains metrics required by InitialResources (default "localhost:8080/api/v1/proxy/namespaces/kube-system/services/monitoring-influxdb:api")
      --ir-namespace-only                                       Whether the estimation should be made only based on data from the same namespace.
      --ir-password string                                      Password used for connecting to InfluxDB (default "root")
      --ir-percentile int                                       Which percentile of samples should InitialResources use when estimating resources. For experiment purposes. (default 90)
      --ir-user string                                          User used for connecting to InfluxDB (default "root")
      --kubelet-certificate-authority string                    Path to a cert file for the certificate authority.
      --kubelet-client-certificate string                       Path to a client cert file for TLS.
      --kubelet-client-key string                               Path to a client key file for TLS.
      --kubelet-https                                           Use https for kubelet connections. (default true)
      --kubelet-preferred-address-types stringSlice             List of the preferred NodeAddressTypes to use for kubelet connections. (default [Hostname,InternalIP,ExternalIP,LegacyHostIP])
      --kubelet-timeout duration                                Timeout for kubelet operations. (default 5s)
      --kubernetes-service-node-port int                        If non-zero, the Kubernetes master service (which apiserver creates/maintains) will be of type NodePort, using this as the value of the port. If zero, the Kubernetes master service will be of type ClusterIP.
      --log-backtrace-at traceLocation                          when logging hits line file:N, emit a stack trace (default :0)
      --log-dir string                                          If non-empty, write log files in this directory
      --log-flush-frequency duration                            Maximum number of seconds between log flushes (default 5s)
      --logtostderr                                             log to standard error instead of files (default true)
      --long-running-request-regexp string                      A regular expression matching long running requests which should be excluded from maximum inflight request handling. (default "(/|^)((watch|proxy)(/|$)|(logs?|portforward|exec|attach)/?$)")
      --master-service-namespace string                         DEPRECATED: the namespace from which the kubernetes master services should be injected into pods. (default "default")
      --max-connection-bytes-per-sec int                        If non-zero, throttle each user connection to this number of bytes/sec. Currently only applies to long-running requests.
      --max-requests-inflight int                               The maximum number of requests in flight at a given time. When the server exceeds this, it rejects requests. Zero for no limit. (default 400)
      --min-request-timeout int                                 An optional field indicating the minimum number of seconds a handler must keep a request open before timing it out. Currently only honored by the watch request handler, which picks a randomized value above this number as the connection timeout, to spread out load. (default 1800)
      --oidc-ca-file string                                     If set, the OpenID server's certificate will be verified by one of the authorities in the oidc-ca-file, otherwise the host's root CA set will be used.
      --oidc-client-id string                                   The client ID for the OpenID Connect client, must be set if oidc-issuer-url is set.
      --oidc-groups-claim string                                If provided, the name of a custom OpenID Connect claim for specifying user groups. The claim value is expected to be a string or array of strings. This flag is experimental, please see the authentication documentation for further details.
      --oidc-issuer-url string                                  The URL of the OpenID issuer, only HTTPS scheme will be accepted. If set, it will be used to verify the OIDC JSON Web Token (JWT).
      --oidc-username-claim string                              The OpenID claim to use as the user name. Note that claims other than the default ('sub') is not guaranteed to be unique and immutable. This flag is experimental, please see the authentication documentation for further details. (default "sub")
      --profiling                                               Enable profiling via web interface host:port/debug/pprof/ (default true)
      --repair-malformed-updates                                If true, server will do its best to fix the update request to pass the validation, e.g., setting empty UID in update request to its existing value. This flag can be turned off after we fix all the clients that send malformed updates. (default true)
      --requestheader-allowed-names stringSlice                 List of client certificate common names to allow to provide usernames in headers specified by --requestheader-username-headers. If empty, any client certificate validated by the authorities in --requestheader-client-ca-file is allowed.
      --requestheader-client-ca-file string                     Root certificate bundle to use to verify client certificates on incoming requests before trusting usernames in headers specified by --requestheader-username-headers
      --requestheader-username-headers stringSlice              List of request headers to inspect for usernames. X-Remote-User is common.
      --runtime-config mapStringString                          A set of key=value pairs that describe runtime configuration that may be passed to apiserver. apis/<groupVersion> key can be used to turn on/off specific api versions. apis/<groupVersion>/<resource> can be used to turn on/off specific resources. api/all and api/legacy are special keys to control all and legacy api versions respectively.
      --secure-port int                                         The port on which to serve HTTPS with authentication and authorization. If 0, don't serve HTTPS at all. (default 6443)
      --service-account-key-file stringArray                    File containing PEM-encoded x509 RSA or ECDSA private or public keys, used to verify ServiceAccount tokens. If unspecified, --tls-private-key-file is used. The specified file can contain multiple keys, and the flag can be specified multiple times with different files.
      --service-account-lookup                                  If true, validate ServiceAccount tokens exist in etcd as part of authentication.
      --service-cluster-ip-range ipNet                          A CIDR notation IP range from which to assign service cluster IPs. This must not overlap with any IP ranges assigned to nodes for pods.
      --service-node-port-range portRange                       A port range to reserve for services with NodePort visibility. Example: '30000-32767'. Inclusive at both ends of the range. (default 30000-32767)
      --ssh-keyfile string                                      If non-empty, use secure SSH proxy to the nodes, using this user keyfile
      --ssh-user string                                         If non-empty, use secure SSH proxy to the nodes, using this user name
      --stderrthreshold severity                                logs at or above this threshold go to stderr (default 2)
      --storage-backend string                                  The storage backend for persistence. Options: 'etcd2' (default), 'etcd3'.
      --storage-media-type string                               The media type to use to store objects in storage. Defaults to application/json. Some resources may only support a specific media type and will ignore this setting. (default "application/json")
      --storage-versions string                                 The per-group version to store resources in. Specified in the format "group1/version1,group2/version2,...". In the case where objects are moved from one group to the other, you may specify the format "group1=group2/v1beta1,group3/v1beta1,...". You only need to pass the groups you wish to change from the defaults. It defaults to a list of preferred versions of all registered groups, which is derived from the KUBE_API_VERSIONS environment variable. (default "apps/v1beta1,authentication.k8s.io/v1beta1,authorization.k8s.io/v1beta1,autoscaling/v1,batch/v1,certificates.k8s.io/v1alpha1,componentconfig/v1alpha1,extensions/v1beta1,imagepolicy.k8s.io/v1alpha1,policy/v1beta1,rbac.authorization.k8s.io/v1alpha1,storage.k8s.io/v1beta1,v1")
      --target-ram-mb int                                       Memory limit for apiserver in MB (used to configure sizes of caches, etc.)
      --tls-ca-file string                                      If set, this certificate authority will used for secure access from Admission Controllers. This must be a valid PEM-encoded CA bundle.
      --tls-cert-file string                                    File containing the default x509 Certificate for HTTPS. (CA cert, if any, concatenated after server cert). If HTTPS serving is enabled, and --tls-cert-file and --tls-private-key-file are not provided, a self-signed certificate and key are generated for the public address and saved to /var/run/kubernetes.
      --tls-private-key-file string                             File containing the default x509 private key matching --tls-cert-file.
      --tls-sni-cert-key namedCertKey                           A pair of x509 certificate and private key file paths, optionally suffixed with a list of domain patterns which are fully qualified domain names, possibly with prefixed wildcard segments. If no domain patterns are provided, the names of the certificate are extracted. Non-wildcard matches trump over wildcard matches, explicit domain patterns trump over extracted names. For multiple key/certificate pairs, use the --tls-sni-cert-key multiple times. Examples: "example.key,example.crt" or "*.foo.com,foo.com:foo.key,foo.crt". (default [])
      --token-auth-file string                                  If set, the file that will be used to secure the secure port of the API server via token authentication.
  -v, --v Level                                                 log level for V logs
      --version version[=true]                                  Print version information and quit
      --vmodule moduleSpec                                      comma-separated list of pattern=N settings for file-filtered logging
      --watch-cache                                             Enable watch caching in the apiserver (default true)
      --watch-cache-sizes stringSlice                           List of watch cache sizes for every resource (pods, nodes, etc.), comma separated. The individual override format: resource#size, where size is a number. It takes effect when watch-cache is enabled.
```

### Command kube-controller-manager

Help of kube-controller-manager

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kube-controller-manager --help
Usage of /opt/kubernetes/server/bin/kube-controller-manager:
      --address ip                                                        The IP address to serve on (set to 0.0.0.0 for all interfaces) (default 0.0.0.0)
      --allocate-node-cidrs                                               Should CIDRs for Pods be allocated and set on the cloud provider.
      --allow-verification-with-non-compliant-keys                        Allow a SignatureVerifier to use keys which are technically non-compliant with RFC6962.
      --alsologtostderr                                                   log to standard error as well as files
      --attach-detach-reconcile-sync-period duration                      The reconciler sync wait time between volume attach detach. This duration must be larger than one second, and increasing this value from the default may allow for volumes to be mismatched with pods. (default 1m0s)
      --cloud-config string                                               The path to the cloud provider configuration file.  Empty string for no configuration file.
      --cloud-provider string                                             The provider for cloud services.  Empty string for no provider.
      --cluster-cidr string                                               CIDR Range for Pods in cluster.
      --cluster-name string                                               The instance prefix for the cluster (default "kubernetes")
      --cluster-signing-cert-file string                                  Filename containing a PEM-encoded X509 CA certificate used to issue cluster-scoped certificates (default "/etc/kubernetes/ca/ca.pem")
      --cluster-signing-key-file string                                   Filename containing a PEM-encoded RSA or ECDSA private key used to sign cluster-scoped certificates (default "/etc/kubernetes/ca/ca.key")
      --concurrent-deployment-syncs int32                                 The number of deployment objects that are allowed to sync concurrently. Larger number = more responsive deployments, but more CPU (and network) load (default 5)
      --concurrent-endpoint-syncs int32                                   The number of endpoint syncing operations that will be done concurrently. Larger number = faster endpoint updating, but more CPU (and network) load (default 5)
      --concurrent-gc-syncs int32                                         The number of garbage collector workers that are allowed to sync concurrently. (default 20)
      --concurrent-namespace-syncs int32                                  The number of namespace objects that are allowed to sync concurrently. Larger number = more responsive namespace termination, but more CPU (and network) load (default 2)
      --concurrent-rc-syncs int32                                         The number of replication controllers that are allowed to sync concurrently. Larger number = more responsive replica management, but more CPU (and network) load (default 5)
      --concurrent-replicaset-syncs int32                                 The number of replica sets that are allowed to sync concurrently. Larger number = more responsive replica management, but more CPU (and network) load (default 5)
      --concurrent-resource-quota-syncs int32                             The number of resource quotas that are allowed to sync concurrently. Larger number = more responsive quota management, but more CPU (and network) load (default 5)
      --concurrent-service-syncs int32                                    The number of services that are allowed to sync concurrently. Larger number = more responsive service management, but more CPU (and network) load (default 1)
      --concurrent-serviceaccount-token-syncs int32                       The number of service account token objects that are allowed to sync concurrently. Larger number = more responsive token generation, but more CPU (and network) load (default 5)
      --configure-cloud-routes                                            Should CIDRs allocated by allocate-node-cidrs be configured on the cloud provider. (default true)
      --controller-start-interval duration                                Interval between starting controller managers.
      --daemonset-lookup-cache-size int32                                 The the size of lookup cache for daemonsets. Larger number = more responsive daemonsets, but more MEM load. (default 1024)
      --deployment-controller-sync-period duration                        Period for syncing the deployments. (default 30s)
      --disable-attach-detach-reconcile-sync                              Disable volume attach detach reconciler sync. Disabling this may cause volumes to be mismatched with pods. Use wisely.
      --enable-dynamic-provisioning                                       Enable dynamic provisioning for environments that support it. (default true)
      --enable-garbage-collector                                          Enables the generic garbage collector. MUST be synced with the corresponding flag of the kube-apiserver. (default true)
      --enable-hostpath-provisioner                                       Enable HostPath PV provisioning when running without a cloud provider. This allows testing and development of provisioning features.  HostPath provisioning is not supported in any way, won't work in a multi-node cluster, and should not be used for anything other than testing or development.
      --feature-gates mapStringBool                                       A set of key=value pairs that describe feature gates for alpha/experimental features. Options are:
AllAlpha=true|false (ALPHA - default=false)
AllowExtTrafficLocalEndpoints=true|false (BETA - default=true)
AppArmor=true|false (BETA - default=true)
DynamicKubeletConfig=true|false (ALPHA - default=false)
DynamicVolumeProvisioning=true|false (ALPHA - default=true)
ExperimentalCriticalPodAnnotation=true|false (ALPHA - default=false)
ExperimentalHostUserNamespaceDefaulting=true|false (ALPHA - default=false)
StreamingProxyRedirects=true|false (ALPHA - default=false)
      --flex-volume-plugin-dir string                                     Full path of the directory in which the flex volume plugin should search for additional third party volume plugins. (default "/usr/libexec/kubernetes/kubelet-plugins/volume/exec/")
      --horizontal-pod-autoscaler-sync-period duration                    The period for syncing the number of pods in horizontal pod autoscaler. (default 30s)
      --insecure-experimental-approve-all-kubelet-csrs-for-group string   The group for which the controller-manager will auto approve all CSRs for kubelet client certificates.
      --kube-api-burst int32                                              Burst to use while talking with kubernetes apiserver (default 30)
      --kube-api-content-type string                                      Content type of requests sent to apiserver. (default "application/vnd.kubernetes.protobuf")
      --kube-api-qps float32                                              QPS to use while talking with kubernetes apiserver (default 20)
      --kubeconfig string                                                 Path to kubeconfig file with authorization and master location information.
      --large-cluster-size-threshold int32                                Number of nodes from which NodeController treats the cluster as large for the eviction logic purposes. --secondary-node-eviction-rate is implicitly overridden to 0 for clusters this size or smaller. (default 50)
      --leader-elect                                                      Start a leader election client and gain leadership before executing the main loop. Enable this when running replicated components for high availability. (default true)
      --leader-elect-lease-duration duration                              The duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. This is only applicable if leader election is enabled. (default 15s)
      --leader-elect-renew-deadline duration                              The interval between attempts by the acting master to renew a leadership slot before it stops leading. This must be less than or equal to the lease duration. This is only applicable if leader election is enabled. (default 10s)
      --leader-elect-retry-period duration                                The duration the clients should wait between attempting acquisition and renewal of a leadership. This is only applicable if leader election is enabled. (default 2s)
      --log-backtrace-at traceLocation                                    when logging hits line file:N, emit a stack trace (default :0)
      --log-dir string                                                    If non-empty, write log files in this directory
      --log-flush-frequency duration                                      Maximum number of seconds between log flushes (default 5s)
      --loglevel int                                                      Log level (0 = DEBUG, 5 = FATAL) (default 1)
      --logtostderr                                                       log to standard error instead of files (default true)
      --master string                                                     The address of the Kubernetes API server (overrides any value in kubeconfig)
      --min-resync-period duration                                        The resync period in reflectors will be random between MinResyncPeriod and 2*MinResyncPeriod (default 12h0m0s)
      --namespace-sync-period duration                                    The period for syncing namespace life-cycle updates (default 5m0s)
      --node-cidr-mask-size int32                                         Mask size for node cidr in cluster. (default 24)
      --node-eviction-rate float32                                        Number of nodes per second on which pods are deleted in case of node failure when a zone is healthy (see --unhealthy-zone-threshold for definition of healthy/unhealthy). Zone refers to entire cluster in non-multizone clusters. (default 0.1)
      --node-monitor-grace-period duration                                Amount of time which we allow running Node to be unresponsive before marking it unhealthy. Must be N times more than kubelet's nodeStatusUpdateFrequency, where N means number of retries allowed for kubelet to post node status. (default 40s)
      --node-monitor-period duration                                      The period for syncing NodeStatus in NodeController. (default 5s)
      --node-startup-grace-period duration                                Amount of time which we allow starting Node to be unresponsive before marking it unhealthy. (default 1m0s)
      --pod-eviction-timeout duration                                     The grace period for deleting pods on failed nodes. (default 5m0s)
      --port int32                                                        The port that the controller-manager's http service runs on (default 10252)
      --profiling                                                         Enable profiling via web interface host:port/debug/pprof/ (default true)
      --pv-recycler-increment-timeout-nfs int32                           the increment of time added per Gi to ActiveDeadlineSeconds for an NFS scrubber pod (default 30)
      --pv-recycler-minimum-timeout-hostpath int32                        The minimum ActiveDeadlineSeconds to use for a HostPath Recycler pod.  This is for development and testing only and will not work in a multi-node cluster. (default 60)
      --pv-recycler-minimum-timeout-nfs int32                             The minimum ActiveDeadlineSeconds to use for an NFS Recycler pod (default 300)
      --pv-recycler-pod-template-filepath-hostpath string                 The file path to a pod definition used as a template for HostPath persistent volume recycling. This is for development and testing only and will not work in a multi-node cluster.
      --pv-recycler-pod-template-filepath-nfs string                      The file path to a pod definition used as a template for NFS persistent volume recycling
      --pv-recycler-timeout-increment-hostpath int32                      the increment of time added per Gi to ActiveDeadlineSeconds for a HostPath scrubber pod.  This is for development and testing only and will not work in a multi-node cluster. (default 30)
      --pvclaimbinder-sync-period duration                                The period for syncing persistent volumes and persistent volume claims (default 15s)
      --replicaset-lookup-cache-size int32                                The the size of lookup cache for replicatsets. Larger number = more responsive replica management, but more MEM load. (default 4096)
      --replication-controller-lookup-cache-size int32                    The the size of lookup cache for replication controllers. Larger number = more responsive replica management, but more MEM load. (default 4096)
      --resource-quota-sync-period duration                               The period for syncing quota usage status in the system (default 5m0s)
      --root-ca-file string                                               If set, this root certificate authority will be included in service account's token secret. This must be a valid PEM-encoded CA bundle.
      --route-reconciliation-period duration                              The period for reconciling routes created for Nodes by cloud provider. (default 10s)
      --secondary-node-eviction-rate float32                              Number of nodes per second on which pods are deleted in case of node failure when a zone is unhealthy (see --unhealthy-zone-threshold for definition of healthy/unhealthy). Zone refers to entire cluster in non-multizone clusters. This value is implicitly overridden to 0 if the cluster size is smaller than --large-cluster-size-threshold. (default 0.01)
      --service-account-private-key-file string                           Filename containing a PEM-encoded private RSA or ECDSA key used to sign service account tokens.
      --service-cluster-ip-range string                                   CIDR Range for Services in cluster.
      --service-sync-period duration                                      The period for syncing services with their external load balancers (default 5m0s)
      --stderrthreshold severity                                          logs at or above this threshold go to stderr (default 2)
      --terminated-pod-gc-threshold int32                                 Number of terminated pods that can exist before the terminated pod garbage collector starts deleting terminated pods. If <= 0, the terminated pod garbage collector is disabled. (default 12500)
      --unhealthy-zone-threshold float32                                  Fraction of Nodes in a zone which needs to be not Ready (minimum 3) for zone to be treated as unhealthy.  (default 0.55)
      --use-service-account-credentials                                   If true, use individual service account credentials for each controller.
  -v, --v Level                                                           log level for V logs
      --version version[=true]                                            Print version information and quit
      --vmodule moduleSpec                                                comma-separated list of pattern=N settings for file-filtered logging
```

### Command kube-scheduler

Help of kube-scheduler

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kube-scheduler --help
Usage of /opt/kubernetes/server/bin/kube-scheduler:
      --address string                           The IP address to serve on (set to 0.0.0.0 for all interfaces) (default "0.0.0.0")
      --algorithm-provider string                The scheduling algorithm provider to use, one of: ClusterAutoscalerProvider | DefaultProvider (default "DefaultProvider")
      --alsologtostderr                          log to standard error as well as files
      --failure-domains string                   Indicate the "all topologies" set for an empty topologyKey when it's used for PreferredDuringScheduling pod anti-affinity. (default "kubernetes.io/hostname,failure-domain.beta.kubernetes.io/zone,failure-domain.beta.kubernetes.io/region")
      --feature-gates mapStringBool              A set of key=value pairs that describe feature gates for alpha/experimental features. Options are:
AllAlpha=true|false (ALPHA - default=false)
AllowExtTrafficLocalEndpoints=true|false (BETA - default=true)
AppArmor=true|false (BETA - default=true)
DynamicKubeletConfig=true|false (ALPHA - default=false)
DynamicVolumeProvisioning=true|false (ALPHA - default=true)
ExperimentalCriticalPodAnnotation=true|false (ALPHA - default=false)
ExperimentalHostUserNamespaceDefaulting=true|false (ALPHA - default=false)
StreamingProxyRedirects=true|false (ALPHA - default=false)
      --hard-pod-affinity-symmetric-weight int   RequiredDuringScheduling affinity is not symmetric, but there is an implicit PreferredDuringScheduling affinity rule corresponding to every RequiredDuringScheduling affinity rule. --hard-pod-affinity-symmetric-weight represents the weight of implicit PreferredDuringScheduling affinity rule. (default 1)
      --kube-api-burst int32                     Burst to use while talking with kubernetes apiserver (default 100)
      --kube-api-content-type string             Content type of requests sent to apiserver. (default "application/vnd.kubernetes.protobuf")
      --kube-api-qps float32                     QPS to use while talking with kubernetes apiserver (default 50)
      --kubeconfig string                        Path to kubeconfig file with authorization and master location information.
      --leader-elect                             Start a leader election client and gain leadership before executing the main loop. Enable this when running replicated components for high availability. (default true)
      --leader-elect-lease-duration duration     The duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. This is only applicable if leader election is enabled. (default 15s)
      --leader-elect-renew-deadline duration     The interval between attempts by the acting master to renew a leadership slot before it stops leading. This must be less than or equal to the lease duration. This is only applicable if leader election is enabled. (default 10s)
      --leader-elect-retry-period duration       The duration the clients should wait between attempting acquisition and renewal of a leadership. This is only applicable if leader election is enabled. (default 2s)
      --log-backtrace-at traceLocation           when logging hits line file:N, emit a stack trace (default :0)
      --log-dir string                           If non-empty, write log files in this directory
      --log-flush-frequency duration             Maximum number of seconds between log flushes (default 5s)
      --logtostderr                              log to standard error instead of files (default true)
      --master string                            The address of the Kubernetes API server (overrides any value in kubeconfig)
      --policy-config-file string                File with scheduler policy configuration
      --port int32                               The port that the scheduler's http service runs on (default 10251)
      --profiling                                Enable profiling via web interface host:port/debug/pprof/ (default true)
      --scheduler-name string                    Name of the scheduler, used to select which pods will be processed by this scheduler, based on pod's annotation with key 'scheduler.alpha.kubernetes.io/name' (default "default-scheduler")
      --stderrthreshold severity                 logs at or above this threshold go to stderr (default 2)
  -v, --v Level                                  log level for V logs
      --version version[=true]                   Print version information and quit
      --vmodule moduleSpec                       comma-separated list of pattern=N settings for file-filtered logging
```

### Command kube-proxy

Help of kube-proxy

```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kube-proxy --help
Usage of /opt/kubernetes/server/bin/kube-proxy:
      --alsologtostderr                              log to standard error as well as files
      --bind-address ip                              The IP address for the proxy server to serve on (set to 0.0.0.0 for all interfaces) (default 0.0.0.0)
      --cleanup-iptables                             If true cleanup iptables rules and exit.
      --cluster-cidr string                          The CIDR range of pods in the cluster. It is used to bridge traffic coming from outside of the cluster. If not provided, no off-cluster bridging will be performed.
      --config-sync-period duration                  How often configuration from the apiserver is refreshed.  Must be greater than 0. (default 15m0s)
      --conntrack-max-per-core int32                 Maximum number of NAT connections to track per CPU core (0 to leave the limit as-is and ignore conntrack-min). (default 32768)
      --conntrack-min int32                          Minimum number of conntrack entries to allocate, regardless of conntrack-max-per-core (set conntrack-max-per-core=0 to leave the limit as-is). (default 131072)
      --conntrack-tcp-timeout-close-wait duration    NAT timeout for TCP connections in the CLOSE_WAIT state (default 1h0m0s)
      --conntrack-tcp-timeout-established duration   Idle timeout for established TCP connections (0 to leave as-is) (default 24h0m0s)
      --feature-gates mapStringBool                  A set of key=value pairs that describe feature gates for alpha/experimental features. Options are:
AllAlpha=true|false (ALPHA - default=false)
AllowExtTrafficLocalEndpoints=true|false (BETA - default=true)
AppArmor=true|false (BETA - default=true)
DynamicKubeletConfig=true|false (ALPHA - default=false)
DynamicVolumeProvisioning=true|false (ALPHA - default=true)
ExperimentalCriticalPodAnnotation=true|false (ALPHA - default=false)
ExperimentalHostUserNamespaceDefaulting=true|false (ALPHA - default=false)
StreamingProxyRedirects=true|false (ALPHA - default=false)
      --healthz-bind-address ip                      The IP address for the health check server to serve on, defaulting to 127.0.0.1 (set to 0.0.0.0 for all interfaces) (default 127.0.0.1)
      --healthz-port int32                           The port to bind the health check server. Use 0 to disable. (default 10249)
      --hostname-override string                     If non-empty, will use this string as identification instead of the actual hostname.
      --iptables-masquerade-bit int32                If using the pure iptables proxy, the bit of the fwmark space to mark packets requiring SNAT with.  Must be within the range [0, 31]. (default 14)
      --iptables-min-sync-period duration            The minimum interval of how often the iptables rules can be refreshed as endpoints and services change (e.g. '5s', '1m', '2h22m').
      --iptables-sync-period duration                The maximum interval of how often iptables rules are refreshed (e.g. '5s', '1m', '2h22m').  Must be greater than 0. (default 30s)
      --kube-api-burst int32                         Burst to use while talking with kubernetes apiserver (default 10)
      --kube-api-content-type string                 Content type of requests sent to apiserver. (default "application/vnd.kubernetes.protobuf")
      --kube-api-qps float32                         QPS to use while talking with kubernetes apiserver (default 5)
      --kubeconfig string                            Path to kubeconfig file with authorization information (the master location is set by the master flag).
      --log-backtrace-at traceLocation               when logging hits line file:N, emit a stack trace (default :0)
      --log-dir string                               If non-empty, write log files in this directory
      --log-flush-frequency duration                 Maximum number of seconds between log flushes (default 5s)
      --logtostderr                                  log to standard error instead of files (default true)
      --masquerade-all                               If using the pure iptables proxy, SNAT everything
      --master string                                The address of the Kubernetes API server (overrides any value in kubeconfig)
      --oom-score-adj int32                          The oom-score-adj value for kube-proxy process. Values must be within the range [-1000, 1000] (default -999)
      --proxy-mode ProxyMode                         Which proxy mode to use: 'userspace' (older) or 'iptables' (faster). If blank, look at the Node object on the Kubernetes API and respect the 'net.experimental.kubernetes.io/proxy-mode' annotation if provided.  Otherwise use the best-available proxy (currently iptables).  If the iptables proxy is selected, regardless of how, but the system's kernel or iptables versions are insufficient, this always falls back to the userspace proxy.
      --proxy-port-range port-range                  Range of host ports (beginPort-endPort, inclusive) that may be consumed in order to proxy service traffic. If unspecified (0-0) then ports will be randomly chosen.
      --stderrthreshold severity                     logs at or above this threshold go to stderr (default 2)
      --udp-timeout duration                         How long an idle UDP connection will be kept open (e.g. '250ms', '2s').  Must be greater than 0. Only applicable for proxy-mode=userspace (default 250ms)
  -v, --v Level                                      log level for V logs
      --version version[=true]                       Print version information and quit
      --vmodule moduleSpec                           comma-separated list of pattern=N settings for file-filtered logging
```

