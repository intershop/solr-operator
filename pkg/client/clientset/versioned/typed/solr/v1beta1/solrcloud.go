/*
Copyright 2019 Bloomberg Finance LP.

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

package v1beta1

import (
	"time"

	v1beta1 "github.com/bloomberg/solr-operator/pkg/apis/solr/v1beta1"
	scheme "github.com/bloomberg/solr-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SolrCloudsGetter has a method to return a SolrCloudInterface.
// A group's client should implement this interface.
type SolrCloudsGetter interface {
	SolrClouds(namespace string) SolrCloudInterface
}

// SolrCloudInterface has methods to work with SolrCloud resources.
type SolrCloudInterface interface {
	Create(*v1beta1.SolrCloud) (*v1beta1.SolrCloud, error)
	Update(*v1beta1.SolrCloud) (*v1beta1.SolrCloud, error)
	UpdateStatus(*v1beta1.SolrCloud) (*v1beta1.SolrCloud, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.SolrCloud, error)
	List(opts v1.ListOptions) (*v1beta1.SolrCloudList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.SolrCloud, err error)
	SolrCloudExpansion
}

// solrClouds implements SolrCloudInterface
type solrClouds struct {
	client rest.Interface
	ns     string
}

// newSolrClouds returns a SolrClouds
func newSolrClouds(c *SolrV1beta1Client, namespace string) *solrClouds {
	return &solrClouds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the solrCloud, and returns the corresponding solrCloud object, and an error if there is any.
func (c *solrClouds) Get(name string, options v1.GetOptions) (result *v1beta1.SolrCloud, err error) {
	result = &v1beta1.SolrCloud{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("solrclouds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SolrClouds that match those selectors.
func (c *solrClouds) List(opts v1.ListOptions) (result *v1beta1.SolrCloudList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SolrCloudList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("solrclouds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested solrClouds.
func (c *solrClouds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("solrclouds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a solrCloud and creates it.  Returns the server's representation of the solrCloud, and an error, if there is any.
func (c *solrClouds) Create(solrCloud *v1beta1.SolrCloud) (result *v1beta1.SolrCloud, err error) {
	result = &v1beta1.SolrCloud{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("solrclouds").
		Body(solrCloud).
		Do().
		Into(result)
	return
}

// Update takes the representation of a solrCloud and updates it. Returns the server's representation of the solrCloud, and an error, if there is any.
func (c *solrClouds) Update(solrCloud *v1beta1.SolrCloud) (result *v1beta1.SolrCloud, err error) {
	result = &v1beta1.SolrCloud{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("solrclouds").
		Name(solrCloud.Name).
		Body(solrCloud).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *solrClouds) UpdateStatus(solrCloud *v1beta1.SolrCloud) (result *v1beta1.SolrCloud, err error) {
	result = &v1beta1.SolrCloud{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("solrclouds").
		Name(solrCloud.Name).
		SubResource("status").
		Body(solrCloud).
		Do().
		Into(result)
	return
}

// Delete takes name of the solrCloud and deletes it. Returns an error if one occurs.
func (c *solrClouds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("solrclouds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *solrClouds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("solrclouds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched solrCloud.
func (c *solrClouds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.SolrCloud, err error) {
	result = &v1beta1.SolrCloud{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("solrclouds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
