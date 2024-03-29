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

// checks if the Allocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Allocation{}

// Allocation struct for Allocation
type Allocation struct {
	EntitlementQuantity *int32 `json:"entitlementQuantity,omitempty"`
	Name *string `json:"name,omitempty"`
	SimpleContentAccess *string `json:"simpleContentAccess,omitempty"`
	Type *string `json:"type,omitempty"`
	Url *string `json:"url,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewAllocation instantiates a new Allocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocation() *Allocation {
	this := Allocation{}
	return &this
}

// NewAllocationWithDefaults instantiates a new Allocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocationWithDefaults() *Allocation {
	this := Allocation{}
	return &this
}

// GetEntitlementQuantity returns the EntitlementQuantity field value if set, zero value otherwise.
func (o *Allocation) GetEntitlementQuantity() int32 {
	if o == nil || IsNil(o.EntitlementQuantity) {
		var ret int32
		return ret
	}
	return *o.EntitlementQuantity
}

// GetEntitlementQuantityOk returns a tuple with the EntitlementQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allocation) GetEntitlementQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.EntitlementQuantity) {
		return nil, false
	}
	return o.EntitlementQuantity, true
}

// HasEntitlementQuantity returns a boolean if a field has been set.
func (o *Allocation) HasEntitlementQuantity() bool {
	if o != nil && !IsNil(o.EntitlementQuantity) {
		return true
	}

	return false
}

// SetEntitlementQuantity gets a reference to the given int32 and assigns it to the EntitlementQuantity field.
func (o *Allocation) SetEntitlementQuantity(v int32) {
	o.EntitlementQuantity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Allocation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allocation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Allocation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Allocation) SetName(v string) {
	o.Name = &v
}

// GetSimpleContentAccess returns the SimpleContentAccess field value if set, zero value otherwise.
func (o *Allocation) GetSimpleContentAccess() string {
	if o == nil || IsNil(o.SimpleContentAccess) {
		var ret string
		return ret
	}
	return *o.SimpleContentAccess
}

// GetSimpleContentAccessOk returns a tuple with the SimpleContentAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allocation) GetSimpleContentAccessOk() (*string, bool) {
	if o == nil || IsNil(o.SimpleContentAccess) {
		return nil, false
	}
	return o.SimpleContentAccess, true
}

// HasSimpleContentAccess returns a boolean if a field has been set.
func (o *Allocation) HasSimpleContentAccess() bool {
	if o != nil && !IsNil(o.SimpleContentAccess) {
		return true
	}

	return false
}

// SetSimpleContentAccess gets a reference to the given string and assigns it to the SimpleContentAccess field.
func (o *Allocation) SetSimpleContentAccess(v string) {
	o.SimpleContentAccess = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Allocation) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allocation) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Allocation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Allocation) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Allocation) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allocation) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Allocation) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Allocation) SetUrl(v string) {
	o.Url = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Allocation) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allocation) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Allocation) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Allocation) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Allocation) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Allocation) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Allocation) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Allocation) SetVersion(v string) {
	o.Version = &v
}

func (o Allocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Allocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntitlementQuantity) {
		toSerialize["entitlementQuantity"] = o.EntitlementQuantity
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SimpleContentAccess) {
		toSerialize["simpleContentAccess"] = o.SimpleContentAccess
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableAllocation struct {
	value *Allocation
	isSet bool
}

func (v NullableAllocation) Get() *Allocation {
	return v.value
}

func (v *NullableAllocation) Set(val *Allocation) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocation(val *Allocation) *NullableAllocation {
	return &NullableAllocation{value: val, isSet: true}
}

func (v NullableAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


