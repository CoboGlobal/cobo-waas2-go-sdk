/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// TransactionDestinationType The transaction destination type. Refer to [Transaction sources and destinations](/v2/guides/sources-and-destinations) for details. 
type TransactionDestinationType string

// List of TransactionDestinationType
const (
	TRANSACTIONDESTINATIONTYPE_ADDRESS TransactionDestinationType = "Address"
	TRANSACTIONDESTINATIONTYPE_EXCHANGE_WALLET TransactionDestinationType = "ExchangeWallet"
	TRANSACTIONDESTINATIONTYPE_EVM_CONTRACT TransactionDestinationType = "EVM_Contract"
	TRANSACTIONDESTINATIONTYPE_EVM_EIP_191_SIGNATURE TransactionDestinationType = "EVM_EIP_191_Signature"
	TRANSACTIONDESTINATIONTYPE_EVM_EIP_712_SIGNATURE TransactionDestinationType = "EVM_EIP_712_Signature"
	TRANSACTIONDESTINATIONTYPE_DEPOSIT_TO_ADDRESS TransactionDestinationType = "DepositToAddress"
	TRANSACTIONDESTINATIONTYPE_DEPOSIT_TO_WALLET TransactionDestinationType = "DepositToWallet"
)

// All allowed values of TransactionDestinationType enum
var AllowedTransactionDestinationTypeEnumValues = []TransactionDestinationType{
	"Address",
	"ExchangeWallet",
	"EVM_Contract",
	"EVM_EIP_191_Signature",
	"EVM_EIP_712_Signature",
	"DepositToAddress",
	"DepositToWallet",
}

func (v *TransactionDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionDestinationType(value)
	for _, existing := range AllowedTransactionDestinationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = TransactionDestinationType("unknown")
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

