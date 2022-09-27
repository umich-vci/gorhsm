# ListSubscriptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]DetailResponse**](DetailResponse.md) | ListResponse is the actual collection of subscription details that gets rendered | [optional] 
**Pagination** | Pointer to [**APIPageParam**](APIPageParam.md) |  | [optional] 

## Methods

### NewListSubscriptions200Response

`func NewListSubscriptions200Response() *ListSubscriptions200Response`

NewListSubscriptions200Response instantiates a new ListSubscriptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptions200ResponseWithDefaults

`func NewListSubscriptions200ResponseWithDefaults() *ListSubscriptions200Response`

NewListSubscriptions200ResponseWithDefaults instantiates a new ListSubscriptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ListSubscriptions200Response) GetBody() []DetailResponse`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ListSubscriptions200Response) GetBodyOk() (*[]DetailResponse, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ListSubscriptions200Response) SetBody(v []DetailResponse)`

SetBody sets Body field to given value.

### HasBody

`func (o *ListSubscriptions200Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPagination

`func (o *ListSubscriptions200Response) GetPagination() APIPageParam`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSubscriptions200Response) GetPaginationOk() (*APIPageParam, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSubscriptions200Response) SetPagination(v APIPageParam)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSubscriptions200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


