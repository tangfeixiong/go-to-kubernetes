# go-for-kubernetes
Learning, exercising and working on Kubernetes

## Install a Docker-based Kubernetes cluster

The installation can be referenced in [kubernetes v1.0 documentation - Getting Started - Running locally via docker](http://kubernetes.io/v1.0/docs/getting-started-guides/docker.html)

* Run etcd for cluster configuration management

>`# docker run --net=host -d gcr.io/google_containers/etcd:2.0.9 /usr/local/bin/etcd --addr=127.0.0.1:4001 --bind-addr=0.0.0.0:4001 --data-dir=/var/etcd/data
`

* Run master

the master include controll, scheduler, POD, api... so enable TCP 8080 port or API will be failed

>`# docker run --net=host -d -v /var/run/docker.sock:/var/run/docker.sock  gcr.io/google_containers/hyperkube:v0.21.2 /hyperkube kubelet --api_servers=http://localhost:8080 --v=2 --address=0.0.0.0 --enable_server --hostname_override=127.0.0.1 --config=/etc/kubernetes/manifests
`

* Run service proxy

it is used to expose service with endpoint for external networking

>`# docker run -d --net=host --privileged gcr.io/google_containers/hyperkube:v0.21.2 /hyperkube proxy --master=http://127.0.0.1:8080 --v=2
`

* Run a nginx demo

>`# kubectl -s http://localhost:8080 run-container nginx --image=nginx --port=80
`

* Expose nginx as networing service

enable TCP 80 port or service can not be created

>`# kubectl expose rc nginx --port=80`


## Now test it out

* via endpoint

>`# /opt/tfx/kubectl get endpoints`

>`NAME         ENDPOINTS`

>`kubernetes   192.168.0.25:6443`

>`nginx        172.17.0.1:80`

>`# curl http://172.17.0.1`

>`<!DOCTYPE html>`

>`<html>`

>`<head>`

>`<title>Welcome to nginx!</title>`

>`<style>`

>`    body {`

>`        width: 35em;`

>`        margin: 0 auto;`

>`        font-family: Tahoma, Verdana, Arial, sans-serif;`

>`    }`

>`</style>`

>`</head>`

>`<body>`

>`<h1>Welcome to nginx!</h1>`

>`<p>If you see this page, the nginx web server is successfully installed and`

>`working. Further configuration is required.</p>`

>`<p>For online documentation and support please refer to`

>`<a href="http://nginx.org/">nginx.org</a>.<br/>`

>`Commercial support is available at`

>`<a href="http://nginx.com/">nginx.com</a>.</p>`

>`<p><em>Thank you for using nginx.</em></p>`

>`</body>`

>`</html>`

* via networking service

>`# /opt/tfx/kubectl get services`

>`NAME         LABELS                                    SELECTOR    IP(S)        PORT(S)`

>`kubernetes   component=apiserver,provider=kubernetes   <none>      10.0.0.1     443/TCP`

>`nginx        run=nginx                                 run=nginx   10.0.0.206   80/TCP`

>`# curl http://10.0.0.206`

>the same result as above

## list container with Docker

>`# docker ps`

>CONTAINER ID        IMAGE                                        COMMAND                CREATED             STATUS              PORTS               NAMES

>58ca7b7bc89a        nginx                                        "nginx -g 'daemon of   59 minutes ago      Up 59 minutes                           k8s_nginx.d7d3eb2f_nginx-syeg7_default_fc945b19-41a1-11e5-b8c4-c4346b46de6e_24ac6b02

>ffed942cfc52        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube controll   59 minutes ago      Up 59 minutes                           k8s_controller-manager.aad1ee8f_k8s-master-127.0.0.1_default_9b44830745c166dfc6d027b0fc2df36d_8441f980

>ac7cbe846daf        gcr.io/google_containers/pause:0.8.0         "/pause"
        59 minutes ago      Up 59 minutes                           k8s_POD.ef28
e851_nginx-syeg7_default_fc945b19-41a1-11e5-b8c4-c4346b46de6e_eeac80e6

>0af559d7e351        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube schedule   59 minutes ago      Up 59 minutes                           k8s_scheduler.b725e775_k8s-master-127.0.0.1_default_9b44830745c166dfc6d027b0fc2df36d_8d259834

>fa8bac1ebcfb        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube apiserve   59 minutes ago      Up 59 minutes                           k8s_apiserver.70750283_k8s-master-127.0.0.1_default_9b44830745c166dfc6d027b0fc2df36d_6c3cd54
a

>608b86a418dd        gcr.io/google_containers/pause:0.8.0         "/pause"
        59 minutes ago      Up 59 minutes                           k8s_POD.e4cc
795_k8s-master-127.0.0.1_default_9b44830745c166dfc6d027b0fc2df36d_4c85ae5c

>706a1dd60111        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube proxy --   10 days ago         Up 59 minutes                           sharp_bell


>8f4172000232        gcr.io/google_containers/hyperkube:v0.21.2   "/hyperkube kubelet    10 days ago         Up About an hour                        suspicious_albattani

>cfb63e647cac        gcr.io/google_containers/etcd:2.0.9          "/usr/local/bin/etcd   10 days ago         Up About an hour                        gloomy_jones


