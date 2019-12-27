/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	v2 "bk-bcs/bcs-services/bcs-webhook-server/pkg/apis/bk-bcs/v2"
	scheme "bk-bcs/bcs-services/bcs-webhook-server/pkg/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BcsDbPrivConfigsGetter has a method to return a BcsDbPrivConfigInterface.
// A group's client should implement this interface.
type BcsDbPrivConfigsGetter interface {
	BcsDbPrivConfigs(namespace string) BcsDbPrivConfigInterface
}

// BcsDbPrivConfigInterface has methods to work with BcsDbPrivConfig resources.
type BcsDbPrivConfigInterface interface {
	Create(*v2.BcsDbPrivConfig) (*v2.BcsDbPrivConfig, error)
	Update(*v2.BcsDbPrivConfig) (*v2.BcsDbPrivConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2.BcsDbPrivConfig, error)
	List(opts v1.ListOptions) (*v2.BcsDbPrivConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.BcsDbPrivConfig, err error)
	BcsDbPrivConfigExpansion
}

// bcsDbPrivConfigs implements BcsDbPrivConfigInterface
type bcsDbPrivConfigs struct {
	client rest.Interface
	ns     string
}

// newBcsDbPrivConfigs returns a BcsDbPrivConfigs
func newBcsDbPrivConfigs(c *BkbcsV2Client, namespace string) *bcsDbPrivConfigs {
	return &bcsDbPrivConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bcsDbPrivConfig, and returns the corresponding bcsDbPrivConfig object, and an error if there is any.
func (c *bcsDbPrivConfigs) Get(name string, options v1.GetOptions) (result *v2.BcsDbPrivConfig, err error) {
	result = &v2.BcsDbPrivConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsdbprivconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BcsDbPrivConfigs that match those selectors.
func (c *bcsDbPrivConfigs) List(opts v1.ListOptions) (result *v2.BcsDbPrivConfigList, err error) {
	result = &v2.BcsDbPrivConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsdbprivconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bcsDbPrivConfigs.
func (c *bcsDbPrivConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bcsdbprivconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a bcsDbPrivConfig and creates it.  Returns the server's representation of the bcsDbPrivConfig, and an error, if there is any.
func (c *bcsDbPrivConfigs) Create(bcsDbPrivConfig *v2.BcsDbPrivConfig) (result *v2.BcsDbPrivConfig, err error) {
	result = &v2.BcsDbPrivConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bcsdbprivconfigs").
		Body(bcsDbPrivConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bcsDbPrivConfig and updates it. Returns the server's representation of the bcsDbPrivConfig, and an error, if there is any.
func (c *bcsDbPrivConfigs) Update(bcsDbPrivConfig *v2.BcsDbPrivConfig) (result *v2.BcsDbPrivConfig, err error) {
	result = &v2.BcsDbPrivConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcsdbprivconfigs").
		Name(bcsDbPrivConfig.Name).
		Body(bcsDbPrivConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the bcsDbPrivConfig and deletes it. Returns an error if one occurs.
func (c *bcsDbPrivConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsdbprivconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bcsDbPrivConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsdbprivconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bcsDbPrivConfig.
func (c *bcsDbPrivConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.BcsDbPrivConfig, err error) {
	result = &v2.BcsDbPrivConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bcsdbprivconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
