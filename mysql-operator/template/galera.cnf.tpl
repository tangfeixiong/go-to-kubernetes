[mysqld]
# binlog_format=ROW — see Binary Log Formats
# default_storage_engine=InnoDB
# innodb_autoinc_lock_mode=2
# innodb_doublewrite=1 (the default) when using Galera provider of version >= 2.0.
# query_cache_size=0 (only for versions prior to 5.5.40-galera, 10.0.14-galera and 10.1.2)
binlog_format=ROW
default-storage-engine=innodb
innodb_autoinc_lock_mode=2
bind-address=0.0.0.0

# Galera Provider Configuration
# wsrep_on=ON — Enable wsrep replication (starting 10.1.1)
# wsrep_provider — Path to the Galera library
wsrep_on=ON
wsrep_provider=/usr/lib/galera/libgalera_smm.so

# Galera Cluster Configuration
#wsrep_cluster_name	example_cluster
# wsrep_cluster_address — see cluster connection URL
wsrep_cluster_name="{{.ClusterName}}"
{{if .DisableClusterAddresses}}#{{end}}wsrep_cluster_address="gcomm://{{with .ClusterAddresses}}{{.}}{{else}}{{.FirstNodeHost}},{{.SecondNodeHost}},{{.ThirdNodeHost}}{{end}}"

# Galera Synchronization Configuration
#wsrep_sst_method	mysqldump
wsrep_sst_method=rsync

# Galera Node Configuration
#wsrep_node_address	host address:default port
#wsrep_node_name	<hostname>
{{if .ThisNodeHost}}wsrep_node_address="{{.ThisNodeHost}}"{{end}}
{{if .ThisNodeName}}wsrep_node_name="{{.ThisNodeName}}"{{end}}