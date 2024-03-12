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

// checks if the ApiRoleMemberMappingDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRoleMemberMappingDTO{}

// ApiRoleMemberMappingDTO struct for ApiRoleMemberMappingDTO
type ApiRoleMemberMappingDTO struct {
	Members []ApiMemberDTO `json:"members,omitempty"`
	RoleId *string `json:"roleId,omitempty"`
}

// NewApiRoleMemberMappingDTO instantiates a new ApiRoleMemberMappingDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRoleMemberMappingDTO() *ApiRoleMemberMappingDTO {
	this := ApiRoleMemberMappingDTO{}
	return &this
}

// NewApiRoleMemberMappingDTOWithDefaults instantiates a new ApiRoleMemberMappingDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRoleMemberMappingDTOWithDefaults() *ApiRoleMemberMappingDTO {
	this := ApiRoleMemberMappingDTO{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ApiRoleMemberMappingDTO) GetMembers() []ApiMemberDTO {
	if o == nil || IsNil(o.Members) {
		var ret []ApiMemberDTO
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleMemberMappingDTO) GetMembersOk() ([]ApiMemberDTO, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ApiRoleMemberMappingDTO) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []ApiMemberDTO and assigns it to the Members field.
func (o *ApiRoleMemberMappingDTO) SetMembers(v []ApiMemberDTO) {
	o.Members = v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *ApiRoleMemberMappingDTO) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleMemberMappingDTO) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ApiRoleMemberMappingDTO) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *ApiRoleMemberMappingDTO) SetRoleId(v string) {
	o.RoleId = &v
}

func (o ApiRoleMemberMappingDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRoleMemberMappingDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	return toSerialize, nil
}

type NullableApiRoleMemberMappingDTO struct {
	value *ApiRoleMemberMappingDTO
	isSet bool
}

func (v NullableApiRoleMemberMappingDTO) Get() *ApiRoleMemberMappingDTO {
	return v.value
}

func (v *NullableApiRoleMemberMappingDTO) Set(val *ApiRoleMemberMappingDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRoleMemberMappingDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRoleMemberMappingDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRoleMemberMappingDTO(val *ApiRoleMemberMappingDTO) *NullableApiRoleMemberMappingDTO {
	return &NullableApiRoleMemberMappingDTO{value: val, isSet: true}
}

func (v NullableApiRoleMemberMappingDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRoleMemberMappingDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


