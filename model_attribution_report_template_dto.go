/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the AttributionReportTemplateDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributionReportTemplateDTO{}

// AttributionReportTemplateDTO struct for AttributionReportTemplateDTO
type AttributionReportTemplateDTO struct {
	DocumentTitle *string `json:"documentTitle,omitempty"`
	Footer *string `json:"footer,omitempty"`
	Header *string `json:"header,omitempty"`
	Id *string `json:"id,omitempty"`
	IncludeAppendix *bool `json:"includeAppendix,omitempty"`
	IncludeInnerSource *bool `json:"includeInnerSource,omitempty"`
	IncludeSonatypeSpecialLicenses *bool `json:"includeSonatypeSpecialLicenses,omitempty"`
	IncludeStandardLicenseTexts *bool `json:"includeStandardLicenseTexts,omitempty"`
	IncludeTableOfContents *bool `json:"includeTableOfContents,omitempty"`
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
}

// NewAttributionReportTemplateDTO instantiates a new AttributionReportTemplateDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributionReportTemplateDTO() *AttributionReportTemplateDTO {
	this := AttributionReportTemplateDTO{}
	return &this
}

// NewAttributionReportTemplateDTOWithDefaults instantiates a new AttributionReportTemplateDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributionReportTemplateDTOWithDefaults() *AttributionReportTemplateDTO {
	this := AttributionReportTemplateDTO{}
	return &this
}

// GetDocumentTitle returns the DocumentTitle field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetDocumentTitle() string {
	if o == nil || IsNil(o.DocumentTitle) {
		var ret string
		return ret
	}
	return *o.DocumentTitle
}

// GetDocumentTitleOk returns a tuple with the DocumentTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetDocumentTitleOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentTitle) {
		return nil, false
	}
	return o.DocumentTitle, true
}

// HasDocumentTitle returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasDocumentTitle() bool {
	if o != nil && !IsNil(o.DocumentTitle) {
		return true
	}

	return false
}

// SetDocumentTitle gets a reference to the given string and assigns it to the DocumentTitle field.
func (o *AttributionReportTemplateDTO) SetDocumentTitle(v string) {
	o.DocumentTitle = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetFooter() string {
	if o == nil || IsNil(o.Footer) {
		var ret string
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetFooterOk() (*string, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given string and assigns it to the Footer field.
func (o *AttributionReportTemplateDTO) SetFooter(v string) {
	o.Footer = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetHeader() string {
	if o == nil || IsNil(o.Header) {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *AttributionReportTemplateDTO) SetHeader(v string) {
	o.Header = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AttributionReportTemplateDTO) SetId(v string) {
	o.Id = &v
}

// GetIncludeAppendix returns the IncludeAppendix field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetIncludeAppendix() bool {
	if o == nil || IsNil(o.IncludeAppendix) {
		var ret bool
		return ret
	}
	return *o.IncludeAppendix
}

// GetIncludeAppendixOk returns a tuple with the IncludeAppendix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetIncludeAppendixOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAppendix) {
		return nil, false
	}
	return o.IncludeAppendix, true
}

// HasIncludeAppendix returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasIncludeAppendix() bool {
	if o != nil && !IsNil(o.IncludeAppendix) {
		return true
	}

	return false
}

// SetIncludeAppendix gets a reference to the given bool and assigns it to the IncludeAppendix field.
func (o *AttributionReportTemplateDTO) SetIncludeAppendix(v bool) {
	o.IncludeAppendix = &v
}

// GetIncludeInnerSource returns the IncludeInnerSource field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetIncludeInnerSource() bool {
	if o == nil || IsNil(o.IncludeInnerSource) {
		var ret bool
		return ret
	}
	return *o.IncludeInnerSource
}

// GetIncludeInnerSourceOk returns a tuple with the IncludeInnerSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetIncludeInnerSourceOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeInnerSource) {
		return nil, false
	}
	return o.IncludeInnerSource, true
}

// HasIncludeInnerSource returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasIncludeInnerSource() bool {
	if o != nil && !IsNil(o.IncludeInnerSource) {
		return true
	}

	return false
}

// SetIncludeInnerSource gets a reference to the given bool and assigns it to the IncludeInnerSource field.
func (o *AttributionReportTemplateDTO) SetIncludeInnerSource(v bool) {
	o.IncludeInnerSource = &v
}

// GetIncludeSonatypeSpecialLicenses returns the IncludeSonatypeSpecialLicenses field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetIncludeSonatypeSpecialLicenses() bool {
	if o == nil || IsNil(o.IncludeSonatypeSpecialLicenses) {
		var ret bool
		return ret
	}
	return *o.IncludeSonatypeSpecialLicenses
}

// GetIncludeSonatypeSpecialLicensesOk returns a tuple with the IncludeSonatypeSpecialLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetIncludeSonatypeSpecialLicensesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSonatypeSpecialLicenses) {
		return nil, false
	}
	return o.IncludeSonatypeSpecialLicenses, true
}

