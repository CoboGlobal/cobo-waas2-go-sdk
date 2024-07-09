/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
)

// checks if the TransactionAddressDestinationUtxoOutputsOutputsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionAddressDestinationUtxoOutputsOutputsInner{}

// TransactionAddressDestinationUtxoOutputsOutputsInner struct for TransactionAddressDestinationUtxoOutputsOutputsInner
type TransactionAddressDestinationUtxoOutputsOutputsInner struct {
	// The destination address.
	Address *string `json:"address,omitempty"`
	// The transaction amount. For example, if you trade 1.5 ETH, then the amount is `1.5`. 
	Amount *string `json:"amount,omitempty"`
	// The script of the output. It is a programmable code fragment that defines the conditions under which the UTXO can be spent.
	Script *string `json:"script,omitempty"`
}

// NewTransactionAddressDestinationUtxoOutputsOutputsInner instantiates a new TransactionAddressDestinationUtxoOutputsOutputsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionAddressDestinationUtxoOutputsOutputsInner() *TransactionAddressDestinationUtxoOutputsOutputsInner {
	this := TransactionAddressDestinationUtxoOutputsOutputsInner{}
	return &this
}

// NewTransactionAddressDestinationUtxoOutputsOutputsInnerWithDefaults instantiates a new TransactionAddressDestinationUtxoOutputsOutputsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionAddressDestinationUtxoOutputsOutputsInnerWithDefaults() *TransactionAddressDestinationUtxoOutputsOutputsInner {
	this := TransactionAddressDestinationUtxoOutputsOutputsInner{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) SetAddress(v string) {
	o.Address = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *TransactionAddressDestinationUtxoOutputsOutputsInner) SetScript(v string) {
	o.Script = &v
}

func (o TransactionAddressDestinationUtxoOutputsOutputsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionAddressDestinationUtxoOutputsOutputsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	return toSerialize, nil
}

type NullableTransactionAddressDestinationUtxoOutputsOutputsInner struct {
	value *TransactionAddressDestinationUtxoOutputsOutputsInner
	isSet bool
}

func (v NullableTransactionAddressDestinationUtxoOutputsOutputsInner) Get() *TransactionAddressDestinationUtxoOutputsOutputsInner {
	return v.value
}

func (v *NullableTransactionAddressDestinationUtxoOutputsOutputsInner) Set(val *TransactionAddressDestinationUtxoOutputsOutputsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionAddressDestinationUtxoOutputsOutputsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionAddressDestinationUtxoOutputsOutputsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionAddressDestinationUtxoOutputsOutputsInner(val *TransactionAddressDestinationUtxoOutputsOutputsInner) *NullableTransactionAddressDestinationUtxoOutputsOutputsInner {
	return &NullableTransactionAddressDestinationUtxoOutputsOutputsInner{value: val, isSet: true}
}

func (v NullableTransactionAddressDestinationUtxoOutputsOutputsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionAddressDestinationUtxoOutputsOutputsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


