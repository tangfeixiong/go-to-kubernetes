package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/helm"
)

var getValuesHelp = `
This command downloads a values file for a given release.
`

type getValuesCmd struct {
	release   string
	allValues bool
	out       io.Writer
	client    helm.Interface
}

func newGetValuesCmd(client helm.Interface, out io.Writer) *cobra.Command {
	get := &getValuesCmd{
		out:    out,
		client: client,
	}
	cmd := &cobra.Command{
		Use:   "values [flags] RELEASE_NAME",
		Short: "download the values file for a named release",
		Long:  getValuesHelp,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errReleaseRequired
			}
			get.release = args[0]
			get.client = ensureHelmClient(get.client)
			return get.run()
		},
	}
	cmd.Flags().BoolVarP(&get.allValues, "all", "a", false, "dump all (computed) values")
	return cmd
}

// getValues implements 'helm get values'
func (g *getValuesCmd) run() error {
	res, err := g.client.ReleaseContent(g.release)
	if err != nil {
		return prettyError(err)
	}

	// If the user wants all values, compute the values and return.
	if g.allValues {
		cfg, err := chartutil.CoalesceValues(res.Release.Chart, res.Release.Config)
		if err != nil {
			return err
		}
		cfgStr, err := cfg.YAML()
		if err != nil {
			return err
		}
		fmt.Fprintln(g.out, cfgStr)
		return nil
	}

	fmt.Fprintln(g.out, res.Release.Config.Raw)
	return nil
}
