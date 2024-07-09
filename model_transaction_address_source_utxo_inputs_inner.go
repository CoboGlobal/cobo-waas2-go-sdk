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

// checks if the TransactionAddressSourceUtxoInputsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionAddressSourceUtxoInputsInner{}

// TransactionAddressSourceUtxoInputsInner struct for TransactionAddressSourceUtxoInputsInner
type TransactionAddressSourceUtxoInputsInner struct {
	// The transaction hash of UTXO that are consumed in the transaction.
	TxHash *string `json:"tx_hash,omitempty"`
}

// NewTransactionAddressSourceUtxoInputsInner instantiates a new TransactionAddressSourceUtxoInputsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionAddressSourceUtxoInputsInner() *TransactionAddressSourceUtxoInputsInner {
	this := TransactionAddressSourceUtxoInputsInner{}
	return &this
}

// NewTransactionAddressSourceUtxoInputsInnerWithDefaults instantiates a new TransactionAddressSourceUtxoInputsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionAddressSourceUtxoInputsInnerWithDefaults() *TransactionAddressSourceUtxoInputsInner {
	this := TransactionAddressSourceUtxoInputsInner{}
	return &this
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *TransactionAddressSourceUtxoInputsInner) GetTxHash() string {
	if o == nil || IsNil(o.TxHash) {
		var ret string
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionAddressSourceUtxoInputsInner) GetTxHashOk() (*string, bool) {
	if o == nil || IsNil(o.TxHash) {
		return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *TransactionAddressSourceUtxoInputsInner) HasTxHash() bool {
	if o != nil && !IsNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given string and assigns it to the TxHash field.
func (o *TransactionAddressSourceUtxoInputsInner) SetTxHash(v string) {
	o.TxHash = &v
}

func (o TransactionAddressSourceUtxoInputsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionAddressSourceUtxoInputsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TxHash) {
		toSerialize["tx_hash"] = o.TxHash
	}
	return toSerialize, nil
}

type NullableTransactionAddressSourceUtxoInputsInner struct {
	value *TransactionAddressSourceUtxoInputsInner
	isSet bool
}

func (v NullableTransactionAddressSourceUtxoInputsInner) Get() *TransactionAddressSourceUtxoInputsInner {
	return v.value
}

func (v *NullableTransactionAddressSourceUtxoInputsInner) Set(val *TransactionAddressSourceUtxoInputsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionAddressSourceUtxoInputsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionAddressSourceUtxoInputsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionAddressSourceUtxoInputsInner(val *TransactionAddressSourceUtxoInputsInner) *NullableTransactionAddressSourceUtxoInputsInner {
	return &NullableTransactionAddressSourceUtxoInputsInner{value: val, isSet: true}
}

func (v NullableTransactionAddressSourceUtxoInputsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionAddressSourceUtxoInputsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

