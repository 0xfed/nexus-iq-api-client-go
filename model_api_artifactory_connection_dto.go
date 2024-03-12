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

// checks if the ApiArtifactoryConnectionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiArtifactoryConnectionDTO{}

// ApiArtifactoryConnectionDTO struct for ApiArtifactoryConnectionDTO
type ApiArtifactoryConnectionDTO struct {
	ArtifactoryConnectionId *string `json:"artifactoryConnectionId,omitempty"`
	BaseUrl *string `json:"baseUrl,omitempty"`
	IsAnonymous *bool `json:"isAnonymous,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerType *string `json:"ownerType,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewApiArtifactoryConnectionDTO instantiates a new ApiArtifactoryConnectionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiArtifactoryConnectionDTO() *ApiArtifactoryConnectionDTO {
	this := ApiArtifactoryConnectionDTO{}
	return &this
}

// NewApiArtifactoryConnectionDTOWithDefaults instantiates a new ApiArtifactoryConnectionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiArtifactoryConnectionDTOWithDefaults() *ApiArtifactoryConnectionDTO {
	this := ApiArtifactoryConnectionDTO{}
	return &this
}

// GetArtifactoryConnectionId returns the ArtifactoryConnectionId field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionDTO) GetArtifactoryConnectionId() string {
	if o == nil || IsNil(o.ArtifactoryConnectionId) {
		var ret string
		return ret
	}
	return *o.ArtifactoryConnectionId
}

// GetArtifactoryConnectionIdOk returns a tuple with the ArtifactoryConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionDTO) GetArtifactoryConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactoryConnectionId) {
		return nil, false
	}
	return o.ArtifactoryConnectionId, true
}

// HasArtifactoryConnectionId returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionDTO) HasArtifactoryConnectionId() bool {
	if o != nil && !IsNil(o.ArtifactoryConnectionId) {
		return true
	}

	return false
}

// SetArtifactoryConnectionId gets a reference to the given string and assigns it to the ArtifactoryConnectionId field.
func (o *ApiArtifactoryConnectionDTO) SetArtifactoryConnectionId(v string) {
	o.ArtifactoryConnectionId = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionDTO) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionDTO) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionDTO) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *ApiArtifactoryConnectionDTO) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetIsAnonymous returns the IsAnonymous field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionDTO) GetIsAnonymous() bool {
	if o == nil || IsNil(o.IsAnonymous) {
		var ret bool
		return ret
	}
	return *o.IsAnonymous
}

// GetIsAnonymousOk returns a tuple with the IsAnonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionDTO) GetIsAnonymousOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnonymous) {
		return nil, false
	}
	return o.IsAnonymous, true
}

// HasIsAnonymous returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionDTO) HasIsAnonymous() bool {
	if o != nil && !IsNil(o.IsAnonymous) {
		return true
	}

	return false
}

// SetIsAnonymous gets a reference to the given bool and assigns it to the IsAnonymous field.
func (o *ApiArtifactoryConnectionDTO) SetIsAnonymous(v bool) {
	o.IsAnonymous = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ApiArtifactoryConnectionDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionDTO) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionDTO) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionDTO) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *ApiArtifactoryConnectionDTO) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionDTO) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionDTO) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionDTO) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiArtifactoryConnectionDTO) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionDTO) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionDTO) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionDTO) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiArtifactoryConnectionDTO) SetUsername(v string) {
	o.Username = &v
}

func (o ApiArtifactoryConnectionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiArtifactoryConnectionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArtifactoryConnectionId) {
		toSerialize["artifactoryConnectionId"] = o.ArtifactoryConnectionId
	}
	if !IsNil(o.BaseUrl) {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	if !IsNil(o.IsAnonymous) {
		toSerialize["isAnonymous"] = o.IsAnonymous
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.OwnerType) {
		toSerialize["ownerType"] = o.OwnerType
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableApiArtifactoryConnectionDTO struct {
	value *ApiArtifactoryConnectionDTO
	isSet bool
}

func (v NullableApiArtifactoryConnectionDTO) Get() *ApiArtifactoryConnectionDTO {
	return v.value
}

func (v *NullableApiArtifactoryConnectionDTO) Set(val *ApiArtifactoryConnectionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiArtifactoryConnectionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiArtifactoryConnectionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiArtifactoryConnectionDTO(val *ApiArtifactoryConnectionDTO) *NullableApiArtifactoryConnectionDTO {
	return &NullableApiArtifactoryConnectionDTO{value: val, isSet: true}
}

func (v NullableApiArtifactoryConnectionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiArtifactoryConnectionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


