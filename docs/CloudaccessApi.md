# \CloudaccessApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProviderAccounts**](CloudaccessApi.md#AddProviderAccounts) | **Post** /cloud_access_providers/{ProviderShortName}/accounts | Add accounts for a provider
[**EnableGoldImages**](CloudaccessApi.md#EnableGoldImages) | **Post** /cloud_access_providers/{ProviderShortName}/goldimage | Enable Gold image access
[**ListEnabledCloudAccessProviders**](CloudaccessApi.md#ListEnabledCloudAccessProviders) | **Get** /cloud_access_providers/enabled | List all enabled cloud access providers for a user
[**RemoveProviderAccount**](CloudaccessApi.md#RemoveProviderAccount) | **Delete** /cloud_access_providers/{ProviderShortName}/accounts | Remove a provider account
[**UpdateProviderAccount**](CloudaccessApi.md#UpdateProviderAccount) | **Put** /cloud_access_providers/{ProviderShortName}/account | Update provider account



## AddProviderAccounts

> AddProviderAccounts(ctx, providerShortName, optional)

Add accounts for a provider

Add up to `100` new provider accounts, with optional nicknames, to a currently-enabled provider for Red Hat Cloud Access. You can find a list of currently-enabled provider accounts and provider short names from the `/v1/cloud_access_providers/enabled` endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string**|  | 
 **optional** | ***AddProviderAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddProviderAccountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**optional.Interface of []AddProviderAccount**](AddProviderAccount.md)|  | 

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

> EnableGoldImages(ctx, providerShortName, optional)

Enable Gold image access

Requests access to Red Hat Gold Images, where available, for currently-enabled products and provider accounts. Customers can request Red Hat Gold Images for account IDs and product image groups listed in the `/v1/cloud_access_providers/enabled` endpoint using the provider short name listed in the same response. After the request has been accepted for processing, gold image status for accounts can be checked in the `/v1/cloud_access_providers/enabled` endpoint response. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string**|  | 
 **optional** | ***EnableGoldImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnableGoldImagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **goldImages** | [**optional.Interface of InlineObject2**](InlineObject2.md)|  | 

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

> InlineResponse2004 ListEnabledCloudAccessProviders(ctx, )

List all enabled cloud access providers for a user

Returns the full list of all enabled Red Hat products and Cloud Access provider accounts/subscriptions associated with the user’s Red Hat account. - For Products that are expired `nextRenewal` field would be omitted and `totalQuantity` would be `0`. - Product objects can have `totalQuantity` field as `-1` indicating `Unlimited` quantity available. - The `nextRenewalDate` field of a Product has the format `YYYY-MM-DD`. - The `goldImageStatus` field is available for an account when the Provider is a certified Gold Image Provider and the account has been requested for gold image access. It's value could be `\"Requested\"`, `\"Granted\"` or `\"Failed\"`. - The `imageGroups` field is available for a product when the Provider is a certified Gold Image Provider, and the product has a gold image group available for it. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProviderAccount

> RemoveProviderAccount(ctx, providerShortName, optional)

Remove a provider account

Removes a currently-enabled provider account, including removing access to Gold Images, where applicable. You can find a list of currently-enabled provider accounts and provider short names from the `/v1/cloud_access_providers/enabled` endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string**|  | 
 **optional** | ***RemoveProviderAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveProviderAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**optional.Interface of InlineObject1**](InlineObject1.md)|  | 

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

> UpdateProviderAccount(ctx, providerShortName, optional)

Update provider account

Updates the account ID and/or nickname for a currently-enabled provider account. You can find a list of currently-enabled provider accounts and provider short names from the `/v1/cloud_access_providers/enabled` endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerShortName** | **string**|  | 
 **optional** | ***UpdateProviderAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateProviderAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

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

