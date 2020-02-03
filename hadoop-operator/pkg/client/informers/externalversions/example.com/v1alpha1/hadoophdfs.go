// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	example_com_v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	versioned "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/listers/example.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HadoopHdfsInformer provides access to a shared informer and lister for
// HadoopHdfses.
type HadoopHdfsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HadoopHdfsLister
}

type hadoopHdfsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHadoopHdfsInformer constructs a new informer for HadoopHdfs type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHadoopHdfsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHadoopHdfsInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHadoopHdfsInformer constructs a new informer for HadoopHdfs type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHadoopHdfsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1alpha1().HadoopHdfses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1alpha1().HadoopHdfses(namespace).Watch(options)
			},
		},
		&example_com_v1alpha1.HadoopHdfs{},
		resyncPeriod,
		indexers,
	)
}

func (f *hadoopHdfsInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHadoopHdfsInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hadoopHdfsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&example_com_v1alpha1.HadoopHdfs{}, f.defaultInformer)
}

func (f *hadoopHdfsInformer) Lister() v1alpha1.HadoopHdfsLister {
	return v1alpha1.NewHadoopHdfsLister(f.Informer().GetIndexer())
}