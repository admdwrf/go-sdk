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

// SmsIncoming Sms history of sms incoming received
type SmsIncoming struct {
	CreationDatetime *time.Time `json:"creationDatetime,omitempty"`

	Credits float64 `json:"credits,omitempty"`

	ID int64 `json:"id,omitempty"`

	Message string `json:"message,omitempty"`

	Sender string `json:"sender,omitempty"`

	Tag string `json:"tag,omitempty"`
}