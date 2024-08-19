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

// checks if the ApiLicenseLegalComponentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseLegalComponentDTO{}

// ApiLicenseLegalComponentDTO struct for ApiLicenseLegalComponentDTO
type ApiLicenseLegalComponentDTO struct {
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Hash *string `json:"hash,omitempty"`
	LicenseLegalData *ApiLicenseLegalDataDTO `json:"licenseLegalData,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	StageScans []ApiLicenseLegalStageScanDTO `json:"stageScans,omitempty"`
}

// NewApiLicenseLegalComponentDTO instantiates a new ApiLicenseLegalComponentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseLegalComponentDTO() *ApiLicenseLegalComponentDTO {
	this := ApiLicenseLegalComponentDTO{}
	return &this
}

// NewApiLicenseLegalComponentDTOWithDefaults instantiates a new ApiLicenseLegalComponentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseLegalComponentDTOWithDefaults() *ApiLicenseLegalComponentDTO {
	this := ApiLicenseLegalComponentDTO{}
	return &this
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiLicenseLegalComponentDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalComponentDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiLicenseLegalComponentDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiLicenseLegalComponentDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApiLicenseLegalComponentDTO) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalComponentDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApiLicenseLegalComponentDTO) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApiLicenseLegalComponentDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiLicenseLegalComponentDTO) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalComponentDTO) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiLicenseLegalComponentDTO) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiLicenseLegalComponentDTO) SetHash(v string) {
	o.Hash = &v
}

// GetLicenseLegalData returns the LicenseLegalData field value if set, zero value otherwise.
func (o *ApiLicenseLegalComponentDTO) GetLicenseLegalData() ApiLicenseLegalDataDTO {
	if o == nil || IsNil(o.LicenseLegalData) {
		var ret ApiLicenseLegalDataDTO
		return ret
	}
	return *o.LicenseLegalData
}

// GetLicenseLegalDataOk returns a tuple with the LicenseLegalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalComponentDTO) GetLicenseLegalDataOk() (*ApiLicenseLegalDataDTO, bool) {
	if o == nil || IsNil(o.LicenseLegalData) {
		return nil, false
	}
	return o.LicenseLegalData, true
}

// HasLicenseLegalData returns a boolean if a field has been set.
func (o *ApiLicenseLegalComponentDTO) HasLicenseLegalData() bool {
	if o != nil && !IsNil(o.LicenseLegalData) {
		return true
	}

	return false
}

// SetLicenseLegalData gets a reference to the given ApiLicenseLegalDataDTO and assigns it to the LicenseLegalData field.
func (o *ApiLicenseLegalComponentDTO) SetLicenseLegalData(v ApiLicenseLegalDataDTO) {
	o.LicenseLegalData = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiLicenseLegalComponentDTO) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalComponentDTO) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiLicenseLegalComponentDTO) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiLicenseLegalComponentDTO) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetStageScans returns the StageScans field value if set, zero value otherwise.
func (o *ApiLicenseLegalComponentDTO) GetStageScans() []ApiLicenseLegalStageScanDTO {
	if o == nil || IsNil(o.StageScans) {
		var ret []ApiLicenseLegalStageScanDTO
		return ret
	}
	return o.StageScans
}

// GetStageScansOk returns a tuple with the StageScans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalComponentDTO) GetStageScansOk() ([]ApiLicenseLegalStageScanDTO, bool) {
	if o == nil || IsNil(o.StageScans) {
		return nil, false
	}
	return o.StageScans, true
}

// HasStageScans returns a boolean if a field has been set.
func (o *ApiLicenseLegalComponentDTO) HasStageScans() bool {
	if o != nil && !IsNil(o.StageScans) {
		return true
	}

	return false
}

// SetStageScans gets a reference to the given []ApiLicenseLegalStageScanDTO and assigns it to the StageScans field.
func (o *ApiLicenseLegalComponentDTO) SetStageScans(v []ApiLicenseLegalStageScanDTO) {
	o.StageScans = v
}

func (o ApiLicenseLegalComponentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseLegalComponentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.LicenseLegalData) {
		toSerialize["licenseLegalData"] = o.LicenseLegalData
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	if !IsNil(o.StageScans) {
		toSerialize["stageScans"] = o.StageScans
	}
	return toSerialize, nil
}

type NullableApiLicenseLegalComponentDTO struct {
	value *ApiLicenseLegalComponentDTO
	isSet bool
}

func (v NullableApiLicenseLegalComponentDTO) Get() *ApiLicenseLegalComponentDTO {
	return v.value
}

func (v *NullableApiLicenseLegalComponentDTO) Set(val *ApiLicenseLegalComponentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseLegalComponentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseLegalComponentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseLegalComponentDTO(val *ApiLicenseLegalComponentDTO) *NullableApiLicenseLegalComponentDTO {
	return &NullableApiLicenseLegalComponentDTO{value: val, isSet: true}
}

func (v NullableApiLicenseLegalComponentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseLegalComponentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


