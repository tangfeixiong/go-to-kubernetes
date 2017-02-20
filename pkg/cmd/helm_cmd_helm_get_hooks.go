package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/helm"
)

const getHooksHelp = `
This command downloads hooks for a given release.

Hooks are formatted in YAML and separated by the YAML '---\n' separator.
`

type getHooksCmd struct {
	release string
	out     io.Writer
	client  helm.Interface
}

func newGetHooksCmd(client helm.Interface, out io.Writer) *cobra.Command {
	ghc := &getHooksCmd{
		out:    out,
		client: client,
	}
	cmd := &cobra.Command{
		Use:   "hooks [flags] RELEASE_NAME",
		Short: "download all hooks for a named release",
		Long:  getHooksHelp,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errReleaseRequired
			}
			ghc.release = args[0]
			ghc.client = ensureHelmClient(ghc.client)
			return ghc.run()
		},
	}
	return cmd
}

func (g *getHooksCmd) run() error {
	res, err := g.client.ReleaseContent(g.release)
	if err != nil {
		fmt.Fprintln(g.out, g.release)
		return prettyError(err)
	}

	for _, hook := range res.Release.Hooks {
		fmt.Fprintf(g.out, "---\n# %s\n%s", hook.Name, hook.Manifest)
	}
	return nil
}
