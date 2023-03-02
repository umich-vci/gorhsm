/*
RHSM-API

API for Red Hat Subscription Management

API version: 1.313.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"encoding/json"
)

// checks if the ErratumDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErratumDetails{}

// ErratumDetails struct for ErratumDetails
type ErratumDetails struct {
	AffectedProducts []string `json:"affectedProducts,omitempty"`
	Bugzillas []Reference `json:"bugzillas,omitempty"`
	CopyrightYear *string `json:"copyrightYear,omitempty"`
	Cves *string `json:"cves,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	// Date represents the date format used for API returns
	Issued *string `json:"issued,omitempty"`
	// Date represents the date format used for API returns
	LastUpdated *string `json:"lastUpdated,omitempty"`
	References []Reference `json:"references,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Solution *string `json:"solution,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Synopsis *string `json:"synopsis,omitempty"`
	Type *string `json:"type,omitempty"`
	TypeSeverity *string `json:"typeSeverity,omitempty"`
}

// NewErratumDetails instantiates a new ErratumDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErratumDetails() *ErratumDetails {
	this := ErratumDetails{}
	return &this
}

// NewErratumDetailsWithDefaults instantiates a new ErratumDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErratumDetailsWithDefaults() *ErratumDetails {
	this := ErratumDetails{}
	return &this
}

// GetAffectedProducts returns the AffectedProducts field value if set, zero value otherwise.
func (o *ErratumDetails) GetAffectedProducts() []string {
	if o == nil || IsNil(o.AffectedProducts) {
		var ret []string
		return ret
	}
	return o.AffectedProducts
}

// GetAffectedProductsOk returns a tuple with the AffectedProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetAffectedProductsOk() ([]string, bool) {
	if o == nil || IsNil(o.AffectedProducts) {
		return nil, false
	}
	return o.AffectedProducts, true
}

// HasAffectedProducts returns a boolean if a field has been set.
func (o *ErratumDetails) HasAffectedProducts() bool {
	if o != nil && !IsNil(o.AffectedProducts) {
		return true
	}

	return false
}

// SetAffectedProducts gets a reference to the given []string and assigns it to the AffectedProducts field.
func (o *ErratumDetails) SetAffectedProducts(v []string) {
	o.AffectedProducts = v
}

// GetBugzillas returns the Bugzillas field value if set, zero value otherwise.
func (o *ErratumDetails) GetBugzillas() []Reference {
	if o == nil || IsNil(o.Bugzillas) {
		var ret []Reference
		return ret
	}
	return o.Bugzillas
}

// GetBugzillasOk returns a tuple with the Bugzillas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetBugzillasOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Bugzillas) {
		return nil, false
	}
	return o.Bugzillas, true
}

// HasBugzillas returns a boolean if a field has been set.
func (o *ErratumDetails) HasBugzillas() bool {
	if o != nil && !IsNil(o.Bugzillas) {
		return true
	}

	return false
}

// SetBugzillas gets a reference to the given []Reference and assigns it to the Bugzillas field.
func (o *ErratumDetails) SetBugzillas(v []Reference) {
	o.Bugzillas = v
}

// GetCopyrightYear returns the CopyrightYear field value if set, zero value otherwise.
func (o *ErratumDetails) GetCopyrightYear() string {
	if o == nil || IsNil(o.CopyrightYear) {
		var ret string
		return ret
	}
	return *o.CopyrightYear
}

// GetCopyrightYearOk returns a tuple with the CopyrightYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetCopyrightYearOk() (*string, bool) {
	if o == nil || IsNil(o.CopyrightYear) {
		return nil, false
	}
	return o.CopyrightYear, true
}

// HasCopyrightYear returns a boolean if a field has been set.
func (o *ErratumDetails) HasCopyrightYear() bool {
	if o != nil && !IsNil(o.CopyrightYear) {
		return true
	}

	return false
}

// SetCopyrightYear gets a reference to the given string and assigns it to the CopyrightYear field.
func (o *ErratumDetails) SetCopyrightYear(v string) {
	o.CopyrightYear = &v
}

// GetCves returns the Cves field value if set, zero value otherwise.
func (o *ErratumDetails) GetCves() string {
	if o == nil || IsNil(o.Cves) {
		var ret string
		return ret
	}
	return *o.Cves
}

// GetCvesOk returns a tuple with the Cves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetCvesOk() (*string, bool) {
	if o == nil || IsNil(o.Cves) {
		return nil, false
	}
	return o.Cves, true
}

// HasCves returns a boolean if a field has been set.
func (o *ErratumDetails) HasCves() bool {
	if o != nil && !IsNil(o.Cves) {
		return true
	}

	return false
}

// SetCves gets a reference to the given string and assigns it to the Cves field.
func (o *ErratumDetails) SetCves(v string) {
	o.Cves = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ErratumDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ErratumDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ErratumDetails) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ErratumDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ErratumDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ErratumDetails) SetId(v string) {
	o.Id = &v
}

// GetIssued returns the Issued field value if set, zero value otherwise.
func (o *ErratumDetails) GetIssued() string {
	if o == nil || IsNil(o.Issued) {
		var ret string
		return ret
	}
	return *o.Issued
}

// GetIssuedOk returns a tuple with the Issued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetIssuedOk() (*string, bool) {
	if o == nil || IsNil(o.Issued) {
		return nil, false
	}
	return o.Issued, true
}

// HasIssued returns a boolean if a field has been set.
func (o *ErratumDetails) HasIssued() bool {
	if o != nil && !IsNil(o.Issued) {
		return true
	}

	return false
}

// SetIssued gets a reference to the given string and assigns it to the Issued field.
func (o *ErratumDetails) SetIssued(v string) {
	o.Issued = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ErratumDetails) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ErratumDetails) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *ErratumDetails) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ErratumDetails) GetReferences() []Reference {
	if o == nil || IsNil(o.References) {
		var ret []Reference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetReferencesOk() ([]Reference, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ErratumDetails) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []Reference and assigns it to the References field.
func (o *ErratumDetails) SetReferences(v []Reference) {
	o.References = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ErratumDetails) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ErratumDetails) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ErratumDetails) SetSeverity(v string) {
	o.Severity = &v
}

// GetSolution returns the Solution field value if set, zero value otherwise.
func (o *ErratumDetails) GetSolution() string {
	if o == nil || IsNil(o.Solution) {
		var ret string
		return ret
	}
	return *o.Solution
}

// GetSolutionOk returns a tuple with the Solution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetSolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Solution) {
		return nil, false
	}
	return o.Solution, true
}

// HasSolution returns a boolean if a field has been set.
func (o *ErratumDetails) HasSolution() bool {
	if o != nil && !IsNil(o.Solution) {
		return true
	}

	return false
}

// SetSolution gets a reference to the given string and assigns it to the Solution field.
func (o *ErratumDetails) SetSolution(v string) {
	o.Solution = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ErratumDetails) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ErratumDetails) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ErratumDetails) SetSummary(v string) {
	o.Summary = &v
}

// GetSynopsis returns the Synopsis field value if set, zero value otherwise.
func (o *ErratumDetails) GetSynopsis() string {
	if o == nil || IsNil(o.Synopsis) {
		var ret string
		return ret
	}
	return *o.Synopsis
}

// GetSynopsisOk returns a tuple with the Synopsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetSynopsisOk() (*string, bool) {
	if o == nil || IsNil(o.Synopsis) {
		return nil, false
	}
	return o.Synopsis, true
}

// HasSynopsis returns a boolean if a field has been set.
func (o *ErratumDetails) HasSynopsis() bool {
	if o != nil && !IsNil(o.Synopsis) {
		return true
	}

	return false
}

// SetSynopsis gets a reference to the given string and assigns it to the Synopsis field.
func (o *ErratumDetails) SetSynopsis(v string) {
	o.Synopsis = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErratumDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErratumDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErratumDetails) SetType(v string) {
	o.Type = &v
}

// GetTypeSeverity returns the TypeSeverity field value if set, zero value otherwise.
func (o *ErratumDetails) GetTypeSeverity() string {
	if o == nil || IsNil(o.TypeSeverity) {
		var ret string
		return ret
	}
	return *o.TypeSeverity
}

// GetTypeSeverityOk returns a tuple with the TypeSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErratumDetails) GetTypeSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.TypeSeverity) {
		return nil, false
	}
	return o.TypeSeverity, true
}

// HasTypeSeverity returns a boolean if a field has been set.
func (o *ErratumDetails) HasTypeSeverity() bool {
	if o != nil && !IsNil(o.TypeSeverity) {
		return true
	}

	return false
}

// SetTypeSeverity gets a reference to the given string and assigns it to the TypeSeverity field.
func (o *ErratumDetails) SetTypeSeverity(v string) {
	o.TypeSeverity = &v
}

func (o ErratumDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErratumDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedProducts) {
		toSerialize["affectedProducts"] = o.AffectedProducts
	}
	if !IsNil(o.Bugzillas) {
		toSerialize["bugzillas"] = o.Bugzillas
	}
	if !IsNil(o.CopyrightYear) {
		toSerialize["copyrightYear"] = o.CopyrightYear
	}
	if !IsNil(o.Cves) {
		toSerialize["cves"] = o.Cves
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Issued) {
		toSerialize["issued"] = o.Issued
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Solution) {
		toSerialize["solution"] = o.Solution
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Synopsis) {
		toSerialize["synopsis"] = o.Synopsis
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeSeverity) {
		toSerialize["typeSeverity"] = o.TypeSeverity
	}
	return toSerialize, nil
}

type NullableErratumDetails struct {
	value *ErratumDetails
	isSet bool
}

func (v NullableErratumDetails) Get() *ErratumDetails {
	return v.value
}

func (v *NullableErratumDetails) Set(val *ErratumDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableErratumDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableErratumDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErratumDetails(val *ErratumDetails) *NullableErratumDetails {
	return &NullableErratumDetails{value: val, isSet: true}
}

func (v NullableErratumDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErratumDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


