# \ImagesApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadImage**](ImagesApi.md#DownloadImage) | **Get** /images/{checksum}/download | Download an image by its SHA256 checksum



## DownloadImage

> DownloadImage(ctx, checksum)

Download an image by its SHA256 checksum

Find an image by its SHA256 checksum and generate a download link with a short-lived expiration. It is expected for users to obtain a new download link every time an image is downloaded and to not store the link for more than several minutes. If the user has \"Download Software and Updates\" permissions and a valid subscription for the image, they will receive a HTTP 307 redirect to the location on the Red Hat CDN. Clients can either follow the HTTP redirect or find the download URL in the response body. 

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

