apiVersion: apps/v1beta2
kind: StatefulSet
metadata:
  labels:
    app: mariadb
    component: mariadb-galera
    example.com/mysql-operator: {{.CustomResourceName}} # required by operator itself!
  name: {{.Name}} # e.g. demo-mariadb-galera
  #namespace: default
spec:
  podManagementPolicy: OrderedReady
  #podManagementPolicy: Parallel
  replicas: {{.Count}}
  selector:
    matchLabels:
      app: mariadb
      component: mariadb-galera
      example.com/mysql-operator: {{.CustomResourceName}} # required by operator itself!
  serviceName: {{.ServiceName}} # e.g. demo-mariadb-galera
  template:
    metadata:
      labels:
        app: mariadb
        component: mariadb-galera
        app-affinity: local-test-affinity
        example.com/go-to-kubernetes: {{.Name}} # required by operator itself!
        example.com/mysql-operator: {{.CustomResourceName}} # required by operator itself!
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
        - mysqld
        command:
        - docker-entrypoint.sh
        env:
        - name: MYSQL_USER
          value: {{.MysqlUser}}  
        - name: MYSQL_PASSWORD
          value: {{.MysqlPassword}}  
        #- name: MYSQL_PASSWORD_FILE
        #  value: /run/secrets/mysql-user  
        - name: MYSQL_DATABASE
          value: {{.MysqlDatabase}}  
        - name: MYSQL_ROOT_PASSWORD
          value: {{.MysqlRootPassword}}  
        #- name: MYSQL_ROOT_PASSWORD_FILE
        #  value: /run/secrets/mysql-root
        - name: POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        image: docker.io/mariadb:10.2
        imagePullPolicy: IfNotPresent
        #livenessProbe:
        #  exec:
        #    command:
        #    - mysqladmin
        #    - --password={{.MysqlRootPassword}}
        #    - ping
        #  failureThreshold: 3
        #  initialDelaySeconds: 5
        #  periodSeconds: 10
        #  successThreshold: 1
        #  timeoutSeconds: 5
        name: mariadb-galera
        ports:
        - containerPort: 3306
          name: client
          protocol: TCP
        #readinessProbe:
        #  exec:
        #    command:
        #    - mysqladmin
        #    - --password={{.MysqlRootPassword}}
        #    - ping
        #  failureThreshold: 3
        #  initialDelaySeconds: 15
        #  periodSeconds: 10
        #  successThreshold: 1
        #  timeoutSeconds: 5
        resources: {}
        securityContext:
          privileged: true
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /etc/mysql/mariadb.conf.d
          name: confd
        - mountPath: /docker-entrypoint-initdb.d
          name: initdb
        - mountPath: /var/lib/mysql
          #name: local-vol
          name: hostpath
        - mountPath: /podinfo
          name: podinfo
      dnsPolicy: ClusterFirst
      initContainers:
      - args:
        - --conf_dir=/etc/mysql/mariadb.conf.d
        - --logtostderr
        - --v=5
        command:
        - /mysql-operator
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
        - name: CLUSTER_NAME
          value: {{.ClusterName}}
        image: docker.io/tangfeixiong/mysql-operator
        #imagePullPolicy: IfNotPresent
        imagePullPolicy: Always
        name: initcnf
        resources: {}
        securityContext:
          privileged: true
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /etc/mysql/mariadb.conf.d
          name: confd
        - mountPath: /docker-entrypoint-initdb.d
          name: initdb
        - mountPath: /tmp/docker-entrypoint-initdb.d
          name: initscripts
        - mountPath: /podinfo
          name: podinfo
      #restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      #securityContext:
      #  fsGroup: 1234
      terminationGracePeriodSeconds: 30
      volumes:
      - emptyDir: {}
        name: confd
      - emptyDir: {}
        name: initdb
      - name: initscripts
        persistentVolumeClaim:
          claimName: example-local-claim
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
  #        storage: 100Mi
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
          storage: 80Mi
    