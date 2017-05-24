Deploy Fabric8 into Kubernetes
==============================

Fabric8 is integrated development platform (iPaaS) for Kubernetes (OpenShift)


### Get bin

Directly download community release, for example: Mac版本程序软件
```
fanhonglingdeMacBook-Pro:bin fanhongling$ curl -jkSL -C- https://github.com/fabric8io/gofabric8/releases/download/v0.4.124/gofabric8-darwin-amd64 -o gofabric8
```

## Version 0.4.124

Make 
```
fanhonglingdeMacBook-Pro:gofabric8 fanhongling$ make bootstrap
GO15VENDOREXPERIMENT=1 go get -u github.com/Masterminds/glide
GO15VENDOREXPERIMENT=1 glide update --strip-vendor --strip-vcs --update-vendored
[WARN]	The --update-vendored flag is deprecated. This now works by default.
[WARN]	The --strip-vcs flag is deprecated. This now works by default.
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for github.com/daviddengcn/go-colortext
[INFO]	--> Fetching updates for github.com/spf13/cobra
[INFO]	--> Fetching updates for github.com/spf13/pflag
[INFO]	--> Fetching updates for github.com/golang/glog
[INFO]	--> Fetching updates for golang.org/x/crypto
[INFO]	--> Fetching updates for k8s.io/kubernetes
[INFO]	--> Fetching updates for github.com/Sirupsen/logrus
[INFO]	--> Fetching updates for github.com/coreos/pkg
[INFO]	--> Fetching updates for github.com/coreos/go-systemd
[INFO]	--> Fetching updates for speter.net/go/exp/math/dec/inf
[INFO]	--> Fetching updates for github.com/stretchr/testify
[INFO]	--> Fetching updates for github.com/pmezard/go-difflib
[INFO]	--> Fetching updates for github.com/pkg/browser
[INFO]	--> Fetching updates for github.com/jimmidyson/minishift
[INFO]	--> Fetching updates for github.com/docker/docker
[INFO]	--> Fetching updates for github.com/openshift/origin
[INFO]	--> Fetching updates for github.com/imdario/mergo
[INFO]	--> Fetching updates for github.com/fabric8io/fabric8-init-tenant
[INFO]	--> Fetching updates for github.com/gogo/protobuf
[INFO]	--> Fetching updates for github.com/gonum/graph
[INFO]	--> Fetching updates for github.com/dgrijalva/jwt-go
[INFO]	--> Fetching updates for github.com/jteeuwen/go-bindata
[INFO]	--> Fetching updates for github.com/elazarl/go-bindata-assetfs
[INFO]	--> Setting version for k8s.io/kubernetes to 52492b4bff99ef3b8ca617d385a3ff0612f9402d.
[INFO]	--> Setting version for github.com/spf13/pflag to c7e63cf4530bcd3ba943729cee0efeff2ebea63f.
[INFO]	--> Setting version for github.com/dgrijalva/jwt-go to 5ca80149b9d3f8b863af0e2bb6742e608603bd99.
[INFO]	--> Setting version for github.com/golang/glog to 335da9dda11408a34b64344f82e9c03779b71673.
[INFO]	--> Setting version for github.com/gonum/graph to bde6d0fbd9dec5a997e906611fe0364001364c41.
[INFO]	--> Setting version for github.com/openshift/origin to v1.3.0.
[INFO]	--> Setting version for github.com/gogo/protobuf to c0656edd0d9eab7c66d1eb0c568f9039345796f7.
[INFO]	--> Setting version for github.com/spf13/cobra to 9c28e4bbd74e5c3ed7aacbc552b2cab7cfdfe744.
[INFO]	--> Setting version for github.com/imdario/mergo to 6633656539c1639d9d78127b7d47c622b5d7b6dc.
[INFO]	--> Setting version for github.com/docker/docker to 0f5c9d301b9b1cca66b3ea0f9dec3b5317d3686d.
[INFO]	--> Setting version for golang.org/x/crypto to c84e1f8e3a7e322d497cd16c0e8a13c7e127baf3.
[INFO]	--> Detected semantic version. Setting version for github.com/jteeuwen/go-bindata to v3.0.7
[INFO]	--> Detected semantic version. Setting version for github.com/Sirupsen/logrus to v0.11.5
[INFO]	Resolving imports
[INFO]	--> Fetching updates for github.com/blang/semver
[INFO]	--> Detected semantic version. Setting version for github.com/blang/semver to v3.5.0
[INFO]	--> Fetching updates for github.com/google/go-github
[INFO]	--> Setting version for github.com/google/go-github to 30a21ee1a3839fb4a408efe331f226b73faac379.
[INFO]	--> Fetching updates for github.com/inconshreveable/go-update
[INFO]	--> Fetching updates for github.com/kardianos/osext
[INFO]	--> Fetching updates for github.com/minishift/minishift
[INFO]	--> Fetching updates for github.com/spf13/viper
[INFO]	--> Setting version for github.com/spf13/viper to 382f87b929b84ce13e9c8a375a4b217f224e6c65.
[INFO]	--> Fetching updates for gopkg.in/cheggaaa/pb.v1
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-github.com-kubernetes-minikube
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Fetching updates for k8s.io/minikube
[INFO]	--> Fetching updates for github.com/inconshreveable/mousetrap
[INFO]	--> Setting version for github.com/inconshreveable/mousetrap to 76626ae9c91c4f2a10f34cad8ce83ea42c93bb75.
[WARN]	Conflict: github.com/spf13/pflag rev is currently c7e63cf4530bcd3ba943729cee0efeff2ebea63f, but github.com/kubernetes/minikube wants 9ff6c6923cfffbcd502984b8e0c80539a94968b7
[INFO]	github.com/spf13/pflag reference c7e63cf4530bcd3ba943729cee0efeff2ebea63f:
[INFO] - author: Eric Paris <eparis@redhat.com>
[INFO] - commit date: Thu, 15 Sep 2016 11:31:01 -0400
[INFO] - subject (first line): Merge pull request #93 from bogem/flagStrings
[INFO]	github.com/spf13/pflag reference 9ff6c6923cfffbcd502984b8e0c80539a94968b7:
[INFO] - author: Ian Campbell <ijc25@users.noreply.github.com>
[INFO] - commit date: Mon, 30 Jan 2017 21:42:45 +0000
[INFO] - subject (first line): Add FlagSet.FlagUsagesWrapped(cols) which wraps to the given column (#105)
[INFO]	Keeping github.com/spf13/pflag c7e63cf4530bcd3ba943729cee0efeff2ebea63f
[INFO]	--> Fetching updates for github.com/fsnotify/fsnotify
[INFO]	--> Setting version for github.com/fsnotify/fsnotify to f12c6236fe7b5cf6bcf30e5935d08cb079d78334.
[INFO]	--> Fetching updates for github.com/hashicorp/hcl
[INFO]	--> Setting version for github.com/hashicorp/hcl to d8c773c4cba11b11539e3d45f93daeaa5dcf1fa1.
[INFO]	--> Fetching updates for github.com/magiconair/properties
[INFO]	--> Setting version for github.com/magiconair/properties to 61b492c03cf472e0c6419be5899b8e0dc28b1b88.
[INFO]	--> Fetching updates for github.com/mitchellh/mapstructure
[INFO]	--> Setting version for github.com/mitchellh/mapstructure to db1efb556f84b25a0a13a04aad883943538ad2e0.
[INFO]	--> Fetching updates for github.com/pelletier/go-toml
[INFO]	--> Setting version for github.com/pelletier/go-toml to 0049ab3dc4c4c70a9eee23087437b69c0dde2130.
[INFO]	--> Fetching updates for github.com/spf13/afero
[INFO]	--> Setting version for github.com/spf13/afero to b28a7effac979219c2a2ed6205a4d70e4b1bcd02.
[INFO]	--> Fetching updates for github.com/spf13/cast
[INFO]	--> Setting version for github.com/spf13/cast to e31f36ffc91a2ba9ddb72a4b6a607ff9b3d3cb63.
[INFO]	--> Fetching updates for github.com/spf13/jwalterweatherman
[INFO]	--> Setting version for github.com/spf13/jwalterweatherman to 33c24e77fb80341fe7130ee7c594256ff08ccc46.
[INFO]	--> Fetching updates for gopkg.in/yaml.v2
[INFO]	--> Setting version for gopkg.in/yaml.v2 to 53feefa2559fb8dfa8d81baad31be332c97d6c77.
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-github.com-openshift-kubernetes.git
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Fetching updates for github.com/emicklei/go-restful
[INFO]	--> Setting version for github.com/emicklei/go-restful to 09691a3b6378b740595c1002f40c34dd5f218a22.
[INFO]	--> Fetching updates for github.com/evanphx/json-patch
[INFO]	--> Setting version for github.com/evanphx/json-patch to ba18e35c5c1b36ef6334cad706eb681153d2d379.
[WARN]	Conflict: k8s.io/kubernetes rev is currently 52492b4bff99ef3b8ca617d385a3ff0612f9402d, but github.com/kubernetes/minikube wants 0480917b552be33e2dba47386e51decb1a211df6
[INFO]	k8s.io/kubernetes reference 52492b4bff99ef3b8ca617d385a3ff0612f9402d:
[INFO] - author: derekwaynecarr <decarr@redhat.com>
[INFO] - commit date: Wed, 06 Jul 2016 13:33:10 -0400
[INFO] - subject (first line): Fix httpclient setup for gcp credential provider to have timeout
[INFO]	Keeping k8s.io/kubernetes 52492b4bff99ef3b8ca617d385a3ff0612f9402d
[INFO]	--> Fetching updates for github.com/google/cadvisor
[INFO]	--> Setting version for github.com/google/cadvisor to 17543becf9053e7e80806a57b05002a88c79ec8a.
[INFO]	--> Fetching updates for github.com/davecgh/go-spew
[INFO]	--> Setting version for github.com/davecgh/go-spew to 5215b55f46b2b919f50a1df0eaa5886afe4e3b3d.
[INFO]	--> Fetching golang.org/x/net
[INFO]	--> Setting version for golang.org/x/net to e90d6d0afc4c315a0d87a568ae68577cc15149a0.
[WARN]	Conflict: github.com/gogo/protobuf rev is currently c0656edd0d9eab7c66d1eb0c568f9039345796f7, but github.com/kubernetes/minikube wants e18d7aa8f8c624c915db340349aad4c49b10d173
[INFO]	github.com/gogo/protobuf reference c0656edd0d9eab7c66d1eb0c568f9039345796f7:
[INFO] - author: Walter Schulze <awalterschulze@gmail.com>
[INFO] - commit date: Thu, 30 Mar 2017 09:10:51 +0200
[INFO] - subject (first line): Update Readme.md
[INFO]	github.com/gogo/protobuf reference e18d7aa8f8c624c915db340349aad4c49b10d173:
[INFO] - author: Walter Schulze <awalterschulze@gmail.com>
[INFO] - commit date: Mon, 18 Jul 2016 18:13:53 +0200
[INFO] - subject (first line): merged 8d92cf5fc15a4382f8964b08e1f42a75c0591aa3 from golang/protobuf
[INFO]	Keeping github.com/gogo/protobuf c0656edd0d9eab7c66d1eb0c568f9039345796f7
[INFO]	--> Fetching updates for gopkg.in/inf.v0
[INFO]	--> Setting version for gopkg.in/inf.v0 to 3887ee99ecf07df5b447e9b00d9c0b2adaa9f3e4.
[INFO]	--> Fetching updates for github.com/google/gofuzz
[INFO]	--> Setting version for github.com/google/gofuzz to 44d81051d367757e1c7c6a5a86423ece9afcf63c.
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-github.com-openshift-origin
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Fetching updates for github.com/almighty/almighty-core
[INFO]	--> Fetching updates for github.com/google/go-querystring
[INFO]	--> Fetching golang.org/x/oauth2
[INFO]	--> Setting version for golang.org/x/oauth2 to 3c3a985cb79f52a3190fbc056984415ca6763d01.
[INFO]	--> Fetching updates for github.com/mattn/go-runewidth
[INFO]	--> Setting version for github.com/mattn/go-runewidth to 737072b4e32b7a5018b4a7125da8d12de90e8045.
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-k8s.io-minikube
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Fetching updates for k8s.io/client-go
[INFO]	--> Fetching golang.org/x/sys
[INFO]	--> Setting version for golang.org/x/sys to 8f0908ab3b2457e2e15403d3697c9ef5cb4b57a9.
[INFO]	--> Fetching updates for github.com/pelletier/go-buffruneio
[INFO]	--> Setting version for github.com/pelletier/go-buffruneio to df1e16fde7fc330a0ca68167c23bf7ed6ac31d6d.
[INFO]	--> Fetching github.com/pkg/sftp
[INFO]	--> Setting version for github.com/pkg/sftp to 4d0e916071f68db74f8a73926335f809396d6b42.
[INFO]	--> Fetching golang.org/x/text
[INFO]	--> Setting version for golang.org/x/text to 2910a502d2bf9e43193af9d68ca516529614eed3.
[INFO]	--> Fetching updates for github.com/golang/groupcache
[INFO]	--> Setting version for github.com/golang/groupcache to 02826c3e79038b59d737d3b1c0a1d937f71a4433.
[INFO]	--> Fetching updates for github.com/ugorji/go
[INFO]	--> Setting version for github.com/ugorji/go to ded73eae5db7e7a0ef6f55aace87a2873c5d2b74.
[INFO]	--> Fetching updates for github.com/prometheus/client_golang
[INFO]	--> Setting version for github.com/prometheus/client_golang to e51041b3fa41cece0dca035740ba6411905be473.
[INFO]	--> Fetching updates for github.com/juju/ratelimit
[INFO]	--> Setting version for github.com/juju/ratelimit to 77ed1c8a01217656d2080ad51981f6e99adaa177.
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-github.com-google-cadvisor
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Fetching updates for github.com/opencontainers/runc
[INFO]	--> Setting version for github.com/opencontainers/runc to 45c30e75abfd52107b53048004a83165403ad0d1.
[INFO]	--> Fetching updates for github.com/pborman/uuid
[INFO]	--> Setting version for github.com/pborman/uuid to ca53cad383cad2479bbba7f7a1a05797ec1386e4.
[INFO]	--> Fetching updates for github.com/docker/distribution
[INFO]	--> Setting version for github.com/docker/distribution to cd27f179f2c10c5d300e6d09025b538c475b0d51.
[INFO]	--> Fetching updates for github.com/fsouza/go-dockerclient
[INFO]	--> Setting version for github.com/fsouza/go-dockerclient to bf97c77db7c945cbcdbf09d56c6f87a66f54537b.
[WARN]	Conflict: github.com/dgrijalva/jwt-go rev is currently 5ca80149b9d3f8b863af0e2bb6742e608603bd99, but github.com/kubernetes/minikube wants 01aeca54ebda6e0fbfafd0a524d234159c05ec20
[INFO]	github.com/dgrijalva/jwt-go reference 5ca80149b9d3f8b863af0e2bb6742e608603bd99:
[INFO] - author: Dave Grijalva <grijalva@gmail.com>
[INFO] - commit date: Mon, 20 Apr 2015 13:43:07 -0700
[INFO] - subject (first line): added reference to gopkg.in
[INFO]	github.com/dgrijalva/jwt-go reference 01aeca54ebda6e0fbfafd0a524d234159c05ec20:
[INFO] - author: Dave Grijalva <grijalva@gmail.com>
[INFO] - commit date: Tue, 05 Jul 2016 13:30:06 -0700
[INFO] - subject (first line): Merge pull request #146 from pkieltyka/master
[INFO]	Keeping github.com/dgrijalva/jwt-go 5ca80149b9d3f8b863af0e2bb6742e608603bd99
[WARN]	Conflict: github.com/docker/docker rev is currently 0f5c9d301b9b1cca66b3ea0f9dec3b5317d3686d, but github.com/kubernetes/minikube wants b9f10c951893f9a00865890a5232e85d770c1087
[INFO]	github.com/docker/docker reference 0f5c9d301b9b1cca66b3ea0f9dec3b5317d3686d:
[INFO] - author: Antonio Murdaca <runcom@linux.com>
[INFO] - commit date: Tue, 21 Jul 2015 19:49:42 +0200
[INFO] - subject (first line): pkg: mount: golint
[INFO]	github.com/docker/docker reference b9f10c951893f9a00865890a5232e85d770c1087 (v1.11.2):
[INFO] - author: Sebastiaan van Stijn <thaJeztah@users.noreply.github.com>
[INFO] - commit date: Wed, 01 Jun 2016 23:11:42 +0200
[INFO] - subject (first line): Merge pull request #23132 from mlaventure/bump-to-1.11.2
[INFO]	Keeping github.com/docker/docker 0f5c9d301b9b1cca66b3ea0f9dec3b5317d3686d
[WARN]	Conflict: github.com/Sirupsen/logrus version is ^0.10.0, but also asked for 51fe59aca108dc5680109e7b2051cbdcfa5a253c
[INFO]	github.com/Sirupsen/logrus reference 51fe59aca108dc5680109e7b2051cbdcfa5a253c:
[INFO] - author: Simon Eskildsen <sirup@sirupsen.com>
[INFO] - commit date: Thu, 18 Dec 2014 10:52:14 -0500
[INFO] - subject (first line): Merge pull request #92 from rasky/no-extra-quoting
[INFO]	Keeping github.com/Sirupsen/logrus ^0.10.0
[INFO]	--> Fetching updates for github.com/goadesign/goa
[INFO]	--> Detected semantic version. Setting version for github.com/goadesign/goa to v1.2.0
[INFO]	--> Fetching google.golang.org/appengine
[INFO]	--> Setting version for google.golang.org/appengine to 4f7eeb5305a4ba1966344836ba4af9996b7b4e05.
[INFO]	--> Fetching updates for github.com/asaskevich/govalidator
[INFO]	--> Detected semantic version. Setting version for github.com/asaskevich/govalidator to v5
[WARN]	Conflict: golang.org/x/crypto rev is currently c84e1f8e3a7e322d497cd16c0e8a13c7e127baf3, but github.com/kubernetes/minikube wants d172538b2cfce0c13cee31e647d0367aa8cd2486
[INFO]	golang.org/x/crypto reference c84e1f8e3a7e322d497cd16c0e8a13c7e127baf3:
[INFO] - author: Dmitry Savintsev <dsavints@gmail.com>
[INFO] - commit date: Wed, 22 Apr 2015 13:17:00 +0200
[INFO] - subject (first line): crypto/ssh: update references to the old code.google.com repo
[INFO]	golang.org/x/crypto reference d172538b2cfce0c13cee31e647d0367aa8cd2486:
[INFO] - author: Paul van Brouwershaven <paul@vanbrouwershaven.com>
[INFO] - commit date: Tue, 02 Jun 2015 14:54:44 +0000
[INFO] - subject (first line): x/crypto/ocsp: Accept status for multiple certificates in response
[INFO]	Keeping golang.org/x/crypto c84e1f8e3a7e322d497cd16c0e8a13c7e127baf3
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-k8s.io-client-go
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Fetching updates for github.com/howeyc/gopass
[INFO]	--> Setting version for github.com/howeyc/gopass to 3ca23474a7c7203e0a0a070fd33508f6efdb9b3d.
[INFO]	--> Fetching updates for k8s.io/apimachinery
[INFO]	--> Setting version for k8s.io/apimachinery to 2de00c78cb6d6127fb51b9531c1b3def1cbcac8c.
[INFO]	--> Fetching github.com/kr/fs
[INFO]	--> Setting version for github.com/kr/fs to 2788f0dbd16903de03cb8186e5c7d97b69ad387b.
[INFO]	--> Fetching updates for github.com/docker/engine-api
[INFO]	--> Setting version for github.com/docker/engine-api to dea108d3aa0c67d7162a3fd8aa65f38a430019fd.
[INFO]	--> Fetching updates for github.com/beorn7/perks
[INFO]	--> Setting version for github.com/beorn7/perks to 3ac7bf7a47d159a033b107610db8a1b6575507a4.
[INFO]	--> Fetching updates for github.com/golang/protobuf
[INFO]	--> Setting version for github.com/golang/protobuf to 8616e8ee5e20a1704615e6c8d7afcdac06087a67.
[INFO]	--> Fetching updates for github.com/prometheus/client_model
[INFO]	--> Setting version for github.com/prometheus/client_model to fa8ad6fec33561be4280a8f0514318c79d7f6cb6.
[INFO]	--> Fetching updates for github.com/prometheus/common
[INFO]	--> Setting version for github.com/prometheus/common to ffe929a3f4c4faeaa10f2b9535c2b1be3ad15650.
[INFO]	--> Fetching updates for github.com/prometheus/procfs
[INFO]	--> Setting version for github.com/prometheus/procfs to 454a56f35412459b5e684fd5ec0f9211b94f002a.
[INFO]	--> Fetching updates for github.com/coreos/go-oidc
[INFO]	--> Setting version for github.com/coreos/go-oidc to be73733bb8cc830d0205609b95d125215f8e9c70.
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-github.com-opencontainers-runc
[INFO]	--> Parsing Godeps metadata...
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-github.com-docker-distribution
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Fetching updates for github.com/docker/libtrust
[INFO]	--> Setting version for github.com/docker/libtrust to c54fbb67c1f1e68d7d6f8d2ad7c9360404616a41.
[INFO]	--> Fetching updates for github.com/gonum/matrix
[INFO]	--> Setting version for github.com/gonum/matrix to fb1396264e2e259ff714a408a7b0142d238b198d.
[INFO]	--> Fetching updates for github.com/MakeNowJust/heredoc
[INFO]	--> Setting version for github.com/MakeNowJust/heredoc to 1d91351acdc1cb2f2c995864674b754134b86ca7.
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-k8s.io-apimachinery
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Fetching updates for github.com/docker/go-connections
[INFO]	--> Setting version for github.com/docker/go-connections to f549a9393d05688dff0992ef3efd8bbe6c628aeb.
[INFO]	--> Fetching updates for github.com/docker/go-units
[INFO]	--> Setting version for github.com/docker/go-units to e30f1e79f3cd72542f2026ceec18d3bd67ab859c.
[INFO]	--> Fetching bitbucket.org/ww/goautoneg
[INFO]	--> Setting version for bitbucket.org/ww/goautoneg to 75cd24fc2f2c2a2088577d12123ddee5f54e0675.
[INFO]	--> Fetching updates for github.com/matttproud/golang_protobuf_extensions
[INFO]	--> Setting version for github.com/matttproud/golang_protobuf_extensions to fc2b8d3a73c4867e51861bbdd5ae3c1f0869dd6a.
[INFO]	--> Fetching cloud.google.com/go
[INFO]	--> Setting version for cloud.google.com/go to 3b1ae45394a234c385be014e9a488f2bb6eef821.
[INFO]	--> Setting version for github.com/coreos/pkg to fa29b1d70f0beaddd4c7021607cc3c3be8ce94b8.
[INFO]	--> Fetching updates for github.com/jonboulle/clockwork
[INFO]	--> Setting version for github.com/jonboulle/clockwork to 72f9bd7c4e0c2a40055ab3d0f09654f730cce982.
[INFO]	--> Fetching updates for github.com/gorilla/mux
[INFO]	--> Setting version for github.com/gorilla/mux to 8096f47503459bcc74d1f4c487b7e6e42e5746b5.
[INFO]	--> Fetching updates for github.com/gonum/blas
[INFO]	--> Setting version for github.com/gonum/blas to 80dca99229cccca259b550ae3f755cf79c65a224.
[INFO]	--> Fetching updates for github.com/gonum/internal
[INFO]	--> Setting version for github.com/gonum/internal to 5b84ddfb9d3e72d73b8de858c97650be140935c0.
[INFO]	--> Fetching updates for github.com/gonum/lapack
[INFO]	--> Setting version for github.com/gonum/lapack to 88ec467285859a6cd23900147d250a8af1f38b10.
[INFO]	--> Fetching updates for github.com/armon/go-metrics
[INFO]	--> Setting version for github.com/armon/go-metrics to 345426c77237ece5dab0e1605c3e4b35c3f54757.
[INFO]	--> Fetching updates for github.com/dimfeld/httptreemux
[INFO]	--> Detected semantic version. Setting version for github.com/dimfeld/httptreemux to v3.8.0
[INFO]	--> Fetching updates for github.com/go-openapi/spec
[INFO]	--> Setting version for github.com/go-openapi/spec to 6aced65f8501fe1217321abf0749d354824ba2ff.
[INFO]	--> Fetching updates for github.com/gorilla/context
[INFO]	--> Setting version for github.com/gorilla/context to 215affda49addc4c8ef7e2534915df2c8c35c6cd.
[INFO]	--> Fetching updates for github.com/satori/go.uuid
[INFO]	--> Fetching updates for github.com/go-openapi/jsonpointer
[INFO]	--> Setting version for github.com/go-openapi/jsonpointer to 46af16f9f7b149af66e5d1bd010e3574dc06de98.
[INFO]	--> Fetching updates for github.com/go-openapi/jsonreference
[INFO]	--> Setting version for github.com/go-openapi/jsonreference to 13c6e3589ad90f49bd3e3bbe2c2cb3d7a4142272.
[INFO]	--> Fetching updates for github.com/go-openapi/swag
[INFO]	--> Setting version for github.com/go-openapi/swag to 1d0bd113de87027671077d3c71eb3ac5d7dbba72.
[INFO]	--> Fetching updates for github.com/PuerkitoBio/purell
[INFO]	--> Setting version for github.com/PuerkitoBio/purell to 8a290539e2e8629dbc4e6bad948158f790ec31f4.
[INFO]	--> Fetching updates for github.com/mailru/easyjson
[INFO]	--> Setting version for github.com/mailru/easyjson to d5b7844b561a7bc640052f1b935f7b800330d7e0.
[INFO]	--> Fetching updates for github.com/PuerkitoBio/urlesc
[INFO]	--> Setting version for github.com/PuerkitoBio/urlesc to 5bd2802263f21d8788851d5305584c82a5c75d7e.
[INFO]	Found Godeps.json file in /Users/fanhongling/.glide/cache/src/https-github.com-stretchr-testify
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Setting version for github.com/pmezard/go-difflib to d8ed2627bdf02c080bf22230dbb337003b7aba2d.
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for github.com/ghodss/yaml
[INFO]	--> Fetching updates for github.com/pkg/errors
[INFO]	--> Fetching updates for github.com/kubernetes/minikube
[INFO]	Setting references for remaining imports
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting github.com/daviddengcn/go-colortext
[INFO]	--> Exporting github.com/spf13/cobra
[INFO]	--> Exporting github.com/spf13/pflag
[INFO]	--> Exporting github.com/pkg/browser
[INFO]	--> Exporting github.com/stretchr/testify
[INFO]	--> Exporting github.com/golang/glog
[INFO]	--> Exporting k8s.io/kubernetes
[INFO]	--> Exporting github.com/coreos/pkg
[INFO]	--> Exporting github.com/fabric8io/fabric8-init-tenant
[INFO]	--> Exporting speter.net/go/exp/math/dec/inf
[INFO]	--> Exporting github.com/jimmidyson/minishift
[INFO]	--> Exporting github.com/coreos/go-systemd
[INFO]	--> Exporting github.com/imdario/mergo
[INFO]	--> Exporting golang.org/x/crypto
[INFO]	--> Exporting github.com/gogo/protobuf
[INFO]	--> Exporting github.com/pmezard/go-difflib
[INFO]	--> Exporting github.com/gonum/graph
[INFO]	--> Exporting github.com/Sirupsen/logrus
[INFO]	--> Exporting github.com/docker/docker
[INFO]	--> Exporting github.com/openshift/origin
[INFO]	--> Exporting github.com/dgrijalva/jwt-go
[INFO]	--> Exporting github.com/jteeuwen/go-bindata
[INFO]	--> Exporting github.com/elazarl/go-bindata-assetfs
[INFO]	--> Exporting github.com/blang/semver
[INFO]	--> Exporting github.com/google/go-github
[INFO]	--> Exporting github.com/inconshreveable/go-update
[INFO]	--> Exporting github.com/kardianos/osext
[INFO]	--> Exporting github.com/minishift/minishift
[INFO]	--> Exporting github.com/spf13/viper
[INFO]	--> Exporting github.com/inconshreveable/mousetrap
[INFO]	--> Exporting github.com/hashicorp/hcl
[INFO]	--> Exporting github.com/fsnotify/fsnotify
[INFO]	--> Exporting github.com/magiconair/properties
[INFO]	--> Exporting github.com/mitchellh/mapstructure
[INFO]	--> Exporting github.com/pelletier/go-toml
[INFO]	--> Exporting github.com/spf13/afero
[INFO]	--> Exporting github.com/spf13/cast
[INFO]	--> Exporting github.com/spf13/jwalterweatherman
[INFO]	--> Exporting github.com/emicklei/go-restful
[INFO]	--> Exporting github.com/evanphx/json-patch
[INFO]	--> Exporting github.com/google/cadvisor
[INFO]	--> Exporting github.com/davecgh/go-spew
[INFO]	--> Exporting k8s.io/minikube
[INFO]	--> Exporting github.com/google/gofuzz
[INFO]	--> Exporting github.com/almighty/almighty-core
[INFO]	--> Exporting github.com/google/go-querystring
[INFO]	--> Exporting golang.org/x/net
[INFO]	--> Exporting github.com/mattn/go-runewidth
[INFO]	--> Exporting github.com/pelletier/go-buffruneio
[INFO]	--> Exporting golang.org/x/oauth2
[INFO]	--> Exporting github.com/pkg/sftp
[INFO]	--> Exporting github.com/golang/groupcache
[INFO]	--> Exporting golang.org/x/sys
[INFO]	--> Exporting github.com/ugorji/go
[INFO]	--> Exporting github.com/prometheus/client_golang
[INFO]	--> Exporting github.com/juju/ratelimit
[INFO]	--> Exporting golang.org/x/text
[INFO]	--> Exporting gopkg.in/yaml.v2
[INFO]	--> Exporting gopkg.in/cheggaaa/pb.v1
[INFO]	--> Exporting github.com/opencontainers/runc
[INFO]	--> Exporting gopkg.in/inf.v0
[INFO]	--> Exporting k8s.io/client-go
[INFO]	--> Exporting github.com/pborman/uuid
[INFO]	--> Exporting github.com/docker/distribution
[INFO]	--> Exporting github.com/fsouza/go-dockerclient
[INFO]	--> Exporting github.com/goadesign/goa
[INFO]	--> Exporting github.com/asaskevich/govalidator
[INFO]	--> Exporting github.com/howeyc/gopass
[INFO]	--> Exporting github.com/kr/fs
[INFO]	--> Exporting google.golang.org/appengine
[INFO]	--> Exporting github.com/docker/engine-api
[INFO]	--> Exporting github.com/beorn7/perks
[INFO]	--> Exporting github.com/golang/protobuf
[INFO]	--> Exporting github.com/prometheus/client_model
[INFO]	--> Exporting github.com/prometheus/common
[INFO]	--> Exporting github.com/prometheus/procfs
[INFO]	--> Exporting github.com/coreos/go-oidc
[INFO]	--> Exporting k8s.io/apimachinery
[INFO]	--> Exporting github.com/docker/libtrust
[INFO]	--> Exporting github.com/MakeNowJust/heredoc
[INFO]	--> Exporting github.com/gonum/matrix
[INFO]	--> Exporting github.com/docker/go-connections
[INFO]	--> Exporting github.com/docker/go-units
[INFO]	--> Exporting github.com/matttproud/golang_protobuf_extensions
[INFO]	--> Exporting github.com/jonboulle/clockwork
[INFO]	--> Exporting github.com/gorilla/mux
[INFO]	--> Exporting github.com/gonum/internal
[INFO]	--> Exporting github.com/gonum/blas
[INFO]	--> Exporting github.com/gonum/lapack
[INFO]	--> Exporting github.com/armon/go-metrics
[INFO]	--> Exporting github.com/dimfeld/httptreemux
[INFO]	--> Exporting github.com/go-openapi/spec
[INFO]	--> Exporting github.com/gorilla/context
[INFO]	--> Exporting github.com/satori/go.uuid
[INFO]	--> Exporting github.com/go-openapi/jsonpointer
[INFO]	--> Exporting github.com/go-openapi/jsonreference
[INFO]	--> Exporting github.com/go-openapi/swag
[INFO]	--> Exporting github.com/PuerkitoBio/purell
[INFO]	--> Exporting github.com/mailru/easyjson
[INFO]	--> Exporting github.com/PuerkitoBio/urlesc
[INFO]	--> Exporting github.com/kubernetes/minikube
[INFO]	--> Exporting cloud.google.com/go
[INFO]	--> Exporting github.com/ghodss/yaml
[INFO]	--> Exporting github.com/pkg/errors
[INFO]	--> Exporting bitbucket.org/ww/goautoneg
[INFO]	Replacing existing vendor dependencies
[INFO]	Project relies on 105 dependencies.
[INFO]	Removing nested vendor and Godeps/_workspace directories...
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/docker/distribution/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/docker/docker/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/google/cadvisor/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/kubernetes/minikube/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/opencontainers/runc/Godeps/_workspace/src/github.com/docker/docker/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/openshift/origin/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/stretchr/testify/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/k8s.io/apimachinery/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/k8s.io/client-go/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/k8s.io/kubernetes/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/k8s.io/minikube/vendor
[INFO]	Removing: /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/opencontainers/runc/Godeps/_workspace
[INFO]	Removing Godep rewrites for /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/opencontainers/runc
cd vendor/github.com/jteeuwen/go-bindata/go-bindata && go build -v
github.com/fabric8io/gofabric8/vendor/github.com/jteeuwen/go-bindata
github.com/fabric8io/gofabric8/vendor/github.com/jteeuwen/go-bindata/go-bindata
export TEAM_VERSION=1.0.158
echo "using team version 1.0.158"
using team version 1.0.158
cd vendor/github.com/fabric8io/fabric8-init-tenant && TEAM_VERSION=1.0.158 go generate template/generate.go
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 11668  100 11668    0     0   5905      0  0:00:01  0:00:01 --:--:--  5904
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 28367  100 28367    0     0  40629      0 --:--:-- --:--:-- --:--:-- 40582
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 16132  100 16132    0     0  25485      0 --:--:-- --:--:-- --:--:-- 25525
cd vendor/github.com/fabric8io/fabric8-init-tenant && /Users/fanhongling/Downloads/workspace-openshift/src/github.com/fabric8io/gofabric8/vendor/github.com/jteeuwen/go-bindata/go-bindata/go-bindata -o template/bindata.go \
	-pkg template \
	-prefix '' \
	-nocompress \
	template
```

