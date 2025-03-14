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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	afsarcfracom "github.com/iomesh/sfs-sdk-go/pkg/apis/arcfra.com/sfs"
)

const (
	CustomResourceGroup = afsarcfracom.CustomResourceGroupName
	Version             = "v1"
)

var SchemeGroupVersion = schema.GroupVersion{Group: afsarcfracom.CustomResourceGroupName, Version: Version}

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	// SchemeBuilder initializes a scheme builder.
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme is a global function that registers this API group & version to a scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&CloudProvider{},
		&CloudProviderList{},
		&Cluster{},
		&ClusterList{},
		&License{},
		&LicenseList{},
		&Namespace{},
		&NamespaceList{},
		&Node{},
		&NodeList{},
		&NodeHealth{},
		&NodeHealthList{},
		&RecoveryDatabase{},
		&RecoveryDatabaseList{},
		&Session{},
		&SessionList{},
		&Shard{},
		&ShardList{},
		&Vip{},
		&VipList{},
		&Volume{},
		&VolumeList{},
		&Metric{},
		&MetricList{},
		&Route{},
		&RouteList{},
		&ManagementVip{},
		&ManagementVipList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
