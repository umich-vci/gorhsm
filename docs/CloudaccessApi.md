# \CloudaccessApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProviderAccounts**](CloudaccessApi.md#AddProviderAccounts) | **Post** /cloud_access_providers/{ProviderShortName}/accounts | Add accounts for a provider
[**EnableGoldImages**](CloudaccessApi.md#EnableGoldImages) | **Post** /cloud_access_providers/{ProviderShortName}/goldimage | Enable Gold image access
[**ListEnabledCloudAccessProviders**](CloudaccessApi.md#ListEnabledCloudAccessProviders) | **Get** /cloud_access_providers/enabled | List all enabled cloud access providers for a user
[**RemoveProviderAccount**](CloudaccessApi.md#RemoveProviderAccount) | **Delete** /cloud_access_providers/{ProviderShortName}/accounts | Remove a provider account
[**UpdateProviderAccount**](CloudaccessApi.md#UpdateProviderAccount) | **Put** /cloud_access_providers/{ProviderShortName}/accounts/{AccountID} | Update provider account
[**UpdateProviderAccountDeprecated**](CloudaccessApi.md#UpdateProviderAccountDeprecated) | **Put** /cloud_access_providers/{ProviderShortName}/account | Update provider account
[**VerifyProviderAccount**](CloudaccessApi.md#VerifyProviderAccount) | **Put** /cloud_access_providers/{ProviderShortName}/accounts/{AccountID}/verification | Verify a provider account



## AddProviderAccounts

> AddProviderAccounts(ctx, providerShortName).Account(account).Execute()

Add accounts for a provider



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
    providerShortName := "providerShortName_example" // string | 
    account := []openapiclient.AddProviderAccount{*openapiclient.NewAddProviderAccount()} // []AddProviderAccount |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudaccessApi.AddProviderAccounts(context.Background(), providerShortName).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudaccessApi.AddProviderAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProviderAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**[]AddProviderAccount**](AddProviderAccount.md) |  | 

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


## EnableGoldImages

> EnableGoldImages(ctx, providerShortName).GoldImages(goldImages).Execute()

Enable Gold image access



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
    providerShortName := "providerShortName_example" // string | 
    goldImages := *openapiclient.NewInlineObject4([]string{"Accounts_example"}, []string{"Images_example"}) // InlineObject4 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudaccessApi.EnableGoldImages(context.Background(), providerShortName).GoldImages(goldImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudaccessApi.EnableGoldImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableGoldImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **goldImages** | [**InlineObject4**](InlineObject4.md) |  | 

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


## ListEnabledCloudAccessProviders

> InlineResponse2006 ListEnabledCloudAccessProviders(ctx).Execute()

List all enabled cloud access providers for a user



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
    resp, r, err := api_client.CloudaccessApi.ListEnabledCloudAccessProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudaccessApi.ListEnabledCloudAccessProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnabledCloudAccessProviders`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `CloudaccessApi.ListEnabledCloudAccessProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEnabledCloudAccessProvidersRequest struct via the builder pattern


### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProviderAccount

> RemoveProviderAccount(ctx, providerShortName).Account(account).Execute()

Remove a provider account



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
    providerShortName := "providerShortName_example" // string | 
    account := *openapiclient.NewInlineObject1("Id_example") // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudaccessApi.RemoveProviderAccount(context.Background(), providerShortName).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudaccessApi.RemoveProviderAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProviderAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**InlineObject1**](InlineObject1.md) |  | 

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


## UpdateProviderAccount

> UpdateProviderAccount(ctx, providerShortName, accountID).Account(account).Execute()

Update provider account



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
    providerShortName := "providerShortName_example" // string | 
    accountID := "accountID_example" // string | 
    account := *openapiclient.NewInlineObject2("Nickname_example") // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudaccessApi.UpdateProviderAccount(context.Background(), providerShortName, accountID).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudaccessApi.UpdateProviderAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string** |  | 
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **account** | [**InlineObject2**](InlineObject2.md) |  | 

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


## UpdateProviderAccountDeprecated

> UpdateProviderAccountDeprecated(ctx, providerShortName).Account(account).Execute()

Update provider account



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
    providerShortName := "providerShortName_example" // string | 
    account := *openapiclient.NewInlineObject("Id_example") // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudaccessApi.UpdateProviderAccountDeprecated(context.Background(), providerShortName).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudaccessApi.UpdateProviderAccountDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderAccountDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**InlineObject**](InlineObject.md) |  | 

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


## VerifyProviderAccount

> VerifyProviderAccount(ctx, providerShortName, accountID).Account(account).Execute()

Verify a provider account



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
    providerShortName := "providerShortName_example" // string | 
    accountID := "accountID_example" // string | 
    account := *openapiclient.NewInlineObject3("Identity_example", "Signature_example") // InlineObject3 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudaccessApi.VerifyProviderAccount(context.Background(), providerShortName, accountID).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudaccessApi.VerifyProviderAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string** |  | 
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyProviderAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **account** | [**InlineObject3**](InlineObject3.md) |  | 

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

