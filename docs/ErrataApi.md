# \ErrataAPI

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListErrata**](ErrataAPI.md#ListErrata) | **Get** /errata | List all errata for a user&#39;s systems
[**ListErrataByContentSetArch**](ErrataAPI.md#ListErrataByContentSetArch) | **Get** /errata/cset/{ContentSet}/arch/{Arch} | Get all the errata for the specified content set and arch
[**ListErratumImages**](ErrataAPI.md#ListErratumImages) | **Get** /errata/{AdvisoryID}/images | List all updated container images for an advisory
[**ListErratumPackages**](ErrataAPI.md#ListErratumPackages) | **Get** /errata/{AdvisoryID}/packages | List all packages for an advisory
[**ListErratumSystems**](ErrataAPI.md#ListErratumSystems) | **Get** /errata/{AdvisoryID}/systems | List all systems for an advisory
[**ShowErratum**](ErrataAPI.md#ShowErratum) | **Get** /errata/{AdvisoryID} | Get the details of an advisory



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
	openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
	limit := int32(56) // int32 | max number of results you want (optional)
	offset := int32(56) // int32 | index from which you want next items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ErrataAPI.ListErrata(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrataAPI.ListErrata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListErrata`: MyErrataListMock
	fmt.Fprintf(os.Stdout, "Response from `ErrataAPI.ListErrata`: %v\n", resp)
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
	openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
	contentSet := "contentSet_example" // string | 
	arch := "arch_example" // string | 
	limit := int32(56) // int32 | max number of results you want (optional)
	offset := int32(56) // int32 | index from which you want next items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ErrataAPI.ListErrataByContentSetArch(context.Background(), contentSet, arch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrataAPI.ListErrataByContentSetArch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListErrataByContentSetArch`: ContentSetArchMock
	fmt.Fprintf(os.Stdout, "Response from `ErrataAPI.ListErrataByContentSetArch`: %v\n", resp)
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


## ListErratumImages

> UpdatedImagesList ListErratumImages(ctx, advisoryID).Execute()

List all updated container images for an advisory



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
	advisoryID := "advisoryID_example" // string | unique identifier for a Red Hat advisory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ErrataAPI.ListErratumImages(context.Background(), advisoryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrataAPI.ListErratumImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListErratumImages`: UpdatedImagesList
	fmt.Fprintf(os.Stdout, "Response from `ErrataAPI.ListErratumImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**advisoryID** | **string** | unique identifier for a Red Hat advisory | 

### Other Parameters

Other parameters are passed through a pointer to a apiListErratumImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdatedImagesList**](UpdatedImagesList.md)

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
	openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
	advisoryID := "advisoryID_example" // string | unique identifier for a Red Hat advisory
	limit := int32(56) // int32 | max number of results you want (optional)
	offset := int32(56) // int32 | index from which you want next items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ErrataAPI.ListErratumPackages(context.Background(), advisoryID).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrataAPI.ListErratumPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListErratumPackages`: PkgListMock
	fmt.Fprintf(os.Stdout, "Response from `ErrataAPI.ListErratumPackages`: %v\n", resp)
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
	openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
	advisoryID := "advisoryID_example" // string | unique identifier for a Red Hat advisory
	limit := int32(56) // int32 | max number of results you want (optional)
	offset := int32(56) // int32 | index from which you want next items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ErrataAPI.ListErratumSystems(context.Background(), advisoryID).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrataAPI.ListErratumSystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListErratumSystems`: SystemListMock
	fmt.Fprintf(os.Stdout, "Response from `ErrataAPI.ListErratumSystems`: %v\n", resp)
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

> ShowErratum200Response ShowErratum(ctx, advisoryID).Execute()

Get the details of an advisory



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
	advisoryID := "advisoryID_example" // string | unique identifier for a Red Hat advisory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ErrataAPI.ShowErratum(context.Background(), advisoryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrataAPI.ShowErratum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowErratum`: ShowErratum200Response
	fmt.Fprintf(os.Stdout, "Response from `ErrataAPI.ShowErratum`: %v\n", resp)
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

[**ShowErratum200Response**](ShowErratum200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

