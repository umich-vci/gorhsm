# VerifyProviderAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | **string** | base64-encoded cloud instance metadata. For GCE, please use the instance identity token (JWT) as it is. | 
**Signature** | **string** | base64-encoded cloud instance metadata signature. For GCE, please omit this field. | 

## Methods

### NewVerifyProviderAccountRequest

`func NewVerifyProviderAccountRequest(identity string, signature string, ) *VerifyProviderAccountRequest`

NewVerifyProviderAccountRequest instantiates a new VerifyProviderAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyProviderAccountRequestWithDefaults

`func NewVerifyProviderAccountRequestWithDefaults() *VerifyProviderAccountRequest`

NewVerifyProviderAccountRequestWithDefaults instantiates a new VerifyProviderAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *VerifyProviderAccountRequest) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VerifyProviderAccountRequest) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VerifyProviderAccountRequest) SetIdentity(v string)`

SetIdentity sets Identity field to given value.


### GetSignature

`func (o *VerifyProviderAccountRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerifyProviderAccountRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerifyProviderAccountRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


