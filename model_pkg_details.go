/*
RHSM-API

API for Red Hat Subscription Management

API version: 1.300.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// PkgDetails struct for PkgDetails
type PkgDetails struct {
	Arch *string `json:"arch,omitempty"`
	// Date represents the date format used for API returns
	BuildDate   *string  `json:"buildDate,omitempty"`
	BuildHost   *string  `json:"buildHost,omitempty"`
	Checksum    *string  `json:"checksum,omitempty"`
	ContentSets []string `json:"contentSets,omitempty"`
	Description *string  `json:"description,omitempty"`
	Epoch       *string  `json:"epoch,omitempty"`
	Group       *string  `json:"group,omitempty"`
	Href        *string  `json:"href,omitempty"`
	License     *string  `json:"license,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Release     *string  `json:"release,omitempty"`
	Size        *int32   `json:"size,omitempty"`
	Summary     *string  `json:"summary,omitempty"`
	Version     *string  `json:"version,omitempty"`
}

// NewPkgDetails instantiates a new PkgDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgDetails() *PkgDetails {
	this := PkgDetails{}
	return &this
}

// NewPkgDetailsWithDefaults instantiates a new PkgDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgDetailsWithDefaults() *PkgDetails {
	this := PkgDetails{}
	return &this
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *PkgDetails) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *PkgDetails) HasArch() bool {
	if o != nil && o.Arch != nil {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *PkgDetails) SetArch(v string) {
	o.Arch = &v
}

// GetBuildDate returns the BuildDate field value if set, zero value otherwise.
func (o *PkgDetails) GetBuildDate() string {
	if o == nil || o.BuildDate == nil {
		var ret string
		return ret
	}
	return *o.BuildDate
}

// GetBuildDateOk returns a tuple with the BuildDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetBuildDateOk() (*string, bool) {
	if o == nil || o.BuildDate == nil {
		return nil, false
	}
	return o.BuildDate, true
}

// HasBuildDate returns a boolean if a field has been set.
func (o *PkgDetails) HasBuildDate() bool {
	if o != nil && o.BuildDate != nil {
		return true
	}

	return false
}

// SetBuildDate gets a reference to the given string and assigns it to the BuildDate field.
func (o *PkgDetails) SetBuildDate(v string) {
	o.BuildDate = &v
}

// GetBuildHost returns the BuildHost field value if set, zero value otherwise.
func (o *PkgDetails) GetBuildHost() string {
	if o == nil || o.BuildHost == nil {
		var ret string
		return ret
	}
	return *o.BuildHost
}

// GetBuildHostOk returns a tuple with the BuildHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetBuildHostOk() (*string, bool) {
	if o == nil || o.BuildHost == nil {
		return nil, false
	}
	return o.BuildHost, true
}

// HasBuildHost returns a boolean if a field has been set.
func (o *PkgDetails) HasBuildHost() bool {
	if o != nil && o.BuildHost != nil {
		return true
	}

	return false
}

// SetBuildHost gets a reference to the given string and assigns it to the BuildHost field.
func (o *PkgDetails) SetBuildHost(v string) {
	o.BuildHost = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *PkgDetails) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *PkgDetails) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *PkgDetails) SetChecksum(v string) {
	o.Checksum = &v
}

// GetContentSets returns the ContentSets field value if set, zero value otherwise.
func (o *PkgDetails) GetContentSets() []string {
	if o == nil || o.ContentSets == nil {
		var ret []string
		return ret
	}
	return o.ContentSets
}

// GetContentSetsOk returns a tuple with the ContentSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetContentSetsOk() ([]string, bool) {
	if o == nil || o.ContentSets == nil {
		return nil, false
	}
	return o.ContentSets, true
}

// HasContentSets returns a boolean if a field has been set.
func (o *PkgDetails) HasContentSets() bool {
	if o != nil && o.ContentSets != nil {
		return true
	}

	return false
}

// SetContentSets gets a reference to the given []string and assigns it to the ContentSets field.
func (o *PkgDetails) SetContentSets(v []string) {
	o.ContentSets = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PkgDetails) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PkgDetails) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PkgDetails) SetDescription(v string) {
	o.Description = &v
}

// GetEpoch returns the Epoch field value if set, zero value otherwise.
func (o *PkgDetails) GetEpoch() string {
	if o == nil || o.Epoch == nil {
		var ret string
		return ret
	}
	return *o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetEpochOk() (*string, bool) {
	if o == nil || o.Epoch == nil {
		return nil, false
	}
	return o.Epoch, true
}

// HasEpoch returns a boolean if a field has been set.
func (o *PkgDetails) HasEpoch() bool {
	if o != nil && o.Epoch != nil {
		return true
	}

	return false
}

// SetEpoch gets a reference to the given string and assigns it to the Epoch field.
func (o *PkgDetails) SetEpoch(v string) {
	o.Epoch = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PkgDetails) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PkgDetails) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *PkgDetails) SetGroup(v string) {
	o.Group = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PkgDetails) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PkgDetails) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PkgDetails) SetHref(v string) {
	o.Href = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *PkgDetails) GetLicense() string {
	if o == nil || o.License == nil {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetLicenseOk() (*string, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *PkgDetails) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *PkgDetails) SetLicense(v string) {
	o.License = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PkgDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PkgDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PkgDetails) SetName(v string) {
	o.Name = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *PkgDetails) GetRelease() string {
	if o == nil || o.Release == nil {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetReleaseOk() (*string, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *PkgDetails) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *PkgDetails) SetRelease(v string) {
	o.Release = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PkgDetails) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PkgDetails) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PkgDetails) SetSize(v int32) {
	o.Size = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PkgDetails) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PkgDetails) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PkgDetails) SetSummary(v string) {
	o.Summary = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PkgDetails) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgDetails) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PkgDetails) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PkgDetails) SetVersion(v string) {
	o.Version = &v
}

func (o PkgDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Arch != nil {
		toSerialize["arch"] = o.Arch
	}
	if o.BuildDate != nil {
		toSerialize["buildDate"] = o.BuildDate
	}
	if o.BuildHost != nil {
		toSerialize["buildHost"] = o.BuildHost
	}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if o.ContentSets != nil {
		toSerialize["contentSets"] = o.ContentSets
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Epoch != nil {
		toSerialize["epoch"] = o.Epoch
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Release != nil {
		toSerialize["release"] = o.Release
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePkgDetails struct {
	value *PkgDetails
	isSet bool
}

func (v NullablePkgDetails) Get() *PkgDetails {
	return v.value
}

func (v *NullablePkgDetails) Set(val *PkgDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgDetails(val *PkgDetails) *NullablePkgDetails {
	return &NullablePkgDetails{value: val, isSet: true}
}

func (v NullablePkgDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
