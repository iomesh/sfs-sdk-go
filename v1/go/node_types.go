// io.iomesh.sfs.v1.NodeList, NodeList is a list of Node
type IoIomeshSfsV1NodeList struct {
	APIVersion *string                                 `json:"apiVersion,omitempty"`// APIVersion defines the versioned schema of this representation of an object. Servers; should convert recognized schemas to the latest internal value, and may reject; unrecognized values. More info:; https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Items      []IoIomeshSfsV1Node                     `json:"items"`               // List of nodes. More info:; https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md
	Kind       *string                                 `json:"kind,omitempty"`      // Kind is a string value representing the REST resource this object represents. Servers may; infer this from the endpoint the client submits requests to. Cannot be updated. In; CamelCase. More info:; https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata   *IoK8SApimachineryPkgApisMetaV1ListMeta `json:"metadata,omitempty"`
}

// io.iomesh.sfs.v1.Node, Auto-generated derived type for NodeSpec via `CustomResource`
type IoIomeshSfsV1Node struct {
	APIVersion *string                                     `json:"apiVersion,omitempty"`// APIVersion defines the versioned schema of this representation of an object. Servers; should convert recognized schemas to the latest internal value, and may reject; unrecognized values. More info:; https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	Kind       *string                                     `json:"kind,omitempty"`      // Kind is a string value representing the REST resource this object represents. Servers may; infer this from the endpoint the client submits requests to. Cannot be updated. In; CamelCase. More info:; https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata   *IoK8SApimachineryPkgApisMetaV1ObjectMetaV2 `json:"metadata,omitempty"`
	Spec       Spec                                        `json:"spec"`
	Status     *Status                                     `json:"status"`
}

