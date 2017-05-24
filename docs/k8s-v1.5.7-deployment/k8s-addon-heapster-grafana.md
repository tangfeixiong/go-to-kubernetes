Kubernetes度量和监控平台
=====================

Table of Contents
-----------------

* Heapster, InfluxDB, Grafana

Content
--------

![屏幕快照 2017-05-16 上午6.58.30.png](./屏幕快照%202017-05-16%20上午6.58.30.png)

![屏幕快照 2017-05-16 上午9.19.30.png](./屏幕快照%202017-05-16%20上午9.19.30.png)

### Heapster and Grafana (InfluxDB)

With _saltbase_
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxvf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/kube-addons/cluster-monitoring/
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/googleinfluxdb/
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/googleinfluxdb/heapster-controller-combined.yaml
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/standalone/
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/standalone/heapster-service.yaml
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/standalone/heapster-controller.yaml
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/influxdb-grafana-controller.yaml
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/grafana-service.yaml
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/heapster-service.yaml
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/heapster-controller.yaml
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/influxdb-service.yaml
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/google/
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/google/heapster-service.yaml
kubernetes/saltbase/salt/kube-addons/cluster-monitoring/google/heapster-controller.yaml
```

Pull images for v1.3.10
```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster:v1.1.0
v1.1.0: Pulling from google_containers/heapster
a01c74355a5f: Pull complete 
a3ed95caeb02: Pull complete 
7e46460a408a: Pull complete 
884fe5582a2b: Pull complete 
d2c7964e7cfd: Pull complete 
Digest: sha256:47220999ca593a35fa15e4bce6df18b6e82c8e1d7c7a00b312560eb5a007e5c4
Status: Downloaded newer image for gcr.io/google_containers/heapster:v1.1.0
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/addon-resizer:1.3
1.3: Pulling from google_containers/addon-resizer
eeee0535bf3c: Already exists 
a3ed95caeb02: Already exists 
3c24a95a48a9: Already exists 
Digest: sha256:ed051a326cede252129913c8e1435e8c7f6c49efaf5ac2643672786e7c315d33
Status: Image is up to date for gcr.io/google_containers/addon-resizer:1.3
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster_influxdb:v0.5
v0.5: Pulling from google_containers/heapster_influxdb
a3ed95caeb02: Already exists 
7059585c469e: Already exists 
782c76bb9e67: Already exists 
706514fbad74: Already exists 
a059048dac62: Already exists 
ae96ab8536d1: Already exists 
Digest: sha256:080b0559ac3b7a01fd92ca7161e4ad130bb16f812d70f71f036e5cf9fddce413
Status: Image is up to date for gcr.io/google_containers/heapster_influxdb:v0.5
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster_grafana:v2.6.0-2
v2.6.0-2: Pulling from google_containers/heapster_grafana
03e1855d4f31: Already exists 
a3ed95caeb02: Already exists 
7f1ce4d71e93: Already exists 
23d149931be4: Already exists 
2e86b9218e3a: Already exists 
db71c66d238d: Already exists 
de3678928269: Already exists 
Digest: sha256:208c98b77d4e18ad7759c0958bf87d467a3243bf75b76f1240a577002e9de277
Status: Image is up to date for gcr.io/google_containers/heapster_grafana:v2.6.0-2
```

Pull images for v1.5.7
```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster:v1.2.0
v1.2.0: Pulling from google_containers/heapster
bfc185be0245: Pull complete 
a3ed95caeb02: Pull complete 
615d924f3dd0: Pull complete 
45b80ce42529: Pull complete 
6e16b98311c7: Pull complete 
Digest: sha256:97485e7168ee127c4fd42fc248b56a50dfbd5db878335c0bd190663c0cad16bc
Status: Downloaded newer image for gcr.io/google_containers/heapster:v1.2.0
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/addon-resizer:1.7
1.7: Pulling from google_containers/addon-resizer
4b0bc1c4050b: Already exists 
4c27efe83392: Pull complete 
Digest: sha256:dcec9a5c2e20b8df19f3e9eeb87d9054a9e94e71479b935d5cfdbede9ce15895
Status: Downloaded newer image for gcr.io/google_containers/addon-resizer:1.7
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster_influxdb:v0.7
v0.7: Pulling from google_containers/heapster_influxdb
5ba4f30e5bea: Pull complete 
9d7d19c9dc56: Pull complete 
ac6ad7efd0f9: Pull complete 
e7491a747824: Pull complete 
a3ed95caeb02: Pull complete 
bc32aa1f608c: Pull complete 
2d1b799bb7c9: Pull complete 
Digest: sha256:746f16791b937af9e2d1871458296a3e0a6aad2fea9a8c8e7b70c7a0c5694928
Status: Downloaded newer image for gcr.io/google_containers/heapster_influxdb:v0.7
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster_grafana:v3.1.1
v3.1.1: Pulling from google_containers/heapster_grafana
fdd5d7827f33: Already exists 
a3ed95caeb02: Pull complete 
375c4f2dd3bf: Pull complete 
28fb221884b7: Pull complete 
e8af4667efd9: Pull complete 
90e39cbcca1c: Pull complete 
2cdad2a6e87d: Pull complete 
Digest: sha256:8183ae0d8d8867a0e66d40deb4f3f62afc7af80f2e06f0422f9f8d5f9bb17e22
Status: Downloaded newer image for gcr.io/google_containers/heapster_grafana:v3.1.1
```

Generate manifests
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/stackdocker/tomcat-docker/download/tomcat$ sudo sed '/{[%#].*[%#]}/d;s/{{ nanny_memory }}/50Mi/g;s/{{ base_metrics_cpu }}/50m/;s/{{ metrics_cpu_per_node }}/0.5/;s/{{ base_metrics_memory }}/50Mi/;s/{{ metrics_memory_per_node }}/4/;s/{{ base_eventer_memory }}/80Mi/;s/{{ eventer_memory_per_node }}/500/' /opt/kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/heapster-controller.yaml | sudo tee /etc/kubernetes/manifests/addons/heapster-controller.yaml

apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: heapster-v1.2.0.1
  namespace: kube-system
  labels:
    k8s-app: heapster
    kubernetes.io/cluster-service: "true"
    version: v1.2.0.1
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: heapster
      version: v1.2.0.1
  template:
    metadata:
      labels:
        k8s-app: heapster
        version: v1.2.0.1
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
        scheduler.alpha.kubernetes.io/tolerations: '[{"key":"CriticalAddonsOnly", "operator":"Exists"}]'
    spec:
      containers:
        - image: gcr.io/google_containers/heapster:v1.2.0
          name: heapster
          livenessProbe:
            httpGet:
              path: /healthz
              port: 8082
              scheme: HTTP
            initialDelaySeconds: 180
            timeoutSeconds: 5
          command:
            - /heapster
            - --source=kubernetes.summary_api:''
            - --sink=influxdb:http://monitoring-influxdb:8086
        - image: gcr.io/google_containers/heapster:v1.2.0
          name: eventer
          command:
            - /eventer
            - --source=kubernetes:''
            - --sink=influxdb:http://monitoring-influxdb:8086
        - image: gcr.io/google_containers/addon-resizer:1.7
          name: heapster-nanny
          resources:
            limits:
              cpu: 50m
              memory: 90Mi
            requests:
              cpu: 50m
              memory: 90Mi
          env:
            - name: MY_POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: MY_POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
          command:
            - /pod_nanny
            - --cpu=80m
            - --extra-cpu=0.5|floatm
            - --memory=100Mi
            - --extra-memory=4Mi
            - --threshold=5
            - --deployment=heapster-v1.2.0.1
            - --container=heapster
            - --poll-period=300000
            - --estimator=exponential
        - image: gcr.io/google_containers/addon-resizer:1.7
          name: eventer-nanny
          resources:
            limits:
              cpu: 50m
              memory: 90Mi
            requests:
              cpu: 50m
              memory: 90Mi
          env:
            - name: MY_POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: MY_POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
          command:
            - /pod_nanny
            - --cpu=100m
            - --extra-cpu=0m
            - --memory=100Mi
            - --extra-memory=500Ki
            - --threshold=5
            - --deployment=heapster-v1.2.0.1
            - --container=eventer
            - --poll-period=300000
            - --estimator=exponential

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/stackdocker/tomcat-docker/download/tomcat$ sudo cat /opt/kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/heapster-service.yaml | sudo tee /etc/kubernetes/manifests/addons/heapster-service.yaml
kind: Service
apiVersion: v1
metadata:
  name: heapster
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: "Heapster"
spec: 
  ports: 
    - port: 80
      targetPort: 8082
  selector: 
    k8s-app: heapster

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/stackdocker/tomcat-docker/download/tomcat$ sudo sed 's/500Mi/100Mi/g' /opt/kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/influxdb-grafana-controller.yaml | sudo tee /etc/kubernetes/manifests/addons/influxdb-grafana-controller.yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: monitoring-influxdb-grafana-v4
  namespace: kube-system
  labels: 
    k8s-app: influxGrafana
    version: v4
    kubernetes.io/cluster-service: "true"
spec: 
  replicas: 1
  selector: 
    k8s-app: influxGrafana
    version: v4
  template: 
    metadata: 
      labels: 
        k8s-app: influxGrafana
        version: v4
        kubernetes.io/cluster-service: "true"
    spec: 
      containers: 
        - image: gcr.io/google_containers/heapster_influxdb:v0.7
          name: influxdb
          resources:
            # keep request = limit to keep this container in guaranteed class
            limits:
              cpu: 100m
              memory: 100Mi
            requests:
              cpu: 100m
              memory: 100Mi
          ports: 
            - containerPort: 8083
            - containerPort: 8086
          volumeMounts:
          - name: influxdb-persistent-storage
            mountPath: /data
        - image: gcr.io/google_containers/heapster_grafana:v3.1.1
          name: grafana
          env:
          resources:
            # keep request = limit to keep this container in guaranteed class
            limits:
              cpu: 100m
              memory: 100Mi
            requests:
              cpu: 100m
              memory: 100Mi
          env:
            # This variable is required to setup templates in Grafana.
            - name: INFLUXDB_SERVICE_URL
              value: http://monitoring-influxdb:8086
              # The following env variables are required to make Grafana accessible via
              # the kubernetes api-server proxy. On production clusters, we recommend
              # removing these env variables, setup auth for grafana, and expose the grafana
              # service using a LoadBalancer or a public IP.
            - name: GF_AUTH_BASIC_ENABLED
              value: "false"
            - name: GF_AUTH_ANONYMOUS_ENABLED
              value: "true"
            - name: GF_AUTH_ANONYMOUS_ORG_ROLE
              value: Admin
            - name: GF_SERVER_ROOT_URL
              value: /api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/
          volumeMounts:
          - name: grafana-persistent-storage
            mountPath: /var
      volumes:
      - name: influxdb-persistent-storage
        emptyDir: {}
      - name: grafana-persistent-storage
        emptyDir: {}

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/stackdocker/tomcat-docker/download/tomcat$ sudo cat /opt/kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/influxdb-service.yaml | sudo tee /etc/kubernetes/manifests/addons/influxdb-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: monitoring-influxdb
  namespace: kube-system
  labels: 
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: "InfluxDB"
spec: 
  ports: 
    - name: http
      port: 8083
      targetPort: 8083
    - name: api
      port: 8086
      targetPort: 8086
  selector: 
    k8s-app: influxGrafana

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/stackdocker/tomcat-docker/download/tomcat$ sudo cat /opt/kubernetes/saltbase/salt/kube-addons/cluster-monitoring/influxdb/grafana-service.yaml | sudo tee /etc/kubernetes/manifests/addons/grafana-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: monitoring-grafana
  namespace: kube-system
  labels: 
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: "Grafana"
spec:
  # On production clusters, consider setting up auth for grafana, and
  # exposing Grafana either using a LoadBalancer or a public IP.
  # type: LoadBalancer
  ports: 
    - port: 80
      targetPort: 3000
  selector: 
    k8s-app: influxGrafana
```

