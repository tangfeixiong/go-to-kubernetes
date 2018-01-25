package svc

import (
	apiv1 "k8s.io/api/core/v1"
)

const (
	DEFAULT_NAMESPACE = "default"
)

var (
	name_default string = "my-redis"
)

type ServiceRecipient interface {
	GenerateArtifact() (*apiv1.Service, error)
}

func findServicePort(svc *apiv1.Service, name string) (bool, *apiv1.ServicePort) {
	var sp *apiv1.ServicePort

	for i := 0; i < len(svc.Spec.Ports); i++ {
		if svc.Spec.Ports[i].Name == name {
			sp = &svc.Spec.Ports[i]
			return true, sp
		}
	}
	return false, nil
}
