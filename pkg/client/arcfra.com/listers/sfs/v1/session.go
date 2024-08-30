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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/iomesh/sfs-sdk-go/pkg/apis/arcfra.com/sfs/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SessionLister helps list Sessions.
// All objects returned here must be treated as read-only.
type SessionLister interface {
	// List lists all Sessions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Session, err error)
	// Sessions returns an object that can list and get Sessions.
	Sessions(namespace string) SessionNamespaceLister
	SessionListerExpansion
}

// sessionLister implements the SessionLister interface.
type sessionLister struct {
	indexer cache.Indexer
}

// NewSessionLister returns a new SessionLister.
func NewSessionLister(indexer cache.Indexer) SessionLister {
	return &sessionLister{indexer: indexer}
}

// List lists all Sessions in the indexer.
func (s *sessionLister) List(selector labels.Selector) (ret []*v1.Session, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Session))
	})
	return ret, err
}

// Sessions returns an object that can list and get Sessions.
func (s *sessionLister) Sessions(namespace string) SessionNamespaceLister {
	return sessionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SessionNamespaceLister helps list and get Sessions.
// All objects returned here must be treated as read-only.
type SessionNamespaceLister interface {
	// List lists all Sessions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Session, err error)
	// Get retrieves the Session from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Session, error)
	SessionNamespaceListerExpansion
}

// sessionNamespaceLister implements the SessionNamespaceLister
// interface.
type sessionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Sessions in the indexer for a given namespace.
func (s sessionNamespaceLister) List(selector labels.Selector) (ret []*v1.Session, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Session))
	})
	return ret, err
}

// Get retrieves the Session from the indexer for a given namespace and name.
func (s sessionNamespaceLister) Get(name string) (*v1.Session, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("session"), name)
	}
	return obj.(*v1.Session), nil
}
