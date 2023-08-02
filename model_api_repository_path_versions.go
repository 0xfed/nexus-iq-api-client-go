/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiRepositoryPathVersions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryPathVersions{}

// ApiRepositoryPathVersions struct for ApiRepositoryPathVersions
type ApiRepositoryPathVersions struct {
	RepositoryComponentPaths []ApiRepositoryComponentPath `json:"repositoryComponentPaths,omitempty"`
	RequestIndex *int32 `json:"requestIndex,omitempty"`
}

// NewApiRepositoryPathVersions instantiates a new ApiRepositoryPathVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryPathVersions() *ApiRepositoryPathVersions {
	this := ApiRepositoryPathVersions{}
	return &this
}

// NewApiRepositoryPathVersionsWithDefaults instantiates a new ApiRepositoryPathVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryPathVersionsWithDefaults() *ApiRepositoryPathVersions {
	this := ApiRepositoryPathVersions{}
	return &this
}

// GetRepositoryComponentPaths returns the RepositoryComponentPaths field value if set, zero value otherwise.
func (o *ApiRepositoryPathVersions) GetRepositoryComponentPaths() []ApiRepositoryComponentPath {
	if o == nil || IsNil(o.RepositoryComponentPaths) {
		var ret []ApiRepositoryComponentPath
		return ret
	}
	return o.RepositoryComponentPaths
}

// GetRepositoryComponentPathsOk returns a tuple with the RepositoryComponentPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryPathVersions) GetRepositoryComponentPathsOk() ([]ApiRepositoryComponentPath, bool) {
	if o == nil || IsNil(o.RepositoryComponentPaths) {
		return nil, false
	}
	return o.RepositoryComponentPaths, true
}

// HasRepositoryComponentPaths returns a boolean if a field has been set.
func (o *ApiRepositoryPathVersions) HasRepositoryComponentPaths() bool {
	if o != nil && !IsNil(o.RepositoryComponentPaths) {
		return true
	}

	return false
}

// SetRepositoryComponentPaths gets a reference to the given []ApiRepositoryComponentPath and assigns it to the RepositoryComponentPaths field.
func (o *ApiRepositoryPathVersions) SetRepositoryComponentPaths(v []ApiRepositoryComponentPath) {
	o.RepositoryComponentPaths = v
}

// GetRequestIndex returns the RequestIndex field value if set, zero value otherwise.
func (o *ApiRepositoryPathVersions) GetRequestIndex() int32 {
	if o == nil || IsNil(o.RequestIndex) {
		var ret int32
		return ret
	}
	return *o.RequestIndex
}

// GetRequestIndexOk returns a tuple with the RequestIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryPathVersions) GetRequestIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.RequestIndex) {
		return nil, false
	}
	return o.RequestIndex, true
}

// HasRequestIndex returns a boolean if a field has been set.
func (o *ApiRepositoryPathVersions) HasRequestIndex() bool {
	if o != nil && !IsNil(o.RequestIndex) {
		return true
	}

	return false
}

// SetRequestIndex gets a reference to the given int32 and assigns it to the RequestIndex field.
func (o *ApiRepositoryPathVersions) SetRequestIndex(v int32) {
	o.RequestIndex = &v
}

func (o ApiRepositoryPathVersions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryPathVersions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepositoryComponentPaths) {
		toSerialize["repositoryComponentPaths"] = o.RepositoryComponentPaths
	}
	if !IsNil(o.RequestIndex) {
		toSerialize["requestIndex"] = o.RequestIndex
	}
	return toSerialize, nil
}

type NullableApiRepositoryPathVersions struct {
	value *ApiRepositoryPathVersions
	isSet bool
}

func (v NullableApiRepositoryPathVersions) Get() *ApiRepositoryPathVersions {
	return v.value
}

func (v *NullableApiRepositoryPathVersions) Set(val *ApiRepositoryPathVersions) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryPathVersions) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryPathVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryPathVersions(val *ApiRepositoryPathVersions) *NullableApiRepositoryPathVersions {
	return &NullableApiRepositoryPathVersions{value: val, isSet: true}
}

func (v NullableApiRepositoryPathVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryPathVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


