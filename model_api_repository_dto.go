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

// checks if the ApiRepositoryDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryDTO{}

// ApiRepositoryDTO struct for ApiRepositoryDTO
type ApiRepositoryDTO struct {
	AuditEnabled *bool `json:"auditEnabled,omitempty"`
	Format *string `json:"format,omitempty"`
	NamespaceConfusionProtectionEnabled *bool `json:"namespaceConfusionProtectionEnabled,omitempty"`
	PolicyCompliantComponentSelectionEnabled *bool `json:"policyCompliantComponentSelectionEnabled,omitempty"`
	PublicId *string `json:"publicId,omitempty"`
	QuarantineEnabled *bool `json:"quarantineEnabled,omitempty"`
	RepositoryId *string `json:"repositoryId,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewApiRepositoryDTO instantiates a new ApiRepositoryDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryDTO() *ApiRepositoryDTO {
	this := ApiRepositoryDTO{}
	return &this
}

// NewApiRepositoryDTOWithDefaults instantiates a new ApiRepositoryDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryDTOWithDefaults() *ApiRepositoryDTO {
	this := ApiRepositoryDTO{}
	return &this
}

// GetAuditEnabled returns the AuditEnabled field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetAuditEnabled() bool {
	if o == nil || IsNil(o.AuditEnabled) {
		var ret bool
		return ret
	}
	return *o.AuditEnabled
}

// GetAuditEnabledOk returns a tuple with the AuditEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetAuditEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AuditEnabled) {
		return nil, false
	}
	return o.AuditEnabled, true
}

// HasAuditEnabled returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasAuditEnabled() bool {
	if o != nil && !IsNil(o.AuditEnabled) {
		return true
	}

	return false
}

// SetAuditEnabled gets a reference to the given bool and assigns it to the AuditEnabled field.
func (o *ApiRepositoryDTO) SetAuditEnabled(v bool) {
	o.AuditEnabled = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ApiRepositoryDTO) SetFormat(v string) {
	o.Format = &v
}

// GetNamespaceConfusionProtectionEnabled returns the NamespaceConfusionProtectionEnabled field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetNamespaceConfusionProtectionEnabled() bool {
	if o == nil || IsNil(o.NamespaceConfusionProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.NamespaceConfusionProtectionEnabled
}

// GetNamespaceConfusionProtectionEnabledOk returns a tuple with the NamespaceConfusionProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetNamespaceConfusionProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NamespaceConfusionProtectionEnabled) {
		return nil, false
	}
	return o.NamespaceConfusionProtectionEnabled, true
}

// HasNamespaceConfusionProtectionEnabled returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasNamespaceConfusionProtectionEnabled() bool {
	if o != nil && !IsNil(o.NamespaceConfusionProtectionEnabled) {
		return true
	}

	return false
}

// SetNamespaceConfusionProtectionEnabled gets a reference to the given bool and assigns it to the NamespaceConfusionProtectionEnabled field.
func (o *ApiRepositoryDTO) SetNamespaceConfusionProtectionEnabled(v bool) {
	o.NamespaceConfusionProtectionEnabled = &v
}

// GetPolicyCompliantComponentSelectionEnabled returns the PolicyCompliantComponentSelectionEnabled field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetPolicyCompliantComponentSelectionEnabled() bool {
	if o == nil || IsNil(o.PolicyCompliantComponentSelectionEnabled) {
		var ret bool
		return ret
	}
	return *o.PolicyCompliantComponentSelectionEnabled
}

// GetPolicyCompliantComponentSelectionEnabledOk returns a tuple with the PolicyCompliantComponentSelectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetPolicyCompliantComponentSelectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PolicyCompliantComponentSelectionEnabled) {
		return nil, false
	}
	return o.PolicyCompliantComponentSelectionEnabled, true
}

// HasPolicyCompliantComponentSelectionEnabled returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasPolicyCompliantComponentSelectionEnabled() bool {
	if o != nil && !IsNil(o.PolicyCompliantComponentSelectionEnabled) {
		return true
	}

	return false
}

// SetPolicyCompliantComponentSelectionEnabled gets a reference to the given bool and assigns it to the PolicyCompliantComponentSelectionEnabled field.
func (o *ApiRepositoryDTO) SetPolicyCompliantComponentSelectionEnabled(v bool) {
	o.PolicyCompliantComponentSelectionEnabled = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetPublicId() string {
	if o == nil || IsNil(o.PublicId) {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetPublicIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicId) {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasPublicId() bool {
	if o != nil && !IsNil(o.PublicId) {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *ApiRepositoryDTO) SetPublicId(v string) {
	o.PublicId = &v
}

// GetQuarantineEnabled returns the QuarantineEnabled field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetQuarantineEnabled() bool {
	if o == nil || IsNil(o.QuarantineEnabled) {
		var ret bool
		return ret
	}
	return *o.QuarantineEnabled
}

// GetQuarantineEnabledOk returns a tuple with the QuarantineEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetQuarantineEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.QuarantineEnabled) {
		return nil, false
	}
	return o.QuarantineEnabled, true
}

// HasQuarantineEnabled returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasQuarantineEnabled() bool {
	if o != nil && !IsNil(o.QuarantineEnabled) {
		return true
	}

	return false
}

// SetQuarantineEnabled gets a reference to the given bool and assigns it to the QuarantineEnabled field.
func (o *ApiRepositoryDTO) SetQuarantineEnabled(v bool) {
	o.QuarantineEnabled = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetRepositoryId() string {
	if o == nil || IsNil(o.RepositoryId) {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetRepositoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryId) {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasRepositoryId() bool {
	if o != nil && !IsNil(o.RepositoryId) {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *ApiRepositoryDTO) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiRepositoryDTO) SetType(v string) {
	o.Type = &v
}

func (o ApiRepositoryDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditEnabled) {
		toSerialize["auditEnabled"] = o.AuditEnabled
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.NamespaceConfusionProtectionEnabled) {
		toSerialize["namespaceConfusionProtectionEnabled"] = o.NamespaceConfusionProtectionEnabled
	}
	if !IsNil(o.PolicyCompliantComponentSelectionEnabled) {
		toSerialize["policyCompliantComponentSelectionEnabled"] = o.PolicyCompliantComponentSelectionEnabled
	}
	if !IsNil(o.PublicId) {
		toSerialize["publicId"] = o.PublicId
	}
	if !IsNil(o.QuarantineEnabled) {
		toSerialize["quarantineEnabled"] = o.QuarantineEnabled
	}
	if !IsNil(o.RepositoryId) {
		toSerialize["repositoryId"] = o.RepositoryId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableApiRepositoryDTO struct {
	value *ApiRepositoryDTO
	isSet bool
}

func (v NullableApiRepositoryDTO) Get() *ApiRepositoryDTO {
	return v.value
}

func (v *NullableApiRepositoryDTO) Set(val *ApiRepositoryDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryDTO(val *ApiRepositoryDTO) *NullableApiRepositoryDTO {
	return &NullableApiRepositoryDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


