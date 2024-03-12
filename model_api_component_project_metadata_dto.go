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

// checks if the ApiComponentProjectMetadataDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentProjectMetadataDTO{}

// ApiComponentProjectMetadataDTO struct for ApiComponentProjectMetadataDTO
type ApiComponentProjectMetadataDTO struct {
	Description *string `json:"description,omitempty"`
	Organization *string `json:"organization,omitempty"`
}

// NewApiComponentProjectMetadataDTO instantiates a new ApiComponentProjectMetadataDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentProjectMetadataDTO() *ApiComponentProjectMetadataDTO {
	this := ApiComponentProjectMetadataDTO{}
	return &this
}

// NewApiComponentProjectMetadataDTOWithDefaults instantiates a new ApiComponentProjectMetadataDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentProjectMetadataDTOWithDefaults() *ApiComponentProjectMetadataDTO {
	this := ApiComponentProjectMetadataDTO{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiComponentProjectMetadataDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectMetadataDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiComponentProjectMetadataDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiComponentProjectMetadataDTO) SetDescription(v string) {
	o.Description = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ApiComponentProjectMetadataDTO) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectMetadataDTO) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ApiComponentProjectMetadataDTO) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *ApiComponentProjectMetadataDTO) SetOrganization(v string) {
	o.Organization = &v
}

func (o ApiComponentProjectMetadataDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentProjectMetadataDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableApiComponentProjectMetadataDTO struct {
	value *ApiComponentProjectMetadataDTO
	isSet bool
}

func (v NullableApiComponentProjectMetadataDTO) Get() *ApiComponentProjectMetadataDTO {
	return v.value
}

func (v *NullableApiComponentProjectMetadataDTO) Set(val *ApiComponentProjectMetadataDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentProjectMetadataDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentProjectMetadataDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentProjectMetadataDTO(val *ApiComponentProjectMetadataDTO) *NullableApiComponentProjectMetadataDTO {
	return &NullableApiComponentProjectMetadataDTO{value: val, isSet: true}
}

func (v NullableApiComponentProjectMetadataDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentProjectMetadataDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


