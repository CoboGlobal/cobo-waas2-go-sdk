/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ExchangeWalletTransactionAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeWalletTransactionAddress{}

// ExchangeWalletTransactionAddress struct for ExchangeWalletTransactionAddress
type ExchangeWalletTransactionAddress struct {
	Type TransactionAddressType `json:"type"`
	// Address
	Address *string `json:"address,omitempty"`
	// Address memo
	Memo *string `json:"memo,omitempty"`
	// Unique id of the wallet.
	WalletId string `json:"wallet_id"`
	ExchangeId ExchangeId `json:"exchange_id"`
	// Exchange trading account or any sub wallet info for transfer.
	SubWalletId *string `json:"sub_wallet_id,omitempty"`
}

type _ExchangeWalletTransactionAddress ExchangeWalletTransactionAddress

// NewExchangeWalletTransactionAddress instantiates a new ExchangeWalletTransactionAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeWalletTransactionAddress(type_ TransactionAddressType, walletId string, exchangeId ExchangeId) *ExchangeWalletTransactionAddress {
	this := ExchangeWalletTransactionAddress{}
	this.Type = type_
	this.WalletId = walletId
	this.ExchangeId = exchangeId
	return &this
}

// NewExchangeWalletTransactionAddressWithDefaults instantiates a new ExchangeWalletTransactionAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeWalletTransactionAddressWithDefaults() *ExchangeWalletTransactionAddress {
	this := ExchangeWalletTransactionAddress{}
	return &this
}

// GetType returns the Type field value
func (o *ExchangeWalletTransactionAddress) GetType() TransactionAddressType {
	if o == nil {
		var ret TransactionAddressType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletTransactionAddress) GetTypeOk() (*TransactionAddressType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExchangeWalletTransactionAddress) SetType(v TransactionAddressType) {
	o.Type = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ExchangeWalletTransactionAddress) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeWalletTransactionAddress) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ExchangeWalletTransactionAddress) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ExchangeWalletTransactionAddress) SetAddress(v string) {
	o.Address = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *ExchangeWalletTransactionAddress) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeWalletTransactionAddress) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *ExchangeWalletTransactionAddress) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *ExchangeWalletTransactionAddress) SetMemo(v string) {
	o.Memo = &v
}

// GetWalletId returns the WalletId field value
func (o *ExchangeWalletTransactionAddress) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletTransactionAddress) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *ExchangeWalletTransactionAddress) SetWalletId(v string) {
	o.WalletId = v
}

// GetExchangeId returns the ExchangeId field value
func (o *ExchangeWalletTransactionAddress) GetExchangeId() ExchangeId {
	if o == nil {
		var ret ExchangeId
		return ret
	}

	return o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value
// and a boolean to check if the value has been set.
func (o *ExchangeWalletTransactionAddress) GetExchangeIdOk() (*ExchangeId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeId, true
}

// SetExchangeId sets field value
func (o *ExchangeWalletTransactionAddress) SetExchangeId(v ExchangeId) {
	o.ExchangeId = v
}

// GetSubWalletId returns the SubWalletId field value if set, zero value otherwise.
func (o *ExchangeWalletTransactionAddress) GetSubWalletId() string {
	if o == nil || IsNil(o.SubWalletId) {
		var ret string
		return ret
	}
	return *o.SubWalletId
}

// GetSubWalletIdOk returns a tuple with the SubWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeWalletTransactionAddress) GetSubWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubWalletId) {
		return nil, false
	}
	return o.SubWalletId, true
}

// HasSubWalletId returns a boolean if a field has been set.
func (o *ExchangeWalletTransactionAddress) HasSubWalletId() bool {
	if o != nil && !IsNil(o.SubWalletId) {
		return true
	}

	return false
}

// SetSubWalletId gets a reference to the given string and assigns it to the SubWalletId field.
func (o *ExchangeWalletTransactionAddress) SetSubWalletId(v string) {
	o.SubWalletId = &v
}

func (o ExchangeWalletTransactionAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeWalletTransactionAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["exchange_id"] = o.ExchangeId
	if !IsNil(o.SubWalletId) {
		toSerialize["sub_wallet_id"] = o.SubWalletId
	}
	return toSerialize, nil
}

func (o *ExchangeWalletTransactionAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"wallet_id",
		"exchange_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExchangeWalletTransactionAddress := _ExchangeWalletTransactionAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeWalletTransactionAddress)

	if err != nil {
		return err
	}

	*o = ExchangeWalletTransactionAddress(varExchangeWalletTransactionAddress)

	return err
}

type NullableExchangeWalletTransactionAddress struct {
	value *ExchangeWalletTransactionAddress
	isSet bool
}

func (v NullableExchangeWalletTransactionAddress) Get() *ExchangeWalletTransactionAddress {
	return v.value
}

func (v *NullableExchangeWalletTransactionAddress) Set(val *ExchangeWalletTransactionAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeWalletTransactionAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeWalletTransactionAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeWalletTransactionAddress(val *ExchangeWalletTransactionAddress) *NullableExchangeWalletTransactionAddress {
	return &NullableExchangeWalletTransactionAddress{value: val, isSet: true}
}

func (v NullableExchangeWalletTransactionAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeWalletTransactionAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


