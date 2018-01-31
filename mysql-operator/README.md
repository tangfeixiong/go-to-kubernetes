# Development Guide

## CI/CD

Artifacts
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go make go-bindata-artifacts
```

Image
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go make docker-
### snip: more packages...
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/initcnf
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/server
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/cmd
github.com/tangfeixiong/go-to-kubernetes/mysql-operator
#@docker build --no-cache --force-rm -t docker.io/tangfeixiong/mysql-operator:1801311008-gitb5fa67c ./
Sending build context to Docker daemon 20.36 MB
Step 1/7 : FROM busybox
 ---> 6ad733544a63
Step 2/7 : LABEL "maintainer" "tangfeixiong <tangfx128@gmail.com>" "project" "https://github.com/tangfeixiong/go-to-kubernetes" "name" "mysql-operator" "version" "0.1" "created-by" '{"go":"v1.9.2","kubernetes":"v1.8","docker":"1.13"}'
 ---> Running in f9b873e6fcea
 ---> cd87e8a21934
Removing intermediate container f9b873e6fcea
Step 3/7 : COPY bin/mysql-operator /
 ---> 82fcd8f235c5
Removing intermediate container 44dd257836db
Step 4/7 : ENV DOCKER_API_VERSION '1.12' DOCKER_CONFIG_JSON '{"auths": {"localhost:5000": {"auth": "","email": ""}}}' REGISTRY_CERTS_JSON '{"localhost:5000": {"ca_base64": "", "crt_base64": "", "key_base64": ""}}'
 ---> Running in e8c80357cafd
 ---> d9985c1aca75
Removing intermediate container e8c80357cafd
Step 5/7 : EXPOSE 10002 10001
 ---> Running in d41cd32bd854
 ---> c67d35e977bd
Removing intermediate container d41cd32bd854
Step 6/7 : ENTRYPOINT /mysql-operator serve
 ---> Running in cb91c65618ef
 ---> ea4905a6ad86
Removing intermediate container cb91c65618ef
Step 7/7 : CMD -v 2 --logtostderr=true
 ---> Running in b2344ac901d5
 ---> 73b28b1bb617
Removing intermediate container b2344ac901d5
Successfully built 73b28b1bb617
The push refers to a repository [docker.io/tangfeixiong/mysql-operator]
bbbd56bf89d5: Pushed 
0271b8eebde3: Mounted from tangfeixiong/redis-operator 
latest: digest: sha256:367926b68253d259d306374a8aca6c8cb8a1d568ab57b017b212253bc709bd8d size: 738
```

## Outside K8s

### Investigate custome resource

Install `spectest`
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go install -v ../cmd/spectest/
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/artifact
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/sts
github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/svc
github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/mysqlorbranches
github.com/tangfeixiong/go-to-kubernetes/cmd/spectest
```
crd
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ spectest create mariadb-crd -kubeconfig -v=5 -logtostderr
Created CRD "mysqlfamilies.example.com".
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl get crd mysqlfamilies.example.com
NAME                        AGE
mysqlfamilies.example.com   1m
```

```
[vagrant@rookdev-172-17-4-63 ~]$ kubectl delete crd mysqlfamilies.example.com        
customresourcedefinition "mysqlfamilies.example.com" deleted
```

local-storage
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ spectest create mariadb-localstorage -kubeconfig -v=5 -logtostderr
Created StorageClass "local-storage".
panic: PersistentVolume "example-local-pv" is invalid: [metadata.annotations: Forbidden: Storage node affinity is disabled by feature-gate, spec.local: Forbidden: Local volumes are disabled by feature-gate]

goroutine 1 [running]:
github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/mysqlorbranches.CreateLocalStorage(0x177c7a0, 0xc4204424e0)
	/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/mysqlorbranches/create.go:76 +0x814
main.main()
	/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/cmd/spectest/main.go:113 +0x921
```

```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ kubectl get storageclass
NAME            PROVISIONER                    AGE
local-storage   kubernetes.io/no-provisioner   1m
```
### Investigate initContainer
```
[vagrant@kubedev-172-17-4-59 mysql-operator]$ mysql-operator init --kubeconfig --logtostderr -v 5 --conf_dir=$HOME --name=my-galera
F0131 09:46:04.208810   19922 cnf.go:99] Get statefulset failed: statefulsets.apps "my-galera" not found
```




