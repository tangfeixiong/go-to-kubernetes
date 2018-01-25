/*
  Inspired by:
    https://github.com/jw-s/redis-operator/blob/master/pkg/operator/controller/reconcile.go
*/

package controller

import (
	//"encoding/json"
	"fmt"
	//"reflect"

	"github.com/golang/glog"

	appsv1 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	//"k8s.io/apimachinery/pkg/types"
	//"k8s.io/apimachinery/pkg/util/strategicpatch"
	//"k8s.io/kubernetes/plugin/pkg/scheduler/api"

	samplev1 "github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1"
	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/sts"
	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/svc"
)

func (c *Controller) handleClusterService(foo *samplev1.Cluster) (*corev1.Service, error) {
	serviceName := foo.Name
	if foo.Spec.RedisTemplate != nil && foo.Spec.RedisTemplate.ServiceName != "" {
		serviceName = foo.Spec.RedisTemplate.ServiceName
	}
	service, err := c.serviceLister.Services(foo.Namespace).Get(serviceName)
	if errors.IsNotFound(err) {
		recipe := svc.NewRedisRecipient(serviceName, foo.Namespace, foo.Name, "client", "gossip", 6379)
		artifact, err := recipe.GenerateArtifact()
		if err != nil {
			runtime.HandleError(fmt.Errorf("Internel error: %s", err.Error))
			return nil, nil
		}
		artifact.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
			*metav1.NewControllerRef(foo, schema.GroupVersionKind{
				Group:   samplev1.SchemeGroupVersion.Group,
				Version: samplev1.SchemeGroupVersion.Version,
				Kind:    "Cluster",
			}),
		}

		service, err = c.kubeclientset.CoreV1().Services(foo.Namespace).Create(artifact)
	}

	if err != nil {
		glog.V(2).Infoln(err)
		return service, err
	}

	if !metav1.IsControlledBy(service, foo) {
		msg := fmt.Sprintf(MessageResourceExists, service.Name)
		c.recorder.Event(foo, corev1.EventTypeWarning, ErrResourceExists, msg)
		return service, fmt.Errorf(msg)
	}
	return service, nil
}

func (c *Controller) handleClusterStatefulSet(foo *samplev1.Cluster) (*appsv1.StatefulSet, error) {
	statefulSetName := foo.Name
	if foo.Spec.RedisTemplate == nil {
		runtime.HandleError(fmt.Errorf("%s: Redis master and slaves must be specified", foo.Name))
		return nil, nil
	}
	replicationSlaves := 1
	if foo.Spec.RedisTemplate.ReplicationSlaves != nil {
		if *foo.Spec.RedisTemplate.ReplicationSlaves < 1 || *foo.Spec.RedisTemplate.ReplicationSlaves > 5 {
			runtime.HandleError(fmt.Errorf("Redis slaves must be between 1 and 5"))
			return nil, nil
		}
	}
	replicationSlaves = int(*foo.Spec.RedisTemplate.ReplicationSlaves)

	var artifact *appsv1.StatefulSet

	if foo.Spec.RedisTemplate.StatefulSetName != "" {
		statefulSetName = foo.Spec.RedisTemplate.StatefulSetName
	}
	statefulSet, err := c.statefulSetLister.StatefulSets(foo.Namespace).Get(statefulSetName)
	if errors.IsNotFound(err) {
		recipe := sts.NewRedisRecipient(statefulSetName, foo.Namespace, foo.Name, foo.Name, replicationSlaves)
		artifact, err = recipe.GenerateArtifact()
		if err != nil {
			runtime.HandleError(fmt.Errorf("Internel error: %s", err.Error))
			return nil, nil
		}
		artifact.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
			*metav1.NewControllerRef(foo, schema.GroupVersionKind{
				Group:   samplev1.SchemeGroupVersion.Group,
				Version: samplev1.SchemeGroupVersion.Version,
				Kind:    "Cluster",
			}),
		}

		statefulSet, err = c.kubeclientset.AppsV1beta2().StatefulSets(foo.Namespace).Create(artifact)
	}

	if err != nil {
		glog.V(2).Infoln(err)
		return statefulSet, err
	}

	if !metav1.IsControlledBy(statefulSet, foo) {
		msg := fmt.Sprintf(MessageResourceExists, statefulSet.Name)
		c.recorder.Event(foo, corev1.EventTypeWarning, ErrResourceExists, msg)
		return statefulSet, fmt.Errorf(msg)
	}

	//	if foo.Spec.RedisTemplate.ReplicationSlaves != nil && *foo.Spec.RedisTemplate.ReplicationSlaves != *statefulSet.Spec.Replicas {
	//		glog.V(4).Infof("Foor: %d, deplR: %d", *foo.Spec.Sentinels.Replicas, *deployment.Spec.Replicas)
	//		statefulSet, err = c.kubeclientset.AppsV1beta2().StatefulSets(foo.Namespace).Update(artifact)
	//	}

	if err != nil {
		return statefulSet, err
	}

	return statefulSet, nil
}

