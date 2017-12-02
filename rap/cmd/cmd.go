package cmd

import (
	"flag"
	"strconv"
	// "fmt"
	// "log"
	// "os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-kubernetes/rap/pkg/server"
)

func RootCommandFor(name string) *cobra.Command {
	var config server.Config
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "Remote desktop and shell web service, for example noVNC, hyperterm ",
		Long: `
        rap
        
        This is a ..., and ...
        ...,
        It is inspired by some github projects.
        `,
	}
	root.AddCommand(createServiceCommand(&config))
	// root.AddCommand(createClientCommand())

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

	command.Flags().StringVar(&config.SecureAddress, "grpc-addr", "0.0.0.0:10071", "Default is 0.0.0.0:10071")
	command.Flags().StringVar(&config.InsecureAddress, "http-addr", "0.0.0.0:10072", "Serve HTTP, or No HTTP if empty")
	command.Flags().BoolVar(&config.SecureHTTP, "secure-http", false, "Default not used, if both HTTP address and HTTPS flag not set, just gRPC noly")
	command.Flags().IntVar(&config.LogLevel, "log-level", 2, "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)

	return command
}
