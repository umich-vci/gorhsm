/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.113.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// EntitlementsAttachedResponseValue EntitlementsAttachedResponseValue represents the Value field in the EntitlementsAttachedResponse
type EntitlementsAttachedResponseValue struct {
	ContractNumber      string `json:"contractNumber,omitempty"`
	EntitlementQuantity int32  `json:"entitlementQuantity,omitempty"`
	Id                  string `json:"id,omitempty"`
	Sku                 string `json:"sku,omitempty"`
}
