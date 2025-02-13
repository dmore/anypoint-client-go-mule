/*
 * Anypoint MQ Exchange Binding specfication
 *
 * Anypoint MQ Exchange Binding API specification
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ame_binding

import (
	"encoding/json"
)

// ExchangeBindingWithRulesAllOf struct for ExchangeBindingWithRulesAllOf
type ExchangeBindingWithRulesAllOf struct {
	Rules *[]map[string]interface{} `json:"rules,omitempty"`
}

// NewExchangeBindingWithRulesAllOf instantiates a new ExchangeBindingWithRulesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeBindingWithRulesAllOf() *ExchangeBindingWithRulesAllOf {
	this := ExchangeBindingWithRulesAllOf{}
	return &this
}

// NewExchangeBindingWithRulesAllOfWithDefaults instantiates a new ExchangeBindingWithRulesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeBindingWithRulesAllOfWithDefaults() *ExchangeBindingWithRulesAllOf {
	this := ExchangeBindingWithRulesAllOf{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ExchangeBindingWithRulesAllOf) GetRules() []map[string]interface{} {
	if o == nil || o.Rules == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeBindingWithRulesAllOf) GetRulesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ExchangeBindingWithRulesAllOf) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []map[string]interface{} and assigns it to the Rules field.
func (o *ExchangeBindingWithRulesAllOf) SetRules(v []map[string]interface{}) {
	o.Rules = &v
}

func (o ExchangeBindingWithRulesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableExchangeBindingWithRulesAllOf struct {
	value *ExchangeBindingWithRulesAllOf
	isSet bool
}

func (v NullableExchangeBindingWithRulesAllOf) Get() *ExchangeBindingWithRulesAllOf {
	return v.value
}

func (v *NullableExchangeBindingWithRulesAllOf) Set(val *ExchangeBindingWithRulesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeBindingWithRulesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeBindingWithRulesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeBindingWithRulesAllOf(val *ExchangeBindingWithRulesAllOf) *NullableExchangeBindingWithRulesAllOf {
	return &NullableExchangeBindingWithRulesAllOf{value: val, isSet: true}
}

func (v NullableExchangeBindingWithRulesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeBindingWithRulesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


