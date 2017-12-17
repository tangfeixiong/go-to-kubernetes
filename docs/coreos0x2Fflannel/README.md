## DevOps

### Table of contents
1. v0.9.1
1. v0.9.0

### v0.9.1

Refer to https://github.com/coreos/flannel for Kubernetes 1.7+

For kubernetes v1.9.0
```
[vagrant@kubedev-172-17-4-59 ~]$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml -o /tmp/kube-flannel.yaml  
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  3215  100  3215    0     0   4040      0 --:--:-- --:--:-- --:--:--  4038
```

```
[vagrant@kubedev-172-17-4-59 ~]$ img=$(cat /tmp/kube-flannel.yaml | grep ' image: ' | head -1 | awk '{print $2}');  ver=$(echo $img | awk -F'[:-]' '{print $2}' | cut -c2-); docker pull $img; mv /tmp/kube-flannel.yaml /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/docs/coreos0x2Fflannel/kube-flannel_${ver}.yaml
Trying to pull repository quay.io/coreos/flannel ... 
sha256:056cf57fd3bbe7264c0be1a3b34ec2e289b33e51c70f332f4e88aa83970ad891: Pulling from quay.io/coreos/flannel
6d987f6f4279: Pull complete 
16a827ca53c8: Pull complete 
8f4dde5859ad: Pull complete 
40db9b39f697: Pull complete 
48f4941e520f: Pull complete 
f77b10f3d93e: Pull complete 
Digest: sha256:056cf57fd3bbe7264c0be1a3b34ec2e289b33e51c70f332f4e88aa83970ad891
Status: Downloaded newer image for quay.io/coreos/flannel:v0.9.1-amd64
```

Legacy
```
[vagrant@kubedev-172-17-4-59 ~]$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/master/Documentation/k8s-manifests/kube-flannel-legacy.yml -o /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/docs/coreos0x2Fflannel/legacy/kube-flannel-legacy_${ver}.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2281  100  2281    0     0   3568      0 --:--:-- --:--:-- --:--:--  3564
```

```
[vagrant@kubedev-172-17-4-59 ~]$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/master/Documentation/k8s-manifests/kube-flannel-rbac.yml -o /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/docs/coreos0x2Fflannel/legacy/kube-flannel-rbac_${ver}.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   857  100   857    0     0   1099      0 --:--:-- --:--:-- --:--:--  1100
```

### v0,9.0

For kubernetes v1.8.4
```
fanhonglingdeMacBook-Pro:coreos0x2Fflannel fanhongling$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml | tee kube-flannel.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2830  100  2830    0     0    186      0  0:00:15  0:00:15 --:--:--   751
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: flannel
rules:
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - get
  - apiGroups:
      - ""
    resources:
      - nodes
    verbs:
      - list
      - watch
  - apiGroups:
      - ""
    resources:
      - nodes/status
    verbs:
      - patch
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: flannel
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: flannel
subjects:
- kind: ServiceAccount
  name: flannel
  namespace: kube-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: flannel
  namespace: kube-system
---
kind: ConfigMap
apiVersion: v1
metadata:
  name: kube-flannel-cfg
  namespace: kube-system
  labels:
    tier: node
    app: flannel
data:
  cni-conf.json: |
    {
      "name": "cbr0",
      "type": "flannel",
      "delegate": {
        "hairpinMode": true,
        "isDefaultGateway": true
      }
    }
  net-conf.json: |
    {
      "Network": "10.244.0.0/16",
      "Backend": {
        "Type": "vxlan"
      }
    }
---
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: kube-flannel-ds
  namespace: kube-system
  labels:
    tier: node
    app: flannel
spec:
  template:
    metadata:
      labels:
        tier: node
        app: flannel
    spec:
      hostNetwork: true
      nodeSelector:
        beta.kubernetes.io/arch: amd64
      tolerations:
      - key: node-role.kubernetes.io/master
        operator: Exists
        effect: NoSchedule
      serviceAccountName: flannel
      initContainers:
      - name: install-cni
        image: quay.io/coreos/flannel:v0.9.0-amd64
        command:
        - cp
        args:
        - -f
        - /etc/kube-flannel/cni-conf.json
        - /etc/cni/net.d/10-flannel.conf
        volumeMounts:
        - name: cni
          mountPath: /etc/cni/net.d
        - name: flannel-cfg
          mountPath: /etc/kube-flannel/
      containers:
      - name: kube-flannel
        image: quay.io/coreos/flannel:v0.9.0-amd64
        command: [ "/opt/bin/flanneld", "--ip-masq", "--kube-subnet-mgr" ]
        securityContext:
          privileged: true
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        volumeMounts:
        - name: run
          mountPath: /run
        - name: flannel-cfg
          mountPath: /etc/kube-flannel/
      volumes:
        - name: run
          hostPath:
            path: /run
        - name: cni
          hostPath:
            path: /etc/cni/net.d
        - name: flannel-cfg
          configMap:
            name: kube-flannel-cfg
```

Legacy, refert to (https://github.com/coreos/flannel/blob/master/Documentation/kubernetes.md
```
fanhonglingdeMacBook-Pro:before1.6 fanhongling$ curl -jsSL https://raw.githubusercontent.com/coreos/flannel/master/Documentation/k8s-manifests/kube-flannel-legacy.yml -O
```

```
fanhonglingdeMacBook-Pro:before1.6 fanhongling$ curl -jsSL https://raw.githubusercontent.com/coreos/flannel/master/Documentation/k8s-manifests/kube-flannel-rbac.yml -O
```

Images
```
[vagrant@localhost go-to-kubernetes]$ docker pull quay.io/coreos/flannel:v0.9.0-amd64
Trying to pull repository quay.io/coreos/flannel ... 
v0.9.0-amd64: Pulling from quay.io/coreos/flannel
6d987f6f4279: Already exists 
16a827ca53c8: Pull complete 
8f4dde5859ad: Pull complete 
2a4b88adf2f4: Pull complete 
60432d0f9d91: Pull complete 
46d9b18a374a: Pull complete 
Digest: sha256:1b401bf0c30bada9a539389c3be652b58fe38463361edf488e6543c8761d4970
Status: Downloaded newer image for quay.io/coreos/flannel:v0.9.0-amd64
```