# ListSubContentSets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]ContentSet**](ContentSet.md) |  | [optional] 
**Pagination** | Pointer to [**APIPageParam**](APIPageParam.md) |  | [optional] 

## Methods

### NewListSubContentSets200Response

`func NewListSubContentSets200Response() *ListSubContentSets200Response`

NewListSubContentSets200Response instantiates a new ListSubContentSets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubContentSets200ResponseWithDefaults

`func NewListSubContentSets200ResponseWithDefaults() *ListSubContentSets200Response`

NewListSubContentSets200ResponseWithDefaults instantiates a new ListSubContentSets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ListSubContentSets200Response) GetBody() []ContentSet`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ListSubContentSets200Response) GetBodyOk() (*[]ContentSet, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ListSubContentSets200Response) SetBody(v []ContentSet)`

SetBody sets Body field to given value.

### HasBody

`func (o *ListSubContentSets200Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPagination

`func (o *ListSubContentSets200Response) GetPagination() APIPageParam`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSubContentSets200Response) GetPaginationOk() (*APIPageParam, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSubContentSets200Response) SetPagination(v APIPageParam)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSubContentSets200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


