/* Inspired by:
   https://github.com/kubernetes/sample-controller/blob/master/controller.go
*/

package controller

import (
	"fmt"
	//"os"
	//"os/signal"
	//"syscall"
	"time"

	"github.com/golang/glog"
	"github.com/sirupsen/logrus"
	//"github.com/streadway/amqp"

	appsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	// "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	kubeinformers "k8s.io/client-go/informers"
	//appsv1beta2informer "k8s.io/client-go/informers/apps/v1beta2"
	//v1informer "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisters "k8s.io/client-go/listers/apps/v1beta2"
	v1lister "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	//"k8s.io/kubernetes/plugin/pkg/scheduler/api"

	samplev1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	clientset "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned"
	samplescheme "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned/scheme"
	informers "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/informers/externalversions"
	listers "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/listers/example.com/v1alpha1"
	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/spec/sts"
	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/spec/svc"
)

const controllerAgentName = "hadoop-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when a Foo is synced
	SuccessSynced = "Synced"
	// ErrResourceExists is used as part of the Event 'reason' when a Foo fails
	// to sync due to a Deployment of the same name already existing.
	ErrResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	MessageResourceExists = "Resource %q already exists and is not managed by Foo"
	// MessageResourceSynced is the message used for an Event fired when a Foo
	// is synced successfully
	MessageResourceSynced = "Foo synced successfully"
)

type Controller struct {
	kubeclientset     kubernetes.Interface
	sampleclientset   clientset.Interface
	podLister         v1lister.PodLister
	serviceLister     v1lister.ServiceLister
	endpointsLister   v1lister.EndpointsLister
	configMapLister   v1lister.ConfigMapLister
	deploymentsLister appslisters.DeploymentLister
	statefulSetLister appslisters.StatefulSetLister
	foosLister        listers.HadoopHdfsLister
	cacheSyncs        []cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
	logger   *logrus.Entry
	//Config
	Config *struct {
		*rest.Config
		DefaultResync time.Duration
	}
	//redises     map[string]*redis.Redis
	//controllerConfig controllerConfigFunc
}

//type controllerConfigFunc func(*Controller)

