Kubernetes日志平台
=================

Table of Contents
-----------------

* EFK for K8s v1.3.10
* ElasticSearch, fluentd, Kibana for K8s v1.5.7

Content
--------

### V1.3.10

![kibana ui](./屏幕快照%202017-05-14%20上午8.13.23.png)

Downloaded
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo mkdir -p /etc/kubernetes/manifests/addons

vagrant@vagrant-ubuntu-trusty-64:~$ sudo curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/v1.3.10/cluster/addons/fluentd-elasticsearch/es-controller.yaml -o /etc/kubernetes/manifests/addons/es-controller.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1077  100  1077    0     0   1215      0 --:--:-- --:--:-- --:--:--  1215
vagrant@vagrant-ubuntu-trusty-64:~$ sudo curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/v1.3.10/cluster/addons/fluentd-elasticsearch/es-service.yaml -o /etc/kubernetes/manifests/addons/es-service.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   335  100   335    0     0    377      0 --:--:-- --:--:-- --:--:--   377
vagrant@vagrant-ubuntu-trusty-64:~$ sudo curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/v1.3.10/cluster/addons/fluentd-elasticsearch/kibana-controller.yaml -o /etc/kubernetes/manifests/addons/kibana-controller.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   893  100   893    0     0    908      0 --:--:-- --:--:-- --:--:--   908
vagrant@vagrant-ubuntu-trusty-64:~$ sudo curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/v1.3.10/cluster/addons/fluentd-elasticsearch/kibana-service.yaml -o /etc/kubernetes/manifests/addons/kibana-service.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   307  100   307    0     0    284      0  0:00:01  0:00:01 --:--:--   285
vagrant@vagrant-ubuntu-trusty-64:~$ sudo curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/v1.3.10/cluster/saltbase/salt/fluentd-es/fluentd-es.yaml -o /etc/kubernetes/manifests/addons/fluented-es.yaml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   700  100   700    0     0    214      0  0:00:03  0:00:03 --:--:--   214
```

Pull
```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/elasticsearch:1.9
1.9: Pulling from google_containers/elasticsearch
51f5c6a04d83: Already exists 
a3ed95caeb02: Pull complete 
7004cfc6e122: Already exists 
aee1f2b2873f: Already exists 
1f7298c48e9b: Already exists 
8ddf451692fc: Pull complete 
e9a817caf28d: Pull complete 
5e5b1cc83d21: Pull complete 
571642fc63a4: Pull complete 
e3f7313bae7e: Pull complete 
12b21c24fe75: Pull complete 
2a0636ccd445: Pull complete 
3cd011c3a971: Pull complete 
Digest: sha256:d635d15e8b78645829120eb5083b6a30f513c41f070ba786b974e060771d197d
Status: Downloaded newer image for gcr.io/google_containers/elasticsearch:1.9
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/fluentd-elasticsearch:1.17
1.17: Pulling from google_containers/fluentd-elasticsearch
8387d9ff0016: Pull complete 
3b52deaaf0ed: Pull complete 
4bd501fad6de: Pull complete 
a3ed95caeb02: Pull complete 
88e724ecff68: Pull complete 
2d53af608a8a: Pull complete 
e0a40554da95: Pull complete 
e10d578debb2: Pull complete 
688e97373c86: Pull complete 
Digest: sha256:fb140aa1756f84a3d3f2bf8d27c3da38b0eab890671d7e134dde95c69ef41dd8
Status: Downloaded newer image for gcr.io/google_containers/fluentd-elasticsearch:1.17
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/kibana:1.3
1.3: Pulling from google_containers/kibana
bb548df4bb25: Pull complete 
a3ed95caeb02: Pull complete 
db7807720e29: Pull complete 
26fd388f3bd9: Pull complete 
bdb4876f5972: Pull complete 
a3d84814df51: Pull complete 
a8579e64f38f: Pull complete 
b10b0415a168: Pull complete 
Digest: sha256:4fbe92e0da989b54070a3a2dd93be7cf5ef049bbfae6b98cde15ac3b99426f12
Status: Downloaded newer image for gcr.io/google_containers/kibana:1.3
```

deploy
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl create -f /etc/kubernetes/manifests/addons/
replicationcontroller "elasticsearch-logging-v1" created
service "elasticsearch-logging" created
pod "fluentd-elasticsearch" created
replicationcontroller "kibana-logging-v1" created
service "kibana-logging" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get all -l version=v1 -o wide
NAME                                READY     STATUS    RESTARTS   AGE       IP            NODE
po/elasticsearch-logging-v1-10flh   1/1       Running   0          2m        172.18.0.15   172.17.4.200
po/elasticsearch-logging-v1-vn881   1/1       Running   0          2m        172.18.0.16   172.17.4.200
po/kibana-logging-v1-8gwfq          1/1       Running   0          2m        172.18.0.17   172.17.4.200

NAME                          DESIRED   CURRENT   READY     AGE       CONTAINER(S)            IMAGE(S)                                     SELECTOR
rc/elasticsearch-logging-v1   2         2         2         2m        elasticsearch-logging   gcr.io/google_containers/elasticsearch:1.9   k8s-app=elasticsearch-logging,version=v1
rc/kibana-logging-v1          1         1         1         2m        kibana-logging          gcr.io/google_containers/kibana:1.3          k8s-app=kibana-logging,version=v1
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get svc -l k8s-app=elasticsearch-logging
NAME                    CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
elasticsearch-logging   10.123.243.22   <none>        9200/TCP   4m
```

Test with [ES cluster APIs](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster.html)
```
vagrant@vagrant-ubuntu-trusty-64:~$ curl http://10.123.243.22:9200/_cluster/health;echo
{"cluster_name":"kubernetes-logging","status":"yellow","timed_out":false,"number_of_nodes":2,"number_of_data_nodes":2,"active_primary_shards":5,"active_shards":5,"relocating_shards":0,"initializing_shards":0,"unassigned_shards":5,"number_of_pending_tasks":0}
vagrant@vagrant-ubuntu-trusty-64:~$ curl http://10.123.243.22:9200/_cluster/stats?human\&pretty
{
  "timestamp" : 1494773695393,
  "cluster_name" : "kubernetes-logging",
  "status" : "yellow",
  "indices" : {
    "count" : 1,
    "shards" : {
      "total" : 5,
      "primaries" : 5,
      "replication" : 0.0,
      "index" : {
        "shards" : {
          "min" : 5,
          "max" : 5,
          "avg" : 5.0
        },
        "primaries" : {
          "min" : 5,
          "max" : 5,
          "avg" : 5.0
        },
        "replication" : {
          "min" : 0.0,
          "max" : 0.0,
          "avg" : 0.0
        }
      }
    },
    "docs" : {
      "count" : 32788,
      "deleted" : 0
    },
    "store" : {
      "size" : "8.3mb",
      "size_in_bytes" : 8736274,
      "throttle_time" : "486.7ms",
      "throttle_time_in_millis" : 486
    },
    "fielddata" : {
      "memory_size" : "0b",
      "memory_size_in_bytes" : 0,
      "evictions" : 0
    },
    "filter_cache" : {
      "memory_size" : "0b",
      "memory_size_in_bytes" : 0,
      "evictions" : 0
    },
    "id_cache" : {
      "memory_size" : "0b",
      "memory_size_in_bytes" : 0
    },
    "completion" : {
      "size" : "0b",
      "size_in_bytes" : 0
    },
    "segments" : {
      "count" : 37,
      "memory" : "254kb",
      "memory_in_bytes" : 260130,
      "index_writer_memory" : "1mb",
      "index_writer_memory_in_bytes" : 1120520,
      "index_writer_max_memory" : "201.4mb",
      "index_writer_max_memory_in_bytes" : 211261848,
      "version_map_memory" : "184.1kb",
      "version_map_memory_in_bytes" : 188612,
      "fixed_bit_set" : "0b",
      "fixed_bit_set_memory_in_bytes" : 0
    },
    "percolate" : {
      "total" : 0,
      "get_time" : "0s",
      "time_in_millis" : 0,
      "current" : 0,
      "memory_size_in_bytes" : -1,
      "memory_size" : "-1b",
      "queries" : 0
    }
  },
  "nodes" : {
    "count" : {
      "total" : 2,
      "master_only" : 0,
      "data_only" : 0,
      "master_data" : 2,
      "client" : 0
    },
    "versions" : [ "1.5.2" ],
    "os" : {
      "available_processors" : 4,
      "mem" : {
        "total" : "9.6gb",
        "total_in_bytes" : 10404675584
      },
      "cpu" : [ {
        "vendor" : "Intel",
        "model" : "Core(TM) i5-5257U CPU @ 2.70GHz",
        "mhz" : 2699,
        "total_cores" : 2,
        "total_sockets" : 1,
        "cores_per_socket" : 2,
        "cache_size" : "3kb",
        "cache_size_in_bytes" : 3072,
        "count" : 2
      } ]
    },
    "process" : {
      "cpu" : {
        "percent" : 18
      },
      "open_file_descriptors" : {
        "min" : 181,
        "max" : 185,
        "avg" : 183
      }
    },
    "jvm" : {
      "max_uptime" : "5.8m",
      "max_uptime_in_millis" : 350087,
      "versions" : [ {
        "version" : "1.7.0_101",
        "vm_name" : "OpenJDK 64-Bit Server VM",
        "vm_version" : "24.95-b01",
        "vm_vendor" : "Oracle Corporation",
        "count" : 2
      } ],
      "mem" : {
        "heap_used" : "416.1mb",
        "heap_used_in_bytes" : 436373528,
        "heap_max" : "1.9gb",
        "heap_max_in_bytes" : 2112618496
      },
      "threads" : 78
    },
    "fs" : {
      "total" : "78.6gb",
      "total_in_bytes" : 84482326528,
      "free" : "14gb",
      "free_in_bytes" : 15074799616,
      "available" : "10.7gb",
      "available_in_bytes" : 11519651840
    },
    "plugins" : [ ]
  }
}
```

Test with Kibana web server
```
vagrant@vagrant-ubuntu-trusty-64:~$ curl http://10.123.249.243:5601
<!DOCTYPE html>
  <!--[if IE 8]>         <html class="no-js lt-ie9" lang="en"> <![endif]-->
  <!--[if gt IE 8]><!--> <html class="no-js" lang="en"> <!--<![endif]-->
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width">
    <link rel="shortcut icon" href="styles/theme/elk.ico">
    <title>Kibana 4</title>
    <link rel="stylesheet" href="styles/main.css?_b=6004">

  </head>
  <body kibana ng-class="'application-' + activeApp.id">

    <div class="row">
      <div class="col-md-offset-4 col-md-4 page-header initial-load">
        <center>
          <img width="128" src="images/initial_load.gif">
          <h1>
            <strong>Kibana</strong>
            <small id="cache-message">is loading. Give me a moment here. I'm loading a whole bunch of code. Don't worry, all this good stuff will be cached up for next time!</small>
          </h1>
        </center>
      </div>
    </div>

    <script>
      window.KIBANA_VERSION='4.0.2';
      window.KIBANA_BUILD_NUM='6004';
      window.KIBANA_COMMIT_SHA='b2861167bf39066c91556fc03298c38b89ad28a3';
    </script>

    <script src="bower_components/requirejs/require.js?_b=6004"></script>
    <script src="require.config.js?_b=6004"></script>
    <script>
      var showCacheMessage = location.href.indexOf('?embed') < 0 && location.href.indexOf('&embed') < 0;
      if (!showCacheMessage) document.getElementById('cache-message').style.display = 'none';

      if (window.KIBANA_BUILD_NUM.substr(0, 2) !== '@@') {
        // only cache bust if this is really the build number
        require.config({ urlArgs: '_b=' + window.KIBANA_BUILD_NUM });
      }

      require(['kibana'], function (kibana) { kibana.init(); });
    </script>

  </body>
</html>
```

