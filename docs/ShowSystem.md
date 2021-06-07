# ShowSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoAttachSetting** | Pointer to **bool** |  | [optional] 
**ComplianceStatus** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** | Date represents the date format used for API returns | [optional] 
**EntitlementStatus** | Pointer to **string** |  | [optional] 
**EntitlementsAttached** | Pointer to [**EntitlementsAttachedResponse**](EntitlementsAttachedResponse.md) |  | [optional] 
**EntitlementsAttachedCount** | Pointer to **int32** |  | [optional] 
**ErrataApplicabilityCounts** | Pointer to [**ErrataApplicabilityCounts**](ErrataApplicabilityCounts.md) |  | [optional] 
**Facts** | Pointer to [**Facts**](Facts.md) |  | [optional] 
**FactsCount** | Pointer to **int32** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**InstalledProducts** | Pointer to [**InstalledProducts**](InstalledProducts.md) |  | [optional] 
**InstalledProductsCount** | Pointer to **int32** |  | [optional] 
**LastCheckin** | Pointer to **string** | Date represents the date format used for API returns | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServiceLevelPreference** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewShowSystem

`func NewShowSystem() *ShowSystem`

NewShowSystem instantiates a new ShowSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowSystemWithDefaults

`func NewShowSystemWithDefaults() *ShowSystem`

NewShowSystemWithDefaults instantiates a new ShowSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoAttachSetting

`func (o *ShowSystem) GetAutoAttachSetting() bool`

GetAutoAttachSetting returns the AutoAttachSetting field if non-nil, zero value otherwise.

### GetAutoAttachSettingOk

`func (o *ShowSystem) GetAutoAttachSettingOk() (*bool, bool)`

GetAutoAttachSettingOk returns a tuple with the AutoAttachSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAttachSetting

`func (o *ShowSystem) SetAutoAttachSetting(v bool)`

SetAutoAttachSetting sets AutoAttachSetting field to given value.

### HasAutoAttachSetting

`func (o *ShowSystem) HasAutoAttachSetting() bool`

HasAutoAttachSetting returns a boolean if a field has been set.

### GetComplianceStatus

`func (o *ShowSystem) GetComplianceStatus() string`

GetComplianceStatus returns the ComplianceStatus field if non-nil, zero value otherwise.

### GetComplianceStatusOk

`func (o *ShowSystem) GetComplianceStatusOk() (*string, bool)`

GetComplianceStatusOk returns a tuple with the ComplianceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceStatus

`func (o *ShowSystem) SetComplianceStatus(v string)`

SetComplianceStatus sets ComplianceStatus field to given value.

### HasComplianceStatus

`func (o *ShowSystem) HasComplianceStatus() bool`

HasComplianceStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ShowSystem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ShowSystem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ShowSystem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ShowSystem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ShowSystem) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ShowSystem) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ShowSystem) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ShowSystem) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetEntitlementStatus

`func (o *ShowSystem) GetEntitlementStatus() string`

GetEntitlementStatus returns the EntitlementStatus field if non-nil, zero value otherwise.

### GetEntitlementStatusOk

`func (o *ShowSystem) GetEntitlementStatusOk() (*string, bool)`

GetEntitlementStatusOk returns a tuple with the EntitlementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementStatus

`func (o *ShowSystem) SetEntitlementStatus(v string)`

SetEntitlementStatus sets EntitlementStatus field to given value.

### HasEntitlementStatus

`func (o *ShowSystem) HasEntitlementStatus() bool`

HasEntitlementStatus returns a boolean if a field has been set.

### GetEntitlementsAttached

`func (o *ShowSystem) GetEntitlementsAttached() EntitlementsAttachedResponse`

GetEntitlementsAttached returns the EntitlementsAttached field if non-nil, zero value otherwise.

### GetEntitlementsAttachedOk

`func (o *ShowSystem) GetEntitlementsAttachedOk() (*EntitlementsAttachedResponse, bool)`

GetEntitlementsAttachedOk returns a tuple with the EntitlementsAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsAttached

`func (o *ShowSystem) SetEntitlementsAttached(v EntitlementsAttachedResponse)`

SetEntitlementsAttached sets EntitlementsAttached field to given value.

### HasEntitlementsAttached

`func (o *ShowSystem) HasEntitlementsAttached() bool`

HasEntitlementsAttached returns a boolean if a field has been set.

### GetEntitlementsAttachedCount

`func (o *ShowSystem) GetEntitlementsAttachedCount() int32`

GetEntitlementsAttachedCount returns the EntitlementsAttachedCount field if non-nil, zero value otherwise.

### GetEntitlementsAttachedCountOk

`func (o *ShowSystem) GetEntitlementsAttachedCountOk() (*int32, bool)`

GetEntitlementsAttachedCountOk returns a tuple with the EntitlementsAttachedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsAttachedCount

`func (o *ShowSystem) SetEntitlementsAttachedCount(v int32)`

SetEntitlementsAttachedCount sets EntitlementsAttachedCount field to given value.

