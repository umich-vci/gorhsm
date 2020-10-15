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

> InlineResponse20013 AttachEntitlement(ctx, systemUUID, pool, optional)

Attach entitlement to system

The default success response will be 200.  Sam & Satellite systems are unsupported system types.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string**|  | 
**pool** | **string**|  | 
 **optional** | ***AttachEntitlementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AttachEntitlementOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **quantity** | **optional.Int32**| quantity you want to attach | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemErrata

> InlineResponse20014 ListSystemErrata(ctx, systemUUID, optional)

List all applicable errata for a system

The default and max number of results in a response are 100.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string**|  | 
 **optional** | ***ListSystemErrataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSystemErrataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemPackages

> InlineResponse20015 ListSystemPackages(ctx, systemUUID, optional)

List all packages for a system

The default and max number of results in a response are 1000.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string**|  | 
 **optional** | ***ListSystemPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSystemPackagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 
 **errataDetail** | **optional.Bool**| Show errata details for packages | 
 **upgradeable** | **optional.Bool**| Show upgradable packages only. Also accepts &#39;upgradable&#39; as valid query. | 
 **filter** | **optional.String**| Filter packages | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemPools

> PoolsListMock ListSystemPools(ctx, systemUUID, optional)

List all pools for a system

The default and max number of results in a response are 50.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string**|  | 
 **optional** | ***ListSystemPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSystemPoolsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 

### Return type

[**PoolsListMock**](poolsListMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystems

> InlineResponse20012 ListSystems(ctx, optional)

List all systems for a user

The default and max number of results in a response are 100.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSystemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 
 **filter** | **optional.String**| Filter Systems by System Name | 
 **username** | **optional.String**| Filter Systems by a valid User Name, where User Name is the system owner and wildcard characters are not allowed | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSystem

> RemoveSystem(ctx, systemUUID)

Remove system profile

The default success response will be 204

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string**|  | 

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

> RemoveSystemEntitlement(ctx, systemUUID, entitlementID)

Remove entitlement from the system

The default success response will be 204.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string**|  | 
**entitlementID** | **string**| Remove an entitlement from a system | 

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

> InlineResponse20013 ShowSystem(ctx, systemUUID, optional)

Get a system specified by UUID.

Sam & Satellite systems are unsupported system types.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemUUID** | **string**|  | 
 **optional** | ***ShowSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ShowSystemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | [**optional.Interface of []string**](string.md)| Show more details about a system | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

