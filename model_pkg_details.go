/*
 * RHSM-API
 *
 * API for Red Hat Subscription Management
 *
 * API version: 1.109.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gorhsm

import (
	"time"
)

// PkgDetails struct for PkgDetails
type PkgDetails struct {
	Arch string `json:"arch,omitempty"`
	// Date represents the date format used for API returns
	BuildDate   time.Time `json:"buildDate,omitempty"`
	BuildHost   string    `json:"buildHost,omitempty"`
	Checksum    string    `json:"checksum,omitempty"`
	ContentSets []string  `json:"contentSets,omitempty"`
	Description string    `json:"description,omitempty"`
	Epoch       string    `json:"epoch,omitempty"`
	Group       string    `json:"group,omitempty"`
	Href        string    `json:"href,omitempty"`
	License     string    `json:"license,omitempty"`
	Name        string    `json:"name,omitempty"`
	Release     string    `json:"release,omitempty"`
	Size        int64     `json:"size,omitempty"`
	Summary     string    `json:"summary,omitempty"`
	Version     string    `json:"version,omitempty"`
}
