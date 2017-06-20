
## Prerequistes

VirtualBox vboxfs into vagrant
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster fanhongling$ cat ~/https%3A%2F%2Fwww.vagrantup.com/https%3A%2F%2Fatlas.hashicorp.com/boxes/ubuntu/trusty64/Vagrantfile | grep synced_folder
  # config.vm.synced_folder "../data", "/vagrant_data"
  # config.vm.synced_folder ".", "/vagrant", disabled: true
  config.vm.synced_folder "/Users/fanhongling/Downloads/99-mirror", "/Users/fanhongling/Downloads/99-mirror"
  config.vm.synced_folder "/Users/fanhongling/Downloads/99-mirror/linux-bin/kubernetes/v1.5.7", "/opt/kubernetes"
  config.vm.synced_folder "/Users/fanhongling/Downloads/go-kubernetes/pkg", "/Users/fanhongling/Downloads/go-kubernetes/pkg"
  config.vm.synced_folder "/Users/fanhongling/Downloads/go-kubernetes/src", "/Users/fanhongling/Downloads/go-kubernetes/src"
  config.vm.synced_folder "/Users/fanhongling/Downloads/go-openshift/pkg", "/Users/fanhongling/Downloads/go-openshift/pkg"
  config.vm.synced_folder "/Users/fanhongling/Downloads/go-openshift/src", "/Users/fanhongling/Downloads/go-openshift/src"
  config.vm.synced_folder "/Users/fanhongling/Downloads/workspace", "/Users/fanhongling/Downloads/workspace"
  config.vm.synced_folder "/Users/fanhongling/go/pkg", "/Users/fanhongling/go/pkg"
  config.vm.synced_folder "/Users/fanhongling/go/src", "/Users/fanhongling/go/src"
```

Before
```
[vagrant@localhost gitlab]$ sudo mkdir -p /srv/docker/gitlab/redis
[vagrant@localhost gitlab]$ sudo mkdir -p /srv/docker/gitlab/postgresql
[vagrant@localhost gitlab]$ sudo mkdir -p /srv/docker/gitlab/gitlab
```

### Installation of `docker-compose`

Refer to https://docs.docker.com/compose/install/

```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster fanhongling$ curl -jkSL https://github.com/docker/compose/releases/download/1.14.0-rc2/docker-compose-Linux-x86_64 > /Users/fanhongling/Downloads/99-mirror/linux-bin/docker-compose
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   617    0   617    0     0     97      0 --:--:--  0:00:06 --:--:--   162
  0     0    0     0    0     0      0      0 --:--:--  0:01:22 --:--:--     0curl: (7) Failed to connect to github-production-release-asset-2e65be.s3.amazonaws.com port 443: Operation timed out
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster fanhongling$ curl -jkSL https://github.com/docker/compose/releases/download/1.14.0-rc2/docker-compose-Linux-x86_64 > /Users/fanhongling/Downloads/99-mirror/linux-bin/docker-compose
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   617    0   617    0     0     34      0 --:--:--  0:00:17 --:--:--   149
 74 8084k   74 5983k    0     0   6510      0  0:21:11  0:15:41  0:05:30     0curl: (56) SSLRead() return error -9806
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster fanhongling$ curl -jkSL -C- https://github.com/docker/compose/releases/download/1.14.0-rc2/docker-compose-Linux-x86_64 -o /Users/fanhongling/Downloads/99-mirror/linux-bin/docker-compose
** Resuming transfer from byte position 6127170
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   617    0   617    0     0     15      0 --:--:--  0:00:38 --:--:--   151
  0     0    0     0    0     0      0      0 --:--:--  0:01:13 --:--:--     0curl: (6) Could not resolve host: github-production-release-asset-2e65be.s3.amazonaws.com
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster fanhongling$ curl -jkSL -C- https://github.com/docker/compose/releases/download/1.14.0-rc2/docker-compose-Linux-x86_64 -o /Users/fanhongling/Downloads/99-mirror/linux-bin/docker-compose
** Resuming transfer from byte position 6127170
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   617    0   617    0     0     14      0 --:--:--  0:00:42 --:--:--   133
  0 2100k    0 16347    0     0     13      0 45:57:33  0:19:47 45:37:46     0curl: (56) SSLRead() return error -36
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster fanhongling$ curl -jkSL -C- https://github.com/docker/compose/releases/download/1.14.0-rc2/docker-compose-Linux-x86_64 -o /Users/fanhongling/Downloads/99-mirror/linux-bin/docker-compose
** Resuming transfer from byte position 6143517
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   617    0   617    0     0    124      0 --:--:--  0:00:04 --:--:--   139
100 2084k  100 2084k    0     0  42061      0  0:00:50  0:00:50 --:--:--   98k
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster fanhongling$ chmod +x /Users/fanhongling/Downloads/99-mirror/linux-bin/docker-compose 
```

version
```
vagrant@vagrant-ubuntu-trusty-64:~$ docker-compose --version
docker-compose version 1.14.0-rc2, build 24dae73
```

Help
```
vagrant@vagrant-ubuntu-trusty-64:~$ docker-compose --help
Define and run multi-container applications with Docker.