Deploy
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl create -f /etc/kubernetes/manifests/addons/influxdb-grafana-controller.yaml -f /etc/kubernetes/manifests/addons/influxdb-service.yaml -f /etc/kubernetes/manifests/addons/grafana-service.yaml -f /etc/kubernetes/manifests/addons/heapster-controller.yaml -f /etc/kubernetes/manifests/addons/heapster-service.yaml
replicationcontroller "monitoring-influxdb-grafana-v4" created
service "monitoring-influxdb" created
service "monitoring-grafana" created
deployment "heapster-v1.2.0.1" created
service "heapster" created

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get all -l k8s-app=influxGrafana
NAME                                      READY     STATUS    RESTARTS   AGE
po/monitoring-influxdb-grafana-v4-kp44p   2/2       Running   0          3h

NAME                                DESIRED   CURRENT   READY     AGE
rc/monitoring-influxdb-grafana-v4   1         1         1         3h

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get all -l k8s-app=heapster
NAME                                    READY     STATUS    RESTARTS   AGE
po/heapster-v1.2.0.1-3282335156-p1l16   0/4       Pending   0          1m

NAME                       DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/heapster-v1.2.0.1   1         1         1            0           1m

NAME                              DESIRED   CURRENT   READY     AGE
rs/heapster-v1.2.0.1-3282335156   1         1         0         1m
rs/heapster-v1.2.0.1-4244670147   0         0         0         1m
```

Test
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get po -o wide -l k8s-app=influxGrafana
NAME                                   READY     STATUS    RESTARTS   AGE       IP           NODE
monitoring-influxdb-grafana-v4-kp44p   2/2       Running   2          6h        172.18.0.6   172.17.4.200
vagrant@vagrant-ubuntu-trusty-64:~$ curl 172.18.0.6:3000
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get service monitoring-influxdb
NAME                  CLUSTER-IP       EXTERNAL-IP   PORT(S)             AGE
monitoring-influxdb   10.123.245.191   <none>        8083/TCP,8086/TCP   6h
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get service monitoring-grafana
NAME                 CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
monitoring-grafana   10.123.249.35   <none>        80/TCP    6h
vagrant@vagrant-ubuntu-trusty-64:~$ curl 10.123.249.35
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width">
    <meta name="theme-color" content="#000">

    <title>Grafana</title>

		<link href='/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/public/css/fonts.min.css' rel='stylesheet' type='text/css'>

		
		  <link rel="stylesheet" href="/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/public/css/grafana.dark.min.4a9eed6d.css">
		

    <link rel="icon" type="image/png" href="/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/public/img/fav32.png">
		<base href="/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/" />

	</head>

	<body ng-cloak>
		<grafana-app class="grafana-app">

			<aside class="sidemenu-wrapper">
				<sidemenu ng-if="contextSrv.sidemenu"></sidemenu>
			</aside>

			<div class="page-alert-list">
				<div ng-repeat='alert in dashAlerts.list' class="alert-{{alert.severity}} alert">
					<button type="button" class="alert-close" ng-click="dashAlerts.clear(alert)">
						<i class="fa fa-times-circle"></i>
					</button>
					<div class="alert-title">{{alert.title}}</div>
					<div ng-bind='alert.text'></div>
				</div>
			</div>

			<div ng-view class="main-view"></div>
			<footer class="footer">
				<div class="row text-center">
					<ul>
						<li>
							<a href="http://docs.grafana.org" target="_blank">
								<i class="fa fa-file-code-o"></i>
								Docs
							</a>
						</li>
						<li>
							<a href="https://grafana.net/support/plans" target="_blank">
								<i class="fa fa-support"></i>
								Support Plans
							</a>
						</li>
						<li>
							<a href="http://grafana.org/community" target="_blank">
								<i class="fa fa-comments-o"></i>
								Community
							</a>
						</li>
						<li>
							<a href="http://grafana.org" target="_blank">Grafana</a>
							<span>v3.1.1 (commit: a4d2708)</span>
						</li>
						<li>
							
						</li>
					</ul>
				</div>
			</footer>
		</grafana-app>
  </body>

	<script>
		window.grafanaBootData = {
			user:{"isSignedIn":false,"id":0,"login":"","email":"","name":"","lightTheme":false,"orgId":1,"orgName":"Main Org.","orgRole":"Admin","isGrafanaAdmin":false,"gravatarUrl":"","timezone":"browser"},
			settings: {"allowOrgCreate":false,"appSubUrl":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana","authProxyEnabled":false,"buildInfo":{"buildstamp":1470046816,"commit":"a4d2708","env":"production","hasUpdate":false,"latestVersion":"","version":"3.1.1"},"datasources":{"-- Grafana --":{"meta":{"type":"datasource","name":"Grafana","id":"grafana","info":{"author":{"name":"","url":""},"description":"","links":null,"logos":{"small":"public/img/icn-datasource.svg","large":"public/img/icn-datasource.svg"},"screenshots":null,"version":"","updated":""},"dependencies":{"grafanaVersion":"*","plugins":[]},"includes":null,"module":"app/plugins/datasource/grafana/module","baseUrl":"public/app/plugins/datasource/grafana","annotations":false,"metrics":true,"builtIn":true,"mixed":false,"app":""},"type":"grafana"},"-- Mixed --":{"meta":{"type":"datasource","name":"Mixed datasource","id":"mixed","info":{"author":{"name":"","url":""},"description":"","links":null,"logos":{"small":"public/img/icn-datasource.svg","large":"public/img/icn-datasource.svg"},"screenshots":null,"version":"","updated":""},"dependencies":{"grafanaVersion":"*","plugins":[]},"includes":null,"module":"app/plugins/datasource/mixed/module","baseUrl":"public/app/plugins/datasource/mixed","annotations":false,"metrics":true,"builtIn":true,"mixed":true,"app":""},"type":"mixed"},"influxdb-datasource":{"database":"k8s","jsonData":{},"meta":{"type":"datasource","name":"InfluxDB","id":"influxdb","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/datasource/influxdb/img/influxdb_logo.svg","large":"public/app/plugins/datasource/influxdb/img/influxdb_logo.svg"},"screenshots":null,"version":"","updated":""},"dependencies":{"grafanaVersion":"*","plugins":[]},"includes":null,"module":"app/plugins/datasource/influxdb/module","baseUrl":"public/app/plugins/datasource/influxdb","annotations":true,"metrics":true,"builtIn":false,"mixed":false,"app":""},"name":"influxdb-datasource","type":"influxdb","url":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/api/datasources/proxy/1"}},"defaultDatasource":"influxdb-datasource","panels":{"dashlist":{"baseUrl":"public/app/plugins/panel/dashlist","id":"dashlist","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/dashlist/img/icn-dashlist-panel.svg","large":"public/app/plugins/panel/dashlist/img/icn-dashlist-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/dashlist/module","name":"Dashboard list"},"graph":{"baseUrl":"public/app/plugins/panel/graph","id":"graph","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/graph/img/icn-graph-panel.svg","large":"public/app/plugins/panel/graph/img/icn-graph-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/graph/module","name":"Graph"},"pluginlist":{"baseUrl":"public/app/plugins/panel/pluginlist","id":"pluginlist","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/pluginlist/img/icn-dashlist-panel.svg","large":"public/app/plugins/panel/pluginlist/img/icn-dashlist-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/pluginlist/module","name":"Plugin list"},"singlestat":{"baseUrl":"public/app/plugins/panel/singlestat","id":"singlestat","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/singlestat/img/icn-singlestat-panel.svg","large":"public/app/plugins/panel/singlestat/img/icn-singlestat-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/singlestat/module","name":"Singlestat"},"table":{"baseUrl":"public/app/plugins/panel/table","id":"table","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/table/img/icn-table-panel.svg","large":"public/app/plugins/panel/table/img/icn-table-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/table/module","name":"Table"},"text":{"baseUrl":"public/app/plugins/panel/text","id":"text","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/text/img/icn-text-panel.svg","large":"public/app/plugins/panel/text/img/icn-text-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/text/module","name":"Text"}}},
			mainNavLinks: [{"text":"Dashboards","icon":"icon-gf icon-gf-dashboard","url":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/","children":[{"text":"Home","url":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/"},{"text":"Playlists","url":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/playlists"},{"text":"Snapshots","url":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/dashboard/snapshots"},{"divider":true},{"text":"New","icon":"fa fa-plus","url":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/dashboard/new"},{"text":"Import","icon":"fa fa-download","url":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/dashboard/new/?editview=import"}]},{"text":"Data Sources","icon":"icon-gf icon-gf-datasources","url":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/datasources"},{"text":"Plugins","icon":"icon-gf icon-gf-apps","url":"/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/plugins"}]
		};
	</script>

	<script src="/api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/public/app/boot.1649cbaa.js"></script>

	

	
</html>
```

