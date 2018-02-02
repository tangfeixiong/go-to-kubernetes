package v1alpha1

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

type Mariadb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              MariadbSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MariadbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Mariadb `json:"items"`
}

type ClusterMode string

const (
	Galera     ClusterMode = "Galera"
	ClusterCGE ClusterMode = "ClusterCGE"
	None       ClusterMode = "None"
)

type MysqlTemplate struct {
	StatefulSetName string `json:"statefulSetName,omitempty"`
	ServiceName     string `json:"serviceName,omitempty"`
}

type GaleraTemplate struct {
	ClusterName string `json:"clusterName,omitempty"`
	Count       *int32 `json:"count,omitempty`
}

type MariadbSpec struct {
	MysqlTemplate  *MysqlTemplate  `json:"mysqlTemplate"`
	GaleraTemplate *GaleraTemplate `json:"galeraTemplate"`
	Mode           ClusterMode     `json:"clusterMode,omitempty"`
	Image          string          `json:"image,omitempty"`

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
	DB              Placement `json:"db,omitempty"`
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
	DB     v1.ResourceRequirements `json:"mgr,omitempty"`
}

/*
  Inspired by:
    https://github.com/jw-s/redis-operator/blob/master/pkg/apis/redis/v1/redis.go
*/

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MysqlList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mysql `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Mysql struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status"`
}

type ServerSpec struct {
	MysqlSpec MysqlSpec `json:"mysqlSpec"`

	ClusterSpec ClusterSpec `json:"clusterSpec"`

	BaseImage string `json:"baseImage,omitempty"`

	Version string `json:"version,omitempty"`

	Paused bool `json:"paused,omitempty"`

	Pod *PodPolicy `json:"pod,omitempty"`
}

type MysqlSpec struct {
	StatefulSetName string    `json:"statefulsetName"`
	ServiceName     string    `json:"serviceName"`
	ConfigMap       ConfigMap `json:"configMap"`
}

type ClusterSpec struct {
	Count     int32     `json:"count"`
	ConfigMap ConfigMap `json:"configMap"`
}

type ConfigMap string

type PodPolicy struct {
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
}

type ServerStatus struct {
	Phase        ServerPhase       `json:"phase"`
	Conditions   []ServerCondition `json:"conditions"`
	SlaveStatus  SlaveStatus       `json:"slaves"`
	MasterStatus MasterStatus      `json:"masters"`
}

type SlaveStatus struct {
	Ready   []string `json:"ready,omitempty"`
	Unready []string `json:"unready,omitempty"`
}

type MasterStatus struct {
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
	ServerConditionJoinMaster       ServerConditionType = "JoiningMaster"
	ServerConditionRemoveMaster     ServerConditionType = "removingMaster"
	ServerConditionAddSlave         ServerConditionType = "AddingSlave"
	ServerConditionRemoveSlave      ServerConditionType = "removingSlave"
	ServerConditionReady            ServerConditionType = "Ready"
)
