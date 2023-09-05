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

// ShardSpec defines the desired state of Shard
type ShardSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	CloudProvider string    `json:"cloud_provider"`// Cloud provider CRD name.
	Export        bool      `json:"export"`        // If this shard is supposed to be export.
	Followers     int64     `json:"followers"`     // How many follower nodes for this shard. e.g. there are `1 + followers` HA group members.
	NSRef         NSRef     `json:"ns_ref"`        // Reference of the ns that this shard belongs.
	Nsid          int64     `json:"nsid"`          // The ns id that this shard belongs.
	ShardID       int64     `json:"shard_id"`      // The shard id.
	ShardType     ShardType `json:"shard_type"`    // Meta or Data.
	Size          int64     `json:"size"`          // Volume size in bytes.
}

// Reference of the ns that this shard belongs.
type NSRef struct {
	NSName      string `json:"ns_name"`
	NSNamespace string `json:"ns_namespace"`
	NSUid       string `json:"ns_uid"`
}

// ShardStatus defines the observed state of Shard
type ShardStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	Exported   bool     `json:"exported"`   // If this shard is exported.
	HaMembers  []string `json:"ha_members"` // HA group members, using `node_name` as identity.
	NodeName   *string  `json:"node_name"`  // Which node updated the status.
	VolumeName *string  `json:"volume_name"`// On-premise storage.
}

// Meta or Data.
type ShardType string
const (
	Data ShardType = "Data"
	Meta ShardType = "Meta"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Shard is the Schema for the shards API
type Shard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShardSpec   `json:"spec,omitempty"`
	Status ShardStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ShardList contains a list of Shard
type ShardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Shard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Shard{}, &ShardList{})
}
