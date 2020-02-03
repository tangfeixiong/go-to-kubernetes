// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Hadoops returns a HadoopInformer.
	Hadoops() HadoopInformer
	// HadoopHdfses returns a HadoopHdfsInformer.
	HadoopHdfses() HadoopHdfsInformer
	// HadoopMapreduces returns a HadoopMapreduceInformer.
	HadoopMapreduces() HadoopMapreduceInformer
	// HadoopNodes returns a HadoopNodeInformer.
	HadoopNodes() HadoopNodeInformer
	// HadoopYarns returns a HadoopYarnInformer.
	HadoopYarns() HadoopYarnInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Hadoops returns a HadoopInformer.
func (v *version) Hadoops() HadoopInformer {
	return &hadoopInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HadoopHdfses returns a HadoopHdfsInformer.
func (v *version) HadoopHdfses() HadoopHdfsInformer {
	return &hadoopHdfsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HadoopMapreduces returns a HadoopMapreduceInformer.
func (v *version) HadoopMapreduces() HadoopMapreduceInformer {
	return &hadoopMapreduceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HadoopNodes returns a HadoopNodeInformer.
func (v *version) HadoopNodes() HadoopNodeInformer {
	return &hadoopNodeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HadoopYarns returns a HadoopYarnInformer.
func (v *version) HadoopYarns() HadoopYarnInformer {
	return &hadoopYarnInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}