Print log
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log elasticsearch-logging-v1-10flh
I0514 14:48:44.008694       9 elasticsearch_logging_discovery.go:42] Kubernetes Elasticsearch logging discovery
I0514 14:48:44.300618       9 elasticsearch_logging_discovery.go:75] Found []
I0514 14:48:54.304979       9 elasticsearch_logging_discovery.go:75] Found ["172.18.0.15" "172.18.0.16"]
I0514 14:49:04.310064       9 elasticsearch_logging_discovery.go:75] Found ["172.18.0.15" "172.18.0.16"]
I0514 14:49:04.310169       9 elasticsearch_logging_discovery.go:87] Endpoints = ["172.18.0.15" "172.18.0.16"]
[2017-05-14 14:49:14,502][INFO ][node                     ] [Gaza] version[1.5.2], pid[15], build[62ff986/2015-04-27T09:21:06Z]
[2017-05-14 14:49:14,502][INFO ][node                     ] [Gaza] initializing ...
[2017-05-14 14:49:14,599][INFO ][plugins                  ] [Gaza] loaded [], sites []
[2017-05-14 14:49:57,302][INFO ][node                     ] [Gaza] initialized
[2017-05-14 14:49:57,303][INFO ][node                     ] [Gaza] starting ...
[2017-05-14 14:49:59,804][INFO ][transport                ] [Gaza] bound_address {inet[/0:0:0:0:0:0:0:0:9300]}, publish_address {inet[/172.18.0.15:9300]}
[2017-05-14 14:50:00,193][INFO ][discovery                ] [Gaza] kubernetes-logging/a1O1Gcd_SRSOReTRDiK1SQ
[2017-05-14 14:50:03,711][INFO ][cluster.service          ] [Gaza] new_master [Gaza][a1O1Gcd_SRSOReTRDiK1SQ][elasticsearch-logging-v1-10flh][inet[/172.18.0.15:9300]]{master=true}, reason: zen-disco-join (elected_as_master)
[2017-05-14 14:50:03,900][INFO ][http                     ] [Gaza] bound_address {inet[/0:0:0:0:0:0:0:0:9200]}, publish_address {inet[/172.18.0.15:9200]}
[2017-05-14 14:50:03,901][INFO ][node                     ] [Gaza] started

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log elasticsearch-logging-v1-vn881
I0514 14:48:44.310818       7 elasticsearch_logging_discovery.go:42] Kubernetes Elasticsearch logging discovery
I0514 14:48:44.512669       7 elasticsearch_logging_discovery.go:75] Found []
I0514 14:48:54.518926       7 elasticsearch_logging_discovery.go:75] Found ["172.18.0.15" "172.18.0.16"]
I0514 14:49:04.522846       7 elasticsearch_logging_discovery.go:75] Found ["172.18.0.15" "172.18.0.16"]
I0514 14:49:04.522990       7 elasticsearch_logging_discovery.go:87] Endpoints = ["172.18.0.15" "172.18.0.16"]
[2017-05-14 14:49:14,696][INFO ][node                     ] [Decay] version[1.5.2], pid[14], build[62ff986/2015-04-27T09:21:06Z]
[2017-05-14 14:49:14,697][INFO ][node                     ] [Decay] initializing ...
[2017-05-14 14:49:14,711][INFO ][plugins                  ] [Decay] loaded [], sites []
[2017-05-14 14:49:55,793][INFO ][node                     ] [Decay] initialized
[2017-05-14 14:49:55,797][INFO ][node                     ] [Decay] starting ...
[2017-05-14 14:49:58,408][INFO ][transport                ] [Decay] bound_address {inet[/0:0:0:0:0:0:0:0:9300]}, publish_address {inet[/172.18.0.16:9300]}
[2017-05-14 14:49:58,596][INFO ][discovery                ] [Decay] kubernetes-logging/vUblylunSiOqDGtERxh6xQ
[2017-05-14 14:50:04,198][INFO ][cluster.service          ] [Decay] detected_master [Gaza][a1O1Gcd_SRSOReTRDiK1SQ][elasticsearch-logging-v1-10flh][inet[/172.18.0.15:9300]]{master=true}, added {[Gaza][a1O1Gcd_SRSOReTRDiK1SQ][elasticsearch-logging-v1-10flh][inet[/172.18.0.15:9300]]{master=true},}, reason: zen-disco-receive(from master [[Gaza][a1O1Gcd_SRSOReTRDiK1SQ][elasticsearch-logging-v1-10flh][inet[/172.18.0.15:9300]]{master=true}])
[2017-05-14 14:50:04,296][INFO ][http                     ] [Decay] bound_address {inet[/0:0:0:0:0:0:0:0:9200]}, publish_address {inet[/172.18.0.16:9200]}
[2017-05-14 14:50:04,393][INFO ][node                     ] [Decay] started

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log kibana-logging-v1-8gwfq
ELASTICSEARCH_URL=http://elasticsearch-logging:9200
{"@timestamp":"2017-05-14T14:48:57.600Z","level":"error","node_env":"production","error":"Request error, retrying -- connect ECONNREFUSED"}
{"@timestamp":"2017-05-14T14:48:57.603Z","level":"warn","message":"Unable to revive connection: http://elasticsearch-logging:9200/","node_env":"production"}
{"@timestamp":"2017-05-14T14:48:57.604Z","level":"warn","message":"No living connections","node_env":"production"}
{"@timestamp":"2017-05-14T14:48:57.606Z","level":"info","message":"Unable to connect to elasticsearch at http://elasticsearch-logging:9200. Retrying in 2.5 seconds.","node_env":"production"}

