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

// checks if the ApiReportComponentPolicyViolationsDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReportComponentPolicyViolationsDTOV2{}

// ApiReportComponentPolicyViolationsDTOV2 struct for ApiReportComponentPolicyViolationsDTOV2
type ApiReportComponentPolicyViolationsDTOV2 struct {
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	DependencyData *ApiDependencyDataDTO `json:"dependencyData,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Hash *string `json:"hash,omitempty"`
	MatchState *string `json:"matchState,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	Pathnames []string `json:"pathnames,omitempty"`
	Proprietary *bool `json:"proprietary,omitempty"`
	Sha256 *string `json:"sha256,omitempty"`
	ThirdParty *bool `json:"thirdParty,omitempty"`
	Violations []ApiReportPolicyViolationDTOV2 `json:"violations,omitempty"`
}

// NewApiReportComponentPolicyViolationsDTOV2 instantiates a new ApiReportComponentPolicyViolationsDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportComponentPolicyViolationsDTOV2() *ApiReportComponentPolicyViolationsDTOV2 {
	this := ApiReportComponentPolicyViolationsDTOV2{}
	return &this
}

// NewApiReportComponentPolicyViolationsDTOV2WithDefaults instantiates a new ApiReportComponentPolicyViolationsDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportComponentPolicyViolationsDTOV2WithDefaults() *ApiReportComponentPolicyViolationsDTOV2 {
	this := ApiReportComponentPolicyViolationsDTOV2{}
	return &this
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetDependencyData returns the DependencyData field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetDependencyData() ApiDependencyDataDTO {
	if o == nil || IsNil(o.DependencyData) {
		var ret ApiDependencyDataDTO
		return ret
	}
	return *o.DependencyData
}

// GetDependencyDataOk returns a tuple with the DependencyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetDependencyDataOk() (*ApiDependencyDataDTO, bool) {
	if o == nil || IsNil(o.DependencyData) {
		return nil, false
	}
	return o.DependencyData, true
}

// HasDependencyData returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasDependencyData() bool {
	if o != nil && !IsNil(o.DependencyData) {
		return true
	}

	return false
}

// SetDependencyData gets a reference to the given ApiDependencyDataDTO and assigns it to the DependencyData field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetDependencyData(v ApiDependencyDataDTO) {
	o.DependencyData = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetHash(v string) {
	o.Hash = &v
}

// GetMatchState returns the MatchState field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetMatchState() string {
	if o == nil || IsNil(o.MatchState) {
		var ret string
		return ret
	}
	return *o.MatchState
}

// GetMatchStateOk returns a tuple with the MatchState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetMatchStateOk() (*string, bool) {
	if o == nil || IsNil(o.MatchState) {
		return nil, false
	}
	return o.MatchState, true
}

// HasMatchState returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasMatchState() bool {
	if o != nil && !IsNil(o.MatchState) {
		return true
	}

	return false
}

// SetMatchState gets a reference to the given string and assigns it to the MatchState field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetMatchState(v string) {
	o.MatchState = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetPathnames returns the Pathnames field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetPathnames() []string {
	if o == nil || IsNil(o.Pathnames) {
		var ret []string
		return ret
	}
	return o.Pathnames
}

// GetPathnamesOk returns a tuple with the Pathnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetPathnamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Pathnames) {
		return nil, false
	}
	return o.Pathnames, true
}

// HasPathnames returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasPathnames() bool {
	if o != nil && !IsNil(o.Pathnames) {
		return true
	}

	return false
}

// SetPathnames gets a reference to the given []string and assigns it to the Pathnames field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetPathnames(v []string) {
	o.Pathnames = v
}

// GetProprietary returns the Proprietary field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetProprietary() bool {
	if o == nil || IsNil(o.Proprietary) {
		var ret bool
		return ret
	}
	return *o.Proprietary
}

// GetProprietaryOk returns a tuple with the Proprietary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetProprietaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Proprietary) {
		return nil, false
	}
	return o.Proprietary, true
}

// HasProprietary returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasProprietary() bool {
	if o != nil && !IsNil(o.Proprietary) {
		return true
	}

	return false
}

// SetProprietary gets a reference to the given bool and assigns it to the Proprietary field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetProprietary(v bool) {
	o.Proprietary = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetSha256() string {
	if o == nil || IsNil(o.Sha256) {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha256) {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasSha256() bool {
	if o != nil && !IsNil(o.Sha256) {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetSha256(v string) {
	o.Sha256 = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetThirdParty() bool {
	if o == nil || IsNil(o.ThirdParty) {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetThirdPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.ThirdParty) {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasThirdParty() bool {
	if o != nil && !IsNil(o.ThirdParty) {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

// GetViolations returns the Violations field value if set, zero value otherwise.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetViolations() []ApiReportPolicyViolationDTOV2 {
	if o == nil || IsNil(o.Violations) {
		var ret []ApiReportPolicyViolationDTOV2
		return ret
	}
	return o.Violations
}

// GetViolationsOk returns a tuple with the Violations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) GetViolationsOk() ([]ApiReportPolicyViolationDTOV2, bool) {
	if o == nil || IsNil(o.Violations) {
		return nil, false
	}
	return o.Violations, true
}

// HasViolations returns a boolean if a field has been set.
func (o *ApiReportComponentPolicyViolationsDTOV2) HasViolations() bool {
	if o != nil && !IsNil(o.Violations) {
		return true
	}

	return false
}

// SetViolations gets a reference to the given []ApiReportPolicyViolationDTOV2 and assigns it to the Violations field.
func (o *ApiReportComponentPolicyViolationsDTOV2) SetViolations(v []ApiReportPolicyViolationDTOV2) {
	o.Violations = v
}

func (o ApiReportComponentPolicyViolationsDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportComponentPolicyViolationsDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.DependencyData) {
		toSerialize["dependencyData"] = o.DependencyData
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.MatchState) {
		toSerialize["matchState"] = o.MatchState
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	if !IsNil(o.Pathnames) {
		toSerialize["pathnames"] = o.Pathnames
	}
	if !IsNil(o.Proprietary) {
		toSerialize["proprietary"] = o.Proprietary
	}
	if !IsNil(o.Sha256) {
		toSerialize["sha256"] = o.Sha256
	}
	if !IsNil(o.ThirdParty) {
		toSerialize["thirdParty"] = o.ThirdParty
	}
	if !IsNil(o.Violations) {
		toSerialize["violations"] = o.Violations
	}
	return toSerialize, nil
}

type NullableApiReportComponentPolicyViolationsDTOV2 struct {
	value *ApiReportComponentPolicyViolationsDTOV2
	isSet bool
}

func (v NullableApiReportComponentPolicyViolationsDTOV2) Get() *ApiReportComponentPolicyViolationsDTOV2 {
	return v.value
}

func (v *NullableApiReportComponentPolicyViolationsDTOV2) Set(val *ApiReportComponentPolicyViolationsDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportComponentPolicyViolationsDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportComponentPolicyViolationsDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportComponentPolicyViolationsDTOV2(val *ApiReportComponentPolicyViolationsDTOV2) *NullableApiReportComponentPolicyViolationsDTOV2 {
	return &NullableApiReportComponentPolicyViolationsDTOV2{value: val, isSet: true}
}

func (v NullableApiReportComponentPolicyViolationsDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportComponentPolicyViolationsDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


