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

// HostingPrivateDatabaseTask Tasks
type HostingPrivateDatabaseTask struct {

	// DatabaseName Database name on which the task is working
	DatabaseName string `json:"databaseName,omitempty"`

	// DoneDate Completion date
	DoneDate *time.Time `json:"doneDate,omitempty"`

	// DumpID DumpId on which the task is working
	DumpID int64 `json:"dumpId,omitempty"`

	// Function Function name
	Function string `json:"function,omitempty"`

	// ID The id of the task
	ID int64 `json:"id,omitempty"`

	// LastUpdate Last update
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// StartDate Task Creation date
	StartDate *time.Time `json:"startDate,omitempty"`

	// Status Task status
	Status string `json:"status,omitempty"`

	// UserName User name on which the task is working
	UserName string `json:"userName,omitempty"`
}
