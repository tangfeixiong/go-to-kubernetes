/*
  Inspired by:
    https://github.com/kubernetes/sample-controller/blob/master/pkg/signals/signal_windows.go
*/

package signals

import (
	"os"
)

var shutdownSignals = []os.Signal{os.Interrupt}
