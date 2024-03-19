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

// checks if the AttachEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachEntitlement{}

// AttachEntitlement struct for AttachEntitlement
type AttachEntitlement struct {
	AutoAttachSetting *bool `json:"autoAttachSetting,omitempty"`
	ComplianceStatus *string `json:"complianceStatus,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	// Date represents the date format used for API returns
	CreatedDate *string `json:"createdDate,omitempty"`
	EntitlementStatus *string `json:"entitlementStatus,omitempty"`
	EntitlementsAttached *EntitlementsAttachedResponse `json:"entitlementsAttached,omitempty"`
	EntitlementsAttachedCount *int32 `json:"entitlementsAttachedCount,omitempty"`
	ErrataApplicabilityCounts *ErrataApplicabilityCounts `json:"errataApplicabilityCounts,omitempty"`
	FactsCount *int32 `json:"factsCount,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	InstalledProductsCount *int32 `json:"installedProductsCount,omitempty"`
	// Date represents the date format used for API returns
	LastCheckin *string `json:"lastCheckin,omitempty"`
	Name *string `json:"name,omitempty"`
	ServiceLevelPreference *string `json:"serviceLevelPreference,omitempty"`
	Type *string `json:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewAttachEntitlement instantiates a new AttachEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachEntitlement() *AttachEntitlement {
	this := AttachEntitlement{}
	return &this
}

// NewAttachEntitlementWithDefaults instantiates a new AttachEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachEntitlementWithDefaults() *AttachEntitlement {
	this := AttachEntitlement{}
	return &this
}

// GetAutoAttachSetting returns the AutoAttachSetting field value if set, zero value otherwise.
func (o *AttachEntitlement) GetAutoAttachSetting() bool {
	if o == nil || IsNil(o.AutoAttachSetting) {
		var ret bool
		return ret
	}
	return *o.AutoAttachSetting
}

// GetAutoAttachSettingOk returns a tuple with the AutoAttachSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetAutoAttachSettingOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoAttachSetting) {
		return nil, false
	}
	return o.AutoAttachSetting, true
}

// HasAutoAttachSetting returns a boolean if a field has been set.
func (o *AttachEntitlement) HasAutoAttachSetting() bool {
	if o != nil && !IsNil(o.AutoAttachSetting) {
		return true
	}

	return false
}

// SetAutoAttachSetting gets a reference to the given bool and assigns it to the AutoAttachSetting field.
func (o *AttachEntitlement) SetAutoAttachSetting(v bool) {
	o.AutoAttachSetting = &v
}

// GetComplianceStatus returns the ComplianceStatus field value if set, zero value otherwise.
func (o *AttachEntitlement) GetComplianceStatus() string {
	if o == nil || IsNil(o.ComplianceStatus) {
		var ret string
		return ret
	}
	return *o.ComplianceStatus
}

// GetComplianceStatusOk returns a tuple with the ComplianceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetComplianceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ComplianceStatus) {
		return nil, false
	}
	return o.ComplianceStatus, true
}

// HasComplianceStatus returns a boolean if a field has been set.
func (o *AttachEntitlement) HasComplianceStatus() bool {
	if o != nil && !IsNil(o.ComplianceStatus) {
		return true
	}

	return false
}

// SetComplianceStatus gets a reference to the given string and assigns it to the ComplianceStatus field.
func (o *AttachEntitlement) SetComplianceStatus(v string) {
	o.ComplianceStatus = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AttachEntitlement) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AttachEntitlement) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AttachEntitlement) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *AttachEntitlement) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AttachEntitlement) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *AttachEntitlement) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetEntitlementStatus returns the EntitlementStatus field value if set, zero value otherwise.
func (o *AttachEntitlement) GetEntitlementStatus() string {
	if o == nil || IsNil(o.EntitlementStatus) {
		var ret string
		return ret
	}
	return *o.EntitlementStatus
}

// GetEntitlementStatusOk returns a tuple with the EntitlementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetEntitlementStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementStatus) {
		return nil, false
	}
	return o.EntitlementStatus, true
}

// HasEntitlementStatus returns a boolean if a field has been set.
func (o *AttachEntitlement) HasEntitlementStatus() bool {
	if o != nil && !IsNil(o.EntitlementStatus) {
		return true
	}

	return false
}

// SetEntitlementStatus gets a reference to the given string and assigns it to the EntitlementStatus field.
func (o *AttachEntitlement) SetEntitlementStatus(v string) {
	o.EntitlementStatus = &v
}

// GetEntitlementsAttached returns the EntitlementsAttached field value if set, zero value otherwise.
func (o *AttachEntitlement) GetEntitlementsAttached() EntitlementsAttachedResponse {
	if o == nil || IsNil(o.EntitlementsAttached) {
		var ret EntitlementsAttachedResponse
		return ret
	}
	return *o.EntitlementsAttached
}

// GetEntitlementsAttachedOk returns a tuple with the EntitlementsAttached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetEntitlementsAttachedOk() (*EntitlementsAttachedResponse, bool) {
	if o == nil || IsNil(o.EntitlementsAttached) {
		return nil, false
	}
	return o.EntitlementsAttached, true
}

// HasEntitlementsAttached returns a boolean if a field has been set.
func (o *AttachEntitlement) HasEntitlementsAttached() bool {
	if o != nil && !IsNil(o.EntitlementsAttached) {
		return true
	}

	return false
}

// SetEntitlementsAttached gets a reference to the given EntitlementsAttachedResponse and assigns it to the EntitlementsAttached field.
func (o *AttachEntitlement) SetEntitlementsAttached(v EntitlementsAttachedResponse) {
	o.EntitlementsAttached = &v
}

// GetEntitlementsAttachedCount returns the EntitlementsAttachedCount field value if set, zero value otherwise.
func (o *AttachEntitlement) GetEntitlementsAttachedCount() int32 {
	if o == nil || IsNil(o.EntitlementsAttachedCount) {
		var ret int32
		return ret
	}
	return *o.EntitlementsAttachedCount
}

// GetEntitlementsAttachedCountOk returns a tuple with the EntitlementsAttachedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetEntitlementsAttachedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EntitlementsAttachedCount) {
		return nil, false
	}
	return o.EntitlementsAttachedCount, true
}

// HasEntitlementsAttachedCount returns a boolean if a field has been set.
func (o *AttachEntitlement) HasEntitlementsAttachedCount() bool {
	if o != nil && !IsNil(o.EntitlementsAttachedCount) {
		return true
	}

	return false
}

// SetEntitlementsAttachedCount gets a reference to the given int32 and assigns it to the EntitlementsAttachedCount field.
func (o *AttachEntitlement) SetEntitlementsAttachedCount(v int32) {
	o.EntitlementsAttachedCount = &v
}

// GetErrataApplicabilityCounts returns the ErrataApplicabilityCounts field value if set, zero value otherwise.
func (o *AttachEntitlement) GetErrataApplicabilityCounts() ErrataApplicabilityCounts {
	if o == nil || IsNil(o.ErrataApplicabilityCounts) {
		var ret ErrataApplicabilityCounts
		return ret
	}
	return *o.ErrataApplicabilityCounts
}

// GetErrataApplicabilityCountsOk returns a tuple with the ErrataApplicabilityCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetErrataApplicabilityCountsOk() (*ErrataApplicabilityCounts, bool) {
	if o == nil || IsNil(o.ErrataApplicabilityCounts) {
		return nil, false
	}
	return o.ErrataApplicabilityCounts, true
}

// HasErrataApplicabilityCounts returns a boolean if a field has been set.
func (o *AttachEntitlement) HasErrataApplicabilityCounts() bool {
	if o != nil && !IsNil(o.ErrataApplicabilityCounts) {
		return true
	}

	return false
}

// SetErrataApplicabilityCounts gets a reference to the given ErrataApplicabilityCounts and assigns it to the ErrataApplicabilityCounts field.
func (o *AttachEntitlement) SetErrataApplicabilityCounts(v ErrataApplicabilityCounts) {
	o.ErrataApplicabilityCounts = &v
}

// GetFactsCount returns the FactsCount field value if set, zero value otherwise.
func (o *AttachEntitlement) GetFactsCount() int32 {
	if o == nil || IsNil(o.FactsCount) {
		var ret int32
		return ret
	}
	return *o.FactsCount
}

// GetFactsCountOk returns a tuple with the FactsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetFactsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FactsCount) {
		return nil, false
	}
	return o.FactsCount, true
}

// HasFactsCount returns a boolean if a field has been set.
func (o *AttachEntitlement) HasFactsCount() bool {
	if o != nil && !IsNil(o.FactsCount) {
		return true
	}

	return false
}

// SetFactsCount gets a reference to the given int32 and assigns it to the FactsCount field.
func (o *AttachEntitlement) SetFactsCount(v int32) {
	o.FactsCount = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *AttachEntitlement) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *AttachEntitlement) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *AttachEntitlement) SetHostname(v string) {
	o.Hostname = &v
}

// GetInstalledProductsCount returns the InstalledProductsCount field value if set, zero value otherwise.
func (o *AttachEntitlement) GetInstalledProductsCount() int32 {
	if o == nil || IsNil(o.InstalledProductsCount) {
		var ret int32
		return ret
	}
	return *o.InstalledProductsCount
}

// GetInstalledProductsCountOk returns a tuple with the InstalledProductsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetInstalledProductsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InstalledProductsCount) {
		return nil, false
	}
	return o.InstalledProductsCount, true
}

// HasInstalledProductsCount returns a boolean if a field has been set.
func (o *AttachEntitlement) HasInstalledProductsCount() bool {
	if o != nil && !IsNil(o.InstalledProductsCount) {
		return true
	}

	return false
}

// SetInstalledProductsCount gets a reference to the given int32 and assigns it to the InstalledProductsCount field.
func (o *AttachEntitlement) SetInstalledProductsCount(v int32) {
	o.InstalledProductsCount = &v
}

// GetLastCheckin returns the LastCheckin field value if set, zero value otherwise.
func (o *AttachEntitlement) GetLastCheckin() string {
	if o == nil || IsNil(o.LastCheckin) {
		var ret string
		return ret
	}
	return *o.LastCheckin
}

// GetLastCheckinOk returns a tuple with the LastCheckin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetLastCheckinOk() (*string, bool) {
	if o == nil || IsNil(o.LastCheckin) {
		return nil, false
	}
	return o.LastCheckin, true
}

// HasLastCheckin returns a boolean if a field has been set.
func (o *AttachEntitlement) HasLastCheckin() bool {
	if o != nil && !IsNil(o.LastCheckin) {
		return true
	}

	return false
}

// SetLastCheckin gets a reference to the given string and assigns it to the LastCheckin field.
func (o *AttachEntitlement) SetLastCheckin(v string) {
	o.LastCheckin = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttachEntitlement) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttachEntitlement) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttachEntitlement) SetName(v string) {
	o.Name = &v
}

// GetServiceLevelPreference returns the ServiceLevelPreference field value if set, zero value otherwise.
func (o *AttachEntitlement) GetServiceLevelPreference() string {
	if o == nil || IsNil(o.ServiceLevelPreference) {
		var ret string
		return ret
	}
	return *o.ServiceLevelPreference
}

// GetServiceLevelPreferenceOk returns a tuple with the ServiceLevelPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetServiceLevelPreferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLevelPreference) {
		return nil, false
	}
	return o.ServiceLevelPreference, true
}

// HasServiceLevelPreference returns a boolean if a field has been set.
func (o *AttachEntitlement) HasServiceLevelPreference() bool {
	if o != nil && !IsNil(o.ServiceLevelPreference) {
		return true
	}

	return false
}

// SetServiceLevelPreference gets a reference to the given string and assigns it to the ServiceLevelPreference field.
func (o *AttachEntitlement) SetServiceLevelPreference(v string) {
	o.ServiceLevelPreference = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AttachEntitlement) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AttachEntitlement) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AttachEntitlement) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *AttachEntitlement) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachEntitlement) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *AttachEntitlement) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *AttachEntitlement) SetUuid(v string) {
	o.Uuid = &v
}

func (o AttachEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoAttachSetting) {
		toSerialize["autoAttachSetting"] = o.AutoAttachSetting
	}
	if !IsNil(o.ComplianceStatus) {
		toSerialize["complianceStatus"] = o.ComplianceStatus
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.EntitlementStatus) {
		toSerialize["entitlementStatus"] = o.EntitlementStatus
	}
	if !IsNil(o.EntitlementsAttached) {
		toSerialize["entitlementsAttached"] = o.EntitlementsAttached
	}
	if !IsNil(o.EntitlementsAttachedCount) {
		toSerialize["entitlementsAttachedCount"] = o.EntitlementsAttachedCount
	}
	if !IsNil(o.ErrataApplicabilityCounts) {
		toSerialize["errataApplicabilityCounts"] = o.ErrataApplicabilityCounts
	}
	if !IsNil(o.FactsCount) {
		toSerialize["factsCount"] = o.FactsCount
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.InstalledProductsCount) {
		toSerialize["installedProductsCount"] = o.InstalledProductsCount
	}
	if !IsNil(o.LastCheckin) {
		toSerialize["lastCheckin"] = o.LastCheckin
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServiceLevelPreference) {
		toSerialize["serviceLevelPreference"] = o.ServiceLevelPreference
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableAttachEntitlement struct {
	value *AttachEntitlement
	isSet bool
}

func (v NullableAttachEntitlement) Get() *AttachEntitlement {
	return v.value
}

func (v *NullableAttachEntitlement) Set(val *AttachEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachEntitlement(val *AttachEntitlement) *NullableAttachEntitlement {
	return &NullableAttachEntitlement{value: val, isSet: true}
}

func (v NullableAttachEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


