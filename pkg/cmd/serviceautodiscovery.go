package cmd

import (
	"flag"
	"fmt"
	"io"
	"log"
	//"os"
	"strconv"

	"github.com/spf13/cobra"

	//"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/kubectl"
	//kubectlcmd "k8s.io/kubernetes/pkg/kubectl/cmd"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	//"k8s.io/kubernetes/pkg/kubectl/resource"
	//"k8s.io/kubernetes/pkg/runtime"
	//utilerrors "k8s.io/kubernetes/pkg/util/errors"
	//"k8s.io/kubernetes/pkg/watch"

	"github.com/tangfeixiong/go-to-kubernetes/pkg/component"
)

// ServiceAutoDiscoveryOptions is the start of the data required to perform the operation.  As new fields are added, add them here instead of
// referencing the cmd.Flags()
type ServiceAutoDiscoveryOptions struct {
	Filenames []string
	Recursive bool

	Raw string
}

const (
	get_sad_long = `Display one or many service auto discovery information.

` + kubectl.PossibleResourceTypes + `

By specifying the output as 'template' and providing a Go template as the value
of the --template flag, you can filter the attributes of the fetched resource(s).`
	get_sad_example = `# List all pods in ps output format.
kubectl get pods

# List all pods in ps output format with more information (such as node name).
kubectl get pods -o wide

# List a single replication controller with specified NAME in ps output format.
kubectl get replicationcontroller web

# List a single pod in JSON output format.
kubectl get -o json pod web-pod-13je7

# List a pod identified by type and name specified in "pod.yaml" in JSON output format.
kubectl get -f pod.yaml -o json

# Return only the phase value of the specified pod.
kubectl get -o template pod/web-pod-13je7 --template={{.status.phase}}

# List all replication controllers and services together in ps output format.
kubectl get rc,services

# List one or more resources by their type and names.
kubectl get rc/web service/frontend pods/web-pod-13je7`
)

// NewCmdServiceAutoDiscovery creates a command object for the generic "get" action, which
// retrieves one or more resources from a server.
func NewCmdServiceAutoDiscovery(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	options := &ServiceAutoDiscoveryOptions{}

	// retrieve a list of handled resources from printer as valid args
	validArgs, argAliases := []string{}, []string{}
	//p, err := f.Printer(nil, false, false, false, false, false, false, []string{})
	p, err := f.Printer(nil, &kubectl.PrintOptions{})
	cmdutil.CheckErr(err)
	if p != nil {
		validArgs = p.HandledResources()
		argAliases = kubectl.ResourceAliases(validArgs)
	}

	cmd := &cobra.Command{
		Use:     "get-service-auto-discovery [(-o|--output=)json|yaml|wide|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=...] (TYPE [NAME | -l label] | TYPE/NAME ...) [flags]",
		Short:   "Display one or many resources",
		Long:    get_sad_long,
		Example: get_sad_example,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunServiceAutoDiscovery(f, out, cmd, args, options)
			cmdutil.CheckErr(err)
		},
		SuggestFor: []string{"list", "ps"},
		ValidArgs:  validArgs,
		ArgAliases: argAliases,
	}
	cmdutil.AddPrinterFlags(cmd)
	cmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on")
	cmd.Flags().BoolP("watch", "w", false, "After listing/getting the requested object, watch for changes.")
	cmd.Flags().Bool("watch-only", false, "Watch for changes to the requested object(s), without listing/getting first.")
	cmd.Flags().Bool("all-namespaces", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	cmd.Flags().Bool("show-kind", false, "If present, list the kind of each requested resource.")
	cmd.Flags().StringSliceP("label-columns", "L", []string{}, "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag statements like -L label1 -L label2...")
	cmd.Flags().Bool("export", false, "If true, use 'export' for the resources.  Exported resources are stripped of cluster-specific information.")
	usage := "Filename, directory, or URL to a file identifying the resource to get from a server."
	kubectl.AddJsonFilenameFlag(cmd, &options.Filenames, usage)
	cmdutil.AddRecursiveFlag(cmd, &options.Recursive)
	cmdutil.AddInclude3rdPartyFlags(cmd)
	cmd.Flags().StringVar(&options.Raw, "raw", options.Raw, "Raw URI to request from the server.  Uses the transport specified by the kubeconfig file.")
	return cmd
}

