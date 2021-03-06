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

// IPLoadBalancingTask List of tasks associated with your IP load balancing
type IPLoadBalancingTask struct {

	// Action The action made
	Action string `json:"action,omitempty"`

	// CreationDate Creation date of your task
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// ID Identifier of your task
	ID int64 `json:"id,omitempty"`

	// Status Current status of your task
	Status string `json:"status,omitempty"`
}
