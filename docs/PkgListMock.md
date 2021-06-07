# PkgListMock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]PackageDetail**](PackageDetail.md) |  | [optional] 
**Pagination** | Pointer to [**APIPageParam**](APIPageParam.md) |  | [optional] 

## Methods

### NewPkgListMock

`func NewPkgListMock() *PkgListMock`

NewPkgListMock instantiates a new PkgListMock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgListMockWithDefaults

`func NewPkgListMockWithDefaults() *PkgListMock`

NewPkgListMockWithDefaults instantiates a new PkgListMock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *PkgListMock) GetBody() []PackageDetail`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PkgListMock) GetBodyOk() (*[]PackageDetail, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PkgListMock) SetBody(v []PackageDetail)`

SetBody sets Body field to given value.

### HasBody

`func (o *PkgListMock) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPagination

`func (o *PkgListMock) GetPagination() APIPageParam`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PkgListMock) GetPaginationOk() (*APIPageParam, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PkgListMock) SetPagination(v APIPageParam)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PkgListMock) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


