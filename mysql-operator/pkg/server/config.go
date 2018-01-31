package server

import (
	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/initcnf"
)

type Config struct {
	SecureAddress   string
	InsecureAddress string
	SecureHTTP      bool
	LogLevel        int
	InitConfig      initcnf.Config
}
