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

// CloudBillingViewHourlyVolume HourlyVolume
type CloudBillingViewHourlyVolume struct {

	// Details Detail about volume consumption
	Details []*CloudBillingViewHourlyVolumeDetail `json:"details,omitempty"`

	Quantity *CloudBillingViewQuantity `json:"quantity,omitempty"`

	// Region Region
	Region string `json:"region,omitempty"`

	// TotalPrice Total price
	TotalPrice float64 `json:"totalPrice,omitempty"`

	// TType Volume type
	TType string `json:"type,omitempty"`
}
