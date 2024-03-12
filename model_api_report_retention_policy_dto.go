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

// checks if the ApiReportRetentionPolicyDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReportRetentionPolicyDTO{}

// ApiReportRetentionPolicyDTO struct for ApiReportRetentionPolicyDTO
type ApiReportRetentionPolicyDTO struct {
	EnablePurging *bool `json:"enablePurging,omitempty"`
	InheritPolicy *bool `json:"inheritPolicy,omitempty"`
	MaxAge *string `json:"maxAge,omitempty"`
	MaxCount *int32 `json:"maxCount,omitempty"`
}

// NewApiReportRetentionPolicyDTO instantiates a new ApiReportRetentionPolicyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportRetentionPolicyDTO() *ApiReportRetentionPolicyDTO {
	this := ApiReportRetentionPolicyDTO{}
	return &this
}

// NewApiReportRetentionPolicyDTOWithDefaults instantiates a new ApiReportRetentionPolicyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportRetentionPolicyDTOWithDefaults() *ApiReportRetentionPolicyDTO {
	this := ApiReportRetentionPolicyDTO{}
	return &this
}

// GetEnablePurging returns the EnablePurging field value if set, zero value otherwise.
func (o *ApiReportRetentionPolicyDTO) GetEnablePurging() bool {
	if o == nil || IsNil(o.EnablePurging) {
		var ret bool
		return ret
	}
	return *o.EnablePurging
}

// GetEnablePurgingOk returns a tuple with the EnablePurging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportRetentionPolicyDTO) GetEnablePurgingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePurging) {
		return nil, false
	}
	return o.EnablePurging, true
}

// HasEnablePurging returns a boolean if a field has been set.
func (o *ApiReportRetentionPolicyDTO) HasEnablePurging() bool {
	if o != nil && !IsNil(o.EnablePurging) {
		return true
	}

	return false
}

// SetEnablePurging gets a reference to the given bool and assigns it to the EnablePurging field.
func (o *ApiReportRetentionPolicyDTO) SetEnablePurging(v bool) {
	o.EnablePurging = &v
}

// GetInheritPolicy returns the InheritPolicy field value if set, zero value otherwise.
func (o *ApiReportRetentionPolicyDTO) GetInheritPolicy() bool {
	if o == nil || IsNil(o.InheritPolicy) {
		var ret bool
		return ret
	}
	return *o.InheritPolicy
}

// GetInheritPolicyOk returns a tuple with the InheritPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportRetentionPolicyDTO) GetInheritPolicyOk() (*bool, bool) {
	if o == nil || IsNil(o.InheritPolicy) {
		return nil, false
	}
	return o.InheritPolicy, true
}

// HasInheritPolicy returns a boolean if a field has been set.
func (o *ApiReportRetentionPolicyDTO) HasInheritPolicy() bool {
	if o != nil && !IsNil(o.InheritPolicy) {
		return true
	}

	return false
}

// SetInheritPolicy gets a reference to the given bool and assigns it to the InheritPolicy field.
func (o *ApiReportRetentionPolicyDTO) SetInheritPolicy(v bool) {
	o.InheritPolicy = &v
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *ApiReportRetentionPolicyDTO) GetMaxAge() string {
	if o == nil || IsNil(o.MaxAge) {
		var ret string
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportRetentionPolicyDTO) GetMaxAgeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *ApiReportRetentionPolicyDTO) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given string and assigns it to the MaxAge field.
func (o *ApiReportRetentionPolicyDTO) SetMaxAge(v string) {
	o.MaxAge = &v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *ApiReportRetentionPolicyDTO) GetMaxCount() int32 {
	if o == nil || IsNil(o.MaxCount) {
		var ret int32
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportRetentionPolicyDTO) GetMaxCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxCount) {
		return nil, false
	}
	return o.MaxCount, true
}

// HasMaxCount returns a boolean if a field has been set.
func (o *ApiReportRetentionPolicyDTO) HasMaxCount() bool {
	if o != nil && !IsNil(o.MaxCount) {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int32 and assigns it to the MaxCount field.
func (o *ApiReportRetentionPolicyDTO) SetMaxCount(v int32) {
	o.MaxCount = &v
}

func (o ApiReportRetentionPolicyDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportRetentionPolicyDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnablePurging) {
		toSerialize["enablePurging"] = o.EnablePurging
	}
	if !IsNil(o.InheritPolicy) {
		toSerialize["inheritPolicy"] = o.InheritPolicy
	}
	if !IsNil(o.MaxAge) {
		toSerialize["maxAge"] = o.MaxAge
	}
	if !IsNil(o.MaxCount) {
		toSerialize["maxCount"] = o.MaxCount
	}
	return toSerialize, nil
}

type NullableApiReportRetentionPolicyDTO struct {
	value *ApiReportRetentionPolicyDTO
	isSet bool
}

func (v NullableApiReportRetentionPolicyDTO) Get() *ApiReportRetentionPolicyDTO {
	return v.value
}

func (v *NullableApiReportRetentionPolicyDTO) Set(val *ApiReportRetentionPolicyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportRetentionPolicyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportRetentionPolicyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportRetentionPolicyDTO(val *ApiReportRetentionPolicyDTO) *NullableApiReportRetentionPolicyDTO {
	return &NullableApiReportRetentionPolicyDTO{value: val, isSet: true}
}

func (v NullableApiReportRetentionPolicyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportRetentionPolicyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


