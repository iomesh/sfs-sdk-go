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

package fake

import (
	"context"

	v1 "github.com/iomesh/sfs-sdk-go/pkg/apis/arcfra.com/sfs/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNamespaces implements NamespaceInterface
type FakeNamespaces struct {
	Fake *FakeSfsV1
	ns   string
}

var namespacesResource = v1.SchemeGroupVersion.WithResource("namespaces")

var namespacesKind = v1.SchemeGroupVersion.WithKind("Namespace")

// Get takes name of the namespace, and returns the corresponding namespace object, and an error if there is any.
func (c *FakeNamespaces) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(namespacesResource, c.ns, name), &v1.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Namespace), err
}

// List takes label and field selectors, and returns the list of Namespaces that match those selectors.
func (c *FakeNamespaces) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(namespacesResource, namespacesKind, c.ns, opts), &v1.NamespaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.NamespaceList{ListMeta: obj.(*v1.NamespaceList).ListMeta}
	for _, item := range obj.(*v1.NamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested namespaces.
func (c *FakeNamespaces) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(namespacesResource, c.ns, opts))

}

// Create takes the representation of a namespace and creates it.  Returns the server's representation of the namespace, and an error, if there is any.
func (c *FakeNamespaces) Create(ctx context.Context, namespace *v1.Namespace, opts metav1.CreateOptions) (result *v1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(namespacesResource, c.ns, namespace), &v1.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Namespace), err
}

// Update takes the representation of a namespace and updates it. Returns the server's representation of the namespace, and an error, if there is any.
func (c *FakeNamespaces) Update(ctx context.Context, namespace *v1.Namespace, opts metav1.UpdateOptions) (result *v1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(namespacesResource, c.ns, namespace), &v1.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Namespace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNamespaces) UpdateStatus(ctx context.Context, namespace *v1.Namespace, opts metav1.UpdateOptions) (*v1.Namespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(namespacesResource, "status", c.ns, namespace), &v1.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Namespace), err
}

// Delete takes name of the namespace and deletes it. Returns an error if one occurs.
func (c *FakeNamespaces) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(namespacesResource, c.ns, name, opts), &v1.Namespace{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNamespaces) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(namespacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.NamespaceList{})
	return err
}

// Patch applies the patch and returns the patched namespace.
func (c *FakeNamespaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(namespacesResource, c.ns, name, pt, data, subresources...), &v1.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Namespace), err
}
