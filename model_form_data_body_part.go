/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.181.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the FormDataBodyPart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormDataBodyPart{}

// FormDataBodyPart struct for FormDataBodyPart
type FormDataBodyPart struct {
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`
	Entity map[string]interface{} `json:"entity,omitempty"`
	FormDataContentDisposition *FormDataContentDisposition `json:"formDataContentDisposition,omitempty"`
	Headers *map[string][]string `json:"headers,omitempty"`
	MediaType *BodyPartMediaType `json:"mediaType,omitempty"`
	MessageBodyWorkers map[string]interface{} `json:"messageBodyWorkers,omitempty"`
	Name *string `json:"name,omitempty"`
	ParameterizedHeaders *map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`
	Parent *MultiPart `json:"parent,omitempty"`
	Providers map[string]interface{} `json:"providers,omitempty"`
	Simple *bool `json:"simple,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewFormDataBodyPart instantiates a new FormDataBodyPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormDataBodyPart() *FormDataBodyPart {
	this := FormDataBodyPart{}
	return &this
}

// NewFormDataBodyPartWithDefaults instantiates a new FormDataBodyPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormDataBodyPartWithDefaults() *FormDataBodyPart {
	this := FormDataBodyPart{}
	return &this
}

// GetContentDisposition returns the ContentDisposition field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetContentDisposition() ContentDisposition {
	if o == nil || IsNil(o.ContentDisposition) {
		var ret ContentDisposition
		return ret
	}
	return *o.ContentDisposition
}

// GetContentDispositionOk returns a tuple with the ContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetContentDispositionOk() (*ContentDisposition, bool) {
	if o == nil || IsNil(o.ContentDisposition) {
		return nil, false
	}
	return o.ContentDisposition, true
}

// HasContentDisposition returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasContentDisposition() bool {
	if o != nil && !IsNil(o.ContentDisposition) {
		return true
	}

	return false
}

// SetContentDisposition gets a reference to the given ContentDisposition and assigns it to the ContentDisposition field.
func (o *FormDataBodyPart) SetContentDisposition(v ContentDisposition) {
	o.ContentDisposition = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetEntity() map[string]interface{} {
	if o == nil || IsNil(o.Entity) {
		var ret map[string]interface{}
		return ret
	}
	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetEntityOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Entity) {
		return map[string]interface{}{}, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given map[string]interface{} and assigns it to the Entity field.
func (o *FormDataBodyPart) SetEntity(v map[string]interface{}) {
	o.Entity = v
}

// GetFormDataContentDisposition returns the FormDataContentDisposition field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetFormDataContentDisposition() FormDataContentDisposition {
	if o == nil || IsNil(o.FormDataContentDisposition) {
		var ret FormDataContentDisposition
		return ret
	}
	return *o.FormDataContentDisposition
}

// GetFormDataContentDispositionOk returns a tuple with the FormDataContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetFormDataContentDispositionOk() (*FormDataContentDisposition, bool) {
	if o == nil || IsNil(o.FormDataContentDisposition) {
		return nil, false
	}
	return o.FormDataContentDisposition, true
}

// HasFormDataContentDisposition returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasFormDataContentDisposition() bool {
	if o != nil && !IsNil(o.FormDataContentDisposition) {
		return true
	}

	return false
}

// SetFormDataContentDisposition gets a reference to the given FormDataContentDisposition and assigns it to the FormDataContentDisposition field.
func (o *FormDataBodyPart) SetFormDataContentDisposition(v FormDataContentDisposition) {
	o.FormDataContentDisposition = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetHeaders() map[string][]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string][]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetHeadersOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string][]string and assigns it to the Headers field.
func (o *FormDataBodyPart) SetHeaders(v map[string][]string) {
	o.Headers = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetMediaType() BodyPartMediaType {
	if o == nil || IsNil(o.MediaType) {
		var ret BodyPartMediaType
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetMediaTypeOk() (*BodyPartMediaType, bool) {
	if o == nil || IsNil(o.MediaType) {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasMediaType() bool {
	if o != nil && !IsNil(o.MediaType) {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given BodyPartMediaType and assigns it to the MediaType field.
func (o *FormDataBodyPart) SetMediaType(v BodyPartMediaType) {
	o.MediaType = &v
}

// GetMessageBodyWorkers returns the MessageBodyWorkers field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetMessageBodyWorkers() map[string]interface{} {
	if o == nil || IsNil(o.MessageBodyWorkers) {
		var ret map[string]interface{}
		return ret
	}
	return o.MessageBodyWorkers
}

// GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetMessageBodyWorkersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MessageBodyWorkers) {
		return map[string]interface{}{}, false
	}
	return o.MessageBodyWorkers, true
}

// HasMessageBodyWorkers returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasMessageBodyWorkers() bool {
	if o != nil && !IsNil(o.MessageBodyWorkers) {
		return true
	}

	return false
}

// SetMessageBodyWorkers gets a reference to the given map[string]interface{} and assigns it to the MessageBodyWorkers field.
func (o *FormDataBodyPart) SetMessageBodyWorkers(v map[string]interface{}) {
	o.MessageBodyWorkers = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FormDataBodyPart) SetName(v string) {
	o.Name = &v
}

// GetParameterizedHeaders returns the ParameterizedHeaders field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetParameterizedHeaders() map[string][]ParameterizedHeader {
	if o == nil || IsNil(o.ParameterizedHeaders) {
		var ret map[string][]ParameterizedHeader
		return ret
	}
	return *o.ParameterizedHeaders
}

// GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool) {
	if o == nil || IsNil(o.ParameterizedHeaders) {
		return nil, false
	}
	return o.ParameterizedHeaders, true
}

// HasParameterizedHeaders returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasParameterizedHeaders() bool {
	if o != nil && !IsNil(o.ParameterizedHeaders) {
		return true
	}

	return false
}

// SetParameterizedHeaders gets a reference to the given map[string][]ParameterizedHeader and assigns it to the ParameterizedHeaders field.
func (o *FormDataBodyPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader) {
	o.ParameterizedHeaders = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetParent() MultiPart {
	if o == nil || IsNil(o.Parent) {
		var ret MultiPart
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetParentOk() (*MultiPart, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given MultiPart and assigns it to the Parent field.
func (o *FormDataBodyPart) SetParent(v MultiPart) {
	o.Parent = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetProviders() map[string]interface{} {
	if o == nil || IsNil(o.Providers) {
		var ret map[string]interface{}
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetProvidersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Providers) {
		return map[string]interface{}{}, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given map[string]interface{} and assigns it to the Providers field.
func (o *FormDataBodyPart) SetProviders(v map[string]interface{}) {
	o.Providers = v
}

// GetSimple returns the Simple field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetSimple() bool {
	if o == nil || IsNil(o.Simple) {
		var ret bool
		return ret
	}
	return *o.Simple
}

// GetSimpleOk returns a tuple with the Simple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetSimpleOk() (*bool, bool) {
	if o == nil || IsNil(o.Simple) {
		return nil, false
	}
	return o.Simple, true
}

// HasSimple returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasSimple() bool {
	if o != nil && !IsNil(o.Simple) {
		return true
	}

	return false
}

// SetSimple gets a reference to the given bool and assigns it to the Simple field.
func (o *FormDataBodyPart) SetSimple(v bool) {
	o.Simple = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FormDataBodyPart) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataBodyPart) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FormDataBodyPart) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FormDataBodyPart) SetValue(v string) {
	o.Value = &v
}

func (o FormDataBodyPart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormDataBodyPart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentDisposition) {
		toSerialize["contentDisposition"] = o.ContentDisposition
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.FormDataContentDisposition) {
		toSerialize["formDataContentDisposition"] = o.FormDataContentDisposition
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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
	if !IsNil(o.Simple) {
		toSerialize["simple"] = o.Simple
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFormDataBodyPart struct {
	value *FormDataBodyPart
	isSet bool
}

func (v NullableFormDataBodyPart) Get() *FormDataBodyPart {
	return v.value
}

func (v *NullableFormDataBodyPart) Set(val *FormDataBodyPart) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDataBodyPart) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDataBodyPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDataBodyPart(val *FormDataBodyPart) *NullableFormDataBodyPart {
	return &NullableFormDataBodyPart{value: val, isSet: true}
}

func (v NullableFormDataBodyPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDataBodyPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


