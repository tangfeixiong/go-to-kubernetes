package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	scheme "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HadoopYarnsGetter has a method to return a HadoopYarnInterface.
// A group's client should implement this interface.
type HadoopYarnsGetter interface {
	HadoopYarns(namespace string) HadoopYarnInterface
}

// HadoopYarnInterface has methods to work with HadoopYarn resources.
type HadoopYarnInterface interface {
	Create(*v1alpha1.HadoopYarn) (*v1alpha1.HadoopYarn, error)
	Update(*v1alpha1.HadoopYarn) (*v1alpha1.HadoopYarn, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.HadoopYarn, error)
	List(opts v1.ListOptions) (*v1alpha1.HadoopYarnList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HadoopYarn, err error)
	HadoopYarnExpansion
}

// hadoopYarns implements HadoopYarnInterface
type hadoopYarns struct {
	client rest.Interface
	ns     string
}

// newHadoopYarns returns a HadoopYarns
func newHadoopYarns(c *ExampleV1alpha1Client, namespace string) *hadoopYarns {
	return &hadoopYarns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hadoopYarn, and returns the corresponding hadoopYarn object, and an error if there is any.
func (c *hadoopYarns) Get(name string, options v1.GetOptions) (result *v1alpha1.HadoopYarn, err error) {
	result = &v1alpha1.HadoopYarn{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoopyarns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HadoopYarns that match those selectors.
func (c *hadoopYarns) List(opts v1.ListOptions) (result *v1alpha1.HadoopYarnList, err error) {
	result = &v1alpha1.HadoopYarnList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoopyarns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hadoopYarns.
func (c *hadoopYarns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hadoopyarns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a hadoopYarn and creates it.  Returns the server's representation of the hadoopYarn, and an error, if there is any.
func (c *hadoopYarns) Create(hadoopYarn *v1alpha1.HadoopYarn) (result *v1alpha1.HadoopYarn, err error) {
	result = &v1alpha1.HadoopYarn{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hadoopyarns").
		Body(hadoopYarn).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hadoopYarn and updates it. Returns the server's representation of the hadoopYarn, and an error, if there is any.
func (c *hadoopYarns) Update(hadoopYarn *v1alpha1.HadoopYarn) (result *v1alpha1.HadoopYarn, err error) {
	result = &v1alpha1.HadoopYarn{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hadoopyarns").
		Name(hadoopYarn.Name).
		Body(hadoopYarn).
		Do().
		Into(result)
	return
}

// Delete takes name of the hadoopYarn and deletes it. Returns an error if one occurs.
func (c *hadoopYarns) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoopyarns").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hadoopYarns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoopyarns").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hadoopYarn.
func (c *hadoopYarns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HadoopYarn, err error) {
	result = &v1alpha1.HadoopYarn{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hadoopyarns").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
