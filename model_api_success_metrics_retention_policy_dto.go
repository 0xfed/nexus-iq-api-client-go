/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiSuccessMetricsRetentionPolicyDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSuccessMetricsRetentionPolicyDTO{}

// ApiSuccessMetricsRetentionPolicyDTO struct for ApiSuccessMetricsRetentionPolicyDTO
type ApiSuccessMetricsRetentionPolicyDTO struct {
	EnablePurging *bool `json:"enablePurging,omitempty"`
	InheritPolicy *bool `json:"inheritPolicy,omitempty"`
	MaxAge *string `json:"maxAge,omitempty"`
}

// NewApiSuccessMetricsRetentionPolicyDTO instantiates a new ApiSuccessMetricsRetentionPolicyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSuccessMetricsRetentionPolicyDTO() *ApiSuccessMetricsRetentionPolicyDTO {
	this := ApiSuccessMetricsRetentionPolicyDTO{}
	return &this
}

// NewApiSuccessMetricsRetentionPolicyDTOWithDefaults instantiates a new ApiSuccessMetricsRetentionPolicyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSuccessMetricsRetentionPolicyDTOWithDefaults() *ApiSuccessMetricsRetentionPolicyDTO {
	this := ApiSuccessMetricsRetentionPolicyDTO{}
	return &this
}

// GetEnablePurging returns the EnablePurging field value if set, zero value otherwise.
func (o *ApiSuccessMetricsRetentionPolicyDTO) GetEnablePurging() bool {
	if o == nil || IsNil(o.EnablePurging) {
		var ret bool
		return ret
	}
	return *o.EnablePurging
}

// GetEnablePurgingOk returns a tuple with the EnablePurging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSuccessMetricsRetentionPolicyDTO) GetEnablePurgingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePurging) {
		return nil, false
	}
	return o.EnablePurging, true
}

// HasEnablePurging returns a boolean if a field has been set.
func (o *ApiSuccessMetricsRetentionPolicyDTO) HasEnablePurging() bool {
	if o != nil && !IsNil(o.EnablePurging) {
		return true
	}

	return false
}

// SetEnablePurging gets a reference to the given bool and assigns it to the EnablePurging field.
func (o *ApiSuccessMetricsRetentionPolicyDTO) SetEnablePurging(v bool) {
	o.EnablePurging = &v
}

// GetInheritPolicy returns the InheritPolicy field value if set, zero value otherwise.
func (o *ApiSuccessMetricsRetentionPolicyDTO) GetInheritPolicy() bool {
	if o == nil || IsNil(o.InheritPolicy) {
		var ret bool
		return ret
	}
	return *o.InheritPolicy
}

// GetInheritPolicyOk returns a tuple with the InheritPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSuccessMetricsRetentionPolicyDTO) GetInheritPolicyOk() (*bool, bool) {
	if o == nil || IsNil(o.InheritPolicy) {
		return nil, false
	}
	return o.InheritPolicy, true
}

// HasInheritPolicy returns a boolean if a field has been set.
func (o *ApiSuccessMetricsRetentionPolicyDTO) HasInheritPolicy() bool {
	if o != nil && !IsNil(o.InheritPolicy) {
		return true
	}

	return false
}

// SetInheritPolicy gets a reference to the given bool and assigns it to the InheritPolicy field.
func (o *ApiSuccessMetricsRetentionPolicyDTO) SetInheritPolicy(v bool) {
	o.InheritPolicy = &v
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *ApiSuccessMetricsRetentionPolicyDTO) GetMaxAge() string {
	if o == nil || IsNil(o.MaxAge) {
		var ret string
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSuccessMetricsRetentionPolicyDTO) GetMaxAgeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *ApiSuccessMetricsRetentionPolicyDTO) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given string and assigns it to the MaxAge field.
func (o *ApiSuccessMetricsRetentionPolicyDTO) SetMaxAge(v string) {
	o.MaxAge = &v
}

func (o ApiSuccessMetricsRetentionPolicyDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSuccessMetricsRetentionPolicyDTO) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableApiSuccessMetricsRetentionPolicyDTO struct {
	value *ApiSuccessMetricsRetentionPolicyDTO
	isSet bool
}

func (v NullableApiSuccessMetricsRetentionPolicyDTO) Get() *ApiSuccessMetricsRetentionPolicyDTO {
	return v.value
}

func (v *NullableApiSuccessMetricsRetentionPolicyDTO) Set(val *ApiSuccessMetricsRetentionPolicyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSuccessMetricsRetentionPolicyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSuccessMetricsRetentionPolicyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSuccessMetricsRetentionPolicyDTO(val *ApiSuccessMetricsRetentionPolicyDTO) *NullableApiSuccessMetricsRetentionPolicyDTO {
	return &NullableApiSuccessMetricsRetentionPolicyDTO{value: val, isSet: true}
}

func (v NullableApiSuccessMetricsRetentionPolicyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSuccessMetricsRetentionPolicyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


