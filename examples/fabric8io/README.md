# Instruction

Deploy

[gofabric8](./gofabric8.md)

Fabric8 system
```
[vagrant@localhost ~]$ kubectl --namespace=fabric8-system get all
NAME                   DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/ingress-nginx   1         1         1            1           180d

NAME                          DESIRED   CURRENT   READY     AGE
rs/ingress-nginx-2447702149   0         0         0         180d
rs/ingress-nginx-775506543    1         1         1         180d

NAME                   DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/ingress-nginx   1         1         1            1           180d

NAME                               READY     STATUS    RESTARTS   AGE
po/ingress-nginx-775506543-vfzgs   1/1       Running   0          1d
```

Components
```
[vagrant@localhost ~]$ kubectl get pods
NAME                                       READY     STATUS     RESTARTS   AGE
configmapcontroller-4273343753-kh4rk       1/1       Running    1          110d
exposecontroller-3102591648-888vh          1/1       Running    3          110d
fabric8-3873669821-lf9qc                   2/2       Running    2          110d
fabric8-docker-registry-3263760418-k4ff7   0/1       Init:0/1   0          110d
fabric8-forge-1088523184-cv7lg             1/1       Running    1          110d
gogs-3486434875-d6pwf                      0/1       Init:0/1   0          110d
jenkins-4278279665-5tt3g                   0/1       Init:0/1   0          110d
nexus-4015917150-204dh                     0/1       Init:0/1   0          110d
```

_Somethins is not working because it is a standalone small node_

From Dashboard

[屏幕快照 2017-12-16 下午1.18.18.png](./屏幕快照%202017-12-16%20下午1.18.18.png)

[屏幕快照 2017-12-16 下午1.23.34.png](./屏幕快照%202017-12-16%20下午1.23.34.png)