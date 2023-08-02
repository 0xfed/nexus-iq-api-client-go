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

// checks if the ApiApplicationBaseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiApplicationBaseDTO{}

// ApiApplicationBaseDTO struct for ApiApplicationBaseDTO
type ApiApplicationBaseDTO struct {
	ContactUserName *string `json:"contactUserName,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	PublicId *string `json:"publicId,omitempty"`
}

// NewApiApplicationBaseDTO instantiates a new ApiApplicationBaseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiApplicationBaseDTO() *ApiApplicationBaseDTO {
	this := ApiApplicationBaseDTO{}
	return &this
}

// NewApiApplicationBaseDTOWithDefaults instantiates a new ApiApplicationBaseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiApplicationBaseDTOWithDefaults() *ApiApplicationBaseDTO {
	this := ApiApplicationBaseDTO{}
	return &this
}

// GetContactUserName returns the ContactUserName field value if set, zero value otherwise.
func (o *ApiApplicationBaseDTO) GetContactUserName() string {
	if o == nil || IsNil(o.ContactUserName) {
		var ret string
		return ret
	}
	return *o.ContactUserName
}

// GetContactUserNameOk returns a tuple with the ContactUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationBaseDTO) GetContactUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactUserName) {
		return nil, false
	}
	return o.ContactUserName, true
}

// HasContactUserName returns a boolean if a field has been set.
func (o *ApiApplicationBaseDTO) HasContactUserName() bool {
	if o != nil && !IsNil(o.ContactUserName) {
		return true
	}

	return false
}

// SetContactUserName gets a reference to the given string and assigns it to the ContactUserName field.
func (o *ApiApplicationBaseDTO) SetContactUserName(v string) {
	o.ContactUserName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiApplicationBaseDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationBaseDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiApplicationBaseDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiApplicationBaseDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiApplicationBaseDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationBaseDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiApplicationBaseDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiApplicationBaseDTO) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ApiApplicationBaseDTO) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationBaseDTO) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ApiApplicationBaseDTO) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ApiApplicationBaseDTO) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *ApiApplicationBaseDTO) GetPublicId() string {
	if o == nil || IsNil(o.PublicId) {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationBaseDTO) GetPublicIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicId) {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *ApiApplicationBaseDTO) HasPublicId() bool {
	if o != nil && !IsNil(o.PublicId) {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *ApiApplicationBaseDTO) SetPublicId(v string) {
	o.PublicId = &v
}

func (o ApiApplicationBaseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiApplicationBaseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactUserName) {
		toSerialize["contactUserName"] = o.ContactUserName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.PublicId) {
		toSerialize["publicId"] = o.PublicId
	}
	return toSerialize, nil
}

type NullableApiApplicationBaseDTO struct {
	value *ApiApplicationBaseDTO
	isSet bool
}

func (v NullableApiApplicationBaseDTO) Get() *ApiApplicationBaseDTO {
	return v.value
}

func (v *NullableApiApplicationBaseDTO) Set(val *ApiApplicationBaseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiApplicationBaseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiApplicationBaseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiApplicationBaseDTO(val *ApiApplicationBaseDTO) *NullableApiApplicationBaseDTO {
	return &NullableApiApplicationBaseDTO{value: val, isSet: true}
}

func (v NullableApiApplicationBaseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiApplicationBaseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


