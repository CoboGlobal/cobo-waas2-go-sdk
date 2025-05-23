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

// TSSCurve The elliptic curve type. Possible values include:  - 0 : The `SECP256K1` elliptic curve. - 2 : The `ED25519` elliptic curve. 
type TSSCurve int32

// List of TSSCurve
const (
	TSSCURVE_SECP256K1 TSSCurve = 0
	TSSCURVE_ED25519 TSSCurve = 2
)

// All allowed values of TSSCurve enum
var AllowedTSSCurveEnumValues = []TSSCurve{
	0,
	2,
}

func (v *TSSCurve) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TSSCurve(value)
	for _, existing := range AllowedTSSCurveEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = TSSCurve(-1)
	return nil
}

// NewTSSCurveFromValue returns a pointer to a valid TSSCurve
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTSSCurveFromValue(v int32) (*TSSCurve, error) {
	ev := TSSCurve(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TSSCurve: valid values are %v", v, AllowedTSSCurveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TSSCurve) IsValid() bool {
	for _, existing := range AllowedTSSCurveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TSSCurve value
func (v TSSCurve) Ptr() *TSSCurve {
	return &v
}

type NullableTSSCurve struct {
	value *TSSCurve
	isSet bool
}

func (v NullableTSSCurve) Get() *TSSCurve {
	return v.value
}

func (v *NullableTSSCurve) Set(val *TSSCurve) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSCurve) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSCurve) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSCurve(val *TSSCurve) *NullableTSSCurve {
	return &NullableTSSCurve{value: val, isSet: true}
}

func (v NullableTSSCurve) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSCurve) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

