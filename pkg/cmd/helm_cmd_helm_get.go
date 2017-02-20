package cmd

// k8s.io/helm/cmd/helm/get.go
//
import (
	"errors"
	"io"
	"text/template"
	"time"

	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/timeconv"
)

var getHelp = `
This command shows the details of a named release.

It can be used to get extended information about the release, including:

  - The values used to generate the release
  - The chart used to generate the release
  - The generated manifest file

By default, this prints a human readable collection of information about the
chart, the supplied values, and the generated manifest file.
`

var errReleaseRequired = errors.New("release name is required")

type getCmd struct {
	release string
	out     io.Writer
	client  helm.Interface
	version int32
}

func newGetCmd(client helm.Interface, out io.Writer) *cobra.Command {
	get := &getCmd{
		out:    out,
		client: client,
	}
	cmd := &cobra.Command{
		Use:               "get [flags] RELEASE_NAME",
		Short:             "download a named release",
		Long:              getHelp,
		PersistentPreRunE: setupConnection,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errReleaseRequired
			}
			get.release = args[0]
			if get.client == nil {
				get.client = helm.NewClient(helm.Host(tillerHost))
			}
			return get.run()
		},
	}

	cmd.PersistentFlags().Int32Var(&get.version, "version", 0, "get the named release with version")

	cmd.AddCommand(newGetValuesCmd(nil, out))
	cmd.AddCommand(newGetManifestCmd(nil, out))
	cmd.AddCommand(newGetHooksCmd(nil, out))
	return cmd
}

var getTemplate = `VERSION: {{.Release.Version}}
RELEASED: {{.ReleaseDate}}
CHART: {{.Release.Chart.Metadata.Name}}-{{.Release.Chart.Metadata.Version}}
USER-SUPPLIED VALUES:
{{.Release.Config.Raw}}
COMPUTED VALUES:
{{.ComputedValues}}
HOOKS:
{{- range .Release.Hooks }}
---
# {{.Name}}
{{.Manifest}}
{{- end }}
MANIFEST:
{{.Release.Manifest}}
`

// getCmd is the command that implements 'helm get'
func (g *getCmd) run() error {
	res, err := g.client.ReleaseContent(g.release, helm.ContentReleaseVersion(g.version))
	if err != nil {
		return prettyError(err)
	}

	cfg, err := chartutil.CoalesceValues(res.Release.Chart, res.Release.Config)
	if err != nil {
		return err
	}
	cfgStr, err := cfg.YAML()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"Release":        res.Release,
		"ComputedValues": cfgStr,
		"ReleaseDate":    timeconv.Format(res.Release.Info.LastDeployed, time.ANSIC),
	}
	return tpl(getTemplate, data, g.out)
}

func tpl(t string, vals map[string]interface{}, out io.Writer) error {
	tt, err := template.New("_").Parse(t)
	if err != nil {
		return err
	}
	return tt.Execute(out, vals)
}

func ensureHelmClient(h helm.Interface) helm.Interface {
	if h != nil {
		return h
	}
	return helm.NewClient(helm.Host(tillerHost))
}
