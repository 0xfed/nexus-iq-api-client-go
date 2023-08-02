/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ContentDisposition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentDisposition{}

// ContentDisposition struct for ContentDisposition
type ContentDisposition struct {
	CreationDate *time.Time `json:"creationDate,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	ModificationDate *time.Time `json:"modificationDate,omitempty"`
	Parameters *map[string]string `json:"parameters,omitempty"`
	ReadDate *time.Time `json:"readDate,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewContentDisposition instantiates a new ContentDisposition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentDisposition() *ContentDisposition {
	this := ContentDisposition{}
	return &this
}

// NewContentDispositionWithDefaults instantiates a new ContentDisposition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentDispositionWithDefaults() *ContentDisposition {
	this := ContentDisposition{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *ContentDisposition) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *ContentDisposition) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *ContentDisposition) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ContentDisposition) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ContentDisposition) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ContentDisposition) SetFileName(v string) {
	o.FileName = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *ContentDisposition) GetModificationDate() time.Time {
	if o == nil || IsNil(o.ModificationDate) {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModificationDate) {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *ContentDisposition) HasModificationDate() bool {
	if o != nil && !IsNil(o.ModificationDate) {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *ContentDisposition) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ContentDisposition) GetParameters() map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ContentDisposition) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *ContentDisposition) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetReadDate returns the ReadDate field value if set, zero value otherwise.
func (o *ContentDisposition) GetReadDate() time.Time {
	if o == nil || IsNil(o.ReadDate) {
		var ret time.Time
		return ret
	}
	return *o.ReadDate
}

// GetReadDateOk returns a tuple with the ReadDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetReadDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReadDate) {
		return nil, false
	}
	return o.ReadDate, true
}

// HasReadDate returns a boolean if a field has been set.
func (o *ContentDisposition) HasReadDate() bool {
	if o != nil && !IsNil(o.ReadDate) {
		return true
	}

	return false
}

// SetReadDate gets a reference to the given time.Time and assigns it to the ReadDate field.
func (o *ContentDisposition) SetReadDate(v time.Time) {
	o.ReadDate = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ContentDisposition) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ContentDisposition) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ContentDisposition) SetSize(v int64) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContentDisposition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContentDisposition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContentDisposition) SetType(v string) {
	o.Type = &v
}

func (o ContentDisposition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentDisposition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.ModificationDate) {
		toSerialize["modificationDate"] = o.ModificationDate
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.ReadDate) {
		toSerialize["readDate"] = o.ReadDate
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableContentDisposition struct {
	value *ContentDisposition
	isSet bool
}

func (v NullableContentDisposition) Get() *ContentDisposition {
	return v.value
}

func (v *NullableContentDisposition) Set(val *ContentDisposition) {
	v.value = val
	v.isSet = true
}

func (v NullableContentDisposition) IsSet() bool {
	return v.isSet
}

func (v *NullableContentDisposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentDisposition(val *ContentDisposition) *NullableContentDisposition {
	return &NullableContentDisposition{value: val, isSet: true}
}

func (v NullableContentDisposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentDisposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


