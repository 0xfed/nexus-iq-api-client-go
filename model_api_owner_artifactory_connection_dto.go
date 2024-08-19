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

// checks if the ApiOwnerArtifactoryConnectionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiOwnerArtifactoryConnectionDTO{}

// ApiOwnerArtifactoryConnectionDTO struct for ApiOwnerArtifactoryConnectionDTO
type ApiOwnerArtifactoryConnectionDTO struct {
	ArtifactoryConnection *ApiArtifactoryConnectionDTO `json:"artifactoryConnection,omitempty"`
	ArtifactoryConnectionStatus *ApiArtifactoryConnectionStatusResponseDTO `json:"artifactoryConnectionStatus,omitempty"`
	OwnerDTO *ApiOwnerDTO `json:"ownerDTO,omitempty"`
}

// NewApiOwnerArtifactoryConnectionDTO instantiates a new ApiOwnerArtifactoryConnectionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiOwnerArtifactoryConnectionDTO() *ApiOwnerArtifactoryConnectionDTO {
	this := ApiOwnerArtifactoryConnectionDTO{}
	return &this
}

// NewApiOwnerArtifactoryConnectionDTOWithDefaults instantiates a new ApiOwnerArtifactoryConnectionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiOwnerArtifactoryConnectionDTOWithDefaults() *ApiOwnerArtifactoryConnectionDTO {
	this := ApiOwnerArtifactoryConnectionDTO{}
	return &this
}

// GetArtifactoryConnection returns the ArtifactoryConnection field value if set, zero value otherwise.
func (o *ApiOwnerArtifactoryConnectionDTO) GetArtifactoryConnection() ApiArtifactoryConnectionDTO {
	if o == nil || IsNil(o.ArtifactoryConnection) {
		var ret ApiArtifactoryConnectionDTO
		return ret
	}
	return *o.ArtifactoryConnection
}

// GetArtifactoryConnectionOk returns a tuple with the ArtifactoryConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOwnerArtifactoryConnectionDTO) GetArtifactoryConnectionOk() (*ApiArtifactoryConnectionDTO, bool) {
	if o == nil || IsNil(o.ArtifactoryConnection) {
		return nil, false
	}
	return o.ArtifactoryConnection, true
}

// HasArtifactoryConnection returns a boolean if a field has been set.
func (o *ApiOwnerArtifactoryConnectionDTO) HasArtifactoryConnection() bool {
	if o != nil && !IsNil(o.ArtifactoryConnection) {
		return true
	}

	return false
}

// SetArtifactoryConnection gets a reference to the given ApiArtifactoryConnectionDTO and assigns it to the ArtifactoryConnection field.
func (o *ApiOwnerArtifactoryConnectionDTO) SetArtifactoryConnection(v ApiArtifactoryConnectionDTO) {
	o.ArtifactoryConnection = &v
}

// GetArtifactoryConnectionStatus returns the ArtifactoryConnectionStatus field value if set, zero value otherwise.
func (o *ApiOwnerArtifactoryConnectionDTO) GetArtifactoryConnectionStatus() ApiArtifactoryConnectionStatusResponseDTO {
	if o == nil || IsNil(o.ArtifactoryConnectionStatus) {
		var ret ApiArtifactoryConnectionStatusResponseDTO
		return ret
	}
	return *o.ArtifactoryConnectionStatus
}

// GetArtifactoryConnectionStatusOk returns a tuple with the ArtifactoryConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOwnerArtifactoryConnectionDTO) GetArtifactoryConnectionStatusOk() (*ApiArtifactoryConnectionStatusResponseDTO, bool) {
	if o == nil || IsNil(o.ArtifactoryConnectionStatus) {
		return nil, false
	}
	return o.ArtifactoryConnectionStatus, true
}

// HasArtifactoryConnectionStatus returns a boolean if a field has been set.
func (o *ApiOwnerArtifactoryConnectionDTO) HasArtifactoryConnectionStatus() bool {
	if o != nil && !IsNil(o.ArtifactoryConnectionStatus) {
		return true
	}

	return false
}

// SetArtifactoryConnectionStatus gets a reference to the given ApiArtifactoryConnectionStatusResponseDTO and assigns it to the ArtifactoryConnectionStatus field.
func (o *ApiOwnerArtifactoryConnectionDTO) SetArtifactoryConnectionStatus(v ApiArtifactoryConnectionStatusResponseDTO) {
	o.ArtifactoryConnectionStatus = &v
}

// GetOwnerDTO returns the OwnerDTO field value if set, zero value otherwise.
func (o *ApiOwnerArtifactoryConnectionDTO) GetOwnerDTO() ApiOwnerDTO {
	if o == nil || IsNil(o.OwnerDTO) {
		var ret ApiOwnerDTO
		return ret
	}
	return *o.OwnerDTO
}

// GetOwnerDTOOk returns a tuple with the OwnerDTO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOwnerArtifactoryConnectionDTO) GetOwnerDTOOk() (*ApiOwnerDTO, bool) {
	if o == nil || IsNil(o.OwnerDTO) {
		return nil, false
	}
	return o.OwnerDTO, true
}

// HasOwnerDTO returns a boolean if a field has been set.
func (o *ApiOwnerArtifactoryConnectionDTO) HasOwnerDTO() bool {
	if o != nil && !IsNil(o.OwnerDTO) {
		return true
	}

	return false
}

// SetOwnerDTO gets a reference to the given ApiOwnerDTO and assigns it to the OwnerDTO field.
func (o *ApiOwnerArtifactoryConnectionDTO) SetOwnerDTO(v ApiOwnerDTO) {
	o.OwnerDTO = &v
}

func (o ApiOwnerArtifactoryConnectionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiOwnerArtifactoryConnectionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArtifactoryConnection) {
		toSerialize["artifactoryConnection"] = o.ArtifactoryConnection
	}
	if !IsNil(o.ArtifactoryConnectionStatus) {
		toSerialize["artifactoryConnectionStatus"] = o.ArtifactoryConnectionStatus
	}
	if !IsNil(o.OwnerDTO) {
		toSerialize["ownerDTO"] = o.OwnerDTO
	}
	return toSerialize, nil
}

type NullableApiOwnerArtifactoryConnectionDTO struct {
	value *ApiOwnerArtifactoryConnectionDTO
	isSet bool
}

func (v NullableApiOwnerArtifactoryConnectionDTO) Get() *ApiOwnerArtifactoryConnectionDTO {
	return v.value
}

func (v *NullableApiOwnerArtifactoryConnectionDTO) Set(val *ApiOwnerArtifactoryConnectionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiOwnerArtifactoryConnectionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiOwnerArtifactoryConnectionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiOwnerArtifactoryConnectionDTO(val *ApiOwnerArtifactoryConnectionDTO) *NullableApiOwnerArtifactoryConnectionDTO {
	return &NullableApiOwnerArtifactoryConnectionDTO{value: val, isSet: true}
}

func (v NullableApiOwnerArtifactoryConnectionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiOwnerArtifactoryConnectionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


