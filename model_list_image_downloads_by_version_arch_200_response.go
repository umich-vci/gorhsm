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

// ListImageDownloadsByVersionArch200Response struct for ListImageDownloadsByVersionArch200Response
type ListImageDownloadsByVersionArch200Response struct {
	Body []ImageInContentSet `json:"body,omitempty"`
}

// NewListImageDownloadsByVersionArch200Response instantiates a new ListImageDownloadsByVersionArch200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImageDownloadsByVersionArch200Response() *ListImageDownloadsByVersionArch200Response {
	this := ListImageDownloadsByVersionArch200Response{}
	return &this
}

// NewListImageDownloadsByVersionArch200ResponseWithDefaults instantiates a new ListImageDownloadsByVersionArch200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImageDownloadsByVersionArch200ResponseWithDefaults() *ListImageDownloadsByVersionArch200Response {
	this := ListImageDownloadsByVersionArch200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ListImageDownloadsByVersionArch200Response) GetBody() []ImageInContentSet {
	if o == nil || o.Body == nil {
		var ret []ImageInContentSet
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageDownloadsByVersionArch200Response) GetBodyOk() ([]ImageInContentSet, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ListImageDownloadsByVersionArch200Response) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given []ImageInContentSet and assigns it to the Body field.
func (o *ListImageDownloadsByVersionArch200Response) SetBody(v []ImageInContentSet) {
	o.Body = v
}

func (o ListImageDownloadsByVersionArch200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableListImageDownloadsByVersionArch200Response struct {
	value *ListImageDownloadsByVersionArch200Response
	isSet bool
}

func (v NullableListImageDownloadsByVersionArch200Response) Get() *ListImageDownloadsByVersionArch200Response {
	return v.value
}

func (v *NullableListImageDownloadsByVersionArch200Response) Set(val *ListImageDownloadsByVersionArch200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListImageDownloadsByVersionArch200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListImageDownloadsByVersionArch200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListImageDownloadsByVersionArch200Response(val *ListImageDownloadsByVersionArch200Response) *NullableListImageDownloadsByVersionArch200Response {
	return &NullableListImageDownloadsByVersionArch200Response{value: val, isSet: true}
}

func (v NullableListImageDownloadsByVersionArch200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListImageDownloadsByVersionArch200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}