/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.143.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

// PackageForSystem package installed on a system
type PackageForSystem struct {
	Advisories  []PackageForSystemAdvisories `json:"advisories,omitempty"`
	Arch        string                       `json:"arch,omitempty"`
	Epoch       int32                        `json:"epoch,omitempty"`
	ErrataCount int32                        `json:"errataCount,omitempty"`
	Name        string                       `json:"name,omitempty"`
	Release     string                       `json:"release,omitempty"`
	Version     string                       `json:"version,omitempty"`
}
