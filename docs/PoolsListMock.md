# PoolsListMock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]PoolDetail**](PoolDetail.md) |  | [optional] 
**Pagination** | Pointer to [**APIPageParam**](APIPageParam.md) |  | [optional] 

## Methods

### NewPoolsListMock

`func NewPoolsListMock() *PoolsListMock`

NewPoolsListMock instantiates a new PoolsListMock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolsListMockWithDefaults

`func NewPoolsListMockWithDefaults() *PoolsListMock`

NewPoolsListMockWithDefaults instantiates a new PoolsListMock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *PoolsListMock) GetBody() []PoolDetail`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PoolsListMock) GetBodyOk() (*[]PoolDetail, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PoolsListMock) SetBody(v []PoolDetail)`

SetBody sets Body field to given value.

### HasBody

`func (o *PoolsListMock) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPagination

`func (o *PoolsListMock) GetPagination() APIPageParam`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PoolsListMock) GetPaginationOk() (*APIPageParam, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PoolsListMock) SetPagination(v APIPageParam)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PoolsListMock) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


