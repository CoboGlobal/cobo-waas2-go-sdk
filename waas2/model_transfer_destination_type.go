/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
	"fmt"
)

// TransferDestinationType The transaction destination type. Possible values include: - `Address`: An address, including an address of Custodial Wallets, MPC Wallets, or Smart Contract Wallets (Safe{Wallet}) and an external address. - `ExchangeWallet`: An Exchange Wallet. 
type TransferDestinationType string

// List of TransferDestinationType
const (
	TRANSFERDESTINATIONTYPE_ADDRESS TransferDestinationType = "Address"
	TRANSFERDESTINATIONTYPE_EXCHANGE_WALLET TransferDestinationType = "ExchangeWallet"
)

// All allowed values of TransferDestinationType enum
var AllowedTransferDestinationTypeEnumValues = []TransferDestinationType{
	"Address",
	"ExchangeWallet",
}

func (v *TransferDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferDestinationType(value)
	for _, existing := range AllowedTransferDestinationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = TransferDestinationType("unknown")
    return nil
}

// NewTransferDestinationTypeFromValue returns a pointer to a valid TransferDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferDestinationTypeFromValue(v string) (*TransferDestinationType, error) {
	ev := TransferDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferDestinationType: valid values are %v", v, AllowedTransferDestinationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferDestinationType) IsValid() bool {
	for _, existing := range AllowedTransferDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferDestinationType value
func (v TransferDestinationType) Ptr() *TransferDestinationType {
	return &v
}

type NullableTransferDestinationType struct {
	value *TransferDestinationType
	isSet bool
}

func (v NullableTransferDestinationType) Get() *TransferDestinationType {
	return v.value
}

func (v *NullableTransferDestinationType) Set(val *TransferDestinationType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDestinationType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDestinationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDestinationType(val *TransferDestinationType) *NullableTransferDestinationType {
	return &NullableTransferDestinationType{value: val, isSet: true}
}

func (v NullableTransferDestinationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDestinationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

