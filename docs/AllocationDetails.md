# AllocationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** | Date represents the date format used for API returns | [optional] 
**EntitlementsAttached** | Pointer to [**EntitlementsAttachedResponse**](EntitlementsAttachedResponse.md) |  | [optional] 
**EntitlementsAttachedQuantity** | Pointer to **int32** |  | [optional] 
**LastModified** | Pointer to **string** | Date represents the date format used for API returns | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SimpleContentAccess** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAllocationDetails

`func NewAllocationDetails() *AllocationDetails`

NewAllocationDetails instantiates a new AllocationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationDetailsWithDefaults

`func NewAllocationDetailsWithDefaults() *AllocationDetails`

NewAllocationDetailsWithDefaults instantiates a new AllocationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *AllocationDetails) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AllocationDetails) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AllocationDetails) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AllocationDetails) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *AllocationDetails) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AllocationDetails) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AllocationDetails) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AllocationDetails) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetEntitlementsAttached

`func (o *AllocationDetails) GetEntitlementsAttached() EntitlementsAttachedResponse`

GetEntitlementsAttached returns the EntitlementsAttached field if non-nil, zero value otherwise.

### GetEntitlementsAttachedOk

`func (o *AllocationDetails) GetEntitlementsAttachedOk() (*EntitlementsAttachedResponse, bool)`

GetEntitlementsAttachedOk returns a tuple with the EntitlementsAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsAttached

`func (o *AllocationDetails) SetEntitlementsAttached(v EntitlementsAttachedResponse)`

SetEntitlementsAttached sets EntitlementsAttached field to given value.

### HasEntitlementsAttached

`func (o *AllocationDetails) HasEntitlementsAttached() bool`

HasEntitlementsAttached returns a boolean if a field has been set.

### GetEntitlementsAttachedQuantity

`func (o *AllocationDetails) GetEntitlementsAttachedQuantity() int32`

GetEntitlementsAttachedQuantity returns the EntitlementsAttachedQuantity field if non-nil, zero value otherwise.

### GetEntitlementsAttachedQuantityOk

`func (o *AllocationDetails) GetEntitlementsAttachedQuantityOk() (*int32, bool)`

GetEntitlementsAttachedQuantityOk returns a tuple with the EntitlementsAttachedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsAttachedQuantity

`func (o *AllocationDetails) SetEntitlementsAttachedQuantity(v int32)`

SetEntitlementsAttachedQuantity sets EntitlementsAttachedQuantity field to given value.

### HasEntitlementsAttachedQuantity

`func (o *AllocationDetails) HasEntitlementsAttachedQuantity() bool`

HasEntitlementsAttachedQuantity returns a boolean if a field has been set.

### GetLastModified

`func (o *AllocationDetails) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AllocationDetails) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AllocationDetails) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AllocationDetails) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetName

`func (o *AllocationDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllocationDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllocationDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllocationDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSimpleContentAccess

`func (o *AllocationDetails) GetSimpleContentAccess() string`

GetSimpleContentAccess returns the SimpleContentAccess field if non-nil, zero value otherwise.

### GetSimpleContentAccessOk

`func (o *AllocationDetails) GetSimpleContentAccessOk() (*string, bool)`

GetSimpleContentAccessOk returns a tuple with the SimpleContentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleContentAccess

`func (o *AllocationDetails) SetSimpleContentAccess(v string)`

SetSimpleContentAccess sets SimpleContentAccess field to given value.

### HasSimpleContentAccess

`func (o *AllocationDetails) HasSimpleContentAccess() bool`

HasSimpleContentAccess returns a boolean if a field has been set.

### GetType

`func (o *AllocationDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AllocationDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AllocationDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AllocationDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *AllocationDetails) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AllocationDetails) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AllocationDetails) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AllocationDetails) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *AllocationDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AllocationDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AllocationDetails) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AllocationDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


