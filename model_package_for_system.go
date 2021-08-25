/*
RHSM-API

API for Red Hat Subscription Management

API version: 1.196.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// PackageForSystem package installed on a system
type PackageForSystem struct {
	Advisories  *[]PackageForSystemAdvisories `json:"advisories,omitempty"`
	Arch        *string                       `json:"arch,omitempty"`
	Epoch       *int32                        `json:"epoch,omitempty"`
	ErrataCount *int32                        `json:"errataCount,omitempty"`
	Name        *string                       `json:"name,omitempty"`
	Release     *string                       `json:"release,omitempty"`
	Version     *string                       `json:"version,omitempty"`
}

// NewPackageForSystem instantiates a new PackageForSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageForSystem() *PackageForSystem {
	this := PackageForSystem{}
	return &this
}

// NewPackageForSystemWithDefaults instantiates a new PackageForSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageForSystemWithDefaults() *PackageForSystem {
	this := PackageForSystem{}
	return &this
}

// GetAdvisories returns the Advisories field value if set, zero value otherwise.
func (o *PackageForSystem) GetAdvisories() []PackageForSystemAdvisories {
	if o == nil || o.Advisories == nil {
		var ret []PackageForSystemAdvisories
		return ret
	}
	return *o.Advisories
}

// GetAdvisoriesOk returns a tuple with the Advisories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageForSystem) GetAdvisoriesOk() (*[]PackageForSystemAdvisories, bool) {
	if o == nil || o.Advisories == nil {
		return nil, false
	}
	return o.Advisories, true
}

// HasAdvisories returns a boolean if a field has been set.
func (o *PackageForSystem) HasAdvisories() bool {
	if o != nil && o.Advisories != nil {
		return true
	}

	return false
}

// SetAdvisories gets a reference to the given []PackageForSystemAdvisories and assigns it to the Advisories field.
func (o *PackageForSystem) SetAdvisories(v []PackageForSystemAdvisories) {
	o.Advisories = &v
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *PackageForSystem) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageForSystem) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *PackageForSystem) HasArch() bool {
	if o != nil && o.Arch != nil {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *PackageForSystem) SetArch(v string) {
	o.Arch = &v
}

// GetEpoch returns the Epoch field value if set, zero value otherwise.
func (o *PackageForSystem) GetEpoch() int32 {
	if o == nil || o.Epoch == nil {
		var ret int32
		return ret
	}
	return *o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageForSystem) GetEpochOk() (*int32, bool) {
	if o == nil || o.Epoch == nil {
		return nil, false
	}
	return o.Epoch, true
}

// HasEpoch returns a boolean if a field has been set.
func (o *PackageForSystem) HasEpoch() bool {
	if o != nil && o.Epoch != nil {
		return true
	}

	return false
}

// SetEpoch gets a reference to the given int32 and assigns it to the Epoch field.
func (o *PackageForSystem) SetEpoch(v int32) {
	o.Epoch = &v
}

// GetErrataCount returns the ErrataCount field value if set, zero value otherwise.
func (o *PackageForSystem) GetErrataCount() int32 {
	if o == nil || o.ErrataCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrataCount
}

// GetErrataCountOk returns a tuple with the ErrataCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageForSystem) GetErrataCountOk() (*int32, bool) {
	if o == nil || o.ErrataCount == nil {
		return nil, false
	}
	return o.ErrataCount, true
}

// HasErrataCount returns a boolean if a field has been set.
func (o *PackageForSystem) HasErrataCount() bool {
	if o != nil && o.ErrataCount != nil {
		return true
	}

	return false
}

// SetErrataCount gets a reference to the given int32 and assigns it to the ErrataCount field.
func (o *PackageForSystem) SetErrataCount(v int32) {
	o.ErrataCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PackageForSystem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageForSystem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PackageForSystem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PackageForSystem) SetName(v string) {
	o.Name = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *PackageForSystem) GetRelease() string {
	if o == nil || o.Release == nil {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageForSystem) GetReleaseOk() (*string, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *PackageForSystem) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *PackageForSystem) SetRelease(v string) {
	o.Release = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PackageForSystem) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageForSystem) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PackageForSystem) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PackageForSystem) SetVersion(v string) {
	o.Version = &v
}

func (o PackageForSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Advisories != nil {
		toSerialize["advisories"] = o.Advisories
	}
	if o.Arch != nil {
		toSerialize["arch"] = o.Arch
	}
	if o.Epoch != nil {
		toSerialize["epoch"] = o.Epoch
	}
	if o.ErrataCount != nil {
		toSerialize["errataCount"] = o.ErrataCount
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Release != nil {
		toSerialize["release"] = o.Release
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePackageForSystem struct {
	value *PackageForSystem
	isSet bool
}

func (v NullablePackageForSystem) Get() *PackageForSystem {
	return v.value
}

func (v *NullablePackageForSystem) Set(val *PackageForSystem) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageForSystem) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageForSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageForSystem(val *PackageForSystem) *NullablePackageForSystem {
	return &NullablePackageForSystem{value: val, isSet: true}
}

func (v NullablePackageForSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageForSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
