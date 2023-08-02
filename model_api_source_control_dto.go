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

// checks if the ApiSourceControlDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSourceControlDTO{}

// ApiSourceControlDTO struct for ApiSourceControlDTO
type ApiSourceControlDTO struct {
	BaseBranch *string `json:"baseBranch,omitempty"`
	CommitStatusEnabled *bool `json:"commitStatusEnabled,omitempty"`
	EnablePullRequests *bool `json:"enablePullRequests,omitempty"`
	EnableStatusChecks *bool `json:"enableStatusChecks,omitempty"`
	Id *string `json:"id,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	Provider *string `json:"provider,omitempty"`
	PullRequestCommentingEnabled *bool `json:"pullRequestCommentingEnabled,omitempty"`
	RemediationPullRequestsEnabled *bool `json:"remediationPullRequestsEnabled,omitempty"`
	RepositoryUrl *string `json:"repositoryUrl,omitempty"`
	SourceControlEvaluationsEnabled *bool `json:"sourceControlEvaluationsEnabled,omitempty"`
	SourceControlScanTarget *string `json:"sourceControlScanTarget,omitempty"`
	SshEnabled *bool `json:"sshEnabled,omitempty"`
	StatusChecksEnabled *bool `json:"statusChecksEnabled,omitempty"`
	Token *string `json:"token,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewApiSourceControlDTO instantiates a new ApiSourceControlDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSourceControlDTO() *ApiSourceControlDTO {
	this := ApiSourceControlDTO{}
	return &this
}

// NewApiSourceControlDTOWithDefaults instantiates a new ApiSourceControlDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSourceControlDTOWithDefaults() *ApiSourceControlDTO {
	this := ApiSourceControlDTO{}
	return &this
}

// GetBaseBranch returns the BaseBranch field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetBaseBranch() string {
	if o == nil || IsNil(o.BaseBranch) {
		var ret string
		return ret
	}
	return *o.BaseBranch
}

// GetBaseBranchOk returns a tuple with the BaseBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetBaseBranchOk() (*string, bool) {
	if o == nil || IsNil(o.BaseBranch) {
		return nil, false
	}
	return o.BaseBranch, true
}

// HasBaseBranch returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasBaseBranch() bool {
	if o != nil && !IsNil(o.BaseBranch) {
		return true
	}

	return false
}

// SetBaseBranch gets a reference to the given string and assigns it to the BaseBranch field.
func (o *ApiSourceControlDTO) SetBaseBranch(v string) {
	o.BaseBranch = &v
}

// GetCommitStatusEnabled returns the CommitStatusEnabled field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetCommitStatusEnabled() bool {
	if o == nil || IsNil(o.CommitStatusEnabled) {
		var ret bool
		return ret
	}
	return *o.CommitStatusEnabled
}

// GetCommitStatusEnabledOk returns a tuple with the CommitStatusEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetCommitStatusEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CommitStatusEnabled) {
		return nil, false
	}
	return o.CommitStatusEnabled, true
}

// HasCommitStatusEnabled returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasCommitStatusEnabled() bool {
	if o != nil && !IsNil(o.CommitStatusEnabled) {
		return true
	}

	return false
}

// SetCommitStatusEnabled gets a reference to the given bool and assigns it to the CommitStatusEnabled field.
func (o *ApiSourceControlDTO) SetCommitStatusEnabled(v bool) {
	o.CommitStatusEnabled = &v
}

// GetEnablePullRequests returns the EnablePullRequests field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetEnablePullRequests() bool {
	if o == nil || IsNil(o.EnablePullRequests) {
		var ret bool
		return ret
	}
	return *o.EnablePullRequests
}

// GetEnablePullRequestsOk returns a tuple with the EnablePullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetEnablePullRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePullRequests) {
		return nil, false
	}
	return o.EnablePullRequests, true
}

// HasEnablePullRequests returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasEnablePullRequests() bool {
	if o != nil && !IsNil(o.EnablePullRequests) {
		return true
	}

	return false
}

// SetEnablePullRequests gets a reference to the given bool and assigns it to the EnablePullRequests field.
func (o *ApiSourceControlDTO) SetEnablePullRequests(v bool) {
	o.EnablePullRequests = &v
}

