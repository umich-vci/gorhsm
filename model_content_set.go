/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.109.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// ContentSet struct for ContentSet
type ContentSet struct {
	Arch    string `json:"arch,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Label   string `json:"label,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
}
