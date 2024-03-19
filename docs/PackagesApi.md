# \PackagesAPI

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadPackage**](PackagesAPI.md#DownloadPackage) | **Get** /packages/{checksum}/download | Download a package by its SHA256 checksum
[**ListPackagesByContentSetArch**](PackagesAPI.md#ListPackagesByContentSetArch) | **Get** /packages/cset/{ContentSet}/arch/{Arch} | Get all the packages for the specified content set and arch.
[**ShowPackage**](PackagesAPI.md#ShowPackage) | **Get** /packages/{Checksum} | Get the details of a package



## DownloadPackage

> DownloadPackage(ctx, checksum).Execute()

Download a package by its SHA256 checksum



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
	checksum := "checksum_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackagesAPI.DownloadPackage(context.Background(), checksum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.DownloadPackage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDownloadPackageRequest struct via the builder pattern


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


## ListPackagesByContentSetArch

> ListPackagesByContentSetArch200Response ListPackagesByContentSetArch(ctx, contentSet, arch).Limit(limit).Offset(offset).Filter(filter).Execute()

Get all the packages for the specified content set and arch.



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
	contentSet := "contentSet_example" // string | 
	arch := "arch_example" // string | 
	limit := int32(56) // int32 | max number of results you want (optional)
	offset := int32(56) // int32 | index from which you want next items (optional)
	filter := []string{"Filter_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.ListPackagesByContentSetArch(context.Background(), contentSet, arch).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.ListPackagesByContentSetArch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPackagesByContentSetArch`: ListPackagesByContentSetArch200Response
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.ListPackagesByContentSetArch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentSet** | **string** |  | 
**arch** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPackagesByContentSetArchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 
 **filter** | **[]string** |  | 

### Return type

[**ListPackagesByContentSetArch200Response**](ListPackagesByContentSetArch200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPackage

> ShowPackage200Response ShowPackage(ctx, checksum).Execute()

Get the details of a package



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
	checksum := "checksum_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.ShowPackage(context.Background(), checksum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.ShowPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPackage`: ShowPackage200Response
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.ShowPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checksum** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShowPackage200Response**](ShowPackage200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

