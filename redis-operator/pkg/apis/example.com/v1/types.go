package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ***************************************************************************
// IMPORTANT FOR CODE GENERATION
// If the types in this file are updated, you will need to run
// `make codegen` to generate the new types under the client/clientset folder.
// ***************************************************************************

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