func (c *Controller) handleSentinelService(foo *samplev1.Cluster) (*corev1.Service, error) {
	serviceName := fmt.Sprintf("%s-ha", foo.Name)
	if foo.Spec.SentinelTemplate != nil && foo.Spec.SentinelTemplate.ServiceName != "" && foo.Spec.SentinelTemplate.ServiceName != foo.Name {
		serviceName = foo.Spec.SentinelTemplate.ServiceName
	}
	service, err := c.serviceLister.Services(foo.Namespace).Get(serviceName)
	if errors.IsNotFound(err) {
		recipe := svc.NewSentinelRecipient(serviceName, foo.Namespace, foo.Name, "client", 26379)
		artifact, err := recipe.GenerateArtifact()
		if err != nil {
			runtime.HandleError(fmt.Errorf("Internel error: %s", err.Error))
			return nil, nil
		}
		artifact.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
			*metav1.NewControllerRef(foo, schema.GroupVersionKind{
				Group:   samplev1.SchemeGroupVersion.Group,
				Version: samplev1.SchemeGroupVersion.Version,
				Kind:    "Cluster",
			}),
		}

		service, err = c.kubeclientset.CoreV1().Services(foo.Namespace).Create(artifact)
	}

	if err != nil {
		glog.V(2).Infoln(err)
		return service, err
	}

	if !metav1.IsControlledBy(service, foo) {
		msg := fmt.Sprintf(MessageResourceExists, service.Name)
		c.recorder.Event(foo, corev1.EventTypeWarning, ErrResourceExists, msg)
		return service, fmt.Errorf(msg)
	}
	return service, nil
}

//func (c *RedisController) reconcile(redis *redis.Redis) error {

//	var masterIP string

//	if !redis.SeedMasterProcessComplete {
//		if err := c.seedMasterPodProcess(redis.Redis); err != nil {
//			return err
//		}

//		seedIP, err := util.GetSeedMasterIP(c.podLister, redis.Redis.Namespace, redis.Redis.Name)

//		if err != nil {
//			return err
//		}

//		masterIP = seedIP

//		redis.SeedMasterProcessComplete = true

//	} else {
//		ip, err := util.GetMasterIPByName(redis.Config.RedisClient, spec.GetRedisMasterName(redis.Redis))

//		if err != nil {
//			// Something went wrong, mark to spin up seed pod on next run
//			redis.SeedMasterProcessComplete = false
//			return err
//		}

//		masterIP = ip

//		deletePolicy := v1.DeletePropagationForeground

//		err = c.kubernetesClient.CoreV1().Pods(redis.Redis.Namespace).Delete(spec.GetMasterPodName(redis.Redis.Name),
//			&metav1.DeleteOptions{
//				PropagationPolicy: &deletePolicy,
//			})

//		if err != nil && !util.ResourceNotFoundError(err) {
//			redis.SeedMasterDeleted = false
//			return err
//		}