Usage:
  docker-compose [-f <arg>...] [options] [COMMAND] [ARGS...]
  docker-compose -h|--help

Options:
  -f, --file FILE             Specify an alternate compose file (default: docker-compose.yml)
  -p, --project-name NAME     Specify an alternate project name (default: directory name)
  --verbose                   Show more output
  -v, --version               Print version and exit
  -H, --host HOST             Daemon socket to connect to

  --tls                       Use TLS; implied by --tlsverify
  --tlscacert CA_PATH         Trust certs signed only by this CA
  --tlscert CLIENT_CERT_PATH  Path to TLS certificate file
  --tlskey TLS_KEY_PATH       Path to TLS key file
  --tlsverify                 Use TLS and verify the remote
  --skip-hostname-check       Don't check the daemon's hostname against the name specified
                              in the client certificate (for example if your docker host
                              is an IP address)
  --project-directory PATH    Specify an alternate working directory
                              (default: the path of the Compose file)

Commands:
  build              Build or rebuild services
  bundle             Generate a Docker bundle from the Compose file
  config             Validate and view the Compose file
  create             Create services
  down               Stop and remove containers, networks, images, and volumes
  events             Receive real time events from containers
  exec               Execute a command in a running container
  help               Get help on a command
  images             List images
  kill               Kill containers
  logs               View output from containers
  pause              Pause services
  port               Print the public port for a port binding
  ps                 List containers
  pull               Pull service images
  push               Push service images
  restart            Restart services
  rm                 Remove stopped containers
  run                Run a one-off command
  scale              Set number of containers for a service
  start              Start services
  stop               Stop services
  top                Display the running processes
  unpause            Unpause services
  up                 Create and start containers
  version            Show the Docker-Compose version information
