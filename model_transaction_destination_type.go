/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
	"fmt"
)

// TransactionDestinationType The transaction destination type. Possible values include:   - `Address`: An external address.    - `ContractCall`: A smart contract.   - `MessageSign`: A message to be signed.    - `CustodialWallet`: A Custodial Wallet.   - `MPCWallet`: An MPC Wallet.   - `SafeWallet`: A Smart Contract Wallet (Safe{Wallet}).   - `ExchangeWallet`: An Exchange Wallet. 
type TransactionDestinationType string

// List of TransactionDestinationType
const (
	TRANSACTIONDESTINATIONTYPE_ADDRESS TransactionDestinationType = "Address"
	TRANSACTIONDESTINATIONTYPE_CONTRACT_CALL TransactionDestinationType = "ContractCall"
	TRANSACTIONDESTINATIONTYPE_MESSAGE_SIGN TransactionDestinationType = "MessageSign"
	TRANSACTIONDESTINATIONTYPE_CUSTODIAL_WALLET TransactionDestinationType = "CustodialWallet"
	TRANSACTIONDESTINATIONTYPE_MPC_WALLET TransactionDestinationType = "MPCWallet"
	TRANSACTIONDESTINATIONTYPE_SAFE_WALLET TransactionDestinationType = "SafeWallet"
	TRANSACTIONDESTINATIONTYPE_EXCHANGE_WALLET TransactionDestinationType = "ExchangeWallet"
)

// All allowed values of TransactionDestinationType enum
var AllowedTransactionDestinationTypeEnumValues = []TransactionDestinationType{
	"Address",
	"ContractCall",
	"MessageSign",
	"CustodialWallet",
	"MPCWallet",
	"SafeWallet",
	"ExchangeWallet",
}

func (v *TransactionDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TransactionDestinationType(value)
	return nil
}

// NewTransactionDestinationTypeFromValue returns a pointer to a valid TransactionDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionDestinationTypeFromValue(v string) (*TransactionDestinationType, error) {
	ev := TransactionDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionDestinationType: valid values are %v", v, AllowedTransactionDestinationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionDestinationType) IsValid() bool {
	for _, existing := range AllowedTransactionDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionDestinationType value
func (v TransactionDestinationType) Ptr() *TransactionDestinationType {
	return &v
}

type NullableTransactionDestinationType struct {
	value *TransactionDestinationType
	isSet bool
}

func (v NullableTransactionDestinationType) Get() *TransactionDestinationType {
	return v.value
}

func (v *NullableTransactionDestinationType) Set(val *TransactionDestinationType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDestinationType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDestinationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDestinationType(val *TransactionDestinationType) *NullableTransactionDestinationType {
	return &NullableTransactionDestinationType{value: val, isSet: true}
}

func (v NullableTransactionDestinationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDestinationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
