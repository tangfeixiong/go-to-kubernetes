package cmd

import (
	"fmt"
	"io"
	"os"

	//"github.com/golang/glog"
	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/helm"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/kubectl"
	kubectlcmd "k8s.io/kubernetes/pkg/kubectl/cmd"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/resource"
)

const (
	create_long = `Create a resource by filename or stdin.

JSON and YAML formats are accepted.`
	create_example = `# Create a pod using the data in pod.json.
kubectl create -f ./pod.json

# Create a pod based on the JSON passed into stdin.
cat pod.json | kubectl create -f -`
)

type CreateFromChartOptions struct {
	Filenames []string
	Recursive bool
}

func NewCmdCreateFromChart(facto *cmdutil.Factory, out io.Writer, c helm.Interface) *cobra.Command {
	options := &CreateFromChartOptions{}
	inst := &installCmd{
		out:    out,
		client: c,
		values: new(values),
	}

	cmd := &cobra.Command{
		Use:     "create-from-chart [CHART]",
		Short:   "Create a resource by filename or stdin",
		Long:    create_long,
		Example: create_example,
		Run: func(cmd *cobra.Command, args []string) {
			/*if len(options.Filenames) == 0 {
				cmd.Help()
				return
			}
			cmdutil.CheckErr(ValidateArgs(cmd, args))
			cmdutil.CheckErr(cmdutil.ValidateOutputArgs(cmd))
			cmdutil.CheckErr(RunCreateFromChart(facto, cmd, out, options))*/

			// k8s.io/helm/cmd/helm/install.go
			//
			if err := checkArgsLength(len(args), "chart name"); err != nil {
				//return err
				cmdutil.CheckErr(err)
			}
			cp, err := locateChartPath(args[0], inst.verify, inst.keyring)
			if err != nil {
				//return err
				cmdutil.CheckErr(err)
			}
			inst.chartPath = cp
			inst.client = ensureHelmClient(inst.client)
			//return inst.run()
			cmdutil.CheckErr(inst.run())
		},
		SilenceUsage:      true,
		PersistentPreRunE: setupConnection,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			teardown()
		},
	}

	// k8s.io/helm/cmd/helm/helm.go
	//
	home := os.Getenv(homeEnvVar)
	if home == "" {
		home = "$HOME/.helm"
	}
	thost := os.Getenv(hostEnvVar)
	p := cmd.PersistentFlags()
	p.StringVar(&helmHome, "home", home, "location of your Helm config. Overrides $HELM_HOME.")
	p.StringVar(&tillerHost, "host", thost, "address of tiller. Overrides $HELM_HOST.")
	p.BoolVarP(&flagDebug, "debug", "", false, "enable verbose output")

	// k8s.io/helm/cmd/helm/install.go
	//
	f := cmd.Flags()
	f.StringVarP(&inst.valuesFile, "values", "", "", "specify values in a YAML file")
	f.StringVarP(&inst.name, "name", "n", "", "the release name. If unspecified, it will autogenerate one for you")
	// TODO use kubeconfig default
	f.StringVar(&inst.namespace, "namespace", "default", "the namespace to install the release into")
	f.BoolVar(&inst.dryRun, "dry-run", false, "simulate an install")
	f.BoolVar(&inst.disableHooks, "no-hooks", false, "prevent hooks from running during install")
	f.BoolVar(&inst.replace, "replace", false, "re-use the given name, even if that name is already used. This is unsafe in production")
	f.Var(inst.values, "set", "set values on the command line. Separate values with commas: key1=val1,key2=val2")
	f.StringVar(&inst.nameTemplate, "name-template", "", "specify template used to name the release")
	f.BoolVar(&inst.verify, "verify", false, "verify the package before installing it")
	f.StringVar(&inst.keyring, "keyring", defaultKeyring(), "location of public keys used for verification")

	// k8s.io/kubernetes/pkg/kubectl/cmd/create.go
	//
	usage := "Filename, directory, or URL to file to use to create the resource"
	kubectl.AddJsonFilenameFlag(cmd, &options.Filenames, usage)
	cmd.MarkFlagRequired("filename")
	cmdutil.AddValidateFlags(cmd)
	cmdutil.AddRecursiveFlag(cmd, &options.Recursive)
	cmdutil.AddOutputFlagsForMutation(cmd)
	cmdutil.AddApplyAnnotationFlags(cmd)
	cmdutil.AddRecordFlag(cmd)
	cmdutil.AddInclude3rdPartyFlags(cmd)

	return cmd
}

func RunCreateFromChart(f *cmdutil.Factory, cmd *cobra.Command, out io.Writer, options *CreateFromChartOptions) error {

	return nil
}

// k8s.io/kubernetes/pkg/kubectl/cmd/create.go
//
func ValidateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return cmdutil.UsageError(cmd, "Unexpected args: %v", args)
	}
	return nil
}

