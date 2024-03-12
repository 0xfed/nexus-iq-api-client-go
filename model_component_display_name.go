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

// checks if the ComponentDisplayName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentDisplayName{}

// ComponentDisplayName struct for ComponentDisplayName
type ComponentDisplayName struct {
	Name *string `json:"name,omitempty"`
	Parts []ComponentDisplayNamePart `json:"parts,omitempty"`
}

// NewComponentDisplayName instantiates a new ComponentDisplayName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentDisplayName() *ComponentDisplayName {
	this := ComponentDisplayName{}
	return &this
}

// NewComponentDisplayNameWithDefaults instantiates a new ComponentDisplayName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentDisplayNameWithDefaults() *ComponentDisplayName {
	this := ComponentDisplayName{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentDisplayName) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDisplayName) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentDisplayName) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentDisplayName) SetName(v string) {
	o.Name = &v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *ComponentDisplayName) GetParts() []ComponentDisplayNamePart {
	if o == nil || IsNil(o.Parts) {
		var ret []ComponentDisplayNamePart
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentDisplayName) GetPartsOk() ([]ComponentDisplayNamePart, bool) {
	if o == nil || IsNil(o.Parts) {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *ComponentDisplayName) HasParts() bool {
	if o != nil && !IsNil(o.Parts) {
		return true
	}

	return false
}

// SetParts gets a reference to the given []ComponentDisplayNamePart and assigns it to the Parts field.
func (o *ComponentDisplayName) SetParts(v []ComponentDisplayNamePart) {
	o.Parts = v
}

func (o ComponentDisplayName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentDisplayName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parts) {
		toSerialize["parts"] = o.Parts
	}
	return toSerialize, nil
}

type NullableComponentDisplayName struct {
	value *ComponentDisplayName
	isSet bool
}

func (v NullableComponentDisplayName) Get() *ComponentDisplayName {
	return v.value
}

func (v *NullableComponentDisplayName) Set(val *ComponentDisplayName) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentDisplayName) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentDisplayName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentDisplayName(val *ComponentDisplayName) *NullableComponentDisplayName {
	return &NullableComponentDisplayName{value: val, isSet: true}
}

func (v NullableComponentDisplayName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentDisplayName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


