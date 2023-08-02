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

// checks if the ApiCompositeValueDTOBoolean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCompositeValueDTOBoolean{}

// ApiCompositeValueDTOBoolean struct for ApiCompositeValueDTOBoolean
type ApiCompositeValueDTOBoolean struct {
	ParentName *string `json:"parentName,omitempty"`
	ParentValue *bool `json:"parentValue,omitempty"`
	Value *bool `json:"value,omitempty"`
}

// NewApiCompositeValueDTOBoolean instantiates a new ApiCompositeValueDTOBoolean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCompositeValueDTOBoolean() *ApiCompositeValueDTOBoolean {
	this := ApiCompositeValueDTOBoolean{}
	return &this
}

// NewApiCompositeValueDTOBooleanWithDefaults instantiates a new ApiCompositeValueDTOBoolean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCompositeValueDTOBooleanWithDefaults() *ApiCompositeValueDTOBoolean {
	this := ApiCompositeValueDTOBoolean{}
	return &this
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *ApiCompositeValueDTOBoolean) GetParentName() string {
	if o == nil || IsNil(o.ParentName) {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeValueDTOBoolean) GetParentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentName) {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *ApiCompositeValueDTOBoolean) HasParentName() bool {
	if o != nil && !IsNil(o.ParentName) {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *ApiCompositeValueDTOBoolean) SetParentName(v string) {
	o.ParentName = &v
}

// GetParentValue returns the ParentValue field value if set, zero value otherwise.
func (o *ApiCompositeValueDTOBoolean) GetParentValue() bool {
	if o == nil || IsNil(o.ParentValue) {
		var ret bool
		return ret
	}
	return *o.ParentValue
}

// GetParentValueOk returns a tuple with the ParentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeValueDTOBoolean) GetParentValueOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentValue) {
		return nil, false
	}
	return o.ParentValue, true
}

// HasParentValue returns a boolean if a field has been set.
func (o *ApiCompositeValueDTOBoolean) HasParentValue() bool {
	if o != nil && !IsNil(o.ParentValue) {
		return true
	}

	return false
}

// SetParentValue gets a reference to the given bool and assigns it to the ParentValue field.
func (o *ApiCompositeValueDTOBoolean) SetParentValue(v bool) {
	o.ParentValue = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiCompositeValueDTOBoolean) GetValue() bool {
	if o == nil || IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeValueDTOBoolean) GetValueOk() (*bool, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiCompositeValueDTOBoolean) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *ApiCompositeValueDTOBoolean) SetValue(v bool) {
	o.Value = &v
}

func (o ApiCompositeValueDTOBoolean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCompositeValueDTOBoolean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParentName) {
		toSerialize["parentName"] = o.ParentName
	}
	if !IsNil(o.ParentValue) {
		toSerialize["parentValue"] = o.ParentValue
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApiCompositeValueDTOBoolean struct {
	value *ApiCompositeValueDTOBoolean
	isSet bool
}

func (v NullableApiCompositeValueDTOBoolean) Get() *ApiCompositeValueDTOBoolean {
	return v.value
}

func (v *NullableApiCompositeValueDTOBoolean) Set(val *ApiCompositeValueDTOBoolean) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCompositeValueDTOBoolean) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCompositeValueDTOBoolean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCompositeValueDTOBoolean(val *ApiCompositeValueDTOBoolean) *NullableApiCompositeValueDTOBoolean {
	return &NullableApiCompositeValueDTOBoolean{value: val, isSet: true}
}

func (v NullableApiCompositeValueDTOBoolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCompositeValueDTOBoolean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


