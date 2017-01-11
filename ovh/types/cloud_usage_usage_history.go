/*
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

import (
	"time"
)

// CloudUsageHistory UsageHistory
type CloudUsageHistory struct {

	// ID Usage id
	ID string `json:"id,omitempty"`

	// LastUpdate Entry last update
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	Period *CloudUsagePeriod `json:"period,omitempty"`
}