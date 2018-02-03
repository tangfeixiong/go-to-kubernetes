package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
  Inspired by:
    https://github.com/jw-s/redis-operator/blob/master/pkg/apis/redis/v1/redis.go
*/

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RabbitmqList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rabbitmq `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Rabbitmq struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Spec   `json:"spec"`
	Status            Status `json:"status"`
}

type Spec struct {
	Count *int32 `json:"count,omitempty"`

	Mode ClusteringMode `json:"mode,omitempty"`

	StatefulSetName string `json:"statefulSetName,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	Pod *PodPolicy `json:"pod,omitempty"`
}

type ConfigMap string

type PodPolicy struct {
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
}

type Status struct {
	Phase      Phase       `json:"phase"`
	Conditions []Condition `json:"conditions"`
	Ready      []string    `json:"ready,omitempty"`
	Unready    []string    `json:"unready,omitempty"`
}

type Phase string

const (
	CreatingPhase Phase = "Creating"
	StoppingPhase Phase = "Stopping"
	RunningPhase  Phase = "Running"
	FailedPhase   Phase = "Failed"
)

type ConditionType string

type Condition struct {
	Type   ConditionType `json:"type"`
	Reason string        `json:"reason,omitempty"`
}

const (
	ConditionAddSeedMaster    ConditionType = "AddingSeedMaster"
	ConditionRemoveSeedMaster ConditionType = "removingSeedMaster"
	ConditionJoinMaster       ConditionType = "JoiningMaster"
	ConditionRemoveMaster     ConditionType = "removingMaster"
	ConditionAddSlave         ConditionType = "AddingSlave"
	ConditionRemoveSlave      ConditionType = "removingSlave"
	ConditionReady            ConditionType = "Ready"
)

type ClusteringMode string

const (
	PeerDiscovery_K8s ClusteringMode = "Peer_Discovery_K8s"
	Native            ClusteringMode = "Native"
	None              ClusteringMode = "None"
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

type Amqp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              QueueSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AmqpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Amqp `json:"items"`
}

type QueueSpec struct {
	Count           *int32         `json:"count,omitempty"`
	StatefulSetName string         `json:"statefulSetName,omitempty"`
	ServiceName     string         `json:"serviceName,omitempty"`
	Mode            ClusteringMode `json:"mode,omitempty"`
	Image           string         `json:"image,omitempty"`

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
	SERVER          Placement `json:"server,omitempty"`
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
	API    v1.ResourceRequirements `json:"api,omitempty"`
	SERVER v1.ResourceRequirements `json:"server,omitempty"`
}
