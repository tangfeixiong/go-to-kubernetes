package operator

import (
	"fmt"

	//"github.com/go-redis/redis"
	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/spec/sts"
)

type redisControllerManager struct {
}

func (ctlman *redisControllerManager) ProvisionMasterSlaves(name, ns, provisioning, svcName string, slaves int) error {
	glog.V(2).Infoln("provisioning StatefulSet for redis")

	recipe := sts.NewRedisRecipient(name, ns, provisioning, svcName, slaves)

	s, err := recipe.GenerateArtifact()
	if err != nil {
		return err
	}

	clientset, err := K8sClientset("")
	if err != nil {
		glog.V(2).Infoln("Get kubernetes clientset failed:", err.Error())
		return fmt.Errorf("Cloud not get Kubernetes clientset, error:", err.Error())
	}

	stsClient := clientset.AppsV1beta2().StatefulSets(ns)

	result, err := stsClient.Create(s)
	if err != nil {
		glog.Errorf("Create StatefulSet failed: %v", err)
		return err
	}
	fmt.Printf("Created statefulset %q.\n", result.GetObjectMeta().GetName())
	return nil
}
