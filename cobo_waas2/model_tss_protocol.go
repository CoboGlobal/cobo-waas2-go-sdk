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

// TSSProtocol The type of TSS protocol. Possible values include:  - 1 : GG18 - 2 : Lindell - 3 : EddsaTSS 
type TSSProtocol int32

// List of TSSProtocol
const (
	TSSPROTOCOL_GG18 TSSProtocol = 1
	TSSPROTOCOL_Lindell TSSProtocol = 2
	TSSPROTOCOL_EddsaTSS TSSProtocol = 3
)

// All allowed values of TSSProtocol enum
var AllowedTSSProtocolEnumValues = []TSSProtocol{
	1,
	2,
	3,
}

func (v *TSSProtocol) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TSSProtocol(value)
	for _, existing := range AllowedTSSProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = TSSProtocol(-1)
	return nil
}

// NewTSSProtocolFromValue returns a pointer to a valid TSSProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTSSProtocolFromValue(v int32) (*TSSProtocol, error) {
	ev := TSSProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TSSProtocol: valid values are %v", v, AllowedTSSProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TSSProtocol) IsValid() bool {
	for _, existing := range AllowedTSSProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TSSProtocol value
func (v TSSProtocol) Ptr() *TSSProtocol {
	return &v
}

type NullableTSSProtocol struct {
	value *TSSProtocol
	isSet bool
}

func (v NullableTSSProtocol) Get() *TSSProtocol {
	return v.value
}

func (v *NullableTSSProtocol) Set(val *TSSProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSProtocol(val *TSSProtocol) *NullableTSSProtocol {
	return &NullableTSSProtocol{value: val, isSet: true}
}

func (v NullableTSSProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

