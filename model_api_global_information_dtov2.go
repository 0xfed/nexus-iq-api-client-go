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

// checks if the ApiGlobalInformationDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiGlobalInformationDTOV2{}

// ApiGlobalInformationDTOV2 struct for ApiGlobalInformationDTOV2
type ApiGlobalInformationDTOV2 struct {
	DataVersionDate *string `json:"dataVersionDate,omitempty"`
}

// NewApiGlobalInformationDTOV2 instantiates a new ApiGlobalInformationDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiGlobalInformationDTOV2() *ApiGlobalInformationDTOV2 {
	this := ApiGlobalInformationDTOV2{}
	return &this
}

// NewApiGlobalInformationDTOV2WithDefaults instantiates a new ApiGlobalInformationDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiGlobalInformationDTOV2WithDefaults() *ApiGlobalInformationDTOV2 {
	this := ApiGlobalInformationDTOV2{}
	return &this
}

// GetDataVersionDate returns the DataVersionDate field value if set, zero value otherwise.
func (o *ApiGlobalInformationDTOV2) GetDataVersionDate() string {
	if o == nil || IsNil(o.DataVersionDate) {
		var ret string
		return ret
	}
	return *o.DataVersionDate
}

// GetDataVersionDateOk returns a tuple with the DataVersionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiGlobalInformationDTOV2) GetDataVersionDateOk() (*string, bool) {
	if o == nil || IsNil(o.DataVersionDate) {
		return nil, false
	}
	return o.DataVersionDate, true
}

// HasDataVersionDate returns a boolean if a field has been set.
func (o *ApiGlobalInformationDTOV2) HasDataVersionDate() bool {
	if o != nil && !IsNil(o.DataVersionDate) {
		return true
	}

	return false
}

// SetDataVersionDate gets a reference to the given string and assigns it to the DataVersionDate field.
func (o *ApiGlobalInformationDTOV2) SetDataVersionDate(v string) {
	o.DataVersionDate = &v
}

func (o ApiGlobalInformationDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiGlobalInformationDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataVersionDate) {
		toSerialize["dataVersionDate"] = o.DataVersionDate
	}
	return toSerialize, nil
}

type NullableApiGlobalInformationDTOV2 struct {
	value *ApiGlobalInformationDTOV2
	isSet bool
}

func (v NullableApiGlobalInformationDTOV2) Get() *ApiGlobalInformationDTOV2 {
	return v.value
}

func (v *NullableApiGlobalInformationDTOV2) Set(val *ApiGlobalInformationDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiGlobalInformationDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiGlobalInformationDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiGlobalInformationDTOV2(val *ApiGlobalInformationDTOV2) *NullableApiGlobalInformationDTOV2 {
	return &NullableApiGlobalInformationDTOV2{value: val, isSet: true}
}

func (v NullableApiGlobalInformationDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiGlobalInformationDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


