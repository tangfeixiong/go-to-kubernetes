package server

import (
	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/preboot"
)

type Config struct {
	SecureAddress   string
	InsecureAddress string
	SecureHTTP      bool
	LogLevel        int
	PrebootCfg      preboot.Config
}
