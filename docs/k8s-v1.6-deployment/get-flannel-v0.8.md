
Download
```
fanhonglingdeMacBook-Pro:linux-bin fanhongling$ curl -jkSL -C- https://github.com/coreos/flannel/releases/download/v0.8.0-rc1/flanneld-amd64 -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   587    0   587    0     0    396      0 --:--:--  0:00:01 --:--:--   396
100 55.4M  100 55.4M    0     0   883k      0  0:01:04  0:01:04 --:--:--  868k


vagrant@vagrant-ubuntu-trusty-64:~$ docker pull quay.io/coreos/flannel:v0.8.0-rc1-amd64
v0.8.0-rc1-amd64: Pulling from coreos/flannel
1e76f742da49: Pull complete 
71c47dd8f443: Pull complete 
18a2f4e9176e: Pull complete 
a0e4c93cae64: Pull complete 
6183ceb2de5c: Pull complete 
33d4719e2f96: Pull complete 
Digest: sha256:181d38565c1633f6da33a604a6a368660ba52fb973d4d63ac0a0f225bb435dfc
Status: Downloaded newer image for quay.io/coreos/flannel:v0.8.0-rc1-amd64


vagrant@vagrant-ubuntu-trusty-64:~$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel-rbac.yml -o-
# Create the clusterrole and clusterrolebinding:
# $ kubectl create -f kube-flannel-rbac.yml
# Create the pod using the same namespace used by the flannel serviceaccount:
# $ kubectl create --namespace kube-system -f kube-flannel.yml
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


vagrant@vagrant-ubuntu-trusty-64:~$ curl -jkSL https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml -o-
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
      containers:
      - name: kube-flannel
        image: quay.io/coreos/flannel:v0.7.1-amd64
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
      - name: install-cni
        image: quay.io/coreos/flannel:v0.7.1-amd64
        command: [ "/bin/sh", "-c", "set -e -x; cp -f /etc/kube-flannel/cni-conf.json /etc/cni/net.d/10-flannel.conf; while true; do sleep 3600; done" ]
        volumeMounts:
        - name: cni
          mountPath: /etc/cni/net.d
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