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

// checks if the ApiLicenseLegalObligationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseLegalObligationDTO{}

// ApiLicenseLegalObligationDTO struct for ApiLicenseLegalObligationDTO
type ApiLicenseLegalObligationDTO struct {
	Comment *string `json:"comment,omitempty"`
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"`
	LastUpdatedByUsername *string `json:"lastUpdatedByUsername,omitempty"`
	Name *string `json:"name,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewApiLicenseLegalObligationDTO instantiates a new ApiLicenseLegalObligationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseLegalObligationDTO() *ApiLicenseLegalObligationDTO {
	this := ApiLicenseLegalObligationDTO{}
	return &this
}

// NewApiLicenseLegalObligationDTOWithDefaults instantiates a new ApiLicenseLegalObligationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseLegalObligationDTOWithDefaults() *ApiLicenseLegalObligationDTO {
	this := ApiLicenseLegalObligationDTO{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiLicenseLegalObligationDTO) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalObligationDTO) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiLicenseLegalObligationDTO) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiLicenseLegalObligationDTO) SetComment(v string) {
	o.Comment = &v
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiLicenseLegalObligationDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalObligationDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiLicenseLegalObligationDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiLicenseLegalObligationDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiLicenseLegalObligationDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalObligationDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiLicenseLegalObligationDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiLicenseLegalObligationDTO) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *ApiLicenseLegalObligationDTO) GetLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.LastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalObligationDTO) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedAt) {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *ApiLicenseLegalObligationDTO) HasLastUpdatedAt() bool {
	if o != nil && !IsNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *ApiLicenseLegalObligationDTO) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

// GetLastUpdatedByUsername returns the LastUpdatedByUsername field value if set, zero value otherwise.
func (o *ApiLicenseLegalObligationDTO) GetLastUpdatedByUsername() string {
	if o == nil || IsNil(o.LastUpdatedByUsername) {
		var ret string
		return ret
	}
	return *o.LastUpdatedByUsername
}

// GetLastUpdatedByUsernameOk returns a tuple with the LastUpdatedByUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalObligationDTO) GetLastUpdatedByUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedByUsername) {
		return nil, false
	}
	return o.LastUpdatedByUsername, true
}

// HasLastUpdatedByUsername returns a boolean if a field has been set.
func (o *ApiLicenseLegalObligationDTO) HasLastUpdatedByUsername() bool {
	if o != nil && !IsNil(o.LastUpdatedByUsername) {
		return true
	}

	return false
}

// SetLastUpdatedByUsername gets a reference to the given string and assigns it to the LastUpdatedByUsername field.
func (o *ApiLicenseLegalObligationDTO) SetLastUpdatedByUsername(v string) {
	o.LastUpdatedByUsername = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiLicenseLegalObligationDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalObligationDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiLicenseLegalObligationDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiLicenseLegalObligationDTO) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ApiLicenseLegalObligationDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalObligationDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ApiLicenseLegalObligationDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ApiLicenseLegalObligationDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiLicenseLegalObligationDTO) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalObligationDTO) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiLicenseLegalObligationDTO) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiLicenseLegalObligationDTO) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiLicenseLegalObligationDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalObligationDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiLicenseLegalObligationDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiLicenseLegalObligationDTO) SetStatus(v string) {
	o.Status = &v
}

func (o ApiLicenseLegalObligationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseLegalObligationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	if !IsNil(o.LastUpdatedByUsername) {
		toSerialize["lastUpdatedByUsername"] = o.LastUpdatedByUsername
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableApiLicenseLegalObligationDTO struct {
	value *ApiLicenseLegalObligationDTO
	isSet bool
}

func (v NullableApiLicenseLegalObligationDTO) Get() *ApiLicenseLegalObligationDTO {
	return v.value
}

func (v *NullableApiLicenseLegalObligationDTO) Set(val *ApiLicenseLegalObligationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseLegalObligationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseLegalObligationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseLegalObligationDTO(val *ApiLicenseLegalObligationDTO) *NullableApiLicenseLegalObligationDTO {
	return &NullableApiLicenseLegalObligationDTO{value: val, isSet: true}
}

func (v NullableApiLicenseLegalObligationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseLegalObligationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


