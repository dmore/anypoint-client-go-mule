/*
 * Organization API
 *
 * Description of the Organization API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// TheDesignCenterSchema An explanation about the purpose of this instance.
type TheDesignCenterSchema struct {
	// An explanation about the purpose of this instance.
	Api *bool `json:"api,omitempty"`
	// An explanation about the purpose of this instance.
	Mozart *bool `json:"mozart,omitempty"`
}

// NewTheDesignCenterSchema instantiates a new TheDesignCenterSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTheDesignCenterSchema() *TheDesignCenterSchema {
	this := TheDesignCenterSchema{}
	var api bool = false
	this.Api = &api
	var mozart bool = false
	this.Mozart = &mozart
	return &this
}

// NewTheDesignCenterSchemaWithDefaults instantiates a new TheDesignCenterSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTheDesignCenterSchemaWithDefaults() *TheDesignCenterSchema {
	this := TheDesignCenterSchema{}
	var api bool = false
	this.Api = &api
	var mozart bool = false
	this.Mozart = &mozart
	return &this
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *TheDesignCenterSchema) GetApi() bool {
	if o == nil || o.Api == nil {
		var ret bool
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TheDesignCenterSchema) GetApiOk() (*bool, bool) {
	if o == nil || o.Api == nil {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *TheDesignCenterSchema) HasApi() bool {
	if o != nil && o.Api != nil {
		return true
	}

	return false
}

// SetApi gets a reference to the given bool and assigns it to the Api field.
func (o *TheDesignCenterSchema) SetApi(v bool) {
	o.Api = &v
}

// GetMozart returns the Mozart field value if set, zero value otherwise.
func (o *TheDesignCenterSchema) GetMozart() bool {
	if o == nil || o.Mozart == nil {
		var ret bool
		return ret
	}
	return *o.Mozart
}

// GetMozartOk returns a tuple with the Mozart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TheDesignCenterSchema) GetMozartOk() (*bool, bool) {
	if o == nil || o.Mozart == nil {
		return nil, false
	}
	return o.Mozart, true
}

// HasMozart returns a boolean if a field has been set.
func (o *TheDesignCenterSchema) HasMozart() bool {
	if o != nil && o.Mozart != nil {
		return true
	}

	return false
}

// SetMozart gets a reference to the given bool and assigns it to the Mozart field.
func (o *TheDesignCenterSchema) SetMozart(v bool) {
	o.Mozart = &v
}

func (o TheDesignCenterSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Api != nil {
		toSerialize["api"] = o.Api
	}
	if o.Mozart != nil {
		toSerialize["mozart"] = o.Mozart
	}
	return json.Marshal(toSerialize)
}

type NullableTheDesignCenterSchema struct {
	value *TheDesignCenterSchema
	isSet bool
}

func (v NullableTheDesignCenterSchema) Get() *TheDesignCenterSchema {
	return v.value
}

func (v *NullableTheDesignCenterSchema) Set(val *TheDesignCenterSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableTheDesignCenterSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableTheDesignCenterSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTheDesignCenterSchema(val *TheDesignCenterSchema) *NullableTheDesignCenterSchema {
	return &NullableTheDesignCenterSchema{value: val, isSet: true}
}

func (v NullableTheDesignCenterSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTheDesignCenterSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


