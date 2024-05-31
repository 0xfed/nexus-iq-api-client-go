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

// checks if the ApiMemberDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMemberDTO{}

// ApiMemberDTO struct for ApiMemberDTO
type ApiMemberDTO struct {
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerType *string `json:"ownerType,omitempty"`
	Type *string `json:"type,omitempty"`
	UserOrGroupName *string `json:"userOrGroupName,omitempty"`
}

// NewApiMemberDTO instantiates a new ApiMemberDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMemberDTO() *ApiMemberDTO {
	this := ApiMemberDTO{}
	return &this
}

// NewApiMemberDTOWithDefaults instantiates a new ApiMemberDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMemberDTOWithDefaults() *ApiMemberDTO {
	this := ApiMemberDTO{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ApiMemberDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMemberDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ApiMemberDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ApiMemberDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *ApiMemberDTO) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMemberDTO) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *ApiMemberDTO) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *ApiMemberDTO) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiMemberDTO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMemberDTO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiMemberDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiMemberDTO) SetType(v string) {
	o.Type = &v
}

// GetUserOrGroupName returns the UserOrGroupName field value if set, zero value otherwise.
func (o *ApiMemberDTO) GetUserOrGroupName() string {
	if o == nil || IsNil(o.UserOrGroupName) {
		var ret string
		return ret
	}
	return *o.UserOrGroupName
}

// GetUserOrGroupNameOk returns a tuple with the UserOrGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMemberDTO) GetUserOrGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserOrGroupName) {
		return nil, false
	}
	return o.UserOrGroupName, true
}

// HasUserOrGroupName returns a boolean if a field has been set.
func (o *ApiMemberDTO) HasUserOrGroupName() bool {
	if o != nil && !IsNil(o.UserOrGroupName) {
		return true
	}

	return false
}

// SetUserOrGroupName gets a reference to the given string and assigns it to the UserOrGroupName field.
func (o *ApiMemberDTO) SetUserOrGroupName(v string) {
	o.UserOrGroupName = &v
}

func (o ApiMemberDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMemberDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.OwnerType) {
		toSerialize["ownerType"] = o.OwnerType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserOrGroupName) {
		toSerialize["userOrGroupName"] = o.UserOrGroupName
	}
	return toSerialize, nil
}

type NullableApiMemberDTO struct {
	value *ApiMemberDTO
	isSet bool
}

func (v NullableApiMemberDTO) Get() *ApiMemberDTO {
	return v.value
}

func (v *NullableApiMemberDTO) Set(val *ApiMemberDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMemberDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMemberDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMemberDTO(val *ApiMemberDTO) *NullableApiMemberDTO {
	return &NullableApiMemberDTO{value: val, isSet: true}
}

func (v NullableApiMemberDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMemberDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


