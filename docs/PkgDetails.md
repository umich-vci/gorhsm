# PkgDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**BuildDate** | Pointer to **string** | Date represents the date format used for API returns | [optional] 
**BuildHost** | Pointer to **string** |  | [optional] 
**Checksum** | Pointer to **string** |  | [optional] 
**ContentSets** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Epoch** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewPkgDetails

`func NewPkgDetails() *PkgDetails`

NewPkgDetails instantiates a new PkgDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgDetailsWithDefaults

`func NewPkgDetailsWithDefaults() *PkgDetails`

NewPkgDetailsWithDefaults instantiates a new PkgDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *PkgDetails) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *PkgDetails) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *PkgDetails) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *PkgDetails) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetBuildDate

`func (o *PkgDetails) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *PkgDetails) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *PkgDetails) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.

### HasBuildDate

`func (o *PkgDetails) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.

### GetBuildHost

`func (o *PkgDetails) GetBuildHost() string`

GetBuildHost returns the BuildHost field if non-nil, zero value otherwise.

### GetBuildHostOk

`func (o *PkgDetails) GetBuildHostOk() (*string, bool)`

GetBuildHostOk returns a tuple with the BuildHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildHost

`func (o *PkgDetails) SetBuildHost(v string)`

SetBuildHost sets BuildHost field to given value.

### HasBuildHost

`func (o *PkgDetails) HasBuildHost() bool`

HasBuildHost returns a boolean if a field has been set.

### GetChecksum

`func (o *PkgDetails) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *PkgDetails) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *PkgDetails) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *PkgDetails) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetContentSets

`func (o *PkgDetails) GetContentSets() []string`

GetContentSets returns the ContentSets field if non-nil, zero value otherwise.

### GetContentSetsOk

`func (o *PkgDetails) GetContentSetsOk() (*[]string, bool)`

GetContentSetsOk returns a tuple with the ContentSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSets

`func (o *PkgDetails) SetContentSets(v []string)`

SetContentSets sets ContentSets field to given value.

### HasContentSets

`func (o *PkgDetails) HasContentSets() bool`

HasContentSets returns a boolean if a field has been set.

### GetDescription

`func (o *PkgDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PkgDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PkgDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PkgDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEpoch

`func (o *PkgDetails) GetEpoch() string`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *PkgDetails) GetEpochOk() (*string, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *PkgDetails) SetEpoch(v string)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *PkgDetails) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetGroup

`func (o *PkgDetails) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PkgDetails) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PkgDetails) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PkgDetails) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHref

`func (o *PkgDetails) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PkgDetails) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PkgDetails) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PkgDetails) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLicense

`func (o *PkgDetails) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *PkgDetails) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *PkgDetails) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *PkgDetails) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *PkgDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkgDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkgDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PkgDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelease

`func (o *PkgDetails) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PkgDetails) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PkgDetails) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PkgDetails) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetSize

`func (o *PkgDetails) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PkgDetails) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PkgDetails) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PkgDetails) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSummary

`func (o *PkgDetails) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PkgDetails) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PkgDetails) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PkgDetails) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetVersion

`func (o *PkgDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PkgDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PkgDetails) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PkgDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


