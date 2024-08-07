/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// MPCVaultType The vault type. Possible values include: - `Org-Controlled`: This vault is a collection of [Organization-Controlled Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#organization-controlled-wallets).  - `User-Controlled`: This vault is a collection of [User-Controlled Wallets](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#user-controlled-wallets). 
type MPCVaultType string

// List of MPCVaultType
const (
	MPCVAULTTYPE_ORG_CONTROLLED MPCVaultType = "Org-Controlled"
	MPCVAULTTYPE_USER_CONTROLLED MPCVaultType = "User-Controlled"
)

// All allowed values of MPCVaultType enum
var AllowedMPCVaultTypeEnumValues = []MPCVaultType{
	"Org-Controlled",
	"User-Controlled",
}

func (v *MPCVaultType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MPCVaultType(value)
	for _, existing := range AllowedMPCVaultTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = MPCVaultType("unknown")
    return nil
}

// NewMPCVaultTypeFromValue returns a pointer to a valid MPCVaultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMPCVaultTypeFromValue(v string) (*MPCVaultType, error) {
	ev := MPCVaultType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MPCVaultType: valid values are %v", v, AllowedMPCVaultTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MPCVaultType) IsValid() bool {
	for _, existing := range AllowedMPCVaultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MPCVaultType value
func (v MPCVaultType) Ptr() *MPCVaultType {
	return &v
}

type NullableMPCVaultType struct {
	value *MPCVaultType
	isSet bool
}

func (v NullableMPCVaultType) Get() *MPCVaultType {
	return v.value
}

func (v *NullableMPCVaultType) Set(val *MPCVaultType) {
	v.value = val
	v.isSet = true
}

func (v NullableMPCVaultType) IsSet() bool {
	return v.isSet
}

func (v *NullableMPCVaultType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMPCVaultType(val *MPCVaultType) *NullableMPCVaultType {
	return &NullableMPCVaultType{value: val, isSet: true}
}

func (v NullableMPCVaultType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMPCVaultType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

