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

// CloudProviderSpec defines the desired state of CloudProvider.
type CloudProviderSpec struct {
	Endpoint    string            `json:"endpoint"`             // Cloud provider endpoint address, e.g. "http://xx.xx.xx.xx:xxxx"
	Parameters  map[string]string `json:"parameters,omitempty"` // Backend storage specific parameters
	StorageType StorageType       `json:"storage_type"`         // Backend storage type.
}

// Backend storage type.
type StorageType string

const (
	StorageTypeELF   StorageType = "Elf"
	StorageTypeIscsi StorageType = "Iscsi"
	StorageTypeLocal StorageType = "Local"
)

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
	ClusterID     string `json:"cluster_id,omitempty"` // Cluster id, updated by sfs-operator
	Capacity      int64  `json:"capacity"`             // Cluster capacity in bytes, updated by sfs-operator
	NextNsid      int64  `json:"next_nsid"`            // Next available ns id in this cluster
	NextSessionID int64  `json:"next_session_id"`      // Next available session id in this cluster
	NextShardid   int64  `json:"next_shardid"`         // Next available shard id in this cluster
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// License is the Schema for the License API.
type License struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LicenseSpec   `json:"spec,omitempty"`
	Status LicenseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LicenseList contains a list of License.
type LicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []License `json:"items"`
}

// LicenseSpec defines the desired state of License.
type LicenseSpec struct {
	// The raw license code provided by SmartX. `None` means trial license.
	Code *string `json:"code"`
}

// LicenseStatus defines the observed state of License.
type LicenseStatus struct {
	// The license information.
	Info *LicenseInfo `json:"info"`
	// The total used capacity of SFS.
	UsedCapacity uint64 `json:"used_capacity"`
	// The license code is verified.
	Valid bool `json:"valid"`
}

type LicenseInfo struct {
	// The license type.
	LicenseType LicenseType `json:"license_type"`
	// The maximum allowed capacity.
	MaxCapacity uint64 `json:"max_capacity"`
	// The license valid period.
	Period metav1.Duration `json:"period"`
	// The serial number of the product instance to which this license can apply.
	SerialNumber string `json:"serial_number"`
	// The license signed date.
	SignDate metav1.Time `json:"sign_date"`
	// The product edition.
	SoftwareEdition SoftwareEdition `json:"software_edition"`
}

// The license type.
type LicenseType string

const (
	LicenseTypePerpetual    LicenseType = "Perpetual"
	LicenseTypeSubscription LicenseType = "Subscription"
	LicenseTypeTrial        LicenseType = "Trial"
)

// The product edition.
type SoftwareEdition string

const (
	SoftwareEditionEnterprise SoftwareEdition = "Enterprise"
	SoftwareEditionStandard   SoftwareEdition = "Standard"
)

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
	ClusterIOMetric      IOMetric              `json:"cluster_io_metric"`                // Cluster-wide IO metric.
	ClusterStatMetric    StatMetric            `json:"cluster_stat_metric"`              // Cluster-wide statistic metric.
	NamespaceIOMetrics   map[string]IOMetric   `json:"namespace_io_metrics,omitempty"`   // IO metrics per namespace.
	NamespaceStatMetrics map[string]StatMetric `json:"namespace_stat_metrics,omitempty"` // Statistic metrics per namespace.
}

type IOMetric struct {
	Read  PerfMetric `json:"read"`  // Read performance metric.
	Write PerfMetric `json:"write"` // Write performance metric.
}

type PerfMetric struct {
	QPS     uint64 `json:"qps"`     // Queries per second.
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
	CloudProvider   string           `json:"cloud_provider"`    // Cloud provider CRD name.
	StoragePolicy   StoragePolicy    `json:"storage_policy"`    // Storage Policy used to create volumes.
	Protocols       Protocols        `json:"protocols"`         // Supported access protocols.
	Capacity        int64            `json:"capacity"`          // Capacity limit in bytes.
	Mshards         int64            `json:"mshards"`           // Number of meta shards.
	Dshards         int64            `json:"dshards"`           // Number of data shards.
	Export          bool             `json:"export"`            // Whether to export the namespace, i.e. whether mountable.
	Followers       int64            `json:"followers"`         // How many follower nodes for an individual shard when exporting. e.g. for one shard HA; group, there are `1 + followers` nodes.
	NFSExportConfig *NFSExportConfig `json:"nfs_export_config"` // Configs of nfs-ganesha-export that sfs supports. It takes effect if NFS is supported and if none, default values are used.
	UserDescription string           `json:"user_description"`  // User-defined description.
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
	SecType           SecType           `json:"sec_type"`        // User authorization managemennt.
	Squash            SquashType        `json:"squash"`          // RootSquash, NoRootSquash or AllSquash.
	AnonymousUserInfo AnonymousUserInfo `json:"anony_user_info"` // Anonymous user info.
	AccessConfig      AccessConfig      `json:"access_config"`   // Access management.
}

// User authorization managemennt.
type SecType string

