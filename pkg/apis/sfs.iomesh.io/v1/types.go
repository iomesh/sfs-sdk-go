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

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudProvider is the Schema for the cloudproviders API.
type CloudProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CloudProviderSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudProviderList contains a list of CloudProvider.
type CloudProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudProvider `json:"items"`
}

// Backend storage type.
type StorageType string

const (
	ELF   StorageType = "Elf"
	Iscsi StorageType = "Iscsi"
	Local StorageType = "Local"
)

// CloudProviderSpec defines the desired state of CloudProvider.
type CloudProviderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	Endpoint    string            `json:"endpoint"`             // Cloud provider endpoint address, e.g. "http://xx.xx.xx.xx:xxxx"
	Parameters  map[string]string `json:"parameters,omitempty"` // Backend storage specific parameters
	StorageType StorageType       `json:"storage_type"`         // Backend storage type.
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Cluster is the Schema for the clusters API.
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ClusterList contains a list of Cluster.
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// ClusterSpec defines the desired state of Cluster.
type ClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	NextNsid      int64 `json:"next_nsid"`       // Next available ns id in this cluster
	NextSessionID int64 `json:"next_session_id"` // Next available session id in this cluster
	NextShardid   int64 `json:"next_shardid"`    // Next available shard id in this cluster
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Metric is the Schema for the Metrics API.
type Metric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetricSpec   `json:"spec,omitempty"`
	Status MetricStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MetricList contains a list of Metric.
type MetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Metric `json:"items"`
}

// MetricSpec defines the desired state of Metric.
type MetricSpec struct {
	ClusterID string `json:"cluster_id"` // Cluster ID.
}

// MetricStatus defines the observed state of Metric.
type MetricStatus struct {
	ClusterIOMetric      IOMetric              `json:"cluster_io_metric"`      // Cluster-wide IO metric.
	ClusterStatMetric    StatMetric            `json:"cluster_stat_metric"`    // Cluster-wide statistic metric.
	NamespaceIOMetrics   map[string]IOMetric   `json:"namespace_io_metrics"`   // IO metrics per namespace.
	NamespaceStatMetrics map[string]StatMetric `json:"namespace_stat_metrics"` // Statistic metrics per namespace.
}

type IOMetric struct {
	Read  PerfMetric `json:"read"`  // Read performance metric.
	Write PerfMetric `json:"write"` // Write performance metric.
}

type PerfMetric struct {
	Qps     uint64 `json:"qps"`     // Queries per second.
	Bps     uint64 `json:"bps"`     // Bytes per second.
	Latency uint64 `json:"latency"` // (average) Nanoseconds per query.
}

type StatMetric struct {
	CapacityUsed uint64 `json:"capacity_used"` // Capacity used in bytes.
	TotalFiles   uint64 `json:"total_files"`   // Total files.
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Namespace is the Schema for the namespaces API.
type Namespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceSpec   `json:"spec,omitempty"`
	Status NamespaceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NamespaceList contains a list of Namespace.
type NamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Namespace `json:"items"`
}

// NamespaceSpec defines the desired state of Namespace.
type NamespaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	CloudProvider   string           `json:"cloud_provider"`              // Cloud provider CRD name.
	StoragePolicy   *StoragePolicy   `json:"storage_policy"`              // Storage Policy used to create volumes.
	Protocols       *Protocols       `json:"protocols"`                   // Supported access protocols.
	Capacity        int64            `json:"capacity"`                    // Capacity limit in bytes.
	Mshards         int64            `json:"mshards"`                     // Number of meta shards.
	Dshards         int64            `json:"dshards"`                     // Number of data shards.
	Export          bool             `json:"export"`                      // Whether to export the namespace, i.e. whether mountable.
	Followers       int64            `json:"followers"`                   // How many follower nodes for an individual shard when exporting. e.g. for one shard HA; group, there are `1 + followers` nodes.
	NFSExportConfig *NFSExportConfig `json:"nfs_export_config,omitempty"` // Configs of nfs-ganesha-export that sfs supports. None if it is not given.
	UserDescription *string          `json:"user_description,omitempty"`  // User-defined description.
}

type StoragePolicy struct {
	Replicas      uint32            `json:"replicas"`             // Backend storage replicas.
	ThinProvision bool              `json:"thin_provision"`       // Whether the backend storage is thin-provisioned.
	Parameters    map[string]string `json:"parameters,omitempty"` // Parameters unique to specific backend storage.
}

type Protocols struct {
	SupportNFS  bool `json:"support_nfs"`  // Whether this namespace can be accessed via NFS.
	SupportHDFS bool `json:"support_hdfs"` // Whether this namespace can be accessed via HDFS.
}

