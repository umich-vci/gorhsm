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

// checks if the ListVersionsAllocation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVersionsAllocation200Response{}

// ListVersionsAllocation200Response struct for ListVersionsAllocation200Response
type ListVersionsAllocation200Response struct {
	Body []AllocationVersion `json:"body,omitempty"`
}

// NewListVersionsAllocation200Response instantiates a new ListVersionsAllocation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVersionsAllocation200Response() *ListVersionsAllocation200Response {
	this := ListVersionsAllocation200Response{}
	return &this
}

// NewListVersionsAllocation200ResponseWithDefaults instantiates a new ListVersionsAllocation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVersionsAllocation200ResponseWithDefaults() *ListVersionsAllocation200Response {
	this := ListVersionsAllocation200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ListVersionsAllocation200Response) GetBody() []AllocationVersion {
	if o == nil || IsNil(o.Body) {
		var ret []AllocationVersion
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVersionsAllocation200Response) GetBodyOk() ([]AllocationVersion, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ListVersionsAllocation200Response) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given []AllocationVersion and assigns it to the Body field.
func (o *ListVersionsAllocation200Response) SetBody(v []AllocationVersion) {
	o.Body = v
}

func (o ListVersionsAllocation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVersionsAllocation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableListVersionsAllocation200Response struct {
	value *ListVersionsAllocation200Response
	isSet bool
}

func (v NullableListVersionsAllocation200Response) Get() *ListVersionsAllocation200Response {
	return v.value
}

func (v *NullableListVersionsAllocation200Response) Set(val *ListVersionsAllocation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListVersionsAllocation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListVersionsAllocation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVersionsAllocation200Response(val *ListVersionsAllocation200Response) *NullableListVersionsAllocation200Response {
	return &NullableListVersionsAllocation200Response{value: val, isSet: true}
}

func (v NullableListVersionsAllocation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVersionsAllocation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