// io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta_v2, ObjectMeta is metadata that all
// persisted resources must have, which includes all objects users must create.
type IoK8SApimachineryPkgApisMetaV1ObjectMetaV2 struct {
	Annotations                map[string]string                                  `json:"annotations,omitempty"`               // Annotations is an unstructured key value map stored with a resource that may be set by; external tools to store and retrieve arbitrary metadata. They are not queryable and; should be preserved when modifying objects. More info:; http://kubernetes.io/docs/user-guide/annotations
	ClusterName                *string                                            `json:"clusterName,omitempty"`               // Deprecated: ClusterName is a legacy field that was always cleared by the system and never; used; it will be removed completely in 1.25.; ; The name in the go struct is changed to help clients detect accidental use.
	CreationTimestamp          *string                                            `json:"creationTimestamp,omitempty"`
	DeletionGracePeriodSeconds *int64                                             `json:"deletionGracePeriodSeconds,omitempty"`// Number of seconds allowed for this object to gracefully terminate before it will be; removed from the system. Only set when deletionTimestamp is also set. May only be; shortened. Read-only.
	DeletionTimestamp          *string                                            `json:"deletionTimestamp,omitempty"`
	Finalizers                 []string                                           `json:"finalizers,omitempty"`                // Must be empty before the object is deleted from the registry. Each entry is an identifier; for the responsible component that will remove the entry from the list. If the; deletionTimestamp of the object is non-nil, entries in this list can only be removed.; Finalizers may be processed and removed in any order.  Order is NOT enforced because it; introduces significant risk of stuck finalizers. finalizers is a shared field, any actor; with permission can reorder it. If the finalizer list is processed in order, then this; can lead to a situation in which the component responsible for the first finalizer in the; list is waiting for a signal (field value, external system, or other) produced by a; component responsible for a finalizer later in the list, resulting in a deadlock. Without; enforced ordering finalizers are free to order amongst themselves and are not vulnerable; to ordering changes in the list.
	GenerateName               *string                                            `json:"generateName,omitempty"`              // GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF; the Name field has not been provided. If this field is used, the name returned to the; client will be different than the name passed. This value will also be combined with a; unique suffix. The provided value has the same validation rules as the Name field, and; may be truncated by the length of the suffix required to make the value unique on the; server.; ; If this field is specified and the generated name exists, the server will return a 409.; ; Applied only if Name is not specified. More info:; https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency
	Generation                 *int64                                             `json:"generation,omitempty"`                // A sequence number representing a specific generation of the desired state. Populated by; the system. Read-only.
	Labels                     map[string]string                                  `json:"labels,omitempty"`                    // Map of string keys and values that can be used to organize and categorize (scope and; select) objects. May match selectors of replication controllers and services. More info:; http://kubernetes.io/docs/user-guide/labels
	ManagedFields              []IoK8SApimachineryPkgApisMetaV1ManagedFieldsEntry `json:"managedFields,omitempty"`             // ManagedFields maps workflow-id and version to the set of fields that are managed by that; workflow. This is mostly for internal housekeeping, and users typically shouldn't need to; set or understand this field. A workflow can be the user's name, a controller's name, or; the name of a specific apply path like "ci-cd". The set of fields is always in the; version that the workflow used when modifying the object.
	Name                       *string                                            `json:"name,omitempty"`                      // Name must be unique within a namespace. Is required when creating resources, although; some resources may allow a client to request the generation of an appropriate name; automatically. Name is primarily intended for creation idempotence and configuration; definition. Cannot be updated. More info:; http://kubernetes.io/docs/user-guide/identifiers#names
	Namespace                  *string                                            `json:"namespace,omitempty"`                 // Namespace defines the space within which each name must be unique. An empty namespace is; equivalent to the "default" namespace, but "default" is the canonical representation. Not; all objects are required to be scoped to a namespace - the value of this field for those; objects will be empty.; ; Must be a DNS_LABEL. Cannot be updated. More info:; http://kubernetes.io/docs/user-guide/namespaces
	OwnerReferences            []IoK8SApimachineryPkgApisMetaV1OwnerReferenceV2   `json:"ownerReferences,omitempty"`           // List of objects depended by this object. If ALL objects in the list have been deleted,; this object will be garbage collected. If this object is managed by a controller, then an; entry in this list will point to this controller, with the controller field set to true.; There cannot be more than one managing controller.
	ResourceVersion            *string                                            `json:"resourceVersion,omitempty"`           // An opaque value that represents the internal version of this object that can be used by; clients to determine when objects have changed. May be used for optimistic concurrency,; change detection, and the watch operation on a resource or set of resources. Clients must; treat these values as opaque and passed unmodified back to the server. They may only be; valid for a particular resource or set of resources.; ; Populated by the system. Read-only. Value must be treated as opaque by clients and . More; info:; https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	SelfLink                   *string                                            `json:"selfLink,omitempty"`                  // Deprecated: selfLink is a legacy read-only field that is no longer populated by the; system.
	Uid                        *string                                            `json:"uid,omitempty"`                       // UID is the unique in time and space value for this object. It is typically generated by; the server on successful creation of a resource and is not allowed to change on PUT; operations.; ; Populated by the system. Read-only. More info:; http://kubernetes.io/docs/user-guide/identifiers#uids
}

// io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry, ManagedFieldsEntry is a
// workflow-id, a FieldSet and the group version of the resource that the fieldset applies
// to.
type IoK8SApimachineryPkgApisMetaV1ManagedFieldsEntry struct {
	APIVersion  *string                `json:"apiVersion,omitempty"` // APIVersion defines the version of this resource that this field set applies to. The; format is "group/version" just like the top-level APIVersion field. It is necessary to; track the version of a field set because it cannot be automatically converted.
	FieldsType  *string                `json:"fieldsType,omitempty"` // FieldsType is the discriminator for the different fields format and version. There is; currently only one possible value: "FieldsV1"
	FieldsV1    map[string]interface{} `json:"fieldsV1,omitempty"`
	Manager     *string                `json:"manager,omitempty"`    // Manager is an identifier of the workflow managing these fields.
	Operation   *string                `json:"operation,omitempty"`  // Operation is the type of operation which lead to this ManagedFieldsEntry being created.; The only valid values for this field are 'Apply' and 'Update'.
	Subresource *string                `json:"subresource,omitempty"`// Subresource is the name of the subresource used to update that object, or empty string if; the object was updated through the main resource. The value of this field is used to; distinguish between managers, even if they share the same name. For example, a status; update will be distinct from a regular update using the same manager name. Note that the; APIVersion field is not related to the Subresource field and it always corresponds to the; version of the main resource.
	Time        *string                `json:"time,omitempty"`
}

