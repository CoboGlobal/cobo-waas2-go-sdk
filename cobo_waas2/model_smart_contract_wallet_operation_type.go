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

// SmartContractWalletOperationType The way you interact with the Smart Contract Wallet.
type SmartContractWalletOperationType string

// List of SmartContractWalletOperationType
const (
	SMARTCONTRACTWALLETOPERATIONTYPE_COBO_SAFE SmartContractWalletOperationType = "CoboSafe"
)

// All allowed values of SmartContractWalletOperationType enum
var AllowedSmartContractWalletOperationTypeEnumValues = []SmartContractWalletOperationType{
	"CoboSafe",
}

func (v *SmartContractWalletOperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SmartContractWalletOperationType(value)
	for _, existing := range AllowedSmartContractWalletOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = SmartContractWalletOperationType("unknown")
    return nil
}

// NewSmartContractWalletOperationTypeFromValue returns a pointer to a valid SmartContractWalletOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSmartContractWalletOperationTypeFromValue(v string) (*SmartContractWalletOperationType, error) {
	ev := SmartContractWalletOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SmartContractWalletOperationType: valid values are %v", v, AllowedSmartContractWalletOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SmartContractWalletOperationType) IsValid() bool {
	for _, existing := range AllowedSmartContractWalletOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SmartContractWalletOperationType value
func (v SmartContractWalletOperationType) Ptr() *SmartContractWalletOperationType {
	return &v
}

type NullableSmartContractWalletOperationType struct {
	value *SmartContractWalletOperationType
	isSet bool
}

func (v NullableSmartContractWalletOperationType) Get() *SmartContractWalletOperationType {
	return v.value
}

func (v *NullableSmartContractWalletOperationType) Set(val *SmartContractWalletOperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractWalletOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractWalletOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractWalletOperationType(val *SmartContractWalletOperationType) *NullableSmartContractWalletOperationType {
	return &NullableSmartContractWalletOperationType{value: val, isSet: true}
}

func (v NullableSmartContractWalletOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractWalletOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

