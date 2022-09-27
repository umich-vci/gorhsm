# ListSystems200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]System**](System.md) | systemList is a System Slice | [optional] 
**Pagination** | Pointer to [**APIPageParam**](APIPageParam.md) |  | [optional] 

## Methods

### NewListSystems200Response

`func NewListSystems200Response() *ListSystems200Response`

NewListSystems200Response instantiates a new ListSystems200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSystems200ResponseWithDefaults

`func NewListSystems200ResponseWithDefaults() *ListSystems200Response`

NewListSystems200ResponseWithDefaults instantiates a new ListSystems200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ListSystems200Response) GetBody() []System`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ListSystems200Response) GetBodyOk() (*[]System, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ListSystems200Response) SetBody(v []System)`

SetBody sets Body field to given value.

### HasBody

`func (o *ListSystems200Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPagination

`func (o *ListSystems200Response) GetPagination() APIPageParam`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSystems200Response) GetPaginationOk() (*APIPageParam, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSystems200Response) SetPagination(v APIPageParam)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSystems200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


