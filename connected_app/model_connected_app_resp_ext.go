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

// ConnectedAppRespExt struct for ConnectedAppRespExt
type ConnectedAppRespExt struct {
	ClientName *string `json:"client_name,omitempty"`
	GrantTypes *[]string `json:"grant_types,omitempty"`
	RedirectUris *[]string `json:"redirect_uris,omitempty"`
	Scopes *[]string `json:"scopes,omitempty"`
	// Application public key (PEM format). Used to validate JWT authorization grants.
	PublicKeys *[]string `json:"public_keys,omitempty"`
	ClientUri *string `json:"client_uri,omitempty"`
	Audience *string `json:"audience,omitempty"`
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

// NewConnectedAppRespExt instantiates a new ConnectedAppRespExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedAppRespExt() *ConnectedAppRespExt {
	this := ConnectedAppRespExt{}
	return &this
}

// NewConnectedAppRespExtWithDefaults instantiates a new ConnectedAppRespExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedAppRespExtWithDefaults() *ConnectedAppRespExt {
	this := ConnectedAppRespExt{}
	return &this
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *ConnectedAppRespExt) SetClientName(v string) {
	o.ClientName = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetGrantTypes() []string {
	if o == nil || o.GrantTypes == nil {
		var ret []string
		return ret
	}
	return *o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetGrantTypesOk() (*[]string, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *ConnectedAppRespExt) SetGrantTypes(v []string) {
	o.GrantTypes = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *ConnectedAppRespExt) SetRedirectUris(v []string) {
	o.RedirectUris = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ConnectedAppRespExt) SetScopes(v []string) {
	o.Scopes = &v
}

// GetPublicKeys returns the PublicKeys field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetPublicKeys() []string {
	if o == nil || o.PublicKeys == nil {
		var ret []string
		return ret
	}
	return *o.PublicKeys
}

// GetPublicKeysOk returns a tuple with the PublicKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetPublicKeysOk() (*[]string, bool) {
	if o == nil || o.PublicKeys == nil {
		return nil, false
	}
	return o.PublicKeys, true
}

// HasPublicKeys returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasPublicKeys() bool {
	if o != nil && o.PublicKeys != nil {
		return true
	}

	return false
}

// SetPublicKeys gets a reference to the given []string and assigns it to the PublicKeys field.
func (o *ConnectedAppRespExt) SetPublicKeys(v []string) {
	o.PublicKeys = &v
}

// GetClientUri returns the ClientUri field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetClientUri() string {
	if o == nil || o.ClientUri == nil {
		var ret string
		return ret
	}
	return *o.ClientUri
}

// GetClientUriOk returns a tuple with the ClientUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetClientUriOk() (*string, bool) {
	if o == nil || o.ClientUri == nil {
		return nil, false
	}
	return o.ClientUri, true
}

// HasClientUri returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasClientUri() bool {
	if o != nil && o.ClientUri != nil {
		return true
	}

	return false
}

// SetClientUri gets a reference to the given string and assigns it to the ClientUri field.
func (o *ConnectedAppRespExt) SetClientUri(v string) {
	o.ClientUri = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *ConnectedAppRespExt) SetAudience(v string) {
	o.Audience = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ConnectedAppRespExt) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ConnectedAppRespExt) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetOwnerOrgId returns the OwnerOrgId field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetOwnerOrgId() string {
	if o == nil || o.OwnerOrgId == nil {
		var ret string
		return ret
	}
	return *o.OwnerOrgId
}

// GetOwnerOrgIdOk returns a tuple with the OwnerOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetOwnerOrgIdOk() (*string, bool) {
	if o == nil || o.OwnerOrgId == nil {
		return nil, false
	}
	return o.OwnerOrgId, true
}

// HasOwnerOrgId returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasOwnerOrgId() bool {
	if o != nil && o.OwnerOrgId != nil {
		return true
	}

	return false
}

// SetOwnerOrgId gets a reference to the given string and assigns it to the OwnerOrgId field.
func (o *ConnectedAppRespExt) SetOwnerOrgId(v string) {
	o.OwnerOrgId = &v
}

