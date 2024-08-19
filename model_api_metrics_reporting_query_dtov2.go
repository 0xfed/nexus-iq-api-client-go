/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.180.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiMetricsReportingQueryDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMetricsReportingQueryDTOV2{}

// ApiMetricsReportingQueryDTOV2 struct for ApiMetricsReportingQueryDTOV2
type ApiMetricsReportingQueryDTOV2 struct {
	ApplicationIds []string `json:"applicationIds,omitempty"`
	FirstTimePeriod *string `json:"firstTimePeriod,omitempty"`
	LastTimePeriod *string `json:"lastTimePeriod,omitempty"`
	OrganizationIds []string `json:"organizationIds,omitempty"`
	TimePeriod *string `json:"timePeriod,omitempty"`
}

// NewApiMetricsReportingQueryDTOV2 instantiates a new ApiMetricsReportingQueryDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMetricsReportingQueryDTOV2() *ApiMetricsReportingQueryDTOV2 {
	this := ApiMetricsReportingQueryDTOV2{}
	return &this
}

// NewApiMetricsReportingQueryDTOV2WithDefaults instantiates a new ApiMetricsReportingQueryDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMetricsReportingQueryDTOV2WithDefaults() *ApiMetricsReportingQueryDTOV2 {
	this := ApiMetricsReportingQueryDTOV2{}
	return &this
}

// GetApplicationIds returns the ApplicationIds field value if set, zero value otherwise.
func (o *ApiMetricsReportingQueryDTOV2) GetApplicationIds() []string {
	if o == nil || IsNil(o.ApplicationIds) {
		var ret []string
		return ret
	}
	return o.ApplicationIds
}

// GetApplicationIdsOk returns a tuple with the ApplicationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricsReportingQueryDTOV2) GetApplicationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ApplicationIds) {
		return nil, false
	}
	return o.ApplicationIds, true
}

// HasApplicationIds returns a boolean if a field has been set.
func (o *ApiMetricsReportingQueryDTOV2) HasApplicationIds() bool {
	if o != nil && !IsNil(o.ApplicationIds) {
		return true
	}

	return false
}

// SetApplicationIds gets a reference to the given []string and assigns it to the ApplicationIds field.
func (o *ApiMetricsReportingQueryDTOV2) SetApplicationIds(v []string) {
	o.ApplicationIds = v
}

// GetFirstTimePeriod returns the FirstTimePeriod field value if set, zero value otherwise.
func (o *ApiMetricsReportingQueryDTOV2) GetFirstTimePeriod() string {
	if o == nil || IsNil(o.FirstTimePeriod) {
		var ret string
		return ret
	}
	return *o.FirstTimePeriod
}

// GetFirstTimePeriodOk returns a tuple with the FirstTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricsReportingQueryDTOV2) GetFirstTimePeriodOk() (*string, bool) {
	if o == nil || IsNil(o.FirstTimePeriod) {
		return nil, false
	}
	return o.FirstTimePeriod, true
}

// HasFirstTimePeriod returns a boolean if a field has been set.
func (o *ApiMetricsReportingQueryDTOV2) HasFirstTimePeriod() bool {
	if o != nil && !IsNil(o.FirstTimePeriod) {
		return true
	}

	return false
}

// SetFirstTimePeriod gets a reference to the given string and assigns it to the FirstTimePeriod field.
func (o *ApiMetricsReportingQueryDTOV2) SetFirstTimePeriod(v string) {
	o.FirstTimePeriod = &v
}

// GetLastTimePeriod returns the LastTimePeriod field value if set, zero value otherwise.
func (o *ApiMetricsReportingQueryDTOV2) GetLastTimePeriod() string {
	if o == nil || IsNil(o.LastTimePeriod) {
		var ret string
		return ret
	}
	return *o.LastTimePeriod
}

// GetLastTimePeriodOk returns a tuple with the LastTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricsReportingQueryDTOV2) GetLastTimePeriodOk() (*string, bool) {
	if o == nil || IsNil(o.LastTimePeriod) {
		return nil, false
	}
	return o.LastTimePeriod, true
}

