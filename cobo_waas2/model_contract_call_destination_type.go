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

// ContractCallDestinationType The type of the contract format. Refer to [Transaction sources and destinations](https://www.cobo.com/developers/v2/guides/transactions/sources-and-destinations) for a detailed introduction about the supported sources and destinations for each transaction type. 
type ContractCallDestinationType string

// List of ContractCallDestinationType
const (
	CONTRACTCALLDESTINATIONTYPE_EVM_CONTRACT ContractCallDestinationType = "EVM_Contract"
	CONTRACTCALLDESTINATIONTYPE_SOL_CONTRACT ContractCallDestinationType = "SOL_Contract"
)

// All allowed values of ContractCallDestinationType enum
var AllowedContractCallDestinationTypeEnumValues = []ContractCallDestinationType{
	"EVM_Contract",
	"SOL_Contract",
}

func (v *ContractCallDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContractCallDestinationType(value)
	for _, existing := range AllowedContractCallDestinationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = ContractCallDestinationType("unknown")
	return nil
}

// NewContractCallDestinationTypeFromValue returns a pointer to a valid ContractCallDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContractCallDestinationTypeFromValue(v string) (*ContractCallDestinationType, error) {
	ev := ContractCallDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContractCallDestinationType: valid values are %v", v, AllowedContractCallDestinationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContractCallDestinationType) IsValid() bool {
	for _, existing := range AllowedContractCallDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContractCallDestinationType value
func (v ContractCallDestinationType) Ptr() *ContractCallDestinationType {
	return &v
}

type NullableContractCallDestinationType struct {
	value *ContractCallDestinationType
	isSet bool
}

func (v NullableContractCallDestinationType) Get() *ContractCallDestinationType {
	return v.value
}

func (v *NullableContractCallDestinationType) Set(val *ContractCallDestinationType) {
	v.value = val
	v.isSet = true
}

func (v NullableContractCallDestinationType) IsSet() bool {
	return v.isSet
}

func (v *NullableContractCallDestinationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractCallDestinationType(val *ContractCallDestinationType) *NullableContractCallDestinationType {
	return &NullableContractCallDestinationType{value: val, isSet: true}
}

func (v NullableContractCallDestinationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractCallDestinationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

