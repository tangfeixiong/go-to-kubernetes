package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	scheme "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HadoopHdfsesGetter has a method to return a HadoopHdfsInterface.
// A group's client should implement this interface.
type HadoopHdfsesGetter interface {
	HadoopHdfses(namespace string) HadoopHdfsInterface
}

// HadoopHdfsInterface has methods to work with HadoopHdfs resources.
type HadoopHdfsInterface interface {
	Create(*v1alpha1.HadoopHdfs) (*v1alpha1.HadoopHdfs, error)
	Update(*v1alpha1.HadoopHdfs) (*v1alpha1.HadoopHdfs, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.HadoopHdfs, error)
	List(opts v1.ListOptions) (*v1alpha1.HadoopHdfsList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HadoopHdfs, err error)
	HadoopHdfsExpansion
}

// hadoopHdfses implements HadoopHdfsInterface
type hadoopHdfses struct {
	client rest.Interface
	ns     string
}

// newHadoopHdfses returns a HadoopHdfses
func newHadoopHdfses(c *ExampleV1alpha1Client, namespace string) *hadoopHdfses {
	return &hadoopHdfses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hadoopHdfs, and returns the corresponding hadoopHdfs object, and an error if there is any.
func (c *hadoopHdfses) Get(name string, options v1.GetOptions) (result *v1alpha1.HadoopHdfs, err error) {
	result = &v1alpha1.HadoopHdfs{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoophdfses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HadoopHdfses that match those selectors.
func (c *hadoopHdfses) List(opts v1.ListOptions) (result *v1alpha1.HadoopHdfsList, err error) {
	result = &v1alpha1.HadoopHdfsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoophdfses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hadoopHdfses.
func (c *hadoopHdfses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hadoophdfses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a hadoopHdfs and creates it.  Returns the server's representation of the hadoopHdfs, and an error, if there is any.
func (c *hadoopHdfses) Create(hadoopHdfs *v1alpha1.HadoopHdfs) (result *v1alpha1.HadoopHdfs, err error) {
	result = &v1alpha1.HadoopHdfs{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hadoophdfses").
		Body(hadoopHdfs).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hadoopHdfs and updates it. Returns the server's representation of the hadoopHdfs, and an error, if there is any.
func (c *hadoopHdfses) Update(hadoopHdfs *v1alpha1.HadoopHdfs) (result *v1alpha1.HadoopHdfs, err error) {
	result = &v1alpha1.HadoopHdfs{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hadoophdfses").
		Name(hadoopHdfs.Name).
		Body(hadoopHdfs).
		Do().
		Into(result)
	return
}

// Delete takes name of the hadoopHdfs and deletes it. Returns an error if one occurs.
func (c *hadoopHdfses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoophdfses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hadoopHdfses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoophdfses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hadoopHdfs.
func (c *hadoopHdfses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HadoopHdfs, err error) {
	result = &v1alpha1.HadoopHdfs{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hadoophdfses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
