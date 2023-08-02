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

// checks if the ApiApplicationEvaluationResultDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiApplicationEvaluationResultDTOV2{}

// ApiApplicationEvaluationResultDTOV2 struct for ApiApplicationEvaluationResultDTOV2
type ApiApplicationEvaluationResultDTOV2 struct {
	EmbeddableReportHtmlUrl *string `json:"embeddableReportHtmlUrl,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ReportDataUrl *string `json:"reportDataUrl,omitempty"`
	ReportHtmlUrl *string `json:"reportHtmlUrl,omitempty"`
	ReportPdfUrl *string `json:"reportPdfUrl,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewApiApplicationEvaluationResultDTOV2 instantiates a new ApiApplicationEvaluationResultDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiApplicationEvaluationResultDTOV2() *ApiApplicationEvaluationResultDTOV2 {
	this := ApiApplicationEvaluationResultDTOV2{}
	return &this
}

// NewApiApplicationEvaluationResultDTOV2WithDefaults instantiates a new ApiApplicationEvaluationResultDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiApplicationEvaluationResultDTOV2WithDefaults() *ApiApplicationEvaluationResultDTOV2 {
	this := ApiApplicationEvaluationResultDTOV2{}
	return &this
}

// GetEmbeddableReportHtmlUrl returns the EmbeddableReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationResultDTOV2) GetEmbeddableReportHtmlUrl() string {
	if o == nil || IsNil(o.EmbeddableReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.EmbeddableReportHtmlUrl
}

// GetEmbeddableReportHtmlUrlOk returns a tuple with the EmbeddableReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationResultDTOV2) GetEmbeddableReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EmbeddableReportHtmlUrl) {
		return nil, false
	}
	return o.EmbeddableReportHtmlUrl, true
}

// HasEmbeddableReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationResultDTOV2) HasEmbeddableReportHtmlUrl() bool {
	if o != nil && !IsNil(o.EmbeddableReportHtmlUrl) {
		return true
	}

	return false
}

// SetEmbeddableReportHtmlUrl gets a reference to the given string and assigns it to the EmbeddableReportHtmlUrl field.
func (o *ApiApplicationEvaluationResultDTOV2) SetEmbeddableReportHtmlUrl(v string) {
	o.EmbeddableReportHtmlUrl = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationResultDTOV2) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationResultDTOV2) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationResultDTOV2) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ApiApplicationEvaluationResultDTOV2) SetReason(v string) {
	o.Reason = &v
}

// GetReportDataUrl returns the ReportDataUrl field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationResultDTOV2) GetReportDataUrl() string {
	if o == nil || IsNil(o.ReportDataUrl) {
		var ret string
		return ret
	}
	return *o.ReportDataUrl
}

// GetReportDataUrlOk returns a tuple with the ReportDataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationResultDTOV2) GetReportDataUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportDataUrl) {
		return nil, false
	}
	return o.ReportDataUrl, true
}

// HasReportDataUrl returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationResultDTOV2) HasReportDataUrl() bool {
	if o != nil && !IsNil(o.ReportDataUrl) {
		return true
	}

	return false
}

// SetReportDataUrl gets a reference to the given string and assigns it to the ReportDataUrl field.
func (o *ApiApplicationEvaluationResultDTOV2) SetReportDataUrl(v string) {
	o.ReportDataUrl = &v
}

// GetReportHtmlUrl returns the ReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationResultDTOV2) GetReportHtmlUrl() string {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.ReportHtmlUrl
}

// GetReportHtmlUrlOk returns a tuple with the ReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationResultDTOV2) GetReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		return nil, false
	}
	return o.ReportHtmlUrl, true
}

// HasReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationResultDTOV2) HasReportHtmlUrl() bool {
	if o != nil && !IsNil(o.ReportHtmlUrl) {
		return true
	}

	return false
}

// SetReportHtmlUrl gets a reference to the given string and assigns it to the ReportHtmlUrl field.
func (o *ApiApplicationEvaluationResultDTOV2) SetReportHtmlUrl(v string) {
	o.ReportHtmlUrl = &v
}

// GetReportPdfUrl returns the ReportPdfUrl field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationResultDTOV2) GetReportPdfUrl() string {
	if o == nil || IsNil(o.ReportPdfUrl) {
		var ret string
		return ret
	}
	return *o.ReportPdfUrl
}

// GetReportPdfUrlOk returns a tuple with the ReportPdfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationResultDTOV2) GetReportPdfUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportPdfUrl) {
		return nil, false
	}
	return o.ReportPdfUrl, true
}

// HasReportPdfUrl returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationResultDTOV2) HasReportPdfUrl() bool {
	if o != nil && !IsNil(o.ReportPdfUrl) {
		return true
	}

	return false
}

// SetReportPdfUrl gets a reference to the given string and assigns it to the ReportPdfUrl field.
func (o *ApiApplicationEvaluationResultDTOV2) SetReportPdfUrl(v string) {
	o.ReportPdfUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationResultDTOV2) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationResultDTOV2) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationResultDTOV2) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiApplicationEvaluationResultDTOV2) SetStatus(v string) {
	o.Status = &v
}

func (o ApiApplicationEvaluationResultDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiApplicationEvaluationResultDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmbeddableReportHtmlUrl) {
		toSerialize["embeddableReportHtmlUrl"] = o.EmbeddableReportHtmlUrl
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.ReportDataUrl) {
		toSerialize["reportDataUrl"] = o.ReportDataUrl
	}
	if !IsNil(o.ReportHtmlUrl) {
		toSerialize["reportHtmlUrl"] = o.ReportHtmlUrl
	}
	if !IsNil(o.ReportPdfUrl) {
		toSerialize["reportPdfUrl"] = o.ReportPdfUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableApiApplicationEvaluationResultDTOV2 struct {
	value *ApiApplicationEvaluationResultDTOV2
	isSet bool
}

func (v NullableApiApplicationEvaluationResultDTOV2) Get() *ApiApplicationEvaluationResultDTOV2 {
	return v.value
}

func (v *NullableApiApplicationEvaluationResultDTOV2) Set(val *ApiApplicationEvaluationResultDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiApplicationEvaluationResultDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiApplicationEvaluationResultDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiApplicationEvaluationResultDTOV2(val *ApiApplicationEvaluationResultDTOV2) *NullableApiApplicationEvaluationResultDTOV2 {
	return &NullableApiApplicationEvaluationResultDTOV2{value: val, isSet: true}
}

func (v NullableApiApplicationEvaluationResultDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiApplicationEvaluationResultDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


