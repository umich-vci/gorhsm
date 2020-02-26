/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.109.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm
import (
	"time"
)
// SystemList SystemList is the result of the get system list API
type SystemList struct {
	ComplianceStatus string `json:"complianceStatus,omitempty"`
	Details string `json:"details,omitempty"`
	EntitlementQuantity int64 `json:"entitlementQuantity,omitempty"`
	// Date represents the date format used for API returns
	LastCheckin time.Time `json:"lastCheckin,omitempty"`
	SystemName string `json:"systemName,omitempty"`
	TotalEntitlementQuantity int64 `json:"totalEntitlementQuantity,omitempty"`
	Type string `json:"type,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
