Hand-on: ElasticSearch With StorageClass and StatefulSet
========================================================

Table of Contents
-----------------

* Reference
* Download
* hostpath-provisioner from kubernetes-incubator external-storage
* deploying ElasticSearch
* test of ElasticSearch

### Reference

https://github.com/pires/kubernetes-elasticsearch-cluster

and

https://github.com/pires/kubernetes-elasticsearch-cluster/tree/master/stateful

![resources](./屏幕快照%202017-05-13%20下午2.38.28.png)

Content
-------

### Download _manifests_

Prerequistes
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples$ mkdir elasticsearch && cd elasticsearch
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch$ mkdir https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch$ cd https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ mkdir ns-efk
```

With `curl` & `sed`
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/LICENSE -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 11325  100 11325    0     0  10738      0  0:00:01  0:00:01 --:--:-- 10744
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/README.md -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 18171  100 18171    0     0  15058      0  0:00:01  0:00:01 --:--:-- 15067
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster fanhongling$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/es-discovery-svc.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   253  100   253    0     0     90      0  0:00:02  0:00:02 --:--:--    90
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-discovery-svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: elasticsearch-discovery
  namespace: efk
  labels:
    component: elasticsearch
    role: master
spec:
  selector:
    component: elasticsearch
    role: master
  ports:
  - name: transport
    port: 9300
    protocol: TCP
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-discovery-svc.yaml >> ns-efk/es-discovery-svc.yaml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/es-svc.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   259  100   259    0     0    161      0  0:00:01  0:00:01 --:--:--   161
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/;s/LoadBalancer/NodePort/' es-svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: elasticsearch
  namespace: efk
  labels:
    component: elasticsearch
    role: client
spec:
  type: NodePort
  selector:
    component: elasticsearch
    role: client
  ports:
  - name: http
    port: 9200
    protocol: TCP
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/;s/LoadBalancer/NodePort/' es-svc.yaml >> ns-efk/es-svc.yaml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/es-master.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1788  100  1788    0     0   1285      0  0:00:01  0:00:01 --:--:--  1285
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-master.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: es-master
  namespace: efk
  labels:
    component: elasticsearch
    role: master
spec:
  replicas: 3
  template:
    metadata:
      labels:
        component: elasticsearch
        role: master
      annotations:
        pod.beta.kubernetes.io/init-containers: '[
          {
          "name": "sysctl",
            "image": "busybox",
            "imagePullPolicy": "IfNotPresent",
            "command": ["sysctl", "-w", "vm.max_map_count=262144"],
            "securityContext": {
              "privileged": true
            }
          }
        ]'
    spec:
      containers:
      - name: es-master
        securityContext:
          privileged: false
          capabilities:
            add:
              - IPC_LOCK
              - SYS_RESOURCE
        image: quay.io/pires/docker-elasticsearch-kubernetes:5.4.0
        imagePullPolicy: Always
        env:
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: "CLUSTER_NAME"
          value: "myesdb"
        - name: "NUMBER_OF_MASTERS"
          value: "2"
        - name: NODE_MASTER
          value: "true"
        - name: NODE_INGEST
          value: "false"
        - name: NODE_DATA
          value: "false"
        - name: HTTP_ENABLE
          value: "false"
        - name: "ES_JAVA_OPTS"
          value: "-Xms256m -Xmx256m"
        ports:
        - containerPort: 9300
          name: transport
          protocol: TCP
        volumeMounts:
        - name: storage
          mountPath: /data
      volumes:
          - emptyDir:
              medium: ""
            name: "storage"
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-master.yaml >> ns-efk/es-master.yaml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/es-client.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1753  100  1753    0     0   1675      0  0:00:01  0:00:01 --:--:--  1675
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-client.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: es-client
  namespace: efk
  labels:
    component: elasticsearch
    role: client
spec:
  replicas: 2
  template:
    metadata:
      labels:
        component: elasticsearch
        role: client
      annotations:
        pod.beta.kubernetes.io/init-containers: '[
          {
          "name": "sysctl",
            "image": "busybox",
            "imagePullPolicy": "IfNotPresent",
            "command": ["sysctl", "-w", "vm.max_map_count=262144"],
            "securityContext": {
              "privileged": true
            }
          }
        ]'
    spec:
      containers:
      - name: es-client
        securityContext:
          privileged: false
          capabilities:
            add:
              - IPC_LOCK
              - SYS_RESOURCE
        image: quay.io/pires/docker-elasticsearch-kubernetes:5.4.0
        imagePullPolicy: Always
        env:
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: "CLUSTER_NAME"
          value: "myesdb"
        - name: NODE_MASTER
          value: "false"
        - name: NODE_DATA
          value: "false"
        - name: HTTP_ENABLE
          value: "true"
        - name: "ES_JAVA_OPTS"
          value: "-Xms256m -Xmx256m"
        ports:
        - containerPort: 9200
          name: http
          protocol: TCP
        - containerPort: 9300
          name: transport
          protocol: TCP
        volumeMounts:
        - name: storage
          mountPath: /data
      volumes:
          - emptyDir:
              medium: ""
            name: "storage"
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-client.yaml >> ns-efk/es-client.yaml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-data.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: es-data
  namespace: efk
  labels:
    component: elasticsearch
    role: data
spec:
  replicas: 2
  template:
    metadata:
      labels:
        component: elasticsearch
        role: data
      annotations:
        pod.beta.kubernetes.io/init-containers: '[
          {
          "name": "sysctl",
            "image": "busybox",
            "imagePullPolicy": "IfNotPresent",
            "command": ["sysctl", "-w", "vm.max_map_count=262144"],
            "securityContext": {
              "privileged": true
            }
          }
        ]'
    spec:
      containers:
      - name: es-data
        securityContext:
          privileged: false
          capabilities:
            add:
              - IPC_LOCK
              - SYS_RESOURCE
        image: quay.io/pires/docker-elasticsearch-kubernetes:5.4.0
        imagePullPolicy: Always
        env:
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: "CLUSTER_NAME"
          value: "myesdb"
        - name: NODE_MASTER
          value: "false"
        - name: NODE_INGEST
          value: "false"
        - name: HTTP_ENABLE
          value: "false"
        - name: "ES_JAVA_OPTS"
          value: "-Xms256m -Xmx256m"
        ports:
        - containerPort: 9300
          name: transport
          protocol: TCP
        volumeMounts:
        - name: storage
          mountPath: /data
      volumes:
          - emptyDir:
              medium: ""
            name: "storage"
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-data.yaml >> ns-efk/es-data.yaml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/es-curator-config.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1383  100  1383    0     0   1338      0  0:00:01  0:00:01 --:--:--  1338
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^data:$\)/  namespace: efk\n\1/' es-curator-config.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: curator-config
  namespace: efk
data:
  action_file.yml: |-
    ---
    # Remember, leave a key empty if there is no value.  None will be a string,
    # not a Python "NoneType"
    #
    # Also remember that all examples have 'disable_action' set to True.  If you
    # want to use this action as a template, be sure to set this to False after
    # copying it.
    actions:
      1:
        action: delete_indices
        description: "Clean up ES by deleting old indices"
        options:
          timeout_override:
          continue_if_exception: False
          disable_action: False
        filters:
        - filtertype: age
          source: name
          direction: older
          timestring: '%Y.%m.%d'
          unit: days
          unit_count: 3
          field:
          stats_result:
          epoch:
          exclude: False
  config.yml: |-
    ---
    # Remember, leave a key empty if there is no value.  None will be a string,
    # not a Python "NoneType"
    client:
      hosts:
        - elasticsearch
      port: 9200
      url_prefix:
      use_ssl: False
      certificate:
      client_cert:
      client_key:
      ssl_no_validate: False
      http_auth:
      timeout: 30
      master_only: False

    logging:
      loglevel: INFO
      logfile:
      logformat: default
      blacklist: ['elasticsearch', 'urllib3']
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^data:$\)/  namespace: efk\n\1/' es-curator-config.yaml >> ns-efk/es-curator-config.yaml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/es-curator.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   571  100   571    0     0    548      0  0:00:01  0:00:01 --:--:--   549
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^spec:$\)/  namespace: efk\n\1/' es-curator.yaml
apiVersion: batch/v2alpha1
kind: CronJob
metadata:
  name: curator
  namespace: efk
spec:
  schedule: 1 0 * * *
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: curator
            image: bobrik/curator
            args: ["--config", "/etc/config/config.yml", "/etc/config/action_file.yml"]
            volumeMounts:
              - name: config-volume
                mountPath: /etc/config
          volumes:
            - name: config-volume
              configMap:
                name: curator-config
          restartPolicy: OnFailure
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ sed 's/\(^spec:$\)/  namespace: efk\n\1/' es-curator.yaml >> ns-efk/es-curator.yaml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster$ mkdir stateful
&& cd stateful
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/stateful/es-data-stateful.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1825  100  1825    0     0   1761      0  0:00:01  0:00:01 --:--:--  1763
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/;s/storage: 12Gi/storage: 100Mi/' es-data-stateful.yaml
apiVersion: apps/v1beta1
kind: StatefulSet
metadata:
  name: es-data
  namespace: efk
  labels:
    component: elasticsearch
    role: data
spec:
  serviceName: elasticsearch-data
  replicas: 3
  template:
    metadata:
      labels:
        component: elasticsearch
        role: data
      annotations:
        pod.beta.kubernetes.io/init-containers: '[
          {
          "name": "sysctl",
            "image": "busybox",
            "imagePullPolicy": "IfNotPresent",
            "command": ["sysctl", "-w", "vm.max_map_count=262144"],
            "securityContext": {
              "privileged": true
            }
          }
        ]'
    spec:
      containers:
      - name: es-data
        securityContext:
          privileged: true
          capabilities:
            add:
              - IPC_LOCK
        image: quay.io/pires/docker-elasticsearch-kubernetes:5.4.0
        imagePullPolicy: Always
        env:
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: "CLUSTER_NAME"
          value: "myesdb"
        - name: NODE_MASTER
          value: "false"
        - name: NODE_INGEST
          value: "false"
        - name: HTTP_ENABLE
          value: "false"
        - name: "ES_JAVA_OPTS"
          value: "-Xms256m -Xmx256m"
        ports:
        - containerPort: 9300
          name: transport
          protocol: TCP
        volumeMounts:
        - name: storage
          mountPath: /data
  volumeClaimTemplates:
  - metadata:
      name: storage
      annotations:
        volume.beta.kubernetes.io/storage-class: standard
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 100Mi
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/;s/storage: 12Gi/storage: 100Mi/' es-data-stateful.yaml >> ../ns-efk/es-data-stateful.yaml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/stateful/es-data-svc.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   244  100   244    0     0    105      0  0:00:02  0:00:02 --:--:--   105
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-data-svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: elasticsearch-data
  namespace: efk
  labels:
    component: elasticsearch
    role: data
spec:
  ports:
  - port: 9300
    name: transport
  clusterIP: None
  selector:
    component: elasticsearch
    role: data
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ sed 's/\(^  labels:$\)/  namespace: efk\n\1/' es-data-svc.yaml > ../ns-efk/es-data-stateful-svc.yaml
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ curl -jkSL https://raw.githubusercontent.com/pires/kubernetes-elasticsearch-cluster/master/stateful/gce-storage-class.yaml -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   166  100   166    0     0    160      0  0:00:01  0:00:01 --:--:--   161
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ sed '/parameters:/,$d;s%kubernetes.io/gce-pd%example.com/hostpath%' gce-storage-class.yaml 
kind: StorageClass
apiVersion: storage.k8s.io/v1beta1
metadata:
  name: standard
provisioner: example.com/hostpath
```

