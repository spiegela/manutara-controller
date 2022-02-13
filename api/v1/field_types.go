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

// FieldKind is the kind of schema field of the field.
type FieldKind int

// Possible FieldKind types
const (
	MutationFieldKind FieldKind = iota
	QueryFieldKind
	SubscriptionFieldKind
)

// FieldSpec defines the desired state of Field
type FieldSpec struct {
	FieldName   string            `json:"fieldName"`
	Description string            `json:"description"`
	Args        map[string]string `json:"args"`
	Kind        FieldKind         `json:"kind"`
	Type        string            `json:"type"`
}

// FieldStatus defines the observed state of Field
type FieldStatus struct {
}

// +kubebuilder:object:root=true

// Field is the Schema for the mutations API
type Field struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FieldSpec   `json:"spec,omitempty"`
	Status FieldStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FieldList contains a list of Field
type FieldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Field `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Field{}, &FieldList{})
}
