/*
	"github.com/rook/rook/pkg/operator/k8sutil"
*/

package kubehelper

const (
	// Namespace for rook
	Namespace = "rook"
	// DefaultNamespace for the cluster
	DefaultNamespace = "default"
	// DataDirVolume data dir volume
	DataDirVolume = "rook-data"
	// DataDir folder
	DataDir = "/var/lib/rook"
	// RookType for the CRD
	RookType = "kubernetes.io/rook"
	// PodNameEnvVar is the env variable for getting the pod name via downward api
	PodNameEnvVar = "POD_NAME"
	// PodNamespaceEnvVar is the env variable for getting the pod namespace via downward api
	PodNamespaceEnvVar = "POD_NAMESPACE"
	// NodeNameEnvVar is the env variable for getting the node via downward api
	NodeNameEnvVar = "NODE_NAME"
)
