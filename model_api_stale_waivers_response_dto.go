/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.175.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiStaleWaiversResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStaleWaiversResponseDTO{}

// ApiStaleWaiversResponseDTO struct for ApiStaleWaiversResponseDTO
type ApiStaleWaiversResponseDTO struct {
	StaleWaivers []ApiStaleWaiverDTO `json:"staleWaivers,omitempty"`
}

// NewApiStaleWaiversResponseDTO instantiates a new ApiStaleWaiversResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStaleWaiversResponseDTO() *ApiStaleWaiversResponseDTO {
	this := ApiStaleWaiversResponseDTO{}
	return &this
}

// NewApiStaleWaiversResponseDTOWithDefaults instantiates a new ApiStaleWaiversResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStaleWaiversResponseDTOWithDefaults() *ApiStaleWaiversResponseDTO {
	this := ApiStaleWaiversResponseDTO{}
	return &this
}

// GetStaleWaivers returns the StaleWaivers field value if set, zero value otherwise.
func (o *ApiStaleWaiversResponseDTO) GetStaleWaivers() []ApiStaleWaiverDTO {
	if o == nil || IsNil(o.StaleWaivers) {
		var ret []ApiStaleWaiverDTO
		return ret
	}
	return o.StaleWaivers
}

// GetStaleWaiversOk returns a tuple with the StaleWaivers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiversResponseDTO) GetStaleWaiversOk() ([]ApiStaleWaiverDTO, bool) {
	if o == nil || IsNil(o.StaleWaivers) {
		return nil, false
	}
	return o.StaleWaivers, true
}

// HasStaleWaivers returns a boolean if a field has been set.
func (o *ApiStaleWaiversResponseDTO) HasStaleWaivers() bool {
	if o != nil && !IsNil(o.StaleWaivers) {
		return true
	}

	return false
}

// SetStaleWaivers gets a reference to the given []ApiStaleWaiverDTO and assigns it to the StaleWaivers field.
func (o *ApiStaleWaiversResponseDTO) SetStaleWaivers(v []ApiStaleWaiverDTO) {
	o.StaleWaivers = v
}

func (o ApiStaleWaiversResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStaleWaiversResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StaleWaivers) {
		toSerialize["staleWaivers"] = o.StaleWaivers
	}
	return toSerialize, nil
}

type NullableApiStaleWaiversResponseDTO struct {
	value *ApiStaleWaiversResponseDTO
	isSet bool
}

func (v NullableApiStaleWaiversResponseDTO) Get() *ApiStaleWaiversResponseDTO {
	return v.value
}

func (v *NullableApiStaleWaiversResponseDTO) Set(val *ApiStaleWaiversResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStaleWaiversResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStaleWaiversResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStaleWaiversResponseDTO(val *ApiStaleWaiversResponseDTO) *NullableApiStaleWaiversResponseDTO {
	return &NullableApiStaleWaiversResponseDTO{value: val, isSet: true}
}

func (v NullableApiStaleWaiversResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStaleWaiversResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