And
```
fanhonglingdeMacBook-Pro:gofabric8 fanhongling$ make templates
make: Nothing to be done for `templates'.
```

From __Trusty (Ubuntu 14.04)__
```
vagrant@vagrant-ubuntu-trusty-64:/data/src/github.com/fabric8io/gofabric8$ make build
rm -rf build
CGO_ENABLED=0 GO15VENDOREXPERIMENT=1 go build -ldflags " -X github.com/fabric8io/gofabric8/version.Version=0.4.124 -X github.com/fabric8io/gofabric8/version.Revision='d215325' -X github.com/fabric8io/gofabric8/version.Branch='master' -X github.com/fabric8io/gofabric8/version.BuildDate='20170524-03:44:33' -X github.com/fabric8io/gofabric8/version.GoVersion='1.8.1'" -o build/gofabric8 gofabric8.go
# github.com/fabric8io/gofabric8/vendor/github.com/goadesign/goa/middleware/security/jwt
vendor/github.com/goadesign/goa/middleware/security/jwt/jwt.go:136: invalid type assertion: token.Claims.(jwt.MapClaims) (non-interface type map[string]interface {} on left)
vendor/github.com/goadesign/goa/middleware/security/jwt/jwt.go:141: use of .(type) outside type switch
# github.com/fabric8io/gofabric8/vendor/k8s.io/kubernetes/pkg/util
vendor/k8s.io/kubernetes/pkg/util/resource_container_linux.go:39: cannot use true (type bool) as type *bool in field value
make: *** [build] Error 2
```

## Version 0.4.71

Build bin
```
[vagrant@localhost gofabric8]$ make build
CGO_ENABLED=0 GO15VENDOREXPERIMENT=1 go build -ldflags " -X github.com/fabric8io/gofabric8/version.Version=0.4.71 -X github.com/fabric8io/gofabric8/version.Revision='ea5306f' -X github.com/fabric8io/gofabric8/version.Branch='master' -X github.com/fabric8io/gofabric8/version.BuildDate='20160921-14:54:24' -X github.com/fabric8io/gofabric8/version.GoVersion='1.6.2'" -o build/gofabric8 gofabric8.go
```

### Deploy

Into Kubernetes
```
[vagrant@localhost gofabric8]$ ./build/gofabric8 deploy --domain=172.17.4.50.xip.io --ingress=false --loadbalancer=false
             ▄▄▄▄▄▄▄         
             ███████         
     ▄▄▄▄▄▄▄ ▄▄▄▄▄▄█         
     ███████ ▀▀▀▀▀▀▀ ▄▄▄▄▄▄  
     ▄▄▄▄▄▄▄         ██████  
   ▄▄▄▄▄             ▄▄▄▄▄▄  
   ▀▄▄▄▄▄      ▄▄▄       ▄▄▄ 
    ▀▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄   ▄▄▄▄▄▀
      ▀▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█▄▄▀ 
        ▀▀▄█▄▄▄▄▄▄▄▄▄▄▄▄▄▀   
          ▄█▄▄▄█▄▄▄█▄▄▄      
          ▄█▄▄▄▄▄▄▄▄▄▄▄      
           ▄▄█▄▄▄▄▄▄▄▄       
            ▄▄▄▄▀▀▄█▄        
         ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄     
         ▀▀▀▀▀▀▀▀▀▀▀▀▀▀      
