# ListSystemErrata200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]ErratumForSystem**](ErratumForSystem.md) |  | [optional] 
**Pagination** | Pointer to [**APIPageParam**](APIPageParam.md) |  | [optional] 

## Methods

### NewListSystemErrata200Response

`func NewListSystemErrata200Response() *ListSystemErrata200Response`

NewListSystemErrata200Response instantiates a new ListSystemErrata200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSystemErrata200ResponseWithDefaults

`func NewListSystemErrata200ResponseWithDefaults() *ListSystemErrata200Response`

NewListSystemErrata200ResponseWithDefaults instantiates a new ListSystemErrata200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ListSystemErrata200Response) GetBody() []ErratumForSystem`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ListSystemErrata200Response) GetBodyOk() (*[]ErratumForSystem, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ListSystemErrata200Response) SetBody(v []ErratumForSystem)`

SetBody sets Body field to given value.

### HasBody

`func (o *ListSystemErrata200Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPagination

`func (o *ListSystemErrata200Response) GetPagination() APIPageParam`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSystemErrata200Response) GetPaginationOk() (*APIPageParam, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSystemErrata200Response) SetPagination(v APIPageParam)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSystemErrata200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


