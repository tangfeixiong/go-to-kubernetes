// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HadoopLister helps list Hadoops.
type HadoopLister interface {
	// List lists all Hadoops in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Hadoop, err error)
	// Hadoops returns an object that can list and get Hadoops.
	Hadoops(namespace string) HadoopNamespaceLister
	HadoopListerExpansion
}

// hadoopLister implements the HadoopLister interface.
type hadoopLister struct {
	indexer cache.Indexer
}

// NewHadoopLister returns a new HadoopLister.
func NewHadoopLister(indexer cache.Indexer) HadoopLister {
	return &hadoopLister{indexer: indexer}
}

// List lists all Hadoops in the indexer.
func (s *hadoopLister) List(selector labels.Selector) (ret []*v1alpha1.Hadoop, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Hadoop))
	})
	return ret, err
}

// Hadoops returns an object that can list and get Hadoops.
func (s *hadoopLister) Hadoops(namespace string) HadoopNamespaceLister {
	return hadoopNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HadoopNamespaceLister helps list and get Hadoops.
type HadoopNamespaceLister interface {
	// List lists all Hadoops in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Hadoop, err error)
	// Get retrieves the Hadoop from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Hadoop, error)
	HadoopNamespaceListerExpansion
}

// hadoopNamespaceLister implements the HadoopNamespaceLister
// interface.
type hadoopNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Hadoops in the indexer for a given namespace.
func (s hadoopNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Hadoop, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Hadoop))
	})
	return ret, err
}

// Get retrieves the Hadoop from the indexer for a given namespace and name.
func (s hadoopNamespaceLister) Get(name string) (*v1alpha1.Hadoop, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hadoop"), name)
	}
	return obj.(*v1alpha1.Hadoop), nil
}