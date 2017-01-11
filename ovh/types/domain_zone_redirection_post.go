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

// DomainZoneRedirectionPost ...
type DomainZoneRedirectionPost struct {
	Description string `json:"description,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	SubDomain string `json:"subDomain,omitempty"`

	Target string `json:"target,omitempty"`

	Title string `json:"title,omitempty"`

	TType string `json:"type,omitempty"`
}