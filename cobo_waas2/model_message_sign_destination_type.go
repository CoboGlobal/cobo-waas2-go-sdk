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

// MessageSignDestinationType The type of the signature format. Refer to [Transaction sources and destinations](/v2/guides/transactions/sources-and-destinations) for a detailed introduction about the supported sources and destinations for each transaction type.  Each signature format type requires a different set of properties. Switch between the above tabs for details. 
type MessageSignDestinationType string

// List of MessageSignDestinationType
const (
	MESSAGESIGNDESTINATIONTYPE__191_SIGNATURE MessageSignDestinationType = "EVM_EIP_191_Signature"
	MESSAGESIGNDESTINATIONTYPE__712_SIGNATURE MessageSignDestinationType = "EVM_EIP_712_Signature"
)

// All allowed values of MessageSignDestinationType enum
var AllowedMessageSignDestinationTypeEnumValues = []MessageSignDestinationType{
	"EVM_EIP_191_Signature",
	"EVM_EIP_712_Signature",
}

func (v *MessageSignDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageSignDestinationType(value)
	for _, existing := range AllowedMessageSignDestinationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = MessageSignDestinationType("unknown")
    return nil
}

// NewMessageSignDestinationTypeFromValue returns a pointer to a valid MessageSignDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageSignDestinationTypeFromValue(v string) (*MessageSignDestinationType, error) {
	ev := MessageSignDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageSignDestinationType: valid values are %v", v, AllowedMessageSignDestinationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageSignDestinationType) IsValid() bool {
	for _, existing := range AllowedMessageSignDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageSignDestinationType value
func (v MessageSignDestinationType) Ptr() *MessageSignDestinationType {
	return &v
}

type NullableMessageSignDestinationType struct {
	value *MessageSignDestinationType
	isSet bool
}

func (v NullableMessageSignDestinationType) Get() *MessageSignDestinationType {
	return v.value
}

func (v *NullableMessageSignDestinationType) Set(val *MessageSignDestinationType) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSignDestinationType) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSignDestinationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSignDestinationType(val *MessageSignDestinationType) *NullableMessageSignDestinationType {
	return &NullableMessageSignDestinationType{value: val, isSet: true}
}

func (v NullableMessageSignDestinationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSignDestinationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

