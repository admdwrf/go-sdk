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

// DedicatedServerIPOrderable A structure describing informations about orderable IP address
type DedicatedServerIPOrderable struct {

	// IPv4 Orderable IP v4 details
	IPv4 []*DedicatedServerIPOrderableDetails `json:"ipv4,omitempty"`

	// IPv6 Orderable IP v6 details
	IPv6 []*DedicatedServerIPOrderableDetails `json:"ipv6,omitempty"`
}