Specify _namespace_
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ cd ../ns-efk/
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ cat > es-namespace.yaml << EOF
> apiVersion: v1
> kind: Namespace
> metadata:
>   name: efk
> EOF
```

### Get hostpath-provisioner from https://github.com/kubernetes-incubator/external-storage

Download _hostpath-provisioner_ StorageClass
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/stateful$ curl -jkSL https://raw.githubusercontent.com/kubernetes-incubator/external-storage/master/docs/demo/hostpath-provisioner/pod.yaml -o hostpath-provisioner-pod.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   497  100   497    0     0    461      0  0:00:01  0:00:01 --:--:--   462
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ cat > hostpath-persisten-volume-claim.yaml << EOF
> kind: PersistentVolumeClaim
> apiVersion: v1
> metadata:
>   name: hostpath
>   annotations:
>     volume.beta.kubernetes.io/storage-class: "standard"
> spec:
>   accessModes:
>     - ReadWriteMany
>   resources:
>     requests:
>       storage: 1Mi
> EOF
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ curl -jkSL https://raw.githubusercontent.com/kubernetes-incubator/external-storage/master/docs/demo/hostpath-provisioner/class.yaml -o hostpath-storageclass.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   123  100   123    0     0    176      0 --:--:-- --:--:-- --:--:--   176
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ sed -i 's/name: example-hostpath/name: standard/' hostpath-storageclass.yaml 
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ curl -jkSL https://raw.githubusercontent.com/kubernetes-incubator/external-storage/master/docs/demo/hostpath-provisioner/claim.yaml -o hostpath-persistentvolumeclaim.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   236  100   236    0     0    256      0 --:--:-- --:--:-- --:--:--   256
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ sed -i 's%volume.beta.kubernetes.io/storage-class: "example-hostpath"%volume.beta.kubernetes.io/storage-class: "standard"%' hostpath-persistentvolumeclaim.yaml
```

