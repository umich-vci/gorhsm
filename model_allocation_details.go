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

// AllocationDetails details of a subscription allocation
type AllocationDetails struct {
	CreatedBy *string `json:"createdBy,omitempty"`
	// Date represents the date format used for API returns
	CreatedDate                  *string                       `json:"createdDate,omitempty"`
	EntitlementsAttached         *EntitlementsAttachedResponse `json:"entitlementsAttached,omitempty"`
	EntitlementsAttachedQuantity *int32                        `json:"entitlementsAttachedQuantity,omitempty"`
	// Date represents the date format used for API returns
	LastModified        *string `json:"lastModified,omitempty"`
	Name                *string `json:"name,omitempty"`
	SimpleContentAccess *string `json:"simpleContentAccess,omitempty"`
	Type                *string `json:"type,omitempty"`
	Uuid                *string `json:"uuid,omitempty"`
	Version             *string `json:"version,omitempty"`
}

// NewAllocationDetails instantiates a new AllocationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocationDetails() *AllocationDetails {
	this := AllocationDetails{}
	return &this
}

// NewAllocationDetailsWithDefaults instantiates a new AllocationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocationDetailsWithDefaults() *AllocationDetails {
	this := AllocationDetails{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AllocationDetails) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AllocationDetails) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AllocationDetails) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *AllocationDetails) GetCreatedDate() string {
	if o == nil || o.CreatedDate == nil {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetCreatedDateOk() (*string, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AllocationDetails) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *AllocationDetails) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetEntitlementsAttached returns the EntitlementsAttached field value if set, zero value otherwise.
func (o *AllocationDetails) GetEntitlementsAttached() EntitlementsAttachedResponse {
	if o == nil || o.EntitlementsAttached == nil {
		var ret EntitlementsAttachedResponse
		return ret
	}
	return *o.EntitlementsAttached
}

// GetEntitlementsAttachedOk returns a tuple with the EntitlementsAttached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetEntitlementsAttachedOk() (*EntitlementsAttachedResponse, bool) {
	if o == nil || o.EntitlementsAttached == nil {
		return nil, false
	}
	return o.EntitlementsAttached, true
}

// HasEntitlementsAttached returns a boolean if a field has been set.
func (o *AllocationDetails) HasEntitlementsAttached() bool {
	if o != nil && o.EntitlementsAttached != nil {
		return true
	}

	return false
}

// SetEntitlementsAttached gets a reference to the given EntitlementsAttachedResponse and assigns it to the EntitlementsAttached field.
func (o *AllocationDetails) SetEntitlementsAttached(v EntitlementsAttachedResponse) {
	o.EntitlementsAttached = &v
}

// GetEntitlementsAttachedQuantity returns the EntitlementsAttachedQuantity field value if set, zero value otherwise.
func (o *AllocationDetails) GetEntitlementsAttachedQuantity() int32 {
	if o == nil || o.EntitlementsAttachedQuantity == nil {
		var ret int32
		return ret
	}
	return *o.EntitlementsAttachedQuantity
}

// GetEntitlementsAttachedQuantityOk returns a tuple with the EntitlementsAttachedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetEntitlementsAttachedQuantityOk() (*int32, bool) {
	if o == nil || o.EntitlementsAttachedQuantity == nil {
		return nil, false
	}
	return o.EntitlementsAttachedQuantity, true
}

// HasEntitlementsAttachedQuantity returns a boolean if a field has been set.
func (o *AllocationDetails) HasEntitlementsAttachedQuantity() bool {
	if o != nil && o.EntitlementsAttachedQuantity != nil {
		return true
	}

	return false
}

// SetEntitlementsAttachedQuantity gets a reference to the given int32 and assigns it to the EntitlementsAttachedQuantity field.
func (o *AllocationDetails) SetEntitlementsAttachedQuantity(v int32) {
	o.EntitlementsAttachedQuantity = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *AllocationDetails) GetLastModified() string {
	if o == nil || o.LastModified == nil {
		var ret string
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetLastModifiedOk() (*string, bool) {
	if o == nil || o.LastModified == nil {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *AllocationDetails) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given string and assigns it to the LastModified field.
func (o *AllocationDetails) SetLastModified(v string) {
	o.LastModified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AllocationDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AllocationDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AllocationDetails) SetName(v string) {
	o.Name = &v
}

// GetSimpleContentAccess returns the SimpleContentAccess field value if set, zero value otherwise.
func (o *AllocationDetails) GetSimpleContentAccess() string {
	if o == nil || o.SimpleContentAccess == nil {
		var ret string
		return ret
	}
	return *o.SimpleContentAccess
}

// GetSimpleContentAccessOk returns a tuple with the SimpleContentAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetSimpleContentAccessOk() (*string, bool) {
	if o == nil || o.SimpleContentAccess == nil {
		return nil, false
	}
	return o.SimpleContentAccess, true
}

// HasSimpleContentAccess returns a boolean if a field has been set.
func (o *AllocationDetails) HasSimpleContentAccess() bool {
	if o != nil && o.SimpleContentAccess != nil {
		return true
	}

	return false
}

// SetSimpleContentAccess gets a reference to the given string and assigns it to the SimpleContentAccess field.
func (o *AllocationDetails) SetSimpleContentAccess(v string) {
	o.SimpleContentAccess = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AllocationDetails) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AllocationDetails) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AllocationDetails) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *AllocationDetails) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *AllocationDetails) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *AllocationDetails) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AllocationDetails) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationDetails) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AllocationDetails) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AllocationDetails) SetVersion(v string) {
	o.Version = &v
}

func (o AllocationDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if o.EntitlementsAttached != nil {
		toSerialize["entitlementsAttached"] = o.EntitlementsAttached
	}
	if o.EntitlementsAttachedQuantity != nil {
		toSerialize["entitlementsAttachedQuantity"] = o.EntitlementsAttachedQuantity
	}
	if o.LastModified != nil {
		toSerialize["lastModified"] = o.LastModified
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SimpleContentAccess != nil {
		toSerialize["simpleContentAccess"] = o.SimpleContentAccess
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableAllocationDetails struct {
	value *AllocationDetails
	isSet bool
}

func (v NullableAllocationDetails) Get() *AllocationDetails {
	return v.value
}

func (v *NullableAllocationDetails) Set(val *AllocationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocationDetails(val *AllocationDetails) *NullableAllocationDetails {
	return &NullableAllocationDetails{value: val, isSet: true}
}

func (v NullableAllocationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
