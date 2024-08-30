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

// checks if the SearchResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResultDTO{}

// SearchResultDTO struct for SearchResultDTO
type SearchResultDTO struct {
	GroupingByDTOS []GroupingByDTO `json:"groupingByDTOS,omitempty"`
	IsExactTotalNumberOfHits *bool `json:"isExactTotalNumberOfHits,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	SearchQuery *string `json:"searchQuery,omitempty"`
	TotalNumberOfHits *int32 `json:"totalNumberOfHits,omitempty"`
}

// NewSearchResultDTO instantiates a new SearchResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResultDTO() *SearchResultDTO {
	this := SearchResultDTO{}
	return &this
}

// NewSearchResultDTOWithDefaults instantiates a new SearchResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultDTOWithDefaults() *SearchResultDTO {
	this := SearchResultDTO{}
	return &this
}

// GetGroupingByDTOS returns the GroupingByDTOS field value if set, zero value otherwise.
func (o *SearchResultDTO) GetGroupingByDTOS() []GroupingByDTO {
	if o == nil || IsNil(o.GroupingByDTOS) {
		var ret []GroupingByDTO
		return ret
	}
	return o.GroupingByDTOS
}

// GetGroupingByDTOSOk returns a tuple with the GroupingByDTOS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultDTO) GetGroupingByDTOSOk() ([]GroupingByDTO, bool) {
	if o == nil || IsNil(o.GroupingByDTOS) {
		return nil, false
	}
	return o.GroupingByDTOS, true
}

// HasGroupingByDTOS returns a boolean if a field has been set.
func (o *SearchResultDTO) HasGroupingByDTOS() bool {
	if o != nil && !IsNil(o.GroupingByDTOS) {
		return true
	}

	return false
}

// SetGroupingByDTOS gets a reference to the given []GroupingByDTO and assigns it to the GroupingByDTOS field.
func (o *SearchResultDTO) SetGroupingByDTOS(v []GroupingByDTO) {
	o.GroupingByDTOS = v
}

// GetIsExactTotalNumberOfHits returns the IsExactTotalNumberOfHits field value if set, zero value otherwise.
func (o *SearchResultDTO) GetIsExactTotalNumberOfHits() bool {
	if o == nil || IsNil(o.IsExactTotalNumberOfHits) {
		var ret bool
		return ret
	}
	return *o.IsExactTotalNumberOfHits
}

// GetIsExactTotalNumberOfHitsOk returns a tuple with the IsExactTotalNumberOfHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultDTO) GetIsExactTotalNumberOfHitsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExactTotalNumberOfHits) {
		return nil, false
	}
	return o.IsExactTotalNumberOfHits, true
}

// HasIsExactTotalNumberOfHits returns a boolean if a field has been set.
func (o *SearchResultDTO) HasIsExactTotalNumberOfHits() bool {
	if o != nil && !IsNil(o.IsExactTotalNumberOfHits) {
		return true
	}

	return false
}

// SetIsExactTotalNumberOfHits gets a reference to the given bool and assigns it to the IsExactTotalNumberOfHits field.
func (o *SearchResultDTO) SetIsExactTotalNumberOfHits(v bool) {
	o.IsExactTotalNumberOfHits = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *SearchResultDTO) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultDTO) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *SearchResultDTO) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *SearchResultDTO) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *SearchResultDTO) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultDTO) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *SearchResultDTO) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *SearchResultDTO) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSearchQuery returns the SearchQuery field value if set, zero value otherwise.
func (o *SearchResultDTO) GetSearchQuery() string {
	if o == nil || IsNil(o.SearchQuery) {
		var ret string
		return ret
	}
	return *o.SearchQuery
}

// GetSearchQueryOk returns a tuple with the SearchQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultDTO) GetSearchQueryOk() (*string, bool) {
	if o == nil || IsNil(o.SearchQuery) {
		return nil, false
	}
	return o.SearchQuery, true
}

// HasSearchQuery returns a boolean if a field has been set.
func (o *SearchResultDTO) HasSearchQuery() bool {
	if o != nil && !IsNil(o.SearchQuery) {
		return true
	}

	return false
}

// SetSearchQuery gets a reference to the given string and assigns it to the SearchQuery field.
func (o *SearchResultDTO) SetSearchQuery(v string) {
	o.SearchQuery = &v
}

// GetTotalNumberOfHits returns the TotalNumberOfHits field value if set, zero value otherwise.
func (o *SearchResultDTO) GetTotalNumberOfHits() int32 {
	if o == nil || IsNil(o.TotalNumberOfHits) {
		var ret int32
		return ret
	}
	return *o.TotalNumberOfHits
}

// GetTotalNumberOfHitsOk returns a tuple with the TotalNumberOfHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultDTO) GetTotalNumberOfHitsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNumberOfHits) {
		return nil, false
	}
	return o.TotalNumberOfHits, true
}

// HasTotalNumberOfHits returns a boolean if a field has been set.
func (o *SearchResultDTO) HasTotalNumberOfHits() bool {
	if o != nil && !IsNil(o.TotalNumberOfHits) {
		return true
	}

	return false
}

// SetTotalNumberOfHits gets a reference to the given int32 and assigns it to the TotalNumberOfHits field.
func (o *SearchResultDTO) SetTotalNumberOfHits(v int32) {
	o.TotalNumberOfHits = &v
}

func (o SearchResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupingByDTOS) {
		toSerialize["groupingByDTOS"] = o.GroupingByDTOS
	}
	if !IsNil(o.IsExactTotalNumberOfHits) {
		toSerialize["isExactTotalNumberOfHits"] = o.IsExactTotalNumberOfHits
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.SearchQuery) {
		toSerialize["searchQuery"] = o.SearchQuery
	}
	if !IsNil(o.TotalNumberOfHits) {
		toSerialize["totalNumberOfHits"] = o.TotalNumberOfHits
	}
	return toSerialize, nil
}

type NullableSearchResultDTO struct {
	value *SearchResultDTO
	isSet bool
}

func (v NullableSearchResultDTO) Get() *SearchResultDTO {
	return v.value
}

func (v *NullableSearchResultDTO) Set(val *SearchResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResultDTO(val *SearchResultDTO) *NullableSearchResultDTO {
	return &NullableSearchResultDTO{value: val, isSet: true}
}

func (v NullableSearchResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


