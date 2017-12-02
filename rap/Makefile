
GOPATHP?=/Users/fanhongling/Downloads/workspace
GOPATHD?=/home/vagrant/go

IMG_NS?=docker.io/tangfeixiong
IMG_REPO?=rap
IMG_TAG?=latest
GIT_COMMIT=$(shell date +%y%m%d%H%M)-git$(shell git rev-parse --short=7 HEAD)
DOCKER_FILE?=Dockerfile.busybox
LOCAL_REGISTRY?=172.17.4.50:5000

all: protoc-grpc docker-push

protoc-grpc:
	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:$(GOPATHP)/src \
		$(GOPATHP)/src/github.com/tangfeixiong/go-to-kubernetes/rap/pb/server.proto
	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:$(GOPATHP)/src \
		$(GOPATHP)/src/github.com/tangfeixiong/go-to-kubernetes/rap/pb/server.proto


go-build:
	@CGO_ENABLED=0 go build -a -v -o $(GOPATHP)/src/github.com/tangfeixiong/go-to-kubernetes/rap/bin/rap --installsuffix cgo -ldflags "-s" ./

docker-build: go-build
	docker build -t $(IMG_NS)/$(IMG_REPO):$(IMG_TAG) ./

docker-push: docker-build
	docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

go-install:
	go install -v ./


.PHONY: all protoc-grpc go-build docker-build docker-push go-install
