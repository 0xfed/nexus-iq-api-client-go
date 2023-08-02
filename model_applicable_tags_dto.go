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

// checks if the ApplicableTagsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicableTagsDTO{}

// ApplicableTagsDTO struct for ApplicableTagsDTO
type ApplicableTagsDTO struct {
	ApplicationCategoriesByOwner []TagsByOwnerDTO `json:"applicationCategoriesByOwner,omitempty"`
}

// NewApplicableTagsDTO instantiates a new ApplicableTagsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicableTagsDTO() *ApplicableTagsDTO {
	this := ApplicableTagsDTO{}
	return &this
}

// NewApplicableTagsDTOWithDefaults instantiates a new ApplicableTagsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicableTagsDTOWithDefaults() *ApplicableTagsDTO {
	this := ApplicableTagsDTO{}
	return &this
}

// GetApplicationCategoriesByOwner returns the ApplicationCategoriesByOwner field value if set, zero value otherwise.
func (o *ApplicableTagsDTO) GetApplicationCategoriesByOwner() []TagsByOwnerDTO {
	if o == nil || IsNil(o.ApplicationCategoriesByOwner) {
		var ret []TagsByOwnerDTO
		return ret
	}
	return o.ApplicationCategoriesByOwner
}

// GetApplicationCategoriesByOwnerOk returns a tuple with the ApplicationCategoriesByOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicableTagsDTO) GetApplicationCategoriesByOwnerOk() ([]TagsByOwnerDTO, bool) {
	if o == nil || IsNil(o.ApplicationCategoriesByOwner) {
		return nil, false
	}
	return o.ApplicationCategoriesByOwner, true
}

// HasApplicationCategoriesByOwner returns a boolean if a field has been set.
func (o *ApplicableTagsDTO) HasApplicationCategoriesByOwner() bool {
	if o != nil && !IsNil(o.ApplicationCategoriesByOwner) {
		return true
	}

	return false
}

// SetApplicationCategoriesByOwner gets a reference to the given []TagsByOwnerDTO and assigns it to the ApplicationCategoriesByOwner field.
func (o *ApplicableTagsDTO) SetApplicationCategoriesByOwner(v []TagsByOwnerDTO) {
	o.ApplicationCategoriesByOwner = v
}

func (o ApplicableTagsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicableTagsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationCategoriesByOwner) {
		toSerialize["applicationCategoriesByOwner"] = o.ApplicationCategoriesByOwner
	}
	return toSerialize, nil
}

type NullableApplicableTagsDTO struct {
	value *ApplicableTagsDTO
	isSet bool
}

func (v NullableApplicableTagsDTO) Get() *ApplicableTagsDTO {
	return v.value
}

func (v *NullableApplicableTagsDTO) Set(val *ApplicableTagsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicableTagsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicableTagsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicableTagsDTO(val *ApplicableTagsDTO) *NullableApplicableTagsDTO {
	return &NullableApplicableTagsDTO{value: val, isSet: true}
}

func (v NullableApplicableTagsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicableTagsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


