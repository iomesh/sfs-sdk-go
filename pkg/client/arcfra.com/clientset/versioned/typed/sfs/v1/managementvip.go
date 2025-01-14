/*
Copyright 2023 The IOMesh Authors.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/iomesh/sfs-sdk-go/pkg/apis/arcfra.com/sfs/v1"
	scheme "github.com/iomesh/sfs-sdk-go/pkg/client/arcfra.com/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ManagementVipsGetter has a method to return a ManagementVipInterface.
// A group's client should implement this interface.
type ManagementVipsGetter interface {
	ManagementVips(namespace string) ManagementVipInterface
}

// ManagementVipInterface has methods to work with ManagementVip resources.
type ManagementVipInterface interface {
	Create(ctx context.Context, managementVip *v1.ManagementVip, opts metav1.CreateOptions) (*v1.ManagementVip, error)
	Update(ctx context.Context, managementVip *v1.ManagementVip, opts metav1.UpdateOptions) (*v1.ManagementVip, error)
	UpdateStatus(ctx context.Context, managementVip *v1.ManagementVip, opts metav1.UpdateOptions) (*v1.ManagementVip, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ManagementVip, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ManagementVipList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ManagementVip, err error)
	ManagementVipExpansion
}

// managementVips implements ManagementVipInterface
type managementVips struct {
	client rest.Interface
	ns     string
}

// newManagementVips returns a ManagementVips
func newManagementVips(c *SfsV1Client, namespace string) *managementVips {
	return &managementVips{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the managementVip, and returns the corresponding managementVip object, and an error if there is any.
func (c *managementVips) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ManagementVip, err error) {
	result = &v1.ManagementVip{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managementvips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManagementVips that match those selectors.
func (c *managementVips) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ManagementVipList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ManagementVipList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managementvips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested managementVips.
func (c *managementVips) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("managementvips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a managementVip and creates it.  Returns the server's representation of the managementVip, and an error, if there is any.
func (c *managementVips) Create(ctx context.Context, managementVip *v1.ManagementVip, opts metav1.CreateOptions) (result *v1.ManagementVip, err error) {
	result = &v1.ManagementVip{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("managementvips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managementVip).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a managementVip and updates it. Returns the server's representation of the managementVip, and an error, if there is any.
func (c *managementVips) Update(ctx context.Context, managementVip *v1.ManagementVip, opts metav1.UpdateOptions) (result *v1.ManagementVip, err error) {
	result = &v1.ManagementVip{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managementvips").
		Name(managementVip.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managementVip).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *managementVips) UpdateStatus(ctx context.Context, managementVip *v1.ManagementVip, opts metav1.UpdateOptions) (result *v1.ManagementVip, err error) {
	result = &v1.ManagementVip{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managementvips").
		Name(managementVip.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managementVip).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the managementVip and deletes it. Returns an error if one occurs.
func (c *managementVips) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managementvips").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *managementVips) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managementvips").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched managementVip.
func (c *managementVips) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ManagementVip, err error) {
	result = &v1.ManagementVip{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("managementvips").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
