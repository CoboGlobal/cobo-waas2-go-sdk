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

// AcquiringType The payment acquisition type. - `Order`: Payers pay by fixed-amount orders. Ideal for specific purchases and one-time transactions. - `TopUp`: Account recharge flow where payers deposit funds to their dedicated top-up addresses. Ideal for flexible or usage-based payment models. 
type AcquiringType string

// List of AcquiringType
const (
	ACQUIRINGTYPE_ORDER AcquiringType = "Order"
	ACQUIRINGTYPE_TOP_UP AcquiringType = "TopUp"
)

// All allowed values of AcquiringType enum
var AllowedAcquiringTypeEnumValues = []AcquiringType{
	"Order",
	"TopUp",
}

func (v *AcquiringType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AcquiringType(value)
	for _, existing := range AllowedAcquiringTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = AcquiringType("unknown")
	return nil
}

// NewAcquiringTypeFromValue returns a pointer to a valid AcquiringType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAcquiringTypeFromValue(v string) (*AcquiringType, error) {
	ev := AcquiringType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AcquiringType: valid values are %v", v, AllowedAcquiringTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AcquiringType) IsValid() bool {
	for _, existing := range AllowedAcquiringTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AcquiringType value
func (v AcquiringType) Ptr() *AcquiringType {
	return &v
}

type NullableAcquiringType struct {
	value *AcquiringType
	isSet bool
}

func (v NullableAcquiringType) Get() *AcquiringType {
	return v.value
}

func (v *NullableAcquiringType) Set(val *AcquiringType) {
	v.value = val
	v.isSet = true
}

func (v NullableAcquiringType) IsSet() bool {
	return v.isSet
}

func (v *NullableAcquiringType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcquiringType(val *AcquiringType) *NullableAcquiringType {
	return &NullableAcquiringType{value: val, isSet: true}
}

func (v NullableAcquiringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcquiringType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

