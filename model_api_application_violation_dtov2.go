/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.180.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiApplicationViolationDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiApplicationViolationDTOV2{}

// ApiApplicationViolationDTOV2 struct for ApiApplicationViolationDTOV2
type ApiApplicationViolationDTOV2 struct {
	Application *ApiApplicationBaseDTO `json:"application,omitempty"`
	PolicyViolations []ApiEnhancedPolicyViolationDTOV2 `json:"policyViolations,omitempty"`
}

// NewApiApplicationViolationDTOV2 instantiates a new ApiApplicationViolationDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiApplicationViolationDTOV2() *ApiApplicationViolationDTOV2 {
	this := ApiApplicationViolationDTOV2{}
	return &this
}

// NewApiApplicationViolationDTOV2WithDefaults instantiates a new ApiApplicationViolationDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiApplicationViolationDTOV2WithDefaults() *ApiApplicationViolationDTOV2 {
	this := ApiApplicationViolationDTOV2{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ApiApplicationViolationDTOV2) GetApplication() ApiApplicationBaseDTO {
	if o == nil || IsNil(o.Application) {
		var ret ApiApplicationBaseDTO
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationViolationDTOV2) GetApplicationOk() (*ApiApplicationBaseDTO, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ApiApplicationViolationDTOV2) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given ApiApplicationBaseDTO and assigns it to the Application field.
func (o *ApiApplicationViolationDTOV2) SetApplication(v ApiApplicationBaseDTO) {
	o.Application = &v
}

// GetPolicyViolations returns the PolicyViolations field value if set, zero value otherwise.
func (o *ApiApplicationViolationDTOV2) GetPolicyViolations() []ApiEnhancedPolicyViolationDTOV2 {
	if o == nil || IsNil(o.PolicyViolations) {
		var ret []ApiEnhancedPolicyViolationDTOV2
		return ret
	}
	return o.PolicyViolations
}

// GetPolicyViolationsOk returns a tuple with the PolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationViolationDTOV2) GetPolicyViolationsOk() ([]ApiEnhancedPolicyViolationDTOV2, bool) {
	if o == nil || IsNil(o.PolicyViolations) {
		return nil, false
	}
	return o.PolicyViolations, true
}

// HasPolicyViolations returns a boolean if a field has been set.
func (o *ApiApplicationViolationDTOV2) HasPolicyViolations() bool {
	if o != nil && !IsNil(o.PolicyViolations) {
		return true
	}

	return false
}

// SetPolicyViolations gets a reference to the given []ApiEnhancedPolicyViolationDTOV2 and assigns it to the PolicyViolations field.
func (o *ApiApplicationViolationDTOV2) SetPolicyViolations(v []ApiEnhancedPolicyViolationDTOV2) {
	o.PolicyViolations = v
}

func (o ApiApplicationViolationDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiApplicationViolationDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.PolicyViolations) {
		toSerialize["policyViolations"] = o.PolicyViolations
	}
	return toSerialize, nil
}

type NullableApiApplicationViolationDTOV2 struct {
	value *ApiApplicationViolationDTOV2
	isSet bool
}

func (v NullableApiApplicationViolationDTOV2) Get() *ApiApplicationViolationDTOV2 {
	return v.value
}

func (v *NullableApiApplicationViolationDTOV2) Set(val *ApiApplicationViolationDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiApplicationViolationDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiApplicationViolationDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiApplicationViolationDTOV2(val *ApiApplicationViolationDTOV2) *NullableApiApplicationViolationDTOV2 {
	return &NullableApiApplicationViolationDTOV2{value: val, isSet: true}
}

func (v NullableApiApplicationViolationDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiApplicationViolationDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