```

Local 
```
fanhonglingdeMacBook-Pro:gitlab fanhongling$ mkdir https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster/
fanhonglingdeMacBook-Pro:gitlab fanhongling$ cd https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster/
```

Download 
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster fanhongling$ curl -jkSLO https://raw.githubusercontent.com/stackdocker/docker-gitlab/master/docker-compose.yml
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  3836  100  3836    0     0   2430      0  0:00:01  0:00:01 --:--:--  2430
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster fanhongling$ cat docker-compose.yml 
version: '2'

services:
  redis:
    restart: always
    image: sameersbn/redis:latest
    command:
    - --loglevel warning
    volumes:
    - /srv/docker/gitlab/redis:/var/lib/redis:Z

  postgresql:
    restart: always
    image: sameersbn/postgresql:9.6-2
    volumes:
    - /srv/docker/gitlab/postgresql:/var/lib/postgresql:Z
    environment:
    - DB_USER=gitlab
    - DB_PASS=password
    - DB_NAME=gitlabhq_production
    - DB_EXTENSION=pg_trgm

  gitlab:
    restart: always
    image: sameersbn/gitlab:9.2.2
    depends_on:
    - redis
    - postgresql
    ports:
    - "10080:80"
    - "10022:22"
    volumes:
    - /srv/docker/gitlab/gitlab:/home/git/data:Z
    environment:
    - DEBUG=false

    - DB_ADAPTER=postgresql
    - DB_HOST=postgresql
    - DB_PORT=5432
    - DB_USER=gitlab
    - DB_PASS=password
    - DB_NAME=gitlabhq_production

    - REDIS_HOST=redis
    - REDIS_PORT=6379

    - TZ=Asia/Kolkata
    - GITLAB_TIMEZONE=Kolkata

    - GITLAB_HTTPS=false
    - SSL_SELF_SIGNED=false

    - GITLAB_HOST=localhost
    - GITLAB_PORT=10080
    - GITLAB_SSH_PORT=10022
    - GITLAB_RELATIVE_URL_ROOT=
    - GITLAB_SECRETS_DB_KEY_BASE=long-and-random-alphanumeric-string
    - GITLAB_SECRETS_SECRET_KEY_BASE=long-and-random-alphanumeric-string
    - GITLAB_SECRETS_OTP_KEY_BASE=long-and-random-alphanumeric-string

    - GITLAB_ROOT_PASSWORD=
    - GITLAB_ROOT_EMAIL=

    - GITLAB_NOTIFY_ON_BROKEN_BUILDS=true
    - GITLAB_NOTIFY_PUSHER=false

    - GITLAB_EMAIL=notifications@example.com
    - GITLAB_EMAIL_REPLY_TO=noreply@example.com
    - GITLAB_INCOMING_EMAIL_ADDRESS=reply@example.com

    - GITLAB_BACKUP_SCHEDULE=daily
    - GITLAB_BACKUP_TIME=01:00

    - SMTP_ENABLED=false
    - SMTP_DOMAIN=www.example.com
    - SMTP_HOST=smtp.gmail.com
    - SMTP_PORT=587
    - SMTP_USER=mailer@example.com
    - SMTP_PASS=password
    - SMTP_STARTTLS=true
    - SMTP_AUTHENTICATION=login

    - IMAP_ENABLED=false
    - IMAP_HOST=imap.gmail.com
    - IMAP_PORT=993
    - IMAP_USER=mailer@example.com
    - IMAP_PASS=password
    - IMAP_SSL=true
    - IMAP_STARTTLS=false

    - OAUTH_ENABLED=false
    - OAUTH_AUTO_SIGN_IN_WITH_PROVIDER=
    - OAUTH_ALLOW_SSO=
    - OAUTH_BLOCK_AUTO_CREATED_USERS=true
    - OAUTH_AUTO_LINK_LDAP_USER=false
    - OAUTH_AUTO_LINK_SAML_USER=false
    - OAUTH_EXTERNAL_PROVIDERS=

    - OAUTH_CAS3_LABEL=cas3
    - OAUTH_CAS3_SERVER=
    - OAUTH_CAS3_DISABLE_SSL_VERIFICATION=false
    - OAUTH_CAS3_LOGIN_URL=/cas/login
    - OAUTH_CAS3_VALIDATE_URL=/cas/p3/serviceValidate
    - OAUTH_CAS3_LOGOUT_URL=/cas/logout

    - OAUTH_GOOGLE_API_KEY=
    - OAUTH_GOOGLE_APP_SECRET=
    - OAUTH_GOOGLE_RESTRICT_DOMAIN=

    - OAUTH_FACEBOOK_API_KEY=
    - OAUTH_FACEBOOK_APP_SECRET=

    - OAUTH_TWITTER_API_KEY=
    - OAUTH_TWITTER_APP_SECRET=

    - OAUTH_GITHUB_API_KEY=
    - OAUTH_GITHUB_APP_SECRET=
    - OAUTH_GITHUB_URL=
    - OAUTH_GITHUB_VERIFY_SSL=

    - OAUTH_GITLAB_API_KEY=
    - OAUTH_GITLAB_APP_SECRET=

    - OAUTH_BITBUCKET_API_KEY=
    - OAUTH_BITBUCKET_APP_SECRET=

    - OAUTH_SAML_ASSERTION_CONSUMER_SERVICE_URL=
    - OAUTH_SAML_IDP_CERT_FINGERPRINT=
    - OAUTH_SAML_IDP_SSO_TARGET_URL=
    - OAUTH_SAML_ISSUER=
    - OAUTH_SAML_LABEL="Our SAML Provider"
    - OAUTH_SAML_NAME_IDENTIFIER_FORMAT=urn:oasis:names:tc:SAML:2.0:nameid-format:transient
    - OAUTH_SAML_GROUPS_ATTRIBUTE=
    - OAUTH_SAML_EXTERNAL_GROUPS=
    - OAUTH_SAML_ATTRIBUTE_STATEMENTS_EMAIL=
    - OAUTH_SAML_ATTRIBUTE_STATEMENTS_NAME=
    - OAUTH_SAML_ATTRIBUTE_STATEMENTS_FIRST_NAME=
    - OAUTH_SAML_ATTRIBUTE_STATEMENTS_LAST_NAME=

    - OAUTH_CROWD_SERVER_URL=
    - OAUTH_CROWD_APP_NAME=
    - OAUTH_CROWD_APP_PASSWORD=

    - OAUTH_AUTH0_CLIENT_ID=
    - OAUTH_AUTH0_CLIENT_SECRET=
    - OAUTH_AUTH0_DOMAIN=

    - OAUTH_AZURE_API_KEY=
    - OAUTH_AZURE_API_SECRET=
    - OAUTH_AZURE_TENANT_ID=
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab/https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster$ sed -i 's%TZ=Asia/Kolkata%TZ=Asia/Shanghai%;s/GITLAB_TIMEZONE=Kolkata/GITLAB_TIMEZONE=UTC/' docker-compose.yml 
```

