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
	certsphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	controlplanephase "k8s.io/kubernetes/cmd/kubeadm/app/phases/controlplane"
	etcdphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/etcd"
	kubeconfigphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/kubeconfig"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

func Test_kubeadm_MasterConfiguration(t *testing.T) {
	cfg := &kubeadmapiext.MasterConfiguration{}
	legacyscheme.Scheme.Default(cfg)
	cfg.API.AdvertiseAddress = "172.17.4.50"
	cfg.API.BindPort = 443
	cfg.KubernetesVersion = "stable-1.8"
	cfg.Networking.ServiceSubnet = "10.96.0.0/12"
	cfg.APIServerCertSANs = []string{"172.17.4.50", "10.96.0.1",
		"kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc.cluster.local"}
	//cfg.Networking.DNSDomain = "cluster.local"

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
	fmt.Println("%+q", i.cfg)
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

	if res, _ := certsphase.UsingExternalCA(i.cfg); !res {

		// PHASE 1: Generate certificates
		if err := certsphase.CreatePKIAssets(i.cfg); err != nil {
			return err
		}

		// PHASE 2: Generate kubeconfig files for the admin and the kubelet
		if err := kubeconfigphase.CreateInitKubeConfigFiles(kubeConfigDir, i.cfg); err != nil {
			return err
		}

	} else {
		fmt.Println("[externalca] The file 'ca.key' was not found, yet all other certificates are present. Using external CA mode - certificates or kubeconfig will not be generated.")
	}

	// Temporarily set cfg.CertificatesDir to the "real value" when writing controlplane manifests
	// This is needed for writing the right kind of manifests
	i.cfg.CertificatesDir = realCertsDir

	// PHASE 3: Bootstrap the control plane
	if err := controlplanephase.CreateInitStaticPodManifestFiles(manifestDir, i.cfg); err != nil {
		return fmt.Errorf("error creating init static pod manifest files: %v", err)
	}
	// Add etcd static pod spec only if external etcd is not configured
	if len(i.cfg.Etcd.Endpoints) == 0 {
		if err := etcdphase.CreateLocalEtcdStaticPodManifestFile(manifestDir, i.cfg); err != nil {
			return fmt.Errorf("error creating local etcd static pod manifest file: %v", err)
		}
	}

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
