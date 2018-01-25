/*
  Build:
    GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go install -v github.com/tangfeixiong/go-to-kubernetes/cmd/spectest-redis

  Run:
    spectest-redis create redis-sts -v 5 -logtostderr

  Check:
    kubectl get pods -l app=redis,component=redis,redis=my-redis -o wide
*/

package main
