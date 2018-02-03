package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/apis/example.com/v1alpha1"
	scheme "github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AmqpsGetter has a method to return a AmqpInterface.
// A group's client should implement this interface.
type AmqpsGetter interface {
	Amqps(namespace string) AmqpInterface
}

// AmqpInterface has methods to work with Amqp resources.
type AmqpInterface interface {
	Create(*v1alpha1.Amqp) (*v1alpha1.Amqp, error)
	Update(*v1alpha1.Amqp) (*v1alpha1.Amqp, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Amqp, error)
	List(opts v1.ListOptions) (*v1alpha1.AmqpList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Amqp, err error)
	AmqpExpansion
}

// amqps implements AmqpInterface
type amqps struct {
	client rest.Interface
	ns     string
}

// newAmqps returns a Amqps
func newAmqps(c *ExampleV1alpha1Client, namespace string) *amqps {
	return &amqps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the amqp, and returns the corresponding amqp object, and an error if there is any.
func (c *amqps) Get(name string, options v1.GetOptions) (result *v1alpha1.Amqp, err error) {
	result = &v1alpha1.Amqp{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("amqps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Amqps that match those selectors.
func (c *amqps) List(opts v1.ListOptions) (result *v1alpha1.AmqpList, err error) {
	result = &v1alpha1.AmqpList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("amqps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested amqps.
func (c *amqps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("amqps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a amqp and creates it.  Returns the server's representation of the amqp, and an error, if there is any.
func (c *amqps) Create(amqp *v1alpha1.Amqp) (result *v1alpha1.Amqp, err error) {
	result = &v1alpha1.Amqp{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("amqps").
		Body(amqp).
		Do().
		Into(result)
	return
}

// Update takes the representation of a amqp and updates it. Returns the server's representation of the amqp, and an error, if there is any.
func (c *amqps) Update(amqp *v1alpha1.Amqp) (result *v1alpha1.Amqp, err error) {
	result = &v1alpha1.Amqp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("amqps").
		Name(amqp.Name).
		Body(amqp).
		Do().
		Into(result)
	return
}

// Delete takes name of the amqp and deletes it. Returns an error if one occurs.
func (c *amqps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("amqps").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *amqps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("amqps").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched amqp.
func (c *amqps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Amqp, err error) {
	result = &v1alpha1.Amqp{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("amqps").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