Expose
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system expose service monitoring-grafana --name=monitoring-grafana-ex --target-port=3000 --type=NodePort
service "monitoring-grafana-ex" exposed
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get service monitoring-grafana-ex
NAME                    CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
monitoring-grafana-ex   10.123.240.77   <nodes>       80:32567/TCP   1m
```

### v1.3.10

From https://github.com/kubernetes/kubernetes/v1.3.10/cluster/addons/cluster-monitoring/influxdb
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/stackdocker/tomcat-docker/download/tomcat$ curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/v1.3.10/cluster/addons/cluster-monitoring/influxdb/heapster-controller.yaml | sudo sed '/{[%#].*[%#]}/d;s/{{ nanny_memory }}/90Mi/g;s/{{ base_metrics_cpu }}/80m/;s/{{ metrics_cpu }}/80m/;s/{{ metrics_cpu_per_node }}/0.5/;s/{{ base_metrics_memory }}/140Mi/;s/{{ metrics_memory }}/140Mi/;s/{{ metrics_memory_per_node }}/4/;s/{{ base_eventer_memory }}/190Mi/;s/{{ eventer_memory }}/190Mi/;s/{{ eventer_memory_per_node }}/500/' | sudo tee /etc/kubernetes/manifests/addons/heapster-controller.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  4279  100  4279    0     0   5142      0 --:--:-- --:--:-- --:--:--  5143

apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: heapster-v1.1.0
  namespace: kube-system
  labels:
    k8s-app: heapster
    kubernetes.io/cluster-service: "true"
    version: v1.1.0
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: heapster
      version: v1.1.0
  template:
    metadata:
      labels:
        k8s-app: heapster
        version: v1.1.0
    spec:
      containers:
        - image: gcr.io/google_containers/heapster:v1.1.0
          name: heapster
          resources:
            # keep request = limit to keep this container in guaranteed class
            limits:
              cpu: 80m
              memory: 140Mi
            requests:
              cpu: 80m
              memory: 140Mi
          command:
            - /heapster
            - --source=kubernetes.summary_api:''
            - --sink=influxdb:http://monitoring-influxdb:8086
        - image: gcr.io/google_containers/heapster:v1.1.0
          name: eventer
          resources:
            # keep request = limit to keep this container in guaranteed class
            limits:
              cpu: 100m
              memory: 190Mi
            requests:
              cpu: 100m
              memory: 190Mi
          command:
            - /eventer
            - --source=kubernetes:''
            - --sink=influxdb:http://monitoring-influxdb:8086
        - image: gcr.io/google_containers/addon-resizer:1.3
          name: heapster-nanny
          resources:
            limits:
              cpu: 50m
              memory: 90Mi
            requests:
              cpu: 50m
              memory: 90Mi
          env:
            - name: MY_POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: MY_POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
          command:
            - /pod_nanny
            - --cpu=80m
            - --extra-cpu=0.5m
            - --memory=140Mi
            - --extra-memory=4Mi
            - --threshold=5
            - --deployment=heapster-v1.1.0
            - --container=heapster
            - --poll-period=300000
            - --estimator=exponential
        - image: gcr.io/google_containers/addon-resizer:1.3
          name: eventer-nanny
          resources:
            limits:
              cpu: 50m
              memory: 90Mi
            requests:
              cpu: 50m
              memory: 90Mi
          env:
            - name: MY_POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: MY_POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
          command:
            - /pod_nanny
            - --cpu=100m
            - --extra-cpu=0m
            - --memory=190Mi
            - --extra-memory=500Ki
            - --threshold=5
            - --deployment=heapster-v1.1.0
            - --container=eventer
            - --poll-period=300000
            - --estimator=exponential

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system create -f /etc/kubernetes/manifests/addons/heapster-controller.yaml 
deployment "heapster-v1.1.0" created
```

