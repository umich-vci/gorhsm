# \AllocationApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachEntitlementAllocation**](AllocationApi.md#AttachEntitlementAllocation) | **Post** /allocations/{AllocationUUID}/entitlements | Attach entitlement to allocation
[**CreateSatellite**](AllocationApi.md#CreateSatellite) | **Post** /allocations | Create Satellite by name
[**ExportAllocation**](AllocationApi.md#ExportAllocation) | **Get** /allocations/{AllocationUUID}/export | Trigger allocation manifest export
[**ExportJobAllocation**](AllocationApi.md#ExportJobAllocation) | **Get** /allocations/{AllocationUUID}/exportJob/{ExportJobID} | Check status of allocation manifest export
[**GetExportAllocation**](AllocationApi.md#GetExportAllocation) | **Get** /allocations/{AllocationUUID}/export/{ExportID} | Download allocation manifest
[**ListAllocationPools**](AllocationApi.md#ListAllocationPools) | **Get** /allocations/{AllocationUUID}/pools | List all pools for an allocation
[**ListAllocations**](AllocationApi.md#ListAllocations) | **Get** /allocations | List all allocations for a user
[**RemoveAllocation**](AllocationApi.md#RemoveAllocation) | **Delete** /allocations/{AllocationUUID} | Remove allocation profile
[**RemoveAllocationEntitlement**](AllocationApi.md#RemoveAllocationEntitlement) | **Delete** /allocations/{AllocationUUID}/{EntitlementID} | Remove entitlement from the allocation
[**ShowAllocation**](AllocationApi.md#ShowAllocation) | **Get** /allocations/{AllocationUUID} | Get an allocation by UUID
[**UpdateEntitlementAllocation**](AllocationApi.md#UpdateEntitlementAllocation) | **Put** /allocations/{AllocationUUID}/entitlements/{EntitlementUUID} | Update attached entitlement to allocation



## AttachEntitlementAllocation

> InlineResponse2001 AttachEntitlementAllocation(ctx, pool, allocationUUID, optional)

Attach entitlement to allocation

The default success response will be 200.  System, RHUI, Hypervisor are unsupported allocation types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pool** | **string**|  | 
**allocationUUID** | **string**|  | 
 **optional** | ***AttachEntitlementAllocationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AttachEntitlementAllocationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **quantity** | **optional.Int32**| quantity you want to attach | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSatellite

> InlineResponse2001 CreateSatellite(ctx, name)

Create Satellite by name

 Version 6.5 (most recent version of Satellite)  The default success response will be 200.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| must be less than 100 characters and use only numbers, letters, underscores, hyphens, and periods | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAllocation

> InlineResponse2002 ExportAllocation(ctx, allocationUUID)

Trigger allocation manifest export

Starts job to generate export for an allocation. To check the status of the export job visit the href in the response.  System, RHUI, Hypervisor are unsupported allocation types. SAM 1.2 or lower, and Satellite 5 versions are unsupported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string**|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportJobAllocation

> InlineResponse2003 ExportJobAllocation(ctx, allocationUUID, exportJobID)

Check status of allocation manifest export

Returns export download link in response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string**|  | 
**exportJobID** | **string**|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportAllocation

> []int32 GetExportAllocation(ctx, allocationUUID, exportID)

Download allocation manifest

Success response contains a zip file. The link is one-time download and expires after one try. Trigger export job to get another download link.  Content-Type: application/zip

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string**|  | 
**exportID** | **string**|  | 

### Return type

**[]int32**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllocationPools

> PoolsListMock ListAllocationPools(ctx, allocationUUID, optional)

List all pools for an allocation

System, RHUI, Hypervisor are unsupported allocation types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string**|  | 
 **optional** | ***ListAllocationPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllocationPoolsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 
 **future** | **optional.Bool**| include future dated pools for satellite 6.3 or higher | 

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


## ListAllocations

> InlineResponse200 ListAllocations(ctx, optional)

List all allocations for a user

The default and max number of results in a response are 100.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAllocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllocationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| max number of results you want | 
 **offset** | **optional.Int32**| index from which you want next items | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAllocation

> RemoveAllocation(ctx, allocationUUID, force)

Remove allocation profile

The default success response will be 204  System, RHUI, Hypervisor are unsupported allocation types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string**|  | 
**force** | **bool**| Deleting a subscription allocation can have significant impacts on your hosts and activation keys. We require a force parameter to make sure the delete operation is intentional. | 

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


## RemoveAllocationEntitlement

> RemoveAllocationEntitlement(ctx, allocationUUID, entitlementID)

Remove entitlement from the allocation

The default success response will be 204.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string**|  | 
**entitlementID** | **string**| Remove an entitlement from an allocation | 

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


## ShowAllocation

> InlineResponse2001 ShowAllocation(ctx, allocationUUID, optional)

Get an allocation by UUID

System, RHUI, Hypervisor are unsupported allocation types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string**|  | 
 **optional** | ***ShowAllocationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ShowAllocationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Show more details about a allocation | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitlementAllocation

> InlineResponse2001 UpdateEntitlementAllocation(ctx, allocationUUID, entitlementUUID, optional)

Update attached entitlement to allocation

The default success response will be 200.  System, RHUI, Hypervisor are unsupported allocation types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allocationUUID** | **string**|  | 
**entitlementUUID** | **string**|  | 
 **optional** | ***UpdateEntitlementAllocationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateEntitlementAllocationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **quantity** | **optional.Int32**| maxItem: quantity must be less than or equal to the maximum number of allowed entitlements in the entitlement pool minItem: 1 | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

