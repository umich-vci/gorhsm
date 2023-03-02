/*
RHSM-API

API for Red Hat Subscription Management

API version: 1.313.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// checks if the ListPackagesByContentSetArch200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPackagesByContentSetArch200Response{}

// ListPackagesByContentSetArch200Response struct for ListPackagesByContentSetArch200Response
type ListPackagesByContentSetArch200Response struct {
	Body []PkgContentSetArch `json:"body,omitempty"`
	Pagination *APIPageParam `json:"pagination,omitempty"`
}

// NewListPackagesByContentSetArch200Response instantiates a new ListPackagesByContentSetArch200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPackagesByContentSetArch200Response() *ListPackagesByContentSetArch200Response {
	this := ListPackagesByContentSetArch200Response{}
	return &this
}

// NewListPackagesByContentSetArch200ResponseWithDefaults instantiates a new ListPackagesByContentSetArch200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPackagesByContentSetArch200ResponseWithDefaults() *ListPackagesByContentSetArch200Response {
	this := ListPackagesByContentSetArch200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ListPackagesByContentSetArch200Response) GetBody() []PkgContentSetArch {
	if o == nil || IsNil(o.Body) {
		var ret []PkgContentSetArch
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPackagesByContentSetArch200Response) GetBodyOk() ([]PkgContentSetArch, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ListPackagesByContentSetArch200Response) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given []PkgContentSetArch and assigns it to the Body field.
func (o *ListPackagesByContentSetArch200Response) SetBody(v []PkgContentSetArch) {
	o.Body = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListPackagesByContentSetArch200Response) GetPagination() APIPageParam {
	if o == nil || IsNil(o.Pagination) {
		var ret APIPageParam
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPackagesByContentSetArch200Response) GetPaginationOk() (*APIPageParam, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListPackagesByContentSetArch200Response) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given APIPageParam and assigns it to the Pagination field.
func (o *ListPackagesByContentSetArch200Response) SetPagination(v APIPageParam) {
	o.Pagination = &v
}

func (o ListPackagesByContentSetArch200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPackagesByContentSetArch200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListPackagesByContentSetArch200Response struct {
	value *ListPackagesByContentSetArch200Response
	isSet bool
}

func (v NullableListPackagesByContentSetArch200Response) Get() *ListPackagesByContentSetArch200Response {
	return v.value
}

func (v *NullableListPackagesByContentSetArch200Response) Set(val *ListPackagesByContentSetArch200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListPackagesByContentSetArch200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListPackagesByContentSetArch200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPackagesByContentSetArch200Response(val *ListPackagesByContentSetArch200Response) *NullableListPackagesByContentSetArch200Response {
	return &NullableListPackagesByContentSetArch200Response{value: val, isSet: true}
}

func (v NullableListPackagesByContentSetArch200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPackagesByContentSetArch200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


