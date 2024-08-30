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

// checks if the ApiRepositoryComponentsInQuarantineDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryComponentsInQuarantineDTO{}

// ApiRepositoryComponentsInQuarantineDTO struct for ApiRepositoryComponentsInQuarantineDTO
type ApiRepositoryComponentsInQuarantineDTO struct {
	Components []ApiRepositoryComponentPolicyViolationDTO `json:"components,omitempty"`
	Repository *ApiRepositoryDTO `json:"repository,omitempty"`
}

// NewApiRepositoryComponentsInQuarantineDTO instantiates a new ApiRepositoryComponentsInQuarantineDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryComponentsInQuarantineDTO() *ApiRepositoryComponentsInQuarantineDTO {
	this := ApiRepositoryComponentsInQuarantineDTO{}
	return &this
}

// NewApiRepositoryComponentsInQuarantineDTOWithDefaults instantiates a new ApiRepositoryComponentsInQuarantineDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryComponentsInQuarantineDTOWithDefaults() *ApiRepositoryComponentsInQuarantineDTO {
	this := ApiRepositoryComponentsInQuarantineDTO{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ApiRepositoryComponentsInQuarantineDTO) GetComponents() []ApiRepositoryComponentPolicyViolationDTO {
	if o == nil || IsNil(o.Components) {
		var ret []ApiRepositoryComponentPolicyViolationDTO
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentsInQuarantineDTO) GetComponentsOk() ([]ApiRepositoryComponentPolicyViolationDTO, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ApiRepositoryComponentsInQuarantineDTO) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ApiRepositoryComponentPolicyViolationDTO and assigns it to the Components field.
func (o *ApiRepositoryComponentsInQuarantineDTO) SetComponents(v []ApiRepositoryComponentPolicyViolationDTO) {
	o.Components = v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ApiRepositoryComponentsInQuarantineDTO) GetRepository() ApiRepositoryDTO {
	if o == nil || IsNil(o.Repository) {
		var ret ApiRepositoryDTO
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentsInQuarantineDTO) GetRepositoryOk() (*ApiRepositoryDTO, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ApiRepositoryComponentsInQuarantineDTO) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given ApiRepositoryDTO and assigns it to the Repository field.
func (o *ApiRepositoryComponentsInQuarantineDTO) SetRepository(v ApiRepositoryDTO) {
	o.Repository = &v
}

func (o ApiRepositoryComponentsInQuarantineDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryComponentsInQuarantineDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

type NullableApiRepositoryComponentsInQuarantineDTO struct {
	value *ApiRepositoryComponentsInQuarantineDTO
	isSet bool
}

func (v NullableApiRepositoryComponentsInQuarantineDTO) Get() *ApiRepositoryComponentsInQuarantineDTO {
	return v.value
}

func (v *NullableApiRepositoryComponentsInQuarantineDTO) Set(val *ApiRepositoryComponentsInQuarantineDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryComponentsInQuarantineDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryComponentsInQuarantineDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryComponentsInQuarantineDTO(val *ApiRepositoryComponentsInQuarantineDTO) *NullableApiRepositoryComponentsInQuarantineDTO {
	return &NullableApiRepositoryComponentsInQuarantineDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryComponentsInQuarantineDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryComponentsInQuarantineDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