Build _hostpath-provisioner_ binary and Docker image
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ cd /work/src/github.com/
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com$ git clone --depth=1 https://github.com/kubernetes-incubator/external-storage kubernetes-incubator/external-storage
Cloning into 'kubernetes-incubator/external-storage'...
remote: Counting objects: 2536, done.
remote: Compressing objects: 100% (1949/1949), done.
remote: Total 2536 (delta 634), reused 1786 (delta 465), pack-reused 0
Receiving objects: 100% (2536/2536), 4.08 MiB | 315.00 KiB/s, done.
Resolving deltas: 100% (634/634), done.
Checking connectivity... done.
Checking out files: 100% (2255/2255), done.
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner$ make hostpath-provisioner
glide install -v --strip-vcs
[INFO] Downloading dependencies. Please wait...

[INFO] Setting references.

[INFO] Removing version control data from vendor directory...

[INFO]	Removing: /work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner/vendor/github.com/kubernetes-incubator/external-storage/vendor
[ERROR]	Unable to strip vendor directories: remove /work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner/vendor/github.com/kubernetes-incubator/external-storage/vendor/golang.org/x/sys/unix: directory not empty
An Error has occurred
make: *** [hostpath-provisioner] Error 2
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner$ glide install -v --strip-vcs
fanhonglingdeMacBook-Pro:Downloads fanhongling$ rm -rf workspace/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner/vendor/github.com/kubernetes-incubator/external-storage/vendor/golang.org/x/sys/unix/...
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner$ CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o hostpath-provisioner .
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner$ ./hostpath-provisioner --help
Usage of ./hostpath-provisioner:
  -alsologtostderr
    	log to standard error as well as files
  -log_backtrace_at value
    	when logging hits line file:N, emit a stack trace
  -log_dir string
    	If non-empty, write log files in this directory
  -logtostderr
    	log to standard error instead of files
  -stderrthreshold value
    	logs at or above this threshold go to stderr
  -v value
    	log level for V logs
  -vmodule value
    	comma-separated list of pattern=N settings for file-filtered logging
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner$ make image
docker build -t hostpath-provisioner -f Dockerfile.scratch .
Sending build context to Docker daemon  44.9 MB
Step 1 : FROM scratch
 ---> 
