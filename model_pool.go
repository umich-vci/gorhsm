/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.109.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm
// Pool Pool represents pool information that matter for the detail list
type Pool struct {
	Consumed int64 `json:"consumed,omitempty"`
	Id string `json:"id,omitempty"`
	Quantity int64 `json:"quantity,omitempty"`
	Type string `json:"type,omitempty"`
}
