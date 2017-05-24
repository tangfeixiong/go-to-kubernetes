Monitoring Apache Kafak with Prometheus and Grafana
===================================================

Inspried from https://github.com/rama-nallamilli/kafka-prometheus-monitoring

Table of contents
-----------------

* Download 
* Build docker image
* Deploy 
* Test

Content
-------

Download
```
fanhonglingdeMacBook-Pro:github.com fanhongling$ cd tangfeixiong/go-to-kubernetes/examples/kafka/
fanhonglingdeMacBook-Pro:kafka fanhongling$ mkdir -p https0x3A0x2F0x2Fraw.githubusercontent.com0x2Frama-nallamilli0x2Fkafka-prometheus-monitoring0x2Fmaster/prometheus-jmx-exporter
fanhonglingdeMacBook-Pro:kafka fanhongling$ cd https0x3A0x2F0x2Fraw.githubusercontent.com0x2Frama-nallamilli0x2Fkafka-prometheus-monitoring0x2Fmaster/prometheus-jmx-exporter
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ curl -jkSL https://raw.githubusercontent.com/rama-nallamilli/kafka-prometheus-monitoring/master/prometheus-jmx-exporter/Dockerfile | tee Dockerfile
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   510  100   510    0     0     73      0  0:00:06  0:00:06 --:--:--   114
FROM java:8

RUN mkdir /opt/jmx_prometheus_httpserver && wget 'http://central.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_httpserver/0.6/jmx_prometheus_httpserver-0.6-jar-with-dependencies.jar' -O /opt/jmx_prometheus_httpserver/jmx_prometheus_httpserver.jar

ADD https://github.com/kelseyhightower/confd/releases/download/v0.11.0/confd-0.11.0-linux-amd64 /usr/local/bin/confd
COPY confd /etc/confd
RUN chmod +x /usr/local/bin/confd

COPY entrypoint.sh /opt/entrypoint.sh
ENTRYPOINT ["/opt/entrypoint.sh"]
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ curl -jkSL https://raw.githubusercontent.com/rama-nalng/master/prometheus-jmx-exporter/entrypoint.sh | tee entrypoint.sh
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    80  100    80    0     0     28      0  0:00:02  0:00:02 --:--:--    28
#!/bin/bash
/usr/local/bin/confd -onetime -backend env
/opt/start-jmx-scraper.sh
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ mkdir confd/conf.d
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ curl -jkSL https://raw.githubusercontent.com/rama-nallamilli/kafka-prometheus-monitoring/master/prometheus-jmx-exporter/confd/conf.d/kafka.yml.toml | tee confd/conf.d/kafka.yml.toml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   123  100   123    0     0     17      0  0:00:07  0:00:06  0:00:01    28
[template]
src   = "kafka.yml.tmpl"
dest = "/opt/jmx_prometheus_httpserver/kafka.yml"
keys = [
 "/jmx/host",
 "/jmx/port"
]
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ curl -jkSL https://raw.githubusercontent.com/rama-nallamilli/kafka-prometheus-monitorng/master/prometheus-jmx-exporter/confd/conf.d/start-jmx-scraper.sh.toml | tee confd/conf.d/start-jmx-scraper.sh.toml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   178  100   178    0     0    141      0  0:00:01  0:00:01 --:--:--   141
[template]
src   = "start-jmx-scraper.sh.tmpl"
dest = "/opt/start-jmx-scraper.sh"
mode = "0550"
keys = [
 "/jmx/host",
 "/jmx/port",
 "/http/port",
 "/jmx/exporter/config/file"
]
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$mkdir confd/templates
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ curl -jkSL https://raw.githubusercontent.com/rama-nallamilli/kafka-prometheus-monitoring/master/prometheus-jmx-exporter/confd/templates/kafka.yml.tmpl | tee confd/templates/kafka.yml.tmpl
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1537  100  1537    0     0    127      0  0:00:12  0:00:12 --:--:--   420
lowercaseOutputName: true
jmxUrl: service:jmx:rmi:///jndi/rmi://{{ getv "/jmx/host" }}:{{ getv "/jmx/port" }}/jmxrmi
rules:
- pattern : kafka.network<type=Processor, name=IdlePercent, networkProcessor=(.+)><>Value
- pattern : kafka.network<type=RequestMetrics, name=RequestsPerSec, request=(.+)><>OneMinuteRate
- pattern : kafka.network<type=SocketServer, name=NetworkProcessorAvgIdlePercent><>Value
- pattern : kafka.server<type=ReplicaFetcherManager, name=MaxLag, clientId=(.+)><>Value
- pattern : kafka.server<type=BrokerTopicMetrics, name=(.+), topic=(.+)><>OneMinuteRate
- pattern : kafka.server<type=KafkaRequestHandlerPool, name=RequestHandlerAvgIdlePercent><>OneMinuteRate
- pattern : kafka.server<type=Produce><>queue-size
- pattern : kafka.server<type=ReplicaManager, name=(.+)><>(Value|OneMinuteRate)
- pattern : kafka.server<type=controller-channel-metrics, broker-id=(.+)><>(.*)
- pattern : kafka.server<type=socket-server-metrics, networkProcessor=(.+)><>(.*)
- pattern : kafka.server<type=Fetch><>queue-size
- pattern : kafka.server<type=SessionExpireListener, name=(.+)><>OneMinuteRate
- pattern : kafka.controller<type=KafkaController, name=(.+)><>Value
- pattern : kafka.controller<type=ControllerStats, name=(.+)><>OneMinuteRate
- pattern : kafka.cluster<type=Partition, name=UnderReplicated, topic=(.+), partition=(.+)><>Value
- pattern : kafka.utils<type=Throttler, name=cleaner-io><>OneMinuteRate
- pattern : kafka.log<type=Log, name=LogEndOffset, topic=(.+), partition=(.+)><>Value
- pattern : java.lang<type=(.*)>fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ 
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ curl -jkSL https://raw.githubusercontent.com/rama-nallamilli/kafka-prometheus-monitoring/master/prometheus-jmx-exporter/confd/templates/start-jmx-scraper.sh.tmpl | tee confd/templates/start-jmx-scraper.sh.tmpl
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   376  100   376    0     0     62      0  0:00:06  0:00:06 --:--:--   108
#!/bin/bash
java -Dcom.sun.management.jmxremote.ssl=false -Djava.rmi.server.hostname={{ getv "/jmx/host" }} -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.port={{ getv "/jmx/port" }} -jar /opt/jmx_prometheus_httpserver/jmx_prometheus_httpserver.jar {{ getv "/http/port" }} /opt/jmx_prometheus_httpserver/{{ getv "/jmx/exporter/config/file" }}

fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ cd ..
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ mkdir -p mount/prometheus
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Frama-nallamilli0x2Fkafka-prometheus-monitoring0x2Fmaster fanhongling$ curl -jkSL https://raw.githubusercontent.com/rama-nallamilli/kafka-prometheus-monitoring/master/mount/prometheus/prometheus.yml | tee mount/prometheus/prometheus.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   437  100   437    0     0     63      0  0:00:06  0:00:06 --:--:--   100
global:
  scrape_interval:     10s # By default, scrape targets every 15 seconds.
  evaluation_interval: 10s # By default, scrape targets every 15 seconds.
  # scrape_timeout is set to the global default (10s).

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  - job_name: 'kafka'
    scrape_interval: 10s
    target_groups:
      - targets: ['kafka-jmx-exporter:8080']
```

Build docker image
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Frama-nallamilli0x2Fkafka-prometheus-monitoring0x2Fmaster/prometheus-jmx-exporter$ docker build -t tangfeixiong/prometheus-jmx-exporter .
Sending build context to Docker daemon  11.2 MB
Step 1 : FROM java:8-jre
 ---> e44d62cf8862
Step 2 : RUN mkdir /opt/jmx_prometheus_httpserver   && wget 'http://central.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_httpserver/0.6/jmx_prometheus_httpserver-0.6-jar-with-dependencies.jar'       -O /opt/jmx_prometheus_httpserver/jmx_prometheus_httpserver.jar   && curl -jkSL https://github.com/kelseyhightower/confd/releases/download/v0.11.0/confd-0.11.0-linux-amd64 -o /usr/local/bin/confd   && chmod +x /usr/local/bin/confd
 ---> Running in c62285c9630a
