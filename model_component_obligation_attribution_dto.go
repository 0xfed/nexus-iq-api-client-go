/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 165
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ComponentObligationAttributionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentObligationAttributionDTO{}

// ComponentObligationAttributionDTO struct for ComponentObligationAttributionDTO
type ComponentObligationAttributionDTO struct {
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	Content *string `json:"content,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"`
	LastUpdatedByUsername *string `json:"lastUpdatedByUsername,omitempty"`
	ObligationName *string `json:"obligationName,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
}

// NewComponentObligationAttributionDTO instantiates a new ComponentObligationAttributionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentObligationAttributionDTO() *ComponentObligationAttributionDTO {
	this := ComponentObligationAttributionDTO{}
	return &this
}

// NewComponentObligationAttributionDTOWithDefaults instantiates a new ComponentObligationAttributionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentObligationAttributionDTOWithDefaults() *ComponentObligationAttributionDTO {
	this := ComponentObligationAttributionDTO{}
	return &this
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ComponentObligationAttributionDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentObligationAttributionDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ComponentObligationAttributionDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ComponentObligationAttributionDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ComponentObligationAttributionDTO) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentObligationAttributionDTO) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ComponentObligationAttributionDTO) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ComponentObligationAttributionDTO) SetContent(v string) {
	o.Content = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentObligationAttributionDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentObligationAttributionDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentObligationAttributionDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComponentObligationAttributionDTO) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *ComponentObligationAttributionDTO) GetLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.LastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentObligationAttributionDTO) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedAt) {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *ComponentObligationAttributionDTO) HasLastUpdatedAt() bool {
	if o != nil && !IsNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *ComponentObligationAttributionDTO) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

// GetLastUpdatedByUsername returns the LastUpdatedByUsername field value if set, zero value otherwise.
func (o *ComponentObligationAttributionDTO) GetLastUpdatedByUsername() string {
	if o == nil || IsNil(o.LastUpdatedByUsername) {
		var ret string
		return ret
	}
	return *o.LastUpdatedByUsername
}

// GetLastUpdatedByUsernameOk returns a tuple with the LastUpdatedByUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentObligationAttributionDTO) GetLastUpdatedByUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedByUsername) {
		return nil, false
	}
	return o.LastUpdatedByUsername, true
}

// HasLastUpdatedByUsername returns a boolean if a field has been set.
func (o *ComponentObligationAttributionDTO) HasLastUpdatedByUsername() bool {
	if o != nil && !IsNil(o.LastUpdatedByUsername) {
		return true
	}

	return false
}

// SetLastUpdatedByUsername gets a reference to the given string and assigns it to the LastUpdatedByUsername field.
func (o *ComponentObligationAttributionDTO) SetLastUpdatedByUsername(v string) {
	o.LastUpdatedByUsername = &v
}

// GetObligationName returns the ObligationName field value if set, zero value otherwise.
func (o *ComponentObligationAttributionDTO) GetObligationName() string {
	if o == nil || IsNil(o.ObligationName) {
		var ret string
		return ret
	}
	return *o.ObligationName
}

// GetObligationNameOk returns a tuple with the ObligationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentObligationAttributionDTO) GetObligationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObligationName) {
		return nil, false
	}
	return o.ObligationName, true
}

// HasObligationName returns a boolean if a field has been set.
func (o *ComponentObligationAttributionDTO) HasObligationName() bool {
	if o != nil && !IsNil(o.ObligationName) {
		return true
	}

	return false
}

// SetObligationName gets a reference to the given string and assigns it to the ObligationName field.
func (o *ComponentObligationAttributionDTO) SetObligationName(v string) {
	o.ObligationName = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ComponentObligationAttributionDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentObligationAttributionDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ComponentObligationAttributionDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ComponentObligationAttributionDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ComponentObligationAttributionDTO) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentObligationAttributionDTO) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ComponentObligationAttributionDTO) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ComponentObligationAttributionDTO) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

func (o ComponentObligationAttributionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentObligationAttributionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	if !IsNil(o.LastUpdatedByUsername) {
		toSerialize["lastUpdatedByUsername"] = o.LastUpdatedByUsername
	}
	if !IsNil(o.ObligationName) {
		toSerialize["obligationName"] = o.ObligationName
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	return toSerialize, nil
}

type NullableComponentObligationAttributionDTO struct {
	value *ComponentObligationAttributionDTO
	isSet bool
}

func (v NullableComponentObligationAttributionDTO) Get() *ComponentObligationAttributionDTO {
	return v.value
}

func (v *NullableComponentObligationAttributionDTO) Set(val *ComponentObligationAttributionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentObligationAttributionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentObligationAttributionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentObligationAttributionDTO(val *ComponentObligationAttributionDTO) *NullableComponentObligationAttributionDTO {
	return &NullableComponentObligationAttributionDTO{value: val, isSet: true}
}

func (v NullableComponentObligationAttributionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentObligationAttributionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