const (
	SecTypeNone SecType = "None"
	SecTypeSys  SecType = "Sys"
)

// RootSquash, NoRootSquash or AllSquash, default is RootSquash.
type SquashType string

const (
	SquashTypeAllSquash    SquashType = "AllSquash"
	SquashTypeNoRootSquash SquashType = "NoRootSquash"
	SquashTypeRootSquash   SquashType = "RootSquash"
)

// Anonymous user info.
type AnonymousUserInfo struct {
	GID int64 `json:"gid"` // Group id, range INT32_MIN to UINT32_MAX according to ganesha manual doc.
	UID int64 `json:"uid"` // User id, range INT32_MIN to UINT32_MAX according to ganesha manual doc.
}

// Access management.
type AccessConfig struct {
	DefaultAccessType  AccessType   `json:"default_access_type"`            // Default access type.
	SpecialAccessRules []AccessRule `json:"special_access_rules,omitempty"` // Rules that diff with `default_access_type`, so access_type in special rules must not be; the same with `default_access_type`.
}

type AccessRule struct {
	AccessType AccessType `json:"access_type"` // Access type for the specific host(s).
	// Client list entries can take on one of the following forms: Match any client:
	//
	// - @name       Netgroup name.
	//
	// - x.x.x.x/y   IPv4 network address.
	//
	// - wildcarded  If the string contains at least one ? or *
	//             character (and is not simply "*"), the string is
	//             used to pattern match host names. Note that [] may
	//             also be used, but the pattern MUST have at least one
	//             ? or *.
	//
	// - hostname    Match a single client (match is by IP address, all
	//             addresses returned by getaddrinfo will match, the
	//             getaddrinfo call is made at config parsing time).
	//
	// - IP address  Match a single client.
	Clients string `json:"clients"`
}

// Access type for the specific host(s).
type AccessType string

const (
	AccessTypeNone AccessType = "None"
	AccessTypeRO   AccessType = "RO"
	AccessTypeRW   AccessType = "RW"
)

// NamespaceStatus defines the observed state of Namespace.
type NamespaceStatus struct {
	DshardIDS    []int64        `json:"dshard_ids,omitempty"` // Data shards assigned to this namespace.
	ExportStatus NsExportStatus `json:"export_status"`        // Export status.
	MshardIDS    []int64        `json:"mshard_ids,omitempty"` // Meta shards assigned to this namespace.
	Nsid         int64          `json:"nsid"`                 // Cluster wide exclusive ID allocated by the ns_manager.
	State        NsState        `json:"state"`                // Namespace FSM state.
}

// Export status.
type NsExportStatus string

const (
	NsExportStatusUnexported NsExportStatus = "Unexported"
	NsExportStatusDegraded   NsExportStatus = "Degraded"
	NsExportStatusExported   NsExportStatus = "Exported"
	NsExportStatusAbnormal   NsExportStatus = "Abnormal"
)

// Namespace FSM state.
type NsState string

const (
	NsStateInitializing NsState = "Initializing"
	NsStateStaging      NsState = "Staging"
	NsStateExporting    NsState = "Exporting"
	NsStateDeleting     NsState = "Deleting"
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
	DataServer *string           `json:"data_server"`           // If there is a DS on this node: ip+port. Updated by end users.
	DataShards map[string]string `json:"data_shards,omitempty"` // Data shards that should be exported on this node. shard_id -> volume name Updated by; shard controller.
	MetaServer *string           `json:"meta_server"`           // If there is a MDS on this node: ip+port. Updated by end users.
	MetaShards map[string]string `json:"meta_shards,omitempty"` // Meta shards that should be exported on this node. shard_id -> volume name Updated by; shard controller.
	NodeUUID   *string           `json:"node_uuid"`             // Uuid of corresponding virtual machine. Currently used by elf cloud provider to find; corresponding virtual machine. Shouldn't be modified by manager or agent.
	Online     bool              `json:"online"`                // If the node should be considered as a candidate when placing shards. Updated by end users.
}

// NodeStatus defines the observed state of Node.
type NodeStatus struct {
	DataShards map[string]string `json:"data_shards,omitempty"` // Mounted data shards. shard_id -> volume_name:volume_target_path NB: this indicates the; mount status of a shard on this node, but does not mean if it should be attached or; detached. Only shard leader should be attached.
	DataStatus *bool             `json:"data_status"`           // Whether DS is healthy when there is a DS on this node: Data server status updated by this; node agent
	LastActive *string           `json:"last_active"`           // Node level status Last heartbeat updated by agent
	MetaShards map[string]string `json:"meta_shards,omitempty"` // Mounted meta shards. shard_id -> volume_name:volume_target_path NB: this indicates the; mount status of a shard on this node, but does not mean if it should be attached or; detached. Only shard leader should be attached.
	MetaStatus *bool             `json:"meta_status"`           // Whether MDS is healthy when there is a MDS on this node: Meta server status updated by; this node agent
	Score      int64             `json:"score"`                 // Load of the node.
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeHealth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeHealthSpec   `json:"spec,omitempty"`
	Status NodeHealthStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeHealthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeHealth `json:"items"`
}