Deploying fabric8 to your Kubernetes installation at https://172.17.4.50:443 for domain 172.17.4.50.xip.io in namespace default

Loading fabric8 releases from maven repository:https://repo1.maven.org/maven2/
Continue? [Y/n] y

Starting fabric8 console deployment using 2.2.182...

fabric8 console...............................................................✔ 
addServiceAccount fluentd.....................................................✔ 
addServiceAccount registry....................................................✔ 
Created fabric8 console
Installing templates!
Downloading apps from: https://repo1.maven.org/maven2/io/fabric8/forge/distro/distro/2.3.48/distro-2.3.48-templates.zip
Loading template kubernetes/main/chat-letschat.yml
Loading template kubernetes/microservices/hubot-notifier.yml
Loading template kubernetes/microservices/gitlab.yml
Loading template kubernetes/main/chat-slack.yml
Loading template kubernetes/microservices/hubot-irc.yml
Loading template kubernetes/microservices/git-collector.yml
Loading template kubernetes/microservices/prometheus-blackbox-exporter.yml
Loading template kubernetes/microservices/prometheus-node-exporter.yml
Loading template kubernetes/microservices/grafana.yml
Loading template kubernetes/microservices/artifactory.yml
Loading template kubernetes/microservices/kiwiirc.yml
Loading template kubernetes/microservices/gogs.yml
Loading template kubernetes/microservices/letschat.yml
Loading template kubernetes/microservices/nexus.yml
Loading template kubernetes/microservices/fabric8-forge.yml
Loading template kubernetes/main/social.yml
Loading template kubernetes/microservices/hubot-letschat.yml
Loading template kubernetes/microservices/fluentd.yml
Loading template kubernetes/main/management.yml
Loading template kubernetes/microservices/hubot-slack.yml
Loading template kubernetes/microservices/kibana.yml
Loading template kubernetes/main/cd-pipeline.yml
Loading template kubernetes/microservices/ingress-nginx.yml
Loading template kubernetes/main/chat-irc.yml
Loading template kubernetes/microservices/exposecontroller.yml
Loading template kubernetes/microservices/gerrit.yml
Loading template kubernetes/main/logging.yml
Loading template kubernetes/microservices/fabric8-docker-registry.yml
Loading template kubernetes/microservices/prometheus.yml
Loading template kubernetes/microservices/chaos-monkey.yml
Loading template kubernetes/microservices/jenkins.yml
Loading template kubernetes/microservices/maven-shell.yml
Loading template kubernetes/microservices/manageiq.yml
Loading template kubernetes/microservices/content-repository.yml
Loading template kubernetes/main/metrics.yml
Loading template kubernetes/microservices/elasticsearch.yml
Install DevOps templates......................................................✔ 
Downloading apps from: https://repo1.maven.org/maven2/io/fabric8/ipaas/distro/distro/2.2.154/distro-2.2.154-templates.zip
Loading template kubernetes/microservices/example-message-consumer.yml
Loading template kubernetes/main/messaging.yml
Loading template kubernetes/microservices/example-message-producer.yml
Loading template kubernetes/microservices/message-broker.yml
Loading template kubernetes/microservices/cassandra.yml
Loading template kubernetes/microservices/message-gateway.yml
Loading template kubernetes/microservices/zookeeper.yml
Loading template kubernetes/microservices/kafka.yml
Install iPaaS templates.......................................................✔ 
Downloading apps from: https://repo1.maven.org/maven2/io/fabric8/kubeflix/distro/distro/1.0.22/distro-1.0.22-templates.zip
Loading template kubernetes/microservices/hystrix-dashboard.yml
Loading template kubernetes/main/kubeflix.yml
Loading template kubernetes/microservices/turbine-server.yml
Install Kubeflix templates....................................................✔ 
Downloading apps from: https://repo1.maven.org/maven2/io/fabric8/zipkin/packages/distro/0.1.5/distro-0.1.5-templates.zip
Loading template main/zipkin-0.1.5.json
Install Zipkin templates......................................................✔ 


