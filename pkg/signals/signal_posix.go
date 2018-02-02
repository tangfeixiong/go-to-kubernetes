/*
  Inspired by:
    https://github.com/kubernetes/sample-controller/blob/master/pkg/signals/signal_posix.go
*/


// +build !windows

package signals

import (
	"os"
	"syscall"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
