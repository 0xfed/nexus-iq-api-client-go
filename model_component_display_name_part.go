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

// checks if the ComponentDisplayNamePart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentDisplayNamePart{}

// ComponentDisplayNamePart struct for ComponentDisplayNamePart
type ComponentDisplayNamePart struct {
	Field *string `json:"field,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewComponentDisplayNamePart instantiates a new ComponentDisplayNamePart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentDisplayNamePart() *ComponentDisplayNamePart {
	this := ComponentDisplayNamePart{}
	return &this
}

// NewComponentDisplayNamePartWithDefaults instantiates a new ComponentDisplayNamePart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentDisplayNamePartWithDefaults() *ComponentDisplayNamePart {
	this := ComponentDisplayNamePart{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ComponentDisplayNamePart) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDisplayNamePart) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ComponentDisplayNamePart) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ComponentDisplayNamePart) SetField(v string) {
	o.Field = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ComponentDisplayNamePart) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDisplayNamePart) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ComponentDisplayNamePart) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ComponentDisplayNamePart) SetValue(v string) {
	o.Value = &v
}

func (o ComponentDisplayNamePart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentDisplayNamePart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableComponentDisplayNamePart struct {
	value *ComponentDisplayNamePart
	isSet bool
}

func (v NullableComponentDisplayNamePart) Get() *ComponentDisplayNamePart {
	return v.value
}

func (v *NullableComponentDisplayNamePart) Set(val *ComponentDisplayNamePart) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentDisplayNamePart) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentDisplayNamePart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentDisplayNamePart(val *ComponentDisplayNamePart) *NullableComponentDisplayNamePart {
	return &NullableComponentDisplayNamePart{value: val, isSet: true}
}

func (v NullableComponentDisplayNamePart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentDisplayNamePart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


