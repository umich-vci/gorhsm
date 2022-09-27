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

// Facts facts give additional details about the system
type Facts struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewFacts instantiates a new Facts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacts() *Facts {
	this := Facts{}
	return &this
}

// NewFactsWithDefaults instantiates a new Facts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFactsWithDefaults() *Facts {
	this := Facts{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Facts) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facts) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Facts) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Facts) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Facts) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facts) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Facts) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Facts) SetValue(v string) {
	o.Value = &v
}

func (o Facts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFacts struct {
	value *Facts
	isSet bool
}

func (v NullableFacts) Get() *Facts {
	return v.value
}

func (v *NullableFacts) Set(val *Facts) {
	v.value = val
	v.isSet = true
}

func (v NullableFacts) IsSet() bool {
	return v.isSet
}

func (v *NullableFacts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacts(val *Facts) *NullableFacts {
	return &NullableFacts{value: val, isSet: true}
}

func (v NullableFacts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
