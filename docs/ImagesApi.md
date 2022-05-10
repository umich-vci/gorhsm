# \ImagesApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadImage**](ImagesApi.md#DownloadImage) | **Get** /images/{checksum}/download | Download an image by its SHA256 checksum
[**ListImageDownloadsByVersionArch**](ImagesApi.md#ListImageDownloadsByVersionArch) | **Get** /images/rhel/{Version}/{Arch} | List RHEL image downloads by version and architecture.
[**ListImagesByContentSet**](ImagesApi.md#ListImagesByContentSet) | **Get** /images/cset/{ContentSet} | List available images in a content set



## DownloadImage

> DownloadImage(ctx, checksum).Execute()

Download an image by its SHA256 checksum



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
    checksum := "checksum_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.DownloadImage(context.Background(), checksum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.DownloadImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checksum** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadImageRequest struct via the builder pattern


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


## ListImageDownloadsByVersionArch

> InlineResponse2009 ListImageDownloadsByVersionArch(ctx, version, arch).Execute()

List RHEL image downloads by version and architecture.



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
    version := "version_example" // string | 
    arch := "arch_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ListImageDownloadsByVersionArch(context.Background(), version, arch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ListImageDownloadsByVersionArch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImageDownloadsByVersionArch`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ListImageDownloadsByVersionArch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 
**arch** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageDownloadsByVersionArchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImagesByContentSet

> InlineResponse2008 ListImagesByContentSet(ctx, contentSet).Limit(limit).Offset(offset).Execute()

List available images in a content set



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
    limit := int32(56) // int32 | maximum number of list items in a page (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ListImagesByContentSet(context.Background(), contentSet).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ListImagesByContentSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImagesByContentSet`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ListImagesByContentSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentSet** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImagesByContentSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum number of list items in a page | 
 **offset** | **int32** | index from which you want next items | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

