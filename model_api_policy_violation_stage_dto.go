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

// checks if the ApiPolicyViolationStageDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPolicyViolationStageDTO{}

// ApiPolicyViolationStageDTO struct for ApiPolicyViolationStageDTO
type ApiPolicyViolationStageDTO struct {
	ComponentPolicyViolations []ApiComponentPolicyViolationDTO `json:"componentPolicyViolations,omitempty"`
	StageId *string `json:"stageId,omitempty"`
}

// NewApiPolicyViolationStageDTO instantiates a new ApiPolicyViolationStageDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPolicyViolationStageDTO() *ApiPolicyViolationStageDTO {
	this := ApiPolicyViolationStageDTO{}
	return &this
}

// NewApiPolicyViolationStageDTOWithDefaults instantiates a new ApiPolicyViolationStageDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPolicyViolationStageDTOWithDefaults() *ApiPolicyViolationStageDTO {
	this := ApiPolicyViolationStageDTO{}
	return &this
}

// GetComponentPolicyViolations returns the ComponentPolicyViolations field value if set, zero value otherwise.
func (o *ApiPolicyViolationStageDTO) GetComponentPolicyViolations() []ApiComponentPolicyViolationDTO {
	if o == nil || IsNil(o.ComponentPolicyViolations) {
		var ret []ApiComponentPolicyViolationDTO
		return ret
	}
	return o.ComponentPolicyViolations
}

// GetComponentPolicyViolationsOk returns a tuple with the ComponentPolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyViolationStageDTO) GetComponentPolicyViolationsOk() ([]ApiComponentPolicyViolationDTO, bool) {
	if o == nil || IsNil(o.ComponentPolicyViolations) {
		return nil, false
	}
	return o.ComponentPolicyViolations, true
}

// HasComponentPolicyViolations returns a boolean if a field has been set.
func (o *ApiPolicyViolationStageDTO) HasComponentPolicyViolations() bool {
	if o != nil && !IsNil(o.ComponentPolicyViolations) {
		return true
	}

	return false
}

// SetComponentPolicyViolations gets a reference to the given []ApiComponentPolicyViolationDTO and assigns it to the ComponentPolicyViolations field.
func (o *ApiPolicyViolationStageDTO) SetComponentPolicyViolations(v []ApiComponentPolicyViolationDTO) {
	o.ComponentPolicyViolations = v
}

// GetStageId returns the StageId field value if set, zero value otherwise.
func (o *ApiPolicyViolationStageDTO) GetStageId() string {
	if o == nil || IsNil(o.StageId) {
		var ret string
		return ret
	}
	return *o.StageId
}

// GetStageIdOk returns a tuple with the StageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyViolationStageDTO) GetStageIdOk() (*string, bool) {
	if o == nil || IsNil(o.StageId) {
		return nil, false
	}
	return o.StageId, true
}

// HasStageId returns a boolean if a field has been set.
func (o *ApiPolicyViolationStageDTO) HasStageId() bool {
	if o != nil && !IsNil(o.StageId) {
		return true
	}

	return false
}

// SetStageId gets a reference to the given string and assigns it to the StageId field.
func (o *ApiPolicyViolationStageDTO) SetStageId(v string) {
	o.StageId = &v
}

func (o ApiPolicyViolationStageDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPolicyViolationStageDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentPolicyViolations) {
		toSerialize["componentPolicyViolations"] = o.ComponentPolicyViolations
	}
	if !IsNil(o.StageId) {
		toSerialize["stageId"] = o.StageId
	}
	return toSerialize, nil
}

type NullableApiPolicyViolationStageDTO struct {
	value *ApiPolicyViolationStageDTO
	isSet bool
}

func (v NullableApiPolicyViolationStageDTO) Get() *ApiPolicyViolationStageDTO {
	return v.value
}

func (v *NullableApiPolicyViolationStageDTO) Set(val *ApiPolicyViolationStageDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPolicyViolationStageDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPolicyViolationStageDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPolicyViolationStageDTO(val *ApiPolicyViolationStageDTO) *NullableApiPolicyViolationStageDTO {
	return &NullableApiPolicyViolationStageDTO{value: val, isSet: true}
}

func (v NullableApiPolicyViolationStageDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPolicyViolationStageDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


