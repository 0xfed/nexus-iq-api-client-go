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

// checks if the SbomsAnalyzedMetricsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SbomsAnalyzedMetricsDTO{}

// SbomsAnalyzedMetricsDTO struct for SbomsAnalyzedMetricsDTO
type SbomsAnalyzedMetricsDTO struct {
	Threshold *int64 `json:"threshold,omitempty"`
	Total *int64 `json:"total,omitempty"`
}

// NewSbomsAnalyzedMetricsDTO instantiates a new SbomsAnalyzedMetricsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSbomsAnalyzedMetricsDTO() *SbomsAnalyzedMetricsDTO {
	this := SbomsAnalyzedMetricsDTO{}
	return &this
}

// NewSbomsAnalyzedMetricsDTOWithDefaults instantiates a new SbomsAnalyzedMetricsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSbomsAnalyzedMetricsDTOWithDefaults() *SbomsAnalyzedMetricsDTO {
	this := SbomsAnalyzedMetricsDTO{}
	return &this
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *SbomsAnalyzedMetricsDTO) GetThreshold() int64 {
	if o == nil || IsNil(o.Threshold) {
		var ret int64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SbomsAnalyzedMetricsDTO) GetThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *SbomsAnalyzedMetricsDTO) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int64 and assigns it to the Threshold field.
func (o *SbomsAnalyzedMetricsDTO) SetThreshold(v int64) {
	o.Threshold = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SbomsAnalyzedMetricsDTO) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SbomsAnalyzedMetricsDTO) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SbomsAnalyzedMetricsDTO) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *SbomsAnalyzedMetricsDTO) SetTotal(v int64) {
	o.Total = &v
}

func (o SbomsAnalyzedMetricsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SbomsAnalyzedMetricsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableSbomsAnalyzedMetricsDTO struct {
	value *SbomsAnalyzedMetricsDTO
	isSet bool
}

func (v NullableSbomsAnalyzedMetricsDTO) Get() *SbomsAnalyzedMetricsDTO {
	return v.value
}

func (v *NullableSbomsAnalyzedMetricsDTO) Set(val *SbomsAnalyzedMetricsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSbomsAnalyzedMetricsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSbomsAnalyzedMetricsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSbomsAnalyzedMetricsDTO(val *SbomsAnalyzedMetricsDTO) *NullableSbomsAnalyzedMetricsDTO {
	return &NullableSbomsAnalyzedMetricsDTO{value: val, isSet: true}
}

func (v NullableSbomsAnalyzedMetricsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSbomsAnalyzedMetricsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


