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

// FakeSessions implements SessionInterface
type FakeSessions struct {
	Fake *FakeSfsV1
	ns   string
}

var sessionsResource = v1.SchemeGroupVersion.WithResource("sessions")

var sessionsKind = v1.SchemeGroupVersion.WithKind("Session")

// Get takes name of the session, and returns the corresponding session object, and an error if there is any.
func (c *FakeSessions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Session, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sessionsResource, c.ns, name), &v1.Session{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Session), err
}

// List takes label and field selectors, and returns the list of Sessions that match those selectors.
func (c *FakeSessions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SessionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sessionsResource, sessionsKind, c.ns, opts), &v1.SessionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SessionList{ListMeta: obj.(*v1.SessionList).ListMeta}
	for _, item := range obj.(*v1.SessionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sessions.
func (c *FakeSessions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sessionsResource, c.ns, opts))

}

// Create takes the representation of a session and creates it.  Returns the server's representation of the session, and an error, if there is any.
func (c *FakeSessions) Create(ctx context.Context, session *v1.Session, opts metav1.CreateOptions) (result *v1.Session, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sessionsResource, c.ns, session), &v1.Session{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Session), err
}

// Update takes the representation of a session and updates it. Returns the server's representation of the session, and an error, if there is any.
func (c *FakeSessions) Update(ctx context.Context, session *v1.Session, opts metav1.UpdateOptions) (result *v1.Session, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sessionsResource, c.ns, session), &v1.Session{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Session), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSessions) UpdateStatus(ctx context.Context, session *v1.Session, opts metav1.UpdateOptions) (*v1.Session, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sessionsResource, "status", c.ns, session), &v1.Session{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Session), err
}

// Delete takes name of the session and deletes it. Returns an error if one occurs.
func (c *FakeSessions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(sessionsResource, c.ns, name, opts), &v1.Session{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSessions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sessionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.SessionList{})
	return err
}

// Patch applies the patch and returns the patched session.
func (c *FakeSessions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Session, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sessionsResource, c.ns, name, pt, data, subresources...), &v1.Session{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Session), err
}
