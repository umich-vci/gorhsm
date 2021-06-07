# EnabledCloudAccessProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]EnabledProviderAccount**](EnabledProviderAccount.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Products** | Pointer to [**[]EnabledProduct**](EnabledProduct.md) |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 

## Methods

### NewEnabledCloudAccessProvider

`func NewEnabledCloudAccessProvider() *EnabledCloudAccessProvider`

NewEnabledCloudAccessProvider instantiates a new EnabledCloudAccessProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnabledCloudAccessProviderWithDefaults

`func NewEnabledCloudAccessProviderWithDefaults() *EnabledCloudAccessProvider`

NewEnabledCloudAccessProviderWithDefaults instantiates a new EnabledCloudAccessProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *EnabledCloudAccessProvider) GetAccounts() []EnabledProviderAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *EnabledCloudAccessProvider) GetAccountsOk() (*[]EnabledProviderAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *EnabledCloudAccessProvider) SetAccounts(v []EnabledProviderAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *EnabledCloudAccessProvider) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetName

`func (o *EnabledCloudAccessProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnabledCloudAccessProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnabledCloudAccessProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnabledCloudAccessProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProducts

`func (o *EnabledCloudAccessProvider) GetProducts() []EnabledProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *EnabledCloudAccessProvider) GetProductsOk() (*[]EnabledProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *EnabledCloudAccessProvider) SetProducts(v []EnabledProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *EnabledCloudAccessProvider) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetShortName

`func (o *EnabledCloudAccessProvider) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *EnabledCloudAccessProvider) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *EnabledCloudAccessProvider) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *EnabledCloudAccessProvider) HasShortName() bool`

HasShortName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


