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

// InlineResponse2009 struct for InlineResponse2009
type InlineResponse2009 struct {
	Body       *[]PkgContentSetArch `json:"body,omitempty"`
	Pagination *APIPageParam        `json:"pagination,omitempty"`
}

// NewInlineResponse2009 instantiates a new InlineResponse2009 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// NewInlineResponse2009WithDefaults instantiates a new InlineResponse2009 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009WithDefaults() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineResponse2009) GetBody() []PkgContentSetArch {
	if o == nil || o.Body == nil {
		var ret []PkgContentSetArch
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetBodyOk() (*[]PkgContentSetArch, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineResponse2009) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given []PkgContentSetArch and assigns it to the Body field.
func (o *InlineResponse2009) SetBody(v []PkgContentSetArch) {
	o.Body = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *InlineResponse2009) GetPagination() APIPageParam {
	if o == nil || o.Pagination == nil {
		var ret APIPageParam
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetPaginationOk() (*APIPageParam, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *InlineResponse2009) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given APIPageParam and assigns it to the Pagination field.
func (o *InlineResponse2009) SetPagination(v APIPageParam) {
	o.Pagination = &v
}

func (o InlineResponse2009) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2009 struct {
	value *InlineResponse2009
	isSet bool
}

func (v NullableInlineResponse2009) Get() *InlineResponse2009 {
	return v.value
}

func (v *NullableInlineResponse2009) Set(val *InlineResponse2009) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2009) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2009) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2009(val *InlineResponse2009) *NullableInlineResponse2009 {
	return &NullableInlineResponse2009{value: val, isSet: true}
}

func (v NullableInlineResponse2009) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2009) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
