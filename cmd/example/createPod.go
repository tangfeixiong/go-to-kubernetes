
package main

import (
    "fmt"
    "net/http"
    "errors"
    "bytes"
    "encoding/json"
    "io"
    "io/ioutil"
    "os"
    //"speter.net/go/exp/math/dec/inf"
)

type V1ObjectMeta struct {
    Name                string          `json:"name,omitempty"`
    GenerateName        string          `json:"generateName,omptempty"`
    Namespace           string          `json:"namespace,omitempty"`
    SelfLink            string          `json:"selfLink,omitempty"`
    Uid                 string          `json:"uid,omitempty"`
    ResourceVersion     string          `json:"resourceVersion,omitempty"`
    Generation          int64           `json:"generation,omitempty"`        
    CreationTimestamp   string          `json:"creationTimestamp,omitempty"`
    DeletionTimestamp   string          `json:"deletionTimestamp,omitempty"`
    Labels              map[string]string `json:"labels,omitempty"`
    Annotations         map[string]string `json:"annotations,omitempty"`
}

type V1HostPathVolumeSource struct {
    Path        string      `json:"path"`
}

type V1EmptyDirVolumeSource struct {
    Medium      string      `json:"medium,omitempty"`
}

type V1GCEPersistentDiskVolumeSource struct {
    PdName      string      `json:"pdName"`
    FsType      string      `json:"fsType"`
    Partition   int32       `json:"partition,omitempty"`
    ReadOnly    bool        `json:"readOnly,omitempty"`
}

type V1AWSElasticBlockStoreVolumeSource struct {
    VolumeID    string      `json:"volumeID"`
    FsType      string      `json:"fsType"`
    Partition   int32       `json:"partition,omitempty"`
    ReadOnly    bool        `json:"readOnly,omitempty"`
}

type V1GitRepoVolumeSource struct {
    Repository  string      `json:"repository"`
    Revision    string      `json:"revision"`
}

type V1SecretVolumeSource struct {
    SecretName  string      `json:"secretName"`
}

type V1NFSVolumeSource struct {
    Server      string      `json:"server"`
    Path        string      `json:"path"`
    ReadOnly    bool        `json:"readOnly,omitempty"`
}

type V1ISCSIVolumeSource struct {
    TargetPortal    string  `json:"targetProtal"`
    Iqn             string  `json:"iqn"`
    Lun             int32   `json:"lun"`
    FsType          string  `json:"fsType"`
    ReadOnly        bool    `json:"readOnly,omitempty"`
}    

type V1GlusterfsVolumeSource struct {
    Endpoints       string  `json:"endpoints"`  
    Path            string  `json:"path"`
    ReadOnly        bool    `json:"readOnly,omitempty"`
}

type V1PersistentVolumeClaimVolumeSource struct {
    ClaimName       string  `json:"claimName"`
    ReadOnly        bool    `json:"readOnly,omitempty"`
}

type V1RBDVolumeSource struct {
    Monitors        []string `json:"monitors"`
    Image           string   `json:"image"`
    FsType          string   `json:"fsType,omitempty"`
    Pool            string   `json:"pool"`
    User            string   `json:"user"`
    Keyring         string   `json:"keyring"`
    SecretRef       *V1LocalObjectReference `json:"secretRef"`
    ReadOnly        bool     `json:"readOnly,omitempty"`
}

