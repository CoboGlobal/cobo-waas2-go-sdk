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

// MessageSignSourceType The wallet subtype of the transaction source. Refer to [Transaction sources and destinations](https://www.cobo.com/developers/v2/guides/transactions/sources-and-destinations) for a detailed introduction about the supported sources and destinations for each transaction type. 
type MessageSignSourceType string

// List of MessageSignSourceType
const (
	MESSAGESIGNSOURCETYPE_WEB3 MessageSignSourceType = "Web3"
	MESSAGESIGNSOURCETYPE_ORG_CONTROLLED MessageSignSourceType = "Org-Controlled"
	MESSAGESIGNSOURCETYPE_USER_CONTROLLED MessageSignSourceType = "User-Controlled"
)

// All allowed values of MessageSignSourceType enum
var AllowedMessageSignSourceTypeEnumValues = []MessageSignSourceType{
	"Web3",
	"Org-Controlled",
	"User-Controlled",
}

func (v *MessageSignSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageSignSourceType(value)
	for _, existing := range AllowedMessageSignSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = MessageSignSourceType("unknown")
	return nil
}

// NewMessageSignSourceTypeFromValue returns a pointer to a valid MessageSignSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageSignSourceTypeFromValue(v string) (*MessageSignSourceType, error) {
	ev := MessageSignSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageSignSourceType: valid values are %v", v, AllowedMessageSignSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageSignSourceType) IsValid() bool {
	for _, existing := range AllowedMessageSignSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageSignSourceType value
func (v MessageSignSourceType) Ptr() *MessageSignSourceType {
	return &v
}

type NullableMessageSignSourceType struct {
	value *MessageSignSourceType
	isSet bool
}

func (v NullableMessageSignSourceType) Get() *MessageSignSourceType {
	return v.value
}

func (v *NullableMessageSignSourceType) Set(val *MessageSignSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSignSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSignSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSignSourceType(val *MessageSignSourceType) *NullableMessageSignSourceType {
	return &NullableMessageSignSourceType{value: val, isSet: true}
}

func (v NullableMessageSignSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSignSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