func RunServiceAutoDiscovery(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, args []string, options *ServiceAutoDiscoveryOptions) error {
	pp := "[component/RunServiceAutoDiscovery]"
	var logger *log.Logger = log.New(out, pp+" ", log.LstdFlags|log.Lshortfile)

	//selector := cmdutil.GetFlagString(cmd, "selector")
	allNamespaces := cmdutil.GetFlagBool(cmd, "all-namespaces")
	//mapper, typer := f.Object(cmdutil.GetIncludeThirdPartyAPIs(cmd))

	cmdNamespace, enforceNamespace, err := f.DefaultNamespace()
	if err != nil {
		fmt.Fprintf(out, "%+s Namespace configuration Failed: %+v", pp, err)
		return err
	}

	if allNamespaces {
		enforceNamespace = false
	}

	if false == flag.Parsed() {
		flag.Parse()
	}
	if llf := flag.Lookup("v"); llf != nil {
		if v, err := strconv.Atoi(llf.Value.String()); err == nil && v > 2 {
			logger.Printf("enforce namespace: %+v", enforceNamespace)
		}
	}

	sad, err := component.ServiceAutoDiscoveryEnvVar(f, cmdNamespace)
	if err != nil {
		fmt.Fprintf(out, "%+s Execution Failed: %+v", pp, err)
		return err
	}

	for i := range sad {
		env := sad[i].Env
		fmt.Fprintf(out, "# Auto discovery %d\n# Service host and first port\n", i+1)
		fmt.Fprintf(out, "%s=%s\n", env.ServiceHostEnvVar.Name, env.ServiceHostEnvVar.Value)
		fmt.Fprintf(out, "%s=%s\n", env.FirstPortEnvVar.Name, env.FirstPortEnvVar.Value)
		if len(env.ServicePortsEnvVar) > 0 {
			fmt.Fprint(out, "\n#    Service ports\n")
			for j := range env.ServicePortsEnvVar {
				fmt.Fprintf(out, "%s=%s\n", env.ServicePortsEnvVar[j].Name, env.ServicePortsEnvVar[j].Value)
			}
		}
		if len(env.HostnameEnvVar) > 0 {
			fmt.Fprint(out, "\n#    PODs HOSTNAME\n")
			for j := range env.HostnameEnvVar {
				fmt.Fprintf(out, "# %s=%s\n", env.HostnameEnvVar[j].Name, env.HostnameEnvVar[j].Value)
			}
		}
		fmt.Fprint(out, "\n#    Docker compatial EnvVar\n")
		fmt.Fprintf(out, "%s=%s\n", env.DockerFirstServiceEnvVar.Name, env.DockerFirstServiceEnvVar.Value)
		for j := 0; j < len(env.DockerServicesEnvVar); j += 4 {
			fmt.Fprintf(out, "\n#    Docker compatial details %d\n", j/4+1)
			fmt.Fprintf(out, "%s=%s\n", env.DockerServicesEnvVar[j].Name, env.DockerServicesEnvVar[j].Value)
			fmt.Fprintf(out, "%s=%s\n", env.DockerServicesEnvVar[j+1].Name, env.DockerServicesEnvVar[j+1].Value)
			fmt.Fprintf(out, "%s=%s\n", env.DockerServicesEnvVar[j+2].Name, env.DockerServicesEnvVar[j+2].Value)
			fmt.Fprintf(out, "%s=%s\n", env.DockerServicesEnvVar[j+3].Name, env.DockerServicesEnvVar[j+3].Value)
		}
		if i+1 < len(sad) {
			fmt.Fprintln(out)
		}
	}
	return nil
}
