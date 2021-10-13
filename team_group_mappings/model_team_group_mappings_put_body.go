/*
 * Team Group Mappings API
 *
 * Description of the Team Group Mappings API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package team_group_mappings

import (
	"encoding/json"
)

// TeamGroupMappingsPutBody struct for TeamGroupMappingsPutBody
type TeamGroupMappingsPutBody struct {
	MembershipType *string `json:"membership_type,omitempty"`
	ExternalGroupName *string `json:"external_group_name,omitempty"`
}

// NewTeamGroupMappingsPutBody instantiates a new TeamGroupMappingsPutBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamGroupMappingsPutBody() *TeamGroupMappingsPutBody {
	this := TeamGroupMappingsPutBody{}
	return &this
}

// NewTeamGroupMappingsPutBodyWithDefaults instantiates a new TeamGroupMappingsPutBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamGroupMappingsPutBodyWithDefaults() *TeamGroupMappingsPutBody {
	this := TeamGroupMappingsPutBody{}
	return &this
}

// GetMembershipType returns the MembershipType field value if set, zero value otherwise.
func (o *TeamGroupMappingsPutBody) GetMembershipType() string {
	if o == nil || o.MembershipType == nil {
		var ret string
		return ret
	}
	return *o.MembershipType
}

// GetMembershipTypeOk returns a tuple with the MembershipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamGroupMappingsPutBody) GetMembershipTypeOk() (*string, bool) {
	if o == nil || o.MembershipType == nil {
		return nil, false
	}
	return o.MembershipType, true
}

// HasMembershipType returns a boolean if a field has been set.
func (o *TeamGroupMappingsPutBody) HasMembershipType() bool {
	if o != nil && o.MembershipType != nil {
		return true
	}

	return false
}

// SetMembershipType gets a reference to the given string and assigns it to the MembershipType field.
func (o *TeamGroupMappingsPutBody) SetMembershipType(v string) {
	o.MembershipType = &v
}

// GetExternalGroupName returns the ExternalGroupName field value if set, zero value otherwise.
func (o *TeamGroupMappingsPutBody) GetExternalGroupName() string {
	if o == nil || o.ExternalGroupName == nil {
		var ret string
		return ret
	}
	return *o.ExternalGroupName
}

// GetExternalGroupNameOk returns a tuple with the ExternalGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamGroupMappingsPutBody) GetExternalGroupNameOk() (*string, bool) {
	if o == nil || o.ExternalGroupName == nil {
		return nil, false
	}
	return o.ExternalGroupName, true
}

// HasExternalGroupName returns a boolean if a field has been set.
func (o *TeamGroupMappingsPutBody) HasExternalGroupName() bool {
	if o != nil && o.ExternalGroupName != nil {
		return true
	}

	return false
}

// SetExternalGroupName gets a reference to the given string and assigns it to the ExternalGroupName field.
func (o *TeamGroupMappingsPutBody) SetExternalGroupName(v string) {
	o.ExternalGroupName = &v
}

func (o TeamGroupMappingsPutBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MembershipType != nil {
		toSerialize["membership_type"] = o.MembershipType
	}
	if o.ExternalGroupName != nil {
		toSerialize["external_group_name"] = o.ExternalGroupName
	}
	return json.Marshal(toSerialize)
}

type NullableTeamGroupMappingsPutBody struct {
	value *TeamGroupMappingsPutBody
	isSet bool
}

func (v NullableTeamGroupMappingsPutBody) Get() *TeamGroupMappingsPutBody {
	return v.value
}

func (v *NullableTeamGroupMappingsPutBody) Set(val *TeamGroupMappingsPutBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamGroupMappingsPutBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamGroupMappingsPutBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamGroupMappingsPutBody(val *TeamGroupMappingsPutBody) *NullableTeamGroupMappingsPutBody {
	return &NullableTeamGroupMappingsPutBody{value: val, isSet: true}
}

func (v NullableTeamGroupMappingsPutBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamGroupMappingsPutBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

