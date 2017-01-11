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

// CloudRegionDetail RegionDetail
type CloudRegionDetail struct {

	// ContinentCode Region continent code
	ContinentCode string `json:"continentCode,omitempty"`

	// DatacenterLocation Location of the datacenter where the region is
	DatacenterLocation string `json:"datacenterLocation,omitempty"`

	// Name Region name
	Name string `json:"name,omitempty"`

	// Services Details about openstack services status
	Services []*CloudRegionOpenstackService `json:"services,omitempty"`
}