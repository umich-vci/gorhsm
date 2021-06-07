/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.182.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// InlineObject struct for InlineObject
type InlineObject struct {
	SimpleContentAccess string `json:"simpleContentAccess"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject(simpleContentAccess string) *InlineObject {
	this := InlineObject{}
	this.SimpleContentAccess = simpleContentAccess
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetSimpleContentAccess returns the SimpleContentAccess field value
func (o *InlineObject) GetSimpleContentAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SimpleContentAccess
}

// GetSimpleContentAccessOk returns a tuple with the SimpleContentAccess field value
// and a boolean to check if the value has been set.
func (o *InlineObject) GetSimpleContentAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SimpleContentAccess, true
}

// SetSimpleContentAccess sets field value
func (o *InlineObject) SetSimpleContentAccess(v string) {
	o.SimpleContentAccess = v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["simpleContentAccess"] = o.SimpleContentAccess
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
