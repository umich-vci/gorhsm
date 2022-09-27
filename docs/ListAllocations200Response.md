# ListAllocations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]Allocation**](Allocation.md) |  | [optional] 
**Pagination** | Pointer to [**APIPageParam**](APIPageParam.md) |  | [optional] 

## Methods

### NewListAllocations200Response

`func NewListAllocations200Response() *ListAllocations200Response`

NewListAllocations200Response instantiates a new ListAllocations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllocations200ResponseWithDefaults

`func NewListAllocations200ResponseWithDefaults() *ListAllocations200Response`

NewListAllocations200ResponseWithDefaults instantiates a new ListAllocations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ListAllocations200Response) GetBody() []Allocation`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ListAllocations200Response) GetBodyOk() (*[]Allocation, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ListAllocations200Response) SetBody(v []Allocation)`

SetBody sets Body field to given value.

### HasBody

`func (o *ListAllocations200Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPagination

`func (o *ListAllocations200Response) GetPagination() APIPageParam`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAllocations200Response) GetPaginationOk() (*APIPageParam, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAllocations200Response) SetPagination(v APIPageParam)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAllocations200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


