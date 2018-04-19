DevOps

Refer to https://github.com/nanoninja/docker-gogs-mysql/blob/master/docker-compose.yml
```
[vagrant@kubedev-172-17-4-59 github.com0x2Fgogits0x2Fgogs]$ curl -jkSLO https://raw.githubusercontent.com/nanoninja/docker-gogs-mysql/master/docker-compose.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   757  100   757    0     0   1145      0 --:--:-- --:--:-- --:--:--  1145
```

env
```
[vagrant@kubedev-172-17-4-59 github.com0x2Fgogits0x2Fgogs]$ curl -jkSLO https://raw.githubusercontent.com/nanoninja/docker-gogs-mysql/master/.env
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   627  100   627    0     0    919      0 --:--:-- --:--:-- --:--:--   918
```

modify env
```
[vagrant@kubedev-172-17-4-59 github.com0x2Fgogits0x2Fgogs]$ sed -i "s/GOGS_HTTP_DOMAIN=localhost/GOGS_HTTP_DOMAIN=172.17.4.59/" .env
```

up
```
[vagrant@kubedev-172-17-4-59 github.com0x2Fgogits0x2Fgogs]$ docker-compose up
Creating network "githubcom0x2fgogits0x2fgogs_default" with the default driver
Creating volume "githubcom0x2fgogits0x2fgogs_gogs-data" with local driver
Creating volume "githubcom0x2fgogits0x2fgogs_db-data" with local driver
Pulling gogsdb (mysql:)...
Trying to pull repository docker.io/library/mysql ... 
sha256:691c55aabb3c4e3b89b953dd2f022f7ea845e5443954767d321d5f5fa394e28c: Pulling from docker.io/library/mysql
2a72cbf407d6: Pull complete
38680a9b47a8: Pull complete
4c732aa0eb1b: Pull complete
c5317a34eddd: Pull complete
f92be680366c: Pull complete
e8ecd8bec5ab: Pull complete
2a650284a6a8: Pull complete
5b5108d08c6d: Pull complete
beaff1261757: Pull complete
c1a55c6375b5: Pull complete
8181cde51c65: Pull complete
Digest: sha256:691c55aabb3c4e3b89b953dd2f022f7ea845e5443954767d321d5f5fa394e28c
Status: Downloaded newer image for docker.io/mysql:latest
Pulling gogsapp (gogs/gogs:)...
Trying to pull repository docker.io/gogs/gogs ... 
sha256:2ee2fc095278e7961f1f50eb7e38dc9815a5d005ea7e495cb5f02772fc5a43c4: Pulling from docker.io/gogs/gogs
550fe1bea624: Pull complete
ea295c89d831: Pull complete
611c78386f26: Pull complete
ddddaba6d049: Pull complete
5b8850ec29ff: Pull complete
236e4f66508d: Pull complete
ee9320cc3d06: Pull complete
e81c2b21b45b: Pull complete
64685142ab1a: Pull complete
f904fa1ed2f8: Pull complete
ae599267dccf: Pull complete
Digest: sha256:2ee2fc095278e7961f1f50eb7e38dc9815a5d005ea7e495cb5f02772fc5a43c4
Status: Downloaded newer image for docker.io/gogs/gogs:latest
Creating gogsdb ... done
Creating githubcom0x2fgogits0x2fgogs_gogsapp_1 ... done
Attaching to gogsdb, githubcom0x2fgogits0x2fgogs_gogsapp_1
gogsdb     | Initializing database
gogsdb     | 2018-04-19T23:06:08.291200Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
gogsdb     | 2018-04-19T23:06:08.871345Z 0 [Warning] InnoDB: New log files created, LSN=45790
gogsdb     | 2018-04-19T23:06:08.947065Z 0 [Warning] InnoDB: Creating foreign key constraint system tables.
gogsapp_1  | usermod: no changes
gogsdb     | 2018-04-19T23:06:09.031417Z 0 [Warning] No existing UUID has been found, so we assume that this is the first time that this server has been started. Generating a new UUID: 3ef250bd-4426-11e8-a96f-0242ac140002.
gogsdb     | 2018-04-19T23:06:09.035127Z 0 [Warning] Gtid table is not ready to be used. Table 'mysql.gtid_executed' cannot be opened.
gogsdb     | 2018-04-19T23:06:09.035508Z 1 [Warning] root@localhost is created with an empty password ! Please consider switching off the --initialize-insecure option.
gogsapp_1  | init:crond  | Cron Daemon (crond) will be run as requested by s6
gogsapp_1  | Apr 19 23:06:09 syslogd started: BusyBox v1.25.1
gogsapp_1  | Apr 19 23:06:09 sshd[31]: Server listening on :: port 22.
gogsapp_1  | Apr 19 23:06:09 sshd[31]: Server listening on 0.0.0.0 port 22.
gogsdb     | 2018-04-19T23:06:13.215694Z 1 [Warning] 'user' entry 'root@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:13.215842Z 1 [Warning] 'user' entry 'mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:13.215865Z 1 [Warning] 'user' entry 'mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:13.215890Z 1 [Warning] 'db' entry 'performance_schema mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:13.215912Z 1 [Warning] 'db' entry 'sys mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:13.215929Z 1 [Warning] 'proxies_priv' entry '@ root@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:13.215966Z 1 [Warning] 'tables_priv' entry 'user mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:13.215982Z 1 [Warning] 'tables_priv' entry 'sys_config mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | Database initialized
gogsdb     | Initializing certificates
gogsdb     | Generating a 2048 bit RSA private key
gogsdb     | ..........+++
gogsdb     | ..........................+++
gogsdb     | unable to write 'random state'
gogsdb     | writing new private key to 'ca-key.pem'
gogsdb     | -----
gogsdb     | Generating a 2048 bit RSA private key
gogsdb     | .................................+++
gogsdb     | ..................................+++
gogsdb     | unable to write 'random state'
gogsdb     | writing new private key to 'server-key.pem'
gogsdb     | -----
gogsdb     | Generating a 2048 bit RSA private key
gogsdb     | .......................................................................+++
gogsdb     | .........+++
gogsdb     | unable to write 'random state'
gogsdb     | writing new private key to 'client-key.pem'
gogsdb     | -----
gogsdb     | Certificates initialized
gogsdb     | MySQL init process in progress...
gogsdb     | 2018-04-19T23:06:18.563479Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
gogsdb     | 2018-04-19T23:06:18.566933Z 0 [Note] mysqld (mysqld 5.7.21) starting as process 88 ...
gogsdb     | 2018-04-19T23:06:18.569142Z 0 [Note] InnoDB: PUNCH HOLE support available
gogsdb     | 2018-04-19T23:06:18.569189Z 0 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
gogsdb     | 2018-04-19T23:06:18.569202Z 0 [Note] InnoDB: Uses event mutexes
gogsdb     | 2018-04-19T23:06:18.569226Z 0 [Note] InnoDB: GCC builtin __atomic_thread_fence() is used for memory barrier
gogsdb     | 2018-04-19T23:06:18.569238Z 0 [Note] InnoDB: Compressed tables use zlib 1.2.3
gogsdb     | 2018-04-19T23:06:18.569248Z 0 [Note] InnoDB: Using Linux native AIO
gogsdb     | 2018-04-19T23:06:18.569421Z 0 [Note] InnoDB: Number of pools: 1
gogsdb     | 2018-04-19T23:06:18.569510Z 0 [Note] InnoDB: Using CPU crc32 instructions
gogsdb     | 2018-04-19T23:06:18.574696Z 0 [Note] InnoDB: Initializing buffer pool, total size = 128M, instances = 1, chunk size = 128M
gogsdb     | 2018-04-19T23:06:18.597549Z 0 [Note] InnoDB: Completed initialization of buffer pool
gogsdb     | 2018-04-19T23:06:18.606292Z 0 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
gogsdb     | 2018-04-19T23:06:18.615569Z 0 [Note] InnoDB: Highest supported file format is Barracuda.
gogsdb     | 2018-04-19T23:06:18.650855Z 0 [Note] InnoDB: Creating shared tablespace for temporary tables
gogsdb     | 2018-04-19T23:06:18.651697Z 0 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
gogsdb     | 2018-04-19T23:06:18.689850Z 0 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
gogsdb     | 2018-04-19T23:06:18.690514Z 0 [Note] InnoDB: 96 redo rollback segment(s) found. 96 redo rollback segment(s) are active.
gogsdb     | 2018-04-19T23:06:18.690544Z 0 [Note] InnoDB: 32 non-redo rollback segment(s) are active.
gogsdb     | 2018-04-19T23:06:18.690744Z 0 [Note] InnoDB: Waiting for purge to start
gogsdb     | 2018-04-19T23:06:18.741277Z 0 [Note] InnoDB: 5.7.21 started; log sequence number 2551166
gogsdb     | 2018-04-19T23:06:18.741632Z 0 [Note] Plugin 'FEDERATED' is disabled.
gogsdb     | 2018-04-19T23:06:18.764428Z 0 [Note] Found ca.pem, server-cert.pem and server-key.pem in data directory. Trying to enable SSL support using them.
gogsdb     | 2018-04-19T23:06:18.766426Z 0 [Warning] CA certificate ca.pem is self signed.
gogsdb     | 2018-04-19T23:06:18.771739Z 0 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
gogsdb     | 2018-04-19T23:06:18.794386Z 0 [Note] InnoDB: Buffer pool(s) load completed at 180419 23:06:18
gogsdb     | 2018-04-19T23:06:18.801243Z 0 [Warning] 'user' entry 'root@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:18.801352Z 0 [Warning] 'user' entry 'mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:18.801388Z 0 [Warning] 'user' entry 'mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:18.801417Z 0 [Warning] 'db' entry 'performance_schema mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:18.801433Z 0 [Warning] 'db' entry 'sys mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:18.801453Z 0 [Warning] 'proxies_priv' entry '@ root@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:18.802401Z 0 [Warning] 'tables_priv' entry 'user mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:18.802443Z 0 [Warning] 'tables_priv' entry 'sys_config mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:18.810631Z 0 [Note] Event Scheduler: Loaded 0 events
gogsdb     | 2018-04-19T23:06:18.810836Z 0 [Note] mysqld: ready for connections.
gogsdb     | Version: '5.7.21'  socket: '/var/run/mysqld/mysqld.sock'  port: 0  MySQL Community Server (GPL)
gogsdb     | Warning: Unable to load '/usr/share/zoneinfo/iso3166.tab' as time zone. Skipping it.
gogsdb     | Warning: Unable to load '/usr/share/zoneinfo/leap-seconds.list' as time zone. Skipping it.
gogsdb     | Warning: Unable to load '/usr/share/zoneinfo/zone.tab' as time zone. Skipping it.
gogsdb     | Warning: Unable to load '/usr/share/zoneinfo/zone1970.tab' as time zone. Skipping it.
gogsdb     | 2018-04-19T23:06:22.586539Z 4 [Warning] 'user' entry 'root@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.586630Z 4 [Warning] 'user' entry 'mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.586656Z 4 [Warning] 'user' entry 'mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.586685Z 4 [Warning] 'db' entry 'performance_schema mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.586701Z 4 [Warning] 'db' entry 'sys mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.586718Z 4 [Warning] 'proxies_priv' entry '@ root@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.586754Z 4 [Warning] 'tables_priv' entry 'user mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.586773Z 4 [Warning] 'tables_priv' entry 'sys_config mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | mysql: [Warning] Using a password on the command line interface can be insecure.
gogsdb     | mysql: [Warning] Using a password on the command line interface can be insecure.
gogsdb     | mysql: [Warning] Using a password on the command line interface can be insecure.
gogsdb     | mysql: [Warning] Using a password on the command line interface can be insecure.
gogsdb     | 2018-04-19T23:06:22.620766Z 8 [Warning] 'user' entry 'root@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.620832Z 8 [Warning] 'user' entry 'mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.620858Z 8 [Warning] 'user' entry 'mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.620902Z 8 [Warning] 'db' entry 'performance_schema mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.620918Z 8 [Warning] 'db' entry 'sys mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.620937Z 8 [Warning] 'proxies_priv' entry '@ root@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.621057Z 8 [Warning] 'tables_priv' entry 'user mysql.session@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 2018-04-19T23:06:22.621081Z 8 [Warning] 'tables_priv' entry 'sys_config mysql.sys@localhost' ignored in --skip-name-resolve mode.
gogsdb     | 
gogsdb     | 2018-04-19T23:06:22.641755Z 0 [Note] Giving 0 client threads a chance to die gracefully
gogsdb     | 2018-04-19T23:06:22.641889Z 0 [Note] Shutting down slave threads
gogsdb     | 2018-04-19T23:06:22.642077Z 0 [Note] Forcefully disconnecting 0 remaining clients
gogsdb     | 2018-04-19T23:06:22.642116Z 0 [Note] Event Scheduler: Purging the queue. 0 events
gogsdb     | 2018-04-19T23:06:22.647376Z 0 [Note] Binlog end
gogsdb     | 2018-04-19T23:06:22.649992Z 0 [Note] Shutting down plugin 'ngram'
gogsdb     | 2018-04-19T23:06:22.650071Z 0 [Note] Shutting down plugin 'partition'
gogsdb     | 2018-04-19T23:06:22.650086Z 0 [Note] Shutting down plugin 'BLACKHOLE'
gogsdb     | 2018-04-19T23:06:22.650099Z 0 [Note] Shutting down plugin 'ARCHIVE'
gogsdb     | 2018-04-19T23:06:22.650107Z 0 [Note] Shutting down plugin 'PERFORMANCE_SCHEMA'
gogsdb     | 2018-04-19T23:06:22.650130Z 0 [Note] Shutting down plugin 'MRG_MYISAM'
gogsdb     | 2018-04-19T23:06:22.650140Z 0 [Note] Shutting down plugin 'MyISAM'
gogsdb     | 2018-04-19T23:06:22.650156Z 0 [Note] Shutting down plugin 'INNODB_SYS_VIRTUAL'
gogsdb     | 2018-04-19T23:06:22.650166Z 0 [Note] Shutting down plugin 'INNODB_SYS_DATAFILES'
gogsdb     | 2018-04-19T23:06:22.650174Z 0 [Note] Shutting down plugin 'INNODB_SYS_TABLESPACES'
gogsdb     | 2018-04-19T23:06:22.650182Z 0 [Note] Shutting down plugin 'INNODB_SYS_FOREIGN_COLS'
gogsdb     | 2018-04-19T23:06:22.650189Z 0 [Note] Shutting down plugin 'INNODB_SYS_FOREIGN'
gogsdb     | 2018-04-19T23:06:22.650197Z 0 [Note] Shutting down plugin 'INNODB_SYS_FIELDS'
gogsdb     | 2018-04-19T23:06:22.650207Z 0 [Note] Shutting down plugin 'INNODB_SYS_COLUMNS'
gogsdb     | 2018-04-19T23:06:22.650215Z 0 [Note] Shutting down plugin 'INNODB_SYS_INDEXES'
gogsdb     | 2018-04-19T23:06:22.650222Z 0 [Note] Shutting down plugin 'INNODB_SYS_TABLESTATS'
gogsdb     | 2018-04-19T23:06:22.650230Z 0 [Note] Shutting down plugin 'INNODB_SYS_TABLES'
gogsdb     | 2018-04-19T23:06:22.650237Z 0 [Note] Shutting down plugin 'INNODB_FT_INDEX_TABLE'
gogsdb     | 2018-04-19T23:06:22.650244Z 0 [Note] Shutting down plugin 'INNODB_FT_INDEX_CACHE'
gogsdb     | 2018-04-19T23:06:22.650252Z 0 [Note] Shutting down plugin 'INNODB_FT_CONFIG'
gogsdb     | 2018-04-19T23:06:22.650259Z 0 [Note] Shutting down plugin 'INNODB_FT_BEING_DELETED'
gogsdb     | 2018-04-19T23:06:22.650267Z 0 [Note] Shutting down plugin 'INNODB_FT_DELETED'
gogsdb     | 2018-04-19T23:06:22.650274Z 0 [Note] Shutting down plugin 'INNODB_FT_DEFAULT_STOPWORD'
gogsdb     | 2018-04-19T23:06:22.650282Z 0 [Note] Shutting down plugin 'INNODB_METRICS'
gogsdb     | 2018-04-19T23:06:22.650289Z 0 [Note] Shutting down plugin 'INNODB_TEMP_TABLE_INFO'
gogsdb     | 2018-04-19T23:06:22.650297Z 0 [Note] Shutting down plugin 'INNODB_BUFFER_POOL_STATS'
gogsdb     | 2018-04-19T23:06:22.650315Z 0 [Note] Shutting down plugin 'INNODB_BUFFER_PAGE_LRU'
gogsdb     | 2018-04-19T23:06:22.650325Z 0 [Note] Shutting down plugin 'INNODB_BUFFER_PAGE'
gogsdb     | 2018-04-19T23:06:22.650333Z 0 [Note] Shutting down plugin 'INNODB_CMP_PER_INDEX_RESET'
gogsdb     | 2018-04-19T23:06:22.650340Z 0 [Note] Shutting down plugin 'INNODB_CMP_PER_INDEX'
gogsdb     | 2018-04-19T23:06:22.650348Z 0 [Note] Shutting down plugin 'INNODB_CMPMEM_RESET'
gogsdb     | 2018-04-19T23:06:22.650355Z 0 [Note] Shutting down plugin 'INNODB_CMPMEM'
gogsdb     | 2018-04-19T23:06:22.650362Z 0 [Note] Shutting down plugin 'INNODB_CMP_RESET'
gogsdb     | 2018-04-19T23:06:22.650370Z 0 [Note] Shutting down plugin 'INNODB_CMP'
gogsdb     | 2018-04-19T23:06:22.650377Z 0 [Note] Shutting down plugin 'INNODB_LOCK_WAITS'
gogsdb     | 2018-04-19T23:06:22.650385Z 0 [Note] Shutting down plugin 'INNODB_LOCKS'
gogsdb     | 2018-04-19T23:06:22.650392Z 0 [Note] Shutting down plugin 'INNODB_TRX'
gogsdb     | 2018-04-19T23:06:22.650400Z 0 [Note] Shutting down plugin 'InnoDB'
gogsdb     | 2018-04-19T23:06:22.650482Z 0 [Note] InnoDB: FTS optimize thread exiting.
gogsdb     | 2018-04-19T23:06:22.650628Z 0 [Note] InnoDB: Starting shutdown...
gogsdb     | 2018-04-19T23:06:22.751290Z 0 [Note] InnoDB: Dumping buffer pool(s) to /var/lib/mysql/ib_buffer_pool
gogsdb     | 2018-04-19T23:06:22.751663Z 0 [Note] InnoDB: Buffer pool(s) dump completed at 180419 23:06:22
gogsdb     | 2018-04-19T23:06:24.612325Z 0 [Note] InnoDB: Shutdown completed; log sequence number 12318903
gogsdb     | 2018-04-19T23:06:24.616924Z 0 [Note] InnoDB: Removed temporary tablespace data file: "ibtmp1"
gogsdb     | 2018-04-19T23:06:24.617008Z 0 [Note] Shutting down plugin 'MEMORY'
gogsdb     | 2018-04-19T23:06:24.617022Z 0 [Note] Shutting down plugin 'CSV'
gogsdb     | 2018-04-19T23:06:24.617040Z 0 [Note] Shutting down plugin 'sha256_password'
gogsdb     | 2018-04-19T23:06:24.617049Z 0 [Note] Shutting down plugin 'mysql_native_password'
gogsdb     | 2018-04-19T23:06:24.617157Z 0 [Note] Shutting down plugin 'binlog'
gogsdb     | 2018-04-19T23:06:24.618247Z 0 [Note] mysqld: Shutdown complete
gogsdb     | 
gogsdb     | 
gogsdb     | MySQL init process done. Ready for start up.
gogsdb     | 
gogsapp_1  | 2018/04/19 23:06:24 [ WARN] Custom config '/data/gogs/conf/app.ini' not found, ignore this if you're running first time
gogsapp_1  | 2018/04/19 23:06:24 [TRACE] Custom path: /data/gogs
gogsapp_1  | 2018/04/19 23:06:24 [TRACE] Log path: /app/gogs/log
gogsapp_1  | 2018/04/19 23:06:24 [TRACE] Build Time: 2018-04-19 07:15:14 UTC
gogsapp_1  | 2018/04/19 23:06:24 [TRACE] Build Git Hash: 
gogsapp_1  | 2018/04/19 23:06:24 [TRACE] Log Mode: Console (Trace)
gogsapp_1  | 2018/04/19 23:06:24 [ INFO] Gogs 0.11.46.0418
gogsapp_1  | 2018/04/19 23:06:24 [ INFO] Cache Service Enabled
gogsapp_1  | 2018/04/19 23:06:24 [ INFO] Session Service Enabled
gogsapp_1  | 2018/04/19 23:06:24 [ INFO] SQLite3 Supported
gogsapp_1  | 2018/04/19 23:06:24 [ INFO] Run Mode: Development
```

Clips

![屏幕快照 2018-04-19 下午4.13.06.png](./屏幕快照%202018-04-19%20下午4.13.06.png)


Documentation

1. https://github.com/nanoninja/docker-gogs-mysql

1. https://github.com/gogits/gogs
