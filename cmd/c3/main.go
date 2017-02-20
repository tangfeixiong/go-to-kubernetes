package main

import (
	"os"
	"runtime"

	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"

	"github.com/tangfeixiong/go-to-kubernetes/pkg/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

type cli struct{}

var app = cli{}

func (app cli) Run() error {
	f := cmdutil.NewFactory(nil)

	c := cmd.NewC3Command(f, os.Stdin, os.Stdout, os.Stderr)

	return c.Execute()
}
