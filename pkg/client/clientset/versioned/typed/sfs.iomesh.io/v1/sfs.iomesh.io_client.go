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
	"net/http"

	v1 "github.com/iomesh/sfs-sdk-go/pkg/apis/sfs.iomesh.io/v1"
	"github.com/iomesh/sfs-sdk-go/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type SfsV1Interface interface {
	RESTClient() rest.Interface
	CloudProvidersGetter
	ClustersGetter
	MetricsGetter
	NamespacesGetter
	NodesGetter
	NodeHealthsGetter
	RecoveryDatabasesGetter
	SessionsGetter
	ShardsGetter
	VipsGetter
	VolumesGetter
}

// SfsV1Client is used to interact with features provided by the sfs.iomesh.io group.
type SfsV1Client struct {
	restClient rest.Interface
}

func (c *SfsV1Client) CloudProviders(namespace string) CloudProviderInterface {
	return newCloudProviders(c, namespace)
}

func (c *SfsV1Client) Clusters(namespace string) ClusterInterface {
	return newClusters(c, namespace)
}

func (c *SfsV1Client) Metrics(namespace string) MetricInterface {
	return newMetrics(c, namespace)
}

func (c *SfsV1Client) Namespaces(namespace string) NamespaceInterface {
	return newNamespaces(c, namespace)
}

func (c *SfsV1Client) Nodes(namespace string) NodeInterface {
	return newNodes(c, namespace)
}

func (c *SfsV1Client) NodeHealths(namespace string) NodeHealthInterface {
	return newNodeHealths(c, namespace)
}

func (c *SfsV1Client) RecoveryDatabases(namespace string) RecoveryDatabaseInterface {
	return newRecoveryDatabases(c, namespace)
}

func (c *SfsV1Client) Sessions(namespace string) SessionInterface {
	return newSessions(c, namespace)
}

func (c *SfsV1Client) Shards(namespace string) ShardInterface {
	return newShards(c, namespace)
}

func (c *SfsV1Client) Vips(namespace string) VipInterface {
	return newVips(c, namespace)
}

func (c *SfsV1Client) Volumes(namespace string) VolumeInterface {
	return newVolumes(c, namespace)
}

// NewForConfig creates a new SfsV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*SfsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new SfsV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*SfsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &SfsV1Client{client}, nil
}

// NewForConfigOrDie creates a new SfsV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SfsV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SfsV1Client for the given RESTClient.
func New(c rest.Interface) *SfsV1Client {
	return &SfsV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SfsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
