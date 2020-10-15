# \PackagesApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadPackage**](PackagesApi.md#DownloadPackage) | **Get** /packages/{checksum}/download | Download a package by its SHA256 checksum
[**ListPackagesByContentSetArch**](PackagesApi.md#ListPackagesByContentSetArch) | **Get** /packages/cset/{ContentSet}/arch/{Arch} | Get all the packages for the specified content set and arch.
[**ShowPackage**](PackagesApi.md#ShowPackage) | **Get** /packages/{Checksum} | Get the details of a package



## DownloadPackage

> DownloadPackage(ctx, checksum)

Download a package by its SHA256 checksum

Find a package by its SHA256 checksum and generate a download link with a short-lived expiration. It is expected for users to obtain a new download link every time a package is downloaded and to not store the link for more than several minutes. If the user has \"Download Software and Updates\" permissions and a valid subscription for the package, they will receive an HTTP 307 redirect to the location on the Red Hat CDN. Clients can either follow the HTTP redirect or find the download URL in the response body. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checksum** | **string**|  | 

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

> InlineResponse2007 ListPackagesByContentSetArch(ctx, contentSet, arch, optional)

Get all the packages for the specified content set and arch.

The default and max results in a response are 50 and 100 respectively.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentSet** | **string**|  | 
**arch** | **string**|  | 
 **optional** | ***ListPackagesByContentSetArchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPackagesByContentSetArchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 
 **filter** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPackage

> InlineResponse2008 ShowPackage(ctx, checksum)

Get the details of a package

This will get the details of a package specified by its checksum.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checksum** | **string**|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