//		redis.SeedMasterDeleted = true
//	}

//	if err := c.masterEndpointProcess(redis.Redis, masterIP); err != nil {
//		return err
//	}

//	if err := c.masterServiceProcess(redis.Redis); err != nil {
//		return err
//	}

//	if err := c.sentinelConfigProcess(redis.Redis); err != nil {
//		return err
//	}

//	if err := c.sentinelServiceProcess(redis.Redis); err != nil {
//		return err
//	}

//	if err := c.sentinelProcess(redis.Redis); err != nil {
//		return err
//	}

//	if err := c.slaveProcess(redis.Redis); err != nil {
//		return err
//	}

//	return nil

//}
//func (c *RedisController) seedMasterPodProcess(redis *redisv1.Redis) error {

//	_, err := c.podLister.Pods(redis.Namespace).Get(spec.GetMasterPodName(redis.Name))

//	if err != nil {
//		if util.ResourceNotFoundError(err) {
//			if _, err := c.kubernetesClient.CoreV1().Pods(redis.Namespace).Create(spec.MasterSeedPod(redis)); err != nil {
//				return err
//			}

//			return nil
//		}

//		return err
//	}

//	return nil

//}

//func (c *RedisController) sentinelProcess(redis *redisv1.Redis) error {

//	actual, err := c.deploymentLister.Deployments(redis.Namespace).Get(spec.GetSentinelDeploymentName(redis.Name))

//	if err != nil {
//		if util.ResourceNotFoundError(err) {

//			return c.createDesiredDeployment(redis, spec.SentinelDeployment)
//		}

//		return err
//	}

//	obj, err := api.Scheme.DeepCopy(actual)

//	if err != nil {
//		return err
//	}

//	return c.updateDeploymentToDesired(redis, obj.(*appsv1beta1.Deployment), spec.SentinelDeployment)

//}

//func (c *RedisController) sentinelConfigProcess(redis *redisv1.Redis) error {

//	configMapName := string(redis.Spec.Sentinels.ConfigMap)

//	_, err := c.configMapLister.ConfigMaps(redis.Namespace).Get(configMapName)

//	if err != nil {
//		if util.ResourceNotFoundError(err) {

//			return c.createDesiredConfigMap(redis, spec.DefaultSentinelConfig)
//		}

//		return err
//	}

//	return nil
//}

//func (c *RedisController) slaveProcess(redis *redisv1.Redis) error {

//	actual, err := c.statefulSetLister.StatefulSets(redis.Namespace).Get(spec.GetSlaveStatefulSetName(redis.Name))

//	if err != nil {
//		if util.ResourceNotFoundError(err) {

//			return c.createDesiredStatefulSet(redis, spec.SlaveStatefulSet)
//		}

//		return err
//	}

//	obj, err := api.Scheme.DeepCopy(actual)

//	if err != nil {
//		return err
//	}

//	return c.updateStateFulSetToDesired(redis, obj.(*appsv1beta1.StatefulSet), spec.SlaveStatefulSet)

//}

//func (c *RedisController) sentinelServiceProcess(redis *redisv1.Redis) error {

//	actual, err := c.serviceLister.Services(redis.Namespace).Get(spec.GetSentinelServiceName(redis.Name))

//	if err != nil {
//		if util.ResourceNotFoundError(err) {

//			return c.createDesiredService(redis, spec.SentinelService)

//		}

//		return err
//	}

//	obj, err := api.Scheme.DeepCopy(actual)

//	if err != nil {
//		return err
//	}

//	return c.updateServiceToDesired(redis, obj.(*corev1.Service), spec.SentinelService)

//}

//func (c *RedisController) masterServiceProcess(redis *redisv1.Redis) error {

//	actual, err := c.serviceLister.Services(redis.Namespace).Get(spec.GetMasterServiceName(redis.Name))

//	if err != nil {
//		if util.ResourceNotFoundError(err) {

//			return c.createDesiredService(redis, spec.MasterService)
//		}

//		return err
//	}

