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

// Queue struct for Queue
type Queue struct {
	QueueId *string `json:"queueId,omitempty"`
	ExchangeId *string `json:"exchangeId,omitempty"`
	DefaultTtl *int32 `json:"defaultTtl,omitempty"`
	DefaultLockTtl *int32 `json:"defaultLockTtl,omitempty"`
	Type *string `json:"type,omitempty"`
	Encrypted *bool `json:"encrypted,omitempty"`
	DefaultDeliveryDelay *int32 `json:"defaultDeliveryDelay,omitempty"`
	DeadLetterQueueId *string `json:"deadLetterQueueId,omitempty"`
	MaxDeliveries *int32 `json:"maxDeliveries,omitempty"`
	Fifo *bool `json:"fifo,omitempty"`
}

// NewQueue instantiates a new Queue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueue() *Queue {
	this := Queue{}
	return &this
}

// NewQueueWithDefaults instantiates a new Queue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueWithDefaults() *Queue {
	this := Queue{}
	return &this
}

// GetQueueId returns the QueueId field value if set, zero value otherwise.
func (o *Queue) GetQueueId() string {
	if o == nil || o.QueueId == nil {
		var ret string
		return ret
	}
	return *o.QueueId
}

// GetQueueIdOk returns a tuple with the QueueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetQueueIdOk() (*string, bool) {
	if o == nil || o.QueueId == nil {
		return nil, false
	}
	return o.QueueId, true
}

// HasQueueId returns a boolean if a field has been set.
func (o *Queue) HasQueueId() bool {
	if o != nil && o.QueueId != nil {
		return true
	}

	return false
}

// SetQueueId gets a reference to the given string and assigns it to the QueueId field.
func (o *Queue) SetQueueId(v string) {
	o.QueueId = &v
}

// GetExchangeId returns the ExchangeId field value if set, zero value otherwise.
func (o *Queue) GetExchangeId() string {
	if o == nil || o.ExchangeId == nil {
		var ret string
		return ret
	}
	return *o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetExchangeIdOk() (*string, bool) {
	if o == nil || o.ExchangeId == nil {
		return nil, false
	}
	return o.ExchangeId, true
}

// HasExchangeId returns a boolean if a field has been set.
func (o *Queue) HasExchangeId() bool {
	if o != nil && o.ExchangeId != nil {
		return true
	}

	return false
}

// SetExchangeId gets a reference to the given string and assigns it to the ExchangeId field.
func (o *Queue) SetExchangeId(v string) {
	o.ExchangeId = &v
}

// GetDefaultTtl returns the DefaultTtl field value if set, zero value otherwise.
func (o *Queue) GetDefaultTtl() int32 {
	if o == nil || o.DefaultTtl == nil {
		var ret int32
		return ret
	}
	return *o.DefaultTtl
}

// GetDefaultTtlOk returns a tuple with the DefaultTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDefaultTtlOk() (*int32, bool) {
	if o == nil || o.DefaultTtl == nil {
		return nil, false
	}
	return o.DefaultTtl, true
}

// HasDefaultTtl returns a boolean if a field has been set.
func (o *Queue) HasDefaultTtl() bool {
	if o != nil && o.DefaultTtl != nil {
		return true
	}

	return false
}

// SetDefaultTtl gets a reference to the given int32 and assigns it to the DefaultTtl field.
func (o *Queue) SetDefaultTtl(v int32) {
	o.DefaultTtl = &v
}

// GetDefaultLockTtl returns the DefaultLockTtl field value if set, zero value otherwise.
func (o *Queue) GetDefaultLockTtl() int32 {
	if o == nil || o.DefaultLockTtl == nil {
		var ret int32
		return ret
	}
	return *o.DefaultLockTtl
}

// GetDefaultLockTtlOk returns a tuple with the DefaultLockTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDefaultLockTtlOk() (*int32, bool) {
	if o == nil || o.DefaultLockTtl == nil {
		return nil, false
	}
	return o.DefaultLockTtl, true
}

// HasDefaultLockTtl returns a boolean if a field has been set.
func (o *Queue) HasDefaultLockTtl() bool {
	if o != nil && o.DefaultLockTtl != nil {
		return true
	}

	return false
}

