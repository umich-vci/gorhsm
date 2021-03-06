/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.143.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// ErrorDetails ErrorDetails details the Error in ErrorResponse
type ErrorDetails struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
