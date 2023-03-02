# \SubscriptionApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSubContentSets**](SubscriptionApi.md#ListSubContentSets) | **Get** /subscriptions/{SubscriptionNumber}/contentSets | List all content sets for a subscription
[**ListSubSystems**](SubscriptionApi.md#ListSubSystems) | **Get** /subscriptions/{SubscriptionNumber}/systems | List all systems consuming a subscription
[**ListSubscriptions**](SubscriptionApi.md#ListSubscriptions) | **Get** /subscriptions | List all subscriptions for a user



## ListSubContentSets

> ListSubContentSets200Response ListSubContentSets(ctx, subscriptionNumber).Limit(limit).Offset(offset).Execute()

List all content sets for a subscription



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
    subscriptionNumber := "subscriptionNumber_example" // string | 
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ListSubContentSets(context.Background(), subscriptionNumber).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ListSubContentSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubContentSets`: ListSubContentSets200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ListSubContentSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubContentSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 

### Return type

[**ListSubContentSets200Response**](ListSubContentSets200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubSystems

> ListSubSystems200Response ListSubSystems(ctx, subscriptionNumber).Limit(limit).Offset(offset).Execute()

List all systems consuming a subscription



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
    subscriptionNumber := "subscriptionNumber_example" // string | 
    limit := int32(56) // int32 | max number of results you want (optional)
    offset := int32(56) // int32 | index from which you want next items (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ListSubSystems(context.Background(), subscriptionNumber).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ListSubSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubSystems`: ListSubSystems200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ListSubSystems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 

### Return type

[**ListSubSystems200Response**](ListSubSystems200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> ListSubscriptions200Response ListSubscriptions(ctx).Limit(limit).Offset(offset).Execute()

List all subscriptions for a user



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
    resp, r, err := apiClient.SubscriptionApi.ListSubscriptions(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ListSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptions`: ListSubscriptions200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ListSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | max number of results you want | 
 **offset** | **int32** | index from which you want next items | 

### Return type

[**ListSubscriptions200Response**](ListSubscriptions200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

