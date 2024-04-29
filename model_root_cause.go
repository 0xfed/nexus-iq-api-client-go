/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.175.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the RootCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RootCause{}

// RootCause struct for RootCause
type RootCause struct {
	ListOfPaths []string `json:"listOfPaths,omitempty"`
	VersionRange *string `json:"versionRange,omitempty"`
}

// NewRootCause instantiates a new RootCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootCause() *RootCause {
	this := RootCause{}
	return &this
}

// NewRootCauseWithDefaults instantiates a new RootCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootCauseWithDefaults() *RootCause {
	this := RootCause{}
	return &this
}

// GetListOfPaths returns the ListOfPaths field value if set, zero value otherwise.
func (o *RootCause) GetListOfPaths() []string {
	if o == nil || IsNil(o.ListOfPaths) {
		var ret []string
		return ret
	}
	return o.ListOfPaths
}

// GetListOfPathsOk returns a tuple with the ListOfPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootCause) GetListOfPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.ListOfPaths) {
		return nil, false
	}
	return o.ListOfPaths, true
}

// HasListOfPaths returns a boolean if a field has been set.
func (o *RootCause) HasListOfPaths() bool {
	if o != nil && !IsNil(o.ListOfPaths) {
		return true
	}

	return false
}

// SetListOfPaths gets a reference to the given []string and assigns it to the ListOfPaths field.
func (o *RootCause) SetListOfPaths(v []string) {
	o.ListOfPaths = v
}

// GetVersionRange returns the VersionRange field value if set, zero value otherwise.
func (o *RootCause) GetVersionRange() string {
	if o == nil || IsNil(o.VersionRange) {
		var ret string
		return ret
	}
	return *o.VersionRange
}

// GetVersionRangeOk returns a tuple with the VersionRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootCause) GetVersionRangeOk() (*string, bool) {
	if o == nil || IsNil(o.VersionRange) {
		return nil, false
	}
	return o.VersionRange, true
}

// HasVersionRange returns a boolean if a field has been set.
func (o *RootCause) HasVersionRange() bool {
	if o != nil && !IsNil(o.VersionRange) {
		return true
	}

	return false
}

// SetVersionRange gets a reference to the given string and assigns it to the VersionRange field.
func (o *RootCause) SetVersionRange(v string) {
	o.VersionRange = &v
}

func (o RootCause) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RootCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListOfPaths) {
		toSerialize["listOfPaths"] = o.ListOfPaths
	}
	if !IsNil(o.VersionRange) {
		toSerialize["versionRange"] = o.VersionRange
	}
	return toSerialize, nil
}

type NullableRootCause struct {
	value *RootCause
	isSet bool
}

func (v NullableRootCause) Get() *RootCause {
	return v.value
}

func (v *NullableRootCause) Set(val *RootCause) {
	v.value = val
	v.isSet = true
}

func (v NullableRootCause) IsSet() bool {
	return v.isSet
}

func (v *NullableRootCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootCause(val *RootCause) *NullableRootCause {
	return &NullableRootCause{value: val, isSet: true}
}

func (v NullableRootCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