func RunCreate(f *cmdutil.Factory, cmd *cobra.Command, out io.Writer, options *kubectlcmd.CreateOptions) error {
	schema, err := f.Validator(cmdutil.GetFlagBool(cmd, "validate"), cmdutil.GetFlagString(cmd, "schema-cache-dir"))
	if err != nil {
		return err
	}

	cmdNamespace, enforceNamespace, err := f.DefaultNamespace()
	if err != nil {
		return err
	}

	mapper, typer := f.Object(cmdutil.GetIncludeThirdPartyAPIs(cmd))
	r := resource.NewBuilder(mapper, typer, resource.ClientMapperFunc(f.ClientForMapping), f.Decoder(true)).
		Schema(schema).
		ContinueOnError().
		NamespaceParam(cmdNamespace).DefaultNamespace().
		FilenameParam(enforceNamespace, options.Recursive, options.Filenames...).
		Flatten().
		Do()
	err = r.Err()
	if err != nil {
		return err
	}

	count := 0
	err = r.Visit(func(info *resource.Info, err error) error {
		if err != nil {
			return err
		}
		if err := kubectl.CreateOrUpdateAnnotation(cmdutil.GetFlagBool(cmd, cmdutil.ApplyAnnotationsFlag), info, f.JSONEncoder()); err != nil {
			return cmdutil.AddSourceToErr("creating", info.Source, err)
		}

		if cmdutil.ShouldRecord(cmd, info) {
			if err := cmdutil.RecordChangeCause(info.Object, f.Command()); err != nil {
				return cmdutil.AddSourceToErr("creating", info.Source, err)
			}
		}

		if err := createAndRefresh(info); err != nil {
			return cmdutil.AddSourceToErr("creating", info.Source, err)
		}

		count++
		shortOutput := cmdutil.GetFlagString(cmd, "output") == "name"
		if !shortOutput {
			f.PrintObjectSpecificMessage(info.Object, out)
		}
		cmdutil.PrintSuccess(mapper, shortOutput, out, info.Mapping.Resource, info.Name, "created")
		return nil
	})
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no objects passed to create")
	}
	return nil
}

// createAndRefresh creates an object from input info and refreshes info with that object
func createAndRefresh(info *resource.Info) error {
	obj, err := resource.NewHelper(info.Client, info.Mapping).Create(info.Namespace, true, info.Object)
	if err != nil {
		return err
	}
	info.Refresh(obj, true)
	return nil
}

// NameFromCommandArgs is a utility function for commands that assume the first argument is a resource name
func NameFromCommandArgs(cmd *cobra.Command, args []string) (string, error) {
	if len(args) == 0 {
		return "", cmdutil.UsageError(cmd, "NAME is required")
	}
	return args[0], nil
}

// CreateSubcommandOptions is an options struct to support create subcommands
type CreateSubcommandOptions struct {
	// Name of resource being created
	Name string
	// StructuredGenerator is the resource generator for the object being created
	StructuredGenerator kubectl.StructuredGenerator
	// DryRun is true if the command should be simulated but not run against the server
	DryRun bool
	// OutputFormat
	OutputFormat string
}

// RunCreateSubcommand executes a create subcommand using the specified options
func RunCreateSubcommand(f *cmdutil.Factory, cmd *cobra.Command, out io.Writer, options *CreateSubcommandOptions) error {
	namespace, _, err := f.DefaultNamespace()
	if err != nil {
		return err
	}
	obj, err := options.StructuredGenerator.StructuredGenerate()
	if err != nil {
		return err
	}
	mapper, typer := f.Object(cmdutil.GetIncludeThirdPartyAPIs(cmd))
	gvks, _, err := typer.ObjectKinds(obj)
	if err != nil {
		return err
	}
	gvk := gvks[0]
	mapping, err := mapper.RESTMapping(unversioned.GroupKind{Group: gvk.Group, Kind: gvk.Kind}, gvk.Version)
	if err != nil {
		return err
	}
	client, err := f.ClientForMapping(mapping)
	if err != nil {
		return err
	}
	resourceMapper := &resource.Mapper{
		ObjectTyper:  typer,
		RESTMapper:   mapper,
		ClientMapper: resource.ClientMapperFunc(f.ClientForMapping),
	}
	info, err := resourceMapper.InfoForObject(obj, nil)
	if err != nil {
		return err
	}
	if err := kubectl.UpdateApplyAnnotation(info, f.JSONEncoder()); err != nil {
		return err
	}
	if !options.DryRun {
		obj, err = resource.NewHelper(client, mapping).Create(namespace, false, info.Object)
		if err != nil {
			return err
		}
	}

	if useShortOutput := options.OutputFormat == "name"; useShortOutput || len(options.OutputFormat) == 0 {
		cmdutil.PrintSuccess(mapper, useShortOutput, out, mapping.Resource, options.Name, "created")
		return nil
	}

	return f.PrintObject(cmd, mapper, obj, out)
}
