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

// checks if the ApiThirdPartyScanResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiThirdPartyScanResultDTO{}

// ApiThirdPartyScanResultDTO struct for ApiThirdPartyScanResultDTO
type ApiThirdPartyScanResultDTO struct {
	ComponentsAffected *ApiEvaluationResultCounterDTO `json:"componentsAffected,omitempty"`
	EmbeddableReportHtmlUrl *string `json:"embeddableReportHtmlUrl,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	GrandfatheredPolicyViolations *int32 `json:"grandfatheredPolicyViolations,omitempty"`
	IsError *bool `json:"isError,omitempty"`
	LegacyViolations *int32 `json:"legacyViolations,omitempty"`
	OpenPolicyViolations *ApiEvaluationResultCounterDTO `json:"openPolicyViolations,omitempty"`
	PolicyAction *string `json:"policyAction,omitempty"`
	ReportDataUrl *string `json:"reportDataUrl,omitempty"`
	ReportHtmlUrl *string `json:"reportHtmlUrl,omitempty"`
	ReportPdfUrl *string `json:"reportPdfUrl,omitempty"`
}

// NewApiThirdPartyScanResultDTO instantiates a new ApiThirdPartyScanResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiThirdPartyScanResultDTO() *ApiThirdPartyScanResultDTO {
	this := ApiThirdPartyScanResultDTO{}
	return &this
}

// NewApiThirdPartyScanResultDTOWithDefaults instantiates a new ApiThirdPartyScanResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiThirdPartyScanResultDTOWithDefaults() *ApiThirdPartyScanResultDTO {
	this := ApiThirdPartyScanResultDTO{}
	return &this
}

// GetComponentsAffected returns the ComponentsAffected field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetComponentsAffected() ApiEvaluationResultCounterDTO {
	if o == nil || IsNil(o.ComponentsAffected) {
		var ret ApiEvaluationResultCounterDTO
		return ret
	}
	return *o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetComponentsAffectedOk() (*ApiEvaluationResultCounterDTO, bool) {
	if o == nil || IsNil(o.ComponentsAffected) {
		return nil, false
	}
	return o.ComponentsAffected, true
}

// HasComponentsAffected returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasComponentsAffected() bool {
	if o != nil && !IsNil(o.ComponentsAffected) {
		return true
	}

	return false
}

// SetComponentsAffected gets a reference to the given ApiEvaluationResultCounterDTO and assigns it to the ComponentsAffected field.
func (o *ApiThirdPartyScanResultDTO) SetComponentsAffected(v ApiEvaluationResultCounterDTO) {
	o.ComponentsAffected = &v
}

// GetEmbeddableReportHtmlUrl returns the EmbeddableReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetEmbeddableReportHtmlUrl() string {
	if o == nil || IsNil(o.EmbeddableReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.EmbeddableReportHtmlUrl
}

// GetEmbeddableReportHtmlUrlOk returns a tuple with the EmbeddableReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetEmbeddableReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EmbeddableReportHtmlUrl) {
		return nil, false
	}
	return o.EmbeddableReportHtmlUrl, true
}

// HasEmbeddableReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasEmbeddableReportHtmlUrl() bool {
	if o != nil && !IsNil(o.EmbeddableReportHtmlUrl) {
		return true
	}

	return false
}

// SetEmbeddableReportHtmlUrl gets a reference to the given string and assigns it to the EmbeddableReportHtmlUrl field.
func (o *ApiThirdPartyScanResultDTO) SetEmbeddableReportHtmlUrl(v string) {
	o.EmbeddableReportHtmlUrl = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ApiThirdPartyScanResultDTO) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetGrandfatheredPolicyViolations returns the GrandfatheredPolicyViolations field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetGrandfatheredPolicyViolations() int32 {
	if o == nil || IsNil(o.GrandfatheredPolicyViolations) {
		var ret int32
		return ret
	}
	return *o.GrandfatheredPolicyViolations
}

// GetGrandfatheredPolicyViolationsOk returns a tuple with the GrandfatheredPolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetGrandfatheredPolicyViolationsOk() (*int32, bool) {
	if o == nil || IsNil(o.GrandfatheredPolicyViolations) {
		return nil, false
	}
	return o.GrandfatheredPolicyViolations, true
}

// HasGrandfatheredPolicyViolations returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasGrandfatheredPolicyViolations() bool {
	if o != nil && !IsNil(o.GrandfatheredPolicyViolations) {
		return true
	}

	return false
}

// SetGrandfatheredPolicyViolations gets a reference to the given int32 and assigns it to the GrandfatheredPolicyViolations field.
func (o *ApiThirdPartyScanResultDTO) SetGrandfatheredPolicyViolations(v int32) {
	o.GrandfatheredPolicyViolations = &v
}

// GetIsError returns the IsError field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetIsError() bool {
	if o == nil || IsNil(o.IsError) {
		var ret bool
		return ret
	}
	return *o.IsError
}

// GetIsErrorOk returns a tuple with the IsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetIsErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsError) {
		return nil, false
	}
	return o.IsError, true
}

// HasIsError returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasIsError() bool {
	if o != nil && !IsNil(o.IsError) {
		return true
	}

	return false
}

// SetIsError gets a reference to the given bool and assigns it to the IsError field.
func (o *ApiThirdPartyScanResultDTO) SetIsError(v bool) {
	o.IsError = &v
}

// GetLegacyViolations returns the LegacyViolations field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetLegacyViolations() int32 {
	if o == nil || IsNil(o.LegacyViolations) {
		var ret int32
		return ret
	}
	return *o.LegacyViolations
}

// GetLegacyViolationsOk returns a tuple with the LegacyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetLegacyViolationsOk() (*int32, bool) {
	if o == nil || IsNil(o.LegacyViolations) {
		return nil, false
	}
	return o.LegacyViolations, true
}

// HasLegacyViolations returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasLegacyViolations() bool {
	if o != nil && !IsNil(o.LegacyViolations) {
		return true
	}

	return false
}

// SetLegacyViolations gets a reference to the given int32 and assigns it to the LegacyViolations field.
func (o *ApiThirdPartyScanResultDTO) SetLegacyViolations(v int32) {
	o.LegacyViolations = &v
}

// GetOpenPolicyViolations returns the OpenPolicyViolations field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetOpenPolicyViolations() ApiEvaluationResultCounterDTO {
	if o == nil || IsNil(o.OpenPolicyViolations) {
		var ret ApiEvaluationResultCounterDTO
		return ret
	}
	return *o.OpenPolicyViolations
}

// GetOpenPolicyViolationsOk returns a tuple with the OpenPolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetOpenPolicyViolationsOk() (*ApiEvaluationResultCounterDTO, bool) {
	if o == nil || IsNil(o.OpenPolicyViolations) {
		return nil, false
	}
	return o.OpenPolicyViolations, true
}

// HasOpenPolicyViolations returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasOpenPolicyViolations() bool {
	if o != nil && !IsNil(o.OpenPolicyViolations) {
		return true
	}

	return false
}

// SetOpenPolicyViolations gets a reference to the given ApiEvaluationResultCounterDTO and assigns it to the OpenPolicyViolations field.
func (o *ApiThirdPartyScanResultDTO) SetOpenPolicyViolations(v ApiEvaluationResultCounterDTO) {
	o.OpenPolicyViolations = &v
}

// GetPolicyAction returns the PolicyAction field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetPolicyAction() string {
	if o == nil || IsNil(o.PolicyAction) {
		var ret string
		return ret
	}
	return *o.PolicyAction
}

// GetPolicyActionOk returns a tuple with the PolicyAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetPolicyActionOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyAction) {
		return nil, false
	}
	return o.PolicyAction, true
}

// HasPolicyAction returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasPolicyAction() bool {
	if o != nil && !IsNil(o.PolicyAction) {
		return true
	}

	return false
}

// SetPolicyAction gets a reference to the given string and assigns it to the PolicyAction field.
func (o *ApiThirdPartyScanResultDTO) SetPolicyAction(v string) {
	o.PolicyAction = &v
}

// GetReportDataUrl returns the ReportDataUrl field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetReportDataUrl() string {
	if o == nil || IsNil(o.ReportDataUrl) {
		var ret string
		return ret
	}
	return *o.ReportDataUrl
}

// GetReportDataUrlOk returns a tuple with the ReportDataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetReportDataUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportDataUrl) {
		return nil, false
	}
	return o.ReportDataUrl, true
}

// HasReportDataUrl returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasReportDataUrl() bool {
	if o != nil && !IsNil(o.ReportDataUrl) {
		return true
	}

	return false
}

// SetReportDataUrl gets a reference to the given string and assigns it to the ReportDataUrl field.
func (o *ApiThirdPartyScanResultDTO) SetReportDataUrl(v string) {
	o.ReportDataUrl = &v
}

// GetReportHtmlUrl returns the ReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetReportHtmlUrl() string {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.ReportHtmlUrl
}

// GetReportHtmlUrlOk returns a tuple with the ReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		return nil, false
	}
	return o.ReportHtmlUrl, true
}

// HasReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasReportHtmlUrl() bool {
	if o != nil && !IsNil(o.ReportHtmlUrl) {
		return true
	}

	return false
}

// SetReportHtmlUrl gets a reference to the given string and assigns it to the ReportHtmlUrl field.
func (o *ApiThirdPartyScanResultDTO) SetReportHtmlUrl(v string) {
	o.ReportHtmlUrl = &v
}

// GetReportPdfUrl returns the ReportPdfUrl field value if set, zero value otherwise.
func (o *ApiThirdPartyScanResultDTO) GetReportPdfUrl() string {
	if o == nil || IsNil(o.ReportPdfUrl) {
		var ret string
		return ret
	}
	return *o.ReportPdfUrl
}

// GetReportPdfUrlOk returns a tuple with the ReportPdfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanResultDTO) GetReportPdfUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportPdfUrl) {
		return nil, false
	}
	return o.ReportPdfUrl, true
}

// HasReportPdfUrl returns a boolean if a field has been set.
func (o *ApiThirdPartyScanResultDTO) HasReportPdfUrl() bool {
	if o != nil && !IsNil(o.ReportPdfUrl) {
		return true
	}

	return false
}

// SetReportPdfUrl gets a reference to the given string and assigns it to the ReportPdfUrl field.
func (o *ApiThirdPartyScanResultDTO) SetReportPdfUrl(v string) {
	o.ReportPdfUrl = &v
}

func (o ApiThirdPartyScanResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiThirdPartyScanResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentsAffected) {
		toSerialize["componentsAffected"] = o.ComponentsAffected
	}
	if !IsNil(o.EmbeddableReportHtmlUrl) {
		toSerialize["embeddableReportHtmlUrl"] = o.EmbeddableReportHtmlUrl
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.GrandfatheredPolicyViolations) {
		toSerialize["grandfatheredPolicyViolations"] = o.GrandfatheredPolicyViolations
	}
	if !IsNil(o.IsError) {
		toSerialize["isError"] = o.IsError
	}
	if !IsNil(o.LegacyViolations) {
		toSerialize["legacyViolations"] = o.LegacyViolations
	}
	if !IsNil(o.OpenPolicyViolations) {
		toSerialize["openPolicyViolations"] = o.OpenPolicyViolations
	}
	if !IsNil(o.PolicyAction) {
		toSerialize["policyAction"] = o.PolicyAction
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
	return toSerialize, nil
}

type NullableApiThirdPartyScanResultDTO struct {
	value *ApiThirdPartyScanResultDTO
	isSet bool
}

func (v NullableApiThirdPartyScanResultDTO) Get() *ApiThirdPartyScanResultDTO {
	return v.value
}

func (v *NullableApiThirdPartyScanResultDTO) Set(val *ApiThirdPartyScanResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiThirdPartyScanResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiThirdPartyScanResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiThirdPartyScanResultDTO(val *ApiThirdPartyScanResultDTO) *NullableApiThirdPartyScanResultDTO {
	return &NullableApiThirdPartyScanResultDTO{value: val, isSet: true}
}

func (v NullableApiThirdPartyScanResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiThirdPartyScanResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


