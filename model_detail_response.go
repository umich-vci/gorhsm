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

// checks if the DetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailResponse{}

// DetailResponse struct for DetailResponse
type DetailResponse struct {
	ContractNumber *string `json:"contractNumber,omitempty"`
	// Date represents the date format used for API returns
	EndDate *string `json:"endDate,omitempty"`
	Pools []Pool `json:"pools,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Sku *string `json:"sku,omitempty"`
	// Date represents the date format used for API returns
	StartDate *string `json:"startDate,omitempty"`
	Status *string `json:"status,omitempty"`
	SubscriptionName *string `json:"subscriptionName,omitempty"`
	SubscriptionNumber *string `json:"subscriptionNumber,omitempty"`
}

// NewDetailResponse instantiates a new DetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailResponse() *DetailResponse {
	this := DetailResponse{}
	return &this
}

// NewDetailResponseWithDefaults instantiates a new DetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailResponseWithDefaults() *DetailResponse {
	this := DetailResponse{}
	return &this
}

// GetContractNumber returns the ContractNumber field value if set, zero value otherwise.
func (o *DetailResponse) GetContractNumber() string {
	if o == nil || IsNil(o.ContractNumber) {
		var ret string
		return ret
	}
	return *o.ContractNumber
}

// GetContractNumberOk returns a tuple with the ContractNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailResponse) GetContractNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ContractNumber) {
		return nil, false
	}
	return o.ContractNumber, true
}

// HasContractNumber returns a boolean if a field has been set.
func (o *DetailResponse) HasContractNumber() bool {
	if o != nil && !IsNil(o.ContractNumber) {
		return true
	}

	return false
}

// SetContractNumber gets a reference to the given string and assigns it to the ContractNumber field.
func (o *DetailResponse) SetContractNumber(v string) {
	o.ContractNumber = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DetailResponse) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailResponse) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DetailResponse) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *DetailResponse) SetEndDate(v string) {
	o.EndDate = &v
}

// GetPools returns the Pools field value if set, zero value otherwise.
func (o *DetailResponse) GetPools() []Pool {
	if o == nil || IsNil(o.Pools) {
		var ret []Pool
		return ret
	}
	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailResponse) GetPoolsOk() ([]Pool, bool) {
	if o == nil || IsNil(o.Pools) {
		return nil, false
	}
	return o.Pools, true
}

// HasPools returns a boolean if a field has been set.
func (o *DetailResponse) HasPools() bool {
	if o != nil && !IsNil(o.Pools) {
		return true
	}

	return false
}

// SetPools gets a reference to the given []Pool and assigns it to the Pools field.
func (o *DetailResponse) SetPools(v []Pool) {
	o.Pools = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *DetailResponse) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailResponse) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *DetailResponse) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *DetailResponse) SetQuantity(v string) {
	o.Quantity = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *DetailResponse) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailResponse) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *DetailResponse) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *DetailResponse) SetSku(v string) {
	o.Sku = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DetailResponse) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailResponse) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DetailResponse) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *DetailResponse) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DetailResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DetailResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DetailResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubscriptionName returns the SubscriptionName field value if set, zero value otherwise.
func (o *DetailResponse) GetSubscriptionName() string {
	if o == nil || IsNil(o.SubscriptionName) {
		var ret string
		return ret
	}
	return *o.SubscriptionName
}

// GetSubscriptionNameOk returns a tuple with the SubscriptionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailResponse) GetSubscriptionNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionName) {
		return nil, false
	}
	return o.SubscriptionName, true
}

// HasSubscriptionName returns a boolean if a field has been set.
func (o *DetailResponse) HasSubscriptionName() bool {
	if o != nil && !IsNil(o.SubscriptionName) {
		return true
	}

	return false
}

// SetSubscriptionName gets a reference to the given string and assigns it to the SubscriptionName field.
func (o *DetailResponse) SetSubscriptionName(v string) {
	o.SubscriptionName = &v
}

// GetSubscriptionNumber returns the SubscriptionNumber field value if set, zero value otherwise.
func (o *DetailResponse) GetSubscriptionNumber() string {
	if o == nil || IsNil(o.SubscriptionNumber) {
		var ret string
		return ret
	}
	return *o.SubscriptionNumber
}

// GetSubscriptionNumberOk returns a tuple with the SubscriptionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailResponse) GetSubscriptionNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionNumber) {
		return nil, false
	}
	return o.SubscriptionNumber, true
}

// HasSubscriptionNumber returns a boolean if a field has been set.
func (o *DetailResponse) HasSubscriptionNumber() bool {
	if o != nil && !IsNil(o.SubscriptionNumber) {
		return true
	}

	return false
}

// SetSubscriptionNumber gets a reference to the given string and assigns it to the SubscriptionNumber field.
func (o *DetailResponse) SetSubscriptionNumber(v string) {
	o.SubscriptionNumber = &v
}

func (o DetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractNumber) {
		toSerialize["contractNumber"] = o.ContractNumber
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Pools) {
		toSerialize["pools"] = o.Pools
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscriptionName) {
		toSerialize["subscriptionName"] = o.SubscriptionName
	}
	if !IsNil(o.SubscriptionNumber) {
		toSerialize["subscriptionNumber"] = o.SubscriptionNumber
	}
	return toSerialize, nil
}

type NullableDetailResponse struct {
	value *DetailResponse
	isSet bool
}

func (v NullableDetailResponse) Get() *DetailResponse {
	return v.value
}

func (v *NullableDetailResponse) Set(val *DetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailResponse(val *DetailResponse) *NullableDetailResponse {
	return &NullableDetailResponse{value: val, isSet: true}
}

func (v NullableDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