Installing: exposecontroller

Processing resource kind: ServiceAccount in namespace default
Processing resource kind: Deployment in namespace default
exposecontroller..............................................................✔ 


Installing: cd-pipeline

Found namespace on kind Secret of user-secrets-source-adminCreating new Namespace: user-secrets-source-admin
Processing resource kind: Secret in namespace user-secrets-source-admin
Processing resource kind: Secret in namespace default
Processing resource kind: Secret in namespace default
Processing resource kind: Secret in namespace default
Processing resource kind: Secret in namespace default
Processing resource kind: Secret in namespace default
Processing resource kind: Secret in namespace default
Processing resource kind: Secret in namespace default
Processing resource kind: ServiceAccount in namespace default
Processing resource kind: ServiceAccount in namespace default
Processing resource kind: Service in namespace default
Processing resource kind: Service in namespace default
Processing resource kind: Service in namespace default
Processing resource kind: Service in namespace default
Processing resource kind: Service in namespace default
Processing resource kind: Service in namespace default
Processing resource kind: Service in namespace default
Processing resource kind: ConfigMap in namespace default
Processing resource kind: ConfigMap in namespace default
Processing resource kind: ConfigMap in namespace default
Processing resource kind: Deployment in namespace default
Processing resource kind: Deployment in namespace default
Converted Deployment to avoid the use of PersistentVolumeClaim
Processing resource kind: Deployment in namespace default
Converted Deployment to avoid the use of PersistentVolumeClaim
Processing resource kind: Deployment in namespace default
Converted Deployment to avoid the use of PersistentVolumeClaim
Processing resource kind: Deployment in namespace default
cd-pipeline...................................................................✔ 

