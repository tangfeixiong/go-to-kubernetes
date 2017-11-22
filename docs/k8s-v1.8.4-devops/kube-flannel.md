__Flannal Networking__

For manifests , refer to [../coreos0x2Fflannel](../coreos0x2Fflannel)

Modify kube-flannel.yaml because it is vagrant environment
```
        command: [ "/opt/bin/flanneld", "--ip-masq", "--kube-subnet-mgr", "--iface=enp0s8" ]
```

Create
```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl create -f kube-flannel.yaml 
clusterrole "flannel" created
clusterrolebinding "flannel" created
serviceaccount "flannel" created
configmap "kube-flannel-cfg" created
daemonset "kube-flannel-ds" created
```

```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl get pods --all-namespaces
NAMESPACE     NAME                                          READY     STATUS    RESTARTS   AGE
kube-system   etcd-kubedev-10-64-33-82                      1/1       Running   7          2h
kube-system   kube-apiserver-kubedev-10-64-33-82            1/1       Running   3          2h
kube-system   kube-controller-manager-kubedev-10-64-33-82   1/1       Running   3          2h
kube-system   kube-dns-545bc4bfd4-2xk7v                     3/3       Running   0          2h
kube-system   kube-flannel-ds-4z8jc                         1/1       Running   0          32s
kube-system   kube-proxy-7fdvp                              1/1       Running   3          2h
kube-system   kube-scheduler-kubedev-10-64-33-82            1/1       Running   4          2h
```

