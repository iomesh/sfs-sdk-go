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

// NodeSpec defines the desired state of Node
type NodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	AgentServer *string           `json:"agent_server"`         // If there is an agent on this node: ip+port.
	DataServer  *string           `json:"data_server"`          // If there is a DS on this node: ip+port. Updated by end users.
	DataShards  map[string]string `json:"data_shards,omitempty"`// Data shards that should be exported on this node. shard_id -> volume name Updated by; shard controller.
	MetaServer  *string           `json:"meta_server"`          // If there is a MDS on this node: ip+port. Updated by end users.
	MetaShards  map[string]string `json:"meta_shards,omitempty"`// Meta shards that should be exported on this node. shard_id -> volume name Updated by; shard controller.
	NodeUUID    *string           `json:"node_uuid"`            // Uuid of corresponding virtual machine. Currently used by elf cloud provider to find; corresponding virtual machine. Shouldn't be modified by manager or agent.
	Online      bool              `json:"online"`               // If the node should be considered as a candidate when placing shards. Updated by end users.
}

// NodeStatus defines the observed state of Node
type NodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	DataShards map[string]string `json:"data_shards"`// Mounted data shards. shard_id -> volume_name:volume_target_path NB: this indicates the; mount status of a shard on this node, but does not mean if it should be attached or; detached. Only shard leader should be attached.
	DataStatus *bool             `json:"data_status"`// Whether DS is healthy when there is a DS on this node: Data server status updated by this; node agent
	LastActive *string           `json:"last_active"`// Node level status Last heartbeat updated by agent
	MetaShards map[string]string `json:"meta_shards"`// Mounted meta shards. shard_id -> volume_name:volume_target_path NB: this indicates the; mount status of a shard on this node, but does not mean if it should be attached or; detached. Only shard leader should be attached.
	MetaStatus *bool             `json:"meta_status"`// Whether MDS is healthy when there is a MDS on this node: Meta server status updated by; this node agent
	Score      int64             `json:"score"`      // Load of the node.
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Node is the Schema for the nodes API
type Node struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeSpec   `json:"spec,omitempty"`
	Status NodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NodeList contains a list of Node
type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Node `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Node{}, &NodeList{})
}
