/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.109.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm
// EntitlementsAttachedResponse EntitlementsAttachedResponse wraps data obtained for EntitlementsAttached and sends metadata about it using helpers.OptionalResult
type EntitlementsAttachedResponse struct {
	Reason string `json:"reason,omitempty"`
	Valid bool `json:"valid,omitempty"`
	Value []EntitlementsAttachedResponseValue `json:"value,omitempty"`
}
