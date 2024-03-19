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

// checks if the ListSubContentSets200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSubContentSets200Response{}

// ListSubContentSets200Response struct for ListSubContentSets200Response
type ListSubContentSets200Response struct {
	Body []ContentSet `json:"body,omitempty"`
	Pagination *APIPageParam `json:"pagination,omitempty"`
}

// NewListSubContentSets200Response instantiates a new ListSubContentSets200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSubContentSets200Response() *ListSubContentSets200Response {
	this := ListSubContentSets200Response{}
	return &this
}

// NewListSubContentSets200ResponseWithDefaults instantiates a new ListSubContentSets200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSubContentSets200ResponseWithDefaults() *ListSubContentSets200Response {
	this := ListSubContentSets200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ListSubContentSets200Response) GetBody() []ContentSet {
	if o == nil || IsNil(o.Body) {
		var ret []ContentSet
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubContentSets200Response) GetBodyOk() ([]ContentSet, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ListSubContentSets200Response) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given []ContentSet and assigns it to the Body field.
func (o *ListSubContentSets200Response) SetBody(v []ContentSet) {
	o.Body = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListSubContentSets200Response) GetPagination() APIPageParam {
	if o == nil || IsNil(o.Pagination) {
		var ret APIPageParam
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubContentSets200Response) GetPaginationOk() (*APIPageParam, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListSubContentSets200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given APIPageParam and assigns it to the Pagination field.
func (o *ListSubContentSets200Response) SetPagination(v APIPageParam) {
	o.Pagination = &v
}

func (o ListSubContentSets200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSubContentSets200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListSubContentSets200Response struct {
	value *ListSubContentSets200Response
	isSet bool
}

func (v NullableListSubContentSets200Response) Get() *ListSubContentSets200Response {
	return v.value
}

func (v *NullableListSubContentSets200Response) Set(val *ListSubContentSets200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSubContentSets200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSubContentSets200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSubContentSets200Response(val *ListSubContentSets200Response) *NullableListSubContentSets200Response {
	return &NullableListSubContentSets200Response{value: val, isSet: true}
}

func (v NullableListSubContentSets200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSubContentSets200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


