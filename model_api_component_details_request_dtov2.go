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

// checks if the ApiComponentDetailsRequestDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentDetailsRequestDTOV2{}

// ApiComponentDetailsRequestDTOV2 struct for ApiComponentDetailsRequestDTOV2
type ApiComponentDetailsRequestDTOV2 struct {
	Components []ApiComponentDTOV2 `json:"components,omitempty"`
}

// NewApiComponentDetailsRequestDTOV2 instantiates a new ApiComponentDetailsRequestDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentDetailsRequestDTOV2() *ApiComponentDetailsRequestDTOV2 {
	this := ApiComponentDetailsRequestDTOV2{}
	return &this
}

// NewApiComponentDetailsRequestDTOV2WithDefaults instantiates a new ApiComponentDetailsRequestDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentDetailsRequestDTOV2WithDefaults() *ApiComponentDetailsRequestDTOV2 {
	this := ApiComponentDetailsRequestDTOV2{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ApiComponentDetailsRequestDTOV2) GetComponents() []ApiComponentDTOV2 {
	if o == nil || IsNil(o.Components) {
		var ret []ApiComponentDTOV2
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentDetailsRequestDTOV2) GetComponentsOk() ([]ApiComponentDTOV2, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ApiComponentDetailsRequestDTOV2) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ApiComponentDTOV2 and assigns it to the Components field.
func (o *ApiComponentDetailsRequestDTOV2) SetComponents(v []ApiComponentDTOV2) {
	o.Components = v
}

func (o ApiComponentDetailsRequestDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentDetailsRequestDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	return toSerialize, nil
}

type NullableApiComponentDetailsRequestDTOV2 struct {
	value *ApiComponentDetailsRequestDTOV2
	isSet bool
}

func (v NullableApiComponentDetailsRequestDTOV2) Get() *ApiComponentDetailsRequestDTOV2 {
	return v.value
}

func (v *NullableApiComponentDetailsRequestDTOV2) Set(val *ApiComponentDetailsRequestDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentDetailsRequestDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentDetailsRequestDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentDetailsRequestDTOV2(val *ApiComponentDetailsRequestDTOV2) *NullableApiComponentDetailsRequestDTOV2 {
	return &NullableApiComponentDetailsRequestDTOV2{value: val, isSet: true}
}

func (v NullableApiComponentDetailsRequestDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentDetailsRequestDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