// SetDefaultLockTtl gets a reference to the given int32 and assigns it to the DefaultLockTtl field.
func (o *Queue) SetDefaultLockTtl(v int32) {
	o.DefaultLockTtl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Queue) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Queue) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Queue) SetType(v string) {
	o.Type = &v
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *Queue) GetEncrypted() bool {
	if o == nil || o.Encrypted == nil {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetEncryptedOk() (*bool, bool) {
	if o == nil || o.Encrypted == nil {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *Queue) HasEncrypted() bool {
	if o != nil && o.Encrypted != nil {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *Queue) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetDefaultDeliveryDelay returns the DefaultDeliveryDelay field value if set, zero value otherwise.
func (o *Queue) GetDefaultDeliveryDelay() int32 {
	if o == nil || o.DefaultDeliveryDelay == nil {
		var ret int32
		return ret
	}
	return *o.DefaultDeliveryDelay
}

// GetDefaultDeliveryDelayOk returns a tuple with the DefaultDeliveryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDefaultDeliveryDelayOk() (*int32, bool) {
	if o == nil || o.DefaultDeliveryDelay == nil {
		return nil, false
	}
	return o.DefaultDeliveryDelay, true
}

// HasDefaultDeliveryDelay returns a boolean if a field has been set.
func (o *Queue) HasDefaultDeliveryDelay() bool {
	if o != nil && o.DefaultDeliveryDelay != nil {
		return true
	}

	return false
}

// SetDefaultDeliveryDelay gets a reference to the given int32 and assigns it to the DefaultDeliveryDelay field.
func (o *Queue) SetDefaultDeliveryDelay(v int32) {
	o.DefaultDeliveryDelay = &v
}

// GetDeadLetterQueueId returns the DeadLetterQueueId field value if set, zero value otherwise.
func (o *Queue) GetDeadLetterQueueId() string {
	if o == nil || o.DeadLetterQueueId == nil {
		var ret string
		return ret
	}
	return *o.DeadLetterQueueId
}

// GetDeadLetterQueueIdOk returns a tuple with the DeadLetterQueueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetDeadLetterQueueIdOk() (*string, bool) {
	if o == nil || o.DeadLetterQueueId == nil {
		return nil, false
	}
	return o.DeadLetterQueueId, true
}

// HasDeadLetterQueueId returns a boolean if a field has been set.
func (o *Queue) HasDeadLetterQueueId() bool {
	if o != nil && o.DeadLetterQueueId != nil {
		return true
	}

	return false
}

// SetDeadLetterQueueId gets a reference to the given string and assigns it to the DeadLetterQueueId field.
func (o *Queue) SetDeadLetterQueueId(v string) {
	o.DeadLetterQueueId = &v
}

// GetMaxDeliveries returns the MaxDeliveries field value if set, zero value otherwise.
func (o *Queue) GetMaxDeliveries() int32 {
	if o == nil || o.MaxDeliveries == nil {
		var ret int32
		return ret
	}
	return *o.MaxDeliveries
}

// GetMaxDeliveriesOk returns a tuple with the MaxDeliveries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetMaxDeliveriesOk() (*int32, bool) {
	if o == nil || o.MaxDeliveries == nil {
		return nil, false
	}
	return o.MaxDeliveries, true
}

// HasMaxDeliveries returns a boolean if a field has been set.
func (o *Queue) HasMaxDeliveries() bool {
	if o != nil && o.MaxDeliveries != nil {
		return true
	}

	return false
}

// SetMaxDeliveries gets a reference to the given int32 and assigns it to the MaxDeliveries field.
func (o *Queue) SetMaxDeliveries(v int32) {
	o.MaxDeliveries = &v
}

// GetFifo returns the Fifo field value if set, zero value otherwise.
func (o *Queue) GetFifo() bool {
	if o == nil || o.Fifo == nil {
		var ret bool
		return ret
	}
	return *o.Fifo
}

// GetFifoOk returns a tuple with the Fifo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Queue) GetFifoOk() (*bool, bool) {
	if o == nil || o.Fifo == nil {
		return nil, false
	}
	return o.Fifo, true
}

// HasFifo returns a boolean if a field has been set.
func (o *Queue) HasFifo() bool {
	if o != nil && o.Fifo != nil {
		return true
	}

	return false
}

// SetFifo gets a reference to the given bool and assigns it to the Fifo field.
func (o *Queue) SetFifo(v bool) {
	o.Fifo = &v
}

func (o Queue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QueueId != nil {
		toSerialize["queueId"] = o.QueueId
	}
	if o.ExchangeId != nil {
		toSerialize["exchangeId"] = o.ExchangeId
	}
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
	if o.DeadLetterQueueId != nil {
		toSerialize["deadLetterQueueId"] = o.DeadLetterQueueId
	}
	if o.MaxDeliveries != nil {
		toSerialize["maxDeliveries"] = o.MaxDeliveries
	}
	if o.Fifo != nil {
		toSerialize["fifo"] = o.Fifo
	}
	return json.Marshal(toSerialize)
}

type NullableQueue struct {
	value *Queue
	isSet bool
}

func (v NullableQueue) Get() *Queue {
	return v.value
}

func (v *NullableQueue) Set(val *Queue) {
	v.value = val
	v.isSet = true
}

func (v NullableQueue) IsSet() bool {
	return v.isSet
}

func (v *NullableQueue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueue(val *Queue) *NullableQueue {
	return &NullableQueue{value: val, isSet: true}
}

func (v NullableQueue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


