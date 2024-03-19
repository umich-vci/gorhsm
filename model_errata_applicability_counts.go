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

// checks if the ErrataApplicabilityCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrataApplicabilityCounts{}

// ErrataApplicabilityCounts Applicable errata details
type ErrataApplicabilityCounts struct {
	Reason *string `json:"reason,omitempty"`
	Valid *bool `json:"valid,omitempty"`
	Value *ErrataCount `json:"value,omitempty"`
}

// NewErrataApplicabilityCounts instantiates a new ErrataApplicabilityCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrataApplicabilityCounts() *ErrataApplicabilityCounts {
	this := ErrataApplicabilityCounts{}
	return &this
}

// NewErrataApplicabilityCountsWithDefaults instantiates a new ErrataApplicabilityCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrataApplicabilityCountsWithDefaults() *ErrataApplicabilityCounts {
	this := ErrataApplicabilityCounts{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ErrataApplicabilityCounts) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataApplicabilityCounts) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ErrataApplicabilityCounts) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ErrataApplicabilityCounts) SetReason(v string) {
	o.Reason = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *ErrataApplicabilityCounts) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataApplicabilityCounts) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *ErrataApplicabilityCounts) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *ErrataApplicabilityCounts) SetValid(v bool) {
	o.Valid = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ErrataApplicabilityCounts) GetValue() ErrataCount {
	if o == nil || IsNil(o.Value) {
		var ret ErrataCount
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataApplicabilityCounts) GetValueOk() (*ErrataCount, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ErrataApplicabilityCounts) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ErrataCount and assigns it to the Value field.
func (o *ErrataApplicabilityCounts) SetValue(v ErrataCount) {
	o.Value = &v
}

func (o ErrataApplicabilityCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrataApplicabilityCounts) ToMap() (map[string]interface{}, error) {
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

type NullableErrataApplicabilityCounts struct {
	value *ErrataApplicabilityCounts
	isSet bool
}

func (v NullableErrataApplicabilityCounts) Get() *ErrataApplicabilityCounts {
	return v.value
}

func (v *NullableErrataApplicabilityCounts) Set(val *ErrataApplicabilityCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableErrataApplicabilityCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableErrataApplicabilityCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrataApplicabilityCounts(val *ErrataApplicabilityCounts) *NullableErrataApplicabilityCounts {
	return &NullableErrataApplicabilityCounts{value: val, isSet: true}
}

func (v NullableErrataApplicabilityCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrataApplicabilityCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


