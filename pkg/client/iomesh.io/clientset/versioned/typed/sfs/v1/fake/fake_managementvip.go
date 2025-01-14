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

	v1 "github.com/iomesh/sfs-sdk-go/pkg/apis/iomesh.io/sfs/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagementVips implements ManagementVipInterface
type FakeManagementVips struct {
	Fake *FakeSfsV1
	ns   string
}

var managementvipsResource = v1.SchemeGroupVersion.WithResource("managementvips")

var managementvipsKind = v1.SchemeGroupVersion.WithKind("ManagementVip")

// Get takes name of the managementVip, and returns the corresponding managementVip object, and an error if there is any.
func (c *FakeManagementVips) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ManagementVip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managementvipsResource, c.ns, name), &v1.ManagementVip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagementVip), err
}

// List takes label and field selectors, and returns the list of ManagementVips that match those selectors.
func (c *FakeManagementVips) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ManagementVipList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managementvipsResource, managementvipsKind, c.ns, opts), &v1.ManagementVipList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ManagementVipList{ListMeta: obj.(*v1.ManagementVipList).ListMeta}
	for _, item := range obj.(*v1.ManagementVipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managementVips.
func (c *FakeManagementVips) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managementvipsResource, c.ns, opts))

}

// Create takes the representation of a managementVip and creates it.  Returns the server's representation of the managementVip, and an error, if there is any.
func (c *FakeManagementVips) Create(ctx context.Context, managementVip *v1.ManagementVip, opts metav1.CreateOptions) (result *v1.ManagementVip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managementvipsResource, c.ns, managementVip), &v1.ManagementVip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagementVip), err
}

// Update takes the representation of a managementVip and updates it. Returns the server's representation of the managementVip, and an error, if there is any.
func (c *FakeManagementVips) Update(ctx context.Context, managementVip *v1.ManagementVip, opts metav1.UpdateOptions) (result *v1.ManagementVip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managementvipsResource, c.ns, managementVip), &v1.ManagementVip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagementVip), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagementVips) UpdateStatus(ctx context.Context, managementVip *v1.ManagementVip, opts metav1.UpdateOptions) (*v1.ManagementVip, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managementvipsResource, "status", c.ns, managementVip), &v1.ManagementVip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagementVip), err
}

// Delete takes name of the managementVip and deletes it. Returns an error if one occurs.
func (c *FakeManagementVips) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(managementvipsResource, c.ns, name, opts), &v1.ManagementVip{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagementVips) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managementvipsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ManagementVipList{})
	return err
}

// Patch applies the patch and returns the patched managementVip.
func (c *FakeManagementVips) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ManagementVip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managementvipsResource, c.ns, name, pt, data, subresources...), &v1.ManagementVip{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagementVip), err
}
