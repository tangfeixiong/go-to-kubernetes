
## Kafka metrics and alarm

Investigation

### Prometheus backend

__https://github.com/rama-nallamilli/kafka-prometheus-monitoring__

![Prometheus dashboard](./prometheus-ui.png)
```
fanhonglingdeMacBook-Pro:docs fanhongling$ curl -L https://raw.githubusercontent.com/rama-nallamilli/kafka-prometheus-monitoring/master/images/prometheus-ui.png -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  121k  100  121k    0     0  13049      0  0:00:09  0:00:09 --:--:-- 31072
```

![Grafana dashboard](./grafana-ui.png)
```
fanhonglingdeMacBook-Pro:docs fanhongling$ curl -L https://raw.githubusercontent.com/rama-nallamilli/kafka-prometheus-monitoring/master/images/grafana-ui.png -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  334k  100  334k    0     0   1120      0  0:05:05  0:05:05 --:--:--  3952
```

__https://www.robustperception.io/monitoring-kafka-with-prometheus/__

Grafana dashboard: [Kafka Overview](https://grafana.com/dashboards/721)
```
fanhonglingdeMacBook-Pro:docs fanhongling$ curl -L https://www.robustperception.io/wp-content/uploads/2016/10/Screenshot-231016-211712.png -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  106k  100  106k    0     0  10656      0  0:00:10  0:00:10 --:--:-- 27056
```

### InfluxDB backend

__https://github.com/amient/kafka-metrics__

![Grafana dashboard](https://raw.githubusercontent.com/amient/kafka-metrics/master/doc/discovery-example-3-brokers.png)
```
fanhonglingdeMacBook-Pro:docs fanhongling$ curl -L https://raw.githubusercontent.com/amient/kafka-metrics/master/doc/discovery-example-3-brokers.png -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  140k  100  140k    0     0   8638      0  0:00:16  0:00:16 --:--:-- 14488
```

### Other 

__https://github.com/yahoo/kafka-manager__

__https://github.com/quantifind/KafkaOffsetMonitor__

__https://github.com/1ambda/kafka-connect-dashboard__