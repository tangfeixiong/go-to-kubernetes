package cmd

import (
	"flag"
	"path/filepath"
	"strconv"
	// "fmt"
	// "log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/operator"
	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/server"
)

func RootCommandFor(name string) *cobra.Command {
	var config server.Config
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "operator server, with gRPC & ReST API",
		Long: `
        redis-operator
        
        This is a ..., and ...
        ...,
        It is inspired by some github projects.
        `,
	}
	root.AddCommand(createServiceCommand(&config))
	root.AddCommand(createRedisHighAvailabilityConfigCommand(&config.RedisBootstrapConfig))

	bc := &config.RedisBootstrapConfig
	root.PersistentFlags().StringVar(&bc.Kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file. it means running out of cluster if supplied")
	if home := homeDir(); home != "" {
		root.PersistentFlags().Lookup("kubeconfig").NoOptDefVal = filepath.Join(home, ".kube", "config")
	}
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return root
}

func createServiceCommand(config *server.Config) *cobra.Command {

	command := &cobra.Command{
		Use:   "serve",
		Short: "Serving with gRPC and a gRPC-Gateway",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", strconv.Itoa(config.LogLevel))
			flag.Parse()
			server.Start(config)
		},
	}

	command.Flags().StringVar(&config.SecureAddress, "grpc_addr", "0.0.0.0:10001", "IP:port format")
	command.Flags().StringVar(&config.InsecureAddress, "http_addr", "0.0.0.0:10002", "IP:port format. Serve HTTP, or No HTTP if empty")
	command.Flags().BoolVar(&config.SecureHTTP, "secure_http", false, "Currently not used, if both HTTP address and HTTPS flag not set, just gRPC noly")
	command.Flags().IntVar(&config.LogLevel, "log_level", 2, "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)

	return command
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func createRedisHighAvailabilityConfigCommand(config *operator.RedisBootstrapConfig) *cobra.Command {

	command := &cobra.Command{
		Use:   "config-ha",
		Short: "config redis.conf or sentinel.conf",
	}
	command.AddCommand(createRedisCommand(config), createSentinelCommand(config))

	command.PersistentFlags().StringVar(&config.Name, "name", "my-redis", "Artifact name")
	command.PersistentFlags().StringVar(&config.Namespace, "namespace", "default", "Application namespace")
	command.PersistentFlags().StringVar(&config.Dir, "conf_dir", "/data", "Directory for redis.conf or sentinel.conf")
	command.PersistentFlags().IntVar(&config.Port, "service_port", 6379, "Service port")
	return command
}

func createRedisCommand(config *operator.RedisBootstrapConfig) *cobra.Command {

	command := &cobra.Command{
		Use:   "redis",
		Short: "config redis.conf",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", strconv.Itoa(config.LogLevel))
			flag.Parse()
			config.Role = operator.RedisHighAvailabilityRedis
			operator.ConfigConf(config)
		},
	}
	command.Flags().StringVar(&config.Kind, "kind", "StatefulSet", "Maybe support Deployment, ReplicationController")

	return command
}

func createSentinelCommand(config *operator.RedisBootstrapConfig) *cobra.Command {

	command := &cobra.Command{
		Use:   "sentinel",
		Short: "config sentinel.conf",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", strconv.Itoa(config.LogLevel))
			flag.Parse()
			config.Role = operator.RedisHighAvailabilitySentinel
			operator.ConfigConf(config)
		},
	}
	command.Flags().StringVar(&config.Kind, "kind", "Deployment", "Maybe support ReplicationController")

	return command
}
