/*
 * Identity Provider Management API
 *
 * Description of Identity Provider API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// Urls1 struct for Urls1
type Urls1 struct {
	SignOn *string `json:"sign_on,omitempty"`
	// only available for SAML
	SignOut *string `json:"sign_out,omitempty"`
}

// NewUrls1 instantiates a new Urls1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrls1() *Urls1 {
	this := Urls1{}
	return &this
}

// NewUrls1WithDefaults instantiates a new Urls1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrls1WithDefaults() *Urls1 {
	this := Urls1{}
	return &this
}

// GetSignOn returns the SignOn field value if set, zero value otherwise.
func (o *Urls1) GetSignOn() string {
	if o == nil || o.SignOn == nil {
		var ret string
		return ret
	}
	return *o.SignOn
}

// GetSignOnOk returns a tuple with the SignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Urls1) GetSignOnOk() (*string, bool) {
	if o == nil || o.SignOn == nil {
		return nil, false
	}
	return o.SignOn, true
}

// HasSignOn returns a boolean if a field has been set.
func (o *Urls1) HasSignOn() bool {
	if o != nil && o.SignOn != nil {
		return true
	}

	return false
}

// SetSignOn gets a reference to the given string and assigns it to the SignOn field.
func (o *Urls1) SetSignOn(v string) {
	o.SignOn = &v
}

// GetSignOut returns the SignOut field value if set, zero value otherwise.
func (o *Urls1) GetSignOut() string {
	if o == nil || o.SignOut == nil {
		var ret string
		return ret
	}
	return *o.SignOut
}

// GetSignOutOk returns a tuple with the SignOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Urls1) GetSignOutOk() (*string, bool) {
	if o == nil || o.SignOut == nil {
		return nil, false
	}
	return o.SignOut, true
}

// HasSignOut returns a boolean if a field has been set.
func (o *Urls1) HasSignOut() bool {
	if o != nil && o.SignOut != nil {
		return true
	}

	return false
}

// SetSignOut gets a reference to the given string and assigns it to the SignOut field.
func (o *Urls1) SetSignOut(v string) {
	o.SignOut = &v
}

func (o Urls1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SignOn != nil {
		toSerialize["sign_on"] = o.SignOn
	}
	if o.SignOut != nil {
		toSerialize["sign_out"] = o.SignOut
	}
	return json.Marshal(toSerialize)
}

type NullableUrls1 struct {
	value *Urls1
	isSet bool
}

func (v NullableUrls1) Get() *Urls1 {
	return v.value
}

func (v *NullableUrls1) Set(val *Urls1) {
	v.value = val
	v.isSet = true
}

func (v NullableUrls1) IsSet() bool {
	return v.isSet
}

func (v *NullableUrls1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrls1(val *Urls1) *NullableUrls1 {
	return &NullableUrls1{value: val, isSet: true}
}

func (v NullableUrls1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrls1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