type V1Volume struct {
    Name                string                  `json:"name"`
    HostPath            *V1HostPathVolumeSource  `json:"hostPath,omitempty"`
    EmptyDir            *V1EmptyDirVolumeSource  `json:"emptyDir,omitempty"`
    GcePersistentDisk *V1GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"`
    AwsElasticBlockStore *V1AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"`
    GitRepo             *V1GitRepoVolumeSource   `json:"gitRepo,omitempty"` 
    Secret              *V1SecretVolumeSource    `json:"secret,omitempty"`
    Nfs                 *V1NFSVolumeSource       `json:"nfs,omitempty"`
    Iscsi               *V1ISCSIVolumeSource     `json:"iscsi,omitempty"`
    Glusterfs           *V1GlusterfsVolumeSource `json:"glusterfs,omitempty"`
    PersistentVolumeClaim *V1PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`
    Rbd                 *V1RBDVolumeSource       `json:"rbd,omitempty"`
}

type V1ContainerPort struct {
    Name                string      `json:"name,omitempty"`     
    HostPort            int32       `json:"hostPort,omitempty"`
    ContainerPort       int32       `json:"containerPort"` 
    Protocol            string      `json:"protocol,omitempty"`
    HostIP              string      `json:"hostIP,omitempty"`
}

type V1ObjectFieldSelector struct {
    ApiVersion      string              `json:"apiVersion,omitempty"`
    FieldPath       string              `json:"fieldPath"`
}

type V1EnvVarSource struct {
    FieldRef        *V1ObjectFieldSelector   `json:"fieldRef"`
}

type V1EnvVar struct {
    Name            string              `json:"name"`
    Value           string              `json:"value,omitempty"`
    ValueFrom       *V1EnvVarSource      `json:"valueFrom,omitempty"`
}

type V1ResourceRequirements struct {
    Limits          map[string]string `json:"limits,omitempty"`
    Requests        map[string]string `json:"requests,omitempty"`
}

type V1VolumeMount struct {
    Name            string              `json:"name"`
    ReadOnly        bool                `json:"readOnly,omitempty"`
    MountPath       string              `json:"mountPath"`
}

type V1ExecAction struct {
    Command         []string    `json:"command,omitempty"`
}

type V1HTTPGetAction struct {
    Path            string      `json:"path,omitempty"`
    Port            string      `json:"port"`
    Host            string      `json:"host,omitempty"`
    Scheme          string      `json:"scheme,omitempty"`
}

type V1TCPSocketAction struct {
    Port            string      `json:"port"`
}

type V1Probe struct {
    Exec                V1ExecAction        `json:"exec,omitempty"`
    HttpGet             V1HTTPGetAction     `json:"httpGet,omitempty"`
    TcpSocket           V1TCPSocketAction   `json:"tcpSocket,omitempty"`
    InitialDelaySeconds int64           `json:"initialDelaySeconds,omitempty"`
    TimeoutSeconds      int64               `json:"timeoutSeconds,omitempty"`
}

type V1Handler struct {
    Exec            *V1ExecAction        `json:"exec,omitempty"`
    HttpGet         *V1HTTPGetAction     `json:"httpGet,omitempty"`
    TcpSocket       *V1TCPSocketAction   `json:"tcpSocket,omitempty"`
}

type V1Lifecycle struct {
    PostStart       *V1Handler       `json:"postStart,omitempty"`
    PreStop         *V1Handler       `json:"preStop,omitempty"`
}    

type V1Capability string

type V1Capabilities struct {
    Add         []V1Capability      `json:"add,omitempty"`
    Drop        []V1Capability      `json:"drop,omitempty"`
}

type V1SELinuxOptions struct {
    User            string      `json:"user,omitempty"`
    Role            string      `json:"role,omitempty"`
    Type            string      `json:"type,omitempty"`
    Level           string      `json:"level,omitempty"`
}

type V1SecurityContext struct {
    Capabilities    *V1Capabilities      `json:"capabilities,omitempty"`
    Privileged      *bool                `json:"privileged,omitempty"`
    SeLinuxOptions  *V1SELinuxOptions    `json:"seLinuxOptions,omitempty"`
    RunAsUser       *int64               `json:"runAsUser,omitempty"`
    RunAsNonRoot    bool `json:"runAsNonRoot,omitempty"`
}

type V1Container struct {
    Name                      string            `json:"name"`
    Image                     string            `json:"image,omitempty"`
    Command                   []string          `json:"command,omitempty"`
    Args                      []string          `json:"args,omitempty"`
    WorkingDir                string            `json:"workingDir,omitempty"`
    Ports                     []V1ContainerPort `json:"ports,omitempty"`
    Env                       []V1EnvVar        `json:"env,omitempty"`
    Resources           V1ResourceRequirements  `json:"resources,omitempty"`
    VolumeMounts        []V1VolumeMount       `json:"volumeMounts,omitempty"`
    LivenessProbe       *V1Probe             `json:"livenessProbe,omitempty"`
    ReadinessProbe      *V1Probe             `json:"readnessProbe,omitempty"`
    Lifecycle                 *V1Lifecycle   `json:"lifecycle,omitempty"`
    TerminationMessagePath    string  `json:"terminationMessagePath,omitempty"`
    ImagePullPolicy           string  `json:"imagePullPolicy,omitempty"`
    SecurityContext     *V1SecurityContext   `json:"securityContext,omitempty"`
    Stdin               bool                `json:"stdin,omitempty"`
    TTY                 bool                `json:"tty,omitempty"`
}

type V1LocalObjectReference  struct {
    Name        string      `json:"name,omitempty"`
}

type V1PodSpec struct {
    Volumns                 []V1Volume      `json:"volumes,omitempty"` 
    Containers              []V1Container   `json:"containers"` 
    RestartPolicy           string          `json:"restartPolicy,omitempty"`
    TerminationGracePeriodSeconds int64     `json:"terminationGracePeriodSeconds,omitempty"`
    ActiveDeadlineSeconds   int64   `json:"activeDeadlineSeconds,omitempty"`
    DnsPolicy               string      `json:"dnsPolicy,omitempty"`
    NodeSelector            map[string]string `json:"nodeSelector,omitemtpy"`
    ServiceAccountName      string      `json:"serviceAccountName,omitempty"`
    DeprecatedServiceAccount string     `json:"serviceAccount,omitempty"`
    ServiceAccount          string      `json:"serviceAccount,omitempty"`
    NodeName                string      `json:"nodeName,omitempty"`
    HostNetwork             bool        `json:"hostNetwork,omitempty"`
    ImagePullSecrets []V1LocalObjectReference `json:imagePullSecrets,omitemtpy"`
}

