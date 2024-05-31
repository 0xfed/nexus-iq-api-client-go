/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.176.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiCompositeSourceControlDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCompositeSourceControlDTO{}

// ApiCompositeSourceControlDTO struct for ApiCompositeSourceControlDTO
type ApiCompositeSourceControlDTO struct {
	BaseBranch *ApiCompositeValueDTOString `json:"baseBranch,omitempty"`
	CommitStatusEnabled *ApiCompositeValueDTOBoolean `json:"commitStatusEnabled,omitempty"`
	Id *string `json:"id,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	Provider *ApiCompositeValueDTOString `json:"provider,omitempty"`
	PullRequestCommentingEnabled *ApiCompositeValueDTOBoolean `json:"pullRequestCommentingEnabled,omitempty"`
	RemediationPullRequestsEnabled *ApiCompositeValueDTOBoolean `json:"remediationPullRequestsEnabled,omitempty"`
	RepositoryUrl *string `json:"repositoryUrl,omitempty"`
	SourceControlEvaluationsEnabled *ApiCompositeValueDTOBoolean `json:"sourceControlEvaluationsEnabled,omitempty"`
	SourceControlScanTarget *ApiCompositeValueDTOString `json:"sourceControlScanTarget,omitempty"`
	SshEnabled *ApiCompositeValueDTOBoolean `json:"sshEnabled,omitempty"`
	StatusChecksEnabled *ApiCompositeValueDTOBoolean `json:"statusChecksEnabled,omitempty"`
	Token *ApiCompositeValueDTOString `json:"token,omitempty"`
	Username *ApiCompositeValueDTOString `json:"username,omitempty"`
}

// NewApiCompositeSourceControlDTO instantiates a new ApiCompositeSourceControlDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCompositeSourceControlDTO() *ApiCompositeSourceControlDTO {
	this := ApiCompositeSourceControlDTO{}
	return &this
}

// NewApiCompositeSourceControlDTOWithDefaults instantiates a new ApiCompositeSourceControlDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCompositeSourceControlDTOWithDefaults() *ApiCompositeSourceControlDTO {
	this := ApiCompositeSourceControlDTO{}
	return &this
}

// GetBaseBranch returns the BaseBranch field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetBaseBranch() ApiCompositeValueDTOString {
	if o == nil || IsNil(o.BaseBranch) {
		var ret ApiCompositeValueDTOString
		return ret
	}
	return *o.BaseBranch
}

// GetBaseBranchOk returns a tuple with the BaseBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetBaseBranchOk() (*ApiCompositeValueDTOString, bool) {
	if o == nil || IsNil(o.BaseBranch) {
		return nil, false
	}
	return o.BaseBranch, true
}

// HasBaseBranch returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasBaseBranch() bool {
	if o != nil && !IsNil(o.BaseBranch) {
		return true
	}

	return false
}

// SetBaseBranch gets a reference to the given ApiCompositeValueDTOString and assigns it to the BaseBranch field.
func (o *ApiCompositeSourceControlDTO) SetBaseBranch(v ApiCompositeValueDTOString) {
	o.BaseBranch = &v
}

// GetCommitStatusEnabled returns the CommitStatusEnabled field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetCommitStatusEnabled() ApiCompositeValueDTOBoolean {
	if o == nil || IsNil(o.CommitStatusEnabled) {
		var ret ApiCompositeValueDTOBoolean
		return ret
	}
	return *o.CommitStatusEnabled
}

// GetCommitStatusEnabledOk returns a tuple with the CommitStatusEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetCommitStatusEnabledOk() (*ApiCompositeValueDTOBoolean, bool) {
	if o == nil || IsNil(o.CommitStatusEnabled) {
		return nil, false
	}
	return o.CommitStatusEnabled, true
}

// HasCommitStatusEnabled returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasCommitStatusEnabled() bool {
	if o != nil && !IsNil(o.CommitStatusEnabled) {
		return true
	}

	return false
}

// SetCommitStatusEnabled gets a reference to the given ApiCompositeValueDTOBoolean and assigns it to the CommitStatusEnabled field.
func (o *ApiCompositeSourceControlDTO) SetCommitStatusEnabled(v ApiCompositeValueDTOBoolean) {
	o.CommitStatusEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiCompositeSourceControlDTO) SetId(v string) {
	o.Id = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ApiCompositeSourceControlDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetProvider() ApiCompositeValueDTOString {
	if o == nil || IsNil(o.Provider) {
		var ret ApiCompositeValueDTOString
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetProviderOk() (*ApiCompositeValueDTOString, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given ApiCompositeValueDTOString and assigns it to the Provider field.
func (o *ApiCompositeSourceControlDTO) SetProvider(v ApiCompositeValueDTOString) {
	o.Provider = &v
}

// GetPullRequestCommentingEnabled returns the PullRequestCommentingEnabled field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetPullRequestCommentingEnabled() ApiCompositeValueDTOBoolean {
	if o == nil || IsNil(o.PullRequestCommentingEnabled) {
		var ret ApiCompositeValueDTOBoolean
		return ret
	}
	return *o.PullRequestCommentingEnabled
}

// GetPullRequestCommentingEnabledOk returns a tuple with the PullRequestCommentingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetPullRequestCommentingEnabledOk() (*ApiCompositeValueDTOBoolean, bool) {
	if o == nil || IsNil(o.PullRequestCommentingEnabled) {
		return nil, false
	}
	return o.PullRequestCommentingEnabled, true
}

// HasPullRequestCommentingEnabled returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasPullRequestCommentingEnabled() bool {
	if o != nil && !IsNil(o.PullRequestCommentingEnabled) {
		return true
	}

	return false
}

// SetPullRequestCommentingEnabled gets a reference to the given ApiCompositeValueDTOBoolean and assigns it to the PullRequestCommentingEnabled field.
func (o *ApiCompositeSourceControlDTO) SetPullRequestCommentingEnabled(v ApiCompositeValueDTOBoolean) {
	o.PullRequestCommentingEnabled = &v
}

// GetRemediationPullRequestsEnabled returns the RemediationPullRequestsEnabled field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetRemediationPullRequestsEnabled() ApiCompositeValueDTOBoolean {
	if o == nil || IsNil(o.RemediationPullRequestsEnabled) {
		var ret ApiCompositeValueDTOBoolean
		return ret
	}
	return *o.RemediationPullRequestsEnabled
}

// GetRemediationPullRequestsEnabledOk returns a tuple with the RemediationPullRequestsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetRemediationPullRequestsEnabledOk() (*ApiCompositeValueDTOBoolean, bool) {
	if o == nil || IsNil(o.RemediationPullRequestsEnabled) {
		return nil, false
	}
	return o.RemediationPullRequestsEnabled, true
}

// HasRemediationPullRequestsEnabled returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasRemediationPullRequestsEnabled() bool {
	if o != nil && !IsNil(o.RemediationPullRequestsEnabled) {
		return true
	}

	return false
}

// SetRemediationPullRequestsEnabled gets a reference to the given ApiCompositeValueDTOBoolean and assigns it to the RemediationPullRequestsEnabled field.
func (o *ApiCompositeSourceControlDTO) SetRemediationPullRequestsEnabled(v ApiCompositeValueDTOBoolean) {
	o.RemediationPullRequestsEnabled = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetRepositoryUrl() string {
	if o == nil || IsNil(o.RepositoryUrl) {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryUrl) {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasRepositoryUrl() bool {
	if o != nil && !IsNil(o.RepositoryUrl) {
		return true
	}

	return false
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *ApiCompositeSourceControlDTO) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// GetSourceControlEvaluationsEnabled returns the SourceControlEvaluationsEnabled field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetSourceControlEvaluationsEnabled() ApiCompositeValueDTOBoolean {
	if o == nil || IsNil(o.SourceControlEvaluationsEnabled) {
		var ret ApiCompositeValueDTOBoolean
		return ret
	}
	return *o.SourceControlEvaluationsEnabled
}

// GetSourceControlEvaluationsEnabledOk returns a tuple with the SourceControlEvaluationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetSourceControlEvaluationsEnabledOk() (*ApiCompositeValueDTOBoolean, bool) {
	if o == nil || IsNil(o.SourceControlEvaluationsEnabled) {
		return nil, false
	}
	return o.SourceControlEvaluationsEnabled, true
}

// HasSourceControlEvaluationsEnabled returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasSourceControlEvaluationsEnabled() bool {
	if o != nil && !IsNil(o.SourceControlEvaluationsEnabled) {
		return true
	}

	return false
}

// SetSourceControlEvaluationsEnabled gets a reference to the given ApiCompositeValueDTOBoolean and assigns it to the SourceControlEvaluationsEnabled field.
func (o *ApiCompositeSourceControlDTO) SetSourceControlEvaluationsEnabled(v ApiCompositeValueDTOBoolean) {
	o.SourceControlEvaluationsEnabled = &v
}

// GetSourceControlScanTarget returns the SourceControlScanTarget field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetSourceControlScanTarget() ApiCompositeValueDTOString {
	if o == nil || IsNil(o.SourceControlScanTarget) {
		var ret ApiCompositeValueDTOString
		return ret
	}
	return *o.SourceControlScanTarget
}

// GetSourceControlScanTargetOk returns a tuple with the SourceControlScanTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetSourceControlScanTargetOk() (*ApiCompositeValueDTOString, bool) {
	if o == nil || IsNil(o.SourceControlScanTarget) {
		return nil, false
	}
	return o.SourceControlScanTarget, true
}

// HasSourceControlScanTarget returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasSourceControlScanTarget() bool {
	if o != nil && !IsNil(o.SourceControlScanTarget) {
		return true
	}

	return false
}

// SetSourceControlScanTarget gets a reference to the given ApiCompositeValueDTOString and assigns it to the SourceControlScanTarget field.
func (o *ApiCompositeSourceControlDTO) SetSourceControlScanTarget(v ApiCompositeValueDTOString) {
	o.SourceControlScanTarget = &v
}

// GetSshEnabled returns the SshEnabled field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetSshEnabled() ApiCompositeValueDTOBoolean {
	if o == nil || IsNil(o.SshEnabled) {
		var ret ApiCompositeValueDTOBoolean
		return ret
	}
	return *o.SshEnabled
}

// GetSshEnabledOk returns a tuple with the SshEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetSshEnabledOk() (*ApiCompositeValueDTOBoolean, bool) {
	if o == nil || IsNil(o.SshEnabled) {
		return nil, false
	}
	return o.SshEnabled, true
}

// HasSshEnabled returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasSshEnabled() bool {
	if o != nil && !IsNil(o.SshEnabled) {
		return true
	}

	return false
}

// SetSshEnabled gets a reference to the given ApiCompositeValueDTOBoolean and assigns it to the SshEnabled field.
func (o *ApiCompositeSourceControlDTO) SetSshEnabled(v ApiCompositeValueDTOBoolean) {
	o.SshEnabled = &v
}

// GetStatusChecksEnabled returns the StatusChecksEnabled field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetStatusChecksEnabled() ApiCompositeValueDTOBoolean {
	if o == nil || IsNil(o.StatusChecksEnabled) {
		var ret ApiCompositeValueDTOBoolean
		return ret
	}
	return *o.StatusChecksEnabled
}

// GetStatusChecksEnabledOk returns a tuple with the StatusChecksEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetStatusChecksEnabledOk() (*ApiCompositeValueDTOBoolean, bool) {
	if o == nil || IsNil(o.StatusChecksEnabled) {
		return nil, false
	}
	return o.StatusChecksEnabled, true
}

// HasStatusChecksEnabled returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasStatusChecksEnabled() bool {
	if o != nil && !IsNil(o.StatusChecksEnabled) {
		return true
	}

	return false
}

// SetStatusChecksEnabled gets a reference to the given ApiCompositeValueDTOBoolean and assigns it to the StatusChecksEnabled field.
func (o *ApiCompositeSourceControlDTO) SetStatusChecksEnabled(v ApiCompositeValueDTOBoolean) {
	o.StatusChecksEnabled = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetToken() ApiCompositeValueDTOString {
	if o == nil || IsNil(o.Token) {
		var ret ApiCompositeValueDTOString
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetTokenOk() (*ApiCompositeValueDTOString, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given ApiCompositeValueDTOString and assigns it to the Token field.
func (o *ApiCompositeSourceControlDTO) SetToken(v ApiCompositeValueDTOString) {
	o.Token = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiCompositeSourceControlDTO) GetUsername() ApiCompositeValueDTOString {
	if o == nil || IsNil(o.Username) {
		var ret ApiCompositeValueDTOString
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCompositeSourceControlDTO) GetUsernameOk() (*ApiCompositeValueDTOString, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiCompositeSourceControlDTO) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given ApiCompositeValueDTOString and assigns it to the Username field.
func (o *ApiCompositeSourceControlDTO) SetUsername(v ApiCompositeValueDTOString) {
	o.Username = &v
}

func (o ApiCompositeSourceControlDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCompositeSourceControlDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseBranch) {
		toSerialize["baseBranch"] = o.BaseBranch
	}
	if !IsNil(o.CommitStatusEnabled) {
		toSerialize["commitStatusEnabled"] = o.CommitStatusEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.PullRequestCommentingEnabled) {
		toSerialize["pullRequestCommentingEnabled"] = o.PullRequestCommentingEnabled
	}
	if !IsNil(o.RemediationPullRequestsEnabled) {
		toSerialize["remediationPullRequestsEnabled"] = o.RemediationPullRequestsEnabled
	}
	if !IsNil(o.RepositoryUrl) {
		toSerialize["repositoryUrl"] = o.RepositoryUrl
	}
	if !IsNil(o.SourceControlEvaluationsEnabled) {
		toSerialize["sourceControlEvaluationsEnabled"] = o.SourceControlEvaluationsEnabled
	}
	if !IsNil(o.SourceControlScanTarget) {
		toSerialize["sourceControlScanTarget"] = o.SourceControlScanTarget
	}
	if !IsNil(o.SshEnabled) {
		toSerialize["sshEnabled"] = o.SshEnabled
	}
	if !IsNil(o.StatusChecksEnabled) {
		toSerialize["statusChecksEnabled"] = o.StatusChecksEnabled
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableApiCompositeSourceControlDTO struct {
	value *ApiCompositeSourceControlDTO
	isSet bool
}

func (v NullableApiCompositeSourceControlDTO) Get() *ApiCompositeSourceControlDTO {
	return v.value
}

func (v *NullableApiCompositeSourceControlDTO) Set(val *ApiCompositeSourceControlDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCompositeSourceControlDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCompositeSourceControlDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCompositeSourceControlDTO(val *ApiCompositeSourceControlDTO) *NullableApiCompositeSourceControlDTO {
	return &NullableApiCompositeSourceControlDTO{value: val, isSet: true}
}

func (v NullableApiCompositeSourceControlDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCompositeSourceControlDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


