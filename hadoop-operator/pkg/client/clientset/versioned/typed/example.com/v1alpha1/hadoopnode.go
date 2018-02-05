package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	scheme "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HadoopNodesGetter has a method to return a HadoopNodeInterface.
// A group's client should implement this interface.
type HadoopNodesGetter interface {
	HadoopNodes(namespace string) HadoopNodeInterface
}

// HadoopNodeInterface has methods to work with HadoopNode resources.
type HadoopNodeInterface interface {
	Create(*v1alpha1.HadoopNode) (*v1alpha1.HadoopNode, error)
	Update(*v1alpha1.HadoopNode) (*v1alpha1.HadoopNode, error)
	UpdateStatus(*v1alpha1.HadoopNode) (*v1alpha1.HadoopNode, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.HadoopNode, error)
	List(opts v1.ListOptions) (*v1alpha1.HadoopNodeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HadoopNode, err error)
	HadoopNodeExpansion
}

// hadoopNodes implements HadoopNodeInterface
type hadoopNodes struct {
	client rest.Interface
	ns     string
}

// newHadoopNodes returns a HadoopNodes
func newHadoopNodes(c *ExampleV1alpha1Client, namespace string) *hadoopNodes {
	return &hadoopNodes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hadoopNode, and returns the corresponding hadoopNode object, and an error if there is any.
func (c *hadoopNodes) Get(name string, options v1.GetOptions) (result *v1alpha1.HadoopNode, err error) {
	result = &v1alpha1.HadoopNode{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoopnodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HadoopNodes that match those selectors.
func (c *hadoopNodes) List(opts v1.ListOptions) (result *v1alpha1.HadoopNodeList, err error) {
	result = &v1alpha1.HadoopNodeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hadoopnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hadoopNodes.
func (c *hadoopNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hadoopnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a hadoopNode and creates it.  Returns the server's representation of the hadoopNode, and an error, if there is any.
func (c *hadoopNodes) Create(hadoopNode *v1alpha1.HadoopNode) (result *v1alpha1.HadoopNode, err error) {
	result = &v1alpha1.HadoopNode{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hadoopnodes").
		Body(hadoopNode).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hadoopNode and updates it. Returns the server's representation of the hadoopNode, and an error, if there is any.
func (c *hadoopNodes) Update(hadoopNode *v1alpha1.HadoopNode) (result *v1alpha1.HadoopNode, err error) {
	result = &v1alpha1.HadoopNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hadoopnodes").
		Name(hadoopNode.Name).
		Body(hadoopNode).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *hadoopNodes) UpdateStatus(hadoopNode *v1alpha1.HadoopNode) (result *v1alpha1.HadoopNode, err error) {
	result = &v1alpha1.HadoopNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hadoopnodes").
		Name(hadoopNode.Name).
		SubResource("status").
		Body(hadoopNode).
		Do().
		Into(result)
	return
}

// Delete takes name of the hadoopNode and deletes it. Returns an error if one occurs.
func (c *hadoopNodes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoopnodes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hadoopNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hadoopnodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hadoopNode.
func (c *hadoopNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HadoopNode, err error) {
	result = &v1alpha1.HadoopNode{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hadoopnodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
