# \SubscriptionApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSubContentSets**](SubscriptionApi.md#ListSubContentSets) | **Get** /subscriptions/{SubscriptionNumber}/contentSets | List all content sets for a subscription
[**ListSubSystems**](SubscriptionApi.md#ListSubSystems) | **Get** /subscriptions/{SubscriptionNumber}/systems | List all systems consuming a subscription
[**ListSubscriptions**](SubscriptionApi.md#ListSubscriptions) | **Get** /subscriptions | List all subscriptions for a user



## ListSubContentSets

> InlineResponse20010 ListSubContentSets(ctx, subscriptionNumber, optional)

List all content sets for a subscription

The default and max results in a response are 1000.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionNumber** | **string**|  | 
 **optional** | ***ListSubContentSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubContentSetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubSystems

> InlineResponse20011 ListSubSystems(ctx, subscriptionNumber, optional)

List all systems consuming a subscription

The default and max results in a response are 100.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionNumber** | **string**|  | 
 **optional** | ***ListSubSystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubSystemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> InlineResponse2009 ListSubscriptions(ctx, optional)

List all subscriptions for a user

The default and max results in a response are 50.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

