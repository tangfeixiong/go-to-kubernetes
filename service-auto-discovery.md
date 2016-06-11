# Deep dive into Kubernetes sevice auto discovery mechanism
__Kubernetes的服务自动发现机制__

## Openshift/Kubernetes development environment
__使用Openshift的vagrant开发环境__

It is openshift vagrant development environment [fork](https://github.com/stackdocker/origin)

IaaS: Mac + VirtualBox + Vagrant

* PODs
__以下为开发环境中运行的Kubernetes应用__

部署架构参考coreos的tectonic

As show in below, Kuberentes are all run in `kube-system` namespace except of core component `kubelet`, as well as other important external serives like `etcd2`, `docker engine` or `rkt`, and container networking for example `flannel`.

Openshift could be run as Kubernetes applications in sepearte `openshift-origin` namespace, that is different with RedHat architecture: Openshift plus Kuberentes are worked as containerized PaaS.

Helm is newly joined project under kubernetes community, it is consisted with tiller server and helm client, now the tiller server is run in own namespace

`default` namespace include experimental applications: nc-http is a hello world http server based on netcat, [forked from appropriate/docker-nc](https://github.com/tangfeixiong/docker-nc), redis and sentinel are HA cluster, [forked from deis/redis-cluster](https://github.com/stackdocker/redis-cluster) 

    [vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig get pods --all-namespaces
    NAMESPACE          NAME                                  READY     STATUS    RESTARTS   AGE
    default            nc-http-1578405182-27w56              1/1       Running   0          9d
    default            redis-40abi                           1/1       Running   0          1m
    default            redis-40abi                           1/1       Running   0          3m
    default            redis-8bt3q                           1/1       Running   0          3m
    default            redis-sentinel-otr7k                  1/1       Running   0          3m
    default            redis-sentinel-skiwv                  1/1       Running   0          3m
    default            redis-sentinel-y2x62                  1/1       Running   0          2m
    helm               openshift-webtty-29ge2                1/1       Running   1          9d
    helm               tiller-rc-a2hri                       1/1       Running   1          9d
    kube-system        kube-apiserver-172.17.4.50            1/1       Running   1          1d
    kube-system        kube-controller-manager-172.17.4.50   1/1       Running   1          9d
    kube-system        kube-dns-v9-0suvp                     4/4       Running   0          4m
    kube-system        kube-proxy-172.17.4.50                1/1       Running   1          9d
    kube-system        kube-scheduler-172.17.4.50            1/1       Running   2          9d
    openshift-origin   etcd-ecy6p                            1/1       Running   1          1d
    openshift-origin   openshift-4jvgb                       1/1       Running   0          9d

* Services
__公开的服务__

The following output shows only `redis` is not exposed as service. Kuberentes service exposing is a function that port forward application listening address for example tcp or udp address and port to cluster service networking: Virtual IP that could be load balanced by all cluster machines. 

服务显示的可端口转发到docker集群网络的VIP（虚地址）

    [vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig get services --all-namespaces
    NAMESPACE          NAME               CLUSTER-IP   EXTERNAL-IP   PORT(S)             AGE
    default            kubernetes         10.3.0.1     <none>        443/TCP             30d
    default            nc-http            10.3.0.15    <none>        80/TCP              30d
    default            redis-sentinel     10.3.0.206   <none>        26379/TCP           7d
    helm               openshift-webtty   10.3.0.151                 9123/TCP            10d
    kube-system        kube-dns           10.3.0.10    <none>        53/UDP,53/TCP       30d
    openshift-origin   etcd               10.3.0.16    <none>        4001/TCP,7001/TCP   30d
    openshift-origin   openshift          10.3.0.179                 8443/TCP            30d 

In advanced, Kubernetes service could also configure to serve out of cluster networking, that is called NodePort service. For example, the following filter show `openshift-webtty` and `openshift` has `NodePort`

其中openshift-webtty和openshift的服务是NodePort类型，即允许从集群主机DNAT到VIP

    [vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig get services --all-namespaces --output='jsonpath={.items[?(@.spec.type=="LoadBalancer")].spec}';echo
    map[selector:map[name:openshift-webtty] clusterIP:10.3.0.151 type:LoadBalancer sessionAffinity:None ports:[map[targetPort:9123 nodePort:32121 name:openshift-webtty protocol:TCP port:9123]]] map[selector:map[name:openshift] clusterIP:10.3.0.179 type:LoadBalancer sessionAffinity:None ports:[map[targetPort:8443 nodePort:30448 name:openshift protocol:TCP port:8443]]]

Wanting to access Kuberentes NodePort services (for example `openshift-webtty`) under machine networking (LAN, VirtualBox host-only), simply do:

因此，可以在集群主机上访问此类服务

    fanhonglingdeMacBook-Pro:single-node fanhongling$ curl -ikL http://172.17.4.50:32121
    HTTP/1.1 200 OK
    X-Powered-By: Express
    Accept-Ranges: bytes
    ETag: "630-1464592794000"
    Date: Fri, 10 Jun 2016 19:59:46 GMT
    Cache-Control: public, max-age=0
    Last-Modified: Mon, 30 May 2016 07:19:54 GMT
    Content-Type: text/html; charset=UTF-8
    Content-Length: 630
    Connection: keep-alive
    
    <!doctype html>
    <html lang="en">
    
    <head>
        <meta charset="UTF-8">
        <title>Wetty - The WebTTY Terminal Emulator</title>
        <script src="/openshift/hterm_all.js"></script>

...

In fact, the previous service types are realy not "NodePort" but "LoadBalancer". Because the `NodePort` is based on `ClusterIP`, e.g. internal load balanced service, as well as `LoadBalancer` is based on `NodePort`. The `LoadBalancer` type should go to port-forarding Node Port onto IaaS application network controller, such as Public Cloud Load Balancer.

In this development platform the IaaS is VirtualBox and Management platform is Vagrant. the `Vagrantfile` shows:

如果Kubernetes配置了相应的IaaS LB组件，则可以配置为NodePort之上的LoadBalance类型，上面显示事实不是NodePort，但只作为NodePort在使用

      if vagrant_openshift_config['private_network_ip']
        config.vm.network "private_network", ip: vagrant_openshift_config['private_network_ip']
      else
        config.vm.network "forwarded_port", guest: 80, host: 1080
        config.vm.network "forwarded_port", guest: 443, host: 1443
        config.vm.network "forwarded_port", guest: 8080, host: 8080
        config.vm.network "forwarded_port", guest: 8443, host: 8443
      end

And box configuration file `.vagrant-openshift.json` is customized:

    "private_network_ip":  "172.17.4.50",

Thus the `LoadBalancer` type won't worked well, but only has availability of "NodePort", for example access Kuberentes api:

    fanhonglingdeMacBook-Pro:single-node fanhongling$ ~/Downloads/workspace/bin/kubectl --certificate-authority=/Users/fanhongling/Downloads/github.com/openshift/origin/etc/kubernetes/cacerts/ca.crt --client-certificate=/Users/fanhongling/Downloads/github.com/openshift/origin/etc/kubernetes/cacerts/kubecfg.crt --client-key=/Users/fanhongling/Downloads/github.com/openshift/origin/etc/kubernetes/cacerts/kubecfg.key --server=https://172.17.4.50:443 get pods
    NAME                       READY     STATUS    RESTARTS   AGE
    nc-http-1578405182-27w56   1/1       Running   0          9d
    redis-40abi                1/1       Running   0          1h
    redis-8bt3q                1/1       Running   0          1h
    redis-sentinel-otr7k       1/1       Running   0          1h
    redis-sentinel-skiwv       1/1       Running   0          1h
    redis-sentinel-y2x62       1/1       Running   0          1h

Or:

    fanhonglingdeMacBook-Pro:single-node fanhongling$ KUBECONFIG=/Users/fanhongling/Downloads/github.com/openshift/origin/kubeconfig ~/Downloads/workspace/bin/oc get namespaces
    NAME               STATUS    AGE
    default            Active    30d
    gogogo             Active    4d
    helm               Active    29d
    kube-system        Active    30d
    openshift          Active    30d
    openshift-infra    Active    30d
    openshift-origin   Active    30d
    tangfeixiong       Active    30d

## In namespacing service auto discovery mechanism

__namespace内的服务自动发现__

Docker compose与之在原理上类似

### Study case: redis cluster

所示redis集群为例，从Docker build清单的bash文件中并未发现如何export服务地址和端口的环境变量（因为部署之前根本无从确认）

The redis cluster solution is from [deis/redis-cluster](https://github.com/deis/redis-cluster)

In Dockerfile context, [the entrypoint script](https://github.com/stackdocker/redis-cluster/blob/upstream/rootfs/run.sh) has `launchsentinel` function like that:

    master=$(redis-cli -h ${REDIS_SENTINEL_SERVICE_HOST} -p ${REDIS_SENTINEL_SERVICE_PORT}

It may be wondered where to export 2 environment variables

Actually, they are provided by Kuberentes service auto discovery controller

From this redis cluster provisioning procedure

通过部署资源步骤

* The first step is creating standalone Redis and Sentinel

    [vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig create -f /data/src/github.com/deis/redis-cluster/manifests/redis-master.yaml 
    pod "redis-master" created

This time the service is not arrived out of Docker networking

* The second step is exposing Sentinel service

    [vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig create -f /data/src/github.com/deis/redis-cluster/manifests/redis-sentinel-service.yaml 
    service "redis-sentinel" created

According service manifest, Kubernetes give all containers a service discovery method through service endpoint environment variable generation, for example:

    [vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig get pods
    NAME                       READY     STATUS    RESTARTS   AGE
    nc-http-1578405182-27w56   1/1       Running   0          9d
    redis-40abi                1/1       Running   0          2h
    redis-8bt3q                1/1       Running   0          2h
    redis-sentinel-otr7k       1/1       Running   0          2h
    redis-sentinel-skiwv       1/1       Running   0          2h
    redis-sentinel-y2x62       1/1       Running   0          2h
    
The nc-http application has run 9 days, while entering its container and do investiagtion:

然后，进入当前namespace内与redis集群无关的nc-http容器，export环境变量，除了Kuberentes的POD地址，服务地址，自身的地址，还注入了Redis的地址

    [vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig exec -t -i nc-http-1578405182-27w56 /bin/ash
    / # export
    export HOME='/root'
    export HOSTNAME='nc-http-1578405182-27w56'
    export KUBERNETES_PORT='tcp://10.3.0.1:443'
    export KUBERNETES_PORT_443_TCP='tcp://10.3.0.1:443'
    export KUBERNETES_PORT_443_TCP_ADDR='10.3.0.1'
    export KUBERNETES_PORT_443_TCP_PORT='443'
    export KUBERNETES_PORT_443_TCP_PROTO='tcp'
    export KUBERNETES_SERVICE_HOST='10.3.0.1'
    export KUBERNETES_SERVICE_PORT='443'
    export KUBERNETES_SERVICE_PORT_HTTPS='443'
    export NC_HTTP_PORT='tcp://10.3.0.15:80'
    export NC_HTTP_PORT_80_TCP='tcp://10.3.0.15:80'
    export NC_HTTP_PORT_80_TCP_ADDR='10.3.0.15'
    export NC_HTTP_PORT_80_TCP_PORT='80'
    export NC_HTTP_PORT_80_TCP_PROTO='tcp'
    export NC_HTTP_SERVICE_HOST='10.3.0.15'
    export NC_HTTP_SERVICE_PORT='80'
    export PATH='/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin'
    export PWD='/'
    export REDIS_SENTINEL_PORT='tcp://10.3.0.206:26379'
    export REDIS_SENTINEL_PORT_26379_TCP='tcp://10.3.0.206:26379'
    export REDIS_SENTINEL_PORT_26379_TCP_ADDR='10.3.0.206'
    export REDIS_SENTINEL_PORT_26379_TCP_PORT='26379'
    export REDIS_SENTINEL_PORT_26379_TCP_PROTO='tcp'
    export REDIS_SENTINEL_SERVICE_HOST='10.3.0.206'
    export REDIS_SENTINEL_SERVICE_PORT='26379'
    export SHLVL='1'

Upon front output, the `REDIS_SENTINEL_SERVICE_HOST` and `REDIS_SENTINEL_SERVICE_PORT` are auto generated and injected into containers in this namespace

Try to access Redis Sentinel

以下示例应用程序访问Sentinel和发现Redis服务端口，因为Redis集群是主从架构，因此提供的是Docker网络地址。实际，也可以为Redis配置ClusterIP服务，但Sentinel并无机制去知道外部赋予的VIP，因此可以不用Sentinel，但Client App需要自行发现主Redis

    / # echo -e "INFO\r\nQUIT\r\n" | nc ${REDIS_SENTINEL_SERVICE_HOST} ${REDIS_SENTINEL_SERVICE_PORT}
    $597
    # Server
    redis_version:2.8.23
    redis_git_sha1:00000000
    redis_git_dirty:0
    redis_build_id:ec084e4e7f0409d
    redis_mode:sentinel
    os:Linux 4.2.3-300.fc23.x86_64 x86_64
    arch_bits:64
    multiplexing_api:epoll
    gcc_version:4.9.2
    process_id:12
    run_id:c4616846f33804846dbf0f7c72d031387bfdf5ff
    tcp_port:26379
    uptime_in_seconds:10003
    uptime_in_days:0
    hz:10
    lru_clock:5977634
    config_file:/data/sentinel.conf

    # Sentinel
    sentinel_masters:1
    sentinel_tilt:0
    sentinel_running_scripts:0
    sentinel_scripts_queue_length:0
    master0:name=mymaster,status=ok,address=172.17.0.5:6379,slaves=2,sentinels=4

    +OK

From last line of output, it could find the Redis master serving address, and play as consuming client:

    / # echo -e "INFO\r\nQUIT\r\n" | nc 172.17.0.5 6379
    $2108
    # Server
    redis_version:2.8.23
    redis_git_sha1:00000000
    redis_git_dirty:0
    redis_build_id:ec084e4e7f0409d
    redis_mode:standalone
    os:Linux 4.2.3-300.fc23.x86_64 x86_64
    arch_bits:64
    multiplexing_api:epoll
    gcc_version:4.9.2
    process_id:14
    run_id:5d2edd47a2091f4b331531bfe7a2fe419177a3bb
    tcp_port:6379
    uptime_in_seconds:10199
    uptime_in_days:0
    hz:10
    lru_clock:5977823
    config_file:/redis-slave/redis.conf
    
    # Clients
    connected_clients:7
    client_longest_output_list:0
    client_biggest_input_buf:7
    blocked_clients:0
    
    # Memory
    used_memory:2053072
    used_memory_human:1.96M
    used_memory_rss:4759552
    used_memory_peak:2105488
    used_memory_peak_human:2.01M
    used_memory_lua:36864
    mem_fragmentation_ratio:2.32
    mem_allocator:jemalloc-3.6.0
    
    # Persistence
    loading:0
    rdb_changes_since_last_save:0
    rdb_bgsave_in_progress:0
    rdb_last_save_time:1465585922
    rdb_last_bgsave_status:ok
    rdb_last_bgsave_time_sec:1
    rdb_current_bgsave_time_sec:-1
    aof_enabled:1
    aof_rewrite_in_progress:0
    aof_rewrite_scheduled:0
    aof_last_rewrite_time_sec:1
    aof_current_rewrite_time_sec:-1
    aof_last_bgrewrite_status:ok
    aof_last_write_status:ok
    aof_current_size:52
    aof_base_size:0
    aof_pending_rewrite:0
    aof_buffer_length:0
    aof_rewrite_buffer_length:0
    aof_pending_bio_fsync:0
    aof_delayed_fsync:0
    
    # Stats
    total_connections_received:23
    total_commands_processed:58582
    instantaneous_ops_per_sec:6
    total_net_input_bytes:2912634
    total_net_output_bytes:17201245
    instantaneous_input_kbps:0.25
    instantaneous_output_kbps:1.96
    rejected_connections:0
    sync_full:1
    sync_partial_ok:0
    sync_partial_err:0
    expired_keys:0
    evicted_keys:0
    keyspace_hits:0
    keyspace_misses:0
    pubsub_channels:1
    pubsub_patterns:0
    latest_fork_usec:866
    
    # Replication
    role:master
    connected_slaves:1
    slave0:ip=172.17.0.2,port=6379,state=online,offset=1959177,lag=1
    master_repl_offset:1959448
    repl_backlog_active:1
    repl_backlog_size:1048576
    repl_backlog_first_byte_offset:910873
    repl_backlog_histlen:1048576
    
    # CPU
    used_cpu_sys:10.29
    used_cpu_user:15.42
    used_cpu_sys_children:0.00
    used_cpu_user_children:0.00
    
    # Keyspace
    
    +OK

From Redis info, the `Replication` section shows the master has one slave: `slave0:ip=172.17.0.2,port=6379,state=online,offset=1959177,lag=1`

* Next steps

由于Redis Master POD只是提供服务自动发现的起点，并不配置到Kubernetes的Replication Controller，因此，如上，已经被删除

This just answer where the previous scripts see environments variables from

So in continuing RC creation, new Redis and Sentinel containers could exactily join the cluster when starting the entrypoint script (run.sh)

* Notice

Every container has environment variables of Kuberentes service

## In clustering service auto discovery mechanism

__其它名字空间内如何发现服务__

使用helm名字空间内的openshift-webtty容器进行验证

首先是使用内置的kubectl命令工具获取所在集群的服务，包括自身，注意Kuberentes为它挂载了一个volume

### Study case: openshift web tty

The web tty is [forked and customized from krishnasrinivas/wetty](https://github.com/stackdocker/wetty)

* Login and run kubectl client

![kubectl](/images/service-auto-discovery/屏幕快照%202016-06-10%20下午3.35.12.png?raw=true "Optional Title")

* From helm namespace to access services of another namespaces

nc-http, redis-sentinel, openshift

![access services](/images/service-auto-discovery/屏幕快照%202016-06-10%20下午3.44.06.png?raw=true "Optional Title")

* Inspect itself

![inspect itself](/images/service-auto-discovery/屏幕快照%202016-06-10%20下午3.57.36.png?raw=true "Optional Title")

### Study case: if have not kubectl in container

其中volume中的证书，名字空间，token以及之前所示的环境变量恰恰是Kubernetes为每个容器以非客户端证书访问kube-apiserver的参数，Kuberentes client api上有实现方法

According first case, a container already know Kubernetes services

According front inspecting, Kubernetes not only inject environments, but also mount a secret volume

![get sa and secret](/images/service-auto-discovery/屏幕快照%202016-06-10%20下午4.14.00.png?raw=true "Optional Title")

The provided service account (namespace, token, ca) is equal to:

![get sa and secret](/images/service-auto-discovery/屏幕快照%202016-06-10%20下午4.25.38.png?raw=true "Optional Title")

This is called in-cluster config, [The golang source code of Kubernetes client api](https://github.com/kubernetes/kubernetes/blob/master/pkg/client/unversioned/clientcmd/client_config.go#L345-L382)

## Other resources of the development environment

* ReplicationControllers

Each application and Kuberentes addon are configured with replication controller. the replication controller mechanisim will keep applicaiton alived although it is only one instance.

    [vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig get rc --all-namespaces
    NAMESPACE          NAME               DESIRED   CURRENT   AGE
    default            redis              2         2         7d
    default            redis-sentinel     3         3         7d
    helm               openshift-webtty   1         1         10d
    helm               tiller-rc          1         1         29d
    kube-system        kube-dns-v9        1         1         30d
    openshift-origin   etcd               1         1         30d
    openshift-origin   openshift          1         1         30d

* PersistenVolumes and PersistentVolumeClaims

Openshift is using these resources:

    vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig get pv
    NAME           CAPACITY   ACCESSMODES   STATUS    CLAIM                           REASON    AGE
    etcd-storage   2Gi        RWO,ROX,RWX   Bound     openshift-origin/etcd-storage             30d

    [vagrant@localhost origin]$ kubectl --kubeconfig=kubeconfig get pvc --all-namespaces
    NAMESPACE          NAME           STATUS    VOLUME         CAPACITY   ACCESSMODES   AGE
    openshift-origin   etcd-storage   Bound     etcd-storage   2Gi        RWO,ROX,RWX   30d

