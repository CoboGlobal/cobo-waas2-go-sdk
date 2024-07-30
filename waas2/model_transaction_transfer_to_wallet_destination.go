/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionTransferToWalletDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionTransferToWalletDestination{}

// TransactionTransferToWalletDestination Information about the transaction destination, which varies depending on whether you are the initiator or the receiver of the transaction.   - As the initiator, you will see detailed information about the transaction destination, and the `destination` will be displayed as one of the following types: `EVM_Contract`, `EVM_EIP_191_Signature`, `EVM_EIP_712_Signature`, `DepositToAddress`, or `DepositToWallet`. `DepositToWallet` indicates the destination is an Exchange Wallet, while `DepositToAddress` indicates the destination is a wallet of other wallet types or an external address. - As the receiver, you will see the `destination` as the type `Address` or `ExchangeWallet`. `Address` indicates the destination is a wallet of other wallet types than Exchange Wallets or an external address. 
type TransactionTransferToWalletDestination struct {
	DestinationType TransactionDestinationType `json:"destination_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	// The exchange trading account or the sub-wallet ID.
	SubWalletId *string `json:"sub_wallet_id,omitempty"`
	ExchangeId *ExchangeId `json:"exchange_id,omitempty"`
	// The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is `1.5`. 
	Amount string `json:"amount"`
}

type _TransactionTransferToWalletDestination TransactionTransferToWalletDestination

// NewTransactionTransferToWalletDestination instantiates a new TransactionTransferToWalletDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionTransferToWalletDestination(destinationType TransactionDestinationType, walletId string, amount string) *TransactionTransferToWalletDestination {
	this := TransactionTransferToWalletDestination{}
	this.DestinationType = destinationType
	this.WalletId = walletId
	this.Amount = amount
	return &this
}

// NewTransactionTransferToWalletDestinationWithDefaults instantiates a new TransactionTransferToWalletDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionTransferToWalletDestinationWithDefaults() *TransactionTransferToWalletDestination {
	this := TransactionTransferToWalletDestination{}
	return &this
}

// GetDestinationType returns the DestinationType field value
func (o *TransactionTransferToWalletDestination) GetDestinationType() TransactionDestinationType {
	if o == nil {
		var ret TransactionDestinationType
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *TransactionTransferToWalletDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *TransactionTransferToWalletDestination) SetDestinationType(v TransactionDestinationType) {
	o.DestinationType = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionTransferToWalletDestination) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionTransferToWalletDestination) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionTransferToWalletDestination) SetWalletId(v string) {
	o.WalletId = v
}

// GetSubWalletId returns the SubWalletId field value if set, zero value otherwise.
func (o *TransactionTransferToWalletDestination) GetSubWalletId() string {
	if o == nil || IsNil(o.SubWalletId) {
		var ret string
		return ret
	}
	return *o.SubWalletId
}

// GetSubWalletIdOk returns a tuple with the SubWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransferToWalletDestination) GetSubWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubWalletId) {
		return nil, false
	}
	return o.SubWalletId, true
}

// HasSubWalletId returns a boolean if a field has been set.
func (o *TransactionTransferToWalletDestination) HasSubWalletId() bool {
	if o != nil && !IsNil(o.SubWalletId) {
		return true
	}

	return false
}

// SetSubWalletId gets a reference to the given string and assigns it to the SubWalletId field.
func (o *TransactionTransferToWalletDestination) SetSubWalletId(v string) {
	o.SubWalletId = &v
}

// GetExchangeId returns the ExchangeId field value if set, zero value otherwise.
func (o *TransactionTransferToWalletDestination) GetExchangeId() ExchangeId {
	if o == nil || IsNil(o.ExchangeId) {
		var ret ExchangeId
		return ret
	}
	return *o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTransferToWalletDestination) GetExchangeIdOk() (*ExchangeId, bool) {
	if o == nil || IsNil(o.ExchangeId) {
		return nil, false
	}
	return o.ExchangeId, true
}

// HasExchangeId returns a boolean if a field has been set.
func (o *TransactionTransferToWalletDestination) HasExchangeId() bool {
	if o != nil && !IsNil(o.ExchangeId) {
		return true
	}

	return false
}

// SetExchangeId gets a reference to the given ExchangeId and assigns it to the ExchangeId field.
func (o *TransactionTransferToWalletDestination) SetExchangeId(v ExchangeId) {
	o.ExchangeId = &v
}

// GetAmount returns the Amount field value
func (o *TransactionTransferToWalletDestination) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionTransferToWalletDestination) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionTransferToWalletDestination) SetAmount(v string) {
	o.Amount = v
}

func (o TransactionTransferToWalletDestination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionTransferToWalletDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["wallet_id"] = o.WalletId
	if !IsNil(o.SubWalletId) {
		toSerialize["sub_wallet_id"] = o.SubWalletId
	}
	if !IsNil(o.ExchangeId) {
		toSerialize["exchange_id"] = o.ExchangeId
	}
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *TransactionTransferToWalletDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_type",
		"wallet_id",
		"amount",
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

	varTransactionTransferToWalletDestination := _TransactionTransferToWalletDestination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionTransferToWalletDestination)

	if err != nil {
		return err
	}

	*o = TransactionTransferToWalletDestination(varTransactionTransferToWalletDestination)

	return err
}

type NullableTransactionTransferToWalletDestination struct {
	value *TransactionTransferToWalletDestination
	isSet bool
}

func (v NullableTransactionTransferToWalletDestination) Get() *TransactionTransferToWalletDestination {
	return v.value
}

func (v *NullableTransactionTransferToWalletDestination) Set(val *TransactionTransferToWalletDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionTransferToWalletDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionTransferToWalletDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionTransferToWalletDestination(val *TransactionTransferToWalletDestination) *NullableTransactionTransferToWalletDestination {
	return &NullableTransactionTransferToWalletDestination{value: val, isSet: true}
}

func (v NullableTransactionTransferToWalletDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionTransferToWalletDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


