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

// checks if the LabelsByOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelsByOwner{}

// LabelsByOwner struct for LabelsByOwner
type LabelsByOwner struct {
	Labels []ApiLabelDTO `json:"labels,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerName *string `json:"ownerName,omitempty"`
	OwnerType *string `json:"ownerType,omitempty"`
}

// NewLabelsByOwner instantiates a new LabelsByOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelsByOwner() *LabelsByOwner {
	this := LabelsByOwner{}
	return &this
}

// NewLabelsByOwnerWithDefaults instantiates a new LabelsByOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelsByOwnerWithDefaults() *LabelsByOwner {
	this := LabelsByOwner{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *LabelsByOwner) GetLabels() []ApiLabelDTO {
	if o == nil || IsNil(o.Labels) {
		var ret []ApiLabelDTO
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsByOwner) GetLabelsOk() ([]ApiLabelDTO, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LabelsByOwner) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []ApiLabelDTO and assigns it to the Labels field.
func (o *LabelsByOwner) SetLabels(v []ApiLabelDTO) {
	o.Labels = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *LabelsByOwner) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsByOwner) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *LabelsByOwner) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *LabelsByOwner) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *LabelsByOwner) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsByOwner) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *LabelsByOwner) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *LabelsByOwner) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *LabelsByOwner) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsByOwner) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *LabelsByOwner) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *LabelsByOwner) SetOwnerType(v string) {
	o.OwnerType = &v
}

func (o LabelsByOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelsByOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !IsNil(o.OwnerType) {
		toSerialize["ownerType"] = o.OwnerType
	}
	return toSerialize, nil
}

type NullableLabelsByOwner struct {
	value *LabelsByOwner
	isSet bool
}

func (v NullableLabelsByOwner) Get() *LabelsByOwner {
	return v.value
}

func (v *NullableLabelsByOwner) Set(val *LabelsByOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelsByOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelsByOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelsByOwner(val *LabelsByOwner) *NullableLabelsByOwner {
	return &NullableLabelsByOwner{value: val, isSet: true}
}

func (v NullableLabelsByOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelsByOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


