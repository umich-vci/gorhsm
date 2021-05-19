# GoldImageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the provider image group | [optional] 
**Name** | Pointer to **string** | Name of the requested provider image group | [optional] 
**Status** | Pointer to **string** | Status of Gold Image Request | [optional] 

## Methods

### NewGoldImageStatus

`func NewGoldImageStatus() *GoldImageStatus`

NewGoldImageStatus instantiates a new GoldImageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoldImageStatusWithDefaults

`func NewGoldImageStatusWithDefaults() *GoldImageStatus`

NewGoldImageStatusWithDefaults instantiates a new GoldImageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GoldImageStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GoldImageStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GoldImageStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GoldImageStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *GoldImageStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoldImageStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoldImageStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GoldImageStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *GoldImageStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GoldImageStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GoldImageStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GoldImageStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


