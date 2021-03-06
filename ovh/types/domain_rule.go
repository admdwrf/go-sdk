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

// DomainRule Rule
type DomainRule struct {
	AllowedValues []string `json:"allowedValues,omitempty"`

	Description string `json:"description,omitempty"`

	Fields []string `json:"fields,omitempty"`

	InnerConfigurations []*DomainRule `json:"innerConfigurations,omitempty"`

	Label string `json:"label,omitempty"`

	Required bool `json:"required,omitempty"`

	TType string `json:"type,omitempty"`
}
