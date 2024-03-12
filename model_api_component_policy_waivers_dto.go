/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.174.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiComponentPolicyWaiversDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentPolicyWaiversDTO{}

// ApiComponentPolicyWaiversDTO struct for ApiComponentPolicyWaiversDTO
type ApiComponentPolicyWaiversDTO struct {
	ComponentPolicyWaivers []ApiPolicyWaiverDTO `json:"componentPolicyWaivers,omitempty"`
}

// NewApiComponentPolicyWaiversDTO instantiates a new ApiComponentPolicyWaiversDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentPolicyWaiversDTO() *ApiComponentPolicyWaiversDTO {
	this := ApiComponentPolicyWaiversDTO{}
	return &this
}

// NewApiComponentPolicyWaiversDTOWithDefaults instantiates a new ApiComponentPolicyWaiversDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentPolicyWaiversDTOWithDefaults() *ApiComponentPolicyWaiversDTO {
	this := ApiComponentPolicyWaiversDTO{}
	return &this
}

// GetComponentPolicyWaivers returns the ComponentPolicyWaivers field value if set, zero value otherwise.
func (o *ApiComponentPolicyWaiversDTO) GetComponentPolicyWaivers() []ApiPolicyWaiverDTO {
	if o == nil || IsNil(o.ComponentPolicyWaivers) {
		var ret []ApiPolicyWaiverDTO
		return ret
	}
	return o.ComponentPolicyWaivers
}

// GetComponentPolicyWaiversOk returns a tuple with the ComponentPolicyWaivers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentPolicyWaiversDTO) GetComponentPolicyWaiversOk() ([]ApiPolicyWaiverDTO, bool) {
	if o == nil || IsNil(o.ComponentPolicyWaivers) {
		return nil, false
	}
	return o.ComponentPolicyWaivers, true
}

// HasComponentPolicyWaivers returns a boolean if a field has been set.
func (o *ApiComponentPolicyWaiversDTO) HasComponentPolicyWaivers() bool {
	if o != nil && !IsNil(o.ComponentPolicyWaivers) {
		return true
	}

	return false
}

// SetComponentPolicyWaivers gets a reference to the given []ApiPolicyWaiverDTO and assigns it to the ComponentPolicyWaivers field.
func (o *ApiComponentPolicyWaiversDTO) SetComponentPolicyWaivers(v []ApiPolicyWaiverDTO) {
	o.ComponentPolicyWaivers = v
}

func (o ApiComponentPolicyWaiversDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentPolicyWaiversDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentPolicyWaivers) {
		toSerialize["componentPolicyWaivers"] = o.ComponentPolicyWaivers
	}
	return toSerialize, nil
}

type NullableApiComponentPolicyWaiversDTO struct {
	value *ApiComponentPolicyWaiversDTO
	isSet bool
}

func (v NullableApiComponentPolicyWaiversDTO) Get() *ApiComponentPolicyWaiversDTO {
	return v.value
}

func (v *NullableApiComponentPolicyWaiversDTO) Set(val *ApiComponentPolicyWaiversDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentPolicyWaiversDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentPolicyWaiversDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentPolicyWaiversDTO(val *ApiComponentPolicyWaiversDTO) *NullableApiComponentPolicyWaiversDTO {
	return &NullableApiComponentPolicyWaiversDTO{value: val, isSet: true}
}

func (v NullableApiComponentPolicyWaiversDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentPolicyWaiversDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


