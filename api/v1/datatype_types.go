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

// DataTypeSpec defines the desired state of DataType
type DataTypeSpec struct {
	// Description is the user presented short description of the data type
	Description string `json:"description"`

	// Fields is the list of fields making up the data type
	Fields DataTypeFields `json:"fields"`
}

// DataTypeFields is a named list of fields used in a query
type DataTypeFields map[string]DataTypeField

type DataTypeBasicFieldType string

const (
	// DataTypeFieldID is a GraphQL ID data type
	DataTypeIDField DataTypeBasicFieldType = "ID"

	// DataTypeFieldString is a GraphQL String data type
	DataTypeStringField DataTypeBasicFieldType = "String"

	// DataTypeFieldInt is a GraphQL Int data type
	DataTypeIntField DataTypeBasicFieldType = "Int"

	// DataTypeFieldFloat is a GraphQL Float data type
	DataTypeFloatField DataTypeBasicFieldType = "Float"

	// DataTypeFieldBoolean is a GraphQL Boolean data type
	DataTypeBooleanField DataTypeBasicFieldType = "Boolean"

	// DataTypeFieldDate is a GraphQL Date data type
	DataTypeDateField DataTypeBasicFieldType = "Date"
)

// DataTypeFieldUnionType is a list of types that are valid as values for a
// field
type DataTypeFieldUnionType []string

// DataTypeFieldEnumType is a list of values that are valid for a field
type DataTypeFieldEnumType []string

// DataTypeField defines a GraphQL field for the data type
type DataTypeField struct {
	// Description is the user presented short description of the field
	Description string `json:"description"`

	// BasicType is the data type of the field. If the field references a user
	// defined type, such as an interface or a GraphQL Data Type, this should be
	// empty
	BasicType DataTypeBasicFieldType `json:"type"`

	// UserDefinedType is a type that is defined within the GraphQL schema and
	// should either be an DataType name or an interface
	UserDefinedType string `json:"userDefinedType"`

	// UnionType is a list of types that are valid for for this field
	//UnionType DataTypeFieldUnionType `json:"union"`

	// EnumType is a list of valid values that are valid for for this field
	//EnumType DataTypeFieldEnumType `json:"enum"`

	// IsList specifies if this type is a list of objects
	IsList bool `json:"isList"`

	// AllowEmpty defines if an empty list is allowed in the case that this is
	// a list field
	AllowEmpty bool `json:"allowEmpty"`

	// AllowNull defines if a null value is allowed for the field
	AllowNull bool `json:"allowNull"`
}

// DataTypeStatus defines the observed state of DataType
type DataTypeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// DataType is the Schema for the datatypes API
type DataType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataTypeSpec   `json:"spec,omitempty"`
	Status DataTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataTypeList contains a list of DataType
type DataTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataType{}, &DataTypeList{})
}