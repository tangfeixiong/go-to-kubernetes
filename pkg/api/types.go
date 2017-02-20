package api

import (
	kapi "k8s.io/kubernetes/pkg/api"
)

type ServiceAutoDiscovery struct {
	Cluster   string
	Domain    string
	Namespace string
	Env       ServiceAutoDiscoveryEnv
	AddonDNS  string
}

type DockerCompatialEnv struct {
	DockerFirstServiceEnvVar kapi.EnvVar
	DockerServicesEnvVar     []kapi.EnvVar
}

type ServiceAutoDiscoveryEnv struct {
	ServiceHostEnvVar  kapi.EnvVar
	FirstPortEnvVar    kapi.EnvVar
	ServicePortsEnvVar []kapi.EnvVar
	HostnameEnvVar     []kapi.EnvVar
	DockerCompatialEnv
}