Docker Image
```
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull sameersbn/redis
Using default tag: latest
latest: Pulling from sameersbn/redis
c60055a51d74: Already exists 
755da0cdb7d2: Already exists 
969d017f67e6: Already exists 
37c9a9113595: Already exists 
a3d9f8479786: Already exists 
e43d9de53575: Already exists 
be39bc918d43: Pull complete 
9fdc82c2ad69: Pull complete 
Digest: sha256:5c587d60c512504c362a4e833c90eab6c783e2b5a8c5c5994872705fa86daa6e
Status: Downloaded newer image for sameersbn/redis:latest
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab$ docker pull sameersbn/postgresql:9.6-2
9.6-2: Pulling from sameersbn/postgresql
c60055a51d74: Already exists 
755da0cdb7d2: Already exists 
969d017f67e6: Already exists 
37c9a9113595: Already exists 
a3d9f8479786: Already exists 
e43d9de53575: Already exists 
cddf24084b61: Pull complete 
f23b95c5f17c: Pull complete 
3965edbc705f: Pull complete 
4e6c393fb7f5: Pull complete 
Digest: sha256:884ccbac79c1d3daa39c07e7e1f2e5cc11ce18e2c5e41ac146258b9715f1d6c4
Status: Downloaded newer image for sameersbn/postgresql:9.6-2
vagrant@vagrant-ubuntu-trusty-64:~$ docker pull sameersbn/gitlab:9.2.2
9.2.2: Pulling from sameersbn/gitlab
8f229c550c2e: Pull complete 
8e1fb71e8df6: Pull complete 
f75a34586856: Pull complete 
8744e322b832: Pull complete 
d5165bfce78f: Pull complete 
73988ed74490: Pull complete 
8c2f0b5b704f: Pull complete 
a586098e0232: Pull complete 
4f059359e123: Pull complete 
af63213f4e40: Pull complete 
b9e3d0a7d2b6: Pull complete 
Digest: sha256:fa16ef20a45ff12bab3ff04ebeb66e2fdb34981ba4122bed10db81d67ac50a59
Status: Downloaded newer image for sameersbn/gitlab:9.2.2
```