### https://github.com/coreos/coreos-kubernetes

From https://github.com/coreos/coreos-kubernetes/blob/v0.8.2/single-node/user-data
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo vi /etc/kubernetes/manifests/addons/coreos-kubernetes.heapster-de.json
{
  "apiVersion": "extensions/v1beta1",
  "kind": "Deployment",
  "metadata": {
    "labels": {
      "k8s-app": "heapster",
      "kubernetes.io/cluster-service": "true",
      "version": "v1.1.0"
    },
    "name": "heapster-v1.1.0",
    "namespace": "kube-system"
  },
  "spec": {
    "replicas": 1,
    "selector": {
      "matchLabels": {
        "k8s-app": "heapster",
        "version": "v1.1.0"
      }
    },
    "template": {
      "metadata": {
        "labels": {
          "k8s-app": "heapster",
          "version": "v1.1.0"
        }
      },
      "spec": {
        "containers": [
          {
            "command": [
              "/heapster",
              "--source=kubernetes.summary_api:''"
            ],
            "image": "gcr.io/google_containers/heapster:v1.1.0",
            "name": "heapster",
            "resources": {
              "limits": {
                "cpu": "100m",
                "memory": "200Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "200Mi"
              }
            }
          },
          {
            "command": [
              "/pod_nanny",
              "--cpu=100m",
              "--extra-cpu=0.5m",
              "--memory=200Mi",
              "--extra-memory=4Mi",
              "--threshold=5",
              "--deployment=heapster-v1.1.0",
              "--container=heapster",
              "--poll-period=300000",
              "--estimator=exponential"
            ],
            "env": [
              {
                "name": "MY_POD_NAME",
                "valueFrom": {
                  "fieldRef": {
                    "fieldPath": "metadata.name"
                  }
                }
              },
              {
                "name": "MY_POD_NAMESPACE",
                "valueFrom": {
                  "fieldRef": {
                    "fieldPath": "metadata.namespace"
                  }
                }
              }
            ],
            "image": "gcr.io/google_containers/addon-resizer:1.3",
            "name": "heapster-nanny",
            "resources": {
              "limits": {
                "cpu": "50m",
                "memory": "100Mi"
              },
              "requests": {
                "cpu": "50m",
                "memory": "100Mi"
              }
            }
          }
        ]
      }
    }
  }
}
vagrant@vagrant-ubuntu-trusty-64:~$ kubectl --namespace=kube-system create -f /etc/kubernetes/manifests/addons/coreos-kubernetes.heapster-de.json 
deployment "heapster-v1.1.0" created

