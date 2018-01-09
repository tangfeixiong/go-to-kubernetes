// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/tangfeixiong/go-to-kubernetes/redis-operator/pkg/apis/example.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RedisLister helps list Redises.
type RedisLister interface {
	// List lists all Redises in the indexer.
	List(selector labels.Selector) (ret []*v1.Redis, err error)
	// Redises returns an object that can list and get Redises.
	Redises(namespace string) RedisNamespaceLister
	RedisListerExpansion
}

// redisLister implements the RedisLister interface.
type redisLister struct {
	indexer cache.Indexer
}

// NewRedisLister returns a new RedisLister.
func NewRedisLister(indexer cache.Indexer) RedisLister {
	return &redisLister{indexer: indexer}
}

// List lists all Redises in the indexer.
func (s *redisLister) List(selector labels.Selector) (ret []*v1.Redis, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Redis))
	})
	return ret, err
}

// Redises returns an object that can list and get Redises.
func (s *redisLister) Redises(namespace string) RedisNamespaceLister {
	return redisNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RedisNamespaceLister helps list and get Redises.
type RedisNamespaceLister interface {
	// List lists all Redises in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Redis, err error)
	// Get retrieves the Redis from the indexer for a given namespace and name.
	Get(name string) (*v1.Redis, error)
	RedisNamespaceListerExpansion
}

// redisNamespaceLister implements the RedisNamespaceLister
// interface.
type redisNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Redises in the indexer for a given namespace.
func (s redisNamespaceLister) List(selector labels.Selector) (ret []*v1.Redis, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Redis))
	})
	return ret, err
}

// Get retrieves the Redis from the indexer for a given namespace and name.
func (s redisNamespaceLister) Get(name string) (*v1.Redis, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("redis"), name)
	}
	return obj.(*v1.Redis), nil
}
