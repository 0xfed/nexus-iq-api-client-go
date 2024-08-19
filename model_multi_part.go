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

// checks if the MultiPart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiPart{}

// MultiPart struct for MultiPart
type MultiPart struct {
	BodyParts []BodyPart `json:"bodyParts,omitempty"`
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`
	Entity map[string]interface{} `json:"entity,omitempty"`
	Headers *map[string][]string `json:"headers,omitempty"`
	MediaType *BodyPartMediaType `json:"mediaType,omitempty"`
	MessageBodyWorkers map[string]interface{} `json:"messageBodyWorkers,omitempty"`
	ParameterizedHeaders *map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`
	Parent *MultiPart `json:"parent,omitempty"`
	Providers map[string]interface{} `json:"providers,omitempty"`
}

// NewMultiPart instantiates a new MultiPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiPart() *MultiPart {
	this := MultiPart{}
	return &this
}

// NewMultiPartWithDefaults instantiates a new MultiPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiPartWithDefaults() *MultiPart {
	this := MultiPart{}
	return &this
}

// GetBodyParts returns the BodyParts field value if set, zero value otherwise.
func (o *MultiPart) GetBodyParts() []BodyPart {
	if o == nil || IsNil(o.BodyParts) {
		var ret []BodyPart
		return ret
	}
	return o.BodyParts
}

// GetBodyPartsOk returns a tuple with the BodyParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetBodyPartsOk() ([]BodyPart, bool) {
	if o == nil || IsNil(o.BodyParts) {
		return nil, false
	}
	return o.BodyParts, true
}

// HasBodyParts returns a boolean if a field has been set.
func (o *MultiPart) HasBodyParts() bool {
	if o != nil && !IsNil(o.BodyParts) {
		return true
	}

	return false
}

// SetBodyParts gets a reference to the given []BodyPart and assigns it to the BodyParts field.
func (o *MultiPart) SetBodyParts(v []BodyPart) {
	o.BodyParts = v
}

// GetContentDisposition returns the ContentDisposition field value if set, zero value otherwise.
func (o *MultiPart) GetContentDisposition() ContentDisposition {
	if o == nil || IsNil(o.ContentDisposition) {
		var ret ContentDisposition
		return ret
	}
	return *o.ContentDisposition
}

// GetContentDispositionOk returns a tuple with the ContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetContentDispositionOk() (*ContentDisposition, bool) {
	if o == nil || IsNil(o.ContentDisposition) {
		return nil, false
	}
	return o.ContentDisposition, true
}

// HasContentDisposition returns a boolean if a field has been set.
func (o *MultiPart) HasContentDisposition() bool {
	if o != nil && !IsNil(o.ContentDisposition) {
		return true
	}

	return false
}

// SetContentDisposition gets a reference to the given ContentDisposition and assigns it to the ContentDisposition field.
func (o *MultiPart) SetContentDisposition(v ContentDisposition) {
	o.ContentDisposition = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *MultiPart) GetEntity() map[string]interface{} {
	if o == nil || IsNil(o.Entity) {
		var ret map[string]interface{}
		return ret
	}
	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetEntityOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Entity) {
		return map[string]interface{}{}, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *MultiPart) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given map[string]interface{} and assigns it to the Entity field.
func (o *MultiPart) SetEntity(v map[string]interface{}) {
	o.Entity = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *MultiPart) GetHeaders() map[string][]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string][]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetHeadersOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *MultiPart) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string][]string and assigns it to the Headers field.