type V1PodCondition struct {
    Type        string      `json:"type"`
    Status      string      `json:"status"`
}

type V1ContainerStateWaiting struct {
    Reason      string      `json:"reason,omitempty"`
}

type V1ContainerStateRunning struct {
    StartedAt   string      `json:"startedAt,omitempty"`
}

type V1ContainerStateTerminated struct {
    ExitCode    int32       `json:"exitCode"`
    Signal      int32       `json:"signal,omitempty"`
    Reason      string      `json:"reason,omitempty"`
    Message     string      `json:"message,omitempty"`
    StartedAt   string      `json:"startedAt,omitempty"`
    FinishedAt  string      `json:"finishedAt,omitempty"`
    ContainerID string      `json:"containerID,omitempty"`
}

type V1ContainerState struct {
    Waiting     V1ContainerStateWaiting     `json:"waiting,omitempty"`
    Running     V1ContainerStateRunning     `json:"running,omitempty"`
    Terminated  V1ContainerStateTerminated  `json:"terminated,omitempty"`
}

type V1ContainerStatus struct {
    Name                string              `json:"name"`
    State               V1ContainerState    `json:"state,omitempty"`
    LastState           V1ContainerState    `json:"lastState,omitempty"`
    Ready               bool                `json:"ready"`
    RestartCount        int32               `json:"restartCount"`
    Image               string              `json:"image"`
    ImageID             string              `json:"imageID"`
    ContainerID         string              `json:"containerID,omitempty"`
}

type V1PodStatus struct {
    Phase               string              `json:"phase,omitempty"`
    Conditions          []V1PodCondition    `json:"conditions,omitempty"`
    Message             string              `json:"message,omitempty"`
    Reason              string              `json:"reason,omitempty"`
    HostIP              string              `json:"hostIP,omitempty"`
    PodIP               string              `json:"podIP,omitempty"`
    StartTime           string              `json:"startTime,omitempty"`
    ContainerStatuses   []V1ContainerStatus `json:"containerStatuses,omitempty"`
}

type V1Pod struct {
    Kind         string         `json:"kind"`
    ApiVersion   string         `json:"apiVersion"`
    Metadata     V1ObjectMeta   `json:"metadata"`
    Spec         V1PodSpec      `json:"spec"`
    Status       V1PodStatus    `json:"status,omitempty"`   
}

type PodCreation struct {
    Pretty       string     `json:"pretty,omitempty"`
    Body         V1Pod      `json:"body"`
    Namespace    string     `json:"namespace,omitempty"`      
}

type ByteList []byte

