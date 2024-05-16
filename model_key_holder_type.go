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

// KeyHolderType the model 'KeyHolderType'
type KeyHolderType string

// List of KeyHolderType
const (
	KEYHOLDERTYPE_COBO KeyHolderType = "COBO"
	KEYHOLDERTYPE_MOBILE KeyHolderType = "MOBILE"
	KEYHOLDERTYPE_API KeyHolderType = "API"
)

// All allowed values of KeyHolderType enum
var AllowedKeyHolderTypeEnumValues = []KeyHolderType{
	"COBO",
	"MOBILE",
	"API",
}

func (v *KeyHolderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyHolderType(value)
	for _, existing := range AllowedKeyHolderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyHolderType", value)
}

// NewKeyHolderTypeFromValue returns a pointer to a valid KeyHolderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyHolderTypeFromValue(v string) (*KeyHolderType, error) {
	ev := KeyHolderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyHolderType: valid values are %v", v, AllowedKeyHolderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyHolderType) IsValid() bool {
	for _, existing := range AllowedKeyHolderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeyHolderType value
func (v KeyHolderType) Ptr() *KeyHolderType {
	return &v
}

type NullableKeyHolderType struct {
	value *KeyHolderType
	isSet bool
}

func (v NullableKeyHolderType) Get() *KeyHolderType {
	return v.value
}

func (v *NullableKeyHolderType) Set(val *KeyHolderType) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyHolderType) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyHolderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyHolderType(val *KeyHolderType) *NullableKeyHolderType {
	return &NullableKeyHolderType{value: val, isSet: true}
}

func (v NullableKeyHolderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyHolderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