func NewController(
	kubeclientset kubernetes.Interface,
	sampleclientset clientset.Interface,
	kubeInformerFactory kubeinformers.SharedInformerFactory,
	sampleInformerFactory informers.SharedInformerFactory,
) *Controller {

	// obtain references to shared index informers for the Deployment and Foo
	// types.
	podInformer := kubeInformerFactory.Core().V1().Pods()
	serviceInformer := kubeInformerFactory.Core().V1().Services()
	endpointInformer := kubeInformerFactory.Core().V1().Endpoints()
	configMapInformer := kubeInformerFactory.Core().V1().ConfigMaps()
	statefulSetInformer := kubeInformerFactory.Apps().V1beta2().StatefulSets()
	deploymentInformer := kubeInformerFactory.Apps().V1beta2().Deployments()

	fooInformer := sampleInformerFactory.Example().V1alpha1().HadoopHdfses()

	// Create event broadcaster
	// Add sample-controller types to the default Kubernetes Scheme so Events can be
	// logged for sample-controller types.
	samplescheme.AddToScheme(scheme.Scheme)
	glog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	cacheSyncs := []cache.InformerSynced{
		podInformer.Informer().HasSynced,
		deploymentInformer.Informer().HasSynced,
		serviceInformer.Informer().HasSynced,
		endpointInformer.Informer().HasSynced,
		configMapInformer.Informer().HasSynced,
		statefulSetInformer.Informer().HasSynced,
		fooInformer.Informer().HasSynced,
	}

	controller := &Controller{
		kubeclientset:     kubeclientset,
		sampleclientset:   sampleclientset,
		podLister:         podInformer.Lister(),
		serviceLister:     serviceInformer.Lister(),
		endpointsLister:   endpointInformer.Lister(),
		configMapLister:   configMapInformer.Lister(),
		deploymentsLister: deploymentInformer.Lister(),
		statefulSetLister: statefulSetInformer.Lister(),
		foosLister:        fooInformer.Lister(),
		cacheSyncs:        cacheSyncs,
		recorder:          recorder,
		workqueue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "hadoophdfses"),
		logger:            logrus.WithField("pkg", "controller"),
		//Config:            cfg,
		//redises: make(map[string]*redis.Redis),
	}

	glog.Info("Setting up event handlers")
	// Set up an event handler for when Foo resources change
	fooInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueFoo,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueFoo(new)
		},
	})

	// Set up an event handler for when Deployment/StatefulSet resources change. This
	// handler will lookup the owner of the given Deployment/StatefulSet, and if it is
	// owned by a Foo resource will enqueue that Foo resource for
	// processing. This way, we don't need to implement custom logic for
	// handling Deployment/StatefulSet resources. More info on this pattern:
	// https://github.com/kubernetes/community/blob/8cafef897a22026d42f5e5bb3f104febe7e29830/contributors/devel/controllers.md
	statefulSetInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleObject,
		UpdateFunc: func(old, new interface{}) {
			newSts := new.(*appsv1beta2.StatefulSet)
			oldSts := old.(*appsv1beta2.StatefulSet)
			if newSts.ResourceVersion == oldSts.ResourceVersion {
				// Periodic resync will send update events for all known Deployments/StatefulSets.
				// Two different versions of the same Deployment/StatefulSet will always have different RVs.
				return
			}
			controller.handleObject(new)
		},
		DeleteFunc: controller.handleObject,
	})

	//	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	//		AddFunc: controller.handleObject,
	//		UpdateFunc: func(old, new interface{}) {
	//			newDepl := new.(*appsv1beta2.Deployment)
	//			oldDepl := old.(*appsv1beta2.Deployment)
	//			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
	//				return
	//			}
	//			controller.handleObject(new)
	//		},
	//		DeleteFunc: controller.handleObject,
	//	})

	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleObject,
		UpdateFunc: func(old, new interface{}) {
			newSvc := new.(*corev1.Service)
			oldSvc := old.(*corev1.Service)
			if newSvc.ResourceVersion == oldSvc.ResourceVersion {
				return
			}
			controller.handleObject(new)
		},
		DeleteFunc: controller.handleObject,
	})

	return controller

}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	//glog.Info("Starting Foo controller")
	c.logger.Info("Starting controller")

	defer c.logger.Info("Exiting controller")

	// Wait for the caches to be synced before starting workers
	glog.Info("Waiting for informer caches to sync")
	if !cache.WaitForCacheSync(stopCh, c.cacheSyncs...) {
		c.logger.Fatal("Timeout waiting for cache to sync")
	}
	c.logger.Info("Sync completed")

	glog.Info("Starting workers")
	// Launch two workers to process Foo resources
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	glog.Info("Started workers")
	<-stopCh
	glog.Info("Shutting down workers")

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		c.logger.Info("Shutting down queue")
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(obj)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}

		// Run the syncHandler, passing it the namespace/name string of the
		// Foo resource to be synced.
		if err := c.syncHandler(key); err != nil {
			return fmt.Errorf("error syncing '%s': %s", key, err.Error())
		}

		c.logger.Debugf("Working on %s", obj)

		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		glog.Infof("Successfully synced '%s'", key)
		return nil

		// err := c.process(k.(string))

		// if err == nil {
		// 	c.logger.Debugf("Finished Working on %s", k)
		// 	c.queue.Forget(k)
		// 	return true
		// }

		// c.logger.Errorf("Re-queueing as encountered an error: %s", err.Error())

		// c.queue.AddRateLimited(k)
	}(obj)

	if err != nil {
		runtime.HandleError(err)
		return true
	}

	return true
}

