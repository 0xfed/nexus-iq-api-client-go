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

// checks if the ConfigurationValidationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationValidationResult{}

// ConfigurationValidationResult struct for ConfigurationValidationResult
type ConfigurationValidationResult struct {
	ConfigurationComplete *ValidationResult `json:"configurationComplete,omitempty"`
	RepoPrivate *ValidationResult `json:"repoPrivate,omitempty"`
	SshConfiguration *ValidationResult `json:"sshConfiguration,omitempty"`
	TokenPermissions *ValidationResult `json:"tokenPermissions,omitempty"`
}

// NewConfigurationValidationResult instantiates a new ConfigurationValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationValidationResult() *ConfigurationValidationResult {
	this := ConfigurationValidationResult{}
	return &this
}

// NewConfigurationValidationResultWithDefaults instantiates a new ConfigurationValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationValidationResultWithDefaults() *ConfigurationValidationResult {
	this := ConfigurationValidationResult{}
	return &this
}

// GetConfigurationComplete returns the ConfigurationComplete field value if set, zero value otherwise.
func (o *ConfigurationValidationResult) GetConfigurationComplete() ValidationResult {
	if o == nil || IsNil(o.ConfigurationComplete) {
		var ret ValidationResult
		return ret
	}
	return *o.ConfigurationComplete
}

// GetConfigurationCompleteOk returns a tuple with the ConfigurationComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationValidationResult) GetConfigurationCompleteOk() (*ValidationResult, bool) {
	if o == nil || IsNil(o.ConfigurationComplete) {
		return nil, false
	}
	return o.ConfigurationComplete, true
}

// HasConfigurationComplete returns a boolean if a field has been set.
func (o *ConfigurationValidationResult) HasConfigurationComplete() bool {
	if o != nil && !IsNil(o.ConfigurationComplete) {
		return true
	}

	return false
}

// SetConfigurationComplete gets a reference to the given ValidationResult and assigns it to the ConfigurationComplete field.
func (o *ConfigurationValidationResult) SetConfigurationComplete(v ValidationResult) {
	o.ConfigurationComplete = &v
}

// GetRepoPrivate returns the RepoPrivate field value if set, zero value otherwise.
func (o *ConfigurationValidationResult) GetRepoPrivate() ValidationResult {
	if o == nil || IsNil(o.RepoPrivate) {
		var ret ValidationResult
		return ret
	}
	return *o.RepoPrivate
}

// GetRepoPrivateOk returns a tuple with the RepoPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationValidationResult) GetRepoPrivateOk() (*ValidationResult, bool) {
	if o == nil || IsNil(o.RepoPrivate) {
		return nil, false
	}
	return o.RepoPrivate, true
}

// HasRepoPrivate returns a boolean if a field has been set.
func (o *ConfigurationValidationResult) HasRepoPrivate() bool {
	if o != nil && !IsNil(o.RepoPrivate) {
		return true
	}

	return false
}

// SetRepoPrivate gets a reference to the given ValidationResult and assigns it to the RepoPrivate field.
func (o *ConfigurationValidationResult) SetRepoPrivate(v ValidationResult) {
	o.RepoPrivate = &v
}

// GetSshConfiguration returns the SshConfiguration field value if set, zero value otherwise.
func (o *ConfigurationValidationResult) GetSshConfiguration() ValidationResult {
	if o == nil || IsNil(o.SshConfiguration) {
		var ret ValidationResult
		return ret
	}
	return *o.SshConfiguration
}

// GetSshConfigurationOk returns a tuple with the SshConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationValidationResult) GetSshConfigurationOk() (*ValidationResult, bool) {
	if o == nil || IsNil(o.SshConfiguration) {
		return nil, false
	}
	return o.SshConfiguration, true
}

// HasSshConfiguration returns a boolean if a field has been set.
func (o *ConfigurationValidationResult) HasSshConfiguration() bool {
	if o != nil && !IsNil(o.SshConfiguration) {
		return true
	}

	return false
}

// SetSshConfiguration gets a reference to the given ValidationResult and assigns it to the SshConfiguration field.
func (o *ConfigurationValidationResult) SetSshConfiguration(v ValidationResult) {
	o.SshConfiguration = &v
}

// GetTokenPermissions returns the TokenPermissions field value if set, zero value otherwise.
func (o *ConfigurationValidationResult) GetTokenPermissions() ValidationResult {
	if o == nil || IsNil(o.TokenPermissions) {
		var ret ValidationResult
		return ret
	}
	return *o.TokenPermissions
}

// GetTokenPermissionsOk returns a tuple with the TokenPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationValidationResult) GetTokenPermissionsOk() (*ValidationResult, bool) {
	if o == nil || IsNil(o.TokenPermissions) {
		return nil, false
	}
	return o.TokenPermissions, true
}

// HasTokenPermissions returns a boolean if a field has been set.
func (o *ConfigurationValidationResult) HasTokenPermissions() bool {
	if o != nil && !IsNil(o.TokenPermissions) {
		return true
	}

	return false
}

// SetTokenPermissions gets a reference to the given ValidationResult and assigns it to the TokenPermissions field.
func (o *ConfigurationValidationResult) SetTokenPermissions(v ValidationResult) {
	o.TokenPermissions = &v
}

func (o ConfigurationValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationValidationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigurationComplete) {
		toSerialize["configurationComplete"] = o.ConfigurationComplete
	}
	if !IsNil(o.RepoPrivate) {
		toSerialize["repoPrivate"] = o.RepoPrivate
	}
	if !IsNil(o.SshConfiguration) {
		toSerialize["sshConfiguration"] = o.SshConfiguration
	}
	if !IsNil(o.TokenPermissions) {
		toSerialize["tokenPermissions"] = o.TokenPermissions
	}
	return toSerialize, nil
}

type NullableConfigurationValidationResult struct {
	value *ConfigurationValidationResult
	isSet bool
}

func (v NullableConfigurationValidationResult) Get() *ConfigurationValidationResult {
	return v.value
}

func (v *NullableConfigurationValidationResult) Set(val *ConfigurationValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationValidationResult(val *ConfigurationValidationResult) *NullableConfigurationValidationResult {
	return &NullableConfigurationValidationResult{value: val, isSet: true}
}

func (v NullableConfigurationValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


