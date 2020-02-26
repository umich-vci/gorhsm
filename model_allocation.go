/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.109.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm
// Allocation struct for Allocation
type Allocation struct {
	EntitlementQuantity int64 `json:"entitlementQuantity,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Version string `json:"version,omitempty"`
}
