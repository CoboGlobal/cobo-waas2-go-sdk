/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the EvmEip1559FeeBasePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvmEip1559FeeBasePrice{}

// EvmEip1559FeeBasePrice The transaction gas price based on the EIP-1559 fee model.
type EvmEip1559FeeBasePrice struct {
	// The maximum gas fee per gas unit used on the chain, in wei.
	MaxFeePerGas *string `json:"max_fee_per_gas,omitempty"`
	// The maximum priority fee per gas unit used, in wei. The maximum priority fee represents the highest amount of miner tips that you are willing to pay for your transaction.
	MaxPriorityFeePerGas *string `json:"max_priority_fee_per_gas,omitempty"`
}

// NewEvmEip1559FeeBasePrice instantiates a new EvmEip1559FeeBasePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmEip1559FeeBasePrice() *EvmEip1559FeeBasePrice {
	this := EvmEip1559FeeBasePrice{}
	return &this
}

// NewEvmEip1559FeeBasePriceWithDefaults instantiates a new EvmEip1559FeeBasePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmEip1559FeeBasePriceWithDefaults() *EvmEip1559FeeBasePrice {
	this := EvmEip1559FeeBasePrice{}
	return &this
}

// GetMaxFeePerGas returns the MaxFeePerGas field value if set, zero value otherwise.
func (o *EvmEip1559FeeBasePrice) GetMaxFeePerGas() string {
	if o == nil || IsNil(o.MaxFeePerGas) {
		var ret string
		return ret
	}
	return *o.MaxFeePerGas
}

// GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeBasePrice) GetMaxFeePerGasOk() (*string, bool) {
	if o == nil || IsNil(o.MaxFeePerGas) {
		return nil, false
	}
	return o.MaxFeePerGas, true
}

// HasMaxFeePerGas returns a boolean if a field has been set.
func (o *EvmEip1559FeeBasePrice) HasMaxFeePerGas() bool {
	if o != nil && !IsNil(o.MaxFeePerGas) {
		return true
	}

	return false
}

// SetMaxFeePerGas gets a reference to the given string and assigns it to the MaxFeePerGas field.
func (o *EvmEip1559FeeBasePrice) SetMaxFeePerGas(v string) {
	o.MaxFeePerGas = &v
}

// GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field value if set, zero value otherwise.
func (o *EvmEip1559FeeBasePrice) GetMaxPriorityFeePerGas() string {
	if o == nil || IsNil(o.MaxPriorityFeePerGas) {
		var ret string
		return ret
	}
	return *o.MaxPriorityFeePerGas
}

// GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeBasePrice) GetMaxPriorityFeePerGasOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPriorityFeePerGas) {
		return nil, false
	}
	return o.MaxPriorityFeePerGas, true
}

// HasMaxPriorityFeePerGas returns a boolean if a field has been set.
func (o *EvmEip1559FeeBasePrice) HasMaxPriorityFeePerGas() bool {
	if o != nil && !IsNil(o.MaxPriorityFeePerGas) {
		return true
	}

	return false
}

// SetMaxPriorityFeePerGas gets a reference to the given string and assigns it to the MaxPriorityFeePerGas field.
func (o *EvmEip1559FeeBasePrice) SetMaxPriorityFeePerGas(v string) {
	o.MaxPriorityFeePerGas = &v
}

func (o EvmEip1559FeeBasePrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvmEip1559FeeBasePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxFeePerGas) {
		toSerialize["max_fee_per_gas"] = o.MaxFeePerGas
	}
	if !IsNil(o.MaxPriorityFeePerGas) {
		toSerialize["max_priority_fee_per_gas"] = o.MaxPriorityFeePerGas
	}
	return toSerialize, nil
}

type NullableEvmEip1559FeeBasePrice struct {
	value *EvmEip1559FeeBasePrice
	isSet bool
}

func (v NullableEvmEip1559FeeBasePrice) Get() *EvmEip1559FeeBasePrice {
	return v.value
}

func (v *NullableEvmEip1559FeeBasePrice) Set(val *EvmEip1559FeeBasePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmEip1559FeeBasePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmEip1559FeeBasePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmEip1559FeeBasePrice(val *EvmEip1559FeeBasePrice) *NullableEvmEip1559FeeBasePrice {
	return &NullableEvmEip1559FeeBasePrice{value: val, isSet: true}
}

func (v NullableEvmEip1559FeeBasePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmEip1559FeeBasePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


