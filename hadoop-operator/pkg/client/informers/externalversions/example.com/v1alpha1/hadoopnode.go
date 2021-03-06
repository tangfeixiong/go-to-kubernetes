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

// HadoopNodeInformer provides access to a shared informer and lister for
// HadoopNodes.
type HadoopNodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HadoopNodeLister
}

type hadoopNodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHadoopNodeInformer constructs a new informer for HadoopNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHadoopNodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHadoopNodeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHadoopNodeInformer constructs a new informer for HadoopNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHadoopNodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1alpha1().HadoopNodes(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1alpha1().HadoopNodes(namespace).Watch(options)
			},
		},
		&example_com_v1alpha1.HadoopNode{},
		resyncPeriod,
		indexers,
	)
}

func (f *hadoopNodeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHadoopNodeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hadoopNodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&example_com_v1alpha1.HadoopNode{}, f.defaultInformer)
}

func (f *hadoopNodeInformer) Lister() v1alpha1.HadoopNodeLister {
	return v1alpha1.NewHadoopNodeLister(f.Informer().GetIndexer())
}
