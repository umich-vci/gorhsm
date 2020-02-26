/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.109.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	Body       []PkgDetails `json:"body,omitempty"`
	Pagination ApiPageParam `json:"pagination,omitempty"`
}