type NFSExportConfig struct {
	AccessConfig  AccessConfig  `json:"access_config"`   // Access management
	AnonyUserInfo AnonyUserInfo `json:"anony_user_info"` // Anonymous user info,
	SECType       SECType       `json:"sec_type"`        // User authorization managemennt, Sys or None, default is Sys
	Squash        Squash        `json:"squash"`          // RootSquash, NoRootSquash or AllSquash, default is RootSquash
}

// Access management.
type AccessConfig struct {
	DefaultAccessType    AccessType            `json:"default_access_type"`    // Default access type of this config
	ExceptionAccessRules []ExceptionAccessRule `json:"exception_access_rules"` // Rules that diff with `default_access_type`, so access_type in exception_rules must not be; the same with `default_access_type`.
}

type ExceptionAccessRule struct {
	AccessType AccessType `json:"access_type"` // Access type for the specific host(s)
	Host       string     `json:"host"`        // Host format:  ip, hostname or CIDR(eg 192.168.1.0/24)
}

// Anonymous user info.
type AnonyUserInfo struct {
	GID int64 `json:"gid"` // Group id, range INT32_MIN to UINT32_MAX according to ganesha manual doc
	UID int64 `json:"uid"` // User id, range INT32_MIN to UINT32_MAX according to ganesha manual doc
}

// NamespaceStatus defines the observed state of Namespace.
type NamespaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	DshardIDS    []int64      `json:"dshard_ids"`    // Data shards assigned to this namespace.
	ExportStatus ExportStatus `json:"export_status"` // Export status.
	MshardIDS    []int64      `json:"mshard_ids"`    // Meta shards assigned to this namespace.
	Nsid         int64        `json:"nsid"`          // Cluster wide exclusive ID allocated by the ns_manager.
	State        State        `json:"state"`         // Namespace FSM state.
}

// Default access type of this config.
//
// Access type for the specific host(s).
type AccessType string

const (
	AccessTypeNone AccessType = "None"
	Ro             AccessType = "RO"
	Rw             AccessType = "RW"
)

// User authorization managemennt, Sys or None, default is Sys.
type SECType string

const (
	SECTypeNone SECType = "None"
	Sys         SECType = "Sys"
)

// RootSquash, NoRootSquash or AllSquash, default is RootSquash.
type Squash string

const (
	AllSquash    Squash = "AllSquash"
	NoRootSquash Squash = "NoRootSquash"
	RootSquash   Squash = "RootSquash"
)

// Export status.
type ExportStatus string

const (
	Degraded   ExportStatus = "Degraded"
	Exported   ExportStatus = "Exported"
	Unexported ExportStatus = "Unexported"
)

// Namespace FSM state.
type State string

const (
	Deleting     State = "Deleting"
	Exporting    State = "Exporting"
	Initializing State = "Initializing"
	Staging      State = "Staging"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Node is the Schema for the nodes API.
type Node struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeSpec   `json:"spec,omitempty"`
	Status NodeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NodeList contains a list of Node.
type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Node `json:"items"`
}

// NodeSpec defines the desired state of Node.
type NodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	AgentServer *string           `json:"agent_server"`          // If there is an agent on this node: ip+port.
	DataServer  *string           `json:"data_server"`           // If there is a DS on this node: ip+port. Updated by end users.
	DataShards  map[string]string `json:"data_shards,omitempty"` // Data shards that should be exported on this node. shard_id -> volume name Updated by; shard controller.
	MetaServer  *string           `json:"meta_server"`           // If there is a MDS on this node: ip+port. Updated by end users.
	MetaShards  map[string]string `json:"meta_shards,omitempty"` // Meta shards that should be exported on this node. shard_id -> volume name Updated by; shard controller.
	NodeUUID    *string           `json:"node_uuid"`             // Uuid of corresponding virtual machine. Currently used by elf cloud provider to find; corresponding virtual machine. Shouldn't be modified by manager or agent.
	Online      bool              `json:"online"`                // If the node should be considered as a candidate when placing shards. Updated by end users.
}

// NodeStatus defines the observed state of Node.
type NodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	DataShards map[string]string `json:"data_shards"` // Mounted data shards. shard_id -> volume_name:volume_target_path NB: this indicates the; mount status of a shard on this node, but does not mean if it should be attached or; detached. Only shard leader should be attached.
	DataStatus *bool             `json:"data_status"` // Whether DS is healthy when there is a DS on this node: Data server status updated by this; node agent
	LastActive *string           `json:"last_active"` // Node level status Last heartbeat updated by agent
	MetaShards map[string]string `json:"meta_shards"` // Mounted meta shards. shard_id -> volume_name:volume_target_path NB: this indicates the; mount status of a shard on this node, but does not mean if it should be attached or; detached. Only shard leader should be attached.
	MetaStatus *bool             `json:"meta_status"` // Whether MDS is healthy when there is a MDS on this node: Meta server status updated by; this node agent
	Score      int64             `json:"score"`       // Load of the node.
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RecoveryDatabase is the Schema for the recoverydatabases API.
type RecoveryDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RecoveryDatabaseSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RecoveryDatabaseList contains a list of RecoveryDatabase.
type RecoveryDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecoveryDatabase `json:"items"`
}

