/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 165
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiComponentEvaluationResultDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentEvaluationResultDTOV2{}

// ApiComponentEvaluationResultDTOV2 struct for ApiComponentEvaluationResultDTOV2
type ApiComponentEvaluationResultDTOV2 struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	EvaluationDate *time.Time `json:"evaluationDate,omitempty"`
	IsError *bool `json:"isError,omitempty"`
	Results []ApiComponentDetailsDTOV2 `json:"results,omitempty"`
	SubmittedDate *time.Time `json:"submittedDate,omitempty"`
}

// NewApiComponentEvaluationResultDTOV2 instantiates a new ApiComponentEvaluationResultDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentEvaluationResultDTOV2() *ApiComponentEvaluationResultDTOV2 {
	this := ApiComponentEvaluationResultDTOV2{}
	return &this
}

// NewApiComponentEvaluationResultDTOV2WithDefaults instantiates a new ApiComponentEvaluationResultDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentEvaluationResultDTOV2WithDefaults() *ApiComponentEvaluationResultDTOV2 {
	this := ApiComponentEvaluationResultDTOV2{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApiComponentEvaluationResultDTOV2) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentEvaluationResultDTOV2) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApiComponentEvaluationResultDTOV2) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApiComponentEvaluationResultDTOV2) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ApiComponentEvaluationResultDTOV2) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentEvaluationResultDTOV2) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ApiComponentEvaluationResultDTOV2) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ApiComponentEvaluationResultDTOV2) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetEvaluationDate returns the EvaluationDate field value if set, zero value otherwise.
func (o *ApiComponentEvaluationResultDTOV2) GetEvaluationDate() time.Time {
	if o == nil || IsNil(o.EvaluationDate) {
		var ret time.Time
		return ret
	}
	return *o.EvaluationDate
}

// GetEvaluationDateOk returns a tuple with the EvaluationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentEvaluationResultDTOV2) GetEvaluationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EvaluationDate) {
		return nil, false
	}
	return o.EvaluationDate, true
}

// HasEvaluationDate returns a boolean if a field has been set.
func (o *ApiComponentEvaluationResultDTOV2) HasEvaluationDate() bool {
	if o != nil && !IsNil(o.EvaluationDate) {
		return true
	}

	return false
}

// SetEvaluationDate gets a reference to the given time.Time and assigns it to the EvaluationDate field.
func (o *ApiComponentEvaluationResultDTOV2) SetEvaluationDate(v time.Time) {
	o.EvaluationDate = &v
}

// GetIsError returns the IsError field value if set, zero value otherwise.
func (o *ApiComponentEvaluationResultDTOV2) GetIsError() bool {
	if o == nil || IsNil(o.IsError) {
		var ret bool
		return ret
	}
	return *o.IsError
}

// GetIsErrorOk returns a tuple with the IsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentEvaluationResultDTOV2) GetIsErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsError) {
		return nil, false
	}
	return o.IsError, true
}

// HasIsError returns a boolean if a field has been set.
func (o *ApiComponentEvaluationResultDTOV2) HasIsError() bool {
	if o != nil && !IsNil(o.IsError) {
		return true
	}

	return false
}

// SetIsError gets a reference to the given bool and assigns it to the IsError field.
func (o *ApiComponentEvaluationResultDTOV2) SetIsError(v bool) {
	o.IsError = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ApiComponentEvaluationResultDTOV2) GetResults() []ApiComponentDetailsDTOV2 {
	if o == nil || IsNil(o.Results) {
		var ret []ApiComponentDetailsDTOV2
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentEvaluationResultDTOV2) GetResultsOk() ([]ApiComponentDetailsDTOV2, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ApiComponentEvaluationResultDTOV2) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ApiComponentDetailsDTOV2 and assigns it to the Results field.
func (o *ApiComponentEvaluationResultDTOV2) SetResults(v []ApiComponentDetailsDTOV2) {
	o.Results = v
}

// GetSubmittedDate returns the SubmittedDate field value if set, zero value otherwise.
func (o *ApiComponentEvaluationResultDTOV2) GetSubmittedDate() time.Time {
	if o == nil || IsNil(o.SubmittedDate) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedDate
}

// GetSubmittedDateOk returns a tuple with the SubmittedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentEvaluationResultDTOV2) GetSubmittedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedDate) {
		return nil, false
	}
	return o.SubmittedDate, true
}

// HasSubmittedDate returns a boolean if a field has been set.
func (o *ApiComponentEvaluationResultDTOV2) HasSubmittedDate() bool {
	if o != nil && !IsNil(o.SubmittedDate) {
		return true
	}

	return false
}

// SetSubmittedDate gets a reference to the given time.Time and assigns it to the SubmittedDate field.
func (o *ApiComponentEvaluationResultDTOV2) SetSubmittedDate(v time.Time) {
	o.SubmittedDate = &v
}

func (o ApiComponentEvaluationResultDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentEvaluationResultDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.EvaluationDate) {
		toSerialize["evaluationDate"] = o.EvaluationDate
	}
	if !IsNil(o.IsError) {
		toSerialize["isError"] = o.IsError
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.SubmittedDate) {
		toSerialize["submittedDate"] = o.SubmittedDate
	}
	return toSerialize, nil
}

type NullableApiComponentEvaluationResultDTOV2 struct {
	value *ApiComponentEvaluationResultDTOV2
	isSet bool
}

func (v NullableApiComponentEvaluationResultDTOV2) Get() *ApiComponentEvaluationResultDTOV2 {
	return v.value
}

func (v *NullableApiComponentEvaluationResultDTOV2) Set(val *ApiComponentEvaluationResultDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentEvaluationResultDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentEvaluationResultDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentEvaluationResultDTOV2(val *ApiComponentEvaluationResultDTOV2) *NullableApiComponentEvaluationResultDTOV2 {
	return &NullableApiComponentEvaluationResultDTOV2{value: val, isSet: true}
}

func (v NullableApiComponentEvaluationResultDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentEvaluationResultDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


