

## releases

get binary packages
```
fanhonglingdeMacBook-Pro:v1.8.4 fanhongling$ curl -jkSL -C- https://dl.k8s.io/v1.8.4/kubernetes-server-linux-amd64.tar.gz -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     27      0  0:00:05  0:00:05 --:--:--    48
100  385M  100  385M    0     0  85228      0  1:18:59  1:18:59 --:--:-- 40587
```

```
fanhonglingdeMacBook-Pro:v1.8.4 fanhongling$ tar -ztf kubernetes-server-linux-amd64.tar.gz 
kubernetes/
kubernetes/server/
kubernetes/server/bin/
kubernetes/server/bin/kube-proxy.tar
kubernetes/server/bin/kubectl
kubernetes/server/bin/kube-apiserver.tar
kubernetes/server/bin/apiextensions-apiserver
kubernetes/server/bin/kubeadm
kubernetes/server/bin/kube-apiserver.docker_tag
kubernetes/server/bin/kube-aggregator.docker_tag
kubernetes/server/bin/hyperkube
kubernetes/server/bin/kube-aggregator.tar
kubernetes/server/bin/cloud-controller-manager
kubernetes/server/bin/kube-scheduler.docker_tag
kubernetes/server/bin/kubefed
kubernetes/server/bin/kube-controller-manager
kubernetes/server/bin/kube-proxy
kubernetes/server/bin/kube-scheduler
kubernetes/server/bin/cloud-controller-manager.tar
kubernetes/server/bin/kube-scheduler.tar
kubernetes/server/bin/kubelet
kubernetes/server/bin/kube-controller-manager.docker_tag
kubernetes/server/bin/kube-apiserver
kubernetes/server/bin/kube-aggregator
kubernetes/server/bin/kube-controller-manager.tar
kubernetes/server/bin/cloud-controller-manager.docker_tag
kubernetes/server/bin/kube-proxy.docker_tag
kubernetes/LICENSES
kubernetes/kubernetes-src.tar.gz
kubernetes/addons/
```

get addons manifests
```
fanhonglingdeMacBook-Pro:v1.6.4 fanhongling$ curl -jkSL -C- https://dl.k8s.io/v1.6.4/kubernetes.tar.gz -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     18      0  0:00:08  0:00:08 --:--:--    35
100 5943k  100 5943k    0     0  27118      0  0:03:44  0:03:44 --:--:-- 57422
```

```
fanhonglingdeMacBook-Pro:v1.6.4 fanhongling$ tar -ztf kubernetes.tar.gz 
### snippets ###
kubernetes/cluster/addons/dashboard/
kubernetes/cluster/addons/dashboard/dashboard-controller.yaml
kubernetes/cluster/addons/dashboard/dashboard-service.yaml
kubernetes/cluster/addons/dashboard/MAINTAINERS.md
kubernetes/cluster/addons/dashboard/README.md
### snippets ###
```
