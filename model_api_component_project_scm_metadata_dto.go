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

// checks if the ApiComponentProjectScmMetadataDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentProjectScmMetadataDTO{}

// ApiComponentProjectScmMetadataDTO struct for ApiComponentProjectScmMetadataDTO
type ApiComponentProjectScmMetadataDTO struct {
	Forks *int32 `json:"forks,omitempty"`
	Stars *int32 `json:"stars,omitempty"`
}

// NewApiComponentProjectScmMetadataDTO instantiates a new ApiComponentProjectScmMetadataDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentProjectScmMetadataDTO() *ApiComponentProjectScmMetadataDTO {
	this := ApiComponentProjectScmMetadataDTO{}
	return &this
}

// NewApiComponentProjectScmMetadataDTOWithDefaults instantiates a new ApiComponentProjectScmMetadataDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentProjectScmMetadataDTOWithDefaults() *ApiComponentProjectScmMetadataDTO {
	this := ApiComponentProjectScmMetadataDTO{}
	return &this
}

// GetForks returns the Forks field value if set, zero value otherwise.
func (o *ApiComponentProjectScmMetadataDTO) GetForks() int32 {
	if o == nil || IsNil(o.Forks) {
		var ret int32
		return ret
	}
	return *o.Forks
}

// GetForksOk returns a tuple with the Forks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectScmMetadataDTO) GetForksOk() (*int32, bool) {
	if o == nil || IsNil(o.Forks) {
		return nil, false
	}
	return o.Forks, true
}

// HasForks returns a boolean if a field has been set.
func (o *ApiComponentProjectScmMetadataDTO) HasForks() bool {
	if o != nil && !IsNil(o.Forks) {
		return true
	}

	return false
}

// SetForks gets a reference to the given int32 and assigns it to the Forks field.
func (o *ApiComponentProjectScmMetadataDTO) SetForks(v int32) {
	o.Forks = &v
}

// GetStars returns the Stars field value if set, zero value otherwise.
func (o *ApiComponentProjectScmMetadataDTO) GetStars() int32 {
	if o == nil || IsNil(o.Stars) {
		var ret int32
		return ret
	}
	return *o.Stars
}

// GetStarsOk returns a tuple with the Stars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectScmMetadataDTO) GetStarsOk() (*int32, bool) {
	if o == nil || IsNil(o.Stars) {
		return nil, false
	}
	return o.Stars, true
}

// HasStars returns a boolean if a field has been set.
func (o *ApiComponentProjectScmMetadataDTO) HasStars() bool {
	if o != nil && !IsNil(o.Stars) {
		return true
	}

	return false
}

// SetStars gets a reference to the given int32 and assigns it to the Stars field.
func (o *ApiComponentProjectScmMetadataDTO) SetStars(v int32) {
	o.Stars = &v
}

func (o ApiComponentProjectScmMetadataDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentProjectScmMetadataDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Forks) {
		toSerialize["forks"] = o.Forks
	}
	if !IsNil(o.Stars) {
		toSerialize["stars"] = o.Stars
	}
	return toSerialize, nil
}

type NullableApiComponentProjectScmMetadataDTO struct {
	value *ApiComponentProjectScmMetadataDTO
	isSet bool
}

func (v NullableApiComponentProjectScmMetadataDTO) Get() *ApiComponentProjectScmMetadataDTO {
	return v.value
}

func (v *NullableApiComponentProjectScmMetadataDTO) Set(val *ApiComponentProjectScmMetadataDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentProjectScmMetadataDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentProjectScmMetadataDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentProjectScmMetadataDTO(val *ApiComponentProjectScmMetadataDTO) *NullableApiComponentProjectScmMetadataDTO {
	return &NullableApiComponentProjectScmMetadataDTO{value: val, isSet: true}
}

func (v NullableApiComponentProjectScmMetadataDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentProjectScmMetadataDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


