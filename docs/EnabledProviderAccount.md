# EnabledProviderAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateAdded** | **string** | Date represents the date format used for API returns | 
**GoldImageStatus** | Pointer to [**[]GoldImageStatus**](GoldImageStatus.md) |  | [optional] 
**Id** | **string** |  | 
**Nickname** | **string** |  | 
**SourceId** | Pointer to **string** | Source ID of linked account (only for accounts created via Sources on cloud.redhat.com) | [optional] 
**Verified** | Pointer to **bool** | verification status for RHSM Auto Registration (only displayed for supported cloud providers) | [optional] 

## Methods

### NewEnabledProviderAccount

`func NewEnabledProviderAccount(dateAdded string, id string, nickname string, ) *EnabledProviderAccount`

NewEnabledProviderAccount instantiates a new EnabledProviderAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnabledProviderAccountWithDefaults

`func NewEnabledProviderAccountWithDefaults() *EnabledProviderAccount`

NewEnabledProviderAccountWithDefaults instantiates a new EnabledProviderAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateAdded

`func (o *EnabledProviderAccount) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *EnabledProviderAccount) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *EnabledProviderAccount) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.


### GetGoldImageStatus

`func (o *EnabledProviderAccount) GetGoldImageStatus() []GoldImageStatus`

GetGoldImageStatus returns the GoldImageStatus field if non-nil, zero value otherwise.

### GetGoldImageStatusOk

`func (o *EnabledProviderAccount) GetGoldImageStatusOk() (*[]GoldImageStatus, bool)`

GetGoldImageStatusOk returns a tuple with the GoldImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoldImageStatus

`func (o *EnabledProviderAccount) SetGoldImageStatus(v []GoldImageStatus)`

SetGoldImageStatus sets GoldImageStatus field to given value.

### HasGoldImageStatus

`func (o *EnabledProviderAccount) HasGoldImageStatus() bool`

HasGoldImageStatus returns a boolean if a field has been set.

### GetId

`func (o *EnabledProviderAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnabledProviderAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnabledProviderAccount) SetId(v string)`

SetId sets Id field to given value.


### GetNickname

`func (o *EnabledProviderAccount) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *EnabledProviderAccount) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *EnabledProviderAccount) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetSourceId

`func (o *EnabledProviderAccount) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *EnabledProviderAccount) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *EnabledProviderAccount) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *EnabledProviderAccount) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetVerified

`func (o *EnabledProviderAccount) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *EnabledProviderAccount) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *EnabledProviderAccount) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *EnabledProviderAccount) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


