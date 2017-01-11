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

// TelephonyOvhPabxMenu IVR Menu
type TelephonyOvhPabxMenu struct {

	// GreetSound The id of the OvhPabxSound played to greet
	GreetSound int64 `json:"greetSound,omitempty"`

	// GreetSoundTts The text to speech sound played to greet
	GreetSoundTts int64 `json:"greetSoundTts,omitempty"`

	// InvalidSound The id of the OvhPabxSound played when the caller uses an invalid DTMF
	InvalidSound int64 `json:"invalidSound,omitempty"`

	// InvalidSoundTts The text to speech sound played when the caller uses an invalid DTMF
	InvalidSoundTts int64 `json:"invalidSoundTts,omitempty"`

	MenuID int64 `json:"menuId,omitempty"`

	// Name The name of the menu
	Name string `json:"name,omitempty"`
}