-------------------------

Default GOGS admin username/password = gogsadmin/RedHat$1

Downloading images and waiting to open the fabric8 console...

-------------------------
..
Opening URL http://fabric8.default.172.17.4.50.xip.io
/bin/xdg-open:行781: www-browser: 未找到命令
/bin/xdg-open:行781: links2: 未找到命令
/bin/xdg-open:行781: elinks: 未找到命令
/bin/xdg-open:行781: links: 未找到命令
/bin/xdg-open:行781: lynx: 未找到命令
/bin/xdg-open:行781: w3m: 未找到命令
xdg-open: no method available for opening 'http://fabric8.default.172.17.4.50.xip.io'


Installing: ingress-nginx

Found namespace on kind ConfigMap of fabric8-systemProcessing resource kind: ConfigMap in namespace fabric8-system
Failed to create ConfigMap: configmaps "nginx-config" already existsingress-nginx.................................................................✘ configmaps "nginx-config" already exists
Found namespace on kind Deployment of fabric8-systemProcessing resource kind: Deployment in namespace fabric8-system
ingress-nginx.................................................................✔ 

Deploying ingress controller on node 172.17.4.50 use its external ip when configuring your wildcard DNS.
To change node move the label: `kubectl label node 172.17.4.50 externalIP- && kubectl label node $YOUR_NEW_NODE externalIP=true`

