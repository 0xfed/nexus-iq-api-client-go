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

// checks if the ApiConstraintViolationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiConstraintViolationDTO{}

// ApiConstraintViolationDTO struct for ApiConstraintViolationDTO
type ApiConstraintViolationDTO struct {
	ConstraintId *string `json:"constraintId,omitempty"`
	ConstraintName *string `json:"constraintName,omitempty"`
	Reasons []ApiConstraintViolationReasonDTO `json:"reasons,omitempty"`
}

// NewApiConstraintViolationDTO instantiates a new ApiConstraintViolationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiConstraintViolationDTO() *ApiConstraintViolationDTO {
	this := ApiConstraintViolationDTO{}
	return &this
}

// NewApiConstraintViolationDTOWithDefaults instantiates a new ApiConstraintViolationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiConstraintViolationDTOWithDefaults() *ApiConstraintViolationDTO {
	this := ApiConstraintViolationDTO{}
	return &this
}

// GetConstraintId returns the ConstraintId field value if set, zero value otherwise.
func (o *ApiConstraintViolationDTO) GetConstraintId() string {
	if o == nil || IsNil(o.ConstraintId) {
		var ret string
		return ret
	}
	return *o.ConstraintId
}

// GetConstraintIdOk returns a tuple with the ConstraintId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConstraintViolationDTO) GetConstraintIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConstraintId) {
		return nil, false
	}
	return o.ConstraintId, true
}

// HasConstraintId returns a boolean if a field has been set.
func (o *ApiConstraintViolationDTO) HasConstraintId() bool {
	if o != nil && !IsNil(o.ConstraintId) {
		return true
	}

	return false
}

// SetConstraintId gets a reference to the given string and assigns it to the ConstraintId field.
func (o *ApiConstraintViolationDTO) SetConstraintId(v string) {
	o.ConstraintId = &v
}

// GetConstraintName returns the ConstraintName field value if set, zero value otherwise.
func (o *ApiConstraintViolationDTO) GetConstraintName() string {
	if o == nil || IsNil(o.ConstraintName) {
		var ret string
		return ret
	}
	return *o.ConstraintName
}

// GetConstraintNameOk returns a tuple with the ConstraintName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConstraintViolationDTO) GetConstraintNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConstraintName) {
		return nil, false
	}
	return o.ConstraintName, true
}

// HasConstraintName returns a boolean if a field has been set.
func (o *ApiConstraintViolationDTO) HasConstraintName() bool {
	if o != nil && !IsNil(o.ConstraintName) {
		return true
	}

	return false
}

// SetConstraintName gets a reference to the given string and assigns it to the ConstraintName field.
func (o *ApiConstraintViolationDTO) SetConstraintName(v string) {
	o.ConstraintName = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *ApiConstraintViolationDTO) GetReasons() []ApiConstraintViolationReasonDTO {
	if o == nil || IsNil(o.Reasons) {
		var ret []ApiConstraintViolationReasonDTO
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConstraintViolationDTO) GetReasonsOk() ([]ApiConstraintViolationReasonDTO, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *ApiConstraintViolationDTO) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []ApiConstraintViolationReasonDTO and assigns it to the Reasons field.
func (o *ApiConstraintViolationDTO) SetReasons(v []ApiConstraintViolationReasonDTO) {
	o.Reasons = v
}

func (o ApiConstraintViolationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiConstraintViolationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConstraintId) {
		toSerialize["constraintId"] = o.ConstraintId
	}
	if !IsNil(o.ConstraintName) {
		toSerialize["constraintName"] = o.ConstraintName
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

type NullableApiConstraintViolationDTO struct {
	value *ApiConstraintViolationDTO
	isSet bool
}

func (v NullableApiConstraintViolationDTO) Get() *ApiConstraintViolationDTO {
	return v.value
}

func (v *NullableApiConstraintViolationDTO) Set(val *ApiConstraintViolationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiConstraintViolationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiConstraintViolationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiConstraintViolationDTO(val *ApiConstraintViolationDTO) *NullableApiConstraintViolationDTO {
	return &NullableApiConstraintViolationDTO{value: val, isSet: true}
}

func (v NullableApiConstraintViolationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiConstraintViolationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


