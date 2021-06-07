# \ErrataApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListErrata**](ErrataApi.md#ListErrata) | **Get** /errata | List all errata for a user&#39;s systems
[**ListErrataByContentSetArch**](ErrataApi.md#ListErrataByContentSetArch) | **Get** /errata/cset/{ContentSet}/arch/{Arch} | Get all the errata for the specified content set and arch
[**ListErratumPackages**](ErrataApi.md#ListErratumPackages) | **Get** /errata/{AdvisoryID}/packages | List all packages for an advisory
[**ListErratumSystems**](ErrataApi.md#ListErratumSystems) | **Get** /errata/{AdvisoryID}/systems | List all systems for an advisory
[**ShowErratum**](ErrataApi.md#ShowErratum) | **Get** /errata/{AdvisoryID} | Get the details of an advisory



## ListErrata

> MyErrataListMock ListErrata(ctx).Limit(limit).Offset(offset).Execute()

List all errata for a user's systems



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrataApi.ListErrata(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrataApi.ListErrata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListErrata`: MyErrataListMock
    fmt.Fprintf(os.Stdout, "Response from `ErrataApi.ListErrata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListErrataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 

### Return type

[**MyErrataListMock**](MyErrataListMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListErrataByContentSetArch

> ContentSetArchMock ListErrataByContentSetArch(ctx, contentSet, arch).Limit(limit).Offset(offset).Execute()

Get all the errata for the specified content set and arch



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
    contentSet := "contentSet_example" // string | 
    arch := "arch_example" // string | 
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrataApi.ListErrataByContentSetArch(context.Background(), contentSet, arch).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrataApi.ListErrataByContentSetArch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListErrataByContentSetArch`: ContentSetArchMock
    fmt.Fprintf(os.Stdout, "Response from `ErrataApi.ListErrataByContentSetArch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentSet** | **string** |  | 
**arch** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListErrataByContentSetArchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 

### Return type

[**ContentSetArchMock**](ContentSetArchMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListErratumPackages

> PkgListMock ListErratumPackages(ctx, advisoryID).Limit(limit).Offset(offset).Execute()

List all packages for an advisory



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
    advisoryID := "advisoryID_example" // string | unique identifier for a Red Hat advisory
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrataApi.ListErratumPackages(context.Background(), advisoryID).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrataApi.ListErratumPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListErratumPackages`: PkgListMock
    fmt.Fprintf(os.Stdout, "Response from `ErrataApi.ListErratumPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**advisoryID** | **string** | unique identifier for a Red Hat advisory | 

### Other Parameters

Other parameters are passed through a pointer to a apiListErratumPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 

### Return type

[**PkgListMock**](PkgListMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListErratumSystems

> SystemListMock ListErratumSystems(ctx, advisoryID).Limit(limit).Offset(offset).Execute()

List all systems for an advisory



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
    advisoryID := "advisoryID_example" // string | unique identifier for a Red Hat advisory
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrataApi.ListErratumSystems(context.Background(), advisoryID).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrataApi.ListErratumSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListErratumSystems`: SystemListMock
    fmt.Fprintf(os.Stdout, "Response from `ErrataApi.ListErratumSystems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**advisoryID** | **string** | unique identifier for a Red Hat advisory | 

### Other Parameters

Other parameters are passed through a pointer to a apiListErratumSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 

### Return type

[**SystemListMock**](SystemListMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowErratum

> InlineResponse2007 ShowErratum(ctx, advisoryID).Execute()

Get the details of an advisory



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
    advisoryID := "advisoryID_example" // string | unique identifier for a Red Hat advisory

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ErrataApi.ShowErratum(context.Background(), advisoryID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ErrataApi.ShowErratum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowErratum`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `ErrataApi.ShowErratum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**advisoryID** | **string** | unique identifier for a Red Hat advisory | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowErratumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

