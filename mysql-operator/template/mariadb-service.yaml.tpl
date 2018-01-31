apiVersion: v1
kind: Service
metadata:
  labels:
    app: mariadb
    component: mariadb-galera
    example.com/mysql-operator: {{.CustomResourceName}} # operator reserved!
  name: {{.Name}} # e.g. demo-mariadb-galera
  #namespace: default
spec:
  #clusterIP: None
  ports:
  - name: client
    port: 3306
    protocol: TCP
    targetPort: 3306
  - name: traffic
    port: 4567
    protocol: TCP
    targetPort: 4567
  - name: trafficu
    port: 4567
    protocol: UDP
    targetPort: 4567
  - name: ist
    port: 4568
    protocol: TCP
    targetPort: 4568
  - name: sst
    port: 4444
    protocol: TCP
    targetPort: 4444
  selector:
    app: mariadb
    component: mariadb-galera
    example.com/mysql-operator: {{.CustomResourceName}} # required by operator itself!
  #sessionAffinity: None
  #sessionAffinity: ClientIP
  #type: ClusterIP