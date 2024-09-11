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

// ContractCallSourceType The type of the source. Refer to [Transaction sources and destinations](/v2/guides/sources-and-destinations) for details. 
type ContractCallSourceType string

// List of ContractCallSourceType
const (
	CONTRACTCALLSOURCETYPE_ORG_CONTROLLED ContractCallSourceType = "Org-Controlled"
	CONTRACTCALLSOURCETYPE_USER_CONTROLLED ContractCallSourceType = "User-Controlled"
	CONTRACTCALLSOURCETYPE_SAFEWALLET ContractCallSourceType = "Safe{Wallet}"
)

// All allowed values of ContractCallSourceType enum
var AllowedContractCallSourceTypeEnumValues = []ContractCallSourceType{
	"Org-Controlled",
	"User-Controlled",
	"Safe{Wallet}",
}

func (v *ContractCallSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContractCallSourceType(value)
	for _, existing := range AllowedContractCallSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = ContractCallSourceType("unknown")
    return nil
}

// NewContractCallSourceTypeFromValue returns a pointer to a valid ContractCallSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContractCallSourceTypeFromValue(v string) (*ContractCallSourceType, error) {
	ev := ContractCallSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContractCallSourceType: valid values are %v", v, AllowedContractCallSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContractCallSourceType) IsValid() bool {
	for _, existing := range AllowedContractCallSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContractCallSourceType value
func (v ContractCallSourceType) Ptr() *ContractCallSourceType {
	return &v
}

type NullableContractCallSourceType struct {
	value *ContractCallSourceType
	isSet bool
}

func (v NullableContractCallSourceType) Get() *ContractCallSourceType {
	return v.value
}

func (v *NullableContractCallSourceType) Set(val *ContractCallSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableContractCallSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableContractCallSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractCallSourceType(val *ContractCallSourceType) *NullableContractCallSourceType {
	return &NullableContractCallSourceType{value: val, isSet: true}
}

func (v NullableContractCallSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractCallSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

