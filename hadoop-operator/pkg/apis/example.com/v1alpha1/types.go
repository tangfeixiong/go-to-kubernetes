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

type Hadoop struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              HadoopSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HadoopList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Hadoop `json:"items"`
}

type HadoopSpec struct {
	Core    string `json:"core"`
	Common  string `json:"common"`
	Version string `json:"version,omitempty"`
	Image   string `json:"image,omitempty"`

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

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HadoopHdfs struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdfsSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HadoopHdfsList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HadoopHdfs `json:"items"`
}

type HdfsSpec struct {
	DataReplication    *int32        `json:"dataReplication,omitempty"`
	ReplicaReplacement string        `json:"replicaReplacement,omitempty"`
	Count              *int32        `json:"count,omitempty"`
	NameNodeTemplate   *NodeTemplate `json:"nameNodeTemplate`
	DataNodeTemplate   *NodeTemplate `json:"nameNodeTemplate`
}

type NodeTemplate struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeSpec `json:"spec"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HadoopNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeSpec   `json:"spec"`
	Status            NodeStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HadoopNodeList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HadoopNode `json:"items"`
}

type NodeType string

const (
	NameNode          NodeType = "namenode"
	DataNode          NodeType = "datanode"
	SecondaryNameNode NodeType = "secondarynamenode"
)

type NodeSpec struct {
	Type NodeType `json:"type"`
}

type NodeStatus struct {
	Phase string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HadoopYarn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              YarnSpec `json:"spec"`
	//Status            Status   `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HadoopYarnList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HadoopYarn `json:"items"`
}

type YarnSpec struct {
	Count *int32 `json:"count,omitempty"`

	StatefulSetName string `json:"statefulSetName,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	Pod *PodPolicy `json:"pod,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HadoopMapreduce struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MapReduceSpec `json:"spec"`
	//Status            Status        `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HadoopMapreduceList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HadoopMapreduce `json:"items"`
}

type MapReduceSpec struct {
	Type  NodeType `json:"type"`
	Count *int32   `json:"count,omitempty"`

	StatefulSetName string `json:"statefulSetName,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	Pod *PodPolicy `json:"pod,omitempty"`
}

type PodPolicy struct {
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
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
