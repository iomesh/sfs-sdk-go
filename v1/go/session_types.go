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

// SessionSpec defines the desired state of Session
type SessionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	ClientAddr   string            `json:"client_addr"`   // Exposed <ip:port>.
	NodeName     string            `json:"node_name"`     // Node associated with the session.
	ProtocolAddr string            `json:"protocol_addr"` // Exposed <ip:port>.
	Stype        Stype             `json:"stype"`         // Session type, e.g. nfs, fuse and etc.
	Vips         map[string]string `json:"vips,omitempty"`// VIP list assigned to this session.
}

// Session type, e.g. nfs, fuse and etc.
type Stype string
const (
	Fuse Stype = "FUSE"
	NFS Stype = "NFS"
)

// SessionStatus defines the observed state of Session
type SessionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	SessionID int64 `json:"session_id"`// Cluster wide exclusive ID.
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Session is the Schema for the sessions API
type Session struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SessionSpec   `json:"spec,omitempty"`
	Status SessionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SessionList contains a list of Session
type SessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Session `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Session{}, &SessionList{})
}
