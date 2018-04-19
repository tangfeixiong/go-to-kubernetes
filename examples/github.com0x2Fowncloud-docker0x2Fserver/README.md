DevOps

Refer to https://github.com/owncloud-docker/server
```
[vagrant@kubedev-172-17-4-59 github.com0x2Fowncloud-docker0x2Fserver]$ curl -jkSLO https://raw.githubusercontent.com/owncloud-docker/server/master/docker-compose.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1762  100  1762    0     0   2309      0 --:--:-- --:--:-- --:--:--  2315
```

env
```
[vagrant@kubedev-172-17-4-59 github.com0x2Fowncloud-docker0x2Fserver]$ cat << EOF > .env
> VERSION=10.0
> DOMAIN=172.17.4.59
> ADMIN_USERNAME=admin
> ADMIN_PASSWORD=admin
> HTTP_PORT=28081
> HTTPS_PORT=44381
> EOF
```

docker-compose, refer to https://github.com/docker/compose/releases
```
[vagrant@kubedev-172-17-4-59 github.com0x2Fowncloud-docker0x2Fserver]$ curl -L https://github.com/docker/compose/releases/download/1.21.0/docker-compose-`uname -s`-`uname -m` -o /Users/fanhongling/Downloads/99-mirror/linux-bin/docker-compose 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   617    0   617    0     0    362      0 --:--:--  0:00:01 --:--:--   362
100 10.3M  100 10.3M    0     0  22718      0  0:07:58  0:07:58 --:--:-- 10806
```

