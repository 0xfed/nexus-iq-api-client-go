/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.175.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiOrganizationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiOrganizationDTO{}

// ApiOrganizationDTO struct for ApiOrganizationDTO
type ApiOrganizationDTO struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ParentOrganizationId *string `json:"parentOrganizationId,omitempty"`
	Tags []ApiTagDTO `json:"tags,omitempty"`
}

// NewApiOrganizationDTO instantiates a new ApiOrganizationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiOrganizationDTO() *ApiOrganizationDTO {
	this := ApiOrganizationDTO{}
	return &this
}

// NewApiOrganizationDTOWithDefaults instantiates a new ApiOrganizationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiOrganizationDTOWithDefaults() *ApiOrganizationDTO {
	this := ApiOrganizationDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiOrganizationDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiOrganizationDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiOrganizationDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiOrganizationDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiOrganizationDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiOrganizationDTO) SetName(v string) {
	o.Name = &v
}

// GetParentOrganizationId returns the ParentOrganizationId field value if set, zero value otherwise.
func (o *ApiOrganizationDTO) GetParentOrganizationId() string {
	if o == nil || IsNil(o.ParentOrganizationId) {
		var ret string
		return ret
	}
	return *o.ParentOrganizationId
}

// GetParentOrganizationIdOk returns a tuple with the ParentOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationDTO) GetParentOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentOrganizationId) {
		return nil, false
	}
	return o.ParentOrganizationId, true
}

// HasParentOrganizationId returns a boolean if a field has been set.
func (o *ApiOrganizationDTO) HasParentOrganizationId() bool {
	if o != nil && !IsNil(o.ParentOrganizationId) {
		return true
	}

	return false
}

// SetParentOrganizationId gets a reference to the given string and assigns it to the ParentOrganizationId field.
func (o *ApiOrganizationDTO) SetParentOrganizationId(v string) {
	o.ParentOrganizationId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApiOrganizationDTO) GetTags() []ApiTagDTO {
	if o == nil || IsNil(o.Tags) {
		var ret []ApiTagDTO
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationDTO) GetTagsOk() ([]ApiTagDTO, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApiOrganizationDTO) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ApiTagDTO and assigns it to the Tags field.
func (o *ApiOrganizationDTO) SetTags(v []ApiTagDTO) {
	o.Tags = v
}

func (o ApiOrganizationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiOrganizationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ParentOrganizationId) {
		toSerialize["parentOrganizationId"] = o.ParentOrganizationId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableApiOrganizationDTO struct {
	value *ApiOrganizationDTO
	isSet bool
}

func (v NullableApiOrganizationDTO) Get() *ApiOrganizationDTO {
	return v.value
}

func (v *NullableApiOrganizationDTO) Set(val *ApiOrganizationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiOrganizationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiOrganizationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiOrganizationDTO(val *ApiOrganizationDTO) *NullableApiOrganizationDTO {
	return &NullableApiOrganizationDTO{value: val, isSet: true}
}

func (v NullableApiOrganizationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiOrganizationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


