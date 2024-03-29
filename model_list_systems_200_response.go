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

// checks if the ListSystems200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSystems200Response{}

// ListSystems200Response struct for ListSystems200Response
type ListSystems200Response struct {
	// systemList is a System Slice
	Body []System `json:"body,omitempty"`
	Pagination *APIPageParam `json:"pagination,omitempty"`
}

// NewListSystems200Response instantiates a new ListSystems200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSystems200Response() *ListSystems200Response {
	this := ListSystems200Response{}
	return &this
}

// NewListSystems200ResponseWithDefaults instantiates a new ListSystems200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSystems200ResponseWithDefaults() *ListSystems200Response {
	this := ListSystems200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ListSystems200Response) GetBody() []System {
	if o == nil || IsNil(o.Body) {
		var ret []System
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSystems200Response) GetBodyOk() ([]System, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ListSystems200Response) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given []System and assigns it to the Body field.
func (o *ListSystems200Response) SetBody(v []System) {
	o.Body = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListSystems200Response) GetPagination() APIPageParam {
	if o == nil || IsNil(o.Pagination) {
		var ret APIPageParam
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSystems200Response) GetPaginationOk() (*APIPageParam, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListSystems200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given APIPageParam and assigns it to the Pagination field.
func (o *ListSystems200Response) SetPagination(v APIPageParam) {
	o.Pagination = &v
}

func (o ListSystems200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSystems200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListSystems200Response struct {
	value *ListSystems200Response
	isSet bool
}

func (v NullableListSystems200Response) Get() *ListSystems200Response {
	return v.value
}

func (v *NullableListSystems200Response) Set(val *ListSystems200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSystems200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSystems200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSystems200Response(val *ListSystems200Response) *NullableListSystems200Response {
	return &NullableListSystems200Response{value: val, isSet: true}
}

func (v NullableListSystems200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSystems200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


