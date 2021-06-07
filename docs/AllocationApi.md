# \AllocationApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachEntitlementAllocation**](AllocationApi.md#AttachEntitlementAllocation) | **Post** /allocations/{AllocationUUID}/entitlements | Attach entitlement to allocation
[**CreateSatellite**](AllocationApi.md#CreateSatellite) | **Post** /allocations | Create Satellite
[**ExportAllocation**](AllocationApi.md#ExportAllocation) | **Get** /allocations/{AllocationUUID}/export | Trigger allocation manifest export
[**ExportJobAllocation**](AllocationApi.md#ExportJobAllocation) | **Get** /allocations/{AllocationUUID}/exportJob/{ExportJobID} | Check status of allocation manifest export
[**GetExportAllocation**](AllocationApi.md#GetExportAllocation) | **Get** /allocations/{AllocationUUID}/export/{ExportID} | Download allocation manifest
[**ListAllocationPools**](AllocationApi.md#ListAllocationPools) | **Get** /allocations/{AllocationUUID}/pools | List all pools for an allocation
[**ListAllocations**](AllocationApi.md#ListAllocations) | **Get** /allocations | List all allocations for a user
[**ListVersionsAllocation**](AllocationApi.md#ListVersionsAllocation) | **Get** /allocations/versions | List Satellite versions
[**RemoveAllocation**](AllocationApi.md#RemoveAllocation) | **Delete** /allocations/{AllocationUUID} | Remove allocation profile
[**RemoveAllocationEntitlement**](AllocationApi.md#RemoveAllocationEntitlement) | **Delete** /allocations/{AllocationUUID}/{EntitlementID} | Remove entitlement from the allocation
[**ShowAllocation**](AllocationApi.md#ShowAllocation) | **Get** /allocations/{AllocationUUID} | Get an allocation by UUID
[**UpdateAllocation**](AllocationApi.md#UpdateAllocation) | **Put** /allocations/{AllocationUUID} | Update an allocation
[**UpdateEntitlementAllocation**](AllocationApi.md#UpdateEntitlementAllocation) | **Put** /allocations/{AllocationUUID}/entitlements/{EntitlementUUID} | Update attached entitlement to allocation



## AttachEntitlementAllocation

> InlineResponse2003 AttachEntitlementAllocation(ctx, allocationUUID).Pool(pool).Quantity(quantity).Execute()

Attach entitlement to allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pool := "pool_example" // string | 
    allocationUUID := "allocationUUID_example" // string | 
    quantity := int32(56) // int32 | quantity you want to attach (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.AttachEntitlementAllocation(context.Background(), allocationUUID).Pool(pool).Quantity(quantity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.AttachEntitlementAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachEntitlementAllocation`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.AttachEntitlementAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachEntitlementAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pool** | **string** |  | 

 **quantity** | **int32** | quantity you want to attach | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSatellite

> InlineResponse2001 CreateSatellite(ctx).Name(name).Version(version).Execute()

Create Satellite



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | must be less than 100 characters and use only numbers, letters, underscores, hyphens, and periods
    version := "version_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.CreateSatellite(context.Background()).Name(name).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.CreateSatellite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSatellite`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.CreateSatellite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSatelliteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | must be less than 100 characters and use only numbers, letters, underscores, hyphens, and periods | 
 **version** | **string** |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAllocation

> InlineResponse2004 ExportAllocation(ctx, allocationUUID).Execute()

Trigger allocation manifest export



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allocationUUID := "allocationUUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.ExportAllocation(context.Background(), allocationUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ExportAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAllocation`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ExportAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportJobAllocation

> InlineResponse2005 ExportJobAllocation(ctx, allocationUUID, exportJobID).Execute()

Check status of allocation manifest export



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allocationUUID := "allocationUUID_example" // string | 
    exportJobID := "exportJobID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.ExportJobAllocation(context.Background(), allocationUUID, exportJobID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ExportJobAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportJobAllocation`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ExportJobAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 
**exportJobID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportJobAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportAllocation

> []int32 GetExportAllocation(ctx, allocationUUID, exportID).Execute()

Download allocation manifest



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allocationUUID := "allocationUUID_example" // string | 
    exportID := "exportID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.GetExportAllocation(context.Background(), allocationUUID, exportID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.GetExportAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportAllocation`: []int32
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.GetExportAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 
**exportID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]int32**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllocationPools

> PoolsListMock ListAllocationPools(ctx, allocationUUID).Limit(limit).Offset(offset).Future(future).Execute()

List all pools for an allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allocationUUID := "allocationUUID_example" // string | 
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)
    future := true // bool | include future dated pools for satellite 6.3 or higher (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.ListAllocationPools(context.Background(), allocationUUID).Limit(limit).Offset(offset).Future(future).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ListAllocationPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllocationPools`: PoolsListMock
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ListAllocationPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllocationPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 
 **future** | **bool** | include future dated pools for satellite 6.3 or higher | 

### Return type

[**PoolsListMock**](PoolsListMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllocations

> InlineResponse200 ListAllocations(ctx).Limit(limit).Offset(offset).Type_(type_).Execute()

List all allocations for a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.ListAllocations(context.Background()).Limit(limit).Offset(offset).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ListAllocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllocations`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ListAllocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 
 **type_** | **string** |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersionsAllocation

> InlineResponse2002 ListVersionsAllocation(ctx).Execute()

List Satellite versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.ListVersionsAllocation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ListVersionsAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVersionsAllocation`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ListVersionsAllocation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionsAllocationRequest struct via the builder pattern


### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAllocation

> RemoveAllocation(ctx, allocationUUID).Force(force).Execute()

Remove allocation profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allocationUUID := "allocationUUID_example" // string | 
    force := true // bool | Deleting a subscription allocation can have significant impacts on your hosts and activation keys. We require a force parameter to make sure the delete operation is intentional.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.RemoveAllocation(context.Background(), allocationUUID).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.RemoveAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Deleting a subscription allocation can have significant impacts on your hosts and activation keys. We require a force parameter to make sure the delete operation is intentional. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAllocationEntitlement

> RemoveAllocationEntitlement(ctx, allocationUUID, entitlementID).Execute()

Remove entitlement from the allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allocationUUID := "allocationUUID_example" // string | 
    entitlementID := "entitlementID_example" // string | Remove an entitlement from an allocation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.RemoveAllocationEntitlement(context.Background(), allocationUUID, entitlementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.RemoveAllocationEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 
**entitlementID** | **string** | Remove an entitlement from an allocation | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAllocationEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAllocation

> InlineResponse2003 ShowAllocation(ctx, allocationUUID).Include(include).Execute()

Get an allocation by UUID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allocationUUID := "allocationUUID_example" // string | 
    include := "include_example" // string | Show more details about a allocation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.ShowAllocation(context.Background(), allocationUUID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ShowAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowAllocation`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ShowAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Show more details about a allocation | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAllocation

> UpdateAllocation(ctx, allocationUUID).Allocation(allocation).Execute()

Update an allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allocationUUID := "allocationUUID_example" // string | 
    allocation := *openapiclient.NewInlineObject("SimpleContentAccess_example") // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.UpdateAllocation(context.Background(), allocationUUID).Allocation(allocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.UpdateAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allocation** | [**InlineObject**](InlineObject.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitlementAllocation

> InlineResponse2003 UpdateEntitlementAllocation(ctx, allocationUUID, entitlementUUID).Quantity(quantity).Execute()

Update attached entitlement to allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    allocationUUID := "allocationUUID_example" // string | 
    entitlementUUID := "entitlementUUID_example" // string | 
    quantity := int32(56) // int32 | maxItem: quantity must be less than or equal to the maximum number of allowed entitlements in the entitlement pool minItem: 1 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllocationApi.UpdateEntitlementAllocation(context.Background(), allocationUUID, entitlementUUID).Quantity(quantity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.UpdateEntitlementAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntitlementAllocation`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.UpdateEntitlementAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string** |  | 
**entitlementUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitlementAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **quantity** | **int32** | maxItem: quantity must be less than or equal to the maximum number of allowed entitlements in the entitlement pool minItem: 1 | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