--2017-05-19 17:20:42--  http://central.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_httpserver/0.6/jmx_prometheus_httpserver-0.6-jar-with-dependencies.jar
Resolving central.maven.org (central.maven.org)... 151.101.40.209
Connecting to central.maven.org (central.maven.org)|151.101.40.209|:80... connected.
HTTP request sent, awaiting response... 200 OK
Length: 1516258 (1.4M) [application/java-archive]
Saving to: ‘/opt/jmx_prometheus_httpserver/jmx_prometheus_httpserver.jar’

     0K .......... .......... .......... .......... ..........  3%  152K 9s
    50K .......... .......... .......... .......... ..........  6%  335K 7s
   100K .......... .......... .......... .......... .......... 10%  766K 5s
   150K .......... .......... .......... .......... .......... 13%  919K 4s
   200K .......... .......... .......... .......... .......... 16%  679K 3s
   250K .......... .......... .......... .......... .......... 20%  891K 3s
   300K .......... .......... .......... .......... .......... 23%  499K 3s
   350K .......... .......... .......... .......... .......... 27% 18.8M 2s
   400K .......... .......... .......... .......... .......... 30% 3.28M 2s
   450K .......... .......... .......... .......... .......... 33% 1.39M 2s
   500K .......... .......... .......... .......... .......... 37%  606K 2s
   550K .......... .......... .......... .......... .......... 40% 3.93M 1s
   600K .......... .......... .......... .......... .......... 43% 2.93M 1s
   650K .......... .......... .......... .......... .......... 47% 1.95M 1s
   700K .......... .......... .......... .......... .......... 50% 1.68M 1s
   750K .......... .......... .......... .......... .......... 54%  648K 1s
   800K .......... .......... .......... .......... .......... 57% 4.58M 1s
   850K .......... .......... .......... .......... .......... 60% 6.36M 1s
   900K .......... .......... .......... .......... .......... 64% 5.57M 1s
   950K .......... .......... .......... .......... .......... 67% 6.09M 1s
  1000K .......... .......... .......... .......... .......... 70% 4.48M 0s
  1050K .......... .......... .......... .......... .......... 74% 9.08M 0s
  1100K .......... .......... .......... .......... .......... 77% 1.30M 0s
  1150K .......... .......... .......... .......... .......... 81%  711K 0s
  1200K .......... .......... .......... .......... .......... 84% 5.70M 0s
  1250K .......... .......... .......... .......... .......... 87% 6.48M 0s
  1300K .......... .......... .......... .......... .......... 91% 5.25M 0s
  1350K .......... .......... .......... .......... .......... 94% 5.27M 0s
  1400K .......... .......... .......... .......... .......... 97% 5.89M 0s
  1450K .......... .......... ..........                      100% 12.1M=1.3s

2017-05-19 17:20:47 (1.09 MB/s) - ‘/opt/jmx_prometheus_httpserver/jmx_prometheus_httpserver.jar’ saved [1516258/1516258]

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   611    0   611    0     0    173      0 --:--:--  0:00:03 --:--:--   173
100 9443k  100 9443k    0     0   348k      0  0:00:27  0:00:27 --:--:--  787k
 ---> 62ef6b30698e
Removing intermediate container c62285c9630a
Step 3 : COPY confd /etc/confd
 ---> 2b721a0e5753
Removing intermediate container 8a21ba340c99
Step 4 : COPY entrypoint.sh /opt/entrypoint.sh
 ---> 955b037bd0f2
Removing intermediate container 341eaab03790
Step 5 : ENTRYPOINT /opt/entrypoint.sh
 ---> Running in b1c14d4d21bf
 ---> 4ea0ca25341e