func (o *MultiPart) SetHeaders(v map[string][]string) {
	o.Headers = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *MultiPart) GetMediaType() BodyPartMediaType {
	if o == nil || IsNil(o.MediaType) {
		var ret BodyPartMediaType
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetMediaTypeOk() (*BodyPartMediaType, bool) {
	if o == nil || IsNil(o.MediaType) {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *MultiPart) HasMediaType() bool {
	if o != nil && !IsNil(o.MediaType) {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given BodyPartMediaType and assigns it to the MediaType field.
func (o *MultiPart) SetMediaType(v BodyPartMediaType) {
	o.MediaType = &v
}

// GetMessageBodyWorkers returns the MessageBodyWorkers field value if set, zero value otherwise.
func (o *MultiPart) GetMessageBodyWorkers() map[string]interface{} {
	if o == nil || IsNil(o.MessageBodyWorkers) {
		var ret map[string]interface{}
		return ret
	}
	return o.MessageBodyWorkers
}

// GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetMessageBodyWorkersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MessageBodyWorkers) {
		return map[string]interface{}{}, false
	}
	return o.MessageBodyWorkers, true
}

// HasMessageBodyWorkers returns a boolean if a field has been set.
func (o *MultiPart) HasMessageBodyWorkers() bool {
	if o != nil && !IsNil(o.MessageBodyWorkers) {
		return true
	}

	return false
}

// SetMessageBodyWorkers gets a reference to the given map[string]interface{} and assigns it to the MessageBodyWorkers field.
func (o *MultiPart) SetMessageBodyWorkers(v map[string]interface{}) {
	o.MessageBodyWorkers = v
}

// GetParameterizedHeaders returns the ParameterizedHeaders field value if set, zero value otherwise.
func (o *MultiPart) GetParameterizedHeaders() map[string][]ParameterizedHeader {
	if o == nil || IsNil(o.ParameterizedHeaders) {
		var ret map[string][]ParameterizedHeader
		return ret
	}
	return *o.ParameterizedHeaders
}

// GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool) {
	if o == nil || IsNil(o.ParameterizedHeaders) {
		return nil, false
	}
	return o.ParameterizedHeaders, true
}

// HasParameterizedHeaders returns a boolean if a field has been set.
func (o *MultiPart) HasParameterizedHeaders() bool {
	if o != nil && !IsNil(o.ParameterizedHeaders) {
		return true
	}

	return false
}

// SetParameterizedHeaders gets a reference to the given map[string][]ParameterizedHeader and assigns it to the ParameterizedHeaders field.
func (o *MultiPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader) {
	o.ParameterizedHeaders = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *MultiPart) GetParent() MultiPart {
	if o == nil || IsNil(o.Parent) {
		var ret MultiPart
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetParentOk() (*MultiPart, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *MultiPart) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given MultiPart and assigns it to the Parent field.
func (o *MultiPart) SetParent(v MultiPart) {
	o.Parent = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *MultiPart) GetProviders() map[string]interface{} {
	if o == nil || IsNil(o.Providers) {
		var ret map[string]interface{}
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetProvidersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Providers) {
		return map[string]interface{}{}, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *MultiPart) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given map[string]interface{} and assigns it to the Providers field.
func (o *MultiPart) SetProviders(v map[string]interface{}) {
	o.Providers = v
}

func (o MultiPart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiPart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BodyParts) {
		toSerialize["bodyParts"] = o.BodyParts
	}
	if !IsNil(o.ContentDisposition) {
		toSerialize["contentDisposition"] = o.ContentDisposition
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.MediaType) {
		toSerialize["mediaType"] = o.MediaType
	}
	if !IsNil(o.MessageBodyWorkers) {
		toSerialize["messageBodyWorkers"] = o.MessageBodyWorkers
	}
	if !IsNil(o.ParameterizedHeaders) {
		toSerialize["parameterizedHeaders"] = o.ParameterizedHeaders
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}
	return toSerialize, nil
}

type NullableMultiPart struct {
	value *MultiPart
	isSet bool
}

func (v NullableMultiPart) Get() *MultiPart {
	return v.value
}

func (v *NullableMultiPart) Set(val *MultiPart) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiPart) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiPart(val *MultiPart) *NullableMultiPart {
	return &NullableMultiPart{value: val, isSet: true}
}

func (v NullableMultiPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


