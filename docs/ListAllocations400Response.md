# ListAllocations400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorDetails**](ErrorDetails.md) |  | [optional] 

## Methods

### NewListAllocations400Response

`func NewListAllocations400Response() *ListAllocations400Response`

NewListAllocations400Response instantiates a new ListAllocations400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllocations400ResponseWithDefaults

`func NewListAllocations400ResponseWithDefaults() *ListAllocations400Response`

NewListAllocations400ResponseWithDefaults instantiates a new ListAllocations400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ListAllocations400Response) GetError() ErrorDetails`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListAllocations400Response) GetErrorOk() (*ErrorDetails, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListAllocations400Response) SetError(v ErrorDetails)`

SetError sets Error field to given value.

### HasError

`func (o *ListAllocations400Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


