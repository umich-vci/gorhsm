/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.113.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// MyErratum MyErratum contains erratum information that affects at least one system
type MyErratum struct {
	AdvisoryId          string `json:"advisoryId,omitempty"`
	AffectedSystemCount int32  `json:"affectedSystemCount,omitempty"`
	Details             string `json:"details,omitempty"`
	// Date represents the date format used for API returns
	PublishDate string `json:"publishDate,omitempty"`
	Synopsis    string `json:"synopsis,omitempty"`
	Systems     string `json:"systems,omitempty"`
	Type        string `json:"type,omitempty"`
}
