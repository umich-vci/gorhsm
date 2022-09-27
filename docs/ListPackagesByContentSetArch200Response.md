# ListPackagesByContentSetArch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]PkgContentSetArch**](PkgContentSetArch.md) |  | [optional] 
**Pagination** | Pointer to [**APIPageParam**](APIPageParam.md) |  | [optional] 

## Methods

### NewListPackagesByContentSetArch200Response

`func NewListPackagesByContentSetArch200Response() *ListPackagesByContentSetArch200Response`

NewListPackagesByContentSetArch200Response instantiates a new ListPackagesByContentSetArch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPackagesByContentSetArch200ResponseWithDefaults

`func NewListPackagesByContentSetArch200ResponseWithDefaults() *ListPackagesByContentSetArch200Response`

NewListPackagesByContentSetArch200ResponseWithDefaults instantiates a new ListPackagesByContentSetArch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ListPackagesByContentSetArch200Response) GetBody() []PkgContentSetArch`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ListPackagesByContentSetArch200Response) GetBodyOk() (*[]PkgContentSetArch, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ListPackagesByContentSetArch200Response) SetBody(v []PkgContentSetArch)`

SetBody sets Body field to given value.

### HasBody

`func (o *ListPackagesByContentSetArch200Response) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPagination

`func (o *ListPackagesByContentSetArch200Response) GetPagination() APIPageParam`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListPackagesByContentSetArch200Response) GetPaginationOk() (*APIPageParam, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListPackagesByContentSetArch200Response) SetPagination(v APIPageParam)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListPackagesByContentSetArch200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


