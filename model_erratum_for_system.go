/*
RHSM-API

API for Red Hat Subscription Management

API version: 1.366.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// checks if the ErratumForSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErratumForSystem{}

// ErratumForSystem an erratum listed for a system
type ErratumForSystem struct {
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	// Date represents the date format used for API returns
	Published *string `json:"published,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Synopsis *string `json:"synopsis,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewErratumForSystem instantiates a new ErratumForSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErratumForSystem() *ErratumForSystem {
	this := ErratumForSystem{}
	return &this
}

// NewErratumForSystemWithDefaults instantiates a new ErratumForSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErratumForSystemWithDefaults() *ErratumForSystem {
	this := ErratumForSystem{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ErratumForSystem) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumForSystem) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ErratumForSystem) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ErratumForSystem) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ErratumForSystem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumForSystem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ErratumForSystem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ErratumForSystem) SetId(v string) {
	o.Id = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *ErratumForSystem) GetPublished() string {
	if o == nil || IsNil(o.Published) {
		var ret string
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumForSystem) GetPublishedOk() (*string, bool) {
	if o == nil || IsNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *ErratumForSystem) HasPublished() bool {
	if o != nil && !IsNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given string and assigns it to the Published field.
func (o *ErratumForSystem) SetPublished(v string) {
	o.Published = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ErratumForSystem) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumForSystem) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ErratumForSystem) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ErratumForSystem) SetSeverity(v string) {
	o.Severity = &v
}

// GetSynopsis returns the Synopsis field value if set, zero value otherwise.
func (o *ErratumForSystem) GetSynopsis() string {
	if o == nil || IsNil(o.Synopsis) {
		var ret string
		return ret
	}
	return *o.Synopsis
}

// GetSynopsisOk returns a tuple with the Synopsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumForSystem) GetSynopsisOk() (*string, bool) {
	if o == nil || IsNil(o.Synopsis) {
		return nil, false
	}
	return o.Synopsis, true
}

// HasSynopsis returns a boolean if a field has been set.
func (o *ErratumForSystem) HasSynopsis() bool {
	if o != nil && !IsNil(o.Synopsis) {
		return true
	}

	return false
}

// SetSynopsis gets a reference to the given string and assigns it to the Synopsis field.
func (o *ErratumForSystem) SetSynopsis(v string) {
	o.Synopsis = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErratumForSystem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumForSystem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErratumForSystem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErratumForSystem) SetType(v string) {
	o.Type = &v
}

func (o ErratumForSystem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErratumForSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Published) {
		toSerialize["published"] = o.Published
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Synopsis) {
		toSerialize["synopsis"] = o.Synopsis
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableErratumForSystem struct {
	value *ErratumForSystem
	isSet bool
}

func (v NullableErratumForSystem) Get() *ErratumForSystem {
	return v.value
}

func (v *NullableErratumForSystem) Set(val *ErratumForSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableErratumForSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableErratumForSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErratumForSystem(val *ErratumForSystem) *NullableErratumForSystem {
	return &NullableErratumForSystem{value: val, isSet: true}
}

func (v NullableErratumForSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErratumForSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


