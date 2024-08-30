/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.181.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiDependencyTreeResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDependencyTreeResponseDTO{}

// ApiDependencyTreeResponseDTO struct for ApiDependencyTreeResponseDTO
type ApiDependencyTreeResponseDTO struct {
	DependencyTree *ApiDependencyTreeNodeDTO `json:"dependencyTree,omitempty"`
}

// NewApiDependencyTreeResponseDTO instantiates a new ApiDependencyTreeResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDependencyTreeResponseDTO() *ApiDependencyTreeResponseDTO {
	this := ApiDependencyTreeResponseDTO{}
	return &this
}

// NewApiDependencyTreeResponseDTOWithDefaults instantiates a new ApiDependencyTreeResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDependencyTreeResponseDTOWithDefaults() *ApiDependencyTreeResponseDTO {
	this := ApiDependencyTreeResponseDTO{}
	return &this
}

// GetDependencyTree returns the DependencyTree field value if set, zero value otherwise.
func (o *ApiDependencyTreeResponseDTO) GetDependencyTree() ApiDependencyTreeNodeDTO {
	if o == nil || IsNil(o.DependencyTree) {
		var ret ApiDependencyTreeNodeDTO
		return ret
	}
	return *o.DependencyTree
}

// GetDependencyTreeOk returns a tuple with the DependencyTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDependencyTreeResponseDTO) GetDependencyTreeOk() (*ApiDependencyTreeNodeDTO, bool) {
	if o == nil || IsNil(o.DependencyTree) {
		return nil, false
	}
	return o.DependencyTree, true
}

// HasDependencyTree returns a boolean if a field has been set.
func (o *ApiDependencyTreeResponseDTO) HasDependencyTree() bool {
	if o != nil && !IsNil(o.DependencyTree) {
		return true
	}

	return false
}

// SetDependencyTree gets a reference to the given ApiDependencyTreeNodeDTO and assigns it to the DependencyTree field.
func (o *ApiDependencyTreeResponseDTO) SetDependencyTree(v ApiDependencyTreeNodeDTO) {
	o.DependencyTree = &v
}

func (o ApiDependencyTreeResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDependencyTreeResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DependencyTree) {
		toSerialize["dependencyTree"] = o.DependencyTree
	}
	return toSerialize, nil
}

type NullableApiDependencyTreeResponseDTO struct {
	value *ApiDependencyTreeResponseDTO
	isSet bool
}

func (v NullableApiDependencyTreeResponseDTO) Get() *ApiDependencyTreeResponseDTO {
	return v.value
}

func (v *NullableApiDependencyTreeResponseDTO) Set(val *ApiDependencyTreeResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDependencyTreeResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDependencyTreeResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDependencyTreeResponseDTO(val *ApiDependencyTreeResponseDTO) *NullableApiDependencyTreeResponseDTO {
	return &NullableApiDependencyTreeResponseDTO{value: val, isSet: true}
}

func (v NullableApiDependencyTreeResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDependencyTreeResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