Step 2 : COPY hostpath-provisioner /
 ---> b637b9385a2a
Removing intermediate container 64b082cbd55e
Step 3 : CMD /hostpath-provisioner
 ---> Running in 4a6baba093f5
 ---> 87568f963447
Removing intermediate container 4a6baba093f5
Successfully built 87568f963447
```

Test _hostpath-provisioner_ 
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner$ cd /work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk/
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl create -f hostpath-provisioner-pod.yaml 
pod "hostpath-provisioner" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner$ kubectl get pods
NAME                   READY     STATUS    RESTARTS   AGE
hostpath-provisioner   1/1       Running   0          2h
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl create -f hostpath-storageclass.yaml 
storageclass "standard" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl get storageclass
NAME       TYPE
standard   example.com/hostpath   
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl create -f hostpath-persistentvolumeclaim.yaml 
persistentvolumeclaim "hostpath" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl get pv,pvc
NAME                                          CAPACITY   ACCESSMODES   RECLAIMPOLICY   STATUS    CLAIM                   REASON    AGE
pv/datadir-kafka-0                            100Mi      RWO           Retain          Bound     kafka/datadir-kafka-0             8d
pv/datadir-kafka-1                            100Mi      RWO           Retain          Bound     kafka/datadir-kafka-1             8d
pv/datadir-kafka-2                            100Mi      RWO           Retain          Bound     kafka/datadir-kafka-2             8d
pv/pvc-02ef2b20-381d-11e7-b4ee-080027360e7e   1Mi        RWX           Delete          Bound     default/hostpath                  28s

NAME           STATUS    VOLUME                                     CAPACITY   ACCESSMODES   AGE
pvc/hostpath   Bound     pvc-02ef2b20-381d-11e7-b4ee-080027360e7e   1Mi        RWX           28s
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ ls /tmp/hostpath-provisioner/
pvc-02ef2b20-381d-11e7-b4ee-080027360e7e
```

