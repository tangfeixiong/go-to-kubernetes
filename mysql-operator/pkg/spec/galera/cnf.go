package galera

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/spec/artifact"
)

type Recipient struct {
	ClusterName                                  string
	ClusterAddresses, WsrepProviderOptions       string
	FirstNodeHost, SecondNodeHost, ThirdNodeHost string
	ThisNodeHost, ThisNodeName                   string
	DisableClusterAddresses                      bool
}

type ConfigOptionFunc func(*Recipient) error

var (
	default_asset_name string      = "template/galera.cnf.tpl"
	logger             *log.Logger = log.New(os.Stderr, fmt.Sprintf("[%s] ", filepath.Base(os.Args[0])), log.LstdFlags|log.Lshortfile)
)

func NewRecipient(disableClusterAddresses bool, options ...ConfigOptionFunc) (*Recipient, error) {
	recipe := &Recipient{
		ClusterName:             "my-cluster",
		DisableClusterAddresses: disableClusterAddresses,
	}

	for _, option := range options {
		if err := option(recipe); err != nil {
			return nil, err
		}
	}

	return recipe, nil
}

func ClusterNameSetter(v string) ConfigOptionFunc {
	return func(recipe *Recipient) error {
		if v != "" {
			recipe.ClusterName = v
		}
		return nil
	}
}

func ClusterAddressesSetter(v ...string) ConfigOptionFunc {
	return func(recipe *Recipient) error {
		if len(v) > 0 {
			recipe.ClusterAddresses = strings.Join(v, ",")
		}
		return nil
	}
}

func WsrepProviderOptionsSetter(v ...string) ConfigOptionFunc {
	return func(recipe *Recipient) error {
		if len(v) > 0 {
			recipe.WsrepProviderOptions = strings.Join(v, "&")
		}
		return nil
	}
}

func MinimalClusterSetter(first, second, third string) ConfigOptionFunc {
	return func(recipe *Recipient) error {
		recipe.ClusterAddresses = ""
		recipe.FirstNodeHost = first
		recipe.SecondNodeHost = second
		recipe.ThirdNodeHost = third
		return nil
	}
}

func ThisNodeSetter(host, name string) ConfigOptionFunc {
	return func(recipe *Recipient) error {
		if host != "" {
			recipe.ThisNodeHost = host
		}
		if name != "" {
			recipe.ThisNodeName = name
		}
		return nil
	}
}

func (recipe *Recipient) parse(tpl []byte) (*bytes.Buffer, error) {
	text := string(tpl)

	te := template.Must(template.New("cnf").Parse(text))
	b := &bytes.Buffer{}
	if err := te.Execute(b, recipe); err != nil {
		return nil, fmt.Errorf("Executing template failed: %v\n", err)
	}

	return b, nil
}

func (recipe *Recipient) Parse() (*bytes.Buffer, error) {
	asset, err := artifact.Asset(default_asset_name)
	if err != nil {
		return nil, err
	}
	return recipe.parse(asset)
}

func (recipe *Recipient) Generate(path string) error {
	b, err := recipe.Parse()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
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

	//err := ioutil.WriteFile(path, b.Bytes(), 0644)

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
	return nil
}
