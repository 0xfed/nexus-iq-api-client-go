/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.176.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiHashComponentIdentifiersDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHashComponentIdentifiersDTO{}

// ApiHashComponentIdentifiersDTO struct for ApiHashComponentIdentifiersDTO
type ApiHashComponentIdentifiersDTO struct {
	ComponentClaims []ApiHashComponentIdentifierDTO `json:"componentClaims,omitempty"`
}

// NewApiHashComponentIdentifiersDTO instantiates a new ApiHashComponentIdentifiersDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHashComponentIdentifiersDTO() *ApiHashComponentIdentifiersDTO {
	this := ApiHashComponentIdentifiersDTO{}
	return &this
}

// NewApiHashComponentIdentifiersDTOWithDefaults instantiates a new ApiHashComponentIdentifiersDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHashComponentIdentifiersDTOWithDefaults() *ApiHashComponentIdentifiersDTO {
	this := ApiHashComponentIdentifiersDTO{}
	return &this
}

// GetComponentClaims returns the ComponentClaims field value if set, zero value otherwise.
func (o *ApiHashComponentIdentifiersDTO) GetComponentClaims() []ApiHashComponentIdentifierDTO {
	if o == nil || IsNil(o.ComponentClaims) {
		var ret []ApiHashComponentIdentifierDTO
		return ret
	}
	return o.ComponentClaims
}

// GetComponentClaimsOk returns a tuple with the ComponentClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHashComponentIdentifiersDTO) GetComponentClaimsOk() ([]ApiHashComponentIdentifierDTO, bool) {
	if o == nil || IsNil(o.ComponentClaims) {
		return nil, false
	}
	return o.ComponentClaims, true
}

// HasComponentClaims returns a boolean if a field has been set.
func (o *ApiHashComponentIdentifiersDTO) HasComponentClaims() bool {
	if o != nil && !IsNil(o.ComponentClaims) {
		return true
	}

	return false
}

// SetComponentClaims gets a reference to the given []ApiHashComponentIdentifierDTO and assigns it to the ComponentClaims field.
func (o *ApiHashComponentIdentifiersDTO) SetComponentClaims(v []ApiHashComponentIdentifierDTO) {
	o.ComponentClaims = v
}

func (o ApiHashComponentIdentifiersDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHashComponentIdentifiersDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentClaims) {
		toSerialize["componentClaims"] = o.ComponentClaims
	}
	return toSerialize, nil
}

type NullableApiHashComponentIdentifiersDTO struct {
	value *ApiHashComponentIdentifiersDTO
	isSet bool
}

func (v NullableApiHashComponentIdentifiersDTO) Get() *ApiHashComponentIdentifiersDTO {
	return v.value
}

func (v *NullableApiHashComponentIdentifiersDTO) Set(val *ApiHashComponentIdentifiersDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHashComponentIdentifiersDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHashComponentIdentifiersDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHashComponentIdentifiersDTO(val *ApiHashComponentIdentifiersDTO) *NullableApiHashComponentIdentifiersDTO {
	return &NullableApiHashComponentIdentifiersDTO{value: val, isSet: true}
}

func (v NullableApiHashComponentIdentifiersDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHashComponentIdentifiersDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


