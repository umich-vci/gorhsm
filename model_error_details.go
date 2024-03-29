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

// checks if the ErrorDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorDetails{}

// ErrorDetails ErrorDetails details the Error in ErrorResponse
type ErrorDetails struct {
	Code *int32 `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewErrorDetails instantiates a new ErrorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetails() *ErrorDetails {
	this := ErrorDetails{}
	return &this
}

// NewErrorDetailsWithDefaults instantiates a new ErrorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailsWithDefaults() *ErrorDetails {
	this := ErrorDetails{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorDetails) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetails) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorDetails) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ErrorDetails) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorDetails) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetails) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorDetails) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorDetails) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableErrorDetails struct {
	value *ErrorDetails
	isSet bool
}

func (v NullableErrorDetails) Get() *ErrorDetails {
	return v.value
}

func (v *NullableErrorDetails) Set(val *ErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetails(val *ErrorDetails) *NullableErrorDetails {
	return &NullableErrorDetails{value: val, isSet: true}
}

func (v NullableErrorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


