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

// checks if the ApiMatchStateSummaryDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMatchStateSummaryDTOV2{}

// ApiMatchStateSummaryDTOV2 struct for ApiMatchStateSummaryDTOV2
type ApiMatchStateSummaryDTOV2 struct {
	KnownComponentCount *int32 `json:"knownComponentCount,omitempty"`
	TotalComponentCount *int32 `json:"totalComponentCount,omitempty"`
}

// NewApiMatchStateSummaryDTOV2 instantiates a new ApiMatchStateSummaryDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMatchStateSummaryDTOV2() *ApiMatchStateSummaryDTOV2 {
	this := ApiMatchStateSummaryDTOV2{}
	return &this
}

// NewApiMatchStateSummaryDTOV2WithDefaults instantiates a new ApiMatchStateSummaryDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMatchStateSummaryDTOV2WithDefaults() *ApiMatchStateSummaryDTOV2 {
	this := ApiMatchStateSummaryDTOV2{}
	return &this
}

// GetKnownComponentCount returns the KnownComponentCount field value if set, zero value otherwise.
func (o *ApiMatchStateSummaryDTOV2) GetKnownComponentCount() int32 {
	if o == nil || IsNil(o.KnownComponentCount) {
		var ret int32
		return ret
	}
	return *o.KnownComponentCount
}

// GetKnownComponentCountOk returns a tuple with the KnownComponentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMatchStateSummaryDTOV2) GetKnownComponentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.KnownComponentCount) {
		return nil, false
	}
	return o.KnownComponentCount, true
}

// HasKnownComponentCount returns a boolean if a field has been set.
func (o *ApiMatchStateSummaryDTOV2) HasKnownComponentCount() bool {
	if o != nil && !IsNil(o.KnownComponentCount) {
		return true
	}

	return false
}

// SetKnownComponentCount gets a reference to the given int32 and assigns it to the KnownComponentCount field.
func (o *ApiMatchStateSummaryDTOV2) SetKnownComponentCount(v int32) {
	o.KnownComponentCount = &v
}

// GetTotalComponentCount returns the TotalComponentCount field value if set, zero value otherwise.
func (o *ApiMatchStateSummaryDTOV2) GetTotalComponentCount() int32 {
	if o == nil || IsNil(o.TotalComponentCount) {
		var ret int32
		return ret
	}
	return *o.TotalComponentCount
}

// GetTotalComponentCountOk returns a tuple with the TotalComponentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMatchStateSummaryDTOV2) GetTotalComponentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalComponentCount) {
		return nil, false
	}
	return o.TotalComponentCount, true
}

// HasTotalComponentCount returns a boolean if a field has been set.
func (o *ApiMatchStateSummaryDTOV2) HasTotalComponentCount() bool {
	if o != nil && !IsNil(o.TotalComponentCount) {
		return true
	}

	return false
}

// SetTotalComponentCount gets a reference to the given int32 and assigns it to the TotalComponentCount field.
func (o *ApiMatchStateSummaryDTOV2) SetTotalComponentCount(v int32) {
	o.TotalComponentCount = &v
}

func (o ApiMatchStateSummaryDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMatchStateSummaryDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KnownComponentCount) {
		toSerialize["knownComponentCount"] = o.KnownComponentCount
	}
	if !IsNil(o.TotalComponentCount) {
		toSerialize["totalComponentCount"] = o.TotalComponentCount
	}
	return toSerialize, nil
}

type NullableApiMatchStateSummaryDTOV2 struct {
	value *ApiMatchStateSummaryDTOV2
	isSet bool
}

func (v NullableApiMatchStateSummaryDTOV2) Get() *ApiMatchStateSummaryDTOV2 {
	return v.value
}

func (v *NullableApiMatchStateSummaryDTOV2) Set(val *ApiMatchStateSummaryDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMatchStateSummaryDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMatchStateSummaryDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMatchStateSummaryDTOV2(val *ApiMatchStateSummaryDTOV2) *NullableApiMatchStateSummaryDTOV2 {
	return &NullableApiMatchStateSummaryDTOV2{value: val, isSet: true}
}

func (v NullableApiMatchStateSummaryDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMatchStateSummaryDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