```

Form https://github.com/coreos/coreos-kubernetes/blob/v0.7.1/single-node/user-data
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo vi /etc/kubernetes/manifests/addons/coreos-kubernetes.heapster-dc.json
{
  "apiVersion": "extensions/v1beta1",
  "kind": "Deployment",
  "metadata": {
    "labels": {
      "k8s-app": "heapster",
      "kubernetes.io/cluster-service": "true",
      "version": "v1.0.2"
    },
    "name": "heapster-v1.0.2",
    "namespace": "kube-system"
  },
  "spec": {
    "replicas": 1,
    "selector": {
      "matchLabels": {
        "k8s-app": "heapster",
        "version": "v1.0.2"
      }
    },
    "template": {
      "metadata": {
        "labels": {
          "k8s-app": "heapster",
          "version": "v1.0.2"
        }
      },
      "spec": {
        "containers": [
          {
            "command": [
              "/heapster",
              "--source=kubernetes.summary_api:''",
              "--metric_resolution=60s"
            ],
            "image": "gcr.io/google_containers/heapster:v1.0.2",
            "name": "heapster",
            "resources": {
              "limits": {
                "cpu": "100m",
                "memory": "250Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "250Mi"
              }
            }
          },
          {
            "command": [
              "/pod_nanny",
              "--cpu=100m",
              "--extra-cpu=0m",
              "--memory=250Mi",
              "--extra-memory=4Mi",
              "--threshold=5",
              "--deployment=heapster-v1.0.2",
              "--container=heapster",
              "--poll-period=300000"
            ],
            "env": [
              {
                "name": "MY_POD_NAME",
                "valueFrom": {
                  "fieldRef": {
                    "fieldPath": "metadata.name"
                  }
                }
              },
              {
                "name": "MY_POD_NAMESPACE",
                "valueFrom": {
                  "fieldRef": {
                    "fieldPath": "metadata.namespace"
                  }
                }
              }
            ],
            "image": "gcr.io/google_containers/addon-resizer:1.0",
            "name": "heapster-nanny",
            "resources": {
              "limits": {
                "cpu": "50m",
                "memory": "100Mi"
              },
              "requests": {
                "cpu": "50m",
                "memory": "100Mi"
              }
            }
          }
        ]
      }
    }
  }
}
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster:v1.0.2
v1.0.2: Pulling from google_containers/heapster
56b26193e0fa: Pull complete 
a3ed95caeb02: Pull complete 
f215c8b87666: Pull complete 
92bdaa3c406e: Pull complete 
715b61202b02: Pull complete 
Digest: sha256:0346e9a67f5168763de12dc611e646ad547302e6639be47104ff4728dac62a3f
Status: Downloaded newer image for gcr.io/google_containers/heapster:v1.0.2
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/addon-resizer:1.0
1.0: Pulling from google_containers/addon-resizer
eeee0535bf3c: Already exists 
a3ed95caeb02: Pull complete 
b1834c5ac950: Pull complete 
Digest: sha256:e77acf80697a70386c04ae3ab494a7b13917cb30de2326dcf1a10a5118eddabe
Status: Downloaded newer image for gcr.io/google_containers/addon-resizer:1.0
vagrant@vagrant-ubuntu-trusty-64:~$ kubectl --namespace=kube-system create -f /etc/kubernetes/manifests/addons/coreos-kubernetes.heapster-dc.json 
deployment "heapster-v1.0.2" created
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster_influxdb:v0.5
v0.5: Pulling from google_containers/heapster_influxdb
a3ed95caeb02: Pull complete 
7059585c469e: Pull complete 
782c76bb9e67: Pull complete 
706514fbad74: Pull complete 
a059048dac62: Pull complete 
ae96ab8536d1: Pull complete 
Digest: sha256:080b0559ac3b7a01fd92ca7161e4ad130bb16f812d70f71f036e5cf9fddce413
Status: Downloaded newer image for gcr.io/google_containers/heapster_influxdb:v0.5
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster_grafana:v2.6.0-2
v2.6.0-2: Pulling from google_containers/heapster_grafana
03e1855d4f31: Already exists 
a3ed95caeb02: Already exists 
7f1ce4d71e93: Already exists 
23d149931be4: Already exists 
2e86b9218e3a: Already exists 
db71c66d238d: Already exists 
de3678928269: Already exists 
Digest: sha256:208c98b77d4e18ad7759c0958bf87d467a3243bf75b76f1240a577002e9de277
Status: Image is up to date for gcr.io/google_containers/heapster_grafana:v2.6.0-2

```

