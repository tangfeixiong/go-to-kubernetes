package exec2hterm

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"

	"github.com/fatih/structs"
	"github.com/golang/glog"
	gotty "github.com/yudai/gotty/app"
)

var (
	_ns        string = "default"
	_pod       string = "netcat-simple"
	_container string = ""
	_command          = []string{"/bin/ash"}
	_stdin     bool   = true
	_tty       bool   = true

	_kubeconfig string = "/data/src/github.com/openshift/origin/openshift.local.config/master/kubeconfig"

	_context   string = "openshift-origin-single"
	_apiserver string = "https://172.17.4.50"

	_timeoutSeconds int64 = 30
)

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

var helpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [options] <command> [<arguments...>]

VERSION:
   {{.Version}}{{if or .Author .Email}}

AUTHOR:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

func GoTTY() {

	flags := []struct {
		name        string
		shortName   string
		description string
	}{
		{"address", "a", "IP address to listen"},
		{"port", "p", "Port number to listen"},
		{"permit-write", "w", "Permit clients to write to the TTY (BE CAREFUL)"},
		{"credential", "c", "Credential for Basic Authentication (ex: user:pass, default disabled)"},
		{"random-url", "r", "Add a random string to the URL"},
		{"random-url-length", "", "Random URL length"},
		{"tls", "t", "Enable TLS/SSL"},
		{"tls-crt", "", "TLS/SSL certificate file path"},
		{"tls-key", "", "TLS/SSL key file path"},
		{"tls-ca-crt", "", "TLS/SSL CA certificate file for client certifications"},
		{"index", "", "Custom index.html file"},
		{"title-format", "", "Title format of browser window"},
		{"reconnect", "", "Enable reconnection"},
		{"reconnect-time", "", "Time to reconnect"},
		{"once", "", "Accept only one client and exit on disconnection"},
		{"permit-arguments", "", "Permit clients to send command line arguments in URL (e.g. http://example.com:8080/?arg=AAA&arg=BBB)"},
		{"close-signal", "", "Signal sent to the command process when gotty close it (default: SIGHUP)"},
	}

	mappingHint := map[string]string{
		"index":      "IndexFile",
		"tls":        "EnableTLS",
		"tls-crt":    "TLSCrtFile",
		"tls-key":    "TLSKeyFile",
		"tls-ca-crt": "TLSCACrtFile",
		"random-url": "EnableRandomUrl",
		"reconnect":  "EnableReconnect",
	}

	options := gotty.DefaultOptions

	var configFile string = "~/.gotty"
	if v, ok := os.LookupEnv("GOTTY_CONFIG"); ok {
		configFile = v
	}
	args := os.Args[1:]
	//si := SliceIndex(len(args), func(i int) bool { return strings.Contains(args[i], "-config") })

	flag.StringVar(&configFile, "config", configFile, "Config file path")

	//cliFlags, err := generateFlags(flags, mappingHint)
	//applyFlags(&options, flags, mappingHint, c)
	o := structs.New(&options)
	for _, f := range flags {
		fieldname := fieldName(f.name, mappingHint)

		field, ok := o.FieldOk(fieldname)
		if !ok {
			glog.Warningf("No such field: %s", fieldname)
			continue
		}

		//flagname := f.name
		//if f.shortName != "" {
		//	flagname += ", " + f.shortName
		//}
		envName := "GOTTY_" + strings.ToUpper(strings.Join(strings.Split(f.name, "-"), "_"))

		//field := o.Field(fieldname)
		field.Set(os.Getenv(envName))
		var val interface{}
		switch field.Kind() {
		case reflect.String:
			val = flag.String(f.name, field.Value().(string), f.description)
		case reflect.Bool:
			val = flag.Bool(f.name, false, f.description)
		case reflect.Int:
			val = flag.Int(f.name, field.Value().(int), f.description)
		}
		field.Set(val)
	}

	flag.Parse()

	_, err := os.Stat(gotty.ExpandHomeDir(configFile))
	if configFile != "~/.gotty" || !os.IsNotExist(err) {
		if err := gotty.ApplyConfigFile(&options, configFile); err != nil {
			exit(err, 2)
		}
	}

	if options.Credential != "" {
		options.EnableBasicAuth = true
	}
	if options.TLSCACrtFile != "" {
		options.EnableTLSClientAuth = true
	}

	if err := gotty.CheckConfig(&options); err != nil {
		exit(err, 6)
	}

	var app *App

	if len(args) == 0 {
		app, err = NewApp([]string{"kubectl", "exec", "-ti", "netcat-simple", "ash"}, &options)
	} else {
		clientConfig := directKClientConfig(_kubeconfig, _context, _apiserver)
		if clientConfig == nil {
			exit(fmt.Errorf("kubeclient required"), 1)
		}
		pty, tty, err := openPTYTTY()
		if err != nil {
			exit(err, 1)
		}
		app, err = NewCmdExec(clientConfig, _ns, _pod, _container, _stdin, _tty, pty, tty, tty, tty, _command, &options)
	}
	if err != nil {
		exit(err, 3)
	}

	registerSignals(app)

	err = app.Run()
	if err != nil {
		exit(err, 4)
	}
}

func registerSignals(app *App) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		for {
			s := <-sigChan
			switch s {
			case syscall.SIGINT, syscall.SIGTERM:
				if app.Exit() {
					fmt.Println("Send ^C to force exit.")
				} else {
					os.Exit(5)
				}
			}
		}
	}()
}

func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

//func generateFlags(flags []flag, hint map[string]string) ([]cli.Flag, error) {}
//func applyFlags(options *app.Options, flags []flag, mappingHint map[string]string, c *cli.Context) {}

func fieldName(name string, hint map[string]string) string {
	if fieldName, ok := hint[name]; ok {
		return fieldName
	}
	nameParts := strings.Split(name, "-")
	for i, part := range nameParts {
		nameParts[i] = strings.ToUpper(part[0:1]) + part[1:]
	}
	return strings.Join(nameParts, "")
}
