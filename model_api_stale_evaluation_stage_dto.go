/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.181.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiStaleEvaluationStageDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStaleEvaluationStageDTO{}

// ApiStaleEvaluationStageDTO struct for ApiStaleEvaluationStageDTO
type ApiStaleEvaluationStageDTO struct {
	LastEvaluationDate *time.Time `json:"lastEvaluationDate,omitempty"`
	StageId *string `json:"stageId,omitempty"`
}

// NewApiStaleEvaluationStageDTO instantiates a new ApiStaleEvaluationStageDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStaleEvaluationStageDTO() *ApiStaleEvaluationStageDTO {
	this := ApiStaleEvaluationStageDTO{}
	return &this
}

// NewApiStaleEvaluationStageDTOWithDefaults instantiates a new ApiStaleEvaluationStageDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStaleEvaluationStageDTOWithDefaults() *ApiStaleEvaluationStageDTO {
	this := ApiStaleEvaluationStageDTO{}
	return &this
}

// GetLastEvaluationDate returns the LastEvaluationDate field value if set, zero value otherwise.
func (o *ApiStaleEvaluationStageDTO) GetLastEvaluationDate() time.Time {
	if o == nil || IsNil(o.LastEvaluationDate) {
		var ret time.Time
		return ret
	}
	return *o.LastEvaluationDate
}

// GetLastEvaluationDateOk returns a tuple with the LastEvaluationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleEvaluationStageDTO) GetLastEvaluationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastEvaluationDate) {
		return nil, false
	}
	return o.LastEvaluationDate, true
}

// HasLastEvaluationDate returns a boolean if a field has been set.
func (o *ApiStaleEvaluationStageDTO) HasLastEvaluationDate() bool {
	if o != nil && !IsNil(o.LastEvaluationDate) {
		return true
	}

	return false
}

// SetLastEvaluationDate gets a reference to the given time.Time and assigns it to the LastEvaluationDate field.
func (o *ApiStaleEvaluationStageDTO) SetLastEvaluationDate(v time.Time) {
	o.LastEvaluationDate = &v
}

// GetStageId returns the StageId field value if set, zero value otherwise.
func (o *ApiStaleEvaluationStageDTO) GetStageId() string {
	if o == nil || IsNil(o.StageId) {
		var ret string
		return ret
	}
	return *o.StageId
}

// GetStageIdOk returns a tuple with the StageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleEvaluationStageDTO) GetStageIdOk() (*string, bool) {
	if o == nil || IsNil(o.StageId) {
		return nil, false
	}
	return o.StageId, true
}

// HasStageId returns a boolean if a field has been set.
func (o *ApiStaleEvaluationStageDTO) HasStageId() bool {
	if o != nil && !IsNil(o.StageId) {
		return true
	}

	return false
}

// SetStageId gets a reference to the given string and assigns it to the StageId field.
func (o *ApiStaleEvaluationStageDTO) SetStageId(v string) {
	o.StageId = &v
}

func (o ApiStaleEvaluationStageDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStaleEvaluationStageDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastEvaluationDate) {
		toSerialize["lastEvaluationDate"] = o.LastEvaluationDate
	}
	if !IsNil(o.StageId) {
		toSerialize["stageId"] = o.StageId
	}
	return toSerialize, nil
}

type NullableApiStaleEvaluationStageDTO struct {
	value *ApiStaleEvaluationStageDTO
	isSet bool
}

func (v NullableApiStaleEvaluationStageDTO) Get() *ApiStaleEvaluationStageDTO {
	return v.value
}

func (v *NullableApiStaleEvaluationStageDTO) Set(val *ApiStaleEvaluationStageDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStaleEvaluationStageDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStaleEvaluationStageDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStaleEvaluationStageDTO(val *ApiStaleEvaluationStageDTO) *NullableApiStaleEvaluationStageDTO {
	return &NullableApiStaleEvaluationStageDTO{value: val, isSet: true}
}

func (v NullableApiStaleEvaluationStageDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStaleEvaluationStageDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


