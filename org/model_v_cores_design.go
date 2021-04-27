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

// VCoresDesign An explanation about the purpose of this instance.
type VCoresDesign struct {
	// An explanation about the purpose of this instance.
	Assigned int32 `json:"assigned"`
	// An explanation about the purpose of this instance.
	Reassigned float32 `json:"reassigned"`
}

// NewVCoresDesign instantiates a new VCoresDesign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVCoresDesign(assigned int32, reassigned float32) *VCoresDesign {
	this := VCoresDesign{}
	this.Assigned = assigned
	this.Reassigned = reassigned
	return &this
}

// NewVCoresDesignWithDefaults instantiates a new VCoresDesign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVCoresDesignWithDefaults() *VCoresDesign {
	this := VCoresDesign{}
	var assigned int32 = 0
	this.Assigned = assigned
	var reassigned float32 = 0.0
	this.Reassigned = reassigned
	return &this
}

// GetAssigned returns the Assigned field value
func (o *VCoresDesign) GetAssigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value
// and a boolean to check if the value has been set.
func (o *VCoresDesign) GetAssignedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Assigned, true
}

// SetAssigned sets field value
func (o *VCoresDesign) SetAssigned(v int32) {
	o.Assigned = v
}

// GetReassigned returns the Reassigned field value
func (o *VCoresDesign) GetReassigned() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Reassigned
}

// GetReassignedOk returns a tuple with the Reassigned field value
// and a boolean to check if the value has been set.
func (o *VCoresDesign) GetReassignedOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reassigned, true
}

// SetReassigned sets field value
func (o *VCoresDesign) SetReassigned(v float32) {
	o.Reassigned = v
}

func (o VCoresDesign) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assigned"] = o.Assigned
	}
	if true {
		toSerialize["reassigned"] = o.Reassigned
	}
	return json.Marshal(toSerialize)
}

type NullableVCoresDesign struct {
	value *VCoresDesign
	isSet bool
}

func (v NullableVCoresDesign) Get() *VCoresDesign {
	return v.value
}

func (v *NullableVCoresDesign) Set(val *VCoresDesign) {
	v.value = val
	v.isSet = true
}

func (v NullableVCoresDesign) IsSet() bool {
	return v.isSet
}

func (v *NullableVCoresDesign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVCoresDesign(val *VCoresDesign) *NullableVCoresDesign {
	return &NullableVCoresDesign{value: val, isSet: true}
}

func (v NullableVCoresDesign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVCoresDesign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


