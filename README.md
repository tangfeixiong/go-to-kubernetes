# go-for-kubernetes
Learning, exercising and working on Kubernetes

## Install a Docker-based Kubernetes cluster

The installation can be referenced in [kubernetes v1.0 documentation - Getting Started - Running locally via docker](http://kubernetes.io/v1.0/docs/getting-started-guides/docker.html)

* Run etcd for cluster configuration management

>`# docker run --net=host -d gcr.io/google_containers/etcd:2.0.9 /usr/local/bin/etcd --addr=127.0.0.1:4001 --bind-addr=0.0.0.0:4001 --data-dir=/var/etcd/data
`

* Run master

the master include scheduler, api... so enable TCP 8080 port or API will be failed

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

