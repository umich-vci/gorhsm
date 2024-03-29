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

// checks if the EntitlementsAttachedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementsAttachedResponse{}

// EntitlementsAttachedResponse EntitlementsAttachedResponse wraps data obtained for EntitlementsAttached and sends metadata about it using helpers.OptionalResult
type EntitlementsAttachedResponse struct {
	Reason *string `json:"reason,omitempty"`
	Valid *bool `json:"valid,omitempty"`
	Value []EntitlementsAttachedResponseValue `json:"value,omitempty"`
}

// NewEntitlementsAttachedResponse instantiates a new EntitlementsAttachedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementsAttachedResponse() *EntitlementsAttachedResponse {
	this := EntitlementsAttachedResponse{}
	return &this
}

// NewEntitlementsAttachedResponseWithDefaults instantiates a new EntitlementsAttachedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementsAttachedResponseWithDefaults() *EntitlementsAttachedResponse {
	this := EntitlementsAttachedResponse{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponse) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EntitlementsAttachedResponse) SetReason(v string) {
	o.Reason = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponse) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponse) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponse) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *EntitlementsAttachedResponse) SetValid(v bool) {
	o.Valid = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponse) GetValue() []EntitlementsAttachedResponseValue {
	if o == nil || IsNil(o.Value) {
		var ret []EntitlementsAttachedResponseValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponse) GetValueOk() ([]EntitlementsAttachedResponseValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []EntitlementsAttachedResponseValue and assigns it to the Value field.
func (o *EntitlementsAttachedResponse) SetValue(v []EntitlementsAttachedResponseValue) {
	o.Value = v
}

func (o EntitlementsAttachedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementsAttachedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableEntitlementsAttachedResponse struct {
	value *EntitlementsAttachedResponse
	isSet bool
}

func (v NullableEntitlementsAttachedResponse) Get() *EntitlementsAttachedResponse {
	return v.value
}

func (v *NullableEntitlementsAttachedResponse) Set(val *EntitlementsAttachedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementsAttachedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementsAttachedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementsAttachedResponse(val *EntitlementsAttachedResponse) *NullableEntitlementsAttachedResponse {
	return &NullableEntitlementsAttachedResponse{value: val, isSet: true}
}

func (v NullableEntitlementsAttachedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementsAttachedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


