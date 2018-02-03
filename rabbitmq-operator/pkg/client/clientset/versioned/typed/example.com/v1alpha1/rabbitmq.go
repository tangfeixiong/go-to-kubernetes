package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/apis/example.com/v1alpha1"
	scheme "github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RabbitmqsGetter has a method to return a RabbitmqInterface.
// A group's client should implement this interface.
type RabbitmqsGetter interface {
	Rabbitmqs(namespace string) RabbitmqInterface
}

// RabbitmqInterface has methods to work with Rabbitmq resources.
type RabbitmqInterface interface {
	Create(*v1alpha1.Rabbitmq) (*v1alpha1.Rabbitmq, error)
	Update(*v1alpha1.Rabbitmq) (*v1alpha1.Rabbitmq, error)
	UpdateStatus(*v1alpha1.Rabbitmq) (*v1alpha1.Rabbitmq, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Rabbitmq, error)
	List(opts v1.ListOptions) (*v1alpha1.RabbitmqList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Rabbitmq, err error)
	RabbitmqExpansion
}

// rabbitmqs implements RabbitmqInterface
type rabbitmqs struct {
	client rest.Interface
	ns     string
}

// newRabbitmqs returns a Rabbitmqs
func newRabbitmqs(c *ExampleV1alpha1Client, namespace string) *rabbitmqs {
	return &rabbitmqs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rabbitmq, and returns the corresponding rabbitmq object, and an error if there is any.
func (c *rabbitmqs) Get(name string, options v1.GetOptions) (result *v1alpha1.Rabbitmq, err error) {
	result = &v1alpha1.Rabbitmq{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rabbitmqs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Rabbitmqs that match those selectors.
func (c *rabbitmqs) List(opts v1.ListOptions) (result *v1alpha1.RabbitmqList, err error) {
	result = &v1alpha1.RabbitmqList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rabbitmqs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rabbitmqs.
func (c *rabbitmqs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rabbitmqs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a rabbitmq and creates it.  Returns the server's representation of the rabbitmq, and an error, if there is any.
func (c *rabbitmqs) Create(rabbitmq *v1alpha1.Rabbitmq) (result *v1alpha1.Rabbitmq, err error) {
	result = &v1alpha1.Rabbitmq{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rabbitmqs").
		Body(rabbitmq).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rabbitmq and updates it. Returns the server's representation of the rabbitmq, and an error, if there is any.
func (c *rabbitmqs) Update(rabbitmq *v1alpha1.Rabbitmq) (result *v1alpha1.Rabbitmq, err error) {
	result = &v1alpha1.Rabbitmq{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rabbitmqs").
		Name(rabbitmq.Name).
		Body(rabbitmq).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *rabbitmqs) UpdateStatus(rabbitmq *v1alpha1.Rabbitmq) (result *v1alpha1.Rabbitmq, err error) {
	result = &v1alpha1.Rabbitmq{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rabbitmqs").
		Name(rabbitmq.Name).
		SubResource("status").
		Body(rabbitmq).
		Do().
		Into(result)
	return
}

// Delete takes name of the rabbitmq and deletes it. Returns an error if one occurs.
func (c *rabbitmqs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rabbitmqs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rabbitmqs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rabbitmqs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched rabbitmq.
func (c *rabbitmqs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Rabbitmq, err error) {
	result = &v1alpha1.Rabbitmq{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rabbitmqs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