Up (gitlab 9.2.2)
```
vagrant@vagrant-ubuntu-trusty-64:~$ cd /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab/
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab$ docker-compose -p gitlab -f https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster/docker-compose.yml up
Creating network "gitlab_default" with the default driver
Creating gitlab_redis_1 ... 
Creating gitlab_postgresql_1 ... 
Creating gitlab_redis_1
Creating gitlab_postgresql_1 ... done
Creating gitlab_gitlab_1 ... 
Creating gitlab_gitlab_1 ... done
Attaching to gitlab_redis_1, gitlab_postgresql_1, gitlab_gitlab_1
redis_1       | Starting redis-server...
redis_1       | [1] 08 Jun 17:36:05.568 # Server started, Redis version 2.8.4
redis_1       | [1] 08 Jun 17:36:05.573 # WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
postgresql_1  | Initializing datadir...
gitlab_1      | Initializing logdir...
gitlab_1      | Initializing datadir...
gitlab_1      | Installing configuration templates...
gitlab_1      | Configuring gitlab...
postgresql_1  | Initializing certdir...
postgresql_1  | Initializing logdir...
postgresql_1  | Initializing rundir...
postgresql_1  | Setting resolv.conf ACLs...
postgresql_1  | Creating database user: gitlab
postgresql_1  | Creating database: gitlabhq_production...
postgresql_1  | ‣ Loading pg_trgm extension...
postgresql_1  | ‣ Granting access to gitlab user...
postgresql_1  | Starting PostgreSQL 9.6...
postgresql_1  | LOG:  database system was shut down at 2017-06-08 17:36:08 UTC
postgresql_1  | LOG:  MultiXact member wraparound protections are now enabled
postgresql_1  | LOG:  database system is ready to accept connections
postgresql_1  | LOG:  autovacuum launcher started
gitlab_1      | Configuring gitlab::database
gitlab_1      | Configuring gitlab::redis
gitlab_1      | Configuring gitlab::secrets...
gitlab_1      | Configuring gitlab::sidekiq...
gitlab_1      | Configuring gitlab::gitlab-workhorse...
gitlab_1      | Configuring gitlab::unicorn...
gitlab_1      | Configuring gitlab::timezone...
gitlab_1      | Configuring gitlab::rack_attack...
gitlab_1      | Configuring gitlab::ci...
gitlab_1      | Configuring gitlab::artifacts...
gitlab_1      | Configuring gitlab::lfs...
gitlab_1      | Configuring gitlab::mattermost...
gitlab_1      | Configuring gitlab::project_features...
gitlab_1      | Configuring gitlab::oauth...
gitlab_1      | Configuring gitlab::ldap...
gitlab_1      | Configuring gitlab::backups...
gitlab_1      | Configuring gitlab::backups::schedule...
gitlab_1      | Configuring gitlab::registry...
gitlab_1      | Configuring gitlab::pages...
gitlab_1      | Configuring gitlab-shell...
gitlab_1      | Configuring nginx...
gitlab_1      | Configuring nginx::gitlab...
gitlab_1      | Migrating database...
gitlab_1      | Missing Rails.application.secrets.jws_private_key for production environment. The secret will be generated and stored in config/secrets.yml.
gitlab_1      | Clearing cache...
gitlab_1      | 2017-06-09 01:37:02,099 CRIT Supervisor running as root (no user in config file)
gitlab_1      | 2017-06-09 01:37:02,100 WARN Included extra file "/etc/supervisor/conf.d/nginx.conf" during parsing
gitlab_1      | 2017-06-09 01:37:02,100 WARN Included extra file "/etc/supervisor/conf.d/unicorn.conf" during parsing
gitlab_1      | 2017-06-09 01:37:02,100 WARN Included extra file "/etc/supervisor/conf.d/cron.conf" during parsing
gitlab_1      | 2017-06-09 01:37:02,101 WARN Included extra file "/etc/supervisor/conf.d/sidekiq.conf" during parsing
gitlab_1      | 2017-06-09 01:37:02,101 WARN Included extra file "/etc/supervisor/conf.d/mail_room.conf" during parsing
gitlab_1      | 2017-06-09 01:37:02,101 WARN Included extra file "/etc/supervisor/conf.d/sshd.conf" during parsing
gitlab_1      | 2017-06-09 01:37:02,101 WARN Included extra file "/etc/supervisor/conf.d/gitlab-workhorse.conf" during parsing
gitlab_1      | 2017-06-09 01:37:02,123 INFO RPC interface 'supervisor' initialized
gitlab_1      | 2017-06-09 01:37:02,124 CRIT Server 'unix_http_server' running without any HTTP authentication checking
gitlab_1      | 2017-06-09 01:37:02,124 INFO supervisord started with pid 1
gitlab_1      | 2017-06-09 01:37:03,127 INFO spawned: 'sidekiq' with pid 521
gitlab_1      | 2017-06-09 01:37:03,128 INFO spawned: 'unicorn' with pid 522
gitlab_1      | 2017-06-09 01:37:03,131 INFO spawned: 'gitlab-workhorse' with pid 523
gitlab_1      | 2017-06-09 01:37:03,133 INFO spawned: 'cron' with pid 524
gitlab_1      | 2017-06-09 01:37:03,143 INFO spawned: 'nginx' with pid 526
gitlab_1      | 2017-06-09 01:37:03,145 INFO spawned: 'sshd' with pid 527
gitlab_1      | 2017-06-09 01:37:04,653 INFO success: sidekiq entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-09 01:37:04,653 INFO success: unicorn entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-09 01:37:04,653 INFO success: gitlab-workhorse entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-09 01:37:04,653 INFO success: cron entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-09 01:37:04,653 INFO success: nginx entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-09 01:37:04,653 INFO success: sshd entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
```

