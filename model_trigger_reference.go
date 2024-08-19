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

// checks if the TriggerReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerReference{}

// TriggerReference struct for TriggerReference
type TriggerReference struct {
	Type *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewTriggerReference instantiates a new TriggerReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerReference() *TriggerReference {
	this := TriggerReference{}
	return &this
}

// NewTriggerReferenceWithDefaults instantiates a new TriggerReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerReferenceWithDefaults() *TriggerReference {
	this := TriggerReference{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TriggerReference) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerReference) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TriggerReference) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TriggerReference) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TriggerReference) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerReference) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TriggerReference) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TriggerReference) SetValue(v string) {
	o.Value = &v
}

func (o TriggerReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTriggerReference struct {
	value *TriggerReference
	isSet bool
}

func (v NullableTriggerReference) Get() *TriggerReference {
	return v.value
}

func (v *NullableTriggerReference) Set(val *TriggerReference) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerReference) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerReference(val *TriggerReference) *NullableTriggerReference {
	return &NullableTriggerReference{value: val, isSet: true}
}

func (v NullableTriggerReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


