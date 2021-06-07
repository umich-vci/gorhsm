# PackageDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**Checksum** | Pointer to **string** |  | [optional] 
**ContentSets** | Pointer to **[]string** |  | [optional] 
**DetailsUrl** | Pointer to **string** |  | [optional] 
**Epoch** | Pointer to **int32** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**RepoTags** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewPackageDetail

`func NewPackageDetail() *PackageDetail`

NewPackageDetail instantiates a new PackageDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDetailWithDefaults

`func NewPackageDetailWithDefaults() *PackageDetail`

NewPackageDetailWithDefaults instantiates a new PackageDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *PackageDetail) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *PackageDetail) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *PackageDetail) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *PackageDetail) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetChecksum

`func (o *PackageDetail) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *PackageDetail) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *PackageDetail) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *PackageDetail) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetContentSets

`func (o *PackageDetail) GetContentSets() []string`

GetContentSets returns the ContentSets field if non-nil, zero value otherwise.

### GetContentSetsOk

`func (o *PackageDetail) GetContentSetsOk() (*[]string, bool)`

GetContentSetsOk returns a tuple with the ContentSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSets

`func (o *PackageDetail) SetContentSets(v []string)`

SetContentSets sets ContentSets field to given value.

### HasContentSets

`func (o *PackageDetail) HasContentSets() bool`

HasContentSets returns a boolean if a field has been set.

### GetDetailsUrl

`func (o *PackageDetail) GetDetailsUrl() string`

GetDetailsUrl returns the DetailsUrl field if non-nil, zero value otherwise.

### GetDetailsUrlOk

`func (o *PackageDetail) GetDetailsUrlOk() (*string, bool)`

GetDetailsUrlOk returns a tuple with the DetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsUrl

`func (o *PackageDetail) SetDetailsUrl(v string)`

SetDetailsUrl sets DetailsUrl field to given value.

### HasDetailsUrl

`func (o *PackageDetail) HasDetailsUrl() bool`

HasDetailsUrl returns a boolean if a field has been set.

### GetEpoch

`func (o *PackageDetail) GetEpoch() int32`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *PackageDetail) GetEpochOk() (*int32, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *PackageDetail) SetEpoch(v int32)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *PackageDetail) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetFilename

`func (o *PackageDetail) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PackageDetail) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PackageDetail) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PackageDetail) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetName

`func (o *PackageDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelease

`func (o *PackageDetail) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PackageDetail) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PackageDetail) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PackageDetail) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRepoTags

`func (o *PackageDetail) GetRepoTags() []string`

GetRepoTags returns the RepoTags field if non-nil, zero value otherwise.

### GetRepoTagsOk

`func (o *PackageDetail) GetRepoTagsOk() (*[]string, bool)`

GetRepoTagsOk returns a tuple with the RepoTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoTags

`func (o *PackageDetail) SetRepoTags(v []string)`

SetRepoTags sets RepoTags field to given value.

### HasRepoTags

`func (o *PackageDetail) HasRepoTags() bool`

HasRepoTags returns a boolean if a field has been set.

### GetSummary

`func (o *PackageDetail) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PackageDetail) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PackageDetail) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PackageDetail) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetVersion

`func (o *PackageDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackageDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackageDetail) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackageDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