### HasEntitlementsAttachedCount

`func (o *ShowSystem) HasEntitlementsAttachedCount() bool`

HasEntitlementsAttachedCount returns a boolean if a field has been set.

### GetErrataApplicabilityCounts

`func (o *ShowSystem) GetErrataApplicabilityCounts() ErrataApplicabilityCounts`

GetErrataApplicabilityCounts returns the ErrataApplicabilityCounts field if non-nil, zero value otherwise.

### GetErrataApplicabilityCountsOk

`func (o *ShowSystem) GetErrataApplicabilityCountsOk() (*ErrataApplicabilityCounts, bool)`

GetErrataApplicabilityCountsOk returns a tuple with the ErrataApplicabilityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrataApplicabilityCounts

`func (o *ShowSystem) SetErrataApplicabilityCounts(v ErrataApplicabilityCounts)`

SetErrataApplicabilityCounts sets ErrataApplicabilityCounts field to given value.

### HasErrataApplicabilityCounts

`func (o *ShowSystem) HasErrataApplicabilityCounts() bool`

HasErrataApplicabilityCounts returns a boolean if a field has been set.

### GetFacts

`func (o *ShowSystem) GetFacts() Facts`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *ShowSystem) GetFactsOk() (*Facts, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *ShowSystem) SetFacts(v Facts)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *ShowSystem) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetFactsCount

`func (o *ShowSystem) GetFactsCount() int32`

GetFactsCount returns the FactsCount field if non-nil, zero value otherwise.

### GetFactsCountOk

`func (o *ShowSystem) GetFactsCountOk() (*int32, bool)`

GetFactsCountOk returns a tuple with the FactsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactsCount

`func (o *ShowSystem) SetFactsCount(v int32)`

SetFactsCount sets FactsCount field to given value.

### HasFactsCount

`func (o *ShowSystem) HasFactsCount() bool`

HasFactsCount returns a boolean if a field has been set.

### GetHostname

`func (o *ShowSystem) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ShowSystem) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ShowSystem) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ShowSystem) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInstalledProducts

`func (o *ShowSystem) GetInstalledProducts() InstalledProducts`

GetInstalledProducts returns the InstalledProducts field if non-nil, zero value otherwise.

### GetInstalledProductsOk

`func (o *ShowSystem) GetInstalledProductsOk() (*InstalledProducts, bool)`

GetInstalledProductsOk returns a tuple with the InstalledProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProducts

`func (o *ShowSystem) SetInstalledProducts(v InstalledProducts)`

SetInstalledProducts sets InstalledProducts field to given value.

### HasInstalledProducts

`func (o *ShowSystem) HasInstalledProducts() bool`

HasInstalledProducts returns a boolean if a field has been set.

### GetInstalledProductsCount

`func (o *ShowSystem) GetInstalledProductsCount() int32`

GetInstalledProductsCount returns the InstalledProductsCount field if non-nil, zero value otherwise.

### GetInstalledProductsCountOk

`func (o *ShowSystem) GetInstalledProductsCountOk() (*int32, bool)`

GetInstalledProductsCountOk returns a tuple with the InstalledProductsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProductsCount

`func (o *ShowSystem) SetInstalledProductsCount(v int32)`

SetInstalledProductsCount sets InstalledProductsCount field to given value.

### HasInstalledProductsCount

`func (o *ShowSystem) HasInstalledProductsCount() bool`

HasInstalledProductsCount returns a boolean if a field has been set.

### GetLastCheckin

`func (o *ShowSystem) GetLastCheckin() string`

GetLastCheckin returns the LastCheckin field if non-nil, zero value otherwise.

### GetLastCheckinOk

`func (o *ShowSystem) GetLastCheckinOk() (*string, bool)`

GetLastCheckinOk returns a tuple with the LastCheckin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckin

`func (o *ShowSystem) SetLastCheckin(v string)`

SetLastCheckin sets LastCheckin field to given value.

### HasLastCheckin

`func (o *ShowSystem) HasLastCheckin() bool`

HasLastCheckin returns a boolean if a field has been set.

### GetName

`func (o *ShowSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShowSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShowSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShowSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceLevelPreference

`func (o *ShowSystem) GetServiceLevelPreference() string`

GetServiceLevelPreference returns the ServiceLevelPreference field if non-nil, zero value otherwise.

### GetServiceLevelPreferenceOk

`func (o *ShowSystem) GetServiceLevelPreferenceOk() (*string, bool)`

GetServiceLevelPreferenceOk returns a tuple with the ServiceLevelPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelPreference

`func (o *ShowSystem) SetServiceLevelPreference(v string)`

SetServiceLevelPreference sets ServiceLevelPreference field to given value.

### HasServiceLevelPreference

`func (o *ShowSystem) HasServiceLevelPreference() bool`

HasServiceLevelPreference returns a boolean if a field has been set.

### GetType

`func (o *ShowSystem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShowSystem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShowSystem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShowSystem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *ShowSystem) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ShowSystem) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ShowSystem) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ShowSystem) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


