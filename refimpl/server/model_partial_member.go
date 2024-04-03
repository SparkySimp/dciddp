/*
DCIDDP

The Decentralised Identity Dispatch Protocol (DCIDDP) is a protocol that allows for dispension  of name and pronoun information in a decentralised manner, for rendering in a user interface.  The protocol is designed to be simple, lightweight, and easy to implement. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PartialMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialMember{}

// PartialMember struct for PartialMember
type PartialMember struct {
	// The name of the member
	Name *string `json:"name,omitempty"`
	Pronouns []Pronoun `json:"pronouns,omitempty"`
}

// NewPartialMember instantiates a new PartialMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialMember() *PartialMember {
	this := PartialMember{}
	return &this
}

// NewPartialMemberWithDefaults instantiates a new PartialMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialMemberWithDefaults() *PartialMember {
	this := PartialMember{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartialMember) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialMember) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartialMember) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartialMember) SetName(v string) {
	o.Name = &v
}

// GetPronouns returns the Pronouns field value if set, zero value otherwise.
func (o *PartialMember) GetPronouns() []Pronoun {
	if o == nil || IsNil(o.Pronouns) {
		var ret []Pronoun
		return ret
	}
	return o.Pronouns
}

// GetPronounsOk returns a tuple with the Pronouns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialMember) GetPronounsOk() ([]Pronoun, bool) {
	if o == nil || IsNil(o.Pronouns) {
		return nil, false
	}
	return o.Pronouns, true
}

// HasPronouns returns a boolean if a field has been set.
func (o *PartialMember) HasPronouns() bool {
	if o != nil && !IsNil(o.Pronouns) {
		return true
	}

	return false
}

// SetPronouns gets a reference to the given []Pronoun and assigns it to the Pronouns field.
func (o *PartialMember) SetPronouns(v []Pronoun) {
	o.Pronouns = v
}

func (o PartialMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartialMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Pronouns) {
		toSerialize["pronouns"] = o.Pronouns
	}
	return toSerialize, nil
}

type NullablePartialMember struct {
	value *PartialMember
	isSet bool
}

func (v NullablePartialMember) Get() *PartialMember {
	return v.value
}

func (v *NullablePartialMember) Set(val *PartialMember) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialMember) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialMember(val *PartialMember) *NullablePartialMember {
	return &NullablePartialMember{value: val, isSet: true}
}

func (v NullablePartialMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


