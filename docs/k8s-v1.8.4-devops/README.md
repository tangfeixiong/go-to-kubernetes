

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
fanhonglingdeMacBook-Pro:v1.8.4 fanhongling$ curl -jkSL -C- https://dl.k8s.io/v1.8.4/kubernetes.tar.gz -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   161  100   161    0     0     17      0  0:00:09  0:00:09 --:--:--    36
100 4029k  100 4029k    0     0  87544      0  0:00:47  0:00:47 --:--:--  189k
```

```
fanhonglingdeMacBook-Pro:v1.8.4 fanhongling$ tar -ztf kubernetes.tar.gz 
### snippets ###
kubernetes/cluster/addons/
kubernetes/cluster/addons/dns/
kubernetes/cluster/addons/dns/kubedns-controller.yaml.in
kubernetes/cluster/addons/dns/kubedns-svc.yaml.base
kubernetes/cluster/addons/dns/OWNERS
kubernetes/cluster/addons/dns/kubedns-controller.yaml.base
kubernetes/cluster/addons/dns/Makefile
kubernetes/cluster/addons/dns/kubedns-controller.yaml.sed
kubernetes/cluster/addons/dns/kubedns-cm.yaml
kubernetes/cluster/addons/dns/transforms2salt.sed
kubernetes/cluster/addons/dns/kubedns-svc.yaml.sed
kubernetes/cluster/addons/dns/transforms2sed.sed
kubernetes/cluster/addons/dns/kubedns-sa.yaml
kubernetes/cluster/addons/dns/kubedns-svc.yaml.in
kubernetes/cluster/addons/dns/README.md
### snippets ###
```
