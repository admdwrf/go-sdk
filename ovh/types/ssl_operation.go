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

// SSLOperation Task on a SSL
type SSLOperation struct {

	// DoneDate Completion date
	DoneDate *time.Time `json:"doneDate,omitempty"`

	// Function Task function name
	Function string `json:"function,omitempty"`

	// LastUpdate Task last update
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// StartDate Task Creation date
	StartDate *time.Time `json:"startDate,omitempty"`

	// Status Task status
	Status string `json:"status,omitempty"`

	TaskID int64 `json:"taskId,omitempty"`
}