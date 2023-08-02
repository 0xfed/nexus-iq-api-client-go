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

// checks if the LegalSourceLinkDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegalSourceLinkDTO{}

// LegalSourceLinkDTO struct for LegalSourceLinkDTO
type LegalSourceLinkDTO struct {
	Content *string `json:"content,omitempty"`
	Id *string `json:"id,omitempty"`
	OriginalContent *string `json:"originalContent,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewLegalSourceLinkDTO instantiates a new LegalSourceLinkDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalSourceLinkDTO() *LegalSourceLinkDTO {
	this := LegalSourceLinkDTO{}
	return &this
}

// NewLegalSourceLinkDTOWithDefaults instantiates a new LegalSourceLinkDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalSourceLinkDTOWithDefaults() *LegalSourceLinkDTO {
	this := LegalSourceLinkDTO{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *LegalSourceLinkDTO) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalSourceLinkDTO) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *LegalSourceLinkDTO) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *LegalSourceLinkDTO) SetContent(v string) {
	o.Content = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LegalSourceLinkDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalSourceLinkDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LegalSourceLinkDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LegalSourceLinkDTO) SetId(v string) {
	o.Id = &v
}

// GetOriginalContent returns the OriginalContent field value if set, zero value otherwise.
func (o *LegalSourceLinkDTO) GetOriginalContent() string {
	if o == nil || IsNil(o.OriginalContent) {
		var ret string
		return ret
	}
	return *o.OriginalContent
}

// GetOriginalContentOk returns a tuple with the OriginalContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalSourceLinkDTO) GetOriginalContentOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalContent) {
		return nil, false
	}
	return o.OriginalContent, true
}

// HasOriginalContent returns a boolean if a field has been set.
func (o *LegalSourceLinkDTO) HasOriginalContent() bool {
	if o != nil && !IsNil(o.OriginalContent) {
		return true
	}

	return false
}

// SetOriginalContent gets a reference to the given string and assigns it to the OriginalContent field.
func (o *LegalSourceLinkDTO) SetOriginalContent(v string) {
	o.OriginalContent = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LegalSourceLinkDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalSourceLinkDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LegalSourceLinkDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LegalSourceLinkDTO) SetStatus(v string) {
	o.Status = &v
}

func (o LegalSourceLinkDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegalSourceLinkDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OriginalContent) {
		toSerialize["originalContent"] = o.OriginalContent
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableLegalSourceLinkDTO struct {
	value *LegalSourceLinkDTO
	isSet bool
}

func (v NullableLegalSourceLinkDTO) Get() *LegalSourceLinkDTO {
	return v.value
}

func (v *NullableLegalSourceLinkDTO) Set(val *LegalSourceLinkDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalSourceLinkDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalSourceLinkDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalSourceLinkDTO(val *LegalSourceLinkDTO) *NullableLegalSourceLinkDTO {
	return &NullableLegalSourceLinkDTO{value: val, isSet: true}
}

func (v NullableLegalSourceLinkDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalSourceLinkDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


