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

// DestinationWalletType The type of the destination wallet.
type DestinationWalletType string

// List of DestinationWalletType
const (
	DESTINATIONWALLETTYPE_EXCHANGES_OR_VASP DestinationWalletType = "EXCHANGES_OR_VASP"
	DESTINATIONWALLETTYPE_SELF_CUSTODY_WALLET DestinationWalletType = "SELF_CUSTODY_WALLET"
)

// All allowed values of DestinationWalletType enum
var AllowedDestinationWalletTypeEnumValues = []DestinationWalletType{
	"EXCHANGES_OR_VASP",
	"SELF_CUSTODY_WALLET",
}

func (v *DestinationWalletType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DestinationWalletType(value)
	for _, existing := range AllowedDestinationWalletTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = DestinationWalletType("unknown")
    return nil
}

// NewDestinationWalletTypeFromValue returns a pointer to a valid DestinationWalletType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDestinationWalletTypeFromValue(v string) (*DestinationWalletType, error) {
	ev := DestinationWalletType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DestinationWalletType: valid values are %v", v, AllowedDestinationWalletTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DestinationWalletType) IsValid() bool {
	for _, existing := range AllowedDestinationWalletTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DestinationWalletType value
func (v DestinationWalletType) Ptr() *DestinationWalletType {
	return &v
}

type NullableDestinationWalletType struct {
	value *DestinationWalletType
	isSet bool
}

func (v NullableDestinationWalletType) Get() *DestinationWalletType {
	return v.value
}

func (v *NullableDestinationWalletType) Set(val *DestinationWalletType) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationWalletType) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationWalletType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationWalletType(val *DestinationWalletType) *NullableDestinationWalletType {
	return &NullableDestinationWalletType{value: val, isSet: true}
}

func (v NullableDestinationWalletType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationWalletType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

