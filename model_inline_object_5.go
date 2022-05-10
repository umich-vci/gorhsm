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

// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	Accounts []string `json:"accounts"`
	Images   []string `json:"images"`
}

// NewInlineObject5 instantiates a new InlineObject5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject5(accounts []string, images []string) *InlineObject5 {
	this := InlineObject5{}
	this.Accounts = accounts
	this.Images = images
	return &this
}

// NewInlineObject5WithDefaults instantiates a new InlineObject5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject5WithDefaults() *InlineObject5 {
	this := InlineObject5{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *InlineObject5) GetAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetAccountsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *InlineObject5) SetAccounts(v []string) {
	o.Accounts = v
}

// GetImages returns the Images field value
func (o *InlineObject5) GetImages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetImagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Images, true
}

// SetImages sets field value
func (o *InlineObject5) SetImages(v []string) {
	o.Images = v
}

func (o InlineObject5) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	if true {
		toSerialize["images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject5 struct {
	value *InlineObject5
	isSet bool
}

func (v NullableInlineObject5) Get() *InlineObject5 {
	return v.value
}

func (v *NullableInlineObject5) Set(val *InlineObject5) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject5) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject5(val *InlineObject5) *NullableInlineObject5 {
	return &NullableInlineObject5{value: val, isSet: true}
}

func (v NullableInlineObject5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}