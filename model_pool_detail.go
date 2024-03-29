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

// checks if the PoolDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolDetail{}

// PoolDetail PoolDetail is an entry in the system/allocation pools listing
type PoolDetail struct {
	ContractNumber *string `json:"contractNumber,omitempty"`
	// Date represents the date format used for API returns
	EndDate *string `json:"endDate,omitempty"`
	EntitlementsAvailable *int32 `json:"entitlementsAvailable,omitempty"`
	Id *string `json:"id,omitempty"`
	ServiceLevel *string `json:"serviceLevel,omitempty"`
	Sku *string `json:"sku,omitempty"`
	// Date represents the date format used for API returns
	StartDate *string `json:"startDate,omitempty"`
	SubscriptionName *string `json:"subscriptionName,omitempty"`
	SubscriptionNumber *string `json:"subscriptionNumber,omitempty"`
}

// NewPoolDetail instantiates a new PoolDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDetail() *PoolDetail {
	this := PoolDetail{}
	return &this
}

// NewPoolDetailWithDefaults instantiates a new PoolDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDetailWithDefaults() *PoolDetail {
	this := PoolDetail{}
	return &this
}

// GetContractNumber returns the ContractNumber field value if set, zero value otherwise.
func (o *PoolDetail) GetContractNumber() string {
	if o == nil || IsNil(o.ContractNumber) {
		var ret string
		return ret
	}
	return *o.ContractNumber
}

// GetContractNumberOk returns a tuple with the ContractNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetail) GetContractNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ContractNumber) {
		return nil, false
	}
	return o.ContractNumber, true
}

// HasContractNumber returns a boolean if a field has been set.
func (o *PoolDetail) HasContractNumber() bool {
	if o != nil && !IsNil(o.ContractNumber) {
		return true
	}

	return false
}

// SetContractNumber gets a reference to the given string and assigns it to the ContractNumber field.
func (o *PoolDetail) SetContractNumber(v string) {
	o.ContractNumber = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *PoolDetail) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetail) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *PoolDetail) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *PoolDetail) SetEndDate(v string) {
	o.EndDate = &v
}

// GetEntitlementsAvailable returns the EntitlementsAvailable field value if set, zero value otherwise.
func (o *PoolDetail) GetEntitlementsAvailable() int32 {
	if o == nil || IsNil(o.EntitlementsAvailable) {
		var ret int32
		return ret
	}
	return *o.EntitlementsAvailable
}

// GetEntitlementsAvailableOk returns a tuple with the EntitlementsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetail) GetEntitlementsAvailableOk() (*int32, bool) {
	if o == nil || IsNil(o.EntitlementsAvailable) {
		return nil, false
	}
	return o.EntitlementsAvailable, true
}

// HasEntitlementsAvailable returns a boolean if a field has been set.
func (o *PoolDetail) HasEntitlementsAvailable() bool {
	if o != nil && !IsNil(o.EntitlementsAvailable) {
		return true
	}

	return false
}

// SetEntitlementsAvailable gets a reference to the given int32 and assigns it to the EntitlementsAvailable field.
func (o *PoolDetail) SetEntitlementsAvailable(v int32) {
	o.EntitlementsAvailable = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoolDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoolDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PoolDetail) SetId(v string) {
	o.Id = &v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *PoolDetail) GetServiceLevel() string {
	if o == nil || IsNil(o.ServiceLevel) {
		var ret string
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetail) GetServiceLevelOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLevel) {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *PoolDetail) HasServiceLevel() bool {
	if o != nil && !IsNil(o.ServiceLevel) {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given string and assigns it to the ServiceLevel field.
func (o *PoolDetail) SetServiceLevel(v string) {
	o.ServiceLevel = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *PoolDetail) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetail) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *PoolDetail) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *PoolDetail) SetSku(v string) {
	o.Sku = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PoolDetail) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetail) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PoolDetail) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *PoolDetail) SetStartDate(v string) {
	o.StartDate = &v
}

// GetSubscriptionName returns the SubscriptionName field value if set, zero value otherwise.
func (o *PoolDetail) GetSubscriptionName() string {
	if o == nil || IsNil(o.SubscriptionName) {
		var ret string
		return ret
	}
	return *o.SubscriptionName
}

// GetSubscriptionNameOk returns a tuple with the SubscriptionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetail) GetSubscriptionNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionName) {
		return nil, false
	}
	return o.SubscriptionName, true
}

// HasSubscriptionName returns a boolean if a field has been set.
func (o *PoolDetail) HasSubscriptionName() bool {
	if o != nil && !IsNil(o.SubscriptionName) {
		return true
	}

	return false
}

// SetSubscriptionName gets a reference to the given string and assigns it to the SubscriptionName field.
func (o *PoolDetail) SetSubscriptionName(v string) {
	o.SubscriptionName = &v
}

// GetSubscriptionNumber returns the SubscriptionNumber field value if set, zero value otherwise.
func (o *PoolDetail) GetSubscriptionNumber() string {
	if o == nil || IsNil(o.SubscriptionNumber) {
		var ret string
		return ret
	}
	return *o.SubscriptionNumber
}

// GetSubscriptionNumberOk returns a tuple with the SubscriptionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetail) GetSubscriptionNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionNumber) {
		return nil, false
	}
	return o.SubscriptionNumber, true
}

// HasSubscriptionNumber returns a boolean if a field has been set.
func (o *PoolDetail) HasSubscriptionNumber() bool {
	if o != nil && !IsNil(o.SubscriptionNumber) {
		return true
	}

	return false
}

// SetSubscriptionNumber gets a reference to the given string and assigns it to the SubscriptionNumber field.
func (o *PoolDetail) SetSubscriptionNumber(v string) {
	o.SubscriptionNumber = &v
}

func (o PoolDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractNumber) {
		toSerialize["contractNumber"] = o.ContractNumber
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.EntitlementsAvailable) {
		toSerialize["entitlementsAvailable"] = o.EntitlementsAvailable
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ServiceLevel) {
		toSerialize["serviceLevel"] = o.ServiceLevel
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
	if !IsNil(o.SubscriptionNumber) {
		toSerialize["subscriptionNumber"] = o.SubscriptionNumber
	}
	return toSerialize, nil
}

type NullablePoolDetail struct {
	value *PoolDetail
	isSet bool
}

func (v NullablePoolDetail) Get() *PoolDetail {
	return v.value
}

func (v *NullablePoolDetail) Set(val *PoolDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDetail(val *PoolDetail) *NullablePoolDetail {
	return &NullablePoolDetail{value: val, isSet: true}
}

func (v NullablePoolDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


