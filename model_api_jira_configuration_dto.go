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

// checks if the ApiJiraConfigurationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiJiraConfigurationDTO{}

// ApiJiraConfigurationDTO struct for ApiJiraConfigurationDTO
type ApiJiraConfigurationDTO struct {
	CustomFields map[string]map[string]interface{} `json:"customFields,omitempty"`
	Password []string `json:"password,omitempty"`
	Url *string `json:"url,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewApiJiraConfigurationDTO instantiates a new ApiJiraConfigurationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiJiraConfigurationDTO() *ApiJiraConfigurationDTO {
	this := ApiJiraConfigurationDTO{}
	return &this
}

// NewApiJiraConfigurationDTOWithDefaults instantiates a new ApiJiraConfigurationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiJiraConfigurationDTOWithDefaults() *ApiJiraConfigurationDTO {
	this := ApiJiraConfigurationDTO{}
	return &this
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ApiJiraConfigurationDTO) GetCustomFields() map[string]map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiJiraConfigurationDTO) GetCustomFieldsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ApiJiraConfigurationDTO) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]map[string]interface{} and assigns it to the CustomFields field.
func (o *ApiJiraConfigurationDTO) SetCustomFields(v map[string]map[string]interface{}) {
	o.CustomFields = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiJiraConfigurationDTO) GetPassword() []string {
	if o == nil || IsNil(o.Password) {
		var ret []string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiJiraConfigurationDTO) GetPasswordOk() ([]string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiJiraConfigurationDTO) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given []string and assigns it to the Password field.
func (o *ApiJiraConfigurationDTO) SetPassword(v []string) {
	o.Password = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ApiJiraConfigurationDTO) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiJiraConfigurationDTO) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApiJiraConfigurationDTO) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ApiJiraConfigurationDTO) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiJiraConfigurationDTO) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiJiraConfigurationDTO) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiJiraConfigurationDTO) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiJiraConfigurationDTO) SetUsername(v string) {
	o.Username = &v
}

func (o ApiJiraConfigurationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiJiraConfigurationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableApiJiraConfigurationDTO struct {
	value *ApiJiraConfigurationDTO
	isSet bool
}

func (v NullableApiJiraConfigurationDTO) Get() *ApiJiraConfigurationDTO {
	return v.value
}

func (v *NullableApiJiraConfigurationDTO) Set(val *ApiJiraConfigurationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiJiraConfigurationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiJiraConfigurationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiJiraConfigurationDTO(val *ApiJiraConfigurationDTO) *NullableApiJiraConfigurationDTO {
	return &NullableApiJiraConfigurationDTO{value: val, isSet: true}
}

func (v NullableApiJiraConfigurationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiJiraConfigurationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


