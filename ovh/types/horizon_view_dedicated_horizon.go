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

// HorizonViewDedicatedHorizon Horizon View as a Service
type HorizonViewDedicatedHorizon struct {
	NumberOfAvailableIP *HorizonViewDedicatedHorizonNumberOfAvailableIP `json:"numberOfAvailableIp,omitempty"`

	// Pack Your Cloud Desktop Infrastructure pack
	Pack string `json:"pack,omitempty"`

	// PrivateCloudName Your Horizon Private Cloud
	PrivateCloudName string `json:"privateCloudName,omitempty"`

	// PrivateCloudZone The location of your datacenter
	PrivateCloudZone string `json:"privateCloudZone,omitempty"`

	// PublicURL Url of your Dedicated Horizon
	PublicURL string `json:"publicUrl,omitempty"`

	// ServiceName Domain of your Dedicated Horizon
	ServiceName string `json:"serviceName,omitempty"`

	// State Current state of your Dedicated Horizon
	State string `json:"state,omitempty"`

	// Version Version of your Dedicated Horizon
	Version string `json:"version,omitempty"`
}