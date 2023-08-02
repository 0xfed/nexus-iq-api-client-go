/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 165
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiPromoteScanRequestDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPromoteScanRequestDTOV2{}

// ApiPromoteScanRequestDTOV2 struct for ApiPromoteScanRequestDTOV2
type ApiPromoteScanRequestDTOV2 struct {
	ScanId *string `json:"scanId,omitempty"`
	SourceStageId *string `json:"sourceStageId,omitempty"`
	TargetStageId *string `json:"targetStageId,omitempty"`
}

// NewApiPromoteScanRequestDTOV2 instantiates a new ApiPromoteScanRequestDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPromoteScanRequestDTOV2() *ApiPromoteScanRequestDTOV2 {
	this := ApiPromoteScanRequestDTOV2{}
	return &this
}

// NewApiPromoteScanRequestDTOV2WithDefaults instantiates a new ApiPromoteScanRequestDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPromoteScanRequestDTOV2WithDefaults() *ApiPromoteScanRequestDTOV2 {
	this := ApiPromoteScanRequestDTOV2{}
	return &this
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *ApiPromoteScanRequestDTOV2) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPromoteScanRequestDTOV2) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *ApiPromoteScanRequestDTOV2) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *ApiPromoteScanRequestDTOV2) SetScanId(v string) {
	o.ScanId = &v
}

// GetSourceStageId returns the SourceStageId field value if set, zero value otherwise.
func (o *ApiPromoteScanRequestDTOV2) GetSourceStageId() string {
	if o == nil || IsNil(o.SourceStageId) {
		var ret string
		return ret
	}
	return *o.SourceStageId
}

// GetSourceStageIdOk returns a tuple with the SourceStageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPromoteScanRequestDTOV2) GetSourceStageIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceStageId) {
		return nil, false
	}
	return o.SourceStageId, true
}

// HasSourceStageId returns a boolean if a field has been set.
func (o *ApiPromoteScanRequestDTOV2) HasSourceStageId() bool {
	if o != nil && !IsNil(o.SourceStageId) {
		return true
	}

	return false
}

// SetSourceStageId gets a reference to the given string and assigns it to the SourceStageId field.
func (o *ApiPromoteScanRequestDTOV2) SetSourceStageId(v string) {
	o.SourceStageId = &v
}

// GetTargetStageId returns the TargetStageId field value if set, zero value otherwise.
func (o *ApiPromoteScanRequestDTOV2) GetTargetStageId() string {
	if o == nil || IsNil(o.TargetStageId) {
		var ret string
		return ret
	}
	return *o.TargetStageId
}

// GetTargetStageIdOk returns a tuple with the TargetStageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPromoteScanRequestDTOV2) GetTargetStageIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetStageId) {
		return nil, false
	}
	return o.TargetStageId, true
}

// HasTargetStageId returns a boolean if a field has been set.
func (o *ApiPromoteScanRequestDTOV2) HasTargetStageId() bool {
	if o != nil && !IsNil(o.TargetStageId) {
		return true
	}

	return false
}

// SetTargetStageId gets a reference to the given string and assigns it to the TargetStageId field.
func (o *ApiPromoteScanRequestDTOV2) SetTargetStageId(v string) {
	o.TargetStageId = &v
}

func (o ApiPromoteScanRequestDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPromoteScanRequestDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScanId) {
		toSerialize["scanId"] = o.ScanId
	}
	if !IsNil(o.SourceStageId) {
		toSerialize["sourceStageId"] = o.SourceStageId
	}
	if !IsNil(o.TargetStageId) {
		toSerialize["targetStageId"] = o.TargetStageId
	}
	return toSerialize, nil
}

type NullableApiPromoteScanRequestDTOV2 struct {
	value *ApiPromoteScanRequestDTOV2
	isSet bool
}

func (v NullableApiPromoteScanRequestDTOV2) Get() *ApiPromoteScanRequestDTOV2 {
	return v.value
}

func (v *NullableApiPromoteScanRequestDTOV2) Set(val *ApiPromoteScanRequestDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPromoteScanRequestDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPromoteScanRequestDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPromoteScanRequestDTOV2(val *ApiPromoteScanRequestDTOV2) *NullableApiPromoteScanRequestDTOV2 {
	return &NullableApiPromoteScanRequestDTOV2{value: val, isSet: true}
}

func (v NullableApiPromoteScanRequestDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPromoteScanRequestDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


