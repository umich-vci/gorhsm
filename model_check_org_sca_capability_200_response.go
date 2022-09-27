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

// CheckOrgSCACapability200Response struct for CheckOrgSCACapability200Response
type CheckOrgSCACapability200Response struct {
	Body *OrgSimpleContentAccess `json:"body,omitempty"`
}

// NewCheckOrgSCACapability200Response instantiates a new CheckOrgSCACapability200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckOrgSCACapability200Response() *CheckOrgSCACapability200Response {
	this := CheckOrgSCACapability200Response{}
	return &this
}

// NewCheckOrgSCACapability200ResponseWithDefaults instantiates a new CheckOrgSCACapability200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckOrgSCACapability200ResponseWithDefaults() *CheckOrgSCACapability200Response {
	this := CheckOrgSCACapability200Response{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CheckOrgSCACapability200Response) GetBody() OrgSimpleContentAccess {
	if o == nil || o.Body == nil {
		var ret OrgSimpleContentAccess
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckOrgSCACapability200Response) GetBodyOk() (*OrgSimpleContentAccess, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CheckOrgSCACapability200Response) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given OrgSimpleContentAccess and assigns it to the Body field.
func (o *CheckOrgSCACapability200Response) SetBody(v OrgSimpleContentAccess) {
	o.Body = &v
}

func (o CheckOrgSCACapability200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableCheckOrgSCACapability200Response struct {
	value *CheckOrgSCACapability200Response
	isSet bool
}

func (v NullableCheckOrgSCACapability200Response) Get() *CheckOrgSCACapability200Response {
	return v.value
}

func (v *NullableCheckOrgSCACapability200Response) Set(val *CheckOrgSCACapability200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckOrgSCACapability200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckOrgSCACapability200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckOrgSCACapability200Response(val *CheckOrgSCACapability200Response) *NullableCheckOrgSCACapability200Response {
	return &NullableCheckOrgSCACapability200Response{value: val, isSet: true}
}

func (v NullableCheckOrgSCACapability200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckOrgSCACapability200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}