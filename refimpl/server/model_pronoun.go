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

// checks if the Pronoun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pronoun{}

// Pronoun struct for Pronoun
type Pronoun struct {
	// The subject pronoun
	Subject *string `json:"subject,omitempty"`
	// The object pronoun
	Object *string `json:"object,omitempty"`
	// The possessive pronoun
	Possessive *string `json:"possessive,omitempty"`
	// The possessive determiner
	PossessiveDeterminer *string `json:"possessive_determiner,omitempty"`
	// The reflexive pronoun
	Reflexive *string `json:"reflexive,omitempty"`
}

// NewPronoun instantiates a new Pronoun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPronoun() *Pronoun {
	this := Pronoun{}
	return &this
}

// NewPronounWithDefaults instantiates a new Pronoun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPronounWithDefaults() *Pronoun {
	this := Pronoun{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Pronoun) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pronoun) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Pronoun) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Pronoun) SetSubject(v string) {
	o.Subject = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Pronoun) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pronoun) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Pronoun) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Pronoun) SetObject(v string) {
	o.Object = &v
}

// GetPossessive returns the Possessive field value if set, zero value otherwise.
func (o *Pronoun) GetPossessive() string {
	if o == nil || IsNil(o.Possessive) {
		var ret string
		return ret
	}
	return *o.Possessive
}

// GetPossessiveOk returns a tuple with the Possessive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pronoun) GetPossessiveOk() (*string, bool) {
	if o == nil || IsNil(o.Possessive) {
		return nil, false
	}
	return o.Possessive, true
}

// HasPossessive returns a boolean if a field has been set.
func (o *Pronoun) HasPossessive() bool {
	if o != nil && !IsNil(o.Possessive) {
		return true
	}

	return false
}

// SetPossessive gets a reference to the given string and assigns it to the Possessive field.
func (o *Pronoun) SetPossessive(v string) {
	o.Possessive = &v
}

// GetPossessiveDeterminer returns the PossessiveDeterminer field value if set, zero value otherwise.
func (o *Pronoun) GetPossessiveDeterminer() string {
	if o == nil || IsNil(o.PossessiveDeterminer) {
		var ret string
		return ret
	}
	return *o.PossessiveDeterminer
}

// GetPossessiveDeterminerOk returns a tuple with the PossessiveDeterminer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pronoun) GetPossessiveDeterminerOk() (*string, bool) {
	if o == nil || IsNil(o.PossessiveDeterminer) {
		return nil, false
	}
	return o.PossessiveDeterminer, true
}

// HasPossessiveDeterminer returns a boolean if a field has been set.
func (o *Pronoun) HasPossessiveDeterminer() bool {
	if o != nil && !IsNil(o.PossessiveDeterminer) {
		return true
	}

	return false
}

// SetPossessiveDeterminer gets a reference to the given string and assigns it to the PossessiveDeterminer field.
func (o *Pronoun) SetPossessiveDeterminer(v string) {
	o.PossessiveDeterminer = &v
}

// GetReflexive returns the Reflexive field value if set, zero value otherwise.
func (o *Pronoun) GetReflexive() string {
	if o == nil || IsNil(o.Reflexive) {
		var ret string
		return ret
	}
	return *o.Reflexive
}

// GetReflexiveOk returns a tuple with the Reflexive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pronoun) GetReflexiveOk() (*string, bool) {
	if o == nil || IsNil(o.Reflexive) {
		return nil, false
	}
	return o.Reflexive, true
}

// HasReflexive returns a boolean if a field has been set.
func (o *Pronoun) HasReflexive() bool {
	if o != nil && !IsNil(o.Reflexive) {
		return true
	}

	return false
}

// SetReflexive gets a reference to the given string and assigns it to the Reflexive field.
func (o *Pronoun) SetReflexive(v string) {
	o.Reflexive = &v
}

func (o Pronoun) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pronoun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Possessive) {
		toSerialize["possessive"] = o.Possessive
	}
	if !IsNil(o.PossessiveDeterminer) {
		toSerialize["possessive_determiner"] = o.PossessiveDeterminer
	}
	if !IsNil(o.Reflexive) {
		toSerialize["reflexive"] = o.Reflexive
	}
	return toSerialize, nil
}

type NullablePronoun struct {
	value *Pronoun
	isSet bool
}

func (v NullablePronoun) Get() *Pronoun {
	return v.value
}

func (v *NullablePronoun) Set(val *Pronoun) {
	v.value = val
	v.isSet = true
}

func (v NullablePronoun) IsSet() bool {
	return v.isSet
}

func (v *NullablePronoun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePronoun(val *Pronoun) *NullablePronoun {
	return &NullablePronoun{value: val, isSet: true}
}

func (v NullablePronoun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePronoun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