//	obj, err := api.Scheme.DeepCopy(actual)

//	if err != nil {
//		return err
//	}

//	return c.updateServiceToDesired(redis, obj.(*corev1.Service), spec.MasterService)
//}

//func (c *RedisController) masterEndpointProcess(redis *redisv1.Redis, ipAddress string) error {

//	actual, err := c.endpointsLister.Endpoints(redis.Namespace).Get(spec.GetMasterServiceName(redis.Name))

//	if err != nil {
//		if util.ResourceNotFoundError(err) {

//			return c.createDesiredEndpoint(redis, ipAddress, spec.MasterServiceEndpoint)
//		}

//		return err
//	}

//	obj, err := api.Scheme.DeepCopy(actual)

//	if err != nil {
//		return err
//	}

//	return c.updateEndpointToDesired(redis, obj.(*corev1.Endpoints), ipAddress, spec.MasterServiceEndpoint)

//}

//func (c *RedisController) createDesiredDeployment(redis *redisv1.Redis, desiredFunction func(*redisv1.Redis) *appsv1beta1.Deployment) (err error) {

//	_, err = c.kubernetesClient.AppsV1beta1().Deployments(redis.Namespace).Create(desiredFunction(redis))

//	return

//}

//func (c *RedisController) createDesiredStatefulSet(redis *redisv1.Redis, desiredFunction func(*redisv1.Redis) *appsv1beta1.StatefulSet) (err error) {

//	_, err = c.kubernetesClient.AppsV1beta1().StatefulSets(redis.Namespace).Create(desiredFunction(redis))

//	return

//}

//func (c *RedisController) createDesiredConfigMap(redis *redisv1.Redis, desiredFunction func(*redisv1.Redis) *corev1.ConfigMap) (err error) {

//	_, err = c.kubernetesClient.CoreV1().ConfigMaps(redis.Namespace).Create(desiredFunction(redis))

//	return
//}

//func (c *RedisController) createDesiredService(redis *redisv1.Redis, desiredFunction func(*redisv1.Redis) *corev1.Service) (err error) {

//	_, err = c.kubernetesClient.CoreV1().Services(redis.Namespace).Create(desiredFunction(redis))

//	return

//}

//func (c *RedisController) createDesiredEndpoint(redis *redisv1.Redis, ipAddress string, desiredFunction func(*redisv1.Redis, string) *corev1.Endpoints) (err error) {

//	_, err = c.kubernetesClient.CoreV1().Endpoints(redis.Namespace).Create(desiredFunction(redis, ipAddress))

//	return

//}

//func (c *RedisController) updateDeploymentToDesired(redis *redisv1.Redis, actual *appsv1beta1.Deployment, desiredFunction func(*redisv1.Redis) *appsv1beta1.Deployment) error {

//	desired := desiredFunction(redis)

//	if !reflect.DeepEqual(actual, desired) {

//		actualJSON, err := json.Marshal(actual)
//		if err != nil {
//			return err
//		}
//		desiredJSON, err := json.Marshal(desired)
//		if err != nil {
//			return err
//		}

//		patch, err := strategicpatch.CreateTwoWayMergePatch(actualJSON, desiredJSON, appsv1beta1.Deployment{})

//		if err != nil {
//			return err
//		}

//		_, err = c.kubernetesClient.AppsV1beta1().Deployments(redis.Namespace).Patch(actual.Name, types.StrategicMergePatchType, patch)

//		return err
//	}

//	return nil
//}

//func (c *RedisController) updateStateFulSetToDesired(redis *redisv1.Redis, actual *appsv1beta1.StatefulSet, desiredFunction func(*redisv1.Redis) *appsv1beta1.StatefulSet) error {

//	desired := desiredFunction(redis)

//	if !reflect.DeepEqual(actual, desired) {

//		actualJSON, err := json.Marshal(actual)
//		if err != nil {
//			return err
//		}
//		desiredJSON, err := json.Marshal(desired)
//		if err != nil {
//			return err
//		}

