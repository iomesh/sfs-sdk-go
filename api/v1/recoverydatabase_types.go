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

// RecoveryDatabaseSpec defines the desired state of RecoveryDatabase
type RecoveryDatabaseSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	Clids           map[string]string   `json:"clids"`           // Client entry records, clientid -> cid_recov_tag.
	ReclaimingClids []int64             `json:"reclaiming_clids"`// Reclaiming client entry records, clientid -> cid_recov_tag.
	Rfhs            map[string][]string `json:"rfhs"`            // Revoked file handle records, cid_recov_tag -> Vec<revoked file handle>.
	Version         int64               `json:"version"`         // Version number.
}

// RecoveryDatabaseStatus defines the observed state of RecoveryDatabase
type RecoveryDatabaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RecoveryDatabase is the Schema for the recoverydatabases API
type RecoveryDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RecoveryDatabaseSpec   `json:"spec,omitempty"`
	Status RecoveryDatabaseStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RecoveryDatabaseList contains a list of RecoveryDatabase
type RecoveryDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecoveryDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RecoveryDatabase{}, &RecoveryDatabaseList{})
}
