package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/apis/example.com/v1alpha1"
	scheme "github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MariadbsGetter has a method to return a MariadbInterface.
// A group's client should implement this interface.
type MariadbsGetter interface {
	Mariadbs(namespace string) MariadbInterface
}

// MariadbInterface has methods to work with Mariadb resources.
type MariadbInterface interface {
	Create(*v1alpha1.Mariadb) (*v1alpha1.Mariadb, error)
	Update(*v1alpha1.Mariadb) (*v1alpha1.Mariadb, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Mariadb, error)
	List(opts v1.ListOptions) (*v1alpha1.MariadbList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Mariadb, err error)
	MariadbExpansion
}

// mariadbs implements MariadbInterface
type mariadbs struct {
	client rest.Interface
	ns     string
}

// newMariadbs returns a Mariadbs
func newMariadbs(c *ExampleV1alpha1Client, namespace string) *mariadbs {
	return &mariadbs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mariadb, and returns the corresponding mariadb object, and an error if there is any.
func (c *mariadbs) Get(name string, options v1.GetOptions) (result *v1alpha1.Mariadb, err error) {
	result = &v1alpha1.Mariadb{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mariadbs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Mariadbs that match those selectors.
func (c *mariadbs) List(opts v1.ListOptions) (result *v1alpha1.MariadbList, err error) {
	result = &v1alpha1.MariadbList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mariadbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mariadbs.
func (c *mariadbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mariadbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a mariadb and creates it.  Returns the server's representation of the mariadb, and an error, if there is any.
func (c *mariadbs) Create(mariadb *v1alpha1.Mariadb) (result *v1alpha1.Mariadb, err error) {
	result = &v1alpha1.Mariadb{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mariadbs").
		Body(mariadb).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mariadb and updates it. Returns the server's representation of the mariadb, and an error, if there is any.
func (c *mariadbs) Update(mariadb *v1alpha1.Mariadb) (result *v1alpha1.Mariadb, err error) {
	result = &v1alpha1.Mariadb{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mariadbs").
		Name(mariadb.Name).
		Body(mariadb).
		Do().
		Into(result)
	return
}

// Delete takes name of the mariadb and deletes it. Returns an error if one occurs.
func (c *mariadbs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mariadbs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mariadbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mariadbs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mariadb.
func (c *mariadbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Mariadb, err error) {
	result = &v1alpha1.Mariadb{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mariadbs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
