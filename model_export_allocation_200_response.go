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

// ExportAllocation200Response struct for ExportAllocation200Response
type ExportAllocation200Response struct {
	Body *ExportResponse `json:"body,omitempty"`
}

// NewExportAllocation200Response instantiates a new ExportAllocation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportAllocation200Response() *ExportAllocation200Response {
	this := ExportAllocation200Response{}
	return &this
}

// NewExportAllocation200ResponseWithDefaults instantiates a new ExportAllocation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportAllocation200ResponseWithDefaults() *ExportAllocation200Response {
	this := ExportAllocation200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ExportAllocation200Response) GetBody() ExportResponse {
	if o == nil || o.Body == nil {
		var ret ExportResponse
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportAllocation200Response) GetBodyOk() (*ExportResponse, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ExportAllocation200Response) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given ExportResponse and assigns it to the Body field.
func (o *ExportAllocation200Response) SetBody(v ExportResponse) {
	o.Body = &v
}

func (o ExportAllocation200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableExportAllocation200Response struct {
	value *ExportAllocation200Response
	isSet bool
}

func (v NullableExportAllocation200Response) Get() *ExportAllocation200Response {
	return v.value
}

func (v *NullableExportAllocation200Response) Set(val *ExportAllocation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableExportAllocation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableExportAllocation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportAllocation200Response(val *ExportAllocation200Response) *NullableExportAllocation200Response {
	return &NullableExportAllocation200Response{value: val, isSet: true}
}

func (v NullableExportAllocation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportAllocation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}