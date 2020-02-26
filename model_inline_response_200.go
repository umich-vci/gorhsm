/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.109.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Body       []Allocation `json:"body,omitempty"`
	Pagination ApiPageParam `json:"pagination,omitempty"`
}
