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

// checks if the UTXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UTXO{}

// UTXO The UTXO information.
type UTXO struct {
	// The transaction hash of the UTXO.
	TxHash *string `json:"tx_hash,omitempty"`
	// The output index of the UTXO.
	VoutN *int32 `json:"vout_n,omitempty"`
	// The address of the UTXO.
	Address *string `json:"address,omitempty"`
	TokenBalances []TokenBalance `json:"token_balances,omitempty"`
	// Whether the UTXO comes from a coinbase transaction.
	IsCoinbase *bool `json:"is_coinbase,omitempty"`
	// Whether the UTXO is locked.
	IsLocked *bool `json:"is_locked,omitempty"`
	// The number of confirmations for the UTXO.
	ConfirmedNumber *int32 `json:"confirmed_number,omitempty"`
}

// NewUTXO instantiates a new UTXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUTXO() *UTXO {
	this := UTXO{}
	return &this
}

// NewUTXOWithDefaults instantiates a new UTXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUTXOWithDefaults() *UTXO {
	this := UTXO{}
	return &this
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *UTXO) GetTxHash() string {
	if o == nil || IsNil(o.TxHash) {
		var ret string
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UTXO) GetTxHashOk() (*string, bool) {
	if o == nil || IsNil(o.TxHash) {
		return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *UTXO) HasTxHash() bool {
	if o != nil && !IsNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given string and assigns it to the TxHash field.
func (o *UTXO) SetTxHash(v string) {
	o.TxHash = &v
}

// GetVoutN returns the VoutN field value if set, zero value otherwise.
func (o *UTXO) GetVoutN() int32 {
	if o == nil || IsNil(o.VoutN) {
		var ret int32
		return ret
	}
	return *o.VoutN
}

// GetVoutNOk returns a tuple with the VoutN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UTXO) GetVoutNOk() (*int32, bool) {
	if o == nil || IsNil(o.VoutN) {
		return nil, false
	}
	return o.VoutN, true
}

// HasVoutN returns a boolean if a field has been set.
func (o *UTXO) HasVoutN() bool {
	if o != nil && !IsNil(o.VoutN) {
		return true
	}

	return false
}

// SetVoutN gets a reference to the given int32 and assigns it to the VoutN field.
func (o *UTXO) SetVoutN(v int32) {
	o.VoutN = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UTXO) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UTXO) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UTXO) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UTXO) SetAddress(v string) {
	o.Address = &v
}

// GetTokenBalances returns the TokenBalances field value if set, zero value otherwise.
func (o *UTXO) GetTokenBalances() []TokenBalance {
	if o == nil || IsNil(o.TokenBalances) {
		var ret []TokenBalance
		return ret
	}
	return o.TokenBalances
}

// GetTokenBalancesOk returns a tuple with the TokenBalances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UTXO) GetTokenBalancesOk() ([]TokenBalance, bool) {
	if o == nil || IsNil(o.TokenBalances) {
		return nil, false
	}
	return o.TokenBalances, true
}

// HasTokenBalances returns a boolean if a field has been set.
func (o *UTXO) HasTokenBalances() bool {
	if o != nil && !IsNil(o.TokenBalances) {
		return true
	}

	return false
}

// SetTokenBalances gets a reference to the given []TokenBalance and assigns it to the TokenBalances field.
func (o *UTXO) SetTokenBalances(v []TokenBalance) {
	o.TokenBalances = v
}

// GetIsCoinbase returns the IsCoinbase field value if set, zero value otherwise.
func (o *UTXO) GetIsCoinbase() bool {
	if o == nil || IsNil(o.IsCoinbase) {
		var ret bool
		return ret
	}
	return *o.IsCoinbase
}

// GetIsCoinbaseOk returns a tuple with the IsCoinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UTXO) GetIsCoinbaseOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCoinbase) {
		return nil, false
	}
	return o.IsCoinbase, true
}

// HasIsCoinbase returns a boolean if a field has been set.
func (o *UTXO) HasIsCoinbase() bool {
	if o != nil && !IsNil(o.IsCoinbase) {
		return true
	}

	return false
}

// SetIsCoinbase gets a reference to the given bool and assigns it to the IsCoinbase field.
func (o *UTXO) SetIsCoinbase(v bool) {
	o.IsCoinbase = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *UTXO) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UTXO) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *UTXO) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *UTXO) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetConfirmedNumber returns the ConfirmedNumber field value if set, zero value otherwise.
func (o *UTXO) GetConfirmedNumber() int32 {
	if o == nil || IsNil(o.ConfirmedNumber) {
		var ret int32
		return ret
	}
	return *o.ConfirmedNumber
}

// GetConfirmedNumberOk returns a tuple with the ConfirmedNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UTXO) GetConfirmedNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfirmedNumber) {
		return nil, false
	}
	return o.ConfirmedNumber, true
}

// HasConfirmedNumber returns a boolean if a field has been set.
func (o *UTXO) HasConfirmedNumber() bool {
	if o != nil && !IsNil(o.ConfirmedNumber) {
		return true
	}

	return false
}

// SetConfirmedNumber gets a reference to the given int32 and assigns it to the ConfirmedNumber field.
func (o *UTXO) SetConfirmedNumber(v int32) {
	o.ConfirmedNumber = &v
}

func (o UTXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UTXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TxHash) {
		toSerialize["tx_hash"] = o.TxHash
	}
	if !IsNil(o.VoutN) {
		toSerialize["vout_n"] = o.VoutN
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.TokenBalances) {
		toSerialize["token_balances"] = o.TokenBalances
	}
	if !IsNil(o.IsCoinbase) {
		toSerialize["is_coinbase"] = o.IsCoinbase
	}
	if !IsNil(o.IsLocked) {
		toSerialize["is_locked"] = o.IsLocked
	}
	if !IsNil(o.ConfirmedNumber) {
		toSerialize["confirmed_number"] = o.ConfirmedNumber
	}
	return toSerialize, nil
}

type NullableUTXO struct {
	value *UTXO
	isSet bool
}

func (v NullableUTXO) Get() *UTXO {
	return v.value
}

func (v *NullableUTXO) Set(val *UTXO) {
	v.value = val
	v.isSet = true
}

func (v NullableUTXO) IsSet() bool {
	return v.isSet
}

func (v *NullableUTXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUTXO(val *UTXO) *NullableUTXO {
	return &NullableUTXO{value: val, isSet: true}
}

func (v NullableUTXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUTXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


