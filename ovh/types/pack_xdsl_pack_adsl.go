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

// PackXdslPackAdsl Pack of xDSL services
type PackXdslPackAdsl struct {
	Capabilities *PackXdslPackCapabilities `json:"capabilities,omitempty"`

	// Description Customer pack description
	Description string `json:"description,omitempty"`

	// OfferDescription Name of the offer
	OfferDescription string `json:"offerDescription,omitempty"`

	OfferPrice *OrderPrice `json:"offerPrice,omitempty"`

	// PackName Name of the xdsl pack
	PackName string `json:"packName,omitempty"`
}