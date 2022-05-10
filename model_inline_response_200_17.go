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

// InlineResponse20017 struct for InlineResponse20017
type InlineResponse20017 struct {
	Body *AttachEntitlement `json:"body,omitempty"`
}

// NewInlineResponse20017 instantiates a new InlineResponse20017 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017() *InlineResponse20017 {
	this := InlineResponse20017{}
	return &this
}

// NewInlineResponse20017WithDefaults instantiates a new InlineResponse20017 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017WithDefaults() *InlineResponse20017 {
	this := InlineResponse20017{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineResponse20017) GetBody() AttachEntitlement {
	if o == nil || o.Body == nil {
		var ret AttachEntitlement
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017) GetBodyOk() (*AttachEntitlement, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineResponse20017) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given AttachEntitlement and assigns it to the Body field.
func (o *InlineResponse20017) SetBody(v AttachEntitlement) {
	o.Body = &v
}

func (o InlineResponse20017) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017 struct {
	value *InlineResponse20017
	isSet bool
}

func (v NullableInlineResponse20017) Get() *InlineResponse20017 {
	return v.value
}

func (v *NullableInlineResponse20017) Set(val *InlineResponse20017) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017(val *InlineResponse20017) *NullableInlineResponse20017 {
	return &NullableInlineResponse20017{value: val, isSet: true}
}

func (v NullableInlineResponse20017) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}