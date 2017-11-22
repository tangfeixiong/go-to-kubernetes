
```
[vagrant@openshiftdev ~]$ echo "
> apiVersion: v1
> kind: List
> items:
>   - apiVersion: v1
>     kind: ServiceAccount
>     metadata:
>       name: weave-net
>       labels:
>         name: weave-net
>       namespace: kube-system
>   - apiVersion: rbac.authorization.k8s.io/v1beta1
>     kind: ClusterRole
>     metadata:
>       name: weave-net
>       labels:
>         name: weave-net
>     rules:
>       - apiGroups:
>           - ''
>         resources:
>           - pods
>           - namespaces
>           - nodes
>         verbs:
>           - get
>           - list
>           - watch
>       - apiGroups:
>           - extensions
>         resources:
>           - networkpolicies
>         verbs:
>           - get
>           - list
>           - watch
>   - apiVersion: rbac.authorization.k8s.io/v1beta1
>     kind: ClusterRoleBinding
>     metadata:
>       name: weave-net
>       labels:
>         name: weave-net
>     roleRef:
>       kind: ClusterRole
>       name: weave-net
>       apiGroup: rbac.authorization.k8s.io
>     subjects:
>       - kind: ServiceAccount
>         name: weave-net
>         namespace: kube-system
>   - apiVersion: rbac.authorization.k8s.io/v1beta1
>     kind: Role
>     metadata:
>       name: weave-net-kube-peer
>       namespace: kube-system
>       labels:
>         name: weave-net-kube-peer
>     rules:
>       - apiGroups:
>           - ''
>         resources:
>           - configmaps
>         resourceNames:
>           - weave-net
>         verbs:
>           - get
>           - update
>       - apiGroups:
>           - ''
>         resources:
>           - configmaps
>         verbs:
>           - create
>   - apiVersion: rbac.authorization.k8s.io/v1beta1
>     kind: RoleBinding
>     metadata:
>       name: weave-net-kube-peer
>       namespace: kube-system
>       labels:
>         name: weave-net-kube-peer
>     roleRef:
>       kind: Role
>       name: weave-net-kube-peer
>       apiGroup: rbac.authorization.k8s.io
>     subjects:
>       - kind: ServiceAccount
>         name: weave-net
>         namespace: kube-system
>   - apiVersion: extensions/v1beta1
>     kind: DaemonSet
>     metadata:
>       name: weave-net
>       labels:
>         name: weave-net
>       namespace: kube-system
>     spec:
>       template:
>         metadata:
>           labels:
>             name: weave-net
>         spec:
>           containers:
>             - name: weave
>               command:
>                 - /home/weave/launch.sh
>               env:
>                 - name: HOSTNAME
>                   valueFrom:
>                     fieldRef:
>                       apiVersion: v1
>                       fieldPath: spec.nodeName
>               image: 'weaveworks/weave-kube:2.1.1'
>               livenessProbe:
>                 httpGet:
>                   host: 127.0.0.1
>                   path: /status
>                   port: 6784
>                 initialDelaySeconds: 30
>               resources:
>                 requests:
>                   cpu: 10m
>               securityContext:
>                 privileged: true
>               volumeMounts:
>                 - name: weavedb
>                   mountPath: /weavedb
>                 - name: cni-bin
>                   mountPath: /host/opt
>                 - name: cni-bin2
>                   mountPath: /host/home
>                 - name: cni-conf
>                   mountPath: /host/etc
>                 - name: dbus
>                   mountPath: /host/var/lib/dbus
>                 - name: lib-modules
>                   mountPath: /lib/modules
>             - name: weave-npc
>               env:
>                 - name: HOSTNAME
>                   valueFrom:
>                     fieldRef:
>                       apiVersion: v1
>                       fieldPath: spec.nodeName
>               image: 'weaveworks/weave-npc:2.1.1'
>               resources:
>                 requests:
>                   cpu: 10m
>               securityContext:
>                 privileged: true
>           hostNetwork: true
>           hostPID: true
>           restartPolicy: Always
>           securityContext:
>             seLinuxOptions: {}
>           serviceAccountName: weave-net
>           tolerations:
>             - effect: NoSchedule
>               operator: Exists
>           volumes:
>             - name: weavedb
>               hostPath:
>                 path: /var/lib/weave
>             - name: cni-bin
>               hostPath:
>                 path: /opt
>             - name: cni-bin2
>               hostPath:
>                 path: /home
>             - name: cni-conf
>               hostPath:
>                 path: /etc
>             - name: dbus
>               hostPath:
>                 path: /var/lib/dbus
>             - name: lib-modules
>               hostPath:
>                 path: /lib/modules
>       updateStrategy:
>         type: RollingUpdate
> " > weave-daemonset-k8s-1.6.yaml
```

