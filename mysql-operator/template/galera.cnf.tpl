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
wsrep_provider=/usr/lib/libgalera_smm.so

# Galera Cluster Configuration
# wsrep_cluster_name - Defines the logical cluster name for the node. default is example_cluster
# wsrep_cluster_address — see cluster connection URL
# Refer to https://mariadb.com/kb/en/library/getting-started-with-mariadb-galera-cluster/#restarting-the-cluster
# wsrep_provider_options="pc.wait_prim=FALSE" - Default is TRUE, the node waits for the pc.wait_prim_timeout time period. Useful to bring up a non-primary component and make it primary with pc.bootstrap.
# wsrep_provider_options="pc.wait_prim_timeout=PT30S" - The period of time to wait for a primary component.
# wsrep_provider_options="pc.bootstrap=TRUE" - If you set this value to TRUE is a signal to turn a NON-PRIMARY component into PRIMARY.
wsrep_cluster_name="{{.ClusterName}}"
{{if .DisableClusterAddresses}}#{{end}}wsrep_cluster_address="gcomm://{{with .ClusterAddresses}}{{.}}{{end}}{{with .WsrepProviderOptions}}?{{.}}{{end}}"

# Galera Synchronization Configuration
#wsrep_sst_method	mysqldump
wsrep_sst_method=rsync

# Galera Node Configuration
#wsrep_node_address	host address:default port
#wsrep_node_name	<hostname>
{{if .ThisNodeHost}}wsrep_node_address="{{.ThisNodeHost}}"{{end}}
{{if .ThisNodeName}}wsrep_node_name="{{.ThisNodeName}}"{{end}}