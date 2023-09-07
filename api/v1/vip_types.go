package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VipSpec defines the desired state of Vip
type VipSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	Addr string `json:"addr"`// The ipv4 addr.
}

// VipStatus defines the observed state of Vip
type VipStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	Ready       bool    `json:"ready"`       // If this vip is ready.
	SessionName *string `json:"session_name"`// Which session this vip is appointed.
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Vip is the Schema for the vips API
type Vip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VipSpec   `json:"spec,omitempty"`
	Status VipStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VipList contains a list of Vip
type VipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vip `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Vip{}, &VipList{})
}
