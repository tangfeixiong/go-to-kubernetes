package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
  Inspired by:
    https://github.com/rook/rook/blob/master/pkg/apis/rook.io/v1alpha1/types.go
*/

// ***************************************************************************
// IMPORTANT FOR CODE GENERATION
// If the types in this file are updated, you will need to run
// `make codegen` to generate the new types under the client/clientset folder.
// ***************************************************************************

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ClusterSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Cluster `json:"items"`
}

type ClusterMode string

const (
	MasterSlaveModel ClusterMode = "ms"
	RedisSentinelHA  ClusterMode = "rs"
)

type RedisTemplate struct {
	ReplicationSlaves *int32 `json:"replicationSlaves,omitempty"`
	HashSlots         *int32 `json:"hashSlots,omitempty"`
	StatefulSetName   string `json:"statefulSetName,omitempty"`
	ServiceName       string `json:"serviceName,omitempty"`
}

type SentinelTemplate struct {
	Quorum          *int32 `json:"quorum,omitempty`
	MasterGroupName string `json:"masterGroupName,omitempty"`
	DeploymentName  string `json:"deploymentName,omitempty"`
	ServiceName     string `json:"serviceName,omitempty"`
}

type ClusterSpec struct {
	RedisTemplate    *RedisTemplate    `json:"redisTemplate"`
	SentinelTemplate *SentinelTemplate `json:"sentinelTemplate"`
	Mode             ClusterMode       `json:"clusterMode,omitempty"`
	Image            string            `json:"image,omitempty"`

	// The type of backend for the cluster (only "ceph" is implemented)
	//Backend string `json:"backend,omitempty"`

	// The path on the host where config and data can be persisted.
	DataDirHostPath string `json:"dataDirHostPath,omitempty"`

	// The placement-related configuration to pass to kubernetes (affinity, node selector, tolerations).
	Placement PlacementSpec `json:"placement,omitempty"`

	// A spec for available storage in the cluster and how it should be used
	//Storage StorageSpec `json:"storage"`

	// HostNetwork to enable host network
	HostNetwork bool `json:"hostNetwork,omitempty"`

	// MonCount sets the mon size
	//MonCount int `json:"monCount"`

	// Resources set resource requests and limits
	Resources ResourceSpec `json:"resources,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PlacementSpec struct {
	metav1.TypeMeta `json:",inline"`
	All             Placement `json:"all,omitempty"`
	API             Placement `json:"api,omitempty"`
	Redis           Placement `json:"redis,omitempty"`
	Sentinel        Placement `json:"sentinel,omitempty"`
}

// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Placement encapsulates the various kubernetes options that control where pods are scheduled and executed.
type Placement struct {
	metav1.TypeMeta `json:",inline"`
	NodeAffinity    *v1.NodeAffinity    `json:"nodeAffinity,omitempty"`
	PodAffinity     *v1.PodAffinity     `json:"podAffinity,omitempty"`
	PodAntiAffinity *v1.PodAntiAffinity `json:"podAntiAffinity,omitempty"`
	Tolerations     []v1.Toleration     `json:"tolerations,omitemtpy"`
}

type ResourceSpec struct {
	API      v1.ResourceRequirements `json:"api,omitempty"`
	Redis    v1.ResourceRequirements `json:"redis,omitempty"`
	Sentinel v1.ResourceRequirements `json:"sentinel,omitempty"`
}

/*
  Inspired by:
    https://github.com/jw-s/redis-operator/blob/master/pkg/apis/redis/v1/redis.go
*/

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RedisList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Redis `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Redis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status"`
}

type ServerSpec struct {
	Sentinels SentinelSpec `json:"sentinels"`

	Slaves SlaveSpec `json:"slaves"`

	BaseImage string `json:"baseImage,omitempty"`

	Version string `json:"version,omitempty"`

	Paused bool `json:"paused,omitempty"`

	Pod *PodPolicy `json:"pod,omitempty"`
}

type SentinelSpec struct {
	Replicas  int32     `json:"replicas"`
	Quorum    int32     `json:"quorum"`
	ConfigMap ConfigMap `json:"configMap"`
}

type SlaveSpec struct {
	Replicas  int32     `json:"replicas"`
	ConfigMap ConfigMap `json:"configMap"`
}

type ConfigMap string

type PodPolicy struct {
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
}

type ServerStatus struct {
	Phase          ServerPhase       `json:"phase"`
	Conditions     []ServerCondition `json:"conditions"`
	SlaveStatus    SlaveStatus       `json:"slaves"`
	SentinelStatus SentinelStatus    `json:"sentinels"`
}

type SlaveStatus struct {
	Ready   []string `json:"ready,omitempty"`
	Unready []string `json:"unready,omitempty"`
}

type SentinelStatus struct {
	Ready   []string `json:"ready,omitempty"`
	Unready []string `json:"unready,omitempty"`
}

type ServerPhase string

const (
	ServerCreatingPhase ServerPhase = "Creating"
	ServerStoppingPhase ServerPhase = "Stopping"
	ServerRunningPhase  ServerPhase = "Running"
	ServerFailedPhase   ServerPhase = "Failed"
)

type ServerConditionType string

type ServerCondition struct {
	Type   ServerConditionType `json:"type"`
	Reason string              `json:"reason,omitempty"`
}

const (
	ServerConditionAddSeedMaster    ServerConditionType = "AddingSeedMaster"
	ServerConditionRemoveSeedMaster ServerConditionType = "removingSeedMaster"
	ServerConditionAddSentinel      ServerConditionType = "AddingSentinel"
	ServerConditionRemoveSentinel   ServerConditionType = "removingSentinel"
	ServerConditionAddSlave         ServerConditionType = "AddingSlave"
	ServerConditionRemoveSlave      ServerConditionType = "removingSlave"
	ServerConditionReady            ServerConditionType = "Ready"
)