### Deploying ElasticSearch

Pull _busybox_ Docker image
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/kubernetes-incubator/external-storage/docs/demo/hostpath-provisioner$ docker pull busybox
Using default tag: latest
latest: Pulling from library/busybox
Digest: sha256:32f093055929dbc23dec4d03e09dfe971f5973a9ca5cf059cbfb644c206aa83f
Status: Image is up to date for busybox:latest
```

Create _namespace_
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl create -f es-namespace.yaml 
namespace "efk" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl get ns
NAME          STATUS    AGE
default       Active    10d
efk           Active    7s
kafka         Active    8d
kube-system   Active    10d
thisisatest   Active    6d
```

Create _elasticsearch_ for _master_
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl create -f es-discovery-svc.yaml -f es-svc.yaml -f es-master.yaml 
service "elasticsearch-discovery" created
service "elasticsearch" created
deployment "es-master" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl --namespace=efk get all --selector=component=elasticsearch
NAME                            READY     STATUS    RESTARTS   AGE
po/es-master-2570037784-43wb5   1/1       Running   0          7m
po/es-master-2570037784-db0q3   1/1       Running   0          7m
po/es-master-2570037784-z0bth   1/1       Running   0          7m

NAME                          CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
svc/elasticsearch             10.123.242.101   <nodes>       9200:30198/TCP   7m
svc/elasticsearch-discovery   10.123.243.183   <none>        9300/TCP         7m

NAME               DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/es-master   3         3         3            3           7m

NAME                      DESIRED   CURRENT   READY     AGE
rs/es-master-2570037784   3         3         3         7m
```

Create _StatefulSet_ of _elasticsearch_ for _data_
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl create -f es-client.yaml -f es-data-stateful-svc.yaml -f es-data-stateful.yaml 
deployment "es-client" created
service "elasticsearch-data" created
statefulset "es-data" created
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl --namespace=efk get all --selector=role=data
NAME           READY     STATUS    RESTARTS   AGE
po/es-data-0   1/1       Running   0          3m
po/es-data-1   1/1       Running   0          1m
po/es-data-2   1/1       Running   0          1m

NAME                     CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
svc/elasticsearch-data   None         <none>        9300/TCP   3m

NAME                   DESIRED   CURRENT   AGE
statefulsets/es-data   3         3         3m
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl --namespace=efk get all --selector=role=client
NAME                            READY     STATUS    RESTARTS   AGE
po/es-client-2818086261-8kh55   1/1       Running   1          8m
po/es-client-2818086261-f46b5   1/1       Running   1          8m

NAME                CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
svc/elasticsearch   10.123.242.101   <nodes>       9200:30198/TCP   17m

NAME               DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/es-client   2         2         2            2           8m

NAME                      DESIRED   CURRENT   READY     AGE
rs/es-client-2818086261   2         2         2         8m
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ ls /tmp/hostpath-provisioner/
pvc-02ef2b20-381d-11e7-b4ee-080027360e7e  pvc-90a1be11-381f-11e7-9502-080027360e7e
pvc-90975014-381f-11e7-9502-080027360e7e  pvc-90a55e0a-381f-11e7-9502-080027360e7e
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/elasticsearch/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fpires0x2Fkubernetes-elasticsearch-cluster0x2Fmaster/ns-efk$ kubectl get pv
NAME                                       CAPACITY   ACCESSMODES   RECLAIMPOLICY   STATUS    CLAIM                   REASON    AGE
datadir-kafka-0                            100Mi      RWO           Retain          Bound     kafka/datadir-kafka-0             8d
datadir-kafka-1                            100Mi      RWO           Retain          Bound     kafka/datadir-kafka-1             8d
datadir-kafka-2                            100Mi      RWO           Retain          Bound     kafka/datadir-kafka-2             8d
pvc-02ef2b20-381d-11e7-b4ee-080027360e7e   1Mi        RWX           Delete          Bound     default/hostpath                  56m
pvc-90975014-381f-11e7-9502-080027360e7e   100Mi      RWO           Delete          Bound     efk/storage-es-data-0             37m
pvc-90a1be11-381f-11e7-9502-080027360e7e   100Mi      RWO           Delete          Bound     efk/storage-es-data-1             37m
pvc-90a55e0a-381f-11e7-9502-080027360e7e   100Mi      RWO           Delete          Bound     efk/storage-es-data-2             37m
```