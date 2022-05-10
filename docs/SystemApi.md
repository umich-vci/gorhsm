# \SystemApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachEntitlement**](SystemApi.md#AttachEntitlement) | **Post** /systems/{SystemUUID}/entitlements | Attach entitlement to system
[**ListSystemErrata**](SystemApi.md#ListSystemErrata) | **Get** /systems/{SystemUUID}/errata | List all applicable errata for a system
[**ListSystemPackages**](SystemApi.md#ListSystemPackages) | **Get** /systems/{SystemUUID}/packages | List all packages for a system
[**ListSystemPools**](SystemApi.md#ListSystemPools) | **Get** /systems/{SystemUUID}/pools | List all pools for a system
[**ListSystems**](SystemApi.md#ListSystems) | **Get** /systems | List all systems for a user
[**RemoveSystem**](SystemApi.md#RemoveSystem) | **Delete** /systems/{SystemUUID} | Remove system profile
[**RemoveSystemEntitlement**](SystemApi.md#RemoveSystemEntitlement) | **Delete** /systems/{SystemUUID}/{EntitlementID} | Remove entitlement from the system
[**ShowSystem**](SystemApi.md#ShowSystem) | **Get** /systems/{SystemUUID} | Get a system specified by UUID.



## AttachEntitlement

> InlineResponse20018 AttachEntitlement(ctx, systemUUID).Pool(pool).Quantity(quantity).Execute()

Attach entitlement to system



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
    systemUUID := "systemUUID_example" // string | 
    pool := "pool_example" // string | 
    quantity := int32(56) // int32 | quantity you want to attach (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.AttachEntitlement(context.Background(), systemUUID).Pool(pool).Quantity(quantity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.AttachEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachEntitlement`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.AttachEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pool** | **string** |  | 
 **quantity** | **int32** | quantity you want to attach | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemErrata

> InlineResponse20019 ListSystemErrata(ctx, systemUUID).Limit(limit).Offset(offset).Execute()

List all applicable errata for a system



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
    systemUUID := "systemUUID_example" // string | 
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.ListSystemErrata(context.Background(), systemUUID).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.ListSystemErrata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystemErrata`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.ListSystemErrata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemErrataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemPackages

> InlineResponse20020 ListSystemPackages(ctx, systemUUID).Limit(limit).Offset(offset).ErrataDetail(errataDetail).Upgradeable(upgradeable).Filter(filter).Execute()

List all packages for a system



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
    systemUUID := "systemUUID_example" // string | 
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)
    errataDetail := true // bool | Show errata details for packages (optional)
    upgradeable := true // bool | Show upgradable packages only. Also accepts 'upgradable' as valid query. (optional)
    filter := "filter_example" // string | Filter packages (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.ListSystemPackages(context.Background(), systemUUID).Limit(limit).Offset(offset).ErrataDetail(errataDetail).Upgradeable(upgradeable).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.ListSystemPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystemPackages`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.ListSystemPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 
 **errataDetail** | **bool** | Show errata details for packages | 
 **upgradeable** | **bool** | Show upgradable packages only. Also accepts &#39;upgradable&#39; as valid query. | 
 **filter** | **string** | Filter packages | 

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemPools

> PoolsListMock ListSystemPools(ctx, systemUUID).Limit(limit).Offset(offset).Execute()

List all pools for a system



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
    systemUUID := "systemUUID_example" // string | 
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.ListSystemPools(context.Background(), systemUUID).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.ListSystemPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystemPools`: PoolsListMock
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.ListSystemPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 

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


## ListSystems

> InlineResponse20016 ListSystems(ctx).Limit(limit).Offset(offset).Filter(filter).Username(username).Execute()

List all systems for a user



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
    filter := "filter_example" // string | Filter Systems by System Name (optional)
    username := "username_example" // string | Filter Systems by a valid User Name, where User Name is the system owner and wildcard characters are not allowed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.ListSystems(context.Background()).Limit(limit).Offset(offset).Filter(filter).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.ListSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystems`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.ListSystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 
 **filter** | **string** | Filter Systems by System Name | 
 **username** | **string** | Filter Systems by a valid User Name, where User Name is the system owner and wildcard characters are not allowed | 

### Return type

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSystem

> RemoveSystem(ctx, systemUUID).Execute()

Remove system profile



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
    systemUUID := "systemUUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.RemoveSystem(context.Background(), systemUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.RemoveSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSystemRequest struct via the builder pattern


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


## RemoveSystemEntitlement

> RemoveSystemEntitlement(ctx, systemUUID, entitlementID).Execute()

Remove entitlement from the system



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
    systemUUID := "systemUUID_example" // string | 
    entitlementID := "entitlementID_example" // string | Remove an entitlement from a system

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.RemoveSystemEntitlement(context.Background(), systemUUID, entitlementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.RemoveSystemEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string** |  | 
**entitlementID** | **string** | Remove an entitlement from a system | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSystemEntitlementRequest struct via the builder pattern


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


## ShowSystem

> InlineResponse20017 ShowSystem(ctx, systemUUID).Include(include).Execute()

Get a system specified by UUID.



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
    systemUUID := "systemUUID_example" // string | 
    include := []string{"Include_example"} // []string | Show more details about a system (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.ShowSystem(context.Background(), systemUUID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.ShowSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowSystem`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.ShowSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Show more details about a system | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