// syncHandler compares the actual state with the desired, and attempts to
// converge the two. It then updates the Status block of the Foo resource
// with the current status of the resource.
func (c *Controller) syncHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	ns, name, err := cache.SplitMetaNamespaceKey(key)

	if err != nil {
		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
		// return err
	}

	// Get the Foo resource with this namespace/name
	foo, err := c.foosLister.HadoopHdfses(ns).Get(name)

	if err != nil {
		// The Foo resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			runtime.HandleError(fmt.Errorf("foo '%s' in work queue no longer exists", key))
			return nil
		}
		// if util.ResourceNotFoundError(err) {
		// 	 delete(c.redises, name)
		// 	 return c.deleteResources(ns, name)
		// }

		return err
	}

	serviceName := foo.Name
	//	if foo.Spec.ServiceName != "" {
	//		serviceName = foo.Spec.ServiceName
	//	}
	service, err := c.serviceLister.Services(foo.Namespace).Get(serviceName)
	if errors.IsNotFound(err) {
		recipe := svc.NewHadoopHdfsNodeServiceRecipient(serviceName, foo.Namespace, foo.Name)
		artifact, err := recipe.New()
		if err != nil {
			runtime.HandleError(fmt.Errorf("Internel error: %s", err.Error))
			return nil
		}
		artifact.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
			*metav1.NewControllerRef(foo, schema.GroupVersionKind{
				Group:   samplev1alpha1.SchemeGroupVersion.Group,
				Version: samplev1alpha1.SchemeGroupVersion.Version,
				Kind:    "HadoopHdfs",
			}),
		}

		service, err = c.kubeclientset.CoreV1().Services(foo.Namespace).Create(artifact)
	}

	if err != nil {
		glog.V(2).Info(err)
		return err
	}

	if !metav1.IsControlledBy(service, foo) {
		msg := fmt.Sprintf(MessageResourceExists, service.Name)
		c.recorder.Event(foo, corev1.EventTypeWarning, ErrResourceExists, msg)
		return fmt.Errorf(msg)
	}

	// deploymentName := foo.Spec.DeploymentName
	statefulSetName := foo.Name
	//if deploymentName == "" {
	if statefulSetName == "" {
		// We choose to absorb the error here as the worker would requeue the
		// resource otherwise. Instead, the next time the resource is updated
		// the resource will be queued again.
		statefulSetName = foo.Name
		//key := foo.Name
		//runtime.HandleError(fmt.Errorf("%s: deployment name must be specified", key))
		//return nil
	}
	count := int32(1)
	//	if foo.Spec.Count != nil {
	//		if *foo.Spec.Count < 3 || *foo.Spec.Count > 15 {
	//			runtime.HandleError(fmt.Errorf("Cluster members must be between 3 and 5"))
	//			return nil
	//		}
	//	}
	//	count = int32(*foo.Spec.Count)

	//var artifact extensions.Deployment
	var artifact *appsv1beta2.StatefulSet

	// Get the deployment with the name specified in Foo.spec
	//deployment, err := c.deploymentsLister.Deployments(foo.Namespace).Get(deploymentName)
	statefulSet, err := c.statefulSetLister.StatefulSets(foo.Namespace).Get(statefulSetName)
	// If the resource doesn't exist, we'll create it
	if errors.IsNotFound(err) {
		recipe := sts.NewHadoopHdfsNodeRecipient(statefulSetName, foo.Namespace, foo.Name, serviceName, "", "namenode", &count)
		artifact, err = recipe.New()
		if err != nil {
			runtime.HandleError(fmt.Errorf("Internel error: %s", err.Error()))
			return nil
		}
		artifact.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
			*metav1.NewControllerRef(foo, schema.GroupVersionKind{
				Group:   samplev1alpha1.SchemeGroupVersion.Group,
				Version: samplev1alpha1.SchemeGroupVersion.Version,
				Kind:    "HadoopHdfs",
			}),
		}

		//deployment, err = c.kubeclientset.ExtensionsV1beta1().Deployments(foo.Namespace).Create(newDeployment(foo))
		//deployment, err = c.kubeclientset.AppsV1beta2().Deployments(foo.Namespace).Create(artifact)
		statefulSet, err = c.kubeclientset.AppsV1beta2().StatefulSets(foo.Namespace).Create(artifact)
	}

	// If an error occurs during Get/Create, we'll requeue the item so we can
	// attempt processing again later. This could have been caused by a
	// temporary network failure, or any other transient reason.
	if err != nil {
		glog.V(2).Info(err)
		return err
	}

	// If the Deployment is not controlled by this Foo resource, we should log
	// a warning to the event recorder and ret
	if !metav1.IsControlledBy( /*deployment*/ statefulSet, foo) {
		msg := fmt.Sprintf(MessageResourceExists /*deployment*/, statefulSet.Name)
		c.recorder.Event(foo, corev1.EventTypeWarning, ErrResourceExists, msg)
		return fmt.Errorf(msg)
	}

	// If this number of the replicas on the Foo resource is specified, and the
	// number does not equal the current desired replicas on the Deployment, we
	// should update the Deployment resource.
	if 1 != *statefulSet.Spec.Replicas {
		glog.V(4).Infof("Replicas: %d", *statefulSet.Spec.Replicas)
		//deployment, err = c.kubeclientset.ExtensionsV1beta1().Deployments(foo.Namespace).Update(newDeployment(foo))
		//deployment, err = c.kubeclientset.AppsV1beta2().Deployments(foo.Namespace).Update(artifact)
		statefulSet, err = c.kubeclientset.AppsV1beta2().StatefulSets(foo.Namespace).Update(artifact)
	}

	// If an error occurs during Update, we'll requeue the item so we can
	// attempt processing again later. THis could have been caused by a
	// temporary network failure, or any other transient reason.
	if err != nil {
		glog.V(2).Info(err)
		return err
	}

	// Finally, we update the status block of the Foo resource to reflect the
	// current state of the world
	//	err = c.updateFooStatus(foo, statefulSet)
	//	if err != nil {
	//		return err
	//	}

	c.recorder.Event(foo, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func (c *Controller) updateFooStatus(foo *samplev1alpha1.HadoopHdfs /*deployment *appsv1beta2.Deployment*/, statefuleSet *appsv1beta2.StatefulSet) error {
	// NEVER modify objects from the store. It's a read-only, local cache.
	// You can use DeepCopy() to make a deep copy of original object and modify this copy
	// Or create a copy manually for better performance
	fooCopy := foo.DeepCopy()
	//fooCopy.Status.AvailableReplicas = deployment.Status.AvailableReplicas

	// Until #38113 is merged, we must use Update instead of UpdateStatus to
	// update the Status block of the Foo resource. UpdateStatus will not
	// allow changes to the Spec of the resource, which is ideal for ensuring
	// nothing other than resource status has been updated.
	_, err := c.sampleclientset.ExampleV1alpha1().HadoopHdfses(foo.Namespace).Update(fooCopy)
	return err
}

// enqueueFoo takes a Foo resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than Foo.
func (c *Controller) enqueueFoo(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}
	c.workqueue.AddRateLimited(key)
}