// GetOwnerUserId returns the OwnerUserId field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetOwnerUserId() string {
	if o == nil || o.OwnerUserId == nil {
		var ret string
		return ret
	}
	return *o.OwnerUserId
}

// GetOwnerUserIdOk returns a tuple with the OwnerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetOwnerUserIdOk() (*string, bool) {
	if o == nil || o.OwnerUserId == nil {
		return nil, false
	}
	return o.OwnerUserId, true
}

// HasOwnerUserId returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasOwnerUserId() bool {
	if o != nil && o.OwnerUserId != nil {
		return true
	}

	return false
}

// SetOwnerUserId gets a reference to the given string and assigns it to the OwnerUserId field.
func (o *ConnectedAppRespExt) SetOwnerUserId(v string) {
	o.OwnerUserId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ConnectedAppRespExt) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPolicyUri returns the PolicyUri field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetPolicyUri() string {
	if o == nil || o.PolicyUri == nil {
		var ret string
		return ret
	}
	return *o.PolicyUri
}

// GetPolicyUriOk returns a tuple with the PolicyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetPolicyUriOk() (*string, bool) {
	if o == nil || o.PolicyUri == nil {
		return nil, false
	}
	return o.PolicyUri, true
}

// HasPolicyUri returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasPolicyUri() bool {
	if o != nil && o.PolicyUri != nil {
		return true
	}

	return false
}

// SetPolicyUri gets a reference to the given string and assigns it to the PolicyUri field.
func (o *ConnectedAppRespExt) SetPolicyUri(v string) {
	o.PolicyUri = &v
}

// GetTosUri returns the TosUri field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetTosUri() string {
	if o == nil || o.TosUri == nil {
		var ret string
		return ret
	}
	return *o.TosUri
}

// GetTosUriOk returns a tuple with the TosUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetTosUriOk() (*string, bool) {
	if o == nil || o.TosUri == nil {
		return nil, false
	}
	return o.TosUri, true
}

// HasTosUri returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasTosUri() bool {
	if o != nil && o.TosUri != nil {
		return true
	}

	return false
}

// SetTosUri gets a reference to the given string and assigns it to the TosUri field.
func (o *ConnectedAppRespExt) SetTosUri(v string) {
	o.TosUri = &v
}

// GetCertExpiry returns the CertExpiry field value if set, zero value otherwise.
func (o *ConnectedAppRespExt) GetCertExpiry() string {
	if o == nil || o.CertExpiry == nil {
		var ret string
		return ret
	}
	return *o.CertExpiry
}

// GetCertExpiryOk returns a tuple with the CertExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppRespExt) GetCertExpiryOk() (*string, bool) {
	if o == nil || o.CertExpiry == nil {
		return nil, false
	}
	return o.CertExpiry, true
}

// HasCertExpiry returns a boolean if a field has been set.
func (o *ConnectedAppRespExt) HasCertExpiry() bool {
	if o != nil && o.CertExpiry != nil {
		return true
	}

	return false
}

// SetCertExpiry gets a reference to the given string and assigns it to the CertExpiry field.
func (o *ConnectedAppRespExt) SetCertExpiry(v string) {
	o.CertExpiry = &v
}

func (o ConnectedAppRespExt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientName != nil {
		toSerialize["client_name"] = o.ClientName
	}
	if o.GrantTypes != nil {
		toSerialize["grant_types"] = o.GrantTypes
	}
	if o.RedirectUris != nil {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.PublicKeys != nil {
		toSerialize["public_keys"] = o.PublicKeys
	}
	if o.ClientUri != nil {
		toSerialize["client_uri"] = o.ClientUri
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
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

type NullableConnectedAppRespExt struct {
	value *ConnectedAppRespExt
	isSet bool
}

func (v NullableConnectedAppRespExt) Get() *ConnectedAppRespExt {
	return v.value
}

func (v *NullableConnectedAppRespExt) Set(val *ConnectedAppRespExt) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedAppRespExt) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedAppRespExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedAppRespExt(val *ConnectedAppRespExt) *NullableConnectedAppRespExt {
	return &NullableConnectedAppRespExt{value: val, isSet: true}
}

func (v NullableConnectedAppRespExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedAppRespExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