From https://github.com/coreos/coreos-kubernetes/blob/v0.6.0/single-node/user-data
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo vi /etc/kubernetes/manifests/addons/coreos-kubernetes.heapster-rc.json
{
  "apiVersion": "v1",
  "kind": "ReplicationController",
  "metadata": {
    "labels": {
      "k8s-app": "heapster",
      "kubernetes.io/cluster-service": "true"
    },
    "name": "heapster-v1.0.2",
    "namespace": "kube-system"
  },
  "spec": {
    "replicas": 1,
    "selector": {
      "k8s-app": "heapster"
    },
    "template": {
      "metadata": {
        "labels": {
          "k8s-app": "heapster",
          "kubernetes.io/cluster-service": "true"
        }
      },
      "spec": {
        "containers": [
          {
            "command": [
              "/heapster",
              "--source=kubernetes.summary_api:''",
              "--metric_resolution=60s"
            ],
            "image": "gcr.io/google_containers/heapster:v1.0.2",
            "name": "heapster",
            "resources": {
              "limits": {
                "cpu": "100m",
                "memory": "208Mi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "208Mi"
              }
            }
          }
        ]
      }
    }
  }
}
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system create -f /etc/kubernetes/manifests/addons/coreos-kubernetes.heapster-rc.json 
replicationcontroller "heapster-v1.0.2" created