---
      apiVersion: "v1"
      kind: "ConfigMap"
      metadata:
        labels:
          provider: "fabric8"
          project: "ingress-nginx"
          version: "2.2.265"
          group: "io.fabric8.devops.apps"
        name: "nginx-config"
        # namespace: "fabric8-system"
      data:
        client-max-body-size: "2000m"
        proxy-connect-timeout: "10s"
        proxy-read-timeout: "10s"
        server-names-hash-bucket-size: "256"
        server-names-hash-max-size: "1024"
---
      apiVersion: "extensions/v1beta1"
      kind: "Deployment"
      metadata:
        annotations:
          fabric8.io/git-commit: "3ed46620b6802ae12c469284f09cf8b799d88411"
          fabric8.io/metrics-path: "dashboard/file/kubernetes-pods.json/?var-project=ingress-nginx&var-version=2.2.265"
          fabric8.io/build-id: "3"
          fabric8.io/build-url: "http://jenkins.ux.fabric8.io/job/oss-parent/3"
          fabric8.io/iconUrl: "https://cdn.rawgit.com/fabric8io/fabric8-devops/master/ingress-nginx/src/main/fabric8/icon.png"
          fabric8.io/git-branch: "release-v2.2.265"
          fabric8.io/git-url: "http://gogs.ux.fabric8.io/gogsadmin/oss-parent/commit/3ed46620b6802ae12c469284f09cf8b799d88411"
        labels:
          provider: "fabric8"
          project: "ingress-nginx"
          version: "2.2.265"
          group: "io.fabric8.devops.apps"
        name: "ingress-nginx"
        # namespace: "fabric8-system"
      spec:
        replicas: 1
        selector:
          matchLabels:
            project: "ingress-nginx"
            provider: "fabric8"
            group: "io.fabric8.devops.apps"
        template:
          metadata:
            annotations:
              fabric8.io/git-commit: "3ed46620b6802ae12c469284f09cf8b799d88411"
              fabric8.io/metrics-path: "dashboard/file/kubernetes-pods.json/?var-project=ingress-nginx&var-version=2.2.265"
              fabric8.io/build-id: "3"
              fabric8.io/build-url: "http://jenkins.ux.fabric8.io/job/oss-parent/3"
              fabric8.io/iconUrl: "https://cdn.rawgit.com/fabric8io/fabric8-devops/master/ingress-nginx/src/main/fabric8/icon.png"
              fabric8.io/git-branch: "release-v2.2.265"
              fabric8.io/git-url: "http://gogs.ux.fabric8.io/gogsadmin/oss-parent/commit/3ed46620b6802ae12c469284f09cf8b799d88411"
            labels:
              provider: "fabric8"
              project: "ingress-nginx"
              version: "2.2.265"
              group: "io.fabric8.devops.apps"
          spec:
            containers:
            - args:
              - "-v=3"
              - "-nginx-configmaps=fabric8-system/nginx-config"
              image: "nginxdemos/nginx-ingress:0.3"
              name: "nginx-ingress"
              ports:
              - containerPort: 80
                hostPort: 80
                name: "http"
                protocol: "TCP"
              # - containerPort: 443
              #   hostPort: 443
              #   name: "https"
              #   protocol: "TCP"
            nodeSelector:
              externalIP: "true"
