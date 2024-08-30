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

// checks if the ApiApplicationTagDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiApplicationTagDTO{}

// ApiApplicationTagDTO struct for ApiApplicationTagDTO
type ApiApplicationTagDTO struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	Id *string `json:"id,omitempty"`
	TagId *string `json:"tagId,omitempty"`
}

// NewApiApplicationTagDTO instantiates a new ApiApplicationTagDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiApplicationTagDTO() *ApiApplicationTagDTO {
	this := ApiApplicationTagDTO{}
	return &this
}

// NewApiApplicationTagDTOWithDefaults instantiates a new ApiApplicationTagDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiApplicationTagDTOWithDefaults() *ApiApplicationTagDTO {
	this := ApiApplicationTagDTO{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApiApplicationTagDTO) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationTagDTO) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApiApplicationTagDTO) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApiApplicationTagDTO) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiApplicationTagDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationTagDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiApplicationTagDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiApplicationTagDTO) SetId(v string) {
	o.Id = &v
}

// GetTagId returns the TagId field value if set, zero value otherwise.
func (o *ApiApplicationTagDTO) GetTagId() string {
	if o == nil || IsNil(o.TagId) {
		var ret string
		return ret
	}
	return *o.TagId
}

// GetTagIdOk returns a tuple with the TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationTagDTO) GetTagIdOk() (*string, bool) {
	if o == nil || IsNil(o.TagId) {
		return nil, false
	}
	return o.TagId, true
}

// HasTagId returns a boolean if a field has been set.
func (o *ApiApplicationTagDTO) HasTagId() bool {
	if o != nil && !IsNil(o.TagId) {
		return true
	}

	return false
}

// SetTagId gets a reference to the given string and assigns it to the TagId field.
func (o *ApiApplicationTagDTO) SetTagId(v string) {
	o.TagId = &v
}

func (o ApiApplicationTagDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiApplicationTagDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TagId) {
		toSerialize["tagId"] = o.TagId
	}
	return toSerialize, nil
}

type NullableApiApplicationTagDTO struct {
	value *ApiApplicationTagDTO
	isSet bool
}

func (v NullableApiApplicationTagDTO) Get() *ApiApplicationTagDTO {
	return v.value
}

func (v *NullableApiApplicationTagDTO) Set(val *ApiApplicationTagDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiApplicationTagDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiApplicationTagDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiApplicationTagDTO(val *ApiApplicationTagDTO) *NullableApiApplicationTagDTO {
	return &NullableApiApplicationTagDTO{value: val, isSet: true}
}

func (v NullableApiApplicationTagDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiApplicationTagDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