vagrant@vagrant-ubuntu-trusty-64:~$ curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.2/cluster/addons/cluster-monitoring/influxdb/influxdb-grafana-controller.yaml | sudo sed 's/value: \/api\/v1\/proxy\/namespaces\/kube-system\/services\/monitoring-grafana\//value: \//' | sudo tee /etc/kubernetes/manifests/addons/influxdb-grafana-controller.yaml 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2423  100  2423    0     0   2315      0  0:00:01  0:00:01 --:--:--  2316
apiVersion: v1
kind: ReplicationController
metadata:
  name: monitoring-influxdb-grafana-v3
  namespace: kube-system
  labels: 
    k8s-app: influxGrafana
    version: v3
    kubernetes.io/cluster-service: "true"
spec: 
  replicas: 1
  selector: 
    k8s-app: influxGrafana
    version: v3
  template: 
    metadata: 
      labels: 
        k8s-app: influxGrafana
        version: v3
        kubernetes.io/cluster-service: "true"
    spec: 
      containers: 
        - image: gcr.io/google_containers/heapster_influxdb:v0.5
          name: influxdb
          resources:
            # keep request = limit to keep this container in guaranteed class
            limits:
              cpu: 100m
              memory: 500Mi
            requests:
              cpu: 100m
              memory: 500Mi
          ports: 
            - containerPort: 8083
            - containerPort: 8086
          volumeMounts:
          - name: influxdb-persistent-storage
            mountPath: /data
        - image: gcr.io/google_containers/heapster_grafana:v2.6.0-2
          name: grafana
          env:
          resources:
            # keep request = limit to keep this container in guaranteed class
            limits:
              cpu: 100m
              memory: 100Mi
            requests:
              cpu: 100m
              memory: 100Mi
          env:
            # This variable is required to setup templates in Grafana.
            - name: INFLUXDB_SERVICE_URL
              value: http://monitoring-influxdb:8086
              # The following env variables are required to make Grafana accessible via
              # the kubernetes api-server proxy. On production clusters, we recommend
              # removing these env variables, setup auth for grafana, and expose the grafana
              # service using a LoadBalancer or a public IP.
            - name: GF_AUTH_BASIC_ENABLED
              value: "false"
            - name: GF_AUTH_ANONYMOUS_ENABLED
              value: "true"
            - name: GF_AUTH_ANONYMOUS_ORG_ROLE
              value: Admin
            - name: GF_SERVER_ROOT_URL
              value: /
          volumeMounts:
          - name: grafana-persistent-storage
            mountPath: /var
      volumes:
      - name: influxdb-persistent-storage
        emptyDir: {}
      - name: grafana-persistent-storage
        emptyDir: {}

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system create -f /etc/kubernetes/manifests/addons/influxdb-grafana-controller.yaml 
replicationcontroller "monitoring-influxdb-grafana-v3" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get rc monitoring-influxdb-grafana-v3
NAME                             DESIRED   CURRENT   READY     AGE
monitoring-influxdb-grafana-v3   1         1         1         47s
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get pods -l k8s-app=influxGrafana
NAME                                   READY     STATUS    RESTARTS   AGE
monitoring-influxdb-grafana-v3-l99s0   2/2       Running   0          1m
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system expose service monitoring-grafana --name=monitoring-grafana-ex --target-port=3000 --type=NodePort
service "monitoring-grafana-ex" exposed
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get svc monitoring-grafana-ex
NAME                    CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
monitoring-grafana-ex   10.123.244.188   <nodes>       80:31346/TCP   7s