// io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference_v2, OwnerReference contains enough
// information to let you identify an owning object. An owning object must be in the same
// namespace as the dependent, or be cluster-scoped, so there is no namespace field.
type IoK8SApimachineryPkgApisMetaV1OwnerReferenceV2 struct {
	APIVersion         string `json:"apiVersion"`                  // API version of the referent.
	BlockOwnerDeletion *bool  `json:"blockOwnerDeletion,omitempty"`// If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot; be deleted from the key-value store until this reference is removed. See; https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion; for how the garbage collector interacts with this field and enforces the foreground; deletion. Defaults to false. To set this field, a user needs "delete" permission of the; owner, otherwise 422 (Unprocessable Entity) will be returned.
	Controller         *bool  `json:"controller,omitempty"`        // If true, this reference points to the managing controller.
	Kind               string `json:"kind"`                        // Kind of the referent. More info:; https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Name               string `json:"name"`                        // Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Uid                string `json:"uid"`                         // UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
}

type Spec struct {
	AgentServer *string           `json:"agent_server"`         // If there is an agent on this node: ip+port.
	DataServer  *string           `json:"data_server"`          // If there is a DS on this node: ip+port. Updated by end users.
	DataShards  map[string]string `json:"data_shards,omitempty"`// Data shards that should be exported on this node. shard_id -> volume name Updated by; shard controller.
	MetaServer  *string           `json:"meta_server"`          // If there is a MDS on this node: ip+port. Updated by end users.
	MetaShards  map[string]string `json:"meta_shards,omitempty"`// Meta shards that should be exported on this node. shard_id -> volume name Updated by; shard controller.
	NodeUUID    *string           `json:"node_uuid"`            // Uuid of corresponding virtual machine. Currently used by elf cloud provider to find; corresponding virtual machine. Shouldn't be modified by manager or agent.
	Online      bool              `json:"online"`               // If the node should be considered as a candidate when placing shards. Updated by end users.
}

type Status struct {
	DataShards map[string]string `json:"data_shards"`// Mounted data shards. shard_id -> volume_name:volume_target_path NB: this indicates the; mount status of a shard on this node, but does not mean if it should be attached or; detached. Only shard leader should be attached.
	DataStatus *bool             `json:"data_status"`// Whether DS is healthy when there is a DS on this node: Data server status updated by this; node agent
	LastActive *string           `json:"last_active"`// Node level status Last heartbeat updated by agent
	MetaShards map[string]string `json:"meta_shards"`// Mounted meta shards. shard_id -> volume_name:volume_target_path NB: this indicates the; mount status of a shard on this node, but does not mean if it should be attached or; detached. Only shard leader should be attached.
	MetaStatus *bool             `json:"meta_status"`// Whether MDS is healthy when there is a MDS on this node: Meta server status updated by; this node agent
	Score      int64             `json:"score"`      // Load of the node.
}

// io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta, ListMeta describes metadata that synthetic
// resources must have, including lists and various status objects. A resource may have only
// one of {ObjectMeta, ListMeta}.
type IoK8SApimachineryPkgApisMetaV1ListMeta struct {
	Continue           *string `json:"continue,omitempty"`          // continue may be set if the user set a limit on the number of items returned, and; indicates that the server has more data available. The value is opaque and may be used to; issue another request to the endpoint that served this list to retrieve the next set of; available objects. Continuing a consistent list may not be possible if the server; configuration has changed or more than a few minutes have passed. The resourceVersion; field returned when using this continue value will be identical to the value in the first; response, unless you have received this token from an error message.
	RemainingItemCount *int64  `json:"remainingItemCount,omitempty"`// remainingItemCount is the number of subsequent items in the list which are not included; in this list response. If the list request contained label or field selectors, then the; number of remaining items is unknown and the field will be left unset and omitted during; serialization. If the list is complete (either because it is not chunking or because this; is the last chunk), then there are no more remaining items and this field will be left; unset and omitted during serialization. Servers older than v1.15 do not set this field.; The intended use of the remainingItemCount is *estimating* the size of a collection.; Clients should not rely on the remainingItemCount to be set or to be exact.
	ResourceVersion    *string `json:"resourceVersion,omitempty"`   // String that identifies the server's internal version of this object that can be used by; clients to determine when objects have changed. Value must be treated as opaque by; clients and passed unmodified back to the server. Populated by the system. Read-only.; More info:; https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	SelfLink           *string `json:"selfLink,omitempty"`          // Deprecated: selfLink is a legacy read-only field that is no longer populated by the; system.
}
