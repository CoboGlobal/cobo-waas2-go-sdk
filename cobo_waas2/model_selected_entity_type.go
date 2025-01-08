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

// SelectedEntityType Specifies the type of entity. This must be `LEGAL`.
type SelectedEntityType string

// List of SelectedEntityType
const (
	SELECTEDENTITYTYPE_LEGAL SelectedEntityType = "LEGAL"
	SELECTEDENTITYTYPE_NATURAL SelectedEntityType = "NATURAL"
)

// All allowed values of SelectedEntityType enum
var AllowedSelectedEntityTypeEnumValues = []SelectedEntityType{
	"LEGAL",
	"NATURAL",
}

func (v *SelectedEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelectedEntityType(value)
	for _, existing := range AllowedSelectedEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = SelectedEntityType("unknown")
    return nil
}

// NewSelectedEntityTypeFromValue returns a pointer to a valid SelectedEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelectedEntityTypeFromValue(v string) (*SelectedEntityType, error) {
	ev := SelectedEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SelectedEntityType: valid values are %v", v, AllowedSelectedEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelectedEntityType) IsValid() bool {
	for _, existing := range AllowedSelectedEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelectedEntityType value
func (v SelectedEntityType) Ptr() *SelectedEntityType {
	return &v
}

type NullableSelectedEntityType struct {
	value *SelectedEntityType
	isSet bool
}

func (v NullableSelectedEntityType) Get() *SelectedEntityType {
	return v.value
}

func (v *NullableSelectedEntityType) Set(val *SelectedEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectedEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectedEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectedEntityType(val *SelectedEntityType) *NullableSelectedEntityType {
	return &NullableSelectedEntityType{value: val, isSet: true}
}

func (v NullableSelectedEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectedEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