//		patch, err := strategicpatch.CreateTwoWayMergePatch(actualJSON, desiredJSON, appsv1beta1.StatefulSet{})

//		if err != nil {
//			return err
//		}

//		_, err = c.kubernetesClient.AppsV1beta1().StatefulSets(redis.Namespace).Patch(actual.Name, types.StrategicMergePatchType, patch)

//		return err
//	}

//	return nil
//}

//func (c *RedisController) updateServiceToDesired(redis *redisv1.Redis, actual *corev1.Service, desiredFunction func(*redisv1.Redis) *corev1.Service) error {

//	desired := desiredFunction(redis)

//	if !reflect.DeepEqual(actual, desired) {

//		desiredJSON, err := json.Marshal(desired)
//		if err != nil {
//			return err
//		}

//		_, err = c.kubernetesClient.CoreV1().Services(redis.Namespace).Patch(actual.Name, types.StrategicMergePatchType, desiredJSON)

//		return err
//	}

//	return nil
//}

//func (c *RedisController) updateConfigMapToDesired(redis *redisv1.Redis, actual *corev1.ConfigMap, desiredFunction func(*redisv1.Redis) *corev1.ConfigMap) error {

//	desired := desiredFunction(redis)

//	if !reflect.DeepEqual(actual, desired) {
//		desiredJSON, err := json.Marshal(desired)
//		if err != nil {
//			return err
//		}

//		_, err = c.kubernetesClient.CoreV1().ConfigMaps(redis.Namespace).Patch(actual.Name, types.StrategicMergePatchType, desiredJSON)

//		return err
//	}

//	return nil
//}

//func (c *RedisController) updateEndpointToDesired(redis *redisv1.Redis, actual *corev1.Endpoints, ipAddress string, desiredFunction func(*redisv1.Redis, string) *corev1.Endpoints) error {

//	desired := desiredFunction(redis, ipAddress)

//	if !reflect.DeepEqual(actual, desired) {
//		desiredJSON, err := json.Marshal(desired)
//		if err != nil {
//			return err
//		}

//		_, err = c.kubernetesClient.CoreV1().Endpoints(redis.Namespace).Patch(actual.Name, types.StrategicMergePatchType, desiredJSON)

//		return err
//	}

//	return nil
//}

//func (c *RedisController) deleteResources(namespace, name string) error {

//	deletePolicy := metav1.DeletePropagationBackground

//	err := c.kubernetesClient.CoreV1().Pods(namespace).Delete(spec.GetMasterPodName(name), &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})

//	if err != nil && !util.ResourceNotFoundError(err) {
//		return err
//	}

//	err = c.kubernetesClient.AppsV1beta1().Deployments(namespace).Delete(spec.GetSlaveStatefulSetName(name), &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})

//	if err != nil && !util.ResourceNotFoundError(err) {
//		return err
//	}

//	err = c.kubernetesClient.AppsV1beta1().Deployments(namespace).Delete(spec.GetSentinelDeploymentName(name), &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})

//	if err != nil && !util.ResourceNotFoundError(err) {
//		return err
//	}

//	err = c.kubernetesClient.CoreV1().Services(namespace).Delete(spec.GetSentinelServiceName(name), &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})

//	if err != nil && !util.ResourceNotFoundError(err) {
//		return err
//	}

//	err = c.kubernetesClient.CoreV1().ConfigMaps(namespace).Delete(spec.GetSentinelConfigMapName(name), &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})

//	if err != nil && !util.ResourceNotFoundError(err) {
//		return err
//	}

//	err = c.kubernetesClient.AppsV1beta1().StatefulSets(namespace).Delete(spec.GetSlaveStatefulSetName(name), &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})

//	if err != nil && !util.ResourceNotFoundError(err) {
//		return err
//	}

//	err = c.kubernetesClient.CoreV1().Services(namespace).Delete(spec.GetMasterServiceName(name), &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})

//	if err != nil && !util.ResourceNotFoundError(err) {
//		return err
//	}

//	return nil
//}
