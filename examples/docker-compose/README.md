# 使用Docker Compose部署开发和测试环境 － Continuosuly Delivery and Deploy for QA within docker-compose

Demo project https://github.com/tangfeixiong/osev3-examples/tree/cicd/spring-boot/sample-microservices-springboot

Docker Compose file https://github.com/tangfeixiong/osev3-examples/blob/cicd/spring-boot/sample-microservices-springboot/docker-compose.yml

## For CI/CD and DevOps team

在开发和测试主机上安装git，[安装docker-compose](https://github.com/docker/compose/releases/tag/1.14.0)

Get yaml
```
$ git clone -b cicd https://github.com/tangfeixiong/osev3-examples tangfeixiong/osev3-examples
```

Compose up
```
sample-microservices-springboot$ docker-compose --project osev3-examples-$(date +%y%m%d-%H%M%S) up --no-build
```

## For QA

使用jenkins job，并将容器部署到QA指定的主机上

### 配置ssh password-less

使jenkins主机能无密码登录到QA主机，见[../../docs/ssh-login-passwordless.md](../../docs/ssh-login-passwordless.md)

### 为job使用docker run bash

如
```
ssh -i ~/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-compose/vagrant vagrant@172.17.4.50 "[[ -n $(docker ps -qf name=springboot-osev3-examples-web) ]] \
    && docker stop springboot-osev3-examples-web \
    && docker rm springboot-osev3-examples-web; \
docker run -d --name springboot-osev3-examples-web -p 8091:8091 172.17.4.50:5000/springboot-osev3-examples:web"
```

### 为job使用docker compose bash

如
```
ssh -i ~/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/docker-compose/vagrant vagrant@172.17.4.50 "if [[ -f tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/docker-compose.yml ]]; then \
    ( \
        set -o pipefail \
        && docker-compose --file tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/docker-compose.yml --project osev3-examples-springboot down \
    ); \
    rm -rf tangfeixiong/osev3-examples; \
fi \
&& git clone -b cicd https://github.com/tangfeixiong/osev3-examples tangfeixiong/osev3-examples \
&& docker-compose --file tangfeixiong/osev3-examples/spring-boot/sample-microservices-springboot/docker-compose.yml --project osev3-examples-springboot up --no-build"
```

### 为job使用jenkins pipeline

TBC