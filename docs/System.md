# System

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementCount** | Pointer to **int32** |  | [optional] 
**EntitlementStatus** | Pointer to **string** |  | [optional] 
**ErrataCounts** | Pointer to [**ErrataCount**](ErrataCount.md) |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**LastCheckin** | Pointer to **string** | Date represents the date format used for API returns | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewSystem

`func NewSystem() *System`

NewSystem instantiates a new System object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemWithDefaults

`func NewSystemWithDefaults() *System`

NewSystemWithDefaults instantiates a new System object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementCount

`func (o *System) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *System) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *System) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *System) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetEntitlementStatus

`func (o *System) GetEntitlementStatus() string`

GetEntitlementStatus returns the EntitlementStatus field if non-nil, zero value otherwise.

### GetEntitlementStatusOk

`func (o *System) GetEntitlementStatusOk() (*string, bool)`

GetEntitlementStatusOk returns a tuple with the EntitlementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementStatus

`func (o *System) SetEntitlementStatus(v string)`

SetEntitlementStatus sets EntitlementStatus field to given value.

### HasEntitlementStatus

`func (o *System) HasEntitlementStatus() bool`

HasEntitlementStatus returns a boolean if a field has been set.

### GetErrataCounts

`func (o *System) GetErrataCounts() ErrataCount`

GetErrataCounts returns the ErrataCounts field if non-nil, zero value otherwise.

### GetErrataCountsOk

`func (o *System) GetErrataCountsOk() (*ErrataCount, bool)`

GetErrataCountsOk returns a tuple with the ErrataCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrataCounts

`func (o *System) SetErrataCounts(v ErrataCount)`

SetErrataCounts sets ErrataCounts field to given value.

### HasErrataCounts

`func (o *System) HasErrataCounts() bool`

HasErrataCounts returns a boolean if a field has been set.

### GetHostname

`func (o *System) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *System) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *System) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *System) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHref

`func (o *System) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *System) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *System) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *System) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLastCheckin

`func (o *System) GetLastCheckin() string`

GetLastCheckin returns the LastCheckin field if non-nil, zero value otherwise.

### GetLastCheckinOk

`func (o *System) GetLastCheckinOk() (*string, bool)`

GetLastCheckinOk returns a tuple with the LastCheckin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckin

`func (o *System) SetLastCheckin(v string)`

SetLastCheckin sets LastCheckin field to given value.

### HasLastCheckin

`func (o *System) HasLastCheckin() bool`

HasLastCheckin returns a boolean if a field has been set.

### GetName

`func (o *System) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *System) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *System) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *System) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *System) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *System) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *System) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *System) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *System) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *System) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *System) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *System) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *System) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *System) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *System) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *System) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


