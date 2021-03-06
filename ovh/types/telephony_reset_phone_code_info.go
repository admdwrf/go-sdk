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

// TelephonyResetPhoneCodeInfo Relevant informations about reset code
type TelephonyResetPhoneCodeInfo struct {

	// ActivationCode Activation code
	ActivationCode string `json:"activationCode,omitempty"`

	// KeyBook Key book url
	KeyBook string `json:"keyBook,omitempty"`

	// ServerURL Server url
	ServerURL string `json:"serverURL,omitempty"`
}
