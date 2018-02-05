package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	scheme "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HadoopMapreducesGetter has a method to return a HadoopMapreduceInterface.
// A group's client should implement this interface.
type HadoopMapreducesGetter interface {
	HadoopMapreduces(namespace string) HadoopMapreduceInterface
}

// HadoopMapreduceInterface has methods to work with HadoopMapreduce resources.
type HadoopMapreduceInterface interface {
	Create(*v1alpha1.HadoopMapreduce) (*v1alpha1.HadoopMapreduce, error)
	Update(*v1alpha1.HadoopMapreduce) (*v1alpha1.HadoopMapreduce, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.HadoopMapreduce, error)
	List(opts v1.ListOptions) (*v1alpha1.HadoopMapreduceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HadoopMapreduce, err error)
	HadoopMapreduceExpansion
}

// hadoopMapreduces implements HadoopMapreduceInterface
type hadoopMapreduces struct {
	client rest.Interface
	ns     string
}

// newHadoopMapreduces returns a HadoopMapreduces
func newHadoopMapreduces(c *ExampleV1alpha1Client, namespace string) *hadoopMapreduces {
	return &hadoopMapreduces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hadoopMapreduce, and returns the corresponding hadoopMapreduce object, and an error if there is any.
func (c *hadoopMapreduces) Get(name string, options v1.GetOptions) (result *v1alpha1.HadoopMapreduce, err error) {
	result = &v1alpha1.HadoopMapreduce{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoopmapreduces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HadoopMapreduces that match those selectors.
func (c *hadoopMapreduces) List(opts v1.ListOptions) (result *v1alpha1.HadoopMapreduceList, err error) {
	result = &v1alpha1.HadoopMapreduceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoopmapreduces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hadoopMapreduces.
func (c *hadoopMapreduces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hadoopmapreduces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a hadoopMapreduce and creates it.  Returns the server's representation of the hadoopMapreduce, and an error, if there is any.
func (c *hadoopMapreduces) Create(hadoopMapreduce *v1alpha1.HadoopMapreduce) (result *v1alpha1.HadoopMapreduce, err error) {
	result = &v1alpha1.HadoopMapreduce{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hadoopmapreduces").
		Body(hadoopMapreduce).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hadoopMapreduce and updates it. Returns the server's representation of the hadoopMapreduce, and an error, if there is any.
func (c *hadoopMapreduces) Update(hadoopMapreduce *v1alpha1.HadoopMapreduce) (result *v1alpha1.HadoopMapreduce, err error) {
	result = &v1alpha1.HadoopMapreduce{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hadoopmapreduces").
		Name(hadoopMapreduce.Name).
		Body(hadoopMapreduce).
		Do().
		Into(result)
	return
}

// Delete takes name of the hadoopMapreduce and deletes it. Returns an error if one occurs.
func (c *hadoopMapreduces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoopmapreduces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hadoopMapreduces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoopmapreduces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hadoopMapreduce.
func (c *hadoopMapreduces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HadoopMapreduce, err error) {
	result = &v1alpha1.HadoopMapreduce{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hadoopmapreduces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
