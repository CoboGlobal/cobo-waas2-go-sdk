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

// TokenizationOperationSourceType The wallet source type. Possible values include: - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets). - `Web3`: Custodial Wallets (Web3 Wallets).  
type TokenizationOperationSourceType string

// List of TokenizationOperationSourceType
const (
	TOKENIZATIONOPERATIONSOURCETYPE_ORG_CONTROLLED TokenizationOperationSourceType = "Org-Controlled"
	TOKENIZATIONOPERATIONSOURCETYPE_WEB3 TokenizationOperationSourceType = "Web3"
)

// All allowed values of TokenizationOperationSourceType enum
var AllowedTokenizationOperationSourceTypeEnumValues = []TokenizationOperationSourceType{
	"Org-Controlled",
	"Web3",
}

func (v *TokenizationOperationSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TokenizationOperationSourceType(value)
	for _, existing := range AllowedTokenizationOperationSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = TokenizationOperationSourceType("unknown")
	return nil
}

// NewTokenizationOperationSourceTypeFromValue returns a pointer to a valid TokenizationOperationSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTokenizationOperationSourceTypeFromValue(v string) (*TokenizationOperationSourceType, error) {
	ev := TokenizationOperationSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TokenizationOperationSourceType: valid values are %v", v, AllowedTokenizationOperationSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TokenizationOperationSourceType) IsValid() bool {
	for _, existing := range AllowedTokenizationOperationSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TokenizationOperationSourceType value
func (v TokenizationOperationSourceType) Ptr() *TokenizationOperationSourceType {
	return &v
}

type NullableTokenizationOperationSourceType struct {
	value *TokenizationOperationSourceType
	isSet bool
}

func (v NullableTokenizationOperationSourceType) Get() *TokenizationOperationSourceType {
	return v.value
}

func (v *NullableTokenizationOperationSourceType) Set(val *TokenizationOperationSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizationOperationSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizationOperationSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizationOperationSourceType(val *TokenizationOperationSourceType) *NullableTokenizationOperationSourceType {
	return &NullableTokenizationOperationSourceType{value: val, isSet: true}
}

func (v NullableTokenizationOperationSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizationOperationSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

