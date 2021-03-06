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

// DBaasQueueKeyWithSecret KeyWithSecret
type DBaasQueueKeyWithSecret struct {

	// HumanAppID Human ID of the key's application
	HumanAppID string `json:"humanAppId,omitempty"`

	// ID Key ID
	ID string `json:"id,omitempty"`

	// Name Key name
	Name string `json:"name,omitempty"`

	// Secret Key secret
	Secret string `json:"secret,omitempty"`
}
