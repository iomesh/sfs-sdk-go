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

// VolumeSpec defines the desired state of Volume
type VolumeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	CloudProvider string    `json:"cloud_provider"`// Cloud provider name
	NSRef         NSRef     `json:"ns_ref"`        // Reference of the ns that this volume belongs.
	ShardID       int64     `json:"shard_id"`      // Shard that this volume is assigned.
	ShardType     ShardType `json:"shard_type"`    // Shard type, Meta or Data.
	Size          int64     `json:"size"`          // Volume size in bytes
}

// VolumeStatus defines the observed state of Volume
type VolumeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Volume is the Schema for the volumes API
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeSpec   `json:"spec,omitempty"`
	Status VolumeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VolumeList contains a list of Volume
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Volume{}, &VolumeList{})
}
