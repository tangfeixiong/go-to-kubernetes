package hadoop

import (
	"bytes"
	"fmt"
	"io/ioutil"
	//"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/spec/artifact"
)

var (
	script_asset_name   string = "template/hadoop-startup.sh.gotpl"
	hadoop_home         string = "/hadoop-3.0.0"
	hdfs_bin_name       string = "hdfs"
	namenode_cmd_name   string = "namenode"
	datanode_cmd_name   string = "datanode"
	yarn_bin_name       string = "yarn"
	resman_cmd_name     string = "resourcemanager"
	nodeman_cmd_name    string = "nodemanager"
	proxysvr_cmd_name   string = "proxyserver"
	mapred_bin_name     string = "mpared"
	historysvr_cmd_name string = "historyserver"
	//logger              *log.Logger = log.New(os.Stderr, fmt.Sprintf("[%s] ", filepath.Base(os.Args[0])), log.LstdFlags|log.Lshortfile)
)

type HadoopStartupRecipient struct {
	hadoopHome, binName string
	BinPath, CmdName    string
	EtcPath             string
}

type StartupOptFunc func(*HadoopStartupRecipient) error

func NewHadoopStartupRecipient(options ...StartupOptFunc) (*HadoopStartupRecipient, error) {
	recipe := &HadoopStartupRecipient{
		hadoopHome: os.Getenv("HADOOP_HOME"),
	}

	for _, option := range options {
		if err := option(recipe); err != nil {
			return nil, err
		}
	}

	if recipe.hadoopHome == "" {
		return nil, fmt.Errorf("Hadoop directory is required")
	}
	fi, err := os.Stat(recipe.hadoopHome)
	if os.IsNotExist(err) || fi == nil {
		return nil, fmt.Errorf("Hadoop directory does not exist")
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("Invalid hadoop home")
	}

	if recipe.binName == "" {
		return nil, fmt.Errorf("Bin name is required, e.g. dfs, yarn, mapred")
	}
	recipe.BinPath = filepath.Join(recipe.hadoopHome, "bin", recipe.binName)
	_, err = os.Stat(recipe.BinPath)
	if err != nil {
		return nil, fmt.Errorf("Stats bin path failed: %v", err)
	}

	recipe.EtcPath = filepath.Join(recipe.hadoopHome, "etc/hadoop")
	_, err = os.Stat(recipe.EtcPath)
	if err != nil {
		return nil, fmt.Errorf("Stats etc path failed: %v", err)
	}

	if recipe.CmdName == "" {
		return nil, fmt.Errorf("CmdName e.g. namenode, datanode, resourcemanager, nodemanager, proxyserver and historyserver is required")
	}

	return recipe, nil
}

func HadoopHomeSetter(v string) StartupOptFunc {
	return func(recipe *HadoopStartupRecipient) error {
		if v == "" {
			return fmt.Errorf("Invalid hadoop home")
		}
		recipe.hadoopHome = v
		return nil
	}
}

func BinNameSetter(v string) StartupOptFunc {
	return func(recipe *HadoopStartupRecipient) error {
		if v == "" {
			return fmt.Errorf("Invalid bin name")
		}
		recipe.binName = v
		return nil
	}
}

func CmdNameSetter(v string) StartupOptFunc {
	return func(recipe *HadoopStartupRecipient) error {
		if v == "" {
			return fmt.Errorf("Invalid cmd name")
		}
		recipe.CmdName = v
		return nil
	}
}

func (recipe *HadoopStartupRecipient) parse(tpl []byte) (*bytes.Buffer, error) {
	text := string(tpl)

	te := template.Must(template.New("sh").Parse(text))
	b := &bytes.Buffer{}
	if err := te.Execute(b, recipe); err != nil {
		return nil, fmt.Errorf("Executing template failed: %v\n", err)
	}

	return b, nil
}

func (recipe *HadoopStartupRecipient) Parse(assetName string) (*bytes.Buffer, error) {
	asset, err := artifact.Asset(assetName)
	if err != nil {
		return nil, err
	}
	return recipe.parse(asset)
}

func (recipe *HadoopStartupRecipient) Generate(dir string) error {
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

	b, err := recipe.Parse(script_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return err
	}

	path := filepath.Join(dir, "hadoop-startup.sh")
	if err := ioutil.WriteFile(path, b.Bytes(), 0755); err != nil {
		logger.Printf("Create %s failed: %s", path, err.Error())
		return fmt.Errorf("Create %s failed: %s", path, err.Error())
	}
	logger.Println("Wrote", path)

	return nil
}
