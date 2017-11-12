# CI on Kubernetes

![web ui](./屏幕快照%202016-07-09%20下午2.04.43.png)

## Tables of content

* [Jul 09, 2016](#kubernetes-application-experimental)

* [Jun 21, 2016](#first-time-working-on-tidb)

## Kubernetes application experimental

* binary

Build bin

    [vagrant@localhost tidb]$ mkdir -p kubernetes-examples/docker-tidb

    [vagrant@localhost tidb]$ make server TARGET=$PWD/kubernetes-examples/docker-tidb/tidb-server
    go get github.com/qiuyesuifeng/goyacc
    go get github.com/qiuyesuifeng/golex
    goyacc -o /dev/null -xegen temp_parser_file parser/parser.y
    Parse table entries: 420495 of 1042566, x 16 bits == 840990 bytes
    goyacc -o parser/parser.go -xe temp_parser_file parser/parser.y 2>&1 | egrep "(shift|reduce)/reduce" | awk '{print} END {if (NR > 0) {print "Find conflict in parser.y. Please check y.output for more information."; system("rm -f temp_parser_file"); exit 1;}}'
    rm -f temp_parser_file
    rm -f y.output
    golex -o parser/scanner.go parser/scanner.l
    rm -rf vendor && ln -s _vendor/vendor vendor
    rm -rf vendor

Bin help

    [vagrant@localhost tidb]$ ./kubernetes-examples/docker-tidb/tidb-server --help
    Welcome to the TiDB.
    Version:
    Git Commit Hash: b3bd0dce4d547ebed68584472a112ed1f1890239
    UTC Build Time:  2016-07-09 09:27:39
    
    Usage of ./kubernetes-examples/docker-tidb/tidb-server:
      -L string
        	log level: info, debug, warn, error, fatal (default "debug")
      -P string
        	mp server port (default "4000")
      -lease int
        	schema lease seconds, very dangerous to change only if you know what you do (default 1)
      -path string
        	tidb storage path (default "/tmp/tidb")
      -socket string
        	The socket file to use for connection.
      -status string
        	tidb server status port (default "10080")
      -store string
        	registered store name, [memory, goleveldb, boltdb, tikv] (default "goleveldb")

* Docker image

Build

    [vagrant@localhost tidb]$ cd kubernetes-examples/docker-tidb
    [vagrant@localhost docker-tidb]$ vim Dockerfile 
    [vagrant@localhost docker-tidb]$ vim entry.sh
    [vagrant@localhost docker-tidb]$ chmod +x entry.sh 
    [vagrant@localhost docker-tidb]$ ls
    Dockerfile  entry.sh  tidb-server
    
    [vagrant@localhost docker-tidb]$ docker build -t tangfeixiong/tidb-server .
    Sending build context to Docker daemon 34.93 MB
    Step 1 : FROM golang:1.6
     ---> 8b9728127e88
    Step 2 : ADD tidb-server entry.sh /
     ---> 9f50b683262f
    Removing intermediate container f8914806ce8d
    Step 3 : VOLUME /tmp/tidb
     ---> Running in 60b1c396d38a
     ---> e3eb979d4fc3
    Removing intermediate container 60b1c396d38a
    Step 4 : EXPOSE 4000 10080
     ---> Running in b652ebfb2598
     ---> 1e290e337a96
    Removing intermediate container b652ebfb2598
    Step 5 : CMD /entry.sh
     ---> Running in f36c7fbdfb3a
     ---> 0556c3ecbd3f
    Removing intermediate container f36c7fbdfb3a
    Successfully built 0556c3ecbd3f

* Kubernetes app

Work with manifests

    [vagrant@localhost docker-tidb]$ cd ..
    [vagrant@localhost kubernetes-examples]$ mkdir k8s-manifests
    [vagrant@localhost kubernetes-examples]$ cd k8s-manifests
    [vagrant@localhost k8s-manifests]$ vim kidb-server-service.yaml
    
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon create -f kidb-server-service.yaml 
    service "tidb-server" created
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon get service --selector=app=tidb-server --show-labels
    NAME                 CLUSTER-IP   EXTERNAL-IP   PORT(S)              AGE       LABELS
    tidb-server          10.3.0.110   <none>        4000/TCP,10080/TCP   3m        app=tidb-server,developer=pingcap,heritage=qingyuancloud,project=tidb

    [vagrant@localhost k8s-manifests]$ vim tidb-server-controller.yaml
    
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon create -f tidb-server-controller.yaml 
    replicationcontroller "tidb-server" created
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon get rc -l app=tidb-server
    NAME          DESIRED   CURRENT   AGE
    tidb-server   1         1         27s
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon get pods -l name=tidb-server
    NAME                READY     STATUS    RESTARTS   AGE
    tidb-server-qp8y7   1/1       Running   0          39s
    
    [vagrant@localhost k8s-manifests]$ ls /tmp/tidb/
    000001.log  CURRENT  LOCK  LOG  MANIFEST-000000
    
    [vagrant@localhost k8s-manifests]$ curl http://10.3.0.110:10080/status;echo
    {"tps":0,"connections":0,"version":"5.5.31-TiDB-1.0"}

Investigate

    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon logs tidb-server-qp8y7
    2016/07/09 10:50:57 bg_worker.go:166: [debug] [ddl] wait 10s to check background job status again
    2016/07/09 10:50:57 txn.go:48: [debug] [kv] Begin txn:384843502761279488
    2016/07/09 10:50:57 txn.go:54: [debug] [kv] get key:6d 42 67 4a 6f 62 4f 77 6e ff 65 72 00 00 00 00 00 00 f9 00 00 00 00 00 00 00 73, txn:384843502761279488
    2016/07/09 10:50:57 txn.go:59: [debug] [kv] set key:6d 42 67 4a 6f 62 4f 77 6e ff 65 72 00 00 00 00 00 00 f9 00 00 00 00 00 00 00 73, txn:384843502761279488
    2016/07/09 10:50:57 ddl_worker.go:148: [debug] [ddl] become background job owner 3d45f653-a83b-4f76-98f5-4bf37cbd0098
    2016/07/09 10:50:57 txn.go:54: [debug] [kv] get key:6d 42 67 4a 6f 62 4c 69 73 ff 74 00 00 00 00 00 00 00 f8 00 00 00 00 00 00 00 4c, txn:384843502761279488
    2016/07/09 10:50:57 txn.go:113: [debug] [kv] commit txn 384843502761279488
    2016/07/09 10:50:57 ddl_worker.go:317: [debug] [ddl] wait 10s to check DDL status again
    2016/07/09 10:50:57 txn.go:48: [debug] [kv] Begin txn:384843502762590208
    2016/07/09 10:50:57 txn.go:54: [debug] [kv] get key:6d 44 44 4c 4a 6f 62 4f 77 ff 6e 65 72 00 00 00 00 00 fa 00 00 00 00 00 00 00 73, txn:384843502762590208
    2016/07/09 10:50:57 txn.go:59: [debug] [kv] set key:6d 44 44 4c 4a 6f 62 4f 77 ff 6e 65 72 00 00 00 00 00 fa 00 00 00 00 00 00 00 73, txn:384843502762590208
    2016/07/09 10:50:57 ddl_worker.go:148: [debug] [ddl] become ddl job owner 3d45f653-a83b-4f76-98f5-4bf37cbd0098
    2016/07/09 10:50:57 txn.go:54: [debug] [kv] get key:6d 44 44 4c 4a 6f 62 4c 69 ff 73 74 00 00 00 00 00 00 f9 00 00 00 00 00 00 00 4c, txn:384843502762590208
    2016/07/09 10:50:57 txn.go:113: [debug] [kv] commit txn 384843502762590208
    2016/07/09 10:51:27 compactor.go:179: [debug] [kv] GC send 2 keys to delete worker
    2016/07/09 10:51:27 compactor.go:179: [debug] [kv] GC send 17 keys to delete worker
    2016/07/09 10:51:27 ddl_worker.go:317: [debug] [ddl] wait 10s to check DDL status again
    2016/07/09 10:51:27 txn.go:48: [debug] [kv] Begin txn:384843510626123776
    2016/07/09 10:51:27 txn.go:54: [debug] [kv] get key:6d 44 44 4c 4a 6f 62 4f 77 ff 6e 65 72 00 00 00 00 00 fa 00 00 00 00 00 00 00 73, txn:384843510626123776
    2016/07/09 10:51:27 txn.go:59: [debug] [kv] set key:6d 44 44 4c 4a 6f 62 4f 77 ff 6e 65 72 00 00 00 00 00 fa 00 00 00 00 00 00 00 73, txn:384843510626123776
    2016/07/09 10:51:27 ddl_worker.go:148: [debug] [ddl] become ddl job owner 3d45f653-a83b-4f76-98f5-4bf37cbd0098
    2016/07/09 10:51:27 txn.go:54: [debug] [kv] get key:6d 44 44 4c 4a 6f 62 4c 69 ff 73 74 00 00 00 00 00 00 f9 00 00 00 00 00 00 00 4c, txn:384843510626123776
    2016/07/09 10:51:27 txn.go:113: [debug] [kv] commit txn 384843510626123776
    
    ...

* Using MySQL client

Docker run

    [vagrant@localhost k8s-manifests]$ docker run --rm -ti --name mysql-client docker.io/mysql /bin/bash
    root@fe826da9f935:/# mysql -h 10.3.0.110 -P 4000 -u root -D test
    Welcome to the MySQL monitor.  Commands end with ; or \g.
    Your MySQL connection id is 10001
    Server version: 5.5.31-TiDB-1.0 MySQL Community Server (GPL)
    
    Copyright (c) 2000, 2016, Oracle and/or its affiliates. All rights reserved.
    
    Oracle is a registered trademark of Oracle Corporation and/or its
    affiliates. Other names may be trademarks of their respective
    owners.
    
    Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
    
    mysql> show databases;
    +--------------------+
    | Database           |
    +--------------------+
    | INFORMATION_SCHEMA |
    | PERFORMANCE_SCHEMA |
    | mysql              |
    | test               |
    +--------------------+
    4 rows in set (0.00 sec)
    
    mysql> use test;
    Database changed
    mysql> create table t1 (
        ->   c1 INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
        ->   c2 VARCHAR(20));
    Query OK, 0 rows affected (0.03 sec)
    
    mysql> insert into t1 (c2) values ("hello");
    Query OK, 1 row affected (0.02 sec)
    
    mysql> select * from t1;
    +----+-------+
    | c1 | c2    |
    +----+-------+
    |  1 | hello |
    +----+-------+
    1 row in set (0.01 sec)
    
    mysql> quit
    Bye
    root@fe826da9f935:/# mysql -h 10.3.0.110 -P 4000 -u root -D test
    Reading table information for completion of table and column names
    You can turn off this feature to get a quicker startup with -A
    
    Welcome to the MySQL monitor.  Commands end with ; or \g.
    Your MySQL connection id is 10002
    Server version: 5.5.31-TiDB-1.0 MySQL Community Server (GPL)
    
    Copyright (c) 2000, 2016, Oracle and/or its affiliates. All rights reserved.
    
    Oracle is a registered trademark of Oracle Corporation and/or its
    affiliates. Other names may be trademarks of their respective
    owners.
    
    Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
    
    mysql> select * from test.t1;
    +----+-------+
    | c1 | c2    |
    +----+-------+
    |  1 | hello |
    +----+-------+
    1 row in set (0.01 sec)
    
    mysql> 

* Using phpmyadmin

Build image with Kubernetes service auto discovery

    [vagrant@localhost k8s-manifests]$ cd ..
    [vagrant@localhost kubernetes-examples]$ mkdir docker-phpmyadmin-kubernetes
    [vagrant@localhost kubernetes-examples]$ cd docker-phpmyadmin-kubernetes
    [vagrant@localhost docker-phpmyadmin-kubernetes]$ vim run-tidb-client.sh 
    [vagrant@localhost docker-phpmyadmin-kubernetes]$ vim Dockerfile.q8sad 
    [vagrant@localhost docker-phpmyadmin-kubernetes]$ docker build -t tangfeixiong/tidb-client-phpmyadmin:4.6.3 -f Dockerfile.q8sad .
    Sending build context to Docker daemon 10.33 MB
    Step 1 : FROM alpine:3.4
     ---> 31b45a1205be
    Step 2 : MAINTAINER tangfeixiong <fxtang@qingyuanos.com>
     ---> Using cache
     ---> 5725fae9db2b
    Step 3 : RUN apk add --no-cache php5-cli php5-mysqli php5-ctype php5-xml php5-gd php5-zlib php5-bz2 php5-zip php5-openssl php5-curl php5-opcache php5-json curl
     ---> Using cache
     ---> 6059de587ef3
    Step 4 : ENV PMA_REL_VER 4.6.3
     ---> Using cache
     ---> fe71580e5a66
    Step 5 : COPY run-tidb-client.sh /run.sh
     ---> 7aa81b5855ca
    Removing intermediate container f37888b90d52
    Step 6 : COPY *.tar.gz /tmp/
     ---> 55d56af3fbca
    Removing intermediate container 8cbe4ec01a5d
    Step 7 : WORKDIR /tmp
     ---> Running in 9fe02117857a
     ---> 72718736b5da
    Removing intermediate container 9fe02117857a
    Step 8 : RUN if [[ -f phpMyAdmin-${PMA_REL_VER}-all-languages.tar.gz ]]; then  tar -zxf phpMyAdmin-${PMA_REL_VER}-all-languages.tar.gz &&  rm phpMyAdmin-${PMA_REL_VER}-all-languages.tar.gz; else  curl --location https://www.phpmyadmin.net/downloads/phpMyAdmin-${PMA_REL_VER}-all-languages.tar.gz | tar xzf -   ; fi  && mv phpMyAdmin* /www  && rm -rf /www/js/jquery/src/ /www/examples /www/po/
     ---> Running in 57be86e375e7
     ---> 4ad23f663688
    Removing intermediate container 57be86e375e7
    Step 9 : COPY config.inc.php /www/
     ---> 133ee08d0f5d
    Removing intermediate container faf470b76712
    Step 10 : RUN chmod u+rwx /run.sh
     ---> Running in c99c5cd8199c
     ---> e56ed182025e
    Removing intermediate container c99c5cd8199c
    Step 11 : VOLUME /sessions
     ---> Running in f273ed149a2a
     ---> 0df89f79e74a
    Removing intermediate container f273ed149a2a
    Step 12 : EXPOSE 80
     ---> Running in f6162188b06d
     ---> ca827f3a1f28
    Removing intermediate container f6162188b06d
    Step 13 : ENV PHP_UPLOAD_MAX_FILESIZE 64M PHP_MAX_INPUT_VARS 2000
     ---> Running in d52908baa9e8
     ---> 6f41a42f872b
    Removing intermediate container d52908baa9e8
    Step 14 : ENTRYPOINT /run.sh
     ---> Running in dc5061d7354d
     ---> 644cd34f680e
    Removing intermediate container dc5061d7354d
    Successfully built 644cd34f680e

* K8s

Play with manifests

    [vagrant@localhost docker-phpmyadmin-kubernetes]$ cd ../k8s-manifests
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon get pods -l name=tidb-server
    NAME                READY     STATUS    RESTARTS   AGE
    tidb-server-qp8y7   1/1       Running   0          7h
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon exec -t -i tidb-server-qp8y7 -- printenv | grep TIDB
    TIDB_SERVER_STATUS_SERVICE_PORT=10080
    TIDB_SERVER_STATUS_SERVICE_HOST=10.3.0.225
    TIDB_SERVER_STATUS_PORT_10080_TCP=tcp://10.3.0.225:10080
    TIDB_SERVER_STATUS_PORT_10080_TCP_PROTO=tcp
    TIDB_SERVER_PORT_10080_TCP_PORT=10080
    TIDB_SERVER_SERVICE_PORT_TIDB_SERVER_STATUS=10080
    TIDB_SERVER_PORT_4000_TCP_PROTO=tcp
    TIDB_SERVER_PORT_4000_TCP_PORT=4000
    TIDB_SERVER_SERVICE_HOST=10.3.0.110
    TIDB_SERVER_PORT=tcp://10.3.0.110:4000
    TIDB_SERVER_PORT_4000_TCP_ADDR=10.3.0.110
    TIDB_SERVER_PORT_10080_TCP_PROTO=tcp
    TIDB_SERVER_STATUS_PORT=tcp://10.3.0.225:10080
    TIDB_SERVER_STATUS_SERVICE_PORT_WUI=10080
    TIDB_SERVER_STATUS_PORT_10080_TCP_ADDR=10.3.0.225
    TIDB_SERVER_SERVICE_PORT_TIDB_SERVER_MP=4000
    TIDB_SERVER_PORT_4000_TCP=tcp://10.3.0.110:4000
    TIDB_SERVER_PORT_10080_TCP=tcp://10.3.0.110:10080
    TIDB_SERVER_SERVICE_PORT=4000
    TIDB_SERVER_PORT_10080_TCP_ADDR=10.3.0.110
    TIDB_SERVER_STATUS_PORT_10080_TCP_PORT=10080

    [vagrant@localhost k8s-manifests]$ vim tidb-client-phpmyadmin-service.yaml
    [vagrant@localhost k8s-manifests]$ vim tidb-client-phpmyadmin-controller.yaml 

    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon create -f tidb-client-phpmyadmin-service.yaml 
    You have exposed your service on an external port on all nodes in your
    cluster.  If you want to expose this service to the external internet, you may
    need to set up firewall rules for the service port(s) (tcp:32132) to serve traffic.
    
    See http://releases.k8s.io/HEAD/docs/user-guide/services-firewalls.md for more details.
    service "tidb-client-phpmyadmin" created
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon get service -l app=tidb-client-phpmyadmin
    NAME                     CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
    tidb-client-phpmyadmin   10.3.0.7     <nodes>       80/TCP    1m


    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon create -f tidb-client-phpmyadmin-controller.yaml 
    replicationcontroller "tidb-client-phpmyadmin" created
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon get rc,pods -l app=tidb-client-phpmyadmin
    NAME                           DESIRED   CURRENT   AGE
    tidb-client-phpmyadmin         1         1         45s
    NAME                           READY     STATUS    RESTARTS   AGE
    tidb-client-phpmyadmin-r99hr   1/1       Running   0          45s
    
    [vagrant@localhost k8s-manifests]$ kubectl --namespace=harpoon logs tidb-client-phpmyadmin-8af05
    [Sat Jul  9 20:59:47 2016] 172.17.0.1:62867 [200]: /sql.php?server=1&db=test&table=t1&pos=0&token=6c9d5c6aa7a81881d475fdb89d2b202d&ajax_request=true&ajax_page_request=true&_nocache=1468097987708759339
    [Sat Jul  9 20:59:47 2016] 172.17.0.1:62869 [200]: /js/get_scripts.js.php?scripts%5B%5D=jquery/jquery.uitablefilter.js&scripts%5B%5D=tbl_change.js&scripts%5B%5D=gis_data_editor.js&scripts%5B%5D=multi_column_sort.js&scripts%5B%5D=makegrid.js&scripts%5B%5D=sql.js&call_done=1&v=4.6.3
    [Sat Jul  9 20:59:47 2016] 172.17.0.1:62870 [200]: /index.php?ajax_request=1&recent_table=1&token=6c9d5c6aa7a81881d475fdb89d2b202d&no_debug=true&_nocache=1468097991437865153
    [Sat Jul  9 20:59:47 2016] 172.17.0.1:62871 [200]: /themes/pmahomme/img/col_pointer.png

## First time working on TiDB

Fork and clone

    [vagrant@localhost pingcap]$ git clone --depth=1 https://github.com/tangfeixiong/tidb tidb
    正克隆到 'tidb'...
    remote: Counting objects: 959, done.
    remote: Compressing objects: 100% (847/847), done.
    remote: Total 959 (delta 78), reused 583 (delta 44), pack-reused 0
    接收对象中: 100% (959/959), 2.33 MiB | 117.00 KiB/s, 完成.
    处理 delta 中: 100% (78/78), 完成.
    检查连接... 完成。
    [vagrant@localhost pingcap]$ cd tidb/
    [vagrant@localhost tidb]$ git remote add upstream https://github.com/pingcap/tidb
    [vagrant@localhost tidb]$ git pull upstream
    remote: Counting objects: 22195, done.
    remote: Compressing objects: 100% (7324/7324), done.
    remote: Total 22195 (delta 15259), reused 21531 (delta 14601), pack-reused 0
    接收对象中: 100% (22195/22195), 10.27 MiB | 140.00 KiB/s, 完成.
    处理 delta 中: 100% (15259/15259), 完成 319 个本地对象.
    来自 https://github.com/pingcap/tidb
     * [新分支]          c4pt0r-patch-roadmap -> upstream/c4pt0r-patch-roadmap
     * [新分支]          c4pt0r/crocks-bridge -> upstream/c4pt0r/crocks-bridge
     * [新分支]          c4pt0r/hbase-metric -> upstream/c4pt0r/hbase-metric
     * [新分支]          c4pt0r/scan-bench -> upstream/c4pt0r/scan-bench
     * [新分支]          disksing/fast-commit -> upstream/disksing/fast-commit
     * [新分支]          disksing/tikv-columns -> upstream/disksing/tikv-columns
     * [新分支]          hanfei/apply -> upstream/hanfei/apply
     * [新分支]          hanfei/limit -> upstream/hanfei/limit
     * [新分支]          hanfei/orderby -> upstream/hanfei/orderby
     * [新分支]          master     -> upstream/master
     * [新分支]          zimuxia/tiny-cleanup -> upstream/zimuxia/tiny-cleanup
     * [新tag]           v0.5-alpha -> v0.5-alpha
    You asked to pull from the remote 'upstream', but did not specify
    a branch. Because this is not the default configured remote
    for your current branch, you must specify a branch on the command line.
    [vagrant@localhost tidb]$ git pull upstream master
    来自 https://github.com/pingcap/tidb
     * branch            master     -> FETCH_HEAD
    更新 fcfc5a7..b3bd0dc
    Fast-forward
     _vendor/Godeps/Godeps.json                               |  12 ++++--
    
    ...

Convert Golang version

    [vagrant@localhost tidb]$ go version
    go version go1.5.3 linux/amd64
    [vagrant@localhost tidb]$ . ~/go.rc
    [vagrant@localhost tidb]$ go version
    go version go1.6.2 linux/amd64

Failure build

    [vagrant@localhost tidb]$ ln -s _vendor/vendor vendor
    [vagrant@localhost tidb]$ go install -v ./...
    github.com/pingcap/tidb/vendor/github.com/juju/errors
    github.com/pingcap/tidb/vendor/github.com/ngaut/log
    github.com/pingcap/tidb/vendor/github.com/petar/GoLLRB/llrb
    github.com/pingcap/tidb/mysql
    github.com/pingcap/tidb/terror
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/util
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/cache
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/comparer
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/storage
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/errors
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/filter
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/iterator
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/journal
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/memdb
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/opt
    github.com/pingcap/tidb/vendor/github.com/golang/snappy
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb/table
    github.com/pingcap/tidb/vendor/github.com/syndtr/goleveldb/leveldb
    github.com/pingcap/tidb/kv
    github.com/pingcap/tidb/context
    github.com/pingcap/tidb/parser/opcode
    github.com/pingcap/tidb/vendor/golang.org/x/text/transform
    github.com/pingcap/tidb/vendor/golang.org/x/text/encoding
    github.com/pingcap/tidb/vendor/golang.org/x/text/encoding/internal/identifier
    github.com/pingcap/tidb/vendor/golang.org/x/text/encoding/internal
    github.com/pingcap/tidb/vendor/golang.org/x/text/encoding/charmap
    github.com/pingcap/tidb/vendor/golang.org/x/text/encoding/japanese
    github.com/pingcap/tidb/vendor/golang.org/x/text/encoding/korean
    github.com/pingcap/tidb/vendor/golang.org/x/text/encoding/simplifiedchinese
    github.com/pingcap/tidb/vendor/golang.org/x/text/encoding/traditionalchinese
    github.com/pingcap/tidb/vendor/golang.org/x/text/encoding/unicode
    github.com/pingcap/tidb/util/charset
    github.com/pingcap/tidb/util/hack
    github.com/pingcap/tidb/util/types
    github.com/pingcap/tidb/model
    github.com/pingcap/tidb/sessionctx/db
    github.com/pingcap/tidb/util/codec
    github.com/pingcap/tidb/util/distinct
    github.com/pingcap/tidb/ast
    github.com/pingcap/tidb/sessionctx/variable
    github.com/pingcap/tidb/util/stringutil
    github.com/pingcap/tidb/evaluator
    github.com/pingcap/tidb/vendor/github.com/golang/protobuf/proto
    github.com/pingcap/tidb/vendor/github.com/pingcap/tipb/go-tipb
    github.com/pingcap/tidb/xapi/tablecodec
    github.com/pingcap/tidb/plan/statistics
    github.com/pingcap/tidb/structure
    github.com/pingcap/tidb/meta
    github.com/pingcap/tidb/meta/autoid
    github.com/pingcap/tidb/table
    github.com/pingcap/tidb/util
    github.com/pingcap/tidb/table/tables
    github.com/pingcap/tidb/perfschema
    github.com/pingcap/tidb/infoschema
    github.com/pingcap/tidb/inspectkv
    github.com/pingcap/tidb/vendor/github.com/twinj/uuid
    github.com/pingcap/tidb/ddl
    github.com/pingcap/tidb/store/localstore/engine
    github.com/pingcap/tidb/util/bytes
    github.com/pingcap/tidb/util/segmentmap
    github.com/pingcap/tidb/xapi/xeval
    github.com/pingcap/tidb/store/localstore
    github.com/pingcap/tidb/domain
    github.com/pingcap/tidb/expression
    github.com/pingcap/tidb/parser
    # github.com/pingcap/tidb/parser
    parser/yy_parser.go:71: undefined: NewLexer
    parser/yy_parser.go:73: undefined: yyParse
    github.com/pingcap/tidb/privilege
    github.com/pingcap/tidb/sessionctx
    github.com/pingcap/tidb/sessionctx/autocommit
    github.com/pingcap/tidb/sessionctx/forupdate
    github.com/pingcap/tidb/util/sqlexec
    github.com/pingcap/tidb/xapi
    github.com/pingcap/tidb/vendor/github.com/rcrowley/go-metrics
    github.com/pingcap/tidb/metric
    github.com/pingcap/tidb/privilege/privileges
    github.com/pingcap/tidb/store/localstore/goleveldb
    github.com/pingcap/tidb/vendor/github.com/peterh/liner
    github.com/pingcap/tidb/util/printer
    github.com/pingcap/tidb/util/arena
    github.com/pingcap/tidb/store
    go build github.com/pingcap/tidb/store: no buildable Go source files in /data/src/github.com/pingcap/tidb/store
    github.com/pingcap/tidb/vendor/github.com/boltdb/bolt
    github.com/pingcap/tidb/store/localstore/boltdb
    github.com/pingcap/tidb/vendor/github.com/ngaut/deadline
    github.com/pingcap/tidb/vendor/github.com/ngaut/sync2
    github.com/pingcap/tidb/vendor/github.com/ngaut/pools
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/metapb
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/errorpb
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/kvrpcpb
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/coprocessor
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/raftpb
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/pdpb
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/raft_cmdpb
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/raft_serverpb
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/msgpb
    github.com/pingcap/tidb/vendor/github.com/pingcap/kvproto/pkg/util
    github.com/pingcap/tidb/vendor/github.com/gogo/protobuf/proto
    github.com/pingcap/tidb/vendor/github.com/coreos/etcd/auth/authpb
    github.com/pingcap/tidb/vendor/golang.org/x/net/context
    github.com/pingcap/tidb/vendor/golang.org/x/net/http2/hpack
    github.com/pingcap/tidb/vendor/golang.org/x/net/http2
    github.com/pingcap/tidb/vendor/golang.org/x/net/internal/timeseries
    github.com/pingcap/tidb/vendor/golang.org/x/net/trace
    github.com/pingcap/tidb/vendor/google.golang.org/grpc/codes
    github.com/pingcap/tidb/vendor/google.golang.org/grpc/credentials
    github.com/pingcap/tidb/vendor/google.golang.org/grpc/grpclog
    github.com/pingcap/tidb/vendor/google.golang.org/grpc/internal
    github.com/pingcap/tidb/vendor/google.golang.org/grpc/metadata
    github.com/pingcap/tidb/vendor/google.golang.org/grpc/naming
    github.com/pingcap/tidb/vendor/google.golang.org/grpc/peer
    github.com/pingcap/tidb/vendor/google.golang.org/grpc/transport
    github.com/pingcap/tidb/vendor/google.golang.org/grpc
    github.com/pingcap/tidb/vendor/github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes
    github.com/pingcap/tidb/vendor/github.com/coreos/etcd/mvcc/mvccpb
    github.com/pingcap/tidb/vendor/github.com/coreos/etcd/etcdserver/etcdserverpb
    github.com/pingcap/tidb/vendor/github.com/coreos/etcd/pkg/tlsutil
    github.com/pingcap/tidb/vendor/github.com/cloudfoundry-incubator/candiedyaml
    github.com/pingcap/tidb/vendor/github.com/ghodss/yaml
    github.com/pingcap/tidb/vendor/github.com/coreos/etcd/clientv3
    github.com/pingcap/tidb/vendor/github.com/pingcap/pd/pd-client
    github.com/pingcap/tidb/store/tikv/mock-tikv
    github.com/pingcap/tidb/store/tikv/oracle
    github.com/pingcap/tidb/store/tikv/oracle/oracles
    github.com/pingcap/tidb/store/tikv
    github.com/pingcap/tidb/util/format
    github.com/pingcap/tidb/util/mock
    github.com/pingcap/tidb/vendor/github.com/pingcap/check
    github.com/pingcap/tidb/util/testleak
    github.com/pingcap/tidb/util/testutil
    [vagrant@localhost tidb]$ unlink vendor

True way

    [vagrant@localhost tidb]$ make server TARGET=/data/bin/tidb-server
    go get github.com/qiuyesuifeng/goyacc
    go get github.com/qiuyesuifeng/golex
    goyacc -o /dev/null -xegen temp_parser_file parser/parser.y
    Parse table entries: 420495 of 1042566, x 16 bits == 840990 bytes
    goyacc -o parser/parser.go -xe temp_parser_file parser/parser.y 2>&1 | egrep "(shift|reduce)/reduce" | awk '{print} END {if (NR > 0) {print "Find conflict in parser.y. Please check y.output for more information."; system("rm -f temp_parser_file"); exit 1;}}'
    rm -f temp_parser_file
    rm -f y.output
    golex -o parser/scanner.go parser/scanner.l
    rm -rf vendor && ln -s _vendor/vendor vendor
    rm -rf vendor
    
    [vagrant@localhost tidb]$ ls /data/bin/
    cobra          genbashcomp         ginkgo      integration              kube-scheduler  openshift-deploy        osc                      tidb-server
    dapper         genconversion       go-bindata  kube-apiserver           lc-tlscert      openshift-docker-build  protoc-gen-go
    dockerbuilder  gendocs             godep       kube-controller-manager  linkcheck       openshift-recycle       rancher-catalog-service
    e2e_node.test  genkubedocs         golex       kubectl                  mungedocs       openshift-router        src
    e2e.test       genman              go-semver   kubelet                  oadm            openshiftstaging        tcbuilddockerimage
    etcd           genswaggertypedocs  gotty       kubemark                 oc              openshift-sti-build     tclogin
    etcdctl        getting-started-1   goyacc      kube-proxy               ociacibuilds    origin                  tcproject
    examples       getting-started-2   hyperkube   kubernetes               openshift       osadm                   tcprojectrequest

Test

    [vagrant@localhost tidb]$ tidb-server
    Welcome to the TiDB.
    Version:
    Git Commit Hash: b3bd0dce4d547ebed68584472a112ed1f1890239
    UTC Build Time:  2016-06-21 04:35:03
    
    2016/06/21 16:36:41 kv.go:290: [info] [kv] New store /tmp/tidb 
    2016/06/21 16:36:41 txn.go:48: [debug] [kv] Begin txn:384441254244515840
    2016/06/21 16:36:41 txn.go:54: [debug] [kv] get key:6d 53 63 68 65 6d 61 56 65 ff 72 73 69 6f 6e 4b 65 79 ff 00 00 00 00 00 00 00 00 f7 00 00 00 00 00 00 00 73, txn:384441254244515840
    2016/06/21 16:36:41 txn.go:69: [debug] [kv] seek key:6d 44 42 73 00 00 00 00 00 fa 00 00 00 00 00 00 00 68, txn:384441254244515840
    2016/06/21 16:36:41 domain.go:88: [info] [ddl] loadInfoSchema 0
    2016/06/21 16:36:41 txn.go:113: [debug] [kv] commit txn 384441254244515840
    
    ...
    
    ^Z  
    [1]+  已停止               tidb-server
    [vagrant@localhost tidb]$ bg %1
    [1]+ tidb-server &
    [vagrant@localhost tidb]$ sudo netstat -tpnl
    Active Internet connections (only servers)
    Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
    tcp        0      0 127.0.0.1:4001          0.0.0.0:*               LISTEN      4745/etcd           
    tcp        0      0 127.0.0.1:10248         0.0.0.0:*               LISTEN      4762/kubelet        
    tcp        0      0 127.0.0.1:10249         0.0.0.0:*               LISTEN      5119/hyperkube      
    tcp        0      0 172.17.4.50:10250       0.0.0.0:*               LISTEN      4762/kubelet        
    tcp        0      0 127.0.0.1:2379          0.0.0.0:*               LISTEN      4745/etcd           
    tcp        0      0 127.0.0.1:2380          0.0.0.0:*               LISTEN      4745/etcd           
    tcp        0      0 172.17.4.50:10255       0.0.0.0:*               LISTEN      4762/kubelet        
    tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      5247/hyperkube      
    tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      707/sshd            
    tcp        0      0 127.0.0.1:7001          0.0.0.0:*               LISTEN      4745/etcd           
    tcp        0      0 172.17.4.50:443         0.0.0.0:*               LISTEN      5247/hyperkube      
    tcp6       0      0 :::32124                :::*                    LISTEN      5119/hyperkube      
    tcp6       0      0 :::32125                :::*                    LISTEN      5119/hyperkube      
    tcp6       0      0 :::32126                :::*                    LISTEN      5119/hyperkube      
    tcp6       0      0 :::32127                :::*                    LISTEN      5119/hyperkube      
    tcp6       0      0 :::10080                :::*                    LISTEN      2624/tidb-server    
    tcp6       0      0 :::4000                 :::*                    LISTEN      2624/tidb-server    
    tcp6       0      0 :::4194                 :::*                    LISTEN      4762/kubelet        
    tcp6       0      0 :::10251                :::*                    LISTEN      5061/hyperkube      
    tcp6       0      0 :::10252                :::*                    LISTEN      5191/hyperkube      
    tcp6       0      0 :::30448                :::*                    LISTEN      5119/hyperkube      
    tcp6       0      0 :::22                   :::*                    LISTEN      707/sshd            
    tcp6       0      0 :::8888                 :::*                    LISTEN      2624/tidb-server    
    tcp6       0      0 :::32121                :::*                    LISTEN      5119/hyperkube      
    tcp6       0      0 :::32122                :::*                    LISTEN      5119/hyperkube      
    tcp6       0      0 :::32123                :::*                    LISTEN      5119/hyperkube      
    [vagrant@localhost tidb]$ fg %1
    # ^C to quit