type NodeHealthSpec struct {
	NodeName string `json:"node_name"`
}

type NodeHealthStatus struct {
	DataService *bool   `json:"data_service"`
	MetaService *bool   `json:"meta_service"`
	RenewTime   *string `json:"renew_time"`
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
	Clids           map[string]string   `json:"clids,omitempty"`            // Client entry records, clientid -> cid_recov_tag.
	ReclaimingClids []int64             `json:"reclaiming_clids,omitempty"` // Reclaiming client entry records, clientid -> cid_recov_tag.
	Rfhs            map[string][]string `json:"rfhs,omitempty"`             // Revoked file handle records, cid_recov_tag -> Vec<revoked file handle>. NOTE: `[]string` must be initialized instead of `nil`.
	Version         int64               `json:"version"`                    // Version number.
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
	NodeName     string            `json:"node_name"`      // Node associated with the session.
	SessionType  SessionType       `json:"stype"`          // Session type, e.g. nfs, fuse and etc.
	ProtocolAddr string            `json:"protocol_addr"`  // Exposed <ip:port>.
	ClientAddr   string            `json:"client_addr"`    // Exposed <ip:port>.
	Vips         map[string]string `json:"vips,omitempty"` // VIP list assigned to this session.
	// MAC address of the network interface to which vips should be attached. If None, vips will
	// be attached to the best matching network interface, i.e. the one with the longest netmask
	// prefix length.
	MACAddr []int64 `json:"mac_addr,omitempty"`
	// Slash-notation of netmask, i.e. prefix_len.
	Netmask *int64 `json:"netmask,omitempty"`
	// Namespace_ids that need be exported by this NFS session, updated by namespace controller
	// in sfs-manager. Hash of nfs_export_config is used by sfs-manager to quickly deliver
	// change of nfs_export_config to ganesha.
	NFSNeedExportNSS map[string]int64 `json:"nfs_need_export_nss,omitempty"`
}

// Session type, e.g. nfs, fuse and etc.
type SessionType string

const (
	SessionTypeFuse SessionType = "FUSE"
	SessionTypeNFS  SessionType = "NFS"
)

// SessionStatus defines the observed state of Session.
type SessionStatus struct {
	SessionID      int64               `json:"session_id"`                 // Cluster wide exclusive ID.
	NFSExportPaths map[string][]string `json:"nfs_export_paths,omitempty"` // Ns_name -> export_paths, updated by session controller in sfs-agent.
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
	ShardID       int64       `json:"shard_id"`       // The shard id.
	ShardType     ShardType   `json:"shard_type"`     // Meta or Data.
	CloudProvider string      `json:"cloud_provider"` // Cloud provider CRD name.
	Size          int64       `json:"size"`           // Volume size in bytes.
	NsReference   NsReference `json:"ns_ref"`         // Reference of the ns that this shard belongs.
	Nsid          int64       `json:"nsid"`           // The ns id that this shard belongs.
	Export        bool        `json:"export"`         // If this shard is supposed to be export.
	Followers     int64       `json:"followers"`      // How many follower nodes for this shard. e.g. there are `1 + followers` HA group members.
}

// Meta or Data.
type ShardType string

const (
	ShardTypeData ShardType = "Data"
	ShardTypeMeta ShardType = "Meta"
)

// Reference of the ns that this shard belongs.
type NsReference struct {
	NsName      string `json:"ns_name"`
	NsNamespace string `json:"ns_namespace"`
	NsUID       string `json:"ns_uid"`
}

// ShardStatus defines the observed state of Shard.
type ShardStatus struct {
	VolumeName    *string  `json:"volume_name"`    // On-premise storage.
	VolumeCreated bool     `json:"volume_created"` // Backend volume is created.
	HAMembers     []string `json:"ha_members"`     // HA group members, using `node_name` as identity.
	NodeName      *string  `json:"node_name"`      // Which node updated the status.
	Exported      bool     `json:"exported"`       // If this shard is exported.
}

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
	Addr string `json:"addr"` // The ipv4 addr.
}

// VipStatus defines the observed state of Vip.
type VipStatus struct {
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
	CloudProvider string      `json:"cloud_provider"` // Cloud provider name
	NsReference   NsReference `json:"ns_ref"`         // Reference of the ns that this volume belongs.
	ShardID       int64       `json:"shard_id"`       // Shard that this volume is assigned.
	ShardType     ShardType   `json:"shard_type"`     // Shard type, Meta or Data.
	Size          int64       `json:"size"`           // Volume size in bytes
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Route is the Schema for the routes API.
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RouteSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RouteList contains a list of Route.
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// RouteSpec defines the desired state of Route.
type RouteSpec struct {
	Routes []RouteInfo `json:"routes"`
}

type RouteInfo struct {
	// The destination network.
	Destination string `json:"destination"`
	// The gateway address.
	Gateway string `json:"gateway"`
	// Netmask for the destination net.
	Genmask int64 `json:"genmask"`
}
