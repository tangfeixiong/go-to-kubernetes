# Learning Kubernetes API v1 

## api versions

>`# /opt/tfx/kubectl api-versions`

>Available Server Api Versions: v1beta3,v1

>`# curl http://127.0.0.1:8080/api; echo`

>{

>  "versions": [

>    "v1beta3",

>    "v1"

>  ]

>}

## get PODs

>`# /opt/tfx/kubectl get pods`

>NAME                   READY     STATUS    RESTARTS   AGE

>k8s-master-127.0.0.1   3/3       Running   6          10d

>nginx-syeg7            1/1       Running   1          10d

The curl command to api is same as `kuberctl get -o json pods`

>`# curl -H "Accept: application/json" http://127.0.0.1:8080/api/v1/pods`

>{

>  "kind": "PodList",

>  "apiVersion": "v1",

>  "metadata": {

>    "selfLink": "/api/v1/pods",

>    "resourceVersion": "56963"

>  },

>  "items": [

>    {

>      "metadata": {

>        "name": "k8s-master-127.0.0.1",

>        "namespace": "default",

>        "selfLink": "/api/v1/namespaces/default/pods/k8s-master-127.0.0.1",

>        "uid": "9d1d876f-41a1-11e5-b8c4-c4346b46de6e",

>        "resourceVersion": "56273",

>        "creationTimestamp": "2015-08-13T09:56:48Z",

>        "annotations": {

>          "kubernetes.io/config.mirror": "mirror",

>          "kubernetes.io/config.seen": "2015-08-13T09:56:04.922025669Z",

>          "kubernetes.io/config.source": "file"

>        }

>      },

>      "spec": {

>        "containers": [

>          {

>            "name": "controller-manager",

>            "image": "gcr.io/google_containers/hyperkube:v0.21.2",

>            "command": [

>              "/hyperkube",

>              "controller-manager",

>              "--master=127.0.0.1:8080",

>              "--v=2"

>            ],

>            "resources": {},

>            "terminationMessagePath": "/dev/termination-log",

>            "imagePullPolicy": "IfNotPresent"

>          },

>          {

>            "name": "apiserver",

>            "image": "gcr.io/google_containers/hyperkube:v0.21.2",

>            "command": [

>              "/hyperkube",

>              "apiserver",
 
>             "--portal-net=10.0.0.1/24",

>             "--address=127.0.0.1",
 
>             "--etcd_servers=http://127.0.0.1:4001",

>              "--cluster_name=kubernetes",

>              "--v=2"

>            ],

>            "resources": {},

>            "terminationMessagePath": "/dev/termination-log",

>            "imagePullPolicy": "IfNotPresent"

>          },

>          {

>            "name": "scheduler",

>            "image": "gcr.io/google_containers/hyperkube:v0.21.2",

>            "command": [

>              "/hyperkube",

>              "scheduler",

>              "--master=127.0.0.1:8080",

>              "--v=2"

>            ],

>            "resources": {},

>            "terminationMessagePath": "/dev/termination-log",

>            "imagePullPolicy": "IfNotPresent"

>          }

>        ],

>        "restartPolicy": "Always",

>        "dnsPolicy": "ClusterFirst",

>        "nodeName": "127.0.0.1",

>        "hostNetwork": true

>      },

>      "status": {

>        "phase": "Running",

>        "conditions": [

>          {

>            "type": "Ready",

>            "status": "True"

>          }

>        ],

>        "hostIP": "127.0.0.1",

>        "podIP": "127.0.0.1",

>        "startTime": "2015-08-13T09:56:48Z",

>        "containerStatuses": [

>          {

>            "name": "apiserver",

>            "state": {

>              "running": {

>                "startedAt": "2015-08-24T07:27:44Z"

>              }

>            },

>            "lastState": {

>              "terminated": {

>                "exitCode": 2,

>                "startedAt": "2015-08-13T09:56:38Z",

>                "finishedAt": "2015-08-19T07:04:08Z",

>                "containerID": "docker://591e5f4e1f14623c93c48f15dec74d3e49034080b91fe0c866ac686b9130a8e1"

>              }

>            },

>            "ready": true,

>            "restartCount": 2,

>            "image": "gcr.io/google_containers/hyperkube:v0.21.2",

>            "imageID": "docker://f1f0bba2640f9143f75120c92f0686f1721b72588eecc7006b5c61b7727287ab",

>            "containerID": "docker://fa8bac1ebcfb188b0e3ae43c3da8759a8217dfd66ec62b0d99d264e854b616d9"

>          },

>          {

>            "name": "controller-manager",

>            "state": {
 
>            "running": {

>                "startedAt": "2015-08-24T07:27:46Z"

>              }

>            },

>            "lastState": {

>              "terminated": {

>                "exitCode": 2,

>                "startedAt": "2015-08-13T09:56:31Z",

>                "finishedAt": "2015-08-19T07:04:08Z",

>                "containerID": "docker://2f51426fec0b21e2333d9d4ed52c54ef205f67124773df0f3ea26e4c2c3e301c"

>              }

>            },

>            "ready": true,

>            "restartCount": 2,

>            "image": "gcr.io/google_containers/hyperkube:v0.21.2",

>            "imageID": "docker://f1f0bba2640f9143f75120c92f0686f1721b72588eecc7006b5c61b7727287ab",

>            "containerID": "docker://ffed942cfc52c9831038c92297a020010c8a3ed43ebdd76347ec1e3554128384"

>          },

>          {

>            "name": "scheduler",

>            "state": {

>              "running": {

>                "startedAt": "2015-08-24T07:27:45Z"

>              }

>            },

>            "lastState": {

>              "terminated": {

>                "exitCode": 2,

>                "startedAt": "2015-08-13T09:56:48Z",

>                "finishedAt": "2015-08-19T07:04:08Z",

>                "containerID": "docker://a6ff3b15a011a3e1c653467b800f8ea52071c699359a6f841f62d41a6d8f9b71"

>              }

>            },

>            "ready": true,

>            "restartCount": 2,

>            "image": "gcr.io/google_containers/hyperkube:v0.21.2",

>            "imageID": "docker://f1f0bba2640f9143f75120c92f0686f1721b72588eecc7006b5c61b7727287ab",

>            "containerID": "docker://0af559d7e35142a490bbe68000ef33ac4f02c68bf5e7079c4d9d2728cbc1c3a9"

>          }

>        ]

>      }

>    },

>    {

>      "metadata": {
 
>      "name": "nginx-syeg7",

>        "generateName": "nginx-",

>        "namespace": "default",

>        "selfLink": "/api/v1/namespaces/default/pods/nginx-syeg7",

>        "uid": "fc945b19-41a1-11e5-b8c4-c4346b46de6e",

>        "resourceVersion": "56291",

>        "creationTimestamp": "2015-08-13T09:59:28Z",

>        "labels": {

>          "run": "nginx"

>        },

>        "annotations": {

>          "kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"nginx\",\"uid\":\"fc8d142b-41a1-11e5-b8c4-c4346b46de6e\",\"apiVersion\":\"v1\",\"resourceVersion\":\"54\"}}"

>        }

>      },

>      "spec": {

>        "containers": [

>          {

>            "name": "nginx",

>            "image": "nginx",

>            "ports": [

>              {

>                "containerPort": 80,

>                "protocol": "TCP"

>              }

>            ],

>            "resources": {},

>            "terminationMessagePath": "/dev/termination-log",

>            "imagePullPolicy": "IfNotPresent"

>          }

>        ],

>        "restartPolicy": "Always",

>        "dnsPolicy": "ClusterFirst",

>        "nodeName": "127.0.0.1"

>      },

>      "status": {

>        "phase": "Running",

>        "conditions": [

>          {

>            "type": "Ready",

>            "status": "True"

>          }

>        ],

>        "hostIP": "127.0.0.1",

>        "podIP": "172.17.0.1",

>        "startTime": "2015-08-24T07:27:44Z",

>        "containerStatuses": [

>          {

>            "name": "nginx",

>            "state": {

>              "running": {

>                "startedAt": "2015-08-24T07:27:55Z"

>              }

>            },

>            "lastState": {

>              "terminated": {

>                "exitCode": 0,

>                "startedAt": "2015-08-13T09:59:39Z",

>                "finishedAt": "2015-08-19T07:04:08Z",

>                "containerID": "docker://7242304a8719b3f91a989052da04aa0e41dd032ef355d1941987da352fd4abdf"

>              }

>            },

>            "ready": true,

>            "restartCount": 1,

>            "image": "nginx",

>            "imageID": "docker://6886fb5a9b8d73b12d842bab8f9a6941c36094c2974abddb685d54c9d99e37da",

>            "containerID": "docker://58ca7b7bc89a1a25499f6c28c87ba533a73f7a27daa9b7216687461a5940ebbb"

>          }

>        ]

>      }

>    }

>  ]

>}


All POD reside in a Namespace, if not specfied, use `default` namespace 




## get Namespace

>`# /opt/tfx/kubectl get namespaces`

>NAME      LABELS    STATUS

>default   <none>    Active

>`# /opt/tfx/kubectl get -o yaml namespaces default`

>apiVersion: v1

>kind: Namespace

>metadata:

>  creationTimestamp: 2015-08-13T09:56:38Z

>  name: default

>  resourceVersion: "5"

>  selfLink: /api/v1/namespaces/default

>  uid: 96fbcab2-41a1-11e5-b8c4-c4346b46de6e

>spec:

>  finalizers:

>  - kubernetes

>status:

>  phase: Active

