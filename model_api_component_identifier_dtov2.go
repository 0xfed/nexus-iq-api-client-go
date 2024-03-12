/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.174.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiComponentIdentifierDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentIdentifierDTOV2{}

// ApiComponentIdentifierDTOV2 struct for ApiComponentIdentifierDTOV2
type ApiComponentIdentifierDTOV2 struct {
	Coordinates *map[string]string `json:"coordinates,omitempty"`
	Format *string `json:"format,omitempty"`
}

// NewApiComponentIdentifierDTOV2 instantiates a new ApiComponentIdentifierDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentIdentifierDTOV2() *ApiComponentIdentifierDTOV2 {
	this := ApiComponentIdentifierDTOV2{}
	return &this
}

// NewApiComponentIdentifierDTOV2WithDefaults instantiates a new ApiComponentIdentifierDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentIdentifierDTOV2WithDefaults() *ApiComponentIdentifierDTOV2 {
	this := ApiComponentIdentifierDTOV2{}
	return &this
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *ApiComponentIdentifierDTOV2) GetCoordinates() map[string]string {
	if o == nil || IsNil(o.Coordinates) {
		var ret map[string]string
		return ret
	}
	return *o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentIdentifierDTOV2) GetCoordinatesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Coordinates) {
		return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *ApiComponentIdentifierDTOV2) HasCoordinates() bool {
	if o != nil && !IsNil(o.Coordinates) {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given map[string]string and assigns it to the Coordinates field.
func (o *ApiComponentIdentifierDTOV2) SetCoordinates(v map[string]string) {
	o.Coordinates = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ApiComponentIdentifierDTOV2) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentIdentifierDTOV2) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ApiComponentIdentifierDTOV2) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ApiComponentIdentifierDTOV2) SetFormat(v string) {
	o.Format = &v
}

func (o ApiComponentIdentifierDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentIdentifierDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Coordinates) {
		toSerialize["coordinates"] = o.Coordinates
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	return toSerialize, nil
}

type NullableApiComponentIdentifierDTOV2 struct {
	value *ApiComponentIdentifierDTOV2
	isSet bool
}

func (v NullableApiComponentIdentifierDTOV2) Get() *ApiComponentIdentifierDTOV2 {
	return v.value
}

func (v *NullableApiComponentIdentifierDTOV2) Set(val *ApiComponentIdentifierDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentIdentifierDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentIdentifierDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentIdentifierDTOV2(val *ApiComponentIdentifierDTOV2) *NullableApiComponentIdentifierDTOV2 {
	return &NullableApiComponentIdentifierDTOV2{value: val, isSet: true}
}

func (v NullableApiComponentIdentifierDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentIdentifierDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


