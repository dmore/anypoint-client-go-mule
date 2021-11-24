/*
 * Dedicated Load Balancer API
 *
 * Description of the DLB API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dlb

import (
	"encoding/json"
)

// Dlb struct for Dlb
type Dlb struct {
	// The dlb Id
	Id *string `json:"id,omitempty"`
	// dlb domain
	Domain *string `json:"domain,omitempty"`
	// the dlb deployment id
	DeploymentId *string `json:"deploymentId,omitempty"`
	InstanceConfig *InstanceConfig `json:"instanceConfig,omitempty"`
	IpAddresses *[]string `json:"ipAddresses,omitempty"`
	StaticIPsDisabled *bool `json:"staticIPsDisabled,omitempty"`
	VpcId *string `json:"vpcId,omitempty"`
	Workers *int32 `json:"workers,omitempty"`
	DefaultCipherSuite *string `json:"defaultCipherSuite,omitempty"`
	KeepUrlEncoding *bool `json:"keepUrlEncoding,omitempty"`
	UpstreamTlsv12 *bool `json:"upstreamTlsv12,omitempty"`
	ProxyReadTimeout *int32 `json:"proxyReadTimeout,omitempty"`
	IpAddressesInfo *[]DlbExtrasIpAddressesInfo `json:"ipAddressesInfo,omitempty"`
	DoubleStaticIps *string `json:"doubleStaticIps,omitempty"`
	Name *string `json:"name,omitempty"`
	// dlb state
	State *string `json:"state,omitempty"`
	IpWhitelist *[]string `json:"ipWhitelist,omitempty"`
	HttpMode *string `json:"httpMode,omitempty"`
	DefaultSslEndpoint *int32 `json:"defaultSslEndpoint,omitempty"`
	Tlsv1 *bool `json:"tlsv1,omitempty"`
	SslEndpoints *[]DlbCoreSslEndpoints `json:"sslEndpoints,omitempty"`
}

// NewDlb instantiates a new Dlb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDlb() *Dlb {
	this := Dlb{}
	var state string = "STOPPED"
	this.State = &state
	var httpMode string = "redirect"
	this.HttpMode = &httpMode
	var defaultSslEndpoint int32 = 0
	this.DefaultSslEndpoint = &defaultSslEndpoint
	return &this
}

// NewDlbWithDefaults instantiates a new Dlb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDlbWithDefaults() *Dlb {
	this := Dlb{}
	var state string = "STOPPED"
	this.State = &state
	var httpMode string = "redirect"
	this.HttpMode = &httpMode
	var defaultSslEndpoint int32 = 0
	this.DefaultSslEndpoint = &defaultSslEndpoint
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dlb) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dlb) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Dlb) SetId(v string) {
	o.Id = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Dlb) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Dlb) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Dlb) SetDomain(v string) {
	o.Domain = &v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *Dlb) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *Dlb) HasDeploymentId() bool {
	if o != nil && o.DeploymentId != nil {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *Dlb) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetInstanceConfig returns the InstanceConfig field value if set, zero value otherwise.
func (o *Dlb) GetInstanceConfig() InstanceConfig {
	if o == nil || o.InstanceConfig == nil {
		var ret InstanceConfig
		return ret
	}
	return *o.InstanceConfig
}

// GetInstanceConfigOk returns a tuple with the InstanceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetInstanceConfigOk() (*InstanceConfig, bool) {
	if o == nil || o.InstanceConfig == nil {
		return nil, false
	}
	return o.InstanceConfig, true
}

// HasInstanceConfig returns a boolean if a field has been set.
func (o *Dlb) HasInstanceConfig() bool {
	if o != nil && o.InstanceConfig != nil {
		return true
	}

	return false
}

// SetInstanceConfig gets a reference to the given InstanceConfig and assigns it to the InstanceConfig field.
func (o *Dlb) SetInstanceConfig(v InstanceConfig) {
	o.InstanceConfig = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *Dlb) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return *o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *Dlb) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *Dlb) SetIpAddresses(v []string) {
	o.IpAddresses = &v
}

// GetStaticIPsDisabled returns the StaticIPsDisabled field value if set, zero value otherwise.
func (o *Dlb) GetStaticIPsDisabled() bool {
	if o == nil || o.StaticIPsDisabled == nil {
		var ret bool
		return ret
	}
	return *o.StaticIPsDisabled
}

// GetStaticIPsDisabledOk returns a tuple with the StaticIPsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetStaticIPsDisabledOk() (*bool, bool) {
	if o == nil || o.StaticIPsDisabled == nil {
		return nil, false
	}
	return o.StaticIPsDisabled, true
}

// HasStaticIPsDisabled returns a boolean if a field has been set.
func (o *Dlb) HasStaticIPsDisabled() bool {
	if o != nil && o.StaticIPsDisabled != nil {
		return true
	}

	return false
}

// SetStaticIPsDisabled gets a reference to the given bool and assigns it to the StaticIPsDisabled field.
func (o *Dlb) SetStaticIPsDisabled(v bool) {
	o.StaticIPsDisabled = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *Dlb) GetVpcId() string {
	if o == nil || o.VpcId == nil {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetVpcIdOk() (*string, bool) {
	if o == nil || o.VpcId == nil {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *Dlb) HasVpcId() bool {
	if o != nil && o.VpcId != nil {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *Dlb) SetVpcId(v string) {
	o.VpcId = &v
}

// GetWorkers returns the Workers field value if set, zero value otherwise.
func (o *Dlb) GetWorkers() int32 {
	if o == nil || o.Workers == nil {
		var ret int32
		return ret
	}
	return *o.Workers
}

// GetWorkersOk returns a tuple with the Workers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetWorkersOk() (*int32, bool) {
	if o == nil || o.Workers == nil {
		return nil, false
	}
	return o.Workers, true
}

// HasWorkers returns a boolean if a field has been set.
func (o *Dlb) HasWorkers() bool {
	if o != nil && o.Workers != nil {
		return true
	}

	return false
}

// SetWorkers gets a reference to the given int32 and assigns it to the Workers field.
func (o *Dlb) SetWorkers(v int32) {
	o.Workers = &v
}

// GetDefaultCipherSuite returns the DefaultCipherSuite field value if set, zero value otherwise.
func (o *Dlb) GetDefaultCipherSuite() string {
	if o == nil || o.DefaultCipherSuite == nil {
		var ret string
		return ret
	}
	return *o.DefaultCipherSuite
}

// GetDefaultCipherSuiteOk returns a tuple with the DefaultCipherSuite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetDefaultCipherSuiteOk() (*string, bool) {
	if o == nil || o.DefaultCipherSuite == nil {
		return nil, false
	}
	return o.DefaultCipherSuite, true
}

// HasDefaultCipherSuite returns a boolean if a field has been set.
func (o *Dlb) HasDefaultCipherSuite() bool {
	if o != nil && o.DefaultCipherSuite != nil {
		return true
	}

	return false
}

// SetDefaultCipherSuite gets a reference to the given string and assigns it to the DefaultCipherSuite field.
func (o *Dlb) SetDefaultCipherSuite(v string) {
	o.DefaultCipherSuite = &v
}

// GetKeepUrlEncoding returns the KeepUrlEncoding field value if set, zero value otherwise.
func (o *Dlb) GetKeepUrlEncoding() bool {
	if o == nil || o.KeepUrlEncoding == nil {
		var ret bool
		return ret
	}
	return *o.KeepUrlEncoding
}

// GetKeepUrlEncodingOk returns a tuple with the KeepUrlEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetKeepUrlEncodingOk() (*bool, bool) {
	if o == nil || o.KeepUrlEncoding == nil {
		return nil, false
	}
	return o.KeepUrlEncoding, true
}

// HasKeepUrlEncoding returns a boolean if a field has been set.
func (o *Dlb) HasKeepUrlEncoding() bool {
	if o != nil && o.KeepUrlEncoding != nil {
		return true
	}

	return false
}

// SetKeepUrlEncoding gets a reference to the given bool and assigns it to the KeepUrlEncoding field.
func (o *Dlb) SetKeepUrlEncoding(v bool) {
	o.KeepUrlEncoding = &v
}

// GetUpstreamTlsv12 returns the UpstreamTlsv12 field value if set, zero value otherwise.
func (o *Dlb) GetUpstreamTlsv12() bool {
	if o == nil || o.UpstreamTlsv12 == nil {
		var ret bool
		return ret
	}
	return *o.UpstreamTlsv12
}

// GetUpstreamTlsv12Ok returns a tuple with the UpstreamTlsv12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetUpstreamTlsv12Ok() (*bool, bool) {
	if o == nil || o.UpstreamTlsv12 == nil {
		return nil, false
	}
	return o.UpstreamTlsv12, true
}

// HasUpstreamTlsv12 returns a boolean if a field has been set.
func (o *Dlb) HasUpstreamTlsv12() bool {
	if o != nil && o.UpstreamTlsv12 != nil {
		return true
	}

	return false
}

// SetUpstreamTlsv12 gets a reference to the given bool and assigns it to the UpstreamTlsv12 field.
func (o *Dlb) SetUpstreamTlsv12(v bool) {
	o.UpstreamTlsv12 = &v
}

// GetProxyReadTimeout returns the ProxyReadTimeout field value if set, zero value otherwise.
func (o *Dlb) GetProxyReadTimeout() int32 {
	if o == nil || o.ProxyReadTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ProxyReadTimeout
}

// GetProxyReadTimeoutOk returns a tuple with the ProxyReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetProxyReadTimeoutOk() (*int32, bool) {
	if o == nil || o.ProxyReadTimeout == nil {
		return nil, false
	}
	return o.ProxyReadTimeout, true
}

// HasProxyReadTimeout returns a boolean if a field has been set.
func (o *Dlb) HasProxyReadTimeout() bool {
	if o != nil && o.ProxyReadTimeout != nil {
		return true
	}

	return false
}

// SetProxyReadTimeout gets a reference to the given int32 and assigns it to the ProxyReadTimeout field.
func (o *Dlb) SetProxyReadTimeout(v int32) {
	o.ProxyReadTimeout = &v
}

// GetIpAddressesInfo returns the IpAddressesInfo field value if set, zero value otherwise.
func (o *Dlb) GetIpAddressesInfo() []DlbExtrasIpAddressesInfo {
	if o == nil || o.IpAddressesInfo == nil {
		var ret []DlbExtrasIpAddressesInfo
		return ret
	}
	return *o.IpAddressesInfo
}

// GetIpAddressesInfoOk returns a tuple with the IpAddressesInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetIpAddressesInfoOk() (*[]DlbExtrasIpAddressesInfo, bool) {
	if o == nil || o.IpAddressesInfo == nil {
		return nil, false
	}
	return o.IpAddressesInfo, true
}

// HasIpAddressesInfo returns a boolean if a field has been set.
func (o *Dlb) HasIpAddressesInfo() bool {
	if o != nil && o.IpAddressesInfo != nil {
		return true
	}

	return false
}

// SetIpAddressesInfo gets a reference to the given []DlbExtrasIpAddressesInfo and assigns it to the IpAddressesInfo field.
func (o *Dlb) SetIpAddressesInfo(v []DlbExtrasIpAddressesInfo) {
	o.IpAddressesInfo = &v
}

// GetDoubleStaticIps returns the DoubleStaticIps field value if set, zero value otherwise.
func (o *Dlb) GetDoubleStaticIps() string {
	if o == nil || o.DoubleStaticIps == nil {
		var ret string
		return ret
	}
	return *o.DoubleStaticIps
}

// GetDoubleStaticIpsOk returns a tuple with the DoubleStaticIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetDoubleStaticIpsOk() (*string, bool) {
	if o == nil || o.DoubleStaticIps == nil {
		return nil, false
	}
	return o.DoubleStaticIps, true
}

// HasDoubleStaticIps returns a boolean if a field has been set.
func (o *Dlb) HasDoubleStaticIps() bool {
	if o != nil && o.DoubleStaticIps != nil {
		return true
	}

	return false
}

// SetDoubleStaticIps gets a reference to the given string and assigns it to the DoubleStaticIps field.
func (o *Dlb) SetDoubleStaticIps(v string) {
	o.DoubleStaticIps = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dlb) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dlb) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dlb) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Dlb) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Dlb) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Dlb) SetState(v string) {
	o.State = &v
}

// GetIpWhitelist returns the IpWhitelist field value if set, zero value otherwise.
func (o *Dlb) GetIpWhitelist() []string {
	if o == nil || o.IpWhitelist == nil {
		var ret []string
		return ret
	}
	return *o.IpWhitelist
}

// GetIpWhitelistOk returns a tuple with the IpWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetIpWhitelistOk() (*[]string, bool) {
	if o == nil || o.IpWhitelist == nil {
		return nil, false
	}
	return o.IpWhitelist, true
}

// HasIpWhitelist returns a boolean if a field has been set.
func (o *Dlb) HasIpWhitelist() bool {
	if o != nil && o.IpWhitelist != nil {
		return true
	}

	return false
}

// SetIpWhitelist gets a reference to the given []string and assigns it to the IpWhitelist field.
func (o *Dlb) SetIpWhitelist(v []string) {
	o.IpWhitelist = &v
}

// GetHttpMode returns the HttpMode field value if set, zero value otherwise.
func (o *Dlb) GetHttpMode() string {
	if o == nil || o.HttpMode == nil {
		var ret string
		return ret
	}
	return *o.HttpMode
}

// GetHttpModeOk returns a tuple with the HttpMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetHttpModeOk() (*string, bool) {
	if o == nil || o.HttpMode == nil {
		return nil, false
	}
	return o.HttpMode, true
}

// HasHttpMode returns a boolean if a field has been set.
func (o *Dlb) HasHttpMode() bool {
	if o != nil && o.HttpMode != nil {
		return true
	}

	return false
}

// SetHttpMode gets a reference to the given string and assigns it to the HttpMode field.
func (o *Dlb) SetHttpMode(v string) {
	o.HttpMode = &v
}

// GetDefaultSslEndpoint returns the DefaultSslEndpoint field value if set, zero value otherwise.
func (o *Dlb) GetDefaultSslEndpoint() int32 {
	if o == nil || o.DefaultSslEndpoint == nil {
		var ret int32
		return ret
	}
	return *o.DefaultSslEndpoint
}

// GetDefaultSslEndpointOk returns a tuple with the DefaultSslEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetDefaultSslEndpointOk() (*int32, bool) {
	if o == nil || o.DefaultSslEndpoint == nil {
		return nil, false
	}
	return o.DefaultSslEndpoint, true
}

// HasDefaultSslEndpoint returns a boolean if a field has been set.
func (o *Dlb) HasDefaultSslEndpoint() bool {
	if o != nil && o.DefaultSslEndpoint != nil {
		return true
	}

	return false
}

// SetDefaultSslEndpoint gets a reference to the given int32 and assigns it to the DefaultSslEndpoint field.
func (o *Dlb) SetDefaultSslEndpoint(v int32) {
	o.DefaultSslEndpoint = &v
}

// GetTlsv1 returns the Tlsv1 field value if set, zero value otherwise.
func (o *Dlb) GetTlsv1() bool {
	if o == nil || o.Tlsv1 == nil {
		var ret bool
		return ret
	}
	return *o.Tlsv1
}

// GetTlsv1Ok returns a tuple with the Tlsv1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetTlsv1Ok() (*bool, bool) {
	if o == nil || o.Tlsv1 == nil {
		return nil, false
	}
	return o.Tlsv1, true
}

// HasTlsv1 returns a boolean if a field has been set.
func (o *Dlb) HasTlsv1() bool {
	if o != nil && o.Tlsv1 != nil {
		return true
	}

	return false
}

// SetTlsv1 gets a reference to the given bool and assigns it to the Tlsv1 field.
func (o *Dlb) SetTlsv1(v bool) {
	o.Tlsv1 = &v
}

// GetSslEndpoints returns the SslEndpoints field value if set, zero value otherwise.
func (o *Dlb) GetSslEndpoints() []DlbCoreSslEndpoints {
	if o == nil || o.SslEndpoints == nil {
		var ret []DlbCoreSslEndpoints
		return ret
	}
	return *o.SslEndpoints
}

// GetSslEndpointsOk returns a tuple with the SslEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dlb) GetSslEndpointsOk() (*[]DlbCoreSslEndpoints, bool) {
	if o == nil || o.SslEndpoints == nil {
		return nil, false
	}
	return o.SslEndpoints, true
}

// HasSslEndpoints returns a boolean if a field has been set.
func (o *Dlb) HasSslEndpoints() bool {
	if o != nil && o.SslEndpoints != nil {
		return true
	}

	return false
}

// SetSslEndpoints gets a reference to the given []DlbCoreSslEndpoints and assigns it to the SslEndpoints field.
func (o *Dlb) SetSslEndpoints(v []DlbCoreSslEndpoints) {
	o.SslEndpoints = &v
}

func (o Dlb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.DeploymentId != nil {
		toSerialize["deploymentId"] = o.DeploymentId
	}
	if o.InstanceConfig != nil {
		toSerialize["instanceConfig"] = o.InstanceConfig
	}
	if o.IpAddresses != nil {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if o.StaticIPsDisabled != nil {
		toSerialize["staticIPsDisabled"] = o.StaticIPsDisabled
	}
	if o.VpcId != nil {
		toSerialize["vpcId"] = o.VpcId
	}
	if o.Workers != nil {
		toSerialize["workers"] = o.Workers
	}
	if o.DefaultCipherSuite != nil {
		toSerialize["defaultCipherSuite"] = o.DefaultCipherSuite
	}
	if o.KeepUrlEncoding != nil {
		toSerialize["keepUrlEncoding"] = o.KeepUrlEncoding
	}
	if o.UpstreamTlsv12 != nil {
		toSerialize["upstreamTlsv12"] = o.UpstreamTlsv12
	}
	if o.ProxyReadTimeout != nil {
		toSerialize["proxyReadTimeout"] = o.ProxyReadTimeout
	}
	if o.IpAddressesInfo != nil {
		toSerialize["ipAddressesInfo"] = o.IpAddressesInfo
	}
	if o.DoubleStaticIps != nil {
		toSerialize["doubleStaticIps"] = o.DoubleStaticIps
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.IpWhitelist != nil {
		toSerialize["ipWhitelist"] = o.IpWhitelist
	}
	if o.HttpMode != nil {
		toSerialize["httpMode"] = o.HttpMode
	}
	if o.DefaultSslEndpoint != nil {
		toSerialize["defaultSslEndpoint"] = o.DefaultSslEndpoint
	}
	if o.Tlsv1 != nil {
		toSerialize["tlsv1"] = o.Tlsv1
	}
	if o.SslEndpoints != nil {
		toSerialize["sslEndpoints"] = o.SslEndpoints
	}
	return json.Marshal(toSerialize)
}

type NullableDlb struct {
	value *Dlb
	isSet bool
}

func (v NullableDlb) Get() *Dlb {
	return v.value
}

func (v *NullableDlb) Set(val *Dlb) {
	v.value = val
	v.isSet = true
}

func (v NullableDlb) IsSet() bool {
	return v.isSet
}

func (v *NullableDlb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlb(val *Dlb) *NullableDlb {
	return &NullableDlb{value: val, isSet: true}
}

func (v NullableDlb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

