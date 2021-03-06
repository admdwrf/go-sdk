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

// LicenseTask licenses Todos
type LicenseTask struct {

	// Action This Task description
	Action string `json:"action,omitempty"`

	// DoneDate When was this Task done
	DoneDate *time.Time `json:"doneDate,omitempty"`

	// LastUpdate The last time this Task was updated
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// Name This Task name
	Name string `json:"name,omitempty"`

	// Status Current Taks status
	Status string `json:"status,omitempty"`

	// TaskID This Task id
	TaskID int64 `json:"taskId,omitempty"`

	// TodoDate When was this Task created
	TodoDate *time.Time `json:"todoDate,omitempty"`
}
