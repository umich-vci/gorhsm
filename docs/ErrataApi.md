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

> MyErrataListMock ListErrata(ctx, optional)

List all errata for a user's systems

The default and max results in a response are 1000.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListErrataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListErrataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| max number of results you want | 
 **offset** | **optional.Int64**| index from which you want next items | 

### Return type

[**MyErrataListMock**](myErrataListMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListErrataByContentSetArch

> ContentSetArchMock ListErrataByContentSetArch(ctx, contentSet, arch, optional)

Get all the errata for the specified content set and arch

Limit is the number of results in a response. The default limit is 50 and max limit is 100.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentSet** | **string**|  | 
**arch** | **string**|  | 
 **optional** | ***ListErrataByContentSetArchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListErrataByContentSetArchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int64**| max number of results you want | 
 **offset** | **optional.Int64**| index from which you want next items | 

### Return type

[**ContentSetArchMock**](contentSetArchMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListErratumPackages

> PkgListMock ListErratumPackages(ctx, advisoryID, optional)

List all packages for an advisory

The default and max results in a response are 50.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**advisoryID** | **string**| unique identifier for a Red Hat advisory | 
 **optional** | ***ListErratumPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListErratumPackagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| max number of results you want | 
 **offset** | **optional.Int64**| index from which you want next items | 

### Return type

[**PkgListMock**](pkgListMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListErratumSystems

> SystemListMock ListErratumSystems(ctx, advisoryID, optional)

List all systems for an advisory

The default and max results in a response are 1000.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**advisoryID** | **string**| unique identifier for a Red Hat advisory | 
 **optional** | ***ListErratumSystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListErratumSystemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| max number of results you want | 
 **offset** | **optional.Int64**| index from which you want next items | 

### Return type

[**SystemListMock**](systemListMock.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowErratum

> InlineResponse2004 ShowErratum(ctx, advisoryID)

Get the details of an advisory

This will get the details of an advisory specified by its advisoryID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**advisoryID** | **string**| unique identifier for a Red Hat advisory | 

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

