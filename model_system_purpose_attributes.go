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

// SystemPurposeAttributes System purpose settings available to an organization
type SystemPurposeAttributes struct {
	Roles        []string `json:"roles,omitempty"`
	ServiceLevel []string `json:"serviceLevel,omitempty"`
	Usage        []string `json:"usage,omitempty"`
}

// NewSystemPurposeAttributes instantiates a new SystemPurposeAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPurposeAttributes() *SystemPurposeAttributes {
	this := SystemPurposeAttributes{}
	return &this
}

// NewSystemPurposeAttributesWithDefaults instantiates a new SystemPurposeAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPurposeAttributesWithDefaults() *SystemPurposeAttributes {
	this := SystemPurposeAttributes{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *SystemPurposeAttributes) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeAttributes) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *SystemPurposeAttributes) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *SystemPurposeAttributes) SetRoles(v []string) {
	o.Roles = v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *SystemPurposeAttributes) GetServiceLevel() []string {
	if o == nil || o.ServiceLevel == nil {
		var ret []string
		return ret
	}
	return o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeAttributes) GetServiceLevelOk() ([]string, bool) {
	if o == nil || o.ServiceLevel == nil {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *SystemPurposeAttributes) HasServiceLevel() bool {
	if o != nil && o.ServiceLevel != nil {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given []string and assigns it to the ServiceLevel field.
func (o *SystemPurposeAttributes) SetServiceLevel(v []string) {
	o.ServiceLevel = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *SystemPurposeAttributes) GetUsage() []string {
	if o == nil || o.Usage == nil {
		var ret []string
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeAttributes) GetUsageOk() ([]string, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *SystemPurposeAttributes) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []string and assigns it to the Usage field.
func (o *SystemPurposeAttributes) SetUsage(v []string) {
	o.Usage = v
}

func (o SystemPurposeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.ServiceLevel != nil {
		toSerialize["serviceLevel"] = o.ServiceLevel
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableSystemPurposeAttributes struct {
	value *SystemPurposeAttributes
	isSet bool
}

func (v NullableSystemPurposeAttributes) Get() *SystemPurposeAttributes {
	return v.value
}

func (v *NullableSystemPurposeAttributes) Set(val *SystemPurposeAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPurposeAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPurposeAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPurposeAttributes(val *SystemPurposeAttributes) *NullableSystemPurposeAttributes {
	return &NullableSystemPurposeAttributes{value: val, isSet: true}
}

func (v NullableSystemPurposeAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPurposeAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
