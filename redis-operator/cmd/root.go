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

	"github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/server"
)

func RootCommandFor(name string) *cobra.Command {
	var config server.Config
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "operator server, with gRPC & ReST API",
		Long: `
        hadoop-hdfs-operator
        
        This is a ..., and ...
        ...,
        It is inspired by some github projects.
        `,
	}
	root.AddCommand(createServiceCommand(&config))
	// root.AddCommand(createClientCommand())

	root.Flags().StringVar(&config.Kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file, default is $HOME/.kube/config")
	if home := homeDir(); home != "" {
		root.Flags().Lookup("kubeconfig").NoOptDefVal = filepath.Join(home, ".kube", "config")
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

	command.Flags().StringVar(&config.SecureAddress, "grpc-addr", "0.0.0.0:10001", "IP:port format")
	command.Flags().StringVar(&config.InsecureAddress, "http-addr", "0.0.0.0:10002", "IP:port format. Serve HTTP, or No HTTP if empty")
	command.Flags().BoolVar(&config.SecureHTTP, "secure-http", false, "Currently not used, if both HTTP address and HTTPS flag not set, just gRPC noly")
	command.Flags().IntVar(&config.LogLevel, "log-level", 2, "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)

	return command
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
