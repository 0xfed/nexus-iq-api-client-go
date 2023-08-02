/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 165
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiComponentProjectScmDetailsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentProjectScmDetailsDTO{}

// ApiComponentProjectScmDetailsDTO struct for ApiComponentProjectScmDetailsDTO
type ApiComponentProjectScmDetailsDTO struct {
	CommitsPerMonth *int32 `json:"commitsPerMonth,omitempty"`
	UniqueDevsPerMonth *int32 `json:"uniqueDevsPerMonth,omitempty"`
}

// NewApiComponentProjectScmDetailsDTO instantiates a new ApiComponentProjectScmDetailsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentProjectScmDetailsDTO() *ApiComponentProjectScmDetailsDTO {
	this := ApiComponentProjectScmDetailsDTO{}
	return &this
}

// NewApiComponentProjectScmDetailsDTOWithDefaults instantiates a new ApiComponentProjectScmDetailsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentProjectScmDetailsDTOWithDefaults() *ApiComponentProjectScmDetailsDTO {
	this := ApiComponentProjectScmDetailsDTO{}
	return &this
}

// GetCommitsPerMonth returns the CommitsPerMonth field value if set, zero value otherwise.
func (o *ApiComponentProjectScmDetailsDTO) GetCommitsPerMonth() int32 {
	if o == nil || IsNil(o.CommitsPerMonth) {
		var ret int32
		return ret
	}
	return *o.CommitsPerMonth
}

// GetCommitsPerMonthOk returns a tuple with the CommitsPerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectScmDetailsDTO) GetCommitsPerMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.CommitsPerMonth) {
		return nil, false
	}
	return o.CommitsPerMonth, true
}

// HasCommitsPerMonth returns a boolean if a field has been set.
func (o *ApiComponentProjectScmDetailsDTO) HasCommitsPerMonth() bool {
	if o != nil && !IsNil(o.CommitsPerMonth) {
		return true
	}

	return false
}

// SetCommitsPerMonth gets a reference to the given int32 and assigns it to the CommitsPerMonth field.
func (o *ApiComponentProjectScmDetailsDTO) SetCommitsPerMonth(v int32) {
	o.CommitsPerMonth = &v
}

// GetUniqueDevsPerMonth returns the UniqueDevsPerMonth field value if set, zero value otherwise.
func (o *ApiComponentProjectScmDetailsDTO) GetUniqueDevsPerMonth() int32 {
	if o == nil || IsNil(o.UniqueDevsPerMonth) {
		var ret int32
		return ret
	}
	return *o.UniqueDevsPerMonth
}

// GetUniqueDevsPerMonthOk returns a tuple with the UniqueDevsPerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectScmDetailsDTO) GetUniqueDevsPerMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueDevsPerMonth) {
		return nil, false
	}
	return o.UniqueDevsPerMonth, true
}

// HasUniqueDevsPerMonth returns a boolean if a field has been set.
func (o *ApiComponentProjectScmDetailsDTO) HasUniqueDevsPerMonth() bool {
	if o != nil && !IsNil(o.UniqueDevsPerMonth) {
		return true
	}

	return false
}

// SetUniqueDevsPerMonth gets a reference to the given int32 and assigns it to the UniqueDevsPerMonth field.
func (o *ApiComponentProjectScmDetailsDTO) SetUniqueDevsPerMonth(v int32) {
	o.UniqueDevsPerMonth = &v
}

func (o ApiComponentProjectScmDetailsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentProjectScmDetailsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommitsPerMonth) {
		toSerialize["commitsPerMonth"] = o.CommitsPerMonth
	}
	if !IsNil(o.UniqueDevsPerMonth) {
		toSerialize["uniqueDevsPerMonth"] = o.UniqueDevsPerMonth
	}
	return toSerialize, nil
}

type NullableApiComponentProjectScmDetailsDTO struct {
	value *ApiComponentProjectScmDetailsDTO
	isSet bool
}

func (v NullableApiComponentProjectScmDetailsDTO) Get() *ApiComponentProjectScmDetailsDTO {
	return v.value
}

func (v *NullableApiComponentProjectScmDetailsDTO) Set(val *ApiComponentProjectScmDetailsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentProjectScmDetailsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentProjectScmDetailsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentProjectScmDetailsDTO(val *ApiComponentProjectScmDetailsDTO) *NullableApiComponentProjectScmDetailsDTO {
	return &NullableApiComponentProjectScmDetailsDTO{value: val, isSet: true}
}

func (v NullableApiComponentProjectScmDetailsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentProjectScmDetailsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


