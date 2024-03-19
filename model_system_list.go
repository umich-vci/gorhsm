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

// checks if the SystemList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemList{}

// SystemList SystemList is the result of the get system list API
type SystemList struct {
	ComplianceStatus *string `json:"complianceStatus,omitempty"`
	Details *string `json:"details,omitempty"`
	EntitlementQuantity *int32 `json:"entitlementQuantity,omitempty"`
	// Date represents the date format used for API returns
	LastCheckin *string `json:"lastCheckin,omitempty"`
	SystemName *string `json:"systemName,omitempty"`
	TotalEntitlementQuantity *int32 `json:"totalEntitlementQuantity,omitempty"`
	Type *string `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewSystemList instantiates a new SystemList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemList() *SystemList {
	this := SystemList{}
	return &this
}

// NewSystemListWithDefaults instantiates a new SystemList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemListWithDefaults() *SystemList {
	this := SystemList{}
	return &this
}

// GetComplianceStatus returns the ComplianceStatus field value if set, zero value otherwise.
func (o *SystemList) GetComplianceStatus() string {
	if o == nil || IsNil(o.ComplianceStatus) {
		var ret string
		return ret
	}
	return *o.ComplianceStatus
}

// GetComplianceStatusOk returns a tuple with the ComplianceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemList) GetComplianceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ComplianceStatus) {
		return nil, false
	}
	return o.ComplianceStatus, true
}

// HasComplianceStatus returns a boolean if a field has been set.
func (o *SystemList) HasComplianceStatus() bool {
	if o != nil && !IsNil(o.ComplianceStatus) {
		return true
	}

	return false
}

// SetComplianceStatus gets a reference to the given string and assigns it to the ComplianceStatus field.
func (o *SystemList) SetComplianceStatus(v string) {
	o.ComplianceStatus = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SystemList) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemList) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SystemList) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SystemList) SetDetails(v string) {
	o.Details = &v
}

// GetEntitlementQuantity returns the EntitlementQuantity field value if set, zero value otherwise.
func (o *SystemList) GetEntitlementQuantity() int32 {
	if o == nil || IsNil(o.EntitlementQuantity) {
		var ret int32
		return ret
	}
	return *o.EntitlementQuantity
}

// GetEntitlementQuantityOk returns a tuple with the EntitlementQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemList) GetEntitlementQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.EntitlementQuantity) {
		return nil, false
	}
	return o.EntitlementQuantity, true
}

// HasEntitlementQuantity returns a boolean if a field has been set.
func (o *SystemList) HasEntitlementQuantity() bool {
	if o != nil && !IsNil(o.EntitlementQuantity) {
		return true
	}

	return false
}

// SetEntitlementQuantity gets a reference to the given int32 and assigns it to the EntitlementQuantity field.
func (o *SystemList) SetEntitlementQuantity(v int32) {
	o.EntitlementQuantity = &v
}

// GetLastCheckin returns the LastCheckin field value if set, zero value otherwise.
func (o *SystemList) GetLastCheckin() string {
	if o == nil || IsNil(o.LastCheckin) {
		var ret string
		return ret
	}
	return *o.LastCheckin
}

// GetLastCheckinOk returns a tuple with the LastCheckin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemList) GetLastCheckinOk() (*string, bool) {
	if o == nil || IsNil(o.LastCheckin) {
		return nil, false
	}
	return o.LastCheckin, true
}

// HasLastCheckin returns a boolean if a field has been set.
func (o *SystemList) HasLastCheckin() bool {
	if o != nil && !IsNil(o.LastCheckin) {
		return true
	}

	return false
}

// SetLastCheckin gets a reference to the given string and assigns it to the LastCheckin field.
func (o *SystemList) SetLastCheckin(v string) {
	o.LastCheckin = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *SystemList) GetSystemName() string {
	if o == nil || IsNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemList) GetSystemNameOk() (*string, bool) {
	if o == nil || IsNil(o.SystemName) {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *SystemList) HasSystemName() bool {
	if o != nil && !IsNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *SystemList) SetSystemName(v string) {
	o.SystemName = &v
}

// GetTotalEntitlementQuantity returns the TotalEntitlementQuantity field value if set, zero value otherwise.
func (o *SystemList) GetTotalEntitlementQuantity() int32 {
	if o == nil || IsNil(o.TotalEntitlementQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalEntitlementQuantity
}

// GetTotalEntitlementQuantityOk returns a tuple with the TotalEntitlementQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemList) GetTotalEntitlementQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalEntitlementQuantity) {
		return nil, false
	}
	return o.TotalEntitlementQuantity, true
}

// HasTotalEntitlementQuantity returns a boolean if a field has been set.
func (o *SystemList) HasTotalEntitlementQuantity() bool {
	if o != nil && !IsNil(o.TotalEntitlementQuantity) {
		return true
	}

	return false
}

// SetTotalEntitlementQuantity gets a reference to the given int32 and assigns it to the TotalEntitlementQuantity field.
func (o *SystemList) SetTotalEntitlementQuantity(v int32) {
	o.TotalEntitlementQuantity = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SystemList) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemList) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SystemList) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SystemList) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SystemList) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemList) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SystemList) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SystemList) SetUuid(v string) {
	o.Uuid = &v
}

func (o SystemList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComplianceStatus) {
		toSerialize["complianceStatus"] = o.ComplianceStatus
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.EntitlementQuantity) {
		toSerialize["entitlementQuantity"] = o.EntitlementQuantity
	}
	if !IsNil(o.LastCheckin) {
		toSerialize["lastCheckin"] = o.LastCheckin
	}
	if !IsNil(o.SystemName) {
		toSerialize["systemName"] = o.SystemName
	}
	if !IsNil(o.TotalEntitlementQuantity) {
		toSerialize["totalEntitlementQuantity"] = o.TotalEntitlementQuantity
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableSystemList struct {
	value *SystemList
	isSet bool
}

func (v NullableSystemList) Get() *SystemList {
	return v.value
}

func (v *NullableSystemList) Set(val *SystemList) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemList) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemList(val *SystemList) *NullableSystemList {
	return &NullableSystemList{value: val, isSet: true}
}

func (v NullableSystemList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


