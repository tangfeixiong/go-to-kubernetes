package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/apis/example.com/v1alpha1"
	scheme "github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MysqlsGetter has a method to return a MysqlInterface.
// A group's client should implement this interface.
type MysqlsGetter interface {
	Mysqls(namespace string) MysqlInterface
}

// MysqlInterface has methods to work with Mysql resources.
type MysqlInterface interface {
	Create(*v1alpha1.Mysql) (*v1alpha1.Mysql, error)
	Update(*v1alpha1.Mysql) (*v1alpha1.Mysql, error)
	UpdateStatus(*v1alpha1.Mysql) (*v1alpha1.Mysql, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Mysql, error)
	List(opts v1.ListOptions) (*v1alpha1.MysqlList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Mysql, err error)
	MysqlExpansion
}

// mysqls implements MysqlInterface
type mysqls struct {
	client rest.Interface
	ns     string
}

// newMysqls returns a Mysqls
func newMysqls(c *ExampleV1alpha1Client, namespace string) *mysqls {
	return &mysqls{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mysql, and returns the corresponding mysql object, and an error if there is any.
func (c *mysqls) Get(name string, options v1.GetOptions) (result *v1alpha1.Mysql, err error) {
	result = &v1alpha1.Mysql{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Mysqls that match those selectors.
func (c *mysqls) List(opts v1.ListOptions) (result *v1alpha1.MysqlList, err error) {
	result = &v1alpha1.MysqlList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mysqls.
func (c *mysqls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a mysql and creates it.  Returns the server's representation of the mysql, and an error, if there is any.
func (c *mysqls) Create(mysql *v1alpha1.Mysql) (result *v1alpha1.Mysql, err error) {
	result = &v1alpha1.Mysql{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqls").
		Body(mysql).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mysql and updates it. Returns the server's representation of the mysql, and an error, if there is any.
func (c *mysqls) Update(mysql *v1alpha1.Mysql) (result *v1alpha1.Mysql, err error) {
	result = &v1alpha1.Mysql{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqls").
		Name(mysql.Name).
		Body(mysql).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mysqls) UpdateStatus(mysql *v1alpha1.Mysql) (result *v1alpha1.Mysql, err error) {
	result = &v1alpha1.Mysql{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqls").
		Name(mysql.Name).
		SubResource("status").
		Body(mysql).
		Do().
		Into(result)
	return
}

// Delete takes name of the mysql and deletes it. Returns an error if one occurs.
func (c *mysqls) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mysqls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mysql.
func (c *mysqls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Mysql, err error) {
	result = &v1alpha1.Mysql{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
