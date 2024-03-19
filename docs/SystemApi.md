# \SystemAPI

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachEntitlement**](SystemAPI.md#AttachEntitlement) | **Post** /systems/{SystemUUID}/entitlements | Attach entitlement to system
[**ListSystemErrata**](SystemAPI.md#ListSystemErrata) | **Get** /systems/{SystemUUID}/errata | List all applicable errata for a system
[**ListSystemPackages**](SystemAPI.md#ListSystemPackages) | **Get** /systems/{SystemUUID}/packages | List all packages for a system
[**ListSystemPools**](SystemAPI.md#ListSystemPools) | **Get** /systems/{SystemUUID}/pools | List all pools for a system
[**ListSystems**](SystemAPI.md#ListSystems) | **Get** /systems | List all systems for a user
[**RemoveSystem**](SystemAPI.md#RemoveSystem) | **Delete** /systems/{SystemUUID} | Remove system profile
[**RemoveSystemEntitlement**](SystemAPI.md#RemoveSystemEntitlement) | **Delete** /systems/{SystemUUID}/{EntitlementID} | Remove entitlement from the system
[**ShowSystem**](SystemAPI.md#ShowSystem) | **Get** /systems/{SystemUUID} | Get a system specified by UUID.



## AttachEntitlement

> AttachEntitlement200Response AttachEntitlement(ctx, systemUUID).Pool(pool).Quantity(quantity).Execute()

Attach entitlement to system



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
	systemUUID := "systemUUID_example" // string | 
	pool := "pool_example" // string | 
	quantity := int32(56) // int32 | quantity you want to attach (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.AttachEntitlement(context.Background(), systemUUID).Pool(pool).Quantity(quantity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.AttachEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachEntitlement`: AttachEntitlement200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.AttachEntitlement`: %v\n", resp)
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

[**AttachEntitlement200Response**](AttachEntitlement200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemErrata

> ListSystemErrata200Response ListSystemErrata(ctx, systemUUID).Limit(limit).Offset(offset).Execute()

List all applicable errata for a system



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
	systemUUID := "systemUUID_example" // string | 
	limit := int32(56) // int32 | max number of results you want (optional)
	offset := int32(56) // int32 | index from which you want next items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.ListSystemErrata(context.Background(), systemUUID).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListSystemErrata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSystemErrata`: ListSystemErrata200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListSystemErrata`: %v\n", resp)
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

[**ListSystemErrata200Response**](ListSystemErrata200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemPackages

> ListSystemPackages200Response ListSystemPackages(ctx, systemUUID).Limit(limit).Offset(offset).ErrataDetail(errataDetail).Upgradeable(upgradeable).Filter(filter).Execute()

List all packages for a system



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
	systemUUID := "systemUUID_example" // string | 
	limit := int32(56) // int32 | max number of results you want (optional)
	offset := int32(56) // int32 | index from which you want next items (optional)
	errataDetail := true // bool | Show errata details for packages (optional)
	upgradeable := true // bool | Show upgradable packages only. Also accepts 'upgradable' as valid query. (optional)
	filter := "filter_example" // string | Filter packages (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.ListSystemPackages(context.Background(), systemUUID).Limit(limit).Offset(offset).ErrataDetail(errataDetail).Upgradeable(upgradeable).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListSystemPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSystemPackages`: ListSystemPackages200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListSystemPackages`: %v\n", resp)
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

[**ListSystemPackages200Response**](ListSystemPackages200Response.md)

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
	openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
	systemUUID := "systemUUID_example" // string | 
	limit := int32(56) // int32 | max number of results you want (optional)
	offset := int32(56) // int32 | index from which you want next items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.ListSystemPools(context.Background(), systemUUID).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListSystemPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSystemPools`: PoolsListMock
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListSystemPools`: %v\n", resp)
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

> ListSystems200Response ListSystems(ctx).Limit(limit).Offset(offset).Filter(filter).Username(username).Execute()

List all systems for a user



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
	filter := "filter_example" // string | Filter Systems by System Name (optional)
	username := "username_example" // string | Filter Systems by a valid User Name, where User Name is the system owner and wildcard characters are not allowed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.ListSystems(context.Background()).Limit(limit).Offset(offset).Filter(filter).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListSystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSystems`: ListSystems200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListSystems`: %v\n", resp)
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

[**ListSystems200Response**](ListSystems200Response.md)

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
	openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
	systemUUID := "systemUUID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemAPI.RemoveSystem(context.Background(), systemUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.RemoveSystem``: %v\n", err)
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
	openapiclient "github.com/umich-vci/gorhsm"
)

func main() {
	systemUUID := "systemUUID_example" // string | 
	entitlementID := "entitlementID_example" // string | Remove an entitlement from a system

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemAPI.RemoveSystemEntitlement(context.Background(), systemUUID, entitlementID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.RemoveSystemEntitlement``: %v\n", err)
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

> ShowSystem200Response ShowSystem(ctx, systemUUID).Include(include).Execute()

Get a system specified by UUID.



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
	systemUUID := "systemUUID_example" // string | 
	include := []string{"Include_example"} // []string | Show more details about a system (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.ShowSystem(context.Background(), systemUUID).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ShowSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowSystem`: ShowSystem200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ShowSystem`: %v\n", resp)
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

[**ShowSystem200Response**](ShowSystem200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

