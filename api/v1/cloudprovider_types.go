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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CloudProviderSpec defines the desired state of CloudProvider
type CloudProviderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	Endpoint    string            `json:"endpoint"`            // Cloud provider endpoint address, e.g. "http://xx.xx.xx.xx:xxxx"
	Parameters  map[string]string `json:"parameters,omitempty"`// Backend storage specific parameters
	StorageType StorageType       `json:"storage_type"`        // Backend storage type.
}

// Backend storage type.
type StorageType string
const (
	ELF StorageType = "Elf"
	Iscsi StorageType = "Iscsi"
	Local StorageType = "Local"
)

// CloudProviderStatus defines the observed state of CloudProvider
type CloudProviderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CloudProvider is the Schema for the cloudproviders API
type CloudProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudProviderSpec   `json:"spec,omitempty"`
	Status CloudProviderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CloudProviderList contains a list of CloudProvider
type CloudProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudProvider{}, &CloudProviderList{})
}
