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

// PoolsListMock struct for PoolsListMock
type PoolsListMock struct {
	Body       []PoolDetail  `json:"body,omitempty"`
	Pagination *APIPageParam `json:"pagination,omitempty"`
}

// NewPoolsListMock instantiates a new PoolsListMock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolsListMock() *PoolsListMock {
	this := PoolsListMock{}
	return &this
}

// NewPoolsListMockWithDefaults instantiates a new PoolsListMock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolsListMockWithDefaults() *PoolsListMock {
	this := PoolsListMock{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *PoolsListMock) GetBody() []PoolDetail {
	if o == nil || o.Body == nil {
		var ret []PoolDetail
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsListMock) GetBodyOk() ([]PoolDetail, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *PoolsListMock) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given []PoolDetail and assigns it to the Body field.
func (o *PoolsListMock) SetBody(v []PoolDetail) {
	o.Body = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *PoolsListMock) GetPagination() APIPageParam {
	if o == nil || o.Pagination == nil {
		var ret APIPageParam
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsListMock) GetPaginationOk() (*APIPageParam, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *PoolsListMock) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given APIPageParam and assigns it to the Pagination field.
func (o *PoolsListMock) SetPagination(v APIPageParam) {
	o.Pagination = &v
}

func (o PoolsListMock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullablePoolsListMock struct {
	value *PoolsListMock
	isSet bool
}

func (v NullablePoolsListMock) Get() *PoolsListMock {
	return v.value
}

func (v *NullablePoolsListMock) Set(val *PoolsListMock) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolsListMock) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolsListMock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolsListMock(val *PoolsListMock) *NullablePoolsListMock {
	return &NullablePoolsListMock{value: val, isSet: true}
}

func (v NullablePoolsListMock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolsListMock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
