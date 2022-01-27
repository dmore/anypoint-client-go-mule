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

// IdpPatchBody struct for IdpPatchBody
type IdpPatchBody struct {
	Name *string `json:"name,omitempty"`
	Type *IdpPatchBodyType `json:"type,omitempty"`
	OidcProvider *OidcProvider1 `json:"oidc_provider,omitempty"`
	AllowUntrustedCertificates *bool `json:"allow_untrusted_certificates,omitempty"`
	Saml *Saml1 `json:"saml,omitempty"`
	ServiceProvider *ServiceProvider1 `json:"service_provider,omitempty"`
}

// NewIdpPatchBody instantiates a new IdpPatchBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpPatchBody() *IdpPatchBody {
	this := IdpPatchBody{}
	return &this
}

// NewIdpPatchBodyWithDefaults instantiates a new IdpPatchBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpPatchBodyWithDefaults() *IdpPatchBody {
	this := IdpPatchBody{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdpPatchBody) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPatchBody) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdpPatchBody) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdpPatchBody) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdpPatchBody) GetType() IdpPatchBodyType {
	if o == nil || o.Type == nil {
		var ret IdpPatchBodyType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPatchBody) GetTypeOk() (*IdpPatchBodyType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdpPatchBody) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given IdpPatchBodyType and assigns it to the Type field.
func (o *IdpPatchBody) SetType(v IdpPatchBodyType) {
	o.Type = &v
}

// GetOidcProvider returns the OidcProvider field value if set, zero value otherwise.
func (o *IdpPatchBody) GetOidcProvider() OidcProvider1 {
	if o == nil || o.OidcProvider == nil {
		var ret OidcProvider1
		return ret
	}
	return *o.OidcProvider
}

// GetOidcProviderOk returns a tuple with the OidcProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPatchBody) GetOidcProviderOk() (*OidcProvider1, bool) {
	if o == nil || o.OidcProvider == nil {
		return nil, false
	}
	return o.OidcProvider, true
}

// HasOidcProvider returns a boolean if a field has been set.
func (o *IdpPatchBody) HasOidcProvider() bool {
	if o != nil && o.OidcProvider != nil {
		return true
	}

	return false
}

// SetOidcProvider gets a reference to the given OidcProvider1 and assigns it to the OidcProvider field.
func (o *IdpPatchBody) SetOidcProvider(v OidcProvider1) {
	o.OidcProvider = &v
}

// GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field value if set, zero value otherwise.
func (o *IdpPatchBody) GetAllowUntrustedCertificates() bool {
	if o == nil || o.AllowUntrustedCertificates == nil {
		var ret bool
		return ret
	}
	return *o.AllowUntrustedCertificates
}

// GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPatchBody) GetAllowUntrustedCertificatesOk() (*bool, bool) {
	if o == nil || o.AllowUntrustedCertificates == nil {
		return nil, false
	}
	return o.AllowUntrustedCertificates, true
}

// HasAllowUntrustedCertificates returns a boolean if a field has been set.
func (o *IdpPatchBody) HasAllowUntrustedCertificates() bool {
	if o != nil && o.AllowUntrustedCertificates != nil {
		return true
	}

	return false
}

// SetAllowUntrustedCertificates gets a reference to the given bool and assigns it to the AllowUntrustedCertificates field.
func (o *IdpPatchBody) SetAllowUntrustedCertificates(v bool) {
	o.AllowUntrustedCertificates = &v
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *IdpPatchBody) GetSaml() Saml1 {
	if o == nil || o.Saml == nil {
		var ret Saml1
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPatchBody) GetSamlOk() (*Saml1, bool) {
	if o == nil || o.Saml == nil {
		return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *IdpPatchBody) HasSaml() bool {
	if o != nil && o.Saml != nil {
		return true
	}

	return false
}

// SetSaml gets a reference to the given Saml1 and assigns it to the Saml field.
func (o *IdpPatchBody) SetSaml(v Saml1) {
	o.Saml = &v
}

// GetServiceProvider returns the ServiceProvider field value if set, zero value otherwise.
func (o *IdpPatchBody) GetServiceProvider() ServiceProvider1 {
	if o == nil || o.ServiceProvider == nil {
		var ret ServiceProvider1
		return ret
	}
	return *o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPatchBody) GetServiceProviderOk() (*ServiceProvider1, bool) {
	if o == nil || o.ServiceProvider == nil {
		return nil, false
	}
	return o.ServiceProvider, true
}

// HasServiceProvider returns a boolean if a field has been set.
func (o *IdpPatchBody) HasServiceProvider() bool {
	if o != nil && o.ServiceProvider != nil {
		return true
	}

	return false
}

// SetServiceProvider gets a reference to the given ServiceProvider1 and assigns it to the ServiceProvider field.
func (o *IdpPatchBody) SetServiceProvider(v ServiceProvider1) {
	o.ServiceProvider = &v
}

func (o IdpPatchBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.OidcProvider != nil {
		toSerialize["oidc_provider"] = o.OidcProvider
	}
	if o.AllowUntrustedCertificates != nil {
		toSerialize["allow_untrusted_certificates"] = o.AllowUntrustedCertificates
	}
	if o.Saml != nil {
		toSerialize["saml"] = o.Saml
	}
	if o.ServiceProvider != nil {
		toSerialize["service_provider"] = o.ServiceProvider
	}
	return json.Marshal(toSerialize)
}

type NullableIdpPatchBody struct {
	value *IdpPatchBody
	isSet bool
}

func (v NullableIdpPatchBody) Get() *IdpPatchBody {
	return v.value
}

func (v *NullableIdpPatchBody) Set(val *IdpPatchBody) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpPatchBody) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpPatchBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpPatchBody(val *IdpPatchBody) *NullableIdpPatchBody {
	return &NullableIdpPatchBody{value: val, isSet: true}
}

func (v NullableIdpPatchBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpPatchBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