// HasLastTimePeriod returns a boolean if a field has been set.
func (o *ApiMetricsReportingQueryDTOV2) HasLastTimePeriod() bool {
	if o != nil && !IsNil(o.LastTimePeriod) {
		return true
	}

	return false
}

// SetLastTimePeriod gets a reference to the given string and assigns it to the LastTimePeriod field.
func (o *ApiMetricsReportingQueryDTOV2) SetLastTimePeriod(v string) {
	o.LastTimePeriod = &v
}

// GetOrganizationIds returns the OrganizationIds field value if set, zero value otherwise.
func (o *ApiMetricsReportingQueryDTOV2) GetOrganizationIds() []string {
	if o == nil || IsNil(o.OrganizationIds) {
		var ret []string
		return ret
	}
	return o.OrganizationIds
}

// GetOrganizationIdsOk returns a tuple with the OrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricsReportingQueryDTOV2) GetOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationIds) {
		return nil, false
	}
	return o.OrganizationIds, true
}

// HasOrganizationIds returns a boolean if a field has been set.
func (o *ApiMetricsReportingQueryDTOV2) HasOrganizationIds() bool {
	if o != nil && !IsNil(o.OrganizationIds) {
		return true
	}

	return false
}

// SetOrganizationIds gets a reference to the given []string and assigns it to the OrganizationIds field.
func (o *ApiMetricsReportingQueryDTOV2) SetOrganizationIds(v []string) {
	o.OrganizationIds = v
}

// GetTimePeriod returns the TimePeriod field value if set, zero value otherwise.
func (o *ApiMetricsReportingQueryDTOV2) GetTimePeriod() string {
	if o == nil || IsNil(o.TimePeriod) {
		var ret string
		return ret
	}
	return *o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricsReportingQueryDTOV2) GetTimePeriodOk() (*string, bool) {
	if o == nil || IsNil(o.TimePeriod) {
		return nil, false
	}
	return o.TimePeriod, true
}

// HasTimePeriod returns a boolean if a field has been set.
func (o *ApiMetricsReportingQueryDTOV2) HasTimePeriod() bool {
	if o != nil && !IsNil(o.TimePeriod) {
		return true
	}

	return false
}

// SetTimePeriod gets a reference to the given string and assigns it to the TimePeriod field.
func (o *ApiMetricsReportingQueryDTOV2) SetTimePeriod(v string) {
	o.TimePeriod = &v
}

func (o ApiMetricsReportingQueryDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMetricsReportingQueryDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationIds) {
		toSerialize["applicationIds"] = o.ApplicationIds
	}
	if !IsNil(o.FirstTimePeriod) {
		toSerialize["firstTimePeriod"] = o.FirstTimePeriod
	}
	if !IsNil(o.LastTimePeriod) {
		toSerialize["lastTimePeriod"] = o.LastTimePeriod
	}
	if !IsNil(o.OrganizationIds) {
		toSerialize["organizationIds"] = o.OrganizationIds
	}
	if !IsNil(o.TimePeriod) {
		toSerialize["timePeriod"] = o.TimePeriod
	}
	return toSerialize, nil
}

type NullableApiMetricsReportingQueryDTOV2 struct {
	value *ApiMetricsReportingQueryDTOV2
	isSet bool
}

func (v NullableApiMetricsReportingQueryDTOV2) Get() *ApiMetricsReportingQueryDTOV2 {
	return v.value
}

func (v *NullableApiMetricsReportingQueryDTOV2) Set(val *ApiMetricsReportingQueryDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMetricsReportingQueryDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMetricsReportingQueryDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMetricsReportingQueryDTOV2(val *ApiMetricsReportingQueryDTOV2) *NullableApiMetricsReportingQueryDTOV2 {
	return &NullableApiMetricsReportingQueryDTOV2{value: val, isSet: true}
}

func (v NullableApiMetricsReportingQueryDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMetricsReportingQueryDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