{"@timestamp":"2017-05-14T14:50:13.799Z","level":"info","message":"No existing kibana index found","node_env":"production"}
{"@timestamp":"2017-05-14T14:50:13.900Z","level":"info","message":"Listening on 0.0.0.0:5601","node_env":"production"}

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log fluentd-elasticsearch
2017-05-14 14:48:44 +0000 [info]: reading config file path="/etc/td-agent/td-agent.conf"
2017-05-14 14:48:44 +0000 [info]: starting fluentd-0.12.19
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-mixin-config-placeholders' version '0.3.0'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-mixin-plaintextformatter' version '0.2.6'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-docker_metadata_filter' version '0.1.1'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-elasticsearch' version '1.3.0'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-kubernetes_metadata_filter' version '0.14.0'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-mongo' version '0.7.11'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-rewrite-tag-filter' version '1.5.3'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-s3' version '0.6.4'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-scribe' version '0.10.14'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-td' version '0.10.28'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-td-monitoring' version '0.2.1'
2017-05-14 14:48:45 +0000 [info]: gem 'fluent-plugin-webhdfs' version '0.4.1'
2017-05-14 14:48:45 +0000 [info]: gem 'fluentd' version '0.12.19'
2017-05-14 14:48:45 +0000 [info]: adding match pattern="fluent.**" type="null"
2017-05-14 14:48:45 +0000 [info]: adding filter pattern="kubernetes.**" type="kubernetes_metadata"
2017-05-14 14:48:45 +0000 [info]: adding match pattern="**" type="elasticsearch"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: adding source type="tail"
2017-05-14 14:48:45 +0000 [info]: using configuration file: <ROOT>
  <match fluent.**>
    type null
  </match>
  <source>
    type tail
    path /var/log/containers/*.log
    pos_file /var/log/es-containers.log.pos
    time_format %Y-%m-%dT%H:%M:%S.%NZ
    tag kubernetes.*
    format json
    read_from_head true
  </source>
  <source>
    type tail
    format /^(?<time>[^ ]* [^ ,]*)[^\[]*\[[^\]]*\]\[(?<severity>[^ \]]*) *\] (?<message>.*)$/
    time_format %Y-%m-%d %H:%M:%S
    path /var/log/salt/minion
    pos_file /var/log/gcp-salt.pos
    tag salt
  </source>
  <source>
    type tail
    format syslog
    path /var/log/startupscript.log
    pos_file /var/log/es-startupscript.log.pos
    tag startupscript
  </source>
  <source>
    type tail
    format /^time="(?<time>[^)]*)" level=(?<severity>[^ ]*) msg="(?<message>[^"]*)"( err="(?<error>[^"]*)")?( statusCode=($<status_code>\d+))?/
    time_format %Y-%m-%dT%H:%M:%S.%NZ
    path /var/log/docker.log
    pos_file /var/log/es-docker.log.pos
    tag docker
  </source>
  <source>
    type tail
    format none
    path /var/log/etcd.log
    pos_file /var/log/es-etcd.log.pos
    tag etcd
  </source>
  <source>
    type tail
    format multiline
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/kubelet.log
    pos_file /var/log/es-kubelet.log.pos
    tag kubelet
  </source>
  <source>
    type tail
    format multiline
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/kube-apiserver.log
    pos_file /var/log/es-kube-apiserver.log.pos
    tag kube-apiserver
  </source>
  <source>
    type tail
    format multiline
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/kube-controller-manager.log
    pos_file /var/log/es-kube-controller-manager.log.pos
    tag kube-controller-manager
  </source>
  <source>
    type tail
    format multiline
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/kube-scheduler.log
    pos_file /var/log/es-kube-scheduler.log.pos
    tag kube-scheduler
  </source>
  <filter kubernetes.**>
    type kubernetes_metadata
  </filter>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/glbc.log
    pos_file /var/log/es-glbc.log.pos
    tag glbc
  </source>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/cluster-autoscaler.log
    pos_file /var/log/es-cluster-autoscaler.log.pos
    tag cluster-autoscaler
  </source>
  <match **>
    type elasticsearch
    log_level info
    include_tag_key true
    host elasticsearch-logging
    port 9200
    logstash_format true
    buffer_chunk_limit 2M
    buffer_queue_limit 32
    flush_interval 5s
    max_retry_wait 30
    disable_retry_limit 
  </match>
</ROOT>
2017-05-14 14:48:45 +0000 [warn]: parameter 'multiline_flush_interval' in <source>
  type tail
  format multiline
  multiline_flush_interval 5s
  format_firstline /^\w\d{4}/
  format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
  time_format %m%d %H:%M:%S.%N
  path /var/log/glbc.log
  pos_file /var/log/es-glbc.log.pos
  tag glbc
</source> is not used.
2017-05-14 14:48:45 +0000 [warn]: parameter 'multiline_flush_interval' in <source>
  type tail
  format multiline
  multiline_flush_interval 5s
  format_firstline /^\w\d{4}/
  format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
  time_format %m%d %H:%M:%S.%N
  path /var/log/cluster-autoscaler.log
  pos_file /var/log/es-cluster-autoscaler.log.pos
  tag cluster-autoscaler
</source> is not used.

2017-05-14 14:50:16 +0000 [info]: Connection opened to Elasticsearch cluster => {:host=>"elasticsearch-logging", :port=>9200, :scheme=>"http"}

```

For browsing Kibana
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system expose service kibana-logging --name=kibana-logging-ex --target-port=5601 --type=NodePort
service "kibana-logging-ex" exposed
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get service kibana-logging-ex
NAME                CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kibana-logging-ex   10.123.255.91   <nodes>       5601:31725/TCP   11s
```

![Intial Page](./屏幕快照%202017-05-14%20上午8.11.00.png)

### V1.5.7

From _saltbase_
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxvf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch
kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/
kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/kibana-controller.yaml
kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/kibana-service.yaml
kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/es-service.yaml
kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/es-controller.yaml
kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/es-image/
kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/es-image/template-k8s-logstash.json
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls /opt/kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch
es-controller.yaml  es-image  es-service.yaml  kibana-controller.yaml  kibana-service.yaml
vagrant@vagrant-ubuntu-trusty-64:~$ sudo tar -C /opt -zxvf /opt/kubernetes/server/kubernetes-salt.tar.gz kubernetes/saltbase/salt/fluentd-es
kubernetes/saltbase/salt/fluentd-es/
kubernetes/saltbase/salt/fluentd-es/fluentd-es.yaml
kubernetes/saltbase/salt/fluentd-es/init.sls
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ls /opt/kubernetes/saltbase/salt/fluentd-es
fluentd-es.yaml  init.sls
```

Pull
```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/elasticsearch:v2.4.1
v2.4.1: Pulling from google_containers/elasticsearch
1fad42e8a0d9: Pull complete 
2ed832147d3e: Pull complete 
3e2e387eb26a: Pull complete 
eef540699244: Pull complete 
1624a2f8d114: Pull complete 
7018f4ec6e0a: Pull complete 
6ca3bc2ad3b3: Pull complete 
64fb22c146a4: Pull complete 
6e8a23b5aaad: Pull complete 
2e53bc96a04d: Pull complete 
06dbadc7e0dd: Pull complete 
2ee41ff03c33: Pull complete 
976b1271c2ec: Pull complete 
f38eda110397: Pull complete 
1351b011b63b: Pull complete 
Digest: sha256:20fd799104ec5d95ebe2e1444ab1ca2aa5d6eb0feb0caa8f7a00f80170c1a504
Status: Image is up to date for gcr.io/google_containers/elasticsearch:v2.4.1
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/kibana:v4.6.1
v4.6.1: Pulling from google_containers/kibana
169bb74c2cc6: Pull complete 
a3ed95caeb02: Pull complete 
49518420053c: Pull complete 
26fcbd3c88d7: Pull complete 
e1cfe75be08b: Pull complete 
Digest: sha256:5de6725b17c9affbff52a1cb6d66cc279dc27a6f94b1c0d3c15dcf5c42ddbdf3
Status: Image is up to date for gcr.io/google_containers/kibana:v4.6.1
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/fluentd-elasticsearch:1.22
1.22: Pulling from google_containers/fluentd-elasticsearch
be4ee6c45489: Pull complete 
186fba63a286: Pull complete 
eb2d94fe8f23: Pull complete 
4a3658c3601f: Pull complete 
Digest: sha256:8b2244d0c10c70fad3dac392aab8efbb235cc84e0d2f3ae6a11291b6b35f0845
Status: Downloaded newer image for gcr.io/google_containers/fluentd-elasticsearch:1.22
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/fluentd-elasticsearch:1.20
1.20: Pulling from google_containers/fluentd-elasticsearch
169bb74c2cc6: Already exists 
a3ed95caeb02: Already exists 
fa2b73488d2b: Pull complete 
8adc08acc1c8: Pull complete 
220ee2c01d13: Pull complete 
Digest: sha256:de60e9048b3b79c4e53e66dc3b8ee58b4be2a88395a4b33407f8bbadc9548dfb
Status: Downloaded newer image for gcr.io/google_containers/fluentd-elasticsearch:1.20
```

Deploy
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl create -f /opt/kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/
replicationcontroller "elasticsearch-logging-v1" created
service "elasticsearch-logging" created
deployment "kibana-logging" created
service "kibana-logging" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo sed 's/200Mi/100Mi/g' /opt/kubernetes/saltbase/salt/fluentd-es/fluentd-es.yaml | sudo tee /etc/kubernetes/manifests/addons/fluentd-es.yaml
apiVersion: v1
kind: Pod
metadata:
  name: fluentd-elasticsearch
  namespace: kube-system
  labels:
    k8s-app: fluentd-logging
spec:
  containers:
  - name: fluentd-elasticsearch
    image: gcr.io/google_containers/fluentd-elasticsearch:1.20
    resources:
      limits:
        memory: 100Mi
      requests:
        cpu: 100m
        memory: 100Mi
    volumeMounts:
    - name: varlog
      mountPath: /var/log
    - name: varlibdockercontainers
      mountPath: /var/lib/docker/containers
      readOnly: true
  terminationGracePeriodSeconds: 30
  volumes:
  - name: varlog
    hostPath:
      path: /var/log
  - name: varlibdockercontainers
    hostPath:
      path: /var/lib/docker/containers

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl create -f /etc/kubernetes/manifests/addons
POD "fluentd-logging" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get svc,po -l k8s-app=elasticsearch-logging
NAME                        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
svc/elasticsearch-logging   10.123.246.99   <none>        9200/TCP   1h

NAME                                READY     STATUS    RESTARTS   AGE
po/elasticsearch-logging-v1-2lmt3   1/1       Running   0          1h
po/elasticsearch-logging-v1-bpvrs   1/1       Running   0          1h
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get pods -l k8s-app=fluentd-logging
NAME                    READY     STATUS    RESTARTS   AGE
fluentd-elasticsearch   1/1       Running   0          2m
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get svc,po -l k8s-app=kibana-logging
NAME                    CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
svc/kibana-logging      10.123.249.244   <none>        5601/TCP         2h

NAME                                 READY     STATUS    RESTARTS   AGE
po/kibana-logging-2924323056-2l63r   1/1       Running   0          2h
```

Test with ES cluster APIs
```
vagrant@vagrant-ubuntu-trusty-64:~$ curl http://10.123.246.99:9200/
{
  "name" : "Robert Hunter",
  "cluster_name" : "kubernetes-logging",
  "cluster_uuid" : "JB0Rn64bRWql5zYIoUdvNg",
  "version" : {
    "number" : "2.4.1",
    "build_hash" : "c67dc32e24162035d18d6fe1e952c4cbcbe79d16",
    "build_timestamp" : "2016-09-27T18:57:55Z",
    "build_snapshot" : false,
    "lucene_version" : "5.5.2"
  },
  "tagline" : "You Know, for Search"
}
vagrant@vagrant-ubuntu-trusty-64:~$ curl 10.123.240.190:9200/_cluster/health
{"cluster_name":"kubernetes-logging","status":"green","timed_out":false,"number_of_nodes":2,"number_of_data_nodes":2,"active_primary_shards":6,"active_shards":12,"relocating_shards":0,"initializing_shards":0,"unassigned_shards":0,"delayed_unassigned_shards":0,"number_of_pending_tasks":0,"number_of_in_flight_fetch":0,"task_max_waiting_in_queue_millis":0,"active_shards_percent_as_number":100.0}
vagrant@vagrant-ubuntu-trusty-64:~$ curl 10.123.240.190:9200/_cluster/state
vagrant@vagrant-ubuntu-trusty-64:~$ curl 10.123.240.190:9200/_cluster/stats?human\&pretty
{
  "timestamp" : 1494757789870,
  "cluster_name" : "kubernetes-logging",
  "status" : "green",
  "indices" : {
    "count" : 2,
    "shards" : {
      "total" : 12,
      "primaries" : 6,
      "replication" : 1.0,
      "index" : {
        "shards" : {
          "min" : 2,
          "max" : 10,
          "avg" : 6.0
        },
        "primaries" : {
          "min" : 1,
          "max" : 5,
          "avg" : 3.0
        },
        "replication" : {
          "min" : 1.0,
          "max" : 1.0,
          "avg" : 1.0
        }
      }
    },
    "docs" : {
      "count" : 23749,
      "deleted" : 0
    },
    "store" : {
      "size" : "9.1mb",
      "size_in_bytes" : 9586017,
      "throttle_time" : "0s",
      "throttle_time_in_millis" : 0
    },
    "fielddata" : {
      "memory_size" : "0b",
      "memory_size_in_bytes" : 0,
      "evictions" : 0
    },
    "query_cache" : {
      "memory_size" : "0b",
      "memory_size_in_bytes" : 0,
      "total_count" : 484,
      "hit_count" : 0,
      "miss_count" : 484,
      "cache_size" : 0,
      "cache_count" : 0,
      "evictions" : 0
    },
    "completion" : {
      "size" : "0b",
      "size_in_bytes" : 0
    },
    "segments" : {
      "count" : 69,
      "memory" : "293.1kb",
      "memory_in_bytes" : 300236,
      "terms_memory" : "226.3kb",
      "terms_memory_in_bytes" : 231816,
      "stored_fields_memory" : "24kb",
      "stored_fields_memory_in_bytes" : 24624,
      "term_vectors_memory" : "0b",
      "term_vectors_memory_in_bytes" : 0,
      "norms_memory" : "33.5kb",
      "norms_memory_in_bytes" : 34304,
      "doc_values_memory" : "9.2kb",
      "doc_values_memory_in_bytes" : 9492,
      "index_writer_memory" : "0b",
      "index_writer_memory_in_bytes" : 0,
      "index_writer_max_memory" : "202.4mb",
      "index_writer_max_memory_in_bytes" : 212285840,
      "version_map_memory" : "0b",
      "version_map_memory_in_bytes" : 0,
      "fixed_bit_set" : "0b",
      "fixed_bit_set_memory_in_bytes" : 0
    },
    "percolate" : {
      "total" : 0,
      "time" : "0s",
      "time_in_millis" : 0,
      "current" : 0,
      "memory_size_in_bytes" : -1,
      "memory_size" : "-1b",
      "queries" : 0
    }
  },
  "nodes" : {
    "count" : {
      "total" : 2,
      "master_only" : 0,
      "data_only" : 0,
      "master_data" : 2,
      "client" : 0
    },
    "versions" : [ "2.4.1" ],
    "os" : {
      "available_processors" : 4,
      "allocated_processors" : 4,
      "mem" : {
        "total" : "295.8mb",
        "total_in_bytes" : 310255616
      },
      "names" : [ {
        "name" : "Linux",
        "count" : 2
      } ]
    },
    "process" : {
      "cpu" : {
        "percent" : 3
      },
      "open_file_descriptors" : {
        "min" : 198,
        "max" : 200,
        "avg" : 199
      }
    },
    "jvm" : {
      "max_uptime" : "28.4m",
      "max_uptime_in_millis" : 1706873,
      "versions" : [ {
        "version" : "1.8.0_111",
        "vm_name" : "OpenJDK 64-Bit Server VM",
        "vm_version" : "25.111-b14",
        "vm_vendor" : "Oracle Corporation",
        "count" : 2
      } ],
      "mem" : {
        "heap_used" : "192.9mb",
        "heap_used_in_bytes" : 202334264,
        "heap_max" : "1.9gb",
        "heap_max_in_bytes" : 2112618496
      },
      "threads" : 79
    },
    "fs" : {
      "total" : "78.6gb",
      "total_in_bytes" : 84482326528,
      "free" : "18.7gb",
      "free_in_bytes" : 20083646464,
      "available" : "15.3gb",
      "available_in_bytes" : 16528498688,
      "spins" : "true"
    },
    "plugins" : [ ]
  }
}
```

Print log
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log po/elasticsearch-logging-v1-2lmt3
[2017-05-14 08:05:48,016][INFO ][node                     ] [Robert Hunter] version[2.4.1], pid[16], build[c67dc32/2016-09-27T18:57:55Z]
[2017-05-14 08:05:48,018][INFO ][node                     ] [Robert Hunter] initializing ...
[2017-05-14 08:05:49,643][INFO ][plugins                  ] [Robert Hunter] modules [reindex, lang-expression, lang-groovy], plugins [], sites []
[2017-05-14 08:05:49,700][INFO ][env                      ] [Robert Hunter] using [1] data paths, mounts [[/data (/dev/sda1)]], net usable_space [8.7gb], net total_space [39.3gb], spins? [possibly], types [ext4]
[2017-05-14 08:05:49,700][INFO ][env                      ] [Robert Hunter] heap size [1007.3mb], compressed ordinary object pointers [true]
[2017-05-14 08:05:56,238][INFO ][node                     ] [Robert Hunter] initialized
[2017-05-14 08:05:56,238][INFO ][node                     ] [Robert Hunter] starting ...
[2017-05-14 08:05:56,581][INFO ][transport                ] [Robert Hunter] publish_address {172.18.0.5:9300}, bound_addresses {[::]:9300}
[2017-05-14 08:05:56,641][INFO ][discovery                ] [Robert Hunter] kubernetes-logging/E-ehvWGARXyhkpGmLdl6Og
[2017-05-14 08:06:12,759][INFO ][cluster.service          ] [Robert Hunter] new_master {Robert Hunter}{E-ehvWGARXyhkpGmLdl6Og}{172.18.0.5}{172.18.0.5:9300}{master=true}, added {{Libra}{yU3-BXyVT16SQ0vZGAHSNg}{172.18.0.7}{172.18.0.7:9300}{master=true},}, reason: zen-disco-join(elected_as_master, [1] joins received)
[2017-05-14 08:06:28,389][WARN ][discovery                ] [Robert Hunter] waited for 30s and no initial state was set by the discovery
[2017-05-14 08:06:32,761][INFO ][http                     ] [Robert Hunter] publish_address {172.18.0.5:9200}, bound_addresses {[::]:9200}
[2017-05-14 08:06:32,772][INFO ][node                     ] [Robert Hunter] started
[2017-05-14 08:06:32,886][WARN ][cluster.service          ] [Robert Hunter] cluster state update task [zen-disco-join(elected_as_master, [1] joins received)] took 31.6s above the warn threshold of 30s
[2017-05-14 08:06:37,386][INFO ][gateway                  ] [Robert Hunter] recovered [0] indices into cluster_state
[2017-05-14 08:21:18,496][INFO ][cluster.metadata         ] [Robert Hunter] [.kibana] creating index, cause [api], templates [], shards [1]/[1], mappings [config]
[2017-05-14 08:21:21,801][INFO ][cluster.routing.allocation] [Robert Hunter] Cluster health status changed from [RED] to [YELLOW] (reason: [shards started [[.kibana][0]] ...]).
[2017-05-14 08:21:23,262][INFO ][cluster.routing.allocation] [Robert Hunter] Cluster health status changed from [YELLOW] to [GREEN] (reason: [shards started [[.kibana][0]] ...]).
[2017-05-14 09:34:30,540][INFO ][cluster.metadata         ] [Robert Hunter] [logstash-2017.05.13] creating index, cause [auto(bulk api)], templates [], shards [5]/[1], mappings []
[2017-05-14 09:34:31,135][INFO ][cluster.metadata         ] [Robert Hunter] [logstash-2017.05.14] creating index, cause [auto(bulk api)], templates [], shards [5]/[1], mappings []
[2017-05-14 09:34:32,451][INFO ][cluster.routing.allocation] [Robert Hunter] Cluster health status changed from [RED] to [YELLOW] (reason: [shards started [[logstash-2017.05.14][1], [logstash-2017.05.14][4], [logstash-2017.05.14][1], [logstash-2017.05.14][3]] ...]).
[2017-05-14 09:34:33,026][INFO ][cluster.metadata         ] [Robert Hunter] [logstash-2017.05.06] creating index, cause [auto(bulk api)], templates [], shards [5]/[1], mappings []
[2017-05-14 09:34:33,876][INFO ][cluster.routing.allocation] [Robert Hunter] Cluster health status changed from [RED] to [YELLOW] (reason: [shards started [[logstash-2017.05.13][0], [logstash-2017.05.14][1], [logstash-2017.05.13][0], [logstash-2017.05.14][1], [logstash-2017.05.06][3], [logstash-2017.05.06][1], [logstash-2017.05.14][2], [logstash-2017.05.14][0], [logstash-2017.05.14][2], [logstash-2017.05.14][0]] ...]).
[2017-05-14 09:34:34,384][INFO ][cluster.metadata         ] [Robert Hunter] [logstash-2017.05.14] create_mapping [fluentd]
[2017-05-14 09:34:34,435][INFO ][cluster.metadata         ] [Robert Hunter] [logstash-2017.05.13] create_mapping [fluentd]
[2017-05-14 09:34:40,178][INFO ][cluster.routing.allocation] [Robert Hunter] Cluster health status changed from [YELLOW] to [GREEN] (reason: [shards started [[logstash-2017.05.13][4]] ...]).
[2017-05-14 09:34:42,080][INFO ][cluster.metadata         ] [Robert Hunter] [logstash-2017.05.06] create_mapping [fluentd]
[2017-05-14 09:35:18,981][WARN ][monitor.jvm              ] [Robert Hunter] [gc][young][5177][9] duration [11.5s], collections [2]/[21.4s], total [11.5s]/[12.1s], memory [195.6mb]->[85.7mb]/[1007.3mb], all_pools {[young] [133.1mb]->[2.3mb]/[133.1mb]}{[survivor] [16.6mb]->[12.7mb]/[16.6mb]}{[old] [45.8mb]->[70.5mb]/[857.6mb]}
[2017-05-14 09:36:56,298][WARN ][monitor.jvm              ] [Robert Hunter] [gc][young][5224][11] duration [1.9s], collections [1]/[2.1s], total [1.9s]/[14.2s], memory [190.5mb]->[59mb]/[1007.3mb], all_pools {[young] [133.1mb]->[274.1kb]/[133.1mb]}{[survivor] [8mb]->[9.4mb]/[16.6mb]}{[old] [49.3mb]->[49.3mb]/[857.6mb]}
[2017-05-14 09:37:12,095][INFO ][cluster.metadata         ] [Robert Hunter] [logstash-2017.05.14] update_mapping [fluentd]
[2017-05-14 09:37:12,194][INFO ][cluster.metadata         ] [Robert Hunter] [logstash-2017.05.14] update_mapping [fluentd]
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log po/elasticsearch-logging-v1-bpvrs
[2017-05-14 08:05:41,192][INFO ][node                     ] [Libra] version[2.4.1], pid[19], build[c67dc32/2016-09-27T18:57:55Z]
[2017-05-14 08:05:41,194][INFO ][node                     ] [Libra] initializing ...
[2017-05-14 08:05:42,264][INFO ][plugins                  ] [Libra] modules [reindex, lang-expression, lang-groovy], plugins [], sites []
[2017-05-14 08:05:42,328][INFO ][env                      ] [Libra] using [1] data paths, mounts [[/data (/dev/sda1)]], net usable_space [8.7gb], net total_space [39.3gb], spins? [possibly], types [ext4]
[2017-05-14 08:05:42,328][INFO ][env                      ] [Libra] heap size [1007.3mb], compressed ordinary object pointers [true]
[2017-05-14 08:05:47,840][INFO ][node                     ] [Libra] initialized
[2017-05-14 08:05:47,843][INFO ][node                     ] [Libra] starting ...
[2017-05-14 08:05:48,144][INFO ][transport                ] [Libra] publish_address {172.18.0.7:9300}, bound_addresses {[::]:9300}
[2017-05-14 08:05:48,169][INFO ][discovery                ] [Libra] kubernetes-logging/yU3-BXyVT16SQ0vZGAHSNg
[2017-05-14 08:06:18,380][WARN ][discovery                ] [Libra] waited for 30s and no initial state was set by the discovery
[2017-05-14 08:06:20,584][INFO ][http                     ] [Libra] publish_address {172.18.0.7:9200}, bound_addresses {[::]:9200}
[2017-05-14 08:06:20,585][INFO ][node                     ] [Libra] started
[2017-05-14 08:06:29,613][INFO ][cluster.service          ] [Libra] detected_master {Robert Hunter}{E-ehvWGARXyhkpGmLdl6Og}{172.18.0.5}{172.18.0.5:9300}{master=true}, added {{Robert Hunter}{E-ehvWGARXyhkpGmLdl6Og}{172.18.0.5}{172.18.0.5:9300}{master=true},}, reason: zen-disco-receive(from master [{Robert Hunter}{E-ehvWGARXyhkpGmLdl6Og}{172.18.0.5}{172.18.0.5:9300}{master=true}])
[2017-05-14 09:35:52,712][WARN ][monitor.jvm              ] [Libra] [gc][young][5219][8] duration [3.8s], collections [1]/[6.9s], total [3.8s]/[5s], memory [185.4mb]->[60.6mb]/[1007.3mb], all_pools {[young] [130.6mb]->[3mb]/[133.1mb]}{[survivor] [16.6mb]->[14.6mb]/[16.6mb]}{[old] [38.2mb]->[45mb]/[857.6mb]}

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log fluentd-elasticsearch
2017-05-14 09:34:28 +0000 [info]: reading config file path="/etc/td-agent/td-agent.conf"
2017-05-14 09:34:28 +0000 [info]: starting fluentd-0.12.29
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-mixin-config-placeholders' version '0.4.0'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-mixin-plaintextformatter' version '0.2.6'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-docker_metadata_filter' version '0.1.3'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-elasticsearch' version '1.5.0'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-kafka' version '0.3.1'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-kubernetes_metadata_filter' version '0.24.0'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-mongo' version '0.7.15'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-rewrite-tag-filter' version '1.5.5'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-s3' version '0.7.1'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-scribe' version '0.10.14'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-td' version '0.10.29'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-td-monitoring' version '0.2.2'
2017-05-14 09:34:28 +0000 [info]: gem 'fluent-plugin-webhdfs' version '0.4.2'
2017-05-14 09:34:28 +0000 [info]: gem 'fluentd' version '0.12.29'
2017-05-14 09:34:28 +0000 [info]: adding match pattern="fluent.**" type="null"
2017-05-14 09:34:28 +0000 [info]: adding match pattern="**" type="elasticsearch"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: adding source type="tail"
2017-05-14 09:34:29 +0000 [info]: using configuration file: <ROOT>
  <match fluent.**>
    type null
  </match>
  <source>
    type tail
    path /var/log/containers/*.log
    pos_file /var/log/es-containers.log.pos
    time_format %Y-%m-%dT%H:%M:%S.%NZ
    tag kubernetes.*
    format json
    read_from_head true
  </source>
  <source>
    type tail
    format /^(?<time>[^ ]* [^ ,]*)[^\[]*\[[^\]]*\]\[(?<severity>[^ \]]*) *\] (?<message>.*)$/
    time_format %Y-%m-%d %H:%M:%S
    path /var/log/salt/minion
    pos_file /var/log/es-salt.pos
    tag salt
  </source>
  <source>
    type tail
    format syslog
    path /var/log/startupscript.log
    pos_file /var/log/es-startupscript.log.pos
    tag startupscript
  </source>
  <source>
    type tail
    format /^time="(?<time>[^)]*)" level=(?<severity>[^ ]*) msg="(?<message>[^"]*)"( err="(?<error>[^"]*)")?( statusCode=($<status_code>\d+))?/
    path /var/log/docker.log
    pos_file /var/log/es-docker.log.pos
    tag docker
  </source>
  <source>
    type tail
    format none
    path /var/log/etcd.log
    pos_file /var/log/es-etcd.log.pos
    tag etcd
  </source>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/kubelet.log
    pos_file /var/log/es-kubelet.log.pos
    tag kubelet
  </source>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/kube-proxy.log
    pos_file /var/log/es-kube-proxy.log.pos
    tag kube-proxy
  </source>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/kube-apiserver.log
    pos_file /var/log/es-kube-apiserver.log.pos
    tag kube-apiserver
  </source>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/kube-controller-manager.log
    pos_file /var/log/es-kube-controller-manager.log.pos
    tag kube-controller-manager
  </source>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/kube-scheduler.log
    pos_file /var/log/es-kube-scheduler.log.pos
    tag kube-scheduler
  </source>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/rescheduler.log
    pos_file /var/log/es-rescheduler.log.pos
    tag rescheduler
  </source>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/glbc.log
    pos_file /var/log/es-glbc.log.pos
    tag glbc
  </source>
  <source>
    type tail
    format multiline
    multiline_flush_interval 5s
    format_firstline /^\w\d{4}/
    format1 /^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/
    time_format %m%d %H:%M:%S.%N
    path /var/log/cluster-autoscaler.log
    pos_file /var/log/es-cluster-autoscaler.log.pos
    tag cluster-autoscaler
  </source>
  <match **>
    type elasticsearch
    log_level info
    include_tag_key true
    host elasticsearch-logging
    port 9200
    logstash_format true
    buffer_chunk_limit 2M
    buffer_queue_limit 32
    flush_interval 5s
    max_retry_wait 30
    disable_retry_limit 
    num_threads 8
  </match>
</ROOT>
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-dns-2185667875-gkjlh_kube-system_healthz-531b7b4c078d731a9ddae2c90a207535a5543013da1f503f56c3c3d5ceb2a161.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kafka-1_kafka_POD-44f0f9431bb396601baea3afad5c2547b2692a7889bbba3ee9bf4b6d206ef22a.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/hostpath-provisioner_default_POD-7401ad08fbf816392bcc64af1f185b3cfe9e6e4040312f7162b3128d960ebb79.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/fluentd-elasticsearch_kube-system_fluentd-elasticsearch-02aed7573e80956c6c43d5512918ea1161cd34c3d6b2e1495a39f2f43beacb16.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-dns-2185667875-gkjlh_kube-system_dnsmasq-59ca5eba331763781a6a56fd8fb21a65dc2d82185de1f637f0cb095ac4fc764c.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-dns-2185667875-gkjlh_kube-system_healthz-e776888834e3423362b4f225c782c534eff5ca388f60cd4486628b639de5fa19.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/nginx-controller-zr4p2_thisisatest_nginx-ef9460f56f77292eb1803214ee680911efacd162eabab57fe6b1479b18358a4b.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/zoo-1_kafka_POD-14a708da672a8d19f1b65e345974577a3484adc4bb623a56a27df09e8b845c0f.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-scheduler-172.17.4.200_kube-system_kube-scheduler-14c826342ea2d80b6fc1c4ea16715c670838c65744b484401f52ec3a99d1dd52.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kubernetes-dashboard-3543765157-87fkr_kube-system_kubernetes-dashboard-e6adfe74a0ce6d191506b8459377cb7c25523a561d24d643bfcbc89f5c443b16.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/zoo-1_kafka_zookeeper-37c6a7217bf8461233184e9d23198936f169d7ba1b42ef748acfb522ef8a5439.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-dns-2185667875-gkjlh_kube-system_POD-97945e96ee788d66ecf2a2e58d26f0072325b95d8173793deaacdc42a4dff12b.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/zoo-4_kafka_POD-0db2d1d9221e05386a13dd8856a4416170e6469cb5b2120992a8430846cd3444.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-controller-manager-172.17.4.200_kube-system_POD-d995bcd8ff31922620dbda7377635f722cfd37113984ac4e8fd2f927e7fdb74b.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-proxy-172.17.4.200_kube-system_POD-c41acb1be9fa4b0fc179af1e9e50e10770a27c91fc87bc4929d339389d6d375d.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/zoo-2_kafka_zookeeper-f5e0010b4a90de7743b31e8944ffc2598c712967f5f058671b9ae0b539729443.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-dns-2185667875-gkjlh_kube-system_kubedns-0692848e9704c42c6d425860221d7231cb4f24933b33d5087d4a6e324c02cd68.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-proxy-172.17.4.200_kube-system_kube-proxy-58424ead3b682c286d24e43f70761f253a042dc15ce54c3bb1d501fe9a69e56e.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/zoo-4_kafka_POD-45244ca2bd38c769018137c23a2f5061e77fcba06fd30c2dbfa7c5fbd199a671.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kube-apiserver-172.17.4.200_kube-system_POD-5e9566f96b10bc6b87520c50ebd2b07895773f9c50db5271f3537d6b5b28a556.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/topic-create-test1-x2zqm_kafka_kafka-744f7340902015181b2bde5a9e27ca2455f984f31758895553e5645f0732d10f.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kubernetes-dashboard-3543765157-87fkr_kube-system_POD-0dcfef3751a713642b79b543b2909b4c69e814874d21c5754cb606278866965d.log
2017-05-14 09:34:29 +0000 [info]: following tail of /var/log/containers/kafka-0_kafka_broker-68f4dc99ae84510d23097e5d2d4e5b5e4923f294f1be285ccc857225e3936d7d.log
2017-05-14 09:34:29 +0000 [info]: Connection opened to Elasticsearch cluster => {:host=>"elasticsearch-logging", :port=>9200, :scheme=>"http"}
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-dns-2185667875-gkjlh_kube-system_kubedns-5f6c4467fbd34a343ed4265a06111160f6025cb9768a980b58e8007518395a57.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/etcd-server-172.17.4.200_kube-system_POD-aa08d65a1053532cfc58b810e6b5ef3727ef2b1a96a4bc70365a651aa7e6fd57.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-controller-manager-172.17.4.200_kube-system_kube-controller-manager-f5616e6edd6a2c78dcaa4f79f6838a5334d0aba05412967c4f5e5b1084757131.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kubernetes-dashboard-3543765157-87fkr_kube-system_kubernetes-dashboard-8620b6be02046c6160d733c7c44fe14d48e00069f8b30022c88db704c1147c4f.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kafka-2_kafka_POD-f6adf67f8fdcdd58751dacab23e14642cb9841cc240d724ba998dfc7b3906372.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-dns-2185667875-gkjlh_kube-system_POD-000633c6edff255bc4346ba4a2213e21be2ded08e2e7f756c8e6b3bd5737715d.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-apiserver-172.17.4.200_kube-system_kube-apiserver-c9a2b09de77dab199b2bb4ed1afe145a9e750d7afc33722a918f0781ac350c3e.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-1_kafka_POD-4878f2a40210a00772c4353d1e6cd45569e7f1e3dffb357003aa90d1012a0b53.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-dns-2185667875-gkjlh_kube-system_dnsmasq-f180ea5263e9c006fd564a250011301a7a9b7db62b1963767b2d8041831d1d05.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/fluentd-elasticsearch_kube-system_POD-043a0dc2ed817d23c05e89dbecfcb01d24881ede9e8403287ed970ab8dc3f854.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/nginx-controller-zr4p2_thisisatest_nginx-d6594192e3e825124a6b351b6d29ef5771e8df943491ad3089c7ae81beac577c.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-2_kafka_POD-e123fc7fa8d8326bc075917965909705234129b5d42c563811320573b5a4b74f.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-3_kafka_POD-1a03f76c97b0654bc5cb374352d9215720564cffbedbae1f36946e4bb725eb45.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/hostpath-provisioner_default_hostpath-provisioner-bdb6003c618d141cda79af352529de61829f708bc200ba55fa7b15069166a827.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-0_kafka_zookeeper-6a65fd076ea3246ad58f4bd87747e14763bdd2f88d35fb2d858b1619dc033440.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kafka-1_kafka_broker-7bb1df313b6f1365303b68e4961e576c6d8d469fcf252c18459c4f0837384c73.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/topic-create-test1-x2zqm_kafka_POD-18b22525a21aaa7c69ec8ce16658aa3422ea5b23cd939c44de804a1f844c9a51.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kafka-0_kafka_broker-15780bf9e45d07097a31cb2ff8dd861aea4ddbf6245b7e3acc82074f89407951.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kibana-logging-2924323056-2l63r_kube-system_POD-85393a04bc36c77b2bf94aa372a82b36b20301dbd18c9c0a1916197a10d4a665.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-1_kafka_zookeeper-86f5da2130160206f6f0c428fcfe551fa23d03afec09d7640aa8ea0150e7141e.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kibana-logging-2924323056-2l63r_kube-system_kibana-logging-e0a58691fd23653a37e14f84a0d69a3f6c6719b6ec809f171623fd7828c96626.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kafka-1_kafka_broker-8df26e9f12c02317c1fc179c51899406e9f58abae2b3ad476d56e37db9a8b503.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/elasticsearch-logging-v1-2lmt3_kube-system_elasticsearch-logging-e60e3111bce0e0748b7d5b06db51f40ac6851eb823991551ac1466f725d09005.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-dns-2185667875-gkjlh_kube-system_dnsmasq-metrics-f71e65d8bb58b441356e31a87a5a160b7549d1c373719539dc867737c1a31231.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/elasticsearch-logging-v1-bpvrs_kube-system_POD-f0a15a5e2ec27f36b3250d358fa4481150c167c42cf0b30592c931bc73563441.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/hostpath-provisioner_default_POD-00a1cccc8295782c6f02fc3abf4721397fe47584bf87d92c7c4ac8fab90a3e4b.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-3_kafka_POD-8754322c2c858c11195a1e790675e16603f08daefe7d881c8a5b23abe92c5835.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kafka-1_kafka_POD-6002e214ce97ce0b6c0bbb1dec07996b11e518dc06b78471893d7897cf444613.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-apiserver-172.17.4.200_kube-system_kube-apiserver-69daa65567c6ecbab029e9b59b6c250a3fca8a2d87466e00556ce75c83892747.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-apiserver-172.17.4.200_kube-system_POD-c13258f312c01346346ff28367e30a6020dd35e313ff3d703e16e89c4ef17844.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/hostpath-provisioner_default_hostpath-provisioner-2c2bea3eb9428c1feba620356ac77466f7090f171a20ee0f774ad538d7fbe27f.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/etcd-server-172.17.4.200_kube-system_POD-7bd924b2d68cde6a1d18e9b4897b3548fa21b65b43f80c4b6f54111aa6f6244a.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-proxy-172.17.4.200_kube-system_POD-524dc9f5c249621c00e0b1399c16f21ba431ef1d2673c8440043fe95833b2a58.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kafka-0_kafka_POD-b62f1faeaf606019dbcf5b931b3e824e859e0f7e8010fa75629434565eae3d30.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/nginx-controller-zr4p2_thisisatest_POD-6329e9326a0696199bb682fcaa0e573b839a83025d35d42a71e6c14887ce7a19.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/etcd-server-172.17.4.200_kube-system_etcd-container-d72f798256a44603771a12f0e04230388e556eec2697c20e0df84cd8f732ac70.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-3_kafka_zookeeper-d525028a251b1a3710251401102dce5f83512c258a28fd16d4805367ffe40567.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/etcd-server-172.17.4.200_kube-system_etcd-container-79984eb8d800545860e73f4b7238704a9742b3a32c238528e09d643aeb651843.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-4_kafka_zookeeper-7fe9456ea40206444cc790d6e1043e7fab37e6129b78fe49cee6e39287e9a434.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kubernetes-dashboard-3543765157-87fkr_kube-system_POD-5af218f732d9753dc017e81196ce285995e97ce44f0a506b7399b2488df10c3f.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-2_kafka_POD-9cfa8315de8d971f7d0e676d3bee2a1cfac5b4197d751a6f45fae47160e03c0c.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-scheduler-172.17.4.200_kube-system_POD-f5caf98454591e0ef5da25330f83f476466db398a9491e0845a9237cf2494430.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/zoo-3_kafka_zookeeper-a060029d8d77e7046a72e4ce9a345e38e72ee1a7f1a315746707c65f4da88846.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kube-proxy-172.17.4.200_kube-system_kube-proxy-5cf46a46791435601e198409a251ce52201491553f508507666a14fa415b1d29.log
2017-05-14 09:34:30 +0000 [info]: following tail of /var/log/containers/kafka-2_kafka_broker-a43699d7e9f5b67842e02be271d60a072ef8233d0291080f426e7aca37675e51.log
2017-05-14 09:34:34 +0000 [info]: following tail of /var/log/containers/kube-controller-manager-172.17.4.200_kube-system_kube-controller-manager-e2254f940567051e5d333c7d7ccc322a830346704d41f81ab8eb366517436355.log
2017-05-14 09:34:35 +0000 [info]: following tail of /var/log/containers/kafka-2_kafka_broker-ea7a52f7ccd679320c48db104467697c36ab01ac8d23bc22677d99482fc7515a.log
2017-05-14 09:34:35 +0000 [info]: process finished code=9
2017-05-14 09:34:35 +0000 [error]: fluentd main process died unexpectedly. restarting.
2017-05-14 09:34:35 +0000 [info]: starting fluentd-0.12.29

2017-05-14 09:37:30 +0000 [info]: Connection opened to Elasticsearch cluster => {:host=>"elasticsearch-logging", :port=>9200, :scheme=>"http"}
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log po/kibana-logging-2924323056-2l63r
ELASTICSEARCH_URL=http://elasticsearch-logging:9200
server.basePath: /api/v1/proxy/namespaces/kube-system/services/kibana-logging
{"type":"log","@timestamp":"2017-05-14T08:13:50Z","tags":["info","optimize"],"pid":7,"message":"Optimizing and caching bundles for kibana and statusPage. This may take a few minutes"}
{"type":"log","@timestamp":"2017-05-14T08:21:02Z","tags":["info","optimize"],"pid":7,"message":"Optimization of bundles for kibana and statusPage complete in 431.21 seconds"}
{"type":"log","@timestamp":"2017-05-14T08:21:02Z","tags":["status","plugin:kibana@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from uninitialized to green - Ready","prevState":"uninitialized","prevMsg":"uninitialized"}
{"type":"log","@timestamp":"2017-05-14T08:21:02Z","tags":["status","plugin:elasticsearch@1.0.0","info"],"pid":7,"state":"yellow","message":"Status changed from uninitialized to yellow - Waiting for Elasticsearch","prevState":"uninitialized","prevMsg":"uninitialized"}
{"type":"log","@timestamp":"2017-05-14T08:21:03Z","tags":["status","plugin:kbn_vislib_vis_types@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from uninitialized to green - Ready","prevState":"uninitialized","prevMsg":"uninitialized"}
{"type":"log","@timestamp":"2017-05-14T08:21:03Z","tags":["status","plugin:markdown_vis@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from uninitialized to green - Ready","prevState":"uninitialized","prevMsg":"uninitialized"}
{"type":"log","@timestamp":"2017-05-14T08:21:03Z","tags":["status","plugin:metric_vis@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from uninitialized to green - Ready","prevState":"uninitialized","prevMsg":"uninitialized"}
{"type":"log","@timestamp":"2017-05-14T08:21:03Z","tags":["status","plugin:spyModes@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from uninitialized to green - Ready","prevState":"uninitialized","prevMsg":"uninitialized"}
{"type":"log","@timestamp":"2017-05-14T08:21:03Z","tags":["status","plugin:statusPage@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from uninitialized to green - Ready","prevState":"uninitialized","prevMsg":"uninitialized"}
{"type":"log","@timestamp":"2017-05-14T08:21:04Z","tags":["status","plugin:table_vis@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from uninitialized to green - Ready","prevState":"uninitialized","prevMsg":"uninitialized"}
{"type":"log","@timestamp":"2017-05-14T08:21:04Z","tags":["listening","info"],"pid":7,"message":"Server running at http://0.0.0.0:5601"}
{"type":"log","@timestamp":"2017-05-14T08:21:10Z","tags":["status","plugin:elasticsearch@1.0.0","info"],"pid":7,"state":"yellow","message":"Status changed from yellow to yellow - No existing Kibana index found","prevState":"yellow","prevMsg":"Waiting for Elasticsearch"}
{"type":"log","@timestamp":"2017-05-14T08:21:26Z","tags":["status","plugin:elasticsearch@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from yellow to green - Kibana index ready","prevState":"yellow","prevMsg":"No existing Kibana index found"}
{"type":"log","@timestamp":"2017-05-14T08:21:35Z","tags":["status","plugin:elasticsearch@1.0.0","error"],"pid":7,"state":"red","message":"Status changed from green to red - Request Timeout after 3000ms","prevState":"green","prevMsg":"Kibana index ready"}
{"type":"log","@timestamp":"2017-05-14T08:21:39Z","tags":["status","plugin:elasticsearch@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from red to green - Kibana index ready","prevState":"red","prevMsg":"Request Timeout after 3000ms"}
{"type":"log","@timestamp":"2017-05-14T08:23:24Z","tags":["status","plugin:elasticsearch@1.0.0","error"],"pid":7,"state":"red","message":"Status changed from green to red - Request Timeout after 3000ms","prevState":"green","prevMsg":"Kibana index ready"}
{"type":"log","@timestamp":"2017-05-14T08:23:30Z","tags":["status","plugin:elasticsearch@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from red to green - Kibana index ready","prevState":"red","prevMsg":"Request Timeout after 3000ms"}
{"type":"log","@timestamp":"2017-05-14T08:31:20Z","tags":["status","plugin:elasticsearch@1.0.0","error"],"pid":7,"state":"red","message":"Status changed from green to red - Request Timeout after 3000ms","prevState":"green","prevMsg":"Kibana index ready"}
{"type":"log","@timestamp":"2017-05-14T08:31:42Z","tags":["status","plugin:elasticsearch@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from red to green - Kibana index ready","prevState":"red","prevMsg":"Request Timeout after 3000ms"}
{"type":"log","@timestamp":"2017-05-14T08:31:48Z","tags":["status","plugin:elasticsearch@1.0.0","error"],"pid":7,"state":"red","message":"Status changed from green to red - Request Timeout after 3000ms","prevState":"green","prevMsg":"Kibana index ready"}
{"type":"log","@timestamp":"2017-05-14T08:31:56Z","tags":["status","plugin:elasticsearch@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from red to green - Kibana index ready","prevState":"red","prevMsg":"Request Timeout after 3000ms"}
{"type":"response","@timestamp":"2017-05-14T08:45:28Z","tags":[],"pid":7,"method":"get","statusCode":200,"req":{"url":"/","method":"get","headers":{"host":"172.17.4.200:30520","connection":"keep-alive","upgrade-insecure-requests":"1","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36","accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8","accept-encoding":"gzip, deflate, sdch","accept-language":"zh-CN,zh;q=0.8,en;q=0.6"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1"},"res":{"statusCode":200,"responseTime":1176,"contentLength":9},"message":"GET / 200 1176ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:45:30Z","tags":[],"pid":7,"method":"get","statusCode":404,"req":{"url":"/api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana","method":"get","headers":{"host":"172.17.4.200:30520","connection":"keep-alive","upgrade-insecure-requests":"1","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36","accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8","referer":"http://172.17.4.200:30520/","accept-encoding":"gzip, deflate, sdch","accept-language":"zh-CN,zh;q=0.8,en;q=0.6"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1","referer":"http://172.17.4.200:30520/"},"res":{"statusCode":404,"responseTime":194,"contentLength":9},"message":"GET /api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana 404 194ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:45:30Z","tags":[],"pid":7,"method":"get","statusCode":404,"req":{"url":"/favicon.ico","method":"get","headers":{"host":"172.17.4.200:30520","connection":"keep-alive","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36","accept":"image/webp,image/*,*/*;q=0.8","referer":"http://172.17.4.200:30520/","accept-encoding":"gzip, deflate, sdch","accept-language":"zh-CN,zh;q=0.8,en;q=0.6"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1","referer":"http://172.17.4.200:30520/"},"res":{"statusCode":404,"responseTime":204,"contentLength":9},"message":"GET /favicon.ico 404 204ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:45:30Z","tags":[],"pid":7,"method":"get","statusCode":404,"req":{"url":"/favicon.ico","method":"get","headers":{"host":"172.17.4.200:30520","connection":"keep-alive","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36","accept":"image/webp,image/*,*/*;q=0.8","referer":"http://172.17.4.200:30520/api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana","accept-encoding":"gzip, deflate, sdch","accept-language":"zh-CN,zh;q=0.8,en;q=0.6"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1","referer":"http://172.17.4.200:30520/api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana"},"res":{"statusCode":404,"responseTime":95,"contentLength":9},"message":"GET /favicon.ico 404 95ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:46:05Z","tags":[],"pid":7,"method":"get","statusCode":200,"req":{"url":"/","method":"get","headers":{"user-agent":"curl/7.35.0","host":"10.123.249.244:5601","accept":"*/*"},"remoteAddress":"10.0.2.15","userAgent":"10.0.2.15"},"res":{"statusCode":200,"responseTime":5,"contentLength":9},"message":"GET / 200 5ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:47:11Z","tags":[],"pid":7,"method":"get","statusCode":200,"req":{"url":"/","method":"get","headers":{"user-agent":"curl/7.35.0","host":"10.123.240.212:5601","accept":"*/*"},"remoteAddress":"10.0.2.15","userAgent":"10.0.2.15"},"res":{"statusCode":200,"responseTime":3,"contentLength":9},"message":"GET / 200 3ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:47:56Z","tags":[],"pid":7,"method":"get","statusCode":200,"req":{"url":"/","method":"get","headers":{"user-agent":"curl/7.35.0","host":"172.17.4.200:30520","accept":"*/*"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1"},"res":{"statusCode":200,"responseTime":4,"contentLength":9},"message":"GET / 200 4ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:48:59Z","tags":[],"pid":7,"method":"get","statusCode":404,"req":{"url":"/api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana","method":"get","headers":{"host":"172.17.4.200:30520","connection":"keep-alive","cache-control":"max-age=0","upgrade-insecure-requests":"1","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36","accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8","referer":"http://172.17.4.200:30520/","accept-encoding":"gzip, deflate, sdch","accept-language":"zh-CN,zh;q=0.8,en;q=0.6"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1","referer":"http://172.17.4.200:30520/"},"res":{"statusCode":404,"responseTime":21,"contentLength":9},"message":"GET /api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana 404 21ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:49:03Z","tags":[],"pid":7,"method":"get","statusCode":404,"req":{"url":"/api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana","method":"get","headers":{"host":"172.17.4.200:30520","connection":"keep-alive","cache-control":"max-age=0","upgrade-insecure-requests":"1","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36","accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8","referer":"http://172.17.4.200:30520/","accept-encoding":"gzip, deflate, sdch","accept-language":"zh-CN,zh;q=0.8,en;q=0.6"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1","referer":"http://172.17.4.200:30520/"},"res":{"statusCode":404,"responseTime":3,"contentLength":9},"message":"GET /api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana 404 3ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:49:04Z","tags":[],"pid":7,"method":"get","statusCode":404,"req":{"url":"/api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana","method":"get","headers":{"host":"172.17.4.200:30520","connection":"keep-alive","cache-control":"max-age=0","upgrade-insecure-requests":"1","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36","accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8","referer":"http://172.17.4.200:30520/","accept-encoding":"gzip, deflate, sdch","accept-language":"zh-CN,zh;q=0.8,en;q=0.6"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1","referer":"http://172.17.4.200:30520/"},"res":{"statusCode":404,"responseTime":8,"contentLength":9},"message":"GET /api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana 404 8ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:49:13Z","tags":[],"pid":7,"method":"get","statusCode":200,"req":{"url":"/","method":"get","headers":{"host":"172.17.4.200:30520","connection":"keep-alive","upgrade-insecure-requests":"1","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36","accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8","accept-encoding":"gzip, deflate, sdch","accept-language":"zh-CN,zh;q=0.8,en;q=0.6"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1"},"res":{"statusCode":200,"responseTime":4,"contentLength":9},"message":"GET / 200 4ms - 9.0B"}
{"type":"response","@timestamp":"2017-05-14T08:49:13Z","tags":[],"pid":7,"method":"get","statusCode":404,"req":{"url":"/api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana","method":"get","headers":{"host":"172.17.4.200:30520","connection":"keep-alive","upgrade-insecure-requests":"1","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36","accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8","referer":"http://172.17.4.200:30520/","accept-encoding":"gzip, deflate, sdch","accept-language":"zh-CN,zh;q=0.8,en;q=0.6"},"remoteAddress":"172.18.0.1","userAgent":"172.18.0.1","referer":"http://172.17.4.200:30520/"},"res":{"statusCode":404,"responseTime":8,"contentLength":9},"message":"GET /api/v1/proxy/namespaces/kube-system/services/kibana-logging/app/kibana 404 8ms - 9.0B"}
{"type":"log","@timestamp":"2017-05-14T09:35:39Z","tags":["status","plugin:elasticsearch@1.0.0","error"],"pid":7,"state":"red","message":"Status changed from green to red - Request Timeout after 30000ms","prevState":"green","prevMsg":"Kibana index ready"}
{"type":"log","@timestamp":"2017-05-14T09:36:15Z","tags":["status","plugin:elasticsearch@1.0.0","info"],"pid":7,"state":"green","message":"Status changed from red to green - Kibana index ready","prevState":"red","prevMsg":"Request Timeout after 30000ms"}
```

For browsing Kibana
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system expose service kibana-logging --name=kibana-logging-ex --target-port=5601 --type=NodePort
service "kibana-logging-ex" exposed
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get svc/kibana-logging-ex
NAME                CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
kibana-logging-ex   10.123.249.230   <nodes>       5601:31224/TCP   10s
```