// GetEnableStatusChecks returns the EnableStatusChecks field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetEnableStatusChecks() bool {
	if o == nil || IsNil(o.EnableStatusChecks) {
		var ret bool
		return ret
	}
	return *o.EnableStatusChecks
}

// GetEnableStatusChecksOk returns a tuple with the EnableStatusChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetEnableStatusChecksOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableStatusChecks) {
		return nil, false
	}
	return o.EnableStatusChecks, true
}

// HasEnableStatusChecks returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasEnableStatusChecks() bool {
	if o != nil && !IsNil(o.EnableStatusChecks) {
		return true
	}

	return false
}

// SetEnableStatusChecks gets a reference to the given bool and assigns it to the EnableStatusChecks field.
func (o *ApiSourceControlDTO) SetEnableStatusChecks(v bool) {
	o.EnableStatusChecks = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiSourceControlDTO) SetId(v string) {
	o.Id = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ApiSourceControlDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ApiSourceControlDTO) SetProvider(v string) {
	o.Provider = &v
}

// GetPullRequestCommentingEnabled returns the PullRequestCommentingEnabled field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetPullRequestCommentingEnabled() bool {
	if o == nil || IsNil(o.PullRequestCommentingEnabled) {
		var ret bool
		return ret
	}
	return *o.PullRequestCommentingEnabled
}

// GetPullRequestCommentingEnabledOk returns a tuple with the PullRequestCommentingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetPullRequestCommentingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PullRequestCommentingEnabled) {
		return nil, false
	}
	return o.PullRequestCommentingEnabled, true
}

// HasPullRequestCommentingEnabled returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasPullRequestCommentingEnabled() bool {
	if o != nil && !IsNil(o.PullRequestCommentingEnabled) {
		return true
	}

	return false
}

// SetPullRequestCommentingEnabled gets a reference to the given bool and assigns it to the PullRequestCommentingEnabled field.
func (o *ApiSourceControlDTO) SetPullRequestCommentingEnabled(v bool) {
	o.PullRequestCommentingEnabled = &v
}

// GetRemediationPullRequestsEnabled returns the RemediationPullRequestsEnabled field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetRemediationPullRequestsEnabled() bool {
	if o == nil || IsNil(o.RemediationPullRequestsEnabled) {
		var ret bool
		return ret
	}
	return *o.RemediationPullRequestsEnabled
}

// GetRemediationPullRequestsEnabledOk returns a tuple with the RemediationPullRequestsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetRemediationPullRequestsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RemediationPullRequestsEnabled) {
		return nil, false
	}
	return o.RemediationPullRequestsEnabled, true
}

// HasRemediationPullRequestsEnabled returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasRemediationPullRequestsEnabled() bool {
	if o != nil && !IsNil(o.RemediationPullRequestsEnabled) {
		return true
	}

	return false
}

// SetRemediationPullRequestsEnabled gets a reference to the given bool and assigns it to the RemediationPullRequestsEnabled field.
func (o *ApiSourceControlDTO) SetRemediationPullRequestsEnabled(v bool) {
	o.RemediationPullRequestsEnabled = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetRepositoryUrl() string {
	if o == nil || IsNil(o.RepositoryUrl) {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryUrl) {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasRepositoryUrl() bool {
	if o != nil && !IsNil(o.RepositoryUrl) {
		return true
	}

	return false
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *ApiSourceControlDTO) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// GetSourceControlEvaluationsEnabled returns the SourceControlEvaluationsEnabled field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetSourceControlEvaluationsEnabled() bool {
	if o == nil || IsNil(o.SourceControlEvaluationsEnabled) {
		var ret bool
		return ret
	}
	return *o.SourceControlEvaluationsEnabled
}

// GetSourceControlEvaluationsEnabledOk returns a tuple with the SourceControlEvaluationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetSourceControlEvaluationsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SourceControlEvaluationsEnabled) {
		return nil, false
	}
	return o.SourceControlEvaluationsEnabled, true
}

// HasSourceControlEvaluationsEnabled returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasSourceControlEvaluationsEnabled() bool {
	if o != nil && !IsNil(o.SourceControlEvaluationsEnabled) {
		return true
	}

	return false
}

