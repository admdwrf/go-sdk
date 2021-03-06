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

// TelephonyOvhPabxHuntingAgentLiveStatus Live statistics of the queue
type TelephonyOvhPabxHuntingAgentLiveStatus struct {

	// AnsweredCalls The number of calls this agent took on the current day
	AnsweredCalls int64 `json:"answeredCalls,omitempty"`

	// LastStatusChange Last status change date
	LastStatusChange *time.Time `json:"lastStatusChange,omitempty"`

	// Status Current status of the agent
	Status string `json:"status,omitempty"`

	// TotalCallDuration The total duration in seconds of the calls this agent took on the current day
	TotalCallDuration int64 `json:"totalCallDuration,omitempty"`
}
