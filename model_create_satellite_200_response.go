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

// checks if the CreateSatellite200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSatellite200Response{}

// CreateSatellite200Response struct for CreateSatellite200Response
type CreateSatellite200Response struct {
	Body *AllocationSummary `json:"body,omitempty"`
}

// NewCreateSatellite200Response instantiates a new CreateSatellite200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSatellite200Response() *CreateSatellite200Response {
	this := CreateSatellite200Response{}
	return &this
}

// NewCreateSatellite200ResponseWithDefaults instantiates a new CreateSatellite200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSatellite200ResponseWithDefaults() *CreateSatellite200Response {
	this := CreateSatellite200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CreateSatellite200Response) GetBody() AllocationSummary {
	if o == nil || IsNil(o.Body) {
		var ret AllocationSummary
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSatellite200Response) GetBodyOk() (*AllocationSummary, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CreateSatellite200Response) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given AllocationSummary and assigns it to the Body field.
func (o *CreateSatellite200Response) SetBody(v AllocationSummary) {
	o.Body = &v
}

func (o CreateSatellite200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSatellite200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableCreateSatellite200Response struct {
	value *CreateSatellite200Response
	isSet bool
}

func (v NullableCreateSatellite200Response) Get() *CreateSatellite200Response {
	return v.value
}

func (v *NullableCreateSatellite200Response) Set(val *CreateSatellite200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSatellite200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSatellite200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSatellite200Response(val *CreateSatellite200Response) *NullableCreateSatellite200Response {
	return &NullableCreateSatellite200Response{value: val, isSet: true}
}

func (v NullableCreateSatellite200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSatellite200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


