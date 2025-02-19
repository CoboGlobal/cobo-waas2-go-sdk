/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the SwapTokenPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwapTokenPair{}

// SwapTokenPair struct for SwapTokenPair
type SwapTokenPair struct {
	// The source token symbol.
	PayTokenId *string `json:"pay_token_id,omitempty"`
	// The target token symbol.
	ReceiveTokenId *string `json:"receive_token_id,omitempty"`
}

// NewSwapTokenPair instantiates a new SwapTokenPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwapTokenPair() *SwapTokenPair {
	this := SwapTokenPair{}
	return &this
}

// NewSwapTokenPairWithDefaults instantiates a new SwapTokenPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwapTokenPairWithDefaults() *SwapTokenPair {
	this := SwapTokenPair{}
	return &this
}

// GetPayTokenId returns the PayTokenId field value if set, zero value otherwise.
func (o *SwapTokenPair) GetPayTokenId() string {
	if o == nil || IsNil(o.PayTokenId) {
		var ret string
		return ret
	}
	return *o.PayTokenId
}

// GetPayTokenIdOk returns a tuple with the PayTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapTokenPair) GetPayTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayTokenId) {
		return nil, false
	}
	return o.PayTokenId, true
}

// HasPayTokenId returns a boolean if a field has been set.
func (o *SwapTokenPair) HasPayTokenId() bool {
	if o != nil && !IsNil(o.PayTokenId) {
		return true
	}

	return false
}

// SetPayTokenId gets a reference to the given string and assigns it to the PayTokenId field.
func (o *SwapTokenPair) SetPayTokenId(v string) {
	o.PayTokenId = &v
}

// GetReceiveTokenId returns the ReceiveTokenId field value if set, zero value otherwise.
func (o *SwapTokenPair) GetReceiveTokenId() string {
	if o == nil || IsNil(o.ReceiveTokenId) {
		var ret string
		return ret
	}
	return *o.ReceiveTokenId
}

// GetReceiveTokenIdOk returns a tuple with the ReceiveTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwapTokenPair) GetReceiveTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveTokenId) {
		return nil, false
	}
	return o.ReceiveTokenId, true
}

// HasReceiveTokenId returns a boolean if a field has been set.
func (o *SwapTokenPair) HasReceiveTokenId() bool {
	if o != nil && !IsNil(o.ReceiveTokenId) {
		return true
	}

	return false
}

// SetReceiveTokenId gets a reference to the given string and assigns it to the ReceiveTokenId field.
func (o *SwapTokenPair) SetReceiveTokenId(v string) {
	o.ReceiveTokenId = &v
}

func (o SwapTokenPair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwapTokenPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayTokenId) {
		toSerialize["pay_token_id"] = o.PayTokenId
	}
	if !IsNil(o.ReceiveTokenId) {
		toSerialize["receive_token_id"] = o.ReceiveTokenId
	}
	return toSerialize, nil
}

type NullableSwapTokenPair struct {
	value *SwapTokenPair
	isSet bool
}

func (v NullableSwapTokenPair) Get() *SwapTokenPair {
	return v.value
}

func (v *NullableSwapTokenPair) Set(val *SwapTokenPair) {
	v.value = val
	v.isSet = true
}

func (v NullableSwapTokenPair) IsSet() bool {
	return v.isSet
}

func (v *NullableSwapTokenPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwapTokenPair(val *SwapTokenPair) *NullableSwapTokenPair {
	return &NullableSwapTokenPair{value: val, isSet: true}
}

func (v NullableSwapTokenPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwapTokenPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