// handleObject will take any resource implementing metav1.Object and attempt
// to find the Foo resource that 'owns' it. It does this by looking at the
// objects metadata.ownerReferences field for an appropriate OwnerReference.
// It then enqueues that Foo resource to be processed. If the object does not
// have an appropriate OwnerReference, it will simply be skipped.
func (c *Controller) handleObject(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			runtime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			runtime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		glog.V(4).Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	glog.V(4).Infof("Processing object: %s", object.GetName())
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		// If this object is not owned by a Foo, we should not do anything more
		// with it.
		if ownerRef.Kind != "Foo" {
			return
		}

		foo, err := c.foosLister.HadoopHdfses(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			glog.V(4).Infof("ignoring orphaned object '%s' of foo '%s'", object.GetSelfLink(), ownerRef.Name)
			return
		}

		c.enqueueFoo(foo)
		return
	}
}

// newDeployment creates a new Deployment for a Foo resource. It also sets
// the appropriate OwnerReferences on the resource so handleObject can discover
// the Foo resource that 'owns' it.
//func newDeployment(foo *samplev1alpha1.Foo) *appsv1beta2.Deployment {
//	labels := map[string]string{
//		"app":        "nginx",
//		"controller": foo.Name,
//	}
//	return &appsv1beta2.Deployment{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      foo.Spec.DeploymentName,
//			Namespace: foo.Namespace,
//			OwnerReferences: []metav1.OwnerReference{
//				*metav1.NewControllerRef(foo, schema.GroupVersionKind{
//					Group:   samplev1alpha1.SchemeGroupVersion.Group,
//					Version: samplev1alpha1.SchemeGroupVersion.Version,
//					Kind:    "Foo",
//				}),
//			},
//		},
//		Spec: appsv1beta2.DeploymentSpec{
//			Replicas: foo.Spec.Replicas,
//			Selector: &metav1.LabelSelector{
//				MatchLabels: labels,
//			},
//			Template: corev1.PodTemplateSpec{
//				ObjectMeta: metav1.ObjectMeta{
//					Labels: labels,
//				},
//				Spec: corev1.PodSpec{
//					Containers: []corev1.Container{
//						{
//							Name:  "nginx",
//							Image: "nginx:latest",
//						},
//					},
//				},
//			},
//		},
//	}
//}
