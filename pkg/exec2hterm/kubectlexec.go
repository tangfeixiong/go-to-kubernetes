/*
Source from:
  k8s.io/kubernetes/pkg/kubectl/cmd/exec.go
*/
package exec2hterm

import (
	"fmt"
	"io"
	"os"

	"github.com/golang/glog"
	"github.com/kr/pty"
	gotty "github.com/yudai/gotty/app"

	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	kclientcmdapi "k8s.io/kubernetes/pkg/client/unversioned/clientcmd/api"
	kubectlcmd "k8s.io/kubernetes/pkg/kubectl/cmd"
)

func openPTYTTY() (ptyIO, tty *os.File, err error) {
	ptyIO, tty, err = pty.Open()
	if err != nil {
		return nil, nil, err
	}
	//defer tty.Close()
	//cmd.Stdout = tty
	//cmd.Stdin = tty
	//cmd.Stderr = tty
	//if cmd.SysProcAttr == nil {
	//	cmd.SysProcAttr = &syscall.SysProcAttr{}
	//}
	//cmd.SysProcAttr.Setctty = true
	//cmd.SysProcAttr.Setsid = true
	//err = cmd.Start()
	//if err != nil {
	//	pty.Close()
	//  return nil, err
	//}
	return
}

func NewCmdExec(clientConfig clientcmd.ClientConfig,
	ns, pod, container string, stdin, tty bool,
	ptyIO *os.File, cmdIn io.Reader, cmdOut, cmdErr io.Writer,
	command []string, options *gotty.Options) (*App, error) {
	app, err := NewApp(command, options)
	if err != nil {
		ptyIO.Close()
		return nil, err
	}
	app.ExecOptions = &kubectlcmd.ExecOptions{
		//Namespace:     ns,
		PodName:       pod,
		ContainerName: container,
		Stdin:         stdin,
		TTY:           tty,
		Command:       command,

		In:  cmdIn,
		Out: cmdOut,
		Err: cmdErr,

		Executor: &kubectlcmd.DefaultRemoteExecutor{},
	}

	if err := app.Complete(clientConfig); err != nil {
		ptyIO.Close()
		return nil, err
	}
	if ns != "" {
		app.ExecOptions.Namespace = ns
	}
	app.PTY = ptyIO
	return app, nil
}

// Complete verifies command line arguments and loads data from the command environment
func (app *App) Complete(clientConfig clientcmd.ClientConfig) error {
	p := app.ExecOptions
	// Let kubectl exec follow rules for `--`, see #13004 issue
	if len(p.PodName) == 0 {
		return fmt.Errorf("POD is required for exec")
	}
	if len(p.Command) < 1 {
		return fmt.Errorf("COMMAND is required for exec")
	}
	if clientConfig == nil {
		return fmt.Errorf("kubeconfig is required for exec")
	}

	namespace, _, err := clientConfig.Namespace() //f.DefaultNamespace()
	if err != nil {
		return err
	}
	p.Namespace = namespace

	config, err := clientConfig.ClientConfig() //f.ClientConfig()
	if err != nil {
		return err
	}
	p.Config = config

	client, err := kclient.New(config) //f.Client()
	if err != nil {
		return err
	}
	p.Client = client

	return nil
}

// k8s.io/kubernetes/pkg/client/unversioned/clientcmd/loader.go
func directKClientConfig(kubeconfigPath string, context, server string) clientcmd.ClientConfig {
	conf, err := clientcmd.LoadFromFile(kubeconfigPath)
	//conf, err := clientcmd.NewDefaultPathOptions().GetStartingConfig()
	if err != nil {
		glog.Errorf("client not configured: %v\n", err)
		return nil
	}
	glog.Infof("kube client configured: %+v\n", conf)

	var clientConfig clientcmd.ClientConfig
	overrides := &clientcmd.ConfigOverrides{}
	if server != "" {
		overrides.ClusterInfo = kclientcmdapi.Cluster{Server: server}
	}
	if context != "" {
		overrides.CurrentContext = context
	}
	clientConfig = clientcmd.NewNonInteractiveClientConfig(*conf,
		context, overrides)
	return clientConfig
}
