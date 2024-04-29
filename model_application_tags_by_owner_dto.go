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

// checks if the ApplicationTagsByOwnerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationTagsByOwnerDTO{}

// ApplicationTagsByOwnerDTO struct for ApplicationTagsByOwnerDTO
type ApplicationTagsByOwnerDTO struct {
	ApplicationTags []ApplicationTag `json:"applicationTags,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerName *string `json:"ownerName,omitempty"`
	OwnerType *string `json:"ownerType,omitempty"`
}

// NewApplicationTagsByOwnerDTO instantiates a new ApplicationTagsByOwnerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationTagsByOwnerDTO() *ApplicationTagsByOwnerDTO {
	this := ApplicationTagsByOwnerDTO{}
	return &this
}

// NewApplicationTagsByOwnerDTOWithDefaults instantiates a new ApplicationTagsByOwnerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationTagsByOwnerDTOWithDefaults() *ApplicationTagsByOwnerDTO {
	this := ApplicationTagsByOwnerDTO{}
	return &this
}

// GetApplicationTags returns the ApplicationTags field value if set, zero value otherwise.
func (o *ApplicationTagsByOwnerDTO) GetApplicationTags() []ApplicationTag {
	if o == nil || IsNil(o.ApplicationTags) {
		var ret []ApplicationTag
		return ret
	}
	return o.ApplicationTags
}

// GetApplicationTagsOk returns a tuple with the ApplicationTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTagsByOwnerDTO) GetApplicationTagsOk() ([]ApplicationTag, bool) {
	if o == nil || IsNil(o.ApplicationTags) {
		return nil, false
	}
	return o.ApplicationTags, true
}

// HasApplicationTags returns a boolean if a field has been set.
func (o *ApplicationTagsByOwnerDTO) HasApplicationTags() bool {
	if o != nil && !IsNil(o.ApplicationTags) {
		return true
	}

	return false
}

// SetApplicationTags gets a reference to the given []ApplicationTag and assigns it to the ApplicationTags field.
func (o *ApplicationTagsByOwnerDTO) SetApplicationTags(v []ApplicationTag) {
	o.ApplicationTags = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ApplicationTagsByOwnerDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTagsByOwnerDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ApplicationTagsByOwnerDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ApplicationTagsByOwnerDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *ApplicationTagsByOwnerDTO) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTagsByOwnerDTO) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *ApplicationTagsByOwnerDTO) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *ApplicationTagsByOwnerDTO) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *ApplicationTagsByOwnerDTO) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTagsByOwnerDTO) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *ApplicationTagsByOwnerDTO) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *ApplicationTagsByOwnerDTO) SetOwnerType(v string) {
	o.OwnerType = &v
}

func (o ApplicationTagsByOwnerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationTagsByOwnerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationTags) {
		toSerialize["applicationTags"] = o.ApplicationTags
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

type NullableApplicationTagsByOwnerDTO struct {
	value *ApplicationTagsByOwnerDTO
	isSet bool
}

func (v NullableApplicationTagsByOwnerDTO) Get() *ApplicationTagsByOwnerDTO {
	return v.value
}

func (v *NullableApplicationTagsByOwnerDTO) Set(val *ApplicationTagsByOwnerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationTagsByOwnerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationTagsByOwnerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationTagsByOwnerDTO(val *ApplicationTagsByOwnerDTO) *NullableApplicationTagsByOwnerDTO {
	return &NullableApplicationTagsByOwnerDTO{value: val, isSet: true}
}

func (v NullableApplicationTagsByOwnerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationTagsByOwnerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