### V1.6+

Pull
```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/elasticsearch:v2.4.1-2
v2.4.1-2: Pulling from google_containers/elasticsearch
5040bd298390: Pull complete 
fce5728aad85: Pull complete 
c42794440453: Pull complete 
0c0da797ba48: Pull complete 
7c9b17433752: Pull complete 
114e02586e63: Pull complete 
e4c663802e9a: Pull complete 
96c525834929: Pull complete 
f001cc4e82e1: Pull complete 
24af501b4bb7: Pull complete 
22158bcd99c5: Pull complete 
a2c240bdd3d1: Pull complete 
18c7b4bd2d8e: Pull complete 
10b1a9e02172: Pull complete 
064c2264dfef: Pull complete 
Digest: sha256:5dbeae2505ff4e64d9feb4b91bc8d6dccad014be3f74b5f9bbc2977d3ac3494b
Status: Downloaded newer image for gcr.io/google_containers/elasticsearch:v2.4.1-2
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull gcr.io/google_containers/kibana:v4.6.1-1
v4.6.1-1: Pulling from google_containers/kibana
be4ee6c45489: Already exists 
f2e515c7820d: Pull complete 
b9b055361fab: Pull complete 
1baf5972be74: Pull complete 
Digest: sha256:73298cb0ee77e5a2707887ee98f81826dfe480f2ddf35864a80170ce5a0791a6
Status: Downloaded newer image for gcr.io/google_containers/kibana:v4.6.1-1
```

