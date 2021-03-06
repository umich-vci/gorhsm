/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.143.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// InlineResponse2009 struct for InlineResponse2009
type InlineResponse2009 struct {
	// ListResponse is the actual collection of subscription details that gets rendered
	Body       []DetailResponse `json:"body,omitempty"`
	Pagination ApiPageParam     `json:"pagination,omitempty"`
}
