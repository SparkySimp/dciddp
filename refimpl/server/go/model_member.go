/*
 * DCIDDP
 *
 * The Decentralised Identity Dispatch Protocol (DCIDDP) is a protocol that allows for dispension  of name and pronoun information in a decentralised manner, for rendering in a user interface.  The protocol is designed to be simple, lightweight, and easy to implement. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Member struct {

	// The name of the member
	Name string `json:"name,omitempty"`

	Pronouns []Pronoun `json:"pronouns,omitempty"`

	// The ID of the member
	Id string `json:"id,omitempty"`
}