// SetSourceControlEvaluationsEnabled gets a reference to the given bool and assigns it to the SourceControlEvaluationsEnabled field.
func (o *ApiSourceControlDTO) SetSourceControlEvaluationsEnabled(v bool) {
	o.SourceControlEvaluationsEnabled = &v
}

// GetSourceControlScanTarget returns the SourceControlScanTarget field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetSourceControlScanTarget() string {
	if o == nil || IsNil(o.SourceControlScanTarget) {
		var ret string
		return ret
	}
	return *o.SourceControlScanTarget
}

// GetSourceControlScanTargetOk returns a tuple with the SourceControlScanTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetSourceControlScanTargetOk() (*string, bool) {
	if o == nil || IsNil(o.SourceControlScanTarget) {
		return nil, false
	}
	return o.SourceControlScanTarget, true
}

// HasSourceControlScanTarget returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasSourceControlScanTarget() bool {
	if o != nil && !IsNil(o.SourceControlScanTarget) {
		return true
	}

	return false
}

// SetSourceControlScanTarget gets a reference to the given string and assigns it to the SourceControlScanTarget field.
func (o *ApiSourceControlDTO) SetSourceControlScanTarget(v string) {
	o.SourceControlScanTarget = &v
}

// GetSshEnabled returns the SshEnabled field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetSshEnabled() bool {
	if o == nil || IsNil(o.SshEnabled) {
		var ret bool
		return ret
	}
	return *o.SshEnabled
}

// GetSshEnabledOk returns a tuple with the SshEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetSshEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SshEnabled) {
		return nil, false
	}
	return o.SshEnabled, true
}

// HasSshEnabled returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasSshEnabled() bool {
	if o != nil && !IsNil(o.SshEnabled) {
		return true
	}

	return false
}

// SetSshEnabled gets a reference to the given bool and assigns it to the SshEnabled field.
func (o *ApiSourceControlDTO) SetSshEnabled(v bool) {
	o.SshEnabled = &v
}

// GetStatusChecksEnabled returns the StatusChecksEnabled field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetStatusChecksEnabled() bool {
	if o == nil || IsNil(o.StatusChecksEnabled) {
		var ret bool
		return ret
	}
	return *o.StatusChecksEnabled
}

// GetStatusChecksEnabledOk returns a tuple with the StatusChecksEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetStatusChecksEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StatusChecksEnabled) {
		return nil, false
	}
	return o.StatusChecksEnabled, true
}

// HasStatusChecksEnabled returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasStatusChecksEnabled() bool {
	if o != nil && !IsNil(o.StatusChecksEnabled) {
		return true
	}

	return false
}

// SetStatusChecksEnabled gets a reference to the given bool and assigns it to the StatusChecksEnabled field.
func (o *ApiSourceControlDTO) SetStatusChecksEnabled(v bool) {
	o.StatusChecksEnabled = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ApiSourceControlDTO) SetToken(v string) {
	o.Token = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiSourceControlDTO) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlDTO) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiSourceControlDTO) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiSourceControlDTO) SetUsername(v string) {
	o.Username = &v
}

func (o ApiSourceControlDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSourceControlDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseBranch) {
		toSerialize["baseBranch"] = o.BaseBranch
	}
	if !IsNil(o.CommitStatusEnabled) {
		toSerialize["commitStatusEnabled"] = o.CommitStatusEnabled
	}
	if !IsNil(o.EnablePullRequests) {
		toSerialize["enablePullRequests"] = o.EnablePullRequests
	}
	if !IsNil(o.EnableStatusChecks) {
		toSerialize["enableStatusChecks"] = o.EnableStatusChecks
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

type NullableApiSourceControlDTO struct {
	value *ApiSourceControlDTO
	isSet bool
}

func (v NullableApiSourceControlDTO) Get() *ApiSourceControlDTO {
	return v.value
}

func (v *NullableApiSourceControlDTO) Set(val *ApiSourceControlDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSourceControlDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSourceControlDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSourceControlDTO(val *ApiSourceControlDTO) *NullableApiSourceControlDTO {
	return &NullableApiSourceControlDTO{value: val, isSet: true}
}

func (v NullableApiSourceControlDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSourceControlDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