Up (gitlb 8.16.3)
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab$ docker-compose -p gitlab -f https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster/docker-compose.yml up
Starting gitlab_postgresql_1 ... 
Starting gitlab_postgresql_1
Starting gitlab_redis_1 ... 
Starting gitlab_redis_1 ... done
Recreating gitlab_gitlab_1 ... 
Recreating gitlab_gitlab_1 ... done
Attaching to gitlab_postgresql_1, gitlab_redis_1, gitlab_gitlab_1
postgresql_1  | Initializing datadir...
postgresql_1  | Initializing certdir...
postgresql_1  | Initializing logdir...
postgresql_1  | Initializing rundir...
postgresql_1  | Setting resolv.conf ACLs...
redis_1       | Starting redis-server...
redis_1       | [1] 08 Jun 08:46:06.032 # Server started, Redis version 2.8.4
redis_1       | [1] 08 Jun 08:46:06.033 # WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
gitlab_1      | Initializing logdir...
gitlab_1      | Initializing datadir...
gitlab_1      | Installing configuration templates...
gitlab_1      | Configuring gitlab...
postgresql_1  | Creating database user: gitlab
postgresql_1  | Creating database: gitlabhq_production...
postgresql_1  | ‣ Loading pg_trgm extension...
postgresql_1  | ‣ Granting access to gitlab user...
postgresql_1  | Starting PostgreSQL 9.6...
postgresql_1  | LOG:  database system was shut down at 2017-06-08 08:46:16 UTC
postgresql_1  | LOG:  MultiXact member wraparound protections are now enabled
postgresql_1  | LOG:  database system is ready to accept connections
postgresql_1  | LOG:  autovacuum launcher started
gitlab_1      | Configuring gitlab::database..
gitlab_1      | Configuring gitlab::redis
gitlab_1      | Configuring gitlab::secrets...
gitlab_1      | Configuring gitlab::sidekiq...
gitlab_1      | Configuring gitlab::gitlab-workhorse...
gitlab_1      | Configuring gitlab::unicorn...
gitlab_1      | Configuring gitlab::timezone...
gitlab_1      | Configuring gitlab::rack_attack...
gitlab_1      | Configuring gitlab::ci...
gitlab_1      | Configuring gitlab::artifacts...
gitlab_1      | Configuring gitlab::lfs...
gitlab_1      | Configuring gitlab::mattermost...
gitlab_1      | Configuring gitlab::project_features...
gitlab_1      | Configuring gitlab::oauth...
gitlab_1      | Configuring gitlab::ldap...
gitlab_1      | Configuring gitlab::backups...
gitlab_1      | Configuring gitlab::backups::schedule...
gitlab_1      | Configuring gitlab::registry...
gitlab_1      | Configuring gitlab-shell...
gitlab_1      | Configuring nginx...
gitlab_1      | Configuring nginx::gitlab...
gitlab_1      | Setting up GitLab for firstrun. Please be patient, this could take a while...
postgresql_1  | ERROR:  database "gitlabhq_production" already exists
postgresql_1  | STATEMENT:  CREATE DATABASE "gitlabhq_production" ENCODING = 'unicode'
gitlab_1      | gitlabhq_production already exists
gitlab_1      | Migrating database...
gitlab_1      | Clearing cache...
gitlab_1      | 2017-06-08 16:47:25,070 CRIT Supervisor running as root (no user in config file)
gitlab_1      | 2017-06-08 16:47:25,070 WARN Included extra file "/etc/supervisor/conf.d/nginx.conf" during parsing
gitlab_1      | 2017-06-08 16:47:25,070 WARN Included extra file "/etc/supervisor/conf.d/unicorn.conf" during parsing
gitlab_1      | 2017-06-08 16:47:25,070 WARN Included extra file "/etc/supervisor/conf.d/cron.conf" during parsing
gitlab_1      | 2017-06-08 16:47:25,070 WARN Included extra file "/etc/supervisor/conf.d/sidekiq.conf" during parsing
gitlab_1      | 2017-06-08 16:47:25,070 WARN Included extra file "/etc/supervisor/conf.d/mail_room.conf" during parsing
gitlab_1      | 2017-06-08 16:47:25,070 WARN Included extra file "/etc/supervisor/conf.d/sshd.conf" during parsing
gitlab_1      | 2017-06-08 16:47:25,070 WARN Included extra file "/etc/supervisor/conf.d/gitlab-workhorse.conf" during parsing
gitlab_1      | 2017-06-08 16:47:25,096 INFO RPC interface 'supervisor' initialized
gitlab_1      | 2017-06-08 16:47:25,096 CRIT Server 'unix_http_server' running without any HTTP authentication checking
gitlab_1      | 2017-06-08 16:47:25,096 INFO supervisord started with pid 1
gitlab_1      | 2017-06-08 16:47:26,099 INFO spawned: 'sidekiq' with pid 503
gitlab_1      | 2017-06-08 16:47:26,100 INFO spawned: 'unicorn' with pid 504
gitlab_1      | 2017-06-08 16:47:26,105 INFO spawned: 'gitlab-workhorse' with pid 505
gitlab_1      | 2017-06-08 16:47:26,117 INFO spawned: 'cron' with pid 508
gitlab_1      | 2017-06-08 16:47:26,119 INFO spawned: 'nginx' with pid 509
gitlab_1      | 2017-06-08 16:47:26,125 INFO spawned: 'sshd' with pid 510
gitlab_1      | 2017-06-08 16:47:27,879 INFO success: sidekiq entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-08 16:47:27,879 INFO success: unicorn entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-08 16:47:27,879 INFO success: gitlab-workhorse entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-08 16:47:27,879 INFO success: cron entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-08 16:47:27,879 INFO success: nginx entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
gitlab_1      | 2017-06-08 16:47:27,880 INFO success: sshd entered RUNNING state, process has stayed up for > than 1 seconds (startsecs)
```

Try, push inoto Local GitLab
```
[vagrant@bogon osev3-examples]$ git remote -v
origin	https://github.com/tangfeixiong/osev3-examples (fetch)
origin	https://github.com/tangfeixiong/osev3-examples (push)
[vagrant@bogon osev3-examples]$ git remote add gitlab-local http://root@172.17.4.50:10080/root/osev3-examples.git
[vagrant@bogon osev3-examples]$ git remote -v
gitlab-local	http://root@172.17.4.50:10080/root/osev3-examples.git (fetch)
gitlab-local	http://root@172.17.4.50:10080/root/osev3-examples.git (push)
origin	https://github.com/tangfeixiong/osev3-examples (fetch)
origin	https://github.com/tangfeixiong/osev3-examples (push)
[vagrant@bogon osev3-examples]$ git push gitlab-local master
Password for 'http://root@172.17.4.50:10080': 
对象计数中: 342, 完成.
Delta compression using up to 2 threads.
压缩对象中: 100% (173/173), 完成.
写入对象中: 100% (342/342), 13.93 MiB | 172.00 KiB/s, 完成.
Total 342 (delta 108), reused 314 (delta 96)
remote: Resolving deltas: 100% (108/108), done.
To http://root@172.17.4.50:10080/root/osev3-examples.git
 * [new branch]      master -> master
