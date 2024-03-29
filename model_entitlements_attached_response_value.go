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

// checks if the EntitlementsAttachedResponseValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementsAttachedResponseValue{}

// EntitlementsAttachedResponseValue EntitlementsAttachedResponseValue represents the Value field in the EntitlementsAttachedResponse
type EntitlementsAttachedResponseValue struct {
	ContractNumber *string `json:"contractNumber,omitempty"`
	// Date represents the date format used for API returns
	EndDate *string `json:"endDate,omitempty"`
	EntitlementQuantity *int32 `json:"entitlementQuantity,omitempty"`
	Id *string `json:"id,omitempty"`
	Sku *string `json:"sku,omitempty"`
	// Date represents the date format used for API returns
	StartDate *string `json:"startDate,omitempty"`
	SubscriptionName *string `json:"subscriptionName,omitempty"`
}

// NewEntitlementsAttachedResponseValue instantiates a new EntitlementsAttachedResponseValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementsAttachedResponseValue() *EntitlementsAttachedResponseValue {
	this := EntitlementsAttachedResponseValue{}
	return &this
}

// NewEntitlementsAttachedResponseValueWithDefaults instantiates a new EntitlementsAttachedResponseValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementsAttachedResponseValueWithDefaults() *EntitlementsAttachedResponseValue {
	this := EntitlementsAttachedResponseValue{}
	return &this
}

// GetContractNumber returns the ContractNumber field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponseValue) GetContractNumber() string {
	if o == nil || IsNil(o.ContractNumber) {
		var ret string
		return ret
	}
	return *o.ContractNumber
}

// GetContractNumberOk returns a tuple with the ContractNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponseValue) GetContractNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ContractNumber) {
		return nil, false
	}
	return o.ContractNumber, true
}

// HasContractNumber returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponseValue) HasContractNumber() bool {
	if o != nil && !IsNil(o.ContractNumber) {
		return true
	}

	return false
}

// SetContractNumber gets a reference to the given string and assigns it to the ContractNumber field.
func (o *EntitlementsAttachedResponseValue) SetContractNumber(v string) {
	o.ContractNumber = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponseValue) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponseValue) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponseValue) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *EntitlementsAttachedResponseValue) SetEndDate(v string) {
	o.EndDate = &v
}

// GetEntitlementQuantity returns the EntitlementQuantity field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponseValue) GetEntitlementQuantity() int32 {
	if o == nil || IsNil(o.EntitlementQuantity) {
		var ret int32
		return ret
	}
	return *o.EntitlementQuantity
}

// GetEntitlementQuantityOk returns a tuple with the EntitlementQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponseValue) GetEntitlementQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.EntitlementQuantity) {
		return nil, false
	}
	return o.EntitlementQuantity, true
}

// HasEntitlementQuantity returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponseValue) HasEntitlementQuantity() bool {
	if o != nil && !IsNil(o.EntitlementQuantity) {
		return true
	}

	return false
}

// SetEntitlementQuantity gets a reference to the given int32 and assigns it to the EntitlementQuantity field.
func (o *EntitlementsAttachedResponseValue) SetEntitlementQuantity(v int32) {
	o.EntitlementQuantity = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponseValue) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponseValue) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponseValue) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementsAttachedResponseValue) SetId(v string) {
	o.Id = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponseValue) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponseValue) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponseValue) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EntitlementsAttachedResponseValue) SetSku(v string) {
	o.Sku = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponseValue) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponseValue) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponseValue) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *EntitlementsAttachedResponseValue) SetStartDate(v string) {
	o.StartDate = &v
}

// GetSubscriptionName returns the SubscriptionName field value if set, zero value otherwise.
func (o *EntitlementsAttachedResponseValue) GetSubscriptionName() string {
	if o == nil || IsNil(o.SubscriptionName) {
		var ret string
		return ret
	}
	return *o.SubscriptionName
}

// GetSubscriptionNameOk returns a tuple with the SubscriptionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementsAttachedResponseValue) GetSubscriptionNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionName) {
		return nil, false
	}
	return o.SubscriptionName, true
}

// HasSubscriptionName returns a boolean if a field has been set.
func (o *EntitlementsAttachedResponseValue) HasSubscriptionName() bool {
	if o != nil && !IsNil(o.SubscriptionName) {
		return true
	}

	return false
}

// SetSubscriptionName gets a reference to the given string and assigns it to the SubscriptionName field.
func (o *EntitlementsAttachedResponseValue) SetSubscriptionName(v string) {
	o.SubscriptionName = &v
}

func (o EntitlementsAttachedResponseValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementsAttachedResponseValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractNumber) {
		toSerialize["contractNumber"] = o.ContractNumber
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.EntitlementQuantity) {
		toSerialize["entitlementQuantity"] = o.EntitlementQuantity
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.SubscriptionName) {
		toSerialize["subscriptionName"] = o.SubscriptionName
	}
	return toSerialize, nil
}

type NullableEntitlementsAttachedResponseValue struct {
	value *EntitlementsAttachedResponseValue
	isSet bool
}

func (v NullableEntitlementsAttachedResponseValue) Get() *EntitlementsAttachedResponseValue {
	return v.value
}

func (v *NullableEntitlementsAttachedResponseValue) Set(val *EntitlementsAttachedResponseValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementsAttachedResponseValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementsAttachedResponseValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementsAttachedResponseValue(val *EntitlementsAttachedResponseValue) *NullableEntitlementsAttachedResponseValue {
	return &NullableEntitlementsAttachedResponseValue{value: val, isSet: true}
}

func (v NullableEntitlementsAttachedResponseValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementsAttachedResponseValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


