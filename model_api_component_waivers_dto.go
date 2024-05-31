/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.176.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiComponentWaiversDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentWaiversDTO{}

// ApiComponentWaiversDTO struct for ApiComponentWaiversDTO
type ApiComponentWaiversDTO struct {
	ApplicationWaivers []ApiApplicationWaiverDTO `json:"applicationWaivers,omitempty"`
	RepositoryWaivers []ApiRepositoryWaiverDTO `json:"repositoryWaivers,omitempty"`
}

// NewApiComponentWaiversDTO instantiates a new ApiComponentWaiversDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentWaiversDTO() *ApiComponentWaiversDTO {
	this := ApiComponentWaiversDTO{}
	return &this
}

// NewApiComponentWaiversDTOWithDefaults instantiates a new ApiComponentWaiversDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentWaiversDTOWithDefaults() *ApiComponentWaiversDTO {
	this := ApiComponentWaiversDTO{}
	return &this
}

// GetApplicationWaivers returns the ApplicationWaivers field value if set, zero value otherwise.
func (o *ApiComponentWaiversDTO) GetApplicationWaivers() []ApiApplicationWaiverDTO {
	if o == nil || IsNil(o.ApplicationWaivers) {
		var ret []ApiApplicationWaiverDTO
		return ret
	}
	return o.ApplicationWaivers
}

// GetApplicationWaiversOk returns a tuple with the ApplicationWaivers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentWaiversDTO) GetApplicationWaiversOk() ([]ApiApplicationWaiverDTO, bool) {
	if o == nil || IsNil(o.ApplicationWaivers) {
		return nil, false
	}
	return o.ApplicationWaivers, true
}

// HasApplicationWaivers returns a boolean if a field has been set.
func (o *ApiComponentWaiversDTO) HasApplicationWaivers() bool {
	if o != nil && !IsNil(o.ApplicationWaivers) {
		return true
	}

	return false
}

// SetApplicationWaivers gets a reference to the given []ApiApplicationWaiverDTO and assigns it to the ApplicationWaivers field.
func (o *ApiComponentWaiversDTO) SetApplicationWaivers(v []ApiApplicationWaiverDTO) {
	o.ApplicationWaivers = v
}

// GetRepositoryWaivers returns the RepositoryWaivers field value if set, zero value otherwise.
func (o *ApiComponentWaiversDTO) GetRepositoryWaivers() []ApiRepositoryWaiverDTO {
	if o == nil || IsNil(o.RepositoryWaivers) {
		var ret []ApiRepositoryWaiverDTO
		return ret
	}
	return o.RepositoryWaivers
}

// GetRepositoryWaiversOk returns a tuple with the RepositoryWaivers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentWaiversDTO) GetRepositoryWaiversOk() ([]ApiRepositoryWaiverDTO, bool) {
	if o == nil || IsNil(o.RepositoryWaivers) {
		return nil, false
	}
	return o.RepositoryWaivers, true
}

// HasRepositoryWaivers returns a boolean if a field has been set.
func (o *ApiComponentWaiversDTO) HasRepositoryWaivers() bool {
	if o != nil && !IsNil(o.RepositoryWaivers) {
		return true
	}

	return false
}

// SetRepositoryWaivers gets a reference to the given []ApiRepositoryWaiverDTO and assigns it to the RepositoryWaivers field.
func (o *ApiComponentWaiversDTO) SetRepositoryWaivers(v []ApiRepositoryWaiverDTO) {
	o.RepositoryWaivers = v
}

func (o ApiComponentWaiversDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentWaiversDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationWaivers) {
		toSerialize["applicationWaivers"] = o.ApplicationWaivers
	}
	if !IsNil(o.RepositoryWaivers) {
		toSerialize["repositoryWaivers"] = o.RepositoryWaivers
	}
	return toSerialize, nil
}

type NullableApiComponentWaiversDTO struct {
	value *ApiComponentWaiversDTO
	isSet bool
}

func (v NullableApiComponentWaiversDTO) Get() *ApiComponentWaiversDTO {
	return v.value
}

func (v *NullableApiComponentWaiversDTO) Set(val *ApiComponentWaiversDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentWaiversDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentWaiversDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentWaiversDTO(val *ApiComponentWaiversDTO) *NullableApiComponentWaiversDTO {
	return &NullableApiComponentWaiversDTO{value: val, isSet: true}
}

func (v NullableApiComponentWaiversDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentWaiversDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


