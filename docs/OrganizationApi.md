# \OrganizationApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOrgSCACapability**](OrganizationApi.md#CheckOrgSCACapability) | **Get** /organization | Get details of the user&#39;s organization



## CheckOrgSCACapability

> InlineResponse20010 CheckOrgSCACapability(ctx).Include(include).Execute()

Get details of the user's organization



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
    include := "include_example" // string | Request for system purpose attributes in response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApi.CheckOrgSCACapability(context.Background()).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.CheckOrgSCACapability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckOrgSCACapability`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.CheckOrgSCACapability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckOrgSCACapabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Request for system purpose attributes in response | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

