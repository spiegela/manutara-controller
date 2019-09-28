/*

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

// MutationSpec defines the desired state of Mutation
type MutationSpec struct {
	FieldName   string          `json:"fieldName"`
	Description string          `json:"description"`
	Args        []DataTypeField `json:"args"`
	Type        DataTypeField   `json:"type"`
}

// MutationStatus defines the observed state of Mutation
type MutationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Mutation is the Schema for the mutations API
type Mutation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MutationSpec   `json:"spec,omitempty"`
	Status MutationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MutationList contains a list of Mutation
type MutationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mutation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mutation{}, &MutationList{})
}
