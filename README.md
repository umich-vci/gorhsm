# Go API client for gorhsm

API for Red Hat Subscription Management

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.177.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://access.redhat.com/support/cases/](https://access.redhat.com/support/cases/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./gorhsm"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.access.redhat.com/management/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AllocationApi* | [**AttachEntitlementAllocation**](docs/AllocationApi.md#attachentitlementallocation) | **Post** /allocations/{AllocationUUID}/entitlements | Attach entitlement to allocation
*AllocationApi* | [**CreateSatellite**](docs/AllocationApi.md#createsatellite) | **Post** /allocations | Create Satellite
*AllocationApi* | [**ExportAllocation**](docs/AllocationApi.md#exportallocation) | **Get** /allocations/{AllocationUUID}/export | Trigger allocation manifest export
*AllocationApi* | [**ExportJobAllocation**](docs/AllocationApi.md#exportjoballocation) | **Get** /allocations/{AllocationUUID}/exportJob/{ExportJobID} | Check status of allocation manifest export
*AllocationApi* | [**GetExportAllocation**](docs/AllocationApi.md#getexportallocation) | **Get** /allocations/{AllocationUUID}/export/{ExportID} | Download allocation manifest
*AllocationApi* | [**ListAllocationPools**](docs/AllocationApi.md#listallocationpools) | **Get** /allocations/{AllocationUUID}/pools | List all pools for an allocation
*AllocationApi* | [**ListAllocations**](docs/AllocationApi.md#listallocations) | **Get** /allocations | List all allocations for a user
*AllocationApi* | [**ListVersionsAllocation**](docs/AllocationApi.md#listversionsallocation) | **Get** /allocations/versions | List Satellite versions
*AllocationApi* | [**RemoveAllocation**](docs/AllocationApi.md#removeallocation) | **Delete** /allocations/{AllocationUUID} | Remove allocation profile
*AllocationApi* | [**RemoveAllocationEntitlement**](docs/AllocationApi.md#removeallocationentitlement) | **Delete** /allocations/{AllocationUUID}/{EntitlementID} | Remove entitlement from the allocation
*AllocationApi* | [**ShowAllocation**](docs/AllocationApi.md#showallocation) | **Get** /allocations/{AllocationUUID} | Get an allocation by UUID
*AllocationApi* | [**UpdateEntitlementAllocation**](docs/AllocationApi.md#updateentitlementallocation) | **Put** /allocations/{AllocationUUID}/entitlements/{EntitlementUUID} | Update attached entitlement to allocation
*CloudaccessApi* | [**AddProviderAccounts**](docs/CloudaccessApi.md#addprovideraccounts) | **Post** /cloud_access_providers/{ProviderShortName}/accounts | Add accounts for a provider
*CloudaccessApi* | [**EnableGoldImages**](docs/CloudaccessApi.md#enablegoldimages) | **Post** /cloud_access_providers/{ProviderShortName}/goldimage | Enable Gold image access
*CloudaccessApi* | [**ListEnabledCloudAccessProviders**](docs/CloudaccessApi.md#listenabledcloudaccessproviders) | **Get** /cloud_access_providers/enabled | List all enabled cloud access providers for a user
*CloudaccessApi* | [**RemoveProviderAccount**](docs/CloudaccessApi.md#removeprovideraccount) | **Delete** /cloud_access_providers/{ProviderShortName}/accounts | Remove a provider account
*CloudaccessApi* | [**UpdateProviderAccount**](docs/CloudaccessApi.md#updateprovideraccount) | **Put** /cloud_access_providers/{ProviderShortName}/accounts/{AccountID} | Update provider account
*CloudaccessApi* | [**UpdateProviderAccountDeprecated**](docs/CloudaccessApi.md#updateprovideraccountdeprecated) | **Put** /cloud_access_providers/{ProviderShortName}/account | Update provider account
*CloudaccessApi* | [**VerifyProviderAccount**](docs/CloudaccessApi.md#verifyprovideraccount) | **Put** /cloud_access_providers/{ProviderShortName}/accounts/{AccountID}/verification | Verify a provider account
*ErrataApi* | [**ListErrata**](docs/ErrataApi.md#listerrata) | **Get** /errata | List all errata for a user&#39;s systems
*ErrataApi* | [**ListErrataByContentSetArch**](docs/ErrataApi.md#listerratabycontentsetarch) | **Get** /errata/cset/{ContentSet}/arch/{Arch} | Get all the errata for the specified content set and arch
*ErrataApi* | [**ListErratumPackages**](docs/ErrataApi.md#listerratumpackages) | **Get** /errata/{AdvisoryID}/packages | List all packages for an advisory
*ErrataApi* | [**ListErratumSystems**](docs/ErrataApi.md#listerratumsystems) | **Get** /errata/{AdvisoryID}/systems | List all systems for an advisory
*ErrataApi* | [**ShowErratum**](docs/ErrataApi.md#showerratum) | **Get** /errata/{AdvisoryID} | Get the details of an advisory
*ImagesApi* | [**DownloadImage**](docs/ImagesApi.md#downloadimage) | **Get** /images/{checksum}/download | Download an image by its SHA256 checksum
*ImagesApi* | [**ListImagesByContentSet**](docs/ImagesApi.md#listimagesbycontentset) | **Get** /images/cset/{ContentSet} | List available images in a content set
*PackagesApi* | [**DownloadPackage**](docs/PackagesApi.md#downloadpackage) | **Get** /packages/{checksum}/download | Download a package by its SHA256 checksum
*PackagesApi* | [**ListPackagesByContentSetArch**](docs/PackagesApi.md#listpackagesbycontentsetarch) | **Get** /packages/cset/{ContentSet}/arch/{Arch} | Get all the packages for the specified content set and arch.
*PackagesApi* | [**ShowPackage**](docs/PackagesApi.md#showpackage) | **Get** /packages/{Checksum} | Get the details of a package
*SubscriptionApi* | [**ListSubContentSets**](docs/SubscriptionApi.md#listsubcontentsets) | **Get** /subscriptions/{SubscriptionNumber}/contentSets | List all content sets for a subscription
*SubscriptionApi* | [**ListSubSystems**](docs/SubscriptionApi.md#listsubsystems) | **Get** /subscriptions/{SubscriptionNumber}/systems | List all systems consuming a subscription
*SubscriptionApi* | [**ListSubscriptions**](docs/SubscriptionApi.md#listsubscriptions) | **Get** /subscriptions | List all subscriptions for a user
*SystemApi* | [**AttachEntitlement**](docs/SystemApi.md#attachentitlement) | **Post** /systems/{SystemUUID}/entitlements | Attach entitlement to system
*SystemApi* | [**ListSystemErrata**](docs/SystemApi.md#listsystemerrata) | **Get** /systems/{SystemUUID}/errata | List all applicable errata for a system
*SystemApi* | [**ListSystemPackages**](docs/SystemApi.md#listsystempackages) | **Get** /systems/{SystemUUID}/packages | List all packages for a system
*SystemApi* | [**ListSystemPools**](docs/SystemApi.md#listsystempools) | **Get** /systems/{SystemUUID}/pools | List all pools for a system
*SystemApi* | [**ListSystems**](docs/SystemApi.md#listsystems) | **Get** /systems | List all systems for a user
*SystemApi* | [**RemoveSystem**](docs/SystemApi.md#removesystem) | **Delete** /systems/{SystemUUID} | Remove system profile
*SystemApi* | [**RemoveSystemEntitlement**](docs/SystemApi.md#removesystementitlement) | **Delete** /systems/{SystemUUID}/{EntitlementID} | Remove entitlement from the system
*SystemApi* | [**ShowSystem**](docs/SystemApi.md#showsystem) | **Get** /systems/{SystemUUID} | Get a system specified by UUID.


## Documentation For Models

 - [APIPageParam](docs/APIPageParam.md)
 - [AddProviderAccount](docs/AddProviderAccount.md)
 - [Allocation](docs/Allocation.md)
 - [AllocationDetails](docs/AllocationDetails.md)
 - [AllocationSummary](docs/AllocationSummary.md)
 - [AllocationVersion](docs/AllocationVersion.md)
 - [AttachEntitlement](docs/AttachEntitlement.md)
 - [ContentSet](docs/ContentSet.md)
 - [ContentSetArchMock](docs/ContentSetArchMock.md)
 - [DetailResponse](docs/DetailResponse.md)
 - [DownloadLink](docs/DownloadLink.md)
 - [EnabledCloudAccessProvider](docs/EnabledCloudAccessProvider.md)
 - [EnabledProduct](docs/EnabledProduct.md)
 - [EnabledProviderAccount](docs/EnabledProviderAccount.md)
 - [EntitlementsAttachedResponse](docs/EntitlementsAttachedResponse.md)
 - [EntitlementsAttachedResponseValue](docs/EntitlementsAttachedResponseValue.md)
 - [ErrataApplicabilityCounts](docs/ErrataApplicabilityCounts.md)
 - [ErrataCount](docs/ErrataCount.md)
 - [ErratumDetails](docs/ErratumDetails.md)
 - [ErratumForSystem](docs/ErratumForSystem.md)
 - [ErratumInContentSet](docs/ErratumInContentSet.md)
 - [ErrorDetails](docs/ErrorDetails.md)
 - [ExportJobResponse](docs/ExportJobResponse.md)
 - [ExportResponse](docs/ExportResponse.md)
 - [Facts](docs/Facts.md)
 - [GoldImageStatus](docs/GoldImageStatus.md)
 - [ImageInContentSet](docs/ImageInContentSet.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse202](docs/InlineResponse202.md)
 - [InlineResponse307](docs/InlineResponse307.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [InstalledProducts](docs/InstalledProducts.md)
 - [MyErrataListMock](docs/MyErrataListMock.md)
 - [MyErratum](docs/MyErratum.md)
 - [OngoingExportJobResponse](docs/OngoingExportJobResponse.md)
 - [PackageDetail](docs/PackageDetail.md)
 - [PackageForSystem](docs/PackageForSystem.md)
 - [PackageForSystemAdvisories](docs/PackageForSystemAdvisories.md)
 - [PkgContentSetArch](docs/PkgContentSetArch.md)
 - [PkgDetails](docs/PkgDetails.md)
 - [PkgListMock](docs/PkgListMock.md)
 - [Pool](docs/Pool.md)
 - [PoolDetail](docs/PoolDetail.md)
 - [PoolsListMock](docs/PoolsListMock.md)
 - [Reference](docs/Reference.md)
 - [ShowSystem](docs/ShowSystem.md)
 - [System](docs/System.md)
 - [SystemList](docs/SystemList.md)
 - [SystemListMock](docs/SystemListMock.md)


## Documentation For Authorization



### Bearer

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



