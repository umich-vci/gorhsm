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
// System struct for System
type System struct {
	EntitlementCount int64 `json:"entitlementCount,omitempty"`
	EntitlementStatus string `json:"entitlementStatus,omitempty"`
	ErrataCounts ErrataCount `json:"errataCounts,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Href string `json:"href,omitempty"`
	// Date represents the date format used for API returns
	LastCheckin time.Time `json:"lastCheckin,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Username string `json:"username,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
