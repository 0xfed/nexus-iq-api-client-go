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

// checks if the BodyPartMediaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BodyPartMediaType{}

// BodyPartMediaType struct for BodyPartMediaType
type BodyPartMediaType struct {
	Parameters *map[string]string `json:"parameters,omitempty"`
	Subtype *string `json:"subtype,omitempty"`
	Type *string `json:"type,omitempty"`
	WildcardSubtype *bool `json:"wildcardSubtype,omitempty"`
	WildcardType *bool `json:"wildcardType,omitempty"`
}

// NewBodyPartMediaType instantiates a new BodyPartMediaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBodyPartMediaType() *BodyPartMediaType {
	this := BodyPartMediaType{}
	return &this
}

// NewBodyPartMediaTypeWithDefaults instantiates a new BodyPartMediaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBodyPartMediaTypeWithDefaults() *BodyPartMediaType {
	this := BodyPartMediaType{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetParameters() map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *BodyPartMediaType) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BodyPartMediaType) SetSubtype(v string) {
	o.Subtype = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BodyPartMediaType) SetType(v string) {
	o.Type = &v
}

// GetWildcardSubtype returns the WildcardSubtype field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetWildcardSubtype() bool {
	if o == nil || IsNil(o.WildcardSubtype) {
		var ret bool
		return ret
	}
	return *o.WildcardSubtype
}

// GetWildcardSubtypeOk returns a tuple with the WildcardSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetWildcardSubtypeOk() (*bool, bool) {
	if o == nil || IsNil(o.WildcardSubtype) {
		return nil, false
	}
	return o.WildcardSubtype, true
}

// HasWildcardSubtype returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasWildcardSubtype() bool {
	if o != nil && !IsNil(o.WildcardSubtype) {
		return true
	}

	return false
}

// SetWildcardSubtype gets a reference to the given bool and assigns it to the WildcardSubtype field.
func (o *BodyPartMediaType) SetWildcardSubtype(v bool) {
	o.WildcardSubtype = &v
}

// GetWildcardType returns the WildcardType field value if set, zero value otherwise.
func (o *BodyPartMediaType) GetWildcardType() bool {
	if o == nil || IsNil(o.WildcardType) {
		var ret bool
		return ret
	}
	return *o.WildcardType
}

// GetWildcardTypeOk returns a tuple with the WildcardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPartMediaType) GetWildcardTypeOk() (*bool, bool) {
	if o == nil || IsNil(o.WildcardType) {
		return nil, false
	}
	return o.WildcardType, true
}

// HasWildcardType returns a boolean if a field has been set.
func (o *BodyPartMediaType) HasWildcardType() bool {
	if o != nil && !IsNil(o.WildcardType) {
		return true
	}

	return false
}

// SetWildcardType gets a reference to the given bool and assigns it to the WildcardType field.
func (o *BodyPartMediaType) SetWildcardType(v bool) {
	o.WildcardType = &v
}

func (o BodyPartMediaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BodyPartMediaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.WildcardSubtype) {
		toSerialize["wildcardSubtype"] = o.WildcardSubtype
	}
	if !IsNil(o.WildcardType) {
		toSerialize["wildcardType"] = o.WildcardType
	}
	return toSerialize, nil
}

type NullableBodyPartMediaType struct {
	value *BodyPartMediaType
	isSet bool
}

func (v NullableBodyPartMediaType) Get() *BodyPartMediaType {
	return v.value
}

func (v *NullableBodyPartMediaType) Set(val *BodyPartMediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableBodyPartMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableBodyPartMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBodyPartMediaType(val *BodyPartMediaType) *NullableBodyPartMediaType {
	return &NullableBodyPartMediaType{value: val, isSet: true}
}

func (v NullableBodyPartMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBodyPartMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


