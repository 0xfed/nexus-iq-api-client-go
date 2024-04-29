/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.175.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiPolicyViolationDiffDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPolicyViolationDiffDTO{}

// ApiPolicyViolationDiffDTO struct for ApiPolicyViolationDiffDTO
type ApiPolicyViolationDiffDTO struct {
	AddedViolations []ApiPolicyViolationForDiffDTO `json:"addedViolations,omitempty"`
	Application *ApiApplicationDTO `json:"application,omitempty"`
	DiffTime *time.Time `json:"diffTime,omitempty"`
	FromCommit *ApiApplicationEvaluationCommitDTO `json:"fromCommit,omitempty"`
	RemovedViolations []ApiPolicyViolationForDiffDTO `json:"removedViolations,omitempty"`
	SameViolations []ApiPolicyViolationForDiffDTO `json:"sameViolations,omitempty"`
	ToCommit *ApiApplicationEvaluationCommitDTO `json:"toCommit,omitempty"`
}

// NewApiPolicyViolationDiffDTO instantiates a new ApiPolicyViolationDiffDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPolicyViolationDiffDTO() *ApiPolicyViolationDiffDTO {
	this := ApiPolicyViolationDiffDTO{}
	return &this
}

// NewApiPolicyViolationDiffDTOWithDefaults instantiates a new ApiPolicyViolationDiffDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPolicyViolationDiffDTOWithDefaults() *ApiPolicyViolationDiffDTO {
	this := ApiPolicyViolationDiffDTO{}
	return &this
}

// GetAddedViolations returns the AddedViolations field value if set, zero value otherwise.
func (o *ApiPolicyViolationDiffDTO) GetAddedViolations() []ApiPolicyViolationForDiffDTO {
	if o == nil || IsNil(o.AddedViolations) {
		var ret []ApiPolicyViolationForDiffDTO
		return ret
	}
	return o.AddedViolations
}

// GetAddedViolationsOk returns a tuple with the AddedViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyViolationDiffDTO) GetAddedViolationsOk() ([]ApiPolicyViolationForDiffDTO, bool) {
	if o == nil || IsNil(o.AddedViolations) {
		return nil, false
	}
	return o.AddedViolations, true
}

// HasAddedViolations returns a boolean if a field has been set.
func (o *ApiPolicyViolationDiffDTO) HasAddedViolations() bool {
	if o != nil && !IsNil(o.AddedViolations) {
		return true
	}

	return false
}

// SetAddedViolations gets a reference to the given []ApiPolicyViolationForDiffDTO and assigns it to the AddedViolations field.
func (o *ApiPolicyViolationDiffDTO) SetAddedViolations(v []ApiPolicyViolationForDiffDTO) {
	o.AddedViolations = v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ApiPolicyViolationDiffDTO) GetApplication() ApiApplicationDTO {
	if o == nil || IsNil(o.Application) {
		var ret ApiApplicationDTO
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyViolationDiffDTO) GetApplicationOk() (*ApiApplicationDTO, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ApiPolicyViolationDiffDTO) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given ApiApplicationDTO and assigns it to the Application field.
func (o *ApiPolicyViolationDiffDTO) SetApplication(v ApiApplicationDTO) {
	o.Application = &v
}

// GetDiffTime returns the DiffTime field value if set, zero value otherwise.
func (o *ApiPolicyViolationDiffDTO) GetDiffTime() time.Time {
	if o == nil || IsNil(o.DiffTime) {
		var ret time.Time
		return ret
	}
	return *o.DiffTime
}

// GetDiffTimeOk returns a tuple with the DiffTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyViolationDiffDTO) GetDiffTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DiffTime) {
		return nil, false
	}
	return o.DiffTime, true
}

// HasDiffTime returns a boolean if a field has been set.
func (o *ApiPolicyViolationDiffDTO) HasDiffTime() bool {
	if o != nil && !IsNil(o.DiffTime) {
		return true
	}

	return false
}

// SetDiffTime gets a reference to the given time.Time and assigns it to the DiffTime field.
func (o *ApiPolicyViolationDiffDTO) SetDiffTime(v time.Time) {
	o.DiffTime = &v
}

// GetFromCommit returns the FromCommit field value if set, zero value otherwise.
func (o *ApiPolicyViolationDiffDTO) GetFromCommit() ApiApplicationEvaluationCommitDTO {
	if o == nil || IsNil(o.FromCommit) {
		var ret ApiApplicationEvaluationCommitDTO
		return ret
	}
	return *o.FromCommit
}

// GetFromCommitOk returns a tuple with the FromCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyViolationDiffDTO) GetFromCommitOk() (*ApiApplicationEvaluationCommitDTO, bool) {
	if o == nil || IsNil(o.FromCommit) {
		return nil, false
	}
	return o.FromCommit, true
}

// HasFromCommit returns a boolean if a field has been set.
func (o *ApiPolicyViolationDiffDTO) HasFromCommit() bool {
	if o != nil && !IsNil(o.FromCommit) {
		return true
	}

	return false
}

// SetFromCommit gets a reference to the given ApiApplicationEvaluationCommitDTO and assigns it to the FromCommit field.
func (o *ApiPolicyViolationDiffDTO) SetFromCommit(v ApiApplicationEvaluationCommitDTO) {
	o.FromCommit = &v
}

