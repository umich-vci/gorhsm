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

// InlineResponse2007 struct for InlineResponse2007
type InlineResponse2007 struct {
	Body *ErratumDetails `json:"body,omitempty"`
}

// NewInlineResponse2007 instantiates a new InlineResponse2007 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007() *InlineResponse2007 {
	this := InlineResponse2007{}
	return &this
}

// NewInlineResponse2007WithDefaults instantiates a new InlineResponse2007 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007WithDefaults() *InlineResponse2007 {
	this := InlineResponse2007{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineResponse2007) GetBody() ErratumDetails {
	if o == nil || o.Body == nil {
		var ret ErratumDetails
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetBodyOk() (*ErratumDetails, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineResponse2007) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given ErratumDetails and assigns it to the Body field.
func (o *InlineResponse2007) SetBody(v ErratumDetails) {
	o.Body = &v
}

func (o InlineResponse2007) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2007 struct {
	value *InlineResponse2007
	isSet bool
}

func (v NullableInlineResponse2007) Get() *InlineResponse2007 {
	return v.value
}

func (v *NullableInlineResponse2007) Set(val *InlineResponse2007) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2007) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2007) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2007(val *InlineResponse2007) *NullableInlineResponse2007 {
	return &NullableInlineResponse2007{value: val, isSet: true}
}

func (v NullableInlineResponse2007) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2007) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
