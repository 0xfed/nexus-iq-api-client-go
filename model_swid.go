/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.180.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the Swid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Swid{}

// Swid struct for Swid
type Swid struct {
	Name *string `json:"name,omitempty"`
	Patch *bool `json:"patch,omitempty"`
	TagId *string `json:"tagId,omitempty"`
	TagVersion *int32 `json:"tagVersion,omitempty"`
	Text *AttachmentText `json:"text,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewSwid instantiates a new Swid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwid() *Swid {
	this := Swid{}
	return &this
}

// NewSwidWithDefaults instantiates a new Swid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwidWithDefaults() *Swid {
	this := Swid{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Swid) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Swid) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Swid) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Swid) SetName(v string) {
	o.Name = &v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *Swid) GetPatch() bool {
	if o == nil || IsNil(o.Patch) {
		var ret bool
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Swid) GetPatchOk() (*bool, bool) {
	if o == nil || IsNil(o.Patch) {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *Swid) HasPatch() bool {
	if o != nil && !IsNil(o.Patch) {
		return true
	}

	return false
}

// SetPatch gets a reference to the given bool and assigns it to the Patch field.
func (o *Swid) SetPatch(v bool) {
	o.Patch = &v
}

// GetTagId returns the TagId field value if set, zero value otherwise.
func (o *Swid) GetTagId() string {
	if o == nil || IsNil(o.TagId) {
		var ret string
		return ret
	}
	return *o.TagId
}

// GetTagIdOk returns a tuple with the TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Swid) GetTagIdOk() (*string, bool) {
	if o == nil || IsNil(o.TagId) {
		return nil, false
	}
	return o.TagId, true
}

// HasTagId returns a boolean if a field has been set.
func (o *Swid) HasTagId() bool {
	if o != nil && !IsNil(o.TagId) {
		return true
	}

	return false
}

// SetTagId gets a reference to the given string and assigns it to the TagId field.
func (o *Swid) SetTagId(v string) {
	o.TagId = &v
}

// GetTagVersion returns the TagVersion field value if set, zero value otherwise.
func (o *Swid) GetTagVersion() int32 {
	if o == nil || IsNil(o.TagVersion) {
		var ret int32
		return ret
	}
	return *o.TagVersion
}

// GetTagVersionOk returns a tuple with the TagVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Swid) GetTagVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.TagVersion) {
		return nil, false
	}
	return o.TagVersion, true
}

// HasTagVersion returns a boolean if a field has been set.
func (o *Swid) HasTagVersion() bool {
	if o != nil && !IsNil(o.TagVersion) {
		return true
	}

	return false
}

// SetTagVersion gets a reference to the given int32 and assigns it to the TagVersion field.
func (o *Swid) SetTagVersion(v int32) {
	o.TagVersion = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Swid) GetText() AttachmentText {
	if o == nil || IsNil(o.Text) {
		var ret AttachmentText
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Swid) GetTextOk() (*AttachmentText, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Swid) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given AttachmentText and assigns it to the Text field.
func (o *Swid) SetText(v AttachmentText) {
	o.Text = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Swid) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Swid) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Swid) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Swid) SetVersion(v string) {
	o.Version = &v
}

func (o Swid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Swid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Patch) {
		toSerialize["patch"] = o.Patch
	}
	if !IsNil(o.TagId) {
		toSerialize["tagId"] = o.TagId
	}
	if !IsNil(o.TagVersion) {
		toSerialize["tagVersion"] = o.TagVersion
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableSwid struct {
	value *Swid
	isSet bool
}

func (v NullableSwid) Get() *Swid {
	return v.value
}

func (v *NullableSwid) Set(val *Swid) {
	v.value = val
	v.isSet = true
}

func (v NullableSwid) IsSet() bool {
	return v.isSet
}

func (v *NullableSwid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwid(val *Swid) *NullableSwid {
	return &NullableSwid{value: val, isSet: true}
}

func (v NullableSwid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


