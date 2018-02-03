package server

import (
	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/rabbitmq"
)

type Config struct {
	SecureAddress   string
	InsecureAddress string
	SecureHTTP      bool
	LogLevel        int
	RabbitmqConfig  rabbitmq.Config
}