```

Push _spring-guides/gs-rest-service_
```
fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git remote --verbose
origin	https://github.com/spring-guides/gs-rest-service (fetch)
origin	https://github.com/spring-guides/gs-rest-service (push)

fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git remote add gitlab-local http://172.17.4.200:10080/root/gs-rest-service

fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git remote --verbose
gitlab-local	http://172.17.4.200:10080/root/gs-rest-service (fetch)
gitlab-local	http://172.17.4.200:10080/root/gs-rest-service (push)
origin	https://github.com/spring-guides/gs-rest-service (fetch)
origin	https://github.com/spring-guides/gs-rest-service (push)

fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git pull gitlab-local
Username for 'http://172.17.4.200:10080': admin@example.com
Password for 'http://admin@example.com@172.17.4.200:10080': 
You asked to pull from the remote 'gitlab-local', but did not specify
a branch. Because this is not the default configured remote
for your current branch, you must specify a branch on the command line.


fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git checkout -b gitlab-local-demo
Switched to a new branch 'gitlab-local-demo'

fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git push gitlab-local
warning: push.default is unset; its implicit value has changed in
Git 2.0 from 'matching' to 'simple'. To squelch this message
and maintain the traditional behavior, use:

  git config --global push.default matching

To squelch this message and adopt the new behavior now, use:

  git config --global push.default simple

