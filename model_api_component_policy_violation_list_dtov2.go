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

// checks if the ApiComponentPolicyViolationListDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentPolicyViolationListDTOV2{}

// ApiComponentPolicyViolationListDTOV2 struct for ApiComponentPolicyViolationListDTOV2
type ApiComponentPolicyViolationListDTOV2 struct {
	PolicyViolations []ApiPolicyViolationDTOV2 `json:"policyViolations,omitempty"`
}

// NewApiComponentPolicyViolationListDTOV2 instantiates a new ApiComponentPolicyViolationListDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentPolicyViolationListDTOV2() *ApiComponentPolicyViolationListDTOV2 {
	this := ApiComponentPolicyViolationListDTOV2{}
	return &this
}

// NewApiComponentPolicyViolationListDTOV2WithDefaults instantiates a new ApiComponentPolicyViolationListDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentPolicyViolationListDTOV2WithDefaults() *ApiComponentPolicyViolationListDTOV2 {
	this := ApiComponentPolicyViolationListDTOV2{}
	return &this
}

// GetPolicyViolations returns the PolicyViolations field value if set, zero value otherwise.
func (o *ApiComponentPolicyViolationListDTOV2) GetPolicyViolations() []ApiPolicyViolationDTOV2 {
	if o == nil || IsNil(o.PolicyViolations) {
		var ret []ApiPolicyViolationDTOV2
		return ret
	}
	return o.PolicyViolations
}

// GetPolicyViolationsOk returns a tuple with the PolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentPolicyViolationListDTOV2) GetPolicyViolationsOk() ([]ApiPolicyViolationDTOV2, bool) {
	if o == nil || IsNil(o.PolicyViolations) {
		return nil, false
	}
	return o.PolicyViolations, true
}

// HasPolicyViolations returns a boolean if a field has been set.
func (o *ApiComponentPolicyViolationListDTOV2) HasPolicyViolations() bool {
	if o != nil && !IsNil(o.PolicyViolations) {
		return true
	}

	return false
}

// SetPolicyViolations gets a reference to the given []ApiPolicyViolationDTOV2 and assigns it to the PolicyViolations field.
func (o *ApiComponentPolicyViolationListDTOV2) SetPolicyViolations(v []ApiPolicyViolationDTOV2) {
	o.PolicyViolations = v
}

func (o ApiComponentPolicyViolationListDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentPolicyViolationListDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyViolations) {
		toSerialize["policyViolations"] = o.PolicyViolations
	}
	return toSerialize, nil
}

type NullableApiComponentPolicyViolationListDTOV2 struct {
	value *ApiComponentPolicyViolationListDTOV2
	isSet bool
}

func (v NullableApiComponentPolicyViolationListDTOV2) Get() *ApiComponentPolicyViolationListDTOV2 {
	return v.value
}

func (v *NullableApiComponentPolicyViolationListDTOV2) Set(val *ApiComponentPolicyViolationListDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentPolicyViolationListDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentPolicyViolationListDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentPolicyViolationListDTOV2(val *ApiComponentPolicyViolationListDTOV2) *NullableApiComponentPolicyViolationListDTOV2 {
	return &NullableApiComponentPolicyViolationListDTOV2{value: val, isSet: true}
}

func (v NullableApiComponentPolicyViolationListDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentPolicyViolationListDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


