# InlineResponse20015

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]System**](System.md) | systemList is a System Slice | [optional] 
**Pagination** | Pointer to [**APIPageParam**](APIPageParam.md) |  | [optional] 

## Methods

### NewInlineResponse20015

`func NewInlineResponse20015() *InlineResponse20015`

NewInlineResponse20015 instantiates a new InlineResponse20015 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20015WithDefaults

`func NewInlineResponse20015WithDefaults() *InlineResponse20015`

NewInlineResponse20015WithDefaults instantiates a new InlineResponse20015 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *InlineResponse20015) GetBody() []System`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineResponse20015) GetBodyOk() (*[]System, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineResponse20015) SetBody(v []System)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineResponse20015) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPagination

`func (o *InlineResponse20015) GetPagination() APIPageParam`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InlineResponse20015) GetPaginationOk() (*APIPageParam, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InlineResponse20015) SetPagination(v APIPageParam)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *InlineResponse20015) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


