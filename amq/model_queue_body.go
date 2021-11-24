/*
 * Anypoint MQ specfication
 *
 * Anypoint MQ API specification
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amq

import (
	"encoding/json"
)

// QueueBody struct for QueueBody
type QueueBody struct {
	DefaultTtl *int32 `json:"defaultTtl,omitempty"`
	DefaultLockTtl *int32 `json:"defaultLockTtl,omitempty"`
	Type *string `json:"type,omitempty"`
	Encrypted *bool `json:"encrypted,omitempty"`
	DefaultDeliveryDelay *int32 `json:"defaultDeliveryDelay,omitempty"`
	Fifo *bool `json:"fifo,omitempty"`
}

// NewQueueBody instantiates a new QueueBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueBody() *QueueBody {
	this := QueueBody{}
	return &this
}

// NewQueueBodyWithDefaults instantiates a new QueueBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueBodyWithDefaults() *QueueBody {
	this := QueueBody{}
	return &this
}

// GetDefaultTtl returns the DefaultTtl field value if set, zero value otherwise.
func (o *QueueBody) GetDefaultTtl() int32 {
	if o == nil || o.DefaultTtl == nil {
		var ret int32
		return ret
	}
	return *o.DefaultTtl
}

// GetDefaultTtlOk returns a tuple with the DefaultTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueBody) GetDefaultTtlOk() (*int32, bool) {
	if o == nil || o.DefaultTtl == nil {
		return nil, false
	}
	return o.DefaultTtl, true
}

// HasDefaultTtl returns a boolean if a field has been set.
func (o *QueueBody) HasDefaultTtl() bool {
	if o != nil && o.DefaultTtl != nil {
		return true
	}

	return false
}

// SetDefaultTtl gets a reference to the given int32 and assigns it to the DefaultTtl field.
func (o *QueueBody) SetDefaultTtl(v int32) {
	o.DefaultTtl = &v
}

// GetDefaultLockTtl returns the DefaultLockTtl field value if set, zero value otherwise.
func (o *QueueBody) GetDefaultLockTtl() int32 {
	if o == nil || o.DefaultLockTtl == nil {
		var ret int32
		return ret
	}
	return *o.DefaultLockTtl
}

// GetDefaultLockTtlOk returns a tuple with the DefaultLockTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueBody) GetDefaultLockTtlOk() (*int32, bool) {
	if o == nil || o.DefaultLockTtl == nil {
		return nil, false
	}
	return o.DefaultLockTtl, true
}

// HasDefaultLockTtl returns a boolean if a field has been set.
func (o *QueueBody) HasDefaultLockTtl() bool {
	if o != nil && o.DefaultLockTtl != nil {
		return true
	}

	return false
}

// SetDefaultLockTtl gets a reference to the given int32 and assigns it to the DefaultLockTtl field.
func (o *QueueBody) SetDefaultLockTtl(v int32) {
	o.DefaultLockTtl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueueBody) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueBody) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueueBody) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueueBody) SetType(v string) {
	o.Type = &v
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *QueueBody) GetEncrypted() bool {
	if o == nil || o.Encrypted == nil {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueBody) GetEncryptedOk() (*bool, bool) {
	if o == nil || o.Encrypted == nil {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *QueueBody) HasEncrypted() bool {
	if o != nil && o.Encrypted != nil {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *QueueBody) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetDefaultDeliveryDelay returns the DefaultDeliveryDelay field value if set, zero value otherwise.
func (o *QueueBody) GetDefaultDeliveryDelay() int32 {
	if o == nil || o.DefaultDeliveryDelay == nil {
		var ret int32
		return ret
	}
	return *o.DefaultDeliveryDelay
}

// GetDefaultDeliveryDelayOk returns a tuple with the DefaultDeliveryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueBody) GetDefaultDeliveryDelayOk() (*int32, bool) {
	if o == nil || o.DefaultDeliveryDelay == nil {
		return nil, false
	}
	return o.DefaultDeliveryDelay, true
}

// HasDefaultDeliveryDelay returns a boolean if a field has been set.
func (o *QueueBody) HasDefaultDeliveryDelay() bool {
	if o != nil && o.DefaultDeliveryDelay != nil {
		return true
	}

	return false
}

// SetDefaultDeliveryDelay gets a reference to the given int32 and assigns it to the DefaultDeliveryDelay field.
func (o *QueueBody) SetDefaultDeliveryDelay(v int32) {
	o.DefaultDeliveryDelay = &v
}

// GetFifo returns the Fifo field value if set, zero value otherwise.
func (o *QueueBody) GetFifo() bool {
	if o == nil || o.Fifo == nil {
		var ret bool
		return ret
	}
	return *o.Fifo
}

// GetFifoOk returns a tuple with the Fifo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueBody) GetFifoOk() (*bool, bool) {
	if o == nil || o.Fifo == nil {
		return nil, false
	}
	return o.Fifo, true
}

// HasFifo returns a boolean if a field has been set.
func (o *QueueBody) HasFifo() bool {
	if o != nil && o.Fifo != nil {
		return true
	}

	return false
}

// SetFifo gets a reference to the given bool and assigns it to the Fifo field.
func (o *QueueBody) SetFifo(v bool) {
	o.Fifo = &v
}

func (o QueueBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultTtl != nil {
		toSerialize["defaultTtl"] = o.DefaultTtl
	}
	if o.DefaultLockTtl != nil {
		toSerialize["defaultLockTtl"] = o.DefaultLockTtl
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Encrypted != nil {
		toSerialize["encrypted"] = o.Encrypted
	}
	if o.DefaultDeliveryDelay != nil {
		toSerialize["defaultDeliveryDelay"] = o.DefaultDeliveryDelay
	}
	if o.Fifo != nil {
		toSerialize["fifo"] = o.Fifo
	}
	return json.Marshal(toSerialize)
}

type NullableQueueBody struct {
	value *QueueBody
	isSet bool
}

func (v NullableQueueBody) Get() *QueueBody {
	return v.value
}

func (v *NullableQueueBody) Set(val *QueueBody) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueBody) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueBody(val *QueueBody) *NullableQueueBody {
	return &NullableQueueBody{value: val, isSet: true}
}

func (v NullableQueueBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

