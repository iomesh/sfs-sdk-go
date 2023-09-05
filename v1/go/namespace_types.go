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

// NamespaceSpec defines the desired state of Namespace
type NamespaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	Capacity        int64             `json:"capacity"`         // Capacity limit in bytes.
	CloudProvider   string            `json:"cloud_provider"`   // Cloud provider CRD name.
	Dshards         int64             `json:"dshards"`          // Number of data shards.
	Export          bool              `json:"export"`           // Whether to export the namespace, i.e. whether mountable.
	Followers       int64             `json:"followers"`        // How many follower nodes for an individual shard when exporting. e.g. for one shard HA; group, there are `1 + followers` nodes.
	Mshards         int64             `json:"mshards"`          // Number of meta shards.
	NFSExportConfig *NFSExportConfig  `json:"nfs_export_config"`// Configs of nfs-ganesha-export that sfs supports. None if it is not given.
	Props           map[string]string `json:"props,omitempty"`  // Optional. Namespace properties.
}

type NFSExportConfig struct {
	AccessConfig  AccessConfig  `json:"access_config"`  // Access management
	AnonyUserInfo AnonyUserInfo `json:"anony_user_info"`// Anonymous user info,
	SECType       SECType       `json:"sec_type"`       // User authorization managemennt, Sys or None, default is Sys
	Squash        Squash        `json:"squash"`         // RootSquash, NoRootSquash or AllSquash, default is RootSquash
}

// Access management
type AccessConfig struct {
	DefaultAccessType    AccessType            `json:"default_access_type"`   // Default access type of this config
	ExceptionAccessRules []ExceptionAccessRule `json:"exception_access_rules"`// Rules that diff with `default_access_type`, so access_type in exception_rules must not be; the same with `default_access_type`.
}

type ExceptionAccessRule struct {
	AccessType AccessType `json:"access_type"`// Access type for the specific host(s)
	Host       string     `json:"host"`       // Host format:  ip, hostname or CIDR(eg 192.168.1.0/24)
}

// Anonymous user info,
type AnonyUserInfo struct {
	Gid int64 `json:"gid"`// Group id, range INT32_MIN to UINT32_MAX according to ganesha manual doc
	Uid int64 `json:"uid"`// User id, range INT32_MIN to UINT32_MAX according to ganesha manual doc
}

// NamespaceStatus defines the observed state of Namespace
type NamespaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	DshardIDS    []int64      `json:"dshard_ids"`   // Data shards assigned to this namespace.
	ExportStatus ExportStatus `json:"export_status"`// Export status.
	MshardIDS    []int64      `json:"mshard_ids"`   // Meta shards assigned to this namespace.
	Nsid         int64        `json:"nsid"`         // Cluster wide exclusive ID allocated by the ns_manager.
	State        State        `json:"state"`        // Namespace FSM state.
}

// Default access type of this config
//
// Access type for the specific host(s)
type AccessType string
const (
	AccessTypeNone AccessType = "None"
	Ro AccessType = "RO"
	Rw AccessType = "RW"
)

// User authorization managemennt, Sys or None, default is Sys
type SECType string
const (
	SECTypeNone SECType = "None"
	Sys SECType = "Sys"
)

// RootSquash, NoRootSquash or AllSquash, default is RootSquash
type Squash string
const (
	AllSquash Squash = "AllSquash"
	NoRootSquash Squash = "NoRootSquash"
	RootSquash Squash = "RootSquash"
)

// Export status.
type ExportStatus string
const (
	Degraded ExportStatus = "Degraded"
	Exported ExportStatus = "Exported"
	Unexported ExportStatus = "Unexported"
)

// Namespace FSM state.
type State string
const (
	Deleting State = "Deleting"
	Exporting State = "Exporting"
	Initializing State = "Initializing"
	Staging State = "Staging"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Namespace is the Schema for the namespaces API
type Namespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceSpec   `json:"spec,omitempty"`
	Status NamespaceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NamespaceList contains a list of Namespace
type NamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Namespace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Namespace{}, &NamespaceList{})
}
