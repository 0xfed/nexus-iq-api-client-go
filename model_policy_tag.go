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

// checks if the PolicyTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyTag{}

// PolicyTag struct for PolicyTag
type PolicyTag struct {
	Id *string `json:"id,omitempty"`
	PolicyId *string `json:"policyId,omitempty"`
	TagId *string `json:"tagId,omitempty"`
}

// NewPolicyTag instantiates a new PolicyTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyTag() *PolicyTag {
	this := PolicyTag{}
	return &this
}

// NewPolicyTagWithDefaults instantiates a new PolicyTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyTagWithDefaults() *PolicyTag {
	this := PolicyTag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PolicyTag) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTag) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PolicyTag) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PolicyTag) SetId(v string) {
	o.Id = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *PolicyTag) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTag) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *PolicyTag) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *PolicyTag) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetTagId returns the TagId field value if set, zero value otherwise.
func (o *PolicyTag) GetTagId() string {
	if o == nil || IsNil(o.TagId) {
		var ret string
		return ret
	}
	return *o.TagId
}

// GetTagIdOk returns a tuple with the TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTag) GetTagIdOk() (*string, bool) {
	if o == nil || IsNil(o.TagId) {
		return nil, false
	}
	return o.TagId, true
}

// HasTagId returns a boolean if a field has been set.
func (o *PolicyTag) HasTagId() bool {
	if o != nil && !IsNil(o.TagId) {
		return true
	}

	return false
}

// SetTagId gets a reference to the given string and assigns it to the TagId field.
func (o *PolicyTag) SetTagId(v string) {
	o.TagId = &v
}

func (o PolicyTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if !IsNil(o.TagId) {
		toSerialize["tagId"] = o.TagId
	}
	return toSerialize, nil
}

type NullablePolicyTag struct {
	value *PolicyTag
	isSet bool
}

func (v NullablePolicyTag) Get() *PolicyTag {
	return v.value
}

func (v *NullablePolicyTag) Set(val *PolicyTag) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyTag) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyTag(val *PolicyTag) *NullablePolicyTag {
	return &NullablePolicyTag{value: val, isSet: true}
}

func (v NullablePolicyTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


