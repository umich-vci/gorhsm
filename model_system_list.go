/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.113.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// SystemList SystemList is the result of the get system list API
type SystemList struct {
	ComplianceStatus    string `json:"complianceStatus,omitempty"`
	Details             string `json:"details,omitempty"`
	EntitlementQuantity int32  `json:"entitlementQuantity,omitempty"`
	// Date represents the date format used for API returns
	LastCheckin              string `json:"lastCheckin,omitempty"`
	SystemName               string `json:"systemName,omitempty"`
	TotalEntitlementQuantity int32  `json:"totalEntitlementQuantity,omitempty"`
	Type                     string `json:"type,omitempty"`
	Uuid                     string `json:"uuid,omitempty"`
}
