/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.181.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the LegacyViolationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegacyViolationsResponse{}

// LegacyViolationsResponse struct for LegacyViolationsResponse
type LegacyViolationsResponse struct {
	AllowChange *bool `json:"allowChange,omitempty"`
	AllowOverride *bool `json:"allowOverride,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	EnabledInParent NullableBool `json:"enabledInParent,omitempty"`
	InheritedFromOrganizationName NullableString `json:"inheritedFromOrganizationName,omitempty"`
}

// NewLegacyViolationsResponse instantiates a new LegacyViolationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyViolationsResponse() *LegacyViolationsResponse {
	this := LegacyViolationsResponse{}
	return &this
}

// NewLegacyViolationsResponseWithDefaults instantiates a new LegacyViolationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyViolationsResponseWithDefaults() *LegacyViolationsResponse {
	this := LegacyViolationsResponse{}
	return &this
}

// GetAllowChange returns the AllowChange field value if set, zero value otherwise.
func (o *LegacyViolationsResponse) GetAllowChange() bool {
	if o == nil || IsNil(o.AllowChange) {
		var ret bool
		return ret
	}
	return *o.AllowChange
}

// GetAllowChangeOk returns a tuple with the AllowChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyViolationsResponse) GetAllowChangeOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowChange) {
		return nil, false
	}
	return o.AllowChange, true
}

// HasAllowChange returns a boolean if a field has been set.
func (o *LegacyViolationsResponse) HasAllowChange() bool {
	if o != nil && !IsNil(o.AllowChange) {
		return true
	}

	return false
}

// SetAllowChange gets a reference to the given bool and assigns it to the AllowChange field.
func (o *LegacyViolationsResponse) SetAllowChange(v bool) {
	o.AllowChange = &v
}

// GetAllowOverride returns the AllowOverride field value if set, zero value otherwise.
func (o *LegacyViolationsResponse) GetAllowOverride() bool {
	if o == nil || IsNil(o.AllowOverride) {
		var ret bool
		return ret
	}
	return *o.AllowOverride
}

// GetAllowOverrideOk returns a tuple with the AllowOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyViolationsResponse) GetAllowOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowOverride) {
		return nil, false
	}
	return o.AllowOverride, true
}

// HasAllowOverride returns a boolean if a field has been set.
func (o *LegacyViolationsResponse) HasAllowOverride() bool {
	if o != nil && !IsNil(o.AllowOverride) {
		return true
	}

	return false
}

// SetAllowOverride gets a reference to the given bool and assigns it to the AllowOverride field.
func (o *LegacyViolationsResponse) SetAllowOverride(v bool) {
	o.AllowOverride = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LegacyViolationsResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyViolationsResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LegacyViolationsResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LegacyViolationsResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnabledInParent returns the EnabledInParent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LegacyViolationsResponse) GetEnabledInParent() bool {
	if o == nil || IsNil(o.EnabledInParent.Get()) {
		var ret bool
		return ret
	}
	return *o.EnabledInParent.Get()
}

// GetEnabledInParentOk returns a tuple with the EnabledInParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LegacyViolationsResponse) GetEnabledInParentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnabledInParent.Get(), o.EnabledInParent.IsSet()
}

// HasEnabledInParent returns a boolean if a field has been set.
func (o *LegacyViolationsResponse) HasEnabledInParent() bool {
	if o != nil && o.EnabledInParent.IsSet() {
		return true
	}

	return false
}

// SetEnabledInParent gets a reference to the given NullableBool and assigns it to the EnabledInParent field.
func (o *LegacyViolationsResponse) SetEnabledInParent(v bool) {
	o.EnabledInParent.Set(&v)
}
// SetEnabledInParentNil sets the value for EnabledInParent to be an explicit nil
func (o *LegacyViolationsResponse) SetEnabledInParentNil() {
	o.EnabledInParent.Set(nil)
}

// UnsetEnabledInParent ensures that no value is present for EnabledInParent, not even an explicit nil
func (o *LegacyViolationsResponse) UnsetEnabledInParent() {
	o.EnabledInParent.Unset()
}

// GetInheritedFromOrganizationName returns the InheritedFromOrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LegacyViolationsResponse) GetInheritedFromOrganizationName() string {
	if o == nil || IsNil(o.InheritedFromOrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.InheritedFromOrganizationName.Get()
}

// GetInheritedFromOrganizationNameOk returns a tuple with the InheritedFromOrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LegacyViolationsResponse) GetInheritedFromOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InheritedFromOrganizationName.Get(), o.InheritedFromOrganizationName.IsSet()
}

// HasInheritedFromOrganizationName returns a boolean if a field has been set.
func (o *LegacyViolationsResponse) HasInheritedFromOrganizationName() bool {
	if o != nil && o.InheritedFromOrganizationName.IsSet() {
		return true
	}

	return false
}

// SetInheritedFromOrganizationName gets a reference to the given NullableString and assigns it to the InheritedFromOrganizationName field.
func (o *LegacyViolationsResponse) SetInheritedFromOrganizationName(v string) {
	o.InheritedFromOrganizationName.Set(&v)
}
// SetInheritedFromOrganizationNameNil sets the value for InheritedFromOrganizationName to be an explicit nil
func (o *LegacyViolationsResponse) SetInheritedFromOrganizationNameNil() {
	o.InheritedFromOrganizationName.Set(nil)
}

// UnsetInheritedFromOrganizationName ensures that no value is present for InheritedFromOrganizationName, not even an explicit nil
func (o *LegacyViolationsResponse) UnsetInheritedFromOrganizationName() {
	o.InheritedFromOrganizationName.Unset()
}

func (o LegacyViolationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegacyViolationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowChange) {
		toSerialize["allowChange"] = o.AllowChange
	}
	if !IsNil(o.AllowOverride) {
		toSerialize["allowOverride"] = o.AllowOverride
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.EnabledInParent.IsSet() {
		toSerialize["enabledInParent"] = o.EnabledInParent.Get()
	}
	if o.InheritedFromOrganizationName.IsSet() {
		toSerialize["inheritedFromOrganizationName"] = o.InheritedFromOrganizationName.Get()
	}
	return toSerialize, nil
}

type NullableLegacyViolationsResponse struct {
	value *LegacyViolationsResponse
	isSet bool
}

func (v NullableLegacyViolationsResponse) Get() *LegacyViolationsResponse {
	return v.value
}

func (v *NullableLegacyViolationsResponse) Set(val *LegacyViolationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyViolationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyViolationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyViolationsResponse(val *LegacyViolationsResponse) *NullableLegacyViolationsResponse {
	return &NullableLegacyViolationsResponse{value: val, isSet: true}
}

func (v NullableLegacyViolationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyViolationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