Removing intermediate container b1c14d4d21bf
Successfully built 4ea0ca25341e
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Frama-nallamilli0x2Fkafka-prometheus-monitoring0x2Fmaster/prometheus-jmx-exporter$ docker images tangfeixiong/prometheus-jmx-exporter
REPOSITORY                                            TAG                                IMAGE ID            CREATED             SIZE
tangfeixiong/prometheus-jmx-exporter                  latest                             4ea0ca25341e        24 seconds ago      322.4 MB
```

Optional download binaries
```
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ curl -jkSL http://central.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_httpserver/0.6/jmx_prometheus_httpserver-0.6-jar-with-dependencies.jar -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 1480k  100 1480k    0     0  18408      0  0:01:22  0:01:22 --:--:-- 43987
fanhonglingdeMacBook-Pro:prometheus-jmx-exporter fanhongling$ curl -jkSL https://github.com/kelseyhightower/confd/releases/download/v0.11.0/confd-0.11.0-linux-amd64 -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   611    0   611    0     0     89      0 --:--:--  0:00:06 --:--:--   184
100 9443k  100 9443k    0     0  54992      0  0:02:55  0:02:55 --:--:-- 76682
```

Modify __prometheus.yml__ into __etc/prometheus/__

Deploying
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ kubectl create -f prometheus-jmx-exporter-0-pod.yaml -f prometheus-jmx-exporter-0-service.yaml 
pod "prometheus-jmx-exporter-0" created
service "prometheus-jmx-exporter-0" created

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ kubectl get pod,service -l app=prometheus-jmx-exporter-0 -o wide
NAME                           READY     STATUS    RESTARTS   AGE       IP            NODE
po/prometheus-jmx-exporter-0   1/1       Running   0          29m       172.18.0.20   172.17.4.200

NAME                            CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE       SELECTOR
svc/prometheus-jmx-exporter-0   None         <none>        8080/TCP   29m       app=prometheus-jmx-exporter-0

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl 172.18.0.20:8080/metrics
# HELP kafka_utils_throttler_oneminuterate Attribute exposed for management (kafka.utils<type=Throttler, name=cleaner-io><>OneMinuteRate)
# TYPE kafka_utils_throttler_oneminuterate gauge
kafka_utils_throttler_oneminuterate{name="cleaner-io",} 0.0
# HELP kafka_server_socket_server_metrics_io_wait_time_ns_avg The average length of time the I/O thread spent waiting for a socket ready for reads or writes in nanoseconds. (kafka.server<type=socket-server-metrics, networkProcessor=0><>io-wait-time-ns-avg)
# TYPE kafka_server_socket_server_metrics_io_wait_time_ns_avg gauge
kafka_server_socket_server_metrics_io_wait_time_ns_avg{networkProcessor="0",} 3.006861860089286E8
kafka_server_socket_server_metrics_io_wait_time_ns_avg{networkProcessor="1",} 1.2556842470982143E8
kafka_server_socket_server_metrics_io_wait_time_ns_avg{networkProcessor="2",} 1.2567337756578948E8
# HELP kafka_server_controller_channel_metrics_incoming_byte_rate Bytes/second read off all sockets (kafka.server<type=controller-channel-metrics, broker-id=0><>incoming-byte-rate)
# TYPE kafka_server_controller_channel_metrics_incoming_byte_rate gauge
kafka_server_controller_channel_metrics_incoming_byte_rate{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_incoming_byte_rate{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_incoming_byte_rate{broker_id="2",} 0.0
# HELP kafka_network_processor_value Attribute exposed for management (kafka.network<type=Processor, name=IdlePercent, networkProcessor=0><>Value)
# TYPE kafka_network_processor_value gauge
kafka_network_processor_value{name="IdlePercent",networkProcessor="0",} 1.0040457773861422
kafka_network_processor_value{name="IdlePercent",networkProcessor="1",} 1.000390284084095
kafka_network_processor_value{name="IdlePercent",networkProcessor="2",} 1.0009738307788385
# HELP java_lang_operatingsystem_committedvirtualmemorysize CommittedVirtualMemorySize (java.lang<type=OperatingSystem><>CommittedVirtualMemorySize)
# TYPE java_lang_operatingsystem_committedvirtualmemorysize gauge
java_lang_operatingsystem_committedvirtualmemorysize 4.785700864E9
# HELP java_lang_memory_objectpendingfinalizationcount ObjectPendingFinalizationCount (java.lang<type=Memory><>ObjectPendingFinalizationCount)
# TYPE java_lang_memory_objectpendingfinalizationcount gauge
java_lang_memory_objectpendingfinalizationcount 0.0
# HELP kafka_server_socket_server_metrics_response_rate Responses received sent per second. (kafka.server<type=socket-server-metrics, networkProcessor=0><>response-rate)
# TYPE kafka_server_socket_server_metrics_response_rate gauge
kafka_server_socket_server_metrics_response_rate{networkProcessor="0",} 0.0
kafka_server_socket_server_metrics_response_rate{networkProcessor="1",} 2.004794072782741
kafka_server_socket_server_metrics_response_rate{networkProcessor="2",} 2.0044883107828397
# HELP java_lang_garbagecollector_lastgcinfo_gcthreadcount CompositeType for GC info for G1 Young Generation (java.lang<type=GarbageCollector, name=G1 Young Generation><LastGcInfo>GcThreadCount)
# TYPE java_lang_garbagecollector_lastgcinfo_gcthreadcount gauge
java_lang_garbagecollector_lastgcinfo_gcthreadcount{name="G1 Young Generation",} 6.0
# HELP java_lang_operatingsystem_systemcpuload SystemCpuLoad (java.lang<type=OperatingSystem><>SystemCpuLoad)
# TYPE java_lang_operatingsystem_systemcpuload gauge
java_lang_operatingsystem_systemcpuload 0.18297315655739768
# HELP kafka_server_socket_server_metrics_io_ratio The fraction of time the I/O thread spent doing I/O (kafka.server<type=socket-server-metrics, networkProcessor=0><>io-ratio)
# TYPE kafka_server_socket_server_metrics_io_ratio gauge
kafka_server_socket_server_metrics_io_ratio{networkProcessor="0",} 5.2219269400932676E-5
kafka_server_socket_server_metrics_io_ratio{networkProcessor="1",} 0.0013977561123279377
kafka_server_socket_server_metrics_io_ratio{networkProcessor="2",} 0.001090867442063839
# HELP java_lang_operatingsystem_availableprocessors AvailableProcessors (java.lang<type=OperatingSystem><>AvailableProcessors)
# TYPE java_lang_operatingsystem_availableprocessors gauge
java_lang_operatingsystem_availableprocessors 2.0
# HELP java_lang_memorypool_collectionusagethresholdcount CollectionUsageThresholdCount (java.lang<type=MemoryPool, name=G1 Eden Space><>CollectionUsageThresholdCount)
# TYPE java_lang_memorypool_collectionusagethresholdcount gauge
java_lang_memorypool_collectionusagethresholdcount{name="G1 Eden Space",} 0.0
java_lang_memorypool_collectionusagethresholdcount{name="G1 Old Gen",} 0.0
java_lang_memorypool_collectionusagethresholdcount{name="G1 Survivor Space",} 0.0
# HELP java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_init java.lang.management.MemoryUsage (java.lang<type=GarbageCollector, name=G1 Young Generation, key=G1 Survivor Space><LastGcInfo, memoryUsageBeforeGc>init)
# TYPE java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_init gauge
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_init{name="G1 Young Generation",key="G1 Survivor Space",} 0.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_init{name="G1 Young Generation",key="Compressed Class Space",} 0.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_init{name="G1 Young Generation",key="Metaspace",} 0.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_init{name="G1 Young Generation",key="G1 Old Gen",} 1.01711872E9
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_init{name="G1 Young Generation",key="G1 Eden Space",} 5.6623104E7
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_init{name="G1 Young Generation",key="Code Cache",} 2555904.0
# HELP java_lang_memorypool_usage_committed java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=Code Cache><Usage>committed)
# TYPE java_lang_memorypool_usage_committed gauge
java_lang_memorypool_usage_committed{name="Code Cache",} 1.7301504E7
java_lang_memorypool_usage_committed{name="Compressed Class Space",} 3670016.0
java_lang_memorypool_usage_committed{name="G1 Eden Space",} 8.8080384E7
java_lang_memorypool_usage_committed{name="G1 Old Gen",} 9.81467136E8
java_lang_memorypool_usage_committed{name="G1 Survivor Space",} 4194304.0
java_lang_memorypool_usage_committed{name="Metaspace",} 2.7262976E7
# HELP java_lang_operatingsystem_totalphysicalmemorysize TotalPhysicalMemorySize (java.lang<type=OperatingSystem><>TotalPhysicalMemorySize)
# TYPE java_lang_operatingsystem_totalphysicalmemorysize gauge
java_lang_operatingsystem_totalphysicalmemorysize 5.202337792E9
# HELP kafka_server_controller_channel_metrics_request_size_avg The average size of all requests in the window.. (kafka.server<type=controller-channel-metrics, broker-id=0><>request-size-avg)
# TYPE kafka_server_controller_channel_metrics_request_size_avg gauge
kafka_server_controller_channel_metrics_request_size_avg{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_request_size_avg{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_request_size_avg{broker_id="2",} 0.0
# HELP java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_max java.lang.management.MemoryUsage (java.lang<type=GarbageCollector, name=G1 Young Generation, key=G1 Survivor Space><LastGcInfo, memoryUsageBeforeGc>max)
# TYPE java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_max gauge
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_max{name="G1 Young Generation",key="G1 Survivor Space",} -1.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_max{name="G1 Young Generation",key="Compressed Class Space",} 1.073741824E9
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_max{name="G1 Young Generation",key="Metaspace",} -1.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_max{name="G1 Young Generation",key="G1 Old Gen",} 1.073741824E9
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_max{name="G1 Young Generation",key="G1 Eden Space",} -1.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_max{name="G1 Young Generation",key="Code Cache",} 2.5165824E8
# HELP java_lang_threading_daemonthreadcount DaemonThreadCount (java.lang<type=Threading><>DaemonThreadCount)
# TYPE java_lang_threading_daemonthreadcount gauge
java_lang_threading_daemonthreadcount 36.0
# HELP java_lang_memorypool_usage_init java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=Code Cache><Usage>init)
# TYPE java_lang_memorypool_usage_init gauge
java_lang_memorypool_usage_init{name="Code Cache",} 2555904.0
java_lang_memorypool_usage_init{name="Compressed Class Space",} 0.0
java_lang_memorypool_usage_init{name="G1 Eden Space",} 5.6623104E7
java_lang_memorypool_usage_init{name="G1 Old Gen",} 1.01711872E9
java_lang_memorypool_usage_init{name="G1 Survivor Space",} 0.0
java_lang_memorypool_usage_init{name="Metaspace",} 0.0
# HELP kafka_server_replicafetchermanager_value Attribute exposed for management (kafka.server<type=ReplicaFetcherManager, name=MaxLag, clientId=Replica><>Value)
# TYPE kafka_server_replicafetchermanager_value gauge
kafka_server_replicafetchermanager_value{name="MaxLag",clientId="Replica",} 0.0
# HELP java_lang_memorypool_collectionusage_max java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=G1 Eden Space><CollectionUsage>max)
# TYPE java_lang_memorypool_collectionusage_max gauge
java_lang_memorypool_collectionusage_max{name="G1 Eden Space",} -1.0
java_lang_memorypool_collectionusage_max{name="G1 Old Gen",} 1.073741824E9
java_lang_memorypool_collectionusage_max{name="G1 Survivor Space",} -1.0
# HELP java_lang_threading_currentthreadusertime CurrentThreadUserTime (java.lang<type=Threading><>CurrentThreadUserTime)
# TYPE java_lang_threading_currentthreadusertime gauge
java_lang_threading_currentthreadusertime 4.0E7
# HELP kafka_server_controller_channel_metrics_io_wait_ratio The fraction of time the I/O thread spent waiting. (kafka.server<type=controller-channel-metrics, broker-id=0><>io-wait-ratio)
# TYPE kafka_server_controller_channel_metrics_io_wait_ratio gauge
kafka_server_controller_channel_metrics_io_wait_ratio{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_io_wait_ratio{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_io_wait_ratio{broker_id="2",} 0.0
# HELP java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_used java.lang.management.MemoryUsage (java.lang<type=GarbageCollector, name=G1 Young Generation, key=G1 Survivor Space><LastGcInfo, memoryUsageAfterGc>used)
# TYPE java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_used gauge
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_used{name="G1 Young Generation",key="G1 Survivor Space",} 4194304.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_used{name="G1 Young Generation",key="Compressed Class Space",} 3525912.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_used{name="G1 Young Generation",key="Metaspace",} 2.6626584E7
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_used{name="G1 Young Generation",key="G1 Old Gen",} 1.55179008E8
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_used{name="G1 Young Generation",key="G1 Eden Space",} 0.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_used{name="G1 Young Generation",key="Code Cache",} 1.6669568E7
# HELP kafka_log_log_value Attribute exposed for management (kafka.log<type=Log, name=LogEndOffset, topic=__consumer_offsets, partition=0><>Value)
# TYPE kafka_log_log_value gauge
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="0",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="test1",partition="0",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="1",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="10",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="11",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="12",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="13",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="14",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="15",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="16",} 3.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="17",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="18",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="19",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="2",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="20",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="21",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="22",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="23",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="24",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="25",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="26",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="27",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="28",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="29",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="3",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="30",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="31",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="32",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="33",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="34",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="35",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="36",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="37",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="38",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="39",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="4",} 8.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="40",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="41",} 12.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="42",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="43",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="44",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="45",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="46",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="47",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="48",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="49",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="5",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="6",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="7",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="8",} 0.0
kafka_log_log_value{name="LogEndOffset",topic="__consumer_offsets",partition="9",} 0.0
# HELP kafka_server_kafkarequesthandlerpool_oneminuterate Attribute exposed for management (kafka.server<type=KafkaRequestHandlerPool, name=RequestHandlerAvgIdlePercent><>OneMinuteRate)
# TYPE kafka_server_kafkarequesthandlerpool_oneminuterate gauge
kafka_server_kafkarequesthandlerpool_oneminuterate{name="RequestHandlerAvgIdlePercent",} 0.9996590142663573
# HELP java_lang_memory_nonheapmemoryusage_committed java.lang.management.MemoryUsage (java.lang<type=Memory><NonHeapMemoryUsage>committed)
# TYPE java_lang_memory_nonheapmemoryusage_committed gauge
java_lang_memory_nonheapmemoryusage_committed 4.8234496E7
# HELP kafka_server_produce_queue_size Tracks the size of the delay queue (kafka.server<type=Produce><>queue-size)
# TYPE kafka_server_produce_queue_size gauge
kafka_server_produce_queue_size 0.0
# HELP java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_committed java.lang.management.MemoryUsage (java.lang<type=GarbageCollector, name=G1 Young Generation, key=G1 Survivor Space><LastGcInfo, memoryUsageBeforeGc>committed)
# TYPE java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_committed gauge
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_committed{name="G1 Young Generation",key="G1 Survivor Space",} 6291456.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_committed{name="G1 Young Generation",key="Compressed Class Space",} 3670016.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_committed{name="G1 Young Generation",key="Metaspace",} 2.7262976E7
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_committed{name="G1 Young Generation",key="G1 Old Gen",} 9.68884224E8
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_committed{name="G1 Young Generation",key="G1 Eden Space",} 9.8566144E7
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_committed{name="G1 Young Generation",key="Code Cache",} 1.7301504E7
# HELP java_lang_operatingsystem_processcputime ProcessCpuTime (java.lang<type=OperatingSystem><>ProcessCpuTime)
# TYPE java_lang_operatingsystem_processcputime gauge
java_lang_operatingsystem_processcputime 7.927E10
# HELP java_lang_compilation_totalcompilationtime TotalCompilationTime (java.lang<type=Compilation><>TotalCompilationTime)
# TYPE java_lang_compilation_totalcompilationtime gauge
java_lang_compilation_totalcompilationtime 38490.0
# HELP java_lang_memory_heapmemoryusage_used java.lang.management.MemoryUsage (java.lang<type=Memory><HeapMemoryUsage>used)
# TYPE java_lang_memory_heapmemoryusage_used gauge
java_lang_memory_heapmemoryusage_used 1.97122048E8
# HELP kafka_server_socket_server_metrics_select_rate Number of times the I/O layer checked for new I/O to perform per second (kafka.server<type=socket-server-metrics, networkProcessor=0><>select-rate)
# TYPE kafka_server_socket_server_metrics_select_rate gauge
kafka_server_socket_server_metrics_select_rate{networkProcessor="0",} 3.348180921347643
kafka_server_socket_server_metrics_select_rate{networkProcessor="1",} 7.977776194885676
kafka_server_socket_server_metrics_select_rate{networkProcessor="2",} 7.975653268968412
# HELP kafka_network_socketserver_value Attribute exposed for management (kafka.network<type=SocketServer, name=NetworkProcessorAvgIdlePercent><>Value)
# TYPE kafka_network_socketserver_value gauge
kafka_network_socketserver_value{name="NetworkProcessorAvgIdlePercent",} 1.0024901782565627
# HELP java_lang_classloading_unloadedclasscount UnloadedClassCount (java.lang<type=ClassLoading><>UnloadedClassCount)
# TYPE java_lang_classloading_unloadedclasscount gauge
java_lang_classloading_unloadedclasscount 0.0
# HELP kafka_server_socket_server_metrics_request_size_avg The average size of all requests in the window.. (kafka.server<type=socket-server-metrics, networkProcessor=0><>request-size-avg)
# TYPE kafka_server_socket_server_metrics_request_size_avg gauge
kafka_server_socket_server_metrics_request_size_avg{networkProcessor="0",} 0.0
kafka_server_socket_server_metrics_request_size_avg{networkProcessor="1",} 346.0
kafka_server_socket_server_metrics_request_size_avg{networkProcessor="2",} 346.0
# HELP kafka_server_replicamanager_value Attribute exposed for management (kafka.server<type=ReplicaManager, name=LeaderCount><>Value)
# TYPE kafka_server_replicamanager_value gauge
kafka_server_replicamanager_value{name="LeaderCount",} 18.0
kafka_server_replicamanager_value{name="PartitionCount",} 51.0
kafka_server_replicamanager_value{name="UnderReplicatedPartitions",} 0.0
# HELP kafka_server_controller_channel_metrics_request_size_max The maximum size of any request sent in the window. (kafka.server<type=controller-channel-metrics, broker-id=0><>request-size-max)
# TYPE kafka_server_controller_channel_metrics_request_size_max gauge
kafka_server_controller_channel_metrics_request_size_max{broker_id="0",} -Inf
kafka_server_controller_channel_metrics_request_size_max{broker_id="1",} -Inf
kafka_server_controller_channel_metrics_request_size_max{broker_id="2",} -Inf
# HELP java_lang_memorypool_collectionusage_committed java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=G1 Eden Space><CollectionUsage>committed)
# TYPE java_lang_memorypool_collectionusage_committed gauge
java_lang_memorypool_collectionusage_committed{name="G1 Eden Space",} 8.8080384E7
java_lang_memorypool_collectionusage_committed{name="G1 Old Gen",} 0.0
java_lang_memorypool_collectionusage_committed{name="G1 Survivor Space",} 4194304.0
# HELP java_lang_memorypool_peakusage_max java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=Code Cache><PeakUsage>max)
# TYPE java_lang_memorypool_peakusage_max gauge
java_lang_memorypool_peakusage_max{name="Code Cache",} 2.5165824E8
java_lang_memorypool_peakusage_max{name="Compressed Class Space",} 1.073741824E9
java_lang_memorypool_peakusage_max{name="G1 Eden Space",} -1.0
java_lang_memorypool_peakusage_max{name="G1 Old Gen",} 1.073741824E9
java_lang_memorypool_peakusage_max{name="G1 Survivor Space",} -1.0
java_lang_memorypool_peakusage_max{name="Metaspace",} -1.0
# HELP kafka_controller_kafkacontroller_value Attribute exposed for management (kafka.controller<type=KafkaController, name=ActiveControllerCount><>Value)
# TYPE kafka_controller_kafkacontroller_value gauge
kafka_controller_kafkacontroller_value{name="ActiveControllerCount",} 1.0
kafka_controller_kafkacontroller_value{name="OfflinePartitionsCount",} 0.0
kafka_controller_kafkacontroller_value{name="PreferredReplicaImbalanceCount",} 0.0
# HELP java_lang_memorypool_peakusage_used java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=Code Cache><PeakUsage>used)
# TYPE java_lang_memorypool_peakusage_used gauge
java_lang_memorypool_peakusage_used{name="Code Cache",} 1.7131776E7
java_lang_memorypool_peakusage_used{name="Compressed Class Space",} 3525912.0
java_lang_memorypool_peakusage_used{name="G1 Eden Space",} 1.27926272E8
java_lang_memorypool_peakusage_used{name="G1 Old Gen",} 1.93214648E8
java_lang_memorypool_peakusage_used{name="G1 Survivor Space",} 7340032.0
java_lang_memorypool_peakusage_used{name="Metaspace",} 2.6679912E7
# HELP java_lang_classloading_totalloadedclasscount TotalLoadedClassCount (java.lang<type=ClassLoading><>TotalLoadedClassCount)
# TYPE java_lang_classloading_totalloadedclasscount gauge
java_lang_classloading_totalloadedclasscount 4400.0
# HELP java_lang_memorypool_collectionusagethreshold CollectionUsageThreshold (java.lang<type=MemoryPool, name=G1 Eden Space><>CollectionUsageThreshold)
# TYPE java_lang_memorypool_collectionusagethreshold gauge
java_lang_memorypool_collectionusagethreshold{name="G1 Eden Space",} 0.0
java_lang_memorypool_collectionusagethreshold{name="G1 Old Gen",} 0.0
java_lang_memorypool_collectionusagethreshold{name="G1 Survivor Space",} 0.0
# HELP kafka_server_controller_channel_metrics_request_rate The average number of requests sent per second. (kafka.server<type=controller-channel-metrics, broker-id=0><>request-rate)
# TYPE kafka_server_controller_channel_metrics_request_rate gauge
kafka_server_controller_channel_metrics_request_rate{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_request_rate{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_request_rate{broker_id="2",} 0.0
# HELP java_lang_garbagecollector_lastgcinfo_starttime CompositeType for GC info for G1 Young Generation (java.lang<type=GarbageCollector, name=G1 Young Generation><LastGcInfo>startTime)
# TYPE java_lang_garbagecollector_lastgcinfo_starttime gauge
java_lang_garbagecollector_lastgcinfo_starttime{name="G1 Young Generation",} 2939337.0
# HELP java_lang_operatingsystem_maxfiledescriptorcount MaxFileDescriptorCount (java.lang<type=OperatingSystem><>MaxFileDescriptorCount)
# TYPE java_lang_operatingsystem_maxfiledescriptorcount gauge
java_lang_operatingsystem_maxfiledescriptorcount 1048576.0
# HELP java_lang_threading_peakthreadcount PeakThreadCount (java.lang<type=Threading><>PeakThreadCount)
# TYPE java_lang_threading_peakthreadcount gauge
java_lang_threading_peakthreadcount 55.0
# HELP java_lang_classloading_loadedclasscount LoadedClassCount (java.lang<type=ClassLoading><>LoadedClassCount)
# TYPE java_lang_classloading_loadedclasscount gauge
java_lang_classloading_loadedclasscount 4400.0
# HELP kafka_server_controller_channel_metrics_outgoing_byte_rate The average number of outgoing bytes sent per second to all servers. (kafka.server<type=controller-channel-metrics, broker-id=0><>outgoing-byte-rate)
# TYPE kafka_server_controller_channel_metrics_outgoing_byte_rate gauge
kafka_server_controller_channel_metrics_outgoing_byte_rate{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_outgoing_byte_rate{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_outgoing_byte_rate{broker_id="2",} 0.0
# HELP java_lang_operatingsystem_freephysicalmemorysize FreePhysicalMemorySize (java.lang<type=OperatingSystem><>FreePhysicalMemorySize)
# TYPE java_lang_operatingsystem_freephysicalmemorysize gauge
java_lang_operatingsystem_freephysicalmemorysize 2.56577536E8
# HELP kafka_server_socket_server_metrics_connection_count The current number of active connections. (kafka.server<type=socket-server-metrics, networkProcessor=0><>connection-count)
# TYPE kafka_server_socket_server_metrics_connection_count gauge
kafka_server_socket_server_metrics_connection_count{networkProcessor="0",} 0.0
kafka_server_socket_server_metrics_connection_count{networkProcessor="1",} 1.0
kafka_server_socket_server_metrics_connection_count{networkProcessor="2",} 1.0
# HELP kafka_server_fetch_queue_size Tracks the size of the delay queue (kafka.server<type=Fetch><>queue-size)
# TYPE kafka_server_fetch_queue_size gauge
kafka_server_fetch_queue_size 0.0
# HELP java_lang_memorypool_collectionusage_used java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=G1 Eden Space><CollectionUsage>used)
# TYPE java_lang_memorypool_collectionusage_used gauge
java_lang_memorypool_collectionusage_used{name="G1 Eden Space",} 0.0
java_lang_memorypool_collectionusage_used{name="G1 Old Gen",} 0.0
java_lang_memorypool_collectionusage_used{name="G1 Survivor Space",} 4194304.0
# HELP java_lang_memorypool_usage_used java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=Code Cache><Usage>used)
# TYPE java_lang_memorypool_usage_used gauge
java_lang_memorypool_usage_used{name="Code Cache",} 1.670848E7
java_lang_memorypool_usage_used{name="Compressed Class Space",} 3525912.0
java_lang_memorypool_usage_used{name="G1 Eden Space",} 3.7748736E7
java_lang_memorypool_usage_used{name="G1 Old Gen",} 1.55179008E8
java_lang_memorypool_usage_used{name="G1 Survivor Space",} 4194304.0
java_lang_memorypool_usage_used{name="Metaspace",} 2.6679912E7
# HELP java_lang_garbagecollector_collectiontime CollectionTime (java.lang<type=GarbageCollector, name=G1 Old Generation><>CollectionTime)
# TYPE java_lang_garbagecollector_collectiontime gauge
java_lang_garbagecollector_collectiontime{name="G1 Old Generation",} 0.0
java_lang_garbagecollector_collectiontime{name="G1 Young Generation",} 1172.0
# HELP kafka_server_socket_server_metrics_outgoing_byte_rate The average number of outgoing bytes sent per second to all servers. (kafka.server<type=socket-server-metrics, networkProcessor=0><>outgoing-byte-rate)
# TYPE kafka_server_socket_server_metrics_outgoing_byte_rate gauge
kafka_server_socket_server_metrics_outgoing_byte_rate{networkProcessor="0",} 0.0
kafka_server_socket_server_metrics_outgoing_byte_rate{networkProcessor="1",} 693.6464575255551
kafka_server_socket_server_metrics_outgoing_byte_rate{networkProcessor="2",} 693.2799013563503
# HELP java_lang_runtime_uptime Uptime (java.lang<type=Runtime><>Uptime)
# TYPE java_lang_runtime_uptime gauge
java_lang_runtime_uptime 2991325.0
# HELP java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_used java.lang.management.MemoryUsage (java.lang<type=GarbageCollector, name=G1 Young Generation, key=G1 Survivor Space><LastGcInfo, memoryUsageBeforeGc>used)
# TYPE java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_used gauge
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_used{name="G1 Young Generation",key="G1 Survivor Space",} 6291456.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_used{name="G1 Young Generation",key="Compressed Class Space",} 3525912.0
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_used{name="G1 Young Generation",key="Metaspace",} 2.6626584E7
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_used{name="G1 Young Generation",key="G1 Old Gen",} 1.54654728E8
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_used{name="G1 Young Generation",key="G1 Eden Space",} 9.3323264E7
java_lang_garbagecollector_lastgcinfo_memoryusagebeforegc_used{name="G1 Young Generation",key="Code Cache",} 1.6669568E7
# HELP java_lang_memory_nonheapmemoryusage_init java.lang.management.MemoryUsage (java.lang<type=Memory><NonHeapMemoryUsage>init)
# TYPE java_lang_memory_nonheapmemoryusage_init gauge
java_lang_memory_nonheapmemoryusage_init 2555904.0
# HELP kafka_server_controller_channel_metrics_select_rate Number of times the I/O layer checked for new I/O to perform per second (kafka.server<type=controller-channel-metrics, broker-id=0><>select-rate)
# TYPE kafka_server_controller_channel_metrics_select_rate gauge
kafka_server_controller_channel_metrics_select_rate{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_select_rate{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_select_rate{broker_id="2",} 0.0
# HELP kafka_server_controller_channel_metrics_response_rate Responses received sent per second. (kafka.server<type=controller-channel-metrics, broker-id=0><>response-rate)
# TYPE kafka_server_controller_channel_metrics_response_rate gauge
kafka_server_controller_channel_metrics_response_rate{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_response_rate{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_response_rate{broker_id="2",} 0.0
# HELP kafka_server_controller_channel_metrics_io_wait_time_ns_avg The average length of time the I/O thread spent waiting for a socket ready for reads or writes in nanoseconds. (kafka.server<type=controller-channel-metrics, broker-id=0><>io-wait-time-ns-avg)
# TYPE kafka_server_controller_channel_metrics_io_wait_time_ns_avg gauge
kafka_server_controller_channel_metrics_io_wait_time_ns_avg{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_io_wait_time_ns_avg{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_io_wait_time_ns_avg{broker_id="2",} 0.0
# HELP kafka_controller_controllerstats_oneminuterate Attribute exposed for management (kafka.controller<type=ControllerStats, name=LeaderElectionRateAndTimeMs><>OneMinuteRate)
# TYPE kafka_controller_controllerstats_oneminuterate gauge
kafka_controller_controllerstats_oneminuterate{name="LeaderElectionRateAndTimeMs",} 1.3272588297965072E-22
kafka_controller_controllerstats_oneminuterate{name="UncleanLeaderElectionsPerSec",} 0.0
# HELP java_lang_operatingsystem_systemloadaverage SystemLoadAverage (java.lang<type=OperatingSystem><>SystemLoadAverage)
# TYPE java_lang_operatingsystem_systemloadaverage gauge
java_lang_operatingsystem_systemloadaverage 1.29
# HELP java_lang_memorypool_peakusage_init java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=Code Cache><PeakUsage>init)
# TYPE java_lang_memorypool_peakusage_init gauge
java_lang_memorypool_peakusage_init{name="Code Cache",} 2555904.0
java_lang_memorypool_peakusage_init{name="Compressed Class Space",} 0.0
java_lang_memorypool_peakusage_init{name="G1 Eden Space",} 5.6623104E7
java_lang_memorypool_peakusage_init{name="G1 Old Gen",} 1.01711872E9
java_lang_memorypool_peakusage_init{name="G1 Survivor Space",} 0.0
java_lang_memorypool_peakusage_init{name="Metaspace",} 0.0
# HELP java_lang_garbagecollector_lastgcinfo_duration CompositeType for GC info for G1 Young Generation (java.lang<type=GarbageCollector, name=G1 Young Generation><LastGcInfo>duration)
# TYPE java_lang_garbagecollector_lastgcinfo_duration gauge
java_lang_garbagecollector_lastgcinfo_duration{name="G1 Young Generation",} 22.0
# HELP java_lang_memory_nonheapmemoryusage_max java.lang.management.MemoryUsage (java.lang<type=Memory><NonHeapMemoryUsage>max)
# TYPE java_lang_memory_nonheapmemoryusage_max gauge
java_lang_memory_nonheapmemoryusage_max -1.0
# HELP kafka_server_controller_channel_metrics_connection_creation_rate New connections established per second in the window. (kafka.server<type=controller-channel-metrics, broker-id=0><>connection-creation-rate)
# TYPE kafka_server_controller_channel_metrics_connection_creation_rate gauge
kafka_server_controller_channel_metrics_connection_creation_rate{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_connection_creation_rate{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_connection_creation_rate{broker_id="2",} 0.0
# HELP java_lang_threading_currentthreadcputime CurrentThreadCpuTime (java.lang<type=Threading><>CurrentThreadCpuTime)
# TYPE java_lang_threading_currentthreadcputime gauge
java_lang_threading_currentthreadcputime 6.7286643E7
# HELP java_lang_memory_heapmemoryusage_committed java.lang.management.MemoryUsage (java.lang<type=Memory><HeapMemoryUsage>committed)
# TYPE java_lang_memory_heapmemoryusage_committed gauge
java_lang_memory_heapmemoryusage_committed 1.073741824E9
# HELP kafka_network_requestmetrics_oneminuterate Attribute exposed for management (kafka.network<type=RequestMetrics, name=RequestsPerSec, request=ApiVersions><>OneMinuteRate)
# TYPE kafka_network_requestmetrics_oneminuterate gauge
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="ApiVersions",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="ControlledShutdown",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="CreateTopics",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="DeleteTopics",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="DescribeGroups",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="Fetch",} 3.9560007313558985
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="FetchConsumer",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="FetchFollower",} 3.9827949626317114
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="GroupCoordinator",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="Heartbeat",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="JoinGroup",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="LeaderAndIsr",} 2.5786999318061392E-20
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="LeaveGroup",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="ListGroups",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="Metadata",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="OffsetCommit",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="OffsetFetch",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="Offsets",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="Produce",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="SaslHandshake",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="StopReplica",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="SyncGroup",} 0.0
kafka_network_requestmetrics_oneminuterate{name="RequestsPerSec",request="UpdateMetadata",} 2.7493204970164856E-20
# HELP java_lang_operatingsystem_processcpuload ProcessCpuLoad (java.lang<type=OperatingSystem><>ProcessCpuLoad)
# TYPE java_lang_operatingsystem_processcpuload gauge
java_lang_operatingsystem_processcpuload 0.011176381795768447
# HELP kafka_server_brokertopicmetrics_oneminuterate Attribute exposed for management (kafka.server<type=BrokerTopicMetrics, name=BytesInPerSec, topic=__consumer_offsets><>OneMinuteRate)
# TYPE kafka_server_brokertopicmetrics_oneminuterate gauge
kafka_server_brokertopicmetrics_oneminuterate{name="BytesInPerSec",topic="__consumer_offsets",} 0.0
kafka_server_brokertopicmetrics_oneminuterate{name="BytesOutPerSec",topic="__consumer_offsets",} 0.0
kafka_server_brokertopicmetrics_oneminuterate{name="BytesRejectedPerSec",topic="__consumer_offsets",} 0.0
kafka_server_brokertopicmetrics_oneminuterate{name="FailedFetchRequestsPerSec",topic="__consumer_offsets",} 0.0
kafka_server_brokertopicmetrics_oneminuterate{name="FailedProduceRequestsPerSec",topic="__consumer_offsets",} 0.0
kafka_server_brokertopicmetrics_oneminuterate{name="MessagesInPerSec",topic="__consumer_offsets",} 0.0
kafka_server_brokertopicmetrics_oneminuterate{name="TotalFetchRequestsPerSec",topic="__consumer_offsets",} 135.36428193209875
kafka_server_brokertopicmetrics_oneminuterate{name="TotalProduceRequestsPerSec",topic="__consumer_offsets",} 0.0
# HELP java_lang_garbagecollector_collectioncount CollectionCount (java.lang<type=GarbageCollector, name=G1 Old Generation><>CollectionCount)
# TYPE java_lang_garbagecollector_collectioncount gauge
java_lang_garbagecollector_collectioncount{name="G1 Old Generation",} 0.0
java_lang_garbagecollector_collectioncount{name="G1 Young Generation",} 47.0
# HELP kafka_server_socket_server_metrics_connection_close_rate Connections closed per second in the window. (kafka.server<type=socket-server-metrics, networkProcessor=0><>connection-close-rate)
# TYPE kafka_server_socket_server_metrics_connection_close_rate gauge
kafka_server_socket_server_metrics_connection_close_rate{networkProcessor="0",} 0.0
kafka_server_socket_server_metrics_connection_close_rate{networkProcessor="1",} 0.0
kafka_server_socket_server_metrics_connection_close_rate{networkProcessor="2",} 0.0
# HELP java_lang_memorypool_peakusage_committed java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=Code Cache><PeakUsage>committed)
# TYPE java_lang_memorypool_peakusage_committed gauge
java_lang_memorypool_peakusage_committed{name="Code Cache",} 1.7301504E7
java_lang_memorypool_peakusage_committed{name="Compressed Class Space",} 3670016.0
java_lang_memorypool_peakusage_committed{name="G1 Eden Space",} 1.4155776E8
java_lang_memorypool_peakusage_committed{name="G1 Old Gen",} 1.01711872E9
java_lang_memorypool_peakusage_committed{name="G1 Survivor Space",} 7340032.0
java_lang_memorypool_peakusage_committed{name="Metaspace",} 2.7262976E7
# HELP kafka_server_socket_server_metrics_io_wait_ratio The fraction of time the I/O thread spent waiting. (kafka.server<type=socket-server-metrics, networkProcessor=0><>io-wait-ratio)
# TYPE kafka_server_socket_server_metrics_io_wait_ratio gauge
kafka_server_socket_server_metrics_io_wait_ratio{networkProcessor="0",} 1.0067216558950138
kafka_server_socket_server_metrics_io_wait_ratio{networkProcessor="1",} 1.0017211131094412
kafka_server_socket_server_metrics_io_wait_ratio{networkProcessor="2",} 1.0020643859833185
# HELP kafka_server_controller_channel_metrics_io_time_ns_avg The average length of time for I/O per select call in nanoseconds. (kafka.server<type=controller-channel-metrics, broker-id=0><>io-time-ns-avg)
# TYPE kafka_server_controller_channel_metrics_io_time_ns_avg gauge
kafka_server_controller_channel_metrics_io_time_ns_avg{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_io_time_ns_avg{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_io_time_ns_avg{broker_id="2",} 0.0
# HELP java_lang_memorypool_usagethreshold UsageThreshold (java.lang<type=MemoryPool, name=Code Cache><>UsageThreshold)
# TYPE java_lang_memorypool_usagethreshold gauge
java_lang_memorypool_usagethreshold{name="Code Cache",} 0.0
java_lang_memorypool_usagethreshold{name="Compressed Class Space",} 0.0
java_lang_memorypool_usagethreshold{name="G1 Old Gen",} 0.0
java_lang_memorypool_usagethreshold{name="Metaspace",} 0.0
# HELP java_lang_memorypool_collectionusage_init java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=G1 Eden Space><CollectionUsage>init)
# TYPE java_lang_memorypool_collectionusage_init gauge
java_lang_memorypool_collectionusage_init{name="G1 Eden Space",} 5.6623104E7
java_lang_memorypool_collectionusage_init{name="G1 Old Gen",} 1.01711872E9
java_lang_memorypool_collectionusage_init{name="G1 Survivor Space",} 0.0
# HELP java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_committed java.lang.management.MemoryUsage (java.lang<type=GarbageCollector, name=G1 Young Generation, key=G1 Survivor Space><LastGcInfo, memoryUsageAfterGc>committed)
# TYPE java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_committed gauge
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_committed{name="G1 Young Generation",key="G1 Survivor Space",} 4194304.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_committed{name="G1 Young Generation",key="Compressed Class Space",} 3670016.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_committed{name="G1 Young Generation",key="Metaspace",} 2.7262976E7
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_committed{name="G1 Young Generation",key="G1 Old Gen",} 9.81467136E8
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_committed{name="G1 Young Generation",key="G1 Eden Space",} 8.8080384E7
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_committed{name="G1 Young Generation",key="Code Cache",} 1.7301504E7
# HELP kafka_server_socket_server_metrics_io_time_ns_avg The average length of time for I/O per select call in nanoseconds. (kafka.server<type=socket-server-metrics, networkProcessor=0><>io-time-ns-avg)
# TYPE kafka_server_socket_server_metrics_io_time_ns_avg gauge
kafka_server_socket_server_metrics_io_time_ns_avg{networkProcessor="0",} 15596.776785714286
kafka_server_socket_server_metrics_io_time_ns_avg{networkProcessor="1",} 175209.35267857142
kafka_server_socket_server_metrics_io_time_ns_avg{networkProcessor="2",} 136777.0745614035
# HELP kafka_server_controller_channel_metrics_connection_count The current number of active connections. (kafka.server<type=controller-channel-metrics, broker-id=0><>connection-count)
# TYPE kafka_server_controller_channel_metrics_connection_count gauge
kafka_server_controller_channel_metrics_connection_count{broker_id="0",} 1.0
kafka_server_controller_channel_metrics_connection_count{broker_id="1",} 1.0
kafka_server_controller_channel_metrics_connection_count{broker_id="2",} 1.0
# HELP kafka_server_socket_server_metrics_incoming_byte_rate Bytes/second read off all sockets (kafka.server<type=socket-server-metrics, networkProcessor=0><>incoming-byte-rate)
# TYPE kafka_server_socket_server_metrics_incoming_byte_rate gauge
kafka_server_socket_server_metrics_incoming_byte_rate{networkProcessor="0",} 0.0
kafka_server_socket_server_metrics_incoming_byte_rate{networkProcessor="1",} 701.6014816428805
kafka_server_socket_server_metrics_incoming_byte_rate{networkProcessor="2",} 701.0668408447638
# HELP java_lang_operatingsystem_totalswapspacesize TotalSwapSpaceSize (java.lang<type=OperatingSystem><>TotalSwapSpaceSize)
# TYPE java_lang_operatingsystem_totalswapspacesize gauge
java_lang_operatingsystem_totalswapspacesize 0.0
# HELP java_lang_garbagecollector_lastgcinfo_id CompositeType for GC info for G1 Young Generation (java.lang<type=GarbageCollector, name=G1 Young Generation><LastGcInfo>id)
# TYPE java_lang_garbagecollector_lastgcinfo_id gauge
java_lang_garbagecollector_lastgcinfo_id{name="G1 Young Generation",} 47.0
# HELP kafka_cluster_partition_value Attribute exposed for management (kafka.cluster<type=Partition, name=UnderReplicated, topic=__consumer_offsets, partition=0><>Value)
# TYPE kafka_cluster_partition_value gauge
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="0",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="test1",partition="0",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="1",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="10",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="11",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="12",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="13",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="14",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="15",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="16",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="17",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="18",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="19",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="2",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="20",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="21",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="22",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="23",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="24",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="25",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="26",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="27",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="28",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="29",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="3",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="30",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="31",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="32",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="33",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="34",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="35",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="36",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="37",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="38",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="39",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="4",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="40",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="41",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="42",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="43",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="44",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="45",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="46",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="47",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="48",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="49",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="5",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="6",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="7",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="8",} 0.0
kafka_cluster_partition_value{name="UnderReplicated",topic="__consumer_offsets",partition="9",} 0.0
# HELP java_lang_memorypool_usagethresholdcount UsageThresholdCount (java.lang<type=MemoryPool, name=Code Cache><>UsageThresholdCount)
# TYPE java_lang_memorypool_usagethresholdcount gauge
java_lang_memorypool_usagethresholdcount{name="Code Cache",} 0.0
java_lang_memorypool_usagethresholdcount{name="Compressed Class Space",} 0.0
java_lang_memorypool_usagethresholdcount{name="G1 Old Gen",} 0.0
java_lang_memorypool_usagethresholdcount{name="Metaspace",} 0.0
# HELP java_lang_threading_totalstartedthreadcount TotalStartedThreadCount (java.lang<type=Threading><>TotalStartedThreadCount)
# TYPE java_lang_threading_totalstartedthreadcount gauge
java_lang_threading_totalstartedthreadcount 66.0
# HELP java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_init java.lang.management.MemoryUsage (java.lang<type=GarbageCollector, name=G1 Young Generation, key=G1 Survivor Space><LastGcInfo, memoryUsageAfterGc>init)
# TYPE java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_init gauge
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_init{name="G1 Young Generation",key="G1 Survivor Space",} 0.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_init{name="G1 Young Generation",key="Compressed Class Space",} 0.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_init{name="G1 Young Generation",key="Metaspace",} 0.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_init{name="G1 Young Generation",key="G1 Old Gen",} 1.01711872E9
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_init{name="G1 Young Generation",key="G1 Eden Space",} 5.6623104E7
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_init{name="G1 Young Generation",key="Code Cache",} 2555904.0
# HELP java_lang_memorypool_usage_max java.lang.management.MemoryUsage (java.lang<type=MemoryPool, name=Code Cache><Usage>max)
# TYPE java_lang_memorypool_usage_max gauge
java_lang_memorypool_usage_max{name="Code Cache",} 2.5165824E8
java_lang_memorypool_usage_max{name="Compressed Class Space",} 1.073741824E9
java_lang_memorypool_usage_max{name="G1 Eden Space",} -1.0
java_lang_memorypool_usage_max{name="G1 Old Gen",} 1.073741824E9
java_lang_memorypool_usage_max{name="G1 Survivor Space",} -1.0
java_lang_memorypool_usage_max{name="Metaspace",} -1.0
# HELP java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_max java.lang.management.MemoryUsage (java.lang<type=GarbageCollector, name=G1 Young Generation, key=G1 Survivor Space><LastGcInfo, memoryUsageAfterGc>max)
# TYPE java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_max gauge
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_max{name="G1 Young Generation",key="G1 Survivor Space",} -1.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_max{name="G1 Young Generation",key="Compressed Class Space",} 1.073741824E9
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_max{name="G1 Young Generation",key="Metaspace",} -1.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_max{name="G1 Young Generation",key="G1 Old Gen",} 1.073741824E9
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_max{name="G1 Young Generation",key="G1 Eden Space",} -1.0
java_lang_garbagecollector_lastgcinfo_memoryusageaftergc_max{name="G1 Young Generation",key="Code Cache",} 2.5165824E8
# HELP kafka_server_socket_server_metrics_request_rate The average number of requests sent per second. (kafka.server<type=socket-server-metrics, networkProcessor=0><>request-rate)
# TYPE kafka_server_socket_server_metrics_request_rate gauge
kafka_server_socket_server_metrics_request_rate{networkProcessor="0",} 0.0
kafka_server_socket_server_metrics_request_rate{networkProcessor="1",} 2.0046702207340177
kafka_server_socket_server_metrics_request_rate{networkProcessor="2",} 2.003169851193097
# HELP java_lang_memory_nonheapmemoryusage_used java.lang.management.MemoryUsage (java.lang<type=Memory><NonHeapMemoryUsage>used)
# TYPE java_lang_memory_nonheapmemoryusage_used gauge
java_lang_memory_nonheapmemoryusage_used 4.6867256E7
# HELP java_lang_operatingsystem_openfiledescriptorcount OpenFileDescriptorCount (java.lang<type=OperatingSystem><>OpenFileDescriptorCount)
# TYPE java_lang_operatingsystem_openfiledescriptorcount gauge
java_lang_operatingsystem_openfiledescriptorcount 186.0
# HELP kafka_server_controller_channel_metrics_connection_close_rate Connections closed per second in the window. (kafka.server<type=controller-channel-metrics, broker-id=0><>connection-close-rate)
# TYPE kafka_server_controller_channel_metrics_connection_close_rate gauge
kafka_server_controller_channel_metrics_connection_close_rate{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_connection_close_rate{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_connection_close_rate{broker_id="2",} 0.0
# HELP kafka_server_socket_server_metrics_network_io_rate The average number of network operations (reads or writes) on all connections per second. (kafka.server<type=socket-server-metrics, networkProcessor=0><>network-io-rate)
# TYPE kafka_server_socket_server_metrics_network_io_rate gauge
kafka_server_socket_server_metrics_network_io_rate{networkProcessor="0",} 0.0
kafka_server_socket_server_metrics_network_io_rate{networkProcessor="1",} 4.009326041008433
kafka_server_socket_server_metrics_network_io_rate{networkProcessor="2",} 4.008801934682673
# HELP java_lang_memory_heapmemoryusage_max java.lang.management.MemoryUsage (java.lang<type=Memory><HeapMemoryUsage>max)
# TYPE java_lang_memory_heapmemoryusage_max gauge
java_lang_memory_heapmemoryusage_max 1.073741824E9
# HELP kafka_server_socket_server_metrics_connection_creation_rate New connections established per second in the window. (kafka.server<type=socket-server-metrics, networkProcessor=0><>connection-creation-rate)
# TYPE kafka_server_socket_server_metrics_connection_creation_rate gauge
kafka_server_socket_server_metrics_connection_creation_rate{networkProcessor="0",} 0.0
kafka_server_socket_server_metrics_connection_creation_rate{networkProcessor="1",} 0.0
kafka_server_socket_server_metrics_connection_creation_rate{networkProcessor="2",} 0.0
# HELP java_lang_garbagecollector_lastgcinfo_endtime CompositeType for GC info for G1 Young Generation (java.lang<type=GarbageCollector, name=G1 Young Generation><LastGcInfo>endTime)
# TYPE java_lang_garbagecollector_lastgcinfo_endtime gauge
java_lang_garbagecollector_lastgcinfo_endtime{name="G1 Young Generation",} 2939359.0
# HELP kafka_server_socket_server_metrics_request_size_max The maximum size of any request sent in the window. (kafka.server<type=socket-server-metrics, networkProcessor=0><>request-size-max)
# TYPE kafka_server_socket_server_metrics_request_size_max gauge
kafka_server_socket_server_metrics_request_size_max{networkProcessor="0",} -Inf
kafka_server_socket_server_metrics_request_size_max{networkProcessor="1",} 346.0
kafka_server_socket_server_metrics_request_size_max{networkProcessor="2",} 346.0
# HELP java_lang_runtime_starttime StartTime (java.lang<type=Runtime><>StartTime)
# TYPE java_lang_runtime_starttime gauge
java_lang_runtime_starttime 1.495220704375E12
# HELP kafka_server_replicamanager_oneminuterate Attribute exposed for management (kafka.server<type=ReplicaManager, name=IsrExpandsPerSec><>OneMinuteRate)
# TYPE kafka_server_replicamanager_oneminuterate gauge
kafka_server_replicamanager_oneminuterate{name="IsrExpandsPerSec",} 2.1131600396985195E-20
kafka_server_replicamanager_oneminuterate{name="IsrShrinksPerSec",} 1.1320500212670665E-20
# HELP java_lang_memory_heapmemoryusage_init java.lang.management.MemoryUsage (java.lang<type=Memory><HeapMemoryUsage>init)
# TYPE java_lang_memory_heapmemoryusage_init gauge
java_lang_memory_heapmemoryusage_init 1.073741824E9
# HELP java_lang_threading_threadcount ThreadCount (java.lang<type=Threading><>ThreadCount)
# TYPE java_lang_threading_threadcount gauge
java_lang_threading_threadcount 55.0
# HELP kafka_server_controller_channel_metrics_network_io_rate The average number of network operations (reads or writes) on all connections per second. (kafka.server<type=controller-channel-metrics, broker-id=0><>network-io-rate)
# TYPE kafka_server_controller_channel_metrics_network_io_rate gauge
kafka_server_controller_channel_metrics_network_io_rate{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_network_io_rate{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_network_io_rate{broker_id="2",} 0.0
# HELP java_lang_operatingsystem_freeswapspacesize FreeSwapSpaceSize (java.lang<type=OperatingSystem><>FreeSwapSpaceSize)
# TYPE java_lang_operatingsystem_freeswapspacesize gauge
java_lang_operatingsystem_freeswapspacesize 0.0
# HELP kafka_server_sessionexpirelistener_oneminuterate Attribute exposed for management (kafka.server<type=SessionExpireListener, name=ZooKeeperAuthFailuresPerSec><>OneMinuteRate)
# TYPE kafka_server_sessionexpirelistener_oneminuterate gauge
kafka_server_sessionexpirelistener_oneminuterate{name="ZooKeeperAuthFailuresPerSec",} 0.0
kafka_server_sessionexpirelistener_oneminuterate{name="ZooKeeperDisconnectsPerSec",} 0.0
kafka_server_sessionexpirelistener_oneminuterate{name="ZooKeeperExpiresPerSec",} 0.0
kafka_server_sessionexpirelistener_oneminuterate{name="ZooKeeperReadOnlyConnectsPerSec",} 0.0
kafka_server_sessionexpirelistener_oneminuterate{name="ZooKeeperSaslAuthenticationsPerSec",} 0.0
kafka_server_sessionexpirelistener_oneminuterate{name="ZooKeeperSyncConnectsPerSec",} 0.0
# HELP kafka_server_controller_channel_metrics_io_ratio The fraction of time the I/O thread spent doing I/O (kafka.server<type=controller-channel-metrics, broker-id=0><>io-ratio)
# TYPE kafka_server_controller_channel_metrics_io_ratio gauge
kafka_server_controller_channel_metrics_io_ratio{broker_id="0",} 0.0
kafka_server_controller_channel_metrics_io_ratio{broker_id="1",} 0.0
kafka_server_controller_channel_metrics_io_ratio{broker_id="2",} 0.0
# HELP jmx_scrape_duration_seconds Time this JMX scrape took, in seconds.
# TYPE jmx_scrape_duration_seconds gauge
jmx_scrape_duration_seconds 2.131619138
# HELP jmx_scrape_error Non-zero if this scrape failed.
# TYPE jmx_scrape_error gauge
jmx_scrape_error 0.0


vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ kubectl create -f prometheus-deployment.yaml -f prometheus-service.yaml 
deployment "prometheus" created
service "prometheus" created

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ kubectl get deployment,pod,service -l app=prometheus -o wide
NAME                DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/prometheus   1         1         1            1           17m

NAME                             READY     STATUS    RESTARTS   AGE       IP            NODE
po/prometheus-1738250332-6zqtm   1/1       Running   0          17m       172.18.0.21   172.17.4.200

NAME             CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE       SELECTOR
svc/prometheus   10.123.247.53   <nodes>       80:30390/TCP   21m       app=prometheus

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -L 10.123.247.53
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>Prometheus Time Series Collection and Processing Server</title>
    <script src="/static/vendor/js/jquery.min.js"></script>
    <script src="/static/vendor/bootstrap-3.3.1/js/bootstrap.min.js"></script>

    <link type="text/css" rel="stylesheet" href="/static/vendor/bootstrap-3.3.1/css/bootstrap.min.css">
    <link type="text/css" rel="stylesheet" href="/static/css/prometheus.css">

    <script>
      var PATH_PREFIX = "";
      $(function () {
        $('[data-toggle="tooltip"]').tooltip()
      })
    </script>

    
    <link type="text/css" rel="stylesheet" href="/static/css/graph.css">

    <link type="text/css" rel="stylesheet" href="/static/vendor/rickshaw/rickshaw.min.css">
    <link type="text/css" rel="stylesheet" href="/static/vendor/bootstrap-datetimepicker/bootstrap-datetimepicker.min.css">

    <script src="/static/vendor/rickshaw/vendor/d3.v3.js"></script>
    <script src="/static/vendor/rickshaw/vendor/d3.layout.min.js"></script>
    <script src="/static/vendor/rickshaw/rickshaw.min.js"></script>
    <script src="/static/vendor/bootstrap-datetimepicker/bootstrap-datetimepicker.js"></script>
    <script src="/static/vendor/bootstrap3-typeahead/bootstrap3-typeahead.min.js"></script>

    <script src="/static/vendor/js/handlebars.js"></script>
    <script src="/static/vendor/js/jquery.selection.js"></script>
    <script src="/static/vendor/js/jquery.hotkeys.js"></script>

    <script src="/static/js/graph.js?v=1"></script>

    <script id="graph_template" type="text/x-handlebars-template"></script>

  </head>

  <body>
    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container-fluid">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="/">Prometheus</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav navbar-left">
            
            
            <li><a href="/alerts">Alerts</a></li>
            <li><a href="/graph">Graph</a></li>
            <li><a href="/status">Status</a></li>
            <li>
              <a href="https://prometheus.io" target="_blank">Help</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>

    
    <div id="graph_container" class="container-fluid">
    </div>
    <div class="container-fluid">
      <div><input class="btn btn-primary" type="submit" value="Add Graph" id="add_graph"></div>
    </div>

  </body>
</html>

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ kubectl create -f grafana-deployment.yaml -f grafana-service.yaml 
deployment "grafana" created
service "grafana" created

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ kubectl get deployment,pod,service -l app=grafana -o wide
NAME             DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/grafana   1         1         1            1           1m

NAME                          READY     STATUS    RESTARTS   AGE       IP            NODE
po/grafana-1785626975-nb46z   1/1       Running   0          1m        172.18.0.22   172.17.4.200

NAME          CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE       SELECTOR
svc/grafana   10.123.245.51   <nodes>       80:30300/TCP   1m        app=grafana

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/examples/kafka$ curl -L 10.123.245.51
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width">
    <meta name="theme-color" content="#000">

    <title>Grafana</title>

		<link href='/public/css/fonts.min.css' rel='stylesheet' type='text/css'>

		
		  <link rel="stylesheet" href="/public/css/grafana.dark.min.23445646.css">
		

    <link rel="icon" type="image/png" href="/public/img/fav32.png">
		<base href="/" />

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
		</grafana-app>
  </body>

	<script>
		window.grafanaBootData = {
			user:{"isSignedIn":false,"id":0,"login":"","email":"","name":"","lightTheme":false,"orgId":0,"orgName":"","orgRole":"","isGrafanaAdmin":false,"gravatarUrl":"","timezone":"browser"},
			settings: {"allowOrgCreate":false,"appSubUrl":"","authProxyEnabled":false,"buildInfo":{"buildstamp":1462955841,"commit":"v3.0.1+2-g549eb15","hasUpdate":true,"latestVersion":"4.2.0","version":"3.0.1"},"datasources":{"-- Grafana --":{"meta":{"type":"datasource","name":"Grafana","id":"grafana","info":{"author":{"name":"","url":""},"description":"","links":null,"logos":{"small":"public/img/icn-datasource.svg","large":"public/img/icn-datasource.svg"},"screenshots":null,"version":"","updated":""},"dependencies":{"grafanaVersion":"*","plugins":[]},"includes":null,"module":"app/plugins/datasource/grafana/module","baseUrl":"public/app/plugins/datasource/grafana","annotations":false,"metrics":true,"builtIn":true,"mixed":false,"app":""},"type":"grafana"},"-- Mixed --":{"meta":{"type":"datasource","name":"Mixed datasource","id":"mixed","info":{"author":{"name":"","url":""},"description":"","links":null,"logos":{"small":"public/img/icn-datasource.svg","large":"public/img/icn-datasource.svg"},"screenshots":null,"version":"","updated":""},"dependencies":{"grafanaVersion":"*","plugins":[]},"includes":null,"module":"app/plugins/datasource/mixed/module","baseUrl":"public/app/plugins/datasource/mixed","annotations":false,"metrics":true,"builtIn":true,"mixed":true,"app":""},"type":"mixed"}},"defaultDatasource":"-- Grafana --","disableUserSignUp":false,"githubAuthEnabled":false,"googleAuthEnabled":false,"loginHint":"email or username","panels":{"dashlist":{"baseUrl":"public/app/plugins/panel/dashlist","id":"dashlist","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/dashlist/img/icn-dashlist-panel.svg","large":"public/app/plugins/panel/dashlist/img/icn-dashlist-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/dashlist/module","name":"Dashboard list"},"graph":{"baseUrl":"public/app/plugins/panel/graph","id":"graph","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/graph/img/icn-graph-panel.svg","large":"public/app/plugins/panel/graph/img/icn-graph-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/graph/module","name":"Graph"},"pluginlist":{"baseUrl":"public/app/plugins/panel/pluginlist","id":"pluginlist","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/pluginlist/img/icn-dashlist-panel.svg","large":"public/app/plugins/panel/pluginlist/img/icn-dashlist-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/pluginlist/module","name":"Plugin list"},"singlestat":{"baseUrl":"public/app/plugins/panel/singlestat","id":"singlestat","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/singlestat/img/icn-singlestat-panel.svg","large":"public/app/plugins/panel/singlestat/img/icn-singlestat-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/singlestat/module","name":"Singlestat"},"table":{"baseUrl":"public/app/plugins/panel/table","id":"table","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/table/img/icn-table-panel.svg","large":"public/app/plugins/panel/table/img/icn-table-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/table/module","name":"Table"},"text":{"baseUrl":"public/app/plugins/panel/text","id":"text","info":{"author":{"name":"Grafana Project","url":"http://grafana.org"},"description":"","links":null,"logos":{"small":"public/app/plugins/panel/text/img/icn-text-panel.svg","large":"public/app/plugins/panel/text/img/icn-text-panel.svg"},"screenshots":null,"version":"","updated":""},"module":"app/plugins/panel/text/module","name":"Text"}}},
			mainNavLinks: [{"text":"Dashboards","icon":"icon-gf icon-gf-dashboard","url":"/","children":[{"text":"Home","url":"/"},{"text":"Playlists","url":"/playlists"},{"text":"Snapshots","url":"/dashboard/snapshots"}]}]
		};
	</script>

	<script src="/public/app/boot.60b87512.js"></script>

	

	
</html>

```

Set up datasource "Prometheus"

Import _kafka.json_ configuration

![屏幕快照 2017-05-19 下午1.12.55.png](./屏幕快照%202017-05-19%20下午1.12.55.png)

![屏幕快照 2017-05-19 下午1.45.39.png](./屏幕快照%202017-05-19%20下午1.45.39.png)

![屏幕快照 2017-05-19 下午1.49.13.png](./屏幕快照%202017-05-19%20下午1.49.13.png)