When push.default is set to 'matching', git will push local branches
to the remote branches that already exist with the same name.

Since Git 2.0, Git defaults to the more conservative 'simple'
behavior, which only pushes the current branch to the corresponding
remote branch that 'git pull' uses to update the current branch.

See 'git help config' and search for 'push.default' for further information.
(the 'simple' mode was introduced in Git 1.7.11. Use the similar mode
'current' instead of 'simple' if you sometimes use older versions of Git)

Counting objects: 1532, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (591/591), done.
Writing objects: 100% (1532/1532), 819.19 KiB | 0 bytes/s, done.
Total 1532 (delta 830), reused 1532 (delta 830)
remote: Resolving deltas: 100% (830/830), done.
To http://172.17.4.200:10080/root/gs-rest-service
 * [new branch]      gitlab-local-demo -> gitlab-local-demo
```

### Inspecting

Networking
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab$ docker network ls
NETWORK ID          NAME                DRIVER
0178d75935e1        gitlab_default      bridge              
352882a33218        bridge              bridge              
590c2ddbaedd        none                null                
14a2869a71b0        host                host                
```

Drop composer
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab$ docker-compose -p gitlab -f https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster/docker-compose.yml down
Stopping gitlab_gitlab_1 ... done
Stopping gitlab_redis_1 ... done
Stopping gitlab_postgresql_1 ... done
Removing gitlab_gitlab_1 ... done
Removing gitlab_redis_1 ... done
Removing gitlab_postgresql_1 ... done
Removing network gitlab_default
```

After drop down
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab$ docker network ls
NETWORK ID          NAME                DRIVER
590c2ddbaedd        none                null                
14a2869a71b0        host                host                
352882a33218        bridge              bridge              
```