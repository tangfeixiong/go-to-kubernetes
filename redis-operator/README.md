

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:$HOME/go make go-install
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go make go-install
github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/server
github.com/tangfeixiong/go-to-kubernetes/redis-operator/cmd
github.com/tangfeixiong/go-to-kubernetes/redis-operator
```

```
[vagrant@kubedev-172-17-4-59 redis-operator]$ ls -lh $GOPATH/bin/redis-operator 
-rwxr-xr-x. 1 vagrant vagrant 66M Jan  2 22:40 /Users/fanhongling/Downloads/workspace/bin/redis-operator
```