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

// SupportMessage Support ticket message
type SupportMessage struct {

	// Body Message body
	Body string `json:"body,omitempty"`

	// CreationDate Message creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// From Message sender type
	From string `json:"from,omitempty"`

	// MessageID Message identifier
	MessageID int64 `json:"messageId,omitempty"`

	// TicketID Ticket identifier
	TicketID int64 `json:"ticketId,omitempty"`

	// UpdateDate Message last update date
	UpdateDate *time.Time `json:"updateDate,omitempty"`
}
