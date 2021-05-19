/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.177.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// MyErratum MyErratum contains erratum information that affects at least one system
type MyErratum struct {
	AdvisoryId          *string `json:"advisoryId,omitempty"`
	AffectedSystemCount *int32  `json:"affectedSystemCount,omitempty"`
	Details             *string `json:"details,omitempty"`
	// Date represents the date format used for API returns
	PublishDate *string `json:"publishDate,omitempty"`
	Synopsis    *string `json:"synopsis,omitempty"`
	Systems     *string `json:"systems,omitempty"`
	Type        *string `json:"type,omitempty"`
}

// NewMyErratum instantiates a new MyErratum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyErratum() *MyErratum {
	this := MyErratum{}
	return &this
}

// NewMyErratumWithDefaults instantiates a new MyErratum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyErratumWithDefaults() *MyErratum {
	this := MyErratum{}
	return &this
}

// GetAdvisoryId returns the AdvisoryId field value if set, zero value otherwise.
func (o *MyErratum) GetAdvisoryId() string {
	if o == nil || o.AdvisoryId == nil {
		var ret string
		return ret
	}
	return *o.AdvisoryId
}

// GetAdvisoryIdOk returns a tuple with the AdvisoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyErratum) GetAdvisoryIdOk() (*string, bool) {
	if o == nil || o.AdvisoryId == nil {
		return nil, false
	}
	return o.AdvisoryId, true
}

// HasAdvisoryId returns a boolean if a field has been set.
func (o *MyErratum) HasAdvisoryId() bool {
	if o != nil && o.AdvisoryId != nil {
		return true
	}

	return false
}

// SetAdvisoryId gets a reference to the given string and assigns it to the AdvisoryId field.
func (o *MyErratum) SetAdvisoryId(v string) {
	o.AdvisoryId = &v
}

// GetAffectedSystemCount returns the AffectedSystemCount field value if set, zero value otherwise.
func (o *MyErratum) GetAffectedSystemCount() int32 {
	if o == nil || o.AffectedSystemCount == nil {
		var ret int32
		return ret
	}
	return *o.AffectedSystemCount
}

// GetAffectedSystemCountOk returns a tuple with the AffectedSystemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyErratum) GetAffectedSystemCountOk() (*int32, bool) {
	if o == nil || o.AffectedSystemCount == nil {
		return nil, false
	}
	return o.AffectedSystemCount, true
}

// HasAffectedSystemCount returns a boolean if a field has been set.
func (o *MyErratum) HasAffectedSystemCount() bool {
	if o != nil && o.AffectedSystemCount != nil {
		return true
	}

	return false
}

// SetAffectedSystemCount gets a reference to the given int32 and assigns it to the AffectedSystemCount field.
func (o *MyErratum) SetAffectedSystemCount(v int32) {
	o.AffectedSystemCount = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MyErratum) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyErratum) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MyErratum) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *MyErratum) SetDetails(v string) {
	o.Details = &v
}

// GetPublishDate returns the PublishDate field value if set, zero value otherwise.
func (o *MyErratum) GetPublishDate() string {
	if o == nil || o.PublishDate == nil {
		var ret string
		return ret
	}
	return *o.PublishDate
}

// GetPublishDateOk returns a tuple with the PublishDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyErratum) GetPublishDateOk() (*string, bool) {
	if o == nil || o.PublishDate == nil {
		return nil, false
	}
	return o.PublishDate, true
}

// HasPublishDate returns a boolean if a field has been set.
func (o *MyErratum) HasPublishDate() bool {
	if o != nil && o.PublishDate != nil {
		return true
	}

	return false
}

// SetPublishDate gets a reference to the given string and assigns it to the PublishDate field.
func (o *MyErratum) SetPublishDate(v string) {
	o.PublishDate = &v
}

// GetSynopsis returns the Synopsis field value if set, zero value otherwise.
func (o *MyErratum) GetSynopsis() string {
	if o == nil || o.Synopsis == nil {
		var ret string
		return ret
	}
	return *o.Synopsis
}

// GetSynopsisOk returns a tuple with the Synopsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyErratum) GetSynopsisOk() (*string, bool) {
	if o == nil || o.Synopsis == nil {
		return nil, false
	}
	return o.Synopsis, true
}

// HasSynopsis returns a boolean if a field has been set.
func (o *MyErratum) HasSynopsis() bool {
	if o != nil && o.Synopsis != nil {
		return true
	}

	return false
}

// SetSynopsis gets a reference to the given string and assigns it to the Synopsis field.
func (o *MyErratum) SetSynopsis(v string) {
	o.Synopsis = &v
}

// GetSystems returns the Systems field value if set, zero value otherwise.
func (o *MyErratum) GetSystems() string {
	if o == nil || o.Systems == nil {
		var ret string
		return ret
	}
	return *o.Systems
}

// GetSystemsOk returns a tuple with the Systems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyErratum) GetSystemsOk() (*string, bool) {
	if o == nil || o.Systems == nil {
		return nil, false
	}
	return o.Systems, true
}

// HasSystems returns a boolean if a field has been set.
func (o *MyErratum) HasSystems() bool {
	if o != nil && o.Systems != nil {
		return true
	}

	return false
}

// SetSystems gets a reference to the given string and assigns it to the Systems field.
func (o *MyErratum) SetSystems(v string) {
	o.Systems = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MyErratum) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyErratum) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MyErratum) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MyErratum) SetType(v string) {
	o.Type = &v
}

func (o MyErratum) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdvisoryId != nil {
		toSerialize["advisoryId"] = o.AdvisoryId
	}
	if o.AffectedSystemCount != nil {
		toSerialize["affectedSystemCount"] = o.AffectedSystemCount
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.PublishDate != nil {
		toSerialize["publishDate"] = o.PublishDate
	}
	if o.Synopsis != nil {
		toSerialize["synopsis"] = o.Synopsis
	}
	if o.Systems != nil {
		toSerialize["systems"] = o.Systems
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMyErratum struct {
	value *MyErratum
	isSet bool
}

func (v NullableMyErratum) Get() *MyErratum {
	return v.value
}

func (v *NullableMyErratum) Set(val *MyErratum) {
	v.value = val
	v.isSet = true
}

func (v NullableMyErratum) IsSet() bool {
	return v.isSet
}

func (v *NullableMyErratum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyErratum(val *MyErratum) *NullableMyErratum {
	return &NullableMyErratum{value: val, isSet: true}
}

func (v NullableMyErratum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyErratum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
