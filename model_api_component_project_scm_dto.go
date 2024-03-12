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

// checks if the ApiComponentProjectScmDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentProjectScmDTO{}

// ApiComponentProjectScmDTO struct for ApiComponentProjectScmDTO
type ApiComponentProjectScmDTO struct {
	ScmDetails *ApiComponentProjectScmDetailsDTO `json:"scmDetails,omitempty"`
	ScmMetadata *ApiComponentProjectScmMetadataDTO `json:"scmMetadata,omitempty"`
	ScmUrl *string `json:"scmUrl,omitempty"`
}

// NewApiComponentProjectScmDTO instantiates a new ApiComponentProjectScmDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentProjectScmDTO() *ApiComponentProjectScmDTO {
	this := ApiComponentProjectScmDTO{}
	return &this
}

// NewApiComponentProjectScmDTOWithDefaults instantiates a new ApiComponentProjectScmDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentProjectScmDTOWithDefaults() *ApiComponentProjectScmDTO {
	this := ApiComponentProjectScmDTO{}
	return &this
}

// GetScmDetails returns the ScmDetails field value if set, zero value otherwise.
func (o *ApiComponentProjectScmDTO) GetScmDetails() ApiComponentProjectScmDetailsDTO {
	if o == nil || IsNil(o.ScmDetails) {
		var ret ApiComponentProjectScmDetailsDTO
		return ret
	}
	return *o.ScmDetails
}

// GetScmDetailsOk returns a tuple with the ScmDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectScmDTO) GetScmDetailsOk() (*ApiComponentProjectScmDetailsDTO, bool) {
	if o == nil || IsNil(o.ScmDetails) {
		return nil, false
	}
	return o.ScmDetails, true
}

// HasScmDetails returns a boolean if a field has been set.
func (o *ApiComponentProjectScmDTO) HasScmDetails() bool {
	if o != nil && !IsNil(o.ScmDetails) {
		return true
	}

	return false
}

// SetScmDetails gets a reference to the given ApiComponentProjectScmDetailsDTO and assigns it to the ScmDetails field.
func (o *ApiComponentProjectScmDTO) SetScmDetails(v ApiComponentProjectScmDetailsDTO) {
	o.ScmDetails = &v
}

// GetScmMetadata returns the ScmMetadata field value if set, zero value otherwise.
func (o *ApiComponentProjectScmDTO) GetScmMetadata() ApiComponentProjectScmMetadataDTO {
	if o == nil || IsNil(o.ScmMetadata) {
		var ret ApiComponentProjectScmMetadataDTO
		return ret
	}
	return *o.ScmMetadata
}

// GetScmMetadataOk returns a tuple with the ScmMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectScmDTO) GetScmMetadataOk() (*ApiComponentProjectScmMetadataDTO, bool) {
	if o == nil || IsNil(o.ScmMetadata) {
		return nil, false
	}
	return o.ScmMetadata, true
}

// HasScmMetadata returns a boolean if a field has been set.
func (o *ApiComponentProjectScmDTO) HasScmMetadata() bool {
	if o != nil && !IsNil(o.ScmMetadata) {
		return true
	}

	return false
}

// SetScmMetadata gets a reference to the given ApiComponentProjectScmMetadataDTO and assigns it to the ScmMetadata field.
func (o *ApiComponentProjectScmDTO) SetScmMetadata(v ApiComponentProjectScmMetadataDTO) {
	o.ScmMetadata = &v
}

// GetScmUrl returns the ScmUrl field value if set, zero value otherwise.
func (o *ApiComponentProjectScmDTO) GetScmUrl() string {
	if o == nil || IsNil(o.ScmUrl) {
		var ret string
		return ret
	}
	return *o.ScmUrl
}

// GetScmUrlOk returns a tuple with the ScmUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectScmDTO) GetScmUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ScmUrl) {
		return nil, false
	}
	return o.ScmUrl, true
}

// HasScmUrl returns a boolean if a field has been set.
func (o *ApiComponentProjectScmDTO) HasScmUrl() bool {
	if o != nil && !IsNil(o.ScmUrl) {
		return true
	}

	return false
}

// SetScmUrl gets a reference to the given string and assigns it to the ScmUrl field.
func (o *ApiComponentProjectScmDTO) SetScmUrl(v string) {
	o.ScmUrl = &v
}

func (o ApiComponentProjectScmDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentProjectScmDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScmDetails) {
		toSerialize["scmDetails"] = o.ScmDetails
	}
	if !IsNil(o.ScmMetadata) {
		toSerialize["scmMetadata"] = o.ScmMetadata
	}
	if !IsNil(o.ScmUrl) {
		toSerialize["scmUrl"] = o.ScmUrl
	}
	return toSerialize, nil
}

type NullableApiComponentProjectScmDTO struct {
	value *ApiComponentProjectScmDTO
	isSet bool
}

func (v NullableApiComponentProjectScmDTO) Get() *ApiComponentProjectScmDTO {
	return v.value
}

func (v *NullableApiComponentProjectScmDTO) Set(val *ApiComponentProjectScmDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentProjectScmDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentProjectScmDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentProjectScmDTO(val *ApiComponentProjectScmDTO) *NullableApiComponentProjectScmDTO {
	return &NullableApiComponentProjectScmDTO{value: val, isSet: true}
}

func (v NullableApiComponentProjectScmDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentProjectScmDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


