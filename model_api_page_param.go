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

// checks if the APIPageParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIPageParam{}

// APIPageParam APIPageParam details the pagination parameters in APIResponse
type APIPageParam struct {
	Count *int32 `json:"count,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
}

// NewAPIPageParam instantiates a new APIPageParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIPageParam() *APIPageParam {
	this := APIPageParam{}
	return &this
}

// NewAPIPageParamWithDefaults instantiates a new APIPageParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIPageParamWithDefaults() *APIPageParam {
	this := APIPageParam{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *APIPageParam) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIPageParam) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *APIPageParam) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *APIPageParam) SetCount(v int32) {
	o.Count = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *APIPageParam) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIPageParam) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *APIPageParam) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *APIPageParam) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *APIPageParam) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIPageParam) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *APIPageParam) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *APIPageParam) SetOffset(v int32) {
	o.Offset = &v
}

func (o APIPageParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIPageParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableAPIPageParam struct {
	value *APIPageParam
	isSet bool
}

func (v NullableAPIPageParam) Get() *APIPageParam {
	return v.value
}

func (v *NullableAPIPageParam) Set(val *APIPageParam) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIPageParam) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIPageParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIPageParam(val *APIPageParam) *NullableAPIPageParam {
	return &NullableAPIPageParam{value: val, isSet: true}
}

func (v NullableAPIPageParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIPageParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