vagrant@vagrant-ubuntu-trusty-64:~$ curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.1/cluster/addons/cluster-monitoring/influxdb/influxdb-grafana-controller.yaml | sudo tee /etc/kubernetes/manifests/addons/influxdb-grafana-controller.yaml 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2196  100  2196    0     0    378      0  0:00:05  0:00:05 --:--:--   505
apiVersion: v1
kind: ReplicationController
metadata:
  name: monitoring-influxdb-grafana-v2
  namespace: kube-system
  labels: 
    k8s-app: influxGrafana
    version: v2
    kubernetes.io/cluster-service: "true"
spec: 
  replicas: 1
  selector: 
    k8s-app: influxGrafana
    version: v2
  template: 
    metadata: 
      labels: 
        k8s-app: influxGrafana
        version: v2
        kubernetes.io/cluster-service: "true"
    spec: 
      containers: 
        - image: gcr.io/google_containers/heapster_influxdb:v0.4
          name: influxdb
          resources:
            limits:
              cpu: 100m
              memory: 200Mi
          ports: 
            - containerPort: 8083
              hostPort: 8083
            - containerPort: 8086
              hostPort: 8086
          volumeMounts:
          - name: influxdb-persistent-storage
            mountPath: /data
        - image: beta.gcr.io/google_containers/heapster_grafana:v2.1.1
          name: grafana
          env:
          resources:
            limits:
              cpu: 100m
              memory: 100Mi
          env:
            # This variable is required to setup templates in Grafana.
            - name: INFLUXDB_SERVICE_URL
              value: http://monitoring-influxdb:8086
              # The following env variables are required to make Grafana accessible via
              # the kubernetes api-server proxy. On production clusters, we recommend
              # removing these env variables, setup auth for grafana, and expose the grafana
              # service using a LoadBalancer or a public IP.
            - name: GF_AUTH_BASIC_ENABLED
              value: "false"
            - name: GF_AUTH_ANONYMOUS_ENABLED
              value: "true"
            - name: GF_AUTH_ANONYMOUS_ORG_ROLE
              value: Admin
            - name: GF_SERVER_ROOT_URL
              value: /api/v1/proxy/namespaces/kube-system/services/monitoring-grafana/
          volumeMounts:
          - name: grafana-persistent-storage
            mountPath: /var
              
      volumes:
      - name: influxdb-persistent-storage
        emptyDir: {}
      - name: grafana-persistent-storage
        emptyDir: {}

vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/heapster_influxdb:v0.4
v0.4: Pulling from google_containers/heapster_influxdb
a3ed95caeb02: Pull complete 
e30706b9b4ff: Pull complete 
82286846aa71: Pull complete 
d433c42f6001: Pull complete 
f1b0b5440e07: Pull complete 
3a0c8989a1d6: Pull complete 
333791cb73c4: Pull complete 
Digest: sha256:4cdc84a811794b4df7bbadf38ddedd8d7df8c84f909c53b65f393b0d1aefb7a2
Status: Downloaded newer image for gcr.io/google_containers/heapster_influxdb:v0.4
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull beta.gcr.io/google_containers/heapster_grafana:v2.1.1
v2.1.1: Pulling from google_containers/heapster_grafana
70c964415e86: Pull complete 
a3ed95caeb02: Pull complete 
3aadf9695c00: Pull complete 
1e6e82e7ad96: Pull complete 
465ddd43a582: Pull complete 
c20a85113271: Pull complete 
726569365009: Pull complete 
4a85898a72ac: Pull complete 
9791abe2a3f3: Pull complete 
Digest: sha256:87a4a16b08ee0ae474cb528b6aa2fa1cf78a79428d247bc215f0031e326b8f9a
Status: Downloaded newer image for beta.gcr.io/google_containers/heapster_grafana:v2.1.1

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system create -f /etc/kubernetes/manifests/addons/influxdb-grafana-controller.yaml 
replicationcontroller "monitoring-influxdb-grafana-v2" created

```

### Links

[Metrics](https://github.com/kubernetes/heapster/blob/master/docs/storage-schema.md#metrics)

[Issue](https://github.com/kubernetes/heapster/issues/1316)