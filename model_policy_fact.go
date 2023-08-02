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

// checks if the PolicyFact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyFact{}

// PolicyFact struct for PolicyFact
type PolicyFact struct {
	ComponentFacts []ComponentFact `json:"componentFacts,omitempty"`
	PolicyId *string `json:"policyId,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	PolicyViolationId *string `json:"policyViolationId,omitempty"`
	ThreatLevel *int32 `json:"threatLevel,omitempty"`
}

// NewPolicyFact instantiates a new PolicyFact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyFact() *PolicyFact {
	this := PolicyFact{}
	return &this
}

// NewPolicyFactWithDefaults instantiates a new PolicyFact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyFactWithDefaults() *PolicyFact {
	this := PolicyFact{}
	return &this
}

// GetComponentFacts returns the ComponentFacts field value if set, zero value otherwise.
func (o *PolicyFact) GetComponentFacts() []ComponentFact {
	if o == nil || IsNil(o.ComponentFacts) {
		var ret []ComponentFact
		return ret
	}
	return o.ComponentFacts
}

// GetComponentFactsOk returns a tuple with the ComponentFacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFact) GetComponentFactsOk() ([]ComponentFact, bool) {
	if o == nil || IsNil(o.ComponentFacts) {
		return nil, false
	}
	return o.ComponentFacts, true
}

// HasComponentFacts returns a boolean if a field has been set.
func (o *PolicyFact) HasComponentFacts() bool {
	if o != nil && !IsNil(o.ComponentFacts) {
		return true
	}

	return false
}

// SetComponentFacts gets a reference to the given []ComponentFact and assigns it to the ComponentFacts field.
func (o *PolicyFact) SetComponentFacts(v []ComponentFact) {
	o.ComponentFacts = v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *PolicyFact) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFact) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *PolicyFact) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *PolicyFact) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *PolicyFact) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFact) GetPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *PolicyFact) HasPolicyName() bool {
	if o != nil && !IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *PolicyFact) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetPolicyViolationId returns the PolicyViolationId field value if set, zero value otherwise.
func (o *PolicyFact) GetPolicyViolationId() string {
	if o == nil || IsNil(o.PolicyViolationId) {
		var ret string
		return ret
	}
	return *o.PolicyViolationId
}

// GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFact) GetPolicyViolationIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyViolationId) {
		return nil, false
	}
	return o.PolicyViolationId, true
}

// HasPolicyViolationId returns a boolean if a field has been set.
func (o *PolicyFact) HasPolicyViolationId() bool {
	if o != nil && !IsNil(o.PolicyViolationId) {
		return true
	}

	return false
}

// SetPolicyViolationId gets a reference to the given string and assigns it to the PolicyViolationId field.
func (o *PolicyFact) SetPolicyViolationId(v string) {
	o.PolicyViolationId = &v
}

// GetThreatLevel returns the ThreatLevel field value if set, zero value otherwise.
func (o *PolicyFact) GetThreatLevel() int32 {
	if o == nil || IsNil(o.ThreatLevel) {
		var ret int32
		return ret
	}
	return *o.ThreatLevel
}

// GetThreatLevelOk returns a tuple with the ThreatLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFact) GetThreatLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ThreatLevel) {
		return nil, false
	}
	return o.ThreatLevel, true
}

// HasThreatLevel returns a boolean if a field has been set.
func (o *PolicyFact) HasThreatLevel() bool {
	if o != nil && !IsNil(o.ThreatLevel) {
		return true
	}

	return false
}

// SetThreatLevel gets a reference to the given int32 and assigns it to the ThreatLevel field.
func (o *PolicyFact) SetThreatLevel(v int32) {
	o.ThreatLevel = &v
}

func (o PolicyFact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyFact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentFacts) {
		toSerialize["componentFacts"] = o.ComponentFacts
	}
	if !IsNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if !IsNil(o.PolicyName) {
		toSerialize["policyName"] = o.PolicyName
	}
	if !IsNil(o.PolicyViolationId) {
		toSerialize["policyViolationId"] = o.PolicyViolationId
	}
	if !IsNil(o.ThreatLevel) {
		toSerialize["threatLevel"] = o.ThreatLevel
	}
	return toSerialize, nil
}

type NullablePolicyFact struct {
	value *PolicyFact
	isSet bool
}

func (v NullablePolicyFact) Get() *PolicyFact {
	return v.value
}

func (v *NullablePolicyFact) Set(val *PolicyFact) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyFact) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyFact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyFact(val *PolicyFact) *NullablePolicyFact {
	return &NullablePolicyFact{value: val, isSet: true}
}

func (v NullablePolicyFact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyFact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


