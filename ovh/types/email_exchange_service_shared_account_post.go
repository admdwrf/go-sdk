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

// EmailExchangeServiceSharedAccountPost ...
type EmailExchangeServiceSharedAccountPost struct {
	DisplayName string `json:"displayName,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	HiddenFromGAL bool `json:"hiddenFromGAL,omitempty"`

	Initials string `json:"initials,omitempty"`

	LastName string `json:"lastName,omitempty"`

	MailingFilter []string `json:"mailingFilter,omitempty"`

	Quota int64 `json:"quota,omitempty"`

	SharedEmailAddress string `json:"sharedEmailAddress,omitempty"`
}