package kubeadm

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmapiext "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
	// "k8s.io/kubernetes/cmd/kubeadm/app/cmd"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

func Test_kubeadm_MasterConfiguration(t *testing.T) {
	cfg := &kubeadmapiext.MasterConfiguration{}
	legacyscheme.Scheme.Default(cfg)
	cfg.API.AdvertiseAddress = "10.64.33.82"
	cfg.API.BindPort = 443
	cfg.KubernetesVersion = "v1.8.2"

	legacyscheme.Scheme.Default(cfg)
	internalcfg := &kubeadmapi.MasterConfiguration{}
	legacyscheme.Scheme.Convert(cfg, internalcfg, nil)
	err := configutil.SetInitDynamicDefaults(internalcfg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", cfg)

	skipTokenPrint := false
	dryRun := true
	i := &Init{cfg: internalcfg, skipTokenPrint: skipTokenPrint, dryRun: dryRun}

	if err := i.Run(os.Stdout); err != nil {
		t.Fatal(err)
	}
}

// Init defines struct used by "kubeadm init" command
type Init struct {
	cfg            *kubeadmapi.MasterConfiguration
	skipTokenPrint bool
	dryRun         bool
}

// Run executes master node provisioning, including certificates, needed static pod manifests, etc.
func (i *Init) Run(out io.Writer) error {
	// Get directories to write files to; can be faked if we're dry-running
	realCertsDir := i.cfg.CertificatesDir
	certsDirToWriteTo, kubeConfigDir, manifestDir, err := getDirectoriesToUse(i.dryRun, i.cfg.CertificatesDir)
	if err != nil {
		return fmt.Errorf("error getting directories to use: %v", err)
	}
	// certsDirToWriteTo is gonna equal cfg.CertificatesDir in the normal case, but gonna be a temp directory if dryrunning
	i.cfg.CertificatesDir = certsDirToWriteTo

	adminKubeConfigPath := filepath.Join(kubeConfigDir, kubeadmconstants.AdminKubeConfigFileName)

	println(realCertsDir, certsDirToWriteTo, kubeConfigDir, manifestDir, adminKubeConfigPath)

	return nil
}

// getDirectoriesToUse returns the (in order) certificates, kubeconfig and Static Pod manifest directories, followed by a possible error
// This behaves differently when dry-running vs the normal flow
func getDirectoriesToUse(dryRun bool, defaultPkiDir string) (string, string, string, error) {
	if dryRun {
		dryRunDir, err := ioutil.TempDir("", "kubeadm-init-dryrun")
		if err != nil {
			return "", "", "", fmt.Errorf("couldn't create a temporary directory: %v", err)
		}
		// Use the same temp dir for all
		return dryRunDir, dryRunDir, dryRunDir, nil
	}

	return defaultPkiDir, kubeadmconstants.KubernetesDir, kubeadmconstants.GetStaticPodDirectory(), nil
}
