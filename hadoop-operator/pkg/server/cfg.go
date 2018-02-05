package server

import (
	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/hadoop"
)

type Config struct {
	SecureAddress   string
	InsecureAddress string
	SecureHTTP      bool
	LogLevel        int
	RuntimeConfig   hadoop.Config
}
