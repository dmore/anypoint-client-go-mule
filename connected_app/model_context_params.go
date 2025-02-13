/*
 * Connected App API
 *
 * Description of the Connected App API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connected_app

import (
	"encoding/json"
)

// ContextParams struct for ContextParams
type ContextParams struct {
	Org *string `json:"org,omitempty"`
	EnvId *string `json:"envId,omitempty"`
}

// NewContextParams instantiates a new ContextParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextParams() *ContextParams {
	this := ContextParams{}
	return &this
}

// NewContextParamsWithDefaults instantiates a new ContextParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextParamsWithDefaults() *ContextParams {
	this := ContextParams{}
	return &this
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *ContextParams) GetOrg() string {
	if o == nil || o.Org == nil {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextParams) GetOrgOk() (*string, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *ContextParams) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *ContextParams) SetOrg(v string) {
	o.Org = &v
}

// GetEnvId returns the EnvId field value if set, zero value otherwise.
func (o *ContextParams) GetEnvId() string {
	if o == nil || o.EnvId == nil {
		var ret string
		return ret
	}
	return *o.EnvId
}

// GetEnvIdOk returns a tuple with the EnvId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextParams) GetEnvIdOk() (*string, bool) {
	if o == nil || o.EnvId == nil {
		return nil, false
	}
	return o.EnvId, true
}

// HasEnvId returns a boolean if a field has been set.
func (o *ContextParams) HasEnvId() bool {
	if o != nil && o.EnvId != nil {
		return true
	}

	return false
}

// SetEnvId gets a reference to the given string and assigns it to the EnvId field.
func (o *ContextParams) SetEnvId(v string) {
	o.EnvId = &v
}

func (o ContextParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	if o.EnvId != nil {
		toSerialize["envId"] = o.EnvId
	}
	return json.Marshal(toSerialize)
}

type NullableContextParams struct {
	value *ContextParams
	isSet bool
}

func (v NullableContextParams) Get() *ContextParams {
	return v.value
}

func (v *NullableContextParams) Set(val *ContextParams) {
	v.value = val
	v.isSet = true
}

func (v NullableContextParams) IsSet() bool {
	return v.isSet
}

func (v *NullableContextParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextParams(val *ContextParams) *NullableContextParams {
	return &NullableContextParams{value: val, isSet: true}
}

func (v NullableContextParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


