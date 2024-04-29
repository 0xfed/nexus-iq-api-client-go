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

// checks if the ApiReportConstraintConditionDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReportConstraintConditionDTOV2{}

// ApiReportConstraintConditionDTOV2 struct for ApiReportConstraintConditionDTOV2
type ApiReportConstraintConditionDTOV2 struct {
	ConditionReason *string `json:"conditionReason,omitempty"`
	ConditionSummary *string `json:"conditionSummary,omitempty"`
}

// NewApiReportConstraintConditionDTOV2 instantiates a new ApiReportConstraintConditionDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportConstraintConditionDTOV2() *ApiReportConstraintConditionDTOV2 {
	this := ApiReportConstraintConditionDTOV2{}
	return &this
}

// NewApiReportConstraintConditionDTOV2WithDefaults instantiates a new ApiReportConstraintConditionDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportConstraintConditionDTOV2WithDefaults() *ApiReportConstraintConditionDTOV2 {
	this := ApiReportConstraintConditionDTOV2{}
	return &this
}

// GetConditionReason returns the ConditionReason field value if set, zero value otherwise.
func (o *ApiReportConstraintConditionDTOV2) GetConditionReason() string {
	if o == nil || IsNil(o.ConditionReason) {
		var ret string
		return ret
	}
	return *o.ConditionReason
}

// GetConditionReasonOk returns a tuple with the ConditionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportConstraintConditionDTOV2) GetConditionReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionReason) {
		return nil, false
	}
	return o.ConditionReason, true
}

// HasConditionReason returns a boolean if a field has been set.
func (o *ApiReportConstraintConditionDTOV2) HasConditionReason() bool {
	if o != nil && !IsNil(o.ConditionReason) {
		return true
	}

	return false
}

// SetConditionReason gets a reference to the given string and assigns it to the ConditionReason field.
func (o *ApiReportConstraintConditionDTOV2) SetConditionReason(v string) {
	o.ConditionReason = &v
}

// GetConditionSummary returns the ConditionSummary field value if set, zero value otherwise.
func (o *ApiReportConstraintConditionDTOV2) GetConditionSummary() string {
	if o == nil || IsNil(o.ConditionSummary) {
		var ret string
		return ret
	}
	return *o.ConditionSummary
}

// GetConditionSummaryOk returns a tuple with the ConditionSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportConstraintConditionDTOV2) GetConditionSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionSummary) {
		return nil, false
	}
	return o.ConditionSummary, true
}

// HasConditionSummary returns a boolean if a field has been set.
func (o *ApiReportConstraintConditionDTOV2) HasConditionSummary() bool {
	if o != nil && !IsNil(o.ConditionSummary) {
		return true
	}

	return false
}

// SetConditionSummary gets a reference to the given string and assigns it to the ConditionSummary field.
func (o *ApiReportConstraintConditionDTOV2) SetConditionSummary(v string) {
	o.ConditionSummary = &v
}

func (o ApiReportConstraintConditionDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportConstraintConditionDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConditionReason) {
		toSerialize["conditionReason"] = o.ConditionReason
	}
	if !IsNil(o.ConditionSummary) {
		toSerialize["conditionSummary"] = o.ConditionSummary
	}
	return toSerialize, nil
}

type NullableApiReportConstraintConditionDTOV2 struct {
	value *ApiReportConstraintConditionDTOV2
	isSet bool
}

func (v NullableApiReportConstraintConditionDTOV2) Get() *ApiReportConstraintConditionDTOV2 {
	return v.value
}

func (v *NullableApiReportConstraintConditionDTOV2) Set(val *ApiReportConstraintConditionDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportConstraintConditionDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportConstraintConditionDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportConstraintConditionDTOV2(val *ApiReportConstraintConditionDTOV2) *NullableApiReportConstraintConditionDTOV2 {
	return &NullableApiReportConstraintConditionDTOV2{value: val, isSet: true}
}

func (v NullableApiReportConstraintConditionDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportConstraintConditionDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


