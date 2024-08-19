/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.180.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiUserTokenDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUserTokenDTO{}

// ApiUserTokenDTO struct for ApiUserTokenDTO
type ApiUserTokenDTO struct {
	PassCode *string `json:"passCode,omitempty"`
	Realm *string `json:"realm,omitempty"`
	UserCode *string `json:"userCode,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewApiUserTokenDTO instantiates a new ApiUserTokenDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserTokenDTO() *ApiUserTokenDTO {
	this := ApiUserTokenDTO{}
	return &this
}

// NewApiUserTokenDTOWithDefaults instantiates a new ApiUserTokenDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserTokenDTOWithDefaults() *ApiUserTokenDTO {
	this := ApiUserTokenDTO{}
	return &this
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *ApiUserTokenDTO) GetPassCode() string {
	if o == nil || IsNil(o.PassCode) {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserTokenDTO) GetPassCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PassCode) {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *ApiUserTokenDTO) HasPassCode() bool {
	if o != nil && !IsNil(o.PassCode) {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *ApiUserTokenDTO) SetPassCode(v string) {
	o.PassCode = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *ApiUserTokenDTO) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserTokenDTO) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *ApiUserTokenDTO) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *ApiUserTokenDTO) SetRealm(v string) {
	o.Realm = &v
}

// GetUserCode returns the UserCode field value if set, zero value otherwise.
func (o *ApiUserTokenDTO) GetUserCode() string {
	if o == nil || IsNil(o.UserCode) {
		var ret string
		return ret
	}
	return *o.UserCode
}

// GetUserCodeOk returns a tuple with the UserCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserTokenDTO) GetUserCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UserCode) {
		return nil, false
	}
	return o.UserCode, true
}

// HasUserCode returns a boolean if a field has been set.
func (o *ApiUserTokenDTO) HasUserCode() bool {
	if o != nil && !IsNil(o.UserCode) {
		return true
	}

	return false
}

// SetUserCode gets a reference to the given string and assigns it to the UserCode field.
func (o *ApiUserTokenDTO) SetUserCode(v string) {
	o.UserCode = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiUserTokenDTO) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserTokenDTO) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiUserTokenDTO) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiUserTokenDTO) SetUsername(v string) {
	o.Username = &v
}

func (o ApiUserTokenDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserTokenDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PassCode) {
		toSerialize["passCode"] = o.PassCode
	}
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !IsNil(o.UserCode) {
		toSerialize["userCode"] = o.UserCode
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableApiUserTokenDTO struct {
	value *ApiUserTokenDTO
	isSet bool
}

func (v NullableApiUserTokenDTO) Get() *ApiUserTokenDTO {
	return v.value
}

func (v *NullableApiUserTokenDTO) Set(val *ApiUserTokenDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserTokenDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserTokenDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserTokenDTO(val *ApiUserTokenDTO) *NullableApiUserTokenDTO {
	return &NullableApiUserTokenDTO{value: val, isSet: true}
}

func (v NullableApiUserTokenDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserTokenDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


