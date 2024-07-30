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

// checks if the TransactionDepositToWalletDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDepositToWalletDestination{}

// TransactionDepositToWalletDestination Information about the transaction destination, which varies depending on whether you are the initiator or the receiver of the transaction.   - As the initiator, you will see detailed information about the transaction destination, and the `destination` will be displayed as one of the following types: `EVM_Contract`, `EVM_EIP_191_Signature`, `EVM_EIP_712_Signature`, `DepositToAddress`, or `DepositToWallet`. `DepositToWallet` indicates the destination is an Exchange Wallet, while `DepositToAddress` indicates the destination is a wallet of other wallet types or an external address. - As the receiver, you will see the `destination` as the type `Address` or `ExchangeWallet`. `Address` indicates the destination is a wallet of other wallet types than Exchange Wallets or an external address. 
type TransactionDepositToWalletDestination struct {
	DestinationType TransactionDestinationType `json:"destination_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	WalletType WalletType `json:"wallet_type"`
	WalletSubtype WalletSubtype `json:"wallet_subtype"`
	// The exchange trading account or the sub-wallet ID.
	SubWalletId *string `json:"sub_wallet_id,omitempty"`
	ExchangeId *ExchangeId `json:"exchange_id,omitempty"`
	// The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is `1.5`. 
	Amount string `json:"amount"`
}

type _TransactionDepositToWalletDestination TransactionDepositToWalletDestination

// NewTransactionDepositToWalletDestination instantiates a new TransactionDepositToWalletDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDepositToWalletDestination(destinationType TransactionDestinationType, walletId string, walletType WalletType, walletSubtype WalletSubtype, amount string) *TransactionDepositToWalletDestination {
	this := TransactionDepositToWalletDestination{}
	this.DestinationType = destinationType
	this.WalletId = walletId
	this.WalletType = walletType
	this.WalletSubtype = walletSubtype
	this.Amount = amount
	return &this
}

// NewTransactionDepositToWalletDestinationWithDefaults instantiates a new TransactionDepositToWalletDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDepositToWalletDestinationWithDefaults() *TransactionDepositToWalletDestination {
	this := TransactionDepositToWalletDestination{}
	return &this
}

// GetDestinationType returns the DestinationType field value
func (o *TransactionDepositToWalletDestination) GetDestinationType() TransactionDestinationType {
	if o == nil {
		var ret TransactionDestinationType
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToWalletDestination) GetDestinationTypeOk() (*TransactionDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *TransactionDepositToWalletDestination) SetDestinationType(v TransactionDestinationType) {
	o.DestinationType = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionDepositToWalletDestination) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToWalletDestination) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionDepositToWalletDestination) SetWalletId(v string) {
	o.WalletId = v
}

// GetWalletType returns the WalletType field value
func (o *TransactionDepositToWalletDestination) GetWalletType() WalletType {
	if o == nil {
		var ret WalletType
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToWalletDestination) GetWalletTypeOk() (*WalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *TransactionDepositToWalletDestination) SetWalletType(v WalletType) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value
func (o *TransactionDepositToWalletDestination) GetWalletSubtype() WalletSubtype {
	if o == nil {
		var ret WalletSubtype
		return ret
	}

	return o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToWalletDestination) GetWalletSubtypeOk() (*WalletSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletSubtype, true
}

// SetWalletSubtype sets field value
func (o *TransactionDepositToWalletDestination) SetWalletSubtype(v WalletSubtype) {
	o.WalletSubtype = v
}

// GetSubWalletId returns the SubWalletId field value if set, zero value otherwise.
func (o *TransactionDepositToWalletDestination) GetSubWalletId() string {
	if o == nil || IsNil(o.SubWalletId) {
		var ret string
		return ret
	}
	return *o.SubWalletId
}

// GetSubWalletIdOk returns a tuple with the SubWalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDepositToWalletDestination) GetSubWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubWalletId) {
		return nil, false
	}
	return o.SubWalletId, true
}

// HasSubWalletId returns a boolean if a field has been set.
func (o *TransactionDepositToWalletDestination) HasSubWalletId() bool {
	if o != nil && !IsNil(o.SubWalletId) {
		return true
	}

	return false
}

// SetSubWalletId gets a reference to the given string and assigns it to the SubWalletId field.
func (o *TransactionDepositToWalletDestination) SetSubWalletId(v string) {
	o.SubWalletId = &v
}

// GetExchangeId returns the ExchangeId field value if set, zero value otherwise.
func (o *TransactionDepositToWalletDestination) GetExchangeId() ExchangeId {
	if o == nil || IsNil(o.ExchangeId) {
		var ret ExchangeId
		return ret
	}
	return *o.ExchangeId
}

// GetExchangeIdOk returns a tuple with the ExchangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDepositToWalletDestination) GetExchangeIdOk() (*ExchangeId, bool) {
	if o == nil || IsNil(o.ExchangeId) {
		return nil, false
	}
	return o.ExchangeId, true
}

// HasExchangeId returns a boolean if a field has been set.
func (o *TransactionDepositToWalletDestination) HasExchangeId() bool {
	if o != nil && !IsNil(o.ExchangeId) {
		return true
	}

	return false
}

// SetExchangeId gets a reference to the given ExchangeId and assigns it to the ExchangeId field.
func (o *TransactionDepositToWalletDestination) SetExchangeId(v ExchangeId) {
	o.ExchangeId = &v
}

// GetAmount returns the Amount field value
func (o *TransactionDepositToWalletDestination) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionDepositToWalletDestination) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionDepositToWalletDestination) SetAmount(v string) {
	o.Amount = v
}

func (o TransactionDepositToWalletDestination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDepositToWalletDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["wallet_type"] = o.WalletType
	toSerialize["wallet_subtype"] = o.WalletSubtype
	if !IsNil(o.SubWalletId) {
		toSerialize["sub_wallet_id"] = o.SubWalletId
	}
	if !IsNil(o.ExchangeId) {
		toSerialize["exchange_id"] = o.ExchangeId
	}
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *TransactionDepositToWalletDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_type",
		"wallet_id",
		"wallet_type",
		"wallet_subtype",
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

	varTransactionDepositToWalletDestination := _TransactionDepositToWalletDestination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionDepositToWalletDestination)

	if err != nil {
		return err
	}

	*o = TransactionDepositToWalletDestination(varTransactionDepositToWalletDestination)

	return err
}

type NullableTransactionDepositToWalletDestination struct {
	value *TransactionDepositToWalletDestination
	isSet bool
}

func (v NullableTransactionDepositToWalletDestination) Get() *TransactionDepositToWalletDestination {
	return v.value
}

func (v *NullableTransactionDepositToWalletDestination) Set(val *TransactionDepositToWalletDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDepositToWalletDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDepositToWalletDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDepositToWalletDestination(val *TransactionDepositToWalletDestination) *NullableTransactionDepositToWalletDestination {
	return &NullableTransactionDepositToWalletDestination{value: val, isSet: true}
}

func (v NullableTransactionDepositToWalletDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDepositToWalletDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


