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

// SecondaryDNSNameServer A structure describing informations about available nameserver for secondary dns
type SecondaryDNSNameServer struct {

	// Hostname the name server
	Hostname string `json:"hostname,omitempty"`

	IP string `json:"ip,omitempty"`

	IPv6 string `json:"ipv6,omitempty"`
}