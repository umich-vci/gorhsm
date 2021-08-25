/*
RHSM-API

API for Red Hat Subscription Management

API version: 1.196.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// InlineResponse2008 struct for InlineResponse2008
type InlineResponse2008 struct {
	Body       *[]ImageInContentSet `json:"body,omitempty"`
	Pagination *APIPageParam        `json:"pagination,omitempty"`
}

// NewInlineResponse2008 instantiates a new InlineResponse2008 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008() *InlineResponse2008 {
	this := InlineResponse2008{}
	return &this
}

// NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008WithDefaults() *InlineResponse2008 {
	this := InlineResponse2008{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineResponse2008) GetBody() []ImageInContentSet {
	if o == nil || o.Body == nil {
		var ret []ImageInContentSet
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008) GetBodyOk() (*[]ImageInContentSet, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineResponse2008) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given []ImageInContentSet and assigns it to the Body field.
func (o *InlineResponse2008) SetBody(v []ImageInContentSet) {
	o.Body = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *InlineResponse2008) GetPagination() APIPageParam {
	if o == nil || o.Pagination == nil {
		var ret APIPageParam
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008) GetPaginationOk() (*APIPageParam, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *InlineResponse2008) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given APIPageParam and assigns it to the Pagination field.
func (o *InlineResponse2008) SetPagination(v APIPageParam) {
	o.Pagination = &v
}

func (o InlineResponse2008) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2008 struct {
	value *InlineResponse2008
	isSet bool
}

func (v NullableInlineResponse2008) Get() *InlineResponse2008 {
	return v.value
}

func (v *NullableInlineResponse2008) Set(val *InlineResponse2008) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008(val *InlineResponse2008) *NullableInlineResponse2008 {
	return &NullableInlineResponse2008{value: val, isSet: true}
}

func (v NullableInlineResponse2008) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
