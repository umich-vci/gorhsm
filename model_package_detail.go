/*
RHSM-API

API for Red Hat Subscription Management

API version: 1.264.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// PackageDetail PackageDetail wraps an errata package and adds a RefURL
type PackageDetail struct {
	Arch        *string  `json:"arch,omitempty"`
	Checksum    *string  `json:"checksum,omitempty"`
	ContentSets []string `json:"contentSets,omitempty"`
	DetailsUrl  *string  `json:"details_url,omitempty"`
	Epoch       *int32   `json:"epoch,omitempty"`
	Filename    *string  `json:"filename,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Release     *string  `json:"release,omitempty"`
	RepoTags    []string `json:"repoTags,omitempty"`
	Summary     *string  `json:"summary,omitempty"`
	Version     *string  `json:"version,omitempty"`
}

// NewPackageDetail instantiates a new PackageDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageDetail() *PackageDetail {
	this := PackageDetail{}
	return &this
}

// NewPackageDetailWithDefaults instantiates a new PackageDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageDetailWithDefaults() *PackageDetail {
	this := PackageDetail{}
	return &this
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *PackageDetail) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *PackageDetail) HasArch() bool {
	if o != nil && o.Arch != nil {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *PackageDetail) SetArch(v string) {
	o.Arch = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *PackageDetail) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *PackageDetail) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *PackageDetail) SetChecksum(v string) {
	o.Checksum = &v
}

// GetContentSets returns the ContentSets field value if set, zero value otherwise.
func (o *PackageDetail) GetContentSets() []string {
	if o == nil || o.ContentSets == nil {
		var ret []string
		return ret
	}
	return o.ContentSets
}

// GetContentSetsOk returns a tuple with the ContentSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetContentSetsOk() ([]string, bool) {
	if o == nil || o.ContentSets == nil {
		return nil, false
	}
	return o.ContentSets, true
}

// HasContentSets returns a boolean if a field has been set.
func (o *PackageDetail) HasContentSets() bool {
	if o != nil && o.ContentSets != nil {
		return true
	}

	return false
}

// SetContentSets gets a reference to the given []string and assigns it to the ContentSets field.
func (o *PackageDetail) SetContentSets(v []string) {
	o.ContentSets = v
}

// GetDetailsUrl returns the DetailsUrl field value if set, zero value otherwise.
func (o *PackageDetail) GetDetailsUrl() string {
	if o == nil || o.DetailsUrl == nil {
		var ret string
		return ret
	}
	return *o.DetailsUrl
}

// GetDetailsUrlOk returns a tuple with the DetailsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetDetailsUrlOk() (*string, bool) {
	if o == nil || o.DetailsUrl == nil {
		return nil, false
	}
	return o.DetailsUrl, true
}

// HasDetailsUrl returns a boolean if a field has been set.
func (o *PackageDetail) HasDetailsUrl() bool {
	if o != nil && o.DetailsUrl != nil {
		return true
	}

	return false
}

// SetDetailsUrl gets a reference to the given string and assigns it to the DetailsUrl field.
func (o *PackageDetail) SetDetailsUrl(v string) {
	o.DetailsUrl = &v
}

// GetEpoch returns the Epoch field value if set, zero value otherwise.
func (o *PackageDetail) GetEpoch() int32 {
	if o == nil || o.Epoch == nil {
		var ret int32
		return ret
	}
	return *o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetEpochOk() (*int32, bool) {
	if o == nil || o.Epoch == nil {
		return nil, false
	}
	return o.Epoch, true
}

// HasEpoch returns a boolean if a field has been set.
func (o *PackageDetail) HasEpoch() bool {
	if o != nil && o.Epoch != nil {
		return true
	}

	return false
}

// SetEpoch gets a reference to the given int32 and assigns it to the Epoch field.
func (o *PackageDetail) SetEpoch(v int32) {
	o.Epoch = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *PackageDetail) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *PackageDetail) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *PackageDetail) SetFilename(v string) {
	o.Filename = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PackageDetail) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PackageDetail) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PackageDetail) SetName(v string) {
	o.Name = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *PackageDetail) GetRelease() string {
	if o == nil || o.Release == nil {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetReleaseOk() (*string, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *PackageDetail) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *PackageDetail) SetRelease(v string) {
	o.Release = &v
}

// GetRepoTags returns the RepoTags field value if set, zero value otherwise.
func (o *PackageDetail) GetRepoTags() []string {
	if o == nil || o.RepoTags == nil {
		var ret []string
		return ret
	}
	return o.RepoTags
}

// GetRepoTagsOk returns a tuple with the RepoTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetRepoTagsOk() ([]string, bool) {
	if o == nil || o.RepoTags == nil {
		return nil, false
	}
	return o.RepoTags, true
}

// HasRepoTags returns a boolean if a field has been set.
func (o *PackageDetail) HasRepoTags() bool {
	if o != nil && o.RepoTags != nil {
		return true
	}

	return false
}

// SetRepoTags gets a reference to the given []string and assigns it to the RepoTags field.
func (o *PackageDetail) SetRepoTags(v []string) {
	o.RepoTags = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PackageDetail) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PackageDetail) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PackageDetail) SetSummary(v string) {
	o.Summary = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PackageDetail) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageDetail) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PackageDetail) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PackageDetail) SetVersion(v string) {
	o.Version = &v
}

func (o PackageDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Arch != nil {
		toSerialize["arch"] = o.Arch
	}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if o.ContentSets != nil {
		toSerialize["contentSets"] = o.ContentSets
	}
	if o.DetailsUrl != nil {
		toSerialize["details_url"] = o.DetailsUrl
	}
	if o.Epoch != nil {
		toSerialize["epoch"] = o.Epoch
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Release != nil {
		toSerialize["release"] = o.Release
	}
	if o.RepoTags != nil {
		toSerialize["repoTags"] = o.RepoTags
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePackageDetail struct {
	value *PackageDetail
	isSet bool
}

func (v NullablePackageDetail) Get() *PackageDetail {
	return v.value
}

func (v *NullablePackageDetail) Set(val *PackageDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageDetail(val *PackageDetail) *NullablePackageDetail {
	return &NullablePackageDetail{value: val, isSet: true}
}

func (v NullablePackageDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
