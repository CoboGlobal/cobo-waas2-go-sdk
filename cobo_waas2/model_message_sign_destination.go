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

// MessageSignDestination - struct for MessageSignDestination
type MessageSignDestination struct {
	EvmEIP191MessageSignDestination *EvmEIP191MessageSignDestination
	EvmEIP712MessageSignDestination *EvmEIP712MessageSignDestination
	RawMessageSignDestination *RawMessageSignDestination
}

// EvmEIP191MessageSignDestinationAsMessageSignDestination is a convenience function that returns EvmEIP191MessageSignDestination wrapped in MessageSignDestination
func EvmEIP191MessageSignDestinationAsMessageSignDestination(v *EvmEIP191MessageSignDestination) MessageSignDestination {
	return MessageSignDestination{
		EvmEIP191MessageSignDestination: v,
	}
}

// EvmEIP712MessageSignDestinationAsMessageSignDestination is a convenience function that returns EvmEIP712MessageSignDestination wrapped in MessageSignDestination
func EvmEIP712MessageSignDestinationAsMessageSignDestination(v *EvmEIP712MessageSignDestination) MessageSignDestination {
	return MessageSignDestination{
		EvmEIP712MessageSignDestination: v,
	}
}

// RawMessageSignDestinationAsMessageSignDestination is a convenience function that returns RawMessageSignDestination wrapped in MessageSignDestination
func RawMessageSignDestinationAsMessageSignDestination(v *RawMessageSignDestination) MessageSignDestination {
	return MessageSignDestination{
		RawMessageSignDestination: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MessageSignDestination) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EVM_EIP_191'
	if jsonDict["destination_type"] == "EVM_EIP_191" {
		// try to unmarshal JSON data into EvmEIP191MessageSignDestination
		err = json.Unmarshal(data, &dst.EvmEIP191MessageSignDestination)
		if err == nil {
			return nil // data stored in dst.EvmEIP191MessageSignDestination, return on the first match
		} else {
			dst.EvmEIP191MessageSignDestination = nil
			return fmt.Errorf("failed to unmarshal MessageSignDestination as EvmEIP191MessageSignDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EVM_EIP_712'
	if jsonDict["destination_type"] == "EVM_EIP_712" {
		// try to unmarshal JSON data into EvmEIP712MessageSignDestination
		err = json.Unmarshal(data, &dst.EvmEIP712MessageSignDestination)
		if err == nil {
			return nil // data stored in dst.EvmEIP712MessageSignDestination, return on the first match
		} else {
			dst.EvmEIP712MessageSignDestination = nil
			return fmt.Errorf("failed to unmarshal MessageSignDestination as EvmEIP712MessageSignDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RAW_MESSAGE'
	if jsonDict["destination_type"] == "RAW_MESSAGE" {
		// try to unmarshal JSON data into RawMessageSignDestination
		err = json.Unmarshal(data, &dst.RawMessageSignDestination)
		if err == nil {
			return nil // data stored in dst.RawMessageSignDestination, return on the first match
		} else {
			dst.RawMessageSignDestination = nil
			return fmt.Errorf("failed to unmarshal MessageSignDestination as RawMessageSignDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EvmEIP191MessageSignDestination'
	if jsonDict["destination_type"] == "EvmEIP191MessageSignDestination" {
		// try to unmarshal JSON data into EvmEIP191MessageSignDestination
		err = json.Unmarshal(data, &dst.EvmEIP191MessageSignDestination)
		if err == nil {
			return nil // data stored in dst.EvmEIP191MessageSignDestination, return on the first match
		} else {
			dst.EvmEIP191MessageSignDestination = nil
			return fmt.Errorf("failed to unmarshal MessageSignDestination as EvmEIP191MessageSignDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EvmEIP712MessageSignDestination'
	if jsonDict["destination_type"] == "EvmEIP712MessageSignDestination" {
		// try to unmarshal JSON data into EvmEIP712MessageSignDestination
		err = json.Unmarshal(data, &dst.EvmEIP712MessageSignDestination)
		if err == nil {
			return nil // data stored in dst.EvmEIP712MessageSignDestination, return on the first match
		} else {
			dst.EvmEIP712MessageSignDestination = nil
			return fmt.Errorf("failed to unmarshal MessageSignDestination as EvmEIP712MessageSignDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RawMessageSignDestination'
	if jsonDict["destination_type"] == "RawMessageSignDestination" {
		// try to unmarshal JSON data into RawMessageSignDestination
		err = json.Unmarshal(data, &dst.RawMessageSignDestination)
		if err == nil {
			return nil // data stored in dst.RawMessageSignDestination, return on the first match
		} else {
			dst.RawMessageSignDestination = nil
			return fmt.Errorf("failed to unmarshal MessageSignDestination as RawMessageSignDestination: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageSignDestination) MarshalJSON() ([]byte, error) {
	if src.EvmEIP191MessageSignDestination != nil {
		return json.Marshal(&src.EvmEIP191MessageSignDestination)
	}

	if src.EvmEIP712MessageSignDestination != nil {
		return json.Marshal(&src.EvmEIP712MessageSignDestination)
	}

	if src.RawMessageSignDestination != nil {
		return json.Marshal(&src.RawMessageSignDestination)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MessageSignDestination) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EvmEIP191MessageSignDestination != nil {
		return obj.EvmEIP191MessageSignDestination
	}

	if obj.EvmEIP712MessageSignDestination != nil {
		return obj.EvmEIP712MessageSignDestination
	}

	if obj.RawMessageSignDestination != nil {
		return obj.RawMessageSignDestination
	}

	// all schemas are nil
	return nil
}

type NullableMessageSignDestination struct {
	value *MessageSignDestination
	isSet bool
}

func (v NullableMessageSignDestination) Get() *MessageSignDestination {
	return v.value
}

func (v *NullableMessageSignDestination) Set(val *MessageSignDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSignDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSignDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSignDestination(val *MessageSignDestination) *NullableMessageSignDestination {
	return &NullableMessageSignDestination{value: val, isSet: true}
}

func (v NullableMessageSignDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSignDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


