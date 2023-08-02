/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiReportPolicyDataDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReportPolicyDataDTOV2{}

// ApiReportPolicyDataDTOV2 struct for ApiReportPolicyDataDTOV2
type ApiReportPolicyDataDTOV2 struct {
	Application *ApiApplicationBaseDTO `json:"application,omitempty"`
	CommitHash *string `json:"commitHash,omitempty"`
	Components []ApiReportComponentPolicyViolationsDTOV2 `json:"components,omitempty"`
	Counts *map[string]int32 `json:"counts,omitempty"`
	Initiator *string `json:"initiator,omitempty"`
	ReportTime *time.Time `json:"reportTime,omitempty"`
	ReportTitle *string `json:"reportTitle,omitempty"`
}

// NewApiReportPolicyDataDTOV2 instantiates a new ApiReportPolicyDataDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportPolicyDataDTOV2() *ApiReportPolicyDataDTOV2 {
	this := ApiReportPolicyDataDTOV2{}
	return &this
}

// NewApiReportPolicyDataDTOV2WithDefaults instantiates a new ApiReportPolicyDataDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportPolicyDataDTOV2WithDefaults() *ApiReportPolicyDataDTOV2 {
	this := ApiReportPolicyDataDTOV2{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ApiReportPolicyDataDTOV2) GetApplication() ApiApplicationBaseDTO {
	if o == nil || IsNil(o.Application) {
		var ret ApiApplicationBaseDTO
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportPolicyDataDTOV2) GetApplicationOk() (*ApiApplicationBaseDTO, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ApiReportPolicyDataDTOV2) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given ApiApplicationBaseDTO and assigns it to the Application field.
func (o *ApiReportPolicyDataDTOV2) SetApplication(v ApiApplicationBaseDTO) {
	o.Application = &v
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise.
func (o *ApiReportPolicyDataDTOV2) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash) {
		var ret string
		return ret
	}
	return *o.CommitHash
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportPolicyDataDTOV2) GetCommitHashOk() (*string, bool) {
	if o == nil || IsNil(o.CommitHash) {
		return nil, false
	}
	return o.CommitHash, true
}

// HasCommitHash returns a boolean if a field has been set.
func (o *ApiReportPolicyDataDTOV2) HasCommitHash() bool {
	if o != nil && !IsNil(o.CommitHash) {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given string and assigns it to the CommitHash field.
func (o *ApiReportPolicyDataDTOV2) SetCommitHash(v string) {
	o.CommitHash = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ApiReportPolicyDataDTOV2) GetComponents() []ApiReportComponentPolicyViolationsDTOV2 {
	if o == nil || IsNil(o.Components) {
		var ret []ApiReportComponentPolicyViolationsDTOV2
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportPolicyDataDTOV2) GetComponentsOk() ([]ApiReportComponentPolicyViolationsDTOV2, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ApiReportPolicyDataDTOV2) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ApiReportComponentPolicyViolationsDTOV2 and assigns it to the Components field.
func (o *ApiReportPolicyDataDTOV2) SetComponents(v []ApiReportComponentPolicyViolationsDTOV2) {
	o.Components = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *ApiReportPolicyDataDTOV2) GetCounts() map[string]int32 {
	if o == nil || IsNil(o.Counts) {
		var ret map[string]int32
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportPolicyDataDTOV2) GetCountsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *ApiReportPolicyDataDTOV2) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given map[string]int32 and assigns it to the Counts field.
func (o *ApiReportPolicyDataDTOV2) SetCounts(v map[string]int32) {
	o.Counts = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *ApiReportPolicyDataDTOV2) GetInitiator() string {
	if o == nil || IsNil(o.Initiator) {
		var ret string
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportPolicyDataDTOV2) GetInitiatorOk() (*string, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *ApiReportPolicyDataDTOV2) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given string and assigns it to the Initiator field.
func (o *ApiReportPolicyDataDTOV2) SetInitiator(v string) {
	o.Initiator = &v
}

// GetReportTime returns the ReportTime field value if set, zero value otherwise.
func (o *ApiReportPolicyDataDTOV2) GetReportTime() time.Time {
	if o == nil || IsNil(o.ReportTime) {
		var ret time.Time
		return ret
	}
	return *o.ReportTime
}

// GetReportTimeOk returns a tuple with the ReportTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportPolicyDataDTOV2) GetReportTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReportTime) {
		return nil, false
	}
	return o.ReportTime, true
}

// HasReportTime returns a boolean if a field has been set.
func (o *ApiReportPolicyDataDTOV2) HasReportTime() bool {
	if o != nil && !IsNil(o.ReportTime) {
		return true
	}

	return false
}

// SetReportTime gets a reference to the given time.Time and assigns it to the ReportTime field.
func (o *ApiReportPolicyDataDTOV2) SetReportTime(v time.Time) {
	o.ReportTime = &v
}

// GetReportTitle returns the ReportTitle field value if set, zero value otherwise.
func (o *ApiReportPolicyDataDTOV2) GetReportTitle() string {
	if o == nil || IsNil(o.ReportTitle) {
		var ret string
		return ret
	}
	return *o.ReportTitle
}

// GetReportTitleOk returns a tuple with the ReportTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportPolicyDataDTOV2) GetReportTitleOk() (*string, bool) {
	if o == nil || IsNil(o.ReportTitle) {
		return nil, false
	}
	return o.ReportTitle, true
}

// HasReportTitle returns a boolean if a field has been set.
func (o *ApiReportPolicyDataDTOV2) HasReportTitle() bool {
	if o != nil && !IsNil(o.ReportTitle) {
		return true
	}

	return false
}

// SetReportTitle gets a reference to the given string and assigns it to the ReportTitle field.
func (o *ApiReportPolicyDataDTOV2) SetReportTitle(v string) {
	o.ReportTitle = &v
}

func (o ApiReportPolicyDataDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportPolicyDataDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.CommitHash) {
		toSerialize["commitHash"] = o.CommitHash
	}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	if !IsNil(o.ReportTime) {
		toSerialize["reportTime"] = o.ReportTime
	}
	if !IsNil(o.ReportTitle) {
		toSerialize["reportTitle"] = o.ReportTitle
	}
	return toSerialize, nil
}

type NullableApiReportPolicyDataDTOV2 struct {
	value *ApiReportPolicyDataDTOV2
	isSet bool
}

func (v NullableApiReportPolicyDataDTOV2) Get() *ApiReportPolicyDataDTOV2 {
	return v.value
}

func (v *NullableApiReportPolicyDataDTOV2) Set(val *ApiReportPolicyDataDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportPolicyDataDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportPolicyDataDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportPolicyDataDTOV2(val *ApiReportPolicyDataDTOV2) *NullableApiReportPolicyDataDTOV2 {
	return &NullableApiReportPolicyDataDTOV2{value: val, isSet: true}
}

func (v NullableApiReportPolicyDataDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportPolicyDataDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


