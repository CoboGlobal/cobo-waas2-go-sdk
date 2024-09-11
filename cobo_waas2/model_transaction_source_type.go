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

// TransactionSourceType The transaction source. Refer to [Transaction sources and destinations](/v2/guides/sources-and-destinations) for details. 
type TransactionSourceType string

// List of TransactionSourceType
const (
	TRANSACTIONSOURCETYPE_ASSET TransactionSourceType = "Asset"
	TRANSACTIONSOURCETYPE_ORG_CONTROLLED TransactionSourceType = "Org-Controlled"
	TRANSACTIONSOURCETYPE_USER_CONTROLLED TransactionSourceType = "User-Controlled"
	TRANSACTIONSOURCETYPE_SAFEWALLET TransactionSourceType = "Safe{Wallet}"
	TRANSACTIONSOURCETYPE_MAIN TransactionSourceType = "Main"
	TRANSACTIONSOURCETYPE_SUB TransactionSourceType = "Sub"
	TRANSACTIONSOURCETYPE_DEPOSIT_FROM_ADDRESS TransactionSourceType = "DepositFromAddress"
	TRANSACTIONSOURCETYPE_DEPOSIT_FROM_WALLET TransactionSourceType = "DepositFromWallet"
	TRANSACTIONSOURCETYPE_DEPOSIT_FROM_LOOP TransactionSourceType = "DepositFromLoop"
)

// All allowed values of TransactionSourceType enum
var AllowedTransactionSourceTypeEnumValues = []TransactionSourceType{
	"Asset",
	"Org-Controlled",
	"User-Controlled",
	"Safe{Wallet}",
	"Main",
	"Sub",
	"DepositFromAddress",
	"DepositFromWallet",
	"DepositFromLoop",
}

func (v *TransactionSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionSourceType(value)
	for _, existing := range AllowedTransactionSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = TransactionSourceType("unknown")
    return nil
}

// NewTransactionSourceTypeFromValue returns a pointer to a valid TransactionSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionSourceTypeFromValue(v string) (*TransactionSourceType, error) {
	ev := TransactionSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionSourceType: valid values are %v", v, AllowedTransactionSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionSourceType) IsValid() bool {
	for _, existing := range AllowedTransactionSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionSourceType value
func (v TransactionSourceType) Ptr() *TransactionSourceType {
	return &v
}

type NullableTransactionSourceType struct {
	value *TransactionSourceType
	isSet bool
}

func (v NullableTransactionSourceType) Get() *TransactionSourceType {
	return v.value
}

func (v *NullableTransactionSourceType) Set(val *TransactionSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionSourceType(val *TransactionSourceType) *NullableTransactionSourceType {
	return &NullableTransactionSourceType{value: val, isSet: true}
}

func (v NullableTransactionSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

