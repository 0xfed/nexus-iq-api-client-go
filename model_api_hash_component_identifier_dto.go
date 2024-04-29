/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.175.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiHashComponentIdentifierDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHashComponentIdentifierDTO{}

// ApiHashComponentIdentifierDTO struct for ApiHashComponentIdentifierDTO
type ApiHashComponentIdentifierDTO struct {
	Comment *string `json:"comment,omitempty"`
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	Hash *string `json:"hash,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
}

// NewApiHashComponentIdentifierDTO instantiates a new ApiHashComponentIdentifierDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHashComponentIdentifierDTO() *ApiHashComponentIdentifierDTO {
	this := ApiHashComponentIdentifierDTO{}
	return &this
}

// NewApiHashComponentIdentifierDTOWithDefaults instantiates a new ApiHashComponentIdentifierDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHashComponentIdentifierDTOWithDefaults() *ApiHashComponentIdentifierDTO {
	this := ApiHashComponentIdentifierDTO{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiHashComponentIdentifierDTO) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHashComponentIdentifierDTO) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiHashComponentIdentifierDTO) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiHashComponentIdentifierDTO) SetComment(v string) {
	o.Comment = &v
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiHashComponentIdentifierDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHashComponentIdentifierDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiHashComponentIdentifierDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiHashComponentIdentifierDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ApiHashComponentIdentifierDTO) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHashComponentIdentifierDTO) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ApiHashComponentIdentifierDTO) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *ApiHashComponentIdentifierDTO) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiHashComponentIdentifierDTO) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHashComponentIdentifierDTO) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiHashComponentIdentifierDTO) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiHashComponentIdentifierDTO) SetHash(v string) {
	o.Hash = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiHashComponentIdentifierDTO) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHashComponentIdentifierDTO) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiHashComponentIdentifierDTO) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiHashComponentIdentifierDTO) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

func (o ApiHashComponentIdentifierDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHashComponentIdentifierDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	return toSerialize, nil
}

type NullableApiHashComponentIdentifierDTO struct {
	value *ApiHashComponentIdentifierDTO
	isSet bool
}

func (v NullableApiHashComponentIdentifierDTO) Get() *ApiHashComponentIdentifierDTO {
	return v.value
}

func (v *NullableApiHashComponentIdentifierDTO) Set(val *ApiHashComponentIdentifierDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHashComponentIdentifierDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHashComponentIdentifierDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHashComponentIdentifierDTO(val *ApiHashComponentIdentifierDTO) *NullableApiHashComponentIdentifierDTO {
	return &NullableApiHashComponentIdentifierDTO{value: val, isSet: true}
}

func (v NullableApiHashComponentIdentifierDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHashComponentIdentifierDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