// RecoveryDatabaseSpec defines the desired state of RecoveryDatabase.
type RecoveryDatabaseSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	Clids           map[string]string   `json:"clids"`            // Client entry records, clientid -> cid_recov_tag.
	ReclaimingClids []int64             `json:"reclaiming_clids"` // Reclaiming client entry records, clientid -> cid_recov_tag.
	Rfhs            map[string][]string `json:"rfhs"`             // Revoked file handle records, cid_recov_tag -> Vec<revoked file handle>.
	Version         int64               `json:"version"`          // Version number.
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Session is the Schema for the sessions API.
type Session struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SessionSpec   `json:"spec,omitempty"`
	Status SessionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SessionList contains a list of Session.
type SessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Session `json:"items"`
}

// SessionSpec defines the desired state of Session.
type SessionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	ClientAddr   string            `json:"client_addr"`    // Exposed <ip:port>.
	NodeName     string            `json:"node_name"`      // Node associated with the session.
	ProtocolAddr string            `json:"protocol_addr"`  // Exposed <ip:port>.
	Stype        Stype             `json:"stype"`          // Session type, e.g. nfs, fuse and etc.
	Vips         map[string]string `json:"vips,omitempty"` // VIP list assigned to this session.
}

// Session type, e.g. nfs, fuse and etc.
type Stype string

const (
	Fuse Stype = "FUSE"
	NFS  Stype = "NFS"
)

// SessionStatus defines the observed state of Session.
type SessionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	SessionID int64 `json:"session_id"` // Cluster wide exclusive ID.
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Shard is the Schema for the shards API.
type Shard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShardSpec   `json:"spec,omitempty"`
	Status ShardStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ShardList contains a list of Shard.
type ShardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Shard `json:"items"`
}

// ShardSpec defines the desired state of Shard.
type ShardSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	CloudProvider string    `json:"cloud_provider"` // Cloud provider CRD name.
	Export        bool      `json:"export"`         // If this shard is supposed to be export.
	Followers     int64     `json:"followers"`      // How many follower nodes for this shard. e.g. there are `1 + followers` HA group members.
	NSRef         NSRef     `json:"ns_ref"`         // Reference of the ns that this shard belongs.
	Nsid          int64     `json:"nsid"`           // The ns id that this shard belongs.
	ShardID       int64     `json:"shard_id"`       // The shard id.
	ShardType     ShardType `json:"shard_type"`     // Meta or Data.
	Size          int64     `json:"size"`           // Volume size in bytes.
}

// Reference of the ns that this shard belongs.
type NSRef struct {
	NSName      string `json:"ns_name"`
	NSNamespace string `json:"ns_namespace"`
	NSUid       string `json:"ns_uid"`
}

// ShardStatus defines the observed state of Shard.
type ShardStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	Exported   bool     `json:"exported"`    // If this shard is exported.
	HaMembers  []string `json:"ha_members"`  // HA group members, using `node_name` as identity.
	NodeName   *string  `json:"node_name"`   // Which node updated the status.
	VolumeName *string  `json:"volume_name"` // On-premise storage.
}

// Meta or Data.
type ShardType string

const (
	Data ShardType = "Data"
	Meta ShardType = "Meta"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Vip is the Schema for the vips API.
type Vip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VipSpec   `json:"spec,omitempty"`
	Status VipStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VipList contains a list of Vip.
type VipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vip `json:"items"`
}

// VipSpec defines the desired state of Vip.
type VipSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	Addr string `json:"addr"` // The ipv4 addr.
}

// VipStatus defines the observed state of Vip.
type VipStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	Ready       bool    `json:"ready"`        // If this vip is ready.
	SessionName *string `json:"session_name"` // Which session this vip is appointed.
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Volume is the Schema for the volumes API.
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VolumeSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VolumeList contains a list of Volume.
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

// VolumeSpec defines the desired state of Volume.
type VolumeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	CloudProvider string    `json:"cloud_provider"` // Cloud provider name
	NSRef         NSRef     `json:"ns_ref"`         // Reference of the ns that this volume belongs.
	ShardID       int64     `json:"shard_id"`       // Shard that this volume is assigned.
	ShardType     ShardType `json:"shard_type"`     // Shard type, Meta or Data.
	Size          int64     `json:"size"`           // Volume size in bytes
}
