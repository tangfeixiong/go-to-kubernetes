package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	scheme "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HadoopsGetter has a method to return a HadoopInterface.
// A group's client should implement this interface.
type HadoopsGetter interface {
	Hadoops(namespace string) HadoopInterface
}

// HadoopInterface has methods to work with Hadoop resources.
type HadoopInterface interface {
	Create(*v1alpha1.Hadoop) (*v1alpha1.Hadoop, error)
	Update(*v1alpha1.Hadoop) (*v1alpha1.Hadoop, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Hadoop, error)
	List(opts v1.ListOptions) (*v1alpha1.HadoopList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Hadoop, err error)
	HadoopExpansion
}

// hadoops implements HadoopInterface
type hadoops struct {
	client rest.Interface
	ns     string
}

// newHadoops returns a Hadoops
func newHadoops(c *ExampleV1alpha1Client, namespace string) *hadoops {
	return &hadoops{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hadoop, and returns the corresponding hadoop object, and an error if there is any.
func (c *hadoops) Get(name string, options v1.GetOptions) (result *v1alpha1.Hadoop, err error) {
	result = &v1alpha1.Hadoop{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoops").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Hadoops that match those selectors.
func (c *hadoops) List(opts v1.ListOptions) (result *v1alpha1.HadoopList, err error) {
	result = &v1alpha1.HadoopList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hadoops.
func (c *hadoops) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hadoops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a hadoop and creates it.  Returns the server's representation of the hadoop, and an error, if there is any.
func (c *hadoops) Create(hadoop *v1alpha1.Hadoop) (result *v1alpha1.Hadoop, err error) {
	result = &v1alpha1.Hadoop{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hadoops").
		Body(hadoop).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hadoop and updates it. Returns the server's representation of the hadoop, and an error, if there is any.
func (c *hadoops) Update(hadoop *v1alpha1.Hadoop) (result *v1alpha1.Hadoop, err error) {
	result = &v1alpha1.Hadoop{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hadoops").
		Name(hadoop.Name).
		Body(hadoop).
		Do().
		Into(result)
	return
}

// Delete takes name of the hadoop and deletes it. Returns an error if one occurs.
func (c *hadoops) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoops").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hadoops) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoops").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hadoop.
func (c *hadoops) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Hadoop, err error) {
	result = &v1alpha1.Hadoop{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hadoops").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
