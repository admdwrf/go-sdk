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

// TelephonyOvhPabxMenuEntryPost ...
type TelephonyOvhPabxMenuEntryPost struct {
	Action string `json:"action,omitempty"`

	ActionParam string `json:"actionParam,omitempty"`

	Dtmf string `json:"dtmf,omitempty"`

	Position int64 `json:"position,omitempty"`
}
