# ErratumDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **[]string** |  | [optional] 
**Bugzillas** | Pointer to [**[]Reference**](Reference.md) |  | [optional] 
**CopyrightYear** | Pointer to **string** |  | [optional] 
**Cves** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Issued** | Pointer to **string** | Date represents the date format used for API returns | [optional] 
**LastUpdated** | Pointer to **string** | Date represents the date format used for API returns | [optional] 
**References** | Pointer to [**[]Reference**](Reference.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Solution** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Synopsis** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TypeSeverity** | Pointer to **string** |  | [optional] 

## Methods

### NewErratumDetails

`func NewErratumDetails() *ErratumDetails`

NewErratumDetails instantiates a new ErratumDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErratumDetailsWithDefaults

`func NewErratumDetailsWithDefaults() *ErratumDetails`

NewErratumDetailsWithDefaults instantiates a new ErratumDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *ErratumDetails) GetAffectedProducts() []string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *ErratumDetails) GetAffectedProductsOk() (*[]string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *ErratumDetails) SetAffectedProducts(v []string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *ErratumDetails) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetBugzillas

`func (o *ErratumDetails) GetBugzillas() []Reference`

GetBugzillas returns the Bugzillas field if non-nil, zero value otherwise.

### GetBugzillasOk

`func (o *ErratumDetails) GetBugzillasOk() (*[]Reference, bool)`

GetBugzillasOk returns a tuple with the Bugzillas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugzillas

`func (o *ErratumDetails) SetBugzillas(v []Reference)`

SetBugzillas sets Bugzillas field to given value.

### HasBugzillas

`func (o *ErratumDetails) HasBugzillas() bool`

HasBugzillas returns a boolean if a field has been set.

### GetCopyrightYear

`func (o *ErratumDetails) GetCopyrightYear() string`

GetCopyrightYear returns the CopyrightYear field if non-nil, zero value otherwise.

### GetCopyrightYearOk

`func (o *ErratumDetails) GetCopyrightYearOk() (*string, bool)`

GetCopyrightYearOk returns a tuple with the CopyrightYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightYear

`func (o *ErratumDetails) SetCopyrightYear(v string)`

SetCopyrightYear sets CopyrightYear field to given value.

### HasCopyrightYear

`func (o *ErratumDetails) HasCopyrightYear() bool`

HasCopyrightYear returns a boolean if a field has been set.

### GetCves

`func (o *ErratumDetails) GetCves() string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *ErratumDetails) GetCvesOk() (*string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *ErratumDetails) SetCves(v string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *ErratumDetails) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetDescription

`func (o *ErratumDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErratumDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErratumDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErratumDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ErratumDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErratumDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErratumDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ErratumDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssued

`func (o *ErratumDetails) GetIssued() string`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *ErratumDetails) GetIssuedOk() (*string, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *ErratumDetails) SetIssued(v string)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *ErratumDetails) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ErratumDetails) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ErratumDetails) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ErratumDetails) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ErratumDetails) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetReferences

`func (o *ErratumDetails) GetReferences() []Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ErratumDetails) GetReferencesOk() (*[]Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ErratumDetails) SetReferences(v []Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ErratumDetails) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *ErratumDetails) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ErratumDetails) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ErratumDetails) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ErratumDetails) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSolution

`func (o *ErratumDetails) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *ErratumDetails) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *ErratumDetails) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *ErratumDetails) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetSummary

`func (o *ErratumDetails) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ErratumDetails) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ErratumDetails) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ErratumDetails) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSynopsis

`func (o *ErratumDetails) GetSynopsis() string`

GetSynopsis returns the Synopsis field if non-nil, zero value otherwise.

### GetSynopsisOk

`func (o *ErratumDetails) GetSynopsisOk() (*string, bool)`

GetSynopsisOk returns a tuple with the Synopsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynopsis

`func (o *ErratumDetails) SetSynopsis(v string)`

SetSynopsis sets Synopsis field to given value.

### HasSynopsis

`func (o *ErratumDetails) HasSynopsis() bool`

HasSynopsis returns a boolean if a field has been set.

### GetType

`func (o *ErratumDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErratumDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErratumDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErratumDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeSeverity

`func (o *ErratumDetails) GetTypeSeverity() string`

GetTypeSeverity returns the TypeSeverity field if non-nil, zero value otherwise.

### GetTypeSeverityOk

`func (o *ErratumDetails) GetTypeSeverityOk() (*string, bool)`

GetTypeSeverityOk returns a tuple with the TypeSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSeverity

`func (o *ErratumDetails) SetTypeSeverity(v string)`

SetTypeSeverity sets TypeSeverity field to given value.

### HasTypeSeverity

`func (o *ErratumDetails) HasTypeSeverity() bool`

HasTypeSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


