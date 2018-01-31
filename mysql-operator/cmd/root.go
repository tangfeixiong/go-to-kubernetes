package cmd

import (
	"flag"
	"path/filepath"
	"strconv"
	// "fmt"
	// "log"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/initcnf"
	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/server"
)

func RootCommandFor(name string) *cobra.Command {
	var config server.Config
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

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
	root.AddCommand(createInitCommand(&config.InitConfig))

	cfg := &config.InitConfig
	root.PersistentFlags().StringVar(&cfg.Kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file. it means running out of cluster if supplied")
	if home := HomeDir(); home != "" {
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
			//server.Start(config)
		},
	}

	command.Flags().StringVar(&config.SecureAddress, "grpc_addr", "0.0.0.0:10001", "IP:port format")
	command.Flags().StringVar(&config.InsecureAddress, "http_addr", "0.0.0.0:10002", "IP:port format. Serve HTTP, or No HTTP if empty")
	command.Flags().BoolVar(&config.SecureHTTP, "secure_http", false, "Currently not used, if both HTTP address and HTTPS flag not set, just gRPC noly")
	command.Flags().IntVar(&config.LogLevel, "log_level", 2, "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)

	return command
}

/*
  Inspired by:
    https://github.com/kubernetes/client-go/blob/master/util/homedir/homedir.go
*/
// HomeDir returns the home directory for the current user
func HomeDir() string {
	if runtime.GOOS == "windows" {

		// First prefer the HOME environmental variable
		if home := os.Getenv("HOME"); len(home) > 0 {
			if _, err := os.Stat(home); err == nil {
				return home
			}
		}
		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			homeDir := homeDrive + homePath
			if _, err := os.Stat(homeDir); err == nil {
				return homeDir
			}
		}
		if userProfile := os.Getenv("USERPROFILE"); len(userProfile) > 0 {
			if _, err := os.Stat(userProfile); err == nil {
				return userProfile
			}
		}
	}
	return os.Getenv("HOME")
}

func createInitCommand(config *initcnf.Config) *cobra.Command {

	command := &cobra.Command{
		Use:   "init",
		Short: "init galera.cnf",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", strconv.Itoa(config.LogLevel))
			flag.Parse()

			config.Clustering = initcnf.MariadbGalera
			initcnf.BuildGaleraCnf(config)
		},
	}

	command.PersistentFlags().StringVar(&config.Name, "name", "", "StatefulSet name, or lookup value via label <crd group>/go-to-kubernetes")
	command.PersistentFlags().StringVar(&config.Namespace, "namespace", "", "Kubernetes namespace, or lookup value from env name POD_NAMESPACE, otherwise 'default'")
	command.PersistentFlags().StringVar(&config.Dir, "conf_dir", "/etc/mysql/conf.d", "Directory of galera.cnf")
	command.PersistentFlags().IntVar(&config.Port, "service_port", 3306, "Service port")
	return command
}
