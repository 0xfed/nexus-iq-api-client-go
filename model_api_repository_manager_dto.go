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

// checks if the ApiRepositoryManagerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryManagerDTO{}

// ApiRepositoryManagerDTO struct for ApiRepositoryManagerDTO
type ApiRepositoryManagerDTO struct {
	Id *string `json:"id,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	Name *string `json:"name,omitempty"`
	ProductName *string `json:"productName,omitempty"`
	ProductVersion *string `json:"productVersion,omitempty"`
}

// NewApiRepositoryManagerDTO instantiates a new ApiRepositoryManagerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryManagerDTO() *ApiRepositoryManagerDTO {
	this := ApiRepositoryManagerDTO{}
	return &this
}

// NewApiRepositoryManagerDTOWithDefaults instantiates a new ApiRepositoryManagerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryManagerDTOWithDefaults() *ApiRepositoryManagerDTO {
	this := ApiRepositoryManagerDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiRepositoryManagerDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryManagerDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiRepositoryManagerDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiRepositoryManagerDTO) SetId(v string) {
	o.Id = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *ApiRepositoryManagerDTO) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryManagerDTO) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ApiRepositoryManagerDTO) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *ApiRepositoryManagerDTO) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiRepositoryManagerDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryManagerDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiRepositoryManagerDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiRepositoryManagerDTO) SetName(v string) {
	o.Name = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *ApiRepositoryManagerDTO) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryManagerDTO) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *ApiRepositoryManagerDTO) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *ApiRepositoryManagerDTO) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *ApiRepositoryManagerDTO) GetProductVersion() string {
	if o == nil || IsNil(o.ProductVersion) {
		var ret string
		return ret
	}
	return *o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryManagerDTO) GetProductVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVersion) {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *ApiRepositoryManagerDTO) HasProductVersion() bool {
	if o != nil && !IsNil(o.ProductVersion) {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.
func (o *ApiRepositoryManagerDTO) SetProductVersion(v string) {
	o.ProductVersion = &v
}

func (o ApiRepositoryManagerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryManagerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.ProductVersion) {
		toSerialize["productVersion"] = o.ProductVersion
	}
	return toSerialize, nil
}

type NullableApiRepositoryManagerDTO struct {
	value *ApiRepositoryManagerDTO
	isSet bool
}

func (v NullableApiRepositoryManagerDTO) Get() *ApiRepositoryManagerDTO {
	return v.value
}

func (v *NullableApiRepositoryManagerDTO) Set(val *ApiRepositoryManagerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryManagerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryManagerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryManagerDTO(val *ApiRepositoryManagerDTO) *NullableApiRepositoryManagerDTO {
	return &NullableApiRepositoryManagerDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryManagerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryManagerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


