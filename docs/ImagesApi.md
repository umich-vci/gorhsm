# \ImagesApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadImage**](ImagesApi.md#DownloadImage) | **Get** /images/{checksum}/download | Download an image by its SHA256 checksum
[**ListImagesByContentSet**](ImagesApi.md#ListImagesByContentSet) | **Get** /images/cset/{ContentSet} | List available images in a content set



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


## ListImagesByContentSet

> InlineResponse2006 ListImagesByContentSet(ctx, contentSet, optional)

List available images in a content set

List all the available images in a given content set. The content set parameter is expected to be  properly formatted (for example rhel-8-for-x86_64-baseos-isos). And the user requesting the endpoint is expected to have Download permission (UGC). If the user is entitled to an image, a \"downloadHref\" attribute is added in the image response object which links to image download API. The list is paginated by default to 25 results in a response and goes to maximum 100 results in a response. Use pagination by setting offset and limit url parameters (valid integer values). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentSet** | **string**|  | 
 **optional** | ***ListImagesByContentSetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListImagesByContentSetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| maximum number of list items in a page | 
 **offset** | **optional.Int32**| index from which you want next items | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

