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

// checks if the ReferenceLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceLink{}

// ReferenceLink struct for ReferenceLink
type ReferenceLink struct {
	ReferenceType *string `json:"referenceType,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewReferenceLink instantiates a new ReferenceLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceLink() *ReferenceLink {
	this := ReferenceLink{}
	return &this
}

// NewReferenceLinkWithDefaults instantiates a new ReferenceLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceLinkWithDefaults() *ReferenceLink {
	this := ReferenceLink{}
	return &this
}

// GetReferenceType returns the ReferenceType field value if set, zero value otherwise.
func (o *ReferenceLink) GetReferenceType() string {
	if o == nil || IsNil(o.ReferenceType) {
		var ret string
		return ret
	}
	return *o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceLink) GetReferenceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceType) {
		return nil, false
	}
	return o.ReferenceType, true
}

// HasReferenceType returns a boolean if a field has been set.
func (o *ReferenceLink) HasReferenceType() bool {
	if o != nil && !IsNil(o.ReferenceType) {
		return true
	}

	return false
}

// SetReferenceType gets a reference to the given string and assigns it to the ReferenceType field.
func (o *ReferenceLink) SetReferenceType(v string) {
	o.ReferenceType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ReferenceLink) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceLink) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ReferenceLink) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ReferenceLink) SetUrl(v string) {
	o.Url = &v
}

func (o ReferenceLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceType) {
		toSerialize["referenceType"] = o.ReferenceType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableReferenceLink struct {
	value *ReferenceLink
	isSet bool
}

func (v NullableReferenceLink) Get() *ReferenceLink {
	return v.value
}

func (v *NullableReferenceLink) Set(val *ReferenceLink) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceLink) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceLink(val *ReferenceLink) *NullableReferenceLink {
	return &NullableReferenceLink{value: val, isSet: true}
}

func (v NullableReferenceLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


