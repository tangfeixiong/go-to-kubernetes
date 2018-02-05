/*
   Inspired by:
    https://github.com/helm/helm-classic/blob/master/kubectl/get.go
*/
// package kubectl
package hadoop

import (
	"fmt"
	"os"
)

func (r RealRunner) Format(hdfs, nodetype, clusterid string) error {
	args := []string{hdfs + " " + nodetype + " -format"}
	if nodetype == "namenode" && clusterid != "" {
		args[0] += " -clusterid " + clusterid
	}

	args = append([]string{"-c"}, args...)
	cmd := command(args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	// return cmd.CombinedOutput()

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("Executing format failed: %v", err)
	}

	//Waiting for command to finish...
	return cmd.Wait()
}

//// Get returns Kubernetes resources
//func (r RealRunner) Get(stdin []byte, ns string) ([]byte, error) {
//	args := []string{"get", "-f", "-"}

//	if ns != "" {
//		args = append([]string{"--namespace=" + ns}, args...)
//	}
//	cmd := command(args...)
//	assignStdin(cmd, stdin)

//	return cmd.CombinedOutput()
//}

//// Get returns the commands to kubectl
//func (r PrintRunner) Get(stdin []byte, ns string) ([]byte, error) {
//	args := []string{"get", "-f", "-"}

//	if ns != "" {
//		args = append([]string{"--namespace=" + ns}, args...)
//	}
//	cmd := command(args...)
//	assignStdin(cmd, stdin)

//	return []byte(cmd.String()), nil
//}
