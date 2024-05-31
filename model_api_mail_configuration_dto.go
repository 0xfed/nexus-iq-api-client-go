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

// checks if the ApiMailConfigurationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMailConfigurationDTO{}

// ApiMailConfigurationDTO struct for ApiMailConfigurationDTO
type ApiMailConfigurationDTO struct {
	Hostname *string `json:"hostname,omitempty"`
	Password *string `json:"password,omitempty"`
	PasswordIsIncluded *bool `json:"passwordIsIncluded,omitempty"`
	Port *int32 `json:"port,omitempty"`
	SslEnabled *bool `json:"sslEnabled,omitempty"`
	StartTlsEnabled *bool `json:"startTlsEnabled,omitempty"`
	SystemEmail *string `json:"systemEmail,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewApiMailConfigurationDTO instantiates a new ApiMailConfigurationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMailConfigurationDTO() *ApiMailConfigurationDTO {
	this := ApiMailConfigurationDTO{}
	return &this
}

// NewApiMailConfigurationDTOWithDefaults instantiates a new ApiMailConfigurationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMailConfigurationDTOWithDefaults() *ApiMailConfigurationDTO {
	this := ApiMailConfigurationDTO{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApiMailConfigurationDTO) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMailConfigurationDTO) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApiMailConfigurationDTO) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApiMailConfigurationDTO) SetHostname(v string) {
	o.Hostname = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiMailConfigurationDTO) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMailConfigurationDTO) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiMailConfigurationDTO) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiMailConfigurationDTO) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordIsIncluded returns the PasswordIsIncluded field value if set, zero value otherwise.
func (o *ApiMailConfigurationDTO) GetPasswordIsIncluded() bool {
	if o == nil || IsNil(o.PasswordIsIncluded) {
		var ret bool
		return ret
	}
	return *o.PasswordIsIncluded
}

// GetPasswordIsIncludedOk returns a tuple with the PasswordIsIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMailConfigurationDTO) GetPasswordIsIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordIsIncluded) {
		return nil, false
	}
	return o.PasswordIsIncluded, true
}

// HasPasswordIsIncluded returns a boolean if a field has been set.
func (o *ApiMailConfigurationDTO) HasPasswordIsIncluded() bool {
	if o != nil && !IsNil(o.PasswordIsIncluded) {
		return true
	}

	return false
}

// SetPasswordIsIncluded gets a reference to the given bool and assigns it to the PasswordIsIncluded field.
func (o *ApiMailConfigurationDTO) SetPasswordIsIncluded(v bool) {
	o.PasswordIsIncluded = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApiMailConfigurationDTO) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMailConfigurationDTO) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApiMailConfigurationDTO) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ApiMailConfigurationDTO) SetPort(v int32) {
	o.Port = &v
}

// GetSslEnabled returns the SslEnabled field value if set, zero value otherwise.
func (o *ApiMailConfigurationDTO) GetSslEnabled() bool {
	if o == nil || IsNil(o.SslEnabled) {
		var ret bool
		return ret
	}
	return *o.SslEnabled
}

// GetSslEnabledOk returns a tuple with the SslEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMailConfigurationDTO) GetSslEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SslEnabled) {
		return nil, false
	}
	return o.SslEnabled, true
}

// HasSslEnabled returns a boolean if a field has been set.
func (o *ApiMailConfigurationDTO) HasSslEnabled() bool {
	if o != nil && !IsNil(o.SslEnabled) {
		return true
	}

	return false
}

// SetSslEnabled gets a reference to the given bool and assigns it to the SslEnabled field.
func (o *ApiMailConfigurationDTO) SetSslEnabled(v bool) {
	o.SslEnabled = &v
}

// GetStartTlsEnabled returns the StartTlsEnabled field value if set, zero value otherwise.
func (o *ApiMailConfigurationDTO) GetStartTlsEnabled() bool {
	if o == nil || IsNil(o.StartTlsEnabled) {
		var ret bool
		return ret
	}
	return *o.StartTlsEnabled
}

// GetStartTlsEnabledOk returns a tuple with the StartTlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMailConfigurationDTO) GetStartTlsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StartTlsEnabled) {
		return nil, false
	}
	return o.StartTlsEnabled, true
}

// HasStartTlsEnabled returns a boolean if a field has been set.
func (o *ApiMailConfigurationDTO) HasStartTlsEnabled() bool {
	if o != nil && !IsNil(o.StartTlsEnabled) {
		return true
	}

	return false
}

// SetStartTlsEnabled gets a reference to the given bool and assigns it to the StartTlsEnabled field.
func (o *ApiMailConfigurationDTO) SetStartTlsEnabled(v bool) {
	o.StartTlsEnabled = &v
}

// GetSystemEmail returns the SystemEmail field value if set, zero value otherwise.
func (o *ApiMailConfigurationDTO) GetSystemEmail() string {
	if o == nil || IsNil(o.SystemEmail) {
		var ret string
		return ret
	}
	return *o.SystemEmail
}

// GetSystemEmailOk returns a tuple with the SystemEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMailConfigurationDTO) GetSystemEmailOk() (*string, bool) {
	if o == nil || IsNil(o.SystemEmail) {
		return nil, false
	}
	return o.SystemEmail, true
}

// HasSystemEmail returns a boolean if a field has been set.
func (o *ApiMailConfigurationDTO) HasSystemEmail() bool {
	if o != nil && !IsNil(o.SystemEmail) {
		return true
	}

	return false
}

// SetSystemEmail gets a reference to the given string and assigns it to the SystemEmail field.
func (o *ApiMailConfigurationDTO) SetSystemEmail(v string) {
	o.SystemEmail = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiMailConfigurationDTO) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMailConfigurationDTO) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiMailConfigurationDTO) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiMailConfigurationDTO) SetUsername(v string) {
	o.Username = &v
}

func (o ApiMailConfigurationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMailConfigurationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PasswordIsIncluded) {
		toSerialize["passwordIsIncluded"] = o.PasswordIsIncluded
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.SslEnabled) {
		toSerialize["sslEnabled"] = o.SslEnabled
	}
	if !IsNil(o.StartTlsEnabled) {
		toSerialize["startTlsEnabled"] = o.StartTlsEnabled
	}
	if !IsNil(o.SystemEmail) {
		toSerialize["systemEmail"] = o.SystemEmail
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableApiMailConfigurationDTO struct {
	value *ApiMailConfigurationDTO
	isSet bool
}

func (v NullableApiMailConfigurationDTO) Get() *ApiMailConfigurationDTO {
	return v.value
}

func (v *NullableApiMailConfigurationDTO) Set(val *ApiMailConfigurationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMailConfigurationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMailConfigurationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMailConfigurationDTO(val *ApiMailConfigurationDTO) *NullableApiMailConfigurationDTO {
	return &NullableApiMailConfigurationDTO{value: val, isSet: true}
}

func (v NullableApiMailConfigurationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMailConfigurationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