up
```
[vagrant@kubedev-172-17-4-59 github.com0x2Fowncloud-docker0x2Fserver]$ docker-compose up
Creating network "githubcom0x2fowncloud-docker0x2fserver_default" with the default driver
Creating volume "githubcom0x2fowncloud-docker0x2fserver_files" with local driver
Creating volume "githubcom0x2fowncloud-docker0x2fserver_mysql" with local driver
Creating volume "githubcom0x2fowncloud-docker0x2fserver_backup" with local driver
Creating volume "githubcom0x2fowncloud-docker0x2fserver_redis" with local driver
Pulling db (webhippie/mariadb:latest)...
Trying to pull repository docker.io/webhippie/mariadb ... 
sha256:a25cf900c3b448636b3a85dc98944a92416fe9e31cd770cd4d11ca20b482c5b9: Pulling from docker.io/webhippie/mariadb
1eae7a7426b0: Pull complete
94b2629b2898: Pull complete
cd08385cd186: Pull complete
226190784501: Pull complete
15b1176bc8ef: Pull complete
Digest: sha256:a25cf900c3b448636b3a85dc98944a92416fe9e31cd770cd4d11ca20b482c5b9
Status: Downloaded newer image for docker.io/webhippie/mariadb:latest
Pulling redis (webhippie/redis:latest)...
Trying to pull repository docker.io/webhippie/redis ... 
sha256:c8f3f8f2a877b692e3298a9fe461d1f9335fff01115c0eca7fe778aa525779ab: Pulling from docker.io/webhippie/redis
1eae7a7426b0: Already exists
94b2629b2898: Already exists
cd08385cd186: Already exists
6f225a002295: Pull complete
e39b826b4dea: Pull complete
Digest: sha256:c8f3f8f2a877b692e3298a9fe461d1f9335fff01115c0eca7fe778aa525779ab
Status: Downloaded newer image for docker.io/webhippie/redis:latest
Pulling owncloud (owncloud/server:10.0)...
Trying to pull repository docker.io/owncloud/server ... 
sha256:5b38b0f5734484cf954a58e3e5418a7bfa91c44336b6f13f7a7a4cafedf8b87f: Pulling from docker.io/owncloud/server
e0a742c2abfd: Pull complete
486cb8339a27: Pull complete
dc6f0d824617: Pull complete
4f7a5649a30e: Pull complete
672363445ad2: Pull complete
5d096b9c72d0: Pull complete
84d0842af6a5: Pull complete
1a777208d429: Pull complete
fde7cdb2bd88: Pull complete
1e589ba689be: Pull complete
063ad84b8e51: Pull complete
05e31948afbd: Pull complete
Digest: sha256:5b38b0f5734484cf954a58e3e5418a7bfa91c44336b6f13f7a7a4cafedf8b87f
Status: Downloaded newer image for docker.io/owncloud/server:10.0
Creating githubcom0x2fowncloud-docker0x2fserver_redis_1 ... done
Creating githubcom0x2fowncloud-docker0x2fserver_db_1    ... done
Creating githubcom0x2fowncloud-docker0x2fserver_owncloud_1 ... done
Attaching to githubcom0x2fowncloud-docker0x2fserver_db_1, githubcom0x2fowncloud-docker0x2fserver_redis_1, githubcom0x2fowncloud-docker0x2fserver_owncloud_1
db_1        | Cron is enabled, launching it!
db_1        | Running mysql_install_db...
db_1        | Installing MariaDB/MySQL system tables in '/var/lib/mysql' ...
db_1        | 2018-04-19 21:42:10 140333224467336 [Warning] InnoDB: Using innodb_file_format is deprecated and the parameter may be removed in future releases. See https://mariadb.com/kb/en/library/xtradbinnodb-file-format/
db_1        | 2018-04-19 21:42:10 140333126654696 [Warning] Failed to load slave replication state from table mysql.gtid_slave_pos: 1146: Table 'mysql.gtid_slave_pos' doesn't exist
redis_1     | 10:C 19 Apr 21:42:11.343 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
redis_1     | 10:C 19 Apr 21:42:11.343 # Redis version=4.0.9, bits=64, commit=757d0c15, modified=0, pid=10, just started
redis_1     | 10:C 19 Apr 21:42:11.343 # Configuration loaded
redis_1     | 10:M 19 Apr 21:42:11.346 * Running mode=standalone, port=6379.
redis_1     | 10:M 19 Apr 21:42:11.351 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
redis_1     | 10:M 19 Apr 21:42:11.353 # Server initialized
redis_1     | 10:M 19 Apr 21:42:11.354 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
redis_1     | 10:M 19 Apr 21:42:11.364 * Ready to accept connections
owncloud_1  | Creating volume folders...
owncloud_1  | Creating hook folders...
owncloud_1  | Waiting for MySQL...
owncloud_1  | wait-for-it: waiting 180 seconds for db:3306
db_1        | OK
db_1        | 
db_1        | To start mysqld at boot time you have to copy
db_1        | support-files/mysql.server to the right place for your system
db_1        | 
db_1        | PLEASE REMEMBER TO SET A PASSWORD FOR THE MariaDB root USER !
db_1        | To do so, start the server, then issue the following commands:
db_1        | 
db_1        | '/usr/bin/mysqladmin' -u root password 'new-password'
db_1        | '/usr/bin/mysqladmin' -u root -h 1ff8f4ae7816 password 'new-password'
db_1        | 
db_1        | Alternatively you can run:
db_1        | '/usr/bin/mysql_secure_installation'
db_1        | 
db_1        | which will also give you the option of removing the test
db_1        | databases and anonymous user created by default.  This is
db_1        | strongly recommended for production servers.
db_1        | 
db_1        | See the MariaDB Knowledgebase at http://mariadb.com/kb or the
db_1        | MySQL manual for more instructions.
db_1        | 
db_1        | You can start the MariaDB daemon with:
db_1        | cd '/usr' ; /usr/bin/mysqld_safe --datadir='/var/lib/mysql'
db_1        | 
db_1        | You can test the MariaDB daemon with mysql-test-run.pl
db_1        | cd '/usr/mysql-test' ; perl mysql-test-run.pl
db_1        | 
db_1        | Please report any problems at http://mariadb.org/jira
db_1        | 
db_1        | The latest information about MariaDB is available at http://mariadb.org/.
db_1        | You can find additional information about the MySQL part at:
db_1        | http://dev.mysql.com
db_1        | Consider joining MariaDB's strong and vibrant community:
db_1        | https://mariadb.org/get-involved/
db_1        | 
db_1        | Finished mysql_install_db
db_1        | 180419 21:42:12 mysqld_safe Logging to '/var/lib/mysql/1ff8f4ae7816.err'.
db_1        | 180419 21:42:12 mysqld_safe Starting mysqld daemon with databases from /var/lib/mysql
owncloud_1  | wait-for-it: db:3306 is available after 2 seconds
owncloud_1  | Removing custom folder...
owncloud_1  | Linking custom folder...
owncloud_1  | Removing config folder...
owncloud_1  | Linking config folder...
owncloud_1  | Copying config file...
owncloud_1  | Copying domains file...
owncloud_1  | Copying db file...
owncloud_1  | Copying utf8mb4 config...
owncloud_1  | Fixing base perms...
owncloud_1  | Fixing data perms...
owncloud_1  | Installing server database...
owncloud_1  | ownCloud was successfully installed
owncloud_1  | Enabling files_external app...
owncloud_1  | ownCloud is already latest version
owncloud_1  | Generating default cert...
owncloud_1  | Generating a 4096 bit RSA private key
owncloud_1  | .............................................++
owncloud_1  | ............................++
owncloud_1  | writing new private key to '/mnt/data/certs/ssl-cert.key'
owncloud_1  | -----
owncloud_1  | Enforcing cli url...
owncloud_1  | System config value overwrite.cli.url set to string http://172.17.4.59/
owncloud_1  | Disabling update checks...
owncloud_1  | System config value updatechecker set to string false
owncloud_1  | Disabling memcached config...
owncloud_1  | System config value memcache.distributed deleted
owncloud_1  | System config value memcache.locking deleted
owncloud_1  | System config value memcached_servers deleted
owncloud_1  | Waiting for Redis...
owncloud_1  | wait-for-it: waiting 60 seconds for redis:6379
owncloud_1  | wait-for-it: redis:6379 is available after 0 seconds
owncloud_1  | Enabling redis config...
owncloud_1  | System config value memcache.distributed set to string \OC\Memcache\Redis
owncloud_1  | System config value memcache.locking set to string \OC\Memcache\Redis
owncloud_1  | System config value redis set to string REDIS
owncloud_1  | Enabling file locking...
owncloud_1  | System config value filelocking.enabled set to string true
owncloud_1  | Configure locale caching...
owncloud_1  | System config value memcache.local set to string \OC\Memcache\APCu
owncloud_1  | Disabling maintenace mode...
owncloud_1  | Maintenance mode disabled
owncloud_1  | Touching log file...
owncloud_1  | Setting log level...
owncloud_1  | System config value loglevel set to string 2
owncloud_1  | Setting default language...
owncloud_1  | System config value default_language set to string en
owncloud_1  | Enabling background cron...
owncloud_1  | Set mode for background jobs to 'cron'
owncloud_1  | Removing object storage...
owncloud_1  | Disabling object storage...
owncloud_1  | System config value objectstore deleted
owncloud_1  | Setting rewrite base...
owncloud_1  | System config value htaccess.RewriteBase set to string /
owncloud_1  | .htaccess has been updated
owncloud_1  | Writing apache config...
owncloud_1  | Writing php config...
owncloud_1  | Starting cron daemon...
owncloud_1  | Starting apache daemon...
owncloud_1  | Finishing ownCloud launch...
owncloud_1  | [Thu Apr 19 21:42:35.880004 2018] [ssl:warn] [pid 347] AH01906: localhost:443:0 server certificate is a CA certificate (BasicConstraints: CA == TRUE !?)
owncloud_1  | [Thu Apr 19 21:42:36.014978 2018] [ssl:warn] [pid 348] AH01906: localhost:443:0 server certificate is a CA certificate (BasicConstraints: CA == TRUE !?)
owncloud_1  | [Thu Apr 19 21:42:36.024138 2018] [mpm_prefork:notice] [pid 348] AH00163: Apache/2.4.18 (Ubuntu) OpenSSL/1.0.2g configured -- resuming normal operations
owncloud_1  | [Thu Apr 19 21:42:36.024173 2018] [core:notice] [pid 348] AH00094: Command line: '/usr/sbin/apache2'
owncloud_1  | ::1 - - [19/Apr/2018:21:42:41 +0000] "GET /status.php HTTP/1.1" 200 1120 "-" "curl/7.47.0"
owncloud_1  | ::1 - - [19/Apr/2018:21:43:12 +0000] "GET /status.php HTTP/1.1" 200 1112 "-" "curl/7.47.0"
owncloud_1  | ::1 - - [19/Apr/2018:21:43:42 +0000] "GET /status.php HTTP/1.1" 200 1124 "-" "curl/7.47.0"
owncloud_1  | ::1 - - [19/Apr/2018:21:44:12 +0000] "GET /status.php HTTP/1.1" 200 1120 "-" "curl/7.47.0"
owncloud_1  | ::1 - - [19/Apr/2018:21:44:42 +0000] "GET /status.php HTTP/1.1" 200 1118 "-" "curl/7.47.0"
owncloud_1  | ::1 - - [19/Apr/2018:21:45:12 +0000] "GET /status.php HTTP/1.1" 200 1110 "-" "curl/7.47.0"
owncloud_1  | ::1 - - [19/Apr/2018:21:45:42 +0000] "GET /status.php HTTP/1.1" 200 1112 "-" "curl/7.47.0"
owncloud_1  | ::1 - - [19/Apr/2018:21:46:12 +0000] "GET /status.php HTTP/1.1" 200 1130 "-" "curl/7.47.0"
owncloud_1  | ::1 - - [19/Apr/2018:21:46:42 +0000] "GET /status.php HTTP/1.1" 200 1116 "-" "curl/7.47.0"
```

clips

![屏幕快照 2018-04-19 下午2.49.57.png](./屏幕快照%202018-04-19%20下午2.49.57.png)

Documentation

1. https://doc.owncloud.org/server/latest/admin_manual/installation/docker/

1. https://docs.docker.com/compose/install/