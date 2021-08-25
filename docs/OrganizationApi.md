# \OrganizationApi

All URIs are relative to *https://api.access.redhat.com/management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOrgSCACapability**](OrganizationApi.md#CheckOrgSCACapability) | **Get** /organization | Get details of the user&#39;s organization



## CheckOrgSCACapability

> InlineResponse2009 CheckOrgSCACapability(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.CheckOrgSCACapability(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.CheckOrgSCACapability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckOrgSCACapability`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.CheckOrgSCACapability`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckOrgSCACapabilityRequest struct via the builder pattern


### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

