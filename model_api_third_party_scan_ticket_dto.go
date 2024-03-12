/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.174.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiThirdPartyScanTicketDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiThirdPartyScanTicketDTO{}

// ApiThirdPartyScanTicketDTO struct for ApiThirdPartyScanTicketDTO
type ApiThirdPartyScanTicketDTO struct {
	StatusUrl *string `json:"statusUrl,omitempty"`
}

// NewApiThirdPartyScanTicketDTO instantiates a new ApiThirdPartyScanTicketDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiThirdPartyScanTicketDTO() *ApiThirdPartyScanTicketDTO {
	this := ApiThirdPartyScanTicketDTO{}
	return &this
}

// NewApiThirdPartyScanTicketDTOWithDefaults instantiates a new ApiThirdPartyScanTicketDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiThirdPartyScanTicketDTOWithDefaults() *ApiThirdPartyScanTicketDTO {
	this := ApiThirdPartyScanTicketDTO{}
	return &this
}

// GetStatusUrl returns the StatusUrl field value if set, zero value otherwise.
func (o *ApiThirdPartyScanTicketDTO) GetStatusUrl() string {
	if o == nil || IsNil(o.StatusUrl) {
		var ret string
		return ret
	}
	return *o.StatusUrl
}

// GetStatusUrlOk returns a tuple with the StatusUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiThirdPartyScanTicketDTO) GetStatusUrlOk() (*string, bool) {
	if o == nil || IsNil(o.StatusUrl) {
		return nil, false
	}
	return o.StatusUrl, true
}

// HasStatusUrl returns a boolean if a field has been set.
func (o *ApiThirdPartyScanTicketDTO) HasStatusUrl() bool {
	if o != nil && !IsNil(o.StatusUrl) {
		return true
	}

	return false
}

// SetStatusUrl gets a reference to the given string and assigns it to the StatusUrl field.
func (o *ApiThirdPartyScanTicketDTO) SetStatusUrl(v string) {
	o.StatusUrl = &v
}

func (o ApiThirdPartyScanTicketDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiThirdPartyScanTicketDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusUrl) {
		toSerialize["statusUrl"] = o.StatusUrl
	}
	return toSerialize, nil
}

type NullableApiThirdPartyScanTicketDTO struct {
	value *ApiThirdPartyScanTicketDTO
	isSet bool
}

func (v NullableApiThirdPartyScanTicketDTO) Get() *ApiThirdPartyScanTicketDTO {
	return v.value
}

func (v *NullableApiThirdPartyScanTicketDTO) Set(val *ApiThirdPartyScanTicketDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiThirdPartyScanTicketDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiThirdPartyScanTicketDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiThirdPartyScanTicketDTO(val *ApiThirdPartyScanTicketDTO) *NullableApiThirdPartyScanTicketDTO {
	return &NullableApiThirdPartyScanTicketDTO{value: val, isSet: true}
}

func (v NullableApiThirdPartyScanTicketDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiThirdPartyScanTicketDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