```
[vagrant@openshiftdev ~]$ kubectl create -f weave-daemonset-k8s-1.6.yaml 
serviceaccount "weave-net" created
clusterrole "weave-net" created
clusterrolebinding "weave-net" created
role "weave-net-kube-peer" created
rolebinding "weave-net-kube-peer" created
daemonset "weave-net" created
```

```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl get all --all-namespaces
NAMESPACE     NAME                                             READY     STATUS             RESTARTS   AGE
kube-system   po/etcd-kubedev-10-64-33-82                      1/1       Running            4          1h
kube-system   po/kube-apiserver-kubedev-10-64-33-82            1/1       Running            4          1h
kube-system   po/kube-controller-manager-kubedev-10-64-33-82   1/1       Running            6          1h
kube-system   po/kube-scheduler-kubedev-10-64-33-82            1/1       Running            4          1h
kube-system   po/weave-net-1l2sr                               1/2       CrashLoopBackOff   6          20m
```

```
[vagrant@kubedev-10-64-33-82 ~]$ kubectl logs po/weave-net-1l2sr --namespace=kube-system
Error from server (BadRequest): a container name must be specified for pod weave-net-1l2sr, choose one of: [weave weave-npc]
[vagrant@kubedev-10-64-33-82 ~]$ kubectl logs po/weave-net-1l2sr --namespace=kube-system -c weave
Error from server: Get https://10.64.33.82:10250/containerLogs/kube-system/weave-net-1l2sr/weave: dial tcp 10.64.33.82:10250: getsockopt: connection refused
[vagrant@kubedev-10-64-33-82 ~]$ kubectl logs po/weave-net-1l2sr --namespace=kube-system -c weave-npc
Error from server: Get https://10.64.33.82:10250/containerLogs/kube-system/weave-net-1l2sr/weave-npc: dial tcp 10.64.33.82:10250: getsockopt: connection refused
```

## Issue

https://github.com/kubernetes/kubernetes/issues/43815

remove $KUBELET_NETWORK_ARGS in /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

https://github.com/weaveworks/weave/issues/2820



