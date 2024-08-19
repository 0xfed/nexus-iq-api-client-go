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

// checks if the Action type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Action{}

// Action struct for Action
type Action struct {
	ActionTypeId *string `json:"actionTypeId,omitempty"`
	Target *string `json:"target,omitempty"`
	TargetType *string `json:"targetType,omitempty"`
}

// NewAction instantiates a new Action object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAction() *Action {
	this := Action{}
	return &this
}

// NewActionWithDefaults instantiates a new Action object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionWithDefaults() *Action {
	this := Action{}
	return &this
}

// GetActionTypeId returns the ActionTypeId field value if set, zero value otherwise.
func (o *Action) GetActionTypeId() string {
	if o == nil || IsNil(o.ActionTypeId) {
		var ret string
		return ret
	}
	return *o.ActionTypeId
}

// GetActionTypeIdOk returns a tuple with the ActionTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetActionTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActionTypeId) {
		return nil, false
	}
	return o.ActionTypeId, true
}

// HasActionTypeId returns a boolean if a field has been set.
func (o *Action) HasActionTypeId() bool {
	if o != nil && !IsNil(o.ActionTypeId) {
		return true
	}

	return false
}

// SetActionTypeId gets a reference to the given string and assigns it to the ActionTypeId field.
func (o *Action) SetActionTypeId(v string) {
	o.ActionTypeId = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Action) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Action) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *Action) SetTarget(v string) {
	o.Target = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *Action) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *Action) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *Action) SetTargetType(v string) {
	o.TargetType = &v
}

func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Action) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionTypeId) {
		toSerialize["actionTypeId"] = o.ActionTypeId
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.TargetType) {
		toSerialize["targetType"] = o.TargetType
	}
	return toSerialize, nil
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


