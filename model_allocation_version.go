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

// checks if the AllocationVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllocationVersion{}

// AllocationVersion List of satellite version
type AllocationVersion struct {
	Description *string `json:"description,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewAllocationVersion instantiates a new AllocationVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocationVersion() *AllocationVersion {
	this := AllocationVersion{}
	return &this
}

// NewAllocationVersionWithDefaults instantiates a new AllocationVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocationVersionWithDefaults() *AllocationVersion {
	this := AllocationVersion{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AllocationVersion) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationVersion) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AllocationVersion) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AllocationVersion) SetDescription(v string) {
	o.Description = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AllocationVersion) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationVersion) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AllocationVersion) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AllocationVersion) SetValue(v string) {
	o.Value = &v
}

func (o AllocationVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllocationVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAllocationVersion struct {
	value *AllocationVersion
	isSet bool
}

func (v NullableAllocationVersion) Get() *AllocationVersion {
	return v.value
}

func (v *NullableAllocationVersion) Set(val *AllocationVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocationVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocationVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocationVersion(val *AllocationVersion) *NullableAllocationVersion {
	return &NullableAllocationVersion{value: val, isSet: true}
}

func (v NullableAllocationVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocationVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


