package cmd

import (
	"flag"
	"path/filepath"
	"strconv"
	// "fmt"
	// "log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-kubernetes/pkg/util"
	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/rabbitmq"
	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/server"
)

func RootCommandFor(name string) *cobra.Command {
	var config server.Config
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr
	cfg := &config.RabbitmqConfig

	root := &cobra.Command{
		Use:   name,
		Short: "operator server, with gRPC & ReST API",
		Long: `
        mysql-operator
        
        This is a ..., and ...
        ...,
        It is inspired by some github projects.
        `,
	}
	root.AddCommand(createServiceCommand(&config))
	root.AddCommand(createInitCommand(cfg))

	root.PersistentFlags().StringVar(&cfg.Kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file. it means running out of cluster if supplied")
	if home := util.HomeDir(); home != "" {
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

func createInitCommand(config *rabbitmq.Config) *cobra.Command {

	command := &cobra.Command{
		Use:   "init",
		Short: "init configurtions of /etc/rabbitmq/",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", strconv.Itoa(config.LogLevel))
			flag.Parse()

			rabbitmq.BuildConfigFile(config)
		},
	}

	command.PersistentFlags().StringVar(&config.Name, "name", "", "StatefulSet name, or lookup value via label <crd group>/go-to-kubernetes")
	command.PersistentFlags().StringVar(&config.Namespace, "namespace", "", "Kubernetes namespace, or lookup value from env name POD_NAMESPACE, otherwise 'default'")
	command.PersistentFlags().StringVar(&config.DomainName, "domain_name", "cluster.local", "Domain name of K8s DNS")
	command.PersistentFlags().StringVar(&config.Mode, "mode", "", "Peer_Discovery_K8s, Or native cluster configuration")
	command.PersistentFlags().StringVar(&config.Dir, "conf_dir", "/etc/rabbitmq", "Directory of galera.cnf")
	command.PersistentFlags().IntVar(&config.Port, "service_port", 5672, "Service port")
	return command
}
