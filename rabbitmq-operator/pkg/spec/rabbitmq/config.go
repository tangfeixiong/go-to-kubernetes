package rabbitmq

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	sampleapiv1alpha1 "github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/apis/example.com/v1alpha1"
	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/spec/artifact"
)

type RabbitmqRecipient struct {
	Namespace, TargetService                     string
	Plugins, Domain                              string
	K8sPort                                      int
	FirstNodeHost, SecondNodeHost, ThirdNodeHost string
	ThisNodeHost, ThisNodeName                   string
	clusteringMode                               sampleapiv1alpha1.ClusteringMode
}

type CfgOptFunc func(*RabbitmqRecipient) error

var (
	config_asset_name   string      = "template/rabbitmq.conf.tpl"
	plugins_asset_name  string      = "template/enabled_plugins.tpl"
	default_domain_name string      = "cluster.local"
	logger              *log.Logger = log.New(os.Stderr, fmt.Sprintf("[%s] ", filepath.Base(os.Args[0])), log.LstdFlags|log.Lshortfile)
)

func NewRabbitmqRecipient(options ...CfgOptFunc) (*RabbitmqRecipient, error) {
	recipe := &RabbitmqRecipient{
		Domain:         default_domain_name,
		K8sPort:        443,
		TargetService:  "my-rabbitmq",
		Namespace:      "default",
		clusteringMode: sampleapiv1alpha1.PeerDiscovery_K8s,
	}

	for _, option := range options {
		if err := option(recipe); err != nil {
			return nil, err
		}
	}

	return recipe, nil
}

func DomainNameSetter(v string) CfgOptFunc {
	return func(recipe *RabbitmqRecipient) error {
		if v != "" {
			recipe.Domain = v
		}
		return nil
	}
}

func K8sPortSetter(v int) CfgOptFunc {
	return func(recipe *RabbitmqRecipient) error {
		if 1024 < v && v < 65535 || v == 443 {
			recipe.K8sPort = v
		}
		return nil
	}
}

func EnabledPluginsSetter(v ...string) CfgOptFunc {
	return func(recipe *RabbitmqRecipient) error {
		if len(v) > 0 {
			recipe.Plugins = strings.Join(v, ",")
		}
		return nil
	}
}

func TargetServiceSetter(v string) CfgOptFunc {
	return func(recipe *RabbitmqRecipient) error {
		if v != "" {
			recipe.TargetService = v
		}
		return nil
	}
}

func NamespaceSetter(v string) CfgOptFunc {
	return func(recipe *RabbitmqRecipient) error {
		if v != "" {
			recipe.Namespace = v
		} else {
			recipe.Namespace = "default"
		}
		return nil
	}
}

func ClusteringModeSetter(v string) CfgOptFunc {
	return func(recipe *RabbitmqRecipient) error {
		recipe.clusteringMode = sampleapiv1alpha1.ClusteringMode(v)
		return nil
	}
}

func (recipe *RabbitmqRecipient) parse(tpl []byte) (*bytes.Buffer, error) {
	text := string(tpl)

	te := template.Must(template.New("rabbit").Parse(text))
	b := &bytes.Buffer{}
	if err := te.Execute(b, recipe); err != nil {
		return nil, fmt.Errorf("Executing template failed: %v\n", err)
	}

	return b, nil
}

func (recipe *RabbitmqRecipient) Parse(assetName string) (*bytes.Buffer, error) {
	asset, err := artifact.Asset(assetName)
	if err != nil {
		return nil, err
	}
	return recipe.parse(asset)
}

func (recipe *RabbitmqRecipient) Generate(dir string) error {
	if fi, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModeDir); err != nil {
				logger.Printf("Make dir %s failed: %s", dir, err.Error())
				return fmt.Errorf("Make dir %s failed: %s", dir, err.Error())
			}
		} else {
			logger.Printf("Stats %s failed: %s. %q", dir, err.Error(), fi)
			return fmt.Errorf("Stats %s failed: %s", dir, err.Error())
		}
	}

	b, err := recipe.Parse(plugins_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return err
	}

	path := filepath.Join(dir, "enabled_plugins")
	if err := ioutil.WriteFile(path, b.Bytes(), 0644); err != nil {
		logger.Printf("Create %s failed: %s", path, err.Error())
		return fmt.Errorf("Create %s failed: %s", path, err.Error())
	}
	logger.Println("Wrote", path)

	switch recipe.clusteringMode {
	case sampleapiv1alpha1.PeerDiscovery_K8s:
		b, err = recipe.Parse(config_asset_name)
		if err != nil {
			logger.Printf("Executing template failed: %s", err.Error())
			return err
		}

		path = filepath.Join(dir, "rabbitmq.conf")
		f, err := os.Create(path)
		if err != nil {
			logger.Printf("Create %s failed: %s", path, err.Error())
			return fmt.Errorf("Create %s failed: %s", path, err.Error())
		}
		defer f.Close()

		if _, err := f.Write(b.Bytes()); err != nil {
			logger.Printf("Writing %s failed: %s", path, err.Error())
			fmt.Errorf("Writing %s failed: %s", path, err.Error())
		}
		logger.Println("Wrote", path)
		break
	case sampleapiv1alpha1.Native:
	case sampleapiv1alpha1.None:
	default:
		return fmt.Errorf("Clustering mode not found: %v", recipe.clusteringMode)
	}

	return nil
}
