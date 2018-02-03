## https://github.com/rabbitmq/rabbitmq-server/blob/master/docs/rabbitmq.conf.example
## https://github.com/rabbitmq/rabbitmq-peer-discovery-k8s/blob/master/priv/schema/rabbitmq_peer_discovery_k8s.schema

## Clustering
cluster_formation.peer_discovery_backend = rabbit_peer_discovery_k8s
cluster_formation.k8s.host = kubernetes.default.svc.{{ .Domain }}
## cluster_formation.k8s.port = {{ .K8sPort }}
## cluster_formation.k8s.scheme = https
## cluster_formation.k8s.token_path  /var/run/secrets/kubernetes.io/serviceaccount/token
## cluster_formation.k8s.cert_path  /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
## cluster_formation.k8s.namespace_path  /var/run/secrets/kubernetes.io/serviceaccount/namespace
## cluster_formation.k8s.service_name = {{ .TargetService }}
cluster_formation.k8s.address_type = ip
## cluster_formation.k8s.hostname_suffix = .{{ .TargetService }}.{{ .Namespace }}.svc.{{ .Domain }}
cluster_formation.node_cleanup.interval = 10
cluster_formation.node_cleanup.only_log_warning = false
cluster_partition_handling = autoheal
## queue master locator 
queue_master_locator=min-masters
## enable guest user  
loopback_users.guest = false
