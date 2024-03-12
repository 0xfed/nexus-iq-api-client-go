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

// checks if the ApiApplicationListDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiApplicationListDTO{}

// ApiApplicationListDTO struct for ApiApplicationListDTO
type ApiApplicationListDTO struct {
	Applications []ApiApplicationDTO `json:"applications,omitempty"`
}

// NewApiApplicationListDTO instantiates a new ApiApplicationListDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiApplicationListDTO() *ApiApplicationListDTO {
	this := ApiApplicationListDTO{}
	return &this
}

// NewApiApplicationListDTOWithDefaults instantiates a new ApiApplicationListDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiApplicationListDTOWithDefaults() *ApiApplicationListDTO {
	this := ApiApplicationListDTO{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *ApiApplicationListDTO) GetApplications() []ApiApplicationDTO {
	if o == nil || IsNil(o.Applications) {
		var ret []ApiApplicationDTO
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationListDTO) GetApplicationsOk() ([]ApiApplicationDTO, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *ApiApplicationListDTO) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []ApiApplicationDTO and assigns it to the Applications field.
func (o *ApiApplicationListDTO) SetApplications(v []ApiApplicationDTO) {
	o.Applications = v
}

func (o ApiApplicationListDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiApplicationListDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	return toSerialize, nil
}

type NullableApiApplicationListDTO struct {
	value *ApiApplicationListDTO
	isSet bool
}

func (v NullableApiApplicationListDTO) Get() *ApiApplicationListDTO {
	return v.value
}

func (v *NullableApiApplicationListDTO) Set(val *ApiApplicationListDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiApplicationListDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiApplicationListDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiApplicationListDTO(val *ApiApplicationListDTO) *NullableApiApplicationListDTO {
	return &NullableApiApplicationListDTO{value: val, isSet: true}
}

func (v NullableApiApplicationListDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiApplicationListDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


