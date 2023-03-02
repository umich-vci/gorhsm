# \AllocationApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachEntitlementAllocation**](AllocationApi.md#AttachEntitlementAllocation) | **Post** /allocations/{uuid}/entitlements | Attach entitlement to allocation
[**CreateSatellite**](AllocationApi.md#CreateSatellite) | **Post** /allocations | Create Satellite
[**ExportAllocation**](AllocationApi.md#ExportAllocation) | **Get** /allocations/{uuid}/export | Trigger allocation manifest export
[**ExportJobAllocation**](AllocationApi.md#ExportJobAllocation) | **Get** /allocations/{uuid}/exportJob/{ExportJobID} | Check status of allocation manifest export
[**GetExportAllocation**](AllocationApi.md#GetExportAllocation) | **Get** /allocations/{uuid}/export/{ExportID} | Download allocation manifest
[**ListAllocationPools**](AllocationApi.md#ListAllocationPools) | **Get** /allocations/{uuid}/pools | List all pools for an allocation
[**ListAllocations**](AllocationApi.md#ListAllocations) | **Get** /allocations | List all allocations for a user
[**ListVersionsAllocation**](AllocationApi.md#ListVersionsAllocation) | **Get** /allocations/versions | List Satellite versions
[**RemoveAllocation**](AllocationApi.md#RemoveAllocation) | **Delete** /allocations/{uuid} | Remove allocation profile
[**RemoveAllocationEntitlement**](AllocationApi.md#RemoveAllocationEntitlement) | **Delete** /allocations/{uuid}/entitlements/{EntitlementID} | Remove entitlement from the allocation
[**RemoveAllocationEntitlementDeprecated**](AllocationApi.md#RemoveAllocationEntitlementDeprecated) | **Delete** /allocations/{uuid}/{EntitlementID} | Remove entitlement from the allocation
[**ShowAllocation**](AllocationApi.md#ShowAllocation) | **Get** /allocations/{uuid} | Get an allocation by UUID
[**UpdateAllocation**](AllocationApi.md#UpdateAllocation) | **Put** /allocations/{uuid} | Update an allocation
[**UpdateEntitlementAllocation**](AllocationApi.md#UpdateEntitlementAllocation) | **Put** /allocations/{uuid}/entitlements/{EntitlementID} | Update attached entitlement to allocation



## AttachEntitlementAllocation

> ShowAllocation200Response AttachEntitlementAllocation(ctx, uuid).Pool(pool).Quantity(quantity).Execute()

Attach entitlement to allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    pool := "pool_example" // string | 
    uuid := "uuid_example" // string | 
    quantity := int32(56) // int32 | quantity you want to attach (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.AttachEntitlementAllocation(context.Background(), uuid).Pool(pool).Quantity(quantity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.AttachEntitlementAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachEntitlementAllocation`: ShowAllocation200Response
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.AttachEntitlementAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachEntitlementAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pool** | **string** |  | 

 **quantity** | **int32** | quantity you want to attach | 

### Return type

[**ShowAllocation200Response**](ShowAllocation200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSatellite

> CreateSatellite200Response CreateSatellite(ctx).Name(name).Version(version).Execute()

Create Satellite



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    name := "name_example" // string | must be less than 100 characters and use only numbers, letters, underscores, hyphens, and periods
    version := "version_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.CreateSatellite(context.Background()).Name(name).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.CreateSatellite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSatellite`: CreateSatellite200Response
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

[**CreateSatellite200Response**](CreateSatellite200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAllocation

> ExportAllocation200Response ExportAllocation(ctx, uuid).Execute()

Trigger allocation manifest export



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.ExportAllocation(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ExportAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAllocation`: ExportAllocation200Response
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ExportAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportAllocation200Response**](ExportAllocation200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportJobAllocation

> ExportJobAllocation200Response ExportJobAllocation(ctx, uuid, exportJobID).Execute()

Check status of allocation manifest export



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 
    exportJobID := "exportJobID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.ExportJobAllocation(context.Background(), uuid, exportJobID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ExportJobAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportJobAllocation`: ExportJobAllocation200Response
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ExportJobAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 
**exportJobID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportJobAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExportJobAllocation200Response**](ExportJobAllocation200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportAllocation

> []int32 GetExportAllocation(ctx, uuid, exportID).Execute()

Download allocation manifest



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 
    exportID := "exportID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.GetExportAllocation(context.Background(), uuid, exportID).Execute()
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
**uuid** | **string** |  | 
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

> PoolsListMock ListAllocationPools(ctx, uuid).Limit(limit).Offset(offset).Future(future).Execute()

List all pools for an allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)
    future := true // bool | include future dated pools for satellite 6.3 or higher (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.ListAllocationPools(context.Background(), uuid).Limit(limit).Offset(offset).Future(future).Execute()
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
**uuid** | **string** |  | 

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

> ListAllocations200Response ListAllocations(ctx).Limit(limit).Offset(offset).Type_(type_).Execute()

List all allocations for a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.ListAllocations(context.Background()).Limit(limit).Offset(offset).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ListAllocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllocations`: ListAllocations200Response
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

[**ListAllocations200Response**](ListAllocations200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersionsAllocation

> ListVersionsAllocation200Response ListVersionsAllocation(ctx).Execute()

List Satellite versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.ListVersionsAllocation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ListVersionsAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVersionsAllocation`: ListVersionsAllocation200Response
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ListVersionsAllocation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionsAllocationRequest struct via the builder pattern


### Return type

[**ListVersionsAllocation200Response**](ListVersionsAllocation200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAllocation

> RemoveAllocation(ctx, uuid).Force(force).Execute()

Remove allocation profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 
    force := true // bool | Deleting a subscription allocation can have significant impacts on your hosts and activation keys. We require a force parameter to make sure the delete operation is intentional.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AllocationApi.RemoveAllocation(context.Background(), uuid).Force(force).Execute()
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
**uuid** | **string** |  | 

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

> RemoveAllocationEntitlement(ctx, uuid, entitlementID).Execute()

Remove entitlement from the allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 
    entitlementID := "entitlementID_example" // string | Remove an entitlement from an allocation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AllocationApi.RemoveAllocationEntitlement(context.Background(), uuid, entitlementID).Execute()
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
**uuid** | **string** |  | 
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


## RemoveAllocationEntitlementDeprecated

> RemoveAllocationEntitlementDeprecated(ctx, uuid, entitlementID).Execute()

Remove entitlement from the allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 
    entitlementID := "entitlementID_example" // string | Remove an entitlement from an allocation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AllocationApi.RemoveAllocationEntitlementDeprecated(context.Background(), uuid, entitlementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.RemoveAllocationEntitlementDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 
**entitlementID** | **string** | Remove an entitlement from an allocation | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAllocationEntitlementDeprecatedRequest struct via the builder pattern


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

> ShowAllocation200Response ShowAllocation(ctx, uuid).Include(include).Execute()

Get an allocation by UUID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 
    include := "include_example" // string | Show more details about a allocation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.ShowAllocation(context.Background(), uuid).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.ShowAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowAllocation`: ShowAllocation200Response
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.ShowAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Show more details about a allocation | 

### Return type

[**ShowAllocation200Response**](ShowAllocation200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAllocation

> UpdateAllocation(ctx, uuid).Allocation(allocation).Execute()

Update an allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 
    allocation := *openapiclient.NewUpdateAllocationRequest("SimpleContentAccess_example") // UpdateAllocationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AllocationApi.UpdateAllocation(context.Background(), uuid).Allocation(allocation).Execute()
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
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allocation** | [**UpdateAllocationRequest**](UpdateAllocationRequest.md) |  | 

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

> ShowAllocation200Response UpdateEntitlementAllocation(ctx, uuid, entitlementID).Quantity(quantity).Execute()

Update attached entitlement to allocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
    uuid := "uuid_example" // string | 
    entitlementID := "entitlementID_example" // string | 
    quantity := int32(56) // int32 | maxItem: quantity must be less than or equal to the maximum number of allowed entitlements in the entitlement pool minItem: 1 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllocationApi.UpdateEntitlementAllocation(context.Background(), uuid, entitlementID).Quantity(quantity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllocationApi.UpdateEntitlementAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntitlementAllocation`: ShowAllocation200Response
    fmt.Fprintf(os.Stdout, "Response from `AllocationApi.UpdateEntitlementAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 
**entitlementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitlementAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **quantity** | **int32** | maxItem: quantity must be less than or equal to the maximum number of allowed entitlements in the entitlement pool minItem: 1 | 

### Return type

[**ShowAllocation200Response**](ShowAllocation200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

