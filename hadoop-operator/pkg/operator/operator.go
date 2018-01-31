/*
   "github.com/rook/rook/pkg/operator"
*/

package operator

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreos/pkg/capnslog"
	opkit "github.com/rook/operator-kit"
	"github.com/rook/rook/pkg/clusterd"
	"github.com/rook/rook/pkg/daemon/agent/flexvolume/attachment"
	"github.com/rook/rook/pkg/operator/agent"
	"github.com/rook/rook/pkg/operator/cluster"
	"github.com/rook/rook/pkg/operator/k8sutil"
	"github.com/rook/rook/pkg/operator/mds"
	"github.com/rook/rook/pkg/operator/pool"
	"github.com/rook/rook/pkg/operator/provisioner"
	"github.com/rook/rook/pkg/operator/provisioner/controller"
	"github.com/rook/rook/pkg/operator/rgw"
	"k8s.io/api/core/v1"
)

const (
	initRetryDelay = 10 * time.Second
)

// volume provisioner constant
const (
	provisionerName = "rook.io/block"
)

var logger = capnslog.NewPackageLogger("github.com/rook/rook", "operator")

type Operator struct {
	agents map[string]*Agent
}

// New creates an operator instance
func New(context *clusterd.Context, volumeAttachmentWrapper attachment.Attachment, rookImage string) *Operator {
	clusterController := cluster.NewClusterController(context, rookImage, volumeAttachmentWrapper)
	volumeProvisioner := provisioner.New(context)

	schemes := []opkit.CustomResource{cluster.ClusterResource, pool.PoolResource, rgw.ObjectStoreResource,
		mds.FilesystemResource, attachment.VolumeAttachmentResource}
	return &Operator{
		context:           context,
		clusterController: clusterController,
		resources:         schemes,
		volumeProvisioner: volumeProvisioner,
		rookImage:         rookImage,
	}
}

// Run the operator instance
func (o *Operator) Run() error {

	namespace := os.Getenv(k8sutil.PodNamespaceEnvVar)
	if namespace == "" {
		return fmt.Errorf("Rook operator namespace is not provided. Expose it via downward API in the rook operator manifest file using environment variable %s", k8sutil.PodNamespaceEnvVar)
	}

	for {
		err := o.initResources()
		if err == nil {
			break
		}
		logger.Errorf("failed to init resources. %+v. retrying...", err)
		<-time.After(initRetryDelay)
	}

	rookAgent := agent.New(o.context.Clientset)

	if err := rookAgent.Start(namespace, o.rookImage); err != nil {
		return fmt.Errorf("Error starting agent daemonset: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	stopChan := make(chan struct{})
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Run volume provisioner
	// The controller needs to know what the server version is because out-of-tree
	// provisioners aren't officially supported until 1.5
	serverVersion, err := o.context.Clientset.Discovery().ServerVersion()
	if err != nil {
		return fmt.Errorf("Error getting server version: %v", err)
	}
	pc := controller.NewProvisionController(
		o.context.Clientset,
		provisionerName,
		o.volumeProvisioner,
		serverVersion.GitVersion,
	)
	go pc.Run(stopChan)
	logger.Infof("rook-provisioner started")

	// watch for changes to the rook clusters
	o.clusterController.StartWatch(v1.NamespaceAll, stopChan)

	for {
		select {
		case <-signalChan:
			logger.Infof("shutdown signal received, exiting...")
			close(stopChan)
			return nil
		}
	}
}

func (o *Operator) initResources() error {
	kitCtx := opkit.Context{
		Clientset:             o.context.Clientset,
		APIExtensionClientset: o.context.APIExtensionClientset,
		Interval:              500 * time.Millisecond,
		Timeout:               60 * time.Second,
	}

	// Create and wait for CRD resources
	err := opkit.CreateCustomResources(kitCtx, o.resources)
	if err != nil {
		return fmt.Errorf("failed to create custom resource. %+v", err)
	}
	return nil
}
