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

// MetricLister helps list Metrics.
// All objects returned here must be treated as read-only.
type MetricLister interface {
	// List lists all Metrics in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Metric, err error)
	// Metrics returns an object that can list and get Metrics.
	Metrics(namespace string) MetricNamespaceLister
	MetricListerExpansion
}

// metricLister implements the MetricLister interface.
type metricLister struct {
	indexer cache.Indexer
}

// NewMetricLister returns a new MetricLister.
func NewMetricLister(indexer cache.Indexer) MetricLister {
	return &metricLister{indexer: indexer}
}

// List lists all Metrics in the indexer.
func (s *metricLister) List(selector labels.Selector) (ret []*v1.Metric, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Metric))
	})
	return ret, err
}

// Metrics returns an object that can list and get Metrics.
func (s *metricLister) Metrics(namespace string) MetricNamespaceLister {
	return metricNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MetricNamespaceLister helps list and get Metrics.
// All objects returned here must be treated as read-only.
type MetricNamespaceLister interface {
	// List lists all Metrics in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Metric, err error)
	// Get retrieves the Metric from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Metric, error)
	MetricNamespaceListerExpansion
}

// metricNamespaceLister implements the MetricNamespaceLister
// interface.
type metricNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Metrics in the indexer for a given namespace.
func (s metricNamespaceLister) List(selector labels.Selector) (ret []*v1.Metric, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Metric))
	})
	return ret, err
}

// Get retrieves the Metric from the indexer for a given namespace and name.
func (s metricNamespaceLister) Get(name string) (*v1.Metric, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("metric"), name)
	}
	return obj.(*v1.Metric), nil
}
