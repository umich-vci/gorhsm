/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.177.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// InlineResponse20014 struct for InlineResponse20014
type InlineResponse20014 struct {
	// systemList is a System Slice
	Body       *[]System     `json:"body,omitempty"`
	Pagination *APIPageParam `json:"pagination,omitempty"`
}

// NewInlineResponse20014 instantiates a new InlineResponse20014 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014() *InlineResponse20014 {
	this := InlineResponse20014{}
	return &this
}

// NewInlineResponse20014WithDefaults instantiates a new InlineResponse20014 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014WithDefaults() *InlineResponse20014 {
	this := InlineResponse20014{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineResponse20014) GetBody() []System {
	if o == nil || o.Body == nil {
		var ret []System
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetBodyOk() (*[]System, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineResponse20014) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given []System and assigns it to the Body field.
func (o *InlineResponse20014) SetBody(v []System) {
	o.Body = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *InlineResponse20014) GetPagination() APIPageParam {
	if o == nil || o.Pagination == nil {
		var ret APIPageParam
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetPaginationOk() (*APIPageParam, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *InlineResponse20014) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given APIPageParam and assigns it to the Pagination field.
func (o *InlineResponse20014) SetPagination(v APIPageParam) {
	o.Pagination = &v
}

func (o InlineResponse20014) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20014 struct {
	value *InlineResponse20014
	isSet bool
}

func (v NullableInlineResponse20014) Get() *InlineResponse20014 {
	return v.value
}

func (v *NullableInlineResponse20014) Set(val *InlineResponse20014) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20014) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20014) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20014(val *InlineResponse20014) *NullableInlineResponse20014 {
	return &NullableInlineResponse20014{value: val, isSet: true}
}

func (v NullableInlineResponse20014) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20014) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