var index int = 0
/*
func (b *ByteList) ReadByte() (c byte, err error) {
    if []byte(*b) == nil { 
        return 0, errors.New("invalid data buffer")
    }
    if len([]byte(*b)) == index {
        return 0, io.EOF
    }
    index += 1
    return (*b)[index-1], nil
} 
*/
func (b ByteList) Read(p []byte) (n int, err error) {
    if []byte(b) == nil { 
        return 0, errors.New("invalid data buffer")
    }
    if len([]byte(b)) == index {
        return 0, io.EOF
    }
    if len([]byte(b)) - index < len(p) {
        n = copy(p, []byte(b)[index:])
        index += n
        return n, nil
    }
    n = copy(p, []byte(b)[index:index+len(p)])
    index += n
    return n, nil   
}

func main() {

    resp, err := http.Get("http://127.0.0.1:8080/api/v1/namespaces/default/pods/nginx-syeg7")

    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        fmt.Println(err)
        return
    }
    //fmt.Printf("%v\n", body)
  
    count := int64(bytes.Index(body, []byte{0}))
    if count == -1 {
        count = int64(len(body))
    }
    s := string(body[:count])
    fmt.Printf("Read string ----\n%q\n", s)

    index = 0
    count, err = io.Copy(os.Stdout, ByteList(body))

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("\nRead chars ----", count)

    var podConfig PodCreation
    var podBody *V1Pod = &podConfig.Body
    
    err = json.Unmarshal(body, podBody)
    
    if err != nil {
        fmt.Println("Unmarshal error: ", err)
        return 
    }
    fmt.Printf("Unmarshal ----\n%q\n", *podBody)
    //fmt.Println("{")

    //fmt.Printf("    %+V\n", *podBody)
    //fmt.Println("}")

    podConfig = *new(PodCreation)
    podBody = &podConfig.Body
    var podMetadata *V1ObjectMeta = &podBody.Metadata
    var podSpec *V1PodSpec = &podBody.Spec
    //var podStatus *V1PodStatus = &podBody.Status

    podConfig.Pretty = "true"
    podBody.ApiVersion = "v1"
    podBody.Kind = "Pod"
    podMetadata.Name = "hello-world-2"
    podSpec.RestartPolicy = "Never"
    podSpec.Containers = append(podSpec.Containers, V1Container{Name:"hello-world-2", Image:"ubuntu:14.04", Command:[]string{"/bin/echo", "hello", "world"}})

    var data []byte
    data, err = json.Marshal(podConfig)
    if err != nil {
        fmt.Println(err)
        return
    }

    count = int64(bytes.Index(data, []byte{0})) 
    if count == -1 {
        count = int64(len(data))
    }
    s = string(data[:count])
    fmt.Printf("Marshal string ----\n%q\n", s)

    index = 0
    count, err = io.Copy(os.Stdout, ByteList(data))
 
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("\nMarshal chars ----", count)
/*
    temp := []byte(`{
        "apiVersion": "v1",
        "kind": "Pod",
        "metadata": {
            "name": "hello-world",
        },
        "spec": { 
            "restartPolicy": "Never",
            "containers": [{
                "name": "hello",
                "image": "ubuntu:14.04",
                "command": ["/bin/echo","hello”,”world"],
            },],
        },
    }`)
*/      
    //index = 0
    //resp, err = http.Post("http://127.0.0.1:8080/api/v1/pods", "application/json", ByteList(temp))

    var podJSON bytes.Buffer
    if err = json.NewEncoder(&podJSON).Encode(*podBody); err != nil {
        fmt.Printf("failed to encode pod in json: %v", err)
        return
    }

    var req *http.Request
    req, err = http.NewRequest("POST", "http://127.0.0.1:8080/api/v1/namespaces/default/pods", &podJSON)
    if err != nil {
        fmt.Printf("failed to create request: POST %v", err)
        return
    }
   
    var client *http.Client = &http.Client{}
    resp, err = client.Do(req)
    if err != nil {
        fmt.Println("post error", err)
        return
    } else if resp.StatusCode >= 400 {
        fmt.Println("api server:", resp.Status)
        return
    }

    defer resp.Body.Close()
    body, err = ioutil.ReadAll(resp.Body)

    if err != nil {
        fmt.Println(err)
        return
    }
    
    podConfig = *new(PodCreation)
    podBody = &podConfig.Body

    err = json.Unmarshal(body, podBody)

    if err != nil {
        fmt.Println("Unmarshal error: ", err)
        return
    }
    fmt.Printf("Unmarshal ----\n%q\n", *podBody)

}   
