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

// checks if the AttachEntitlement200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachEntitlement200Response{}

// AttachEntitlement200Response struct for AttachEntitlement200Response
type AttachEntitlement200Response struct {
	Body *AttachEntitlement `json:"body,omitempty"`
}

// NewAttachEntitlement200Response instantiates a new AttachEntitlement200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachEntitlement200Response() *AttachEntitlement200Response {
	this := AttachEntitlement200Response{}
	return &this
}

// NewAttachEntitlement200ResponseWithDefaults instantiates a new AttachEntitlement200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachEntitlement200ResponseWithDefaults() *AttachEntitlement200Response {
	this := AttachEntitlement200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *AttachEntitlement200Response) GetBody() AttachEntitlement {
	if o == nil || IsNil(o.Body) {
		var ret AttachEntitlement
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement200Response) GetBodyOk() (*AttachEntitlement, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *AttachEntitlement200Response) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given AttachEntitlement and assigns it to the Body field.
func (o *AttachEntitlement200Response) SetBody(v AttachEntitlement) {
	o.Body = &v
}

func (o AttachEntitlement200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachEntitlement200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableAttachEntitlement200Response struct {
	value *AttachEntitlement200Response
	isSet bool
}

func (v NullableAttachEntitlement200Response) Get() *AttachEntitlement200Response {
	return v.value
}

func (v *NullableAttachEntitlement200Response) Set(val *AttachEntitlement200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachEntitlement200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachEntitlement200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachEntitlement200Response(val *AttachEntitlement200Response) *NullableAttachEntitlement200Response {
	return &NullableAttachEntitlement200Response{value: val, isSet: true}
}

func (v NullableAttachEntitlement200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachEntitlement200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


