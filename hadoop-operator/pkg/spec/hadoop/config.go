package hadoop

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sampleapiv1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	"github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/spec/artifact"
)

type HadoopRecipient struct {
	CustomResourceName, ResourceName, ServiceName, Namespace, BaseDomain string
	NameNodeHost                                                         string
	NameNodePort, DfsReplication                                         int
	HdfsNodeType                                                         sampleapiv1alpha1.NodeType
	HdfsClusterID                                                        string
	ResourceCount                                                        int
}

type CfgOptFunc func(*HadoopRecipient) error

const (
	domain_example = "cluster.local"
)

var (
	core_asset_name   string      = "template/core-site.xml.gotpl"
	hdfs_asset_name   string      = "template/hdfs-site.xml.gotpl"
	yarn_asset_name   string      = "template/yarn-site.xml.gotpl"
	mapred_asset_name string      = "template/mapred-site.xml.gotpl"
	logger            *log.Logger = log.New(os.Stderr, fmt.Sprintf("[%s] ", filepath.Base(os.Args[0])), log.LstdFlags|log.Lshortfile)
)

func NewHadoopRecipient(options ...CfgOptFunc) (*HadoopRecipient, error) {
	recipe := &HadoopRecipient{
		Namespace:      apiv1.NamespaceDefault,
		BaseDomain:     domain_example,
		NameNodePort:   9000,
		DfsReplication: 3,
		HdfsNodeType:   sampleapiv1alpha1.NameNode,
	}

	for _, option := range options {
		if err := option(recipe); err != nil {
			return nil, err
		}
	}

	if recipe.NameNodeHost == "" {
		return nil, fmt.Errorf("NameNode e.g. IP or FQDN is required")
	}

	if recipe.CustomResourceName == "" {
		return nil, fmt.Errorf("CustomResourceName is required")
	}

	if recipe.ResourceName == "" {
		return nil, fmt.Errorf("ResourceName e.g. StatefulSet, DaemonSet, Deployment, ReplicationController and so on is required")
	}

	return recipe, nil
}

func CustomResourceNameSetter(v string) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		if v == "" {
			return fmt.Errorf("Invalid custom resource name")
		}
		recipe.CustomResourceName = v
		return nil
	}
}

func ResourceNameSetter(v string) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		if v == "" {
			return fmt.Errorf("Invalid resource name (e.g. StatefulSet, DaemonSet, Deployment, ReplicationController...)")
		}
		recipe.ResourceName = v
		return nil
	}
}

func NameNodeFqdnSetter(v ...string) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		if len(v) > 0 {
			recipe.NameNodeHost = strings.Join(v, ".")
		}
		return nil
	}
}

func NameNodeHostSetter(v string) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		if v != "" {
			recipe.NameNodeHost = v
		}
		return nil
	}
}

func ServiceNameSetter(v string) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		if v == "" {
			return fmt.Errorf("Invalid service name")
		}
		recipe.ServiceName = v
		return nil
	}
}

func NamespaceSetter(v string) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		if v != "" {
			recipe.Namespace = v
		}
		return nil
	}
}

func BaseDomainSetter(v string) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		if v != "" {
			recipe.BaseDomain = v
		}
		return nil
	}
}

func NameNodePortSetter(v int) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		if 1024 < v && v < 65535 || v == 443 {
			recipe.NameNodePort = v
		}
		return nil
	}
}

func HdfsNodeTypeSetter(v string) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		recipe.HdfsNodeType = sampleapiv1alpha1.NodeType(v)
		return nil
	}
}

func HdfsClusterIdSetter(v string) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		recipe.HdfsClusterID = v
		return nil
	}
}

func ResourceCountSetter(v int) CfgOptFunc {
	return func(recipe *HadoopRecipient) error {
		recipe.ResourceCount = v
		return nil
	}
}

func (recipe *HadoopRecipient) parse(tpl []byte) (*bytes.Buffer, error) {
	text := string(tpl)

	te := template.Must(template.New("xml").Parse(text))
	b := &bytes.Buffer{}
	if err := te.Execute(b, recipe); err != nil {
		return nil, fmt.Errorf("Executing template failed: %v\n", err)
	}

	return b, nil
}

func (recipe *HadoopRecipient) Parse(assetName string) (*bytes.Buffer, error) {
	asset, err := artifact.Asset(assetName)
	if err != nil {
		return nil, err
	}
	return recipe.parse(asset)
}

func (recipe *HadoopRecipient) New() (*apiv1.ConfigMap, error) {
	cm := apiv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      strings.Join([]string{recipe.CustomResourceName, recipe.ResourceName}, "--"),
			Namespace: recipe.Namespace,
		},
	}

	data := make(map[string]string)

	b, err := recipe.Parse(core_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return nil, err
	}
	data["core-site.xml"] = b.String()

	b, err = recipe.Parse(hdfs_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return nil, err
	}
	data["hdfs-site.xml"] = b.String()

	b, err = recipe.Parse(yarn_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return nil, err
	}
	data["yarn-site.xml"] = b.String()

	b, err = recipe.Parse(mapred_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return nil, err
	}
	data["mapred-site.xml"] = b.String()

	cm.Data = data

	return &cm, nil
}

func (recipe *HadoopRecipient) Generate(dir string) error {
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

	b, err := recipe.Parse(core_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return err
	}

	path := filepath.Join(dir, "core-site.xml")
	if err := ioutil.WriteFile(path, b.Bytes(), 0644); err != nil {
		logger.Printf("Create %s failed: %s", path, err.Error())
		return fmt.Errorf("Create %s failed: %s", path, err.Error())
	}
	logger.Println("Wrote", path)

	b, err = recipe.Parse(hdfs_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return err
	}

	path = filepath.Join(dir, "hdfs-site.xml")
	if err := ioutil.WriteFile(path, b.Bytes(), 0644); err != nil {
		logger.Printf("Create %s failed: %s", path, err.Error())
		return fmt.Errorf("Create %s failed: %s", path, err.Error())
	}
	logger.Println("Wrote", path)

	b, err = recipe.Parse(yarn_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return err
	}

	path = filepath.Join(dir, "yarn-site.xml")
	if err := ioutil.WriteFile(path, b.Bytes(), 0644); err != nil {
		logger.Printf("Create %s failed: %s", path, err.Error())
		return fmt.Errorf("Create %s failed: %s", path, err.Error())
	}
	logger.Println("Wrote", path)

	b, err = recipe.Parse(mapred_asset_name)
	if err != nil {
		logger.Printf("Executing template failed: %s", err.Error())
		return err
	}

	path = filepath.Join(dir, "mapred-site.xml")
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
