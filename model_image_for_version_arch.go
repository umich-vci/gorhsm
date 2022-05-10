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

// ImageForVersionArch Image Details for provided version and architecture.
type ImageForVersionArch struct {
	Arch     *string `json:"arch,omitempty"`
	Checksum *string `json:"checksum,omitempty"`
	// Date represents the date format used for API returns
	DatePublished *string `json:"datePublished,omitempty"`
	DownloadHref  *string `json:"downloadHref,omitempty"`
	Filename      *string `json:"filename,omitempty"`
	ImageName     *string `json:"imageName,omitempty"`
}

// NewImageForVersionArch instantiates a new ImageForVersionArch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageForVersionArch() *ImageForVersionArch {
	this := ImageForVersionArch{}
	return &this
}

// NewImageForVersionArchWithDefaults instantiates a new ImageForVersionArch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageForVersionArchWithDefaults() *ImageForVersionArch {
	this := ImageForVersionArch{}
	return &this
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *ImageForVersionArch) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageForVersionArch) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *ImageForVersionArch) HasArch() bool {
	if o != nil && o.Arch != nil {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *ImageForVersionArch) SetArch(v string) {
	o.Arch = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *ImageForVersionArch) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageForVersionArch) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *ImageForVersionArch) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *ImageForVersionArch) SetChecksum(v string) {
	o.Checksum = &v
}

// GetDatePublished returns the DatePublished field value if set, zero value otherwise.
func (o *ImageForVersionArch) GetDatePublished() string {
	if o == nil || o.DatePublished == nil {
		var ret string
		return ret
	}
	return *o.DatePublished
}

// GetDatePublishedOk returns a tuple with the DatePublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageForVersionArch) GetDatePublishedOk() (*string, bool) {
	if o == nil || o.DatePublished == nil {
		return nil, false
	}
	return o.DatePublished, true
}

// HasDatePublished returns a boolean if a field has been set.
func (o *ImageForVersionArch) HasDatePublished() bool {
	if o != nil && o.DatePublished != nil {
		return true
	}

	return false
}

// SetDatePublished gets a reference to the given string and assigns it to the DatePublished field.
func (o *ImageForVersionArch) SetDatePublished(v string) {
	o.DatePublished = &v
}

// GetDownloadHref returns the DownloadHref field value if set, zero value otherwise.
func (o *ImageForVersionArch) GetDownloadHref() string {
	if o == nil || o.DownloadHref == nil {
		var ret string
		return ret
	}
	return *o.DownloadHref
}

// GetDownloadHrefOk returns a tuple with the DownloadHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageForVersionArch) GetDownloadHrefOk() (*string, bool) {
	if o == nil || o.DownloadHref == nil {
		return nil, false
	}
	return o.DownloadHref, true
}

// HasDownloadHref returns a boolean if a field has been set.
func (o *ImageForVersionArch) HasDownloadHref() bool {
	if o != nil && o.DownloadHref != nil {
		return true
	}

	return false
}

// SetDownloadHref gets a reference to the given string and assigns it to the DownloadHref field.
func (o *ImageForVersionArch) SetDownloadHref(v string) {
	o.DownloadHref = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ImageForVersionArch) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageForVersionArch) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ImageForVersionArch) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ImageForVersionArch) SetFilename(v string) {
	o.Filename = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *ImageForVersionArch) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageForVersionArch) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *ImageForVersionArch) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *ImageForVersionArch) SetImageName(v string) {
	o.ImageName = &v
}

func (o ImageForVersionArch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Arch != nil {
		toSerialize["arch"] = o.Arch
	}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if o.DatePublished != nil {
		toSerialize["datePublished"] = o.DatePublished
	}
	if o.DownloadHref != nil {
		toSerialize["downloadHref"] = o.DownloadHref
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.ImageName != nil {
		toSerialize["imageName"] = o.ImageName
	}
	return json.Marshal(toSerialize)
}

type NullableImageForVersionArch struct {
	value *ImageForVersionArch
	isSet bool
}

func (v NullableImageForVersionArch) Get() *ImageForVersionArch {
	return v.value
}

func (v *NullableImageForVersionArch) Set(val *ImageForVersionArch) {
	v.value = val
	v.isSet = true
}

func (v NullableImageForVersionArch) IsSet() bool {
	return v.isSet
}

func (v *NullableImageForVersionArch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageForVersionArch(val *ImageForVersionArch) *NullableImageForVersionArch {
	return &NullableImageForVersionArch{value: val, isSet: true}
}

func (v NullableImageForVersionArch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageForVersionArch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}