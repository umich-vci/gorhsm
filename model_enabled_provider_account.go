/*
RHSM-API

API for Red Hat Subscription Management

API version: 1.366.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EnabledProviderAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnabledProviderAccount{}

// EnabledProviderAccount Enabled Provider Account represents a cloud access provider account
type EnabledProviderAccount struct {
	// Date represents the date format used for API returns
	DateAdded string `json:"dateAdded"`
	GoldImageStatus []GoldImageStatus `json:"goldImageStatus,omitempty"`
	Id string `json:"id"`
	Nickname string `json:"nickname"`
	// Source ID of linked account (only for accounts created via Sources on cloud.redhat.com)
	SourceId *string `json:"sourceId,omitempty"`
	// verification status for RHSM Auto Registration (only displayed for supported cloud providers)
	Verified *bool `json:"verified,omitempty"`
}

type _EnabledProviderAccount EnabledProviderAccount

// NewEnabledProviderAccount instantiates a new EnabledProviderAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnabledProviderAccount(dateAdded string, id string, nickname string) *EnabledProviderAccount {
	this := EnabledProviderAccount{}
	this.DateAdded = dateAdded
	this.Id = id
	this.Nickname = nickname
	return &this
}

// NewEnabledProviderAccountWithDefaults instantiates a new EnabledProviderAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnabledProviderAccountWithDefaults() *EnabledProviderAccount {
	this := EnabledProviderAccount{}
	return &this
}

// GetDateAdded returns the DateAdded field value
func (o *EnabledProviderAccount) GetDateAdded() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value
// and a boolean to check if the value has been set.
func (o *EnabledProviderAccount) GetDateAddedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateAdded, true
}

// SetDateAdded sets field value
func (o *EnabledProviderAccount) SetDateAdded(v string) {
	o.DateAdded = v
}

// GetGoldImageStatus returns the GoldImageStatus field value if set, zero value otherwise.
func (o *EnabledProviderAccount) GetGoldImageStatus() []GoldImageStatus {
	if o == nil || IsNil(o.GoldImageStatus) {
		var ret []GoldImageStatus
		return ret
	}
	return o.GoldImageStatus
}

// GetGoldImageStatusOk returns a tuple with the GoldImageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledProviderAccount) GetGoldImageStatusOk() ([]GoldImageStatus, bool) {
	if o == nil || IsNil(o.GoldImageStatus) {
		return nil, false
	}
	return o.GoldImageStatus, true
}

// HasGoldImageStatus returns a boolean if a field has been set.
func (o *EnabledProviderAccount) HasGoldImageStatus() bool {
	if o != nil && !IsNil(o.GoldImageStatus) {
		return true
	}

	return false
}

// SetGoldImageStatus gets a reference to the given []GoldImageStatus and assigns it to the GoldImageStatus field.
func (o *EnabledProviderAccount) SetGoldImageStatus(v []GoldImageStatus) {
	o.GoldImageStatus = v
}

// GetId returns the Id field value
func (o *EnabledProviderAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnabledProviderAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnabledProviderAccount) SetId(v string) {
	o.Id = v
}

// GetNickname returns the Nickname field value
func (o *EnabledProviderAccount) GetNickname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value
// and a boolean to check if the value has been set.
func (o *EnabledProviderAccount) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nickname, true
}

// SetNickname sets field value
func (o *EnabledProviderAccount) SetNickname(v string) {
	o.Nickname = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *EnabledProviderAccount) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledProviderAccount) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *EnabledProviderAccount) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *EnabledProviderAccount) SetSourceId(v string) {
	o.SourceId = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *EnabledProviderAccount) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledProviderAccount) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *EnabledProviderAccount) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *EnabledProviderAccount) SetVerified(v bool) {
	o.Verified = &v
}

func (o EnabledProviderAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnabledProviderAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dateAdded"] = o.DateAdded
	if !IsNil(o.GoldImageStatus) {
		toSerialize["goldImageStatus"] = o.GoldImageStatus
	}
	toSerialize["id"] = o.Id
	toSerialize["nickname"] = o.Nickname
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	return toSerialize, nil
}

func (o *EnabledProviderAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dateAdded",
		"id",
		"nickname",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEnabledProviderAccount := _EnabledProviderAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnabledProviderAccount)

	if err != nil {
		return err
	}

	*o = EnabledProviderAccount(varEnabledProviderAccount)

	return err
}

type NullableEnabledProviderAccount struct {
	value *EnabledProviderAccount
	isSet bool
}

func (v NullableEnabledProviderAccount) Get() *EnabledProviderAccount {
	return v.value
}

func (v *NullableEnabledProviderAccount) Set(val *EnabledProviderAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableEnabledProviderAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableEnabledProviderAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnabledProviderAccount(val *EnabledProviderAccount) *NullableEnabledProviderAccount {
	return &NullableEnabledProviderAccount{value: val, isSet: true}
}

func (v NullableEnabledProviderAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnabledProviderAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