// HasIncludeSonatypeSpecialLicenses returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasIncludeSonatypeSpecialLicenses() bool {
	if o != nil && !IsNil(o.IncludeSonatypeSpecialLicenses) {
		return true
	}

	return false
}

// SetIncludeSonatypeSpecialLicenses gets a reference to the given bool and assigns it to the IncludeSonatypeSpecialLicenses field.
func (o *AttributionReportTemplateDTO) SetIncludeSonatypeSpecialLicenses(v bool) {
	o.IncludeSonatypeSpecialLicenses = &v
}

// GetIncludeStandardLicenseTexts returns the IncludeStandardLicenseTexts field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetIncludeStandardLicenseTexts() bool {
	if o == nil || IsNil(o.IncludeStandardLicenseTexts) {
		var ret bool
		return ret
	}
	return *o.IncludeStandardLicenseTexts
}

// GetIncludeStandardLicenseTextsOk returns a tuple with the IncludeStandardLicenseTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetIncludeStandardLicenseTextsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeStandardLicenseTexts) {
		return nil, false
	}
	return o.IncludeStandardLicenseTexts, true
}

// HasIncludeStandardLicenseTexts returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasIncludeStandardLicenseTexts() bool {
	if o != nil && !IsNil(o.IncludeStandardLicenseTexts) {
		return true
	}

	return false
}

// SetIncludeStandardLicenseTexts gets a reference to the given bool and assigns it to the IncludeStandardLicenseTexts field.
func (o *AttributionReportTemplateDTO) SetIncludeStandardLicenseTexts(v bool) {
	o.IncludeStandardLicenseTexts = &v
}

// GetIncludeTableOfContents returns the IncludeTableOfContents field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetIncludeTableOfContents() bool {
	if o == nil || IsNil(o.IncludeTableOfContents) {
		var ret bool
		return ret
	}
	return *o.IncludeTableOfContents
}

// GetIncludeTableOfContentsOk returns a tuple with the IncludeTableOfContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetIncludeTableOfContentsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeTableOfContents) {
		return nil, false
	}
	return o.IncludeTableOfContents, true
}

// HasIncludeTableOfContents returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasIncludeTableOfContents() bool {
	if o != nil && !IsNil(o.IncludeTableOfContents) {
		return true
	}

	return false
}

// SetIncludeTableOfContents gets a reference to the given bool and assigns it to the IncludeTableOfContents field.
func (o *AttributionReportTemplateDTO) SetIncludeTableOfContents(v bool) {
	o.IncludeTableOfContents = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.LastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedAt) {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasLastUpdatedAt() bool {
	if o != nil && !IsNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *AttributionReportTemplateDTO) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *AttributionReportTemplateDTO) GetTemplateName() string {
	if o == nil || IsNil(o.TemplateName) {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributionReportTemplateDTO) GetTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateName) {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *AttributionReportTemplateDTO) HasTemplateName() bool {
	if o != nil && !IsNil(o.TemplateName) {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *AttributionReportTemplateDTO) SetTemplateName(v string) {
	o.TemplateName = &v
}

func (o AttributionReportTemplateDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributionReportTemplateDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentTitle) {
		toSerialize["documentTitle"] = o.DocumentTitle
	}
	if !IsNil(o.Footer) {
		toSerialize["footer"] = o.Footer
	}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IncludeAppendix) {
		toSerialize["includeAppendix"] = o.IncludeAppendix
	}
	if !IsNil(o.IncludeInnerSource) {
		toSerialize["includeInnerSource"] = o.IncludeInnerSource
	}
	if !IsNil(o.IncludeSonatypeSpecialLicenses) {
		toSerialize["includeSonatypeSpecialLicenses"] = o.IncludeSonatypeSpecialLicenses
	}
	if !IsNil(o.IncludeStandardLicenseTexts) {
		toSerialize["includeStandardLicenseTexts"] = o.IncludeStandardLicenseTexts
	}
	if !IsNil(o.IncludeTableOfContents) {
		toSerialize["includeTableOfContents"] = o.IncludeTableOfContents
	}
	if !IsNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	if !IsNil(o.TemplateName) {
		toSerialize["templateName"] = o.TemplateName
	}
	return toSerialize, nil
}

type NullableAttributionReportTemplateDTO struct {
	value *AttributionReportTemplateDTO
	isSet bool
}

func (v NullableAttributionReportTemplateDTO) Get() *AttributionReportTemplateDTO {
	return v.value
}

func (v *NullableAttributionReportTemplateDTO) Set(val *AttributionReportTemplateDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributionReportTemplateDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributionReportTemplateDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributionReportTemplateDTO(val *AttributionReportTemplateDTO) *NullableAttributionReportTemplateDTO {
	return &NullableAttributionReportTemplateDTO{value: val, isSet: true}
}

func (v NullableAttributionReportTemplateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributionReportTemplateDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


