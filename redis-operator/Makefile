
GOPATHP?=$(PWD)/../../../../..
GOPATHD?=$(HOME)/go

IMG_NS?=docker.io/tangfeixiong
IMG_REPO?=redis-operator
IMG_TAG?=latest

GIT_COMMIT=$(shell date +%y%m%d%H%M)-git$(shell git rev-parse --short=7 HEAD)

DOCKER_FILE?=Dockerfile.busybox
REGISTRY_HOST?=172.17.4.59:5000

all: protoc-grpc docker-push

protoc-grpc:
	@if [[ ! -e $(GOPATHD) ]] ; \
	then \
		ln -sf /Users/fanhongling/go $(GOPATHD) ; \
	fi

	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATHD}/src/github.com/gogo/protobuf \
		-I${GOPATHD}/src \
		--gogo_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,Mpb/fake.proto=github.com/tangfeixiong/go-to-kubernetes/redis-operator/pb,plugins=grpc:. \
		pb/service.proto pb/fake.proto
	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATHD}/src/github.com/gogo/protobuf \
		-I${GOPATHD}/src \
		--grpc-gateway_out=logtostderr=true:. \
		pb/service.proto

go-bindata:
	@pkg=webapp; src=static/...; output_file=pkg/ui/data/$${pkg}/datafile.go; \
		go-bindata -nocompress -o $${output_file} -prefix $${PWD} -pkg $${pkg} $${src}

go-build:
	@CGO_ENABLED=0 go build -a -v -o ./bin/redis-operator --installsuffix cgo -ldflags "-s" ./

docker-push: docker-build
	@docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

go-install:
	@go install -v ./

.PHONY: all protoc-grpc go-bindata go-build docker-build docker-push docker-build-exporter docker-build-collector docker-cgo docker-run go-install