// GetRemovedViolations returns the RemovedViolations field value if set, zero value otherwise.
func (o *ApiPolicyViolationDiffDTO) GetRemovedViolations() []ApiPolicyViolationForDiffDTO {
	if o == nil || IsNil(o.RemovedViolations) {
		var ret []ApiPolicyViolationForDiffDTO
		return ret
	}
	return o.RemovedViolations
}

// GetRemovedViolationsOk returns a tuple with the RemovedViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyViolationDiffDTO) GetRemovedViolationsOk() ([]ApiPolicyViolationForDiffDTO, bool) {
	if o == nil || IsNil(o.RemovedViolations) {
		return nil, false
	}
	return o.RemovedViolations, true
}

// HasRemovedViolations returns a boolean if a field has been set.
func (o *ApiPolicyViolationDiffDTO) HasRemovedViolations() bool {
	if o != nil && !IsNil(o.RemovedViolations) {
		return true
	}

	return false
}

// SetRemovedViolations gets a reference to the given []ApiPolicyViolationForDiffDTO and assigns it to the RemovedViolations field.
func (o *ApiPolicyViolationDiffDTO) SetRemovedViolations(v []ApiPolicyViolationForDiffDTO) {
	o.RemovedViolations = v
}

// GetSameViolations returns the SameViolations field value if set, zero value otherwise.
func (o *ApiPolicyViolationDiffDTO) GetSameViolations() []ApiPolicyViolationForDiffDTO {
	if o == nil || IsNil(o.SameViolations) {
		var ret []ApiPolicyViolationForDiffDTO
		return ret
	}
	return o.SameViolations
}

// GetSameViolationsOk returns a tuple with the SameViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyViolationDiffDTO) GetSameViolationsOk() ([]ApiPolicyViolationForDiffDTO, bool) {
	if o == nil || IsNil(o.SameViolations) {
		return nil, false
	}
	return o.SameViolations, true
}

// HasSameViolations returns a boolean if a field has been set.
func (o *ApiPolicyViolationDiffDTO) HasSameViolations() bool {
	if o != nil && !IsNil(o.SameViolations) {
		return true
	}

	return false
}

// SetSameViolations gets a reference to the given []ApiPolicyViolationForDiffDTO and assigns it to the SameViolations field.
func (o *ApiPolicyViolationDiffDTO) SetSameViolations(v []ApiPolicyViolationForDiffDTO) {
	o.SameViolations = v
}

// GetToCommit returns the ToCommit field value if set, zero value otherwise.
func (o *ApiPolicyViolationDiffDTO) GetToCommit() ApiApplicationEvaluationCommitDTO {
	if o == nil || IsNil(o.ToCommit) {
		var ret ApiApplicationEvaluationCommitDTO
		return ret
	}
	return *o.ToCommit
}

// GetToCommitOk returns a tuple with the ToCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyViolationDiffDTO) GetToCommitOk() (*ApiApplicationEvaluationCommitDTO, bool) {
	if o == nil || IsNil(o.ToCommit) {
		return nil, false
	}
	return o.ToCommit, true
}

// HasToCommit returns a boolean if a field has been set.
func (o *ApiPolicyViolationDiffDTO) HasToCommit() bool {
	if o != nil && !IsNil(o.ToCommit) {
		return true
	}

	return false
}

// SetToCommit gets a reference to the given ApiApplicationEvaluationCommitDTO and assigns it to the ToCommit field.
func (o *ApiPolicyViolationDiffDTO) SetToCommit(v ApiApplicationEvaluationCommitDTO) {
	o.ToCommit = &v
}

func (o ApiPolicyViolationDiffDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPolicyViolationDiffDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddedViolations) {
		toSerialize["addedViolations"] = o.AddedViolations
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.DiffTime) {
		toSerialize["diffTime"] = o.DiffTime
	}
	if !IsNil(o.FromCommit) {
		toSerialize["fromCommit"] = o.FromCommit
	}
	if !IsNil(o.RemovedViolations) {
		toSerialize["removedViolations"] = o.RemovedViolations
	}
	if !IsNil(o.SameViolations) {
		toSerialize["sameViolations"] = o.SameViolations
	}
	if !IsNil(o.ToCommit) {
		toSerialize["toCommit"] = o.ToCommit
	}
	return toSerialize, nil
}

type NullableApiPolicyViolationDiffDTO struct {
	value *ApiPolicyViolationDiffDTO
	isSet bool
}

func (v NullableApiPolicyViolationDiffDTO) Get() *ApiPolicyViolationDiffDTO {
	return v.value
}

func (v *NullableApiPolicyViolationDiffDTO) Set(val *ApiPolicyViolationDiffDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPolicyViolationDiffDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPolicyViolationDiffDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPolicyViolationDiffDTO(val *ApiPolicyViolationDiffDTO) *NullableApiPolicyViolationDiffDTO {
	return &NullableApiPolicyViolationDiffDTO{value: val, isSet: true}
}

func (v NullableApiPolicyViolationDiffDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPolicyViolationDiffDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


