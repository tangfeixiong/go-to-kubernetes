

## releases

get binary packages
```
fanhonglingdeMacBook-Pro:v1.6.4 fanhongling$ curl -jkSL -C- https://dl.k8s.io/v1.6.4/kubernetes-server-linux-amd64.tar.gz -O

```

```
fanhonglingdeMacBook-Pro:v1.6.4 fanhongling$ tar -ztf kubernetes-server-linux-amd64.tar.gz 

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
