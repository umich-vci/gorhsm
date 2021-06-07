# DownloadLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to **string** | time at which the download link expires (in UTC) | [optional] 
**Filename** | Pointer to **string** | filename of the file on the download link | [optional] 
**Href** | Pointer to **string** | URL to obtain the image | [optional] 

## Methods

### NewDownloadLink

`func NewDownloadLink() *DownloadLink`

NewDownloadLink instantiates a new DownloadLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadLinkWithDefaults

`func NewDownloadLinkWithDefaults() *DownloadLink`

NewDownloadLinkWithDefaults instantiates a new DownloadLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *DownloadLink) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *DownloadLink) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *DownloadLink) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *DownloadLink) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetFilename

`func (o *DownloadLink) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *DownloadLink) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *DownloadLink) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *DownloadLink) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetHref

`func (o *DownloadLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DownloadLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DownloadLink) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DownloadLink) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


