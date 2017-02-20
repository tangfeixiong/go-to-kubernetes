package client

import (
	"log"
	"os"
	"testing"
)

var (
	fake_kubeconfig_path string = "/data/src/github.com/openshift/origin/etc/kubeconfig"
	fake_kubectl_context string = "openshift-origin-single"
	fake_kube_apiserver  string = "https://172.17.4.50"

	fake_cf *CmdUtilFactory
)

func TestMain(m *testing.M) {
	/*
	   flag.Parse()

	   if user == "" || pass == "" {
	       fmt.Errorf("no enough args")
	       os.Exit(1)
	   }

	   c, err := NewClientWrapper(server, user, pass)
	   if err != nil {
	       fmt.Errorf(err.Error())
	       os.Exit(1)
	   }

	   if c == nil {
	       fmt.Errorf("no client")
	       os.Exit(1)
	   }
	*/
	var err error
	fake_cf, err = NewCmdUtilFactory(fake_kubeconfig_path, fake_kube_apiserver, fake_kubectl_context)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ret := m.Run()

	os.Exit(ret)
}
