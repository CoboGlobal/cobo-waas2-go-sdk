/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the RetryWebhookEventById201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetryWebhookEventById201Response{}

// RetryWebhookEventById201Response struct for RetryWebhookEventById201Response
type RetryWebhookEventById201Response struct {
	// Whether the retry request has been successfully sent.
	Retried *bool `json:"retried,omitempty"`
}

// NewRetryWebhookEventById201Response instantiates a new RetryWebhookEventById201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetryWebhookEventById201Response() *RetryWebhookEventById201Response {
	this := RetryWebhookEventById201Response{}
	return &this
}

// NewRetryWebhookEventById201ResponseWithDefaults instantiates a new RetryWebhookEventById201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetryWebhookEventById201ResponseWithDefaults() *RetryWebhookEventById201Response {
	this := RetryWebhookEventById201Response{}
	return &this
}

// GetRetried returns the Retried field value if set, zero value otherwise.
func (o *RetryWebhookEventById201Response) GetRetried() bool {
	if o == nil || IsNil(o.Retried) {
		var ret bool
		return ret
	}
	return *o.Retried
}

// GetRetriedOk returns a tuple with the Retried field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryWebhookEventById201Response) GetRetriedOk() (*bool, bool) {
	if o == nil || IsNil(o.Retried) {
		return nil, false
	}
	return o.Retried, true
}

// HasRetried returns a boolean if a field has been set.
func (o *RetryWebhookEventById201Response) HasRetried() bool {
	if o != nil && !IsNil(o.Retried) {
		return true
	}

	return false
}

// SetRetried gets a reference to the given bool and assigns it to the Retried field.
func (o *RetryWebhookEventById201Response) SetRetried(v bool) {
	o.Retried = &v
}

func (o RetryWebhookEventById201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetryWebhookEventById201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Retried) {
		toSerialize["retried"] = o.Retried
	}
	return toSerialize, nil
}

type NullableRetryWebhookEventById201Response struct {
	value *RetryWebhookEventById201Response
	isSet bool
}

func (v NullableRetryWebhookEventById201Response) Get() *RetryWebhookEventById201Response {
	return v.value
}

func (v *NullableRetryWebhookEventById201Response) Set(val *RetryWebhookEventById201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRetryWebhookEventById201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRetryWebhookEventById201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetryWebhookEventById201Response(val *RetryWebhookEventById201Response) *NullableRetryWebhookEventById201Response {
	return &NullableRetryWebhookEventById201Response{value: val, isSet: true}
}

func (v NullableRetryWebhookEventById201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetryWebhookEventById201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