Deploy
```
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl create -f workspace/src/k8s.io/kubernetes/cluster/addons/fluentd-elasticsearch/
replicationcontroller "elasticsearch-logging-v1" created
service "elasticsearch-logging" created
daemonset "fluentd-es-v1.22" created
deployment "kibana-logging" created
service "kibana-logging" created

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system delete svc/kibana-logging-ex
service "kibana-logging-ex" deleted
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system expose service kibana-logging --name=kibana-logging-ex --target-port=5601 --type=NodePort
service "kibana-logging-ex" exposed
```

Validation
```
vagrant@vagrant-ubuntu-trusty-64:~$ curl 10.123.252.1:9200
{
  "name" : "elasticsearch-logging-v1-hz7tk",
  "cluster_name" : "kubernetes-logging",
  "cluster_uuid" : "lN67UzdTR_2Wb3j8S7SG2A",
  "version" : {
    "number" : "2.4.1",
    "build_hash" : "c67dc32e24162035d18d6fe1e952c4cbcbe79d16",
    "build_timestamp" : "2016-09-27T18:57:55Z",
    "build_snapshot" : false,
    "lucene_version" : "5.5.2"
  },
  "tagline" : "You Know, for Search"
}
vagrant@vagrant-ubuntu-trusty-64:~$ curl 10.123.252.1:9200/_cluster/health
{"cluster_name":"kubernetes-logging","status":"green","timed_out":false,"number_of_nodes":2,"number_of_data_nodes":2,"active_primary_shards":0,"active_shards":0,"relocating_shards":0,"initializing_shards":0,"unassigned_shards":0,"delayed_unassigned_shards":0,"number_of_pending_tasks":0,"number_of_in_flight_fetch":0,"task_max_waiting_in_queue_millis":0,"active_shards_percent_as_number":100.0}vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ curl 10.123.252.1:9200/_cluster/state
{"cluster_name":"kubernetes-logging","version":2,"state_uuid":"HeG0UEiESLmGXVpLL7V8yA","master_node":"Os2phPq5RICYQ5irk1QshA","blocks":{},"nodes":{"tseOZ-0tTEiVEDwcQBlX9A":{"name":"elasticsearch-logging-v1-8n3pj","transport_address":"172.18.0.7:9300","attributes":{"master":"true"}},"Os2phPq5RICYQ5irk1QshA":{"name":"elasticsearch-logging-v1-hz7tk","transport_address":"172.18.0.5:9300","attributes":{"master":"true"}}},"metadata":{"cluster_uuid":"lN67UzdTR_2Wb3j8S7SG2A","templates":{},"indices":{}},"routing_table":{"indices":{}},"routing_nodes":{"unassigned":[],"nodes":{"Os2phPq5RICYQ5irk1QshA":[],"tseOZ-0tTEiVEDwcQBlX9A":[]}}}vagrant@vagrant-ubuntu-trusty-64:~$ 
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl get po -l k8s-app=elasticsearch-logging
No resources found.
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get po -l k8s-app=elasticsearch-logging
NAME                             READY     STATUS    RESTARTS   AGE
elasticsearch-logging-v1-8n3pj   1/1       Running   0          2m
elasticsearch-logging-v1-hz7tk   1/1       Running   0          2m
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log elasticsearch-logging-v1-8n3pj
[2017-05-14 12:47:37,072][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] version[2.4.1], pid[1], build[c67dc32/2016-09-27T18:57:55Z]
[2017-05-14 12:47:37,076][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] initializing ...
[2017-05-14 12:47:38,253][INFO ][plugins                  ] [elasticsearch-logging-v1-8n3pj] modules [reindex, lang-expression, lang-groovy], plugins [], sites []
[2017-05-14 12:47:38,306][INFO ][env                      ] [elasticsearch-logging-v1-8n3pj] using [1] data paths, mounts [[/data (/dev/sda1)]], net usable_space [6.9gb], net total_space [39.3gb], spins? [possibly], types [ext4]
[2017-05-14 12:47:38,307][INFO ][env                      ] [elasticsearch-logging-v1-8n3pj] heap size [1007.3mb], compressed ordinary object pointers [true]
[2017-05-14 12:47:44,906][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] initialized
[2017-05-14 12:47:44,909][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] starting ...
[2017-05-14 12:47:45,216][INFO ][transport                ] [elasticsearch-logging-v1-8n3pj] publish_address {172.18.0.7:9300}, bound_addresses {[::]:9300}
[2017-05-14 12:47:45,236][INFO ][discovery                ] [elasticsearch-logging-v1-8n3pj] kubernetes-logging/tseOZ-0tTEiVEDwcQBlX9A
[2017-05-14 12:47:48,477][INFO ][cluster.service          ] [elasticsearch-logging-v1-8n3pj] detected_master {elasticsearch-logging-v1-hz7tk}{Os2phPq5RICYQ5irk1QshA}{172.18.0.5}{172.18.0.5:9300}{master=true}, added {{elasticsearch-logging-v1-hz7tk}{Os2phPq5RICYQ5irk1QshA}{172.18.0.5}{172.18.0.5:9300}{master=true},}, reason: zen-disco-receive(from master [{elasticsearch-logging-v1-hz7tk}{Os2phPq5RICYQ5irk1QshA}{172.18.0.5}{172.18.0.5:9300}{master=true}])
[2017-05-14 12:47:48,513][INFO ][http                     ] [elasticsearch-logging-v1-8n3pj] publish_address {172.18.0.7:9200}, bound_addresses {[::]:9200}
[2017-05-14 12:47:48,513][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] started

vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get po -l k8s-app=elasticsearch-logging
NAME                             READY     STATUS    RESTARTS   AGE
elasticsearch-logging-v1-8n3pj   1/1       Running   0          3m
elasticsearch-logging-v1-hz7tk   1/1       Running   0          3m
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log elasticsearch-logging-v1-hz7tk
[2017-05-14 12:47:36,893][INFO ][node                     ] [elasticsearch-logging-v1-hz7tk] version[2.4.1], pid[1], build[c67dc32/2016-09-27T18:57:55Z]
[2017-05-14 12:47:36,897][INFO ][node                     ] [elasticsearch-logging-v1-hz7tk] initializing ...
[2017-05-14 12:47:38,147][INFO ][plugins                  ] [elasticsearch-logging-v1-hz7tk] modules [reindex, lang-expression, lang-groovy], plugins [], sites []
[2017-05-14 12:47:38,196][INFO ][env                      ] [elasticsearch-logging-v1-hz7tk] using [1] data paths, mounts [[/data (/dev/sda1)]], net usable_space [6.9gb], net total_space [39.3gb], spins? [possibly], types [ext4]
[2017-05-14 12:47:38,200][INFO ][env                      ] [elasticsearch-logging-v1-hz7tk] heap size [1007.3mb], compressed ordinary object pointers [true]
[2017-05-14 12:47:44,122][INFO ][node                     ] [elasticsearch-logging-v1-hz7tk] initialized
[2017-05-14 12:47:44,124][INFO ][node                     ] [elasticsearch-logging-v1-hz7tk] starting ...
[2017-05-14 12:47:44,396][INFO ][transport                ] [elasticsearch-logging-v1-hz7tk] publish_address {172.18.0.5:9300}, bound_addresses {[::]:9300}
[2017-05-14 12:47:44,423][INFO ][discovery                ] [elasticsearch-logging-v1-hz7tk] kubernetes-logging/Os2phPq5RICYQ5irk1QshA
[2017-05-14 12:47:48,436][INFO ][cluster.service          ] [elasticsearch-logging-v1-hz7tk] new_master {elasticsearch-logging-v1-hz7tk}{Os2phPq5RICYQ5irk1QshA}{172.18.0.5}{172.18.0.5:9300}{master=true}, added {{elasticsearch-logging-v1-8n3pj}{tseOZ-0tTEiVEDwcQBlX9A}{172.18.0.7}{172.18.0.7:9300}{master=true},}, reason: zen-disco-join(elected_as_master, [1] joins received)
[2017-05-14 12:47:48,578][INFO ][http                     ] [elasticsearch-logging-v1-hz7tk] publish_address {172.18.0.5:9200}, bound_addresses {[::]:9200}
[2017-05-14 12:47:48,578][INFO ][node                     ] [elasticsearch-logging-v1-hz7tk] started
[2017-05-14 12:47:48,643][INFO ][gateway                  ] [elasticsearch-logging-v1-hz7tk] recovered [0] indices into cluster_state
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system log elasticsearch-logging-v1-8n3pj
[2017-05-14 12:47:37,072][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] version[2.4.1], pid[1], build[c67dc32/2016-09-27T18:57:55Z]
[2017-05-14 12:47:37,076][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] initializing ...
[2017-05-14 12:47:38,253][INFO ][plugins                  ] [elasticsearch-logging-v1-8n3pj] modules [reindex, lang-expression, lang-groovy], plugins [], sites []
[2017-05-14 12:47:38,306][INFO ][env                      ] [elasticsearch-logging-v1-8n3pj] using [1] data paths, mounts [[/data (/dev/sda1)]], net usable_space [6.9gb], net total_space [39.3gb], spins? [possibly], types [ext4]
[2017-05-14 12:47:38,307][INFO ][env                      ] [elasticsearch-logging-v1-8n3pj] heap size [1007.3mb], compressed ordinary object pointers [true]
[2017-05-14 12:47:44,906][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] initialized
[2017-05-14 12:47:44,909][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] starting ...
[2017-05-14 12:47:45,216][INFO ][transport                ] [elasticsearch-logging-v1-8n3pj] publish_address {172.18.0.7:9300}, bound_addresses {[::]:9300}
[2017-05-14 12:47:45,236][INFO ][discovery                ] [elasticsearch-logging-v1-8n3pj] kubernetes-logging/tseOZ-0tTEiVEDwcQBlX9A
[2017-05-14 12:47:48,477][INFO ][cluster.service          ] [elasticsearch-logging-v1-8n3pj] detected_master {elasticsearch-logging-v1-hz7tk}{Os2phPq5RICYQ5irk1QshA}{172.18.0.5}{172.18.0.5:9300}{master=true}, added {{elasticsearch-logging-v1-hz7tk}{Os2phPq5RICYQ5irk1QshA}{172.18.0.5}{172.18.0.5:9300}{master=true},}, reason: zen-disco-receive(from master [{elasticsearch-logging-v1-hz7tk}{Os2phPq5RICYQ5irk1QshA}{172.18.0.5}{172.18.0.5:9300}{master=true}])
[2017-05-14 12:47:48,513][INFO ][http                     ] [elasticsearch-logging-v1-8n3pj] publish_address {172.18.0.7:9200}, bound_addresses {[::]:9200}
[2017-05-14 12:47:48,513][INFO ][node                     ] [elasticsearch-logging-v1-8n3pj] started

vagrant@vagrant-ubuntu-trusty-64:~$ sed 's/200Mi/100Mi/g' workspace/src/k8s.io/kubernetes/cluster/addons/fluentd-elasticsearch/fluentd-es-ds.yaml | sudo tee /etc/kubernetes/manifests/addons/fluentd-es-ds.yaml
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: fluentd-es-v1.22
  namespace: kube-system
  labels:
    k8s-app: fluentd-es
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    version: v1.22
spec:
  template:
    metadata:
      labels:
        k8s-app: fluentd-es
        kubernetes.io/cluster-service: "true"
        version: v1.22
      # This annotation ensures that fluentd does not get evicted if the node
      # supports critical pod annotation based priority scheme.
      # Note that this does not guarantee admission on the nodes (#40573).
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      containers:
      - name: fluentd-es
        image: gcr.io/google_containers/fluentd-elasticsearch:1.22
        command:
          - '/bin/sh'
          - '-c'
          - '/usr/sbin/td-agent 2>&1 >> /var/log/fluentd.log'
        resources:
          limits:
            memory: 100Mi
          requests:
            cpu: 100m
            memory: 100Mi
        volumeMounts:
        - name: varlog
          mountPath: /var/log
        - name: varlibdockercontainers
          mountPath: /var/lib/docker/containers
          readOnly: true
      nodeSelector:
        beta.kubernetes.io/fluentd-ds-ready: "true"
      terminationGracePeriodSeconds: 30
      volumes:
      - name: varlog
        hostPath:
          path: /var/log
      - name: varlibdockercontainers
        hostPath:
          path: /var/lib/docker/containers
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl create -f /etc/kubernetes/manifests/addons/fluentd-es-ds.yaml 
daemonset "fluentd-es-v1.22" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get pods
NAME                                    READY     STATUS    RESTARTS   AGE
elasticsearch-logging-v1-8n3pj          1/1       Running   0          13m
elasticsearch-logging-v1-hz7tk          1/1       Running   0          13m
etcd-server-172.17.4.200                1/1       Running   9          9d
kibana-logging-1918083406-x5p53         1/1       Running   0          13m
kube-apiserver-172.17.4.200             1/1       Running   29         9d
kube-controller-manager-172.17.4.200    1/1       Running   13         9d
kube-dns-2185667875-gkjlh               4/4       Running   39         10d
kube-proxy-172.17.4.200                 1/1       Running   6          9d
kube-scheduler-172.17.4.200             1/1       Running   12         9d
kubernetes-dashboard-3543765157-87fkr   1/1       Running   6          9d
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get ds
NAME               DESIRED   CURRENT   READY     NODE-SELECTOR                              AGE
fluentd-es-v1.22   0         0         0         beta.kubernetes.io/fluentd-ds-ready=true   22s
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system delete -f workspace/src/k8s.io/kubernetes/cluster/addons/fluentd-elasticsearch/
replicationcontroller "elasticsearch-logging-v1" deleted
service "elasticsearch-logging" deleted
daemonset "fluentd-es-v1.22" deleted
deployment "kibana-logging" deleted
service "kibana-logging" deleted
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system create -f /opt/kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/
replicationcontroller "elasticsearch-logging-v1" created
service "elasticsearch-logging" created
deployment "kibana-logging" created
service "kibana-logging" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system create -f /opt/kubernetes/saltbase/salt/fluentd-es/
pod "fluentd-elasticsearch" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/v1.3.7/cluster/addons/fluentd-elasticsearch/kibana-controller.yaml -o /etc/kubernetes/manifests/addons/kibana-controller.yaml 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   893  100   893    0     0    152      0  0:00:05  0:00:05 --:--:--   205
vagrant@vagrant-ubuntu-trusty-64:~$ sudo curl -jkSL https://raw.githubusercontent.com/kubernetes/kubernetes/v1.3.7/cluster/addons/fluentd-elasticsearch/kibana-service.yaml -o /etc/kubernetes/manifests/addons/kibana-service.yaml 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   307  100   307    0     0    255      0  0:00:01  0:00:01 --:--:--   255
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system delete -f /opt/kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/kibana-service.yaml -f /opt/kubernetes/saltbase/salt/kube-addons/fluentd-elasticsearch/kibana-controller.yaml
service "kibana-logging" deleted
deployment "kibana-logging" deleted
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system create -f /etc/kubernetes/manifests/addons/ 
replicationcontroller "kibana-logging-v1" created
service "kibana-logging" created
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get all -l k8s-app=kibana-logging
NAME                         READY     STATUS    RESTARTS   AGE
po/kibana-logging-v1-zgbsc   1/1       Running   0          29s

NAME                   DESIRED   CURRENT   READY     AGE
rc/kibana-logging-v1   1         1         1         29s

NAME                    CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
svc/kibana-logging      10.123.251.25    <none>        5601/TCP         29s
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system expose service kibana-logging --name=kibana-logging-ex --target-port=5601 --type=NodePort
service "kibana-logging-ex" exposed
vagrant@vagrant-ubuntu-trusty-64:~$ sudo /opt/kubernetes/server/bin/kubectl --namespace=kube-system get service -l k8s-app=kibana-logging
NAME                CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kibana-logging      10.123.251.25   <none>        5601/TCP         4m
kibana-logging-ex   10.123.245.63   <nodes>       5601:30288/TCP   32s
```

### Links

[使用Fluentd和ElasticSearch Stack实现Kubernetes的集群Logging](http://tonybai.com/2017/03/03/implement-kubernetes-cluster-level-logging-with-fluentd-and-elasticsearch-stack/) - Tony Bai