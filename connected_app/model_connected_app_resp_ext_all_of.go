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

// ConnectedAppRespExtAllOf struct for ConnectedAppRespExtAllOf
type ConnectedAppRespExtAllOf struct {
	// connected app client id
	ClientId *string `json:"client_id,omitempty"`
	// connected app generated secret
	ClientSecret *string `json:"client_secret,omitempty"`
	// connected app owner organization id
	OwnerOrgId *string `json:"owner_org_id,omitempty"`
	// connected app owner user id
	OwnerUserId *string `json:"owner_user_id,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	PolicyUri *string `json:"policy_uri,omitempty"`
	TosUri *string `json:"tos_uri,omitempty"`
	CertExpiry *string `json:"cert_expiry,omitempty"`
}

// NewConnectedAppRespExtAllOf instantiates a new ConnectedAppRespExtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedAppRespExtAllOf() *ConnectedAppRespExtAllOf {
	this := ConnectedAppRespExtAllOf{}
	return &this
}

// NewConnectedAppRespExtAllOfWithDefaults instantiates a new ConnectedAppRespExtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedAppRespExtAllOfWithDefaults() *ConnectedAppRespExtAllOf {
	this := ConnectedAppRespExtAllOf{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ConnectedAppRespExtAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExtAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ConnectedAppRespExtAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ConnectedAppRespExtAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ConnectedAppRespExtAllOf) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExtAllOf) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ConnectedAppRespExtAllOf) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ConnectedAppRespExtAllOf) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetOwnerOrgId returns the OwnerOrgId field value if set, zero value otherwise.
func (o *ConnectedAppRespExtAllOf) GetOwnerOrgId() string {
	if o == nil || o.OwnerOrgId == nil {
		var ret string
		return ret
	}
	return *o.OwnerOrgId
}

// GetOwnerOrgIdOk returns a tuple with the OwnerOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExtAllOf) GetOwnerOrgIdOk() (*string, bool) {
	if o == nil || o.OwnerOrgId == nil {
		return nil, false
	}
	return o.OwnerOrgId, true
}

// HasOwnerOrgId returns a boolean if a field has been set.
func (o *ConnectedAppRespExtAllOf) HasOwnerOrgId() bool {
	if o != nil && o.OwnerOrgId != nil {
		return true
	}

	return false
}

// SetOwnerOrgId gets a reference to the given string and assigns it to the OwnerOrgId field.
func (o *ConnectedAppRespExtAllOf) SetOwnerOrgId(v string) {
	o.OwnerOrgId = &v
}

// GetOwnerUserId returns the OwnerUserId field value if set, zero value otherwise.
func (o *ConnectedAppRespExtAllOf) GetOwnerUserId() string {
	if o == nil || o.OwnerUserId == nil {
		var ret string
		return ret
	}
	return *o.OwnerUserId
}

// GetOwnerUserIdOk returns a tuple with the OwnerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExtAllOf) GetOwnerUserIdOk() (*string, bool) {
	if o == nil || o.OwnerUserId == nil {
		return nil, false
	}
	return o.OwnerUserId, true
}

// HasOwnerUserId returns a boolean if a field has been set.
func (o *ConnectedAppRespExtAllOf) HasOwnerUserId() bool {
	if o != nil && o.OwnerUserId != nil {
		return true
	}

	return false
}

// SetOwnerUserId gets a reference to the given string and assigns it to the OwnerUserId field.
func (o *ConnectedAppRespExtAllOf) SetOwnerUserId(v string) {
	o.OwnerUserId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ConnectedAppRespExtAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExtAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ConnectedAppRespExtAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ConnectedAppRespExtAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPolicyUri returns the PolicyUri field value if set, zero value otherwise.
func (o *ConnectedAppRespExtAllOf) GetPolicyUri() string {
	if o == nil || o.PolicyUri == nil {
		var ret string
		return ret
	}
	return *o.PolicyUri
}

// GetPolicyUriOk returns a tuple with the PolicyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExtAllOf) GetPolicyUriOk() (*string, bool) {
	if o == nil || o.PolicyUri == nil {
		return nil, false
	}
	return o.PolicyUri, true
}

// HasPolicyUri returns a boolean if a field has been set.
func (o *ConnectedAppRespExtAllOf) HasPolicyUri() bool {
	if o != nil && o.PolicyUri != nil {
		return true
	}

	return false
}

// SetPolicyUri gets a reference to the given string and assigns it to the PolicyUri field.
func (o *ConnectedAppRespExtAllOf) SetPolicyUri(v string) {
	o.PolicyUri = &v
}

// GetTosUri returns the TosUri field value if set, zero value otherwise.
func (o *ConnectedAppRespExtAllOf) GetTosUri() string {
	if o == nil || o.TosUri == nil {
		var ret string
		return ret
	}
	return *o.TosUri
}

// GetTosUriOk returns a tuple with the TosUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExtAllOf) GetTosUriOk() (*string, bool) {
	if o == nil || o.TosUri == nil {
		return nil, false
	}
	return o.TosUri, true
}

// HasTosUri returns a boolean if a field has been set.
func (o *ConnectedAppRespExtAllOf) HasTosUri() bool {
	if o != nil && o.TosUri != nil {
		return true
	}

	return false
}

// SetTosUri gets a reference to the given string and assigns it to the TosUri field.
func (o *ConnectedAppRespExtAllOf) SetTosUri(v string) {
	o.TosUri = &v
}

// GetCertExpiry returns the CertExpiry field value if set, zero value otherwise.
func (o *ConnectedAppRespExtAllOf) GetCertExpiry() string {
	if o == nil || o.CertExpiry == nil {
		var ret string
		return ret
	}
	return *o.CertExpiry
}

// GetCertExpiryOk returns a tuple with the CertExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExtAllOf) GetCertExpiryOk() (*string, bool) {
	if o == nil || o.CertExpiry == nil {
		return nil, false
	}
	return o.CertExpiry, true
}

// HasCertExpiry returns a boolean if a field has been set.
func (o *ConnectedAppRespExtAllOf) HasCertExpiry() bool {
	if o != nil && o.CertExpiry != nil {
		return true
	}

	return false
}

// SetCertExpiry gets a reference to the given string and assigns it to the CertExpiry field.
func (o *ConnectedAppRespExtAllOf) SetCertExpiry(v string) {
	o.CertExpiry = &v
}

func (o ConnectedAppRespExtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.OwnerOrgId != nil {
		toSerialize["owner_org_id"] = o.OwnerOrgId
	}
	if o.OwnerUserId != nil {
		toSerialize["owner_user_id"] = o.OwnerUserId
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.PolicyUri != nil {
		toSerialize["policy_uri"] = o.PolicyUri
	}
	if o.TosUri != nil {
		toSerialize["tos_uri"] = o.TosUri
	}
	if o.CertExpiry != nil {
		toSerialize["cert_expiry"] = o.CertExpiry
	}
	return json.Marshal(toSerialize)
}

type NullableConnectedAppRespExtAllOf struct {
	value *ConnectedAppRespExtAllOf
	isSet bool
}

func (v NullableConnectedAppRespExtAllOf) Get() *ConnectedAppRespExtAllOf {
	return v.value
}

func (v *NullableConnectedAppRespExtAllOf) Set(val *ConnectedAppRespExtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedAppRespExtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedAppRespExtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedAppRespExtAllOf(val *ConnectedAppRespExtAllOf) *NullableConnectedAppRespExtAllOf {
	return &NullableConnectedAppRespExtAllOf{value: val, isSet: true}
}

func (v NullableConnectedAppRespExtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedAppRespExtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