```
[vagrant@openshiftdev ~]$ echo "
> apiVersion: v1
> kind: List
> items:
>   - apiVersion: v1
>     kind: ServiceAccount
>     metadata:
>       name: weave-net
>       labels:
>         name: weave-net
>       namespace: kube-system
>   - apiVersion: rbac.authorization.k8s.io/v1beta1
>     kind: ClusterRole
>     metadata:
>       name: weave-net
>       labels:
>         name: weave-net
>     rules:
>       - apiGroups:
>           - ''
>         resources:
>           - pods
>           - namespaces
>           - nodes
>         verbs:
>           - get
>           - list
>           - watch
>       - apiGroups:
>           - extensions
>         resources:
>           - networkpolicies
>         verbs:
>           - get
>           - list
>           - watch
>       - apiGroups:
>           - 'networking.k8s.io'
>         resources:
>           - networkpolicies
>         verbs:
>           - get
>           - list
>           - watch
>   - apiVersion: rbac.authorization.k8s.io/v1beta1
>     kind: ClusterRoleBinding
>     metadata:
>       name: weave-net
>       labels:
>         name: weave-net
>     roleRef:
>       kind: ClusterRole
>       name: weave-net
>       apiGroup: rbac.authorization.k8s.io
>     subjects:
>       - kind: ServiceAccount
>         name: weave-net
>         namespace: kube-system
>   - apiVersion: rbac.authorization.k8s.io/v1beta1
>     kind: Role
>     metadata:
>       name: weave-net
>       namespace: kube-system
>       labels:
>         name: weave-net
>     rules:
>       - apiGroups:
>           - ''
>         resources:
>           - configmaps
>         resourceNames:
>           - weave-net
>         verbs:
>           - get
>           - update
>       - apiGroups:
>           - ''
>         resources:
>           - configmaps
>         verbs:
>           - create
>   - apiVersion: rbac.authorization.k8s.io/v1beta1
>     kind: RoleBinding
>     metadata:
>       name: weave-net
>       namespace: kube-system
>       labels:
>         name: weave-net
>     roleRef:
>       kind: Role
>       name: weave-net
>       apiGroup: rbac.authorization.k8s.io
>     subjects:
>       - kind: ServiceAccount
>         name: weave-net
>         namespace: kube-system
>   - apiVersion: extensions/v1beta1
>     kind: DaemonSet
>     metadata:
>       name: weave-net
>       labels:
>         name: weave-net
>       namespace: kube-system
>     spec:
>       template:
>         metadata:
>           labels:
>             name: weave-net
>         spec:
>           containers:
>             - name: weave
>               command:
>                 - /home/weave/launch.sh
>               env:
>                 - name: HOSTNAME
>                   valueFrom:
>                     fieldRef:
>                       apiVersion: v1
>                       fieldPath: spec.nodeName
>               image: 'weaveworks/weave-kube:2.1.1'
>               livenessProbe:
>                 httpGet:
>                   host: 127.0.0.1
>                   path: /status
>                   port: 6784
>                 initialDelaySeconds: 30
>               resources:
>                 requests:
>                   cpu: 10m
>               securityContext:
>                 privileged: true
>               volumeMounts:
>                 - name: weavedb
>                   mountPath: /weavedb
>                 - name: cni-bin
>                   mountPath: /host/opt
>                 - name: cni-bin2
>                   mountPath: /host/home
>                 - name: cni-conf
>                   mountPath: /host/etc
>                 - name: dbus
>                   mountPath: /host/var/lib/dbus
>                 - name: lib-modules
>                   mountPath: /lib/modules
>                 - name: xtables-lock
>                   mountPath: /run/xtables.lock
>                   readOnly: false
>             - name: weave-npc
>               env:
>                 - name: HOSTNAME
>                   valueFrom:
>                     fieldRef:
>                       apiVersion: v1
>                       fieldPath: spec.nodeName
>               image: 'weaveworks/weave-npc:2.1.1'
> #npc-args
>               resources:
>                 requests:
>                   cpu: 10m
>               securityContext:
>                 privileged: true
>               volumeMounts:
>                 - name: xtables-lock
>                   mountPath: /run/xtables.lock
>                   readOnly: false
>           hostNetwork: true
>           hostPID: true
>           restartPolicy: Always
>           securityContext:
>             seLinuxOptions: {}
>           serviceAccountName: weave-net
>           tolerations:
>             - effect: NoSchedule
>               operator: Exists
>           volumes:
>             - name: weavedb
>               hostPath:
>                 path: /var/lib/weave
>             - name: cni-bin
>               hostPath:
>                 path: /opt
>             - name: cni-bin2
>               hostPath:
>                 path: /home
>             - name: cni-conf
>               hostPath:
>                 path: /etc
>             - name: dbus
>               hostPath:
>                 path: /var/lib/dbus
>             - name: lib-modules
>               hostPath:
>                 path: /lib/modules
>             - name: xtables-lock
>               hostPath:
>                 path: /run/xtables.lock
>       updateStrategy:
>         type: RollingUpdate
> " > weave-daemonset-k8s-1.7.yaml
```

```
[vagrant@openshiftdev ~]$ kubectl apply -f weave-daemonset-k8s-1.7.yaml 
serviceaccount "weave-net" created
clusterrole "weave-net" created
clusterrolebinding "weave-net" created
role "weave-net" created
rolebinding "weave-net" created
daemonset "weave-net" created
```


```
[vagrant@kubedev-10-64-33-82 ~]$ ls /opt/cni/bin/
bridge  cnitool  dhcp  flannel  host-local  ipvlan  loopback  macvlan  noop  ptp  tuning
```