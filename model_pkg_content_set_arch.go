/*
RHSM-API

API for Red Hat Subscription Management

API version: 1.313.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// checks if the PkgContentSetArch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PkgContentSetArch{}

// PkgContentSetArch struct for PkgContentSetArch
type PkgContentSetArch struct {
	Arch *string `json:"arch,omitempty"`
	// Date represents the date format used for API returns
	BuildDate *string `json:"buildDate,omitempty"`
	BuildHost *string `json:"buildHost,omitempty"`
	Checksum *string `json:"checksum,omitempty"`
	ContentSets []string `json:"contentSets,omitempty"`
	Description *string `json:"description,omitempty"`
	DownloadHref *string `json:"downloadHref,omitempty"`
	Epoch *string `json:"epoch,omitempty"`
	Group *string `json:"group,omitempty"`
	Href *string `json:"href,omitempty"`
	License *string `json:"license,omitempty"`
	Name *string `json:"name,omitempty"`
	Release *string `json:"release,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewPkgContentSetArch instantiates a new PkgContentSetArch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgContentSetArch() *PkgContentSetArch {
	this := PkgContentSetArch{}
	return &this
}

// NewPkgContentSetArchWithDefaults instantiates a new PkgContentSetArch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgContentSetArchWithDefaults() *PkgContentSetArch {
	this := PkgContentSetArch{}
	return &this
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetArch() string {
	if o == nil || IsNil(o.Arch) {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetArchOk() (*string, bool) {
	if o == nil || IsNil(o.Arch) {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasArch() bool {
	if o != nil && !IsNil(o.Arch) {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *PkgContentSetArch) SetArch(v string) {
	o.Arch = &v
}

// GetBuildDate returns the BuildDate field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetBuildDate() string {
	if o == nil || IsNil(o.BuildDate) {
		var ret string
		return ret
	}
	return *o.BuildDate
}

// GetBuildDateOk returns a tuple with the BuildDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetBuildDateOk() (*string, bool) {
	if o == nil || IsNil(o.BuildDate) {
		return nil, false
	}
	return o.BuildDate, true
}

// HasBuildDate returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasBuildDate() bool {
	if o != nil && !IsNil(o.BuildDate) {
		return true
	}

	return false
}

// SetBuildDate gets a reference to the given string and assigns it to the BuildDate field.
func (o *PkgContentSetArch) SetBuildDate(v string) {
	o.BuildDate = &v
}

// GetBuildHost returns the BuildHost field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetBuildHost() string {
	if o == nil || IsNil(o.BuildHost) {
		var ret string
		return ret
	}
	return *o.BuildHost
}

// GetBuildHostOk returns a tuple with the BuildHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetBuildHostOk() (*string, bool) {
	if o == nil || IsNil(o.BuildHost) {
		return nil, false
	}
	return o.BuildHost, true
}

// HasBuildHost returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasBuildHost() bool {
	if o != nil && !IsNil(o.BuildHost) {
		return true
	}

	return false
}

// SetBuildHost gets a reference to the given string and assigns it to the BuildHost field.
func (o *PkgContentSetArch) SetBuildHost(v string) {
	o.BuildHost = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *PkgContentSetArch) SetChecksum(v string) {
	o.Checksum = &v
}

// GetContentSets returns the ContentSets field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetContentSets() []string {
	if o == nil || IsNil(o.ContentSets) {
		var ret []string
		return ret
	}
	return o.ContentSets
}

// GetContentSetsOk returns a tuple with the ContentSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetContentSetsOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentSets) {
		return nil, false
	}
	return o.ContentSets, true
}

// HasContentSets returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasContentSets() bool {
	if o != nil && !IsNil(o.ContentSets) {
		return true
	}

	return false
}

// SetContentSets gets a reference to the given []string and assigns it to the ContentSets field.
func (o *PkgContentSetArch) SetContentSets(v []string) {
	o.ContentSets = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PkgContentSetArch) SetDescription(v string) {
	o.Description = &v
}

// GetDownloadHref returns the DownloadHref field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetDownloadHref() string {
	if o == nil || IsNil(o.DownloadHref) {
		var ret string
		return ret
	}
	return *o.DownloadHref
}

// GetDownloadHrefOk returns a tuple with the DownloadHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetDownloadHrefOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadHref) {
		return nil, false
	}
	return o.DownloadHref, true
}

// HasDownloadHref returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasDownloadHref() bool {
	if o != nil && !IsNil(o.DownloadHref) {
		return true
	}

	return false
}

// SetDownloadHref gets a reference to the given string and assigns it to the DownloadHref field.
func (o *PkgContentSetArch) SetDownloadHref(v string) {
	o.DownloadHref = &v
}

// GetEpoch returns the Epoch field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetEpoch() string {
	if o == nil || IsNil(o.Epoch) {
		var ret string
		return ret
	}
	return *o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetEpochOk() (*string, bool) {
	if o == nil || IsNil(o.Epoch) {
		return nil, false
	}
	return o.Epoch, true
}

// HasEpoch returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasEpoch() bool {
	if o != nil && !IsNil(o.Epoch) {
		return true
	}

	return false
}

// SetEpoch gets a reference to the given string and assigns it to the Epoch field.
func (o *PkgContentSetArch) SetEpoch(v string) {
	o.Epoch = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *PkgContentSetArch) SetGroup(v string) {
	o.Group = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PkgContentSetArch) SetHref(v string) {
	o.Href = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetLicense() string {
	if o == nil || IsNil(o.License) {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetLicenseOk() (*string, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *PkgContentSetArch) SetLicense(v string) {
	o.License = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PkgContentSetArch) SetName(v string) {
	o.Name = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetRelease() string {
	if o == nil || IsNil(o.Release) {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *PkgContentSetArch) SetRelease(v string) {
	o.Release = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PkgContentSetArch) SetSize(v int32) {
	o.Size = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PkgContentSetArch) SetSummary(v string) {
	o.Summary = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PkgContentSetArch) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgContentSetArch) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PkgContentSetArch) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PkgContentSetArch) SetVersion(v string) {
	o.Version = &v
}

func (o PkgContentSetArch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PkgContentSetArch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Arch) {
		toSerialize["arch"] = o.Arch
	}
	if !IsNil(o.BuildDate) {
		toSerialize["buildDate"] = o.BuildDate
	}
	if !IsNil(o.BuildHost) {
		toSerialize["buildHost"] = o.BuildHost
	}
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	if !IsNil(o.ContentSets) {
		toSerialize["contentSets"] = o.ContentSets
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DownloadHref) {
		toSerialize["downloadHref"] = o.DownloadHref
	}
	if !IsNil(o.Epoch) {
		toSerialize["epoch"] = o.Epoch
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullablePkgContentSetArch struct {
	value *PkgContentSetArch
	isSet bool
}

func (v NullablePkgContentSetArch) Get() *PkgContentSetArch {
	return v.value
}

func (v *NullablePkgContentSetArch) Set(val *PkgContentSetArch) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgContentSetArch) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgContentSetArch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgContentSetArch(val *PkgContentSetArch) *NullablePkgContentSetArch {
	return &NullablePkgContentSetArch{value: val, isSet: true}
}

func (v NullablePkgContentSetArch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgContentSetArch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


