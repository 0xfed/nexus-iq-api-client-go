/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.181.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the SbomComponentListDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SbomComponentListDTO{}

// SbomComponentListDTO struct for SbomComponentListDTO
type SbomComponentListDTO struct {
	Results []SbomComponentDTO `json:"results,omitempty"`
	TotalResultsCount *int32 `json:"totalResultsCount,omitempty"`
}

// NewSbomComponentListDTO instantiates a new SbomComponentListDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSbomComponentListDTO() *SbomComponentListDTO {
	this := SbomComponentListDTO{}
	return &this
}

// NewSbomComponentListDTOWithDefaults instantiates a new SbomComponentListDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSbomComponentListDTOWithDefaults() *SbomComponentListDTO {
	this := SbomComponentListDTO{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SbomComponentListDTO) GetResults() []SbomComponentDTO {
	if o == nil || IsNil(o.Results) {
		var ret []SbomComponentDTO
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SbomComponentListDTO) GetResultsOk() ([]SbomComponentDTO, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SbomComponentListDTO) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SbomComponentDTO and assigns it to the Results field.
func (o *SbomComponentListDTO) SetResults(v []SbomComponentDTO) {
	o.Results = v
}

// GetTotalResultsCount returns the TotalResultsCount field value if set, zero value otherwise.
func (o *SbomComponentListDTO) GetTotalResultsCount() int32 {
	if o == nil || IsNil(o.TotalResultsCount) {
		var ret int32
		return ret
	}
	return *o.TotalResultsCount
}

// GetTotalResultsCountOk returns a tuple with the TotalResultsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SbomComponentListDTO) GetTotalResultsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResultsCount) {
		return nil, false
	}
	return o.TotalResultsCount, true
}

// HasTotalResultsCount returns a boolean if a field has been set.
func (o *SbomComponentListDTO) HasTotalResultsCount() bool {
	if o != nil && !IsNil(o.TotalResultsCount) {
		return true
	}

	return false
}

// SetTotalResultsCount gets a reference to the given int32 and assigns it to the TotalResultsCount field.
func (o *SbomComponentListDTO) SetTotalResultsCount(v int32) {
	o.TotalResultsCount = &v
}

func (o SbomComponentListDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SbomComponentListDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.TotalResultsCount) {
		toSerialize["totalResultsCount"] = o.TotalResultsCount
	}
	return toSerialize, nil
}

type NullableSbomComponentListDTO struct {
	value *SbomComponentListDTO
	isSet bool
}

func (v NullableSbomComponentListDTO) Get() *SbomComponentListDTO {
	return v.value
}

func (v *NullableSbomComponentListDTO) Set(val *SbomComponentListDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSbomComponentListDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSbomComponentListDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSbomComponentListDTO(val *SbomComponentListDTO) *NullableSbomComponentListDTO {
	return &NullableSbomComponentListDTO{value: val, isSet: true}
}

func (v NullableSbomComponentListDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSbomComponentListDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


