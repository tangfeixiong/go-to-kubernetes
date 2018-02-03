# inspired by https://github.com/rabbitmq/rabbitmq-peer-discovery-k8s/blob/master/examples/k8s_statefulsets/rabbitmq_statefulsets.yaml
apiVersion: apps/v1beta2
kind: StatefulSet
metadata:
  labels:
    app: rabbitmq
    example.com/rabbitmq-operator: "{{ .CustomResourceName }}" # required by operator itself!
    github.com/go-to-kubernetes: rabbitmq-operator
  name: "{{ .Name }}" # e.g. demo-rabbitmq-galera
  #namespace: default
spec:
  podManagementPolicy: OrderedReady
  #podManagementPolicy: Parallel
  replicas: {{ .Count }}
  selector:
    matchLabels:
      app: rabbitmq
      example.com/rabbitmq-operator: "{{ .CustomResourceName }}" # required by operator itself!
      github.com/go-to-kubernetes: rabbitmq-operator
  serviceName: "{{ .ServiceName }}" # e.g. demo-rabbitmq-galera
  template:
    metadata:
      labels:
        app: rabbitmq
        app-affinity: local-test-affinity
        example.com/rabbitmq-operator: "{{ .CustomResourceName }}" # required by operator itself!
        github.com/go-to-kubernetes: rabbitmq-operator
        rabbitmq-operator: "{{ .Name }}" # required by operator itself!
    spec:
      #affinity:
      #  podAffinity:
      #    requiredDuringSchedulingIgnoredDuringExecution:
      #    - labelSelector:
      #        matchExpressions:
      #        - key: app-affinity
      #          operator: In
      #          values:
      #          - local-test-affinity
      #      topologyKey: kubernetes.io/hostname
      #affinity:
      #  podAntiAffinity:
      #    requiredDuringSchedulingIgnoredDuringExecution:
      #    - labelSelector:
      #        matchExpressions:
      #        - key: app-affinity
      #          operator: In
      #          values:
      #          - local-test-anti-affinity
      #      topologyKey: kubernetes.io/hostname
      containers:
      - args:
        - rabbitmq-server
        command:
        - docker-entrypoint.sh
        env:
        - name: MY_POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        - name: MY_POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: RABBITMQ_USE_LONGNAME
          value: "true"
        - name: RABBITMQ_NODENAME
          value: "rabbit@$(MY_POD_IP)"
        - name: K8S_SERVICE_NAME
          value: "{{ .ServiceName }}"
        - name: RABBITMQ_ERLANG_COOKIE
          value: "mycookie" 
        image: docker.io/rabbitmq:3.7
        imagePullPolicy: IfNotPresent
        #livenessProbe:
        #  exec:
        #    command:
        #    - rabbitmqctl
        #    - status
        #  failureThreshold: 6
        #  initialDelaySeconds: 30
        #  periodSeconds: 10
        #  successThreshold: 1
        #  timeoutSeconds: 10
        name: rabbitmq
        ports:
        - containerPort: 5672
          name: amqp
          protocol: TCP
        - containerPort: 15672
          name: http
          protocol: TCP
        - containerPort: 4369
          name: epmd
        - containerPort: 25672
          name: dist
        #readinessProbe:
        #  exec:
        #    command:
        #    - rabbitmqctl
        #    - status
        #  failureThreshold: 3
        #  initialDelaySeconds: 10
        #  periodSeconds: 3
        #  successThreshold: 1
        #  timeoutSeconds: 10
        resources: {}
        securityContext: {}
        #serviceAccountName: rabbitmq
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /etc/rabbitmq
          name: conf
        - mountPath: /var/lib/rabbitmq
          #name: local-vol
          name: hostpath
      dnsPolicy: ClusterFirst
      initContainers:
      - args:
        - --conf_dir=/etc/rabbitmq
        - --logtostderr
        - --v=5
        command:
        - /rabbitmq-operator
        - init
        env:
        - name: MY_NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        - name: MY_POD_NAME  # important, required to discover StatefulSet name by operator itself!
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: MY_POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: MY_POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        #- name: MY_POD_SERVICE_ACCOUNT
        #  valueFrom:
        #    fieldRef:
        #      fieldPath: spec.serviceAccountName
        image: docker.io/tangfeixiong/rabbitmq-operator
        #imagePullPolicy: IfNotPresent
        imagePullPolicy: Always
        name: initcnf
        resources: {}
        securityContext: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /etc/rabbitmq
          name: conf
        - mountPath: /podinfo
          name: podinfo
      #restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      #securityContext:
      #  fsGroup: 1234
      serviceAccountName: rabbitmq
      terminationGracePeriodSeconds: 30
      volumes:
      - emptyDir: {}
        name: conf
      - downwardAPI:
          defaultMode: 420
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.annotations
            path: annotations
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.labels
            path: labels
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
            path: name
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
            path: namespace
        name: podinfo
  updateStrategy:
    #type: RollingUpdate
    type: OnDelete
  volumeClaimTemplates:
  #- metadata:
  #    name: local-vol
  #  spec:
  #    accessModes: [ "ReadWriteOnce" ]
  #    storageClassName: "local-storage"
  #    resources:
  #      requests:
  #        storage: 50Mi
  - metadata:
      name: hostpath
      #annotations:
      #  volume.beta.kubernetes.io/storage-class: "example-hostpath"
    spec:
      accessModes:
        - ReadWriteOnce
        #- ReadOnlyMany
        #- ReadWriteMany
      storageClassName: "example-hostpath"
      resources:
        requests:
          storage: 30Mi
    