```

### Set up DNS

Using __skydns__ and __etcd__ running up DNS
```
[vagrant@localhost skydns]$ curl -XPUT http://172.17.4.50:30001/v2/keys/skydns/f8/vagrant/dns/ns/ns1 -d value='{"host":"172.17.4.50"}'
{"action":"set","node":{"key":"/skydns/f8/vagrant/dns/ns/ns1","value":"{\"host\":\"172.17.4.50\"}","modifiedIndex":10,"createdIndex":10}}

[vagrant@localhost skydns]$ dig @localhost SRV ns1.ns.dns.vagrant.f8

; <<>> DiG 9.10.3-P3-RedHat-9.10.3-10.P3.fc23 <<>> @localhost SRV ns1.ns.dns.vagrant.f8
; (2 servers found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 35360
;; flags: qr aa rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; QUESTION SECTION:
;ns1.ns.dns.vagrant.f8.		IN	SRV

;; ANSWER SECTION:
ns1.ns.dns.vagrant.f8.	3600	IN	SRV	10 100 0 ns1.ns.dns.vagrant.f8.

;; ADDITIONAL SECTION:
ns1.ns.dns.vagrant.f8.	3600	IN	A	172.17.4.50

;; Query time: 4 msec
;; SERVER: ::1#53(::1)
;; WHEN: 六 10月 01 10:42:36 UTC 2016
;; MSG SIZE  rcvd: 96

[vagrant@localhost skydns]$ curl -XPUT http://172.17.4.50:30001/v2/keys/skydns/f8/vagrant/default/fabric8 -d value='{"host":"172.17.4.50","port":80}'
{"action":"set","node":{"key":"/skydns/f8/vagrant/default/fabric8","value":"{\"host\":\"172.17.4.50\",\"port\":80}","modifiedIndex":12,"createdIndex":12}}

[vagrant@localhost skydns]$ curl -XPUT http://172.17.4.50:30001/v2/keys/skydns/f8/vagrant/default/gogs -d value='{"host":"172.17.4.50","port":80}'
{"action":"set","node":{"key":"/skydns/f8/vagrant/default/gogs","value":"{\"host\":\"172.17.4.50\",\"port\":80}","modifiedIndex":13,"createdIndex":13}}

[vagrant@localhost skydns]$ curl -XPUT http://172.17.4.50:30001/v2/keys/skydns/f8/vagrant/default/nexus -d value='{"host":"172.17.4.50","port":80}'
{"action":"set","node":{"key":"/skydns/f8/vagrant/default/nexus","value":"{\"host\":\"172.17.4.50\",\"port\":80}","modifiedIndex":14,"createdIndex":14}}

[vagrant@localhost skydns]$ curl -XPUT http://172.17.4.50:30001/v2/keys/skydns/f8/vagrant/default/jenkins -d value='{"host":"172.17.4.50","port":80}'
{"action":"set","node":{"key":"/skydns/f8/vagrant/default/jenkins","value":"{\"host\":\"172.17.4.50\",\"port\":80}","modifiedIndex":15,"createdIndex":15}}

[vagrant@localhost skydns]$ dig @localhost A local.dns.vagrant.f8

; <<>> DiG 9.10.3-P3-RedHat-9.10.3-10.P3.fc23 <<>> @localhost A local.dns.vagrant.f8
; (2 servers found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 10193
;; flags: qr aa rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;local.dns.vagrant.f8.		IN	A

;; ANSWER SECTION:
local.dns.vagrant.f8.	3600	IN	A	172.17.4.50

;; Query time: 16 msec
;; SERVER: ::1#53(::1)
;; WHEN: 六 10月 01 10:54:49 UTC 2016
;; MSG SIZE  rcvd: 54

[vagrant@localhost skydns]$ dig @localhost SRV default.vagrant.f8

; <<>> DiG 9.10.3-P3-RedHat-9.10.3-10.P3.fc23 <<>> @localhost SRV default.vagrant.f8
; (2 servers found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 22141
;; flags: qr aa rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; QUESTION SECTION:
;default.vagrant.f8.		IN	SRV

;; ANSWER SECTION:
default.vagrant.f8.	3600	IN	SRV	10 100 80 gogs.default.vagrant.f8.

;; ADDITIONAL SECTION:
gogs.default.vagrant.f8. 3600	IN	A	172.17.4.50

;; Query time: 9 msec
;; SERVER: ::1#53(::1)
;; WHEN: 六 10月 01 11:21:42 UTC 2016
;; MSG SIZE  rcvd: 100


[vagrant@localhost skydns]$ cat /etc/resolv.conf
nameserver 127.0.0.1
# Generated by NetworkManager
nameserver 10.0.2.3
```