/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.109.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// InlineResponse20010 struct for InlineResponse20010
type InlineResponse20010 struct {
	// systemList is a System Slice
	Body       []System     `json:"body,omitempty"`
	Pagination ApiPageParam `json:"pagination,omitempty"`
}
